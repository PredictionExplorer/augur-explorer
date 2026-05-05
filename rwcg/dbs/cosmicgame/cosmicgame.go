package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"

	p "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
)
type SQLStorageWrapper struct {
	S					*SQLStorage
}
func (sw *SQLStorageWrapper) Get_cosmic_game_contract_addrs() p.CosmicGameContractAddrs {

	var query string
	query = "SELECT "+
				"cg.cosmic_game_addr,"+
				"cg.cosmic_signature_addr,"+
				"cg.cosmic_token_addr,"+
				"cg.cosmic_dao_addr,"+
				"cg.charity_wallet_addr, "+
				"cg.prizes_wallet_addr, "+
				"cg.random_walk_addr, "+
				"cg.staking_wallet_cst_addr, "+
				"cg.staking_wallet_rwalk_addr, "+
				"cg.marketing_wallet_addr, "+
				"COALESCE((SELECT marketplace_addr FROM "+sw.S.SchemaName()+".rw_contracts LIMIT 1), '') AS marketplace_addr, "+
				"cg.implementation_addr "+
			"FROM "+sw.S.SchemaName()+".cg_contracts cg"
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
	var marketplace_addr string
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
		&marketplace_addr,
		&implementation_addr,
	);
	if (err!=nil) {
		err_msg := fmt.Sprintf("Error in Get_cosmic_game_contract_addrs(): %v, q=%v",err,query)
		sw.S.Log_msg(err_msg)
		fmt.Printf("\nFATAL: %s\n", err_msg)
		fmt.Printf("HINT: If you don't need CosmicGame, set ENABLE_ROUTES_COSMICGAME=false in websrv .env\n\n")
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
	output.MarketplaceAddr = marketplace_addr
	output.ImplementationAddr = implementation_addr
	return output
}
func (sw *SQLStorageWrapper) Get_cosmic_game_processing_status() p.CosmicGameProcStatus {

	var output p.CosmicGameProcStatus
	var null_id sql.NullInt64
	var null_block sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id, last_block_num FROM "+sw.S.SchemaName()+".cg_proc_status"

		res := sw.S.Db().QueryRow(query)
		err := res.Scan(&null_id, &null_block)
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
	if null_block.Valid {
		output.LastBlockNum = null_block.Int64
	}
	return output
}
func (sw *SQLStorageWrapper) Update_cosmic_game_process_status(status *p.CosmicGameProcStatus) {

	var query string
	query = "UPDATE "+sw.S.SchemaName()+".cg_proc_status SET last_evt_id = $1, last_block_num = $2"

	_,err := sw.S.Db().Exec(query, status.LastEvtIdProcessed, status.LastBlockNum)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
