// Package scan is the freezer-scan orchestration: it reads geth freezer
// receipts through internal/freezer's parallel readers, decodes each block's
// logs, filters them by contract address and event signature, and writes
// matching entries as JSONL — resumable, chunk-parallel and
// interrupt-tolerant.
//
// cmd/freezer-scan is a thin flag wrapper around Run; everything here is
// injectable (output writers, progress logging, temp directory) so the whole
// pipeline is tested against synthetic freezer fixtures.
package scan

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"

	freezerscanner "github.com/PredictionExplorer/augur-explorer/internal/freezer"
	"github.com/PredictionExplorer/augur-explorer/internal/freezer/decode"
	"github.com/PredictionExplorer/augur-explorer/internal/freezer/output"
)

// DefaultChunkSize is the number of blocks each worker chunk covers.
const DefaultChunkSize = 50_000

// Options configures one scan run.
type Options struct {
	// StartBlock is the first block to scan.
	StartBlock uint64
	// EndBlock bounds the scan (exclusive); 0 means "all available data",
	// and any value is clamped to the reader's available data.
	EndBlock uint64
	// Contracts filters logs by emitting address; empty means all.
	Contracts map[common.Address]bool
	// EventSigs filters logs by topic0; empty means all.
	EventSigs map[common.Hash]bool
	// OutPath receives the JSONL output. "-" writes to Stdout; an existing
	// directory gets an events_<start>_<end>.jsonl file inside.
	OutPath string
	// Stdout receives the output when OutPath is "-"; defaults to os.Stdout.
	Stdout io.Writer
	// ErrorLogPath appends per-block read/decode errors when set.
	ErrorLogPath string
	// IncludeData embeds the full hex log data in each record.
	IncludeData bool
	// BestEffort continues past per-block read/decode errors.
	BestEffort bool
	// NoResume starts fresh even when OutPath already has scanned blocks.
	NoResume bool
	// Workers is the parallel chunk worker count; <=0 selects a CPU-based
	// default via DefaultWorkers.
	Workers int
	// ChunkSize is the blocks-per-chunk unit of parallelism; <=0 uses
	// DefaultChunkSize.
	ChunkSize uint64
	// Logf receives progress lines; defaults to discarding them.
	Logf func(format string, args ...any)
	// ProgressEvery spaces the progress reports; defaults to 5s.
	ProgressEvery time.Duration
}

// Stats summarizes a scan run.
type Stats struct {
	// StartBlock and EndBlock are the effective range after resume/clamping.
	StartBlock uint64
	EndBlock   uint64
	// Resumed reports whether the run continued an existing output file.
	Resumed bool
	// OutputPath is the effective output destination ("-" for stdout).
	OutputPath string
	// ProcessedBlocks counts blocks read (including empty and errored ones).
	ProcessedBlocks uint64
	// TotalLogs counts decoded logs before filtering.
	TotalLogs uint64
	// MatchedLogs counts logs written to the output.
	MatchedLogs uint64
	// Errors counts per-block read/decode failures (best-effort mode).
	Errors int64
}

// DefaultWorkers derives the worker count from the CPU count, clamped to
// [16, 64] like the legacy tool.
func DefaultWorkers(numCPU int) int {
	workers := numCPU * 2
	if workers < 16 {
		workers = 16
	}
	if workers > 64 {
		workers = 64
	}
	return workers
}

// ParseContracts parses a comma-separated list of hex contract addresses.
func ParseContracts(csv string) (map[common.Address]bool, error) {
	out := make(map[common.Address]bool)
	for _, addr := range strings.Split(csv, ",") {
		addr = strings.TrimSpace(addr)
		if addr == "" {
			continue
		}
		if !common.IsHexAddress(addr) {
			return nil, fmt.Errorf("invalid contract address: %s", addr)
		}
		out[common.HexToAddress(addr)] = true
	}
	return out, nil
}

// ParseEventSigs parses a comma-separated list of 32-byte hex topic0 hashes.
func ParseEventSigs(csv string) (map[common.Hash]bool, error) {
	out := make(map[common.Hash]bool)
	for _, sig := range strings.Split(csv, ",") {
		sig = strings.TrimSpace(sig)
		if sig == "" {
			continue
		}
		sig = strings.TrimPrefix(sig, "0x")
		if len(sig) != 64 {
			return nil, fmt.Errorf("invalid event signature (must be 32 bytes hex): %s", sig)
		}
		out[common.HexToHash(sig)] = true
	}
	return out, nil
}

