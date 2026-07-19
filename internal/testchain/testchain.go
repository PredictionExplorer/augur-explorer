// Package testchain provides a deterministic in-memory fake Ethereum node for
// integration tests. It serves just enough of the JSON-RPC surface for the
// indexing pipeline (internal/indexer and the cmd/*-etl binaries) to run
// against it exactly as it would against a real node:
//
//   - eth_getBlockByNumber returns real types.Header objects whose hashes are
//     computed by go-ethereum, so hash-verification paths (insertBlockFromChain)
//     behave as in production
//   - eth_getTransactionByHash / eth_getTransactionReceipt serve properly signed
//     transactions, so sender recovery in Insert_transaction works
//   - eth_getLogs filters registered logs by block range and address
//   - eth_call dispatches to registrable per-contract handlers, so handlers that
//     read contract state (donation info records, token URIs) are testable
//
// The chain is mutable: Reorg re-mines a block range with different hashes,
// which lets tests drive the chain-split handling end to end.
//
// Everything is deterministic: one fixed private key signs all transactions,
// block timestamps derive from the block number, and no randomness is used.
package testchain

import (
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// BaseTime is the timestamp of block 0: 2026-01-01 00:00:00 UTC, matching the
// fixture conventions of the API parity suite. Block N gets BaseTime + N*100.
const BaseTime = 1767225600

// BlockTime returns the deterministic timestamp of a block.
func BlockTime(blockNum int64) uint64 {
	return uint64(BaseTime + blockNum*100) // #nosec G115 -- small positive test block numbers
}

// BlockTimeInt64 is BlockTime for assertions against int64-typed production
// values; the arithmetic never leaves the int64 domain.
func BlockTimeInt64(blockNum int64) int64 {
	return BaseTime + blockNum*100
}

// chainIDValue is an arbitrary fixed chain id for the fake chain.
const chainIDValue = 1337

// signerKeyHex is the fixed private key that signs every transaction. Its
// address is deterministic and exposed via Sender().
const signerKeyHex = "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"

// CallHandler computes the return data of an eth_call to one contract.
// input is the full calldata (4-byte selector + ABI-encoded arguments).
type CallHandler func(input []byte) ([]byte, error)

// txEntry is one transaction known to the fake chain.
type txEntry struct {
	tx       *types.Transaction
	blockNum int64
	txIndex  uint
	logs     []*types.Log
	status   uint64
	pending  bool // receipt withheld until ReleasePendingTxs
	// contractAddr is the deployment address for transactions without a
	// `to`, served in the receipt's contractAddress field.
	contractAddr common.Address
}

// Chain is the in-memory chain state plus the JSON-RPC server exposing it.
type Chain struct {
	mu      sync.Mutex
	chainID *big.Int
	key     *ecdsa.PrivateKey
	sender  common.Address

	headers map[int64]*types.Header
	forkSeq map[int64]uint64 // bumped by Reorg to change block hashes
	txs     map[common.Hash]*txEntry
	logs    []types.Log
	tip     int64

	// timeOffset shifts the timestamps of newly built blocks; accumulated by
	// evm_increaseTime (devtime.go). Atomic so eth_call stub handlers, which
	// run under the chain lock, can read it.
	timeOffset atomic.Uint64

	calls map[common.Address]CallHandler

	// Transaction-submission state (sendtx.go).
	gasPrice       *big.Int
	balances       map[common.Address]*big.Int
	nonces         map[common.Address]uint64
	nextTxReverted bool
	nextTxPending  bool
	nextSendError  string
	// rpcFailures maps a JSON-RPC method to a scheduled one-shot failure
	// (FailNextRPC / FailRPCAfter).
	rpcFailures map[string]*rpcFailure
	// submittedTxs is atomic (not mutex-guarded) so eth_call stub handlers,
	// which run under the chain lock, can read it to script reactions to
	// mined transactions.
	submittedTxs atomic.Int64
	minedTxLogs  func(tx *types.Transaction, blockNum int64) []*types.Log

	server *httptest.Server
}

// New starts a fake chain and registers cleanup with t.
func New(t *testing.T) *Chain {
	t.Helper()
	c, stop := Start()
	t.Cleanup(stop)
	return c
}

// Start starts a fake chain for TestMain-style callers. The returned stop
// function shuts the JSON-RPC server down.
func Start() (*Chain, func()) {
	key, err := crypto.HexToECDSA(signerKeyHex)
	if err != nil {
		panic(fmt.Sprintf("testchain: parsing built-in key: %v", err))
	}
	c := &Chain{
		chainID: big.NewInt(chainIDValue),
		key:     key,
		sender:  crypto.PubkeyToAddress(key.PublicKey),
		headers: make(map[int64]*types.Header),
		forkSeq: make(map[int64]uint64),
		txs:     make(map[common.Hash]*txEntry),
		calls:   make(map[common.Address]CallHandler),
	}
	c.server = httptest.NewServer(http.HandlerFunc(c.handleRPC))
	return c, c.server.Close
}

// URL returns the JSON-RPC endpoint of the fake node.
func (c *Chain) URL() string { return c.server.URL }

// ChainID returns the fixed chain id.
func (c *Chain) ChainID() *big.Int { return new(big.Int).Set(c.chainID) }

// Sender returns the address that signs (and therefore sends) every
// transaction created by AddTx.
func (c *Chain) Sender() common.Address { return c.sender }

// EnsureBlock creates the header for blockNum if it does not exist yet and
// returns it. Parent linkage is best-effort: if block N-1 exists its real hash
// is used, otherwise a deterministic placeholder stands in, so sparse block
// numbers are fine for tests that only fetch individual headers.
func (c *Chain) EnsureBlock(blockNum int64) *types.Header {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ensureBlockLocked(blockNum)
}

func (c *Chain) ensureBlockLocked(blockNum int64) *types.Header {
	if h, ok := c.headers[blockNum]; ok {
		return h
	}
	h := c.buildHeaderLocked(blockNum)
	c.headers[blockNum] = h
	if blockNum > c.tip {
		c.tip = blockNum
	}
	return h
}

func (c *Chain) buildHeaderLocked(blockNum int64) *types.Header {
	parentHash := common.BigToHash(big.NewInt(blockNum - 1))
	if parent, ok := c.headers[blockNum-1]; ok {
		parentHash = parent.Hash()
	}
	extra := make([]byte, 8)
	binary.BigEndian.PutUint64(extra, c.forkSeq[blockNum])
	return &types.Header{
		ParentHash:  parentHash,
		UncleHash:   types.EmptyUncleHash,
		Coinbase:    common.Address{},
		Root:        types.EmptyRootHash,
		TxHash:      types.EmptyTxsHash,
		ReceiptHash: types.EmptyReceiptsHash,
		Bloom:       types.Bloom{},
		Difficulty:  big.NewInt(1),
		Number:      big.NewInt(blockNum),
		GasLimit:    30_000_000,
		GasUsed:     0,
		Time:        BlockTime(blockNum) + c.timeOffset.Load(),
		Extra:       extra,
		MixDigest:   common.Hash{},
		Nonce:       types.BlockNonce{},
	}
}

// BlockHash returns the hash of the block, creating the header if needed.
func (c *Chain) BlockHash(blockNum int64) common.Hash {
	return c.EnsureBlock(blockNum).Hash()
}

// AddTx creates a signed legacy transaction addressed to `to` with the given
// calldata, records it in the given block, and returns it. The nonce derives
// from the block number and the transaction's position within the block, so a
// transaction re-added after Reorg (which drops the block's transactions)
// reproduces its original hash — the property reorg-replay tests rely on.
func (c *Chain) AddTx(blockNum int64, to common.Address, data []byte) *types.Transaction {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ensureBlockLocked(blockNum)

	var txIndex uint
	for _, e := range c.txs {
		if e.blockNum == blockNum {
			txIndex++
		}
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    uint64(blockNum)*1000 + uint64(txIndex), // #nosec G115 -- small positive test block numbers
		GasPrice: big.NewInt(1_000_000_000),
		Gas:      500_000,
		To:       &to,
		Value:    big.NewInt(0),
		Data:     data,
	})
	signed, err := types.SignTx(tx, types.LatestSignerForChainID(c.chainID), c.key)
	if err != nil {
		panic(fmt.Sprintf("testchain: signing tx: %v", err))
	}
	c.txs[signed.Hash()] = &txEntry{tx: signed, blockNum: blockNum, txIndex: txIndex, status: types.ReceiptStatusSuccessful}
	return signed
}

