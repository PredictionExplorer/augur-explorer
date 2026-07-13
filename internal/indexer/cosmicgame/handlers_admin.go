// Admin and parameter-change events: percentages, durations, divisors, address
// changes, ownership transfers, proxy upgrades and related system-management events.

package cosmicgame

import (
	"context"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// adminEventBase fills the fields every admin event row shares.
func adminEventBase(lg *types.Log, elog *store.EthereumEventLog) (evtID, blockNum, txID, timeStamp int64, contract string) {
	return elog.EvtId, elog.BlockNum, elog.TxId, elog.TimeStamp, lg.Address.String()
}

// --- CST-bid-reward changes: three event types share cg_adm_erc20_reward ---

// storeCstRewardForBiddingChange persists a CST-bid-reward change; three
// event types (CstRewardAmountForBiddingChanged, BidCstRewardAmountChanged,
// BidCstRewardAmountMultiplierChanged) share the cg_adm_erc20_reward table.
// label names the concrete event in the log record.
func (h *Handlers) storeCstRewardForBiddingChange(label string) func(context.Context, *cgmodel.CGCstRewardForBiddingChanged) error {
	return func(ctx context.Context, evt *cgmodel.CGCstRewardForBiddingChanged) error {
		h.log.Info(label, "evt_id", evt.EvtId, "new_reward", evt.NewReward)

		if err := h.repo.DeleteCstRewardForBiddingChange(ctx, evt.EvtId); err != nil {
			return err
		}
		return h.repo.InsertCstRewardForBiddingChange(ctx, evt)
	}
}

func (h *Handlers) decodeCstRewardAmountForBiddingChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCstRewardForBiddingChanged, error) {
	var ethEvt cgc.CosmicSignatureGameCstRewardAmountForBiddingChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "CstRewardAmountForBiddingChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCstRewardForBiddingChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewReward = ethEvt.NewValue.String()
	return evt, nil
}

// decodeBidCstRewardAmountChanged decodes BidCstRewardAmountChanged(uint256),
// a legacy contract event no current ABI defines: the single non-indexed
// uint256 comes from the raw data words (unpacking a name absent from the
// ABI made the legacy handler terminate the process on every occurrence).
func (h *Handlers) decodeBidCstRewardAmountChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCstRewardForBiddingChanged, error) {
	newValue, err := adminUint256FromLogData(lg.Data)
	if err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCstRewardForBiddingChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewReward = newValue.String()
	return evt, nil
}

func (h *Handlers) decodeBidCstRewardAmountMultiplierChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCstRewardForBiddingChanged, error) {
	var ethEvt cgc.CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged
	if err := h.gameV2ABI.UnpackIntoInterface(&ethEvt, "BidCstRewardAmountMultiplierChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCstRewardForBiddingChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewReward = ethEvt.NewValue.String()
	return evt, nil
}

// --- CST auction-length changes: V1 divisor and V2 duration share cg_adm_cst_auclen ---

// storeCstAuctionLengthChange persists a CST auction-length change; the V1
// divisor and V2 duration events share the cg_adm_cst_auclen table.
func (h *Handlers) storeCstAuctionLengthChange(label string) func(context.Context, *cgmodel.CGCstDutchAuctionDurationDivisorChanged) error {
	return func(ctx context.Context, evt *cgmodel.CGCstDutchAuctionDurationDivisorChanged) error {
		h.log.Info(label, "evt_id", evt.EvtId, "new_value", evt.NewValue)

		if err := h.repo.DeleteCstAuctionLengthChange(ctx, evt.EvtId); err != nil {
			return err
		}
		return h.repo.InsertCstAuctionLengthChange(ctx, evt)
	}
}

func (h *Handlers) decodeCstDutchAuctionDurationDivisorChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCstDutchAuctionDurationDivisorChanged, error) {
	var ethEvt cgc.CosmicSignatureGameCstDutchAuctionDurationDivisorChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "CstDutchAuctionDurationDivisorChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCstDutchAuctionDurationDivisorChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewValue = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) decodeCstDutchAuctionDurationChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCstDutchAuctionDurationDivisorChanged, error) {
	var ethEvt cgc.CosmicSignatureGameV2CstDutchAuctionDurationChanged
	if err := h.gameV2ABI.UnpackIntoInterface(&ethEvt, "CstDutchAuctionDurationChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCstDutchAuctionDurationDivisorChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewValue = ethEvt.NewValue.String()
	return evt, nil
}

// --- CharityAddressChanged: one topic0, two contracts, two meanings ---

// decodeCharityReceiverChanged handles the CharityWallet's event: it sets
// who receives charity funds.
func (h *Handlers) decodeCharityReceiverChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCharityUpdatedEvent, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CharityWalletCharityAddressChanged
	if err := h.charityWalletABI.UnpackIntoInterface(&ethEvt, "CharityAddressChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCharityUpdatedEvent{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCharityAddr = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storeCharityReceiverChanged(ctx context.Context, evt *cgmodel.CGCharityUpdatedEvent) error {
	h.log.Info("CharityAddressChanged (CharityWallet receiver)", "evt_id", evt.EvtId, "new_charity", evt.NewCharityAddr)

	if err := h.repo.DeleteCharityReceiverChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertCharityReceiverChange(ctx, evt)
}

// decodeCharityWalletChanged handles the game's event of the same signature:
// it sets which CharityWallet contract the game uses.
func (h *Handlers) decodeCharityWalletChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCharityAddressChanged, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameCharityAddressChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "CharityAddressChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCharityAddressChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewCharity = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storeCharityWalletChanged(ctx context.Context, evt *cgmodel.CGCharityAddressChanged) error {
	h.log.Info("CharityAddressChanged (game wallet)", "evt_id", evt.EvtId, "new_wallet", evt.NewCharity)

	if err := h.repo.DeleteCharityWalletAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertCharityWalletAddressChange(ctx, evt)
}

