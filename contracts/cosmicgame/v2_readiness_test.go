package cosmicgame

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func topicHash(signature string) common.Hash {
	return crypto.Keccak256Hash([]byte(signature))
}

func mustABI(t *testing.T, jsonABI string) abi.ABI {
	t.Helper()
	parsed, err := abi.JSON(strings.NewReader(jsonABI))
	if err != nil {
		t.Fatalf("parse ABI: %v", err)
	}
	return parsed
}

func TestV2TopicHashesMatchBindings(t *testing.T) {
	cases := []struct {
		name      string
		signature string
		wantHex   string
	}{
		{"BidPlaced V1", "BidPlaced(uint256,address,int256,int256,int256,string,uint256)", "bcb004d688d0951e50c218ded0d0d574bde915630e29b92987b1f2eab9556549"},
		{"BidPlaced V2", "BidPlaced(uint256,address,int256,int256,int256,string,uint256,uint256,uint256)", "1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec"},
		{"BidCstRewardAmountChanged", "BidCstRewardAmountChanged(uint256)", "96978b83addd498dff54ab50bf4ed5b62e543d07c7935099eafe180248efe4b4"},
		{"BidCstRewardAmountMultiplierChanged", "BidCstRewardAmountMultiplierChanged(uint256)", "40b9c59af8c486ccf8c7cc73df5a51e7cc29747ea7d39f99632ecaf9caa2ed1f"},
		{"CstDutchAuctionDurationChanged", "CstDutchAuctionDurationChanged(uint256)", "4abea08c196329c357e3175d011af39a8625be99ef0ba5a0f3547a95534fedb7"},
		{"CstDutchAuctionDurationChangeDivisorChanged", "CstDutchAuctionDurationChangeDivisorChanged(uint256)", "acbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f"},
		{"CstDutchAuctionDurationDivisorChanged V1", "CstDutchAuctionDurationDivisorChanged(uint256)", "c95d03f6c735a9e59c760fdb88e585aafe0a31b5c034fc7838155287ee32212f"},
		{"CstRewardAmountForBiddingChanged V1", "CstRewardAmountForBiddingChanged(uint256)", "70ad04ce09c925ea466a5f603054f310bba5b7484bba77b382aade0bf93b55d0"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := topicHash(tc.signature).Hex()[2:]
			if got != tc.wantHex {
				t.Fatalf("topic mismatch: got %s want %s", got, tc.wantHex)
			}
		})
	}
}

func TestV2BidPlacedABIContainsEvent(t *testing.T) {
	v2ABI := mustABI(t, CosmicSignatureGameV2ABI)
	ev, ok := v2ABI.Events["BidPlaced"]
	if !ok {
		t.Fatal("CosmicSignatureGameV2ABI missing BidPlaced event")
	}
	if len(ev.Inputs) != 9 {
		t.Fatalf("BidPlaced V2 expected 9 inputs, got %d", len(ev.Inputs))
	}
	got := ev.ID.Hex()[2:]
	want := "1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec"
	if got != want {
		t.Fatalf("BidPlaced event id: got %s want %s", got, want)
	}
}

func TestV2AdminEventsInABI(t *testing.T) {
	v2ABI := mustABI(t, CosmicSignatureGameV2ABI)
	for _, name := range []string{
		"BidCstRewardAmountMultiplierChanged",
		"CstDutchAuctionDurationChanged",
		"CstDutchAuctionDurationChangeDivisorChanged",
	} {
		if _, ok := v2ABI.Events[name]; !ok {
			t.Fatalf("CosmicSignatureGameV2ABI missing event %s", name)
		}
	}
}

