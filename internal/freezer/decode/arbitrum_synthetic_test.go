package decode

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/snappy"
)

// Synthetic fixtures for DecodeArbitrumReceipts, complementing the
// mainnet-data integration tests in arbitrum_decode_test.go (which skip when
// the freezer directory is absent). These pin every structural branch:
// standard vs 7-field extended receipts, typed prefixes, the fallback
// log-field scan, skip-on-bad-receipt and the smartDecompress format chain.

var arbLog = Log{
	Address: common.HexToAddress("0x5555555555555555555555555555555555555555"),
	Topics:  []common.Hash{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
	Data:    []byte{0x01, 0x02},
}

// encodeArbReceipt builds one receipt as a raw field list with the logs list
// at the given index; fields before it are scalar placeholders.
func encodeArbReceipt(t *testing.T, totalFields, logsIdx int, logs []Log) []byte {
	t.Helper()
	fields := make([]interface{}, totalFields)
	for i := range fields {
		fields[i] = []byte{byte(i + 1)}
	}
	fields[logsIdx] = logs
	return mustEncode(t, fields)
}

func TestDecodeArbitrumReceiptsStandardFormat(t *testing.T) {
	// Standard: [status, gas, bloom, logs] — logs at index 3.
	receipt := encodeArbReceipt(t, 4, 3, []Log{arbLog})
	blob := mustEncode(t, []rlp.RawValue{receipt})

	logs, err := DecodeArbitrumReceipts(blob)
	if err != nil {
		t.Fatalf("DecodeArbitrumReceipts: %v", err)
	}
	if len(logs) != 1 || logs[0].Address != arbLog.Address {
		t.Fatalf("logs = %+v, want the fixture log", logs)
	}
	if len(logs[0].Topics) != 1 || logs[0].Topics[0] != arbLog.Topics[0] {
		t.Errorf("topics = %v", logs[0].Topics)
	}
}

func TestDecodeArbitrumReceiptsExtendedFormat(t *testing.T) {
	// Arbitrum extended: 7 fields with the logs list at index 6.
	receipt := encodeArbReceipt(t, 7, 6, []Log{arbLog})
	blob := mustEncode(t, []rlp.RawValue{receipt})

	logs, err := DecodeArbitrumReceipts(blob)
	if err != nil {
		t.Fatalf("DecodeArbitrumReceipts: %v", err)
	}
	if len(logs) != 1 || logs[0].Address != arbLog.Address {
		t.Fatalf("extended-format logs = %+v", logs)
	}
}

func TestDecodeArbitrumReceiptsTypedPrefix(t *testing.T) {
	// A tx-type byte spliced before a receipt parses as its own single-byte
	// list element (RLP: bytes < 0x80 encode as themselves). The decoder must
	// take the type-strip branch on it, fail to decode the empty remainder,
	// skip it, and still return the following receipt's logs.
	receipt := encodeArbReceipt(t, 4, 3, []Log{arbLog})
	typed := append([]byte{0x68}, receipt...) // Arbitrum tx type byte < 0x80
	blob := mustEncode(t, []rlp.RawValue{typed})

	logs, err := DecodeArbitrumReceipts(blob)
	if err != nil {
		t.Fatalf("DecodeArbitrumReceipts: %v", err)
	}
	if len(logs) != 1 {
		t.Fatalf("typed receipt yielded %d logs, want 1", len(logs))
	}
}

func TestDecodeArbitrumReceiptsFallbackFieldScan(t *testing.T) {
	// Five fields with the logs list at index 4: index 3 is a scalar, so the
	// primary position fails and the fallback scan must locate the logs.
	receipt := encodeArbReceipt(t, 5, 4, []Log{arbLog})
	blob := mustEncode(t, []rlp.RawValue{receipt})

	logs, err := DecodeArbitrumReceipts(blob)
	if err != nil {
		t.Fatalf("DecodeArbitrumReceipts: %v", err)
	}
	if len(logs) != 1 {
		t.Fatalf("fallback scan yielded %d logs, want 1", len(logs))
	}
}

func TestDecodeArbitrumReceiptsSkipsBadEntries(t *testing.T) {
	good := encodeArbReceipt(t, 4, 3, []Log{arbLog})
	notAList := mustEncode(t, []byte{0xaa})              // receipt payload is a string
	tooFewFields := encodeArbReceipt(t, 3, 2, []Log{})   // no field 3 at all
	noLogsAnywhere := encodeArbReceipt(t, 5, 4, []Log{}) // empty logs everywhere
	blob := mustEncode(t, []rlp.RawValue{notAList, tooFewFields, noLogsAnywhere, good})

	logs, err := DecodeArbitrumReceipts(blob)
	if err != nil {
		t.Fatalf("DecodeArbitrumReceipts: %v", err)
	}
	if len(logs) != 1 || logs[0].Address != arbLog.Address {
		t.Fatalf("bad receipts must be skipped, good one kept; got %+v", logs)
	}
}

func TestDecodeArbitrumReceiptsMalformedLogSkipped(t *testing.T) {
	// A log with fewer than 3 fields fails DecodeArbitrumLog and is dropped.
	badLog := mustEncode(t, []interface{}{[]byte{0x01}})
	goodLog := mustEncode(t, arbLog)
	logsList := mustEncode(t, []rlp.RawValue{badLog, goodLog})

	fields := []interface{}{[]byte{1}, []byte{2}, []byte{3}, rlp.RawValue(logsList)}
	receipt := mustEncode(t, fields)
	blob := mustEncode(t, []rlp.RawValue{receipt})

	logs, err := DecodeArbitrumReceipts(blob)
	if err != nil {
		t.Fatalf("DecodeArbitrumReceipts: %v", err)
	}
	if len(logs) != 1 {
		t.Fatalf("malformed log must be dropped; got %d logs", len(logs))
	}
}

func TestDecodeArbitrumReceiptsSnappy(t *testing.T) {
	receipt := encodeArbReceipt(t, 4, 3, []Log{arbLog})
	blob := mustEncode(t, []rlp.RawValue{receipt})

	logs, err := DecodeArbitrumReceipts(snappy.Encode(nil, blob))
	if err != nil {
		t.Fatalf("DecodeArbitrumReceipts(snappy): %v", err)
	}
	if len(logs) != 1 {
		t.Fatalf("snappy round trip yielded %d logs, want 1", len(logs))
	}
}

func TestDecodeArbitrumReceiptsNitroHeader(t *testing.T) {
	receipt := encodeArbReceipt(t, 4, 3, []Log{arbLog})
	payload := mustEncode(t, []rlp.RawValue{receipt})

	// varint(size) + junk header + raw RLP.
	blob := append([]byte{byte(len(payload)), 0x00, 0x01}, payload...)
	logs, err := DecodeArbitrumReceipts(blob)
	if err != nil {
		t.Fatalf("DecodeArbitrumReceipts(nitro header): %v", err)
	}
	if len(logs) != 1 {
		t.Fatalf("nitro header format yielded %d logs, want 1", len(logs))
	}

	// Oversized declared length: the tail from the RLP prefix is used.
	blob = append([]byte{0x7f, 0x00, 0x01}, payload...)
	if _, err := DecodeArbitrumReceipts(blob); err != nil {
		t.Fatalf("oversized varint should still decode the tail: %v", err)
	}
}

func TestSmartDecompressErrors(t *testing.T) {
	cases := []struct {
		name    string
		in      []byte
		wantErr string
	}{
		{"empty", nil, ""}, // empty passes through, fails at list decode
		{"too short", []byte{0x01, 0x02}, "decompress"},
		{"zero varint", []byte{0x00, 0x01, 0x02, 0x03, 0x04}, "decompress"},
		{"no rlp after header", []byte{0x05, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, "decompress"},
		{"non-terminating varint", []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, "decompress"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := DecodeArbitrumReceipts(tc.in)
			if err == nil {
				t.Fatal("expected error")
			}
			if tc.wantErr != "" && !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("error = %v, want substring %q", err, tc.wantErr)
			}
		})
	}
}

