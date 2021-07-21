package dbs
import (
	"os"
	"fmt"
	"database/sql"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_polymarkets_processing_status() p.PolymarketProcStatus {

	var output p.PolymarketProcStatus
	var null_id sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM poly_proc_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO poly_proc_status DEFAULT VALUES"
				_,err := ss.db.Exec(query)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
				os.Exit(1)
			}
		} else {
			break
		}
	}
	if null_id.Valid {
		output.LastIdProcessed = null_id.Int64
	}
	return output
}
func (ss *SQLStorage) Update_polymarkets_process_status(status *p.PolymarketProcStatus) {

	var query string
	query = "UPDATE poly_proc_status SET last_evt_id = $1"

	_,err := ss.db.Exec(query,status.LastIdProcessed)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_condition_preparation(evt *p.Pol_ConditionPreparation) {


	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	oracle_aid:=ss.Lookup_or_create_address(evt.OracleAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_cond_prep (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"oracle_aid,condition_id,question_id,outcome_slot_count" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		oracle_aid,
		evt.ConditionId,
		evt.QuestionId,
		evt.OutcomeSlotCount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_cond_prep table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_condition_resolution(evt *p.Pol_ConditionResolution) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	oracle_aid :=ss.Lookup_or_create_address(evt.OracleAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_cond_res (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"oracle_aid,condition_id,question_id,outcome_slot_count,payout_numerators" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		oracle_aid,
		evt.ConditionId,
		evt.QuestionId,
		evt.OutcomeSlotCount,
		evt.PayoutNumerators,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_cond_res table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_position_split(evt *p.Pol_PositionSplit) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	stakeholder_aid:=ss.Lookup_or_create_address(evt.StakeHolderAddr,evt.BlockNum,evt.TxId)
	collateral_aid:=ss.Lookup_or_create_address(evt.CollateralToken,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_pos_split (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"stakeholder_aid,collateral_aid,parent_coll_id,condition_id,partition,amount" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10,$11"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		stakeholder_aid,
		collateral_aid,
		evt.ParentCollectionId,
		evt.ConditionId,
		evt.Partition,
		evt.Amount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_pos_split table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_position_merge(evt *p.Pol_PositionMerge) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	stakeholder_aid:=ss.Lookup_or_create_address(evt.StakeHolderAddr,evt.BlockNum,evt.TxId)
	collateral_aid:=ss.Lookup_or_create_address(evt.CollateralToken,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_pos_merge ("+
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"stakeholder_aid,collateral_aid,parent_coll_id,condition_id,partition,amount" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10,$11"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		stakeholder_aid,
		collateral_aid,
		evt.ParentCollectionId,
		evt.ConditionId,
		evt.Partition,
		evt.Amount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_pos_merge table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_payout_redemption(evt *p.Pol_PayoutRedemption) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	redeemer_aid:=ss.Lookup_or_create_address(evt.Redeemer,evt.BlockNum,evt.TxId)
	collateral_aid :=ss.Lookup_or_create_address(evt.CollateralToken,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_pay_redem (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"redeemer_aid,collateral_aid,parent_coll_id,index_sets,payout" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,"+
				"$6,$7,$8,$9,$10"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		redeemer_aid,
		collateral_aid,
		evt.ParentCollectionId,
		evt.ConditionId,
		evt.IndexSets,
		evt.Payout,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_cond_res table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_URI(evt *p.Pol_URI) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_uri ("+
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"uri_id,value" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.Id,
		evt.Value,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_uri table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_funding_added(evt *p.Pol_FundingAdded) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	funder_aid:=ss.Lookup_or_create_address(evt.Funder,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_fund_add (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"funder_aid,amounts_added,shares_minted" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,"+
				"$6,$7,$8"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		funder_aid,
		evt.AmountsAdded,
		evt.SharesMinted,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_fund_add table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_funding_removed(evt *p.Pol_FundingRemoved) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	funder_aid:=ss.Lookup_or_create_address(evt.Funder,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_fund_rem (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"funder_aid,amounts_removed,shares_burnt,collateral_removed" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,"+
				"$6,$7,$8"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		funder_aid,
		evt.AmountsRemoved,
		evt.SharesBurnt,
		evt.CollateralRemoved,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_fund_rem table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_fpmm_buy(evt *p.Pol_Buy) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	buyer_aid:=ss.Lookup_or_create_address(evt.Buyer,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_buy (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"buyer_aid,outcome_idx,investment_amount,fee_amount,tokens_bought" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,"+
				"$6,$7,$8,$9,$10"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		buyer_aid,
		evt.OutcomeIdx,
		evt.InvestmentAmount,
		evt.FeeAmount,
		evt.TokensBought,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_buy table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_fpmm_sell(evt *p.Pol_Sell) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	seller_aid:=ss.Lookup_or_create_address(evt.Seller,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_sell (" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"seller_aid,outcome_idx,return_amount,fee_amount,tokens_sold" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,"+
				"$6,$7,$8,$9,$10"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		seller_aid,
		evt.OutcomeIdx,
		evt.ReturnAmount,
		evt.TokensSold,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_fund_rem table: %v\n",err))
		os.Exit(1)
	}
}

