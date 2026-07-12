package scan

// The whole pipeline runs against synthetic freezer fixtures: a receipts.cidx
// index plus receipts.NNNN.cdat data files built in a temp dir, with each
// block's payload being real RLP-encoded storage receipts. No mainnet data
// is required.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	freezerscanner "github.com/PredictionExplorer/augur-explorer/internal/freezer"
	"github.com/PredictionExplorer/augur-explorer/internal/freezer/decode"
)

var (
	contractA = common.HexToAddress("0xaa00000000000000000000000000000000000001")
	contractB = common.HexToAddress("0xbb00000000000000000000000000000000000002")
	topicX    = common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
	topicY    = common.HexToHash("0x2222222222222222222222222222222222222222222222222222222222222222")
)

// receiptsBlob RLP-encodes one block's receipts, each carrying the given
// logs, in the standard 4-field storage format the Arbitrum decoder handles.
func receiptsBlob(t *testing.T, logs []*types.Log) []byte {
	t.Helper()
	receipt := decode.ReceiptForStorage{
		PostStateOrStatus: []byte{1},
		CumulativeGasUsed: 21000,
		Bloom:             types.Bloom{},
		Logs:              logs,
	}
	encoded, err := rlp.EncodeToBytes([]decode.ReceiptForStorage{receipt})
	if err != nil {
		t.Fatalf("encoding receipts: %v", err)
	}
	return encoded
}

// simpleLog builds a log for addr with topic0 and one data byte marking the
// block, so records are distinguishable in the output.
func simpleLog(addr common.Address, topic0 common.Hash, marker byte) *types.Log {
	return &types.Log{
		Address: addr,
		Topics:  []common.Hash{topic0},
		Data:    []byte{marker},
	}
}

// buildFreezer writes receipts.cidx + receipts.0000.cdat for the given
// per-block payloads (nil payload = empty block) and opens a reader on them.
func buildFreezer(t *testing.T, blocks [][]byte) *freezerscanner.ParallelReader {
	t.Helper()
	dir := t.TempDir()

	var data []byte
	offsets := make([]uint64, 0, len(blocks))
	for _, b := range blocks {
		offsets = append(offsets, uint64(len(data)))
		data = append(data, b...)
	}

	index := make([]byte, 0, len(offsets)*6)
	for _, off := range offsets {
		entry := []byte{
			byte(off >> 40), byte(off >> 32), byte(off >> 24),
			byte(off >> 16), byte(off >> 8), byte(off),
		}
		index = append(index, entry...)
	}
	if err := os.WriteFile(filepath.Join(dir, "receipts.cidx"), index, 0o600); err != nil {
		t.Fatalf("writing cidx: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "receipts.0000.cdat"), data, 0o600); err != nil {
		t.Fatalf("writing cdat: %v", err)
	}

	reader, err := freezerscanner.NewParallelReader(dir)
	if err != nil {
		t.Fatalf("NewParallelReader: %v", err)
	}
	return reader
}

// readOutput reads a scan output file written under t.TempDir.
func readOutput(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path) //nolint:gosec // test output path under t.TempDir
	if err != nil {
		t.Fatalf("reading output %s: %v", path, err)
	}
	return data
}

// standardBlocks builds a five-block fixture:
//
//	block 0: contractA/topicX
//	block 1: empty
//	block 2: contractB/topicY
//	block 3: contractA/topicY + contractB/topicX (two logs)
//	block 4: tail data so MaxAvailableBlock covers block 3
func standardBlocks(t *testing.T) [][]byte {
	t.Helper()
	return [][]byte{
		receiptsBlob(t, []*types.Log{simpleLog(contractA, topicX, 0)}),
		nil,
		receiptsBlob(t, []*types.Log{simpleLog(contractB, topicY, 2)}),
		receiptsBlob(t, []*types.Log{
			simpleLog(contractA, topicY, 3),
			simpleLog(contractB, topicX, 3),
		}),
		receiptsBlob(t, []*types.Log{simpleLog(contractA, topicX, 4)}),
	}
}

