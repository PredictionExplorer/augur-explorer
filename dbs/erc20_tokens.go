package dbs

import (
	"fmt"
	"os"
	"math/big"
	"database/sql"
	_  "github.com/lib/pq"


	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_erc20_process_status() p.ERC20ProcessStatus {

	var output p.ERC20ProcessStatus
	var null_last_evtid sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM erc20_proc_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_last_evtid)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO erc20_proc_status DEFAULT VALUES"
				_,err := ss.db.Exec(query)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			}
		} else {
			break
		}
	}
	if null_last_evtid.Valid {
		output.LastEvtId  = null_last_evtid.Int64
	}
	return output
}
func (ss *SQLStorage) Update_erc20_process_status(status *p.ERC20ProcessStatus) {

	var query string
	query = "UPDATE erc20_proc_status SET last_evt_id = $1"

	_,err := ss.db.Exec(query,status.LastEvtId)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}

func (ss *SQLStorage) Update_erc20_token_balances_backwards(last_block_num int64,contract_aid,acct_aid int64,eth_balance *big.Int) int {
	// Note: we are using block_hash in WHERE conditions to prevent balance corruption during chain split
	var updated_rows  int =0
	var query string

	query = "SELECT id FROM erc20_bal "+
				"WHERE (contract_aid=$1) AND (aid=$2) AND (processed=FALSE) "+
				"ORDER BY id DESC LIMIT 1"
	row:=ss.db.QueryRow(query,contract_aid,acct_aid)
	var null_id sql.NullInt64
	var stopping_id int64 = 0
	var block_hash string
	err := row.Scan(&null_id)
	if err == nil {
		if null_id.Valid {
			stopping_id = null_id.Int64
		}
	}
	ss.Info.Printf("balance_updater(): update backwards: stopping ID=%v\n",stopping_id)

	query = "SELECT " +
				"id," +
				"balance::text as balance," +
				"amount::text as amount," +
				"processed, " +
				"b.block_hash " +
			"FROM erc20_bal AS db " +
				"JOIN block AS b on db.block_num = b.block_num " +
			"WHERE " +
				"(db.aid = $1) AND " +
				"(db.contract_aid=$2) AND "+
				"(db.block_num <= $3) " +
			"ORDER BY db.id DESC"
	rows,err := ss.db.Query(query,acct_aid,contract_aid,last_block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	correct_balance := new(big.Int)
	correct_balance.Set(eth_balance)
	ss.Info.Printf("balance_updater(): Entering update_erc20_token_balances_backwards() with eth_balance=%v correct_balace=%v\n",eth_balance.String(),correct_balance.String())
	var row_count = 0;
	defer rows.Close()
	for rows.Next() {
		row_count++
		var id int64
		var balance_str string
		var amount_str string
		var processed bool
		err = rows.Scan(&id,&balance_str,&amount_str,&processed,&block_hash)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		db_balance := new(big.Int)
		db_balance.SetString(balance_str,10)
		amount := new(big.Int)
		amount.SetString(amount_str,10)
		tmp_int := new(big.Int)
		tmp_int.Set(correct_balance)
		correct_balance.Sub(tmp_int,amount)	// inverse operation to Add()
		cmp_res := correct_balance.Cmp(db_balance)
		ss.Info.Printf("balance_updater(): aid=%v,id=%v,correct=%v,db=%v,amount=%v,cmp_res=%v\n",
					acct_aid,id,correct_balance.String(),db_balance.String(),amount.String(),cmp_res)
		if cmp_res != 0 {	// incorrect balance, update it
			ss.Info.Printf("balance_updater(): incorrect balance, setting correct balance to %v for id=%v\n",
				correct_balance.String(),id)
			query = "UPDATE erc20_bal AS db " +
					"SET balance=$3,processed = true " +
					"FROM block AS b " +
					" WHERE db.block_num=b.block_num AND b.block_hash=$2 AND db.id=$1"
			ss.Info.Printf("query = %v\n",query)
			_,err = ss.db.Exec(query,id,block_hash,correct_balance.String())
			if (err!=nil) {
				ss.Log_msg(fmt.Sprintf("DB Error: %v",err));
				os.Exit(1)
			}
			updated_rows++
		} else {
			if !processed {
				query = "UPDATE erc20_bal AS db " +
						"SET processed = true " +
						"FROM block AS b " +
						"WHERE db.block_num=b.block_num AND b.block_hash=$2 AND db.id=$1"
				_,err = ss.db.Exec(query,id,block_hash)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB Error: %v",err));
					os.Exit(1)
				}
				updated_rows++
			}
			if id <= stopping_id {
				return updated_rows
			}
		}
	}
	return updated_rows
}
func (ss *SQLStorage) Set_erc20_balance(id int64,block_hash string,balance string) {
	//Note: we are using block hash in WHERE condition because during balance update process
	//		chain split could occur and the block hash can change
	var query string
	query = "UPDATE erc20_bal AS db " +
				"SET balance = $3, " +
				"processed=true " +
			"FROM block AS b " +
			"WHERE db.block_num=b.block_num AND b.block_hash=$2 AND db.id=$1"

	_,err := ss.db.Exec(query,id,block_hash,balance)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB Error: %v",err));
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_ERC20_token_transfer(contract_addr string,evt *p.ETransfer,block_num,tx_id int64,evtlog_id int64) {


	contract_aid := ss.Lookup_or_create_address(contract_addr,block_num,tx_id)
	from_aid := ss.Lookup_or_create_address(evt.From.String(),block_num,tx_id)
	to_aid := ss.Lookup_or_create_address(evt.To.String(),block_num,tx_id)
	amount := evt.Value.String()

	var query string
	query = "INSERT INTO erc20_transf("+
				"evtlog_id,block_num,tx_id,contract_aid,from_aid,to_aid,amount" +
			") " +
			"VALUES($1,$2,$3,$4,$5,$6,$7::DECIMAL)"
	_,err := ss.db.Exec(query,evtlog_id,block_num,tx_id,contract_aid,from_aid,to_aid,amount)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_unprocessed_erc20_balances(below_id int64) []p.ERC20B {

	var id_condition string
	if below_id > 0 {
		id_condition = fmt.Sprintf(" AND (db.id > %v) ",below_id)
	}
	records := make([]p.ERC20B,0,8)
	var query string
	query = "SELECT " +
				"db.id," +
				"db.aid," +
				"db.parent_id," +
				"db.contract_aid,"+
				"a.addr," +
				"amount," +
				"balance as balance," +
				"db.block_num, " +
				"b.block_hash, " +
				"ca.addr " +
			"FROM erc20_bal db " +
				"JOIN block AS b ON db.block_num = b.block_num " +
				"LEFT JOIN address a ON db.aid=a.address_id " +
				"LEFT JOIN address ca ON db.contract_aid=ca.address_id " +
			"WHERE (processed = false) " + id_condition +
			"ORDER by db.id " +
			"LIMIT 10"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.ERC20B
		err=rows.Scan(
			&rec.Id,
			&rec.Aid,
			&rec.ParentId,
			&rec.ContractAid,
			&rec.Address,
			&rec.Amount,
			&rec.Balance,
			&rec.BlockNum,
			&rec.BlockHash,
			&rec.ContractAddr,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Error at Scan: %v\n",err))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_previous_erc20_balance_from_DB(id int64,contract_aid,acct_aid int64) (string,error) {

	var query string
	query = "SELECT balance::text,processed FROM erc20_bal " +
			"WHERE (contract_aid=$1) AND (aid=$2) AND (id<$3) ORDER BY id DESC LIMIT 1"

	res := ss.db.QueryRow(query,contract_aid,acct_aid,id)
	var balance string
	var processed bool
	err := res.Scan(&balance,&processed)
	if err != nil {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		return balance,err
	}
	if !processed {
		return "",ErrUnprocessedBalances
	}
	return balance,err
}
func (ss *SQLStorage) Get_erc20_operations(factory_aid,market_id int64,offset,limit int) []p.AMM_ERC20_Op {

	records := make([]p.AMM_ERC20_Op,0,128)
	var query string

	query = "SELECT count(*) AS total FROM erc20_transf WHERE "
	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.AMM_ERC20_Op
		err=rows.Scan(&rec.FromAddr)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Error on Scan() at Get_erc20_operations(): %v\n",err))
			os.Exit(1)
		}
	}
	return records
}
