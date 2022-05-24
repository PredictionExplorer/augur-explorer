
package balancerv2

import (
	"os"
	"fmt"
	"time"
	"math"

	"database/sql"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives/balancerv2"
	pr "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (sw *SQLStorageWrapper) Get_first_last_swap_timestamp(pool_aid int64, is_last bool) int64 {

	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts " +
			"FROM swf_hist "+
			"WHERE pool_aid=$1 "+
			"ORDER BY time_stamp "
	if is_last {
		query = query +	"DESC "
	} else {
		query = query + "ASC "
	}
	query = query + "LIMIT 1"

	row := sw.S.Db().QueryRow(query,pool_aid)
	var err error
	var ts int64
	err=row.Scan(&ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_first_last_swap_timestamp(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return ts
}
func (sw *SQLStorageWrapper) Get_first_last_swap_timestamp_all_pools(is_last bool) int64 {

	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts " +
			"FROM swf_hist "+
			"ORDER BY time_stamp "
	if is_last {
		query = query +	"DESC "
	} else {
		query = query + "ASC "
	}
	query = query + "LIMIT 1"

	row := sw.S.Db().QueryRow(query)
	var err error
	var ts int64
	err=row.Scan(&ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_first_last_swap_timestamp_all_pools(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return ts
}
func (sw *SQLStorageWrapper) Get_pool_info(pool_id string) p.BalV2PoolInfo {

	var output p.BalV2PoolInfo
	var query string
	query = "SELECT " +
				"block_num,"+
				"p.pool_aid,"+
				"pa.addr,"+
				"specialization "+
			"FROM "+sw.S.SchemaName()+".pool_reg p "+
				"JOIN "+sw.S.SchemaName()+".addr pa ON p.pool_aid=pa.address_id "+
			"WHERE p.pool_id=$1"


	row := sw.S.Db().QueryRow(query,pool_id)
	var err error
	err=row.Scan(
		&output.BlockNum,
		&output.PoolAid,
		&output.PoolAddr,
		&output.Specialization,
	);
	output.PoolId=pool_id
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_pool_info(): %v, q=%v",err,query))
		os.Exit(1)
	}
	query = "SELECT comments FROM unhandled WHERE pool_id=$1"
	row = sw.S.Db().QueryRow(query,pool_id)
	var comments string
	err=row.Scan(&comments)
	if err != nil {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_pool_info(): %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		output.Unhandled = true
		output.UnhandledComments = comments
	}
	output.FirstSwapTs = sw.Get_first_last_swap_timestamp(output.PoolAid,false)
	output.LastSwapTs = sw.Get_first_last_swap_timestamp(output.PoolAid,true)
	return output
}
func (sw *SQLStorageWrapper) Get_pool_total_swaps(pool_id string) int64 {

	var query string
	query = "SELECT "+
					"count(*) totswaps "+
				"FROM "+sw.S.SchemaName()+".swap "+
				"WHERE pool_id=$1"

	row := sw.S.Db().QueryRow(query,pool_id)
	var err error
	var total_swaps int64
	err=row.Scan(&total_swaps)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_pool_info(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return total_swaps
}
func (sw *SQLStorageWrapper) Get_pool_swap_history_backwards(pool_aid,offset,limit int64) []p.BalV2SwapRecordInfo{

	records := make([]p.BalV2SwapRecordInfo,0,32)
	var query string
	query = "SELECT "+
				"sh.block_num,"+
				"EXTRACT(EPOCH FROM sh.time_stamp)::BIGINT ts," +
				"sh.time_stamp,"+
				"s.token_in_aid,"+
				"s.token_out_aid,"+
				"ia.addr,"+
				"oa.addr,"+
				"s.amount_in,"+
				"s.amount_in,"+
				"s.amount_out, "+
				"s.amount_out, "+
				"sh.swap_fee_usd, "+
				"ii.decimals, "+
				"io.decimals, "+
				"ii.symbol,"+
				"io.symbol "+
			"FROM "+sw.S.SchemaName()+".swf_hist sh "+
				"JOIN "+sw.S.SchemaName()+".swap s ON (s.block_num=sh.block_num) AND (s.tx_index=sh.tx_index) AND (s.log_index=sh.log_index) "+
				"JOIN "+sw.S.SchemaName()+".tok_bal tb ON (s.block_num=tb.block_num) AND (s.tx_index=tb.tx_index) AND (s.log_index=tb.log_index) AND (s.token_in_aid=tb.tok_aid) "+
				"JOIN "+sw.S.SchemaName()+".addr ia ON s.token_in_aid=ia.address_id "+
				"JOIN "+sw.S.SchemaName()+".addr oa ON s.token_out_aid=oa.address_id "+
				"LEFT JOIN erc20_info ii ON ii.token_aid=s.token_in_aid "+
				"LEFT JOIN erc20_info io ON io.token_aid=s.token_out_aid "+
			"WHERE sh.pool_aid=$1 "+
			"ORDER BY sh.id DESC "+
			"OFFSET $2 LIMIT $3"

	rows,err := sw.S.Db().Query(query,pool_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.BalV2SwapRecordInfo
		var amount_in_f64,amount_out_f64 float64
		var n_decimals_in,n_decimals_out sql.NullInt64
		var n_sym_in,n_sym_out sql.NullString
		err=rows.Scan(
			&rec.BlockNum,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TokenInAid,
			&rec.TokenOutAid,
			&rec.TokenInAddr,
			&rec.TokenOutAddr,
			&rec.AmountIn,
			&amount_in_f64,
			&rec.AmountOut,
			&amount_out_f64,
			&rec.USDValue,
			&n_decimals_in,
			&n_decimals_out,
			&n_sym_in,
			&n_sym_out,
		)
		if err!=nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if amount_in_f64 == 0.0 {
			continue // skip invalid swaps (and avoid division by 0)
		}
		if n_decimals_in.Valid { rec.DecimalsTokIn = n_decimals_in.Int64 }
		if n_decimals_out.Valid { rec.DecimalsTokOut = n_decimals_out.Int64 }
		if n_sym_in.Valid { rec.SymbolIn = n_sym_in.String }
		if n_sym_out.Valid { rec.SymbolOut = n_sym_out.String }
		rec.TokenInAddrShort=pr.Short_address(rec.TokenInAddr)
		rec.TokenOutAddrShort=pr.Short_address(rec.TokenOutAddr)
		if rec.DecimalsTokIn > 0 {
			var denominator float64 = math.Pow(10,float64(rec.DecimalsTokIn))
			var result = amount_in_f64/denominator
			rec.AmountInFmt = fmt.Sprintf("%.2f",result)
		}
		if rec.DecimalsTokOut > 0 {
			var denominator float64 = math.Pow(10,float64(rec.DecimalsTokOut))
			var result = amount_out_f64/denominator
			rec.AmountOutFmt = fmt.Sprintf("%.2f",result)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_token_latest_balance(pool_aid,token_aid int64) float64 {

	var query string
	query = "SELECT "+
				"balance "+
			"FROM tok_bal "+
			"WHERE pool_aid=$1 AND tok_aid=$2 "+
			"ORDER BY time_stamp DESC,id DESC "+
			"LIMIT 1"
	row := sw.S.Db().QueryRow(query,pool_aid,token_aid)
	var err error
	var balance float64
	err=row.Scan(&balance)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0.0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_token_latest_balance(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return balance

}
func (sw *SQLStorageWrapper) Get_pool_registered_tokens(ethprice_storage *SQLStorage,pool_aid,weth_aid int64) []p.BalV2PoolToken {

	records := make([]p.BalV2PoolToken,0,32)
	var query string
	query = "WITH toks AS ("+
				"SELECT DISTINCT tok_aid FROM "+sw.S.SchemaName()+".tok_bal WHERE pool_aid=$1"+
			") "+
			"SELECT " +
				"tb.tok_aid,"+
				"ta.addr, "+
				"i.symbol,"+
				"i.name, "+
				"i.decimals "+
			"FROM toks tb "+
				"JOIN "+sw.S.SchemaName()+".addr ta ON tb.tok_aid=ta.address_id " +
				"LEFT JOIN erc20_info i ON tb.tok_aid=i.token_aid "

	rows,err := sw.S.Db().Query(query,pool_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
 	cur_ts :=  time.Now().Unix()
	defer rows.Close()
	for rows.Next() {
		var rec p.BalV2PoolToken
		var n_name,n_sym sql.NullString
		var n_decimals sql.NullInt64
		err=rows.Scan(
			&rec.Token.TokenAid,
			&rec.Token.TokenAddr,
			&n_name,
			&n_sym,
			&n_decimals,
		)
		if err!=nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if n_name.Valid { rec.Token.Name = n_name.String }
		if n_sym.Valid { rec.Token.Symbol = n_sym.String }
		swap_price_factor,found := sw.Get_latest_eth_swap_price_for_token(rec.Token.TokenAid,weth_aid,cur_ts)
		if rec.Token.TokenAid == weth_aid {
			swap_price_factor = 1.0
			found = true
		}
		fmt.Printf("Pool %v, token_aid %v, swap_price_factor %v\n",pool_aid,rec.Token.TokenAid,swap_price_factor)
		if found {
			ethusd_price,got_price := ethprice_storage.Ethprice_get_ethusd_price_closest_to_timestamp(cur_ts)
			if got_price {
				fmt.Printf("ethusd_price = %v\n",ethusd_price)
				cur_bal := sw.Get_token_latest_balance(pool_aid,rec.Token.TokenAid)
				fmt.Printf("cur_bal = %v\n",cur_bal)
				rec.Token.CurBalanceUSD = cur_bal * ethusd_price * swap_price_factor
				rec.Token.USDBalanceAvailable = true
			}
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_pool_token_balance_history(pool_aid,token_aid int64) []p.BalV2PoolTokBalanceHistory {

	records := make([]p.BalV2PoolTokBalanceHistory,0,32)

	var query string
	query = "SELECT "+
				"b.block_num,"+
				"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT ts," +
				"b.time_stamp,"+
				"b.swf_hist_id,"+
				"s.token_in_aid,"+
				"s.token_out_aid,"+
				"b.op_sign,"+
				"b.amount, "+
				"b.balance "+
			"FROM "+sw.S.SchemaName()+". tok_bal b "+
				"LEFT JOIN swf_hist h ON b.swf_hist_id=h.id "+
				"LEFT JOIN swap s ON (h.block_num=s.block_num AND h.tx_index=s.tx_index AND h.log_index=s.log_index)" +
			"WHERE b.pool_aid=$1 AND b.tok_aid=$2 " +
			"ORDER BY b.block_num,b.tx_index,b.log_index"

	rows,err := sw.S.Db().Query(query,pool_aid,token_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.BalV2PoolTokBalanceHistory
		var swaphist_id int64
		var n_in_aid,n_out_aid sql.NullInt64
		err=rows.Scan(
			&rec.BlockNum,
			&rec.TimeStamp,
			&rec.DateTime,
			&swaphist_id,
			&n_in_aid,
			&n_out_aid,
			&rec.OpSign,
			&rec.Amount,
			&rec.Balance,
		)
		if err!=nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if swaphist_id > 0 { rec.IsSwap=true }
		if n_in_aid.Valid { rec.TokenInAid = n_in_aid.Int64 }
		if n_out_aid.Valid { rec.TokenOutAid = n_out_aid.Int64 }
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_pool_swap_fee_profits(pool_aid int64,ts_ini,ts_fin int64) float64 {

	var query string
	query = "SELECT sum(swap_fee_usd) AS total "+
			"FROM "+sw.S.SchemaName()+".swf_hist "+
			"WHERE (pool_aid=$1) AND "+
				"(TO_TIMESTAMP($2)<=time_stamp) AND "+
				"(time_stamp<TO_TIMESTAMP($3))"
	row := sw.S.Db().QueryRow(query,pool_aid,ts_ini,ts_fin)
	var err error
	var total sql.NullFloat64
	err=row.Scan(&total)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0.0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_pool_swap_fee_profits(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return total.Float64
}
func (sw *SQLStorageWrapper) Get_top_profitable_pools(tf_code,ini_ts,fin_ts int64) []p.BalV2PoolProfit{

	records := make([]p.BalV2PoolProfit,0,16)
	var query string
	query = "SELECT "+
				"sa.pool_aid,"+
				"pa.addr,"+
				"p.pool_id,"+
				"amount_usd "+
				//"amount "+
			"FROM "+sw.S.SchemaName()+".swap_accum sa "+
				"JOIN "+sw.S.SchemaName()+".pool_reg p ON sa.pool_aid=p.pool_aid "+
				"JOIN "+sw.S.SchemaName()+".addr pa ON sa.pool_aid=pa.address_id "+
			"WHERE sa.tf_code=$1 "
	var many_params bool = false
	if ini_ts != 0 && fin_ts != 0 {
		query = query + "AND sa.time_stamp>=TO_TIMESTAMP($2) "
		query = query + "AND sa.time_stamp<TO_TIMESTAMP($3) "
		many_params = true
	}
	query = query + "ORDER BY amount_usd DESC"
	//query = query + "ORDER BY amount DESC"

	var err error
	var rows *sql.Rows
	if many_params {
		rows,err = sw.S.Db().Query(query,tf_code,ini_ts,fin_ts)
	} else {
		rows,err = sw.S.Db().Query(query,tf_code)
	}
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.BalV2PoolProfit
		err=rows.Scan(
			&rec.PoolAid,
			&rec.PoolAddr,
			&rec.PoolId,
			&rec.AmountUSD,
		)
		if err!=nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_top_profitable_pools_v2(tf_code,ini_ts,fin_ts int64) []p.BalV2PoolProfit{

	records := make([]p.BalV2PoolProfit,0,16)
	var query string
	query = "WITH maxes AS ("+
				"SELECT MAX(amount_usd) max_usd,pool_aid "+
				"FROM swap_accum "+
				"WHERE tf_code=$1 AND "+
						"time_stamp>=TO_TIMESTAMP($2) AND "+
						"time_stamp<TO_TIMESTAMP($3) "+
				"GROUP BY pool_aid "+
			") "+
			"SELECT "+
				"p.pool_aid,"+
				"pa.addr,"+
				"p.pool_id,"+
				"max_usd "+
			"FROM maxes " +
				"JOIN "+sw.S.SchemaName()+".pool_reg p ON maxes.pool_aid=p.pool_aid "+
				"JOIN "+sw.S.SchemaName()+".addr pa ON maxes.pool_aid=pa.address_id "+
			"ORDER BY max_usd DESC"

	var err error
	var rows *sql.Rows
	rows,err = sw.S.Db().Query(query,tf_code,ini_ts,fin_ts)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.BalV2PoolProfit
		err=rows.Scan(
			&rec.PoolAid,
			&rec.PoolAddr,
			&rec.PoolId,
			&rec.AmountUSD,
		)
		if err!=nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_pool_swap_fee_returns_by_timeframe_code(pool_aid,tf_code,offset,limit int64) []p.BalV2FeeReturns {


	records := make([]p.BalV2FeeReturns,0,16)
	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"time_stamp," +
				"amount_usd "+
			"FROM "+sw.S.SchemaName()+".swap_accum sa "+
			"WHERE sa.pool_aid=$1 AND sa.tf_code=$2 " +
			"OFFSET $3 LIMIT $4"

	d_query := "SELECT "+
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"time_stamp," +
				"amount_usd "+
			"FROM "+sw.S.SchemaName()+".swap_accum sa "+
			"WHERE sa.pool_aid="+fmt.Sprintf("%v",pool_aid)+" AND sa.tf_code="+fmt.Sprintf("%v",tf_code)+
			"OFFSET "+fmt.Sprintf("%v",offset)+" LIMIT "+fmt.Sprintf("%v",limit)
	fmt.Printf("%v",d_query)
	var err error
	var rows *sql.Rows
	rows,err = sw.S.Db().Query(query,pool_aid,tf_code,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.BalV2FeeReturns
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.FeeReturnsUSD,
		)
		if err!=nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_pool_liquidity_provider_distrib(pool_aid int64) []p.BalV2LiqProvDistrib {

	records := make([]p.BalV2LiqProvDistrib,0,16)
	var query string
	query = "SELECT "+
				"b.aid,"+
				"a.addr, "+
				"b.balance "+
			"FROM "+sw.S.SchemaName()+".bpt_bal b "+
				"LEFT JOIN "+sw.S.SchemaName()+".addr a ON b.aid=a.address_id "+
			"WHERE b.pool_aid=$1 AND b.balance>0 " +
			"ORDER BY b.balance DESC "

	rows,err := sw.S.Db().Query(query,pool_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	var total float64
	for rows.Next() {
		var rec p.BalV2LiqProvDistrib
		err=rows.Scan(
			&rec.FunderAid,
			&rec.FunderAddr,
			&rec.Balance,
		)
		if err!=nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		total = total + rec.Balance
		records = append(records,rec)
	}
	fmt.Printf("total =%v\n",total)
	for i:=0; i<len(records);i++ {
		var rec p.BalV2LiqProvDistrib
		rec = records[i]
		var percent float64
		percent = rec.Balance/total
		percent = percent * 100
		fmt.Printf("percent = %v\n",percent)
		rec.Percentage = percent
		records[i]=rec
	}
	return records
}
