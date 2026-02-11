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
