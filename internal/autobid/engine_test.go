package autobid

// Engine tests: scripted CosmicGame rounds against the deterministic fake
// chain (internal/testchain + ContractStub) through the production wiring —
// real abigen bindings, EIP-155 signing, receipt polling. The Sleep seam is
// a no-op so scripted rounds run instantly; stub handlers script contract
// state as a function of how many transactions the engine has submitted.

import (
	"bytes"
	"context"
	"errors"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	rwc "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// Test key and deterministic contract addresses.
const testKeyHex = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

var (
	botAddr = func() common.Address {
		key, err := crypto.HexToECDSA(testKeyHex)
		if err != nil {
			panic(err)
		}
		return crypto.PubkeyToAddress(key.PublicKey)
	}()
	gameAddr   = common.HexToAddress("0x1000000000000000000000000000000000000001")
	rwalkAddr  = common.HexToAddress("0x2000000000000000000000000000000000000002")
	tokenAddr  = common.HexToAddress("0x3000000000000000000000000000000000000003")
	prizesAddr = common.HexToAddress("0x4000000000000000000000000000000000000004")
)

// gameWorld wires the four contract stubs onto a fake chain and provides
// mutable knobs the tests script against the submitted-transaction count.
type gameWorld struct {
	chain  *testchain.Chain
	game   *testchain.ContractStub
	rwalk  *testchain.ContractStub
	token  *testchain.ContractStub
	prizes *testchain.ContractStub

	// Scripted state (atomics: stub handlers run on the RPC server
	// goroutines).
	roundNum       atomic.Int64
	lastBidder     atomic.Pointer[common.Address]
	timeUntilPrize atomic.Int64
	cstPrice       atomic.Pointer[big.Int]
	ethBidPrice    atomic.Pointer[big.Int]
	rwalkMintPrice atomic.Pointer[big.Int]
	cstBalance     atomic.Pointer[big.Int]
	winner         atomic.Pointer[common.Address]
	prizeTime      atomic.Int64
	claimTimeout   atomic.Int64
}

func (w *gameWorld) setLastBidder(a common.Address) { w.lastBidder.Store(&a) }
func (w *gameWorld) setWinner(a common.Address)     { w.winner.Store(&a) }
func (w *gameWorld) setCstPrice(v *big.Int)         { w.cstPrice.Store(v) }
func (w *gameWorld) setEthBidPrice(v *big.Int)      { w.ethBidPrice.Store(v) }

// newGameWorld boots the fake chain with a funded bot account and all four
// contract stubs answering from the world's scripted state.
func newGameWorld(t *testing.T) *gameWorld {
	t.Helper()
	w := &gameWorld{chain: testchain.New(t)}
	w.chain.EnsureBlock(100)
	w.chain.SetBalance(botAddr, eth(50))
	w.chain.SetGasPrice(big.NewInt(1_000_000_000))

	w.setLastBidder(common.Address{})
	w.setWinner(common.Address{})
	w.setCstPrice(eth(20))
	w.setEthBidPrice(eth(0.05))
	w.rwalkMintPrice.Store(eth(0.02))
	w.cstBalance.Store(eth(100))
	w.timeUntilPrize.Store(1000)
	w.prizeTime.Store(1 << 40) // far future: no timeout claim by default
	w.claimTimeout.Store(3600)

	w.game = testchain.MustContractStub(cgc.CosmicSignatureGameABI)
	w.rwalk = testchain.MustContractStub(rwc.RWalkABI)
	w.token = testchain.MustContractStub(cgc.ERC20ABI)
	w.prizes = testchain.MustContractStub(cgc.PrizesWalletABI)

	w.game.Return("randomWalkNft", rwalkAddr)
	w.game.Return("prizesWallet", prizesAddr)
	w.game.Return("token", tokenAddr)
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		return []any{big.NewInt(w.roundNum.Load())}, nil
	})
	w.game.Handle("lastBidderAddress", func([]any) ([]any, error) {
		return []any{*w.lastBidder.Load()}, nil
	})
	w.game.Handle("getDurationUntilMainPrize", func([]any) ([]any, error) {
		return []any{big.NewInt(w.timeUntilPrize.Load())}, nil
	})
	w.game.Handle("getNextCstBidPrice", func([]any) ([]any, error) {
		return []any{new(big.Int).Set(w.cstPrice.Load())}, nil
	})
	w.game.Handle("getNextEthBidPrice", func([]any) ([]any, error) {
		return []any{new(big.Int).Set(w.ethBidPrice.Load())}, nil
	})
	w.game.Handle("mainPrizeTime", func([]any) ([]any, error) {
		return []any{big.NewInt(w.prizeTime.Load())}, nil
	})
	w.game.Handle("timeoutDurationToClaimMainPrize", func([]any) ([]any, error) {
		return []any{big.NewInt(w.claimTimeout.Load())}, nil
	})
	w.game.Return("usedRandomWalkNfts", big.NewInt(0))

	w.rwalk.Handle("getMintPrice", func([]any) ([]any, error) {
		return []any{new(big.Int).Set(w.rwalkMintPrice.Load())}, nil
	})
	w.rwalk.Return("nextTokenId", big.NewInt(0))

	w.token.Handle("balanceOf", func([]any) ([]any, error) {
		return []any{new(big.Int).Set(w.cstBalance.Load())}, nil
	})

	w.prizes.Handle("mainPrizeBeneficiaryAddresses", func([]any) ([]any, error) {
		return []any{*w.winner.Load()}, nil
	})

	w.chain.RegisterCall(gameAddr, w.game.Handler())
	w.chain.RegisterCall(rwalkAddr, w.rwalk.Handler())
	w.chain.RegisterCall(tokenAddr, w.token.Handler())
	w.chain.RegisterCall(prizesAddr, w.prizes.Handler())
	return w
}

