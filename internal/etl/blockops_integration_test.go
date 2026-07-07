//go:build integration

// Integration tests for the shared ETL block operations against a real
// migrated Postgres (testcontainers) and the deterministic fake Ethereum node
// (internal/testchain): block insertion and hash verification, chain-split
// rollback, the transaction three-level fallback (RPC, archive, minimal) and
// event-log deduplication.
package common

import (
	"context"
	"io"
	"log"
	"math/big"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	dbs "github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

// env is one isolated test environment: fresh database, fresh fake chain.
type env struct {
	ctx     *ETLContext
	db      *testdb.DB
	chain   *testchain.Chain
	storage *dbs.SQLStorage
}

func newEnv(t *testing.T) *env {
	t.Helper()
	db := testdb.New(t)
	chain := testchain.New(t)

	rpcClient, err := rpc.DialContext(context.Background(), chain.URL())
	if err != nil {
		t.Fatalf("dialing fake chain: %v", err)
	}
	ethClient := ethclient.NewClient(rpcClient)
	t.Cleanup(ethClient.Close)

	// The address-id cache is package state in internal/store; every fresh
	// database must start with a fresh cache.
	dbs.ResetAddressCacheForTests()

	storage := dbs.NewSQLStorageFromDB(db.SQL, log.New(io.Discard, "", 0))
	storage.Db_set_schema_name("public")

	return &env{
		ctx: &ETLContext{
			Storage:   storage,
			EthClient: ethClient,
			RpcClient: rpcClient,
			Info:      log.New(io.Discard, "", 0),
			Error:     log.New(io.Discard, "", 0),
		},
		db:      db,
		chain:   chain,
		storage: storage,
	}
}

func (e *env) blockHashInDB(t *testing.T, blockNum int64) string {
	t.Helper()
	hash, err := e.storage.Get_block_hash(blockNum)
	if err != nil {
		t.Fatalf("Get_block_hash(%d): %v", blockNum, err)
	}
	return hash
}

func TestEnsureBlockExistsInsertsAndIsIdempotent(t *testing.T) {
	e := newEnv(t)
	hash := e.chain.BlockHash(100).Hex()

	inserted, err := EnsureBlockExists(e.ctx, 100, hash)
	if err != nil {
		t.Fatalf("EnsureBlockExists: %v", err)
	}
	if !inserted {
		t.Error("first call did not insert the block")
	}
	if got := e.blockHashInDB(t, 100); got != hash {
		t.Errorf("stored hash = %s, want %s", got, hash)
	}
	last, err := e.storage.Get_last_block_num()
	if err != nil {
		t.Fatalf("Get_last_block_num: %v", err)
	}
	if last != 100 {
		t.Errorf("last_block watermark = %d, want 100 (requires the migration seeding the last_block row)", last)
	}

	inserted, err = EnsureBlockExists(e.ctx, 100, hash)
	if err != nil {
		t.Fatalf("EnsureBlockExists (second call): %v", err)
	}
	if inserted {
		t.Error("second call re-inserted an existing block")
	}
}

func TestEnsureBlockExistsRejectsUnverifiableHash(t *testing.T) {
	e := newEnv(t)
	e.chain.EnsureBlock(100)

	// The fetched log claims a hash the chain does not confirm: the block
	// must not be inserted.
	bogus := ethcommon.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
	if _, err := EnsureBlockExists(e.ctx, 100, bogus.Hex()); err == nil {
		t.Fatal("expected hash-mismatch error, got nil")
	}
	if _, err := e.storage.Get_block_hash(100); err == nil {
		t.Error("block was inserted despite failing hash verification")
	}
}

func TestChainSplitRollsBackAndCascades(t *testing.T) {
	e := newEnv(t)

	// Build blocks 100..105 with one tx + evt_log in block 103.
	for n := int64(100); n <= 105; n++ {
		if _, err := EnsureBlockExists(e.ctx, n, e.chain.BlockHash(n).Hex()); err != nil {
			t.Fatalf("seeding block %d: %v", n, err)
		}
	}
	to := ethcommon.HexToAddress("0x2000000000000000000000000000000000000002")
	tx := e.chain.AddTx(103, to, nil)
	e.chain.AttachLogs(tx.Hash(), []*types.Log{{
		Address:     to,
		Topics:      []ethcommon.Hash{ethcommon.HexToHash("0x01")},
		BlockNumber: 103,
		TxHash:      tx.Hash(),
		Index:       0,
	}})
	txID, _, err := EnsureTransactionExists(e.ctx, tx.Hash(), 103)
	if err != nil {
		t.Fatalf("EnsureTransactionExists: %v", err)
	}
	evtID, err := InsertEventLog(e.ctx, types.Log{
		Address:     to,
		Topics:      []ethcommon.Hash{ethcommon.HexToHash("0x01")},
		BlockNumber: 103,
		Index:       0,
	}, txID)
	if err != nil {
		t.Fatalf("InsertEventLog: %v", err)
	}

	oldHash := e.chain.BlockHash(102).Hex()
	e.chain.Reorg(102)
	newHash := e.chain.BlockHash(102).Hex()
	if oldHash == newHash {
		t.Fatal("Reorg did not change the block hash")
	}

	inserted, err := EnsureBlockExists(e.ctx, 102, newHash)
	if err != nil {
		t.Fatalf("EnsureBlockExists after reorg: %v", err)
	}
	if !inserted {
		t.Error("divergent block was not re-inserted")
	}

	// Blocks 102.. were rolled back; 102 was re-inserted with the fork hash.
	if got := e.blockHashInDB(t, 102); got != newHash {
		t.Errorf("block 102 hash = %s, want fork hash %s", got, newHash)
	}
	for n := int64(103); n <= 105; n++ {
		if _, err := e.storage.Get_block_hash(n); err == nil {
			t.Errorf("block %d survived the chain split", n)
		}
	}
	// The tx and evt_log cascade away with their block.
	if _, err := e.storage.Get_transaction_id_by_hash(tx.Hash().Hex()); err == nil {
		t.Error("transaction survived the chain split")
	}
	if _, err := e.storage.Get_event_log(evtID); err == nil {
		t.Error("event log survived the chain split")
	}
	// Watermark: HandleChainSplit rewinds to divergent-1, the re-inserted
	// block advances it to the divergent height.
	last, err := e.storage.Get_last_block_num()
	if err != nil {
		t.Fatalf("Get_last_block_num: %v", err)
	}
	if last != 102 {
		t.Errorf("last_block after split = %d, want 102", last)
	}
	// Blocks below the split are untouched.
	if _, err := e.storage.Get_block_hash(101); err != nil {
		t.Errorf("block 101 lost in the chain split: %v", err)
	}
}

func TestHandleChainSplitBelowWatermarkIsNoop(t *testing.T) {
	e := newEnv(t)
	if _, err := EnsureBlockExists(e.ctx, 100, e.chain.BlockHash(100).Hex()); err != nil {
		t.Fatalf("seeding block: %v", err)
	}
	// Divergent block above the watermark: nothing to delete.
	if err := HandleChainSplit(e.ctx, 200); err != nil {
		t.Fatalf("HandleChainSplit: %v", err)
	}
	if _, err := e.storage.Get_block_hash(100); err != nil {
		t.Errorf("block 100 deleted by a no-op chain split: %v", err)
	}
	last, err := e.storage.Get_last_block_num()
	if err != nil {
		t.Fatalf("Get_last_block_num: %v", err)
	}
	if last != 100 {
		t.Errorf("watermark moved by a no-op chain split: %d", last)
	}
}

func TestEnsureTransactionExistsFromRPC(t *testing.T) {
	e := newEnv(t)
	if _, err := EnsureBlockExists(e.ctx, 100, e.chain.BlockHash(100).Hex()); err != nil {
		t.Fatalf("seeding block: %v", err)
	}
	to := ethcommon.HexToAddress("0x2000000000000000000000000000000000000002")
	tx := e.chain.AddTx(100, to, []byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee})

	txID, inserted, err := EnsureTransactionExists(e.ctx, tx.Hash(), 100)
	if err != nil {
		t.Fatalf("EnsureTransactionExists: %v", err)
	}
	if !inserted || txID == 0 {
		t.Errorf("tx not inserted from RPC: id=%d inserted=%v", txID, inserted)
	}

	// Second call resolves from the database without inserting.
	txID2, inserted2, err := EnsureTransactionExists(e.ctx, tx.Hash(), 100)
	if err != nil {
		t.Fatalf("EnsureTransactionExists (second call): %v", err)
	}
	if inserted2 || txID2 != txID {
		t.Errorf("second call: id=%d inserted=%v, want cached id=%d", txID2, inserted2, txID)
	}

	// The full insert path stored sender, input signature and gas data.
	var fromAddr, inputSig string
	var gasUsed int64
	err = e.db.SQL.QueryRow(`
		SELECT a.addr, t.input_sig, t.gas_used FROM transaction t
		JOIN address a ON a.address_id = t.from_aid WHERE t.id = $1`, txID).
		Scan(&fromAddr, &inputSig, &gasUsed)
	if err != nil {
		t.Fatalf("reading inserted tx: %v", err)
	}
	if fromAddr != e.chain.Sender().Hex() {
		t.Errorf("stored sender = %s, want %s (signature recovery)", fromAddr, e.chain.Sender().Hex())
	}
	if inputSig != "0xaabbccdd" {
		t.Errorf("stored input_sig = %q, want 0xaabbccdd", inputSig)
	}
	if gasUsed != 21000 {
		t.Errorf("stored gas_used = %d, want 21000", gasUsed)
	}
}

