// Fuzz targets for freezer index/data reading (MODERNIZATION.md §4.4).
// Corrupt .cidx/.cdat bytes must produce errors, never panics, and reads must
// stay inside the data files.
package freezerscanner

import (
	"os"
	"path/filepath"
	"testing"
)

// parse6 mirrors the 48-bit big-endian offset composition used by
// FreezerReader.ReadOffset and ParallelReader.getOffset.
func parse6(b []byte) uint64 {
	return uint64(b[0])<<40 | uint64(b[1])<<32 | uint64(b[2])<<24 |
		uint64(b[3])<<16 | uint64(b[4])<<8 | uint64(b[5])
}

func FuzzUint48RoundTrip(f *testing.F) {
	for _, seed := range []uint64{0, 1, 255, 256, 0x123456789ABC, 0xFFFFFFFFFFFF, 1 << 47} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint64) {
		b := Uint48ToBytes(v)
		if len(b) != 6 {
			t.Fatalf("Uint48ToBytes(%d) returned %d bytes, want 6", v, len(b))
		}
		got := parse6(b)
		want := v & 0xFFFFFFFFFFFF // values above 48 bits truncate by design
		if got != want {
			t.Fatalf("round-trip mismatch: Uint48ToBytes(%d) parsed back as %d, want %d", v, got, want)
		}
	})
}

func FuzzFreezerIndexRead(f *testing.F) {
	// Monotonic 3-entry index over a 30-byte data file.
	goodIdx := append(append(Uint48ToBytes(0), Uint48ToBytes(5)...), Uint48ToBytes(15)...)
	f.Add(goodIdx, make([]byte, 30), uint64(1))
	// Offset regression (10 -> 5).
	f.Add(append(append(Uint48ToBytes(0), Uint48ToBytes(10)...), Uint48ToBytes(5)...), make([]byte, 30), uint64(0))
	// Offsets pointing far beyond the data file.
	f.Add(append(Uint48ToBytes(0), Uint48ToBytes(0xFFFFFFFFFFFF)...), []byte{1, 2, 3}, uint64(0))
	// Truncated index (not a multiple of 6) and empty everything.
	f.Add([]byte{1, 2, 3}, []byte(nil), uint64(0))
	f.Add([]byte(nil), []byte(nil), uint64(0))
	f.Fuzz(func(t *testing.T, cidx []byte, cdat []byte, blockNum uint64) {
		// Keep per-iteration file IO bounded.
		if len(cidx) > 4096 || len(cdat) > 65536 {
			t.Skip("oversized input")
		}
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "receipts.cidx"), cidx, 0o600); err != nil {
			t.Fatalf("write cidx: %v", err)
		}
		if err := os.WriteFile(filepath.Join(dir, "receipts.0000.cdat"), cdat, 0o600); err != nil {
			t.Fatalf("write cdat: %v", err)
		}

		fr, err := NewFreezerReader(dir)
		if err != nil {
			return // malformed layout rejected cleanly
		}
		defer func() { _ = fr.Close() }()

		if fr.ItemCount() != uint64(len(cidx))/IndexEntrySize {
			t.Fatalf("ItemCount() = %d for %d index bytes", fr.ItemCount(), len(cidx))
		}

		// Offsets read back must match the raw index bytes.
		if blockNum < fr.ItemCount() {
			off, err := fr.ReadOffset(blockNum)
			if err != nil {
				t.Fatalf("ReadOffset(%d) failed on a well-formed index: %v", blockNum, err)
			}
			if want := parse6(cidx[blockNum*IndexEntrySize:]); off != want {
				t.Fatalf("ReadOffset(%d) = %d, want %d", blockNum, off, want)
			}
		} else if _, err := fr.ReadOffset(blockNum); err == nil {
			t.Fatalf("ReadOffset(%d) succeeded beyond item count %d", blockNum, fr.ItemCount())
		}

		// ReadItem must never panic and never hand back more bytes than exist.
		data, err := fr.ReadItem(blockNum)
		if err == nil && len(data) > len(cdat) {
			t.Fatalf("ReadItem(%d) returned %d bytes from a %d-byte data file", blockNum, len(data), len(cdat))
		}

		// Validation walks are panic-free on arbitrary indexes.
		_ = fr.ValidateIndexRange(0, fr.ItemCount())
	})
}
