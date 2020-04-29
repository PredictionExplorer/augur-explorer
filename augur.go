package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/hex"
	"math/big"

//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
func load_abi(fname string) *abi.ABI {

	abi_data, err := ioutil.ReadFile(fname)
	check(err)
	abi_rdr := bytes.NewReader(abi_data)
	check(err)
	abi,err := abi.JSON(abi_rdr)
	check(err)
	return &abi
}
func augur_init() {

	// Augur service involves 39 contracts in total. We only use a few of them

	// Load main Agur contract ABI
	augur_abi = load_abi("./abis/main-abis/Augur.abi")

	// Load AugurTrading contract
	trading_abi = load_abi("./abis/trading-abis/AugurTrading.abi")

	zerox_abi = load_abi("./abis/trading-abis/ZeroXTrade.abi")
	cash_abi = load_abi("./abis/trading-abis/Cash.abi")
	exchange_abi = load_abi("./abis/exchange/Exchange.abi")

	wallet_abi = load_abi("./abis/main-abis/AugurWalletRegistry.abi")
}
/* Discontinued
func get_universe_and_market(log *types.Log) (string,string) {

	bytes := common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	a_universe := bytes.String()
	bytes = common.BytesToAddress(log.Topics[2][12:])	// extract market addr
	a_market := bytes.String()
	return a_universe,a_market
}
*/
func process_event(log *types.Log) {

	if log == nil {
		Fatalf("process_event() received null pointer")
	}
	num_topics := len(log.Topics)
	if num_topics > 0 {
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_exchange_fill) {
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
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc20_transfer) {
			var mevt Transfer
			mevt.From= common.BytesToAddress(log.Topics[1][12:])
			mevt.To= common.BytesToAddress(log.Topics[2][12:])
			err := cash_abi.Unpack(&mevt,"Transfer",log.Data)
			if err != nil {
				Fatalf("Event ERC20_Transfer for Cash decode error: %v",err)
			} else {
				fmt.Printf("ERC20_Transfer event found (block=%v) :\n",log.BlockNumber)
				mevt.Dump()
			}
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_profit_loss_changed) {
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
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer_single) {
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
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer_batch) {
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
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_tokens_transferred) {
			var mevt TokensTransferred
			mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
			mevt.From= common.BytesToAddress(log.Topics[2][12:])	// extract From
			mevt.To= common.BytesToAddress(log.Topics[3][12:])	// extract To
			err := augur_abi.Unpack(&mevt,"TokensTransferred",log.Data)
			if err != nil {
				Fatalf("Event TokensTransferred decode error: %v",err)
			} else {
				fmt.Printf("TokensTransferred event found (block=%v) :\n",log.BlockNumber)
				mevt.Dump()
			}
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_token_balance_changed) {
			var mevt TokenBalanceChanged
			mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
			mevt.Owner= common.BytesToAddress(log.Topics[2][12:])
			err := augur_abi.Unpack(&mevt,"TokenBalanceChanged",log.Data)
			if err != nil {
				Fatalf("Event TokenBalanceChanged decode error: %v",err)
			} else {
				fmt.Printf("TokenBalanceChanged event found (block=%v) :\n",log.BlockNumber)
				mevt.Dump()
			}
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_share_token_balance_changed) {
			var mevt ShareTokenBalanceChanged
			mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
			mevt.Account= common.BytesToAddress(log.Topics[2][12:])
			mevt.Market = common.BytesToAddress(log.Topics[3][12:])
			err := augur_abi.Unpack(&mevt,"ShareTokenBalanceChanged",log.Data)
			if err != nil {
				Fatalf("Event ShareTokenBalanceChanged decode error: %v\n",err)
			} else {
				fmt.Printf("ShareTokenBalanceChanged event found (block=%v) :\n",log.BlockNumber)
				mevt.Dump()
				storage.insert_share_balance_changed_evt(&mevt)
			}
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_order) {
			var mevt MktOrderEvt
			mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
			mevt.Market = common.BytesToAddress(log.Topics[2][12:])
			mevt.EventType = log.Topics[3][31];	// EventType (uint8) which we label as OrderAction
			err := trading_abi.Unpack(&mevt,"OrderEvent",log.Data)
			if err != nil {
				Fatalf("Event OrderEvent decode error: %v",err)
			} else {
				fmt.Printf("OrderEvent event found (block=%v) : \n",log.BlockNumber)
				mevt.Dump()
				storage.insert_market_order_evt(BlockNumber(log.BlockNumber),&mevt)
			}
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_cancel_0x_order) {
			var mevt CancelZeroXOrder 
			mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
			mevt.Market = common.BytesToAddress(log.Topics[2][12:])
			mevt.Account = common.BytesToAddress(log.Topics[3][12:]);
			err := trading_abi.Unpack(&mevt,"CancelZeroXOrder",log.Data)
			if err != nil {
				Fatalf("Event CancelZeroXOrder decode error: %v",err)
			} else {
				fmt.Printf("CancelZeroXOrder event found (block=%v) : \n",log.BlockNumber)
				mevt.Dump()
				storage.insert_cancel_0x_order_evt(&mevt)
			}
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_oi_changed) {
			var mevt MarketOIChangedEvt
			err := augur_abi.Unpack(&mevt,"MarketOIChanged",log.Data)
			if err != nil {
				Fatalf("Event decode error: %v",err)
			} else {
				fmt.Printf("MarketOIChanged event found (block=%v) : \n",log.BlockNumber)
				mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
				mevt.Market =common.BytesToAddress(log.Topics[2][12:])
				mevt.Dump()
				storage.insert_market_oi_changed_evt(&mevt)
			}
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_finalized) {
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
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_initial_report_submitted) {
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
				storage.insert_initial_report_evt(&mevt)
			}
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_dispute_crowd_contrib) {
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
				storage.insert_dispute_crowd_contrib(&mevt)
			}
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_volume_changed) {
			var mevt MktVolumeChangedEvt
			err := trading_abi.Unpack(&mevt,"MarketVolumeChanged",log.Data)
			if err != nil {
				Fatalf("Event MarketVolumeChanged decode error: %v\n",err)
			} else {
				fmt.Printf("MarketVolumeChanged event found (block=%v): \n",log.BlockNumber)
				mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
				mevt.Market = common.BytesToAddress(log.Topics[2][12:])
				mevt.Dump()
				storage.insert_market_volume_changed_evt(&mevt)
			}
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_created) {
			var mevt MarketCreatedEvt
			mevt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
			mevt.MarketCreator = common.BytesToAddress(log.Topics[2][12:])	// extract crator addr
			err := augur_abi.Unpack(&mevt,"MarketCreated",log.Data)
			if err != nil {
				Fatalf("Event MarketCreated decode error: %v",err)
			} else {
				fmt.Printf("MarketCreated event found (block=%v)\n",log.BlockNumber)
				mevt.Dump()
				storage.insert_market_created_evt(&mevt)
			}
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
		fmt.Printf("\tto: %v",input_data.To.String())
		fmt.Printf("\tdata: %v",hex.EncodeToString(input_data.Data[:]))
		fmt.Printf("\tvalue: %v",input_data.Value.String())
		fmt.Printf("\tpayment: %v",input_data.Payment.String())
		fmt.Printf("\treferralAddress:  %v",input_data.ReferralAddress.String())
		fmt.Printf("\tfingerprint: %v",hex.EncodeToString(input_data.Fingerprint[:]))
		fmt.Printf("\tdesiredSignerBalance: %v",input_data.DesiredSignerBalance.String())
		fmt.Printf("\tmaxExchangeRateInDai: %v",input_data.MaxExchangeRateInDai.String())
		fmt.Printf("\trevertOnFaliure: %v",input_data.RevertOnFailure)
		fmt.Printf("\t}")
	} else {
		fmt.Printf("dump_tx_input: input sig (%v) != (%v) registered sig ",
			hex.EncodeToString(input_sig[:]),hex.EncodeToString(decoded_sig))
	}
}
