package dbs

import (
	"fmt"
	"os"

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
	if rec.EvtId == 0 {	// initial load, we don't have the Block in 'block' table
		query = "INSERT INTO ens_name(owner_aid,label,cost) " +
				"VALUES($1,$2,$3::DECIMAL/1e+18)"
		_,err = ss.db.Exec(query,owner_aid,rec.Label,rec.Cost)
	} else {
		query = "INSERT INTO ens_name(evtlog_id,block_num,tx_id,time_stamp,owner_aid,label,cost) " +
				"VALUES($1,$2,$3,$4,$5,$6,$7::DECIMAL/1e+18)"
		_,err = ss.db.Exec(query,
			rec.EvtId,
			rec.BlockNum,
			rec.TxId,
			rec.TimeStamp,
			owner_aid,
			rec.Label,
			rec.Cost,
		)
	}
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
