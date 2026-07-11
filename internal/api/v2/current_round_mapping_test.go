package v2

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

func TestMapCurrentRound(t *testing.T) {
	t.Parallel()

	live, err := normalizeCurrentRoundSnapshot(validCurrentRoundSnapshot())
	if err != nil {
		t.Fatalf("normalizeCurrentRoundSnapshot: %v", err)
	}
	stats := validCurrentRoundStats()
	stats.TotalBids = 99 // The direct count is authoritative for the live resource.

	got, err := mapCurrentRound(live, stats, 3)
	if err != nil {
		t.Fatalf("mapCurrentRound: %v", err)
	}
	if got.Round != 3 || got.Status != Open || got.Statistics.TotalBids != 3 {
		t.Fatalf("current round identity/statistics = %+v", got)
	}
	if got.NextEthBidPriceWei != "1010000000000000" ||
		got.MainPrizeAmountWei != "2500000000000000000" ||
		got.RafflePrizeAmountWei != "50000000000000000" ||
		got.StakingRewardAmountWei != "100000000000000000" ||
		got.MainPrizeTimeIncrementMicroseconds != "3600000000" {
		t.Errorf("current round amounts = %+v", got)
	}
	if got.SecondsUntilMainPrize != 3600 {
		t.Errorf("SecondsUntilMainPrize = %d, want 3600", got.SecondsUntilMainPrize)
	}
	if got.LastBidderAddress == nil ||
		*got.LastBidderAddress != "0x2100000000000000000000000000000000000021" {
		t.Errorf("LastBidderAddress = %v", got.LastBidderAddress)
	}
	if got.Statistics.ActivatedAt == nil ||
		got.Statistics.RoundStartedAt == nil ||
		got.Statistics.ActivatedAt.Location().String() != "UTC" ||
		got.Statistics.RoundStartedAt.Location().String() != "UTC" {
		t.Errorf("current round timing = %+v", got.Statistics)
	}
}

func TestMapCurrentRoundOmitsLastBidderBeforeFirstBid(t *testing.T) {
	t.Parallel()

	snapshot := validCurrentRoundSnapshot()
	snapshot.LastBidder = ethcommon.Address{}
	live, err := normalizeCurrentRoundSnapshot(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	stats := validCurrentRoundStats()
	stats.TotalBids = 0

	got, err := mapCurrentRound(live, stats, 0)
	if err != nil {
		t.Fatalf("mapCurrentRound: %v", err)
	}
	if got.LastBidderAddress != nil {
		t.Fatalf("LastBidderAddress = %q, want omitted", *got.LastBidderAddress)
	}

	encoded, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"lastBidderAddress", "BidPriceEth", "PrizeAmountEth"} {
		if strings.Contains(string(encoded), forbidden) {
			t.Errorf("current round JSON leaked %q: %s", forbidden, encoded)
		}
	}
}

func TestNormalizeCurrentRoundSnapshotRejectsUnavailableState(t *testing.T) {
	t.Parallel()

	tests := map[string]func(*contractstate.Snapshot){
		"round sentinel":        func(s *contractstate.Snapshot) { s.RoundNum = -1 },
		"duration sentinel":     func(s *contractstate.Snapshot) { s.PrizeClaimTimestamp = -1 },
		"empty bid price":       func(s *contractstate.Snapshot) { s.BlockPinnedBidPrice = "" },
		"failed bid price":      func(s *contractstate.Snapshot) { s.BlockPinnedBidPrice = "error" },
		"empty main prize":      func(s *contractstate.Snapshot) { s.PrizeAmount = "" },
		"failed main prize":     func(s *contractstate.Snapshot) { s.PrizeAmount = "error" },
		"empty raffle prize":    func(s *contractstate.Snapshot) { s.RaffleAmount = "" },
		"failed raffle prize":   func(s *contractstate.Snapshot) { s.RaffleAmount = "error" },
		"empty staking reward":  func(s *contractstate.Snapshot) { s.StakingAmount = "" },
		"failed staking reward": func(s *contractstate.Snapshot) { s.StakingAmount = "error" },
		"empty time increment":  func(s *contractstate.Snapshot) { s.MainPrizeTimeIncrement = "" },
		"failed time increment": func(s *contractstate.Snapshot) { s.MainPrizeTimeIncrement = "error" },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			snapshot := validCurrentRoundSnapshot()
			mutate(&snapshot)
			_, err := normalizeCurrentRoundSnapshot(snapshot)
			if !errors.Is(err, errCurrentRoundUnavailable) {
				t.Fatalf("error = %v, want errCurrentRoundUnavailable", err)
			}
		})
	}
}

