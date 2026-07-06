// Shared decode/lookup helpers: RLP event-log decoding and locating related
// events (token transfers, mints, prize claims) within a transaction.
package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func find_cosmic_token_transfer(bid_evtlog_id, tx_id int64, bidder_addr string) (string, error) {
	// Returns the CST bid reward (the amount minted to the bidder) for this bid transaction.
	//
	// We locate the reward by MATCHING the correct Transfer event rather than by positional
	// offset from the bid event. Per bid the only CST mint is the bid reward credited to the
	// bidder, i.e. an ERC20 Transfer with from == zero address (mint) and to == bidder.
	// (A CST bid also burns the paid price, emitting Transfer(bidder -> 0x0); that is ignored
	// by this match. The marketing-wallet CST contribution is minted at round end, not per bid.)
	//
	// In V2 the reward is dynamic (BiddingV2.getBidCstRewardAmountAdvanced) and is 0 for the
	// earliest bid(s) of a round. When it is 0 the contract performs no mint, so there is no
	// matching Transfer and we correctly return "0" instead of failing.
	elog_rlps, err := storagew.S.Get_specific_event_logs_by_tx_backwards_from_id(tx_id, cosmic_tok_aid, bid_evtlog_id, hex.EncodeToString(evt_transfer[:4]))
	if err != nil {
		return "", fmt.Errorf("find_cosmic_token_transfer(): %w", err)
	}
	bidder := common.HexToAddress(bidder_addr)
	for _, raw := range elog_rlps {
		var log types.Log
		err := rlp.DecodeBytes(raw, &log)
		if err != nil {
			err_str := fmt.Sprintf("RLP Decode error at find_cosmic_token_transfer(): %v\n", err)
			Info.Print(err_str)
			Error.Print(err_str)
			os.Exit(1)
		}
		// ERC20 Transfer indexes `from` and `to`, so they live in the topics, not in Data.
		if len(log.Topics) < 3 {
			continue
		}
		from := common.BytesToAddress(log.Topics[1][12:])
		to := common.BytesToAddress(log.Topics[2][12:])
		if from != (common.Address{}) || to != bidder {
			continue // not the bid reward mint to the bidder
		}
		var eth_evt ERC20Transfer
		err = erc20_abi.UnpackIntoInterface(&eth_evt, "Transfer", log.Data)
		if err != nil {
			err_str := fmt.Sprintf("Event Transfer decode error at find_cosmic_token_transfer(): %v\n", err)
			Error.Print(err_str)
			Info.Print(err_str)
			Info.Printf("%+v", log)
			Error.Printf("%+v", log)
			os.Exit(1)
		}
		return eth_evt.Value.String(), nil
	}
	// No mint to the bidder in this transaction => the dynamic CST bid reward was 0.
	return "0", nil
}
func find_cosmic_token_721_transfer(bid_evtlog_id int64) (int64, error) {
	// fetches the ERC721::Transfer event which has the id=evtlog-1 because it is
	//		inserted right before RaffleNFTClaim event
	//		this function aborts in case of a decode failure because that would be an invalid database state
	ee, err := storagew.S.Get_event_log(bid_evtlog_id - 1) // ERC20 tansfer is always 1 less than Bid id
	if err != nil {
		return 0, fmt.Errorf("find_cosmic_token_721_transfer(): %w", err)
	}
	var log types.Log
	err = rlp.DecodeBytes(ee.RlpLog, &log)
	if err != nil {
		err_str := fmt.Sprintf("RLP Decode error at find_cosmic_token_721_transfer(): %v", err)
		Info.Print(err_str)
		Error.Print(err_str)
		os.Exit(1)
	}
	var eth_evt IERC721Transfer
	err = erc721_abi.UnpackIntoInterface(&eth_evt, "Transfer", log.Data)
	if err != nil {
		err_str := fmt.Sprintf("Event Transfer decode error at find_cosmic_token_721_transfer(): %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		Info.Printf("%+v", log)
		Error.Printf("%+v", log)
		os.Exit(1)
	}
	if eth_evt.TokenId != nil {
		Info.Printf("token id=%v\n", eth_evt.TokenId.Int64())
	}
	if len(log.Topics) < 3 {
		err_str := fmt.Sprintf("Event ERC721 Transfer doesn't have 4 topics")
		Error.Print(err_str)
		Info.Print(err_str)
		Info.Printf("%+v", log)
		Error.Printf("%+v", log)
		os.Exit(1)
	}
	return log.Topics[1].Big().Int64(), nil
}
func find_cosmic_token_721_mint_event(contract_aid, tx_id, claim_prize_evtlog_id int64) (int64, error) {

	mint_evt_list, err := storagew.S.Get_specific_event_logs_by_tx_backwards_from_id(tx_id, contract_aid, claim_prize_evtlog_id, hex.EncodeToString(evt_mint_event[0:4]))
	if err != nil {
		return 0, fmt.Errorf("find_cosmic_token_721_mint_event(): %w", err)
	}
	if len(mint_evt_list) == 0 {
		err_str := fmt.Sprintf("find_cosmic_token_721_mint_event() couldn't find corresponding MintEvent()")
		Info.Print(err_str)
		Error.Print(err_str)
		os.Exit(1)
	}
	mint_location := len(mint_evt_list) - 1
	var log types.Log
	err = rlp.DecodeBytes(mint_evt_list[mint_location], &log)
	if err != nil {
		err_str := fmt.Sprintf("RLP Decode error at find_cosmic_token_721_mint_event(): %v", err)
		Info.Print(err_str)
		Error.Print(err_str)
		os.Exit(1)
	}
	if len(log.Topics) < 2 {
		err_str := fmt.Sprintf("Event ERC721 MintEvent doesn't have 3 topics")
		Error.Print(err_str)
		Info.Print(err_str)
		Info.Printf("%+v", log)
		Error.Printf("%+v", log)
		os.Exit(1)
	}
	return log.Topics[1].Big().Int64(), nil
}
func find_prize_num(tx_id int64) int64 {

	evt_list, err := storagew.S.Get_events_by_sig_and_tx_id(tx_id, hex.EncodeToString(evt_prize_claim_event[0:4]))
	if err != nil {

		err_str := fmt.Sprintf("find_prize_num()() error: %v", err)
		Info.Print(err_str)
		Error.Print(err_str)
		os.Exit(1)
	}
	if len(evt_list) == 0 {
		return -1
	}
	if len(evt_list) != 1 {
		err_str := fmt.Sprintf("find_prize_num() there is more than 1 PrizeClaim in this transaction()")
		Info.Print(err_str)
		Error.Print(err_str)
		os.Exit(1)
	}
	var log types.Log
	err = rlp.DecodeBytes(evt_list[0].RlpLog, &log)
	if err != nil {
		err_str := fmt.Sprintf("RLP Decode error at find_prize_num(): %v", err)
		Info.Print(err_str)
		Error.Print(err_str)
		os.Exit(1)
	}
	if len(log.Topics) < 2 {
		err_str := fmt.Sprintf("Event PrizeClaimEvent doesn't have 3 topics")
		Error.Print(err_str)
		Info.Print(err_str)
		Info.Printf("%+v", log)
		Error.Printf("%+v", log)
		os.Exit(1)
	}
	return log.Topics[1].Big().Int64()
}
func admin_uint256_from_log_data(data []byte) (*big.Int, error) {
	if len(data) < 32 {
		return nil, fmt.Errorf("admin event data too short: %d bytes", len(data))
	}
	return new(big.Int).SetBytes(data[len(data)-32:]), nil
}
