package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_balancer_status() p.BalancerStatus {

	var output p.BalancerStatus
	var null_id sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM balancer_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO balancer_status DEFAULT VALUES"
				_,err := ss.db.Exec(query)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
				os.Exit(1)
			}
		} else {
			break
		}
	}
	if null_id.Valid {
		output.LastEvtId = null_id.Int64
	}
	return output
}
func (ss *SQLStorage) Update_balancer_status(status *p.BalancerStatus) {

	var query string
	query = "UPDATE balancer_status SET last_evt_id = $1"

	_,err := ss.db.Exec(query,status.LastEvtId)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_balancer_pool_created_evt(evt *p.BalancerNewPool) {

	pool_aid := ss.Lookup_or_create_address(evt.PoolAddr,evt.BlockNum,evt.TxId)
	caller_aid := ss.Lookup_or_create_address(evt.CallerAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO bpool (" +
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,caller_aid" +
			") VALUES ($1,$2,$3,$4,$5,$6)"

	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		pool_aid,
		caller_aid,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_balancer_pool_join_evt(evt *p.BalancerJoin) {

	pool_aid := ss.Lookup_or_create_address(evt.PoolAddr,evt.BlockNum,evt.TxId)
	caller_aid := ss.Lookup_or_create_address(evt.CallerAddr,evt.BlockNum,evt.TxId)
	token_aid := ss.Lookup_or_create_address(evt.TokenInAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO bjoin (" +
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,caller_aid,token_aid,amount_in" +
			") VALUES ($1,$2,$3,$4,$5,$6,$7,$8::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		pool_aid,
		caller_aid,
		token_aid,
		evt.AmountIn,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_balancer_pool_exit_evt(evt *p.BalancerExit) {

	pool_aid := ss.Lookup_or_create_address(evt.PoolAddr,evt.BlockNum,evt.TxId)
	caller_aid := ss.Lookup_or_create_address(evt.CallerAddr,evt.BlockNum,evt.TxId)
	token_aid := ss.Lookup_or_create_address(evt.TokenOutAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO bexit (" +
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,caller_aid,token_aid,amount_out" +
			") VALUES ($1,$2,$3,$4,$5,$6,$7,$8::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		pool_aid,
		caller_aid,
		token_aid,
		evt.AmountOut,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_balancer_swap_evt(evt *p.BalancerSwap) {

	pool_aid := ss.Lookup_or_create_address(evt.PoolAddr,evt.BlockNum,evt.TxId)
	caller_aid := ss.Lookup_or_create_address(evt.CallerAddr,evt.BlockNum,evt.TxId)
	token_in_aid := ss.Lookup_or_create_address(evt.TokenInAddr,evt.BlockNum,evt.TxId)
	token_out_aid := ss.Lookup_or_create_address(evt.TokenOutAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO bswap (" +
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,caller_aid,"+
				"token_in_aid,token_out_aid,amount_in,amount_out" +
			") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9::DECIMAL/1e+18,$10::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		pool_aid,
		caller_aid,
		token_in_aid,
		token_out_aid,
		evt.AmountIn,
		evt.AmountOut,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_balancer_contracts() (string,string,string) {

	var query string
	query="SELECT pool_token,factory,xchg_proxy FROM balancer_contracts";
	row := ss.db.QueryRow(query)
	var null_pool,null_factory,null_xchg sql.NullString
	var err error
	err=row.Scan(&null_pool,&null_factory,&null_xchg);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Balancer contracts are not defined in 'balancer_contracts' table"))
			os.Exit(1)
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return null_pool.String,null_factory.String,null_xchg.String
}
