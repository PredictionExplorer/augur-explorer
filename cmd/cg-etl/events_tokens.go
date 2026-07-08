// Token events: CosmicSignature NFT mint/name/transfer, CosmicToken (ERC20)
// transfers and MarketingWallet reward payments.
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

func proc_token_name_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGTokenNameEvent
	var eth_evt ICosmicSignatureNftNftNameChanged

	Info.Printf("Processing TokenNameEvent event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_signature_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt, "NftNameChanged", log.Data)
	if err != nil {
		return fmt.Errorf("NftNameChanged (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = log.Topics[1].Big().Int64()
	evt.TokenName = eth_evt.NftName

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("TokenNameEvent {\n")
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("\tTokenName: %v\n", evt.TokenName)
	Info.Printf("}\n")

	if err := cgRepo.DeleteTokenName(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertTokenName(ctx, &evt)
}
func proc_mint_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGMintEvent
	var eth_evt CosmicSignatureNftNftMinted

	Info.Printf("Processing MintEvent event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_signature_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt, "NftMinted", log.Data)
	if err != nil {
		return fmt.Errorf("NftMinted (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = log.Topics[3].Big().Int64()
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.OwnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	// NftSeed is uint256; big.Int.Bytes() drops leading zero bytes, so format
	// with a fixed 64-hex-char width (32 bytes) to preserve leading zeros.
	evt.Seed = fmt.Sprintf("%064x", eth_evt.NftSeed)

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("MintEvent{\n")
	Info.Printf("\tRoundNum: %v\n", evt.RoundNum)
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("\tOwner:%v\n", evt.OwnerAddr)
	Info.Printf("\tSeed: %v\n", evt.Seed)
	Info.Printf("}\n")

	if err := cgRepo.DeleteMint(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertMint(ctx, &evt)
}
func proc_marketing_reward_paid_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGMarketingRewardPaid
	var eth_evt MarketingWalletRewardPaid

	Info.Printf("Processing MarketingWallet RewardPaid event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), marketing_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := marketing_wallet_abi.UnpackIntoInterface(&eth_evt, "RewardPaid", log.Data)
	if err != nil {
		return fmt.Errorf("RewardPaid (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Marketer = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("Marketing RewardPaid{\n")
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("\tMarketer: %v\n", evt.Marketer)
	Info.Printf("}\n")

	if err := cgRepo.DeleteMarketingRewardPaid(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertMarketingRewardPaid(ctx, &evt)
}
func proc_cosmic_signature_transfer_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGERC721Transfer

	if !bytes.Equal(log.Address.Bytes(), cosmic_signature_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing Token Transfer event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.From = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.To = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenId = log.Topics[3].Big().Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("Transfer {\n")
	Info.Printf("\tFrom: %v\n", evt.From)
	Info.Printf("\tTo: %v\n", evt.To)
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("}\n")

	if err := cgRepo.DeleteCosmicSignatureTransfer(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertCosmicSignatureTransfer(ctx, &evt)
}
func proc_cosmic_token_transfer_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGERC20Transfer
	var eth_evt ERC20Transfer

	if !bytes.Equal(log.Address.Bytes(), cosmic_token_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing ERC20 Transfer event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := erc20_abi.UnpackIntoInterface(&eth_evt, "Transfer", log.Data)
	if err != nil {
		return fmt.Errorf("ERC20 Transfer (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.From = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.To = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Value = eth_evt.Value.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("Transfer {\n")
	Info.Printf("\tFrom: %v\n", evt.From)
	Info.Printf("\tTo: %v\n", evt.To)
	Info.Printf("\tValue: %v\n", evt.Value)
	Info.Printf("}\n")

	if err := cgRepo.DeleteCosmicTokenTransfer(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertCosmicTokenTransfer(ctx, &evt)
}
func proc_transfer_event_common(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	if bytes.Equal(log.Address.Bytes(), cosmic_signature_addr.Bytes()) {
		if err := proc_cosmic_signature_transfer_event(ctx, log, elog); err != nil {
			return err
		}
	}
	if bytes.Equal(log.Address.Bytes(), cosmic_token_addr.Bytes()) {
		if err := proc_cosmic_token_transfer_event(ctx, log, elog); err != nil {
			return err
		}
	}
	return nil
}
