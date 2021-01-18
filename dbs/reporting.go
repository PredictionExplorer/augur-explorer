package dbs

import (
	"fmt"
	"os"
	_  "github.com/lib/pq"

	//"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Insert_initial_report_evt(agtx *p.AugurTx,evt *p.EInitialReportSubmitted) {

	universe_id,err := ss.lookup_universe_id(evt.Universe.String())
	if err != nil {
		ss.Info.Printf("universe_mismatch: Dropping InitialReportSubmitted event for mismatch in Universe: %v",evt.Universe.String())
		os.Exit(1)
	}
	_ = universe_id
	market_aid := ss.Lookup_address_id(evt.Market.String())
	reporter_aid := ss.Lookup_or_create_address(evt.Reporter.String(),agtx.BlockNum,agtx.TxId)
	ini_reporter_aid := ss.Lookup_or_create_address(evt.InitialReporter.String(),agtx.BlockNum,agtx.TxId)

	amount_staked := evt.AmountStaked.String()
	payout_numerators := p.Bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	next_win_start := evt.NextWindowStartTime.Int64()
	next_win_end := evt.NextWindowEndTime.Int64()
	rpt_timestamp := evt.Timestamp.Int64()

	ss.Info.Printf("insert_initial_report_evt(): market_aid=%v, reporter_id=%v\n",market_aid,reporter_aid)

	market_type,mticks,_ := ss.get_market_type_and_ticks(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.PayoutNumerators)

	var query string
	query = `
		INSERT INTO report (
			block_num,
			tx_id,
			market_aid,
			aid,
			ini_reporter_aid,
			outcome_idx,
			is_initial,
			is_designated,
			amount_staked,
			pnumerators,
			description,
			next_win_start,
			next_win_end,
			rpt_timestamp
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,(` + amount_staked + `/1e+18),$9,$10,
			TO_TIMESTAMP($11),
			TO_TIMESTAMP($12),
			TO_TIMESTAMP($13)
		)`
	result,err := ss.db.Exec(query,
			agtx.BlockNum,
			agtx.TxId,
			market_aid,
			reporter_aid,
			ini_reporter_aid,
			reported_outcome,
			true,
			evt.IsDesignatedReporter,
			payout_numerators,
			evt.Description,
			next_win_start,
			next_win_end,
			rpt_timestamp)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into report table: %v,q=%v",err,query))
		os.Exit(1)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected > 0 {
		//break
	} else {
		ss.Log_msg(fmt.Sprintf("DB error: couldn't insert into InitialReport table. Rows affeced = 0"))
	}
	// set 'Reporting' status
	// ToDo: possibly migrate to triggers (or maybe not)
	ss.update_market_status(market_aid,p.MktStatusReported)
}
func (ss *SQLStorage) Delete_report_evt(tx_id int64) {

	var query string
	query = "DELETE FROM report WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_dispute_crowd_contrib(agtx *p.AugurTx,evt *p.EDisputeCrowdsourcerContribution) {

	_,err := ss.lookup_universe_id(evt.Universe.String())
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Universe %v not found on DisputeCrowdsourcererContribution event\n",evt.Universe.String()))
		os.Exit(1)
	}
	market_aid := ss.Lookup_address_id(evt.Market.String())
	reporter_aid := ss.Lookup_or_create_address(evt.Reporter.String(),agtx.BlockNum,agtx.TxId)
	signer_aid := ss.Lookup_or_create_address(agtx.From,agtx.BlockNum,agtx.TxId)
	disputed_aid := ss.Lookup_or_create_address(evt.DisputeCrowdsourcer.String(),agtx.BlockNum,agtx.TxId)

	amount_staked := evt.AmountStaked.String()
	payout_numerators := p.Bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	cur_stake := evt.CurrentStake.String()
	stake_remaining := evt.StakeRemaining.String()
	dispute_round := evt.DisputeRound.Int64()
	rpt_timestamp := evt.Timestamp.Int64()

	ss.Info.Printf("insert_dispute_crows_contrib(): market_aid=%v, reporter_id=%v, signer_aid=%v",
					market_aid,reporter_aid,signer_aid)

	market_type,mticks,_ := ss.get_market_type_and_ticks(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.PayoutNumerators)

	var query string
	query = `
		INSERT INTO report (
			block_num,
			tx_id,
			market_aid,
			aid,
			disputed_aid,
			dispute_round,
			outcome_idx,
			amount_staked,
			description,
			pnumerators,
			current_stake,
			stake_remaining,
			rpt_timestamp
		) VALUES ($1,$2,$3,$4,$5,$6,$7,`+amount_staked+`/1e+18,$8,$9,
				`+cur_stake+`/1e+18,`+stake_remaining+`/1e+18,TO_TIMESTAMP($10))`
	result,err := ss.db.Exec(query,
			agtx.BlockNum,
			agtx.TxId,
			market_aid,
			reporter_aid,
			disputed_aid,
			dispute_round,
			reported_outcome,
			evt.Description,
			payout_numerators,
			rpt_timestamp)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert dispute into report table: %v; q=%v",err,query))
		os.Exit(1)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected == 0 {
		ss.Log_msg(fmt.Sprintf("DB error: couldn't insert dispute into Report table. Rows affeced = 0"))
	}
	ss.update_market_status(market_aid,p.MktStatusDisputing)
}
