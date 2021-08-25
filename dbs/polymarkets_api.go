package dbs
import (
	"os"
	"fmt"
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
				"prep.question_id, "+
				"prep.outcome_slot_count, " +
				"prep.tx_hash " +
			"FROM pol_market pm " +
				"JOIN address ma ON pm.mkt_mkr_aid=ma.address_id " +
				"JOIN pol_mkt_stats mst ON pm.mkt_mkr_aid=mst.contract_aid " +
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
	return rec,nil
}
func (ss *SQLStorage) Get_polymarkets_markets(status,sort int) []p.API_Pol_MarketInfo {
	// status: 0 - all markets, 1 - not finalized, 2 - finalized
	// sort : 0 - by trading volume, 1 - by liquidity invested, 2-by creation date, 3-by resolution date, 4 - fees collected
	var where_condition string
	if status == 1 {
		where_condition = "WHERE res.id IS NULL "
	}
	if status == 2 {
		where_condition = "WHERE res.id IS NOT NULL "
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
	rows,err := ss.db.Query(query)
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
		if n_open_interest.Valid { rec.OpenInterest = -n_open_interest.Float64 }
		if n_liquidity.Valid { rec.Liquidity = n_liquidity.Float64 }
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
			fmt.Printf("Adjusting prices, addr %v, Current Price=%v\n",records[i].UserAddr,records[i].CurrentPrice)
			pos_value := records[i].CurrentPrice * records[i].CurrentBalance
			fmt.Printf("position value = %v realized profit = %v\n",pos_value,records[i].RealizedProfit)
			profit := pos_value + records[i].TotalProfit // it is a + because User's deposits are negative
			fmt.Printf("Profit = %v\n",profit)
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
				"ptk.token_id_hex,"+
				"eh.cur_balance::TEXT " +
			"FROM pol_tok_ids ptk "+
				"JOIN erc1155_tok et ON ptk.token_id_hex=et.token_id_hex " +
				"JOIN erc1155_holder eh ON (et.token_id=eh.token_id AND eh.aid=ptk.contract_aid) " +
			"WHERE "+
				"ptk.contract_aid=$1"

	fmt.Printf("Query = %v (contract=%v)\n",query,contract_aid)

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
			&rec.TokenIdHex,
			&rec.TokenBalanceStr,
		)
		fmt.Printf("Adding outcome %v token hex %v and balance %v\n",rec.OutcomeIdx,rec.TokenIdHex,rec.TokenBalanceStr)
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
	fmt.Printf("rec_len = %v\n",rec_len)
	prices := make([]*big.Float,0,rec_len)
	for i:=0; i<rec_len ; i++ {
		numerator := odds_weight_for_outcome_func(i,balances)
		fmt.Printf("Numerator = %v\n",numerator.String())
		denominator := big.NewInt(0)
		for j:=0; j<rec_len; j++ {
			odds_outcome := odds_weight_for_outcome_func(j,balances)
			fmt.Printf("Adding %v to denominator\n",odds_outcome.String())
			denominator.Add(denominator,odds_outcome)
		}
		fmt.Printf("Denominator = %v\n",denominator.String())
		p := new(big.Float)
		numerator_float := new(big.Float)
		denominator_float := new(big.Float)
		numerator_float.SetInt(numerator)
		denominator_float.SetInt(denominator)
		p.Quo(numerator_float,denominator_float)
		fmt.Printf("Price for outcome %v = %v\n",i,p.String())
		prices = append(prices,p)
	}
	for i:=0; i<rec_len; i++ {
		f,_ := prices[i].Float64()
		records[i].TokenPrice = f
		fmt.Printf("Adjusted token price = %v\n",records[i].TokenPrice)
		f,_= strconv.ParseFloat(records[i].TokenBalanceStr,64)
		records[i].TokenBalance = f
		fmt.Printf("Price converted to float = %v\n",f)
	}
	fmt.Printf("exciting Calculate_prices()\n")
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
				"FROM poly_uranks AS r " +
					"JOIN  pol_ustats AS s ON r.aid=s.aid " +
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
