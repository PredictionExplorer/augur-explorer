// Package decode provides decoding utilities for geth freezer receipts
// It handles snappy decompression and RLP decoding of receipt data.
package decode

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/snappy"
)

// DecodedLog represents a single log entry with all relevant fields
type DecodedLog struct {
	ReceiptIndex uint // Index of receipt within the block
	LogIndex     uint // Index of log within the receipt
	Address      common.Address
	Topics       []common.Hash
	Data         []byte
	DataKeccak   common.Hash
}

// DecodedReceipt represents a decoded receipt with its logs
type DecodedReceipt struct {
	Index             uint
	Status            uint64
	CumulativeGasUsed uint64
	Logs              []*DecodedLog
}

// BlockReceipts represents all receipts for a block
type BlockReceipts struct {
	Receipts []*DecodedReceipt
	AllLogs  []*DecodedLog // Flattened list of all logs
}

// DecodeReceipts decodes a raw freezer receipts blob into structured receipts
// It handles snappy decompression and RLP decoding
func DecodeReceipts(data []byte) (*BlockReceipts, error) {
	if len(data) == 0 {
		return &BlockReceipts{}, nil
	}

	// Try snappy decompression first
	decompressed, err := trySnappyDecompress(data)
	if err != nil {
		return nil, fmt.Errorf("decompression failed: %w", err)
	}

	// Decode RLP into receipts
	receipts, err := decodeRLPReceipts(decompressed)
	if err != nil {
		return nil, fmt.Errorf("RLP decode failed: %w", err)
	}

	return receipts, nil
}

// trySnappyDecompress attempts to decompress data using snappy.
// If the data is not snappy-compressed, it handles the Arbitrum Nitro format
// which has a header before raw RLP data.
func trySnappyDecompress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return data, nil
	}

	// If data already starts with valid RLP, return as-is
	if isValidRLPPrefix(data[0]) {
		return data, nil
	}

	// First, try standard snappy decode
	decoded, err := snappy.Decode(nil, data)
	if err == nil {
		return decoded, nil
	}

	// Snappy decode failed. Check for Arbitrum Nitro format:
	// - 2 bytes: varint for RLP size
	// - 2 bytes: unknown header (possibly flags/CRC)
	// - N bytes: raw RLP data
	
	if len(data) < 5 {
		return data, nil
	}

	// Parse the varint at the start
	rlpSize, varintBytes := decodeVarint(data)
	if varintBytes == 0 || rlpSize == 0 {
		return data, nil
	}

	// Find where valid RLP starts (scan for RLP list prefix after varint)
	for offset := varintBytes; offset < len(data) && offset < varintBytes+10; offset++ {
		if isValidRLPPrefix(data[offset]) {
			// Found RLP start, check if the size makes sense
			remainingData := len(data) - offset
			if uint64(remainingData) >= rlpSize {
				// Size matches, extract exact amount
				return data[offset : offset+int(rlpSize)], nil
			}
			// Size doesn't match exactly, return from this offset
			return data[offset:], nil
		}
	}

	// No valid RLP found after header, return original
	return data, nil
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

// isValidRLPPrefix checks if a byte could be a valid RLP prefix
func isValidRLPPrefix(b byte) bool {
	// RLP encoding rules:
	// - 0x00-0x7f: single byte
	// - 0x80-0xb7: short string (0-55 bytes)
	// - 0xb8-0xbf: long string
	// - 0xc0-0xf7: short list (0-55 bytes total)
	// - 0xf8-0xff: long list
	// For receipts, we expect a list, so 0xc0 and above
	return b >= 0xc0
}

// ReceiptForStorage is a minimal struct for RLP decoding storage receipts
// This matches geth's internal storage format
type ReceiptForStorage struct {
	PostStateOrStatus []byte
	CumulativeGasUsed uint64
	Bloom             types.Bloom
	Logs              []*types.Log
}

