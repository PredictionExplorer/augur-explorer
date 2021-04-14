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
