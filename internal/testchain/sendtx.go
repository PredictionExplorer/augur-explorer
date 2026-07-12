package testchain

// Transaction-submission support: the JSON-RPC surface a signing client
// (internal/ethtx, the operator CLIs) needs on top of the read-only indexing
// surface — eth_gasPrice, eth_getTransactionCount, eth_getBalance with
// settable per-address balances, and eth_sendRawTransaction, which mines the
// submitted transaction into a fresh block and serves its receipt.

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// defaultGasPrice is the eth_gasPrice answer until SetGasPrice overrides it.
var defaultGasPrice = big.NewInt(1_000_000_000) // 1 gwei

// SetGasPrice overrides the value served by eth_gasPrice.
func (c *Chain) SetGasPrice(wei *big.Int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.gasPrice = new(big.Int).Set(wei)
}

// SetBalance sets the eth_getBalance answer for one address. Addresses
// without an entry keep the default zero balance.
func (c *Chain) SetBalance(addr common.Address, wei *big.Int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.balances == nil {
		c.balances = make(map[common.Address]*big.Int)
	}
	c.balances[addr] = new(big.Int).Set(wei)
}

// SetNonce sets the eth_getTransactionCount answer for one address. A
// subsequent eth_sendRawTransaction from that address advances it to the
// submitted transaction's nonce + 1, like a real node's pending count.
func (c *Chain) SetNonce(addr common.Address, nonce uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.nonces == nil {
		c.nonces = make(map[common.Address]uint64)
	}
	c.nonces[addr] = nonce
}

// MarkNextTxReverted makes the next submitted transaction's receipt carry
// status 0 (on-chain revert). One-shot.
func (c *Chain) MarkNextTxReverted() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nextTxReverted = true
}

// MarkNextTxPending makes the next submitted transaction stay pending:
// eth_getTransactionReceipt answers null until ReleasePendingTxs is called,
// so receipt-wait timeout paths are testable. One-shot per transaction.
func (c *Chain) MarkNextTxPending() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nextTxPending = true
}

// ReleasePendingTxs makes the receipts of all held transactions visible.
func (c *Chain) ReleasePendingTxs() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, entry := range c.txs {
		entry.pending = false
	}
}

// RejectNextSendWith makes the next eth_sendRawTransaction fail with the
// given message, as a node would reject an underpriced or invalid
// transaction. One-shot.
func (c *Chain) RejectNextSendWith(msg string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nextSendError = msg
}

// SubmittedTxCount returns how many transactions have been submitted through
// eth_sendRawTransaction. It is safe to call from eth_call stub handlers
// (which run under the chain lock): tests script contract state that
// "reacts" to mined transactions with it.
func (c *Chain) SubmittedTxCount() int {
	return int(c.submittedTxs.Load())
}

// SetMinedTxLogs installs a hook that produces the receipt logs of every
// subsequently submitted transaction (the fake chain does not execute EVM
// code, so tests that need event logs on submitted transactions script them
// here). The hook runs under the chain lock and must not call back into the
// chain.
func (c *Chain) SetMinedTxLogs(hook func(tx *types.Transaction, blockNum int64) []*types.Log) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.minedTxLogs = hook
}

// gasPriceLocked returns the configured or default gas price.
func (c *Chain) gasPriceLocked() *big.Int {
	if c.gasPrice != nil {
		return c.gasPrice
	}
	return defaultGasPrice
}

func (c *Chain) getBalance(params []json.RawMessage) (any, error) {
	addr, err := paramAddress(params)
	if err != nil {
		return nil, err
	}
	if bal, ok := c.balances[addr]; ok {
		return "0x" + bal.Text(16), nil
	}
	return "0x0", nil
}

func (c *Chain) getTransactionCount(params []json.RawMessage) (any, error) {
	addr, err := paramAddress(params)
	if err != nil {
		return nil, err
	}
	return hexUint64(c.nonces[addr]), nil
}

func paramAddress(params []json.RawMessage) (common.Address, error) {
	if len(params) < 1 {
		return common.Address{}, fmt.Errorf("testchain: missing address param")
	}
	var s string
	if err := json.Unmarshal(params[0], &s); err != nil {
		return common.Address{}, fmt.Errorf("testchain: bad address param: %w", err)
	}
	if !common.IsHexAddress(s) {
		return common.Address{}, fmt.Errorf("testchain: invalid address %q", s)
	}
	return common.HexToAddress(s), nil
}

// sendRawTransaction decodes a signed transaction, mines it into a fresh
// block at the tip, records its receipt, and returns the transaction hash.
func (c *Chain) sendRawTransaction(params []json.RawMessage) (any, error) {
	if c.nextSendError != "" {
		msg := c.nextSendError
		c.nextSendError = ""
		return nil, fmt.Errorf("%s", msg)
	}
	if len(params) < 1 {
		return nil, fmt.Errorf("testchain: eth_sendRawTransaction needs raw tx bytes")
	}
	var rawHex string
	if err := json.Unmarshal(params[0], &rawHex); err != nil {
		return nil, fmt.Errorf("testchain: bad raw tx param: %w", err)
	}
	raw, err := hex.DecodeString(strings.TrimPrefix(rawHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("testchain: bad raw tx hex: %w", err)
	}
	tx := new(types.Transaction)
	if err := tx.UnmarshalBinary(raw); err != nil {
		return nil, fmt.Errorf("testchain: decoding raw tx: %w", err)
	}
	sender, err := types.Sender(types.LatestSignerForChainID(c.chainID), tx)
	if err != nil {
		return nil, fmt.Errorf("testchain: recovering sender: %w", err)
	}

	blockNum := c.tip + 1
	c.ensureBlockLocked(blockNum)

	status := types.ReceiptStatusSuccessful
	if c.nextTxReverted {
		status = types.ReceiptStatusFailed
		c.nextTxReverted = false
	}
	pending := c.nextTxPending
	c.nextTxPending = false

	entry := &txEntry{
		tx:       tx,
		blockNum: blockNum,
		txIndex:  0,
		status:   status,
		pending:  pending,
	}
	if tx.To() == nil {
		entry.contractAddr = crypto.CreateAddress(sender, tx.Nonce())
	}
	if c.minedTxLogs != nil {
		for _, l := range c.minedTxLogs(tx, blockNum) {
			if l.Topics == nil {
				l.Topics = []common.Hash{}
			}
			entry.logs = append(entry.logs, l)
			c.logs = append(c.logs, *l)
		}
	}
	c.txs[tx.Hash()] = entry
	c.submittedTxs.Add(1)
	if c.nonces == nil {
		c.nonces = make(map[common.Address]uint64)
	}
	c.nonces[sender] = tx.Nonce() + 1
	return tx.Hash().Hex(), nil
}