// records parses JSONL output lines.
func records(t *testing.T, jsonl []byte) []map[string]any {
	t.Helper()
	var out []map[string]any
	for _, line := range bytes.Split(bytes.TrimSpace(jsonl), []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		var m map[string]any
		if err := json.Unmarshal(line, &m); err != nil {
			t.Fatalf("parsing output line %q: %v", line, err)
		}
		out = append(out, m)
	}
	return out
}

func TestScanToStdoutMatchesAllLogs(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	var out bytes.Buffer

	stats, err := Run(context.Background(), reader, Options{
		StartBlock: 0,
		EndBlock:   0, // clamp to available data
		OutPath:    "-",
		Stdout:     &out,
		Workers:    2,
		ChunkSize:  2,
	})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}

	// Blocks 0..3 are scannable (EndBlock clamps to MaxAvailableBlock 4,
	// exclusive), containing 4 logs total.
	if stats.EndBlock != 4 {
		t.Errorf("effective end = %d, want clamp to 4", stats.EndBlock)
	}
	if stats.ProcessedBlocks != 4 || stats.TotalLogs != 4 || stats.MatchedLogs != 4 || stats.Errors != 0 {
		t.Errorf("stats = %+v, want 4 blocks / 4 logs / 4 matched", stats)
	}

	recs := records(t, out.Bytes())
	if len(recs) != 4 {
		t.Fatalf("records = %d, want 4", len(recs))
	}
	// Chunk-ordered merge keeps blocks ascending.
	wantBlocks := []float64{0, 2, 3, 3}
	for i, rec := range recs {
		if rec["blockNumber"] != wantBlocks[i] {
			t.Errorf("record %d block = %v, want %v", i, rec["blockNumber"], wantBlocks[i])
		}
	}
	if recs[0]["contract"] != contractA.Hex() || recs[0]["topic0"] != topicX.Hex() {
		t.Errorf("record 0 = %+v", recs[0])
	}
}

func TestScanFiltersByContractAndTopic(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))

	t.Run("contract filter", func(t *testing.T) {
		var out bytes.Buffer
		stats, err := Run(context.Background(), reader, Options{
			OutPath:   "-",
			Stdout:    &out,
			Contracts: map[common.Address]bool{contractA: true},
			Workers:   1,
		})
		if err != nil {
			t.Fatalf("Run: %v", err)
		}
		recs := records(t, out.Bytes())
		if stats.MatchedLogs != 2 || len(recs) != 2 {
			t.Fatalf("matched = %d, records = %d, want 2 contractA logs", stats.MatchedLogs, len(recs))
		}
		for _, rec := range recs {
			if rec["contract"] != contractA.Hex() {
				t.Errorf("record = %+v, want only contractA", rec)
			}
		}
	})

	t.Run("topic filter", func(t *testing.T) {
		var out bytes.Buffer
		stats, err := Run(context.Background(), reader, Options{
			OutPath:   "-",
			Stdout:    &out,
			EventSigs: map[common.Hash]bool{topicY: true},
			Workers:   1,
		})
		if err != nil {
			t.Fatalf("Run: %v", err)
		}
		if stats.MatchedLogs != 2 {
			t.Errorf("matched = %d, want the two topicY logs", stats.MatchedLogs)
		}
	})

	t.Run("combined filter", func(t *testing.T) {
		var out bytes.Buffer
		stats, err := Run(context.Background(), reader, Options{
			OutPath:   "-",
			Stdout:    &out,
			Contracts: map[common.Address]bool{contractA: true},
			EventSigs: map[common.Hash]bool{topicY: true},
			Workers:   1,
		})
		if err != nil {
			t.Fatalf("Run: %v", err)
		}
		recs := records(t, out.Bytes())
		if stats.MatchedLogs != 1 || len(recs) != 1 {
			t.Fatalf("matched = %d, want exactly contractA+topicY", stats.MatchedLogs)
		}
		if recs[0]["blockNumber"] != float64(3) {
			t.Errorf("record = %+v, want block 3", recs[0])
		}
	})
}

