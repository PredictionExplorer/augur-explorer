package biddingwar

import (

	p "github.com/PredictionExplorer/augur-explorer/primitives/biddinwar"
)
type SQLStorageWrapper struct {
	S					*SQLStorage
}
func (sw *SQLStorageWrapper) BiddingWar_get_contract_addrs() p.BiddingWarContractAddrs {

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
	var output p.BiddingWarContractaddrs
	output.BiddingWarAddr = bidding_war_addr
	output.CosmicSignatureAddr = cosmic_signature_addr
	output.CosmicTokenAddr = cosmic_token_addr
	output.CharityWalletAddr = charity_wallet_addr
	return output
}