// decodeRLPReceipts decodes RLP-encoded receipts
func decodeRLPReceipts(data []byte) (*BlockReceipts, error) {
	// First, try to decode as a list of raw values
	var rawReceipts []rlp.RawValue
	if err := rlp.DecodeBytes(data, &rawReceipts); err != nil {
		return nil, fmt.Errorf("failed to decode receipt list: %w", err)
	}

	result := &BlockReceipts{
		Receipts: make([]*DecodedReceipt, 0, len(rawReceipts)),
		AllLogs:  make([]*DecodedLog, 0),
	}

	for i, rawReceipt := range rawReceipts {
		receipt, err := decodeStorageReceipt(rawReceipt, uint(i))
		if err != nil {
			return nil, fmt.Errorf("failed to decode receipt %d: %w", i, err)
		}
		result.Receipts = append(result.Receipts, receipt)
		result.AllLogs = append(result.AllLogs, receipt.Logs...)
	}

	return result, nil
}

// decodeStorageReceipt decodes a single receipt from storage format
func decodeStorageReceipt(data []byte, index uint) (*DecodedReceipt, error) {
	// Handle typed receipts (EIP-2718): first byte is tx type if < 0x80
	payload := data
	if len(data) > 0 && data[0] < 0x80 {
		// Typed receipt - skip the type byte
		payload = data[1:]
	}

	// Decode as storage receipt
	var storedReceipt ReceiptForStorage
	if err := rlp.DecodeBytes(payload, &storedReceipt); err != nil {
		// Try alternative decoding - some formats differ
		return decodeReceiptAlternative(payload, index)
	}

	receipt := &DecodedReceipt{
		Index:             index,
		CumulativeGasUsed: storedReceipt.CumulativeGasUsed,
		Logs:              make([]*DecodedLog, 0, len(storedReceipt.Logs)),
	}

	// Parse status from PostStateOrStatus
	if len(storedReceipt.PostStateOrStatus) == 1 {
		receipt.Status = uint64(storedReceipt.PostStateOrStatus[0])
	} else if len(storedReceipt.PostStateOrStatus) > 1 {
		// Post-state root (pre-Byzantium)
		receipt.Status = 1 // Assume success for pre-Byzantium
	}

	// Convert logs
	for j, log := range storedReceipt.Logs {
		decodedLog := &DecodedLog{
			ReceiptIndex: index,
			LogIndex:     uint(j),
			Address:      log.Address,
			Topics:       log.Topics,
			Data:         log.Data,
		}
		if len(log.Data) > 0 {
			decodedLog.DataKeccak = crypto.Keccak256Hash(log.Data)
		}
		receipt.Logs = append(receipt.Logs, decodedLog)
	}

	return receipt, nil
}

// StoredLog is a minimal log structure for alternative decoding
type StoredLog struct {
	Address common.Address
	Topics  []common.Hash
	Data    []byte
}

// decodeReceiptAlternative tries an alternative decoding format
func decodeReceiptAlternative(data []byte, index uint) (*DecodedReceipt, error) {
	// Try decoding as a simple struct without bloom
	type SimpleReceipt struct {
		Status            uint64
		CumulativeGasUsed uint64
		Logs              []*StoredLog
	}

	var simple SimpleReceipt
	if err := rlp.DecodeBytes(data, &simple); err != nil {
		// Last resort: try to extract just the logs
		return decodeReceiptLogsOnly(data, index)
	}

	receipt := &DecodedReceipt{
		Index:             index,
		Status:            simple.Status,
		CumulativeGasUsed: simple.CumulativeGasUsed,
		Logs:              make([]*DecodedLog, 0, len(simple.Logs)),
	}

	for j, log := range simple.Logs {
		decodedLog := &DecodedLog{
			ReceiptIndex: index,
			LogIndex:     uint(j),
			Address:      log.Address,
			Topics:       log.Topics,
			Data:         log.Data,
		}
		if len(log.Data) > 0 {
			decodedLog.DataKeccak = crypto.Keccak256Hash(log.Data)
		}
		receipt.Logs = append(receipt.Logs, decodedLog)
	}

	return receipt, nil
}

