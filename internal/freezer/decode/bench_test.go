package decode

// Benchmarks for the freezer receipts decoder (§4.5 of docs/MODERNIZATION.md).
// Baselines live in docs/benchmarks.md; re-run with:
//
//	go test ./internal/freezer/decode/ -bench BenchmarkReceiptsDecode -benchmem -count=6

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/snappy"
)

// buildReceiptsBlob encodes a realistic block payload: 10 receipts with 3
// Transfer-shaped logs each.
func buildReceiptsBlob(b *testing.B) []byte {
	b.Helper()
	receipts := make([]ReceiptForStorage, 0, 10)
	for i := byte(0); i < 10; i++ {
		logs := make([]*types.Log, 0, 3)
		for j := byte(0); j < 3; j++ {
			logs = append(logs, &types.Log{
				Address: common.BytesToAddress([]byte{0x20, i, j}),
				Topics: []common.Hash{
					common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
					common.BytesToHash([]byte{0x11, i}),
					common.BytesToHash([]byte{0x22, j}),
				},
				Data: make([]byte, 32),
			})
		}
		receipts = append(receipts, ReceiptForStorage{
			PostStateOrStatus: []byte{1},
			CumulativeGasUsed: 21000 * uint64(i+1),
			Logs:              logs,
		})
	}
	encoded, err := rlp.EncodeToBytes(receipts)
	if err != nil {
		b.Fatalf("encoding receipts: %v", err)
	}
	return encoded
}

func BenchmarkReceiptsDecode(b *testing.B) {
	raw := buildReceiptsBlob(b)
	compressed := snappy.Encode(nil, raw)

	b.Run("raw_rlp", func(b *testing.B) {
		b.SetBytes(int64(len(raw)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out, err := DecodeReceipts(raw)
			if err != nil {
				b.Fatalf("decode: %v", err)
			}
			if len(out.Receipts) != 10 {
				b.Fatalf("got %d receipts, want 10", len(out.Receipts))
			}
		}
	})

	b.Run("snappy", func(b *testing.B) {
		b.SetBytes(int64(len(compressed)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out, err := DecodeReceipts(compressed)
			if err != nil {
				b.Fatalf("decode: %v", err)
			}
			if len(out.Receipts) != 10 {
				b.Fatalf("got %d receipts, want 10", len(out.Receipts))
			}
		}
	})
}
