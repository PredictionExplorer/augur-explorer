//go:build integration

package cosmicgame

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

// TestRLPCorpusReplaysThroughRealHandlers proves the committed JSONL sample
// survives archive-format parsing, is installed with its exact exported RLP
// bytes, and dispatches through the production CosmicGame registry. The
// admin-event trigger must update cg_glob_stats as it does in live ingestion.
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
		if entry.Project != "cosmicgame" {
			t.Fatalf("entry %d project = %q, want cosmicgame", i, entry.Project)
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
		adminRows int
		reward    string
	)
	if err := testDB.Pool.QueryRow(
		context.Background(),
		"SELECT COUNT(*) FROM cg_adm_erc20_reward",
	).Scan(&adminRows); err != nil {
		t.Fatalf("count admin rows: %v", err)
	}
	if err := testDB.Pool.QueryRow(
		context.Background(),
		"SELECT cst_reward_for_bidding::text FROM cg_glob_stats",
	).Scan(&reward); err != nil {
		t.Fatalf("read trigger-updated reward: %v", err)
	}
	if adminRows != 1 || reward != eth(123).String() {
		t.Fatalf("corpus result: admin rows=%d reward=%s", adminRows, reward)
	}
}
