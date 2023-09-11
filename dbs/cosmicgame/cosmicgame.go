package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"

	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
type SQLStorageWrapper struct {
	S					*SQLStorage
}
func (sw *SQLStorageWrapper) Get_cosmic_game_contract_addrs() p.CosmicGameContractAddrs {

	var query string
	query = "SELECT "+
				"cosmic_game_addr,"+
				"cosmic_signature_addr,"+
				"cosmic_token_addr,"+
				"cosmic_dao_addr,"+
				"charity_wallet_addr, "+
				"raffle_wallet_addr, "+
				"random_walk_addr "+
			"FROM "+sw.S.SchemaName()+".cg_contracts"
	row := sw.S.Db().QueryRow(query)
	var cosmic_game_addr string
	var cosmic_signature_addr string
	var cosmic_token_addr string
	var cosmic_dao_addr string
	var charity_wallet_addr string
	var raffle_wallet_addr string
	var random_walk_addr string
	var err error
	err=row.Scan(
		&cosmic_game_addr,
		&cosmic_signature_addr,
		&cosmic_token_addr,
		&cosmic_dao_addr,
		&charity_wallet_addr,
		&raffle_wallet_addr,
		&random_walk_addr,
	);
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_contract_addrs(): %v, q=%v",err,query))
		os.Exit(1)
	}
	var output p.CosmicGameContractAddrs
	output.CosmicGameAddr = cosmic_game_addr
	output.CosmicSignatureAddr = cosmic_signature_addr
	output.CosmicTokenAddr = cosmic_token_addr
	output.CosmicDaoAddr = cosmic_dao_addr
	output.CharityWalletAddr = charity_wallet_addr
	output.RaffleWalletAddr = raffle_wallet_addr
	output.RandomWalkAddr = random_walk_addr
	return output
}
func (sw *SQLStorageWrapper) Get_cosmic_game_processing_status() p.CosmicGameProcStatus {

	var output p.CosmicGameProcStatus
	var null_id sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM "+sw.S.SchemaName()+".cg_proc_status"

		res := sw.S.Db().QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO "+sw.S.SchemaName()+".cg_proc_status DEFAULT VALUES"
				_,err := sw.S.Db().Exec(query)
				if (err!=nil) {
					sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
				os.Exit(1)
			}
		} else {
			break
		}
	}
	if null_id.Valid {
		output.LastEvtIdProcessed = null_id.Int64
	}
	return output
}
func (sw *SQLStorageWrapper) Update_cosmic_game_process_status(status *p.CosmicGameProcStatus) {

	var query string
	query = "UPDATE "+sw.S.SchemaName()+".cg_proc_status SET last_evt_id = $1"

	_,err := sw.S.Db().Exec(query,status.LastEvtIdProcessed)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Get_cosmic_game_bid_by_evtlog_id(bid_evtlog_id int64) int64 {

	var query string
	query = "SELECT id FROM "+sw.S.SchemaName()+".cg_bid WHERE evtlog_id=$1"
	res := sw.S.Db().QueryRow(query,bid_evtlog_id)
	var null_id sql.NullInt64
	err := res.Scan(&null_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0	// if bid wasn't found there wasn't any bid but pure Donate() instead,
						//	so we return 0 as Id
		}
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return null_id.Int64
}
func (sw *SQLStorageWrapper) Get_donation_received_evt_id(tx_id,starting_id int64,sig string) int64 {

	var query string 
	query = "SELECT "+
				"d.evtlog_id "+
			"FROM "+
				"evt_log e "+
				"LEFT JOIN cg_donation_received d ON e.id=d.evtlog_id "+
			"WHERE "+
				"(e.tx_id=$1) AND "+
				"(e.topic0_sig=$2) AND "+
				"(e.id<$3) "+
			"ORDER BY e.id DESC LIMIT 1"
	res := sw.S.Db().QueryRow(query,tx_id,sig,starting_id)
	var null_id sql.NullInt64
	err := res.Scan(&null_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return null_id.Int64
}
