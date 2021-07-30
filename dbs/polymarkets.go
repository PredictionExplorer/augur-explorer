package dbs
import (
	"os"
	"fmt"
	"database/sql"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_polymarkets_processing_status() p.PolymarketProcStatus {

	var output p.PolymarketProcStatus
	var null_id,null_block sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id,last_block FROM pol_proc_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id,&null_block)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO pol_proc_status DEFAULT VALUES"
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
	if null_block.Valid {
		output.LastBlockNum = null_block.Int64
	}
	return output
}
func (ss *SQLStorage) Update_polymarkets_process_status(status *p.PolymarketProcStatus) {

	var query string
	query = "UPDATE pol_proc_status SET last_evt_id = $1,last_block=$2"

	_,err := ss.db.Exec(query,status.LastIdProcessed,status.LastBlockNum)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_condition_preparation(evt *p.Pol_ConditionPreparation) {

	evt_log_field := "evtlog_id,"
	evt_log_value := fmt.Sprintf("%v,",evt.EvtId)
	if evt.EvtId == 0 {
		evt_log_field = ""
		evt_log_value = ""
	}
	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	oracle_aid:=ss.Lookup_or_create_address(evt.OracleAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_cond_prep (" +
				evt_log_field+"block_num,tx_id,time_stamp,contract_aid, "+
				"oracle_aid,condition_id,question_id,outcome_slot_count" +
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8"+
			")"
	_,err := ss.db.Exec(query,
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

	evt_log_field := "evtlog_id,"
	evt_log_value := fmt.Sprintf("%v,",evt.EvtId)
	if evt.EvtId == 0 {
		evt_log_field = ""
		evt_log_value = ""
	}
	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	oracle_aid :=ss.Lookup_or_create_address(evt.OracleAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_cond_res (" +
				evt_log_field+"block_num,tx_id,time_stamp,contract_aid, "+
				"oracle_aid,condition_id,question_id,outcome_slot_count,payout_numerators" +
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9"+
			")"
	_,err := ss.db.Exec(query,
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

	evt_log_field := "evtlog_id,"
	evt_log_value := fmt.Sprintf("%v,",evt.EvtId)
	if evt.EvtId == 0 {
		evt_log_field = ""
		evt_log_value = ""
	}
	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	stakeholder_aid:=ss.Lookup_or_create_address(evt.StakeHolderAddr,evt.BlockNum,evt.TxId)
	collateral_aid:=ss.Lookup_or_create_address(evt.CollateralToken,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_pos_split (" +
				evt_log_field+"block_num,tx_id,time_stamp,contract_aid, "+
				"stakeholder_aid,collateral_aid,parent_coll_id,condition_id,partition,amount" +
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10"+
			")"
	_,err := ss.db.Exec(query,
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

	evt_log_field := "evtlog_id,"
	evt_log_value := fmt.Sprintf("%v,",evt.EvtId)
	if evt.EvtId == 0 {
		evt_log_field = ""
		evt_log_value = ""
	}
	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	stakeholder_aid:=ss.Lookup_or_create_address(evt.StakeHolderAddr,evt.BlockNum,evt.TxId)
	collateral_aid:=ss.Lookup_or_create_address(evt.CollateralToken,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_pos_merge ("+
				evt_log_field+"block_num,tx_id,time_stamp,contract_aid, "+
				"stakeholder_aid,collateral_aid,parent_coll_id,condition_id,partition,amount" +
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10"+
			")"
	_,err := ss.db.Exec(query,
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

	evt_log_field := "evtlog_id,"
	evt_log_value := fmt.Sprintf("%v,",evt.EvtId)
	if evt.EvtId == 0 {
		evt_log_field = ""
		evt_log_value = ""
	}
	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	redeemer_aid:=ss.Lookup_or_create_address(evt.Redeemer,evt.BlockNum,evt.TxId)
	collateral_aid :=ss.Lookup_or_create_address(evt.CollateralToken,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_pay_redem (" +
				evt_log_field + "block_num,tx_id,time_stamp,contract_aid, "+
				"redeemer_aid,collateral_aid,parent_coll_id,index_sets,payout" +
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,"+
				"$5,$6,$7,$8,$9"+
			")"
	_,err := ss.db.Exec(query,
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

	evt_log_field := "evtlog_id,"
	evt_log_value := fmt.Sprintf("%v,",evt.EvtId)
	if evt.EvtId == 0 {
		evt_log_field = ""
		evt_log_value = ""
	}
	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_uri ("+
				evt_log_field + "block_num,tx_id,time_stamp,contract_aid, "+
				"uri_id,value" +
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,$5,$6"+
			")"
	_,err := ss.db.Exec(query,
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

	evt_log_field := "evtlog_id,"
	evt_log_value := fmt.Sprintf("%v,",evt.EvtId)
	if evt.EvtId == 0 {
		evt_log_field = ""
		evt_log_value = ""
	}
	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	funder_aid:=ss.Lookup_or_create_address(evt.Funder,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_fund_addrem (" +
				evt_log_field+"block_num,tx_id,time_stamp,contract_aid, "+
				"funder_aid,op_type,amounts,sum_amounts,shares" +
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,"+
				"$5,$6,$7,$8,$9"+
			")"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		funder_aid,
		0,
		evt.AmountsAdded,
		evt.AllAmountsSummed,
		evt.SharesMinted,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert Liquidity Added event: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_funding_removed(evt *p.Pol_FundingRemoved) {

	evt_log_field := "evtlog_id,"
	evt_log_value := fmt.Sprintf("%v,",evt.EvtId)
	if evt.EvtId == 0 {
		evt_log_field = ""
		evt_log_value = ""
	}
	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	funder_aid:=ss.Lookup_or_create_address(evt.Funder,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_fund_addrem (" +
				evt_log_field+"block_num,tx_id,time_stamp,contract_aid, "+
				"funder_aid,op_type,amounts,sum_amounts,shares,collateral_removed" +
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,"+
				"$5,$6,$7,$8,$9,$10"+
			")"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		funder_aid,
		1,
		evt.AmountsRemoved,
		evt.AllAmountsSummed,
		evt.SharesBurnt,
		evt.CollateralRemoved,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert Liquidity Removed event: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_fpmm_buy(evt *p.Pol_Buy) {

	evt_log_field := "evtlog_id,"
	evt_log_value := fmt.Sprintf("%v,",evt.EvtId)
	if evt.EvtId == 0 {
		evt_log_field = ""
		evt_log_value = ""
	}
	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	buyer_aid:=ss.Lookup_or_create_address(evt.Buyer,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_buysell (" +
				evt_log_field+"block_num,tx_id,time_stamp,contract_aid, "+
				"user_aid,outcome_idx,collateral_amount,fee_amount,token_amount,op_type" +
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,"+
				"$5,$6,$7,$8,$9,$10"+
			")"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		buyer_aid,
		evt.OutcomeIdx,
		evt.InvestmentAmount,
		evt.FeeAmount,
		evt.TokensBought,
		0,	// BUY
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert BUY op into pol_buysell table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_fpmm_sell(evt *p.Pol_Sell) {

	evt_log_field := "evtlog_id,"
	evt_log_value := fmt.Sprintf("%v,",evt.EvtId)
	if evt.EvtId == 0 {
		evt_log_field = ""
		evt_log_value = ""
	}
	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	seller_aid:=ss.Lookup_or_create_address(evt.Seller,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO pol_buysell (" +
				evt_log_field+"block_num,tx_id,time_stamp,contract_aid, "+
				"user_aid,outcome_idx,collateral_amount,fee_amount,token_amount,op_type" +
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,"+
				"$5,$6,$7,$8,$9,$10"+
			")"
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		seller_aid,
		evt.OutcomeIdx,
		evt.ReturnAmount,
		evt.FeeAmount,
		evt.TokensSold,
		1,	// SELL
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert SELL op into pol_buysell table: %v\n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_fpmm_contract_aid(poly_market_id int64) int64 {

	var query string
	query = "SELECT mkt_mkr_aid FROM pol_market WHERE market_id=$1"

	var contract_aid int64
	res := ss.db.QueryRow(query,poly_market_id)
	err := res.Scan(&contract_aid)
	if err != nil {
		if err == sql.ErrNoRows {
			contract_aid = 0
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return contract_aid
}
func (ss *SQLStorage) Get_polymarkets_buysell_operations(contract_aid int64,offset,limit int) []p.API_Pol_BuySell_Op {

	records := make([]p.API_Pol_BuySell_Op,0,64)
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM bs.time_stamp)::BIGINT as ts," +
				"bs.time_stamp,"+
				"bs.block_num," +
				"bs.op_type," +
				"bs.outcome_idx," +
				"bs.collateral_amount,"+
				"bs.fee_amount,"+
				"bs.token_amount/1e+9,"+
				"bs.user_aid," +
				"ba.addr " +
			"FROM pol_buysell bs " +
				"JOIN address ba ON bs.user_aid=ba.address_id " +
				"wHERE bs.contract_aid = $1 "+
			"ORDER BY bs.time_stamp DESC "+
			"OFFSET $2 LIMIT $3"

	rows,err := ss.db.Query(query,contract_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_BuySell_Op
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.OperationType,
			&rec.OutcomeIdx,
			&rec.CollateralAmount,
			&rec.FeeAmount,
			&rec.TokenAmount,
			&rec.UserAid,
			&rec.UserAddr,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_polymarkets_liquidity_operations(contract_aid int64,offset,limit int) []p.API_Pol_Liquidity_Op {

	records := make([]p.API_Pol_Liquidity_Op,0,64)
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM liq.time_stamp)::BIGINT as ts," +
				"liq.time_stamp,"+
				"liq.block_num," +
				"liq.op_type," +
				"liq.shares,"+
				"liq.funder_aid, " +
				"la.addr " +
			"FROM pol_fund_addrem liq " +
				"JOIN address la ON liq.funder_aid=la.address_id " +
				"wHERE contract_aid = $1 "+
			"ORDER BY liq.time_stamp DESC "+
			"OFFSET $2 LIMIT $3"

	rows,err := ss.db.Query(query,contract_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_Liquidity_Op
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.OperationType,
			&rec.CollateralAmount,
			&rec.FunderAid,
			&rec.FunderAddr,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_erc1155_transfers(tx_id,contract_aid int64,signature string) []string {

	records := make([]string,0,4)
	var query string 
	query = "SELECT log_rlp FROM evt_log WHERE contract_aid=$1 AND tx_id=$2 AND signature=$3  ORDER BY id"

	rows,err := ss.db.Query(query,tx_id,contract_aid,signature)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rlp_encoded_log string
		err=rows.Scan(&rlp_encoded_log)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rlp_encoded_log)
	}
	return records
}
func (ss *SQLStorage) Get_buysell_operations(market_id int64,offset,limit int) []p.API_Pol_BuySell_Op {

	records := make([]p.API_Pol_BuySell_Op,0,64)
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM bs.time_stamp)::BIGINT as ts," +
				"bs.time_stamp,"+
				"bs.block_num," +
				"bs.op_type," +
				"bs.outcome_idx," +
				"bs.collateral_amount,"+
				"bs.fee_amount,"+
				"bs.token_amount,"+
				"bs.user_aid, " +
				"ba.addr " +
			"FROM pol_buysell bs " +
				"JOIN pol_market mkt ON p.contract_aid=mkt.mkt_mkr_aid " +
				"JOIN address ba ON bs.buyer_aid=ba.address_id " +
				"wHERE mkt.market_id = $1 "+
			"ORDER BY bs.time_stamp DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := ss.db.Query(query,market_id,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_BuySell_Op
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.OperationType,
			&rec.OutcomeIdx,
			&rec.CollateralAmount,
			&rec.FeeAmount,
			&rec.TokenAmount,
			&rec.UserAid,
			&rec.UserAddr,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_poly_market_stats(market_id int64) (p.API_Pol_MarketStats,error) {

	var output p.API_Pol_MarketStats
	var query string
	query = "SELECT " +
				"open_interest/1e+18," +
				"num_liquidity_ops," +
				"num_trades," +
				"total_volume/1e+18," +
				"total_fees/1e+18 " +
			"FROM pol_mkt_stats " +
			"WHERE market_id = $1"

	res := ss.db.QueryRow(query,market_id)
	err := res.Scan(
		&output.OpenInterest,
		&output.NumLiquidityOps,
		&output.NumTrades,
		&output.TotalVolume,
		&output.TotalFeesCollected,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return output,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return output,nil
}
func (ss *SQLStorage) Get_poly_market_info(market_id int64) (p.API_Pol_MarketInfo,error) {

	var output p.API_Pol_MarketInfo
	var query string
	query = "SELECT " +
				"question," +
				"condition_id," +
				"slug," +
				"resolution_source,"+
				"EXTRACT(EPOCH FROM created_at_ts)::BIGINT,"+
				"created_at_date," +
				"EXTRACT(EPOCH FROM end_date_ts)::BIGINT," +
				"end_date," +
				"category," +
				"fee/1e+18," +
				"market_type,"+
				"image," +
				"icon," +
				"description," +
				"outcomes,"+
				"ma.addr " +
			"FROM pol_market pm " +
				"JOIN address ma ON pm.mkt_mkr_aid=ma.address_id " +
			"WHERE pm.market_id=$1"

	res := ss.db.QueryRow(query,market_id)
	err := res.Scan(
		&output.Question,
		&output.ConditionId,
		&output.Slug,
		&output.ResolutionSource,
		&output.CreatedAtTs,
		&output.CreatedAtDate,
		&output.EndDateTs,
		&output.EndDate,
		&output.Category,
		&output.Fee,
		&output.MarketType,
		&output.Image,
		&output.Icon,
		&output.Description,
		&output.Outcomes,
		&output.MarketMakerAddr,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	output.MarketId = market_id
	return output,nil
}
func (ss *SQLStorage) Get_polymarket_global_liquidity_history(init_ts int,fin_ts int,interval int) []p.API_Pol_GlobalLiquidityHistoryEntry {

	var query string
	query = "WITH periods AS (" +
				"SELECT * FROM (" +
					"SELECT " +
						"generate_series AS start_ts,"+
						"TO_TIMESTAMP(EXTRACT(EPOCH FROM generate_series) + $3) AS end_ts "+
					"FROM (" +
						"SELECT * " +
							"FROM generate_series(" +
								"TO_TIMESTAMP($1)," +
								"TO_TIMESTAMP($2)," +
								"TO_TIMESTAMP($3)-TO_TIMESTAMP(0)) " +
					") AS i" +
				") AS data " +
			") " +
			"SELECT " +
				"COALESCE(COUNT(pr.id),0) as num_rows, " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				"SUM(sumamounts) AS sum_amounts," +
				"SUM(shares) AS sum_shares," +
				"SUM(collateral_removed) as collateral_removed "+
			"FROM periods AS p " +
				"LEFT JOIN pol_addrem_AS liq ON (" +
					"p.start_ts <= liq.time_stamp AND "+
					"liq.time_stamp < p.end_ts AND " +
			") " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"

	rows,err := ss.db.Query(query,init_ts,fin_ts,interval)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.API_Pol_GlobalLiquidityHistoryEntry,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_GlobalLiquidityHistoryEntry
		var sum_amounts,sum_shares,sum_collateral_removed sql.NullFloat64
		var num_rows int
		err=rows.Scan(
			&num_rows,
			&rec.StartTs,
			&sum_amounts,
			&sum_shares,
			&sum_collateral_removed,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		if sum_amounts.Valid {
			rec.SumAmounts = sum_amounts.Float64
		}
		if sum_shares.Valid {
			rec.SumShares = sum_shares.Float64
		}
		if sum_collateral_removed.Valid {
			rec.SumCollateralRemoved = sum_collateral_removed.Float64
		}
		records = append(records,rec)
	}
	return records

}
func (ss *SQLStorage) Get_market_liquidity_history() {

}
