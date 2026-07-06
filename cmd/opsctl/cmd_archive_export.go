package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/spf13/cobra"
)

// archiveExportBatchSize is the row batch used for the evt_log, transaction
// and block export queries.
const archiveExportBatchSize = 20000

// newArchiveExportCmd builds `opsctl archive export`, the replacement for the
// standalone archive_export tool.
func newArchiveExportCmd() *cobra.Command {
	var (
		srcConn     string
		dstConn     string
		projectType string
	)
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Copy evt_log / transaction / block rows for project contracts into arch_* tables",
		Long: `Copies evt_log rows for a project's contracts from a source (production)
database into the destination's arch_evtlog table, then fills arch_tx and
arch_block for every referenced transaction and block.

arch_evtlog rows are keyed by (tx_hash, log_index) so archives stay valid if
evt_log ids change; incremental export resumes from the per-contract minimum
MAX(evt_id) already archived on the destination.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			projects, err := resolveProjects(projectType)
			if err != nil {
				return err
			}

			srcDB, err := sql.Open("postgres", srcConn)
			if err != nil {
				return fmt.Errorf("connect to source: %w", err)
			}
			defer srcDB.Close()
			srcDB.SetMaxOpenConns(2)

			dstDB, err := sql.Open("postgres", dstConn)
			if err != nil {
				return fmt.Errorf("connect to destination: %w", err)
			}
			defer dstDB.Close()
			dstDB.SetMaxOpenConns(2)

			for _, p := range projects {
				if err := runArchiveExport(srcDB, dstDB, p); err != nil {
					return err
				}
			}

			log.Println("=== All exports complete ===")
			return nil
		},
	}
	cmd.Flags().StringVar(&srcConn, "src", "", "Source DB connection string (production)")
	cmd.Flags().StringVar(&dstConn, "dst", "", "Destination DB connection string (dev)")
	cmd.Flags().StringVar(&projectType, "project", "", "Project: randomwalk | cosmicgame | both (runs cosmicgame then randomwalk)")
	_ = cmd.MarkFlagRequired("src")
	_ = cmd.MarkFlagRequired("dst")
	_ = cmd.MarkFlagRequired("project")
	return cmd
}

func init() { archiveCmd.AddCommand(newArchiveExportCmd()) }

// runArchiveExport copies evt_log into arch_evtlog, then arch_tx and
// arch_block, for one project.
func runArchiveExport(srcDB, dstDB *sql.DB, project string) error {
	log.Printf("Project type: %s", project)

	contractAids, contractAddrs, err := projectContracts(srcDB, project)
	if err != nil {
		return err
	}
	log.Printf("Found %d contract address IDs: %v", len(contractAids), contractAids)

	if err := exportEventLogs(srcDB, dstDB, contractAids, contractAddrs); err != nil {
		return err
	}
	if err := exportTransactions(srcDB, dstDB); err != nil {
		return err
	}
	if err := exportBlocks(srcDB, dstDB); err != nil {
		return err
	}

	log.Printf("=== Export complete (%s) ===", project)
	return nil
}

// eventLogResumeFloor returns min(COALESCE(MAX(evt_id),0)) over the project
// contracts in arch_evtlog. Using the global MAX(evt_id) would skip older
// logs for another project sharing the same archive.
func eventLogResumeFloor(dstDB *sql.DB, contractAddrs []string) (int64, error) {
	if len(contractAddrs) == 0 {
		return 0, nil
	}
	log.Println("Per-contract MAX(evt_id) already in arch_evtlog (0 = no rows for that contract):")
	var floor int64 = -1
	for _, ca := range contractAddrs {
		var maxForContract int64
		err := dstDB.QueryRow(`
			SELECT COALESCE(MAX(evt_id), 0) FROM arch_evtlog WHERE contract_addr = $1
		`, ca).Scan(&maxForContract)
		if err != nil {
			return 0, fmt.Errorf("read resume position for contract %s: %w", ca, err)
		}
		log.Printf("  %s -> %d", ca, maxForContract)
		if floor < 0 || maxForContract < floor {
			floor = maxForContract
		}
	}
	if floor < 0 {
		floor = 0
	}
	log.Printf("Resume floor evt_id (min over contracts) = %d — exporting source evt_log rows with id > %d", floor, floor)
	return floor, nil
}

// exportEventLogs copies missing evt_log rows for the given contracts into
// arch_evtlog. Duplicate chain logs are suppressed by (tx_hash, log_index),
// not evt_id.
func exportEventLogs(srcDB, dstDB *sql.DB, contractAids []int64, contractAddrs []string) error {
	log.Println("=== Exporting event logs for contracts ===")

	log.Println("Counting events...")
	countStart := time.Now()
	var totalEvents int64
	err := srcDB.QueryRow("SELECT COUNT(*) FROM evt_log WHERE contract_aid = ANY($1)", pq.Array(contractAids)).Scan(&totalEvents)
	if err != nil {
		return fmt.Errorf("count events: %w", err)
	}
	log.Printf("Total events on source for these contracts: %d (counted in %.1f ms)", totalEvents, time.Since(countStart).Seconds()*1000)

	currentID, err := eventLogResumeFloor(dstDB, contractAddrs)
	if err != nil {
		return err
	}

	insertStmt, err := dstDB.Prepare(`
		INSERT INTO arch_evtlog (block_num, evt_id, log_index, tx_hash, contract_addr, topic0_sig, log_rlp)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (tx_hash, log_index) DO NOTHING
	`)
	if err != nil {
		return fmt.Errorf("prepare arch_evtlog insert: %w", err)
	}
	defer insertStmt.Close()

	txIDs := make(map[int64]bool)
	totalInserted := int64(0)
	startTime := time.Now()

	for {
		batchStart := time.Now()
		log.Printf("Querying events where id > %d (limit %d)...", currentID, archiveExportBatchSize)

		batchCount, lastID, err := exportEventLogBatch(srcDB, insertStmt, contractAids, currentID, txIDs)
		if err != nil {
			return err
		}
		if batchCount == 0 {
			break
		}

		currentID = lastID
		totalInserted += int64(batchCount)
		elapsed := time.Since(startTime).Seconds()
		rate := float64(totalInserted) / elapsed
		pct := float64(totalInserted) / float64(totalEvents) * 100

		log.Printf("EventLogs: %d rows (%.1f ms) | Total: %d/%d | %.1f/sec | %.1f%%",
			batchCount,
			time.Since(batchStart).Seconds()*1000,
			totalInserted, totalEvents, rate, pct)
	}

	log.Printf("EventLogs export complete. Total: %d, Unique tx_ids: %d", totalInserted, len(txIDs))
	return nil
}

// exportEventLogBatch copies one id-ordered batch of evt_log rows, recording
// the tx ids it saw, and returns the number of rows processed plus the last
// evt_log id of the batch.
func exportEventLogBatch(srcDB *sql.DB, insertStmt *sql.Stmt, contractAids []int64, afterID int64, txIDs map[int64]bool) (int, int64, error) {
	queryStart := time.Now()
	rows, err := srcDB.Query(`
		SELECT e.block_num, e.id, e.tx_id, e.log_index, t.tx_hash, a.addr, e.topic0_sig, e.log_rlp
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
		JOIN address a ON e.contract_aid = a.address_id
		WHERE e.contract_aid = ANY($1) AND e.id > $2
		ORDER BY e.id
		LIMIT $3
	`, pq.Array(contractAids), afterID, archiveExportBatchSize)
	if err != nil {
		return 0, 0, fmt.Errorf("query evt_log batch: %w", err)
	}
	defer rows.Close()
	log.Printf("Query returned in %.1f ms, processing rows...", time.Since(queryStart).Seconds()*1000)

	batchCount := 0
	var lastID int64
	for rows.Next() {
		var blockNum, evtID, txID int64
		var logIndex int
		var txHash, contractAddr, topic0Sig string
		var logRLP []byte

		if err := rows.Scan(&blockNum, &evtID, &txID, &logIndex, &txHash, &contractAddr, &topic0Sig, &logRLP); err != nil {
			return 0, 0, fmt.Errorf("scan evt_log row: %w", err)
		}
		if _, err := insertStmt.Exec(blockNum, evtID, logIndex, txHash, contractAddr, topic0Sig, logRLP); err != nil {
			return 0, 0, fmt.Errorf("insert arch_evtlog evt_id %d tx %s log_index %d: %w", evtID, txHash, logIndex, err)
		}

		txIDs[txID] = true
		lastID = evtID
		batchCount++
	}
	return batchCount, lastID, rows.Err()
}

// exportTransactions copies into arch_tx every transaction referenced by
// arch_evtlog that the archive does not have yet.
func exportTransactions(srcDB, dstDB *sql.DB) error {
	log.Println("=== Exporting transactions ===")

	log.Println("Querying destination for tx_hashes missing from arch_tx...")
	missingTxHashes, err := queryStringColumn(dstDB, `
		SELECT DISTINCT e.tx_hash
		FROM arch_evtlog e
		LEFT JOIN arch_tx tx ON e.tx_hash = tx.tx_hash
		WHERE tx.tx_hash IS NULL
	`)
	if err != nil {
		return fmt.Errorf("find missing transactions: %w", err)
	}

	log.Printf("Transactions to export: %d", len(missingTxHashes))
	if len(missingTxHashes) == 0 {
		log.Println("No missing transactions to export")
		return nil
	}

	insertStmt, err := dstDB.Prepare(`
		INSERT INTO arch_tx (block_num, from_aid, to_aid, gas_used, tx_index, num_logs, ctrct_create, value, gas_price, tx_hash, input_sig)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (tx_hash) DO NOTHING
	`)
	if err != nil {
		return fmt.Errorf("prepare arch_tx insert: %w", err)
	}
	defer insertStmt.Close()

	blockNums := make(map[int64]bool)
	totalInserted := int64(0)
	startTime := time.Now()

	for i := 0; i < len(missingTxHashes); i += archiveExportBatchSize {
		batchStart := time.Now()
		end := min(i+archiveExportBatchSize, len(missingTxHashes))
		batch := missingTxHashes[i:end]

		log.Printf("Querying %d transactions...", len(batch))
		batchCount, err := exportTransactionBatch(srcDB, insertStmt, batch, blockNums)
		if err != nil {
			return err
		}

		totalInserted += int64(batchCount)
		elapsed := time.Since(startTime).Seconds()
		rate := float64(totalInserted) / elapsed
		pct := float64(i+batchCount) / float64(len(missingTxHashes)) * 100

		log.Printf("Transactions: %d rows (%.1f ms) | Total: %d/%d | %.1f/sec | %.1f%%",
			batchCount,
			time.Since(batchStart).Seconds()*1000,
			totalInserted, len(missingTxHashes), rate, pct)
	}

	log.Printf("Transactions export complete. Total: %d, Unique blocks: %d", totalInserted, len(blockNums))
	return nil
}

// exportTransactionBatch copies one batch of transactions by tx_hash,
// recording the block numbers it saw.
func exportTransactionBatch(srcDB *sql.DB, insertStmt *sql.Stmt, txHashes []string, blockNums map[int64]bool) (int, error) {
	queryStart := time.Now()
	rows, err := srcDB.Query(`
		SELECT block_num, from_aid, to_aid, gas_used, tx_index, num_logs, ctrct_create, value, gas_price, tx_hash, input_sig
		FROM transaction
		WHERE tx_hash = ANY($1)
	`, pq.Array(txHashes))
	if err != nil {
		return 0, fmt.Errorf("query transaction batch: %w", err)
	}
	defer rows.Close()
	log.Printf("Query returned in %.1f ms", time.Since(queryStart).Seconds()*1000)

	batchCount := 0
	for rows.Next() {
		var blockNum, fromAid, toAid, gasUsed int64
		var txIndex, numLogs int
		var ctrctCreate bool
		var value, gasPrice, txHash string
		var inputSig sql.NullString

		if err := rows.Scan(&blockNum, &fromAid, &toAid, &gasUsed, &txIndex, &numLogs, &ctrctCreate, &value, &gasPrice, &txHash, &inputSig); err != nil {
			return 0, fmt.Errorf("scan transaction row: %w", err)
		}

		var inputSigVal any
		if inputSig.Valid {
			inputSigVal = inputSig.String
		}
		if _, err := insertStmt.Exec(blockNum, fromAid, toAid, gasUsed, txIndex, numLogs, ctrctCreate, value, gasPrice, txHash, inputSigVal); err != nil {
			return 0, fmt.Errorf("insert arch_tx %s: %w", txHash, err)
		}

		blockNums[blockNum] = true
		batchCount++
	}
	return batchCount, rows.Err()
}

// exportBlocks copies into arch_block every block referenced by arch_tx that
// the archive does not have yet.
func exportBlocks(srcDB, dstDB *sql.DB) error {
	log.Println("=== Exporting blocks ===")

	log.Println("Finding missing blocks in archive...")
	missingBlockNums, err := queryInt64Column(dstDB, `
		SELECT DISTINCT tx.block_num
		FROM arch_tx tx
		LEFT JOIN arch_block b ON tx.block_num = b.block_num
		WHERE b.block_num IS NULL
	`)
	if err != nil {
		return fmt.Errorf("find missing blocks: %w", err)
	}

	log.Printf("Blocks to export: %d", len(missingBlockNums))
	if len(missingBlockNums) == 0 {
		log.Println("No missing blocks to export")
		return nil
	}

	insertStmt, err := dstDB.Prepare(`
		INSERT INTO arch_block (block_num, num_tx, ts, cash_flow, block_hash, parent_hash)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (block_hash) DO NOTHING
	`)
	if err != nil {
		return fmt.Errorf("prepare arch_block insert: %w", err)
	}
	defer insertStmt.Close()

	totalInserted := int64(0)
	startTime := time.Now()

	for i := 0; i < len(missingBlockNums); i += archiveExportBatchSize {
		batchStart := time.Now()
		end := min(i+archiveExportBatchSize, len(missingBlockNums))
		batch := missingBlockNums[i:end]

		log.Printf("Querying %d blocks...", len(batch))
		batchCount, err := exportBlockBatch(srcDB, insertStmt, batch)
		if err != nil {
			return err
		}

		totalInserted += int64(batchCount)
		elapsed := time.Since(startTime).Seconds()
		rate := float64(totalInserted) / elapsed
		pct := float64(i+batchCount) / float64(len(missingBlockNums)) * 100

		log.Printf("Blocks: %d rows (%.1f ms) | Total: %d/%d | %.1f/sec | %.1f%%",
			batchCount,
			time.Since(batchStart).Seconds()*1000,
			totalInserted, len(missingBlockNums), rate, pct)
	}

	log.Printf("Blocks export complete. Total: %d", totalInserted)
	return nil
}

// exportBlockBatch copies one batch of blocks by block_num.
func exportBlockBatch(srcDB *sql.DB, insertStmt *sql.Stmt, blockNums []int64) (int, error) {
	queryStart := time.Now()
	rows, err := srcDB.Query(`
		SELECT block_num, num_tx, ts, cash_flow, block_hash, parent_hash
		FROM block
		WHERE block_num = ANY($1)
	`, pq.Array(blockNums))
	if err != nil {
		return 0, fmt.Errorf("query block batch: %w", err)
	}
	defer rows.Close()
	log.Printf("Query returned in %.1f ms", time.Since(queryStart).Seconds()*1000)

	batchCount := 0
	for rows.Next() {
		var blockNum, numTx int64
		var ts time.Time
		var cashFlow, blockHash, parentHash string

		if err := rows.Scan(&blockNum, &numTx, &ts, &cashFlow, &blockHash, &parentHash); err != nil {
			return 0, fmt.Errorf("scan block row: %w", err)
		}
		if _, err := insertStmt.Exec(blockNum, numTx, ts, cashFlow, blockHash, parentHash); err != nil {
			return 0, fmt.Errorf("insert arch_block %d: %w", blockNum, err)
		}
		batchCount++
	}
	return batchCount, rows.Err()
}

// queryStringColumn returns the first column of every result row as strings.
func queryStringColumn(db *sql.DB, query string) ([]string, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []string
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, rows.Err()
}

// queryInt64Column returns the first column of every result row as int64s.
func queryInt64Column(db *sql.DB, query string) ([]int64, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []int64
	for rows.Next() {
		var v int64
		if err := rows.Scan(&v); err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, rows.Err()
}
