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
				"prizes_wallet_addr, "+
				"random_walk_addr, "+
				"staking_wallet_cst_addr, "+
				"staking_wallet_rwalk_addr, "+
				"marketing_wallet_addr, "+
				"implementation_addr "+
			"FROM "+sw.S.SchemaName()+".cg_contracts"
	row := sw.S.Db().QueryRow(query)
	var cosmic_game_addr string
	var cosmic_signature_addr string
	var cosmic_token_addr string
	var cosmic_dao_addr string
	var charity_wallet_addr string
	var prizes_wallet_addr string
	var random_walk_addr string
	var staking_wallet_cst_addr string
	var staking_wallet_rwalk_addr string
	var marketing_wallet_addr string
	var implementation_addr string
	var err error
	err=row.Scan(
		&cosmic_game_addr,
		&cosmic_signature_addr,
		&cosmic_token_addr,
		&cosmic_dao_addr,
		&charity_wallet_addr,
		&prizes_wallet_addr,
		&random_walk_addr,
		&staking_wallet_cst_addr,
		&staking_wallet_rwalk_addr,
		&marketing_wallet_addr,
		&implementation_addr,
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
	output.PrizesWalletAddr = prizes_wallet_addr
	output.RandomWalkAddr = random_walk_addr
	output.StakingWalletCSTAddr = staking_wallet_cst_addr
	output.StakingWalletRWalkAddr = staking_wallet_rwalk_addr
	output.MarketingWalletAddr = marketing_wallet_addr
	output.ImplementationAddr = implementation_addr
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
