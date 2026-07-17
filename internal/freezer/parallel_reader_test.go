package freezerscanner

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writeCidx writes a receipts.cidx file with one 6-byte big-endian entry per
// offset, mirroring the on-disk geth freezer index format.
func writeCidx(t *testing.T, dir string, offsets []uint64) {
	t.Helper()
	index := make([]byte, 0, len(offsets)*IndexEntrySize)
	for _, off := range offsets {
		index = append(index, Uint48ToBytes(off)...)
	}
	if err := os.WriteFile(filepath.Join(dir, "receipts.cidx"), index, 0o600); err != nil {
		t.Fatal(err)
	}
}

func writeCdat(t *testing.T, dir string, index int, data []byte) {
	t.Helper()
	name := filepath.Join(dir, fmt.Sprintf("receipts.%04d.cdat", index))
	if err := os.WriteFile(name, data, 0o600); err != nil {
		t.Fatal(err)
	}
}

// seqBytes returns n sequential bytes starting at start, so tests can assert
// exactly which byte range a read returned.
func seqBytes(start, n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte((start + i) & 0xff)
	}
	return b
}

// newSmallChunkReader builds a ParallelReader over the fixture directory with
// a tiny chunk size, so multi-cdat spanning reads are exercisable with
// kilobyte fixtures instead of geth's 2GB chunks. It goes through the same
// index-loading and cdat-discovery code as NewParallelReader.
func newSmallChunkReader(t *testing.T, dir string, chunkSize uint64) *ParallelReader {
	t.Helper()
	indexData, err := os.ReadFile(filepath.Join(dir, "receipts.cidx")) //nolint:gosec // test fixture path under t.TempDir
	if err != nil {
		t.Fatal(err)
	}
	pr := &ParallelReader{
		chunkSize: chunkSize,
		cdatDir:   dir,
		indexData: indexData,
		itemCount: uint64(len(indexData) / IndexEntrySize),
	}
	if err := pr.indexCdatFiles(); err != nil {
		t.Fatal(err)
	}
	return pr
}

// TestParallelReaderReadItem covers the production read path of freezer-scan:
// per-block reads through WorkerReader.ReadItem, including a block spanning
// two cdat files, an empty block, the final block (which reads to the data
// end) and file-handle caching.
func TestParallelReaderReadItem(t *testing.T) {
	dir := t.TempDir()
	const chunkSize = 16

	// Two cdat files: file 0 is exactly one chunk (16 bytes), file 1 has 8
	// bytes. Block 2 spans the file boundary.
	writeCdat(t, dir, 0, seqBytes(0, 16))
	writeCdat(t, dir, 1, seqBytes(16, 8))
	// Blocks: [0,5) [5,12) [12,20) [20,20)=empty [20,24)=last.
	writeCidx(t, dir, []uint64{0, 5, 12, 20, 20})

	pr := newSmallChunkReader(t, dir, chunkSize)
	wr := pr.NewWorkerReader()
	defer func() {
		if err := wr.Close(); err != nil {
			t.Errorf("Close: %v", err)
		}
	}()

	cases := []struct {
		block uint64
		want  []byte
	}{
		{0, seqBytes(0, 5)},
		{1, seqBytes(5, 7)},
		{2, seqBytes(12, 8)}, // spans receipts.0000.cdat -> receipts.0001.cdat
		{3, nil},             // empty block (offset equals next offset)
		{4, seqBytes(20, 4)}, // last block reads to the end of the data
	}
	for _, tc := range cases {
		got, err := wr.ReadItem(tc.block)
		if err != nil {
			t.Errorf("ReadItem(%d): %v", tc.block, err)
			continue
		}
		if string(got) != string(tc.want) {
			t.Errorf("ReadItem(%d) = %v, want %v", tc.block, got, tc.want)
		}
	}

	// The spanning read must have opened (and cached) both cdat handles.
	if len(wr.fileHandles) != 2 {
		t.Errorf("worker caches %d file handles, want 2", len(wr.fileHandles))
	}

	if _, err := wr.ReadItem(5); err == nil {
		t.Error("ReadItem beyond item count should fail")
	}

	// Close resets the cache; a subsequent read reopens handles.
	if err := wr.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if len(wr.fileHandles) != 0 {
		t.Errorf("Close left %d cached handles", len(wr.fileHandles))
	}
	if _, err := wr.ReadItem(0); err != nil {
		t.Errorf("ReadItem after Close: %v", err)
	}
}

