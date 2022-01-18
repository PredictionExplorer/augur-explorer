package dbs
import (
	"os"
	"fmt"
	"strings"
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
				"stakeholder_aid,collateral_aid,parent_coll_id,condition_id,partition,amount," +
				"tok_ids,tok_froms,tok_tos,tok_amounts"+
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,"+
				"$11,$12,$13,$14"+
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
		evt.TokenIds,
		evt.TokenFroms,
		evt.TokenTos,
		evt.TokenAmounts,
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
				"stakeholder_aid,collateral_aid,parent_coll_id,condition_id,partition,amount," +
				"tok_ids,tok_froms,tok_tos,tok_amounts"+
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,"+
				"$11,$12,$13,$14"+
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
		evt.TokenIds,
		evt.TokenFroms,
		evt.TokenTos,
		evt.TokenAmounts,
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
				"redeemer_aid,collateral_aid,parent_coll_id,condition_id,index_sets,payout," +
				"tok_ids,tok_froms,tok_tos,tok_amounts"+
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,"+
				"$5,$6,$7,$8,$9,$10,"+
				"$11,$12,$13,$14"+
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
		evt.TokenIds,
		evt.TokenFroms,
		evt.TokenTos,
		evt.TokenAmounts,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pol_pay_redem table: %v\n",err))
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
				"funder_aid,op_type,amounts,sum_amounts,shares,transfer_amount" +
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
		0,
		evt.AmountsAdded,
		evt.AllAmountsSummed,
		evt.SharesMinted,
		evt.ERC20Value,
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
				"funder_aid,op_type,amounts,sum_amounts,shares,collateral_removed,transfer_amount" +
			") VALUES (" +
				evt_log_value+"$1,$2,TO_TIMESTAMP($3),$4,"+
				"$5,$6,$7,$8,$9,$10,$11"+
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
		evt.ERC20Value,
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
func (ss *SQLStorage) Get_condition_id(poly_market_id int64) string {

	var query string
	query = "SELECT condition_id FROM pol_market WHERE market_id=$1"

	var condition_id string
	res := ss.db.QueryRow(query,poly_market_id)
	err := res.Scan(&condition_id)
	if err != nil {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	if len(condition_id)>2 {
		condition_id = condition_id[2:] // strip 0x
	}
	return condition_id
}
func (ss *SQLStorage) Get_polymarkets_buysell_operations(market_info *p.API_Pol_MarketInfo,contract_aid int64,offset,limit int) []p.API_Pol_BuySell_Op {

	records := make([]p.API_Pol_BuySell_Op,0,64)
	var query string
	query = "SELECT " +
				"bs.id," +
				"EXTRACT(EPOCH FROM bs.time_stamp)::BIGINT ts," +
				"TO_CHAR(bs.time_stamp,'DD-MM-YYYY HH::MM') date,"+
				"bs.block_num," +
				"bs.op_type," +
				"bs.outcome_idx," +
				"bs.collateral_amount/1e+6,"+
				"bs.fee_amount/1e+6,"+
				"bs.token_amount/1e+6,"+
				"bs.collateral_amount/COALESCE(NULLIF(bs.token_amount,0), 1) as price,"+
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

	outcomes := strings.Split(market_info.Outcomes,",")

	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_BuySell_Op
		err=rows.Scan(
			&rec.Id,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.OperationType,
			&rec.OutcomeIdx,
			&rec.CollateralAmount,
			&rec.FeeAmount,
			&rec.TokenAmount,
			&rec.Price,
			&rec.UserAid,
			&rec.UserAddr,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if rec.OutcomeIdx < len(outcomes) {
			rec.OutcomeStr = outcomes[rec.OutcomeIdx]
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
				"TO_CHAR(liq.time_stamp,'DD-MM-YYYY HH::MM') date,"+
				"liq.block_num," +
				"liq.tx_id,"+
				"liq.op_type," +
				"-liq.norm_collateral/1e+6 collateral,"+
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
			&rec.TxId,
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
func (ss *SQLStorage) Get_poly_market_stats(contract_aid int64) (p.API_Pol_MarketStats,error) {

	var output p.API_Pol_MarketStats
	var query string
	query = "SELECT " +
				"ABS(open_interest/1e+6) as open_interest," +
					// negative value for OI represents user deposits, so it will always be negative
				"num_liq_ops," +
				"num_trades," +
				"total_volume/1e+6," +
				"total_fees/1e+6 " +
			"FROM pol_mkt_stats " +
			"WHERE contract_aid = $1"

	res := ss.db.QueryRow(query,contract_aid)
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
func (ss *SQLStorage) Get_polymarket_global_liquidity_history(init_ts int,fin_ts int,interval int) []p.API_Pol_GlobalLiquidityHistoryEntry {

	var query string
	query = "SELECT sum(norm_collateral)/1e+6 AS accum_liq FROM pol_fund_addrem liq " +
				"JOIN pol_market pm ON liq.contract_aid=pm.mkt_mkr_aid " +
				"WHERE liq.time_stamp < TO_TIMESTAMP($1)"
	var initial_accum_collateral sql.NullFloat64
	err := ss.db.QueryRow(query,init_ts).Scan(&initial_accum_collateral)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Get_polymarket_global_liquidity_history(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}

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
				"COALESCE(COUNT(liq.id),0) as num_rows, " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				//"SUM(sumamounts) AS sum_amounts," +
				//"SUM(shares) AS sum_shares," +
				"SUM(norm_collateral)/1e+6 as collateral "+
			"FROM periods AS p " +
				"LEFT JOIN LATERAL ( "+
					"SELECT liq.id,liq.time_stamp,liq.norm_collateral "+
						"FROM  pol_fund_addrem liq "+
						"JOIN pol_market pm ON liq.contract_aid=pm.mkt_mkr_aid "+
				") liq ON " +
					"(p.start_ts <= liq.time_stamp) AND "+
					"(liq.time_stamp < p.end_ts) " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"

	rows,err := ss.db.Query(query,init_ts,fin_ts,interval)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.API_Pol_GlobalLiquidityHistoryEntry,0,8)
	var accum_liq float64 = 0.0
	if initial_accum_collateral.Valid {
		accum_liq = initial_accum_collateral.Float64
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_GlobalLiquidityHistoryEntry
		var /*sum_amounts,sum_shares,*/sum_collateral sql.NullFloat64
		var num_rows int
		err=rows.Scan(
			&num_rows,
			&rec.StartTs,
			//&sum_amounts,
			//&sum_shares,
			&sum_collateral,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		/*if sum_amounts.Valid {
			rec.SumAmounts = sum_amounts.Float64
		}*/
		/*if sum_shares.Valid {
			rec.SumShares = sum_shares.Float64
		}*/
		if sum_collateral.Valid {
			// revert the sign because negative liquidity is an expense for the user
			rec.Liquidity= -sum_collateral.Float64
			accum_liq = accum_liq + rec.Liquidity
		}
		rec.LiquidityAccum = accum_liq

		records = append(records,rec)
	}
	return records

}
func (ss *SQLStorage) Get_polymarket_market_liquidity_history(contract_aid int64,init_ts int,fin_ts int,interval int) []p.API_Pol_MarketLiquidityHistoryEntry {

	var query string
	query = "SELECT sum(norm_collateral/1e+6) AS accum_liq FROM pol_fund_addrem liq " +
				"JOIN pol_market pm ON liq.contract_aid=pm.mkt_mkr_aid " +
				"WHERE (pm.mkt_mkr_aid=$1) AND (liq.time_stamp < TO_TIMESTAMP($2))"
	var initial_accum_collateral sql.NullFloat64
	err := ss.db.QueryRow(query,contract_aid,init_ts).Scan(&initial_accum_collateral)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}

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
				"COALESCE(COUNT(liq.id),0) as num_rows, " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				//"SUM(sumamounts) AS sum_amounts," +
				//"SUM(shares) AS sum_shares," +
				"SUM(norm_collateral)/1e+6 as collateral "+
			"FROM periods AS p " +
				"LEFT JOIN LATERAL ( "+
					"SELECT liq.id,liq.time_stamp,liq.norm_collateral "+
						"FROM pol_fund_addrem liq "+
						"JOIN pol_market pm ON liq.contract_aid=pm.mkt_mkr_aid "+
						"WHERE (pm.mkt_mkr_aid=$4) " +
				") liq ON " +
					"(p.start_ts <= liq.time_stamp) AND "+
					"(liq.time_stamp < p.end_ts) " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"

	rows,err := ss.db.Query(query,init_ts,fin_ts,interval,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.API_Pol_MarketLiquidityHistoryEntry,0,8)
	var accum_liq float64 = 0.0
	if initial_accum_collateral.Valid {
		accum_liq = initial_accum_collateral.Float64
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_MarketLiquidityHistoryEntry
		var sum_collateral sql.NullFloat64
		var num_rows int64
		err=rows.Scan(
			&num_rows,
			&rec.StartTs,
			&sum_collateral,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		rec.NumOperations = num_rows
		if sum_collateral.Valid {
			rec.Liquidity= -sum_collateral.Float64
			accum_liq = accum_liq + rec.Liquidity
		}
		rec.LiquidityAccum = accum_liq

		records = append(records,rec)
	}
	return records

}
func (ss *SQLStorage) Update_polymarkets_unique_addresses(ts int64,num_addrs,num_funders,num_traders int64) {
	var query string
	query = "UPDATE pol_unique_addrs "+
				"SET "+
					"num_addrs = $2,"+
					"num_funders = $3,"+
					"num_traders = $4 "+ 
				"WHERE day=to_timestamp($1)"
	res,err:=ss.db.Exec(query,ts,num_addrs,num_funders,num_traders)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Update_polymarket_unique_addresses_entry() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf(
			"Error getting RowsAffected in Update_polymarket_unique_addresses_entry(): %v",err,
		))
		os.Exit(1)
	}
	if affected_rows == 0 {
		query = "INSERT INTO pol_unique_addrs(day,num_addrs,num_funders,num_traders) "+
					"VALUES(to_timestamp($1),$2,$3,$4)"
		_,err := ss.db.Exec(query,ts,num_addrs,num_funders,num_traders)
		if (err!=nil) {
			ss.Log_msg(
				fmt.Sprintf(
					"DB Error on INSERT in Update_polymarket_unique_addresses_entry(): %v q=%v",
					err,query,
				),
			);
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Calc_polymarkets_unique_addresses(ts_from int64,ts_to int64) (int64,int64,int64,bool) {

	var no_rows bool = true
	var query string
	query = "SELECT count(*) FROM ( " +
				"SELECT "+
						"DISTINCT bs.user_aid "+
					"FROM pol_buysell bs " +
						"JOIN pol_market pm ON pm.mkt_mkr_aid=bs.contract_aid " +
					"WHERE "+
						"bs.time_stamp >= TO_TIMESTAMP($1) AND "+
						"bs.time_stamp < TO_TIMESTAMP($2) "+
			") data"

	var num_traders sql.NullInt64
	err := ss.db.QueryRow(query,ts_from,ts_to).Scan(&num_traders)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Calc_polymarkets_unique_addresses(): %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		no_rows = false
	}

	query = "SELECT count(*) FROM ( " +
				"SELECT "+
						"DISTINCT liq.funder_aid AS user_aid "+
					"FROM pol_fund_addrem liq " +
						"JOIN pol_market pm ON pm.mkt_mkr_aid=liq.contract_aid " +
					"WHERE "+
						"liq.time_stamp >= TO_TIMESTAMP($1) AND "+
						"liq.time_stamp < TO_TIMESTAMP($2) "+
			") data"

	var num_funders sql.NullInt64
	err = ss.db.QueryRow(query,ts_from,ts_to).Scan(&num_funders)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Calc_polymarkets_unique_addresses(): %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		no_rows = false
	}

	query = "SELECT count(*) FROM ( " +
				"SELECT "+
						"DISTINCT user_aid "+
					"FROM ( "+
						"(SELECT " +
							"DISTINCT funder_aid AS user_aid "+
							"FROM pol_fund_addrem liq "+
								"JOIN pol_market pm ON pm.mkt_mkr_aid=liq.contract_aid " +
								"WHERE "+
									"liq.time_stamp >= TO_TIMESTAMP($1) AND "+
									"liq.time_stamp < TO_TIMESTAMP($2) "+
						") " +
						" UNION ALL " +
						"(SELECT " +
							"DISTINCT user_aid "+
							"FROM pol_buysell bs "+
								"JOIN pol_market pm ON pm.mkt_mkr_aid=bs.contract_aid " +
								"WHERE "+
									"bs.time_stamp >= TO_TIMESTAMP($1) AND "+
									"bs.time_stamp < TO_TIMESTAMP($2) "+
						") " +
					") data" +
			") result"

	var num_addrs sql.NullInt64
	err = ss.db.QueryRow(query,ts_from,ts_to).Scan(&num_addrs)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Calc_polymarkets_unique_addresses(): %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		no_rows = false
	}

	return num_addrs.Int64,num_funders.Int64,num_traders.Int64,no_rows
}
func (ss *SQLStorage) Get_polymarket_global_trading_history(init_ts int,fin_ts int,interval int) []p.API_Pol_GlobalTradingHistoryEntry {

	var query string
	query = "SELECT sum(normalized_amount)/1e+6 AS accum_trading FROM pol_buysell tr " +
				"JOIN pol_market pm ON tr.contract_aid=pm.mkt_mkr_aid " +
				"WHERE tr.time_stamp < TO_TIMESTAMP($1)"
	var initial_accum_collateral sql.NullFloat64
	err := ss.db.QueryRow(query,init_ts).Scan(&initial_accum_collateral)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}

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
				"COALESCE(COUNT(tr.id),0) as num_rows, " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				"SUM(normalized_amount)/1e+6 as collateral "+
			"FROM periods AS p " +
				"LEFT JOIN LATERAL ( "+
					"SELECT tr.id,tr.time_stamp,tr.normalized_amount "+
						"FROM  pol_buysell tr "+
						"JOIN pol_market pm ON tr.contract_aid=pm.mkt_mkr_aid "+
				") tr ON " +
					"(p.start_ts <= tr.time_stamp) AND "+
					"(tr.time_stamp < p.end_ts) " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"

	rows,err := ss.db.Query(query,init_ts,fin_ts,interval)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.API_Pol_GlobalTradingHistoryEntry,0,8)
	var accum_volume float64 = 0.0
	if initial_accum_collateral.Valid {
		accum_volume = initial_accum_collateral.Float64
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_GlobalTradingHistoryEntry
		var sum_collateral sql.NullFloat64
		var num_rows sql.NullInt64
		err=rows.Scan(
			&num_rows,
			&rec.StartTs,
			&sum_collateral,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		if num_rows.Valid {
			rec.NumOperations = num_rows.Int64
		}
		if sum_collateral.Valid {
			rec.TradingVol= -sum_collateral.Float64
			accum_volume = accum_volume + rec.TradingVol
		}
		rec.TradingVolAccum = accum_volume

		records = append(records,rec)
	}
	return records

}
func (ss *SQLStorage) Get_polymarket_market_trading_history(contract_aid int64,init_ts int,fin_ts int,interval int) []p.API_Pol_MarketTradingHistoryEntry {

	var query string
	query = "SELECT sum(normalized_amount)/1e+6 AS accum_vol FROM pol_buysell tr " +
				"JOIN pol_market pm ON tr.contract_aid=pm.mkt_mkr_aid " +
				"WHERE (pm.mkt_mkr_aid=$1) AND (tr.time_stamp < TO_TIMESTAMP($2))"
	var initial_accum_collateral sql.NullFloat64
	err := ss.db.QueryRow(query,contract_aid,init_ts).Scan(&initial_accum_collateral)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}

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
				"COALESCE(COUNT(tr.id),0) as num_rows, " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				"SUM(normalized_amount)/1e+6 as collateral "+
			"FROM periods AS p " +
				"LEFT JOIN LATERAL ( "+
					"SELECT tr.id,tr.time_stamp,tr.normalized_amount "+
						"FROM pol_buysell tr "+
						"JOIN pol_market pm ON tr.contract_aid=pm.mkt_mkr_aid "+
						"WHERE (pm.mkt_mkr_aid=$4) " +
				") tr ON " +
					"(p.start_ts <= tr.time_stamp) AND "+
					"(tr.time_stamp < p.end_ts) " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"

	rows,err := ss.db.Query(query,init_ts,fin_ts,interval,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.API_Pol_MarketTradingHistoryEntry,0,8)
	var accum_vol float64 = 0.0
	if initial_accum_collateral.Valid {
		accum_vol = initial_accum_collateral.Float64
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_MarketTradingHistoryEntry
		var sum_collateral sql.NullFloat64
		var num_rows sql.NullInt64
		err=rows.Scan(
			&num_rows,
			&rec.StartTs,
			&sum_collateral,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		if num_rows.Valid {
			rec.NumOperations = num_rows.Int64
		}
		if sum_collateral.Valid {
			rec.TradingVol = -sum_collateral.Float64
			accum_vol = accum_vol + rec.TradingVol
		}
		rec.TradingVolAccum = accum_vol

		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Polymarkets_data_feed(evtlog_id int64) (int64,[]p.API_Pol_DataFeed) {

	evtlog_id = 0 // for Development purposes (should be removed upon production release)
	var query string
	query = "SELECT " +
				"bs.evtlog_id,"+
				"EXTRACT(EPOCH FROM bs.time_stamp)::BIGINT as ts," +
				"bs.time_stamp,"+
				"bs.op_type," +
				"bs.outcome_idx," +
				"bs.collateral_amount/1e+6,"+
				"bs.fee_amount/1e+6,"+
				"bs.user_aid,"+
				"mkt.market_id," +
				"mkt.mkt_mkr_aid," +
				"ua.addr,"+
				"mkt.question, "+
				"mkt_mkr_addr " +
			"FROM pol_buysell bs " +
				"JOIN LATERAL ( "+ 
					"SELECT " +
						"pm.mkt_mkr_aid," +
						"ma.addr mkt_mkr_addr," +
						"pm.market_id,"+
						"pm.question " +
					"FROM pol_market pm " +
					"JOIN address ma ON pm.mkt_mkr_aid=ma.address_id " +
				") AS mkt ON bs.contract_aid=mkt.mkt_mkr_aid " +
				"JOIN address ua ON bs.user_aid=ua.address_id " +
				"WHERE bs.evtlog_id > $1 " +
			"ORDER BY bs.time_stamp DESC LIMIT 5"

	rows,err := ss.db.Query(query,evtlog_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.API_Pol_DataFeed,0,8)
	var data_evtlog_id int64 = 0
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_DataFeed
		err=rows.Scan(
			&rec.EvtlogId,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.OperationType,
			&rec.OutcomeIdx,
			&rec.Collateral,
			&rec.Fee,
			&rec.UserAid,
			&rec.MarketId,
			&rec.MarketMakerAid,
			&rec.UserAddr,
			&rec.MarketQuestion,
			&rec.MarketMakerAddr,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		if data_evtlog_id < rec.EvtlogId {
			data_evtlog_id = rec.EvtlogId
		}


		records = append(records,rec)
	}
	new_evtlog_id := evtlog_id
	if data_evtlog_id > new_evtlog_id {
		new_evtlog_id = data_evtlog_id
	}
	return new_evtlog_id,records

}
func (ss *SQLStorage) Update_data_feed_status(last_evt_id int64) {
	// sets event id to the latest id that was fed to the client
	var query string
	query = "UPDATE pol_data_feed " +
			"SET last_evt_id = $1 "

	result,err := ss.db.Exec(query,last_evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	rows_affected,err:=result.RowsAffected()
	if rows_affected == 0 {
		query = "INSERT INTO pol_data_feed(last_evt_id) VALUES($1)"
		_,err := ss.db.Exec(query,last_evt_id)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Get_data_feed_status() int64 {

	var query string
	query = "SELECT last_evt_id FROM pol_data_feed "
	var null_id sql.NullInt64
	res := ss.db.QueryRow(query)
	err := res.Scan(&null_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return null_id.Int64
}
func (ss *SQLStorage) Get_poly_market_trader_operations(market_info *p.API_Pol_MarketInfo,contract_aid,user_aid int64,offset,limit int) []p.API_Pol_TraderOp{

	records := make([]p.API_Pol_TraderOp,0,64)
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM bs.time_stamp)::BIGINT as ts," +
				"TO_CHAR(bs.time_stamp,'DD-MM-YYYY HH::MM') date,"+
				"bs.block_num," +
				"bs.op_type," +
				"bs.outcome_idx," +
				"bs.collateral_amount/1e+6,"+
				"bs.fee_amount/1e+6,"+
				"bs.token_amount/1e+6 "+
			"FROM pol_buysell bs " +
				"JOIN pol_market mkt ON bs.contract_aid=mkt.mkt_mkr_aid " +
				"JOIN address ba ON bs.user_aid=ba.address_id " +
				"wHERE (bs.contract_aid = $1) AND (user_aid=$2) "+
			"ORDER BY bs.id "+
			"OFFSET $3 LIMIT $4"
	rows,err := ss.db.Query(query,contract_aid,user_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	outcomes := strings.Split(market_info.Outcomes,",")
	var accum_pl float64 = 0.0
	var accum_collateral float64 = 0.0
	var profit_loss float64 = 0.0
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_TraderOp
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.OperationType,
			&rec.OutcomeIdx,
			&rec.CollateralAmount,
			&rec.FeeAmount,
			&rec.TokenAmount,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if rec.OutcomeIdx < len(outcomes) {
			rec.OutcomeStr = outcomes[rec.OutcomeIdx]
		}
		var prev_accum_collateral = accum_collateral
		if rec.OperationType == 0 { // buy
			accum_collateral = accum_collateral - rec.CollateralAmount
		} else {
			accum_collateral = accum_collateral + rec.CollateralAmount
		}
		profit_loss = accum_collateral - prev_accum_collateral
		accum_pl = accum_pl + profit_loss
		rec.ProfitLoss = profit_loss
		rec.AccumProfitLoss = accum_pl
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_poly_market_funder_operations(contract_aid,user_aid int64,offset,limit int) []p.API_Pol_Liquidity_Op {

	records := make([]p.API_Pol_Liquidity_Op,0,64)
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM liq.time_stamp)::BIGINT as ts," +
				"liq.time_stamp,"+
				"liq.block_num," +
				"liq.op_type," +
				"liq.norm_collateral/1e+6 "+
			"FROM pol_fund_addrem liq  " +
				"JOIN pol_market mkt ON liq.contract_aid=mkt.mkt_mkr_aid " +
				"JOIN address ba ON liq.funder_aid=ba.address_id " +
				"wHERE (liq.contract_aid = $1) AND (liq.funder_aid=$2) "+
			"ORDER BY liq.id "+
			"OFFSET $3 LIMIT $4"
	rows,err := ss.db.Query(query,contract_aid,user_aid,offset,limit)
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
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.CollateralAmount = - rec.CollateralAmount
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_polymarkets_market_user_list(contract_aid int64) []p.API_Pol_TraderListEntry {

	records := make([]p.API_Pol_TraderListEntry,0,512)
	var query string
	query = "SELECT " +
				"ums.user_aid,"+
				"ua.addr,"+
				"ums.tot_trades," +
				"ums.tot_liq_ops," +
				"ums.tot_volume/1e+6,"+
				"ums.tot_liq_given/1e+6,"+
				"ums.tot_fees/1e+6,"+
				"ums.profit/1e+6 "+
			"FROM pol_ustats_mkt ums " +
				"JOIN address ua ON ums.user_aid=ua.address_id "+
			"wHERE ums.contract_aid = $1 " +
			"ORDER BY ums.tot_trades DESC"

	rows,err := ss.db.Query(query,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_TraderListEntry
		err=rows.Scan(
			&rec.UserAid,
			&rec.UserAddr,
			&rec.NumTrades,
			&rec.NumLiquidityOps,
			&rec.TotalTradeVolume,
			&rec.TotalLiquidityVol,
			&rec.TotalFeesPaid,
			&rec.TotalProfitLoss,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.TotalLiquidityVol = -rec.TotalLiquidityVol

		records = append(records,rec)
	}
	return records

}
func (ss *SQLStorage) Get_gnosis_erc1155_transfer_events(tx_id int64,topping_evtlog_id int64) []p.EthereumEventLog {

	records := make([]p.EthereumEventLog,0,16)
	var query string
	query = "SELECT " +
				"id,contract_aid,topic0_sig,log_rlp " +
			"FROM evt_log " +
			"WHERE " +
				"(tx_id=$1) AND " +
				"(id < $2) " +
			"ORDER by id DESC"

	rows,err := ss.db.Query(query,tx_id,topping_evtlog_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.EthereumEventLog
		err=rows.Scan(
			&rec.EvtId,
			&rec.ContractAid,
			&rec.Topic0_Sig,
			&rec.RlpLog,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage) Get_poly_market_outcome_price_history(contract_aid int64,outcome_idx int32) []p.API_Pol_OutcomePriceHistoryEntry {

	var query string
	query = "SELECT " +
				"bs.id, "+
				"EXTRACT(EPOCH FROM bs.time_stamp)::BIGINT," +
				"op_type,"+
				"bs.collateral_amount/COALESCE(NULLIF(bs.token_amount,0), 1) as price "+
			"FROM pol_buysell bs "+
			"WHERE bs.contract_aid=$1 AND outcome_idx=$2 " +
			"ORDER BY bs.time_stamp"

	rows,err := ss.db.Query(query,contract_aid,outcome_idx)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	records := make([]p.API_Pol_OutcomePriceHistoryEntry,0,256)
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_OutcomePriceHistoryEntry
		err=rows.Scan(
			&rec.OperationId,
			&rec.TimeStamp,
			&rec.OperationType,
			&rec.Price,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Error in Get_poly_market_outcome_price_history(): %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_polymarket_top_profit_makers() []p.ProfitMaker {

	var query string
	query = "SELECT "+
				"a.addr,"+
				"r.top_profit,"+
				"r.profit/1e+6 " +
			"FROM pol_uranks AS r " +
				"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_profit ASC,r.profit DESC LIMIT 100"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.ProfitMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec p.ProfitMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.ProfitLoss)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_polymarket_top_trade_makers() []p.TradeMaker {

	var query string
	query = "SELECT " +
				"a.addr," +
				"r.top_trades,"+
				"r.total_trades " +
			"FROM pol_uranks AS r " +
				"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_trades ASC,r.total_trades DESC LIMIT 100"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.TradeMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec p.TradeMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.TotalTrades)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_polymarket_top_volume_makers() []p.VolumeMaker {

	var query string
	query = "SELECT "+
				"a.addr," +
				"r.top_volume,"+
				"r.volume/1e+6 "+
			"FROM pol_uranks AS r " +
				"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY r.top_volume ASC,r.volume DESC LIMIT 100"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.VolumeMaker,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec p.VolumeMaker
		err=rows.Scan(&rec.Addr,&rec.Percentage,&rec.Volume)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Update_polymarket_top_profit_rank(aid int64,value float64,profit float64) int64 {

	var query string
	query = "UPDATE pol_uranks SET top_profit = $2,profit=$3 WHERE aid = $1"
	res,err:=ss.db.Exec(query,aid,value,profit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Update_polymarket_top_profit_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_profit_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO pol_uranks(aid,top_profit,profit) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,aid,value,profit)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("Update_polymarket_top_profit_rank() failed: %v, q=%v",err,query))
		}
	}
	return affected_rows
}
func (ss *SQLStorage) Update_polymarket_top_total_trades_rank(aid int64,value float64,total_trades int64) int64 {

	var query string
	query = "UPDATE pol_uranks SET top_trades = $2,total_trades=$3 WHERE aid = $1"
	res,err:=ss.db.Exec(query,aid,value,total_trades)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Uppdate_polymarket_top_total_trades_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in update_top_total_trades_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO pol_uranks(aid,top_trades,total_trades) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,aid,value,total_trades)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("Update_polymarket_top_total_trades_rank() failed: value=%v, err: %v, q=%v",value,err,query))
			os.Exit(1)
		}

	}
	return affected_rows
}
func (ss *SQLStorage) Update_polymarket_top_volume_rank(aid int64,value float64,volume float64) int64 {

	var query string
	query = "UPDATE pol_uranks SET top_volume = $2,volume=$3 WHERE aid = $1"
	res,err:=ss.db.Exec(query,aid,value,volume)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Update_polymarket_top_volume_rank() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in Update_polymarket_top_volume_rank(): %v",err))
	}
	if affected_rows == 0 {
		query = "INSERT INTO pol_uranks(aid,top_volume,volume) VALUES($1,$2,$3)"
		_,err:=ss.db.Exec(query,aid,value,volume)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("Update_polymarket_top_volume_rank() failed: value=%v, err: %v, q=%v",value,err,query))
			os.Exit(1)
		}

	}
	return affected_rows
}
func (ss *SQLStorage) Get_polymarket_ranking_data_for_all_users() []p.RankStats {

	var query string
	query = "SELECT user_aid,tot_trades,profit,tot_volume FROM pol_ustats"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.RankStats,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec p.RankStats
		err=rows.Scan(&rec.Aid,&rec.TotalTrades,&rec.ProfitLoss,&rec.VolumeTraded)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_polymarket_contract_addresses() p.Pol_ContractAddresses {

	var output p.Pol_ContractAddresses
	var query string
	query = "SELECT "+
				"cond_tok_addr,usdc_addr," +
				"ct_a.address_id,usd_a.address_id "+
			"FROM pol_contracts ca " +
				"JOIN address ct_a ON ca.cond_tok_addr=ct_a.addr " +
				"JOIN address usd_a ON ca.usdc_addr=usd_a.addr "

	res := ss.db.QueryRow(query)
	err := res.Scan(
		&output.ConditionalToken,
		&output.USDC,
		&output.CondTokAid,
		&output.USDCAid,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output
		}
		ss.Log_msg(fmt.Sprintf("Get_polymarket_contract_addresses() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	return output
}
func (ss *SQLStorage) Search_polymarket_keywords(keywords string) []p.PolTextSearchResult {

	var query string
	query = "WITH search_tokens AS (" +
				"SELECT market_id,contract_aid,tok_type "+
				"FROM pol_mkt_words,plainto_tsquery($1) AS q " +
				"WHERE tokens @@ q" +
			") " +
			"SELECT " +
				"st.market_id," +
				"st.contract_aid,"+
				"st.tok_type, " +
				"ca.addr," +
				"pm.question," +
				"pm.description," +
				"stat.total_volume " +
			"FROM search_tokens AS st " +
			"LEFT JOIN pol_market AS pm ON st.market_id=pm.market_id " +
			"LEFT JOIN address AS ca ON pm.mkt_mkr_aid=ca.address_id " +
			"LEFT JOIN pol_mkt_stats stat ON pm.mkt_mkr_aid=stat.contract_aid " +
			"ORDER BY st.tok_type DESC,total_volume DESC"

	rows,err := ss.db.Query(query,keywords)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.PolTextSearchResult,0,8)

	defer rows.Close()
	for rows.Next() {
		var description,title sql.NullString
		var contract_addr sql.NullString
		var tot_vol sql.NullFloat64
		var rec p.PolTextSearchResult
		err=rows.Scan(
			&rec.MarketId,
			&rec.ContractAid,
			&rec.ObjType,
			&contract_addr,
			&title,
			&description,
			&tot_vol,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if contract_addr.Valid { rec.ContractAddr = contract_addr.String }
		if tot_vol.Valid { rec.Volume = tot_vol.Float64 }
		if title.Valid { rec.Title = title.String }
		if description.Valid { rec.Description = description.String }
		records = append(records,rec)
	}
	return records
}
