// freezer-scan scans geth freezer receipts to extract event logs
// matching specified contract addresses and event signatures.
//
// Usage:
//
//	freezer-scan --ancientDir /path/to/mainnet \
//	             --startBlock 0 --endBlock 1000000 \
//	             --contracts 0x123...,0x456... \
//	             --eventSigs 0xabc...,0xdef... \
//	             --out events.jsonl \
//	             --workers 16
//
// Resume: If the output file exists and is not empty, the scan will
// automatically resume from the last block found in the file.
//
// The scanning pipeline lives in internal/freezer/scan; this binary parses
// flags and handles SIGINT/SIGTERM (the scan stops gracefully and merges the
// chunks that completed).
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	freezerscanner "github.com/PredictionExplorer/augur-explorer/internal/freezer"
	"github.com/PredictionExplorer/augur-explorer/internal/freezer/scan"
)

var (
	// Command line flags
	ancientDir   = flag.String("ancientDir", "", "Path to directory containing freezer data (receipts.cidx)")
	receiptsCidx = flag.String("receiptsCidx", "", "Explicit path to receipts.cidx (optional)")
	startBlock   = flag.Uint64("startBlock", 0, "Starting block number")
	endBlock     = flag.Uint64("endBlock", 0, "Ending block number (0 = scan to end)")
	contracts    = flag.String("contracts", "", "Comma-separated list of contract addresses (hex)")
	eventSigs    = flag.String("eventSigs", "", "Comma-separated list of event signature topic0 hashes (hex)")
	outPath      = flag.String("out", "-", "Output file path (- for stdout)")
	errorLogPath = flag.String("errorLog", "", "Path to error log file (errors are appended)")
	includeData  = flag.Bool("includeData", false, "Include full hex data in output (large)")
	bestEffort   = flag.Bool("bestEffort", false, "Continue on decode errors")
	showInfo     = flag.Bool("info", false, "Show information about the freezer data")
	noResume     = flag.Bool("noResume", false, "Don't resume from existing file, start fresh")
	numWorkers   = flag.Int("workers", 0, "Number of parallel chunk workers (0 = auto, typically 16-64)")
	chunkSize    = flag.Uint64("chunkSize", 50000, "Blocks per worker chunk")
)

func main() {
	flag.Parse()
	// SIGINT/SIGTERM cancel the scan; completed chunks still merge.
	ctx, stopSignals := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopSignals()
	if err := run(ctx, os.Stdout); err != nil {
		log.Fatalf("%v", err)
	}
}

// run executes the scan described by the package flags, writing --info and
// stdout-mode records to infoOut.
func run(ctx context.Context, infoOut io.Writer) error {
	if *ancientDir == "" && *receiptsCidx == "" {
		flag.Usage()
		return fmt.Errorf("--ancientDir or --receiptsCidx is required")
	}

	// Determine ancient directory
	dir := *ancientDir
	if dir == "" {
		dir = strings.TrimSuffix(*receiptsCidx, "/receipts.cidx")
		dir = strings.TrimSuffix(dir, "/ancient/receipts.cidx")
	}

	contractSet, err := scan.ParseContracts(*contracts)
	if err != nil {
		return err
	}
	eventSigSet, err := scan.ParseEventSigs(*eventSigs)
	if err != nil {
		return err
	}

	workers := *numWorkers
	log.Printf("Starting freezer-scan with %d parallel workers", workers)

	// Open parallel reader (loads index into memory)
	log.Println("Loading index into memory...")
	loadStart := time.Now()
	reader, err := freezerscanner.NewParallelReader(dir)
	if err != nil {
		return fmt.Errorf("failed to open freezer: %w", err)
	}
	totalBlocks, indexSizeMB, cdatCount := reader.GetIndexStats()
	log.Printf("Index loaded in %v: %d blocks (%.1f MB), %d cdat files",
		time.Since(loadStart), totalBlocks, indexSizeMB, cdatCount)

	if *showInfo {
		scan.PrintInfo(infoOut, reader)
		return nil
	}

	if _, err := scan.Run(ctx, reader, scan.Options{
		StartBlock:   *startBlock,
		EndBlock:     *endBlock,
		Contracts:    contractSet,
		EventSigs:    eventSigSet,
		OutPath:      *outPath,
		Stdout:       infoOut,
		ErrorLogPath: *errorLogPath,
		IncludeData:  *includeData,
		BestEffort:   *bestEffort,
		NoResume:     *noResume,
		Workers:      workers,
		ChunkSize:    *chunkSize,
		Logf:         log.Printf,
	}); err != nil {
		return fmt.Errorf("scan failed: %w", err)
	}
	return nil
}
