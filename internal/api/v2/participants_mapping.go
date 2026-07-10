package v2

import (
	"errors"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func mapBidderParticipant(record cgstore.BidderParticipantRecord) (BidderParticipant, error) {
	address, err := participantAddress(record.BidderAid, record.Address)
	if err != nil {
		return BidderParticipant{}, fmt.Errorf("bidder: %w", err)
	}
	if record.BidCount < 1 {
		return BidderParticipant{}, errors.New("bidder has no bids")
	}
	maxBid, err := optionalAmount(record.MaxBidWei)
	if err != nil {
		return BidderParticipant{}, fmt.Errorf("bidder max bid: %w", err)
	}
	return BidderParticipant{
		BidCount:      record.BidCount,
		BidderAddress: address,
		MaxBidWei:     maxBid,
	}, nil
}

func mapWinnerParticipant(record cgstore.WinnerParticipantRecord) (WinnerParticipant, error) {
	address, err := participantAddress(record.WinnerAid, record.Address)
	if err != nil {
		return WinnerParticipant{}, fmt.Errorf("winner: %w", err)
	}
	if err := nonNegativeParticipantCounts(map[string]int64{
		"winner prize":         record.PrizeCount,
		"winner CST prize":     record.CSTPrizeCount,
		"winner NFT prize":     record.NFTPrizeCount,
		"winner unclaimed NFT": record.UnclaimedNFTCount,
	}); err != nil {
		return WinnerParticipant{}, err
	}
	if record.PrizeCount < 1 {
		return WinnerParticipant{}, errors.New("winner has no prizes")
	}
	maxWin, err := requiredAmount(record.MaxMainPrizeETHWei)
	if err != nil {
		return WinnerParticipant{}, fmt.Errorf("winner max win: %w", err)
	}
	totalWon, err := requiredAmount(record.TotalETHWonWei)
	if err != nil {
		return WinnerParticipant{}, fmt.Errorf("winner total won: %w", err)
	}
	totalSpent, err := requiredAmount(record.TotalETHSpentWei)
	if err != nil {
		return WinnerParticipant{}, fmt.Errorf("winner total spent: %w", err)
	}
	return WinnerParticipant{
		CstPrizeCount:      record.CSTPrizeCount,
		MaxMainPrizeEthWei: maxWin,
		NftPrizeCount:      record.NFTPrizeCount,
		PrizeCount:         record.PrizeCount,
		TotalEthSpentWei:   totalSpent,
		TotalEthWonWei:     totalWon,
		UnclaimedNftCount:  record.UnclaimedNFTCount,
		WinnerAddress:      address,
	}, nil
}

func mapDonorParticipant(record cgstore.DonorParticipantRecord) (DonorParticipant, error) {
	address, err := participantAddress(record.DonorAid, record.Address)
	if err != nil {
		return DonorParticipant{}, fmt.Errorf("donor: %w", err)
	}
	if record.DonationCount < 1 {
		return DonorParticipant{}, errors.New("donor has no donations")
	}
	total, err := requiredAmount(record.TotalDonatedWei)
	if err != nil {
		return DonorParticipant{}, fmt.Errorf("donor total donated: %w", err)
	}
	return DonorParticipant{
		DonationCount:   record.DonationCount,
		DonorAddress:    address,
		TotalDonatedWei: total,
	}, nil
}

func mapCSTStakerParticipant(record cgstore.CSTStakerParticipantRecord) (CstStakerParticipant, error) {
	address, err := participantAddress(record.StakerAid, record.Address)
	if err != nil {
		return CstStakerParticipant{}, fmt.Errorf("CST staker: %w", err)
	}
	staking, err := mapCSTParticipantStats(
		record.StakedTokenCount,
		record.StakeActionCount,
		record.UnstakeActionCount,
		record.TotalRewardWei,
		record.UnclaimedRewardWei,
	)
	if err != nil {
		return CstStakerParticipant{}, fmt.Errorf("CST staker: %w", err)
	}
	if record.StakeActionCount < 1 {
		return CstStakerParticipant{}, errors.New("CST staker has no stake actions")
	}
	return CstStakerParticipant{StakerAddress: address, Staking: staking}, nil
}

func mapRandomWalkStakerParticipant(
	record cgstore.RandomWalkStakerParticipantRecord,
) (RandomWalkStakerParticipant, error) {
	address, err := participantAddress(record.StakerAid, record.Address)
	if err != nil {
		return RandomWalkStakerParticipant{}, fmt.Errorf("RandomWalk staker: %w", err)
	}
	if record.StakeActionCount < 1 {
		return RandomWalkStakerParticipant{}, errors.New("RandomWalk staker has no stake actions")
	}
	staking, err := mapRandomWalkParticipantStats(
		record.StakedTokenCount,
		record.StakeActionCount,
		record.UnstakeActionCount,
		record.MintedTokenCount,
	)
	if err != nil {
		return RandomWalkStakerParticipant{}, fmt.Errorf("RandomWalk staker: %w", err)
	}
	return RandomWalkStakerParticipant{StakerAddress: address, Staking: staking}, nil
}

func mapDualStakerParticipant(
	record cgstore.DualStakerParticipantRecord,
) (DualStakerParticipant, error) {
	address, err := participantAddress(record.StakerAid, record.Address)
	if err != nil {
		return DualStakerParticipant{}, fmt.Errorf("dual staker: %w", err)
	}
	if record.CSTStakedTokenCount < 1 ||
		record.RandomWalkStakedTokenCount < 1 ||
		record.TotalStakedTokenCount < 2 ||
		record.TotalStakedTokenCount != record.CSTStakedTokenCount+record.RandomWalkStakedTokenCount {
		return DualStakerParticipant{}, errors.New("invalid dual-staker total")
	}
	cst, err := mapCSTParticipantStats(
		record.CSTStakedTokenCount,
		record.CSTStakeActionCount,
		record.CSTUnstakeActionCount,
		record.CSTTotalRewardWei,
		record.CSTUnclaimedRewardWei,
	)
	if err != nil {
		return DualStakerParticipant{}, fmt.Errorf("dual staker CST statistics: %w", err)
	}
	randomWalk, err := mapRandomWalkParticipantStats(
		record.RandomWalkStakedTokenCount,
		record.RandomWalkStakeActionCount,
		record.RandomWalkUnstakeActionCount,
		record.RandomWalkMintedTokenCount,
	)
	if err != nil {
		return DualStakerParticipant{}, fmt.Errorf("dual staker RandomWalk statistics: %w", err)
	}
	return DualStakerParticipant{
		Cst:                   cst,
		RandomWalk:            randomWalk,
		StakerAddress:         address,
		TotalStakedTokenCount: record.TotalStakedTokenCount,
	}, nil
}

func mapCSTParticipantStats(
	staked, stakes, unstakes int64,
	totalReward, unclaimedReward string,
) (CstStakingParticipantStats, error) {
	if err := nonNegativeParticipantCounts(map[string]int64{
		"CST staked token":   staked,
		"CST stake action":   stakes,
		"CST unstake action": unstakes,
	}); err != nil {
		return CstStakingParticipantStats{}, err
	}
	total, err := requiredAmount(totalReward)
	if err != nil {
		return CstStakingParticipantStats{}, fmt.Errorf("total reward: %w", err)
	}
	unclaimed, err := requiredAmount(unclaimedReward)
	if err != nil {
		return CstStakingParticipantStats{}, fmt.Errorf("unclaimed reward: %w", err)
	}
	return CstStakingParticipantStats{
		StakeActionCount:   stakes,
		StakedTokenCount:   staked,
		TotalRewardWei:     total,
		UnclaimedRewardWei: unclaimed,
		UnstakeActionCount: unstakes,
	}, nil
}

func mapRandomWalkParticipantStats(
	staked, stakes, unstakes, minted int64,
) (RandomWalkStakingParticipantStats, error) {
	if err := nonNegativeParticipantCounts(map[string]int64{
		"RandomWalk staked token":   staked,
		"RandomWalk stake action":   stakes,
		"RandomWalk unstake action": unstakes,
		"RandomWalk minted token":   minted,
	}); err != nil {
		return RandomWalkStakingParticipantStats{}, err
	}
	return RandomWalkStakingParticipantStats{
		MintedTokenCount:   minted,
		StakeActionCount:   stakes,
		StakedTokenCount:   staked,
		UnstakeActionCount: unstakes,
	}, nil
}

func participantAddress(id int64, address string) (string, error) {
	if id < 1 {
		return "", errors.New("invalid internal address ID")
	}
	if !ethcommon.IsHexAddress(address) {
		return "", errors.New("invalid participant address")
	}
	return ethcommon.HexToAddress(address).Hex(), nil
}

func nonNegativeParticipantCounts(counts map[string]int64) error {
	for name, value := range counts {
		if value < 0 {
			return fmt.Errorf("negative %s count", name)
		}
	}
	return nil
}
