// Package cstscan scans CstDutchAuctionDurationChangeDivisorChanged events and
// optionally cross-checks them against PostgreSQL-derived keys.
package cstscan

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer/logscan"
)

// Client is the narrow Ethereum JSON-RPC surface used by Scan.
type Client interface {
	logscan.Filterer
	BlockNumber(context.Context) (uint64, error)
}

// Logger receives progress and summary lines. *log.Logger satisfies it.
type Logger interface {
	Printf(string, ...any)
}

// EventKey identifies one chain event in the database cross-check.
type EventKey struct {
	TxHash   common.Hash
	LogIndex uint64
}

// KeySource loads event keys already stored in the application database.
type KeySource interface {
	LoadKeys(context.Context) (map[EventKey]struct{}, error)
}

// KeySourceFunc adapts a function to KeySource.
type KeySourceFunc func(context.Context) (map[EventKey]struct{}, error)

// LoadKeys implements KeySource.
func (f KeySourceFunc) LoadKeys(ctx context.Context) (map[EventKey]struct{}, error) {
	return f(ctx)
}

// Querier is the narrow pgx query surface used by PostgresKeySource.
// *pgxpool.Pool, *pgx.Conn and pgx.Tx satisfy it.
type Querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

// PostgresKeySource loads keys from cg_adm_cst_auclen_chg_div. The caller owns
// opening and closing DB.
type PostgresKeySource struct {
	DB Querier
}

// LoadKeys implements KeySource.
func (source PostgresKeySource) LoadKeys(ctx context.Context) (map[EventKey]struct{}, error) {
	if source.DB == nil {
		return nil, errors.New("cstscan: database is required")
	}
	const query = `
		SELECT lower(t.tx_hash), e.log_index
		FROM cg_adm_cst_auclen_chg_div r
		JOIN evt_log e ON e.id = r.evtlog_id
		JOIN transaction t ON t.id = e.tx_id`
	rows, err := source.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	keys := make(map[EventKey]struct{})
	for rows.Next() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		var (
			txHash   string
			logIndex int64
		)
		if err := rows.Scan(&txHash, &logIndex); err != nil {
			return nil, fmt.Errorf("db scan: %w", err)
		}
		if logIndex < 0 {
			return nil, fmt.Errorf("db scan: negative log index %d", logIndex)
		}
		keys[EventKey{
			TxHash:   common.HexToHash(txHash),
			LogIndex: uint64(logIndex),
		}] = struct{}{}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return keys, nil
}

// Config contains Scan's stable dependencies and event filter.
type Config struct {
	Client    Client
	Contract  common.Address
	Topic0    common.Hash
	KeySource KeySource
	Output    io.Writer
	Logger    Logger
}

// Options bounds the inclusive scan. ToBlock zero selects the latest block.
type Options struct {
	FromBlock    uint64
	ToBlock      uint64
	InitialBatch uint64
	MinBatch     uint64
	RetryDelay   time.Duration

	Sleep      logscan.SleepFunc
	OnProgress logscan.ProgressFunc
}

// Stats summarizes one completed or interrupted scan.
type Stats struct {
	FromBlock     uint64
	ToBlock       uint64
	BlocksScanned uint64
	Events        uint64
	InDB          uint64
	MissingFromDB uint64
	FilterErrors  uint64
	RemovedLogs   uint64
}

// DecodeAuctionLength decodes the event's single non-indexed uint256 value.
func DecodeAuctionLength(data []byte) (*big.Int, error) {
	if len(data) != common.HashLength {
		return nil, fmt.Errorf("cstscan: auction length data is %d bytes, want 32", len(data))
	}
	return new(big.Int).SetBytes(data), nil
}

