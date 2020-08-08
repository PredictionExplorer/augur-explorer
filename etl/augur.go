package main

import (
	"time"
	"bytes"
	"encoding/hex"
	//"encoding/json"
	"math/big"
	"context"
	"os"
	"errors"
	"fmt"
//	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/common/hexutil"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
type EventSequencer struct {	// determines the order for contained events
	unordered_list		[]*types.Log
}
type InputStruct struct {
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
func (sequencer *EventSequencer) append_event(new_log *types.Log) {

	sequencer.unordered_list = append(sequencer.unordered_list,new_log)
}
func (sequencer *EventSequencer) get_ordered_event_list() []*types.Log {
	// determines the correct event sequence for different event combinations
//	return sequencer.unordered_list
	// at this moment we just reverse the events. more logic will follow later if needed
	output := make([]*types.Log,0,8)
	for i := len(sequencer.unordered_list) - 1; i >= 0; i-- {
		output = append(output,sequencer.unordered_list[i])
	}
	return output
}
func (sequencer *EventSequencer) get_events_for_market_finalized_case() []*types.Log {
	// we must move TradingProceedsClaimed events to the end, so they execut before ProfitLoss events

	output := make([]*types.Log,0,8)
	var market_finalized *types.Log
	profit_loss_events := make([]*types.Log,0,8)
	proceed_events := make([]*types.Log,0,8)
	other_events := make([]*types.Log,0,8)

	for i := 0 ; i < len(sequencer.unordered_list) ; i++ {
		if len(sequencer.unordered_list[i].Topics) == 0 {
			other_events = append(other_events,sequencer.unordered_list[i])
			continue
		}
		if 0 == bytes.Compare(sequencer.unordered_list[i].Topics[0].Bytes(),evt_profit_loss_changed) {
			profit_loss_events = append(profit_loss_events,sequencer.unordered_list[i])
			continue
		}
		if 0 == bytes.Compare(sequencer.unordered_list[i].Topics[0].Bytes(),evt_trading_proceeds_claimed) {
			proceed_events = append(proceed_events,sequencer.unordered_list[i])
			continue
		}
		if 0 == bytes.Compare(sequencer.unordered_list[i].Topics[0].Bytes(),evt_market_finalized) {
			market_finalized = sequencer.unordered_list[i]
			continue
		}
		other_events = append(other_events,sequencer.unordered_list[i])
	}
	output = append(output,market_finalized)
	for i := 0; i < len(profit_loss_events) ; i++ {
		output = append(output,profit_loss_events[i])
	}
	for i := 0; i < len(proceed_events); i++ {
		output = append(output,proceed_events[i])
	}
	for i := 0; i < len(other_events); i++ {
		output = append(output,other_events[i])
	}
	return output
}
func augur_init(addresses *ContractAddresses,contracts *map[string]interface{}) {

	//Init_contract_addresses(addresses)

	all_contracts = Load_all_artifacts("./abis/augur-artifacts-abi.json")
	//dump_all_artifacts()

	// Augur service involves 39 contracts in total. We only use a few of them
	augur_abi = Abi_from_artifacts(contracts,"Augur")
	trading_abi = Abi_from_artifacts(contracts,"AugurTrading")
	zerox_abi = Abi_from_artifacts(contracts,"ZeroXTrade")
	cash_abi = Abi_from_artifacts(contracts,"Cash")
	exchange_abi = Abi_from_artifacts(contracts,"Exchange")
	wallet_abi = Abi_from_artifacts(contracts,"AugurWalletRegistry")

	build_list_of_inspected_events()

	var err error
	ctrct_wallet_registry,err = NewAugurWalletRegistry(addresses.WalletReg, eclient)
	if err != nil {
		Fatalf("Failed to instantiate a AugurWalletRegistry contract: %v", err)
	}
	ctrct_zerox, err = NewZeroX(addresses.Zerox,eclient)
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
							evt_zerox_approval_for_all,
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
func proc_approval(log *types.Log) {

	if len(log.Topics)!=3 {
		Info.Printf("ERC20_Approval event is not compliant log.Topics!=3. Tx hash=%v\n",log.TxHash.String())
		return
	}
	var mevt Approval
	mevt.Owner= common.BytesToAddress(log.Topics[1][12:])
	mevt.Spender = common.BytesToAddress(log.Topics[2][12:])
	err := cash_abi.Unpack(&mevt,"Approval",log.Data)
	if err != nil {
		Fatalf("Event ERC20_Approval Cash decode error: %v",err)
	} else {
		Info.Printf("ERC20_Approval event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump(Info)
	}
}
func proc_approval_for_all(log *types.Log) {

	if len(log.Topics)!=3 {
		Info.Printf("ERC20_ApprovalForAll event is not compliant log.Topics!=3. Tx hash=%v\n",log.TxHash.String())
		return
	}
	var mevt ApprovalForAll
	mevt.Owner= common.BytesToAddress(log.Topics[1][12:])
	mevt.Operator= common.BytesToAddress(log.Topics[2][12:])
	err := zerox_abi.Unpack(&mevt,"ApprovalForAll",log.Data)
	if err != nil {
		Fatalf("Event ApprovalForAll decode error: %v",err)
	} else {
		Info.Printf("ApprovalForAll event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump(Info)
	}
}
func proc_trading_proceeds_claimed(agtx *AugurTx,timestamp int64,log *types.Log) {

	var mevt TradingProceedsClaimed
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
	var mevt FillEvt
	mevt.MakerAddress= common.BytesToAddress(log.Topics[1][12:])
	mevt.FeeRecipientAddress= common.BytesToAddress(log.Topics[2][12:])
	mevt.OrderHash= log.Topics[3]
	err := exchange_abi.Unpack(&mevt,"Fill",log.Data)
	if err != nil {
		Fatalf("Event Fill for 0x decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.AugurTrading.Bytes()) {
		return
	}
	Info.Printf("Fill event found (block=%v) :\n",log.BlockNumber)
	mevt.Dump(Info)
	// we need to locate order id because Profit Loss events are linked to this Order 
	fill_order_id = storage.Locate_fill_event_order(&mevt)
}
func proc_erc20_transfer(log *types.Log,agtx *AugurTx) {
	var mevt Transfer
	/*
	*/
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
	var mevt ProfitLossChanged
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
//		Info.Printf("position_changes len=%v\n",lken(position_changes)
	eoa_aid := get_eoa_aid(&mevt.Account,agtx.BlockNum,agtx.TxId)
	id = storage.Insert_profit_loss_evt(agtx,eoa_aid,&mevt)
	return id
}
func proc_transfer_single(log *types.Log) {
	var mevt TransferSingle
	mevt.Operator= common.BytesToAddress(log.Topics[1][12:])
	mevt.From= common.BytesToAddress(log.Topics[2][12:])
	mevt.To= common.BytesToAddress(log.Topics[3][12:])
	err := zerox_abi.Unpack(&mevt,"TransferSingle",log.Data)
	if err != nil {
		Fatalf("Event TransferSingle decode error: %v",err)
	} else {
		Info.Printf("TransferSingle event found (block=%v) :\n",log.BlockNumber)
		mevt.Dump(Info)
	}
}
func proc_transfer_batch(log *types.Log) {
	var mevt TransferBatch
	mevt.Operator= common.BytesToAddress(log.Topics[1][12:])
	mevt.From= common.BytesToAddress(log.Topics[2][12:])
	mevt.To= common.BytesToAddress(log.Topics[3][12:])
	err := zerox_abi.Unpack(&mevt,"TransferBatch",log.Data)
	if err != nil {
		Fatalf("Event TransferBatch decode error: %v",err)
	} else {
		Info.Printf("TransferBatch event found (block=%v) :\n",log.BlockNumber)
		mevt.Dump(ctrct_zerox,Info)
	}
}
func proc_tokens_transferred(agtx *AugurTx, log *types.Log) {
	var mevt TokensTransferred
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
	var mevt TokenBalanceChanged
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
	var mevt ShareTokenBalanceChanged
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
func proc_market_order_event(agtx *AugurTx,log *types.Log) {

	var mevt MktOrderEvt
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
	//storage.Insert_market_order_evt(BlockNumber(log.BlockNumber),tx_id,signer,eoa_aid,&mevt)
	storage.Insert_market_order_evt(agtx,eoa_aid,eoa_fill_aid,&mevt)
}
func proc_cancel_zerox_order(log *types.Log) {
	var mevt CancelZeroXOrder
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
	storage.Delete_open_0x_order(ohash_str)
}
func proc_market_oi_changed(block *types.Header, agtx *AugurTx, log *types.Log) {
	var mevt MarketOIChangedEvt
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

	var mevt MktFinalizedEvt
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
	var mevt MktFinalizedEvt
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
	var mevt InitialReportSubmittedEvt
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
	var mevt DisputeCrowdsourcerContributionEvt
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
	var mevt MktVolumeChangedEvt_v1
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
	var mevt MktVolumeChangedEvt_v2
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
	var mevt MarketCreatedEvt
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
	var mevt MarketCreatedEvt
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
	var evt ExecuteTransactionStatus
	err := wallet_abi.Unpack(&evt,"ExecuteTransactionStatus",log.Data)
	if err != nil {
		Fatalf("Event decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		return
	}
	Info.Printf("ExecuteTransactionStatus event found (block=%v) : \n",log.BlockNumber)
	evt.Dump(Info)
	eoa_aid := get_eoa_aid(&common.Address{},agtx.BlockNum,agtx.TxId)
	_ = eoa_aid
	storage.Insert_augur_transaction_status(agtx,&evt)
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
			proc_approval(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_zerox_approval_for_all) {
			proc_approval_for_all(log)
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
			proc_market_order_event(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_cancel_0x_order) {
			proc_cancel_zerox_order(log)
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
			var transf_evt Transfer
			tr_idx := lidx + 1	// the offset to ERC20 event (as they fired by contracts)
			err := cash_abi.Unpack(&transf_evt,"Transfer",(*logs)[tr_idx].Data)
			if err == nil {
				validity_bond = transf_evt.Value.String()
				Info.Printf("extracted validity bond = %v\n",validity_bond)
			}
			tx_insert_if_needed(agtx)
			proc_market_created(agtx,log,validity_bond)
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
func process_block(bnum int64,update_last_block bool) error {

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
	Info.Printf("block hash = %v, num_tx=%v\n",block_hash_str,num_transactions)
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
	err = storage.Insert_block(block_hash_str,header)
	if err != nil {
		err = roll_back_blocks(header)
		if err == nil {
			return nil
		}
		Error.Printf("Unable to recover from chainsplit: %v. Aborting",err)
		os.Exit(1)
	}
	if num_transactions == 0 {
		Info.Printf("block_proc: block: %v EMPTY\n",bnum)
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
		dump_tx_input_if_known(agtx.Input)
		Info.Printf("\ttx: %v of %v : %v at blockNum=%v\n",tnum,num_transactions,agtx.TxHash,bnum)
		Info.Printf("\t from=%v\n",agtx.From)
		Info.Printf("\t to=%v for $%v (%v bytes data)\n",
						agtx.To,agtx.Value,len(agtx.Input))
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
		contains_market_finalized:=false
		sequencer := new(EventSequencer)
		num_logs := len(rcpt.Logs)
		for i:=0 ; i<num_logs ; i++ {
			if len(rcpt.Logs[i].Topics) > 0 {
				Info.Printf(
					"\t\t\tlog %v\n\t\t\t\t\t\t for contract %v (%v of %v items)\n",
					rcpt.Logs[i].Topics[0].String(),rcpt.Logs[i].Address.String(),(i+1),len(rcpt.Logs))
				if 0 == bytes.Compare(rcpt.Logs[i].Topics[0].Bytes(),evt_market_finalized) {
					if is_warp_sync_event(rcpt.Logs[i]) {
						// WarpSync market emits 2 events MarketFFinalized and MarketCreated
						// MarketFinalized doesn't have ProfitLoss events, so we can process it
						// just using inverse order (i.e. considering it as non-MarketFinalized)
					} else {
						contains_market_finalized = true
					}
				}
			}
			sequencer.append_event(rcpt.Logs[i])
		}
		var ordered_list []*types.Log
		if contains_market_finalized {
			// logs with Market finalized event need to have special order
			ordered_list = sequencer.get_events_for_market_finalized_case()
		} else {
			ordered_list = sequencer.get_ordered_event_list()
		}
		num_logs = len(ordered_list)
		pl_entries := make([]int64,0,2);// profit loss entries
		market_order_id = 0
		fill_order_id = 0
		for i:=0 ; i < num_logs ; i++ {
			if len(ordered_list[i].Topics) > 0 {
				Info.Printf(
					"\t\t\tchecking log with sig %v\n\t\t\t\t\t\t for contract %v\n",
					ordered_list[i].Topics[0].String(),
					ordered_list[i].Address.String())
				id := process_event(header,agtx,&ordered_list,i)
				//var id int64 =0
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
		var input_data InputStruct
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
		}
	} else {
		Info.Printf("dump_tx_input: input sig: %v\n",hex.EncodeToString(input_sig[:]))
	}
}
func scan_profit_loss_data_for_debugging(block_num int64,position_changes *[]*PosChg) {
	// this function makes direct calls to contracts to get changes in profit loss and record them
	// right after each block is processed (developed for debugging purposes)

	var copts = new(bind.CallOpts)
	//Info.Printf("position_changes len=%v\n",len(*position_changes))
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
