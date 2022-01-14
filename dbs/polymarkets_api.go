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
		err = rows.Scan(
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
				"mkt_group_question,"+
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
			&rec.MktGroupQuestion,
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
	if len(rec.MktGroupQuestion)>0 { rec.IsGroupMarket = true }
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
func (ss *SQLStorage) Get_poly_user_open_positions(user_aid int64) ([]p.API_Pol_MarketUserOpenPosition) {

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
/*
DISCONTINUED, to be deleted
func (ss *SQLStorage) Get_polymarket_open_interest_history(usdc_aid,condtok_aid,contract_aid int64,condition_id string,offset,limit int) []p.API_Pol_OpenInterestHistory {

	records := make([]p.API_Pol_OpenInterestHistory,0,512)
	var query string
	query = "WITH b AS (" +
				"SELECT "+
					"DISTINCT e20b.id bal_id, "+
					"e20b.parent_id, "+
					"e20b.tx_id, "+
					"e20b.contract_aid, "+
					"e20b.aid user_aid,"+
					"EXTRACT(EPOCH FROM e20b.time_stamp)::BIGINT AS ts,"+
					"TO_CHAR(e20b.time_stamp,'DD-MM-YYYY HH::MM') datetime,"+
					//"e20b.time_stamp datetime,"+
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
				"ORDER BY bal_id" +
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
				"b.amount/1e+6,"+
				"b.amount," +
				"b.bal_usd, "+
				"b.balance "+
			"FROM b "+
				"JOIN transaction tx ON b.tx_id=tx.id "+
				"JOIN erc20_transf e20t ON b.parent_id=e20t.id "+
				"JOIN address fa ON e20t.from_aid=fa.address_id "+
				"JOIN address ta ON e20t.to_aid=ta.address_id "+
				"LEFT JOIN pol_buysell bs ON b.tx_id=bs.tx_id "+
				"LEFT JOIN pol_fund_addrem f ON b.tx_id=f.tx_id "+
				"LEFT JOIN pol_pay_redem red ON b.tx_id=red.tx_id " +
			"ORDER BY bal_id "

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
			&rec.FromAid,
			&rec.ToAid,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TxId,
			&rec.TxHash,
			&rec.FromAddr,
			&rec.ToAddr,
			&rec.BalChgId,
			&n_bs_id,
			&n_bs_optype,
			&n_far_id,
			&n_far_optype,
			&n_red_id,
			&rec.Amount,
			&rec.IntegerAmount,
			&rec.Balance,
			&rec.IntegerBalance,
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
		if (rec.ToAid == contract_aid) && (rec.Amount<0.0) { 
			ss.Info.Printf("Skipping tx_id=%v (rule 1)\n",rec.TxId)
			continue
		}
		if (rec.FromAid == contract_aid) && (rec.ToAid==condtok_aid) && (rec.Amount<0.0) { 
			ss.Info.Printf("Skipping tx_id=%v (rule 2)\n",rec.TxId)
			continue
		}
		if rec.FromAddr == "0x0000000000000000000000000000000000000000" { 
			ss.Info.Printf("Skipping tx_id=%v (rule 3)\n",rec.TxId)
			continue
		}
		/// filter for duplicate ends
		if n_bs_id.Valid {
			if (rec.ToAid == condtok_aid) && (rec.FromAid == contract_aid) && (rec.BuySellOpType == 0){
				// buy op
				if prev_trf_to_aid == contract_aid {
					rec.IntegerFee = prev_trf_amount - (rec.IntegerAmount)
					rec.Fee = rec.IntegerFee/1000000.0
					fee_accum = fee_accum + rec.IntegerFee
					open_interest = open_interest +  (rec.IntegerAmount)
				}
			}
			if (rec.FromAid == contract_aid) && (rec.FromAid != condtok_aid) && (rec.BuySellOpType == 1) {
				// sell op
				if prev_trf_to_aid == contract_aid {
					rec.IntegerFee = prev_trf_amount - (-rec.IntegerAmount)
					rec.Fee = rec.IntegerFee / 1000000.0
					fee_accum = fee_accum + rec.IntegerFee
					open_interest = open_interest - (rec.IntegerAmount) - rec.IntegerFee
				}
			}
		}
		//ss.Info.Printf("n_far_id=%v contract_aid=%v\n",n_far_id.Int64,contract_aid)
		if n_far_id.Valid {
			if rec.FundOpType == 0 { // add funds
				if (rec.ToAid == condtok_aid) && (rec.FromAid == contract_aid) {
					open_interest = open_interest + (rec.IntegerAmount)
				}
			}
			if rec.FundOpType == 1 { //withdraw funds
				if (rec.FromAid == contract_aid) && (rec.ToAid != condtok_aid) {
					open_interest = open_interest - (rec.IntegerAmount)
				} else {
				}
			}
		}
		rec.AdjustedBalance = (rec.IntegerBalance - rec.IntegerAmount)/1000000.0
		rec.FeeAccum = fee_accum / 1000000.0
		rec.IntegerFeeAccum = fee_accum
		rec.OpenInterest = open_interest / 1000000.0
		prev_trf_to_aid = rec.ToAid
		prev_trf_amount = rec.IntegerAmount
//		ss.Info.Printf("Amount %v: fundtype %v , bs_id=%v, rec.ToAid = %v , recFromAid = %v, OI=%v\n",
//			rec.Amount,rec.FundOpType,rec.BuySellOpId,rec.ToAid,rec.FromAid,open_interest)
		records = append(records,rec)
	}
	ss.Info.Printf("rows returned = %v\n",len(records))
	return records
}
*/
func (ss *SQLStorage) Get_polymarket_open_interst_history_v2(usdc_aid,condtok_aid,contract_aid int64,condition_id string,offset,limit int) (p.API_Pol_OI_HistoryTotals,[]p.API_Pol_OpenInterestHistory) {
	// another version of history for testing

	var totals p.API_Pol_OI_HistoryTotals
	records := make([]p.API_Pol_OpenInterestHistory,0,512)
	var query string

	var resolution_evtlog_id int64
	var payout_numerators string
	var resolution_date string
	query = "SELECT evtlog_id,payout_numerators,time_stamp "+
			"FROM pol_cond_res WHERE condition_id=$1"
	res := ss.db.QueryRow(query,condition_id)
	var n_elog_id sql.NullInt64
	var n_resolution_date,n_numerators sql.NullString
	err := res.Scan(&n_elog_id,&n_numerators,&n_resolution_date)
	if err != nil {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		resolution_evtlog_id=n_elog_id.Int64
		payout_numerators=n_numerators.String
		resolution_date= n_resolution_date.String
	}

	query = "WITH b AS (" +
				"SELECT "+
					"e20b.id bal_id, "+
					"e20b.parent_id, "+
					"e20b.tx_id, "+
					"e20b.contract_aid, "+
					"e20b.aid user_aid,"+
					"EXTRACT(EPOCH FROM e20b.time_stamp)::BIGINT AS ts,"+
					"TO_CHAR(e20b.time_stamp,'DD-MM-YYYY HH::MM') datetime,"+
					"e20b.amount,"+
					"e20b.balance,"+
					"e20b.balance/1e+6 as bal_usd "+
				"FROM erc20_bal e20b "+
				"WHERE e20b.id IN( SELECT * FROM oi_history_transactions($1,$2,$3)) " +
			") " +
			"SELECT " +
				"e20t.evtlog_id,"+
				"e20t.from_aid,"+
				"e20t.to_aid, "+
				"b.user_aid," +
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
				"b.amount/1e+6,"+
				"b.amount," +
				"b.bal_usd, "+
				"b.balance "+
			"FROM b "+
				"JOIN transaction tx ON b.tx_id=tx.id "+
				"JOIN erc20_transf e20t ON b.parent_id=e20t.id "+
				"JOIN address fa ON e20t.from_aid=fa.address_id "+
				"JOIN address ta ON e20t.to_aid=ta.address_id "+
				"LEFT JOIN pol_buysell bs ON b.tx_id=bs.tx_id "+
				"LEFT JOIN pol_fund_addrem f ON b.tx_id=f.tx_id "+
				"LEFT JOIN pol_pay_redem red ON ((b.tx_id=red.tx_id) AND ((e20t.evtlog_id+1)=red.evtlog_id))" +
			"ORDER BY bal_id "

	ss.Info.Printf("usdc=%v, contract_aid=%v\n",usdc_aid,contract_aid)
	ss.Info.Printf("query : %v\n",query)
	rows,err := ss.db.Query(query,condition_id,usdc_aid,contract_aid)
//	rows,err := ss.db.Query(query,condition_id)
//	rows,err := ss.db.Query(query,usdc_aid,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	var separator_was_added bool
	var prev_trf_to_aid int64 = -1
	var prev_trf_from_aid int64 = -1
	var prev_trf_amount float64 = -1.0
	var fee_accum float64 = 0.0
	var open_interest float64 = 0.0
	for rows.Next() {
		var rec p.API_Pol_OpenInterestHistory
		var n_bs_id,n_far_id,n_red_id sql.NullInt64
		var n_bs_optype,n_far_optype sql.NullInt32
		var evtlog_id int64
		err=rows.Scan(
			&evtlog_id,
			&rec.FromAid,
			&rec.ToAid,
			&rec.UserAid,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TxId,
			&rec.TxHash,
			&rec.FromAddr,
			&rec.ToAddr,
			&rec.BalChgId,
			&n_bs_id,
			&n_bs_optype,
			&n_far_id,
			&n_far_optype,
			&n_red_id,
			&rec.Amount,
			&rec.IntegerAmount,
			&rec.Balance,
			&rec.IntegerBalance,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		ss.Info.Printf("tx_id=%v\n",rec.TxId)
		if (resolution_evtlog_id>0) && (evtlog_id>resolution_evtlog_id) {
			if !separator_was_added {
				var resolution_rec p.API_Pol_OpenInterestHistory
				resolution_rec.TxId = -1
				resolution_rec.PayoutNumerators = payout_numerators
				resolution_rec.DateTime = resolution_date
				records = append(records,resolution_rec)
				separator_was_added=true
				ss.Info.Printf("\t appending separator\n")
			}
		}
		if n_bs_id.Valid { rec.BuySellOpId = n_bs_id.Int64 } else { rec.BuySellOpId = -1 }
		if n_bs_optype.Valid { rec.BuySellOpType  = n_bs_optype.Int32 } else {rec.BuySellOpType = -1 }
		if n_far_id.Valid { rec.FundOpId = n_far_id.Int64 } else { rec.FundOpId = -1 }
		if n_far_optype.Valid { rec.FundOpType = n_far_optype.Int32 } else { rec.FundOpType = -1 }
		if n_red_id.Valid { rec.RedeemId = n_red_id.Int64} else { rec.RedeemId = -1 }
		/// filter for duplicates begin
		if (rec.ToAid == contract_aid) && (rec.Amount<0.0) { 
			ss.Info.Printf("Skipping tx_id=%v (rule 1)\n",rec.TxId)
			continue
		}
		if (rec.FromAid == contract_aid) && (rec.ToAid==condtok_aid) && (rec.Amount<0.0) { 
			ss.Info.Printf("Skipping tx_id=%v (rule 2)\n",rec.TxId)
			continue
		}
		if (rec.FromAid == contract_aid) && (rec.UserAid == contract_aid) && (rec.FundOpType!=1) {
			//prev_trf_to_aid = rec.ToAid
			//prev_trf_from_aid = rec.FromAid
			//prev_trf_amount = rec.IntegerAmount
			ss.Info.Printf("Skipping tx_id=%v (rule 3)\n",rec.TxId)
			continue
		}
		if rec.FromAddr == "0x0000000000000000000000000000000000000000" {
			ss.Info.Printf("Skipping tx_id=%v (rule 4)\n",rec.TxId)
			continue
		}
		if (rec.BuySellOpType == -1) && (rec.FundOpType == -1) && (rec.RedeemId == -1) {
			if rec.IntegerAmount < 0 {
				ss.Info.Printf("Skipping out-of Polymarket transaction to save fees txid=%v",rec.TxId)
				continue
			} else {
				// user is making transaction outside of FPMM contract (to avoid paying fees)
				open_interest = open_interest + rec.IntegerAmount
			}
			open_interest = open_interest + rec.IntegerAmount
		}
		/// filter for duplicate ends
		if n_bs_id.Valid {
			if (rec.ToAid == condtok_aid) && (rec.FromAid == contract_aid) && (rec.BuySellOpType == 0){
				// buy op
				if prev_trf_to_aid == contract_aid {
					rec.IntegerFee = prev_trf_amount - (rec.IntegerAmount)
					rec.Fee = rec.IntegerFee/1000000.0
					fee_accum = fee_accum + rec.IntegerFee
					open_interest = open_interest +  (rec.IntegerAmount)
				}
			}
			if (rec.FromAid == contract_aid) && (rec.FromAid != condtok_aid) && (rec.BuySellOpType == 1) {
				// sell op
				ss.Info.Printf("Entering sell op, bal_id=%v tx_id=%v prev_trf_from_aid=%v\n",rec.BalChgId,rec.TxId,prev_trf_from_aid)
				if prev_trf_from_aid == condtok_aid {
					ss.Info.Printf("Condition passed, calculating fees now prev_trf_amount=%v, IntegerAmount=%v\n",prev_trf_amount,rec.IntegerAmount)
					rec.IntegerFee = prev_trf_amount - rec.IntegerAmount
					rec.Fee = rec.IntegerFee / 1000000.0
					fee_accum = fee_accum + rec.IntegerFee
					open_interest = open_interest - (rec.IntegerAmount) 
				}
			}
		}
		//ss.Info.Printf("n_far_id=%v contract_aid=%v\n",n_far_id.Int64,contract_aid)
		if n_far_id.Valid {
			if rec.FundOpType == 0 { // add funds
				if (rec.ToAid == condtok_aid) && (rec.FromAid == contract_aid) {
					open_interest = open_interest + (rec.IntegerAmount)
				}
			}
			if rec.FundOpType == 1 { //withdraw funds
				if (rec.FromAid == contract_aid) && (rec.ToAid != condtok_aid) {
					open_interest = open_interest - (-rec.IntegerAmount)
				} else {
				}
			}
		}
		if n_red_id.Valid {// Payout redemption
			if (rec.UserAid != condtok_aid) && (rec.UserAid != contract_aid) {
				open_interest = open_interest - rec.IntegerAmount
			} else {
				ss.Info.Printf("Skipping unnecesary payout redemption record, tx_id=%v\n",rec.TxId)
				continue
			}
		}
		rec.AdjustedBalance = (rec.IntegerBalance - rec.IntegerAmount)/1000000.0
		rec.FeeAccum = fee_accum / 1000000.0
		rec.IntegerFeeAccum = fee_accum
		rec.OpenInterest = open_interest / 1000000.0
		rec.OIVerif = rec.OpenInterest + rec.FeeAccum
		totals.FinalOpenInterest = rec.OpenInterest
		totals.FinalFees = rec.FeeAccum
		prev_trf_to_aid = rec.ToAid
		prev_trf_from_aid = rec.FromAid
		prev_trf_amount = rec.IntegerAmount
//		ss.Info.Printf("Amount %v: fundtype %v , bs_id=%v, rec.ToAid = %v , recFromAid = %v, OI=%v\n",
//			rec.Amount,rec.FundOpType,rec.BuySellOpId,rec.ToAid,rec.FromAid,open_interest)
		records = append(records,rec)
	}
	ss.Info.Printf("rows returned = %v\n",len(records))
	return totals,records
}
func (ss *SQLStorage) Get_polymarket_open_interst_history_v3(usdc_aid,condtok_aid,contract_aid int64,condition_id string,offset,limit int) (p.API_Pol_OI_HistoryTotals,[]p.API_Pol_OpenInterestHistory) {
	// another version of history for testing

	var totals p.API_Pol_OI_HistoryTotals
	records := make([]p.API_Pol_OpenInterestHistory,0,512)
	red_buysell := make(map[int64]struct{},0)

	var query string

	var resolution_evtlog_id int64
	var payout_numerators string
	var resolution_date string
	query = "SELECT evtlog_id,payout_numerators,time_stamp "+
			"FROM pol_cond_res WHERE condition_id=$1"
	res := ss.db.QueryRow(query,condition_id)
	var n_elog_id sql.NullInt64
	var n_resolution_date,n_numerators sql.NullString
	err := res.Scan(&n_elog_id,&n_numerators,&n_resolution_date)
	if err != nil {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		resolution_evtlog_id=n_elog_id.Int64
		payout_numerators=n_numerators.String
		resolution_date= n_resolution_date.String
	}

	query = "WITH b AS (" +
				"SELECT "+
					"e20b.id bal_id, "+
					"e20b.parent_id, "+
					"e20b.tx_id, "+
					"e20b.contract_aid, "+
					"e20b.aid user_aid,"+
					"EXTRACT(EPOCH FROM e20b.time_stamp)::BIGINT AS ts,"+
					"TO_CHAR(e20b.time_stamp,'DD-MM-YYYY HH::MM') datetime,"+
					"e20b.amount,"+
					"e20b.balance,"+
					"e20b.balance/1e+6 as bal_usd "+
				"FROM erc20_bal e20b "+
				"WHERE e20b.id IN( SELECT * FROM oi_history_transactions($1,$2,$3)) " +
			") " +
			"SELECT " +
				"e20t.evtlog_id,"+
				"e20t.from_aid,"+
				"e20t.to_aid, "+
				"b.user_aid," +
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
				"red.id red_id,"+
				"b.amount/1e+6,"+
				"b.amount," +
				"b.bal_usd, "+
				"b.balance "+
			"FROM b "+
				"JOIN transaction tx ON b.tx_id=tx.id "+
				"JOIN erc20_transf e20t ON b.parent_id=e20t.id "+
				"JOIN address fa ON e20t.from_aid=fa.address_id "+
				"JOIN address ta ON e20t.to_aid=ta.address_id "+
			//	"LEFT JOIN pol_pos_split spl ON "+
				"LEFT JOIN pol_buysell bs ON b.tx_id=bs.tx_id "+
				"LEFT JOIN pol_fund_addrem f ON b.tx_id=f.tx_id "+
//				"LEFT JOIN pol_pay_redem red ON ((b.tx_id=red.tx_id) AND ((e20t.evtlog_id+1)=red.evtlog_id))" +
				"LEFT JOIN pol_pay_redem red ON (b.tx_id=red.tx_id) "+
			"ORDER BY bal_id,red_id "

	ss.Info.Printf("usdc=%v, contract_aid=%v\n",usdc_aid,contract_aid)
	ss.Info.Printf("query : %v\n",query)
	rows,err := ss.db.Query(query,condition_id,usdc_aid,contract_aid)
//	rows,err := ss.db.Query(query,condition_id)
//	rows,err := ss.db.Query(query,usdc_aid,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	var separator_was_added bool
	var prev_trf_to_aid int64 = -1
	var prev_trf_from_aid int64 = -1
	var prev_trf_amount float64 = -1.0
	var fee_accum float64 = 0.0
	var open_interest float64 = 0.0
	for rows.Next() {
		var rec p.API_Pol_OpenInterestHistory
		var n_bs_id,n_far_id,n_red_id sql.NullInt64
		var n_bs_optype,n_far_optype sql.NullInt32
		var evtlog_id int64
		err=rows.Scan(
			&evtlog_id,
			&rec.FromAid,
			&rec.ToAid,
			&rec.UserAid,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TxId,
			&rec.TxHash,
			&rec.FromAddr,
			&rec.ToAddr,
			&rec.BalChgId,
			&n_bs_id,
			&n_bs_optype,
			&n_far_id,
			&n_far_optype,
			&n_red_id,
			&rec.Amount,
			&rec.IntegerAmount,
			&rec.Balance,
			&rec.IntegerBalance,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		ss.Info.Printf("tx_id=%v\n",rec.TxId)
		if (resolution_evtlog_id>0) && (evtlog_id>resolution_evtlog_id) {
			if !separator_was_added {
				var resolution_rec p.API_Pol_OpenInterestHistory
				resolution_rec.TxId = -1
				resolution_rec.PayoutNumerators = payout_numerators
				resolution_rec.DateTime = resolution_date
				records = append(records,resolution_rec)
				separator_was_added=true
				ss.Info.Printf("\t appending separator\n")
			}
		}
		if n_bs_id.Valid { rec.BuySellOpId = n_bs_id.Int64 } else { rec.BuySellOpId = -1 }
		if n_bs_optype.Valid { rec.BuySellOpType  = n_bs_optype.Int32 } else {rec.BuySellOpType = -1 }
		if n_far_id.Valid { rec.FundOpId = n_far_id.Int64 } else { rec.FundOpId = -1 }
		if n_far_optype.Valid { rec.FundOpType = n_far_optype.Int32 } else { rec.FundOpType = -1 }
		if n_red_id.Valid { rec.RedeemId = n_red_id.Int64} else { rec.RedeemId = -1 }
		/// filter for duplicates begin
		if rec.RedeemId == -1 { // non Redeem type transactions
			if (rec.ToAid == contract_aid) && (rec.Amount<0.0) { 
				ss.Info.Printf("Skipping tx_id=%v (rule 1)\n",rec.TxId)
				continue
			}
			if (rec.FromAid == contract_aid) && (rec.ToAid==condtok_aid) && (rec.Amount<0.0) { 
				ss.Info.Printf("Skipping tx_id=%v (rule 2)\n",rec.TxId)
				continue
			}
			if (rec.FromAid == contract_aid) && (rec.UserAid == contract_aid) && (rec.FundOpType!=1) {
				//prev_trf_to_aid = rec.ToAid
				//prev_trf_from_aid = rec.FromAid
				//prev_trf_amount = rec.IntegerAmount
				ss.Info.Printf("Skipping tx_id=%v (rule 3)\n",rec.TxId)
				continue
			}
			if rec.FromAddr == "0x0000000000000000000000000000000000000000" {
				ss.Info.Printf("Skipping tx_id=%v (rule 4)\n",rec.TxId)
				continue
			}
			if (rec.BuySellOpType == -1) && (rec.FundOpType == -1) {
				if rec.IntegerAmount < 0 {
					ss.Info.Printf("Skipping out-of Polymarket transaction to save fees txid=%v",rec.TxId)
					continue
				} else {
					// user is making transaction outside of FPMM contract (to avoid paying fees)
					open_interest = open_interest + rec.IntegerAmount
				}
			}
			/// filter for duplicate ends
			if n_bs_id.Valid {
				if (rec.ToAid == condtok_aid) && (rec.FromAid == contract_aid) && (rec.BuySellOpType == 0){
					// buy op
					if prev_trf_to_aid == contract_aid {
						rec.IntegerFee = prev_trf_amount - (rec.IntegerAmount)
						rec.Fee = rec.IntegerFee/1000000.0
						fee_accum = fee_accum + rec.IntegerFee
						open_interest = open_interest +  (rec.IntegerAmount)
					}
				}
				if (rec.FromAid == contract_aid) && (rec.FromAid != condtok_aid) && (rec.BuySellOpType == 1) {
					// sell op
					ss.Info.Printf("Entering sell op, bal_id=%v tx_id=%v prev_trf_from_aid=%v\n",rec.BalChgId,rec.TxId,prev_trf_from_aid)
					if prev_trf_from_aid == condtok_aid {
						ss.Info.Printf("Condition passed, calculating fees now prev_trf_amount=%v, IntegerAmount=%v\n",prev_trf_amount,rec.IntegerAmount)
						rec.IntegerFee = prev_trf_amount - rec.IntegerAmount
						rec.Fee = rec.IntegerFee / 1000000.0
						fee_accum = fee_accum + rec.IntegerFee
						open_interest = open_interest - (rec.IntegerAmount) 
					}
				}
			}
			//ss.Info.Printf("n_far_id=%v contract_aid=%v\n",n_far_id.Int64,contract_aid)
			if n_far_id.Valid {
				if rec.FundOpType == 0 { // add funds
					if (rec.ToAid == condtok_aid) && (rec.FromAid == contract_aid) {
						open_interest = open_interest + (rec.IntegerAmount)
					}
				}
				if rec.FundOpType == 1 { //withdraw funds
					if (rec.FromAid == contract_aid) && (rec.ToAid != condtok_aid) {
						open_interest = open_interest - (-rec.IntegerAmount)
					} else {
					}
				}
			}
		} else { // rec.ReddemId > -1	// redeem type transactions
	/*		if (rec.ToAid== condtok_aid) && (rec.IntegerAmount<0) {
				ss.Info.Printf("Skipping duplicated transfer to condtok in redemption transfer (red_id=%v)\n",rec.RedeemId)
				continue
			}*/
			_,exists := red_buysell[rec.BuySellOpId]
			if !exists {
				red_buysell[rec.BuySellOpId]=struct{}{}
				open_interest = open_interest - rec.IntegerAmount
			} else {
				ss.Info.Printf("Skipping duplicated transfer to condtok (with buysell) in redemption transfer (red_id=%v)\n",rec.RedeemId)
				continue
			}
			/*if (rec.UserAid != condtok_aid) && (rec.UserAid != contract_aid) {
				open_interest = open_interest - rec.IntegerAmount
			} else {
				ss.Info.Printf("Skipping unnecesary payout redemption record, tx_id=%v\n",rec.TxId)
				continue
			}*/
		}
		rec.AdjustedBalance = (rec.IntegerBalance - rec.IntegerAmount)/1000000.0
		rec.FeeAccum = fee_accum / 1000000.0
		rec.IntegerFeeAccum = fee_accum
		rec.OpenInterest = open_interest / 1000000.0
		rec.OIVerif = rec.OpenInterest + rec.FeeAccum
		totals.FinalOpenInterest = rec.OpenInterest
		totals.FinalFees = rec.FeeAccum
		prev_trf_to_aid = rec.ToAid
		prev_trf_from_aid = rec.FromAid
		prev_trf_amount = rec.IntegerAmount
//		ss.Info.Printf("Amount %v: fundtype %v , bs_id=%v, rec.ToAid = %v , recFromAid = %v, OI=%v\n",
//			rec.Amount,rec.FundOpType,rec.BuySellOpId,rec.ToAid,rec.FromAid,open_interest)
		records = append(records,rec)
	}
	ss.Info.Printf("rows returned = %v\n",len(records))
	return totals,records
}
type TmpBuySellOp struct {
	EvtLogId		int64
	BuySellId		int64
	BuySellType		int32
	IntegerFee		float64
}
func (ss *SQLStorage)  build_buysell_operations(condition_id string,usdc_aid,contract_aid int64) map[int64]TmpBuySellOp {
	var query string
	/// build buysell operations
	// later this map is used to lookup for buysell operation 
	buysell_ops := make(map[int64]TmpBuySellOp)
	query = "SELECT "+
				"bs.evtlog_id,"+
				"bs.id,"+
				"bs.op_type, "+
				"bs.fee_amount " +
			"FROM pol_buysell bs " +
				"WHERE bs.tx_id IN( SELECT * FROM oi_history_transaction_ids($1,$2,$3)) "

	rows_buysell,err := ss.db.Query(query,condition_id,usdc_aid,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows_buysell.Close()
	for rows_buysell.Next() {
		var rec TmpBuySellOp
		err=rows_buysell.Scan(
			&rec.EvtLogId,
			&rec.BuySellId,
			&rec.BuySellType,
			&rec.IntegerFee,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		buysell_ops[rec.EvtLogId] = rec
	}
	return buysell_ops
}
type TmpFundOp struct {
	EvtLogId		int64
	FundOpId		int64
	FundType		int32
}
func (ss *SQLStorage)  build_fund_operations(condition_id string,usdc_aid,contract_aid int64) map[int64]TmpFundOp {
	var query string
	/// build buysell operations
	// later this map is used to lookup for buysell operation 
	fund_ops := make(map[int64]TmpFundOp)
	query = "SELECT "+
				"f.evtlog_id,"+
				"f.id," +
				"f.op_type "+
			"FROM pol_fund_addrem f " +
				"WHERE f.tx_id IN( SELECT * FROM oi_history_transaction_ids($1,$2,$3)) " 

	rows_fund_ops,err := ss.db.Query(query,condition_id,usdc_aid,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows_fund_ops.Close()
	for rows_fund_ops.Next() {
		var rec TmpFundOp
		err=rows_fund_ops.Scan(
			&rec.EvtLogId,
			&rec.FundOpId,
			&rec.FundType,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		fund_ops[rec.EvtLogId] = rec
	}
	return fund_ops
}
type TmpRedeemOp struct {
	EvtLogId		int64
	RedeemOpId		int64
	IntegerAmount	float64
	Amount			float64
}
func (ss *SQLStorage)  build_redeem_operations(condition_id string,usdc_aid,contract_aid int64) map[int64]TmpRedeemOp {
	var query string
	/// build buysell operations
	// later this map is used to lookup for buysell operation 
	redeem_ops := make(map[int64]TmpRedeemOp)
	query = "SELECT "+
				"red.evtlog_id,"+
				"red.id red_id,"+
				"red.payout,"+
				"red.payout/1e+6 "+
			"FROM pol_pay_redem red " +
				"WHERE red.tx_id IN( SELECT * FROM oi_history_transaction_ids($1,$2,$3)) "

	rows_redeem_ops,err := ss.db.Query(query,condition_id,usdc_aid,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows_redeem_ops.Close()
	for rows_redeem_ops.Next() {
		var rec TmpRedeemOp
		err=rows_redeem_ops.Scan(
			&rec.EvtLogId,
			&rec.RedeemOpId,
			&rec.IntegerAmount,
			&rec.Amount,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		redeem_ops[rec.EvtLogId] = rec
	}
	return redeem_ops
}
type TmpSplitOp struct {
	EvtLogId		int64
	SplitOpId		int64
	TxId			int64
	ReferenceId		int64
	ReferenceType	int32	//0-Undefined; 1-BuySell; 2-Fund Add/Remove
	AddressId		int64
	IntegerAmount	float64
	Amount			float64
}
func (ss *SQLStorage)  build_pos_split_operations(condition_id string,usdc_aid,contract_aid int64) map[int64]TmpSplitOp {
	var query string
	/// build buysell operations
	// later this map is used to lookup for buysell operation 
	split_ops := make(map[int64]TmpSplitOp)
	query = "SELECT "+
				"s.evtlog_id,"+
				"s.id split_id,"+
				"s.tx_id,"+
				"s.stakeholder_aid,"+
				"s.amount,"+
				"s.amount/1e+6 "+
			"FROM pol_pos_split s " +
				"WHERE s.tx_id IN( SELECT * FROM oi_history_transaction_ids($1,$2,$3)) "

	rows_split_ops,err := ss.db.Query(query,condition_id,usdc_aid,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows_split_ops.Close()
	for rows_split_ops.Next() {
		var rec TmpSplitOp
		err=rows_split_ops.Scan(
			&rec.EvtLogId,
			&rec.SplitOpId,
			&rec.TxId,
			&rec.AddressId,
			&rec.IntegerAmount,
			&rec.Amount,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		split_ops[rec.EvtLogId] = rec
	}
	return split_ops
}
type TmpMergeOp struct {
	EvtLogId		int64
	MergeOpId		int64
	TxId			int64
	ReferenceId		int64
	ReferenceType	int32	//0-Undefined; 1-BuySell; 2-Fund Add/Remove
	AddressId		int64
	IntegerAmount	float64
	Amount			float64
}
func (ss *SQLStorage)  build_pos_merge_operations(condition_id string,usdc_aid,contract_aid int64) map[int64]TmpMergeOp {
	var query string
	/// build buysell operations
	// later this map is used to lookup for buysell operation 
	merge_ops := make(map[int64]TmpMergeOp)
	query = "SELECT "+
				"m.evtlog_id,"+
				"m.id split_id,"+
				"m.tx_id,"+
				"m.stakeholder_aid,"+
				"m.amount,"+
				"m.amount/1e+6 "+
			"FROM pol_pos_merge m " +
				"WHERE m.tx_id IN( SELECT * FROM oi_history_transaction_ids($1,$2,$3)) "

	rows_merge_ops,err := ss.db.Query(query,condition_id,usdc_aid,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows_merge_ops.Close()
	for rows_merge_ops.Next() {
		var rec TmpMergeOp
		err=rows_merge_ops.Scan(
			&rec.EvtLogId,
			&rec.MergeOpId,
			&rec.TxId,
			&rec.AddressId,
			&rec.IntegerAmount,
			&rec.Amount,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		merge_ops[rec.EvtLogId] = rec
	}
	return merge_ops
}
func (ss *SQLStorage) Get_polymarket_open_interst_history_v4(usdc_aid,condtok_aid,contract_aid int64,condition_id string,offset,limit int) (p.API_Pol_OI_HistoryTotals,[]p.API_Pol_OpenInterestHistory) {
	// another version of history for testing

	var totals p.API_Pol_OI_HistoryTotals
	records := make([]p.API_Pol_OpenInterestHistory,0,512)

	var query string

	var resolution_evtlog_id int64
	var payout_numerators string
	var resolution_date string
	query = "SELECT evtlog_id,payout_numerators,time_stamp "+
			"FROM pol_cond_res WHERE condition_id=$1"
	res := ss.db.QueryRow(query,condition_id)
	var n_elog_id sql.NullInt64
	var n_resolution_date,n_numerators sql.NullString
	err := res.Scan(&n_elog_id,&n_numerators,&n_resolution_date)
	if err != nil {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		resolution_evtlog_id=n_elog_id.Int64
		payout_numerators=n_numerators.String
		resolution_date= n_resolution_date.String
	}

	buysell_ops := ss.build_buysell_operations(condition_id,usdc_aid,contract_aid)
	fund_ops := ss.build_fund_operations(condition_id,usdc_aid,contract_aid)
	redeem_ops := ss.build_redeem_operations(condition_id,usdc_aid,contract_aid)
	split_ops := ss.build_pos_split_operations(condition_id,usdc_aid,contract_aid)
	merge_ops := ss.build_pos_merge_operations(condition_id,usdc_aid,contract_aid)
	_=split_ops
	_=merge_ops

	query = "WITH b AS (" +
				"SELECT "+
					"e20b.id bal_id, "+
					"e20b.parent_id, "+
					"e20b.tx_id, "+
					"e20b.contract_aid, "+
					"e20b.aid user_aid,"+
					"EXTRACT(EPOCH FROM e20b.time_stamp)::BIGINT AS ts,"+
					"TO_CHAR(e20b.time_stamp,'DD-MM-YYYY HH::MM') datetime,"+
					"e20b.amount,"+
					"e20b.balance,"+
					"e20b.balance/1e+6 as bal_usd "+
				"FROM erc20_bal e20b "+
				"WHERE e20b.id IN( SELECT * FROM oi_history_transactions($1,$2,$3)) " +
			") " +
			"SELECT " +
				"e20t.evtlog_id,"+
				"e20t.from_aid,"+
				"e20t.to_aid, "+
				"b.user_aid," +
				"b.ts," +
				"b.datetime,"+
				"b.tx_id,"+
				"tx.tx_hash,"+
				"fa.addr from_addr,"+
				"ta.addr, " +
				"b.bal_id,"+
				"b.amount/1e+6,"+
				"b.amount," +
				"b.bal_usd, "+
				"b.balance "+
			"FROM b "+
				"JOIN transaction tx ON b.tx_id=tx.id "+
				"JOIN erc20_transf e20t ON b.parent_id=e20t.id "+
				"JOIN address fa ON e20t.from_aid=fa.address_id "+
				"JOIN address ta ON e20t.to_aid=ta.address_id "+
	//			"LEFT JOIN pol_buysell bs ON b.tx_id=bs.tx_id "+
	//			"LEFT JOIN pol_fund_addrem f ON b.tx_id=f.tx_id "+
//				"LEFT JOIN pol_pay_redem red ON ((b.tx_id=red.tx_id) AND ((e20t.evtlog_id+1)=red.evtlog_id))" +
	//			"LEFT JOIN pol_pay_redem red ON (b.tx_id=red.tx_id) "+
			"ORDER BY bal_id"

	ss.Info.Printf("usdc=%v, contract_aid=%v\n",usdc_aid,contract_aid)
	ss.Info.Printf("query : %v\n",query)
	rows,err := ss.db.Query(query,condition_id,usdc_aid,contract_aid)
//	rows,err := ss.db.Query(query,condition_id)
//	rows,err := ss.db.Query(query,usdc_aid,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	var separator_was_added bool
	var prev_trf_to_aid int64 = -1
	var prev_trf_from_aid int64 = -1
	var prev_trf_amount float64 = -1.0
	var fee_accum float64 = 0.0
	var open_interest float64 = 0.0
	for rows.Next() {
		var rec p.API_Pol_OpenInterestHistory
		var evtlog_id int64
		err=rows.Scan(
			&evtlog_id,
			&rec.FromAid,
			&rec.ToAid,
			&rec.UserAid,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TxId,
			&rec.TxHash,
			&rec.FromAddr,
			&rec.ToAddr,
			&rec.BalChgId,
			&rec.Amount,
			&rec.IntegerAmount,
			&rec.Balance,
			&rec.IntegerBalance,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		rec.EvtlogId=evtlog_id
		ss.Info.Printf("tx_id=%v\n",rec.TxId)
		if (resolution_evtlog_id>0) && (evtlog_id>resolution_evtlog_id) {
			if !separator_was_added {
				var resolution_rec p.API_Pol_OpenInterestHistory
				resolution_rec.TxId = -1
				resolution_rec.PayoutNumerators = payout_numerators
				resolution_rec.DateTime = resolution_date
				records = append(records,resolution_rec)
				separator_was_added=true
				ss.Info.Printf("\t appending separator\n")
			}
		}
		bs_rec,bs_exists := buysell_ops[evtlog_id]
		if !bs_exists {
			// note: all offsets were empirically found
			bs_rec,bs_exists = buysell_ops[evtlog_id+1]
			if !bs_exists {
				bs_rec,bs_exists = buysell_ops[evtlog_id+8]
				if !bs_exists {
					bs_rec,bs_exists = buysell_ops[evtlog_id+5]
					if !bs_exists {
						bs_rec,bs_exists = buysell_ops[evtlog_id+3]
						if !bs_exists {
							bs_rec,bs_exists = buysell_ops[evtlog_id+9]
						}
					}
				}
			}
		}
		if bs_exists {
			rec.BuySellOpId = bs_rec.BuySellId
			rec.BuySellOpType = bs_rec.BuySellType
		} else{
			rec.BuySellOpId = -1
		}
		//if n_bs_id.Valid { rec.BuySellOpId = n_bs_id.Int64 } else { rec.BuySellOpId = -1 }
		//if n_bs_optype.Valid { rec.BuySellOpType  = n_bs_optype.Int32 } else {rec.BuySellOpType = -1 }
		fund_rec,fund_exists := fund_ops[evtlog_id]
		if !fund_exists {
			fund_rec,fund_exists = fund_ops[evtlog_id+6]
			if !fund_exists {
				fund_rec,fund_exists = fund_ops[evtlog_id+9]
			}
		}
		if fund_exists {
			rec.FundOpId = fund_rec.FundOpId
			rec.FundOpType = fund_rec.FundType
		} else{
			rec.FundOpId = -1
		}
		//if n_far_id.Valid { rec.FundOpId = n_far_id.Int64 } else { rec.FundOpId = -1 }
		//if n_far_optype.Valid { rec.FundOpType = n_far_optype.Int32 } else { rec.FundOpType = -1 }

		redeem_rec,redeem_exists := redeem_ops[evtlog_id]
		if !redeem_exists {
			redeem_rec,redeem_exists = redeem_ops[evtlog_id+1]
		}
		if redeem_exists {
			rec.RedeemId = redeem_rec.RedeemOpId
			rec.RedeemIntegerPayout= redeem_rec.IntegerAmount
			rec.RedeemPayout = redeem_rec.Amount
		} else{
			rec.RedeemId = -1
		}
		//if n_red_id.Valid { rec.RedeemId = n_red_id.Int64} else { rec.RedeemId = -1 }
		/// filter for duplicates begin
		if rec.RedeemId == -1 { // non Redeem type transactions
			if (rec.ToAid == contract_aid) && (rec.Amount<0.0) { 
				ss.Info.Printf("Skipping tx_id=%v (rule 1)\n",rec.TxId)
				continue
			}
			if (rec.FromAid == contract_aid) && (rec.ToAid==condtok_aid) && (rec.Amount<0.0) { 
				ss.Info.Printf("Skipping tx_id=%v (rule 2)\n",rec.TxId)
				continue
			}
			if (rec.FromAid == contract_aid) && (rec.UserAid == contract_aid) && (rec.FundOpType!=1) {
				//prev_trf_to_aid = rec.ToAid
				//prev_trf_from_aid = rec.FromAid
				//prev_trf_amount = rec.IntegerAmount
				ss.Info.Printf("Skipping tx_id=%v (rule 3)\n",rec.TxId)
				continue
			}
			if rec.FromAddr == "0x0000000000000000000000000000000000000000" {
				ss.Info.Printf("Skipping tx_id=%v (rule 4)\n",rec.TxId)
				continue
			}
			if (rec.BuySellOpType == -1) && (rec.FundOpType == -1) {
				if rec.IntegerAmount < 0 {
					ss.Info.Printf("Skipping out-of Polymarket transaction to save fees txid=%v",rec.TxId)
					continue
				} else {
					// user is making transaction outside of FPMM contract (to avoid paying fees)
					open_interest = open_interest + rec.IntegerAmount
				}
			}
			/// filter for duplicate ends
			if rec.BuySellOpId != -1 {
				if (rec.ToAid == condtok_aid) && (rec.FromAid == contract_aid) && (rec.BuySellOpType == 0){
					// buy op
					if prev_trf_to_aid == contract_aid {
						rec.IntegerFee = prev_trf_amount - (rec.IntegerAmount)
						rec.Fee = rec.IntegerFee/1000000.0
						fee_accum = fee_accum + rec.IntegerFee
						open_interest = open_interest +  (rec.IntegerAmount)
					}
				}
				if (rec.FromAid == contract_aid) && (rec.FromAid != condtok_aid) && (rec.BuySellOpType == 1) {
					// sell op
					ss.Info.Printf("Entering sell op, bal_id=%v tx_id=%v prev_trf_from_aid=%v\n",rec.BalChgId,rec.TxId,prev_trf_from_aid)
					if prev_trf_from_aid == condtok_aid {
						ss.Info.Printf("Condition passed, calculating fees now prev_trf_amount=%v, IntegerAmount=%v\n",prev_trf_amount,rec.IntegerAmount)
						rec.IntegerFee = prev_trf_amount - rec.IntegerAmount
						rec.Fee = rec.IntegerFee / 1000000.0
						fee_accum = fee_accum + rec.IntegerFee
						open_interest = open_interest - (rec.IntegerAmount) 
					}
				}
			}
			//ss.Info.Printf("n_far_id=%v contract_aid=%v\n",n_far_id.Int64,contract_aid)
			if rec.FundOpId != -1 {
				if rec.FundOpType == 0 { // add funds
					if (rec.ToAid == condtok_aid) && (rec.FromAid == contract_aid) {
						open_interest = open_interest + (rec.IntegerAmount)
					}
				}
				if rec.FundOpType == 1 { //withdraw funds
					if (rec.FromAid == contract_aid) && (rec.ToAid != condtok_aid) {
						open_interest = open_interest - (-rec.IntegerAmount)
					} else {
					}
				}
			}
		} else { // rec.ReddemId > -1	// redeem type transactions
	/*		if (rec.ToAid== condtok_aid) && (rec.IntegerAmount<0) {
				ss.Info.Printf("Skipping duplicated transfer to condtok in redemption transfer (red_id=%v)\n",rec.RedeemId)
				continue
			}*/
	/*		_,exists := red_buysell[rec.BuySellOpId]
			if !exists {
				red_buysell[rec.BuySellOpId]=struct{}{}
				open_interest = open_interest - rec.IntegerAmount
			} else {
				ss.Info.Printf("Skipping duplicated transfer to condtok (with buysell) in redemption transfer (red_id=%v)\n",rec.RedeemId)
				continue
			}*/
			/*if (rec.UserAid != condtok_aid) && (rec.UserAid != contract_aid) {
				open_interest = open_interest - rec.IntegerAmount
			} else {
				ss.Info.Printf("Skipping unnecesary payout redemption record, tx_id=%v\n",rec.TxId)
				continue
			}*/
		}
		rec.AdjustedBalance = (rec.IntegerBalance - rec.IntegerAmount)/1000000.0
		rec.FeeAccum = fee_accum / 1000000.0
		rec.IntegerFeeAccum = fee_accum
		rec.OpenInterest = open_interest / 1000000.0
		rec.OIVerif = rec.OpenInterest + rec.FeeAccum
		totals.FinalOpenInterest = rec.OpenInterest
		totals.FinalFees = rec.FeeAccum
		prev_trf_to_aid = rec.ToAid
		prev_trf_from_aid = rec.FromAid
		prev_trf_amount = rec.IntegerAmount
//		ss.Info.Printf("Amount %v: fundtype %v , bs_id=%v, rec.ToAid = %v , recFromAid = %v, OI=%v\n",
//			rec.Amount,rec.FundOpType,rec.BuySellOpId,rec.ToAid,rec.FromAid,open_interest)
		records = append(records,rec)
	}
	ss.Info.Printf("rows returned = %v\n",len(records))
	return totals,records
}
func make_note(tx_id int64,split_id,merge_id int64,buysell_id int64,buysell_op_type int32,fund_id int64,fund_type int32,redeem_id int64,from_aid,to_aid,mkt_mkr_aid int64,condtok_aid int64,payout_numerators string) string {

	var output string

	if tx_id > 0 {
		if buysell_id > 0 {
			if buysell_op_type == 0 {
				output = "Buy. "
				if from_aid == mkt_mkr_aid {
					output = fmt.Sprintf("%v Market Maker contract is sending funds to conditional token contract",output)
				} else {
					output = fmt.Sprintf("%v User is sending funds to Market Maker contract",output)
				}
			} else {
				output = "Sell. "
				if to_aid == mkt_mkr_aid {
					output = fmt.Sprintf("%v Conditional Token contract is sending funds to Market Maker contract",output)
				} else {
					output = fmt.Sprintf("%v Market Maker contract is sending funds to the User",output)
				}
			}
		}
		if fund_id > 0 {
			if fund_type == 0 {
				output = "Add Funds. "
			} else {
				output = "Remove funds. "
			}
			if from_aid == condtok_aid {
				output = fmt.Sprintf("%v Conditional Token contract is sending funds to the User",output)
			}
			if to_aid == mkt_mkr_aid {
				output = fmt.Sprintf("%v User is sending funds to Market Maker contract",output)
			}
			if (to_aid != condtok_aid) && (to_aid != mkt_mkr_aid) {
				output = fmt.Sprintf("%v Market Maker is sending funds to the User",output)
			}
			if to_aid == condtok_aid {
				output = fmt.Sprintf("%v Market Maker is sending Funds to Conditional Token contract",output)
			}
		}
		if (buysell_id == 0) && (fund_id==0) && (redeem_id==0) {
			if split_id > 0 {
				output = "Split. "
			} else {
				output = "Merge. "
			}
			output = fmt.Sprintf("%v User is sending the operation to Conditional Tokens contract, buypassing Market Maker (probably to avoid paying fees)",output)
		}
		if redeem_id > 0 {
			output = fmt.Sprintf("Payout Redemption. Conditional Token contract is sending funds to the User")
		}
	} else {
		output = fmt.Sprintf("Payout numerators are: %v",payout_numerators)
	}
	return output
}
func (ss *SQLStorage) Get_polymarket_open_interst_history_v5(usdc_aid,condtok_aid,contract_aid int64,condition_id string,offset,limit int) (p.API_Pol_OI_HistoryTotals,[]p.API_Pol_OpenInterestHistory) {
	// another version of history for testing

	var totals p.API_Pol_OI_HistoryTotals
	records := make([]p.API_Pol_OpenInterestHistory,0,512)

	var query string

	var resolution_evtlog_id int64
	var payout_numerators string
	var resolution_date string
	query = "SELECT evtlog_id,payout_numerators,time_stamp "+
			"FROM pol_cond_res WHERE condition_id=$1"
	res := ss.db.QueryRow(query,condition_id)
	var n_elog_id sql.NullInt64
	var n_resolution_date,n_numerators sql.NullString
	err := res.Scan(&n_elog_id,&n_numerators,&n_resolution_date)
	if err != nil {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		resolution_evtlog_id=n_elog_id.Int64
		payout_numerators=n_numerators.String
		resolution_date= n_resolution_date.String
	}

	buysell_ops := ss.build_buysell_operations(condition_id,usdc_aid,contract_aid)
	fund_ops := ss.build_fund_operations(condition_id,usdc_aid,contract_aid)
	redeem_ops := ss.build_redeem_operations(condition_id,usdc_aid,contract_aid)
	split_ops := ss.build_pos_split_operations(condition_id,usdc_aid,contract_aid)
	merge_ops := ss.build_pos_merge_operations(condition_id,usdc_aid,contract_aid)
	ss.Info.Printf("buysell_ops = %v elts\nfund_ops = %v elts\nredeem_ops = %v elts\nsplit_ops = %v elts\nmerge_ops = %v elts\n",len(buysell_ops),len(fund_ops),len(redeem_ops),len(split_ops),len(merge_ops))

	query =
			"SELECT " +
				"e20b.id,"+
				"e.contract_aid,"+
				"e.id,"+
				"e20t.from_aid,"+
				"e20t.to_aid, "+
				"e20b.aid user_aid," +
				"EXTRACT(EPOCH FROM e20b.time_stamp)::BIGINT AS ts,"+
				"TO_CHAR(e20b.time_stamp,'DD-MM-YYYY HH::MM') datetime,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"e20b.amount/1e+6,"+
				"e20b.amount," +
				"e20b.balance/1e+6 as bal_usd,"+
				"e20b.balance, "+
				"bs.id bs_id, "+
				"bs.user_aid,"+
				"bs.op_type,"+
				"f.id fund_id," +
				"f.op_type,"+
				"f.funder_aid,"+
				"red.id red_id, "+
				"red.payout " +
			"FROM transaction tx "+
				"JOIN evt_log e ON tx.id=e.tx_id "+
				"LEFT JOIN erc20_transf e20t ON e.id=e20t.evtlog_id "+
				"LEFT JOIN erc20_bal e20b ON e20b.parent_id=e20t.id "+
		//		"JOIN address fa ON e20t.from_aid=fa.address_id "+
		//		"JOIN address ta ON e20t.to_aid=ta.address_id "+
				"LEFT JOIN pol_buysell bs ON e.id=bs.evtlog_id "+
				"LEFT JOIN pol_fund_addrem f ON e.id=f.evtlog_id "+
				"LEFT JOIN pol_pay_redem red ON e.id=red.evtlog_id " +
			"WHERE tx.id IN ("+
				"SELECT tx_id FROM oi_history_transaction_ids($1,$2,$3) data " +
			")"+
			"ORDER BY e.id"

	ss.Info.Printf("usdc=%v, contract_aid=%v\n",usdc_aid,contract_aid)
	d_query := strings.ReplaceAll(query,"$1",fmt.Sprintf("'%v'",condition_id))
	d_query = strings.ReplaceAll(d_query,"$2",fmt.Sprintf("%d",usdc_aid))
	d_query = strings.ReplaceAll(d_query,"$3",fmt.Sprintf("%d",contract_aid))
	ss.Info.Printf("query : %v\n",d_query)
	rows,err := ss.db.Query(query,condition_id,usdc_aid,contract_aid)
//	rows,err := ss.db.Query(query,condition_id)
//	rows,err := ss.db.Query(query,usdc_aid,contract_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	var separator_was_added bool
	var fee_accum float64 = 0.0
	var open_interest float64 = 0.0
	type Boundary struct {
		EvtLogId		int64
		OpType			int32 //0-undefined,1-PosSplit,2-PosMerge,3-BuySell,4-Fund,5-Redeem,6-Resolved
		Offset			int32 // offset to the main array of records
		Len				int32 // length (how many records follow after the Offset)
	}
	boundaries := make([]Boundary,0,256)
	counter := int32(0)
	addresses := make(map[int64]string)
	var b Boundary
	var last_split_op TmpSplitOp
	var last_merge_op TmpMergeOp
	var boundary_tx int64 = 0

	for rows.Next() {
		var rec p.API_Pol_OpenInterestHistory
		//var evtlog_id int64
		var n_evtlog_id,n_bal_id,n_from_aid,n_to_aid,n_user_aid,n_timestamp sql.NullInt64
		var n_tx_id sql.NullInt64
		var n_datetime,n_txhash sql.NullString
		var n_amount,n_int_amount,n_bal_usd,n_balance,n_red_payout sql.NullFloat64
		var n_bs_id,n_fund_id,n_red_id,n_bs_user_aid,n_funder_aid sql.NullInt64
		var n_bs_op_type,n_fund_op_type sql.NullInt32
		err=rows.Scan(
			&n_bal_id,
			&rec.ContractAid,
			&n_evtlog_id,
			&n_from_aid,
			&n_to_aid,
			&n_user_aid,
			&n_timestamp,
			&n_datetime,
			&n_tx_id,
			&n_txhash,
			&n_amount,
			&n_int_amount,
			&n_bal_usd,
			&n_balance,
			&n_bs_id,
			&n_bs_user_aid,
			&n_bs_op_type,
			&n_fund_id,
			&n_fund_op_type,
			&n_funder_aid,
			&n_red_id,
			&n_red_payout,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}

		rec.EvtlogId=n_evtlog_id.Int64
		if n_tx_id.Int64 != boundary_tx {
			// reset current boundary on transaction change
			b.EvtLogId = 0
			b.OpType = 0
			b.Offset = counter
			b.Len = 0
			boundary_tx = n_tx_id.Int64
			ss.Info.Printf("Event boundary was reset at TxId %v\n",n_tx_id.Int64)
		}
		counter++
		if n_bal_id.Valid { rec.BalChgId = n_bal_id.Int64 }
		if n_from_aid.Valid { rec.FromAid = n_from_aid.Int64 }
		if n_to_aid.Valid { rec.ToAid = n_to_aid.Int64 }
		if n_user_aid.Valid { rec.UserAid = n_user_aid.Int64 }
		if n_timestamp.Valid { rec.TimeStamp = n_timestamp.Int64 }
		if n_datetime.Valid { rec.DateTime = n_datetime.String }
		if n_tx_id.Valid { rec.TxId = n_tx_id.Int64 }
		if n_txhash.Valid { rec.TxHash = n_txhash.String }
		if n_amount.Valid { rec.Amount = n_amount.Float64 }
		if n_int_amount.Valid { rec.IntegerAmount = n_int_amount.Float64 }
		if n_bal_usd.Valid { rec.Balance = n_bal_usd.Float64 }
		if n_balance.Valid { rec.IntegerBalance = n_balance.Float64 }
		if n_bs_id.Valid { rec.BuySellOpId = n_bs_id.Int64 }
		if n_bs_op_type.Valid { rec.BuySellOpType = n_bs_op_type.Int32 }
		if n_fund_id.Valid { rec.FundOpId = n_fund_id.Int64 }
		if n_fund_op_type.Valid { rec.FundOpType = n_fund_op_type.Int32 }
		if n_red_id.Valid { rec.RedeemId = n_red_id.Int64 }
		if n_red_payout.Valid { rec.RedeemPayout = n_red_payout.Float64 }
		// push key of the address for subsequent fetch of the address itself
		if n_from_aid.Valid { addresses[n_from_aid.Int64] = "" }
		if n_to_aid.Valid { addresses[n_to_aid.Int64] = "" }
		ss.Info.Printf("tx_id %v (boundary tx_id=%v), evtlog_id %v, bal id=%v fund_id=%v, bs_id=%v bs_type=%v\n",n_tx_id.Int64,boundary_tx,n_evtlog_id.Int64,n_bal_id.Int64,n_fund_id.Int64,n_bs_id.Int64,n_bs_op_type.Int32)

		tmp_split_op,split_exists := split_ops[n_evtlog_id.Int64]
		if split_exists {
			ss.Info.Printf("Split exists (evtlog_id=%v)\n",n_evtlog_id.Int64)
			last_split_op = tmp_split_op
			var empty_op TmpMergeOp
			last_merge_op = empty_op
			b.EvtLogId = last_split_op.EvtLogId
			b.OpType = 1		// Split
			b.Len = counter - b.Offset
			ss.Info.Printf("Added split with boundary len = %v last evtlog_id= %v (type=%v)\n",b.Len,n_evtlog_id.Int64,b.OpType)
			boundaries = append(boundaries,b)
			b.Offset = counter; b.Len = 0; b.OpType = 0;
		}
		tmp_merge_op,merge_exists := merge_ops[n_evtlog_id.Int64]
		if merge_exists {
			ss.Info.Printf("Merge exists (evtlog_id=%v\n",n_evtlog_id.Int64)
			last_merge_op = tmp_merge_op
			var empty_op TmpSplitOp
			last_split_op = empty_op
			b.EvtLogId = last_merge_op.EvtLogId
			b.OpType = 2		// Merge
			b.Len = counter - b.Offset
			ss.Info.Printf("Added merge with boundary len = %v last evtlog_id= %v (type=%v)\n",b.Len,n_evtlog_id.Int64,b.OpType)
			boundaries = append(boundaries,b)
			b.Offset = counter; b.Len = 0; b.OpType = 0;
		}
		bs_op,buysell_exists := buysell_ops[n_evtlog_id.Int64]
		if buysell_exists {
			ss.Info.Printf("BuySell exists (evtlog_id=%v)\n",n_evtlog_id.Int64)
			b.EvtLogId = bs_op.EvtLogId
			b.OpType = 3		// Buy/Sell operation
			b.Len = counter - b.Offset
			boundaries = append(boundaries,b)
			ss.Info.Printf("Added Sell boundary len = %v , last evtlog_id=%v (type=%v)\n",b.Len,n_evtlog_id.Int64,b.OpType)
			b.Offset = counter; b.Len = 0; b.OpType = 0;
			ss.Info.Printf(
				"last_split_op.AddressId=%v,n_bs_user_aid=%v,last_split_op.TxId=%v,n_tx_id=%v\n",
				last_split_op.AddressId,n_bs_user_aid.Int64,last_split_op.TxId,n_tx_id.Int64,
			)
			if ((last_split_op.AddressId==n_bs_user_aid.Int64) && (last_split_op.TxId==n_tx_id.Int64)) ||
				((last_split_op.AddressId==contract_aid) && (last_split_op.TxId==n_tx_id.Int64)) {
				tmp_entry,_ := split_ops[last_split_op.EvtLogId]
				tmp_entry.ReferenceId=bs_op.EvtLogId
				tmp_entry.ReferenceType=1//Buy/Sell
				ss.Info.Printf("Set ReferenceId %v for split op of event %v (ReferenceType=1 Buy/Sell)\n",tmp_entry.ReferenceId,last_split_op.EvtLogId)
				split_ops[last_split_op.EvtLogId]=tmp_entry
			}
			ss.Info.Printf(
				"last_merge_op.AddressId=%v,n_bs_user_aid=%v,last_merge_op.TxId=%v,n_tx_id=%v\n",
				last_merge_op.AddressId,n_bs_user_aid.Int64,last_merge_op.TxId,n_tx_id.Int64,
			)
			if ((last_merge_op.AddressId==n_bs_user_aid.Int64) && (last_merge_op.TxId==n_tx_id.Int64)) ||
				((last_merge_op.AddressId==contract_aid) && (last_merge_op.TxId==n_tx_id.Int64)) {
				tmp_entry,_ := merge_ops[last_merge_op.EvtLogId]
				tmp_entry.ReferenceId=bs_op.EvtLogId
				tmp_entry.ReferenceType=1//Buy/Sell
				ss.Info.Printf("Set ReferenceId %v for merge op of event %v (ReferenceType=1 Buy/Sell)\n",tmp_entry.ReferenceId,last_merge_op.EvtLogId)
				merge_ops[last_merge_op.EvtLogId]=tmp_entry
			}
		}
		fund_op,fundop_exists := fund_ops[n_evtlog_id.Int64]
		if fundop_exists {
			ss.Info.Printf("Fund Add/Rem exists (evtlog_id=%v)\n",n_evtlog_id.Int64)
			b.EvtLogId = fund_op.EvtLogId
			b.OpType = 4		// Fund Add/Remove operation
			b.Len = counter - b.Offset
			boundaries = append(boundaries,b)
			b.Offset = counter; b.Len = 0; b.OpType = 0;
			ss.Info.Printf(
				"last_split_op.AddressId=%v,n_funder_aid=%v,last_split_op.TxId=%v,n_tx_id=%v\n",
				last_split_op.AddressId,n_funder_aid.Int64,last_split_op.TxId,n_tx_id.Int64,
			)
			if ((last_split_op.AddressId==n_funder_aid.Int64) && (last_split_op.TxId==n_tx_id.Int64)) ||
				((last_split_op.AddressId==contract_aid) && (last_split_op.TxId==n_tx_id.Int64)) {
				tmp_entry,_ := split_ops[last_split_op.EvtLogId]
				tmp_entry.ReferenceId=fund_op.EvtLogId
				tmp_entry.ReferenceType=2//Fund Add/Remove
				ss.Info.Printf("Set ReferenceId %v for split op of event %v (ReferenceType=2 Fund Add/Remove)\n",tmp_entry.ReferenceId,last_split_op.EvtLogId)
				split_ops[last_split_op.EvtLogId]=tmp_entry
			}
			ss.Info.Printf(
				"last_merge_op.AddressId=%v,n_funder_aid=%v,last_merge_op.TxId=%v,n_tx_id=%v\n",
				last_merge_op.AddressId,n_funder_aid.Int64,last_merge_op.TxId,n_tx_id.Int64,
			)
			if ((last_merge_op.AddressId==n_funder_aid.Int64) && (last_merge_op.TxId==n_tx_id.Int64)) ||
				((last_merge_op.AddressId==contract_aid) && (last_merge_op.TxId==n_tx_id.Int64)) {
				tmp_entry,_ := merge_ops[last_merge_op.EvtLogId]
				tmp_entry.ReferenceId=fund_op.EvtLogId
				tmp_entry.ReferenceType=2//Fund Add/Remove
				ss.Info.Printf("Set ReferenceId %v for merge op of event %v (ReferenceType=2 Fund Add/Remove)\n",tmp_entry.ReferenceId,last_merge_op.EvtLogId)
				merge_ops[last_merge_op.EvtLogId]=tmp_entry
			}
		}
		redeem_op,redeem_exists := redeem_ops[n_evtlog_id.Int64]
		if redeem_exists {
			b.EvtLogId = redeem_op.EvtLogId
			b.OpType = 5		// Redeem operation
			b.Len = counter - b.Offset
			ss.Info.Printf("Added redeem with boundary len = %v last evtlog_id= %v (type=%v)\n",b.Len,n_evtlog_id.Int64,b.OpType)
			boundaries = append(boundaries,b)
			b.Offset = counter; b.Len = 0; b.OpType = 0;
		}
		if (resolution_evtlog_id>0) && (n_evtlog_id.Int64>resolution_evtlog_id) {
			if !separator_was_added {
				b.EvtLogId = n_evtlog_id.Int64
				b.OpType = 6		// Market Resolution
				b.Len = counter - b.Offset
				boundaries = append(boundaries,b)
				b.Offset = counter; b.Len = 0; b.OpType = 0;
				separator_was_added = true
			}
		}
		records = append(records,rec)
	}

	// We query addresses separately because doing this in the main
	// query makes it run into infinity due to sequential scans
	// (bad postgres query optimization is the cause)
	var keys string
	for k,_ := range addresses {
		if len(keys) > 0 { keys = keys + "," }
		keys = keys + fmt.Sprintf("%v",k)
	}
	if len(keys) >0 {
		query_addrs := "SELECT address_id,addr FROM address WHERE address_id IN("+keys+")"
		rows_addrs,err := ss.db.Query(query_addrs)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query_addrs))
			os.Exit(1)
		}
		defer rows_addrs.Close()
		for rows_addrs.Next() {
			var addr string
			var aid int64
			err=rows_addrs.Scan(&aid,&addr)
			if err != nil {
				ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query_addrs))
				os.Exit(1)
			}
			addresses[aid]=addr
		}
	}
	for i:=0; i<len(records); i++ {
		rec := records[i]
		rec.FromAddr = addresses[rec.FromAid]
		rec.ToAddr = addresses[rec.ToAid]
		records[i]=rec
	}

	output := make([]p.API_Pol_OpenInterestHistory,0,512)
	num_boundaries := int32(len(boundaries))
	ss.Info.Printf("Num boundaries=%v\n",num_boundaries)
	counter = 0
	for {
		if counter>=num_boundaries {
			break;
		}
		b := &boundaries[counter]
		counter++
		ss.Info.Printf("boundary type = %v, counter=%v offset=%v, len=%v\n",b.OpType,counter,b.Len,b.Offset)
		switch b.OpType {
			case 1:		// Split
				ss.Info.Printf("Processing Split\n")
				split_op := split_ops[b.EvtLogId]
				ss.Info.Printf("Ref for split evtlog_id=%v is %v\n",b.EvtLogId,split_op.ReferenceId)
				for i:=int32(0); i<b.Len; i++ {
					idx := b.Offset + i
					in_rec := &records[idx]
					ss.Info.Printf("Processing evtlog_id %v (i=%v, balChgId=%v, tx_id=%v)\n",in_rec.EvtlogId,i,in_rec.BalChgId,in_rec.TxId)
					if in_rec.Amount<0.0 {
						ss.Info.Printf("Skipping evtlog_id %v (rule 1)\n",in_rec.EvtlogId)
						continue
					}
					if (in_rec.FromAid == contract_aid) && (in_rec.ToAid==condtok_aid) &&
						(in_rec.Amount<0.0) {
						ss.Info.Printf("Skipping evtlog_id %v (rule 2)\n",in_rec.EvtlogId)
						continue
					}
					if in_rec.FromAddr == "0x0000000000000000000000000000000000000000" {
						ss.Info.Printf("Skipping evtlog_id %v (rule 4)\n",in_rec.EvtlogId)
						continue
					}
					if in_rec.BalChgId == 0 {
						ss.Info.Printf("Skipping evtlog_id %v BalChgId=0\n",in_rec.EvtlogId)
						continue
					}
					if in_rec.ContractAid != usdc_aid {
						ss.Info.Printf("Skipping evtlkog_id %v ContractAid is not USDC\n",in_rec.EvtlogId)
						continue
					}
					ss.Info.Printf("Displaying evtlog %v\n",in_rec.EvtlogId)
					var out_rec p.API_Pol_OpenInterestHistory
					copy_all_fields(&out_rec,in_rec)

					ss.Info.Printf("split_op.ReferenceType = %v\n",split_op.ReferenceType)
					if split_op.ReferenceType == 1 { //Buy/Sell
						op,_ := buysell_ops[split_op.ReferenceId]
						out_rec.BuySellOpId = op.BuySellId
						out_rec.BuySellOpType = op.BuySellType
						ss.Info.Printf("op.BuySellType = %v\n",op.BuySellType)
						if op.BuySellType == 0 { // Buy
							ss.Info.Printf("Detected referenced buysell op\n")
							if (out_rec.ToAid == condtok_aid ) &&
									(out_rec.FromAid == contract_aid) && (op.BuySellType == 0) {
								// buy op
								out_rec.IntegerFee = op.IntegerFee
								out_rec.Fee = out_rec.IntegerFee/1000000.0
								fee_accum = fee_accum + out_rec.IntegerFee
								open_interest = open_interest +  out_rec.IntegerAmount
							} else {
								ss.Info.Printf("Condition #1 for identifying transfer wasn't met\n")
							}
						}
						if op.BuySellType == 1 {
							ss.Info.Printf("Detected referenced buysell op\n")
							if (out_rec.ToAid == contract_aid) &&
									(out_rec.FromAid == condtok_aid) && (op.BuySellType == 1) {
								// sell op
								ss.Info.Printf("Entering sell op, bal_id=%v tx_id=%v\n",out_rec.BalChgId,out_rec.TxId)
								ss.Info.Printf("Condition passed, calculating fees now IntegerAmount=%v\n",out_rec.IntegerAmount)
								out_rec.IntegerFee = op.IntegerFee
								out_rec.Fee = out_rec.IntegerFee / 1000000.0
								fee_accum = fee_accum + out_rec.IntegerFee
								open_interest = open_interest - (out_rec.IntegerAmount)
							} else {
								ss.Info.Printf("Condition for identifying transfer wasn't met\n")
							}
						}
					}
					if split_op.ReferenceType == 2 { //Fund Add/Remove
						op,_ := fund_ops[split_op.ReferenceId]
						ss.Info.Printf("op.FundType=%v\n",op.FundType)
						if op.FundType == 0 {// Add
							ss.Info.Printf("It is an add, ToAid=%v, condtok_aid=%v\n",out_rec.ToAid,condtok_aid)
							out_rec.FundOpId = op.FundOpId
							out_rec.FundOpType = op.FundType
							if out_rec.ToAid == condtok_aid {
								ss.Info.Printf("Set fund add id %v\n",op.FundOpId)
							}
							if (out_rec.ToAid == condtok_aid) && (out_rec.FromAid == contract_aid) {
								open_interest = open_interest + (out_rec.IntegerAmount)
							}
						}
						if op.FundType == 1 {// Remove
							ss.Info.Printf("Its an Remove, FromAid=%v, condtok_aid=%v\n",out_rec.FromAid,condtok_aid)
							out_rec.FundOpId = op.FundOpId
							out_rec.FundOpType = op.FundType
							if out_rec.FromAid == condtok_aid {
								ss.Info.Printf("Set fund remove id %v\n",op.FundOpId)
							}
							if (out_rec.FromAid == contract_aid) && (out_rec.ToAid != condtok_aid) {
								open_interest = open_interest - (-out_rec.IntegerAmount)
							} else {
							}
						}
					}
					if split_op.ReferenceType == 0 { // not linked with BuySell or Fund operation
						ss.Info.Printf("Outside Polymarket Buy (pure split) operation (evtlog_id=%v. bal_id=%v)\n",out_rec.EvtlogId,out_rec.BalChgId)
						ss.Info.Printf("buysell_id=%v, fund_id=%v, redeem_id=%v\n",out_rec.BuySellOpId,out_rec.FundOpId,out_rec.RedeemId)
						if (out_rec.FromAid != condtok_aid) && (out_rec.ToAid == condtok_aid) {
							// Sell
							open_interest = open_interest + (out_rec.IntegerAmount)
						} else {
							ss.Info.Printf("Couldn't calculate open interest (condition mismatch)\n")
						}
					}
					out_rec.AdjustedBalance = (out_rec.IntegerBalance - out_rec.IntegerAmount)/1000000.0
					out_rec.FeeAccum = fee_accum / 1000000.0
					out_rec.IntegerFeeAccum = fee_accum
					out_rec.OpenInterest = open_interest / 1000000.0
					out_rec.OIVerif = out_rec.OpenInterest + out_rec.FeeAccum
					totals.FinalOpenInterest = out_rec.OpenInterest
					totals.FinalFees = out_rec.FeeAccum
					out_rec.Note = make_note(out_rec.TxId,split_op.EvtLogId,0,out_rec.BuySellOpId,out_rec.BuySellOpType,out_rec.FundOpId,out_rec.FundOpType,out_rec.RedeemId,out_rec.FromAid,out_rec.ToAid,contract_aid,condtok_aid,out_rec.PayoutNumerators)
					output = append(output,out_rec)
				}
				ss.Info.Printf("Split boundary ended\n\n")
			case 2:		// Merge
				ss.Info.Printf("Processing merge\n")
				merge_op := merge_ops[b.EvtLogId]
				ss.Info.Printf("Ref for merge evtlog_id=%v is %v\n",b.EvtLogId,merge_op.ReferenceId)
				//var prev_trf_to_aid int64 = -1
				var prev_trf_from_aid int64 = -1
				var prev_trf_amount float64 = -1.0
				for i:=int32(0); i<b.Len; i++ {
					idx := b.Offset + i
					in_rec := &records[idx]
					ss.Info.Printf("Processing evtlog_id %v (i=%v, balChgId=%v, tx_id=%v)\n",in_rec.EvtlogId,i,in_rec.BalChgId,in_rec.TxId)
					if in_rec.BalChgId == 0 {
						ss.Info.Printf("Skipping evtlog_id %v BalChgId=0\n",in_rec.EvtlogId)
						continue
					}
					if in_rec.ContractAid != usdc_aid {
						ss.Info.Printf("Skipping evtlkog_id %v ContractAid is not USDC\n",in_rec.EvtlogId)
						continue
					}
					if (in_rec.Amount<0.0) {
						ss.Info.Printf("Skipping evtlog_id %v (negative balance)\n",in_rec.EvtlogId)
						prev_trf_from_aid = in_rec.FromAid
						prev_trf_amount = -in_rec.IntegerAmount // cancel sign
						continue
					}
					var out_rec p.API_Pol_OpenInterestHistory
					copy_all_fields(&out_rec,in_rec)
					ss.Info.Printf("merge_op.ReferenceType = %v\n",merge_op.ReferenceType)
					if merge_op.ReferenceType == 1 { //Buy/Sell
						op,_ := buysell_ops[merge_op.ReferenceId]
						ss.Info.Printf("op.BuySellType = %v\n",op.BuySellType)
						if (out_rec.FromAid == condtok_aid) && 
								(out_rec.ToAid == contract_aid) && (out_rec.BuySellOpType == 0) {
							// buy op
							out_rec.BuySellOpId = op.BuySellId
							out_rec.BuySellOpType = op.BuySellType
							out_rec.IntegerFee = prev_trf_amount - (out_rec.IntegerAmount)
							out_rec.Fee = out_rec.IntegerFee/1000000.0
							fee_accum = fee_accum + out_rec.IntegerFee
							open_interest = open_interest +  (out_rec.IntegerAmount)
						}
						if (out_rec.ToAid == contract_aid) &&
								(out_rec.FromAid == condtok_aid) && (out_rec.BuySellOpType == 1) {
							// sell op
							out_rec.BuySellOpId = op.BuySellId
							out_rec.BuySellOpType = op.BuySellType
							ss.Info.Printf("Entering sell op, bal_id=%v tx_id=%v prev_trf_from_aid=%v\n",out_rec.BalChgId,out_rec.TxId,prev_trf_from_aid)
							ss.Info.Printf("Condition passed, calculating fees now prev_trf_amount=%v, IntegerAmount=%v\n",prev_trf_amount,out_rec.IntegerAmount)
							//out_rec.IntegerFee = prev_trf_amount - out_rec.IntegerAmount
							//out_rec.Fee = out_rec.IntegerFee / 1000000.0
							//fee_accum = fee_accum + out_rec.IntegerFee
							open_interest = open_interest - (out_rec.IntegerAmount)
						}
					}
					if merge_op.ReferenceType == 2 { //Fund Add/Remove
						op,_ := fund_ops[merge_op.ReferenceId]
						ss.Info.Printf("op.FundType=%v\n",op.FundType)
						if op.FundType == 0 {// Add
							ss.Info.Printf("It is an add, ToAid=%v, condtok_aid=%v\n",out_rec.ToAid,condtok_aid)
							if out_rec.ToAid == condtok_aid {
								out_rec.FundOpId = op.FundOpId
								out_rec.FundOpType = op.FundType
								ss.Info.Printf("Set fund add id %v\n",op.FundOpId)
							}
							if (out_rec.ToAid == condtok_aid) && (out_rec.FromAid == contract_aid) {
								open_interest = open_interest + (out_rec.IntegerAmount)
							}
						}
						if op.FundType == 1 {// Remove
							ss.Info.Printf("Its an Remove, FromAid=%v, condtok_aid=%v\n",out_rec.FromAid,condtok_aid)
							if out_rec.FromAid == condtok_aid {
								out_rec.FundOpId = op.FundOpId
								out_rec.FundOpType = op.FundType
								ss.Info.Printf("Set fund remove id %v\n",op.FundOpId)
							}
							if (out_rec.FromAid == contract_aid) && (out_rec.ToAid != condtok_aid) {
								open_interest = open_interest - (-out_rec.IntegerAmount)
							} else {
							}
						}
					}
					if merge_op.ReferenceType == 0 { // not linked with BuySell or Fund operation
						ss.Info.Printf("Outside Polymarket Sell (pure merge) operation (evtlog_id=%v. bal_id=%v)\n",out_rec.EvtlogId,out_rec.BalChgId)
						if (out_rec.FromAid == condtok_aid) && (out_rec.ToAid != condtok_aid) {
							// Sell
							open_interest = open_interest - (out_rec.IntegerAmount)
						} else {
							ss.Info.Printf("Couldn't calculate open interest (condition mismatch)\n")
						}
					}
					out_rec.AdjustedBalance = (out_rec.IntegerBalance - out_rec.IntegerAmount)/1000000.0
					out_rec.FeeAccum = fee_accum / 1000000.0
					out_rec.IntegerFeeAccum = fee_accum
					out_rec.OpenInterest = open_interest / 1000000.0
					out_rec.OIVerif = out_rec.OpenInterest + out_rec.FeeAccum
					totals.FinalOpenInterest = out_rec.OpenInterest
					totals.FinalFees = out_rec.FeeAccum
					//prev_trf_to_aid = out_rec.ToAid
					//prevtrf_from_aid = out_rec.FromAid
					prev_trf_amount = out_rec.IntegerAmount
					ss.Info.Printf("EvtlogId=%v ids: buysell_op_id=%v, fund_op_id=%v, redeem_id=%v, split_id=%v, merge_id=%v\n",out_rec.EvtlogId,out_rec.BuySellOpId,out_rec.FundOpId,out_rec.RedeemId,0,merge_op.EvtLogId)
					out_rec.Note = make_note(out_rec.TxId,0,merge_op.EvtLogId,out_rec.BuySellOpId,out_rec.BuySellOpType,out_rec.FundOpId,out_rec.FundOpType,out_rec.RedeemId,out_rec.FromAid,out_rec.ToAid,contract_aid,condtok_aid,out_rec.PayoutNumerators)
					output = append(output,out_rec)
				}
				ss.Info.Printf("Merge boundary ended\n\n")
			case 3:		// Buy/Sell
				buysell_op := buysell_ops[b.EvtLogId]
				ss.Info.Printf("processing bysell op, b=%+v\n",*b)
				ss.Info.Printf("op.BuySellType = %v\n",buysell_op.BuySellType)
				var prev_trf_to_aid int64 = -1
				//var prev_trf_from_aid int64 = -1
				var prev_trf_amount float64 = -1.0
				for i:=int32(0); i<b.Len; i++ {
					idx := b.Offset + i
					in_rec := &records[idx]
					ss.Info.Printf("Processing evtlog_id %v (i=%v, balChgId=%v)\n",in_rec.EvtlogId,i,in_rec.BalChgId)
					ss.Info.Printf("iteration %v, BuySellOpType = %v\n",(i+1),in_rec.BuySellOpType)
					if in_rec.BalChgId == 0 {
						ss.Info.Printf("Skipping evtlog_id %v BalChgId=0\n",in_rec.EvtlogId)
						continue
					}
					if in_rec.ContractAid != usdc_aid {
						ss.Info.Printf("Skipping evtlkog_id %v ContractAid is not USDC\n",in_rec.EvtlogId)
						continue
					}
					if (in_rec.Amount<0.0) {
						ss.Info.Printf("Skipping evtlog_id %v (negative balance)\n",in_rec.EvtlogId)
						prev_trf_to_aid = in_rec.ToAid
						prev_trf_amount = -in_rec.IntegerAmount // cancel sign
						ss.Info.Printf("Setting previous amount to %v\n",prev_trf_amount)
						continue
					}
					var out_rec p.API_Pol_OpenInterestHistory
					copy_all_fields(&out_rec,in_rec)
					ss.Info.Printf("Sell op, integer amount = %v, optype=%v\n",out_rec.IntegerAmount,out_rec.BuySellOpType)
					if (out_rec.ToAid == condtok_aid) && 
						(out_rec.FromAid == contract_aid) && (buysell_op.BuySellType  == 0) {
						// buy op
						out_rec.BuySellOpId = buysell_op.BuySellId
						out_rec.BuySellOpType = buysell_op.BuySellType
						out_rec.IntegerFee = prev_trf_amount - (out_rec.IntegerAmount)
						out_rec.Fee = out_rec.IntegerFee/1000000.0
						fee_accum = fee_accum + out_rec.IntegerFee
						open_interest = open_interest +  (out_rec.IntegerAmount)
					}
					if (out_rec.FromAid == contract_aid) &&
							(out_rec.ToAid != condtok_aid) && (buysell_op.BuySellType == 1) {
						// sell op
						out_rec.BuySellOpId = buysell_op.BuySellId
						out_rec.BuySellOpType = buysell_op.BuySellType
						ss.Info.Printf("Entering sell op, bal_id=%v tx_id=%v prev_trf_to_aid=%v\n",out_rec.BalChgId,out_rec.TxId,prev_trf_to_aid)
						ss.Info.Printf("Condition passed, calculating fees now prev_trf_amount=%v, IntegerAmount=%v\n",prev_trf_amount,out_rec.IntegerAmount)
						out_rec.IntegerFee = buysell_op.IntegerFee
						out_rec.Fee = out_rec.IntegerFee / 1000000.0
						fee_accum = fee_accum + out_rec.IntegerFee
						open_interest = open_interest - (out_rec.IntegerAmount)
					}
					out_rec.AdjustedBalance = (out_rec.IntegerBalance - out_rec.IntegerAmount)/1000000.0
					out_rec.FeeAccum = fee_accum / 1000000.0
					out_rec.IntegerFeeAccum = fee_accum
					out_rec.OpenInterest = open_interest / 1000000.0
					out_rec.OIVerif = out_rec.OpenInterest + out_rec.FeeAccum
					totals.FinalOpenInterest = out_rec.OpenInterest
					totals.FinalFees = out_rec.FeeAccum
					//prev_trf_to_aid = out_rec.ToAid
					prev_trf_to_aid = out_rec.ToAid
					prev_trf_amount = out_rec.IntegerAmount
					ss.Info.Printf("EvtlogId=%v ids: buysell_op_id=%v, fund_op_id=%v, redeem_id=%v, split_id=%v, sellop_id=%v\n",out_rec.EvtlogId,out_rec.BuySellOpId,out_rec.FundOpId,out_rec.RedeemId,0,buysell_op.EvtLogId)
					out_rec.Note = make_note(out_rec.TxId,0,0,out_rec.BuySellOpId,out_rec.BuySellOpType,out_rec.FundOpId,out_rec.FundOpType,out_rec.RedeemId,out_rec.FromAid,out_rec.ToAid,contract_aid,condtok_aid,out_rec.PayoutNumerators)
					output = append(output,out_rec)
				}
				ss.Info.Printf("Buy/Sell boundary ended\n\n")
			case 4:		// Fund Add/Remove
				ss.Info.Printf("Entering Fund Add/Remove\n")
				fund_op := fund_ops[b.EvtLogId]
				for i:=int32(0); i<b.Len; i++ {
					idx := b.Offset + i
					in_rec := &records[idx]
					ss.Info.Printf("Processing evtlog_id %v (i=%v, balChgId=%v)\n",in_rec.EvtlogId,i,in_rec.BalChgId)
					if in_rec.BalChgId == 0 {
						ss.Info.Printf("Skipping evtlog_id %v BalChgId=0\n",in_rec.EvtlogId)
						continue
					}
					if in_rec.ContractAid != usdc_aid {
						ss.Info.Printf("Skipping evtlkog_id %v ContractAid is not USDC\n",in_rec.EvtlogId)
						continue
					}
					if (in_rec.Amount<0.0) {
						ss.Info.Printf("Skipping evtlog_id %v (negative balance)\n",in_rec.EvtlogId)
						continue
					}
					var out_rec p.API_Pol_OpenInterestHistory
					copy_all_fields(&out_rec,in_rec)
					if (out_rec.FromAid == contract_aid) &&
							(out_rec.UserAid == contract_aid) && (fund_op.FundType!=1) {
						ss.Info.Printf("Skipping tx_id=%v (rule 3)\n",out_rec.TxId)
						continue
					}
					out_rec.FundOpId = fund_op.FundOpId
					out_rec.FundOpType = fund_op.FundType
					if fund_op.FundType == 0 { // add funds
						ss.Info.Printf("Add funds operation (IntegerAmount=%v)\n",out_rec.IntegerAmount)
						if (out_rec.ToAid == condtok_aid) && (out_rec.FromAid == contract_aid) {
							open_interest = open_interest + out_rec.IntegerAmount
							ss.Info.Printf("Open Interest updated (addition)\n")
						} else {
							ss.Info.Printf("Condition to update open interest unmet\n")
						}

					}
					if fund_op.FundType == 1 { //withdraw funds
						ss.Info.Printf("Remove funds operation (IntegerAmount=%v)\n",out_rec.IntegerAmount)
						if (out_rec.FromAid == contract_aid) && (out_rec.ToAid != condtok_aid) {
							open_interest = open_interest - out_rec.IntegerAmount
							ss.Info.Printf("Open interest updated (subtraction)\n")
						} else {
							ss.Info.Printf("Condition to update open interest unmet\n")
						}
					}
					out_rec.AdjustedBalance = (out_rec.IntegerBalance - out_rec.IntegerAmount)/1000000.0
					out_rec.FeeAccum = fee_accum / 1000000.0
					out_rec.IntegerFeeAccum = fee_accum
					out_rec.OpenInterest = open_interest / 1000000.0
					out_rec.OIVerif = out_rec.OpenInterest + out_rec.FeeAccum
					totals.FinalOpenInterest = out_rec.OpenInterest
					totals.FinalFees = out_rec.FeeAccum
					out_rec.Note = make_note(out_rec.TxId,0,0,out_rec.BuySellOpId,out_rec.BuySellOpType,out_rec.FundOpId,out_rec.FundOpType,out_rec.RedeemId,out_rec.FromAid,out_rec.ToAid,contract_aid,condtok_aid,out_rec.PayoutNumerators)
					output = append(output,out_rec)
				}
				ss.Info.Printf("Fund Add/Remove boundary ended\n\n")
			case 5:		// Redeem
				ss.Info.Printf("Entering Redeem operation\n")
				redeem_op := redeem_ops[b.EvtLogId]
				for i:=int32(0); i<b.Len; i++ {
					idx := b.Offset + i
					in_rec := &records[idx]
					ss.Info.Printf("Processing evtlog_id %v (i=%v, balChgId=%v)\n",in_rec.EvtlogId,i,in_rec.BalChgId)
					if in_rec.BalChgId == 0 {
						ss.Info.Printf("Skipping evtlog_id %v BalChgId=0\n",in_rec.EvtlogId)
						continue
					}
					if in_rec.ContractAid != usdc_aid {
						ss.Info.Printf("Skipping evtlkog_id %v ContractAid is not USDC\n",in_rec.EvtlogId)
						continue
					}
					if (in_rec.Amount<0.0) {
						ss.Info.Printf("Skipping evtlog_id %v (negative balance)\n",in_rec.EvtlogId)
						continue
					}
					var out_rec p.API_Pol_OpenInterestHistory
					copy_all_fields(&out_rec,in_rec)

					open_interest = open_interest - out_rec.IntegerAmount
					out_rec.RedeemId = redeem_op.RedeemOpId

					out_rec.AdjustedBalance = (out_rec.IntegerBalance - out_rec.IntegerAmount)/1000000.0
					out_rec.FeeAccum = fee_accum / 1000000.0
					out_rec.IntegerFeeAccum = fee_accum
					out_rec.OpenInterest = open_interest / 1000000.0
					out_rec.OIVerif = out_rec.OpenInterest + out_rec.FeeAccum
					totals.FinalOpenInterest = out_rec.OpenInterest
					totals.FinalFees = out_rec.FeeAccum
					out_rec.Note = make_note(out_rec.TxId,0,0,out_rec.BuySellOpId,out_rec.BuySellOpType,out_rec.FundOpId,out_rec.FundOpType,out_rec.RedeemId,out_rec.FromAid,out_rec.ToAid,contract_aid,condtok_aid,out_rec.PayoutNumerators)
					output = append(output,out_rec)
				}
			case 6:		// Separator (market resolved event)
				idx := b.Offset
				in_rec := &records[idx]
				var out_rec p.API_Pol_OpenInterestHistory
				copy_all_fields(&out_rec,in_rec)
				var resolution_rec p.API_Pol_OpenInterestHistory
				resolution_rec.TxId = -1
				resolution_rec.PayoutNumerators = payout_numerators
				resolution_rec.DateTime = resolution_date
				resolution_rec.Note = make_note(-1,0,0,resolution_rec.BuySellOpId,resolution_rec.BuySellOpType,resolution_rec.FundOpId,resolution_rec.FundOpType,resolution_rec.RedeemId,resolution_rec.FromAid,resolution_rec.ToAid,contract_aid,condtok_aid,resolution_rec.PayoutNumerators)
				output = append(output,resolution_rec)
				separator_was_added=true
			default:
				ss.Info.Printf("Unkown boundary operation type %v\n",b.OpType)
		}
	}
	ss.Info.Printf("len(output)=%v\n",len(output))
	return totals,output
}
func copy_all_fields(out_rec,in_rec *p.API_Pol_OpenInterestHistory) {

	out_rec.EvtlogId =				in_rec.EvtlogId
	out_rec.ContractAid =			in_rec.ContractAid
	out_rec.TimeStamp =				in_rec.TimeStamp
	out_rec.DateTime =				in_rec.DateTime
	out_rec.TxId =					in_rec.TxId
	out_rec.TxHash =				in_rec.TxHash
	out_rec.FromAid =				in_rec.FromAid
	out_rec.FromAddr =				in_rec.FromAddr
	out_rec.ToAid =					in_rec.ToAid
	out_rec.UserAid =				in_rec.UserAid
	out_rec.ToAddr =				in_rec.ToAddr
	out_rec.PayoutNumerators =		in_rec.PayoutNumerators
	out_rec.BalChgId =				in_rec.BalChgId
	out_rec.BuySellOpId =			in_rec.BuySellOpId
	out_rec.BuySellOpType =			in_rec.BuySellOpType
	out_rec.FundOpId =				in_rec.FundOpId
	out_rec.FundOpType =			in_rec.FundOpType
	out_rec.RedeemId =				in_rec.RedeemId
	out_rec.RedeemIntegerPayout =	in_rec.RedeemIntegerPayout
	out_rec.RedeemPayout =			in_rec.RedeemPayout
	out_rec.Amount =				in_rec.Amount
	out_rec.IntegerAmount =			in_rec.IntegerAmount
	out_rec.Balance =				in_rec.Balance
	out_rec.AdjustedBalance =		in_rec.AdjustedBalance
	out_rec.IntegerBalance =		in_rec.IntegerBalance
	out_rec.OpenInterest =			in_rec.OpenInterest
	out_rec.OIVerif =				in_rec.OIVerif
	out_rec.IntegerFee =			in_rec.IntegerFee
	out_rec.Fee =					in_rec.Fee
	out_rec.FeeAccum =				in_rec.FeeAccum
	out_rec.IntegerFeeAccum =		in_rec.IntegerFeeAccum
	out_rec.ContractBalance =		in_rec.ContractBalance
	out_rec.ContractBalanceAccum =	in_rec.ContractBalanceAccum
	out_rec.Note =					in_rec.Note
}
func update_fields(rec *p.API_Pol_OpenInterestHistory,op_type int32) {

}
func (ss *SQLStorage) Get_polymarket_user_info(user_aid int64) (p.API_Pol_UserInfo,error){

	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM reg_time_stamp)::BIGINT AS reg_ts,"+
				"s.reg_time_stamp reg_datetime, " +
				"s.markets_count," +
				"s.tot_trades,"+
				"s.tot_liq_ops,"+
				"s.tot_volume/1e+6,"+
				"s.tot_liq_given/1e+6,"+
				"s.tot_fees/1e+6,"+
				"s.profit/1e+6, "+
				"ua.addr "+
			"FROM pol_ustats s "+
				"LEFT JOIN address ua ON s.user_aid=ua.address_id "+
			"WHERE user_aid=$1"

	res := ss.db.QueryRow(query,user_aid)
	var rec p.API_Pol_UserInfo
	var n_tot_mkt_count,n_tot_trades, n_tot_liq_ops,n_timestamp sql.NullInt64
	var n_volume,n_liq_given,n_fees,n_profit sql.NullFloat64
	var n_datetime sql.NullString
	err := res.Scan(
		&n_timestamp,
		&n_datetime,
		&n_tot_mkt_count,
		&n_tot_trades,
		&n_tot_liq_ops,
		&n_volume,
		&n_liq_given,
		&n_fees,
		&n_profit,
		&rec.Address,
	)
	rec.Aid= user_aid
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return rec,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	if n_timestamp.Valid { rec.TimeStampRegistered = n_timestamp.Int64 }
	if n_datetime.Valid { rec.DateTimeRegistered = n_datetime.String }
	if n_tot_mkt_count.Valid { rec.TotalMarketsTraded = n_tot_mkt_count.Int64 }
	if n_tot_trades.Valid { rec.TotalTrades = n_tot_trades.Int64 }
	if n_tot_liq_ops.Valid { rec.TotalLiquidityOps  = n_tot_liq_ops.Int64 }
	if n_volume.Valid { rec.TotalVolume = n_volume.Float64 }
	if n_liq_given.Valid { rec.TotalLiquidityFunded = n_liq_given.Float64 }
	if n_fees.Valid { rec.TotalFees = n_fees.Float64 }
	if n_profit.Valid { rec.TotalProfit = n_profit.Float64 }
	return rec,nil
}
func (ss *SQLStorage) Get_polymarket_markets_by_user(user_aid int64) []p.API_Pol_MarketsByUser {

	records := make([]p.API_Pol_MarketsByUser,0,32)
	var query string
	query = "SELECT "+
				"s.contract_aid, " +
				"m.market_id," +
				"m.question," +
				"EXTRACT(EPOCH FROM start_date_ts)::BIGINT AS start_date_ts," +
				"m.start_date," +
				"s.tot_volume/1e+6 vol, " +
				"s.tot_trades, "+
				"s.tot_liq_ops "+
			"FROM pol_ustats_mkt s " +
				"JOIN pol_market m ON s.contract_aid=m.mkt_mkr_aid " +
			"WHERE s.user_aid=$1 " +
			"ORDER BY contract_aid"

	rows,err := ss.db.Query(query,user_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_MarketsByUser
		err=rows.Scan(
			&rec.MarketMakerAid,
			&rec.MarketId,
			&rec.Question,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TotalVolume,
			&rec.TotalTrades,
			&rec.TotalLiquidityOperations,
		)
		records = append(records,rec)
	}
	return records

}