func TestDecodeArbitrumLogErrors(t *testing.T) {
	if _, err := DecodeArbitrumLog([]byte{0xaa}); err == nil {
		t.Error("non-list log must fail")
	}
	short := mustEncode(t, []interface{}{[]byte{1}, []byte{2}})
	if _, err := DecodeArbitrumLog(short); err == nil || !strings.Contains(err.Error(), "expected 3 fields") {
		t.Errorf("two-field log: %v", err)
	}
	// Field decoders: address must be 20 bytes, topics a list of hashes,
	// data a byte string.
	badAddr := mustEncode(t, []interface{}{[]byte{1, 2, 3}, []common.Hash{}, []byte{}})
	if _, err := DecodeArbitrumLog(badAddr); err == nil || !strings.Contains(err.Error(), "decode address") {
		t.Errorf("bad address: %v", err)
	}
	badTopics := mustEncode(t, []interface{}{arbLog.Address, []byte{0x01}, []byte{}})
	if _, err := DecodeArbitrumLog(badTopics); err == nil || !strings.Contains(err.Error(), "decode topics") {
		t.Errorf("bad topics: %v", err)
	}
	badData := mustEncode(t, []interface{}{arbLog.Address, []common.Hash{}, []common.Hash{{}}})
	if _, err := DecodeArbitrumLog(badData); err == nil || !strings.Contains(err.Error(), "decode data") {
		t.Errorf("bad data: %v", err)
	}
}
