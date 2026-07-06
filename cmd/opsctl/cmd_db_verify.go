package main

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
	"github.com/lib/pq"
	"github.com/spf13/cobra"
)

// newDBVerifyCmd builds `opsctl db verify`, the replacement for the
// standalone db_verify tool.
func newDBVerifyCmd() *cobra.Command {
	var (
		primaryConn   string
		secondaryConn string
	)
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Compare evt_log / transaction / block between two databases (randomwalk contracts)",
		Long: `Loads the RandomWalk-contract event logs, transactions and blocks from the
primary (gold standard) database and checks that the secondary database holds
exactly the same records.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			primaryDB, err := sql.Open("postgres", primaryConn)
			if err != nil {
				return fmt.Errorf("connect to primary: %w", err)
			}
			defer primaryDB.Close()
			log.Println("Connected to primary database")

			secondaryDB, err := sql.Open("postgres", secondaryConn)
			if err != nil {
				return fmt.Errorf("connect to secondary: %w", err)
			}
			defer secondaryDB.Close()
			log.Println("Connected to secondary database")

			contractAids, err := toolutil.GetContractAddressIds(primaryDB, toolutil.ProjectRandomWalk)
			if err != nil {
				return fmt.Errorf("contract aids: %w", err)
			}
			if len(contractAids) == 0 {
				return errors.New("no contract addresses found in rw_contracts")
			}
			log.Printf("Found %d contract address IDs: %v", len(contractAids), contractAids)

			evtOK, err := compareEventLogs(primaryDB, secondaryDB, contractAids)
			if err != nil {
				return err
			}
			txOK, err := compareTransactions(primaryDB, secondaryDB, contractAids)
			if err != nil {
				return err
			}
			blockOK, err := compareBlocks(primaryDB, secondaryDB, contractAids)
			if err != nil {
				return err
			}

			log.Println("")
			log.Println("=== FINAL SUMMARY ===")
			if evtOK && txOK && blockOK {
				log.Println("✓ All tables match perfectly!")
				return nil
			}
			if !evtOK {
				log.Println("✗ evt_log: MISMATCH")
			}
			if !txOK {
				log.Println("✗ transaction: MISMATCH")
			}
			if !blockOK {
				log.Println("✗ block: MISMATCH")
			}
			return errors.New("verification FAILED")
		},
	}
	cmd.Flags().StringVar(&primaryConn, "primary", "", "Primary DB connection string (production - gold standard)")
	cmd.Flags().StringVar(&secondaryConn, "secondary", "", "Secondary DB connection string (new rwcg)")
	_ = cmd.MarkFlagRequired("primary")
	_ = cmd.MarkFlagRequired("secondary")
	return cmd
}

func init() { dbCmd.AddCommand(newDBVerifyCmd()) }

// eventRecord is one evt_log row reduced to the fields compared across DBs.
type eventRecord struct {
	blockNum  int64
	txHash    string
	logRLPHex string // hex-encoded log_rlp, used as the content-based map key
}

func compareEventLogs(primaryDB, secondaryDB *sql.DB, contractAids []int64) (bool, error) {
	log.Println("")
	log.Println("=== Comparing evt_log tables ===")

	// Primary is filtered by contract (production schema); secondary holds
	// only the project's data, so all rows are loaded.
	log.Println("Loading events from primary (filtered by contract)...")
	primaryEvents, err := loadEventRecords(primaryDB, contractAids)
	if err != nil {
		return false, fmt.Errorf("query evt_log (primary): %w", err)
	}
	log.Printf("Primary: %d events", len(primaryEvents))

	log.Println("Loading events from secondary...")
	secondaryEvents, err := loadEventRecords(secondaryDB, nil)
	if err != nil {
		return false, fmt.Errorf("query evt_log (secondary): %w", err)
	}
	log.Printf("Secondary: %d events", len(secondaryEvents))

	return compareEventSets(primaryEvents, secondaryEvents), nil
}

// loadEventRecords loads evt_log rows keyed by log_rlp content. A nil
// contractAids loads every row (secondary DB); otherwise rows are filtered by
// contract (primary DB).
func loadEventRecords(db *sql.DB, contractAids []int64) (map[string]eventRecord, error) {
	query := `
		SELECT e.block_num, t.tx_hash, e.log_rlp
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
	`
	var rows *sql.Rows
	var err error
	if contractAids != nil {
		rows, err = db.Query(query+" WHERE e.contract_aid = ANY($1)", pq.Array(contractAids))
	} else {
		rows, err = db.Query(query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make(map[string]eventRecord)
	for rows.Next() {
		var rec eventRecord
		var logRLP []byte
		if err := rows.Scan(&rec.blockNum, &rec.txHash, &logRLP); err != nil {
			return nil, fmt.Errorf("scan event: %w", err)
		}
		rec.logRLPHex = hex.EncodeToString(logRLP)
		events[rec.logRLPHex] = rec
	}
	return events, rows.Err()
}

func compareEventSets(primary, secondary map[string]eventRecord) bool {
	missing := 0
	extra := 0

	for key, rec := range primary {
		if _, exists := secondary[key]; !exists {
			missing++
			if missing <= 10 {
				log.Printf("MISSING in secondary: block=%d tx=%s rlp=%s...", rec.blockNum, rec.txHash, key[:16])
			}
		}
	}
	for key, rec := range secondary {
		if _, exists := primary[key]; !exists {
			extra++
			if extra <= 10 {
				log.Printf("EXTRA in secondary: block=%d tx=%s rlp=%s...", rec.blockNum, rec.txHash, key[:16])
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

// txRecord is one transaction row reduced to the fields compared across DBs.
type txRecord struct {
	blockNum int64
	txHash   string
	gasUsed  int64
	numLogs  int
}

func compareTransactions(primaryDB, secondaryDB *sql.DB, contractAids []int64) (bool, error) {
	log.Println("")
	log.Println("=== Comparing transaction tables ===")

	log.Println("Getting transaction hashes from primary events...")
	primaryTxHashes, err := txHashesFromEvents(primaryDB, contractAids)
	if err != nil {
		return false, fmt.Errorf("get tx_hashes: %w", err)
	}
	log.Printf("Primary has %d unique transactions for our contracts", len(primaryTxHashes))

	log.Println("Loading transactions from primary...")
	primaryTxs, err := loadTransactions(primaryDB, primaryTxHashes)
	if err != nil {
		return false, fmt.Errorf("query transactions (primary): %w", err)
	}
	log.Printf("Primary: %d transactions loaded", len(primaryTxs))

	log.Println("Loading transactions from secondary...")
	secondaryTxs, err := loadTransactions(secondaryDB, nil)
	if err != nil {
		return false, fmt.Errorf("query transactions (secondary): %w", err)
	}
	log.Printf("Secondary: %d transactions", len(secondaryTxs))

	return compareTxSets(primaryTxs, secondaryTxs), nil
}

// txHashesFromEvents returns the distinct tx_hash values referenced by the
// contracts' event logs.
func txHashesFromEvents(db *sql.DB, contractAids []int64) ([]string, error) {
	rows, err := db.Query(`
		SELECT DISTINCT t.tx_hash
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
		WHERE e.contract_aid = ANY($1)
	`, pq.Array(contractAids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hashes []string
	for rows.Next() {
		var h string
		if err := rows.Scan(&h); err != nil {
			return nil, fmt.Errorf("scan tx_hash: %w", err)
		}
		hashes = append(hashes, h)
	}
	return hashes, rows.Err()
}

// loadTransactions loads transactions keyed by tx_hash. A nil txHashes loads
// every row (secondary DB); otherwise only the listed hashes (primary DB).
func loadTransactions(db *sql.DB, txHashes []string) (map[string]txRecord, error) {
	txs := make(map[string]txRecord)
	query := `
		SELECT block_num, tx_hash, gas_used, num_logs
		FROM transaction
	`
	var rows *sql.Rows
	var err error
	if txHashes != nil {
		if len(txHashes) == 0 {
			return txs, nil
		}
		rows, err = db.Query(query+" WHERE tx_hash = ANY($1)", pq.Array(txHashes))
	} else {
		rows, err = db.Query(query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rec txRecord
		if err := rows.Scan(&rec.blockNum, &rec.txHash, &rec.gasUsed, &rec.numLogs); err != nil {
			return nil, fmt.Errorf("scan transaction: %w", err)
		}
		txs[rec.txHash] = rec
	}
	return txs, rows.Err()
}

func compareTxSets(primary, secondary map[string]txRecord) bool {
	missing := 0
	extra := 0

	for txHash, rec := range primary {
		if _, exists := secondary[txHash]; !exists {
			missing++
			if missing <= 10 {
				log.Printf("MISSING in secondary: tx=%s block=%d", txHash, rec.blockNum)
			}
		}
	}
	for txHash, rec := range secondary {
		if _, exists := primary[txHash]; !exists {
			extra++
			if extra <= 10 {
				log.Printf("EXTRA in secondary: tx=%s block=%d", txHash, rec.blockNum)
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

// blockRecord is one block row reduced to the fields compared across DBs.
type blockRecord struct {
	blockNum   int64
	blockHash  string
	parentHash string
	numTx      int64
}

func compareBlocks(primaryDB, secondaryDB *sql.DB, contractAids []int64) (bool, error) {
	log.Println("")
	log.Println("=== Comparing block tables ===")

	log.Println("Getting block numbers from primary transactions...")
	primaryBlockNums, err := blockNumsFromTransactions(primaryDB, contractAids)
	if err != nil {
		return false, fmt.Errorf("get block_nums: %w", err)
	}
	log.Printf("Primary has %d unique blocks for our contracts", len(primaryBlockNums))

	log.Println("Loading blocks from primary...")
	primaryBlocks, err := loadBlocks(primaryDB, primaryBlockNums)
	if err != nil {
		return false, fmt.Errorf("query blocks (primary): %w", err)
	}
	log.Printf("Primary: %d blocks loaded", len(primaryBlocks))

	log.Println("Loading blocks from secondary...")
	secondaryBlocks, err := loadBlocks(secondaryDB, nil)
	if err != nil {
		return false, fmt.Errorf("query blocks (secondary): %w", err)
	}
	log.Printf("Secondary: %d blocks", len(secondaryBlocks))

	return compareBlockSets(primaryBlocks, secondaryBlocks), nil
}

// blockNumsFromTransactions returns the distinct block numbers referenced by
// the contracts' event logs.
func blockNumsFromTransactions(db *sql.DB, contractAids []int64) ([]int64, error) {
	rows, err := db.Query(`
		SELECT DISTINCT t.block_num
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
		WHERE e.contract_aid = ANY($1)
	`, pq.Array(contractAids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nums []int64
	for rows.Next() {
		var bn int64
		if err := rows.Scan(&bn); err != nil {
			return nil, fmt.Errorf("scan block_num: %w", err)
		}
		nums = append(nums, bn)
	}
	return nums, rows.Err()
}

// loadBlocks loads blocks keyed by block_hash. A nil blockNums loads every
// row (secondary DB); otherwise only the listed block numbers (primary DB).
func loadBlocks(db *sql.DB, blockNums []int64) (map[string]blockRecord, error) {
	blocks := make(map[string]blockRecord)
	query := `
		SELECT block_num, block_hash, parent_hash, num_tx
		FROM block
	`
	var rows *sql.Rows
	var err error
	if blockNums != nil {
		if len(blockNums) == 0 {
			return blocks, nil
		}
		rows, err = db.Query(query+" WHERE block_num = ANY($1)", pq.Array(blockNums))
	} else {
		rows, err = db.Query(query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rec blockRecord
		if err := rows.Scan(&rec.blockNum, &rec.blockHash, &rec.parentHash, &rec.numTx); err != nil {
			return nil, fmt.Errorf("scan block: %w", err)
		}
		blocks[rec.blockHash] = rec
	}
	return blocks, rows.Err()
}

func compareBlockSets(primary, secondary map[string]blockRecord) bool {
	missing := 0
	extra := 0

	for blockHash, rec := range primary {
		if _, exists := secondary[blockHash]; !exists {
			missing++
			if missing <= 10 {
				log.Printf("MISSING in secondary: block_num=%d hash=%s", rec.blockNum, blockHash)
			}
		}
	}
	for blockHash, rec := range secondary {
		if _, exists := primary[blockHash]; !exists {
			extra++
			if extra <= 10 {
				log.Printf("EXTRA in secondary: block_num=%d hash=%s", rec.blockNum, blockHash)
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
