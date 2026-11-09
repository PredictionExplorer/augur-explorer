// Insert methods for the cz_ event tables, one per dispatched ChaosZero
// event type. Like the RandomWalk store steps, inserts run without a
// preceding delete-by-evtlog-id: re-processing the same evt_log row violates
// the tables' UNIQUE(evtlog_id) constraints, and legitimate re-processing
// only happens through the reorg path, which cascade-deletes the rows first.
//
// uint256 values arrive as decimal strings (pgx encodes them into DECIMAL
// columns); uint256[8] arrays arrive as []string and are passed through a
// TEXT[]::NUMERIC[] cast.
package chaoszero

import (
	"context"

	czmodel "github.com/PredictionExplorer/augur-explorer/internal/model/chaoszero"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// InsertRoundInitialized records a market RoundInitialized event.
func (r *Repo) InsertRoundInitialized(ctx context.Context, evt *czmodel.RoundInitialized) error {
	const op = "insert into cz_round_initialized"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO cz_round_initialized(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,round_id
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6)`
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.RoundId,
	)
	return store.WrapError(op, err)
}

// InsertRoundFunded records a market RoundFunded event.
func (r *Repo) InsertRoundFunded(ctx context.Context, evt *czmodel.RoundFunded) error {
	const op = "insert into cz_round_funded"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	callerAid, err := r.addrID(ctx, evt.Caller, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO cz_round_funded(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,caller_aid,donated,transfers_used,carry_used,allocations
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10,$11::TEXT[]::NUMERIC[])`
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.RoundId, callerAid, evt.Donated, evt.TransfersUsed, evt.CarryUsed,
		evt.Allocations,
	)
	return store.WrapError(op, err)
}

// InsertTrade records a market Trade event.
func (r *Repo) InsertTrade(ctx context.Context, evt *czmodel.Trade) error {
	const op = "insert into cz_trade"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	userAid, err := r.addrID(ctx, evt.User, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO cz_trade(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,user_aid,question,yes,buy,cash,units
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10,$11,$12)`
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.RoundId, userAid, evt.Question, evt.Yes, evt.Buy, evt.Cash, evt.Units,
	)
	return store.WrapError(op, err)
}

// InsertPositionSplit records a market PositionSplit event.
func (r *Repo) InsertPositionSplit(ctx context.Context, evt *czmodel.PositionSplit) error {
	const op = "insert into cz_position_split"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	userAid, err := r.addrID(ctx, evt.User, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO cz_position_split(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,user_aid,question,amount
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9)`
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.RoundId, userAid, evt.Question, evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertPositionsMerged records a market PositionsMerged event.
func (r *Repo) InsertPositionsMerged(ctx context.Context, evt *czmodel.PositionsMerged) error {
	const op = "insert into cz_positions_merged"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	userAid, err := r.addrID(ctx, evt.User, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO cz_positions_merged(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,user_aid,amounts,collateral_out
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8::TEXT[]::NUMERIC[],$9)`
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.RoundId, userAid, evt.Amounts, evt.CollateralOut,
	)
	return store.WrapError(op, err)
}

// InsertRoundSettled records a market RoundSettled event.
func (r *Repo) InsertRoundSettled(ctx context.Context, evt *czmodel.RoundSettled) error {
	const op = "insert into cz_round_settled"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO cz_round_settled(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,values_mask,answers,escrowed,carried
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9::TEXT[]::NUMERIC[],$10)`
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.RoundId, evt.ValuesMask, evt.Answers, evt.Escrowed, evt.Carried,
	)
	return store.WrapError(op, err)
}

// InsertClaim records a market Claimed event.
func (r *Repo) InsertClaim(ctx context.Context, evt *czmodel.Claim) error {
	const op = "insert into cz_claim"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	userAid, err := r.addrID(ctx, evt.User, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO cz_claim(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,user_aid,amounts,total
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8::TEXT[]::NUMERIC[],$9)`
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.RoundId, userAid, evt.Amounts, evt.Total,
	)
	return store.WrapError(op, err)
}

// InsertAnswersObserved records a market AnswersObserved event.
func (r *Repo) InsertAnswersObserved(ctx context.Context, evt *czmodel.AnswersObserved) error {
	const op = "insert into cz_answers_observed"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO cz_answers_observed(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,known_mask,values_mask
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8)`
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.RoundId, evt.KnownMask, evt.ValuesMask,
	)
	return store.WrapError(op, err)
}

// InsertCycleEscrowed records a vault CycleEscrowed event.
func (r *Repo) InsertCycleEscrowed(ctx context.Context, evt *czmodel.CycleEscrowed) error {
	const op = "insert into cz_cycle_escrowed"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO cz_cycle_escrowed(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,amounts
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7::TEXT[]::NUMERIC[])`
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.RoundId, evt.Amounts,
	)
	return store.WrapError(op, err)
}

// InsertPaid records a vault Paid event.
func (r *Repo) InsertPaid(ctx context.Context, evt *czmodel.Paid) error {
	const op = "insert into cz_paid"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	recipientAid, err := r.addrID(ctx, evt.Recipient, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO cz_paid(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,question,recipient_aid,amount
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9)`
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		evt.RoundId, evt.Question, recipientAid, evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertPositionTransfer records one decoded TransferSingle or TransferBatch
// as one cz_position_transfer row per (id, value) pair, batch_index ordering
// them within the event.
func (r *Repo) InsertPositionTransfer(ctx context.Context, evt *czmodel.PositionTransfer) error {
	const op = "insert into cz_position_transfer"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	operatorAid, err := r.addrID(ctx, evt.Operator, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	fromAid, err := r.addrID(ctx, evt.From, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	toAid, err := r.addrID(ctx, evt.To, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO cz_position_transfer(
			evtlog_id,batch_index,block_num,tx_id,time_stamp,contract_aid,
			operator_aid,from_aid,to_aid,position_id,round_id,question,yes,amount
		) VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9,$10,$11,$12,$13,$14)`
	for i, item := range evt.Items {
		_, err = r.q(ctx).Exec(ctx, query,
			evt.EvtId, i, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
			operatorAid, fromAid, toAid,
			item.PositionId, item.RoundId, item.Question, item.Yes, item.Amount,
		)
		if err != nil {
			return store.WrapError(op, err)
		}
	}
	return nil
}

// InsertApprovalForAll records an ERC-1155 ApprovalForAll on the market.
func (r *Repo) InsertApprovalForAll(ctx context.Context, evt *czmodel.ApprovalForAll) error {
	const op = "insert into cz_approval_for_all"
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
	query := `INSERT INTO cz_approval_for_all(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			owner_aid,operator_aid,approved
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8)`
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid,
		ownerAid, operatorAid, evt.Approved,
	)
	return store.WrapError(op, err)
}
