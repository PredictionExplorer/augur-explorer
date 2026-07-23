//go:build integration

package cosmicgame

import (
	"context"
	"log/slog"
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

func driftV1GameStub() *testchain.ContractStub {
	return testchain.MustContractStub(
		cgc.CosmicSignatureGameABI,
		cgc.CosmicSignatureGameV2ABI,
		cgc.CosmicSignatureGameV3ABI,
	).
		Return("cstDutchAuctionDurationDivisor", big.NewInt(400)).
		Return("cstRewardAmountForBidding", eth(100)).
		Return("timeoutDurationToClaimMainPrize", big.NewInt(86400)).
		Return("ethBidPriceIncreaseDivisor", big.NewInt(125)).
		Return("mainPrizeTimeIncrementIncreaseDivisor", big.NewInt(50)).
		Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(3_600_000_000)).
		Return("initialDurationUntilMainPrizeDivisor", big.NewInt(200)).
		Return("ethDutchAuctionDurationDivisor", big.NewInt(40)).
		Return("ethDutchAuctionEndingBidPriceDivisor", big.NewInt(10)).
		Return("cstDutchAuctionBeginningBidPriceMinLimit", eth(200)).
		Return("marketingWalletCstContributionAmount", eth(300)).
		Return("delayDurationBeforeRoundActivation", big.NewInt(1800))
}

func driftV3GameStub() *testchain.ContractStub {
	return testchain.MustContractStub(
		cgc.CosmicSignatureGameABI,
		cgc.CosmicSignatureGameV2ABI,
		cgc.CosmicSignatureGameV3ABI,
	).
		Return("mainPrizeNumCosmicSignatureNfts", big.NewInt(3)).
		Return("cstDutchAuctionDurationChangeDivisor", big.NewInt(33)).
		Return("bidCstRewardAmountMultiplier", eth(100)).
		Return("roundLateBidDurationDivisor", big.NewInt(4)).
		Return("roundLateBidPricePremiumAmountBaseMultiplier", big.NewInt(2)).
		Return("roundLateBidPricePremiumAmountExponent", big.NewInt(3)).
		Return("lastBidderBidCstRewardAmountPercentage", big.NewInt(90)).
		Return("timeoutDurationToClaimMainPrize", big.NewInt(86400)).
		Return("ethBidPriceIncreaseDivisor", big.NewInt(100)).
		Return("mainPrizeTimeIncrementIncreaseDivisor", big.NewInt(50)).
		Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(3_600_000_000)).
		Return("initialDurationUntilMainPrizeDivisor", big.NewInt(200)).
		Return("ethDutchAuctionDurationDivisor", big.NewInt(40)).
		Return("ethDutchAuctionEndingBidPriceDivisor", big.NewInt(10)).
		Return("cstDutchAuctionBeginningBidPriceMinLimit", eth(200)).
		Return("marketingWalletCstContributionAmount", eth(300)).
		Return("delayDurationBeforeRoundActivation", big.NewInt(1800))
}

func prizesWalletDriftStub() *testchain.ContractStub {
	return testchain.MustContractStub(cgc.PrizesWalletABI).
		Return("timeoutDurationToWithdrawPrizes", big.NewInt(604800))
}

func TestContractDriftAuditIsReadOnlyAndPreservesSentinelRows(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)
	testChain.EnsureBlock(500)

	ingestTx(t, 500, addr(fxGameAddr), 0, []*types.Log{
		buildLog(t, gameABI, "EthBidPriceIncreaseDivisorChanged", addr(fxGameAddr), nil, []any{big.NewInt(100)}),
	})
	const sentinel = "0xcccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"
	if _, err := testDB.SQL.Exec(`UPDATE transaction SET tx_hash=$1
		WHERE id=(SELECT tx_id FROM cg_adm_price_inc LIMIT 1)`, sentinel); err != nil {
		t.Fatalf("marking historical sentinel transaction: %v", err)
	}

	stub := driftV1GameStub()
	testChain.RegisterCall(addr(fxGameAddr), stub.Handler())
	testChain.RegisterCall(addr(fxPrizesAddr), prizesWalletDriftStub().Handler())
	before := snapshot(t)

	drifted, err := CheckContractParamsDrift(
		context.Background(), cgRepo, eclient,
		fxGameAddr, fxPrizesAddr, nil,
	)
	if err != nil {
		t.Fatalf("drift audit: %v", err)
	}
	if drifted != 1 {
		t.Fatalf("drifted = %d, want exactly the price divisor", drifted)
	}
	requireNoDiff(t, before, snapshot(t), "read-only drift audit")

	var sentinelRows int
	if err := testDB.SQL.QueryRow(
		"SELECT COUNT(*) FROM transaction WHERE tx_hash=$1", sentinel,
	).Scan(&sentinelRows); err != nil {
		t.Fatal(err)
	}
	if sentinelRows != 1 {
		t.Fatalf("sentinel transaction rows = %d, want preserved", sentinelRows)
	}

	stub.Return("ethBidPriceIncreaseDivisor", big.NewInt(100))
	drifted, err = CheckContractParamsDrift(
		context.Background(), cgRepo, eclient,
		fxGameAddr, fxPrizesAddr, slog.New(slog.DiscardHandler),
	)
	if err != nil || drifted != 0 {
		t.Fatalf("matching drift audit = %d, %v", drifted, err)
	}
	requireNoDiff(t, before, snapshot(t), "matching read-only drift audit")
}

