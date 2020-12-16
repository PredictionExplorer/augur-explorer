package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_ethusd_price_events(from_evt_id int64) []p.EthUsdPriceEvt {

	var usdc_pair_aid int64 = ss.Lookup_address_id("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc")
	var dai_pair_aid int64 = ss.Lookup_address_id("0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11")
	// USDC (pair_aid=233)
	//		token0	USDC	0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48	aid=50
	//		token1	WETH	0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2	aid=8

	// DAI (pair_aid=576)
	//		token0	DAI		0x6B175474E89094C44Da98b954EedeAC495271d0F	aid=53
	//		token1	WETH	0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2	aid=8

	records := make([]p.EthUsdPriceEvt,0,8)
	var where_cond string = fmt.Sprintf("%v,%v",usdc_pair_aid,dai_pair_aid)
	var query string
	query = "SELECT " +
					"s.evtlog_id,s.tx_id,s.block_num," +
					"FLOOR(EXTRACT(EPOCH FROM s.time_stamp))::BIGINT, " +
					"amount0_in,amount1_in,amount0_out,amount1_out " +
			"FROM uswap1 AS s " +
			"JOIN evt_log AS e ON e.id=s.evtlog_id " +
			"WHERE pair_aid IN("+where_cond+") AND s.evtlog_id > $1 "+
			"LIMIT 2000"
	rows,err := ss.db.Query(query,from_evt_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return records
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.EthUsdPriceEvt
		err=rows.Scan(
			&rec.EvtId,
			&rec.TxId,
			&rec.BlockNum,
			&rec.TimeStamp,
			&rec.Amount0In,
			&rec.Amount1In,
			&rec.Amount0Out,
			&rec.Amount1Out,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_ethusd_price_process_status() p.EthUsdProcessStatus {

	var output p.EthUsdProcessStatus
	var null_last_evtid sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM ethusd_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_last_evtid)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO ethusd_status DEFAULT VALUES"
				_,err := ss.db.Exec(query)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			}
		} else {
			break
		}
	}
	if null_last_evtid.Valid {
		output.LastEvtId  = null_last_evtid.Int64
	}
	return output
}
func (ss *SQLStorage) Update_ethusd_process_status(status *p.EthUsdProcessStatus) {

	var query string
	query = "UPDATE ethusd_status SET last_evt_id = $1"

	_,err := ss.db.Exec(query,status.LastEvtId)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_ethusd_price_evt(evtlog_id int64) {

	var query string
	query = "DELETE FROM ethusd_price WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evtlog_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_ethusd_price_evt(pr *p.EthUsdPriceEvt) {

	var query string
	query = "INSERT INTO ethusd_price( " +
				"evtlog_id,block_num,tx_id,time_stamp,eth_price" +
			") VALUES($1,$2,$3,TO_TIMESTAMP($4),$5)"

	_,err := ss.db.Exec(query,
		pr.EvtId,
		pr.BlockNum,
		pr.TxId,
		pr.TimeStamp,
		pr.EthUsd,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_ethusd_price_history(init_ts,fin_ts int) (int,int,[]p.EthUsdPrice) {

	var query string
	if fin_ts == 2147483647 {
		query = "SELECT  EXTRACT(EPOCH FROM time_stamp)::BIGINT AS ending_ts "+
				"FROM ethusd_price " +
				"ORDER BY ending_ts DESC LIMIT 1"

		var null_ts sql.NullInt64
		err := ss.db.QueryRow(query).Scan(&null_ts)
		ss.adjust_ts(&fin_ts,err,&null_ts)
		fin_ts++
	}
	if init_ts == 0 {
		query = "SELECT  EXTRACT(EPOCH FROM time_stamp)::BIGINT AS starting_ts "+
				"FROM ethusd_price " +
				"ORDER BY starting_ts LIMIT 1"
		var null_ts sql.NullInt64
		err := ss.db.QueryRow(query).Scan(&null_ts)
		ss.adjust_ts(&init_ts,err,&null_ts)

	}
/*	query =	"SELECT EXTRACT(EPOCH FROM time_stamp)::BIGINT AS ts,eth_price FROM ethusd_price " +
			"WHERE time_stamp >= TO_TIMESTAMP($1) AND time_stamp < TO_TIMESTAMP($2) " +
			"ORDER by time_stamp "*/
	query =	"SELECT " +
				"AVG(p.eth_price), "+
				"(EXTRACT(EPOCH FROM time_stamp)::BIGINT/300)*300 AS ts "+
				"FROM ethusd_price AS p " +
			"WHERE time_stamp >= TO_TIMESTAMP($1) AND time_stamp < TO_TIMESTAMP($2) " +
			"GROUP BY ts " +
			"ORDER by ts "

	records := make([]p.EthUsdPrice,0,8)
	rows,err := ss.db.Query(query,init_ts,fin_ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return init_ts,fin_ts,records
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.EthUsdPrice
		err=rows.Scan(
			&rec.Price,
			&rec.TimeStamp,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return init_ts,fin_ts,records
}
func (ss *SQLStorage) Get_last_ethusd_price() (float64,error) {


	var query string
	query = "SELECT eth_price FROM ethusd_price " +
			"ORDER by block_num DESC,tx_id DESC,evtlog_id DESC " +
			"LIMIT 1"
	res := ss.db.QueryRow(query)
	var null_price sql.NullFloat64
	err := res.Scan(&null_price)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0.0,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return null_price.Float64,nil
}
