package autobid

// Error-branch tests: every per-call failure of engine construction, market
// refresh, transaction submission and the timeout-claim probe, driven
// through partial contract stubs and the fake chain's failure injectors.

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

func TestNewEngineContractFailures(t *testing.T) {
	cases := []struct {
		name    string
		prepare func(w *gameWorld)
		want    string
	}{
		{
			name: "randomWalkNft read fails",
			prepare: func(w *gameWorld) {
				w.game.Handle("randomWalkNft", func([]any) ([]any, error) {
					return nil, errors.New("no rwalk")
				})
			},
			want: "getting RWalk addr",
		},
		{
			name: "prizesWallet read fails",
			prepare: func(w *gameWorld) {
				w.game.Handle("prizesWallet", func([]any) ([]any, error) {
					return nil, errors.New("no wallet")
				})
			},
			want: "fetching PrizesWallet address",
		},
		{
			name: "roundNum read fails",
			prepare: func(w *gameWorld) {
				w.game.Handle("roundNum", func([]any) ([]any, error) {
					return nil, errors.New("no round")
				})
			},
			want: "getting roundNum",
		},
		{
			name: "initial balance read fails",
			prepare: func(w *gameWorld) {
				w.chain.FailNextRPC("eth_getBalance", "balance down")
			},
			want: "getting initial balance",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := newGameWorld(t)
			tc.prepare(w)
			_, err := New(context.Background(), Config{
				RPCURL: w.chain.URL(), PrivateKeyHex: testKeyHex, GameAddr: gameAddr,
				Out: &syncBuffer{},
			})
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("New = %v, want %q", err, tc.want)
			}
		})
	}
}

func TestNewEngineChainIDFailure(t *testing.T) {
	w := newGameWorld(t)
	w.chain.FailNextRPC("eth_chainId", "id service down")
	_, err := New(context.Background(), Config{
		RPCURL: w.chain.URL(), PrivateKeyHex: testKeyHex, GameAddr: gameAddr,
		Out: &syncBuffer{},
	})
	if err == nil || !strings.Contains(err.Error(), "getting chain id") {
		t.Errorf("New = %v, want chain-id failure", err)
	}
}

func TestRefreshMarketPerCallFailures(t *testing.T) {
	cases := []struct {
		name    string
		prepare func(w *gameWorld)
		want    string
	}{
		{"cst price", func(w *gameWorld) {
			w.game.Handle("getNextCstBidPrice", func([]any) ([]any, error) { return nil, errors.New("x") })
		}, "getting CST bid price"},
		{"eth price", func(w *gameWorld) {
			w.game.Handle("getNextEthBidPrice", func([]any) ([]any, error) { return nil, errors.New("x") })
		}, "getting ETH bid price"},
		{"time until prize", func(w *gameWorld) {
			w.game.Handle("getDurationUntilMainPrize", func([]any) ([]any, error) { return nil, errors.New("x") })
		}, "getting time until prize"},
		{"token address", func(w *gameWorld) {
			w.game.Handle("token", func([]any) ([]any, error) { return nil, errors.New("x") })
		}, "getting token address"},
		{"cst balance", func(w *gameWorld) {
			w.token.Handle("balanceOf", func([]any) ([]any, error) { return nil, errors.New("x") })
		}, "getting CST balance"},
		{"last bidder", func(w *gameWorld) {
			w.game.Handle("lastBidderAddress", func([]any) ([]any, error) { return nil, errors.New("x") })
		}, "getting last bidder"},
		{"eth balance", func(w *gameWorld) {
			w.chain.FailNextRPC("eth_getBalance", "x")
		}, "getting ETH balance"},
		{"rwalk mint price", func(w *gameWorld) {
			w.rwalk.Handle("getMintPrice", func([]any) ([]any, error) { return nil, errors.New("x") })
		}, "getting RWalk mint price"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := newGameWorld(t)
			e, _ := newTestEngine(t, w, nil)
			tc.prepare(w)
			err := e.refreshMarket(context.Background())
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("refreshMarket = %v, want %q", err, tc.want)
			}
		})
	}
}

