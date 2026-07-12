package testchain

// Development-node time manipulation: the evm_increaseTime and evm_mine
// JSON-RPC methods that Hardhat/Ganache expose. The operator CLIs use them
// (via ethtx.Session.AdvanceDevChainTime) to wait out contract timers on
// local chains, so the fake node supports them too: evm_increaseTime shifts
// the timestamp of every subsequently built block, and evm_mine mines an
// empty block at the new tip.

import (
	"encoding/json"
	"fmt"
	"math/big"
)

// SetChainID overrides the fake chain's id (default 1337). Dev-chain-only
// code paths key on well-known ids (Hardhat's 31337), so tests exercising
// them need to pick the id. Must be called before any transaction exists:
// signatures embed the chain id, and re-signing is not supported.
func (c *Chain) SetChainID(id int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.txs) > 0 {
		panic("testchain: SetChainID called after transactions were created")
	}
	c.chainID = big.NewInt(id)
}

// TimeOffset returns the accumulated evm_increaseTime shift applied to the
// timestamps of newly built blocks. It is safe to call from eth_call stub
// handlers (which run under the chain lock).
func (c *Chain) TimeOffset() uint64 {
	return c.timeOffset.Load()
}

func (c *Chain) evmIncreaseTime(params []json.RawMessage) (any, error) {
	if len(params) < 1 {
		return nil, fmt.Errorf("testchain: evm_increaseTime needs a seconds param")
	}
	var secs int64
	if err := json.Unmarshal(params[0], &secs); err != nil {
		return nil, fmt.Errorf("testchain: bad evm_increaseTime param: %w", err)
	}
	if secs < 0 {
		return nil, fmt.Errorf("testchain: evm_increaseTime seconds must be non-negative")
	}
	// Hardhat answers with the total adjustment in seconds.
	return c.timeOffset.Add(uint64(secs)), nil
}

func (c *Chain) evmMine() any {
	c.ensureBlockLocked(c.tip + 1)
	return "0x0"
}
