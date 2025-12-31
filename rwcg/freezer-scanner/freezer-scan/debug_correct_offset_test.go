package main

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/snappy"
)

func TestReadCorrectBlock(t *testing.T) {
	// Block 16550977: offset 12884902456 to 12884903004, length 548
	// File 6 starts at 12000000000
	// File offset = 12884902456 - 12000000000 = 884902456

	f, err := os.Open("../../mainnet/ancient/receipts.0006.cdat")
	if err != nil {
		t.Skipf("Skipping: %v", err)
	}
	defer f.Close()

	fileOffset := int64(884902456)
	length := 548

	data := make([]byte, length)
	n, err := f.ReadAt(data, fileOffset)
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if n != length {
		t.Fatalf("Short read: got %d, want %d", n, length)
	}

	t.Logf("Read %d bytes from offset %d", n, fileOffset)
	t.Logf("First 50 bytes: %x", data[:min(50, len(data))])
	t.Logf("First byte: 0x%02x", data[0])

	// Try snappy decode
	decoded, err := snappy.Decode(nil, data)
	if err != nil {
		t.Logf("Snappy decode failed: %v", err)
		// Maybe it's not compressed, try RLP directly
		t.Log("Trying direct RLP decode...")
		var receipts []rlp.RawValue
		if err := rlp.DecodeBytes(data, &receipts); err != nil {
			t.Logf("Direct RLP decode failed: %v", err)
		} else {
			t.Logf("Direct RLP: %d receipts", len(receipts))
		}
	} else {
		t.Logf("Snappy decoded: %d bytes -> %d bytes", len(data), len(decoded))
		t.Logf("Decoded first 50 bytes: %x", decoded[:min(50, len(decoded))])

		// Try RLP decode
		var receipts []rlp.RawValue
		if err := rlp.DecodeBytes(decoded, &receipts); err != nil {
			t.Logf("RLP decode failed: %v", err)
		} else {
			t.Logf("RLP decoded: %d receipts", len(receipts))
		}
	}

	// Check snappy DecodedLen
	decodedLen, err := snappy.DecodedLen(data)
	if err != nil {
		t.Logf("snappy.DecodedLen failed: %v", err)
	} else {
		t.Logf("snappy.DecodedLen: %d", decodedLen)
	}
}

func TestMultipleBlocks(t *testing.T) {
	f, err := os.Open("../../mainnet/ancient/receipts.0006.cdat")
	if err != nil {
		t.Skipf("Skipping: %v", err)
	}
	defer f.Close()

	// Test several blocks
	blocks := []struct {
		block      uint64
		fileOffset int64
		length     int
	}{
		{16550977, 884902456, 548},
		{16550978, 884903004, 853},
		{16550979, 884903857, 132},
		{16550980, 884903989, 671},
	}

	for _, b := range blocks {
		data := make([]byte, b.length)
		n, err := f.ReadAt(data, b.fileOffset)
		if err != nil || n != b.length {
			t.Errorf("Block %d: read error", b.block)
			continue
		}

		decoded, err := snappy.Decode(nil, data)
		if err != nil {
			t.Logf("Block %d: snappy failed: %v, trying direct RLP", b.block, err)
			// Try direct RLP
			var receipts []rlp.RawValue
			if err := rlp.DecodeBytes(data, &receipts); err == nil {
				t.Logf("Block %d: direct RLP OK, %d receipts", b.block, len(receipts))
			} else {
				t.Logf("Block %d: RLP also failed: %v", b.block, err)
			}
		} else {
			var receipts []rlp.RawValue
			if err := rlp.DecodeBytes(decoded, &receipts); err == nil {
				t.Logf("Block %d: snappy+RLP OK, %d receipts", b.block, len(receipts))
			} else {
				t.Logf("Block %d: snappy OK but RLP failed: %v", b.block, err)
			}
		}
	}
}