func TestScanIncludeDataEmbedsHex(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	var out bytes.Buffer
	if _, err := Run(context.Background(), reader, Options{
		OutPath: "-", Stdout: &out, IncludeData: true, Workers: 1,
	}); err != nil {
		t.Fatalf("Run: %v", err)
	}
	recs := records(t, out.Bytes())
	if recs[0]["dataHex"] != "0x00" {
		t.Errorf("dataHex = %v, want embedded 0x00 marker", recs[0]["dataHex"])
	}

	out.Reset()
	if _, err := Run(context.Background(), reader, Options{
		OutPath: "-", Stdout: &out, Workers: 1,
	}); err != nil {
		t.Fatalf("Run: %v", err)
	}
	recs = records(t, out.Bytes())
	if _, ok := recs[0]["dataHex"]; ok {
		t.Errorf("dataHex present without IncludeData: %+v", recs[0])
	}
}

func TestScanWritesFileThenResumesAppending(t *testing.T) {
	blocks := standardBlocks(t)
	reader := buildFreezer(t, blocks)
	outPath := filepath.Join(t.TempDir(), "events.jsonl")

	// First run covers blocks 0..3 only.
	stats, err := Run(context.Background(), reader, Options{
		StartBlock: 0, EndBlock: 3, OutPath: outPath, Workers: 1,
	})
	if err != nil {
		t.Fatalf("first Run: %v", err)
	}
	if stats.Resumed {
		t.Error("first run reported resumed")
	}

	// Second run resumes after block 2 (the last block in the file).
	var logLines []string
	stats, err = Run(context.Background(), reader, Options{
		StartBlock: 0, EndBlock: 0, OutPath: outPath, Workers: 1,
		Logf: func(format string, args ...any) { logLines = append(logLines, fmt.Sprintf(format, args...)) },
	})
	if err != nil {
		t.Fatalf("second Run: %v", err)
	}
	if !stats.Resumed || stats.StartBlock != 3 {
		t.Errorf("stats = %+v, want resume from block 3", stats)
	}
	found := false
	for _, l := range logLines {
		if strings.Contains(l, "Resuming from block 3") {
			found = true
		}
	}
	if !found {
		t.Errorf("resume log line missing: %v", logLines)
	}

	recs := records(t, readOutput(t, outPath))
	if len(recs) != 4 {
		t.Fatalf("records after resume = %d, want 4 with no duplicates", len(recs))
	}
	wantBlocks := []float64{0, 2, 3, 3}
	for i, rec := range recs {
		if rec["blockNumber"] != wantBlocks[i] {
			t.Errorf("record %d block = %v, want %v", i, rec["blockNumber"], wantBlocks[i])
		}
	}
}

func TestScanNoResumeTruncates(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	outPath := filepath.Join(t.TempDir(), "events.jsonl")

	if _, err := Run(context.Background(), reader, Options{OutPath: outPath, Workers: 1}); err != nil {
		t.Fatalf("first Run: %v", err)
	}
	stats, err := Run(context.Background(), reader, Options{OutPath: outPath, NoResume: true, Workers: 1})
	if err != nil {
		t.Fatalf("second Run: %v", err)
	}
	if stats.Resumed {
		t.Error("NoResume run reported resumed")
	}
	if recs := records(t, readOutput(t, outPath)); len(recs) != 4 {
		t.Errorf("records = %d, want a fresh full scan (4), not appended duplicates", len(recs))
	}
}

func TestScanAlreadyCompleteIsNoOp(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	outPath := filepath.Join(t.TempDir(), "events.jsonl")

	if _, err := Run(context.Background(), reader, Options{OutPath: outPath, Workers: 1}); err != nil {
		t.Fatalf("first Run: %v", err)
	}
	before := readOutput(t, outPath)

	stats, err := Run(context.Background(), reader, Options{OutPath: outPath, Workers: 1})
	if err != nil {
		t.Fatalf("second Run: %v", err)
	}
	if stats.ProcessedBlocks != 0 {
		t.Errorf("processed = %d, want 0 for an already-complete scan", stats.ProcessedBlocks)
	}
	after := readOutput(t, outPath)
	if !bytes.Equal(before, after) {
		t.Error("already-complete run modified the output file")
	}
}

