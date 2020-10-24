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
func (ss *SQLStorage) Delete_balancer_pool_created_evt(evt_id int64) {

	var query string
	query = "DELETE FROM bpool WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
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
func (ss *SQLStorage) Delete_balancer_join_evt(evt_id int64) {

	var query string
	query = "DELETE FROM bjoin WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
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
func (ss *SQLStorage) Delete_balancer_exit_evt(evt_id int64) {

	var query string
	query = "DELETE FROM bexit WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
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
func (ss *SQLStorage) Delete_balancer_swap_evt(evt_id int64) {

	var query string
	query = "DELETE FROM bswap WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
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
func (ss *SQLStorage) Get_pool_info(pool_aid int64) p.BalancerNewPool {

	var output p.BalancerNewPool
	var query string
	query = "SELECT " +
//				"FLOOR(EXTRACT(EPOCH FROM p.time_stamp))::BIGINT AS ts, " +
				"p.time_stamp,"+
				"p.pool_aid," +
				"p.caller_aid," +
				"pa.addr," +
				"ca.addr," +
				"p.block_num " +
			"FROM bpool AS p " +
				"LEFT JOIN address AS pa ON p.pool_aid=pa.address_id " +
				"LEFT JOIN address AS ca ON p.caller_aid=ca.address_id " +
			"WHERE p.pool_aid=$1 "

	row := ss.db.QueryRow(query,pool_aid)
	err := row.Scan(
			&output.TimeStamp,
			&output.PoolAid,
			&output.CallerAid,
			&output.PoolAddr,
			&output.CallerAddr,
			&output.BlockNum,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output
		} else {
			ss.Log_msg(fmt.Sprintf("Error Augur Foundry contract address is not set: %v",err))
			os.Exit(1)
		}
	}
	return output
}
func (ss *SQLStorage) Get_pool_swaps(pool_aid int64) []p.BalancerSwap {

	records := make([]p.BalancerSwap,0,64)
	var query string
	query = "SELECT " +
//				"FLOOR(EXTRACT(EPOCH FROM s.time_stamp))::BIGINT AS ts, " +
				"s.time_stamp AS ts,"+
				"s.block_num," +
				"s.tx_id,"+
				"ca.addr,"+
				"tia.addr," +
				"toa.addr," +
				"s.amount_in, " +
				"s.amount_out " +
			"FROM bswap AS s " +
				"LEFT JOIN address AS ca ON s.caller_aid=ca.address_id " +
				"LEFT JOIN address AS tia ON s.token_in_aid=tia.address_id " +
				"LEFT JOIN address AS toa ON s.token_out_aid=toa.address_id " +
			"WHERE s.pool_aid=$1 " +
			"ORDER BY ts"
	rows,err := ss.db.Query(query,pool_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.BalancerSwap
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.BlockNum,
			&rec.TxId,
			&rec.CallerAddr,
			&rec.TokenInAddr,
			&rec.TokenOutAddr,
			&rec.AmountInF,
			&rec.AmountOutF,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}


