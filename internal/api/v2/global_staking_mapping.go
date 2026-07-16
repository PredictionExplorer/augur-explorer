package v2

import (
	"errors"
	"fmt"
	"math/big"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func mapGlobalCstStakingAction(
	record cgstore.GlobalStakingActionRecord,
) (GlobalCstStakingAction, error) {
	if record.ActionID < 0 || record.TokenID < 0 || record.RoundNum < 0 ||
		record.TotalStakedNfts < 0 {
		return GlobalCstStakingAction{}, errors.New("invalid global CST staking action identity")
	}
	kind, err := mapStakingActionKind(record.Kind)
	if err != nil {
		return GlobalCstStakingAction{}, err
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return GlobalCstStakingAction{}, fmt.Errorf("global CST staking action transaction: %w", err)
	}
	address, err := canonicalNonZeroAddress("CST staker", record.StakerAddress)
	if err != nil {
		return GlobalCstStakingAction{}, err
	}
	result := GlobalCstStakingAction{
		ActionId:        record.ActionID,
		ActionType:      kind,
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		NftTokenId:      record.TokenID,
		OccurredAt:      transaction.OccurredAt,
		Round:           record.RoundNum,
		StakerAddress:   address,
		TotalStakedNfts: record.TotalStakedNfts,
		TransactionHash: transaction.TransactionHash,
	}
	switch kind {
	case Stake:
		if record.RewardWei != "" || record.RewardPerTokenWei != "" {
			return GlobalCstStakingAction{}, errors.New("CST stake action carries a reward")
		}
	case Unstake:
		reward, err := requiredAmount(record.RewardWei)
		if err != nil {
			return GlobalCstStakingAction{}, fmt.Errorf("CST unstake reward: %w", err)
		}
		perToken, err := requiredAmount(record.RewardPerTokenWei)
		if err != nil {
			return GlobalCstStakingAction{}, fmt.Errorf("CST unstake per-token reward: %w", err)
		}
		result.RewardWei = &reward
		result.RewardPerTokenWei = &perToken
	}
	return result, nil
}

func mapGlobalRandomWalkStakingAction(
	record cgstore.GlobalStakingActionRecord,
) (GlobalRandomWalkStakingAction, error) {
	if record.ActionID < 0 || record.TokenID < 0 || record.RoundNum < 0 ||
		record.TotalStakedNfts < 0 {
		return GlobalRandomWalkStakingAction{}, errors.New("invalid global RandomWalk staking action identity")
	}
	kind, err := mapStakingActionKind(record.Kind)
	if err != nil {
		return GlobalRandomWalkStakingAction{}, err
	}
	if record.RewardWei != "" || record.RewardPerTokenWei != "" {
		return GlobalRandomWalkStakingAction{}, errors.New("RandomWalk staking action carries a reward")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return GlobalRandomWalkStakingAction{}, fmt.Errorf("global RandomWalk staking action transaction: %w", err)
	}
	address, err := canonicalNonZeroAddress("RandomWalk staker", record.StakerAddress)
	if err != nil {
		return GlobalRandomWalkStakingAction{}, err
	}
	return GlobalRandomWalkStakingAction{
		ActionId:        record.ActionID,
		ActionType:      kind,
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		NftTokenId:      record.TokenID,
		OccurredAt:      transaction.OccurredAt,
		Round:           record.RoundNum,
		StakerAddress:   address,
		TotalStakedNfts: record.TotalStakedNfts,
		TransactionHash: transaction.TransactionHash,
	}, nil
}

func mapGlobalCstStakingActionDetail(
	record cgmodel.CGStakeUnstakeCombined,
) (GlobalCstStakingActionDetail, error) {
	stakeRecord := globalStakingRecordFromStake(record.Stake)
	stake, err := mapGlobalCstStakingAction(stakeRecord)
	if err != nil {
		return GlobalCstStakingActionDetail{}, fmt.Errorf("stake: %w", err)
	}
	result := GlobalCstStakingActionDetail{Stake: stake}
	if record.Unstake.Tx.EvtLogId == 0 {
		return result, nil
	}
	unstakeRecord := globalStakingRecordFromUnstake(record.Unstake, true)
	if err := sameStakingLifecycle(stakeRecord, unstakeRecord); err != nil {
		return GlobalCstStakingActionDetail{}, err
	}
	unstake, err := mapGlobalCstStakingAction(unstakeRecord)
	if err != nil {
		return GlobalCstStakingActionDetail{}, fmt.Errorf("unstake: %w", err)
	}
	result.Unstake = &unstake
	return result, nil
}

func mapGlobalRandomWalkStakingActionDetail(
	record cgmodel.CGStakeUnstakeCombined,
) (GlobalRandomWalkStakingActionDetail, error) {
	stakeRecord := globalStakingRecordFromStake(record.Stake)
	stake, err := mapGlobalRandomWalkStakingAction(stakeRecord)
	if err != nil {
		return GlobalRandomWalkStakingActionDetail{}, fmt.Errorf("stake: %w", err)
	}
	result := GlobalRandomWalkStakingActionDetail{Stake: stake}
	if record.Unstake.Tx.EvtLogId == 0 {
		return result, nil
	}
	unstakeRecord := globalStakingRecordFromUnstake(record.Unstake, false)
	if err := sameStakingLifecycle(stakeRecord, unstakeRecord); err != nil {
		return GlobalRandomWalkStakingActionDetail{}, err
	}
	unstake, err := mapGlobalRandomWalkStakingAction(unstakeRecord)
	if err != nil {
		return GlobalRandomWalkStakingActionDetail{}, fmt.Errorf("unstake: %w", err)
	}
	result.Unstake = &unstake
	return result, nil
}

func globalStakingRecordFromStake(record cgmodel.CGStakeActionInfoRec) cgstore.GlobalStakingActionRecord {
	return cgstore.GlobalStakingActionRecord{
		Tx:              record.Tx,
		Kind:            cgstore.UserStakingActionStake,
		StakerAid:       record.StakerAid,
		StakerAddress:   record.StakerAddr,
		ActionID:        record.ActionId,
		TokenID:         record.TokenId,
		RoundNum:        record.RoundNum,
		TotalStakedNfts: record.NumStakedNFTs,
	}
}

func globalStakingRecordFromUnstake(
	record cgmodel.CGUnstakeActionInfoRec,
	withReward bool,
) cgstore.GlobalStakingActionRecord {
	result := cgstore.GlobalStakingActionRecord{
		Tx:              record.Tx,
		Kind:            cgstore.UserStakingActionUnstake,
		StakerAid:       record.StakerAid,
		StakerAddress:   record.StakerAddr,
		ActionID:        record.ActionId,
		TokenID:         record.TokenId,
		RoundNum:        record.RoundNum,
		TotalStakedNfts: record.NumStakedNFTs,
	}
	if withReward {
		result.RewardWei = record.RewardAmount
		result.RewardPerTokenWei = record.RewardPerToken
	}
	return result
}

func sameStakingLifecycle(stake, unstake cgstore.GlobalStakingActionRecord) error {
	if stake.ActionID != unstake.ActionID ||
		stake.TokenID != unstake.TokenID ||
		stake.StakerAid != unstake.StakerAid ||
		stake.StakerAddress != unstake.StakerAddress {
		return errors.New("stake and unstake rows describe different lifecycles")
	}
	return nil
}

func mapGlobalStakedCstToken(
	record cgstore.GlobalStakedCstTokenRecord,
) (GlobalStakedCstToken, error) {
	if record.ActionID < 0 || record.TokenID < 0 || record.MintRound < 0 {
		return GlobalStakedCstToken{}, errors.New("invalid globally staked CST token identity")
	}
	if record.Seed == "" {
		return GlobalStakedCstToken{}, errors.New("globally staked CST token misses its mint seed")
	}
	transaction, err := mapClaimTransaction(record.StakeTx)
	if err != nil {
		return GlobalStakedCstToken{}, fmt.Errorf("global CST stake transaction: %w", err)
	}
	address, err := canonicalNonZeroAddress("CST staker", record.StakerAddress)
	if err != nil {
		return GlobalStakedCstToken{}, err
	}
	result := GlobalStakedCstToken{
		ActionId:        record.ActionID,
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		MintRound:       record.MintRound,
		NftTokenId:      record.TokenID,
		Seed:            record.Seed,
		StakedAt:        transaction.OccurredAt,
		StakerAddress:   address,
		TransactionHash: transaction.TransactionHash,
	}
	if record.TokenName != "" {
		name := record.TokenName
		result.TokenName = &name
	}
	return result, nil
}

func mapGlobalStakedRandomWalkToken(
	record cgstore.GlobalStakedRwalkTokenRecord,
) (GlobalStakedRandomWalkToken, error) {
	if record.ActionID < 0 || record.TokenID < 0 {
		return GlobalStakedRandomWalkToken{}, errors.New("invalid globally staked RandomWalk token identity")
	}
	transaction, err := mapClaimTransaction(record.StakeTx)
	if err != nil {
		return GlobalStakedRandomWalkToken{}, fmt.Errorf("global RandomWalk stake transaction: %w", err)
	}
	address, err := canonicalNonZeroAddress("RandomWalk staker", record.StakerAddress)
	if err != nil {
		return GlobalStakedRandomWalkToken{}, err
	}
	return GlobalStakedRandomWalkToken{
		ActionId:        record.ActionID,
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		NftTokenId:      record.TokenID,
		StakedAt:        transaction.OccurredAt,
		StakerAddress:   address,
		TransactionHash: transaction.TransactionHash,
	}, nil
}

func mapGlobalStakingDeposit(
	record cgstore.GlobalStakingDepositRecord,
) (GlobalStakingDeposit, error) {
	if record.DepositID < 0 || record.RoundNum < 0 || record.TotalStakedNfts < 1 ||
		record.RewardCount < 0 || record.PendingRewardCount < 0 ||
		record.PendingRewardCount > record.RewardCount {
		return GlobalStakingDeposit{}, errors.New("invalid global staking deposit identity")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return GlobalStakingDeposit{}, fmt.Errorf("global staking deposit transaction: %w", err)
	}
	total, err := requiredAmount(record.TotalDepositWei)
	if err != nil {
		return GlobalStakingDeposit{}, fmt.Errorf("global staking deposit total: %w", err)
	}
	perToken, err := requiredAmount(record.AmountPerTokenWei)
	if err != nil {
		return GlobalStakingDeposit{}, fmt.Errorf("global staking deposit per-token amount: %w", err)
	}
	collected, err := requiredAmount(record.CollectedWei)
	if err != nil {
		return GlobalStakingDeposit{}, fmt.Errorf("global staking deposit collected amount: %w", err)
	}
	pending, err := requiredAmount(record.PendingWei)
	if err != nil {
		return GlobalStakingDeposit{}, fmt.Errorf("global staking deposit pending amount: %w", err)
	}
	remainder, err := requiredAmount(record.RemainderWei)
	if err != nil {
		return GlobalStakingDeposit{}, fmt.Errorf("global staking deposit remainder: %w", err)
	}
	if !amountEqualsSum(total, collected, pending, remainder) {
		return GlobalStakingDeposit{}, errors.New("global staking deposit amounts do not close")
	}
	return GlobalStakingDeposit{
		AmountPerTokenWei:  perToken,
		BlockNumber:        transaction.BlockNumber,
		CollectedWei:       collected,
		DepositId:          record.DepositID,
		EventLogId:         transaction.EventLogId,
		FullyClaimed:       record.PendingRewardCount == 0,
		OccurredAt:         transaction.OccurredAt,
		PendingRewardCount: record.PendingRewardCount,
		PendingWei:         pending,
		RemainderWei:       remainder,
		RewardCount:        record.RewardCount,
		Round:              record.RoundNum,
		TotalDepositWei:    total,
		TotalStakedNfts:    record.TotalStakedNfts,
		TransactionHash:    transaction.TransactionHash,
	}, nil
}

func mapRoundStakingReward(record cgstore.RoundStakingRewardRecord) (RoundStakingReward, error) {
	if record.DepositID < 0 || record.RoundNum < 0 || record.StakerAid < 1 ||
		record.StakedNftCount < 1 {
		return RoundStakingReward{}, errors.New("invalid round staking reward identity")
	}
	address, err := canonicalNonZeroAddress("staking reward staker", record.StakerAddress)
	if err != nil {
		return RoundStakingReward{}, err
	}
	reward, err := requiredAmount(record.RewardWei)
	if err != nil {
		return RoundStakingReward{}, fmt.Errorf("round staking reward total: %w", err)
	}
	collected, err := requiredAmount(record.CollectedWei)
	if err != nil {
		return RoundStakingReward{}, fmt.Errorf("round staking reward collected amount: %w", err)
	}
	pending, err := requiredAmount(record.PendingWei)
	if err != nil {
		return RoundStakingReward{}, fmt.Errorf("round staking reward pending amount: %w", err)
	}
	if !amountEqualsSum(reward, collected, pending) {
		return RoundStakingReward{}, errors.New("round staking reward amounts do not close")
	}
	return RoundStakingReward{
		CollectedWei:   collected,
		DepositId:      record.DepositID,
		FullyClaimed:   pending == "0",
		PendingWei:     pending,
		RewardWei:      reward,
		Round:          record.RoundNum,
		StakedNftCount: record.StakedNftCount,
		StakerAddress:  address,
	}, nil
}

func amountEqualsSum(total string, parts ...string) bool {
	totalInt, ok := new(big.Int).SetString(total, 10)
	if !ok {
		return false
	}
	sum := new(big.Int)
	for _, part := range parts {
		value, ok := new(big.Int).SetString(part, 10)
		if !ok {
			return false
		}
		sum.Add(sum, value)
	}
	return totalInt.Cmp(sum) == 0
}
