package ethtx

import (
	"bytes"
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

func TestIsDevChain(t *testing.T) {
	if (&Network{ChainID: big.NewInt(1)}).IsDevChain() {
		t.Error("mainnet id reported as dev chain")
	}
	if !(&Network{ChainID: big.NewInt(HardhatChainID)}).IsDevChain() {
		t.Error("hardhat id not reported as dev chain")
	}
	if (&Network{}).IsDevChain() {
		t.Error("nil chain id reported as dev chain")
	}
}

func TestAdvanceDevChainTime(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(5)
	net, err := Connect(context.Background(), chain.URL())
	if err != nil {
		t.Fatalf("Connect: %v", err)
	}
	blockBefore := net.BlockNum.Int64()
	timeBefore := net.BlockTime

	if err := net.AdvanceDevChainTime(context.Background(), 600); err != nil {
		t.Fatalf("AdvanceDevChainTime: %v", err)
	}
	if net.BlockNum.Int64() != blockBefore+1 {
		t.Errorf("block after advance = %d, want %d", net.BlockNum.Int64(), blockBefore+1)
	}
	// The new block carries the base progression (+100) plus the offset.
	if want := timeBefore + 100 + 600; net.BlockTime != want {
		t.Errorf("block time after advance = %d, want %d", net.BlockTime, want)
	}

	// Zero and negative are no-ops.
	if err := net.AdvanceDevChainTime(context.Background(), 0); err != nil {
		t.Errorf("zero seconds: %v", err)
	}
	if err := net.AdvanceDevChainTime(context.Background(), -5); err != nil {
		t.Errorf("negative seconds: %v", err)
	}
	if got := chain.TimeOffset(); got != 600 {
		t.Errorf("chain offset = %d, want 600 (no-ops must not accumulate)", got)
	}
}

func TestAdvanceDevChainTimeUnsupportedNode(t *testing.T) {
	// The failingRPC helper answers eth_* but not evm_increaseTime, like a
	// real (non-dev) node.
	net, err := Connect(context.Background(), failingRPC(t, "nothing"))
	if err != nil {
		t.Fatalf("Connect: %v", err)
	}
	if err := net.AdvanceDevChainTime(context.Background(), 10); err == nil {
		t.Error("AdvanceDevChainTime against a node without evm_increaseTime succeeded")
	}
}

func TestSessionGasPriceMultiplier(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	chain.SetBalance(testKeyAddr, big.NewInt(1_000_000_000_000_000_000))
	chain.SetGasPrice(big.NewInt(1_000_000_000))

	newSession := func(mult float64) *Session {
		s, err := New(context.Background(), Options{
			RPCURL:             chain.URL(),
			PrivateKeyHex:      testKeyHex,
			Out:                &bytes.Buffer{},
			GasPriceMultiplier: mult,
		})
		if err != nil {
			t.Fatalf("New(mult=%v): %v", mult, err)
		}
		return s
	}

	// Default (unset) applies the legacy 2.0.
	if got := newSession(0).AdjustedGasPrice(); got.Cmp(big.NewInt(2_000_000_000)) != 0 {
		t.Errorf("default multiplier price = %s, want 2 gwei", got)
	}
	// Custom multiplier flows into TransactOpts.
	s := newSession(1.5)
	if got := s.AdjustedGasPrice(); got.Cmp(big.NewInt(1_500_000_000)) != 0 {
		t.Errorf("1.5 multiplier price = %s, want 1.5 gwei", got)
	}
	to := common.HexToAddress("0x2200000000000000000000000000000000000022")
	opts := s.TransactOpts(nil, GasLimitContractCall)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    opts.Nonce.Uint64(),
		GasPrice: opts.GasPrice,
		Gas:      opts.GasLimit,
		To:       &to,
	})
	signed, err := opts.Signer(opts.From, tx)
	if err != nil {
		t.Fatalf("signing: %v", err)
	}
	if signed.GasPrice().Cmp(big.NewInt(1_500_000_000)) != 0 {
		t.Errorf("signed gas price = %s, want 1.5 gwei", signed.GasPrice())
	}
}

