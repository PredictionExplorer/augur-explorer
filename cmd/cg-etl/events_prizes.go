// Prize events: main prize claims, raffle ETH/NFT prizes, endurance champion,
// chrono warrior and last-CST-bidder prizes, PrizesWallet ETH deposits/withdrawals,
// and the fund/ERC20 transfer-failure diagnostics.
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

func proc_prize_claim_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGPrizeClaimEvent
	var eth_evt CosmicSignatureGameMainPrizeClaimed

	Info.Printf("Processing PrizeClaim event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "MainPrizeClaimed", log.Data)
	if err != nil {
		return fmt.Errorf("MainPrizeClaimed (evt id %v): decode: %w", elog.EvtId, err)
	}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Amount = eth_evt.EthPrizeAmount.String()
	evt.CstAmount = eth_evt.CstPrizeAmount.String()
	evt.TokenId = log.Topics[3].Big().Int64()
	evt.Timeout = eth_evt.TimeoutTimeToWithdrawSecondaryPrizes.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("MainPrizeClaimed {\n")
	Info.Printf("\tRoundNum: %v\n", evt.RoundNum)
	Info.Printf("\tWinner%v\n", evt.WinnerAddr)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("\tCstAmount: %v\n", evt.CstAmount)
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("\tTimeout to withdraw: %v\n", evt.Timeout)
	Info.Printf("}\n")

	if err := cgRepo.DeletePrizeClaim(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertPrizeClaim(ctx, &evt)
}
func proc_prizes_eth_deposit_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGPrizesEthDeposit
	var eth_evt IPrizesWalletEthReceived

	Info.Printf("Processing PrizeReceived event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt, "EthReceived", log.Data)
	if err != nil {
		return fmt.Errorf("EthReceived (evt id %v): decode: %w", elog.EvtId, err)
	}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Round = log.Topics[1].Big().Int64()
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.WinnerIndex = eth_evt.PrizeWinnerIndex.Int64()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("EthReceived{\n")
	Info.Printf("\tWinnerAddr: %v\n", evt.WinnerAddr)
	Info.Printf("\tRound:%v\n", evt.Round)
	Info.Printf("\tWinnerIndex:%v\n", evt.WinnerIndex)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("}\n")

	if err := cgRepo.DeletePrizeDeposit(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertPrizeDeposit(ctx, &evt)
}
func proc_eth_prize_withdrawal_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGPrizesEthWithdrawal
	var eth_evt IPrizesWalletEthWithdrawn

	Info.Printf("Processing RaffleWithdrawalevent id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt, "EthWithdrawn", log.Data)
	if err != nil {
		return fmt.Errorf("EthWithdrawn (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Round = log.Topics[1].Big().Int64()
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.BeneficiaryAddr = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("PrizeWithdrawn {\n")
	Info.Printf("\tRound: %v\n", evt.Round)
	Info.Printf("\tWinnerAddr: %v\n", evt.WinnerAddr)
	Info.Printf("\tBeneficiaryAddr: %v\n", evt.BeneficiaryAddr)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("}\n")

	if err := cgRepo.DeletePrizeWithdrawal(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertPrizeWithdrawal(ctx, &evt)
}
func proc_raffle_nft_winner_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGRaffleNFTWinner
	var eth_evt CosmicSignatureGameRaffleWinnerPrizePaid

	Info.Printf("Processing RaffleNFTWinner event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "RaffleWinnerPrizePaid", log.Data)
	if err != nil {
		return fmt.Errorf("RaffleWinnerPrizePaid (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Round = log.Topics[1].Big().Int64()
	evt.TokenId = log.Topics[3].Big().Int64()
	evt.WinnerIndex = eth_evt.WinnerIndex.Int64()
	evt.CstAmount = eth_evt.CstPrizeAmount.String()
	evt.IsRandomWalk = eth_evt.WinnerIsRandomWalkNftStaker
	evt.IsStaker = evt.IsRandomWalk

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("RaffleNftWinnerEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n", evt.WinnerAddr)
	Info.Printf("\tRound:%v\n", evt.Round)
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("\tWinnerIndex: %v\n", evt.WinnerIndex)
	Info.Printf("\tCstAmount: %v\n", evt.CstAmount)
	Info.Printf("\tIsStaker: %v\n", evt.IsStaker)
	Info.Printf("\tIsRandomWalk: %v\n", evt.IsRandomWalk)
	Info.Printf("}\n")

	if err := cgRepo.DeleteRaffleNFTWinner(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertRaffleNFTWinner(ctx, &evt)
}
func proc_raffle_eth_winner_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGRaffleETHWinner
	var eth_evt CosmicSignatureGameRaffleWinnerBidderEthPrizeAllocated

	Info.Printf("Processing RaffleETHWinner event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "RaffleWinnerBidderEthPrizeAllocated", log.Data)
	if err != nil {
		return fmt.Errorf("RaffleWinnerBidderEthPrizeAllocated (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Round = log.Topics[1].Big().Int64()
	evt.WinnerIndex = eth_evt.WinnerIndex.Int64()
	evt.Amount = eth_evt.EthPrizeAmount.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("RaffleWinnerBidderEthPrizeAllocated{\n")
	Info.Printf("\tWinnerAddr: %v\n", evt.WinnerAddr)
	Info.Printf("\tRound:%v\n", evt.Round)
	Info.Printf("\tWinnerIndex: %v\n", evt.WinnerIndex)
	Info.Printf("}\n")

	if err := cgRepo.DeleteRaffleETHWinner(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertRaffleETHWinner(ctx, &evt)
}
func proc_endurance_winner_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGEnduranceWinner
	var eth_evt ICosmicSignatureGameEnduranceChampionPrizePaid

	Info.Printf("Processing Endurance winner event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "EnduranceChampionPrizePaid", log.Data)
	if err != nil {
		return fmt.Errorf("EnduranceChampionPrizePaid (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Round = log.Topics[1].Big().Int64()
	evt.Erc721TokenId = log.Topics[3].Big().Int64()
	evt.Erc20Amount = eth_evt.CstPrizeAmount.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("EnduranceChampionPrizePaid {\n")
	Info.Printf("\tEnduranceChampion : %v\n", evt.WinnerAddr)
	Info.Printf("\tRound:%v\n", evt.Round)
	Info.Printf("\tErc721TokenId: %v\n", evt.Erc721TokenId)
	Info.Printf("\tErc20Amount: %v\n", evt.Erc20Amount)
	Info.Printf("}\n")

	if err := cgRepo.DeleteEnduranceWinner(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertEnduranceWinner(ctx, &evt)
}
func proc_lastcst_bidder_winner_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGLastBidderWinner
	var eth_evt CosmicSignatureGameLastCstBidderPrizePaid
	Info.Printf("Processing LastCstBidderwinner event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "LastCstBidderPrizePaid", log.Data)
	if err != nil {
		return fmt.Errorf("LastCstBidderPrizePaid (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Round = log.Topics[1].Big().Int64()
	evt.Erc721TokenId = log.Topics[3].Big().Int64()
	evt.Erc20Amount = eth_evt.CstPrizeAmount.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("LastCstBidderPrizePaidEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n", evt.WinnerAddr)
	Info.Printf("\tRound:%v\n", evt.Round)
	Info.Printf("\tErc721TokenId: %v\n", evt.Erc721TokenId)
	Info.Printf("\tErc20TokenId: %v\n", evt.Erc20Amount)
	Info.Printf("\tWinnerIndex: %v\n", evt.WinnerIndex)
	Info.Printf("}\n")

	if err := cgRepo.DeleteLastCstBidderWinner(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertLastCstBidderWinner(ctx, &evt)
}
func proc_chrono_warrior_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGChronoWarrior
	var eth_evt ICosmicSignatureGameChronoWarriorPrizePaid
	Info.Printf("Processing ChronoWarrior prize event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "ChronoWarriorPrizePaid", log.Data)
	if err != nil {
		return fmt.Errorf("ChronoWarriorPrizePaid (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
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

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("ChronoWarriorPrizePaid {\n")
	Info.Printf("\tWinnerAddr: %v\n", evt.WinnerAddr)
	Info.Printf("\tRound:%v\n", evt.Round)
	Info.Printf("\tWinnerIndex:%v\n", evt.WinnerIndex)
	Info.Printf("\tEthAmount: %v\n", evt.EthAmount)
	Info.Printf("\tCstAmount: %v\n", evt.CstAmount)
	Info.Printf("\tNftId: %v\n", evt.NftId)
	Info.Printf("}\n")

	if err := cgRepo.DeleteChronoWarrior(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertChronoWarrior(ctx, &evt)
}
func proc_fund_transfer_failed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGFundTransferFailed
	var eth_evt CosmicSignatureGameFundTransferFailed

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing FundTransferFailed event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "FundTransferFailed", log.Data)
	if err != nil {
		return fmt.Errorf("FundTransferFailed (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Destination = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("FundTransferFailed {\n")
	Info.Printf("\tDestination: %v\n", evt.Destination)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("}\n")

	if err := cgRepo.DeleteFundTransferFailed(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertFundTransferFailed(ctx, &evt)
}

// proc_erc20_transfer_failed_event handles ICosmicSignatureErrors.sol:
// ERC20TransferFailed(string errStr, address indexed destination, uint256
// amount), emitted by the game when a CST transfer cannot complete. The
// registry inspected this topic but never dispatched it, so the events were
// fetched and silently dropped while cg_erc20_transf_err stayed empty. The
// event is absent from every generated ABI, so the amount is decoded from
// the raw data words (see erc20_transfer_failed_amount).
func proc_erc20_transfer_failed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGErc20TransferFailed

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing ERC20TransferFailed event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	if len(log.Topics) < 2 {
		return fmt.Errorf("ERC20TransferFailed (evt id %v): %d topics, want 2 (destination is indexed)", elog.EvtId, len(log.Topics))
	}
	amount, err := erc20_transfer_failed_amount(log.Data)
	if err != nil {
		return fmt.Errorf("ERC20TransferFailed (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Destination = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = amount.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("ERC20TransferFailed {\n")
	Info.Printf("\tDestination: %v\n", evt.Destination)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("}\n")

	if err := cgRepo.DeleteERC20TransferFailed(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertERC20TransferFailed(ctx, &evt)
}