// syncBuffer is a bytes.Buffer safe for concurrent writes by the engine and
// reads by the test goroutine.
type syncBuffer struct {
	mu      sync.Mutex
	buf     bytes.Buffer
	changed chan struct{}
}

func (b *syncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	n, err := b.buf.Write(p)
	b.mu.Unlock()
	if b.changed != nil {
		select {
		case b.changed <- struct{}{}:
		default:
		}
	}
	return n, err
}

func (b *syncBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

func (b *syncBuffer) waitFor(ctx context.Context, text string) bool {
	for {
		if strings.Contains(b.String(), text) {
			return true
		}
		select {
		case <-ctx.Done():
			return false
		case <-b.changed:
		}
	}
}

// newTestEngine builds an engine over the world with no-op sleeps and the
// output captured.
func newTestEngine(t *testing.T, w *gameWorld, mutate func(*Config)) (*Engine, *syncBuffer) {
	t.Helper()
	out := &syncBuffer{changed: make(chan struct{}, 1)}
	cfg := Config{
		RPCURL:        w.chain.URL(),
		PrivateKeyHex: testKeyHex,
		GameAddr:      gameAddr,
		Limits: Limits{
			MaxEthBid:       eth(5),
			MaxCstBid:       eth(9),
			RWalkMinPrice:   eth(0.1),
			TimeBeforePrize: 15,
			CstBidAnyway:    true,
		},
		Out:   out,
		Sleep: func(ctx context.Context, d time.Duration) error { return ctx.Err() },
	}
	if mutate != nil {
		mutate(&cfg)
	}
	e, err := New(context.Background(), cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	t.Cleanup(e.Close)
	return e, out
}

// runWithTimeout drives Run with a hard test deadline so a scripting mistake
// cannot hang the suite.
func runWithTimeout(t *testing.T, e *Engine, d time.Duration) error {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	err := e.Run(ctx)
	if ctx.Err() != nil && err == nil {
		t.Fatal("Run ended by test timeout, not by script — check the scenario")
	}
	return err
}

func TestEngineBidsEthWinsAndClaims(t *testing.T) {
	w := newGameWorld(t)
	// Script: nobody has bid, 10s until prize (inside TimeBeforePrize) →
	// first bid must be ETH. After the bid lands (1 tx), the bot is last
	// bidder with the timer at zero → claim (2 txs). After the claim, the
	// round advances and the bot is recorded as the round winner → exit.
	w.timeUntilPrize.Store(10)
	w.game.Handle("lastBidderAddress", func([]any) ([]any, error) {
		if w.chain.SubmittedTxCount() >= 1 {
			return []any{botAddr}, nil
		}
		return []any{common.Address{}}, nil
	})
	w.game.Handle("getDurationUntilMainPrize", func([]any) ([]any, error) {
		if w.chain.SubmittedTxCount() >= 1 {
			return []any{big.NewInt(0)}, nil
		}
		return []any{big.NewInt(10)}, nil
	})
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		if w.chain.SubmittedTxCount() >= 2 {
			return []any{big.NewInt(1)}, nil
		}
		return []any{big.NewInt(0)}, nil
	})
	w.setWinner(botAddr)

	e, out := newTestEngine(t, w, nil)
	if err := runWithTimeout(t, e, 30*time.Second); err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}

	if got := w.chain.SubmittedTxCount(); got != 2 {
		t.Errorf("submitted txs = %d, want 2 (bid + claim)", got)
	}
	if e.stats.ethBidCount != 1 {
		t.Errorf("ethBidCount = %d, want 1", e.stats.ethBidCount)
	}
	if e.stats.prizesClaimed != 1 {
		t.Errorf("prizesClaimed = %d, want 1", e.stats.prizesClaimed)
	}
	text := out.String()
	for _, want := range []string{
		"first bid of round - must use ETH",
		"ETH bid tx",
		"I am last bidder after ETH bid",
		"prize timer expired, claiming",
		"Prize claimed successfully!",
		"I am the winner of round 0!",
		"Round ended, exiting...",
		"SESSION SUMMARY",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("output missing %q\noutput:\n%s", want, text)
		}
	}
}

