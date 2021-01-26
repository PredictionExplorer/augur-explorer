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
func (ss *SQLStorage) Insert_dispute_crowdsourcer_created(agtx *p.AugurTx,timestamp int64,evt *p.EDisputeCrowdsourcerCreated) {

	market_aid:=ss.Lookup_or_create_address(evt.Market.String(),agtx.BlockNum,agtx.TxId)
	dispute_aid:=ss.Lookup_or_create_address(evt.DisputeCrowdsourcer.String(),agtx.BlockNum,agtx.TxId)
	payouts := p.Bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	var query string
	query = "INSERT INTO crowdsourcer_created (" +
				"block_num,tx_id,time_stamp,market_aid,dispute_aid,dispute_round,payout_numerators,size" +
				") VALUES ($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		timestamp,
		market_aid,
		dispute_aid,
		evt.DisputeRound.Int64(),
		payouts,
		evt.Size.String(),
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into dispute_created table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_dispute_crowdsourcer_created(tx_id int64) {

	var query string
	query = "DELETE FROM dispute_created WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_dispute_window_created(agtx *p.AugurTx,evt *p.EDisputeWindowCreated) {

	universe_id:=ss.Lookup_address_id(evt.Universe.String())
	window_aid:=ss.Lookup_or_create_address(evt.DisputeWindow.String(),agtx.BlockNum,agtx.TxId)
	var query string
	query = "INSERT INTO dispute_window (" +
				"block_num,tx_id,universe_id,wid,window_aid,start_time,end_time,initial" +
			") VALUES ($1,$2,$3,$4,$5,TO_TIMESTAMP($6),TO_TIMESTAMP($7),$8)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		universe_id,
		evt.Id.Int64(),
		window_aid,
		evt.StartTime.Int64(),
		evt.EndTime.Int64(),
		evt.Initial,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into dispute_created table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_dispute_window_created(tx_id int64) {

	var query string
	query = "DELETE FROM dispute_window WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_designated_report_stake_changed(agtx *p.AugurTx,evt *p.EDesignatedReportStakeChanged) {

	universe_id:=ss.Lookup_address_id(evt.Universe.String())
	var query string
	query = "INSERT INTO rep_stake_chg (" +
				"block_num,tx_id,universe_id,rep_stake" +
				") VALUES ($1,$2,$3,$4::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		universe_id,
		evt.DesignatedReportStake.String(),
		)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into rep_stake_chg table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_designated_report_stake_changed(tx_id int64) {

	var query string
	query = "DELETE FROM rep_stake_chg WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_initial_reporter_redeemed(agtx *p.AugurTx,evt *p.EInitialReporterRedeemed) {

	market_aid := ss.Lookup_address_id(evt.Market.String())
	reporter_aid := ss.Lookup_or_create_address(evt.Reporter.String(),agtx.BlockNum,agtx.TxId)
	ini_rep_aid := ss.Lookup_or_create_address(evt.InitialReporter.String(),agtx.BlockNum,agtx.TxId)
	payout_numerators := p.Bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	var query string
	query = "INSERT INTO irep_redeem (" +
				"block_num,tx_id,market_aid,reporter_aid,ini_rep_aid,time_stamp,amount,rep,payout_numerators" +
				") VALUES ($1,$2,$3,$4,$5,TO_TIMESTAMP($6),$7::DECIMAL/1e+18,$8::DECIMAL/1e+18,$9)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		market_aid,
		reporter_aid,
		ini_rep_aid,
		evt.Timestamp.Int64(),
		evt.AmountRedeemed.String(),
		evt.RepReceived.String(),
		payout_numerators,
		)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into irep_redeem table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_initial_reporter_redeemed(tx_id int64) {

	var query string
	query = "DELETE FROM irep_redeem WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_dispute_crowdsourcer_redeemed(agtx *p.AugurTx,evt *p.EDisputeCrowdsourcerRedeemed) {

	ss.Lookup_address_id(evt.Universe.String())
	market_aid := ss.Lookup_address_id(evt.Market.String())
	reporter_aid := ss.Lookup_or_create_address(evt.Reporter.String(),agtx.BlockNum,agtx.TxId)
	crowdsourcer_aid := ss.Lookup_or_create_address(evt.DisputeCrowdsourcer.String(),agtx.BlockNum,agtx.TxId)
	payout_numerators := p.Bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	var query string
	query = "INSERT INTO crowdsourcer_redeemed (" +
				"block_num,tx_id,market_aid,reporter_aid,crowdsourcer_aid," +
				"time_stamp,amount,rep,payout_numerators" +
			") VALUES ($1,$2,$3,$4,$5,TO_TIMESTAMP($6),$7::DECIMAL/1e+18,$8::DECIMAL/1e+18,$9)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		market_aid,
		reporter_aid,
		crowdsourcer_aid,
		evt.Timestamp.Int64(),
		evt.AmountRedeemed.String(),
		evt.RepReceived.String(),
		payout_numerators,
		)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into crowdsourcer_redeemed : %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_dispute_crowdsourcer_redeemed(tx_id int64) {

	var query string
	query = "DELETE FROM crowdsourcer_redeemed WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_dispute_crowdsourcer_completed(agtx *p.AugurTx,evt *p.EDisputeCrowdsourcerCompleted) {

	ss.Lookup_address_id(evt.Universe.String())
	market_aid := ss.Lookup_address_id(evt.Market.String())
	crowdsourcer_aid := ss.Lookup_or_create_address(evt.DisputeCrowdsourcer.String(),agtx.BlockNum,agtx.TxId)
	payout_numerators := p.Bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	var query string
	query = "INSERT INTO crowdsourcer_completed (" +
				"block_num,tx_id,time_stamp,market_aid,crowdsrc_aid,next_win_start,next_win_end," +
				"dispute_round,pacing_on,tot_rep_payout,tot_rep_market,payout_numerators" +
			") VALUES (" +
				"$1,$2,TO_TIMESTAMP($3),$4,$5,TO_TIMESTAMP($6),TO_TIMESTAMP($7),"+
				"$8,$9,$10::DECIMAL/1e+18,$11::DECIMAL/1e+18,$12"+
			")"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		evt.Timestamp.Int64(),
		market_aid,
		crowdsourcer_aid,
		evt.NextWindowStartTime.Int64(),
		evt.NextWindowEndTime.Int64(),
		evt.DisputeRound.Int64(),
		evt.PacingOn,
		evt.TotalRepStakedInPayout.String(),
		evt.TotalRepStakedInMarket.String(),
		payout_numerators,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into crowdsourcer_completed : %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_dispute_crowdsourcer_completed(tx_id int64) {

	var query string
	query = "DELETE FROM crowdsourcer_completed WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}


func (ss *SQLStorage) Get_reporting_status(market_aid int64) p.ReportingStatus {

	var repst p.ReportingStatus
	return repst
}