func TestParallelReaderOffsetRegression(t *testing.T) {
	dir := t.TempDir()
	writeCdat(t, dir, 0, seqBytes(0, 16))
	writeCidx(t, dir, []uint64{0, 10, 5})

	pr := newSmallChunkReader(t, dir, 16)
	wr := pr.NewWorkerReader()
	defer func() { _ = wr.Close() }()

	if _, err := wr.ReadItem(1); err == nil || !strings.Contains(err.Error(), "offset regression") {
		t.Errorf("want offset regression error, got %v", err)
	}
}

func TestParallelReaderMissingCdatForOffset(t *testing.T) {
	dir := t.TempDir()
	// Only file 1 exists: offsets in chunk 0 have no covering cdat file,
	// while the data end (chunk 1 start + size) still bounds the read.
	writeCdat(t, dir, 1, seqBytes(0, 8))
	writeCidx(t, dir, []uint64{0, 4})

	pr := newSmallChunkReader(t, dir, 16)
	wr := pr.NewWorkerReader()
	defer func() { _ = wr.Close() }()

	if _, err := wr.ReadItem(0); !errors.Is(err, ErrDataFileMissing) {
		t.Errorf("want ErrDataFileMissing, got %v", err)
	}
}

// TestParallelReaderShortReadOnTruncatedFile pins the short-read guard: a
// cdat file that shrinks after indexing (the entry keeps the stat-time
// size) must produce a loud error instead of returning partial data.
func TestParallelReaderShortReadOnTruncatedFile(t *testing.T) {
	dir := t.TempDir()
	writeCdat(t, dir, 0, seqBytes(0, 10))
	writeCidx(t, dir, []uint64{0, 8})

	pr := newSmallChunkReader(t, dir, 16)
	wr := pr.NewWorkerReader()
	defer func() { _ = wr.Close() }()

	if err := os.Truncate(filepath.Join(dir, "receipts.0000.cdat"), 4); err != nil {
		t.Fatal(err)
	}
	if _, err := wr.ReadItem(0); err == nil || !strings.Contains(err.Error(), "short read") {
		t.Fatalf("ReadItem(truncated) = %v, want short-read error", err)
	}
}

func TestParallelReaderLastBlockBeyondData(t *testing.T) {
	dir := t.TempDir()
	writeCdat(t, dir, 0, seqBytes(0, 8))
	// Last block's offset lies beyond the end of the data.
	writeCidx(t, dir, []uint64{0, 20})

	pr := newSmallChunkReader(t, dir, 16)
	wr := pr.NewWorkerReader()
	defer func() { _ = wr.Close() }()

	if _, err := wr.ReadItem(1); err == nil || !strings.Contains(err.Error(), "beyond data end") {
		t.Errorf("want beyond-data-end error, got %v", err)
	}
}

