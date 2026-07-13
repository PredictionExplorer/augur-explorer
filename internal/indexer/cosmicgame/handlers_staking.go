// Staking events for the CST and RandomWalk staking wallets: stake/unstake
// actions and staking-reward ETH deposits.

package cosmicgame

import (
	"context"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func (h *Handlers) decodeNftStakedCST(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGNftStakedCst, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.IStakingWalletCosmicSignatureNftNftUnstaked
	if err := h.stakingCSTABI.UnpackIntoInterface(&ethEvt, "NftStaked", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGNftStakedCst{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = lg.Topics[1].Big().Int64()
	evt.NftId = lg.Topics[2].Big().Int64()
	evt.StakerAddress = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	evt.NumStakedNfts = ethEvt.NumStakedNfts.Int64()
	evt.RewardPerStaker = ethEvt.RewardAmountPerStakedNft.String()
	return evt, nil
}

func (h *Handlers) storeNftStakedCST(ctx context.Context, evt *cgmodel.CGNftStakedCst) error {
	h.log.Info("CST NftStaked",
		"evt_id", evt.EvtId, "action_id", evt.ActionId, "nft_id", evt.NftId,
		"staker", evt.StakerAddress, "num_staked", evt.NumStakedNfts,
		"reward_per_staker", evt.RewardPerStaker)

	if err := h.repo.DeleteNftStakedCST(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertNftStakedCST(ctx, evt)
}

func (h *Handlers) decodeNftStakedRWalk(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGNftStakedRWalk, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.IStakingWalletRandomWalkNftNftStaked
	if err := h.stakingRWalkABI.UnpackIntoInterface(&ethEvt, "NftStaked", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGNftStakedRWalk{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = lg.Topics[1].Big().Int64()
	evt.NftId = lg.Topics[2].Big().Int64()
	evt.StakerAddress = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	evt.NumStakedNfts = ethEvt.NumStakedNfts.Int64()
	return evt, nil
}

func (h *Handlers) storeNftStakedRWalk(ctx context.Context, evt *cgmodel.CGNftStakedRWalk) error {
	h.log.Info("RWalk NftStaked",
		"evt_id", evt.EvtId, "action_id", evt.ActionId, "nft_id", evt.NftId,
		"staker", evt.StakerAddress, "num_staked", evt.NumStakedNfts)

	if err := h.repo.DeleteNftStakedRWalk(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertNftStakedRWalk(ctx, evt)
}

func (h *Handlers) decodeStakingEthDeposit(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGEthDeposit, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.IStakingWalletCosmicSignatureNftEthDepositReceived
	if err := h.stakingCSTABI.UnpackIntoInterface(&ethEvt, "EthDepositReceived", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGEthDeposit{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DepositTime = elog.TimeStamp
	evt.DepositId = ethEvt.ActionCounter.Int64()
	evt.NumStakedNfts = ethEvt.NumStakedNfts.Int64()
	evt.Amount = ethEvt.DepositAmount.String()
	evt.AccumModulo = "0" // pending for resolution regarding StakingWalletCST refactoring
	evt.RoundNum = lg.Topics[1].Big().Int64()
	divres := big.NewInt(0)
	rem := big.NewInt(0)
	divres.QuoRem(ethEvt.DepositAmount, ethEvt.NumStakedNfts, rem)
	evt.AmountPerStaker = divres.String()
	evt.Modulo = rem.String()
	return evt, nil
}

func (h *Handlers) storeStakingEthDeposit(ctx context.Context, evt *cgmodel.CGEthDeposit) error {
	h.log.Info("EthDepositReceived",
		"evt_id", evt.EvtId, "deposit_id", evt.DepositId, "round", evt.RoundNum,
		"num_staked", evt.NumStakedNfts, "amount", evt.Amount,
		"amount_per_staker", evt.AmountPerStaker, "modulo", evt.Modulo)

	if err := h.repo.DeleteStakingEthDeposit(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertStakingEthDeposit(ctx, evt)
}

func (h *Handlers) decodeNftUnstakedRWalk(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGNftUnstakedRWalk, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.IStakingWalletRandomWalkNftNftUnstaked
	if err := h.stakingRWalkABI.UnpackIntoInterface(&ethEvt, "NftUnstaked", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGNftUnstakedRWalk{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = lg.Topics[1].Big().Int64()
	evt.NftId = lg.Topics[2].Big().Int64()
	evt.StakerAddress = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	evt.NumStakedNfts = ethEvt.NumStakedNfts.Int64()
	return evt, nil
}

func (h *Handlers) storeNftUnstakedRWalk(ctx context.Context, evt *cgmodel.CGNftUnstakedRWalk) error {
	h.log.Info("RWalk NftUnstaked",
		"evt_id", evt.EvtId, "action_id", evt.ActionId, "nft_id", evt.NftId,
		"staker", evt.StakerAddress, "num_staked", evt.NumStakedNfts)

	if err := h.repo.DeleteNftUnstakedRWalk(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertNftUnstakedRWalk(ctx, evt)
}

func (h *Handlers) decodeNftUnstakedCST(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGNftUnstakedCst, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.IStakingWalletCosmicSignatureNftNftUnstaked
	if err := h.stakingCSTABI.UnpackIntoInterface(&ethEvt, "NftUnstaked", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGNftUnstakedCst{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = lg.Topics[1].Big().Int64()
	evt.NftId = lg.Topics[2].Big().Int64()
	evt.StakerAddress = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	evt.NumStakedNfts = ethEvt.NumStakedNfts.Int64()
	evt.RewardAmount = ethEvt.RewardAmount.String()
	evt.RewardPerToken = ethEvt.RewardAmountPerStakedNft.String()
	evt.ActionCounter = ethEvt.ActionCounter.Int64()
	return evt, nil
}

func (h *Handlers) storeNftUnstakedCST(ctx context.Context, evt *cgmodel.CGNftUnstakedCst) error {
	h.log.Info("CST NftUnstaked",
		"evt_id", evt.EvtId, "action_id", evt.ActionId, "nft_id", evt.NftId,
		"staker", evt.StakerAddress, "num_staked", evt.NumStakedNfts,
		"reward_amount", evt.RewardAmount, "reward_per_token", evt.RewardPerToken,
		"action_counter", evt.ActionCounter)

	if err := h.repo.DeleteNftUnstakedCST(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertNftUnstakedCST(ctx, evt)
}
