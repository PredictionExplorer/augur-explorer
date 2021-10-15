package main

import (
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
/*
	var evt Pol_ConditionPreparation
	var eth_evt EConditionPreparation

	eth_evt.ConditionId = log.Topics[1]
	eth_evt.Oracle = common.BytesToAddress(log.Topics[2][12:])
	eth_evt.QuestionId = log.Topics[3]

	Info.Printf("Processing ConditionPreparation event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := condtoken_abi.Unpack(&eth_evt,"ConditionPreparation",log.Data)
	if err != nil {
		Error.Printf("Event ConditionPreparation decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ConditionId = hex.EncodeToString(eth_evt.ConditionId[:])
	evt.OracleAddr = eth_evt.Oracle.String()
	evt.QuestionId = hex.EncodeToString(eth_evt.QuestionId[:])
	evt.OutcomeSlotCount = eth_evt.OutcomeSlotCount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("ConditionPreparation{\n")
	Info.Printf("\tConditionId: %v\n",evt.ConditionId)
	Info.Printf("\tOracle: %v\n",evt.OracleAddr)
	Info.Printf("\tQuestionId: %v\n",evt.QuestionId)
	Info.Printf("\tOutcomeSlotCount: %v\n",evt.OutcomeSlotCount)
	Info.Printf("}\n")

	storage.Insert_condition_preparation(&evt)
	*/
}
func proc_item_bought(log *types.Log,elog *EthereumEventLog) {

}
func proc_offer_cancelled(log *types.Log,elog *EthereumEventLog) {

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
