package dbs

import (
	"fmt"
	"os"
	"strings"

	"database/sql"
	_  "github.com/lib/pq"

//	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Layer1_lookup_or_insert_address_id(addr string) int64 {

	var query string
	var aid int64
	query="SELECT address_id FROM "+ss.schema_name+".addr WHERE addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&aid);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			aid = ss.Layer1_insert_address(addr)
			return aid
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v ,q=%v",query))
		os.Exit(1)
	}
	return aid
}
func (ss *SQLStorage) Layer1_lookup_address_id(addr string) (int64,error) {

	var query string
	var aid int64
	query="SELECT address_id FROM "+ss.schema_name+".addr WHERE addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&aid);
	if (err!=nil) {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v ,q=%v",query))
			os.Exit(1)
		}
		return 0,err
	}
	return aid,nil
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
func (ss *SQLStorage) Layer1_block_delete_with_everything(block_num int64) {

	// deletes block table and all the dependent tables receieve cascaded DELETEs also
	var query string
	query = "DELETE FROM "+ss.schema_name+".block WHERE block_num = $1"
	_,err := ss.db.Exec(query,block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (block_num=%v, %v)",err,block_num,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Layer1_insert_block(hash_str string,block *types.Header,num_tx int,no_chainsplit_check bool) error {

	var query string
	var parent_block_num int64
	parent_hash := block.ParentHash.String()

	query="SELECT block_num FROM "+ss.schema_name+".block WHERE block_hash=$1"
	err:=ss.db.QueryRow(query,parent_hash).Scan(&parent_block_num);
	if no_chainsplit_check {
		err = nil // clear error as we don't need to validate the chain
		parent_block_num = block.Number.Int64()-1
	}
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			query = "SELECT count(*) FROM "+ss.schema_name+".block"
			row := ss.db.QueryRow(query)
			var block_count int64
			err := row.Scan(&block_count)
			if (err!=nil) {
				ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
				os.Exit(1)
			}
			if block_count > 0 {
				starting_block:=ss.Bigstats_get_starting_block_from_config()
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
		INSERT INTO `+ss.schema_name+`.block(
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
