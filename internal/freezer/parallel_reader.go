package freezerscanner

import (
	"cmp"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"sync"
)

// ParallelReader provides high-performance parallel access to freezer data
// by memory-mapping the index and using multiple goroutines for reading.
type ParallelReader struct {
	indexData []byte       // Memory-mapped or loaded index data
	itemCount uint64       // Number of blocks in the index
	cdatDir   string       // Directory containing cdat files
	cdatFiles []*cdatEntry // Sorted list of cdat files
	chunkSize uint64       // Size of each cdat chunk (for offset calculation)
}

// NewParallelReader creates a new parallel freezer reader with the index loaded into memory.
func NewParallelReader(ancientDir string) (*ParallelReader, error) {
	pr := &ParallelReader{
		chunkSize: DefaultChunkSize,
	}

	// Find cidx file
	cidxPath := filepath.Join(ancientDir, "receipts.cidx")
	if _, err := os.Stat(cidxPath); err != nil {
		cidxPath = filepath.Join(ancientDir, "ancient", "receipts.cidx")
		if _, err := os.Stat(cidxPath); err != nil {
			return nil, fmt.Errorf("receipts.cidx not found in %s: %w", ancientDir, err)
		}
	}

	// Find cdat directory
	cdatDir := filepath.Dir(cidxPath)
	ancientSubdir := filepath.Join(cdatDir, "ancient")
	if info, err := os.Stat(ancientSubdir); err == nil && info.IsDir() {
		cdatDir = ancientSubdir
	}
	pr.cdatDir = cdatDir

	// Load entire index into memory - this is critical for performance
	// 416M blocks * 6 bytes = 2.5GB - large but manageable
	indexData, err := os.ReadFile(filepath.Clean(cidxPath))
	if err != nil {
		return nil, fmt.Errorf("failed to read index file: %w", err)
	}

	if len(indexData)%IndexEntrySize != 0 {
		return nil, fmt.Errorf("invalid index size: %d not divisible by %d", len(indexData), IndexEntrySize)
	}

	pr.indexData = indexData
	pr.itemCount = uint64(len(indexData) / IndexEntrySize)

	// Index cdat files
	if err := pr.indexCdatFiles(); err != nil {
		return nil, err
	}

	return pr, nil
}

// indexCdatFiles discovers and indexes all cdat files.
func (pr *ParallelReader) indexCdatFiles() error {
	pattern := filepath.Join(pr.cdatDir, "receipts.*.cdat")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("glob error: %w", err)
	}

	if len(matches) == 0 {
		return fmt.Errorf("no cdat files found matching %s", pattern)
	}

	var entries []*cdatEntry
	for _, path := range matches {
		var index uint64
		base := filepath.Base(path)
		// The unsigned target makes Sscanf reject negative indexes, which
		// the offset math below would otherwise wrap.
		if _, err := fmt.Sscanf(base, "receipts.%d.cdat", &index); err != nil {
			continue
		}

		info, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("failed to stat %s: %w", path, err)
		}
		size := info.Size()
		if size < 0 {
			return fmt.Errorf("stat %s: negative size %d", path, size)
		}

		entries = append(entries, &cdatEntry{
			index: index,
			path:  path,
			size:  uint64(size),
		})
	}

	slices.SortFunc(entries, func(a, b *cdatEntry) int {
		return cmp.Compare(a.index, b.index)
	})

	for _, entry := range entries {
		entry.startOffset = entry.index * pr.chunkSize
	}

	pr.cdatFiles = entries
	return nil
}

// getOffset reads the offset for a block directly from the in-memory index (no syscall!)
func (pr *ParallelReader) getOffset(blockNum uint64) uint64 {
	pos := blockNum * IndexEntrySize
	return uint64(pr.indexData[pos])<<40 |
		uint64(pr.indexData[pos+1])<<32 |
		uint64(pr.indexData[pos+2])<<24 |
		uint64(pr.indexData[pos+3])<<16 |
		uint64(pr.indexData[pos+4])<<8 |
		uint64(pr.indexData[pos+5])
}

// BlockRange represents a range of blocks to process.
type BlockRange struct {
	Start uint64
	End   uint64
}

// BlockData contains raw data for a single block.
type BlockData struct {
	BlockNum uint64
	Data     []byte
	Err      error
}

// WorkerReader is a reader instance for a single worker goroutine
// Each worker has its own file handles to avoid contention.
type WorkerReader struct {
	pr          *ParallelReader
	fileHandles map[uint64]*os.File
	mu          sync.Mutex
}

// NewWorkerReader creates a reader for a worker goroutine.
func (pr *ParallelReader) NewWorkerReader() *WorkerReader {
	return &WorkerReader{
		pr:          pr,
		fileHandles: make(map[uint64]*os.File),
	}
}

