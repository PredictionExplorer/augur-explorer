package cosmicgame

import (
	"math/big"
	"testing"
)

func TestV3TopicHashesMatchBindings(t *testing.T) {
	cases := []struct {
		signature string
		wantHex   string
	}{
		{
			"MainPrizeClaimed(uint256,address,uint256,uint256,uint256,uint256,uint256)",
			"9314e78538382a9fc6cd54cee107c8a7721d172c43b2e48c2e7a51f200fc3788",
		},
		{
			"RoundLateBidDurationDivisorChanged(uint256)",
			"7acba37d1b2d934e554139ff1d470d2cce50a7b6e56870577249b87564a95a69",
		},
		{
			"RoundLateBidPricePremiumAmountBaseMultiplierChanged(uint256)",
			"169f25ec19cc5b518dc57adf05bb2d85155b1d6c60767e9f0e3ced18eac0ab77",
		},
		{
			"RoundLateBidPricePremiumAmountExponentChanged(uint256)",
			"cb78cca7628d232a9c7beef53b62f7204d9eacb44de85a8f593e6b0bb72a1621",
		},
		{
			"CstBidPriceDeclineMultiplierChanged(uint256)",
			"5a7755107bc57392a4597c870f95d2c7cd7802e3e92c8aa549888e0b4a66d19c",
		},
		{
			"CstBidPriceDeclineMultiplierChangeDivisorChanged(uint256)",
			"dca564eaecef774ca88453bd3a492a26c40497ab897aa02ad10eb030e126c5ce",
		},
		{
			"MainPrizeNumCosmicSignatureNftsChanged(uint256)",
			"616bfcaa6490f55f6e57a4deedac1db04d0d6826deb84fad86cc43439bcf3564",
		},
	}
	v3ABI := mustABI(t, CosmicSignatureGameV3ABI)
	for _, tc := range cases {
		t.Run(tc.signature, func(t *testing.T) {
			if got := topicHash(tc.signature).Hex()[2:]; got != tc.wantHex {
				t.Fatalf("topic hash = %s, want %s", got, tc.wantHex)
			}
			name := tc.signature[:len(tc.signature)-len("(uint256)")]
			if tc.signature[:16] == "MainPrizeClaimed" {
				name = "MainPrizeClaimed"
			}
			event, ok := v3ABI.Events[name]
			if !ok {
				t.Fatalf("CosmicSignatureGameV3 ABI missing %s", name)
			}
			if got := event.ID.Hex()[2:]; got != tc.wantHex {
				t.Fatalf("ABI event ID = %s, want %s", got, tc.wantHex)
			}
		})
	}
}

func TestV3MainPrizeClaimedShapeAndDecode(t *testing.T) {
	v3ABI := mustABI(t, CosmicSignatureGameV3ABI)
	event, ok := v3ABI.Events["MainPrizeClaimed"]
	if !ok {
		t.Fatal("CosmicSignatureGameV3 ABI missing MainPrizeClaimed")
	}
	if len(event.Inputs) != 7 {
		t.Fatalf("MainPrizeClaimed inputs = %d, want 7", len(event.Inputs))
	}
	var indexed int
	for _, input := range event.Inputs {
		if input.Indexed {
			indexed++
		}
	}
	if indexed != 3 {
		t.Fatalf("MainPrizeClaimed indexed inputs = %d, want 3", indexed)
	}

	ethAmount := big.NewInt(11)
	cstAmount := big.NewInt(22)
	numNFTs := big.NewInt(3)
	timeout := big.NewInt(44)
	data, err := event.Inputs.NonIndexed().Pack(ethAmount, cstAmount, numNFTs, timeout)
	if err != nil {
		t.Fatalf("pack V3 MainPrizeClaimed: %v", err)
	}
	var decoded CosmicSignatureGameV3MainPrizeClaimed
	if err := v3ABI.UnpackIntoInterface(&decoded, "MainPrizeClaimed", data); err != nil {
		t.Fatalf("unpack V3 MainPrizeClaimed: %v", err)
	}
	if decoded.EthPrizeAmount.Cmp(ethAmount) != 0 ||
		decoded.CstPrizeAmount.Cmp(cstAmount) != 0 ||
		decoded.PrizeNumCosmicSignatureNfts.Cmp(numNFTs) != 0 ||
		decoded.TimeoutTimeToWithdrawSecondaryPrizes.Cmp(timeout) != 0 {
		t.Fatalf("decoded V3 claim = %+v", decoded)
	}
}

func TestV3LatestReadMethodsPresent(t *testing.T) {
	v3ABI := mustABI(t, CosmicSignatureGameV3ABI)
	for _, method := range []string{
		"mainPrizeNumCosmicSignatureNfts",
		"roundLateBidDurationDivisor",
		"roundLateBidPricePremiumAmountBaseMultiplier",
		"roundLateBidPricePremiumAmountExponent",
		"getRoundLateBidDuration",
		"championDurations",
		"getCstDutchAuctionDurations",
		"cstDutchAuctionBeginningBidPriceMinLimit",
		"cstBidPriceDeclineMultiplier",
		"cstBidPriceDeclineMultiplierChangeDivisor",
		"bidRaffleCumulativeWeights",
	} {
		if _, ok := v3ABI.Methods[method]; !ok {
			t.Errorf("CosmicSignatureGameV3 ABI missing method %s", method)
		}
	}
	if _, ok := v3ABI.Methods["bidCstRewardAmountPerMinute"]; ok {
		t.Fatal("V3 ABI contains obsolete bidCstRewardAmountPerMinute getter")
	}
	if _, ok := v3ABI.Events["BidCstRewardAmountPerMinuteChanged"]; ok {
		t.Fatal("V3 ABI contains obsolete BidCstRewardAmountPerMinuteChanged event")
	}
	// Retired on the v3-2026-07-24 contracts: the bid CST reward is minted
	// entirely to the outbid previous bidder, and the CST Dutch auction is
	// driven by cstBidPriceDeclineMultiplier instead of a stored duration.
	for _, method := range []string{
		"lastBidderBidCstRewardAmountPercentage",
		"getCstDutchAuctionBeginningBidPriceMinLimit",
		"getBidCstRewardAmountPerMainPrizeTimeIncrement",
	} {
		if _, ok := v3ABI.Methods[method]; ok {
			t.Errorf("V3 ABI contains retired method %s", method)
		}
	}
	if _, ok := v3ABI.Events["LastBidderBidCstRewardAmountPercentageChanged"]; ok {
		t.Fatal("V3 ABI contains retired LastBidderBidCstRewardAmountPercentageChanged event")
	}
	if CosmicSignatureGameV3Bin == "" {
		t.Fatal("CosmicSignatureGameV3 creation bytecode is empty")
	}
}
