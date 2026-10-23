package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

func writeReceiptFixture(t *testing.T, dir, block, txHash string, blob []byte) {
	t.Helper()
	blockDir := filepath.Join(dir, block)
	if err := os.MkdirAll(blockDir, 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(blockDir, txHash+"_receipt.rlp"), blob, 0o600); err != nil {
		t.Fatal(err)
	}
}

func receiptBlob(t *testing.T, addrs ...common.Address) []byte {
	t.Helper()
	r := &types.Receipt{Status: types.ReceiptStatusSuccessful}
	for _, a := range addrs {
		r.Logs = append(r.Logs, &types.Log{Address: a})
	}
	blob, err := toolutil.EncodeBackupReceiptRLP(r)
	if err != nil {
		t.Fatal(err)
	}
	return blob
}

func TestArchiveLogCounts(t *testing.T) {
	tracked1 := common.HexToAddress("0x1111111111111111111111111111111111111111")
	tracked2 := common.HexToAddress("0x2222222222222222222222222222222222222222")
	ignored := common.HexToAddress("0x3333333333333333333333333333333333333333")
	tracked := map[string]struct{}{
		"0x1111111111111111111111111111111111111111": {},
		"0x2222222222222222222222222222222222222222": {},
	}

	dir := t.TempDir()
	// Block 100: two logs for tracked1 plus one untracked in one tx,
	// one tracked2 log in another.
	writeReceiptFixture(t, dir, "100", "0xaa", receiptBlob(t, tracked1, ignored, tracked1))
	writeReceiptFixture(t, dir, "100", "0xbb", receiptBlob(t, tracked2))
	// Block 101: a legacy-format (go-ethereum consensus RLP) receipt.
	legacy, err := rlp.EncodeToBytes(&types.Receipt{
		Status: types.ReceiptStatusSuccessful,
		Logs:   []*types.Log{{Address: tracked2}},
	})
	if err != nil {
		t.Fatal(err)
	}
	writeReceiptFixture(t, dir, "101", "0xcc", legacy)
	// Block 999 is beyond dbHead and must be skipped entirely.
	writeReceiptFixture(t, dir, "999", "0xdd", receiptBlob(t, tracked1))

	counts, nReceipts, decodeErrs, err := archiveLogCounts(dir, tracked, 200)
	if err != nil {
		t.Fatal(err)
	}
	if len(decodeErrs) != 0 {
		t.Fatalf("decode errors: %v", decodeErrs)
	}
	if nReceipts != 3 {
		t.Fatalf("nReceipts = %d, want 3 (block 999 excluded)", nReceipts)
	}
	want := map[txAddrKey]uint64{
		{"0xaa", "0x1111111111111111111111111111111111111111"}: 2,
		{"0xbb", "0x2222222222222222222222222222222222222222"}: 1,
		{"0xcc", "0x2222222222222222222222222222222222222222"}: 1,
	}
	if !reflect.DeepEqual(counts, want) {
		t.Fatalf("counts = %v, want %v", counts, want)
	}
}

func TestArchiveLogCountsReportsCorruptReceipts(t *testing.T) {
	dir := t.TempDir()
	writeReceiptFixture(t, dir, "100", "0xaa", []byte("not rlp"))
	_, nReceipts, decodeErrs, err := archiveLogCounts(dir, map[string]struct{}{}, 200)
	if err != nil {
		t.Fatal(err)
	}
	if nReceipts != 1 || len(decodeErrs) != 1 {
		t.Fatalf("nReceipts=%d decodeErrs=%v, want 1 receipt with 1 decode error",
			nReceipts, decodeErrs)
	}
}

func TestChainScanStateRoundtrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sub", "state.json")
	in := chainScanState{
		"42161:platform": {From: 10, ScannedTo: 99,
			Counts: map[string]uint64{"0xabc": 7}},
	}
	if err := saveChainScanState(path, in); err != nil {
		t.Fatal(err)
	}
	out := loadChainScanState(path)
	if !reflect.DeepEqual(in, out) {
		t.Fatalf("roundtrip mismatch: %v != %v", out, in)
	}
	if got := loadChainScanState(filepath.Join(t.TempDir(), "missing.json")); len(got) != 0 {
		t.Fatalf("missing file should yield empty state, got %v", got)
	}
}
