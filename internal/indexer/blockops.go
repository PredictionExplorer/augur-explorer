// Block, transaction and event-log pipeline: every fetched log passes through
// EnsureBlockExists -> EnsureTransactionExists -> InsertEventLog before its
// stored form is dispatched to the processor.

package indexer

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// logBlockNum converts an RPC-supplied log block number to the int64 the
// store schema uses, rejecting values beyond int64 (corrupt node data)
// instead of wrapping them negative.
func logBlockNum(lg *types.Log) (int64, error) {
	if lg.BlockNumber > math.MaxInt64 {
		return 0, fmt.Errorf("log block number %d overflows int64", lg.BlockNumber)
	}
	return int64(lg.BlockNumber), nil
}

// EnsureBlockExists verifies the block exists in the database with the
// expected hash, inserting it from the chain when missing. A stored hash that
// disagrees with expectedHash means a chain split: the divergent range is
// rolled back (HandleChainSplit) and the replacement block inserted.
// It reports whether the block was newly inserted.
func (e *Engine) EnsureBlockExists(ctx context.Context, blockNum int64, expectedHash string) (bool, error) {
	existingHash, err := e.store.BlockHash(ctx, blockNum)
	if errors.Is(err, store.ErrNotFound) {
		return e.insertBlockFromChain(ctx, blockNum, expectedHash)
	}
	if err != nil {
		return false, fmt.Errorf("block hash check failed for block %d: %w", blockNum, err)
	}

	if existingHash == expectedHash {
		return false, nil
	}

	// Hash mismatch: chain split detected.
	e.log.Info("chain split detected",
		"block", blockNum, "db_hash", existingHash, "event_hash", expectedHash)

	if err := e.HandleChainSplit(ctx, blockNum); err != nil {
		return false, fmt.Errorf("HandleChainSplit failed: %w", err)
	}

	return e.insertBlockFromChain(ctx, blockNum, expectedHash)
}

// insertBlockFromChain fetches the block header from the chain, verifies it
// carries the hash the fetched log claims, and inserts it.
func (e *Engine) insertBlockFromChain(ctx context.Context, blockNum int64, expectedHash string) (bool, error) {
	header, err := e.client.HeaderByNumber(ctx, big.NewInt(blockNum))
	if err != nil {
		return false, fmt.Errorf("header fetch failed for block %d: %w", blockNum, err)
	}

	if actualHash := header.Hash().Hex(); actualHash != expectedHash {
		return false, fmt.Errorf("block %d hash mismatch: chain has %s, event has %s",
			blockNum, actualHash, expectedHash)
	}

	if err := e.store.InsertBlock(ctx, header); err != nil {
		return false, fmt.Errorf("block insert failed for block %d: %w", blockNum, err)
	}
	return true, nil
}

// EnsureTransactionExists returns the database id of the transaction,
// fetching and inserting it when missing. Transactions the RPC node has
// pruned fall back to the archive tables, and as a last resort a minimal
// record (hash + block) is created so the event log can still reference it.
// It reports the transaction id and whether a row was newly inserted.
func (e *Engine) EnsureTransactionExists(ctx context.Context, txHash common.Hash, blockNum int64) (int64, bool, error) {
	txID, err := e.store.TransactionIDByHash(ctx, txHash.Hex())
	if err == nil && txID > 0 {
		return txID, false, nil
	}
	if err != nil && !errors.Is(err, store.ErrNotFound) {
		return 0, false, fmt.Errorf("transaction id check failed for %s: %w", txHash.Hex(), err)
	}

	tx, isPending, err := e.client.TransactionByHash(ctx, txHash)
	if err != nil {
		if isNotFoundRPCError(err) {
			// Pruned from the RPC node: try the archive.
			e.log.Info("transaction not in RPC node, checking archive", "tx", txHash.Hex())
			return e.fetchTransactionFromArchive(ctx, txHash.Hex(), blockNum)
		}
		// Connection failures etc. must not degrade to the archive path.
		return 0, false, fmt.Errorf("TransactionByHash failed for %s: %w", txHash.Hex(), err)
	}
	if isPending {
		return 0, false, fmt.Errorf("transaction %s is still pending", txHash.Hex())
	}

	receipt, err := e.client.TransactionReceipt(ctx, txHash)
	if err != nil {
		if isNotFoundRPCError(err) {
			e.log.Info("receipt not in RPC node, checking archive", "tx", txHash.Hex())
			return e.fetchTransactionFromArchive(ctx, txHash.Hex(), blockNum)
		}
		return 0, false, fmt.Errorf("TransactionReceipt failed for %s: %w", txHash.Hex(), err)
	}

	txID, err = e.store.InsertTransaction(ctx, tx, blockNum, receipt)
	if err != nil {
		return 0, false, fmt.Errorf("transaction insert failed for %s: %w", txHash.Hex(), err)
	}
	e.log.Debug("transaction fetched from RPC node", "tx", txHash.Hex())
	return txID, true, nil
}

// isNotFoundRPCError matches the responses go-ethereum surfaces when a node
// no longer carries a transaction or receipt (pruned/unindexed).
func isNotFoundRPCError(err error) bool {
	msg := err.Error()
	return msg == "not found" || msg == "transaction not found"
}

// fetchTransactionFromArchive restores a pruned transaction from the archive
// tables, degrading to a minimal record when the archive misses it too.
func (e *Engine) fetchTransactionFromArchive(ctx context.Context, txHash string, blockNum int64) (int64, bool, error) {
	archTx, err := e.store.ArchivedTransactionByHash(ctx, txHash)
	if err != nil {
		return 0, false, fmt.Errorf("archive lookup failed for %s: %w", txHash, err)
	}
	if archTx == nil {
		e.log.Warn("transaction not in RPC or archive, creating minimal record", "tx", txHash)
		txID, err := e.store.InsertMinimalTransaction(ctx, txHash, blockNum)
		if err != nil {
			return 0, false, fmt.Errorf("minimal transaction insert failed for %s: %w", txHash, err)
		}
		return txID, true, nil
	}

	e.log.Info("transaction restored from archive", "tx", txHash)
	txID, err := e.store.InsertTransactionFromArchive(ctx, archTx)
	if err != nil {
		return 0, false, fmt.Errorf("archive transaction insert failed for %s: %w", txHash, err)
	}
	return txID, true, nil
}

// InsertEventLog stores one fetched log in evt_log (resolving the emitting
// contract's address id first) and returns the new row's id.
func (e *Engine) InsertEventLog(ctx context.Context, log types.Log, txID int64) (int64, error) {
	blockNum, err := logBlockNum(&log)
	if err != nil {
		return 0, err
	}
	contractAid, err := e.store.LookupOrCreateAddress(ctx, log.Address.Hex(), blockNum, txID)
	if err != nil {
		return 0, fmt.Errorf("contract address lookup failed for %s: %w", log.Address.Hex(), err)
	}

	evtID, err := e.store.InsertEventLog(ctx, log, txID, contractAid)
	if err != nil {
		return 0, fmt.Errorf("event log insert failed: %w", err)
	}
	return evtID, nil
}
