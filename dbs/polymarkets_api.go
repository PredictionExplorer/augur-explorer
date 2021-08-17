package dbs
import (
	"os"
	"fmt"
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
				"condition_id," +
				"slug," +
				"resolution_source,"+
				"EXTRACT(EPOCH FROM created_at_ts)::BIGINT AS created_at_ts,"+
				"created_at_date, " +
				"EXTRACT(EPOCH FROM end_date_ts)::BIGINT AS ts_end," +
				"end_date," +
				"EXTRACT(EPOCH FROM start_date_ts)::BIGINT AS ts_start," +
				"start_date," +
				"category," +
				"image," +
				"icon," +
				"description," +
				"tags," +
				"outcomes," +
				"active," +
				"market_type," +
				"market_type_code," +
				"closed," +
				"mkt_mkr_aid," +
				"ma.addr AS mkt_mkr_addr " +
			"FROM pol_market pm " +
				"JOIN address ma ON pm.mkt_mkr_aid=ma.address_id " +
			"WHERE pm.market_id=$1"

	res := ss.db.QueryRow(query,market_id)
	err := res.Scan(
			&rec.Question,
			&rec.ConditionId,
			&rec.Slug,
			&rec.ResolutionSource,
			&rec.CreatedAtTs,
			&rec.CreatedAtDate,
			&rec.EndDateTs,
			&rec.EndDate,
			&rec.StartDateTs,
			&rec.StartDate,
			&rec.Category,
			&rec.Image,
			&rec.Icon,
			&rec.Description,
			&rec.Tags,
			&rec.Outcomes,
			&rec.Active,
			&rec.MarketType,
			&rec.MarketTypeCode,
			&rec.Closed,
			&rec.MarketMakerAid,
			&rec.MarketMakerAddr,
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
	return rec,nil
}
func (ss *SQLStorage) Get_polymarkets_markets(status int) []p.API_Pol_MarketInfo {
	// status: 0 - all markets, 1 - not finalized, 2 - finalized

	var where_condition string
	if status == 1 {
		where_condition = "WHERE closed=FALSE "
	}
	if status == 2 {
		where_condition = "WHERE closed=TRUE "
	}
	records := make([]p.API_Pol_MarketInfo,0,32)
	var query string
	query = "SELECT " +
				"market_id," +
				"question," +
				"condition_id," +
				"slug," +
				"resolution_source,"+
				"EXTRACT(EPOCH FROM created_at_ts)::BIGINT AS created_at_ts,"+
				"created_at_date, " +
				"EXTRACT(EPOCH FROM end_date_ts)::BIGINT AS ts_end," +
				"end_date," +
				"EXTRACT(EPOCH FROM start_date_ts)::BIGINT AS ts_start," +
				"start_date," +
				"category," +
				"image," +
				"icon," +
				"description," +
				"tags," +
				"outcomes," +
				"active," +
				"market_type," +
				"market_type_code," +
				"closed," +
				"mkt_mkr_aid," +
				"ma.addr AS mkt_mkr_addr " +
			"FROM pol_market pm " +
				"JOIN address ma ON pm.mkt_mkr_aid=ma.address_id " +
			where_condition +
			"ORDER BY market_id DESC"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.API_Pol_MarketInfo
		err=rows.Scan(
			&rec.MarketId,
			&rec.Question,
			&rec.ConditionId,
			&rec.Slug,
			&rec.ResolutionSource,
			&rec.CreatedAtTs,
			&rec.CreatedAtDate,
			&rec.EndDateTs,
			&rec.EndDate,
			&rec.StartDateTs,
			&rec.StartDate,
			&rec.Category,
			&rec.Image,
			&rec.Icon,
			&rec.Description,
			&rec.Tags,
			&rec.Outcomes,
			&rec.Active,
			&rec.MarketType,
			&rec.MarketTypeCode,
			&rec.Closed,
			&rec.MarketMakerAid,
			&rec.MarketMakerAddr,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Error in Get_polymarkets_markets(): %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_poly_market_open_positions(contract_aid int64) []p.API_Pol_MarketOpenPosition {

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
			"WHERE (ptk.contract_aid = $1) AND (ROUND(eh.cur_balance/1e+4)>0) " +
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
			&rec.TotalProfit,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Error in Get_polymarkets_markets(): %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_poly_market_user_open_positions(user_aid int64) []p.API_Pol_MarketUserOpenPosition {

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
	return nil
}
