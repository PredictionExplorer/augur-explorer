package archive

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
)

// DefaultExportBatchSize is the row batch used by archive export.
const DefaultExportBatchSize = 20_000

// Logger is the small logging surface used by archive operations.
type Logger interface {
	Printf(format string, args ...any)
	Println(args ...any)
}

// ExportStats summarizes one project's export work.
type ExportStats struct {
	EventLogsProcessed int64
	TransactionsCopied int64
	BlocksCopied       int64
}

// ExportResult associates a project's export stats with its project name.
type ExportResult struct {
	Project string
	Stats   ExportStats
}

// ProjectExporter is implemented by SQLExporter and is intentionally narrow
// so project ordering, cancellation, and error propagation can be unit tested.
type ProjectExporter interface {
	ExportProject(ctx context.Context, project string) (ExportStats, error)
}

// ExportProjects exports projects in the supplied order and stops at the
// first cancellation or project error.
func ExportProjects(ctx context.Context, projects []string, exporter ProjectExporter) ([]ExportResult, error) {
	results := make([]ExportResult, 0, len(projects))
	if exporter == nil {
		return results, errors.New("archive export: exporter is required")
	}
	for _, project := range projects {
		if err := ctx.Err(); err != nil {
			return results, err
		}
		stats, err := exporter.ExportProject(ctx, project)
		if err != nil {
			return results, err
		}
		results = append(results, ExportResult{Project: project, Stats: stats})
	}
	return results, nil
}

// SQLExporter copies archive data between two database/sql handles. The
// caller owns both handles.
type SQLExporter struct {
	Source      *sql.DB
	Destination *sql.DB
	BatchSize   int
	Logger      Logger
}

// ExportProject copies evt_log, transaction, and block rows for one project.
func (e *SQLExporter) ExportProject(ctx context.Context, project string) (ExportStats, error) {
	var stats ExportStats
	if e == nil || e.Source == nil || e.Destination == nil {
		return stats, errors.New("archive export: source and destination databases are required")
	}
	e.printf("Project type: %s", project)

	contracts, err := LoadProjectContracts(ctx, e.Source, project)
	if err != nil {
		return stats, err
	}
	e.printf("Found %d contract address IDs: %v", len(contracts.AddressIDs), contracts.AddressIDs)

	stats.EventLogsProcessed, err = e.exportEventLogs(ctx, contracts)
	if err != nil {
		return stats, err
	}
	stats.TransactionsCopied, err = e.exportTransactions(ctx)
	if err != nil {
		return stats, err
	}
	stats.BlocksCopied, err = e.exportBlocks(ctx)
	if err != nil {
		return stats, err
	}

	e.printf("=== Export complete (%s) ===", project)
	return stats, nil
}

// ResumePosition is one contract's archive watermark.
type ResumePosition struct {
	ContractAddress string
	MaxEventID      int64
}

// ResumeFloor returns the minimum contract watermark, or zero for an empty
// input. This is the safe project resume point.
func ResumeFloor(positions []ResumePosition) int64 {
	if len(positions) == 0 {
		return 0
	}
	floor := positions[0].MaxEventID
	for _, position := range positions[1:] {
		if position.MaxEventID < floor {
			floor = position.MaxEventID
		}
	}
	if floor < 0 {
		return 0
	}
	return floor
}

// EventLogResumeFloor loads each contract's independent archive watermark and
// returns their minimum. A global maximum is unsafe when projects share an
// archive.
func EventLogResumeFloor(ctx context.Context, db *sql.DB, contractAddresses []string) (int64, []ResumePosition, error) {
	positions := make([]ResumePosition, 0, len(contractAddresses))
	for _, address := range contractAddresses {
		var maxEventID int64
		err := db.QueryRowContext(ctx, `
			SELECT COALESCE(MAX(evt_id), 0)
			FROM arch_evtlog
			WHERE contract_addr = $1
		`, address).Scan(&maxEventID)
		if err != nil {
			return 0, positions, fmt.Errorf("read resume position for contract %s: %w", address, err)
		}
		positions = append(positions, ResumePosition{
			ContractAddress: address,
			MaxEventID:      maxEventID,
		})
	}
	return ResumeFloor(positions), positions, nil
}

