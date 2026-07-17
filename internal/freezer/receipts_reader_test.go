package freezerscanner

import (
	"errors"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUint48ToBytes(t *testing.T) {
	tests := []struct {
		input    uint64
		expected []byte
	}{
		{0, []byte{0, 0, 0, 0, 0, 0}},
		{1, []byte{0, 0, 0, 0, 0, 1}},
		{255, []byte{0, 0, 0, 0, 0, 255}},
		{256, []byte{0, 0, 0, 0, 1, 0}},
		{0x123456789ABC, []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}},
		{0xFFFFFFFFFFFF, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
	}

	for _, tt := range tests {
		got := Uint48ToBytes(tt.input)
		if len(got) != 6 {
			t.Errorf("Uint48ToBytes(%d) returned %d bytes, want 6", tt.input, len(got))
			continue
		}
		for i := range got {
			if got[i] != tt.expected[i] {
				t.Errorf("Uint48ToBytes(%d) = %v, want %v", tt.input, got, tt.expected)
				break
			}
		}
	}
}

// TestFileReadOffset pins the checked uint64→int64 conversion feeding
// os.File.ReadAt: in-range offsets pass through exactly, and values beyond
// int64 (corrupt bookkeeping — real offsets are bounded by the 2GB chunk
// size) are rejected instead of wrapping into a negative seek.
func TestFileReadOffset(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		in   uint64
		want int64
	}{
		{0, 0},
		{1, 1},
		{DefaultChunkSize, int64(DefaultChunkSize)},
		{math.MaxInt64, math.MaxInt64},
	} {
		got, err := fileReadOffset(tc.in)
		if err != nil {
			t.Errorf("fileReadOffset(%d): unexpected error %v", tc.in, err)
			continue
		}
		if got != tc.want {
			t.Errorf("fileReadOffset(%d) = %d, want %d", tc.in, got, tc.want)
		}
	}

	for _, in := range []uint64{math.MaxInt64 + 1, math.MaxUint64} {
		if _, err := fileReadOffset(in); err == nil {
			t.Errorf("fileReadOffset(%d): want overflow error, got nil", in)
		}
	}
}

func TestIndexEntryParsing(t *testing.T) {
	// Create a temporary directory with mock index file
	tmpDir := t.TempDir()
	cidxPath := filepath.Join(tmpDir, "receipts.cidx")
	cdatDir := filepath.Join(tmpDir, "ancient")
	if err := os.MkdirAll(cdatDir, 0o750); err != nil {
		t.Fatalf("Failed to create cdat dir: %v", err)
	}

	// Create mock cidx file with 10 entries
	// Each entry is 6 bytes, offsets: 0, 10, 20, 30, 40, 50, 60, 70, 80, 90
	cidxData := make([]byte, 0, 10*IndexEntrySize)
	for i := range 10 {
		offset := uint64(i * 10)
		cidxData = append(cidxData, Uint48ToBytes(offset)...)
	}
	if err := os.WriteFile(cidxPath, cidxData, 0o600); err != nil {
		t.Fatalf("Failed to create cidx file: %v", err)
	}

	// Create mock cdat file
	cdatPath := filepath.Join(cdatDir, "receipts.0000.cdat")
	cdatData := make([]byte, 100)
	for i := range cdatData {
		cdatData[i] = byte(i)
	}
	if err := os.WriteFile(cdatPath, cdatData, 0o600); err != nil {
		t.Fatalf("Failed to create cdat file: %v", err)
	}
	// A negative-index cdat file must be skipped during discovery: parsed
	// into the unsigned index it is rejected, where it would previously
	// wrap into an astronomical startOffset shadowing real data.
	if err := os.WriteFile(filepath.Join(cdatDir, "receipts.-1.cdat"), []byte{1}, 0o600); err != nil {
		t.Fatalf("Failed to create negative-index cdat file: %v", err)
	}

	// Open reader
	reader, err := NewFreezerReader(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}
	defer func() { _ = reader.Close() }()

	if got := len(reader.cdatFiles); got != 1 {
		t.Errorf("indexed %d cdat files, want 1 (negative index skipped)", got)
	}

	// Verify item count
	if reader.ItemCount() != 10 {
		t.Errorf("ItemCount() = %d, want 10", reader.ItemCount())
	}

	// Verify offsets
	for i := range uint64(10) {
		offset, err := reader.ReadOffset(i)
		if err != nil {
			t.Errorf("ReadOffset(%d) error: %v", i, err)
			continue
		}
		expected := i * 10
		if offset != expected {
			t.Errorf("ReadOffset(%d) = %d, want %d", i, offset, expected)
		}
	}

	// Test out of bounds
	_, err = reader.ReadOffset(10)
	if err == nil {
		t.Error("ReadOffset(10) should return error for out of bounds")
	}
}

