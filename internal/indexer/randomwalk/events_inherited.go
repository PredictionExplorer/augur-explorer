// Decode/store pairs for the OpenZeppelin events RandomWalkNFT inherits
// (ERC721Enumerable approvals, Ownable ownership transfers).

package randomwalk

import (
	"context"
	"errors"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// errTokenIDOverflow rejects a tokenId topic that does not fit the BIGINT
// column; decode must fail the batch, never wrap.
var errTokenIDOverflow = errors.New("approval tokenId exceeds int64")

// decodeApproval handles IERC721.Approval(owner indexed, approved indexed,
// tokenId indexed); the log carries no data words.
func (h *Handlers) decodeApproval(lg *types.Log, elog *store.EthereumEventLog) (*rwmodel.Approval, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	tokenID := lg.Topics[3].Big()
	if !tokenID.IsInt64() {
		return nil, errTokenIDOverflow
	}
	evt := &rwmodel.Approval{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Owner = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Approved = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.TokenId = tokenID.Int64()
	return evt, nil
}

func (h *Handlers) storeApproval(ctx context.Context, evt *rwmodel.Approval) error {
	h.log.Info("Approval",
		"evt_id", evt.EvtId, "owner", evt.Owner, "approved", evt.Approved, "token_id", evt.TokenId)

	return h.repo.InsertApproval(ctx, evt)
}

func (h *Handlers) decodeApprovalForAll(lg *types.Log, elog *store.EthereumEventLog) (*rwmodel.ApprovalForAll, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt rwmodel.EApprovalForAll
	if err := h.rwalkABI.UnpackIntoInterface(&ethEvt, "ApprovalForAll", lg.Data); err != nil {
		return nil, err
	}
	evt := &rwmodel.ApprovalForAll{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Owner = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Operator = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Approved = ethEvt.Approved
	return evt, nil
}

func (h *Handlers) storeApprovalForAll(ctx context.Context, evt *rwmodel.ApprovalForAll) error {
	h.log.Info("ApprovalForAll",
		"evt_id", evt.EvtId, "owner", evt.Owner, "operator", evt.Operator, "approved", evt.Approved)

	return h.repo.InsertApprovalForAll(ctx, evt)
}

// decodeOwnershipTransferred handles Ownable.OwnershipTransferred(
// previousOwner indexed, newOwner indexed).
func (h *Handlers) decodeOwnershipTransferred(lg *types.Log, elog *store.EthereumEventLog) (*rwmodel.OwnershipTransferred, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	evt := &rwmodel.OwnershipTransferred{}
	evt.EvtId = elog.EvtID
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxID
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.PrevOwner = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.NewOwner = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	return evt, nil
}

func (h *Handlers) storeOwnershipTransferred(ctx context.Context, evt *rwmodel.OwnershipTransferred) error {
	h.log.Info("OwnershipTransferred",
		"evt_id", evt.EvtId, "prev_owner", evt.PrevOwner, "new_owner", evt.NewOwner)

	return h.repo.InsertOwnershipTransferred(ctx, evt)
}
