package v2

import (
	"strings"
	"testing"
	"time"

	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestMapGlobalStatistics(t *testing.T) {
	t.Parallel()
	record := validGlobalStatisticsRecord()
	got, err := mapGlobalStatistics(record)
	if err != nil {
		t.Fatalf("mapGlobalStatistics: %v", err)
	}
	if got.TotalBids != 10 || got.TotalPrizesPaidWei != "1000" ||
		got.CstStaking.TotalRewardWei != "40" ||
		len(got.DonatedTokenDistribution) != 1 ||
		got.DonatedTokenDistribution[0].TokenAddress != "0x3333333333333333333333333333333333333333" {
		t.Fatalf("global statistics = %+v", got)
	}
}

func TestMapGlobalStatisticsRejectsInvalidRecords(t *testing.T) {
	t.Parallel()
	tests := map[string]func(*cgstore.GlobalStatisticsRecord){
		"uint overflow": func(record *cgstore.GlobalStatisticsRecord) {
			record.TotalBids = ^uint64(0)
		},
		"negative count": func(record *cgstore.GlobalStatisticsRecord) {
			record.UniqueDonors = -1
		},
		"invalid amount": func(record *cgstore.GlobalStatisticsRecord) {
			record.TotalEthDonatedWei = "1.5"
		},
		"invalid distribution address": func(record *cgstore.GlobalStatisticsRecord) {
			record.DonatedTokenDistribution[0].ContractAddr = "bad"
		},
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validGlobalStatisticsRecord()
			mutate(&record)
			if _, err := mapGlobalStatistics(record); err == nil {
				t.Fatal("mapper accepted invalid statistics")
			}
		})
	}
}

func TestMapCounters(t *testing.T) {
	t.Parallel()
	got, err := mapCounters(cgprimitives.CGRecordCounters{
		TotalBids: 4, TotalPrizes: 2, TotalDonatedNFTs: 1,
	})
	if err != nil || got.TotalBids != 4 || got.CompletedRounds != 2 || got.DonatedNfts != 1 {
		t.Fatalf("counters = %+v, err=%v", got, err)
	}
	if _, err := mapCounters(cgprimitives.CGRecordCounters{TotalBids: -1}); err == nil {
		t.Fatal("accepted negative counters")
	}
}

func TestMapROILeaderboardEntry(t *testing.T) {
	t.Parallel()
	record := validROILeaderboardRecord()
	record.NetProfitWei = "-0005"
	record.ROIRatio = "-00.5000"
	record.WinRateRatio = "01.2500"
	got, err := mapROILeaderboardEntry(record)
	if err != nil {
		t.Fatalf("mapROILeaderboardEntry: %v", err)
	}
	if got.NetProfitWei != "-5" || got.RoiRatio != "-0.5" ||
		got.WinRateRatio != "1.25" || got.BidderAddress == "" {
		t.Fatalf("ROI entry = %+v", got)
	}
}

func TestMapROILeaderboardEntryRejectsInvalidRecords(t *testing.T) {
	t.Parallel()
	tests := map[string]func(*cgstore.ROILeaderboardRecord){
		"invalid bidder":    func(record *cgstore.ROILeaderboardRecord) { record.BidderAddr = "bad" },
		"negative count":    func(record *cgstore.ROILeaderboardRecord) { record.NumBids = -1 },
		"invalid amount":    func(record *cgstore.ROILeaderboardRecord) { record.TotalEthSpentWei = "bad" },
		"invalid ratio":     func(record *cgstore.ROILeaderboardRecord) { record.ROIRatio = "1e3" },
		"negative win rate": func(record *cgstore.ROILeaderboardRecord) { record.WinRateRatio = "-1" },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validROILeaderboardRecord()
			mutate(&record)
			if _, err := mapROILeaderboardEntry(record); err == nil {
				t.Fatal("mapper accepted invalid ROI record")
			}
		})
	}
}

func TestMapClaimSummary(t *testing.T) {
	t.Parallel()
	record := validClaimSummaryRecord(2)
	now := time.Unix(record.ClaimWindowTimeout+1, 0)
	got, err := mapClaimSummary(record, now)
	if err != nil {
		t.Fatalf("mapClaimSummary: %v", err)
	}
	if got.Round != 2 || !got.IsExpired || got.TotalAwarded != 3 ||
		got.UnclaimedEthAmountWei != "10" ||
		!got.AwardedAt.Equal(time.Unix(record.AwardedTimestamp, 0).UTC()) {
		t.Fatalf("claim summary = %+v", got)
	}
	record.TotalAwarded++
	if _, err := mapClaimSummary(record, now); err == nil {
		t.Fatal("accepted inconsistent claim totals")
	}
}