func TestSessionRefresh(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	chain.SetBalance(testKeyAddr, big.NewInt(1_000_000_000_000_000_000))
	chain.SetNonce(testKeyAddr, 1)
	chain.SetGasPrice(big.NewInt(1_000_000_000))

	s, err := New(context.Background(), Options{
		RPCURL:        chain.URL(),
		PrivateKeyHex: testKeyHex,
		Out:           &bytes.Buffer{},
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	chain.EnsureBlock(5)
	chain.SetNonce(testKeyAddr, 7)
	chain.SetBalance(testKeyAddr, big.NewInt(42))
	chain.SetGasPrice(big.NewInt(3_000_000_000))

	if err := s.Refresh(context.Background()); err != nil {
		t.Fatalf("Refresh: %v", err)
	}
	if s.Net.BlockNum.Int64() != 5 {
		t.Errorf("refreshed block = %d, want 5", s.Net.BlockNum.Int64())
	}
	if s.Acc.Nonce != 7 {
		t.Errorf("refreshed nonce = %d, want 7", s.Acc.Nonce)
	}
	if s.Acc.Balance.Cmp(big.NewInt(42)) != 0 {
		t.Errorf("refreshed balance = %s, want 42", s.Acc.Balance)
	}
	if s.Net.GasPrice.Cmp(big.NewInt(3_000_000_000)) != 0 {
		t.Errorf("refreshed gas price = %s, want 3 gwei", s.Net.GasPrice)
	}
	// The refreshed nonce flows into the next TransactOpts.
	if got := s.TransactOpts(nil, GasLimitContractCall).Nonce.Uint64(); got != 7 {
		t.Errorf("TransactOpts nonce after refresh = %d, want 7", got)
	}
}

func TestSessionRefreshPerCallFailures(t *testing.T) {
	newSession := func(t *testing.T) (*testchain.Chain, *Session) {
		t.Helper()
		chain := testchain.New(t)
		chain.EnsureBlock(1)
		chain.SetBalance(testKeyAddr, big.NewInt(1_000_000_000_000_000_000))
		s, err := New(context.Background(), Options{
			RPCURL:        chain.URL(),
			PrivateKeyHex: testKeyHex,
			Out:           &bytes.Buffer{},
		})
		if err != nil {
			t.Fatalf("New: %v", err)
		}
		return chain, s
	}
	cases := []struct {
		rpc  string
		want string
	}{
		{"eth_gasPrice", "refreshing gas price"},
		{"eth_getBlockByNumber", "refreshing latest block"},
		{"eth_getTransactionCount", "refreshing account nonce"},
		{"eth_getBalance", "refreshing account balance"},
	}
	for _, tc := range cases {
		t.Run(tc.rpc, func(t *testing.T) {
			chain, s := newSession(t)
			chain.FailNextRPC(tc.rpc, "injected failure")
			if err := s.Refresh(context.Background()); err == nil ||
				!strings.Contains(err.Error(), tc.want) {
				t.Errorf("Refresh with failing %s = %v, want %q", tc.rpc, err, tc.want)
			}
		})
	}
}

func TestAdvanceDevChainTimePerCallFailures(t *testing.T) {
	newNet := func(t *testing.T) (*testchain.Chain, *Network) {
		t.Helper()
		chain := testchain.New(t)
		chain.EnsureBlock(1)
		net, err := Connect(context.Background(), chain.URL())
		if err != nil {
			t.Fatalf("Connect: %v", err)
		}
		return chain, net
	}
	t.Run("evm_mine fails", func(t *testing.T) {
		chain, net := newNet(t)
		chain.FailNextRPC("evm_mine", "mining off")
		if err := net.AdvanceDevChainTime(context.Background(), 10); err == nil ||
			!strings.Contains(err.Error(), "evm_mine") {
			t.Errorf("AdvanceDevChainTime = %v, want evm_mine failure", err)
		}
	})
	t.Run("header refresh fails", func(t *testing.T) {
		chain, net := newNet(t)
		chain.FailNextRPC("eth_getBlockByNumber", "header gone")
		if err := net.AdvanceDevChainTime(context.Background(), 10); err == nil ||
			!strings.Contains(err.Error(), "refresh block header") {
			t.Errorf("AdvanceDevChainTime = %v, want header failure", err)
		}
	})
}

func TestNetworkBalance(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	addr := common.HexToAddress("0x3300000000000000000000000000000000000033")
	chain.SetBalance(addr, big.NewInt(777))
	net, err := Connect(context.Background(), chain.URL())
	if err != nil {
		t.Fatalf("Connect: %v", err)
	}
	bal, err := net.Balance(context.Background(), addr)
	if err != nil {
		t.Fatalf("Balance: %v", err)
	}
	if bal.Cmp(big.NewInt(777)) != 0 {
		t.Errorf("balance = %s, want 777", bal)
	}
}