// chunkResult is the outcome of one worker chunk.
type chunkResult struct {
	chunkIdx int
	tempFile string
	err      error
}

// Run executes the scan. Cancelling ctx stops the workers gracefully: chunks
// already finished are still merged into the output, and Run returns the
// partial Stats with a nil error (the legacy interrupt behavior).
func Run(ctx context.Context, reader *freezerscanner.ParallelReader, opts Options) (Stats, error) {
	logf := opts.Logf
	if logf == nil {
		logf = func(string, ...any) {}
	}
	stdout := opts.Stdout
	if stdout == nil {
		stdout = os.Stdout
	}
	chunkSize := opts.ChunkSize
	if chunkSize == 0 {
		chunkSize = DefaultChunkSize
	}
	workers := opts.Workers
	if workers <= 0 {
		workers = DefaultWorkers(runtime.NumCPU())
	}
	progressEvery := opts.ProgressEvery
	if progressEvery <= 0 {
		progressEvery = 5 * time.Second
	}

	stats := Stats{StartBlock: opts.StartBlock, EndBlock: opts.EndBlock, OutputPath: opts.OutPath}

	// Clamp the end to the data that actually exists.
	maxAvailable := reader.MaxAvailableBlock()
	if stats.EndBlock == 0 || stats.EndBlock > maxAvailable {
		stats.EndBlock = maxAvailable
		logf("Limiting end block to %d (max available data)", stats.EndBlock)
	}

	// A directory output path gets a range-named file inside.
	if stats.OutputPath != "-" {
		if info, err := os.Stat(stats.OutputPath); err == nil && info.IsDir() {
			stats.OutputPath = filepath.Join(stats.OutputPath, fmt.Sprintf("events_%d_%d.jsonl", stats.StartBlock, stats.EndBlock))
			logf("Output directory detected, using: %s", stats.OutputPath)
		}
	}

	// Resume from the last block already in the output file.
	appendMode := false
	if stats.OutputPath != "-" && !opts.NoResume {
		if resumeBlock, found, warn := lastScannedBlock(stats.OutputPath); found && resumeBlock >= stats.StartBlock {
			stats.StartBlock = resumeBlock + 1
			stats.Resumed = true
			appendMode = true
			logf("Resuming from block %d (last block in file: %d)", stats.StartBlock, resumeBlock)
		} else if warn != nil {
			logf("Warning: %v", warn)
		}
	}

	if stats.StartBlock >= stats.EndBlock {
		logf("Already complete: start block %d >= end block %d", stats.StartBlock, stats.EndBlock)
		return stats, nil
	}

	var errorLog *ErrorLog
	if opts.ErrorLogPath != "" {
		var err error
		errorLog, err = NewErrorLog(opts.ErrorLogPath)
		if err != nil {
			return stats, fmt.Errorf("failed to create error log: %w", err)
		}
		defer func() { _ = errorLog.Close() }() // best-effort close of the error log
		logf("Error log: %s", opts.ErrorLogPath)
	}

	startTime := time.Now()
	totalBlocks := stats.EndBlock - stats.StartBlock
	logf("Parallel scan: blocks %d to %d (%d blocks), %d workers", stats.StartBlock, stats.EndBlock, totalBlocks, workers)

	tempDir, err := os.MkdirTemp("", "freezer-scan-*")
	if err != nil {
		return stats, fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }() // best-effort temp dir cleanup

	var chunks []freezerscanner.BlockRange
	for chunkStart := stats.StartBlock; chunkStart < stats.EndBlock; chunkStart += chunkSize {
		chunkEnd := chunkStart + chunkSize
		if chunkEnd > stats.EndBlock {
			chunkEnd = stats.EndBlock
		}
		chunks = append(chunks, freezerscanner.BlockRange{Start: chunkStart, End: chunkEnd})
	}
	logf("  Chunks: %d (each %d blocks)", len(chunks), chunkSize)

	var processedBlocks, totalMatched, totalLogs atomic.Uint64
	var totalErrors atomic.Int64

	progressDone := make(chan struct{})
	go func() {
		ticker := time.NewTicker(progressEvery)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				processed := processedBlocks.Load()
				elapsed := time.Since(startTime)
				if elapsed.Seconds() < 1 {
					continue
				}
				blocksPerSec := float64(processed) / elapsed.Seconds()
				remaining := float64(totalBlocks-processed) / blocksPerSec
				pct := float64(processed) * 100 / float64(totalBlocks)
				logf("Progress: %d/%d blocks (%.1f%%), %.0f blocks/sec, ETA: %v, matched: %d/%d logs, errors: %d",
					processed, totalBlocks, pct, blocksPerSec,
					time.Duration(remaining)*time.Second, totalMatched.Load(), totalLogs.Load(), totalErrors.Load())
			case <-progressDone:
				return
			}
		}
	}()

	chunkChan := make(chan int, len(chunks))
	resultChan := make(chan chunkResult, workers)

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workerReader := reader.NewWorkerReader()
			defer func() { _ = workerReader.Close() }() // best-effort close of read-only handles

			for chunkIdx := range chunkChan {
				if ctx.Err() != nil {
					return
				}
				chunk := chunks[chunkIdx]
				tempFile := filepath.Join(tempDir, fmt.Sprintf("chunk-%08d.jsonl", chunkIdx))
				err := processChunk(ctx, workerReader, chunk, tempFile, opts,
					&processedBlocks, &totalMatched, &totalLogs, &totalErrors, errorLog)
				resultChan <- chunkResult{chunkIdx: chunkIdx, tempFile: tempFile, err: err}
			}
		}()
	}

	go func() {
		defer close(chunkChan)
		for i := range chunks {
			select {
			case chunkChan <- i:
			case <-ctx.Done():
				return // stop queueing further chunks on interrupt
			}
		}
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	results := make([]chunkResult, 0, len(chunks))
	for result := range resultChan {
		results = append(results, result)
		if result.err != nil && !opts.BestEffort {
			close(progressDone)
			return stats, fmt.Errorf("chunk %d failed: %w", result.chunkIdx, result.err)
		}
	}
	close(progressDone)

	// Merge in chunk order so the output stays block-sorted.
	sort.Slice(results, func(i, j int) bool { return results[i].chunkIdx < results[j].chunkIdx })
	logf("Merging results...")
	if err := mergeChunkFiles(results, stats.OutputPath, appendMode, stdout); err != nil {
		return stats, fmt.Errorf("merge failed: %w", err)
	}

	stats.ProcessedBlocks = processedBlocks.Load()
	stats.TotalLogs = totalLogs.Load()
	stats.MatchedLogs = totalMatched.Load()
	stats.Errors = totalErrors.Load()

	elapsed := time.Since(startTime)
	logf("Scan complete: %d blocks in %v (%.0f blocks/sec)",
		stats.ProcessedBlocks, elapsed, float64(stats.ProcessedBlocks)/elapsed.Seconds())
	logf("Total logs: %d, Matched: %d, Errors: %d", stats.TotalLogs, stats.MatchedLogs, stats.Errors)
	return stats, nil
}

