package v2

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func mapGlobalStatistics(record cgstore.GlobalStatisticsRecord) (CosmicGameGlobalStatistics, error) {
	totalBids, err := statisticsCount(record.TotalBids, "total bids")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	currentRoundBids, err := statisticsCount(record.CurrentRoundBids, "current-round bids")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	completedRounds, err := statisticsCount(record.CompletedRounds, "completed rounds")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	totalPrizeAwards, err := statisticsCount(record.TotalPrizeAwards, "total prize awards")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	prizeRegistryRows, err := statisticsCount(record.PrizeRegistryRows, "prize registry rows")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	uniqueBidders, err := statisticsCount(record.UniqueBidders, "unique bidders")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	uniqueWinners, err := statisticsCount(record.UniqueWinners, "unique winners")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	uniqueCSTStakers, err := statisticsCount(record.UniqueCSTStakers, "unique CST stakers")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	uniqueRandomWalkStakers, err := statisticsCount(record.UniqueRandomWalkStakers, "unique RandomWalk stakers")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	uniqueDualStakers, err := statisticsCount(record.UniqueDualStakers, "unique dual stakers")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	voluntaryDonationCount, err := statisticsCount(record.VoluntaryDonationCount, "voluntary donations")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	cosmicGameDonationCount, err := statisticsCount(record.CosmicGameDonationCount, "CosmicGame donations")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	charityWithdrawalCount, err := statisticsCount(record.CharityWithdrawalCount, "charity withdrawals")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	randomWalkTokensUsed, err := statisticsCount(record.RandomWalkTokensUsedInBids, "RandomWalk tokens used")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	donatedNFTCount, err := statisticsCount(record.DonatedNFTCount, "donated NFTs")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	cosmicSignatureMints, err := statisticsCount(record.CosmicSignatureMints, "Cosmic Signature mints")
	if err != nil {
		return CosmicGameGlobalStatistics{}, err
	}
	for name, value := range map[string]int64{
		"unique donors":                          record.UniqueDonors,
		"direct donations":                       record.DirectDonationCount,
		"named tokens":                           record.NamedTokens,
		"pending raffle withdrawals":             record.WinnersWithPendingRaffleWithdrawal,
		"CST bids":                               record.CSTBidCount,
		"marketing rewards":                      record.MarketingRewardCount,
		"CST staking total tokens":               record.CSTStaking.TotalTokensStaked,
		"CST staking active stakers":             record.CSTStaking.NumActiveStakers,
		"CST staking deposits":                   record.CSTStaking.NumDeposits,
		"RandomWalk staking total tokens":        record.RandomWalkStaking.TotalTokensStaked,
		"RandomWalk staking active stakers":      record.RandomWalkStaking.NumActiveStakers,
		"RandomWalk staking total tokens minted": record.RandomWalkStaking.TotalTokensMinted,
	} {
		if value < 0 {
			return CosmicGameGlobalStatistics{}, fmt.Errorf("%s is negative", name)
		}
	}
	amounts := map[string]string{
		"total prizes paid":             record.TotalPrizesPaidWei,
		"total ETH donated":             record.TotalEthDonatedWei,
		"voluntary donations":           record.VoluntaryDonationsTotalWei,
		"CosmicGame donations":          record.CosmicGameDonationsTotalWei,
		"direct donations":              record.DirectDonationsTotalWei,
		"charity withdrawals":           record.CharityWithdrawalsTotalWei,
		"raffle ETH deposits":           record.RaffleEthDepositsTotalWei,
		"raffle ETH withdrawn":          record.RaffleEthWithdrawnTotalWei,
		"chrono warrior ETH deposits":   record.ChronoWarriorEthDepositsTotalWei,
		"CST given in prizes":           record.CSTGivenInPrizesTotalWei,
		"CST consumed":                  record.CSTConsumedTotalWei,
		"marketing rewards":             record.MarketingRewardsTotalWei,
		"CST staking rewards":           record.CSTStaking.TotalReward,
		"CST staking unclaimed rewards": record.CSTStaking.UnclaimedReward,
	}
	canonical := make(map[string]string, len(amounts))
	for name, value := range amounts {
		amount, err := requiredAmount(value)
		if err != nil {
			return CosmicGameGlobalStatistics{}, fmt.Errorf("%s amount: %w", name, err)
		}
		canonical[name] = amount
	}
	distribution := make([]DonatedTokenStatistic, 0, len(record.DonatedTokenDistribution))
	for i := range record.DonatedTokenDistribution {
		item := record.DonatedTokenDistribution[i]
		if !ethcommon.IsHexAddress(item.ContractAddr) || item.NumDonatedTokens < 0 {
			return CosmicGameGlobalStatistics{}, errors.New("invalid donated-token distribution")
		}
		distribution = append(distribution, DonatedTokenStatistic{
			DonatedCount: item.NumDonatedTokens,
			TokenAddress: ethcommon.HexToAddress(item.ContractAddr).Hex(),
		})
	}
	return CosmicGameGlobalStatistics{
		CharityWithdrawalCount:           charityWithdrawalCount,
		CharityWithdrawalsTotalWei:       canonical["charity withdrawals"],
		ChronoWarriorEthDepositsTotalWei: canonical["chrono warrior ETH deposits"],
		CompletedRounds:                  completedRounds,
		CosmicGameDonationCount:          cosmicGameDonationCount,
		CosmicGameDonationsTotalWei:      canonical["CosmicGame donations"],
		CosmicSignatureMints:             cosmicSignatureMints,
		CstBidCount:                      record.CSTBidCount,
		CstConsumedTotalWei:              canonical["CST consumed"],
		CstGivenInPrizesTotalWei:         canonical["CST given in prizes"],
		CstStaking: CstStakingStatistics{
			ActiveStakers:      record.CSTStaking.NumActiveStakers,
			Deposits:           record.CSTStaking.NumDeposits,
			TotalRewardWei:     canonical["CST staking rewards"],
			TotalTokensStaked:  record.CSTStaking.TotalTokensStaked,
			UnclaimedRewardWei: canonical["CST staking unclaimed rewards"],
		},
		CurrentRoundBids:           currentRoundBids,
		DirectDonationCount:        record.DirectDonationCount,
		DirectDonationsTotalWei:    canonical["direct donations"],
		DonatedNftCount:            donatedNFTCount,
		DonatedTokenDistribution:   distribution,
		MarketingRewardCount:       record.MarketingRewardCount,
		MarketingRewardsTotalWei:   canonical["marketing rewards"],
		NamedTokens:                record.NamedTokens,
		PrizeRegistryRows:          prizeRegistryRows,
		RaffleEthDepositsTotalWei:  canonical["raffle ETH deposits"],
		RaffleEthWithdrawnTotalWei: canonical["raffle ETH withdrawn"],
		RandomWalkStaking: RandomWalkStakingStatistics{
			ActiveStakers:     record.RandomWalkStaking.NumActiveStakers,
			TotalTokensMinted: record.RandomWalkStaking.TotalTokensMinted,
			TotalTokensStaked: record.RandomWalkStaking.TotalTokensStaked,
		},
		RandomWalkTokensUsedInBids:         randomWalkTokensUsed,
		TotalBids:                          totalBids,
		TotalEthDonatedWei:                 canonical["total ETH donated"],
		TotalPrizeAwards:                   totalPrizeAwards,
		TotalPrizesPaidWei:                 canonical["total prizes paid"],
		UniqueBidders:                      uniqueBidders,
		UniqueCstStakers:                   uniqueCSTStakers,
		UniqueDonors:                       record.UniqueDonors,
		UniqueDualStakers:                  uniqueDualStakers,
		UniqueRandomWalkStakers:            uniqueRandomWalkStakers,
		UniqueWinners:                      uniqueWinners,
		VoluntaryDonationCount:             voluntaryDonationCount,
		VoluntaryDonationsTotalWei:         canonical["voluntary donations"],
		WinnersWithPendingRaffleWithdrawal: record.WinnersWithPendingRaffleWithdrawal,
	}, nil
}

