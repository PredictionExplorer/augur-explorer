// The typed event-handler registry: each chain event is described by an
// EventHandler (topic, metric name, emitting contracts, a pure decode step
// and a persisting store step). A Registry dispatches stored logs to every
// handler registered for the log's topic0 whose source-contract filter
// matches, and LogProcessor adapts a Registry to the engine's ProcessFunc.
//
// Decode is deliberately free of database and RPC work — it turns a raw log
// into a typed domain event and nothing else — so decoders can be fuzzed and
// unit-tested without infrastructure. Everything that touches the store, the
// chain, or sibling events of the transaction belongs in Store.

package indexer

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// EventHandler consumes one event type emitted by one (or a fixed set of)
// contract(s). Implementations come from NewHandler, which preserves the
// concrete event type between Decode and Store.
type EventHandler interface {
	// Topic returns the topic0 hash (event signature) this handler consumes.
	Topic() common.Hash
	// Name returns the human-readable event name; it labels the
	// rwcg_etl_events_total metric, so handlers sharing a topic0 must agree
	// on it (enforced by NewRegistry).
	Name() string
	// Sources returns the contracts expected to emit the event. The registry
	// invokes the handler only for logs whose address is in the set; an
	// empty set accepts the event from any contract.
	Sources() []common.Address
	// Decode turns the raw log into a typed domain event. It must be total:
	// malformed input yields an error, never a panic (fuzzed per handler
	// package). No database or RPC access.
	Decode(lg *types.Log, elog *store.EthereumEventLog) (any, error)
	// Store persists the decoded event: enrichment queries, contract reads
	// and the domain writes. event is exactly the value Decode returned.
	Store(ctx context.Context, event any) error
}

// handler is the generic EventHandler adapter built by NewHandler.
type handler[E any] struct {
	topic   common.Hash
	name    string
	sources []common.Address
	decode  func(*types.Log, *store.EthereumEventLog) (E, error)
	store   func(context.Context, E) error
}

// NewHandler pairs a typed decode function with its store function. The
// concrete event type E is erased only at the EventHandler boundary; Store
// re-asserts it, so a registry can never feed a handler another handler's
// event.
func NewHandler[E any](
	topic common.Hash,
	name string,
	sources []common.Address,
	decode func(*types.Log, *store.EthereumEventLog) (E, error),
	store func(context.Context, E) error,
) EventHandler {
	return handler[E]{topic: topic, name: name, sources: sources, decode: decode, store: store}
}

func (h handler[E]) Topic() common.Hash        { return h.topic }
func (h handler[E]) Name() string              { return h.name }
func (h handler[E]) Sources() []common.Address { return h.sources }

func (h handler[E]) Decode(lg *types.Log, elog *store.EthereumEventLog) (any, error) {
	return h.decode(lg, elog)
}

func (h handler[E]) Store(ctx context.Context, event any) error {
	evt, ok := event.(E)
	if !ok {
		return fmt.Errorf("handler %s: Store got %T, want %T", h.name, event, evt)
	}
	return h.store(ctx, evt)
}

// acceptsSource reports whether the handler consumes events emitted by addr.
func acceptsSource(h EventHandler, addr common.Address) bool {
	sources := h.Sources()
	if len(sources) == 0 {
		return true
	}
	for _, s := range sources {
		if s == addr {
			return true
		}
	}
	return false
}

// Registry maps topic0 hashes to their event handlers. Multiple handlers may
// share a topic (the same signature emitted by different contracts); the
// source filter decides which of them run for a given log.
type Registry struct {
	byTopic map[common.Hash][]EventHandler
	names   map[common.Hash]string
	all     []EventHandler
}

// NewRegistry validates and indexes the handler set: every handler needs a
// name, and handlers sharing a topic0 must carry the same name (the metric
// label of a topic must not depend on registration order).
func NewRegistry(handlers ...EventHandler) (*Registry, error) {
	r := &Registry{
		byTopic: make(map[common.Hash][]EventHandler, len(handlers)),
		names:   make(map[common.Hash]string, len(handlers)),
		all:     handlers,
	}
	for _, h := range handlers {
		if h.Name() == "" {
			return nil, fmt.Errorf("indexer: handler for topic %s has an empty event name", h.Topic().Hex())
		}
		if prev, ok := r.names[h.Topic()]; ok && prev != h.Name() {
			return nil, fmt.Errorf("indexer: topic %s carries two event names: %q and %q",
				h.Topic().Hex(), prev, h.Name())
		}
		r.names[h.Topic()] = h.Name()
		r.byTopic[h.Topic()] = append(r.byTopic[h.Topic()], h)
	}
	return r, nil
}

// Handlers returns every registered handler in registration order (used by
// the per-package decode fuzz targets).
func (r *Registry) Handlers() []EventHandler { return r.all }

// TopicName resolves the metric label of a topic0 hash ("" = unknown); it
// plugs directly into Config.TopicName.
func (r *Registry) TopicName(topic common.Hash) string { return r.names[topic] }

// ProcessLog dispatches one log to every registered handler of its topic0
// whose source filter matches. Unknown topics and unknown sources are
// no-ops (FilterLogs fetches whole contracts, so unhandled events are
// routine). Any decode or store error stops the dispatch and is returned to
// the polling loop, which leaves the batch unacknowledged for re-processing.
func (r *Registry) ProcessLog(ctx context.Context, lg *types.Log, elog *store.EthereumEventLog) error {
	if len(lg.Topics) == 0 {
		return nil
	}
	for _, h := range r.byTopic[lg.Topics[0]] {
		if !acceptsSource(h, lg.Address) {
			continue
		}
		event, err := h.Decode(lg, elog)
		if err != nil {
			return fmt.Errorf("%s (evt id %v): decode: %w", h.Name(), elog.EvtID, err)
		}
		if err := h.Store(ctx, event); err != nil {
			return err
		}
	}
	return nil
}

// EventLogSource loads one stored event-log row by id. *store.Store
// satisfies it; unit tests substitute fakes.
type EventLogSource interface {
	EventLog(ctx context.Context, evtlogID int64) (store.EthereumEventLog, error)
}

// LogProcessor builds the engine's ProcessFunc: load the stored evt_log row,
// reconstruct the Ethereum log from its RLP and dispatch it through the
// registry. All failures — a missing row, a corrupt RLP payload or a handler
// error — are returned to the caller.
func LogProcessor(src EventLogSource, registry *Registry) ProcessFunc {
	return func(ctx context.Context, evtID int64) error {
		evtlog, err := src.EventLog(ctx, evtID)
		if err != nil {
			return fmt.Errorf("processing event %v: %w", evtID, err)
		}
		var lg types.Log
		if err := rlp.DecodeBytes(evtlog.RlpLog, &lg); err != nil {
			return fmt.Errorf("processing event %v: RLP decode: %w", evtID, err)
		}
		// The RLP payload carries only the consensus fields; the derived
		// fields ride in the evt_log columns.
		lg.BlockNumber = uint64(evtlog.BlockNum)
		lg.TxHash.SetBytes(common.HexToHash(evtlog.TxHash).Bytes())
		lg.Address.SetBytes(common.HexToHash(evtlog.ContractAddress).Bytes())
		return registry.ProcessLog(ctx, &lg, &evtlog)
	}
}