// --- Percentage changes on the game contract ---

func (h *Handlers) decodeCharityPercentageChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCharityPercentageChanged, error) {
	var ethEvt cgc.CosmicSignatureGameCharityEthDonationAmountPercentageChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "CharityEthDonationAmountPercentageChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCharityPercentageChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewCharityPercentage = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeCharityPercentageChanged(ctx context.Context, evt *cgmodel.CGCharityPercentageChanged) error {
	h.log.Info("CharityEthDonationAmountPercentageChanged", "evt_id", evt.EvtId, "new_percentage", evt.NewCharityPercentage)

	if err := h.repo.DeleteCharityPercentageChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertCharityPercentageChange(ctx, evt)
}

func (h *Handlers) decodePrizePercentageChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGPrizePercentageChanged, error) {
	var ethEvt cgc.CosmicSignatureGameMainEthPrizeAmountPercentageChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "MainEthPrizeAmountPercentageChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGPrizePercentageChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewPrizePercentage = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storePrizePercentageChanged(ctx context.Context, evt *cgmodel.CGPrizePercentageChanged) error {
	h.log.Info("MainEthPrizeAmountPercentageChanged", "evt_id", evt.EvtId, "new_percentage", evt.NewPrizePercentage)

	if err := h.repo.DeletePrizePercentageChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertPrizePercentageChange(ctx, evt)
}

func (h *Handlers) decodeRafflePercentageChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGRafflePercentageChanged, error) {
	var ethEvt cgc.CosmicSignatureGameRaffleTotalEthPrizeAmountForBiddersPercentageChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGRafflePercentageChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewRafflePercentage = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeRafflePercentageChanged(ctx context.Context, evt *cgmodel.CGRafflePercentageChanged) error {
	h.log.Info("RaffleTotalEthPrizeAmountForBiddersPercentageChanged", "evt_id", evt.EvtId, "new_percentage", evt.NewRafflePercentage)

	if err := h.repo.DeleteRafflePercentageChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertRafflePercentageChange(ctx, evt)
}

func (h *Handlers) decodeStakingPercentageChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGStakingPercentageChanged, error) {
	var ethEvt cgc.CosmicSignatureGameCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGStakingPercentageChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewStakingPercentage = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeStakingPercentageChanged(ctx context.Context, evt *cgmodel.CGStakingPercentageChanged) error {
	h.log.Info("CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", "evt_id", evt.EvtId, "new_percentage", evt.NewStakingPercentage)

	if err := h.repo.DeleteStakingPercentageChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertStakingPercentageChange(ctx, evt)
}

func (h *Handlers) decodeChronoPercentageChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGChronoPercentageChanged, error) {
	var ethEvt cgc.ISystemEventsChronoWarriorEthPrizeAmountPercentageChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "ChronoWarriorEthPrizeAmountPercentageChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGChronoPercentageChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewChronoPercentage = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeChronoPercentageChanged(ctx context.Context, evt *cgmodel.CGChronoPercentageChanged) error {
	h.log.Info("ChronoWarriorEthPrizeAmountPercentageChanged", "evt_id", evt.EvtId, "new_percentage", evt.NewChronoPercentage)

	if err := h.repo.DeleteChronoPercentageChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertChronoPercentageChange(ctx, evt)
}

