// Fuzz targets for the freezer receipt decoders (MODERNIZATION.md §4.4).
// These functions consume raw chain data straight from freezer files, so the
// core invariant is: arbitrary bytes must never panic or corrupt the result
// structure — only return an error.
package decode

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/snappy"
)

// encodedReceiptsSeed builds a valid RLP blob of one storage receipt with one log.
func encodedReceiptsSeed(tb testing.TB) []byte {
	tb.Helper()
	receipts := []ReceiptForStorage{{
		PostStateOrStatus: []byte{1},
		CumulativeGasUsed: 21000,
		Logs: []*types.Log{{
			Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
			Topics: []common.Hash{
				common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
				common.HexToHash("0x0000000000000000000000001111111111111111111111111111111111111111"),
			},
			Data: []byte{0x00, 0x01},
		}},
	}}
	encoded, err := rlp.EncodeToBytes(receipts)
	if err != nil {
		tb.Fatalf("encode seed receipts: %v", err)
	}
	return encoded
}

func FuzzReceiptsDecode(f *testing.F) {
	valid := encodedReceiptsSeed(f)
	f.Add([]byte(nil))
	f.Add([]byte{0xc0}) // empty RLP list
	f.Add(valid)
	f.Add(snappy.Encode(nil, valid))
	// Arbitrum Nitro style: varint(size) + 2 header bytes + raw RLP.
	nitro := append([]byte{byte(len(valid)), 0x00, 0x00}, valid...)
	f.Add(nitro)
	f.Add([]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05})
	f.Add([]byte{0xff, 0xff, 0xff, 0xff})
	f.Fuzz(func(t *testing.T, data []byte) {
		result, err := Receipts(data)
		if err != nil {
			return
		}
		if result == nil {
			t.Fatal("Receipts returned nil result with nil error")
		}
		// Structural consistency: AllLogs is exactly the receipts' logs, in order,
		// and every log's indices point back at its receipt.
		total := 0
		for i, r := range result.Receipts {
			if r == nil {
				t.Fatalf("nil receipt at index %d", i)
			}
			for j, lg := range r.Logs {
				if lg == nil {
					t.Fatalf("nil log at receipt %d log %d", i, j)
				}
				if lg.ReceiptIndex != uint(i) || lg.LogIndex != uint(j) {
					t.Fatalf("log index mismatch: receipt %d log %d carries (%d,%d)",
						i, j, lg.ReceiptIndex, lg.LogIndex)
				}
			}
			total += len(r.Logs)
		}
		if len(result.AllLogs) != total {
			t.Fatalf("AllLogs has %d entries, receipts carry %d", len(result.AllLogs), total)
		}
	})
}

func FuzzArbitrumLegacyDecode(f *testing.F) {
	valid := encodedReceiptsSeed(f)
	f.Add([]byte(nil))
	f.Add([]byte{0xc0})
	f.Add(valid)
	f.Add(snappy.Encode(nil, valid))
	f.Add([]byte{byte(len(valid)), 0x00, 0x00, 0xc1, 0x80})
	f.Add([]byte{0x01, 0x02, 0x03, 0x04})
	f.Fuzz(func(t *testing.T, data []byte) {
		logs, err := ArbitrumReceipts(data)
		if err != nil {
			return
		}
		for i, lg := range logs {
			if lg == nil {
				t.Fatalf("ArbitrumReceipts returned nil log at %d", i)
			}
		}
	})
}

func FuzzArbitrumLogDecode(f *testing.F) {
	// A well-formed log: [address, topics, data]
	good, err := rlp.EncodeToBytes([]interface{}{
		common.HexToAddress("0x1234567890123456789012345678901234567890"),
		[]common.Hash{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
		[]byte{0x01, 0x02},
	})
	if err != nil {
		f.Fatalf("encode seed log: %v", err)
	}
	f.Add(good)
	f.Add([]byte(nil))
	f.Add([]byte{0xc0})
	f.Add([]byte{0xc3, 0x01, 0x02, 0x03})
	f.Fuzz(func(t *testing.T, data []byte) {
		lg, err := ArbitrumLog(data)
		if err != nil {
			return
		}
		if lg == nil {
			t.Fatal("ArbitrumLog returned nil log with nil error")
		}
	})
}
