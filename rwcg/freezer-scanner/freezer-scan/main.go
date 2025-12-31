// freezer-scan scans geth freezer receipts to extract event logs
// matching specified contract addresses and event signatures.
//
// Usage:
//
//	freezer-scan --ancientDir /path/to/mainnet \
//	             --startBlock 0 --endBlock 1000000 \
//	             --contracts 0x123...,0x456... \
//	             --eventSigs 0xabc...,0xdef... \
//	             --out events.jsonl --format jsonl
//
// Resume: If the output file exists and is not empty, the scan will
// automatically resume from the last block found in the file.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/rwcg/freezer-scanner/decode"
	freezerscanner "github.com/PredictionExplorer/augur-explorer/rwcg/freezer-scanner"
	"github.com/PredictionExplorer/augur-explorer/rwcg/freezer-scanner/output"
)

var (
	// Command line flags
	ancientDir      = flag.String("ancientDir", "", "Path to directory containing freezer data (receipts.cidx)")
	receiptsCidx    = flag.String("receiptsCidx", "", "Explicit path to receipts.cidx (optional)")
	receiptsCdatDir = flag.String("receiptsCdatDir", "", "Explicit path to folder containing receipts.*.cdat (optional)")
	startBlock      = flag.Uint64("startBlock", 0, "Starting block number")
	endBlock        = flag.Uint64("endBlock", 0, "Ending block number (0 = scan to end)")
	contracts       = flag.String("contracts", "", "Comma-separated list of contract addresses (hex)")
	eventSigs       = flag.String("eventSigs", "", "Comma-separated list of event signature topic0 hashes (hex)")
	outPath         = flag.String("out", "-", "Output file path (- for stdout)")
	errorLogPath    = flag.String("errorLog", "", "Path to error log file (errors are appended)")
	format          = flag.String("format", "jsonl", "Output format: jsonl or csv")
	includeData     = flag.Bool("includeData", false, "Include full hex data in output (large)")
	progressEvery   = flag.Uint64("progressEvery", 10000, "Log progress every N blocks")
	bestEffort      = flag.Bool("bestEffort", false, "Continue on decode errors")
	validate        = flag.Bool("validate", false, "Validate index before scanning")
	validateOnly    = flag.Bool("validateOnly", false, "Only validate index, don't scan")
	showInfo        = flag.Bool("info", false, "Show information about the freezer data")
	noResume        = flag.Bool("noResume", false, "Don't resume from existing file, start fresh")
)

func main() {
	flag.Parse()

	// Validate required flags
	if *ancientDir == "" && *receiptsCidx == "" {
		fmt.Fprintln(os.Stderr, "Error: --ancientDir or --receiptsCidx is required")
		flag.Usage()
		os.Exit(1)
	}

	// Determine ancient directory
	dir := *ancientDir
	if dir == "" {
		// Derive from cidx path
		dir = strings.TrimSuffix(*receiptsCidx, "/receipts.cidx")
		dir = strings.TrimSuffix(dir, "/ancient/receipts.cidx")
	}

	// Open freezer reader
	reader, err := freezerscanner.NewFreezerReader(dir)
	if err != nil {
		log.Fatalf("Failed to open freezer: %v", err)
	}
	defer reader.Close()

	// Show info if requested
	if *showInfo {
		printInfo(reader)
		if !*validate && !*validateOnly {
			return
		}
	}

	// Validate if requested
	if *validate || *validateOnly {
		if err := validateIndex(reader); err != nil {
			log.Fatalf("Validation failed: %v", err)
		}
		if *validateOnly {
			return
		}
	}

	// Parse contract addresses
	contractSet := make(map[common.Address]bool)
	if *contracts != "" {
		for _, addr := range strings.Split(*contracts, ",") {
			addr = strings.TrimSpace(addr)
			if addr == "" {
				continue
			}
			if !common.IsHexAddress(addr) {
				log.Fatalf("Invalid contract address: %s", addr)
			}
			contractSet[common.HexToAddress(addr)] = true
		}
	}

	// Parse event signatures
	eventSigSet := make(map[common.Hash]bool)
	if *eventSigs != "" {
		for _, sig := range strings.Split(*eventSigs, ",") {
			sig = strings.TrimSpace(sig)
			if sig == "" {
				continue
			}
			// Remove 0x prefix if present
			sig = strings.TrimPrefix(sig, "0x")
			if len(sig) != 64 {
				log.Fatalf("Invalid event signature (must be 32 bytes hex): %s", sig)
			}
			eventSigSet[common.HexToHash(sig)] = true
		}
	}

	// Determine start block (check for resume)
	actualStart := *startBlock
	appendMode := false
	
	if *outPath != "-" && !*noResume {
		if resumeBlock, found := getResumeBlock(*outPath); found {
			if resumeBlock >= actualStart {
			actualStart = resumeBlock + 1
			appendMode = true
			log.Printf("Resuming from block %d (last block in file: %d)", actualStart, resumeBlock)
			}
		}
	}

	// Determine end block
	end := *endBlock
	if end == 0 || end > reader.ItemCount() {
		end = reader.ItemCount()
	}
	if actualStart >= end {
		log.Printf("Already complete: start block %d >= end block %d", actualStart, end)
		return
	}

	// Create output writer (append mode if resuming)
	writer, err := output.NewWriterWithMode(*format, *outPath, appendMode)
	if err != nil {
		log.Fatalf("Failed to create output writer: %v", err)
	}
	defer writer.Close()

	// Create error logger if specified
	var errorLogger *ErrorLogger
	if *errorLogPath != "" {
		errorLogger, err = NewErrorLogger(*errorLogPath)
		if err != nil {
			log.Fatalf("Failed to create error log: %v", err)
		}
		defer errorLogger.Close()
		log.Printf("Error log: %s", *errorLogPath)
	}

	// Handle interrupt
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan struct{})

	go func() {
		<-sigChan
		log.Println("Interrupt received, shutting down gracefully...")
		close(done)
	}()

	// Scan blocks
	if err := scanBlocks(reader, writer, errorLogger, actualStart, end, contractSet, eventSigSet, done); err != nil {
		log.Fatalf("Scan failed: %v", err)
	}

	if err := writer.Flush(); err != nil {
		log.Fatalf("Failed to flush output: %v", err)
	}
}