func statisticsCount(value uint64, name string) (int64, error) {
	if value > maxSignedInt64 {
		return 0, fmt.Errorf("%s exceeds int64", name)
	}
	return int64(value), nil
}

func mapCounters(record cgprimitives.CGRecordCounters) (CosmicGameCounters, error) {
	if record.TotalBids < 0 || record.TotalPrizes < 0 || record.TotalDonatedNFTs < 0 {
		return CosmicGameCounters{}, errors.New("negative record counter")
	}
	return CosmicGameCounters{
		CompletedRounds: record.TotalPrizes,
		DonatedNfts:     record.TotalDonatedNFTs,
		TotalBids:       record.TotalBids,
	}, nil
}

func mapROILeaderboardEntry(record cgstore.ROILeaderboardRecord) (RoiLeaderboardEntry, error) {
	if record.BidderAid < 1 || record.NumBids < 0 || record.RoundsParticipated < 0 ||
		record.RoundsWon < 0 || record.PrizesCount < 0 ||
		record.CSTPrizesCount < 0 || record.NFTPrizesCount < 0 {
		return RoiLeaderboardEntry{}, errors.New("invalid ROI leaderboard identity or counts")
	}
	if !ethcommon.IsHexAddress(record.BidderAddr) {
		return RoiLeaderboardEntry{}, errors.New("invalid ROI bidder address")
	}
	totalETH, err := requiredAmount(record.TotalEthSpentWei)
	if err != nil {
		return RoiLeaderboardEntry{}, fmt.Errorf("total ETH spent: %w", err)
	}
	totalCST, err := requiredAmount(record.TotalCSTSpentWei)
	if err != nil {
		return RoiLeaderboardEntry{}, fmt.Errorf("total CST spent: %w", err)
	}
	ethWon, err := requiredAmount(record.EthWonWei)
	if err != nil {
		return RoiLeaderboardEntry{}, fmt.Errorf("ETH won: %w", err)
	}
	netProfit, err := canonicalSignedInteger(record.NetProfitWei)
	if err != nil {
		return RoiLeaderboardEntry{}, fmt.Errorf("net profit: %w", err)
	}
	roi, err := canonicalDecimal(record.ROIRatio, true)
	if err != nil {
		return RoiLeaderboardEntry{}, fmt.Errorf("ROI ratio: %w", err)
	}
	winRate, err := canonicalDecimal(record.WinRateRatio, false)
	if err != nil {
		return RoiLeaderboardEntry{}, fmt.Errorf("win-rate ratio: %w", err)
	}
	return RoiLeaderboardEntry{
		BidCount:           record.NumBids,
		BidderAddress:      ethcommon.HexToAddress(record.BidderAddr).Hex(),
		CstPrizeCount:      record.CSTPrizesCount,
		EthWonWei:          ethWon,
		NetProfitWei:       netProfit,
		NftPrizeCount:      record.NFTPrizesCount,
		PrizeCount:         record.PrizesCount,
		RoiRatio:           roi,
		RoundsParticipated: record.RoundsParticipated,
		RoundsWon:          record.RoundsWon,
		TotalCstSpentWei:   totalCST,
		TotalEthSpentWei:   totalETH,
		WinRateRatio:       winRate,
	}, nil
}

