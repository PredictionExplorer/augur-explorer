// Prize events: main prize claims, raffle ETH/NFT prizes, endurance champion,
// chrono warrior and last-CST-bidder prizes, PrizesWallet ETH deposits/withdrawals,
// and the fund/ERC20 transfer-failure diagnostics.

package cosmicgame

import (
	"context"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func (h *Handlers) decodeMainPrizeClaimed(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGPrizeClaimEvent, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameMainPrizeClaimed
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "MainPrizeClaimed", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGPrizeClaimEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.RoundNum = lg.Topics[1].Big().Int64()
	evt.WinnerAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Amount = ethEvt.EthPrizeAmount.String()
	evt.CstAmount = ethEvt.CstPrizeAmount.String()
	evt.TokenId = lg.Topics[3].Big().Int64()
	evt.Timeout = ethEvt.TimeoutTimeToWithdrawSecondaryPrizes.Int64()
	return evt, nil
}

func (h *Handlers) storeMainPrizeClaimed(ctx context.Context, evt *cgmodel.CGPrizeClaimEvent) error {
	h.log.Info("MainPrizeClaimed",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "winner", evt.WinnerAddr,
		"amount", evt.Amount, "cst_amount", evt.CstAmount, "token_id", evt.TokenId,
		"withdraw_timeout", evt.Timeout)

	if err := h.repo.DeletePrizeClaim(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertPrizeClaim(ctx, evt)
}

func (h *Handlers) decodePrizesEthReceived(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGPrizesEthDeposit, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt cgc.IPrizesWalletEthReceived
	if err := h.prizesWalletABI.UnpackIntoInterface(&ethEvt, "EthReceived", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGPrizesEthDeposit{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Round = lg.Topics[1].Big().Int64()
	evt.WinnerAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.WinnerIndex = ethEvt.PrizeWinnerIndex.Int64()
	evt.Amount = ethEvt.Amount.String()
	return evt, nil
}

func (h *Handlers) storePrizesEthReceived(ctx context.Context, evt *cgmodel.CGPrizesEthDeposit) error {
	h.log.Info("Prizes EthReceived",
		"evt_id", evt.EvtId, "round", evt.Round, "winner", evt.WinnerAddr,
		"winner_index", evt.WinnerIndex, "amount", evt.Amount)

	if err := h.repo.DeletePrizeDeposit(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertPrizeDeposit(ctx, evt)
}

func (h *Handlers) decodePrizesEthWithdrawn(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGPrizesEthWithdrawal, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.IPrizesWalletEthWithdrawn
	if err := h.prizesWalletABI.UnpackIntoInterface(&ethEvt, "EthWithdrawn", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGPrizesEthWithdrawal{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Round = lg.Topics[1].Big().Int64()
	evt.WinnerAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.BeneficiaryAddr = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	evt.Amount = ethEvt.Amount.String()
	return evt, nil
}

func (h *Handlers) storePrizesEthWithdrawn(ctx context.Context, evt *cgmodel.CGPrizesEthWithdrawal) error {
	h.log.Info("Prizes EthWithdrawn",
		"evt_id", evt.EvtId, "round", evt.Round, "winner", evt.WinnerAddr,
		"beneficiary", evt.BeneficiaryAddr, "amount", evt.Amount)

	if err := h.repo.DeletePrizeWithdrawal(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertPrizeWithdrawal(ctx, evt)
}

func (h *Handlers) decodeRaffleWinnerPrizePaid(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGRaffleNFTWinner, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameRaffleWinnerPrizePaid
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "RaffleWinnerPrizePaid", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGRaffleNFTWinner{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Round = lg.Topics[1].Big().Int64()
	evt.TokenId = lg.Topics[3].Big().Int64()
	evt.WinnerIndex = ethEvt.WinnerIndex.Int64()
	evt.CstAmount = ethEvt.CstPrizeAmount.String()
	evt.IsRandomWalk = ethEvt.WinnerIsRandomWalkNftStaker
	evt.IsStaker = evt.IsRandomWalk
	return evt, nil
}

func (h *Handlers) storeRaffleWinnerPrizePaid(ctx context.Context, evt *cgmodel.CGRaffleNFTWinner) error {
	h.log.Info("RaffleWinnerPrizePaid",
		"evt_id", evt.EvtId, "round", evt.Round, "winner", evt.WinnerAddr,
		"token_id", evt.TokenId, "winner_index", evt.WinnerIndex,
		"cst_amount", evt.CstAmount, "is_staker", evt.IsStaker)

	if err := h.repo.DeleteRaffleNFTWinner(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertRaffleNFTWinner(ctx, evt)
}

func (h *Handlers) decodeRaffleEthAllocated(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGRaffleETHWinner, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameRaffleWinnerBidderEthPrizeAllocated
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "RaffleWinnerBidderEthPrizeAllocated", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGRaffleETHWinner{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Round = lg.Topics[1].Big().Int64()
	evt.WinnerIndex = ethEvt.WinnerIndex.Int64()
	evt.Amount = ethEvt.EthPrizeAmount.String()
	return evt, nil
}

func (h *Handlers) storeRaffleEthAllocated(ctx context.Context, evt *cgmodel.CGRaffleETHWinner) error {
	h.log.Info("RaffleWinnerBidderEthPrizeAllocated",
		"evt_id", evt.EvtId, "round", evt.Round, "winner", evt.WinnerAddr,
		"winner_index", evt.WinnerIndex, "amount", evt.Amount)

	if err := h.repo.DeleteRaffleETHWinner(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertRaffleETHWinner(ctx, evt)
}

func (h *Handlers) decodeEnduranceChampionPrizePaid(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGEnduranceWinner, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.ICosmicSignatureGameEnduranceChampionPrizePaid
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "EnduranceChampionPrizePaid", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGEnduranceWinner{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Round = lg.Topics[1].Big().Int64()
	evt.Erc721TokenId = lg.Topics[3].Big().Int64()
	evt.Erc20Amount = ethEvt.CstPrizeAmount.String()
	return evt, nil
}

func (h *Handlers) storeEnduranceChampionPrizePaid(ctx context.Context, evt *cgmodel.CGEnduranceWinner) error {
	h.log.Info("EnduranceChampionPrizePaid",
		"evt_id", evt.EvtId, "round", evt.Round, "champion", evt.WinnerAddr,
		"erc721_token_id", evt.Erc721TokenId, "erc20_amount", evt.Erc20Amount)

	if err := h.repo.DeleteEnduranceWinner(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertEnduranceWinner(ctx, evt)
}

func (h *Handlers) decodeLastCstBidderPrizePaid(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGLastBidderWinner, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameLastCstBidderPrizePaid
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "LastCstBidderPrizePaid", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGLastBidderWinner{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Round = lg.Topics[1].Big().Int64()
	evt.Erc721TokenId = lg.Topics[3].Big().Int64()
	evt.Erc20Amount = ethEvt.CstPrizeAmount.String()
	return evt, nil
}

func (h *Handlers) storeLastCstBidderPrizePaid(ctx context.Context, evt *cgmodel.CGLastBidderWinner) error {
	h.log.Info("LastCstBidderPrizePaid",
		"evt_id", evt.EvtId, "round", evt.Round, "winner", evt.WinnerAddr,
		"erc721_token_id", evt.Erc721TokenId, "erc20_amount", evt.Erc20Amount)

	if err := h.repo.DeleteLastCstBidderWinner(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertLastCstBidderWinner(ctx, evt)
}

func (h *Handlers) decodeChronoWarriorPrizePaid(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGChronoWarrior, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.ICosmicSignatureGameChronoWarriorPrizePaid
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "ChronoWarriorPrizePaid", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGChronoWarrior{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Round = lg.Topics[1].Big().Int64()
	evt.WinnerIndex = ethEvt.WinnerIndex.Int64()
	evt.EthAmount = ethEvt.EthPrizeAmount.String()
	evt.CstAmount = ethEvt.CstPrizeAmount.String()
	evt.NftId = lg.Topics[3].Big().Int64()
	return evt, nil
}

func (h *Handlers) storeChronoWarriorPrizePaid(ctx context.Context, evt *cgmodel.CGChronoWarrior) error {
	h.log.Info("ChronoWarriorPrizePaid",
		"evt_id", evt.EvtId, "round", evt.Round, "winner", evt.WinnerAddr,
		"winner_index", evt.WinnerIndex, "eth_amount", evt.EthAmount,
		"cst_amount", evt.CstAmount, "nft_id", evt.NftId)

	if err := h.repo.DeleteChronoWarrior(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertChronoWarrior(ctx, evt)
}

func (h *Handlers) decodeFundTransferFailed(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGFundTransferFailed, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameFundTransferFailed
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "FundTransferFailed", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGFundTransferFailed{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Destination = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Amount = ethEvt.Amount.String()
	return evt, nil
}

func (h *Handlers) storeFundTransferFailed(ctx context.Context, evt *cgmodel.CGFundTransferFailed) error {
	h.log.Info("FundTransferFailed", "evt_id", evt.EvtId, "destination", evt.Destination, "amount", evt.Amount)

	if err := h.repo.DeleteFundTransferFailed(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertFundTransferFailed(ctx, evt)
}

// decodeERC20TransferFailed handles ICosmicSignatureErrors.sol:
// ERC20TransferFailed(string errStr, address indexed destination, uint256
// amount), emitted by the game when a CST transfer cannot complete. The
// event is absent from every generated ABI, so the amount is decoded from
// the raw data words (see erc20TransferFailedAmount).
func (h *Handlers) decodeERC20TransferFailed(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGErc20TransferFailed, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	amount, err := erc20TransferFailedAmount(lg.Data)
	if err != nil {
		return nil, err
	}

	evt := &cgmodel.CGErc20TransferFailed{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Destination = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Amount = amount.String()
	return evt, nil
}

func (h *Handlers) storeERC20TransferFailed(ctx context.Context, evt *cgmodel.CGErc20TransferFailed) error {
	h.log.Info("ERC20TransferFailed", "evt_id", evt.EvtId, "destination", evt.Destination, "amount", evt.Amount)

	if err := h.repo.DeleteERC20TransferFailed(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertERC20TransferFailed(ctx, evt)
}
