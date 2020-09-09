package dbs

import (
	"fmt"
	"os"
	"bytes"
	"strings"
	"encoding/hex"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_contract_addresses() (p.ContractAddresses,error) {

	var query string
	query="SELECT " +
				"chain_id,"+
				"augur,augur_trading,profit_loss,dai_cash,rep_token,zerox_trade,zerox_xchg,wallet_reg,"+
				"wallet_reg2,fill_order,eth_xchg,share_token,universe,create_order, "+
				"leg_rep_token,buy_part_tok,redeem_stake,warp_sync,hot_loading,affiliates," +
				"affiliate_val,ctime,cancel_order,orders,sim_trade,trade,oi_cash,uniswap_v2_fact,"+
				"uniswap_v2_r2,audit_funds,weth9,usdc,usdt,relay_hub_v2,account_loader " +
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
		leg_rep_token string
		buy_part_tok string
		redeem_stake string
		warp_sync string
		hot_loading string
		affiliates string
		affiliate_val string
		ctime string
		cancel_order string
		orders string
		sim_trade string
		trade string
		oi_cash string
		uniswap_v2_fact string
		uniswap_v2_r2 string
		audit_funds string
		weth9 string
		usdc string
		usdt string
		relay_hub_v2 string
		account_loader string
	)
	err=row.Scan(
		&c_addresses.ChainId,
		&augur,&augur_trading,&pl,&dai,&rep,&zerox_trade,&zerox_xchg,&wallet_reg,&wallet_reg2,&fill_order,
		&eth_xchg,&share_token,&universe,&create_order,&leg_rep_token,&buy_part_tok,&redeem_stake,
		&warp_sync,&hot_loading,&affiliates,&affiliate_val,&ctime,&cancel_order,&orders,&sim_trade,&trade,
		&oi_cash,&uniswap_v2_fact,&uniswap_v2_r2,&audit_funds,&weth9,&usdc,&usdt,&relay_hub_v2,
		&account_loader,
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
	c_addresses.GenesisUniverse= common.HexToAddress(universe)
	c_addresses.CreateOrder=common.HexToAddress(create_order)
	c_addresses.LegacyReputationToken=common.HexToAddress(leg_rep_token)
	c_addresses.BuyParticipationTokens=common.HexToAddress(buy_part_tok)
	c_addresses.RedeemStake=common.HexToAddress(redeem_stake)
	c_addresses.WarpSync=common.HexToAddress(warp_sync)
	c_addresses.HotLoading=common.HexToAddress(hot_loading)
	c_addresses.Affiliates=common.HexToAddress(affiliates)
	c_addresses.AffiliateValidator=common.HexToAddress(affiliate_val)
	c_addresses.Time=common.HexToAddress(ctime)
	c_addresses.CancelOrder=common.HexToAddress(cancel_order)
	c_addresses.Orders=common.HexToAddress(orders)
	c_addresses.SimulateTrade=common.HexToAddress(sim_trade)
	c_addresses.Trade=common.HexToAddress(trade)
	c_addresses.OICash=common.HexToAddress(oi_cash)
	c_addresses.UniswapV2Factory=common.HexToAddress(uniswap_v2_fact)
	c_addresses.UniswapV2Router02=common.HexToAddress(uniswap_v2_r2)
	c_addresses.AuditFunds=common.HexToAddress(audit_funds)
	c_addresses.WETH9=common.HexToAddress(weth9)
	c_addresses.USDC=common.HexToAddress(usdc)
	c_addresses.USDT=common.HexToAddress(usdt)
	c_addresses.RelayHubV2=common.HexToAddress(relay_hub_v2)
	c_addresses.AccountLoader=common.HexToAddress(account_loader)
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
func (ss *SQLStorage) Insert_register_contract_event(agtx *p.AugurTx,evt *p.ERegisterContract) {

	length := bytes.Index(evt.Key[:],[]byte{0})
	key := string(evt.Key[:length])

	var query string
	query = "INSERT INTO register_contract(" +
				"block_num,tx_id,addr,key" +
			") VALUES ($1,$2,$3,$4)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,agtx.TxId,evt.ContractAddress.String(),key,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into register_contract table: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_universe_created_event(agtx *p.AugurTx,evt *p.EUniverseCreated) {

	payout_numerators:=p.Bigint_ptr_slice_to_str(&evt.PayoutNumerators,",")
	parent_id:=ss.Lookup_or_create_address(evt.ParentUniverse.String(),agtx.BlockNum,agtx.TxId)
	universe_id:=ss.Lookup_or_create_address(evt.ChildUniverse.String(),agtx.BlockNum,agtx.TxId)

	var query string
	query = "INSERT INTO universe(" +
				"block_num,tx_id,universe_id,parent_id,creation_ts,universe_addr,payout_numerators" +
			") VALUES ($1,$2,$3,$4,TO_TIMESTAMP($5),$6,$7)"

	_,err := ss.db.Exec(query,
		agtx.BlockNum,
		agtx.TxId,
		universe_id,
		parent_id,
		evt.CreationTimestamp.Int64(),
		evt.ChildUniverse.String(),
		payout_numerators,
	)
	if err!= nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert Universe : %v; q=%v",err,query))
		os.Exit(1)
	}
	query = "INSERT INTO main_stats(universe_id) VALUES($1)"
	_,err = ss.db.Exec(query,universe_id)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert record in main_stats table: %v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Update_contract_addresses(caddrs *p.ContractAddresses) {

	var query string
	query = "UPDATE contract_addresses SET "+
				"augur = '"				+ caddrs.Augur.String() + "'," +
				"augur_trading = '"		+ caddrs.AugurTrading.String() + "'," +
				"profit_loss = '"		+ caddrs.PL.String() + "'," +
				"dai_cash = '"			+ caddrs.Dai.String() + "'," +
				"zerox_trade = '"		+ caddrs.ZeroxTrade.String() + "'," +
				"zerox_xchg = '"		+ caddrs.ZeroxXchg.String() + "'," +
				"rep_token = '"			+ caddrs.Reputation.String() + "'," +
				"wallet_reg = '"		+ caddrs.WalletReg.String() + "'," +
				"wallet_reg2 = '"		+ caddrs.WalletReg2.String() + "'," +
				"fill_order = '"		+ caddrs.FillOrder.String() + "'," +
				"eth_xchg = '"			+ caddrs.EthXchg.String() + "'," +
				"share_token = '"		+ caddrs.ShareToken.String() + "'," +
				"universe = '"			+ caddrs.GenesisUniverse.String() + "',"+
				"create_order = '"		+ caddrs.CreateOrder.String() + "'," +
				"leg_rep_token = '"		+ caddrs.LegacyReputationToken.String() + "'," +
				"buy_part_tok = '"		+ caddrs.BuyParticipationTokens.String() + "'," +
				"redeem_stake = '"		+ caddrs.RedeemStake.String() + "'," +
				"warp_sync = '"			+ caddrs.WarpSync.String() + "'," +
				"hot_loading = '"		+ caddrs.HotLoading.String() + "'," +
				"affiliates = '"		+ caddrs.Affiliates.String() + "'," +
				"affiliate_val ='"		+ caddrs.AffiliateValidator.String() + "'," +
				"ctime = '"				+ caddrs.Time.String() + "'," +
				"cancel_order = '"		+ caddrs.CancelOrder.String() + "'," +
				"orders = '"			+ caddrs.Orders.String() + "'," +
				"sim_trade = '"			+ caddrs.SimulateTrade.String() + "'," +
				"trade = '"				+ caddrs.Trade.String() + "'," +
				"oi_cash = '"			+ caddrs.OICash.String() + "'," +
				"uniswap_v2_fact ='"	+ caddrs.UniswapV2Factory.String() + "'," +
				"uniswap_v2_r2 ='"		+ caddrs.UniswapV2Router02.String() + "'," +
				"audit_funds = '"		+ caddrs.AuditFunds.String() + "'," +
				"weth9 = '"				+ caddrs.WETH9.String() + "'," +
				"usdc = '"				+ caddrs.USDC.String() + "'," +
				"usdt = '"				+ caddrs.USDT.String() + "'," +
				"relay_hub_v2 = '"		+ caddrs.RelayHubV2.String() + "'," +
				"account_loader = '"	+ caddrs.AccountLoader.String() + "'"

	_,err := ss.db.Exec(query)
	if err!= nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't update ContractAddresses: %v; q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Set_augur_flag(address *common.Address,agtx *p.AugurTx,flag_name string) {

	aid := ss.Lookup_or_create_address(address.String(),agtx.BlockNum,agtx.TxId)

	infinite := false
	var query string
  again:
	query = "UPDATE augur_flag SET "+flag_name+"=TRUE WHERE aid=$1"
	result,err := ss.db.Exec(query,aid)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
		os.Exit(1)
	}
	rows_affected,err:=result.RowsAffected()
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	if rows_affected == 0 {
		query = "INSERT INTO augur_flag(aid) VALUES($1)"
		_,err := ss.db.Exec(query,aid)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
			os.Exit(1)
		}
		if !infinite {
			infinite = true
			goto again
		} else {
			ss.Log_msg(fmt.Sprintf("Infinite loop at Set_augur_flag"))
			os.Exit(1)
		}
	}
	query = "SELECT " +
				"ap_0xtrade_on_cash,ap_fill_on_cash,ap_fill_on_shtok " +
			"FROM augur_flag " +
			"WHERE aid = $1"

	var ap_0xtrade_on_cash,ap_fill_on_cash,ap_fill_on_shtok bool
	row := ss.db.QueryRow(query,aid)
	err=row.Scan(
		&ap_0xtrade_on_cash,
		&ap_fill_on_cash,
		&ap_fill_on_shtok,
	)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Error in Set_augur_flag(): %v, q=%v",err,query))
		os.Exit(1)
	}
	if  ap_0xtrade_on_cash && ap_fill_on_cash && ap_fill_on_shtok  {
		query = "UPDATE augur_flag SET act_block_num=$2 WHERE aid=$1"
		_,err := ss.db.Exec(query,aid,agtx.BlockNum)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v; q=%v",err,query))
			os.Exit(1)
		}
		ss.Link_eoa_and_wallet_contract(aid,aid)
	}
}
func (ss *SQLStorage) Is_augur_activated(address string) bool {

	aid,err :=	ss.Nonfatal_lookup_address_id(address)
	if err!=nil {
		return false
	}
	var query string
	query = "SELECT act_block_num FROM augur_flag WHERE aid = $1"
	var null_block sql.NullInt64
	row := ss.db.QueryRow(query,aid)
	err=row.Scan(&null_block)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
		os.Exit(1)
	}
	if null_block.Valid {
		return true
	}
	return false
}
func (ss *SQLStorage) Delete_abi_artifacts() {

	var query string
	query = "DELETE FROM abi_funcs"
	_,err:=ss.db.Exec(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Delete_abi_artifacts() failed at function deletion: %v \n",err))
		os.Exit(1)
	}
	query = "DELETE FROM abi_events"
	_,err=ss.db.Exec(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Delete_abi_artifacts() failed at event deletion: %v \n",err))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Insert_function_signature(signature,name,contract string) {

	var query string
	query = "INSERT INTO abi_funcs(signature,func_name,contracts) VALUES($1,$2,$3)"
	_,err:=ss.db.Exec(query,signature,name,contract)
	if (err!=nil) {
		if strings.Contains(err.Error(),"duplicate key value") {
			query = "UPDATE abi_funcs SET contracts=concat(contracts,',',$2::text) WHERE signature=$1"
			_,err:=ss.db.Exec(query,signature,contract)
			if (err!=nil) {
				ss.Log_msg(fmt.Sprintf(
					"Insert_function_signature(): failed to update contract name %v %v\n",contract,err,
				))
				os.Exit(1)
			}
		} else {
			ss.Log_msg(
				fmt.Sprintf(
					"Insert_function_signature() function %v INSERT failed, sig %v ; Error : %v \n",
					name,signature,err,
				),
			)
			os.Exit(1)
		}
	}

}
func (ss *SQLStorage) Insert_event_signature(signature,name,contract string) {

	var query string
	query = "INSERT INTO abi_events(signature,evt_name,contracts) VALUES($1,$2,$3)"
	_,err:=ss.db.Exec(query,signature,name,contract)
	if (err!=nil) {
		if strings.Contains(err.Error(),"duplicate key value") {
			query = "UPDATE abi_events SET contracts=concat(contracts,',',$2::text) WHERE signature=$1"
			_,err:=ss.db.Exec(query,signature,contract)
			if (err!=nil) {
				ss.Log_msg(fmt.Sprintf(
					"Insert_function_signature(): failed to update contract name %v : %v\n",contract,err,
				))
				os.Exit(1)
			}
		} else {
			ss.Log_msg(fmt.Sprintf("Insert_event_signature() failed at event insertion: %v \n",err))
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Get_augur_tx_from_db(tx_id int64) *p.AugurTx {

	output := new(p.AugurTx)

	var query string
	query = "SELECT block_num,fa.addr,ta.addr,ctrct_create,value*1e+18::text AS value,tx_hash,td.data " +
				"FROM transaction as tx " +
				"LEFT JOIN address AS fa ON tx.from_aid=fa.address_id " +
				"LEFT JOIN address AS ta ON tx.to_aid=ta.address_id " +
				"LEFT JOIN tx_input AS td on td.tx_id=tx.id " +
			"WHERE tx.id=$1"

	var null_tx_data sql.NullString
	res := ss.db.QueryRow(query,tx_id)
	err := res.Scan(&output.BlockNum,&output.From,&output.To,&output.CtrctCreate,&output.Value,&output.TxHash,&null_tx_data)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Get_augur_tx_from_Db() failed : %v \n",err))
			os.Exit(1)
		}
	}
	if null_tx_data.Valid {
		data,err:=hex.DecodeString(null_tx_data.String)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("Get_augur_tx_from_Db() error decoding tx_data: %v \n",err))
			os.Exit(1)
		}

		output.Input = data
	}
	return output
}