func TestCheckTimeoutClaimProbeFailures(t *testing.T) {
	w := newGameWorld(t)
	e, _ := newTestEngine(t, w, nil)
	e.market = Market{LastBidder: otherAddrEng()}

	t.Run("zero last bidder skips", func(t *testing.T) {
		e.market.LastBidder = common.Address{}
		if e.checkTimeoutClaim(context.Background()) {
			t.Error("timeout claim with no bidder")
		}
		e.market.LastBidder = otherAddrEng()
	})
	t.Run("timeout read fails", func(t *testing.T) {
		w.game.Handle("timeoutDurationToClaimMainPrize", func([]any) ([]any, error) { return nil, errors.New("x") })
		if e.checkTimeoutClaim(context.Background()) {
			t.Error("timeout claim despite failed timeout read")
		}
		w.game.Handle("timeoutDurationToClaimMainPrize", func([]any) ([]any, error) {
			return []any{big.NewInt(w.claimTimeout.Load())}, nil
		})
	})
	t.Run("prize time read fails", func(t *testing.T) {
		w.game.Handle("mainPrizeTime", func([]any) ([]any, error) { return nil, errors.New("x") })
		if e.checkTimeoutClaim(context.Background()) {
			t.Error("timeout claim despite failed prize-time read")
		}
		w.game.Handle("mainPrizeTime", func([]any) ([]any, error) {
			return []any{big.NewInt(w.prizeTime.Load())}, nil
		})
	})
	t.Run("pending block read fails", func(t *testing.T) {
		w.chain.FailNextRPC("eth_getBlockByNumber", "x")
		if e.checkTimeoutClaim(context.Background()) {
			t.Error("timeout claim despite failed block read")
		}
	})
	t.Run("window still open", func(t *testing.T) {
		// prizeTime is far future by default: not expired.
		if e.checkTimeoutClaim(context.Background()) {
			t.Error("timeout claim before the window expired")
		}
	})
}

// otherAddrEng avoids clashing with the decide_test other var.
func otherAddrEng() common.Address {
	return common.HexToAddress("0x00000000000000000000000000000000000000cc")
}

func TestSubmissionErrorBranches(t *testing.T) {
	// Every bid/claim submission error path: the node rejects the send once
	// (error branch logged, nothing pending), the next loop iteration
	// succeeds and the round ends.
	cases := []struct {
		name    string
		mutate  func(w *gameWorld, e *Engine)
		wantLog string
	}{
		{
			name: "eth bid rejected",
			mutate: func(w *gameWorld, e *Engine) {
				w.timeUntilPrize.Store(10)
				w.setLastBidder(otherAddrEng())
			},
			wantLog: "BidWithEth error:",
		},
		{
			name: "cst bid rejected",
			mutate: func(w *gameWorld, e *Engine) {
				w.setCstPrice(eth(5))
				w.setLastBidder(otherAddrEng())
			},
			wantLog: "BidWithCST error:",
		},
		{
			name: "rwalk bid rejected",
			mutate: func(w *gameWorld, e *Engine) {
				w.timeUntilPrize.Store(10)
				w.setLastBidder(otherAddrEng())
				w.setEthBidPrice(eth(1))
				w.rwalkMintPrice.Store(eth(0.1))
				e.nextRWalkTokenID.Store(3)
			},
			wantLog: "BidWithEth (RWalk) error:",
		},
		{
			name: "rwalk mint rejected",
			mutate: func(w *gameWorld, e *Engine) {
				w.timeUntilPrize.Store(10)
				w.setLastBidder(otherAddrEng())
				w.setEthBidPrice(eth(1))
				w.rwalkMintPrice.Store(eth(0.1))
			},
			wantLog: "RWalk mint error:",
		},
		{
			name: "claim rejected",
			mutate: func(w *gameWorld, e *Engine) {
				w.timeUntilPrize.Store(0)
				w.setLastBidder(botAddr)
			},
			wantLog: "ClaimMainPrize error:",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := newGameWorld(t)
			w.game.Handle("roundNum", func([]any) ([]any, error) {
				if w.chain.SubmittedTxCount() >= 1 {
					return []any{big.NewInt(1)}, nil
				}
				return []any{big.NewInt(0)}, nil
			})
			e, out := newTestEngine(t, w, nil)
			tc.mutate(w, e)
			w.chain.RejectNextSendWith("injected send rejection")

			if err := runWithTimeout(t, e, 30*time.Second); err != nil {
				t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
			}
			if !strings.Contains(out.String(), tc.wantLog) {
				t.Errorf("output missing %q:\n%s", tc.wantLog, out.String())
			}
			if !strings.Contains(out.String(), "injected send rejection") {
				t.Errorf("output missing the node's rejection message:\n%s", out.String())
			}
		})
	}
}

