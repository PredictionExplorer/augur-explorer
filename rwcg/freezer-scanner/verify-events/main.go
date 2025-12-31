// verify-events compares events from a JSONL file with events stored in a PostgreSQL database.
//
// Usage:
//
//	verify-events --input events.jsonl --db "user=cosmicgame dbname=cosmicgame sslmode=disable"
package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lib/pq"
)

var (
	inputFile  = flag.String("input", "", "Path to JSONL file with extracted events")
	dbConnStr  = flag.String("db", "user=cosmicgame dbname=cosmicgame sslmode=disable", "PostgreSQL connection string")
	tableName  = flag.String("table", "evt_log", "Event log table name")
	verbose    = flag.Bool("verbose", false, "Show detailed comparison results")
	maxMissing = flag.Int("maxMissing", 10, "Maximum missing events to display")
	maxExtra   = flag.Int("maxExtra", 10, "Maximum extra events to display")
)

// LogRecord matches the output format from freezer-scan
type LogRecord struct {
	BlockNumber  uint64   `json:"blockNumber"`
	TxIndex      uint     `json:"txIndex"`
	ReceiptIndex uint     `json:"receiptIndex"`
	LogIndex     uint     `json:"logIndex"`
	Contract     string   `json:"contract"`
	Topic0       string   `json:"topic0"`
	Topics       []string `json:"topics"`
	DataKeccak   string   `json:"dataKeccak"`
	DataLen      int      `json:"dataLen"`
	DataHex      string   `json:"dataHex,omitempty"`
}

type eventKey struct {
	blockNum int64
	topic0   string // first 8 chars (4 bytes) of topic0
	contract string
}

