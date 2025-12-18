package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type EventLog struct {
	ID          int64
	BlockNum    int64
	TxID        int64
	ContractAid int64
	Topic0Sig   string
	LogIndex    int
	LogRlp      []byte
}

func main() {
	primary := flag.String("primary", "", "Primary DB connection string (gold standard)")
	secondary := flag.String("secondary", "", "Secondary DB connection string (to verify)")
	table := flag.String("table", "evt_log", "Table to compare (default: evt_log)")
	limit := flag.Int("limit", 0, "Limit comparison to first N records (0 = all)")
	flag.Parse()

	if *primary == "" || *secondary == "" {
		log.Fatal("Usage: evtlog_diff -primary 'postgres://...' -secondary 'postgres://...' [-table evt_log] [-limit N]")
	}

	// Connect to primary
	primaryDB, err := sql.Open("postgres", *primary)
	if err != nil {
		log.Fatalf("Failed to connect to primary: %v", err)
	}
	defer primaryDB.Close()
	log.Println("Connected to primary database")

	// Connect to secondary
	secondaryDB, err := sql.Open("postgres", *secondary)
	if err != nil {
		log.Fatalf("Failed to connect to secondary: %v", err)
	}
	defer secondaryDB.Close()
	log.Println("Connected to secondary database")

	// Get counts
	var primaryCount, secondaryCount int64
	primaryDB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", *table)).Scan(&primaryCount)
	secondaryDB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", *table)).Scan(&secondaryCount)
	log.Printf("Primary has %d records, Secondary has %d records", primaryCount, secondaryCount)

	// Step 1: Check for records in primary that are missing in secondary
	log.Println("\n=== Step 1: Checking for missing records in secondary ===")
	missingCount := checkMissingRecords(primaryDB, secondaryDB, *table, *limit)

	// Step 2: Check for records in secondary that are not in primary
	log.Println("\n=== Step 2: Checking for extra records in secondary ===")
	extraCount := checkExtraRecords(primaryDB, secondaryDB, *table, *limit)

	// Step 3: Check for field mismatches on common records
	log.Println("\n=== Step 3: Checking for field mismatches ===")
	mismatchCount := checkFieldMismatches(primaryDB, secondaryDB, *table, *limit)

	// Summary
	log.Println("\n=== SUMMARY ===")
	log.Printf("Missing in secondary: %d records", missingCount)
	log.Printf("Extra in secondary: %d records", extraCount)
	log.Printf("Field mismatches: %d records", mismatchCount)

	if missingCount == 0 && extraCount == 0 && mismatchCount == 0 {
		log.Println("✓ Databases match perfectly!")
	} else {
		log.Println("✗ Databases have differences!")
	}
}

func checkMissingRecords(primaryDB, secondaryDB *sql.DB, table string, limit int) int {
	// Get all IDs from primary
	limitClause := ""
	if limit > 0 {
		limitClause = fmt.Sprintf(" LIMIT %d", limit)
	}

	rows, err := primaryDB.Query(fmt.Sprintf("SELECT id FROM %s ORDER BY id%s", table, limitClause))
	if err != nil {
		log.Fatalf("Failed to query primary IDs: %v", err)
	}
	defer rows.Close()

	var primaryIDs []int64
	for rows.Next() {
		var id int64
		rows.Scan(&id)
		primaryIDs = append(primaryIDs, id)
	}
	log.Printf("Checking %d records from primary...", len(primaryIDs))

	// Check which IDs exist in secondary (batch query)
	missingCount := 0
	batchSize := 1000
	for i := 0; i < len(primaryIDs); i += batchSize {
		end := i + batchSize
		if end > len(primaryIDs) {
			end = len(primaryIDs)
		}
		batch := primaryIDs[i:end]

		// Query secondary for these IDs
		secRows, err := secondaryDB.Query(
			fmt.Sprintf("SELECT id FROM %s WHERE id = ANY($1)", table),
			pq.Array(batch),
		)
		if err != nil {
			log.Fatalf("Failed to query secondary: %v", err)
		}

		foundIDs := make(map[int64]bool)
		for secRows.Next() {
			var id int64
			secRows.Scan(&id)
			foundIDs[id] = true
		}
		secRows.Close()

		// Report missing
		for _, id := range batch {
			if !foundIDs[id] {
				log.Printf("ERROR: Secondary DB is missing record id=%d", id)
				missingCount++
			}
		}
	}

	return missingCount
}

