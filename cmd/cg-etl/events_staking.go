// Staking events for the CST and RandomWalk staking wallets: stake/unstake
// actions and staking-reward ETH deposits.
package main

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

func proc_cst_nft_staked_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGNftStakedCst
	var eth_evt IStakingWalletCosmicSignatureNftNftUnstaked

	Info.Printf("Processing CST NftStaked event id=%v, (block %v) txhash %v\n", elog.EvtId, elog.BlockNum, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), staking_wallet_cst_addr.Bytes()) {
		if !bytes.Equal(log.Address.Bytes(), staking_wallet_rwalk_addr.Bytes()) {
			Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
			return nil
		}
	}
	err := staking_wallet_cst_abi.UnpackIntoInterface(&eth_evt, "NftStaked", log.Data)
	if err != nil {
		return fmt.Errorf("CST NftStaked (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.NftId = log.Topics[2].Big().Int64()
	evt.StakerAddress = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.NumStakedNfts = eth_evt.NumStakedNfts.Int64()
	evt.RewardPerStaker = eth_evt.RewardAmountPerStakedNft.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("CST NftStaked event{\n")
	Info.Printf("\tActionId: %v\n", evt.ActionId)
	Info.Printf("\tTokenId: %v\n", evt.NftId)
	Info.Printf("\tTotalNFTs: %v\n", evt.NumStakedNfts)
	Info.Printf("\tRewardPerStaker: %v\n", evt.RewardPerStaker)
	Info.Printf("\tStaker: %v\n", evt.StakerAddress)
	Info.Printf("}\n")

	if err := cgRepo.DeleteNftStakedCST(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertNftStakedCST(ctx, &evt)
}
func proc_rwalk_nft_staked_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGNftStakedRWalk
	var eth_evt IStakingWalletRandomWalkNftNftStaked

	Info.Printf("Processing RWalk NftStaked event id=%v, (block %v) txhash %v\n", elog.EvtId, elog.BlockNum, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), staking_wallet_cst_addr.Bytes()) {
		if !bytes.Equal(log.Address.Bytes(), staking_wallet_rwalk_addr.Bytes()) {
			Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
			return nil
		}
	}
	err := staking_wallet_rwalk_abi.UnpackIntoInterface(&eth_evt, "NftStaked", log.Data)
	if err != nil {
		return fmt.Errorf("RWalk NftStaked (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.NftId = log.Topics[2].Big().Int64()
	evt.StakerAddress = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.NumStakedNfts = eth_evt.NumStakedNfts.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("RWalk NftStaked event{\n")
	Info.Printf("\tActionId: %v\n", evt.ActionId)
	Info.Printf("\tTokenId: %v\n", evt.NftId)
	Info.Printf("\tTotalNFTs: %v\n", evt.NumStakedNfts)
	Info.Printf("\tStaker: %v\n", evt.StakerAddress)
	Info.Printf("}\n")

	if err := cgRepo.DeleteNftStakedRWalk(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertNftStakedRWalk(ctx, &evt)
}
func proc_staking_eth_deposit_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGEthDeposit
	var eth_evt IStakingWalletCosmicSignatureNftEthDepositReceived

	Info.Printf("Processing EthDepositReceived event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), staking_wallet_cst_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := staking_wallet_cst_abi.UnpackIntoInterface(&eth_evt, "EthDepositReceived", log.Data)
	if err != nil {
		return fmt.Errorf("EthDepositReceived (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DepositTime = elog.TimeStamp
	evt.DepositId = eth_evt.ActionCounter.Int64()
	evt.NumStakedNfts = eth_evt.NumStakedNfts.Int64()
	evt.Amount = eth_evt.DepositAmount.String()
	evt.AccumModulo = "0" // pending for resolution regarding StakingWalletCST refactoring
	evt.RoundNum = log.Topics[1].Big().Int64()
	divres := big.NewInt(0)
	rem := big.NewInt(0)
	divres.QuoRem(eth_evt.DepositAmount, eth_evt.NumStakedNfts, rem)
	evt.AmountPerStaker = divres.String()
	evt.Modulo = rem.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("EthDepositReceived {\n")
	Info.Printf("\tDepositTime: %v\n", evt.DepositTime)
	Info.Printf("\tDepositId: %v\n", evt.DepositId)
	Info.Printf("\tRoundNum: %v\n", evt.RoundNum)
	Info.Printf("\tNumStakedNFTs: %v\n", evt.NumStakedNfts)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("\tAmountPerStaker: %v\n", evt.AmountPerStaker)
	Info.Printf("\tModulo: %v\n", evt.Modulo)
	Info.Printf("\tAccumModulo: %v\n", evt.AccumModulo)
	Info.Printf("}\n")
	if err := cgRepo.DeleteStakingEthDeposit(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertStakingEthDeposit(ctx, &evt)
}
func proc_nft_unstaked_rwalk_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGNftUnstakedRWalk
	var eth_evt IStakingWalletRandomWalkNftNftUnstaked

	Info.Printf("Processing NftUnstaked event id=%v, (block %v) txhash %v\n", elog.EvtId, elog.BlockNum, elog.TxHash)
	if !bytes.Equal(log.Address.Bytes(), staking_wallet_rwalk_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := staking_wallet_rwalk_abi.UnpackIntoInterface(&eth_evt, "NftUnstaked", log.Data)
	if err != nil {
		return fmt.Errorf("RWalk NftUnstaked (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.NftId = log.Topics[2].Big().Int64()
	evt.StakerAddress = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.NumStakedNfts = eth_evt.NumStakedNfts.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("NftUnstaked RWalk{\n")
	Info.Printf("\tStakeActionId: %v\n", evt.ActionId)
	Info.Printf("\tNftId: %v\n", evt.NftId)
	Info.Printf("\tNumStakedNfts: %v\n", evt.NumStakedNfts)
	Info.Printf("\tStakerAddress: %v\n", evt.StakerAddress)
	Info.Printf("}\n")

	if err := cgRepo.DeleteNftUnstakedRWalk(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertNftUnstakedRWalk(ctx, &evt)
}
func proc_nft_unstaked_cst_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGNftUnstakedCst
	var eth_evt IStakingWalletCosmicSignatureNftNftUnstaked

	Info.Printf("Processing NftUnstaked CST event id=%v, (block %v) txhash %v\n", elog.EvtId, elog.BlockNum, elog.TxHash)
	if !bytes.Equal(log.Address.Bytes(), staking_wallet_cst_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := staking_wallet_cst_abi.UnpackIntoInterface(&eth_evt, "NftUnstaked", log.Data)
	if err != nil {
		return fmt.Errorf("CST NftUnstaked (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.NftId = log.Topics[2].Big().Int64()
	evt.StakerAddress = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.NumStakedNfts = eth_evt.NumStakedNfts.Int64()
	evt.RewardAmount = eth_evt.RewardAmount.String()
	evt.RewardPerToken = eth_evt.RewardAmountPerStakedNft.String()
	evt.ActionCounter = eth_evt.ActionCounter.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("NftUnstaked Cst {\n")
	Info.Printf("\tStakeActionId: %v\n", evt.ActionId)
	Info.Printf("\tNftId: %v\n", evt.NftId)
	Info.Printf("\tNumStakedNfts: %v\n", evt.NumStakedNfts)
	Info.Printf("\tStakerAddress: %v\n", evt.StakerAddress)
	Info.Printf("\tRewardAmount (total): %v\n", evt.RewardAmount)
	Info.Printf("\tRewardPerToken: %v\n", evt.RewardPerToken)
	Info.Printf("\tActionCounter: %v\n", evt.ActionCounter)
	Info.Printf("}\n")

	if err := cgRepo.DeleteNftUnstakedCST(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertNftUnstakedCST(ctx, &evt)
}