func TestEngineBidsCSTWhenCheap(t *testing.T) {
	w := newGameWorld(t)
	// Someone else is the last bidder and CST is cheap: CstBidAnyway bids
	// CST immediately. After the bid lands the round advances → exit.
	w.setLastBidder(common.HexToAddress("0x00000000000000000000000000000000000000cc"))
	w.setCstPrice(eth(5))
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		if w.chain.SubmittedTxCount() >= 1 {
			return []any{big.NewInt(1)}, nil
		}
		return []any{big.NewInt(0)}, nil
	})

	e, out := newTestEngine(t, w, nil)
	if err := runWithTimeout(t, e, 30*time.Second); err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}
	if e.stats.cstBidCount != 1 {
		t.Errorf("cstBidCount = %d, want 1", e.stats.cstBidCount)
	}
	if !strings.Contains(out.String(), "CST bid tx") {
		t.Errorf("output missing CST bid:\n%s", out.String())
	}
	if !strings.Contains(out.String(), "I am not the winner of round 0") {
		t.Errorf("output missing loser report:\n%s", out.String())
	}
}

func TestEngineRWalkBidWithPreOwnedToken(t *testing.T) {
	w := newGameWorld(t)
	// RWalk economics: ETH price 1 ETH (above the 0.1 floor), mint 0.1 →
	// mint + half = 0.6 < 1 → RWalk bid. The bot already owns token 7.
	w.setLastBidder(common.HexToAddress("0x00000000000000000000000000000000000000cc"))
	w.timeUntilPrize.Store(10)
	w.setEthBidPrice(eth(1))
	w.rwalkMintPrice.Store(eth(0.1))
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		if w.chain.SubmittedTxCount() >= 1 {
			return []any{big.NewInt(1)}, nil
		}
		return []any{big.NewInt(0)}, nil
	})

	e, out := newTestEngine(t, w, nil)
	e.nextRWalkTokenID.Store(7) // pre-owned unused token

	if err := runWithTimeout(t, e, 30*time.Second); err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}
	if e.stats.rwalkBidCount != 1 {
		t.Errorf("rwalkBidCount = %d, want 1", e.stats.rwalkBidCount)
	}
	if got := e.nextRWalkTokenID.Load(); got != -1 {
		t.Errorf("nextRWalkTokenID after bid = %d, want -1 (consumed)", got)
	}
	if !strings.Contains(out.String(), "RWalk bid tx") || !strings.Contains(out.String(), "(token 7)") {
		t.Errorf("output missing RWalk bid with token 7:\n%s", out.String())
	}
	// Half the 1 ETH price is charged to the stats.
	if e.stats.totalEthSpent.Cmp(eth(0.5)) != 0 {
		t.Errorf("totalEthSpent = %s, want 0.5 ETH", e.stats.totalEthSpent)
	}
}

