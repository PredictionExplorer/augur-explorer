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
	market_type,mticks,_ := ss.get_market_type_and_ticks(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.PayoutNumerators)
	var query string
	query = "INSERT INTO crowdsourcer_created (" +
				"block_num,tx_id,time_stamp,market_aid,crowdsrc_aid," +
				"dispute_round,outcome_idx,payout_numerators,size" +
			") VALUES ($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		timestamp,
		market_aid,
		dispute_aid,
		evt.DisputeRound.Int64(),
		reported_outcome,
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
	market_type,mticks,_ := ss.get_market_type_and_ticks(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.PayoutNumerators)
	var query string
	query = "INSERT INTO crowdsourcer_redeemed (" +
				"block_num,tx_id,market_aid,reporter_aid,crowdsourcer_aid," +
				"time_stamp,outcome_idx,amount,rep,payout_numerators" +
			") VALUES ($1,$2,$3,$4,$5,TO_TIMESTAMP($6),$7,$8::DECIMAL/1e+18,$9::DECIMAL/1e+18,$10)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		market_aid,
		reporter_aid,
		crowdsourcer_aid,
		evt.Timestamp.Int64(),
		reported_outcome,
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
	market_type,mticks,_ := ss.get_market_type_and_ticks(market_aid)
	reported_outcome := get_outcome_idx_from_numerators(market_type,mticks,evt.PayoutNumerators)
	var query string
	query = "INSERT INTO crowdsourcer_completed (" +
				"block_num,tx_id,time_stamp,market_aid,crowdsrc_aid,next_win_start,next_win_end," +
				"dispute_round,outcome_idx,pacing_on,tot_rep_payout,tot_rep_market,payout_numerators" +
			") VALUES (" +
				"$1,$2,TO_TIMESTAMP($3),$4,$5,TO_TIMESTAMP($6),TO_TIMESTAMP($7),"+
				"$8,$9,$10,$11::DECIMAL/1e+18,$12::DECIMAL/1e+18,$13"+
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
				"cc.outcome_idx," +
				"cc.amount_staked," +
				"cc.current_stake," +
				"cc.stake_remaining " +
			"FROM crowdsourcer_contrib cc " +
				"JOIN transaction tx ON cc.tx_id=tx.id " +
				"LEFT JOIN address ra ON reporter_aid=ra.address_id " +
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
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TxHash,
			&rec.ReporterAid,
			&rec.ReporterAddr,
			&rec.DisputeRound,
			&rec.OutcomeIdx,
			&rec.AmountStaked,
			&rec.CurrentStake,
			&rec.StakeRemaining,
		)

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
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"time_stamp," +
				"tx.tx_hash," +
				"ira.address_id," +
				"ira.addr," +
				"ara.address_id," +
				"ara.addr," +
				"r.outcome_idx," +
				"r.is_designated," +
				"r.amount_staked AS amount_staked "+
			"FROM initial_report r " +
				"JOIN transaction tx ON r.tx_id=tx.id " +
				"LEFT JOIN address ira ON ini_reporter_aid=ira.address_id " +
				"LEFT JOIN address ara ON reporter_aid=ara.address_id "+
			"WHERE market_aid=$1"
	row := ss.db.QueryRow(query,market_aid)
	err=row.Scan(
		&rst.InitialReport.TimeStamp,
		&rst.InitialReport.DateTime,
		&rst.InitialReport.TxHash,
		&rst.InitialReport.InitialReporterAid,
		&rst.InitialReport.InitialReporterAddr,
		&rst.InitialReport.ActualReporterAid,
		&rst.InitialReport.ActualReporterAddr,
		&rst.InitialReport.OutcomeIdx,
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
	rst.InitialReport.TxHashSh = p.Short_hash(rst.InitialReport.TxHash)
	rst.InitialReport.OutcomeStr=get_outcome_str(uint8(mkt_type),rst.InitialReport.OutcomeIdx,&outcomes)

	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"cc.time_stamp," +
				"tx.tx_hash, " +
				"cc.crowdsrc_aid, " +
				"ca.addr, " +
				"cc.dispute_round, " +
				"cc.payout_numerators, " +
				"cc.size " +
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
			&rec.Size,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.TxHashSh = p.Short_hash(rec.TxHash)
		big_pnumerators := payout_numerators_text_to_big(payout_numerators)
		outcidx := get_outcome_idx_from_numerators(mkt_type,int64(num_ticks),big_pnumerators)
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),outcidx,&outcomes)
		rec.Contributions = ss.get_dispute_contributions(rec.CrowdsourcerAid,uint8(mkt_type),num_ticks,&outcomes)
		disputes = append(disputes,rec)
	}
	rst.Disputes = disputes

	return rst,nil
}
func (ss *SQLStorage) Get_round_table(market_aid int64) ([]p.RoundsRow,int,string) {

	_,mkt_type,outcomes,_ := ss.Get_key_market_fields(market_aid)
	fmt.Printf("market outcomes = %v\n",outcomes)
	outcomes = adjust_outcomes_str(mkt_type,outcomes)
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"cc.time_stamp," +
				"tot_rep_payout," +
				"tot_rep_market," +
				"dispute_round, " +
				"pacing_on, " +
				"outcome_idx " +
			"FROM crowdsourcer_completed cc " +
			"WHERE market_aid=$1 " +
			"ORDER BY time_stamp"
	fmt.Printf("query=%v\n",query)
	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	num_outcomes := 3	// Invalid,No,Yes
	if mkt_type == 1 { // categorical
		num_outcomes = len(strings.Split(outcomes,","))
	}
	if mkt_type == 2 { // scalar
		num_outcomes = 2
	}
	rounds := make([]p.RoundsRow,0,8)

	defer rows.Close()
	for rows.Next() {
		var rr p.RoundsRow
		var rec p.DisputeRound
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.RepPayout,
			&rec.MarketRep,
			&rec.RoundNum,
			&rec.PacingOn,
			&rec.OutcomeIdx,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		fmt.Printf("dump rec : %+v\n",rec)
		rec.OutcomeStr = get_outcome_str(uint8(mkt_type),rec.OutcomeIdx,&outcomes)
		rec.RoundNum--

		fmt.Printf("Timestamp=%v, Date=%v, outcomestr=%v, outcomeidx=%v\n",rec.TimeStamp,rec.DateTime,rec.OutcomeStr,rec.OutcomeIdx)
		//var outc_rounds p.OutcomeRounds
		rr.Rounds.RoundNum = rec.RoundNum
		rr.Rounds.ORounds = make([]p.DisputeRound,0,8)
		for i:=0; i<num_outcomes ; i++ {
			var empty_rec p.DisputeRound
			if i==rec.OutcomeIdx {
				rec.Color = true
				rr.Rounds.ORounds = append(rr.Rounds.ORounds,rec)
				fmt.Printf("Appending good record at i=%v, date=%v\n",i,rec.DateTime)
			} else {
				if len(rounds) > 0 {
					prev_rec := &rounds[len(rounds)-1].Rounds.ORounds[i]
					if prev_rec.TimeStamp != 0 {
						empty_rec = *prev_rec
						empty_rec.Color = false
						fmt.Printf("copying rec from previous: %+v\n",empty_rec)
					}
				}
				rr.Rounds.ORounds = append(rr.Rounds.ORounds,empty_rec)
				fmt.Printf("Appending empty record at i=%v\n",i)
			}
		}
		rounds = append(rounds,rr)
	}
	return rounds,num_outcomes,outcomes
}