func TestMapClaimAssetVariants(t *testing.T) {
	t.Parallel()
	ethAmount := "10"
	tokenAddress := "0x3333333333333333333333333333333333333333"
	tokenID := int64(7)
	baseUnits := "20"
	for _, tc := range []struct {
		name   string
		record cgstore.ClaimTransactionRecord
	}{
		{
			name: "eth",
			record: validClaimTransactionRecord(cgstore.ClaimAssetETH, func(record *cgstore.ClaimTransactionRecord) {
				record.EthAmountWei = &ethAmount
			}),
		},
		{
			name: "erc721",
			record: validClaimTransactionRecord(cgstore.ClaimAssetERC721, func(record *cgstore.ClaimTransactionRecord) {
				record.TokenAddr, record.TokenID = &tokenAddress, &tokenID
			}),
		},
		{
			name: "erc20",
			record: validClaimTransactionRecord(cgstore.ClaimAssetERC20, func(record *cgstore.ClaimTransactionRecord) {
				record.TokenAddr, record.AmountBaseUnits = &tokenAddress, &baseUnits
			}),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := mapAssetClaimTransaction(tc.record)
			if err != nil || !got.AssetType.Valid() {
				t.Fatalf("claim transaction = %+v, err=%v", got, err)
			}
		})
	}
}

func TestMapAttachedAndUnclaimedAssets(t *testing.T) {
	t.Parallel()
	tokenAddress := "0x3333333333333333333333333333333333333333"
	tokenID := int64(7)
	baseUnits := "20"
	ethAmount := "10"
	attachedNFT := validAttachedTokenRecord(cgstore.ClaimAssetERC721)
	attachedNFT.TokenID = &tokenID
	if _, err := mapAttachedToken(attachedNFT); err != nil {
		t.Fatalf("map attached NFT: %v", err)
	}
	attachedERC20 := validAttachedTokenRecord(cgstore.ClaimAssetERC20)
	attachedERC20.AmountBaseUnits = &baseUnits
	if _, err := mapAttachedToken(attachedERC20); err != nil {
		t.Fatalf("map attached ERC20: %v", err)
	}
	for _, record := range []cgstore.UnclaimedItemRecord{
		validUnclaimedItemRecord(cgstore.ClaimAssetETH, func(record *cgstore.UnclaimedItemRecord) {
			record.EthAmountWei = &ethAmount
		}),
		validUnclaimedItemRecord(cgstore.ClaimAssetERC721, func(record *cgstore.UnclaimedItemRecord) {
			record.TokenAddr, record.TokenID = &tokenAddress, &tokenID
		}),
		validUnclaimedItemRecord(cgstore.ClaimAssetERC20, func(record *cgstore.UnclaimedItemRecord) {
			record.TokenAddr, record.AmountBaseUnits = &tokenAddress, &baseUnits
		}),
	} {
		if _, err := mapUnclaimedItem(record); err != nil {
			t.Fatalf("map unclaimed %s: %v", record.AssetType, err)
		}
	}
}

func TestCanonicalDecimal(t *testing.T) {
	t.Parallel()
	tests := map[string]string{
		"0":       "0",
		"000":     "0",
		"01.2300": "1.23",
		"-00.50":  "-0.5",
		"-0.0":    "0",
	}
	for input, want := range tests {
		got, err := canonicalDecimal(input, true)
		if err != nil || got != want {
			t.Errorf("canonicalDecimal(%q) = %q, %v; want %q", input, got, err, want)
		}
	}
	for _, input := range []string{"", ".1", "1.", "1e2", "--1", "1.2.3"} {
		if _, err := canonicalDecimal(input, true); err == nil {
			t.Errorf("canonicalDecimal accepted %q", input)
		}
	}
}

