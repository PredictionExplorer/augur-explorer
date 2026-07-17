package main

// Error-branch tests: every per-call contract read failure and transaction
// failure of the ported commands, driven by partial stubs (un-stubbed
// methods revert) and the fake chain's failure injectors.

import (
	"errors"
	"math/big"
	"strings"
	"sync/atomic"
	"testing"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// failOne re-stubs one method to revert.
func failOne(stub *testchain.ContractStub, method string) *testchain.ContractStub {
	stub.Handle(method, func([]any) ([]any, error) {
		return nil, errors.New("injected revert")
	})
	return stub
}

func TestReadCommandsConnectFailure(t *testing.T) {
	// Every read command propagates a failed network connection. The chain
	// answers nothing (eth_chainId fails), so Connect errors.
	cases := []struct {
		name string
		run  func(t *testing.T) error
	}{
		{"info", func(t *testing.T) error {
			t.Helper()
			_, err := executeCmd(t, newInfoCmd(), testGameAddr.Hex())
			return err
		}},
		{"owner", func(t *testing.T) error {
			t.Helper()
			_, err := executeCmd(t, newOwnerCmd(), testGameAddr.Hex())
			return err
		}},
		{"donation-records", func(t *testing.T) error {
			t.Helper()
			_, err := executeCmd(t, newDonationRecordsCmd(), testGameAddr.Hex())
			return err
		}},
		{"erc20 balance", func(t *testing.T) error {
			t.Helper()
			_, err := executeCmd(t, newERC20BalanceCmd(), testTokenAddr.Hex(), otherAddr.Hex())
			return err
		}},
		{"erc20 allowance", func(t *testing.T) error {
			t.Helper()
			_, err := executeCmd(t, newERC20AllowanceCmd(), testTokenAddr.Hex(), otherAddr.Hex(), testSignerAddr.Hex())
			return err
		}},
		{"nft approved", func(t *testing.T) error {
			t.Helper()
			_, err := executeCmd(t, newNFTApprovedCmd(), testNFTAddr.Hex(), "7")
			return err
		}},
		{"nft is-approved-for-all", func(t *testing.T) error {
			t.Helper()
			_, err := executeCmd(t, newNFTIsApprovedForAllCmd(), testNFTAddr.Hex(), otherAddr.Hex(), testSignerAddr.Hex())
			return err
		}},
		{"nft owner-of", func(t *testing.T) error {
			t.Helper()
			_, err := executeCmd(t, newNFTOwnerOfCmd(), testNFTAddr.Hex(), "7")
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			chain := startReadChain(t)
			chain.FailNextRPC("eth_chainId", "node down")
			err := tc.run(t)
			if err == nil || !strings.Contains(err.Error(), "network connection failed") {
				t.Errorf("%s with dead node = %v", tc.name, err)
			}
		})
	}
}

func TestReadRoundStateFailures(t *testing.T) {
	// Every contract read of the claim-and-set planner's state snapshot has
	// its own error wrapper.
	cases := []struct {
		method string
		want   string
	}{
		{"roundNum", "roundNum:"},
		{"roundActivationTime", "roundActivationTime:"},
		{"getTotalNumBids", "getTotalNumBids:"},
		{"lastBidderAddress", "lastBidderAddress:"},
		{"getDurationUntilMainPrize", "getDurationUntilMainPrize:"},
		{"getDurationUntilMainPrizeRaw", "getDurationUntilMainPrizeRaw:"},
		{"timeoutDurationToClaimMainPrize", "timeoutDurationToClaimMainPrize:"},
		{"mainPrizeTimeIncrementInMicroSeconds", "mainPrizeTimeIncrementInMicroSeconds:"},
		{"delayDurationBeforeRoundActivation", "delayDurationBeforeRoundActivation:"},
		{"owner", "owner:"},
	}
	for _, tc := range cases {
		t.Run(tc.method, func(t *testing.T) {
			chain := startFundedChain(t)
			registerGame(t, chain, failOne(plannerStub(chain), tc.method))
			_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("planner with failing %s = %v, want %q", tc.method, err, tc.want)
			}
		})
	}
}

