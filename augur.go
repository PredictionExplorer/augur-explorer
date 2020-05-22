package main

import (
	"fmt"
	"os"
	"bytes"
	"io/ioutil"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)
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
	return sequencer.unordered_list
/* temporarily disabled
	// at this moment we just reverse the events. more logic will follow later if needed
	output := make([]*types.Log,0,8)
	for i := len(sequencer.unordered_list) - 1; i >= 0; i-- {
		output = append(output,sequencer.unordered_list[i])
	}
	return output
*/
}
func dump_abi_events(a *abi.ABI) {

	fmt.Printf("Events:\n")
	for evt:=range a.Events {
		fmt.Printf("\t%v\t%v\n",a.Events[evt].ID().String(),evt)
	}

}
func dump_abi_methods(a *abi.ABI) {
	fmt.Printf("Methods:\n")
	for meth := range a.Methods {
		fmt.Printf("\t%v\t%v\n",hex.EncodeToString(a.Methods[meth].ID()),meth)
	}
}
func dump_all_artifacts() {

	for contract_name , _ := range all_contracts {
		fmt.Printf("Contract: %v\n",contract_name)
		abi:=abi_from_artifacts(contract_name)
		dump_abi_events(abi)
		dump_abi_methods(abi)
	}
}
func load_all_artifacts(filename string) map[string]interface{} {

	abi_data, err := ioutil.ReadFile("./abis/augur-artifacts-abi.json")
	check(err)
	all_abis_rdr := bytes.NewReader(abi_data)
	check(err)
	byte_data, err := ioutil.ReadAll(all_abis_rdr)
	check(err)
	var all_contracts map[string]interface{}
	json.Unmarshal([]byte(byte_data), &all_contracts)
	return all_contracts
}
func abi_from_artifacts(contract string) *abi.ABI {

	contract_abi:=all_contracts[contract]
	contract_bytes, _ := json.Marshal(contract_abi) // convert back to JSON so Ethereum package can work
	rdr := bytes.NewReader(contract_bytes)
	ctrct_abi,err := abi.JSON(rdr)
	check(err)
	return &ctrct_abi
}
func load_abi(fname string) *abi.ABI {

	abi_data, err := ioutil.ReadFile(fname)
	check(err)
	abi_rdr := bytes.NewReader(abi_data)
	check(err)
	abi,err := abi.JSON(abi_rdr)
	check(err)
	return &abi
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
func dev_init() {
	dai_addr = common.HexToAddress("5f3341EA5989aD3129E325027b8d908b63709A00")
}
func prod_init() {
	dai_addr = common.HexToAddress("6B175474E89094C44Da98b954EedeAC495271d0F")
}
func augur_init() {

	augur_prod := os.Getenv("AUGUR_PROD")
	if len(augur_prod) > 0 {
		prod_init()
	} else {
		dev_init()
	}
	fmt.Printf("Initialization...\n")
	fmt.Printf("DAI (Cash) contract address: %v\n",dai_addr.String())

	all_contracts = load_all_artifacts("./abis/augur-artifacts-abi.json")
	//dump_all_artifacts()

	// Augur service involves 39 contracts in total. We only use a few of them
	augur_abi = abi_from_artifacts("Augur")
	trading_abi = abi_from_artifacts("AugurTrading")
	zerox_abi = abi_from_artifacts("ZeroXTrade")
	cash_abi = abi_from_artifacts("Cash")
	exchange_abi = abi_from_artifacts("Exchange")
	wallet_abi = abi_from_artifacts("AugurWalletRegistry")

	build_list_of_inspected_events()
}
func proc_approval(log *types.Log) {

	var mevt Approval
	mevt.Owner= common.BytesToAddress(log.Topics[1][12:])
	mevt.Spender = common.BytesToAddress(log.Topics[2][12:])
	err := cash_abi.Unpack(&mevt,"Approval",log.Data)
	if err != nil {
		Fatalf("Event ERC20_Approval Cash decode error: %v",err)
	} else {
		fmt.Printf("ERC20_Approval event for contract %v (block=%v) :\n",
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
		fmt.Printf("ApprovalForAll event for contract %v (block=%v) :\n",
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
		fmt.Printf("TradingProceedsClaimed event found (block=%v) :\n",log.BlockNumber)
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
		fmt.Printf("Fill event found (block=%v) :\n",log.BlockNumber)
		mevt.Dump()
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
func proc_erc20_transfer(log *types.Log) {
	var mevt Transfer
	mevt.From= common.BytesToAddress(log.Topics[1][12:])
	mevt.To= common.BytesToAddress(log.Topics[2][12:])
	err := cash_abi.Unpack(&mevt,"Transfer",log.Data)
	if err != nil {
		Fatalf("Event ERC20_Transfer, decode error: %v",err)
	} else {
		fmt.Printf("ERC20_Transfer event, contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		if bytes.Equal(dai_addr.Bytes(), log.Address.Bytes()) {	// this is DAI contract
			storage.process_DAI_token_transfer(&mevt)
		}
	}
}
func proc_profit_loss_changed(block_num BlockNumber,tx_id int64,log *types.Log) {
	var mevt ProfitLossChanged
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Market = common.BytesToAddress(log.Topics[2][12:])
	mevt.Account= common.BytesToAddress(log.Topics[3][12:])
	err := trading_abi.Unpack(&mevt,"ProfitLossChanged",log.Data)
	if err != nil {
		Fatalf("Event ProfitLossChanged decode error: %v",err)
	} else {
		fmt.Printf("ProfitLossChanged event found (block=%v) :\n",log.BlockNumber)
		mevt.Dump()
		eoa_aid := get_eoa_aid(&mevt.Account)
		storage.insert_profit_loss_evt(block_num,tx_id,eoa_aid,&mevt)
	}
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
		fmt.Printf("TransferSingle event found (block=%v) :\n",log.BlockNumber)
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
		fmt.Printf("TransferBatch event found (block=%v) :\n",log.BlockNumber)
		mevt.Dump()
	}
}
func proc_tokens_transferred(log *types.Log) {
	var mevt TokensTransferred
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.From= common.BytesToAddress(log.Topics[2][12:])	// extract From
	mevt.To= common.BytesToAddress(log.Topics[3][12:])	// extract To
	err := augur_abi.Unpack(&mevt,"TokensTransferred",log.Data)
	if err != nil {
		Fatalf("Event TokensTransferred decode error: %v",err)
	} else {
		fmt.Printf("TokensTransferred event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
	}
}
func proc_token_balance_changed(log *types.Log) {
	var mevt TokenBalanceChanged
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.Owner= common.BytesToAddress(log.Topics[2][12:])
	err := augur_abi.Unpack(&mevt,"TokenBalanceChanged",log.Data)
	if err != nil {
		Fatalf("Event TokenBalanceChanged decode error: %v",err)
	} else {
		fmt.Printf("TokenBalanceChanged event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
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
		fmt.Printf("ShareTokenBalanceChanged event for contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		storage.insert_share_balance_changed_evt(block_num,tx_id,&mevt)
	}
}
func get_eoa_aid(addr *common.Address) int64 {

	var eoa_aid int64 = 0
	wallet_aid,err := storage.nonfatal_lookup_address(addr.String())
	if err == nil {
		eoa_aid,err = storage.lookup_eoa_aid(wallet_aid)
		if err != nil {
			num:=big.NewInt(int64(1))   // 1 is the offset at Storage where EOA is stored
			key:=common.BigToHash(num)
			fmt.Printf("daitok: Looking up eoa addr via RPC: %v\n",addr.String())
			eoa,err := rpcclient.StorageAt(context.Background(),*addr,key,nil)
			fmt.Printf("daitok: output of rpc: %v\n",hex.EncodeToString(eoa))
			var eoa_addr_str string
			if err == nil {
				eoa_addr_str = common.BytesToAddress(eoa[12:]).String()
			} else {
				fmt.Printf("daitok: error at rpc call: %v\n",err)
			}
			eoa_aid = storage.lookup_or_create_address(eoa_addr_str)
		}
	} else {
			// copied from above, ToDo: generalize
			num:=big.NewInt(int64(1))   // 1 is the offset at Storage where EOA is stored
			key:=common.BigToHash(num)
			fmt.Printf("daitok: Looking up eoa addr via RPC: %v\n",addr.String())
			eoa,err := rpcclient.StorageAt(context.Background(),*addr,key,nil)
			fmt.Printf("daitok: output of rpc: %v\n",hex.EncodeToString(eoa))
			var eoa_addr_str string
			if err == nil {
				eoa_addr_str = common.BytesToAddress(eoa[12:]).String()
			} else {
				fmt.Printf("daitok: error at rpc call: %v\n",err)
			}
			eoa_aid = storage.lookup_or_create_address(eoa_addr_str)
	}
	fmt.Printf("daitok: Getting eoa_aid for address %v, eoa_aid = %v\n",addr.String(),eoa_aid)
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
		fmt.Printf("OrderEvent event for contract %v (block=%v) : \n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		eoa_aid := get_eoa_aid(&mevt.AddressData[0])
		storage.insert_market_order_evt(BlockNumber(log.BlockNumber),tx_id,signer,eoa_aid,&mevt)
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
		fmt.Printf("CancelZeroXOrder event for contract %v (block=%v) : \n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		storage.delete_open_0x_order(ohash_str)
	}
}
func proc_market_oi_changed(block *types.Header, log *types.Log) {
	var mevt MarketOIChangedEvt
	err := augur_abi.Unpack(&mevt,"MarketOIChanged",log.Data)
	if err != nil {
		Fatalf("Event decode error: %v",err)
	} else {
		fmt.Printf("MarketOIChanged event found (block=%v) : \n",log.BlockNumber)
		mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
		mevt.Market =common.BytesToAddress(log.Topics[2][12:])
		mevt.Dump()
		storage.insert_market_oi_changed_evt(block,&mevt)
	}
}
func proc_market_finalized_evt(log *types.Log) {
	var mevt MktFinalizedEvt
	err := augur_abi.Unpack(&mevt,"MarketFinalized",log.Data)
	if err != nil {
		Fatalf("Event MktFinalizedEvt decode error: %v\n",err)
	} else {
		fmt.Printf("MarketFinalized event found (block=%v) : \n",log.BlockNumber)
		mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
		mevt.Market = common.BytesToAddress(log.Topics[2][12:])	// extract universe addr
		mevt.Dump()
		storage.insert_market_finalized_evt(&mevt)
	}
}
func proc_initial_report_submitted(block_num BlockNumber,tx_id int64, log *types.Log,signer common.Address) {
	var mevt InitialReportSubmittedEvt
	err := augur_abi.Unpack(&mevt,"InitialReportSubmitted",log.Data)
	if err != nil {
		Fatalf("Event InitialReportSubmittedEvt decode error: %v\n",err)
	} else {
		fmt.Printf("InitialReportSubmitted event found (block=%v) : \n",log.BlockNumber)
		mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
		mevt.Reporter= common.BytesToAddress(log.Topics[2][12:])
		mevt.Market = common.BytesToAddress(log.Topics[3][12:])
		mevt.Dump()
		storage.insert_initial_report_evt(block_num,tx_id,signer,&mevt)
	}
}
func proc_dispute_crowdsourcerer_contribution(block_num BlockNumber,tx_id int64,log *types.Log,signer common.Address) {
	var mevt DisputeCrowdsourcerContributionEvt
	err := augur_abi.Unpack(&mevt,"DisputeCrowdsourcerContribution",log.Data)
	if err != nil {
		Fatalf("Event DisputeCrowdsourcerContribution decode error: %v\n",err)
	} else {
		fmt.Printf("DisputeCrowdsourcerContribution event found (block %v) : \n",log.BlockNumber)
		mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
		mevt.Reporter= common.BytesToAddress(log.Topics[2][12:])
		mevt.Market = common.BytesToAddress(log.Topics[3][12:])
		mevt.Dump()
		storage.insert_dispute_crowd_contrib(block_num,tx_id,signer,&mevt)
	}
}
func proc_market_volume_changed(block_num BlockNumber, tx_id int64, log *types.Log) {
	var mevt MktVolumeChangedEvt
	err := trading_abi.Unpack(&mevt,"MarketVolumeChanged",log.Data)
	if err != nil {
		Fatalf("Event MarketVolumeChanged decode error: %v\n",err)
	} else {
		fmt.Printf("MarketVolumeChanged event found (block=%v): \n",log.BlockNumber)
		mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
		mevt.Market = common.BytesToAddress(log.Topics[2][12:])
		mevt.Dump()
		storage.insert_market_volume_changed_evt(block_num,tx_id,&mevt)
	}
}
func proc_market_created(block_num BlockNumber,tx_id int64,log *types.Log,signer common.Address) {
	var mevt MarketCreatedEvt
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	mevt.MarketCreator = common.BytesToAddress(log.Topics[2][12:])	// extract crator addr
	err := augur_abi.Unpack(&mevt,"MarketCreated",log.Data)
	if err != nil {
		Fatalf("Event MarketCreated decode error: %v",err)
	} else {
		fmt.Printf("MarketCreated event found (block=%v)\n",log.BlockNumber)
		mevt.Dump()
		fmt.Printf("getwallet: Looking wallet addr via RPC for: %v\n",mevt.MarketCreator.String())
		var copts = new(bind.CallOpts)
		copts.Pending = true
		wallet,err := ctrct_wallet_registry.GetWallet(copts,mevt.MarketCreator)
		fmt.Printf("getwallet: addr is : %v\n",wallet.String())
		var wallet_addr_str string
		var wallet_aid int64 = 0
		if err == nil {
			wallet_addr_str = wallet.String()
		} else {
			fmt.Printf("getwallet: error at rpc call: %v\n",err)
		}
		if !eth_addr_is_zero(&wallet) {
			wallet_aid = storage.lookup_or_create_address(wallet_addr_str)
		}
		fmt.Printf("getwallet: got wallet_aid = %v for wallet addr %v\n",wallet_aid,wallet_addr_str)
		storage.insert_market_created_evt(block_num,tx_id,signer,wallet_aid,&mevt)
	}
}
func process_event(block *types.Header, tx_id int64, signer common.Address, log *types.Log) {

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
			proc_erc20_transfer(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_profit_loss_changed) {
			proc_profit_loss_changed(block_num,tx_id,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer_single) {
			proc_transfer_single(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer_batch) {
			proc_transfer_batch(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_tokens_transferred) {
			proc_tokens_transferred(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_token_balance_changed) {
			proc_token_balance_changed(log)
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
			proc_market_created(block_num,tx_id,log,signer)
		}
	}
	for j:=1; j < num_topics ; j++ {
		fmt.Printf("\t\t\t\tLog Topic %v , %v \n",j,log.Topics[j].String())
	}
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
		fmt.Printf("Skipping setTimestamp() transaction\n")
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
		fmt.Printf("ExecuteWalletTransaction {\n")
		fmt.Printf("\tto: %v\n",input_data.To.String())
		fmt.Printf("\tdata: %v\n",hex.EncodeToString(input_data.Data[:]))
		fmt.Printf("\tvalue: %v\n",input_data.Value.String())
		fmt.Printf("\tpayment: %v\n",input_data.Payment.String())
		fmt.Printf("\treferralAddress:  %v\n",input_data.ReferralAddress.String())
		fmt.Printf("\tfingerprint: %v\n",hex.EncodeToString(input_data.Fingerprint[:]))
		fmt.Printf("\tdesiredSignerBalance: %v\n",input_data.DesiredSignerBalance.String())
		fmt.Printf("\tmaxExchangeRateInDai: %v\n",input_data.MaxExchangeRateInDai.String())
		fmt.Printf("\trevertOnFaliure: %v\n",input_data.RevertOnFailure)
		fmt.Printf("}\n")
	} else {
		fmt.Printf("dump_tx_input: input sig: %v\n",hex.EncodeToString(input_sig[:]))
	}
}