// writeFreezerFixture lays out a receipts.cidx with the given offsets and
// one receipts.0000.cdat of dataSize bytes, returning the directory.
func writeFreezerFixture(t *testing.T, offsets []uint64, dataSize int) string {
	t.Helper()
	dir := t.TempDir()
	cidxData := make([]byte, 0, len(offsets)*IndexEntrySize)
	for _, offset := range offsets {
		cidxData = append(cidxData, Uint48ToBytes(offset)...)
	}
	if err := os.WriteFile(filepath.Join(dir, "receipts.cidx"), cidxData, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "receipts.0000.cdat"), make([]byte, dataSize), 0o600); err != nil {
		t.Fatal(err)
	}
	return dir
}

// TestFreezerReaderLastBlockReadsToDataEnd pins the last-block branch: with
// no next index entry, the read extends to the end of the final cdat file.
func TestFreezerReaderLastBlockReadsToDataEnd(t *testing.T) {
	dir := writeFreezerFixture(t, []uint64{0, 6}, 10)
	reader, err := NewFreezerReader(dir)
	if err != nil {
		t.Fatalf("NewFreezerReader: %v", err)
	}
	defer func() { _ = reader.Close() }()

	data, err := reader.ReadItem(1)
	if err != nil {
		t.Fatalf("ReadItem(last): %v", err)
	}
	if len(data) != 4 { // bytes [6, 10)
		t.Errorf("last block read %d bytes, want 4", len(data))
	}
}

// TestFreezerReaderLastBlockBeyondDataEnd pins the corrupt-index guard on
// the last-block branch: an offset past the data end must error, not read.
func TestFreezerReaderLastBlockBeyondDataEnd(t *testing.T) {
	dir := writeFreezerFixture(t, []uint64{0, 20}, 10)
	reader, err := NewFreezerReader(dir)
	if err != nil {
		t.Fatalf("NewFreezerReader: %v", err)
	}
	defer func() { _ = reader.Close() }()

	if _, err := reader.ReadItem(1); !errors.Is(err, ErrInvalidOffset) {
		t.Fatalf("ReadItem(beyond end) = %v, want ErrInvalidOffset", err)
	}
}

// TestFreezerReaderShortReadOnTruncatedFile pins the short-read guard: a
// cdat file that shrinks after indexing (the entry keeps the stat-time
// size) must produce a loud error instead of returning partial data.
func TestFreezerReaderShortReadOnTruncatedFile(t *testing.T) {
	dir := writeFreezerFixture(t, []uint64{0, 8}, 10)
	reader, err := NewFreezerReader(dir)
	if err != nil {
		t.Fatalf("NewFreezerReader: %v", err)
	}
	defer func() { _ = reader.Close() }()

	if err := os.Truncate(filepath.Join(dir, "receipts.0000.cdat"), 4); err != nil {
		t.Fatal(err)
	}
	if _, err := reader.ReadItem(0); err == nil || !strings.Contains(err.Error(), "short read") {
		t.Fatalf("ReadItem(truncated) = %v, want short-read error", err)
	}
}

func TestValidateIndexRange(t *testing.T) {
	tmpDir := t.TempDir()
	cidxPath := filepath.Join(tmpDir, "receipts.cidx")
	cdatDir := filepath.Join(tmpDir, "ancient")
	if err := os.MkdirAll(cdatDir, 0o750); err != nil {
		t.Fatalf("Failed to create cdat dir: %v", err)
	}

	// Create valid monotonic index
	offsets := []uint64{0, 10, 20, 30, 40, 50}
	cidxData := make([]byte, 0, len(offsets)*IndexEntrySize)
	for _, offset := range offsets {
		cidxData = append(cidxData, Uint48ToBytes(offset)...)
	}
	if err := os.WriteFile(cidxPath, cidxData, 0o600); err != nil {
		t.Fatalf("Failed to create cidx file: %v", err)
	}

	// Create mock cdat file
	cdatPath := filepath.Join(cdatDir, "receipts.0000.cdat")
	if err := os.WriteFile(cdatPath, make([]byte, 100), 0o600); err != nil {
		t.Fatalf("Failed to create cdat file: %v", err)
	}

	reader, err := NewFreezerReader(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}
	defer func() { _ = reader.Close() }()

	// Validate should succeed for monotonic offsets
	if err := reader.ValidateIndexRange(0, 6); err != nil {
		t.Errorf("ValidateIndexRange failed for valid data: %v", err)
	}
}

