// Events the platform contracts inherit from OpenZeppelin (ERC-721/ERC-20
// approvals, ERC-20 Votes delegation, EIP-712 domain changes) and the
// Governor events of CosmicSignatureDao. Every event a deployed contract can
// emit is dispatched; TestEveryABIEventHasHandler enforces it.

package cosmicgame

import (
	"context"
	"fmt"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// int64Of converts a uint256 that the schema stores in a BIGINT column;
// out-of-range values fail the batch instead of wrapping.
func int64Of(name string, v *big.Int) (int64, error) {
	if v == nil || !v.IsInt64() || v.Sign() < 0 {
		return 0, fmt.Errorf("%s %v does not fit int64", name, v)
	}
	return v.Int64(), nil
}

// --- Cosmic Signature NFT (ERC-721) ---

// decodeNftApproval handles IERC721.Approval(owner indexed, approved indexed,
// tokenId indexed). The topic0 is shared with the ERC-20 Approval; the
// registry's source filter routes the NFT contract here.
func (h *Handlers) decodeNftApproval(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGNftApproval, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	tokenID, err := int64Of("tokenId", lg.Topics[3].Big())
	if err != nil {
		return nil, err
	}
	evt := &cgmodel.CGNftApproval{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.Owner = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Approved = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.TokenId = tokenID
	return evt, nil
}

func (h *Handlers) storeNftApproval(ctx context.Context, evt *cgmodel.CGNftApproval) error {
	h.log.Info("CosmicSignature Approval",
		"evt_id", evt.EvtId, "owner", evt.Owner, "approved", evt.Approved, "token_id", evt.TokenId)

	if err := h.repo.DeleteNftApproval(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertNftApproval(ctx, evt)
}

func (h *Handlers) decodeNftApprovalForAll(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGNftApprovalForAll, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureNftApprovalForAll
	if err := h.signatureABI.UnpackIntoInterface(&ethEvt, "ApprovalForAll", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGNftApprovalForAll{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.Owner = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Operator = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Approved = ethEvt.Approved
	return evt, nil
}

func (h *Handlers) storeNftApprovalForAll(ctx context.Context, evt *cgmodel.CGNftApprovalForAll) error {
	h.log.Info("CosmicSignature ApprovalForAll",
		"evt_id", evt.EvtId, "owner", evt.Owner, "operator", evt.Operator, "approved", evt.Approved)

	if err := h.repo.DeleteNftApprovalForAll(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertNftApprovalForAll(ctx, evt)
}

// --- CosmicToken (ERC-20, ERC-20 Votes, ERC-20 Permit) ---

// decodeTokenApproval handles IERC20.Approval(owner indexed, spender indexed,
// value).
func (h *Handlers) decodeTokenApproval(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGTokenApproval, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureTokenApproval
	if err := h.tokenABI.UnpackIntoInterface(&ethEvt, "Approval", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGTokenApproval{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.Owner = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Spender = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Value = ethEvt.Value.String()
	return evt, nil
}

func (h *Handlers) storeTokenApproval(ctx context.Context, evt *cgmodel.CGTokenApproval) error {
	h.log.Info("CosmicToken Approval",
		"evt_id", evt.EvtId, "owner", evt.Owner, "spender", evt.Spender, "value", evt.Value)

	if err := h.repo.DeleteTokenApproval(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertTokenApproval(ctx, evt)
}

func (h *Handlers) decodeDelegateChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDelegateChanged, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGDelegateChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.Delegator = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.FromDelegate = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.ToDelegate = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	return evt, nil
}

func (h *Handlers) storeDelegateChanged(ctx context.Context, evt *cgmodel.CGDelegateChanged) error {
	h.log.Info("CosmicToken DelegateChanged",
		"evt_id", evt.EvtId, "delegator", evt.Delegator, "from", evt.FromDelegate, "to", evt.ToDelegate)

	if err := h.repo.DeleteDelegateChanged(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertDelegateChanged(ctx, evt)
}

func (h *Handlers) decodeDelegateVotesChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDelegateVotesChanged, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureTokenDelegateVotesChanged
	if err := h.tokenABI.UnpackIntoInterface(&ethEvt, "DelegateVotesChanged", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGDelegateVotesChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.Delegate = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.PreviousVotes = ethEvt.PreviousVotes.String()
	evt.NewVotes = ethEvt.NewVotes.String()
	return evt, nil
}

func (h *Handlers) storeDelegateVotesChanged(ctx context.Context, evt *cgmodel.CGDelegateVotesChanged) error {
	h.log.Info("CosmicToken DelegateVotesChanged",
		"evt_id", evt.EvtId, "delegate", evt.Delegate, "previous_votes", evt.PreviousVotes, "new_votes", evt.NewVotes)

	if err := h.repo.DeleteDelegateVotesChanged(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertDelegateVotesChanged(ctx, evt)
}

// decodeEIP712DomainChanged handles the parameterless IERC5267 event emitted
// by the token (permit domain) and the DAO (vote-by-signature domain).
func (h *Handlers) decodeEIP712DomainChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGEIP712DomainChanged, error) {
	evt := &cgmodel.CGEIP712DomainChanged{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	return evt, nil
}

func (h *Handlers) storeEIP712DomainChanged(ctx context.Context, evt *cgmodel.CGEIP712DomainChanged) error {
	h.log.Info("EIP712DomainChanged", "evt_id", evt.EvtId, "contract", evt.Contract)

	if err := h.repo.DeleteEIP712DomainChanged(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertEIP712DomainChanged(ctx, evt)
}

// --- CosmicSignatureDao (OpenZeppelin Governor) ---

func (h *Handlers) decodeDaoProposalCreated(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDaoProposalCreated, error) {
	var ethEvt cgc.CosmicSignatureDaoProposalCreated
	if err := h.daoABI.UnpackIntoInterface(&ethEvt, "ProposalCreated", lg.Data); err != nil {
		return nil, err
	}
	voteStart, err := int64Of("voteStart", ethEvt.VoteStart)
	if err != nil {
		return nil, err
	}
	voteEnd, err := int64Of("voteEnd", ethEvt.VoteEnd)
	if err != nil {
		return nil, err
	}
	evt := &cgmodel.CGDaoProposalCreated{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.ProposalId = ethEvt.ProposalId.String()
	evt.Proposer = ethEvt.Proposer.String()
	evt.Targets = make([]string, len(ethEvt.Targets))
	for i, a := range ethEvt.Targets {
		evt.Targets[i] = a.String()
	}
	evt.Values = make([]string, len(ethEvt.Values))
	for i, v := range ethEvt.Values {
		evt.Values[i] = v.String()
	}
	evt.Signatures = ethEvt.Signatures
	evt.Calldatas = ethEvt.Calldatas
	evt.VoteStart = voteStart
	evt.VoteEnd = voteEnd
	evt.Description = ethEvt.Description
	return evt, nil
}

func (h *Handlers) storeDaoProposalCreated(ctx context.Context, evt *cgmodel.CGDaoProposalCreated) error {
	h.log.Info("DAO ProposalCreated",
		"evt_id", evt.EvtId, "proposal_id", evt.ProposalId, "proposer", evt.Proposer,
		"num_targets", len(evt.Targets), "vote_start", evt.VoteStart, "vote_end", evt.VoteEnd)

	if err := h.repo.DeleteDaoProposalCreated(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertDaoProposalCreated(ctx, evt)
}

func (h *Handlers) daoProposalState(lg *types.Log, elog *store.EthereumEventLog, state int64, proposalID, eta *big.Int) (*cgmodel.CGDaoProposalStateChange, error) {
	evt := &cgmodel.CGDaoProposalStateChange{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.ProposalId = proposalID.String()
	evt.State = state
	if eta != nil {
		v, err := int64Of("etaSeconds", eta)
		if err != nil {
			return nil, err
		}
		evt.EtaSeconds = v
	}
	return evt, nil
}

func (h *Handlers) decodeDaoProposalCanceled(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDaoProposalStateChange, error) {
	var ethEvt cgc.CosmicSignatureDaoProposalCanceled
	if err := h.daoABI.UnpackIntoInterface(&ethEvt, "ProposalCanceled", lg.Data); err != nil {
		return nil, err
	}
	return h.daoProposalState(lg, elog, cgmodel.DaoProposalCanceled, ethEvt.ProposalId, nil)
}

func (h *Handlers) decodeDaoProposalExecuted(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDaoProposalStateChange, error) {
	var ethEvt cgc.CosmicSignatureDaoProposalExecuted
	if err := h.daoABI.UnpackIntoInterface(&ethEvt, "ProposalExecuted", lg.Data); err != nil {
		return nil, err
	}
	return h.daoProposalState(lg, elog, cgmodel.DaoProposalExecuted, ethEvt.ProposalId, nil)
}

func (h *Handlers) decodeDaoProposalQueued(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDaoProposalStateChange, error) {
	var ethEvt cgc.CosmicSignatureDaoProposalQueued
	if err := h.daoABI.UnpackIntoInterface(&ethEvt, "ProposalQueued", lg.Data); err != nil {
		return nil, err
	}
	return h.daoProposalState(lg, elog, cgmodel.DaoProposalQueued, ethEvt.ProposalId, ethEvt.EtaSeconds)
}

func (h *Handlers) storeDaoProposalStateChange(ctx context.Context, evt *cgmodel.CGDaoProposalStateChange) error {
	h.log.Info("DAO proposal state change",
		"evt_id", evt.EvtId, "proposal_id", evt.ProposalId, "state", evt.State, "eta_seconds", evt.EtaSeconds)

	if err := h.repo.DeleteDaoProposalStateChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertDaoProposalStateChange(ctx, evt)
}

func (h *Handlers) decodeDaoVoteCast(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDaoVoteCast, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureDaoVoteCast
	if err := h.daoABI.UnpackIntoInterface(&ethEvt, "VoteCast", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGDaoVoteCast{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.Voter = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.ProposalId = ethEvt.ProposalId.String()
	evt.Support = int64(ethEvt.Support)
	evt.Weight = ethEvt.Weight.String()
	evt.Reason = ethEvt.Reason
	return evt, nil
}

func (h *Handlers) decodeDaoVoteCastWithParams(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDaoVoteCast, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureDaoVoteCastWithParams
	if err := h.daoABI.UnpackIntoInterface(&ethEvt, "VoteCastWithParams", lg.Data); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGDaoVoteCast{}
	evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
	evt.Voter = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.ProposalId = ethEvt.ProposalId.String()
	evt.Support = int64(ethEvt.Support)
	evt.Weight = ethEvt.Weight.String()
	evt.Reason = ethEvt.Reason
	evt.Params = ethEvt.Params
	if evt.Params == nil {
		evt.Params = []byte{}
	}
	return evt, nil
}

func (h *Handlers) storeDaoVoteCast(ctx context.Context, evt *cgmodel.CGDaoVoteCast) error {
	h.log.Info("DAO VoteCast",
		"evt_id", evt.EvtId, "voter", evt.Voter, "proposal_id", evt.ProposalId,
		"support", evt.Support, "weight", evt.Weight, "with_params", evt.Params != nil)

	if err := h.repo.DeleteDaoVoteCast(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertDaoVoteCast(ctx, evt)
}

// decodeDaoSettingChanged builds the decoder for one of the Governor's
// (oldValue, newValue) events; all four share the same data layout.
func (h *Handlers) decodeDaoSettingChanged(eventName string, setting int64) func(*types.Log, *store.EthereumEventLog) (*cgmodel.CGDaoSettingChanged, error) {
	return func(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDaoSettingChanged, error) {
		values, err := h.daoABI.Unpack(eventName, lg.Data)
		if err != nil {
			return nil, err
		}
		if len(values) != 2 {
			return nil, fmt.Errorf("%s: %d values, want 2", eventName, len(values))
		}
		oldValue, ok1 := values[0].(*big.Int)
		newValue, ok2 := values[1].(*big.Int)
		if !ok1 || !ok2 {
			return nil, fmt.Errorf("%s: unexpected value types %T, %T", eventName, values[0], values[1])
		}
		evt := &cgmodel.CGDaoSettingChanged{}
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, evt.Contract = adminEventBase(lg, elog)
		evt.Setting = setting
		evt.OldValue = oldValue.String()
		evt.NewValue = newValue.String()
		return evt, nil
	}
}

func (h *Handlers) storeDaoSettingChanged(ctx context.Context, evt *cgmodel.CGDaoSettingChanged) error {
	h.log.Info("DAO setting changed",
		"evt_id", evt.EvtId, "setting", evt.Setting, "old_value", evt.OldValue, "new_value", evt.NewValue)

	if err := h.repo.DeleteDaoSettingChanged(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertDaoSettingChanged(ctx, evt)
}
