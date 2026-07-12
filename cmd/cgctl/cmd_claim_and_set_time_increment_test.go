package main

// Planner tests for claim-and-set-time-increment: all four documented paths
// (idempotent exit, inactive-direct, claim-then-set, defer-then-set) run
// end-to-end against scripted contract state, with contract reactions to
// mined transactions keyed on the chain's submitted-transaction count.

import (
	"math/big"
	"strings"
	"testing"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

const (
	targetIncrementSec = 3600
	targetDelaySec     = 300
)

// plannerStub returns a game stub with the planner's baseline reads:
// round 3, target increment NOT set, target delay set, no bids, inactive.
func plannerStub(chain *testchain.Chain) *testchain.ContractStub {
	blockTime := int64(testchain.BlockTime(100))
	stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI)
	stub.Return("roundNum", big.NewInt(3))
	stub.Return("getTotalNumBids", big.NewInt(0))
	stub.Return("lastBidderAddress", zeroAddress())
	stub.Return("getDurationUntilMainPrize", big.NewInt(0))
	stub.Return("getDurationUntilMainPrizeRaw", big.NewInt(0))
	stub.Return("timeoutDurationToClaimMainPrize", big.NewInt(3600))
	stub.Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(1)) // differs from target
	stub.Return("delayDurationBeforeRoundActivation", big.NewInt(targetDelaySec))
	stub.Return("owner", testSignerAddr)
	stub.Return("roundActivationTime", big.NewInt(blockTime+1000)) // inactive
	_ = chain
	return stub
}

func zeroAddress() (a [20]byte) { return }

func runPlanner(t *testing.T, verbose bool, args ...string) (string, error) {
	t.Helper()
	full := args
	if verbose {
		full = append([]string{"-i"}, args...)
	}
	return executeCmd(t, newClaimAndSetTimeIncrementCmd(), full...)
}

func TestPlannerPathAIdempotent(t *testing.T) {
	chain := startFundedChain(t)
	stub := plannerStub(chain)
	stub.Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(targetIncrementSec*1_000_000))
	registerGame(t, chain, stub)
	// Any transaction attempt fails the test.
	chain.RejectNextSendWith("unexpected transaction in idempotent path")

	out, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
	if err != nil {
		t.Fatalf("planner A: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "Success. Already configured (increment=3600 sec, delay=300 sec, round=3).") {
		t.Errorf("output = %q, want idempotent success", out)
	}
	if got := chain.SubmittedTxCount(); got != 0 {
		t.Errorf("submitted txs = %d, want 0", got)
	}
}