func TestScanOutputDirectoryGetsRangeNamedFile(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	outDir := t.TempDir()

	stats, err := Run(context.Background(), reader, Options{OutPath: outDir, Workers: 1})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	want := filepath.Join(outDir, "events_0_4.jsonl")
	if stats.OutputPath != want {
		t.Errorf("output path = %s, want %s", stats.OutputPath, want)
	}
	if _, err := os.Stat(want); err != nil {
		t.Errorf("range-named output file missing: %v", err)
	}
}

func TestScanBestEffortVsFailFast(t *testing.T) {
	blocks := standardBlocks(t)
	blocks[2] = []byte{0xff, 0x01, 0x02} // undecodable payload
	reader := buildFreezer(t, blocks)

	t.Run("fail fast", func(t *testing.T) {
		var out bytes.Buffer
		_, err := Run(context.Background(), reader, Options{OutPath: "-", Stdout: &out, Workers: 1})
		if err == nil || !strings.Contains(err.Error(), "block 2 decode") {
			t.Errorf("Run = %v, want block-2 decode failure", err)
		}
	})

	t.Run("best effort with error log", func(t *testing.T) {
		var out bytes.Buffer
		errLogPath := filepath.Join(t.TempDir(), "errors.log")
		stats, err := Run(context.Background(), reader, Options{
			OutPath: "-", Stdout: &out, BestEffort: true, ErrorLogPath: errLogPath, Workers: 1,
		})
		if err != nil {
			t.Fatalf("Run: %v", err)
		}
		if stats.Errors != 1 {
			t.Errorf("errors = %d, want 1", stats.Errors)
		}
		if stats.MatchedLogs != 3 {
			t.Errorf("matched = %d, want the 3 logs outside the corrupt block", stats.MatchedLogs)
		}
		logData, err := os.ReadFile(errLogPath) //nolint:gosec // test path under t.TempDir
		if err != nil {
			t.Fatalf("reading error log: %v", err)
		}
		text := string(logData)
		if !strings.Contains(text, "Scan session started") ||
			!strings.Contains(text, "block=2 type=decode") ||
			!strings.Contains(text, "Scan session ended") {
			t.Errorf("error log = %q, want session frame + block 2 decode entry", text)
		}
	})
}

func TestScanCancellationReturnsPartialResults(t *testing.T) {
	// Many blocks and one worker with tiny chunks: cancellation between
	// chunks must stop the scan without an error.
	var blocks [][]byte
	for i := 0; i < 50; i++ {
		blocks = append(blocks, receiptsBlob(t, []*types.Log{simpleLog(contractA, topicX, byte(i))}))
	}
	reader := buildFreezer(t, blocks)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancelled before the scan starts: nothing gets queued

	var out bytes.Buffer
	stats, err := Run(ctx, reader, Options{OutPath: "-", Stdout: &out, Workers: 1, ChunkSize: 5})
	if err != nil {
		t.Fatalf("Run after cancel: %v", err)
	}
	if stats.ProcessedBlocks != 0 {
		t.Errorf("processed = %d, want 0 for pre-cancelled scan", stats.ProcessedBlocks)
	}
}

func TestParseContracts(t *testing.T) {
	got, err := ParseContracts(fmt.Sprintf(" %s , %s ,", contractA.Hex(), contractB.Hex()))
	if err != nil {
		t.Fatalf("ParseContracts: %v", err)
	}
	if len(got) != 2 || !got[contractA] || !got[contractB] {
		t.Errorf("parsed = %v", got)
	}
	if _, err := ParseContracts("0x123"); err == nil {
		t.Error("short address accepted")
	}
	if got, err := ParseContracts(""); err != nil || len(got) != 0 {
		t.Errorf("empty input = (%v, %v), want empty set", got, err)
	}
}

func TestParseEventSigs(t *testing.T) {
	got, err := ParseEventSigs(topicX.Hex() + "," + strings.TrimPrefix(topicY.Hex(), "0x"))
	if err != nil {
		t.Fatalf("ParseEventSigs: %v", err)
	}
	if len(got) != 2 || !got[topicX] || !got[topicY] {
		t.Errorf("parsed = %v", got)
	}
	if _, err := ParseEventSigs("0xdead"); err == nil {
		t.Error("short signature accepted")
	}
	if got, err := ParseEventSigs(" , "); err != nil || len(got) != 0 {
		t.Errorf("blank input = (%v, %v), want empty set", got, err)
	}
}

