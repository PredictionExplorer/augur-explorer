package v2

import (
	"encoding/json"
	"math"
	"strings"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

func TestMapRoundSummary(t *testing.T) {
	t.Parallel()

	got, err := mapRoundSummary(validRoundRecord())
	if err != nil {
		t.Fatalf("mapRoundSummary: %v", err)
	}
	if got.Round != 0 || got.Status != Completed || got.TotalBids != 4 {
		t.Fatalf("summary identity/statistics = %+v", got)
	}
	if got.CompletedAt.String() != "2026-01-01 05:16:40 +0000 UTC" {
		t.Errorf("CompletedAt = %s", got.CompletedAt)
	}
	if got.MainPrize.EthAmountWei != "500000000000000000" ||
		got.MainPrize.CstAmountWei == nil ||
		*got.MainPrize.CstAmountWei != "100000000000000000000" {
		t.Errorf("MainPrize = %+v", got.MainPrize)
	}
	if got.RaffleEthDepositsWei != "100000000000000000" ||
		got.RaffleNftCount != 2 ||
		got.DonatedNftCount != 1 ||
		got.Erc20DonationCount != 1 {
		t.Errorf("summary aggregates = %+v", got)
	}
}

func TestMapRoundSummaryRejectsInvalidStoreDataExactly(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		mutate func(*cgmodel.CGRoundRec)
		want   string
	}{
		{
			name:   "round overflow",
			mutate: func(record *cgmodel.CGRoundRec) { record.RoundNum = math.MaxUint64 },
			want:   "round number exceeds int64",
		},
		{
			name: "claim identity",
			mutate: func(record *cgmodel.CGRoundRec) {
				record.ClaimPrizeTx.Tx.EvtLogId = 0
			},
			want: "invalid claim transaction identity",
		},
		{
			name: "main prize",
			mutate: func(record *cgmodel.CGRoundRec) {
				record.MainPrize.WinnerAddr = "not-an-address"
			},
			want: "invalid main-prize winner address",
		},
		{
			name: "aggregate count",
			mutate: func(record *cgmodel.CGRoundRec) {
				record.RoundStats.TotalRaffleNFTs = -1
			},
			want: "negative round aggregate",
		},
		{
			name: "raffle amount",
			mutate: func(record *cgmodel.CGRoundRec) {
				record.RoundStats.TotalRaffleEthDeposits = "1.5"
			},
			want: `raffle ETH deposits: invalid non-negative decimal "1.5"`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			record := validRoundRecord()
			test.mutate(&record)
			_, err := mapRoundSummary(record)
			if err == nil || err.Error() != test.want {
				t.Fatalf("error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestMapRoundDetail(t *testing.T) {
	t.Parallel()

	got, err := mapRound(validRoundRecord())
	if err != nil {
		t.Fatalf("mapRound: %v", err)
	}
	if got.Round != 0 || got.Status != Completed || got.Claim.EventLogId != 5018 {
		t.Fatalf("round identity = %+v", got)
	}
	if got.Charity == nil || got.Charity.AmountWei != "90000000000000000" ||
		len(got.Charity.CharityAddresses) != 1 {
		t.Errorf("Charity = %+v", got.Charity)
	}
	if got.Staking == nil || got.Staking.DepositId != 501 ||
		got.Staking.AmountPerTokenWei != "1000000000000000000" {
		t.Errorf("Staking = %+v", got.Staking)
	}
	if got.EnduranceChampion == nil || got.EnduranceChampion.NftTokenId != 5 {
		t.Errorf("EnduranceChampion = %+v", got.EnduranceChampion)
	}
	if got.LastCstBidder == nil || got.LastCstBidder.NftTokenId != 4 {
		t.Errorf("LastCstBidder = %+v", got.LastCstBidder)
	}
	if got.ChronoWarrior == nil || got.ChronoWarrior.EthAmountWei != "80000000000000000" {
		t.Errorf("ChronoWarrior = %+v", got.ChronoWarrior)
	}
	if got.Statistics.ActivatedAt == nil ||
		got.Statistics.RoundStartedAt == nil ||
		got.Statistics.RoundEndedAt == nil ||
		got.Statistics.RoundDurationSeconds == nil ||
		*got.Statistics.RoundDurationSeconds != 900 {
		t.Errorf("Statistics timing = %+v", got.Statistics)
	}
}

func TestMapRoundV3MultiNFTAndChampionDurations(t *testing.T) {
	t.Parallel()
	record := validRoundRecord()
	first := int64(record.MainPrize.NftTokenId) // #nosec G115 -- bounded fixture token ID
	record.MainPrize.NumCSNfts = 3
	record.MainPrize.NftTokenIds = []int64{first, first + 1, first + 2}
	record.RoundStats.EnduranceChampionDuration = 700
	record.RoundStats.ChronoWarriorDuration = 900

	got, err := mapRound(record)
	if err != nil {
		t.Fatalf("mapRound V3: %v", err)
	}
	if got.MainPrize.NumCosmicSignatureNfts == nil ||
		*got.MainPrize.NumCosmicSignatureNfts != 3 ||
		got.MainPrize.NftTokenIds == nil ||
		len(*got.MainPrize.NftTokenIds) != 3 {
		t.Fatalf("V3 main prize = %+v", got.MainPrize)
	}
	if got.Statistics.EnduranceChampionDurationSeconds == nil ||
		*got.Statistics.EnduranceChampionDurationSeconds != 700 ||
		got.Statistics.ChronoWarriorDurationSeconds == nil ||
		*got.Statistics.ChronoWarriorDurationSeconds != 900 {
		t.Fatalf("V3 round statistics = %+v", got.Statistics)
	}
	record.MainPrize.NftTokenIds = nil
	generated, err := mapRound(record)
	if err != nil || generated.MainPrize.NftTokenIds == nil ||
		len(*generated.MainPrize.NftTokenIds) != 3 {
		t.Fatalf("generated V3 token IDs = %+v, %v", generated.MainPrize, err)
	}
}

func TestMapRoundOmitsSentinelsAndLegacyCollections(t *testing.T) {
	t.Parallel()

	record := validRoundRecord()
	record.MainPrize.CstAmount = ""
	record.MainPrize.Seed = "???"
	record.MainPrize.TimeoutTs = 0
	record.CharityDeposit = cgmodel.CGCharityDeposit{}
	record.StakingDeposit.StakingDepositId = -1
	record.EnduranceChampion = cgmodel.CGEnduranceChampionPrize{}
	record.LastCstBidder = cgmodel.CGLastCSTBidderPrize{}
	record.ChronoWarrior = cgmodel.CGChronoWarriorPrize{}
	record.RoundStats.TotalDonatedAmount = ""
	record.RoundStats.ParamWindowStartTime = ""
	record.RoundStats.ActivationTime = 0
	record.RoundStats.ParamWindowDurationSeconds = 0
	record.RoundStats.RoundStartTime = ""
	record.RoundStats.RoundEndTime = ""
	record.RoundStats.RoundDurationSeconds = 0
	record.RaffleNFTWinners = []cgmodel.CGRaffleNFTWinnerRec{{TokenId: 1}}
	record.AllPrizes = []cgmodel.CGPrizeHistory{{RecordType: 1}}

	got, err := mapRound(record)
	if err != nil {
		t.Fatalf("mapRound: %v", err)
	}
	if got.MainPrize.CstAmountWei != nil ||
		got.MainPrize.Seed != nil ||
		got.MainPrize.SecondaryPrizeClaimDeadline != nil ||
		got.Charity != nil ||
		got.Staking != nil ||
		got.EnduranceChampion != nil ||
		got.LastCstBidder != nil ||
		got.ChronoWarrior != nil {
		t.Fatalf("sentinel fields were not omitted: %+v", got)
	}
	if got.Statistics.DonationAmountWei != "0" {
		t.Errorf("empty aggregate amount = %q, want 0", got.Statistics.DonationAmountWei)
	}

	encoded, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	for _, legacy := range []string{"allPrizes", "raffleNFTWinners", "stakingNFTWinners", "raffleETHDeposits", "winnerAid", "AmountEth"} {
		if strings.Contains(string(encoded), legacy) {
			t.Errorf("lean v2 round leaked legacy field %q: %s", legacy, encoded)
		}
	}
}

func TestMapRoundRejectsInvalidStoreData(t *testing.T) {
	t.Parallel()

	tests := map[string]func(*cgmodel.CGRoundRec){
		"round overflow":      func(r *cgmodel.CGRoundRec) { r.RoundNum = math.MaxUint64 },
		"claim identity":      func(r *cgmodel.CGRoundRec) { r.ClaimPrizeTx.Tx.EvtLogId = 0 },
		"claim hash":          func(r *cgmodel.CGRoundRec) { r.ClaimPrizeTx.Tx.TxHash = "bad" },
		"claim timestamp":     func(r *cgmodel.CGRoundRec) { r.ClaimPrizeTx.Tx.DateTime = "bad" },
		"winner address":      func(r *cgmodel.CGRoundRec) { r.MainPrize.WinnerAddr = "bad" },
		"main eth amount":     func(r *cgmodel.CGRoundRec) { r.MainPrize.EthAmount = "-1" },
		"main cst amount":     func(r *cgmodel.CGRoundRec) { r.MainPrize.CstAmount = "wat" },
		"main token overflow": func(r *cgmodel.CGRoundRec) { r.MainPrize.NftTokenId = math.MaxUint64 },
		"main count negative": func(r *cgmodel.CGRoundRec) { r.MainPrize.NumCSNfts = -1 },
		"main range overflow": func(r *cgmodel.CGRoundRec) {
			r.MainPrize.NftTokenId = math.MaxInt64
			r.MainPrize.NumCSNfts = 2
		},
		"main count mismatch": func(r *cgmodel.CGRoundRec) {
			r.MainPrize.NumCSNfts = 3
			r.MainPrize.NftTokenIds = []int64{1, 2}
		},
		"main IDs nonsequential": func(r *cgmodel.CGRoundRec) {
			r.MainPrize.NumCSNfts = 3
			firstTokenID := int64(r.MainPrize.NftTokenId) // #nosec G115 -- bounded fixture token ID
			r.MainPrize.NftTokenIds = []int64{firstTokenID, firstTokenID + 2, firstTokenID + 3}
		},
		"main deadline":        func(r *cgmodel.CGRoundRec) { r.MainPrize.TimeoutTs = -1 },
		"negative aggregate":   func(r *cgmodel.CGRoundRec) { r.RoundStats.TotalBids = -1 },
		"raffle amount":        func(r *cgmodel.CGRoundRec) { r.RoundStats.TotalRaffleEthDeposits = "1.2" },
		"donation amount":      func(r *cgmodel.CGRoundRec) { r.RoundStats.TotalDonatedAmount = "-1" },
		"timing timestamp":     func(r *cgmodel.CGRoundRec) { r.RoundStats.RoundStartTime = "bad" },
		"timing duration":      func(r *cgmodel.CGRoundRec) { r.RoundStats.RoundDurationSeconds = -1 },
		"champion duration":    func(r *cgmodel.CGRoundRec) { r.RoundStats.EnduranceChampionDuration = -1 },
		"charity address":      func(r *cgmodel.CGRoundRec) { r.CharityDeposit.CharityAddress = "bad" },
		"staking identity":     func(r *cgmodel.CGRoundRec) { r.StakingDeposit.StakingDepositId = -2 },
		"staking amount":       func(r *cgmodel.CGRoundRec) { r.StakingDeposit.StakingDepositAmount = "" },
		"endurance address":    func(r *cgmodel.CGRoundRec) { r.EnduranceChampion.WinnerAddr = "bad" },
		"last cst amount":      func(r *cgmodel.CGRoundRec) { r.LastCstBidder.CstAmount = "" },
		"chrono token":         func(r *cgmodel.CGRoundRec) { r.ChronoWarrior.NftTokenId = -1 },
		"chrono eth amount":    func(r *cgmodel.CGRoundRec) { r.ChronoWarrior.EthAmount = "wat" },
		"activation timestamp": func(r *cgmodel.CGRoundRec) { r.RoundStats.ActivationTime = -1 },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validRoundRecord()
			mutate(&record)
			if _, err := mapRound(record); err == nil {
				t.Fatal("mapRound accepted invalid store data")
			}
		})
	}
}

func TestMapCharityAllocationSplitsAddresses(t *testing.T) {
	t.Parallel()

	allocation, err := mapCharityAllocation(cgmodel.CGCharityDeposit{
		CharityAddress: "0x2200000000000000000000000000000000000022, 0x2100000000000000000000000000000000000021",
		CharityAmount:  "10",
	})
	if err != nil {
		t.Fatalf("mapCharityAllocation: %v", err)
	}
	if len(allocation.CharityAddresses) != 2 ||
		allocation.CharityAddresses[0] != "0x2100000000000000000000000000000000000021" {
		t.Fatalf("addresses = %v", allocation.CharityAddresses)
	}
}

func TestOptionalTimestampAcceptsRepositoryFormats(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"RFC3339":                    "2026-01-01T00:41:40-05:00",
		"PostgreSQL hour offset":     "2026-01-01 05:41:40+00",
		"PostgreSQL minute offset":   "2026-01-01 05:41:40+00:00",
		"PostgreSQL fractional time": "2026-01-01 05:41:40.123456+00",
	}
	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := optionalTimestamp(input)
			if err != nil {
				t.Fatalf("optionalTimestamp(%q): %v", input, err)
			}
			if got == nil || got.Location().String() != "UTC" {
				t.Fatalf("timestamp = %v, want UTC value", got)
			}
		})
	}
}