func TestPlannerPathBInactiveDirect(t *testing.T) {
	chain := startFundedChain(t)
	registerGame(t, chain, plannerStub(chain))

	out, err := runPlanner(t, true, testGameAddr.Hex(), "3600", "300")
	if err != nil {
		t.Fatalf("planner B: %v\noutput: %s", err, out)
	}
	if got := chain.SubmittedTxCount(); got != 1 {
		t.Errorf("submitted txs = %d, want 1 (increment only)", got)
	}
	for _, want := range []string{
		"PLAN",
		"Round inactive              = true",
		"SET TIME INCREMENT (INACTIVE ROUND)",
		"Status                      = Time increment updated",
		"Round activation deferred; bidding resumes after the delay",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestPlannerPathCClaimThenSet(t *testing.T) {
	chain := startFundedChain(t)
	blockTime := int64(testchain.BlockTime(100))
	stub := plannerStub(chain)
	// Active round with the signer as claimable last bidder; delay differs
	// so the planner sends delay → claim → increment (3 txs).
	stub.Return("getTotalNumBids", big.NewInt(5))
	stub.Return("lastBidderAddress", testSignerAddr)
	stub.Handle("delayDurationBeforeRoundActivation", func([]any) ([]any, error) {
		if chain.SubmittedTxCount() >= 1 {
			return []any{big.NewInt(targetDelaySec)}, nil
		}
		return []any{big.NewInt(0)}, nil
	})
	stub.Handle("roundActivationTime", func([]any) ([]any, error) {
		// The round reads active until the claim (tx 2) opens the window.
		if chain.SubmittedTxCount() >= 2 {
			return []any{big.NewInt(blockTime + 100000)}, nil
		}
		return []any{big.NewInt(blockTime - 600)}, nil
	})
	registerGame(t, chain, stub)

	out, err := runPlanner(t, true, testGameAddr.Hex(), "3600", "300")
	if err != nil {
		t.Fatalf("planner C: %v\noutput: %s", err, out)
	}
	if got := chain.SubmittedTxCount(); got != 3 {
		t.Errorf("submitted txs = %d, want 3 (delay + claim + increment)", got)
	}
	for _, want := range []string{
		"Can claim now               = true (last bidder, prize timer expired)",
		"SET DELAY BEFORE NEXT ROUND",
		"CLAIM MAIN PRIZE",
		"SET TIME INCREMENT (POST-CLAIM INACTIVE WINDOW)",
		"Prize was claimed; next round activates after the delay",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestPlannerPathDDeferThenSet(t *testing.T) {
	chain := startFundedChain(t)
	blockTime := int64(testchain.BlockTime(100))
	stub := plannerStub(chain)
	// Active round, other bidder, prize not claimable.
	stub.Return("getTotalNumBids", big.NewInt(5))
	stub.Return("lastBidderAddress", otherAddr)
	stub.Return("getDurationUntilMainPrize", big.NewInt(500))
	stub.Return("getDurationUntilMainPrizeRaw", big.NewInt(500))
	stub.Handle("roundActivationTime", func([]any) ([]any, error) {
		// Active until the defer transaction moves activation to the future.
		if chain.SubmittedTxCount() >= 1 {
			return []any{big.NewInt(blockTime + 100000)}, nil
		}
		return []any{big.NewInt(blockTime - 600)}, nil
	})
	registerGame(t, chain, stub)

	out, err := runPlanner(t, true, testGameAddr.Hex(), "3600", "300")
	if err != nil {
		t.Fatalf("planner D: %v\noutput: %s", err, out)
	}
	if got := chain.SubmittedTxCount(); got != 2 {
		t.Errorf("submitted txs = %d, want 2 (defer + increment)", got)
	}
	for _, want := range []string{
		"Can claim now               = false (prize not claimable yet)",
		"DEFER ROUND ACTIVATION (NO CLAIMABLE PRIZE)",
		"SET TIME INCREMENT (DEFERRED INACTIVE WINDOW)",
		"Round activation deferred; bidding resumes after the delay",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestPlannerDeferExhaustionFails(t *testing.T) {
	chain := startFundedChain(t)
	blockTime := int64(testchain.BlockTime(100))
	stub := plannerStub(chain)
	stub.Return("getTotalNumBids", big.NewInt(5))
	stub.Return("lastBidderAddress", otherAddr)
	stub.Return("getDurationUntilMainPrize", big.NewInt(500))
	stub.Return("getDurationUntilMainPrizeRaw", big.NewInt(500))
	// Activation never moves to the future: every defer attempt fails to
	// open the window.
	stub.Return("roundActivationTime", big.NewInt(blockTime-600))
	registerGame(t, chain, stub)

	_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
	if err == nil || !strings.Contains(err.Error(), "round is still active after 3 defer attempts") {
		t.Fatalf("planner exhaustion = %v", err)
	}
	if got := chain.SubmittedTxCount(); got != 3 {
		t.Errorf("submitted txs = %d, want 3 defer attempts", got)
	}
}

func TestPlannerHardhatTimeAdvanceForClaim(t *testing.T) {
	chain := testchain.New(t)
	chain.SetChainID(31337) // Hardhat: the planner may advance time
	chain.EnsureBlock(100)
	chain.SetBalance(testSignerAddr, eth(100))
	chain.SetGasPrice(big.NewInt(1_000_000_000))
	t.Setenv("RPC_URL", chain.URL())
	t.Setenv("PKEY_HEX", testSignerKey)
	t.Setenv("GAS_PRICE_MULTIPLIER", "")

	blockTime := int64(testchain.BlockTime(100))
	stub := plannerStub(chain)
	stub.Return("getTotalNumBids", big.NewInt(5))
	stub.Return("lastBidderAddress", otherAddr)
	// Claimable through the anyone-timeout, but the clamped timer still
	// shows 5 seconds: on Hardhat the planner advances block time.
	stub.Return("getDurationUntilMainPrizeRaw", big.NewInt(-4000))
	stub.Handle("getDurationUntilMainPrize", func([]any) ([]any, error) {
		if chain.TimeOffset() > 0 {
			return []any{big.NewInt(0)}, nil
		}
		return []any{big.NewInt(5)}, nil
	})
	stub.Handle("roundActivationTime", func([]any) ([]any, error) {
		if chain.SubmittedTxCount() >= 1 { // after the claim
			return []any{big.NewInt(blockTime + 100000)}, nil
		}
		return []any{big.NewInt(blockTime - 600)}, nil
	})
	registerGame(t, chain, stub)

	out, err := runPlanner(t, true, testGameAddr.Hex(), "3600", "300")
	if err != nil {
		t.Fatalf("planner hardhat: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "ADVANCE HARDHAT TIME FOR CLAIM") {
		t.Errorf("output missing hardhat advance:\n%s", out)
	}
	if got := chain.TimeOffset(); got != 6 {
		t.Errorf("chain time offset = %d, want 6 (5s remaining + 1)", got)
	}
	if got := chain.SubmittedTxCount(); got != 2 {
		t.Errorf("submitted txs = %d, want 2 (claim + increment)", got)
	}
}

func TestPlannerRealNetworkRefusesEarlyClaim(t *testing.T) {
	chain := startFundedChain(t) // chain id 1337: not a dev chain
	blockTime := int64(testchain.BlockTime(100))
	stub := plannerStub(chain)
	stub.Return("getTotalNumBids", big.NewInt(5))
	stub.Return("lastBidderAddress", otherAddr)
	stub.Return("getDurationUntilMainPrizeRaw", big.NewInt(-4000))
	stub.Return("getDurationUntilMainPrize", big.NewInt(5))
	stub.Return("roundActivationTime", big.NewInt(blockTime-600))
	registerGame(t, chain, stub)

	_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
	if err == nil || !strings.Contains(err.Error(), "prize not claimable yet (5 seconds remaining)") {
		t.Fatalf("early claim on real network = %v", err)
	}
	if got := chain.SubmittedTxCount(); got != 0 {
		t.Errorf("submitted txs = %d, want 0", got)
	}
}

func TestPlannerRequiresOwnerForChanges(t *testing.T) {
	chain := startFundedChain(t)
	stub := plannerStub(chain)
	stub.Return("owner", otherAddr)
	registerGame(t, chain, stub)

	_, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "300")
	if err == nil || !strings.Contains(err.Error(), "PKEY_HEX must be the contract owner") {
		t.Fatalf("non-owner planner = %v", err)
	}
}

func TestPlannerDelayOnlyWhenIncrementAlreadySet(t *testing.T) {
	chain := startFundedChain(t)
	stub := plannerStub(chain)
	stub.Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(targetIncrementSec*1_000_000))
	stub.Handle("delayDurationBeforeRoundActivation", func([]any) ([]any, error) {
		if chain.SubmittedTxCount() >= 1 {
			return []any{big.NewInt(targetDelaySec)}, nil
		}
		return []any{big.NewInt(0)}, nil
	})
	registerGame(t, chain, stub)

	out, err := runPlanner(t, true, testGameAddr.Hex(), "3600", "300")
	if err != nil {
		t.Fatalf("planner delay-only: %v\noutput: %s", err, out)
	}
	if got := chain.SubmittedTxCount(); got != 1 {
		t.Errorf("submitted txs = %d, want 1 (delay only)", got)
	}
	if !strings.Contains(out, "Delay updated; increment already matches target") {
		t.Errorf("output missing delay-only summary:\n%s", out)
	}
}

func TestPlannerArgumentValidation(t *testing.T) {
	startFundedChain(t)
	if _, err := runPlanner(t, false, testGameAddr.Hex(), "0"); err == nil ||
		!strings.Contains(err.Error(), "time_increment_seconds must be a positive integer") {
		t.Errorf("zero increment = %v", err)
	}
	if _, err := runPlanner(t, false, testGameAddr.Hex(), "3600", "0"); err == nil ||
		!strings.Contains(err.Error(), "delay_seconds must be a positive integer") {
		t.Errorf("zero delay = %v", err)
	}
	if _, err := runPlanner(t, false, "nope", "3600"); err == nil ||
		!strings.Contains(err.Error(), "cosmicgame-addr") {
		t.Errorf("bad address = %v", err)
	}
}