func mapClaimSummary(record cgstore.ClaimSummaryRecord, now time.Time) (ClaimSummary, error) {
	if record.RoundNum < 0 || record.EventLogID < 1 ||
		record.ClaimWindowTimeout < 0 || record.AwardedTimestamp < 0 ||
		record.EthAwarded < 0 || record.EthUnclaimed < 0 ||
		record.NFTAwarded < 0 || record.NFTUnclaimed < 0 ||
		record.ERC20Awarded < 0 || record.ERC20Unclaimed < 0 ||
		record.TotalAwarded < 0 || record.TotalUnclaimed < 0 ||
		record.AvgClaimLatencySecs < 0 {
		return ClaimSummary{}, errors.New("invalid claim summary")
	}
	if record.TotalAwarded != record.EthAwarded+record.NFTAwarded+record.ERC20Awarded ||
		record.TotalUnclaimed != record.EthUnclaimed+record.NFTUnclaimed+record.ERC20Unclaimed {
		return ClaimSummary{}, errors.New("inconsistent claim summary totals")
	}
	amount, err := requiredAmount(record.UnclaimedEthAmountWei)
	if err != nil {
		return ClaimSummary{}, fmt.Errorf("unclaimed ETH amount: %w", err)
	}
	return ClaimSummary{
		AvgClaimLatencySeconds: record.AvgClaimLatencySecs,
		AwardedAt:              time.Unix(record.AwardedTimestamp, 0).UTC(),
		ClaimWindowExpiresAt:   time.Unix(record.ClaimWindowTimeout, 0).UTC(),
		Erc20Awarded:           record.ERC20Awarded,
		Erc20Unclaimed:         record.ERC20Unclaimed,
		EthAwarded:             record.EthAwarded,
		EthUnclaimed:           record.EthUnclaimed,
		IsExpired:              !now.Before(time.Unix(record.ClaimWindowTimeout, 0)),
		NftAwarded:             record.NFTAwarded,
		NftUnclaimed:           record.NFTUnclaimed,
		Round:                  record.RoundNum,
		TotalAwarded:           record.TotalAwarded,
		TotalUnclaimed:         record.TotalUnclaimed,
		UnclaimedEthAmountWei:  amount,
	}, nil
}