// Close closes all file handles for this worker.
func (wr *WorkerReader) Close() error {
	wr.mu.Lock()
	defer wr.mu.Unlock()

	var firstErr error
	for _, f := range wr.fileHandles {
		if err := f.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	wr.fileHandles = make(map[uint64]*os.File)
	return firstErr
}

// ReadItem reads the raw receipts data for a block.
func (wr *WorkerReader) ReadItem(blockNum uint64) ([]byte, error) {
	if blockNum >= wr.pr.itemCount {
		return nil, fmt.Errorf("block %d >= item count %d", blockNum, wr.pr.itemCount)
	}

	// Get offsets directly from memory (no syscall!)
	offset := wr.pr.getOffset(blockNum)

	var length uint64
	if blockNum+1 < wr.pr.itemCount {
		nextOffset := wr.pr.getOffset(blockNum + 1)
		if nextOffset < offset {
			return nil, fmt.Errorf("offset regression at block %d: %d -> %d", blockNum, offset, nextOffset)
		}
		length = nextOffset - offset
	} else {
		// Last block
		if len(wr.pr.cdatFiles) == 0 {
			return nil, errors.New("no cdat files available")
		}
		lastEntry := wr.pr.cdatFiles[len(wr.pr.cdatFiles)-1]
		endOffset := lastEntry.startOffset + lastEntry.size
		if offset > endOffset {
			return nil, fmt.Errorf("offset %d beyond data end %d", offset, endOffset)
		}
		length = endOffset - offset
	}

	if length == 0 {
		return nil, nil // Empty block
	}

	return wr.readBytes(offset, length, blockNum)
}

// dataEndOffset returns the global offset one past the last byte of data.
func (pr *ParallelReader) dataEndOffset() uint64 {
	if len(pr.cdatFiles) == 0 {
		return 0
	}
	last := pr.cdatFiles[len(pr.cdatFiles)-1]
	return last.startOffset + last.size
}

// readBytes reads data spanning potentially multiple cdat files.
func (wr *WorkerReader) readBytes(offset, length uint64, blockNum uint64) ([]byte, error) {
	// Bound the allocation by the data that actually exists: a corrupt index
	// must fail with an error, not make(2^48 bytes) and OOM the process.
	if end := wr.pr.dataEndOffset(); offset+length > end || offset+length < offset {
		return nil, fmt.Errorf("%w: block %d wants bytes [%d, %d) but data ends at %d",
			ErrReadBeyondBounds, blockNum, offset, offset+length, end)
	}
	result := make([]byte, length)
	remaining := length
	resultPos := uint64(0)
	currentOffset := offset

	for remaining > 0 {
		entry, err := wr.findCdatForOffset(currentOffset)
		if err != nil {
			return nil, fmt.Errorf("block %d: %w", blockNum, err)
		}

		fileOffset := currentOffset - entry.startOffset
		availableInFile := entry.size - fileOffset

		toRead := min(remaining, availableInFile)

		f, err := wr.getFileHandle(entry)
		if err != nil {
			return nil, err
		}

		readAt, err := fileReadOffset(fileOffset)
		if err != nil {
			return nil, fmt.Errorf("block %d in %s: %w", blockNum, entry.path, err)
		}
		dst := result[resultPos : resultPos+toRead]
		n, err := f.ReadAt(dst, readAt)
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("read error in %s at offset %d: %w", entry.path, fileOffset, err)
		}
		if n != len(dst) {
			return nil, fmt.Errorf("short read in %s: wanted %d, got %d", entry.path, toRead, n)
		}

		remaining -= toRead
		resultPos += toRead
		currentOffset += toRead
	}

	return result, nil
}

// findCdatForOffset finds which cdat file contains the offset.
func (wr *WorkerReader) findCdatForOffset(offset uint64) (*cdatEntry, error) {
	for _, entry := range wr.pr.cdatFiles {
		endOffset := entry.startOffset + entry.size
		if offset >= entry.startOffset && offset < endOffset {
			return entry, nil
		}
	}
	return nil, fmt.Errorf("%w: no cdat file for offset %d", ErrDataFileMissing, offset)
}

// getFileHandle gets or opens a file handle for this worker.
func (wr *WorkerReader) getFileHandle(entry *cdatEntry) (*os.File, error) {
	wr.mu.Lock()
	defer wr.mu.Unlock()

	if f, ok := wr.fileHandles[entry.index]; ok {
		return f, nil
	}

	f, err := os.Open(entry.path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", entry.path, err)
	}

	wr.fileHandles[entry.index] = f
	return f, nil
}

// CdatFileInfo returns info about cdat files.
func (pr *ParallelReader) CdatFileInfo() []string {
	info := make([]string, 0, len(pr.cdatFiles))
	for _, entry := range pr.cdatFiles {
		info = append(info, fmt.Sprintf("index=%d path=%s startOffset=%d size=%d",
			entry.index, entry.path, entry.startOffset, entry.size))
	}
	return info
}

// MaxAvailableBlock returns the highest block number that has data available.
func (pr *ParallelReader) MaxAvailableBlock() uint64 {
	if len(pr.cdatFiles) == 0 {
		return 0
	}

	// Find the last cdat file's end offset
	lastEntry := pr.cdatFiles[len(pr.cdatFiles)-1]
	maxDataOffset := lastEntry.startOffset + lastEntry.size

	// Binary search to find the last block that fits within available data
	low, high := uint64(0), pr.itemCount
	for low < high {
		mid := (low + high) / 2
		offset := pr.getOffset(mid)
		if offset < maxDataOffset {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low > 0 {
		return low - 1
	}
	return 0
}

// ReadOffsetBatch reads multiple offsets efficiently (for debugging/validation).
func (pr *ParallelReader) ReadOffsetBatch(start, count uint64) []uint64 {
	offsets := make([]uint64, count)
	for i := uint64(0); i < count && start+i < pr.itemCount; i++ {
		offsets[i] = pr.getOffset(start + i)
	}
	return offsets
}

// GetIndexStats returns statistics about the index.
func (pr *ParallelReader) GetIndexStats() (totalBlocks uint64, indexSizeMB float64, cdatFiles int) {
	return pr.itemCount, float64(len(pr.indexData)) / (1024 * 1024), len(pr.cdatFiles)
}
