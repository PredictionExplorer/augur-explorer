
package balancerv2

import (
	"os"
	"fmt"

	"database/sql"
	//. "github.com/PredictionExplorer/augur-explorer/dbs"
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
func (sw *SQLStorageWrapper) Get_pool_registered_tokens(pool_aid int64) []p.BalV2PoolToken {

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
	query = "SELECT sum(swap_fee) AS total "+
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
				//"amountUSD "+
				"amount "+
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
	//query = query + "ORDER BY amountUSD DESC"
	query = query + "ORDER BY amount DESC"

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
