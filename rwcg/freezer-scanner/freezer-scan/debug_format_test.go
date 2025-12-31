package main

import (
	"encoding/binary"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/snappy"
)

func TestDebugFormat(t *testing.T) {
	// Read from start of file 6
	f, err := os.Open("../../mainnet/ancient/receipts.0006.cdat")
	if err != nil {
		t.Skipf("Skipping: %v", err)
	}
	defer f.Close()

	// Read first 1000 bytes
	data := make([]byte, 1000)
	n, _ := f.Read(data)
	data = data[:n]

	t.Logf("First 50 bytes: %x", data[:50])
	t.Logf("First byte: 0x%02x (dec: %d)", data[0], data[0])

	// Check various interpretations:
	// 1. Is first byte a snappy varint length?
	// Snappy varint: if < 128, it's the length directly
	// 0x93 = 147 > 128, so it uses multi-byte varint
	if data[0] >= 0x80 {
		// Multi-byte varint
		len1 := int(data[0]&0x7f) | int(data[1])<<7
		t.Logf("Snappy varint length (2 bytes): %d", len1)
	}

	// 2. Check if there's a 4-byte length prefix
	len32 := binary.BigEndian.Uint32(data[:4])
	t.Logf("Big-endian 32-bit prefix: %d (0x%08x)", len32, len32)
	len32le := binary.LittleEndian.Uint32(data[:4])
	t.Logf("Little-endian 32-bit prefix: %d (0x%08x)", len32le, len32le)

	// 3. Try snappy decode on data starting at various offsets
	for offset := 0; offset < 10; offset++ {
		decoded, err := snappy.Decode(nil, data[offset:])
		if err == nil {
			t.Logf("Snappy decode succeeded at offset %d, decoded %d bytes", offset, len(decoded))
			t.Logf("Decoded first 30 bytes: %x", decoded[:min(30, len(decoded))])
			// Try RLP decode
			var raw []rlp.RawValue
			if err := rlp.DecodeBytes(decoded, &raw); err == nil {
				t.Logf("RLP decoded as list with %d items", len(raw))
			} else {
				t.Logf("RLP decode failed: %v", err)
			}
			break
		}
	}

	// 4. Try to interpret as RLP directly
	var raw []rlp.RawValue
	if err := rlp.DecodeBytes(data, &raw); err != nil {
		t.Logf("Direct RLP decode failed: %v", err)
	} else {
		t.Logf("Direct RLP: %d items", len(raw))
	}

	// 5. Check if first byte indicates a typed receipt (0x01, 0x02, etc)
	// Or if it's EIP-2718 encoding
	t.Logf("If interpreted as typed tx, type would be: %d", data[0])

	// 6. Look for RLP list markers in the data
	for i := 0; i < min(100, len(data)); i++ {
		if data[i] >= 0xc0 && data[i] <= 0xf7 {
			t.Logf("Possible RLP short list at offset %d: 0x%02x", i, data[i])
		}
		if data[i] >= 0xf8 && data[i] <= 0xff {
			t.Logf("Possible RLP long list at offset %d: 0x%02x", i, data[i])
		}
	}
}

func TestDebugSnappyMagic(t *testing.T) {
	// Check for snappy stream/framing format magic bytes
	f, err := os.Open("../../mainnet/ancient/receipts.0006.cdat")
	if err != nil {
		t.Skipf("Skipping: %v", err)
	}
	defer f.Close()

	header := make([]byte, 10)
	f.Read(header)

	// Snappy framing format starts with stream identifier: 0xff 0x06 0x00 0x00 's' 'N' 'a' 'P' 'p' 'Y'
	if header[0] == 0xff {
		t.Log("Possible snappy framing format")
	}

	t.Logf("Header bytes: %x", header)
}

