// Unit tests (no Docker) for the startup contract-parameter sync's chain-read
// logic: contract-mechanics detection (V1 divisor model vs V2 duration model)
// and the versioned read fallbacks, exercised against real abigen bindings
// over the deterministic fake chain (internal/testchain + ContractStub).
package cosmicgame

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// syncTestGameAddr is the contract address the unit-test bindings target.
var syncTestGameAddr = ethcommon.HexToAddress("0x2000000000000000000000000000000000000002")

// newSyncBindings starts a fake chain serving stub at the game address and
// returns real V1/V2 abigen bindings over it. Methods left un-stubbed answer
// with an error, exactly like a revert on a real node.
func newSyncBindings(t *testing.T, stub *testchain.ContractStub) (*cgc.CosmicSignatureGame, *cgc.CosmicSignatureGameV2) {
	t.Helper()
	chain := testchain.New(t)
	chain.RegisterCall(syncTestGameAddr, stub.Handler())

	rpcClient, err := rpc.DialContext(context.Background(), chain.URL())
	if err != nil {
		t.Fatalf("dialing fake chain: %v", err)
	}
	client := ethclient.NewClient(rpcClient)
	t.Cleanup(client.Close)

	v1, err := cgc.NewCosmicSignatureGame(syncTestGameAddr, client)
	if err != nil {
		t.Fatalf("binding V1 game: %v", err)
	}
	v2, err := cgc.NewCosmicSignatureGameV2(syncTestGameAddr, client)
	if err != nil {
		t.Fatalf("binding V2 game: %v", err)
	}
	return v1, v2
}

func gameStub() *testchain.ContractStub {
	return testchain.MustContractStub(cgc.CosmicSignatureGameABI, cgc.CosmicSignatureGameV2ABI)
}

func TestProbeContractMechanicsUnknownWithoutBindings(t *testing.T) {
	got := probeContractMechanics(nil, nil, &bind.CallOpts{})
	if got != contractMechanicsUnknown {
		t.Fatalf("expected unknown mechanics, got %d", got)
	}
}

func TestProbeContractMechanicsV2ByDuration(t *testing.T) {
	stub := gameStub().Return("cstDutchAuctionDuration", big.NewInt(1800))
	v1, v2 := newSyncBindings(t, stub)
	if got := probeContractMechanics(v1, v2, &bind.CallOpts{}); got != contractMechanicsV2 {
		t.Fatalf("mechanics = %d, want V2 (%d)", got, contractMechanicsV2)
	}
}

func TestProbeContractMechanicsV1ByDivisor(t *testing.T) {
	stub := gameStub().Return("cstDutchAuctionDurationDivisor", big.NewInt(400))
	v1, v2 := newSyncBindings(t, stub)
	if got := probeContractMechanics(v1, v2, &bind.CallOpts{}); got != contractMechanicsV1 {
		t.Fatalf("mechanics = %d, want V1 (%d)", got, contractMechanicsV1)
	}
}

func TestProbeContractMechanicsV2ByMultiplierFallback(t *testing.T) {
	// Neither duration nor divisor answers, but the V2-only reward
	// multiplier does: the third probe identifies V2.
	stub := gameStub().Return("bidCstRewardAmountMultiplier", big.NewInt(10))
	v1, v2 := newSyncBindings(t, stub)
	if got := probeContractMechanics(v1, v2, &bind.CallOpts{}); got != contractMechanicsV2 {
		t.Fatalf("mechanics = %d, want V2 (%d)", got, contractMechanicsV2)
	}
}

func TestProbeContractMechanicsUnknownWhenNothingAnswers(t *testing.T) {
	v1, v2 := newSyncBindings(t, gameStub())
	if got := probeContractMechanics(v1, v2, &bind.CallOpts{}); got != contractMechanicsUnknown {
		t.Fatalf("mechanics = %d, want unknown (%d)", got, contractMechanicsUnknown)
	}
}

func TestReadCstRewardPrefersV2Multiplier(t *testing.T) {
	stub := gameStub().
		Return("bidCstRewardAmountMultiplier", big.NewInt(7)).
		Return("cstRewardAmountForBidding", big.NewInt(5))
	v1, v2 := newSyncBindings(t, stub)

	got, err := readCstReward(v1, v2, &bind.CallOpts{}, contractMechanicsV2)
	if err != nil {
		t.Fatalf("readCstReward(V2): %v", err)
	}
	if got != "7" {
		t.Errorf("readCstReward(V2) = %s, want 7 (the V2 multiplier)", got)
	}

	got, err = readCstReward(v1, v2, &bind.CallOpts{}, contractMechanicsV1)
	if err != nil {
		t.Fatalf("readCstReward(V1): %v", err)
	}
	if got != "5" {
		t.Errorf("readCstReward(V1) = %s, want 5 (the V1 reward)", got)
	}
}