func TestReadItem(t *testing.T) {
	tmpDir := t.TempDir()
	cidxPath := filepath.Join(tmpDir, "receipts.cidx")
	cdatDir := filepath.Join(tmpDir, "ancient")
	if err := os.MkdirAll(cdatDir, 0o750); err != nil {
		t.Fatalf("Failed to create cdat dir: %v", err)
	}

	// Create index with offsets: 0, 5, 15, 30
	offsets := []uint64{0, 5, 15, 30}
	cidxData := make([]byte, 0, len(offsets)*IndexEntrySize)
	for _, offset := range offsets {
		cidxData = append(cidxData, Uint48ToBytes(offset)...)
	}
	if err := os.WriteFile(cidxPath, cidxData, 0o600); err != nil {
		t.Fatalf("Failed to create cidx file: %v", err)
	}

	// Create cdat file with data
	// Block 0: bytes 0-4 (5 bytes): [0,1,2,3,4]
	// Block 1: bytes 5-14 (10 bytes): [5,6,7,8,9,10,11,12,13,14]
	// Block 2: bytes 15-29 (15 bytes): [15,16,...,29]
	cdatPath := filepath.Join(cdatDir, "receipts.0000.cdat")
	cdatData := make([]byte, 30)
	for i := range cdatData {
		cdatData[i] = byte(i)
	}
	if err := os.WriteFile(cdatPath, cdatData, 0o600); err != nil {
		t.Fatalf("Failed to create cdat file: %v", err)
	}

	reader, err := NewFreezerReader(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}
	defer func() { _ = reader.Close() }()

	// Test reading each block
	tests := []struct {
		block    uint64
		expected []byte
	}{
		{0, []byte{0, 1, 2, 3, 4}},
		{1, []byte{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
		{2, []byte{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}},
	}

	for _, tt := range tests {
		data, err := reader.ReadItem(tt.block)
		if err != nil {
			t.Errorf("ReadItem(%d) error: %v", tt.block, err)
			continue
		}
		if len(data) != len(tt.expected) {
			t.Errorf("ReadItem(%d) returned %d bytes, want %d", tt.block, len(data), len(tt.expected))
			continue
		}
		for i := range data {
			if data[i] != tt.expected[i] {
				t.Errorf("ReadItem(%d)[%d] = %d, want %d", tt.block, i, data[i], tt.expected[i])
				break
			}
		}
	}
}

// TestReadItemCorruptIndexHugeOffset pins the fix for a crash found by
// FuzzFreezerIndexRead: an index entry pointing far beyond the data files made
// ReadItem allocate up to 2^48 bytes (length = nextOffset - offset) and abort
// the whole process with OOM. Corrupt indexes must yield an error instead.
func TestReadItemCorruptIndexHugeOffset(t *testing.T) {
	tmpDir := t.TempDir()

	// Block 0 spans [0, 2^48-1) according to the corrupt index.
	cidxData := make([]byte, 0, 2*IndexEntrySize)
	for _, offset := range []uint64{0, 0xFFFFFFFFFFFF} {
		cidxData = append(cidxData, Uint48ToBytes(offset)...)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "receipts.cidx"), cidxData, 0o600); err != nil {
		t.Fatalf("Failed to create cidx file: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "receipts.0000.cdat"), make([]byte, 10), 0o600); err != nil {
		t.Fatalf("Failed to create cdat file: %v", err)
	}

	reader, err := NewFreezerReader(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}
	defer func() { _ = reader.Close() }()

	if _, err := reader.ReadItem(0); err == nil {
		t.Fatal("ReadItem(0) must fail when the index points beyond the data end")
	}

	// Same guard for the parallel/worker reader path.
	pr, err := NewParallelReader(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create parallel reader: %v", err)
	}
	wr := pr.NewWorkerReader()
	defer func() { _ = wr.Close() }()
	if _, err := wr.ReadItem(0); err == nil {
		t.Fatal("WorkerReader.ReadItem(0) must fail when the index points beyond the data end")
	}
}

func TestEmptyBlock(t *testing.T) {
	tmpDir := t.TempDir()
	cidxPath := filepath.Join(tmpDir, "receipts.cidx")
	cdatDir := filepath.Join(tmpDir, "ancient")
	if err := os.MkdirAll(cdatDir, 0o750); err != nil {
		t.Fatalf("Failed to create cdat dir: %v", err)
	}

	// Create index with repeated offset (empty block at index 1)
	// Block 0: offset 0-5 (5 bytes)
	// Block 1: offset 5-5 (0 bytes - empty!)
	// Block 2: offset 5-10 (5 bytes)
	offsets := []uint64{0, 5, 5, 10}
	cidxData := make([]byte, 0, len(offsets)*IndexEntrySize)
	for _, offset := range offsets {
		cidxData = append(cidxData, Uint48ToBytes(offset)...)
	}
	if err := os.WriteFile(cidxPath, cidxData, 0o600); err != nil {
		t.Fatalf("Failed to create cidx file: %v", err)
	}

	cdatPath := filepath.Join(cdatDir, "receipts.0000.cdat")
	if err := os.WriteFile(cdatPath, make([]byte, 15), 0o600); err != nil {
		t.Fatalf("Failed to create cdat file: %v", err)
	}

	reader, err := NewFreezerReader(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}
	defer func() { _ = reader.Close() }()

	// Block 1 should return empty data (nil or empty slice)
	data, err := reader.ReadItem(1)
	if err != nil {
		t.Errorf("ReadItem(1) error: %v", err)
	}
	if len(data) != 0 {
		t.Errorf("ReadItem(1) returned %d bytes, want 0 (empty block)", len(data))
	}

	// Block 2 should return 5 bytes
	data2, err := reader.ReadItem(2)
	if err != nil {
		t.Errorf("ReadItem(2) error: %v", err)
	}
	if len(data2) != 5 {
		t.Errorf("ReadItem(2) returned %d bytes, want 5", len(data2))
	}
}
