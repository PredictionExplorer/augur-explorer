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
		query = "INSERT INTO ens_name_reg1(" +
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
		query = "INSERT INTO ens_name_reg1(" +
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
func (ss *SQLStorage) Insert_name_registered2(rec *p.ENS_Name2) {

	var query string
	var err error
	owner_aid := ss.Lookup_or_create_address(rec.Owner,rec.BlockNum,rec.TxId)
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_name_reg2(" +
					"tx_hash,time_stamp,block_num,contract_aid,owner_aid,node,label,fqdn,expires" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,TO_TIMESTAMP($9))"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			owner_aid,
			rec.Node,
			rec.Label,
			rec.FQDN,
			rec.Expires,
		)
	} else {
		query = "INSERT INTO ens_name_reg2(" +
					"evtlog_id,block_num,tx_id,contract_aid,owner_aid,"+
					"time_stamp,node,label,fqdn,tx_hash,expires" +
				") VALUES(" +
					"$1,$2,$3,$4,$5,TO_TIMESTAMP($6),$7,$8,$9,$10,TO_TIMESTAMP($11)"+
				")"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			owner_aid,
			rec.TimeStamp,
			rec.Node,
			rec.Label,
			rec.FQDN,
			rec.TxHash,
			rec.Expires,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_name_registered3(rec *p.ENS_Name3) {

	var query string
	var err error
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	beneficiary_aid := ss.Lookup_or_create_address(rec.Beneficiary,rec.BlockNum,rec.TxId)
	caller_aid := ss.Lookup_or_create_address(rec.Caller,rec.BlockNum,rec.TxId)
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_name_reg3(" +
					"tx_hash,time_stamp,block_num,contract_aid,caller_aid,beneficiary_aid," +
					"subdomain,node,label,fqdn,created_date" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10,TO_TIMESTAMP($11))"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			caller_aid,
			beneficiary_aid,
			rec.Subdomain,
			rec.Node,
			rec.Label,
			rec.FQDN,
			rec.CreatedDate,
		)
	} else {
		query = "INSERT INTO ens_name_reg3(" +
					"evtlog_id,time_stamp,block_num,tx_id,contract_aid,caller_aid,"+
					"beneficiary_aid,subdomain,node,label,fqdn,tx_hash,created_date" +
				") VALUES(" +
					"$1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,TO_TIMESTAMP($13)"+
				")"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.TimeStamp,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			caller_aid,
			beneficiary_aid,
			rec.Subdomain,
			rec.Node,
			rec.Label,
			rec.FQDN,
			rec.TxHash,
			rec.CreatedDate,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_name_renewed(rec *p.ENS_NameRenewed) {

	var query string
	var err error
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_name_renewed(" +
					"tx_hash,time_stamp,block_num,contract_aid," +
					"label,node,fqdn,name,cost,expires" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9::DECIMAL/1e+18,TO_TIMESTAMP($10))"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			rec.Label,
			rec.Node,
			rec.FQDN,
			rec.Name,
			rec.Cost,
			rec.Expires,
		)
	} else {
		query = "INSERT INTO ens_name_renewed(" +
					"evtlog_id,block_num,tx_id,contract_aid,"+
					"time_stamp,label,node,fqdn,name,tx_hash,cost,expires" +
				") VALUES(" +
					"$1,$2,$3,$4,$5,TO_TIMESTAMP($6),$7,$8,$9,$10,"+
					"$11::DECIMAL/1e+18,TO_TIMESTAMP($12)"+
				")"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
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
				"label,COALESCE(name,'') " +
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
func (ss *SQLStorage) Insert_address_changed1(rec *p.ENS_AddrChanged) {

	var query string
	var err error
	aid := ss.Lookup_or_create_address(rec.Address,rec.BlockNum,rec.TxId)
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_addr1(" +
					"tx_hash,time_stamp,block_num,contract_aid,aid,fqdn" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6)"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			aid,
			rec.FQDN,
		)
	} else {
		query = "INSERT INTO ens_addr1(" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,aid,tx_hash,fqdn" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			aid,
			rec.TxHash,
			rec.FQDN,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_address_changed2(rec *p.ENS_AddressChanged) {

	var query string
	var err error
	aid := ss.Lookup_or_create_address(rec.Address,rec.BlockNum,rec.TxId)
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_addr2(" +
					"tx_hash,time_stamp,block_num,contract_aid,aid,fqdn,coin_type" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7)"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			aid,
			rec.FQDN,
			rec.CoinType,
		)
	} else {
		query = "INSERT INTO ens_addr2(" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,aid,tx_hash,fqdn,coin_type" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			aid,
			rec.TxHash,
			rec.FQDN,
			rec.CoinType,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_name_changed(rec *p.ENS_NameChanged) {

	var query string
	var err error
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_rev_name(" +
					"tx_hash,time_stamp,block_num,contract_aid,node,name" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6)"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			rec.Node,
			rec.Name,
		)
	} else {
		query = "INSERT INTO ens_rev_name(" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,tx_hash,node,name" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			rec.TxHash,
			rec.Node,
			rec.Name,
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
	name_aid := ss.Lookup_or_create_address(rec.NameAddr,rec.BlockNum,rec.TxId)
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	var query string
	var err error
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_new_resolver(tx_hash,time_stamp,block_num,contract_aid,node,aid,name_aid) " +
		"VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7)"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			rec.Node,
			aid,
			name_aid,
		)
	} else {
		query = "INSERT INTO ens_new_resolver (" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,aid,name_aid,tx_hash,node" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			aid,
			name_aid,
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
func (ss *SQLStorage) Insert_registrar_transfer(rec *p.ENS_RegistrarTransfer) {

	from_aid := ss.Lookup_or_create_address(rec.From,rec.BlockNum,rec.TxId)
	to_aid := ss.Lookup_or_create_address(rec.To,rec.BlockNum,rec.TxId)
	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	var query string
	var err error
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_rstr_transf(" +
					"tx_hash,time_stamp,block_num,contract_aid,label,node,fqdn,from_aid,to_aid" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8,$9)"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			rec.Label,
			rec.Node,
			rec.FQDN,
			from_aid,
			to_aid,
		)
	} else {
		query = "INSERT INTO ens_reg_transf (" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,from_aid,to_aid,tx_hash,label,node,fqdn" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9,$10,$11)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			from_aid,
			to_aid,
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
				"WHERE a.addr=$1 " +
				"ORDER BY r.time_stamp DESC LIMIT 1"
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
func (ss *SQLStorage) Insert_pubkey_changed(rec *p.ENS_PubkeyChanged) {

	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	var query string
	var err error
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_pkey("+
					"tx_hash,time_stamp,block_num,contract_aid,node,x,y,derived_addr" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6,$7,$8)"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			rec.Node,
			rec.X,
			rec.Y,
			rec.DerivedAddr,
		)
	} else {
		query = "INSERT INTO ens_pkey(" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,tx_hash,node,x,y,derived_addr" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9,$10)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			rec.TxHash,
			rec.Node,
			rec.X,
			rec.Y,
			rec.DerivedAddr,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_contenthash_changed(rec *p.ENS_ContentHashChanged) {

	contract_aid := ss.Lookup_or_create_address(rec.Contract,rec.BlockNum,rec.TxId)
	var query string
	var err error
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_hash("+
					"tx_hash,time_stamp,block_num,contract_aid,node,hash" +
				") VALUES($1,TO_TIMESTAMP($2),$3,$4,$5,$6)"
		_,err = ss.db.Exec(query,
			rec.TxHash,
			rec.TimeStamp,
			rec.BlockNum,
			contract_aid,
			rec.Node,
			rec.Hash,
		)
	} else {
		query = "INSERT INTO ens_hash(" +
					"evtlog_id,block_num,tx_id,contract_aid,time_stamp,tx_hash,node,hash" +
				") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			contract_aid,
			rec.TimeStamp,
			rec.TxHash,
			rec.Node,
			rec.Hash,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_ens_node_addresses(node string) []p.ENS_NodeAddr{

	var query string
	query = "SELECT DISTINCT aid FROM ens_addr1 a1 WHERE a1.fqdn=$1"
	rows,err := ss.db.Query(query,node)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	records := make([]p.ENS_NodeAddr,0,4)
	defer rows.Close()
	for rows.Next() {
		var aid int64
		err=rows.Scan(&aid)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		query = "SELECT " +
					"EXTRACT(EPOCH FROM ea.time_stamp)::BIGINT AS ts," +
					"ea.time_stamp," +
					"ea.block_num," +
					"ea.tx_hash, " +
					"a.addr " +
				"FROM ens_addr1 ea " +
					"JOIN address a ON aid=address_id " +
				"WHERE ea.fqdn=$1 AND ea.aid=$2 " +
				"ORDER BY ts DESC " +
				"LIMIT 1"
		subrows,err := ss.db.Query(query,node,aid)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		defer subrows.Close()
		for subrows.Next() {
			var rec p.ENS_NodeAddr
			err=subrows.Scan(
				&rec.AddressSetTs,
				&rec.AddressSetDate,
				&rec.BlockNum,
				&rec.TxHash,
				&rec.Address,
			)
			if (err!=nil) {
				ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
				os.Exit(1)
			}
			if rec.Address == "0x0000000000000000000000000000000000000000" {
			}
			records = append(records,rec)
		}
	}
	return records
}
func (ss *SQLStorage) Get_node_lot(start_id,limit int64) []p.ENS_NodeShort{

	records := make([]p.ENS_NodeShort,0,32)
	var query string
	query = "SELECT id,node,label,fqdn,fqdn_words FROM ens_node " +
			"WHERE id>$1 ORDER BY id LIMIT $2"

	rows,err := ss.db.Query(query,start_id,limit)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	defer rows.Close()
	for rows.Next() {
		var rec p.ENS_NodeShort
		err=rows.Scan(&rec.Id,&rec.Node,&rec.Label,&rec.FQDN,&rec.FQDN_Words)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}

	return records
}
func (ss *SQLStorage)  Get_last_owner_addr(fqdn string) (string,string,error) {

	var addr_owner,addr_assigned sql.NullString
	var query string
	query =	"SELECT a.addr addr_owner " +
				"FROM ens_new_owner o " +
				"JOIN address a ON o.owner_aid=a.address_id " +
				"WHERE fqdn=$1 ORDER BY id DESC LIMIT 1"
	res := ss.db.QueryRow(query,fqdn)
	err := res.Scan(&addr_owner)
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	query = "SELECT a.addr addr_assigned "+
				"FROM ens_addr1 ac " +
				"JOIN address a ON ac.aid=a.address_id " +
				"WHERE ac.fqdn=$1 ORDER BY id DESC LIMIT 1"
	res = ss.db.QueryRow(query,fqdn)
	err = res.Scan(&addr_assigned)
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	return addr_owner.String,addr_assigned.String,nil
}
func (ss *SQLStorage) Get_last_new_resolver_name_addr(fqdn string) (string,error) {
	// returns address of the name on last resolver change
	var addr sql.NullString
	var query string
	query =	"SELECT a.addr " +
				"FROM ens_new_resolver r " +
				"JOIN address a ON r.name_aid=a.address_id " +
				"WHERE node=$1 ORDER BY id DESC LIMIT 1"
	res := ss.db.QueryRow(query,fqdn)
	err := res.Scan(&addr)
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	return addr.String,nil
}
func (ss *SQLStorage) Get_new_owner_events(fqdn string) []p.ENS_NewOwner {

	records := make([]p.ENS_NewOwner,0,24)
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT, " +
				"o.time_stamp," +
				"o.block_num,"+
				"o.label,"+
				"o.node," +
				"o.fqdn," +
				"o.tx_hash," +
				"n.fqdn_words," +
				"a.addr addr_owner " +
			"FROM ens_new_owner o " +
				"LEFT JOIN ens_node n ON o.fqdn=n.fqdn " +
				"JOIN address a ON o.owner_aid=a.address_id " +
				"WHERE o.fqdn=$1 ORDER BY o.id"

	rows,err := ss.db.Query(query,fqdn)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	defer rows.Close()
	for rows.Next() {
		var rec p.ENS_NewOwner
		var null_name sql.NullString
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.Label,
			&rec.Node,
			&rec.FQDN,
			&rec.TxHash,
			&null_name,
			&rec.Owner,
		)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_ens_record_info(fqdn string) (p.ENS_Info,error) {

	var output p.ENS_Info
	last_owner_addr,last_assigned_addr,err := ss.Get_last_owner_addr(fqdn)
	if err != nil {
		return output,err
	}
	fmt.Printf("last owner=%v, last assigned=%v\n",last_owner_addr,last_assigned_addr)
	if len(last_assigned_addr) > 0 {
		output.LastOwnerAddr = last_assigned_addr
	} else {
		output.LastOwnerAddr = last_owner_addr
	}
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM n.time_stamp)::BIGINT AS time_stamp, " +
				"n.time_stamp," +
				"n.label," +
				"n.node," +
				"n.fqdn," +
				"n.fqdn_words," +
				"txt.num_keys " +
			"FROM ens_node n " +
				"LEFT JOIN ens_text txt ON n.fqdn=txt.node " +
			"WHERE n.fqdn=$1 "
	res := ss.db.QueryRow(query,fqdn)
	var null_num_keys sql.NullInt64
	err = res.Scan(
		&output.FirstRegisteredTs,
		&output.FirstRegisteredDate,
		&output.Label,
		&output.Node,
		&output.FQDN,
		&output.ENS_Name,
		&null_num_keys,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output,nil
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	if null_num_keys.Valid {
		output.NumTextKeyValuePairs = null_num_keys.Int64
	}
	query = "SELECT hash FROM ens_hash WHERE node=$1 ORDER BY id DESC LIMIT 1"
	res = ss.db.QueryRow(query,fqdn)
	var null_content_hash sql.NullString
	err = res.Scan(&null_content_hash)
	if (err!=nil) {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	output.ContentHash = null_content_hash.String

	query = "SELECT x,y FROM ens_pkey WHERE node=$1 ORDER BY id DESC LIMIT 1"
	res = ss.db.QueryRow(query,fqdn)
	var null_pkey_x,null_pkey_y sql.NullString
	err = res.Scan(&null_pkey_x,&null_pkey_y)
	if (err!=nil) {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	output.PublicKey_X = null_pkey_x.String
	output.PublicKey_Y = null_pkey_y.String

	if len(output.ENS_Name) == 0 {
		output.ENS_Name = p.ENS_NOT_PUBLIC
	}
	output.ContentHash = null_content_hash.String
	output.AddressChangeHistory = ss.Get_ens_node_addresses(fqdn)
	var metainfo p.ENS_NodeTextInfo
	metainfo.FQDN,metainfo.TextMetaInfo = ss.Get_node_text_key_values(fqdn)
	output.NameTextMetaInfo = metainfo
	output.OwnershipChangeHistory = ss.Get_new_owner_events(fqdn)
	return output,nil
}
func (ss *SQLStorage) Lookup_ens_names(aid int64) (string,int64){

	var output string
	names,total_names:=ss.Get_user_ens_names_active(aid,0,10000000)
	for _,record := range names {
		if len(output)>0 {
			output = output + ","
		}
		if len(record.ENS_Name) > 0 {
			if record.ENS_Name != p.ENS_NOT_PUBLIC {
				output = output + record.ENS_Name
			}
		}
	}
	return output,total_names
}