func TestBidFunctionsTxOptsFailures(t *testing.T) {
	// Every submission helper logs and stays idle when transaction options
	// cannot be built (nonce service down).
	w := newGameWorld(t)
	e, out := newTestEngine(t, w, nil)
	e.market = Market{
		CstPrice:       eth(5),
		EthBidPrice:    eth(0.05),
		RWalkMintPrice: eth(0.02),
	}
	ctx := context.Background()

	run := func(name string, fn func()) {
		w.chain.FailNextRPC("eth_getTransactionCount", "nonce down for "+name)
		fn()
		if e.pendingKind != pendingNone {
			t.Errorf("%s left a pending tx: %s", name, e.pendingKind)
		}
		if !strings.Contains(out.String(), "nonce down for "+name) {
			t.Errorf("%s output missing the nonce failure:\n%s", name, out.String())
		}
	}

	run("cst", func() { e.bidCST(ctx) })
	run("eth", func() { e.bidETH(ctx) })
	e.nextRWalkTokenID.Store(-1)
	run("mint", func() { e.bidRWalk(ctx) })
	e.nextRWalkTokenID.Store(4)
	run("rwalkbid", func() { e.sendRWalkBid(ctx) })
	run("claim", func() { e.claimPrize(ctx) })

	if got := w.chain.SubmittedTxCount(); got != 0 {
		t.Errorf("submitted txs = %d, want 0", got)
	}
}

func TestLogAfterBidReadFailure(t *testing.T) {
	w := newGameWorld(t)
	e, out := newTestEngine(t, w, nil)
	w.game.Handle("lastBidderAddress", func([]any) ([]any, error) {
		return nil, errors.New("bidder read down")
	})
	e.logAfterBid("ETH")
	if !strings.Contains(out.String(), "Error checking last bidder: ") {
		t.Errorf("output missing bidder read failure:\n%s", out.String())
	}
}

func TestCreateTransactOptsFailures(t *testing.T) {
	w := newGameWorld(t)
	e, _ := newTestEngine(t, w, nil)

	w.chain.FailNextRPC("eth_getTransactionCount", "nonce down")
	if _, err := e.createTransactOpts(context.Background(), nil, bidGasLimit); err == nil ||
		!strings.Contains(err.Error(), "getting nonce") {
		t.Errorf("nonce failure = %v", err)
	}

	w.chain.FailNextRPC("eth_gasPrice", "gas down")
	if _, err := e.createTransactOpts(context.Background(), nil, bidGasLimit); err == nil ||
		!strings.Contains(err.Error(), "getting gas price") {
		t.Errorf("gas price failure = %v", err)
	}
}

func TestMintReceiptWithoutEventLogged(t *testing.T) {
	w := newGameWorld(t)
	w.setLastBidder(otherAddrEng())
	w.timeUntilPrize.Store(10)
	w.setEthBidPrice(eth(1))
	w.rwalkMintPrice.Store(eth(0.1))
	// No SetMinedTxLogs: the mint receipt carries no logs.
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
	if !strings.Contains(out.String(), "RWalk mint event not found") {
		t.Errorf("output missing mint-event miss:\n%s", out.String())
	}
}

