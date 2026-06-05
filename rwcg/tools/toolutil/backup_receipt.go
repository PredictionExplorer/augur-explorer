package toolutil

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

const backupReceiptVersion = 1

// BackupReceipt is our on-disk receipt RLP format. It preserves block log indices
// (types.Log.Index is not part of consensus log RLP inside a standard receipt).
type BackupReceipt struct {
	Version           uint8
	Type              uint8
	Status            uint64
	CumulativeGasUsed uint64
	Bloom             types.Bloom
	Logs              []BackupLog
}

// BackupLog is one event log with its block-level log index.
type BackupLog struct {
	Address common.Address
	Topics  []common.Hash
	Data    []byte
	Index   uint
}

// ReceiptToBackup copies an RPC receipt (with derived log fields) into BackupReceipt.
func ReceiptToBackup(r *types.Receipt) *BackupReceipt {
	if r == nil {
		return nil
	}
	out := &BackupReceipt{
		Version:           backupReceiptVersion,
		Type:              r.Type,
		Status:            r.Status,
		CumulativeGasUsed: r.CumulativeGasUsed,
		Bloom:             r.Bloom,
		Logs:              make([]BackupLog, 0, len(r.Logs)),
	}
	for _, lg := range r.Logs {
		if lg == nil {
			continue
		}
		out.Logs = append(out.Logs, BackupLog{
			Address: lg.Address,
			Topics:  lg.Topics,
			Data:    lg.Data,
			Index:   lg.Index,
		})
	}
	return out
}

// EncodeBackupReceiptRLP stores a receipt for backup with log indices preserved.
func EncodeBackupReceiptRLP(r *types.Receipt) ([]byte, error) {
	return rlp.EncodeToBytes(ReceiptToBackup(r))
}

// DecodeBackupReceiptRLP reads our backup receipt format.
func DecodeBackupReceiptRLP(data []byte) (*BackupReceipt, error) {
	var br BackupReceipt
	if err := rlp.DecodeBytes(data, &br); err != nil {
		return nil, err
	}
	return &br, nil
}

// ToTypesLog builds a types.Log for RLP comparison with evt_log.log_rlp.
func (l *BackupLog) ToTypesLog() *types.Log {
	return &types.Log{
		Address: l.Address,
		Topics:  l.Topics,
		Data:    l.Data,
		Index:   l.Index,
	}
}

// TryDecodeReceiptRLP decodes v1 backup format, or falls back to legacy types.Receipt RLP.
// legacyFormat is true when the blob was standard geth receipt RLP (log Index not stored).
func TryDecodeReceiptRLP(data []byte) (br *BackupReceipt, legacyFormat bool, err error) {
	if br, err = DecodeBackupReceiptRLP(data); err == nil && br.Version == backupReceiptVersion {
		return br, false, nil
	}
	var legacy types.Receipt
	if err := rlp.DecodeBytes(data, &legacy); err != nil {
		return nil, false, err
	}
	return ReceiptToBackup(&legacy), true, nil
}
