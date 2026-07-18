//go:build integration

// Atomicity matrix for the per-block ingestion transaction (ADR-0010):
// inject a failure at each isolatable pipeline stage of a multi-write block
// and prove that nothing of the block persisted — no block, transaction or
// evt_log rows, no domain writes, no address rows, no watermark movement —
// then prove the retry converges to exactly the state a clean run produces.
// The pre-transactional engine left partially applied blocks behind on
// every one of these failures and relied on idempotent replay to repair
// them; these tests pin the stronger guarantee.
package indexer

import (
	"context"
	"errors"
	"strings"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// rowCount counts the rows of one table through the raw connection,
// bypassing the engine and every cache.
func (e *env) rowCount(t *testing.T, table string) int {
	t.Helper()
	var n int
	// The table name comes from the test's own literal set below.
	if err := e.db.SQL.QueryRow("SELECT COUNT(*) FROM " + table).Scan(&n); err != nil {
		t.Fatalf("counting %s rows: %v", table, err)
	}
	return n
}

// requireBlock100Absent asserts that no trace of the tests' block 100
// survived a rolled back ingestion transaction: no block row, no
// transactions, no event logs, no address rows and no watermark movement.
func (e *env) requireBlock100Absent(t *testing.T, progress *fakeProgress) {
	t.Helper()
	ctx := context.Background()
	if _, err := e.st.BlockHash(ctx, 100); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("block 100 row survived the rollback (err=%v)", err)
	}
	if got := e.rowCount(t, "transaction"); got != 0 {
		t.Errorf("transaction rows = %d, want 0 after rollback", got)
	}
	if got := len(e.liveEvtIDs(t)); got != 0 {
		t.Errorf("evt_log rows = %d, want 0 after rollback", got)
	}
	if got := e.rowCount(t, "address"); got != 0 {
		t.Errorf("address rows = %d, want 0 after rollback", got)
	}
	last, err := e.st.LastBlockNum(ctx)
	if err != nil || last != 0 {
		t.Errorf("last_block after rollback = (%d, %v), want (0, nil)", last, err)
	}
	if writes := progress.writesCopy(); len(writes) != 0 {
		t.Errorf("progress writes after rollback = %v, want none", writes)
	}
}

// TestProcessBlockRollsBackOnTransactionStageFailure injects a hard RPC
// failure on the second transaction of a block whose first log has already
// been stored inside the transaction. The whole block must vanish — the
// legacy pipeline left the first log's block/transaction/evt_log/address
// rows committed. The retry must then succeed from scratch, which also
// proves the address-cache overlay was discarded with the rollback (a
// poisoned cache would resolve the contract to a rolled-back address id and
// fail the evt_log insert's foreign key).
func TestProcessBlockRollsBackOnTransactionStageFailure(t *testing.T) {
	e := newEnv(t)
	ctx := context.Background()
	e.addLog(t, 100, topicA, 0)
	txB := e.addLog(t, 100, topicB, 1) // same block, second transaction

	flaky := &flakyTxClient{Client: e.client, failures: map[ethcommon.Hash]int{txB.Hash(): 1}}
	progress := &fakeProgress{last: 99}
	proc := &recordingProcessor{}
	engine := loopEngine(t, e, progress, proc.process, flaky)

	logs, err := FetchLogs(ctx, e.client, 100, 100, []ethcommon.Address{watchedContract})
	if err != nil {
		t.Fatalf("FetchLogs: %v", err)
	}
	if len(logs) != 2 {
		t.Fatalf("fetched %d logs, want 2", len(logs))
	}

	lastCompleted, stage, err := engine.processBatch(ctx, logs)
	if err == nil {
		t.Fatal("processBatch succeeded, want the injected transaction-stage failure")
	}
	if lastCompleted != 0 || stage != "transaction" {
		t.Errorf("processBatch = (last %d, stage %q), want (0, transaction)", lastCompleted, stage)
	}
	e.requireBlock100Absent(t, progress)

	// The retry (the RPC blip is gone) commits the whole block atomically.
	lastCompleted, stage, err = engine.processBatch(ctx, logs)
	if err != nil {
		t.Fatalf("retry processBatch = (stage %q, %v), want success", stage, err)
	}
	if lastCompleted != 100 {
		t.Errorf("retry lastCompleted = %d, want 100", lastCompleted)
	}
	if got := len(e.liveEvtIDs(t)); got != 2 {
		t.Errorf("evt_log rows after retry = %d, want 2", got)
	}
	if writes := progress.writesCopy(); len(writes) != 1 || writes[0] != 100 {
		t.Errorf("progress writes after retry = %v, want [100]", writes)
	}
	requireAllLiveProcessed(t, e, proc)
}