// --- Raffle winner-count changes ---

func (h *Handlers) decodeNumRaffleETHWinnersBiddingChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGNumRaffleETHWinnersBiddingChanged, error) {
	var ethEvt cgc.CosmicSignatureGameNumRaffleEthPrizesForBiddersChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "NumRaffleEthPrizesForBiddersChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGNumRaffleETHWinnersBiddingChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewNumRaffleETHWinnersBidding = ethEvt.NewValue.Int64()
	return evt, nil
}

func (h *Handlers) storeNumRaffleETHWinnersBiddingChanged(ctx context.Context, evt *cgmodel.CGNumRaffleETHWinnersBiddingChanged) error {
	h.log.Info("NumRaffleEthPrizesForBiddersChanged", "evt_id", evt.EvtId, "new_value", evt.NewNumRaffleETHWinnersBidding)

	if err := h.repo.DeleteNumRaffleETHWinnersBiddingChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertNumRaffleETHWinnersBiddingChange(ctx, evt)
}

func (h *Handlers) decodeNumRaffleNFTWinnersBiddingChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGNumRaffleNFTWinnersBiddingChanged, error) {
	var ethEvt cgc.ISystemManagementNumRaffleCosmicSignatureNftsForBiddersChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "NumRaffleCosmicSignatureNftsForBiddersChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGNumRaffleNFTWinnersBiddingChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewNumRaffleNFTWinnersBidding = ethEvt.NewValue.Int64()
	return evt, nil
}

func (h *Handlers) storeNumRaffleNFTWinnersBiddingChanged(ctx context.Context, evt *cgmodel.CGNumRaffleNFTWinnersBiddingChanged) error {
	h.log.Info("NumRaffleCosmicSignatureNftsForBiddersChanged", "evt_id", evt.EvtId, "new_value", evt.NewNumRaffleNFTWinnersBidding)

	if err := h.repo.DeleteNumRaffleNFTWinnersBiddingChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertNumRaffleNFTWinnersBiddingChange(ctx, evt)
}

func (h *Handlers) decodeNumRaffleNFTWinnersStakingRWalkChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGNumRaffleNFTWinnersStakingRWalkChanged, error) {
	var ethEvt cgc.ISystemManagementNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGNumRaffleNFTWinnersStakingRWalkChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewNumRaffleNFTWinnersStakingRWalk = ethEvt.NewValue.Int64()
	return evt, nil
}

func (h *Handlers) storeNumRaffleNFTWinnersStakingRWalkChanged(ctx context.Context, evt *cgmodel.CGNumRaffleNFTWinnersStakingRWalkChanged) error {
	h.log.Info("NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", "evt_id", evt.EvtId, "new_value", evt.NewNumRaffleNFTWinnersStakingRWalk)

	if err := h.repo.DeleteNumRaffleNFTWinnersStakingRWalkChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertNumRaffleNFTWinnersStakingRWalkChange(ctx, evt)
}

// --- Platform-address changes (indexed address, empty data) ---

func (h *Handlers) decodeRandomWalkAddressChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGRandomWalkAddressChanged, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGRandomWalkAddressChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewRandomWalk = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storeRandomWalkAddressChanged(ctx context.Context, evt *cgmodel.CGRandomWalkAddressChanged) error {
	h.log.Info("RandomWalkNftAddressChanged", "evt_id", evt.EvtId, "new_address", evt.NewRandomWalk)

	if err := h.repo.DeleteRandomWalkAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertRandomWalkAddressChange(ctx, evt)
}

func (h *Handlers) decodePrizesWalletAddressChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGPrizeWalletAddressChanged, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGPrizeWalletAddressChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewPrizeWallet = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storePrizesWalletAddressChanged(ctx context.Context, evt *cgmodel.CGPrizeWalletAddressChanged) error {
	h.log.Info("PrizesWalletAddressChanged", "evt_id", evt.EvtId, "new_address", evt.NewPrizeWallet)

	if err := h.repo.DeletePrizesWalletAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertPrizesWalletAddressChange(ctx, evt)
}

