// Historical evt_log backfill for monitored contracts.

package indexer

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// BackfillStats summarizes a contract evt_log backfill run.
type BackfillStats struct {
	LogsSeen int
	Inserted int
	Skipped  int
}

func (s *BackfillStats) add(other BackfillStats) {
	s.LogsSeen += other.LogsSeen
	s.Inserted += other.Inserted
	s.Skipped += other.Skipped
}

// BackfillContractEvtLogs inserts missing evt_log rows for contract logs in
// [fromBlock, toBlock]. Domain event processors are not invoked; callers use
// this for contracts whose events are stored in evt_log only. Each block is
// committed atomically, and returned statistics describe committed work only.
func (e *Engine) BackfillContractEvtLogs(
	ctx context.Context,
	contracts []common.Address,
	fromBlock, toBlock, batchSize uint64,
) (BackfillStats, error) {
	var st BackfillStats
	if len(contracts) == 0 || fromBlock > toBlock {
		return st, nil
	}
	if batchSize == 0 {
		batchSize = 100_000
	}

	for from := fromBlock; from <= toBlock; {
		to := min(from+batchSize-1, toBlock)

		logs, err := FetchLogs(ctx, e.client, from, to, contracts)
		if err != nil {
			return st, fmt.Errorf("FilterLogs [%d..%d]: %w", from, to, err)
		}

		active := make([]types.Log, 0, len(logs))
		for _, log := range logs {
			if !log.Removed {
				active = append(active, log)
			}
		}
		slices.SortStableFunc(active, func(a, b types.Log) int {
			return cmp.Or(
				cmp.Compare(a.BlockNumber, b.BlockNumber),
				cmp.Compare(a.TxIndex, b.TxIndex),
				cmp.Compare(a.Index, b.Index),
			)
		})

		for first := 0; first < len(active); {
			last := first + 1
			for last < len(active) && active[last].BlockNumber == active[first].BlockNumber {
				last++
			}
			blockStats, err := e.backfillBlock(ctx, active[first:last])
			if err != nil {
				return st, fmt.Errorf("backfill block %d: %w", active[first].BlockNumber, err)
			}
			st.add(blockStats)
			first = last
		}

		from = to + 1
	}

	return st, nil
}

func (e *Engine) backfillBlock(ctx context.Context, logs []types.Log) (BackfillStats, error) {
	if len(logs) == 0 {
		return BackfillStats{}, nil
	}

	blockNum, err := logBlockNum(&logs[0])
	if err != nil {
		return BackfillStats{}, err
	}
	blockHash := logs[0].BlockHash
	for i := range logs {
		currentNum, err := logBlockNum(&logs[i])
		if err != nil {
			return BackfillStats{}, err
		}
		if currentNum != blockNum || logs[i].BlockHash != blockHash {
			return BackfillStats{}, fmt.Errorf(
				"inconsistent block identity at log %d: got block=%d hash=%s, want block=%d hash=%s",
				i,
				currentNum,
				logs[i].BlockHash.Hex(),
				blockNum,
				blockHash.Hex(),
			)
		}
	}

	var blockStats BackfillStats
	err = e.store.InTx(ctx, func(txCtx context.Context) error {
		if _, err := e.EnsureBlockExists(txCtx, blockNum, blockHash.Hex()); err != nil {
			return fmt.Errorf("EnsureBlockExists block=%d: %w", blockNum, err)
		}

		pending := BackfillStats{LogsSeen: len(logs)}
		for i := range logs {
			log := &logs[i]
			txID, err := e.store.TransactionIDByHash(txCtx, log.TxHash.Hex())
			switch {
			case errors.Is(err, store.ErrNotFound):
				txID, _, err = e.EnsureTransactionExists(txCtx, log.TxHash, blockNum)
				if err != nil {
					return fmt.Errorf("EnsureTransactionExists tx=%s: %w", log.TxHash.Hex(), err)
				}
			case err != nil:
				return fmt.Errorf("transaction id lookup tx=%s: %w", log.TxHash.Hex(), err)
			}

			exists, err := e.store.EvtLogExists(txCtx, blockNum, txID, int(log.Index))
			if err != nil {
				return fmt.Errorf(
					"evt_log existence check block=%d tx=%d log_index=%d: %w",
					blockNum,
					txID,
					log.Index,
					err,
				)
			}
			if exists {
				pending.Skipped++
				continue
			}

			if _, err := e.InsertEventLog(txCtx, *log, txID); err != nil {
				return fmt.Errorf(
					"InsertEventLog block=%d tx=%s log_index=%d: %w",
					blockNum,
					log.TxHash.Hex(),
					log.Index,
					err,
				)
			}
			pending.Inserted++
		}
		blockStats = pending
		return nil
	})
	if err != nil {
		return BackfillStats{}, err
	}
	return blockStats, nil
}
