//go:build integration

package cosmicgame

import (
	"bytes"
	"context"
	"os"
	"slices"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/rlpcorpus"
	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

const cosmicCorpusBidTx = "0x0000000000000000000000000000000000000000000000000000000000002330"

// TestRLPCorpusReplaysThroughRealHandlers proves the committed JSONL sample
// survives archive-format parsing, is installed with its exact exported RLP
// bytes, and dispatches through the production CosmicGame registry. The
// admin-event trigger must update cg_glob_stats as it does in live ingestion.
func TestRLPCorpusReplaysThroughRealHandlers(t *testing.T) {
	resetDB(t)
	entries := loadCosmicCorpus(t)

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
	requireCosmicCorpusState(t)
}

// TestRLPCorpusIngestsThroughEngineRun proves the same archive-shaped payloads
// survive the production RPC -> InsertEventLog -> registry path. Canonical RLP
// bytes and all handler/trigger outcomes must match direct archive replay.
func TestRLPCorpusIngestsThroughEngineRun(t *testing.T) {
	resetDB(t)
	entries := loadCosmicCorpus(t)
	firstBlock := entries[0].BlockNum
	lastBlock := entries[len(entries)-1].BlockNum
	testChain.Reorg(firstBlock) // make -count=N runs reuse-safe
	stageCosmicCorpus(t, entries)

	batchSize := uint64(lastBlock - firstBlock + 1) // #nosec G115 -- validated corpus block numbers are non-negative and ordered
	if err := runHarnessRange(t, firstBlock, lastBlock, batchSize); err != nil {
		t.Fatalf("Engine.Run corpus ingest: %v", err)
	}
	requireCosmicCorpusRLP(t, entries)
	requireCosmicCorpusState(t)
}

func loadCosmicCorpus(t *testing.T) []rlpcorpus.Entry {
	t.Helper()
	corpusBytes, err := os.ReadFile("testdata/rlp_corpus.jsonl")
	if err != nil {
		t.Fatalf("read corpus: %v", err)
	}
	entries, err := rlpcorpus.Load(bytes.NewReader(corpusBytes))
	if err != nil {
		t.Fatalf("rlpcorpus.Load: %v", err)
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
	if err := rlpcorpus.Write(&canonical, entries); err != nil {
		t.Fatalf("rlpcorpus.Write: %v", err)
	}
	if !bytes.Equal(corpusBytes, canonical.Bytes()) {
		t.Fatal("committed corpus is not in canonical JSONL form")
	}

	expectedBid := scriptedRound(6000)[1]
	bidEntries := corpusEntriesForTx(entries, cosmicCorpusBidTx)
	if len(bidEntries) != len(expectedBid.logs) {
		t.Fatalf("bid corpus has %d sibling logs, want %d", len(bidEntries), len(expectedBid.logs))
	}
	for i, fixture := range expectedBid.logs {
		got, err := bidEntries[i].DecodedLog()
		if err != nil {
			t.Fatalf("decode bid corpus log %d: %v", i, err)
		}
		if want := fixture.build(t); !sameCorpusLog(got, *want) {
			t.Fatalf("bid corpus log %d drifted from fixture:\ngot  %+v\nwant %+v", i, got, *want)
		}
	}
	return entries
}

func sameCorpusLog(a, b types.Log) bool {
	return a.Address == b.Address &&
		slices.Equal(a.Topics, b.Topics) &&
		bytes.Equal(a.Data, b.Data)
}

func corpusEntriesForTx(entries []rlpcorpus.Entry, txHash string) []rlpcorpus.Entry {
	result := make([]rlpcorpus.Entry, 0)
	for _, entry := range entries {
		if entry.TxHash == txHash {
			result = append(result, entry)
		}
	}
	return result
}

func stageCosmicCorpus(t *testing.T, entries []rlpcorpus.Entry) {
	t.Helper()
	for first := 0; first < len(entries); {
		last := first + 1
		for last < len(entries) && entries[last].TxHash == entries[first].TxHash {
			last++
		}
		logs := make([]*types.Log, 0, last-first)
		for i := first; i < last; i++ {
			if entries[i].BlockNum != entries[first].BlockNum ||
				entries[i].LogIndex != entries[first].LogIndex+i-first {
				t.Fatalf("non-canonical corpus transaction at entry %d", i)
			}
			log, err := entries[i].DecodedLog()
			if err != nil {
				t.Fatalf("decode corpus entry %d: %v", i, err)
			}
			logs = append(logs, &log)
		}
		stageTx(
			t,
			entries[first].BlockNum,
			addr(entries[first].ContractAddress),
			uint(entries[first].LogIndex), // #nosec G115 -- loader bounds non-negative indexes
			logs,
		)
		first = last
	}
}

func requireCosmicCorpusRLP(t *testing.T, entries []rlpcorpus.Entry) {
	t.Helper()
	for i, entry := range entries {
		want, err := entry.RLPBytes()
		if err != nil {
			t.Fatalf("entry %d RLPBytes: %v", i, err)
		}
		var got []byte
		if err := testDB.Pool.QueryRow(
			context.Background(),
			"SELECT log_rlp FROM evt_log WHERE block_num=$1 AND log_index=$2",
			entry.BlockNum,
			entry.LogIndex,
		).Scan(&got); err != nil {
			t.Fatalf("read production-ingested RLP %d: %v", i, err)
		}
		if !bytes.Equal(got, want) {
			t.Fatalf("entry %d RLP changed during production ingestion", i)
		}
	}
}

func requireCosmicCorpusState(t *testing.T) {
	t.Helper()
	var (
		adminRows    int
		firstBidRows int
		bidRows      int
		transferRows int
		reward       string
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
	for query, destination := range map[string]*int{
		"SELECT COUNT(*) FROM cg_first_bid":      &firstBidRows,
		"SELECT COUNT(*) FROM cg_bid":            &bidRows,
		"SELECT COUNT(*) FROM cg_erc20_transfer": &transferRows,
	} {
		if err := testDB.Pool.QueryRow(context.Background(), query).Scan(destination); err != nil {
			t.Fatalf("read corpus outcome %q: %v", query, err)
		}
	}
	if adminRows != 1 || firstBidRows != 1 || bidRows != 1 || transferRows != 1 ||
		reward != eth(123).String() {
		t.Fatalf(
			"corpus result: admin=%d firstBid=%d bid=%d transfer=%d reward=%s",
			adminRows,
			firstBidRows,
			bidRows,
			transferRows,
			reward,
		)
	}
}
