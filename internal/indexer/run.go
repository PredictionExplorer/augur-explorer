// The polling loop: fetch a block range of contract logs, pipe every log
// through the storage pipeline and the processor, advance the watermark,
// adapt the batch size, and retry failures with backoff until the circuit
// breaker trips.

package indexer

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

// Run polls the chain until ctx is canceled (returns nil — user-requested
// shutdown) or a batch keeps failing MaxConsecutiveFailures times in a row
// (returns the last error — the process exits non-zero and the supervisor
// restarts it).
//
// Shutdown semantics: cancellation is honored between batches and during
// backoff/caught-up waits. The database work of an in-flight batch runs on a
// context that inherits values but not cancellation, so a SIGTERM arriving
// mid-batch still gets the promised "finish batch, write status, exit 0".
//
// Failure semantics: any failed batch is retried from the last fully
// completed block with exponential backoff. The watermark only ever advances
// past blocks whose events were all processed, so a mid-block failure
// re-fetches that whole block; processors stay idempotent under that replay.
func (e *Engine) Run(ctx context.Context) error {
	if e.progress == nil || e.process == nil {
		return fmt.Errorf("indexer: Run requires Config.Progress and Config.Process")
	}
	if len(e.contracts) == 0 {
		return fmt.Errorf("indexer: Run requires a non-empty Config.Contracts (an empty FilterLogs address list would match every contract)")
	}

	// In-flight batch DB work must complete even when shutdown has been
	// signaled; it runs on a context that inherits values but not
	// cancellation. Cancellation still applies to chain RPC reads.
	dbCtx := context.WithoutCancel(ctx)

	batch := newBatchPolicy(e.batch)
	failures := 0
	e.log.Info("indexer started",
		"contracts", len(e.contracts),
		"batch_initial", batch.size, "batch_min", batch.min, "batch_max", batch.max,
		"max_consecutive_failures", e.retry.MaxConsecutiveFailures)

	// The watermark is resolved once at startup and tracked in memory from
	// then on. Re-reading it every iteration (the legacy loops did) is a
	// trap on a fresh status row: a failed batch leaves its inserted blocks
	// behind, advancing the store's block watermark, and the 0-status
	// fallback would resume *past* the events the batch still owes.
	lastProcessed := int64(-1)

	for {
		if ctx.Err() != nil {
			e.log.Info("exiting by user request")
			return nil
		}

		if lastProcessed < 0 {
			// Startup watermark read. A failure here is a database failure:
			// retry with backoff, crash only when persistent.
			lp, err := e.lastProcessedBlock(dbCtx)
			if err != nil {
				if stop := e.batchFailure(ctx, &failures, "watermark", err); stop != nil {
					return stop
				}
				continue
			}
			lastProcessed = lp
			e.log.Info("resuming after last processed block", "block", lastProcessed)
		}

		head, err := e.client.BlockNumber(ctx)
		if err != nil {
			if ctx.Err() != nil {
				e.log.Info("exiting by user request")
				return nil
			}
			if stop := e.batchFailure(ctx, &failures, "chain_head", err); stop != nil {
				return stop
			}
			continue
		}

		fromBlock := uint64(lastProcessed + 1)
		if fromBlock > head {
			// Caught up: wait for new blocks with small real-time batches.
			batch.onCaughtUp()
			if !e.sleep(ctx, e.caughtUpDelay) {
				e.log.Info("exiting by user request")
				return nil
			}
			continue
		}
		toBlock := fromBlock + batch.size - 1
		if toBlock > head {
			toBlock = head
		}

		started := time.Now()
		logs, err := FetchLogs(ctx, e.client, fromBlock, toBlock, e.contracts)
		if err != nil {
			if ctx.Err() != nil {
				e.log.Info("exiting by user request")
				return nil
			}
			// Oversized ranges are the usual cause: shrink and retry.
			batch.onFetchError()
			if stop := e.batchFailure(ctx, &failures, "fetch", err); stop != nil {
				return stop
			}
			continue
		}
		e.log.Info("fetched events",
			"from_block", fromBlock, "to_block", toBlock,
			"batch_size", batch.size, "events", len(logs))

		if lastCompleted, stage, err := e.processBatch(dbCtx, logs); err != nil {
			// Acknowledge the blocks that completed before the failing one,
			// then retry the rest of the batch. Never acknowledge the
			// failing block itself: a block boundary is the only safe
			// watermark (advancing to the failing log's own block would
			// permanently skip that block's remaining logs).
			if lastCompleted > lastProcessed {
				if ackErr := e.setLastBlock(dbCtx, lastCompleted); ackErr != nil {
					e.log.Error("partial watermark write failed", "block", lastCompleted, "err", ackErr)
				} else {
					lastProcessed = lastCompleted
				}
			}
			if stop := e.batchFailure(ctx, &failures, stage, err); stop != nil {
				return stop
			}
			continue
		}

		if err := e.setLastBlock(dbCtx, int64(toBlock)); err != nil {
			if stop := e.batchFailure(ctx, &failures, "watermark", err); stop != nil {
				return stop
			}
			continue
		}
		lastProcessed = int64(toBlock)

		failures = 0
		e.metrics.batchProcessed(time.Since(started).Seconds())
		if len(logs) == 0 {
			batch.onEmpty()
		} else {
			batch.onEvents()
		}
	}
}

