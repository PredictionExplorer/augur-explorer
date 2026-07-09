//go:build integration

// Integration tests for Run: the full polling loop against a real database
// and the fake chain — batch processing end to end, retry after transient
// failures, the mid-block watermark regression, shutdown-mid-batch semantics,
// reorg handling from inside the loop, and the evt_log backfill.
package indexer

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"testing"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/prometheus/client_golang/prometheus"
	promtestutil "github.com/prometheus/client_golang/prometheus/testutil"
)

var watchedContract = ethcommon.HexToAddress("0x2000000000000000000000000000000000000002")

// topicA/topicB are arbitrary event signatures for loop tests.
var (
	topicA = ethcommon.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	topicB = ethcommon.HexToHash("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
)

// addLog records one transaction carrying one log at the given block and
// returns the transaction.
func (e *env) addLog(t *testing.T, block int64, topic ethcommon.Hash, logIndex uint) *types.Transaction {
	t.Helper()
	tx := e.chain.AddTx(block, watchedContract, nil)
	e.chain.AttachLogs(tx.Hash(), []*types.Log{{
		Address:     watchedContract,
		Topics:      []ethcommon.Hash{topic},
		BlockNumber: uint64(block),
		BlockHash:   e.chain.BlockHash(block),
		TxHash:      tx.Hash(),
		Index:       logIndex,
	}})
	return tx
}

// recordingProcessor tracks every processed evt_log id.
type recordingProcessor struct {
	mu      sync.Mutex
	ids     []int64
	failIDs map[int64]int // remaining failures per evt id
	failN   int           // fail the first N calls regardless of id
	onCall  func(callNum int)
}

func (p *recordingProcessor) process(ctx context.Context, evtID int64) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	call := len(p.ids) + 1
	if p.onCall != nil {
		p.onCall(call)
	}
	if p.failN > 0 {
		p.failN--
		return errors.New("simulated processor failure")
	}
	if p.failIDs != nil && p.failIDs[evtID] > 0 {
		p.failIDs[evtID]--
		return errors.New("simulated processor failure")
	}
	p.ids = append(p.ids, evtID)
	return nil
}

func (p *recordingProcessor) processed() []int64 {
	p.mu.Lock()
	defer p.mu.Unlock()
	out := make([]int64, len(p.ids))
	copy(out, p.ids)
	return out
}

// loopEngine rebuilds the env's engine with loop dependencies and fast test
// timings.
func loopEngine(t *testing.T, e *env, progress Progress, proc ProcessFunc, client Client) *Engine {
	t.Helper()
	if client == nil {
		client = e.client
	}
	metrics := NewMetrics(prometheus.NewRegistry())
	engine, err := New(Config{
		Store:         e.st,
		Client:        client,
		Progress:      progress,
		Process:       proc,
		Contracts:     []ethcommon.Address{watchedContract},
		Logger:        slog.New(slog.DiscardHandler),
		Metrics:       metrics,
		Batch:         BatchConfig{Initial: 1000, Min: 10, Max: 10000},
		Retry:         RetryConfig{MaxConsecutiveFailures: 25, MinDelay: time.Millisecond, MaxDelay: 5 * time.Millisecond},
		CaughtUpDelay: time.Millisecond,
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	e.metrics = metrics
	return engine
}

// runUntil runs the engine and cancels its context once cond() holds (checked
// every millisecond). It fails the test if Run errors or cond never holds.
func runUntil(t *testing.T, engine *Engine, cond func() bool) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	go func() { done <- engine.Run(ctx) }()
	deadline := time.After(60 * time.Second)
	for {
		select {
		case err := <-done:
			t.Fatalf("Run returned before the condition held: %v", err)
		case <-deadline:
			t.Fatal("condition not reached within 60s")
		case <-time.After(time.Millisecond):
			if cond() {
				cancel()
				if err := <-done; err != nil {
					t.Fatalf("Run = %v, want nil after cancellation", err)
				}
				return
			}
		}
	}
}

// liveEvtIDs returns the current evt_log ids for the watched contract range.
func (e *env) liveEvtIDs(t *testing.T) map[int64]bool {
	t.Helper()
	rows, err := e.db.SQL.Query(`SELECT id FROM evt_log`)
	if err != nil {
		t.Fatalf("listing evt_log ids: %v", err)
	}
	defer func() { _ = rows.Close() }()
	out := make(map[int64]bool)
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			t.Fatalf("scanning evt_log id: %v", err)
		}
		out[id] = true
	}
	return out
}

func requireAllLiveProcessed(t *testing.T, e *env, proc *recordingProcessor) {
	t.Helper()
	processed := make(map[int64]bool)
	for _, id := range proc.processed() {
		processed[id] = true
	}
	for id := range e.liveEvtIDs(t) {
		if !processed[id] {
			t.Errorf("evt_log id %d exists in the database but was never processed", id)
		}
	}
}