func (h *Handlers) decodeStakingWalletCSTAddressChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGStakingWalletCSTAddressChanged, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGStakingWalletCSTAddressChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewStakingWalletCST = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storeStakingWalletCSTAddressChanged(ctx context.Context, evt *cgmodel.CGStakingWalletCSTAddressChanged) error {
	h.log.Info("StakingWalletCosmicSignatureNftAddressChanged", "evt_id", evt.EvtId, "new_address", evt.NewStakingWalletCST)

	if err := h.repo.DeleteStakingWalletCSTAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertStakingWalletCSTAddressChange(ctx, evt)
}

func (h *Handlers) decodeStakingWalletRWalkAddressChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGStakingWalletRWalkAddressChanged, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGStakingWalletRWalkAddressChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewStakingWalletRWalk = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storeStakingWalletRWalkAddressChanged(ctx context.Context, evt *cgmodel.CGStakingWalletRWalkAddressChanged) error {
	h.log.Info("StakingWalletRandomWalkNftAddressChanged", "evt_id", evt.EvtId, "new_address", evt.NewStakingWalletRWalk)

	if err := h.repo.DeleteStakingWalletRWalkAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertStakingWalletRWalkAddressChange(ctx, evt)
}

func (h *Handlers) decodeMarketingWalletAddressChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGMarketingWalletAddressChanged, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGMarketingWalletAddressChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewMarketingWallet = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storeMarketingWalletAddressChanged(ctx context.Context, evt *cgmodel.CGMarketingWalletAddressChanged) error {
	h.log.Info("MarketingWalletAddressChanged", "evt_id", evt.EvtId, "new_address", evt.NewMarketingWallet)

	if err := h.repo.DeleteMarketingWalletAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertMarketingWalletAddressChange(ctx, evt)
}

func (h *Handlers) decodeTreasurerAddressChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGTreasurerAddressChanged, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGTreasurerAddressChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewTreasurer = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storeTreasurerAddressChanged(ctx context.Context, evt *cgmodel.CGTreasurerAddressChanged) error {
	h.log.Info("TreasurerAddressChanged", "evt_id", evt.EvtId, "new_treasurer", evt.NewTreasurer)

	if err := h.repo.DeleteTreasurerAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertTreasurerAddressChange(ctx, evt)
}

func (h *Handlers) decodeCosmicTokenAddressChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCosmicTokenAddressChanged, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameCosmicSignatureTokenAddressChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "CosmicSignatureTokenAddressChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCosmicTokenAddressChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewCosmicToken = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storeCosmicTokenAddressChanged(ctx context.Context, evt *cgmodel.CGCosmicTokenAddressChanged) error {
	h.log.Info("CosmicSignatureTokenAddressChanged", "evt_id", evt.EvtId, "new_address", evt.NewCosmicToken)

	if err := h.repo.DeleteCosmicTokenAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertCosmicTokenAddressChange(ctx, evt)
}

func (h *Handlers) decodeCosmicSignatureAddressChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCosmicSignatureAddressChanged, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCosmicSignatureAddressChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewCosmicSignature = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storeCosmicSignatureAddressChanged(ctx context.Context, evt *cgmodel.CGCosmicSignatureAddressChanged) error {
	h.log.Info("CosmicSignatureNftAddressChanged", "evt_id", evt.EvtId, "new_address", evt.NewCosmicSignature)

	if err := h.repo.DeleteCosmicSignatureAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertCosmicSignatureAddressChange(ctx, evt)
}

// --- Proxy lifecycle ---

