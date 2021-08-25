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
func (ss *SQLStorage) Insert_block(hash_str string,block *types.Header,num_tx int,no_chainsplit_check bool) error {

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
func (ss *SQLStorage) Insert_transaction(agtx *p.AugurTx) int64 {

	// Note: contract addresses have To as their created address + ctrct_create flag set to 'true'
	var query string
	var tx_id int64

	//ss.Info.Printf("Insert_transaction: from: %v, to: %v\n",agtx.From,agtx.To)

	from_aid := ss.Lookup_or_create_address(agtx.From,agtx.BlockNum,0)
	to_aid := ss.Lookup_or_create_address(agtx.To,agtx.BlockNum,0)
	query = "INSERT INTO transaction ("+
				"block_num,from_aid,to_aid,value,tx_hash,ctrct_create,gas_used,gas_price,tx_index,input_sig,num_logs" +
			") " +
			"VALUES ($1,$2,$3,("+agtx.Value+"/1e+18),$4,$5,$6,"+agtx.GasPrice+"/1e+18,$7,$8,$9) " +
			"RETURNING id"

	var sig string
	if len(agtx.Input) >=4 {
		sig = hex.EncodeToString(agtx.Input[:4])
	}
	row := ss.db.QueryRow(query,agtx.BlockNum,from_aid,to_aid,agtx.TxHash,agtx.CtrctCreate,agtx.GasUsed,agtx.TxIndex,sig,agtx.NumLogs)
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
	var err error

	query = "SELECT " +
				"EXTRACT(EPOCH FROM b1.ts)::BIGINT AS ini_ts," +
				"b1.ts AS ini_date," +
				"EXTRACT(EPOCH FROM b2.ts)::BIGINT AS fin_ts," +
				"b2.ts AS fin_date " +
			"FROM block b1 LEFT JOIN block b2 ON $2=b2.block_num " +
			"WHERE b1.block_num=$1"
	var nts1,nts2 sql.NullInt64
	var nd1,nd2 sql.NullString
	row := ss.db.QueryRow(query,binfo.BlockNumFrom,binfo.BlockNumTo)
	err=row.Scan(&nts1,&nd1,&nts2,&nd2)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return binfo,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	binfo.FromTimeStamp = nts1.Int64
	binfo.ToTimeStamp = nts2.Int64
	binfo.FromDate = nd1.String
	binfo.ToDate = nd2.String

	query = "SELECT " +
				"COALESCE(SUM(num_tx),0) AS num_tx," +
				"COALESCE(SUM(num_events),0) AS num_events," +
				"COALESCE(SUM(num_agtx_evts),0) AS num_tx," +
				"COALESCE(SUM(num_defi_evts),0) AS num_defi_evts," +
				"COALESCE(SUM(num_other_evts),0) AS num_other_events, " +
				"COALESCE(SUM(num_bal_swaps),0) AS num_bal_swaps," +
				"COALESCE(SUM(num_uni_swaps),0) AS num_uni_swaps " +
			"FROM agblk WHERE (block_num >= $1) AND (block_num <= $2)"

	row = ss.db.QueryRow(query,block_num_from,block_num_to)
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
			d_query:=strings.ReplaceAll(query,`$1`,fmt.Sprintf("%v",binfo.BlockNumFrom))
			d_query = strings.ReplaceAll(d_query,`$2`,fmt.Sprintf("%v",binfo.BlockNumTo))
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,d_query))
			os.Exit(1)
		}
	}

	query = "SELECT " +
				"SUM(gas_used) AS gas_used, " +
				"SUM(gas_used*gas_price) AS tx_cost_eth " +
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
	query = "SELECT DISTINCT order_hash FROM agtx_evt e " +
			"JOIN mktord o ON e.ref_id=o.id " +
			"WHERE (e.block_num >= $1) AND (e.block_num <= $2)"

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

	// get MARKETS created
	records_markets_created := make([]p.MarketVeryShortInfo,0,8)
	query = "SELECT " +
				"DISTINCT a.addr, " +
				"m.extra_info::json->>'description' AS descr " +
			"FROM market m " +
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
		var minfo p.MarketVeryShortInfo
		var market_desc sql.NullString
		err=rows.Scan(&minfo.MktAddr,&market_desc)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		minfo.MktDesc = market_desc.String
		records_markets_created = append(records_markets_created,minfo)
	}
	binfo.MarketsCreated = records_markets_created
	binfo.NumMarketsCreated=int64(len(binfo.MarketsCreated))

	// get MARKETS traded
	records_markets_traded := make([]p.MarketVeryShortInfo,0,8)
	query = "WITH aids AS (" +
				"SELECT DISTINCT aid FROM ( "+
					"SELECT DISTINCT market_aid AS aid FROM agtx_evt " +
					"WHERE (block_num >= $1) AND (block_num <= $2)" +
				") AS d " +
			") " +
			"SELECT " +
				"a.addr, " +
				"m.extra_info::json->>'description' AS descr  " +
			"FROM market AS m " +
				"JOIN aids ON aids.aid=m.market_aid " +
				"JOIN address AS a ON a.address_id=m.market_aid"

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
		var minfo p.MarketVeryShortInfo
		var market_desc sql.NullString
		err=rows.Scan(&minfo.MktAddr,&market_desc)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		minfo.MktDesc = market_desc.String
		records_markets_traded = append(records_markets_traded,minfo)
	}
	binfo.MarketsTraded = records_markets_traded
	binfo.NumMarketsTraded=int64(len(binfo.MarketsTraded))

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
				"t.id," +
				"t.block_num," +
				"sa.addr," +
				"ra.addr," +
				"t.value, " +
				"t.from_aid,"+
				"t.to_aid," +
				"t.gas_used," +
				"t.gas_used*t.gas_price AS tx_fee," +
				"t.tx_hash," +
				"at.num_events,"+
				"at.num_defi_evts,"+
				"at.num_agtx_evts,"+
				"at.num_other_evts,"+
				"at.num_bal_swaps,"+
				"at.num_uni_swaps "+
			"FROM transaction t " +
				"LEFT JOIN agtx AS at ON t.id=at.tx_id " +
				"LEFT JOIN address sa ON t.from_aid = sa.address_id " +
				"LEFT JOIN address ra ON t.to_aid = ra.address_id " +
			"WHERE t.tx_hash=$1"

	row := ss.db.QueryRow(query,tx_hash)
	var num_events,num_defi_events,num_agtx_events,num_other_events,num_bal_swaps,num_uni_swaps sql.NullInt64
	err := row.Scan(
		&ti.TxId,
		&ti.BlockNum,
		&ti.From,
		&ti.To,
		&ti.Value,
		&ti.FromAid,
		&ti.ToAid,
		&ti.GasUsed,
		&ti.TxFeeEth,
		&ti.Hash,
		&num_events,
		&num_defi_events,
		&num_agtx_events,
		&num_other_events,
		&num_bal_swaps,
		&num_uni_swaps,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return ti,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	ti.TotalEvents=int(num_events.Int64)
	ti.NumDeFiEvents=int(num_defi_events.Int64)
	ti.NumAugurEvents=int(num_agtx_events.Int64)
	ti.NumOtherEvents=int(num_other_events.Int64)
	ti.NumBalancerSwaps=int(num_bal_swaps.Int64)
	ti.NumUniswapSwaps=int(num_uni_swaps.Int64)

	query = "SELECT DISTINCT p.pool_aid,a.addr,num_swaps,num_holders " +
			"FROM agtx_evt AS e " +
				"JOIN bswap s ON e.ref_id=s.evtlog_id " +
				"JOIN bpool p ON p.pool_aid=s.pool_aid " +
				"JOIN address AS a ON p.pool_aid=a.address_id " +
			"WHERE e.tx_id=$1"

	d_query := fmt.Sprintf("SELECT DISTINCT p.pool_aid,a.addr,num_swaps,num_holders " +
			"FROM agtx_evt AS e " +
				"JOIN bswap s ON e.ref_id=s.evtlog_id " +
				"JOIN bpool p ON p.pool_aid=s.pool_aid " +
				"JOIN address AS a ON p.pool_aid=a.address_id " +
			"WHERE e.tx_id=%v",ti.TxId)
	ss.Info.Printf("d_query= %v\n",d_query)
	rows,err := ss.db.Query(query,ti.TxId)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	} else {
		defer rows.Close()
		for rows.Next() {
			var rec p.PoolVeryShortInfo
			err=rows.Scan(&rec.PoolAid,&rec.PoolAddr,&rec.NumSwaps,&rec.NumHolders)
			if err!=nil {
				ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
				os.Exit(1)
			}
			ti.BalancerPools = append(ti.BalancerPools,rec)
		}
	}

	balancer_records := make([]p.BalancerSwap,0,64)
	query = "SELECT " +
				"s.id," +
				"s.pool_aid,"+
				"pa.addr, " +
				"FLOOR(EXTRACT(EPOCH FROM s.time_stamp))::BIGINT AS ts, " +
				"s.time_stamp as datetime,"+
				"s.block_num," +
				"s.tx_id,"+
				"ca.addr,"+
				"tia.addr," +
				"toa.addr," +
				"s.token_in_aid," +
				"s.token_out_aid," +
				"e_in.Symbol,"+
				"e_out.Symbol," +
				"s.amount_in, " +
				"s.amount_out " +
			"FROM bswap AS s " +
				"LEFT JOIN address AS pa ON s.pool_aid=pa.address_id " +
				"LEFT JOIN address AS ca ON s.caller_aid=ca.address_id " +
				"LEFT JOIN address AS tia ON s.token_in_aid=tia.address_id " +
				"LEFT JOIN address AS toa ON s.token_out_aid=toa.address_id " +
				"LEFT JOIN erc20_info e_in ON s.token_in_aid=e_in.aid " +
				"LEFT JOIN erc20_info e_out ON s.token_out_aid=e_out.aid " +
			"WHERE s.tx_id=$1 "
	rows,err = ss.db.Query(query,ti.TxId)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.BalancerSwap
		var symbol_in,symbol_out sql.NullString
		err=rows.Scan(
			&rec.Id,
			&rec.PoolAid,
			&rec.PoolAddr,
			&rec.TimeStamp,
			&rec.Date,
			&rec.BlockNum,
			&rec.TxId,
			&rec.CallerAddr,
			&rec.TokenInAddr,
			&rec.TokenOutAddr,
			&rec.TokenInAid,
			&rec.TokenOutAid,
			&symbol_in,
			&symbol_out,
			&rec.AmountInF,
			&rec.AmountOutF,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.SymbolIn = symbol_in.String
		rec.SymbolOut = symbol_out.String
		balancer_records = append(balancer_records,rec)
	}
	ti.BalancerSwaps=balancer_records

	// Uniswap Swaps
	uniswap_records := make([]p.UniswapSwap,0,128)
	query = "SELECT " +
				"sw.id,"+
				"sw.pair_aid," +
				"pa.addr," +
				"sw.block_num," +
				"sw.amount0_in, " +
				"sw.amount1_in," +
				"sw.amount0_out," +
				"sw.amount1_out," +
				"sw.time_stamp," +
				"EXTRACT(EPOCH FROM sw.time_stamp)::BIGINT AS created_ts, "+
				"ra.addr AS recipient_addr," +
				"sw.recipient_aid, " +
				"e1.symbol,"+
				"e2.symbol "+
			"FROM uswap1 AS sw " +
			"JOIN upair AS p ON sw.pair_aid=p.pair_aid " +
			"LEFT JOIN address AS ra ON sw.recipient_aid=ra.address_id " +
			"LEFT JOIN address AS pa ON sw.pair_aid=pa.address_id " +
			"LEFT JOIN erc20_info e1 ON p.token0_aid=e1.aid " +
			"LEFT JOIN erc20_info e2 ON p.token1_aid=e2.aid " +
			"WHERE sw.tx_id=$1 " +
			"ORDER BY sw.id"

	rows,err = ss.db.Query(query,ti.TxId)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.UniswapSwap
		err=rows.Scan(
			&rec.Id,
			&rec.PairAid,
			&rec.PairAddr,
			&rec.BlockNum,
			&rec.Amount0_In,
			&rec.Amount1_In,
			&rec.Amount0_Out,
			&rec.Amount1_Out,
			&rec.CreatedDate,
			&rec.CreatedTs,
			&rec.RequesterAddr,
			&rec.RequesterAid,
			&rec.Symbol0,
			&rec.Symbol1,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		uniswap_records = append(uniswap_records,rec)
	}
	ti.UniswapSwaps = uniswap_records

	query = "SELECT DISTINCT p.pair_aid,a.addr,total_swaps " +
			"FROM agtx_evt AS e " +
				"JOIN uswap1 s ON e.ref_id=s.evtlog_id " +
				"JOIN upair p ON p.pair_aid=s.pair_aid " +
				"JOIN address AS a ON p.pair_aid=a.address_id " +
			"WHERE e.tx_id=$1"

	rows,err = ss.db.Query(query,ti.TxId)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	} else {
		defer rows.Close()
		for rows.Next() {
			var rec p.PairVeryShortInfo
			err=rows.Scan(&rec.PairAid,&rec.PairAddr,&rec.TotalSwaps)
			if err!=nil {
				ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
				os.Exit(1)
			}
			ti.UniswapPairs = append(ti.UniswapPairs,rec)
		}
	}

	query = "SELECT " +
				"evt_type," +
				"ref_id," +
				"account_aid," +
				"a.addr," +
				"e.market_aid," +
				"ma.addr," +
				"defi_platform, " +
				"o.order_hash, " +
				"m.extra_info::json->>'description' AS descr, " +
				"bsw.id AS bal_swap_id, " +
				"usw.id AS uni_swap_id " +
			"FROM agtx_evt e " +
				"LEFT JOIN address a ON e.account_aid=a.address_id " +
				"LEFT JOIN address ma ON e.market_aid=ma.address_id " +
				"LEFT JOIN mktord AS o ON (e.ref_id=o.id AND e.evt_type=3) " +
				"LEFT JOIN bswap AS bsw ON (e.ref_id=bsw.evtlog_id AND e.evt_type=1 AND defi_platform=2) " +
				"LEFT JOIN uswap1 AS usw ON (e.ref_id=usw.evtlog_id AND e.evt_type=1 AND defi_platform=1) " +
				"LEFT JOIN market AS m ON e.market_aid=m.market_aid " +
			"WHERE e.tx_id = $1"

	event_records := make([]p.AgtxEvent,0,32)
	rows,err = ss.db.Query(query,ti.TxId)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	} else {
		defer rows.Close()
		for rows.Next() {
			var rec p.AgtxEvent
			var balancer_swap_id,uniswap_swap_id sql.NullInt64
			var order_hash,mkt_descr sql.NullString
			err=rows.Scan(
				&rec.EvtType,
				&rec.ReferenceId,
				&rec.Aid,
				&rec.Addr,
				&rec.MktAid,
				&rec.MktAddr,
				&rec.DeFiPlatformCode,
				&order_hash,
				&mkt_descr,
				&balancer_swap_id,
				&uniswap_swap_id,
			)
			if err!=nil {
				ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
				os.Exit(1)
			}
			if order_hash.Valid { rec.OrderHash = order_hash.String }
			if mkt_descr.Valid { rec.MktDescr = mkt_descr.String }
			if rec.DeFiPlatformCode == 1 { // Uniswap
				if uniswap_swap_id.Valid { rec.DeFiSwapId = uniswap_swap_id.Int64 }
			}
			if rec.DeFiPlatformCode == 2 { // Balancer
				if balancer_swap_id.Valid { rec.DeFiSwapId = balancer_swap_id.Int64 }
			}
			event_records = append(event_records,rec)
		}
	}
	ti.FullEventList = event_records


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
	query = "SELECT " +
				"e.block_num," +
				"EXTRACT(EPOCH FROM b.ts)::BIGINT AS ts, "+
				"e.tx_id," +
				"tx.tx_hash," +
				"e.contract_aid," +
				"ca.addr, " +
				"e.topic0_sig," +
				"e.log_rlp " +
			"FROM evt_log e " +
				"JOIN block b ON e.block_num=b.block_num " +
				"JOIN transaction tx ON e.tx_id=tx.id " +
				"JOIN address ca ON e.contract_aid=ca.address_id " +
			"WHERE e.id=$1"
	res := ss.db.QueryRow(query,evtlog_id)
	err := res.Scan(
		&evtlog.BlockNum,
		&evtlog.TimeStamp,
		&evtlog.TxId,
		&evtlog.TxHash,
		&evtlog.ContractAid,
		&evtlog.ContractAddress,
		&evtlog.Topic0_Sig,
		&evtlog.RlpLog,
	)
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
/*WAY TOO SLOW	query = "SELECT DISTINCT tx_id FROM evt_log " +
				"WHERE (tx_id > $1) AND (tx_id <= $2) " +
						"AND (contract_aid=$3) " +
						"AND (topic0_sig=$4) " +
				"ORDER BY tx_id "*/
	query = "WITH elogs AS (" +
				"SELECT tx_id,contract_aid,id AS evtlog_id " +
					"FROM evt_log " +
					"WHERE topic0_sig=$4" +
				")" +
			"SELECT " +
					"tx_id " +
				"FROM elogs AS e " +
				"WHERE " +
					"(tx_id > $1 AND tx_id <= $2) AND " +
					"contract_aid = $3"

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
func (ss *SQLStorage) Get_last_evtlog_id() (int64,error) {

	var query string
	query = "SELECT id FROM evt_log ORDER BY id DESC LIMIT 1"
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
func (ss *SQLStorage) Get_erc20_info(tok_addr string) (p.ERC20Info,error) {

	var query string
	query = "SELECT aid,decimals,total_supply,name,symbol " +
			"FROM erc20_info inf,address a WHERE inf.aid=a.address_id and a.addr=$1"

	row := ss.db.QueryRow(query,tok_addr)
	var info p.ERC20Info
	var err error
	err=row.Scan(&info.Aid,&info.Decimals,&info.TotalSupply,&info.Name,&info.Symbol);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return info,err
		}
		ss.Log_msg(fmt.Sprintf("Error in Get_first_event_log(): %v",err))
		os.Exit(1)
	}
	return info,nil
}
func (ss *SQLStorage) Get_block_range_for_whats_new(interval_code p.WhatsNewAugurCode) (int,int,error) {

	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM b.ts)::BIGINT AS ts,lb.block_num " +
			"FROM block AS b,last_block AS lb WHERE lb.block_num=b.block_num"

	row := ss.db.QueryRow(query)
	var err error
	var to_ts,to_block_num sql.NullInt64
	err=row.Scan(&to_ts,&to_block_num)
	if (err!=nil) {
		return 0,0,err
	}
	var timestamp_q int
	switch interval_code {
	case p.WNA_6Hours:
		timestamp_q=int(to_ts.Int64) - 6*60*60
	case p.WNA_12Hours:
		timestamp_q=int(to_ts.Int64) - 12*60*60
	case p.WNA_1Day:
		timestamp_q=int(to_ts.Int64) - 24*60*60
	case p.WNA_2Days:
		timestamp_q=int(to_ts.Int64) - 48*60*60
	case p.WNA_3Days:
		timestamp_q=int(to_ts.Int64) - 72*60*60
	case p.WNA_1Week:
		timestamp_q=int(to_ts.Int64) - 7*24*60*60
	case p.WNA_2Weeks:
		timestamp_q=int(to_ts.Int64) - 2*7*24*60*60
	default:
		panic("WhatsNewAugurCode with invalid value")
	}
	query = "SELECT " +
				"block_num " +
			"FROM block AS b " +
			"WHERE ts < TO_TIMESTAMP($1) " +
			"ORDER BY ts DESC " +
			"LIMIT 1"
	row = ss.db.QueryRow(query,timestamp_q)
	var from_block_num sql.NullInt64
	err=row.Scan(&from_block_num)
	if (err!=nil) {
		return 0,0,err
	}
	return int(from_block_num.Int64),int(to_block_num.Int64),nil
}
func (ss *SQLStorage) Get_events_by_sig_and_tx_id(tx_id int64,sig string) ([]p.EthereumEventLog,error) {

	var query string
	records := make([]p.EthereumEventLog,0,4)
	query = "SELECT id,block_num,contract_aid,topic0_sig,log_rlp FROM evt_log " +
			"WHERE tx_id=$1 AND topic0_sig=$2"

	rows,err := ss.db.Query(query,tx_id,sig)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return records,nil
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var event p.EthereumEventLog
		err=rows.Scan(&event.EvtId,&event.BlockNum,&event.ContractAid,&event.Topic0_Sig,&event.RlpLog)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,event)
	}
	return records,nil
}
