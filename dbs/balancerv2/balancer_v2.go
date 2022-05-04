package balancerv2

import (
	"os"
	"fmt"
	"errors"

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
				"proto_fee_amounts "+
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
func (sw *SQLStorageWrapper) Lookup_pool_id_by_addr_id(pool_aid int64) (string,error) {

	var query string
	query = "SELECT pool_id FROM "+sw.S.SchemaName()+".pool_reg WHERE pool_aid=$1"

	row := sw.S.Db().QueryRow(query,pool_aid)
	var pool_id string
	var err error
	err=row.Scan(&pool_id);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return "",nil
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Lookup_pool_id_by_addr_id(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return pool_id,nil
}
func (sw *SQLStorageWrapper) Is_balancer_pool_address(addr string) int64 {

	var query string
	query = "SELECT pool_aid " +
			"FROM "+sw.S.SchemaName()+".pool_reg p "+
			"JOIN "+sw.S.SchemaName()+".addr a ON p.pool_aid=a.address_id "+
			"WHERE a.addr=$1"

	row := sw.S.Db().QueryRow(query,addr)
	var pool_aid int64
	var err error
	err=row.Scan(&pool_aid);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Is_balancer_pool_address(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return pool_aid

}
func (sw *SQLStorageWrapper) Is_pool_unhandled(pool_id string) bool {

	var query string
	query = "SELECT pool_aid FROM unhandled WHERE pool_id=$1"
	row := sw.S.Db().QueryRow(query,pool_id)
	var pool_aid int64
	var err error
	err=row.Scan(&pool_aid);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Is_pool_unhandled(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return true
}
func (sw *SQLStorageWrapper) Get_lowest_pool_aid() int64 {
	// Used for generating accumlated swap fees per timeframe
	var query string
	query = "SELECT pool_aid FROM pool_reg ORDER by pool_aid LIMIT 1"
	row := sw.S.Db().QueryRow(query)
	var pool_aid int64 = 0
	var err error
	err=row.Scan(&pool_aid);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_lowest_pool_aid(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return pool_aid

}
func (sw *SQLStorageWrapper) Get_greater_pool_aid(from_pool_aid int64) int64 {


	// Used for generating accumlated swap fees per timeframe
	var query string
	query = "SELECT "+
				"pool_aid "+
			"FROM pool_reg "+
			"WHERE pool_aid > $1 "+
			"ORDER by pool_aid "+
			"LIMIT 1"
	row := sw.S.Db().QueryRow(query,from_pool_aid)
	var pool_aid int64 = 0
	var err error
	err=row.Scan(&pool_aid);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_greater_pool_aid(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return pool_aid

}
func (sw *SQLStorageWrapper) Get_swaps_for_period(pool_aid,ini_ts,fin_ts int64) (string,int64,error) {

	var query string
	query = "SELECT "+
				"SUM(swap_fee) AS swap_fees, "+
				"MAX(id) AS id "+
			"FROM swf_hist "+
			"WHERE   (pool_aid=$1) AND "+
					"(TO_TIMESTAMP($2)<=time_stamp) AND "+
					"(time_stamp < TO_TIMESTAMP($3)) "

	row := sw.S.Db().QueryRow(query,pool_aid,ini_ts,fin_ts)
	var swap_fees sql.NullString
	var max_id sql.NullInt64
	var err error
	err=row.Scan(&swap_fees,&max_id);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return "",0,err
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_swaps_for_period(): %v, q=%v",err,query))
		os.Exit(1)
	} else {
		if !swap_fees.Valid {
			return "",0,errors.New("No swap fees registered")
		}
	}
	return swap_fees.String,max_id.Int64,nil
}
func (sw *SQLStorageWrapper) Get_timestamp_of_latest_swap_record(pool_aid int64) int64 {

	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts " +
			"FROM swf_hist "+
			"WHERE pool_aid=$1 "+
			"ORDER BY time_stamp DESC "+
			"LIMIT 1"

	row := sw.S.Db().QueryRow(query,pool_aid)
	var ts int64
	var err error
	err=row.Scan(&ts);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_timestamp_of_latest_swap_record(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return ts


}
func (sw *SQLStorageWrapper) Get_last_swap_accum_record(pool_aid,tf_code int64) (p.BalV2SwapAccumRec,error) {

	var query string
	query = "SELECT "+
				"id,"+
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"amount "+
			"FROM swap_accum "+
			"WHERE (pool_aid=$1) AND "+
				"(tf_code=$2) "+
			"ORDER BY time_stamp DESC "+
			"LIMIT 1"

	row := sw.S.Db().QueryRow(query,pool_aid,tf_code)
	var id,ts int64
	var swap_fees string
	var err error
	var output p.BalV2SwapAccumRec
	output.PoolAid = pool_aid
	output.TfCode = tf_code
	err=row.Scan(&id,&ts,&swap_fees);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output,err
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_last_swap_accum_record(): %v, q=%v",err,query))
		os.Exit(1)
	}
	output.Id = id
	output.Amount = swap_fees
	output.TimeStamp = ts
	return output,nil
}
func (sw *SQLStorageWrapper) Get_timestamp_of_first_swap_fee_hist_record(pool_aid int64) int64 {

	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts " +
			"FROM swf_hist "+
			"WHERE pool_aid=$1 "+
			"ORDER BY time_stamp "+
			"LIMIT 1"
	row := sw.S.Db().QueryRow(query,pool_aid)
	var ts int64
	var err error
	err=row.Scan(&ts);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_timestamp_of_first_swap_fee_hist_rec(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return ts
}
func (sw *SQLStorageWrapper) Get_latest_eth_swap_price_for_token(token_aid,weth_aid int64,ts int64) (float64,bool) {
	// returns price in ETH
	// Return values:
	//		float64		- ETH price for token
	//		bool		- true if swap record was found
	var (
		eth_out			bool = false
		eth_in			bool = false
		ts_eth_out		int64
		ts_eth_in		int64
		amount_in1		float64
		amount_out1		float64
		amount_in2		float64
		amount_out2		float64
		eth_price		float64
		block_num		int64
		tx_index		int64
		log_index		int64
	)

	var query string
	query = "SELECT " +
				"block_num,tx_index,log_index,"+
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"amount_in,"+
				"amount_out "+
			"FROM swap s "+
			"WHERE (token_in_aid=$1) AND (token_out_aid=$2) AND (time_stamp<=TO_TIMESTAMP($3)) "+
			"ORDER BY time_stamp DESC "+
			"LIMIT 1"

	q1 := strings.Replace(query,"$1",fmt.Sprintf("%v",token_aid))
	q1 = strings.Replace(q1,"$2",fmt.Sprintf("%v",weth_aid))
	q1 = strings.Replace(q1,"$3",fmt.Sprintf("%v",ts))
	Info.Printf("query1 = %v\n",q1)
	q2 := strings.Replace(query,"$1",fmt.Sprintf("%v",weth_aid))
	q2 = strings.Replace(q2,"$2",fmt.Sprintf("%v",token_aid))
	q2 = strings.Replace(q2,"$3",fmt.Sprintf("%v",ts))
	Info.Printf("query2 = %v\n",q1)
	row := sw.S.Db().QueryRow(query,token_aid,weth_aid,ts)
	var err error
	err=row.Scan(&block_num,&tx_index,&log_index,&ts_eth_out,&amount_in1,&amount_out1)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_latest_eth_swap_price_for_token(): %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		eth_out = true	// ETH is asked
	}
	sw.S.Info.Printf("Get_latest_eth_swap_price_for_token(): call1: block_num=%v,tx_index=%v,log_index=%v ts=%v: amount_in=%v amount_out=%v\n",block_num,tx_index,log_index,ts,amount_in1,amount_out1)
	row = sw.S.Db().QueryRow(query,weth_aid,token_aid,ts)
	err=row.Scan(&block_num,&tx_index,&log_index,&ts_eth_in,&amount_in2,&amount_out2);
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_latest_eth_swap_price_for_token(): %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		eth_in = true	// ETH is given
	}
	sw.S.Info.Printf("Get_latest_eth_swap_price_for_token(): call2: block_num=%v,tx_index=%v,log_index=%v : amount_in=%v amount_out=%v\n",block_num,tx_index,log_index,amount_in2,amount_out2)

	if !(eth_in || eth_out) {
		sw.S.Info.Printf("Get_latest_eth_swap_price_for_token(): no record found, returning\n")
		return 0.0,false	// no record was found, no swap exist
	}
	if eth_in && eth_out {
		ts1_diff := ts_eth_out - ts
		if ts1_diff < 0 { ts1_diff = -ts1_diff }
		ts2_diff := ts_eth_in - ts
		if ts2_diff < 0 { ts2_diff = -ts2_diff }
		if ts2_diff < ts1_diff { // eth_in
			eth_price = amount_in2/amount_out2
			sw.S.Info.Printf("Get_latest_eth_swap_price_for_token(): amount_in=%v, amount_out=%v amount_out/amount_in = %v\n",amount_in2,amount_out2,eth_price)
		} else {	// eth_out
			eth_price = amount_out1/amount_in1
			sw.S.Info.Printf("Get_latest_eth_swap_price_for_token(): amount_in=%v, amount_out=%v amount_out/amount_in = %v\n",amount_in1,amount_out1,eth_price)
		}
	} else {
		if eth_in {
			eth_price = amount_in2/amount_out2
			sw.S.Info.Printf("Get_latest_eth_swap_price_for_token(): amount_in=%v, amount_out=%v amount_out/amount_in = %v\n",amount_in2,amount_out2,eth_price)
		}
		if eth_out {
			eth_price = amount_out1/amount_in1
			sw.S.Info.Printf("Get_latest_eth_swap_price_for_token(): amount_in=%v, amount_out=%v amount_out/amount_in = %v\n",amount_in1,amount_out1,eth_price)
		}
	}
	return eth_price,true
}
func (sw *SQLStorageWrapper) Get_wrapped_eth_contract_address() string {

	var query string
	query = "SELECT weth_addr FROM config"

	var addr string
	row := sw.S.Db().QueryRow(query)
	var err error
	err=row.Scan(&addr);
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_wrapped_eth_contract_address(): %v, q=%v",err,query))
			os.Exit(1)
		}
		return ""
	}

	return addr
}
