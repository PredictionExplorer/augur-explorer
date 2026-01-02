// freezer-scan scans geth freezer receipts to extract event logs
// matching specified contract addresses and event signatures.
//
// Usage:
//
//	freezer-scan --ancientDir /path/to/mainnet \
//	             --startBlock 0 --endBlock 1000000 \
//	             --contracts 0x123...,0x456... \
//	             --eventSigs 0xabc...,0xdef... \
//	             --out events.jsonl --format jsonl \
//	             --workers 16
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
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"

	freezerscanner "github.com/PredictionExplorer/augur-explorer/rwcg/freezer-scanner"
	"github.com/PredictionExplorer/augur-explorer/rwcg/freezer-scanner/decode"
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
	progressEvery   = flag.Uint64("progressEvery", 100000, "Log progress every N blocks")
	bestEffort      = flag.Bool("bestEffort", false, "Continue on decode errors")
	validate        = flag.Bool("validate", false, "Validate index before scanning")
	validateOnly    = flag.Bool("validateOnly", false, "Only validate index, don't scan")
	showInfo        = flag.Bool("info", false, "Show information about the freezer data")
	noResume        = flag.Bool("noResume", false, "Don't resume from existing file, start fresh")
	numWorkers      = flag.Int("workers", 0, "Number of parallel chunk workers (0 = auto, typically 16-64)")
	chunkSize       = flag.Uint64("chunkSize", 50000, "Blocks per worker chunk")
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
		dir = strings.TrimSuffix(*receiptsCidx, "/receipts.cidx")
		dir = strings.TrimSuffix(dir, "/ancient/receipts.cidx")
	}

	// Determine number of workers (chunk workers processing in parallel)
	// Each worker processes blocks sequentially within its chunk
	// Parallelism comes from multiple chunks running concurrently
	workers := *numWorkers
	if workers <= 0 {
		workers = runtime.NumCPU() * 2
		if workers < 16 {
			workers = 16
		}
		if workers > 64 {
			workers = 64
		}
	}

	log.Printf("Starting freezer-scan with %d parallel workers", workers)

	// Open parallel reader (loads index into memory)
	log.Println("Loading index into memory...")
	loadStart := time.Now()
	reader, err := freezerscanner.NewParallelReader(dir)
	if err != nil {
		log.Fatalf("Failed to open freezer: %v", err)
	}
	
	totalBlocks, indexSizeMB, cdatCount := reader.GetIndexStats()
	log.Printf("Index loaded in %v: %d blocks (%.1f MB), %d cdat files",
		time.Since(loadStart), totalBlocks, indexSizeMB, cdatCount)

	// Show info if requested
	if *showInfo {
		printInfo(reader)
		if !*validate && !*validateOnly {
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
			sig = strings.TrimPrefix(sig, "0x")
			if len(sig) != 64 {
				log.Fatalf("Invalid event signature (must be 32 bytes hex): %s", sig)
			}
			eventSigSet[common.HexToHash(sig)] = true
		}
	}

	// Determine end block - limit to available data
	end := *endBlock
	maxAvailable := reader.MaxAvailableBlock()
	if end == 0 || end > maxAvailable {
		end = maxAvailable
		log.Printf("Limiting end block to %d (max available data)", end)
	}

	// Handle output path - if it's a directory, create a filename
	outputPath := *outPath
	if outputPath != "-" {
		if info, err := os.Stat(outputPath); err == nil && info.IsDir() {
			// It's a directory - create filename with block range
			outputPath = filepath.Join(outputPath, fmt.Sprintf("events_%d_%d.jsonl", *startBlock, end))
			log.Printf("Output directory detected, using: %s", outputPath)
		}
	}

	// Determine start block (check for resume)
	actualStart := *startBlock
	appendMode := false

	if outputPath != "-" && !*noResume {
		if resumeBlock, found := getResumeBlock(outputPath); found {
			if resumeBlock >= actualStart {
				actualStart = resumeBlock + 1
				appendMode = true
				log.Printf("Resuming from block %d (last block in file: %d)", actualStart, resumeBlock)
			}
		}
	}
	
	if actualStart >= end {
		log.Printf("Already complete: start block %d >= end block %d", actualStart, end)
		return
	}

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
	var interrupted atomic.Bool

	go func() {
		<-sigChan
		log.Println("Interrupt received, shutting down gracefully...")
		interrupted.Store(true)
		close(done)
	}()

	// Run parallel scan
	if err := parallelScan(reader, actualStart, end, workers,
		contractSet, eventSigSet, outputPath, appendMode, errorLogger, done, &interrupted); err != nil {
		log.Fatalf("Scan failed: %v", err)
	}
}

