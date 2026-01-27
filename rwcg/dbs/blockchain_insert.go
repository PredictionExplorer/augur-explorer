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
		block_num, block_hash, ts, parent_hash
	) VALUES (
		$1, $2, TO_TIMESTAMP($3), $4
	) ON CONFLICT (block_num) DO NOTHING`

	_, err := ss.db.Exec(query,
		header.Number.Int64(),
		header.Hash().Hex(),
		header.Time,
		header.ParentHash.Hex(),
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

// Insert_minimal_transaction inserts a minimal transaction record when full tx data is unavailable
// Used when RPC node doesn't have historical transaction data (non-archive node)
// Returns the transaction ID
func (ss *SQLStorage) Insert_minimal_transaction(txHash string, blockNum int64) (int64, error) {
	var query string
	query = `INSERT INTO transaction (
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
	RETURNING id`

	var txId int64
	err := ss.db.QueryRow(query, blockNum, txHash).Scan(&txId)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Insert_minimal_transaction failed: %v", err))
		return 0, err
	}

	return txId, nil
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
	var toAid int64
	if tx.To() != nil {
		toAid = ss.Lookup_or_create_address(tx.To().Hex(), blockNum, 0)
	}

	// Calculate gas price
	var gasPrice *big.Int
	if tx.Type() == types.DynamicFeeTxType {
		gasPrice = tx.GasFeeCap()
	} else {
		gasPrice = tx.GasPrice()
	}

	// Get input signature (first 4 bytes of input data)
	var inputSig string
	if len(tx.Data()) >= 4 {
		inputSig = "0x" + hex.EncodeToString(tx.Data()[:4])
	}

	var query string
	query = `INSERT INTO transaction (
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
	RETURNING id`

	var txId int64
	err = ss.db.QueryRow(query,
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
	).Scan(&txId)

	if err != nil {
		ss.Log_msg(fmt.Sprintf("Insert_transaction failed: %v", err))
		return 0, err
	}

	return txId, nil
}

// Insert_event_log inserts an event log into the evt_log table
// Uses INSERT ... ON CONFLICT to handle idempotent re-processing
// Returns the event log ID
func (ss *SQLStorage) Insert_event_log(log types.Log, txId int64, contractAid int64) (int64, error) {

	// Get topic0 signature (first 4 bytes = 8 hex chars of first topic)
	var topic0Sig string
	if len(log.Topics) > 0 {
		fullSig := log.Topics[0].Hex()[2:] // Remove 0x prefix
		if len(fullSig) >= 8 {
			topic0Sig = fullSig[:8] // First 4 bytes (8 hex chars)
		} else {
			topic0Sig = fullSig
		}
	}

	// RLP encode the log (need pointer for EncodeRLP method)
	rlpLog, err := rlp.EncodeToBytes(&log)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Failed to RLP encode log: %v", err))
		return 0, err
	}

	// Use ON CONFLICT to handle idempotent re-processing
	// First try to delete any existing record with same (block_num, log_index)
	// This handles the case where tx_id might have changed due to chain reorganization
	var deleteQuery string
	deleteQuery = `DELETE FROM evt_log WHERE block_num = $1 AND log_index = $2`
	_, err = ss.db.Exec(deleteQuery, log.BlockNumber, log.Index)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Delete before insert failed for evt_log (block=%d, log_index=%d): %v", 
			log.BlockNumber, log.Index, err))
		return 0, err
	}

	var query string
	query = `INSERT INTO evt_log (
		block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp
	) VALUES (
		$1, $2, $3, $4, $5, $6
	) RETURNING id`

	var evtId int64
	err = ss.db.QueryRow(query,
		log.BlockNumber,
		txId,
		contractAid, 
		topic0Sig,
		log.Index, // Log index within the block
		rlpLog,
	).Scan(&evtId)

	if err != nil {
		ss.Log_msg(fmt.Sprintf("Insert_event_log failed: %v", err))
		return 0, err
	}

	return evtId, nil
}