// processBatch pipes every log through the storage pipeline and the
// processor. On failure it returns the last block whose logs all completed
// (0 when none did), the failed pipeline stage and the error.
func (e *Engine) processBatch(ctx context.Context, logs []types.Log) (lastCompleted int64, stage string, err error) {
	var currentBlock int64
	for i := range logs {
		log := logs[i]
		blockNum := int64(log.BlockNumber)
		if currentBlock != 0 && blockNum > currentBlock {
			lastCompleted = currentBlock
		}
		currentBlock = blockNum

		if _, err := e.EnsureBlockExists(ctx, blockNum, log.BlockHash.Hex()); err != nil {
			return lastCompleted, "block", err
		}

		txID, _, err := e.EnsureTransactionExists(ctx, log.TxHash, blockNum)
		if err != nil {
			return lastCompleted, "transaction", err
		}

		evtID, err := e.InsertEventLog(ctx, log, txID)
		if err != nil {
			return lastCompleted, "event_log", err
		}

		if err := e.process(ctx, evtID); err != nil {
			return lastCompleted, "process", fmt.Errorf("processing event %d: %w", evtID, err)
		}
		e.metrics.eventProcessed(e.eventTypeLabel(&log))
	}
	return 0, "", nil
}

// lastProcessedBlock reads the processing watermark, falling back to the
// store's block watermark when the status row reports 0 (an ETL that has
// never run resumes from the last block the database knows).
func (e *Engine) lastProcessedBlock(ctx context.Context) (int64, error) {
	last, err := e.progress.LastBlock(ctx)
	if err != nil {
		return 0, fmt.Errorf("reading processing status: %w", err)
	}
	if last != 0 {
		return last, nil
	}
	last, err = e.store.LastBlockNum(ctx)
	if err != nil {
		return 0, fmt.Errorf("reading last block watermark: %w", err)
	}
	return last, nil
}

// setLastBlock persists the watermark and mirrors it to the gauge.
func (e *Engine) setLastBlock(ctx context.Context, block int64) error {
	if err := e.progress.SetLastBlock(ctx, block); err != nil {
		return fmt.Errorf("updating processing status: %w", err)
	}
	e.metrics.watermark(block)
	return nil
}

// batchFailure records one failed batch attempt: it increments the
// consecutive-failure count, trips the circuit breaker when the limit is
// reached (non-nil return: Run exits with it) and otherwise sleeps the
// backoff delay.
func (e *Engine) batchFailure(ctx context.Context, failures *int, stage string, err error) error {
	*failures++
	e.metrics.batchFailed(stage)
	delay := backoffDelay(*failures, e.retry.MinDelay, e.retry.MaxDelay, randFloat)
	e.log.Error("batch failed",
		"stage", stage,
		"consecutive_failures", *failures,
		"max_consecutive_failures", e.retry.MaxConsecutiveFailures,
		"retry_in", delay,
		"err", err)
	if *failures >= e.retry.MaxConsecutiveFailures {
		return fmt.Errorf("giving up after %d consecutive batch failures (stage %s): %w",
			*failures, stage, err)
	}
	e.sleep(ctx, delay)
	return nil
}

// sleep waits d unless ctx is canceled first; it reports whether the full
// wait elapsed.
func (e *Engine) sleep(ctx context.Context, d time.Duration) bool {
	if d <= 0 {
		return ctx.Err() == nil
	}
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}

// eventTypeLabel resolves the metric label of one processed log.
func (e *Engine) eventTypeLabel(log *types.Log) string {
	if len(log.Topics) == 0 {
		return "none"
	}
	if e.topicName != nil {
		if name := e.topicName(log.Topics[0]); name != "" {
			return name
		}
	}
	return "other"
}
