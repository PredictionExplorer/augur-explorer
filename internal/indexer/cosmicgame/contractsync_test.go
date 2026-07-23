package cosmicgame

import (
	"context"
	"errors"
	"math/big"
	"slices"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

var syncTestGameAddr = ethcommon.HexToAddress("0x2000000000000000000000000000000000000002")

func newSyncBindings(
	t *testing.T,
	stub *testchain.ContractStub,
) (*cgc.CosmicSignatureGame, *cgc.CosmicSignatureGameV2, *cgc.CosmicSignatureGameV3) {
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
		t.Fatal(err)
	}
	v2, err := cgc.NewCosmicSignatureGameV2(syncTestGameAddr, client)
	if err != nil {
		t.Fatal(err)
	}
	v3, err := cgc.NewCosmicSignatureGameV3(syncTestGameAddr, client)
	if err != nil {
		t.Fatal(err)
	}
	return v1, v2, v3
}

func gameStub() *testchain.ContractStub {
	return testchain.MustContractStub(
		cgc.CosmicSignatureGameABI,
		cgc.CosmicSignatureGameV2ABI,
		cgc.CosmicSignatureGameV3ABI,
	)
}

func TestProbeContractMechanicsNewestFirst(t *testing.T) {
	if got := probeContractMechanics(nil, nil, nil, &bind.CallOpts{}); got != contractMechanicsUnknown {
		t.Fatalf("nil mechanics = %d", got)
	}

	v1, v2, v3 := newSyncBindings(t, gameStub().
		Return("mainPrizeNumCosmicSignatureNfts", big.NewInt(3)).
		Return("cstDutchAuctionDurationChangeDivisor", big.NewInt(25)).
		Return("cstDutchAuctionDurationDivisor", big.NewInt(400)))
	if got := probeContractMechanics(v1, v2, v3, &bind.CallOpts{}); got != contractMechanicsV3 {
		t.Fatalf("mechanics = %d, want V3", got)
	}

	v1, v2, v3 = newSyncBindings(t, gameStub().
		Return("cstDutchAuctionDurationChangeDivisor", big.NewInt(25)).
		Return("cstDutchAuctionDurationDivisor", big.NewInt(400)))
	if got := probeContractMechanics(v1, v2, v3, &bind.CallOpts{}); got != contractMechanicsV2 {
		t.Fatalf("mechanics = %d, want V2", got)
	}

	v1, v2, v3 = newSyncBindings(t, gameStub().
		Return("cstDutchAuctionDurationDivisor", big.NewInt(400)))
	if got := probeContractMechanics(v1, v2, v3, &bind.CallOpts{}); got != contractMechanicsV1 {
		t.Fatalf("mechanics = %d, want V1", got)
	}
}

func TestReadCstRewardPerMechanics(t *testing.T) {
	v1, v2, v3 := newSyncBindings(t, gameStub().
		Return("bidCstRewardAmountMultiplier", big.NewInt(7)).
		Return("cstRewardAmountForBidding", big.NewInt(5)))
	opts := &bind.CallOpts{}
	for _, test := range []struct {
		mechanics int64
		want      string
	}{
		{contractMechanicsV1, "5"},
		{contractMechanicsV2, "7"},
		{contractMechanicsV3, "7"},
	} {
		got, err := readCstReward(v1, v2, v3, opts, test.mechanics)
		if err != nil || got != test.want {
			t.Errorf("readCstReward(v%d) = %q, %v; want %q", test.mechanics, got, err, test.want)
		}
	}
}

func TestReadDelayDurationPrefersV3(t *testing.T) {
	v1, v2, v3 := newSyncBindings(t,
		gameStub().Return("delayDurationBeforeRoundActivation", big.NewInt(1234)))
	got, err := readDelayDuration(v1, v2, v3, &bind.CallOpts{})
	if err != nil || got.Int64() != 1234 {
		t.Fatalf("readDelayDuration = %v, %v", got, err)
	}
}

