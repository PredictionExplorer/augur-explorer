package v2

import (
	"errors"
	"fmt"
	"math"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

var errCachedLiveUnavailable = errors.New("cached live resource is unavailable")

func mapContractAddressRegistry(record cgmodel.ContractAddrs) (ContractAddressRegistry, error) {
	values := []struct {
		name  string
		value string
	}{
		{"CosmicGame", record.CosmicGameAddr},
		{"CosmicSignature", record.CosmicSignatureAddr},
		{"CosmicToken", record.CosmicTokenAddr},
		{"CosmicDAO", record.CosmicDaoAddr},
		{"charity wallet", record.CharityWalletAddr},
		{"prizes wallet", record.PrizesWalletAddr},
		{"RandomWalk", record.RandomWalkAddr},
		{"CST staking wallet", record.StakingWalletCSTAddr},
		{"RandomWalk staking wallet", record.StakingWalletRWalkAddr},
		{"marketing wallet", record.MarketingWalletAddr},
		{"marketplace", record.MarketplaceAddr},
		{"implementation", record.ImplementationAddr},
	}
	normalized := make([]string, len(values))
	for i := range values {
		address, err := canonicalNonZeroAddress(values[i].name, values[i].value)
		if err != nil {
			return ContractAddressRegistry{}, err
		}
		normalized[i] = address
	}
	return ContractAddressRegistry{
		CosmicGameAddress:              normalized[0],
		CosmicSignatureAddress:         normalized[1],
		CosmicTokenAddress:             normalized[2],
		CosmicDaoAddress:               normalized[3],
		CharityWalletAddress:           normalized[4],
		PrizesWalletAddress:            normalized[5],
		RandomWalkAddress:              normalized[6],
		CstStakingWalletAddress:        normalized[7],
		RandomWalkStakingWalletAddress: normalized[8],
		MarketingWalletAddress:         normalized[9],
		MarketplaceAddress:             normalized[10],
		ImplementationAddress:          normalized[11],
	}, nil
}

func mapContractConfiguration(snapshot contractstate.Snapshot) (ContractConfiguration, error) {
	if !snapshot.ConfigurationReady {
		return ContractConfiguration{}, errCachedLiveUnavailable
	}
	charity, err := canonicalSnapshotAddress("charity", snapshot.CharityAddr)
	if err != nil {
		return ContractConfiguration{}, err
	}
	treasurer, err := canonicalSnapshotAddress("treasurer", snapshot.TreasurerAddr)
	if err != nil {
		return ContractConfiguration{}, err
	}
	priceIncrease, err := requiredAmount(snapshot.PriceIncrease)
	if err != nil {
		return ContractConfiguration{}, fmt.Errorf("ETH bid-price divisor: %w", err)
	}
	timeIncrease, err := requiredAmount(snapshot.TimeIncrease)
	if err != nil {
		return ContractConfiguration{}, fmt.Errorf("time-increment divisor: %w", err)
	}
	if !validPercent(snapshot.CharityPercentage) ||
		!validPercent(snapshot.PrizePercentage) ||
		!validPercent(snapshot.RafflePercentage) ||
		!validPercent(snapshot.ChronoPercentage) ||
		!validPercent(snapshot.StakingPercentage) ||
		snapshot.RaffleEthWinnersBidding < 0 ||
		snapshot.RaffleNFTWinnersBidding < 0 ||
		snapshot.RaffleNFTWinnersStakingRWalk < 0 ||
		snapshot.InitialSecondsUntilPrize <= 0 ||
		snapshot.TimeoutClaimPrize < 0 ||
		snapshot.RoundStartAuctionLength <= 0 {
		return ContractConfiguration{}, errors.New("cached contract configuration is inconsistent")
	}
	result := ContractConfiguration{
		CharityAddress:                        charity,
		CharityDonationPercentage:             snapshot.CharityPercentage,
		ChronoWarriorPercentage:               snapshot.ChronoPercentage,
		CstRoundStartAuctionValue:             snapshot.RoundStartAuctionLength,
		EthBidPriceIncreaseDivisor:            priceIncrease,
		InitialDurationUntilMainPrizeDivisor:  snapshot.InitialSecondsUntilPrize,
		MainPrizePercentage:                   snapshot.PrizePercentage,
		MainPrizeTimeIncrementIncreaseDivisor: timeIncrease,
		RaffleEthWinnerCount:                  snapshot.RaffleEthWinnersBidding,
		RaffleNftBidderWinnerCount:            snapshot.RaffleNFTWinnersBidding,
		RaffleNftRandomWalkStakerWinnerCount:  snapshot.RaffleNFTWinnersStakingRWalk,
		RafflePercentage:                      snapshot.RafflePercentage,
		StakingPercentage:                     snapshot.StakingPercentage,
		TimeoutMainPrizeClaimSeconds:          snapshot.TimeoutClaimPrize,
		TreasurerAddress:                      treasurer,
	}
	switch snapshot.MechanicsVersion {
	case 1:
		if snapshot.CSTAuctionDurationChangeDivisor != -1 {
			return ContractConfiguration{}, errors.New("v1 mechanics has a v2 auction-change divisor")
		}
		result.MechanicsVersion = ContractMechanicsVersionV1
		result.CstBidRewardMode = CstBidRewardFixed
		result.CstRoundStartAuctionMode = CstRoundStartAuctionModeDivisor
		fixedReward, err := requiredAmount(snapshot.FixedCSTBidReward)
		if err != nil {
			return ContractConfiguration{}, fmt.Errorf("fixed CST bid reward: %w", err)
		}
		result.FixedCstBidRewardWei = &fixedReward
	case 2:
		if snapshot.CSTAuctionDurationChangeDivisor <= 0 {
			return ContractConfiguration{}, errors.New("v2 mechanics lacks an auction-change divisor")
		}
		divisor := snapshot.CSTAuctionDurationChangeDivisor
		result.MechanicsVersion = ContractMechanicsVersionV2
		result.CstBidRewardMode = CstBidRewardDynamic
		result.CstRoundStartAuctionMode = CstRoundStartAuctionModeDurationSeconds
		result.CstDutchAuctionDurationChangeDivisor = &divisor
		multiplier, err := requiredAmount(snapshot.BidCSTRewardMultiplier)
		if err != nil {
			return ContractConfiguration{}, fmt.Errorf("CST bid reward multiplier: %w", err)
		}
		result.CstBidRewardMultiplier = &multiplier
	case 3:
		if snapshot.CSTAuctionDurationChangeDivisor <= 0 {
			return ContractConfiguration{}, errors.New("v3 mechanics lacks an auction-change divisor")
		}
		result.MechanicsVersion = ContractMechanicsVersionV3
		result.CstBidRewardMode = CstBidRewardDynamic
		result.CstRoundStartAuctionMode = CstRoundStartAuctionModeDurationSeconds
		divisor := snapshot.CSTAuctionDurationChangeDivisor
		result.CstDutchAuctionDurationChangeDivisor = &divisor
		multiplier, err := requiredAmount(snapshot.BidCSTRewardMultiplier)
		if err != nil {
			return ContractConfiguration{}, fmt.Errorf("CST bid reward multiplier: %w", err)
		}
		result.CstBidRewardMultiplier = &multiplier

		lateBidDivisor, err := requiredAmount(snapshot.V3.RoundLateBidDurationDivisor)
		if err != nil {
			return ContractConfiguration{}, fmt.Errorf("late-bid duration divisor: %w", err)
		}
		premiumBase, err := requiredAmount(snapshot.V3.RoundLateBidPricePremiumAmountBaseMultiplier)
		if err != nil {
			return ContractConfiguration{}, fmt.Errorf("late-bid premium base multiplier: %w", err)
		}
		auctionFloor, err := requiredAmount(snapshot.V3.CstDutchAuctionBeginningBidPriceMinLimit)
		if err != nil {
			return ContractConfiguration{}, fmt.Errorf("CST auction beginning-price floor: %w", err)
		}
		rewardPerIncrement, err := requiredAmount(snapshot.V3.BidCstRewardAmountPerMainPrizeTimeIncrement)
		if err != nil {
			return ContractConfiguration{}, fmt.Errorf("CST reward per main-prize increment: %w", err)
		}
		if snapshot.V3.RoundLateBidDurationSeconds <= 0 ||
			snapshot.V3.RoundLateBidPricePremiumAmountExponent < 0 ||
			!validPercent(snapshot.V3.LastBidderBidCstRewardAmountPercentage) ||
			snapshot.V3.MainPrizeNumCosmicSignatureNfts <= 0 {
			return ContractConfiguration{}, errors.New("cached v3 contract configuration is inconsistent")
		}
		result.RoundLateBidDurationDivisor = &lateBidDivisor
		result.RoundLateBidDurationSeconds = &snapshot.V3.RoundLateBidDurationSeconds
		result.RoundLateBidPricePremiumAmountBaseMultiplier = &premiumBase
		result.RoundLateBidPricePremiumAmountExponent = &snapshot.V3.RoundLateBidPricePremiumAmountExponent
		result.LastBidderBidCstRewardAmountPercentage = &snapshot.V3.LastBidderBidCstRewardAmountPercentage
		result.MainPrizeNumCosmicSignatureNfts = &snapshot.V3.MainPrizeNumCosmicSignatureNfts
		result.CstDutchAuctionBeginningBidPriceMinLimitWei = &auctionFloor
		result.BidCstRewardAmountPerMainPrizeTimeIncrementWei = &rewardPerIncrement
	default:
		return ContractConfiguration{}, errors.New("unknown contract mechanics version")
	}
	return result, nil
}

func mapContractBalances(snapshot contractstate.Snapshot) (ContractBalances, error) {
	if !snapshot.BalancesReady {
		return ContractBalances{}, errCachedLiveUnavailable
	}
	if snapshot.BalanceCharityAddr != snapshot.CharityAddr {
		return ContractBalances{}, errors.New("charity balance belongs to another address")
	}
	charity, err := canonicalSnapshotAddress("charity", snapshot.CharityAddr)
	if err != nil {
		return ContractBalances{}, err
	}
	gameBalance, err := requiredAmount(snapshot.CosmicGameBalance)
	if err != nil {
		return ContractBalances{}, fmt.Errorf("CosmicGame balance: %w", err)
	}
	charityBalance, err := requiredAmount(snapshot.CharityBalance)
	if err != nil {
		return ContractBalances{}, fmt.Errorf("charity balance: %w", err)
	}
	return ContractBalances{
		CharityAddress:       charity,
		CharityBalanceWei:    charityBalance,
		CosmicGameBalanceWei: gameBalance,
	}, nil
}

func mapCurrentBidPrices(snapshot contractstate.Snapshot) (CurrentBidPrices, error) {
	if !snapshot.BidPricesReady {
		return CurrentBidPrices{}, errCachedLiveUnavailable
	}
	ethPrice, err := requiredAmount(snapshot.BlockPinnedBidPrice)
	if err != nil {
		return CurrentBidPrices{}, fmt.Errorf("next ETH bid price: %w", err)
	}
	cstPrice, err := requiredAmount(snapshot.NextCSTBidPrice)
	if err != nil {
		return CurrentBidPrices{}, fmt.Errorf("next CST bid price: %w", err)
	}
	cstReward, err := requiredAmount(snapshot.NextCSTBidReward)
	if err != nil {
		return CurrentBidPrices{}, fmt.Errorf("next CST bid reward: %w", err)
	}
	if !validAuctionProgress(snapshot.ETHAuctionDuration, snapshot.ETHAuctionElapsed) ||
		!validAuctionProgress(snapshot.CSTAuctionDuration, snapshot.CSTAuctionElapsed) {
		return CurrentBidPrices{}, errors.New("cached Dutch-auction progress is inconsistent")
	}
	return CurrentBidPrices{
		CstAuctionDurationSeconds: snapshot.CSTAuctionDuration,
		CstAuctionElapsedSeconds:  snapshot.CSTAuctionElapsed,
		EthAuctionDurationSeconds: snapshot.ETHAuctionDuration,
		EthAuctionElapsedSeconds:  snapshot.ETHAuctionElapsed,
		NextCstBidPriceWei:        cstPrice,
		NextCstBidRewardWei:       cstReward,
		NextEthBidPriceWei:        ethPrice,
	}, nil
}

func mapCurrentSpecialWinners(snapshot contractstate.Snapshot) (CurrentSpecialWinners, error) {
	if !snapshot.SpecialWinnersReady {
		return CurrentSpecialWinners{}, errCachedLiveUnavailable
	}
	record := snapshot.SpecialWinners
	if record.Err != nil || record.RoundNum < 0 ||
		record.SourceBlockNumber > math.MaxInt64 || record.SourceBlockTimeStamp < 0 {
		return CurrentSpecialWinners{}, errors.New("cached special-winner source is inconsistent")
	}
	result := CurrentSpecialWinners{
		Round:                record.RoundNum,
		SourceBlockNumber:    int64(record.SourceBlockNumber), // #nosec G115 -- checked above
		SourceBlockTimestamp: record.SourceBlockTimeStamp,
	}
	if record.EnduranceChampionAddress != (ethcommon.Address{}).Hex() {
		address, err := canonicalNonZeroAddress("endurance champion", record.EnduranceChampionAddress)
		if err != nil {
			return CurrentSpecialWinners{}, err
		}
		if record.EnduranceChampionDuration < 0 ||
			record.EnduranceChampionStartTimeStamp < 0 ||
			record.PrevEnduranceChampionDuration < 0 {
			return CurrentSpecialWinners{}, errors.New("invalid endurance-champion timing")
		}
		result.EnduranceChampion = &EnduranceChampionStanding{
			Address:                 address,
			DurationSeconds:         record.EnduranceChampionDuration,
			PreviousDurationSeconds: record.PrevEnduranceChampionDuration,
			StartedAt:               record.EnduranceChampionStartTimeStamp,
		}
	}
	if record.ChronoWarriorAddress != (ethcommon.Address{}).Hex() {
		address, err := canonicalNonZeroAddress("chrono warrior", record.ChronoWarriorAddress)
		if err != nil {
			return CurrentSpecialWinners{}, err
		}
		if record.ChronoWarriorDuration < 0 {
			return CurrentSpecialWinners{}, errors.New("invalid chrono-warrior duration")
		}
		result.ChronoWarrior = &ChronoWarriorStanding{
			Address:         address,
			DurationSeconds: record.ChronoWarriorDuration,
			IsLive:          record.ChronoWarriorIsLive,
		}
	}
	if record.LastBidderAddress != (ethcommon.Address{}).Hex() {
		address, err := canonicalNonZeroAddress("last bidder", record.LastBidderAddress)
		if err != nil {
			return CurrentSpecialWinners{}, err
		}
		if record.LastBidderLastBidTime < 0 {
			return CurrentSpecialWinners{}, errors.New("invalid last-bidder timestamp")
		}
		result.LastBidder = &LastBidderStanding{
			Address:   address,
			LastBidAt: record.LastBidderLastBidTime,
		}
	}
	lastCSTZero := record.LastCstBidderAddress == (ethcommon.Address{}).Hex()
	if lastCSTZero {
		if record.HasLastCstBidderLastBidTime || record.HasLastCstBidEventLogID {
			return CurrentSpecialWinners{}, errors.New("zero last-CST bidder has attached values")
		}
	} else {
		address, err := canonicalNonZeroAddress("last CST bidder", record.LastCstBidderAddress)
		if err != nil {
			return CurrentSpecialWinners{}, err
		}
		standing := &LastCstBidderStanding{Address: address}
		if record.HasLastCstBidderLastBidTime {
			if record.LastCstBidderLastBidTime < 0 {
				return CurrentSpecialWinners{}, errors.New("invalid last-CST-bidder timestamp")
			}
			if record.LastCstBidderLastBidTime > 0 {
				lastBidAt := record.LastCstBidderLastBidTime
				standing.LastBidAt = &lastBidAt
			}
		}
		if record.HasLastCstBidEventLogID {
			if record.LastCstBidEventLogID < 1 {
				return CurrentSpecialWinners{}, errors.New("invalid last CST bid event-log ID")
			}
			eventLogID := record.LastCstBidEventLogID
			standing.EventLogId = &eventLogID
		}
		result.LastCstBidder = standing
	}
	return result, nil
}

func canonicalNonZeroAddress(name, value string) (string, error) {
	if !ethcommon.IsHexAddress(value) {
		return "", fmt.Errorf("invalid %s address", name)
	}
	address := ethcommon.HexToAddress(value)
	if address == (ethcommon.Address{}) {
		return "", fmt.Errorf("%s address is zero", name)
	}
	return address.Hex(), nil
}

func canonicalSnapshotAddress(name string, value ethcommon.Address) (string, error) {
	if value == (ethcommon.Address{}) {
		return "", fmt.Errorf("%s address is zero", name)
	}
	return value.Hex(), nil
}

func validPercent(value int64) bool {
	return value >= 0 && value <= 100
}

func validAuctionProgress(duration, elapsed int64) bool {
	return duration >= 0 && elapsed >= 0 && elapsed <= duration
}