// TestProcessBatchFailingBlockDoesNotDisturbCommittedBlocks drives a
// two-block batch whose second block's processor fails: the first block
// must stay committed with its own watermark write, the second block must
// leave nothing behind, and the retry of only the uncommitted remainder
// (exactly what Run refetches) must converge without touching the first
// block's rows.
func TestProcessBatchFailingBlockDoesNotDisturbCommittedBlocks(t *testing.T) {
	e := newEnv(t)
	ctx := context.Background()
	e.addLog(t, 100, topicA, 0)
	e.addLog(t, 100, topicB, 1)
	e.addLog(t, 103, topicA, 0)

	progress := &fakeProgress{last: 99}
	// Block 100 commits ids 1 and 2; block 103's log gets id 3 and fails.
	proc := &recordingProcessor{failIDs: map[int64]int{3: 1}}
	engine := loopEngine(t, e, progress, proc.process, nil)

	logs, err := FetchLogs(ctx, e.client, 100, 105, []ethcommon.Address{watchedContract})
	if err != nil {
		t.Fatalf("FetchLogs: %v", err)
	}
	if len(logs) != 3 {
		t.Fatalf("fetched %d logs, want 3", len(logs))
	}

	lastCompleted, stage, err := engine.processBatch(ctx, logs)
	if err == nil {
		t.Fatal("processBatch succeeded, want the injected processor failure")
	}
	if lastCompleted != 100 || stage != "process" {
		t.Errorf("processBatch = (last %d, stage %q), want (100, process)", lastCompleted, stage)
	}

	// Block 100 committed atomically with its watermark; block 103 vanished.
	if got := len(e.liveEvtIDs(t)); got != 2 {
		t.Errorf("evt_log rows = %d, want block 100's 2", got)
	}
	if _, err := e.st.BlockHash(ctx, 103); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("block 103 row survived its rollback (err=%v)", err)
	}
	last, err := e.st.LastBlockNum(ctx)
	if err != nil || last != 100 {
		t.Errorf("last_block = (%d, %v), want (100, nil)", last, err)
	}
	if writes := progress.writesCopy(); len(writes) != 1 || writes[0] != 100 {
		t.Errorf("progress writes = %v, want [100] (committed block acknowledged itself)", writes)
	}

	// Run resumes after the committed boundary: refetch only 101+.
	remainder, err := FetchLogs(ctx, e.client, 101, 105, []ethcommon.Address{watchedContract})
	if err != nil {
		t.Fatalf("FetchLogs remainder: %v", err)
	}
	before := e.liveEvtIDs(t)
	lastCompleted, stage, err = engine.processBatch(ctx, remainder)
	if err != nil {
		t.Fatalf("retry processBatch = (stage %q, %v), want success", stage, err)
	}
	if lastCompleted != 103 {
		t.Errorf("retry lastCompleted = %d, want 103", lastCompleted)
	}
	after := e.liveEvtIDs(t)
	for id := range before {
		if !after[id] {
			t.Errorf("committed evt_log id %d disappeared during the remainder retry", id)
		}
	}
	if got := len(after); got != 3 {
		t.Errorf("evt_log rows after retry = %d, want 3", got)
	}
	if writes := progress.writesCopy(); len(writes) != 2 || writes[1] != 103 {
		t.Errorf("progress writes = %v, want [100 103]", writes)
	}
	requireAllLiveProcessed(t, e, proc)
}

// TestProcessBlockRollsBackOnBlockStageFailure drives the block stage's
// error arm inside the transaction: the fetched log claims a hash the
// chain does not confirm, so EnsureBlockExists fails and nothing commits.
func TestProcessBlockRollsBackOnBlockStageFailure(t *testing.T) {
	e := newEnv(t)
	ctx := context.Background()
	e.addLog(t, 100, topicA, 0)

	logs, err := FetchLogs(ctx, e.client, 100, 100, []ethcommon.Address{watchedContract})
	if err != nil {
		t.Fatalf("FetchLogs: %v", err)
	}
	logs[0].BlockHash = ethcommon.HexToHash(
		"0x1111111111111111111111111111111111111111111111111111111111111111")

	progress := &fakeProgress{last: 99}
	proc := &recordingProcessor{}
	engine := loopEngine(t, e, progress, proc.process, nil)

	lastCompleted, stage, err := engine.processBatch(ctx, logs)
	if err == nil || !strings.Contains(err.Error(), "hash mismatch") {
		t.Fatalf("processBatch = %v, want the block hash mismatch", err)
	}
	if lastCompleted != 0 || stage != "block" {
		t.Errorf("processBatch = (last %d, stage %q), want (0, block)", lastCompleted, stage)
	}
	e.requireBlock100Absent(t, progress)
}

