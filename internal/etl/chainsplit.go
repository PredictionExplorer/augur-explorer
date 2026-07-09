// Package common - Chain split detection and handling
package common

import (
	"context"
	"fmt"
)

// HandleChainSplit handles a chain reorganization by deleting blocks from the divergent point
// It deletes blocks from the highest block number down to the divergent block
// This order is important because DELETE triggers decrement cumulative statistics
func HandleChainSplit(ctx context.Context, etl *ETLContext, divergentBlock int64) error {
	etl.Info.Printf("Handling chain split from block %d\n", divergentBlock)

	// Get the highest block we have in the database
	lastBlock, err := etl.Store.LastBlockNum(ctx)
	if err != nil {
		return fmt.Errorf("last block lookup failed: %w", err)
	}

	if lastBlock < divergentBlock {
		// No blocks to delete
		etl.Info.Printf("No blocks to delete (last block %d < divergent block %d)\n", lastBlock, divergentBlock)
		return nil
	}

	// Delete blocks from highest to divergent (order matters for trigger decrements)
	etl.Info.Printf("Deleting blocks from %d down to %d\n", lastBlock, divergentBlock)

	for blockNum := lastBlock; blockNum >= divergentBlock; blockNum-- {
		err := etl.Store.DeleteBlock(ctx, blockNum)
		if err != nil {
			return fmt.Errorf("block delete failed for block %d: %w", blockNum, err)
		}
		etl.Info.Printf("Deleted block %d\n", blockNum)
	}

	// Update last_block table to point to block before divergent
	if divergentBlock > 0 {
		err = etl.Store.SetLastBlockNum(ctx, divergentBlock-1)
		if err != nil {
			return fmt.Errorf("watermark update failed: %w", err)
		}
	}

	etl.Info.Printf("Chain split handled, deleted %d blocks\n", lastBlock-divergentBlock+1)
	return nil
}