func TestLastScannedBlock(t *testing.T) {
	dir := t.TempDir()

	t.Run("missing file", func(t *testing.T) {
		if _, found, _ := lastScannedBlock(filepath.Join(dir, "nope.jsonl")); found {
			t.Error("missing file reported found")
		}
	})
	t.Run("empty file", func(t *testing.T) {
		path := filepath.Join(dir, "empty.jsonl")
		if err := os.WriteFile(path, nil, 0o600); err != nil {
			t.Fatal(err)
		}
		if _, found, _ := lastScannedBlock(path); found {
			t.Error("empty file reported found")
		}
	})
	t.Run("small file", func(t *testing.T) {
		path := filepath.Join(dir, "small.jsonl")
		if err := os.WriteFile(path, []byte(`{"blockNumber":7}`+"\n"+`{"blockNumber":9}`+"\n"), 0o600); err != nil {
			t.Fatal(err)
		}
		block, found, warn := lastScannedBlock(path)
		if !found || block != 9 || warn != nil {
			t.Errorf("= (%d, %v, %v), want (9, true, nil)", block, found, warn)
		}
	})
	t.Run("large file uses tail", func(t *testing.T) {
		path := filepath.Join(dir, "large.jsonl")
		f, err := os.Create(path) //nolint:gosec // test path under t.TempDir
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 40_000; i++ { // ~1.2MB, past the 1MB whole-file limit
			fmt.Fprintf(f, `{"blockNumber":%d,"pad":"%s"}`+"\n", i, strings.Repeat("x", 8))
		}
		_ = f.Close()
		block, found, warn := lastScannedBlock(path)
		if !found || block != 39_999 || warn != nil {
			t.Errorf("= (%d, %v, %v), want (39999, true, nil)", block, found, warn)
		}
	})
	t.Run("unparseable last line warns", func(t *testing.T) {
		path := filepath.Join(dir, "bad.jsonl")
		if err := os.WriteFile(path, []byte("not json\n"), 0o600); err != nil {
			t.Fatal(err)
		}
		_, found, warn := lastScannedBlock(path)
		if found || warn == nil {
			t.Errorf("= (%v, %v), want not-found with warning", found, warn)
		}
	})
}

func TestDefaultWorkers(t *testing.T) {
	if got := DefaultWorkers(1); got != 16 {
		t.Errorf("DefaultWorkers(1) = %d, want floor 16", got)
	}
	if got := DefaultWorkers(20); got != 40 {
		t.Errorf("DefaultWorkers(20) = %d, want 40", got)
	}
	if got := DefaultWorkers(100); got != 64 {
		t.Errorf("DefaultWorkers(100) = %d, want cap 64", got)
	}
}

func TestPrintInfo(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	var out bytes.Buffer
	PrintInfo(&out, reader)
	text := out.String()
	for _, want := range []string{
		"Freezer Information:",
		"Item count (blocks): 5",
		"Max available block: 4",
		"CDAT files: 1",
		"receipts.0000.cdat",
		"First 10 index entries:",
		"Block 0: offset 0",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("PrintInfo output missing %q\noutput:\n%s", want, text)
		}
	}
}

// --- additional branch coverage ---

func TestScanDefaultWorkerCount(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	var out bytes.Buffer
	if _, err := Run(context.Background(), reader, Options{OutPath: "-", Stdout: &out, Workers: 0}); err != nil {
		t.Fatalf("Run with auto workers: %v", err)
	}
	if len(records(t, out.Bytes())) != 4 {
		t.Error("auto-worker scan missed records")
	}
}