func (h *Handlers) decodeUpgraded(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGUpgraded, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameUpgraded
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "Upgraded", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGUpgraded{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.Implementation = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storeUpgraded(ctx context.Context, evt *cgmodel.CGUpgraded) error {
	h.log.Info("Upgraded", "evt_id", evt.EvtId, "implementation", evt.Implementation)

	if err := h.repo.DeleteUpgraded(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertUpgraded(ctx, evt)
}

// decodeAdminChanged decodes the ERC-1967 proxy event; it is absent from the
// game ABI, so it must be decoded with the IERC1967 ABI (using the game ABI
// here made the legacy handler terminate the process on every AdminChanged
// event).
func (h *Handlers) decodeAdminChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGAdminChanged, error) {
	var ethEvt cgc.IERC1967AdminChanged
	if err := h.erc1967ABI.UnpackIntoInterface(&ethEvt, "AdminChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGAdminChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.OldAdmin = ethEvt.PreviousAdmin.String()
	evt.NewAdmin = ethEvt.NewAdmin.String()
	return evt, nil
}

func (h *Handlers) storeAdminChanged(ctx context.Context, evt *cgmodel.CGAdminChanged) error {
	h.log.Info("AdminChanged", "evt_id", evt.EvtId, "old_admin", evt.OldAdmin, "new_admin", evt.NewAdmin)

	if err := h.repo.DeleteAdminChanged(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertAdminChanged(ctx, evt)
}

// --- Timing and pricing parameters ---

// decodeTimeIncreaseChanged decodes TimeIncreaseChanged(uint256), a legacy
// contract event no current ABI defines: the single non-indexed uint256
// comes from the raw data (unpacking a name absent from the ABI made the
// legacy handler terminate the process on every occurrence).
func (h *Handlers) decodeTimeIncreaseChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGTimeIncreaseChanged, error) {
	newValue, err := adminUint256FromLogData(lg.Data)
	if err != nil {
		return nil, err
	}
	evt := &cgmodel.CGTimeIncreaseChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewTimeIncrease = newValue.String()
	return evt, nil
}

func (h *Handlers) storeTimeIncreaseChanged(ctx context.Context, evt *cgmodel.CGTimeIncreaseChanged) error {
	h.log.Info("MainPrizeTimeIncrementIncreaseDivisorChanged", "evt_id", evt.EvtId, "new_value", evt.NewTimeIncrease)

	if err := h.repo.DeleteTimeIncreaseChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertTimeIncreaseChange(ctx, evt)
}

func (h *Handlers) decodeTimeoutClaimPrizeChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGTimeoutClaimPrizeChanged, error) {
	var ethEvt cgc.CosmicSignatureGameTimeoutDurationToClaimMainPrizeChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "TimeoutDurationToClaimMainPrizeChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGTimeoutClaimPrizeChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewTimeout = ethEvt.NewValue.Int64()
	return evt, nil
}

func (h *Handlers) storeTimeoutClaimPrizeChanged(ctx context.Context, evt *cgmodel.CGTimeoutClaimPrizeChanged) error {
	h.log.Info("TimeoutDurationToClaimMainPrizeChanged", "evt_id", evt.EvtId, "new_timeout", evt.NewTimeout)

	if err := h.repo.DeleteTimeoutClaimPrizeChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertTimeoutClaimPrizeChange(ctx, evt)
}

func (h *Handlers) decodeTimeoutToWithdrawPrizesChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGTimeoutToWithdrawPrizeChanged, error) {
	var ethEvt cgc.IPrizesWalletTimeoutDurationToWithdrawPrizesChanged
	if err := h.prizesWalletABI.UnpackIntoInterface(&ethEvt, "TimeoutDurationToWithdrawPrizesChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGTimeoutToWithdrawPrizeChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewTimeout = ethEvt.NewValue.Int64()
	return evt, nil
}

func (h *Handlers) storeTimeoutToWithdrawPrizesChanged(ctx context.Context, evt *cgmodel.CGTimeoutToWithdrawPrizeChanged) error {
	h.log.Info("TimeoutDurationToWithdrawPrizesChanged", "evt_id", evt.EvtId, "new_timeout", evt.NewTimeout)

	if err := h.repo.DeleteTimeoutToWithdrawPrizesChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertTimeoutToWithdrawPrizesChange(ctx, evt)
}

func (h *Handlers) decodePriceIncreaseChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGPriceIncreaseChanged, error) {
	var ethEvt cgc.CosmicSignatureGameEthBidPriceIncreaseDivisorChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "EthBidPriceIncreaseDivisorChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGPriceIncreaseChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewPriceIncrease = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storePriceIncreaseChanged(ctx context.Context, evt *cgmodel.CGPriceIncreaseChanged) error {
	h.log.Info("EthBidPriceIncreaseDivisorChanged", "evt_id", evt.EvtId, "new_value", evt.NewPriceIncrease)

	if err := h.repo.DeletePriceIncreaseChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertPriceIncreaseChange(ctx, evt)
}

func (h *Handlers) decodeMainPrizeMicrosecondsChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGMainPrizeMicroSecondsIncreaseChanged, error) {
	var ethEvt cgc.CosmicSignatureGameMainPrizeTimeIncrementInMicroSecondsChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "MainPrizeTimeIncrementInMicroSecondsChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGMainPrizeMicroSecondsIncreaseChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewMicroseconds = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeMainPrizeMicrosecondsChanged(ctx context.Context, evt *cgmodel.CGMainPrizeMicroSecondsIncreaseChanged) error {
	h.log.Info("MainPrizeTimeIncrementInMicroSecondsChanged", "evt_id", evt.EvtId, "new_value", evt.NewMicroseconds)

	if err := h.repo.DeleteMainPrizeMicrosecondsChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertMainPrizeMicrosecondsChange(ctx, evt)
}

