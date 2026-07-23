package v2

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

const maxSignedInt64 = uint64(1<<63 - 1)

func mapRoundSummary(record cgmodel.CGRoundRec) (CosmicGameRoundSummary, error) {
	roundNum, err := roundNumber(record.RoundNum)
	if err != nil {
		return CosmicGameRoundSummary{}, err
	}
	claim, err := mapClaimTransaction(record.ClaimPrizeTx.Tx)
	if err != nil {
		return CosmicGameRoundSummary{}, err
	}
	mainPrize, err := mapMainPrize(record.MainPrize)
	if err != nil {
		return CosmicGameRoundSummary{}, err
	}
	if err := validateRoundCounts(record.RoundStats); err != nil {
		return CosmicGameRoundSummary{}, err
	}
	raffleAmount, err := amountOrZero(record.RoundStats.TotalRaffleEthDeposits)
	if err != nil {
		return CosmicGameRoundSummary{}, fmt.Errorf("raffle ETH deposits: %w", err)
	}

	return CosmicGameRoundSummary{
		CompletedAt:          claim.OccurredAt,
		DonatedNftCount:      record.RoundStats.TotalDonatedNFTs,
		Erc20DonationCount:   record.RoundStats.NumERC20Donations,
		MainPrize:            mainPrize,
		RaffleEthDepositsWei: raffleAmount,
		RaffleNftCount:       record.RoundStats.TotalRaffleNFTs,
		Round:                roundNum,
		Status:               Completed,
		TotalBids:            record.RoundStats.TotalBids,
	}, nil
}

func mapRound(record cgmodel.CGRoundRec) (CosmicGameRound, error) {
	roundNum, err := roundNumber(record.RoundNum)
	if err != nil {
		return CosmicGameRound{}, err
	}
	claim, err := mapClaimTransaction(record.ClaimPrizeTx.Tx)
	if err != nil {
		return CosmicGameRound{}, err
	}
	mainPrize, err := mapMainPrize(record.MainPrize)
	if err != nil {
		return CosmicGameRound{}, err
	}
	statistics, err := mapRoundStatistics(record.RoundStats)
	if err != nil {
		return CosmicGameRound{}, err
	}
	charity, err := mapCharityAllocation(record.CharityDeposit)
	if err != nil {
		return CosmicGameRound{}, err
	}
	staking, err := mapStakingAllocation(record.StakingDeposit)
	if err != nil {
		return CosmicGameRound{}, err
	}
	endurance, err := mapTokenPrize("endurance champion", record.EnduranceChampion.WinnerAddr,
		record.EnduranceChampion.NftTokenId, record.EnduranceChampion.CstAmount)
	if err != nil {
		return CosmicGameRound{}, err
	}
	lastCst, err := mapTokenPrize("last CST bidder", record.LastCstBidder.WinnerAddr,
		record.LastCstBidder.NftTokenId, record.LastCstBidder.CstAmount)
	if err != nil {
		return CosmicGameRound{}, err
	}
	chrono, err := mapChronoWarrior(record.ChronoWarrior)
	if err != nil {
		return CosmicGameRound{}, err
	}

	return CosmicGameRound{
		Charity:           charity,
		ChronoWarrior:     chrono,
		Claim:             claim,
		EnduranceChampion: endurance,
		LastCstBidder:     lastCst,
		MainPrize:         mainPrize,
		Round:             roundNum,
		Staking:           staking,
		Statistics:        statistics,
		Status:            Completed,
	}, nil
}

func mapClaimTransaction(tx cgmodel.Transaction) (ClaimTransaction, error) {
	if tx.EvtLogId < 1 || tx.BlockNum < 0 {
		return ClaimTransaction{}, errors.New("invalid claim transaction identity")
	}
	if !isTransactionHash(tx.TxHash) {
		return ClaimTransaction{}, errors.New("invalid claim transaction hash")
	}
	occurredAt, err := time.Parse(time.RFC3339Nano, tx.DateTime)
	if err != nil {
		return ClaimTransaction{}, fmt.Errorf("parse claim timestamp: %w", err)
	}
	return ClaimTransaction{
		BlockNumber:     tx.BlockNum,
		EventLogId:      tx.EvtLogId,
		OccurredAt:      occurredAt.UTC(),
		TransactionHash: strings.ToLower(tx.TxHash),
	}, nil
}

