package v2

import (
	"testing"
)

// TestGeneratedEnumDomainsMatchSpec pins the generated enum helpers against
// the documented value sets: every spec value is a member and everything
// else is rejected. A regenerated contract that drops or renames a value
// fails here before any handler misbehaves.
func TestGeneratedEnumDomainsMatchSpec(t *testing.T) {
	t.Parallel()

	type membership struct {
		valid   bool
		invalid bool
	}
	cases := map[string]membership{
		"BidType": {
			valid:   BidType("eth").Valid() && BidType("randomWalk").Valid() && BidType("cst").Valid() && BidType("unknown").Valid(),
			invalid: BidType("fiat").Valid(),
		},
		"ClaimAssetType": {
			valid:   ClaimAssetType("eth").Valid() && ClaimAssetType("erc20").Valid() && ClaimAssetType("erc721").Valid(),
			invalid: ClaimAssetType("gold").Valid(),
		},
		"EthDepositSource": {
			valid:   EthDepositSource("raffle").Valid() && EthDepositSource("chronoWarrior").Valid(),
			invalid: EthDepositSource("faucet").Valid(),
		},
		"RoundEthDonationKind": {
			valid:   RoundEthDonationKind("plain").Valid() && RoundEthDonationKind("withInfo").Valid(),
			invalid: RoundEthDonationKind("anonymous").Valid(),
		},
		"RoundStatus": {
			valid:   RoundStatus("open").Valid() && RoundStatus("completed").Valid(),
			invalid: RoundStatus("paused").Valid(),
		},
		"StakingActionType": {
			valid:   StakingActionType("stake").Valid() && StakingActionType("unstake").Valid(),
			invalid: StakingActionType("restake").Valid(),
		},
		"StakerRafflePool": {
			valid:   StakerRafflePool("cst").Valid() && StakerRafflePool("randomWalk").Valid(),
			invalid: StakerRafflePool("bidder").Valid(),
		},
		"RandomWalkTokenSort": {
			valid:   RandomWalkTokenSort("tokenId").Valid() && RandomWalkTokenSort("mostTraded").Valid(),
			invalid: RandomWalkTokenSort("priceAsc").Valid(),
		},
		"RandomWalkOfferSort": {
			valid:   RandomWalkOfferSort("newest").Valid() && RandomWalkOfferSort("oldest").Valid() && RandomWalkOfferSort("priceAsc").Valid() && RandomWalkOfferSort("priceDesc").Valid(),
			invalid: RandomWalkOfferSort("tokenId").Valid(),
		},
		"RandomWalkOfferSide": {
			valid:   RandomWalkOfferSide("sell").Valid() && RandomWalkOfferSide("buy").Valid(),
			invalid: RandomWalkOfferSide("bid").Valid(),
		},
		"RandomWalkOfferStatus": {
			valid:   RandomWalkOfferStatus("active").Valid() && RandomWalkOfferStatus("bought").Valid() && RandomWalkOfferStatus("canceled").Valid(),
			invalid: RandomWalkOfferStatus("expired").Valid(),
		},
		"RandomWalkTokenEventType": {
			valid: RandomWalkTokenEventType("mint").Valid() && RandomWalkTokenEventType("transfer").Valid() &&
				RandomWalkTokenEventType("nameChange").Valid() && RandomWalkTokenEventType("listed").Valid() &&
				RandomWalkTokenEventType("offerCanceled").Valid() && RandomWalkTokenEventType("purchase").Valid(),
			invalid: RandomWalkTokenEventType("burn").Valid(),
		},
		"ContractConfiguration0CstBidRewardMode": {
			valid:   ContractConfiguration0CstBidRewardMode("fixed").Valid(),
			invalid: ContractConfiguration0CstBidRewardMode("dynamic").Valid(),
		},
		"ContractConfiguration0CstRoundStartAuctionMode": {
			valid:   ContractConfiguration0CstRoundStartAuctionMode("divisor").Valid(),
			invalid: ContractConfiguration0CstRoundStartAuctionMode("durationSeconds").Valid(),
		},
		"ContractConfiguration0MechanicsVersion": {
			valid:   ContractConfiguration0MechanicsVersion("v1").Valid(),
			invalid: ContractConfiguration0MechanicsVersion("v2").Valid(),
		},
		"ContractConfiguration1CstBidRewardMode": {
			valid:   ContractConfiguration1CstBidRewardMode("dynamic").Valid(),
			invalid: ContractConfiguration1CstBidRewardMode("fixed").Valid(),
		},
		"ContractConfiguration1CstRoundStartAuctionMode": {
			valid:   ContractConfiguration1CstRoundStartAuctionMode("durationSeconds").Valid(),
			invalid: ContractConfiguration1CstRoundStartAuctionMode("divisor").Valid(),
		},
		"ContractConfiguration1MechanicsVersion": {
			valid:   ContractConfiguration1MechanicsVersion("v2").Valid(),
			invalid: ContractConfiguration1MechanicsVersion("v1").Valid(),
		},
	}
	for name, got := range cases {
		if !got.valid {
			t.Errorf("%s rejected a documented spec value", name)
		}
		if got.invalid {
			t.Errorf("%s accepted an undocumented value", name)
		}
	}
}