// TestNewParallelReader covers the production constructor: cidx discovery in
// the root and ancient/ subdirectory, index validation and cdat indexing,
// plus the debug helpers freezer-scan --info prints.
func TestNewParallelReader(t *testing.T) {
	dir := t.TempDir()
	ancient := filepath.Join(dir, "ancient")
	if err := os.Mkdir(ancient, 0o750); err != nil {
		t.Fatal(err)
	}
	writeCidx(t, dir, []uint64{0, 4, 8})
	writeCdat(t, ancient, 0, seqBytes(0, 10))
	// A file that does not match the receipts.%d.cdat pattern is skipped.
	if err := os.WriteFile(filepath.Join(ancient, "receipts.junk.cdat"), []byte{1}, 0o600); err != nil {
		t.Fatal(err)
	}
	// A negative index is skipped too: it would otherwise wrap into an
	// astronomical startOffset (index is unsigned by construction).
	if err := os.WriteFile(filepath.Join(ancient, "receipts.-1.cdat"), []byte{1}, 0o600); err != nil {
		t.Fatal(err)
	}

	pr, err := NewParallelReader(dir)
	if err != nil {
		t.Fatalf("NewParallelReader: %v", err)
	}

	totalBlocks, indexSizeMB, cdatFiles := pr.GetIndexStats()
	if totalBlocks != 3 {
		t.Errorf("totalBlocks = %d, want 3", totalBlocks)
	}
	if wantMB := float64(3*IndexEntrySize) / (1024 * 1024); indexSizeMB != wantMB {
		t.Errorf("indexSizeMB = %v, want %v", indexSizeMB, wantMB)
	}
	if cdatFiles != 1 {
		t.Errorf("cdatFiles = %d, want 1 (junk and negative-index files must be skipped)", cdatFiles)
	}

	if got := pr.MaxAvailableBlock(); got != 2 {
		t.Errorf("MaxAvailableBlock() = %d, want 2", got)
	}

	offsets := pr.ReadOffsetBatch(0, 5)
	want := []uint64{0, 4, 8, 0, 0} // beyond-count entries stay zero
	for i, off := range offsets {
		if off != want[i] {
			t.Errorf("ReadOffsetBatch[%d] = %d, want %d", i, off, want[i])
		}
	}

	info := pr.CdatFileInfo()
	if len(info) != 1 || !strings.Contains(info[0], "receipts.0000.cdat") {
		t.Errorf("CdatFileInfo() = %v", info)
	}

	wr := pr.NewWorkerReader()
	defer func() { _ = wr.Close() }()
	if got, err := wr.ReadItem(1); err != nil || string(got) != string(seqBytes(4, 4)) {
		t.Errorf("ReadItem(1) = %v, %v", got, err)
	}
}

func TestNewParallelReaderErrors(t *testing.T) {
	t.Run("missing cidx", func(t *testing.T) {
		if _, err := NewParallelReader(t.TempDir()); err == nil ||
			!strings.Contains(err.Error(), "receipts.cidx not found") {
			t.Errorf("want missing-cidx error, got %v", err)
		}
	})

	t.Run("misaligned index", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "receipts.cidx"), []byte{1, 2, 3}, 0o600); err != nil {
			t.Fatal(err)
		}
		if _, err := NewParallelReader(dir); err == nil ||
			!strings.Contains(err.Error(), "invalid index size") {
			t.Errorf("want invalid-index-size error, got %v", err)
		}
	})

	t.Run("no cdat files", func(t *testing.T) {
		dir := t.TempDir()
		writeCidx(t, dir, []uint64{0})
		if _, err := NewParallelReader(dir); err == nil ||
			!strings.Contains(err.Error(), "no cdat files") {
			t.Errorf("want no-cdat-files error, got %v", err)
		}
	})
}

// TestMaxAvailableBlockPartialData pins the binary search when the index has
// more blocks than the cdat files hold data for (a mid-copy freezer).
func TestMaxAvailableBlockPartialData(t *testing.T) {
	dir := t.TempDir()
	writeCdat(t, dir, 0, seqBytes(0, 8)) // data ends at offset 8
	writeCidx(t, dir, []uint64{0, 4, 8, 12, 16})

	// A different chunk size than the multi-file test: the mapping only
	// affects files beyond index 0, so the data end stays at 8.
	pr := newSmallChunkReader(t, dir, 32)
	// Offsets 0 and 4 are < 8; block 2 starts exactly at the data end, so the
	// last block with any data is block 1.
	if got := pr.MaxAvailableBlock(); got != 1 {
		t.Errorf("MaxAvailableBlock() = %d, want 1", got)
	}
}

func TestMaxAvailableBlockNoCdatFiles(t *testing.T) {
	pr := &ParallelReader{chunkSize: 16}
	if got := pr.MaxAvailableBlock(); got != 0 {
		t.Errorf("MaxAvailableBlock() with no data = %d, want 0", got)
	}
}
