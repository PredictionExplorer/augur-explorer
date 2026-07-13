// Package logscan provides a reusable adaptive scanner for inclusive Ethereum
// log ranges.
package logscan

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Filterer is the narrow JSON-RPC surface needed by Scan.
type Filterer interface {
	FilterLogs(context.Context, ethereum.FilterQuery) ([]types.Log, error)
}

// HandleFunc handles one canonical (non-removed) log.
type HandleFunc func(context.Context, types.Log) error

// SleepFunc waits between retries at the minimum batch size.
type SleepFunc func(context.Context, time.Duration) error

// ProgressFunc observes every range immediately before it is fetched.
type ProgressFunc func(context.Context, Progress) error

// Progress describes one FilterLogs attempt.
type Progress struct {
	FromBlock uint64
	ToBlock   uint64
	BatchSize uint64
	Attempt   uint64
}

// Options configures one inclusive range scan.
type Options struct {
	FromBlock uint64
	ToBlock   uint64
	Query     ethereum.FilterQuery

	InitialBatch uint64
	MinBatch     uint64
	RetryDelay   time.Duration

	Sleep      SleepFunc
	OnProgress ProgressFunc
}

// Stats describes completed work and retry activity. BlocksScanned and
// RangesScanned count only fully handled ranges. BlocksScanned saturates at
// math.MaxUint64 for the unrepresentable full [0, math.MaxUint64] domain.
type Stats struct {
	BlocksScanned uint64
	RangesScanned uint64
	FetchAttempts uint64
	FilterErrors  uint64
	LogsSeen      uint64
	RemovedLogs   uint64
}

// Scan fetches and handles all logs in [FromBlock, ToBlock]. A failed fetch is
// retried at the same start block. The batch is halved down to MinBatch; once
// at MinBatch, retries wait RetryDelay. Scan continues until success or context
// cancellation.
func Scan(ctx context.Context, client Filterer, opts Options, handle HandleFunc) (Stats, error) {
	var stats Stats
	if err := validate(client, opts, handle); err != nil {
		return stats, err
	}

	sleep := opts.Sleep
	if sleep == nil {
		sleep = sleepContext
	}

	query := cloneQuery(opts.Query)
	batch := opts.InitialBatch
	from := opts.FromBlock
	var attempt uint64

	for {
		if err := ctx.Err(); err != nil {
			return stats, err
		}

		to := inclusiveEnd(from, opts.ToBlock, batch)
		attempt++
		progress := Progress{
			FromBlock: from,
			ToBlock:   to,
			BatchSize: batch,
			Attempt:   attempt,
		}
		if opts.OnProgress != nil {
			if err := opts.OnProgress(ctx, progress); err != nil {
				return stats, fmt.Errorf("logscan: progress [%d..%d]: %w", from, to, err)
			}
		}
		if err := ctx.Err(); err != nil {
			return stats, err
		}

		request := cloneQuery(query)
		request.FromBlock = new(big.Int).SetUint64(from)
		request.ToBlock = new(big.Int).SetUint64(to)
		stats.FetchAttempts++
		logs, err := client.FilterLogs(ctx, request)
		if err != nil {
			stats.FilterErrors++
			if ctxErr := ctx.Err(); ctxErr != nil {
				return stats, fmt.Errorf("logscan: FilterLogs [%d..%d]: %w", from, to, ctxErr)
			}
			if batch > opts.MinBatch {
				batch = max(opts.MinBatch, batch/2)
				continue
			}
			if err := sleep(ctx, opts.RetryDelay); err != nil {
				return stats, fmt.Errorf("logscan: retry delay after FilterLogs [%d..%d]: %w", from, to, err)
			}
			continue
		}
		if err := ctx.Err(); err != nil {
			return stats, err
		}

		for i := range logs {
			if err := ctx.Err(); err != nil {
				return stats, err
			}
			if logs[i].Removed {
				stats.RemovedLogs++
				continue
			}
			stats.LogsSeen++
			if err := handle(ctx, logs[i]); err != nil {
				return stats, fmt.Errorf(
					"logscan: handle log block=%d index=%d tx=%s: %w",
					logs[i].BlockNumber,
					logs[i].Index,
					logs[i].TxHash.Hex(),
					err,
				)
			}
		}

		stats.RangesScanned++
		stats.BlocksScanned = addBlocksSaturating(stats.BlocksScanned, from, to)
		if to == opts.ToBlock {
			return stats, nil
		}
		from = to + 1
		attempt = 0
	}
}

func validate(client Filterer, opts Options, handle HandleFunc) error {
	switch {
	case client == nil:
		return errors.New("logscan: client is required")
	case handle == nil:
		return errors.New("logscan: handle callback is required")
	case opts.FromBlock > opts.ToBlock:
		return fmt.Errorf("logscan: invalid range %d..%d", opts.FromBlock, opts.ToBlock)
	case opts.InitialBatch == 0:
		return errors.New("logscan: initial batch must be greater than zero")
	case opts.MinBatch == 0:
		return errors.New("logscan: minimum batch must be greater than zero")
	case opts.InitialBatch < opts.MinBatch:
		return fmt.Errorf(
			"logscan: initial batch %d is smaller than minimum batch %d",
			opts.InitialBatch,
			opts.MinBatch,
		)
	case opts.RetryDelay <= 0:
		return errors.New("logscan: retry delay must be greater than zero")
	case opts.Query.BlockHash != nil:
		return errors.New("logscan: block-hash queries cannot be range scanned")
	default:
		return nil
	}
}

func inclusiveEnd(from, end, batch uint64) uint64 {
	if batch-1 >= end-from {
		return end
	}
	return from + batch - 1
}

func addBlocksSaturating(current, from, to uint64) uint64 {
	const maxUint64 = ^uint64(0)
	delta := to - from
	if delta == maxUint64 {
		return maxUint64
	}
	count := delta + 1
	if count > maxUint64-current {
		return maxUint64
	}
	return current + count
}

func cloneQuery(query ethereum.FilterQuery) ethereum.FilterQuery {
	cloned := query
	cloned.Addresses = append(cloned.Addresses[:0:0], query.Addresses...)
	if query.Topics != nil {
		cloned.Topics = make([][]common.Hash, len(query.Topics))
		for i := range query.Topics {
			cloned.Topics[i] = append(cloned.Topics[i], query.Topics[i]...)
		}
	}
	return cloned
}

func sleepContext(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