func TestV3DriftRegistryUsesFinalNames(t *testing.T) {
	parameters := buildContractParamSyncList(contractMechanicsV3)
	names := make([]string, 0, len(parameters))
	tables := make([]string, 0, len(parameters))
	for _, parameter := range parameters {
		names = append(names, parameter.name)
		tables = append(tables, parameter.table)
	}
	for _, want := range []string{
		"round_late_bid_duration_divisor",
		"round_late_bid_price_premium_base_multiplier",
		"round_late_bid_price_premium_exponent",
		"last_bidder_bid_cst_reward_amount_percentage",
		"main_prize_num_cosmic_signature_nfts",
	} {
		if !slices.Contains(names, want) {
			t.Errorf("V3 drift registry missing %s", want)
		}
	}
	if slices.Contains(names, "bid_cst_reward_amount_per_minute") ||
		slices.Contains(tables, "cg_adm_bid_cst_reward_per_min") {
		t.Fatal("V3 drift registry contains obsolete per-minute reward naming")
	}
	if slices.Contains(names, "cst_dutch_auction_duration") {
		t.Fatal("V3 drift audit compares the inert V2 duration slot")
	}
}

func TestV3DriftReaders(t *testing.T) {
	v1, v2, v3 := newSyncBindings(t, gameStub().
		Return("roundLateBidDurationDivisor", big.NewInt(4)).
		Return("roundLateBidPricePremiumAmountBaseMultiplier", big.NewInt(2)).
		Return("roundLateBidPricePremiumAmountExponent", big.NewInt(3)).
		Return("lastBidderBidCstRewardAmountPercentage", big.NewInt(90)).
		Return("mainPrizeNumCosmicSignatureNfts", big.NewInt(3)))
	opts := &bind.CallOpts{}
	got := map[string]string{}
	for _, parameter := range buildContractParamSyncList(contractMechanicsV3) {
		switch parameter.table {
		case "cg_adm_late_bid_dur_divisor",
			"cg_adm_late_bid_premium_base_mul",
			"cg_adm_late_bid_premium_exponent",
			"cg_adm_last_bidder_reward_pct",
			"cg_adm_main_prize_num_nfts":
			value, err := parameter.read(v1, v2, v3, opts)
			if err != nil {
				t.Fatalf("%s reader: %v", parameter.name, err)
			}
			got[parameter.table] = value
		}
	}
	want := map[string]string{
		"cg_adm_late_bid_dur_divisor":      "4",
		"cg_adm_late_bid_premium_base_mul": "2",
		"cg_adm_late_bid_premium_exponent": "3",
		"cg_adm_last_bidder_reward_pct":    "90",
		"cg_adm_main_prize_num_nfts":       "3",
	}
	for table, value := range want {
		if got[table] != value {
			t.Errorf("%s = %q, want %q", table, got[table], value)
		}
	}
}

func TestContractDriftReaderFailurePaths(t *testing.T) {
	opts := &bind.CallOpts{}
	if _, err := readDelayDuration(nil, nil, nil, opts); err == nil {
		t.Fatal("readDelayDuration without bindings succeeded")
	}
	if _, err := readCstReward(nil, nil, nil, opts, contractMechanicsUnknown); err == nil {
		t.Fatal("readCstReward without bindings succeeded")
	}

	v1Big := buildContractParamSyncList(contractMechanicsV1)[1]
	if _, err := v1Big.read(nil, nil, nil, opts); err == nil {
		t.Fatal("V1 parameter reader without V1 binding succeeded")
	}

	v1, v2, v3 := newSyncBindings(t, gameStub().
		Return("cstDutchAuctionDuration", big.NewInt(1800)).
		Return("cstDutchAuctionDurationChangeDivisor", big.NewInt(25)))
	var durationParam contractParamSync
	for _, parameter := range buildContractParamSyncList(contractMechanicsV2) {
		if parameter.name == "cst_dutch_auction_duration" {
			durationParam = parameter
			break
		}
	}
	if got, err := durationParam.read(v1, v2, v3, opts); err != nil || got != "1800" {
		t.Fatalf("V2 duration reader = %q, %v", got, err)
	}

	failingV1, failingV2, failingV3 := newSyncBindings(t, gameStub())
	if _, err := durationParam.read(failingV1, failingV2, failingV3, opts); err == nil {
		t.Fatal("unreadable V2 duration succeeded")
	}
	v3Parameter := v3ContractParam("test", "test_table",
		func(*cgc.CosmicSignatureGameV3, *bind.CallOpts) (*big.Int, error) {
			return nil, errors.New("forced read failure")
		})
	if _, err := v3Parameter.read(nil, nil, nil, opts); err == nil {
		t.Fatal("V3 parameter reader without V3 binding succeeded")
	}
	if _, err := v3Parameter.read(v1, v2, v3, opts); err == nil {
		t.Fatal("V3 parameter reader swallowed getter failure")
	}
}
