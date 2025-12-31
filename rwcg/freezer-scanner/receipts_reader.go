// Package freezer provides a reader for geth freezer receipts data
// It reads the .cidx index file and .cdat data files to retrieve receipts
// for specific block numbers without requiring a running geth node.
package freezerscanner

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

const (
	// IndexEntrySize is the size of each entry in the .cidx file (6 bytes)
	IndexEntrySize = 6

	// DefaultChunkSize is the typical size of each .cdat file (~2GB)
	DefaultChunkSize = uint64(2_000_000_000)
)

var (
	ErrBlockNotFound     = errors.New("block not found in index")
	ErrInvalidIndex      = errors.New("invalid index file")
	ErrInvalidOffset     = errors.New("invalid offset in index")
	ErrDataFileMissing   = errors.New("data file missing")
	ErrReadBeyondBounds  = errors.New("read beyond file bounds")
)

// FreezerReader provides access to geth freezer receipts data
type FreezerReader struct {
	cidxPath    string        // Path to receipts.cidx
	cdatDir     string        // Directory containing receipts.*.cdat files
	cidxFile    *os.File      // Open handle to cidx
	cidxSize    int64         // Size of cidx file
	itemCount   uint64        // Number of blocks in the index
	chunkSize   uint64        // Size of each cdat chunk
	cdatFiles   []*cdatEntry  // Sorted list of cdat files with their ranges
	fileHandles map[int]*os.File
	mu          sync.RWMutex
}

// cdatEntry represents a .cdat file with its global offset range
type cdatEntry struct {
	index       int    // File index (e.g., 6 for receipts.0006.cdat)
	path        string
	startOffset uint64 // Global offset where this file starts
	size        int64  // Size of this file
}

// NewFreezerReader creates a new freezer reader
// ancientDir should be the directory containing receipts.cidx and optionally an "ancient" subdirectory
func NewFreezerReader(ancientDir string) (*FreezerReader, error) {
	fr := &FreezerReader{
		chunkSize:   DefaultChunkSize,
		fileHandles: make(map[int]*os.File),
	}

	// Try to find cidx file
	cidxPath := filepath.Join(ancientDir, "receipts.cidx")
	if _, err := os.Stat(cidxPath); err != nil {
		// Try in ancient subdirectory
		cidxPath = filepath.Join(ancientDir, "ancient", "receipts.cidx")
		if _, err := os.Stat(cidxPath); err != nil {
			return nil, fmt.Errorf("receipts.cidx not found in %s: %w", ancientDir, err)
		}
	}
	fr.cidxPath = cidxPath

	// Find cdat directory
	cdatDir := filepath.Dir(cidxPath)
	ancientSubdir := filepath.Join(cdatDir, "ancient")
	if info, err := os.Stat(ancientSubdir); err == nil && info.IsDir() {
		cdatDir = ancientSubdir
	}
	fr.cdatDir = cdatDir

	// Open and validate cidx file
	cidxFile, err := os.Open(cidxPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open cidx: %w", err)
	}
	fr.cidxFile = cidxFile

	info, err := cidxFile.Stat()
	if err != nil {
		cidxFile.Close()
		return nil, fmt.Errorf("failed to stat cidx: %w", err)
	}
	fr.cidxSize = info.Size()

	if fr.cidxSize%IndexEntrySize != 0 {
		cidxFile.Close()
		return nil, fmt.Errorf("%w: cidx size %d not divisible by %d", ErrInvalidIndex, fr.cidxSize, IndexEntrySize)
	}
	fr.itemCount = uint64(fr.cidxSize / IndexEntrySize)

	// Discover and index cdat files
	if err := fr.indexCdatFiles(); err != nil {
		cidxFile.Close()
		return nil, fmt.Errorf("failed to index cdat files: %w", err)
	}

	return fr, nil
}

