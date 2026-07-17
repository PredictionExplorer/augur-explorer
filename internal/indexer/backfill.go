// Historical evt_log backfill for monitored contracts.

package indexer

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// BackfillStats summarizes a contract evt_log backfill run.
type BackfillStats struct {
	LogsSeen int
	Inserted int
	Skipped  int
}

// BackfillContractEvtLogs inserts missing evt_log rows for contract logs in
// [fromBlock, toBlock]. Domain event processors are not invoked; callers use
// this for contracts whose events are stored in evt_log only.
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
		to := from + batchSize - 1
		if to > toBlock {
			to = toBlock
		}

		logs, err := FetchLogs(ctx, e.client, from, to, contracts)
		if err != nil {
			return st, fmt.Errorf("FilterLogs [%d..%d]: %w", from, to, err)
		}

		for i := range logs {
			log := &logs[i]
			if log.Removed {
				continue
			}
			st.LogsSeen++

			blockNum, err := logBlockNum(log)
			if err != nil {
				return st, err
			}

			txID, err := e.store.TransactionIDByHash(ctx, log.TxHash.Hex())
			if err != nil && !errors.Is(err, store.ErrNotFound) {
				return st, fmt.Errorf("transaction id lookup tx=%s: %w", log.TxHash.Hex(), err)
			}
			if err == nil && txID > 0 {
				exists, err := e.store.EvtLogExists(ctx, blockNum, txID, int(log.Index))
				if err != nil {
					return st, fmt.Errorf("evt_log existence check block=%d tx=%d log_index=%d: %w",
						log.BlockNumber, txID, log.Index, err)
				}
				if exists {
					st.Skipped++
					continue
				}
			}

			if _, err := e.EnsureBlockExists(ctx, blockNum, log.BlockHash.Hex()); err != nil {
				return st, fmt.Errorf("EnsureBlockExists block=%d: %w", log.BlockNumber, err)
			}

			txID, _, err = e.EnsureTransactionExists(ctx, log.TxHash, blockNum)
			if err != nil {
				return st, fmt.Errorf("EnsureTransactionExists tx=%s: %w", log.TxHash.Hex(), err)
			}

			exists, err := e.store.EvtLogExists(ctx, blockNum, txID, int(log.Index))
			if err != nil {
				return st, fmt.Errorf("evt_log existence check block=%d tx=%d log_index=%d: %w",
					log.BlockNumber, txID, log.Index, err)
			}
			if exists {
				st.Skipped++
				continue
			}

			if _, err := e.InsertEventLog(ctx, *log, txID); err != nil {
				return st, fmt.Errorf("InsertEventLog block=%d tx=%s log_index=%d: %w",
					log.BlockNumber, log.TxHash.Hex(), log.Index, err)
			}
			st.Inserted++
		}

		from = to + 1
	}

	return st, nil
}
