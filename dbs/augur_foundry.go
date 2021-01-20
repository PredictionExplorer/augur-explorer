package dbs

import (
	"fmt"
	"os"
	"encoding/hex"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_augur_foundry_contract_addr() string {

	var query string
	query = "SELECT augur_foundry_addr FROM af_addr"
	var addr string
	row := ss.db.QueryRow(query)
	err := row.Scan(&addr)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Error Augur Foundry contract address is not set: %v",err))
			os.Exit(1)
		}
		ss.Log_msg(fmt.Sprintf("Error in Get_upload_block(): %v",err))
		os.Exit(1)
	}
	return addr
}
func (ss *SQLStorage) Get_augur_foundry_status() p.AugurFoundryStatus {

	var output p.AugurFoundryStatus
	var null_id sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM af_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO af_status DEFAULT VALUES"
				_,err := ss.db.Exec(query)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
				os.Exit(1)
			}
		} else {
			break
		}
	}
	if null_id.Valid {
		output.LastEvtId = null_id.Int64
	}
	return output
}
func (ss *SQLStorage) Update_augur_foundry_status(status *p.AugurFoundryStatus) {

	var query string
	query = "UPDATE af_status SET last_evt_id = $1"

	_,err := ss.db.Exec(query,status.LastEvtId)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_wrapper_created_evt(evtlog_id int64,timestamp int64,agtx *p.AugurTx,evt *p.EAugurFoundryWrapperCreated,name,symbol string,decimals int,market_addr *common.Address,outcome_idx uint8) {

	token_id := hex.EncodeToString(evt.TokenId.Bytes())
	wrapper_aid := ss.Lookup_or_create_address(evt.TokenAddress.String(),agtx.BlockNum,agtx.TxId)
	market_aid := ss.Lookup_or_create_address(market_addr.String(),0,0)
	var query string
	query = "INSERT INTO af_wrapper (" +
				"evtlog_id,block_num,tx_id,token_id,wrapper_aid,time_stamp,name,symbol,decimals,"+
				"market_aid,outcome_idx" +
			") VALUES ($1,$2,$3,$4,$5,TO_TIMESTAMP($6),$7,$8,$9,$10,$11)"

	_,err := ss.db.Exec(query,
		evtlog_id,
		agtx.BlockNum,
		agtx.TxId,
		token_id,
		wrapper_aid,
		timestamp,
		name,
		symbol,
		decimals,
		market_aid,
		outcome_idx,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Delete_wrapper_created_evt(evtlog_id int64) {

	var query string
	query = "DELETE FROM af_wrapper WHERE evtlog_id=$1"
	_,err := ss.db.Exec(query,evtlog_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_erc20wrapped_sharetoken_contracts() []p.ERC20ShTokContract {

	records := make([]p.ERC20ShTokContract,0,32)
	var query string
	query = "SELECT " +
				"wrapper_aid,market_aid,last_evt_id,outcome_idx,decimals " +
			"FROM af_wrapper " +
			"ORDER BY block_num,tx_id"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.ERC20ShTokContract
		err=rows.Scan(
			&rec.WrapperAid,
			&rec.MarketAid,
			&rec.LastEvtId,
			&rec.OutcomeIdx,
			&rec.Decimals,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Update_wrapped_token_event_id_status(wrapper_aid,evtlog_id int64) {

	var query string
	query = "UPDATE af_wrapper SET last_evt_id = $2 WHERE wrapper_aid=$1"

	_,err := ss.db.Exec(query,wrapper_aid,evtlog_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_augur_foundry_transfer_evt(t *p.WShTokTransfer,decimals int) {

	from_aid := ss.Lookup_or_create_address(t.From,t.BlockNum,t.TxId)
	to_aid := ss.Lookup_or_create_address(t.To,t.BlockNum,t.TxId)
	var query string
	query = "INSERT INTO wstok_transf( " +
				"evtlog_id,wrapper_aid,block_num,tx_id,time_stamp,from_aid,to_aid,amount" +
			") VALUES($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,($8::DECIMAL/1e+"+fmt.Sprintf("%v",decimals)+"))"

	_,err := ss.db.Exec(query,
		t.EvtLogId,
		t.WrapperAid,
		t.BlockNum,
		t.TxId,
		t.TimeStamp,
		from_aid,
		to_aid,
		t.AmountStr,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_wrapped_tokens_for_market(market_aid int64) []p.ERC20ShTokContract {

	records := make([]p.ERC20ShTokContract,0,4)
	var query string
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM w.time_stamp))::BIGINT AS ts, " +
				"w.wrapper_aid," +
				"w.outcome_idx," +
				"wa.addr," +
				"w.decimals," +
				"w.name, " +
				"w.symbol " +
			"FROM af_wrapper AS w " +
				"LEFT JOIN address AS wa ON w.wrapper_aid=wa.address_id " +
			"WHERE w.market_aid=$1 " +
			"ORDER BY outcome_idx"

	rows,err := ss.db.Query(query,market_aid)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.ERC20ShTokContract
		rec.MarketAid=market_aid
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.WrapperAid,
			&rec.OutcomeIdx,
			&rec.Address,
			&rec.Decimals,
			&rec.Name,
			&rec.Symbol,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_wrapped_token_info(wrapper_aid int64) (p.ERC20ShTokContract,error) {

	var output p.ERC20ShTokContract
	var query string
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM w.time_stamp))::BIGINT AS ts, " +
				"w.wrapper_aid," +
				"w.market_aid," +
				"w.outcome_idx," +
				"wa.addr," +
				"ma.addr," +
				"w.decimals," +
				"w.name, " +
				"w.symbol " +
			"FROM af_wrapper AS w " +
				"LEFT JOIN address AS wa ON w.wrapper_aid=wa.address_id " +
				"LEFT JOIN address AS ma ON w.market_aid=ma.address_id " +
			"WHERE w.wrapper_aid=$1 " +
			"ORDER BY outcome_idx"


	row := ss.db.QueryRow(query,wrapper_aid)
	err := row.Scan(
			&output.TimeStamp,
			&output.WrapperAid,
			&output.MarketAid,
			&output.OutcomeIdx,
			&output.Address,
			&output.MktAddr,
			&output.Decimals,
			&output.Name,
			&output.Symbol,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output,err
		} else {
			ss.Log_msg(fmt.Sprintf("Error Augur Foundry contract address is not set: %v",err))
			os.Exit(1)
		}
	}
	return output,nil
}
func (ss *SQLStorage) Get_wrapped_token_transfers(wrapper_aid int64,offset,limit int) ([]p.WShTokTransfer,int64) {

	records := make([]p.WShTokTransfer,0,64)
	var query string

	var total_rows int64
	query = "SELECT COUNT(*) AS total_rows FROM wstok_transf t " +
			"WHERE t.wrapper_aid=$1"

	var null_counter sql.NullInt64
	var err error
	err=ss.db.QueryRow(query,wrapper_aid).Scan(&null_counter);
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v , q=%v",err,query))
			os.Exit(1)
		}
	}
	total_rows=null_counter.Int64

	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM b.ts))::BIGINT AS ts, " +
				"t.block_num," +
				"t.tx_id,"+
				"fbp.pool_aid,"+
				"tbp.pool_aid,"+
				"fa.addr," +
				"ta.addr," +
				"t.amount " +
			"FROM wstok_transf AS t " +
				"JOIN block AS b ON t.block_num=b.block_num " +
				"LEFT JOIN address AS fa ON t.from_aid=fa.address_id " +
				"LEFT JOIN address AS ta ON t.to_aid=ta.address_id " +
				"LEFT JOIN bpool AS fbp ON t.from_aid=fbp.pool_aid " +
				"LEFT JOIN bpool AS tbp ON t.to_aid=tbp.pool_aid " +
			"WHERE t.wrapper_aid=$1 " +
			"ORDER BY ts DESC " +
			"OFFSET $2 LIMIT $3"
	rows,err := ss.db.Query(query,wrapper_aid,offset,limit)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.WShTokTransfer
		var null_from_pool,null_to_pool sql.NullInt64
		rec.WrapperAid=wrapper_aid
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.BlockNum,
			&rec.TxId,
			&null_from_pool,
			&null_to_pool,
			&rec.From,
			&rec.To,
			&rec.Amount,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		if null_from_pool.Valid {
			rec.FromPool = true
		}
		if null_to_pool.Valid {
			rec.ToPool = true
		}
		if (!rec.FromPool) && (!rec.ToPool) {
			rec.NonPoolTransfer = true
		}
		records = append(records,rec)
	}
	return records,total_rows
}
func (ss *SQLStorage) Get_augur_foundry_wrapper_list() []p.ERC20ShTokContract {

	records := make([]p.ERC20ShTokContract,0,4)
	var query string
	query = "SELECT " +
				"w.wrapper_aid, " +
				"wa.addr, " +
				"ma.addr," +
				"FLOOR(EXTRACT(EPOCH FROM w.time_stamp))::BIGINT AS ts, " +
				"w.time_stamp," +
				"w.outcome_idx," +
				"w.market_aid," +
				"m.extra_info::json->>'description'," +
				"m.market_type," +
				"m.outcomes," +
				"w.name," +
				"w.symbol, " +
				"w.decimals " +
			"FROM " +
				"af_wrapper AS w " +
				"JOIN address AS wa ON w.wrapper_aid=wa.address_id " +
				"LEFT JOIN market AS m ON m.market_aid=w.market_aid " +
				"LEFT JOIN address AS ma on w.market_aid=ma.address_id " +
			"ORDER BY w.time_stamp "

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var rec p.ERC20ShTokContract
		var market_type int64
		var null_outcomes sql.NullString
		err=rows.Scan(
			&rec.WrapperAid,
			&rec.Address,
			&rec.MktAddr,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.OutcomeIdx,
			&rec.MarketAid,
			&rec.MktDescr,
			&market_type,
			&null_outcomes,
			&rec.Name,
			&rec.Symbol,
			&rec.Decimals,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		rec.Outcome = get_outcome_str(uint8(market_type),rec.OutcomeIdx,&null_outcomes.String)
		records = append(records,rec)
	}
	return records
}
