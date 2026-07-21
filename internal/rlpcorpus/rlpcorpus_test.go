package rlpcorpus

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"math"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

var errCorpusWrite = errors.New("write failed")

type failingCorpusWriter struct{}

func (failingCorpusWriter) Write([]byte) (int, error) { return 0, errCorpusWrite }

func validEntry(tb testing.TB, project string, index int) Entry {
	tb.Helper()
	address := common.HexToAddress("0x2000000000000000000000000000000000000002")
	topic := common.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	lg := &types.Log{
		Address: address,
		Topics:  []common.Hash{topic, common.BigToHash(big.NewInt(int64(index + 1)))},
		Data:    []byte{0x01, 0x02, byte(index)}, // #nosec G115 -- test indexes are 0 and 1
	}
	raw, err := toolutil.EncodeLogRLP(lg)
	if err != nil {
		tb.Fatalf("EncodeLogRLP: %v", err)
	}
	return Entry{
		Project:         project,
		BlockNum:        int64(100 + index),
		EventID:         int64(index + 1),
		LogIndex:        index,
		TxHash:          common.BigToHash(big.NewInt(int64(index + 10))).Hex(),
		ContractAddress: address.Hex(),
		Topic0Sig:       toolutil.Topic0Sig(lg),
		LogRLP:          "0x" + hex.EncodeToString(raw),
	}
}

