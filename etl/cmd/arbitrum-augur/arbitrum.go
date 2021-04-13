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
			ContractAid: storage.Lookup_or_create_address(caddrs.AMM_Factory.String(),0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_new_hatchery[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_turbo_created[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_complete_sets_minted[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_complete_sets_burned[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_claim[:4]),
			ContractAid: 0,
		},
	)
	return inspected_events
}
func proc_pool_created(log *types.Log,elog *EthereumEventLog) {

	var evt AA_PoolCreated
	hatchery_addr := common.BytesToAddress(log.Topics[1][12:])
	creator_addr := common.BytesToAddress(log.Topics[3][12:])
	Info.Printf("Processing PoolCreated event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))
	Info.Printf("topics[2] = %v\n",hex.EncodeToString(log.Topics[2][:]))
	pool_addr := common.BytesToAddress(log.Data[12:])

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.PoolAddr = pool_addr.String()
	evt.HatcheryAddr = hatchery_addr.String()
	evt.TurboId = hex.EncodeToString(log.Topics[2][:])
	evt.CreatorAddr = creator_addr.String()

	Info.Printf("PoolCreated {\n")
	Info.Printf("\tPoolAddr: %v\n",pool_addr.String())
	Info.Printf("\tHatchery: %v\n",hatchery_addr.String())
	Info.Printf("\tCreator: %v\n",creator_addr.String())
	Info.Printf("\tTurbo Id: %v\n",evt.TurboId)
	Info.Printf("}\n")

	storage.Insert_aa_pool_created_event(&evt)
}
func proc_new_hatchery(log *types.Log,elog *EthereumEventLog) {

	var evt AA_NewHatchery
	hatchery_addr := common.BytesToAddress(log.Data[12:32])
	collateral_addr := common.BytesToAddress(log.Topics[1][12:])
	sharetoken_addr:= common.BytesToAddress(log.Data[32+12:64])
	feepot_addr:= common.BytesToAddress(log.Data[64+12:64+32])

	Info.Printf("Processing PoolCreated event, txhash %v\n",elog.TxHash)
	Info.Printf("log.Data = %v\n",hex.EncodeToString(log.Data[:]))

//	err := cash_abi.Unpack(&mevt,"Approval",log.Data)
//	if err != nil {
//		Fatalf("Event ERC20_Approval Cash decode error: %v",err)
	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.HatcheryAddr = hatchery_addr.String()
	evt.CollateralAddr= collateral_addr.String()
	evt.FeePotAddr= feepot_addr.String()

	Info.Printf("NewHatchery{\n")
	Info.Printf("\tHatcheryAddress: %v\n",hatchery_addr.String())
	Info.Printf("\tCollateral: %v\n",collateral_addr.String())
	Info.Printf("\tShareToken: %v\n",sharetoken_addr.String())
	Info.Printf("\tFeePot: %v\n",feepot_addr.String())
	Info.Printf("}\n")

	storage.Insert_aa_new_hatchery_event(&evt)
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
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_new_hatchery) {
			proc_new_hatchery(&log,&evtlog)
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
