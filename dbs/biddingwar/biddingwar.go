package biddingwar

import (
	"os"
	"fmt"
	"database/sql"

	p "github.com/PredictionExplorer/augur-explorer/primitives/biddingwar"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
type SQLStorageWrapper struct {
	S					*SQLStorage
}
func (sw *SQLStorageWrapper) Get_biddingwar_contract_addrs() p.BiddingWarContractAddrs {

	var query string
	query = "SELECT "+
				"bidding_war_addr,"+
				"cosmic_signature_addr,"+
				"cosmic_token_addr,"+
				"charity_wallet_addr "+
			"FROM "+sw.S.SchemaName()+".bw_contracts"
	row := sw.S.Db().QueryRow(query)
	var bidding_war_addr string
	var cosmic_signature_addr string
	var cosmic_token_addr string
	var charity_wallet_addr string
	var err error
	err=row.Scan(
		&bidding_war_addr,
		&cosmic_signature_addr,
		&cosmic_token_addr,
		&charity_wallet_addr,
	);
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in BiddingWar_get_contract_addrs(): %v, q=%v",err,query))
		os.Exit(1)
	}
	var output p.BiddingWarContractAddrs
	output.BiddingWarAddr = bidding_war_addr
	output.CosmicSignatureAddr = cosmic_signature_addr
	output.CosmicSignatureTokenAddr = cosmic_token_addr
	output.CharityWalletAddr = charity_wallet_addr
	return output
}
func (sw *SQLStorageWrapper) Get_biddingwar_processing_status() p.BiddingWarProcStatus {

	var output p.BiddingWarProcStatus
	var null_id sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id FROM "+sw.S.SchemaName()+".bw_proc_status"

		res := sw.S.Db().QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO "+sw.S.SchemaName()+".bw_proc_status DEFAULT VALUES"
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
func (sw *SQLStorageWrapper) Update_biddingwar_process_status(status *p.BiddingWarProcStatus) {

	var query string
	query = "UPDATE "+sw.S.SchemaName()+".bw_proc_status SET last_evt_id = $1"

	_,err := sw.S.Db().Exec(query,status.LastEvtIdProcessed)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