func TestEnsureTransactionExistsArchiveFallback(t *testing.T) {
	e := newEnv(t)
	if _, err := EnsureBlockExists(e.ctx, 100, e.chain.BlockHash(100).Hex()); err != nil {
		t.Fatalf("seeding block: %v", err)
	}

	// The tx exists only in the archive: the RPC node reports "not found".
	txHash := "0xcccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc01"
	fromAid, err := e.storage.Lookup_or_create_address("0x2100000000000000000000000000000000000021", 100, 0)
	if err != nil {
		t.Fatalf("creating archive from-address: %v", err)
	}
	_, err = e.db.SQL.Exec(`
		INSERT INTO arch_tx(block_num, tx_hash, tx_index, from_aid, to_aid, value,
		                    gas_used, gas_price, input_sig, num_logs, ctrct_create)
		VALUES (100, $1, 3, $2, 0, 0, 42000, 2, '0xdeadbeef', 1, false)`, txHash, fromAid)
	if err != nil {
		t.Fatalf("seeding archive tx: %v", err)
	}

	txID, inserted, err := EnsureTransactionExists(e.ctx, ethcommon.HexToHash(txHash), 100)
	if err != nil {
		t.Fatalf("EnsureTransactionExists (archive): %v", err)
	}
	if !inserted || txID == 0 {
		t.Errorf("archive tx not inserted: id=%d inserted=%v", txID, inserted)
	}
	var gasUsed, gotFrom int64
	if err := e.db.SQL.QueryRow(`SELECT gas_used, from_aid FROM transaction WHERE id=$1`, txID).Scan(&gasUsed, &gotFrom); err != nil {
		t.Fatalf("reading archived insert: %v", err)
	}
	if gasUsed != 42000 || gotFrom != fromAid {
		t.Errorf("archive data not preserved: gas_used=%d from_aid=%d", gasUsed, gotFrom)
	}
}