func validRoundRecord() cgmodel.CGRoundRec {
	// #nosec G101 -- deterministic chain fixture values, not credentials.
	return cgmodel.CGRoundRec{
		RoundNum: 0,
		ClaimPrizeTx: cgmodel.CGClaimPrizeTx{Tx: cgmodel.Transaction{
			EvtLogId: 5018,
			BlockNum: 110,
			TxHash:   "0xf000000000000000000000000000000000000000000000000000000000001011",
			DateTime: "2026-01-01T00:16:40-05:00",
		}},
		MainPrize: cgmodel.CGMainPrizeInfo{
			WinnerAddr: "0x2100000000000000000000000000000000000021",
			TimeoutTs:  1768089600,
			EthAmount:  "500000000000000000",
			CstAmount:  "100000000000000000000",
			NftTokenId: 1,
			Seed:       "seed0000000000000000000000000000000000000000000000000000000001",
		},
		CharityDeposit: cgmodel.CGCharityDeposit{
			CharityAddress: "0x6000000000000000000000000000000000000006",
			CharityAmount:  "90000000000000000",
		},
		StakingDeposit: cgmodel.CGStakingDeposit{
			StakingDepositId:       501,
			StakingDepositAmount:   "2000000000000000000",
			StakingPerToken:        "1000000000000000000",
			StakingNumStakedTokens: 2,
		},
		EnduranceChampion: cgmodel.CGEnduranceChampionPrize{
			WinnerAddr: "0x2200000000000000000000000000000000000022",
			NftTokenId: 5,
			CstAmount:  "45000000000000000000",
		},
		LastCstBidder: cgmodel.CGLastCSTBidderPrize{
			WinnerAddr: "0x2300000000000000000000000000000000000023",
			NftTokenId: 4,
			CstAmount:  "40000000000000000000",
		},
		ChronoWarrior: cgmodel.CGChronoWarriorPrize{
			WinnerAddr: "0x2100000000000000000000000000000000000021",
			EthAmount:  "80000000000000000",
			CstAmount:  "35000000000000000000",
			NftTokenId: 6,
		},
		RoundStats: cgmodel.CGRoundStats{
			TotalBids:                  4,
			TotalDonatedNFTs:           1,
			NumERC20Donations:          1,
			TotalRaffleEthDeposits:     "100000000000000000",
			TotalRaffleNFTs:            2,
			TotalDonatedCount:          2,
			TotalDonatedAmount:         "500000000000000000",
			ParamWindowStartTime:       "2025-12-31T23:00:00-05:00",
			ActivationTime:             1767225650,
			ParamWindowDurationSeconds: 60,
			RoundStartTime:             "2026-01-01T00:01:40-05:00",
			RoundEndTime:               "2026-01-01T00:16:40-05:00",
			RoundDurationSeconds:       900,
		},
	}
}
