package decode

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/snappy"
)

func TestDecodeEmptyReceipts(t *testing.T) {
	// Empty data should return empty receipts
	result, err := DecodeReceipts(nil)
	if err != nil {
		t.Errorf("DecodeReceipts(nil) error: %v", err)
	}
	if result == nil {
		t.Fatal("DecodeReceipts(nil) returned nil")
	}
	if len(result.Receipts) != 0 {
		t.Errorf("Expected 0 receipts, got %d", len(result.Receipts))
	}
}

func TestSnappyDecompression(t *testing.T) {
	// Create a simple RLP-encoded empty list
	emptyList := []byte{0xc0} // RLP for empty list

	// Compress it with snappy
	compressed := snappy.Encode(nil, emptyList)

	// Should be able to decode the compressed data
	result, err := DecodeReceipts(compressed)
	if err != nil {
		t.Errorf("DecodeReceipts(snappy) error: %v", err)
	}
	if result == nil {
		t.Fatal("DecodeReceipts(snappy) returned nil")
	}
}

func TestDecodeRLPReceipts(t *testing.T) {
	// Create a mock storage receipt
	log := &types.Log{
		Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Topics: []common.Hash{
			common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"), // Transfer event
			common.HexToHash("0x0000000000000000000000001111111111111111111111111111111111111111"),
			common.HexToHash("0x0000000000000000000000002222222222222222222222222222222222222222"),
		},
		Data: []byte{0x00, 0x00, 0x00, 0x01}, // Some data
	}

	// Create a storage receipt structure
	receipt := ReceiptForStorage{
		PostStateOrStatus: []byte{1}, // Success status
		CumulativeGasUsed: 21000,
		Bloom:             types.Bloom{},
		Logs:              []*types.Log{log},
	}

	// RLP encode as a list of receipts
	receipts := []ReceiptForStorage{receipt}
	encoded, err := rlp.EncodeToBytes(receipts)
	if err != nil {
		t.Fatalf("Failed to encode receipts: %v", err)
	}

	// Decode
	result, err := DecodeReceipts(encoded)
	if err != nil {
		t.Fatalf("DecodeReceipts error: %v", err)
	}

	if len(result.Receipts) != 1 {
		t.Fatalf("Expected 1 receipt, got %d", len(result.Receipts))
	}

	if len(result.Receipts[0].Logs) != 1 {
		t.Fatalf("Expected 1 log, got %d", len(result.Receipts[0].Logs))
	}

	decodedLog := result.Receipts[0].Logs[0]
	if decodedLog.Address != log.Address {
		t.Errorf("Log address mismatch: got %s, want %s", decodedLog.Address, log.Address)
	}

	if len(decodedLog.Topics) != 3 {
		t.Errorf("Expected 3 topics, got %d", len(decodedLog.Topics))
	}
}

func TestFilterLogs(t *testing.T) {
	logs := []*DecodedLog{
		{
			Address: common.HexToAddress("0x1111111111111111111111111111111111111111"),
			Topics:  []common.Hash{common.HexToHash("0xaaaa")},
		},
		{
			Address: common.HexToAddress("0x2222222222222222222222222222222222222222"),
			Topics:  []common.Hash{common.HexToHash("0xbbbb")},
		},
		{
			Address: common.HexToAddress("0x1111111111111111111111111111111111111111"),
			Topics:  []common.Hash{common.HexToHash("0xbbbb")},
		},
	}

	// Filter by contract
	contracts := map[common.Address]bool{
		common.HexToAddress("0x1111111111111111111111111111111111111111"): true,
	}
	filtered := FilterLogs(logs, contracts, nil)
	if len(filtered) != 2 {
		t.Errorf("Contract filter: expected 2 logs, got %d", len(filtered))
	}

	// Filter by event signature
	eventSigs := map[common.Hash]bool{
		common.HexToHash("0xbbbb"): true,
	}
	filtered = FilterLogs(logs, nil, eventSigs)
	if len(filtered) != 2 {
		t.Errorf("Event filter: expected 2 logs, got %d", len(filtered))
	}

	// Filter by both
	filtered = FilterLogs(logs, contracts, eventSigs)
	if len(filtered) != 1 {
		t.Errorf("Combined filter: expected 1 log, got %d", len(filtered))
	}
}

func TestComputeLogIdentityHash(t *testing.T) {
	log1 := &DecodedLog{
		Address:    common.HexToAddress("0x1111111111111111111111111111111111111111"),
		Topics:     []common.Hash{common.HexToHash("0xaaaa")},
		DataKeccak: common.HexToHash("0xcccc"),
	}

	log2 := &DecodedLog{
		Address:    common.HexToAddress("0x1111111111111111111111111111111111111111"),
		Topics:     []common.Hash{common.HexToHash("0xaaaa")},
		DataKeccak: common.HexToHash("0xcccc"),
	}

	log3 := &DecodedLog{
		Address:    common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Topics:     []common.Hash{common.HexToHash("0xaaaa")},
		DataKeccak: common.HexToHash("0xcccc"),
	}

	// Same log at same position should have same hash
	hash1 := ComputeLogIdentityHash(100, 0, 0, log1)
	hash2 := ComputeLogIdentityHash(100, 0, 0, log2)
	if hash1 != hash2 {
		t.Error("Identical logs should have same identity hash")
	}

	// Different address should have different hash
	hash3 := ComputeLogIdentityHash(100, 0, 0, log3)
	if hash1 == hash3 {
		t.Error("Logs with different addresses should have different identity hash")
	}

	// Different block should have different hash
	hash4 := ComputeLogIdentityHash(101, 0, 0, log1)
	if hash1 == hash4 {
		t.Error("Logs at different blocks should have different identity hash")
	}
}

func TestIsValidRLPPrefix(t *testing.T) {
	tests := []struct {
		b        byte
		expected bool
	}{
		{0x00, false},
		{0x7f, false},
		{0x80, false},
		{0xbf, false},
		{0xc0, true}, // Empty list
		{0xc1, true}, // Short list
		{0xf7, true}, // Max short list
		{0xf8, true}, // Long list
		{0xff, true}, // Long list
	}

	for _, tt := range tests {
		got := isValidRLPPrefix(tt.b)
		if got != tt.expected {
			t.Errorf("isValidRLPPrefix(0x%02x) = %v, want %v", tt.b, got, tt.expected)
		}
	}
}


