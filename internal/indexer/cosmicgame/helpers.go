// Enrichment helpers the store steps share: locating related events of the
// same transaction (CST reward mints, prize claims, bids), reading contract
// state (donation-info records, donated-NFT token URIs) and decoding the raw
// data words of legacy events no generated ABI defines.

package cosmicgame

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// cstBidReward returns the CST bid reward (the amount minted to the bidder)
// for this bid transaction.
//
// The reward is located by MATCHING the correct Transfer event rather than by
// positional offset from the bid event. Per bid the only CST mint is the bid
// reward credited to the bidder, i.e. an ERC20 Transfer with from == zero
// address (mint) and to == bidder. (A CST bid also burns the paid price,
// emitting Transfer(bidder -> 0x0); that is ignored by this match. The
// marketing-wallet CST contribution is minted at round end, not per bid.)
//
// In V2 the reward is dynamic (BiddingV2.getBidCstRewardAmountAdvanced) and
// is 0 for the earliest bid(s) of a round. When it is 0 the contract performs
// no mint, so there is no matching Transfer and "0" is correctly reported
// instead of failing.
func (h *Handlers) cstBidReward(ctx context.Context, bidEvtlogID, txID int64, bidderAddr string) (string, error) {
	elogRLPs, err := h.store.EventLogRLPsBefore(ctx, txID, h.c.CosmicTokenAid, bidEvtlogID, TRANSFER_EVT[:8])
	if err != nil {
		return "", fmt.Errorf("cstBidReward(): %w", err)
	}
	bidder := ethcommon.HexToAddress(bidderAddr)
	for _, raw := range elogRLPs {
		var lg types.Log
		if err := rlp.DecodeBytes(raw, &lg); err != nil {
			// A stored evt_log row that fails to decode is an invalid
			// database state; abort the batch instead of guessing.
			return "", fmt.Errorf("cstBidReward(): RLP decode: %w", err)
		}
		// ERC20 Transfer indexes `from` and `to`, so they live in the topics, not in Data.
		if len(lg.Topics) < 3 {
			continue
		}
		from := ethcommon.BytesToAddress(lg.Topics[1][12:])
		to := ethcommon.BytesToAddress(lg.Topics[2][12:])
		if from != (ethcommon.Address{}) || to != bidder {
			continue // not the bid reward mint to the bidder
		}
		var ethEvt cgc.ERC20Transfer
		if err := h.erc20ABI.UnpackIntoInterface(&ethEvt, "Transfer", lg.Data); err != nil {
			return "", fmt.Errorf("cstBidReward(): Transfer decode (%+v): %w", lg, err)
		}
		return ethEvt.Value.String(), nil
	}
	// No mint to the bidder in this transaction => the dynamic CST bid reward was 0.
	return "0", nil
}

// prizeRoundInTx returns the round number of the MainPrizeClaimed event in
// txID, or -1 when the transaction contains none (a standalone donation).
func (h *Handlers) prizeRoundInTx(ctx context.Context, txID int64) (int64, error) {
	evtList, err := h.store.EventsBySigAndTx(ctx, txID, PRIZE_CLAIM_EVENT[:8])
	if err != nil {
		return 0, fmt.Errorf("prizeRoundInTx(): %w", err)
	}
	if len(evtList) == 0 {
		return -1, nil
	}
	if len(evtList) != 1 {
		return 0, fmt.Errorf("prizeRoundInTx(): more than one PrizeClaim in tx %v", txID)
	}
	var lg types.Log
	if err := rlp.DecodeBytes(evtList[0].RlpLog, &lg); err != nil {
		return 0, fmt.Errorf("prizeRoundInTx(): RLP decode: %w", err)
	}
	if len(lg.Topics) < 2 {
		return 0, fmt.Errorf("prizeRoundInTx(): PrizeClaim event has %d topics, want at least 2 (%+v)", len(lg.Topics), lg)
	}
	return lg.Topics[1].Big().Int64(), nil
}

// bidIDByEvtlog adapts Repo.BidIDByEvtlog to the handler contract: -1 when
// the event log carries no bid, real DB errors propagate.
func (h *Handlers) bidIDByEvtlog(ctx context.Context, evtlogID int64) (int64, error) {
	id, err := h.repo.BidIDByEvtlog(ctx, evtlogID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return -1, nil
		}
		return 0, fmt.Errorf("BidIDByEvtlog(%v): %w", evtlogID, err)
	}
	return id, nil
}

// fetchDonationInfo reads the JSON data of one EthDonationWithInfo record
// from the game contract.
func (h *Handlers) fetchDonationInfo(ctx context.Context, recordID int64) (string, error) {
	game, err := cgc.NewCosmicSignatureGameCaller(h.c.Game, h.caller)
	if err != nil {
		return "", err
	}
	rec, err := game.EthDonationWithInfoRecords(&bind.CallOpts{Context: ctx}, big.NewInt(recordID))
	if err != nil {
		return "", err
	}
	return rec.Data, nil
}

// fetchTokenURI reads tokenURI for a donated NFT. A failing contract call
// yields "" (the donation is stored without a URI, exactly like the legacy
// path); only a broken contract binding is an error.
func (h *Handlers) fetchTokenURI(ctx context.Context, tokenID int64, contractAddr ethcommon.Address) (string, error) {
	c, err := cgc.NewCosmicSignatureNftCaller(contractAddr, h.caller) // CosmicSignatureNft doubles as the generic ERC721 binding
	if err != nil {
		return "", fmt.Errorf("fetchTokenURI(): contract creation: %w", err)
	}
	tokURI, err := c.TokenURI(&bind.CallOpts{Context: ctx}, big.NewInt(tokenID))
	if err != nil {
		h.log.Error("fetchTokenURI: tokenURI call failed", "token_id", tokenID, "contract", contractAddr.String(), "err", err)
		return "", nil
	}
	return tokURI, nil
}

// adminUint256FromLogData extracts the single non-indexed uint256 of a
// legacy admin event whose signature no current ABI defines.
func adminUint256FromLogData(data []byte) (*big.Int, error) {
	if len(data) < 32 {
		return nil, fmt.Errorf("admin event data too short: %d bytes", len(data))
	}
	return new(big.Int).SetBytes(data[len(data)-32:]), nil
}

// erc20TransferFailedAmount extracts the uint256 amount from an
// ERC20TransferFailed(string errStr, address indexed destination,
// uint256 amount) log body. The event lives in ICosmicSignatureErrors.sol
// and is absent from every generated ABI, so the head words are decoded
// manually: word 0 is the offset of the string tail, word 1 the amount.
func erc20TransferFailedAmount(data []byte) (*big.Int, error) {
	if len(data) < 64 {
		return nil, fmt.Errorf("ERC20TransferFailed data too short: %d bytes", len(data))
	}
	return new(big.Int).SetBytes(data[32:64]), nil
}
