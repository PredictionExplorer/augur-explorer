//go:build integration

package randomwalk

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

const (
	randomWalkCorpusOfferTx = "0x0000000000000000000000000000000000000000000000000000000000002331"
	randomWalkCorpusSaleTx  = "0x0000000000000000000000000000000000000000000000000000000000002332"
)

// TestRLPCorpusReplaysThroughRealHandlers proves the committed JSONL sample
// survives archive-format parsing, is installed with its exact exported RLP
// bytes, and dispatches through the production RandomWalk registry. The mint
// triggers must create the token and update global statistics.
func TestRLPCorpusReplaysThroughRealHandlers(t *testing.T) {
	resetDB(t)
	entries := loadRandomWalkCorpus(t)

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
	requireRandomWalkCorpusState(t)
}

// TestRLPCorpusIngestsThroughEngineRun proves the same archive-shaped payloads
// survive the production RPC -> InsertEventLog -> registry path. Canonical RLP
// bytes and all mint/marketplace/transfer outcomes must match direct replay.
func TestRLPCorpusIngestsThroughEngineRun(t *testing.T) {
	resetDB(t)
	entries := loadRandomWalkCorpus(t)
	firstBlock := entries[0].BlockNum
	lastBlock := entries[len(entries)-1].BlockNum
	testChain.Reorg(firstBlock) // make -count=N runs reuse-safe
	stageRandomWalkCorpus(t, entries)

	batchSize := uint64(lastBlock - firstBlock + 1) // #nosec G115 -- validated corpus block numbers are non-negative and ordered
	if err := runHarnessRange(t, firstBlock, lastBlock, batchSize); err != nil {
		t.Fatalf("Engine.Run corpus ingest: %v", err)
	}
	requireRandomWalkCorpusRLP(t, entries)
	requireRandomWalkCorpusState(t)
}

func loadRandomWalkCorpus(t *testing.T) []rlpcorpus.Entry {
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
		if entry.Project != "randomwalk" {
			t.Fatalf("entry %d project = %q, want randomwalk", i, entry.Project)
		}
	}
	var canonical bytes.Buffer
	if err := rlpcorpus.Write(&canonical, entries); err != nil {
		t.Fatalf("rlpcorpus.Write: %v", err)
	}
	if !bytes.Equal(corpusBytes, canonical.Bytes()) {
		t.Fatal("committed corpus is not in canonical JSONL form")
	}

	offerEntries := randomWalkCorpusEntriesForTx(entries, randomWalkCorpusOfferTx)
	requireRandomWalkFixtureLogs(
		t,
		offerEntries,
		[]*types.Log{offerLog(1, 99, fxBob, fxZero, 3)(t)},
		"offer",
	)
	saleEntries := randomWalkCorpusEntriesForTx(entries, randomWalkCorpusSaleTx)
	requireRandomWalkFixtureLogs(t, saleEntries, []*types.Log{
		buildLog(t, fxMarketABI, "ItemBought", addr(fxMarketplaceAddr),
			[]any{bigInt(1), addr(fxBob), addr(fxDave)}, nil),
		transferLog(fxBob, fxDave, 99)(t),
	}, "sale")
	return entries
}

func randomWalkCorpusEntriesForTx(entries []rlpcorpus.Entry, txHash string) []rlpcorpus.Entry {
	result := make([]rlpcorpus.Entry, 0)
	for _, entry := range entries {
		if entry.TxHash == txHash {
			result = append(result, entry)
		}
	}
	return result
}

func requireRandomWalkFixtureLogs(
	t *testing.T,
	entries []rlpcorpus.Entry,
	expected []*types.Log,
	label string,
) {
	t.Helper()
	if len(entries) != len(expected) {
		t.Fatalf("%s corpus has %d sibling logs, want %d", label, len(entries), len(expected))
	}
	for i, entry := range entries {
		got, err := entry.DecodedLog()
		if err != nil {
			t.Fatalf("decode %s corpus log %d: %v", label, i, err)
		}
		want := expected[i] // #nosec G602 -- equal lengths are required above
		if !sameCorpusLog(got, *want) {
			t.Fatalf(
				"%s corpus log %d drifted from fixture:\ngot  %+v\nwant %+v",
				label,
				i,
				got,
				*want,
			)
		}
	}
}

func sameCorpusLog(a, b types.Log) bool {
	return a.Address == b.Address &&
		slices.Equal(a.Topics, b.Topics) &&
		bytes.Equal(a.Data, b.Data)
}

func stageRandomWalkCorpus(t *testing.T, entries []rlpcorpus.Entry) {
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

func requireRandomWalkCorpusRLP(t *testing.T, entries []rlpcorpus.Entry) {
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

func requireRandomWalkCorpusState(t *testing.T) {
	t.Helper()
	var (
		mintRows     int
		tokenRows    int
		offerRows    int
		purchaseRows int
		transferRows int
		total        int64
		owner        string
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
	for query, destination := range map[string]*int{
		"SELECT COUNT(*) FROM rw_new_offer":   &offerRows,
		"SELECT COUNT(*) FROM rw_item_bought": &purchaseRows,
		"SELECT COUNT(*) FROM rw_transfer":    &transferRows,
	} {
		if err := testDB.Pool.QueryRow(context.Background(), query).Scan(destination); err != nil {
			t.Fatalf("read corpus outcome %q: %v", query, err)
		}
	}
	if err := testDB.Pool.QueryRow(context.Background(), `
		SELECT BTRIM(a.addr)
		FROM rw_token t
		JOIN address a ON a.address_id=t.cur_owner_aid
		WHERE t.token_id=99
	`).Scan(&owner); err != nil {
		t.Fatalf("read corpus token owner: %v", err)
	}
	if mintRows != 1 || tokenRows != 1 || offerRows != 1 || purchaseRows != 1 ||
		transferRows != 1 || total != 1 || owner != fxDave {
		t.Fatalf(
			"corpus result: mint=%d token=%d offer=%d purchase=%d transfer=%d total=%d owner=%s",
			mintRows,
			tokenRows,
			offerRows,
			purchaseRows,
			transferRows,
			total,
			owner,
		)
	}
}
