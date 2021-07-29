package main

import (
	"os"
	"fmt"
	"time"
	"bytes"
	"math/big"
	"context"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"

	ethereum "github.com/ethereum/go-ethereum"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func build_list_of_inspected_events_layer1() []InspectedEvent {

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
		InspectedEvent {
			Signature: hex.EncodeToString(evt_funding_added[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_funding_removed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_fpmm_buy[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_fpmm_sell[:4]),
			ContractAid: 0,
		},
	)
	return inspected_events
}
func build_list_of_inspected_events_filter_logs() []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events= make([]InspectedEvent,0,32)
	inspected_events = append(inspected_events,
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_condition_preparation),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_condition_resolution),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_payout_redemption),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_position_split),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_position_merge),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_uri),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_funding_added),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_funding_removed),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_fpmm_buy),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_fpmm_sell),
			ContractAid: 0,
		},
	)
	return inspected_events
}
func filter_log_query(p_signature []byte,block_num_from,block_num_to int64) ([]types.Log,error){

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
	topics := make([]common.Hash,0,1)
	sig := common.BytesToHash(p_signature)
	topics = append(topics,sig)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(sig.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("block range: %v - %v\n",block_num_from,block_num_to)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	return logs,err
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
func build_erc1155_transfers_for_transaction(tx_id int64,Address,contract_aid int64,sender_aid int64,signature []byte) {
/*
	logs := storage.Get_erc1155_transfers(tx_id,contract_aid,signature)
	for i:=0; i<len(logs); i++ {
		var log types.Log
		err := rlp.DecodeBytes(evtlog.RlpLog,&log)
		if err != nil {
			Error.Printf("Error decoding RLP of event id=%v: %v\n",evtlog.EvtId)
			os.Exit(1)
		}

	}
	*/
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
func proc_funding_added(log *types.Log,elog *EthereumEventLog) {

	var evt Pol_FundingAdded
	var eth_evt EFundingAdded

	eth_evt.Funder = common.BytesToAddress(log.Topics[1][12:])

	Info.Printf("Processing FundingAdded event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := fpmm_abi.Unpack(&eth_evt,"FPMMFundingAdded",log.Data)
	if err != nil {
		Error.Printf("Event FPMMFundingAdded decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Funder = eth_evt.Funder.String()
	evt.AmountsAdded = Bigint_ptr_slice_to_str(&eth_evt.AmountsAdded,",")
	sum_amounts := big.NewInt(0)
	for i:=0; i<len(eth_evt.AmountsAdded) ; i++ {
		sum_amounts.Add(sum_amounts,eth_evt.AmountsAdded[i])
	}
	evt.AllAmountsSummed = sum_amounts.String()
	evt.SharesMinted = eth_evt.SharesMinted.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("FPMMFundingAdded{\n")
	Info.Printf("\tFunder: %v\n",evt.Funder)
	Info.Printf("\tAmountsAdded: %v\n",evt.AmountsAdded)
	Info.Printf("\tAllAmountsSummed: %v\n",evt.AllAmountsSummed)
	Info.Printf("\tSharesMinted: %v\n",evt.SharesMinted)
	Info.Printf("}\n")

	storage.Insert_funding_added(&evt)
}
func proc_funding_removed(log *types.Log,elog *EthereumEventLog) {

	var evt Pol_FundingRemoved
	var eth_evt EFundingRemoved

	eth_evt.Funder = common.BytesToAddress(log.Topics[1][12:])

	Info.Printf("Processing FundingRemoved event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := fpmm_abi.Unpack(&eth_evt,"FPMMFundingRemoved",log.Data)
	if err != nil {
		Error.Printf("Event FPMMFundingAdded decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Funder = eth_evt.Funder.String()
	evt.AmountsRemoved = Bigint_ptr_slice_to_str(&eth_evt.AmountsRemoved,",")
	sum_amounts := big.NewInt(0)
	for i:=0; i<len(eth_evt.AmountsRemoved); i++ {
		sum_amounts.Add(sum_amounts,eth_evt.AmountsRemoved[i])
	}
	evt.AllAmountsSummed = sum_amounts.String()
	evt.SharesBurnt = eth_evt.SharesBurnt.String()
	evt.CollateralRemoved = eth_evt.CollateralRemovedFromFeePool.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("FPMMFundingRemoved {\n")
	Info.Printf("\tFunder: %v\n",evt.Funder)
	Info.Printf("\tAmountsRemoved: %v\n",evt.AmountsRemoved)
	Info.Printf("\tAllAmountsSummed: %v\n",evt.AllAmountsSummed)
	Info.Printf("\tSharesBurnt: %v\n",evt.SharesBurnt)
	Info.Printf("\tCollateralRemovedFromFeePool: %v\n",evt.CollateralRemoved)
	Info.Printf("}\n")

	storage.Insert_funding_removed(&evt)
}
func proc_fpmm_buy(log *types.Log,elog *EthereumEventLog) {

	var evt Pol_Buy
	var eth_evt EBuy


	Info.Printf("Processing FPMMBuy event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := fpmm_abi.Unpack(&eth_evt,"FPMMBuy",log.Data)
	if err != nil {
		Error.Printf("Event FPMMBuy decode error: %v",err)
		os.Exit(1)
	}
	eth_evt.Buyer = common.BytesToAddress(log.Topics[1][12:])
	eth_evt.OutcomeIndex = log.Topics[2].Big()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Buyer= eth_evt.Buyer.String()
	evt.InvestmentAmount = eth_evt.InvestmentAmount.String()
	evt.FeeAmount = eth_evt.FeeAmount.String()
	evt.TokensBought = eth_evt.OutcomeTokensBought.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("FPMMBuy{\n")
	Info.Printf("\tBuyer: %v\n",evt.Buyer)
	Info.Printf("\tInvestmentAmount: %v\n",evt.InvestmentAmount)
	Info.Printf("\tFeeAmount: %v\n",evt.FeeAmount)
	Info.Printf("\tOutcomeIdx: %v\n",evt.OutcomeIdx)
	Info.Printf("\tTokensBought: %v\n",evt.TokensBought)
	Info.Printf("}\n")

	storage.Insert_fpmm_buy(&evt)
}
func proc_fpmm_sell(log *types.Log,elog *EthereumEventLog) {

	var evt Pol_Sell
	var eth_evt ESell


	Info.Printf("Processing FPMMBuy event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	err := fpmm_abi.Unpack(&eth_evt,"FPMMSell",log.Data)
	if err != nil {
		Error.Printf("Event FPMMBuy decode error: %v",err)
		os.Exit(1)
	}
	eth_evt.Seller = common.BytesToAddress(log.Topics[1][12:])
	eth_evt.OutcomeIndex = log.Topics[2].Big()

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Seller= eth_evt.Seller.String()
	evt.ReturnAmount = eth_evt.ReturnAmount.String()
	evt.FeeAmount = eth_evt.FeeAmount.String()
	evt.OutcomeIdx = eth_evt.OutcomeIndex.Int64()
	evt.TokensSold = eth_evt.OutcomeTokensSold.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("FPMMSell {\n")
	Info.Printf("\tSeller: %v\n",evt.Seller)
	Info.Printf("\tReturnAmount: %v\n",evt.ReturnAmount)
	Info.Printf("\tFeeAmount: %v\n",evt.FeeAmount)
	Info.Printf("\tOutcomeIdx: %v\n",evt.OutcomeIdx)
	Info.Printf("\tTokensSold: %v\n",evt.TokensSold)
	Info.Printf("}\n")

	storage.Insert_fpmm_sell(&evt)
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
func select_event_and_process(log *types.Log,evtlog *EthereumEventLog) {

	Info.Printf("processing event with sig = %v\n",log.Topics[0].String())
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_condition_preparation) {
		proc_condition_preparation(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_condition_resolution) {
		proc_condition_resolution(log,evtlog)
	}
	/*if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_payout_redemption) {
		proc_payout_redemption(&log,&evtlog)
	}*/
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_position_split) {
		proc_position_split(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_position_merge) {
		proc_position_merge(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_uri) {
		proc_uri(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_funding_added) {
		proc_funding_added(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_funding_removed) {
		proc_funding_removed(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_fpmm_buy) {
		proc_fpmm_buy(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_fpmm_sell) {
		proc_fpmm_sell(log,evtlog)
	}
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
		select_event_and_process(&log,&evtlog)
	}
	return nil
}
func fetch_and_process_filtered_events(exit_chan chan bool) {

	block_range := int64(100000 - 1)
	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
		status := storage.Get_polymarkets_processing_status()
		latestBlock, err := eclient.BlockByNumber(context.Background(), nil)
		if err != nil {
			Error.Printf("Error getting last block number from Geth: %v\n",err)
			time.Sleep(5 * time.Second)
			continue
		}
		bnum_high := latestBlock.Number().Int64()
		from_block := status.LastBlockNum + 1
		to_block := from_block + block_range
		if to_block > bnum_high {
			to_block = bnum_high
		}
		for i:=0; i<len(inspected_events); i++ {
			esig := inspected_events[i].Signature
			esig_bytes,_ := hex.DecodeString(esig)
			logs,err := filter_log_query(esig_bytes,from_block,to_block)
			if err != nil {
				Error.Printf("Error getting logs: %v\n",err)
			} else {
				for _,log := range logs {
					if log.Removed {
						continue
					}
					var eth_evt	EthereumEventLog
					eth_evt.BlockNum = int64(log.BlockNumber)
					eth_evt.ContractAddress = log.Address.String()
					select_event_and_process(&log,&eth_evt)
				}
			}
		}
		status.LastBlockNum = to_block
		storage.Update_polymarkets_process_status(&status)
	}
}
