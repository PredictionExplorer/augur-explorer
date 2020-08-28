package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_contract_addresses() (p.ContractAddresses,error) {

	var query string
	query="SELECT " +
				"augur,augur_trading,profit_loss,dai_cash,rep_token,zerox_trade,zerox_xchg,wallet_reg,"+
				"wallet_reg2,fill_order,eth_xchg,share_token,universe,create_order "+
			"FROM contract_addresses";
	row := ss.db.QueryRow(query)
	var c_addresses p.ContractAddresses
	var err error
	var (
		augur string
		augur_trading string
		pl string
		dai string
		rep string
		zerox_trade string
		zerox_xchg string
		wallet_reg string
		wallet_reg2 string
		fill_order string
		eth_xchg string
		share_token string
		universe string
		create_order string
	)
	err=row.Scan(
		&augur,&augur_trading,&pl,&dai,&rep,&zerox_trade,&zerox_xchg,&wallet_reg,&wallet_reg2,&fill_order,
		&eth_xchg,&share_token,&universe,&create_order,
	);
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_contract_addresses(): %v",err))
			os.Exit(1)
		}
		return c_addresses,err
	}
	c_addresses.Augur=common.HexToAddress(augur)
	c_addresses.AugurTrading=common.HexToAddress(augur_trading)
	c_addresses.PL=common.HexToAddress(pl)
	c_addresses.Dai=common.HexToAddress(dai)
	c_addresses.Reputation=common.HexToAddress(rep)
	c_addresses.ZeroxTrade=common.HexToAddress(zerox_trade)
	c_addresses.ZeroxXchg=common.HexToAddress(zerox_xchg)
	c_addresses.WalletReg=common.HexToAddress(wallet_reg)
	c_addresses.WalletReg2=common.HexToAddress(wallet_reg2)
	c_addresses.FillOrder=common.HexToAddress(fill_order)
	c_addresses.EthXchg=common.HexToAddress(eth_xchg)
	c_addresses.ShareToken=common.HexToAddress(share_token)
	c_addresses.Universe= common.HexToAddress(universe)
	c_addresses.CreateOrder=common.HexToAddress(create_order)
	return c_addresses,nil
}
func (ss *SQLStorage) Get_augur_blocks(market_aid int64) []int64 {

	var where_cond string = ""
	if market_aid > 0 {
		where_cond = fmt.Sprintf(" WHERE market_aid = %v ",market_aid)
	}
	var query string = ""
	query = "SELECT DISTINCT block_num FROM (" +
				"(SELECT DISTINCT block_num FROM mktord " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM market " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM mkt_fin " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM claim_funds " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM sbalances " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM volume " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM oi_chg " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM tbc " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM profit_loss " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM report " + where_cond + ") " +
			") AS block_numbers ORDER BY block_num"
	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]int64,0,4096)

	defer rows.Close()
	for rows.Next() {
		var block_num int64
		err=rows.Scan(&block_num)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,block_num)
	}
	return records
}
func (ss *SQLStorage) Get_upload_block() int64 {

	var err error
	var block_num int64
	var query string
	query="SELECT upload_block FROM contract_addresses";
	row := ss.db.QueryRow(query)
	err=row.Scan(&block_num)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Error in Get_upload_block(): %v",err))
		os.Exit(1)
	}
	return block_num
}
func (ss *SQLStorage) Insert_augur_transaction_status(agtx *p.AugurTx,evt *p.EExecuteTransactionStatus) {

	//eoa_aid,err := ss.Lookup_eoa_aid(wallet_aid)
	eoa_aid:=0
	wallet_aid:=0
	var query string
	query = "INSERT INTO agtx_status(block_num,tx_id,eoa_aid,wallet_aid,success,funding_success) " +
			"VALUES($1,$2,$3,$4,$5,$6)"
	_,err := ss.db.Exec(query,agtx.BlockNum,agtx.TxId,eoa_aid,wallet_aid,evt.Success,evt.FundingSuccess)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into agtx_status table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Register_eoa_and_wallet(eoa,wallet string,block_num int64,tx_id int64) {

	eoa_aid := ss.Lookup_or_create_address(eoa,block_num,tx_id)
	wallet_aid := ss.Lookup_or_create_address(wallet,block_num,tx_id)
	ss.Info.Printf(
		"Registering eoa=%v (eoa_aid=%v) with wallet %v (wallet_aid=%v)\n",
		eoa,eoa_aid,wallet,wallet_aid,
	)
	ss.Link_eoa_and_wallet_contract(eoa_aid,wallet_aid)
}
func (ss *SQLStorage) Insert_execute_wallet_tx(eoa_aid int64,wallet_aid int64,agtx *p.AugurTx,wtx *p.ExecuteWalletTx) {

	to_aid := ss.Lookup_or_create_address(agtx.To,agtx.BlockNum,agtx.TxId)
	var referral_aid int64  = 0
	if wtx.ReferralAddress != "0x0000000000000000000000000000000000000000" {
		referral_aid = ss.Lookup_or_create_address(wtx.ReferralAddress,agtx.BlockNum,agtx.TxId)
	}
	var query string
	query = "INSERT INTO exec_wtx(" +
				"block_num,tx_id,"+
				"eoa_aid,wallet_aid,to_aid,referral_aid,"+
				"value,payment,desired_signer_bal,"+
				"input_sig,"+
				"fingerprint,"+
				"call_data,"+
				"revert_on_failure" +
			") VALUES (" +
				"$1,$2,"+
				"$3,$4,$5,$6," +
				wtx.Value+"/1e+18,"+wtx.Payment+"/1e+18,"+wtx.DesiredSignerBalance+"/1e+18,"+
				"$7,$8,$9,$10"+
			")"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,agtx.TxId,
		eoa_aid,wallet_aid,to_aid,referral_aid,
		wtx.InputSig,
		wtx.Fingerprint,
		wtx.CallData,
		wtx.RevertOnFailure,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into agtx_status table: %v; q=%v",err,query))
		os.Exit(1)
	}

}