func TestContractDriftAuditV3Configuration(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)
	testChain.EnsureBlock(510)
	ingestTx(t, 510, addr(fxGameAddr), 0, []*types.Log{
		buildLog(t, gameV3ABI, "RoundLateBidDurationDivisorChanged", addr(fxGameAddr), nil, []any{big.NewInt(4)}),
		buildLog(t, gameV3ABI, "RoundLateBidPricePremiumAmountBaseMultiplierChanged", addr(fxGameAddr), nil, []any{big.NewInt(2)}),
		buildLog(t, gameV3ABI, "RoundLateBidPricePremiumAmountExponentChanged", addr(fxGameAddr), nil, []any{big.NewInt(3)}),
		buildLog(t, gameV3ABI, "LastBidderBidCstRewardAmountPercentageChanged", addr(fxGameAddr), nil, []any{big.NewInt(90)}),
		buildLog(t, gameV3ABI, "MainPrizeNumCosmicSignatureNftsChanged", addr(fxGameAddr), nil, []any{big.NewInt(3)}),
	})
	adminEvents, err := cgRepo.AdminEventsInRange(context.Background(), 0, math.MaxInt64)
	if err != nil {
		t.Fatal(err)
	}
	seenTypes := make(map[int64]bool)
	for _, event := range adminEvents {
		seenTypes[event.RecordType] = true
	}
	for recordType := int64(40); recordType <= 44; recordType++ {
		if !seenTypes[recordType] {
			t.Errorf("admin event record type %d was not queryable", recordType)
		}
	}
	testChain.RegisterCall(addr(fxGameAddr), driftV3GameStub().Handler())
	testChain.RegisterCall(addr(fxPrizesAddr), prizesWalletDriftStub().Handler())
	before := snapshot(t)

	drifted, err := CheckContractParamsDrift(
		context.Background(), cgRepo, eclient,
		fxGameAddr, fxPrizesAddr, slog.New(slog.DiscardHandler),
	)
	if err != nil || drifted != 0 {
		t.Fatalf("V3 drift audit = %d, %v", drifted, err)
	}
	requireNoDiff(t, before, snapshot(t), "V3 drift audit")
}

func TestContractDriftAuditFailuresAndUnreadableGetters(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)
	logger := slog.New(slog.DiscardHandler)
	if _, err := CheckContractParamsDrift(
		context.Background(), cgRepo, nil,
		fxGameAddr, fxPrizesAddr, logger,
	); err == nil || !strings.Contains(err.Error(), "eth client is nil") {
		t.Fatalf("nil client error = %v", err)
	}

	testChain.EnsureBlock(520)
	testChain.RegisterCall(addr(fxGameAddr), testchain.MustContractStub(
		cgc.CosmicSignatureGameABI,
		cgc.CosmicSignatureGameV2ABI,
		cgc.CosmicSignatureGameV3ABI,
	).Return("cstDutchAuctionDurationDivisor", big.NewInt(400)).
		Return("cstRewardAmountForBidding", eth(100)).Handler())
	testChain.RegisterCall(addr(fxPrizesAddr), testchain.MustContractStub(cgc.PrizesWalletABI).Handler())
	before := snapshot(t)
	if _, err := CheckContractParamsDrift(
		context.Background(), cgRepo, eclient,
		fxGameAddr, fxPrizesAddr, logger,
	); err != nil {
		t.Fatalf("unreadable getters should be skipped: %v", err)
	}
	requireNoDiff(t, before, snapshot(t), "degraded drift audit")
}
