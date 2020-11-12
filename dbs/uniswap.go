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
func (ss *SQLStorage) Insert_uniswap_pair_swap_evt(pair,user *common.Address,bci *p.BasicChainInfo,evt *p.UPairSwap) {

	// sender and recipient do not contain meaningful data, it is the contract address
	sender_aid := ss.Lookup_or_create_address(evt.Sender.String(),bci.BlockNum,bci.TxId)
	recipient_aid := ss.Lookup_or_create_address(evt.To.String(),bci.BlockNum,bci.TxId)

	pair_aid := ss.Lookup_or_create_address(pair.String(),bci.BlockNum,bci.TxId)
	aid := ss.Lookup_or_create_address(user.String(),bci.BlockNum,bci.TxId)

	amount0_in := evt.Amount0In.String()
	amount1_in := evt.Amount1In.String()
	amount0_out := evt.Amount0Out.String()
	amount1_out := evt.Amount1Out.String()

	var query string
	query = "INSERT INTO uswap1(" +
				"evtlog_id,block_num,tx_id,time_stamp,"+
				"pair_aid,sender_aid,recipient_aid,aid," +
				"amount0_in,amount1_in,amount0_out,amount1_out"+
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),"+
				"$5,$6,$7,$8,"+
				"$9::DECIMAL/1e+18,$10::DECIMAL/1e+18,$11::DECIMAL/1e+18,$12::DECIMAL/1e+18" +
			")"

	_,err := ss.db.Exec(query,
		bci.EvtId,bci.BlockNum,bci.TxId,bci.TimeStamp,
		pair_aid,sender_aid,recipient_aid,aid,
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