// AttachLogs records the logs emitted by a transaction: they become visible to
// eth_getLogs and are embedded in the transaction's receipt.
func (c *Chain) AttachLogs(txHash common.Hash, logs []*types.Log) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.txs[txHash]
	if !ok {
		panic("testchain: AttachLogs: unknown tx " + txHash.Hex())
	}
	for _, l := range logs {
		if l.Topics == nil {
			// The receipt wire format requires the topics field even when
			// empty; a nil slice would marshal to JSON null and break the
			// client-side unmarshal.
			l.Topics = []common.Hash{}
		}
		entry.logs = append(entry.logs, l)
		c.logs = append(c.logs, *l)
	}
}

// Reorg re-mines every block from fromBlock to the tip: their hashes change,
// and all transactions (and logs) recorded in that range are dropped, as if a
// competing chain segment had replaced them. Tests then re-add transactions to
// the new blocks.
func (c *Chain) Reorg(fromBlock int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for num := fromBlock; num <= c.tip; num++ {
		if _, ok := c.headers[num]; !ok {
			continue
		}
		c.forkSeq[num]++
		delete(c.headers, num)
	}
	// Rebuild in ascending order so parent hashes chain correctly.
	for num := fromBlock; num <= c.tip; num++ {
		c.headers[num] = c.buildHeaderLocked(num)
	}
	for hash, entry := range c.txs {
		if entry.blockNum >= fromBlock {
			delete(c.txs, hash)
		}
	}
	kept := c.logs[:0]
	for _, l := range c.logs {
		if int64(l.BlockNumber) < fromBlock { // #nosec G115 -- fake-chain logs carry small positive block numbers
			kept = append(kept, l)
		}
	}
	c.logs = kept
}