func TestReadCstRewardFallsBackToV2WhenV1Reverts(t *testing.T) {
	stub := gameStub().Return("bidCstRewardAmountMultiplier", big.NewInt(9))
	v1, v2 := newSyncBindings(t, stub)

	got, err := readCstReward(v1, v2, &bind.CallOpts{}, contractMechanicsUnknown)
	if err != nil {
		t.Fatalf("readCstReward(fallback): %v", err)
	}
	if got != "9" {
		t.Errorf("readCstReward(fallback) = %s, want 9 (the V2 multiplier)", got)
	}

	// Nothing answers at all: the error propagates.
	v1none, v2none := newSyncBindings(t, gameStub())
	if _, err := readCstReward(v1none, v2none, &bind.CallOpts{}, contractMechanicsUnknown); err == nil {
		t.Error("readCstReward with no readable method returned nil error")
	}
}

func TestReadRoundStartCSTAuctionPerMechanics(t *testing.T) {
	stub := gameStub().
		Return("cstDutchAuctionDuration", big.NewInt(1800)).
		Return("cstDutchAuctionDurationDivisor", big.NewInt(400))
	v1, v2 := newSyncBindings(t, stub)

	got, err := readRoundStartCSTAuction(v1, v2, &bind.CallOpts{}, contractMechanicsV2)
	if err != nil || got != "1800" {
		t.Errorf("readRoundStartCSTAuction(V2) = %s, %v; want 1800 (duration)", got, err)
	}
	got, err = readRoundStartCSTAuction(v1, v2, &bind.CallOpts{}, contractMechanicsV1)
	if err != nil || got != "400" {
		t.Errorf("readRoundStartCSTAuction(V1) = %s, %v; want 400 (divisor)", got, err)
	}
}

func TestReadRoundStartCSTAuctionFallbackTail(t *testing.T) {
	// The divisor read reverts under V1 mechanics, but the final V2
	// duration retry answers — the mechanics probe can be stale across a
	// contract upgrade.
	stub := gameStub().Return("cstDutchAuctionDuration", big.NewInt(2500))
	v1, v2 := newSyncBindings(t, stub)
	got, err := readRoundStartCSTAuction(v1, v2, &bind.CallOpts{}, contractMechanicsV1)
	if err != nil || got != "2500" {
		t.Errorf("readRoundStartCSTAuction(stale V1) = %s, %v; want the V2 duration 2500", got, err)
	}

	// Nothing answers at all: the error propagates.
	v1none, v2none := newSyncBindings(t, gameStub())
	if _, err := readRoundStartCSTAuction(v1none, v2none, &bind.CallOpts{}, contractMechanicsV1); err == nil {
		t.Error("readRoundStartCSTAuction with no readable method returned nil error")
	}
}

func TestReadDelayDuration(t *testing.T) {
	stub := gameStub().Return("delayDurationBeforeRoundActivation", big.NewInt(1234))
	v1, v2 := newSyncBindings(t, stub)
	got, err := readDelayDuration(v1, v2, &bind.CallOpts{})
	if err != nil || got != 1234 {
		t.Errorf("readDelayDuration = %d, %v; want 1234", got, err)
	}

	v1none, v2none := newSyncBindings(t, gameStub())
	if _, err := readDelayDuration(v1none, v2none, &bind.CallOpts{}); err == nil {
		t.Error("readDelayDuration with no readable method returned nil error")
	}
}

func TestReadCSTAuctionDurationChangeDivisor(t *testing.T) {
	// V1 mechanics never carry the divisor, regardless of bindings.
	if got := readCSTAuctionDurationChangeDivisor(nil, &bind.CallOpts{}, contractMechanicsV1); got != -1 {
		t.Fatalf("expected -1 on V1, got %d", got)
	}

	stub := gameStub().Return("cstDutchAuctionDurationChangeDivisor", big.NewInt(25))
	_, v2 := newSyncBindings(t, stub)
	if got := readCSTAuctionDurationChangeDivisor(v2, &bind.CallOpts{}, contractMechanicsV2); got != 25 {
		t.Errorf("readCSTAuctionDurationChangeDivisor(V2) = %d, want 25", got)
	}

	_, v2none := newSyncBindings(t, gameStub())
	if got := readCSTAuctionDurationChangeDivisor(v2none, &bind.CallOpts{}, contractMechanicsV2); got != -1 {
		t.Errorf("readCSTAuctionDurationChangeDivisor(unreadable) = %d, want -1", got)
	}
}