func TestRunProcessesBatchesEndToEnd(t *testing.T) {
	e := newEnv(t)
	e.addLog(t, 100, topicA, 0)
	e.addLog(t, 100, topicB, 1) // second tx in the same block
	e.addLog(t, 103, topicA, 0)
	e.chain.EnsureBlock(105) // chain head above the last event

	progress := &fakeProgress{last: 99}
	proc := &recordingProcessor{}
	engine := loopEngine(t, e, progress, proc.process, nil)

	runUntil(t, engine, func() bool {
		progress.mu.Lock()
		defer progress.mu.Unlock()
		return progress.last >= 105
	})

	if got := len(proc.processed()); got != 3 {
		t.Errorf("processed %d events, want 3", got)
	}
	requireAllLiveProcessed(t, e, proc)

	// All three logs are stored, blocks inserted, watermark at the head.
	if got := len(e.liveEvtIDs(t)); got != 3 {
		t.Errorf("evt_log rows = %d, want 3", got)
	}
	if got := promtestutil.ToFloat64(e.metrics.lastBlock); got != 105 {
		t.Errorf("rwcg_etl_last_block = %v, want 105", got)
	}
	// No topic-name resolver was configured: everything counts as "other".
	if got := promtestutil.ToFloat64(e.metrics.eventsTotal.WithLabelValues("other")); got != 3 {
		t.Errorf(`rwcg_etl_events_total{type="other"} = %v, want 3`, got)
	}
}

func TestRunTopicNameLabelsEventMetrics(t *testing.T) {
	e := newEnv(t)
	e.addLog(t, 100, topicA, 0)
	e.addLog(t, 101, topicB, 0)

	progress := &fakeProgress{last: 99}
	proc := &recordingProcessor{}
	engine := loopEngine(t, e, progress, proc.process, nil)
	engine.topicName = func(topic ethcommon.Hash) string {
		if topic == topicA {
			return "EventAlpha"
		}
		return ""
	}

	runUntil(t, engine, func() bool {
		progress.mu.Lock()
		defer progress.mu.Unlock()
		return progress.last >= 101
	})

	if got := promtestutil.ToFloat64(e.metrics.eventsTotal.WithLabelValues("EventAlpha")); got != 1 {
		t.Errorf(`events_total{type="EventAlpha"} = %v, want 1`, got)
	}
	if got := promtestutil.ToFloat64(e.metrics.eventsTotal.WithLabelValues("other")); got != 1 {
		t.Errorf(`events_total{type="other"} = %v, want 1 (unmapped topic)`, got)
	}
}

func TestRunRetriesBatchAfterTransientProcessorFailure(t *testing.T) {
	e := newEnv(t)
	e.addLog(t, 100, topicA, 0)
	e.addLog(t, 101, topicB, 0)

	progress := &fakeProgress{last: 99}
	proc := &recordingProcessor{failN: 2} // the first two process calls blip
	engine := loopEngine(t, e, progress, proc.process, nil)

	runUntil(t, engine, func() bool {
		progress.mu.Lock()
		defer progress.mu.Unlock()
		return progress.last >= 101
	})

	// The loop retried in-process instead of crashing, and every surviving
	// event row went through the processor.
	requireAllLiveProcessed(t, e, proc)
	if got := promtestutil.ToFloat64(e.metrics.failuresTotal.WithLabelValues("process")); got < 1 {
		t.Errorf(`failures_total{stage="process"} = %v, want >= 1`, got)
	}
	if got := len(e.liveEvtIDs(t)); got != 2 {
		t.Errorf("evt_log rows = %d, want 2", got)
	}
}

// flakyTxClient fails TransactionByHash a fixed number of times per hash with
// a hard (non-"not found") error, then delegates to the real client.
type flakyTxClient struct {
	Client
	mu       sync.Mutex
	failures map[ethcommon.Hash]int
}

func (f *flakyTxClient) TransactionByHash(ctx context.Context, hash ethcommon.Hash) (*types.Transaction, bool, error) {
	f.mu.Lock()
	if f.failures[hash] > 0 {
		f.failures[hash]--
		f.mu.Unlock()
		return nil, false, errors.New("simulated connection failure")
	}
	f.mu.Unlock()
	return f.Client.TransactionByHash(ctx, hash)
}

