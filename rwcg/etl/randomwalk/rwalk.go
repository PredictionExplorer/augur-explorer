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
	. "github.com/PredictionExplorer/augur-explorer/rwcg/primitives"
	rwp "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/randomwalk"
)
func build_list_of_inspected_events_layer1(rwalk_aid int64) []InspectedEvent {

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
		InspectedEvent {
			Signature: hex.EncodeToString(evt_transfer[:4]),
			ContractAid: rwalk_aid,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_mint_event[:4]),
			ContractAid: 0,
		},
	)
	return inspected_events
}
func proc_new_offer(log *types.Log,elog *EthereumEventLog) {

	var evt rwp.NewOffer
	var eth_evt rwp.ENewOffer

	/*if !bytes.Equal(log.Address.Bytes(),market_addr.Bytes()) {
		Info.Printf("Skipping another instance of MarketPlace contract %v\n",log.Address.String())
		return
	}*/
	Info.Printf("Processing NewOffer event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := marketplace_abi.UnpackIntoInterface(&eth_evt,"NewOffer",log.Data)
	if err != nil {
		Error.Printf("Event NewOffer decode error: %v",err)
		os.Exit(1)
	}

	if !bytes.Equal(log.Address.Bytes(),market_addr.Bytes()) {
		Info.Printf("Event doesn't belong to know address set (addr=%v), skipping\n",log.Address.String())
		return
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Seller = eth_evt.Seller.String()
	evt.Buyer = eth_evt.Buyer.String()
	evt.Price = eth_evt.Price.String()
	evt.RWalkAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.OfferId = log.Topics[2].Big().Int64()
	evt.TokenId = log.Topics[3].Big().Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NewOffer {\n")
	Info.Printf("\tNFT addr: %v\n",evt.RWalkAddr)
	Info.Printf("\tOfferId: %v\n",evt.OfferId)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tSeller: %v\n",evt.Seller)
	Info.Printf("\tBuyer: %v\n",evt.Buyer)
	Info.Printf("}\n")

	storagew.Insert_new_offer(&evt)

}
func proc_item_bought(log *types.Log,elog *EthereumEventLog) {

	var evt rwp.ItemBought

	/*if !bytes.Equal(log.Address.Bytes(),market_addr.Bytes()) {
		Info.Printf("Skipping another instance of MarketPlace contract %v\n",log.Address.String())
		return
	}*/
	Info.Printf("Processing ItemBought id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),market_addr.Bytes()) {
		Info.Printf("Event doesn't belong to know address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.OfferId = log.Topics[1].Big().Int64()
	evt.SellerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.BuyerAddr = common.BytesToAddress(log.Topics[3][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("ItemBought {\n")
	Info.Printf("\tOfferId: %v\n",evt.OfferId)
	Info.Printf("\tSeller: %v\n",evt.SellerAddr)
	Info.Printf("\tBuyer: %v\n",evt.BuyerAddr)
	Info.Printf("}\n")

	storagew.Insert_item_bought(&evt)
}
func proc_offer_cancelled(log *types.Log,elog *EthereumEventLog) {

	var evt rwp.OfferCanceled

	if !bytes.Equal(log.Address.Bytes(),market_addr.Bytes()) {
		Info.Printf("Event doesn't belong to know address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.OfferId = log.Topics[1].Big().Int64()

	Info.Printf("Processing OfferCanceled id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	if !storagew.Offer_exists(log.Address.String(),evt.OfferId) {
		Info.Printf(
			"Skipping OfferCanceled : offer %v for contract %v does not exist, skipping\n",
			evt.OfferId,log.Address.String(),
		)
		return
	}

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("OfferCanceled {\n")
	Info.Printf("\tOfferId: %v\n",evt.OfferId)
	Info.Printf("}\n")

	storagew.Insert_offer_canceled(&evt)
}
func proc_withdrawal(log *types.Log,elog *EthereumEventLog) {

	var evt rwp.Withdrawal
	var eth_evt rwp.EWithdrawalEvent

	Info.Printf("Processing WithdrawalEvent id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	if !bytes.Equal(log.Address.Bytes(),rwalk_addr.Bytes()) {
		Info.Printf("Event doesn't belong to know address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := randomwalk_abi.UnpackIntoInterface(&eth_evt,"WithdrawalEvent",log.Data)
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

	storagew.Insert_withdrawal(&evt)
}
func proc_token_name(log *types.Log,elog *EthereumEventLog) {

	var evt rwp.TokenName
	var eth_evt rwp.ETokenNameEvent

	Info.Printf("Processing TokenName id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	if !bytes.Equal(log.Address.Bytes(),rwalk_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := randomwalk_abi.UnpackIntoInterface(&eth_evt,"TokenNameEvent",log.Data)
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

	if !storagew.RWalk_token_exists(log.Address.String(),evt.TokenId) {
		Info.Printf("Token name event skipped, token contract %v is not registered\n",log.Address.String())
		return
	}

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TokenName {\n")
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tNewName: %v\n",evt.NewName)
	Info.Printf("}\n")

	storagew.Insert_token_name(&evt)
}
func proc_mint_event(log *types.Log,elog *EthereumEventLog) {

	var evt rwp.MintEvent
	var eth_evt rwp.EMintEvent

	Info.Printf("Processing MintEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),rwalk_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := randomwalk_abi.UnpackIntoInterface(&eth_evt,"MintEvent",log.Data)
	if err != nil {
		Error.Printf("Event MintEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId= log.Topics[1].Big().Int64()
	evt.Owner= common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Seed= hex.EncodeToString(eth_evt.Seed[:])
	evt.SeedNum = common.BytesToHash(eth_evt.Seed[:]).Big().String()
	evt.Price = eth_evt.Price.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("MintEvent {\n")
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tOwner %v\n",evt.Owner)
	Info.Printf("\tSeed: %v\n",evt.Seed)
	Info.Printf("\tSeed Numeric: %v\n",evt.SeedNum)
	Info.Printf("\tPrice: %v\n",evt.Price)
	Info.Printf("}\n")

	storagew.Insert_mint_event(&evt)

}
func proc_transfer_event(log *types.Log,elog *EthereumEventLog) {

	var evt rwp.Transfer

	if !bytes.Equal(log.Address.Bytes(),rwalk_addr.Bytes()) {
		Info.Printf("Skipping another instance of RandomWalk contract %v\n",log.Address.String())
		return
	}
	Info.Printf("Processing Token Transfer event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.From = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.To = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenId = log.Topics[3].Big().Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("Transfer {\n")
	Info.Printf("\tFrom: %v\n",evt.From)
	Info.Printf("\tTo: %v\n",evt.To)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("}\n")

	storagew.Insert_token_transfer_event(&evt)
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
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer) {
		proc_transfer_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_mint_event) {
		proc_mint_event(log,evtlog)
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
