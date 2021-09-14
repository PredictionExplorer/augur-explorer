package dbs
import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"math/big"
	"database/sql"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_polymarkets_unique_users_stats(ts_day_from int ,ts_day_to int) []p.API_Pol_Unique_Users {

	records := make([]p.API_Pol_Unique_Users,0,32)
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM day)::BIGINT,"+
				"num_addrs,"+
				"num_funders,"+
				"num_traders "+
			"FROM pol_unique_addrs " +
			"WHERE (TO_TIMESTAMP($1) <= day) AND (day < TO_TIMESTAMP($2))" +
			"ORDER BY day"
	rows,err := ss.db.Query(query,ts_day_from,ts_day_to)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_Unique_Users
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.NumFunders,
			&rec.NumTraders,
			&rec.NumTotal,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Error in Get_polymarkets_unique_users_stats(): %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_market_liquidity_history() {

}
func (ss *SQLStorage) Get_poly_market_info(market_id int64) (p.API_Pol_MarketInfo,error) {

	var rec p.API_Pol_MarketInfo
	var query string
	query = "SELECT " +
				"question," +
				"pm.condition_id," +
				"slug," +
				"resolution_source,"+
				"EXTRACT(EPOCH FROM prep.time_stamp)::BIGINT AS created_at_ts,"+
				"prep.time_stamp, " +
				"EXTRACT(EPOCH FROM end_date_ts)::BIGINT AS ts_end," +
				"end_date," +
				"EXTRACT(EPOCH FROM res.time_stamp)::BIGINT as resolution_ts," +
				"res.time_stamp,"+
				"category," +
				"image," +
				"icon," +
				"description," +
				"tags," +
				"outcomes," +
				"market_type," +
				"market_type_code," +
				"mkt_mkr_aid," +
				"ma.addr AS mkt_mkr_addr, " +
				"mst.total_volume,"+
				"mst.open_interest," +
				"mst.total_liquidity,"+
				"mst.total_fees," +
				"mst.num_trades, " +
				"mst.num_liq_ops," +
				"res.id, " +
				"res.payout_numerators,"+
				"prep.question_id, "+
				"prep.outcome_slot_count, " +
				"prep.tx_hash " +
			"FROM pol_market pm " +
				"JOIN address ma ON pm.mkt_mkr_aid=ma.address_id " +
				"LEFT JOIN pol_mkt_stats mst ON pm.mkt_mkr_aid=mst.contract_aid " +
				"LEFT JOIN LATERAL ("+
					"SELECT prep.question_id,prep.outcome_slot_count,tx_hash,prep.condition_id,prep.time_stamp " +
					"FROM pol_cond_prep AS prep "+
						"JOIN transaction tx ON prep.tx_id=tx.id " +
				") AS prep ON pm.condition_id=CONCAT('0x',prep.condition_id) " +
				"LEFT JOIN pol_cond_res AS res ON pm.condition_id=CONCAT('0x',res.condition_id) " +

			"WHERE pm.market_id=$1"

	var n_created_ts,n_resolved_ts sql.NullInt64
	var n_created_date,n_resolved_date sql.NullString
	var n_volume,n_open_interest,n_liquidity,n_fees sql.NullFloat64
	var n_num_trades,n_num_liq_ops sql.NullInt64
	var n_resolution_id sql.NullInt64
	var n_question_id sql.NullString
	var n_outcome_slot_count sql.NullInt64
	var n_cond_prep_tx_hash sql.NullString
	var n_numerators sql.NullString
	res := ss.db.QueryRow(query,market_id)
	err := res.Scan(
			&rec.Question,
			&rec.ConditionId,
			&rec.Slug,
			&rec.ResolutionSource,
			&n_created_ts,
			&n_created_date,
			&rec.EndDateTs,
			&rec.EndDate,
			&n_resolved_ts,
			&n_resolved_date,
			&rec.Category,
			&rec.Image,
			&rec.Icon,
			&rec.Description,
			&rec.Tags,
			&rec.Outcomes,
			&rec.MarketType,
			&rec.MarketTypeCode,
			&rec.MarketMakerAid,
			&rec.MarketMakerAddr,
			&n_volume,
			&n_open_interest,
			&n_liquidity,
			&n_fees,
			&n_num_trades,
			&n_num_liq_ops,
			&n_resolution_id,
			&n_numerators,
			&n_question_id,
			&n_outcome_slot_count,
			&n_cond_prep_tx_hash,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return rec,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	rec.MarketId = market_id
	if n_created_ts.Valid { rec.CreatedAtTs = n_created_ts.Int64 }
	if n_created_date.Valid {rec.CreatedAtDate = n_created_date.String }
	if n_resolved_ts.Valid { rec.ResolvedTs = n_resolved_ts.Int64 }
	if n_resolved_date.Valid { rec.ResolvedDate = n_resolved_date.String }
	if n_volume.Valid { rec.Volume = n_volume.Float64 }
	if n_open_interest.Valid { rec.OpenInterest = n_open_interest.Float64 }
	if n_liquidity.Valid { rec.Liquidity = n_liquidity.Float64 }
	if n_fees.Valid { rec.TotalFeesCollected = n_fees.Float64 }
	if n_num_trades.Valid { rec.NumTrades = n_num_trades.Int64 }
	if n_num_liq_ops.Valid { rec.NumLiquidityOps = n_num_liq_ops.Int64 }
	if n_resolution_id.Valid {rec.WasResolved = true }
	if n_question_id.Valid {rec.QuestionId = n_question_id.String }
	if n_outcome_slot_count.Valid { rec.OutcomeSlotCount = n_outcome_slot_count.Int64 }
	if n_cond_prep_tx_hash.Valid { rec.CondPrepTxHash = n_cond_prep_tx_hash.String }
	if n_numerators.Valid {
		numerators_list := strings.Split(n_numerators.String,",")
		var fsum float64
		for n :=0 ; n < len(numerators_list) ; n ++ {
			if len(numerators_list[n])>0 {
				f,_:= strconv.ParseFloat(numerators_list[n], 64)
				fsum = fsum + f
			}
		}
		for n :=0 ; n < len(numerators_list) ; n ++ {
			if len(numerators_list[n])>0 {
				f,_:= strconv.ParseFloat(numerators_list[n], 64)
				rec.PayoutNumerators = append(rec.PayoutNumerators,f)
				if len(rec.PayoutNumeratorsStr) > 0 {
					rec.PayoutNumeratorsStr += ","
				}
				f = 100*f/fsum
				rec.PayoutNumeratorsStr += fmt.Sprintf("%.1f",f)+"%"
			}
		}
	}
	return rec,nil
}
func (ss *SQLStorage) Get_polymarkets_markets(status,sort int,category string) []p.API_Pol_MarketInfo {
	// status: 0 - all markets, 1 - not finalized, 2 - finalized
	// sort : 0 - by trading volume, 1 - by liquidity invested, 2-by creation date, 3-by resolution date, 4 - fees collected
	var where_condition string
	if status == 1 {
		where_condition = "WHERE (res.id IS NULL) "
	}
	if status == 2 {
		where_condition = "WHERE (res.id IS NOT NULL) "
	}
	if len(category) > 0 {
		if len(where_condition)==0 {
			where_condition += "WHERE "
		} else {
			where_condition += " AND "
		}
		where_condition += "(pm.category=$1) "
	}
	var sort_condition string = "ORDER BY mst.total_volume DESC NULLS LAST "
	if sort == 1 {
		sort_condition = "ORDER BY mst.open_interest ASC NULLS LAST" // ASC because we have OI negative
	}
	if sort == 2 {
		sort_condition = "ORDER BY prep.time_stamp DESC NULLS LAST"
	}
	if sort == 3 {
		sort_condition = "ORDER BY res.time_stamp DESC NULLS LAST"
	}
	if sort == 4 {
		sort_condition = "ORDER BY mst.total_fees DESC NULLS LAST"
	}

	records := make([]p.API_Pol_MarketInfo,0,32)
	var query string
	query = "SELECT " +
				"market_id," +
				"question," +
				"pm.condition_id," +
				"slug," +
				"resolution_source,"+
				"EXTRACT(EPOCH FROM prep.time_stamp)::BIGINT AS created_at_ts,"+
				"prep.time_stamp, " +
				"EXTRACT(EPOCH FROM end_date_ts)::BIGINT AS ts_end," +
				"end_date," +
				"EXTRACT(EPOCH FROM res.time_stamp)::BIGINT as resolution_ts," +
				"res.time_stamp,"+
				"category," +
				"image," +
				"icon," +
				"description," +
				"tags," +
				"outcomes," +
				"market_type," +
				"market_type_code,"+
				"mkt_mkr_aid," +
				"ma.addr AS mkt_mkr_addr," +
				"mst.total_volume/1e+6,"+
				"mst.open_interest/1e+6,"+
				"mst.total_liquidity,"+
				"mst.total_fees/1e+6," +
				"mst.num_trades, " +
				"mst.num_liq_ops, " +
				"res.id resolution_id, " +
				"prep.question_id, " +
				"prep.outcome_slot_count " +
			"FROM pol_market pm " +
				"JOIN address ma ON pm.mkt_mkr_aid=ma.address_id " +
				"LEFT JOIN pol_mkt_stats mst ON pm.mkt_mkr_aid=mst.contract_aid " +
				"LEFT JOIN pol_cond_prep AS prep ON pm.condition_id=CONCAT('0x',prep.condition_id) " +
				"LEFT JOIN pol_cond_res AS res ON pm.condition_id=CONCAT('0x',res.condition_id) " +
			where_condition +
			sort_condition

	fmt.Printf("query = %v\n",query)
	var err error
	var rows *sql.Rows
	if len(category) > 0 {
		rows,err = ss.db.Query(query,category)
	} else {
		rows,err = ss.db.Query(query)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var n_created_ts,n_resolved_ts sql.NullInt64
		var n_created_date,n_resolved_date sql.NullString
		var n_volume,n_open_interest,n_liquidity,n_fees sql.NullFloat64
		var n_num_trades,n_num_liq_ops sql.NullInt64
		var n_resolution_id sql.NullInt64
		var n_question_id sql.NullString
		var n_outcome_slot_count sql.NullInt64
		var rec p.API_Pol_MarketInfo
		err=rows.Scan(
			&rec.MarketId,
			&rec.Question,
			&rec.ConditionId,
			&rec.Slug,
			&rec.ResolutionSource,
			&n_created_ts,
			&n_created_date,
			&rec.EndDateTs,
			&rec.EndDate,
			&n_resolved_ts,
			&n_resolved_date,
			&rec.Category,
			&rec.Image,
			&rec.Icon,
			&rec.Description,
			&rec.Tags,
			&rec.Outcomes,
			&rec.MarketType,
			&rec.MarketTypeCode,
			&rec.MarketMakerAid,
			&rec.MarketMakerAddr,
			&n_volume,
			&n_open_interest,
			&n_liquidity,
			&n_fees,
			&n_num_trades,
			&n_num_liq_ops,
			&n_resolution_id,
			&n_question_id,
			&n_outcome_slot_count,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Error in Get_polymarkets_markets(): %v, q=%v",err,query))
			os.Exit(1)
		}
	//	fmt.Printf("n.created_ts.Int64 = %v (%v), created date=%v\n",n_created_ts.Int64,n_created_ts.Valid,n_created_date.String)
		if n_created_ts.Valid { rec.CreatedAtTs = n_created_ts.Int64 }
		if n_created_date.Valid {rec.CreatedAtDate = n_created_date.String }
		if n_resolved_ts.Valid { rec.ResolvedTs = n_resolved_ts.Int64 }
		if n_resolved_date.Valid { rec.ResolvedDate = n_resolved_date.String }
		if n_volume.Valid { rec.Volume = n_volume.Float64 }
		if n_open_interest.Valid { rec.OpenInterest = n_open_interest.Float64 }
		if n_liquidity.Valid { rec.Liquidity = -n_liquidity.Float64 }
		if n_fees.Valid { rec.TotalFeesCollected = n_fees.Float64 }
		if n_num_trades.Valid { rec.NumTrades = n_num_trades.Int64 }
		if n_num_liq_ops.Valid { rec.NumLiquidityOps = n_num_liq_ops.Int64 }
		if n_resolution_id.Valid {rec.WasResolved = true }
		if n_question_id.Valid { rec.QuestionId = n_question_id.String }
		if n_outcome_slot_count.Valid { rec.OutcomeSlotCount = n_outcome_slot_count.Int64 }
//		fmt.Printf("before append n.created_ts.Int64 = %v , created date=%v\n",rec.CreatedAtTs,rec.CreatedAtDate)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_poly_market_open_positions(contract_aid int64) ([]p.API_Pol_MarketOpenPosition,[]p.Pol_CondTokPrices) {

	records := make([]p.API_Pol_MarketOpenPosition,0,1024)
	var query string
	query = "SELECT "+
				"eh.cur_balance/1e+6," +
				"ptk.outcome_idx," +
				"ptk.token_id_hex," +
				"user_aid," +
				"ua.addr,"+
				"ms.tot_trades,"+
				"ms.tot_volume/1e+6,"+
				"ms.tot_liq_ops,"+
				"ms.tot_fees/1e+6,"+
				"ms.profit/1e+6 " +
			"FROM pol_tok_ids ptk " +
				"JOIN erc1155_tok et ON ptk.token_id_hex = et.token_id_hex "+
				"JOIN erc1155_holder eh ON et.token_id=eh.token_id " +
				"JOIN pol_ustats_mkt ms ON (eh.aid=ms.user_aid) AND (ms.contract_aid=ptk.contract_aid) " +
				"JOIN address ua ON eh.aid=ua.address_id "+
			"WHERE (ptk.contract_aid = $1) AND (eh.cur_balance>0) " +
			"ORDER BY ms.tot_volume DESC"

	rows,err := ss.db.Query(query,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_MarketOpenPosition
		err=rows.Scan(
			&rec.CurrentBalance,
			&rec.OutcomeIdx,
			&rec.TokenId,
			&rec.UserAid,
			&rec.UserAddr,
			&rec.NumTrades,
			&rec.TotalVolume,
			&rec.NumLiquidityOps,
			&rec.TotalFeesPaid,
			&rec.RealizedProfit,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Error in Get_polymarkets_markets(): %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	prices := ss.Calculate_prices(contract_aid)
	for i:=0; i<len(records); i++ {
		if int(records[i].OutcomeIdx) < len(prices) {
			records[i].CurrentPrice = prices[records[i].OutcomeIdx].TokenPrice
			pos_value := records[i].CurrentPrice * records[i].CurrentBalance
			profit := pos_value + records[i].TotalProfit // it is a + because User's deposits are negative
			records[i].UnrealizedProfit = profit
			records[i].TotalProfit = records[i].RealizedProfit + records[i].UnrealizedProfit
		}
	}
	return records,prices
}
func (ss *SQLStorage) Get_poly_market_user_open_positions(user_aid int64) ([]p.API_Pol_MarketUserOpenPosition) {

	records := make([]p.API_Pol_MarketUserOpenPosition,0,1024)
	var query string
	query = "SELECT "+
				"eh.cur_balance/1e+6," +
				"ptk.outcome_idx," +
				"ptk.token_id_hex," +
				"pm.market_id," +
				"pm.question," +
				"ms.tot_trades,"+
				"ms.tot_volume/1e+6,"+
				"ms.tot_liq_ops,"+
				"ms.tot_fees/1e+6,"+
				"ms.profit/1e+6 " +
			"FROM pol_tok_ids ptk " +
				"JOIN erc1155_tok et ON ptk.token_id_hex = et.token_id_hex "+
				"JOIN erc1155_holder eh ON et.token_id=eh.token_id " +
				"JOIN pol_ustats_mkt ms ON (eh.aid=ms.user_aid) AND (ms.contract_aid=ptk.contract_aid) " +
				"JOIN pol_market pm ON ptk.contract_aid = pm.mkt_mkr_aid " +
				"JOIN address ua ON eh.aid=ua.address_id "+
			"WHERE (ms.user_aid=$1) AND (eh.cur_balance>0) " +
			"ORDER BY ms.tot_volume DESC"
	fmt.Printf("q = %v \n",query)
	fmt.Printf("user_aid=%v\n",user_aid)
	rows,err := ss.db.Query(query,user_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_MarketUserOpenPosition
		err=rows.Scan(
			&rec.CollateralInvested,
			&rec.OutcomeIdx,
			&rec.TokenId,
			&rec.MarketId,
			&rec.MarketQuestion,
			&rec.NumTrades,
			&rec.TotalVolume,
			&rec.NumLiquidityOps,
			&rec.TotalFeesPaid,
			&rec.TotalProfit,
		)
		fmt.Printf("Adding market %v : %v\n",rec.MarketId,rec.MarketQuestion)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Error in Get_poly_market_user_open_positions (): %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Calculate_prices(contract_aid int64) []p.Pol_CondTokPrices {

	var query string
	query = "SELECT " +
				"ptk.outcome_idx,"+
				"et.token_id,"+
				"ptk.token_id_hex,"+
				"eh.cur_balance::TEXT " +
			"FROM pol_tok_ids ptk "+
				"JOIN erc1155_tok et ON ptk.token_id_hex=et.token_id_hex " +
				"JOIN erc1155_holder eh ON (et.token_id=eh.token_id AND eh.aid=ptk.contract_aid) " +
			"WHERE "+
				"ptk.contract_aid=$1"

	rows,err := ss.db.Query(query,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	records := make([]p.Pol_CondTokPrices,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec p.Pol_CondTokPrices
		err=rows.Scan(
			&rec.OutcomeIdx,
			&rec.TokenId,
			&rec.TokenIdHex,
			&rec.TokenBalanceStr,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Error in Calculate_prices(): %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	odds_weight_for_outcome_func :=  func(outc int,b []*big.Int) *big.Int {
		odds := big.NewInt(1)
		rec_len :=len(b)
		for j:=0; j< rec_len; j++ {
			if outc==j { continue; }
			odds.Mul(odds,b[j])
		}
		return odds
	}
	rec_len :=len(records)
	balances := make([]*big.Int,0,rec_len)
	for i:=0; i< rec_len; i++ {
		bal := big.NewInt(0)
		_,valid := bal.SetString(records[i].TokenBalanceStr,10)
		if !valid {
			ss.Log_msg(fmt.Sprintf("Error calculating gnosis token prices: bad decimal value for token %v: %v\n",records[i].TokenIdHex,records[i].TokenBalance))
			os.Exit(1)
		}
		balances = append(balances,bal)
	}
	prices := make([]*big.Float,0,rec_len)
	for i:=0; i<rec_len ; i++ {
		numerator := odds_weight_for_outcome_func(i,balances)
		denominator := big.NewInt(1)
		for j:=0; j<rec_len; j++ {
			odds_outcome := odds_weight_for_outcome_func(j,balances)
			denominator.Add(denominator,odds_outcome)
		}
		p := new(big.Float)
		numerator_float := new(big.Float)
		denominator_float := new(big.Float)
		numerator_float.SetInt(numerator)
		denominator_float.SetInt(denominator)
		p.Quo(numerator_float,denominator_float)
		prices = append(prices,p)
	}
	for i:=0; i<rec_len; i++ {
		f,_ := prices[i].Float64()
		records[i].TokenPrice = f
		f,_= strconv.ParseFloat(records[i].TokenBalanceStr,64)
		records[i].TokenBalance = f
	}
	return records

}
func (ss *SQLStorage) Get_poly_liquidity_provider_share_ratio(contract_aid int64) []p.API_Pol_LiquidityShareRatio {
	var query string
	query = "SELECT " +
				"eh.cur_balance, "+
				"eh.aid AS funder_aid," +
				"ua.addr," +
				"s.tot_liq_ops " +
			"FROM erc20_holder eh "+
				"JOIN pol_ustats_mkt s ON (eh.contract_aid=s.contract_aid AND eh.aid=s.user_aid) " +
				"JOIN address ua ON eh.aid=ua.address_id " +
			"WHERE eh.contract_aid=$1 " +
			"ORDER BY eh.cur_balance DESC"

	rows,err := ss.db.Query(query,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	records := make([]p.API_Pol_LiquidityShareRatio,0,128)
	defer rows.Close()
	var total_supply float64
	for rows.Next() {
		var rec p.API_Pol_LiquidityShareRatio
		err=rows.Scan(
			&rec.Balance,
			&rec.FunderAid,
			&rec.FunderAddr,
			&rec.TotalLiquidityOps,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Error in Get_liquidity_provider_share_ratio(): %v, q=%v",err,query))
			os.Exit(1)
		}
		total_supply = total_supply + rec.Balance
		records = append(records,rec)
	}
	for i:=0; i< len(records); i++ {
		records[i].ShareRatio = 100*records[i].Balance/total_supply
	}
	return records
}
func (ss *SQLStorage) Get_buysell_operation_info(id int64) (p.API_Pol_BuySell_Op,error) {

	var query string
	query = "SELECT " +
				"bs.id," +
				"EXTRACT(EPOCH FROM bs.time_stamp)::BIGINT as ts," +
				"bs.time_stamp,"+
				"bs.block_num," +
				"bs.op_type," +
				"bs.outcome_idx," +
				"bs.collateral_amount/1e+6,"+
				"bs.fee_amount/1e+6,"+
				"(bs.fee_amount/1e+6)*(bs.collateral_amount/COALESCE(NULLIF(bs.token_amount,0), 1)) fee_col," +
				"bs.token_amount/1e+6,"+
				"bs.collateral_amount/COALESCE(NULLIF(bs.token_amount,0), 1) as price,"+
				"bs.user_aid," +
				"ba.addr " +
			"FROM pol_buysell bs " +
				"JOIN address ba ON bs.user_aid=ba.address_id " +
				"wHERE bs.id = $1"

	res := ss.db.QueryRow(query,id)
	var rec p.API_Pol_BuySell_Op
	err := res.Scan(
			&rec.Id,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.OperationType,
			&rec.OutcomeIdx,
			&rec.CollateralAmount,
			&rec.FeeAmount,
			&rec.FeeInCollateral,
			&rec.TokenAmount,
			&rec.Price,
			&rec.UserAid,
			&rec.UserAddr,
	)
	if err!=nil {
		if err == sql.ErrNoRows {
			return rec,err
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	return rec,nil
}
func (ss *SQLStorage) Get_polymarket_user_ranks(sort int,order int) []p.UserRank {

	records := make([]p.UserRank,0,256)
	var query string
	var order_field string
	var order_dir string = "DESC"

	switch (sort) {
	case 0: order_field = "r.profit"
	case 1: order_field = "r.volume"
	case 2: order_field = "r.total_trades"
	default:
		return records
	}
	if order!=0 {
		order_dir="ASC"
	}

	query = "SELECT " +
				"r.aid,a.addr,r.profit,r.total_trades,r.volume " +
				"FROM pol_uranks AS r " +
					"JOIN  pol_ustats AS s ON r.aid=s.user_aid " +
			"LEFT JOIN address AS a ON r.aid = a.address_id " +
			"ORDER BY "+order_field+" "+order_dir

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.UserRank
		err=rows.Scan(
			&rec.Aid,
			&rec.Addr,
			&rec.ProfitLoss,
			&rec.TotalTrades,
			&rec.VolumeTraded,
		)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_polymarket_market_redemptions(condition_id string,offset,limit int) []p.API_Pol_MarketRedemption {

	records := make([]p.API_Pol_MarketRedemption,0,256)
	var query string
	query = "SELECT " +
				"r.id," +
				"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT as ts," +
				"r.time_stamp,"+
				"r.block_num," +
				"r.redeemer_aid," +
				"ra.addr," +
				"r.index_sets," +
				"r.payout/1e+6 " +
			"FROM pol_pay_redem r " +
				"LEFT JOIN address ra ON r.redeemer_aid = ra.address_id " +
			"WHERE r.condition_id=$1::TEXT " +
			"OFFSET $2 LIMIT $3"
	fmt.Printf("q= %v\n",query)
	fmt.Printf("condition_id= %v\n",condition_id)
	rows,err := ss.db.Query(query,condition_id,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_MarketRedemption
		err=rows.Scan(
			&rec.Id,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.RedeemerAid,
			&rec.RedeemerAddr,
			&rec.Outcomes,
			&rec.Payout,
		)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_polymarket_user_redemptions(user_aid int64) {


}
func (ss *SQLStorage) Get_polymarket_categories() []p.API_Pol_MarketCategory {

	records := make([]p.API_Pol_MarketCategory,0,128)
	var query string
	query = "SELECT " +
				"category,count(*) as total " +
			"FROM pol_market " +
			"GROUP BY category"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_MarketCategory
		err=rows.Scan(&rec.Category,&rec.NumMarkets)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_polymarket_erc1155_transfers(contract_aid int64,offset,limit int) []p.API_Pol_MarketERC1155Transfer {

	prices := ss.Calculate_prices(contract_aid)
	var token_ids string
	for i:=0; i<len(prices); i++ {
		if len(token_ids) >0 { token_ids = token_ids + "," }
		token_ids = token_ids + fmt.Sprintf("%v",prices[i].TokenId)
	}
	records := make([]p.API_Pol_MarketERC1155Transfer,0,512)
	var query string
	query = "SELECT " +
				"eb.id,"+
				"eb.token_id,"+
				"et.token_id_hex,"+
				"EXTRACT(EPOCH FROM eb.time_stamp)::BIGINT AS created_at_ts,"+
				"eb.time_stamp,"+
				"eb.parent_id,"+
				"eb.batch_id," +
				"eb.amount/1e+6,"+
				"eb.balance/1e+6,"+
				"eb.aid, "+
				"ea.addr," +
				"bs.user_aid,"+
				"bs.addr," +
				"bs.op_type," +
				"bs.amount, " +
				"f.funder_aid," +
				"f.addr, "+
				"f.op_type,"+
				"f.amount, " +
				"tx.tx_hash " +
			"FROM erc1155_bal eb " +
				"JOIN address ea ON eb.aid = ea.address_id " +
				"JOIN erc1155_tok et ON eb.token_id=et.token_id " +
				"LEFT JOIN LATERAL ("+
					"SELECT "+
						"bs.tx_id,bs.user_aid,bs.op_type,"+
						"bsa.addr,bs.collateral_amount amount " +
					"FROM pol_buysell bs "+
						"JOIN address bsa ON bs.user_aid=bsa.address_id "+
				") AS bs ON eb.tx_id=bs.tx_id " +
				"LEFT JOIN LATERAL (" +
					"SELECT f.tx_id,f.funder_aid,op_type,fa.addr,f.shares/1e+6 amount "+
					"FROM pol_fund_addrem f "+
						"JOIN address fa ON f.funder_aid=fa.address_id "+
				") AS f ON eb.tx_id=f.tx_id " +
				"JOIN transaction tx ON eb.tx_id=tx.id " +
			"WHERE eb.token_id IN("+token_ids+") " +
			"ORDER by eb.id "
	ss.Info.Printf("token_ids= %v\n",token_ids)
	ss.Info.Printf("query : %v\n",query)
	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_MarketERC1155Transfer
		var n_bs_addr,n_f_addr sql.NullString
		var n_batch_id,n_parent_id,n_bs_user_aid,n_f_funder_aid sql.NullInt64
		var n_buysell_op_type,n_fund_op_type sql.NullInt32
		var n_b_amount,n_f_amount sql.NullFloat64
		err=rows.Scan(
			&rec.BalOpId,
			&rec.TokenId,
			&rec.TokenIdHex,
			&rec.TimeStamp,
			&rec.DateTime,
			&n_parent_id,
			&n_batch_id,
			&rec.Amount,
			&rec.Balance,
			&rec.BalChgAid,
			&rec.BalChgAddr,
			&n_bs_user_aid,
			&n_bs_addr,
			&n_buysell_op_type,
			&n_b_amount,
			&n_f_funder_aid,
			&n_f_addr,
			&n_fund_op_type,
			&n_f_amount,
			&rec.TxHash,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if n_parent_id.Valid { rec.ParentId = n_parent_id.Int64 }
		if n_batch_id.Valid { rec.IsBatch = true; rec.BatchId=n_batch_id.Int64 }
		if n_bs_user_aid.Valid { rec.BuySellAid = n_bs_user_aid.Int64 }
		if n_bs_addr.Valid { rec.BuySellAddr = n_bs_addr.String }
		if n_buysell_op_type.Valid { rec.BuySellOpType = n_buysell_op_type.Int32 }
		if n_b_amount.Valid { rec.BuySellAmount = n_b_amount.Float64 }
		if n_f_funder_aid.Valid { rec.FunderAid = n_f_funder_aid.Int64 }
		if n_f_addr.Valid { rec.FunderAddr = n_f_addr.String }
		if n_fund_op_type.Valid { rec.FundOpType = n_fund_op_type.Int32 }
		if n_f_amount.Valid { rec.FunderAmount = n_f_amount.Float64 }

		records = append(records,rec)
	}
	ss.Info.Printf("rows returned = %v\n",len(records))
	return records
}
func (ss *SQLStorage) Get_polymarket_open_interst_history(usdc_aid,condtok_aid,contract_aid int64,condition_id string,offset,limit int) []p.API_Pol_OpenInterestHistory {

	records := make([]p.API_Pol_OpenInterestHistory,0,512)
	var query string
/*	query = "WITH usdc AS ("+
				"SELECT "+
					"EXTRACT(EPOCH FROM e20b.time_stamp)::BIGINT AS created_at_ts,"+
					"e20b.time_stamp datetime,"+
					"e20t.from_aid,"+
					"e20t.to_aid," +
					"e20b.id bal_id,"+
					"e20b.tx_id,"+
					"e20b.amount/1e+6 amount," +
					"balance/1e+6 as bal_usd "+
				"FROM erc20_bal e20b "+
					"JOIN erc20_transf e20t ON e20b.parent_id=e20t.id " +
				"WHERE "+
					"e20b.contract_aid=$1 AND "+
					"e20b.aid=$2" +
			") "+
			"SELECT "+
				"usdc.created_at_ts," +
				"usdc.datetime,"+
				"usdc.from_aid,"+
				"usdc.to_aid," +
				"usdc.tx_id,"+
				"tx.tx_hash,"+
				"fa.addr from_addr,"+
				"ta.addr, " +
				"usdc.bal_id,"+
				"bs.id,"+
				"bs.op_type,"+
				"f.id," +
				"f.op_type,"+
				"red.id,"+
				"usdc.amount,"+
				"usdc.bal_usd "+
			"FROM usdc "+
				"JOIN transaction tx ON usdc.tx_id=tx.id "+
				"JOIN address fa ON usdc.from_aid=fa.address_id "+
				"JOIN address ta ON usdc.to_aid=ta.address_id "+
				"LEFT JOIN pol_buysell bs ON usdc.tx_id=bs.tx_id "+
				"LEFT JOIN pol_fund_addrem f ON usdc.tx_id=f.tx_id "+
				"LEFT JOIN pol_pay_redem red ON usdc.tx_id=red.tx_id "+
			"ORDER by bal_id"
			$1 = usdc
			$2 = contract_aid
*/
/*
	query = "WITH b AS (" +
				"SELECT "+
					"tops.tx_id, "+
					"tops.parent_split_id,"+
					"tops.parent_merge_id,"+
					"tops.parent_redeem_id,"+
					"e20b.id bal_id,"+
					"EXTRACT(EPOCH FROM e20b.time_stamp)::BIGINT AS ts,"+
					"e20b.time_stamp datetime,"+
					"e20b.amount/1e+6 amount,"+
					"e20b.balance,"+
					"e20b.balance/1e+6 as bal_usd, "+
					"e20b.parent_id,"+
					"e20b.aid user_aid,"+
					"e20b.contract_aid, "+
					"e20t.from_aid,"+
					"e20t.to_aid "+
				"FROM pol_tok_id_ops tops "+
				"CROSS JOIN erc20_bal e20b "+
				"CROSS JOIN erc20_transf e20t "+
//				"JOIN erc20_bal e20b ON tops.tx_id=e20b.tx_id "+
//				"JOIN erc20_transf e20t ON e20t.id=e20b.parent_id "+
				"WHERE "+
					"tops.tx_id=e20b.tx_id AND "+
					"e20t.id=e20b.parent_id AND "+
					"tops.condition_id = $1 " +
//					"e20b.aid != $2 " +
//					"((e20b.contract_aid=$1 AND e20b.aid=$2)) " +
				"ORDER BY bal_id" +
			") " +
			"SELECT " +
				"DISTINCT b.bal_id bal_id,"+
				"b.ts," +
				"b.datetime,"+
				"b.from_aid,"+
				"b.to_aid," +
				"b.tx_id,"+
				"tx.tx_hash,"+
				"fa.addr from_addr,"+
				"ta.addr, " +
				"bs.id,"+
				"bs.op_type,"+
				"f.id," +
				"f.op_type,"+
				"red.id,"+
				"b.amount,"+
				"b.bal_usd "+
			"FROM b "+
				"JOIN transaction tx ON b.tx_id=tx.id "+
				"JOIN address fa ON b.from_aid=fa.address_id "+
				"JOIN address ta ON b.to_aid=ta.address_id "+
				"LEFT JOIN pol_buysell bs ON b.tx_id=bs.tx_id "+
				"LEFT JOIN pol_fund_addrem f ON b.tx_id=f.tx_id "+
				"LEFT JOIN pol_pay_redem red ON b.tx_id=red.tx_id " +
			"ORDER BY bal_id"
*/
	query = "WITH b AS (" +
				"SELECT "+
					"DISTINCT e20b.id bal_id, "+
					"e20b.parent_id, "+
					"e20b.tx_id, "+
					"e20b.contract_aid, "+
					"e20b.aid user_aid,"+
					"EXTRACT(EPOCH FROM e20b.time_stamp)::BIGINT AS ts,"+
					"e20b.time_stamp datetime,"+
					"e20b.amount,"+
					"e20b.balance,"+
					"e20b.balance/1e+6 as bal_usd, "+
					"tops.parent_split_id,"+
					"tops.parent_merge_id,"+
					"tops.parent_redeem_id "+
				"FROM pol_tok_id_ops tops "+
				"CROSS JOIN erc20_bal e20b "+
				"WHERE "+
					"tops.tx_id=e20b.tx_id AND "+
					"tops.condition_id = $1 " +
				"ORDER BY e20b.id" +
			") " +
			"SELECT " +
				"e20t.from_aid,"+
				"e20t.to_aid, "+
				"b.ts," +
				"b.datetime,"+
				"b.tx_id,"+
				"tx.tx_hash,"+
				"fa.addr from_addr,"+
				"ta.addr, " +
				"b.bal_id,"+
				"bs.id,"+
				"bs.op_type,"+
				"f.id," +
				"f.op_type,"+
				"red.id,"+
				"b.amount,"+
				"b.bal_usd "+
			"FROM b "+
				"JOIN transaction tx ON b.tx_id=tx.id "+
				"JOIN erc20_transf e20t ON b.parent_id=e20t.id "+
				"JOIN address fa ON e20t.from_aid=fa.address_id "+
				"JOIN address ta ON e20t.to_aid=ta.address_id "+
				"LEFT JOIN pol_buysell bs ON b.tx_id=bs.tx_id "+
				"LEFT JOIN pol_fund_addrem f ON b.tx_id=f.tx_id "+
				"LEFT JOIN pol_pay_redem red ON b.tx_id=red.tx_id "

	ss.Info.Printf("query : %v\n",query)
	rows,err := ss.db.Query(query,condition_id)
//	rows,err := ss.db.Query(query,usdc_aid,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	var prev_trf_to_aid int64 = -1
	var prev_trf_amount float64 = -1.0
	var fee_accum float64 = 0.0
	var open_interest float64 = 0.0
	for rows.Next() {
		var rec p.API_Pol_OpenInterestHistory
		var n_bs_id,n_far_id,n_red_id sql.NullInt64
		var n_bs_optype,n_far_optype sql.NullInt32
		err=rows.Scan(
			&rec.BalChgId,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.FromAid,
			&rec.ToAid,
			&rec.TxId,
			&rec.TxHash,
			&rec.FromAddr,
			&rec.ToAddr,
			&n_bs_id,
			&n_bs_optype,
			&n_far_id,
			&n_far_optype,
			&n_red_id,
			&rec.Amount,
			&rec.Balance,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if n_bs_id.Valid { rec.BuySellOpId = n_bs_id.Int64 } else { rec.BuySellOpId = -1 }
		if n_bs_optype.Valid { rec.BuySellOpType  = n_bs_optype.Int32 } else {rec.BuySellOpType = -1 }
		if n_far_id.Valid { rec.FundOpId = n_far_id.Int64 } else { rec.FundOpId = -1 }
		if n_far_optype.Valid { rec.FundOpType = n_far_optype.Int32 } else { rec.FundOpType = -1 }
		if n_red_id.Valid { rec.RedeemId = n_red_id.Int64} else { rec.RedeemId = -1 }
		/// filter for duplicates begin
		if (rec.ToAid == contract_aid) && (rec.Amount<0.0) { continue }
		if (rec.FromAid == contract_aid) && (rec.ToAid==condtok_aid) && (rec.Amount<0.0) { continue }
		if rec.FromAddr == "0x0000000000000000000000000000000000000000" { continue }
		/// filter for duplicate ends
		if n_bs_id.Valid {
			if (rec.ToAid == condtok_aid) && (rec.FromAid == contract_aid) && (rec.BuySellOpType == 0){
				// buy op
				if prev_trf_to_aid == contract_aid {
					rec.Fee = prev_trf_amount - (rec.Amount)
					fee_accum = fee_accum + rec.Fee
					open_interest = open_interest +  (rec.Amount)
				}
			}
			if (rec.FromAid == contract_aid) && (rec.FromAid != condtok_aid) && (rec.BuySellOpType == 1) {
				// sell op
				if prev_trf_to_aid == contract_aid {
					rec.Fee = prev_trf_amount - (rec.Amount)
					fee_accum = fee_accum + rec.Fee
					open_interest = open_interest - (rec.Amount) - rec.Fee
				}
			}
		}
		ss.Info.Printf("Amount %v: fundtype %v , rec.ToAid = %v , recFromAid = %v\n",
			rec.Amount,rec.FundOpType,rec.ToAid,rec.FromAid)
		if n_far_id.Valid {
			if rec.FundOpType == 0 { // add funds
				if (rec.ToAid == condtok_aid) && (rec.FromAid == contract_aid) {
					open_interest = open_interest + (rec.Amount)
				}
			}
			if rec.FundOpType == 1 { //withdraw funds
				if (rec.FromAid == contract_aid) && (rec.ToAid != condtok_aid) {
					open_interest = open_interest - (rec.Amount)
					ss.Info.Printf("new open interest is %v\n",open_interest)
				} else {
					ss.Info.Printf("condition to withdraw funds isn't met\n")
				}
			}
		}
		rec.FeeAccum = fee_accum
		rec.OpenInterest = open_interest
		prev_trf_to_aid = rec.ToAid
		prev_trf_amount = rec.Amount
		records = append(records,rec)
	}
	ss.Info.Printf("rows returned = %v\n",len(records))
	return records
}
