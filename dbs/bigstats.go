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
func (ss *SQLStorage) Set_schema_name(name string) {
	ss.schema_name = name
}
func (ss *SQLStorage) Bigstats_get_stored_chain_id() int64 {

	var query string
	query = "SELECT chain_id FROM "+ss.schema_name+".bs_config"
	row := ss.db.QueryRow(query)
	var null_chain_id sql.NullInt64
	var err error
	err=row.Scan(&null_chain_id);
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Error in Get_stored_chain_id(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return null_chain_id.Int64
}
func (ss *SQLStorage) Bigstats_set_chain_id(chain_id int64) {

	var query string = "UPDATE "+ss.schema_name+".bs_config SET chain_id=$1"
	_,err:=ss.db.Exec(query,chain_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Set_chain_id() failed: %v",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Bigstats_get_last_block_num() (int64,bool) {

	var query string
	query="SELECT last_block FROM "+ss.schema_name+".bs_config LIMIT 1";
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
		return null_block_num.Int64,true
	} else {
		return -1,false
	}
}
func (ss *SQLStorage) Bigstats_set_last_block_num(block_num int64) {

	bnum := int64(block_num)
	var query string = "UPDATE "+ss.schema_name+".bs_config SET last_block=$1"
	res,err:=ss.db.Exec(query,bnum)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("set_last_block_num() failed: %v",err))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	ss.Info.Printf("update_last_block: affected_rows=%v, err=%v\n",affected_rows,err)
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
func (ss *SQLStorage) Bigstats_get_block_num_by_hash(block_hash string) (int64,error) {

	var query string
	query = "SELECT block_num FROM "+ss.schema_name+".bs_block WHERE block_hash=$1"

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
func (ss *SQLStorage) Bigstats_chainsplit_delete_blocks(starting_block_num int64) {

	var err error
	var query string
	// Note: We must delete in reverse order of block creation because the triggers
	//			in the DB have made cumulative operations
	query = "DELETE FROM "+ss.schema_name+".bs_block WHERE block_num IN (" +
				"SELECT block_num FROM "+ss.schema_name+".block WHERE block_num>$1 ORDER BY block_num DESC" +
			")"
	_,err = ss.db.Exec(query,starting_block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v, block_num=%v",err,query,starting_block_num))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Bigstats_block_delete_with_everything(block_num int64) {

	// deletes block table and all the dependent tables receieve cascaded DELETEs also
	var query string
	query = "DELETE FROM "+ss.schema_name+".bs_block WHERE block_num = $1"
	_,err := ss.db.Exec(query,block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (block_num=%v, %v)",err,block_num,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Bigstats_insert_block(hash_str string,block *types.Header,num_tx int,no_chainsplit_check bool) error {

	var query string
	var parent_block_num int64
	parent_hash := block.ParentHash.String()

	query="SELECT block_num FROM "+ss.schema_name+".bs_block WHERE block_hash=$1"
	err:=ss.db.QueryRow(query,parent_hash).Scan(&parent_block_num);
	if no_chainsplit_check {
		err = nil // clear error as we don't need to validate the chain
		parent_block_num = block.Number.Int64()-1
	}
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			query = "SELECT count(*) FROM "+ss.schema_name+".bs_block"
			row := ss.db.QueryRow(query)
			var block_count int64
			err := row.Scan(&block_count)
			if (err!=nil) {
				ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
				os.Exit(1)
			}
			if block_count > 0 {
				starting_block:=ss.Get_upload_block()
				if block.Number.Int64() == starting_block {
					// this is the first block that will be processed (we aren't starting from block 0)
					// allow
				} else {
					ss.Info.Printf(
						fmt.Sprintf(
							"Insert_block() Can't insert block (block_num=%v, block_hash=%v, parent_hash=%v"+
							"), parent not found. Chain split, need recovery procedure. (CHAIN_SPLIT)",
							block.Number.Int64(),hash_str,parent_hash,
						),
					);
					return p.ErrChainSplit // chain split occured (parent block wasn't found)
				}
			} else {
				// database is empty, continue
			}
		} else {
			ss.Log_msg(fmt.Sprintf("DB Error: %v; query=%v",err,query));
			os.Exit(1)
		}
	} else {
		if (parent_block_num + 1) != block.Number.Int64() {
			ss.Info.Printf(
				fmt.Sprintf(
					"Insert_block() Can't insert block (block_num=%v, block_hash=%v, parent_hash=%v"+
					"), block found as parent has non-consecutive number (parent_block_num=%v). " +
					"Chain split, need recovery procedure. (CHAIN_SPLIT)",
					parent_block_num,block.Number.Int64(),hash_str,parent_hash,
				),
			);
			return p.ErrChainSplit // chain split occurred (parent's block num isn't consecutive)
		}
	}

	block_num := int64(block.Number.Uint64())
	query = `
		INSERT INTO `+ss.schema_name+`.bs_block(
			block_num,
			block_hash,
			ts,
			parent_hash,
			num_tx
		) VALUES ($1,$2,TO_TIMESTAMP($3),$4,$5)`

	result,err := ss.db.Exec(query,
			block_num,
			hash_str,
			block.Time,
			parent_hash,
			num_tx,
	)
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't insert into block table block_num=%v: %v, q=%v",
				block.Number.Int64(),err,query,
			),
		)
		os.Exit(1)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected > 0 {
		return nil
	}
	ss.Log_msg(fmt.Sprintf("DB error: couldn't insert into block table. Rows affeced = 0"))
	os.Exit(1)
	return nil
}
func (ss *SQLStorage) Update_block_eth_transferred(block_num int64,total_eth *big.Int) {

	var query string
	query = "UPDATE "+ss.schema_name+".bs_block SET total_eth=$2 WHERE block_num=$1"
	_,err := ss.db.Exec(query,block_num,total_eth.String())
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't update address metadata for block %v : %v: q=%v",
				block_num,err,query,
			),
		)
		os.Exit(1)
	}
}
func (ss *SQLStorage) Bigstats_lookup_or_create_address(addr string,is_contract bool) int64 {

	var aid int64
	var query string
	query="SELECT address_id FROM "+ss.schema_name+".bs_addr WHERE addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&aid);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			aid = ss.Bigstats_insert_address(addr,is_contract)
			return aid
		} else {
			ss.Log_msg(fmt.Sprintf("DB error in getting address id : %v",err))
		}
	}
	return aid
}
func (ss *SQLStorage) Bigstats_insert_address(addr string,is_contract bool) int64 {

	var addr_id int64;
	var query string
	if len(addr) == 0 {
		ss.Log_msg(fmt.Sprintf("Attempt to insert address with len=0"))
		os.Exit(1)
	}
	query = "INSERT INTO "+ss.schema_name+".bs_addr(addr,is_contract) "+
				"VALUES($1,$2) RETURNING address_id"
	row:=ss.db.QueryRow(query,addr,is_contract);
	err:=row.Scan(&addr_id)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error in address insertion: %v : %v",query,err))
		os.Exit(1)
	}
	if addr_id==0 {
		ss.Log_msg(fmt.Sprintf("DB error, addr_id after INSERT is 0"))
		os.Exit(1)
	}

	return addr_id
}
func (ss *SQLStorage) Bigstats_lookup_address_id(addr string) (int64,bool,error) {

	var query string
	var aid int64
	var is_contract bool
	query="SELECT address_id,is_contract FROM "+ss.schema_name+".bs_addr WHERE addr=$1"
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
func (ss *SQLStorage) Insert_all_addr_stat_logs(entries []p.AddrStatsLog) {

	if len(entries) == 0 {
		ss.Log_msg(fmt.Sprintf("Insert of address stats into log with empty array\n"))
		os.Exit(1)
	}
	var query strings.Builder
	query.WriteString("INSERT INTO "+ss.schema_name+".bs_log(block_num,tx_index,aid) VALUES")

	for i,eel := range(entries) {
		if i > 0 {
			query.WriteString(",")
		}
		query.WriteString(fmt.Sprintf(
			"(%v,%v,%v)",
			eel.BlockNum,
			eel.TxIndex,
			eel.Aid,
		))
	}
	query.WriteString(" ON CONFLICT DO NOTHING")
	var err error
	_,err = ss.db.Exec(query.String())
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Insert_all_addr_stat_logs() failed: %v, q=%v\n",err,query.String()))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_unique_accounts_counter_by_type(ts,duration int64,is_contract bool) {

	var query string
	query = "WITH data AS ("+
					"SELECT " +
						"DISTINCT aid aid "+
					"FROM bs_log log "+
						"JOIN bs_block b ON log.block_num=bs.block_num " +
						"JOIN bs_addr a ON log.aid=a.address_id "+
					"WHERE (b.ts <= TO_TIMESTAMP($1)) AND (b.ts < TO_TIMESTAMP($2)" +
						"AND (a.is_contract = $3)"+
				")"+
			"SELECT COUNT(aid) FROM data"

	ts_end := ts + duration
	var num_rows int64
	err:=ss.db.QueryRow(query,ts,ts_end,is_contract).Scan(&num_rows);
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error in getting address id : %v",err))
		os.Exit(1)
	}
	return num_rows
}
func (ss *SQLStorage) Close_period(ts,duration int64) {


	human_account_count := ss.Get_unique_accounts_counter_by_type(ts,duration,false)
	contract_account_count := ss.Get_unique_accounts_counter_by_type(ts,duration,true)

	var query string
	query = "INSERT INTO bs_period("+
					"time_stamp,unique_addrs_eoa,unique_addrs_code,eth_transferred"+
			") VALUES ($1,$2,$3,$4)"+

	result,err := ss.db.Exec(query,
			block_num,
			hash_str,
			block.Time,
			parent_hash,
			num_tx,
	)
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't insert into bs_period table: %v  q=%v",
				err,query,
			),
		)
		os.Exit(1)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
	}
	if rows_affected == 0 {
		ss.Log_msg(fmt.Sprintf("Couldnt insert record in bs_period table"))
		os.Exit(1)
	}
}
