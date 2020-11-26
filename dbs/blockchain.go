package dbs

import (
	"fmt"
	"os"
	"encoding/hex"
	"strings"

	"database/sql"
	_  "github.com/lib/pq"

//	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
// Notes on design:
//		Layer1 is checking block hash sequence to always be valid by matching parent hash
//		Blocks are deleted during chain reorg in reverse order (from highest block to lowest)
//		When block is deleted from `block` table, the DELETEs are cascaded to `transaction` table
//		and to any tables on Layer2, therefore DELETEs on higher layers occur automatically
//		During incremental data population higher layers are using ID fields to select event records
//		for processing, since ID fields are never decreasing the data being processed will be
//		always newer than previously processed data
//		Processing granularity on layer2 is by transaction, while on layer1 it is by block,
//		Event atomicity: when event logs of a Tx are added, they are INSERTed all at once to ensure
//		atomic appearance of the data of all event log records and avoid inconsistency that can
//		occurr on layer2 Unix processes which are running separately (independently from layer1)

func (ss *SQLStorage) Get_last_block_num() (int64,bool) {

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
		return null_block_num.Int64,true
	} else {
		return -1,false
	}
}
func (ss *SQLStorage) Get_first_block_num() int64 {

	var query string
	query="SELECT block_num FROM block ORDER by block_num LIMIT 1";
	row := ss.db.QueryRow(query)
	var null_block_num sql.NullInt64
	var err error
	err=row.Scan(&null_block_num);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("No blocks on Layer1, aborting."))
			os.Exit(1)
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_first_block_num(): %v",err))
			os.Exit(1)
		}
	}
	if (!null_block_num.Valid) {
		ss.Log_msg(fmt.Sprintf("First block is null on Layer1"))
		os.Exit(1)
	}
	return null_block_num.Int64
}
func (ss *SQLStorage) Set_last_block_num(block_num int64) {

	bnum := int64(block_num)
	var query string = "UPDATE last_block SET block_num=$1"
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
func (ss *SQLStorage) Insert_block(hash_str string,block *types.Header,no_chainsplit_check bool) error {

	var query string
	var parent_block_num int64
	parent_hash := block.ParentHash.String()

	query="SELECT block_num FROM block WHERE block_hash=$1"
	err:=ss.db.QueryRow(query,parent_hash).Scan(&parent_block_num);
	if no_chainsplit_check {
		err = nil // clear error as we don't need to validate the chain
		parent_block_num = block.Number.Int64()-1
	}
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			query = "SELECT count(*) FROM block"
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
func (ss *SQLStorage) Insert_transaction(agtx *p.AugurTx) int64 {

	// Note: contract addresses have To as their created address + ctrct_create flag set to 'true'
	var query string
	var tx_id int64

	ss.Info.Printf("Insert_transaction: from: %v, to: %v\n",agtx.From,agtx.To)

	query = "INSERT INTO transaction ("+
				"block_num,value,tx_hash,ctrct_create,gas_used,gas_price,tx_index,input_sig" +
			") " +
			"VALUES ($1,("+agtx.Value+"/1e+18),$2,$3,$4,"+agtx.GasPrice+"/1e+18,$5,$6) " +
			"RETURNING id"

	var sig string
	if len(agtx.Input) >=4 {
		sig = hex.EncodeToString(agtx.Input[:4])
	}
	row := ss.db.QueryRow(query,agtx.BlockNum,agtx.TxHash,agtx.CtrctCreate,agtx.GasUsed,agtx.TxIndex,sig)
	err := row.Scan(&tx_id)
	if err != nil {
		if !strings.Contains(
			err.Error(),
			`duplicate key value violates unique constraint "transaction_tx_hash_key"`,
		) {
			if false {
				tx_id,err = ss.Get_tx_id_by_hash(agtx.TxHash)
				if err != nil {
					return tx_id
				}
			}
		}
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: tx_hash=%v; can't insert into transactions table: %v, q=%v",
				agtx.TxHash,err,query,
			),
		)
		os.Exit(1)
	}

	from_aid := ss.Lookup_or_create_address(agtx.From,agtx.BlockNum,tx_id)
	to_aid := ss.Lookup_or_create_address(agtx.To,agtx.BlockNum,tx_id)
	// this update is needed because 'address' table holds tx_id of account creation
	query = "UPDATE transaction set from_aid=$2 , to_aid=$3 where id = $1"
	_,err = ss.db.Exec(query,tx_id,from_aid,to_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error on tx_hash=%v; %v, q=%v\n",agtx.TxHash,err,query))
		os.Exit(1)
	}

	return tx_id
}
func (ss *SQLStorage) Update_address_metadata(aid int64,agtx *p.AugurTx) {

	// When a market address is inserted into 'address' table before market itself (this can
	// happen if 0x Mesh listener processes events before processing the blockhain data)
	// block_num and tx_id of market address will be 0. This function updates real block_num & tx_id

	var query string
	query = "UPDATE address SET block_num=$2,tx_id=$3 WHERE address_id=$1"
	_,err := ss.db.Exec(query,aid,agtx.BlockNum,agtx.TxId)
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf(
				"DB error: can't update address metadata for address %v : %v: q=%v",
				agtx.BlockNum,err,query,
			),
		)
		os.Exit(1)
	}
}
func (ss *SQLStorage) Chainsplit_delete_blocks(starting_block_num int64) {

	var err error
	var query string
	// Note: We must delete in reverse order of block creation because the triggers
	//			in the DB have made cumulative operations
	query = "DELETE FROM block WHERE block_num IN (" +
				"SELECT block_num FROM block WHERE block_num > $1 ORDER BY block_num DESC" +
			")"
	_,err = ss.db.Exec(query,starting_block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v, block_num=%v",err,query,starting_block_num))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Block_delete_with_everything(block_num int64) {

	// deletes block table and all the dependent tables receieve cascaded DELETEs also
	var query string
	query = "DELETE FROM block WHERE block_num = $1"
	_,err := ss.db.Exec(query,block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (block_num=%v, %v)",err,block_num,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_block_info(block_num_from,block_num_to int64) (p.BlockInfo,error) {

	if block_num_to == 0 {
		var success bool
		block_num_to,success = ss.Get_last_block_num()
		if !success {
			block_num_to = block_num_from
		}
	}

	var binfo p.BlockInfo
	binfo.BlockNumFrom = block_num_from
	binfo.BlockNumTo = block_num_to
	binfo.NumBlocks = block_num_to - block_num_from
	if binfo.NumBlocks < 0 { binfo.NumBlocks = 0 }


	var query string
	query = "SELECT " +
				"SUM(num_tx) AS num_tx," +
				"SUM(num_events) AS num_events," +
				"SUM(num_agtx_evts) AS num_tx," +
				"SUM(num_defi_evts) AS num_defi_evts," +
				"SUM(num_other_evts) AS num_other_events, " +
				"SUM(num_bal_swaps) AS num_bal_swaps," +
				"SUM(num_uni_swaps) AS num_uni_swaps " +
			"FROM agblk WHERE (block_num >= $1) AND (block_num <= $2)"

	row := ss.db.QueryRow(query,block_num_from,block_num_to)
	var err error
	err=row.Scan(
		&binfo.NumAugurTx,
		&binfo.NumEvents,
		&binfo.NumAugurEvents,
		&binfo.NumDefiEvents,
		&binfo.NumOtherEvents,
		&binfo.NumBalSwaps,
		&binfo.NumUniSwaps,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return binfo,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}

	query = "SELECT " +
				"SUM(gas_used) AS gas_used, " +
				"(SUM(gas_used*gas_price)/1e+18)::DECIMAL(64,18) AS tx_cost_eth " +
			"FROM agtx WHERE (block_num >= $1) AND (block_num <= $2)"
	row = ss.db.QueryRow(query,block_num_from,block_num_to)
	var null_gas sql.NullInt64
	var null_tx_cost sql.NullFloat64
	err=row.Scan(&null_gas,&null_tx_cost)
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		if null_gas.Valid {
			binfo.GasUsed = null_gas.Int64
		}
		if null_tx_cost.Valid {
			binfo.TxCostEth = null_tx_cost.Float64
		}
	}
	// get TRANSACTIONS
	records_transactions := make([]string,0,8)
	query = "SELECT tx_hash FROM agtx agt " +
			"JOIN transaction t ON agt.tx_id=t.id " +
			"WHERE (agt.block_num >= $1) AND (agt.block_num <= $2)"

	var rows *sql.Rows
	rows,err = ss.db.Query(query,block_num_from,block_num_to)
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

	// get ORDERs
	records_orders:= make([]string,0,8)
	query = "SELECT DISTINCT order_hash FROM agtx agt " +
			"JOIN mktord o ON agt.tx_id=o.id " +
			"WHERE (agt.block_num >= $1) AND (agt.block_num <= $2)"

	rows,err = ss.db.Query(query,block_num_from,block_num_to)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return binfo,err
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var order_hash string
		err=rows.Scan(&order_hash)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records_orders = append(records_orders,order_hash)
	}
	binfo.Orders= records_orders
	binfo.NumUniqueOrders = int64(len(binfo.Orders))

	// get MARKETS
	records_market := make([]string,0,8)
	query = "SELECT DISTINCT a.addr FROM market m " +
			"LEFT JOIN address a ON m.market_aid=a.address_id " +
			"WHERE (m.block_num >= $1) AND (m.block_num <= $2)"

	rows,err = ss.db.Query(query,block_num_from,block_num_to)
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
		err=rows.Scan(&market_addr)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records_market = append(records_market,market_addr)
	}
	binfo.Markets = records_market
	binfo.NumUniqueMarkets=int64(len(binfo.Markets))

	// get Active addresses
	records_addresses := make([]string,0,8)
	query = "WITH aids AS (" +
				"SELECT DISTINCT aid FROM ( "+
					"(" +
						"SELECT DISTINCT account_aid AS aid FROM agtx_evt " +
						"WHERE (block_num >= $1) AND (block_num <= $2)" +
					") UNION ALL (" +
						"SELECT DISTINCT aid FROM mktord " + // agtx_evt only stored Fill account addr
						"WHERE (block_num >= $1) AND (block_num <= $2)" +
					") " +
				") AS d " +
			") " +
			"SELECT a.addr FROM address AS a " +
				"JOIN aids ON aids.aid=a.address_id "

	rows,err = ss.db.Query(query,block_num_from,block_num_to)
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
	binfo.ActiveAddresses= records_addresses
	binfo.NumUniqueAddresses=int64(len(binfo.ActiveAddresses))

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
func (ss *SQLStorage) Get_tx_hash_by_id(tx_id int64) (string,int64,error) {

	var tx_hash string
	var block_num int64
	var query string
	query = "SELECT tx_hash,block_num FROM transaction WHERE id=$1"
	row := ss.db.QueryRow(query,tx_id)
	err := row.Scan(&tx_hash,&block_num)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return "",0,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return tx_hash,block_num,nil
}
func (ss *SQLStorage) Get_tx_hash_with_timestamp_by_id(tx_id int64) (string,int64,int64,error) {

	var tx_hash string
	var block_num int64
	var timestamp int64
	var query string
	query = "SELECT t.tx_hash,t.block_num,EXTRACT(EPOCH FROM b.ts)::BIGINT AS ts " +
			"FROM transaction AS t,block AS b "+
			"WHERE t.id=$1 AND t.block_num=b.block_num"
	row := ss.db.QueryRow(query,tx_id)
	err := row.Scan(&tx_hash,&block_num,&timestamp)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return "",0,0,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return tx_hash,block_num,timestamp,nil
}
func (ss *SQLStorage) Get_tx_hash_and_input_sig(tx_id int64) (string,string,error) {

	var tx_hash string
	var input_sig string
	var query string
	query = "select tx_hash,input_sig from transaction where id=$1"
	row := ss.db.QueryRow(query,tx_id)
	err := row.Scan(&tx_hash,&input_sig)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return "","",err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return tx_hash,input_sig,nil
}
func (ss *SQLStorage) Get_tx_id_by_hash(tx_hash string) (int64,error) {

	var tx_id int64
	var query string
	query = "select id from transaction where tx_hash=$1"
	row := ss.db.QueryRow(query,tx_hash)
	err := row.Scan(&tx_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return tx_id,nil
}
func (ss *SQLStorage) Tx_exists(tx_hash string) bool {

	var query string
	query = "SELECT id FROM transaction WHERE tx_hash=$1"

	row := ss.db.QueryRow(query,tx_hash)
	var unused int64
	err := row.Scan(&unused)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return true
}
func (ss *SQLStorage) Get_last_tx_id() (int64,error) {

	var query string
	query = "SELECT id FROM transaction ORDER BY id DESC LIMIT 1"
	res := ss.db.QueryRow(query)
	var null_id sql.NullInt64
	err := res.Scan(&null_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	}
	return null_id.Int64,nil
}
func (ss *SQLStorage) Get_block_timestamp(block_num int64) (int64,error) {

	var query string
	query = "SELECT FLOOR(EXTRACT(EPOCH FROM block.ts))::BIGINT AS ts " +
			"FROM block WHERE block_num=$1"
	row := ss.db.QueryRow(query,block_num)
	var ts int64
	var err error
	err=row.Scan(&ts);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,err
		}
		ss.Log_msg(fmt.Sprintf("Error in Get_block_timestamp(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return ts,nil
}
func (ss *SQLStorage) Get_last_block_timestamp() int64 {

	var query string
	query = "SELECT FLOOR(EXTRACT(EPOCH FROM block.ts))::BIGINT AS ts " +
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
func (ss *SQLStorage) Get_block_num_by_hash(block_hash string) (int64,error) {

	var query string
	query = "SELECT block_num FROM block WHERE block_hash=$1"

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
func (ss *SQLStorage) Set_cash_flow_value(block_num int64,new_cash_flow float64) {

	var query string
	query = "UPDATE block SET cash_flow=$2 WHERE block_num=$1"
	res,err:=ss.db.Exec(query,block_num,new_cash_flow)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Set_cash_flow_value() failed: %v, q=%v, b=%v",err,query,block_num))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in Set_cash_flow_value(): %v",err))
		os.Exit(1)
	}
	if affected_rows>0 {
		// break
	} else {
		ss.Log_msg(fmt.Sprintf("Set_cash_flow_Value() failed on UPDATe: 0 affected rows"));
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_all_address_ids() []int64 {
	// Used by dai_bal verification process

	var query string
	query = "SELECT address_id FROM address ORDER by address_id DESC"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	records := make([]int64,0,512)	// allocating one page for the whole array (4096 bytes)
	defer rows.Close()
	for rows.Next() {
		var aid int64
		err=rows.Scan(&aid)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,aid)
	}
	return records
}
func (ss *SQLStorage) Get_stored_chain_id() int64 {

	var query string
	query = "SELECT chain_id FROM contract_addresses"
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
func (ss *SQLStorage) Set_chain_id(chain_id int64) {

	var query string = "UPDATE contract_addresses SET chain_id=$1"
	_,err:=ss.db.Exec(query,chain_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Set_chain_id() failed: %v",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_tx_event_log(eel *p.EthereumEventLog) int64 {

	//contract_aid := ss.Lookup_or_create_address(eel.ContractAddress,eel.BlockNum,eel.TxId)
//	ss.Info.Printf("topic0_sig=%v, len=%v\n",eel.Topic0_Sig,len(eel.Topic0_Sig))
	var query string
	query = "INSERT INTO evt_log(block_num,tx_id,contract_aid,topic0_sig,log_rlp) " +
				"VALUES($1,$2,$3,$4,$5) RETURNING id"

	//_,err:=ss.db.Exec(query,eel.BlockNum,eel.TxId,contract_aid,eel.Data)
	var err error
	var null_id sql.NullInt64
	row := ss.db.QueryRow(query,eel.BlockNum,eel.TxId,eel.ContractAid,eel.Topic0_Sig,eel.RlpLog)
	err=row.Scan(&null_id);
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Insert_tx_event_log() failed: %v, q=%v, b=%v\n",err,query,eel.BlockNum))
		os.Exit(1)
	}
	if !null_id.Valid {
		ss.Log_msg(fmt.Sprintf("Insert_tx_event_log() failed: null id returned at block %v\n",eel.BlockNum))
		os.Exit(1)
	}
	return null_id.Int64
}
func (ss *SQLStorage) Insert_all_tx_event_logs(eelogs []p.EthereumEventLog) {

	var query strings.Builder
	query.WriteString("INSERT INTO evt_log(block_num,tx_id,contract_aid,topic0_sig,log_rlp) VALUES")

	for i,eel := range(eelogs) {
		if i > 0 {
			query.WriteString(",")
		}
		query.WriteString(fmt.Sprintf(
			"(%v,%v,%v,'%v',decode('%v','hex'))",
			eel.BlockNum,
			eel.TxId,
			eel.ContractAid,
			eel.Topic0_Sig,
			hex.EncodeToString(eel.RlpLog),
		))
	}
	var err error
	_,err = ss.db.Exec(query.String())
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Insert_all_tx_event_logs() failed: %v, q=%v\n",err,query.String()))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_event_log_topic(eet *p.EthereumEventTopic) {

	var query string
	query = "INSERT INTO evt_topic(block_num,tx_id,evtlog_id,contract_aid,pos,value) " +
				"VALUES($1,$2,$3,$4,$5,$6)"
	_,err:=ss.db.Exec(query,
			eet.BlockNum,eet.TxId,eet.EventLogId,eet.ContractAid,eet.Pos,eet.Value,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Insert_event_log_topic() failed: %v at block\n",err,eet.BlockNum))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_chain_reorg_event(co *p.ChainReorg) {

	var query string
	query = "INSERT INTO chain_reorg(block_num,hash) VALUES($1,$2)"

	_,err:=ss.db.Exec(query,co.BlockNum,co.Hash)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Insert_chain_reorg_event() failed: %v at block\n",err,co.BlockNum))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_event_log(evtlog_id int64) p.EthereumEventLog {

	var evtlog p.EthereumEventLog
	evtlog.EvtId = evtlog_id
	var query string
	query = "SELECT block_num,tx_id,contract_aid,topic0_sig,log_rlp FROM evt_log WHERE id=$1"
	res := ss.db.QueryRow(query,evtlog_id)
	err := res.Scan(&evtlog.BlockNum,&evtlog.TxId,&evtlog.ContractAid,&evtlog.Topic0_Sig,&evtlog.RlpLog)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}

	return evtlog
}
func (ss *SQLStorage) Get_augur_transaction(tx_id int64) *p.AugurTx {

	agtx := new(p.AugurTx)
	var query string
	query = "SELECT t.block_num,tx_hash,EXTRACT(EPOCH FROM b.ts)::BIGINT AS ts " +
			"FROM transaction t,block b " +
			"WHERE t.id=$1 AND b.block_num=t.block_num"

	res := ss.db.QueryRow(query,tx_id)
	err := res.Scan(&agtx.BlockNum,&agtx.TxHash,&agtx.TimeStamp)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	agtx.TxId=tx_id
	return agtx
}
func (ss *SQLStorage) Get_tx_ids_from_evt_logs_by_signature(sig string,contract_aid int64,from_tx_id int64,to_tx_id int64) []int64 {


	output := make([]int64,0,1024)

	var query string
	query = "SELECT DISTINCT tx_id FROM evt_log " +
				"WHERE (tx_id > $1) AND (tx_id <= $2) " +
						"AND (contract_aid=$3) " +
						"AND (topic0_sig=$4) " +
				"ORDER BY tx_id "

	rows,err := ss.db.Query(query,from_tx_id,to_tx_id,contract_aid,sig)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var tx_id int64 
		err=rows.Scan(&tx_id)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
		output = append(output,tx_id)
	}
	return output
}
func (ss *SQLStorage) Get_evt_log_ids_by_signature(sig string,contract_aids string,from_evt_id int64,num_rows int64) []p.ShortEvtLog {


	output := make([]p.ShortEvtLog,0,1024)

	var query string
	query = "SELECT id,tx_id FROM evt_log " +
				"WHERE id > $1 " +
						"AND (contract_aid IN ("+contract_aids+")) " +
						"AND (topic0_sig=$2) " +
				"ORDER BY id "+
				"LIMIT $3"

	rows,err := ss.db.Query(query,from_evt_id,sig,num_rows)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var evt_log p.ShortEvtLog
		err=rows.Scan(&evt_log.EvtId,&evt_log.TxId)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
		output = append(output,evt_log)
	}
	return output
}
func (ss *SQLStorage) Get_evtlogs_by_signature_in_range(sig string,contract_aids string,from_evt_id,to_evt_id int64) []int64 {


	output := make([]int64,0,1024)

	var query string
	query = "SELECT id FROM evt_log " +
				"WHERE (id > $1) AND (id<=$2) " +
						"AND (contract_aid IN ("+contract_aids+")) " +
						"AND (topic0_sig=$3) " +
				"ORDER BY id "

	rows,err := ss.db.Query(query,from_evt_id,to_evt_id,sig)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var evtlog_id int64
		err=rows.Scan(&evtlog_id)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
		output = append(output,evtlog_id)
	}
	return output
}
func (ss *SQLStorage) Get_evtlogs_by_signature_only_in_range(sig string,from_evt_id,to_evt_id int64) []int64 {


	output := make([]int64,0,1024)

	var query string
	query = "SELECT id FROM evt_log " +
				"WHERE (id > $1) AND (id<=$2) " +
						"AND (topic0_sig=$3) " +
				"ORDER BY id "

	rows,err := ss.db.Query(query,from_evt_id,to_evt_id,sig)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var evtlog_id int64
		err=rows.Scan(&evtlog_id)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
		output = append(output,evtlog_id)
	}
	return output
}
func (ss *SQLStorage) Get_LOG_CALL_evtlogs(sig string,from_evt_id,to_evt_id int64) []int64 {

	// this function scans transaction input signatures, but returns event_ids
	//			(using evtlog_ids is required for ID merge-sort functions)
	output := make([]int64,0,1024)

	var query string
	/*DISCONTINUED, doesn't fetch proxied controlling transactions
	query = "SELECT e.id FROM evt_log e " +
				"JOIN transaction AS t ON e.tx_id=t.id " +
				"WHERE (e.id > $1) AND (e.id<=$2) " +
						"AND (t.input_sig=$3) " +
						"AND (e.topic0_sig=$3)" + // this is the patern of anonymous LOG_CALL event
				"ORDER BY e.id "
	*/
	query = "SELECT e.id FROM evt_log e " +
				"WHERE (e.id > $1) AND (e.id<=$2) " +
						"AND (e.topic0_sig=$3)" + // this is the patern of anonymous LOG_CALL event
				"ORDER BY e.id "

	rows,err := ss.db.Query(query,from_evt_id,to_evt_id,sig)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var evtlog_id int64
		err=rows.Scan(&evtlog_id)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
		output = append(output,evtlog_id)
	}
	return output
}
func (ss *SQLStorage) Delete_transaction_related_data(tx_id int64) {

	// Note: the list of DELETEs must match the list of event signatures
	//			built in built_list_of_expected_events() function

	ss.Delete_market_created_evt(tx_id)
	ss.Delete_market_oi_changed_evt(tx_id)
	ss.Delete_market_order_evt(tx_id)
	ss.Delete_market_finalized_evt(tx_id)
	ss.Delete_report_evt(tx_id)
	ss.Delete_market_vol_changed_evt(tx_id)
	ss.Delete_token_balance_changed_evt(tx_id)
	ss.Delete_share_balance_changed_evt(tx_id)
	ss.Delete_cancel_open_order_evt(tx_id)
	ss.Delete_profit_loss_evt(tx_id)
	ss.Delete_trading_proceeds_claimed_evt(tx_id)
	ss.Delete_register_contract_evt(tx_id)
	ss.Delete_claim_funds(tx_id)
}
func (ss *SQLStorage) Insert_dummy_block(block_num int64) {

	var query string
	query = 
		"INSERT INTO block(block_num,block_hash,ts,parent_hash) " +
		"VALUES ($1,'hash',NOW(),'parent_hash')"

	_,err := ss.db.Exec(query,block_num)
	if err != nil {
		ss.Log_msg(
			fmt.Sprintf("DB error: %v, q=%v",err,query),
		)
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_first_event_log() p.BasicChainInfo {

	var query string
	query = "SELECT id,block_num,tx_id FROM evt_log ORDER BY id DESC LIMIT 1"

	row := ss.db.QueryRow(query)
	var bci p.BasicChainInfo
	var err error
	err=row.Scan(&bci.EvtId,&bci.BlockNum,&bci.TxId);
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Error in Get_first_event_log(): %v",err))
		os.Exit(1)
	}
	return bci
}