func TestUnpackBidPlacedV2Data(t *testing.T) {
	v2ABI := mustABI(t, CosmicSignatureGameV2ABI)
	reward := big.NewInt(1000000000000000000) // 1 CST
	duration := big.NewInt(3600)
	prizeTime := big.NewInt(1700000000)
	data, err := v2ABI.Events["BidPlaced"].Inputs.NonIndexed().Pack(
		big.NewInt(-1),
		big.NewInt(-1),
		"hello v2",
		reward,
		duration,
		prizeTime,
	)
	if err != nil {
		t.Fatalf("pack: %v", err)
	}
	var out CosmicSignatureGameV2BidPlaced
	if err := v2ABI.UnpackIntoInterface(&out, "BidPlaced", data); err != nil {
		t.Fatalf("unpack: %v", err)
	}
	if out.Message != "hello v2" {
		t.Fatalf("message: %q", out.Message)
	}
	if out.BidCstRewardAmount.Cmp(reward) != 0 {
		t.Fatalf("reward: %v", out.BidCstRewardAmount)
	}
	if out.CstDutchAuctionDuration.Cmp(duration) != 0 {
		t.Fatalf("duration: %v", out.CstDutchAuctionDuration)
	}
}

func TestUnpackBidPlacedV1DataStillWorks(t *testing.T) {
	v1ABI := mustABI(t, CosmicSignatureGameABI)
	prizeTime := big.NewInt(1700000000)
	data, err := v1ABI.Events["BidPlaced"].Inputs.NonIndexed().Pack(
		big.NewInt(1000000000000000000),
		big.NewInt(-1),
		"legacy bid",
		prizeTime,
	)
	if err != nil {
		t.Fatalf("pack: %v", err)
	}
	var out CosmicSignatureGameBidPlaced
	if err := v1ABI.UnpackIntoInterface(&out, "BidPlaced", data); err != nil {
		t.Fatalf("unpack: %v", err)
	}
	if out.Message != "legacy bid" {
		t.Fatalf("message: %q", out.Message)
	}
}

func TestAdminUint256EventsUnpack(t *testing.T) {
	v2ABI := mustABI(t, CosmicSignatureGameV2ABI)
	val := big.NewInt(42)
	cases := []struct {
		name string
		out  interface{}
	}{
		{"BidCstRewardAmountMultiplierChanged", &CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged{}},
		{"CstDutchAuctionDurationChanged", &CosmicSignatureGameV2CstDutchAuctionDurationChanged{}},
		{"CstDutchAuctionDurationChangeDivisorChanged", &CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := v2ABI.Events[tc.name].Inputs.Pack(val)
			if err != nil {
				t.Fatalf("pack: %v", err)
			}
			if err := v2ABI.UnpackIntoInterface(tc.out, tc.name, data); err != nil {
				t.Fatalf("unpack: %v", err)
			}
		})
	}
}

func TestV2WrapperCoexistsWithV1InPackage(t *testing.T) {
	if CosmicSignatureGameABI == "" || CosmicSignatureGameV2ABI == "" {
		t.Fatal("both V1 and V2 ABIs must be non-empty")
	}
	if strings.Contains(CosmicSignatureGameV2ABI, "CstDutchAuctionDurationDivisor") {
		t.Fatal("V2 ABI should not expose removed CstDutchAuctionDurationDivisor symbol")
	}
	_ = hex.EncodeToString(topicHash("BidPlaced(uint256,address,int256,int256,int256,string,uint256,uint256,uint256)").Bytes())
}

func TestV2LiveReadMethodsPresent(t *testing.T) {
	v2ABI := mustABI(t, CosmicSignatureGameV2ABI)
	for _, method := range []string{
		"getNextCstBidPrice",
		"getCstDutchAuctionDurations",
		"cstDutchAuctionDuration",
		"bidCstRewardAmountMultiplier",
	} {
		if _, ok := v2ABI.Methods[method]; !ok {
			t.Fatalf("V2 ABI missing method %s", method)
		}
	}
	if _, ok := v2ABI.Methods["cstDutchAuctionDurationDivisor"]; ok {
		t.Fatal("V2 should not have cstDutchAuctionDurationDivisor getter")
	}
}
