package toolutil

import (
	"bytes"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func TestBackupReceiptRoundTrip(t *testing.T) {
	t.Parallel()
	address := ethcommon.HexToAddress("0x2100000000000000000000000000000000000021")
	topic := ethcommon.HexToHash("0x1234")
	receipt := &types.Receipt{
		Type:              types.DynamicFeeTxType,
		Status:            types.ReceiptStatusSuccessful,
		CumulativeGasUsed: 12345,
		Logs: []*types.Log{
			nil,
			{Address: address, Topics: []ethcommon.Hash{topic}, Data: []byte{1, 2}, Index: 7},
		},
	}
	backup := ReceiptToBackup(receipt)
	if backup.Version != backupReceiptVersion || len(backup.Logs) != 1 ||
		backup.Logs[0].Index != 7 {
		t.Fatalf("backup = %+v", backup)
	}
	encoded, err := EncodeBackupReceiptRLP(receipt)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := DecodeBackupReceiptRLP(encoded)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Version != backupReceiptVersion ||
		decoded.Type != receipt.Type ||
		decoded.CumulativeGasUsed != receipt.CumulativeGasUsed ||
		len(decoded.Logs) != 1 {
		t.Fatalf("decoded = %+v", decoded)
	}
	log := decoded.Logs[0].ToTypesLog()
	if log.Address != address || log.Index != 7 ||
		!bytes.Equal(log.Data, []byte{1, 2}) {
		t.Fatalf("types log = %+v", log)
	}

	auto, legacy, err := TryDecodeReceiptRLP(encoded)
	if err != nil || legacy || auto.Version != backupReceiptVersion {
		t.Fatalf("TryDecode backup = %+v legacy=%v err=%v", auto, legacy, err)
	}
}

func TestTryDecodeLegacyReceiptRLP(t *testing.T) {
	t.Parallel()
	legacyReceipt := &types.Receipt{
		Status:            types.ReceiptStatusFailed,
		CumulativeGasUsed: 99,
		Logs: []*types.Log{{
			Address: ethcommon.HexToAddress("0x2200000000000000000000000000000000000022"),
			Index:   8,
		}},
	}
	encoded, err := rlp.EncodeToBytes(legacyReceipt)
	if err != nil {
		t.Fatal(err)
	}
	decoded, legacy, err := TryDecodeReceiptRLP(encoded)
	if err != nil || !legacy || decoded.Status != types.ReceiptStatusFailed ||
		len(decoded.Logs) != 1 {
		t.Fatalf("decoded = %+v legacy=%v err=%v", decoded, legacy, err)
	}
	// Standard receipt RLP omits derived block-level log indexes.
	if decoded.Logs[0].Index != 0 {
		t.Fatalf("legacy log index = %d, want 0", decoded.Logs[0].Index)
	}
}

func TestBackupReceiptRejectsInvalidData(t *testing.T) {
	t.Parallel()
	if ReceiptToBackup(nil) != nil {
		t.Fatal("nil receipt produced a backup")
	}
	if _, err := DecodeBackupReceiptRLP([]byte{0xff}); err == nil {
		t.Fatal("invalid backup RLP was accepted")
	}
	if _, _, err := TryDecodeReceiptRLP([]byte{0xff}); err == nil {
		t.Fatal("invalid receipt RLP was accepted")
	}
	wrongVersion, err := rlp.EncodeToBytes(&BackupReceipt{Version: 2})
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := TryDecodeReceiptRLP(wrongVersion); err == nil {
		t.Fatal("unsupported backup version was accepted as a receipt")
	}
}