func TestEnsureTransactionExistsMinimalFallback(t *testing.T) {
	e := newEnv(t)
	if _, err := EnsureBlockExists(e.ctx, 100, e.chain.BlockHash(100).Hex()); err != nil {
		t.Fatalf("seeding block: %v", err)
	}

	// Neither the RPC node nor the archive knows this tx: a minimal record
	// is created as the last resort.
	txHash := ethcommon.HexToHash("0xdddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd02")
	txID, inserted, err := EnsureTransactionExists(e.ctx, txHash, 100)
	if err != nil {
		t.Fatalf("EnsureTransactionExists (minimal): %v", err)
	}
	if !inserted || txID == 0 {
		t.Errorf("minimal tx not inserted: id=%d inserted=%v", txID, inserted)
	}
	var fromAid int64
	var value string
	if err := e.db.SQL.QueryRow(`SELECT from_aid, value FROM transaction WHERE id=$1`, txID).Scan(&fromAid, &value); err != nil {
		t.Fatalf("reading minimal insert: %v", err)
	}
	if fromAid != 0 {
		t.Errorf("minimal record from_aid = %d, want 0", fromAid)
	}
}

func TestInsertEventLogDeduplicatesByBlockAndIndex(t *testing.T) {
	e := newEnv(t)
	if _, err := EnsureBlockExists(e.ctx, 100, e.chain.BlockHash(100).Hex()); err != nil {
		t.Fatalf("seeding block: %v", err)
	}
	to := ethcommon.HexToAddress("0x2000000000000000000000000000000000000002")
	tx := e.chain.AddTx(100, to, nil)
	txID, _, err := EnsureTransactionExists(e.ctx, tx.Hash(), 100)
	if err != nil {
		t.Fatalf("EnsureTransactionExists: %v", err)
	}

	logTemplate := types.Log{
		Address:     to,
		Topics:      []ethcommon.Hash{ethcommon.BigToHash(big.NewInt(7))},
		Data:        []byte{0x01},
		BlockNumber: 100,
		Index:       4,
	}
	firstID, err := InsertEventLog(e.ctx, logTemplate, txID)
	if err != nil {
		t.Fatalf("InsertEventLog: %v", err)
	}
	// Re-inserting the same (block, log index) replaces the previous row —
	// the reorg-tolerant dedup path.
	secondID, err := InsertEventLog(e.ctx, logTemplate, txID)
	if err != nil {
		t.Fatalf("InsertEventLog (repeat): %v", err)
	}
	if secondID == firstID {
		t.Errorf("expected a fresh evt_log id after dedup-replace, got the same id %d", firstID)
	}
	if _, err := e.storage.Get_event_log(firstID); err == nil {
		t.Error("old evt_log row survived the dedup-replace")
	}
	var count int
	if err := e.db.SQL.QueryRow(`SELECT COUNT(*) FROM evt_log WHERE block_num=100 AND log_index=4`).Scan(&count); err != nil {
		t.Fatalf("counting evt_log rows: %v", err)
	}
	if count != 1 {
		t.Errorf("evt_log rows for (100,4) = %d, want 1", count)
	}
}

