// Unit tests (no Docker) for the event-handler registry and the stored-log
// processor: dispatch by topic, source filtering, error propagation, name
// validation and the RLP reconstruction path.

package indexer

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/PredictionExplorer/augur-explorer/internal/primitives"
)

var (
	handlerTopicA = common.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	handlerTopicB = common.HexToHash("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")

	contractX = common.HexToAddress("0x1000000000000000000000000000000000000001")
	contractY = common.HexToAddress("0x2000000000000000000000000000000000000002")
)

// testEvent is the domain type the test handlers decode into.
type testEvent struct {
	EvtID int64
	Body  string
}

// recordingHandler builds a handler that records decode/store invocations.
type recordingHandler struct {
	decoded []int64
	stored  []testEvent
}

func (rec *recordingHandler) handler(topic common.Hash, name string, sources ...common.Address) EventHandler {
	return NewHandler(topic, name, sources,
		func(lg *types.Log, elog *primitives.EthereumEventLog) (*testEvent, error) {
			rec.decoded = append(rec.decoded, elog.EvtId)
			return &testEvent{EvtID: elog.EvtId, Body: string(lg.Data)}, nil
		},
		func(_ context.Context, evt *testEvent) error {
			rec.stored = append(rec.stored, *evt)
			return nil
		})
}

func logWith(topic common.Hash, addr common.Address, data string) *types.Log {
	return &types.Log{Address: addr, Topics: []common.Hash{topic}, Data: []byte(data)}
}

func elogWith(id int64) *primitives.EthereumEventLog {
	return &primitives.EthereumEventLog{EvtId: id}
}

func TestRegistryDispatchesByTopicAndSource(t *testing.T) {
	var forX, forY, forB recordingHandler
	reg, err := NewRegistry(
		forX.handler(handlerTopicA, "EventA", contractX),
		forY.handler(handlerTopicA, "EventA", contractY),
		forB.handler(handlerTopicB, "EventB", contractX),
	)
	if err != nil {
		t.Fatalf("NewRegistry: %v", err)
	}

	ctx := context.Background()
	if err := reg.ProcessLog(ctx, logWith(handlerTopicA, contractX, "x"), elogWith(1)); err != nil {
		t.Fatalf("ProcessLog: %v", err)
	}
	if err := reg.ProcessLog(ctx, logWith(handlerTopicA, contractY, "y"), elogWith(2)); err != nil {
		t.Fatalf("ProcessLog: %v", err)
	}

	if len(forX.stored) != 1 || forX.stored[0] != (testEvent{EvtID: 1, Body: "x"}) {
		t.Errorf("contract-X handler stored %+v, want exactly event 1", forX.stored)
	}
	if len(forY.stored) != 1 || forY.stored[0] != (testEvent{EvtID: 2, Body: "y"}) {
		t.Errorf("contract-Y handler stored %+v, want exactly event 2", forY.stored)
	}
	if len(forB.stored) != 0 {
		t.Errorf("topic-B handler stored %+v for topic-A logs", forB.stored)
	}
}

func TestRegistrySkipsUnknownTopicSourceAndEmptyTopics(t *testing.T) {
	var rec recordingHandler
	reg, err := NewRegistry(rec.handler(handlerTopicA, "EventA", contractX))
	if err != nil {
		t.Fatalf("NewRegistry: %v", err)
	}

	ctx := context.Background()
	// Unknown topic: no handler consulted.
	if err := reg.ProcessLog(ctx, logWith(handlerTopicB, contractX, ""), elogWith(1)); err != nil {
		t.Fatalf("unknown topic: %v", err)
	}
	// Known topic, unknown source: filtered before Decode.
	if err := reg.ProcessLog(ctx, logWith(handlerTopicA, contractY, ""), elogWith(2)); err != nil {
		t.Fatalf("unknown source: %v", err)
	}
	// No topics at all (anonymous event): no-op.
	if err := reg.ProcessLog(ctx, &types.Log{Address: contractX}, elogWith(3)); err != nil {
		t.Fatalf("zero topics: %v", err)
	}
	if len(rec.decoded) != 0 {
		t.Errorf("Decode invoked %d times, want 0 (ids %v)", len(rec.decoded), rec.decoded)
	}
}

