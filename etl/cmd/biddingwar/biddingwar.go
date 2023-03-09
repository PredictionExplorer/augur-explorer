package main

import (
	"os"
	"fmt"
	"bytes"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func build_list_of_inspected_events_layer1(rwalk_aid int64) []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events= make([]InspectedEvent,0, 32)
	inspected_events = append(inspected_events,
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_prize_claim_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_bid_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_donation_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_donation_received_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_donation_sent_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_charity_updated[:4]),
			ContractAid: rwalk_aid,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_token_name[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_mint_event[:4]),
			ContractAid: 0,
		},
	)
	return inspected_events
}
func proc_prize_claim_event(log *types.Log,elog *EthereumEventLog) {

	var evt BwPrizeClaimEvent
	var eth_evt BiddingWarPrizeClaimEvent

	Info.Printf("Processing PrizeClaim event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),biddingwar_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := biddingwar_abi.UnpackIntoInterface(&eth_evt,"PrizeClaimEvent",log.Data)
	if err != nil {
		Error.Printf("Event PrizeClaimEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.PrizeNum= log.Topics[1].Big().Int64()
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PrizeClaimEvent {\n")
	Info.Printf("\tPrizeNum: %v\n",evt.PrizeNum)
	Info.Printf("\tWinner%v\n",evt.Address)
	Info.Printf("\tAmount: %v\n",evt.Seed)
	Info.Printf("}\n")

	storage.Insert_prize_claim_event(&evt)
}
func proc_bid_event(log *types.Log,elog *EthereumEventLog) {

	var evt BwBidEvent
	var eth_evt BiddingWarBidEvent

	Info.Printf("Processing Bid event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),biddingwar_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := biddingwar_abi.UnpackIntoInterface(&eth_evt,"BidEvent",log.Data)
	if err != nil {
		Error.Printf("Event BidEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.LastBidderAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.BidPrice = eth_evt.BidPrice.String()
	evt.RandomWalkTokenId = eth_evt.RandomWalkNFTID.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("BidEvent {\n")
	Info.Printf("\tLastBidder: %v\n",evt.LastBidderAddr)
	Info.Printf("\tBidPrice%v\n",evt.BidPrice)
	Info.Printf("\tRandomWalkTokenId: %v\n",evt.RandomWalkTokenId)
	Info.Printf("}\n")

	storage.Insert_mint_event(&evt)
}
func proc_donation_event_event(log *types.Log,elog *EthereumEventLog) {

	var evt BwDonationEvent
	var eth_evt BiddingWarDonationEvent

	Info.Printf("Processing DonationEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),biddingwar_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := biddingwar_abi.UnpackIntoInterface(&eth_evt,"DonationEvent",log.Data)
	if err != nil {
		Error.Printf("Event DonationEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Donor = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonationEvent {\n")
	Info.Printf("\tDonor: %v\n",evt.Donor)
	Info.Printf("\tAmount%v\n",evt.Amount)
	Info.Printf("}\n")

	storage.Insert_donation(&evt)
}
func proc_donation_received_event(log *types.Log,elog *EthereumEventLog) {

	var evt BwDonationReceivedEvent
	var eth_evt CharityWalletDonationReceivedEvent

	Info.Printf("Processing DonationReceivedEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),charitywallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := charity_wallet_abi.UnpackIntoInterface(&eth_evt,"DonationReceivedEvent",log.Data)
	if err != nil {
		Error.Printf("Event DonationReceivedEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Donor = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonationReceivedEvent {\n")
	Info.Printf("\tDonor: %v\n",evt.Donor)
	Info.Printf("\tAmount%v\n",evt.Amount)
	Info.Printf("}\n")

	storage.Insert_donation_received(&evt)
}
func proc_donation_sent_event(log *types.Log,elog *EthereumEventLog) {

	var evt BwDonationSentEvent
	var eth_evt CharityWalletDonationSentEvent

	Info.Printf("Processing DonationSentEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),charitywallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := charity_wallet_abi.UnpackIntoInterface(&eth_evt,"DonationSentEvent",log.Data)
	if err != nil {
		Error.Printf("Event DonationSentEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.CharityAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonationSentEvent {\n")
	Info.Printf("\tCharity: %v\n",evt.CharityAddr)
	Info.Printf("\tAmount%v\n",evt.Amount)
	Info.Printf("}\n")

	storage.Insert_donation_sent(&evt)
}
func proc_charity_updated_event(log *types.Log,elog *EthereumEventLog) {

	var evt BwCharityUpdatedEvent
	var eth_evt CharityWalletCharityUpdatedEvent

	Info.Printf("Processing CharityUpdatedEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),charitywallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := charitywallet_abi.UnpackIntoInterface(&eth_evt,"CharityUpdatedEvent",log.Data)
	if err != nil {
		Error.Printf("Event CharityUpdatedEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCharityAddr = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CharityUpdatedEvent {\n")
	Info.Printf("\tNewCharity: %v\n",evt.CharityAddr)
	Info.Printf("\tAmount%v\n",evt.Amount)
	Info.Printf("}\n")

	storage.Insert_charity_updated_event(&evt)
}
func proc_token_name_event(log *types.Log,elog *EthereumEventLog) {

	var evt BwTokenNameEvent
	var eth_evt CosmicSignatureTokenNameEvent

	Info.Printf("Processing TokenNameEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt,"TokenNameEvent",log.Data)
	if err != nil {
		Error.Printf("Event TokenNameEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = eth_evt.TokenId.Int64()
	evt.TokenName = eth_evt.NewName

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TokenNameEvent {\n")
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tTokenName%v\n",evt.TokenName)
	Info.Printf("}\n")

	storage.Insert_token_name_event(&evt)
}
func proc_mint_event(log *types.Log,elog *EthereumEventLog) {

	var evt BwMintEvent
	var eth_evt CosmicSignatureMintEvent

	Info.Printf("Processing MintEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt,"MintEvent",log.Data)
	if err != nil {
		Error.Printf("Event MintEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenID = log.Topics[1].Big().Int64()
	evt.Owner = common.BytesToAddress(log.Topics[2][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("MintEvent{\n")
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tOwner:%v\n",evt.Owner)
	Info.Printf("\tSeed: %v\n",evt.Seed)
	Info.Printf("}\n")

	storage.Insert_token_name_event(&evt)
}
func select_event_and_process(log *types.Log,evtlog *EthereumEventLog) {

	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_new_offer) {
		proc_prize_claim_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_item_bought) {
		proc_bid_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_offer_canceled) {
		proc_donation_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_withdrawal) {
		proc_donation_received_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_token_name) {
		proc_donation_sent_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer) {
		proc_charity_updated_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_token_name) {
		proc_token_name_event(log,evtlog)
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
