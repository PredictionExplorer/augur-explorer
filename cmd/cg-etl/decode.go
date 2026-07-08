// Shared decode/lookup helpers: RLP event-log decoding and locating related
// events (token transfers, prize claims) within a transaction.
package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

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
	elog_rlps, err := storage.Get_specific_event_logs_by_tx_backwards_from_id(tx_id, cosmic_tok_aid, bid_evtlog_id, hex.EncodeToString(evt_transfer[:4]))
	if err != nil {
		return "", fmt.Errorf("find_cosmic_token_transfer(): %w", err)
	}
	bidder := common.HexToAddress(bidder_addr)
	for _, raw := range elog_rlps {
		var log types.Log
		err := rlp.DecodeBytes(raw, &log)
		if err != nil {
			// A stored evt_log row that fails to decode is an invalid
			// database state; abort the batch instead of guessing.
			return "", fmt.Errorf("find_cosmic_token_transfer(): RLP decode: %w", err)
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
			return "", fmt.Errorf("find_cosmic_token_transfer(): Transfer decode (%+v): %w", log, err)
		}
		return eth_evt.Value.String(), nil
	}
	// No mint to the bidder in this transaction => the dynamic CST bid reward was 0.
	return "0", nil
}

// find_prize_num returns the round number of the MainPrizeClaimed event in
// tx_id, or -1 when the transaction contains none (a standalone donation).
func find_prize_num(tx_id int64) (int64, error) {
	evt_list, err := storage.Get_events_by_sig_and_tx_id(tx_id, hex.EncodeToString(evt_prize_claim_event[0:4]))
	if err != nil {
		return 0, fmt.Errorf("find_prize_num(): %w", err)
	}
	if len(evt_list) == 0 {
		return -1, nil
	}
	if len(evt_list) != 1 {
		return 0, fmt.Errorf("find_prize_num(): more than one PrizeClaim in tx %v", tx_id)
	}
	var log types.Log
	err = rlp.DecodeBytes(evt_list[0].RlpLog, &log)
	if err != nil {
		return 0, fmt.Errorf("find_prize_num(): RLP decode: %w", err)
	}
	if len(log.Topics) < 2 {
		return 0, fmt.Errorf("find_prize_num(): PrizeClaim event has %d topics, want at least 2 (%+v)", len(log.Topics), log)
	}
	return log.Topics[1].Big().Int64(), nil
}

// admin_uint256_from_log_data extracts the single non-indexed uint256 of a
// legacy admin event whose signature no current ABI defines.
func admin_uint256_from_log_data(data []byte) (*big.Int, error) {
	if len(data) < 32 {
		return nil, fmt.Errorf("admin event data too short: %d bytes", len(data))
	}
	return new(big.Int).SetBytes(data[len(data)-32:]), nil
}

// erc20_transfer_failed_amount extracts the uint256 amount from an
// ERC20TransferFailed(string errStr, address indexed destination,
// uint256 amount) log body. The event lives in ICosmicSignatureErrors.sol
// and is absent from every generated ABI, so the head words are decoded
// manually: word 0 is the offset of the string tail, word 1 the amount.
func erc20_transfer_failed_amount(data []byte) (*big.Int, error) {
	if len(data) < 64 {
		return nil, fmt.Errorf("ERC20TransferFailed data too short: %d bytes", len(data))
	}
	return new(big.Int).SetBytes(data[32:64]), nil
}
