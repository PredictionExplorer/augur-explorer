// Admin and parameter-change events: percentages, durations, divisors, address
// changes, ownership transfers, proxy upgrades and related system-management events.
package main

import (
	"os"
	"bytes"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)
func store_cst_reward_for_bidding_changed(log *types.Log, elog *EthereumEventLog, newReward string, label string) {
	var evt CGCstRewardForBiddingChanged

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewReward = newReward

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("%s {\n", label)
	Info.Printf("\tNewReward: %v\n", evt.NewReward)
	Info.Printf("}\n")

	storagew.Delete_erc20_token_reward_changed_event(evt.EvtId)
	storagew.Insert_erc20_token_reward_changed_event(&evt)
}
func store_cst_dutch_auction_length_changed(log *types.Log, elog *EthereumEventLog, newValue string, label string) {
	var evt CGCstDutchAuctionDurationDivisorChanged

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewValue = newValue

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("%s {\n", label)
	Info.Printf("\tNewAuctionLength: %v\n", evt.NewValue)
	Info.Printf("}\n")

	storagew.Delete_round_start_cst_auction_length_changed_event(evt.EvtId)
	storagew.Insert_round_start_cst_auction_length_changed_event(&evt)
}
func proc_charity_address_changed_unified(log *types.Log,elog *EthereumEventLog) {
	// This event is emitted by TWO different contracts with the same signature
	// Distinguish by contract address
	
	if bytes.Equal(log.Address.Bytes(),charity_wallet_addr.Bytes()) {
		// CharityWallet contract: Sets who receives charity funds
		var evt CGCharityUpdatedEvent
		var eth_evt CharityWalletCharityAddressChanged
		
		Info.Printf("Processing CharityReceiverChanged (from CharityWallet) event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
		
		err := charity_wallet_abi.UnpackIntoInterface(&eth_evt,"CharityAddressChanged",log.Data)
		if err != nil {
			Error.Printf("Event CharityAddressChanged (CharityWallet) decode error: %v",err)
			os.Exit(1)
		}
		
		evt.EvtId=elog.EvtId
		evt.BlockNum = elog.BlockNum
		evt.TxId = elog.TxId
		evt.ContractAddr = log.Address.String()
		evt.TimeStamp = elog.TimeStamp
		evt.NewCharityAddr = common.BytesToAddress(log.Topics[1][12:]).String()
		
		Info.Printf("CharityReceiverChanged: %v\n",evt.NewCharityAddr)
		
		storagew.Delete_charity_updated(evt.EvtId)
		storagew.Insert_charity_updated_event(&evt)
		
	} else if bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		// CosmicGame contract: Sets which CharityWallet contract to use
		var evt CGCharityAddressChanged
		var eth_evt CosmicSignatureGameCharityAddressChanged
		
		Info.Printf("Processing CharityWalletChanged (from CosmicGame) event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
		
		err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CharityAddressChanged",log.Data)
		if err != nil {
			Error.Printf("Event CharityAddressChanged (CosmicGame) decode error: %v",err)
			os.Exit(1)
		}
		
		evt.EvtId=elog.EvtId
		evt.BlockNum = elog.BlockNum
		evt.TxId = elog.TxId
		evt.Contract = log.Address.String()
		evt.TimeStamp = elog.TimeStamp
		evt.NewCharity = common.BytesToAddress(log.Topics[1][12:]).String()
		
		Info.Printf("CharityWalletChanged: %v\n",evt.NewCharity)
		
		storagew.Delete_cosmic_game_charity_address_changed_event(evt.EvtId)
		storagew.Insert_cosmic_game_charity_address_changed_event(&evt)
	} else {
		Info.Printf("CharityAddressChanged from unknown contract %v, skipping\n",log.Address.String())
		return
	}
}
func proc_charity_percentage_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCharityPercentageChanged
	var eth_evt CosmicSignatureGameCharityEthDonationAmountPercentageChanged 

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing CharityPercentageChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CharityEthDonationAmountPercentageChanged",log.Data)
	if err != nil {
		Error.Printf("Event CharityPercentageChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCharityPercentage= eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CharityPercentageChanged {\n")
	Info.Printf("\tNewCharityPercentage: %v\n",evt.NewCharityPercentage)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_charity_percentage_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_charity_percentage_changed_event(&evt)
}
func proc_prize_percentage_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGPrizePercentageChanged
	var eth_evt CosmicSignatureGameMainEthPrizeAmountPercentageChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing PrizePercentageChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"MainEthPrizeAmountPercentageChanged",log.Data)
	if err != nil {
		Error.Printf("Event PrizePercentageChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewPrizePercentage= eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PrizePercentageChanged {\n")
	Info.Printf("\tNewPrizePercentage: %v\n",evt.NewPrizePercentage)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_prize_percentage_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_prize_percentage_changed_event(&evt)
}
func proc_raffle_percentage_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRafflePercentageChanged
	var eth_evt CosmicSignatureGameRaffleTotalEthPrizeAmountForBiddersPercentageChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing RafflePercentageChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"RaffleTotalEthPrizeAmountForBiddersPercentageChanged",log.Data)
	if err != nil {
		Error.Printf("Event RaffleTotalEthPrizeAmountForBiddersPercentageChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewRafflePercentage= eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RaffleTotalEthPrizeAmountForBiddersPercentageChanged{\n")
	Info.Printf("\tNewRafflePercentage: %v\n",evt.NewRafflePercentage)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_raffle_percentage_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_raffle_percentage_changed_event(&evt)
}
func proc_staking_percentage_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGStakingPercentageChanged
	var eth_evt CosmicSignatureGameCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing StakingPercentageChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged",log.Data)
	if err != nil {
		Error.Printf("Event CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewStakingPercentage= eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged{\n")
	Info.Printf("\tNewStakingPercentage: %v\n",evt.NewStakingPercentage)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_staking_percentage_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_staking_percentage_changed_event(&evt)
}
func proc_chrono_percentage_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGChronoPercentageChanged
	var eth_evt ISystemEventsChronoWarriorEthPrizeAmountPercentageChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing ChronoWarriorEthPrizePercentageChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"ChronoWarriorEthPrizeAmountPercentageChanged",log.Data)
	if err != nil {
		Error.Printf("Event ChronoWarriorEthPrizePercentageChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewChronoPercentage= eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("ChronoWarriorEthPrizePercentageChanged{\n")
	Info.Printf("\tNewPercentage: %v\n",evt.NewChronoPercentage)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_chrono_percentage_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_chrono_percentage_changed_event(&evt)
}
func proc_num_raffle_eth_winners_bidding_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNumRaffleETHWinnersBiddingChanged 
	var eth_evt CosmicSignatureGameNumRaffleEthPrizesForBiddersChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing NumRaffleETHWinnersBiddingChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"NumRaffleEthPrizesForBiddersChanged",log.Data)
	if err != nil {
		Error.Printf("Event NumRaffleEthPrizesForBiddersChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewNumRaffleETHWinnersBidding = eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NumRaffleEthPrizesForBiddersChanged{\n")
	Info.Printf("\tNewNumRaffleETHWinnersBidding: %v\n",evt.NewNumRaffleETHWinnersBidding)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_num_raffle_eth_winners_bidding_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_num_raffle_eth_winners_bidding_changed_event(&evt)
}
func proc_num_raffle_nft_winners_bidding_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNumRaffleNFTWinnersBiddingChanged 
	var eth_evt ISystemManagementNumRaffleCosmicSignatureNftsForBiddersChanged 

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing NumRaffleNftWinnersBiddingChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"NumRaffleCosmicSignatureNftsForBiddersChanged",log.Data)
	if err != nil {
		Error.Printf("Event NumRaffleNftWinnersBiddingChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewNumRaffleNFTWinnersBidding  = eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NumRaffleNftWinnersBiddingChanged{\n")
	Info.Printf("\tNewNumRaffleNFTWinnersBidding: %v\n",evt.NewNumRaffleNFTWinnersBidding)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_num_raffle_nft_winners_bidding_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_num_raffle_nft_winners_bidding_changed_event(&evt)
}
func proc_num_raffle_nft_winners_staking_rwalk_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNumRaffleNFTWinnersStakingRWalkChanged
	var eth_evt ISystemManagementNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing NumRaffleNFTWinnersStakingRWalkChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged",log.Data)
	if err != nil {
		Error.Printf("Event NumRaffleNFTWinnersStakingRWalkChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewNumRaffleNFTWinnersStakingRWalk = eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NumRaffleNFTWinnersStakingRWalkChanged{\n")
	Info.Printf("\tNewNumRaffleNFTWinnersStakingRWalk: %v\n",evt.NewNumRaffleNFTWinnersStakingRWalk)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_num_raffle_nft_winners_staking_rwalk_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_num_raffle_nft_winners_staking_rwalk_changed_event(&evt)
}
func proc_random_walk_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRandomWalkAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing RandomWalkAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewRandomWalk = common.BytesToAddress(log.Topics[1][12:]).String()
	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RandomWalkAddressChanged{\n")
	Info.Printf("\tNewRandomWalk: %v\n",evt.NewRandomWalk)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_random_walk_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_random_walk_address_changed_event(&evt)
}
func proc_raffle_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGPrizeWalletAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing EthPrizesWalletAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewPrizeWallet = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("EthPrizesWalletAddressChanged{\n")
	Info.Printf("\tNewEthPrizesWallet: %v\n",evt.NewPrizeWallet)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_prize_wallet_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_prize_wallet_address_changed_event(&evt)
}
func proc_staking_wallet_cst_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGStakingWalletCSTAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing StakingWalletCSTAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewStakingWalletCST = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("StakingWalletCSTAddressChanged{\n")
	Info.Printf("\tNewStakingWalletCST: %v\n",evt.NewStakingWalletCST)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_staking_wallet_cst_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_staking_wallet_cst_address_changed_event(&evt)
}
func proc_staking_wallet_rwalk_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGStakingWalletRWalkAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing StakingWalletRWalkAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewStakingWalletRWalk = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("StakingWalletRWalkAddressChanged{\n")
	Info.Printf("\tNewStakingWalletRWalk: %v\n",evt.NewStakingWalletRWalk)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_staking_wallet_rwalk_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_staking_wallet_rwalk_address_changed_event(&evt)
}
func proc_marketing_wallet_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGMarketingWalletAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing MarketingWalletAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewMarketingWallet = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("MarketingWalletAddressChanged{\n")
	Info.Printf("\tNewMarketingWallet: %v\n",evt.NewMarketingWallet)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_marketing_wallet_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_marketing_wallet_address_changed_event(&evt)
}
func proc_treasurer_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGTreasurerAddressChanged

	if !bytes.Equal(log.Address.Bytes(),marketing_wallet_addr.Bytes()) {
		return
	}
	Info.Printf("Processing TreasurerAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewTreasurer = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TreasurerAddressChanged{\n")
	Info.Printf("\tNewTreasurer: %v\n",evt.NewTreasurer)
	Info.Printf("}\n")

	storagew.Delete_treasurer_address_changed_event(evt.EvtId)
    storagew.Insert_treasurer_address_changed_event(&evt)
}
func proc_cosmic_token_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCosmicTokenAddressChanged
	var eth_evt CosmicSignatureGameCosmicSignatureTokenAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing CosmicSignatureTokenAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CosmicSignatureTokenAddressChanged",log.Data)
	if err != nil {
		Error.Printf("Event CosmicSignatureTokenAddressChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCosmicToken = eth_evt.NewValue.String()
	evt.NewCosmicToken = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CosmicSignatureTokenAddressChanged{\n")
	Info.Printf("\tNewCosmicToken: %v\n",evt.NewCosmicToken)
	Info.Printf("}\n")

	storagew.Delete_cosmic_token_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_token_address_changed_event(&evt)
}
func proc_cosmic_signature_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCosmicSignatureAddressChanged
	var eth_evt CosmicSignatureGameCosmicSignatureNftAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing CosmicSignatureAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCosmicSignature = common.BytesToAddress(log.Topics[1][12:]).String()
	_ = eth_evt.NewValue.String() // blank-assigned to keep go vet clean (pre-existing dead call)

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CosmicSignatureAddressChanged{\n")
	Info.Printf("\tNewCosmicSignatureWallet: %v\n",evt.NewCosmicSignature)
	Info.Printf("}\n")

	storagew.Delete_cosmic_signature_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_signature_address_changed_event(&evt)
}
func proc_proxy_upgraded_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGUpgraded
	var eth_evt CosmicSignatureGameUpgraded

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event Upgraded doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing Upgraded event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"Upgraded",log.Data)
	if err != nil {
		Error.Printf("Event Upgraded decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Implementation = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("(Proxy) Upgraded{\n")
	Info.Printf("\tImplementation: %v\n",evt.Implementation)
	Info.Printf("}\n")

	storagew.Delete_upgraded_event(evt.EvtId)
    storagew.Insert_upgraded_event(&evt)
}
func proc_admin_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGAdminChanged
	var eth_evt IERC1967AdminChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event AdminChanged doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing AdminChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"AdminChanged",log.Data)
	if err != nil {
		Error.Printf("Event AdminChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.OldAdmin = eth_evt.PreviousAdmin.String()
	evt.NewAdmin = eth_evt.NewAdmin.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("AdminChanged{\n")
	Info.Printf("\tOldAdmin: %v\n",evt.OldAdmin)
	Info.Printf("\tNewAdmin: %v\n",evt.NewAdmin)
	Info.Printf("}\n")

	storagew.Delete_admin_changed_event(evt.EvtId)
    storagew.Insert_admin_changed_event(&evt)
}
func proc_time_increase_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGTimeIncreaseChanged
	var eth_evt CosmicSignatureGameMainPrizeTimeIncrementInMicroSecondsChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing TimeIncreaseChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"TimeIncreaseChanged",log.Data)
	if err != nil {
		Error.Printf("Event TimeIncreaseChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewTimeIncrease = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TimeIncreaseChanged{\n")
	Info.Printf("\tNewTimeIncrease: %v\n",evt.NewTimeIncrease)
	Info.Printf("}\n")

	storagew.Delete_time_increase_changed_event(evt.EvtId)
    storagew.Insert_time_increase_changed_event(&evt)
}
func proc_timeout_claimprize_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGTimeoutClaimPrizeChanged
	var eth_evt CosmicSignatureGameTimeoutDurationToClaimMainPrizeChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing TimeoutClaimPrizeChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"TimeoutDurationToClaimMainPrizeChanged",log.Data)
	if err != nil {
		Error.Printf("Event TimeoutClaimPrizeChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewTimeout = eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TimeoutDurationToClaimMainPrizeChanged{\n")
	Info.Printf("\tNewTimeout: %v\n",evt.NewTimeout)
	Info.Printf("}\n")

	storagew.Delete_timeout_claimprize_changed_event(evt.EvtId)
    storagew.Insert_timeout_claimprize_changed_event(&evt)
}
func proc_timeout_duration_to_withdraw_prize_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGTimeoutToWithdrawPrizeChanged
	var eth_evt IPrizesWalletTimeoutDurationToWithdrawPrizesChanged

	if !bytes.Equal(log.Address.Bytes(),prizes_wallet_addr.Bytes()) {
		return
	}
	Info.Printf("Processing TimeoutDurationToWithdrawPrizesChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt,"TimeoutDurationToWithdrawPrizesChanged",log.Data)
	if err != nil {
		Error.Printf("Event TimeoutDurationToWithdrawPrizesChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewTimeout = eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TimeoutDurationToWithdrawPrizesChanged {\n")
	Info.Printf("\tNewTimeout: %v\n",evt.NewTimeout)
	Info.Printf("}\n")

	storagew.Delete_timeout_to_withdraw_prizes_changed_event(evt.EvtId)
    storagew.Insert_timeout_to_withdraw_prizes_changed_event(&evt)
}
func proc_price_increase_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGPriceIncreaseChanged
	var eth_evt CosmicSignatureGameEthBidPriceIncreaseDivisorChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing PriceIncreaseChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"EthBidPriceIncreaseDivisorChanged",log.Data)
	if err != nil {
		Error.Printf("Event EthBidPriceIncreaseDivisorChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewPriceIncrease = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("EthBidPriceIncreaseDivisorChanged{\n")
	Info.Printf("\tNewPriceIncreasse: %v\n",evt.NewPriceIncrease)
	Info.Printf("}\n")

	storagew.Delete_price_increase_changed_event(evt.EvtId)
    storagew.Insert_price_increase_changed_event(&evt)
}
func proc_mainprize_microsecond_increase_changed(log *types.Log,elog *EthereumEventLog) {

	var evt CGMainPrizeMicroSecondsIncreaseChanged
	var eth_evt CosmicSignatureGameMainPrizeTimeIncrementInMicroSecondsChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing MainPrizeTimeIncrementInMicroSecondsChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"MainPrizeTimeIncrementInMicroSecondsChanged",log.Data)
	if err != nil {
		Error.Printf("Event MainPrizeTimeIncrementInMicroSecondsChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewMicroseconds= eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("MainPrizeTimeIncrementInMicroSecondsChanged{\n")
	Info.Printf("\tNewMicroseconds: %v\n",evt.NewMicroseconds)
	Info.Printf("}\n")

	storagew.Delete_mainprize_microseconds_increase_changed_event(evt.EvtId)
    storagew.Insert_mainprize_microseconds_increase_changed_event(&evt)
}
func proc_initial_seconds_until_prize_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGInitialSecondsUntilPrizeChanged
	var eth_evt CosmicSignatureGameInitialDurationUntilMainPrizeDivisorChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing InitialDurationUntilMainPrizeDivisorChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"InitialDurationUntilMainPrizeDivisorChanged",log.Data)
	if err != nil {
		Error.Printf("Event InitialDurationUntilMainPrizeDivisorChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewInitialSecondsUntilPrize = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("InitialDurationUntilMainPrizeDivisorChanged{\n")
	Info.Printf("\tNewInitialSecondsUntilPrize: %v\n",evt.NewInitialSecondsUntilPrize)
	Info.Printf("}\n")

	storagew.Delete_initial_seconds_until_prize_changed_event(evt.EvtId)
    storagew.Insert_initial_seconds_until_prize_changed_event(&evt)
}
func proc_activation_time_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGActivationTimeChanged
	var eth_evt BiddingBaseRoundActivationTimeChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing RoundActivationTimeChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"RoundActivationTimeChanged",log.Data)
	if err != nil {
		Error.Printf("Event RoundActivationTimeChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewActivationTime = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("ActivationTimeChanged {\n")
	Info.Printf("\tNewActivationTime: %v\n",evt.NewActivationTime)
	Info.Printf("}\n")

	storagew.Delete_activation_time_changed_event(evt.EvtId)
    storagew.Insert_activation_time_changed_event(&evt)
}
func proc_cst_dutch_auction_duration_divisor_changed_event(log *types.Log,elog *EthereumEventLog) {

	var eth_evt CosmicSignatureGameCstDutchAuctionDurationDivisorChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing CstDutchAuctionDurationDivisorChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CstDutchAuctionDurationDivisorChanged",log.Data)
	if err != nil {
		Error.Printf("Event CstDutchAuctionDurationDivisorChanged decode error: %v",err)
		os.Exit(1)
	}

	store_cst_dutch_auction_length_changed(log, elog, eth_evt.NewValue.String(), "CstDutchAuctionDurationDivisorChanged")
}
func proc_cst_dutch_auction_duration_changed_event(log *types.Log,elog *EthereumEventLog) {

	var eth_evt CosmicSignatureGameV2CstDutchAuctionDurationChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing CstDutchAuctionDurationChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_v2_abi.UnpackIntoInterface(&eth_evt,"CstDutchAuctionDurationChanged",log.Data)
	if err != nil {
		Error.Printf("Event CstDutchAuctionDurationChanged decode error: %v",err)
		os.Exit(1)
	}

	store_cst_dutch_auction_length_changed(log, elog, eth_evt.NewValue.String(), "CstDutchAuctionDurationChanged")
}
func proc_cst_dutch_auction_duration_change_divisor_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCstDutchAuctionDurationChangeDivisorChanged
	var eth_evt CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing CstDutchAuctionDurationChangeDivisorChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_v2_abi.UnpackIntoInterface(&eth_evt,"CstDutchAuctionDurationChangeDivisorChanged",log.Data)
	if err != nil {
		Error.Printf("Event CstDutchAuctionDurationChangeDivisorChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewValue = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CstDutchAuctionDurationChangeDivisorChanged {\n")
	Info.Printf("\tNewDivisor: %v\n",evt.NewValue)
	Info.Printf("}\n")

	storagew.Delete_cst_dutch_auction_duration_change_divisor_changed_event(evt.EvtId)
	storagew.Insert_cst_dutch_auction_duration_change_divisor_changed_event(&evt)
}
func proc_eth_dutch_auction_duration_divisor_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGEthDutchAuctionDurationDivisorChanged
	var eth_evt CosmicSignatureGameEthDutchAuctionDurationDivisorChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing EthDutchAuctionDurationDivisorChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"EthDutchAuctionDurationDivisorChanged",log.Data)
	if err != nil {
		Error.Printf("Event EthDutchAuctionDurationDivisorChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewValue  = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("EthDutchAuctionDurationDivisorChanged {\n")
	Info.Printf("\tNewDivisor: %v\n",evt.NewValue)
	Info.Printf("}\n")

	storagew.Delete_eth_dutch_auction_duration_divisor_changed_event(evt.EvtId)
    storagew.Insert_eth_auction_duration_divisor_changed_event(&evt)
}
func proc_eth_dutch_auction_ending_bid_price_divisor_changed__event(log *types.Log,elog *EthereumEventLog) {

	var evt CGEthDutchAuctionEndingBidPriceDivisorChanged 
	var eth_evt CosmicSignatureGameEthDutchAuctionEndingBidPriceDivisorChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing EthDutchAuctionEndingBidPriceDivisorChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"EthDutchAuctionEndingBidPriceDivisorChanged",log.Data)
	if err != nil {
		Error.Printf("Event EthDutchAuctionEndingBidPriceDivisorChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewValue  = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("EthDutchAuctionEndingBidPriceDivisorChanged{\n")
	Info.Printf("\tNewDivisor: %v\n",evt.NewValue)
	Info.Printf("}\n")

	storagew.Delete_eth_dutch_auction_ending_bidprice_divisor_changed_event(evt.EvtId)
    storagew.Insert_eth_dutch_auction_ending_bidprice_divisor_changed_event(&evt)
}
func proc_marketing_reward_changed(log *types.Log,elog *EthereumEventLog) {

	var evt CGMarketingRewardChanged
	var eth_evt CosmicSignatureGameMarketingWalletCstContributionAmountChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing MarketingWalletCstContributionAmountChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"MarketingWalletCstContributionAmountChanged",log.Data)
	if err != nil {
		Error.Printf("Event MarketingWalletCstContributionAmountChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewReward= eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("MarketingWalletCstContributionAmountChanged{\n")
	Info.Printf("\tNewReward: %v\n",evt.NewReward)
	Info.Printf("}\n")

	storagew.Delete_marketing_reward_changed_event(evt.EvtId)
    storagew.Insert_marketing_reward_changed_event(&evt)
}
func proc_erc20_token_reward_changed_event(log *types.Log,elog *EthereumEventLog) {

	var eth_evt CosmicSignatureGameCstRewardAmountForBiddingChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing CstRewardAmountForBiddingChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CstRewardAmountForBiddingChanged",log.Data)
	if err != nil {
		Error.Printf("Event CstRewardAmountForBiddingChanged decode error: %v",err)
		os.Exit(1)
	}

	store_cst_reward_for_bidding_changed(log, elog, eth_evt.NewValue.String(), "CstRewardAmountForBiddingChanged")
}
func proc_bid_cst_reward_amount_changed_event(log *types.Log,elog *EthereumEventLog) {

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing BidCstRewardAmountChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	newValue, err := admin_uint256_from_log_data(log.Data)
	if err != nil {
		Error.Printf("Event BidCstRewardAmountChanged decode error: %v",err)
		os.Exit(1)
	}

	store_cst_reward_for_bidding_changed(log, elog, newValue.String(), "BidCstRewardAmountChanged")
}
func proc_bid_cst_reward_amount_multiplier_changed_event(log *types.Log,elog *EthereumEventLog) {

	var eth_evt CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing BidCstRewardAmountMultiplierChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_v2_abi.UnpackIntoInterface(&eth_evt,"BidCstRewardAmountMultiplierChanged",log.Data)
	if err != nil {
		Error.Printf("Event BidCstRewardAmountMultiplierChanged decode error: %v",err)
		os.Exit(1)
	}

	store_cst_reward_for_bidding_changed(log, elog, eth_evt.NewValue.String(), "BidCstRewardAmountMultiplierChanged")
}
func proc_static_cst_reward_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGStaticCstReward
	var eth_evt BiddingCstPrizeAmountChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing CstPrizeAmountChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CstPrizeAmountChanged",log.Data)
	if err != nil {
		Error.Printf("Event CstPrizeAmountChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewReward= eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CstPrizeAmountChanged{\n")
	Info.Printf("\tNewReward: %v\n",evt.NewReward)
	Info.Printf("}\n")

	storagew.Delete_static_cst_reward_changed_event(evt.EvtId)
    storagew.Insert_static_cst_reward_changed_event(&evt)
}
func proc_max_msg_length_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGMaxMessageLengthChanged 
	var eth_evt CosmicSignatureGameBidMessageLengthMaxLimitChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing BidMessageLengthMaxLimitChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"BidMessageLengthMaxLimitChanged",log.Data)
	if err != nil {
		Error.Printf("Event BidMessageLengthMaxLimitChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewMessageLength = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("BidMessageLengthMaxLimitChanged{\n")
	Info.Printf("\tNewMessageLength: %v\n",evt.NewMessageLength)
	Info.Printf("}\n")

	storagew.Delete_max_message_length_changed_event(evt.EvtId)
    storagew.Insert_max_message_length_changed_event(&evt)
}
func proc_token_generation_script_url_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGTokenGenerationScriptURL
	var eth_evt ICosmicSignatureNftNftGenerationScriptUriChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		return
	}
	Info.Printf("Processing TokenGenerationScriptURLEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt,"NftGenerationScriptUriChanged",log.Data)
	if err != nil {
		Error.Printf("Event TokenGenerationScriptURLEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewURL = eth_evt.NewValue

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TokenGenerationScriptURLEvent{\n")
	Info.Printf("\tNewURL: %v\n",evt.NewURL)
	Info.Printf("}\n")

	storagew.Delete_max_message_length_changed_event(evt.EvtId)
    storagew.Insert_token_generation_script_url_event(&evt)
}
func proc_base_uri_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGBaseURIEvent
	var eth_evt CosmicSignatureNftNftBaseUriChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		return
	}
	Info.Printf("Processing BaseURIEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt,"NftBaseUriChanged",log.Data)
	if err != nil {
		Error.Printf("Event BaseURIEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewURI = eth_evt.NewValue

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("BaseURIEvent{\n")
	Info.Printf("\tNewURI: %v\n",evt.NewURI)
	Info.Printf("}\n")

	storagew.Delete_base_uri_event(evt.EvtId)
    storagew.Insert_base_uri_event(&evt)
}
func proc_ownership_transferred_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGOwnershipTransferred
	var eth_evt CosmicSignatureGameOwnershipTransferred

	contract_code := int64(0);
	if bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		contract_code = 1
	}
	if bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		contract_code = 2
	}
	if bytes.Equal(log.Address.Bytes(),cosmic_token_addr.Bytes()) {
		contract_code = 3
	}
	if bytes.Equal(log.Address.Bytes(),charity_wallet_addr.Bytes()) {
		contract_code = 4 
	}
	if bytes.Equal(log.Address.Bytes(),prizes_wallet_addr.Bytes()) {
		contract_code = 5
	}
	if bytes.Equal(log.Address.Bytes(),staking_wallet_cst_addr.Bytes()) {
		contract_code = 6
	}
	if bytes.Equal(log.Address.Bytes(),staking_wallet_rwalk_addr.Bytes()) {
		contract_code = 7
	}
	if bytes.Equal(log.Address.Bytes(),marketing_wallet_addr.Bytes()) {
		contract_code = 8
	}
	if bytes.Equal(log.Address.Bytes(),cosmic_dao_addr.Bytes()) {
		contract_code = 9 
	}
	if contract_code == 0 {
		return
	}
	Info.Printf("Processing OwnershipTransferred event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt,"OwnershipTransferred",log.Data)
	if err != nil {
		Error.Printf("Event OwnershipTransferred decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.PrevOwner= common.BytesToAddress(log.Topics[1][12:]).String()
	evt.NewOwner= common.BytesToAddress(log.Topics[2][12:]).String()
	evt.ContractCode = contract_code

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("OwnershipTransferred{\n")
	Info.Printf("\tPrevOwner: %v\n",evt.PrevOwner)
	Info.Printf("\tNewOwner: %v\n",evt.NewOwner)
	Info.Printf("}\n")

	storagew.Delete_ownership_transferred_event(evt.EvtId)
    storagew.Insert_ownership_transferred_event(&evt)
}
// isCgInitializedContract reports whether log.Address is a CosmicGame platform
// contract that may emit OpenZeppelin Initializable:Initialized.
func isCgInitializedContract(addr common.Address) bool {
	switch addr {
	case cosmic_game_addr, cosmic_signature_addr, cosmic_token_addr, cosmic_dao_addr,
		charity_wallet_addr, prizes_wallet_addr, staking_wallet_cst_addr,
		staking_wallet_rwalk_addr, marketing_wallet_addr, implementation_addr:
		return true
	default:
		return false
	}
}

func proc_initialized_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGInitialized
	var eth_evt CosmicSignatureGameInitialized 

	if !isCgInitializedContract(log.Address) {
		return
	}
	Info.Printf("Processing Initialized event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"Initialized",log.Data)
	if err != nil {
		Error.Printf("Event Initialized decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Version = int64(eth_evt.Version)

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("Initialized {\n")
	Info.Printf("\tVersion: %v\n",evt.Version)
	Info.Printf("}\n")

	storagew.Delete_initialized_event(evt.EvtId)
    storagew.Insert_initialized_event(&evt)
}
func proc_starting_bid_price_cst_min_limit_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCstMinLimit
	var eth_evt BiddingCstDutchAuctionBeginningBidPriceMinLimitChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing CstDutchAuctionBeginningBidPriceMinLimitChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CstDutchAuctionBeginningBidPriceMinLimitChanged",log.Data)
	if err != nil {
		Error.Printf("Event CstDutchAuctionBeginningBidPriceMinLimitChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.CstMinLimit = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CstDutchAuctionBeginningBidPriceMinLimitChanged {\n")
	Info.Printf("\tNewStartingBidPriceCSTMinLimit: %v\n",evt.CstMinLimit)
	Info.Printf("}\n")

	storagew.Delete_cst_min_limit_event(evt.EvtId)
    storagew.Insert_cst_min_limit_event(&evt)
}
func proc_delay_duration_before_next_round_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNextRoundDelayDuration
	var eth_evt CosmicSignatureGameDelayDurationBeforeRoundActivationChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing DelayDurationBeforeRoundActivationChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"DelayDurationBeforeRoundActivationChanged",log.Data)
	if err != nil {
		Error.Printf("Event DelayDurationBeforeRoundActivationChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewValue= eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DelayDurationBeforeRoundActivationChanged{\n")
	Info.Printf("\tNewValue: %v\n",evt.NewValue)
	Info.Printf("}\n")

    storagew.Delete_delay_duration_before_next_round_changed_event(evt.EvtId)
	storagew.Insert_delay_duration_before_next_round_changed_event(&evt)
}