// indexCdatFiles discovers all .cdat files and builds the offset map
func (fr *FreezerReader) indexCdatFiles() error {
	pattern := filepath.Join(fr.cdatDir, "receipts.*.cdat")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("glob error: %w", err)
	}

	if len(matches) == 0 {
		return fmt.Errorf("no cdat files found matching %s", pattern)
	}

	// Parse file indices and build entries
	var entries []*cdatEntry
	for _, path := range matches {
		var index int
		base := filepath.Base(path)
		// Parse receipts.XXXX.cdat format
		if _, err := fmt.Sscanf(base, "receipts.%d.cdat", &index); err != nil {
			continue // Skip files that don't match the pattern
		}

		info, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("failed to stat %s: %w", path, err)
		}

		entries = append(entries, &cdatEntry{
			index: index,
			path:  path,
			size:  info.Size(),
		})
	}

	// Sort by index
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].index < entries[j].index
	})

	// Calculate global offsets for each file
	// Note: In geth's freezer, file index * chunkSize gives the starting offset
	for _, entry := range entries {
		entry.startOffset = uint64(entry.index) * fr.chunkSize
	}

	fr.cdatFiles = entries
	return nil
}

// ItemCount returns the number of blocks indexed
func (fr *FreezerReader) ItemCount() uint64 {
	return fr.itemCount
}

// ReadOffset reads the offset for a given block number from the index
func (fr *FreezerReader) ReadOffset(blockNum uint64) (uint64, error) {
	if blockNum >= fr.itemCount {
		return 0, fmt.Errorf("%w: block %d >= item count %d", ErrBlockNotFound, blockNum, fr.itemCount)
	}

	buf := make([]byte, IndexEntrySize)
	pos := int64(blockNum) * IndexEntrySize

	n, err := fr.cidxFile.ReadAt(buf, pos)
	if err != nil && err != io.EOF {
		return 0, fmt.Errorf("failed to read index at block %d: %w", blockNum, err)
	}
	if n != IndexEntrySize {
		return 0, fmt.Errorf("short read at block %d: got %d bytes", blockNum, n)
	}

	// Index entries are 6 bytes big-endian (48-bit offset)
	offset := uint64(buf[0])<<40 | uint64(buf[1])<<32 | uint64(buf[2])<<24 |
		uint64(buf[3])<<16 | uint64(buf[4])<<8 | uint64(buf[5])

	return offset, nil
}

// ReadItem reads the raw receipts blob for a given block number
func (fr *FreezerReader) ReadItem(blockNum uint64) ([]byte, error) {
	// Get offset for this block
	offset, err := fr.ReadOffset(blockNum)
	if err != nil {
		return nil, err
	}

	// Get offset for next block (or use end of data)
	var length uint64
	if blockNum+1 < fr.itemCount {
		nextOffset, err := fr.ReadOffset(blockNum + 1)
		if err != nil {
			return nil, fmt.Errorf("failed to read next offset: %w", err)
		}
		if nextOffset < offset {
			return nil, fmt.Errorf("%w: offset regression at block %d (%d -> %d)", ErrInvalidOffset, blockNum, offset, nextOffset)
		}
		length = nextOffset - offset
	} else {
		// Last block - read to end of last cdat file
		lastEntry := fr.cdatFiles[len(fr.cdatFiles)-1]
		endOffset := lastEntry.startOffset + uint64(lastEntry.size)
		if offset > endOffset {
			return nil, fmt.Errorf("%w: offset %d beyond data end %d", ErrInvalidOffset, offset, endOffset)
		}
		length = endOffset - offset
	}

	if length == 0 {
		return nil, nil // Empty receipts (no transactions in block)
	}

	// Read the data
	return fr.readBytes(offset, length, blockNum)
}

