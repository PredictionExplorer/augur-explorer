package v2

import (
	"errors"
	"fmt"
	"math/big"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func mapStakingActionKind(kind cgstore.UserStakingActionKind) (StakingActionType, error) {
	switch kind {
	case cgstore.UserStakingActionStake:
		return Stake, nil
	case cgstore.UserStakingActionUnstake:
		return Unstake, nil
	default:
		return "", fmt.Errorf("unknown staking action kind %q", kind)
	}
}

func mapUserCstStakingAction(record cgstore.UserStakingActionRecord) (UserCstStakingAction, error) {
	if record.ActionID < 0 || record.TokenID < 0 || record.TotalStakedNfts < 0 {
		return UserCstStakingAction{}, errors.New("invalid staking action identity")
	}
	kind, err := mapStakingActionKind(record.Kind)
	if err != nil {
		return UserCstStakingAction{}, err
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return UserCstStakingAction{}, fmt.Errorf("staking action transaction: %w", err)
	}

	result := UserCstStakingAction{
		ActionId:        record.ActionID,
		ActionType:      kind,
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		NftTokenId:      record.TokenID,
		OccurredAt:      transaction.OccurredAt,
		TotalStakedNfts: record.TotalStakedNfts,
		TransactionHash: transaction.TransactionHash,
	}
	switch kind {
	case Stake:
		if record.RewardWei != "" {
			return UserCstStakingAction{}, errors.New("stake action carries a reward")
		}
	case Unstake:
		reward, err := requiredAmount(record.RewardWei)
		if err != nil {
			return UserCstStakingAction{}, fmt.Errorf("unstake reward: %w", err)
		}
		result.RewardWei = &reward
	}
	return result, nil
}

func mapUserRandomWalkStakingAction(record cgstore.UserStakingActionRecord) (UserRandomWalkStakingAction, error) {
	if record.ActionID < 0 || record.TokenID < 0 || record.TotalStakedNfts < 0 {
		return UserRandomWalkStakingAction{}, errors.New("invalid staking action identity")
	}
	kind, err := mapStakingActionKind(record.Kind)
	if err != nil {
		return UserRandomWalkStakingAction{}, err
	}
	if record.RewardWei != "" {
		return UserRandomWalkStakingAction{}, errors.New("RandomWalk staking action carries a reward")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return UserRandomWalkStakingAction{}, fmt.Errorf("staking action transaction: %w", err)
	}
	return UserRandomWalkStakingAction{
		ActionId:        record.ActionID,
		ActionType:      kind,
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		NftTokenId:      record.TokenID,
		OccurredAt:      transaction.OccurredAt,
		TotalStakedNfts: record.TotalStakedNfts,
		TransactionHash: transaction.TransactionHash,
	}, nil
}

func mapUserStakedCstToken(record cgstore.UserStakedCstTokenRecord) (UserStakedCstToken, error) {
	if record.ActionID < 0 || record.TokenID < 0 || record.MintRound < 0 {
		return UserStakedCstToken{}, errors.New("invalid staked token identity")
	}
	if record.Seed == "" {
		return UserStakedCstToken{}, errors.New("staked token misses its mint seed")
	}
	transaction, err := mapClaimTransaction(record.StakeTx)
	if err != nil {
		return UserStakedCstToken{}, fmt.Errorf("stake transaction: %w", err)
	}
	result := UserStakedCstToken{
		ActionId:        record.ActionID,
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		MintRound:       record.MintRound,
		NftTokenId:      record.TokenID,
		Seed:            record.Seed,
		StakedAt:        transaction.OccurredAt,
		TransactionHash: transaction.TransactionHash,
	}
	if record.TokenName != "" {
		name := record.TokenName
		result.TokenName = &name
	}
	return result, nil
}

func mapUserStakedRandomWalkToken(record cgstore.UserStakedRwalkTokenRecord) (UserStakedRandomWalkToken, error) {
	if record.ActionID < 0 || record.TokenID < 0 {
		return UserStakedRandomWalkToken{}, errors.New("invalid staked token identity")
	}
	transaction, err := mapClaimTransaction(record.StakeTx)
	if err != nil {
		return UserStakedRandomWalkToken{}, fmt.Errorf("stake transaction: %w", err)
	}
	return UserStakedRandomWalkToken{
		ActionId:        record.ActionID,
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		NftTokenId:      record.TokenID,
		StakedAt:        transaction.OccurredAt,
		TransactionHash: transaction.TransactionHash,
	}, nil
}

func mapUserStakingDeposit(record cgstore.UserStakingDepositRecord) (UserStakingDeposit, error) {
	if record.DepositID < 0 || record.RoundNum < 0 {
		return UserStakingDeposit{}, errors.New("invalid staking deposit identity")
	}
	if record.TotalStakedNfts < 1 || record.StakedNftCount < 1 {
		return UserStakingDeposit{}, errors.New("staking deposit without staked tokens")
	}
	if record.ClaimedNftCount < 0 || record.PendingNftCount < 0 ||
		record.ClaimedNftCount+record.PendingNftCount != record.StakedNftCount {
		return UserStakingDeposit{}, errors.New("staking deposit token counts are inconsistent")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return UserStakingDeposit{}, fmt.Errorf("staking deposit transaction: %w", err)
	}
	totalDeposit, err := requiredAmount(record.TotalDepositWei)
	if err != nil {
		return UserStakingDeposit{}, fmt.Errorf("staking deposit total: %w", err)
	}
	amountPerToken, err := requiredAmount(record.AmountPerTokenWei)
	if err != nil {
		return UserStakingDeposit{}, fmt.Errorf("staking deposit per-token amount: %w", err)
	}
	reward, err := requiredAmount(record.AmountDepositedWei)
	if err != nil {
		return UserStakingDeposit{}, fmt.Errorf("staking deposit reward: %w", err)
	}
	pendingAccumulator, err := requiredAmount(record.AmountToClaimWei)
	if err != nil {
		return UserStakingDeposit{}, fmt.Errorf("staking deposit pending accumulator: %w", err)
	}
	claimed, err := requiredAmount(record.ClaimedRewardWei)
	if err != nil {
		return UserStakingDeposit{}, fmt.Errorf("staking deposit claimed total: %w", err)
	}
	pending, err := requiredAmount(record.PendingRewardWei)
	if err != nil {
		return UserStakingDeposit{}, fmt.Errorf("staking deposit pending total: %w", err)
	}
	// The smallest reward units must add up to the per-staker accumulator,
	// and their uncollected part must equal its amount_to_claim column;
	// divergence means the trigger bookkeeping drifted.
	if !amountsAddUp(reward, claimed, pending) {
		return UserStakingDeposit{}, errors.New("staking deposit totals are inconsistent")
	}
	if pending != pendingAccumulator {
		return UserStakingDeposit{}, errors.New("staking deposit pending accumulator diverges")
	}

	return UserStakingDeposit{
		AmountPerTokenWei: amountPerToken,
		BlockNumber:       transaction.BlockNumber,
		ClaimedNftCount:   record.ClaimedNftCount,
		ClaimedWei:        claimed,
		DepositId:         record.DepositID,
		EventLogId:        transaction.EventLogId,
		FullyClaimed:      record.PendingNftCount == 0,
		OccurredAt:        transaction.OccurredAt,
		PendingNftCount:   record.PendingNftCount,
		PendingWei:        pending,
		RewardWei:         reward,
		Round:             record.RoundNum,
		StakedNftCount:    record.StakedNftCount,
		TotalDepositWei:   totalDeposit,
		TotalStakedNfts:   record.TotalStakedNfts,
		TransactionHash:   transaction.TransactionHash,
	}, nil
}

// amountsAddUp proves total = left + right over the exact integer values.
func amountsAddUp(total, left, right string) bool {
	totalInt, ok := new(big.Int).SetString(total, 10)
	if !ok {
		return false
	}
	leftInt, ok := new(big.Int).SetString(left, 10)
	if !ok {
		return false
	}
	rightInt, ok := new(big.Int).SetString(right, 10)
	if !ok {
		return false
	}
	return totalInt.Cmp(new(big.Int).Add(leftInt, rightInt)) == 0
}

func mapUserStakingDepositReward(record cgstore.UserStakingDepositRewardRecord) (UserStakingDepositReward, error) {
	if record.ActionID < 0 || record.TokenID < 0 {
		return UserStakingDepositReward{}, errors.New("invalid staking reward identity")
	}
	reward, err := requiredAmount(record.RewardWei)
	if err != nil {
		return UserStakingDepositReward{}, fmt.Errorf("staking reward amount: %w", err)
	}
	return UserStakingDepositReward{
		ActionId:   record.ActionID,
		Claimed:    record.Claimed,
		NftTokenId: record.TokenID,
		RewardWei:  reward,
	}, nil
}

func mapUserStakingTokenReward(record cgstore.UserStakingTokenRewardRecord) (UserStakingTokenReward, error) {
	if record.TokenID < 0 {
		return UserStakingTokenReward{}, errors.New("invalid staking token reward identity")
	}
	total, err := requiredAmount(record.TotalWei)
	if err != nil {
		return UserStakingTokenReward{}, fmt.Errorf("staking token reward total: %w", err)
	}
	collected, err := requiredAmount(record.CollectedWei)
	if err != nil {
		return UserStakingTokenReward{}, fmt.Errorf("staking token reward collected total: %w", err)
	}
	pending, err := requiredAmount(record.PendingWei)
	if err != nil {
		return UserStakingTokenReward{}, fmt.Errorf("staking token reward pending total: %w", err)
	}
	if !amountsAddUp(total, collected, pending) {
		return UserStakingTokenReward{}, errors.New("staking token reward totals are inconsistent")
	}
	return UserStakingTokenReward{
		CollectedWei: collected,
		NftTokenId:   record.TokenID,
		PendingWei:   pending,
		TotalWei:     total,
	}, nil
}

func mapUserStakingTokenRewardDeposit(record cgstore.UserStakingTokenRewardDepositRecord) (UserStakingTokenRewardDeposit, error) {
	if record.DepositID < 0 || record.RoundNum < 0 {
		return UserStakingTokenRewardDeposit{}, errors.New("invalid staking token deposit identity")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return UserStakingTokenRewardDeposit{}, fmt.Errorf("staking deposit transaction: %w", err)
	}
	reward, err := requiredAmount(record.RewardWei)
	if err != nil {
		return UserStakingTokenRewardDeposit{}, fmt.Errorf("staking reward amount: %w", err)
	}
	return UserStakingTokenRewardDeposit{
		BlockNumber:     transaction.BlockNumber,
		Claimed:         record.Claimed,
		DepositId:       record.DepositID,
		EventLogId:      transaction.EventLogId,
		OccurredAt:      transaction.OccurredAt,
		RewardWei:       reward,
		Round:           record.RoundNum,
		TransactionHash: transaction.TransactionHash,
	}, nil
}
