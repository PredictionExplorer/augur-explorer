// Prize events: main prize claims, raffle ETH/NFT prizes, endurance champion,
// chrono warrior and last-CST-bidder prizes, PrizesWallet ETH deposits/withdrawals.
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
func proc_prize_claim_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGPrizeClaimEvent
	var eth_evt CosmicSignatureGameMainPrizeClaimed

	Info.Printf("Processing PrizeClaim event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"MainPrizeClaimed",log.Data)
	if err != nil {
		Error.Printf("Event MainPrizeClaimed decode error: %v",err)
		os.Exit(1)
	}
	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.RoundNum= log.Topics[1].Big().Int64()
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Amount = eth_evt.EthPrizeAmount.String()
	evt.CstAmount = eth_evt.CstPrizeAmount.String()
	evt.TokenId = log.Topics[3].Big().Int64()
	evt.Timeout = eth_evt.TimeoutTimeToWithdrawSecondaryPrizes.Int64()
//	find_cosmic_token_721_mint_event(cosmic_sig_aid,evt.TxId,evt.EvtId)
//	evt.DonationEvtId = storagew.Get_donation_received_evt_id_by_tx_id(evt.TxId,hex.EncodeToString(evt_donation_received_event[:4]))
//	if evt.DonationEvtId == 0 {
//		Error.Printf("Failed to fetch donation received event id for txid=%v\n",evt,TxId)
//		Info.Printf("Failed to fetch donation received event id for txid=%v\n",evt.TxId)
//		os.Exit(1)
//	}

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("MainPrizeClaimed {\n")
	Info.Printf("\tRoundNum: %v\n",evt.RoundNum)
	Info.Printf("\tWinner%v\n",evt.WinnerAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tCstAmount: %v\n",evt.CstAmount)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tTimeout to withdraw: %v\n",evt.Timeout)
	Info.Printf("}\n")

	storagew.Delete_prize_claim_event(evt.EvtId)
	storagew.Insert_prize_claim_event(&evt)
}
func proc_prizes_eth_deposit_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGPrizesEthDeposit
	var eth_evt IPrizesWalletEthReceived  

	Info.Printf("Processing PrizeReceived event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt,"EthReceived",log.Data)
	if err != nil {
		Error.Printf("Event PrizeReceived decode error: %v",err)
		os.Exit(1)
	}
	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Round = log.Topics[1].Big().Int64()
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.WinnerIndex = eth_evt.PrizeWinnerIndex.Int64()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("EthReceived{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tWinnerIndex:%v\n",evt.WinnerIndex)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Delete_prize_deposit(evt.EvtId)
	storagew.Insert_prize_deposit(&evt)
}
func proc_eth_prize_withdrawal_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGPrizesEthWithdrawal
	var eth_evt IPrizesWalletEthWithdrawn

	Info.Printf("Processing RaffleWithdrawalevent id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt,"EthWithdrawn",log.Data)
	if err != nil {
		Error.Printf("Event PrizeWithdrawn decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Round = log.Topics[1].Big().Int64()
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.BeneficiaryAddr = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PrizeWithdrawn {\n")
	Info.Printf("\tRound: %v\n",evt.Round)
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tBeneficiaryAddr: %v\n",evt.BeneficiaryAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Delete_prize_withdrawal(evt.EvtId)
	storagew.Insert_prize_withdrawal(&evt)
}
func proc_raffle_nft_winner_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRaffleNFTWinner
	var eth_evt CosmicSignatureGameRaffleWinnerPrizePaid

	Info.Printf("Processing RaffleNFTWinner event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"RaffleWinnerPrizePaid",log.Data)
	if err != nil {
		Error.Printf("Event RaffleWinnerCosmicSignatureNftAwarded decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Round = log.Topics[1].Big().Int64()
	evt.TokenId = log.Topics[3].Big().Int64()
	evt.WinnerIndex= eth_evt.WinnerIndex.Int64()
	evt.CstAmount = eth_evt.CstPrizeAmount.String()
	evt.IsRandomWalk = eth_evt.WinnerIsRandomWalkNftStaker
	evt.IsStaker = evt.IsRandomWalk

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RaffleNftWinnerEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tWinnerIndex: %v\n",evt.WinnerIndex)
	Info.Printf("\tCstAmount: %v\n",evt.CstAmount)
	Info.Printf("\tIsStaker: %v\n",evt.IsStaker);
	Info.Printf("\tIsRandomWalk: %v\n",evt.IsRandomWalk)
	Info.Printf("}\n")

	storagew.Delete_raffle_nft_winner(evt.EvtId)
	storagew.Insert_raffle_nft_winner(&evt)
}
func proc_raffle_eth_winner_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRaffleETHWinner
	var eth_evt CosmicSignatureGameRaffleWinnerBidderEthPrizeAllocated

	Info.Printf("Processing RaffleETHWinner event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"RaffleWinnerBidderEthPrizeAllocated",log.Data)
	if err != nil {
		Error.Printf("Event RaffleWinnerBidderEthPrizeAllocated decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Round = log.Topics[1].Big().Int64()
	evt.WinnerIndex= eth_evt.WinnerIndex.Int64()
	evt.Amount = eth_evt.EthPrizeAmount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RaffleWinnerBidderEthPrizeAllocated{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tWinnerIndex: %v\n",evt.WinnerIndex)
	Info.Printf("}\n")

	storagew.Delete_raffle_eth_winner(evt.EvtId)
	storagew.Insert_raffle_eth_winner(&evt)
}
func proc_endurance_winner_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGEnduranceWinner
	var eth_evt ICosmicSignatureGameEnduranceChampionPrizePaid 

	Info.Printf("Processing Endurance winner event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"EnduranceChampionPrizePaid",log.Data)
	if err != nil {
		Error.Printf("Event EnduranceChampionPrizePaid decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Round = log.Topics[1].Big().Int64()
	evt.Erc721TokenId = log.Topics[3].Big().Int64()
	evt.Erc20Amount = eth_evt.CstPrizeAmount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("EnduranceChampionPrizePaid {\n")
	Info.Printf("\tEnduranceChampion : %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tErc721TokenId: %v\n",evt.Erc721TokenId)
	Info.Printf("\tErc20Amount: %v\n",evt.Erc20Amount)
	Info.Printf("}\n")

	storagew.Delete_endurance_winner(evt.EvtId)
	storagew.Insert_endurance_winner(&evt)
}
func proc_lastcst_bidder_winner_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGLastBidderWinner
	var eth_evt CosmicSignatureGameLastCstBidderPrizePaid 
	Info.Printf("Processing LastCstBidderwinner event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"LastCstBidderPrizePaid",log.Data)
	if err != nil {
		Error.Printf("Event LastCstBidderPrizePaid decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Round = log.Topics[1].Big().Int64()
	evt.Erc721TokenId = log.Topics[3].Big().Int64()
	evt.Erc20Amount = eth_evt.CstPrizeAmount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("LastCstBidderPrizePaidEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tErc721TokenId: %v\n",evt.Erc721TokenId)
	Info.Printf("\tErc20TokenId: %v\n",evt.Erc20Amount)
	Info.Printf("\tWinnerIndex: %v\n",evt.WinnerIndex)
	Info.Printf("}\n")

	storagew.Delete_lastcst_bidder_winner(evt.EvtId)
	storagew.Insert_lastcst_bidder_winner(&evt)
}
func proc_chrono_warrior_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGChronoWarrior
	var eth_evt ICosmicSignatureGameChronoWarriorPrizePaid
	Info.Printf("Processing ChronoWarrior prize event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"ChronoWarriorPrizePaid",log.Data)
	if err != nil {
		Error.Printf("Event ChronoWarriorPrizePaid decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Round = log.Topics[1].Big().Int64()
	evt.WinnerIndex = eth_evt.WinnerIndex.Int64()
	evt.EthAmount = eth_evt.EthPrizeAmount.String()
	evt.CstAmount = eth_evt.CstPrizeAmount.String()
	evt.NftId = log.Topics[3].Big().Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("ChronoWarriorPrizePaid {\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tWinnerIndex:%v\n",evt.WinnerIndex)
	Info.Printf("\tEthAmount: %v\n",evt.EthAmount)
	Info.Printf("\tCstAmount: %v\n",evt.CstAmount)
	Info.Printf("\tNftId: %v\n",evt.NftId)
	Info.Printf("}\n")

	storagew.Delete_chrono_warrior_event(evt.EvtId)
	storagew.Insert_chrono_warrior_event(&evt)
}
func proc_fund_transfer_failed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGFundTransferFailed
	var eth_evt CosmicSignatureGameFundTransferFailed

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing Initialized event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"FundTransferFailed",log.Data)
	if err != nil {
		Error.Printf("Event Initialized decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Destination= common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount= eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("FundTransferFailed {\n")
	Info.Printf("\tDestination: %v\n",evt.Destination)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Delete_fund_transfer_failed_event(evt.EvtId)
    storagew.Insert_fund_transfer_failed_event(&evt)
}