// processChunk scans one block range sequentially into a temp JSONL file.
// Parallelism comes from multiple chunks running concurrently.
func processChunk(ctx context.Context, reader *freezerscanner.WorkerReader, chunk freezerscanner.BlockRange,
	tempFile string, opts Options,
	processedBlocks, totalMatched, totalLogs *atomic.Uint64, totalErrors *atomic.Int64,
	errorLog *ErrorLog,
) error {
	f, err := os.Create(filepath.Clean(tempFile))
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	// Best-effort close for early-return paths; the success path below
	// flushes and closes explicitly so write errors are not lost.
	defer func() { _ = f.Close() }()

	writer := bufio.NewWriterSize(f, 256*1024)

	for block := chunk.Start; block < chunk.End; block++ {
		if ctx.Err() != nil {
			break
		}

		data, err := reader.ReadItem(block)
		processedBlocks.Add(1)
		if err != nil {
			totalErrors.Add(1)
			errorLog.Log(block, "read", err.Error())
			if !opts.BestEffort {
				return fmt.Errorf("block %d read: %w", block, err)
			}
			continue
		}
		if len(data) == 0 {
			continue
		}

		logs, err := decode.ArbitrumReceipts(data)
		if err != nil {
			totalErrors.Add(1)
			errorLog.Log(block, "decode", err.Error())
			if !opts.BestEffort {
				return fmt.Errorf("block %d decode: %w", block, err)
			}
			continue
		}

		for logIdx, logEntry := range logs {
			totalLogs.Add(1)
			if len(opts.Contracts) > 0 && !opts.Contracts[logEntry.Address] {
				continue
			}
			if len(opts.EventSigs) > 0 {
				if len(logEntry.Topics) == 0 || !opts.EventSigs[logEntry.Topics[0]] {
					continue
				}
			}
			totalMatched.Add(1)

			record := output.LogEntryToRecord(block, uint(logIdx), logEntry, opts.IncludeData)
			jsonBytes, err := json.Marshal(record)
			if err != nil {
				return fmt.Errorf("marshal record: %w", err)
			}
			if _, err := writer.Write(jsonBytes); err != nil {
				return fmt.Errorf("write temp file: %w", err)
			}
			if err := writer.WriteByte('\n'); err != nil {
				return fmt.Errorf("write temp file: %w", err)
			}
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush temp file: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close temp file: %w", err)
	}
	return nil
}

// mergeChunkFiles concatenates the chunk temp files into the final output.
func mergeChunkFiles(results []chunkResult, outPath string, appendMode bool, stdout io.Writer) error {
	if outPath == "-" {
		for _, r := range results {
			if r.tempFile == "" {
				continue
			}
			data, err := os.ReadFile(r.tempFile)
			if err != nil {
				return err
			}
			if _, err := stdout.Write(data); err != nil {
				return err
			}
		}
		return nil
	}

	flags := os.O_CREATE | os.O_WRONLY
	if appendMode {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}
	out, err := os.OpenFile(filepath.Clean(outPath), flags, 0o644) //nolint:gosec // operator-chosen output path; 0644 is intentional for shareable output
	if err != nil {
		return err
	}
	// Best-effort close for early-return paths; the success path below
	// flushes and closes explicitly so write errors are not lost.
	defer func() { _ = out.Close() }()

	writer := bufio.NewWriterSize(out, 1024*1024)
	for _, r := range results {
		if r.tempFile == "" {
			continue
		}
		f, err := os.Open(r.tempFile)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		_, err = io.Copy(writer, f)
		_ = f.Close() // best-effort close on read path
		if err != nil {
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	return out.Close()
}

// lastScannedBlock reads the last JSONL line of path and extracts its block
// number. warn carries a non-fatal parse problem worth logging.
func lastScannedBlock(path string) (block uint64, found bool, warn error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return 0, false, nil
	}
	defer func() { _ = file.Close() }() // best-effort close on read path

	stat, err := file.Stat()
	if err != nil || stat.Size() == 0 {
		return 0, false, nil
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
		if _, err := file.Seek(seekPos, io.SeekStart); err != nil {
			return 0, false, nil
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lastLine = scanner.Text()
		}
	}
	if lastLine == "" {
		return 0, false, nil
	}

	var record struct {
		BlockNumber uint64 `json:"blockNumber"`
	}
	if err := json.Unmarshal([]byte(lastLine), &record); err != nil {
		return 0, false, fmt.Errorf("could not parse last line of %s: %w", path, err)
	}
	return record.BlockNumber, true, nil
}

// PrintInfo writes the freezer index summary the --info flag shows.
func PrintInfo(w io.Writer, reader *freezerscanner.ParallelReader) {
	printf := func(format string, args ...any) { _, _ = fmt.Fprintf(w, format, args...) }
	totalBlocks, indexSizeMB, cdatCount := reader.GetIndexStats()
	printf("Freezer Information:\n")
	printf("  Item count (blocks): %d\n", totalBlocks)
	printf("  Index size: %.1f MB (loaded in memory)\n", indexSizeMB)
	printf("  Max available block: %d\n", reader.MaxAvailableBlock())
	printf("  CDAT files: %d\n", cdatCount)
	printf("\n  CDAT file details:\n")
	for _, info := range reader.CdatFileInfo() {
		printf("    %s\n", info)
	}

	printf("\nFirst 10 index entries:\n")
	offsets := reader.ReadOffsetBatch(0, 10)
	for i, offset := range offsets {
		printf("  Block %d: offset %d\n", i, offset)
	}
	if totalBlocks > 10 {
		printf("\nLast 5 index entries:\n")
		offsets = reader.ReadOffsetBatch(totalBlocks-5, 5)
		for i, offset := range offsets {
			printf("  Block %d: offset %d\n", totalBlocks-5+uint64(i), offset)
		}
	}
}
