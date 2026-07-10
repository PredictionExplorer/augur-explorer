package v2

import (
	"errors"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

var errCurrentRoundUnavailable = errors.New("current round live state is unavailable")

type currentRoundSnapshot struct {
	round                              int64
	nextEthBidPriceWei                 string
	secondsUntilMainPrize              int64
	mainPrizeAmountWei                 string
	rafflePrizeAmountWei               string
	stakingRewardAmountWei             string
	mainPrizeTimeIncrementMicroseconds string
	lastBidder                         ethcommon.Address
}

func normalizeCurrentRoundSnapshot(snapshot contractstate.Snapshot) (currentRoundSnapshot, error) {
	if snapshot.RoundNum < 0 {
		return currentRoundSnapshot{}, fmt.Errorf("round number: %w", errCurrentRoundUnavailable)
	}
	if snapshot.PrizeClaimTimestamp < 0 {
		return currentRoundSnapshot{}, fmt.Errorf("main-prize duration: %w", errCurrentRoundUnavailable)
	}

	nextBidPrice, err := normalizeLiveDecimal("next ETH bid price", snapshot.BidPrice)
	if err != nil {
		return currentRoundSnapshot{}, err
	}
	mainPrizeAmount, err := normalizeLiveDecimal("main-prize amount", snapshot.PrizeAmount)
	if err != nil {
		return currentRoundSnapshot{}, err
	}
	rafflePrizeAmount, err := normalizeLiveDecimal("raffle-prize amount", snapshot.RaffleAmount)
	if err != nil {
		return currentRoundSnapshot{}, err
	}
	stakingRewardAmount, err := normalizeLiveDecimal("staking-reward amount", snapshot.StakingAmount)
	if err != nil {
		return currentRoundSnapshot{}, err
	}
	timeIncrement, err := normalizeLiveDecimal("main-prize time increment", snapshot.MainPrizeTimeIncrement)
	if err != nil {
		return currentRoundSnapshot{}, err
	}

	return currentRoundSnapshot{
		round:                              snapshot.RoundNum,
		nextEthBidPriceWei:                 nextBidPrice,
		secondsUntilMainPrize:              snapshot.PrizeClaimTimestamp,
		mainPrizeAmountWei:                 mainPrizeAmount,
		rafflePrizeAmountWei:               rafflePrizeAmount,
		stakingRewardAmountWei:             stakingRewardAmount,
		mainPrizeTimeIncrementMicroseconds: timeIncrement,
		lastBidder:                         snapshot.LastBidder,
	}, nil
}

func normalizeLiveDecimal(name, value string) (string, error) {
	if value == "" || value == "error" {
		return "", fmt.Errorf("%s: %w", name, errCurrentRoundUnavailable)
	}
	normalized, err := requiredAmount(value)
	if err != nil {
		return "", fmt.Errorf("%s: %w", name, err)
	}
	return normalized, nil
}

func mapCurrentRound(
	live currentRoundSnapshot,
	stats cgprimitives.CGRoundStats,
	bidCount int64,
) (CosmicGameCurrentRound, error) {
	if stats.RoundNum != live.round {
		return CosmicGameCurrentRound{}, errors.New("repository returned statistics for another round")
	}
	if bidCount < 0 {
		return CosmicGameCurrentRound{}, errors.New("repository returned a negative bid count")
	}
	if err := validateRoundCounts(stats); err != nil {
		return CosmicGameCurrentRound{}, err
	}

	stats.TotalBids = bidCount
	mappedStats, err := mapRoundStatistics(stats)
	if err != nil {
		return CosmicGameCurrentRound{}, err
	}

	result := CosmicGameCurrentRound{
		MainPrizeAmountWei:                 live.mainPrizeAmountWei,
		MainPrizeTimeIncrementMicroseconds: live.mainPrizeTimeIncrementMicroseconds,
		NextEthBidPriceWei:                 live.nextEthBidPriceWei,
		RafflePrizeAmountWei:               live.rafflePrizeAmountWei,
		Round:                              live.round,
		SecondsUntilMainPrize:              live.secondsUntilMainPrize,
		StakingRewardAmountWei:             live.stakingRewardAmountWei,
		Statistics:                         mappedStats,
		Status:                             Open,
	}

	if live.lastBidder != (ethcommon.Address{}) {
		address := live.lastBidder.Hex()
		result.LastBidderAddress = &address
	} else if bidCount > 0 {
		return CosmicGameCurrentRound{}, fmt.Errorf("last bidder: %w", errCurrentRoundUnavailable)
	}

	return result, nil
}
