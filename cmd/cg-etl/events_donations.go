// Donation events: ETH/ERC20/NFT donations, donated token/NFT claims and
// charity wallet traffic (donations received/sent, funds transferred to charity).
package main

import (
	"os"
	"fmt"
	"math/big"
	"bytes"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)
func proc_donation_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGDonationEvent
	var eth_evt CosmicSignatureGameEthDonated

	Info.Printf("Processing DonationEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"EthDonated",log.Data)
	if err != nil {
		Error.Printf("Event EthDonaed decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Amount = eth_evt.Amount.String()
	evt.RoundNum = log.Topics[1].Big().Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonationEvent {\n")
	Info.Printf("\tDonor: %v\n",evt.DonorAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tRound: %v\n",evt.RoundNum)
	Info.Printf("}\n")

	storagew.Delete_donation_event(evt.EvtId)
	storagew.Insert_donation_event(&evt)
}
func get_donation_data(record_id int64) (string,error) {

	cosmic_game_ctrct,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err != nil {
		return "",err
	}
	var copts bind.CallOpts
	dinfo_rec,err := cosmic_game_ctrct.EthDonationWithInfoRecords(&copts,big.NewInt(record_id))
	if err != nil {
		return "",err
	}
	return dinfo_rec.Data,err
}
func proc_donation_with_info_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGDonationWithInfoEvent
	var eth_evt CosmicSignatureGameEthDonatedWithInfo

	Info.Printf("Processing DonationWithInfoEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"EthDonatedWithInfo",log.Data)
	if err != nil {
		Error.Printf("Event DonationWithInfoEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.RecordId = log.Topics[3].Big().Int64()
	evt.Amount = eth_evt.Amount.String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	data_json,err := get_donation_data(evt.RecordId)
	if err != nil {
		Info.Printf("Failure to fetch donation info record: %v\n",err.Error())
		Error.Printf("Failure to fetch donation info record: %v\n",err.Error())
		os.Exit(1)
	}

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonationWithInfoEvent {\n")
	Info.Printf("\tDonor: %v\n",evt.DonorAddr)
	Info.Printf("\tRecordId: %v\n",evt.RecordId)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tDaa JSON: %v\n",data_json)
	Info.Printf("}\n")

	storagew.Delete_donation_with_info_event(evt.EvtId)
	storagew.Insert_donation_with_info_event(&evt)
	storagew.Insert_donation_wi_data_json(evt.RecordId,data_json);
}
func proc_donation_received_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGDonationReceivedEvent
	var eth_evt CharityWalletDonationReceived

	Info.Printf("Processing DonationReceivedEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),charity_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := charity_wallet_abi.UnpackIntoInterface(&eth_evt,"DonationReceived",log.Data)
	if err != nil {
		Error.Printf("Event DonationReceivedEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()
	evt.RoundNum = find_prize_num(evt.TxId)

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonationReceivedEvent {\n")
	Info.Printf("\tDonor: %v\n",evt.DonorAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tRoundNum: %v\n",evt.RoundNum)
	Info.Printf("}\n")

	storagew.Delete_donation_received(evt.EvtId)
	storagew.Insert_donation_received(&evt)
}
func proc_donation_sent_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGDonationSentEvent
	var eth_evt CharityWalletFundsTransferredToCharity

	Info.Printf("Processing DonationSentEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),charity_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := charity_wallet_abi.UnpackIntoInterface(&eth_evt,"FundsTransferredToCharity",log.Data)
	if err != nil {
		Error.Printf("Event FundsTransferredToCharity decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.CharityAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("FundsTransferredToCharity{\n")
	Info.Printf("\tCharity: %v\n",evt.CharityAddr)
	Info.Printf("\tAmount%v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Delete_donation_sent(evt.EvtId)
	storagew.Insert_donation_sent(&evt)
}
func get_token_uri(token_id int64,contract_addr common.Address) string {

	c,err := NewCosmicSignatureNft(contract_addr,eclient) // we use cosmicsiangature because its ERC721
	if err != nil {
		err_str := fmt.Sprintf("Error at get_token_uri() during contract creation: %v",err)
		Info.Print(err_str)
		Error.Print(err_str)
		os.Exit(1)
	}
	var copts bind.CallOpts
	tok_uri,err := c.TokenURI(&copts,big.NewInt(token_id))
	if err != nil {
		err_str := fmt.Sprintf("Error at get_token_uri() during GetTokenURI() call: %v",err)
		Info.Print(err_str)
		Error.Print(err_str)
		return ""
	}
	return tok_uri
}
func proc_erc20_donated_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGERC20DonationEvent
	var eth_evt IPrizesWalletTokenDonated

	Info.Printf("Processing TokenDonated event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt,"TokenDonated",log.Data)
	if err != nil {
		Error.Printf("Event TokenDonated decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.DonorAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenAddr = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.Amount = eth_evt.Amount.String()
	evt.BidId = storagew.Get_bid_id_by_evtlog(evt.EvtId-1)
	if evt.BidId == -1 {	// if BidId = -1 , it could be that EvtId - 1 falls on Approval event, so Bid event will be EvtId - 2
		evt.BidId = storagew.Get_bid_id_by_evtlog(evt.EvtId-2)
	}

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TokenDonated{\n")
	Info.Printf("\tRoundNum: %v\n",evt.RoundNum)
	Info.Printf("\tDonor: %v\n",evt.DonorAddr)

	Info.Printf("\tTokenAddr: %v\n",evt.TokenAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount);
	Info.Printf("\tLinking to bid id: %v\n",evt.BidId)
	Info.Printf("}\n")

	storagew.Delete_erc20_donated_event(evt.EvtId)
	storagew.Insert_erc20_donated_event(&evt)
}
func proc_nft_donation_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNFTDonationEvent
	var eth_evt IPrizesWalletNftDonated 

	Info.Printf("Processing NftDonated event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt,"NftDonated",log.Data)
	if err != nil {
		Error.Printf("Event NFTDonationEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenAddr = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.TokenId = eth_evt.NftId.Int64()
	evt.BidId = storagew.Get_cosmic_game_bid_by_evtlog_id(evt.EvtId-1)
	evt.NFTTokenURI = get_token_uri(evt.TokenId,common.HexToAddress(evt.TokenAddr))
	evt.Index = eth_evt.Index.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NFTDonationEvent {\n")
	Info.Printf("\tDonor: %v\n",evt.DonorAddr)
	Info.Printf("\tNFTAddress: %v\n",evt.TokenAddr)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tNFTTokenURI: %v\n",evt.NFTTokenURI)
	Info.Printf("}\n")

	storagew.Delete_nft_donation_event(evt.EvtId)
	storagew.Insert_nft_donation_event(&evt)
}
func proc_donated_token_claimed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGDonatedTokenClaimed
	var eth_evt PrizesWalletDonatedTokenClaimed 

	Info.Printf("Processing DonatedTokenClaimed event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt,"DonatedTokenClaimed",log.Data)
	if err != nil {
		Error.Printf("Event DonatedTokenClaimedEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenAddr = eth_evt.TokenAddress.String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.Amount = eth_evt.Amount.String()
	evt.BeneficiaryAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenAddr = common.BytesToAddress(log.Topics[3][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonatedTokenClaimedEvent{\n")
	Info.Printf("\tRoundNum: %v\n",evt.RoundNum)
	Info.Printf("\tBeneficiary: %v\n",evt.BeneficiaryAddr)

	Info.Printf("\tTokenAddr: %v\n",evt.TokenAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount);
	Info.Printf("}\n")

	storagew.Delete_donated_token_claimed(evt.EvtId)
	storagew.Insert_donated_token_claimed(&evt)
}
func proc_donated_nft_claimed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGDonatedNFTClaimed
	var eth_evt PrizesWalletDonatedNftClaimed 

	Info.Printf("Processing DonatedNFTClaimed event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt,"DonatedNftClaimed",log.Data)
	if err != nil {
		Error.Printf("Event DonatedNFTClaimedEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.BeneficiaryAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenAddr = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.TokenId = eth_evt.NftId.String()
	evt.Index = eth_evt.Index.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonatedNFTClaimedEvent{\n")
	Info.Printf("\tRoundNum: %v\n",evt.RoundNum)
	Info.Printf("\tIndex: %v\n",evt.Index)
	Info.Printf("\tBeneficiary: %v\n",evt.BeneficiaryAddr)

	Info.Printf("\tTokenAddr: %v\n",evt.TokenAddr)
	Info.Printf("\tTokenId: %v\n",evt.TokenId);
	Info.Printf("}\n")

	storagew.Delete_donated_nft_claimed(evt.EvtId)
	storagew.Insert_donated_nft_claimed(&evt)
}
func proc_funds_transferred_to_charity_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGFundsToCharity
	var eth_evt CosmicSignatureGameFundsTransferredToCharity

	if !bytes.Equal(log.Address.Bytes(),marketing_wallet_addr.Bytes()) {
		return
	}
	Info.Printf("Processing FundsTransferredToCharity event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"FundsTransferredToCharity",log.Data)
	if err != nil {
		Error.Printf("Event FundsTransferredToCharity decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.CharityAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount= eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("FundsTransferredToCharity{\n")
	Info.Printf("\tCharityAddress: %v\n",evt.CharityAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Delete_funds_transferred_to_charity_event(evt.EvtId)
    storagew.Insert_funds_transferred_to_charity_event(&evt)
}
