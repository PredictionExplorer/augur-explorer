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
		sw.S.Log_msg(fmt.Sprintf("Error in Get_next_block(): %v, q=%v",err,query))
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
		sw.S.Log_msg(fmt.Sprintf("Error in Get_next_block(): %v, q=%v",err,query))
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
				"JOIN block b ON h.block_num=b.block_num " +
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
func (sw *SQLStorageWrapper) Insert_swap_fee_history(rec *p.BalV2SwapHist) {

	var query string
	query = "INSERT INTO "+sw.S.SchemaName()+".swf_hist("+
				"block_num,time_stamp,tx_index,log_index,contract_aid,"+
				"pool_id,swap_fee,protocol_fee,accum_swap_fee,accum_proto_fee"+
			") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10)"
	_,err := sw.S.Db().Exec(query,
		rec.BlockNum,
		rec.TimeStamp,
		rec.TxIndex,
		rec.LogIndex,
		rec.ContractAid,
		rec.PoolId,
		rec.SwapFee,
		rec.ProtocolFee,
		rec.AccumSwapFee,
		rec.AccumProtoFee,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: can't insert into pool_created table: %v\n",err))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Get_swaps_for_block(block_num int64,block_hash string) []p.BalV2Swap {

	records := make([]p.BalV2Swap,0,8)
	var query string
	query = "SELECT "+
				"pool_id," +
				"EXTRACT(EPOCH FROM bs.time_stamp)::BIGINT ts," +
				"tx_index,"+
				"log_index,"+
				"token_in_aid,"+
				"token_out_aid,"+
				"amount_in,"+
				"amount_out"+
			"FROM "+sw.S.SchemaName()+".swap s " +
				"JOIN block b ON s.block_num=b.block_num "+
			"WHERE block_num = $1 AND b.block_hash=$2 "+
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
func (sw *SQLStorageWrapper) Get_pool_fee_in_timeframe(ts_ini,ts_fin int64) (string,int64,bool) {


	var query string
	query = "SELECT "+
				"swap_fee,"+
				"EXTRACT(EPOCH FROM bs.time_stamp)::BIGINT ts " +
			"FROM "+sw.S.SchemaName()+".swap_fee "+
			"WHERE  (TO_TIMESTAMP($1) <time_stamp) AND "+
						"time_satmp < (TO_TIMESTAMP($2) "+
			"ORDER BY time_stamp DESC "+
			"IMIT 1"

	row := sw.S.Db().QueryRow(query,ts_ini,ts_fin)
	var err error
	var fee string
	var ts int64
	err=row.Scan(&fee,&ts);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return "",0,false
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_last_block_swf_hist(): %v, q=%v",err,query))
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
