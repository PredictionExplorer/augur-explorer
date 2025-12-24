package main

import (
	"database/sql"
	"flag"
	"log"
	"time"

	"github.com/lib/pq"
)

const BATCH_SIZE = 20000

func main() {
	srcConn := flag.String("src", "", "Source DB connection string (production)")
	dstConn := flag.String("dst", "", "Destination DB connection string (dev)")
	flag.Parse()

	if *srcConn == "" || *dstConn == "" {
		log.Fatal("Usage: archive_export -src 'postgres://...' -dst 'postgres://...'")
	}

	// Connect to source
	srcDB, err := sql.Open("postgres", *srcConn)
	if err != nil {
		log.Fatalf("Failed to connect to source: %v", err)
	}
	defer srcDB.Close()
	srcDB.SetMaxOpenConns(2)

	// Connect to destination
	dstDB, err := sql.Open("postgres", *dstConn)
	if err != nil {
		log.Fatalf("Failed to connect to destination: %v", err)
	}
	defer dstDB.Close()
	dstDB.SetMaxOpenConns(2)

	// Get contract address IDs from rw_contracts
	contractAids := getContractAddressIds(srcDB)
	if len(contractAids) == 0 {
		log.Fatal("No contract addresses found in rw_contracts")
	}
	log.Printf("Found %d contract address IDs: %v", len(contractAids), contractAids)

	// Export in order: events -> transactions -> blocks
	// (events reference tx, tx references blocks)
	txIds := exportEventLogs(srcDB, dstDB, contractAids)
	blockNums := exportTransactions(srcDB, dstDB, txIds)
	exportBlocks(srcDB, dstDB, blockNums)

	log.Println("=== Export complete ===")
}

func getContractAddressIds(srcDB *sql.DB) []int64 {
	// Get address IDs for contracts in rw_contracts
	rows, err := srcDB.Query(`
		SELECT a.address_id 
		FROM address a
		JOIN rw_contracts rc ON a.addr = rc.randomwalk_addr OR a.addr = rc.marketplace_addr
	`)
	if err != nil {
		log.Fatalf("Failed to get contract addresses: %v", err)
	}
	defer rows.Close()

	var aids []int64
	for rows.Next() {
		var aid int64
		if err := rows.Scan(&aid); err != nil {
			log.Fatalf("Failed to scan address_id: %v", err)
		}
		aids = append(aids, aid)
	}
	return aids
}

