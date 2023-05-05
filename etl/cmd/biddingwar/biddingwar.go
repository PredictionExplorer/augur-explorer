package main

import (
	"os"
	"fmt"
	"math/big"
	"bytes"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/primitives/biddingwar"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
func build_list_of_inspected_events_layer1(cosmic_sig_aid int64) []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events= make([]InspectedEvent,0, 32)
	inspected_events = append(inspected_events,
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_prize_claim_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_bid_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_donation_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_donation_received_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_donation_sent_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_nft_donation_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_charity_updated[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_token_name_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_mint_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_raffle_deposit[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_raffle_nft_winner[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_raffle_nft_claimed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_transfer[:4]),
			ContractAid: cosmic_sig_aid,
		},
	)
	return inspected_events
}
func proc_prize_claim_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWPrizeClaimEvent
	var eth_evt CosmicGamePrizeClaimEvent

	Info.Printf("Processing PrizeClaim event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"PrizeClaimEvent",log.Data)
	if err != nil {
		Error.Printf("Event PrizeClaimEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.PrizeNum= log.Topics[1].Big().Int64()
	evt.WinnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PrizeClaimEvent {\n")
	Info.Printf("\tPrizeNum: %v\n",evt.PrizeNum)
	Info.Printf("\tWinner%v\n",evt.WinnerAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Insert_prize_claim_event(&evt)
}
func find_cosmic_token_transfer(bid_evtlog_id int64) string {
	// fetches the ERC20::Transfer event which has the id=evtlog-1 because it is
	//		inserted right before Bid event
	//		this function panics in case of failure because that would be an invalid database state
	ee := storagew.S.Get_event_log(bid_evtlog_id-1)	// ERC20 tansfer is always 1 less than Bid id
	var log types.Log
	err := rlp.DecodeBytes(ee.RlpLog,&log)
	if err!= nil {
		err_str := fmt.Sprintf("RLP Decode error at find_cosmic_signature_token_transfer(): %v",err)
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	var eth_evt ERC20Transfer
	err = erc20_abi.UnpackIntoInterface(&eth_evt,"Transfer",log.Data)
	if err != nil {
		err_str := fmt.Sprintf("Event Transfer decode error at find_cosmic_signature_token_transfer(): %v",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
		Info.Printf("%+v",log)
		Error.Printf("%+v",log)
		os.Exit(1)
	}
	return eth_evt.Value.String()
}
func find_cosmic_token_721_transfer(bid_evtlog_id int64) int64 {
	// fetches the ERC721::Transfer event which has the id=evtlog-1 because it is
	//		inserted right before RaffleNFTClaim event
	//		this function panics in case of failure because that would be an invalid database state
	ee := storagew.S.Get_event_log(bid_evtlog_id-1)	// ERC20 tansfer is always 1 less than Bid id
	var log types.Log
	err := rlp.DecodeBytes(ee.RlpLog,&log)
	if err!= nil {
		err_str := fmt.Sprintf("RLP Decode error at find_cosmic_token_721_transfer(): %v",err)
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	var eth_evt ERC721Transfer
	err = erc721_abi.UnpackIntoInterface(&eth_evt,"Transfer",log.Data)
	if err != nil {
		err_str := fmt.Sprintf("Event Transfer decode error at find_cosmic_token_721_transfer(): %v",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
		Info.Printf("%+v",log)
		Error.Printf("%+v",log)
		os.Exit(1)
	}
	if eth_evt.TokenId != nil {
		Info.Printf("token id=%v\n",eth_evt.TokenId.Int64())
	}
	if len(log.Topics) < 3 {
		err_str := fmt.Sprintf("Event ERC721 Transfer doesn't have 4 topics")
		Error.Printf(err_str)
		Info.Printf(err_str)
		Info.Printf("%+v",log)
		Error.Printf("%+v",log)
		os.Exit(1)
	}
	return log.Topics[1].Big().Int64()
}
func proc_bid_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWBidEvent
	var eth_evt CosmicGameBidEvent

	Info.Printf("Processing Bid event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"BidEvent",log.Data)
	if err != nil {
		Error.Printf("Event BidEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.LastBidderAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.BidPrice = eth_evt.BidPrice.String()
	evt.RandomWalkTokenId = eth_evt.RandomWalkNFTId.Int64()
	evt.ERC20_Value = find_cosmic_token_transfer(evt.EvtId)
	evt.PrizeTime = eth_evt.PrizeTime.Int64()
	evt.Message = eth_evt.Message

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("BidEvent {\n")
	Info.Printf("\tLastBidder: %v\n",evt.LastBidderAddr)
	Info.Printf("\tBidPrice: %v\n",evt.BidPrice)
	Info.Printf("\tRandomWalkTokenId: %v\n",evt.RandomWalkTokenId)
	Info.Printf("\tPrizeTime: %v\n",evt.PrizeTime)
	Info.Printf("\tMessage: %v\n",evt.Message)
	Info.Printf("}\n")

	storagew.Insert_bid_event(&evt)
}
func proc_donation_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWDonationEvent
	var eth_evt CosmicGameDonationEvent

	Info.Printf("Processing DonationEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"DonationEvent",log.Data)
	if err != nil {
		Error.Printf("Event DonationEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonationEvent {\n")
	Info.Printf("\tDonor: %v\n",evt.DonorAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Insert_donation(&evt)
}
func proc_donation_received_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWDonationReceivedEvent
	var eth_evt CharityWalletDonationReceivedEvent

	Info.Printf("Processing DonationReceivedEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),charity_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := charity_wallet_abi.UnpackIntoInterface(&eth_evt,"DonationReceivedEvent",log.Data)
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

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonationReceivedEvent {\n")
	Info.Printf("\tDonor: %v\n",evt.DonorAddr)
	Info.Printf("\tAmount%v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Insert_donation_received(&evt)
}
func proc_donation_sent_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWDonationSentEvent
	var eth_evt CharityWalletDonationSentEvent

	Info.Printf("Processing DonationSentEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),charity_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := charity_wallet_abi.UnpackIntoInterface(&eth_evt,"DonationSentEvent",log.Data)
	if err != nil {
		Error.Printf("Event DonationSentEvent decode error: %v",err)
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
	Info.Printf("DonationSentEvent {\n")
	Info.Printf("\tCharity: %v\n",evt.CharityAddr)
	Info.Printf("\tAmount%v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Insert_donation_sent(&evt)
}
func get_token_uri(token_id int64,contract_addr common.Address) string {

	c,err := NewCosmicSignatureNFT(contract_addr,eclient) // we use cosmicsiangature because its ERC721
	if err != nil {
		err_str := fmt.Sprintf("Error at get_token_uri() during contract creation: %v",err)
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	var copts bind.CallOpts
	tok_uri,err := c.TokenURI(&copts,big.NewInt(token_id))
	if err != nil {
		err_str := fmt.Sprintf("Error at get_token_uri() during GetTokenURI() call: %v",err)
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	return tok_uri
}
func proc_nft_donation_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWNFTDonationEvent
	var eth_evt CosmicGameNFTDonationEvent

	Info.Printf("Processing NFTDonationEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"NFTDonationEvent",log.Data)
	if err != nil {
		Error.Printf("Event NFTDonationEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.TokenAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenId = eth_evt.TokenId.Int64()
	evt.BidId = storagew.Get_biddingwar_bid_by_evtlog_id(evt.EvtId-2)
	evt.NFTTokenURI = get_token_uri(evt.TokenId,common.HexToAddress(evt.TokenAddr))

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NFTDonationEvent {\n")
	Info.Printf("\tDonor: %v\n",evt.DonorAddr)
	Info.Printf("\tNFTAddress: %v\n",evt.TokenAddr)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tNFTTokenURI: %v\n",evt.NFTTokenURI)
	Info.Printf("}\n")

	storagew.Insert_nft_donation_event(&evt)
}
func proc_charity_updated_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWCharityUpdatedEvent
	var eth_evt CharityWalletCharityUpdatedEvent

	Info.Printf("Processing CharityUpdatedEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),charity_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := charity_wallet_abi.UnpackIntoInterface(&eth_evt,"CharityUpdatedEvent",log.Data)
	if err != nil {
		Error.Printf("Event CharityUpdatedEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCharityAddr = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CharityUpdatedEvent {\n")
	Info.Printf("\tNewCharity: %v\n",evt.NewCharityAddr)
	Info.Printf("}\n")

	storagew.Insert_charity_updated_event(&evt)
}
func proc_token_name_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWTokenNameEvent
	var eth_evt CosmicSignatureNFTTokenNameEvent

	Info.Printf("Processing TokenNameEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt,"TokenNameEvent",log.Data)
	if err != nil {
		Error.Printf("Event TokenNameEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = eth_evt.TokenId.Int64()
	evt.TokenName = eth_evt.NewName

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TokenNameEvent {\n")
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tTokenName: %v\n",evt.TokenName)
	Info.Printf("}\n")

	storagew.Insert_token_name_event(&evt)
}
func proc_mint_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWMintEvent
	var eth_evt CosmicSignatureNFTMintEvent

	Info.Printf("Processing MintEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_signature_abi.UnpackIntoInterface(&eth_evt,"MintEvent",log.Data)
	if err != nil {
		Error.Printf("Event MintEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = log.Topics[1].Big().Int64()
	evt.OwnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Seed = hex.EncodeToString(eth_evt.Seed[:])

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("MintEvent{\n")
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tOwner:%v\n",evt.OwnerAddr)
	Info.Printf("\tSeed: %v\n",evt.Seed)
	Info.Printf("}\n")

	storagew.Insert_mint_event(&evt)
}
func proc_raffle_deposit_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWRaffleDeposit
	var eth_evt RaffleWalletRaffleDepositEvent 

	Info.Printf("Processing RaffleDeposit event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),raffle_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := raffle_wallet_abi.UnpackIntoInterface(&eth_evt,"RaffleDepositEvent",log.Data)
	if err != nil {
		Error.Printf("Event RaffleDepositEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Round = log.Topics[2].Big().Int64()
	evt.DepositId = eth_evt.DepositId.Int64()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RaffleDepositEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tDepositId: %v\n",evt.DepositId)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Insert_raffle_deposit(&evt)
}
func proc_raffle_nft_winner_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWRaffleNFTWinner
	var eth_evt CosmicGameRaffleNFTWinnerEvent

	Info.Printf("Processing RaffleNFTWinner event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"RaffleNFTWinnerEvent",log.Data)
	if err != nil {
		Error.Printf("Event RaffleNFTWinnerEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Round = log.Topics[2].Big().Int64()
	evt.WinnerIndex= eth_evt.WinnerIndex.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RaffleNFTWinnerEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tWinnerIndex: %v\n",evt.WinnerIndex)
	Info.Printf("}\n")

	storagew.Insert_raffle_nft_winner(&evt)
}
func proc_raffle_nft_claimed_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWRaffleNFTClaimed
	var eth_evt CosmicGameRaffleNFTWinnerEvent

	Info.Printf("Processing RaffleNFTClaimed event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"RaffleNFTClaimedEvent",log.Data)
	if err != nil {
		Error.Printf("Event RaffleNFTClaimedEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.TokenId = find_cosmic_token_721_transfer(evt.EvtId)

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RaffleNFTClaimedEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tTokenId: %v\n",evt.TokenId);
	Info.Printf("}\n")

	storagew.Insert_raffle_nft_claimed(&evt)
}
func proc_cosmic_sig_transfer_event(log *types.Log,elog *EthereumEventLog) {

	var evt BWERC721Transfer

	if !bytes.Equal(log.Address.Bytes(),cosmic_signature_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
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

    storagew.Insert_token_transfer_event(&evt)
}

func select_event_and_process(log *types.Log,evtlog *EthereumEventLog) {

	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_prize_claim_event) {
		proc_prize_claim_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_bid_event) {
		proc_bid_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_donation_event) {
		proc_donation_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_donation_received_event) {
		proc_donation_received_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_donation_sent_event) {
		proc_donation_sent_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_nft_donation_event) {
		proc_nft_donation_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_charity_updated) {
		proc_charity_updated_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_token_name_event) {
		proc_token_name_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_mint_event) {
		proc_mint_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_raffle_deposit) {
		proc_raffle_deposit_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_raffle_nft_winner) {
		proc_raffle_nft_winner_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_raffle_nft_claimed) {
		proc_raffle_nft_claimed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer) {
		proc_cosmic_sig_transfer_event(log,evtlog)
	}
}
func process_single_event(evt_id int64) error {

	evtlog := storagew.S.Get_event_log(evt_id)
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