func TestEngineRWalkMintPathExtractsTokenFromReceipt(t *testing.T) {
	w := newGameWorld(t)
	w.setLastBidder(common.HexToAddress("0x00000000000000000000000000000000000000cc"))
	w.timeUntilPrize.Store(10)
	w.setEthBidPrice(eth(1))
	w.rwalkMintPrice.Store(eth(0.1))
	// The first submitted tx is the mint: its receipt carries the RandomWalk
	// mint event for token 42.
	w.chain.SetMinedTxLogs(func(tx *types.Transaction, blockNum int64) []*types.Log {
		if w.chain == nil || tx.To() == nil || *tx.To() != rwalkAddr {
			return nil
		}
		return []*types.Log{{
			Address:     rwalkAddr,
			Topics:      []common.Hash{rwalkMintEventTopic, common.BigToHash(big.NewInt(42))},
			BlockNumber: uint64(blockNum), // #nosec G115 -- positive fake-chain block number
			TxHash:      tx.Hash(),
		}}
	})
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		if w.chain.SubmittedTxCount() >= 2 { // mint + bid
			return []any{big.NewInt(1)}, nil
		}
		return []any{big.NewInt(0)}, nil
	})

	e, out := newTestEngine(t, w, nil)
	if err := runWithTimeout(t, e, 30*time.Second); err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}
	text := out.String()
	for _, want := range []string{"Need to mint RWalk token first", "Minted RWalk token 42", "RWalk bid tx", "(token 42)"} {
		if !strings.Contains(text, want) {
			t.Errorf("output missing %q\noutput:\n%s", want, text)
		}
	}
	if e.stats.rwalkBidCount != 1 {
		t.Errorf("rwalkBidCount = %d, want 1", e.stats.rwalkBidCount)
	}
}

func TestEngineTimeoutClaim(t *testing.T) {
	w := newGameWorld(t)
	// Someone else is the last bidder and won, but never claimed: the
	// pending block timestamp is past mainPrizeTime + timeout, so the bot
	// claims the abandoned prize.
	w.setLastBidder(common.HexToAddress("0x00000000000000000000000000000000000000cc"))
	w.timeUntilPrize.Store(500) // not bidding time yet
	w.prizeTime.Store(1)        // far past
	w.claimTimeout.Store(1)
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		if w.chain.SubmittedTxCount() >= 1 {
			return []any{big.NewInt(1)}, nil
		}
		return []any{big.NewInt(0)}, nil
	})

	e, out := newTestEngine(t, w, nil)
	if err := runWithTimeout(t, e, 30*time.Second); err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}
	text := out.String()
	for _, want := range []string{
		"Winner didn't claim prize during claim window, I am going to claim it",
		"ClaimPrize tx",
		"Prize claimed successfully!",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("output missing %q\noutput:\n%s", want, text)
		}
	}
}

func TestEngineChainResetAborts(t *testing.T) {
	w := newGameWorld(t)
	w.roundNum.Store(5)

	e, out := newTestEngine(t, w, nil)
	// After startup (round 5), the chain "resets" to round 1.
	w.roundNum.Store(1)

	err := runWithTimeout(t, e, 30*time.Second)
	if !errors.Is(err, ErrChainReset) {
		t.Fatalf("Run = %v, want ErrChainReset", err)
	}
	if !strings.Contains(out.String(), "blockchain was reset") {
		t.Errorf("output missing reset report:\n%s", out.String())
	}
}