func TestFindRWalkTokenIDErrorPaths(t *testing.T) {
	w := newGameWorld(t)
	e, _ := newTestEngine(t, w, nil)

	t.Run("nextTokenId fails", func(t *testing.T) {
		w.rwalk.Handle("nextTokenId", func([]any) ([]any, error) { return nil, errors.New("x") })
		e.findRWalkTokenID()
		if got := e.nextRWalkTokenID.Load(); got != -1 {
			t.Errorf("token after failed nextTokenId = %d", got)
		}
		w.rwalk.Return("nextTokenId", big.NewInt(3))
	})

	t.Run("ownerOf fails aborts scan", func(t *testing.T) {
		w.rwalk.Handle("ownerOf", func([]any) ([]any, error) { return nil, errors.New("x") })
		e.prevRWalkTokenID.Store(-1)
		e.findRWalkTokenID()
		if got := e.nextRWalkTokenID.Load(); got != -1 {
			t.Errorf("token after failed ownerOf = %d", got)
		}
	})

	t.Run("concurrent search is single flight", func(t *testing.T) {
		e.rwalkSearching.Store(true)
		e.findRWalkTokenID() // returns immediately
		e.rwalkSearching.Store(false)
	})
}

func TestReconnectBindFailure(t *testing.T) {
	// After a reconnect the game contract is gone (a fresh chain without the
	// stub): bindContracts fails and the engine keeps retrying until the
	// attempt cap, then exits with the reconnect-exhaustion error.
	w := newGameWorld(t)
	bare, stopBare := testchain.Start()
	t.Cleanup(stopBare)
	bare.EnsureBlock(1)

	w.game.Handle("getNextCstBidPrice", func([]any) ([]any, error) {
		return nil, errors.New("node degraded")
	})

	var dials atomic.Int32
	e, out := newTestEngine(t, w, func(cfg *Config) {
		cfg.Dial = func(ctx context.Context, url string) (*ethrpc.Client, error) {
			if dials.Add(1) >= 2 {
				return ethrpc.DialContext(ctx, bare.URL())
			}
			return ethrpc.DialContext(ctx, url)
		}
	})

	err := runWithTimeout(t, e, 60*time.Second)
	if err == nil || !strings.Contains(err.Error(), "max reconnection attempts reached") {
		t.Fatalf("Run = %v, want reconnect exhaustion", err)
	}
	if !strings.Contains(out.String(), "reconnect failed") {
		t.Errorf("output missing bind failure:\n%s", out.String())
	}
}

func TestReconnectDialFailure(t *testing.T) {
	w := newGameWorld(t)
	w.game.Handle("getNextCstBidPrice", func([]any) ([]any, error) {
		return nil, errors.New("node gone")
	})

	var dials atomic.Int32
	e, out := newTestEngine(t, w, func(cfg *Config) {
		cfg.Dial = func(ctx context.Context, url string) (*ethrpc.Client, error) {
			if dials.Add(1) >= 2 {
				return nil, errors.New("connection refused")
			}
			return ethrpc.DialContext(ctx, url)
		}
	})

	err := runWithTimeout(t, e, 60*time.Second)
	if err == nil || !strings.Contains(err.Error(), "max reconnection attempts reached") {
		t.Fatalf("Run = %v, want reconnect exhaustion", err)
	}
	if !strings.Contains(out.String(), "can't connect to ETH RPC") {
		t.Errorf("output missing dial failure:\n%s", out.String())
	}
}

func TestRoundChangeReadFailureContinues(t *testing.T) {
	// A failed roundNum read logs and keeps the loop alive; the next read
	// ends the round normally.
	w := newGameWorld(t)
	var calls atomic.Int64
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		switch calls.Add(1) {
		case 1: // engine startup
			return []any{big.NewInt(0)}, nil
		case 2: // first loop check fails
			return nil, errors.New("round read glitch")
		default: // then the round is over
			return []any{big.NewInt(1)}, nil
		}
	})

	e, out := newTestEngine(t, w, nil)
	if err := runWithTimeout(t, e, 30*time.Second); err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}
	if !strings.Contains(out.String(), "Error getting roundNum: ") {
		t.Errorf("output missing round-read failure:\n%s", out.String())
	}
	if !strings.Contains(out.String(), "Round ended, exiting...") {
		t.Errorf("output missing round end:\n%s", out.String())
	}
}