// parallelScan processes blocks in parallel using chunk workers
// Each worker processes a chunk of blocks sequentially
// Parallelism comes from multiple workers processing different chunks concurrently
func parallelScan(reader *freezerscanner.ParallelReader, start, end uint64, numWorkers int,
	contracts map[common.Address]bool, eventSigs map[common.Hash]bool,
	outPath string, appendMode bool, errorLogger *ErrorLogger,
	done <-chan struct{}, interrupted *atomic.Bool) error {

	startTime := time.Now()
	totalBlocks := end - start

	log.Printf("Parallel scan: blocks %d to %d (%d blocks), %d workers", start, end, totalBlocks, numWorkers)

	// Create temp directory for chunk output files
	tempDir, err := os.MkdirTemp("", "freezer-scan-*")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create work chunks for ordered output
	var chunks []freezerscanner.BlockRange
	for chunkStart := start; chunkStart < end; chunkStart += *chunkSize {
		chunkEnd := chunkStart + *chunkSize
		if chunkEnd > end {
			chunkEnd = end
		}
		chunks = append(chunks, freezerscanner.BlockRange{Start: chunkStart, End: chunkEnd})
	}

	log.Printf("  Chunks: %d (each %d blocks)", len(chunks), *chunkSize)

	// Progress tracking
	var processedBlocks atomic.Uint64
	var totalMatched atomic.Uint64
	var totalLogs atomic.Uint64
	var totalErrors atomic.Int64

	// Start progress reporter
	progressDone := make(chan struct{})
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				processed := processedBlocks.Load()
				matched := totalMatched.Load()
				logs := totalLogs.Load()
				errors := totalErrors.Load()
				elapsed := time.Since(startTime)
				if elapsed.Seconds() < 1 {
					continue
				}
				blocksPerSec := float64(processed) / elapsed.Seconds()
				remaining := float64(totalBlocks-processed) / blocksPerSec
				pct := float64(processed) * 100 / float64(totalBlocks)

				log.Printf("Progress: %d/%d blocks (%.1f%%), %.0f blocks/sec, ETA: %v, matched: %d/%d logs, errors: %d",
					processed, totalBlocks, pct, blocksPerSec,
					time.Duration(remaining)*time.Second, matched, logs, errors)
			case <-progressDone:
				return
			}
		}
	}()

	// Process chunks in parallel
	chunkChan := make(chan int, len(chunks))
	resultChan := make(chan ChunkResult, numWorkers)

	// Start workers - each processes chunks sequentially
	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			// Create worker-specific reader with its own file handles
			workerReader := reader.NewWorkerReader()
			defer workerReader.Close()

			for chunkIdx := range chunkChan {
				if interrupted.Load() {
					return
				}

				chunk := chunks[chunkIdx]
				tempFile := filepath.Join(tempDir, fmt.Sprintf("chunk-%08d.jsonl", chunkIdx))

				err := processChunk(workerReader, chunk, tempFile,
					contracts, eventSigs, errorLogger,
					&processedBlocks, &totalMatched, &totalLogs, &totalErrors,
					interrupted)

				resultChan <- ChunkResult{
					ChunkIdx: chunkIdx,
					TempFile: tempFile,
					Err:      err,
				}
			}
		}(w)
	}

	// Send chunks to workers
	go func() {
		for i := range chunks {
			select {
			case chunkChan <- i:
			case <-done:
				break
			}
		}
		close(chunkChan)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	results := make([]ChunkResult, 0, len(chunks))
	for result := range resultChan {
		results = append(results, result)
		if result.Err != nil && !*bestEffort {
			close(progressDone)
			return fmt.Errorf("chunk %d failed: %v", result.ChunkIdx, result.Err)
		}
	}

	close(progressDone)

	// Sort by chunk index for ordered merge
	sort.Slice(results, func(i, j int) bool {
		return results[i].ChunkIdx < results[j].ChunkIdx
	})

	log.Println("Merging results...")
	if err := mergeChunkFiles(results, outPath, appendMode); err != nil {
		return fmt.Errorf("merge failed: %v", err)
	}

	// Final statistics
	elapsed := time.Since(startTime)
	processed := processedBlocks.Load()
	matched := totalMatched.Load()
	logs := totalLogs.Load()
	errors := totalErrors.Load()

	log.Printf("Scan complete: %d blocks in %v (%.0f blocks/sec)",
		processed, elapsed, float64(processed)/elapsed.Seconds())
	log.Printf("Total logs: %d, Matched: %d, Errors: %d", logs, matched, errors)

	return nil
}

// processChunk processes a chunk sequentially (simple and memory-efficient)
func processChunk(reader *freezerscanner.WorkerReader, chunk freezerscanner.BlockRange,
	tempFile string,
	contracts map[common.Address]bool, eventSigs map[common.Hash]bool,
	errorLogger *ErrorLogger,
	processedBlocks, totalMatched, totalLogs *atomic.Uint64, totalErrors *atomic.Int64,
	interrupted *atomic.Bool) error {

	// Create output file
	f, err := os.Create(tempFile)
	if err != nil {
		return fmt.Errorf("create temp file: %v", err)
	}
	defer f.Close()

	writer := bufio.NewWriterSize(f, 256*1024) // 256KB buffer
	defer writer.Flush()

	// Process blocks sequentially within chunk - simple and memory efficient
	// Parallelism comes from multiple chunks being processed concurrently
	for block := chunk.Start; block < chunk.End; block++ {
		if interrupted.Load() {
			break
		}

		data, err := reader.ReadItem(block)
		processedBlocks.Add(1)

		if err != nil {
			totalErrors.Add(1)
			if errorLogger != nil {
				errorLogger.Log(block, "read", err.Error())
			}
			if !*bestEffort {
				return fmt.Errorf("block %d read: %v", block, err)
			}
			continue
		}

		if len(data) == 0 {
			continue
		}

		// Decode
		logs, err := decode.DecodeArbitrumReceipts(data)
		if err != nil {
			totalErrors.Add(1)
			if errorLogger != nil {
				errorLogger.Log(block, "decode", err.Error())
			}
			if !*bestEffort {
				return fmt.Errorf("block %d decode: %v", block, err)
			}
			continue
		}

		// Filter and write
		for logIdx, logEntry := range logs {
			totalLogs.Add(1)

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

			totalMatched.Add(1)

			record := output.LogEntryToRecord(block, uint(logIdx), logEntry, *includeData)
			jsonBytes, _ := json.Marshal(record)
			writer.Write(jsonBytes)
			writer.WriteByte('\n')
		}
	}

	return nil
}

