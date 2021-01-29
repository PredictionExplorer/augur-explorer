package dbs

import (
	"fmt"
	"os"
	"errors"
	"database/sql"
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
		INSERT INTO initial_report (
			block_num,
			tx_id,
			time_stamp,
			market_aid,
			reporter_aid,
			ini_reporter_aid,
			outcome_idx,
			is_designated,
			amount_staked,
			pnumerators,
			description,
			next_win_start,
			next_win_end
		) VALUES (
			$1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9::DECIMAL/1e+18,$10,$11,
			TO_TIMESTAMP($12),
			TO_TIMESTAMP($13)
		)`
	result,err := ss.db.Exec(query,
			agtx.BlockNum,
			agtx.TxId,
			rpt_timestamp,
			market_aid,
			reporter_aid,
			ini_reporter_aid,
			reported_outcome,
			evt.IsDesignatedReporter,
			amount_staked,
			payout_numerators,
			evt.Description,
			next_win_start,
			next_win_end,
	)
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
func (ss *SQLStorage) Delete_initial_report_evt(tx_id int64) {

	var query string
	query = "DELETE FROM initial_report WHERE tx_id=$1"
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
	crowdsrc_aid := ss.Lookup_or_create_address(evt.DisputeCrowdsourcer.String(),agtx.BlockNum,agtx.TxId)

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
		INSERT INTO crowdsourcer_contrib (
			block_num,
			tx_id,
			time_stamp,
			market_aid,
			reporter_aid,
			crowdsrc_aid,
			dispute_round,
			outcome_idx,
			amount_staked,
			description,
			pnumerators,
			current_stake,
			stake_remaining
		) VALUES (
			$1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,
			$9::DECIMAL/1e+18,$10,$11,$12::DECIMAL/1e+18,$13::DECIMAL/1e+18
		)`
	result,err := ss.db.Exec(query,
			agtx.BlockNum,
			agtx.TxId,
			rpt_timestamp,
			market_aid,
			reporter_aid,
			crowdsrc_aid,
			dispute_round,
			reported_outcome,
			amount_staked,
			evt.Description,
			payout_numerators,
			cur_stake,
			stake_remaining,
	)
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
func (ss *SQLStorage) Delete_crowdsourcer_contrib(tx_id int64) {

	var query string
	query = "DELETE FROM crowdsourcer_contrib WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_dispute_crowdsourcer_created(agtx *p.AugurTx,timestamp int64,evt *p.EDisputeCrowdsourcerCreated) {

	market_aid:=ss.Lookup_or_create_address(evt.Market.String(),agtx.BlockNum,agtx.TxId)
	dispute_aid:=ss.Lookup_or_create_address(evt.DisputeCrowdsourcer.String(),agtx.BlockNum,agtx.TxId)
	payouts := p.Bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	var query string
	query = "INSERT INTO crowdsourcer_created (" +
				"block_num,tx_id,time_stamp,market_aid,crowdsrc_aid,dispute_round,payout_numerators,size" +
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
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into 'crowdsourcer_created': %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_dispute_crowdsourcer_created(tx_id int64) {

	var query string
	query = "DELETE FROM crowdsourcer_created WHERE tx_id=$1"
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
func (ss *SQLStorage) Insert_reporting_participant_disavowed(agtx *p.AugurTx,timestamp int64,evt *p.EReportingParticipantDisavowed) {

	ss.Lookup_address_id(evt.Universe.String())
	reporter_aid := ss.Lookup_address_id(evt.ReportingParticipant.String())
	var query string
	query = "INSERT INTO reporter_disavowed (block_num,tx_id,time_stamp,reporter_aid) " +
				"VALUES ($1,$2,TO_TIMESTAMP($3),$4)" 

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		timestamp,
		reporter_aid,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into reporter_disavowed : %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_reporting_participant_disavowed(tx_id int64) {

	var query string
	query = "DELETE FROM reporter_disavowed WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_reporting_fee_changed(agtx *p.AugurTx,timestamp int64,evt *p.EReportingFeeChanged) {

	ss.Lookup_address_id(evt.Universe.String())
	var query string
	query = "INSERT INTO reporting_fee(block_num,tx_id,time_stamp,fee_divisor) " +
			"VALUES ($1,$2,TO_TIMESTAMP($3),$4::DECIMAL)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		timestamp,
		evt.ReportingFee.String(),
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into reporting_fee: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_reporting_fee(tx_id int64) {

	var query string
	query = "DELETE FROM reporting_fee WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_participation_tokens_redeemed(agtx *p.AugurTx,evt *p.EParticipationTokensRedeemed) {

	ss.Lookup_address_id(evt.Universe.String())
	dispwin_aid := ss.Lookup_or_create_address(evt.DisputeWindow.String(),agtx.BlockNum,agtx.TxId)
	account_aid := ss.Lookup_or_create_address(evt.Account.String(),agtx.BlockNum,agtx.TxId)

	var query string
	query = "INSERT INTO rep_tok_redeem(block_num,tx_id,time_stamp,dispwin_aid,aid,ptokens,fee_payout) " +
			"VALUES ($1,$2,TO_TIMESTAMP($3),$4,$5,$6::DECIMAL/1e+18,$7::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		evt.Timestamp.Int64(),
		dispwin_aid,
		account_aid,
		evt.AttoParticipationTokens.String(),
		evt.FeePayoutShare.String(),
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into rep_tok_redeem: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_participation_tokens_redeemed(tx_id int64) {

	var query string
	query = "DELETE FROM rep_tok_redeem WHERE tx_id=$1"
	_,err := ss.db.Exec(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_reporting_table(market_aid int64) (p.ReportingStatus,error) {

	var rst p.ReportingStatus

	var query string
	query = "SELECT num_ticks,market_type,outcomes FROM market WHERE market_aid=$1"
	var num_ticks,mkt_type int
	var outcomes string
	row := ss.db.QueryRow(query,market_aid)
	err := row.Scan(&num_ticks,&mkt_type,&outcomes)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Get_upload_block(): %v",err))
			os.Exit(1)
		}
		return rst,errors.New(fmt.Sprintf("Market not found: %v",err.Error()))
	}

	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts" +
				"time_stamp," +
				"tx.tx_hash," +
				"ira.address_id," +
				"ira.addr," +
				"ara.address_id," +
				"ara.addr," +
				"r.outcome_idx," +
				"r.is_designated," +
				"r.amount_staked/1e+18 AS amount_staked "+
			"FROM initial_report r " +
				"JOIN transaction tx ON r.tx_id=tx.id " +
				"LEFT JOIN address ira ON ini_reporter_aid=ira.address_id " +
				"LEFT JOIN address ara ON reporter_aid=ara.address_id "+
			"WHERE market_aid=$1"
	row = ss.db.QueryRow(query,market_aid)
	err=row.Scan(&num_ticks,mkt_type,outcomes)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Get_upload_block(): %v",err))
			os.Exit(1)
		}
		return rst,err
	}
	rst.InitialReport.OutcomeStr=get_outcome_str(uint8(mkt_type),rst.InitialReport.OutcomeIdx,&outcomes)

	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts" +
				"cc.time_stamp," +
				"tx.tx_hash, " +
				"cc.crowdsrc_aid, " +
				"ca.addr, " +
				"cc.dispute_round, " +
				"cc.payout_numerators, " +
				"cc.size/1e+18 " +
			"FROM crowdsourcer_created cc " +
				"JOIN transaction tx ON cc.tx_id=tx.id " +
				"LEFT JOIN address ca ON cc.crowdsrc_aid=ca.address_id " +
			"WHERE market_aid = $1 " +
			"ORDER BY cc.time_stamp"

	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	disputes := make([]p.DisputeInfo,0,8)

	defer rows.Close()
	for rows.Next() {
		var rec p.DisputeInfo
		var payout_numerators string
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TxHash,
			&rec.CrowdsourcerAid,
			&rec.CrowdsourcerAddr,
			&rec.DisputeRound,
			&payout_numerators,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		outcidx := get_outcome_idx_from_numerators(mkt_type,int64(num_ticks),payout_numerators)
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),outcidx,&outcomes)
		disputes = append(disputes,rec)
	}
	rst.Disputes = disputes

	return rst,nil
}
