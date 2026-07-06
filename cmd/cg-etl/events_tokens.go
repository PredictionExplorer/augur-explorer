// Token events: CosmicSignature NFT mint/name/transfer, CosmicToken (ERC20)
// transfers and MarketingWallet reward payments.
package main

import (
	"os"
	"fmt"
	"bytes"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)
func proc_token_name_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGTokenNameEvent
	var eth_evt ICosmicSignatureNftNftNameChanged

	Info.Printf("Processing TokenNameEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt,"NftNameChanged",log.Data)
	if err != nil {
		Error.Printf("Event TokenNameEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = log.Topics[1].Big().Int64()
	evt.TokenName = eth_evt.NftName

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TokenNameEvent {\n")
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tTokenName: %v\n",evt.TokenName)
	Info.Printf("}\n")

	storagew.Delete_token_name(evt.EvtId)
	storagew.Insert_token_name_event(&evt)
}
func proc_mint_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGMintEvent
	var eth_evt CosmicSignatureNftNftMinted

	Info.Printf("Processing MintEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt,"NftMinted",log.Data)
	if err != nil {
		Error.Printf("Event MintEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
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

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("MintEvent{\n")
	Info.Printf("\tRoundNum: %v\n",evt.RoundNum)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tOwner:%v\n",evt.OwnerAddr)
	Info.Printf("\tSeed: %v\n",evt.Seed)
	Info.Printf("}\n")

	storagew.Delete_mint_event(evt.EvtId)
	storagew.Insert_mint_event(&evt)
	/*Temporarily disabled
	cmd_str := fmt.Sprintf("%v/%v %v %v",os.Getenv("HOME"),IMGGEN_PATH,evt.TokenId,evt.Seed)
	Info.Printf("Executing %v\n",cmd_str)
	cmd := exec.Command(cmd_str)
	err = cmd.Run()
	if err != nil {
		Info.Printf("Error executing image generation: %v\n",err)
		Error.Printf("Error executing image generation: %v\n",err)
	}
	*/
}
func proc_marketing_reward_paid_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGMarketingRewardPaid
	var eth_evt MarketingWalletRewardPaid

	Info.Printf("Processing MarketingWallet RewardPaid event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),marketing_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := marketing_wallet_abi.UnpackIntoInterface(&eth_evt,"RewardPaid",log.Data)
	if err != nil {
		Error.Printf("Event RewardSentEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Marketer = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("Marketing RewardPaid{\n")
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tMarketer: %v\n",evt.Marketer)
	Info.Printf("}\n")

	storagew.Delete_marketing_reward_paid_event(evt.EvtId)
	storagew.Insert_marketing_reward_paid_event(&evt)
}
func proc_cosmic_signature_transfer_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGERC721Transfer

	if !bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		return
	}
	Info.Printf("Processing Token Transfer event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.From = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.To = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenId = log.Topics[3].Big().Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("Transfer {\n")
	Info.Printf("\tFrom: %v\n",evt.From)
	Info.Printf("\tTo: %v\n",evt.To)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("}\n")

	storagew.Delete_cosmic_signature_transfer_event(evt.EvtId)
    storagew.Insert_cosmic_signature_transfer_event(&evt)
}
func proc_cosmic_token_transfer_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGERC20Transfer 
	var eth_evt ERC20Transfer

	if !bytes.Equal(log.Address.Bytes(),cosmic_token_addr.Bytes()) {
		return
	}
	Info.Printf("Processing ERC20 Transfer event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := erc20_abi.UnpackIntoInterface(&eth_evt,"Transfer",log.Data)
	if err != nil {
		Error.Printf("Event ERC20Transfer decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.From = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.To = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Value = eth_evt.Value.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("Transfer {\n")
	Info.Printf("\tFrom: %v\n",evt.From)
	Info.Printf("\tTo: %v\n",evt.To)
	Info.Printf("\tValue: %v\n",evt.Value)
	Info.Printf("}\n")

	storagew.Delete_cosmic_token_transfer_event(evt.EvtId)
    storagew.Insert_cosmic_token_transfer_event(&evt)
}
func proc_transfer_event_common(log *types.Log,elog *EthereumEventLog) {

	if bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		proc_cosmic_signature_transfer_event(log,elog)
	}
	if bytes.Equal(log.Address.Bytes(),cosmic_token_addr.Bytes()) {
		proc_cosmic_token_transfer_event(log,elog)
	}

}