// Scan prints matching events as tab-separated rows and optionally annotates
// whether each event exists in the injected key source.
func Scan(ctx context.Context, cfg Config, opts Options) (Stats, error) {
	stats := Stats{FromBlock: opts.FromBlock}
	if cfg.Client == nil {
		return stats, errors.New("cstscan: client is required")
	}
	if cfg.Output == nil {
		return stats, errors.New("cstscan: output writer is required")
	}
	if err := validateOptions(opts); err != nil {
		return stats, err
	}
	if err := ctx.Err(); err != nil {
		return stats, err
	}

	end := opts.ToBlock
	if end == 0 {
		var err error
		end, err = cfg.Client.BlockNumber(ctx)
		if err != nil {
			return stats, fmt.Errorf("cstscan: BlockNumber: %w", err)
		}
	}
	stats.ToBlock = end
	if opts.FromBlock > end {
		return stats, fmt.Errorf("cstscan: invalid range %d..%d", opts.FromBlock, end)
	}

	crossCheck := cfg.KeySource != nil
	var dbKeys map[EventKey]struct{}
	if crossCheck {
		var err error
		dbKeys, err = cfg.KeySource.LoadKeys(ctx)
		if err != nil {
			return stats, fmt.Errorf("cstscan: load database keys: %w", err)
		}
		if dbKeys == nil {
			dbKeys = make(map[EventKey]struct{})
		}
		if err := ctx.Err(); err != nil {
			return stats, err
		}
		logf(cfg.Logger, "Loaded %d existing rows from cg_adm_cst_auclen_chg_div", len(dbKeys))
	}

	logf(
		cfg.Logger,
		"Scanning %s for topic %s, blocks %d..%d",
		cfg.Contract.Hex(),
		cfg.Topic0.Hex(),
		opts.FromBlock,
		end,
	)
	header := "block_num\ttx_hash\tlog_index\tnew_len\n"
	if crossCheck {
		header = "block_num\ttx_hash\tlog_index\tnew_len\tin_db\n"
	}
	if _, err := io.WriteString(cfg.Output, header); err != nil {
		return stats, fmt.Errorf("cstscan: write header: %w", err)
	}

	lastBatch := opts.InitialBatch
	progress := func(ctx context.Context, progress logscan.Progress) error {
		if progress.BatchSize < lastBatch {
			logf(cfg.Logger, "Reducing batch to %d blocks", progress.BatchSize)
		}
		lastBatch = progress.BatchSize
		if opts.OnProgress != nil {
			return opts.OnProgress(ctx, progress)
		}
		return nil
	}

	scanStats, err := logscan.Scan(
		ctx,
		loggingFilterer{client: cfg.Client, logger: cfg.Logger},
		logscan.Options{
			FromBlock: opts.FromBlock,
			ToBlock:   end,
			Query: ethereum.FilterQuery{
				Addresses: []common.Address{cfg.Contract},
				Topics:    [][]common.Hash{{cfg.Topic0}},
			},
			InitialBatch: opts.InitialBatch,
			MinBatch:     opts.MinBatch,
			RetryDelay:   opts.RetryDelay,
			Sleep:        opts.Sleep,
			OnProgress:   progress,
		},
		func(_ context.Context, log types.Log) error {
			value, err := DecodeAuctionLength(log.Data)
			if err != nil {
				return err
			}
			line := fmt.Sprintf(
				"%d\t%s\t%d\t%s",
				log.BlockNumber,
				log.TxHash.Hex(),
				log.Index,
				value.String(),
			)
			stats.Events++
			if crossCheck {
				if _, exists := dbKeys[EventKey{TxHash: log.TxHash, LogIndex: uint64(log.Index)}]; exists {
					line += "\tyes"
					stats.InDB++
				} else {
					line += "\tNO"
					stats.MissingFromDB++
				}
			}
			if _, err := fmt.Fprintln(cfg.Output, line); err != nil {
				return fmt.Errorf("write event: %w", err)
			}
			return nil
		},
	)
	stats.BlocksScanned = scanStats.BlocksScanned
	stats.FilterErrors = scanStats.FilterErrors
	stats.RemovedLogs = scanStats.RemovedLogs
	if err != nil {
		return stats, err
	}

	if crossCheck {
		logf(
			cfg.Logger,
			"Done. on_chain_events=%d in_db=%d MISSING_FROM_DB=%d",
			stats.Events,
			stats.InDB,
			stats.MissingFromDB,
		)
	} else {
		logf(
			cfg.Logger,
			"Done. on_chain_events=%d blocks=%d..%d",
			stats.Events,
			opts.FromBlock,
			end,
		)
	}
	return stats, nil
}

func validateOptions(opts Options) error {
	switch {
	case opts.InitialBatch == 0:
		return errors.New("cstscan: initial batch must be greater than zero")
	case opts.MinBatch == 0:
		return errors.New("cstscan: minimum batch must be greater than zero")
	case opts.InitialBatch < opts.MinBatch:
		return fmt.Errorf(
			"cstscan: initial batch %d is smaller than minimum batch %d",
			opts.InitialBatch,
			opts.MinBatch,
		)
	case opts.RetryDelay <= 0:
		return errors.New("cstscan: retry delay must be greater than zero")
	default:
		return nil
	}
}

type loggingFilterer struct {
	client logscan.Filterer
	logger Logger
}

func (filterer loggingFilterer) FilterLogs(
	ctx context.Context,
	query ethereum.FilterQuery,
) ([]types.Log, error) {
	logs, err := filterer.client.FilterLogs(ctx, query)
	if err != nil {
		logf(
			filterer.logger,
			"FilterLogs error [%d..%d]: %v",
			query.FromBlock.Uint64(),
			query.ToBlock.Uint64(),
			err,
		)
		return nil, err
	}
	logf(
		filterer.logger,
		"scanned %d..%d (%d events)",
		query.FromBlock.Uint64(),
		query.ToBlock.Uint64(),
		len(logs),
	)
	return logs, nil
}

func logf(logger Logger, format string, args ...any) {
	if logger != nil {
		logger.Printf(format, args...)
	}
}