func TestScanResumeWarnsOnUnparseableOutput(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	outPath := filepath.Join(t.TempDir(), "events.jsonl")
	if err := os.WriteFile(outPath, []byte("garbage last line\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	var logLines []string
	if _, err := Run(context.Background(), reader, Options{
		OutPath: outPath, Workers: 1,
		Logf: func(format string, args ...any) { logLines = append(logLines, fmt.Sprintf(format, args...)) },
	}); err != nil {
		t.Fatalf("Run: %v", err)
	}
	found := false
	for _, l := range logLines {
		if strings.Contains(l, "Warning: could not parse last line") {
			found = true
		}
	}
	if !found {
		t.Errorf("resume warning missing from log: %v", logLines)
	}
	// The unparseable file is treated as no-resume: a fresh scan replaces it.
	if recs := records(t, readOutput(t, outPath)); len(recs) != 4 {
		t.Errorf("records = %d, want a full fresh scan", len(recs))
	}
}

func TestScanErrorLogOpenFailureAborts(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	var out bytes.Buffer
	_, err := Run(context.Background(), reader, Options{
		OutPath: "-", Stdout: &out, Workers: 1,
		ErrorLogPath: filepath.Join(t.TempDir(), "no", "dir", "err.log"),
	})
	if err == nil || !strings.Contains(err.Error(), "failed to create error log") {
		t.Errorf("Run = %v, want error-log failure", err)
	}
}

func TestScanTempDirCreationFailure(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	t.Setenv("TMPDIR", filepath.Join(t.TempDir(), "does", "not", "exist"))
	var out bytes.Buffer
	_, err := Run(context.Background(), reader, Options{OutPath: "-", Stdout: &out, Workers: 1})
	if err == nil || !strings.Contains(err.Error(), "failed to create temp dir") {
		t.Errorf("Run = %v, want temp-dir failure", err)
	}
}

func TestScanMergeFailureOnUnwritableOutput(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	_, err := Run(context.Background(), reader, Options{
		OutPath: filepath.Join(t.TempDir(), "missing-dir", "events.jsonl"),
		Workers: 1,
	})
	if err == nil || !strings.Contains(err.Error(), "merge failed") {
		t.Errorf("Run = %v, want merge failure", err)
	}
}

// failingWriter errors on every write.
type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) { return 0, errors.New("stdout gone") }

func TestScanStdoutWriteFailure(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	_, err := Run(context.Background(), reader, Options{OutPath: "-", Stdout: failingWriter{}, Workers: 1})
	if err == nil || !strings.Contains(err.Error(), "merge failed") {
		t.Errorf("Run = %v, want merge failure on stdout writes", err)
	}
}

// buildRegressingFreezer writes an index whose block-0 offset exceeds the
// block-1 offset — the offset-regression corruption ReadItem detects.
func buildRegressingFreezer(t *testing.T) *freezerscanner.ParallelReader {
	t.Helper()
	dir := t.TempDir()
	offsets := []uint64{50, 10, 60}
	var index []byte
	for _, off := range offsets {
		index = append(index,
			byte(off>>40), byte(off>>32), byte(off>>24), byte(off>>16), byte(off>>8), byte(off))
	}
	if err := os.WriteFile(filepath.Join(dir, "receipts.cidx"), index, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "receipts.0000.cdat"), make([]byte, 60), 0o600); err != nil {
		t.Fatal(err)
	}
	reader, err := freezerscanner.NewParallelReader(dir)
	if err != nil {
		t.Fatalf("NewParallelReader: %v", err)
	}
	return reader
}

func TestScanReadErrorFailFastAndBestEffort(t *testing.T) {
	t.Run("fail fast", func(t *testing.T) {
		reader := buildRegressingFreezer(t)
		var out bytes.Buffer
		_, err := Run(context.Background(), reader, Options{OutPath: "-", Stdout: &out, Workers: 1})
		if err == nil || !strings.Contains(err.Error(), "block 0 read") {
			t.Errorf("Run = %v, want block-0 read failure", err)
		}
	})
	t.Run("best effort logs and continues", func(t *testing.T) {
		reader := buildRegressingFreezer(t)
		errLogPath := filepath.Join(t.TempDir(), "errors.log")
		var out bytes.Buffer
		stats, err := Run(context.Background(), reader, Options{
			OutPath: "-", Stdout: &out, Workers: 1, BestEffort: true, ErrorLogPath: errLogPath,
		})
		if err != nil {
			t.Fatalf("Run: %v", err)
		}
		if stats.Errors == 0 {
			t.Error("read errors not counted")
		}
		logData := readOutput(t, errLogPath)
		if !strings.Contains(string(logData), "type=read") {
			t.Errorf("error log = %q, want a read entry", logData)
		}
	})
}

