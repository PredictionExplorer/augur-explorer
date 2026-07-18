// Block, transaction and event-log insertion for the ETL pipeline.

package store

import (
	"context"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// InsertBlock inserts a block header (idempotent on block_num) and advances
// the last_block watermark when the block is higher than the stored one.
func (s *Store) InsertBlock(ctx context.Context, header *types.Header) error {
	blockNum := header.Number.Int64()
	op := fmt.Sprintf("insert block %v", blockNum)
	_, err := s.q(ctx).Exec(ctx, `INSERT INTO block (
			block_num, block_hash, ts, parent_hash
		) VALUES (
			$1, $2, TO_TIMESTAMP($3), $4
		) ON CONFLICT (block_num) DO NOTHING`,
		blockNum,
		header.Hash().Hex(),
		header.Time,
		header.ParentHash.Hex(),
	)
	if err != nil {
		return WrapError(op, err)
	}

	lastBlock, err := s.LastBlockNum(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if blockNum > lastBlock {
		if err := s.SetLastBlockNum(ctx, blockNum); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}
	return nil
}

// InsertMinimalTransaction inserts a placeholder transaction record when the
// full data is unavailable (RPC node without historical transaction data and
// no archive entry). Returns the transaction id.
func (s *Store) InsertMinimalTransaction(ctx context.Context, txHash string, blockNum int64) (int64, error) {
	var txID int64
	err := s.q(ctx).QueryRow(ctx, `INSERT INTO transaction (
			block_num, tx_hash, tx_index,
			from_aid, to_aid, value,
			gas_used, gas_price,
			input_sig, num_logs, ctrct_create
		) VALUES (
			$1, $2, 0,
			0, 0, '0',
			0, '0',
			'', 0, false
		) ON CONFLICT (tx_hash) DO UPDATE SET id = transaction.id
		RETURNING id`, blockNum, txHash).Scan(&txID)
	if err != nil {
		return 0, WrapError(fmt.Sprintf("insert minimal transaction %v", txHash), err)
	}
	return txID, nil
}

// InsertTransaction inserts a full transaction record (idempotent on
// tx_hash), resolving the from/to addresses through the address cache.
// Returns the transaction id.
func (s *Store) InsertTransaction(ctx context.Context, tx *types.Transaction, blockNum int64, receipt *types.Receipt) (int64, error) {
	op := fmt.Sprintf("insert transaction %v", tx.Hash().Hex())

	// Recover the sender; pre-EIP-155 transactions need the Homestead signer.
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		from, err = types.Sender(types.HomesteadSigner{}, tx)
		if err != nil {
			return 0, fmt.Errorf("%s: recover sender: %w", op, err)
		}
	}

	fromAid, err := s.LookupOrCreateAddress(ctx, from.Hex(), blockNum, 0)
	if err != nil {
		return 0, fmt.Errorf("%s: from address: %w", op, err)
	}

	// To is nil for contract creation.
	var toAid int64
	if tx.To() != nil {
		toAid, err = s.LookupOrCreateAddress(ctx, tx.To().Hex(), blockNum, 0)
		if err != nil {
			return 0, fmt.Errorf("%s: to address: %w", op, err)
		}
	}

	var gasPrice *big.Int
	if tx.Type() == types.DynamicFeeTxType {
		gasPrice = tx.GasFeeCap()
	} else {
		gasPrice = tx.GasPrice()
	}

	// Input signature: first 4 bytes of the call data.
	var inputSig string
	if len(tx.Data()) >= 4 {
		inputSig = "0x" + hex.EncodeToString(tx.Data()[:4])
	}

	var txID int64
	err = s.q(ctx).QueryRow(ctx, `INSERT INTO transaction (
			block_num, tx_hash, tx_index,
			from_aid, to_aid, value,
			gas_used, gas_price,
			input_sig, num_logs, ctrct_create
		) VALUES (
			$1, $2, $3,
			$4, $5, $6,
			$7, $8,
			$9, $10, $11
		) ON CONFLICT (tx_hash) DO UPDATE SET id = transaction.id
		RETURNING id`,
		blockNum,
		tx.Hash().Hex(),
		receipt.TransactionIndex,
		fromAid,
		toAid,
		tx.Value().String(),
		receipt.GasUsed,
		gasPrice.String(),
		inputSig,
		len(receipt.Logs),
		tx.To() == nil, // contract creation if To is nil
	).Scan(&txID)
	if err != nil {
		return 0, WrapError(op, err)
	}
	return txID, nil
}

// NextEventLogIndex returns the first unused log index at or above minimum
// for blockNum. Synthetic events use this to avoid deleting an earlier event
// through InsertEventLog's delete-before-insert replay semantics.
func (s *Store) NextEventLogIndex(ctx context.Context, blockNum int64, minimum uint) (uint, error) {
	if blockNum < 0 || uint64(minimum) > math.MaxInt32 {
		return 0, fmt.Errorf("next event log index: invalid block or minimum")
	}
	var next int64
	err := s.q(ctx).QueryRow(ctx, `SELECT GREATEST(
			COALESCE(MAX(log_index)::BIGINT + 1, $2::BIGINT),
			$2::BIGINT
		)
		FROM evt_log
		WHERE block_num=$1`, blockNum, minimum).Scan(&next)
	if err != nil {
		return 0, WrapError("next event log index", err)
	}
	if next < 0 || next > math.MaxInt32 {
		return 0, fmt.Errorf("next event log index: exhausted int32 range")
	}
	return uint(next), nil
}

// InsertEventLog stores an Ethereum log in evt_log (RLP-encoded payload
// included) and returns the new row id. Any existing row with the same
// (block_num, log_index) is replaced first, which makes re-processing after
// a chain reorganization idempotent even when the tx id changed.
func (s *Store) InsertEventLog(ctx context.Context, log types.Log, txID, contractAid int64) (int64, error) {
	op := fmt.Sprintf("insert event log block=%d log_index=%d", log.BlockNumber, log.Index)

	// topic0 signature: first 4 bytes (8 hex chars) of the first topic.
	var topic0Sig string
	if len(log.Topics) > 0 {
		fullSig := log.Topics[0].Hex()[2:]
		if len(fullSig) >= 8 {
			topic0Sig = fullSig[:8]
		} else {
			topic0Sig = fullSig
		}
	}

	rlpLog, err := rlp.EncodeToBytes(&log)
	if err != nil {
		return 0, fmt.Errorf("%s: RLP encode: %w", op, err)
	}

	if _, err := s.q(ctx).Exec(ctx,
		"DELETE FROM evt_log WHERE block_num = $1 AND log_index = $2",
		log.BlockNumber, log.Index); err != nil {
		return 0, WrapError(op+": delete before insert", err)
	}

	var evtID int64
	err = s.q(ctx).QueryRow(ctx, `INSERT INTO evt_log (
			block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp
		) VALUES (
			$1, $2, $3, $4, $5, $6
		) RETURNING id`,
		log.BlockNumber,
		txID,
		contractAid,
		topic0Sig,
		log.Index,
		rlpLog,
	).Scan(&evtID)
	if err != nil {
		return 0, WrapError(op, err)
	}
	return evtID, nil
}
