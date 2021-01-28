package main

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"context"
	"os"
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
func build_list_of_inspected_events() []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events= make([]InspectedEvent,0,32)
	inspected_events = append(inspected_events,
	/*
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_market_created[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_market_oi_changed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_market_order[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.AugurTrading.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_market_finalized[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_market_volume_changed_v1[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.AugurTrading.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_market_volume_changed_v2[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.AugurTrading.String(),0,0),
		},
		InspectedEvent {
			Signature: 	hex.EncodeToString(evt_share_token_balance_changed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_cancel_0x_order[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.AugurTrading.String(),0,0),
		},
		//InspectedEvent {
		//	Signature: hex.EncodeToString(evt_transfer_batch[:4]),
		//	ContractAid: storage.Lookup_or_create_address(caddrs.ZeroxTrade.String(),0,0),
		//},
		//InspectedEvent {
		//	Signature: hex.EncodeToString(evt_transfer_single[:4]),
		//	ContractAid: storage.Lookup_or_create_address(caddrs.ZeroxTrade.String(),0,0),
		//},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_profit_loss_changed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.ZeroxTrade.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_trading_proceeds_claimed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_erc1155_approval_for_all[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.ShareToken.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_universe_created[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_register_contract[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_tokens_transferred[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_validity_bond_changed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_noshow_bond_changed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_designated_report_stake_changed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_complete_sets_purchased[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_complete_sets_sold[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_initial_reporter_redeemed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_dispute_crowdsourcer_created[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_dispute_window_created[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
	*/
	/*
		InspectedEvent {
			Signature: hex.EncodeToString(evt_dispute_crowdsourcer_completed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
*/
	/*	InspectedEvent {
			Signature: 	hex.EncodeToString(evt_initial_report_submitted[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_dispute_crowd_contrib[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},*/
		/*InspectedEvent {
			Signature: hex.EncodeToString(evt_token_balance_changed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},*/
		/*InspectedEvent {
			Signature: hex.EncodeToString(evt_reporting_participant_disavowed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},*/
		/*InspectedEvent {
			Signature: hex.EncodeToString(evt_reporting_fee_changed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},*/
		InspectedEvent {
			Signature: hex.EncodeToString(evt_participation_tokens_redeemed[:4]),
			ContractAid: storage.Lookup_or_create_address(caddrs.Augur.String(),0,0),
		},

	)
	return inspected_events
}
func delete_augur_transaction_related_data(tx_id int64) {

	// Note: the list of DELETEs must match the list of event signatures
	//          built in built_list_of_inspected_events() function

	storage.Delete_market_created_evt(tx_id)
	storage.Delete_market_oi_changed_evt(tx_id)
	storage.Delete_market_order_evt(tx_id)
	storage.Delete_market_finalized_evt(tx_id)
	storage.Delete_report_evt(tx_id)
	storage.Delete_market_vol_changed_evt(tx_id)
	storage.Delete_token_balance_changed_evt(tx_id)
	storage.Delete_share_balance_changed_evt(tx_id)
	storage.Delete_cancel_open_order_evt(tx_id)
	storage.Delete_profit_loss_evt(tx_id)
	storage.Delete_trading_proceeds_claimed_evt(tx_id)
	storage.Delete_register_contract_evt(tx_id)
	storage.Delete_claim_funds(tx_id)
	storage.Delete_validity_bond_changed_event(tx_id)
	storage.Delete_noshow_bond_changed_event(tx_id)
	storage.Delete_dispute_crowdsourcer_created(tx_id)
	storage.Delete_dispute_window_created(tx_id)
	storage.Delete_complete_sets_purchased(tx_id)
	storage.Delete_complete_sets_sold(tx_id)
	storage.Delete_initial_reporter_redeemed(tx_id)
	storage.Delete_dispute_crowdsourcer_redeemed(tx_id)
	storage.Delete_reporting_participant_disavowed(tx_id)
	storage.Delete_reporting_fee(tx_id)
	storage.Delete_participation_tokens_redeemed(tx_id)
}
func get_eoa_aid(wallet_addr *common.Address,block_num int64,tx_id int64) int64 {

	var eoa_aid int64 = 0
	wallet_aid,err := storage.Nonfatal_lookup_address_id(wallet_addr.String())
	if err == nil {
		eoa_aid,err = storage.Lookup_eoa_aid(wallet_aid)
		if err == nil {
			return eoa_aid
		}
		// not found
		_,err = storage.Lookup_wallet_aid(wallet_aid)
		if err == nil {
			// it was an EOA , not Wallet addr
			return wallet_aid
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
				tx_lookup_if_needed(*agtx_ptr)
				storage.Set_augur_flag(&mevt.Owner,*agtx_ptr,"ap_0xtrade_on_cash")
				Discover_augur_account(&mevt.Owner,caddrs,*agtx_ptr)
			}
		}
		if bytes.Equal(log.Address.Bytes(),caddrs.Dai.Bytes()) {
			if bytes.Equal(mevt.Spender.Bytes(),caddrs.FillOrder.Bytes()) {
				tx_lookup_if_needed(*agtx_ptr)
				storage.Set_augur_flag(&mevt.Owner,*agtx_ptr,"ap_fill_on_cash")
				Discover_augur_account(&mevt.Owner,caddrs,*agtx_ptr)
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
				tx_lookup_if_needed(*agtx_ptr)
				storage.Set_augur_flag(&mevt.Owner,*agtx_ptr,"ap_fill_on_shtok")
				Discover_augur_account(&mevt.Owner,caddrs,*agtx_ptr)
			}
		}
	}
}
func proc_trading_proceeds_claimed(agtx *AugurTx,timestamp int64,log *types.Log) {

	var mevt ETradingProceedsClaimed
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
	mevt.Shareholder = common.BytesToAddress(log.Topics[2][12:])
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
	storage.Insert_trading_proceeds_claimed_evt(agtx,&mevt)
	storage.Update_claim_status(agtx,&mevt,timestamp)
	Discover_augur_account(&mevt.Shareholder,caddrs,nil)
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
	id = storage.Insert_profit_loss_evt(agtx,&mevt)
	Discover_augur_account(&mevt.Account,caddrs,nil)
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
	outside_augur_ui := true
	if (bytes.Equal(trade_sig,agtx.Input[0:4])) || (bytes.Equal(exec_wtx_sig,agtx.Input[0:4])) ||
		(bytes.Equal(claim_proceeds_sig,agtx.Input[0:4])) {
		outside_augur_ui = false
	}
	storage.Insert_share_balance_changed_evt(agtx,&mevt,outside_augur_ui)
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

	orders,ospecs := extract_orders_from_input(agtx.Input)
	storage.Insert_market_order_evt(agtx,timestamp,&mevt,orders,ospecs)
	Discover_augur_account(&mevt.AddressData[0],caddrs,nil)
	Discover_augur_account(&mevt.AddressData[1],caddrs,nil)
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
		Info.Printf("Couldn't extract fill amount from Tx input: tx hash=%v block %v. Aborting.",agtx.TxHash,agtx.BlockNum)
		Error.Printf("Couldn't extract fill amount from Tx input: tx hash=%v block %v. Aborting.",agtx.TxHash,agtx.BlockNum)
		os.Exit(1)
	}
