
package balancerv2

import (
	"os"
	"fmt"
	"time"

	"database/sql"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives/balancerv2"
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
				"ta.addr "+
			"FROM toks tb "+
				"JOIN "+sw.S.SchemaName()+".addr ta ON tb.tok_aid=ta.address_id"

	rows,err := sw.S.Db().Query(query,pool_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
 	cur_ts :=  time.Now().Unix()
	defer rows.Close()
	for rows.Next() {
		var rec p.BalV2PoolToken
		err=rows.Scan(
			&rec.Token.TokenAid,
			&rec.Token.TokenAddr,
		)
		if err!=nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
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
				"amount, "+
				"balance "+
			"FROM "+sw.S.SchemaName()+". tok_bal b "+
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
		err=rows.Scan(
			&rec.BlockNum,
			&rec.TimeStamp,
			&rec.DateTime,
			&swaphist_id,
			&rec.Amount,
			&rec.Balance,
		)
		if err!=nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if swaphist_id > 0 { rec.IsSwap=true }
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
			&rec.PoolId,
			&rec.PoolAddr,
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
			&rec.PoolId,
			&rec.PoolAddr,
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
