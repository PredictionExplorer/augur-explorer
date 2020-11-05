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
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,caller_aid,controller_aid" +
			") VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7)"

	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		pool_aid,
		caller_aid,
		caller_aid,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v; for evt_id=%v q=%v",err,evt.EvtId,query))
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
			") VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8::DECIMAL/1e+18)"

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
		ss.Log_msg(fmt.Sprintf("DB error: %v; for evt_id=%v q=%v",err,query,evt.EvtId))
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
			") VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8::DECIMAL/1e+18)"

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
		ss.Log_msg(fmt.Sprintf("DB error: %v; for evt_id=%v q=%v",err,evt.EvtId,query))
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
			") VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9::DECIMAL/1e+18,$10::DECIMAL/1e+18)"

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
		ss.Log_msg(fmt.Sprintf("DB error: %v; for evt_id=%v q=%v",err,evt.EvtId,query))
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
func (ss *SQLStorage) Insert_set_swap_fee(f *p.SetSwapFee) {

	pool_aid := ss.Lookup_or_create_address(f.PoolAddr,f.BlockNum,f.TxId)
	var query string
	query = "INSERT INTO b_set_swap_fee(evtlog_id,block_num,tx_id,time_stamp,pool_aid,fee) "+
			"VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6::DECIMAL/1e+18)"
	_,err := ss.db.Exec(query,
		f.EvtId,
		f.BlockNum,
		f.TxId,
		f.TimeStamp,
		pool_aid,
		f.FeeStr,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v for evt_id=%v q=%v",err,f.EvtId,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_set_swap_fee(evt_id int64) {
	var query string
	query = "DELETE FROM b_set_swap_fee WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_set_controller(c *p.SetController) {

	var query string

	pool_aid := ss.Lookup_or_create_address(c.PoolAddr,c.BlockNum,c.TxId)

	query = "SELECT controller_aid FROM bpool WHERE pool_aid=$1"
	row := ss.db.QueryRow(query,pool_aid)
	var null_controller_aid sql.NullInt64
	var err error
	err=row.Scan(&null_controller_aid)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(
				fmt.Sprintf("set_controller call for unregistered pool %v (evt_id=%v)",c.PoolAddr,c.EvtId),
			)
			return
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, for evt_id=%v q=%v",err,c.EvtId,query))
			os.Exit(1)
		}
	}
	old_controller_aid := null_controller_aid.Int64
	controller_aid := ss.Lookup_or_create_address(c.ControllerAddr,c.BlockNum,c.TxId)

	query = "INSERT INTO b_set_controller(" +
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,controller_aid,old_controller_aid" +
			") VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7)"
	_,err = ss.db.Exec(query,
		c.EvtId,
		c.BlockNum,
		c.TxId,
		c.TimeStamp,
		pool_aid,
		controller_aid,
		old_controller_aid,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v for evt_id=%v q=%v",err,c.EvtId,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_set_controller(evt_id int64) {
	var query string
	query = "DELETE FROM b_set_controller WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_set_public(p *p.SetPublic) {

	var query string

	pool_aid := ss.Lookup_or_create_address(p.PoolAddr,p.BlockNum,p.TxId)

	query = "SELECT " +
				"is_public,went_public,EXTRACT(EPOCH FROM went_public_ts)::BIGINT " +
			"FROM bpool WHERE pool_aid=$1"
	row := ss.db.QueryRow(query,pool_aid)
	var null_went_public,null_went_public_ts sql.NullInt64
	var null_is_public sql.NullBool
	var err error
	err=row.Scan(&null_is_public,&null_went_public,&null_went_public_ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(
				fmt.Sprintf("set_public call for unregistered pool %v (evt_id=%v)",p.PoolAddr,p.EvtId),
			)
			return
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	old_is_public := null_is_public.Bool
	old_went_public := null_went_public.Int64
	old_went_public_ts := null_went_public.Int64

	query = "INSERT INTO b_set_public(" +
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,is_public," +
				"old_is_public,old_went_public,old_went_public_ts "+
			") VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9)"
	_,err = ss.db.Exec(query,
		p.EvtId,
		p.BlockNum,
		p.TxId,
		p.TimeStamp,
		pool_aid,
		p.Public,
		old_is_public,
		old_went_public,
		old_went_public_ts,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v for evt_id=%v q=%v",err,p.EvtId,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_set_public(evt_id int64) {
	var query string
	query = "DELETE FROM b_set_public WHERE evtlog_id=$1"
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
func (ss *SQLStorage) Insert_pool_finalize(f *p.Finalize) {

	pool_aid := ss.Lookup_or_create_address(f.PoolAddr,f.BlockNum,f.TxId)
	var query string
	query = "INSERT INTO b_finalized(evtlog_id,block_num,tx_id,time_stamp,pool_aid) "+
			"VALUES($1,$2,$3,TO_TIMESTAMP($4),$5)"
	_,err := ss.db.Exec(query,
		f.EvtId,
		f.BlockNum,
		f.TxId,
		f.TimeStamp,
		pool_aid,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v for evt_id=%v q=%v",err,f.EvtId,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_pool_finalize(evt_id int64) {
	var query string
	query = "DELETE FROM b_finalized WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_pool_bind(b *p.PoolBind) {

	pool_aid := ss.Lookup_or_create_address(b.PoolAddr,b.BlockNum,b.TxId)
	token_aid := ss.Lookup_or_create_address(b.TokenAddr,b.BlockNum,b.TxId)
	var query string
	query = "INSERT INTO b_bind(evtlog_id,block_num,tx_id,time_stamp,pool_aid,token_aid,denorm,balance) "+
			"VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8::DECIMAL/1e+18)"
	_,err := ss.db.Exec(query,
		b.EvtId,
		b.BlockNum,
		b.TxId,
		b.TimeStamp,
		pool_aid,
		token_aid,
		b.Denorm,
		b.Balance,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v for evt_id=%v q=%v",err,b.EvtId,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_pool_bind(evt_id int64) {
	var query string
	query = "DELETE FROM b_bind WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_pool_unbind(u *p.PoolUnBind) {

	pool_aid := ss.Lookup_or_create_address(u.PoolAddr,u.BlockNum,u.TxId)
	token_aid := ss.Lookup_or_create_address(u.TokenAddr,u.BlockNum,u.TxId)
	var query string

	query = "SELECT denorm,(balance*1e+18)::TEXT FROM btoken WHERE pool_aid=$1 AND token_aid=$2"
	row := ss.db.QueryRow(query,pool_aid,token_aid)
	var null_denorm sql.NullInt64
	var null_balance sql.NullString
	var err error
	err=row.Scan(&null_denorm,&null_balance)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(
				fmt.Sprintf("unbind for unregistered pool %v (evt_id=%v)",u.PoolAddr,u.EvtId),
			)
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	var saved_denorm int = 0
	var saved_balance string = "0"
	if null_denorm.Valid {
		saved_denorm = int(null_denorm.Int64)
	}
	if null_balance.Valid {
		saved_balance = null_balance.String
	}

	query = "INSERT INTO b_unbind(" +
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,token_aid,saved_denorm,saved_balance) "+
			"VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8::DECIMAL/1e+18)"
	_,err = ss.db.Exec(query,
		u.EvtId,
		u.BlockNum,
		u.TxId,
		u.TimeStamp,
		pool_aid,
		token_aid,
		saved_denorm,
		saved_balance,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v for evt_id=%v q=%v",err,u.EvtId,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_pool_unbind(evt_id int64) {
	var query string
	query = "DELETE FROM b_unbind WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_pool_rebind(r *p.PoolReBind) {

	pool_aid := ss.Lookup_or_create_address(r.PoolAddr,r.BlockNum,r.TxId)
	token_aid := ss.Lookup_or_create_address(r.TokenAddr,r.BlockNum,r.TxId)
	var query string

	query = "SELECT denorm,(balance*1e+18)::TEXT FROM btoken WHERE pool_aid=$1 AND token_aid=$2"
	row := ss.db.QueryRow(query,pool_aid,token_aid)
	var null_denorm sql.NullInt64
	var null_balance sql.NullString
	var err error
	err=row.Scan(&null_denorm,&null_balance)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(
				fmt.Sprintf("rebind for unregistered pool %v (evt_id=%v)",r.PoolAddr,r.EvtId),
			)
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	var saved_denorm int = 0
	var saved_balance string = "0"
	if null_denorm.Valid {
		saved_denorm = int(null_denorm.Int64)
	}
	if null_balance.Valid {
		saved_balance = null_balance.String
	}

	query = "INSERT INTO b_rebind(" +
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,token_aid," +
				"denorm,balance,saved_denorm,saved_balance) "+
			"VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8::DECIMAL/1e+18,$9,$10::DECIMAL/1e+18)"
	_,err = ss.db.Exec(query,
		r.EvtId,
		r.BlockNum,
		r.TxId,
		r.TimeStamp,
		pool_aid,
		token_aid,
		r.Denorm,
		r.Balance,
		saved_denorm,
		saved_balance,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v for evt_id=%v q=%v",err,r.EvtId,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_pool_rebind(evt_id int64) {
	var query string
	query = "DELETE FROM b_rebind WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_pool_gulp(g *p.PoolGulp) {

	pool_aid := ss.Lookup_or_create_address(g.PoolAddr,g.BlockNum,g.TxId)
	token_aid := ss.Lookup_or_create_address(g.TokenAddr,g.BlockNum,g.TxId)
	var query string
	query = "INSERT INTO b_gulp(evtlog_id,block_num,tx_id,time_stamp,pool_aid,token_aid,abs_balance) "+
			"VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7::DECIMAL/1e+18)"
	_,err := ss.db.Exec(query,
		g.EvtId,
		g.BlockNum,
		g.TxId,
		g.TimeStamp,
		pool_aid,
		token_aid,
		"0",	// ToDo: extract this value from ERC20 transfer event
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_pool_gulp(evt_id int64) {
	var query string
	query = "DELETE FROM b_gulpWHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
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
func (ss *SQLStorage) Get_market_balancer_pools(market_aid int64) []p.PoolInfo {

	records := make([]p.PoolInfo,0,16)
	var query string
	query = "WITH pool_ids AS (" +
				"SELECT DISTINCT pool_aid FROM af_wrapper "+
				"WHERE market_aid=$1"+
			")" +
			"SELECT " +
//				"FLOOR()EXTRACT(EPOCH FROM s.time_stamp)::BIGINT AS ts, " +
				"p.time_stamp AS ts,"+
				"p.block_num," +
				"pa.addr," +
				"p.num_swaps,"+
				"p.num_holders,"+
				"p.num_tokens," +
				"p.swap_fee," +
				"EXTRACT(EPOCH FROM p.went_puplic_ts)::BIGINT, " +
				"EXTRACT(EPOCH FROM p.finalized_ts)::BIGINT, " +
				"p.usd_liquidity " +
			"FROM pool_ids AS ids " +
				"JOIN pool AS p ON p.pool_aid=ids.pool_aid " +
				"LEFT JOIN address AS pa ON p.pool_aid=pa.address_id " +
			"ORDER BY ts"
	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.PoolInfo
		err=rows.Scan(
			&rec.CreatedTs,
			&rec.CreatedBlockNum,
			&rec.WrapperAddr,
			&rec.NumSwaps,
			&rec.NumHolders,
			&rec.NumTokens,
			&rec.SwapFee,
			&rec.WentPublicTs,
			&rec.FinalizedTs,
			&rec.UsdLiquidity,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}

	return records

}