func TestEngineContextCancellationStopsRun(t *testing.T) {
	w := newGameWorld(t)
	sleepEntered := make(chan struct{}, 1)
	e, out := newTestEngine(t, w, func(cfg *Config) {
		cfg.Sleep = func(ctx context.Context, _ time.Duration) error {
			select {
			case sleepEntered <- struct{}{}:
			default:
			}
			<-ctx.Done()
			return ctx.Err()
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- e.Run(ctx) }()
	select {
	case <-sleepEntered:
		cancel()
	case <-time.After(5 * time.Second):
		t.Fatal("Run did not reach its cancellable loop wait")
	}

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Run after cancel = %v, want nil", err)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("Run did not stop after context cancellation")
	}
	if !strings.Contains(out.String(), "SESSION SUMMARY") {
		t.Errorf("cancelled run should print the session summary:\n%s", out.String())
	}
	if !strings.Contains(out.String(), "Current Balance:   50.000000000000000000 ETH") {
		t.Errorf("cancelled run should report the last known balance:\n%s", out.String())
	}
}

func TestEngineReconnectsAfterConsecutiveErrors(t *testing.T) {
	w := newGameWorld(t)
	// getNextCstBidPrice fails until the engine has reconnected once.
	var dials atomic.Int32
	var failing atomic.Bool
	failing.Store(true)
	w.game.Handle("getNextCstBidPrice", func([]any) ([]any, error) {
		if failing.Load() {
			return nil, errors.New("injected refresh failure")
		}
		return []any{new(big.Int).Set(w.cstPrice.Load())}, nil
	})
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		if !failing.Load() {
			return []any{big.NewInt(1)}, nil // heal → round over → exit
		}
		return []any{big.NewInt(0)}, nil
	})

	e, out := newTestEngine(t, w, func(cfg *Config) {
		cfg.Dial = func(ctx context.Context, url string) (*ethrpc.Client, error) {
			if dials.Add(1) >= 2 {
				failing.Store(false) // the reconnect "fixes" the node
			}
			return ethrpc.DialContext(ctx, url)
		}
	})
	if err := runWithTimeout(t, e, 30*time.Second); err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}
	if dials.Load() < 2 {
		t.Errorf("dial count = %d, want >= 2 (a reconnect)", dials.Load())
	}
	if !strings.Contains(out.String(), "[RECONNECT] Successfully reconnected to RPC") {
		t.Errorf("output missing reconnect report:\n%s", out.String())
	}
	if e.stats.reconnectCount < 1 {
		t.Errorf("stats.reconnectCount = %d, want >= 1", e.stats.reconnectCount)
	}
}

func TestEngineReconnectChainIDMismatchIsFatal(t *testing.T) {
	w := newGameWorld(t)
	// A second chain with a different chain id stands in after the failure.
	wrongChain, stopWrong := testchain.Start()
	t.Cleanup(stopWrong)
	wrongChain.SetChainID(999)
	wrongChain.EnsureBlock(1)

	w.game.Handle("getNextCstBidPrice", func([]any) ([]any, error) {
		return nil, errors.New("node gone")
	})

	var dials atomic.Int32
	e, out := newTestEngine(t, w, func(cfg *Config) {
		cfg.Dial = func(ctx context.Context, url string) (*ethrpc.Client, error) {
			if dials.Add(1) >= 2 {
				return ethrpc.DialContext(ctx, wrongChain.URL())
			}
			return ethrpc.DialContext(ctx, url)
		}
	})

	err := runWithTimeout(t, e, 60*time.Second)
	if err == nil || !strings.Contains(err.Error(), "max reconnection attempts reached") {
		t.Fatalf("Run = %v, want reconnect exhaustion", err)
	}
	if !strings.Contains(out.String(), "chain ID mismatch after reconnect") {
		t.Errorf("output missing chain-id mismatch:\n%s", out.String())
	}
}

func TestEngineInitialBiddingOnNewRound(t *testing.T) {
	w := newGameWorld(t)
	// Round 0 ends immediately (round handler jumps to 1); with
	// InitialBidPrice set the engine bids the new round up to the price
	// level: prices rise 0.01 → 0.02 → 0.03; limit 0.025 stops after 2 bids.
	w.setLastBidder(common.HexToAddress("0x00000000000000000000000000000000000000cc"))

	var roundCalls atomic.Int64
	w.game.Handle("getNextEthBidPrice", func([]any) ([]any, error) {
		// Rise by 0.01 ETH per submitted bid.
		bids := w.chain.SubmittedTxCount()
		return []any{eth(0.01 * float64(bids+1))}, nil
	})
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		// The engine's startup read sees round 0; every later read sees the
		// next round, triggering the initial-bidding opener.
		if roundCalls.Add(1) == 1 {
			return []any{big.NewInt(0)}, nil
		}
		return []any{big.NewInt(1)}, nil
	})

	e, out := newTestEngine(t, w, func(cfg *Config) {
		cfg.InitialBidPrice = eth(0.025)
		// Cancel the run once initial bidding is done: bound loop count.
		var sleeps atomic.Int64
		cfg.Sleep = func(ctx context.Context, d time.Duration) error {
			if sleeps.Add(1) > 500 {
				return context.Canceled
			}
			return ctx.Err()
		}
	})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	done := make(chan error, 1)
	go func() { done <- e.Run(ctx) }()
	if !out.waitFor(ctx, "Initial bidding finished") {
		t.Fatalf("initial bidding did not finish:\n%s", out.String())
	}
	cancel()
	if err := <-done; err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}

	text := out.String()
	for _, want := range []string{
		"Playing new round with initial bids",
		"Initial bid #1 tx",
		"Initial bid #2 tx",
		"Reached bid price limit, stopping initial bidding",
		"Initial bidding finished: 2 bids",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("output missing %q\noutput:\n%s", want, text)
		}
	}
	if got := w.chain.SubmittedTxCount(); got != 2 {
		t.Errorf("submitted txs = %d, want 2 initial bids", got)
	}
}

