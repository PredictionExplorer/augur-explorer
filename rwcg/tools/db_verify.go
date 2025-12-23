package main

import (
	"database/sql"
	"encoding/hex"
	"flag"
	"log"

	"github.com/lib/pq"
)

func main() {
	primaryConn := flag.String("primary", "", "Primary DB connection string (production - gold standard)")
	secondaryConn := flag.String("secondary", "", "Secondary DB connection string (new rwcg)")
	flag.Parse()

	if *primaryConn == "" || *secondaryConn == "" {
		log.Fatal("Usage: db_verify -primary 'postgres://...' -secondary 'postgres://...'")
	}

	// Connect to databases
	primaryDB, err := sql.Open("postgres", *primaryConn)
	if err != nil {
		log.Fatalf("Failed to connect to primary: %v", err)
	}
	defer primaryDB.Close()
	log.Println("Connected to primary database")

	secondaryDB, err := sql.Open("postgres", *secondaryConn)
	if err != nil {
		log.Fatalf("Failed to connect to secondary: %v", err)
	}
	defer secondaryDB.Close()
	log.Println("Connected to secondary database")

	// Get contract address IDs from primary (production)
	contractAids := getContractAddressIds(primaryDB)
	if len(contractAids) == 0 {
		log.Fatal("No contract addresses found in rw_contracts")
	}
	log.Printf("Found %d contract address IDs: %v", len(contractAids), contractAids)

	// Compare tables
	evtOk := compareEventLogs(primaryDB, secondaryDB, contractAids)
	txOk := compareTransactions(primaryDB, secondaryDB, contractAids)
	blockOk := compareBlocks(primaryDB, secondaryDB, contractAids)

	// Summary
	log.Println("")
	log.Println("=== FINAL SUMMARY ===")
	if evtOk && txOk && blockOk {
		log.Println("✓ All tables match perfectly!")
	} else {
		if !evtOk {
			log.Println("✗ evt_log: MISMATCH")
		}
		if !txOk {
			log.Println("✗ transaction: MISMATCH")
		}
		if !blockOk {
			log.Println("✗ block: MISMATCH")
		}
		log.Fatal("Verification FAILED")
	}
}

func getContractAddressIds(db *sql.DB) []int64 {
	rows, err := db.Query(`
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

// EventRecord represents an event log for comparison
type EventRecord struct {
	BlockNum  int64
	TxHash    string
	LogRlpHex string // hex-encoded for map key
}

func compareEventLogs(primaryDB, secondaryDB *sql.DB, contractAids []int64) bool {
	log.Println("")
	log.Println("=== Comparing evt_log tables ===")

	// Load from primary (filtered by contract) - production schema (no log_index)
	log.Println("Loading events from primary (filtered by contract)...")
	primaryEvents := loadEventLogsPrimary(primaryDB, contractAids)
	log.Printf("Primary: %d events", len(primaryEvents))

	// Load from secondary (all events) - new schema (has log_index)
	log.Println("Loading events from secondary...")
	secondaryEvents := loadEventLogsSecondary(secondaryDB)
	log.Printf("Secondary: %d events", len(secondaryEvents))

	// Compare using log_rlp as unique key
	return compareEventSets(primaryEvents, secondaryEvents)
}

// loadEventLogsPrimary loads from production DB (no log_index column)
func loadEventLogsPrimary(db *sql.DB, contractAids []int64) map[string]EventRecord {
	query := `
		SELECT e.block_num, t.tx_hash, e.log_rlp
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
		WHERE e.contract_aid = ANY($1)
	`
	rows, err := db.Query(query, pq.Array(contractAids))
	if err != nil {
		log.Fatalf("Failed to query evt_log (primary): %v", err)
	}
	defer rows.Close()

	events := make(map[string]EventRecord)
	for rows.Next() {
		var rec EventRecord
		var logRlp []byte
		if err := rows.Scan(&rec.BlockNum, &rec.TxHash, &logRlp); err != nil {
			log.Fatalf("Failed to scan event: %v", err)
		}
		rec.LogRlpHex = hex.EncodeToString(logRlp)
		// Use log_rlp as unique key (content-based)
		events[rec.LogRlpHex] = rec
	}
	return events
}

// loadEventLogsSecondary loads from new rwcg DB (has log_index column)
func loadEventLogsSecondary(db *sql.DB) map[string]EventRecord {
	query := `
		SELECT e.block_num, t.tx_hash, e.log_rlp
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to query evt_log (secondary): %v", err)
	}
	defer rows.Close()

	events := make(map[string]EventRecord)
	for rows.Next() {
		var rec EventRecord
		var logRlp []byte
		if err := rows.Scan(&rec.BlockNum, &rec.TxHash, &logRlp); err != nil {
			log.Fatalf("Failed to scan event: %v", err)
		}
		rec.LogRlpHex = hex.EncodeToString(logRlp)
		// Use log_rlp as unique key (content-based)
		events[rec.LogRlpHex] = rec
	}
	return events
}

