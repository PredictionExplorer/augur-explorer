package txcollector

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

// Querier is the narrow pgx query surface used by LoadEventRows.
// *pgxpool.Pool, *pgx.Conn and pgx.Tx satisfy it.
type Querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

// ErrVerificationFailed is returned when a blocking backup mismatch is found.
var ErrVerificationFailed = errors.New("tx-collector verification failed")

// EventRow is one evt_log row checked against the RLP backup.
type EventRow struct {
	BlockNum     int64
	LogIndex     int
	TxHash       string
	ContractAddr string
	Topic0Sig    string
	LogRLP       []byte
}

// VerifyConfig supplies backup rows, reporting limits, and output location.
type VerifyConfig struct {
	OutputDir string
	Rows      []EventRow
	MaxReport int
	Logger    Logger
}

// VerifyStats accumulates counters for one verification run.
type VerifyStats struct {
	EvtRowsTotal       int64
	TxDistinct         int64
	MissingReceiptFile int64
	MissingTxFile      int64
	ReceiptDecodeErr   int64
	TxDecodeErr        int64
	TxHashMismatch     int64
	LogNotInReceipt    int64
	LogIndexMismatch   int64
	LogRLPMismatch     int64
	LogVerifiedOK      int64
	LegacyReceiptBlobs int64
	SQLTxMissingOnDisk int64
	DiskTxExtra        int64
	Reported           int64
}

// HasFailures reports whether a blocking mismatch category is non-zero.
func (s VerifyStats) HasFailures() bool {
	return s.MissingReceiptFile > 0 ||
		s.MissingTxFile > 0 ||
		s.ReceiptDecodeErr > 0 ||
		s.TxDecodeErr > 0 ||
		s.TxHashMismatch > 0 ||
		s.LogNotInReceipt > 0 ||
		s.LogIndexMismatch > 0 ||
		s.LogRLPMismatch > 0 ||
		s.SQLTxMissingOnDisk > 0
}

// HasMismatches reports whether only non-blocking backup-only transactions
// were observed.
func (s VerifyStats) HasMismatches() bool {
	return s.DiskTxExtra > 0
}

