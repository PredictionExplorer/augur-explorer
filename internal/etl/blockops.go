// Package common - Block and transaction operations
package common

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// EnsureBlockExists verifies block exists in DB with correct hash, or inserts it
// If hash mismatch (chain split), it handles the reorg
// Returns: block was newly inserted (true) or already existed (false), error
func EnsureBlockExists(ctx context.Context, etl *ETLContext, blockNum int64, expectedHash string) (bool, error) {
	// Check if block exists in DB
	existingHash, err := etl.Store.BlockHash(ctx, blockNum)
	if errors.Is(err, store.ErrNotFound) {
		// Block doesn't exist, fetch and insert it
		return insertBlockFromChain(ctx, etl, blockNum, expectedHash)
	}
	if err != nil {
		return false, fmt.Errorf("block hash check failed for block %d: %w", blockNum, err)
	}

	// Block exists, verify hash matches
	if existingHash == expectedHash {
		// Hash matches, block is valid
		return false, nil
	}

	// Hash mismatch - chain split detected!
	etl.Info.Printf("Chain split detected at block %d: DB hash %s != event hash %s\n",
		blockNum, existingHash, expectedHash)

	// Handle chain split: delete from this block onwards (from end first)
	err = HandleChainSplit(ctx, etl, blockNum)
	if err != nil {
		return false, fmt.Errorf("HandleChainSplit failed: %w", err)
	}

	// Now insert the correct block
	return insertBlockFromChain(ctx, etl, blockNum, expectedHash)
}

// insertBlockFromChain fetches block header from chain and inserts into DB
func insertBlockFromChain(ctx context.Context, etl *ETLContext, blockNum int64, expectedHash string) (bool, error) {
	// Fetch block header from chain
	header, err := GetBlockHeader(ctx, etl.EthClient, blockNum)
	if err != nil {
		return false, fmt.Errorf("GetBlockHeader failed for block %d: %w", blockNum, err)
	}

	// Verify the hash matches what we expect from the event
	actualHash := header.Hash().Hex()
	if actualHash != expectedHash {
		return false, fmt.Errorf("block %d hash mismatch: chain has %s, event has %s",
			blockNum, actualHash, expectedHash)
	}

	// Insert block into DB
	err = etl.Store.InsertBlock(ctx, header)
	if err != nil {
		return false, fmt.Errorf("block insert failed for block %d: %w", blockNum, err)
	}

	return true, nil
}

// GetBlockHeader fetches a block header from the chain by number
func GetBlockHeader(ctx context.Context, client *ethclient.Client, blockNum int64) (*types.Header, error) {
	header, err := client.HeaderByNumber(ctx, big.NewInt(blockNum))
	if err != nil {
		return nil, err
	}
	return header, nil
}

// EnsureTransactionExists verifies transaction exists in DB, or fetches and inserts it
// If RPC returns "not found", tries to fetch from archive tables
// Returns: transaction ID, whether it was newly inserted, error
func EnsureTransactionExists(ctx context.Context, etl *ETLContext, txHash common.Hash, blockNum int64) (int64, bool, error) {
	// Check if transaction exists in DB
	txId, err := etl.Store.TransactionIDByHash(ctx, txHash.Hex())
	if err == nil && txId > 0 {
		// Transaction exists
		return txId, false, nil
	}
	if err != nil && !errors.Is(err, store.ErrNotFound) {
		return 0, false, fmt.Errorf("transaction id check failed for %s: %w", txHash.Hex(), err)
	}

	// Try to fetch full transaction from chain
	tx, isPending, err := etl.EthClient.TransactionByHash(ctx, txHash)
	if err != nil {
		// Check if it's a "not found" error - only then try archive
		errStr := err.Error()
		if errStr == "not found" || errStr == "transaction not found" {
			// Transaction pruned from RPC node - try archive
			etl.Info.Printf("[RPC-MISS] Transaction %s not found in RPC, checking archive...\n", txHash.Hex())
			return fetchTransactionFromArchive(ctx, etl, txHash.Hex(), blockNum)
		}
		// Other error (connection refused, etc.) - don't use archive, propagate error
		return 0, false, fmt.Errorf("TransactionByHash failed for %s: %w", txHash.Hex(), err)
	}
	if isPending {
		return 0, false, fmt.Errorf("transaction %s is still pending", txHash.Hex())
	}

	// Get transaction receipt for gas used info
	receipt, err := etl.EthClient.TransactionReceipt(ctx, txHash)
	if err != nil {
		// Check if it's a "not found" error
		errStr := err.Error()
		if errStr == "not found" || errStr == "transaction not found" {
			// Receipt pruned - try archive
			etl.Info.Printf("[RPC-MISS] Receipt %s not found in RPC, checking archive...\n", txHash.Hex())
			return fetchTransactionFromArchive(ctx, etl, txHash.Hex(), blockNum)
		}
		// Other error - propagate
		return 0, false, fmt.Errorf("TransactionReceipt failed for %s: %w", txHash.Hex(), err)
	}

	// Insert full transaction into DB
	txId, err = etl.Store.InsertTransaction(ctx, tx, blockNum, receipt)
	if err != nil {
		return 0, false, fmt.Errorf("transaction insert failed for %s: %w", txHash.Hex(), err)
	}

	etl.Info.Printf("[RPC] Transaction %s fetched from RPC node\n", txHash.Hex())
	return txId, true, nil
}

// fetchTransactionFromArchive reads transaction from archive and inserts into main table
func fetchTransactionFromArchive(ctx context.Context, etl *ETLContext, txHash string, blockNum int64) (int64, bool, error) {
	// Try to get transaction from archive
	archTx, err := etl.Store.ArchivedTransactionByHash(ctx, txHash)
	if err != nil {
		return 0, false, fmt.Errorf("archive lookup failed for %s: %w", txHash, err)
	}
	if archTx == nil {
		// Not in archive either - create minimal record as last resort
		etl.Info.Printf("[MINIMAL] Transaction %s not in RPC or archive, creating minimal record\n", txHash)
		txId, err := etl.Store.InsertMinimalTransaction(ctx, txHash, blockNum)
		if err != nil {
			return 0, false, fmt.Errorf("minimal transaction insert failed for %s: %w", txHash, err)
		}
		return txId, true, nil
	}

	// Insert transaction from archive data
	etl.Info.Printf("[ARCHIVE] Transaction %s fetched from archive database\n", txHash)
	txId, err := etl.Store.InsertTransactionFromArchive(ctx, archTx)
	if err != nil {
		return 0, false, fmt.Errorf("archive transaction insert failed for %s: %w", txHash, err)
	}

	return txId, true, nil
}

// InsertEventLog inserts an event log into the evt_log table
// Returns: event log ID, error
func InsertEventLog(ctx context.Context, etl *ETLContext, log types.Log, txId int64) (int64, error) {
	// Look up or create contract address
	contractAid, err := etl.Store.LookupOrCreateAddress(ctx, log.Address.Hex(), int64(log.BlockNumber), txId)
	if err != nil {
		return 0, fmt.Errorf("contract address lookup failed for %s: %w", log.Address.Hex(), err)
	}

	// Insert the event log
	evtId, err := etl.Store.InsertEventLog(ctx, log, txId, contractAid)
	if err != nil {
		return 0, fmt.Errorf("event log insert failed: %w", err)
	}

	return evtId, nil
}