func mapMainPrize(prize cgmodel.CGMainPrizeInfo) (MainPrize, error) {
	if !ethcommon.IsHexAddress(prize.WinnerAddr) {
		return MainPrize{}, errors.New("invalid main-prize winner address")
	}
	ethAmount, err := requiredAmount(prize.EthAmount)
	if err != nil {
		return MainPrize{}, fmt.Errorf("main-prize ETH amount: %w", err)
	}
	if prize.NftTokenId > maxSignedInt64 {
		return MainPrize{}, errors.New("main-prize token id exceeds int64")
	}
	if prize.TimeoutTs < 0 {
		return MainPrize{}, errors.New("negative secondary-prize claim deadline")
	}
	if prize.NumCSNfts < 0 {
		return MainPrize{}, errors.New("negative main-prize NFT count")
	}

	result := MainPrize{
		EthAmountWei:  ethAmount,
		NftTokenId:    int64(prize.NftTokenId),
		WinnerAddress: ethcommon.HexToAddress(prize.WinnerAddr).Hex(),
	}
	if prize.CstAmount != "" {
		cstAmount, err := requiredAmount(prize.CstAmount)
		if err != nil {
			return MainPrize{}, fmt.Errorf("main-prize CST amount: %w", err)
		}
		result.CstAmountWei = &cstAmount
	}
	if prize.Seed != "" && prize.Seed != "???" {
		seed := prize.Seed
		result.Seed = &seed
	}
	if prize.TimeoutTs > 0 {
		deadline := time.Unix(prize.TimeoutTs, 0).UTC()
		result.SecondaryPrizeClaimDeadline = &deadline
	}
	if prize.NumCSNfts > 1 {
		if prize.NftTokenId > maxSignedInt64-uint64(prize.NumCSNfts-1) {
			return MainPrize{}, errors.New("main-prize token range exceeds int64")
		}
		tokenIDs := prize.NftTokenIds
		if len(tokenIDs) == 0 {
			tokenIDs = make([]int64, prize.NumCSNfts)
			for i := range tokenIDs {
				tokenIDs[i] = int64(prize.NftTokenId) + int64(i)
			}
		}
		if int64(len(tokenIDs)) != prize.NumCSNfts {
			return MainPrize{}, errors.New("main-prize token count does not match token IDs")
		}
		for i, tokenID := range tokenIDs {
			if tokenID != int64(prize.NftTokenId)+int64(i) {
				return MainPrize{}, errors.New("main-prize token IDs are not sequential")
			}
		}
		count := prize.NumCSNfts
		copied := append([]int64(nil), tokenIDs...)
		result.NumCosmicSignatureNfts = &count
		result.NftTokenIds = &copied
	}
	return result, nil
}

func mapRoundStatistics(stats cgmodel.CGRoundStats) (RoundStatistics, error) {
	if err := validateRoundCounts(stats); err != nil {
		return RoundStatistics{}, err
	}
	raffleAmount, err := amountOrZero(stats.TotalRaffleEthDeposits)
	if err != nil {
		return RoundStatistics{}, fmt.Errorf("raffle ETH deposits: %w", err)
	}
	donationAmount, err := amountOrZero(stats.TotalDonatedAmount)
	if err != nil {
		return RoundStatistics{}, fmt.Errorf("donation amount: %w", err)
	}
	result := RoundStatistics{
		DonatedNftCount:      stats.TotalDonatedNFTs,
		DonationAmountWei:    donationAmount,
		DonationCount:        stats.TotalDonatedCount,
		Erc20DonationCount:   stats.NumERC20Donations,
		RaffleEthDepositsWei: raffleAmount,
		RaffleNftCount:       stats.TotalRaffleNFTs,
		TotalBids:            stats.TotalBids,
	}

	if result.ParameterWindowStartedAt, err = optionalTimestamp(stats.ParamWindowStartTime); err != nil {
		return RoundStatistics{}, fmt.Errorf("parameter-window start: %w", err)
	}
	if stats.ActivationTime < 0 {
		return RoundStatistics{}, errors.New("negative activation timestamp")
	}
	if stats.ActivationTime > 0 {
		value := time.Unix(stats.ActivationTime, 0).UTC()
		result.ActivatedAt = &value
	}
	if stats.ParamWindowDurationSeconds < 0 ||
		stats.RoundDurationSeconds < 0 ||
		stats.EnduranceChampionDuration < 0 ||
		stats.ChronoWarriorDuration < 0 {
		return RoundStatistics{}, errors.New("negative round duration")
	}
	if stats.ParamWindowDurationSeconds > 0 {
		value := stats.ParamWindowDurationSeconds
		result.ParameterWindowDurationSeconds = &value
	}
	if result.RoundStartedAt, err = optionalTimestamp(stats.RoundStartTime); err != nil {
		return RoundStatistics{}, fmt.Errorf("round start: %w", err)
	}
	if result.RoundEndedAt, err = optionalTimestamp(stats.RoundEndTime); err != nil {
		return RoundStatistics{}, fmt.Errorf("round end: %w", err)
	}
	if stats.RoundDurationSeconds > 0 {
		value := stats.RoundDurationSeconds
		result.RoundDurationSeconds = &value
	}
	if stats.EnduranceChampionDuration > 0 {
		value := stats.EnduranceChampionDuration
		result.EnduranceChampionDurationSeconds = &value
	}
	if stats.ChronoWarriorDuration > 0 {
		value := stats.ChronoWarriorDuration
		result.ChronoWarriorDurationSeconds = &value
	}
	return result, nil
}

func validateRoundCounts(stats cgmodel.CGRoundStats) error {
	if stats.TotalBids < 0 ||
		stats.TotalDonatedNFTs < 0 ||
		stats.NumERC20Donations < 0 ||
		stats.TotalRaffleNFTs < 0 ||
		stats.TotalDonatedCount < 0 {
		return errors.New("negative round aggregate")
	}
	return nil
}

