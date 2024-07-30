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
	. "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
func build_list_of_inspected_events_layer1(cosmic_sig_aid int64) []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events= make([]InspectedEvent,0, 32)
	inspected_events = append(inspected_events,
		// this list matches the order of main.go event variables in `var` declaration
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
			Signature: hex.EncodeToString(evt_donation_with_info_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_nft_donation_event[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_raffle_nft_winner[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_endurance_winner[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_stellar_winner[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_donated_nft_claimed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_charity_percentage_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_prize_percentage_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_raffle_percentage_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_staking_percentage_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_num_raffle_eth_winners_bidding_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_num_raffle_nft_winners_bidding_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_num_raffle_nft_winners_staking_rwalk_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_charity_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_rwalk_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_raffle_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_staking_wallet_cst_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_staking_wallet_rwalk_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_marketing_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_costok_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_cossig_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_blogic_address_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_time_increase_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_timeout_claimprize_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_price_increase_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_nanoseconds_extra_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_initial_seconds_until_prize_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_initial_bid_amount_fraction_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_activation_time_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_ethcst_bid_ratio_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_round_start_auction_length_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_system_mode_changed[:4]),
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
			Signature: hex.EncodeToString(evt_raffle_withdrawal[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_transfer[:4]),
			ContractAid: cosmic_sig_aid,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_stake_action[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_unstake_action[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_claim_reward[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_eth_deposit[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_marketing_reward_sent[:4]),
			ContractAid: 0,
		},
	)
	return inspected_events
}
func proc_prize_claim_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGPrizeClaimEvent
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
	evt.TokenId = find_cosmic_token_721_mint_event(cosmic_sig_aid,evt.TxId,evt.EvtId)
	evt.DonationEvtId = storagew.Get_donation_received_evt_id(evt.TxId,evt.EvtId,hex.EncodeToString(evt_donation_received_event[:4]))
	if evt.DonationEvtId == 0 {
		Error.Printf("Failed to fetch donation received event id\n")
		Info.Printf("Failed to fetch donation received event id\n")
		os.Exit(1)
	}

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PrizeClaimEvent {\n")
	Info.Printf("\tPrizeNum: %v\n",evt.PrizeNum)
	Info.Printf("\tWinner%v\n",evt.WinnerAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tDonationEvtId: %v\n",evt.DonationEvtId)
	Info.Printf("}\n")

	storagew.Delete_prize_claim_event(evt.EvtId)
	storagew.Insert_prize_claim_event(&evt)
}
func find_cosmic_token_transfer(bid_evtlog_id int64) string {
	// fetches the ERC20::Transfer event which has the id=evtlog-1 because it is
	//		inserted right before Bid event
	//		this function panics in case of failure because that would be an invalid database state
	ee := storagew.S.Get_event_log(bid_evtlog_id-2)	// ERC20 tansfer is always 2 less than the bid (-1 is for marketing reward but -2 is the bid reward)
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
	var eth_evt IERC721Transfer
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
func find_cosmic_token_721_mint_event(contract_aid,tx_id,claim_prize_evtlog_id int64) int64 {

	mint_evt_list := storagew.S.Get_specific_event_logs_by_tx_backwards_from_id(tx_id,contract_aid,claim_prize_evtlog_id,hex.EncodeToString(evt_mint_event[0:4]))
	if len(mint_evt_list) == 0 {
		err_str := fmt.Sprintf("find_cosmic_token_721_mint_event() couldn't find corresponding MintEvent()")
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	mint_location := len(mint_evt_list)-1
	var log types.Log
	err := rlp.DecodeBytes(mint_evt_list[mint_location],&log)
	if err!= nil {
		err_str := fmt.Sprintf("RLP Decode error at find_cosmic_token_721_mint_event(): %v",err)
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	if len(log.Topics) < 2 {
		err_str := fmt.Sprintf("Event ERC721 MintEvent doesn't have 3 topics")
		Error.Printf(err_str)
		Info.Printf(err_str)
		Info.Printf("%+v",log)
		Error.Printf("%+v",log)
		os.Exit(1)
	}
	return log.Topics[1].Big().Int64()
}
func find_prize_num(tx_id int64) int64 {

	evt_list,err := storagew.S.Get_events_by_sig_and_tx_id(tx_id,hex.EncodeToString(evt_prize_claim_event[0:4]))
	if err != nil {

		err_str := fmt.Sprintf("find_prize_num()() error: %v",err)
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	if len(evt_list) == 0 {
		return -1
	}
	if len(evt_list) != 1 {
		err_str := fmt.Sprintf("find_prize_num() there is more than 1 PrizeClaim in this transaction()")
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	var log types.Log
	err = rlp.DecodeBytes(evt_list[0].RlpLog,&log)
	if err!= nil {
		err_str := fmt.Sprintf("RLP Decode error at find_prize_num(): %v",err)
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	if len(log.Topics) < 2 {
		err_str := fmt.Sprintf("Event PrizeClaimEvent doesn't have 3 topics")
		Error.Printf(err_str)
		Info.Printf(err_str)
		Info.Printf("%+v",log)
		Error.Printf("%+v",log)
		os.Exit(1)
	}
	return log.Topics[1].Big().Int64()
}
func proc_bid_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGBidEvent
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
	evt.RoundNum = log.Topics[2].Big().Int64()
	evt.BidPrice = eth_evt.BidPrice.String()
	evt.BidType = 0; // ETH
	evt.RandomWalkTokenId = eth_evt.RandomWalkNFTId.Int64()
	evt.ERC20_Value = find_cosmic_token_transfer(evt.EvtId)
	evt.NumCSTTokens = eth_evt.NumCSTTokens.String()
	if evt.RandomWalkTokenId > -1 {
		evt.BidType = 1;	// RandomWalk	
	} else {
		if evt.NumCSTTokens != "-1" { evt.BidType = 2; } // Cosmic Signature Token (ERC20) bid
	}
	evt.PrizeTime = eth_evt.PrizeTime.Int64()
	evt.Message = eth_evt.Message

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("BidEvent {\n")
	Info.Printf("\tLastBidder: %v\n",evt.LastBidderAddr)
	Info.Printf("\tRoundNum: %v\n",evt.RoundNum)
	Info.Printf("\tBidPrice: %v\n",evt.BidPrice)
	Info.Printf("\tNumCSTTokens: %v\n",evt.NumCSTTokens);
	Info.Printf("\tRandomWalkTokenId: %v\n",evt.RandomWalkTokenId)
	Info.Printf("\tPrizeTime: %v\n",evt.PrizeTime)
	Info.Printf("\tMessage: %v\n",evt.Message)
	Info.Printf("}\n")

	storagew.Delete_bid(evt.EvtId)
	storagew.Insert_bid_event(&evt)
}
func proc_donation_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGDonationEvent
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

	storagew.Delete_donation_event(evt.EvtId)
	storagew.Insert_donation_event(&evt)
}
func get_donation_data(record_id int64) (string,error) {

	cosmic_game_ctrct,err := NewCosmicGame(cosmic_game_addr,eclient)
	if err != nil {
		return "",err
	}
	fmt.Printf("record id to query = %v\n",record_id)
	var copts bind.CallOpts
	dinfo_rec,err := cosmic_game_ctrct.DonationInfoRecords(&copts,big.NewInt(record_id))
	if err != nil {
		return "",err
	}
	fmt.Printf("donation data: \n%v\n",dinfo_rec);
	Info.Printf("donation data: \n%v\n",dinfo_rec);
	return dinfo_rec.Data,err
}
func proc_donation_with_info_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGDonationWithInfoEvent
	var eth_evt CosmicGameDonationWithInfoEvent

	Info.Printf("Processing DonationWithInfoEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"DonationWithInfoEvent",log.Data)
	if err != nil {
		Error.Printf("Event DonationWithInfoEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.RecordId = eth_evt.RecordId.Int64();
	evt.Amount = eth_evt.Amount.String()
	data_json,err := get_donation_data(evt.RecordId)
	fmt.Printf("data_json=%v, err=%v\n",data_json,err)
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
	Info.Printf("}\n")

	storagew.Delete_donation_with_info_event(evt.EvtId)
	storagew.Insert_donation_with_info_event(&evt)
	storagew.Insert_donation_wi_data_json(evt.RecordId,data_json);
}
func proc_donation_received_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGDonationReceivedEvent
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
	evt.RoundNum = find_prize_num(evt.TxId)

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonationReceivedEvent {\n")
	Info.Printf("\tDonor: %v\n",evt.DonorAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tPrizeNum: %v\n",evt.RoundNum)
	Info.Printf("}\n")

	storagew.Delete_donation_received(evt.EvtId)
	storagew.Insert_donation_received(&evt)
}
func proc_donation_sent_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGDonationSentEvent
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

	storagew.Delete_donation_sent(evt.EvtId)
	storagew.Insert_donation_sent(&evt)
}
func get_token_uri(token_id int64,contract_addr common.Address) string {

	c,err := NewCosmicSignature(contract_addr,eclient) // we use cosmicsiangature because its ERC721
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
		//os.Exit(1)
		return ""
	}
	return tok_uri
}
func proc_nft_donation_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNFTDonationEvent
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
	evt.RoundNum = log.Topics[3].Big().Int64()
	evt.TokenId = eth_evt.TokenId.Int64()
	evt.BidId = storagew.Get_cosmic_game_bid_by_evtlog_id(evt.EvtId-2)
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
func proc_charity_updated_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCharityUpdatedEvent
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

	storagew.Delete_charity_updated(evt.EvtId)
	storagew.Insert_charity_updated_event(&evt)
}
func proc_token_name_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGTokenNameEvent
	var eth_evt CosmicSignatureTokenNameEvent

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
	evt.TokenId = log.Topics[1].Big().Int64()
	evt.TokenName = eth_evt.NewName

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
	var eth_evt CosmicSignatureMintEvent

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
	evt.RoundNum = log.Topics[3].Big().Int64()
	evt.OwnerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Seed = hex.EncodeToString(eth_evt.Seed[:])

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
func proc_raffle_deposit_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRaffleDeposit
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

	prize_num := find_prize_num(elog.TxId)
	if prize_num == -1 {
		err_str := fmt.Sprintf("find_prize_num() couldn't find corresponding PrizeClaimEvent()")
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Round = prize_num
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RaffleDepositEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Delete_raffle_deposit(evt.EvtId)
	storagew.Insert_raffle_deposit(&evt)
}
func proc_raffle_withdrawal_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRaffleWithdrawal
	var eth_evt RaffleWalletRaffleWithdrawalEvent 

	Info.Printf("Processing RaffleWithdrawalevent id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),raffle_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := raffle_wallet_abi.UnpackIntoInterface(&eth_evt,"RaffleWithdrawalEvent",log.Data)
	if err != nil {
		Error.Printf("Event RaffleWithdrawalEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RaffleWithdrawalEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("}\n")

	storagew.Delete_raffle_withdrawal(evt.EvtId)
	storagew.Insert_raffle_withdrawal(&evt)
}
func proc_raffle_nft_winner_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRaffleNFTWinner
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
	evt.TokenId = log.Topics[3].Big().Int64()
	evt.WinnerIndex= eth_evt.WinnerIndex.Int64()
	evt.IsRandomWalk = eth_evt.IsRWalk
	evt.IsStaker = eth_evt.IsStaker

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RaffleNFTWinnerEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tWinnerIndex: %v\n",evt.WinnerIndex)
	Info.Printf("\tIsStaker: %v\n",evt.IsStaker);
	Info.Printf("\tIsRandomWalk: %v\n",evt.IsRandomWalk)
	Info.Printf("}\n")

	storagew.Delete_raffle_nft_winner(evt.EvtId)
	storagew.Insert_raffle_nft_winner(&evt)
}
func proc_endurance_winner_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGEnduranceWinner
	var eth_evt BusinessLogicEnduranceChampionWinnerEvent

	Info.Printf("Processing Endurance winner event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := blogic_abi.UnpackIntoInterface(&eth_evt,"EnduranceChampionWinnerEvent",log.Data)
	if err != nil {
		Error.Printf("Event EnduranceChampionWinnerEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Round = log.Topics[2].Big().Int64()
	evt.Erc721TokenId = log.Topics[3].Big().Int64()
	evt.Erc20Amount = eth_evt.Erc20TokenAmount.String()
	evt.WinnerIndex= eth_evt.WinnerIndex.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("EnduranceNFTWinnerEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tErc721TokenId: %v\n",evt.Erc721TokenId)
	Info.Printf("\tErc20Amount: %v\n",evt.Erc20Amount)
	Info.Printf("\tWinnerIndex: %v\n",evt.WinnerIndex)
	Info.Printf("}\n")

	storagew.Delete_endurance_winner(evt.EvtId)
	storagew.Insert_endurance_winner(&evt)
}
func proc_stellar_winner_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGStellarWinner
	var eth_evt BusinessLogicStellarSpenderWinnerEvent

	Info.Printf("Processing StellarSpender winner event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := blogic_abi.UnpackIntoInterface(&eth_evt,"StellarSpenderWinnerEvent",log.Data)
	if err != nil {
		Error.Printf("Event StellarSpender decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.WinnerAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Round = log.Topics[2].Big().Int64()
	evt.Erc721TokenId = log.Topics[3].Big().Int64()
	evt.Erc20Amount = eth_evt.Erc20TokenAmount.String()
	evt.WinnerIndex= eth_evt.WinnerIndex.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("StellarSpenderWinnerEvent{\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
	Info.Printf("\tRound:%v\n",evt.Round)
	Info.Printf("\tErc721TokenId: %v\n",evt.Erc721TokenId)
	Info.Printf("\tErc20TokenId: %v\n",evt.Erc20Amount)
	Info.Printf("\tWinnerIndex: %v\n",evt.WinnerIndex)
	Info.Printf("}\n")

	storagew.Delete_stellar_winner(evt.EvtId)
	storagew.Insert_stellar_winner(&evt)
}
func proc_donated_nft_claimed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGDonatedNFTClaimed
	var eth_evt CosmicGameDonatedNFTClaimedEvent

	Info.Printf("Processing DonatedNFTClaimed event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"DonatedNFTClaimedEvent",log.Data)
	if err != nil {
		Error.Printf("Event DonatedNFTClaimedEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenAddr = eth_evt.NftAddressdonatedNFTs.String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.TokenId = eth_evt.TokenId.String()
	evt.Index = eth_evt.Index.Int64()
	evt.WinnerAddr = eth_evt.Winner.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("DonatedNFTClaimedEvent{\n")
	Info.Printf("\tRoundNum: %v\n",evt.RoundNum)
	Info.Printf("\tIndex: %v\n",evt.Index)
	Info.Printf("\tWinner: %v\n",evt.WinnerAddr)

	Info.Printf("\tTokenAddr: %v\n",evt.TokenAddr)
	Info.Printf("\tTokenId: %v\n",evt.TokenId);
	Info.Printf("}\n")

	storagew.Delete_donated_nft_claimed(evt.EvtId)
	storagew.Insert_donated_nft_claimed(&evt)
}
func proc_stake_action_cst_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGStakeActionCST
	var eth_evt StakingWalletCSTStakeActionEvent

	Info.Printf("Processing StakeAction event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),staking_wallet_cst_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := staking_wallet_cst_abi.UnpackIntoInterface(&eth_evt,"StakeActionEvent",log.Data)
	if err != nil {
		Error.Printf("CST Event StakeAction decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.TokenId = log.Topics[2].Big().Int64()
	evt.Staker = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.TotalNfts = eth_evt.TotalNFTs.Int64()
	evt.UnstakeTime = eth_evt.UnstakeTime.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CST StakeActionEvent{\n")
	Info.Printf("\tActionId: %v\n",evt.ActionId)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tTotalNFTs: %v\n",evt.TotalNfts)
	Info.Printf("\tUnstakeTime: %v\n",evt.UnstakeTime)
	Info.Printf("\tStaker: %v\n",evt.Staker)
	Info.Printf("}\n")

	storagew.Delete_stake_action_cst_event(evt.EvtId)
	storagew.Insert_stake_action_cst_event(&evt)
}
func proc_unstake_action_cst_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGUnstakeActionCST
	var eth_evt StakingWalletCSTUnstakeActionEvent

	Info.Printf("Processing UnstakeAction event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	if !bytes.Equal(log.Address.Bytes(),staking_wallet_cst_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := staking_wallet_cst_abi.UnpackIntoInterface(&eth_evt,"UnstakeActionEvent",log.Data)
	if err != nil {
		Error.Printf("CST Event UnstakeAction decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.TokenId = log.Topics[2].Big().Int64()
	evt.Staker = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.TotalNfts = eth_evt.TotalNFTs.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CST UnstakeActionEvent{\n")
	Info.Printf("\tActionId: %v\n",evt.ActionId)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tTotalNFTs: %v\n",evt.TotalNfts)
	Info.Printf("\tStaker: %v\n",evt.Staker)
	Info.Printf("}\n")

	storagew.Delete_unstake_action_cst_event(evt.EvtId)
	storagew.Insert_unstake_action_cst_event(&evt)
}
func proc_eth_deposit_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGEthDeposit
	var eth_evt StakingWalletCSTEthDepositEvent

	Info.Printf("Processing EthDeposit event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),staking_wallet_cst_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := staking_wallet_cst_abi.UnpackIntoInterface(&eth_evt,"EthDepositEvent",log.Data)
	if err != nil {
		Error.Printf("Event EthDepositEvent decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DepositTime = log.Topics[1].Big().Int64()
	evt.DepositNum = eth_evt.DepositNum.Int64()
	evt.NumStakedNfts = eth_evt.NumStakedNFTs.Int64()
	evt.Amount = eth_evt.Amount.String()
	evt.AccumModulo = eth_evt.Modulo.String()
	evt.RoundNum = find_prize_num(evt.TxId)
	if evt.RoundNum == -1 {
		Error.Printf("Failed to gather round_num variable")
		Info.Printf("Failed to gather round_num variable")
		os.Exit(1)
	}
	divres:=big.NewInt(0)
	rem:=big.NewInt(0)
	divres.QuoRem(eth_evt.Amount,eth_evt.NumStakedNFTs,rem);
	evt.AmountPerStaker = divres.String()
	evt.Modulo = rem.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("EthDepositEvent{\n")
	Info.Printf("\tDepositTime: %v\n",evt.DepositTime)
	Info.Printf("\tDepositNum: %v\n",evt.DepositNum)
	Info.Printf("\tRoundNum: %v\n",evt.RoundNum)
	Info.Printf("\tNumStakedNFTs: %v\n",evt.NumStakedNfts)
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tAmountPerStaker: %v\n",evt.AmountPerStaker)
	Info.Printf("\tModulo: %v\n",evt.Modulo)
	Info.Printf("\tAccumModulo: %v\n",evt.AccumModulo)
	Info.Printf("}\n")
	storagew.Delete_eth_deposit_event(evt.EvtId)
	storagew.Insert_eth_deposit_event(&evt)
}
func proc_claim_reward_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGClaimReward
	var eth_evt StakingWalletCSTClaimRewardEvent

	Info.Printf("Processing ClaimReward event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),staking_wallet_cst_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := staking_wallet_cst_abi.UnpackIntoInterface(&eth_evt,"ClaimRewardEvent",log.Data)
	if err != nil {
		Error.Printf("Event ClaimReward decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.DepositId = log.Topics[2].Big().Int64()
	evt.Staker = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.Reward = eth_evt.Reward.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("ClaimReward{\n")
	Info.Printf("\tActionId: %v\n",evt.ActionId)
	Info.Printf("\tDepositId: %v\n",evt.DepositId)
	Info.Printf("\tReward: %v\n",evt.Reward)
	Info.Printf("\tStaker: %v\n",evt.Staker)
	Info.Printf("}\n")

	storagew.Delete_claim_reward_event(evt.EvtId)
	storagew.Insert_claim_reward_event(&evt)
}
func proc_stake_action_rwalk_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGStakeActionRWalk
	var eth_evt StakingWalletRWalkStakeActionEvent

	Info.Printf("Processing StakeAction event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),staking_wallet_rwalk_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := staking_wallet_rwalk_abi.UnpackIntoInterface(&eth_evt,"StakeActionEvent",log.Data)
	if err != nil {
		Error.Printf("RWalk Event StakeAction decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.TokenId = log.Topics[2].Big().Int64()
	evt.Staker = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.TotalNfts = eth_evt.TotalNFTs.Int64()
	evt.UnstakeTime = eth_evt.UnstakeTime.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RWalk StakeActionEvent{\n")
	Info.Printf("\tActionId: %v\n",evt.ActionId)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tTotalNFTs: %v\n",evt.TotalNfts)
	Info.Printf("\tUnstakeTime: %v\n",evt.UnstakeTime)
	Info.Printf("\tStaker: %v\n",evt.Staker)
	Info.Printf("}\n")

	storagew.Delete_stake_action_rwalk_event(evt.EvtId)
	storagew.Insert_stake_action_rwalk_event(&evt)
}
func proc_unstake_action_rwalk_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGUnstakeActionRWalk
	var eth_evt StakingWalletRWalkUnstakeActionEvent

	Info.Printf("Processing UnstakeAction event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	if !bytes.Equal(log.Address.Bytes(),staking_wallet_rwalk_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := staking_wallet_rwalk_abi.UnpackIntoInterface(&eth_evt,"UnstakeActionEvent",log.Data)
	if err != nil {
		Error.Printf("RWalk Event UnstakeAction decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.TokenId = log.Topics[2].Big().Int64()
	evt.Staker = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.TotalNfts = eth_evt.TotalNFTs.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RWalk UnstakeActionEvent{\n")
	Info.Printf("\tActionId: %v\n",evt.ActionId)
	Info.Printf("\tTokenId: %v\n",evt.TokenId)
	Info.Printf("\tTotalNFTs: %v\n",evt.TotalNfts)
	Info.Printf("\tStaker: %v\n",evt.Staker)
	Info.Printf("}\n")

	storagew.Delete_unstake_action_rwalk_event(evt.EvtId)
	storagew.Insert_unstake_action_rwalk_event(&evt)
}
func select_and_proc_stake_action_event(log *types.Log,elog *EthereumEventLog) {

	Info.Printf("Processing StakeAction event id=%v, txhash %v (selection)\n",elog.EvtId,elog.TxHash)

	if bytes.Equal(log.Address.Bytes(),staking_wallet_cst_addr.Bytes()) {
		proc_stake_action_cst_event(log,elog)
	} else {
		if bytes.Equal(log.Address.Bytes(),staking_wallet_rwalk_addr.Bytes()) {
			proc_stake_action_rwalk_event(log,elog)
		} else {
			Info.Printf("StakeAction event doesn't belong to our address %v\n",log.Address.String())
		}
	}
}
func select_and_proc_unstake_action_event(log *types.Log,elog *EthereumEventLog) {

	Info.Printf("Processing UnstakeAction event id=%v, txhash %v (selection)\n",elog.EvtId,elog.TxHash)

	if bytes.Equal(log.Address.Bytes(),staking_wallet_cst_addr.Bytes()) {
		proc_unstake_action_cst_event(log,elog)
	} else {
		if bytes.Equal(log.Address.Bytes(),staking_wallet_rwalk_addr.Bytes()) {
			proc_unstake_action_rwalk_event(log,elog)
		} else {
			Info.Printf("StakeAction event doesn't belong to our address %v\n",log.Address.String())
		}
	}
}
func proc_marketing_reward_sent_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGMarketingRewardSent
	var eth_evt MarketingWalletRewardSentEvent

	Info.Printf("Processing MarketingWallet RewardSent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),marketing_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := marketing_wallet_abi.UnpackIntoInterface(&eth_evt,"RewardSentEvent",log.Data)
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
	Info.Printf("Marketing RwardSentEvent{\n")
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tMarketer: %v\n",evt.Marketer)
	Info.Printf("}\n")

	storagew.Delete_marketing_reward_sent_event(evt.EvtId)
	storagew.Insert_marketing_reward_sent_event(&evt)
}
func proc_cosmic_signature_transfer_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGERC721Transfer

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

	storagew.Delete_cosmic_signature_transfer_event(evt.EvtId)
    storagew.Insert_cosmic_signature_transfer_event(&evt)
}
func proc_cosmic_token_transfer_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGERC20Transfer 
	var eth_evt ERC20Transfer

	if !bytes.Equal(log.Address.Bytes(),cosmic_token_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
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
func proc_charity_percentage_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCharityPercentageChanged
	var eth_evt CosmicGameCharityPercentageChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing CharityPercentageChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CharityPercentageChanged",log.Data)
	if err != nil {
		Error.Printf("Event CharityPercentageChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCharityPercentage= eth_evt.NewCharityPercentage.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CharityPercentageChanged {\n")
	Info.Printf("\tNewCharityPercentage: %v\n",evt.NewCharityPercentage)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_charity_percentage_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_charity_percentage_changed_event(&evt)
}
func proc_prize_percentage_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGPrizePercentageChanged
	var eth_evt CosmicGamePrizePercentageChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing PrizePercentageChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"PrizePercentageChanged",log.Data)
	if err != nil {
		Error.Printf("Event PrizePercentageChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewPrizePercentage= eth_evt.NewPrizePercentage.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PrizePercentageChanged {\n")
	Info.Printf("\tNewPrizePercentage: %v\n",evt.NewPrizePercentage)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_prize_percentage_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_prize_percentage_changed_event(&evt)
}
func proc_raffle_percentage_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRafflePercentageChanged
	var eth_evt CosmicGameRafflePercentageChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing RafflePercentageChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"RafflePercentageChanged",log.Data)
	if err != nil {
		Error.Printf("Event RafflePercentageChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewRafflePercentage= eth_evt.NewRafflePercentage.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RafflePercentageChanged {\n")
	Info.Printf("\tNewRafflePercentage: %v\n",evt.NewRafflePercentage)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_raffle_percentage_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_raffle_percentage_changed_event(&evt)
}
func proc_staking_percentage_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGStakingPercentageChanged
	var eth_evt CosmicGameStakingPercentageChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing StakingPercentageChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"StakingPercentageChanged",log.Data)
	if err != nil {
		Error.Printf("Event StakingPercentageChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewStakingPercentage= eth_evt.NewStakingPercentage.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("StakingPercentageChanged {\n")
	Info.Printf("\tNewStakingPercentage: %v\n",evt.NewStakingPercentage)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_staking_percentage_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_staking_percentage_changed_event(&evt)
}
func proc_num_raffle_eth_winners_bidding_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNumRaffleETHWinnersBiddingChanged 
	var eth_evt CosmicGameNumRaffleETHWinnersBiddingChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing NumRaffleETHWinnersBiddingChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"NumRaffleETHWinnersBiddingChanged",log.Data)
	if err != nil {
		Error.Printf("Event NumRaffleETHWinnersBiddingChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewNumRaffleETHWinnersBidding = eth_evt.NewNumRaffleETHWinnersBidding.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NumRaffleETHWinnersBiddingChanged{\n")
	Info.Printf("\tNewNumRaffleETHWinnersBidding: %v\n",evt.NewNumRaffleETHWinnersBidding)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_num_raffle_eth_winners_bidding_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_num_raffle_eth_winners_bidding_changed_event(&evt)
}
func proc_num_raffle_nft_winners_bidding_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNumRaffleNFTWinnersBiddingChanged 
	var eth_evt CosmicGameNumRaffleNFTWinnersBiddingChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing NumRaffleNFTWinnersBiddingChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"NumRaffleNFTWinnersBiddingChanged",log.Data)
	if err != nil {
		Error.Printf("Event NumRaffleNFTWinnersBiddingChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewNumRaffleNFTWinnersBidding  = eth_evt.NewNumRaffleNFTWinnersBidding.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NumRaffleNFTWinnersBiddingChanged{\n")
	Info.Printf("\tNewNumRaffleNFTWinnersBidding: %v\n",evt.NewNumRaffleNFTWinnersBidding)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_num_raffle_nft_winners_bidding_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_num_raffle_nft_winners_bidding_changed_event(&evt)
}
func proc_num_raffle_nft_winners_staking_rwalk_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNumRaffleNFTWinnersStakingRWalkChanged
	var eth_evt CosmicGameNumRaffleNFTWinnersStakingRWalkChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing NumRaffleNFTWinnersStakingRWalkChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"NumRaffleNFTWinnersStakingRWalkChanged",log.Data)
	if err != nil {
		Error.Printf("Event NumRaffleNFTWinnersStakingRWalkChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewNumRaffleNFTWinnersStakingRWalk = eth_evt.NewNumRaffleNFTWinnersStakingRWalk.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NumRaffleNFTWinnersStakingRWalkChanged{\n")
	Info.Printf("\tNewNumRaffleNFTWinnersStakingRWalk: %v\n",evt.NewNumRaffleNFTWinnersStakingRWalk)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_num_raffle_nft_winners_staking_rwalk_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_num_raffle_nft_winners_staking_rwalk_changed_event(&evt)
}
func proc_system_mode_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGSystemModeChanged
	var eth_evt CosmicGameSystemModeChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing SystemModeChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"SystemModeChanged",log.Data)
	if err != nil {
		Error.Printf("Event SystemModeChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewSystemMode = eth_evt.NewSystemMode.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("SystemModeChanged {\n")
	Info.Printf("\tNewSystemMode: %v\n",evt.NewSystemMode)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_system_mode_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_system_mode_changed_event(&evt)
}
func proc_charity_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCharityAddressChanged
	var eth_evt CosmicGameCharityAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing CharityAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CharityAddressChanged",log.Data)
	if err != nil {
		Error.Printf("Event CharityAddressChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCharity = eth_evt.NewCharity.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CharityAddressChanged{\n")
	Info.Printf("\tNewCharity: %v\n",evt.NewCharity)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_charity_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_charity_address_changed_event(&evt)
}
func proc_random_walk_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRandomWalkAddressChanged
	var eth_evt CosmicGameRandomWalkAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing RandomWalkAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"RandomWalkAddressChanged",log.Data)
	if err != nil {
		Error.Printf("Event RandomWalkAddressChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewRandomWalk = eth_evt.NewRandomWalk.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RandomWalkAddressChanged{\n")
	Info.Printf("\tNewRandomWalk: %v\n",evt.NewRandomWalk)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_random_walk_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_random_walk_address_changed_event(&evt)
}
func proc_raffle_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRaffleWalletAddressChanged
	var eth_evt CosmicGameRaffleWalletAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing RaffleWalletAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"RaffleWalletAddressChanged",log.Data)
	if err != nil {
		Error.Printf("Event RaffleWalletAddressChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewRaffleWallet = eth_evt.NewRaffleWallet.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RaffleWalletAddressChanged{\n")
	Info.Printf("\tNewRaffleWallet: %v\n",evt.NewRaffleWallet)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_raffle_wallet_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_raffle_wallet_address_changed_event(&evt)
}
func proc_staking_wallet_cst_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGStakingWalletCSTAddressChanged
	var eth_evt CosmicGameStakingWalletCSTAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing StakingWalletCSTAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"StakingWalletCSTAddressChanged",log.Data)
	if err != nil {
		Error.Printf("Event StakingWalletCSTAddressChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewStakingWalletCST = eth_evt.NewStakingWalletCST.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("StakingWalletCSTAddressChanged{\n")
	Info.Printf("\tNewStakingWalletCST: %v\n",evt.NewStakingWalletCST)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_staking_wallet_cst_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_staking_wallet_cst_address_changed_event(&evt)
}
func proc_staking_wallet_rwalk_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGStakingWalletRWalkAddressChanged
	var eth_evt CosmicGameStakingWalletRWalkAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing StakingWalletRWalkAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"StakingWalletRWalkAddressChanged",log.Data)
	if err != nil {
		Error.Printf("Event StakingWalletRWalkAddressChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewStakingWalletRWalk = eth_evt.NewStakingWalletRWalk.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("StakingWalletRWalkAddressChanged{\n")
	Info.Printf("\tNewStakingWalletRWalk: %v\n",evt.NewStakingWalletRWalk)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_staking_wallet_rwalk_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_staking_wallet_rwalk_address_changed_event(&evt)
}
func proc_marketing_wallet_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGMarketingWalletAddressChanged
	var eth_evt CosmicGameMarketingWalletAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing MarketingWalletAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"MarketingWalletAddressChanged",log.Data)
	if err != nil {
		Error.Printf("Event MarketingWalletAddressChanged decode error: %v",err)
		os.Exit(1)
	}
	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewMarketingWallet = eth_evt.NewMarketingWallet.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("MarketingWalletAddressChanged{\n")
	Info.Printf("\tNewMarketingWallet: %v\n",evt.NewMarketingWallet)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_marketing_wallet_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_marketing_wallet_address_changed_event(&evt)
}
func proc_cosmic_token_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCosmicTokenAddressChanged
	var eth_evt CosmicGameCosmicTokenAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing CosmicTokenAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CosmicTokenAddressChanged",log.Data)
	if err != nil {
		Error.Printf("Event CosmicTokenAddressChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCosmicToken= eth_evt.NewCosmicToken.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CosmicTokenAddressChanged{\n")
	Info.Printf("\tNewCosmicToken: %v\n",evt.NewCosmicToken)
	Info.Printf("}\n")

	storagew.Delete_cosmic_token_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_token_address_changed_event(&evt)
}
func proc_cosmic_signature_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCosmicSignatureAddressChanged
	var eth_evt CosmicGameCosmicSignatureAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing CosmicSignatureAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CosmicSignatureAddressChanged",log.Data)
	if err != nil {
		Error.Printf("Event CosmicSignatureAddressChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCosmicSignature= eth_evt.NewCosmicSignature.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CosmicSignatureAddressChanged{\n")
	Info.Printf("\tNewCosmicSignatureWallet: %v\n",evt.NewCosmicSignature)
	Info.Printf("}\n")

	storagew.Delete_cosmic_signature_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_signature_address_changed_event(&evt)
}
func proc_business_logic_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGBusinessLogicAddressChanged
	var eth_evt CosmicGameBusinessLogicAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing BusinessLogicAddressChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"BusinessLogicAddressChanged",log.Data)
	if err != nil {
		Error.Printf("Event BusinessLogicAddressChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewContractAddress = eth_evt.NewContractAddress.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("BusinessLogicAddressChanged{\n")
	Info.Printf("\tNewContractAddress: %v\n",evt.NewContractAddress)
	Info.Printf("}\n")

	storagew.Delete_business_logic_address_changed_event(evt.EvtId)
    storagew.Insert_business_logic_address_changed_event(&evt)
}
func proc_time_increase_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGTimeIncreaseChanged
	var eth_evt CosmicGameTimeIncreaseChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
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
	evt.NewTimeIncrease = eth_evt.NewTimeIncrease.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TimeIncreaseChanged{\n")
	Info.Printf("\tNewTimeIncrease: %v\n",evt.NewTimeIncrease)
	Info.Printf("}\n")

	storagew.Delete_time_increase_changed_event(evt.EvtId)
    storagew.Insert_time_increase_changed_event(&evt)
}
func proc_timeout_claimprize_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGTimeoutClaimPrizeChanged
	var eth_evt CosmicGameTimeoutClaimPrizeChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing TimeoutClaimPrizeChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"TimeoutClaimPrizeChanged",log.Data)
	if err != nil {
		Error.Printf("Event TimeoutClaimPrizeChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewTimeout = eth_evt.NewTimeout.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("TimeIncreaseChanged{\n")
	Info.Printf("\tNewTimeout: %v\n",evt.NewTimeout)
	Info.Printf("}\n")

	storagew.Delete_timeout_claimprize_changed_event(evt.EvtId)
    storagew.Insert_timeout_claimprize_changed_event(&evt)
}
func proc_price_increase_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGPriceIncreaseChanged
	var eth_evt CosmicGamePriceIncreaseChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing PriceIncreaseChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"PriceIncreaseChanged",log.Data)
	if err != nil {
		Error.Printf("Event PriceIncreaseChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewPriceIncrease = eth_evt.NewPriceIncrease.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PriceIncreaseChanged{\n")
	Info.Printf("\tNewPriceIncreasse: %v\n",evt.NewPriceIncrease)
	Info.Printf("}\n")

	storagew.Delete_price_increase_changed_event(evt.EvtId)
    storagew.Insert_price_increase_changed_event(&evt)
}
func proc_nanoseconds_extra_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNanoSecondsExtraChanged
	var eth_evt CosmicGameNanoSecondsExtraChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing NanoSecondsExtraChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"NanoSecondsExtraChanged",log.Data)
	if err != nil {
		Error.Printf("Event NanoSecondsExtraChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewNanoSecondsExtra = eth_evt.NewNanoSecondsExtra.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NanoSecondsExtraChanged{\n")
	Info.Printf("\tNewNanoSecondsExtraChanged: %v\n",evt.NewNanoSecondsExtra)
	Info.Printf("}\n")

	storagew.Delete_nanoseconds_extra_changed_event(evt.EvtId)
    storagew.Insert_nanoseconds_extra_changed_event(&evt)
}
func proc_initial_seconds_until_prize_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGInitialSecondsUntilPrizeChanged
	var eth_evt CosmicGameInitialSecondsUntilPrizeChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing InitialSecondsUntilPrizeChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"InitialSecondsUntilPrizeChanged",log.Data)
	if err != nil {
		Error.Printf("Event InitialSecondsUntilPrizeChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewInitialSecondsUntilPrize = eth_evt.NewInitialSecondsUntilPrize.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("InitialSecondsUntilPrizeChanged{\n")
	Info.Printf("\tNewInitialSecondsUntilPrize: %v\n",evt.NewInitialSecondsUntilPrize)
	Info.Printf("}\n")

	storagew.Delete_initial_seconds_until_prize_changed_event(evt.EvtId)
    storagew.Insert_initial_seconds_until_prize_changed_event(&evt)
}
func proc_initial_bid_amount_fraction_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGInitialBidAmountFractionChanged
	var eth_evt CosmicGameInitialBidAmountFractionChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing InitialBidAmountFractionChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"InitialBidAmountFractionChanged",log.Data)
	if err != nil {
		Error.Printf("Event InitialBidAmountFractionChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewInitialBidAmountFraction = eth_evt.NewInitialBidAmountFraction.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("InitialBidAmountFractionChanged {\n")
	Info.Printf("\tNewInitialBidAmountFraction: %v\n",evt.NewInitialBidAmountFraction)
	Info.Printf("}\n")

	storagew.Delete_initial_bid_amount_fraction_changed_event(evt.EvtId)
    storagew.Insert_initial_bid_amount_fraction_changed_event(&evt)
}
func proc_activation_time_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGActivationTimeChanged
	var eth_evt CosmicGameActivationTimeChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing ActivationTimeChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"ActivationTimeChanged",log.Data)
	if err != nil {
		Error.Printf("Event ActivationTimeChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewActivationTime = eth_evt.NewActivationTime.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("ActivationTimeChanged {\n")
	Info.Printf("\tNewActivationTime: %v\n",evt.NewActivationTime)
	Info.Printf("}\n")

	storagew.Delete_activation_time_changed_event(evt.EvtId)
    storagew.Insert_activation_time_changed_event(&evt)
}
func proc_round_start_cst_auction_length_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRoundStartCSTAuctionLengthChanged
	var eth_evt CosmicGameRoundStartCSTAuctionLengthChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing RoundStartCSTAuctionLengthChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"RoundStartCSTAuctionLengthChanged",log.Data)
	if err != nil {
		Error.Printf("Event RoundStartCSTAuctionLengthChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewAuctionLength = eth_evt.NewAuctionLength.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RoundStartCSTAuctionLengthChanged {\n")
	Info.Printf("\tNewAuctionLength: %v\n",evt.NewAuctionLength)
	Info.Printf("}\n")

	storagew.Delete_round_start_cst_auction_length_changed_event(evt.EvtId)
    storagew.Insert_round_start_cst_auction_length_changed_event(&evt)
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
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_donation_with_info_event) {
		proc_donation_with_info_event(log,evtlog)
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
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_raffle_withdrawal) {
		proc_raffle_withdrawal_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_raffle_nft_winner) {
		proc_raffle_nft_winner_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_endurance_winner) {
	proc_endurance_winner_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_stellar_winner) {
		proc_stellar_winner_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_donated_nft_claimed) {
		proc_donated_nft_claimed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer) {
		proc_transfer_event_common(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_stake_action) {
		select_and_proc_stake_action_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_unstake_action) {
		select_and_proc_unstake_action_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_eth_deposit) {
		proc_eth_deposit_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_claim_reward) {
		proc_claim_reward_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_marketing_reward_sent) {
		proc_marketing_reward_sent_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_charity_percentage_changed) {
		proc_charity_percentage_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_prize_percentage_changed) {
		proc_prize_percentage_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_raffle_percentage_changed) {
		proc_raffle_percentage_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_staking_percentage_changed) {
		proc_staking_percentage_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_num_raffle_eth_winners_bidding_changed) {
		proc_num_raffle_eth_winners_bidding_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_num_raffle_nft_winners_bidding_changed) {
		proc_num_raffle_nft_winners_bidding_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_num_raffle_nft_winners_staking_rwalk_changed) {
		proc_num_raffle_nft_winners_staking_rwalk_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_system_mode_changed) {
		proc_system_mode_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_charity_address_changed) {
		proc_charity_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_rwalk_address_changed) {
		proc_random_walk_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_raffle_address_changed) {
		proc_raffle_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_staking_wallet_cst_address_changed) {
		proc_staking_wallet_cst_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_staking_wallet_rwalk_address_changed) {
		proc_staking_wallet_rwalk_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_marketing_address_changed) {
		proc_marketing_wallet_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_costok_address_changed) {
		proc_cosmic_token_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_cossig_address_changed) {
		proc_cosmic_signature_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_blogic_address_changed) {
		proc_business_logic_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_time_increase_changed) {
		proc_time_increase_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_timeout_claimprize_changed) {
		proc_timeout_claimprize_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_price_increase_changed) {
		proc_price_increase_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_nanoseconds_extra_changed) {
		proc_nanoseconds_extra_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_initial_seconds_until_prize_changed) {
		proc_initial_seconds_until_prize_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_initial_bid_amount_fraction_changed) {
		proc_initial_bid_amount_fraction_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_activation_time_changed) {
		proc_activation_time_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_round_start_auction_length_changed) {
		proc_round_start_cst_auction_length_changed_event(log,evtlog)
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