// TestContractConfigurationVariantAccessors pins the generated oneOf
// surface: each mechanics variant survives a From/As round trip and the
// two variants stay mutually exclusive in their enum domains.
func TestContractConfigurationVariantAccessors(t *testing.T) {
	t.Parallel()

	v1Mode := Fixed
	v1Auction := ContractConfiguration0CstRoundStartAuctionModeDivisor
	v1Mechanics := ContractConfiguration0MechanicsVersionV1
	v1Variant := ContractConfiguration0{
		CstBidRewardMode:         &v1Mode,
		CstRoundStartAuctionMode: &v1Auction,
		MechanicsVersion:         &v1Mechanics,
	}
	var union ContractConfiguration
	if err := union.FromContractConfiguration0(v1Variant); err != nil {
		t.Fatalf("FromContractConfiguration0: %v", err)
	}
	roundTripped, err := union.AsContractConfiguration0()
	if err != nil {
		t.Fatalf("AsContractConfiguration0: %v", err)
	}
	if roundTripped.CstBidRewardMode == nil || *roundTripped.CstBidRewardMode != Fixed ||
		roundTripped.MechanicsVersion == nil || !roundTripped.MechanicsVersion.Valid() {
		t.Fatalf("v1 variant round trip = %+v", roundTripped)
	}

	v2Mode := ContractConfiguration1CstBidRewardModeDynamic
	v2Auction := ContractConfiguration1CstRoundStartAuctionModeDurationSeconds
	v2Mechanics := V2
	v2Variant := ContractConfiguration1{
		CstBidRewardMode:         &v2Mode,
		CstRoundStartAuctionMode: &v2Auction,
		MechanicsVersion:         &v2Mechanics,
	}
	if err := union.MergeContractConfiguration1(v2Variant); err != nil {
		t.Fatalf("MergeContractConfiguration1: %v", err)
	}
	merged, err := union.AsContractConfiguration1()
	if err != nil {
		t.Fatalf("AsContractConfiguration1: %v", err)
	}
	if merged.CstBidRewardMode == nil || *merged.CstBidRewardMode != ContractConfiguration1CstBidRewardModeDynamic ||
		merged.CstRoundStartAuctionMode == nil || *merged.CstRoundStartAuctionMode != ContractConfiguration1CstRoundStartAuctionModeDurationSeconds {
		t.Fatalf("v2 variant merge = %+v", merged)
	}

	v3Mode := ContractConfiguration2CstBidRewardModeDynamic
	v3Auction := ContractConfiguration2CstRoundStartAuctionModeDurationSeconds
	v3Mechanics := V3
	v3Variant := ContractConfiguration2{
		CstBidRewardMode:         &v3Mode,
		CstRoundStartAuctionMode: &v3Auction,
		MechanicsVersion:         &v3Mechanics,
	}
	if err := union.MergeContractConfiguration2(v3Variant); err != nil {
		t.Fatalf("MergeContractConfiguration2: %v", err)
	}
	v3RoundTripped, err := union.AsContractConfiguration2()
	if err != nil || v3RoundTripped.MechanicsVersion == nil ||
		*v3RoundTripped.MechanicsVersion != V3 {
		t.Fatalf("v3 variant merge = %+v, err=%v", v3RoundTripped, err)
	}

	var fresh ContractConfiguration
	if err := fresh.FromContractConfiguration1(v2Variant); err != nil {
		t.Fatalf("FromContractConfiguration1: %v", err)
	}
	if err := fresh.MergeContractConfiguration0(v1Variant); err != nil {
		t.Fatalf("MergeContractConfiguration0: %v", err)
	}
	overwritten, err := fresh.AsContractConfiguration0()
	if err != nil || overwritten.MechanicsVersion == nil ||
		*overwritten.MechanicsVersion != ContractConfiguration0MechanicsVersionV1 {
		t.Fatalf("merged v1 variant = %+v, err=%v", overwritten, err)
	}
}