func TestPlannerTransactionFailures(t *testing.T) {
	t.Run("increment tx reverts", func(t *testing.T) {
		chain := startFundedChain(t)
		registerGame(t, chain, plannerStub(chain)) // path B: inactive direct
		chain.MarkNextTxReverted()
		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "setMainPrizeTimeIncrementInMicroSeconds") {
			t.Errorf("reverted increment = %v", err)
		}
	})

	t.Run("delay tx rejected", func(t *testing.T) {
		chain := startFundedChain(t)
		stub := plannerStub(chain)
		stub.Return("delayDurationBeforeRoundActivation", big.NewInt(0)) // delay differs
		registerGame(t, chain, stub)
		chain.RejectNextSendWith("underpriced")
		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "setDelayDurationBeforeRoundActivation") {
			t.Errorf("rejected delay = %v", err)
		}
	})

	t.Run("claim tx reverts", func(t *testing.T) {
		chain := startFundedChain(t)
		blockTime := testchain.BlockTimeInt64(100)
		stub := plannerStub(chain)
		stub.Return("getTotalNumBids", big.NewInt(5))
		stub.Return("lastBidderAddress", testSignerAddr)
		stub.Return("roundActivationTime", big.NewInt(blockTime-600)) // active
		registerGame(t, chain, stub)
		chain.MarkNextTxReverted() // the claim is the first transaction
		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "claimMainPrize") {
			t.Errorf("reverted claim = %v", err)
		}
	})

	t.Run("defer tx reverts", func(t *testing.T) {
		chain := startFundedChain(t)
		blockTime := testchain.BlockTimeInt64(100)
		stub := plannerStub(chain)
		stub.Return("getTotalNumBids", big.NewInt(5))
		stub.Return("lastBidderAddress", otherAddr)
		stub.Return("getDurationUntilMainPrize", big.NewInt(500))
		stub.Return("getDurationUntilMainPrizeRaw", big.NewInt(500))
		stub.Return("roundActivationTime", big.NewInt(blockTime-600)) // active
		registerGame(t, chain, stub)
		chain.MarkNextTxReverted()
		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "setRoundActivationTime failed on attempt 1") {
			t.Errorf("reverted defer = %v", err)
		}
	})

	t.Run("post-claim window still active", func(t *testing.T) {
		chain := startFundedChain(t)
		blockTime := testchain.BlockTimeInt64(100)
		stub := plannerStub(chain)
		stub.Return("getTotalNumBids", big.NewInt(5))
		stub.Return("lastBidderAddress", testSignerAddr)
		// Active before AND after the claim: ensureInactiveRound fails.
		stub.Return("roundActivationTime", big.NewInt(blockTime-600))
		registerGame(t, chain, stub)
		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "round is still active") {
			t.Errorf("active post-claim window = %v", err)
		}
	})

	t.Run("refresh fails after delay tx", func(t *testing.T) {
		chain := startFundedChain(t)
		stub := plannerStub(chain)
		stub.Return("delayDurationBeforeRoundActivation", big.NewInt(0))
		registerGame(t, chain, stub)
		// Call 1 of eth_gasPrice is the session construction; call 2 is the
		// post-delay Refresh — that one fails.
		chain.FailRPCAfter("eth_gasPrice", 1, "gas oracle down")
		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "network refresh failed after delay tx") {
			t.Errorf("refresh failure = %v", err)
		}
	})

	t.Run("verbose idempotent summary", func(t *testing.T) {
		chain := startFundedChain(t)
		stub := plannerStub(chain)
		stub.Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(targetIncrementSec*1_000_000))
		registerGame(t, chain, stub)
		out, err := runPlanner(t, true, testGameAddr.Hex(), "3600", "300")
		if err != nil {
			t.Fatalf("verbose idempotent: %v", err)
		}
		if !strings.Contains(out, "Nothing to do — increment and delay already match") {
			t.Errorf("output missing idempotent summary:\n%s", out)
		}
	})

	t.Run("path B refresh fails", func(t *testing.T) {
		chain := startFundedChain(t)
		registerGame(t, chain, plannerStub(chain))
		chain.FailRPCAfter("eth_gasPrice", 1, "gas oracle down")
		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "network refresh failed") {
			t.Errorf("path B refresh failure = %v", err)
		}
	})

	t.Run("path B round activates before increment", func(t *testing.T) {
		chain := startFundedChain(t)
		blockTime := testchain.BlockTimeInt64(100)
		stub := plannerStub(chain)
		var reads atomic.Int64
		stub.Handle("roundActivationTime", func([]any) ([]any, error) {
			// Inactive at the plan read, active at the ensureInactiveRound
			// re-check: the guard refuses to send the increment.
			if reads.Add(1) == 1 {
				return []any{big.NewInt(blockTime + 1000)}, nil
			}
			return []any{big.NewInt(blockTime - 10)}, nil
		})
		registerGame(t, chain, stub)
		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "cannot set increment") {
			t.Errorf("activation race = %v", err)
		}
	})

	t.Run("path B ensureInactiveRound read fails", func(t *testing.T) {
		chain := startFundedChain(t)
		stub := plannerStub(chain)
		var reads atomic.Int64
		blockTime := testchain.BlockTimeInt64(100)
		stub.Handle("roundActivationTime", func([]any) ([]any, error) {
			if reads.Add(1) == 1 {
				return []any{big.NewInt(blockTime + 1000)}, nil
			}
			return nil, errors.New("activation read revert")
		})
		registerGame(t, chain, stub)
		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "activation read revert") {
			t.Errorf("ensureInactiveRound read failure = %v", err)
		}
	})

	t.Run("path B increment send rejected", func(t *testing.T) {
		chain := startFundedChain(t)
		registerGame(t, chain, plannerStub(chain))
		chain.RejectNextSendWith("underpriced increment")
		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "setMainPrizeTimeIncrementInMicroSeconds") {
			t.Errorf("rejected increment = %v", err)
		}
	})

	t.Run("path C refresh failures", func(t *testing.T) {
		newActiveClaimable := func(t *testing.T) *testchain.Chain {
			t.Helper()
			chain := startFundedChain(t)
			blockTime := testchain.BlockTimeInt64(100)
			stub := plannerStub(chain)
			stub.Return("getTotalNumBids", big.NewInt(5))
			stub.Return("lastBidderAddress", testSignerAddr)
			stub.Handle("roundActivationTime", func([]any) ([]any, error) {
				if chain.SubmittedTxCount() >= 1 {
					return []any{big.NewInt(blockTime + 100000)}, nil
				}
				return []any{big.NewInt(blockTime - 600)}, nil
			})
			registerGame(t, chain, stub)
			return chain
		}

		chain := newActiveClaimable(t)
		chain.FailRPCAfter("eth_gasPrice", 1, "down before claim")
		if _, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300"); err == nil ||
			!strings.Contains(err.Error(), "network refresh failed before claim") {
			t.Errorf("refresh before claim = %v", err)
		}

		chain = newActiveClaimable(t)
		chain.FailRPCAfter("eth_gasPrice", 2, "down after claim")
		if _, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300"); err == nil ||
			!strings.Contains(err.Error(), "network refresh failed after claim") {
			t.Errorf("refresh after claim = %v", err)
		}

		// The post-claim increment send is the second transaction.
		chain = newActiveClaimable(t)
		chain.FailRPCAfter("eth_sendRawTransaction", 1, "increment rejected")
		if _, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300"); err == nil ||
			!strings.Contains(err.Error(), "setMainPrizeTimeIncrementInMicroSeconds") {
			t.Errorf("post-claim increment rejection = %v", err)
		}
	})

	t.Run("path D window read failures", func(t *testing.T) {
		newActiveNotClaimable := func(t *testing.T) (*testchain.Chain, *testchain.ContractStub) {
			t.Helper()
			chain := startFundedChain(t)
			blockTime := testchain.BlockTimeInt64(100)
			stub := plannerStub(chain)
			stub.Return("getTotalNumBids", big.NewInt(5))
			stub.Return("lastBidderAddress", otherAddr)
			stub.Return("getDurationUntilMainPrize", big.NewInt(500))
			stub.Return("getDurationUntilMainPrizeRaw", big.NewInt(500))
			stub.Return("roundActivationTime", big.NewInt(blockTime-600))
			registerGame(t, chain, stub)
			return chain, stub
		}

		// Refresh inside openInactiveWindow fails (call 2 of eth_gasPrice).
		chain, _ := newActiveNotClaimable(t)
		chain.FailRPCAfter("eth_gasPrice", 1, "window refresh down")
		if _, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300"); err == nil ||
			!strings.Contains(err.Error(), "window refresh down") {
			t.Errorf("window refresh failure = %v", err)
		}

		// readRoundState inside openInactiveWindow fails (roundNum call 2).
		_, stub := newActiveNotClaimable(t)
		var roundReads atomic.Int64
		stub.Handle("roundNum", func([]any) ([]any, error) {
			if roundReads.Add(1) >= 2 {
				return nil, errors.New("round read revert")
			}
			return []any{big.NewInt(3)}, nil
		})
		if _, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300"); err == nil ||
			!strings.Contains(err.Error(), "round read revert") {
			t.Errorf("window state read failure = %v", err)
		}

		// The deferred increment send is the second transaction.
		chain2, stub2 := newActiveNotClaimable(t)
		blockTime := testchain.BlockTimeInt64(100)
		stub2.Handle("roundActivationTime", func([]any) ([]any, error) {
			if chain2.SubmittedTxCount() >= 1 {
				return []any{big.NewInt(blockTime + 100000)}, nil
			}
			return []any{big.NewInt(blockTime - 600)}, nil
		})
		chain2.FailRPCAfter("eth_sendRawTransaction", 1, "deferred increment rejected")
		if _, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300"); err == nil ||
			!strings.Contains(err.Error(), "setMainPrizeTimeIncrementInMicroSeconds") {
			t.Errorf("deferred increment rejection = %v", err)
		}
	})

	t.Run("window opens on the final post-loop check", func(t *testing.T) {
		// Every defer attempt reads active, but the post-loop re-read sees
		// the window open: openInactiveWindow succeeds after 3 transactions.
		chain := startFundedChain(t)
		blockTime := testchain.BlockTimeInt64(100)
		stub := plannerStub(chain)
		stub.Return("getTotalNumBids", big.NewInt(5))
		stub.Return("lastBidderAddress", otherAddr)
		stub.Return("getDurationUntilMainPrize", big.NewInt(500))
		stub.Return("getDurationUntilMainPrizeRaw", big.NewInt(500))
		stub.Handle("roundActivationTime", func([]any) ([]any, error) {
			if chain.SubmittedTxCount() >= 3 {
				return []any{big.NewInt(blockTime + 100000)}, nil
			}
			return []any{big.NewInt(blockTime - 600)}, nil
		})
		registerGame(t, chain, stub)

		out, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err != nil {
			t.Fatalf("late window: %v\noutput: %s", err, out)
		}
		if got := chain.SubmittedTxCount(); got != 4 {
			t.Errorf("submitted txs = %d, want 4 (3 defers + increment)", got)
		}
	})

	t.Run("hardhat re-read fails after advance", func(t *testing.T) {
		chain := testchain.New(t)
		chain.SetChainID(31337)
		chain.EnsureBlock(100)
		chain.SetBalance(testSignerAddr, eth(100))
		t.Setenv("RPC_URL", chain.URL())
		t.Setenv("PKEY_HEX", testSignerKey)
		t.Setenv("GAS_PRICE_MULTIPLIER", "")

		blockTime := testchain.BlockTimeInt64(100)
		stub := plannerStub(chain)
		stub.Return("getTotalNumBids", big.NewInt(5))
		stub.Return("lastBidderAddress", otherAddr)
		stub.Return("getDurationUntilMainPrizeRaw", big.NewInt(-4000))
		stub.Handle("getDurationUntilMainPrize", func([]any) ([]any, error) {
			if chain.TimeOffset() > 0 {
				return nil, errors.New("post-advance read revert")
			}
			return []any{big.NewInt(5)}, nil
		})
		stub.Return("roundActivationTime", big.NewInt(blockTime-600))
		registerGame(t, chain, stub)

		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "post-advance read revert") {
			t.Errorf("post-advance read failure = %v", err)
		}
	})

	t.Run("hardhat advance fails", func(t *testing.T) {
		chain := testchain.New(t)
		chain.SetChainID(31337)
		chain.EnsureBlock(100)
		chain.SetBalance(testSignerAddr, eth(100))
		t.Setenv("RPC_URL", chain.URL())
		t.Setenv("PKEY_HEX", testSignerKey)
		t.Setenv("GAS_PRICE_MULTIPLIER", "")

		blockTime := testchain.BlockTimeInt64(100)
		stub := plannerStub(chain)
		stub.Return("getTotalNumBids", big.NewInt(5))
		stub.Return("lastBidderAddress", otherAddr)
		stub.Return("getDurationUntilMainPrizeRaw", big.NewInt(-4000))
		stub.Return("getDurationUntilMainPrize", big.NewInt(5))
		stub.Return("roundActivationTime", big.NewInt(blockTime-600))
		registerGame(t, chain, stub)
		chain.FailNextRPC("evm_increaseTime", "not hardhat after all")

		_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
		if err == nil || !strings.Contains(err.Error(), "advance hardhat time") {
			t.Errorf("failed time advance = %v", err)
		}
	})
}

