//go:build integration

package archive_test

import (
	"bytes"
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/archive"
	"github.com/PredictionExplorer/augur-explorer/internal/rlpcorpus"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

func TestCorpusExportReadsCompleteArchiveTransaction(t *testing.T) {
	db := testdb.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	txHash := "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	contract := common.HexToAddress("0x2000000000000000000000000000000000000002")
	topic := common.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	rawLogs := make([][]byte, 2)
	for index := range rawLogs {
		raw, err := toolutil.EncodeLogRLP(&types.Log{
			Address: contract,
			Topics:  []common.Hash{topic, common.BigToHash(newInt(int64(index + 1)))},
			Data:    []byte{byte(index + 1)}, // #nosec G115 -- two fixed test indexes
		})
		if err != nil {
			t.Fatal(err)
		}
		rawLogs[index] = raw
		if _, err := db.Pool.Exec(ctx, `INSERT INTO arch_evtlog (
				block_num, evt_id, log_index, tx_hash, contract_addr, topic0_sig, log_rlp
			) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			int64(100),
			int64(index+10),
			index,
			txHash,
			contract.Hex(),
			toolutil.Topic0Sig(&types.Log{Topics: []common.Hash{topic}}),
			raw,
		); err != nil {
			t.Fatalf("insert archive log %d: %v", index, err)
		}
	}

	var out bytes.Buffer
	stats, err := archive.ExportCorpus(
		ctx,
		db.Pool,
		archive.CorpusExportOptions{
			Project:  archive.ProjectCosmicGame,
			TxHashes: []string{txHash},
		},
		&out,
	)
	if err != nil {
		t.Fatalf("ExportCorpus: %v", err)
	}
	if stats != (archive.CorpusExportStats{Transactions: 1, EventLogs: 2}) {
		t.Fatalf("stats = %+v", stats)
	}

	entries, err := rlpcorpus.Load(&out)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(entries) != 2 {
		t.Fatalf("entries = %d", len(entries))
	}
	for i, entry := range entries {
		raw, err := entry.RLPBytes()
		if err != nil {
			t.Fatalf("entry %d RLP: %v", i, err)
		}
		if entry.LogIndex != i || !bytes.Equal(raw, rawLogs[i]) {
			t.Errorf("entry %d changed: %+v", i, entry)
		}
	}
}

func newInt(value int64) *big.Int {
	return big.NewInt(value)
}