// getResumeBlock reads the last line of a JSONL file to find the last block number
func getResumeBlock(path string) (uint64, bool) {
	file, err := os.Open(path)
	if err != nil {
		return 0, false // File doesn't exist
	}
	defer file.Close()

	// Check if file is empty
	stat, err := file.Stat()
	if err != nil || stat.Size() == 0 {
		return 0, false
	}

	// Read last line efficiently by seeking from end
	var lastLine string
	
	// For small files, just scan all lines
	if stat.Size() < 1024*1024 { // < 1MB
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lastLine = scanner.Text()
		}
	} else {
		// For large files, seek to near end and find last line
		const tailSize = 4096
		seekPos := stat.Size() - tailSize
		if seekPos < 0 {
			seekPos = 0
		}
		file.Seek(seekPos, io.SeekStart)
		
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lastLine = scanner.Text()
		}
	}

	if lastLine == "" {
		return 0, false
	}

	// Parse the JSON to get blockNumber
	var record struct {
		BlockNumber uint64 `json:"blockNumber"`
	}
	if err := json.Unmarshal([]byte(lastLine), &record); err != nil {
		log.Printf("Warning: could not parse last line of %s: %v", path, err)
		return 0, false
	}

	return record.BlockNumber, true
}

func printInfo(reader *freezerscanner.FreezerReader) {
	fmt.Printf("Freezer Information:\n")
	fmt.Printf("  Item count (blocks): %d\n", reader.ItemCount())
	fmt.Printf("  CDAT files:\n")
	for _, info := range reader.CdatFileInfo() {
		fmt.Printf("    %s\n", info)
	}

	// Sample first few offsets
	fmt.Printf("\nFirst 10 index entries:\n")
	for i := uint64(0); i < 10 && i < reader.ItemCount(); i++ {
		offset, err := reader.ReadOffset(i)
		if err != nil {
			fmt.Printf("  Block %d: error %v\n", i, err)
		} else {
			fmt.Printf("  Block %d: offset %d\n", i, offset)
		}
	}

	// Sample last few offsets
	if reader.ItemCount() > 10 {
		fmt.Printf("\nLast 5 index entries:\n")
		start := reader.ItemCount() - 5
		for i := start; i < reader.ItemCount(); i++ {
			offset, err := reader.ReadOffset(i)
			if err != nil {
				fmt.Printf("  Block %d: error %v\n", i, err)
			} else {
				fmt.Printf("  Block %d: offset %d\n", i, offset)
			}
		}
	}
}

func validateIndex(reader *freezerscanner.FreezerReader) error {
	log.Printf("Validating index for %d blocks...", reader.ItemCount())
	start := time.Now()

	if err := reader.ValidateIndexRange(0, reader.ItemCount()); err != nil {
		return err
	}

	log.Printf("Index validated successfully in %v", time.Since(start))
	return nil
}

