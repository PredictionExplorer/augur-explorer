package decode

import (
	"encoding/binary"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

// Fallback-chain tests: decodeStorageReceipt falls back to
// decodeReceiptAlternative (no-bloom format), which falls back to
// decodeReceiptLogsOnly (stream scan). Each layer is pinned with a synthetic
// RLP fixture of the shape that triggers it.

func mustEncode(t *testing.T, v any) []byte {
	t.Helper()
	b, err := rlp.EncodeToBytes(v)
	if err != nil {
		t.Fatalf("rlp encode: %v", err)
	}
	return b
}

var fallbackLog = &StoredLog{
	Address: common.HexToAddress("0x3333333333333333333333333333333333333333"),
	Topics:  []common.Hash{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
	Data:    []byte{0xde, 0xad},
}

// TestDecodeReceiptAlternative covers the no-bloom [status, gas, logs]
// receipt format: it does not decode as ReceiptForStorage (the logs list
// cannot fill the 256-byte Bloom), so the alternative decoder must handle it.
func TestDecodeReceiptAlternative(t *testing.T) {
	type simpleReceipt struct {
		Status            uint64
		CumulativeGasUsed uint64
		Logs              []*StoredLog
	}
	blob := mustEncode(t, []any{
		simpleReceipt{Status: 1, CumulativeGasUsed: 42000, Logs: []*StoredLog{fallbackLog}},
	})

	result, err := Receipts(blob)
	if err != nil {
		t.Fatalf("Receipts: %v", err)
	}
	if len(result.Receipts) != 1 {
		t.Fatalf("got %d receipts, want 1", len(result.Receipts))
	}
	r := result.Receipts[0]
	if r.Status != 1 || r.CumulativeGasUsed != 42000 {
		t.Errorf("status/gas = %d/%d, want 1/42000", r.Status, r.CumulativeGasUsed)
	}
	if len(r.Logs) != 1 {
		t.Fatalf("got %d logs, want 1", len(r.Logs))
	}
	log := r.Logs[0]
	if log.Address != fallbackLog.Address {
		t.Errorf("address = %s", log.Address)
	}
	if log.DataKeccak != crypto.Keccak256Hash(fallbackLog.Data) {
		t.Errorf("DataKeccak not computed for alternative-format log")
	}
}

// TestDecodeReceiptLogsOnly covers the last-resort stream scan: a four-field
// receipt whose third field is a non-bloom string defeats both struct
// decoders, but the scanner skips the scalars and still extracts the logs.
func TestDecodeReceiptLogsOnly(t *testing.T) {
	blob := mustEncode(t, []any{
		[]any{
			[]byte{1},          // status-ish scalar
			uint64(21000),      // gas
			[]byte("notbloom"), // wrong-size bloom: kills ReceiptForStorage; 4th field kills the alternative
			[]*StoredLog{fallbackLog, {Address: fallbackLog.Address}},
		},
	})

	result, err := Receipts(blob)
	if err != nil {
		t.Fatalf("Receipts: %v", err)
	}
	if len(result.Receipts) != 1 {
		t.Fatalf("got %d receipts, want 1", len(result.Receipts))
	}
	r := result.Receipts[0]
	// The logs-only path cannot recover status/gas; it must still find both logs.
	if len(r.Logs) != 2 {
		t.Fatalf("got %d logs, want 2", len(r.Logs))
	}
	if r.Logs[0].Address != fallbackLog.Address || r.Logs[0].LogIndex != 0 || r.Logs[1].LogIndex != 1 {
		t.Errorf("logs = %+v", r.Logs)
	}
	// The second log has no data, so its keccak stays zero.
	if r.Logs[1].DataKeccak != (common.Hash{}) {
		t.Errorf("empty-data log should have zero DataKeccak")
	}
}

// TestDecodeReceiptLogsOnlyNoLogs pins the soft-empty result when the
// unrecognized receipt has no list field at all.
func TestDecodeReceiptLogsOnlyNoLogs(t *testing.T) {
	blob := mustEncode(t, []any{
		[]any{[]byte("only"), []byte("scalar"), []byte("fields"), []byte("here")},
	})
	result, err := Receipts(blob)
	if err != nil {
		t.Fatalf("Receipts: %v", err)
	}
	if len(result.Receipts) != 1 || len(result.Receipts[0].Logs) != 0 {
		t.Fatalf("want one log-less receipt, got %+v", result.Receipts)
	}
}

// TestDecodeReceiptLogsOnlyNotAList pins the terminal error when a receipt
// element is not even an RLP list.
func TestDecodeReceiptLogsOnlyNotAList(t *testing.T) {
	blob := mustEncode(t, []any{[]byte{0xaa, 0xbb}})
	if _, err := Receipts(blob); err == nil || !strings.Contains(err.Error(), "expected list") {
		t.Fatalf("want expected-list error, got %v", err)
	}
}

// TestDecodeStorageReceiptPreByzantium pins the post-state-root branch: a
// 32-byte PostStateOrStatus means pre-Byzantium, reported as Status 1.
func TestDecodeStorageReceiptPreByzantium(t *testing.T) {
	root := make([]byte, 32)
	root[0] = 0xab
	blob := mustEncode(t, []ReceiptForStorage{{
		PostStateOrStatus: root,
		CumulativeGasUsed: 100,
		Bloom:             types.Bloom{},
		Logs:              nil,
	}})
	result, err := Receipts(blob)
	if err != nil {
		t.Fatalf("Receipts: %v", err)
	}
	if got := result.Receipts[0].Status; got != 1 {
		t.Errorf("pre-Byzantium receipt status = %d, want 1", got)
	}
}

// TestDecodeStorageReceiptTypeByteElement pins the type-byte strip branch:
// a bare tx-type byte spliced before a receipt parses as its own single-byte
// list element. The strip leaves an empty payload that defeats all three
// decoders, so the legacy decoder rejects the whole blob (unlike
// ArbitrumReceipts, which skips the stray element).
func TestDecodeStorageReceiptTypeByteElement(t *testing.T) {
	inner := mustEncode(t, ReceiptForStorage{
		PostStateOrStatus: []byte{1},
		CumulativeGasUsed: 500,
		Bloom:             types.Bloom{},
	})
	blob := mustEncode(t, []rlp.RawValue{{0x02}, inner})

	if _, err := Receipts(blob); err == nil || !strings.Contains(err.Error(), "failed to decode receipt 0") {
		t.Fatalf("stray type byte should fail the legacy decoder, got %v", err)
	}
}

func TestReceiptsMalformedTopLevel(t *testing.T) {
	// Valid RLP header covering the payload, but the payload is a string,
	// not a receipts list.
	if _, err := Receipts([]byte{0x81, 0x80}); err == nil {
		t.Fatal("string payload should fail the top-level list decode")
	}
}

func TestSkipStreamValueRejectsList(t *testing.T) {
	stream := rlp.NewStream(strings.NewReader(string([]byte{0xc1, 0x01})), 0)
	if err := skipStreamValue(stream); err == nil {
		t.Fatal("skipStreamValue must refuse to skip a list")
	}
}

func TestDecodeVarint(t *testing.T) {
	cases := []struct {
		in        []byte
		wantVal   uint64
		wantBytes int
	}{
		{[]byte{0x05}, 5, 1},
		{[]byte{0x80, 0x01}, 128, 2},                 // two-byte varint
		{[]byte{0xff, 0xff}, 0x3fff, 0},              // never terminates -> 0 consumed
		{[]byte{}, 0, 0},                             // empty input
		{[]byte{0xac, 0x02}, 0x12c, 2},               // 300
		{[]byte{0x80, 0x80, 0x80, 0x01}, 1 << 21, 4}, // multi-byte continuation
	}
	for _, tc := range cases {
		gotVal, gotBytes := decodeVarint(tc.in)
		if gotBytes != tc.wantBytes || (gotBytes > 0 && gotVal != tc.wantVal) {
			t.Errorf("decodeVarint(%x) = (%d, %d), want (%d, %d)", tc.in, gotVal, gotBytes, tc.wantVal, tc.wantBytes)
		}
	}
}

// TestTrySnappyDecompressNitroHeader covers the Arbitrum Nitro fallback in
// the legacy decoder: varint size + header bytes + raw RLP.
func TestTrySnappyDecompressNitroHeader(t *testing.T) {
	payload := mustEncode(t, []ReceiptForStorage{{
		PostStateOrStatus: []byte{1},
		CumulativeGasUsed: 1,
		Bloom:             types.Bloom{},
	}})

	// varint(len(payload)) + two junk header bytes + payload
	blob := append(binary.AppendUvarint(nil, uint64(len(payload))), 0x00, 0x01)
	blob = append(blob, payload...)
	result, err := Receipts(blob)
	if err != nil {
		t.Fatalf("Receipts(nitro header): %v", err)
	}
	if len(result.Receipts) != 1 {
		t.Fatalf("got %d receipts, want 1", len(result.Receipts))
	}

	// Declared size larger than the remaining data: decoder returns the tail
	// from the RLP prefix onward, which still decodes.
	blob = append(binary.AppendUvarint(nil, uint64(len(payload)*10)), 0x00, 0x01)
	blob = append(blob, payload...)
	if _, err := Receipts(blob); err != nil {
		t.Fatalf("Receipts(nitro header, oversized varint): %v", err)
	}

	// No RLP prefix within 10 bytes of the varint: data returned unchanged
	// and the top-level decode fails.
	junk := []byte{0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	if _, err := Receipts(junk); err == nil {
		t.Fatal("junk without RLP should fail to decode")
	}

	// Non-terminating varint: data returned unchanged, decode fails.
	if _, err := Receipts([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}); err == nil {
		t.Fatal("non-terminating varint should fail to decode")
	}
}
