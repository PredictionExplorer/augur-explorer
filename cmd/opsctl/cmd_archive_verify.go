package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
	"github.com/spf13/cobra"
)

// newArchiveVerifyCmd builds `opsctl archive verify`, the replacement for the
// standalone arch_verify tool.
func newArchiveVerifyCmd() *cobra.Command {
	var (
		dbConn          string
		projectType     string
		strictBlockMeta bool
		strictTxNumLogs bool
	)
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Check live evt_log / transaction / block against the arch_* tables",
		Long: `Runs archival consistency checks between the live tables and the arch_*
tables in the same database, per project. Exits non-zero when a blocking
mismatch is found; metadata drift is a warning unless a --strict-* flag is set.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			projects, err := resolveProjects(projectType)
			if err != nil {
				return err
			}

			db, err := sql.Open("postgres", dbConn)
			if err != nil {
				return fmt.Errorf("connect: %w", err)
			}
			defer db.Close()
			db.SetMaxOpenConns(2)

			allOK := true
			for _, p := range projects {
				ok, err := verifyArchiveProject(db, p, strictBlockMeta, strictTxNumLogs)
				if err != nil {
					return err
				}
				allOK = allOK && ok
			}

			log.Println("")
			log.Println("=== SUMMARY ===")
			if allOK {
				log.Println("OK — no blocking mismatches for selected project(s). Review any warnings above.")
				return nil
			}
			log.Println("FAILED — see details above.")
			return errors.New("archive verification failed")
		},
	}
	cmd.Flags().StringVar(&dbConn, "db", "", "PostgreSQL connection string (same DB holds live + arch_* tables)")
	cmd.Flags().StringVar(&projectType, "project", "", "Project: randomwalk | cosmicgame | both (same order as archive export: cosmicgame then randomwalk)")
	cmd.Flags().BoolVar(&strictBlockMeta, "strict-arch-block-metadata", false,
		"If set, require arch_block num_tx, ts, cash_flow to match live block (default: only block_hash and parent_hash must match).")
	cmd.Flags().BoolVar(&strictTxNumLogs, "strict-arch-tx-num-logs", false,
		"If set, require arch_tx.num_logs to match transaction.num_logs (default: ignore num_logs drift; indexer may refresh it after archival).")
	_ = cmd.MarkFlagRequired("db")
	_ = cmd.MarkFlagRequired("project")
	return cmd
}

func init() { archiveCmd.AddCommand(newArchiveVerifyCmd()) }

// verifyArchiveProject runs every archive consistency check for one project
// and reports whether the project passed under the selected strictness.
func verifyArchiveProject(db *sql.DB, project string, strictBlockMeta, strictTxNumLogs bool) (bool, error) {
	log.Printf("")
	log.Printf("=== Verifying project: %s ===", project)

	aids, addrs, err := projectContracts(db, project)
	if err != nil {
		return false, err
	}
	log.Printf("Contract AIDs (%d): %v", len(aids), aids)

	ok := true

	c, err := countRow(db, `
		SELECT COUNT(*) FROM evt_log e
		INNER JOIN transaction t ON e.tx_id = t.id
		WHERE e.contract_aid = ANY($1)
		AND NOT EXISTS (
			SELECT 1 FROM arch_evtlog ae
			WHERE ae.tx_hash = t.tx_hash AND ae.log_index = e.log_index
		)
	`, pq.Array(aids))
	if err != nil {
		return false, err
	}
	log.Printf("evt_log rows missing from arch_evtlog (by tx_hash + log_index): %d", c)
	if c != 0 {
		ok = false
	}

	c, err = countRow(db, `
		SELECT COUNT(*) FROM arch_evtlog ae
		WHERE ae.contract_addr = ANY($1)
		AND NOT EXISTS (
			SELECT 1 FROM evt_log e
			INNER JOIN transaction t ON e.tx_id = t.id
			WHERE t.tx_hash = ae.tx_hash AND e.log_index = ae.log_index
		)
	`, pq.Array(addrs))
	if err != nil {
		return false, err
	}
	log.Printf("arch_evtlog orphan rows (no matching live log for tx_hash + log_index): %d", c)
	if c != 0 {
		ok = false
	}

	c, err = countRow(db, `
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
	`, pq.Array(addrs))
	if err != nil {
		return false, err
	}
	log.Printf("arch_evtlog rows that disagree with live evt_log/tx/address: %d", c)
	if c != 0 {
		ok = false
	}

	c, err = countRow(db, `
		SELECT COUNT(DISTINCT t.tx_hash) FROM evt_log e
		INNER JOIN transaction t ON e.tx_id = t.id
		WHERE e.contract_aid = ANY($1)
		AND NOT EXISTS (SELECT 1 FROM arch_tx at WHERE at.tx_hash = t.tx_hash)
	`, pq.Array(aids))
	if err != nil {
		return false, err
	}
	log.Printf("Distinct tx_hash from project evt_log missing in arch_tx: %d", c)
	if c != 0 {
		ok = false
	}

	c, err = countRow(db, `
		SELECT COUNT(*) FROM arch_evtlog ae
		WHERE ae.contract_addr = ANY($1)
		AND NOT EXISTS (SELECT 1 FROM arch_tx at WHERE at.tx_hash = ae.tx_hash)
	`, pq.Array(addrs))
	if err != nil {
		return false, err
	}
	log.Printf("arch_evtlog rows whose tx_hash is missing from arch_tx: %d", c)
	if c != 0 {
		ok = false
	}

	txBad, txNumLogsSkew, err := txMismatchStats(db, aids, strictTxNumLogs)
	if err != nil {
		return false, err
	}
	log.Printf("arch_tx vs transaction (project tx_hashes): failing rows: %d", txBad)
	if !strictTxNumLogs && txNumLogsSkew > 0 {
		log.Printf("  … rows that match on core fields but differ only on num_logs: %d (not failing; pass --strict-arch-tx-num-logs to count as failure)", txNumLogsSkew)
	}
	if txBad != 0 {
		ok = false
	}

	c, err = countRow(db, `
		SELECT COUNT(DISTINCT at.block_num)
		FROM arch_tx at
		WHERE at.tx_hash IN (
			SELECT DISTINCT ae.tx_hash FROM arch_evtlog ae WHERE ae.contract_addr = ANY($1)
		)
		AND NOT EXISTS (SELECT 1 FROM arch_block ab WHERE ab.block_num = at.block_num)
	`, pq.Array(addrs))
	if err != nil {
		return false, err
	}
	log.Printf("block_num values (via arch_tx for project logs) missing from arch_block: %d", c)
	if c != 0 {
		ok = false
	}

	hashBad, metaBad, err := blockMismatchStats(db, addrs)
	if err != nil {
		return false, err
	}
	log.Printf("arch_block vs live block (same block_num): hash/parent mismatch: %d", hashBad)
	log.Printf("arch_block vs live block (same block_num): metadata mismatch (num_tx, ts, cash_flow): %d", metaBad)
	if metaBad > 0 && !strictBlockMeta {
		log.Printf("  … metadata drift not failing (arch is snapshot; live block may be refreshed — pass --strict-arch-block-metadata to fail on this)")
	}
	if hashBad != 0 {
		ok = false
	}
	if strictBlockMeta && metaBad != 0 {
		ok = false
	}

	if ok {
		log.Printf("--- project %s: PASS (under selected strictness) ---", project)
	} else {
		log.Printf("--- project %s: FAIL ---", project)
	}
	return ok, nil
}

// txMismatchStats returns the failing arch_tx row count under the chosen
// rules and, when strictTxNumLogs is false, the number of rows that differ on
// num_logs only.
func txMismatchStats(db *sql.DB, aids []int64, strictTxNumLogs bool) (failCount, numLogsSkew int64, err error) {
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

	var nStrict, nCore int64
	if err := db.QueryRow(`SELECT COUNT(*) `+base+` WHERE `+whereStrict, pq.Array(aids)).Scan(&nStrict); err != nil {
		return 0, 0, fmt.Errorf("tx mismatch (strict): %w", err)
	}
	if err := db.QueryRow(`SELECT COUNT(*) `+base+` WHERE `+whereCore, pq.Array(aids)).Scan(&nCore); err != nil {
		return 0, 0, fmt.Errorf("tx mismatch (core): %w", err)
	}
	if strictTxNumLogs {
		return nStrict, 0, nil
	}
	// Rows failing the full compare but matching on everything except possibly num_logs.
	skew := nStrict - nCore
	if skew < 0 {
		skew = 0
	}
	return nCore, skew, nil
}

// blockMismatchStats counts arch_block rows for project-related blocks joined
// to the live block table on block_num, split into hash/parent mismatches and
// metadata-only mismatches.
func blockMismatchStats(db *sql.DB, addrs []string) (hashBad, metaBad int64, err error) {
	sub := `
		SELECT DISTINCT at.block_num FROM arch_tx at
		WHERE at.tx_hash IN (
			SELECT DISTINCT ae.tx_hash FROM arch_evtlog ae WHERE ae.contract_addr = ANY($1)
		)`
	qHash := `
		SELECT COUNT(*) FROM arch_block ab
		INNER JOIN block b ON b.block_num = ab.block_num
		WHERE ab.block_num IN (` + sub + `)
		AND (
			ab.block_hash IS DISTINCT FROM b.block_hash
			OR ab.parent_hash IS DISTINCT FROM b.parent_hash
		)`
	qMeta := `
		SELECT COUNT(*) FROM arch_block ab
		INNER JOIN block b ON b.block_num = ab.block_num
		WHERE ab.block_num IN (` + sub + `)
		AND (
			ab.num_tx IS DISTINCT FROM b.num_tx
			OR ab.ts IS DISTINCT FROM b.ts
			OR ab.cash_flow IS DISTINCT FROM b.cash_flow
		)
		AND ab.block_hash IS NOT DISTINCT FROM b.block_hash
		AND ab.parent_hash IS NOT DISTINCT FROM b.parent_hash`

	if err := db.QueryRow(qHash, pq.Array(addrs)).Scan(&hashBad); err != nil {
		return 0, 0, fmt.Errorf("block hash mismatch: %w", err)
	}
	if err := db.QueryRow(qMeta, pq.Array(addrs)).Scan(&metaBad); err != nil {
		return 0, 0, fmt.Errorf("block meta mismatch: %w", err)
	}
	return hashBad, metaBad, nil
}

// countRow runs a single COUNT(*)-style query and returns the value.
func countRow(db *sql.DB, query string, args ...any) (int64, error) {
	var n int64
	if err := db.QueryRow(query, args...).Scan(&n); err != nil {
		return 0, fmt.Errorf("count query failed: %w\n%s", err, query)
	}
	return n, nil
}
