// Chain reorganization handling: rolling the database back to the last
// common ancestor so the replacement fork can be re-processed.

package indexer

import (
	"context"
	"fmt"
)

// HandleChainSplit rolls back a chain reorganization by deleting every stored
// block from the tip down to divergentBlock (highest first — the DELETE
// triggers reverse cumulative statistics, so order matters) and rewinding the
// block watermark to just before the divergence.
func (e *Engine) HandleChainSplit(ctx context.Context, divergentBlock int64) error {
	e.log.Info("handling chain split", "divergent_block", divergentBlock)

	lastBlock, err := e.store.LastBlockNum(ctx)
	if err != nil {
		return fmt.Errorf("last block lookup failed: %w", err)
	}

	if lastBlock < divergentBlock {
		e.log.Info("no blocks to delete",
			"last_block", lastBlock, "divergent_block", divergentBlock)
		return nil
	}

	for blockNum := lastBlock; blockNum >= divergentBlock; blockNum-- {
		if err := e.store.DeleteBlock(ctx, blockNum); err != nil {
			return fmt.Errorf("block delete failed for block %d: %w", blockNum, err)
		}
	}

	if divergentBlock > 0 {
		if err := e.store.SetLastBlockNum(ctx, divergentBlock-1); err != nil {
			return fmt.Errorf("watermark update failed: %w", err)
		}
	}

	e.metrics.reorgHandled()
	e.log.Info("chain split handled",
		"deleted_blocks", lastBlock-divergentBlock+1, "divergent_block", divergentBlock)
	return nil
}
