package archive

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

// VerifyOptions controls which snapshot metadata differences are blocking.
type VerifyOptions struct {
	StrictBlockMetadata bool
	StrictTxNumLogs     bool
}

// MismatchStats contains the independently measured archive mismatch
// categories for one project.
type MismatchStats struct {
	EventLogsMissingFromArchive int64
	ArchiveEventLogOrphans      int64
	EventLogDataMismatches      int64
	ProjectTransactionsMissing  int64
	ArchiveLogsMissingTx        int64
	TxCoreMismatches            int64
	TxNumLogsOnlyMismatches     int64
	ProjectBlocksMissing        int64
	BlockHashMismatches         int64
	BlockMetadataMismatches     int64
}

// Passed applies strictness policy to independently loaded mismatch counts.
func (s MismatchStats) Passed(options VerifyOptions) bool {
	if s.EventLogsMissingFromArchive != 0 ||
		s.ArchiveEventLogOrphans != 0 ||
		s.EventLogDataMismatches != 0 ||
		s.ProjectTransactionsMissing != 0 ||
		s.ArchiveLogsMissingTx != 0 ||
		s.TxCoreMismatches != 0 ||
		s.ProjectBlocksMissing != 0 ||
		s.BlockHashMismatches != 0 {
		return false
	}
	if options.StrictTxNumLogs && s.TxNumLogsOnlyMismatches != 0 {
		return false
	}
	if options.StrictBlockMetadata && s.BlockMetadataMismatches != 0 {
		return false
	}
	return true
}

// ProjectVerification is one project's archive verification result.
type ProjectVerification struct {
	Project string
	Stats   MismatchStats
	Passed  bool
}

// VerificationReport is the ordered result for all selected projects.
type VerificationReport struct {
	Projects []ProjectVerification
	Passed   bool
}

// ProjectVerifier is the narrow orchestration seam used by unit tests.
type ProjectVerifier interface {
	VerifyProject(ctx context.Context, project string, options VerifyOptions) (ProjectVerification, error)
}

// VerifyProjects verifies projects in order and stops on cancellation or a
// query error. Mismatches are represented in the returned report, not as an
// operational error.
func VerifyProjects(
	ctx context.Context,
	projects []string,
	options VerifyOptions,
	verifier ProjectVerifier,
) (VerificationReport, error) {
	report := VerificationReport{
		Projects: make([]ProjectVerification, 0, len(projects)),
		Passed:   true,
	}
	if verifier == nil {
		return report, errors.New("archive verify: verifier is required")
	}
	for _, project := range projects {
		if err := ctx.Err(); err != nil {
			return report, err
		}
		result, err := verifier.VerifyProject(ctx, project, options)
		if err != nil {
			return report, err
		}
		report.Projects = append(report.Projects, result)
		report.Passed = report.Passed && result.Passed
	}
	return report, nil
}

// SQLVerifier runs archive consistency queries against one database. The
// caller owns DB.
type SQLVerifier struct {
	DB     *sql.DB
	Logger Logger
}

