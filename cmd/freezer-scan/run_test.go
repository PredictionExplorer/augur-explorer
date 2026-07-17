package main

// run() tests over a synthetic freezer directory: flag validation, --info
// mode and an end-to-end scan through the flag surface. The scan pipeline
// itself is tested exhaustively in internal/freezer/scan.

import (
	"bytes"
	"context"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	freezerscanner "github.com/PredictionExplorer/augur-explorer/internal/freezer"
	"github.com/PredictionExplorer/augur-explorer/internal/freezer/decode"
)

// setFlags applies flag values for one test and restores defaults after.
func setFlags(t *testing.T, values map[string]string) {
	t.Helper()
	for name, value := range values {
		f := flag.Lookup(name)
		if f == nil {
			t.Fatalf("unknown flag %q", name)
		}
		old := f.Value.String()
		if err := f.Value.Set(value); err != nil {
			t.Fatalf("setting flag %q: %v", name, err)
		}
		t.Cleanup(func() {
			_ = f.Value.Set(old)
		})
	}
}

// writeTestFreezer builds a one-cdat freezer with two blocks of receipts.
func writeTestFreezer(t *testing.T) (dir string, addr common.Address) {
	t.Helper()
	dir = t.TempDir()
	addr = common.HexToAddress("0xaa00000000000000000000000000000000000001")

	blob := func(marker byte) []byte {
		receipt := decode.ReceiptForStorage{
			PostStateOrStatus: []byte{1},
			CumulativeGasUsed: 21000,
			Logs: []*types.Log{{
				Address: addr,
				Topics:  []common.Hash{common.HexToHash("0x01")},
				Data:    []byte{marker},
			}},
		}
		encoded, err := rlp.EncodeToBytes([]decode.ReceiptForStorage{receipt})
		if err != nil {
			t.Fatalf("encoding receipts: %v", err)
		}
		return encoded
	}

	blobs := [][]byte{blob(0), blob(1), blob(2)}
	var data []byte
	index := make([]byte, 0, len(blobs)*freezerscanner.IndexEntrySize)
	for _, b := range blobs {
		index = append(index, freezerscanner.Uint48ToBytes(uint64(len(data)))...)
		data = append(data, b...)
	}
	if err := os.WriteFile(filepath.Join(dir, "receipts.cidx"), index, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "receipts.0000.cdat"), data, 0o600); err != nil {
		t.Fatal(err)
	}
	return dir, addr
}

func TestRunRequiresDirFlag(t *testing.T) {
	setFlags(t, map[string]string{"ancientDir": "", "receiptsCidx": ""})
	err := run(context.Background(), &bytes.Buffer{})
	if err == nil || !strings.Contains(err.Error(), "--ancientDir or --receiptsCidx is required") {
		t.Errorf("run = %v, want missing-flag error", err)
	}
}

func TestRunRejectsBadFilterFlags(t *testing.T) {
	dir, _ := writeTestFreezer(t)
	setFlags(t, map[string]string{"ancientDir": dir, "contracts": "0x123"})
	if err := run(context.Background(), &bytes.Buffer{}); err == nil ||
		!strings.Contains(err.Error(), "invalid contract address") {
		t.Errorf("run = %v, want contract parse error", err)
	}

	setFlags(t, map[string]string{"ancientDir": dir, "contracts": "", "eventSigs": "0xbad"})
	if err := run(context.Background(), &bytes.Buffer{}); err == nil ||
		!strings.Contains(err.Error(), "invalid event signature") {
		t.Errorf("run = %v, want event sig parse error", err)
	}
}

func TestRunFailsOnMissingFreezer(t *testing.T) {
	setFlags(t, map[string]string{"ancientDir": t.TempDir()})
	if err := run(context.Background(), &bytes.Buffer{}); err == nil ||
		!strings.Contains(err.Error(), "failed to open freezer") {
		t.Errorf("run = %v, want open failure", err)
	}
}

func TestRunInfoMode(t *testing.T) {
	dir, _ := writeTestFreezer(t)
	setFlags(t, map[string]string{"ancientDir": dir, "info": "true"})
	var out bytes.Buffer
	if err := run(context.Background(), &out); err != nil {
		t.Fatalf("run --info: %v", err)
	}
	for _, want := range []string{"Freezer Information:", "Item count (blocks): 3"} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("--info output missing %q:\n%s", want, out.String())
		}
	}
}

func TestRunScansThroughFlagSurface(t *testing.T) {
	dir, addr := writeTestFreezer(t)
	setFlags(t, map[string]string{
		"ancientDir": dir,
		"contracts":  addr.Hex(),
		"out":        "-",
		"workers":    "1",
	})
	var out bytes.Buffer
	if err := run(context.Background(), &out); err != nil {
		t.Fatalf("run: %v", err)
	}
	// Blocks 0 and 1 are scannable (end clamps to MaxAvailableBlock 2).
	lines := strings.Count(strings.TrimSpace(out.String()), "\n") + 1
	if lines != 2 {
		t.Errorf("output lines = %d, want 2 records:\n%s", lines, out.String())
	}
	if !strings.Contains(out.String(), strings.ToLower(addr.Hex()[2:6])) &&
		!strings.Contains(out.String(), addr.Hex()) {
		t.Errorf("output missing contract address:\n%s", out.String())
	}
}

func TestRunWrapsScanFailure(t *testing.T) {
	dir, _ := writeTestFreezer(t)
	setFlags(t, map[string]string{
		"ancientDir": dir,
		"out":        filepath.Join(t.TempDir(), "missing-dir", "out.jsonl"),
		"workers":    "1",
	})
	err := run(context.Background(), &bytes.Buffer{})
	if err == nil || !strings.Contains(err.Error(), "scan failed") {
		t.Errorf("run = %v, want wrapped scan failure", err)
	}
}

func TestRunResolvesDirFromCidxPath(t *testing.T) {
	dir, _ := writeTestFreezer(t)
	setFlags(t, map[string]string{
		"ancientDir":   "",
		"receiptsCidx": dir + "/receipts.cidx",
		"info":         "true",
	})
	var out bytes.Buffer
	if err := run(context.Background(), &out); err != nil {
		t.Fatalf("run with --receiptsCidx: %v", err)
	}
	if !strings.Contains(out.String(), "Item count (blocks): 3") {
		t.Errorf("--receiptsCidx info output:\n%s", out.String())
	}
}
