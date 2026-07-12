package ethtx

import (
	"context"
	"fmt"
)

// HardhatChainID is the well-known chain id of Hardhat/Anvil development
// nodes. Commands that manipulate block time gate on it so the evm_* methods
// are never sent to a real network.
const HardhatChainID = 31337

// IsDevChain reports whether the connected network is a local development
// node whose block time can be advanced with AdvanceDevChainTime.
func (n *Network) IsDevChain() bool {
	return n.ChainID != nil && n.ChainID.Int64() == HardhatChainID
}

// AdvanceDevChainTime advances block time on Hardhat/Ganache development
// nodes (evm_increaseTime + evm_mine) and refreshes the cached block number
// and timestamp on success. Calling it with zero or negative seconds is a
// no-op.
func (n *Network) AdvanceDevChainTime(ctx context.Context, seconds int64) error {
	if seconds <= 0 {
		return nil
	}
	rpc := n.Client.Client()
	var result any
	if err := rpc.CallContext(ctx, &result, "evm_increaseTime", seconds); err != nil {
		return fmt.Errorf("evm_increaseTime(%d): %w", seconds, err)
	}
	if err := rpc.CallContext(ctx, &result, "evm_mine"); err != nil {
		return fmt.Errorf("evm_mine: %w", err)
	}
	header, err := n.Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("refresh block header: %w", err)
	}
	n.BlockNum = header.Number
	n.BlockTime = header.Time
	return nil
}