// TestRunMidBlockFailureDoesNotSkipRemainingLogs is the regression test for a
// bug in the legacy loops: after a pipeline failure on the second log of a
// block whose first log had succeeded, the watermark advanced to that block,
// so the failed log was never fetched again — silent event loss. The engine
// only ever acknowledges completed block boundaries.
func TestRunMidBlockFailureDoesNotSkipRemainingLogs(t *testing.T) {
	e := newEnv(t)
	e.addLog(t, 100, topicA, 0)
	txB := e.addLog(t, 100, topicB, 1) // same block, second transaction

	flaky := &flakyTxClient{Client: e.client, failures: map[ethcommon.Hash]int{txB.Hash(): 1}}
	progress := &fakeProgress{last: 99}
	proc := &recordingProcessor{}
	engine := loopEngine(t, e, progress, proc.process, flaky)

	runUntil(t, engine, func() bool {
		progress.mu.Lock()
		defer progress.mu.Unlock()
		return progress.last >= 100
	})

	// Both logs of block 100 made it into evt_log and through the processor.
	if got := len(e.liveEvtIDs(t)); got != 2 {
		t.Fatalf("evt_log rows for block 100 = %d, want 2 (the failed log was skipped forever)", got)
	}
	requireAllLiveProcessed(t, e, proc)
	if got := promtestutil.ToFloat64(e.metrics.failuresTotal.WithLabelValues("transaction")); got != 1 {
		t.Errorf(`failures_total{stage="transaction"} = %v, want 1`, got)
	}
	// Exactly one acknowledgment: the retried batch that completed both
	// logs. (The legacy loop would have acknowledged block 100 after the
	// first failure — with only one of its two logs stored — and never
	// fetched the block again; the evt_log count above pins the fix.)
	if writes := progress.writesCopy(); len(writes) != 1 || writes[0] != 100 {
		t.Errorf("watermark writes = %v, want [100]", writes)
	}
}

func TestRunShutdownMidBatchCompletesTheBatch(t *testing.T) {
	e := newEnv(t)
	e.addLog(t, 100, topicA, 0)
	e.addLog(t, 101, topicB, 0)
	e.addLog(t, 102, topicA, 0)

	ctx, cancel := context.WithCancel(context.Background())
	progress := &fakeProgress{last: 99}
	proc := &recordingProcessor{}
	// Cancel while the first event of the batch is being processed: the
	// remaining events and the watermark write must still complete.
	proc.onCall = func(call int) {
		if call == 1 {
			cancel()
		}
	}
	engine := loopEngine(t, e, progress, proc.process, nil)

	done := make(chan error, 1)
	go func() { done <- engine.Run(ctx) }()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Run = %v, want nil (graceful shutdown)", err)
		}
	case <-time.After(60 * time.Second):
		t.Fatal("Run did not return within 60s")
	}

	if got := len(proc.processed()); got != 3 {
		t.Errorf("processed %d events before exit, want the whole batch of 3", got)
	}
	if got := progress.writesCopy(); len(got) != 1 || got[0] != 102 {
		t.Errorf("watermark writes = %v, want [102] (batch acknowledged before exit)", got)
	}
}

func TestRunHandlesReorgDetectedByTheLoop(t *testing.T) {
	e := newEnv(t)
	ctx := context.Background()

	// The database knows block 100 of the old fork.
	if _, err := e.engine.EnsureBlockExists(ctx, 100, e.chain.BlockHash(100).Hex()); err != nil {
		t.Fatalf("seeding old-fork block: %v", err)
	}

	// The chain reorganizes; the new fork carries an event in block 100.
	e.chain.Reorg(100)
	e.addLog(t, 100, topicA, 0)

	progress := &fakeProgress{last: 99}
	proc := &recordingProcessor{}
	engine := loopEngine(t, e, progress, proc.process, nil)

	runUntil(t, engine, func() bool {
		progress.mu.Lock()
		defer progress.mu.Unlock()
		return progress.last >= 100
	})

	if got := promtestutil.ToFloat64(e.metrics.reorgsTotal); got != 1 {
		t.Errorf("rwcg_etl_reorgs_total = %v, want 1", got)
	}
	if got := e.blockHashInDB(t, 100); got != e.chain.BlockHash(100).Hex() {
		t.Errorf("block 100 hash = %s, want the new fork hash", got)
	}
	requireAllLiveProcessed(t, e, proc)
}

func TestBackfillContractEvtLogsInsertsAndSkips(t *testing.T) {
	e := newEnv(t)
	ctx := context.Background()
	e.addLog(t, 100, topicA, 0)
	e.addLog(t, 102, topicB, 0)

	stats, err := e.engine.BackfillContractEvtLogs(ctx, []ethcommon.Address{watchedContract}, 90, 110, 5)
	if err != nil {
		t.Fatalf("BackfillContractEvtLogs: %v", err)
	}
	if stats.LogsSeen != 2 || stats.Inserted != 2 || stats.Skipped != 0 {
		t.Errorf("first run stats = %+v, want 2 seen / 2 inserted / 0 skipped", stats)
	}
	if got := len(e.liveEvtIDs(t)); got != 2 {
		t.Errorf("evt_log rows = %d, want 2", got)
	}

	// A second pass finds everything present and inserts nothing.
	stats, err = e.engine.BackfillContractEvtLogs(ctx, []ethcommon.Address{watchedContract}, 90, 110, 5)
	if err != nil {
		t.Fatalf("BackfillContractEvtLogs (repeat): %v", err)
	}
	if stats.LogsSeen != 2 || stats.Inserted != 0 || stats.Skipped != 2 {
		t.Errorf("second run stats = %+v, want 2 seen / 0 inserted / 2 skipped", stats)
	}
	if got := len(e.liveEvtIDs(t)); got != 2 {
		t.Errorf("evt_log rows after repeat = %d, want 2 (no duplicates)", got)
	}
}
