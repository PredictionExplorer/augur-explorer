package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/snappy"
)

func readOffset(path string, block uint64) (uint64, error) {
	f, _ := os.Open(path)
	defer f.Close()
	buf := make([]byte, 6)
	f.ReadAt(buf, int64(block*6))
	return uint64(buf[0])<<40 | uint64(buf[1])<<32 | uint64(buf[2])<<24 |
		uint64(buf[3])<<16 | uint64(buf[4])<<8 | uint64(buf[5]), nil
}

func TestDebugBlock(t *testing.T) {
	cidx := "../../mainnet/receipts.cidx"
	block := uint64(16550977)

	offset, _ := readOffset(cidx, block)
	nextOffset, _ := readOffset(cidx, block+1)
	length := nextOffset - offset

	fileIdx := 6
	fileStart := uint64(fileIdx) * 2000000000
	fileOffset := offset - fileStart

	t.Logf("Block %d: global offset %d, file offset %d, length %d", block, offset, fileOffset, length)

	f, err := os.Open(fmt.Sprintf("../../mainnet/ancient/receipts.%04d.cdat", fileIdx))
	if err != nil {
		t.Skipf("Skipping: %v", err)
	}
	defer f.Close()

	// Check file size
	info, _ := f.Stat()
	t.Logf("File size: %d, requested offset: %d", info.Size(), fileOffset)
	
	if fileOffset >= uint64(info.Size()) {
		t.Logf("PROBLEM: File offset %d is beyond file size %d!", fileOffset, info.Size())
		t.Logf("This means the data at this offset doesn't exist in the sparse file")
		return
	}

	data := make([]byte, length)
	n, _ := f.ReadAt(data, int64(fileOffset))
	t.Logf("Read %d bytes", n)

	if n > 0 {
		t.Logf("Raw data first 20 bytes: %x", data[:min(20, n)])
		t.Logf("First byte: 0x%02x", data[0])
		
		// Try snappy decompress
		decoded, err := snappy.Decode(nil, data[:n])
		if err != nil {
			t.Logf("Snappy decode error: %v", err)
		} else {
			t.Logf("Snappy decoded OK, %d bytes -> %d bytes", n, len(decoded))
		}
		
		// Try parsing as Arbitrum format
		if n > 4 {
			size, vlen := decodeVarint(data)
			t.Logf("Varint at start: size=%d, bytes=%d", size, vlen)
			
			// Try RLP from different offsets
			for off := 0; off <= 4; off++ {
				if off < n && data[off] >= 0xc0 {
					var rawReceipts []rlp.RawValue
					err := rlp.DecodeBytes(data[off:n], &rawReceipts)
					if err == nil {
						t.Logf("RLP decode from offset %d SUCCESS! %d receipts", off, len(rawReceipts))
						break
					}
				}
			}
		}
	}
}

func TestDebugFileStart(t *testing.T) {
	f, err := os.Open("../../mainnet/ancient/receipts.0006.cdat")
	if err != nil {
		t.Skipf("Skipping: %v", err)
	}
	defer f.Close()
	
	// Read first 100000 bytes of the file
	data := make([]byte, 100000)
	n, _ := f.Read(data)
	data = data[:n]
	
	t.Logf("File 6 first 50 bytes: %x", data[:50])
	
	// Examine boundary between block 0 and block 1
	// Block 0: pos 0, RLP size 4627
	// RLP starts at pos 4 (after 2-byte varint + 2-byte header)
	// RLP ends at pos 4 + 4627 = 4631
	
	t.Logf("\n--- Boundary analysis ---")
	t.Logf("Data at pos 4627-4650: %x", data[4627:min(4650, len(data))])
	
	// Look at where block 0 actually ends
	// Block 0 RLP: f9 12 10 ... (at pos 4)
	// RLP list header: f9 = 2-byte length, 12 10 = 4624 bytes
	// Total RLP = 3 + 4624 = 4627 bytes
	// RLP ends at 4 + 4627 = 4631
	
	// So data at 4631 should be start of block 1
	t.Logf("Data at pos 4631: %x", data[4631:4645])
	
	// Block 1 structure:
	// - If format is same: varint + 2-byte header + RLP
	// At 4631: 82 7c = varint 15874
	// At 4633: 00 4e = header?
	// At 4635: f0 04 86... - 0xf0 is RLP list with payload 48 bytes... doesn't match 15874!
	
	// Let me try different header sizes
	for headerSkip := 0; headerSkip <= 4; headerSkip++ {
		pos := 4631
		size, vlen := decodeVarint(data[pos:])
		rlpStart := pos + vlen + headerSkip
		
		if rlpStart < len(data) && data[rlpStart] >= 0xc0 {
			t.Logf("With header=%d: varint=%d at pos %d, RLP byte at %d is 0x%02x", 
				headerSkip, size, pos, rlpStart, data[rlpStart])
			
			// Try to decode
			if int(size) < len(data)-rlpStart {
				rlpData := data[rlpStart : rlpStart+int(size)]
				var rawReceipts []rlp.RawValue
				err := rlp.DecodeBytes(rlpData, &rawReceipts)
				if err == nil {
					t.Logf("  SUCCESS! %d receipts", len(rawReceipts))
				} else {
					t.Logf("  Failed: %v", err)
				}
			}
		}
	}
	
	// Try snappy on block 1
	t.Logf("\n--- Trying snappy on block 1 ---")
	pos := 4631
	size, vlen := decodeVarint(data[pos:])
	t.Logf("Block 1 varint: %d (in %d bytes)", size, vlen)
	
	// Maybe the varint is the UNCOMPRESSED size, and the data is snappy compressed
	// Let's try snappy decode from different starting positions
	for hdrSkip := 0; hdrSkip <= 4; hdrSkip++ {
		compStart := pos + vlen + hdrSkip
		// Try snappy decode with various chunk sizes
		for chunkSize := 100; chunkSize <= 10000; chunkSize += 500 {
			if compStart+chunkSize > len(data) {
				break
			}
			chunk := data[compStart : compStart+chunkSize]
			decoded, err := snappy.Decode(nil, chunk)
			if err == nil && len(decoded) > 0 {
				t.Logf("Snappy SUCCESS with hdr=%d, chunk=%d: %d -> %d bytes", 
					hdrSkip, chunkSize, len(chunk), len(decoded))
				t.Logf("Decoded first 20 bytes: %x", decoded[:min(20, len(decoded))])
				break
			}
		}
	}
	
	// Alternative theory: Block 0 is special (uncompressed), other blocks are compressed
	// Let's check the DecodedLen for block 1
	compData := data[pos+vlen : min(pos+vlen+20000, len(data))]
	decodedLen, err := snappy.DecodedLen(compData)
	t.Logf("Snappy DecodedLen from pos %d: %d, err: %v", pos+vlen, decodedLen, err)
	
	// If DecodedLen matches varint size, data is snappy compressed
	if decodedLen == int(size) {
		t.Logf("DecodedLen matches varint! Data is likely snappy compressed")
		decoded, err := snappy.Decode(nil, compData[:decodedLen])
		if err != nil {
			t.Logf("But snappy decode failed: %v", err)
		} else {
			t.Logf("Snappy decoded: %d bytes, first 30: %x", len(decoded), decoded[:min(30, len(decoded))])
		}
	}
}

// decodeVarint decodes a varint from data, returns value and bytes consumed
func decodeVarint(data []byte) (uint64, int) {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

