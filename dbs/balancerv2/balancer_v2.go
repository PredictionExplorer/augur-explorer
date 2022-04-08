package balancerv2

import (
	"os"
	"fmt"

	"database/sql"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives/balancerv2"
)
type SQLStorageWrapper struct {
	S					*SQLStorage
}

func (sw *SQLStorageWrapper) Get_first_block_for_swap_history() (int64,string,bool) {

	var query string
	query = "SELECT "+
				"block_num,block_hash "+
			"FROM "+sw.S.SchemaName()+".block "+
			"ORDER BY block_num LIMIT 1"

	row := sw.S.Db().QueryRow(query)
	var err error
	var next_block_num int64
	var block_hash string
	err=row.Scan(&next_block_num,&block_hash);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,"",false
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_first_block_for_swap_history(): %v, q=%v",err,query))
		os.Exit(1)
	}

	return next_block_num,block_hash,true
}
func (sw *SQLStorageWrapper) Get_next_block_for_swap_history(block_num int64,parent_hash string) (int64,string,bool) {

	var query string
	query = "SELECT "+
				"block_num,block_hash "+
			"FROM "+sw.S.SchemaName()+".block "+
			"WHERE block_num=$1 AND block_hash=$2"

	row := sw.S.Db().QueryRow(query,block_num,parent_hash)
	var err error
	var next_block_num int64
	var block_hash string
	err=row.Scan(&next_block_num,&block_hash);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,"",false
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_next_block_for_swap_history(): %v, q=%v",err,query))
		os.Exit(1)
	}

	return next_block_num,block_hash,true
}
func (sw *SQLStorageWrapper) Get_last_block_for_swap_history() (int64,string,bool) {

	var query string
	query = "SELECT "+
				"h.block_num,"+
				"b.block_hash "+
			"FROM "+sw.S.SchemaName()+".swf_hist h "+
				"JOIN "+sw.S.SchemaName()+".block b ON h.block_num=b.block_num " +
			"ORDER BY h.block_num DESC "+
			"LIMIT 1"

	row := sw.S.Db().QueryRow(query)
	var err error
	var block_num int64
	var block_hash string
	err=row.Scan(&block_num,&block_hash);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,"",false
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_last_block_swf_hist(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return block_num,block_hash,true
}
func (sw *SQLStorageWrapper) Get_swaps_for_block(block_num int64,block_hash string) []p.BalV2Swap {

	records := make([]p.BalV2Swap,0,8)
	var query string
	query = "SELECT "+
				"s.pool_id," +
				"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT ts," +
				"s.block_num,"+
				"s.tx_index,"+
				"s.log_index,"+
				"token_in_aid,"+
				"token_out_aid,"+
				"amount_in,"+
				"amount_out "+
			"FROM "+sw.S.SchemaName()+".swap s " +
				"JOIN "+sw.S.SchemaName()+".block b ON s.block_num=b.block_num "+
			"WHERE s.block_num = $1 AND b.block_hash=$2 "+
			"ORDER BY tx_index,log_index"

	rows,err := sw.S.Db().Query(query,block_num,block_hash)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.BalV2Swap
		err=rows.Scan(
			&rec.PoolId,
			&rec.TimeStamp,
			&rec.BlockNum,
			&rec.TxIndex,
			&rec.LogIndex,
			&rec.TokenInAid,
			&rec.TokenOutAid,
			&rec.AmountIn,
			&rec.AmountOut,
		)
		if err!=nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_balance_changes_for_block(block_num int64,block_hash string) []p.BalV2PoolBalanceChanged {

	records := make([]p.BalV2PoolBalanceChanged,0,8)
	var query string
	query = "SELECT "+
				"c.pool_id," +
				"EXTRACT(EPOCH FROM c.time_stamp)::BIGINT ts," +
				"c.block_num,"+
				"c.tx_index,"+
				"c.log_index,"+
				"liqprov_aid,"+
				"tokens,"+
				"deltas, "+
				"fee_amounts "+
			"FROM "+sw.S.SchemaName()+".pool_bal c " +
				"JOIN "+sw.S.SchemaName()+".block b ON c.block_num=b.block_num "+
			"WHERE c.block_num = $1 AND b.block_hash=$2 "+
			"ORDER BY tx_index,log_index"

	rows,err := sw.S.Db().Query(query,block_num,block_hash)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.BalV2PoolBalanceChanged
		err=rows.Scan(
			&rec.PoolId,
			&rec.TimeStamp,
			&rec.BlockNum,
			&rec.TxIndex,
			&rec.LogIndex,
			&rec.LiqProvAid,
			&rec.Tokens,
			&rec.Deltas,
			&rec.ProtocolFeeAmounts,
		)
		if err!=nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_pool_fee_in_timeframe(ts_ini,ts_fin int64) (string,int64,bool) {


	var query string
	query = "SELECT "+
				"swap_fee,"+
				"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT ts " +
			"FROM "+sw.S.SchemaName()+".swap_fee s"+
			"WHERE  (TO_TIMESTAMP($1) <time_stamp) AND "+
						"time_satmp < (TO_TIMESTAMP($2) "+
			"ORDER BY time_stamp DESC "+
			"LIMIT 1"

	row := sw.S.Db().QueryRow(query,ts_ini,ts_fin)
	var err error
	var fee string
	var ts int64
	err=row.Scan(&fee,&ts);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return "",0,false
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_pool_fee_in_timeframe(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return fee,ts,true
}
func (sw *SQLStorageWrapper) Get_pool_fee_by_timestamp(p_contract_aid,p_ts,p_block_num,p_tx_index int64) (string,int64,bool) {

	var query string
	query = "SELECT "+
				"swap_fee,"+
				"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT ts, " +
				"s.block_num,"+
				"tx_index "+
			"FROM "+sw.S.SchemaName()+".swap_fee s "+
			"WHERE  "+
					"(time_stamp <= TO_TIMESTAMP($1)) AND "+
					"(contract_aid = $2) AND "+
					"("+
						// we must exclude fee record if it occurs whithin the same
						// block as our transaction, in the case transaction index is higher
						"NOT ("+
							"(s.block_num=$3) AND "+
							"(tx_index>$4)"+
						")"+
					") "+
			"ORDER BY time_stamp DESC "+
			"LIMIT 1"

	row := sw.S.Db().QueryRow(query,p_ts,p_contract_aid,p_block_num,p_tx_index)
	var err error
	var fee string
	var block_num,tx_index,ts int64
	err=row.Scan(&fee,&ts,&ts,&block_num,&tx_index);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return "",0,false
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_pool_fee_by_timestamp(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return fee,ts,true

}
func (sw *SQLStorageWrapper) Get_pool_fee_by_block_num(p_contract_aid,p_block_num,p_tx_index int64) (string,int64,bool) {

	var query string
	query = "SELECT "+
				"swap_fee,"+
				"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT ts, " +
				"s.block_num,"+
				"tx_index "+
			"FROM "+sw.S.SchemaName()+".swap_fee s "+
			"WHERE  "+
					"(block_num <= $1) AND "+
					"(contract_aid = $2) AND "+
					"("+
						// we must exclude fee record if it occurs whithin the same
						// block as our transaction, in the case transaction index is higher
						"NOT ("+
							"(s.block_num=$1) AND "+
							"(tx_index>$3)"+
						")"+
					") "+
			"ORDER BY block_num DESC "+
			"LIMIT 1"

	row := sw.S.Db().QueryRow(query,p_block_num,p_contract_aid,p_tx_index)
	var err error
	var fee string
	var block_num,tx_index,ts int64
	err=row.Scan(&fee,&ts,&block_num,&tx_index);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return "",0,false
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_pool_fee_by_block_num(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return fee,ts,true

}
func (sw *SQLStorageWrapper) Balancer_get_contract_addrs() p.BalV2ContractAddrs {

	var query string
	query = "SELECT factory_addr,vault_addr FROM "+sw.S.SchemaName()+".config"
	row := sw.S.Db().QueryRow(query)
	var factory_addr,vault_addr string
	var err error
	err=row.Scan(&factory_addr,&vault_addr);
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Balancer_get_contract_addrs(): %v, q=%v",err,query))
		os.Exit(1)
	}
	var output p.BalV2ContractAddrs
	output.FactoryAddr = factory_addr
	output.VaultAddr = vault_addr
	return output
}
func (sw *SQLStorageWrapper) Lookup_pool_address_id(pool_id string) (int64,error) {

	var query string
	query = "SELECT pool_aid FROM "+sw.S.SchemaName()+".pool_reg WHERE pool_id=$1"
	row := sw.S.Db().QueryRow(query,pool_id)
	var pool_aid int64
	var err error
	err=row.Scan(&pool_aid);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,nil
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Lookup_pool_address_id(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return pool_aid,nil
}