func mapAssetClaimTransaction(record cgstore.ClaimTransactionRecord) (AssetClaimTransaction, error) {
	if record.RoundNum < 0 || record.EventLogID < 1 ||
		record.ClaimedAfterSecs < 0 || record.ClaimedTimestamp < 0 {
		return AssetClaimTransaction{}, errors.New("invalid claim transaction identity")
	}
	if !ethcommon.IsHexAddress(record.RecipientAddr) ||
		!ethcommon.IsHexAddress(record.BeneficiaryAddr) ||
		!isTransactionHash(record.TxHash) {
		return AssetClaimTransaction{}, errors.New("invalid claim transaction address or hash")
	}
	result := AssetClaimTransaction{
		AssetType:           claimAssetType(record.AssetType),
		BeneficiaryAddress:  ethcommon.HexToAddress(record.BeneficiaryAddr).Hex(),
		ClaimedAfterSeconds: record.ClaimedAfterSecs,
		ClaimedAt:           time.Unix(record.ClaimedTimestamp, 0).UTC(),
		RecipientAddress:    ethcommon.HexToAddress(record.RecipientAddr).Hex(),
		TransactionHash:     strings.ToLower(record.TxHash),
	}
	switch record.AssetType {
	case cgstore.ClaimAssetETH:
		if record.EthAmountWei == nil || record.TokenAddr != nil ||
			record.TokenID != nil || record.AmountBaseUnits != nil {
			return AssetClaimTransaction{}, errors.New("invalid ETH claim fields")
		}
		amount, err := requiredAmount(*record.EthAmountWei)
		if err != nil {
			return AssetClaimTransaction{}, err
		}
		result.EthAmountWei = &amount
	case cgstore.ClaimAssetERC721:
		if record.EthAmountWei != nil || record.TokenAddr == nil ||
			record.TokenID == nil || *record.TokenID < 0 || record.AmountBaseUnits != nil ||
			!ethcommon.IsHexAddress(*record.TokenAddr) {
			return AssetClaimTransaction{}, errors.New("invalid ERC-721 claim fields")
		}
		tokenAddress := ethcommon.HexToAddress(*record.TokenAddr).Hex()
		tokenID := *record.TokenID
		result.TokenAddress = &tokenAddress
		result.TokenId = &tokenID
	case cgstore.ClaimAssetERC20:
		if record.EthAmountWei != nil || record.TokenAddr == nil ||
			record.TokenID != nil || record.AmountBaseUnits == nil ||
			!ethcommon.IsHexAddress(*record.TokenAddr) {
			return AssetClaimTransaction{}, errors.New("invalid ERC-20 claim fields")
		}
		amount, err := requiredAmount(*record.AmountBaseUnits)
		if err != nil {
			return AssetClaimTransaction{}, err
		}
		tokenAddress := ethcommon.HexToAddress(*record.TokenAddr).Hex()
		result.TokenAddress = &tokenAddress
		result.AmountBaseUnits = &amount
	default:
		return AssetClaimTransaction{}, errors.New("unknown claim asset type")
	}
	return result, nil
}

func mapAttachedToken(record cgstore.AttachedTokenRecord) (AttachedToken, error) {
	if record.RoundNum < 0 || record.EventLogID < 1 || record.OccurredAt < 0 ||
		!ethcommon.IsHexAddress(record.ContributorAddr) ||
		!ethcommon.IsHexAddress(record.TokenAddr) ||
		!isTransactionHash(record.TxHash) {
		return AttachedToken{}, errors.New("invalid attached-token record")
	}
	result := AttachedToken{
		AssetType:          claimAssetType(record.AssetType),
		ContributorAddress: ethcommon.HexToAddress(record.ContributorAddr).Hex(),
		OccurredAt:         time.Unix(record.OccurredAt, 0).UTC(),
		TokenAddress:       ethcommon.HexToAddress(record.TokenAddr).Hex(),
		TransactionHash:    strings.ToLower(record.TxHash),
	}
	switch record.AssetType {
	case cgstore.ClaimAssetERC721:
		if record.TokenID == nil || *record.TokenID < 0 || record.AmountBaseUnits != nil {
			return AttachedToken{}, errors.New("invalid attached ERC-721 fields")
		}
		tokenID := *record.TokenID
		result.TokenId = &tokenID
	case cgstore.ClaimAssetERC20:
		if record.TokenID != nil || record.AmountBaseUnits == nil {
			return AttachedToken{}, errors.New("invalid attached ERC-20 fields")
		}
		amount, err := requiredAmount(*record.AmountBaseUnits)
		if err != nil {
			return AttachedToken{}, err
		}
		result.AmountBaseUnits = &amount
	default:
		return AttachedToken{}, errors.New("invalid attached-token asset type")
	}
	return result, nil
}