func TestRegistryEmptySourcesAcceptAnyContract(t *testing.T) {
	var rec recordingHandler
	reg, err := NewRegistry(rec.handler(handlerTopicA, "EventA"))
	if err != nil {
		t.Fatalf("NewRegistry: %v", err)
	}
	for i, addr := range []common.Address{contractX, contractY, {}} {
		if err := reg.ProcessLog(context.Background(), logWith(handlerTopicA, addr, ""), elogWith(int64(i))); err != nil {
			t.Fatalf("ProcessLog(%v): %v", addr, err)
		}
	}
	if len(rec.stored) != 3 {
		t.Errorf("stored %d events, want 3", len(rec.stored))
	}
}

func TestRegistryDecodeErrorStopsDispatch(t *testing.T) {
	decodeErr := errors.New("bad payload")
	var after recordingHandler
	reg, err := NewRegistry(
		NewHandler(handlerTopicA, "EventA", nil,
			func(*types.Log, *primitives.EthereumEventLog) (*testEvent, error) { return nil, decodeErr },
			func(context.Context, *testEvent) error { t.Fatal("store called after decode error"); return nil }),
		after.handler(handlerTopicA, "EventA"),
	)
	if err != nil {
		t.Fatalf("NewRegistry: %v", err)
	}

	err = reg.ProcessLog(context.Background(), logWith(handlerTopicA, contractX, ""), elogWith(7))
	if !errors.Is(err, decodeErr) {
		t.Fatalf("ProcessLog error = %v, want wrapped %v", err, decodeErr)
	}
	if !strings.Contains(err.Error(), "EventA (evt id 7): decode") {
		t.Errorf("decode error lacks handler context: %v", err)
	}
	if len(after.decoded) != 0 {
		t.Error("handler after the failing one still ran")
	}
}

func TestRegistryStoreErrorPropagates(t *testing.T) {
	storeErr := errors.New("insert failed")
	reg, err := NewRegistry(
		NewHandler(handlerTopicA, "EventA", nil,
			func(_ *types.Log, elog *primitives.EthereumEventLog) (*testEvent, error) {
				return &testEvent{EvtID: elog.EvtId}, nil
			},
			func(context.Context, *testEvent) error { return storeErr }),
	)
	if err != nil {
		t.Fatalf("NewRegistry: %v", err)
	}
	if err := reg.ProcessLog(context.Background(), logWith(handlerTopicA, contractX, ""), elogWith(1)); !errors.Is(err, storeErr) {
		t.Fatalf("ProcessLog error = %v, want %v", err, storeErr)
	}
}

func TestRegistryContinuesPastFilteredHandler(t *testing.T) {
	// Two handlers on one topic: a log matching only the second must still
	// reach it (the first is source-filtered, not an error).
	var first, second recordingHandler
	reg, err := NewRegistry(
		first.handler(handlerTopicA, "EventA", contractX),
		second.handler(handlerTopicA, "EventA", contractY),
	)
	if err != nil {
		t.Fatalf("NewRegistry: %v", err)
	}
	if err := reg.ProcessLog(context.Background(), logWith(handlerTopicA, contractY, "y"), elogWith(9)); err != nil {
		t.Fatalf("ProcessLog: %v", err)
	}
	if len(first.stored) != 0 || len(second.stored) != 1 {
		t.Errorf("stored: first %d, second %d; want 0 and 1", len(first.stored), len(second.stored))
	}
}

func TestNewRegistryRejectsBadNames(t *testing.T) {
	var rec recordingHandler
	if _, err := NewRegistry(rec.handler(handlerTopicA, "")); err == nil {
		t.Error("empty event name accepted")
	}
	if _, err := NewRegistry(
		rec.handler(handlerTopicA, "EventA", contractX),
		rec.handler(handlerTopicA, "SomethingElse", contractY),
	); err == nil {
		t.Error("conflicting names for one topic accepted")
	}
}

