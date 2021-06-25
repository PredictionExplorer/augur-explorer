package main

import (
	"bytes"
	"encoding/hex"
	"os"
	"fmt"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func build_list_of_inspected_events() []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events= make([]InspectedEvent,0,32)
	inspected_events = append(inspected_events,
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_pool_created[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_liquidity_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_sports_market_created[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_price_market_created[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_trusted_market_created[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_shares_minted[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_shares_burned[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_shares_swapped[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_settlement_fee_claimed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_protocol_fee_claimed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_winnings_claimed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_market_resolved[:4]),
			ContractAid: 0,
		},
/*		InspectedEvent {
			Signature: hex.EncodeToString(evt_erc20_transfer[:4]),
			ContractAid: 0,
		},*/
	)
	return inspected_events
}
func proc_pool_created(log *types.Log,elog *EthereumEventLog) {

	var evt AA_PoolCreated
	var eth_evt EAMMFactoryPoolCreated

	eth_evt.MarketFactory = common.BytesToAddress(log.Topics[1][12:])
	eth_evt.MarketId = log.Topics[2].Big()
	eth_evt.Creator = common.BytesToAddress(log.Topics[3][12:])

	Info.Printf("Processing PoolCreated event, txhash %v\n",elog.TxHash)

	err := aa_abi.Unpack(&eth_evt,"PoolCreated",log.Data)
	if err != nil {
		Error.Printf("Event PoolCreated decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.PoolAddr = eth_evt.Pool.String()
	evt.TokenRecipientAddr = eth_evt.LpTokenRecipient.String()
	evt.MarketId = eth_evt.MarketId.Int64()
	evt.CreatorAddr = eth_evt.Creator.String()
	evt.FactoryAddr = eth_evt.MarketFactory.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PoolCreated {\n")
	Info.Printf("\tPoolAddr: %v\n",evt.PoolAddr)
	Info.Printf("\tMarketFactory: %v\n",evt.FactoryAddr)
	Info.Printf("\tCreator: %v\n",evt.CreatorAddr)
	Info.Printf("\tMarket Id: %v\n",evt.MarketId)
	Info.Printf("\tlpTokenRecipient: %v\n",evt.TokenRecipientAddr)
	Info.Printf("}\n")

	storage.Insert_aa_pool_created_event(&evt)
}
func proc_liquidity_changed(log *types.Log,elog *EthereumEventLog) {

	var evt AA_LiquidityChanged
	var eth_evt LiquidityChanged

	err := aa_abi.Unpack(&eth_evt,"LiquidityChanged",log.Data)
	if err != nil {
		Error.Printf("Error unpacking LiquidityChanged event: %v\n",err)
		os.Exit(1)
	}
	eth_evt.MarketFactory = common.BytesToAddress(log.Topics[1][12:])
	eth_evt.MarketId = log.Topics[2].Big()
	eth_evt.User = common.BytesToAddress(log.Topics[3][12:])

	Info.Printf("Processing LiquidityChanged event, txhash %v\n",elog.TxHash)
	Info.Printf("Contract: %v\n",log.Address.String())
	evt.MarketId = eth_evt.MarketId.Int64()
	evt.RecipientAddr = eth_evt.Recipient.String()
	evt.UserAddr = eth_evt.User.String()
	evt.MarketFactoryAddr = eth_evt.MarketFactory.String()
	evt.Collateral = eth_evt.Collateral.String()
	evt.LpTokens = eth_evt.LpTokens.String()
	evt.SharesReturned = Bigint_ptr_slice_to_str(&eth_evt.SharesReturned,",")

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("LiquidityChanged {\n")
	Info.Printf("\tMarketId: %v\n",evt.MarketId)
	Info.Printf("\tMarketFactory: %v\n",evt.MarketFactoryAddr)
	Info.Printf("\tRecipientAddr: %v\n",evt.RecipientAddr)
	Info.Printf("\tCollateral: %v\n",evt.Collateral)
	Info.Printf("\tLpTokens: %v\n",evt.LpTokens)
	Info.Printf("\tSharesReturned: %v\n",evt.SharesReturned)
	Info.Printf("}\n")

	storage.Insert_aa_liquidity_changed_event(&evt)
}
func proc_price_market(log *types.Log,elog *EthereumEventLog) {

	var evt AA_PriceMarket
	var eth_evt PriceMarketCreated

	Info.Printf("Processing NewHatchery event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))

	err := aa_abi.Unpack(&eth_evt,"PriceMarketCreated",log.Data)
	if err != nil {
		Error.Printf("Event PriceMarketCreated decode error: %v",err)
		os.Exit(1)
	}
	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.MarketId = eth_evt.Id.Int64()
	evt.CreatorAddr = eth_evt.Creator.String()
	evt.EndTime = eth_evt.EndTime.Int64()
	evt.SpotPrice = eth_evt.SpotPrice.String()

	Info.Printf("PriceMarketCreated{\n")
	Info.Printf("\tId: %v\n",evt.MarketId)
	Info.Printf("\tCreator: %v\n",evt.CreatorAddr)
	Info.Printf("\tEndTime: %v\n",evt.EndTime)
	Info.Printf("\tSpotPrice: %v\n",evt.SpotPrice)
	Info.Printf("}\n")

	storage.Insert_aa_price_market_event(&evt)
}
func proc_sports_market(log *types.Log,elog *EthereumEventLog) {

	var evt AA_SportsMarket
	var eth_evt SportsLinkMarketCreated

	err := aa_abi.Unpack(&eth_evt,"SportsMarketCreated",log.Data)
	if err != nil {
		Error.Printf("Error unpacking SportsMarketCreated event: %v\n",err)
		os.Exit(1)
	}

	event_id := log.Topics[1].Big()
	evt.EventId = event_id.Int64()
	evt.CreatorAddr = eth_evt.Creator.String()
	evt.EndTime = eth_evt.EndTime.Int64()
	evt.MarketId = eth_evt.Id.Int64()
	evt.MarketType = int(eth_evt.MarketType)
	evt.HomeTeamId = eth_evt.HomeTeamId.Int64()
	evt.AwayTeamId = eth_evt.AwayTeamId.Int64()
	evt.EstimatedStarTime = eth_evt.EstimatedStarTime.Int64()
	evt.Score = eth_evt.Score.Int64()

	Info.Printf("Processing SportsMarketCreated event, txhash %v\n",elog.TxHash)
	Info.Printf("Contract = %v\n",log.Address.String())

//	err := cash_abi.Unpack(&mevt,"Approval",log.Data)
//	if err != nil {
//		Fatalf("Event ERC20_Approval Cash decode error: %v",err)
	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	factory,err := NewSportsLinkMarketFactory(factory_addr,eclient)
	if err != nil {
		Error.Printf("Failed to instantiate Factory contract: %v\n",err)
		os.Exit(1)
	}
	var copts = new(bind.CallOpts)
	market_obj,err:=factory.GetMarket(copts,eth_evt.Id)
	if err  != nil {
		Error.Printf("Failed to execute GetMarket(): %v\n",err)
		os.Exit(1)
	}
	var sharetokens string
	for i:=0; i<len(market_obj.ShareTokens); i++ {
		if len(sharetokens) > 0 {
			sharetokens = sharetokens + ","
		}
		sharetokens = sharetokens + market_obj.ShareTokens.String()
	}
	evt.ShareTokens = shartokens

	collateral,err:=factory.Collateral(copts)
	if err != nil {
		Error.Printf("Failed to execute Collateral(): %v\n",err)
		os.Exit(1)
	}
	evt.CollateralAddr = collateral.String()

	sharefactor,err:=factory.ShareFactor(copts)
	if err != nil {
		Error.Printf("Failed to execute ShareFactor(): %v",err)
		os.Exit(1)
	}
	evt.ShareFactor = sharefactor.String()


	Info.Printf("SportsMarketCreated{\n")
	Info.Printf("\tId: %v\n",evt.MarketId)
	Info.Printf("\tCreator: %v\n",evt.CreatorAddr)
	Info.Printf("\tEndTime: %v\n",evt.EndTime)
	Info.Printf("\tMarketType: %v\n",evt.MarketType)
	Info.Printf("\tEventId: %v\n",evt.EventId)
	Info.Printf("\tHomeTeamId: %v\n",evt.HomeTeamId)
	Info.Printf("\tAwayTeamId: %v\n",evt.AwayTeamId)
	Info.Printf("\tEstimatedStarTime: %v\n",evt.EstimatedStarTime)
	Info.Printf("\tScore: %v\n",evt.Score)
	Info.Printf("}\n")

	storage.Insert_aa_sports_market_event(&evt)
}
func proc_trusted_market(log *types.Log,elog *EthereumEventLog) {

	var evt AA_TrustedMarket
	var eth_evt TrustedMarketCreated

	err := aa_abi.Unpack(&eth_evt,"TrustedMarketCreated",log.Data)
	if err != nil {
		Error.Printf("Error unpacking TrustedMarketCreated event: %v\n",err)
		os.Exit(1)
	}

	evt.MarketId= eth_evt.Id.Int64()
	evt.CreatorAddr = eth_evt.Creator.String()
	evt.EndTime = eth_evt.EndTime.Int64()
	evt.Description = eth_evt.Description
	var outcomes string
	for _,outc:= range eth_evt.Outcomes {
		if len(outcomes) > 0 {
			outcomes=outc+ ","
		}
		outcomes = outcomes + outc
	}
	evt.Outcomes = outcomes

	Info.Printf("Processing TrustedMarketCreated event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("TrustedMarketreated{\n")
	Info.Printf("\tId: %v\n",evt.MarketId)
	Info.Printf("\tCreator: %v\n",evt.CreatorAddr)
	Info.Printf("\tEndTime: %v\n",evt.EndTime)
	Info.Printf("\tDescription: %v\n",evt.Description)
	Info.Printf("\tOutcomes: %v\n",evt.Outcomes)
	Info.Printf("}\n")

	storage.Insert_aa_trusted_market_event(&evt)
}
func proc_shares_minted(log *types.Log,elog *EthereumEventLog) {

	var evt AA_SharesMinted
	var eth_evt SharesMinted

	err := aa_abi.Unpack(&eth_evt,"SharesMinted",log.Data)
	if err != nil {
		Error.Printf("Error unpacking SharesMinted event: %v\n",err)
		os.Exit(1)
	}

	Info.Printf("Processing SharesMinted event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))
	evt.Amount = eth_evt.Amount.String()
	evt.MarketId = eth_evt.Id.Int64()
	evt.ReceiverAddr = eth_evt.Receiver.String()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("SharesMinted{\n")
	Info.Printf("\tId: %v\n",evt.MarketId)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tReceiver: %v\n",evt.ReceiverAddr)
	Info.Printf("}\n")

	storage.Insert_aa_shares_minted_event(&evt)
}
func proc_shares_burned(log *types.Log,elog *EthereumEventLog) {

	var evt AA_SharesBurned
	var eth_evt SharesBurned

	err := aa_abi.Unpack(&eth_evt,"SharesBurned",log.Data)
	if err != nil {
		Error.Printf("Error unpacking SharesBurned event: %v\n",err)
		os.Exit(1)
	}

	Info.Printf("Processing SharesBurned event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))
	evt.Amount = eth_evt.Amount.String()
	evt.MarketId = eth_evt.Id.Int64()
	evt.ReceiverAddr = eth_evt.Receiver.String()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("SharesBurned{\n")
	Info.Printf("\tId: %v\n",evt.MarketId)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tReceiver: %v\n",evt.ReceiverAddr)
	Info.Printf("}\n")

	storage.Insert_aa_shares_burned_event(&evt)
}
func proc_shares_swapped(log *types.Log,elog *EthereumEventLog) {

	var evt AA_SharesSwapped
	var eth_evt SharesSwapped

	err := aa_abi.Unpack(&eth_evt,"SharesSwapped",log.Data)
	if err != nil {
		Error.Printf("Error unpacking SharesSwapped event: %v\n",err)
		os.Exit(1)
	}

	eth_evt.MarketFactory = common.BytesToAddress(log.Topics[1][12:])
	eth_evt.MarketId = log.Topics[2].Big()
	eth_evt.User = common.BytesToAddress(log.Topics[3][12:])

	Info.Printf("Processing SharesSwapped event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))
	evt.MarketId = eth_evt.MarketId.Int64()
	evt.UserAddr = eth_evt.User.String()
	evt.MarketFactoryAddr = eth_evt.MarketFactory.String()
	evt.Outcome = eth_evt.Outcome.Int64()
	evt.Collateral = eth_evt.Collateral.String()
	evt.Shares = eth_evt.Shares.String()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("SharesSwapped{\n")
	Info.Printf("\tMarketId: %v\n",evt.MarketId)
	Info.Printf("\tMarketFactory: %v\n",evt.MarketFactoryAddr)
	Info.Printf("\tUserAddr: %v\n",evt.UserAddr)
	Info.Printf("\tOutcome: %v\n",evt.Outcome)
	Info.Printf("\tCollateral: %v\n",evt.Collateral)
	Info.Printf("\tShares: %v\n",evt.Shares)
	Info.Printf("}\n")

	storage.Insert_aa_shares_swapped_event(&evt)
}
func proc_settlement_fee_claimed(log *types.Log,elog *EthereumEventLog) {

	var evt AA_SettlementFeeClaimed
	var eth_evt SettlementFeeClaimed

	err := aa_abi.Unpack(&eth_evt,"SettlementFeeClaimed",log.Data)
	if err != nil {
		Error.Printf("Error unpacking SettlementFeeClaimed event: %v\n",err)
		os.Exit(1)
	}

	eth_evt.Receiver = common.BytesToAddress(log.Topics[1][12:])

	Info.Printf("Processing SettlementFeeClaimed event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))
	evt.SettlementAddr = eth_evt.SettlementAddress.String()
	evt.ReceiverAddr= eth_evt.Receiver.String()
	evt.Amount = eth_evt.Amount.String()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("SettlementFeeClaimed {\n")
	Info.Printf("\tSettlementAddress: %v\n",evt.SettlementAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tReceiver: %v\n",evt.ReceiverAddr)
	Info.Printf("}\n")

	storage.Insert_aa_settlement_fee_claimed_event(&evt)
}
func proc_protocol_fee_claimed(log *types.Log,elog *EthereumEventLog) {

	var evt AA_ProtocolFeeClaimed
	var eth_evt ProtocolFeeClaimed

	err := aa_abi.Unpack(&eth_evt,"ProtocolFeeClaimed",log.Data)
	if err != nil {
		Error.Printf("Error unpacking SettlementFeeClaimed event: %v\n",err)
		os.Exit(1)
	}


	Info.Printf("Processing ProtocolFeeClaimed event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))
	evt.ProtocolAddr = eth_evt.Protocol.String()
	evt.Amount = eth_evt.Amount.String()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("ProtocolFeeClaimed {\n")
	Info.Printf("\tProtocol: %v\n",evt.ProtocolAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storage.Insert_aa_protocol_fee_claimed_event(&evt)
}
func proc_protocol_changed(log *types.Log,elog *EthereumEventLog) {

	var evt AA_ProtocolChanged
	var eth_evt ProtocolChanged

	err := aa_abi.Unpack(&eth_evt,"ProtocolChanged",log.Data)
	if err != nil {
		Error.Printf("Error unpacking ProtocolChanged event: %v\n",err)
		os.Exit(1)
	}

	Info.Printf("Processing ProtocolChanged event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))
	evt.ProtocolAddr = eth_evt.Protocol.String()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("ProtocolChanged {\n")
	Info.Printf("\tProtocol: %v\n",evt.ProtocolAddr)
	Info.Printf("}\n")

	storage.Insert_aa_protocol_changed_event(&evt)
}
func proc_protocol_fee_changed(log *types.Log,elog *EthereumEventLog) {

	var evt AA_ProtocolFeeChanged
	var eth_evt ProtocolFeeChanged

	err := aa_abi.Unpack(&eth_evt,"ProtocolFeeChanged",log.Data)
	if err != nil {
		Error.Printf("Error unpacking ProtocolFeeChanged event: %v\n",err)
		os.Exit(1)
	}

	Info.Printf("Processing ProtocolFeeChanged event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))
	evt.Fee = eth_evt.Fee.String()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("ProtocolFeeChanged {\n")
	Info.Printf("\tFee: %v\n",evt.Fee)
	Info.Printf("}\n")

	storage.Insert_aa_protocol_fee_changed_event(&evt)
}
func proc_settlement_fee_changed(log *types.Log,elog *EthereumEventLog) {

	var evt AA_SettlementFeeChanged
	var eth_evt SettlementFeeChanged

	err := aa_abi.Unpack(&eth_evt,"SettlementFeeChanged",log.Data)
	if err != nil {
		Error.Printf("Error unpacking SettlementFeeChanged event: %v\n",err)
		os.Exit(1)
	}

	Info.Printf("Processing SettlementFeeChanged event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))
	evt.Fee = eth_evt.Fee.String()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("SettlementFeeChanged {\n")
	Info.Printf("\tFee: %v\n",evt.Fee)
	Info.Printf("}\n")

	storage.Insert_aa_settlement_fee_changed_event(&evt)
}
func proc_staker_fee_changed(log *types.Log,elog *EthereumEventLog) {

	var evt AA_StakerFeeChanged
	var eth_evt StakerFeeChanged

	err := aa_abi.Unpack(&eth_evt,"StakerFeeChanged",log.Data)
	if err != nil {
		Error.Printf("Error unpacking StakerFeeChanged event: %v\n",err)
		os.Exit(1)
	}

	Info.Printf("Processing StakerFeeChanged event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))
	evt.Fee = eth_evt.Fee.String()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("StakerFeeChanged {\n")
	Info.Printf("\tFee: %v\n",evt.Fee)
	Info.Printf("}\n")

	storage.Insert_aa_staker_fee_changed_event(&evt)
}
func proc_winnings_claimed(log *types.Log,elog *EthereumEventLog) {

	var evt AA_WinningsClaimed
	var eth_evt WinningsClaimed

	eth_evt.Receiver = common.BytesToAddress(log.Topics[1][12:])
	err := aa_abi.Unpack(&eth_evt,"WinningsClaimed",log.Data)
	if err != nil {
		Error.Printf("Error unpacking WinningsClaimed event: %v\n",err)
		os.Exit(1)
	}

	Info.Printf("Processing WinningsClaimed event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))
	evt.MarketId = eth_evt.Id.Int64()
	evt.WinningOutcomeAddr = eth_evt.WinningOutcome.String()
	evt.Amount=eth_evt.Amount.String()
	evt.SettlementFee = eth_evt.SettlementFee.String()
	evt.Payout = eth_evt.Payout.String()
	evt.ReceiverAddr=eth_evt.Receiver.String()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("WinningsClaimed {\n")
	Info.Printf("\tId: %v\n",evt.MarketId)
	Info.Printf("\tWinningOutcome: %v\n",evt.WinningOutcomeAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tSettlementFee: %v\n",evt.SettlementFee)
	Info.Printf("\tPayout: %v\n",evt.Payout)
	Info.Printf("\tReceiver: %v\n",evt.ReceiverAddr)
	Info.Printf("}\n")

	storage.Insert_aa_winnings_claimed_event(&evt)
}
func proc_market_resolved(log *types.Log,elog *EthereumEventLog) {

	var evt AA_MarketResolved
	var eth_evt MarketResolved

	err := aa_abi.Unpack(&eth_evt,"MarketResolved",log.Data)
	if err != nil {
		Error.Printf("Error unpacking MarketResolved event: %v\n",err)
		os.Exit(1)
	}

	Info.Printf("Processing MarketResolved event, txhash %v\n",elog.TxHash)
	Info.Printf("Contract: %v\n",log.Address.String())
	evt.MarketId = eth_evt.Id.Int64()
	evt.WinnerAddr = eth_evt.Winner.String()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp

	Info.Printf("MarketResolved{\n")
	Info.Printf("\tId: %v\n",evt.MarketId)
	Info.Printf("\tWinnnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("}\n")

	storage.Insert_aa_market_resolved_event(&evt)
}
func proc_erc20_transfer(log *types.Log,elog *EthereumEventLog) {
	var evt AA_FeePotTransfer
	var eth_evt ETransfer
	if len(log.Topics)!=3 {
		Info.Printf(
			"ERC20 transfer event is not compliant log.Topics!=3. Tx (id=%v) hash=%v\n",
			elog.TxId,log.TxHash.String(),
		)
		return
	}
	eth_evt.From= common.BytesToAddress(log.Topics[1][12:])
	eth_evt.To= common.BytesToAddress(log.Topics[2][12:])
	err := aa_abi.Unpack(&eth_evt,"Transfer",log.Data)
	if err != nil {
		Error.Printf("Event ERC20_Transfer, decode error: %v",err)
	}
	evt.Contract = log.Address.String()
	Info.Printf("Processing FeePot ERC20 transfer event, txhash %v\n",elog.TxHash)
	is_feepot := storage.Is_feepot(evt.Contract)
	if !is_feepot {
		Info.Printf("\t not a FeePot address, skipping\n")
		return
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.TimeStamp = elog.TimeStamp
	evt.From=eth_evt.From.String()
	evt.To=eth_evt.To.String()
	evt.Value=eth_evt.Value.String()

	Info.Printf("FeePot Transfer {\n")
	Info.Printf("\tFrom: %v\n",evt.From)
	Info.Printf("\tTo: %v\n",evt.To)
	Info.Printf("\tValue: %v\n",evt.Value)
	Info.Print("}\n")

	storage.Insert_aa_feepot_transfer_event(&evt)
}
func process_arbitrum_augur_event(evt_id int64) error {

	evtlog := storage.Get_event_log(evt_id)
	var log types.Log
	err := rlp.DecodeBytes(evtlog.RlpLog,&log)
	if err!= nil {
		panic(fmt.Sprintf("RLP Decode error: %v",err))
	}
	log.BlockNumber=uint64(evtlog.BlockNum)
	log.TxHash.SetBytes(common.HexToHash(evtlog.TxHash).Bytes())
	log.Address.SetBytes(common.HexToHash(evtlog.ContractAddress).Bytes())
	num_topics := len(log.Topics)
	if num_topics > 0 {
		Info.Printf("found event with sig = %v\n",log.Topics[0].String())
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_pool_created) {
			proc_pool_created(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_liquidity_changed) {
			proc_liquidity_changed(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_price_market_created) {
			proc_price_market(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_sports_market_created) {
			proc_sports_market(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_trusted_market_created) {
			proc_trusted_market(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_shares_minted) {
			proc_shares_minted(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_shares_burned) {
			proc_shares_burned(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_shares_swapped) {
			proc_shares_swapped(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_settlement_fee_claimed) {
			proc_shares_swapped(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_protocol_fee_claimed) {
			proc_shares_swapped(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_protocol_changed) {
			proc_protocol_changed(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_protocol_fee_changed) {
			proc_protocol_fee_changed(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_settlement_fee_changed) {
			proc_settlement_fee_changed(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_staker_fee_changed) {
			proc_staker_fee_changed(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_winnings_claimed) {
			proc_winnings_claimed(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_market_resolved) {
			proc_market_resolved(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc20_transfer) {
			proc_erc20_transfer(&log,&evtlog)
		}
	}
	return nil
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
