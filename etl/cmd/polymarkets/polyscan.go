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
func build_list_of_inspected_events() []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events= make([]InspectedEvent,0,32)
	inspected_events = append(inspected_events,
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_condition_preparation[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_condition_resolution[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_payout_redemption[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_position_split[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_position_merge[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_uri[:4]),
			ContractAid: 0,
		},
	)
	return inspected_events
}
func proc_condition_preparation(log *types.Log,elog *EthereumEventLog) {

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
}
func proc_condition_resolution(log *types.Log,elog *EthereumEventLog) {

	var evt Pol_ConditionResolution
	var eth_evt EConditionResolution

	eth_evt.ConditionId = log.Topics[1]
	eth_evt.Oracle = common.BytesToAddress(log.Topics[2][12:])
	eth_evt.QuestionId = log.Topics[3]

	Info.Printf("Processing ConditionResolution event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := condtoken_abi.Unpack(&eth_evt,"ConditionResolution",log.Data)
	if err != nil {
		Error.Printf("Event ConditionResolution decode error: %v",err)
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
	Info.Printf("ConditionResolution{\n")
	Info.Printf("\tConditionId: %v\n",evt.ConditionId)
	Info.Printf("\tOracle: %v\n",evt.OracleAddr)
	Info.Printf("\tQuestionId: %v\n",evt.QuestionId)
	Info.Printf("\tOutcomeSlotCount: %v\n",evt.OutcomeSlotCount)
	Info.Printf("}\n")

	storage.Insert_condition_resolution(&evt)
}
func proc_position_split(log *types.Log,elog *EthereumEventLog) {

	var evt Pol_PositionSplit
	var eth_evt EPositionSplit

	eth_evt.Stakeholder = common.BytesToAddress(log.Topics[1][12:])
	eth_evt.ParentCollectionId = log.Topics[2]
	eth_evt.ConditionId = log.Topics[3]

	Info.Printf("Processing PositionSplit event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := condtoken_abi.Unpack(&eth_evt,"PositionSplit",log.Data)
	if err != nil {
		Error.Printf("Event PositionSplit decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ConditionId = hex.EncodeToString(eth_evt.ConditionId[:])
	evt.StakeHolderAddr = eth_evt.Stakeholder.String()
	evt.CollateralToken = eth_evt.CollateralToken.String()
	evt.ParentCollectionId = hex.EncodeToString(eth_evt.ParentCollectionId[:])
	evt.Partition = Bigint_ptr_slice_to_str(&eth_evt.Partition,",")
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PositionSplit {\n")
	Info.Printf("\tStakeholder: %v\n",evt.StakeHolderAddr)
	Info.Printf("\tConditionId: %v\n",evt.ConditionId)
	Info.Printf("\tCollateralToken: %v\n",evt.CollateralToken)
	Info.Printf("\tParentCollectionId: %v\n",evt.ParentCollectionId)
	Info.Printf("\tPartition: %v\n",evt.Partition)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storage.Insert_position_split(&evt)
}
func proc_position_merge(log *types.Log,elog *EthereumEventLog) {

	var evt Pol_PositionMerge
	var eth_evt EPositionsMerge

	eth_evt.Stakeholder = common.BytesToAddress(log.Topics[1][12:])
	eth_evt.ParentCollectionId = log.Topics[2]
	eth_evt.ConditionId = log.Topics[3]

	Info.Printf("Processing PositionMerge event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := condtoken_abi.Unpack(&eth_evt,"PositionMerge",log.Data)
	if err != nil {
		Error.Printf("Event PositionMerge decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ConditionId = hex.EncodeToString(eth_evt.ConditionId[:])
	evt.StakeHolderAddr = eth_evt.Stakeholder.String()
	evt.CollateralToken = eth_evt.CollateralToken.String()
	evt.ParentCollectionId = hex.EncodeToString(eth_evt.ParentCollectionId[:])
	evt.Partition = Bigint_ptr_slice_to_str(&eth_evt.Partition,",")
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PositionMerge {\n")
	Info.Printf("\tStakeholder: %v\n",evt.StakeHolderAddr)
	Info.Printf("\tConditionId: %v\n",evt.ConditionId)
	Info.Printf("\tCollateralToken: %v\n",evt.CollateralToken)
	Info.Printf("\tParentCollectionId: %v\n",evt.ParentCollectionId)
	Info.Printf("\tPartition: %v\n",evt.Partition)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storage.Insert_position_merge(&evt)
}
func proc_payout_redemption(log *types.Log,elog *EthereumEventLog) {

	var evt Pol_PayoutRedemption
	var eth_evt EPayoutRedemption

	eth_evt.Redeemer = common.BytesToAddress(log.Topics[1][12:])
	eth_evt.CollateralToken = common.BytesToAddress(log.Topics[3][12:])
	eth_evt.ParentCollectionId = log.Topics[3]

	Info.Printf("Processing PayoutRedemption event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := condtoken_abi.Unpack(&eth_evt,"PayoutRedemption",log.Data)
	if err != nil {
		Error.Printf("Event PayoutRedemption decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ConditionId = hex.EncodeToString(eth_evt.ConditionId[:])
	evt.Redeemer = eth_evt.Redeemer.String()
	evt.CollateralToken = eth_evt.CollateralToken.String()
	evt.ParentCollectionId = hex.EncodeToString(eth_evt.ParentCollectionId[:])
	evt.IndexSets = Bigint_ptr_slice_to_str(&eth_evt.IndexSets,",")
	evt.Payout = eth_evt.Payout.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PayoutRedemption{\n")
	Info.Printf("\tRedemer: %v\n",evt.Redeemer)
	Info.Printf("\tConditionId: %v\n",evt.ConditionId)
	Info.Printf("\tCollateralToken: %v\n",evt.CollateralToken)
	Info.Printf("\tParentCollectionId: %v\n",evt.ParentCollectionId)
	Info.Printf("\tIndexSets: %v\n",evt.IndexSets)
	Info.Printf("\tPayout: %v\n",evt.Payout)
	Info.Printf("}\n")

	storage.Insert_payout_redemption(&evt)
}
func proc_uri(log *types.Log,elog *EthereumEventLog) {

	var evt Pol_URI
	var eth_evt EURI

	eth_evt.Id = common.BytesToHash(log.Topics[1][:]).Big()

	Info.Printf("Processing URI event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := condtoken_abi.Unpack(&eth_evt,"URI",log.Data)
	if err != nil {
		Error.Printf("Event URI decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Id = eth_evt.Id.String()
	evt.Value = eth_evt.Value

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("URI {\n")
	Info.Printf("\tId: %v\n",evt.Id)
	Info.Printf("\tValue: %v\n",evt.Value)
	Info.Printf("}\n")

	storage.Insert_URI(&evt)
}
func process_polymarket_event(evt_id int64) error {

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
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_condition_preparation) {
			proc_condition_preparation(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_condition_resolution) {
			proc_condition_resolution(&log,&evtlog)
		}
		/*if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_payout_redemption) {
			proc_payout_redemption(&log,&evtlog)
		}*/
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_position_split) {
			proc_position_split(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_position_merge) {
			proc_position_merge(&log,&evtlog)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_uri) {
			proc_uri(&log,&evtlog)
		}
	}
	return nil
}