/*	eoa_aid := get_eoa_aid(&mevt.Account,agtx.BlockNum,agtx.TxId)
	wallet_aid,err := storage.Lookup_wallet_aid(eoa_aid)
	if err!=nil {
		Error.Printf("Lookup of wallet_aid failed for CancelOrder (eoa_aid=%v): %v\n",eoa_aid,err)
	}*/
	aid := storage.Lookup_or_create_address(mevt.Account.String(),agtx.BlockNum,agtx.TxId)
	storage.Cancel_open_order(aid,orders,ospecs,ohash_str,timestamp)
	Discover_augur_account(&mevt.Account,caddrs,nil)
}
func proc_market_oi_changed(timestamp int64, agtx *AugurTx, log *types.Log) {
	var mevt EMarketOIChanged
	err := augur_abi.Unpack(&mevt,"MarketOIChanged",log.Data)
	if err != nil {
		Fatalf("Event decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		return
	}
	Info.Printf("MarketOIChanged event found (block=%v) : \n",agtx.BlockNum)
	mevt.Universe = common.BytesToAddress(log.Topics[1][12:])
	mevt.Market =common.BytesToAddress(log.Topics[2][12:])
	mevt.Dump(Info)
	storage.Insert_market_oi_changed_evt(timestamp,agtx,&mevt)
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
	Discover_augur_account(&mevt.Reporter,caddrs,nil)
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
	Discover_augur_account(&mevt.Reporter,caddrs,nil)
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
	//eoa_aid := get_eoa_aid(&mevt.MarketCreator,agtx.BlockNum,agtx.TxId)
	storage.Insert_market_created_evt(agtx,validity_bond,&mevt)
	Discover_augur_account(&mevt.MarketCreator,caddrs,nil)
}
func proc_transaction_status(agtx *AugurTx, log *types.Log,relayed_from_addr *common.Address) {
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
	var relayed_status string
	if relayed_from_addr == nil {
		relayed_status = "No relay addr."
	} else {
		relayed_status = fmt.Sprintf("Relayed tx, from= %v\n",relayed_from_addr.String())
	}
	Info.Printf("ExecuteTransactionStatus event found (block=%v) %v : \n",log.BlockNumber,relayed_status)
	evt.Dump(Info)
	var owner_addr common.Address
	if relayed_from_addr != nil { // Transaction was relayed via GSN v2
		owner_addr.SetBytes(relayed_from_addr.Bytes())
	} else {
		owner_addr = common.HexToAddress(agtx.From)
	}
	var copts = new(bind.CallOpts)
	wallet_addr,err:=ctrct_wallet_registry.GetWallet(copts,owner_addr)
	if err != nil {
		Error.Printf("Couldn't locate wallet contract for owner %v : %v\n",owner_addr.String(),err)
		os.Exit(1)
	}
	Discover_augur_account(&owner_addr,caddrs,nil)
	wallet_aid := storage.Lookup_address_id(wallet_addr.String())
	storage.Delete_augur_transaction_status(agtx.TxId)
	storage.Insert_augur_transaction_status(wallet_aid,agtx,&evt)
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
	// no DELETE statement here because table 'universe' isn't linked by tx_id
	storage.Insert_universe_created_event(agtx,&evt)
}
func proc_validity_bond_changed(agtx *AugurTx,log *types.Log) {
	var evt EValidityBondChanged
	err := augur_abi.Unpack(&evt,"ValidityBondChanged",log.Data)
	if err != nil {
		Fatalf("Event ValidityBondChanged decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"ValidityBondChanged event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	Info.Printf("ValidityBondChanged event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_validity_bond_changed_event(agtx,&evt)
}
func proc_noshow_bond_changed(agtx *AugurTx,log *types.Log) {
	var evt ENoShowBondChanged
	err := augur_abi.Unpack(&evt,"NoShowBondChanged",log.Data)
	if err != nil {
		Fatalf("Event NoShowBondChanged decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"NoShowBondChanged event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	Info.Printf("NoShowBondChanged event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_noshow_bond_changed_event(agtx,&evt)
}
func proc_dispute_crowdsourcer_created(agtx *AugurTx,timestamp int64,log *types.Log) {
	var evt EDisputeCrowdsourcerCreated
	err := augur_abi.Unpack(&evt,"DisputeCrowdsourcerCreated",log.Data)
	if err != nil {
		Fatalf("Event DisputeCrowdsourcerCreated decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"DisputeCrowdsourcerCreated event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	evt.Market = common.BytesToAddress(log.Topics[2][12:])	// extract universe addr
	Info.Printf("DisputeCrowdsourcerCreated event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_dispute_crowdsourcer_created(agtx,timestamp,&evt)
}
func proc_dispute_window_created(agtx *AugurTx,log *types.Log) {
	var evt EDisputeWindowCreated
	err := augur_abi.Unpack(&evt,"DisputeWindowCreated",log.Data)
	if err != nil {
		Fatalf("Event DisputeWindowCreated decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"DisputeWindowCreated event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	Info.Printf("DisputeWindowCreated event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_dispute_window_created(agtx,&evt)
}
func proc_designated_report_stake_changed(agtx *AugurTx,log *types.Log) {
	var evt EDesignatedReportStakeChanged
	err := augur_abi.Unpack(&evt,"DesignatedReportStakeChanged",log.Data)
	if err != nil {
		Fatalf("Event DesignatedReportStakeChanged decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"DesignatedReportStakeChanged event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	Info.Printf("DesignatedReportStakeChanged event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_designated_report_stake_changed(agtx,&evt)
}
func proc_complete_sets_purchased(agtx *AugurTx,log *types.Log) {
	var evt ECompleteSetsPurchased
	err := augur_abi.Unpack(&evt,"CompleteSetsPurchased",log.Data)
	if err != nil {
		Fatalf("Event CompleteSetsPurchased decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"CompleteSetsPurchased event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	evt.Market = common.BytesToAddress(log.Topics[2][12:])
	evt.Account = common.BytesToAddress(log.Topics[1][12:])
	Info.Printf("CompleteSetsPurchased event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_complete_sets_purchased(agtx,&evt)
}
func proc_complete_sets_sold(agtx *AugurTx,log *types.Log) {
	var evt ECompleteSetsSold
	err := augur_abi.Unpack(&evt,"CompleteSetsSold",log.Data)
	if err != nil {
		Fatalf("Event CompleteSetsSold decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"CompleteSetsSold event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	evt.Market = common.BytesToAddress(log.Topics[2][12:])
	evt.Account = common.BytesToAddress(log.Topics[1][12:])
	Info.Printf("CompleteSetsSold event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_complete_sets_sold(agtx,&evt)
}
func proc_initial_reporter_redeemed(agtx *AugurTx,log *types.Log) {
	var evt EInitialReporterRedeemed
	err := augur_abi.Unpack(&evt,"InitialReporterRedeemed",log.Data)
	if err != nil {
		Fatalf("Event InitialReporterRedeemed decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"InitialReporterRedeemed event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	evt.Reporter = common.BytesToAddress(log.Topics[2][12:])
	evt.Market = common.BytesToAddress(log.Topics[3][12:])
	Info.Printf("InitialReporterRedeemed event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_initial_reporter_redeemed(agtx,&evt)
}
func proc_dispute_crowdsourcer_redeemed(agtx *AugurTx,log *types.Log) {
	var evt EDisputeCrowdsourcerRedeemed
	err := augur_abi.Unpack(&evt,"DisputeCrowdsourcerRedeemed",log.Data)
	if err != nil {
		Fatalf("Event DisputeCrowdsourcerRedeemed decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"DisputeCrowdsourcerRedeemed event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	evt.Reporter = common.BytesToAddress(log.Topics[2][12:])
	evt.Market = common.BytesToAddress(log.Topics[3][12:])
	Info.Printf("DisputeCrowdsourcerRedeemed event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_dispute_crowdsourcer_redeemed(agtx,&evt)
}
func proc_dispute_crowdsourcer_completed(agtx *AugurTx,log *types.Log) {
	var evt EDisputeCrowdsourcerCompleted
	err := augur_abi.Unpack(&evt,"DisputeCrowdsourcerCompleted",log.Data)
	if err != nil {
		Fatalf("Event DisputeCrowdsourcerCompleted decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"DisputeCrowdsourcerCompleted event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	evt.Market = common.BytesToAddress(log.Topics[2][12:])
	Info.Printf("DisputeCrowdsourcerCompleted event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_dispute_crowdsourcer_completed(agtx,&evt)
}
func proc_reporting_participant_disavowed(agtx *AugurTx,timestamp int64,log *types.Log) {
	var evt EReportingParticipantDisavowed
	err := augur_abi.Unpack(&evt,"ReportingParticipantDisavowed",log.Data)
	if err != nil {
		Fatalf("Event ReportingParticipantDisavowed decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"ReportingParticipantDisavowed event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	evt.Market = common.BytesToAddress(log.Topics[2][12:])
	Info.Printf("ReportingParticipantDisavowed event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_reporting_participant_disavowed(agtx,timestamp,&evt)
}
func proc_reporting_fee_changed(agtx *AugurTx,timestamp int64,log *types.Log) {
	var evt EReportingFeeChanged
	err := augur_abi.Unpack(&evt,"ReportingFeeChanged",log.Data)
	if err != nil {
		Fatalf("Event ReportingFeeChanged decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"ReportingFeeChanged event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	Info.Printf("ReportingFeeChanged event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_reporting_fee_changed(agtx,timestamp,&evt)
}
func proc_participation_tokens_redeemed(agtx *AugurTx,log *types.Log) {
	var evt EParticipationTokensRedeemed
	err := augur_abi.Unpack(&evt,"ParticipationTokensRedeemed",log.Data)
	if err != nil {
		Fatalf("Event ParticipationTokensRedeemed decode error: %v",err)
		return
	}
	if !bytes.Equal(log.Address.Bytes(),caddrs.Augur.Bytes()) {
		evt.Dump(Info)
		Info.Printf(
			"ParticipationTokensRedeemed event received and ignored "+
			"(belongs to different contract: %v) at block %v (EVENT_IGNORE)",
			log.Address.String(),agtx.BlockNum,
		)
		return
	}
	evt.Universe = common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
	evt.DisputeWindow = common.BytesToAddress(log.Topics[2][12:])
	evt.Account = common.BytesToAddress(log.Topics[3][12:])
	Info.Printf("ParticipationTokensRedeemed event for contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	evt.Dump(Info)
	storage.Insert_participation_tokens_redeemed(agtx,&evt)
}
func get_tx_relayed_from_addr(logs *[]*types.Log) (*common.Address) {

	var output *common.Address = nil
	for i:=0 ; i < len(*logs) ; i++ {
		log := (*logs)[i]
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_tx_relayed) {
			output := new(common.Address)
			output.SetBytes(log.Topics[2].Bytes())
			return output
		}
	}
	return output
}
func process_event(timestamp int64,agtx *AugurTx,logs *[]*types.Log,lidx int) int64 {
	// Return Value: id of the record inserted (if aplicable, or 0)

	log := &(*(*logs)[lidx])	// we are getting full array of logs (some events need adjacent event data)
	var id int64 = 0
	num_topics := len(log.Topics)
	if num_topics > 0 {
/*
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc20_approval) {
			proc_approval(log,&agtx)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc1155_approval_for_all) {
			proc_approval_for_all(log,&agtx)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_trading_proceeds_claimed) {
			tx_lookup_if_needed(agtx)
			proc_trading_proceeds_claimed(agtx,timestamp,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_exchange_fill) {
			proc_fill_evt(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc20_transfer) {
			tx_lookup_if_needed(agtx)
			proc_erc20_transfer(log,agtx)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_profit_loss_changed) {
			tx_lookup_if_needed(agtx)
			id = proc_profit_loss_changed(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer_single) {
			proc_transfer_single(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer_batch) {
			proc_transfer_batch(log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_share_token_balance_changed) {
			tx_lookup_if_needed(agtx)
			proc_share_token_balance_changed(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_order) {
			tx_lookup_if_needed(agtx)
			proc_market_order_event(agtx,log,timestamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_cancel_0x_order) {
			proc_cancel_zerox_order(agtx,log,timestamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_oi_changed) {
			tx_lookup_if_needed(agtx)
			proc_market_oi_changed(timestamp,agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_finalized) {
			tx_lookup_if_needed(agtx)
			proc_market_finalized_evt(agtx,timestamp,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_volume_changed_v1) {
			tx_lookup_if_needed(agtx)
			proc_market_volume_changed_v2(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_volume_changed_v2) {
			tx_lookup_if_needed(agtx)
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
			tx_lookup_if_needed(agtx)
			proc_market_created(agtx,log,validity_bond)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_execute_tx_status) {
			tx_lookup_if_needed(agtx)
			tx_relayed_from := get_tx_relayed_from_addr(logs)
			proc_transaction_status(agtx,log,tx_relayed_from)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_register_contract) {
			tx_lookup_if_needed(agtx)
			proc_register_contract(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_universe_created) {
			tx_lookup_if_needed(agtx)
			proc_universe_created(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_tokens_transferred) {
			tx_lookup_if_needed(agtx)
			proc_tokens_transferred(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_validity_bond_changed) {
			tx_lookup_if_needed(agtx)
			proc_validity_bond_changed(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_noshow_bond_changed) {
			tx_lookup_if_needed(agtx)
			proc_noshow_bond_changed(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_designated_report_stake_changed) {
			tx_lookup_if_needed(agtx)
			proc_designated_report_stake_changed(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_complete_sets_purchased) {
			tx_lookup_if_needed(agtx)
			proc_complete_sets_sold(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_complete_sets_sold) {
			tx_lookup_if_needed(agtx)
			proc_complete_sets_sold(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_dispute_crowdsourcer_created) {
			tx_lookup_if_needed(agtx)
			proc_dispute_crowdsourcer_created(agtx,timestamp,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_initial_reporter_redeemed) {
			tx_lookup_if_needed(agtx)
			proc_initial_reporter_redeemed(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_dispute_window_created) {
			tx_lookup_if_needed(agtx)
			proc_dispute_window_created(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_dispute_crowdsourcer_redeemed) {
			tx_lookup_if_needed(agtx)
			proc_dispute_crowdsourcer_redeemed(agtx,log)
		}
		*/
		/*if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_dispute_crowdsourcer_completed) {
			tx_lookup_if_needed(agtx)
			proc_dispute_crowdsourcer_completed(agtx,log)
		}*/
		/*if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_initial_report_submitted) {
			tx_lookup_if_needed(agtx)
			proc_initial_report_submitted(agtx,log)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_dispute_crowd_contrib) {
			tx_lookup_if_needed(agtx)
			proc_dispute_crowdsourcerer_contribution(agtx,log)
		}*/
		/*if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_reporting_participant_disavowed) {
			tx_lookup_if_needed(agtx)
			proc_reporting_participant_disavowed(agtx,timestamp,log)
		}*/
		/*if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_reporting_fee_changed) {
			tx_lookup_if_needed(agtx)
			proc_reporting_fee_changed(agtx,timestamp,log)
		}*/
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_participation_tokens_redeemed) {
			tx_lookup_if_needed(agtx)
			proc_participation_tokens_redeemed(agtx,log)
		}
	}
	for j:=1; j < num_topics ; j++ {
		Info.Printf("\t\t\t\tLog Topic %v , %v \n",j,log.Topics[j].String())
	}
	return id
}
func process_transaction(tx_id int64) error {

	tx_hash,bnum,err := storage.Get_tx_hash_by_id(tx_id)
	if err != nil {
		Info.Printf("Can't get transaction by id, id=%v\n",tx_id)
		os.Exit(1)
	}
	hash:=common.HexToHash(tx_hash)
	agtx,err := rpc_get_transaction(hash)
	if err!= nil {
		Info.Printf("Error getting transaction %v: %v\n",tx_hash,err)
		os.Exit(1)
	}
	agtx.TxId = tx_id
	Info.Printf("BLOCK: %v TRANSACTION %v, TXID: %v, input len=%v\n",bnum,hash.String(),tx_id,len(agtx.Input))
//	orders,ospecs := extract_orders_from_input(agtx.Input)
	ctx := context.Background()
	rcpt,err := eclient.TransactionReceipt(ctx,hash)
	if err != nil {
		Info.Printf("Error getting receipt for %v: %v\n",tx_hash,err)
		os.Exit(1)
	}
	if rcpt.Status != types.ReceiptStatusSuccessful {
		Info.Printf("Failed transaction, skipping\n")
		return nil
	}
	timestamp,err := storage.Get_block_timestamp(bnum)
	if err != nil {
		Info.Printf("Error getting timestamp for block %v : %v\n",bnum,err)
		os.Exit(1)
	}
	dump_tx_input_if_known(agtx.Input)
	//Info.Printf("Receipt has %v logs\n",len(receipt.Logs))
	if rcpt.BlockNumber.Int64() != bnum {
		Error.Printf(
			"Transaction's receipt doesn't match current block number. (block possibly changed)" +
			" cur_block_num=%v, receipt.block_num=%v\n",
			bnum,rcpt.BlockNumber.Int64(),
		)
		os.Exit(1)
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
		tx_lookup_if_needed(agtx)
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
			tx_lookup_if_needed(agtx)
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
			storage.Delete_exec_wtx(agtx.TxId)
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

//	delete_augur_transaction_related_data(agtx.TxId)
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
				ordered_list[i].Address.String(),
			)
			id := process_event(timestamp,agtx,&ordered_list,i)
			if 0 == bytes.Compare(ordered_list[i].Topics[0].Bytes(),evt_profit_loss_changed) {
				pl_entries = append(pl_entries,id)
			}
		}
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
		Info.Printf("profit_loss debug: processing pl for %v\n",pchg.Addr.String())
		pl,err := ctrct_pl.GetRealizedProfit(copts,pchg.Mkt_addr,pchg.Addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetRealizedProfit for addr %v outcome %v: %v\n",
							pchg.Addr.String(),pchg.Outcome.String(),err)
			continue
		} else {
			pchg.ProfitLoss = pl
		}
		ff,err := ctrct_pl.GetFrozenFunds(copts,pchg.Mkt_addr,pchg.Addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetFrozenFunds for addr %v outcome %v: %v\n",
							pchg.Addr.String(),pchg.Outcome.String(),err)
			continue
		} else {
			pchg.FrozenFunds = ff
		}
		psize,err := ctrct_pl.GetNetPosition(copts,pchg.Mkt_addr,pchg.Addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetNetPosition for addr %v outcome %v: %v\n",
							pchg.Addr.String(),pchg.Outcome.String(),err)
			continue
		} else {
			pchg.NetPos = psize
		}
		price,err := ctrct_pl.GetAvgPrice(copts,pchg.Mkt_addr,pchg.Addr,pchg.Outcome)
		if err != nil {
			Error.Printf("Failure to GetAvgPricefor addr %v outcome %v: %v\n",
							pchg.Addr.String(),pchg.Outcome.String(),err)
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
			if 0 == bytes.Compare(input_sig,zeroex_cancel_sig) {
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
func Discover_augur_account(addr *common.Address,caddrs *ContractAddresses,agtx_ptr *AugurTx) bool {

	// checks if provided address has a set of approvals required to consider an account Augur-enabled
	var copts = new(bind.CallOpts)
	var agtx AugurTx	// a bogus transaction with empty values
	if agtx_ptr !=  nil {
		agtx.BlockNum = agtx_ptr.BlockNum
		agtx.TxId = agtx_ptr.TxId
	}
	allow_amount,err := ctrct_dai_token.Allowance(copts,*addr,caddrs.ZeroxTrade)
	if err != nil {
		Info.Printf("Discover_augur_account(): allowance() at DAI for ZeroxTrade failed for %v: %v\n",addr.String(),err)
		return false
	}
	z:=big.NewInt(0)
	if z.Cmp(allow_amount) >= 0 {
		Info.Printf("Discover_augur_account(): allowance() amount at DAI for ZeroxTrade is 0 for %v\n",addr.String())
		return false
	}
	storage.Set_augur_flag(addr,&agtx,"ap_0xtrade_on_cash")
	allow_amount,err = ctrct_dai_token.Allowance(copts,*addr,caddrs.FillOrder)
	if err != nil {
		Info.Printf("Discover_augur_account(): allowance() at DAI for FillOrder failed for %v: %v\n",addr.String(),err)
		return false
	}
	if z.Cmp(allow_amount) >= 0 {
		Info.Printf("Discover_augur_account(): allowance() amount at DAI for FillOrder is 0 for %v\n",addr.String())
		return false
	}
	storage.Set_augur_flag(addr,&agtx,"ap_fill_on_cash")

	ctrct_share_token,err := NewShareToken(caddrs.ShareToken,eclient)
	if err != nil {
		Info.Printf("Discover_augur_account(): instantiation of ShareToken contract failed for %v: %v\n",*addr,err)
		return false
	}
	approved4all,err := ctrct_share_token.IsApprovedForAll(copts,*addr,caddrs.FillOrder)
	if err != nil {
		Info.Printf("Discover_augur_account(): isApprovedForAll failed for %v : %v\n",addr.String(),err)
		return false
	}
	if !approved4all {
		Info.Printf("Discover_augur_account(): isApprovedForAll set to false for %v\n",addr.String())
		return false
	}

	storage.Set_augur_flag(addr,&agtx,"ap_fill_on_shtok")
	if storage.Is_augur_activated(addr.String()) {
		Info.Printf("Discovered address %v as augur-enabled (block=%v,tx_id=%v)\n",addr.String(),agtx.BlockNum,agtx.TxId)
		return true
	}
	Info.Printf("Discovery of address %v as augur-enabled not successfull\n",*addr)
	return false
}
func tx_lookup_if_needed(agtx *AugurTx) {
	if agtx.TxId == 0 {
		var err error
		agtx.TxId,err = storage.Get_tx_id_by_hash(agtx.TxHash)
		if err!=nil {
			Info.Printf("Tx lookup failed: txhash=%v\n",agtx.TxHash)
			Error.Printf("Tx lookup failed: txhash=%v\n",agtx.TxHash)
			os.Exit(1)
		}
	}
}