// TestProcessBlockRollsBackOnEventLogStageFailure hides the evt_log table
// for one attempt: the block and transaction stages succeed inside the
// transaction, the event-log insert fails, and the rollback leaves no
// block, transaction or address rows behind.
func TestProcessBlockRollsBackOnEventLogStageFailure(t *testing.T) {
	e := newEnv(t)
	ctx := context.Background()
	e.addLog(t, 100, topicA, 0)

	logs, err := FetchLogs(ctx, e.client, 100, 100, []ethcommon.Address{watchedContract})
	if err != nil {
		t.Fatalf("FetchLogs: %v", err)
	}

	progress := &fakeProgress{last: 99}
	proc := &recordingProcessor{}
	engine := loopEngine(t, e, progress, proc.process, nil)

	if _, err := e.db.SQL.Exec(`ALTER TABLE evt_log RENAME TO evt_log_hidden`); err != nil {
		t.Fatalf("hiding evt_log: %v", err)
	}
	restored := false
	restore := func() {
		if restored {
			return
		}
		restored = true
		if _, err := e.db.SQL.Exec(`ALTER TABLE evt_log_hidden RENAME TO evt_log`); err != nil {
			t.Fatalf("restoring evt_log: %v", err)
		}
	}
	defer restore()

	lastCompleted, stage, err := engine.processBatch(ctx, logs)
	if err == nil {
		t.Fatal("processBatch succeeded with evt_log hidden")
	}
	if lastCompleted != 0 || stage != "event_log" {
		t.Errorf("processBatch = (last %d, stage %q), want (0, event_log)", lastCompleted, stage)
	}
	restore()
	e.requireBlock100Absent(t, progress)

	// With the table back, the same batch commits cleanly.
	if _, stage, err := engine.processBatch(ctx, logs); err != nil {
		t.Fatalf("retry processBatch = (stage %q, %v), want success", stage, err)
	}
	if got := len(e.liveEvtIDs(t)); got != 1 {
		t.Errorf("evt_log rows after retry = %d, want 1", got)
	}
}

// TestProcessBlockReportsCommitStageOnBeginFailure pins the "commit"
// failure stage: when the transaction itself cannot be opened (or
// committed), the failure is attributed to the transaction machinery, not
// to a pipeline stage.
func TestProcessBlockReportsCommitStageOnBeginFailure(t *testing.T) {
	e := newEnv(t)
	ctx := context.Background()
	e.addLog(t, 100, topicA, 0)

	logs, err := FetchLogs(ctx, e.client, 100, 100, []ethcommon.Address{watchedContract})
	if err != nil {
		t.Fatalf("FetchLogs: %v", err)
	}

	progress := &fakeProgress{last: 99}
	proc := &recordingProcessor{}
	// A second store over a closed pool: InTx fails at Begin before any
	// pipeline stage runs.
	closedPool, err := pgxpool.New(ctx, e.db.ConnString)
	if err != nil {
		t.Fatalf("building closed pool: %v", err)
	}
	closedStore := store.NewFromPool(closedPool)
	closedStore.Close()
	engine := loopEngine(t, e, progress, proc.process, nil)
	engine.store = closedStore

	lastCompleted, stage, err := engine.processBatch(ctx, logs)
	if err == nil || !strings.Contains(err.Error(), "begin transaction") {
		t.Fatalf("processBatch = %v, want wrapped begin failure", err)
	}
	if lastCompleted != 0 || stage != "commit" {
		t.Errorf("processBatch = (last %d, stage %q), want (0, commit)", lastCompleted, stage)
	}
	if got := len(proc.processed()); got != 0 {
		t.Errorf("processor ran %d events despite the failed begin", got)
	}
}

// TestProcessBlockRollsBackWhenWatermarkWriteFails pins the headline
// guarantee of ADR-0010: the watermark write is part of the block's
// transaction, so a block whose data landed but whose acknowledgment failed
// commits neither — the watermark can never disagree with the data.
func TestProcessBlockRollsBackWhenWatermarkWriteFails(t *testing.T) {
	e := newEnv(t)
	ctx := context.Background()
	e.addLog(t, 100, topicA, 0)

	writeErr := errors.New("status table on fire")
	progress := &failingProgress{fakeProgress: fakeProgress{last: 99}, writeErr: writeErr}
	proc := &recordingProcessor{}
	engine := loopEngine(t, e, progress, proc.process, nil)

	logs, err := FetchLogs(ctx, e.client, 100, 100, []ethcommon.Address{watchedContract})
	if err != nil {
		t.Fatalf("FetchLogs: %v", err)
	}

	lastCompleted, stage, err := engine.processBatch(ctx, logs)
	if !errors.Is(err, writeErr) {
		t.Fatalf("processBatch = %v, want the watermark write failure", err)
	}
	if lastCompleted != 0 || stage != "watermark" {
		t.Errorf("processBatch = (last %d, stage %q), want (0, watermark)", lastCompleted, stage)
	}
	e.requireBlock100Absent(t, &progress.fakeProgress)
}
