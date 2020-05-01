package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/hex"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
func augur_init() {

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
	}
}
func proc_profit_loss_changed(log *types.Log) {
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
		storage.insert_market_order_evt(BlockNumber(log.BlockNumber),tx_id,signer,&mevt)
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
		fmt.Printf("CancelZeroXOrder event for contract %v (block=%v) : \n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump()
		storage.insert_cancel_0x_order_evt(&mevt)
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
		storage.insert_market_created_evt(block_num,tx_id,signer,&mevt)
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
			proc_profit_loss_changed(log)
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
		fmt.Printf("dump_tx_input: tx_data < 32 len")
		return
	}
	input_sig := tx_data[:4]
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
		fmt.Printf("dump_tx_input: input sig (%v) != (%v) registered sig ",
			hex.EncodeToString(input_sig[:]),hex.EncodeToString(decoded_sig))
	}
}