func TestWinnerCheckFailureStillExits(t *testing.T) {
	// The MainPrizeBeneficiaryAddresses read failing must not block the
	// round-end exit.
	w := newGameWorld(t)
	w.prizes.Handle("mainPrizeBeneficiaryAddresses", func([]any) ([]any, error) {
		return nil, errors.New("no winner data")
	})
	var calls atomic.Int64
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		if calls.Add(1) == 1 {
			return []any{big.NewInt(0)}, nil
		}
		return []any{big.NewInt(1)}, nil
	})

	e, out := newTestEngine(t, w, nil)
	if err := runWithTimeout(t, e, 30*time.Second); err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}
	if strings.Contains(out.String(), "I am the winner") || strings.Contains(out.String(), "I am not the winner") {
		t.Errorf("winner report printed despite failed read:\n%s", out.String())
	}
}

func TestNewEngineDialFailure(t *testing.T) {
	// The Dial seam fails outright (http dialing is lazy, so this branch
	// needs the injected dialer).
	_, err := New(context.Background(), Config{
		RPCURL: "http://ignored", PrivateKeyHex: testKeyHex, GameAddr: gameAddr,
		Out: &syncBuffer{},
		Dial: func(ctx context.Context, url string) (*ethrpc.Client, error) {
			return nil, errors.New("refused")
		},
	})
	if err == nil || !strings.Contains(err.Error(), "can't connect to ETH RPC") {
		t.Errorf("New with failing dial = %v", err)
	}
}

func TestNewEngineDefaultsOutToStdout(t *testing.T) {
	// Out nil defaults to os.Stdout; the engine must still construct.
	w := newGameWorld(t)
	e, err := New(context.Background(), Config{
		RPCURL: w.chain.URL(), PrivateKeyHex: testKeyHex, GameAddr: gameAddr,
	})
	if err != nil {
		t.Fatalf("New with nil Out: %v", err)
	}
	if e.cfg.Out == nil {
		t.Error("Out not defaulted")
	}
}

func TestFmtEthNil(t *testing.T) {
	if got := fmtEth(nil); got != "0" {
		t.Errorf("fmtEth(nil) = %q, want 0", got)
	}
}

func TestReconnectDirectBranches(t *testing.T) {
	t.Run("attempts exceeded", func(t *testing.T) {
		w := newGameWorld(t)
		e, _ := newTestEngine(t, w, nil)
		e.reconnectCount = maxReconnectAttempts
		err := e.reconnect(context.Background())
		if err == nil || !strings.Contains(err.Error(), "max reconnection attempts (10) exceeded") {
			t.Errorf("reconnect over cap = %v", err)
		}
	})

	t.Run("sleep cancelled", func(t *testing.T) {
		w := newGameWorld(t)
		e, _ := newTestEngine(t, w, func(cfg *Config) {
			cfg.Sleep = sleepContext // real ctx-aware sleep
		})
		cancelled, cancel := context.WithCancel(context.Background())
		cancel()
		if err := e.reconnect(cancelled); !errors.Is(err, context.Canceled) {
			t.Errorf("reconnect with cancelled ctx = %v", err)
		}
	})

	t.Run("chain id fetch fails after redial", func(t *testing.T) {
		w := newGameWorld(t)
		e, _ := newTestEngine(t, w, nil)
		w.chain.FailNextRPC("eth_chainId", "id gone")
		err := e.reconnect(context.Background())
		if err == nil || !strings.Contains(err.Error(), "can't get chain ID") {
			t.Errorf("reconnect with failing chain id = %v", err)
		}
	})
}

func TestFindRWalkTokenIDKeepsValidToken(t *testing.T) {
	w := newGameWorld(t)
	e, _ := newTestEngine(t, w, nil)
	// usedRandomWalkNfts answers 0 (unused): the held token stays.
	e.nextRWalkTokenID.Store(5)
	e.findRWalkTokenID()
	if got := e.nextRWalkTokenID.Load(); got != 5 {
		t.Errorf("valid token reset: %d, want 5", got)
	}
}