func (h *Handlers) decodeInitialSecondsUntilPrizeChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGInitialSecondsUntilPrizeChanged, error) {
	var ethEvt cgc.CosmicSignatureGameInitialDurationUntilMainPrizeDivisorChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "InitialDurationUntilMainPrizeDivisorChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGInitialSecondsUntilPrizeChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewInitialSecondsUntilPrize = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeInitialSecondsUntilPrizeChanged(ctx context.Context, evt *cgmodel.CGInitialSecondsUntilPrizeChanged) error {
	h.log.Info("InitialDurationUntilMainPrizeDivisorChanged", "evt_id", evt.EvtId, "new_value", evt.NewInitialSecondsUntilPrize)

	if err := h.repo.DeleteInitialSecondsUntilPrizeChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertInitialSecondsUntilPrizeChange(ctx, evt)
}

func (h *Handlers) decodeRoundActivationTimeChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGActivationTimeChanged, error) {
	var ethEvt cgc.BiddingBaseRoundActivationTimeChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "RoundActivationTimeChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGActivationTimeChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewActivationTime = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeRoundActivationTimeChanged(ctx context.Context, evt *cgmodel.CGActivationTimeChanged) error {
	h.log.Info("RoundActivationTimeChanged", "evt_id", evt.EvtId, "new_activation_time", evt.NewActivationTime)

	if err := h.repo.DeleteActivationTimeChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertActivationTimeChange(ctx, evt)
}

func (h *Handlers) decodeCstAuctionDurationChangeDivisorChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCstDutchAuctionDurationChangeDivisorChanged, error) {
	var ethEvt cgc.CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged
	if err := h.gameV2ABI.UnpackIntoInterface(&ethEvt, "CstDutchAuctionDurationChangeDivisorChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCstDutchAuctionDurationChangeDivisorChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewValue = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeCstAuctionDurationChangeDivisorChanged(ctx context.Context, evt *cgmodel.CGCstDutchAuctionDurationChangeDivisorChanged) error {
	h.log.Info("CstDutchAuctionDurationChangeDivisorChanged", "evt_id", evt.EvtId, "new_divisor", evt.NewValue)

	if err := h.repo.DeleteCstAuctionDurationChangeDivisorChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertCstAuctionDurationChangeDivisorChange(ctx, evt)
}

func (h *Handlers) decodeEthAuctionDurationDivisorChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGEthDutchAuctionDurationDivisorChanged, error) {
	var ethEvt cgc.CosmicSignatureGameEthDutchAuctionDurationDivisorChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "EthDutchAuctionDurationDivisorChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGEthDutchAuctionDurationDivisorChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewValue = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeEthAuctionDurationDivisorChanged(ctx context.Context, evt *cgmodel.CGEthDutchAuctionDurationDivisorChanged) error {
	h.log.Info("EthDutchAuctionDurationDivisorChanged", "evt_id", evt.EvtId, "new_divisor", evt.NewValue)

	if err := h.repo.DeleteEthAuctionDurationDivisorChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertEthAuctionDurationDivisorChange(ctx, evt)
}

