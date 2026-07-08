// Admin and parameter-change events: percentages, durations, divisors, address
// changes, ownership transfers, proxy upgrades and related system-management events.
package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

// store_cst_reward_for_bidding_changed persists a CST-bid-reward change; three
// event types (CstRewardAmountForBiddingChanged, BidCstRewardAmountChanged,
// BidCstRewardAmountMultiplierChanged) share the cg_adm_erc20_reward table.
func store_cst_reward_for_bidding_changed(ctx context.Context, log *types.Log, elog *EthereumEventLog, newReward string, label string) error {
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

	if err := cgRepo.DeleteCstRewardForBiddingChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertCstRewardForBiddingChange(ctx, &evt)
}

// store_cst_dutch_auction_length_changed persists a CST auction-length
// change; the V1 divisor and V2 duration events share the cg_adm_cst_auclen
// table.
func store_cst_dutch_auction_length_changed(ctx context.Context, log *types.Log, elog *EthereumEventLog, newValue string, label string) error {
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

	if err := cgRepo.DeleteCstAuctionLengthChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertCstAuctionLengthChange(ctx, &evt)
}
func proc_charity_address_changed_unified(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {
	// This event is emitted by TWO different contracts with the same signature
	// Distinguish by contract address

	if bytes.Equal(log.Address.Bytes(), charity_wallet_addr.Bytes()) {
		// CharityWallet contract: Sets who receives charity funds
		var evt CGCharityUpdatedEvent
		var eth_evt CharityWalletCharityAddressChanged

		Info.Printf("Processing CharityReceiverChanged (from CharityWallet) event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

		err := charity_wallet_abi.UnpackIntoInterface(&eth_evt, "CharityAddressChanged", log.Data)
		if err != nil {
			return fmt.Errorf("CharityAddressChanged/CharityWallet (evt id %v): decode: %w", elog.EvtId, err)
		}

		evt.EvtId = elog.EvtId
		evt.BlockNum = elog.BlockNum
		evt.TxId = elog.TxId
		evt.ContractAddr = log.Address.String()
		evt.TimeStamp = elog.TimeStamp
		evt.NewCharityAddr = common.BytesToAddress(log.Topics[1][12:]).String()

		Info.Printf("CharityReceiverChanged: %v\n", evt.NewCharityAddr)

		if err := cgRepo.DeleteCharityReceiverChange(ctx, evt.EvtId); err != nil {
			return err
		}
		return cgRepo.InsertCharityReceiverChange(ctx, &evt)

	} else if bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		// CosmicGame contract: Sets which CharityWallet contract to use
		var evt CGCharityAddressChanged
		var eth_evt CosmicSignatureGameCharityAddressChanged

		Info.Printf("Processing CharityWalletChanged (from CosmicGame) event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

		err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "CharityAddressChanged", log.Data)
		if err != nil {
			return fmt.Errorf("CharityAddressChanged/CosmicGame (evt id %v): decode: %w", elog.EvtId, err)
		}

		evt.EvtId = elog.EvtId
		evt.BlockNum = elog.BlockNum
		evt.TxId = elog.TxId
		evt.Contract = log.Address.String()
		evt.TimeStamp = elog.TimeStamp
		evt.NewCharity = common.BytesToAddress(log.Topics[1][12:]).String()

		Info.Printf("CharityWalletChanged: %v\n", evt.NewCharity)

		if err := cgRepo.DeleteCharityWalletAddressChange(ctx, evt.EvtId); err != nil {
			return err
		}
		return cgRepo.InsertCharityWalletAddressChange(ctx, &evt)
	}
	Info.Printf("CharityAddressChanged from unknown contract %v, skipping\n", log.Address.String())
	return nil
}
func proc_charity_percentage_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGCharityPercentageChanged
	var eth_evt CosmicSignatureGameCharityEthDonationAmountPercentageChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing CharityPercentageChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "CharityEthDonationAmountPercentageChanged", log.Data)
	if err != nil {
		return fmt.Errorf("CharityPercentageChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCharityPercentage = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("CharityPercentageChanged {\n")
	Info.Printf("\tNewCharityPercentage: %v\n", evt.NewCharityPercentage)
	Info.Printf("}\n")

	if err := cgRepo.DeleteCharityPercentageChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertCharityPercentageChange(ctx, &evt)
}
func proc_prize_percentage_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGPrizePercentageChanged
	var eth_evt CosmicSignatureGameMainEthPrizeAmountPercentageChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing PrizePercentageChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "MainEthPrizeAmountPercentageChanged", log.Data)
	if err != nil {
		return fmt.Errorf("PrizePercentageChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewPrizePercentage = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("PrizePercentageChanged {\n")
	Info.Printf("\tNewPrizePercentage: %v\n", evt.NewPrizePercentage)
	Info.Printf("}\n")

	if err := cgRepo.DeletePrizePercentageChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertPrizePercentageChange(ctx, &evt)
}
func proc_raffle_percentage_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGRafflePercentageChanged
	var eth_evt CosmicSignatureGameRaffleTotalEthPrizeAmountForBiddersPercentageChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing RafflePercentageChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged", log.Data)
	if err != nil {
		return fmt.Errorf("RafflePercentageChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewRafflePercentage = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("RaffleTotalEthPrizeAmountForBiddersPercentageChanged{\n")
	Info.Printf("\tNewRafflePercentage: %v\n", evt.NewRafflePercentage)
	Info.Printf("}\n")

	if err := cgRepo.DeleteRafflePercentageChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertRafflePercentageChange(ctx, &evt)
}
func proc_staking_percentage_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGStakingPercentageChanged
	var eth_evt CosmicSignatureGameCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing StakingPercentageChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", log.Data)
	if err != nil {
		return fmt.Errorf("StakingPercentageChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewStakingPercentage = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged{\n")
	Info.Printf("\tNewStakingPercentage: %v\n", evt.NewStakingPercentage)
	Info.Printf("}\n")

	if err := cgRepo.DeleteStakingPercentageChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertStakingPercentageChange(ctx, &evt)
}
func proc_chrono_percentage_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGChronoPercentageChanged
	var eth_evt ISystemEventsChronoWarriorEthPrizeAmountPercentageChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	Info.Printf("Processing ChronoWarriorEthPrizePercentageChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "ChronoWarriorEthPrizeAmountPercentageChanged", log.Data)
	if err != nil {
		return fmt.Errorf("ChronoWarriorEthPrizePercentageChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewChronoPercentage = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("ChronoWarriorEthPrizePercentageChanged{\n")
	Info.Printf("\tNewPercentage: %v\n", evt.NewChronoPercentage)
	Info.Printf("}\n")

	if err := cgRepo.DeleteChronoPercentageChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertChronoPercentageChange(ctx, &evt)
}
func proc_num_raffle_eth_winners_bidding_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGNumRaffleETHWinnersBiddingChanged
	var eth_evt CosmicSignatureGameNumRaffleEthPrizesForBiddersChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing NumRaffleETHWinnersBiddingChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "NumRaffleEthPrizesForBiddersChanged", log.Data)
	if err != nil {
		return fmt.Errorf("NumRaffleEthPrizesForBiddersChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewNumRaffleETHWinnersBidding = eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("NumRaffleEthPrizesForBiddersChanged{\n")
	Info.Printf("\tNewNumRaffleETHWinnersBidding: %v\n", evt.NewNumRaffleETHWinnersBidding)
	Info.Printf("}\n")

	if err := cgRepo.DeleteNumRaffleETHWinnersBiddingChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertNumRaffleETHWinnersBiddingChange(ctx, &evt)
}
func proc_num_raffle_nft_winners_bidding_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGNumRaffleNFTWinnersBiddingChanged
	var eth_evt ISystemManagementNumRaffleCosmicSignatureNftsForBiddersChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing NumRaffleNftWinnersBiddingChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "NumRaffleCosmicSignatureNftsForBiddersChanged", log.Data)
	if err != nil {
		return fmt.Errorf("NumRaffleNftWinnersBiddingChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewNumRaffleNFTWinnersBidding = eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("NumRaffleNftWinnersBiddingChanged{\n")
	Info.Printf("\tNewNumRaffleNFTWinnersBidding: %v\n", evt.NewNumRaffleNFTWinnersBidding)
	Info.Printf("}\n")

	if err := cgRepo.DeleteNumRaffleNFTWinnersBiddingChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertNumRaffleNFTWinnersBiddingChange(ctx, &evt)
}
func proc_num_raffle_nft_winners_staking_rwalk_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGNumRaffleNFTWinnersStakingRWalkChanged
	var eth_evt ISystemManagementNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing NumRaffleNFTWinnersStakingRWalkChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", log.Data)
	if err != nil {
		return fmt.Errorf("NumRaffleNFTWinnersStakingRWalkChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewNumRaffleNFTWinnersStakingRWalk = eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("NumRaffleNFTWinnersStakingRWalkChanged{\n")
	Info.Printf("\tNewNumRaffleNFTWinnersStakingRWalk: %v\n", evt.NewNumRaffleNFTWinnersStakingRWalk)
	Info.Printf("}\n")

	if err := cgRepo.DeleteNumRaffleNFTWinnersStakingRWalkChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertNumRaffleNFTWinnersStakingRWalkChange(ctx, &evt)
}
func proc_random_walk_address_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGRandomWalkAddressChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing RandomWalkAddressChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewRandomWalk = common.BytesToAddress(log.Topics[1][12:]).String()
	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("RandomWalkAddressChanged{\n")
	Info.Printf("\tNewRandomWalk: %v\n", evt.NewRandomWalk)
	Info.Printf("}\n")

	if err := cgRepo.DeleteRandomWalkAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertRandomWalkAddressChange(ctx, &evt)
}
func proc_raffle_address_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGPrizeWalletAddressChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing EthPrizesWalletAddressChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewPrizeWallet = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("EthPrizesWalletAddressChanged{\n")
	Info.Printf("\tNewEthPrizesWallet: %v\n", evt.NewPrizeWallet)
	Info.Printf("}\n")

	if err := cgRepo.DeletePrizesWalletAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertPrizesWalletAddressChange(ctx, &evt)
}
func proc_staking_wallet_cst_address_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGStakingWalletCSTAddressChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing StakingWalletCSTAddressChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewStakingWalletCST = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("StakingWalletCSTAddressChanged{\n")
	Info.Printf("\tNewStakingWalletCST: %v\n", evt.NewStakingWalletCST)
	Info.Printf("}\n")

	if err := cgRepo.DeleteStakingWalletCSTAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertStakingWalletCSTAddressChange(ctx, &evt)
}
func proc_staking_wallet_rwalk_address_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGStakingWalletRWalkAddressChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing StakingWalletRWalkAddressChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewStakingWalletRWalk = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("StakingWalletRWalkAddressChanged{\n")
	Info.Printf("\tNewStakingWalletRWalk: %v\n", evt.NewStakingWalletRWalk)
	Info.Printf("}\n")

	if err := cgRepo.DeleteStakingWalletRWalkAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertStakingWalletRWalkAddressChange(ctx, &evt)
}
func proc_marketing_wallet_address_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGMarketingWalletAddressChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing MarketingWalletAddressChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewMarketingWallet = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("MarketingWalletAddressChanged{\n")
	Info.Printf("\tNewMarketingWallet: %v\n", evt.NewMarketingWallet)
	Info.Printf("}\n")

	if err := cgRepo.DeleteMarketingWalletAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertMarketingWalletAddressChange(ctx, &evt)
}
func proc_treasurer_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGTreasurerAddressChanged

	if !bytes.Equal(log.Address.Bytes(), marketing_wallet_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing TreasurerAddressChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewTreasurer = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("TreasurerAddressChanged{\n")
	Info.Printf("\tNewTreasurer: %v\n", evt.NewTreasurer)
	Info.Printf("}\n")

	if err := cgRepo.DeleteTreasurerAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertTreasurerAddressChange(ctx, &evt)
}
func proc_cosmic_token_address_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGCosmicTokenAddressChanged
	var eth_evt CosmicSignatureGameCosmicSignatureTokenAddressChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing CosmicSignatureTokenAddressChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "CosmicSignatureTokenAddressChanged", log.Data)
	if err != nil {
		return fmt.Errorf("CosmicSignatureTokenAddressChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCosmicToken = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("CosmicSignatureTokenAddressChanged{\n")
	Info.Printf("\tNewCosmicToken: %v\n", evt.NewCosmicToken)
	Info.Printf("}\n")

	if err := cgRepo.DeleteCosmicTokenAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertCosmicTokenAddressChange(ctx, &evt)
}
func proc_cosmic_signature_address_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGCosmicSignatureAddressChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing CosmicSignatureAddressChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCosmicSignature = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("CosmicSignatureAddressChanged{\n")
	Info.Printf("\tNewCosmicSignatureWallet: %v\n", evt.NewCosmicSignature)
	Info.Printf("}\n")

	if err := cgRepo.DeleteCosmicSignatureAddressChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertCosmicSignatureAddressChange(ctx, &evt)
}
func proc_proxy_upgraded_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGUpgraded
	var eth_evt CosmicSignatureGameUpgraded

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event Upgraded doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	Info.Printf("Processing Upgraded event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "Upgraded", log.Data)
	if err != nil {
		return fmt.Errorf("decoding Upgraded (evt id %v): %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Implementation = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("(Proxy) Upgraded{\n")
	Info.Printf("\tImplementation: %v\n", evt.Implementation)
	Info.Printf("}\n")

	if err := cgRepo.DeleteUpgraded(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertUpgraded(ctx, &evt)
}
func proc_admin_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGAdminChanged
	var eth_evt IERC1967AdminChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event AdminChanged doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	Info.Printf("Processing AdminChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	// AdminChanged is an ERC-1967 proxy event; it is absent from the game ABI,
	// so it must be decoded with the IERC1967 ABI (using the game ABI here made
	// this handler terminate the process on every AdminChanged event).
	err := erc1967_abi.UnpackIntoInterface(&eth_evt, "AdminChanged", log.Data)
	if err != nil {
		return fmt.Errorf("AdminChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.OldAdmin = eth_evt.PreviousAdmin.String()
	evt.NewAdmin = eth_evt.NewAdmin.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("AdminChanged{\n")
	Info.Printf("\tOldAdmin: %v\n", evt.OldAdmin)
	Info.Printf("\tNewAdmin: %v\n", evt.NewAdmin)
	Info.Printf("}\n")

	if err := cgRepo.DeleteAdminChanged(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertAdminChanged(ctx, &evt)
}
func proc_time_increase_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGTimeIncreaseChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing TimeIncreaseChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	// TimeIncreaseChanged(uint256) is a legacy contract event: no current ABI
	// defines it, so decode the single non-indexed uint256 from the raw data
	// (unpacking a name absent from the ABI made this handler terminate the
	// process on every TimeIncreaseChanged event).
	newValue, err := admin_uint256_from_log_data(log.Data)
	if err != nil {
		return fmt.Errorf("TimeIncreaseChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewTimeIncrease = newValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("TimeIncreaseChanged{\n")
	Info.Printf("\tNewTimeIncrease: %v\n", evt.NewTimeIncrease)
	Info.Printf("}\n")

	if err := cgRepo.DeleteTimeIncreaseChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertTimeIncreaseChange(ctx, &evt)
}
func proc_timeout_claimprize_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGTimeoutClaimPrizeChanged
	var eth_evt CosmicSignatureGameTimeoutDurationToClaimMainPrizeChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing TimeoutClaimPrizeChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "TimeoutDurationToClaimMainPrizeChanged", log.Data)
	if err != nil {
		return fmt.Errorf("TimeoutClaimPrizeChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewTimeout = eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("TimeoutDurationToClaimMainPrizeChanged{\n")
	Info.Printf("\tNewTimeout: %v\n", evt.NewTimeout)
	Info.Printf("}\n")

	if err := cgRepo.DeleteTimeoutClaimPrizeChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertTimeoutClaimPrizeChange(ctx, &evt)
}
func proc_timeout_duration_to_withdraw_prize_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGTimeoutToWithdrawPrizeChanged
	var eth_evt IPrizesWalletTimeoutDurationToWithdrawPrizesChanged

	if !bytes.Equal(log.Address.Bytes(), prizes_wallet_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing TimeoutDurationToWithdrawPrizesChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt, "TimeoutDurationToWithdrawPrizesChanged", log.Data)
	if err != nil {
		return fmt.Errorf("TimeoutDurationToWithdrawPrizesChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewTimeout = eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("TimeoutDurationToWithdrawPrizesChanged {\n")
	Info.Printf("\tNewTimeout: %v\n", evt.NewTimeout)
	Info.Printf("}\n")

	if err := cgRepo.DeleteTimeoutToWithdrawPrizesChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertTimeoutToWithdrawPrizesChange(ctx, &evt)
}
func proc_price_increase_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGPriceIncreaseChanged
	var eth_evt CosmicSignatureGameEthBidPriceIncreaseDivisorChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing PriceIncreaseChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "EthBidPriceIncreaseDivisorChanged", log.Data)
	if err != nil {
		return fmt.Errorf("EthBidPriceIncreaseDivisorChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewPriceIncrease = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("EthBidPriceIncreaseDivisorChanged{\n")
	Info.Printf("\tNewPriceIncreasse: %v\n", evt.NewPriceIncrease)
	Info.Printf("}\n")

	if err := cgRepo.DeletePriceIncreaseChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertPriceIncreaseChange(ctx, &evt)
}
func proc_mainprize_microsecond_increase_changed(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGMainPrizeMicroSecondsIncreaseChanged
	var eth_evt CosmicSignatureGameMainPrizeTimeIncrementInMicroSecondsChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing MainPrizeTimeIncrementInMicroSecondsChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "MainPrizeTimeIncrementInMicroSecondsChanged", log.Data)
	if err != nil {
		return fmt.Errorf("MainPrizeTimeIncrementInMicroSecondsChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewMicroseconds = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("MainPrizeTimeIncrementInMicroSecondsChanged{\n")
	Info.Printf("\tNewMicroseconds: %v\n", evt.NewMicroseconds)
	Info.Printf("}\n")

	if err := cgRepo.DeleteMainPrizeMicrosecondsChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertMainPrizeMicrosecondsChange(ctx, &evt)
}
func proc_initial_seconds_until_prize_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGInitialSecondsUntilPrizeChanged
	var eth_evt CosmicSignatureGameInitialDurationUntilMainPrizeDivisorChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing InitialDurationUntilMainPrizeDivisorChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "InitialDurationUntilMainPrizeDivisorChanged", log.Data)
	if err != nil {
		return fmt.Errorf("InitialDurationUntilMainPrizeDivisorChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewInitialSecondsUntilPrize = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("InitialDurationUntilMainPrizeDivisorChanged{\n")
	Info.Printf("\tNewInitialSecondsUntilPrize: %v\n", evt.NewInitialSecondsUntilPrize)
	Info.Printf("}\n")

	if err := cgRepo.DeleteInitialSecondsUntilPrizeChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertInitialSecondsUntilPrizeChange(ctx, &evt)
}
func proc_activation_time_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGActivationTimeChanged
	var eth_evt BiddingBaseRoundActivationTimeChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing RoundActivationTimeChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "RoundActivationTimeChanged", log.Data)
	if err != nil {
		return fmt.Errorf("RoundActivationTimeChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewActivationTime = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("ActivationTimeChanged {\n")
	Info.Printf("\tNewActivationTime: %v\n", evt.NewActivationTime)
	Info.Printf("}\n")

	if err := cgRepo.DeleteActivationTimeChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertActivationTimeChange(ctx, &evt)
}
func proc_cst_dutch_auction_duration_divisor_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var eth_evt CosmicSignatureGameCstDutchAuctionDurationDivisorChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing CstDutchAuctionDurationDivisorChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "CstDutchAuctionDurationDivisorChanged", log.Data)
	if err != nil {
		return fmt.Errorf("CstDutchAuctionDurationDivisorChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	return store_cst_dutch_auction_length_changed(ctx, log, elog, eth_evt.NewValue.String(), "CstDutchAuctionDurationDivisorChanged")
}
func proc_cst_dutch_auction_duration_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var eth_evt CosmicSignatureGameV2CstDutchAuctionDurationChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing CstDutchAuctionDurationChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_v2_abi.UnpackIntoInterface(&eth_evt, "CstDutchAuctionDurationChanged", log.Data)
	if err != nil {
		return fmt.Errorf("CstDutchAuctionDurationChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	return store_cst_dutch_auction_length_changed(ctx, log, elog, eth_evt.NewValue.String(), "CstDutchAuctionDurationChanged")
}
func proc_cst_dutch_auction_duration_change_divisor_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGCstDutchAuctionDurationChangeDivisorChanged
	var eth_evt CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing CstDutchAuctionDurationChangeDivisorChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_v2_abi.UnpackIntoInterface(&eth_evt, "CstDutchAuctionDurationChangeDivisorChanged", log.Data)
	if err != nil {
		return fmt.Errorf("CstDutchAuctionDurationChangeDivisorChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewValue = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("CstDutchAuctionDurationChangeDivisorChanged {\n")
	Info.Printf("\tNewDivisor: %v\n", evt.NewValue)
	Info.Printf("}\n")

	if err := cgRepo.DeleteCstAuctionDurationChangeDivisorChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertCstAuctionDurationChangeDivisorChange(ctx, &evt)
}
func proc_eth_dutch_auction_duration_divisor_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGEthDutchAuctionDurationDivisorChanged
	var eth_evt CosmicSignatureGameEthDutchAuctionDurationDivisorChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing EthDutchAuctionDurationDivisorChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "EthDutchAuctionDurationDivisorChanged", log.Data)
	if err != nil {
		return fmt.Errorf("EthDutchAuctionDurationDivisorChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewValue = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("EthDutchAuctionDurationDivisorChanged {\n")
	Info.Printf("\tNewDivisor: %v\n", evt.NewValue)
	Info.Printf("}\n")

	if err := cgRepo.DeleteEthAuctionDurationDivisorChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertEthAuctionDurationDivisorChange(ctx, &evt)
}
func proc_eth_dutch_auction_ending_bid_price_divisor_changed__event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGEthDutchAuctionEndingBidPriceDivisorChanged
	var eth_evt CosmicSignatureGameEthDutchAuctionEndingBidPriceDivisorChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing EthDutchAuctionEndingBidPriceDivisorChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "EthDutchAuctionEndingBidPriceDivisorChanged", log.Data)
	if err != nil {
		return fmt.Errorf("EthDutchAuctionEndingBidPriceDivisorChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewValue = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("EthDutchAuctionEndingBidPriceDivisorChanged{\n")
	Info.Printf("\tNewDivisor: %v\n", evt.NewValue)
	Info.Printf("}\n")

	if err := cgRepo.DeleteEthAuctionEndingBidPriceDivisorChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertEthAuctionEndingBidPriceDivisorChange(ctx, &evt)
}
func proc_marketing_reward_changed(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGMarketingRewardChanged
	var eth_evt CosmicSignatureGameMarketingWalletCstContributionAmountChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing MarketingWalletCstContributionAmountChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "MarketingWalletCstContributionAmountChanged", log.Data)
	if err != nil {
		return fmt.Errorf("MarketingWalletCstContributionAmountChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewReward = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("MarketingWalletCstContributionAmountChanged{\n")
	Info.Printf("\tNewReward: %v\n", evt.NewReward)
	Info.Printf("}\n")

	if err := cgRepo.DeleteMarketingRewardChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertMarketingRewardChange(ctx, &evt)
}
func proc_erc20_token_reward_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var eth_evt CosmicSignatureGameCstRewardAmountForBiddingChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing CstRewardAmountForBiddingChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "CstRewardAmountForBiddingChanged", log.Data)
	if err != nil {
		return fmt.Errorf("CstRewardAmountForBiddingChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	return store_cst_reward_for_bidding_changed(ctx, log, elog, eth_evt.NewValue.String(), "CstRewardAmountForBiddingChanged")
}
func proc_bid_cst_reward_amount_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing BidCstRewardAmountChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	newValue, err := admin_uint256_from_log_data(log.Data)
	if err != nil {
		return fmt.Errorf("BidCstRewardAmountChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	return store_cst_reward_for_bidding_changed(ctx, log, elog, newValue.String(), "BidCstRewardAmountChanged")
}
func proc_bid_cst_reward_amount_multiplier_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var eth_evt CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing BidCstRewardAmountMultiplierChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_v2_abi.UnpackIntoInterface(&eth_evt, "BidCstRewardAmountMultiplierChanged", log.Data)
	if err != nil {
		return fmt.Errorf("BidCstRewardAmountMultiplierChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	return store_cst_reward_for_bidding_changed(ctx, log, elog, eth_evt.NewValue.String(), "BidCstRewardAmountMultiplierChanged")
}
func proc_static_cst_reward_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGStaticCstReward
	var eth_evt BiddingCstPrizeAmountChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing CstPrizeAmountChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "CstPrizeAmountChanged", log.Data)
	if err != nil {
		return fmt.Errorf("CstPrizeAmountChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewReward = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("CstPrizeAmountChanged{\n")
	Info.Printf("\tNewReward: %v\n", evt.NewReward)
	Info.Printf("}\n")

	if err := cgRepo.DeleteStaticCstRewardChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertStaticCstRewardChange(ctx, &evt)
}
func proc_max_msg_length_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGMaxMessageLengthChanged
	var eth_evt CosmicSignatureGameBidMessageLengthMaxLimitChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing BidMessageLengthMaxLimitChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "BidMessageLengthMaxLimitChanged", log.Data)
	if err != nil {
		return fmt.Errorf("BidMessageLengthMaxLimitChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewMessageLength = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("BidMessageLengthMaxLimitChanged{\n")
	Info.Printf("\tNewMessageLength: %v\n", evt.NewMessageLength)
	Info.Printf("}\n")

	if err := cgRepo.DeleteMaxMessageLengthChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertMaxMessageLengthChange(ctx, &evt)
}
func proc_token_generation_script_url_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGTokenGenerationScriptURL
	var eth_evt ICosmicSignatureNftNftGenerationScriptUriChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_signature_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing TokenGenerationScriptURLEvent event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt, "NftGenerationScriptUriChanged", log.Data)
	if err != nil {
		return fmt.Errorf("TokenGenerationScriptURLEvent (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewURL = eth_evt.NewValue

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("TokenGenerationScriptURLEvent{\n")
	Info.Printf("\tNewURL: %v\n", evt.NewURL)
	Info.Printf("}\n")

	// Must delete from cg_adm_script_url (this event's own table); deleting
	// from the message-length table here made every re-processed script-URL
	// event abort on the cg_adm_script_url unique constraint.
	if err := cgRepo.DeleteTokenGenerationScriptURL(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertTokenGenerationScriptURL(ctx, &evt)
}
func proc_base_uri_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGBaseURIEvent
	var eth_evt CosmicSignatureNftNftBaseUriChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_signature_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing BaseURIEvent event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt, "NftBaseUriChanged", log.Data)
	if err != nil {
		return fmt.Errorf("BaseURIEvent (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewURI = eth_evt.NewValue

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("BaseURIEvent{\n")
	Info.Printf("\tNewURI: %v\n", evt.NewURI)
	Info.Printf("}\n")

	if err := cgRepo.DeleteBaseURI(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertBaseURI(ctx, &evt)
}
func proc_ownership_transferred_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGOwnershipTransferred
	var eth_evt CosmicSignatureGameOwnershipTransferred

	contract_code := int64(0)
	if bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		contract_code = 1
	}
	if bytes.Equal(log.Address.Bytes(), cosmic_signature_addr.Bytes()) {
		contract_code = 2
	}
	if bytes.Equal(log.Address.Bytes(), cosmic_token_addr.Bytes()) {
		contract_code = 3
	}
	if bytes.Equal(log.Address.Bytes(), charity_wallet_addr.Bytes()) {
		contract_code = 4
	}
	if bytes.Equal(log.Address.Bytes(), prizes_wallet_addr.Bytes()) {
		contract_code = 5
	}
	if bytes.Equal(log.Address.Bytes(), staking_wallet_cst_addr.Bytes()) {
		contract_code = 6
	}
	if bytes.Equal(log.Address.Bytes(), staking_wallet_rwalk_addr.Bytes()) {
		contract_code = 7
	}
	if bytes.Equal(log.Address.Bytes(), marketing_wallet_addr.Bytes()) {
		contract_code = 8
	}
	if bytes.Equal(log.Address.Bytes(), cosmic_dao_addr.Bytes()) {
		contract_code = 9
	}
	if contract_code == 0 {
		return nil
	}
	Info.Printf("Processing OwnershipTransferred event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt, "OwnershipTransferred", log.Data)
	if err != nil {
		return fmt.Errorf("OwnershipTransferred (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.PrevOwner = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.NewOwner = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.ContractCode = contract_code

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("OwnershipTransferred{\n")
	Info.Printf("\tPrevOwner: %v\n", evt.PrevOwner)
	Info.Printf("\tNewOwner: %v\n", evt.NewOwner)
	Info.Printf("}\n")

	if err := cgRepo.DeleteOwnershipTransfer(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertOwnershipTransfer(ctx, &evt)
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

func proc_initialized_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGInitialized
	var eth_evt CosmicSignatureGameInitialized

	if !isCgInitializedContract(log.Address) {
		return nil
	}
	Info.Printf("Processing Initialized event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "Initialized", log.Data)
	if err != nil {
		return fmt.Errorf("decoding Initialized (evt id %v): %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Version = int64(eth_evt.Version)

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("Initialized {\n")
	Info.Printf("\tVersion: %v\n", evt.Version)
	Info.Printf("}\n")

	if err := cgRepo.DeleteInitialized(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertInitialized(ctx, &evt)
}
func proc_starting_bid_price_cst_min_limit_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGCstMinLimit
	var eth_evt BiddingCstDutchAuctionBeginningBidPriceMinLimitChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing CstDutchAuctionBeginningBidPriceMinLimitChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "CstDutchAuctionBeginningBidPriceMinLimitChanged", log.Data)
	if err != nil {
		return fmt.Errorf("CstDutchAuctionBeginningBidPriceMinLimitChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.CstMinLimit = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("CstDutchAuctionBeginningBidPriceMinLimitChanged {\n")
	Info.Printf("\tNewStartingBidPriceCSTMinLimit: %v\n", evt.CstMinLimit)
	Info.Printf("}\n")

	if err := cgRepo.DeleteCstMinLimit(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertCstMinLimit(ctx, &evt)
}
func proc_delay_duration_before_next_round_changed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGNextRoundDelayDuration
	var eth_evt CosmicSignatureGameDelayDurationBeforeRoundActivationChanged

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing DelayDurationBeforeRoundActivationChanged event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "DelayDurationBeforeRoundActivationChanged", log.Data)
	if err != nil {
		return fmt.Errorf("DelayDurationBeforeRoundActivationChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewValue = eth_evt.NewValue.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("DelayDurationBeforeRoundActivationChanged{\n")
	Info.Printf("\tNewValue: %v\n", evt.NewValue)
	Info.Printf("}\n")

	if err := cgRepo.DeleteNextRoundDelayDurationChange(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertNextRoundDelayDurationChange(ctx, &evt)
}
