// V4 (in development) decode/store pairs: the game proxy's art-auction
// wiring and ChaosZero cycle-funding events, plus the full event set of the
// non-upgradeable CosmicSignatureArtAuction contract. The ABIs are the
// hand-written fragments in contracts/cosmicgame/v4abi.go; the topic tests
// pin every constant against them.

package cosmicgame

import (
	"context"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Game events.

func (h *Handlers) decodeArtAuctionConfigured(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGArtAuctionConfiguredEvent, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGArtAuctionConfiguredEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.ArtAuctionAddr = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	return evt, nil
}

func (h *Handlers) storeArtAuctionConfigured(ctx context.Context, evt *cgmodel.CGArtAuctionConfiguredEvent) error {
	h.log.Info("ArtAuctionConfigured",
		"evt_id", evt.EvtId, "art_auction", evt.ArtAuctionAddr)

	if err := h.repo.DeleteArtAuctionConfigured(ctx, evt.EvtId); err != nil {
		return err
	}
	if err := h.repo.InsertArtAuctionConfigured(ctx, evt); err != nil {
		return err
	}
	h.noteArtAuction(ctx, ethcommon.HexToAddress(evt.ArtAuctionAddr), evt.BlockNum)
	return nil
}

func (h *Handlers) decodeArtAuctionOpeningBidChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGArtAuctionOpeningBidChangedEvent, error) {
	if err := requireTopics(lg, 1); err != nil {
		return nil, err
	}
	var raw struct {
		OpeningBid *big.Int
	}
	if err := h.gameV4ABI.UnpackIntoInterface(&raw, "ArtAuctionOpeningBidChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGArtAuctionOpeningBidChangedEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.OpeningBid = raw.OpeningBid.String()
	return evt, nil
}

func (h *Handlers) storeArtAuctionOpeningBidChanged(ctx context.Context, evt *cgmodel.CGArtAuctionOpeningBidChangedEvent) error {
	h.log.Info("ArtAuctionOpeningBidChanged",
		"evt_id", evt.EvtId, "opening_bid", evt.OpeningBid)

	if err := h.repo.DeleteArtAuctionOpeningBidChanged(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertArtAuctionOpeningBidChanged(ctx, evt)
}

func (h *Handlers) decodeArtAuctionNftsMinted(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGArtAuctionNftsMintedEvent, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var raw struct {
		NumNfts *big.Int
	}
	if err := h.gameV4ABI.UnpackIntoInterface(&raw, "ArtAuctionNftsMinted", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGArtAuctionNftsMintedEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.RoundNum = lg.Topics[1].Big().Int64()
	evt.ArtAuctionAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.FirstNftId = lg.Topics[3].Big().Int64()
	evt.NumNfts = raw.NumNfts.Int64()
	return evt, nil
}

func (h *Handlers) storeArtAuctionNftsMinted(ctx context.Context, evt *cgmodel.CGArtAuctionNftsMintedEvent) error {
	h.log.Info("ArtAuctionNftsMinted",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "first_nft_id", evt.FirstNftId, "num_nfts", evt.NumNfts)

	if err := h.repo.DeleteArtAuctionNftsMinted(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertArtAuctionNftsMinted(ctx, evt)
}

func (h *Handlers) decodeChaosZeroCycleFunded(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGChaosZeroCycleFundedEvent, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var raw struct {
		Amount           *big.Int
		CumulativeAmount *big.Int
	}
	if err := h.gameV4ABI.UnpackIntoInterface(&raw, "ChaosZeroCycleFunded", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGChaosZeroCycleFundedEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.MarketAddr = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Cycle = lg.Topics[2].Big().Int64()
	evt.Amount = raw.Amount.String()
	evt.CumulativeAmount = raw.CumulativeAmount.String()
	return evt, nil
}

func (h *Handlers) storeChaosZeroCycleFunded(ctx context.Context, evt *cgmodel.CGChaosZeroCycleFundedEvent) error {
	h.log.Info("ChaosZeroCycleFunded",
		"evt_id", evt.EvtId, "market", evt.MarketAddr, "cycle", evt.Cycle, "amount", evt.Amount)

	if err := h.repo.DeleteChaosZeroCycleFunded(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertChaosZeroCycleFunded(ctx, evt)
}

// Art-auction contract events.

func (h *Handlers) decodeAuctionBatchRegistered(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGArtAuctionBatchRegisteredEvent, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var raw struct {
		Count      *big.Int
		EligibleAt *big.Int
		OpeningBid *big.Int
	}
	if err := h.artAuctionABI.UnpackIntoInterface(&raw, "AuctionBatchRegistered", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGArtAuctionBatchRegisteredEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.RoundNum = lg.Topics[1].Big().Int64()
	evt.FirstNftId = lg.Topics[2].Big().Int64()
	evt.NumNfts = raw.Count.Int64()
	evt.EligibleAt = raw.EligibleAt.Int64()
	evt.OpeningBid = raw.OpeningBid.String()
	return evt, nil
}

func (h *Handlers) storeAuctionBatchRegistered(ctx context.Context, evt *cgmodel.CGArtAuctionBatchRegisteredEvent) error {
	h.log.Info("AuctionBatchRegistered",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "first_nft_id", evt.FirstNftId,
		"num_nfts", evt.NumNfts, "opening_bid", evt.OpeningBid)

	if err := h.repo.DeleteArtAuctionBatchRegistered(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertArtAuctionBatchRegistered(ctx, evt)
}

func (h *Handlers) decodeAuctionBidPlaced(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGArtAuctionBidEvent, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var raw struct {
		Amount  *big.Int
		EthPaid *big.Int
		EndAt   *big.Int
	}
	if err := h.artAuctionABI.UnpackIntoInterface(&raw, "AuctionBidPlaced", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGArtAuctionBidEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.NftId = lg.Topics[1].Big().Int64()
	evt.RoundNum = lg.Topics[2].Big().Int64()
	evt.BidderAddr = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	evt.Amount = raw.Amount.String()
	evt.EthPaid = raw.EthPaid.String()
	evt.EndAt = raw.EndAt.Int64()
	return evt, nil
}

func (h *Handlers) storeAuctionBidPlaced(ctx context.Context, evt *cgmodel.CGArtAuctionBidEvent) error {
	h.log.Info("AuctionBidPlaced",
		"evt_id", evt.EvtId, "nft_id", evt.NftId, "bidder", evt.BidderAddr, "amount", evt.Amount)

	if err := h.repo.DeleteArtAuctionBid(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertArtAuctionBid(ctx, evt)
}

func (h *Handlers) decodeAuctionRefundCredited(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGArtAuctionRefundCreditedEvent, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var raw struct {
		Amount *big.Int
	}
	if err := h.artAuctionABI.UnpackIntoInterface(&raw, "AuctionRefundCredited", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGArtAuctionRefundCreditedEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.BidderAddr = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.NftId = lg.Topics[2].Big().Int64()
	evt.Amount = raw.Amount.String()
	return evt, nil
}

func (h *Handlers) storeAuctionRefundCredited(ctx context.Context, evt *cgmodel.CGArtAuctionRefundCreditedEvent) error {
	h.log.Info("AuctionRefundCredited",
		"evt_id", evt.EvtId, "bidder", evt.BidderAddr, "nft_id", evt.NftId, "amount", evt.Amount)

	if err := h.repo.DeleteArtAuctionRefundCredited(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertArtAuctionRefundCredited(ctx, evt)
}

func (h *Handlers) decodeAuctionRefundWithdrawn(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGArtAuctionRefundWithdrawnEvent, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var raw struct {
		Amount *big.Int
	}
	if err := h.artAuctionABI.UnpackIntoInterface(&raw, "AuctionRefundWithdrawn", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGArtAuctionRefundWithdrawnEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.BidderAddr = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.RecipientAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Amount = raw.Amount.String()
	return evt, nil
}

func (h *Handlers) storeAuctionRefundWithdrawn(ctx context.Context, evt *cgmodel.CGArtAuctionRefundWithdrawnEvent) error {
	h.log.Info("AuctionRefundWithdrawn",
		"evt_id", evt.EvtId, "bidder", evt.BidderAddr, "recipient", evt.RecipientAddr, "amount", evt.Amount)

	if err := h.repo.DeleteArtAuctionRefundWithdrawn(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertArtAuctionRefundWithdrawn(ctx, evt)
}

func (h *Handlers) decodeAuctionSettled(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGArtAuctionSettledEvent, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var raw struct {
		Amount *big.Int
	}
	if err := h.artAuctionABI.UnpackIntoInterface(&raw, "AuctionSettled", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGArtAuctionSettledEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.NftId = lg.Topics[1].Big().Int64()
	evt.RoundNum = lg.Topics[2].Big().Int64()
	evt.WinnerAddr = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	evt.Amount = raw.Amount.String()
	return evt, nil
}

func (h *Handlers) storeAuctionSettled(ctx context.Context, evt *cgmodel.CGArtAuctionSettledEvent) error {
	h.log.Info("AuctionSettled",
		"evt_id", evt.EvtId, "nft_id", evt.NftId, "winner", evt.WinnerAddr, "amount", evt.Amount)

	if err := h.repo.DeleteArtAuctionSettled(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertArtAuctionSettled(ctx, evt)
}

func (h *Handlers) decodeAuctionNftDelivered(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGArtAuctionNftDeliveredEvent, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGArtAuctionNftDeliveredEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.NftId = lg.Topics[1].Big().Int64()
	evt.WinnerAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.RecipientAddr = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	return evt, nil
}

func (h *Handlers) storeAuctionNftDelivered(ctx context.Context, evt *cgmodel.CGArtAuctionNftDeliveredEvent) error {
	h.log.Info("AuctionNftDelivered",
		"evt_id", evt.EvtId, "nft_id", evt.NftId, "winner", evt.WinnerAddr, "recipient", evt.RecipientAddr)

	if err := h.repo.DeleteArtAuctionNftDelivered(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertArtAuctionNftDelivered(ctx, evt)
}

func (h *Handlers) decodeAuctionProceedsForwarded(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGArtAuctionProceedsForwardedEvent, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var raw struct {
		Amount *big.Int
	}
	if err := h.artAuctionABI.UnpackIntoInterface(&raw, "AuctionProceedsForwarded", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGArtAuctionProceedsForwardedEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.NftId = lg.Topics[1].Big().Int64()
	evt.RoundNum = lg.Topics[2].Big().Int64()
	evt.Amount = raw.Amount.String()
	return evt, nil
}

func (h *Handlers) storeAuctionProceedsForwarded(ctx context.Context, evt *cgmodel.CGArtAuctionProceedsForwardedEvent) error {
	h.log.Info("AuctionProceedsForwarded",
		"evt_id", evt.EvtId, "nft_id", evt.NftId, "round", evt.RoundNum, "amount", evt.Amount)

	if err := h.repo.DeleteArtAuctionProceedsForwarded(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertArtAuctionProceedsForwarded(ctx, evt)
}

func (h *Handlers) decodeAuctionSurplusForwarded(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGArtAuctionSurplusForwardedEvent, error) {
	if err := requireTopics(lg, 1); err != nil {
		return nil, err
	}
	var raw struct {
		Amount *big.Int
	}
	if err := h.artAuctionABI.UnpackIntoInterface(&raw, "AuctionSurplusForwarded", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGArtAuctionSurplusForwardedEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.TimeStamp = elog.TimeStamp
	evt.ContractAddr = lg.Address.String()
	evt.Amount = raw.Amount.String()
	return evt, nil
}

func (h *Handlers) storeAuctionSurplusForwarded(ctx context.Context, evt *cgmodel.CGArtAuctionSurplusForwardedEvent) error {
	h.log.Info("AuctionSurplusForwarded",
		"evt_id", evt.EvtId, "amount", evt.Amount)

	if err := h.repo.DeleteArtAuctionSurplusForwarded(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertArtAuctionSurplusForwarded(ctx, evt)
}