func (h *Handlers) decodeEthAuctionEndingBidPriceDivisorChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGEthDutchAuctionEndingBidPriceDivisorChanged, error) {
	var ethEvt cgc.CosmicSignatureGameEthDutchAuctionEndingBidPriceDivisorChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "EthDutchAuctionEndingBidPriceDivisorChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGEthDutchAuctionEndingBidPriceDivisorChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewValue = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeEthAuctionEndingBidPriceDivisorChanged(ctx context.Context, evt *cgmodel.CGEthDutchAuctionEndingBidPriceDivisorChanged) error {
	h.log.Info("EthDutchAuctionEndingBidPriceDivisorChanged", "evt_id", evt.EvtId, "new_divisor", evt.NewValue)

	if err := h.repo.DeleteEthAuctionEndingBidPriceDivisorChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertEthAuctionEndingBidPriceDivisorChange(ctx, evt)
}

// --- Marketing / message / URI parameters ---

func (h *Handlers) decodeMarketingRewardChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGMarketingRewardChanged, error) {
	var ethEvt cgc.CosmicSignatureGameMarketingWalletCstContributionAmountChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "MarketingWalletCstContributionAmountChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGMarketingRewardChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewReward = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeMarketingRewardChanged(ctx context.Context, evt *cgmodel.CGMarketingRewardChanged) error {
	h.log.Info("MarketingWalletCstContributionAmountChanged", "evt_id", evt.EvtId, "new_reward", evt.NewReward)

	if err := h.repo.DeleteMarketingRewardChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertMarketingRewardChange(ctx, evt)
}

func (h *Handlers) decodeStaticCstRewardChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGStaticCstReward, error) {
	var ethEvt cgc.BiddingCstPrizeAmountChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "CstPrizeAmountChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGStaticCstReward{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewReward = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeStaticCstRewardChanged(ctx context.Context, evt *cgmodel.CGStaticCstReward) error {
	h.log.Info("CstPrizeAmountChanged", "evt_id", evt.EvtId, "new_reward", evt.NewReward)

	if err := h.repo.DeleteStaticCstRewardChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertStaticCstRewardChange(ctx, evt)
}

func (h *Handlers) decodeMaxMessageLengthChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGMaxMessageLengthChanged, error) {
	var ethEvt cgc.CosmicSignatureGameBidMessageLengthMaxLimitChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "BidMessageLengthMaxLimitChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGMaxMessageLengthChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewMessageLength = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeMaxMessageLengthChanged(ctx context.Context, evt *cgmodel.CGMaxMessageLengthChanged) error {
	h.log.Info("BidMessageLengthMaxLimitChanged", "evt_id", evt.EvtId, "new_length", evt.NewMessageLength)

	if err := h.repo.DeleteMaxMessageLengthChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertMaxMessageLengthChange(ctx, evt)
}

