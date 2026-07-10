// Decode/store pairs for the seven dispatched RandomWalk events.

package randomwalk

import (
	"context"
	"encoding/hex"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/primitives"
	rwp "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
)

func (h *Handlers) decodeNewOffer(lg *types.Log, elog *primitives.EthereumEventLog) (*rwp.NewOffer, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt rwp.ENewOffer
	if err := h.marketABI.UnpackIntoInterface(&ethEvt, "NewOffer", lg.Data); err != nil {
		return nil, err
	}

	evt := &rwp.NewOffer{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Seller = ethEvt.Seller.String()
	evt.Buyer = ethEvt.Buyer.String()
	evt.Price = ethEvt.Price.String()
	evt.RWalkAddr = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.OfferId = lg.Topics[2].Big().Int64()
	evt.TokenId = lg.Topics[3].Big().Int64()
	return evt, nil
}

func (h *Handlers) storeNewOffer(ctx context.Context, evt *rwp.NewOffer) error {
	h.log.Info("NewOffer",
		"evt_id", evt.EvtId, "nft", evt.RWalkAddr, "offer_id", evt.OfferId,
		"token_id", evt.TokenId, "seller", evt.Seller, "buyer", evt.Buyer, "price", evt.Price)

	return h.repo.InsertNewOffer(ctx, evt)
}

func (h *Handlers) decodeItemBought(lg *types.Log, elog *primitives.EthereumEventLog) (*rwp.ItemBought, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	evt := &rwp.ItemBought{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.OfferId = lg.Topics[1].Big().Int64()
	evt.SellerAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.BuyerAddr = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	return evt, nil
}

func (h *Handlers) storeItemBought(ctx context.Context, evt *rwp.ItemBought) error {
	h.log.Info("ItemBought",
		"evt_id", evt.EvtId, "offer_id", evt.OfferId, "seller", evt.SellerAddr, "buyer", evt.BuyerAddr)

	return h.repo.InsertItemBought(ctx, evt)
}

func (h *Handlers) decodeOfferCanceled(lg *types.Log, elog *primitives.EthereumEventLog) (*rwp.OfferCanceled, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	evt := &rwp.OfferCanceled{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.OfferId = lg.Topics[1].Big().Int64()
	return evt, nil
}

func (h *Handlers) storeOfferCanceled(ctx context.Context, evt *rwp.OfferCanceled) error {
	// An OfferCanceled for an offer this database never saw (pre-genesis or
	// foreign deployment) is skipped; a failing existence check aborts the
	// batch (treating it as "does not exist" silently lost events).
	exists, err := h.repo.OfferExists(ctx, evt.Contract, evt.OfferId)
	if err != nil {
		return fmt.Errorf("offer exists check for OfferCanceled: %w", err)
	}
	if !exists {
		h.log.Info("OfferCanceled skipped: unknown offer", "evt_id", evt.EvtId, "offer_id", evt.OfferId, "contract", evt.Contract)
		return nil
	}

	h.log.Info("OfferCanceled", "evt_id", evt.EvtId, "offer_id", evt.OfferId)

	return h.repo.InsertOfferCanceled(ctx, evt)
}

func (h *Handlers) decodeWithdrawal(lg *types.Log, elog *primitives.EthereumEventLog) (*rwp.Withdrawal, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt rwp.EWithdrawalEvent
	if err := h.rwalkABI.UnpackIntoInterface(&ethEvt, "WithdrawalEvent", lg.Data); err != nil {
		return nil, err
	}

	evt := &rwp.Withdrawal{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = lg.Topics[1].Big().Int64()
	evt.Destination = ethEvt.Destination.String()
	evt.Amount = ethEvt.Amount.String()
	return evt, nil
}

func (h *Handlers) storeWithdrawal(ctx context.Context, evt *rwp.Withdrawal) error {
	h.log.Info("WithdrawalEvent",
		"evt_id", evt.EvtId, "token_id", evt.TokenId, "destination", evt.Destination, "amount", evt.Amount)

	return h.repo.InsertWithdrawal(ctx, evt)
}

func (h *Handlers) decodeTokenName(lg *types.Log, elog *primitives.EthereumEventLog) (*rwp.TokenName, error) {
	var ethEvt rwp.ETokenNameEvent
	if err := h.rwalkABI.UnpackIntoInterface(&ethEvt, "TokenNameEvent", lg.Data); err != nil {
		return nil, err
	}

	evt := &rwp.TokenName{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = ethEvt.TokenId.Int64()
	evt.NewName = ethEvt.NewName
	return evt, nil
}

func (h *Handlers) storeTokenName(ctx context.Context, evt *rwp.TokenName) error {
	// A name change for a token this database never saw is skipped; a
	// failing existence check aborts the batch.
	exists, err := h.repo.TokenExists(ctx, evt.Contract, evt.TokenId)
	if err != nil {
		return fmt.Errorf("token exists check for TokenName: %w", err)
	}
	if !exists {
		h.log.Info("TokenNameEvent skipped: unknown token", "evt_id", evt.EvtId, "token_id", evt.TokenId, "contract", evt.Contract)
		return nil
	}

	h.log.Info("TokenNameEvent", "evt_id", evt.EvtId, "token_id", evt.TokenId, "new_name", evt.NewName)

	return h.repo.InsertTokenName(ctx, evt)
}

func (h *Handlers) decodeMintEvent(lg *types.Log, elog *primitives.EthereumEventLog) (*rwp.MintEvent, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt rwp.EMintEvent
	if err := h.rwalkABI.UnpackIntoInterface(&ethEvt, "MintEvent", lg.Data); err != nil {
		return nil, err
	}

	evt := &rwp.MintEvent{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = lg.Topics[1].Big().Int64()
	evt.Owner = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Seed = hex.EncodeToString(ethEvt.Seed[:])
	evt.SeedNum = ethcommon.BytesToHash(ethEvt.Seed[:]).Big().String()
	evt.Price = ethEvt.Price.String()
	return evt, nil
}

func (h *Handlers) storeMintEvent(ctx context.Context, evt *rwp.MintEvent) error {
	h.log.Info("MintEvent",
		"evt_id", evt.EvtId, "token_id", evt.TokenId, "owner", evt.Owner,
		"seed", evt.Seed, "price", evt.Price)

	return h.repo.InsertMint(ctx, evt)
}

// decodeTransfer handles the RandomWalk ERC721 Transfer. Only the canonical
// RandomWalk deployment is registered as a source; transfers of other
// instances of the contract are filtered before decode.
func (h *Handlers) decodeTransfer(lg *types.Log, elog *primitives.EthereumEventLog) (*rwp.Transfer, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	evt := &rwp.Transfer{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.From = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.To = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.TokenId = lg.Topics[3].Big().Int64()
	return evt, nil
}

func (h *Handlers) storeTransfer(ctx context.Context, evt *rwp.Transfer) error {
	h.log.Info("Transfer",
		"evt_id", evt.EvtId, "from", evt.From, "to", evt.To, "token_id", evt.TokenId)

	return h.repo.InsertTransfer(ctx, evt)
}
