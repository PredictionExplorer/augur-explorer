package testutil

import (
	"bytes"
	"context"
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

func validCorpusEntry(tb testing.TB, project string, index int) RLPCorpusEntry {
	tb.Helper()
	address := common.HexToAddress("0x2000000000000000000000000000000000000002")
	topic := common.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	lg := &types.Log{
		Address: address,
		Topics:  []common.Hash{topic, common.BigToHash(newBigInt(int64(index + 1)))},
		Data:    []byte{0x01, 0x02, byte(index)}, // #nosec G115 -- test indexes are 0 and 1
	}
	raw, err := toolutil.EncodeLogRLP(lg)
	if err != nil {
		tb.Fatalf("EncodeLogRLP: %v", err)
	}
	return RLPCorpusEntry{
		Project:         project,
		BlockNum:        int64(100 + index),
		EventID:         int64(index + 1),
		LogIndex:        index,
		TxHash:          common.BigToHash(newBigInt(int64(index + 10))).Hex(),
		ContractAddress: address.Hex(),
		Topic0Sig:       toolutil.Topic0Sig(lg),
		LogRLP:          "0x" + hex.EncodeToString(raw),
	}
}

func newBigInt(value int64) *big.Int {
	return big.NewInt(value)
}

func TestRLPCorpusRoundTrip(t *testing.T) {
	entries := []RLPCorpusEntry{
		validCorpusEntry(t, "cosmicgame", 0),
		validCorpusEntry(t, "randomwalk", 1),
	}
	var out bytes.Buffer
	if err := WriteRLPCorpus(&out, entries); err != nil {
		t.Fatalf("WriteRLPCorpus: %v", err)
	}
	if strings.Count(out.String(), "\n") != len(entries) {
		t.Fatalf("JSONL = %q, want one line per entry", out.String())
	}

	loaded, err := LoadRLPCorpus(strings.NewReader("\n" + out.String() + "\n"))
	if err != nil {
		t.Fatalf("LoadRLPCorpus: %v", err)
	}
	if len(loaded) != len(entries) {
		t.Fatalf("loaded %d entries, want %d", len(loaded), len(entries))
	}
	for i := range entries {
		if loaded[i] != entries[i] {
			t.Errorf("entry %d = %+v, want %+v", i, loaded[i], entries[i])
		}
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

func TestLoadRLPCorpusRejectsInvalidEntries(t *testing.T) {
	base := validCorpusEntry(t, "cosmicgame", 0)
	tests := []struct {
		name   string
		mutate func(*RLPCorpusEntry)
		want   string
	}{
		{"project", func(e *RLPCorpusEntry) { e.Project = "other" }, "project"},
		{"negative block", func(e *RLPCorpusEntry) { e.BlockNum = -1 }, "blockNum"},
		{"negative event", func(e *RLPCorpusEntry) { e.EventID = -1 }, "eventId"},
		{"negative index", func(e *RLPCorpusEntry) { e.LogIndex = -1 }, "logIndex"},
		{"large index", func(e *RLPCorpusEntry) { e.LogIndex = math.MaxInt32 + 1 }, "integer range"},
		{"short tx hash", func(e *RLPCorpusEntry) { e.TxHash = "0x01" }, "txHash"},
		{"tx hash prefix", func(e *RLPCorpusEntry) { e.TxHash = strings.TrimPrefix(e.TxHash, "0x") }, "0x-prefixed"},
		{"contract", func(e *RLPCorpusEntry) { e.ContractAddress = "0x1234" }, "contractAddress"},
		{"topic length", func(e *RLPCorpusEntry) { e.Topic0Sig = "0x01" }, "topic0Sig"},
		{"topic mismatch", func(e *RLPCorpusEntry) { e.Topic0Sig = "bbbbbbbb" }, "disagrees"},
		{"empty RLP", func(e *RLPCorpusEntry) { e.LogRLP = "0x" }, "empty"},
		{"bad RLP hex", func(e *RLPCorpusEntry) { e.LogRLP = "0xzz" }, "invalid hex"},
		{"bad RLP", func(e *RLPCorpusEntry) { e.LogRLP = "0xff" }, "types.Log RLP"},
		{"address mismatch", func(e *RLPCorpusEntry) {
			e.ContractAddress = "0x3000000000000000000000000000000000000003"
		}, "disagrees"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			entry := base
			test.mutate(&entry)
			var data bytes.Buffer
			enc := json.NewEncoder(&data)
			if err := enc.Encode(entry); err != nil {
				t.Fatal(err)
			}
			_, err := LoadRLPCorpus(&data)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("LoadRLPCorpus error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestLoadRLPCorpusRejectsMalformedJSONAndDuplicates(t *testing.T) {
	entry := validCorpusEntry(t, "cosmicgame", 0)
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
			if _, err := LoadRLPCorpus(strings.NewReader(input)); err == nil {
				t.Fatal("malformed corpus was accepted")
			}
		})
	}

	if _, err := LoadRLPCorpus(strings.NewReader(valid.String() + valid.String())); err == nil ||
		!strings.Contains(err.Error(), "duplicate chain identity") {
		t.Fatalf("duplicate error = %v", err)
	}
}

func TestRLPCorpusIOFailuresAndBounds(t *testing.T) {
	if _, err := LoadRLPCorpus(nil); err == nil {
		t.Fatal("nil reader accepted")
	}
	if err := WriteRLPCorpus(nil, nil); err == nil {
		t.Fatal("nil writer accepted")
	}
	if _, err := InstallRLPCorpus(context.Background(), nil, nil); err == nil {
		t.Fatal("nil Store accepted")
	}
	if err := WriteRLPCorpus(io.Discard, []RLPCorpusEntry{{}}); err == nil {
		t.Fatal("invalid entry accepted by writer")
	}
	duplicate := validCorpusEntry(t, "cosmicgame", 0)
	if err := WriteRLPCorpus(io.Discard, []RLPCorpusEntry{duplicate, duplicate}); err == nil ||
		!strings.Contains(err.Error(), "duplicate") {
		t.Fatalf("duplicate writer error = %v", err)
	}
	if err := WriteRLPCorpus(failingCorpusWriter{}, []RLPCorpusEntry{duplicate}); !errors.Is(err, errCorpusWrite) {
		t.Fatalf("write error = %v, want %v", err, errCorpusWrite)
	}

	oversized := strings.Repeat("x", maxRLPCorpusLine+1)
	if _, err := LoadRLPCorpus(strings.NewReader(oversized)); err == nil ||
		!strings.Contains(err.Error(), "token too long") {
		t.Fatalf("oversized line error = %v", err)
	}
}

func FuzzLoadRLPCorpus(f *testing.F) {
	valid := validCorpusEntry(f, "cosmicgame", 0)
	var seed bytes.Buffer
	if err := WriteRLPCorpus(&seed, []RLPCorpusEntry{valid}); err != nil {
		f.Fatalf("building seed corpus: %v", err)
	}
	f.Add(seed.Bytes())
	f.Add([]byte{})
	f.Add([]byte(`{"project":"cosmicgame","logRlp":"0xff"}`))

	f.Fuzz(func(t *testing.T, data []byte) {
		first, firstErr := LoadRLPCorpus(bytes.NewReader(data))
		second, secondErr := LoadRLPCorpus(bytes.NewReader(data))
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
		if err := WriteRLPCorpus(&roundTrip, first); err != nil {
			t.Fatalf("valid corpus failed to write: %v", err)
		}
		reloaded, err := LoadRLPCorpus(&roundTrip)
		if err != nil {
			t.Fatalf("written corpus failed to reload: %v", err)
		}
		if !reflect.DeepEqual(reloaded, first) {
			t.Fatalf("round trip changed entries: got=%+v want=%+v", reloaded, first)
		}
	})
}
