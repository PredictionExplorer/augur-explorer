// Command freezer-verify compares events extracted from a node freezer (the
// JSONL output of freezer-scan) with events stored in a PostgreSQL database.
//
// Usage:
//
//	freezer-verify --input events.jsonl --db "user=cosmicgame dbname=cosmicgame sslmode=disable"
//
// It exits non-zero when fewer than 99% of the database's
// (block, topic0, contract) combinations match the freezer extract.
// The comparison engine lives in internal/freezer/verify.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/freezer/verify"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	passed, err := run(ctx, os.Args[1:], os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "freezer-verify: %v\n", err)
		os.Exit(1)
	}
	if !passed {
		os.Exit(1)
	}
}

// run executes the verification and reports whether it met the match
// threshold.
func run(ctx context.Context, args []string, out, errOut io.Writer) (bool, error) {
	flags := flag.NewFlagSet("freezer-verify", flag.ContinueOnError)
	flags.SetOutput(errOut)
	inputFile := flags.String("input", "", "Path to JSONL file with extracted events")
	dbConnStr := flags.String("db", "user=cosmicgame dbname=cosmicgame sslmode=disable", "PostgreSQL connection string")
	tableName := flags.String("table", "evt_log", "Event log table name")
	verbose := flags.Bool("verbose", false, "Show detailed comparison results")
	maxMissing := flags.Int("maxMissing", 10, "Maximum missing events to display")
	maxExtra := flags.Int("maxExtra", 10, "Maximum extra events to display")
	if err := flags.Parse(args); err != nil {
		return false, err
	}

	if *inputFile == "" {
		return false, fmt.Errorf("--input is required")
	}
	if !verify.ValidTableName(*tableName) {
		return false, fmt.Errorf("invalid table name %q: must be a plain SQL identifier", *tableName)
	}

	logger := slog.New(slog.NewTextHandler(errOut, nil))

	// Load events from the JSONL file.
	logger.Info(fmt.Sprintf("Loading events from %s...", *inputFile))
	file, err := os.Open(filepath.Clean(*inputFile))
	if err != nil {
		return false, fmt.Errorf("opening input: %w", err)
	}
	defer func() { _ = file.Close() }() // best-effort close on read path

	freezerEvents, contracts, err := verify.LoadJSONL(file)
	if err != nil {
		return false, fmt.Errorf("failed to load events: %w", err)
	}
	logger.Info(fmt.Sprintf("Loaded %d distinct (block, topic0, contract) combinations from file", len(freezerEvents)))
	logger.Info(fmt.Sprintf("Contracts found: %v", contracts))

	minBlock, maxBlock, ok := verify.BlockRange(freezerEvents)
	if !ok {
		return false, fmt.Errorf("no events found in input file")
	}
	logger.Info(fmt.Sprintf("Block range in file: %d - %d", minBlock, maxBlock))

	// Connect to the database.
	conn, err := pgx.Connect(ctx, *dbConnStr)
	if err != nil {
		return false, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer func() { _ = conn.Close(ctx) }() // best-effort close on process exit
	logger.Info("Connected to database")

	// Resolve contract address ids.
	contractAddrs := make([]string, 0, len(contracts))
	for c := range contracts {
		contractAddrs = append(contractAddrs, c)
	}
	contractAids, err := verify.ContractAids(ctx, conn, contractAddrs)
	if err != nil {
		return false, fmt.Errorf("failed to query address table: %w", err)
	}
	if len(contractAids) == 0 {
		return false, fmt.Errorf("none of the contracts found in database address table: %v", contractAddrs)
	}
	logger.Info(fmt.Sprintf("Found %d contract address IDs in database", len(contractAids)))

	// Query database events and compare.
	logger.Info(fmt.Sprintf("Querying database for events in block range %d - %d...", minBlock, maxBlock))
	dbEvents, err := verify.DBEvents(ctx, conn, *tableName, minBlock, maxBlock, contractAids)
	if err != nil {
		return false, fmt.Errorf("failed to query events: %w", err)
	}
	logger.Info(fmt.Sprintf("Found %d distinct (block, topic0, contract) combinations in database", len(dbEvents)))

	report := verify.Compare(freezerEvents, dbEvents, *maxMissing, *maxExtra)
	report.Write(out, *verbose)

	return report.Passed(), nil
}
