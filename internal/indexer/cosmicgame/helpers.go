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

type bidRewardMint struct {
	to     ethcommon.Address
	amount *big.Int
}

// classifyBidRewardMints attributes the CST mints of a bid transaction by
// recipient: V1/V2 mint the bid reward to the bidder placing the bid, while
// V3 mints the whole reward to the outbid previous bidder (and nothing on a
// round's first bid). A transaction can carry more than one mint log for the
// same recipient (e.g. a zero-price CST bid emits a zero-amount mint next to
// the reward mint), so amounts are summed per recipient.
func classifyBidRewardMints(mints []bidRewardMint, bidderAddr ethcommon.Address) (thisReward, previousReward, previousAddr string) {
	this := new(big.Int)
	previous := new(big.Int)
	var previousTo ethcommon.Address
	for i := range mints {
		if mints[i].to == bidderAddr {
			this.Add(this, mints[i].amount)
		} else {
			previous.Add(previous, mints[i].amount)
			previousTo = mints[i].to
		}
	}
	if previousTo == (ethcommon.Address{}) {
		return this.String(), "0", ""
	}
	return this.String(), previous.String(), previousTo.String()
}

// cstBidRewards derives the reward attribution from the CST mint Transfer
// logs preceding the bid. V1/V2 mint to the current bidder; V3 mints the
// whole reward to the outbid previous bidder (and none on a round's first
// bid). The mint recipients disambiguate the shapes.
func (h *Handlers) cstBidRewards(
	ctx context.Context,
	bidEvtlogID, txID int64,
	bidderAddr string,
) (thisReward, previousReward, previousAddr string, err error) {
	elogRLPs, err := h.store.EventLogRLPsBefore(ctx, txID, h.c.CosmicTokenAid, bidEvtlogID, TopicTransferEvt[:8])
	if err != nil {
		return "", "", "", fmt.Errorf("cstBidRewards(): %w", err)
	}
	mints := make([]bidRewardMint, 0, 2)
	for _, raw := range elogRLPs {
		var lg types.Log
		if err := rlp.DecodeBytes(raw, &lg); err != nil {
			return "", "", "", fmt.Errorf("cstBidRewards(): RLP decode: %w", err)
		}
		if len(lg.Topics) < 3 || ethcommon.BytesToAddress(lg.Topics[1][12:]) != (ethcommon.Address{}) {
			continue
		}
		var transfer cgc.ERC20Transfer
		if err := h.erc20ABI.UnpackIntoInterface(&transfer, "Transfer", lg.Data); err != nil {
			return "", "", "", fmt.Errorf("cstBidRewards(): Transfer decode: %w", err)
		}
		amount := new(big.Int).Set(transfer.Value)
		mints = append(mints, bidRewardMint{
			to:     ethcommon.BytesToAddress(lg.Topics[2][12:]),
			amount: amount,
		})
	}
	thisReward, previousReward, previousAddr = classifyBidRewardMints(mints, ethcommon.HexToAddress(bidderAddr))
	return thisReward, previousReward, previousAddr, nil
}

// prizeRoundInTx returns the round number of the MainPrizeClaimed event in
// txID, or -1 when the transaction contains none (a standalone donation).
func (h *Handlers) prizeRoundInTx(ctx context.Context, txID int64) (int64, error) {
	evtList := make([]store.EthereumEventLog, 0, 1)
	for _, topic := range []string{TopicPrizeClaimEvent, TopicPrizeClaimEventV3} {
		events, err := h.store.EventsBySigAndTx(ctx, txID, topic[:8])
		if err != nil {
			return 0, fmt.Errorf("prizeRoundInTx(): %w", err)
		}
		evtList = append(evtList, events...)
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