func exportEventLogs(srcDB, dstDB *sql.DB, contractAids []int64) map[int64]bool {
	log.Println("=== Exporting event logs for contracts ===")

	// Count total events for these contracts
	log.Println("Counting events...")
	countStart := time.Now()
	var totalEvents int64
	err := srcDB.QueryRow("SELECT COUNT(*) FROM evt_log WHERE contract_aid = ANY($1)", pq.Array(contractAids)).Scan(&totalEvents)
	if err != nil {
		log.Fatalf("Failed to count events: %v", err)
	}
	log.Printf("Total events to export: %d (counted in %.1f ms)", totalEvents, time.Since(countStart).Seconds()*1000)

	// Check resume point
	var currentId int64
	dstDB.QueryRow("SELECT COALESCE(MAX(evt_id), 0) FROM rw_arch_evtlog").Scan(&currentId)
	if currentId > 0 {
		log.Printf("Resuming from evt_id = %d", currentId)
	}

	// Prepare insert
	insertStmt, err := dstDB.Prepare(`
		INSERT INTO rw_arch_evtlog (block_num, evt_id, tx_hash, contract_addr, topic0_sig, log_rlp)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (evt_id) DO NOTHING
	`)
	if err != nil {
		log.Fatalf("Failed to prepare insert: %v", err)
	}
	defer insertStmt.Close()

	// Track unique tx_ids for transaction export
	txIds := make(map[int64]bool)
	totalInserted := int64(0)
	startTime := time.Now()

	// Query using contract_aid index, ordered by id for batching
	for {
		batchStart := time.Now()

		log.Printf("Querying events where id > %d (limit %d)...", currentId, BATCH_SIZE)

		// Build query with contract filter - use pq.Array for ANY clause
		query := `
			SELECT e.block_num, e.id, e.tx_id, t.tx_hash, a.addr, e.topic0_sig, e.log_rlp
			FROM evt_log e
			JOIN transaction t ON e.tx_id = t.id
			JOIN address a ON e.contract_aid = a.address_id
			WHERE e.contract_aid = ANY($1) AND e.id > $2
			ORDER BY e.id
			LIMIT $3
		`

		rows, err := srcDB.Query(query, pq.Array(contractAids), currentId, BATCH_SIZE)
		if err != nil {
			log.Fatalf("Query failed: %v", err)
		}
		log.Printf("Query returned in %.1f ms, processing rows...", time.Since(batchStart).Seconds()*1000)

		batchCount := 0
		var lastId int64
		for rows.Next() {
			var blockNum, evtId, txId int64
			var txHash, contractAddr, topic0Sig string
			var logRlp []byte

			err = rows.Scan(&blockNum, &evtId, &txId, &txHash, &contractAddr, &topic0Sig, &logRlp)
			if err != nil {
				rows.Close()
				log.Fatalf("Scan failed: %v", err)
			}

			_, err = insertStmt.Exec(blockNum, evtId, txHash, contractAddr, topic0Sig, logRlp)
			if err != nil {
				rows.Close()
				log.Fatalf("Insert failed for evt_id %d: %v", evtId, err)
			}

			txIds[txId] = true
			lastId = evtId
			batchCount++
		}
		rows.Close()

		if batchCount == 0 {
			break // No more records
		}

		currentId = lastId
		totalInserted += int64(batchCount)
		elapsed := time.Since(startTime).Seconds()
		rate := float64(totalInserted) / elapsed
		pct := float64(totalInserted) / float64(totalEvents) * 100

		log.Printf("EventLogs: %d rows (%.1f ms) | Total: %d/%d | %.1f/sec | %.1f%%",
			batchCount,
			time.Since(batchStart).Seconds()*1000,
			totalInserted, totalEvents, rate, pct)
	}

	log.Printf("EventLogs export complete. Total: %d, Unique tx_ids: %d", totalInserted, len(txIds))
	return txIds
}