// LoadEventRows loads normalized evt_log rows for the selected contracts.
// The caller owns opening and closing db.
func LoadEventRows(
	ctx context.Context,
	db Querier,
	contractAddrs []string,
	fromBlock uint64,
) ([]EventRow, error) {
	if db == nil {
		return nil, errors.New("txcollector: database is required")
	}
	if len(contractAddrs) == 0 {
		return nil, errors.New("txcollector: at least one contract is required")
	}

	const query = `
		SELECT e.block_num, e.log_index, t.tx_hash, a.addr, e.topic0_sig, e.log_rlp
		FROM evt_log e
		INNER JOIN transaction t ON e.tx_id = t.id
		INNER JOIN address a ON e.contract_aid = a.address_id
		WHERE a.addr = ANY($1)
		AND e.block_num >= $2
		ORDER BY t.tx_hash, e.log_index
	`
	rows, err := db.Query(ctx, query, contractAddrs, fromBlock)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []EventRow
	for rows.Next() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		var row EventRow
		if err := rows.Scan(
			&row.BlockNum,
			&row.LogIndex,
			&row.TxHash,
			&row.ContractAddr,
			&row.Topic0Sig,
			&row.LogRLP,
		); err != nil {
			return nil, err
		}
		row.TxHash = toolutil.NormalizeTxHash(row.TxHash)
		row.ContractAddr = toolutil.NormalizeAddr(row.ContractAddr)
		result = append(result, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// Verify compares evt_log rows with receipt and transaction RLP files, then
// checks SQL/on-disk transaction coverage in both directions.
func Verify(ctx context.Context, cfg VerifyConfig) (VerifyStats, error) {
	var stats VerifyStats
	if cfg.OutputDir == "" {
		return stats, errors.New("txcollector: output directory is required")
	}
	if cfg.MaxReport < 0 {
		return stats, errors.New("txcollector: max report must not be negative")
	}
	if err := ctx.Err(); err != nil {
		return stats, err
	}

	var err error
	stats, err = verifyRows(ctx, cfg)
	if err != nil {
		return stats, err
	}
	if err := ctx.Err(); err != nil {
		return stats, err
	}
	onDisk, err := toolutil.BackupTxOnDisk(cfg.OutputDir)
	if err != nil {
		return stats, fmt.Errorf("walk backup: %w", err)
	}
	logf(cfg.Logger, "On-disk tx blobs: %d", len(onDisk))
	if err := checkBackupCoverage(ctx, cfg.Rows, onDisk, &stats, cfg.MaxReport, cfg.Logger); err != nil {
		return stats, err
	}

	LogVerifySummary(cfg.Logger, stats)
	if stats.HasFailures() {
		return stats, ErrVerificationFailed
	}
	return stats, nil
}

type verifyTxKey struct {
	hash     string
	blockNum uint64
}

func verifyRows(ctx context.Context, cfg VerifyConfig) (VerifyStats, error) {
	return verifyRowsWithEncoder(ctx, cfg, toolutil.EncodeLogRLP)
}

func verifyRowsWithEncoder(
	ctx context.Context,
	cfg VerifyConfig,
	encodeLog func(*types.Log) ([]byte, error),
) (VerifyStats, error) {
	stats := VerifyStats{EvtRowsTotal: int64(len(cfg.Rows))}
	byTx := make(map[verifyTxKey][]EventRow)
	for _, input := range cfg.Rows {
		if err := ctx.Err(); err != nil {
			return stats, err
		}
		row := input
		row.TxHash = toolutil.NormalizeTxHash(row.TxHash)
		row.ContractAddr = toolutil.NormalizeAddr(row.ContractAddr)
		if row.BlockNum < 0 {
			// A negative block number is corrupt input; converting it
			// would wrap into an astronomical backup path.
			return stats, fmt.Errorf("txcollector: event row %s has negative block number %d", row.TxHash, row.BlockNum)
		}
		key := verifyTxKey{hash: row.TxHash, blockNum: uint64(row.BlockNum)}
		byTx[key] = append(byTx[key], row)
	}
	stats.TxDistinct = int64(len(byTx))

	keys := make([]verifyTxKey, 0, len(byTx))
	for key := range byTx {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].blockNum != keys[j].blockNum {
			return keys[i].blockNum < keys[j].blockNum
		}
		return keys[i].hash < keys[j].hash
	})
	for key := range byTx {
		sort.SliceStable(byTx[key], func(i, j int) bool {
			return byTx[key][i].LogIndex < byTx[key][j].LogIndex
		})
	}

	report := func(format string, args ...any) {
		if cfg.MaxReport > 0 && stats.Reported >= int64(cfg.MaxReport) {
			return
		}
		logf(cfg.Logger, format, args...)
		stats.Reported++
	}

	for _, key := range keys {
		if err := ctx.Err(); err != nil {
			return stats, err
		}
		eventRows := byTx[key]
		receiptPath := toolutil.ReceiptRLPPath(cfg.OutputDir, key.blockNum, key.hash)
		txPath := toolutil.TxRLPPath(cfg.OutputDir, key.blockNum, key.hash)

		// #nosec G304 -- the path is rooted in operator config and the hash is normalized to fixed hex.
		receiptData, err := os.ReadFile(receiptPath)
		if err != nil {
			stats.MissingReceiptFile++
			report("MISSING receipt file: %s (evt_log rows=%d)", receiptPath, len(eventRows))
			continue
		}

		backupReceipt, legacyFormat, err := toolutil.TryDecodeReceiptRLP(receiptData)
		if err != nil {
			stats.ReceiptDecodeErr++
			report("DECODE receipt %s: %v", receiptPath, err)
			continue
		}
		if legacyFormat {
			stats.LegacyReceiptBlobs++
		}

		// #nosec G304 -- the path is rooted in operator config and the hash is normalized to fixed hex.
		if txData, err := os.ReadFile(txPath); err != nil {
			stats.MissingTxFile++
			report("MISSING tx file: %s", txPath)
		} else {
			var tx types.Transaction
			if err := rlp.DecodeBytes(txData, &tx); err != nil {
				stats.TxDecodeErr++
				report("DECODE tx %s: %v", txPath, err)
			} else if tx.Hash().Hex() != key.hash {
				stats.TxHashMismatch++
				report("TX HASH tx blob %s: blob=%s sql=%s", txPath, tx.Hash().Hex(), key.hash)
			}
		}

		for _, row := range eventRows {
			if err := ctx.Err(); err != nil {
				return stats, err
			}
			log, ok := findBackupLog(backupReceipt, row, legacyFormat)
			if !ok {
				stats.LogNotInReceipt++
				report(
					"LOG missing in receipt: tx=%s log_index=%d contract=%s topic0=%s",
					row.TxHash,
					row.LogIndex,
					row.ContractAddr,
					row.Topic0Sig,
				)
				continue
			}
			if !legacyFormat && int(log.Index) != row.LogIndex {
				stats.LogIndexMismatch++
				report(
					"LOG INDEX tx=%s sql_index=%d receipt_index=%d contract=%s",
					row.TxHash,
					row.LogIndex,
					log.Index,
					row.ContractAddr,
				)
				continue
			}
			encoded, err := encodeLog(log)
			if err != nil {
				report("RLP encode log tx=%s index=%d: %v", row.TxHash, row.LogIndex, err)
				continue
			}
			if !bytes.Equal(encoded, row.LogRLP) {
				stats.LogRLPMismatch++
				report(
					"LOG RLP mismatch: tx=%s log_index=%d contract=%s topic0=%s",
					row.TxHash,
					row.LogIndex,
					row.ContractAddr,
					row.Topic0Sig,
				)
				continue
			}
			stats.LogVerifiedOK++
		}
	}
	return stats, nil
}

