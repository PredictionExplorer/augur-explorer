// Event-row inserts for the OpenZeppelin events the platform contracts
// inherit (ERC-721/ERC-20 approvals, ERC-20 Votes delegation, EIP-712 domain
// changes) and for the Governor events of CosmicSignatureDao. Same contract
// as inserts.go: resolve address ids, one INSERT, return errors.

package cosmicgame

import (
	"context"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// InsertNftApproval records an ERC-721 Approval of the Cosmic Signature NFT.
func (r *Repo) InsertNftApproval(ctx context.Context, evt *cgmodel.CGNftApproval) error {
	const op = "insert into cg_nft_approval"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	ownerAid, err := r.addrID(ctx, evt.Owner, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	approvedAid, err := r.addrID(ctx, evt.Approved, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_nft_approval(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"owner_aid,approved_aid,token_id" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		ownerAid, approvedAid, evt.TokenId,
	)
	return store.WrapError(op, err)
}

// InsertNftApprovalForAll records an ERC-721 ApprovalForAll of the Cosmic
// Signature NFT.
func (r *Repo) InsertNftApprovalForAll(ctx context.Context, evt *cgmodel.CGNftApprovalForAll) error {
	const op = "insert into cg_nft_approval_for_all"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	ownerAid, err := r.addrID(ctx, evt.Owner, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	operatorAid, err := r.addrID(ctx, evt.Operator, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_nft_approval_for_all(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"owner_aid,operator_aid,approved" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		ownerAid, operatorAid, evt.Approved,
	)
	return store.WrapError(op, err)
}

// InsertTokenApproval records an ERC-20 Approval of the CosmicToken.
func (r *Repo) InsertTokenApproval(ctx context.Context, evt *cgmodel.CGTokenApproval) error {
	const op = "insert into cg_token_approval"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	ownerAid, err := r.addrID(ctx, evt.Owner, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	spenderAid, err := r.addrID(ctx, evt.Spender, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_token_approval(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"owner_aid,spender_aid,value" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		ownerAid, spenderAid, evt.Value,
	)
	return store.WrapError(op, err)
}

// InsertDelegateChanged records an ERC-20 Votes DelegateChanged of the
// CosmicToken.
func (r *Repo) InsertDelegateChanged(ctx context.Context, evt *cgmodel.CGDelegateChanged) error {
	const op = "insert into cg_token_delegate_changed"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	delegatorAid, err := r.addrID(ctx, evt.Delegator, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	fromAid, err := r.addrID(ctx, evt.FromDelegate, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	toAid, err := r.addrID(ctx, evt.ToDelegate, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_token_delegate_changed(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"delegator_aid,from_delegate_aid,to_delegate_aid" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		delegatorAid, fromAid, toAid,
	)
	return store.WrapError(op, err)
}

// InsertDelegateVotesChanged records an ERC-20 Votes DelegateVotesChanged of
// the CosmicToken.
func (r *Repo) InsertDelegateVotesChanged(ctx context.Context, evt *cgmodel.CGDelegateVotesChanged) error {
	const op = "insert into cg_token_delegate_votes_changed"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	delegateAid, err := r.addrID(ctx, evt.Delegate, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_token_delegate_votes_changed(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"delegate_aid,previous_votes,new_votes" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		delegateAid, evt.PreviousVotes, evt.NewVotes,
	)
	return store.WrapError(op, err)
}

// InsertEIP712DomainChanged records an EIP712DomainChanged event.
func (r *Repo) InsertEIP712DomainChanged(ctx context.Context, evt *cgmodel.CGEIP712DomainChanged) error {
	const op = "insert into cg_eip712_domain_changed"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_eip712_domain_changed(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5)"
	_, err = r.q(ctx).Exec(ctx, query, evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid)
	return store.WrapError(op, err)
}

// InsertDaoProposalCreated records a Governor ProposalCreated event.
func (r *Repo) InsertDaoProposalCreated(ctx context.Context, evt *cgmodel.CGDaoProposalCreated) error {
	const op = "insert into cg_dao_proposal_created"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	proposerAid, err := r.addrID(ctx, evt.Proposer, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	// Non-nil slices keep pgx from encoding NULL arrays into NOT NULL columns.
	targets := evt.Targets
	if targets == nil {
		targets = []string{}
	}
	values := evt.Values
	if values == nil {
		values = []string{}
	}
	signatures := evt.Signatures
	if signatures == nil {
		signatures = []string{}
	}
	calldatas := evt.Calldatas
	if calldatas == nil {
		calldatas = [][]byte{}
	}
	query := "INSERT INTO cg_dao_proposal_created(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"proposal_id,proposer_aid,targets,call_values,signatures,calldatas,vote_start,vote_end,description" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9::DECIMAL[],$10,$11,$12,$13,$14)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.ProposalId, proposerAid, targets, values, signatures, calldatas,
		evt.VoteStart, evt.VoteEnd, evt.Description,
	)
	return store.WrapError(op, err)
}

// InsertDaoProposalStateChange records ProposalQueued / ProposalExecuted /
// ProposalCanceled.
func (r *Repo) InsertDaoProposalStateChange(ctx context.Context, evt *cgmodel.CGDaoProposalStateChange) error {
	const op = "insert into cg_dao_proposal_state"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	var eta *int64
	if evt.State == cgmodel.DaoProposalQueued {
		eta = &evt.EtaSeconds
	}
	query := "INSERT INTO cg_dao_proposal_state(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"proposal_id,state,eta_seconds" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.ProposalId, evt.State, eta,
	)
	return store.WrapError(op, err)
}

// InsertDaoVoteCast records a Governor VoteCast / VoteCastWithParams event.
func (r *Repo) InsertDaoVoteCast(ctx context.Context, evt *cgmodel.CGDaoVoteCast) error {
	const op = "insert into cg_dao_vote_cast"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	voterAid, err := r.addrID(ctx, evt.Voter, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_dao_vote_cast(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"voter_aid,proposal_id,support,weight,reason,params" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10,$11)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		voterAid, evt.ProposalId, evt.Support, evt.Weight, evt.Reason, evt.Params,
	)
	return store.WrapError(op, err)
}

// InsertDaoSettingChanged records one of the Governor (old, new) parameter
// events.
func (r *Repo) InsertDaoSettingChanged(ctx context.Context, evt *cgmodel.CGDaoSettingChanged) error {
	const op = "insert into cg_dao_setting_changed"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_dao_setting_changed(" +
		"evtlog_id,block_num,tx_id,time_stamp,contract_aid," +
		"setting,old_value,new_value" +
		") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.Setting, evt.OldValue, evt.NewValue,
	)
	return store.WrapError(op, err)
}
