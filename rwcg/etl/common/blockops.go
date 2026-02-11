// Package common - Block and transaction operations
package common

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EnsureBlockExists verifies block exists in DB with correct hash, or inserts it
// If hash mismatch (chain split), it handles the reorg
// Returns: block was newly inserted (true) or already existed (false), error
func EnsureBlockExists(ctx *ETLContext, blockNum int64, expectedHash string) (bool, error) {
	// Check if block exists in DB
	existingHash, err := ctx.Storage.Get_block_hash(blockNum)
	if err != nil {
		// Block doesn't exist, fetch and insert it
		return insertBlockFromChain(ctx, blockNum, expectedHash)
	}

	// Block exists, verify hash matches
	if existingHash == expectedHash {
		// Hash matches, block is valid
		return false, nil
	}

	// Hash mismatch - chain split detected!
	ctx.Info.Printf("Chain split detected at block %d: DB hash %s != event hash %s\n",
		blockNum, existingHash, expectedHash)

	// Handle chain split: delete from this block onwards (from end first)
	err = HandleChainSplit(ctx, blockNum)
	if err != nil {
		return false, fmt.Errorf("HandleChainSplit failed: %v", err)
	}

	// Now insert the correct block
	return insertBlockFromChain(ctx, blockNum, expectedHash)
}

// insertBlockFromChain fetches block header from chain and inserts into DB
func insertBlockFromChain(ctx *ETLContext, blockNum int64, expectedHash string) (bool, error) {
	// Fetch block header from chain
	header, err := GetBlockHeader(ctx.EthClient, blockNum)
	if err != nil {
		return false, fmt.Errorf("GetBlockHeader failed for block %d: %v", blockNum, err)
	}

	// Verify the hash matches what we expect from the event
	actualHash := header.Hash().Hex()
	if actualHash != expectedHash {
		return false, fmt.Errorf("block %d hash mismatch: chain has %s, event has %s",
			blockNum, actualHash, expectedHash)
	}

	// Insert block into DB
	err = ctx.Storage.Insert_block(header)
	if err != nil {
		return false, fmt.Errorf("Insert_block failed for block %d: %v", blockNum, err)
	}

	return true, nil
}

// GetBlockHeader fetches a block header from the chain by number
func GetBlockHeader(client *ethclient.Client, blockNum int64) (*types.Header, error) {
	header, err := client.HeaderByNumber(context.Background(), big.NewInt(blockNum))
	if err != nil {
		return nil, err
	}
	return header, nil
}

// EnsureTransactionExists verifies transaction exists in DB, or fetches and inserts it
// If RPC returns "not found", tries to fetch from archive tables
// Returns: transaction ID, whether it was newly inserted, error
func EnsureTransactionExists(ctx *ETLContext, txHash common.Hash, blockNum int64) (int64, bool, error) {
	// Check if transaction exists in DB
	txId, err := ctx.Storage.Get_transaction_id_by_hash(txHash.Hex())
	if err == nil && txId > 0 {
		// Transaction exists
		return txId, false, nil
	}

	// Try to fetch full transaction from chain
	tx, isPending, err := ctx.EthClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		// Check if it's a "not found" error - only then try archive
		errStr := err.Error()
		if errStr == "not found" || errStr == "transaction not found" {
			// Transaction pruned from RPC node - try archive
			ctx.Info.Printf("[RPC-MISS] Transaction %s not found in RPC, checking archive...\n", txHash.Hex())
			return fetchTransactionFromArchive(ctx, txHash.Hex(), blockNum)
		}
		// Other error (connection refused, etc.) - don't use archive, propagate error
		return 0, false, fmt.Errorf("TransactionByHash failed for %s: %v", txHash.Hex(), err)
	}
	if isPending {
		return 0, false, fmt.Errorf("transaction %s is still pending", txHash.Hex())
	}

	// Get transaction receipt for gas used info
	receipt, err := ctx.EthClient.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		// Check if it's a "not found" error
		errStr := err.Error()
		if errStr == "not found" || errStr == "transaction not found" {
			// Receipt pruned - try archive
			ctx.Info.Printf("[RPC-MISS] Receipt %s not found in RPC, checking archive...\n", txHash.Hex())
			return fetchTransactionFromArchive(ctx, txHash.Hex(), blockNum)
		}
		// Other error - propagate
		return 0, false, fmt.Errorf("TransactionReceipt failed for %s: %v", txHash.Hex(), err)
	}

	// Insert full transaction into DB
	txId, err = ctx.Storage.Insert_transaction(tx, blockNum, receipt)
	if err != nil {
		return 0, false, fmt.Errorf("Insert_transaction failed for %s: %v", txHash.Hex(), err)
	}

	ctx.Info.Printf("[RPC] Transaction %s fetched from RPC node\n", txHash.Hex())
	return txId, true, nil
}

// fetchTransactionFromArchive reads transaction from archive and inserts into main table
func fetchTransactionFromArchive(ctx *ETLContext, txHash string, blockNum int64) (int64, bool, error) {
	// Try to get transaction from archive
	archTx, err := ctx.Storage.Get_archived_transaction(txHash)
	if err != nil {
		return 0, false, fmt.Errorf("archive lookup failed for %s: %v", txHash, err)
	}
	if archTx == nil {
		// Not in archive either - create minimal record as last resort
		ctx.Info.Printf("[MINIMAL] Transaction %s not in RPC or archive, creating minimal record\n", txHash)
		txId, err := ctx.Storage.Insert_minimal_transaction(txHash, blockNum)
		if err != nil {
			return 0, false, fmt.Errorf("Insert_minimal_transaction failed for %s: %v", txHash, err)
		}
		return txId, true, nil
	}

	// Insert transaction from archive data
	ctx.Info.Printf("[ARCHIVE] Transaction %s fetched from archive database\n", txHash)
	txId, err := ctx.Storage.Insert_transaction_from_archive(archTx)
	if err != nil {
		return 0, false, fmt.Errorf("Insert_transaction_from_archive failed for %s: %v", txHash, err)
	}

	return txId, true, nil
}

// InsertEventLog inserts an event log into the evt_log table
// Returns: event log ID, error
func InsertEventLog(ctx *ETLContext, log types.Log, txId int64) (int64, error) {
	// Look up or create contract address
	contractAid := ctx.Storage.Lookup_or_create_address(log.Address.Hex(), int64(log.BlockNumber), txId)

	// Insert the event log
	evtId, err := ctx.Storage.Insert_event_log(log, txId, contractAid)
	if err != nil {
		return 0, fmt.Errorf("Insert_event_log failed: %v", err)
	}

	return evtId, nil
}
