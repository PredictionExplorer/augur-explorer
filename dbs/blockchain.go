package dbs

import (
	"fmt"
	"os"
	"encoding/hex"

	"database/sql"
	_  "github.com/lib/pq"

	//"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
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
func (ss *SQLStorage) Get_block_info(block_num int64) (p.BlockInfo,error) {

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

	contract_aid := ss.Lookup_or_create_address(eel.ContractAddress,eel.BlockNum,eel.TxId)

	var query string
	query = "INSERT INTO evt_log(block_num,tx_id,contract_aid,data) " +
				"VALUES($1,$2,$3,$4) RETURNING id"

	//_,err:=ss.db.Exec(query,eel.BlockNum,eel.TxId,contract_aid,eel.Data)
	var err error
	var null_id sql.NullInt64
	row := ss.db.QueryRow(query,eel.BlockNum,eel.TxId,contract_aid,eel.Data)
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
func (ss *SQLStorage) Insert_event_log_topic(eet *p.EthereumEventTopic) {

	var query string
	query = "INSERT INTO evt_topic(block_num,tx_id,evtlog_id,pos,signature) " +
				"VALUES($1,$2,$3,$4,$5)"
	_,err:=ss.db.Exec(query,eet.BlockNum,eet.TxId,eet.EventLogId,eet.Pos,eet.Signature)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Insert_event_log_topic() failed: %v at block\n",err,eet.BlockNum))
		os.Exit(1)
	}
}
