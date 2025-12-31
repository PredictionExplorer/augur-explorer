// Package freezer provides a sequential reader for geth freezer receipts data
// This reader scans .cdat files directly without relying on the .cidx index
package freezerscanner

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/golang/snappy"
)

// SequentialReader reads receipts sequentially from .cdat files
type SequentialReader struct {
	cdatDir   string
	cdatFiles []string
	currentFile *os.File
	currentIdx  int
	fileOffset  int64
	blockCount  uint64
}

// NewSequentialReader creates a reader that scans cdat files sequentially
func NewSequentialReader(ancientDir string) (*SequentialReader, error) {
	// Find cdat directory
	cdatDir := ancientDir
	ancientSubdir := filepath.Join(cdatDir, "ancient")
	if info, err := os.Stat(ancientSubdir); err == nil && info.IsDir() {
		cdatDir = ancientSubdir
	}

	// Find all cdat files
	pattern := filepath.Join(cdatDir, "receipts.*.cdat")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("glob error: %w", err)
	}
	if len(matches) == 0 {
		return nil, fmt.Errorf("no cdat files found in %s", cdatDir)
	}

	// Sort files by index
	sort.Strings(matches)

	return &SequentialReader{
		cdatDir:    cdatDir,
		cdatFiles:  matches,
		currentIdx: -1,
	}, nil
}

// CdatFiles returns the list of discovered cdat files
func (sr *SequentialReader) CdatFiles() []string {
	return sr.cdatFiles
}

// OpenFile opens a specific cdat file for reading
func (sr *SequentialReader) OpenFile(index int) error {
	if index < 0 || index >= len(sr.cdatFiles) {
		return fmt.Errorf("file index %d out of range (0-%d)", index, len(sr.cdatFiles)-1)
	}

	if sr.currentFile != nil {
		sr.currentFile.Close()
	}

	f, err := os.Open(sr.cdatFiles[index])
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", sr.cdatFiles[index], err)
	}

	sr.currentFile = f
	sr.currentIdx = index
	sr.fileOffset = 0
	return nil
}

// ReadNextBlock reads the next block's raw data from the current file
// Returns the raw data (after header), or io.EOF when file is exhausted
func (sr *SequentialReader) ReadNextBlock() ([]byte, error) {
	if sr.currentFile == nil {
		return nil, fmt.Errorf("no file open")
	}

	// Read header bytes using ReadAt for precise positioning
	header := make([]byte, 20)
	n, err := sr.currentFile.ReadAt(header, sr.fileOffset)
	if err == io.EOF || n == 0 {
		return nil, io.EOF
	}
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("read header error at offset %d: %w", sr.fileOffset, err)
	}

	// Parse varint (uncompressed/RLP size)
	uncompSize, varintLen := decodeVarintBytes(header)
	if varintLen == 0 || uncompSize == 0 {
		return nil, io.EOF
	}
	if uncompSize > 100_000_000 { // 100MB max
		return nil, fmt.Errorf("invalid size: %d at offset %d", uncompSize, sr.fileOffset)
	}

	// Check what comes after varint
	// If byte after varint is 0xc0-0xff, it's RLP (uncompressed with 2-byte header)
	// If not, it's snappy compressed data
	afterVarint := header[varintLen]
	
	if afterVarint >= 0xc0 || (varintLen+2 < n && header[varintLen+2] >= 0xc0) {
		// Uncompressed format: varint + 2-byte header + RLP
		headerSize := varintLen + 2
		
		// Read RLP data
		rlpData := make([]byte, uncompSize)
		rlpOffset := sr.fileOffset + int64(headerSize)
		bytesRead, err := sr.currentFile.ReadAt(rlpData, rlpOffset)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("read RLP data error at offset %d: %w", rlpOffset, err)
		}
		if bytesRead < int(uncompSize) {
			return nil, fmt.Errorf("short read: wanted %d bytes, got %d at offset %d", uncompSize, bytesRead, rlpOffset)
		}

		sr.fileOffset = rlpOffset + int64(uncompSize)
		sr.blockCount++
		return rlpData, nil
	}

	// Snappy compressed format: varint (uncomp size) + compressed data
	// Need to find where compressed data ends
	// Snappy has its own length encoding - read and decode
	compressedStart := sr.fileOffset + int64(varintLen)
	
	// Read a large enough buffer to hold compressed data (estimate: 2x uncomp size for safety)
	estimatedCompSize := int(uncompSize) // compressed is usually smaller
	if estimatedCompSize < 10000 {
		estimatedCompSize = 10000
	}
	compData := make([]byte, estimatedCompSize)
	bytesRead, err := sr.currentFile.ReadAt(compData, compressedStart)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("read compressed data error at offset %d: %w", compressedStart, err)
	}
	
	// Try to decode snappy - it will only consume what it needs
	decoded, err := snappy.Decode(nil, compData[:bytesRead])
	if err != nil {
		return nil, fmt.Errorf("snappy decode failed at offset %d: %w", sr.fileOffset, err)
	}
	
	// Calculate actual compressed size by re-encoding (snappy doesn't expose consumed bytes)
	actualCompLen := len(snappy.Encode(nil, decoded))
	
	sr.fileOffset = compressedStart + int64(actualCompLen)
	sr.blockCount++
	
	return decoded, nil
}

// BlockCount returns number of blocks read so far
func (sr *SequentialReader) BlockCount() uint64 {
	return sr.blockCount
}

// FileOffset returns current position in the file
func (sr *SequentialReader) FileOffset() int64 {
	return sr.fileOffset
}

// CurrentFile returns path of current file being read
func (sr *SequentialReader) CurrentFile() string {
	if sr.currentIdx >= 0 && sr.currentIdx < len(sr.cdatFiles) {
		return sr.cdatFiles[sr.currentIdx]
	}
	return ""
}

// Close closes any open file
func (sr *SequentialReader) Close() error {
	if sr.currentFile != nil {
		err := sr.currentFile.Close()
		sr.currentFile = nil
		return err
	}
	return nil
}

// decodeVarintBytes decodes a varint from bytes
func decodeVarintBytes(data []byte) (uint64, int) {
	var value uint64
	for i := 0; i < len(data) && i < 10; i++ {
		b := data[i]
		value |= uint64(b&0x7f) << (7 * i)
		if b&0x80 == 0 {
			return value, i + 1
		}
	}
	return value, 0
}