// VerifyProject loads and evaluates every mismatch category for project.
func (v *SQLVerifier) VerifyProject(
	ctx context.Context,
	project string,
	options VerifyOptions,
) (ProjectVerification, error) {
	result := ProjectVerification{Project: project}
	if v == nil || v.DB == nil {
		return result, errors.New("archive verify: database is required")
	}
	v.println("")
	v.printf("=== Verifying project: %s ===", project)

	contracts, err := LoadProjectContracts(ctx, v.DB, project)
	if err != nil {
		return result, err
	}
	v.printf("Contract AIDs (%d): %v", len(contracts.AddressIDs), contracts.AddressIDs)

	stats := MismatchStats{}
	stats.EventLogsMissingFromArchive, err = countRow(ctx, v.DB, `
		SELECT COUNT(*) FROM evt_log e
		INNER JOIN transaction t ON e.tx_id = t.id
		WHERE e.contract_aid = ANY($1)
		AND NOT EXISTS (
			SELECT 1 FROM arch_evtlog ae
			WHERE ae.tx_hash = t.tx_hash AND ae.log_index = e.log_index
		)
	`, pq.Array(contracts.AddressIDs))
	if err != nil {
		return result, err
	}
	v.printf("evt_log rows missing from arch_evtlog (by tx_hash + log_index): %d", stats.EventLogsMissingFromArchive)

	stats.ArchiveEventLogOrphans, err = countRow(ctx, v.DB, `
		SELECT COUNT(*) FROM arch_evtlog ae
		WHERE ae.contract_addr = ANY($1)
		AND NOT EXISTS (
			SELECT 1 FROM evt_log e
			INNER JOIN transaction t ON e.tx_id = t.id
			WHERE t.tx_hash = ae.tx_hash AND e.log_index = ae.log_index
		)
	`, pq.Array(contracts.Addresses))
	if err != nil {
		return result, err
	}
	v.printf("arch_evtlog orphan rows (no matching live log for tx_hash + log_index): %d", stats.ArchiveEventLogOrphans)

	stats.EventLogDataMismatches, err = countRow(ctx, v.DB, `
		SELECT COUNT(*) FROM arch_evtlog ae
		INNER JOIN transaction t ON t.tx_hash = ae.tx_hash
		INNER JOIN evt_log e ON e.tx_id = t.id AND e.log_index = ae.log_index
		INNER JOIN address a ON e.contract_aid = a.address_id
		WHERE ae.contract_addr = ANY($1)
		AND (
			ae.block_num IS DISTINCT FROM e.block_num
			OR ae.tx_hash IS DISTINCT FROM t.tx_hash
			OR ae.contract_addr IS DISTINCT FROM a.addr
			OR ae.topic0_sig IS DISTINCT FROM e.topic0_sig
			OR ae.log_rlp IS DISTINCT FROM e.log_rlp
			OR (ae.evt_id IS NOT NULL AND ae.evt_id IS DISTINCT FROM e.id)
		)
	`, pq.Array(contracts.Addresses))
	if err != nil {
		return result, err
	}
	v.printf("arch_evtlog rows that disagree with live evt_log/tx/address: %d", stats.EventLogDataMismatches)

	stats.ProjectTransactionsMissing, err = countRow(ctx, v.DB, `
		SELECT COUNT(DISTINCT t.tx_hash) FROM evt_log e
		INNER JOIN transaction t ON e.tx_id = t.id
		WHERE e.contract_aid = ANY($1)
		AND NOT EXISTS (SELECT 1 FROM arch_tx at WHERE at.tx_hash = t.tx_hash)
	`, pq.Array(contracts.AddressIDs))
	if err != nil {
		return result, err
	}
	v.printf("Distinct tx_hash from project evt_log missing in arch_tx: %d", stats.ProjectTransactionsMissing)

	stats.ArchiveLogsMissingTx, err = countRow(ctx, v.DB, `
		SELECT COUNT(*) FROM arch_evtlog ae
		WHERE ae.contract_addr = ANY($1)
		AND NOT EXISTS (SELECT 1 FROM arch_tx at WHERE at.tx_hash = ae.tx_hash)
	`, pq.Array(contracts.Addresses))
	if err != nil {
		return result, err
	}
	v.printf("arch_evtlog rows whose tx_hash is missing from arch_tx: %d", stats.ArchiveLogsMissingTx)

	stats.TxCoreMismatches, stats.TxNumLogsOnlyMismatches, err = txMismatchStats(ctx, v.DB, contracts.AddressIDs)
	if err != nil {
		return result, err
	}
	txFailing := stats.TxCoreMismatches
	if options.StrictTxNumLogs {
		txFailing += stats.TxNumLogsOnlyMismatches
	}
	v.printf("arch_tx vs transaction (project tx_hashes): failing rows: %d", txFailing)
	if !options.StrictTxNumLogs && stats.TxNumLogsOnlyMismatches > 0 {
		v.printf("  … rows that match on core fields but differ only on num_logs: %d (not failing; pass --strict-arch-tx-num-logs to count as failure)", stats.TxNumLogsOnlyMismatches)
	}

	stats.ProjectBlocksMissing, err = countRow(ctx, v.DB, `
		SELECT COUNT(DISTINCT at.block_num)
		FROM arch_tx at
		WHERE at.tx_hash IN (
			SELECT DISTINCT ae.tx_hash FROM arch_evtlog ae WHERE ae.contract_addr = ANY($1)
		)
		AND NOT EXISTS (SELECT 1 FROM arch_block ab WHERE ab.block_num = at.block_num)
	`, pq.Array(contracts.Addresses))
	if err != nil {
		return result, err
	}
	v.printf("block_num values (via arch_tx for project logs) missing from arch_block: %d", stats.ProjectBlocksMissing)

	stats.BlockHashMismatches, stats.BlockMetadataMismatches, err = blockMismatchStats(ctx, v.DB, contracts.Addresses)
	if err != nil {
		return result, err
	}
	v.printf("arch_block vs live block (same block_num): hash/parent mismatch: %d", stats.BlockHashMismatches)
	v.printf("arch_block vs live block (same block_num): metadata mismatch (num_tx, ts, cash_flow): %d", stats.BlockMetadataMismatches)
	if stats.BlockMetadataMismatches > 0 && !options.StrictBlockMetadata {
		v.println("  … metadata drift not failing (arch is snapshot; live block may be refreshed — pass --strict-arch-block-metadata to fail on this)")
	}

	result.Stats = stats
	result.Passed = stats.Passed(options)
	if result.Passed {
		v.printf("--- project %s: PASS (under selected strictness) ---", project)
	} else {
		v.printf("--- project %s: FAIL ---", project)
	}
	return result, nil
}