func (e *SQLExporter) exportEventLogs(ctx context.Context, contracts Contracts) (int64, error) {
	e.println("=== Exporting event logs for contracts ===")
	e.println("Counting events...")
	countStart := time.Now()
	var totalEvents int64
	err := e.Source.QueryRowContext(ctx,
		"SELECT COUNT(*) FROM evt_log WHERE contract_aid = ANY($1)",
		pq.Array(contracts.AddressIDs),
	).Scan(&totalEvents)
	if err != nil {
		return 0, fmt.Errorf("count events: %w", err)
	}
	e.printf("Total events on source for these contracts: %d (counted in %.1f ms)",
		totalEvents, time.Since(countStart).Seconds()*1000)

	currentID, positions, err := EventLogResumeFloor(ctx, e.Destination, contracts.Addresses)
	if err != nil {
		return 0, err
	}
	e.println("Per-contract MAX(evt_id) already in arch_evtlog (0 = no rows for that contract):")
	for _, position := range positions {
		e.printf("  %s -> %d", position.ContractAddress, position.MaxEventID)
	}
	e.printf("Resume floor evt_id (min over contracts) = %d — exporting source evt_log rows with id > %d", currentID, currentID)

	insertStmt, err := e.Destination.PrepareContext(ctx, `
		INSERT INTO arch_evtlog (block_num, evt_id, log_index, tx_hash, contract_addr, topic0_sig, log_rlp)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (tx_hash, log_index) DO NOTHING
	`)
	if err != nil {
		return 0, fmt.Errorf("prepare arch_evtlog insert: %w", err)
	}
	defer func() { _ = insertStmt.Close() }()

	txIDs := make(map[int64]struct{})
	var totalProcessed int64
	startTime := time.Now()
	for {
		if err := ctx.Err(); err != nil {
			return totalProcessed, err
		}
		batchStart := time.Now()
		e.printf("Querying events where id > %d (limit %d)...", currentID, e.batchSize())
		batchCount, lastID, err := e.exportEventLogBatch(ctx, insertStmt, contracts.AddressIDs, currentID, txIDs)
		if err != nil {
			return totalProcessed, err
		}
		if batchCount == 0 {
			break
		}
		currentID = lastID
		totalProcessed += int64(batchCount)
		e.printf("EventLogs: %d rows (%.1f ms) | Total: %d/%d | %.1f/sec | %.1f%%",
			batchCount,
			time.Since(batchStart).Seconds()*1000,
			totalProcessed,
			totalEvents,
			perSecond(totalProcessed, startTime),
			percentage(totalProcessed, totalEvents),
		)
	}
	e.printf("EventLogs export complete. Total: %d, Unique tx_ids: %d", totalProcessed, len(txIDs))
	return totalProcessed, nil
}

func (e *SQLExporter) exportEventLogBatch(
	ctx context.Context,
	insertStmt *sql.Stmt,
	contractAIDs []int64,
	afterID int64,
	txIDs map[int64]struct{},
) (int, int64, error) {
	queryStart := time.Now()
	rows, err := e.Source.QueryContext(ctx, `
		SELECT e.block_num, e.id, e.tx_id, e.log_index, t.tx_hash, a.addr, e.topic0_sig, e.log_rlp
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
		JOIN address a ON e.contract_aid = a.address_id
		WHERE e.contract_aid = ANY($1) AND e.id > $2
		ORDER BY e.id
		LIMIT $3
	`, pq.Array(contractAIDs), afterID, e.batchSize())
	if err != nil {
		return 0, 0, fmt.Errorf("query evt_log batch: %w", err)
	}
	defer func() { _ = rows.Close() }()
	e.printf("Query returned in %.1f ms, processing rows...", time.Since(queryStart).Seconds()*1000)

	var (
		batchCount int
		lastID     int64
	)
	for rows.Next() {
		var (
			blockNum   int64
			eventID    int64
			txID       int64
			logIndex   int
			txHash     string
			contract   string
			topic0     string
			encodedLog []byte
		)
		if err := rows.Scan(&blockNum, &eventID, &txID, &logIndex, &txHash, &contract, &topic0, &encodedLog); err != nil {
			return 0, 0, fmt.Errorf("scan evt_log row: %w", err)
		}
		if _, err := insertStmt.ExecContext(ctx,
			blockNum, eventID, logIndex, txHash, contract, topic0, encodedLog,
		); err != nil {
			return 0, 0, fmt.Errorf("insert arch_evtlog evt_id %d tx %s log_index %d: %w", eventID, txHash, logIndex, err)
		}
		txIDs[txID] = struct{}{}
		lastID = eventID
		batchCount++
	}
	if err := rows.Err(); err != nil {
		return 0, 0, fmt.Errorf("iterate evt_log batch: %w", err)
	}
	return batchCount, lastID, nil
}

