package dbs

import (
	"fmt"
	"os"
	"strings"
	"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Ethprice_get_last_processed_block() int64 {

	var query string
	quer = "SELECT " +
				"block_num "+
			"FROM ep_swap "+
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
	return block_num
}
func (ss *SQLStorage) Ethprice_insert_swap_event(evt *EthpriceSwap) {

	var query string
	query = "INSERT INTO ep_swap(tx_hash,time_stamp,block_num,log_idx,token_code,sender,recipient,"+
									"amount0,amount1,sqrt_price,liquidity,tick)" +
			"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)"

	_,err := ss.db.Exec(query,
		evt.TxHash
		evt.TimeStampm,
		evt.BlockNum,
		evt.LogIdx,
		evt.TokenCode,
		evt.Sender,
		evt.Recipient,
		evt.Amount0,
		evt.Amount1,
		evt.SqrtPrice,
		evt.Liquidity,
		evt.Tick,
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