func TestPrintStatsBalanceFailure(t *testing.T) {
	w := newGameWorld(t)
	e, out := newTestEngine(t, w, nil)
	w.chain.FailNextRPC("eth_getBalance", "down")
	e.printStats()
	if !strings.Contains(out.String(), "Current Balance:   0.000000000000000000 ETH") {
		t.Errorf("stats with failed balance read:\n%s", out.String())
	}
}

func TestRunExitsGracefullyWhenCancelledDuringReconnect(t *testing.T) {
	w := newGameWorld(t)
	w.game.Handle("getNextCstBidPrice", func([]any) ([]any, error) {
		return nil, errors.New("node degraded")
	})

	ctx, cancel := context.WithCancel(context.Background())
	e, out := newTestEngine(t, w, func(cfg *Config) {
		cfg.Sleep = func(sctx context.Context, d time.Duration) error {
			if d == reconnectDelay {
				// Cancel mid-reconnect: Run must exit gracefully.
				cancel()
			}
			return sctx.Err()
		}
	})

	if err := e.Run(ctx); err != nil {
		t.Fatalf("Run cancelled during reconnect = %v", err)
	}
	if !strings.Contains(out.String(), "Shutting down gracefully...") {
		t.Errorf("output missing graceful shutdown:\n%s", out.String())
	}
}

func TestSleepContext(t *testing.T) {
	if err := sleepContext(context.Background(), 0); err != nil {
		t.Errorf("zero sleep = %v", err)
	}
	if err := sleepContext(context.Background(), time.Millisecond); err != nil {
		t.Errorf("short sleep = %v", err)
	}
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if err := sleepContext(cancelled, time.Hour); !errors.Is(err, context.Canceled) {
		t.Errorf("cancelled sleep = %v, want context.Canceled", err)
	}
}

