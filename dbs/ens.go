package dbs

import (
	"fmt"
	"os"
	"log"
	"encoding/hex"

	"database/sql"
	_  "github.com/lib/pq"

//	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_ens_processing_status() p.EnsProcStatus {

	var output p.EnsProcStatus
	var null_id,null_block_num sql.NullInt64

	var query string
	for {
		query = "SELECT block_num_limit,last_evt_id FROM ens_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_block_num,&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO ens_status DEFAULT VALUES"
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
	if null_block_num.Valid {
		output.IniLoadBlockNumLimit = null_block_num.Int64
	}
	return output
}
func (ss *SQLStorage) Insert_name_registered1(rec *p.ENS_Name1) {

	var query string
	var err error
	owner_aid := ss.Lookup_or_create_address(rec.Owner,rec.BlockNum,rec.TxId)
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_name(" +
					"tx_hash,time_stamp,block_num,contract_aid,owner_aid," +
					"label,node,fqdn,name,cost,expires" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10::DECIMAL/1e+18,TO_TIMESTAMP($11))"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			owner_aid,
			rec.Label,
			rec.Node,
			rec.FQDN,
			rec.Name,
			rec.Cost,
			rec.Expires,
		)
	} else {
		query = "INSERT INTO ens_name(" +
					"evtlog_id,block_num,tx_id,contract_aid,owner_aid,"+
					"time_stamp,label,node,fqdn,name,tx_hash,cost,expires" +
				") VALUES(" +
					"$1,$2,$3,$4,$5,TO_TIMESTAMP($6),$7,$8,$9,$10,$11,"+
					"$12::DECIMAL/1e+18,TO_TIMESTAMP($13)"+
				")"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			owner_aid,
			rec.Label,
			rec.Node,
			rec.FQDN,
			rec.Name,
			rec.TxHash,
			rec.Cost,
			rec.Expires,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Expire_ens_names(l *log.Logger) {

	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM expires)::BIGINT, " +
				"label,name " +
			"FROM active_name " +
			"WHERE expires < (NOW() + interval '90 day')"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return
		}
		ss.Log_msg(fmt.Sprintf("Error in get_last_block_num(): %v",err))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var ts int64
		var label,name string
		err=rows.Scan(&ts,&label,&name)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		} else {
			_,err := ss.db.Exec("DELETE FROM active_name WHERE label=$1",label)
			if (err!=nil) {
				ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
				os.Exit(1)
			}
			if l != nil {
				l.Printf("Expiring ENS name %v with label %v\n",name,label)
			}
		}
	}

}
func (ss *SQLStorage) Get_count_of_active_names() int64 {

	var null_count sql.NullInt64
	var query string
	query = "SELECT count(*) AS total FROM active_name"
	_,err := ss.db.Exec(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	if !null_count.Valid {
		ss.Log_msg(fmt.Sprintf("Can't get number of active names"))
		os.Exit(1)
	}
	return null_count.Int64
}
func (ss *SQLStorage) Insert_new_owner(rec *p.ENS_NewOwner) {

	var query string
	var err error
	owner_aid := ss.Lookup_or_create_address(rec.Owner,rec.BlockNum,rec.TxId)
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_new_owner(" +
					"tx_hash,time_stamp,block_num,contract_aid,owner_aid,label,node,fqdn" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8)"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			owner_aid,
			rec.Label,
			rec.Node,
			rec.FQDN,
		)
	} else {
		query = "INSERT INTO ens_new_owner(" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,owner_aid,tx_hash,label,node,fqdn" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9,$10)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			owner_aid,
			rec.TxHash,
			rec.Label,
			rec.Node,
			rec.FQDN,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_hash_invalidated(rec *p.ENS_HashInvalidated) {

	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	var query string
	var err error
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_hash_inval("+
					"tx_hash,time_stamp,block_num,contract_aid,hash,name,value,reg_date" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7::DECIMAL/1e+18,TO_TIMESTAMP($8))"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			rec.Hash,
			rec.Name,
			rec.Value,
			rec.RegistrationDate,
		)
	} else {
		query = "INSERT INTO ens_hash_inval(" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,reg_date,tx_hash,hash,name,value" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),TO_TIMESTAMP($6),$7,$8,$9,$10)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			rec.RegistrationDate,
			rec.TxHash,
			rec.Hash,
			rec.Name,
			rec.Value,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_hash_registered(rec *p.ENS_HashRegistered) {

	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	owner_aid := ss.Lookup_or_create_address(rec.Owner,rec.BlockNum,rec.TxId)
	var query string
	var err error
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_hash_reg("+
					"tx_hash,time_stamp,block_num,contract_aid,hash,owner_id,value,reg_date" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7::DECIMAL/1e+18,TO_TIMESTAMP($8))"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			rec.Hash,
			owner_aid,
			rec.Value,
			rec.RegistrationDate,
		)
	} else {
		query = "INSERT INTO ens_hash_reg(" +
					"evtlog_id,block_num,tx_id,contract_aid,owner_aid,"+
					"time_stamp,reg_date,tx_hash,hash,value" +
				") VALUES($1,$2,$3,$4,$5,TO_TIMESTAMP($6),TO_TIMESTAMP($7),$8,$9,$10)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			owner_aid,
			rec.TimeStamp,
			rec.RegistrationDate,
			rec.TxHash,
			rec.Hash,
			rec.Value,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_new_resolver(rec *p.ENS_NewResolver) {

	aid := ss.Lookup_or_create_address(rec.Address,rec.BlockNum,rec.TxId)
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	var query string
	var err error
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_new_resolver(tx_hash,time_stamp,block_num,contract_aid,node,aid) " +
		"VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6)"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			rec.Node,
			aid,
		)
	} else {
		query = "INSERT INTO ens_new_resolver (" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,aid,tx_hash,node" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			aid,
			rec.TxHash,
			rec.Node,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_all_ens_labels_from_owners(output *map[string]struct{}) {

	var query string
	query = "SELECT DISTINCT label AS label FROM ens_new_owner ORDER BY label"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var label string
		err=rows.Scan(&label)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		(*output)[label]=struct{}{}
	}
}
func (ss *SQLStorage) Insert_word_for_ens_label(word string,label string) {

	var query string
	query = "INSERT INTO ens_label(word,label) VALUES($1,$2) ON CONFLICT DO NOTHING"
	_,err := ss.db.Exec(query,word,label)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v with label %v for word %v ; q=%v",err,label,word,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Label_exists_in_ens_labels(label string) bool {

	var query string
	query = "SELECT label FROM ens_label WHERE label=$1"
	res := ss.db.QueryRow(query,label)
	var null_label sql.NullString
	err := res.Scan(&null_label)
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
func (ss *SQLStorage) Reverse_lookup_registration_exists(address string,node string) bool {
	// Used by the process that scans reverse names, to make sure the label we are about
	//	to insert in the 'ens_labels' is the correct one

//	aid := ss.Lookup_address_id(address)
	var query string
	query = "SELECT id FROM ens_new_owner WHERE fqdn=$1"
	res := ss.db.QueryRow(query,node)
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

func (ss *SQLStorage) Insert_registry_transfer(rec *p.ENS_RegistryTransfer) {

	aid := ss.Lookup_or_create_address(rec.Owner,rec.BlockNum,rec.TxId)
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	var query string
	var err error
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_reg_transf(tx_hash,time_stamp,block_num,contract_aid,node,aid) " +
		"VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6)"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			rec.Node,
			aid,
		)
	} else {
		query = "INSERT INTO ens_reg_transf (" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,aid,tx_hash,node" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			aid,
			rec.TxHash,
			rec.Node,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_text_changed(rec *p.ENS_TextChanged) {

	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	var query string
	var err error
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_text_chg(tx_hash,time_stamp,block_num,contract_aid,node,key,value) " +
		"VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7)"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			rec.Node,
			rec.Key,
			rec.Value,
		)
	} else {
		query = "INSERT INTO ens_text_chg (" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,tx_hash,node,key,value" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			rec.TxHash,
			rec.Node,
			rec.Key,
			rec.Value,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_addrs_with_reverse_name() []string {

	var query string
	query = "SELECT DISTINCT a.addr AS addr " +
				"FROM ens_new_owner AS o " +
				"JOIN address AS a ON o.owner_aid = a.address_id "
		//		"WHERE node='91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e2'"


	rows,err := ss.db.Query(query)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB Error: %v",err))
			os.Exit(1)
		}
	}
	records := make([]string,0,256)
	defer rows.Close()
	for rows.Next() {
		var addr string
		err=rows.Scan(&addr)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,addr)
	}
	return records
}
func (ss *SQLStorage) Select_TLDs() []string {

	var query string
	query = "SELECT DISTINCT a.addr AS addr " +
				"FROM ens_new_owner AS o " +
				"JOIN address AS a ON o.owner_aid = a.address_id "
		//		"WHERE node='91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e2'"


	rows,err := ss.db.Query(query)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB Error: %v",err))
			os.Exit(1)
		}
	}
	records := make([]string,0,256)
	defer rows.Close()
	for rows.Next() {
		var addr string
		err=rows.Scan(&addr)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,addr)
	}
	return records
}
func (ss *SQLStorage) Get_node_text_key_values(node string) (string,[]p.ENS_TextKeyValue) {

	var query string

	query = "SELECT fqdn_words FROM ens_node WHERE node=$1"
	res := ss.db.QueryRow(query,node)
	var null_fqdn sql.NullString
	err := res.Scan(&null_fqdn)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			// nothing, we will return empty name
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	fqdn_words := null_fqdn.String

	query = "SELECT " +
				"kv.key," +
				"kv.value " +
			"FROM ens_text_key AS kv " +
			"WHERE kv.node=$1 " +
			"ORDER BY kv.key "

	rows,err := ss.db.Query(query,node)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.ENS_TextKeyValue,0,8)
	defer rows.Close()
	for rows.Next() {
		var rec p.ENS_TextKeyValue
		err=rows.Scan(&rec.Key,&rec.Value)
		decoded_txt,err_decode := hex.DecodeString(rec.Value)
		if err_decode != nil {
			ss.Log_msg(fmt.Sprintf(
				"Error decoding hex for key-value pairs of node %v, key %v : %v\n",
				node,rec.Key,err_decode,
			))
		}
		rec.Value = string(decoded_txt)
		records = append(records,rec)
	}
	return fqdn_words,records
}
func (ss *SQLStorage) Update_ens_proc_status(status *p.EnsProcStatus) {

	var query string
	query = "UPDATE ens_status SET last_evt_id = $1"

	_,err := ss.db.Exec(query,status.LastEvtId)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_ens_proc_status() p.EnsProcStatus {

	var output p.EnsProcStatus
	var null_id sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM ens_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO ens_status DEFAULT VALUES"
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
func (ss *SQLStorage) Get_ens_resolver(address string) (p.ENSResolver,error) {

	var output p.ENSResolver
	var query string
	query = "SELECT node,a.address_id FROM ens_new_resolver r " +
				"JOIN address a ON r.aid=a.address_id " +
				"WHERE a.addr=$1"
	res := ss.db.QueryRow(query,address)
	err := res.Scan(&output.Node,&output.Aid)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return output,nil
}
