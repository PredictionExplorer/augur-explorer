package dbs

import (
	"fmt"
	"os"
	"math/big"
	"errors"
	"strings"
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

	market_type,mticks,lo_price,_ := ss.get_market_type_ticks_lo_price(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.PayoutNumerators)
	scalar_val := float64(0)
	if market_type == 2 { // scalar
		if reported_outcome != 0 {
			scalar_val = (float64(lo_price) + float64(evt.PayoutNumerators[2].Int64()))
		}
	}

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
			scalar_val,
			is_designated,
			amount_staked,
			pnumerators,
			description,
			next_win_start,
			next_win_end
		) VALUES (
			$1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10::DECIMAL/1e+18,$11,$12,
			TO_TIMESTAMP($13),
			TO_TIMESTAMP($14)
		)`
	result,err := ss.db.Exec(query,
			agtx.BlockNum,
			agtx.TxId,
			rpt_timestamp,
			market_aid,
			reporter_aid,
			ini_reporter_aid,
			reported_outcome,
			scalar_val,
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

	market_type,mticks,lo_price,_ := ss.get_market_type_ticks_lo_price(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.PayoutNumerators)
	scalar_val := float64(0)
	if market_type == 2 { // scalar
		if reported_outcome != 0 {
			scalar_val = (float64(lo_price) + float64(evt.PayoutNumerators[2].Int64()))
		}
	}

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
			scalar_val,
			amount_staked,
			description,
			pnumerators,
			current_stake,
			stake_remaining
		) VALUES (
			$1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,
			$10::DECIMAL/1e+18,$11,$12,$13::DECIMAL/1e+18,$14::DECIMAL/1e+18
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
			scalar_val,
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
	market_type,mticks,lo_price,_ := ss.get_market_type_ticks_lo_price(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.PayoutNumerators)
	scalar_val := float64(0)
	if market_type == 2 { // scalar
		if reported_outcome != 0 {
			scalar_val = (float64(lo_price) + float64(evt.PayoutNumerators[2].Int64()))
		}
	}
	var query string
	query = "INSERT INTO crowdsourcer_created (" +
				"block_num,tx_id,time_stamp,market_aid,crowdsrc_aid," +
				"dispute_round,outcome_idx,scalar_val,payout_numerators,size" +
			") VALUES ($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		timestamp,
		market_aid,
		dispute_aid,
		evt.DisputeRound.Int64(),
		reported_outcome,
		scalar_val,
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
	market_type,mticks,lo_price,_ := ss.get_market_type_ticks_lo_price(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.PayoutNumerators)
	scalar_val := float64(0)
	if market_type == 2 { // scalar
		if reported_outcome != 0 {
			scalar_val = (float64(lo_price) + float64(evt.PayoutNumerators[2].Int64()))
		}
	}

	var query string
	query = "INSERT INTO irep_redeem (" +
				"block_num,tx_id,market_aid,reporter_aid,ini_rep_aid,time_stamp,"+
				"outcome_idx,scalar_val,amount,rep,payout_numerators" +
				") VALUES ($1,$2,$3,$4,$5,TO_TIMESTAMP($6),$7,$8,$9::DECIMAL/1e+18,$10::DECIMAL/1e+18,$11)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		market_aid,
		reporter_aid,
		ini_rep_aid,
		evt.Timestamp.Int64(),
		reported_outcome,
		scalar_val,
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
	market_type,mticks,lo_price,_ := ss.get_market_type_ticks_lo_price(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.PayoutNumerators)
	scalar_val := float64(0)
	if market_type == 2 { // scalar
		if reported_outcome != 0 {
			scalar_val = (float64(lo_price) + float64(evt.PayoutNumerators[2].Int64()))
		}
	}

	var query string
	query = "INSERT INTO crowdsourcer_redeemed (" +
				"block_num,tx_id,market_aid,reporter_aid,crowdsourcer_aid," +
				"time_stamp,outcome_idx,scalar_val,amount,rep,payout_numerators" +
			") VALUES ($1,$2,$3,$4,$5,TO_TIMESTAMP($6),$7,$8,$9::DECIMAL/1e+18,$10::DECIMAL/1e+18,$11)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		market_aid,
		reporter_aid,
		crowdsourcer_aid,
		evt.Timestamp.Int64(),
		reported_outcome,
		scalar_val,
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
	market_type,mticks,lo_price,_ := ss.get_market_type_ticks_lo_price(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.PayoutNumerators)
	scalar_val := float64(0)
	if market_type == 2 { // scalar
		if reported_outcome != 0 {
			scalar_val = (float64(lo_price) + float64(evt.PayoutNumerators[2].Int64()))
		}
	}
	var query string
	query = "INSERT INTO crowdsourcer_completed (" +
				"block_num,tx_id,time_stamp,market_aid,crowdsrc_aid,next_win_start,next_win_end," +
				"dispute_round,outcome_idx,scalar_val,pacing_on,tot_rep_payout,tot_rep_market,payout_numerators" +
			") VALUES (" +
				"$1,$2,TO_TIMESTAMP($3),$4,$5,TO_TIMESTAMP($6),TO_TIMESTAMP($7),"+
				"$8,$9,$10,$11,$12::DECIMAL/1e+18,$13::DECIMAL/1e+18,$14"+
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
		reported_outcome,
		scalar_val,
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
func payout_numerators_text_to_big(payout_numerators_str string) []*big.Int {

	output := make([]*big.Int,0,8)
	tokens := strings.Split(payout_numerators_str,",")
	for _,t := range tokens {
		b := new(big.Int)
		b.SetString(t,10)
		output = append(output,b)
	}
	return output
}
func (ss *SQLStorage) get_dispute_contributions(crowdsourcer_aid int64,mkt_type uint8,num_ticks int,outcomes *string) []p.DisputeContribution {

	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"time_stamp," +
				"tx.tx_hash," +
				"ra.address_id," +
				"ra.addr," +
				"cc.dispute_round," +
				"m.market_type," +
				"m.decimals," +
				"cc.outcome_idx," +
				"cc.scalar_val," +
				"cc.amount_staked," +
				"cc.current_stake," +
				"cc.stake_remaining " +
			"FROM crowdsourcer_contrib cc " +
				"JOIN transaction tx ON cc.tx_id=tx.id " +
				"JOIN market m ON cc.market_aid=m.market_aid "+
				"LEFT JOIN address ra ON cc.reporter_aid=ra.address_id " +
			"WHERE crowdsrc_aid = $1 " +
			"ORDER BY cc.time_stamp"
	
	rows,err := ss.db.Query(query,crowdsourcer_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	contribs := make([]p.DisputeContribution,0,8)

	defer rows.Close()
	for rows.Next() {
		var rec p.DisputeContribution
		var decimals int
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TxHash,
			&rec.ReporterAid,
			&rec.ReporterAddr,
			&rec.DisputeRound,
			&rec.MktType,
			&decimals,
			&rec.OutcomeIdx,
			&rec.ScalarValue,
			&rec.AmountStaked,
			&rec.CurrentStake,
			&rec.StakeRemaining,
		)

		p.Augur_UI_price_adjustments(&rec.ScalarValue,nil,rec.MktType,decimals)
		rec.TxHashSh = p.Short_hash(rec.TxHash)
		rec.ReporterAddrSh = p.Short_address(rec.ReporterAddr)
		rec.OutcomeStr=get_outcome_str(mkt_type,rec.OutcomeIdx,outcomes)
		contribs = append(contribs,rec)
	}
	return contribs
}
func (ss *SQLStorage) Get_key_market_fields(market_aid int64) (num_ticks int,mkt_type int,outcomes string,err error) {

	var query string
	query = "SELECT num_ticks,market_type,outcomes FROM market WHERE market_aid=$1"
	row := ss.db.QueryRow(query,market_aid)
	err = row.Scan(&num_ticks,&mkt_type,&outcomes)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Get_reporting_table(): %v : %v",err,query))
			os.Exit(1)
		}
		err = errors.New(fmt.Sprintf("Market not found: %v",err.Error()))
		return
	}
	return
}
func (ss *SQLStorage) Get_reporting_table(market_aid int64) (p.ReportingStatus,error) {

	var rst p.ReportingStatus

	num_ticks,mkt_type,outcomes,err := ss.Get_key_market_fields(market_aid)
	if err != nil {
		return rst,err
	}


	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts," +
				"r.time_stamp," +
				"tx.tx_hash," +
				"ira.address_id," +
				"ira.addr," +
				"ara.address_id," +
				"ara.addr," +
				"m.market_type," +
				"m.decimals," +
				"r.outcome_idx," +
				"r.scalar_val," +
				"r.is_designated," +
				"r.amount_staked AS amount_staked "+
			"FROM initial_report r " +
				"JOIN transaction tx ON r.tx_id=tx.id " +
				"JOIN market m ON r.market_aid=m.market_aid " +
				"LEFT JOIN address ira ON r.ini_reporter_aid=ira.address_id " +
				"LEFT JOIN address ara ON r.reporter_aid=ara.address_id "+
			"WHERE r.market_aid=$1"
	row := ss.db.QueryRow(query,market_aid)
	var decimals int
	err=row.Scan(
		&rst.InitialReport.TimeStamp,
		&rst.InitialReport.DateTime,
		&rst.InitialReport.TxHash,
		&rst.InitialReport.InitialReporterAid,
		&rst.InitialReport.InitialReporterAddr,
		&rst.InitialReport.ActualReporterAid,
		&rst.InitialReport.ActualReporterAddr,
		&rst.InitialReport.MktType,
		&decimals,
		&rst.InitialReport.OutcomeIdx,
		&rst.InitialReport.ScalarValue,
		&rst.InitialReport.IsDesignated,
		&rst.InitialReport.AmountStaked,
	)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Get_reporting_table(): %v : %v",err,query))
			os.Exit(1)
		}
		return rst,err
	}
	p.Augur_UI_price_adjustments(&rst.InitialReport.ScalarValue,nil,rst.InitialReport.MktType,decimals)
	rst.InitialReport.TxHashSh = p.Short_hash(rst.InitialReport.TxHash)
	rst.InitialReport.OutcomeStr=get_outcome_str(uint8(mkt_type),rst.InitialReport.OutcomeIdx,&outcomes)
	query = "SELECT " +
				"EXTRACT(EPOCH FROM cr.time_stamp)::BIGINT created_ts," +
				"TO_CHAR(cr.time_stamp,'dd/mm/yyyy HH:ii')," +
				"EXTRACT(EPOCH FROM co.time_stamp)::BIGINT completed_ts," +
				"TO_CHAR(co.time_stamp,'dd/mm/yyyy HH:ii')," +
				"cr_addr.addr, "+
				"cr_tx.tx_hash," +
				"co_tx.tx_hash," +
				"cr.crowdsrc_aid," +
				"cr.dispute_round rstart,"+	// round start (round in Augur = number of participants)
				"co.dispute_round rend," +	// round end
				"m.market_type," +
				"m.decimals," +
				"cr.outcome_idx," +
				"cr.scalar_val," +
				"round(cr.size,5) min_size," +
				"co.pacing_on p," +
				"co.tot_rep_payout," +
				"co.tot_rep_market, " +
				"EXTRACT(EPOCH FROM d.start_time)::BIGINT dwin_start_ts," +
				"d.start_time dwin_start," +
				"EXTRACT(EPOCH FROM d.end_time)::BIGINT dwin_end_ts," +
				"d.end_time dwin_end," +
				"d.window_aid dwin_aid " +
			"FROM crowdsourcer_created cr " +
				"JOIN market m ON cr.market_aid=m.market_aid " +
				"LEFT JOIN transaction co_tx ON cr.tx_id=co_tx.id " +
				"LEFT JOIN transaction cr_tx ON cr.tx_id=cr_tx.id " +
				"LEFT JOIN dispute_window d ON cr.dispute_win_id = d.id " +
				"LEFT JOIN crowdsourcer_completed co ON cr.crowdsrc_aid=co.crowdsrc_aid " +
				"LEFT JOIN address cr_addr ON cr.crowdsrc_aid=cr_addr.address_id " +
			"WHERE cr.market_aid=$1 " +
			"ORDER BY rend,co.time_stamp,rstart,cr.time_stamp"

	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	disputes := make([]p.DisputeInfo,0,8)


	defer rows.Close()
	for rows.Next() {
		var rec p.DisputeInfo
		var null_completed_ts,null_round_end sql.NullInt64
		var null_completed_date,null_tx_hash sql.NullString
		var null_rep_payout,null_rep_market sql.NullFloat64
		var null_pacing sql.NullBool
		var decimals int
		err=rows.Scan(
			&rec.CreatedTs,
			&rec.CreatedDate,
			&null_completed_ts,
			&null_completed_date,
			&rec.CrowdsourcerAddr,
			&rec.CreatedTxHash,
			&null_tx_hash,
			&rec.CrowdsourcerAid,
			&rec.DisputeRoundStart,
			&null_round_end,
			&rec.MktType,
			&decimals,
			&rec.OutcomeIdx,
			&rec.ScalarValue,
			&rec.MinDisputeSize,
			&null_pacing,
			&null_rep_payout,
			&null_rep_market,
			&rec.WindowStartTs,
			&rec.WindowStartDate,
			&rec.WindowEndTs,
			&rec.WindowEndDate,
			&rec.DisputeWindowAid,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.ScalarValue,nil,rec.MktType,decimals)
		if null_completed_ts.Valid { rec.CompletedTs = null_completed_ts.Int64; rec.Completed=true }
		if null_completed_date.Valid { rec.CompletedDate = null_completed_date.String }
		if null_tx_hash.Valid { rec.CompletedTxHash = null_tx_hash.String }
		if null_round_end.Valid { rec.DisputeRoundEnd = int(null_round_end.Int64) }
		if null_rep_payout.Valid { rec.TotalRepPayout = null_rep_payout.Float64 }
		if null_rep_market.Valid { rec.RepInMarket = null_rep_market.Float64 }
		if null_pacing.Valid { rec.PacingOn = null_pacing.Bool }
		rec.CreatedTxHashSh = p.Short_hash(rec.CreatedTxHash)
		rec.CompletedTxHashSh = p.Short_hash(rec.CompletedTxHash)
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.OutcomeIdx,&outcomes)
		rec.Contributions = ss.get_dispute_contributions(rec.CrowdsourcerAid,uint8(mkt_type),num_ticks,&outcomes)
		disputes = append(disputes,rec)
	}

	rst.Disputes = disputes

	return rst,nil
}
func update_current_round(rr *p.RoundsRow,rec *p.DisputeRound) {
		rr.Rounds.MarketRep = rec.MarketRep
		rr.Rounds.TimeStamp = rec.TimeStamp
		rr.Rounds.DateTime = rec.DateTime
		rr.Rounds.WindowStartDate = rec.WindowStartDate
		rr.Rounds.WindowEndDate = rec.WindowEndDate
		rr.Rounds.WindowStartTs = rec.WindowStartTs
		rr.Rounds.WindowEndTs = rec.WindowEndTs
		rr.Rounds.CompletedTs = rec.CompletedTs
		rr.Rounds.CompletedDate = rec.CompletedDate
		rr.Rounds.MktType = rec.MktType
		rr.Rounds.ScalarValue = rec.ScalarValue
}
func (ss *SQLStorage) Get_round_table(market_aid int64) ([]p.RoundsRow,int,string,[]float64) {

	scalar_values := make([]float64,0,4)
	_,mkt_type,outcomes,_ := ss.Get_key_market_fields(market_aid)
	outcomes = adjust_outcomes_str(mkt_type,outcomes)
	num_outcomes := 3	// Invalid,No,Yes
	if mkt_type == 1 { // categorical
		num_outcomes = len(strings.Split(outcomes,","))
	}
	if mkt_type == 2 { // scalar
		num_outcomes = 2
	}

	rounds := make([]p.RoundsRow,0,8)
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts," +
				"TO_CHAR(time_stamp, 'dd/mm/yyyy HH:ii')," +
				"tx.tx_hash," +
				"ira.address_id," +
				"ira.addr," +
				"ara.address_id," +
				"ara.addr," +
				"m.market_type," +
				"m.decimals," +
				"r.outcome_idx," +
				"r.scalar_val," +
				"r.is_designated," +
				"r.amount_staked AS amount_staked,"+
				"EXTRACT(EPOCH FROM m.end_time)::BIGINT win_start_ts," +
				"TO_CHAR(m.end_time, 'dd/mm/yyyy HH:ii') win_start_date," +
				"EXTRACT(EPOCH FROM m.end_time+interval '1' day)::BIGINT win_end_ts," +
				"TO_CHAR(m.end_time+interval '1' day, 'dd/mm/yyyy HH:ii') win_end_date " +
			"FROM initial_report r " +
				"JOIN transaction tx ON r.tx_id=tx.id " +
				"JOIN market m ON r.market_aid=m.market_aid " +
				"LEFT JOIN address ira ON r.ini_reporter_aid=ira.address_id " +
				"LEFT JOIN address ara ON r.reporter_aid=ara.address_id "+
			"WHERE r.market_aid=$1"
	row := ss.db.QueryRow(query,market_aid)
	var inirep p.InitialReportInfo
	var decimals int
	err := row.Scan(
		&inirep.TimeStamp,
		&inirep.DateTime,
		&inirep.TxHash,
		&inirep.InitialReporterAid,
		&inirep.InitialReporterAddr,
		&inirep.ActualReporterAid,
		&inirep.ActualReporterAddr,
		&inirep.MktType,
		&decimals,
		&inirep.OutcomeIdx,
		&inirep.ScalarValue,
		&inirep.IsDesignated,
		&inirep.AmountStaked,
		&inirep.WinStartTs,
		&inirep.WinStartDate,
		&inirep.WinEndTs,
		&inirep.WinEndDate,
	)
	if (err == nil) {
		p.Augur_UI_price_adjustments(&inirep.ScalarValue,nil,inirep.MktType,decimals)
		if inirep.TimeStamp != 0 {
			var rr p.RoundsRow
			var rec p.DisputeRound
			rr.Rounds.RoundNum = -1
			rr.Rounds.ORounds = make([]p.DisputeRound,0,8)
			rr.Rounds.MktType = inirep.MktType

			for i:=0; i<num_outcomes ; i++ {
				var empty_rec p.DisputeRound
				rec.OutcomeIdx=inirep.OutcomeIdx
				if (inirep.MktType == 2) && (rec.OutcomeIdx==2){ // scalar
					rec.OutcomeIdx=1
				}
				if i==rec.OutcomeIdx {
					rec.TimeStamp = inirep.TimeStamp
					rec.DateTime = inirep.DateTime
					rec.OutcomeIdx = inirep.OutcomeIdx
					rec.RepPayout = inirep.AmountStaked
					rec.ScalarValue = inirep.ScalarValue
					rec.RoundNum = -1
					rec.Color = true
					rec.Completed = true
					rr.Rounds.ORounds = append(rr.Rounds.ORounds,rec)
					rr.Rounds.MarketRep = inirep.AmountStaked
					rr.Rounds.TimeStamp = inirep.TimeStamp
					rr.Rounds.DateTime = inirep.DateTime
					rr.Rounds.WindowStartTs = inirep.WinStartTs
					rr.Rounds.WindowEndTs = inirep.WinEndTs
					rr.Rounds.WindowStartDate = inirep.WinStartDate
					rr.Rounds.WindowEndDate = inirep.WinEndDate
					rr.Rounds.CompletedTs = rec.TimeStamp
					rr.Rounds.CompletedDate = rec.DateTime
					rr.Rounds.Completed = true
					rr.Rounds.MktType = inirep.MktType
					rr.Rounds.ScalarValue = inirep.ScalarValue
					scalar_values = append(scalar_values,inirep.ScalarValue)
				} else {
					rr.Rounds.ORounds = append(rr.Rounds.ORounds,empty_rec)
				}
			}
			rounds = append(rounds,rr)
		}
	} else {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	query = "SELECT " +
				"EXTRACT(EPOCH FROM cr.time_stamp)::BIGINT ts," +
				"TO_CHAR(cr.time_stamp, 'dd/mm/yyyy HH:ii')," +
				"EXTRACT(EPOCH FROM co.time_stamp)::BIGINT ts," +
				"TO_CHAR(co.time_stamp, 'dd/mm/yyyy HH:ii')," +
				"cr.size," +
				"co.tot_rep_payout," +
				"co.tot_rep_market," +
				"cr.dispute_round," +
				"co.dispute_round," +
				"co.pacing_on, " +
				"m.market_type, " +
				"m.decimals," +
				"cr.outcome_idx, " +
				"cr.scalar_val, " +
				"EXTRACT(EPOCH FROM d.start_time)::BIGINT dwin_start_ts," +
				"TO_CHAR(d.start_time,'dd/mm/yyyy HH:ii') dwin_start," +
				"EXTRACT(EPOCH FROM d.end_time)::BIGINT dwin_end_ts," +
				"TO_CHAR(d.end_time,'dd/mm/yyyy HH:ii') dwin_end," +
				"d.window_aid dwin_aid " +
			"FROM crowdsourcer_created cr " +
				"JOIN dispute_window d ON cr.dispute_win_id=d.id " +
				"JOIN market m ON cr.market_aid = m.market_aid " +
				"LEFT JOIN crowdsourcer_completed co ON co.crowdsrc_aid=cr.crowdsrc_aid " +
			"WHERE cr.market_aid=$1 " +
			"ORDER BY co.dispute_round,cr.dispute_round"
	fmt.Printf("query=%v\n",query)
	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}


	defer rows.Close()
	for rows.Next() {
		var rr p.RoundsRow
		var rec p.DisputeRound
		var null_rep_payout,null_rep_market sql.NullFloat64
		var null_dispute_round,null_ts sql.NullInt64
		var null_pacing sql.NullBool
		var null_date sql.NullString
		var decimals int
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&null_ts,
			&null_date,
			&rec.MinDisputeSize,
			&null_rep_payout,
			&null_rep_market,
			&rec.RoundNum,
			&null_dispute_round,
			&null_pacing,
			&rec.MktType,
			&decimals,
			&rec.OutcomeIdx,
			&rec.ScalarValue,
			&rec.WindowStartTs,
			&rec.WindowStartDate,
			&rec.WindowEndTs,
			&rec.WindowEndDate,
			&rec.DisputeWinAid,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		p.Augur_UI_price_adjustments(&rec.ScalarValue,nil,rec.MktType,decimals)

		if null_ts.Valid {
			rec.CompletedTs = null_ts.Int64
			rec.CompletedDate = null_date.String
		}
		if null_rep_payout.Valid { rec.RepPayout = null_rep_payout.Float64 }
		if null_rep_market.Valid { rec.MarketRep = null_rep_market.Float64 }
		if null_dispute_round.Valid { rec.RoundNum = int(null_dispute_round.Int64) }
		if null_pacing.Valid { rec.PacingOn = null_pacing.Bool }
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.OutcomeIdx,&outcomes)

		if null_rep_payout.Valid {
			rr.Rounds.Completed = true
		}
		//var outc_rounds p.OutcomeRounds
		rr.Rounds.RoundNum = rec.RoundNum
		rr.Rounds.ORounds = make([]p.DisputeRound,0,8)
		if rec.MktType == 2  {
			current_seen := false
			for _,entry := range scalar_values {
				if rec.ScalarValue == entry {
					current_seen = true
				}
			}
			if !current_seen {
				// we append at the beginning because we need the loop to include current rec
				if rec.OutcomeIdx != 0 { // if is not invalid
					scalar_values = append(scalar_values,rec.ScalarValue)
				}
			}
			// for scalar markets outcome_idx is ignored because it is always 2
			//	but Invalid outcome is present, so we will check only that one
			if rec.OutcomeIdx == 0 { // invalid
				rec.Color = true
				rr.Rounds.ORounds = append(rr.Rounds.ORounds,rec)
				update_current_round(&rr,&rec)
			} else {
				var empty_rec p.DisputeRound
				// add empty record for invalid entry
				if len(rounds) > 0 {
					prev_rec := &rounds[len(rounds)-1].Rounds.ORounds[0]
					if prev_rec.TimeStamp != 0 {
						empty_rec = *prev_rec
						empty_rec.Color = false
						empty_rec.PacingOn = false
						empty_rec.Completed = false
					}
				}
				rr.Rounds.ORounds = append(rr.Rounds.ORounds,empty_rec)
			}
			// now we parse entries using scalar value as distinctive property
			// (because the outcome = 2 for scalar markets)
			num_svalues := len(scalar_values)
			for j:=0; j<num_svalues; j++ {
				if (scalar_values[j] == rec.ScalarValue) && (rec.OutcomeIdx!=0) {
					rec.Color = true
					rr.Rounds.ORounds = append(rr.Rounds.ORounds,rec)
					update_current_round(&rr,&rec)
				} else {
					var empty_rec p.DisputeRound
					if len(rounds) > 0 {
						prev_rec := &rounds[len(rounds)-1].Rounds.ORounds[j+1]// +1 for invalid
						if prev_rec.TimeStamp != 0 {
							empty_rec = *prev_rec
							empty_rec.Color = false
							empty_rec.PacingOn = false
							empty_rec.Completed = false
						}
					}
					rr.Rounds.ORounds = append(rr.Rounds.ORounds,empty_rec)
				}
			}
		} else {
			for i:=0; i<num_outcomes ; i++ {
				if i==rec.OutcomeIdx {
					rec.Color = true
					rr.Rounds.ORounds = append(rr.Rounds.ORounds,rec)
					// update current
					update_current_round(&rr,&rec)
				} else {
					var empty_rec p.DisputeRound
					if len(rounds) > 0 {
						prev_rec := &rounds[len(rounds)-1].Rounds.ORounds[i]
						if prev_rec.TimeStamp != 0 {
							empty_rec = *prev_rec
							empty_rec.Color = false
							empty_rec.PacingOn = false
							empty_rec.Completed = false
						}
					}
					rr.Rounds.ORounds = append(rr.Rounds.ORounds,empty_rec)
				}
			}
		}
		rounds = append(rounds,rr)
	}
	var prev_win_start int64 = 0
	var prev_win_end int64 = 0
	var widx int = 0
	// calculate window number
	var i int = 0
	var wcounter int = 1
	for ; i< len(rounds); i++ {
		if prev_win_start == 0 {
			prev_win_start = rounds[i].Rounds.WindowStartTs
			prev_win_end = rounds[i].Rounds.WindowEndTs
		}
		if (prev_win_start == rounds[i].Rounds.WindowStartTs) && (prev_win_end == rounds[i].Rounds.WindowEndTs) {
			continue // same window 
		}
		// window changed
		rounds[widx].Rounds.WindowSpan = i-widx
		rounds[widx].Rounds.WindowNum = wcounter
		widx = i
		prev_win_start = rounds[widx].Rounds.WindowStartTs
		prev_win_end = rounds[widx].Rounds.WindowEndTs
		wcounter++
	}
	rounds[widx].Rounds.WindowSpan = i-widx
	rounds[widx].Rounds.WindowNum = wcounter
	return rounds,num_outcomes,outcomes,scalar_values
}
func (ss *SQLStorage) Get_initial_report_redeemed_record(market_aid int64) *p.IniRepRedeemed {

	rec := new(p.IniRepRedeemed)
	var query string
	var decimals int
	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"TO_CHAR(ir.time_stamp, 'dd/mm/yyyy HH:ii')," +
				"ir.ini_rep_aid," +
				"aini.addr," +
				"ir.reporter_aid," +
				"arep.addr," +
				"m.market_type," +
				"m.decimals," +
				"outcome_idx," +
				"scalar_val," +
				"amount, " +
				"rep," +
				"tx.tx_hash " +
			"FROM irep_redeem ir " +
				"JOIN transaction tx ON ir.tx_id=tx.id " +
				"JOIN market m ON ir.market_aid=m.market_aid " +
				"LEFT JOIN address aini ON ir.ini_rep_aid=aini.address_id " +
				"LEFT JOIN address arep ON ir.reporter_aid=arep.address_id " +
			"WHERE ir.market_aid=$1"
	row := ss.db.QueryRow(query,market_aid)
	err := row.Scan(
		&rec.TimeStamp,
		&rec.DateTime,
		&rec.InitialReporterAid,
		&rec.InitialReporterAddr,
		&rec.ReporterAid,
		&rec.ReporterAddr,
		&rec.MktType,
		&decimals,
		&rec.OutcomeIdx,
		&rec.ScalarValue,
		&rec.Amount,
		&rec.RepReceived,
		&rec.TxHash,
	)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Get_reporting_table(): %v : %v",err,query))
			os.Exit(1)
		}
		return nil
	}
	p.Augur_UI_price_adjustments(&rec.ScalarValue,nil,rec.MktType,decimals)
	_,mkt_type,outcomes,_ := ss.Get_key_market_fields(market_aid)
	rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.OutcomeIdx,&outcomes)
	rec.TxHashSh = p.Short_hash(rec.TxHash)
	return rec
}
func (ss *SQLStorage) Get_redeemed_participants(market_aid int64) []p.RedeemedParticipant {

	records := make([]p.RedeemedParticipant,0,8)
	_,mkt_type,outcomes,_ := ss.Get_key_market_fields(market_aid)

	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"TO_CHAR(c.time_stamp, 'dd/mm/yyyy HH:ii')," +
				"c.reporter_aid," +
				"ra.addr," +
				"m.market_type," +
				"m.decimals," +
				"outcome_idx," +
				"c.scalar_val," +
				"amount, " +
				"rep," +
				"tx.tx_hash " +
			"FROM crowdsourcer_redeemed c " +
				"JOIN transaction tx ON tx_id=tx.id " +
				"JOIN market m ON c.market_aid=m.market_aid " +
				"LEFT JOIN address ra ON c.reporter_aid=ra.address_id " +
			"WHERE c.market_aid=$1 AND (amount < rep)"

	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RedeemedParticipant
		var decimals int
		err := rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.ReporterAid,
			&rec.ReporterAddr,
			&rec.MktType,
			&decimals,
			&rec.OutcomeIdx,
			&rec.ScalarValue,
			&rec.RepInvested,
			&rec.RepReturned,
			&rec.TxHash,
		)
		if (err!=nil) {
			if err != sql.ErrNoRows {
				ss.Log_msg(fmt.Sprintf("Error in Get_reporting_table(): %v : %v",err,query))
				os.Exit(1)
			}
			return records
		}
		p.Augur_UI_price_adjustments(&rec.ScalarValue,nil,rec.MktType,decimals)
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.OutcomeIdx,&outcomes)
		rec.TxHashSh = p.Short_hash(rec.TxHash)
		rec.Profit = rec.RepReturned - rec.RepInvested
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_losing_rep_participants(market_aid int64) []p.RepLosingParticipant{

	records := make([]p.RepLosingParticipant,0,8)
	_,mkt_type,outcomes,_ := ss.Get_key_market_fields(market_aid)

	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM c.time_stamp)::BIGINT ts," +
				"TO_CHAR(c.time_stamp, 'dd/mm/yyyy HH:ii')," +
				"c.reporter_aid," +
				"ra.addr," +
				"m.winning_outcome," +
				"c.outcome_idx," +
				"c.amount_staked, " +
				"tx.tx_hash " +
			"FROM crowdsourcer_contrib c " +
				"JOIN transaction tx ON tx_id=tx.id " +
				"JOIN market m ON c.market_aid=m.market_aid " +
				"LEFT JOIN address ra ON c.reporter_aid=ra.address_id " +
				"LEFT JOIN crowdsourcer_redeemed r ON (" +
					"c.reporter_aid=r.reporter_aid AND " +
					"c.market_aid=r.market_aid AND " +
					"c.outcome_idx=r.outcome_idx" +
				") " +
			"WHERE c.market_aid=$1 AND r.id IS NULL AND m.status>3 " + // status: Fin or Fin invalid
			"ORDER BY c.time_stamp"

	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.RepLosingParticipant
		var winning_outc int
		err := rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.ReporterAid,
			&rec.ReporterAddr,
			&winning_outc,
			&rec.OutcomeIdx,
			&rec.RepLost,
			&rec.TxHash,
		)
		if (err!=nil) {
			if err != sql.ErrNoRows {
				ss.Log_msg(fmt.Sprintf("Error in Get_reporting_table(): %v : %v",err,query))
				os.Exit(1)
			}
			return records
		}
		if mkt_type == 2 {
		} else {
			if rec.OutcomeIdx == winning_outc {
				continue // winning report records are inserted into 'crowdsourcer_redeemed' table
			}
		}
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.OutcomeIdx,&outcomes)
		rec.TxHashSh = p.Short_hash(rec.TxHash)
		records = append(records,rec)
	}
	return records
}
