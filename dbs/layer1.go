package dbs

import (
	"fmt"
	"os"
	"strings"
	"math/big"

	"database/sql"
	_  "github.com/lib/pq"

//	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Layer1_lookup_address_id(addr string) (int64,bool,error) {

	var query string
	var aid int64
	var is_contract bool
	query="SELECT address_id,is_contract FROM "+ss.schema_name+".addr WHERE addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&aid,&is_contract);
	if (err!=nil) {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v ,q=%v",query))
			os.Exit(1)
		}
		return 0,false,err
	}
	return aid,is_contract,nil
}
func (ss *SQLStorage) Layer1_insert_address(addr string) int64 {

	var addr_id int64;
	var query string
	if len(addr) == 0 {
		ss.Log_msg(fmt.Sprintf("Attempt to insert address with len=0"))
		os.Exit(1)
	}
	query = "INSERT INTO "+ss.schema_name+".addr(addr) "+
				"VALUES($1) RETURNING address_id"
	row:=ss.db.QueryRow(query,addr);
	err:=row.Scan(&addr_id)
	if err!=nil {
		if strings.Contains(err.Error(),"duplicate key value") {
			query="SELECT address_id FROM "+ss.schema_name+".bs_addr WHERE addr=$1"
			err:=ss.db.QueryRow(query,addr).Scan(&addr_id);
			if (err!=nil) {
				ss.Log_msg(fmt.Sprintf("DB error in address insertion on second attempt: %v : %v",query,err))
				os.Exit(1)
			}
		} else {
			ss.Log_msg(fmt.Sprintf("DB error in address insertion: %v : %v",query,err))
			os.Exit(1)
		}
	}
	if addr_id==0 {
		ss.Log_msg(fmt.Sprintf("DB error, addr_id after INSERT is 0"))
		os.Exit(1)
	}

	return addr_id
}
func (ss *SQLStorage) Layer1_chainsplit_delete_blocks(starting_block_num int64) {

	var err error
	var query string
	// Note: We must delete in reverse order of block creation because the triggers
	//			in the DB have made cumulative operations
	query = "DELETE FROM "+ss.schema_name+".block WHERE block_num IN (" +
				"SELECT block_num FROM "+ss.schema_name+".block WHERE block_num>$1 ORDER BY block_num DESC" +
			")"
	_,err = ss.db.Exec(query,starting_block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v, block_num=%v",err,query,starting_block_num))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Layer1_set_last_block_num(block_num int64) {

	bnum := int64(block_num)
	var query string = "UPDATE "+ss.schema_name+".config SET last_block=$1"
	res,err:=ss.db.Exec(query,bnum)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("set_last_block_num() failed: %v",err))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in set_last_block(): %v",err))
		os.Exit(1)
	}
	if affected_rows>0 {
		// break
	} else {
		ss.Log_msg(fmt.Sprintf("set_last_block_num() no default record in bs_config: %v",err));
		os.Exit(1)
	}
}
func (ss *SQLStorage) Layer1_get_block_num_by_hash(block_hash string) (int64,error) {

	var query string
	query = "SELECT block_num FROM "+ss.schema_name+".block WHERE block_hash=$1"

	row := ss.db.QueryRow(query,block_hash)
	var block_num int64
	var err error
	err=row.Scan(&block_num);
	if (err!=nil) {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error in Get_block_num_by_hash(): %v, q=%v,h=%v",err,query,block_hash))
			os.Exit(1)
		}
		return 0,err
	}
	return block_num,nil
}
