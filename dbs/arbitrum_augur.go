package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_arbitrum_augur_contract_addresses() (p.AA_ContractAddrs,error) {

	var query string
	query="SELECT " +
				"amm_factory,hatchery_reg "+
			"FROM aa_caddrs";
	row := ss.db.QueryRow(query)
	var c_addrs p.AA_ContractAddrs
	var err error
	var (
		amm_factory string
		hatchery_reg string
	)
	err=row.Scan(
		&amm_factory,&hatchery_reg,
	);
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_arbitrum_augur_contract_addresses(): %v",err))
			os.Exit(1)
		}
		return c_addrs,err
	}
	c_addrs.AMM_Factory=common.HexToAddress(amm_factory)
	c_addrs.HatcheryRegistry=common.HexToAddress(hatchery_reg)
	return c_addrs,nil
}
func (ss *SQLStorage) Update_arbitrum_augur_process_status(status *p.ArbitrumAugurProcessStatus) {

	var query string
	query = "UPDATE aa_proc_status SET last_evt_id = $1"

	_,err := ss.db.Exec(query,status.LastEvtId)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_pool_created_event(evt *p.AA_PoolCreated) {


	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	pool_aid:=ss.Lookup_or_create_address(evt.PoolAddr,evt.BlockNum,evt.TxId)
	hatchery_aid:=ss.Lookup_or_create_address(evt.HatcheryAddr,evt.BlockNum,evt.TxId)
	creator_aid:=ss.Lookup_or_create_address(evt.CreatorAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_pool_created(" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"pool_aid,hatchery_aid,creator_aid,turbo_id" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			pool_aid,
			hatchery_aid,
			creator_aid,
			evt.TurboId,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_pool_created table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_arbitrum_augur_processing_status() p.ArbitrumAugurProcessStatus {

	var output p.ArbitrumAugurProcessStatus
	var null_id sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM aa_proc_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO aa_proc_status DEFAULT VALUES"
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
func (ss *SQLStorage) Insert_aa_new_hatchery_event(evt *p.AA_NewHatchery) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	hatchery_aid:=ss.Lookup_or_create_address(evt.HatcheryAddr,evt.BlockNum,evt.TxId)
	collateral_aid:=ss.Lookup_or_create_address(evt.CollateralAddr,evt.BlockNum,evt.TxId)
	shtok_aid:=ss.Lookup_or_create_address(evt.ShareTokenAddr,evt.BlockNum,evt.TxId)
	feepot_aid:=ss.Lookup_or_create_address(evt.FeePotAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_new_hatchery(" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"hatchery_aid,collateral_aid,shtok_aid,feepot_aid" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8,$9)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			hatchery_aid,
			collateral_aid,
			shtok_aid,
			feepot_aid,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_new_hatchery table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_turbo_created_event(evt *p.AA_TurboCreated) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	arbiter_aid:=ss.Lookup_or_create_address(evt.ArbiterAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_turbo_created (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"num_ticks,arbiter_aid,creator_fee,tindex,outcome_symbols,outcome_names,arbiter_config" +
				") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8::DECIMAL/1e+19,$9,$10,$11,$12)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.NumTicks,
			arbiter_aid,
			evt.CreatorFee,
			evt.Index,
			evt.OutcomeSymbols,
			evt.OutcomeNames,
			evt.ArbiterConfiguration,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_turbo_created table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_complete_sets_minted_event(evt *p.AA_CompleteSetsMinted) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	aid:=ss.Lookup_or_create_address(evt.TargetAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_sets_minted (" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"aid,turbo_id,amount" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			aid,
			evt.TurboId,
			evt.Amount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_sets_minted table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_complete_sets_burned_event(evt *p.AA_CompleteSetsBurned) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	aid:=ss.Lookup_or_create_address(evt.TargetAddr,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_sets_burned(" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"aid,turbo_id,amount" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			aid,
			evt.TurboId,
			evt.Amount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_sets_burned table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_aa_claim_event(evt *p.AA_Claim) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_claim(" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,turbo_id" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			evt.TurboId,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_claim table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_arbitrum_augur_pools() []p.AA_Pool {

	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM time_stamp)::BIGINT AS created_ts, " +
				"time_stamp," +
				"p.block_num, " +
				"tx.tx_hash," +
				"pa.addr," +
				"ha.addr," +
				"ca.addr," +
				"turbo_id " +
			"FROM aa_pool_created AS p " +
				"LEFT JOIN address pa ON p.pool_aid=pa.address_id " +
				"LEFT JOIN address ha ON p.hatchery_aid=ha.address_id " +
				"LEFT JOIN address ca ON p.creator_aid=ca.address_id " +
				"JOIN transaction tx ON p.tx_id=tx.id " +
			"ORDER BY p.time_stamp"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.AA_Pool,0,32)

	defer rows.Close()
	for rows.Next() {
		var rec p.AA_Pool
		err=rows.Scan(
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.TxHash,
			&rec.PoolAddr,
			&rec.HatcheryAddr,
			&rec.CreatorAddr,
			&rec.TurboId,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records

}
func (ss *SQLStorage) Insert_aa_feepot_transfer_event(evt *p.AA_FeePotTransfer) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	from_aid:=ss.Lookup_or_create_address(evt.From,evt.BlockNum,evt.TxId)
	to_aid:=ss.Lookup_or_create_address(evt.To,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO aa_feepot_trsf(" +
				"evtlog_id,block_num,tx_id,contract_aid,time_stamp,"+
				"from_aid,to_aid,value" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7,$8::DECIMAL/1e+18)"

	_,err := ss.db.Exec(query,
			evt.EvtId,
			evt.BlockNum,
			evt.TxId,
			contract_aid,
			evt.TimeStamp,
			from_aid,
			to_aid,
			evt.Value,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into aa_sets_burned table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Is_feepot(addr string) bool {

	var query string
	query = "SELECT feepot_aid FROM aa_new_hatchery h "+
			"JOIN address a ON h.feepot_aid=a.address_id "+
			"WHERE a.addr=$1"+
			"LIMIT 1"
	row := ss.db.QueryRow(query,addr)
	var null_id sql.NullInt64
	err := row.Scan(&addr)
	if (err!=nil) {
		if err==sql.ErrNoRows {
			return false
		}
		ss.Log_msg(fmt.Sprintf("Error in Is_feepot(): %v",err))
		os.Exit(1)
	}
	_=null_id
	return true
}