func TestScanMidRunCancellation(t *testing.T) {
	// Enough work that cancellation lands mid-scan: many blocks with real
	// receipts, one worker, small chunks.
	var blocks [][]byte
	payload := receiptsBlob(t, []*types.Log{simpleLog(contractA, topicX, 1)})
	for i := 0; i < 3000; i++ {
		blocks = append(blocks, payload)
	}
	reader := buildFreezer(t, blocks)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(3 * time.Millisecond)
		cancel()
	}()
	var out bytes.Buffer
	stats, err := Run(ctx, reader, Options{OutPath: "-", Stdout: &out, Workers: 1, ChunkSize: 50})
	if err != nil {
		t.Fatalf("Run after mid-scan cancel: %v", err)
	}
	// Whatever completed before the cancel is merged; no error either way.
	t.Logf("processed %d blocks before cancellation", stats.ProcessedBlocks)
}

func TestLastScannedBlockBlankLine(t *testing.T) {
	path := filepath.Join(t.TempDir(), "blank.jsonl")
	if err := os.WriteFile(path, []byte("\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, found, warn := lastScannedBlock(path); found || warn != nil {
		t.Errorf("blank-line file = (%v, %v), want silent not-found", found, warn)
	}
}

func TestPrintInfoTailEntriesForLargeIndex(t *testing.T) {
	blocks := make([][]byte, 12)
	blocks[11] = receiptsBlob(t, []*types.Log{simpleLog(contractA, topicX, 1)})
	reader := buildFreezer(t, blocks)
	var out bytes.Buffer
	PrintInfo(&out, reader)
	if !strings.Contains(out.String(), "Last 5 index entries:") {
		t.Errorf("PrintInfo output missing tail entries:\n%s", out.String())
	}
}

func TestErrorLogNilIsSafe(t *testing.T) {
	var el *ErrorLog
	el.Log(1, "read", "boom") // must not panic
	if err := el.Close(); err != nil {
		t.Errorf("nil Close = %v", err)
	}
}

func TestErrorLogAppendsAcrossSessions(t *testing.T) {
	path := filepath.Join(t.TempDir(), "err.log")
	for i := 0; i < 2; i++ {
		el, err := NewErrorLog(path)
		if err != nil {
			t.Fatalf("NewErrorLog: %v", err)
		}
		el.Log(uint64(i), "read", "x")
		if err := el.Close(); err != nil {
			t.Fatalf("Close: %v", err)
		}
	}
	data, err := os.ReadFile(path) //nolint:gosec // test path under t.TempDir
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if got := strings.Count(string(data), "Scan session started"); got != 2 {
		t.Errorf("session headers = %d, want 2 (append mode)", got)
	}
}

func TestErrorLogOpenFailure(t *testing.T) {
	if _, err := NewErrorLog(filepath.Join(t.TempDir(), "no", "dir", "err.log")); err == nil {
		t.Error("NewErrorLog into missing directory succeeded")
	}
}

func TestScanProgressReporting(t *testing.T) {
	reader := buildFreezer(t, standardBlocks(t))
	var mu sync.Mutex
	var lines []string
	var out bytes.Buffer
	_, err := Run(context.Background(), reader, Options{
		OutPath: "-", Stdout: &out, Workers: 1,
		ProgressEvery: time.Nanosecond, // fire as often as possible
		Logf: func(format string, args ...any) {
			mu.Lock()
			defer mu.Unlock()
			lines = append(lines, fmt.Sprintf(format, args...))
		},
	})
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	mu.Lock()
	defer mu.Unlock()
	joined := strings.Join(lines, "\n")
	for _, want := range []string{"Parallel scan: blocks 0 to 4", "Chunks:", "Merging results...", "Scan complete:"} {
		if !strings.Contains(joined, want) {
			t.Errorf("log lines missing %q:\n%s", want, joined)
		}
	}
}