// ChunkResult contains the result from processing a chunk
type ChunkResult struct {
	ChunkIdx int
	TempFile string
	Err      error
}

// mergeChunkFiles merges temp files into the final output
func mergeChunkFiles(results []ChunkResult, outPath string, appendMode bool) error {
	if outPath == "-" {
		for _, r := range results {
			if r.TempFile == "" {
				continue
			}
			data, err := os.ReadFile(r.TempFile)
			if err != nil {
				return err
			}
			os.Stdout.Write(data)
		}
		return nil
	}

	flags := os.O_CREATE | os.O_WRONLY
	if appendMode {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}

	out, err := os.OpenFile(outPath, flags, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	writer := bufio.NewWriterSize(out, 1024*1024) // 1MB buffer
	defer writer.Flush()

	for _, r := range results {
		if r.TempFile == "" {
			continue
		}

		f, err := os.Open(r.TempFile)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}

		_, err = io.Copy(writer, f)
		f.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// getResumeBlock reads the last line of a JSONL file to find the last block number
func getResumeBlock(path string) (uint64, bool) {
	file, err := os.Open(path)
	if err != nil {
		return 0, false
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil || stat.Size() == 0 {
		return 0, false
	}

	var lastLine string
	if stat.Size() < 1024*1024 {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lastLine = scanner.Text()
		}
	} else {
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

	var record struct {
		BlockNumber uint64 `json:"blockNumber"`
	}
	if err := json.Unmarshal([]byte(lastLine), &record); err != nil {
		log.Printf("Warning: could not parse last line of %s: %v", path, err)
		return 0, false
	}

	return record.BlockNumber, true
}

func printInfo(reader *freezerscanner.ParallelReader) {
	totalBlocks, indexSizeMB, cdatCount := reader.GetIndexStats()
	fmt.Printf("Freezer Information:\n")
	fmt.Printf("  Item count (blocks): %d\n", totalBlocks)
	fmt.Printf("  Index size: %.1f MB (loaded in memory)\n", indexSizeMB)
	fmt.Printf("  Max available block: %d\n", reader.MaxAvailableBlock())
	fmt.Printf("  CDAT files: %d\n", cdatCount)
	fmt.Printf("\n  CDAT file details:\n")
	for _, info := range reader.CdatFileInfo() {
		fmt.Printf("    %s\n", info)
	}

	// Sample offsets
	fmt.Printf("\nFirst 10 index entries:\n")
	offsets := reader.ReadOffsetBatch(0, 10)
	for i, offset := range offsets {
		fmt.Printf("  Block %d: offset %d\n", i, offset)
	}

	if totalBlocks > 10 {
		fmt.Printf("\nLast 5 index entries:\n")
		offsets = reader.ReadOffsetBatch(totalBlocks-5, 5)
		for i, offset := range offsets {
			fmt.Printf("  Block %d: offset %d\n", totalBlocks-5+uint64(i), offset)
		}
	}
}

// ErrorLogger writes errors to a log file
type ErrorLogger struct {
	file *os.File
	mu   sync.Mutex
}

// NewErrorLogger creates a new error logger (appends to file)
func NewErrorLogger(path string) (*ErrorLogger, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(f, "\n=== Scan session started at %s ===\n", time.Now().Format(time.RFC3339))
	return &ErrorLogger{file: f}, nil
}

// Log writes an error entry (thread-safe)
func (el *ErrorLogger) Log(blockNum uint64, errType, errMsg string) {
	if el == nil || el.file == nil {
		return
	}
	el.mu.Lock()
	defer el.mu.Unlock()
	fmt.Fprintf(el.file, "[%s] block=%d type=%s error=%s\n",
		time.Now().Format("2006-01-02 15:04:05"), blockNum, errType, errMsg)
}

// Close closes the error log file
func (el *ErrorLogger) Close() error {
	if el == nil || el.file == nil {
		return nil
	}
	el.mu.Lock()
	defer el.mu.Unlock()
	fmt.Fprintf(el.file, "=== Scan session ended at %s ===\n", time.Now().Format(time.RFC3339))
	return el.file.Close()
}