func TestNormalizeCurrentRoundSnapshotRejectsMalformedData(t *testing.T) {
	t.Parallel()

	tests := map[string]func(*contractstate.Snapshot){
		"negative bid price":     func(s *contractstate.Snapshot) { s.BlockPinnedBidPrice = "-1" },
		"fractional main prize":  func(s *contractstate.Snapshot) { s.PrizeAmount = "1.5" },
		"invalid raffle prize":   func(s *contractstate.Snapshot) { s.RaffleAmount = "wat" },
		"negative staking prize": func(s *contractstate.Snapshot) { s.StakingAmount = "-10" },
		"invalid time increment": func(s *contractstate.Snapshot) { s.MainPrizeTimeIncrement = "1e6" },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			snapshot := validCurrentRoundSnapshot()
			mutate(&snapshot)
			_, err := normalizeCurrentRoundSnapshot(snapshot)
			if err == nil {
				t.Fatal("normalizeCurrentRoundSnapshot accepted malformed data")
			}
			if errors.Is(err, errCurrentRoundUnavailable) {
				t.Fatalf("malformed data was classified as unavailable: %v", err)
			}
		})
	}
}

func TestMapCurrentRoundRejectsInconsistentData(t *testing.T) {
	t.Parallel()

	validLive, err := normalizeCurrentRoundSnapshot(validCurrentRoundSnapshot())
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		mutateLive  func(*currentRoundSnapshot)
		mutateStats func(*cgprimitives.CGRoundStats)
		bidCount    int64
		unavailable bool
	}{
		"round mismatch": {
			mutateStats: func(s *cgprimitives.CGRoundStats) { s.RoundNum = 2 },
			bidCount:    3,
		},
		"negative direct bid count": {
			bidCount: -1,
		},
		"negative stored count": {
			mutateStats: func(s *cgprimitives.CGRoundStats) { s.TotalBids = -1 },
			bidCount:    3,
		},
		"malformed aggregate amount": {
			mutateStats: func(s *cgprimitives.CGRoundStats) { s.TotalDonatedAmount = "1.5" },
			bidCount:    3,
		},
		"malformed round timestamp": {
			mutateStats: func(s *cgprimitives.CGRoundStats) { s.RoundStartTime = "yesterday" },
			bidCount:    3,
		},
		"missing bidder after bids": {
			mutateLive:  func(l *currentRoundSnapshot) { l.lastBidder = ethcommon.Address{} },
			bidCount:    3,
			unavailable: true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			live := validLive
			stats := validCurrentRoundStats()
			if tc.mutateLive != nil {
				tc.mutateLive(&live)
			}
			if tc.mutateStats != nil {
				tc.mutateStats(&stats)
			}
			_, err := mapCurrentRound(live, stats, tc.bidCount)
			if err == nil {
				t.Fatal("mapCurrentRound accepted inconsistent data")
			}
			if errors.Is(err, errCurrentRoundUnavailable) != tc.unavailable {
				t.Fatalf("unavailable classification = %v, error=%v", errors.Is(err, errCurrentRoundUnavailable), err)
			}
		})
	}
}

func validCurrentRoundSnapshot() contractstate.Snapshot {
	return contractstate.Snapshot{
		BidPrice:                 "001010000000000000",
		BlockPinnedBidPrice:      "001010000000000000",
		PrizeClaimTimestamp:      3600,
		PrizeAmount:              "02500000000000000000",
		RaffleAmount:             "050000000000000000",
		StakingAmount:            "0100000000000000000",
		RoundNum:                 3,
		MainPrizeTimeIncrement:   "03600000000",
		LastBidder:               ethcommon.HexToAddress("0x2100000000000000000000000000000000000021"),
		InitialSecondsUntilPrize: 2,
	}
}

func validCurrentRoundStats() cgprimitives.CGRoundStats {
	return cgprimitives.CGRoundStats{
		RoundNum:                   3,
		TotalBids:                  3,
		TotalDonatedNFTs:           1,
		NumERC20Donations:          2,
		TotalRaffleEthDeposits:     "50000000000000000",
		TotalRaffleNFTs:            4,
		TotalDonatedCount:          5,
		TotalDonatedAmount:         "60000000000000000",
		ParamWindowStartTime:       "2026-01-01 05:41:40+00",
		ActivationTime:             1767228300,
		ParamWindowDurationSeconds: 200,
		RoundStartTime:             "2026-01-01 05:45:00+00",
	}
}
