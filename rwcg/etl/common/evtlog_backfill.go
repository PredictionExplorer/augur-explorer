// Package common - Historical evt_log backfill for monitored contracts
package common

import (
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	ctx *ETLContext,
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

		logs, err := FetchEvents(client, from, to, contracts)
		if err != nil {
			return st, fmt.Errorf("FilterLogs [%d..%d]: %w", from, to, err)
		}

		for i := range logs {
			log := &logs[i]
			if log.Removed {
				continue
			}
			st.LogsSeen++

			txId, err := ctx.Storage.Get_transaction_id_by_hash(log.TxHash.Hex())
			if err == nil && txId > 0 {
				exists, err := ctx.Storage.Evt_log_exists(int64(log.BlockNumber), txId, int(log.Index))
				if err != nil {
					return st, fmt.Errorf("Evt_log_exists block=%d tx=%d log_index=%d: %w",
						log.BlockNumber, txId, log.Index, err)
				}
				if exists {
					st.Skipped++
					continue
				}
			}

			_, err = EnsureBlockExists(ctx, int64(log.BlockNumber), log.BlockHash.Hex())
			if err != nil {
				return st, fmt.Errorf("EnsureBlockExists block=%d: %w", log.BlockNumber, err)
			}

			txId, _, err = EnsureTransactionExists(ctx, log.TxHash, int64(log.BlockNumber))
			if err != nil {
				return st, fmt.Errorf("EnsureTransactionExists tx=%s: %w", log.TxHash.Hex(), err)
			}

			exists, err := ctx.Storage.Evt_log_exists(int64(log.BlockNumber), txId, int(log.Index))
			if err != nil {
				return st, fmt.Errorf("Evt_log_exists block=%d tx=%d log_index=%d: %w",
					log.BlockNumber, txId, log.Index, err)
			}
			if exists {
				st.Skipped++
				continue
			}

			_, err = InsertEventLog(ctx, *log, txId)
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