func TestEngineAbandonsPendingTxAfterRetries(t *testing.T) {
	w := newGameWorld(t)
	w.setLastBidder(common.HexToAddress("0x00000000000000000000000000000000000000cc"))
	w.timeUntilPrize.Store(10)
	w.chain.MarkNextTxPending() // the bid's receipt never arrives
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		// End the round after the engine gives up on the receipt (1 tx
		// submitted, then 3 failed polls; use the poll counter indirectly by
		// ending once a tx exists and the pending state cleared).
		if w.chain.SubmittedTxCount() >= 1 {
			return []any{big.NewInt(1)}, nil
		}
		return []any{big.NewInt(0)}, nil
	})

	e, out := newTestEngine(t, w, nil)
	if err := runWithTimeout(t, e, 30*time.Second); err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}
	if !strings.Contains(out.String(), "Max retries reached for ETHBid tx") {
		t.Errorf("output missing receipt-retry abandonment:\n%s", out.String())
	}
	if e.stats.ethBidCount != 0 {
		t.Errorf("ethBidCount = %d, want 0 (no confirmed bid)", e.stats.ethBidCount)
	}
}

func TestEngineFailedBidTxCountsAsFailed(t *testing.T) {
	w := newGameWorld(t)
	w.setLastBidder(common.HexToAddress("0x00000000000000000000000000000000000000cc"))
	w.timeUntilPrize.Store(10)
	w.chain.MarkNextTxReverted()
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		if w.chain.SubmittedTxCount() >= 1 {
			return []any{big.NewInt(1)}, nil
		}
		return []any{big.NewInt(0)}, nil
	})

	e, out := newTestEngine(t, w, nil)
	if err := runWithTimeout(t, e, 30*time.Second); err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}
	if e.stats.failedTxCount != 1 {
		t.Errorf("failedTxCount = %d, want 1", e.stats.failedTxCount)
	}
	if !strings.Contains(out.String(), "ETHBid tx failed") {
		t.Errorf("output missing failed-tx report:\n%s", out.String())
	}
}

func TestFindRWalkTokenID(t *testing.T) {
	w := newGameWorld(t)
	// Tokens 0..2 exist; the bot owns token 1 and it is unused.
	w.rwalk.Return("nextTokenId", big.NewInt(3))
	w.rwalk.Handle("ownerOf", func(args []any) ([]any, error) {
		id, _ := args[0].(*big.Int)
		if id.Int64() == 1 {
			return []any{botAddr}, nil
		}
		return []any{common.HexToAddress("0x00000000000000000000000000000000000000dd")}, nil
	})

	e, _ := newTestEngine(t, w, nil)
	e.findRWalkTokenID(t.Context())
	if got := e.nextRWalkTokenID.Load(); got != 1 {
		t.Errorf("nextRWalkTokenID = %d, want 1", got)
	}

	// A found token that has since been used resets and re-searches.
	w.game.Return("usedRandomWalkNfts", big.NewInt(1)) // everything used now
	e.findRWalkTokenID(t.Context())
	if got := e.nextRWalkTokenID.Load(); got != -1 {
		t.Errorf("nextRWalkTokenID after all used = %d, want -1", got)
	}
}

func TestNewEngineFailures(t *testing.T) {
	w := newGameWorld(t)
	t.Run("bad key", func(t *testing.T) {
		_, err := New(context.Background(), Config{
			RPCURL: w.chain.URL(), PrivateKeyHex: "zz", GameAddr: gameAddr,
			Out: &bytes.Buffer{},
		})
		if err == nil || !strings.Contains(err.Error(), "private key") {
			t.Errorf("New with bad key = %v", err)
		}
	})
	t.Run("unreachable rpc", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_, err := New(ctx, Config{
			RPCURL: "http://127.0.0.1:1", PrivateKeyHex: testKeyHex, GameAddr: gameAddr,
			Out: &bytes.Buffer{},
		})
		if err == nil {
			t.Error("New against unreachable RPC succeeded")
		}
	})
}
