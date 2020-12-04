package dbs

import (
	"fmt"
	"os"
	"strings"
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
	token1_aid := token_in_aid
	token2_aid := token_out_aid
	token1_amount := evt.AmountIn
	token2_amount := evt.AmountOut
	decimals1 := evt.DecimalsIn
	decimals2 := evt.DecimalsOut
	if token1_aid > token2_aid {
		token1_aid = token_out_aid
		token2_aid = token_in_aid
		token1_amount = evt.AmountOut
		token2_amount = evt.AmountIn
		decimals1 = evt.DecimalsOut
		decimals2 = evt.DecimalsIn
	}
	var query string
	query = "INSERT INTO bswap (" +
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,caller_aid,"+
				"token_in_aid,token_out_aid,amount_in,amount_out," +
				"token1_aid,token1_amount,token2_aid,token2_amount"+
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,"+
				"$9::DECIMAL/1e+" + fmt.Sprintf("%v",evt.DecimalsIn) + "," +
				"$10::DECIMAL/1e+" + fmt.Sprintf("%v",evt.DecimalsOut) + "," +
				"$11,$12::DECIMAL/1e+" + fmt.Sprintf("%v",decimals1) + "," +
				"$13,$14::DECIMAL/1e+" + fmt.Sprintf("%v",decimals2) + 
			")"

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
		token1_aid,token1_amount,
		token2_aid,token2_amount,
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
			") VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,TO_TIMESTAMP($9))"
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
	query = "INSERT INTO b_finalized(evtlog_id,block_num,tx_id,time_stamp,pool_aid) " +
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
			"VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7::DECIMAL/1e+18,$8::DECIMAL/1e+18)"
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
		if strings.Contains(err.Error(),"duplicate key value") {
			// its ok
			// this happens because bind() calls rebind() at the end so we get duplicated events
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v for evt_id=%v q=%v",err,b.EvtId,query))
			os.Exit(1)
		}
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

	query = "SELECT (denorm*1e+18)::TEXT,(balance*1e+18)::TEXT " +
			"FROM btoken WHERE pool_aid=$1 AND token_aid=$2"
	row := ss.db.QueryRow(query,pool_aid,token_aid)
	var null_denorm,null_balance sql.NullString
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
	var saved_denorm string = "0"
	var saved_balance string = "0"
	if null_denorm.Valid {
		saved_denorm = null_denorm.String
	}
	if null_balance.Valid {
		saved_balance = null_balance.String
	}

	query = "INSERT INTO b_unbind(" +
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,token_aid,saved_denorm,saved_balance" +
			") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7::DECIMAL/1e+18,$8::DECIMAL/1e+18)"
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

	query = "SELECT (denorm*1e+18)::TEXT,(balance*1e+18)::TEXT " +
			"FROM btoken WHERE pool_aid=$1 AND token_aid=$2"
	row := ss.db.QueryRow(query,pool_aid,token_aid)
	var null_denorm,null_balance sql.NullString
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
	var saved_denorm string = "0"
	var saved_balance string = "0"
	if null_denorm.Valid {
		saved_denorm = null_denorm.String
	}
	if null_balance.Valid {
		saved_balance = null_balance.String
	}

	query = "INSERT INTO b_rebind(" +
				"evtlog_id,block_num,tx_id,time_stamp,pool_aid,token_aid," +
				"denorm,balance,saved_denorm,saved_balance) "+
			"VALUES("+
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6," +
				"$7::DECIMAL/1e+18,$8::DECIMAL/1e+18,$9::DECIMAL/1e+18,$10::DECIMAL/1e+18"+
			")"
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
	query = "DELETE FROM b_gulp WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_pool_info(pool_aid int64) (p.BalancerNewPool,error) {

	var output p.BalancerNewPool
	var query string
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM p.time_stamp))::BIGINT AS ts, " +
				"p.time_stamp,"+
				"p.pool_aid," +
				"p.caller_aid," +
				"pa.addr," +
				"ca.addr," +
				"co.addr," +
				"p.block_num," +
				"total_weight, " +
				"num_swaps," +
				"num_holders," +
				"num_tokens," +
				"swap_fee " +
			"FROM bpool AS p " +
				"LEFT JOIN address AS pa ON p.pool_aid=pa.address_id " +
				"LEFT JOIN address AS ca ON p.caller_aid=ca.address_id " +
				"LEFT JOIN address AS co ON p.controller_aid=co.address_id " +
			"WHERE p.pool_aid=$1 "

	row := ss.db.QueryRow(query,pool_aid)
	var total_weight float64
	err := row.Scan(
			&output.TimeStamp,
			&output.CreatedDate,
			&output.PoolAid,
			&output.CallerAid,
			&output.PoolAddr,
			&output.CallerAddr,
			&output.ControllerAddr,
			&output.BlockNum,
			&total_weight,
			&output.NumSwaps,
			&output.NumHolders,
			&output.NumTokens,
			&output.SwapFee,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output,err
		} else {
			ss.Log_msg(fmt.Sprintf("Error Augur Foundry contract address is not set: %v",err))
			os.Exit(1)
		}
	}
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM t.time_stamp))::BIGINT AS ts, " +
				"t.time_stamp,"+
				"token_aid," +
				"denorm," +
				"balance," +
				"ta.addr, " +
				"w.mkt_addr," +
				"w.outcome_idx,"+
				"inf.name,"+
				"inf.symbol "+
			"FROM btoken AS t " +
				"LEFT JOIN address AS ta ON t.token_aid=ta.address_id " +
				"LEFT JOIN LATERAL (" +
					"SELECT wrapper_aid,ma.addr AS mkt_addr,outcome_idx "+
					"FROM af_wrapper AS afw " +
						"LEFT JOIN address AS ma ON afw.market_aid=ma.address_id "+
				") AS w ON w.wrapper_aid=t.token_aid " +
				"LEFT JOIN erc20_info AS inf ON t.token_aid=inf.aid " +
			"WHERE pool_aid=$1 ORDER BY t.block_num,t.id"
	rows,err := ss.db.Query(query,pool_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	tokens := make([]p.BalancerToken,0,4)
	defer rows.Close()
	for rows.Next() {
		var rec p.BalancerToken
		var mkt_addr,name,symbol sql.NullString
		var outc sql.NullInt64
		err=rows.Scan(
			&rec.TimeStampAdded,
			&rec.DateAdded,
			&rec.TokenAid,
			&rec.Denorm,
			&rec.Balance,
			&rec.TokenAddr,
			&mkt_addr,
			&outc,
			&name,
			&symbol,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.Weight = 100*rec.Denorm/total_weight
		if mkt_addr.Valid {
			rec.WrappingContract.MktAddr = mkt_addr.String
		}
		if outc.Valid {
			rec.WrappingContract.OutcomeIdx = int(outc.Int64)
		}
		if name.Valid {
			rec.WrappingContract.Name = name.String
		}
		if symbol.Valid {
			rec.WrappingContract.Symbol = symbol.String
		}
		tokens = append(tokens,rec)
	}
	output.NumAugurTokens,_ = ss.Get_pool_augur_tokens(pool_aid)
	output.Tokens = tokens
	return output,nil
}
func (ss *SQLStorage) Get_pool_swaps(pool_aid int64,offset int,limit int) []p.BalancerSwap {

	records := make([]p.BalancerSwap,0,64)
	var query string
	query = "SELECT " +
				"s.id,"+
				"FLOOR(EXTRACT(EPOCH FROM s.time_stamp))::BIGINT AS ts, " +
				"s.time_stamp as datetime,"+
				"s.block_num," +
				"s.tx_id,"+
				"ca.addr,"+
				"tia.addr," +
				"toa.addr," +
				"s.token_in_aid," +
				"s.token_out_aid," +
				"e_in.symbol,"+
				"e_out.symbol," +
				"s.amount_in, " +
				"s.amount_out " +
			"FROM bswap AS s " +
				"LEFT JOIN address AS ca ON s.caller_aid=ca.address_id " +
				"LEFT JOIN address AS tia ON s.token_in_aid=tia.address_id " +
				"LEFT JOIN address AS toa ON s.token_out_aid=toa.address_id " +
				"LEFT JOIN erc20_info AS e_in ON s.token_in_aid=e_in.aid " +
				"LEFT JOIN erc20_info AS e_out ON s.token_out_aid=e_out.aid " +
			"WHERE s.pool_aid=$1 " +
			"ORDER BY ts DESC OFFSET $2 LIMIT $3"
	rows,err := ss.db.Query(query,pool_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.BalancerSwap
		var symbol_in,symbol_out sql.NullString
		err=rows.Scan(
			&rec.Id,
			&rec.TimeStamp,
			&rec.Date,
			&rec.BlockNum,
			&rec.TxId,
			&rec.CallerAddr,
			&rec.TokenInAddr,
			&rec.TokenOutAddr,
			&rec.TokenInAid,
			&rec.TokenOutAid,
			&symbol_in,
			&symbol_out,
			&rec.AmountInF,
			&rec.AmountOutF,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.SymbolIn = symbol_in.String
		rec.SymbolOut = symbol_out.String
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_market_balancer_pools(market_aid int64) []p.PoolInfo {

	records := make([]p.PoolInfo,0,16)
	var query string
	query = "WITH pool_ids AS (" +
				"SELECT DISTINCT t.pool_aid FROM af_wrapper AS w "+
					"JOIN btoken AS t on w.wrapper_aid=t.token_aid " +
				"WHERE w.market_aid=$1"+
			")" +
			"SELECT " +
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS ts, " +
//				"p.time_stamp AS ts,"+
				"p.block_num," +
				"pa.addr," +
				"p.num_swaps,"+
				"p.num_holders,"+
				"p.num_tokens," +
				"p.swap_fee," +
				"p.is_public," +
				"(p.was_finalized>0),"+
				"EXTRACT(EPOCH FROM p.went_public_ts)::BIGINT, " +
				"EXTRACT(EPOCH FROM p.finalized_ts)::BIGINT, " +
				"p.usd_liquidity " +
			"FROM pool_ids AS ids " +
				"JOIN bpool AS p ON p.pool_aid=ids.pool_aid " +
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
			&rec.PoolAddr,
			&rec.NumSwaps,
			&rec.NumHolders,
			&rec.NumTokens,
			&rec.SwapFee,
			&rec.IsPublic,
			&rec.WasFinalized,
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
func (ss *SQLStorage) Get_last_evtlog_id() (int64,error) {

	var query string
	query = "SELECT id FROM evt_log ORDER BY id DESC LIMIT 1"
	res := ss.db.QueryRow(query)
	var null_id sql.NullInt64
	err := res.Scan(&null_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return null_id.Int64,nil
}
func (ss *SQLStorage) Is_address_balancer_pool(addr string) bool {


	var query string
	query = "SELECT p.pool_aid FROM bpool AS p " +
				"JOIN address AS a ON p.pool_aid=a.address_id " +
			"WHERE a.addr=$1"

	res := ss.db.QueryRow(query,addr)
	var null_id sql.NullInt64
	err := res.Scan(&null_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return true
}
func (ss *SQLStorage) Get_balancer_volume(market_aid int64,outc int,init_ts,fin_ts,interval int) []p.TradingVolume {

	records := make([]p.TradingVolume,0,64)
	var query string
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM bs.time_stamp))::BIGINT AS ts " +
			"FROM bswap bs " +
			"JOIN af_wrapper w ON (bs.token_in_aid=w.wrapper_aid) OR (bs.token_out_aid=w.wrapper_aid)"+
			"WHERE w.market_aid=$1 AND w.outcome_idx=$2 " +
			"ORDER BY bs.time_stamp ASC LIMIT 1"

	res := ss.db.QueryRow(query,market_aid,outc)
	var null_ts sql.NullInt64
	err := res.Scan(&null_ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	} else {
		if null_ts.Valid {
			if init_ts < int(null_ts.Int64) {
				init_ts = int(null_ts.Int64)
			}
		}
	}
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM bs.time_stamp))::BIGINT AS ts " +
			"FROM bswap bs " +
			"JOIN af_wrapper w ON (bs.token_in_aid=w.wrapper_aid) OR (bs.token_out_aid=w.wrapper_aid)"+
			"WHERE w.market_aid=$1 AND w.outcome_idx=$2 " +
			"ORDER BY bs.time_stamp DESC LIMIT 1"

	res = ss.db.QueryRow(query,market_aid,outc)
	err = res.Scan(&null_ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	} else {
		if null_ts.Valid {
			if fin_ts > int(null_ts.Int64) {
				fin_ts = int(null_ts.Int64)
			}
		}
	}

	query = 
			"WITH periods AS (" +
				"SELECT * FROM (" +
					"SELECT " +
						"generate_series AS start_ts,"+
						"TO_TIMESTAMP(EXTRACT(EPOCH FROM generate_series) + $3) AS end_ts "+
					"FROM (" +
						"SELECT * " +
							"FROM generate_series(" +
								"TO_TIMESTAMP($1)," +
								"TO_TIMESTAMP($2)," +
								"TO_TIMESTAMP($3)-TO_TIMESTAMP(0)) " +
					") AS i" +
				") AS data " +
			") " +
			"SELECT " +
				"COALESCE(COUNT(sw.id),0) as num_rows, " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				"SUM(amount) AS volume " +
			"FROM periods AS p " +
				"LEFT JOIN (" +
					"(" +
						"SELECT bs.id,amount_in AS amount,bs.time_stamp AS ts " +
						"FROM bswap bs " +
						"JOIN af_wrapper w ON bs.token_in_aid=w.wrapper_aid " +
						"WHERE w.market_aid=$4 AND w.outcome_idx=$5 " +
					") UNION ALL (" +
						"SELECT bs.id,amount_out AS amount,bs.time_stamp AS ts " +
						"FROM bswap bs " +
						"JOIN af_wrapper w ON bs.token_out_aid=w.wrapper_aid " +
						"WHERE w.market_aid=$4 AND w.outcome_idx=$5 " +
					")" +
				") AS sw ON " +
					"p.start_ts <= sw.ts AND "+
					"sw.ts < p.end_ts " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"

	rows,err := ss.db.Query(query,init_ts,fin_ts,interval,market_aid,outc)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.TradingVolume
		var null_amount sql.NullFloat64
		var null_ts,null_num_rows sql.NullInt64
		rows.Scan(&null_num_rows,&null_ts,&null_amount)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		if null_num_rows.Valid {
			rec.NumRecords = null_num_rows.Int64
		}
		if null_amount.Valid {
			rec.Amount= null_amount.Float64
		}
		if null_ts.Valid {
			rec.TimeStamp= null_ts.Int64
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_wrapped_transfers_volume(wrapper_aid int64,init_ts,fin_ts,interval int) []p.TradingVolume {

	records := make([]p.TradingVolume,0,64)
	var query string
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM b.ts))::BIGINT AS ts " +
			"FROM wstok_transf t " +
			"JOIN block AS b on t.block_num=b.block_num " +
			"WHERE t.wrapper_aid=$1 " +
			"ORDER BY ts ASC LIMIT 1"

	res := ss.db.QueryRow(query,wrapper_aid)
	var null_ts sql.NullInt64
	err := res.Scan(&null_ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	} else {
		if null_ts.Valid {
			if init_ts < int(null_ts.Int64) {
				init_ts = int(null_ts.Int64)
			}
		}
	}
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM b.ts))::BIGINT AS ts " +
			"FROM wstok_transf t " +
			"JOIN block AS b on t.block_num=b.block_num " +
			"WHERE t.wrapper_aid=$1 " +
			"ORDER BY ts DESC LIMIT 1"

	res = ss.db.QueryRow(query,wrapper_aid)
	err = res.Scan(&null_ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	} else {
		if null_ts.Valid {
			if fin_ts > int(null_ts.Int64) {
				fin_ts = int(null_ts.Int64)
			}
		}
	}

	query = 
			"WITH periods AS (" +
				"SELECT * FROM (" +
					"SELECT " +
						"generate_series AS start_ts,"+
						"TO_TIMESTAMP(EXTRACT(EPOCH FROM generate_series) + $3) AS end_ts "+
					"FROM (" +
						"SELECT * " +
							"FROM generate_series(" +
								"TO_TIMESTAMP($1)," +
								"TO_TIMESTAMP($2)," +
								"TO_TIMESTAMP($3)-TO_TIMESTAMP(0)) " +
					") AS i" +
				") AS data " +
			") " +
			"SELECT " +
				"COALESCE(COUNT(d.id),0) as num_rows, " +
				"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
				"SUM(amount) AS volume " +
			"FROM periods AS p " +
				"LEFT JOIN (" +
					"SELECT " +
						"tr.id," +
						"b.ts, " +
						"amount " +
					"FROM wstok_transf AS tr " +
					"JOIN block AS b ON tr.block_num=b.block_num " +
					"WHERE tr.wrapper_aid=$4" +
				") AS d ON (" +
					"p.start_ts <= d.ts AND " +
					"d.ts < p.end_ts " +
				") " +
			"GROUP BY start_ts " +
			"ORDER BY start_ts"

	rows,err := ss.db.Query(query,init_ts,fin_ts,interval,wrapper_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.TradingVolume
		var null_amount sql.NullFloat64
		var null_ts,null_num_rows sql.NullInt64
		rows.Scan(&null_num_rows,&null_ts,&null_amount)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v",err))
			os.Exit(1)
		}
		if null_num_rows.Valid {
			rec.NumRecords = null_num_rows.Int64
		}
		if null_amount.Valid {
			rec.Amount= null_amount.Float64
		}
		if null_ts.Valid {
			rec.TimeStamp= null_ts.Int64
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_balancer_token_prices(pool_aid,token1_aid,token2_aid int64,init_ts,fin_ts int) []p.BSwapPrice {

	var query string
	query = "SELECT * FROM ("+
				"(" +
					"SELECT id,EXTRACT(EPOCH FROM time_stamp)::BIGINT ts,time_stamp,amount_in/amount_out AS price " +
					"FROM bswap WHERE token_in_aid=$1 AND token_out_aid=$2 AND pool_aid=$3 " +
						"AND time_stamp >= TO_TIMESTAMP($4) AND time_stamp < TO_TIMESTAMP($5) " +
				") UNION ALL (" +
					"SELECT id,EXTRACT(EPOCH FROM time_stamp)::BIGINT ts,time_stamp,amount_out/amount_in AS price " +
					"FROM bswap WHERE token_out_aid=$1 AND token_in_aid=$2 AND pool_aid=$3 " +
						"AND time_stamp >= TO_TIMESTAMP($4) AND time_stamp < TO_TIMESTAMP($5) " +
				")" +
			") AS s " +
				"ORDER BY time_stamp" 

	d_query := fmt.Sprintf("SELECT * FROM ("+
				"(" +
					"SELECT id,EXTRACT(EPOCH FROM time_stamp)::BIGINT ts,time_stamp,amount_in/amount_out AS price " +
					"FROM bswap WHERE token_in_aid=%v AND token_out_aid=%v AND pool_aid=%v " +
						"AND time_stamp >= TO_TIMESTAMP(%v) AND time_stamp < TO_TIMESTAMP(%v) " +
				") UNION ALL (" +
					"SELECT id,EXTRACT(EPOCH FROM time_stamp)::BIGINT ts,time_stamp,amount_out/amount_in AS price " +
					"FROM bswap WHERE token_out_aid=%v AND token_in_aid=%v AND pool_aid=%v " +
						"AND time_stamp >= TO_TIMESTAMP(%v) AND time_stamp < TO_TIMESTAMP(%v) " +
				")" +
			") AS s" +
			"ORDER BY time_stamp",
			token1_aid,token2_aid,pool_aid,init_ts,fin_ts,token1_aid,token2_aid,pool_aid,init_ts,fin_ts)
	ss.Info.Printf("dquery= %v\n",d_query)

	records := make([]p.BSwapPrice,0,128)

	rows,err := ss.db.Query(query,token1_aid,token2_aid,pool_aid,init_ts,fin_ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.BSwapPrice
		err=rows.Scan(
			&rec.Id,
			&rec.TimeStamp,
			&rec.Date,
			&rec.Price,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	ss.Info.Printf("returning %v records for price chart\n",len(records))
	return records

}
func (ss *SQLStorage) Get_bpool_token_info(pool_aid,token_aid int64) (p.BalancerToken,error) {

	var query string
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM t.time_stamp))::BIGINT AS ts, " +
				"t.time_stamp,"+
				"token_aid," +
				"p.total_weight,"+
				"denorm," +
				"balance," +
				"ta.addr, " +
				"w.mkt_addr," +
				"w.outcome_idx,"+
				"inf.name,"+
				"inf.symbol "+
			"FROM btoken AS t " +
				"LEFT JOIN address AS ta ON t.token_aid=ta.address_id " +
				"LEFT JOIN LATERAL (" +
					"SELECT wrapper_aid,ma.addr AS mkt_addr,outcome_idx "+
					"FROM af_wrapper AS afw " +
						"LEFT JOIN address AS ma ON afw.market_aid=ma.address_id "+
				") AS w ON w.wrapper_aid=t.token_aid " +
				"LEFT JOIN erc20_info AS inf ON t.token_aid=inf.aid " +
				"JOIN bpool AS p on t.pool_aid=p.pool_aid " +
			"WHERE t.pool_aid=$1 and t.token_aid=$2"

	var rec p.BalancerToken
	var mkt_addr,name,symbol sql.NullString
	var outc sql.NullInt64
	var total_weight sql.NullFloat64
	res := ss.db.QueryRow(query,pool_aid,token_aid)
	err := res.Scan(
		&rec.TimeStampAdded,
		&rec.DateAdded,
		&rec.TokenAid,
		&total_weight,
		&rec.Denorm,
		&rec.Balance,
		&rec.TokenAddr,
		&mkt_addr,
		&outc,
		&name,
		&symbol,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return rec,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	rec.Weight = 100*rec.Denorm/total_weight.Float64
	if mkt_addr.Valid {
		rec.WrappingContract.MktAddr = mkt_addr.String
	}
	if outc.Valid {
		rec.WrappingContract.OutcomeIdx = int(outc.Int64)
	}
	if name.Valid {
		rec.WrappingContract.Name = name.String
	}
	if symbol.Valid {
		rec.WrappingContract.Symbol = symbol.String
	}
	return rec,nil
}
func (ss *SQLStorage) Get_balancer_swap_by_id(id int64) (p.BalancerSwap,error) {

	var rec p.BalancerSwap
	var query string

	query = "SELECT " +
				"s.Id," +
				"s.pool_aid,"+
				"pa.addr, " +
				"FLOOR(EXTRACT(EPOCH FROM s.time_stamp))::BIGINT AS ts, " +
				"s.time_stamp as datetime,"+
				"s.block_num," +
				"s.tx_id,"+
				"ca.addr,"+
				"tia.addr," +
				"toa.addr," +
				"s.token_in_aid," +
				"s.token_out_aid," +
				"e_in.Symbol,"+
				"e_out.Symbol," +
				"s.amount_in, " +
				"s.amount_out " +
			"FROM bswap AS s " +
				"LEFT JOIN address AS pa ON s.pool_aid=pa.address_id " +
				"LEFT JOIN address AS ca ON s.caller_aid=ca.address_id " +
				"LEFT JOIN address AS tia ON s.token_in_aid=tia.address_id " +
				"LEFT JOIN address AS toa ON s.token_out_aid=toa.address_id " +
				"LEFT JOIN erc20_info e_in ON s.token_in_aid=e_in.aid " +
				"LEFT JOIN erc20_info e_out ON s.token_out_aid=e_out.aid " +
			"WHERE s.id=$1 "

	res := ss.db.QueryRow(query,id)
	var symbol_in,symbol_out sql.NullString
	err := res.Scan(
		&rec.Id,
		&rec.PoolAid,
		&rec.PoolAddr,
		&rec.TimeStamp,
		&rec.Date,
		&rec.BlockNum,
		&rec.TxId,
		&rec.CallerAddr,
		&rec.TokenInAddr,
		&rec.TokenOutAddr,
		&rec.TokenInAid,
		&rec.TokenOutAid,
		&symbol_in,
		&symbol_out,
		&rec.AmountInF,
		&rec.AmountOutF,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return rec,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	rec.SymbolIn = symbol_in.String
	rec.SymbolOut = symbol_out.String

	return rec,nil
}
func (ss *SQLStorage) Get_balancer_pool_tokens_for_slippage(pool_aid int64) []p.TokenSlippage {

	var query string
	records := make([]p.TokenSlippage,0,8)
	query = "SELECT " +
				"inf1.decimals AS decimals1," +
				"inf2.decimals AS decimals2," +
				"sp.num_swaps, " +
				"pa.addr, " +
				"t1a.addr," +
				"t2a.addr, " +
				"inf1.symbol,"+
				"inf2.symbol " +
			"FROM b_swaps_per_pair AS sp " +
			"JOIN erc20_info AS inf1 ON sp.token1_aid=inf1.aid " +
			"JOIN erc20_info AS inf2 ON sp.token2_aid=inf2.aid " +
			"JOIN address AS pa ON sp.pool_aid=pa.address_id " +
			"JOIN address AS t1a ON sp.token1_aid=t1a.address_id " +
			"JOIN address AS t2a ON sp.token2_aid=t2a.address_id " +
			"WHERE sp.pool_aid=$1 " +
			"ORDER BY sp.num_swaps DESC " +
			"LIMIT 3"
	rows,err := ss.db.Query(query,pool_aid)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return records
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.TokenSlippage
		err=rows.Scan(
			&rec.Decimals1,
			&rec.Decimals2,
			&rec.NumSwaps,
			&rec.PoolAddr,
			&rec.Token1Addr,
			&rec.Token2Addr,
			&rec.Token1Symbol,
			&rec.Token2Symbol,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_pool_augur_tokens(pool_aid int64) (int64,error) {
	// returns number of Augur-related tokens
	var query string
	query = "SELECT " +
				"count(*) AS num_toks " +
			"FROM btoken AS t " +
				"JOIN af_wrapper aw ON t.token_aid=aw.wrapper_aid " +
			"WHERE t.pool_aid=$1"

	row := ss.db.QueryRow(query,pool_aid)
	var null_count sql.NullInt64	
	err := row.Scan(&null_count)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,err
		} else {
			ss.Log_msg(fmt.Sprintf("Error : %v",err))
			os.Exit(1)
		}
	}
	return null_count.Int64,nil
}