// readBytes reads a range of bytes from the cdat files, handling boundary spanning
func (fr *FreezerReader) readBytes(offset, length uint64, blockNum uint64) ([]byte, error) {
	result := make([]byte, length)
	remaining := length
	resultPos := uint64(0)
	currentOffset := offset

	for remaining > 0 {
		// Find which file contains this offset
		entry, err := fr.findCdatForOffset(currentOffset)
		if err != nil {
			return nil, fmt.Errorf("block %d: %w", blockNum, err)
		}

		// Calculate position within this file
		fileOffset := currentOffset - entry.startOffset
		availableInFile := uint64(entry.size) - fileOffset

		// Determine how much to read from this file
		toRead := remaining
		if toRead > availableInFile {
			toRead = availableInFile
		}

		// Get file handle
		f, err := fr.getFileHandle(entry)
		if err != nil {
			return nil, err
		}

		// Read the data
		n, err := f.ReadAt(result[resultPos:resultPos+toRead], int64(fileOffset))
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("read error in %s at offset %d: %w", entry.path, fileOffset, err)
		}
		if uint64(n) != toRead {
			return nil, fmt.Errorf("short read in %s: wanted %d, got %d", entry.path, toRead, n)
		}

		remaining -= toRead
		resultPos += toRead
		currentOffset += toRead
	}

	return result, nil
}

// findCdatForOffset finds the cdat file containing the given global offset
func (fr *FreezerReader) findCdatForOffset(offset uint64) (*cdatEntry, error) {
	fr.mu.RLock()
	defer fr.mu.RUnlock()

	for _, entry := range fr.cdatFiles {
		endOffset := entry.startOffset + uint64(entry.size)
		if offset >= entry.startOffset && offset < endOffset {
			return entry, nil
		}
	}

	return nil, fmt.Errorf("%w: no cdat file for offset %d", ErrDataFileMissing, offset)
}

// getFileHandle gets or opens a file handle for the given cdat entry
func (fr *FreezerReader) getFileHandle(entry *cdatEntry) (*os.File, error) {
	fr.mu.Lock()
	defer fr.mu.Unlock()

	if f, ok := fr.fileHandles[entry.index]; ok {
		return f, nil
	}

	f, err := os.Open(entry.path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", entry.path, err)
	}

	fr.fileHandles[entry.index] = f
	return f, nil
}

// Close closes all open file handles
func (fr *FreezerReader) Close() error {
	fr.mu.Lock()
	defer fr.mu.Unlock()

	var firstErr error

	if fr.cidxFile != nil {
		if err := fr.cidxFile.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}

	for _, f := range fr.fileHandles {
		if err := f.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}

	fr.fileHandles = make(map[int]*os.File)
	return firstErr
}

// ValidateIndexRange validates that offsets are monotonically non-decreasing
// for the given block range. Returns an error if any regression is found.
func (fr *FreezerReader) ValidateIndexRange(startBlock, endBlock uint64) error {
	if endBlock > fr.itemCount {
		endBlock = fr.itemCount
	}
	if startBlock >= endBlock {
		return nil
	}

	prevOffset, err := fr.ReadOffset(startBlock)
	if err != nil {
		return err
	}

	for block := startBlock + 1; block < endBlock; block++ {
		offset, err := fr.ReadOffset(block)
		if err != nil {
			return err
		}
		if offset < prevOffset {
			return fmt.Errorf("offset regression at block %d: %d < %d", block, offset, prevOffset)
		}
		prevOffset = offset
	}

	return nil
}

// CdatFileInfo returns information about discovered cdat files
func (fr *FreezerReader) CdatFileInfo() []string {
	fr.mu.RLock()
	defer fr.mu.RUnlock()

	var info []string
	for _, entry := range fr.cdatFiles {
		info = append(info, fmt.Sprintf("index=%d path=%s startOffset=%d size=%d",
			entry.index, entry.path, entry.startOffset, entry.size))
	}
	return info
}

// Uint48ToBytes converts a 48-bit value to 6 bytes (for testing)
func Uint48ToBytes(v uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, v)
	return buf[2:] // Return last 6 bytes
}

