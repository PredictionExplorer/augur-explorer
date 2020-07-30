package dbs

import (
	"fmt"
	"os"
	"errors"
	"bytes"
	"math/big"
	"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Process_REP_token_transfer(evt *p.Transfer,agtx *p.AugurTx) {

	from_aid := ss.Lookup_or_create_address(evt.From.String(),agtx.BlockNum,agtx.TxId)
	to_aid := ss.Lookup_or_create_address(evt.To.String(),agtx.BlockNum,agtx.TxId)
	amount := evt.Value.String()

	var query string
	query = "INSERT INTO rep_transf(block_num,tx_id,from_aid,to_aid,amount) VALUES($1,$2,$3,$4,$5/1e+18)"
	_,err := ss.db.Exec(query,agtx.BlockNum,agtx.TxId,from_aid,to_aid,amount)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_token_balance_changed_evt(evt *p.TokenBalanceChanged,block_num int64,tx_id int64) {

	market_aid := ss.Lookup_or_create_address(evt.Market.String(),block_num,tx_id)
	owner_aid := ss.Lookup_or_create_address(evt.Owner.String(),block_num,tx_id)
	token_aid := ss.Lookup_or_create_address(evt.Token.String(),block_num,tx_id)
	outcome_idx := evt.Outcome.Int64()
	balance := evt.Balance.String()

	var query string
	query = "INSERT INTO tbc(block_num,tx_id,market_aid,owner_aid,token_aid,token_type,outcome,balance) " +
				"VALUES($1,$2,$3,$4,$5,$6,$7,("+balance+"/1e+18))"
	_,err := ss.db.Exec(query,
							block_num,
							tx_id,
							market_aid,
							owner_aid,
							token_aid,
							evt.TokenType,
							outcome_idx,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v tx_id=%v q=%v",err,tx_id,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_token_transf_evt(evt *p.TokensTransferred,agtx *p.AugurTx) {

	market_aid := ss.Lookup_or_create_address(evt.Market.String(),agtx.BlockNum,agtx.TxId)
	token_aid := ss.Lookup_or_create_address(evt.Token.String(),agtx.BlockNum,agtx.TxId)
	from_aid := ss.Lookup_or_create_address(evt.From.String(),agtx.BlockNum,agtx.TxId)
	to_aid := ss.Lookup_or_create_address(evt.To.String(),agtx.BlockNum,agtx.TxId)
	value := evt.Value.String()

	var query string
	query = "INSERT INTO tok_transf(block_num,tx_id,market_aid,token_aid,from_aid,to_aid,token_type,value) " +
				"VALUES($1,$2,$3,$4,$5,$6,$7,("+value+"/1e+18))"
	_,err := ss.db.Exec(query,
							agtx.BlockNum,
							agtx.TxId,
							market_aid,
							token_aid,
							from_aid,
							to_aid,
							evt.TokenType,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) get_previous_profit_and_ff(market_aid int64,eoa_aid int64,outcome_idx int) (string,string) {

	var previous_realized_profit string
	var previous_frozen_funds string
	var query string
	query = "SELECT  round(pl.realized_profit*1e+36) AS rpl," +
					"round(pl.frozen_funds*1e+36) as ff " +
			"FROM profit_loss AS pl " +
			"WHERE  market_aid=$1 AND eoa_aid=$2 AND outcome_idx=$3 " +
			"ORDER BY pl.id DESC LIMIT 1"

	res := ss.db.QueryRow(query,market_aid,eoa_aid,outcome_idx)
	var null_pl,null_ff sql.NullString
	err := res.Scan(&null_pl,&null_ff)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	} else {
		previous_realized_profit = null_pl.String
		previous_frozen_funds = null_ff.String
	}
	return previous_realized_profit,previous_frozen_funds
}
func (ss *SQLStorage) Get_unprocessed_dai_balances() []p.DaiB {

	records := make([]p.DaiB,0,8)
	var query string
	query = "SELECT " +
				"db.id," +
				"db.aid," +
				"db.dai_transf_id," +
				"a.addr," +
				"ROUND(amount*1e+18) as amount," +
				"ROUND(balance*1e+18) as balance," +
				"db.block_num " +
			"FROM dai_bal db " +
				"LEFT JOIN address a ON db.aid=a.address_id " +
			"WHERE processed = false " +
			"ORDER by db.id " +
			"LIMIT 10"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.DaiB
		err=rows.Scan(
			&rec.Id,
			&rec.Aid,
			&rec.DaiTransfId,
			&rec.Address,
			&rec.Amount,
			&rec.Balance,
			&rec.BlockNum,
		)
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_previous_balance_from_DB(id int64,aid int64) (string,error) {

	var query string
	query = "SELECT ROUND(balance*1e+18)::text,processed FROM dai_bal " +
			"WHERE (aid=$1) and (id<$2) ORDER BY id DESC LIMIT 1"

	res := ss.db.QueryRow(query,aid,id)
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
		return "",errors.New("Unprocessed balance on past blocks")
	}
	return balance,err
}
func (ss *SQLStorage) Update_dai_token_balances_backwards(last_block_num int64,aid int64,eth_balance *big.Int) int {

	var updated_rows  int =0
	var query string

	query = "SELECT id FROM dai_bal WHERE aid=$1 AND processed=FALSE ORDER BY id DESC LIMIT 1"
	row:=ss.db.QueryRow(query,aid)
	var null_id sql.NullInt64
	var stopping_id int64 = 0
	err := row.Scan(&null_id)
	if err == nil {
		if null_id.Valid {
			stopping_id = null_id.Int64
		}
	}
	ss.Info.Printf("balance_updater(): update backwards: stopping ID=%v\n",stopping_id)

	query = "SELECT id,ROUND(balance*1e+18)::text as balance,ROUND(amount*1e+18)::text as amount,processed FROM dai_bal " +
			"WHERE " +
				"(aid = $1) AND " +
				"(block_num <= $2) " +
			"ORDER BY id DESC"
	rows,err := ss.db.Query(query,aid,last_block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	correct_balance := new(big.Int)
	correct_balance.Set(eth_balance)
	ss.Info.Printf("balance_updater(): Entering update_dai_token_balances() with eth_balance=%v correct_balace=%v\n",eth_balance.String(),correct_balance.String())
	var row_count = 0;
	defer rows.Close()
	for rows.Next() {
		row_count++
		var id int64
		var balance_str string
		var amount_str string
		var processed bool
		err = rows.Scan(&id,&balance_str,&amount_str,&processed)
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
					aid,id,correct_balance.String(),db_balance.String(),amount.String(),cmp_res)
		if cmp_res != 0 {	// incorrect balance, update it
			ss.Info.Printf("balance_updater(): incorrect balance, setting correct balance to %v for id=%v\n",
				correct_balance.String(),id)
			query = "UPDATE dai_bal " +
					"SET balance=("+correct_balance.String()+"/1e+18)," +
						"processed = true " +
					" WHERE id=$1"
			ss.Info.Printf("query = %v\n",query)
			_,err = ss.db.Exec(query,id)
			if (err!=nil) {
				p.Fatalf(fmt.Sprintf("DB Error: %v",err));
				os.Exit(1)
			}
			updated_rows++
		} else {
			if !processed {
				query = "UPDATE dai_bal " +
						"SET processed = true " +
						" WHERE id=$1"
				_,err = ss.db.Exec(query,id)
				if (err!=nil) {
					p.Fatalf(fmt.Sprintf("DB Error: %v",err));
					os.Exit(1)
				}
				updated_rows++
			}
			if id <= stopping_id {
				return updated_rows
			}
		}
	}
	if row_count == 0 {
		d_query := fmt.Sprintf("SELECT id,balance,amount,processed FROM dai_bal " +
			"WHERE " +
				"(aid = %v) AND " +
				"(block_num <= %v) " +
			"ORDER BY id DESC",aid,last_block_num)
		ss.Info.Printf("balance_updater(): query returns no rows: %v\n",d_query)
	}
	return updated_rows
}
func (ss *SQLStorage) Set_dai_balance(id int64,balance string) {

	var query string
	query = "UPDATE dai_bal SET balance = ("+balance+"/1e+18),processed=true WHERE id=$1"
	d_query := fmt.Sprintf("UPDATE dai_bal SET balance = (%v/1e+18),processed=true WHERE id=%v",balance,id)
	ss.Info.Printf("balance_updater(): Set_dai_balance: %v\n",d_query)
	_,err := ss.db.Exec(query,id)
	if (err!=nil) {
		p.Fatalf(fmt.Sprintf("DB Error: %v",err));
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_deposits_withdrawals(wallet_aid int64) []p.DaiOp{

	var query string
	query = "SELECT " +
				"db.block_num," +
//				"FLOOR(EXTRACT(EPOCH FROM b.ts))::date," +
				"b.ts::date, " +
				"db.amount as amount_float," +
				"round(db.amount,2)::text, " +
				"fa.addr AS from_addr," +
				"ta.addr AS to_addr, " +
				"dt.from_aid, " +
				"dt.to_aid " +
			"FROM dai_bal AS db " +
				"JOIN dai_transf AS dt ON db.dai_transf_id=dt.id " +
				"JOIN block AS b ON b.block_num = db.block_num " +
				"LEFT JOIN address AS fa ON dt.from_aid=fa.address_id " +
				"LEFT JOIN address AS ta ON dt.to_aid=ta.address_id " +
			"WHERE " +
				"db.aid = $1 AND " +
				"db.amount != 0 AND " +
				"db.internal = false " +
			"ORDER BY db.block_num,db.id"

	rows,err := ss.db.Query(query,wallet_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	ss.Info.Printf("Get_deposits_withdrawals: query=%v\n",query)
	records := make([]p.DaiOp,0,32)

	defer rows.Close()
	for rows.Next() {
		var rec p.DaiOp
		var amount_str string
		var amount_float float64
		var from_aid int64
		var to_aid int64
		err=rows.Scan(&rec.BlockNum,&rec.Date,&amount_float,&amount_str,&rec.FromAddr,
																&rec.ToAddr,&from_aid,&to_aid)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if amount_float < 0 {
			rec.Withdrawal = amount_str
		} else {
			rec.Deposit = amount_str
		}
		if from_aid == wallet_aid {
			rec.FromAddr = ""
		}
		if to_aid == wallet_aid {
			rec.ToAddr = ""
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) is_dai_transfer_internal(evt *p.Transfer,ca *p.ContractAddresses) (bool,bool) {

	var from_internal bool = false
	var to_internal bool = false
	_,err:=ss.lookup_market_by_addr_str(evt.From.String())
	if err == nil {
		from_internal = true
	}
	_,err=ss.lookup_market_by_addr_str(evt.To.String())
	if err == nil {
		to_internal = true
	}

	if 0 == bytes.Compare(evt.From.Bytes(),ca.Zerox.Bytes()) {
		from_internal = true
	}
	if 0 == bytes.Compare(evt.To.Bytes(),ca.Zerox.Bytes()) {
		to_internal = true
	}
	if 0 == bytes.Compare(evt.From.Bytes(),ca.FillOrder.Bytes()) {
		from_internal = true
	}
	if 0 == bytes.Compare(evt.To.Bytes(),ca.FillOrder.Bytes()) {
		to_internal = true
	}
	if 0 == bytes.Compare(evt.From.Bytes(),ca.EthXchg.Bytes()) {
		from_internal = true
	}
	if 0 == bytes.Compare(evt.To.Bytes(),ca.EthXchg.Bytes()) {
		to_internal = true
	}
	if 0 == bytes.Compare(evt.From.Bytes(),ca.ShareToken.Bytes()) {
		from_internal = true
	}
	if 0 == bytes.Compare(evt.To.Bytes(),ca.ShareToken.Bytes()) {
		to_internal = true
	}
	if 0 == bytes.Compare(evt.From.Bytes(),ca.Universe.Bytes()) {
		from_internal = true
	}
	if 0 == bytes.Compare(evt.To.Bytes(),ca.Universe.Bytes()) {
		to_internal = true
	}
	return from_internal,to_internal // its a Market in To
}
func (ss *SQLStorage) Process_DAI_token_transfer(evt *p.Transfer,ca *p.ContractAddresses,agtx *p.AugurTx) {

	from_aid := ss.Lookup_or_create_address(evt.From.String(),agtx.BlockNum,agtx.TxId)
	to_aid := ss.Lookup_or_create_address(evt.To.String(),agtx.BlockNum,agtx.TxId)
	amount := evt.Value.String()

	from_internal,to_internal := ss.is_dai_transfer_internal(evt,ca)

	var query string
	query = "INSERT INTO dai_transf(block_num,tx_id,from_aid,to_aid,amount,from_internal,to_internal) " +
			"VALUES($1,$2,$3,$4,(" + amount +"/1e+18),$5,$6)"
	_,err := ss.db.Exec(query,agtx.BlockNum,agtx.TxId,from_aid,to_aid,from_internal,to_internal)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