func mapUnclaimedItem(record cgstore.UnclaimedItemRecord) (UnclaimedItem, error) {
	if record.RoundNum < 0 || record.Segment < 0 || record.Segment > 2 ||
		record.Key < 1 || !ethcommon.IsHexAddress(record.RecipientAddr) {
		return UnclaimedItem{}, errors.New("invalid unclaimed-item identity")
	}
	result := UnclaimedItem{
		AssetType:        claimAssetType(record.AssetType),
		RecipientAddress: ethcommon.HexToAddress(record.RecipientAddr).Hex(),
	}
	switch record.AssetType {
	case cgstore.ClaimAssetETH:
		if record.EthAmountWei == nil || record.TokenAddr != nil ||
			record.TokenID != nil || record.AmountBaseUnits != nil {
			return UnclaimedItem{}, errors.New("invalid unclaimed ETH fields")
		}
		amount, err := requiredAmount(*record.EthAmountWei)
		if err != nil {
			return UnclaimedItem{}, err
		}
		result.EthAmountWei = &amount
	case cgstore.ClaimAssetERC721:
		if record.EthAmountWei != nil || record.TokenAddr == nil ||
			record.TokenID == nil || *record.TokenID < 0 || record.AmountBaseUnits != nil ||
			!ethcommon.IsHexAddress(*record.TokenAddr) {
			return UnclaimedItem{}, errors.New("invalid unclaimed ERC-721 fields")
		}
		tokenAddress := ethcommon.HexToAddress(*record.TokenAddr).Hex()
		tokenID := *record.TokenID
		result.TokenAddress = &tokenAddress
		result.TokenId = &tokenID
	case cgstore.ClaimAssetERC20:
		if record.EthAmountWei != nil || record.TokenAddr == nil ||
			record.TokenID != nil || record.AmountBaseUnits == nil ||
			!ethcommon.IsHexAddress(*record.TokenAddr) {
			return UnclaimedItem{}, errors.New("invalid unclaimed ERC-20 fields")
		}
		amount, err := requiredAmount(*record.AmountBaseUnits)
		if err != nil {
			return UnclaimedItem{}, err
		}
		tokenAddress := ethcommon.HexToAddress(*record.TokenAddr).Hex()
		result.TokenAddress = &tokenAddress
		result.AmountBaseUnits = &amount
	default:
		return UnclaimedItem{}, errors.New("unknown unclaimed asset type")
	}
	return result, nil
}

func claimAssetType(asset cgstore.ClaimAssetType) ClaimAssetType {
	switch asset {
	case cgstore.ClaimAssetETH:
		return ClaimAssetEth
	case cgstore.ClaimAssetERC721:
		return ClaimAssetErc721
	case cgstore.ClaimAssetERC20:
		return ClaimAssetErc20
	default:
		return ClaimAssetType(asset)
	}
}

func canonicalSignedInteger(value string) (string, error) {
	if value == "" {
		return "", errors.New("integer is empty")
	}
	number, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "", fmt.Errorf("invalid signed integer %q", value)
	}
	return number.String(), nil
}

func canonicalDecimal(value string, allowNegative bool) (string, error) {
	if value == "" {
		return "", errors.New("decimal is empty")
	}
	negative := strings.HasPrefix(value, "-")
	if negative {
		if !allowNegative {
			return "", errors.New("decimal is negative")
		}
		value = value[1:]
	}
	parts := strings.Split(value, ".")
	if len(parts) > 2 || parts[0] == "" || (len(parts) == 2 && parts[1] == "") {
		return "", errors.New("invalid decimal syntax")
	}
	for _, part := range parts {
		for _, char := range part {
			if char < '0' || char > '9' {
				return "", errors.New("invalid decimal digit")
			}
		}
	}
	integer := strings.TrimLeft(parts[0], "0")
	if integer == "" {
		integer = "0"
	}
	fraction := ""
	if len(parts) == 2 {
		fraction = strings.TrimRight(parts[1], "0")
	}
	result := integer
	if fraction != "" {
		result += "." + fraction
	}
	if negative && result != "0" {
		result = "-" + result
	}
	return result, nil
}

func compareDecimal(left, right string) (int, error) {
	leftValue, ok := new(big.Rat).SetString(left)
	if !ok {
		return 0, fmt.Errorf("invalid decimal %q", left)
	}
	rightValue, ok := new(big.Rat).SetString(right)
	if !ok {
		return 0, fmt.Errorf("invalid decimal %q", right)
	}
	return leftValue.Cmp(rightValue), nil
}
