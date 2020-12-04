package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_uniswap_status() p.UniswapStatus {

	var output p.UniswapStatus
	var null_id sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM uniswap_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO uniswap_status DEFAULT VALUES"
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
func (ss *SQLStorage) Update_uniswap_status(status *p.UniswapStatus) {

	var query string
	query = "UPDATE uniswap_status SET last_evt_id = $1"

	_,err := ss.db.Exec(query,status.LastEvtId)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_uniswap_pair_created_evt(bci *p.BasicChainInfo,evt *p.UPairCreated) {

	pair_aid := ss.Lookup_or_create_address(evt.Pair.String(),bci.BlockNum,bci.TxId)
	token0_aid := ss.Lookup_or_create_address(evt.Token0.String(),bci.BlockNum,bci.TxId)
	token1_aid := ss.Lookup_or_create_address(evt.Token1.String(),bci.BlockNum,bci.TxId)
	pair_seq := evt.PairSeq.Int64()
	var query string
	query = "INSERT INTO upair(" +
				"evtlog_id,block_num,tx_id,time_stamp,pair_aid,token0_aid,token1_aid,pair_seq" +
			") VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8)"

	_,err := ss.db.Exec(query,
		bci.EvtId,
		bci.BlockNum,
		bci.TxId,
		bci.TimeStamp,
		pair_aid,
		token0_aid,
		token1_aid,
		pair_seq,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v; for evt_id=%v q=%v",err,bci.EvtId,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_uniswap_pair_created_evt(evt_id int64) {

	var query string
	query = "DELETE FROM upair WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Pair_exists(pair_addr string) bool {

	var query string
	query = "SELECT id FROM upair AS p,address AS a " +
			"WHERE a.addr=$1 AND a.address_id=p.pair_aid"
	row := ss.db.QueryRow(query,pair_addr)
	var id int64
	var err error
	err=row.Scan(&id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return true
}
func (ss *SQLStorage) Get_pair_tokens(pair_addr string) (p.UPairTokens,error) {

	var query string
	query = "SELECT t0a.addr,t1a.addr,inf0.decimals,inf1.decimals " +
			"FROM upair AS p " +
			"JOIN address AS a ON p.pair_aid=a.address_id " +
			"JOIN address AS t0a ON p.token0_aid=t0a.address_id " +
			"JOIN address AS t1a ON p.token1_aid=t1a.address_id " +
			"LEFT JOIN erc20_info AS inf0 ON inf0.aid=p.token0_aid " +
			"LEFT JOIN erc20_info AS inf1 ON inf1.aid=p.token1_aid " +
			"WHERE a.addr=$1 AND a.address_id=p.pair_aid"
	row := ss.db.QueryRow(query,pair_addr)
	var err error
	var ptoks p.UPairTokens
	var decimals0,decimals1 sql.NullInt64
	var t0addr,t1addr string
	err=row.Scan(&t0addr,&t1addr,&decimals0,&decimals1)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return ptoks,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	ptoks.Token0Addr = common.HexToAddress(t0addr)
	ptoks.Token1Addr = common.HexToAddress(t1addr)
	return ptoks,nil
}
func (ss *SQLStorage) Insert_uniswap_pair_swap_evt(pair *common.Address,bci *p.BasicChainInfo,evt *p.UPairSwap) {

	// sender and recipient do not contain meaningful data, it is the contract address
	sender_aid := ss.Lookup_or_create_address(evt.Sender.String(),bci.BlockNum,bci.TxId)
	recipient_aid := ss.Lookup_or_create_address(evt.To.String(),bci.BlockNum,bci.TxId)

	pair_aid := ss.Lookup_or_create_address(pair.String(),bci.BlockNum,bci.TxId)

	amount0_in := evt.Amount0In.String()
	amount1_in := evt.Amount1In.String()
	amount0_out := evt.Amount0Out.String()
	amount1_out := evt.Amount1Out.String()

	var query string
	query = "INSERT INTO uswap1(" +
				"evtlog_id,block_num,tx_id,time_stamp,"+
				"pair_aid,sender_aid,recipient_aid," +
				"amount0_in,amount1_in,amount0_out,amount1_out"+
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),"+
				"$5,$6,$7,"+
				"$8::DECIMAL/1e+"+fmt.Sprintf("%v",evt.Decimals0)+"," +
				"$9::DECIMAL/1e+"+fmt.Sprintf("%v",evt.Decimals1)+"," +
				"$10::DECIMAL/1e+"+fmt.Sprintf("%v",evt.Decimals0)+"," +
				"$11::DECIMAL/1e+"+fmt.Sprintf("%v",evt.Decimals1)+"" +
			")"

	_,err := ss.db.Exec(query,
		bci.EvtId,bci.BlockNum,bci.TxId,bci.TimeStamp,
		pair_aid,sender_aid,recipient_aid,
		amount0_in,amount1_in,amount0_out,amount1_out,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v; for evt_id=%v q=%v",err,bci.EvtId,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_uniswap_pair_swap_evt(evt_id int64) {

	var query string
	query = "DELETE FROM uswap1 WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evt_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_uniswap_contracts() (string,string,string) {

	var query string
	query="SELECT uniswap_factory,uniswap_router01,uniswap_router02 "+
			"FROM uniswap_contracts";
	row := ss.db.QueryRow(query)
	var router1,router2,factory string
	var err error
	err=row.Scan(&factory,&router1,&router2)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Uniswap contracts are not defined in 'uniswap_contracts' table"))
			os.Exit(1)
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return factory,router1,router2
}
func (ss *SQLStorage) Find_uniswap_transfer_events(tx_id int64) [][]byte {

	var query string
	query = "SELECT log_rlp FROM evt_log WHERE tx_id=$1 AND topic0_sig='ddf252ad' ORDER BY id"

	rows,err := ss.db.Query(query,tx_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	output := make([][]byte,0,8)
	defer rows.Close()
	for rows.Next() {
		var rlp_data []byte
		err=rows.Scan(&rlp_data)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		output = append(output,rlp_data)
	}
	return output
}
func (ss *SQLStorage) Get_market_uniswap_pairs(market_aid int64) []p.MarketUPair {

	records := make([]p.MarketUPair,0,32)
	var query string
	query = "SELECT market_type,outcomes FROM market WHERE market_aid=$1"
	res := ss.db.QueryRow(query,market_aid)
	var outcomes string
	var mkt_type int64
	err := res.Scan(&mkt_type,&outcomes)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return records
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	query = "SELECT " +
				"ma.addr AS mkt_addr," +
				"w.market_aid," +
				"w.outcome_idx, " +
				"p.pair_aid," +
				"p.token0_aid," +
				"p.token1_aid," +
				"pa.addr AS pair_addr," +
				"t0a.addr AS token0_addr," +
				"t1a.addr AS token1_addr," +
				"p.total_swaps, " +
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS created_ts, "+
				"p.time_stamp," +
				"inf0.decimals," +
				"inf0.name," +
				"inf0.symbol, " +
				"inf1.decimals," +
				"inf1.name," +
				"inf1.symbol " +
			"FROM af_wrapper AS w " +
			"JOIN upair AS p ON (w.wrapper_aid = p.token0_aid) OR (w.wrapper_aid=p.token1_aid) " +
			"LEFT JOIN address AS ma ON w.market_aid=ma.address_id " +
			"LEFT JOIN address AS pa ON p.pair_aid=pa.address_id " +
			"LEFT JOIN address AS t0a ON p.token0_aid=t0a.address_id " +
			"LEFT JOIN address AS t1a ON p.token1_aid=t1a.address_id " +
			"LEFT JOIN erc20_info AS inf0 ON p.token0_aid=inf0.aid " +
			"LEFT JOIN erc20_info AS inf1 ON p.token1_aid=inf1.aid " +
			"WHERE w.market_aid=$1 "

	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.MarketUPair
		var decimals0,decimals1 sql.NullInt64
		var name0,symbol0,name1,symbol1 sql.NullString
		err=rows.Scan(
			&rec.MktAddr,
			&rec.MktAid,
			&rec.OutcomeIdx,
			&rec.PairAid,
			&rec.Token0Aid,
			&rec.Token1Aid,
			&rec.PairAddr,
			&rec.Token0Addr,
			&rec.Token1Addr,
			&rec.TotalSwaps,
			&rec.CreatedTs,
			&rec.CreatedDate,
			&decimals0,
			&name0,
			&symbol0,
			&decimals1,
			&name1,
			&symbol1,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if decimals0.Valid {rec.Token0Decimals = int(decimals0.Int64)	}
		if decimals1.Valid {rec.Token1Decimals = int(decimals1.Int64) }
		if name0.Valid { rec.Token0Name = name0.String }
		if name1.Valid { rec.Token1Name = name1.String }
		if symbol0.Valid { rec.Token0Symbol = symbol0.String }
		if symbol1.Valid { rec.Token1Symbol = symbol1.String }
		rec.Outcome = get_outcome_str(uint8(mkt_type),int(rec.OutcomeIdx),&outcomes)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_uniswap_pair_info(pair_aid int64) (p.MarketUPair,error) {

	var query string
	query = "SELECT " +
				"p.pair_aid," +
				"p.token0_aid," +
				"p.token1_aid," +
				"pa.addr AS pair_addr," +
				"t0a.addr AS token0_addr," +
				"t1a.addr AS token1_addr," +
				"p.total_swaps, " +
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS created_ts, "+
				"p.time_stamp," +
				"inf0.decimals," +
				"inf0.name," +
				"inf0.symbol, " +
				"inf1.decimals," +
				"inf1.name," +
				"inf1.symbol " +
			"FROM upair AS p " +
			"LEFT JOIN address AS pa ON p.pair_aid=pa.address_id " +
			"LEFT JOIN address AS t0a ON p.token0_aid=t0a.address_id " +
			"LEFT JOIN address AS t1a ON p.token1_aid=t1a.address_id " +
			"LEFT JOIN erc20_info AS inf0 ON p.token0_aid=inf0.aid " +
			"LEFT JOIN erc20_info AS inf1 ON p.token1_aid=inf1.aid " +
			"WHERE p.pair_aid=$1 "

	res := ss.db.QueryRow(query,pair_aid)
	var rec p.MarketUPair
	var decimals0,decimals1 sql.NullInt64
	var name0,symbol0,name1,symbol1 sql.NullString
	err := res.Scan(
		&rec.PairAid,
		&rec.Token0Aid,
		&rec.Token1Aid,
		&rec.PairAddr,
		&rec.Token0Addr,
		&rec.Token1Addr,
		&rec.TotalSwaps,
		&rec.CreatedTs,
		&rec.CreatedDate,
		&decimals0,
		&name0,
		&symbol0,
		&decimals1,
		&name1,
		&symbol1,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return rec,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	if decimals0.Valid {rec.Token0Decimals = int(decimals0.Int64)	}
	if decimals1.Valid {rec.Token1Decimals = int(decimals1.Int64) }
	if name0.Valid { rec.Token0Name = name0.String }
	if name1.Valid { rec.Token1Name = name1.String }
	if symbol0.Valid { rec.Token0Symbol = symbol0.String }
	if symbol1.Valid { rec.Token1Symbol = symbol1.String }
	rec.NumAugurTokens,_ = ss.Get_uniswap_augur_tokens(pair_aid)
	return rec,nil
}
func (ss *SQLStorage) Get_uniswap_swaps(pair_aid int64,offset int,limit int) []p.UniswapSwap {

	records := make([]p.UniswapSwap,0,128)
	var query string
	query = "SELECT " +
				"sw.id,"+
				"sw.block_num," +
				"sw.amount0_in, " +
				"sw.amount1_in," +
				"sw.amount0_out," +
				"sw.amount1_out," +
				"sw.time_stamp," +
				"EXTRACT(EPOCH FROM sw.time_stamp)::BIGINT AS created_ts, "+
				"ra.addr AS recipient_addr," +
				"sw.recipient_aid " +
			"FROM uswap1 AS sw " +
			"LEFT JOIN address AS ra ON sw.recipient_aid=ra.address_id " +
			"WHERE sw.pair_aid=$1 " +
			"ORDER BY sw.time_stamp DESC "+
			"OFFSET $2 LIMIT $3"

	rows,err := ss.db.Query(query,pair_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.UniswapSwap
		err=rows.Scan(
			&rec.Id,
			&rec.BlockNum,
			&rec.Amount0_In,
			&rec.Amount1_In,
			&rec.Amount0_Out,
			&rec.Amount1_Out,
			&rec.CreatedDate,
			&rec.CreatedTs,
			&rec.RequesterAddr,
			&rec.RequesterAid,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_uniswap_volume(market_aid int64,outc int,init_ts,fin_ts,interval int) []p.TradingVolume {

	records := make([]p.TradingVolume,0,64)
	var query string
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM sw.time_stamp))::BIGINT AS ts " +
			"FROM uswap1 sw " +
			"JOIN upair AS sp ON sw.pair_aid=sp.pair_aid " +
			"JOIN af_wrapper w ON (sp.token0_aid=w.wrapper_aid) OR (sp.token1_aid=w.wrapper_aid)"+
			"WHERE w.market_aid=$1 AND w.outcome_idx=$2 " +
			"ORDER BY sw.time_stamp ASC LIMIT 1"

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
				"FLOOR(EXTRACT(EPOCH FROM sw.time_stamp))::BIGINT AS ts " +
			"FROM uswap1 sw " +
			"JOIN upair AS sp ON sw.pair_aid=sp.pair_aid " +
			"JOIN af_wrapper w ON (sp.token0_aid=w.wrapper_aid) OR (sp.token1_aid=w.wrapper_aid) "+
			"WHERE w.market_aid=$1 AND w.outcome_idx=$2 " +
			"ORDER BY sw.time_stamp DESC LIMIT 1"

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
				"SUM(ABS(amount)) AS volume " +
			"FROM periods AS p " +
				"LEFT JOIN (" +
						"SELECT s.id,(amount*POWER(10,3)) AS amount,s.time_stamp AS ts " +
						"FROM uswap2 s " +
						"JOIN af_wrapper w ON s.token_aid=w.wrapper_aid " +
						"WHERE w.market_aid=$4 AND w.outcome_idx=$5 " +
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
func (ss *SQLStorage) Get_uniswap_token_prices(pair_aid int64,inverse bool,init_ts,fin_ts int) []p.UPairPrice {

	var price_field = "amount1_in/amount0_out AS price "
	var price_cond = " amount0_out > 0 "
	if inverse {
		price_field = "amount0_in/amount1_out AS price "
		price_cond =" amount1_out > 0 "
	}
	var query string
	query =	"SELECT " +
				"id," +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts,"+
				"time_stamp," +
				price_field +
			"FROM uswap1 " +
			"WHERE pair_aid=$1 AND " + price_cond + " AND " +
				"time_stamp >= TO_TIMESTAMP($2) AND " +
				"time_stamp < TO_TIMESTAMP($3) " +
			"ORDER BY time_stamp"

	records := make([]p.UPairPrice,0,128)

	rows,err := ss.db.Query(query,pair_aid,init_ts,fin_ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.UPairPrice
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
func (ss *SQLStorage) Get_uniswap_swap_by_id(id int64) (p.UniswapSwap,error) {

	var rec p.UniswapSwap
	var query string
	query = "SELECT " +
				"sw.id,"+
				"sw.pair_aid," +
				"pa.addr," +
				"sw.block_num," +
				"sw.amount0_in, " +
				"sw.amount1_in," +
				"sw.amount0_out," +
				"sw.amount1_out," +
				"sw.time_stamp," +
				"EXTRACT(EPOCH FROM sw.time_stamp)::BIGINT AS created_ts, "+
				"ra.addr AS recipient_addr," +
				"sw.recipient_aid, " +
				"e1.symbol,"+
				"e2.symbol "+
			"FROM uswap1 AS sw " +
			"JOIN upair AS p ON sw.pair_aid=p.pair_aid " +
			"LEFT JOIN address AS ra ON sw.recipient_aid=ra.address_id " +
			"LEFT JOIN address AS pa ON sw.pair_aid=pa.address_id " +
			"LEFT JOIN erc20_info e1 ON p.token0_aid=e1.aid " +
			"LEFT JOIN erc20_info e2 ON p.token1_aid=e2.aid " +
			"WHERE sw.id=$1 " +
			"ORDER BY sw.id"

	res := ss.db.QueryRow(query,id)
	err := res.Scan(
		&rec.Id,
		&rec.PairAid,
		&rec.PairAddr,
		&rec.BlockNum,
		&rec.Amount0_In,
		&rec.Amount1_In,
		&rec.Amount0_Out,
		&rec.Amount1_Out,
		&rec.CreatedDate,
		&rec.CreatedTs,
		&rec.RequesterAddr,
		&rec.RequesterAid,
		&rec.Symbol0,
		&rec.Symbol1,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return rec,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}

	return rec,nil
}
func (ss *SQLStorage) Get_uniswap_augur_tokens(pair_aid int64) (int64,error) {
	// returns number of Augur-related tokens
	var query string
	query = "SELECT " +
				"count(*) AS num_toks " +
			"FROM upair AS p " +
				"JOIN af_wrapper aw ON (p.token0_aid=aw.wrapper_aid OR p.token1_aid=aw.wrapper_aid) " +
			"WHERE p.pair_aid=$1"

	row := ss.db.QueryRow(query,pair_aid)
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