func TestFetchEventsFiltersByRangeAndContract(t *testing.T) {
	e := newEnv(t)
	watched := ethcommon.HexToAddress("0x2000000000000000000000000000000000000002")
	other := ethcommon.HexToAddress("0x3000000000000000000000000000000000000003")

	tx1 := e.chain.AddTx(100, watched, nil)
	e.chain.AttachLogs(tx1.Hash(), []*types.Log{{Address: watched, Topics: []ethcommon.Hash{{}}, BlockNumber: 100, TxHash: tx1.Hash()}})
	tx2 := e.chain.AddTx(101, other, nil)
	e.chain.AttachLogs(tx2.Hash(), []*types.Log{{Address: other, Topics: []ethcommon.Hash{{}}, BlockNumber: 101, TxHash: tx2.Hash()}})

	logs, err := FetchEvents(e.ctx.EthClient, 90, 110, []ethcommon.Address{watched})
	if err != nil {
		t.Fatalf("FetchEvents: %v", err)
	}
	if len(logs) != 1 || logs[0].Address != watched {
		t.Errorf("FetchEvents = %+v, want exactly the watched-contract log", logs)
	}

	tip, err := GetCurrentBlockNumber(e.ctx.EthClient)
	if err != nil {
		t.Fatalf("GetCurrentBlockNumber: %v", err)
	}
	if tip != 101 {
		t.Errorf("GetCurrentBlockNumber = %d, want 101", tip)
	}
}
