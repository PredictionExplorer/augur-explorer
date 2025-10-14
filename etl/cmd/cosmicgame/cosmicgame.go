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
			Signature: hex.EncodeToString(evt_erc20_donated[:4]),
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
			Signature: hex.EncodeToString(evt_raffle_eth_winner[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_endurance_winner[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_lastcst_bidder_winner[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_donated_token_claimed[:4]),
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
			Signature: hex.EncodeToString(evt_chrono_percentage_changed[:4]),
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
			Signature: hex.EncodeToString(evt_prizes_wallet_address_changed[:4]),
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
			Signature: hex.EncodeToString(evt_treasurer_changed[:4]),
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
			Signature: hex.EncodeToString(evt_time_increase_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_timeout_claimprize_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_timeout_to_withdraw_prize[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_price_increase_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_prize_microsecond_increase_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_initial_seconds_until_prize_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_activation_time_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_cst_dutch_auction_duration_divisor_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_eth_dutch_auction_duration_divisor_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_eth_dutch_auction_ending_bidprice_divisor[:4]),
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
			Signature: hex.EncodeToString(evt_eth_prize_deposit[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_staking_eth_deposit[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_eth_prize_withdrawal[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_transfer[:4]),
			ContractAid: cosmic_sig_aid,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_cst_nft_staked[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_rwalk_nft_staked[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_nft_unstaked_rwalk[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_nft_unstaked_cst[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_marketing_reward_sent[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_cst_reward_for_bidding_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_erc20_reward_mult[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_max_msg_length_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_token_script_url[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_base_uri[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_proxy_upgraded[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_admin_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_marketing_reward_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_ownership_transferred[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_initialized[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_chrono_warrior[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_cst_min_limit[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_fund_transf_err[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_erc20_transf_err[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_funds2charity[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_delay_duration_round[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_round_started[:4]),
			ContractAid: 0,
		},
	)
	return inspected_events
}
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
	Info.Printf("\tDonationEvtId: %v\n",evt.DonationEvtId)
	Info.Printf("\tTimeout to withdraw: %v\n",evt.Timeout)
	Info.Printf("}\n")

	storagew.Delete_prize_claim_event(evt.EvtId)
	storagew.Insert_prize_claim_event(&evt)
}
func find_cosmic_token_transfer(bid_evtlog_id,tx_id int64) string {
	// fetches the ERC20::Transfer event which has the id=evtlog-1 because it is
	//		inserted right before Bid event
	//		this function panics in case of failure because that would be an invalid database state
	elog_ids := storagew.S.Get_specific_event_logs_by_tx_backwards_from_id(tx_id,cosmic_tok_aid,bid_evtlog_id,hex.EncodeToString(evt_transfer[:4]))
//	elog_ids,err := storagew.S.Get_events_by_sig_and_tx_id(tx_id,bid_evtlog_id,hex.EncodeToString(evt_transfer[:4]))	// ERC20 tansfer is always 2 less than the bid (-1 is for marketing reward but -2 is the bid reward)
	if len(elog_ids) == 0 {
		err_str := fmt.Sprintf("Couldn't find any event logs of erc transfer in BidPlaced event (tx_id=%v)\n",tx_id)
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	ee := elog_ids[0]
	var log types.Log
	err := rlp.DecodeBytes(ee,&log)
	if err!= nil {
		err_str := fmt.Sprintf("RLP Decode error at find_cosmic_signature_token_transfer(): %v\n",err)
		Info.Printf(err_str)
		Error.Printf(err_str)
		os.Exit(1)
	}
	var eth_evt ERC20Transfer
	err = erc20_abi.UnpackIntoInterface(&eth_evt,"Transfer",log.Data)
	if err != nil {
		err_str := fmt.Sprintf("Event Transfer decode error at find_cosmic_signature_token_transfer(): %v\n",err)
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
	var eth_evt CosmicSignatureGameBidPlaced

	Info.Printf("Processing Bid event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"BidPlaced",log.Data)
	if err != nil {
		Error.Printf("Event BidPlaced decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.LastBidderAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.BidPrice = eth_evt.PaidEthPrice.String()
	evt.BidType = 0; // ETH
	evt.RandomWalkTokenId = log.Topics[3].Big().Int64()
	evt.ERC20_Value = find_cosmic_token_transfer(evt.EvtId,evt.TxId)
	evt.CstPrice = eth_evt.PaidCstPrice.String()
	if evt.RandomWalkTokenId > -1 {
		evt.BidType = 1;	// RandomWalk	
	} else {
		if evt.CstPrice != "-1" { evt.BidType = 2; } // Cosmic Signature Token (ERC20) bid
	}
	evt.PrizeTime = eth_evt.MainPrizeTime.Int64()
	evt.Message = eth_evt.Message

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("BidPlaced {\n")
	Info.Printf("\tLastBidder: %v\n",evt.LastBidderAddr)
	Info.Printf("\tRoundNum: %v\n",evt.RoundNum)
	Info.Printf("\tBidPrice: %v\n",evt.BidPrice)
	Info.Printf("\tCstPrice: %v\n",evt.CstPrice);
	Info.Printf("\tRandomWalkTokenId: %v\n",evt.RandomWalkTokenId)
	Info.Printf("\tPrizeTime: %v\n",evt.PrizeTime)
	Info.Printf("\tMessage: %v\n",evt.Message)
	Info.Printf("}\n")

	storagew.Delete_bid(evt.EvtId)
	storagew.Insert_bid_event(&evt)
}
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
	Info.Printf("\tLinking to bid id: \n",evt.BidId)
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
func proc_charity_updated_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCharityUpdatedEvent
	var eth_evt CharityWalletCharityAddressChanged

	Info.Printf("Processing CharityUpdatedEvent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),charity_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := charity_wallet_abi.UnpackIntoInterface(&eth_evt,"CharityAddressChanged",log.Data)
	if err != nil {
		Error.Printf("Event CharityAddressChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewCharityAddr = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CharityAddressChanged {\n")
	Info.Printf("\tNewValue: %v\n",evt.NewCharityAddr)
	Info.Printf("}\n")

	storagew.Delete_charity_updated(evt.EvtId)
	storagew.Insert_charity_updated_event(&evt)
}
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
	evt.Seed = hex.EncodeToString(eth_evt.NftSeed.Bytes())

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
	evt.WinnerAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("PrizeWithdrawn {\n")
	Info.Printf("\tWinnerAddr: %v\n",evt.WinnerAddr)
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
func proc_cst_nft_staked_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNftStakedCst
	var eth_evt IStakingWalletCosmicSignatureNftNftUnstaked 

	Info.Printf("Processing CST NftStaked event id=%v, (block %v) txhash %v\n",elog.EvtId,elog.BlockNum,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),staking_wallet_cst_addr.Bytes()) {
		if !bytes.Equal(log.Address.Bytes(),staking_wallet_rwalk_addr.Bytes()) {
			Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
			return
		}
	}
	err := staking_wallet_cst_abi.UnpackIntoInterface(&eth_evt,"NftStaked",log.Data)
	if err != nil {
		Error.Printf("CST Event CST NFT Staked decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.NftId = log.Topics[2].Big().Int64()
	evt.StakerAddress = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.NumStakedNfts = eth_evt.NumStakedNfts.Int64()
	evt.RewardPerStaker = eth_evt.RewardAmountPerStakedNft.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CST NftStaked event{\n")
	Info.Printf("\tActionId: %v\n",evt.ActionId)
	Info.Printf("\tTokenId: %v\n",evt.NftId)
	Info.Printf("\tTotalNFTs: %v\n",evt.NumStakedNfts)
	Info.Printf("\tRewardPerStaker: %v\n",evt.RewardPerStaker)
	Info.Printf("\tStaker: %v\n",evt.StakerAddress)
	Info.Printf("}\n")

	storagew.Delete_nft_staked_cst_event(evt.EvtId)
	storagew.Insert_nft_staked_cst_event(&evt)
}
func proc_rwalk_nft_staked_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNftStakedRWalk
	var eth_evt IStakingWalletRandomWalkNftNftStaked

	Info.Printf("Processing RWalk NftStaked event id=%v, (block %v) txhash %v\n",elog.EvtId,elog.BlockNum,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),staking_wallet_cst_addr.Bytes()) {
		if !bytes.Equal(log.Address.Bytes(),staking_wallet_rwalk_addr.Bytes()) {
			Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
			return
		}
	}
	err := staking_wallet_rwalk_abi.UnpackIntoInterface(&eth_evt,"NftStaked",log.Data)
	if err != nil {
		Error.Printf("RWalk NftStaked decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.NftId = log.Topics[2].Big().Int64()
	evt.StakerAddress = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.NumStakedNfts = eth_evt.NumStakedNfts.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("RWalk NftStaked event{\n")
	Info.Printf("\tActionId: %v\n",evt.ActionId)
	Info.Printf("\tTokenId: %v\n",evt.NftId)
	Info.Printf("\tTotalNFTs: %v\n",evt.NumStakedNfts)
	Info.Printf("\tStaker: %v\n",evt.StakerAddress)
	Info.Printf("}\n")

	storagew.Delete_nft_staked_rwalk_event(evt.EvtId)
	storagew.Insert_nft_staked_rwalk_event(&evt)
}
func proc_staking_eth_deposit_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGEthDeposit
	var eth_evt IStakingWalletCosmicSignatureNftEthDepositReceived

	Info.Printf("Processing EthDepositReceived event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(),staking_wallet_cst_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := staking_wallet_cst_abi.UnpackIntoInterface(&eth_evt,"EthDepositReceived",log.Data)
	if err != nil {
		Error.Printf("Event EthDepositReceived Event decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DepositTime = elog.TimeStamp
	evt.DepositId  = eth_evt.ActionCounter.Int64()
	evt.NumStakedNfts = eth_evt.NumStakedNfts.Int64()
	evt.Amount = eth_evt.DepositAmount.String()
	evt.AccumModulo = "0";	// pending for resolution regarding StakingWalletCST refactoring
	evt.RoundNum = log.Topics[1].Big().Int64()
	divres:=big.NewInt(0)
	rem:=big.NewInt(0)
	divres.QuoRem(eth_evt.DepositAmount,eth_evt.NumStakedNfts,rem);
	evt.AmountPerStaker = divres.String()
	evt.Modulo = rem.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("EthDepositReceived {\n")
	Info.Printf("\tDepositTime: %v\n",evt.DepositTime)
	Info.Printf("\tDepositId: %v\n",evt.DepositId)
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
func proc_nft_unstaked_rwalk_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNftUnstakedRWalk
	var eth_evt IStakingWalletRandomWalkNftNftUnstaked

	Info.Printf("Processing NftUnstaked event id=%v, (block %v) txhash %v\n",elog.EvtId,elog.BlockNum,elog.TxHash)
	if !bytes.Equal(log.Address.Bytes(),staking_wallet_rwalk_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := staking_wallet_rwalk_abi.UnpackIntoInterface(&eth_evt,"NftUnstaked",log.Data)
	if err != nil {
		Error.Printf("RWalk Event NftUnstaked decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.ActionId = log.Topics[1].Big().Int64()
	evt.NftId = log.Topics[2].Big().Int64()
	evt.StakerAddress = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.NumStakedNfts = eth_evt.NumStakedNfts.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NftUnstaked RWalk{\n")
	Info.Printf("\tStakeActionId: %v\n",evt.ActionId)
	Info.Printf("\tNftId: %v\n",evt.NftId)
	Info.Printf("\tNumStakedNfts: %v\n",evt.NumStakedNfts)
	Info.Printf("\tStakerAddress: %v\n",evt.StakerAddress)
	Info.Printf("}\n")

	storagew.Delete_nft_unstaked_rwalk_event(evt.EvtId)
	storagew.Insert_nft_unstaked_rwalk_event(&evt)
}
func proc_nft_unstaked_cst_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGNftUnstakedCst
	var eth_evt IStakingWalletCosmicSignatureNftNftUnstaked

	Info.Printf("Processing NftUnstaked CST event id=%v, (block %v) txhash %v\n",elog.EvtId,elog.BlockNum,elog.TxHash)
	if !bytes.Equal(log.Address.Bytes(),staking_wallet_cst_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	err := staking_wallet_cst_abi.UnpackIntoInterface(&eth_evt,"NftUnstaked",log.Data)
	if err != nil {
		Error.Printf("RWalk Event NftUnstaked decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
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

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("NftUnstaked Cst {\n")
	Info.Printf("\tStakeActionId: %v\n",evt.ActionId)
	Info.Printf("\tNftId: %v\n",evt.NftId)
	Info.Printf("\tNumStakedNfts: %v\n",evt.NumStakedNfts)
	Info.Printf("\tStakerAddress: %v\n",evt.StakerAddress)
	Info.Printf("\tRewardAmount (total): %v\n",evt.RewardAmount)
	Info.Printf("\tRewardPerToken: %v\n",evt.RewardPerToken)
	Info.Printf("\tActionCounter: %v\n",evt.ActionCounter)
	Info.Printf("}\n")

	storagew.Delete_nft_unstaked_cst_event(evt.EvtId)
	storagew.Insert_nft_unstaked_cst_event(&evt)
}
func proc_marketing_reward_sent_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGMarketingRewardSent
	var eth_evt MarketingWalletRewardPaid

	Info.Printf("Processing MarketingWallet RewardSent event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)

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
	Info.Printf("Marketing RwardPaid{\n")
	Info.Printf("\tAmount: %v\n",evt.Amount)
	Info.Printf("\tMarketer: %v\n",evt.Marketer)
	Info.Printf("}\n")

	storagew.Delete_marketing_reward_sent_event(evt.EvtId)
	storagew.Insert_marketing_reward_sent_event(&evt)
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
func proc_charity_address_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGCharityAddressChanged
	var eth_evt CosmicSignatureGameCharityAddressChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
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
	evt.NewCharity = common.BytesToAddress(log.Topics[1][12:]).String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CharityAddressChanged{\n")
	Info.Printf("\tNewCharity: %v\n",evt.NewCharity)
	Info.Printf("}\n")

	storagew.Delete_cosmic_game_charity_address_changed_event(evt.EvtId)
    storagew.Insert_cosmic_game_charity_address_changed_event(&evt)
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
	eth_evt.NewValue.String()

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

	var evt CGCstDutchAuctionDurationDivisorChanged
	var eth_evt CosmicSignatureGameCstDutchAuctionDurationDivisorChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing RoundStartCstAuctionLengthChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CstDutchAuctionDurationDivisorChanged",log.Data)
	if err != nil {
		Error.Printf("Event CstDutchAuctionDurationDivisorChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewValue  = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CstDutchAuctionDurationDivisorChanged {\n")
	Info.Printf("\tNewAuctionLength: %v\n",evt.NewValue)
	Info.Printf("}\n")

	storagew.Delete_round_start_cst_auction_length_changed_event(evt.EvtId)
    storagew.Insert_round_start_cst_auction_length_changed_event(&evt)
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

	var evt CGCstRewardForBiddingChanged 
	var eth_evt CosmicSignatureGameMarketingWalletCstContributionAmountChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing CstRewardForBiddingChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CstRewardAmountForBiddingChanged",log.Data)
	if err != nil {
		Error.Printf("Event CstRewardAmountForBiddingChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewReward = eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("CstRewardAmountForBiddingChanged{\n")
	Info.Printf("\tNewReward: %v\n",evt.NewReward)
	Info.Printf("}\n")

	storagew.Delete_erc20_token_reward_changed_event(evt.EvtId)
    storagew.Insert_erc20_token_reward_changed_event(&evt)
}
func proc_erc20_reward_multiplier_changed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGERC20RewardMultiplierChanged
	var eth_evt BiddingCstPrizeAmountChanged

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing CstPrizeAmountChanged event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"CstPrizeAmountChanged",log.Data)
	if err != nil {
		Error.Printf("Event CstPrizeAmountMultiplierChanged decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.NewMultiplier= eth_evt.NewValue.String()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("cstPrizeMultiplierChanged{\n")
	Info.Printf("\tNewMultiplier: %v\n",evt.NewMultiplier)
	Info.Printf("}\n")

	storagew.Delete_erc20_reward_multiplier_changed_event(evt.EvtId)
    storagew.Insert_erc20_reward_multiplier_changed_event(&evt)
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
func proc_initialized_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGInitialized
	var eth_evt CosmicSignatureGameInitialized 

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
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
/* DISCONTINUED , pending for revision that in Solidity we report a failure if ERC20 mint fails 
func proc_erc20_transfer_failed_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGErc20TransferFailed
	var eth_evt CosmicSignatureEventsERC20TransferFailed

	if !bytes.Equal(log.Address.Bytes(),marketing_wallet_addr.Bytes()) {
		//Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n",log.Address.String())
		return
	}
	Info.Printf("Processing Initialized event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"ERC20TransferFailed",log.Data)
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

	storagew.Delete_erc20_transfer_failed_event(evt.EvtId)
    storagew.Insert_erc20_transfer_failed_event(&evt)
}*/
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
func proc_round_started_event(log *types.Log,elog *EthereumEventLog) {

	var evt CGRoundStarted
	var eth_evt CosmicSignatureGameFirstBidPlacedInRound

	if !bytes.Equal(log.Address.Bytes(),cosmic_game_addr.Bytes()) {
		return
	}
	Info.Printf("Processing FirstBidPlacedInRound  event id=%v, txhash %v\n",elog.EvtId,elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt,"FirstBidPlacedInRound",log.Data)
	if err != nil {
		Error.Printf("Event FirstBidPlacedInRound decode error: %v",err)
		os.Exit(1)
	}

	evt.EvtId=elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.StartTimestamp= eth_evt.BlockTimeStamp.Int64()

	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("FirstBidPlacedInRound{\n")
	Info.Printf("\tRoundNum: %v\n",evt.RoundNum)
	Info.Printf("\tStart TS: %v\n",evt.StartTimestamp)
	Info.Printf("}\n")

    storagew.Delete_round_started_event(evt.EvtId)
	storagew.Insert_round_started_event(&evt)
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
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc20_donated) {
		proc_erc20_donated_event(log,evtlog)
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
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_eth_prize_deposit) {
		proc_prizes_eth_deposit_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_eth_prize_withdrawal) {
		proc_eth_prize_withdrawal_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_raffle_eth_winner) {
		proc_raffle_eth_winner_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_raffle_nft_winner) {
		proc_raffle_nft_winner_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_endurance_winner) {
	proc_endurance_winner_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_lastcst_bidder_winner) {
		proc_lastcst_bidder_winner_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_chrono_warrior) {
		proc_chrono_warrior_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_donated_token_claimed) {
		proc_donated_token_claimed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_donated_nft_claimed) {
		proc_donated_nft_claimed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_transfer) {
		proc_transfer_event_common(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_cst_nft_staked) {
		proc_cst_nft_staked_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_rwalk_nft_staked) {
		proc_rwalk_nft_staked_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_nft_unstaked_rwalk) {
		proc_nft_unstaked_rwalk_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_nft_unstaked_cst) {
		proc_nft_unstaked_cst_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_staking_eth_deposit) {
		proc_staking_eth_deposit_event(log,evtlog)
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
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_chrono_percentage_changed) {
		proc_chrono_percentage_changed_event(log,evtlog)
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
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_charity_address_changed) {
		proc_charity_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_rwalk_address_changed) {
		proc_random_walk_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_prizes_wallet_address_changed) {
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
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_treasurer_changed) {
		proc_treasurer_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_costok_address_changed) {
		proc_cosmic_token_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_cossig_address_changed) {
		proc_cosmic_signature_address_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_proxy_upgraded) {
		proc_proxy_upgraded_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_admin_changed) {
		proc_admin_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_time_increase_changed) {
		proc_time_increase_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_timeout_claimprize_changed) {
		proc_timeout_claimprize_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_timeout_to_withdraw_prize) {
		proc_timeout_duration_to_withdraw_prize_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_price_increase_changed) {
		proc_price_increase_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_prize_microsecond_increase_changed) {
		proc_mainprize_microsecond_increase_changed(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_initial_seconds_until_prize_changed) {
		proc_initial_seconds_until_prize_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_activation_time_changed) {
		proc_activation_time_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_cst_dutch_auction_duration_divisor_changed) {
		proc_cst_dutch_auction_duration_divisor_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_eth_dutch_auction_duration_divisor_changed) {
		proc_eth_dutch_auction_duration_divisor_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_eth_dutch_auction_ending_bidprice_divisor) {
		proc_eth_dutch_auction_ending_bid_price_divisor_changed__event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_cst_reward_for_bidding_changed) {
		proc_erc20_token_reward_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc20_reward_mult) {
		proc_erc20_reward_multiplier_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_max_msg_length_changed) {
		proc_max_msg_length_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_token_script_url) {
		proc_token_generation_script_url_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_base_uri) {
		proc_base_uri_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_marketing_reward_changed) {
		proc_marketing_reward_changed(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_ownership_transferred) {
		proc_ownership_transferred_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_initialized) {
		proc_initialized_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_cst_min_limit) {
		proc_starting_bid_price_cst_min_limit_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_fund_transf_err) {
		proc_fund_transfer_failed_event(log,evtlog)
	}
	/* DISCONTINUED
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_erc20_transf_err) {
		proc_erc20_transfer_failed_event(log,evtlog)
	} */
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_funds2charity) {
		proc_funds_transferred_to_charity_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_delay_duration_round) {
		proc_delay_duration_before_next_round_changed_event(log,evtlog)
	}
	if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_round_started) {
		proc_round_started_event(log,evtlog)
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