func compareEventSets(primary, secondary map[string]EventRecord) bool {
	missing := 0
	extra := 0

	// Check for records missing in secondary
	for key, rec := range primary {
		if _, exists := secondary[key]; !exists {
			missing++
			if missing <= 10 {
				log.Printf("MISSING in secondary: block=%d tx=%s rlp=%s...", rec.BlockNum, rec.TxHash, key[:16])
			}
		}
	}

	// Check for extra records in secondary
	for key, rec := range secondary {
		if _, exists := primary[key]; !exists {
			extra++
			if extra <= 10 {
				log.Printf("EXTRA in secondary: block=%d tx=%s rlp=%s...", rec.BlockNum, rec.TxHash, key[:16])
			}
		}
	}

	if missing > 10 {
		log.Printf("... and %d more missing records", missing-10)
	}
	if extra > 10 {
		log.Printf("... and %d more extra records", extra-10)
	}

	log.Printf("evt_log: Missing=%d, Extra=%d", missing, extra)
	return missing == 0 && extra == 0
}

// TxRecord represents a transaction for comparison
type TxRecord struct {
	BlockNum int64
	TxHash   string
	GasUsed  int64
	NumLogs  int
}

func compareTransactions(primaryDB, secondaryDB *sql.DB, contractAids []int64) bool {
	log.Println("")
	log.Println("=== Comparing transaction tables ===")

	// Get tx_hashes from events for our contracts (primary)
	log.Println("Getting transaction hashes from primary events...")
	primaryTxHashes := getTxHashesFromEvents(primaryDB, contractAids, true)
	log.Printf("Primary has %d unique transactions for our contracts", len(primaryTxHashes))

	// Load transactions from primary
	log.Println("Loading transactions from primary...")
	primaryTxs := loadTransactions(primaryDB, primaryTxHashes)
	log.Printf("Primary: %d transactions loaded", len(primaryTxs))

	// Load all transactions from secondary
	log.Println("Loading transactions from secondary...")
	secondaryTxs := loadAllTransactions(secondaryDB)
	log.Printf("Secondary: %d transactions", len(secondaryTxs))

	// Compare using tx_hash as unique key
	return compareTxSets(primaryTxs, secondaryTxs)
}

func getTxHashesFromEvents(db *sql.DB, contractAids []int64, filterByContract bool) []string {
	var query string
	var rows *sql.Rows
	var err error

	if filterByContract && len(contractAids) > 0 {
		query = `
			SELECT DISTINCT t.tx_hash
			FROM evt_log e
			JOIN transaction t ON e.tx_id = t.id
			WHERE e.contract_aid = ANY($1)
		`
		rows, err = db.Query(query, pq.Array(contractAids))
	} else {
		query = `SELECT DISTINCT tx_hash FROM transaction`
		rows, err = db.Query(query)
	}

	if err != nil {
		log.Fatalf("Failed to get tx_hashes: %v", err)
	}
	defer rows.Close()

	var hashes []string
	for rows.Next() {
		var h string
		if err := rows.Scan(&h); err != nil {
			log.Fatalf("Failed to scan tx_hash: %v", err)
		}
		hashes = append(hashes, h)
	}
	return hashes
}