func exportTransactions(srcDB, dstDB *sql.DB, _ map[int64]bool) map[int64]bool {
	log.Println("=== Exporting transactions ===")

	// Get ALL tx_hashes from archived events that don't have a corresponding archived tx
	log.Println("Finding missing transactions in archive...")
	rows, err := srcDB.Query(`
		SELECT DISTINCT t.id
		FROM transaction t
		WHERE t.tx_hash IN (
			SELECT DISTINCT e.tx_hash FROM rw_arch_evtlog e
			LEFT JOIN rw_arch_tx tx ON e.tx_hash = tx.tx_hash
			WHERE tx.tx_hash IS NULL
		)
	`)
	// Note: This query runs on srcDB but uses destination tables via dblink or requires adjustment
	// Actually, we need to query destination for missing tx_hashes, then lookup ids in source

	// Close that query attempt - we need a different approach
	if rows != nil {
		rows.Close()
	}

	// Get all tx_hashes from destination that are missing in rw_arch_tx
	log.Println("Querying destination for tx_hashes missing from rw_arch_tx...")
	missingRows, err := dstDB.Query(`
		SELECT DISTINCT e.tx_hash
		FROM rw_arch_evtlog e
		LEFT JOIN rw_arch_tx tx ON e.tx_hash = tx.tx_hash
		WHERE tx.tx_hash IS NULL
	`)
	if err != nil {
		log.Fatalf("Failed to find missing transactions: %v", err)
	}

	var missingTxHashes []string
	for missingRows.Next() {
		var txHash string
		if err := missingRows.Scan(&txHash); err != nil {
			missingRows.Close()
			log.Fatalf("Scan failed: %v", err)
		}
		missingTxHashes = append(missingTxHashes, txHash)
	}
	missingRows.Close()

	log.Printf("Transactions to export: %d", len(missingTxHashes))
	if len(missingTxHashes) == 0 {
		log.Println("No missing transactions to export")
		// Still need to return block_nums from existing transactions
		return getBlockNumsFromArchiveTx(dstDB)
	}

	// Prepare insert
	insertStmt, err := dstDB.Prepare(`
		INSERT INTO rw_arch_tx (block_num, from_aid, to_aid, gas_used, tx_index, num_logs, ctrct_create, value, gas_price, tx_hash, input_sig)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (tx_hash) DO NOTHING
	`)
	if err != nil {
		log.Fatalf("Failed to prepare insert: %v", err)
	}
	defer insertStmt.Close()

	// Track unique block_nums for block export
	blockNums := make(map[int64]bool)
	totalInserted := int64(0)
	startTime := time.Now()

	// Use tx_hashes directly
	txHashSlice := missingTxHashes

	// Process in batches
	for i := 0; i < len(txHashSlice); i += BATCH_SIZE {
		batchStart := time.Now()
		end := i + BATCH_SIZE
		if end > len(txHashSlice) {
			end = len(txHashSlice)
		}
		batch := txHashSlice[i:end]

		log.Printf("Querying %d transactions...", len(batch))

		rows, err := srcDB.Query(`
			SELECT block_num, from_aid, to_aid, gas_used, tx_index, num_logs, ctrct_create, value, gas_price, tx_hash, input_sig
			FROM transaction
			WHERE tx_hash = ANY($1)
		`, pq.Array(batch))
		if err != nil {
			log.Fatalf("Query failed: %v", err)
		}
		log.Printf("Query returned in %.1f ms", time.Since(batchStart).Seconds()*1000)

		batchCount := 0
		for rows.Next() {
			var blockNum, fromAid, toAid, gasUsed int64
			var txIndex, numLogs int
			var ctrctCreate bool
			var value, gasPrice, txHash string
			var inputSig sql.NullString

			err = rows.Scan(&blockNum, &fromAid, &toAid, &gasUsed, &txIndex, &numLogs, &ctrctCreate, &value, &gasPrice, &txHash, &inputSig)
			if err != nil {
				rows.Close()
				log.Fatalf("Scan failed: %v", err)
			}

			var inputSigVal interface{}
			if inputSig.Valid {
				inputSigVal = inputSig.String
			}

			_, err = insertStmt.Exec(blockNum, fromAid, toAid, gasUsed, txIndex, numLogs, ctrctCreate, value, gasPrice, txHash, inputSigVal)
			if err != nil {
				rows.Close()
				log.Fatalf("Insert failed for tx %s: %v", txHash, err)
			}

			blockNums[blockNum] = true
			batchCount++
		}
		rows.Close()

		totalInserted += int64(batchCount)
		elapsed := time.Since(startTime).Seconds()
		rate := float64(totalInserted) / elapsed
		pct := float64(i+batchCount) / float64(len(txHashSlice)) * 100

		log.Printf("Transactions: %d rows (%.1f ms) | Total: %d/%d | %.1f/sec | %.1f%%",
			batchCount,
			time.Since(batchStart).Seconds()*1000,
			totalInserted, len(txHashSlice), rate, pct)
	}

	log.Printf("Transactions export complete. Total: %d, Unique blocks: %d", totalInserted, len(blockNums))

	// Merge with any existing block_nums from already-exported transactions
	existingBlockNums := getBlockNumsFromArchiveTx(dstDB)
	for bn := range existingBlockNums {
		blockNums[bn] = true
	}

	return blockNums
}

