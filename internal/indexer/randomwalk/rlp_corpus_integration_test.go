//go:build integration

package randomwalk

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

// TestRLPCorpusReplaysThroughRealHandlers proves the committed JSONL sample
// survives archive-format parsing, is installed with its exact exported RLP
// bytes, and dispatches through the production RandomWalk registry. The mint
// triggers must create the token and update global statistics.
func TestRLPCorpusReplaysThroughRealHandlers(t *testing.T) {
	resetDB(t)
	corpusBytes, err := os.ReadFile("testdata/rlp_corpus.jsonl")
	if err != nil {
		t.Fatalf("read corpus: %v", err)
	}
	entries, err := testutil.LoadRLPCorpus(bytes.NewReader(corpusBytes))
	if err != nil {
		t.Fatalf("LoadRLPCorpus: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("committed corpus is empty")
	}
	for i, entry := range entries {
		if entry.Project != "randomwalk" {
			t.Fatalf("entry %d project = %q, want randomwalk", i, entry.Project)
		}
	}
	var canonical bytes.Buffer
	if err := testutil.WriteRLPCorpus(&canonical, entries); err != nil {
		t.Fatalf("WriteRLPCorpus: %v", err)
	}
	if !bytes.Equal(corpusBytes, canonical.Bytes()) {
		t.Fatal("committed corpus is not in canonical JSONL form")
	}

	ids, err := testutil.InstallRLPCorpus(context.Background(), dbStore, entries)
	if err != nil {
		t.Fatalf("InstallRLPCorpus: %v", err)
	}
	if len(ids) != len(entries) {
		t.Fatalf("installed %d events, want %d", len(ids), len(entries))
	}
	for i, id := range ids {
		var stored []byte
		if err := testDB.Pool.QueryRow(
			context.Background(),
			"SELECT log_rlp FROM evt_log WHERE id=$1",
			id,
		).Scan(&stored); err != nil {
			t.Fatalf("read installed RLP %d: %v", id, err)
		}
		want, err := entries[i].RLPBytes()
		if err != nil {
			t.Fatalf("entry %d RLPBytes: %v", i, err)
		}
		if !bytes.Equal(stored, want) {
			t.Fatalf("entry %d stored RLP changed during installation", i)
		}
		if err := testProcess(context.Background(), id); err != nil {
			t.Fatalf("process corpus event %d: %v", id, err)
		}
	}

	var (
		mintRows  int
		tokenRows int
		total     int64
	)
	if err := testDB.Pool.QueryRow(
		context.Background(),
		"SELECT COUNT(*) FROM rw_mint_evt",
	).Scan(&mintRows); err != nil {
		t.Fatalf("count mint rows: %v", err)
	}
	if err := testDB.Pool.QueryRow(
		context.Background(),
		"SELECT COUNT(*) FROM rw_token WHERE token_id=99",
	).Scan(&tokenRows); err != nil {
		t.Fatalf("count token rows: %v", err)
	}
	if err := testDB.Pool.QueryRow(
		context.Background(),
		"SELECT total_num_toks FROM rw_stats",
	).Scan(&total); err != nil {
		t.Fatalf("read trigger-updated stats: %v", err)
	}
	if mintRows != 1 || tokenRows != 1 || total != 1 {
		t.Fatalf("corpus result: mint rows=%d token rows=%d total=%d", mintRows, tokenRows, total)
	}
}