func (e *SQLExporter) exportTransactions(ctx context.Context) (int64, error) {
	e.println("=== Exporting transactions ===")
	e.println("Querying destination for tx_hashes missing from arch_tx...")
	missingTxHashes, err := queryStringColumn(ctx, e.Destination, `
		SELECT DISTINCT e.tx_hash
		FROM arch_evtlog e
		LEFT JOIN arch_tx tx ON e.tx_hash = tx.tx_hash
		WHERE tx.tx_hash IS NULL
		ORDER BY e.tx_hash
	`)
	if err != nil {
		return 0, fmt.Errorf("find missing transactions: %w", err)
	}
	e.printf("Transactions to export: %d", len(missingTxHashes))
	if len(missingTxHashes) == 0 {
		e.println("No missing transactions to export")
		return 0, nil
	}

	insertStmt, err := e.Destination.PrepareContext(ctx, `
		INSERT INTO arch_tx (block_num, from_aid, to_aid, gas_used, tx_index, num_logs, ctrct_create, value, gas_price, tx_hash, input_sig)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (tx_hash) DO NOTHING
	`)
	if err != nil {
		return 0, fmt.Errorf("prepare arch_tx insert: %w", err)
	}
	defer func() { _ = insertStmt.Close() }()

	blockNums := make(map[int64]struct{})
	var totalProcessed int64
	startTime := time.Now()
	for i := 0; i < len(missingTxHashes); i += e.batchSize() {
		if err := ctx.Err(); err != nil {
			return totalProcessed, err
		}
		batchStart := time.Now()
		end := min(i+e.batchSize(), len(missingTxHashes))
		batch := missingTxHashes[i:end]
		e.printf("Querying %d transactions...", len(batch))
		batchCount, err := e.exportTransactionBatch(ctx, insertStmt, batch, blockNums)
		if err != nil {
			return totalProcessed, err
		}
		totalProcessed += int64(batchCount)
		e.printf("Transactions: %d rows (%.1f ms) | Total: %d/%d | %.1f/sec | %.1f%%",
			batchCount,
			time.Since(batchStart).Seconds()*1000,
			totalProcessed,
			len(missingTxHashes),
			perSecond(totalProcessed, startTime),
			percentage(int64(end), int64(len(missingTxHashes))),
		)
	}
	e.printf("Transactions export complete. Total: %d, Unique blocks: %d", totalProcessed, len(blockNums))
	return totalProcessed, nil
}

func (e *SQLExporter) exportTransactionBatch(
	ctx context.Context,
	insertStmt *sql.Stmt,
	txHashes []string,
	blockNums map[int64]struct{},
) (int, error) {
	queryStart := time.Now()
	rows, err := e.Source.QueryContext(ctx, `
		SELECT block_num, from_aid, to_aid, gas_used, tx_index, num_logs, ctrct_create, value, gas_price, tx_hash, input_sig
		FROM transaction
		WHERE tx_hash = ANY($1)
	`, pq.Array(txHashes))
	if err != nil {
		return 0, fmt.Errorf("query transaction batch: %w", err)
	}
	defer func() { _ = rows.Close() }()
	e.printf("Query returned in %.1f ms", time.Since(queryStart).Seconds()*1000)

	batchCount := 0
	for rows.Next() {
		var (
			blockNum     int64
			fromAID      int64
			toAID        int64
			gasUsed      int64
			txIndex      int
			numLogs      int
			contractMade bool
			value        string
			gasPrice     string
			txHash       string
			inputSig     sql.NullString
		)
		if err := rows.Scan(
			&blockNum, &fromAID, &toAID, &gasUsed, &txIndex, &numLogs,
			&contractMade, &value, &gasPrice, &txHash, &inputSig,
		); err != nil {
			return 0, fmt.Errorf("scan transaction row: %w", err)
		}
		var inputSigValue any
		if inputSig.Valid {
			inputSigValue = inputSig.String
		}
		if _, err := insertStmt.ExecContext(ctx,
			blockNum, fromAID, toAID, gasUsed, txIndex, numLogs, contractMade,
			value, gasPrice, txHash, inputSigValue,
		); err != nil {
			return 0, fmt.Errorf("insert arch_tx %s: %w", txHash, err)
		}
		blockNums[blockNum] = struct{}{}
		batchCount++
	}
	if err := rows.Err(); err != nil {
		return 0, fmt.Errorf("iterate transaction batch: %w", err)
	}
	return batchCount, nil
}

