
package balancerv2

import (
	"os"
	"fmt"

	"database/sql"
	//. "github.com/PredictionExplorer/augur-explorer/dbs"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives/balancerv2"
)
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