// RegisterCall installs the eth_call handler for one contract address.
func (c *Chain) RegisterCall(to common.Address, handler CallHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.calls[to] = handler
}

// rpcFailure is one scheduled per-method failure: skip calls pass through
// before the failing one.
type rpcFailure struct {
	skip int
	msg  string
}

// FailNextRPC makes the next call of the given JSON-RPC method (e.g.
// "eth_gasPrice", "eth_getTransactionCount") fail once with the message, so
// per-call error branches of RPC consumers are testable.
func (c *Chain) FailNextRPC(method, msg string) {
	c.FailRPCAfter(method, 0, msg)
}

// FailRPCAfter schedules a one-shot failure of the given JSON-RPC method
// after skip successful calls (skip=0 fails the next call). Useful when the
// code under test makes the same call several times and a later one must
// fail.
func (c *Chain) FailRPCAfter(method string, skip int, msg string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.rpcFailures == nil {
		c.rpcFailures = make(map[string]*rpcFailure)
	}
	c.rpcFailures[method] = &rpcFailure{skip: skip, msg: msg}
}

// --- JSON-RPC plumbing ---

type rpcRequest struct {
	ID     json.RawMessage   `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

func (c *Chain) handleRPC(w http.ResponseWriter, r *http.Request) {
	var raw json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	single := true
	var reqs []rpcRequest
	if len(raw) > 0 && raw[0] == '[' {
		single = false
		if err := json.Unmarshal(raw, &reqs); err != nil {
			http.Error(w, "bad batch", http.StatusBadRequest)
			return
		}
	} else {
		var one rpcRequest
		if err := json.Unmarshal(raw, &one); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		reqs = []rpcRequest{one}
	}

	resps := make([]map[string]any, 0, len(reqs))
	for _, req := range reqs {
		resp := map[string]any{"jsonrpc": "2.0", "id": req.ID}
		result, err := c.dispatch(req)
		if err != nil {
			resp["error"] = map[string]any{"code": -32000, "message": err.Error()}
		} else {
			resp["result"] = result
		}
		resps = append(resps, resp)
	}

	w.Header().Set("Content-Type", "application/json")
	// The response maps hold only JSON-safe values (strings, numbers and
	// nested json.Unmarshal output); a write failure means the test client
	// disconnected, which the client side reports on its own.
	if single {
		_ = json.NewEncoder(w).Encode(resps[0]) //nolint:errchkjson // JSON-safe by construction; client reports transport failures
		return
	}
	_ = json.NewEncoder(w).Encode(resps) //nolint:errchkjson // JSON-safe by construction; client reports transport failures
}

func (c *Chain) dispatch(req rpcRequest) (any, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if failure, ok := c.rpcFailures[req.Method]; ok {
		if failure.skip > 0 {
			failure.skip--
		} else {
			delete(c.rpcFailures, req.Method)
			return nil, fmt.Errorf("%s", failure.msg)
		}
	}
	switch req.Method {
	case "eth_chainId":
		return hexUint64(c.chainID.Uint64()), nil
	case "eth_blockNumber":
		return hexUint64(uint64(c.tip)), nil // #nosec G115 -- the fake chain's tip is a small positive test height
	case "eth_getBlockByNumber":
		return c.getBlockByNumber(req.Params)
	case "eth_getTransactionByHash":
		return c.getTransactionByHash(req.Params)
	case "eth_getTransactionReceipt":
		return c.getTransactionReceipt(req.Params)
	case "eth_getLogs":
		return c.getLogs(req.Params)
	case "eth_call":
		return c.ethCall(req.Params)
	case "eth_getBalance":
		return c.getBalance(req.Params)
	case "eth_getCode":
		return "0x", nil
	case "eth_gasPrice":
		return "0x" + c.gasPriceLocked().Text(16), nil
	case "eth_getTransactionCount":
		return c.getTransactionCount(req.Params)
	case "eth_sendRawTransaction":
		return c.sendRawTransaction(req.Params)
	case "evm_increaseTime":
		return c.evmIncreaseTime(req.Params)
	case "evm_mine":
		return c.evmMine(), nil
	default:
		return nil, fmt.Errorf("testchain: method %s not supported", req.Method)
	}
}

func hexUint64(v uint64) string { return fmt.Sprintf("0x%x", v) }

// parseBlockTag converts "latest"/"pending" or a hex quantity to a height.
func (c *Chain) parseBlockTag(raw json.RawMessage) (int64, error) {
	var tag string
	if err := json.Unmarshal(raw, &tag); err != nil {
		return 0, fmt.Errorf("testchain: bad block tag: %w", err)
	}
	switch tag {
	case "latest", "pending", "safe", "finalized":
		return c.tip, nil
	case "earliest":
		return 0, nil
	}
	v, err := hexToUint64(tag)
	if err != nil {
		return 0, err
	}
	if v > math.MaxInt64 {
		return 0, fmt.Errorf("testchain: block tag %s overflows int64", tag)
	}
	return int64(v), nil
}

func hexToUint64(s string) (uint64, error) {
	trimmed := strings.TrimPrefix(s, "0x")
	v, ok := new(big.Int).SetString(trimmed, 16)
	if !ok {
		return 0, fmt.Errorf("testchain: bad hex quantity %q", s)
	}
	return v.Uint64(), nil
}

func (c *Chain) getBlockByNumber(params []json.RawMessage) (any, error) {
	if len(params) < 1 {
		return nil, errors.New("testchain: eth_getBlockByNumber needs params")
	}
	num, err := c.parseBlockTag(params[0])
	if err != nil {
		return nil, err
	}
	header, ok := c.headers[num]
	if !ok {
		return nil, nil // JSON null => ethereum.NotFound at the client
	}
	// types.Header marshals to the wire format go-ethereum's client
	// unmarshals; the client recomputes the hash from the fields, so the
	// derived "hash" value we add is informational.
	encoded, err := json.Marshal(header)
	if err != nil {
		return nil, fmt.Errorf("testchain: marshaling header: %w", err)
	}
	var m map[string]any
	if err := json.Unmarshal(encoded, &m); err != nil {
		return nil, err
	}
	m["hash"] = header.Hash().Hex()
	// Block-only fields expected by clients that parse full blocks.
	m["transactions"] = []any{}
	m["uncles"] = []any{}
	m["size"] = "0x0"
	return m, nil
}

func (c *Chain) getTransactionByHash(params []json.RawMessage) (any, error) {
	hash, err := paramHash(params)
	if err != nil {
		return nil, err
	}
	entry, ok := c.txs[hash]
	if !ok {
		return nil, nil // JSON null => ethereum.NotFound ("not found")
	}
	encoded, err := json.Marshal(entry.tx)
	if err != nil {
		return nil, fmt.Errorf("testchain: marshaling tx: %w", err)
	}
	var m map[string]any
	if err := json.Unmarshal(encoded, &m); err != nil {
		return nil, err
	}
	header := c.ensureBlockLocked(entry.blockNum)
	m["blockHash"] = header.Hash().Hex()
	m["blockNumber"] = hexUint64(uint64(entry.blockNum)) // #nosec G115 -- small positive test block numbers
	m["transactionIndex"] = hexUint64(uint64(entry.txIndex))
	m["from"] = c.sender.Hex()
	return m, nil
}

func (c *Chain) getTransactionReceipt(params []json.RawMessage) (any, error) {
	hash, err := paramHash(params)
	if err != nil {
		return nil, err
	}
	entry, ok := c.txs[hash]
	if !ok || entry.pending {
		return nil, nil
	}
	header := c.ensureBlockLocked(entry.blockNum)
	logs := entry.logs
	if logs == nil {
		logs = []*types.Log{}
	}
	receipt := &types.Receipt{
		Type:              entry.tx.Type(),
		Status:            entry.status,
		CumulativeGasUsed: 21_000,
		Bloom:             types.Bloom{},
		Logs:              logs,
		TxHash:            entry.tx.Hash(),
		ContractAddress:   entry.contractAddr,
		GasUsed:           21_000,
		EffectiveGasPrice: entry.tx.GasPrice(),
		BlockHash:         header.Hash(),
		BlockNumber:       big.NewInt(entry.blockNum),
		TransactionIndex:  entry.txIndex,
	}
	return json.RawMessage(mustMarshal(receipt)), nil
}

func mustMarshal(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Sprintf("testchain: marshal: %v", err))
	}
	return b
}

func paramHash(params []json.RawMessage) (common.Hash, error) {
	if len(params) < 1 {
		return common.Hash{}, errors.New("testchain: missing hash param")
	}
	var s string
	if err := json.Unmarshal(params[0], &s); err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(s), nil
}

type logFilter struct {
	FromBlock string          `json:"fromBlock"`
	ToBlock   string          `json:"toBlock"`
	Address   json.RawMessage `json:"address"`
	Topics    [][]common.Hash `json:"topics"`
}

func (c *Chain) getLogs(params []json.RawMessage) (any, error) {
	if len(params) < 1 {
		return nil, errors.New("testchain: eth_getLogs needs a filter")
	}
	var f logFilter
	if err := json.Unmarshal(params[0], &f); err != nil {
		return nil, fmt.Errorf("testchain: bad filter: %w", err)
	}
	from, err := c.parseBlockTag(mustMarshal(orLatest(f.FromBlock)))
	if err != nil {
		return nil, err
	}
	to, err := c.parseBlockTag(mustMarshal(orLatest(f.ToBlock)))
	if err != nil {
		return nil, err
	}
	addrs, err := filterAddresses(f.Address)
	if err != nil {
		return nil, err
	}
	out := make([]types.Log, 0)
	for _, l := range c.logs {
		if int64(l.BlockNumber) < from || int64(l.BlockNumber) > to { // #nosec G115 -- fake-chain logs carry small positive block numbers
			continue
		}
		if len(addrs) > 0 && !containsAddress(addrs, l.Address) {
			continue
		}
		out = append(out, l)
	}
	return out, nil
}

func orLatest(tag string) string {
	if tag == "" {
		return "latest"
	}
	return tag
}

func filterAddresses(raw json.RawMessage) ([]common.Address, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	var one string
	if err := json.Unmarshal(raw, &one); err == nil {
		return []common.Address{common.HexToAddress(one)}, nil
	}
	var many []string
	if err := json.Unmarshal(raw, &many); err != nil {
		return nil, fmt.Errorf("testchain: bad address filter: %w", err)
	}
	out := make([]common.Address, 0, len(many))
	for _, s := range many {
		out = append(out, common.HexToAddress(s))
	}
	return out, nil
}

func containsAddress(addrs []common.Address, a common.Address) bool {
	return slices.Contains(addrs, a)
}

type callObject struct {
	To   *common.Address `json:"to"`
	Data string          `json:"data"`
	// Some clients send calldata as "input" instead of "data".
	Input string `json:"input"`
}

func (c *Chain) ethCall(params []json.RawMessage) (any, error) {
	if len(params) < 1 {
		return nil, errors.New("testchain: eth_call needs a call object")
	}
	var call callObject
	if err := json.Unmarshal(params[0], &call); err != nil {
		return nil, fmt.Errorf("testchain: bad call object: %w", err)
	}
	if call.To == nil {
		return nil, errors.New("testchain: eth_call without `to` not supported")
	}
	handler, ok := c.calls[*call.To]
	if !ok {
		return nil, fmt.Errorf("testchain: no call handler registered for %s", call.To.Hex())
	}
	dataHex := call.Data
	if dataHex == "" {
		dataHex = call.Input
	}
	input, err := hex.DecodeString(strings.TrimPrefix(dataHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("testchain: bad calldata: %w", err)
	}
	output, err := handler(input)
	if err != nil {
		return nil, err
	}
	return "0x" + hex.EncodeToString(output), nil
}
