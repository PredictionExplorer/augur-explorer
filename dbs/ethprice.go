package dbs

import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Ethprice_get_last_processed_block() int64 {

	var query string
	query = "SELECT " +
				"block_num "+
			"FROM "+ss.schema_name+".ep_swap "+
			"ORDER BY block_num DESC "+
			"LIMIT 1"

	var block_num sql.NullInt64 
	res := ss.db.QueryRow(query)
	err := res.Scan(&block_num)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return block_num.Int64
}
func (ss *SQLStorage) Ethprice_insert_swap_event(evt *p.EthpriceSwap) {

	var ethusd_price float64
	eth_amount := big.NewFloat(0.0)
	eth_amount.SetString(evt.Amount1)
	eth_amount.Abs(eth_amount)
	eth_divisor := big.NewFloat(0.0)
	eth_divisor.SetString("1000000000000000000")
	eth_amount.Quo(eth_amount,eth_divisor)
	dollar_amount := big.NewFloat(0.0)
	dollar_amount.SetString(evt.Amount0)
	dollar_amount.Abs(dollar_amount)
	dollar_divisor := big.NewFloat(1e+6)
	if evt.TokenCode == 1 {// DAI
		dollar_divisor.SetString("1000000000000000000")
	}
	dollar_amount.Quo(dollar_amount,dollar_divisor)

	dollar_amount.Quo(dollar_amount,eth_amount) /// calculate ethusd price
	ethusd_price,_ = dollar_amount.Float64()

	var query string
	query = "INSERT INTO "+ss.schema_name+".ep_swap(tx_hash,time_stamp,block_num,tx_idx,log_idx,token_code,"+
						"sender,recipient,amount0,amount1,sqrt_price,liquidity,tick,ethusd_price)" +
			"VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)"

	_,err := ss.db.Exec(query,
		evt.TxHash,
		evt.TimeStamp,
		evt.BlockNum,
		evt.TxIdx,
		evt.LogIdx,
		evt.TokenCode,
		evt.Sender,
		evt.Recipient,
		evt.Amount0,
		evt.Amount1,
		evt.SqrtPrice,
		evt.Liquidity,
		evt.Tick,
		ethusd_price,
	)
	if (err!=nil) {
		unique := strings.Contains(err.Error(),"duplicate key value")
		if unique {
			ss.Log_msg(
				fmt.Sprintf(
					"Attempt to insert event that was already exists" +
					", blocknum=%v, tx_hash %v, log_idx %v",
					evt.BlockNum,evt.TxHash,evt.LogIdx,
				),
			)
			return
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Ethprice_get_ethusd_price_closest_to_timestamp(ts int64) (float64,bool) {

	var (
		price1		float64
		price2		float64
		ts1			int64
		ts2			int64
		have_first	bool = false
		have_second	bool = false
	)
	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"ethusd_price "+
			"FROM "+ss.schema_name+".ep_swap "+
			"WHERE ts<TO_TIMESTAMP($1) "+
			"ORDER BY ts DESC "+
			"LIMIT 1"


	row = sw.S.Db().QueryRow(query,ts)
	err=row.Scan(&ts1,&price1);
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Ethprice_get_ethusd_price_closest_to_timestamp(): %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		have_first = true
	}

	query = "SELECT "+
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT ts," +
				"ethusd_price "+
			"FROM "+ss.schema_name+".ep_swap "+
			"WHERE ts>TO_TIMESTAMP($1)"
			"ORDER BY ts ASC "+
			"LIMIT 1"

	row = sw.S.Db().QueryRow(query,ts)
	err=row.Scan(&ts2,&price2);
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Ethprice_get_ethusd_price_closest_to_timestamp(): %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		have_second = true
	}
	diff1 := ts - ts1
	if diff1 < 0 { diff1 = -diff1 }
	diff2 := ts - ts2
	if diff2 < 0 { diff2 = -diff2 }

	// we pick the record that is the closest to the timestamp asked
	if have_first && have_second {
		if diff1 < diff2 {
			return price1,true
		}
		return price2
	} else {
		if have_first {
			return price1
		}
		if have_second {
			return price2
		}
	}
	return 0.0,false
}