func (e *SQLExporter) exportBlocks(ctx context.Context) (int64, error) {
	e.println("=== Exporting blocks ===")
	e.println("Finding missing blocks in archive...")
	missingBlockNums, err := queryInt64Column(ctx, e.Destination, `
		SELECT DISTINCT tx.block_num
		FROM arch_tx tx
		LEFT JOIN arch_block b ON tx.block_num = b.block_num
		WHERE b.block_num IS NULL
		ORDER BY tx.block_num
	`)
	if err != nil {
		return 0, fmt.Errorf("find missing blocks: %w", err)
	}
	e.printf("Blocks to export: %d", len(missingBlockNums))
	if len(missingBlockNums) == 0 {
		e.println("No missing blocks to export")
		return 0, nil
	}

	insertStmt, err := e.Destination.PrepareContext(ctx, `
		INSERT INTO arch_block (block_num, num_tx, ts, cash_flow, block_hash, parent_hash)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (block_hash) DO NOTHING
	`)
	if err != nil {
		return 0, fmt.Errorf("prepare arch_block insert: %w", err)
	}
	defer func() { _ = insertStmt.Close() }()

	var totalProcessed int64
	startTime := time.Now()
	for i := 0; i < len(missingBlockNums); i += e.batchSize() {
		if err := ctx.Err(); err != nil {
			return totalProcessed, err
		}
		batchStart := time.Now()
		end := min(i+e.batchSize(), len(missingBlockNums))
		batch := missingBlockNums[i:end]
		e.printf("Querying %d blocks...", len(batch))
		batchCount, err := e.exportBlockBatch(ctx, insertStmt, batch)
		if err != nil {
			return totalProcessed, err
		}
		totalProcessed += int64(batchCount)
		e.printf("Blocks: %d rows (%.1f ms) | Total: %d/%d | %.1f/sec | %.1f%%",
			batchCount,
			time.Since(batchStart).Seconds()*1000,
			totalProcessed,
			len(missingBlockNums),
			perSecond(totalProcessed, startTime),
			percentage(int64(end), int64(len(missingBlockNums))),
		)
	}
	e.printf("Blocks export complete. Total: %d", totalProcessed)
	return totalProcessed, nil
}

func (e *SQLExporter) exportBlockBatch(ctx context.Context, insertStmt *sql.Stmt, blockNums []int64) (int, error) {
	queryStart := time.Now()
	rows, err := e.Source.QueryContext(ctx, `
		SELECT block_num, num_tx, ts, cash_flow, block_hash, parent_hash
		FROM block
		WHERE block_num = ANY($1)
	`, pq.Array(blockNums))
	if err != nil {
		return 0, fmt.Errorf("query block batch: %w", err)
	}
	defer func() { _ = rows.Close() }()
	e.printf("Query returned in %.1f ms", time.Since(queryStart).Seconds()*1000)

	batchCount := 0
	for rows.Next() {
		var (
			blockNum  int64
			numTx     int64
			timestamp time.Time
			cashFlow  string
			blockHash string
			parent    string
		)
		if err := rows.Scan(&blockNum, &numTx, &timestamp, &cashFlow, &blockHash, &parent); err != nil {
			return 0, fmt.Errorf("scan block row: %w", err)
		}
		if _, err := insertStmt.ExecContext(ctx, blockNum, numTx, timestamp, cashFlow, blockHash, parent); err != nil {
			return 0, fmt.Errorf("insert arch_block %d: %w", blockNum, err)
		}
		batchCount++
	}
	if err := rows.Err(); err != nil {
		return 0, fmt.Errorf("iterate block batch: %w", err)
	}
	return batchCount, nil
}

func queryStringColumn(ctx context.Context, db *sql.DB, query string) ([]string, error) {
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var values []string
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, rows.Err()
}

func queryInt64Column(ctx context.Context, db *sql.DB, query string) ([]int64, error) {
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var values []int64
	for rows.Next() {
		var value int64
		if err := rows.Scan(&value); err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, rows.Err()
}

func (e *SQLExporter) batchSize() int {
	if e.BatchSize > 0 {
		return e.BatchSize
	}
	return DefaultExportBatchSize
}

func (e *SQLExporter) printf(format string, args ...any) {
	if e.Logger != nil {
		e.Logger.Printf(format, args...)
	}
}

func (e *SQLExporter) println(args ...any) {
	if e.Logger != nil {
		e.Logger.Println(args...)
	}
}

func perSecond(count int64, started time.Time) float64 {
	elapsed := time.Since(started).Seconds()
	if elapsed <= 0 {
		return 0
	}
	return float64(count) / elapsed
}

func percentage(count, total int64) float64 {
	if total <= 0 {
		return 0
	}
	return float64(count) / float64(total) * 100
}
