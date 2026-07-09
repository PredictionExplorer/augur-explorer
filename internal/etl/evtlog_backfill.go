// Package common - Historical evt_log backfill for monitored contracts
package common

import (
	"context"
	"errors"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// EvtLogBackfillStats summarizes a contract evt_log backfill run.
type EvtLogBackfillStats struct {
	LogsSeen int
	Inserted int
	Skipped  int
}

// BackfillContractEvtLogs inserts missing evt_log rows for contract logs in
// [fromBlock, toBlock]. Layer-2 handlers are not invoked; callers use this for
// contracts whose events are stored in evt_log only.
func BackfillContractEvtLogs(
	ctx context.Context,
	etl *ETLContext,
	client *ethclient.Client,
	contracts []ethcommon.Address,
	fromBlock, toBlock, batchSize uint64,
) (EvtLogBackfillStats, error) {
	var st EvtLogBackfillStats
	if len(contracts) == 0 {
		return st, nil
	}
	if fromBlock > toBlock {
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

		logs, err := FetchEvents(ctx, client, from, to, contracts)
		if err != nil {
			return st, fmt.Errorf("FilterLogs [%d..%d]: %w", from, to, err)
		}

		for i := range logs {
			log := &logs[i]
			if log.Removed {
				continue
			}
			st.LogsSeen++

			txId, err := etl.Store.TransactionIDByHash(ctx, log.TxHash.Hex())
			if err != nil && !errors.Is(err, store.ErrNotFound) {
				return st, fmt.Errorf("transaction id lookup tx=%s: %w", log.TxHash.Hex(), err)
			}
			if err == nil && txId > 0 {
				exists, err := etl.Store.EvtLogExists(ctx, int64(log.BlockNumber), txId, int(log.Index))
				if err != nil {
					return st, fmt.Errorf("evt_log existence check block=%d tx=%d log_index=%d: %w",
						log.BlockNumber, txId, log.Index, err)
				}
				if exists {
					st.Skipped++
					continue
				}
			}

			_, err = EnsureBlockExists(ctx, etl, int64(log.BlockNumber), log.BlockHash.Hex())
			if err != nil {
				return st, fmt.Errorf("EnsureBlockExists block=%d: %w", log.BlockNumber, err)
			}

			txId, _, err = EnsureTransactionExists(ctx, etl, log.TxHash, int64(log.BlockNumber))
			if err != nil {
				return st, fmt.Errorf("EnsureTransactionExists tx=%s: %w", log.TxHash.Hex(), err)
			}

			exists, err := etl.Store.EvtLogExists(ctx, int64(log.BlockNumber), txId, int(log.Index))
			if err != nil {
				return st, fmt.Errorf("evt_log existence check block=%d tx=%d log_index=%d: %w",
					log.BlockNumber, txId, log.Index, err)
			}
			if exists {
				st.Skipped++
				continue
			}

			_, err = InsertEventLog(ctx, etl, *log, txId)
			if err != nil {
				return st, fmt.Errorf("InsertEventLog block=%d tx=%s log_index=%d: %w",
					log.BlockNumber, log.TxHash.Hex(), log.Index, err)
			}
			st.Inserted++
		}

		from = to + 1
	}

	return st, nil
}
