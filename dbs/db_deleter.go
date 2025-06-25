package dbs

import (
	"os"
	"fmt"
	"database/sql"
	 "github.com/ethereum/go-ethereum/common"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_deleter_contracts() []p.Deleter_contracts{

	var query string
	query = "SELECT "+
				"a.addr," +
				"a.address_id, "+
				"s.info "+
			"FROM d_contracts AS s " +
				"JOIN address a ON a.addr=s.contract_addr "+
			"WHERE a.addr = s.contract_addr"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.Deleter_contracts,0,16)
	defer rows.Close()
	for rows.Next() {
		var rec p.Deleter_contracts
		err=rows.Scan(
			&rec.ContractAddr,
			&rec.ContractAid,
			&rec.Info,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
		rec.ContractEthAddr = common.HexToAddress(rec.ContractAddr)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_deleter_status() (int64) {

	var query string
	query="SELECT last_block_num FROM d_status";
	row := ss.db.QueryRow(query)
	var block_num int64
	var err error
	err=row.Scan(&block_num);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return -1
		}
		ss.Log_msg(fmt.Sprintf("Error in Get_deleter_status(): %v",err))
		os.Exit(1)
	}
	return block_num
}

func (ss *SQLStorage) Get_deleter_count_non_deleteable_transactions_by_tx_to(block_num int64,contracts string) (int64) {

	var query string
	query="SELECT COUNT(*) FROM transaction t WHERE (block_num=$1) AND (t.to_aid IN ("+contracts+"))";
	row := ss.db.QueryRow(query,block_num)
	var count int64
	var err error
	err=row.Scan(&count);
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Error in Get_deleter_count_non_deleteable_by_tx_to(): %v",err))
		os.Exit(1)
	}
	return count
}
func (ss *SQLStorage) Get_deleter_count_non_deleteable_transactions_by_events_emitted(block_num int64,contracts string) (int64) {

	var query string
	query="SELECT "+
				"COUNT(*) "+
			"FROM transaction t "+
				"INNER JOIN evt_log e ON e.tx_id=t.id "+
			"WHERE (t.block_num=$1) AND (e.contract_aid IN ("+contracts+"))";
	row := ss.db.QueryRow(query,block_num)
	var count int64
	var err error
	err=row.Scan(&count);
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Error in Get_deleter_count_non_deleteable_by_events_emitted(): %v",err))
		os.Exit(1)
	}
	return count
}
func (ss *SQLStorage) Update_deleter_status(block_num int64) {

	var query string = "UPDATE d_status SET last_block_num=$1"
	res,err:=ss.db.Exec(query,block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Update_deleter_status() failed: %v",err))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in Update_deleter_status(): %v",err))
		os.Exit(1)
	}
	if affected_rows>0 {
		// break
	} else {
		query = "INSERT INTO d_status(last_block_num)  VALUES($1)"
		_,err := ss.db.Exec(query,block_num)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("Update_deleter_status() failed on INSERT: %v",err));
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Deleter_do_delete_block_transactions(block_num int64) {

	// deletes transactions and all the dependent tables receieve cascaded DELETEs also
	var query string
	query = "DELETE FROM transaction WHERE block_num = $1"
	_,err := ss.db.Exec(query,block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (block_num=%v, %v) in Deleter_do_delete_block_transactions()",err,block_num,query))
		os.Exit(1)
	}
}