func main() {
	flag.Parse()

	if *inputFile == "" {
		fmt.Fprintln(os.Stderr, "Error: --input is required")
		flag.Usage()
		os.Exit(1)
	}

	// Connect to database
	db, err := sql.Open("postgres", *dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Printf("Connected to database")

	// Load events from JSONL file
	log.Printf("Loading events from %s...", *inputFile)
	freezerEvents, contracts, err := loadJSONLEvents(*inputFile)
	if err != nil {
		log.Fatalf("Failed to load events: %v", err)
	}
	log.Printf("Loaded %d distinct (block, topic0, contract) combinations from file", len(freezerEvents))
	log.Printf("Contracts found: %v", contracts)

	if len(freezerEvents) == 0 {
		log.Fatal("No events found in input file")
	}

	// Get block range from freezer events
	var minBlock, maxBlock int64 = 1<<62, 0
	for key := range freezerEvents {
		if key.blockNum < minBlock {
			minBlock = key.blockNum
		}
		if key.blockNum > maxBlock {
			maxBlock = key.blockNum
		}
	}
	log.Printf("Block range in file: %d - %d", minBlock, maxBlock)

	// Get contract address IDs from database
	contractAddrs := make([]string, 0, len(contracts))
	for c := range contracts {
		contractAddrs = append(contractAddrs, c)
	}

	var contractAids []int64
	rows, err := db.Query(`
		SELECT address_id FROM address 
		WHERE LOWER(addr) = ANY($1)
	`, pq.Array(contractAddrs))
	if err != nil {
		log.Fatalf("Failed to query address table: %v", err)
	}
	for rows.Next() {
		var aid int64
		if err := rows.Scan(&aid); err != nil {
			log.Fatalf("Failed to scan address: %v", err)
		}
		contractAids = append(contractAids, aid)
	}
	rows.Close()

	if len(contractAids) == 0 {
		log.Fatalf("None of the contracts found in database address table: %v", contractAddrs)
	}
	log.Printf("Found %d contract address IDs in database", len(contractAids))

	// Query database for events
	log.Printf("Querying database for events in block range %d - %d...", minBlock, maxBlock)
	dbRows, err := db.Query(`
		SELECT e.block_num, e.topic0_sig, a.addr, COUNT(*) as cnt
		FROM `+*tableName+` e
		JOIN address a ON e.contract_aid = a.address_id
		WHERE e.block_num BETWEEN $1 AND $2
		  AND e.contract_aid = ANY($3)
		GROUP BY e.block_num, e.topic0_sig, a.addr
		ORDER BY e.block_num
	`, minBlock, maxBlock, pq.Array(contractAids))
	if err != nil {
		log.Fatalf("Failed to query events: %v", err)
	}
	defer dbRows.Close()

	dbEvents := make(map[eventKey]int)
	for dbRows.Next() {
		var blockNum int64
		var topic0, contract string
		var cnt int
		if err := dbRows.Scan(&blockNum, &topic0, &contract, &cnt); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		key := eventKey{
			blockNum: blockNum,
			topic0:   strings.ToLower(topic0),
			contract: strings.ToLower(contract),
		}
		dbEvents[key] = cnt
	}
	log.Printf("Found %d distinct (block, topic0, contract) combinations in database", len(dbEvents))

	// Compare
	var match, missing, extra int
	var missingDetails, extraDetails []string

	for key, dbCount := range dbEvents {
		if fCount, ok := freezerEvents[key]; ok {
			if fCount == dbCount {
				match++
			} else {
				// Count mismatch
				if *verbose {
					log.Printf("Count mismatch: block %d topic0 %s contract %s: db=%d freezer=%d",
						key.blockNum, key.topic0, key.contract, dbCount, fCount)
				}
			}
		} else {
			missing++
			if len(missingDetails) < *maxMissing {
				missingDetails = append(missingDetails, 
					fmt.Sprintf("block=%d topic0=%s contract=%s (db count=%d)",
						key.blockNum, key.topic0, key.contract, dbCount))
			}
		}
	}

	for key, fCount := range freezerEvents {
		if _, ok := dbEvents[key]; !ok {
			extra++
			if len(extraDetails) < *maxExtra {
				extraDetails = append(extraDetails,
					fmt.Sprintf("block=%d topic0=%s contract=%s (freezer count=%d)",
						key.blockNum, key.topic0, key.contract, fCount))
			}
		}
	}

	// Report results
	fmt.Println()
	fmt.Println("=== Verification Results ===")
	fmt.Printf("Freezer events:  %d distinct (block, topic0, contract)\n", len(freezerEvents))
	fmt.Printf("Database events: %d distinct (block, topic0, contract)\n", len(dbEvents))
	fmt.Println()
	fmt.Printf("Matching:          %d\n", match)
	fmt.Printf("Missing (in DB, not in freezer): %d\n", missing)
	fmt.Printf("Extra (in freezer, not in DB):   %d\n", extra)

	total := len(dbEvents)
	if total > 0 {
		matchRate := float64(match) / float64(total) * 100
		fmt.Printf("\nMatch rate: %.2f%%\n", matchRate)
	}

	if len(missingDetails) > 0 {
		fmt.Println("\nMissing from freezer (sample):")
		for _, d := range missingDetails {
			fmt.Printf("  %s\n", d)
		}
		if missing > len(missingDetails) {
			fmt.Printf("  ... and %d more\n", missing-len(missingDetails))
		}
	}

	if len(extraDetails) > 0 {
		fmt.Println("\nExtra in freezer (sample):")
		for _, d := range extraDetails {
			fmt.Printf("  %s\n", d)
		}
		if extra > len(extraDetails) {
			fmt.Printf("  ... and %d more\n", extra-len(extraDetails))
		}
	}

	// Exit with error if match rate is below threshold
	if total > 0 && float64(match)/float64(total) < 0.99 {
		os.Exit(1)
	}
}

func loadJSONLEvents(path string) (map[eventKey]int, map[string]bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	events := make(map[eventKey]int)
	contracts := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	
	// Increase buffer size for long lines
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if line == "" {
			continue
		}

		var record LogRecord
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			return nil, nil, fmt.Errorf("line %d: invalid JSON: %w", lineNum, err)
		}

		// Extract first 8 chars of topic0 (matching DB format)
		topic0 := strings.ToLower(record.Topic0)
		if strings.HasPrefix(topic0, "0x") && len(topic0) >= 10 {
			topic0 = topic0[2:10] // Get first 4 bytes (8 hex chars)
		}

		contract := strings.ToLower(record.Contract)
		contracts[contract] = true

		key := eventKey{
			blockNum: int64(record.BlockNumber),
			topic0:   topic0,
			contract: contract,
		}
		events[key]++
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("scanner error: %w", err)
	}

	return events, contracts, nil
}

