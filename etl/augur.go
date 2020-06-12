package main

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "augur-extractor/primitives"
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
func augur_init(addresses *ContractAddresses,contracts *map[string]interface{}) {

	Init_contract_addresses(addresses)

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
	ctrct_wallet_registry,err = NewAugurWalletRegistry(common.HexToAddress("0x1FD9274a2FE0E86f5A7b5Bde57b93C8C9b62e21d"), rpcclient)
	if err != nil {
		Fatalf("Failed to instantiate a AugurWalletRegistry contract: %v", err)
	}
	ctrct_zerox, err = NewZeroX(zerox_addr,rpcclient)
	if err != nil {
		Fatalf("Failed to instantiate a ZeroX contract: %v", err)
	}

}
func balance_updater() {
	// go-routine that updates balances of DAI tokens
	// when the record is inserted into dai_transf table it is inserted with balance = 0 because
	// we don't have the previous balance (and we can't get it during the processing because we are
	// processing finalized  blocks and at this stage of the process the order of transfers was lost)
	// Therefore the only way to calculate valid balances for all the accounts involved is to get the
	// balance on the previous block, and run the sequence of balance changes
	// The order of insertion into dai_transf table is valid and we can use it to reproduce the history

	// in order to avoid being a bottleneck this process must run as an independent thread

	transfers := storage.Get_unprocessed_dai_transfers()
	for i := 0 ; i<len(transfers) ; i++ {
		dai_transfer := &transfers[i]
		prev_balance_db,err := storage.Get_previous_balance_from_DB(dai_transfer.Id,dai_transfer.FromAid)
		if err != nil {
			// no balance locally, get it from RPC
			_ = prev_balance_db
		}
		prev_balance_db,err = storage.Get_previous_balance_from_DB(dai_transfer.Id,dai_transfer.ToAid)
		if err != nil {
			// no balance locally, get it from RPC
		}
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
							evt_market_volume_changed,
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
func proc_approval(log *types.Log) {

	var mevt Approval
	mevt.Owner= common.BytesToAddress(log.Topics[1][12:])
	mevt.Spender = common.BytesToAddress(log.Topics[2][12:])
	err := cash_abi.Unpack(&mevt,"Approval",log.Data)
	if err != nil {
		Fatalf("Event ERC20_Approval Cash decode error: %v",err)
	} else {
		Info.Printf("ERC20_Approval event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
	}
}
func proc_approval_for_all(log *types.Log) {

	var mevt ApprovalForAll
	mevt.Owner= common.BytesToAddress(log.Topics[1][12:])
	mevt.Operator= common.BytesToAddress(log.Topics[2][12:])
	err := zerox_abi.Unpack(&mevt,"ApprovalForAll",log.Data)
	if err != nil {
		Fatalf("Event ApprovalForAll decode error: %v",err)
	} else {
		Info.Printf("ApprovalForAll event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
	}
}
func proc_trading_proceeds_claimed(log *types.Log) {

	var mevt TradingProceedsClaimed
	mevt.Universe= common.BytesToAddress(log.Topics[1][12:])
	mevt.Sender= common.BytesToAddress(log.Topics[2][12:])
	err := augur_abi.Unpack(&mevt,"TradingProceedsClaimed",log.Data)
	if err != nil {
		Fatalf("EventTradingProceedsClaimed error: %v",err)
	} else {
		Info.Printf("TradingProceedsClaimed event found (block=%v) :\n",log.BlockNumber)
		mevt.Dump()
	}
}
func proc_fill_evt(log *types.Log) {
	var mevt FillEvt
	mevt.MakerAddress= common.BytesToAddress(log.Topics[1][12:])
	mevt.FeeRecipientAddress= common.BytesToAddress(log.Topics[2][12:])
	mevt.OrderHash= log.Topics[3]
	err := exchange_abi.Unpack(&mevt,"Fill",log.Data)
	if err != nil {
		Fatalf("Event Fill for 0x decode error: %v",err)
	} else {
		Info.Printf("Fill event found (block=%v) :\n",log.BlockNumber)
		mevt.Dump()
		// we need to locate order id because Profit Loss events are linked to this Order 
		fill_order_id = storage.Locate_fill_event_order(&mevt)
	}
}
/*
pending for removal
func proc_erc20_transfer(log *types.Log) {
	var mevt Transfer
	mevt.From= common.BytesToAddress(log.Topics[1][12:])
	mevt.To= common.BytesToAddress(log.Topics[2][12:])
	err := cash_abi.Unpack(&mevt,"Transfer",log.Data)
	if err != nil {
		Fatalf("Event ERC20_Transfer for Cash decode error: %v",err)
	} else {
		fmt.Printf("ERC20_Transfer event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		from_eoa_aid := get_eoa_aid(&mevt.From)
		to_eoa_aid := get_eoa_aid(&mevt.To)
		storage.process_DAI_token_transfer(from_eoa_aid,to_eoa_aid,&mevt)
	}
}
*/
func proc_erc20_transfer(log *types.Log,block_num BlockNumber,tx_id int64) {
	var mevt Transfer
	mevt.From= common.BytesToAddress(log.Topics[1][12:])
	mevt.To= common.BytesToAddress(log.Topics[2][12:])
	err := cash_abi.Unpack(&mevt,"Transfer",log.Data)
	if err != nil {
		Fatalf("Event ERC20_Transfer, decode error: %v",err)
	} else {
		Info.Printf("ERC20_Transfer event, contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		if bytes.Equal(dai_addr.Bytes(), log.Address.Bytes()) {	// this is DAI contract
			storage.Process_DAI_token_transfer(&mevt,block_num,tx_id)
		}
		if bytes.Equal(rep_addr.Bytes(), log.Address.Bytes()) {	// this is DAI contract
			storage.Process_REP_token_transfer(&mevt,block_num,tx_id)
		}
	}
}
func proc_profit_loss_changed(block_num BlockNumber,tx_id int64,log *types.Log) int64  {
	var id int64 = 0
	var mevt ProfitLossChanged
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Market = common.BytesToAddress(log.Topics[2][12:])
	mevt.Account= common.BytesToAddress(log.Topics[3][12:])
	err := trading_abi.Unpack(&mevt,"ProfitLossChanged",log.Data)
	if err != nil {
		Fatalf("Event ProfitLossChanged decode error: %v",err)
	} else {
		Info.Printf("ProfitLossChanged event found (block=%v) :\n",log.BlockNumber)
		mevt.Dump()
		eoa_aid := get_eoa_aid(&mevt.Account,block_num,tx_id)
		id = storage.Insert_profit_loss_evt(block_num,tx_id,eoa_aid,&mevt)
	}
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
		mevt.Dump()
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
		mevt.Dump(ctrct_zerox)
	}
}
func proc_tokens_transferred(block_num BlockNumber,tx_id int64, log *types.Log) {
	var mevt TokensTransferred
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.From= common.BytesToAddress(log.Topics[2][12:])	// extract From
	mevt.To= common.BytesToAddress(log.Topics[3][12:])	// extract To
	err := augur_abi.Unpack(&mevt,"TokensTransferred",log.Data)
	if err != nil {
		Fatalf("Event TokensTransferred decode error: %v",err)
	} else {
		Info.Printf("TokensTransferred event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		storage.Insert_token_transf_evt(&mevt,block_num,tx_id)
	}
}
func proc_token_balance_changed(block_num BlockNumber,tx_id int64,log *types.Log) {
	var mevt TokenBalanceChanged
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Owner= common.BytesToAddress(log.Topics[2][12:])
	err := augur_abi.Unpack(&mevt,"TokenBalanceChanged",log.Data)
	if err != nil {
		Fatalf("Event TokenBalanceChanged decode error: %v",err)
	} else {
		Info.Printf("TokenBalanceChanged event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		storage.Insert_token_balance_changed_evt(&mevt,block_num,tx_id)
	}
}
func proc_share_token_balance_changed(block_num BlockNumber,tx_id int64,log *types.Log) {
	var mevt ShareTokenBalanceChanged
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Account= common.BytesToAddress(log.Topics[2][12:])
	mevt.Market = common.BytesToAddress(log.Topics[3][12:])
	err := augur_abi.Unpack(&mevt,"ShareTokenBalanceChanged",log.Data)
	if err != nil {
		Fatalf("Event ShareTokenBalanceChanged decode error: %v\n",err)
	} else {
		Info.Printf("ShareTokenBalanceChanged event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		storage.Insert_share_balance_changed_evt(block_num,tx_id,&mevt)
	}
}
func get_eoa_aid(addr *common.Address,block_num BlockNumber,tx_id int64) int64 {

	var eoa_aid int64 = 0
	wallet_aid,err := storage.Nonfatal_lookup_address_id(addr.String())
	if err == nil {
		eoa_aid,err = storage.Lookup_eoa_aid(wallet_aid)
		if err != nil {
			num:=big.NewInt(int64(owner_fld_offset))   // 1 is the offset at Storage where EOA is stored
			key:=common.BigToHash(num)
			Info.Printf("daitok: Looking up eoa addr via RPC: %v\n",addr.String())
			eoa,err := rpcclient.StorageAt(context.Background(),*addr,key,nil)
			Info.Printf("daitok: output of rpc: %v\n",hex.EncodeToString(eoa))
			var eoa_addr_str string
			if err == nil {
				eoa_addr_str = common.BytesToAddress(eoa[12:]).String()
			} else {
				Info.Printf("daitok: error at rpc call: %v\n",err)
			}
			eoa_aid = storage.Lookup_or_create_address(eoa_addr_str,block_num,tx_id)
		}
	} else {
			// copied from above, ToDo: generalize
			num:=big.NewInt(int64(owner_fld_offset))   // 1 is the offset at Storage where EOA is stored
			key:=common.BigToHash(num)
			Info.Printf("daitok: Looking up eoa addr via RPC: %v\n",addr.String())
			eoa,err := rpcclient.StorageAt(context.Background(),*addr,key,nil)
			Info.Printf("daitok: output of rpc: %v\n",hex.EncodeToString(eoa))
			var eoa_addr_str string
			if err == nil {
				eoa_addr_str = common.BytesToAddress(eoa[12:]).String()
			} else {
				Info.Printf("daitok: error at rpc call: %v\n",err)
			}
			eoa_aid = storage.Lookup_or_create_address(eoa_addr_str,block_num,tx_id)
	}
	Info.Printf("daitok: Getting eoa_aid for address %v, eoa_aid = %v\n",addr.String(),eoa_aid)
	return eoa_aid
}
func proc_market_order_event(block_num BlockNumber,tx_id int64,log *types.Log,signer common.Address) {

	var mevt MktOrderEvt
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Market = common.BytesToAddress(log.Topics[2][12:])
	mevt.EventType = log.Topics[3][31];	// EventType (uint8) which we label as OrderAction
	err := trading_abi.Unpack(&mevt,"OrderEvent",log.Data)
	if err != nil {
		Fatalf("Event OrderEvent decode error: %v",err)
	} else {
		Info.Printf("OrderEvent event for contract %v (block=%v) : \n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		eoa_aid := get_eoa_aid(&mevt.AddressData[0],block_num,tx_id)
		storage.Insert_market_order_evt(BlockNumber(log.BlockNumber),tx_id,signer,eoa_aid,&mevt)
	}
}
func proc_cancel_zerox_order(log *types.Log) {
	var mevt CancelZeroXOrder
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Market = common.BytesToAddress(log.Topics[2][12:])
	mevt.Account = common.BytesToAddress(log.Topics[3][12:]);
	err := trading_abi.Unpack(&mevt,"CancelZeroXOrder",log.Data)
	if err != nil {
		Fatalf("Event CancelZeroXOrder decode error: %v",err)
	} else {
		ohash := common.BytesToHash(mevt.OrderHash[:])
		ohash_str := ohash.String()
		Info.Printf("CancelZeroXOrder event for contract %v (block=%v) : \n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		storage.Delete_open_0x_order(ohash_str)
	}
}
func proc_market_oi_changed(block *types.Header, log *types.Log) {
	var mevt MarketOIChangedEvt
	err := augur_abi.Unpack(&mevt,"MarketOIChanged",log.Data)
	if err != nil {
		Fatalf("Event decode error: %v",err)
	} else {
		Info.Printf("MarketOIChanged event found (block=%v) : \n",log.BlockNumber)
		mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
		mevt.Market =common.BytesToAddress(log.Topics[2][12:])
		mevt.Dump()
		storage.Insert_market_oi_changed_evt(block,&mevt)
	}
}
func proc_market_finalized_evt(log *types.Log) {
	var mevt MktFinalizedEvt
	err := augur_abi.Unpack(&mevt,"MarketFinalized",log.Data)
	if err != nil {
		Fatalf("Event MktFinalizedEvt decode error: %v\n",err)
	} else {
		Info.Printf("MarketFinalized event found (block=%v) : \n",log.BlockNumber)
		mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
		mevt.Market = common.BytesToAddress(log.Topics[2][12:])	// extract universe addr
		mevt.Dump()
		storage.Insert_market_finalized_evt(&mevt)
	}
}
func proc_initial_report_submitted(block_num BlockNumber,tx_id int64, log *types.Log,signer common.Address) {
	var mevt InitialReportSubmittedEvt
	err := augur_abi.Unpack(&mevt,"InitialReportSubmitted",log.Data)
	if err != nil {
		Fatalf("Event InitialReportSubmittedEvt decode error: %v\n",err)
	} else {
		Info.Printf("InitialReportSubmitted event found (block=%v) : \n",log.BlockNumber)
		mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
		mevt.Reporter= common.BytesToAddress(log.Topics[2][12:])
		mevt.Market = common.BytesToAddress(log.Topics[3][12:])
		mevt.Dump()
		storage.Insert_initial_report_evt(block_num,tx_id,signer,&mevt)
	}
}
func proc_dispute_crowdsourcerer_contribution(block_num BlockNumber,tx_id int64,log *types.Log,signer common.Address) {
	var mevt DisputeCrowdsourcerContributionEvt
	err := augur_abi.Unpack(&mevt,"DisputeCrowdsourcerContribution",log.Data)
	if err != nil {
		Fatalf("Event DisputeCrowdsourcerContribution decode error: %v\n",err)
	} else {
		Info.Printf("DisputeCrowdsourcerContribution event found (block %v) : \n",log.BlockNumber)
		mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
		mevt.Reporter= common.BytesToAddress(log.Topics[2][12:])
		mevt.Market = common.BytesToAddress(log.Topics[3][12:])
		mevt.Dump()
		storage.Insert_dispute_crowd_contrib(block_num,tx_id,signer,&mevt)
	}
}
func proc_market_volume_changed(block_num BlockNumber, tx_id int64, log *types.Log) {
	var mevt MktVolumeChangedEvt
	err := trading_abi.Unpack(&mevt,"MarketVolumeChanged",log.Data)
	if err != nil {
		Fatalf("Event MarketVolumeChanged decode error: %v\n",err)
	} else {
		Info.Printf("MarketVolumeChanged event found (block=%v): \n",log.BlockNumber)
		mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
		mevt.Market = common.BytesToAddress(log.Topics[2][12:])
		mevt.Dump()
		storage.Insert_market_volume_changed_evt(block_num,tx_id,&mevt)
	}
}
func proc_market_created(block_num BlockNumber,tx_id int64,log *types.Log,signer common.Address,validity_bond string) {
	var mevt MarketCreatedEvt
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.MarketCreator = common.BytesToAddress(log.Topics[2][12:])	// extract crator addr
	err := augur_abi.Unpack(&mevt,"MarketCreated",log.Data)
	if err != nil {
		Fatalf("Event MarketCreated decode error: %v",err)
	} else {
		Info.Printf("MarketCreated event found (block=%v)\n",log.BlockNumber)
		mevt.Dump()
		Info.Printf("getwallet: Looking wallet addr via RPC for: %v\n",mevt.MarketCreator.String())
		var copts = new(bind.CallOpts)
		wallet,err := ctrct_wallet_registry.GetWallet(copts,mevt.MarketCreator)
		Info.Printf("getwallet: addr is : %v\n",wallet.String())
		var wallet_addr_str string
		var wallet_aid int64 = 0
		if err == nil {
			wallet_addr_str = wallet.String()
		} else {
			Info.Printf("getwallet: error at rpc call: %v\n",err)
		}
		if !Eth_addr_is_zero(&wallet) {
			wallet_aid = storage.Lookup_or_create_address(wallet_addr_str,block_num,tx_id)
		}
		Info.Printf("getwallet: got wallet_aid = %v for wallet addr %v\n",wallet_aid,wallet_addr_str)
		storage.Insert_market_created_evt(block_num,tx_id,signer,wallet_aid,validity_bond,&mevt)
	}
}
// DISCONTINUED func process_event(block *types.Header, tx_id int64, signer common.Address, log *types.Log) int64 {
func process_event(block *types.Header, tx_id int64, signer common.Address, logs *[]*types.Log,lidx int) int64 {
	// Return Value: id of the record inserted (if aplicable, or 0)

	log := &(*(*logs)[lidx])	// we are getting full array of logs (some events need adjacent event data)

	var id int64 = 0
	block_num := BlockNumber(block.Number.Uint64())
	if log == nil {
		Fatalf("process_event() received null pointer")
	}
	num_topics := len(log.Topics)
	if num_topics > 0 {

		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc20_approval) {
			proc_approval(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_zerox_approval_for_all) {
			proc_approval_for_all(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_trading_proceeds_claimed) {
			proc_trading_proceeds_claimed(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_exchange_fill) {
			proc_fill_evt(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc20_transfer) {
			proc_erc20_transfer(log,block_num,tx_id)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_profit_loss_changed) {
			id = proc_profit_loss_changed(block_num,tx_id,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer_single) {
			proc_transfer_single(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer_batch) {
			proc_transfer_batch(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_tokens_transferred) {
			proc_tokens_transferred(block_num,tx_id,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_token_balance_changed) {
			proc_token_balance_changed(block_num,tx_id,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_share_token_balance_changed) {
			proc_share_token_balance_changed(block_num,tx_id,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_order) {
			proc_market_order_event(block_num,tx_id,log,signer)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_cancel_0x_order) {
			proc_cancel_zerox_order(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_oi_changed) {
			proc_market_oi_changed(block,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_finalized) {
			proc_market_finalized_evt(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_initial_report_submitted) {
			proc_initial_report_submitted(block_num,tx_id,log,signer)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_dispute_crowd_contrib) {
			proc_dispute_crowdsourcerer_contribution(block_num,tx_id,log,signer)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_volume_changed) {
			proc_market_volume_changed(block_num,tx_id,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_created) {
			// we have inverted the events, so the validity bond amount is stored in
			// ERC20 transfer event (it is transfered to the Universe)
			var validity_bond string
			var transf_evt Transfer
			tr_idx := lidx + 1	// the offset to ERC20 event (as they fired by contracts)
			err := cash_abi.Unpack(&transf_evt,"Transfer",(*logs)[tr_idx].Data)
			if err == nil {
				validity_bond = transf_evt.Value.String()
				Info.Printf("extracted validity bond = %v\n",validity_bond)
			}
			proc_market_created(block_num,tx_id,log,signer,validity_bond)
		}
	}
	for j:=1; j < num_topics ; j++ {
		Info.Printf("\t\t\t\tLog Topic %v , %v \n",j,log.Topics[j].String())
	}
	return id
}
func dump_tx_input_if_known(tx *types.Transaction) {

	tx_data:=tx.Data()
	if len(tx_data) < 32 {
		return
	}
	input_sig := tx_data[:4]
	decoded_sig ,_ := hex.DecodeString("78dc0eed")
	set_timestamp_sig ,_ := hex.DecodeString("a0a2b573")
	if 0 == bytes.Compare(input_sig,set_timestamp_sig) {
		Info.Printf("Skipping setTimestamp() transaction\n")
		return
	}
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
	} else {
		Info.Printf("dump_tx_input: input sig: %v\n",hex.EncodeToString(input_sig[:]))
	}
}
