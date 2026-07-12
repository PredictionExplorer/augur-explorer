package testchain

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

func dialChain(t *testing.T, c *Chain) *ethclient.Client {
	t.Helper()
	ec, err := ethclient.Dial(c.URL())
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(ec.Close)
	return ec
}

func TestEvmIncreaseTimeShiftsNewBlocks(t *testing.T) {
	c := New(t)
	c.EnsureBlock(10)
	baseTime := c.EnsureBlock(10).Time

	ec := dialChain(t, c)
	rpc := ec.Client()

	var total uint64
	if err := rpc.CallContext(context.Background(), &total, "evm_increaseTime", 500); err != nil {
		t.Fatalf("evm_increaseTime: %v", err)
	}
	if total != 500 {
		t.Errorf("total adjustment = %d, want 500", total)
	}
	var mineResult string
	if err := rpc.CallContext(context.Background(), &mineResult, "evm_mine"); err != nil {
		t.Fatalf("evm_mine: %v", err)
	}

	// The pre-existing block keeps its timestamp; the mined block shifts.
	if got := c.EnsureBlock(10).Time; got != baseTime {
		t.Errorf("existing block time changed: %d -> %d", baseTime, got)
	}
	header, err := ec.HeaderByNumber(context.Background(), big.NewInt(11))
	if err != nil {
		t.Fatalf("header 11: %v", err)
	}
	want := BlockTime(11) + 500
	if header.Time != want {
		t.Errorf("mined block time = %d, want %d", header.Time, want)
	}

	// Offsets accumulate.
	if err := rpc.CallContext(context.Background(), &total, "evm_increaseTime", 100); err != nil {
		t.Fatalf("second evm_increaseTime: %v", err)
	}
	if total != 600 {
		t.Errorf("accumulated adjustment = %d, want 600", total)
	}
	if got := c.TimeOffset(); got != 600 {
		t.Errorf("TimeOffset() = %d, want 600", got)
	}
}

func TestEvmIncreaseTimeRejectsBadParams(t *testing.T) {
	c := New(t)
	ec := dialChain(t, c)
	rpc := ec.Client()

	var out uint64
	if err := rpc.CallContext(context.Background(), &out, "evm_increaseTime", -5); err == nil {
		t.Error("negative seconds accepted")
	}
	if err := rpc.CallContext(context.Background(), &out, "evm_increaseTime", "soon"); err == nil {
		t.Error("non-numeric seconds accepted")
	}
	if err := rpc.CallContext(context.Background(), &out, "evm_increaseTime"); err == nil {
		t.Error("missing param accepted")
	}
}

func TestPendingBlockTagServesTip(t *testing.T) {
	c := New(t)
	c.EnsureBlock(7)

	ec := dialChain(t, c)
	rpc := ec.Client()

	var result map[string]any
	if err := rpc.CallContext(context.Background(), &result, "eth_getBlockByNumber", "pending", false); err != nil {
		t.Fatalf("eth_getBlockByNumber(pending): %v", err)
	}
	if result == nil {
		t.Fatal("pending block is null")
	}
	if got, want := result["number"], "0x7"; got != want {
		t.Errorf("pending block number = %v, want %v", got, want)
	}
	if _, ok := result["timestamp"].(string); !ok {
		t.Errorf("pending block timestamp missing or not a string: %v", result["timestamp"])
	}
}

func TestFailNextRPC(t *testing.T) {
	c := New(t)
	c.EnsureBlock(1)
	ec := dialChain(t, c)

	c.FailNextRPC("eth_gasPrice", "injected gas failure")
	if _, err := ec.SuggestGasPrice(context.Background()); err == nil ||
		!strings.Contains(err.Error(), "injected gas failure") {
		t.Errorf("first eth_gasPrice = %v, want injected failure", err)
	}
	// One-shot: the next call succeeds.
	if _, err := ec.SuggestGasPrice(context.Background()); err != nil {
		t.Errorf("second eth_gasPrice = %v, want success", err)
	}
}

func TestSetChainID(t *testing.T) {
	c := New(t)
	c.SetChainID(31337)

	ec := dialChain(t, c)
	id, err := ec.ChainID(context.Background())
	if err != nil {
		t.Fatalf("ChainID: %v", err)
	}
	if id.Int64() != 31337 {
		t.Errorf("chain id = %v, want 31337", id)
	}

	c.AddTx(1, c.Sender(), nil)
	defer func() {
		if recover() == nil {
			t.Error("SetChainID after AddTx did not panic")
		}
	}()
	c.SetChainID(1)
}
