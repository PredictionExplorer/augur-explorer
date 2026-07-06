// Bid events for the CosmicGame ETL: BidPlaced (v1 and v2) and
// FirstBidPlacedInRound (round start).
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

func proc_bid_event_v1(log *types.Log, elog *EthereumEventLog) error {

	var evt CGBidEvent
	var eth_evt CosmicSignatureGameBidPlaced

	Info.Printf("Processing bid v1 id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "BidPlaced", log.Data)
	if err != nil {
		Error.Printf("bid v1 decode error: %v", err)
		os.Exit(1)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.LastBidderAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.EthPrice = eth_evt.PaidEthPrice.String()
	evt.BidType = 0 // ETH
	evt.RandomWalkTokenId = log.Topics[3].Big().Int64()
	evt.ERC20_Value, err = find_cosmic_token_transfer(evt.EvtId, evt.TxId, evt.LastBidderAddr)
	if err != nil {
		return fmt.Errorf("bid v1 (evt id %v): %w", elog.EvtId, err)
	}
	evt.CstPrice = eth_evt.PaidCstPrice.String()
	if evt.RandomWalkTokenId > -1 {
		evt.BidType = 1 // RandomWalk
	} else {
		if evt.CstPrice != "-1" {
			evt.BidType = 2
		} // Cosmic Signature Token (ERC20) bid
	}
	evt.PrizeTime = eth_evt.MainPrizeTime.Int64()
	evt.Message = eth_evt.Message
	evt.BidCstRewardAmount = "-1"
	evt.CstDutchAuctionDuration = "-1"

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("bid v1 {\n")
	Info.Printf("\tLastBidder: %v\n", evt.LastBidderAddr)
	Info.Printf("\tRoundNum: %v\n", evt.RoundNum)
	Info.Printf("\tBidPrice: %v\n", evt.EthPrice)
	Info.Printf("\tCstPrice: %v\n", evt.CstPrice)
	Info.Printf("\tRandomWalkTokenId: %v\n", evt.RandomWalkTokenId)
	Info.Printf("\tPrizeTime: %v\n", evt.PrizeTime)
	Info.Printf("\tMessage: %v\n", evt.Message)
	Info.Printf("}\n")

	storagew.Delete_bid(evt.EvtId)
	storagew.Insert_bid_event(&evt)
	return nil
}
func proc_bid_event_v2(log *types.Log, elog *EthereumEventLog) error {

	var evt CGBidEvent
	var eth_evt CosmicSignatureGameV2BidPlaced

	Info.Printf("Processing bid v2 id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_game_v2_abi.UnpackIntoInterface(&eth_evt, "BidPlaced", log.Data)
	if err != nil {
		Error.Printf("bid v2 decode error: %v", err)
		os.Exit(1)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.LastBidderAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.EthPrice = eth_evt.PaidEthPrice.String()
	evt.BidType = 0 // ETH
	evt.RandomWalkTokenId = log.Topics[3].Big().Int64()
	evt.ERC20_Value, err = find_cosmic_token_transfer(evt.EvtId, evt.TxId, evt.LastBidderAddr)
	if err != nil {
		return fmt.Errorf("bid v2 (evt id %v): %w", elog.EvtId, err)
	}
	evt.CstPrice = eth_evt.PaidCstPrice.String()
	if evt.RandomWalkTokenId > -1 {
		evt.BidType = 1 // RandomWalk
	} else {
		if evt.CstPrice != "-1" {
			evt.BidType = 2
		} // Cosmic Signature Token (ERC20) bid
	}
	evt.PrizeTime = eth_evt.MainPrizeTime.Int64()
	evt.Message = eth_evt.Message
	evt.BidCstRewardAmount = eth_evt.BidCstRewardAmount.String()
	evt.CstDutchAuctionDuration = eth_evt.CstDutchAuctionDuration.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("bid v2 {\n")
	Info.Printf("\tLastBidder: %v\n", evt.LastBidderAddr)
	Info.Printf("\tRoundNum: %v\n", evt.RoundNum)
	Info.Printf("\tBidPrice: %v\n", evt.EthPrice)
	Info.Printf("\tCstPrice: %v\n", evt.CstPrice)
	Info.Printf("\tRandomWalkTokenId: %v\n", evt.RandomWalkTokenId)
	Info.Printf("\tPrizeTime: %v\n", evt.PrizeTime)
	Info.Printf("\tBidCstRewardAmount: %v\n", evt.BidCstRewardAmount)
	Info.Printf("\tCstDutchAuctionDuration: %v\n", evt.CstDutchAuctionDuration)
	Info.Printf("\tMessage: %v\n", evt.Message)
	Info.Printf("}\n")

	storagew.Delete_bid(evt.EvtId)
	storagew.Insert_bid_event(&evt)
	return nil
}
func proc_round_started_event(log *types.Log, elog *EthereumEventLog) {

	var evt CGRoundStarted
	var eth_evt CosmicSignatureGameFirstBidPlacedInRound

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing FirstBidPlacedInRound  event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "FirstBidPlacedInRound", log.Data)
	if err != nil {
		Error.Printf("Event FirstBidPlacedInRound decode error: %v", err)
		os.Exit(1)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.StartTimestamp = eth_evt.BlockTimeStamp.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("FirstBidPlacedInRound{\n")
	Info.Printf("\tRoundNum: %v\n", evt.RoundNum)
	Info.Printf("\tStart TS: %v\n", evt.StartTimestamp)
	Info.Printf("}\n")

	storagew.Delete_round_started_event(evt.EvtId)
	storagew.Insert_round_started_event(&evt)
}