func validGlobalStatisticsRecord() cgstore.GlobalStatisticsRecord {
	// #nosec G101 -- deterministic chain fixture values, not credentials.
	return cgstore.GlobalStatisticsRecord{
		TotalBids:                          10,
		CurrentRoundBids:                   2,
		CompletedRounds:                    3,
		TotalPrizeAwards:                   4,
		PrizeRegistryRows:                  5,
		UniqueBidders:                      6,
		UniqueWinners:                      7,
		UniqueDonors:                       8,
		UniqueCSTStakers:                   9,
		UniqueRandomWalkStakers:            10,
		UniqueDualStakers:                  4,
		TotalPrizesPaidWei:                 "1000",
		TotalEthDonatedWei:                 "2000",
		VoluntaryDonationCount:             2,
		VoluntaryDonationsTotalWei:         "20",
		CosmicGameDonationCount:            3,
		CosmicGameDonationsTotalWei:        "30",
		DirectDonationCount:                4,
		DirectDonationsTotalWei:            "40",
		CharityWithdrawalCount:             5,
		CharityWithdrawalsTotalWei:         "50",
		RandomWalkTokensUsedInBids:         6,
		DonatedNFTCount:                    7,
		CosmicSignatureMints:               8,
		NamedTokens:                        9,
		RaffleEthDepositsTotalWei:          "60",
		RaffleEthWithdrawnTotalWei:         "61",
		ChronoWarriorEthDepositsTotalWei:   "62",
		CSTGivenInPrizesTotalWei:           "70",
		WinnersWithPendingRaffleWithdrawal: 1,
		CSTConsumedTotalWei:                "80",
		CSTBidCount:                        3,
		MarketingRewardsTotalWei:           "90",
		MarketingRewardCount:               2,
		DonatedTokenDistribution: []cgprimitives.CGDonatedTokenDistrRec{{
			ContractAddr:     "0x3333333333333333333333333333333333333333",
			NumDonatedTokens: 2,
		}},
		CSTStaking: cgprimitives.CGStakeStatsCST{
			TotalTokensStaked: 3,
			TotalReward:       "40",
			UnclaimedReward:   "10",
			NumActiveStakers:  2,
			NumDeposits:       4,
		},
		RandomWalkStaking: cgprimitives.CGStakeStatsRWalk{
			TotalTokensStaked: 2,
			TotalTokensMinted: 1,
			NumActiveStakers:  1,
		},
	}
}

func validROILeaderboardRecord() cgstore.ROILeaderboardRecord {
	return cgstore.ROILeaderboardRecord{
		BidderAid:          2,
		BidderAddr:         "0x2222222222222222222222222222222222222222",
		NumBids:            10,
		RoundsParticipated: 4,
		RoundsWon:          2,
		WinRateRatio:       "0.5",
		TotalEthSpentWei:   "100",
		TotalCSTSpentWei:   "200",
		EthWonWei:          "150",
		PrizesCount:        3,
		CSTPrizesCount:     1,
		NFTPrizesCount:     1,
		NetProfitWei:       "50",
		ROIRatio:           "0.5",
	}
}

func validClaimSummaryRecord(round int64) cgstore.ClaimSummaryRecord {
	return cgstore.ClaimSummaryRecord{
		RoundNum:              round,
		EventLogID:            100 + round,
		ClaimWindowTimeout:    1_767_227_000,
		AwardedTimestamp:      1_767_226_000,
		EthAwarded:            1,
		EthUnclaimed:          1,
		UnclaimedEthAmountWei: "10",
		NFTAwarded:            1,
		NFTUnclaimed:          0,
		ERC20Awarded:          1,
		ERC20Unclaimed:        0,
		TotalAwarded:          3,
		TotalUnclaimed:        1,
		AvgClaimLatencySecs:   50,
	}
}

func validClaimTransactionRecord(
	asset cgstore.ClaimAssetType,
	mutate func(*cgstore.ClaimTransactionRecord),
) cgstore.ClaimTransactionRecord {
	record := cgstore.ClaimTransactionRecord{
		EventLogID:       100,
		RoundNum:         1,
		AssetType:        asset,
		RecipientAddr:    "0x2222222222222222222222222222222222222222",
		BeneficiaryAddr:  "0x2222222222222222222222222222222222222222",
		ClaimedAfterSecs: 10,
		ClaimedTimestamp: 1_767_226_000,
		TxHash:           "0x" + strings.Repeat("ab", 32),
	}
	mutate(&record)
	return record
}

func validAttachedTokenRecord(asset cgstore.ClaimAssetType) cgstore.AttachedTokenRecord {
	return cgstore.AttachedTokenRecord{
		EventLogID:      100,
		RoundNum:        1,
		AssetType:       asset,
		ContributorAddr: "0x2222222222222222222222222222222222222222",
		TokenAddr:       "0x3333333333333333333333333333333333333333",
		OccurredAt:      1_767_226_000,
		TxHash:          "0x" + strings.Repeat("ab", 32),
	}
}

func validUnclaimedItemRecord(
	asset cgstore.ClaimAssetType,
	mutate func(*cgstore.UnclaimedItemRecord),
) cgstore.UnclaimedItemRecord {
	record := cgstore.UnclaimedItemRecord{
		Segment:       0,
		Key:           100,
		RoundNum:      1,
		AssetType:     asset,
		RecipientAddr: "0x2222222222222222222222222222222222222222",
	}
	mutate(&record)
	return record
}
