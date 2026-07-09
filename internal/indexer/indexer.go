// Package indexer is the shared chain-indexing engine behind cmd/cg-etl and
// cmd/rw-etl. The Engine polls an Ethereum node for contract logs
// (eth_getLogs), persists blocks / transactions / event logs through the
// store, dispatches every stored event to a caller-supplied processor and
// advances the caller's processing watermark — with adaptive batch sizing,
// exponential-backoff retries, a consecutive-failure circuit breaker,
// structured slog logging and Prometheus metrics.
//
// The two ETL binaries differ only in configuration: which contracts to
// filter, how to process one event, and where the watermark lives. Everything
// else — chain-split (reorg) handling, the transaction three-level fallback
// (RPC, archive, minimal record), event-log deduplication, batch/retry policy
// — is engine code, tested once.
package indexer

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Client is the Ethereum JSON-RPC surface the engine consumes.
// *ethclient.Client satisfies it; tests substitute deterministic fakes.
type Client interface {
	// BlockNumber returns the current chain head.
	BlockNumber(ctx context.Context) (uint64, error)
	// HeaderByNumber returns the header of the given block (nil = latest).
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	// FilterLogs executes an eth_getLogs query.
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	// TransactionByHash returns the transaction and whether it is pending.
	TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error)
	// TransactionReceipt returns the receipt of a mined transaction.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// Progress persists the ETL's processing watermark: the last block whose
// events have all been processed. Each binary adapts its domain status row
// (cg_proc_status, rw_proc_status) to this interface.
type Progress interface {
	// LastBlock returns the last fully processed block, 0 when the ETL has
	// never run (the engine then falls back to the store's block watermark).
	LastBlock(ctx context.Context) (int64, error)
	// SetLastBlock records that every event up to and including block has
	// been processed.
	SetLastBlock(ctx context.Context, block int64) error
}

// ProcessFunc consumes one stored event log (identified by its evt_log id):
// it decodes the event and applies the domain writes. A returned error marks
// the batch as failed; the engine retries it, so implementations must stay
// idempotent (the CosmicGame and RandomWalk handlers delete-then-insert or
// skip duplicates by design, pinned by the fixture replay tests).
type ProcessFunc func(ctx context.Context, evtID int64) error

// BatchConfig bounds the adaptive FilterLogs block-range sizing.
// Zero values select the defaults.
type BatchConfig struct {
	// Initial is the first batch size (default 100,000 blocks).
	Initial uint64
	// Min is the size used while events are flowing (default 1,000).
	Min uint64
	// Max caps growth while scanning empty ranges (default 1,000,000).
	Max uint64
}

func (c BatchConfig) withDefaults() BatchConfig {
	if c.Initial == 0 {
		c.Initial = 100_000
	}
	if c.Min == 0 {
		c.Min = 1_000
	}
	if c.Max == 0 {
		c.Max = 1_000_000
	}
	return c
}

// RetryConfig controls failure handling in Run. Zero values select the
// defaults.
type RetryConfig struct {
	// MaxConsecutiveFailures is the circuit breaker: after this many batch
	// failures in a row Run returns the last error and the process exits
	// (systemd restarts it). Default 10.
	MaxConsecutiveFailures int
	// MinDelay is the backoff delay after the first failure (default 1s).
	MinDelay time.Duration
	// MaxDelay caps the exponential growth (default 60s).
	MaxDelay time.Duration
}

func (c RetryConfig) withDefaults() RetryConfig {
	if c.MaxConsecutiveFailures == 0 {
		c.MaxConsecutiveFailures = 10
	}
	if c.MinDelay == 0 {
		c.MinDelay = time.Second
	}
	if c.MaxDelay == 0 {
		c.MaxDelay = time.Minute
	}
	return c
}

// DefaultCaughtUpDelay is how long Run waits for new blocks once the chain
// head has been reached.
const DefaultCaughtUpDelay = 2 * time.Second

// Config assembles an Engine. Store, Client and Logger are required for the
// pipeline methods; Run additionally requires Progress, Process and a
// non-empty Contracts list.
type Config struct {
	// Store owns the database pool the pipeline writes through.
	Store *store.Store
	// Client is the Ethereum node connection.
	Client Client
	// Progress persists the processing watermark (Run only).
	Progress Progress
	// Process handles one stored event (Run only).
	Process ProcessFunc
	// Contracts are the addresses FilterLogs watches (Run only).
	Contracts []common.Address
	// Logger receives structured engine logs; nil discards them.
	Logger *slog.Logger
	// Metrics receives engine metrics; nil disables them.
	Metrics *Metrics
	// TopicName maps a topic0 hash to a human-readable event name for the
	// events_total metric label; nil or "" labels the event "other".
	TopicName func(common.Hash) string
	// Batch tunes the adaptive batch sizing.
	Batch BatchConfig
	// Retry tunes backoff and the circuit breaker.
	Retry RetryConfig
	// CaughtUpDelay overrides DefaultCaughtUpDelay (tests use ~1ms).
	CaughtUpDelay time.Duration
}

// Engine is the indexing pipeline plus its polling loop. Create one with New;
// it is safe to run exactly one Run per Engine.
type Engine struct {
	store     *store.Store
	client    Client
	progress  Progress
	process   ProcessFunc
	contracts []common.Address
	log       *slog.Logger
	metrics   *Metrics
	topicName func(common.Hash) string

	batch         BatchConfig
	retry         RetryConfig
	caughtUpDelay time.Duration
}

// New validates the configuration and builds an Engine.
func New(cfg Config) (*Engine, error) {
	if cfg.Store == nil {
		return nil, fmt.Errorf("indexer: Config.Store is required")
	}
	if cfg.Client == nil {
		return nil, fmt.Errorf("indexer: Config.Client is required")
	}
	logger := cfg.Logger
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}
	caughtUp := cfg.CaughtUpDelay
	if caughtUp == 0 {
		caughtUp = DefaultCaughtUpDelay
	}
	return &Engine{
		store:         cfg.Store,
		client:        cfg.Client,
		progress:      cfg.Progress,
		process:       cfg.Process,
		contracts:     cfg.Contracts,
		log:           logger,
		metrics:       cfg.Metrics,
		topicName:     cfg.TopicName,
		batch:         cfg.Batch.withDefaults(),
		retry:         cfg.Retry.withDefaults(),
		caughtUpDelay: caughtUp,
	}, nil
}

// FetchLogs retrieves the logs emitted by the given contracts in
// [fromBlock, toBlock] via eth_getLogs. It is exposed as a package function
// because operational tools (opsctl) reuse it outside an Engine.
func FetchLogs(ctx context.Context, client Client, fromBlock, toBlock uint64, contracts []common.Address) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: contracts,
	}
	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("FilterLogs failed: %w", err)
	}
	return logs, nil
}