func getBlockNumsFromArchiveTx(dstDB *sql.DB) map[int64]bool {
	blockNums := make(map[int64]bool)
	rows, err := dstDB.Query("SELECT DISTINCT block_num FROM rw_arch_tx")
	if err != nil {
		log.Printf("Warning: Failed to get existing block_nums: %v", err)
		return blockNums
	}
	defer rows.Close()

	for rows.Next() {
		var bn int64
		if err := rows.Scan(&bn); err == nil {
			blockNums[bn] = true
		}
	}
	return blockNums
}

func exportBlocks(srcDB, dstDB *sql.DB, _ map[int64]bool) {
	log.Println("=== Exporting blocks ===")

	// Get all block_nums from archived transactions that don't have a corresponding archived block
	log.Println("Finding missing blocks in archive...")
	missingRows, err := dstDB.Query(`
		SELECT DISTINCT tx.block_num
		FROM rw_arch_tx tx
		LEFT JOIN rw_arch_block b ON tx.block_num = b.block_num
		WHERE b.block_num IS NULL
	`)
	if err != nil {
		log.Fatalf("Failed to find missing blocks: %v", err)
	}

	var missingBlockNums []int64
	for missingRows.Next() {
		var bn int64
		if err := missingRows.Scan(&bn); err != nil {
			missingRows.Close()
			log.Fatalf("Scan failed: %v", err)
		}
		missingBlockNums = append(missingBlockNums, bn)
	}
	missingRows.Close()

	log.Printf("Blocks to export: %d", len(missingBlockNums))
	if len(missingBlockNums) == 0 {
		log.Println("No missing blocks to export")
		return
	}

	// Prepare insert
	insertStmt, err := dstDB.Prepare(`
		INSERT INTO rw_arch_block (block_num, num_tx, ts, cash_flow, block_hash, parent_hash)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (block_hash) DO NOTHING
	`)
	if err != nil {
		log.Fatalf("Failed to prepare insert: %v", err)
	}
	defer insertStmt.Close()

	// Use missingBlockNums directly
	blockNumSlice := missingBlockNums

	totalInserted := int64(0)
	startTime := time.Now()

	// Process in batches
	for i := 0; i < len(blockNumSlice); i += BATCH_SIZE {
		batchStart := time.Now()
		end := i + BATCH_SIZE
		if end > len(blockNumSlice) {
			end = len(blockNumSlice)
		}
		batch := blockNumSlice[i:end]

		log.Printf("Querying %d blocks...", len(batch))

		rows, err := srcDB.Query(`
			SELECT block_num, num_tx, ts, cash_flow, block_hash, parent_hash
			FROM block
			WHERE block_num = ANY($1)
		`, pq.Array(batch))
		if err != nil {
			log.Fatalf("Query failed: %v", err)
		}
		log.Printf("Query returned in %.1f ms", time.Since(batchStart).Seconds()*1000)

		batchCount := 0
		for rows.Next() {
			var blockNum, numTx int64
			var ts time.Time
			var cashFlow, blockHash, parentHash string

			err = rows.Scan(&blockNum, &numTx, &ts, &cashFlow, &blockHash, &parentHash)
			if err != nil {
				rows.Close()
				log.Fatalf("Scan failed: %v", err)
			}

			_, err = insertStmt.Exec(blockNum, numTx, ts, cashFlow, blockHash, parentHash)
			if err != nil {
				rows.Close()
				log.Fatalf("Insert failed for block %d: %v", blockNum, err)
			}

			batchCount++
		}
		rows.Close()

		totalInserted += int64(batchCount)
		elapsed := time.Since(startTime).Seconds()
		rate := float64(totalInserted) / elapsed
		pct := float64(i+batchCount) / float64(len(blockNumSlice)) * 100

		log.Printf("Blocks: %d rows (%.1f ms) | Total: %d/%d | %.1f/sec | %.1f%%",
			batchCount,
			time.Since(batchStart).Seconds()*1000,
			totalInserted, len(blockNumSlice), rate, pct)
	}

	log.Printf("Blocks export complete. Total: %d", totalInserted)
}

