// Package dbs - Block and transaction insertion methods
package dbs

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// Insert_block inserts a block header into the database
func (ss *SQLStorage) Insert_block(header *types.Header) error {
	var query string
	query = `INSERT INTO block (
		block_num, block_hash, ts, parent_hash, 
		nonce, miner, difficulty, gas_limit, gas_used,
		tx_root, state_root, receipts_root, uncle_hash
	) VALUES (
		$1, $2, TO_TIMESTAMP($3), $4, 
		$5, $6, $7, $8, $9,
		$10, $11, $12, $13
	) ON CONFLICT (block_num) DO NOTHING`

	_, err := ss.db.Exec(query,
		header.Number.Int64(),
		header.Hash().Hex(),
		header.Time,
		header.ParentHash.Hex(),
		fmt.Sprintf("%d", header.Nonce.Uint64()),
		header.Coinbase.Hex(),
		header.Difficulty.String(),
		header.GasLimit,
		header.GasUsed,
		header.TxHash.Hex(),
		header.Root.Hex(),
		header.ReceiptHash.Hex(),
		header.UncleHash.Hex(),
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Insert_block failed: %v", err))
		return err
	}

	// Update last_block if this is a higher block
	lastBlock, _ := ss.Get_last_block_num()
	if header.Number.Int64() > lastBlock {
		ss.Set_last_block_num(header.Number.Int64())
	}

	return nil
}

// Insert_transaction inserts a transaction into the database
// Returns the transaction ID
func (ss *SQLStorage) Insert_transaction(tx *types.Transaction, blockNum int64, receipt *types.Receipt) (int64, error) {
	// Get sender address
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		// Try legacy signer if chain ID signer fails
		from, err = types.Sender(types.HomesteadSigner{}, tx)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Failed to get sender for tx %s: %v", tx.Hash().Hex(), err))
			return 0, err
		}
	}

	// Lookup or create from address
	fromAid := ss.Lookup_or_create_address(from.Hex(), blockNum, 0)

	// Get to address (may be nil for contract creation)
	var toAddr string
	var toAid int64
	if tx.To() != nil {
		toAddr = tx.To().Hex()
		toAid = ss.Lookup_or_create_address(toAddr, blockNum, 0)
	}

	// Calculate gas price
	var gasPrice *big.Int
	if tx.Type() == types.DynamicFeeTxType {
		gasPrice = tx.GasFeeCap()
	} else {
		gasPrice = tx.GasPrice()
	}

	var query string
	query = `INSERT INTO transaction (
		block_num, tx_hash, tx_index, 
		from_aid, to_aid, value, 
		gas, gas_used, gas_price,
		input, num_logs, ctrct_create
	) VALUES (
		$1, $2, $3, 
		$4, $5, $6, 
		$7, $8, $9,
		$10, $11, $12
	) ON CONFLICT (tx_hash) DO UPDATE SET id = transaction.id
	RETURNING id`

	var txId int64
	err = ss.db.QueryRow(query,
		blockNum,
		tx.Hash().Hex(),
		receipt.TransactionIndex,
		fromAid,
		toAid,
		tx.Value().String(),
		tx.Gas(),
		receipt.GasUsed,
		gasPrice.String(),
		hex.EncodeToString(tx.Data()),
		len(receipt.Logs),
		tx.To() == nil, // contract creation if To is nil
	).Scan(&txId)

	if err != nil {
		ss.Log_msg(fmt.Sprintf("Insert_transaction failed: %v", err))
		return 0, err
	}

	return txId, nil
}

// Insert_event_log inserts an event log into the evt_log table
// Returns the event log ID
func (ss *SQLStorage) Insert_event_log(log types.Log, txId int64, contractAid int64) (int64, error) {
	// Get topic0 signature (first 4 bytes of first topic, or full topic if available)
	var topic0Sig string
	if len(log.Topics) > 0 {
		topic0Sig = log.Topics[0].Hex()[2:] // Remove 0x prefix
	}

	// RLP encode the log
	rlpLog, err := rlp.EncodeToBytes(log)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Failed to RLP encode log: %v", err))
		return 0, err
	}

	var query string
	query = `INSERT INTO evt_log (
		block_num, tx_id, contract_aid, topic0_sig, log_rlp
	) VALUES (
		$1, $2, $3, $4, $5
	) RETURNING id`

	var evtId int64
	err = ss.db.QueryRow(query,
		log.BlockNumber,
		txId,
		contractAid,
		topic0Sig,
		rlpLog,
	).Scan(&evtId)

	if err != nil {
		ss.Log_msg(fmt.Sprintf("Insert_event_log failed: %v", err))
		return 0, err
	}

	return evtId, nil
}