func findBackupLog(receipt *toolutil.BackupReceipt, row EventRow, legacyFormat bool) (*types.Log, bool) {
	var indexMismatch *types.Log
	for i := range receipt.Logs {
		backupLog := &receipt.Logs[i]
		if toolutil.NormalizeAddr(backupLog.Address.Hex()) != row.ContractAddr {
			continue
		}
		log := backupLog.ToTypesLog()
		if !legacyFormat && int(backupLog.Index) == row.LogIndex {
			return log, true
		}
		encoded, err := toolutil.EncodeLogRLP(log)
		if err != nil || !bytes.Equal(encoded, row.LogRLP) {
			continue
		}
		if legacyFormat {
			return log, true
		}
		if indexMismatch == nil {
			indexMismatch = log
		}
	}
	return indexMismatch, indexMismatch != nil
}

func checkBackupCoverage(
	ctx context.Context,
	rows []EventRow,
	onDisk map[string]uint64,
	stats *VerifyStats,
	maxReport int,
	logger Logger,
) error {
	sqlTxs := make(map[string]uint64)
	for _, row := range rows {
		if err := ctx.Err(); err != nil {
			return err
		}
		hash := toolutil.NormalizeTxHash(row.TxHash)
		if _, exists := sqlTxs[hash]; !exists {
			sqlTxs[hash] = uint64(row.BlockNum) // #nosec G115 -- rows already validated non-negative by verifyRowsWithEncoder
		}
	}

	reported := 0
	report := func(format string, args ...any) {
		if maxReport > 0 && reported >= maxReport {
			return
		}
		logf(logger, format, args...)
		reported++
	}

	sqlHashes := sortedHashes(sqlTxs)
	for _, hash := range sqlHashes {
		if err := ctx.Err(); err != nil {
			return err
		}
		if _, exists := onDisk[hash]; !exists {
			stats.SQLTxMissingOnDisk++
			report("SQL tx missing on disk: %s block=%d", hash, sqlTxs[hash])
		}
	}

	diskHashes := sortedHashes(onDisk)
	for _, hash := range diskHashes {
		if err := ctx.Err(); err != nil {
			return err
		}
		if _, exists := sqlTxs[hash]; !exists {
			stats.DiskTxExtra++
			report(
				"MISMATCH backup-only tx (on disk, not in evt_log): %s block=%d",
				hash,
				onDisk[hash],
			)
		}
	}
	return nil
}

func sortedHashes(values map[string]uint64) []string {
	hashes := make([]string, 0, len(values))
	for hash := range values {
		hashes = append(hashes, hash)
	}
	sort.Strings(hashes)
	return hashes
}

// LogVerifySummary writes the operator-facing verification summary.
func LogVerifySummary(logger Logger, stats VerifyStats) {
	logf(logger, "")
	logf(logger, "=== SUMMARY ===")
	logf(logger, "evt_log rows:              %d", stats.EvtRowsTotal)
	logf(logger, "distinct tx (SQL):         %d", stats.TxDistinct)
	logf(logger, "logs verified OK:          %d", stats.LogVerifiedOK)
	logf(logger, "legacy receipt blobs:      %d (log index verified via log_rlp only)", stats.LegacyReceiptBlobs)
	logf(logger, "missing receipt file:      %d", stats.MissingReceiptFile)
	logf(logger, "missing tx file:           %d", stats.MissingTxFile)
	logf(logger, "receipt decode errors:     %d", stats.ReceiptDecodeErr)
	logf(logger, "tx decode errors:          %d", stats.TxDecodeErr)
	logf(logger, "tx hash mismatches:        %d", stats.TxHashMismatch)
	logf(logger, "log not in receipt:        %d", stats.LogNotInReceipt)
	logf(logger, "log index mismatches:      %d", stats.LogIndexMismatch)
	logf(logger, "log_rlp mismatches:        %d", stats.LogRLPMismatch)
	logf(logger, "SQL tx missing on disk:    %d", stats.SQLTxMissingOnDisk)
	logf(
		logger,
		"backup-only tx (mismatch): %d (on disk, no matching evt_log — often DAO/admin deploy events)",
		stats.DiskTxExtra,
	)
	switch {
	case stats.HasFailures():
		logf(logger, "RESULT: FAILED")
	case stats.HasMismatches():
		logf(logger, "RESULT: OK — SQL evt_log matches backup; backup-only txs noted above (expected for unindexed contracts/events)")
	default:
		logf(logger, "RESULT: OK")
	}
}