func TestClaimPrizeDelayErrorPaths(t *testing.T) {
	t.Run("missing env", func(t *testing.T) {
		t.Setenv("RPC_URL", "")
		t.Setenv("PKEY_HEX", "")
		_, err := executeCmd(t, newClaimPrizeCmd(), "--delay=60", testGameAddr.Hex())
		if err == nil || !strings.Contains(err.Error(), "RPC_URL") {
			t.Errorf("missing env = %v", err)
		}
	})

	t.Run("refresh fails between steps", func(t *testing.T) {
		chain := startFundedChain(t)
		registerGame(t, chain, gameStubFull())
		chain.FailRPCAfter("eth_gasPrice", 1, "gas oracle down")
		_, err := executeCmd(t, newClaimPrizeCmd(), "--delay=60", testGameAddr.Hex())
		if err == nil || !strings.Contains(err.Error(), "network refresh failed") {
			t.Errorf("refresh failure = %v", err)
		}
	})

	t.Run("step 2 claim rejected", func(t *testing.T) {
		chain := startFundedChain(t)
		registerGame(t, chain, gameStubFull())
		chain.FailRPCAfter("eth_sendRawTransaction", 1, "claim rejected")
		out, err := executeCmd(t, newClaimPrizeCmd(), "--delay=60", testGameAddr.Hex())
		if err == nil || !strings.Contains(err.Error(), "transaction did not succeed") {
			t.Fatalf("step 2 rejection = %v", err)
		}
		if !strings.Contains(out, "claim rejected") {
			t.Errorf("output missing rejection message:\n%s", out)
		}
	})
}

