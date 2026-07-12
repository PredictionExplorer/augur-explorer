package autobid

// The pure decision core of the bidding bot: given one snapshot of market
// state and the operator's limits, decide what to do next. No RPC, no
// state — every rule the legacy bot applied is a total function here so the
// whole decision table is unit-testable.

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Market is one refresh of the contract and account state the bidding
// decisions read.
type Market struct {
	// CstPrice is the next CST bid price (wei-denominated CST).
	CstPrice *big.Int
	// EthBidPrice is the next ETH bid price in wei.
	EthBidPrice *big.Int
	// TimeUntilPrize is the seconds until the main prize can be claimed.
	TimeUntilPrize *big.Int
	// CstBalance is the bot's CST token balance.
	CstBalance *big.Int
	// EthBalance is the bot's ETH balance in wei.
	EthBalance *big.Int
	// LastBidder is the current last bidder (zero address = no bids yet).
	LastBidder common.Address
	// RWalkMintPrice is the current RandomWalk mint price in wei.
	RWalkMintPrice *big.Int
}

// Limits are the operator-configured bidding limits.
type Limits struct {
	// MaxEthBid is the most wei a single ETH bid may cost.
	MaxEthBid *big.Int
	// MaxCstBid is the highest CST price the bot will pay.
	MaxCstBid *big.Int
	// RWalkMinPrice is the ETH bid price above which RandomWalk bidding is
	// considered.
	RWalkMinPrice *big.Int
	// TimeBeforePrize is how many seconds before the prize deadline the bot
	// starts bidding.
	TimeBeforePrize int64
	// CstBidAnyway keeps bidding with cheap CST even while last bidder.
	CstBidAnyway bool
}

// Action is what the decision core wants the engine to do next.
type Action int

// Decision actions in priority order.
const (
	// ActionWait does nothing this iteration.
	ActionWait Action = iota
	// ActionBidCST bids with CST tokens.
	ActionBidCST
	// ActionBidETH bids with plain ETH.
	ActionBidETH
	// ActionBidRWalk bids with a RandomWalk token (half ETH price).
	ActionBidRWalk
	// ActionClaimPrize claims the main prize.
	ActionClaimPrize
)

// String names the action for logs.
func (a Action) String() string {
	switch a {
	case ActionWait:
		return "Wait"
	case ActionBidCST:
		return "BidCST"
	case ActionBidETH:
		return "BidETH"
	case ActionBidRWalk:
		return "BidRWalk"
	case ActionClaimPrize:
		return "ClaimPrize"
	default:
		return fmt.Sprintf("Unknown(%d)", int(a))
	}
}

// isZeroAddress reports whether addr is the zero address (no bids yet).
func isZeroAddress(addr common.Address) bool {
	return addr == (common.Address{})
}

// canBidWithCST reports whether the CST bidding conditions are met: the
// balance covers the configured maximum and the price is at or below it.
// (Faithful port of the legacy rule, which compared the balance against the
// configured maximum rather than the current price.)
func canBidWithCST(m Market, l Limits) bool {
	return l.MaxCstBid.Cmp(m.CstBalance) <= 0 &&
		m.CstPrice.Cmp(l.MaxCstBid) <= 0
}

// Decide picks the bot's next action from one market snapshot. It routes on
// whether myAddr currently holds the last-bidder position and reproduces the
// legacy bot's decision order exactly:
//
//	not last bidder: cheap-CST-anyway (never as the round's first bid) →
//	  wait until TimeBeforePrize → cheapest allowed bid method
//	last bidder: cheap-CST-anyway → claim when the timer hits zero → wait
func Decide(m Market, l Limits, myAddr common.Address) (Action, string) {
	if m.LastBidder != myAddr || isZeroAddress(m.LastBidder) {
		// Cheap CST is taken any time — except on the first bid of a round,
		// which the contract requires to be ETH.
		if l.CstBidAnyway && canBidWithCST(m, l) && !isZeroAddress(m.LastBidder) {
			return ActionBidCST, "CST price below limit, bidding with CST"
		}
		if m.TimeUntilPrize.Cmp(big.NewInt(l.TimeBeforePrize)) <= 0 {
			return decideBidType(m, l)
		}
		return ActionWait, fmt.Sprintf("not my time to bid yet (timeUntilPrize = %v)", m.TimeUntilPrize.Int64())
	}
	if l.CstBidAnyway && canBidWithCST(m, l) {
		return ActionBidCST, "CST price low, bidding again"
	}
	if m.TimeUntilPrize.Sign() == 0 {
		return ActionClaimPrize, "prize timer expired, claiming"
	}
	return ActionWait, "waiting for the prize timer"
}

// decideBidType picks the cheapest allowed bid method once it is time to bid.
func decideBidType(m Market, l Limits) (Action, string) {
	// First bid of round must be ETH (contract requirement).
	if isZeroAddress(m.LastBidder) {
		return tryPlainEthBid(m, l, "first bid of round - must use ETH")
	}

	if canBidWithCST(m, l) {
		return ActionBidCST, "CST price below limit, bidding with CST"
	}

	// RandomWalk bidding halves the ETH price but requires minting (or
	// owning) a token; only worthwhile above the configured price floor.
	if l.RWalkMinPrice.Cmp(m.EthBidPrice) < 0 {
		rwalkDiscountedPrice := new(big.Int).Quo(m.EthBidPrice, big.NewInt(2))
		bidWithRwalkPrice := new(big.Int).Add(m.RWalkMintPrice, rwalkDiscountedPrice)

		if m.EthBidPrice.Cmp(bidWithRwalkPrice) <= 0 {
			return tryPlainEthBid(m, l, "plain ETH cheaper than RWalk mint + half price")
		}
		if bidWithRwalkPrice.Cmp(l.MaxEthBid) < 0 {
			return ActionBidRWalk, "bidding with RWalk (cheaper)"
		}
		return ActionWait, "out of funds even with RWalk"
	}

	return tryPlainEthBid(m, l, "bidding with plain ETH")
}

// tryPlainEthBid checks the ETH limits and either bids or waits.
func tryPlainEthBid(m Market, l Limits, reason string) (Action, string) {
	if l.MaxEthBid.Cmp(m.EthBidPrice) < 0 {
		return ActionWait, "ETH bid price above limit"
	}
	if m.EthBidPrice.Cmp(m.EthBalance) >= 0 {
		return ActionWait, "insufficient ETH balance"
	}
	return ActionBidETH, reason
}

// timeoutClaimExpired reports whether the winner's exclusive claim window has
// passed, letting anyone (including the bot) claim the prize.
func timeoutClaimExpired(prizeTime, timeoutDuration, blockTimestamp *big.Int) bool {
	deadline := new(big.Int).Add(prizeTime, timeoutDuration)
	return deadline.Cmp(blockTimestamp) < 0
}
