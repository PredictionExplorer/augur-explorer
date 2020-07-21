// Data Base Storage
package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	p "augur-extractor/primitives"
)
func (ss *SQLStorage) Get_last_block_num() (p.BlockNumber,bool) {

	var query string
	query="SELECT block_num FROM last_block LIMIT 1";
	row := ss.db.QueryRow(query)
	var null_block_num sql.NullInt64
	var err error
	err=row.Scan(&null_block_num);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return -1,false
		} else {
			ss.Log_msg(fmt.Sprintf("Error in get_last_block_num(): %v",err))
			os.Exit(1)
		}
	}
	if (null_block_num.Valid) {
		return p.BlockNumber(null_block_num.Int64),true
	} else {
		return -1,false
	}
}
func (ss *SQLStorage) Set_last_block_num(block_num p.BlockNumber) {

	bnum := int64(block_num)
	var query string = "UPDATE last_block SET block_num=$1 WHERE block_num < $1"
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
		query = "INSERT INTO last_block VALUES($1)"
		_,err := ss.db.Exec(query,bnum)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("set_last_block_num() failed on INSERT: %v",err));
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Insert_block(hash_str string,block *types.Header)  bool {

	var query string
	var parent_block_num int64
	parent_hash := block.ParentHash.String()

	query="SELECT block_num,parent_hash FROM block WHERE hash=$1"
	err:=ss.db.QueryRow(query,parent_hash).Scan(&parent_block_num);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			if block.Number.Uint64() == 0 {
				// Genesis. Allow.
			} else {
				if (parent_block_num + 1) != int64(block.Number.Uint64()) {
					return false
				}
			}
		}
	}

	block_num := int64(block.Number.Uint64())
	query = `
		INSERT INTO block(
			block_num,
			block_hash,
			ts,
			parent_hash
		) VALUES ($1,$2,TO_TIMESTAMP($3),$4)`

	result,err := ss.db.Exec(query,
			block_num,
			hash_str,
			block.Time,
			parent_hash)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into block  table: %v, q=%v",err,query))
		os.Exit(1)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected > 0 {
		return true
	}
	ss.Log_msg(fmt.Sprintf("DB error: couldn't insert into block table. Rows affeced = 0"))
	return false
}
func (ss *SQLStorage) Insert_transaction(block_num p.BlockNumber,tx_hash string,tx *types.Message) int64 {

	var query string
	var tx_id int64


	query = "INSERT INTO transaction (block_num,value,tx_hash) " +
			"VALUES ($1,("+tx.Value().String()+"/1e+18),$2) RETURNING id"

	row := ss.db.QueryRow(query,block_num,tx_hash)
	err := row.Scan(&tx_id)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into transactions table: %v, q=%v",err,query))
		os.Exit(1)
	}

	from_aid := ss.Lookup_or_create_address(tx.From().String(),block_num,tx_id)
	var to_aid int64 = 0
	if tx.To() == nil {	// case for calling contract creation
		zero_addr := common.BigToAddress(zero)
		to_aid = ss.Lookup_or_create_address(zero_addr.String(),block_num,tx_id)
	} else {
		to_aid = ss.Lookup_or_create_address(tx.To().String(),block_num,tx_id)
	}
	query = "UPDATE transaction set from_aid=$2 , to_aid=$3 where id = $1"
	_,err = ss.db.Exec(query,tx_id,from_aid,to_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
		os.Exit(1)
	}

	return tx_id
}
func (ss *SQLStorage) Fix_chainsplit(block *types.Header) p.BlockNumber {

	var query string
	var my_block_num int64
	parent_hash := block.ParentHash.String()
	query = "SELECT block_num FROM block WHERE block_hash = $1"
	row := ss.db.QueryRow(query,parent_hash)
	err := row.Scan(&my_block_num);
	if (err!=nil) {
		if err==sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Chainsplit detected, I don't have the parent hash %v, exiting. ",parent_hash))
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	cur_block_num := int64(block.Number.Uint64())
	if cur_block_num > (my_block_num + p.MAX_BLOCKS_CHAIN_SPLIT) {
		ss.Log_msg(fmt.Sprintf("Chainsplit detected, and it is more than %v blocks, aborting.",p.MAX_BLOCKS_CHAIN_SPLIT))
	}
	query = "DELETE FROM block WHERE block_num > $1 CASCADE"
	_,err = ss.db.Exec(query,my_block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v, block_num=%v",err,query,my_block_num))
		os.Exit(1)
	}
	return p.BlockNumber(my_block_num + 1)	// parent + 1 = current
}
func (ss *SQLStorage) Block_delete_with_everything(block_num p.BlockNumber) {

	// deletes block table and all the other tables receieve cascaded DELETEs also
	var query string
	query = "DELETE FROM block WHERE block_num = $1"
	_,err := ss.db.Exec(query,block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (block_num=%v, %v)",err,block_num,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_block_info(block_num p.BlockNumber) (p.BlockInfo,error) {

	var binfo p.BlockInfo
	records_market := make([]string,0,8)
	records_addresses := make([]string,0,8)
	records_transactions := make([]string,0,8)

	var query string
	query = "SELECT block_num,num_tx FROM block WHERE block_num = $1"

	row := ss.db.QueryRow(query,block_num)
	var null_bnum sql.NullInt64
	var null_num_tx sql.NullInt64
	var err error
	err=row.Scan(&null_bnum,&null_num_tx);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return binfo,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	// get TRANSACTIONS
	query = "SELECT tx_hash FROM transaction WHERE block_num = $1"

	var rows *sql.Rows
	rows,err = ss.db.Query(query,block_num)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return binfo,err
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var tx_hash string
		err=rows.Scan(&tx_hash)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records_transactions = append(records_transactions,tx_hash)
	}
	binfo.Transactions = records_transactions

	// get MARKETS
	query = "SELECT a.addr,u.addr FROM market m " +
			"LEFT JOIN address a ON m.market_aid=a.address_id " +
			"LEFT JOIN address u ON m.eoa_aid=u.address_id " +
			"WHERE m.block_num = $1"

	rows,err = ss.db.Query(query,block_num)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return binfo,err
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var market_addr string
		var creator_addr string
		err=rows.Scan(&market_addr,&creator_addr)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records_market = append(records_market,market_addr)
		records_addresses = append(records_addresses,creator_addr)
	}
	binfo.Markets = records_market

	// get Active addresses
	query = "SELECT DISTINCT addr FROM " +
			"(" +
				"(" +
					"SELECT addr FROM mktord,address " +
					"WHERE mktord.eoa_aid = address.address_id AND mktord.block_num=$1" +
				")" +
				" UNION ALL "+
				"(" +
					"SELECT addr FROM mktord,address " +
					"WHERE mktord.eoa_fill_aid = address.address_id AND mktord.block_num=$1" +
				")" +
			") AS records"

	rows,err = ss.db.Query(query,block_num)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return binfo,err
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var active_addr string
		err=rows.Scan(&active_addr)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records_addresses = append(records_addresses,active_addr)
	}
	binfo.Addresses= records_addresses

	binfo.BlockNum = block_num
	binfo.NumTx=int64(len(binfo.Transactions))
	binfo.NumAddresses=int64(len(binfo.Addresses))
	binfo.NumMarkets=int64(len(binfo.Markets))

	return binfo,nil
}
func (ss *SQLStorage) Get_transaction(tx_hash string) (p.TxInfo,error) {

	var ti p.TxInfo
	ti.Hash = tx_hash
	var query string
	query = "SELECT " +
				"t.block_num," +
				"sa.addr," +
				"ra.addr," +
				"t.value " +
			"FROM transaction t " +
				"LEFT JOIN address sa ON t.from_aid = sa.address_id " +
				"LEFT JOIN address ra ON t.to_aid = ra.address_id " +
			"WHERE t.tx_hash=$1"

	row := ss.db.QueryRow(query,tx_hash)
	err := row.Scan(
				&ti.BlockNum,
				&ti.From,
				&ti.To,
				&ti.Value,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return ti,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return ti,err
}
func (ss *SQLStorage) Get_last_block_timestamp() int64 {

	var query string
	query = "SELECT FLOOR(EXTRACT(EPOCH FROM block.ts))::BIGINT AS ts " +
//	EXTRACT(EPOCH FROM block.ts)::BIGINT AS ts "+
			"FROM block,last_block WHERE last_block.block_num=block.block_num"
	row := ss.db.QueryRow(query)
	var ts int64
	var err error
	err=row.Scan(&ts);
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Error in Get_last_block_timestamp(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return ts
}
func (ss *SQLStorage) Get_first_block_timestamp() int64 {

	var query string
	query = "SELECT FLOOR(EXTRACT(EPOCH FROM block.ts))::BIGINT AS ts " +
			"FROM block ORDER BY block_num LIMIT 1"
	row := ss.db.QueryRow(query)
	var ts int64
	var err error
	err=row.Scan(&ts);
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Error in Get_last_block_timestamp(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return ts
}
