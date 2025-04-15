package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/PredictionExplorer/augur-explorer/dbs"
)

func main() {
	// Set up logging
	log_dir := filepath.Join(os.TempDir(), "deleter_logs")
	os.MkdirAll(log_dir, 0755)
	log_file := filepath.Join(log_dir, "deleter.log")

	// Create logger
	Info := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	// Initialize storage
	storage := dbs.Connect_to_storage(Info)
	storage.Db_set_schema_name("public")
	storage.Init_log(log_file)
	storage.Log_msg("Log initialized\n")

	// Check if status record exists
	var exists bool
	check_query := `SELECT EXISTS(SELECT 1 FROM d_status WHERE id = 1)`
	err := storage.Db().QueryRow(check_query).Scan(&exists)
	if err != nil {
		fmt.Printf("Error checking status record: %v\n", err)
		os.Exit(1)
	}
	if !exists {
		fmt.Printf("Error: no status record found in d_status table. Please initialize the table first.\n")
		os.Exit(1)
	}

	// Event signatures to exclude (these events will be kept)
	keep_signatures := []string{
		// Add your event signatures here, e.g.:
		// "0x1234...",
	}

	// Prepare the query to fetch events
	// We'll fetch events where topic0_sig is NOT IN our keep list
	query_str := `
		SELECT id, block_num, tx_id, topic0_sig 
		FROM evt_log 
		WHERE topic0_sig NOT IN ($1)
		ORDER BY id
		LIMIT 1000
	`

	// Execute the query
	rows, err := storage.Db().Query(query_str, keep_signatures[0]) // For now, just using first signature as example
	if err != nil {
		fmt.Printf("Error querying database: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	// Process results
	for rows.Next() {
		var (
			row_id       int64
			block_num    int64
			tx_id        int64
			topic0_sig   string
		)
		err := rows.Scan(&row_id, &block_num, &tx_id, &topic0_sig)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			continue
		}
		
		// Print the event details
		fmt.Printf("Found event: ID=%d, Block=%d, TxID=%d, Topic0=%s\n", 
			row_id, block_num, tx_id, topic0_sig)
		
		// Update status in d_status table
		update_query := `
			UPDATE d_status
			SET last_evtlog_id = $1, block_num = $2
			WHERE id = 1
		`
		_, err = storage.Db().Exec(update_query, row_id, block_num)
		if err != nil {
			fmt.Printf("Error updating status: %v\n", err)
			os.Exit(1)
		}
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("Error iterating rows: %v\n", err)
		os.Exit(1)
	}
}
