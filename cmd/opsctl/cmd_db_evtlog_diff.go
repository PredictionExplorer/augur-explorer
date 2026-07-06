package main

import (
	"bytes"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
	"github.com/lib/pq"
	"github.com/spf13/cobra"
)

// evtlogRecord is one evt_log row loaded for the diff, keyed by log_rlp.
type evtlogRecord struct {
	blockNum     int64
	txHash       string
	contractAddr string
	topic0Sig    string
	logRLP       []byte
}

// newDBEvtlogDiffCmd builds `opsctl db evtlog-diff`, the replacement for the
// standalone evtlog_diff tool.
func newDBEvtlogDiffCmd() *cobra.Command {
	var (
		primaryConn   string
		secondaryConn string
		limit         int
	)
	cmd := &cobra.Command{
		Use:   "evtlog-diff",
		Short: "Field-level diff of evt_log between two databases (randomwalk contracts)",
		Long: `Loads the RandomWalk-contract event logs from both databases, indexes them
by log_rlp content, and reports records missing from or extra in the
secondary plus any field mismatches on matching records.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runEvtlogDiff(primaryConn, secondaryConn, limit)
		},
	}
	cmd.Flags().StringVar(&primaryConn, "primary", "", "Primary DB connection string (gold standard)")
	cmd.Flags().StringVar(&secondaryConn, "secondary", "", "Secondary DB connection string (to verify)")
	cmd.Flags().IntVar(&limit, "limit", 0, "Limit comparison to first N records (0 = all)")
	_ = cmd.MarkFlagRequired("primary")
	_ = cmd.MarkFlagRequired("secondary")
	return cmd
}

func init() { dbCmd.AddCommand(newDBEvtlogDiffCmd()) }

func runEvtlogDiff(primaryConn, secondaryConn string, limit int) error {
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

	var primaryCount, secondaryCount int64
	_ = primaryDB.QueryRow("SELECT COUNT(*) FROM evt_log WHERE contract_aid = ANY($1)", pq.Array(contractAids)).Scan(&primaryCount)
	_ = secondaryDB.QueryRow("SELECT COUNT(*) FROM evt_log").Scan(&secondaryCount)
	log.Printf("Primary has %d records for our contracts, Secondary has %d records total", primaryCount, secondaryCount)

	log.Println("\n=== Loading events from primary (filtered by contract) ===")
	primaryEvents, err := loadEvtlogRecords(primaryDB, contractAids, limit)
	if err != nil {
		return fmt.Errorf("query primary events: %w", err)
	}
	log.Printf("Loaded %d events from primary", len(primaryEvents))

	log.Println("\n=== Loading events from secondary ===")
	secondaryEvents, err := loadEvtlogRecords(secondaryDB, nil, limit)
	if err != nil {
		return fmt.Errorf("query secondary events: %w", err)
	}
	log.Printf("Loaded %d events from secondary", len(secondaryEvents))

	log.Println("\n=== Comparing events ===")

	primaryByRLP := indexByLogRLP(primaryEvents)
	secondaryByRLP := indexByLogRLP(secondaryEvents)

	// Report the first 20 records missing from the secondary.
	missingCount := 0
	for key, evt := range primaryByRLP {
		if _, found := secondaryByRLP[key]; !found {
			log.Printf("ERROR: Secondary missing event: block=%d tx=%s topic0=%s",
				evt.blockNum, evt.txHash, evt.topic0Sig)
			missingCount++
			if missingCount >= 20 {
				log.Printf("... (showing first 20 missing)")
				break
			}
		}
	}

	// Report the first 20 extra records in the secondary.
	extraCount := 0
	for key, evt := range secondaryByRLP {
		if _, found := primaryByRLP[key]; !found {
			log.Printf("ERROR: Secondary has extra event: block=%d tx=%s topic0=%s",
				evt.blockNum, evt.txHash, evt.topic0Sig)
			extraCount++
			if extraCount >= 20 {
				log.Printf("... (showing first 20 extra)")
				break
			}
		}
	}

	// Count the actual totals.
	actualMissing := 0
	for key := range primaryByRLP {
		if _, found := secondaryByRLP[key]; !found {
			actualMissing++
		}
	}
	actualExtra := 0
	for key := range secondaryByRLP {
		if _, found := primaryByRLP[key]; !found {
			actualExtra++
		}
	}

	// Field mismatches on matching events.
	mismatchCount := 0
	for key, pri := range primaryByRLP {
		sec, found := secondaryByRLP[key]
		if !found {
			continue
		}

		var mismatches []string
		if pri.blockNum != sec.blockNum {
			mismatches = append(mismatches, fmt.Sprintf("block_num: %d vs %d", pri.blockNum, sec.blockNum))
		}
		if pri.txHash != sec.txHash {
			mismatches = append(mismatches, fmt.Sprintf("tx_hash: %s vs %s", pri.txHash, sec.txHash))
		}
		if pri.topic0Sig != sec.topic0Sig {
			mismatches = append(mismatches, fmt.Sprintf("topic0_sig: %s vs %s", pri.topic0Sig, sec.topic0Sig))
		}
		if !bytes.Equal(pri.logRLP, sec.logRLP) {
			mismatches = append(mismatches, fmt.Sprintf("log_rlp: len %d vs %d", len(pri.logRLP), len(sec.logRLP)))
		}

		if len(mismatches) > 0 {
			log.Printf("ERROR: Mismatch: block=%d tx=%s: %v", pri.blockNum, pri.txHash, mismatches)
			mismatchCount++
		}
	}

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
	return nil
}

// loadEvtlogRecords loads evt_log rows joined with transaction and address.
// A nil contractAids loads every row (secondary DB); otherwise rows are
// filtered by contract (primary DB). limit > 0 caps the result.
func loadEvtlogRecords(db *sql.DB, contractAids []int64, limit int) ([]evtlogRecord, error) {
	limitClause := ""
	if limit > 0 {
		limitClause = fmt.Sprintf(" LIMIT %d", limit)
	}

	var rows *sql.Rows
	var err error
	if contractAids != nil {
		rows, err = db.Query(fmt.Sprintf(`
			SELECT e.block_num, t.tx_hash, a.addr, e.topic0_sig, e.log_rlp
			FROM evt_log e
			JOIN transaction t ON e.tx_id = t.id
			JOIN address a ON e.contract_aid = a.address_id
			WHERE e.contract_aid = ANY($1)
			ORDER BY e.block_num, e.id%s
		`, limitClause), pq.Array(contractAids))
	} else {
		rows, err = db.Query(fmt.Sprintf(`
			SELECT e.block_num, t.tx_hash, a.addr, e.topic0_sig, e.log_rlp
			FROM evt_log e
			JOIN transaction t ON e.tx_id = t.id
			JOIN address a ON e.contract_aid = a.address_id
			ORDER BY e.block_num, e.id%s
		`, limitClause))
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []evtlogRecord
	for rows.Next() {
		var evt evtlogRecord
		if err := rows.Scan(&evt.blockNum, &evt.txHash, &evt.contractAddr, &evt.topic0Sig, &evt.logRLP); err != nil {
			return nil, fmt.Errorf("scan event: %w", err)
		}
		events = append(events, evt)
	}
	return events, rows.Err()
}

// indexByLogRLP builds a content-based lookup keyed by hex-encoded log_rlp.
func indexByLogRLP(events []evtlogRecord) map[string]evtlogRecord {
	byRLP := make(map[string]evtlogRecord, len(events))
	for _, evt := range events {
		byRLP[hex.EncodeToString(evt.logRLP)] = evt
	}
	return byRLP
}