func TestInfoCommandZeroPriceAnnotations(t *testing.T) {
	chain := startReadChain(t)
	stub := infoGameStub()
	stub.Return("ethDutchAuctionBeginningBidPrice", big.NewInt(0))
	stub.Return("cstDutchAuctionBeginningBidPrice", big.NewInt(0))
	// The CST auction's second value is a start timestamp (Unix seconds),
	// exercising the elapsed-from-start rendering.
	blockTime := testchain.BlockTimeInt64(100)
	stub.Return("getCstDutchAuctionDurations", big.NewInt(1800), big.NewInt(blockTime-120))
	registerInfoWorld(t, chain, stub)

	out, err := executeCmd(t, newInfoCmd(), testGameAddr.Hex())
	if err != nil {
		t.Fatalf("info: %v\noutput: %s", err, out)
	}
	for _, want := range []string{
		"when 0, contract uses a minimum floor",
		"stored begin price is only set after first bid",
		"elapsed from start_ts; raw 2nd value was start timestamp",
		"when 0, effective CST price = getNextCstBidPrice()",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
	// The elapsed seconds derive from block time minus the start timestamp.
	if !strings.Contains(out, "= 120 / 1800 (elapsed from start_ts") {
		t.Errorf("output missing derived elapsed seconds:\n%s", out)
	}
}

func TestInfoCommandNegativeStartTimestampClamps(t *testing.T) {
	chain := startReadChain(t)
	stub := infoGameStub()
	blockTime := testchain.BlockTimeInt64(100)
	// A start timestamp in the future clamps the derived elapsed to zero.
	stub.Return("getCstDutchAuctionDurations", big.NewInt(1800), big.NewInt(blockTime+500))
	registerInfoWorld(t, chain, stub)

	out, err := executeCmd(t, newInfoCmd(), testGameAddr.Hex())
	if err != nil {
		t.Fatalf("info: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "= 0 / 1800 (elapsed from start_ts") {
		t.Errorf("output missing clamped elapsed:\n%s", out)
	}
}

func TestInfoCommandPerReadFailures(t *testing.T) {
	// Each info section wraps its contract reads with the method name; the
	// stub answers everything except the method under test.
	methods := []struct {
		method string
		want   string
	}{
		{"roundNum", "RoundNum()"},
		{"delayDurationBeforeRoundActivation", "DelayDurationBeforeRoundActivation()"},
		{"getTotalNumBids", "GetTotalNumBids()"},
		{"bidderAddresses", "BidderAddresses()"},
		{"getDurationUntilMainPrize", "GetDurationUntilMainPrize()"},
		{"getDurationUntilMainPrizeRaw", "GetDurationUntilMainPrizeRaw()"},
		{"mainPrizeTime", "MainPrizeTime()"},
		{"timeoutDurationToClaimMainPrize", "TimeoutDurationToClaimMainPrize()"},
		{"mainPrizeTimeIncrementInMicroSeconds", "MainPrizeTimeIncrementInMicroSeconds()"},
		{"getMainPrizeTimeIncrement", "GetMainPrizeTimeIncrement()"},
		{"mainPrizeTimeIncrementIncreaseDivisor", "MainPrizeTimeIncrementIncreaseDivisor()"},
		{"initialDurationUntilMainPrizeDivisor", "InitialDurationUntilMainPrizeDivisor()"},
		{"getInitialDurationUntilMainPrize", "GetInitialDurationUntilMainPrize()"},
		{"nextEthBidPrice", "NextEthBidPrice()"},
		{"getNextEthBidPrice", "GetNextEthBidPrice()"},
		{"getNextCstBidPrice", "GetNextCstBidPrice()"},
		{"ethBidPriceIncreaseDivisor", "EthBidPriceIncreaseDivisor()"},
		{"cstRewardAmountForBidding", "CstRewardAmountForBidding()"},
		{"getCstDutchAuctionDurations", "GetCstDutchAuctionDurations()"},
		{"getEthDutchAuctionDurations", "GetEthDutchAuctionDurations()"},
		{"cstDutchAuctionBeginningBidPrice", "CstDutchAuctionBeginningBidPrice()"},
		{"ethDutchAuctionBeginningBidPrice", "EthDutchAuctionBeginningBidPrice()"},
		{"ethDutchAuctionEndingBidPriceDivisor", "EthDutchAuctionEndingBidPriceDivisor()"},
		{"lastBidderAddress", "LastBidderAddress()"},
		{"lastCstBidderAddress", "LastCstBidderAddress()"},
		{"tryGetCurrentChampions", "TryGetCurrentChampions()"},
		{"enduranceChampionStartTimeStamp", "EnduranceChampionStartTimeStamp()"},
		{"prevEnduranceChampionDuration", "PrevEnduranceChampionDuration()"},
		{"getMainEthPrizeAmount", "GetMainEthPrizeAmount()"},
		{"mainEthPrizeAmountPercentage", "MainEthPrizeAmountPercentage()"},
		{"getCharityEthDonationAmount", "GetCharityEthDonationAmount()"},
		{"charityEthDonationAmountPercentage", "CharityEthDonationAmountPercentage()"},
		{"getRaffleTotalEthPrizeAmountForBidders", "GetRaffleTotalEthPrizeAmountForBidders()"},
		{"raffleTotalEthPrizeAmountForBiddersPercentage", "RaffleTotalEthPrizeAmountForBiddersPercentage()"},
		{"getChronoWarriorEthPrizeAmount", "GetChronoWarriorEthPrizeAmount()"},
		{"chronoWarriorEthPrizeAmountPercentage", "ChronoWarriorEthPrizeAmountPercentage()"},
		{"getCosmicSignatureNftStakingTotalEthRewardAmount", "GetCosmicSignatureNftStakingTotalEthRewardAmount()"},
		{"cosmicSignatureNftStakingTotalEthRewardAmountPercentage", "CosmicSignatureNftStakingTotalEthRewardAmountPercentage()"},
		{"cstPrizeAmount", "CstPrizeAmount()"},
		{"numRaffleEthPrizesForBidders", "NumRaffleEthPrizesForBidders()"},
		{"numRaffleCosmicSignatureNftsForBidders", "NumRaffleCosmicSignatureNftsForBidders()"},
		{"numRaffleCosmicSignatureNftsForRandomWalkNftStakers", "NumRaffleCosmicSignatureNftsForRandomWalkNftStakers()"},
		{"prizesWallet", "PrizesWallet()"},
		{"charityAddress", "CharityAddress()"},
		{"nft", "Nft()"},
		{"token", "Token()"},
		{"randomWalkNft", "RandomWalkNft()"},
		{"stakingWalletCosmicSignatureNft", "StakingWalletCosmicSignatureNft()"},
		{"stakingWalletRandomWalkNft", "StakingWalletRandomWalkNft()"},
		{"marketingWallet", "MarketingWallet()"},
		{"marketingWalletCstContributionAmount", "MarketingWalletCstContributionAmount()"},
		{"bidMessageLengthMaxLimit", "BidMessageLengthMaxLimit()"},
	}
	for _, tc := range methods {
		t.Run(tc.method, func(t *testing.T) {
			chain := startReadChain(t)
			registerInfoWorld(t, chain, failOne(infoGameStub(), tc.method))
			_, err := executeCmd(t, newInfoCmd(), testGameAddr.Hex())
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("info with failing %s = %v, want %q", tc.method, err, tc.want)
			}
		})
	}

	t.Run("owner read fails", func(t *testing.T) {
		chain := startReadChain(t)
		stub := infoGameStub()
		stub.Handle("owner", func([]any) ([]any, error) { return nil, errors.New("x") })
		registerInfoWorld(t, chain, stub)
		_, err := executeCmd(t, newInfoCmd(), testGameAddr.Hex())
		if err == nil || !strings.Contains(err.Error(), "Owner()") {
			t.Errorf("info with failing owner = %v", err)
		}
	})

	t.Run("prizes wallet reads fail", func(t *testing.T) {
		chain := startReadChain(t)
		chain.RegisterCall(testGameAddr, infoGameStub().Handler())
		// PrizesWallet stub without methods: NextDonatedNftIndex reverts.
		chain.RegisterCall(testPrizesAddr, testchain.MustContractStub(cgc.PrizesWalletABI).Handler())
		charity := testchain.MustContractStub(cgc.CharityWalletABI)
		charity.Return("charityAddress", otherAddr)
		chain.RegisterCall(testCharityAddr, charity.Handler())

		_, err := executeCmd(t, newInfoCmd(), testGameAddr.Hex())
		if err == nil || !strings.Contains(err.Error(), "NextDonatedNftIndex()") {
			t.Errorf("info with failing prizes wallet = %v", err)
		}
	})

	t.Run("charity wallet read fails", func(t *testing.T) {
		chain := startReadChain(t)
		chain.RegisterCall(testGameAddr, infoGameStub().Handler())
		prizes := testchain.MustContractStub(cgc.PrizesWalletABI)
		prizes.Return("nextDonatedNftIndex", big.NewInt(6))
		prizes.Return("timeoutDurationToWithdrawPrizes", big.NewInt(600))
		chain.RegisterCall(testPrizesAddr, prizes.Handler())
		chain.RegisterCall(testCharityAddr, testchain.MustContractStub(cgc.CharityWalletABI).Handler())

		_, err := executeCmd(t, newInfoCmd(), testGameAddr.Hex())
		if err == nil || !strings.Contains(err.Error(), "calling CharityAddress()") {
			t.Errorf("info with failing charity wallet = %v", err)
		}
	})

	t.Run("game balance read fails", func(t *testing.T) {
		chain := startReadChain(t)
		registerInfoWorld(t, chain, infoGameStub())
		chain.FailNextRPC("eth_getBalance", "balance service down")
		_, err := executeCmd(t, newInfoCmd(), testGameAddr.Hex())
		if err == nil || !strings.Contains(err.Error(), "BalanceAt()") {
			t.Errorf("info with failing balance = %v", err)
		}
	})
}

func TestCommandReadFailures(t *testing.T) {
	t.Run("bid preflight reads", func(t *testing.T) {
		for _, tc := range []struct {
			method string
			want   string
		}{
			{"roundNum", "getting round number"},
			{"getNextEthBidPrice", "getting bid price"},
			{"lastBidderAddress", "getting last bidder"},
			{"getTotalNumBids", "getting total bids"},
		} {
			chain := startFundedChain(t)
			registerGame(t, chain, failOne(gameStubFull(), tc.method))
			_, err := executeCmd(t, newBidCmd(), testGameAddr.Hex())
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("bid with failing %s = %v, want %q", tc.method, err, tc.want)
			}
		}
	})

	t.Run("claim-prize preflight reads", func(t *testing.T) {
		for _, tc := range []struct {
			method string
			want   string
		}{
			{"roundNum", "getting round number"},
			{"lastBidderAddress", "getting last bidder"},
			{"getMainEthPrizeAmount", "getting prize amount"},
			{"getDurationUntilMainPrize", "getting duration until prize"},
		} {
			chain := startFundedChain(t)
			registerGame(t, chain, failOne(gameStubFull(), tc.method))
			_, err := executeCmd(t, newClaimPrizeCmd(), testGameAddr.Hex())
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("claim with failing %s = %v, want %q", tc.method, err, tc.want)
			}
		}
	})

	t.Run("claim-prize delay preflight reads", func(t *testing.T) {
		for _, tc := range []struct {
			method string
			want   string
		}{
			{"delayDurationBeforeRoundActivation", "getting current delay"},
			{"getMainEthPrizeAmount", "getting prize amount"},
		} {
			chain := startFundedChain(t)
			registerGame(t, chain, failOne(gameStubFull(), tc.method))
			_, err := executeCmd(t, newClaimPrizeCmd(), "--delay=60", testGameAddr.Hex())
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("claim --delay with failing %s = %v, want %q", tc.method, err, tc.want)
			}
		}
	})

	t.Run("donate contract balance read", func(t *testing.T) {
		chain := startFundedChain(t)
		registerGame(t, chain, gameStubFull())
		chain.FailNextRPC("eth_getBalance", "down")
		// The session's own balance read happens at construction, before the
		// contract balance: fail the third eth_getBalance call instead is
		// impossible with the one-shot injector, so the session read fails —
		// still an error path worth pinning.
		_, err := executeCmd(t, newDonateCmd(), testGameAddr.Hex(), "100")
		if err == nil {
			t.Error("donate with failing balance read succeeded")
		}
	})

	t.Run("setter read failures", func(t *testing.T) {
		for _, tc := range []struct {
			build  func() interface{ Execute() error }
			method string
		}{} {
			_ = tc
		}
		for _, tc := range []struct {
			name   string
			method string
			want   string
			args   []string
		}{
			{"set-activation-delay current", "delayDurationBeforeRoundActivation", "getting current delay", []string{testGameAddr.Hex(), "60"}},
			{"set-activation-delay owner", "owner", "getting contract owner", []string{testGameAddr.Hex(), "60"}},
		} {
			chain := startFundedChain(t)
			registerGame(t, chain, failOne(gameStubFull(), tc.method))
			_, err := executeCmd(t, newSetActivationDelayCmd(), tc.args...)
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("%s = %v, want %q", tc.name, err, tc.want)
			}
		}

		chain := startFundedChain(t)
		registerGame(t, chain, failOne(gameStubFull(), "roundActivationTime"))
		if _, err := executeCmd(t, newSetRoundActivationCmd(), testGameAddr.Hex(), "1900000000"); err == nil ||
			!strings.Contains(err.Error(), "getting current activation time") {
			t.Errorf("set-round-activation read failure = %v", err)
		}

		chain = startFundedChain(t)
		registerGame(t, chain, failOne(gameStubFull(), "mainPrizeTimeIncrementInMicroSeconds"))
		if _, err := executeCmd(t, newSetTimeIncrementCmd(), testGameAddr.Hex(), "3600"); err == nil ||
			!strings.Contains(err.Error(), "reading mainPrizeTimeIncrementInMicroSeconds") {
			t.Errorf("set-time-increment read failure = %v", err)
		}

		chain = startFundedChain(t)
		registerGame(t, chain, failOne(gameStubFull(), "initialDurationUntilMainPrizeDivisor"))
		if _, err := executeCmd(t, newSetInitialDurationDivisorCmd(), testGameAddr.Hex(), "100"); err == nil ||
			!strings.Contains(err.Error(), "getting current divisor") {
			t.Errorf("set-initial-duration-divisor read failure = %v", err)
		}
	})

	t.Run("erc20 required reads", func(t *testing.T) {
		newStub := func() *testchain.ContractStub {
			stub := testchain.MustContractStub(cgc.ERC20ABI)
			stub.Return("symbol", "CST")
			stub.Return("decimals", uint8(18))
			stub.Return("name", "Cosmic")
			stub.Return("totalSupply", eth(100))
			stub.Return("balanceOf", eth(5))
			stub.Return("allowance", big.NewInt(0))
			return stub
		}
		chain := startReadChain(t)
		chain.RegisterCall(testTokenAddr, failOne(newStub(), "totalSupply").Handler())
		if _, err := executeCmd(t, newERC20BalanceCmd(), testTokenAddr.Hex(), otherAddr.Hex()); err == nil ||
			!strings.Contains(err.Error(), "getting total supply") {
			t.Errorf("balance totalSupply failure = %v", err)
		}

		chain = startReadChain(t)
		chain.RegisterCall(testTokenAddr, failOne(newStub(), "balanceOf").Handler())
		if _, err := executeCmd(t, newERC20BalanceCmd(), testTokenAddr.Hex(), otherAddr.Hex()); err == nil ||
			!strings.Contains(err.Error(), "getting balance") {
			t.Errorf("balance balanceOf failure = %v", err)
		}
		chain = startReadChain(t)
		chain.RegisterCall(testTokenAddr, failOne(newStub(), "allowance").Handler())
		if _, err := executeCmd(t, newERC20AllowanceCmd(), testTokenAddr.Hex(), otherAddr.Hex(), testSignerAddr.Hex()); err == nil ||
			!strings.Contains(err.Error(), "getting allowance") {
			t.Errorf("allowance failure = %v", err)
		}

		chain = startFundedChain(t)
		chain.RegisterCall(testTokenAddr, failOne(newStub(), "allowance").Handler())
		if _, err := executeCmd(t, newERC20ApproveCmd(), testTokenAddr.Hex(), otherAddr.Hex()); err == nil ||
			!strings.Contains(err.Error(), "getting current allowance") {
			t.Errorf("approve allowance failure = %v", err)
		}
		chain = startFundedChain(t)
		chain.RegisterCall(testTokenAddr, failOne(newStub(), "balanceOf").Handler())
		if _, err := executeCmd(t, newERC20ApproveCmd(), testTokenAddr.Hex(), otherAddr.Hex()); err == nil ||
			!strings.Contains(err.Error(), "getting token balance") {
			t.Errorf("approve balance failure = %v", err)
		}
		chain = startFundedChain(t)
		chain.RegisterCall(testTokenAddr, failOne(newStub(), "allowance").Handler())
		if _, err := executeCmd(t, newERC20RevokeCmd(), testTokenAddr.Hex(), otherAddr.Hex()); err == nil ||
			!strings.Contains(err.Error(), "getting current allowance") {
			t.Errorf("revoke allowance failure = %v", err)
		}
	})

	t.Run("nft reads", func(t *testing.T) {
		newStub := func() *testchain.ContractStub {
			stub := testchain.MustContractStub(cgc.CosmicSignatureNftABI)
			stub.Return("getApproved", otherAddr)
			stub.Return("ownerOf", otherAddr)
			stub.Return("isApprovedForAll", true)
			stub.Return("balanceOf", big.NewInt(1))
			return stub
		}
		chain := startReadChain(t)
		chain.RegisterCall(testNFTAddr, failOne(newStub(), "getApproved").Handler())
		if _, err := executeCmd(t, newNFTApprovedCmd(), testNFTAddr.Hex(), "7"); err == nil ||
			!strings.Contains(err.Error(), "GetApproved()") {
			t.Errorf("getApproved failure = %v", err)
		}
		chain = startReadChain(t)
		chain.RegisterCall(testNFTAddr, failOne(newStub(), "ownerOf").Handler())
		if _, err := executeCmd(t, newNFTApprovedCmd(), testNFTAddr.Hex(), "7"); err == nil ||
			!strings.Contains(err.Error(), "OwnerOf()") {
			t.Errorf("ownerOf failure = %v", err)
		}
		chain = startReadChain(t)
		chain.RegisterCall(testNFTAddr, failOne(newStub(), "isApprovedForAll").Handler())
		if _, err := executeCmd(t, newNFTIsApprovedForAllCmd(), testNFTAddr.Hex(), otherAddr.Hex(), testSignerAddr.Hex()); err == nil ||
			!strings.Contains(err.Error(), "IsApprovedForAll()") {
			t.Errorf("isApprovedForAll failure = %v", err)
		}
	})

	t.Run("donation records reads", func(t *testing.T) {
		chain := startReadChain(t)
		chain.RegisterCall(testGameAddr, testchain.MustContractStub(cgc.CosmicSignatureGameABI).Handler())
		if _, err := executeCmd(t, newDonationRecordsCmd(), testGameAddr.Hex()); err == nil ||
			!strings.Contains(err.Error(), "getting NumEthDonationWithInfoRecords") {
			t.Errorf("record count failure = %v", err)
		}

		// A per-record read failure is reported inline and the dump goes on.
		chain = startReadChain(t)
		stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI)
		stub.Return("numEthDonationWithInfoRecords", big.NewInt(1))
		chain.RegisterCall(testGameAddr, stub.Handler())
		out, err := executeCmd(t, newDonationRecordsCmd(), testGameAddr.Hex())
		if err != nil {
			t.Fatalf("donation-records with failing record = %v", err)
		}
		if !strings.Contains(out, "Error at record 0") {
			t.Errorf("output missing per-record error:\n%s", out)
		}
	})

	t.Run("owner balance failure is soft", func(t *testing.T) {
		chain := startReadChain(t)
		stub := testchain.MustContractStub(cgc.OwnableABI)
		stub.Return("owner", otherAddr)
		chain.RegisterCall(testGameAddr, stub.Handler())
		chain.FailNextRPC("eth_getBalance", "down")
		out, err := executeCmd(t, newOwnerCmd(), testGameAddr.Hex())
		if err != nil {
			t.Fatalf("owner with failing balance = %v", err)
		}
		if strings.Contains(out, "Owner Balance") {
			t.Errorf("balance printed despite failed read:\n%s", out)
		}
	})
}
