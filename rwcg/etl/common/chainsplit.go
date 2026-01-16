// Package common - Chain split detection and handling
package common

import (
	"fmt"
)

// HandleChainSplit handles a chain reorganization by deleting blocks from the divergent point
// It deletes blocks from the highest block number down to the divergent block
// This order is important because DELETE triggers decrement cumulative statistics
func HandleChainSplit(ctx *ETLContext, divergentBlock int64) error {
	ctx.Info.Printf("Handling chain split from block %d\n", divergentBlock)

	// Get the highest block we have in the database
	lastBlock, err := ctx.Storage.Get_last_block_num()
	if err != nil {
		return fmt.Errorf("Get_last_block_num failed: %v", err)
	}

	if lastBlock < divergentBlock {
		// No blocks to delete
		ctx.Info.Printf("No blocks to delete (last block %d < divergent block %d)\n", lastBlock, divergentBlock)
		return nil
	}

	// Delete blocks from highest to divergent (order matters for trigger decrements)
	ctx.Info.Printf("Deleting blocks from %d down to %d\n", lastBlock, divergentBlock)

	for blockNum := lastBlock; blockNum >= divergentBlock; blockNum-- {
		err := ctx.Storage.Delete_block(blockNum)
		if err != nil {
			return fmt.Errorf("Delete_block failed for block %d: %v", blockNum, err)
		}
		ctx.Info.Printf("Deleted block %d\n", blockNum)
	}

	// Update last_block table to point to block before divergent
	if divergentBlock > 0 {
		err = ctx.Storage.Set_last_block_num(divergentBlock - 1)
		if err != nil {
			return fmt.Errorf("Set_last_block_num failed: %v", err)
		}
	}

	ctx.Info.Printf("Chain split handled, deleted %d blocks\n", lastBlock-divergentBlock+1)
	return nil
}

// ValidateBlockHash checks if the block hash in the database matches the expected hash
// Returns: matches (true/false), error if block doesn't exist
func ValidateBlockHash(ctx *ETLContext, blockNum int64, expectedHash string) (bool, error) {
	existingHash, err := ctx.Storage.Get_block_hash(blockNum)
	if err != nil {
		// Block doesn't exist
		return false, err
	}

	return existingHash == expectedHash, nil
}

// DeleteFutureBlocks deletes all blocks greater than the specified block number
// Used when a chain split is detected to clean up orphaned blocks
func DeleteFutureBlocks(ctx *ETLContext, fromBlock int64) error {
	lastBlock, err := ctx.Storage.Get_last_block_num()
	if err != nil {
		return fmt.Errorf("Get_last_block_num failed: %v", err)
	}

	if lastBlock <= fromBlock {
		return nil // No future blocks to delete
	}

	// Delete from highest to fromBlock+1
	for blockNum := lastBlock; blockNum > fromBlock; blockNum-- {
		err := ctx.Storage.Delete_block(blockNum)
		if err != nil {
			return fmt.Errorf("Delete_block failed for block %d: %v", blockNum, err)
		}
	}

	// Update last_block
	err = ctx.Storage.Set_last_block_num(fromBlock)
	if err != nil {
		return fmt.Errorf("Set_last_block_num failed: %v", err)
	}

	return nil
}