// ErrorLogger writes errors to a log file
type ErrorLogger struct {
	file *os.File
}

// NewErrorLogger creates a new error logger (appends to file)
func NewErrorLogger(path string) (*ErrorLogger, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	// Write session header
	fmt.Fprintf(f, "\n=== Scan session started at %s ===\n", time.Now().Format(time.RFC3339))
	return &ErrorLogger{file: f}, nil
}

// Log writes an error entry
func (el *ErrorLogger) Log(blockNum uint64, errType, errMsg string) {
	if el == nil || el.file == nil {
		return
	}
	fmt.Fprintf(el.file, "[%s] block=%d type=%s error=%s\n",
		time.Now().Format("2006-01-02 15:04:05"), blockNum, errType, errMsg)
}

// Close closes the error log file
func (el *ErrorLogger) Close() error {
	if el == nil || el.file == nil {
		return nil
	}
	fmt.Fprintf(el.file, "=== Scan session ended at %s ===\n", time.Now().Format(time.RFC3339))
	return el.file.Close()
}

func scanBlocks(reader *freezerscanner.FreezerReader, writer output.Writer, errorLogger *ErrorLogger,
	start, end uint64, contracts map[common.Address]bool, eventSigs map[common.Hash]bool,
	done <-chan struct{}) error {

	startTime := time.Now()
	var totalLogs, matchedLogs uint64
	var errorCount int
	var lastBlock uint64

	log.Printf("Scanning blocks %d to %d (filters: %d contracts, %d event sigs)",
		start, end, len(contracts), len(eventSigs))

	for block := start; block < end; block++ {
		lastBlock = block
		
		// Check for interrupt
		select {
		case <-done:
			log.Printf("Scan interrupted at block %d (progress saved)", block)
			return nil
		default:
		}

		// Progress logging
		if *progressEvery > 0 && block > start && (block-start)%*progressEvery == 0 {
			elapsed := time.Since(startTime)
			blocksScanned := block - start
			blocksPerSec := float64(blocksScanned) / elapsed.Seconds()
			remaining := float64(end-block) / blocksPerSec
			log.Printf("Progress: block %d/%d (%.1f%%), %.1f blocks/sec, ETA: %v, matched: %d/%d logs, errors: %d",
				block, end, float64(block-start)*100/float64(end-start),
				blocksPerSec, time.Duration(remaining)*time.Second,
				matchedLogs, totalLogs, errorCount)
		}

		// Read raw receipts data
		data, err := reader.ReadItem(block)
		if err != nil {
			errorCount++
			if errorLogger != nil {
				errorLogger.Log(block, "read", err.Error())
			}
			if *bestEffort {
				continue
			}
			return fmt.Errorf("Block %d: read error: %v", block, err)
		}

		if len(data) == 0 {
			continue // Empty block (no transactions)
		}

		// Decode receipts (Arbitrum Nitro format)
		logs, err := decode.DecodeArbitrumReceipts(data)
		if err != nil {
			errorCount++
			if errorLogger != nil {
				errorLogger.Log(block, "decode", err.Error())
			}
			if *bestEffort {
				continue
			}
			return fmt.Errorf("Block %d: decode error: %v", block, err)
		}

		// Process logs
		for logIdx, logEntry := range logs {
			totalLogs++

			// Apply filters
			if len(contracts) > 0 && !contracts[logEntry.Address] {
				continue
			}
			if len(eventSigs) > 0 {
				if len(logEntry.Topics) == 0 {
					continue
				}
				if !eventSigs[logEntry.Topics[0]] {
					continue
				}
			}

			matchedLogs++

			// Write record
			record := output.LogEntryToRecord(block, uint(logIdx), logEntry, *includeData)
			if err := writer.Write(record); err != nil {
				errorCount++
				if errorLogger != nil {
					errorLogger.Log(block, "write", err.Error())
				}
				return fmt.Errorf("Block %d: write error: %v", block, err)
			}
		}
	}

	elapsed := time.Since(startTime)
	blocksScanned := end - start
	log.Printf("Scan complete: %d blocks in %v (%.1f blocks/sec)",
		blocksScanned, elapsed, float64(blocksScanned)/elapsed.Seconds())
	log.Printf("Total logs: %d, Matched: %d, Last block: %d, Errors: %d", 
		totalLogs, matchedLogs, lastBlock, errorCount)

	return nil
}