func loadTransactions(db *sql.DB, txHashes []string) map[string]TxRecord {
	txs := make(map[string]TxRecord)
	if len(txHashes) == 0 {
		return txs
	}

	rows, err := db.Query(`
		SELECT block_num, tx_hash, gas_used, num_logs
		FROM transaction
		WHERE tx_hash = ANY($1)
	`, pq.Array(txHashes))
	if err != nil {
		log.Fatalf("Failed to query transactions: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var rec TxRecord
		if err := rows.Scan(&rec.BlockNum, &rec.TxHash, &rec.GasUsed, &rec.NumLogs); err != nil {
			log.Fatalf("Failed to scan transaction: %v", err)
		}
		txs[rec.TxHash] = rec
	}
	return txs
}

func loadAllTransactions(db *sql.DB) map[string]TxRecord {
	rows, err := db.Query(`
		SELECT block_num, tx_hash, gas_used, num_logs
		FROM transaction
	`)
	if err != nil {
		log.Fatalf("Failed to query transactions: %v", err)
	}
	defer rows.Close()

	txs := make(map[string]TxRecord)
	for rows.Next() {
		var rec TxRecord
		if err := rows.Scan(&rec.BlockNum, &rec.TxHash, &rec.GasUsed, &rec.NumLogs); err != nil {
			log.Fatalf("Failed to scan transaction: %v", err)
		}
		txs[rec.TxHash] = rec
	}
	return txs
}

func compareTxSets(primary, secondary map[string]TxRecord) bool {
	missing := 0
	extra := 0

	// Check for records missing in secondary
	for txHash, rec := range primary {
		if _, exists := secondary[txHash]; !exists {
			missing++
			if missing <= 10 {
				log.Printf("MISSING in secondary: tx=%s block=%d", txHash, rec.BlockNum)
			}
		}
	}

	// Check for extra records in secondary
	for txHash, rec := range secondary {
		if _, exists := primary[txHash]; !exists {
			extra++
			if extra <= 10 {
				log.Printf("EXTRA in secondary: tx=%s block=%d", txHash, rec.BlockNum)
			}
		}
	}

	if missing > 10 {
		log.Printf("... and %d more missing transactions", missing-10)
	}
	if extra > 10 {
		log.Printf("... and %d more extra transactions", extra-10)
	}

	log.Printf("transaction: Missing=%d, Extra=%d", missing, extra)
	return missing == 0 && extra == 0
}

func compareBlocks(primaryDB, secondaryDB *sql.DB, contractAids []int64) bool {
	log.Println("")
	log.Println("=== Comparing block tables ===")

	// Get block_nums from transactions for our contracts (primary)
	log.Println("Getting block numbers from primary transactions...")
	primaryBlockNums := getBlockNumsFromTransactions(primaryDB, contractAids)
	log.Printf("Primary has %d unique blocks for our contracts", len(primaryBlockNums))

	// Load blocks from primary
	log.Println("Loading blocks from primary...")
	primaryBlocks := loadBlocks(primaryDB, primaryBlockNums)
	log.Printf("Primary: %d blocks loaded", len(primaryBlocks))

	// Load all blocks from secondary
	log.Println("Loading blocks from secondary...")
	secondaryBlocks := loadAllBlocks(secondaryDB)
	log.Printf("Secondary: %d blocks", len(secondaryBlocks))

	// Compare using block_hash as unique key
	return compareBlockSets(primaryBlocks, secondaryBlocks)
}

func getBlockNumsFromTransactions(db *sql.DB, contractAids []int64) []int64 {
	query := `
		SELECT DISTINCT t.block_num
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
		WHERE e.contract_aid = ANY($1)
	`
	rows, err := db.Query(query, pq.Array(contractAids))
	if err != nil {
		log.Fatalf("Failed to get block_nums: %v", err)
	}
	defer rows.Close()

	var nums []int64
	for rows.Next() {
		var bn int64
		if err := rows.Scan(&bn); err != nil {
			log.Fatalf("Failed to scan block_num: %v", err)
		}
		nums = append(nums, bn)
	}
	return nums
}

// BlockRecord represents a block for comparison
type BlockRecord struct {
	BlockNum   int64
	BlockHash  string
	ParentHash string
	NumTx      int64
}

func loadBlocks(db *sql.DB, blockNums []int64) map[string]BlockRecord {
	blocks := make(map[string]BlockRecord)
	if len(blockNums) == 0 {
		return blocks
	}

	rows, err := db.Query(`
		SELECT block_num, block_hash, parent_hash, num_tx
		FROM block
		WHERE block_num = ANY($1)
	`, pq.Array(blockNums))
	if err != nil {
		log.Fatalf("Failed to query blocks: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var rec BlockRecord
		if err := rows.Scan(&rec.BlockNum, &rec.BlockHash, &rec.ParentHash, &rec.NumTx); err != nil {
			log.Fatalf("Failed to scan block: %v", err)
		}
		blocks[rec.BlockHash] = rec
	}
	return blocks
}

func loadAllBlocks(db *sql.DB) map[string]BlockRecord {
	rows, err := db.Query(`
		SELECT block_num, block_hash, parent_hash, num_tx
		FROM block
	`)
	if err != nil {
		log.Fatalf("Failed to query blocks: %v", err)
	}
	defer rows.Close()

	blocks := make(map[string]BlockRecord)
	for rows.Next() {
		var rec BlockRecord
		if err := rows.Scan(&rec.BlockNum, &rec.BlockHash, &rec.ParentHash, &rec.NumTx); err != nil {
			log.Fatalf("Failed to scan block: %v", err)
		}
		blocks[rec.BlockHash] = rec
	}
	return blocks
}

func compareBlockSets(primary, secondary map[string]BlockRecord) bool {
	missing := 0
	extra := 0

	// Check for records missing in secondary
	for blockHash, rec := range primary {
		if _, exists := secondary[blockHash]; !exists {
			missing++
			if missing <= 10 {
				log.Printf("MISSING in secondary: block_num=%d hash=%s", rec.BlockNum, blockHash)
			}
		}
	}

	// Check for extra records in secondary
	for blockHash, rec := range secondary {
		if _, exists := primary[blockHash]; !exists {
			extra++
			if extra <= 10 {
				log.Printf("EXTRA in secondary: block_num=%d hash=%s", rec.BlockNum, blockHash)
			}
		}
	}

	if missing > 10 {
		log.Printf("... and %d more missing blocks", missing-10)
	}
	if extra > 10 {
		log.Printf("... and %d more extra blocks", extra-10)
	}

	log.Printf("block: Missing=%d, Extra=%d", missing, extra)
	return missing == 0 && extra == 0
}