func TestRegistryTopicName(t *testing.T) {
	var rec recordingHandler
	reg, err := NewRegistry(rec.handler(handlerTopicA, "EventA", contractX))
	if err != nil {
		t.Fatalf("NewRegistry: %v", err)
	}
	if got := reg.TopicName(handlerTopicA); got != "EventA" {
		t.Errorf("TopicName(topicA) = %q, want EventA", got)
	}
	if got := reg.TopicName(handlerTopicB); got != "" {
		t.Errorf("TopicName(unknown) = %q, want empty", got)
	}
}

func TestHandlerStoreRejectsForeignEventType(t *testing.T) {
	h := NewHandler(handlerTopicA, "EventA", nil,
		func(*types.Log, *primitives.EthereumEventLog) (*testEvent, error) { return &testEvent{}, nil },
		func(context.Context, *testEvent) error { return nil })
	if err := h.Store(context.Background(), "not a testEvent"); err == nil {
		t.Error("Store accepted a value of the wrong type")
	}
}

// fakeEventLogSource serves canned evt_log rows to LogProcessor.
type fakeEventLogSource map[int64]primitives.EthereumEventLog

func (f fakeEventLogSource) EventLog(_ context.Context, id int64) (primitives.EthereumEventLog, error) {
	row, ok := f[id]
	if !ok {
		return primitives.EthereumEventLog{}, fmt.Errorf("row %d not found", id)
	}
	return row, nil
}

func mustRLP(t *testing.T, lg *types.Log) []byte {
	t.Helper()
	raw, err := rlp.EncodeToBytes(lg)
	if err != nil {
		t.Fatalf("encoding log: %v", err)
	}
	return raw
}

func TestLogProcessorReconstructsAndDispatches(t *testing.T) {
	var rec recordingHandler
	reg, err := NewRegistry(rec.handler(handlerTopicA, "EventA", contractX))
	if err != nil {
		t.Fatalf("NewRegistry: %v", err)
	}

	raw := mustRLP(t, &types.Log{
		// The stored RLP deliberately carries a different address: the
		// evt_log columns are authoritative, exactly like the legacy path.
		Address: contractY,
		Topics:  []common.Hash{handlerTopicA},
		Data:    []byte("payload"),
	})
	src := fakeEventLogSource{
		42: {
			EvtId:           42,
			BlockNum:        1234,
			TxHash:          "0x00000000000000000000000000000000000000000000000000000000000000aa",
			ContractAddress: contractX.Hex(),
			RlpLog:          raw,
		},
	}

	process := LogProcessor(src, reg)
	if err := process(context.Background(), 42); err != nil {
		t.Fatalf("process: %v", err)
	}
	if len(rec.stored) != 1 || rec.stored[0] != (testEvent{EvtID: 42, Body: "payload"}) {
		t.Fatalf("stored %+v, want event 42 with the RLP payload", rec.stored)
	}
}

func TestLogProcessorZeroTopicsIsNoOp(t *testing.T) {
	var rec recordingHandler
	reg, err := NewRegistry(rec.handler(handlerTopicA, "EventA"))
	if err != nil {
		t.Fatalf("NewRegistry: %v", err)
	}
	src := fakeEventLogSource{
		1: {EvtId: 1, ContractAddress: contractX.Hex(), RlpLog: mustRLP(t, &types.Log{Address: contractX, Data: []byte{1}})},
	}
	if err := LogProcessor(src, reg)(context.Background(), 1); err != nil {
		t.Fatalf("process: %v", err)
	}
	if len(rec.decoded) != 0 {
		t.Error("zero-topic log reached a handler")
	}
}

func TestLogProcessorErrors(t *testing.T) {
	var rec recordingHandler
	reg, err := NewRegistry(rec.handler(handlerTopicA, "EventA"))
	if err != nil {
		t.Fatalf("NewRegistry: %v", err)
	}
	src := fakeEventLogSource{
		1: {EvtId: 1, ContractAddress: contractX.Hex(), RlpLog: []byte{0xff, 0x00, 0x13}},
	}
	process := LogProcessor(src, reg)

	if err := process(context.Background(), 999); err == nil {
		t.Error("missing evt_log row: want error")
	}
	err = process(context.Background(), 1)
	if err == nil || !strings.Contains(err.Error(), "RLP decode") {
		t.Errorf("corrupt RLP: got %v, want RLP decode error", err)
	}
}