func mapCharityAllocation(deposit cgmodel.CGCharityDeposit) (*CharityAllocation, error) {
	if deposit.CharityAddress == "" && deposit.CharityAmount == "" {
		return nil, nil
	}
	if deposit.CharityAddress == "" {
		return nil, errors.New("charity allocation has no address")
	}
	amount, err := amountOrZero(deposit.CharityAmount)
	if err != nil {
		return nil, fmt.Errorf("charity amount: %w", err)
	}
	parts := strings.Split(deposit.CharityAddress, ",")
	addresses := make([]string, 0, len(parts))
	for _, part := range parts {
		address := strings.TrimSpace(part)
		if !ethcommon.IsHexAddress(address) {
			return nil, fmt.Errorf("invalid charity address %q", address)
		}
		addresses = append(addresses, ethcommon.HexToAddress(address).Hex())
	}
	slices.Sort(addresses)
	return &CharityAllocation{AmountWei: amount, CharityAddresses: addresses}, nil
}

func mapStakingAllocation(deposit cgmodel.CGStakingDeposit) (*StakingAllocation, error) {
	if deposit.StakingDepositId == -1 {
		return nil, nil
	}
	if deposit.StakingDepositId < -1 || deposit.StakingNumStakedTokens < 0 {
		return nil, errors.New("invalid staking allocation identity")
	}
	amount, err := requiredAmount(deposit.StakingDepositAmount)
	if err != nil {
		return nil, fmt.Errorf("staking amount: %w", err)
	}
	perToken, err := requiredAmount(deposit.StakingPerToken)
	if err != nil {
		return nil, fmt.Errorf("staking per-token amount: %w", err)
	}
	return &StakingAllocation{
		AmountPerTokenWei: perToken,
		AmountWei:         amount,
		DepositId:         deposit.StakingDepositId,
		StakedTokenCount:  deposit.StakingNumStakedTokens,
	}, nil
}

func mapTokenPrize(name, winner string, tokenID int64, amount string) (*TokenPrize, error) {
	if winner == "" && amount == "" && tokenID == 0 {
		return nil, nil
	}
	if !ethcommon.IsHexAddress(winner) {
		return nil, fmt.Errorf("invalid %s winner address", name)
	}
	if tokenID < 0 {
		return nil, fmt.Errorf("negative %s token id", name)
	}
	cstAmount, err := requiredAmount(amount)
	if err != nil {
		return nil, fmt.Errorf("%s CST amount: %w", name, err)
	}
	return &TokenPrize{
		CstAmountWei:  cstAmount,
		NftTokenId:    tokenID,
		WinnerAddress: ethcommon.HexToAddress(winner).Hex(),
	}, nil
}

func mapChronoWarrior(prize cgmodel.CGChronoWarriorPrize) (*ChronoWarriorPrize, error) {
	if prize.WinnerAddr == "" && prize.EthAmount == "" && prize.CstAmount == "" && prize.NftTokenId == 0 {
		return nil, nil
	}
	if !ethcommon.IsHexAddress(prize.WinnerAddr) {
		return nil, errors.New("invalid chrono-warrior winner address")
	}
	if prize.NftTokenId < 0 {
		return nil, errors.New("negative chrono-warrior token id")
	}
	ethAmount, err := requiredAmount(prize.EthAmount)
	if err != nil {
		return nil, fmt.Errorf("chrono-warrior ETH amount: %w", err)
	}
	cstAmount, err := requiredAmount(prize.CstAmount)
	if err != nil {
		return nil, fmt.Errorf("chrono-warrior CST amount: %w", err)
	}
	return &ChronoWarriorPrize{
		CstAmountWei:  cstAmount,
		EthAmountWei:  ethAmount,
		NftTokenId:    prize.NftTokenId,
		WinnerAddress: ethcommon.HexToAddress(prize.WinnerAddr).Hex(),
	}, nil
}

func optionalTimestamp(value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}
	parsed, err := time.Parse(time.RFC3339Nano, value)
	if err == nil {
		parsed = parsed.UTC()
		return &parsed, nil
	}

	// CosmicGameRoundStatistics is shared with frozen v1 and therefore
	// retains PostgreSQL's timestamptz text representation. Accept both of
	// PostgreSQL's UTC-offset widths here while keeping the v2 wire format
	// strictly RFC 3339.
	for _, layout := range [...]string{
		"2006-01-02 15:04:05Z07",
		"2006-01-02 15:04:05.999999999Z07",
		"2006-01-02 15:04:05Z07:00",
		"2006-01-02 15:04:05.999999999Z07:00",
	} {
		parsed, err = time.Parse(layout, value)
		if err == nil {
			parsed = parsed.UTC()
			return &parsed, nil
		}
	}
	return nil, err
}

func amountOrZero(value string) (string, error) {
	if value == "" {
		return "0", nil
	}
	return requiredAmount(value)
}

func roundNumber(value uint64) (int64, error) {
	if value > maxSignedInt64 {
		return 0, errors.New("round number exceeds int64")
	}
	return int64(value), nil
}