func TestInitialBiddingDelayedReceipt(t *testing.T) {
	// The first initial bid's receipt arrives only on a later poll: the
	// receipt-wait loop retries, then counts the confirmed bid.
	w := newGameWorld(t)
	w.setLastBidder(otherAddrEng())
	var roundCalls, priceCalls atomic.Int64
	w.game.Handle("roundNum", func([]any) ([]any, error) {
		if roundCalls.Add(1) == 1 {
			return []any{big.NewInt(0)}, nil
		}
		return []any{big.NewInt(1)}, nil
	})
	w.game.Handle("getNextEthBidPrice", func([]any) ([]any, error) {
		if priceCalls.Add(1) <= 2 { // refresh + first initial bid
			return []any{eth(0.01)}, nil
		}
		return []any{eth(10)}, nil // above the limit: stop after one bid
	})
	w.chain.MarkNextTxPending()

	e, out := newTestEngine(t, w, func(cfg *Config) {
		cfg.InitialBidPrice = eth(1)
		var sleeps atomic.Int64
		cfg.Sleep = func(ctx context.Context, d time.Duration) error {
			// Release the held receipt after a few polls.
			if sleeps.Add(1) == 4 {
				w.chain.ReleasePendingTxs()
			}
			return ctx.Err()
		}
	})
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()
	go func() {
		for ctx.Err() == nil {
			if strings.Contains(out.String(), "Initial bidding finished") {
				cancel()
				return
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()
	if err := e.Run(ctx); err != nil {
		t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
	}
	text := out.String()
	if !strings.Contains(text, "Initial bid successful") {
		t.Errorf("output missing confirmed initial bid:\n%s", text)
	}
	if !strings.Contains(text, "Initial bidding finished: 1 bids") {
		t.Errorf("output missing one-bid finish:\n%s", text)
	}
}

func TestInitialBiddingErrorPaths(t *testing.T) {
	t.Run("price read failures then success", func(t *testing.T) {
		w := newGameWorld(t)
		w.setLastBidder(otherAddrEng())
		var roundCalls, priceCalls atomic.Int64
		w.game.Handle("roundNum", func([]any) ([]any, error) {
			if roundCalls.Add(1) == 1 {
				return []any{big.NewInt(0)}, nil
			}
			return []any{big.NewInt(1)}, nil
		})
		w.game.Handle("getNextEthBidPrice", func([]any) ([]any, error) {
			// Call 1 is the loop's market refresh; call 2 is the first read
			// inside the initial-bidding loop — that one glitches.
			if priceCalls.Add(1) == 2 {
				return nil, errors.New("price glitch")
			}
			return []any{eth(0.05)}, nil // above the 0.01 limit: stop
		})

		e, out := newTestEngine(t, w, func(cfg *Config) {
			cfg.InitialBidPrice = eth(0.01)
		})
		ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
		defer cancel()
		go func() {
			for ctx.Err() == nil {
				if strings.Contains(out.String(), "Initial bidding finished") {
					cancel()
					return
				}
				time.Sleep(5 * time.Millisecond)
			}
		}()
		if err := e.Run(ctx); err != nil {
			t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
		}
		text := out.String()
		if !strings.Contains(text, "Error getting bid price: ") {
			t.Errorf("output missing price glitch:\n%s", text)
		}
		if !strings.Contains(text, "Reached bid price limit, stopping initial bidding") {
			t.Errorf("output missing limit stop:\n%s", text)
		}
		if !strings.Contains(text, "Initial bidding finished: 0 bids") {
			t.Errorf("output missing zero-bid finish:\n%s", text)
		}
	})

	t.Run("insufficient balance stops", func(t *testing.T) {
		w := newGameWorld(t)
		w.setLastBidder(otherAddrEng())
		w.chain.SetBalance(botAddr, eth(0.001)) // below the first bid price
		var roundCalls atomic.Int64
		w.game.Handle("roundNum", func([]any) ([]any, error) {
			if roundCalls.Add(1) == 1 {
				return []any{big.NewInt(0)}, nil
			}
			return []any{big.NewInt(1)}, nil
		})

		e, out := newTestEngine(t, w, func(cfg *Config) {
			cfg.InitialBidPrice = eth(1)
		})
		ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
		defer cancel()
		go func() {
			for ctx.Err() == nil {
				if strings.Contains(out.String(), "Initial bidding finished") {
					cancel()
					return
				}
				time.Sleep(5 * time.Millisecond)
			}
		}()
		if err := e.Run(ctx); err != nil {
			t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
		}
		if !strings.Contains(out.String(), "Insufficient ETH balance for initial bid, stopping") {
			t.Errorf("output missing balance stop:\n%s", out.String())
		}
	})

	t.Run("safety cap stops runaway bidding", func(t *testing.T) {
		w := newGameWorld(t)
		w.setLastBidder(otherAddrEng())
		w.setEthBidPrice(eth(0.001))
		var roundCalls atomic.Int64
		w.game.Handle("roundNum", func([]any) ([]any, error) {
			if roundCalls.Add(1) == 1 {
				return []any{big.NewInt(0)}, nil
			}
			return []any{big.NewInt(1)}, nil
		})
		w.game.Handle("getNextEthBidPrice", func([]any) ([]any, error) {
			return []any{eth(0.001)}, nil // never reaches the limit
		})

		e, out := newTestEngine(t, w, func(cfg *Config) {
			cfg.InitialBidPrice = eth(100)
		})
		ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
		defer cancel()
		go func() {
			for ctx.Err() == nil {
				if strings.Contains(out.String(), "Initial bidding finished") {
					cancel()
					return
				}
				time.Sleep(5 * time.Millisecond)
			}
		}()
		if err := e.Run(ctx); err != nil {
			t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
		}
		if !strings.Contains(out.String(), "SAFETY: Reached max initial bids (20), stopping") {
			t.Errorf("output missing safety stop:\n%s", out.String())
		}
		if got := w.chain.SubmittedTxCount(); got != maxInitialBids {
			t.Errorf("submitted txs = %d, want %d", got, maxInitialBids)
		}
	})

	t.Run("mid-loop nonce and send failures retry", func(t *testing.T) {
		w := newGameWorld(t)
		w.setLastBidder(otherAddrEng())
		var roundCalls, priceCalls atomic.Int64
		w.game.Handle("roundNum", func([]any) ([]any, error) {
			if roundCalls.Add(1) == 1 {
				return []any{big.NewInt(0)}, nil
			}
			return []any{big.NewInt(1)}, nil
		})
		w.game.Handle("getNextEthBidPrice", func([]any) ([]any, error) {
			// refresh + three initial-loop reads (two failures + a success),
			// then above the limit to stop.
			if priceCalls.Add(1) <= 4 {
				return []any{eth(0.01)}, nil
			}
			return []any{eth(10)}, nil
		})
		// First initial bid: nonce fetch fails. Second: send rejected.
		w.chain.FailNextRPC("eth_getTransactionCount", "nonce service down")
		w.chain.RejectNextSendWith("mempool full")

		e, out := newTestEngine(t, w, func(cfg *Config) {
			cfg.InitialBidPrice = eth(0.02)
		})
		ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
		defer cancel()
		go func() {
			for ctx.Err() == nil {
				if strings.Contains(out.String(), "Initial bidding finished") {
					cancel()
					return
				}
				time.Sleep(5 * time.Millisecond)
			}
		}()
		if err := e.Run(ctx); err != nil {
			t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
		}
		text := out.String()
		if !strings.Contains(text, "Bid error: ") {
			t.Errorf("output missing send rejection:\n%s", text)
		}
		if !strings.Contains(text, "Initial bidding finished: 1 bids") {
			t.Errorf("output missing eventual success:\n%s", text)
		}
	})

	t.Run("cancel during receipt wait returns", func(t *testing.T) {
		w := newGameWorld(t)
		w.setLastBidder(otherAddrEng())
		w.setEthBidPrice(eth(0.01))
		var roundCalls atomic.Int64
		w.game.Handle("roundNum", func([]any) ([]any, error) {
			if roundCalls.Add(1) == 1 {
				return []any{big.NewInt(0)}, nil
			}
			return []any{big.NewInt(1)}, nil
		})
		w.game.Handle("getNextEthBidPrice", func([]any) ([]any, error) {
			return []any{eth(0.01)}, nil
		})
		w.chain.MarkNextTxPending() // the receipt never arrives

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var sleeps atomic.Int64
		e, _ := newTestEngine(t, w, func(cfg *Config) {
			cfg.InitialBidPrice = eth(1)
			cfg.Sleep = func(sctx context.Context, d time.Duration) error {
				if sleeps.Add(1) > 5 {
					cancel() // cancel while polling for the receipt
				}
				return sctx.Err()
			}
		})
		done := make(chan error, 1)
		go func() { done <- e.Run(ctx) }()
		select {
		case err := <-done:
			if err != nil {
				t.Fatalf("Run = %v, want graceful nil", err)
			}
		case <-time.After(20 * time.Second):
			t.Fatal("Run hung during cancelled receipt wait")
		}
	})

	t.Run("failed bid tx stops", func(t *testing.T) {
		w := newGameWorld(t)
		w.setLastBidder(otherAddrEng())
		w.setEthBidPrice(eth(0.01))
		var roundCalls atomic.Int64
		w.game.Handle("roundNum", func([]any) ([]any, error) {
			if roundCalls.Add(1) == 1 {
				return []any{big.NewInt(0)}, nil
			}
			return []any{big.NewInt(1)}, nil
		})
		w.game.Handle("getNextEthBidPrice", func([]any) ([]any, error) {
			return []any{eth(0.01)}, nil
		})

		e, out := newTestEngine(t, w, func(cfg *Config) {
			cfg.InitialBidPrice = eth(5)
		})
		w.chain.MarkNextTxReverted()
		ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
		defer cancel()
		go func() {
			for ctx.Err() == nil {
				if strings.Contains(out.String(), "Initial bidding finished") {
					cancel()
					return
				}
				time.Sleep(5 * time.Millisecond)
			}
		}()
		if err := e.Run(ctx); err != nil {
			t.Fatalf("Run: %v\noutput:\n%s", err, out.String())
		}
		if !strings.Contains(out.String(), "Initial bid tx failed, stopping") {
			t.Errorf("output missing failed-bid stop:\n%s", out.String())
		}
		if e.stats.failedTxCount != 1 {
			t.Errorf("failedTxCount = %d, want 1", e.stats.failedTxCount)
		}
	})
}
