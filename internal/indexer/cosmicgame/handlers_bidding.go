// Bid events: BidPlaced (v1 and v2) and FirstBidPlacedInRound (round start).

package cosmicgame

import (
	"context"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func (h *Handlers) decodeBidPlacedV1(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGBidEvent, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameBidPlaced
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "BidPlaced", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGBidEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.LastBidderAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.RoundNum = lg.Topics[1].Big().Int64()
	evt.EthPrice = ethEvt.PaidEthPrice.String()
	evt.BidType = 0 // ETH
	evt.RandomWalkTokenId = lg.Topics[3].Big().Int64()
	evt.CstPrice = ethEvt.PaidCstPrice.String()
	if evt.RandomWalkTokenId > -1 {
		evt.BidType = 1 // RandomWalk
	} else if evt.CstPrice != "-1" {
		evt.BidType = 2 // Cosmic Signature Token (ERC20) bid
	}
	evt.PrizeTime = ethEvt.MainPrizeTime.Int64()
	evt.Message = ethEvt.Message
	evt.BidCstRewardAmount = "-1"
	evt.CstDutchAuctionDuration = "-1"
	return evt, nil
}

func (h *Handlers) decodeBidPlacedV2(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGBidEvent, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameV2BidPlaced
	if err := h.gameV2ABI.UnpackIntoInterface(&ethEvt, "BidPlaced", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGBidEvent{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.LastBidderAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.RoundNum = lg.Topics[1].Big().Int64()
	evt.EthPrice = ethEvt.PaidEthPrice.String()
	evt.BidType = 0 // ETH
	evt.RandomWalkTokenId = lg.Topics[3].Big().Int64()
	evt.CstPrice = ethEvt.PaidCstPrice.String()
	if evt.RandomWalkTokenId > -1 {
		evt.BidType = 1 // RandomWalk
	} else if evt.CstPrice != "-1" {
		evt.BidType = 2 // Cosmic Signature Token (ERC20) bid
	}
	evt.PrizeTime = ethEvt.MainPrizeTime.Int64()
	evt.Message = ethEvt.Message
	evt.BidCstRewardAmount = ethEvt.BidCstRewardAmount.String()
	evt.CstDutchAuctionDuration = ethEvt.CstDutchAuctionDuration.String()
	return evt, nil
}

// storeBid persists a decoded bid: the CST mint split is resolved from the
// transaction's ERC20 Transfer logs, then the row and normalized reward
// shares are written delete-then-insert.
func (h *Handlers) storeBid(ctx context.Context, evt *cgmodel.CGBidEvent) error {
	thisReward, previousReward, previousAddr, err := h.cstBidRewards(
		ctx, evt.EvtId, evt.TxId, evt.LastBidderAddr,
	)
	if err != nil {
		return fmt.Errorf("bid (evt id %v): %w", evt.EvtId, err)
	}
	evt.ThisBidderReward = thisReward
	evt.PrevBidderReward = previousReward
	evt.PrevBidderAddr = previousAddr
	evt.ERC20Value = evt.ThisBidderReward

	h.log.Info("BidPlaced",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "bidder", evt.LastBidderAddr,
		"bid_type", evt.BidType, "eth_price", evt.EthPrice, "cst_price", evt.CstPrice,
		"rwalk_token_id", evt.RandomWalkTokenId,
		"this_bidder_reward", evt.ThisBidderReward,
		"previous_bidder_reward", evt.PrevBidderReward,
		"previous_bidder", evt.PrevBidderAddr,
		"prize_time", evt.PrizeTime, "message", evt.Message)

	if err := h.repo.DeleteBid(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertBid(ctx, evt)
}

func (h *Handlers) decodeFirstBidPlacedInRound(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGRoundStarted, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameFirstBidPlacedInRound
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "FirstBidPlacedInRound", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGRoundStarted{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.Contract = lg.Address.String()
	evt.RoundNum = lg.Topics[1].Big().Int64()
	evt.StartTimestamp = ethEvt.BlockTimeStamp.Int64()
	return evt, nil
}

func (h *Handlers) storeFirstBidPlacedInRound(ctx context.Context, evt *cgmodel.CGRoundStarted) error {
	h.log.Info("FirstBidPlacedInRound",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "start_ts", evt.StartTimestamp)

	if err := h.repo.DeleteRoundStarted(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertRoundStarted(ctx, evt)
}