// decodeReceiptLogsOnly extracts logs using stream decoding when struct decode fails
func decodeReceiptLogsOnly(data []byte, index uint) (*DecodedReceipt, error) {
	receipt := &DecodedReceipt{
		Index: index,
		Logs:  make([]*DecodedLog, 0),
	}

	// Use stream decoder to navigate the receipt
	stream := rlp.NewStream(bytes.NewReader(data), 0)

	// Try to read as list
	if _, err := stream.List(); err != nil {
		return nil, fmt.Errorf("expected list: %w", err)
	}

	// Skip first few fields (status, cumulativeGas, bloom) until we find logs
	fieldsToSkip := 0
	for fieldsToSkip < 3 {
		kind, _, err := stream.Kind()
		if err != nil {
			break
		}
		if kind == rlp.List {
			// This might be the logs array
			break
		}
		if err := skipStreamValue(stream); err != nil {
			break
		}
		fieldsToSkip++
	}

	// Try to decode the logs list
	if _, err := stream.List(); err != nil {
		// No logs found
		return receipt, nil
	}

	logIndex := uint(0)
	for {
		// Try to decode each log
		var log StoredLog
		if err := stream.Decode(&log); err != nil {
			break
		}

		decodedLog := &DecodedLog{
			ReceiptIndex: index,
			LogIndex:     logIndex,
			Address:      log.Address,
			Topics:       log.Topics,
			Data:         log.Data,
		}
		if len(log.Data) > 0 {
			decodedLog.DataKeccak = crypto.Keccak256Hash(log.Data)
		}
		receipt.Logs = append(receipt.Logs, decodedLog)
		logIndex++
	}

	return receipt, nil
}

// skipStreamValue skips a single value in the RLP stream
func skipStreamValue(stream *rlp.Stream) error {
	kind, size, err := stream.Kind()
	if err != nil {
		return err
	}

	switch kind {
	case rlp.Byte, rlp.String:
		_, err = stream.Bytes()
	case rlp.List:
		if _, err = stream.List(); err != nil {
			return err
		}
		for {
			if err := skipStreamValue(stream); err != nil {
				break
			}
		}
		err = stream.ListEnd()
	}
	_ = size
	return err
}

// FilterLogs filters logs by contract addresses and event signatures
func FilterLogs(logs []*DecodedLog, contracts map[common.Address]bool, eventSigs map[common.Hash]bool) []*DecodedLog {
	var result []*DecodedLog

	for _, log := range logs {
		// Check contract address if filter is set
		if len(contracts) > 0 && !contracts[log.Address] {
			continue
		}

		// Check event signature (topic0) if filter is set
		if len(eventSigs) > 0 {
			if len(log.Topics) == 0 {
				continue
			}
			if !eventSigs[log.Topics[0]] {
				continue
			}
		}

		result = append(result, log)
	}

	return result
}

// ComputeLogIdentityHash computes a unique hash for a log entry
// This can be used for deduplication and comparison
func ComputeLogIdentityHash(blockNum uint64, txIndex, logIndex uint, log *DecodedLog) common.Hash {
	var buf bytes.Buffer

	// Write block number (8 bytes)
	b := make([]byte, 8)
	b[0] = byte(blockNum >> 56)
	b[1] = byte(blockNum >> 48)
	b[2] = byte(blockNum >> 40)
	b[3] = byte(blockNum >> 32)
	b[4] = byte(blockNum >> 24)
	b[5] = byte(blockNum >> 16)
	b[6] = byte(blockNum >> 8)
	b[7] = byte(blockNum)
	buf.Write(b)

	// Write tx index (4 bytes)
	buf.WriteByte(byte(txIndex >> 24))
	buf.WriteByte(byte(txIndex >> 16))
	buf.WriteByte(byte(txIndex >> 8))
	buf.WriteByte(byte(txIndex))

	// Write log index (4 bytes)
	buf.WriteByte(byte(logIndex >> 24))
	buf.WriteByte(byte(logIndex >> 16))
	buf.WriteByte(byte(logIndex >> 8))
	buf.WriteByte(byte(logIndex))

	// Write address
	buf.Write(log.Address.Bytes())

	// Write topics
	for _, topic := range log.Topics {
		buf.Write(topic.Bytes())
	}

	// Write data keccak
	buf.Write(log.DataKeccak.Bytes())

	return crypto.Keccak256Hash(buf.Bytes())
}

