package main

import (
	"os"
	"fmt"
	"bytes"
	//"context"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"

	//ethereum "github.com/ethereum/go-ethereum"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func build_list_of_inspected_events_layer1() []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events= make([]InspectedEvent,0,32)
	inspected_events = append(inspected_events,
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_new_offer[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_item_bought[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_offer_canceled[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_withdrawal[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_token_name[:4]),
			ContractAid: 0,
		},
	)
	return inspected_events
}
func proc_new_offer(log *types.Log,elog *EthereumEventLog) {

	var evt RW_NewOffer
	var eth_evt ERandomWalk_NewOffer

	if !bytes.Equal(log.Address.Bytes(),market_addr.Bytes()) {
		Info.Printf("Skipping another instance of MarketPlace contract %v\n",log.Address.String())
		return
	}
	Info.Printf("Processing NewOffer event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := marketplace_abi.Unpack(&eth_evt,"NewOffer",log.Data)
	if err != nil {
		Error.Printf("Event NewOffer decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Seller = eth_evt.Seller.String()
	evt.Buyer = eth_evt.Buyer.String()
	evt.Price = eth_evt.Price.String()
	evt.OfferId = log.Topics[1].Big().Int64()
	evt.TokenId = log.Topics[2].Big().Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NewOffer {\n")
	Info.Printf("\tOfferId: %v\n",evt.OfferId)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tSeller: %v\n",evt.Seller)
	Info.Printf("\tBuyer: %v\n",evt.Buyer)
	Info.Printf("}\n")

	storage.Insert_new_offer(&evt)

}
func proc_item_bought(log *types.Log,elog *EthereumEventLog) {

	var evt RW_ItemBought

	if !bytes.Equal(log.Address.Bytes(),market_addr.Bytes()) {
		Info.Printf("Skipping another instance of MarketPlace contract %v\n",log.Address.String())
		return
	}
	Info.Printf("Processing ItemBought id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.OfferId = log.Topics[1].Big().Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("ItemBought {\n")
	Info.Printf("\tOfferId: %v\n",evt.OfferId)
	Info.Printf("}\n")

	storage.Insert_item_bought(&evt)
}
func proc_offer_cancelled(log *types.Log,elog *EthereumEventLog) {

	var evt RW_OfferCanceled

	if !bytes.Equal(log.Address.Bytes(),market_addr.Bytes()) {
		Info.Printf("Skipping another instance of MarketPlace contract %v\n",log.Address.String())
		return
	}
	Info.Printf("Processing OfferCanceled id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.OfferId = log.Topics[1].Big().Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("OfferCanceled {\n")
	Info.Printf("\tOfferId: %v\n",evt.OfferId)
	Info.Printf("}\n")

	storage.Insert_offer_canceled(&evt)
}
func proc_withdrawal(log *types.Log,elog *EthereumEventLog) {

	var evt RW_Withdrawal
	var eth_evt ERandomWalk_WithdrawalEvent

	if !bytes.Equal(log.Address.Bytes(),rwalk_addr.Bytes()) {
		Info.Printf("Skipping another instance of RandomWalk contract %v\n",log.Address.String())
		return
	}
	Info.Printf("Processing WithdrawalEvent id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := randomwalk_abi.Unpack(&eth_evt,"WithdrawalEvent",log.Data)
	if err != nil {
		Error.Printf("Event WithdrawalEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = eth_evt.TokenId.Int64()
	evt.Destination = eth_evt.Destination.String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("Withdrawal {\n")
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tDestination: %v\n",evt.Destination)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storage.Insert_withdrawal(&evt)
}
func proc_token_name(log *types.Log,elog *EthereumEventLog) {

	var evt RW_TokenName
	var eth_evt ERandomWalk_TokenNameEvent

	if !bytes.Equal(log.Address.Bytes(),rwalk_addr.Bytes()) {
		Info.Printf("Skipping another instance of RandomWalk contract %v\n",log.Address.String())
		return
	}
	Info.Printf("Processing TokenName id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := randomwalk_abi.Unpack(&eth_evt,"TokenNameEvent",log.Data)
	if err != nil {
		Error.Printf("Event TokenName decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = eth_evt.TokenId.Int64()
	evt.NewName= eth_evt.NewName

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TokenName {\n")
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tNewName: %v\n",evt.NewName)
	Info.Printf("}\n")

	storage.Insert_token_name(&evt)
}
func select_event_and_process(log *types.Log,evtlog *EthereumEventLog) {

	Info.Printf("processing event with sig = %v\n",log.Topics[0].String())
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_new_offer) {
		proc_new_offer(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_item_bought) {
		proc_item_bought(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_offer_canceled) {
		proc_offer_cancelled(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_withdrawal) {
		proc_withdrawal(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_token_name) {
		proc_token_name(log,evtlog)
	}
}
func process_single_event(evt_id int64) error {

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
		select_event_and_process(&log,&evtlog)
	}
	return nil
}
