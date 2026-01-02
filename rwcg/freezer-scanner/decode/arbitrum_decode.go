// Package decode provides decoders for Arbitrum Nitro freezer receipts
package decode

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/snappy"
)

// Log represents an Ethereum log entry
type Log struct {
	Address common.Address
	Topics  []common.Hash
	Data    []byte
}

// DecodeArbitrumLog decodes a single log from raw RLP
func DecodeArbitrumLog(rawLog []byte) (*Log, error) {
	var fields []rlp.RawValue
	if err := rlp.DecodeBytes(rawLog, &fields); err != nil {
		return nil, err
	}
	if len(fields) < 3 {
		return nil, fmt.Errorf("expected 3 fields, got %d", len(fields))
	}

	log := &Log{}
	if err := rlp.DecodeBytes(fields[0], &log.Address); err != nil {
		return nil, fmt.Errorf("decode address: %w", err)
	}
	if err := rlp.DecodeBytes(fields[1], &log.Topics); err != nil {
		return nil, fmt.Errorf("decode topics: %w", err)
	}
	if err := rlp.DecodeBytes(fields[2], &log.Data); err != nil {
		return nil, fmt.Errorf("decode data: %w", err)
	}
	return log, nil
}

// smartDecompress handles multiple data formats:
// 1. Raw RLP (starts with 0xc0+)
// 2. Snappy compressed
// 3. Varint header + raw RLP (Arbitrum Nitro uncompressed format)
func smartDecompress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return data, nil
	}

	// If data already starts with valid RLP list prefix, return as-is
	if data[0] >= 0xc0 {
		return data, nil
	}

	// Try standard snappy decode first
	decoded, err := snappy.Decode(nil, data)
	if err == nil {
		return decoded, nil
	}

	// Snappy failed - try Arbitrum Nitro header format:
	// varint(size) + optional header bytes + raw RLP
	if len(data) < 5 {
		return nil, fmt.Errorf("data too short for header format: %d bytes", len(data))
	}

	// Parse varint at start
	rlpSize, varintBytes := decodeVarintLocal(data)
	if varintBytes == 0 || rlpSize == 0 {
		return nil, fmt.Errorf("invalid varint header")
	}

	// Scan for RLP list prefix after varint
	for offset := varintBytes; offset < len(data) && offset < varintBytes+10; offset++ {
		if data[offset] >= 0xc0 {
			// Found RLP start
			remaining := len(data) - offset
			if uint64(remaining) >= rlpSize {
				return data[offset : offset+int(rlpSize)], nil
			}
			return data[offset:], nil
		}
	}

	return nil, fmt.Errorf("no valid RLP found after header")
}

// decodeVarintLocal decodes a varint, returns value and bytes consumed
func decodeVarintLocal(data []byte) (uint64, int) {
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

// DecodeArbitrumReceipts decodes Arbitrum Nitro receipts (snappy or raw)
// and returns all logs from all receipts in the block
func DecodeArbitrumReceipts(compressedData []byte) ([]*Log, error) {
	// Smart decompress - handles snappy, raw RLP, and header formats
	data, err := smartDecompress(compressedData)
	if err != nil {
		return nil, fmt.Errorf("decompress: %w", err)
	}

	// Decode as list of receipts
	var rawReceipts []rlp.RawValue
	if err := rlp.DecodeBytes(data, &rawReceipts); err != nil {
		return nil, fmt.Errorf("decode receipts list: %w", err)
	}

	var allLogs []*Log
	for _, rawReceipt := range rawReceipts {
		// Handle typed receipts (EIP-2718): first byte < 0x80 is tx type
		payload := []byte(rawReceipt)
		if len(payload) > 0 && payload[0] < 0x80 {
			payload = payload[1:] // Skip type byte
		}

		// Try to decode receipt fields
		var fields []rlp.RawValue
		if err := rlp.DecodeBytes(payload, &fields); err != nil {
			continue
		}

		// Find logs - they're typically in field 3 or 6 depending on format
		// Try field 3 first (standard format: status, gasUsed, bloom, logs)
		logsFieldIdx := 3
		if len(fields) >= 7 {
			// Arbitrum extended format with 7 fields - logs in field 6
			logsFieldIdx = 6
		}
		if logsFieldIdx >= len(fields) {
			continue
		}

		var rawLogs []rlp.RawValue
		if err := rlp.DecodeBytes(fields[logsFieldIdx], &rawLogs); err != nil {
			// Try other positions
			for idx := 3; idx < len(fields); idx++ {
				if err := rlp.DecodeBytes(fields[idx], &rawLogs); err == nil && len(rawLogs) > 0 {
					break
				}
			}
			if len(rawLogs) == 0 {
				continue
			}
		}

		for _, rawLog := range rawLogs {
			log, err := DecodeArbitrumLog(rawLog)
			if err == nil {
				allLogs = append(allLogs, log)
			}
		}
	}

	return allLogs, nil
}


