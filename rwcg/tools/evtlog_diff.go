package main

import (
	"bytes"
	"database/sql"
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type EventLog struct {
	BlockNum    int64
	TxHash      string
	ContractAddr string
	Topic0Sig   string
	LogRlp      []byte
}

func main() {
	primary := flag.String("primary", "", "Primary DB connection string (gold standard)")
	secondary := flag.String("secondary", "", "Secondary DB connection string (to verify)")
	limit := flag.Int("limit", 0, "Limit comparison to first N records (0 = all)")
	flag.Parse()

	if *primary == "" || *secondary == "" {
		log.Fatal("Usage: evtlog_diff -primary 'postgres://...' -secondary 'postgres://...' [-limit N]")
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

	// Get contract address IDs from rw_contracts in primary DB
	contractAids := getContractAddressIds(primaryDB)
	if len(contractAids) == 0 {
		log.Fatal("No contract addresses found in rw_contracts")
	}
	log.Printf("Found %d contract address IDs: %v", len(contractAids), contractAids)

	// Get counts for our contracts only
	var primaryCount, secondaryCount int64
	primaryDB.QueryRow("SELECT COUNT(*) FROM evt_log WHERE contract_aid = ANY($1)", pq.Array(contractAids)).Scan(&primaryCount)
	secondaryDB.QueryRow("SELECT COUNT(*) FROM evt_log").Scan(&secondaryCount)
	log.Printf("Primary has %d records for our contracts, Secondary has %d records total", primaryCount, secondaryCount)

	// Load all events from primary (for our contracts)
	log.Println("\n=== Loading events from primary (filtered by contract) ===")
	primaryEvents := loadPrimaryEvents(primaryDB, contractAids, *limit)
	log.Printf("Loaded %d events from primary", len(primaryEvents))

	// Load all events from secondary
	log.Println("\n=== Loading events from secondary ===")
	secondaryEvents := loadSecondaryEvents(secondaryDB, *limit)
	log.Printf("Loaded %d events from secondary", len(secondaryEvents))

	// Compare: create lookup by (block_num, tx_hash, log_rlp_hash)
	log.Println("\n=== Comparing events ===")

	// Build primary index by log_rlp (the unique content)
	primaryByRlp := make(map[string]EventLog)
	for _, evt := range primaryEvents {
		key := hex.EncodeToString(evt.LogRlp)
		primaryByRlp[key] = evt
	}

	// Build secondary index by log_rlp
	secondaryByRlp := make(map[string]EventLog)
	for _, evt := range secondaryEvents {
		key := hex.EncodeToString(evt.LogRlp)
		secondaryByRlp[key] = evt
	}

	// Check for events in primary missing from secondary
	missingCount := 0
	for key, evt := range primaryByRlp {
		if _, found := secondaryByRlp[key]; !found {
			log.Printf("ERROR: Secondary missing event: block=%d tx=%s topic0=%s",
				evt.BlockNum, evt.TxHash, evt.Topic0Sig)
			missingCount++
			if missingCount >= 20 {
				log.Printf("... (showing first 20 missing)")
				break
			}
		}
	}

	// Check for events in secondary not in primary
	extraCount := 0
	for key, evt := range secondaryByRlp {
		if _, found := primaryByRlp[key]; !found {
			log.Printf("ERROR: Secondary has extra event: block=%d tx=%s topic0=%s",
				evt.BlockNum, evt.TxHash, evt.Topic0Sig)
			extraCount++
			if extraCount >= 20 {
				log.Printf("... (showing first 20 extra)")
				break
			}
		}
	}

	// Count actual totals
	actualMissing := 0
	for key := range primaryByRlp {
		if _, found := secondaryByRlp[key]; !found {
			actualMissing++
		}
	}
	actualExtra := 0
	for key := range secondaryByRlp {
		if _, found := primaryByRlp[key]; !found {
			actualExtra++
		}
	}

	// Check for field mismatches on matching events
	mismatchCount := 0
	for key, pri := range primaryByRlp {
		sec, found := secondaryByRlp[key]
		if !found {
			continue
		}

		mismatches := []string{}

		if pri.BlockNum != sec.BlockNum {
			mismatches = append(mismatches, fmt.Sprintf("block_num: %d vs %d", pri.BlockNum, sec.BlockNum))
		}
		if pri.TxHash != sec.TxHash {
			mismatches = append(mismatches, fmt.Sprintf("tx_hash: %s vs %s", pri.TxHash, sec.TxHash))
		}
		if pri.Topic0Sig != sec.Topic0Sig {
			mismatches = append(mismatches, fmt.Sprintf("topic0_sig: %s vs %s", pri.Topic0Sig, sec.Topic0Sig))
		}
		if !bytes.Equal(pri.LogRlp, sec.LogRlp) {
			mismatches = append(mismatches, fmt.Sprintf("log_rlp: len %d vs %d", len(pri.LogRlp), len(sec.LogRlp)))
		}

		if len(mismatches) > 0 {
			log.Printf("ERROR: Mismatch: block=%d tx=%s: %v", pri.BlockNum, pri.TxHash, mismatches)
			mismatchCount++
		}
	}

	// Summary
	log.Println("\n=== SUMMARY ===")
	log.Printf("Primary events (for our contracts): %d", len(primaryEvents))
	log.Printf("Secondary events: %d", len(secondaryEvents))
	log.Printf("Missing in secondary: %d records", actualMissing)
	log.Printf("Extra in secondary: %d records", actualExtra)
	log.Printf("Field mismatches: %d records", mismatchCount)

	if actualMissing == 0 && actualExtra == 0 && mismatchCount == 0 {
		log.Println("✓ Databases match perfectly!")
	} else {
		log.Println("✗ Databases have differences!")
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

func loadPrimaryEvents(db *sql.DB, contractAids []int64, limit int) []EventLog {
	limitClause := ""
	if limit > 0 {
		limitClause = fmt.Sprintf(" LIMIT %d", limit)
	}

	// Primary DB uses contract_aid, join with address and transaction
	rows, err := db.Query(fmt.Sprintf(`
		SELECT e.block_num, t.tx_hash, a.addr, e.topic0_sig, e.log_rlp
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
		JOIN address a ON e.contract_aid = a.address_id
		WHERE e.contract_aid = ANY($1)
		ORDER BY e.block_num, e.id%s
	`, limitClause), pq.Array(contractAids))
	if err != nil {
		log.Fatalf("Failed to query primary events: %v", err)
	}
	defer rows.Close()

	var events []EventLog
	for rows.Next() {
		var evt EventLog
		err := rows.Scan(&evt.BlockNum, &evt.TxHash, &evt.ContractAddr, &evt.Topic0Sig, &evt.LogRlp)
		if err != nil {
			log.Fatalf("Failed to scan primary event: %v", err)
		}
		events = append(events, evt)
	}
	return events
}

func loadSecondaryEvents(db *sql.DB, limit int) []EventLog {
	limitClause := ""
	if limit > 0 {
		limitClause = fmt.Sprintf(" LIMIT %d", limit)
	}

	// Secondary DB - join with address and transaction
	rows, err := db.Query(fmt.Sprintf(`
		SELECT e.block_num, t.tx_hash, a.addr, e.topic0_sig, e.log_rlp
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
		JOIN address a ON e.contract_aid = a.address_id
		ORDER BY e.block_num, e.id%s
	`, limitClause))
	if err != nil {
		log.Fatalf("Failed to query secondary events: %v", err)
	}
	defer rows.Close()

	var events []EventLog
	for rows.Next() {
		var evt EventLog
		err := rows.Scan(&evt.BlockNum, &evt.TxHash, &evt.ContractAddr, &evt.Topic0Sig, &evt.LogRlp)
		if err != nil {
			log.Fatalf("Failed to scan secondary event: %v", err)
		}
		events = append(events, evt)
	}
	return events
}
