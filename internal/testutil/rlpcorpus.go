package testutil

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/PredictionExplorer/augur-explorer/internal/rlpcorpus"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// InstallRLPCorpus atomically inserts archive-compatible layer-1 scaffolding
// and the corpus's exact LogRLP bytes into a test Store. It deliberately
// bypasses Store.InsertEventLog's re-encoding so integration tests exercise
// exported bytes. Returned evt_log IDs preserve corpus order; callers dispatch
// them only after every sibling log has been installed and committed.
func InstallRLPCorpus(ctx context.Context, st *store.Store, entries []rlpcorpus.Entry) ([]int64, error) {
	if st == nil || st.Pool() == nil {
		return nil, errors.New("RLP corpus Store is required")
	}
	if err := rlpcorpus.Write(io.Discard, entries); err != nil {
		return nil, err
	}

	var ids []int64
	err := st.InTx(ctx, func(txCtx context.Context) error {
		pending := make([]int64, 0, len(entries))
		for i, entry := range entries {
			if err := txCtx.Err(); err != nil {
				return err
			}
			blockHash := crypto.Keccak256Hash(fmt.Appendf(nil, "rlp-corpus:block:%d", entry.BlockNum))
			parentHash := crypto.Keccak256Hash(fmt.Appendf(nil, "rlp-corpus:block:%d", entry.BlockNum-1))
			if _, err := st.Querier(txCtx).Exec(txCtx, `INSERT INTO block (
					block_num, ts, block_hash, parent_hash
				) VALUES ($1, TO_TIMESTAMP($2), $3, $4)
				ON CONFLICT (block_num) DO NOTHING`,
				entry.BlockNum,
				entry.BlockNum,
				blockHash.Hex(),
				parentHash.Hex(),
			); err != nil {
				return fmt.Errorf("install RLP corpus entry %d block: %w", i, err)
			}
			txID, err := st.InsertMinimalTransaction(txCtx, entry.TxHash, entry.BlockNum)
			if err != nil {
				return fmt.Errorf("install RLP corpus entry %d transaction: %w", i, err)
			}
			contractAID, err := st.LookupOrCreateAddress(
				txCtx,
				common.HexToAddress(entry.ContractAddress).Hex(),
				entry.BlockNum,
				txID,
			)
			if err != nil {
				return fmt.Errorf("install RLP corpus entry %d contract: %w", i, err)
			}
			raw, err := entry.RLPBytes()
			if err != nil {
				return fmt.Errorf("install RLP corpus entry %d: %w", i, err)
			}
			var evtID int64
			err = st.Querier(txCtx).QueryRow(txCtx, `INSERT INTO evt_log (
					block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp
				) VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id`,
				entry.BlockNum,
				txID,
				contractAID,
				entry.Topic0Sig,
				entry.LogIndex,
				raw,
			).Scan(&evtID)
			if err != nil {
				return fmt.Errorf("install RLP corpus entry %d event log: %w", i, err)
			}
			pending = append(pending, evtID)
		}
		ids = pending
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ids, nil
}
