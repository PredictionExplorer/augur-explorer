package dbs

import (
	"fmt"
	"os"
	"math/big"
	"bytes"
	"encoding/hex"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_erc20_process_status() p.ERC1155ProcStatus {

	var output p.ERC1155ProcStatus
	var null_last_evtid sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM erc1155_proc_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_last_evtid)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO erc1155_proc_status DEFAULT VALUES"
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
func (ss *SQLStorage) Update_erc1155_process_status(status *p.ERC20ProcessStatus) {

	var query string
	query = "UPDATE erc1155_proc_status SET last_evt_id = $1"

	_,err := ss.db.Exec(query,status.LastEvtId)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Lookup_or_insert_token_id(contract_aid,token_id_hex string) int64 {

	var query string
	query = "SELECT token_id FROM erc1155_tok WHERE contract_aid=$1 AND token_id_hex=$2"
	res := ss.db.QueryRow(query,contract_aid,token_id_hex)
	var null_id sql.NullInt64
	err := res.Scan(&null_id)
	if err != nil {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		query = "INSERT INTO erc1155_tok(contract_aid,token_id_hex) VALUES($1,$2) RETURNING token_id"
		res := ss.db.QueryRow(query,contract_aid,token_id_hex)
		err := res.Scan(&null_id)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	return null_id.Int64
}
func (ss *SQLStorage) Insert_ERC1155_transfer_single(
	evt *p.ERC115TransferSingle,
	contract_addr string,
	block_num,
	tx_id int64,
	evtlog_id int64,
	timestamp int64,
) {

	token_id_hex := hex.EncodeToString(common.BigToHash(evt.Id).Bytes())
	contract_aid := ss.Lookup_or_create_address(contract_addr,block_num,tx_id)
	operator_aid := ss.Lookup_or_create_address(evt.Operator.String(),block_num,tx_id)
	from_aid := ss.Lookup_or_create_address(evt.From.String(),block_num,tx_id)
	to_aid := ss.Lookup_or_create_address(evt.To.String(),block_num,tx_id)
	token_id := ss.Lookup_or_insert_token_id(contract_aid,token_id_hex)
	amount := evt.Value.String()
	op_type := int(0)
	var zero_addr common.Address{}
	if bytes.Equal(zero_addr.Bytes(),evt.From.Bytes()) {
		op_type = 1 // mint
	}
	if bytes.Equal(zero_addr.Bytes(),evt.To.Bytes()) {
		op_type = 2 // burn
	}

	var query string
	query = "INSERT INTO erc1155_transf("+
				"evtlog_id,time_stamp,block_num,tx_id,contract_aid,"+
				"operator_aid,token_id,from_aid,to_aid,op_type,amount" +
			") VALUES("+
				"$1,TO_TIMESTAMP($2),$3,$4,$5,"+
				"$6,$7,$9,$9,$10,$11::DECIMAL"+
			")"
	_,err := ss.db.Exec(query,
		evtlog_id,
		timestamp,
		block_num,
		tx_id,
		contract_aid,
		operator_aid,
		token_id,
		from_aid,
		to_aid,
		op_type,
		amount,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_ERC1155_transfer_batch (
	evt *p.ERC115TransferBatch,
	contract_addr string,
	block_num,
	tx_id int64,
	evtlog_id int64,
	timestamp int64,
) {

	contract_aid := ss.Lookup_or_create_address(contract_addr,block_num,tx_id)
	operator_aid := ss.Lookup_or_create_address(evt.Operator.String(),block_num,tx_id)
	from_aid := ss.Lookup_or_create_address(evt.From.String(),block_num,tx_id)
	to_aid := ss.Lookup_or_create_address(evt.To.String(),block_num,tx_id)
	var token_ids string
	for i:=0 ; i<len(evt.Ids); i++ {
		token_id_hex := hex.EncodeToString(common.BigToHash(evt.Ids[i]).Bytes())
		token_id := ss.Lookup_or_insert_token_id(contract_aid,token_id_hex)
		if len(token_ids) > 0 {
			token_ids = token_ids + ","
		}
		token_ids = token_ids + fmt.Sprintf("%v",token_id)
	}
	amounts := Bigint_ptr_slice_to_str(evt.Values,",")
	op_type := int(0)
	var zero_addr common.Address{}
	if bytes.Equal(zero_addr.Bytes(),evt.From.Bytes()) {
		op_type = 1 // mint
	}
	if bytes.Equal(zero_addr.Bytes(),evt.To.Bytes()) {
		op_type = 2 // burn
	}

	var query string
	query = "INSERT INTO erc1155_batch("+
				"evtlog_id,time_stamp,block_num,tx_id,contract_aid,"+
				"operator_aid,token_ids,from_aid,to_aid,op_type,amounts" +
			") VALUES("+
				"$1,TO_TIMESTAMP($2),$3,$4,$5,"+
				"$6,$7,$9,$9,$10,$11"+
			")"
	_,err := ss.db.Exec(query,
		evtlog_id,
		timestamp,
		block_num,
		tx_id,
		contract_aid,
		operator_aid,
		token_ids,
		from_aid,
		to_aid,
		op_type,
		amounts,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
