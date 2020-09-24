package main

import (
	"time"
	"bytes"
	"encoding/hex"
	"math/big"
	"context"
	"os"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/0xProject/0x-mesh/zeroex"

	ztypes "github.com/0xProject/0x-mesh/common/types"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
type ExecWalletTxInputStruct struct {
	To common.Address `abi:"_to"`
	Data []byte `abi:"_data"`
	Value *big.Int `abi:"_value"`
	Payment *big.Int `abi:"_payment"`
	ReferralAddress common.Address `abi:"_referralAddress"`
	Fingerprint [32]byte `abi:"_fingerprint"`
	DesiredSignerBalance *big.Int `abi:"_desiredSignerBalance"`
	MaxExchangeRateInDai *big.Int `abi:"_maxExchangeRateInDai"`
	RevertOnFailure bool `abi:"_revertOnFailure"`
}
type TradeInputStruct struct {
	RequestedFillAmount		*big.Int `abi:"_requestedFillAmount"`
	Fingerprint				[32]byte `abi:"_fingerprint"`
	TradeGroupId			[32]byte `abi:"_tradeGroupId"`
	MaxProtocolFeeDai		*big.Int `abi:"_maxProtocolFeeDai"`
	MaxTrades				*big.Int `abi:"_maxTrades"`
	Orders					[]IExchangeOrder `abi:"_orders"`
	Signatures				[][]byte `abi:"_signatures"`
}
type CancelPrdersInputStruct struct {
	Orders					[]IExchangeOrder `abi:"_orders"`
	Signatures				[][]byte `abi:"_signatures"`
	MaxProtocolFeeDai		*big.Int `abi:"_maxProtocolFeeDai"`

}
func augur_init(addresses *ContractAddresses,contracts *map[string]interface{}) {

	all_contracts = Load_all_artifacts("./abis/augur-artifacts-abi.json")

	// Augur service involves 39 contracts in total. We only use a few of them
	augur_abi = Abi_from_artifacts(contracts,"Augur")
	trading_abi = Abi_from_artifacts(contracts,"AugurTrading")
	zerox_trade_abi = Abi_from_artifacts(contracts,"ZeroXTrade")
	cash_abi = Abi_from_artifacts(contracts,"Cash")
	exchange_abi = Abi_from_artifacts(contracts,"Exchange")
	wallet_abi = Abi_from_artifacts(contracts,"AugurWalletRegistry")

	build_list_of_inspected_events()

	var err error
	ctrct_wallet_registry,err = NewAugurWalletRegistry(addresses.WalletReg, eclient)
	if err != nil {
		Fatalf("Failed to instantiate a AugurWalletRegistry contract: %v", err)
	}
	ctrct_zerox_trade, err = NewZeroX(addresses.ZeroxTrade,eclient)
	if err != nil {
		Fatalf("Failed to instantiate a ZeroX contract: %v", err)
	}
	ctrct_dai_token,err = NewDAICash(addresses.Dai,eclient)
	if err != nil {
		Fatalf("Couldn't initialize DAI Cash contract: %v\n",err)
	}
	ctrct_pl,err = NewProfitLoss(addresses.PL,eclient)
	if err != nil {
		Fatalf("Couldn't initialize Profit Loss contract: %v\n",err)
	}
}
func build_list_of_inspected_events() {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events= make([][]byte,0,64)
	inspected_events = append(inspected_events,
							evt_market_created,
							evt_market_oi_changed,
							evt_market_order,
							evt_market_finalized,
							evt_initial_report_submitted,
							evt_market_volume_changed_v1,
							evt_market_volume_changed_v2,
							evt_dispute_crowd_contrib,
							evt_tokens_transferred,
							evt_token_balance_changed,
							evt_share_token_balance_changed,
							evt_cancel_0x_order,
							evt_transfer_batch,
							evt_transfer_single,
							evt_profit_loss_changed,
							evt_erc20_transfer,
							evt_exchange_fill,
							evt_trading_proceeds_claimed,
							evt_erc1155_approval_for_all,
							evt_erc20_approval,
	)
}
func get_eoa_aid(wallet_addr *common.Address,block_num int64,tx_id int64) int64 {

	var eoa_aid int64 = 0
	wallet_aid,err := storage.Nonfatal_lookup_address_id(wallet_addr.String())
	if err == nil {
		eoa_aid,err = storage.Lookup_eoa_aid(wallet_aid)
		if err == nil {
			storage.Link_eoa_and_wallet_contract(eoa_aid,wallet_aid)
			return eoa_aid
		}
	} else {
		wallet_aid = storage.Lookup_or_create_address(wallet_addr.String(),block_num,tx_id)
	}
	num:=big.NewInt(int64(owner_fld_offset)) 
	key:=common.BigToHash(num)
	Info.Printf("get_eoa_aid: Looking up eoa addr via RPC: wallet addr = %v\n",wallet_addr.String())
	eoa,err := eclient.StorageAt(context.Background(),*wallet_addr,key,nil)
	Info.Printf("get_eoa_aid: output of rpc: %v\n",hex.EncodeToString(eoa))
	var eoa_addr_str string
	if err == nil {
		Info.Printf("get_eoa_aid: got address from RPC successfully")
		eth_addr := common.BytesToAddress(eoa[12:])
		if !Eth_addr_is_zero(&eth_addr) {
			eoa_addr_str = eth_addr.String()
			Info.Printf("get_eoa_aid: Eth addr = %v\n",eoa_addr_str)
			eoa_aid = storage.Lookup_or_create_address(eoa_addr_str,block_num,tx_id)
			Info.Printf("eoa_aid for %v = %v\n",eoa_addr_str,eoa_aid)
			Info.Printf("wallet_aid=%v\n",wallet_aid)
		} else {
			// EOA addr is zero, this means (probably) Wallet Addr is EOA account addr
			Info.Printf("The wallet addr is not contract but EOA account, setting eoa_aid=wallet_aid\n")
			eoa_aid = wallet_aid
		}
		storage.Link_eoa_and_wallet_contract(eoa_aid,wallet_aid)
		Info.Printf("get_eoa_aid: eoa_addr_str=%v\n",eoa_addr_str)
	} else {
		Info.Printf("get_eoa_aid: error at rpc call: %v. Aborting & exiting. No recovery planned\n",err)
		os.Exit(1)// it is easier to relaunch ETL process than designing RPC failure recovery process
	}
	Info.Printf(
		"get_eoa_aid: Success. Getting eoa_aid for address %v, eoa_aid = %v, wallet_aid=%v\n",
		wallet_addr.String(),eoa_aid,wallet_aid,
	)
	return eoa_aid
}
func proc_approval(log *types.Log,agtx_ptr **AugurTx) {

	if len(log.Topics)!=3 {
		Info.Printf("ERC20_Approval event is not compliant log.Topics!=3. Tx hash=%v\n",log.TxHash.String())
		return
	}
	var mevt EApproval
	mevt.Owner= common.BytesToAddress(log.Topics[1][12:])
	mevt.Spender = common.BytesToAddress(log.Topics[2][12:])
	err := cash_abi.Unpack(&mevt,"Approval",log.Data)
	if err != nil {
		Fatalf("Event ERC20_Approval Cash decode error: %v",err)
	} else {
		Info.Printf("ERC20_Approval event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump(Info)
		if bytes.Equal(log.Address.Bytes(),caddrs.Dai.Bytes()) {
			if bytes.Equal(mevt.Spender.Bytes(),caddrs.ZeroxTrade.Bytes()) {
				tx_insert_if_needed(*agtx_ptr)
				storage.Set_augur_flag(&mevt.Owner,*agtx_ptr,"ap_0xtrade_on_cash")
			}
		}
		if bytes.Equal(log.Address.Bytes(),caddrs.Dai.Bytes()) {
			if bytes.Equal(mevt.Spender.Bytes(),caddrs.FillOrder.Bytes()) {
				tx_insert_if_needed(*agtx_ptr)
				storage.Set_augur_flag(&mevt.Owner,*agtx_ptr,"ap_fill_on_cash")
			}
		}
	}
}
func proc_approval_for_all(log *types.Log,agtx_ptr **AugurTx) {

	if len(log.Topics)!=3 {
		Info.Printf("ERC20_ApprovalForAll event is not compliant log.Topics!=3. Tx hash=%v\n",log.TxHash.String())
		return
	}
	var mevt EApprovalForAll
	mevt.Owner= common.BytesToAddress(log.Topics[1][12:])
	mevt.Operator= common.BytesToAddress(log.Topics[2][12:])
	err := zerox_trade_abi.Unpack(&mevt,"ApprovalForAll",log.Data)
	if err != nil {
		Fatalf("Event ApprovalForAll decode error: %v",err)
	} else {
		Info.Printf("ApprovalForAll event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump(Info)
		if bytes.Equal(log.Address.Bytes(),caddrs.ShareToken.Bytes()) {
			if bytes.Equal(mevt.Operator.Bytes(),caddrs.FillOrder.Bytes()) {
				tx_insert_if_needed(*agtx_ptr)
				storage.Set_augur_flag(&mevt.Owner,*agtx_ptr,"ap_fill_on_shtok")
			}
		}
	}
}
func proc_trading_proceeds_claimed(agtx *AugurTx,timestamp int64,log *types.Log) {

	var mevt ETradingProceedsClaimed
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
	mevt.Sender = common.BytesToAddress(log.Topics[2][12:])
	err := augur_abi.Unpack(&mevt,"TradingProceedsClaimed",log.Data)
	if err != nil {
		Fatalf("EventTradingProceedsClaimed error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		return
	}

	Info.Printf("TradingProceedsClaimed event found (block=%v) :\n",log.BlockNumber)
	mevt.Dump(Info)
	pchg := new(PosChg)
	pchg.Mkt_addr = mevt.Market
	pchg.Wallet_addr = mevt.Sender
	pchg.Outcome = new(big.Int)
	pchg.Outcome.Set(mevt.Outcome)
	position_changes = append(position_changes,pchg)
	storage.Update_claim_status(agtx,&mevt,timestamp)
}
func proc_fill_evt(log *types.Log) {
	var mevt EFill
	mevt.MakerAddress= common.BytesToAddress(log.Topics[1][12:])
	mevt.FeeRecipientAddress= common.BytesToAddress(log.Topics[2][12:])
	mevt.OrderHash= log.Topics[3]
	err := exchange_abi.Unpack(&mevt,"Fill",log.Data)
	if err != nil {
		Fatalf("Event Fill for 0x decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.ZeroxXchg.Bytes()) {
		return
	}
	Info.Printf("Fill event found (block=%v) :\n",log.BlockNumber)
	mevt.Dump(Info)
}
func proc_erc20_transfer(log *types.Log,agtx *AugurTx) {
	var mevt ETransfer
	if len(log.Topics)!=3 {
		Info.Printf("ERC20 transfer event is not compliant log.Topics!=3. Tx hash=%v\n",log.TxHash.String())
		return
	}
	mevt.From= common.BytesToAddress(log.Topics[1][12:])
	mevt.To= common.BytesToAddress(log.Topics[2][12:])
	err := cash_abi.Unpack(&mevt,"Transfer",log.Data)
	if err != nil {
		Error.Printf("signature=%v\n",log.Topics[0].String())
		Error.Printf("address=%v\n",log.Address.String())
		Error.Printf("tx hash = %v\n",log.TxHash.String())
		Error.Printf("log.Data=%v, data len=%v\n",hex.EncodeToString(log.Data[:]),len(log.Data[:]))
		Error.Printf("Event ERC20_Transfer, decode error: %v",err)
	} else {
		Info.Printf("ERC20_Transfer event, contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump(Info)
		if bytes.Equal(caddrs.Dai.Bytes(), log.Address.Bytes()) {	// this is DAI contract
			storage.Process_DAI_token_transfer(&mevt,caddrs,agtx)
		}
		if bytes.Equal(caddrs.Reputation.Bytes(), log.Address.Bytes()) {	// this is DAI contract
			storage.Process_REP_token_transfer(&mevt,agtx)
		}
	}
}
func proc_profit_loss_changed(agtx *AugurTx,log *types.Log) int64  {
	var id int64 = 0
	var mevt EProfitLossChanged
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Market = common.BytesToAddress(log.Topics[2][12:])
	mevt.Account= common.BytesToAddress(log.Topics[3][12:])
	err := trading_abi.Unpack(&mevt,"ProfitLossChanged",log.Data)
	if err != nil {
		Fatalf("Event ProfitLossChanged decode error: %v",err)
		return 0
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.AugurTrading.Bytes()) {
		return 0
	}
	Info.Printf("ProfitLossChanged event found (block=%v) :\n",log.BlockNumber)
	mevt.Dump(Info)
	pchg:=new(PosChg)
	pchg.Mkt_addr=mevt.Market
	pchg.Wallet_addr=mevt.Account
	pchg.Outcome = new(big.Int)
	pchg.Outcome.Set(mevt.Outcome)
	position_changes = append(position_changes,pchg)
	eoa_aid := get_eoa_aid(&mevt.Account,agtx.BlockNum,agtx.TxId)
	id = storage.Insert_profit_loss_evt(agtx,eoa_aid,&mevt)
	return id
}
func proc_transfer_single(log *types.Log) {
	var mevt ETransferSingle
	mevt.Operator= common.BytesToAddress(log.Topics[1][12:])
	mevt.From= common.BytesToAddress(log.Topics[2][12:])
	mevt.To= common.BytesToAddress(log.Topics[3][12:])
	err := zerox_trade_abi.Unpack(&mevt,"TransferSingle",log.Data)
	if err != nil {
		Fatalf("Event TransferSingle decode error: %v",err)
	} else {
		Info.Printf("TransferSingle event found (block=%v) :\n",log.BlockNumber)
		mevt.Dump(Info)
	}
}
func proc_transfer_batch(log *types.Log) {
	var mevt ETransferBatch
	mevt.Operator= common.BytesToAddress(log.Topics[1][12:])
	mevt.From= common.BytesToAddress(log.Topics[2][12:])
	mevt.To= common.BytesToAddress(log.Topics[3][12:])
	err := zerox_trade_abi.Unpack(&mevt,"TransferBatch",log.Data)
	if err != nil {
		Fatalf("Event TransferBatch decode error: %v",err)
	} else {
		Info.Printf("TransferBatch event found (block=%v) :\n",log.BlockNumber)
		mevt.Dump(ctrct_zerox_trade,Info)
	}
}
func proc_tokens_transferred(agtx *AugurTx, log *types.Log) {
	var mevt ETokensTransferred
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.From= common.BytesToAddress(log.Topics[2][12:])	// extract From
	mevt.To= common.BytesToAddress(log.Topics[3][12:])	// extract To
	err := augur_abi.Unpack(&mevt,"TokensTransferred",log.Data)
	if err != nil {
		Fatalf("Event TokensTransferred decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		mevt.Dump(Info)
		Info.Printf(
			"TokensTransferred event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	Info.Printf("TokensTransferred event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	mevt.Dump(Info)
	storage.Insert_token_transf_evt(&mevt,agtx)
}
func proc_token_balance_changed(agtx *AugurTx,log *types.Log) {
	var mevt ETokenBalanceChanged
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Owner= common.BytesToAddress(log.Topics[2][12:])
	err := augur_abi.Unpack(&mevt,"TokenBalanceChanged",log.Data)
	if err != nil {
		Fatalf("Event TokenBalanceChanged decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		mevt.Dump(Info)
		Info.Printf(
			"TokenBalanceChanged event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	Info.Printf("TokenBalanceChanged event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	mevt.Dump(Info)
	storage.Insert_token_balance_changed_evt(&mevt,agtx.BlockNum,agtx.TxId)
}
func proc_share_token_balance_changed(agtx *AugurTx,log *types.Log) {
	var mevt EShareTokenBalanceChanged
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Account= common.BytesToAddress(log.Topics[2][12:])
	mevt.Market = common.BytesToAddress(log.Topics[3][12:])
	err := augur_abi.Unpack(&mevt,"ShareTokenBalanceChanged",log.Data)
	if err != nil {
		Fatalf("Event ShareTokenBalanceChanged decode error: %v\n",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		mevt.Dump(Info)
		Info.Printf(
			"ShareTokenBalance Changed event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	Info.Printf("ShareTokenBalanceChanged event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
	mevt.Dump(Info)
	storage.Insert_share_balance_changed_evt(agtx,&mevt)
}
func proc_market_order_event(agtx *AugurTx,log *types.Log,timestamp int64) {

	var mevt EOrderEvent
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Market = common.BytesToAddress(log.Topics[2][12:])
	mevt.EventType = log.Topics[3][31];	// EventType (uint8) which we label as OrderAction
	err := trading_abi.Unpack(&mevt,"OrderEvent",log.Data)
	if err != nil {
		Fatalf("Event OrderEvent decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.AugurTrading.Bytes()) {
		Info.Printf(
			"OrderEvent received and ignored (belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	Info.Printf("OrderEvent event for contract %v (block=%v) : \n",
								log.Address.String(),log.BlockNumber)
	mevt.Dump(Info)
	eoa_aid := get_eoa_aid(&mevt.AddressData[0],agtx.BlockNum,agtx.TxId)
	eoa_fill_aid := get_eoa_aid(&mevt.AddressData[1],agtx.BlockNum,agtx.TxId)

	orders,ospecs := extract_orders_from_input(agtx.Input)
	storage.Insert_market_order_evt(agtx,timestamp,eoa_aid,eoa_fill_aid,&mevt,orders,ospecs)
}
func proc_cancel_zerox_order(agtx *AugurTx,log *types.Log,timestamp int64) {
	var mevt ECancelZeroXOrder
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Market = common.BytesToAddress(log.Topics[2][12:])
	mevt.Account = common.BytesToAddress(log.Topics[3][12:]);
	err := trading_abi.Unpack(&mevt,"CancelZeroXOrder",log.Data)
	if err != nil {
		Fatalf("Event CancelZeroXOrder decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.AugurTrading.Bytes()) {
		return
	}
	ohash := common.BytesToHash(mevt.OrderHash[:])
	ohash_str := ohash.String()
	Info.Printf("CancelZeroXOrder event for contract %v (block=%v) : \n",
								log.Address.String(),log.BlockNumber)
	mevt.Dump(Info)
	orders,ospecs := extract_orders_from_input(agtx.Input)
	if len(orders) == 0 {
		Fatalf("Couldn't extract fill amount from Tx input. Aborting.")
	}
	storage.Cancel_open_order(orders,ospecs,ohash_str,timestamp)
}
func proc_market_oi_changed(block *types.Header, agtx *AugurTx, log *types.Log) {
	var mevt EMarketOIChanged
	err := augur_abi.Unpack(&mevt,"MarketOIChanged",log.Data)
	if err != nil {
		Fatalf("Event decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		return
	}
	Info.Printf("MarketOIChanged event found (block=%v) : \n",log.BlockNumber)
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
	mevt.Market =common.BytesToAddress(log.Topics[2][12:])
	mevt.Dump(Info)
	storage.Insert_market_oi_changed_evt(block,agtx,&mevt)
}
func is_warp_sync_event(log *types.Log) bool {

	var mevt EMarketFinalized
	err := augur_abi.Unpack(&mevt,"MarketFinalized",log.Data)
	if err != nil {
		Fatalf("Event MktFinalizedEvt decode error: %v\n",err)
		return false
	}
	if len(mevt.WinningPayoutNumerators) != 3 {
		return false
	}
	big_num := big.NewInt(0)
	big_num.SetString("25908241534181278443886245536264200757048797036780741184539099179687432804533",10)
	if big_num.Cmp(mevt.WinningPayoutNumerators[1]) ==  0 {
		return true
	}
	return false
}
func proc_market_finalized_evt(agtx *AugurTx,timestamp int64,log *types.Log) {
	var mevt EMarketFinalized
	err := augur_abi.Unpack(&mevt,"MarketFinalized",log.Data)
	if err != nil {
		Fatalf("Event MktFinalizedEvt decode error: %v\n",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		return
	}
	Info.Printf("MarketFinalized event found (block=%v) : \n",log.BlockNumber)
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Market = common.BytesToAddress(log.Topics[2][12:])	// extract universe addr
	mevt.Dump(Info)
	storage.Insert_market_finalized_evt(agtx,timestamp,&mevt)
	pchanges:=storage.Get_mkt_participant_outcomes(&mevt.Market)
	Info.Printf("position_changes: Market finalized, added %v address entries for PL scan\n",len(pchanges))
	for i:=0 ; i<len(pchanges) ; i++ {
		position_changes = append(position_changes,pchanges[i])
	}
}
func proc_initial_report_submitted(agtx *AugurTx, log *types.Log) {
	var mevt EInitialReportSubmitted
	err := augur_abi.Unpack(&mevt,"InitialReportSubmitted",log.Data)
	if err != nil {
		Fatalf("Event InitialReportSubmittedEvt decode error: %v\n",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		return
	}
	Info.Printf("InitialReportSubmitted event found (block=%v) : \n",log.BlockNumber)
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
	mevt.Reporter= common.BytesToAddress(log.Topics[2][12:])
	mevt.Market = common.BytesToAddress(log.Topics[3][12:])
	mevt.Dump(Info)
	storage.Insert_initial_report_evt(agtx,&mevt)
}
func proc_dispute_crowdsourcerer_contribution(agtx *AugurTx,log *types.Log) {
	var mevt EDisputeCrowdsourcerContribution
	err := augur_abi.Unpack(&mevt,"DisputeCrowdsourcerContribution",log.Data)
	if err != nil {
		Fatalf("Event DisputeCrowdsourcerContribution decode error: %v\n",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		return
	}
	Info.Printf("DisputeCrowdsourcerContribution event found (block %v) : \n",log.BlockNumber)
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
	mevt.Reporter= common.BytesToAddress(log.Topics[2][12:])
	mevt.Market = common.BytesToAddress(log.Topics[3][12:])
	mevt.Dump(Info)
	storage.Insert_dispute_crowd_contrib(agtx,&mevt)
}
func proc_market_volume_changed_v1(agtx *AugurTx, log *types.Log) {
	var mevt EMarketVolumeChanged_v1
	// Note: the ./abis/augur-artifacts-abi-26jun.json file was altered to add this (old version of) event
	err := trading_abi.Unpack(&mevt,"MarketVolumeChanged_v1",log.Data)
	if err != nil {
		Fatalf("Event MarketVolumeChanged_v1 decode error: %v\n",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.AugurTrading.Bytes()) {
		return
	}
	Info.Printf("MarketVolumeChanged_v1 event found (block=%v): \n",log.BlockNumber)
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
	mevt.Market = common.BytesToAddress(log.Topics[2][12:])
	mevt.Dump(Info)
	storage.Insert_market_volume_changed_evt_v1(agtx,&mevt)
}
func proc_market_volume_changed_v2(agtx *AugurTx, log *types.Log) {
	var mevt EMarketVolumeChanged_v2
	err := trading_abi.Unpack(&mevt,"MarketVolumeChanged",log.Data)
	if err != nil {
		Fatalf("Event MarketVolumeChanged_v2 decode error: %v\n",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.AugurTrading.Bytes()) {
		return
	}
	Info.Printf("MarketVolumeChanged_v2 event found (block=%v): \n",log.BlockNumber)
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
	mevt.Market = common.BytesToAddress(log.Topics[2][12:])
	mevt.Dump(Info)
	storage.Insert_market_volume_changed_evt_v2(agtx,&mevt)
}
func show_market_created_evt(agtx *AugurTx,log *types.Log) {
	// function used for debugging
	var mevt EMarketCreated
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])   // extract universe addr
	mevt.MarketCreator = common.BytesToAddress(log.Topics[2][12:])  // extract crator addr
	err := augur_abi.Unpack(&mevt,"MarketCreated",log.Data)
	if err != nil {
		Fatalf("Event MarketCreated decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		Info.Printf(
			"MarketCreated event received and ignored (belongs to different contract: %v) " +
			"at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	Info.Printf("MarketCreated event found (block=%v)\n",log.BlockNumber)
	mevt.Dump(Info)
}

func proc_market_created(agtx *AugurTx,log *types.Log,validity_bond string) {
	var mevt EMarketCreated
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.MarketCreator = common.BytesToAddress(log.Topics[2][12:])	// extract crator addr
	err := augur_abi.Unpack(&mevt,"MarketCreated",log.Data)
	if err != nil {
		Fatalf("Event MarketCreated decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		Info.Printf(
			"MarketCreated event received and ignored (belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	Info.Printf("MarketCreated event found (block=%v)\n",log.BlockNumber)
	mevt.Dump(Info)
	eoa_aid := get_eoa_aid(&mevt.MarketCreator,agtx.BlockNum,agtx.TxId)
	storage.Insert_market_created_evt(agtx,eoa_aid,validity_bond,&mevt)
}
func proc_transaction_status(agtx *AugurTx, log *types.Log) {
	var evt EExecuteTransactionStatus
	err := wallet_abi.Unpack(&evt,"ExecuteTransactionStatus",log.Data)
	if err != nil {
		Fatalf("Event decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.WalletReg.Bytes()) {
		Info.Printf(
			"ExecuteTransactionStatus event received and ignored (belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	Info.Printf("ExecuteTransactionStatus event found (block=%v) : \n",log.BlockNumber)
	evt.Dump(Info)
// finding EOA-Wallet link is pending and not clear how it is going to be done
//	eoa_aid := get_eoa_aid(&common.Address{},agtx.BlockNum,agtx.TxId)
//	_ = eoa_aid
	storage.Insert_augur_transaction_status(agtx,&evt)
}
func proc_register_contract(agtx *AugurTx,log *types.Log) {
	var evt ERegisterContract
	err := augur_abi.Unpack(&evt,"RegisterContract",log.Data)
	if err != nil {
		Fatalf("Event Register contract decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"RegisterContract event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	Info.Printf("RegisterContract event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_register_contract_event(agtx,&evt)
}
func proc_universe_created(agtx *AugurTx,log *types.Log) {
	var evt EUniverseCreated
	evt.ParentUniverse = common.BytesToAddress(log.Topics[1][12:])
	evt.ChildUniverse= common.BytesToAddress(log.Topics[2][12:])
	err := augur_abi.Unpack(&evt,"UniverseCreated",log.Data)
	if err != nil {
		Fatalf("Event UniverseCreated decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"UniverseCreated event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	Info.Printf("UniverseCreated event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_universe_created_event(agtx,&evt)
}
func tx_insert_if_needed(agtx *AugurTx) {
	if agtx.TxId == 0 {
		agtx.TxId=storage.Insert_transaction(agtx)
	}
}
func process_event(block *types.Header, agtx *AugurTx,logs *[]*types.Log,lidx int) int64 {
	// Return Value: id of the record inserted (if aplicable, or 0)

	log := &(*(*logs)[lidx])	// we are getting full array of logs (some events need adjacent event data)

	var id int64 = 0
	timestamp := int64(block.Time)
	num_topics := len(log.Topics)
	if num_topics > 0 {

		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc20_approval) {
			proc_approval(log,&agtx)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc1155_approval_for_all) {
			proc_approval_for_all(log,&agtx)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_trading_proceeds_claimed) {
			tx_insert_if_needed(agtx)
			proc_trading_proceeds_claimed(agtx,timestamp,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_exchange_fill) {
			proc_fill_evt(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc20_transfer) {
			tx_insert_if_needed(agtx)
			proc_erc20_transfer(log,agtx)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_profit_loss_changed) {
			tx_insert_if_needed(agtx)
			id = proc_profit_loss_changed(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer_single) {
			proc_transfer_single(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer_batch) {
			proc_transfer_batch(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_tokens_transferred) {
			tx_insert_if_needed(agtx)
			proc_tokens_transferred(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_token_balance_changed) {
			tx_insert_if_needed(agtx)
			proc_token_balance_changed(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_share_token_balance_changed) {
			tx_insert_if_needed(agtx)
			proc_share_token_balance_changed(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_order) {
			tx_insert_if_needed(agtx)
			proc_market_order_event(agtx,log,timestamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_cancel_0x_order) {
			proc_cancel_zerox_order(agtx,log,timestamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_oi_changed) {
			tx_insert_if_needed(agtx)
			proc_market_oi_changed(block,agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_finalized) {
			tx_insert_if_needed(agtx)
			proc_market_finalized_evt(agtx,timestamp,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_initial_report_submitted) {
			tx_insert_if_needed(agtx)
			proc_initial_report_submitted(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_dispute_crowd_contrib) {
			tx_insert_if_needed(agtx)
			proc_dispute_crowdsourcerer_contribution(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_volume_changed_v1) {
			tx_insert_if_needed(agtx)
			proc_market_volume_changed_v2(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_volume_changed_v2) {
			tx_insert_if_needed(agtx)
			proc_market_volume_changed_v2(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_created) {
			// we have inverted the events, so the validity bond amount is stored in
			// ERC20 transfer event (it is transfered to the Universe)
			show_market_created_evt(agtx,log)
			var validity_bond string
			var transf_evt ETransfer
			tr_idx := lidx - 1	// the offset to ERC20 event (as they fired by contracts)
			err := cash_abi.Unpack(&transf_evt,"Transfer",(*logs)[tr_idx].Data)
			if err == nil {
				validity_bond = transf_evt.Value.String()
				Info.Printf("extracted validity bond = %v\n",validity_bond)
			}
			tx_insert_if_needed(agtx)
			proc_market_created(agtx,log,validity_bond)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_execute_tx_status) {
			tx_insert_if_needed(agtx)
			proc_transaction_status(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_register_contract) {
			tx_insert_if_needed(agtx)
			proc_register_contract(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_universe_created) {
			tx_insert_if_needed(agtx)
			proc_universe_created(agtx,log)
		}
	}
	for j:=1; j < num_topics ; j++ {
		Info.Printf("\t\t\t\tLog Topic %v , %v \n",j,log.Topics[j].String())
	}
	return id
}
func roll_back_blocks(diverging_block *types.Header) error {
	// Finds the block from which the fork started
	ctx := context.Background()
	block_num:=diverging_block.Number.Int64()
	starting_block_num:=block_num
	for {
		big_block_num := big.NewInt(block_num)
		block, err := eclient.BlockByNumber(ctx,big_block_num)
		if err != nil {
			return err
		}
		if block == nil {
			e:=errors.New(fmt.Sprintf("ETH client api returned NULL block object (bnum=%v)",block_num))
			return e
		}
		block_hash:=block.Hash().String()
		my_block_num,err := storage.Get_block_num_by_hash(block_hash)
		Info.Printf("Chainsplit fix: hash %v, my_block_num=%v err=%v\n",block_hash,my_block_num,err)
		if err == nil {
			if my_block_num == block.Number().Int64() {
				Info.Printf(
					"Chainsplit fix: deleting blocks higher than %v ; good block hash = %v\n",
					my_block_num,block_hash,
				)
				storage.Chainsplit_delete_blocks(my_block_num)
				storage.Set_last_block_num(my_block_num)
				return nil
			}
		} else {
			Info.Printf(
				"Chainsplit fix: block %v donesn't fit, block_hash=%v not found in my DB.\n",
				block_num,block_hash,
			)
		}
		block_num--
		if (starting_block_num - block_num) > MAX_BLOCKS_CHAIN_SPLIT {
			Info.Printf(
				"Chainsplit fix: Chain split is longer than reasonal length, aborting. " +
				"(starting_block_num=%v, cur_block_num=%v",
				starting_block_num,block_num,
			)
			return errors.New("Chain split max size overflow")
		}
	}
	return errors.New("Chainsplit fix: Undefined behaviour")
}
func process_block(bnum int64,update_last_block bool,no_chainsplit_check bool) error {

	block_hash_str,err:=get_block_hash(bnum)
	if err!=nil {
		return err
	}
	big_bnum:=big.NewInt(int64(bnum))
	block_hash,header,transactions,err := get_full_block(bnum)
	if err!=nil {
		Info.Printf("Can't decode Block object received on RPC: %v. Aborting.\n",err)
		return err
	}
	num_transactions := len(transactions)
	Info.Printf("block %v hash = %v, num_tx=%v\n",bnum,block_hash_str,num_transactions)
	if bnum!=header.Number.Int64() {
		Info.Printf("Retrieved block number %v but Block object contains another number (%v)",bnum,header.Number.Int64())
		Error.Printf("Retrieved block number %v but Block object contains another number (%v)",bnum,header.Number.Int64())
		return errors.New("Block object inconsistency")
	}
	storage.Block_delete_with_everything(big_bnum.Int64())
	receipt_calls := make([]*receiptCallResult,num_transactions,num_transactions)
	for i,tx := range transactions {
		hash := common.HexToHash(tx.TxHash)
		go get_receipt_async(i,hash,&receipt_calls)
	}
	err = storage.Insert_block(block_hash_str,header,no_chainsplit_check)
	if err != nil {
		err = roll_back_blocks(header)
		if err == nil {
			return nil
		}
		Error.Printf("Unable to recover from chainsplit: %v. Aborting",err)
		os.Exit(1)
	}
	if num_transactions == 0 {
		if update_last_block {
			storage.Set_last_block_num(bnum)
		}
		return nil
	}
	for tnum,agtx := range transactions {
		// wait for receipt to arrive
		for {
			if receipt_calls[tnum] != nil {
				break	// receipt arrived from the net, stop waiting
			}
			time.Sleep(1 * time.Millisecond)
		}
		if receipt_calls[tnum].err != nil {
			Info.Printf(
				"Failed to get Tx Receipt for %v, block num=%v. Aborting block processing: %v\n",
				agtx.TxHash,bnum,err,
			)
			Error.Printf(
				"Failed to get Tx Receipt for %v, block num=%v. Aborting block processing: %v\n",
				agtx.TxHash,bnum,err,
			)
			return receipt_calls[tnum].err
		}
		rcpt := receipt_calls[tnum].receipt
		Info.Printf("\ttx: %v of %v : %v at blockNum=%v\n",tnum,num_transactions,agtx.TxHash,bnum)
		Info.Printf("\t from=%v\n",agtx.From)
		Info.Printf("\t to=%v for $%v (%v bytes data)\n",
						agtx.To,agtx.Value,len(agtx.Input))
		if rcpt.Status == types.ReceiptStatusFailed {
			Info.Printf("\t Status: Failed. Skipping this transaciton.\n")
			continue	// transaction failed (i.e. Out of Gas, etc)
		}
		dump_tx_input_if_known(agtx.Input)
		if rcpt.BlockNumber.Int64() != bnum {
			Error.Printf(
				"Transaction's receipt doesn't match current block number. (block possibly changed)" +
				" cur_block_num=%v, receipt.block_num=%v\n",
				bnum,rcpt.BlockNumber.Int64(),
			)
			return errors.New("Block changed during processing")
		}
		agtx.TxId = 0
		if agtx.CtrctCreate == true {
			agtx.To = rcpt.ContractAddress.String()
		}
		agtx.GasUsed = int64(rcpt.GasUsed)
		agtx.TxIndex = int32(tnum)

		if len(agtx.Input) >= 4 {
			input_sig := agtx.Input[:4]
			if bytes.Equal(input_sig,sig_set_referrer) {
				sender := common.HexToAddress(agtx.From)
				tx_insert_if_needed(agtx)
				storage.Set_augur_flag(&sender,agtx,"set_referrer")
			}
		}
		sequencer := new(EventSequencer)
		num_logs := len(rcpt.Logs)
		var agtx_type int = AgTxType_Unclassified
		// Step 1: First detect what kind of Augur Transaction we are dealing with
		for i:=0 ; i<num_logs ; i++ {
			if len(rcpt.Logs[i].Topics) > 0 {
				Info.Printf(
					"\t\tlog %v\t for contract %v (%v of %v items)\n",
					hex.EncodeToString(rcpt.Logs[i].Topics[0][0:4]),
					rcpt.Logs[i].Address.String(),(i+1),len(rcpt.Logs))
				if 0 == bytes.Compare(rcpt.Logs[i].Topics[0].Bytes(),evt_market_finalized) {
					if is_warp_sync_event(rcpt.Logs[i]) {
						// WarpSync market emits 2 events MarketFFinalized and MarketCreated
						// MarketFinalized doesn't have ProfitLoss events, so we can process it
						// just using inverse order (i.e. considering it as non-MarketFinalized)
					} else {
						agtx_type = AgTxType_MarketFinalized
					}
				}
				if 0 == bytes.Compare(rcpt.Logs[i].Topics[0].Bytes(),evt_market_order) {
					agtx_type = AgTxType_MarketOrder
				}
			}
			sequencer.append_event(rcpt.Logs[i])
		}
		// Step 1.1 If a Wallet contract has been created, register EOA-Wallet link
		wallet_created,wallet_addr,possible_eoa_addr := was_wallet_created(caddrs,rcpt.Logs)
		if wallet_created {
			tx_insert_if_needed(agtx)
			var from_addr *string = &agtx.From
			var possible_eoa_str string
			if possible_eoa_addr != nil {
				possible_eoa_str = possible_eoa_addr.String()
				from_addr = &possible_eoa_str
			}
			storage.Register_eoa_and_wallet(*from_addr,wallet_addr.String(),agtx.BlockNum,agtx.TxId)
		}
		// Step 1.2 If transaction contains executeWalletTransaction, store it in the DB
		if (agtx.To == caddrs.WalletReg.String()) || (agtx.To == caddrs.WalletReg2.String()) {
			exec_wtx := contains_execute_wallet_transaction_call(agtx.Input)
			if exec_wtx != nil {
				tx_insert_if_needed(agtx)
				// Note: Tests are pending for transactions going through GSN (EOA address must be extracted
				//			from the tx.Data(), in the first 20 bytes
				eoa_aid := storage.Lookup_or_create_address(agtx.From,agtx.BlockNum,agtx.TxId)
				wallet_aid,err := storage.Lookup_wallet_aid(eoa_aid)
				if err != nil {
					Info.Printf(
						"executeWalletTransaction(): wallet_aid=0 for eoa=%v (id=%v)\n",
						agtx.From,eoa_aid,
					)
				}
				exec_wtx.Dump(Info)
				storage.Insert_execute_wallet_tx(eoa_aid,wallet_aid,agtx,exec_wtx)
			}
		}
		// Step 2: Knowing what kind of Augur Transaction, we are sorting events in an order
		//			that is convinient for us to process the event series
		var ordered_list []*types.Log
		switch agtx_type {
			case AgTxType_Unclassified:
				ordered_list = sequencer.get_ordered_event_list()
			case AgTxType_MarketFinalized:
				// logs with Market finalized event need to have special order
				ordered_list = sequencer.get_events_for_market_finalized_case()
			case AgTxType_MarketOrder:
				ordered_list = sequencer.get_events_for_market_order_case()
			default:
				Info.Printf("Undefined behaviour in detecting Augur Transaction type")
				os.Exit(1)
		}
		num_logs = len(ordered_list)
		pl_entries := make([]int64,0,2);// profit loss entries
		// before processing events we need to reset these global vars as they accumulate some data
		market_order_id = 0
		//
		// Step 3: Execute events using ordered list prepared in previous step
		for i:=0 ; i < num_logs ; i++ {
			if len(ordered_list[i].Topics) > 0 {
				Info.Printf(
					"\t\tchecking log with sig %v\t for contract %v\n",
					hex.EncodeToString(ordered_list[i].Topics[0][0:4]),
					ordered_list[i].Address.String())
				id := process_event(header,agtx,&ordered_list,i)
				if 0 == bytes.Compare(ordered_list[i].Topics[0].Bytes(),evt_profit_loss_changed) {
					pl_entries = append(pl_entries,id)
				}
			}
		}
	}
	Info.Printf("block_proc: %v %v ; %v transactions\n",bnum,block_hash.String(),num_transactions)
	if update_last_block {
		storage.Set_last_block_num(bnum)
	}
	return nil
}
func get_order_data(o *IExchangeOrder) (zeroex.Order,common.Hash) {

	var zero_order zeroex.Order
	zero_order.ChainID=big.NewInt(caddrs.ChainId)
	zero_order.ExchangeAddress.SetBytes(caddrs.ZeroxXchg.Bytes())
	zero_order.MakerAddress.SetBytes(o.MakerAddress.Bytes())
	zero_order.MakerAssetData = make([]byte,len(o.MakerAssetData))
	copy(zero_order.MakerAssetData,o.MakerAssetData)
	zero_order.MakerFeeAssetData = make([]byte,len(o.MakerFeeAssetData))
	copy(zero_order.MakerFeeAssetData,o.MakerFeeAssetData)
	zero_order.MakerAssetAmount = new(big.Int)
	zero_order.MakerAssetAmount.Set(o.MakerAssetAmount)
	zero_order.MakerFee = new(big.Int)
	zero_order.MakerFee.Set(o.MakerFee)
	zero_order.TakerAddress.SetBytes(o.TakerAddress.Bytes())
	zero_order.TakerAssetData = make([]byte,len(o.TakerAssetData))
	copy(zero_order.TakerAssetData,o.TakerAssetData)
	zero_order.TakerFeeAssetData = make([]byte,len(o.TakerFeeAssetData))
	copy(zero_order.TakerFeeAssetData,o.TakerFeeAssetData)
	zero_order.TakerAssetAmount = new(big.Int)
	zero_order.TakerAssetAmount.Set(o.TakerAssetAmount)
	zero_order.TakerFee = new(big.Int)
	zero_order.TakerFee.Set(o.TakerFee)
	zero_order.SenderAddress.SetBytes(o.SenderAddress.Bytes())
	zero_order.FeeRecipientAddress.SetBytes(o.FeeRecipientAddress.Bytes())
	zero_order.ExpirationTimeSeconds = new(big.Int)
	zero_order.ExpirationTimeSeconds.Set(o.ExpirationTimeSeconds)
	zero_order.Salt = new(big.Int)
	zero_order.Salt.Set(o.Salt)
	hash,err:=zero_order.ComputeOrderHash()
	if err!=nil {
		Fatalf("can't compute ZeroX order hash: %v\n",err)
	}
	Info.Printf("get_order_hash() returning %v\n",hash.String())
	return zero_order,hash
}
func decode_original_fill_amount(input_data []byte,method_sig []byte) map[string]*big.Int {
	output := make(map[string]*big.Int,0)
	var input_data_decoded TradeInputStruct
	method, err := zerox_trade_abi.MethodById(method_sig)
	if err != nil {
		Fatalf("Method not found")
	}
	err = method.Inputs.Unpack(&input_data_decoded, input_data)
	if err != nil {
		Fatalf("Couldn't decode input of tx: %v",err)
	}
	if len(input_data_decoded.Orders) > 0 {
		Info.Printf("Requested fill amount = %v\n",input_data_decoded.RequestedFillAmount.String())
		Info.Printf("num orders=%v\n",len(input_data_decoded.Orders))
		for i,order := range input_data_decoded.Orders {
			_,h := get_order_data(&order)
			hash_str := h.String()
			Info.Printf(
				"Order %v (%v), maker amount = %v, taker amount=%v\n",
				i,hash_str,order.MakerAssetAmount,order.TakerAssetAmount,
			)
			initial_amount := big.NewInt(0)
			initial_amount.Set(order.MakerAssetAmount)
			output[hash_str]=initial_amount
		}
	} else {
		Error.Printf("Undefined behavior: no orders detected on the input of ZeroXTrade::trade()")
		os.Exit(1)
	}

	return output
}
func decode_0x_orders(input_data []byte,method_sig []byte) (map[string]*ztypes.OrderInfo,map[string]*ZxMeshOrderSpec) {

	output1 := make(map[string]*ztypes.OrderInfo,0)
	output2 := make(map[string]*ZxMeshOrderSpec,0)

	var trade_input_data_decoded TradeInputStruct
	var cancel_order_input_data_decoded CancelPrdersInputStruct
	var decoded_orders []IExchangeOrder
	var decoded_signatures [][]byte

	zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
	if 0 == bytes.Compare(method_sig,zeroex_trade_sig) {
		method, err := zerox_trade_abi.MethodById(method_sig)
		if err != nil {
			Fatalf("Method not found")
		}
		err = method.Inputs.Unpack(&trade_input_data_decoded, input_data)
		if err != nil {
			Fatalf("Couldn't decode input of tx: %v",err)
		}
		decoded_orders = trade_input_data_decoded.Orders
		decoded_signatures = trade_input_data_decoded.Signatures
	}
	zeroex_cancel_sig,_ := hex.DecodeString("4ea96c30")
	if 0 == bytes.Compare(method_sig,zeroex_cancel_sig) {
		method, err := zerox_trade_abi.MethodById(method_sig)
		if err != nil {
			Fatalf("Method not found")
		}
		err = method.Inputs.Unpack(&cancel_order_input_data_decoded, input_data)
		if err != nil {
			Fatalf("Couldn't decode input of tx: %v",err)
		}
		decoded_orders=cancel_order_input_data_decoded.Orders
		decoded_signatures=cancel_order_input_data_decoded.Signatures
	}
	if len(decoded_orders) > 0 {
		for i,order := range decoded_orders {
			ord,h := get_order_data(&order)
			hash_str := h.String()
			Info.Printf(
				"Order %v (%v), maker amount = %v, taker amount=%v\n",
				i,hash_str,order.MakerAssetAmount,order.TakerAssetAmount,
			)
			order_info := new(ztypes.OrderInfo)
			order_info.OrderHash.SetBytes(h.Bytes())
			order_info.SignedOrder=new(zeroex.SignedOrder)
			order_info.SignedOrder.Signature=make([]byte,len(decoded_signatures[i]))
			order_info.SignedOrder.Order = ord
			order_info.FillableTakerAssetAmount = big.NewInt(0) // this value is incorrect, but we don't have the correct one
			copy(order_info.SignedOrder.Signature,decoded_signatures[i])
			output1[hash_str]=order_info
			ospec := get_ospec(ord.MakerAssetData,&hash_str)
			output2[hash_str] = ospec
		}
	} else {
		Error.Printf("Undefined behavior: no orders detected on the input of ZeroXTrade::trade()")
		os.Exit(1)
	}

	return output1,output2
}
func dump_tx_input_if_known(tx_data []byte) {

	if len(tx_data) < 32 {
		Info.Printf("dump_tx_input: input sig: %v\n",hex.EncodeToString(tx_data[:]))
		return
	}
	input_sig := tx_data[:4]
	set_timestamp_sig ,_ := hex.DecodeString("a0a2b573")
	if 0 == bytes.Compare(input_sig,set_timestamp_sig) {
		Info.Printf("augur_wallet_call: Skipping setTimestamp() transaction\n")
		return
	}
	decoded_sig ,_ := hex.DecodeString("78dc0eed")
	if 0 == bytes.Compare(input_sig,decoded_sig) {
		input_data_raw:= tx_data[4:]
		var input_data ExecWalletTxInputStruct
		method, err := wallet_abi.MethodById(decoded_sig)
		if err != nil {
			Fatalf("Method not found")
		}
		err = method.Inputs.Unpack(&input_data, input_data_raw)
		if err != nil {
			Fatalf("Couldn't decode input of tx %v",err)
		}
		Info.Printf("ExecuteWalletTransaction {\n")
		Info.Printf("\tto: %v\n",input_data.To.String())
		Info.Printf("\tdata: %v\n",hex.EncodeToString(input_data.Data[:]))
		Info.Printf("\tvalue: %v\n",input_data.Value.String())
		Info.Printf("\tpayment: %v\n",input_data.Payment.String())
		Info.Printf("\treferralAddress:  %v\n",input_data.ReferralAddress.String())
		Info.Printf("\tfingerprint: %v\n",hex.EncodeToString(input_data.Fingerprint[:]))
		Info.Printf("\tdesiredSignerBalance: %v\n",input_data.DesiredSignerBalance.String())
		Info.Printf("\tmaxExchangeRateInDai: %v\n",input_data.MaxExchangeRateInDai.String())
		Info.Printf("\trevertOnFaliure: %v\n",input_data.RevertOnFailure)
		Info.Printf("}\n")

		// check for internal transactions for the Wallet Registry contract
		if len(input_data.Data) >= 4 {
			input_sig := input_data.Data[:4]
			market_proceeds_sig ,_ := hex.DecodeString("db754422")
			if 0 == bytes.Compare(input_sig,market_proceeds_sig) {
				Info.Printf("augur_wallet_call: claimMarketProceeds()\n")
				return
			}
			trading_proceeds_sig ,_ := hex.DecodeString("efd342c1")
			if 0 == bytes.Compare(input_sig,trading_proceeds_sig) {
				Info.Printf("augur_wallet_call: claimTradingProceeds()\n")
				return
			}
			zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				Info.Printf("augur_wallet_call: ZeroEx::trade()\n")
				amounts := decode_original_fill_amount(input_data.Data[4:],zeroex_trade_sig)
				for h,a := range amounts {
					if a == nil {
						Fatalf("amounts map contains null initial_order bigint")
					}
					Info.Printf("o %v, amount = %v\n",h,a.String())
				}
				return
			}
		}
	} else {
		Info.Printf("dump_tx_input: input sig: %v\n",hex.EncodeToString(input_sig[:]))
		if len(input_sig) >= 4 {
			input_data_raw:= tx_data[4:]
			Info.Printf("tx input= %v\n",hex.EncodeToString(input_data_raw))
			zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				Info.Printf("direct call to ZeroEx::trade()\n")
				amounts := decode_original_fill_amount(input_data_raw,zeroex_trade_sig)
				for h,a := range amounts {
					if a == nil {
						Fatalf("amounts map contains null initial_order bigint")
					}
					Info.Printf("o %v, amount = %v\n",h,a.String())
				}
				return
			}
		}
	}
}
func scan_profit_loss_data_for_debugging(block_num int64,position_changes *[]*PosChg) {
	// this function makes direct calls to contracts to get changes in profit loss and record them
	// right after each block is processed (developed for debugging purposes)

	var copts = new(bind.CallOpts)
	for i:=0 ; i<len(*position_changes) ; i++ {
		pchg := (*position_changes)[i]
		Info.Printf("profit_loss debug: processing pl for %v\n",pchg.Wallet_addr.String())
		pl,err := ctrct_pl.GetRealizedProfit(copts,pchg.Mkt_addr,pchg.Wallet_addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetRealizedProfit for addr %v outcome %v: %v\n",
							pchg.Wallet_addr.String(),pchg.Outcome.String(),err)
			continue
		} else {
			pchg.ProfitLoss = pl
		}
		ff,err := ctrct_pl.GetFrozenFunds(copts,pchg.Mkt_addr,pchg.Wallet_addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetFrozenFunds for addr %v outcome %v: %v\n",
							pchg.Wallet_addr.String(),pchg.Outcome.String(),err)
			continue
		} else {
			pchg.FrozenFunds = ff
		}
		psize,err := ctrct_pl.GetNetPosition(copts,pchg.Mkt_addr,pchg.Wallet_addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetNetPosition for addr %v outcome %v: %v\n",
							pchg.Wallet_addr.String(),pchg.Outcome.String(),err)
			continue
		} else {
			pchg.NetPos = psize
		}
		price,err := ctrct_pl.GetAvgPrice(copts,pchg.Mkt_addr,pchg.Wallet_addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetAvgPricefor addr %v outcome %v: %v\n",
							pchg.Wallet_addr.String(),pchg.Outcome.String(),err)
			continue
		} else {
			pchg.AvgPrice = price
		}
		pchg.BlockNum=int64(block_num)
		Info.Printf("inserting pchg %+v\n",*pchg)
		storage.Insert_profit_loss_debug_rec(pchg)
	}
}
func was_wallet_created(caddrs *ContractAddresses,event_list []*types.Log) (bool,common.Address,*common.Address){

	var wallet_addr common.Address
	var eoa_addr *common.Address = nil
	// detects pattern of Wallet creation, function AugurWalletFactory::createAugurWallet()
	// scans events and determines if the set of events has correct properties to consider
	// that this is indeed a pattern of Wallet creation.

	var augur_approval bool = false // Pattern 1: Approval event for AugurContract for MAX_APPROVAL_AMOUNT
	var cash_create_approval bool = false // Pattern 2: Approval event for CreateOrder for MAX_APPROVAL_AMOUNT
	var sharetoken_create_approval bool = false // Pattern 3: ApprovalForAll for ShareToken
	var cash_fill_approval bool = false // Pattern 4: Approval event for FillOrder contract
	var sharetoken_fill_approval bool = false // Pattern 3: ApprovalForAll for FillOrder
	var cash_zerox_trade_approval bool = false // Pattern 5: Approval event for ZeroexTrade contract
	var wtx_status bool = false
	var tx_relayed bool = false

	// First we compare contract addresses, and after that for event signature because event signatures
	//		appear more frequently than an event coming for a specific address owner
	for _,log := range event_list {
		if len(log.Topics) < 1 {
			continue
		}
		if bytes.Equal(log.Address.Bytes(),caddrs.WalletReg.Bytes()) {
			if bytes.Equal(log.Topics[0].Bytes(),evt_execute_tx_status) {
				wtx_status = true
				continue
			}
		}
		if bytes.Equal(log.Address.Bytes(),caddrs.WalletReg2.Bytes()) {
			if bytes.Equal(log.Topics[0].Bytes(),evt_execute_tx_status) {
				wtx_status = true
				continue
			}
		}
		if len(log.Topics) < 3 {
			continue		// an event with no topics, not our use case
		}

		addr := common.BytesToAddress(log.Topics[2][12:])
		Info.Printf("checking addr %v, contract=%v\n",addr.String(),log.Address.String())
		if bytes.Equal(log.Topics[0].Bytes(),evt_tx_relayed) {
			eoa_addr = new(common.Address)
			eoa_addr.SetBytes(addr.Bytes())
			tx_relayed = true
			continue
		}
		if bytes.Equal(addr.Bytes(),caddrs.Augur.Bytes()) {
			if bytes.Equal(log.Address.Bytes(),caddrs.Dai.Bytes()) {
				if bytes.Equal(log.Topics[0].Bytes(),evt_erc20_approval) {
					augur_approval = true
					continue
				}
			}
		}
		if bytes.Equal(addr.Bytes(),caddrs.CreateOrder.Bytes()) {
			if bytes.Equal(log.Address.Bytes(),caddrs.Dai.Bytes()) {
				if bytes.Equal(log.Topics[0].Bytes(),evt_erc20_approval) {
					cash_create_approval = true
					continue
				}
			}
			if bytes.Equal(log.Address.Bytes(),caddrs.ShareToken.Bytes()) {
				if bytes.Equal(log.Topics[0].Bytes(),evt_erc1155_approval_for_all) {
					if len(log.Topics)!=3 {
						Info.Printf("ERC20_ApprovalForAll not compliant log.Topics!=3. Aborting for debugging\n")
						Error.Printf("ERC20_ApprovalForAll not compliant log.Topics!=3. Aborting for debugging\n")
						os.Exit(1)
					}
					wallet_addr = common.BytesToAddress(log.Topics[1][12:])
					sharetoken_create_approval = true
					continue
				}
			}
		}
		if bytes.Equal(addr.Bytes(),caddrs.FillOrder.Bytes()) {
			if bytes.Equal(log.Address.Bytes(),caddrs.Dai.Bytes()) {
				if bytes.Equal(log.Topics[0].Bytes(),evt_erc20_approval) {
					cash_fill_approval = true
					continue
				}
			}
			if bytes.Equal(log.Address.Bytes(),caddrs.ShareToken.Bytes()) {
				if bytes.Equal(log.Topics[0].Bytes(),evt_erc1155_approval_for_all) {
					if len(log.Topics)!=3 {
						Info.Printf("ERC20_ApprovalForAll not compliant log.Topics!=3. Aborting for debugging\n")
						Error.Printf("ERC20_ApprovalForAll not compliant log.Topics!=3. Aborting for debugging\n")
						os.Exit(1)
					}
					sharetoken_fill_approval = true
					continue
				}
			}
		}
		if bytes.Equal(addr.Bytes(),caddrs.ZeroxTrade.Bytes()) {
			if bytes.Equal(log.Address.Bytes(),caddrs.Dai.Bytes()) {
				if bytes.Equal(log.Topics[0].Bytes(),evt_erc20_approval) {
					cash_zerox_trade_approval = true
					continue
				}
			}
		}
	}
	var output bool = false
	if	augur_approval &&
		cash_create_approval && sharetoken_create_approval &&
		cash_fill_approval && sharetoken_fill_approval &&
		cash_zerox_trade_approval &&
		wtx_status {
			output = true
	}
	Info.Printf("tx_relayed=%v\n",tx_relayed)
	Info.Printf("augur_approval=%v\n",augur_approval)
	Info.Printf("cash_create_approval=%v\n",cash_create_approval)
	Info.Printf("sharetoken_create_approval=%v\n",sharetoken_create_approval)
	Info.Printf("cash_fill_approval=%v\n",cash_fill_approval)
	Info.Printf("sharetoken_fill_approval=%v\n",sharetoken_fill_approval)
	Info.Printf("cash_zerox_trade_approval=%v\n",cash_zerox_trade_approval)
	Info.Printf("wtx_status=%v\n",wtx_status)
	Info.Printf("wallet_addr=%v\n",wallet_addr.String())
	if eoa_addr != nil {
		Info.Printf("eoa_addr=%v\n",eoa_addr.String())
	}
	Info.Printf("output=%v\n",output)
	return output,wallet_addr,eoa_addr
}
func contains_execute_wallet_transaction_call(tx_data []byte) *ExecuteWalletTx {

	if len(tx_data) < 32 {
		return nil
	}
	input_sig := tx_data[:4]
	if 0 == bytes.Compare(input_sig,exec_wtx_sig) {
		input_data_raw:= tx_data[4:]
		var input_data ExecWalletTxInputStruct
		method, err := wallet_abi.MethodById(exec_wtx_sig)
		if err != nil {
			Fatalf("Method not found")
		}
		err = method.Inputs.Unpack(&input_data, input_data_raw)
		if err != nil {
			Fatalf("Couldn't decode input of tx %v",err)
		}
		exec_wtx:=new(ExecuteWalletTx)
		exec_wtx.To=input_data.To.String()
		exec_wtx.CallData=hex.EncodeToString(input_data.Data[:])
		if len(input_data.Data)>=4 {
			exec_wtx.InputSig=hex.EncodeToString(input_data.Data[:4])
		}
		exec_wtx.Value=input_data.Value.String()
		exec_wtx.Payment=input_data.Payment.String()
		exec_wtx.ReferralAddress=input_data.ReferralAddress.String()
		exec_wtx.Fingerprint=hex.EncodeToString(input_data.Fingerprint[:])
		exec_wtx.DesiredSignerBalance=input_data.DesiredSignerBalance.String()
		exec_wtx.MaxExchangeRateInDAI=input_data.MaxExchangeRateInDai.String()
		exec_wtx.RevertOnFailure=input_data.RevertOnFailure
		return exec_wtx
	}
	return nil
}
/* DISCONTINUED
func extract_original_fill_amount(tx_data []byte) map[string]*big.Int {

	if len(tx_data) < 32 {
		return make(map[string]*big.Int,0)
	}
	input_sig := tx_data[:4]
	decoded_sig ,_ := hex.DecodeString("78dc0eed")
	if 0 == bytes.Compare(input_sig,decoded_sig) {
		input_data_raw:= tx_data[4:]
		var input_data ExecWalletTxInputStruct
		method, err := wallet_abi.MethodById(decoded_sig)
		if err != nil {
			Fatalf("Method not found")
		}
		err = method.Inputs.Unpack(&input_data, input_data_raw)
		if err != nil {
			Fatalf("Couldn't decode input of tx %v",err)
		}

		// check for internal transactions for the Wallet Registry contract
		if len(input_data.Data) >= 4 {
			input_sig := input_data.Data[:4]
			zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				Info.Printf("augur_wallet_call: ZeroEx::trade()\n")
				return decode_original_fill_amount(input_data.Data[4:],zeroex_trade_sig)
			}
		}
	} else {
		if len(input_sig) >= 4 {
			input_data_raw:= tx_data[4:]
			Info.Printf("tx input= %v\n",hex.EncodeToString(input_data_raw))
			zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				return decode_original_fill_amount(input_data_raw,zeroex_trade_sig)
			}
		}
	}
	return make(map[string]*big.Int,0)
}*/
func extract_orders_from_input(tx_data []byte) (map[string]*ztypes.OrderInfo,map[string]*ZxMeshOrderSpec) {
	// returns orders in one map and initial amounts in another map
	if len(tx_data) < 32 {
		return make(map[string]*ztypes.OrderInfo,0),make(map[string]*ZxMeshOrderSpec,0)

	}
	input_sig := tx_data[:4]
	decoded_sig ,_ := hex.DecodeString("78dc0eed")
	if 0 == bytes.Compare(input_sig,decoded_sig) {
		input_data_raw:= tx_data[4:]
		var input_data ExecWalletTxInputStruct
		method, err := wallet_abi.MethodById(decoded_sig)
		if err != nil {
			Fatalf("Method not found")
		}
		err = method.Inputs.Unpack(&input_data, input_data_raw)
		if err != nil {
			Fatalf("Couldn't decode input of tx %v",err)
		}

		// check for internal transactions for the Wallet Registry contract
		if len(input_data.Data) >= 4 {
			input_sig := input_data.Data[:4]
			zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				Info.Printf("augur_wallet_call: ZeroEx::trade()\n")
				return decode_0x_orders(input_data.Data[4:],zeroex_trade_sig)
			}
			zeroex_cancel_sig,_ := hex.DecodeString("4ea96c30")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				Info.Printf("augur_wallet_call: ZeroEx::cancelOrder()\n")
				return decode_0x_orders(input_data.Data[4:],zeroex_cancel_sig)
			}
		}
	} else {
		if len(input_sig) >= 4 {
			input_data_raw:= tx_data[4:]
			Info.Printf("tx input= %v\n",hex.EncodeToString(input_data_raw))
			zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				return decode_0x_orders(input_data_raw,zeroex_trade_sig)
			}
			zeroex_cancel_sig,_ := hex.DecodeString("4ea96c30")
			if 0 == bytes.Compare(input_sig,zeroex_cancel_sig) {
				return decode_0x_orders(input_data_raw,zeroex_cancel_sig)
			}
		}
	}
	return make(map[string]*ztypes.OrderInfo,0),make(map[string]*ZxMeshOrderSpec,0)
}
func get_ospec(maker_asset_data []byte,order_hash *string) *ZxMeshOrderSpec {

	var copts = new(bind.CallOpts)
	adata,err := ctrct_zerox_trade.DecodeAssetData(copts,maker_asset_data)
	if err!=nil {
		Info.Printf("couldn't decode asset data for order %v : %v\n",*order_hash,err)
		Error.Printf("couldn't decode asset data for order %v : %v\n",*order_hash,err)
		os.Exit(1)
	}
	unpacked_id,err := ctrct_zerox_trade.UnpackTokenId(copts,adata.TokenIds[0])
	if err!=nil {
		Info.Printf("Unpack token id failed for order %v: %v\n",*order_hash,err)
		Error.Printf("Unpack token id failed for order %v: %v\n",*order_hash,err)
		os.Exit(1)
	}
	return &unpacked_id
}