func txMismatchStats(ctx context.Context, db *sql.DB, addressIDs []int64) (core, numLogsOnly int64, err error) {
	base := `
		FROM (
			SELECT DISTINCT t.tx_hash AS tx_hash
			FROM evt_log e
			INNER JOIN transaction t ON e.tx_id = t.id
			WHERE e.contract_aid = ANY($1)
		) x
		INNER JOIN transaction t ON t.tx_hash = x.tx_hash
		INNER JOIN arch_tx at ON at.tx_hash = t.tx_hash
	`
	whereStrict := `
		at.block_num IS DISTINCT FROM t.block_num
			OR at.from_aid IS DISTINCT FROM t.from_aid
			OR at.to_aid IS DISTINCT FROM t.to_aid
			OR at.gas_used IS DISTINCT FROM t.gas_used
			OR at.tx_index IS DISTINCT FROM t.tx_index
			OR at.num_logs IS DISTINCT FROM t.num_logs
			OR at.ctrct_create IS DISTINCT FROM t.ctrct_create
			OR at.input_sig IS DISTINCT FROM t.input_sig
			OR at.value IS DISTINCT FROM t.value
			OR at.gas_price IS DISTINCT FROM t.gas_price`
	whereCore := `
		at.block_num IS DISTINCT FROM t.block_num
			OR at.from_aid IS DISTINCT FROM t.from_aid
			OR at.to_aid IS DISTINCT FROM t.to_aid
			OR at.gas_used IS DISTINCT FROM t.gas_used
			OR at.tx_index IS DISTINCT FROM t.tx_index
			OR at.ctrct_create IS DISTINCT FROM t.ctrct_create
			OR at.input_sig IS DISTINCT FROM t.input_sig
			OR at.value IS DISTINCT FROM t.value
			OR at.gas_price IS DISTINCT FROM t.gas_price`

	var strict int64
	if err := db.QueryRowContext(ctx,
		`SELECT COUNT(*) `+base+` WHERE `+whereStrict,
		pq.Array(addressIDs),
	).Scan(&strict); err != nil {
		return 0, 0, fmt.Errorf("tx mismatch (strict): %w", err)
	}
	if err := db.QueryRowContext(ctx,
		`SELECT COUNT(*) `+base+` WHERE `+whereCore,
		pq.Array(addressIDs),
	).Scan(&core); err != nil {
		return 0, 0, fmt.Errorf("tx mismatch (core): %w", err)
	}
	numLogsOnly = strict - core
	if numLogsOnly < 0 {
		numLogsOnly = 0
	}
	return core, numLogsOnly, nil
}

func blockMismatchStats(ctx context.Context, db *sql.DB, addresses []string) (hash, metadata int64, err error) {
	subquery := `
		SELECT DISTINCT at.block_num FROM arch_tx at
		WHERE at.tx_hash IN (
			SELECT DISTINCT ae.tx_hash FROM arch_evtlog ae WHERE ae.contract_addr = ANY($1)
		)`
	hashQuery := `
		SELECT COUNT(*) FROM arch_block ab
		INNER JOIN block b ON b.block_num = ab.block_num
		WHERE ab.block_num IN (` + subquery + `)
		AND (
			ab.block_hash IS DISTINCT FROM b.block_hash
			OR ab.parent_hash IS DISTINCT FROM b.parent_hash
		)`
	metadataQuery := `
		SELECT COUNT(*) FROM arch_block ab
		INNER JOIN block b ON b.block_num = ab.block_num
		WHERE ab.block_num IN (` + subquery + `)
		AND (
			ab.num_tx IS DISTINCT FROM b.num_tx
			OR ab.ts IS DISTINCT FROM b.ts
			OR ab.cash_flow IS DISTINCT FROM b.cash_flow
		)
		AND ab.block_hash IS NOT DISTINCT FROM b.block_hash
		AND ab.parent_hash IS NOT DISTINCT FROM b.parent_hash`

	if err := db.QueryRowContext(ctx, hashQuery, pq.Array(addresses)).Scan(&hash); err != nil {
		return 0, 0, fmt.Errorf("block hash mismatch: %w", err)
	}
	if err := db.QueryRowContext(ctx, metadataQuery, pq.Array(addresses)).Scan(&metadata); err != nil {
		return 0, 0, fmt.Errorf("block meta mismatch: %w", err)
	}
	return hash, metadata, nil
}

func countRow(ctx context.Context, db *sql.DB, query string, args ...any) (int64, error) {
	var count int64
	if err := db.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("count query failed: %w\n%s", err, query)
	}
	return count, nil
}

func (v *SQLVerifier) printf(format string, args ...any) {
	if v.Logger != nil {
		v.Logger.Printf(format, args...)
	}
}

func (v *SQLVerifier) println(args ...any) {
	if v.Logger != nil {
		v.Logger.Println(args...)
	}
}