func TestRoundTrip(t *testing.T) {
	entries := []Entry{
		validEntry(t, "cosmicgame", 0),
		validEntry(t, "randomwalk", 1),
	}
	var out bytes.Buffer
	if err := Write(&out, entries); err != nil {
		t.Fatalf("Write: %v", err)
	}
	if strings.Count(out.String(), "\n") != len(entries) {
		t.Fatalf("JSONL = %q, want one line per entry", out.String())
	}

	loaded, err := Load(strings.NewReader("\n" + out.String() + "\n"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !reflect.DeepEqual(loaded, entries) {
		t.Fatalf("loaded entries = %+v, want %+v", loaded, entries)
	}
	for i := range entries {
		lg, err := loaded[i].DecodedLog()
		if err != nil {
			t.Fatalf("DecodedLog(%d): %v", i, err)
		}
		if lg.Address != common.HexToAddress(entries[i].ContractAddress) ||
			toolutil.Topic0Sig(&lg) != entries[i].Topic0Sig {
			t.Errorf("decoded log %d = %+v", i, lg)
		}
	}
}

func TestLoadRejectsInvalidEntries(t *testing.T) {
	t.Parallel()
	base := validEntry(t, "cosmicgame", 0)
	tests := []struct {
		name   string
		mutate func(*Entry)
		want   string
	}{
		{"project", func(e *Entry) { e.Project = "other" }, "project"},
		{"negative block", func(e *Entry) { e.BlockNum = -1 }, "blockNum"},
		{"negative event", func(e *Entry) { e.EventID = -1 }, "eventId"},
		{"negative index", func(e *Entry) { e.LogIndex = -1 }, "logIndex"},
		{"large index", func(e *Entry) { e.LogIndex = math.MaxInt32 + 1 }, "integer range"},
		{"short tx hash", func(e *Entry) { e.TxHash = "0x01" }, "txHash"},
		{"tx hash prefix", func(e *Entry) { e.TxHash = strings.TrimPrefix(e.TxHash, "0x") }, "0x-prefixed"},
		{"contract", func(e *Entry) { e.ContractAddress = "0x1234" }, "contractAddress"},
		{"topic length", func(e *Entry) { e.Topic0Sig = "0x01" }, "topic0Sig"},
		{"topic mismatch", func(e *Entry) { e.Topic0Sig = "bbbbbbbb" }, "disagrees"},
		{"empty RLP", func(e *Entry) { e.LogRLP = "0x" }, "empty"},
		{"bad RLP hex", func(e *Entry) { e.LogRLP = "0xzz" }, "invalid hex"},
		{"bad RLP", func(e *Entry) { e.LogRLP = "0xff" }, "types.Log RLP"},
		{"address mismatch", func(e *Entry) {
			e.ContractAddress = "0x3000000000000000000000000000000000000003"
		}, "disagrees"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			entry := base
			test.mutate(&entry)
			var data bytes.Buffer
			if err := json.NewEncoder(&data).Encode(entry); err != nil {
				t.Fatal(err)
			}
			_, err := Load(&data)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Load error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestLoadRejectsMalformedJSONAndDuplicates(t *testing.T) {
	t.Parallel()
	entry := validEntry(t, "cosmicgame", 0)
	var valid bytes.Buffer
	if err := json.NewEncoder(&valid).Encode(entry); err != nil {
		t.Fatal(err)
	}

	for name, input := range map[string]string{
		"unknown field": strings.TrimSuffix(valid.String(), "}\n") + `,"extra":true}` + "\n",
		"multiple JSON": strings.TrimSpace(valid.String()) + ` {}` + "\n",
		"syntax":        `{"project":`,
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := Load(strings.NewReader(input)); err == nil {
				t.Fatal("malformed corpus was accepted")
			}
		})
	}

	if _, err := Load(strings.NewReader(valid.String() + valid.String())); err == nil ||
		!strings.Contains(err.Error(), "duplicate chain identity") {
		t.Fatalf("duplicate error = %v", err)
	}

	sibling := validEntry(t, "cosmicgame", 1)
	sibling.BlockNum = entry.BlockNum
	sibling.TxHash = entry.TxHash
	sibling.LogIndex = 1
	foreign := validEntry(t, "randomwalk", 1)
	if _, err := Load(bytes.NewReader(entriesJSON(t, entry, foreign, sibling))); err == nil ||
		!strings.Contains(err.Error(), "not contiguous") {
		t.Fatalf("non-contiguous transaction error = %v", err)
	}

	later := entry
	later.LogIndex = 2
	if _, err := Load(bytes.NewReader(entriesJSON(t, later, sibling))); err == nil ||
		!strings.Contains(err.Error(), "not strictly after") {
		t.Fatalf("descending log-index error = %v", err)
	}
}

func TestIOFailuresAndBounds(t *testing.T) {
	if _, err := Load(nil); err == nil {
		t.Fatal("nil reader accepted")
	}
	if err := Write(nil, nil); err == nil {
		t.Fatal("nil writer accepted")
	}
	if err := Write(io.Discard, []Entry{{}}); err == nil {
		t.Fatal("invalid entry accepted by writer")
	}
	duplicate := validEntry(t, "cosmicgame", 0)
	if err := Write(io.Discard, []Entry{duplicate, duplicate}); err == nil ||
		!strings.Contains(err.Error(), "duplicate") {
		t.Fatalf("duplicate writer error = %v", err)
	}
	if err := Write(failingCorpusWriter{}, []Entry{duplicate}); !errors.Is(err, errCorpusWrite) {
		t.Fatalf("write error = %v, want %v", err, errCorpusWrite)
	}

	oversized := strings.Repeat("x", maxLine+1)
	if _, err := Load(strings.NewReader(oversized)); err == nil ||
		!strings.Contains(err.Error(), "token too long") {
		t.Fatalf("oversized line error = %v", err)
	}
}

func FuzzLoad(f *testing.F) {
	valid := validEntry(f, "cosmicgame", 0)
	var seed bytes.Buffer
	if err := Write(&seed, []Entry{valid}); err != nil {
		f.Fatalf("building seed corpus: %v", err)
	}
	f.Add(seed.Bytes())
	sibling := validEntry(f, "cosmicgame", 1)
	sibling.BlockNum = valid.BlockNum
	sibling.TxHash = valid.TxHash
	sibling.LogIndex = 1
	var multiLogSeed bytes.Buffer
	if err := Write(&multiLogSeed, []Entry{valid, sibling}); err != nil {
		f.Fatalf("building multi-log seed corpus: %v", err)
	}
	f.Add(multiLogSeed.Bytes())
	f.Add([]byte{})
	f.Add([]byte(`{"project":"cosmicgame","logRlp":"0xff"}`))

	f.Fuzz(func(t *testing.T, data []byte) {
		first, firstErr := Load(bytes.NewReader(data))
		second, secondErr := Load(bytes.NewReader(data))
		if !reflect.DeepEqual(first, second) {
			t.Fatalf("non-deterministic entries: first=%+v second=%+v", first, second)
		}
		if (firstErr == nil) != (secondErr == nil) {
			t.Fatalf("non-deterministic errors: first=%v second=%v", firstErr, secondErr)
		}
		if firstErr != nil {
			if firstErr.Error() != secondErr.Error() {
				t.Fatalf("non-deterministic error text: first=%q second=%q", firstErr, secondErr)
			}
			return
		}
		var roundTrip bytes.Buffer
		if err := Write(&roundTrip, first); err != nil {
			t.Fatalf("valid corpus failed to write: %v", err)
		}
		reloaded, err := Load(&roundTrip)
		if err != nil {
			t.Fatalf("written corpus failed to reload: %v", err)
		}
		if !reflect.DeepEqual(reloaded, first) {
			t.Fatalf("round trip changed entries: got=%+v want=%+v", reloaded, first)
		}
	})
}

func entriesJSON(t *testing.T, entries ...Entry) []byte {
	t.Helper()
	var out bytes.Buffer
	enc := json.NewEncoder(&out)
	for _, entry := range entries {
		if err := enc.Encode(entry); err != nil {
			t.Fatal(err)
		}
	}
	return out.Bytes()
}