func checkExtraRecords(primaryDB, secondaryDB *sql.DB, table string, limit int) int {
	// Get all IDs from secondary
	limitClause := ""
	if limit > 0 {
		limitClause = fmt.Sprintf(" LIMIT %d", limit)
	}

	rows, err := secondaryDB.Query(fmt.Sprintf("SELECT id FROM %s ORDER BY id%s", table, limitClause))
	if err != nil {
		log.Fatalf("Failed to query secondary IDs: %v", err)
	}
	defer rows.Close()

	var secondaryIDs []int64
	for rows.Next() {
		var id int64
		rows.Scan(&id)
		secondaryIDs = append(secondaryIDs, id)
	}
	log.Printf("Checking %d records from secondary...", len(secondaryIDs))

	// Check which IDs exist in primary (batch query)
	extraCount := 0
	batchSize := 1000
	for i := 0; i < len(secondaryIDs); i += batchSize {
		end := i + batchSize
		if end > len(secondaryIDs) {
			end = len(secondaryIDs)
		}
		batch := secondaryIDs[i:end]

		// Query primary for these IDs
		priRows, err := primaryDB.Query(
			fmt.Sprintf("SELECT id FROM %s WHERE id = ANY($1)", table),
			pq.Array(batch),
		)
		if err != nil {
			log.Fatalf("Failed to query primary: %v", err)
		}

		foundIDs := make(map[int64]bool)
		for priRows.Next() {
			var id int64
			priRows.Scan(&id)
			foundIDs[id] = true
		}
		priRows.Close()

		// Report extra
		for _, id := range batch {
			if !foundIDs[id] {
				log.Printf("ERROR: Secondary DB has record not in primary id=%d", id)
				extraCount++
			}
		}
	}

	return extraCount
}

func checkFieldMismatches(primaryDB, secondaryDB *sql.DB, table string, limit int) int {
	limitClause := ""
	if limit > 0 {
		limitClause = fmt.Sprintf(" LIMIT %d", limit)
	}

	// Query all records from primary
	rows, err := primaryDB.Query(fmt.Sprintf(`
		SELECT id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp
		FROM %s ORDER BY id%s
	`, table, limitClause))
	if err != nil {
		log.Fatalf("Failed to query primary: %v", err)
	}
	defer rows.Close()

	mismatchCount := 0
	checkedCount := 0

	for rows.Next() {
		var pri EventLog
		err := rows.Scan(&pri.ID, &pri.BlockNum, &pri.TxID, &pri.ContractAid, &pri.Topic0Sig, &pri.LogIndex, &pri.LogRlp)
		if err != nil {
			log.Fatalf("Failed to scan primary row: %v", err)
		}

		// Query corresponding record from secondary
		var sec EventLog
		err = secondaryDB.QueryRow(fmt.Sprintf(`
			SELECT id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp
			FROM %s WHERE id = $1
		`, table), pri.ID).Scan(&sec.ID, &sec.BlockNum, &sec.TxID, &sec.ContractAid, &sec.Topic0Sig, &sec.LogIndex, &sec.LogRlp)

		if err == sql.ErrNoRows {
			// Already reported in missing check
			continue
		}
		if err != nil {
			log.Fatalf("Failed to query secondary for id=%d: %v", pri.ID, err)
		}

		// Compare fields
		mismatches := []string{}

		if pri.BlockNum != sec.BlockNum {
			mismatches = append(mismatches, fmt.Sprintf("block_num: %d vs %d", pri.BlockNum, sec.BlockNum))
		}
		if pri.TxID != sec.TxID {
			mismatches = append(mismatches, fmt.Sprintf("tx_id: %d vs %d", pri.TxID, sec.TxID))
		}
		if pri.ContractAid != sec.ContractAid {
			mismatches = append(mismatches, fmt.Sprintf("contract_aid: %d vs %d", pri.ContractAid, sec.ContractAid))
		}
		if pri.Topic0Sig != sec.Topic0Sig {
			mismatches = append(mismatches, fmt.Sprintf("topic0_sig: %s vs %s", pri.Topic0Sig, sec.Topic0Sig))
		}
		if pri.LogIndex != sec.LogIndex {
			mismatches = append(mismatches, fmt.Sprintf("log_index: %d vs %d", pri.LogIndex, sec.LogIndex))
		}
		if string(pri.LogRlp) != string(sec.LogRlp) {
			mismatches = append(mismatches, fmt.Sprintf("log_rlp: len %d vs %d", len(pri.LogRlp), len(sec.LogRlp)))
		}

		if len(mismatches) > 0 {
			log.Printf("ERROR: Mismatch at id=%d: %v", pri.ID, mismatches)
			mismatchCount++
		}

		checkedCount++
		if checkedCount%10000 == 0 {
			log.Printf("Checked %d records for field mismatches...", checkedCount)
		}
	}

	log.Printf("Checked %d records for field mismatches", checkedCount)
	return mismatchCount
}