func (h *Handlers) decodeNftGenerationScriptURLChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGTokenGenerationScriptURL, error) {
	var ethEvt cgc.ICosmicSignatureNftNftGenerationScriptUriChanged
	if err := h.signatureABI.UnpackIntoInterface(&ethEvt, "NftGenerationScriptUriChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGTokenGenerationScriptURL{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewURL = ethEvt.NewValue
	return evt, nil
}

func (h *Handlers) storeNftGenerationScriptURLChanged(ctx context.Context, evt *cgmodel.CGTokenGenerationScriptURL) error {
	h.log.Info("NftGenerationScriptUriChanged", "evt_id", evt.EvtId, "new_url", evt.NewURL)

	// Must delete from cg_adm_script_url (this event's own table); deleting
	// from the message-length table here made every re-processed script-URL
	// event abort on the cg_adm_script_url unique constraint.
	if err := h.repo.DeleteTokenGenerationScriptURL(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertTokenGenerationScriptURL(ctx, evt)
}

func (h *Handlers) decodeNftBaseURIChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGBaseURIEvent, error) {
	var ethEvt cgc.CosmicSignatureNftNftBaseUriChanged
	if err := h.signatureABI.UnpackIntoInterface(&ethEvt, "NftBaseUriChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGBaseURIEvent{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewURI = ethEvt.NewValue
	return evt, nil
}

func (h *Handlers) storeNftBaseURIChanged(ctx context.Context, evt *cgmodel.CGBaseURIEvent) error {
	h.log.Info("NftBaseUriChanged", "evt_id", evt.EvtId, "new_uri", evt.NewURI)

	if err := h.repo.DeleteBaseURI(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertBaseURI(ctx, evt)
}

// --- Ownership / initialization across the platform contracts ---

// ownershipContractCode maps the emitting contract to the numeric code the
// cg_ownership_transferred table stores. The registry's source filter admits
// only these contracts, so the fallthrough 0 is unreachable in dispatch.
func (h *Handlers) ownershipContractCode(addr ethcommon.Address) int64 {
	switch addr {
	case h.c.Game:
		return 1
	case h.c.Signature:
		return 2
	case h.c.Token:
		return 3
	case h.c.CharityWallet:
		return 4
	case h.c.PrizesWallet:
		return 5
	case h.c.StakingCST:
		return 6
	case h.c.StakingRWalk:
		return 7
	case h.c.MarketingWallet:
		return 8
	case h.c.Dao:
		return 9
	default:
		return 0
	}
}

// ownershipSources lists the contracts whose OwnershipTransferred events are
// recorded (the legacy contract-code set; the implementation contract is
// deliberately absent).
func (h *Handlers) ownershipSources() []ethcommon.Address {
	return []ethcommon.Address{
		h.c.Game, h.c.Signature, h.c.Token, h.c.CharityWallet, h.c.PrizesWallet,
		h.c.StakingCST, h.c.StakingRWalk, h.c.MarketingWallet, h.c.Dao,
	}
}

func (h *Handlers) decodeOwnershipTransferred(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGOwnershipTransferred, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameOwnershipTransferred
	if err := h.signatureABI.UnpackIntoInterface(&ethEvt, "OwnershipTransferred", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGOwnershipTransferred{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.PrevOwner = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.NewOwner = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.ContractCode = h.ownershipContractCode(lg.Address)
	return evt, nil
}

func (h *Handlers) storeOwnershipTransferred(ctx context.Context, evt *cgmodel.CGOwnershipTransferred) error {
	h.log.Info("OwnershipTransferred",
		"evt_id", evt.EvtId, "contract_code", evt.ContractCode,
		"prev_owner", evt.PrevOwner, "new_owner", evt.NewOwner)

	if err := h.repo.DeleteOwnershipTransfer(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertOwnershipTransfer(ctx, evt)
}

// initializedSources lists the platform contracts that may emit OpenZeppelin
// Initializable:Initialized (the ownership set plus the implementation).
func (h *Handlers) initializedSources() []ethcommon.Address {
	return append(h.ownershipSources(), h.c.Implementation)
}

func (h *Handlers) decodeInitialized(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGInitialized, error) {
	var ethEvt cgc.CosmicSignatureGameInitialized
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "Initialized", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGInitialized{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.Version = int64(ethEvt.Version)
	return evt, nil
}

func (h *Handlers) storeInitialized(ctx context.Context, evt *cgmodel.CGInitialized) error {
	h.log.Info("Initialized", "evt_id", evt.EvtId, "version", evt.Version)

	if err := h.repo.DeleteInitialized(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertInitialized(ctx, evt)
}

// --- Remaining single-value parameters ---

func (h *Handlers) decodeCstMinLimitChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGCstMinLimit, error) {
	var ethEvt cgc.BiddingCstDutchAuctionBeginningBidPriceMinLimitChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "CstDutchAuctionBeginningBidPriceMinLimitChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGCstMinLimit{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.CstMinLimit = ethEvt.NewValue.String()
	return evt, nil
}

func (h *Handlers) storeCstMinLimitChanged(ctx context.Context, evt *cgmodel.CGCstMinLimit) error {
	h.log.Info("CstDutchAuctionBeginningBidPriceMinLimitChanged", "evt_id", evt.EvtId, "min_limit", evt.CstMinLimit)

	if err := h.repo.DeleteCstMinLimit(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertCstMinLimit(ctx, evt)
}

func (h *Handlers) decodeDelayDurationChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGNextRoundDelayDuration, error) {
	var ethEvt cgc.CosmicSignatureGameDelayDurationBeforeRoundActivationChanged
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "DelayDurationBeforeRoundActivationChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGNextRoundDelayDuration{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.NewValue = ethEvt.NewValue.Int64()
	return evt, nil
}

func (h *Handlers) storeDelayDurationChanged(ctx context.Context, evt *cgmodel.CGNextRoundDelayDuration) error {
	h.log.Info("DelayDurationBeforeRoundActivationChanged", "evt_id", evt.EvtId, "new_value", evt.NewValue)

	if err := h.repo.DeleteNextRoundDelayDurationChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertNextRoundDelayDurationChange(ctx, evt)
}
