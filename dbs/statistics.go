package dbs

import (
	"fmt"
	"os"
	"strings"
	"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
/* DISCONTINUED, to be deleted
func (ss *SQLStorage) Check_main_stats() {

	var query string
	query="SELECT id FROM main_stats LIMIT 1";
	row := ss.db.QueryRow(query)
	var null_id sql.NullInt64
	var err error
	err=row.Scan(&null_id);
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("Error in check_main_stats(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
}
*/
func (ss *SQLStorage) Get_main_stats() p.MainStats {

	var query string
	query = "SELECT " +
				"markets_count," +
				"yesno_count," +
				"categ_count," +
				"scalar_count," +
				"active_count," +
				"money_at_stake," +
				"trades_count " +
			"FROM main_stats "

	row := ss.db.QueryRow(query)
	var err error
	var s p.MainStats
	err=row.Scan(
				&s.MarketsCount,
				&s.YesNoCount,
				&s.CategCount,
				&s.ScalarCount,
				&s.ActiveCount,
				&s.MoneyAtStake,
				&s.TradesCount,
	);
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
			os.Exit(1)
		}
	}
	s.FinalizedCount = (s.YesNoCount + s.CategCount + s.ScalarCount) - s.ActiveCount
	s.LastBlockNum,_ =  ss.Get_last_block_num()
	query = "SELECT count(*) AS total FROM market WHERE status=5"
	row = ss.db.QueryRow(query)
	var null_count sql.NullInt64
	err = row.Scan(&null_count)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	if null_count.Valid {
		s.InvalidCount = null_count.Int64
	}

	return s
}
func (ss *SQLStorage) Get_front_page_stats() p.FrontPageStats {

	var stats p.FrontPageStats
	var query string
	query = "SELECT markets_count,money_at_stake,trades_count " +
			"FROM main_stats LIMIT 1" // ToDo: we need support for multiple Universes
	row := ss.db.QueryRow(query)
	err := row.Scan(
				&stats.MarketsCreated,
				&stats.MoneyAtStake,
				&stats.TradesCount,
	)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	return stats
}
func (ss *SQLStorage) Get_last_unique_addr_day() int64 {

	var query string
	query = "SELECT EXTRACT(EPOCH FROM day::TIMESTAMP)::BIGINT AS ts FROM unique_addrs ORDER BY day DESC LIMIT 1"
	row := ss.db.QueryRow(query)
	var ts int64
	var err error
	err=row.Scan(&ts);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_last_unique_addr_day(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return ts
}
func (ss *SQLStorage) Update_unique_addresses_entry(ts int64,num_addrs int64) {
	var query string
	query = "UPDATE unique_addrs SET num_addrs = $2 WHERE day=to_timestamp($1)"
	res,err:=ss.db.Exec(query,ts,num_addrs)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Update_unique_addresses_entry() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in Update_unique_addresses_entry(): %v",err))
		os.Exit(1)
	}
	if affected_rows == 0 {
		query = "INSERT INTO unique_addrs(day,num_addrs) VALUES(to_timestamp($1),$2)"
		_,err := ss.db.Exec(query,ts,num_addrs)
		if (err!=nil) {
			ss.Log_msg(
				fmt.Sprintf(
					"DB Error on INSERT in Update_unique_addresses_entry(): %v q=%v",
					err,query,
				),
			);
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Get_unique_addresses() []p.UniqueAddrEntry {

	records := make([]p.UniqueAddrEntry,0,365)
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM day::TIMESTAMP)::BIGINT AS ts,"+
				"day," +
				"num_addrs "+
			"FROM unique_addrs ORDER BY day"
	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var accumulator int64 = 0
	defer rows.Close()
	for rows.Next() {
		var rec p.UniqueAddrEntry
		err=rows.Scan(&rec.Ts,&rec.Day,&rec.NumAddrs)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		accumulator = accumulator + rec.NumAddrs
		rec.NumAddrsAccum = accumulator
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Calc_unique_addresses(ts_from int64,ts_to int64) (int64,bool) {

	var query string
	query = "SELECT count(*) FROM ( " +
				"SELECT DISTINCT u.eoa_aid FROM address a " +
				"JOIN ustats u ON u.eoa_aid=a.address_id " +
				"JOIN block b ON a.block_num=b.block_num " +
				"WHERE b.ts >= to_timestamp($1) AND b.ts < to_timestamp($2)" +
			") AS s"
	row := ss.db.QueryRow(query,ts_from,ts_to)
	var null_counter sql.NullInt64
	var err error
	err=row.Scan(&null_counter);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,true	// this will never happen
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Calc_unique_addresses(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	query = "SELECT b.block_num FROM block b " +
			"WHERE b.ts >= to_timestamp($1) AND b.ts < to_timestamp($2)" +
			"LIMIT 1"
	row = ss.db.QueryRow(query,ts_from,ts_to)
	var no_rows bool = false
	var null_block_num sql.NullInt64
	err = row.Scan(&null_block_num)
	if err!=nil {
		if err == sql.ErrNoRows {
			no_rows=true
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Calc_unique_addresses(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}

	return null_counter.Int64,no_rows
}
func (ss *SQLStorage) Link_eoa_and_wallet_contract(eoa_aid, wallet_aid int64) {

	var query string
	query = "INSERT INTO ustats(eoa_aid,wallet_aid) " +
			"VALUES($1,$2)"

	res,err:=ss.db.Exec(query,eoa_aid,wallet_aid)
	ss.Info.Printf("reporting eoa2wallet link sql error: %v  eoa=%v wallet=%v\n",err,eoa_aid,wallet_aid)
	if (err!=nil) {
		unique_wallet := strings.Contains(err.Error(),`duplicate key value violates unique constraint "ustats2_idx"`)
		unique_eoa := strings.Contains(err.Error(),`duplicate key value violates unique constraint "ustats_pkey"`)
		if unique_wallet || unique_eoa {
			if eoa_aid != wallet_aid {
				// In rare cases we can have a record with eoa_aid=wallet_aid 
				//  and it may be preventing the INSERT. If this is the case check if we can fix it
				query = "SELECT wallet_aid FROM ustats where eoa_aid=$1"
				var stored_wallet_aid int64 = 0
				row := ss.db.QueryRow(query,eoa_aid)
				err := row.Scan(&stored_wallet_aid)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("Error fixing wallet_aid: %v\n",err))
					os.Exit(1)
				}
				if stored_wallet_aid == eoa_aid {
					// rare case confirmed, we have eoa_aid=wallet_aid in ustats, so we can update
					// wallet_aid with the new value
					query = "UPDATE ustats SET wallet_aid = $3 WHERE eoa_aid=$1 AND wallet_aid=$2"
					_,err:=ss.db.Exec(query,eoa_aid,eoa_aid,wallet_aid)
					if (err!=nil) {
						ss.Log_msg(fmt.Sprintf("Update ustats failed: %v, q=%v",err,query))
						os.Exit(1)
					}
					ss.Info.Printf(
						"eoa2wallet link UPDATE: new wallet_id=%v was set for eoa_aid=%v\n",
						wallet_aid,eoa_aid,
					)
				}
			} else {
				// EOA=Wallet , we already have this record, duplicated registration ignored
			}
		} else {
			ss.Info.Printf(
				"eoa2wallet link sql error other than dup-key: %v  eoa=%v wallet=%v\n",
				err,eoa_aid,wallet_aid,
			)
			os.Exit(1)
		}
	} else {
		affected_rows,err:=res.RowsAffected()
		if err == nil {
			if affected_rows > 0 {
				ss.Info.Printf("eoa2wallet link success: eoa=%v wallet=%v\n",eoa_aid,wallet_aid)
			} else {
				ss.Info.Printf(
					"eoa2wallet link without effect (affected rows=0): eoa=%v wallet=%v\n",
					eoa_aid,wallet_aid,
				)
			}
		} else {
			ss.Log_msg(fmt.Sprintf("DB error on getting affected rows: %v\n",err))
			os.Exit(1)
		}
	}
}
func gas_usage_query(table_name string) string {

	var query string
	date_cond := "((b.ts >= to_timestamp($1)) AND (b.ts < to_timestamp($2)))"

	query =
		"SELECT " +
			"SUM(gas_used::decimal)::text AS sum," +
			"SUM(gas_used::decimal * gas_price)::text AS eth_sum,"+
			"count(*) AS num_rows " +
		"FROM (" +
			"SELECT DISTINCT tx_id FROM (" +
				"(SELECT t.tx_id FROM "+table_name+" AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
			") AS d " +
		") AS i " +
		"JOIN transaction AS t ON i.tx_id = t.id " +
		"JOIN block AS b ON t.block_num=b.block_num "

	return query
}
func (ss *SQLStorage) gas_usage_get_results(query *string,ts_from int64,ts_to int64,value *string,eth_value *string,counter *int64) {

	row := ss.db.QueryRow(*query,ts_from,ts_to)
	var null_val sql.NullString
	var null_eth_val sql.NullString
	var null_counter sql.NullInt64
	var err error
	err=row.Scan(&null_val,&null_eth_val,&null_counter);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Calc_gas_global(): %v, q=%v",err,*query))
			os.Exit(1)
		}
	}
	if null_val.Valid {
		*value = null_val.String
	} else {
		*value = "0"
	}
	if null_eth_val.Valid {
		*eth_value = null_eth_val.String
	} else {
		*eth_value = "0"
	}
	if null_counter.Valid {
		*counter= null_counter.Int64
	}
}
func (ss *SQLStorage) Calc_gas_usage_global(from int64,to int64) p.GasSpent {

	var output p.GasSpent

	// TRADING
	var query string
	date_cond := "((b.ts >= to_timestamp($1)) AND (b.ts < to_timestamp($2)))"
	// in this query we are picking only Augur transactions
	query =
		"SELECT " +
				"SUM(gas_used::decimal)::text AS sum," +
				"SUM(gas_used::decimal * gas_price)::text AS eth_sum,"+
				"count(*) AS num_rows " +
		"FROM (" +
			"SELECT DISTINCT tx_id FROM (" +
				"(SELECT t.tx_id FROM mktord AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
					"UNION ALL" +
				"(SELECT t.tx_id FROM profit_loss AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
			") AS d " +
		") AS i " +
		"JOIN transaction AS t ON i.tx_id = t.id " +
		"JOIN block AS b ON t.block_num=b.block_num "

	ss.gas_usage_get_results(&query,from,to,&output.Trading,&output.EthTrading,&output.Num_trading)

	// REPORTING 
	query = gas_usage_query("report")
	ss.gas_usage_get_results(&query,from,to,&output.Reporting,&output.EthReporting,&output.Num_reporting)

	// MARKETS CREATED
	query = gas_usage_query("market")
	ss.gas_usage_get_results(&query,from,to,&output.Markets,&output.EthMarkets,&output.Num_markets)

	// EVERYTHING
	query =
		"SELECT "+
				"SUM(gas_used::decimal)::text AS sum,"+
				"SUM(gas_used::decimal * gas_price)::text AS eth_sum,"+
				"count(*) AS num_rows "+
			"FROM (" +
			"SELECT DISTINCT tx_id FROM (" +
				"(SELECT t.tx_id FROM mktord AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
					"UNION ALL" +
				"(SELECT t.tx_id FROM market AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
					"UNION ALL" +
				"(SELECT t.tx_id FROM mkt_fin AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
					"UNION ALL" +
				"(SELECT t.tx_id FROM claim_funds AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
					"UNION ALL" +
				"(SELECT t.tx_id FROM sbalances AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
					"UNION ALL" +
				"(SELECT t.tx_id FROM volume AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
					"UNION ALL" +
				"(SELECT t.tx_id FROM oi_chg AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
					"UNION ALL" +
				"(SELECT t.tx_id FROM profit_loss AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
					"UNION ALL" +
				"(SELECT t.tx_id FROM report AS t,block AS b WHERE (b.block_num=t.block_num) AND " + date_cond + " )" +
			") AS d " +
		") AS i " +
		"JOIN transaction AS t ON i.tx_id = t.id " +
		"JOIN block AS b ON t.block_num=b.block_num "
	ss.gas_usage_get_results(&query,from,to,&output.Total,&output.EthTotal,&output.Num_total)

	return output
}
func (ss *SQLStorage) Update_global_gas_stats(day int64,stats *p.GasSpent) {
	var query string
	query = "UPDATE gas_spent " +
		"SET trading="+stats.Trading+",num_trading=$2,"+
			"reporting="+stats.Reporting+",num_reporting=$3,"+
			"markets="+stats.Markets+",num_markets=$4,"+
			"total="+stats.Total+",num_total=$5, " +
			"eth_trading="+stats.EthTrading+","+
			"eth_reporting="+stats.EthReporting+","+
			"eth_markets="+stats.EthMarkets+"," +
			"eth_total="+stats.EthTotal+" "+
		"WHERE day=to_timestamp($1)"
	res,err:=ss.db.Exec(query,
		day,
		stats.Num_trading,
		stats.Num_reporting,
		stats.Num_markets,
		stats.Num_total,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Update_global_gas_stats() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in Update_global_gas_stats(): %v",err))
		os.Exit(1)
	}
	if affected_rows == 0 {
		query = "INSERT INTO gas_spent(" +
					"day,"+
					"trading,num_trading,"+
					"reporting,num_reporting,"+
					"markets,num_markets,"+
					"total,num_total," +
					"eth_trading,"+
					"eth_reporting,"+
					"eth_markets,"+
					"eth_total"+
				") "+
				"VALUES(" +
					"to_timestamp($1)," +
					stats.Trading + ",$2,"+
					stats.Reporting + ",$3,"+
					stats.Markets + ",$4,"+
					stats.Total + ",$5,"+
					stats.EthTrading + ","+
					stats.EthReporting + "," +
					stats.EthMarkets + "," +
					stats.EthTotal + 
				")"
		_,err := ss.db.Exec(query,
				day,
				stats.Num_trading,
				stats.Num_reporting,
				stats.Num_markets,
				stats.Num_total,
		)
		if err!=nil {
			ss.Log_msg(
				fmt.Sprintf(
					"DB Error on INSERT in Update_global_gas_stats() for day %v: %v q=%v",
					day,err,query,
				),
			);
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Get_gas_usage_global() []p.GasSpent {

	var query string
	query =
		"SELECT " +
			"EXTRACT(EPOCH FROM day::TIMESTAMP)::BIGINT AS ts,"+
			"trading,reporting,markets,total," +
			"eth_trading,eth_reporting,eth_markets,eth_total,"+
			"num_trading,num_reporting,num_markets,num_total "+
		"FROM gas_spent "+
		"ORDER BY day"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.GasSpent,0,256)
	defer rows.Close()
	for rows.Next() {
		var rec p.GasSpent
		err=rows.Scan(
			&rec.Ts,
			&rec.Trading,&rec.Reporting,&rec.Markets,&rec.Total,
			&rec.EthTrading,&rec.EthReporting,&rec.EthMarkets,&rec.EthTotal,
			&rec.Num_trading,&rec.Num_reporting,&rec.Num_markets,&rec.Num_total,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records

}
func (ss *SQLStorage) Get_gas_used() []p.GasUsed {

	var query string
	query =
		"SELECT " +
			"EXTRACT(EPOCH FROM day::TIMESTAMP)::BIGINT AS ts,"+
			"trading,reporting,markets,total," +
			"num_trading,num_reporting,num_markets,num_total "+
		"FROM gas_spent "+
		"ORDER BY day"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.GasUsed,0,256)
	defer rows.Close()
	for rows.Next() {
		var rec p.GasUsed
		err=rows.Scan(
			&rec.Ts,
			&rec.Trading,&rec.Reporting,&rec.Markets,&rec.Total,
			&rec.Num_trading,&rec.Num_reporting,&rec.Num_markets,&rec.Num_total,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records

}
func (ss *SQLStorage) Get_txcost(init_ts,fin_ts int) []p.TxCost {

	var query string
	query =
		"SELECT " +
			"EXTRACT(EPOCH FROM day::TIMESTAMP)::BIGINT AS ts,"+
			"eth_trading,eth_reporting,eth_markets,eth_total,"+
			"num_trading,num_reporting,num_markets,num_total "+
		"FROM gas_spent " +
		"WHERE day >= TO_TIMESTAMP($1)::DATE AND day < TO_TIMESTAMP($2)::DATE " +
		"ORDER BY day"

	rows,err := ss.db.Query(query,init_ts,fin_ts)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.TxCost,0,256)
	defer rows.Close()
	for rows.Next() {
		var rec p.TxCost
		err=rows.Scan(
			&rec.Ts,
			&rec.EthTrading,&rec.EthReporting,&rec.EthMarkets,&rec.EthTotal,
			&rec.Num_trading,&rec.Num_reporting,&rec.Num_markets,&rec.Num_total,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) get_initial_gas_usage(table_name string,init_ts int,everything bool) (int64,int64,float64) {

	var query string
	query = 
		"SELECT " +
			"COALESCE(COUNT(tx_id),0) as num_rows, " +
			"SUM(gas_used::decimal) AS sum," +
			"SUM(gas_used::decimal * gas_price) as tx_fee " +
		"FROM (" +
			"SELECT tx_id,b.ts,gas_used,gas_price " +
			"FROM ("

				if everything {
					query = query +
					"SELECT DISTINCT tx_id FROM (" +
						"(SELECT t.tx_id FROM mktord AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM market AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM mkt_fin AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM claim_funds AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM sbalances AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM volume AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM oi_chg AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM profit_loss AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM report AS t,block AS b WHERE (b.block_num=t.block_num))" +
					") AS d "
				} else {
					query = query +
					"SELECT DISTINCT tx_id FROM (" +
						"("+

							"SELECT t.tx_id FROM "+table_name+" AS t,block AS b " +
							"WHERE (b.block_num=t.block_num)" +
						")" +
					") AS d "
				}
	query = query +
			") AS i " +
			"JOIN transaction AS t ON i.tx_id = t.id " +
			"JOIN block AS b ON t.block_num=b.block_num " +
		") AS t " +
		"WHERE t.ts < TO_TIMESTAMP($1) "

	row := ss.db.QueryRow(query,init_ts)
	var null_count,null_gas_used sql.NullInt64
	var null_txfee sql.NullFloat64
	var err error
	err=row.Scan(&null_count,&null_gas_used,&null_txfee);
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return null_count.Int64,null_gas_used.Int64,null_txfee.Float64
}
func gas_usage_accum_query(table_name string,everything bool) string {

	var query string

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
			"COALESCE(COUNT(tx_id),0) as num_rows, " +
			"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
			"SUM(gas_used::decimal) AS sum," +
			"SUM(gas_used::decimal * gas_price) as tx_fee " +
		"FROM periods AS p " +
		"LEFT JOIN (" +
			"SELECT tx_id,b.ts,gas_used,gas_price " +
			"FROM ("

				if everything {
					query = query +
					"SELECT DISTINCT tx_id FROM (" +
						"(SELECT t.tx_id FROM mktord AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM market AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM mkt_fin AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM claim_funds AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM sbalances AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM volume AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM oi_chg AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM profit_loss AS t,block AS b WHERE (b.block_num=t.block_num))" +
							"UNION ALL" +
						"(SELECT t.tx_id FROM report AS t,block AS b WHERE (b.block_num=t.block_num))" +
					") AS d " 
				} else {
					query = query +
					"SELECT DISTINCT tx_id FROM (" +
						"("+
							"SELECT t.tx_id FROM "+table_name+" AS t,block AS b " +
							"WHERE (b.block_num=t.block_num)" +
						")" +
					") AS d "
				}

	query = query +
			") AS i " +
			"JOIN transaction AS t ON i.tx_id = t.id " +
			"JOIN block AS b ON t.block_num=b.block_num " +
		") AS t ON (" +
			"p.start_ts <= t.ts AND "+
			"t.ts < p.end_ts " +
		") " +
		"GROUP BY start_ts " +
		"ORDER BY start_ts"

	return query
}
func (ss *SQLStorage) get_counter_data(table_name string,init_ts,fin_ts,interval int,everything bool) []p.GasCounter {

	var query string
	query = gas_usage_accum_query(table_name,everything)

	rows,err := ss.db.Query(query,init_ts,fin_ts,interval)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	count,gas_used,txfee := ss.get_initial_gas_usage(table_name,init_ts,everything)
	records := make([]p.GasCounter,0,256)
	defer rows.Close()
	for rows.Next() {
		var rec p.GasCounter
		var null_gas_used sql.NullInt64
		var null_txcost sql.NullFloat64
		err=rows.Scan(
			&rec.NumRecs,&rec.TimeStamp,&null_gas_used,&null_txcost,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		count = count + rec.NumRecs
		rec.NumRecs = count
		if null_gas_used.Valid {
			gas_used = gas_used + null_gas_used.Int64
		}
		rec.GasUsed = gas_used
		if null_txcost.Valid {
			txfee = txfee + null_txcost.Float64
		}
		rec.TxCost = txfee
		records = append(records,rec)
	}
	return records

}
func (ss *SQLStorage) Get_gasused_accum(init_ts,fin_ts,interval int) []p.AccumGasUsed {

	records := make([]p.AccumGasUsed,0,256)
	{
		stats := ss.get_counter_data("mktord",init_ts,fin_ts,interval,false)
		for _,elt := range stats {
			var rec p.AccumGasUsed
			rec.Ts = elt.TimeStamp
			rec.Trading = elt.GasUsed
			rec.Num_trading = elt.NumRecs
			records = append(records,rec)
		}
	}
	{
		stats := ss.get_counter_data("report",init_ts,fin_ts,interval,false)
		for i,elt := range stats {
			records[i].Reporting = elt.GasUsed
			records[i].Num_reporting = elt.NumRecs
		}
	}
	{
		stats := ss.get_counter_data("market",init_ts,fin_ts,interval,false)
		for i,elt := range stats {
			records[i].Markets = elt.GasUsed
			records[i].Num_markets = elt.NumRecs
		}
	}
	{
		stats := ss.get_counter_data("",init_ts,fin_ts,interval,true)
		for i,elt := range stats {
			records[i].Total = elt.GasUsed
			records[i].Num_total= elt.NumRecs
		}
	}
	return records
}
func (ss *SQLStorage) Get_txcost_accum(init_ts,fin_ts,interval int) []p.AccumTxCost {

	records := make([]p.AccumTxCost,0,256)
	{
		stats := ss.get_counter_data("mktord",init_ts,fin_ts,interval,false)
		for _,elt := range stats {
			var rec p.AccumTxCost
			rec.Ts = elt.TimeStamp
			rec.EthTrading = elt.TxCost
			rec.Num_trading = elt.NumRecs
			records = append(records,rec)
		}
	}
	{
		stats := ss.get_counter_data("report",init_ts,fin_ts,interval,false)
		for i,elt := range stats {
			records[i].EthReporting = elt.TxCost
			records[i].Num_reporting = elt.NumRecs
		}
	}
	{
		stats := ss.get_counter_data("market",init_ts,fin_ts,interval,false)
		for i,elt := range stats {
			records[i].EthMarkets = elt.TxCost
			records[i].Num_markets = elt.NumRecs
		}
	}
	{
		stats := ss.get_counter_data("",init_ts,fin_ts,interval,true)
		for i,elt := range stats {
			records[i].EthTotal = elt.TxCost
			records[i].Num_total= elt.NumRecs
		}
	}
	return records
}
