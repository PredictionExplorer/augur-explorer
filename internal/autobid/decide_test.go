package autobid

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var (
	me    = common.HexToAddress("0x00000000000000000000000000000000000000aa")
	other = common.HexToAddress("0x00000000000000000000000000000000000000bb")
)

// eth converts a float ETH amount to wei for test readability.
func eth(f float64) *big.Int {
	wei := new(big.Float).Mul(big.NewFloat(f), big.NewFloat(1e18))
	out, _ := wei.Int(nil)
	return out
}

// baseLimits is the default operator configuration used by the tests: 5 ETH
// max bid, 9 CST max price, RWalk above 0.1 ETH, bid 15s before the prize.
func baseLimits() Limits {
	return Limits{
		MaxEthBid:       eth(5),
		MaxCstBid:       eth(9),
		RWalkMinPrice:   eth(0.1),
		TimeBeforePrize: 15,
		CstBidAnyway:    true,
	}
}

// baseMarket is a market snapshot where nothing is urgent: someone else is
// the last bidder, the prize is far away, and every price is payable.
func baseMarket() Market {
	return Market{
		CstPrice:       eth(20), // above MaxCstBid: no CST bidding by default
		EthBidPrice:    eth(0.05),
		TimeUntilPrize: big.NewInt(1000),
		CstBalance:     eth(100),
		EthBalance:     eth(50),
		LastBidder:     other,
		RWalkMintPrice: eth(0.02),
	}
}

func TestDecideTable(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(*Market, *Limits)
		want   Action
	}{
		{
			name:   "far from prize, nothing to do",
			mutate: func(m *Market, l *Limits) {},
			want:   ActionWait,
		},
		{
			name: "cheap CST bids anytime when not last bidder",
			mutate: func(m *Market, l *Limits) {
				m.CstPrice = eth(5) // below MaxCstBid 9
			},
			want: ActionBidCST,
		},
		{
			name: "cheap CST not taken as the round's first bid",
			mutate: func(m *Market, l *Limits) {
				m.CstPrice = eth(5)
				m.LastBidder = common.Address{} // no bids yet
				m.TimeUntilPrize = big.NewInt(1000)
			},
			want: ActionWait,
		},
		{
			name: "first bid of round must be ETH even when CST is cheap",
			mutate: func(m *Market, l *Limits) {
				m.CstPrice = eth(5)
				m.LastBidder = common.Address{}
				m.TimeUntilPrize = big.NewInt(10) // time to bid
			},
			want: ActionBidETH,
		},
		{
			name: "cheap CST needs the balance to cover the configured max",
			mutate: func(m *Market, l *Limits) {
				m.CstPrice = eth(5)
				m.CstBalance = eth(6) // below MaxCstBid 9
				m.TimeUntilPrize = big.NewInt(1000)
			},
			want: ActionWait,
		},
		{
			name: "CST anyway disabled waits until bid time",
			mutate: func(m *Market, l *Limits) {
				m.CstPrice = eth(5)
				l.CstBidAnyway = false
			},
			want: ActionWait,
		},
		{
			name: "time to bid, CST preferred when cheap even with CstBidAnyway off",
			mutate: func(m *Market, l *Limits) {
				m.CstPrice = eth(5)
				l.CstBidAnyway = false
				m.TimeUntilPrize = big.NewInt(15) // == TimeBeforePrize
			},
			want: ActionBidCST,
		},
		{
			name: "time to bid, plain ETH below RWalk floor",
			mutate: func(m *Market, l *Limits) {
				m.TimeUntilPrize = big.NewInt(10)
				m.EthBidPrice = eth(0.05) // below RWalkMinPrice 0.1
			},
			want: ActionBidETH,
		},
		{
			name: "RWalk wins when mint + half price beats plain ETH",
			mutate: func(m *Market, l *Limits) {
				m.TimeUntilPrize = big.NewInt(10)
				m.EthBidPrice = eth(1)      // above the 0.1 floor
				m.RWalkMintPrice = eth(0.1) // 0.1 + 0.5 = 0.6 < 1
			},
			want: ActionBidRWalk,
		},
		{
			name: "plain ETH wins when the mint makes RWalk more expensive",
			mutate: func(m *Market, l *Limits) {
				m.TimeUntilPrize = big.NewInt(10)
				m.EthBidPrice = eth(1)
				m.RWalkMintPrice = eth(0.6) // 0.6 + 0.5 = 1.1 > 1
			},
			want: ActionBidETH,
		},
		{
			name: "RWalk total above MaxEthBid waits",
			mutate: func(m *Market, l *Limits) {
				m.TimeUntilPrize = big.NewInt(10)
				m.EthBidPrice = eth(12)   // plain ETH already above MaxEthBid 5
				m.RWalkMintPrice = eth(3) // 3 + 6 = 9 > 5
			},
			want: ActionWait,
		},
		{
			name: "RWalk cheaper and within budget while plain ETH is over it",
			mutate: func(m *Market, l *Limits) {
				m.TimeUntilPrize = big.NewInt(10)
				m.EthBidPrice = eth(6)      // above MaxEthBid 5
				m.RWalkMintPrice = eth(0.5) // 0.5 + 3 = 3.5 < 5
			},
			want: ActionBidRWalk,
		},
		{
			name: "ETH price above limit waits",
			mutate: func(m *Market, l *Limits) {
				m.TimeUntilPrize = big.NewInt(10)
				m.EthBidPrice = eth(0.05)
				l.MaxEthBid = eth(0.01)
			},
			want: ActionWait,
		},
		{
			name: "insufficient ETH balance waits",
			mutate: func(m *Market, l *Limits) {
				m.TimeUntilPrize = big.NewInt(10)
				m.EthBidPrice = eth(0.05)
				m.EthBalance = eth(0.05) // price >= balance
			},
			want: ActionWait,
		},
		{
			name: "last bidder waits for the timer",
			mutate: func(m *Market, l *Limits) {
				m.LastBidder = me
				m.TimeUntilPrize = big.NewInt(30)
			},
			want: ActionWait,
		},
		{
			name: "last bidder claims at zero",
			mutate: func(m *Market, l *Limits) {
				m.LastBidder = me
				m.TimeUntilPrize = big.NewInt(0)
			},
			want: ActionClaimPrize,
		},
		{
			name: "last bidder keeps bidding cheap CST instead of claiming",
			mutate: func(m *Market, l *Limits) {
				m.LastBidder = me
				m.TimeUntilPrize = big.NewInt(0)
				m.CstPrice = eth(5)
			},
			want: ActionBidCST,
		},
		{
			name: "last bidder with CstBidAnyway off claims at zero",
			mutate: func(m *Market, l *Limits) {
				m.LastBidder = me
				m.TimeUntilPrize = big.NewInt(0)
				m.CstPrice = eth(5)
				l.CstBidAnyway = false
			},
			want: ActionClaimPrize,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := baseMarket()
			m.EthBidPrice = eth(0.05)
			l := baseLimits()
			tc.mutate(&m, &l)
			got, reason := Decide(m, l, me)
			if got != tc.want {
				t.Errorf("Decide = %s (%s), want %s", got, reason, tc.want)
			}
			if reason == "" {
				t.Error("Decide returned an empty reason")
			}
		})
	}
}

func TestCanBidWithCSTBoundaries(t *testing.T) {
	l := baseLimits() // MaxCstBid 9
	m := baseMarket()

	m.CstPrice = eth(9)
	m.CstBalance = eth(9)
	if !canBidWithCST(m, l) {
		t.Error("price == max, balance == max should allow CST bidding")
	}
	m.CstPrice = eth(9.000001)
	if canBidWithCST(m, l) {
		t.Error("price just above max should not allow CST bidding")
	}
	m.CstPrice = eth(9)
	m.CstBalance = eth(8.999999)
	if canBidWithCST(m, l) {
		t.Error("balance below the configured max should not allow CST bidding")
	}
}

func TestTimeoutClaimExpired(t *testing.T) {
	prizeTime := big.NewInt(1000)
	timeout := big.NewInt(100)
	if timeoutClaimExpired(prizeTime, timeout, big.NewInt(1100)) {
		t.Error("deadline == timestamp should not be expired (strict inequality)")
	}
	if !timeoutClaimExpired(prizeTime, timeout, big.NewInt(1101)) {
		t.Error("timestamp past the deadline should be expired")
	}
	if timeoutClaimExpired(prizeTime, timeout, big.NewInt(500)) {
		t.Error("timestamp before the deadline should not be expired")
	}
}

func TestActionAndPendingKindStrings(t *testing.T) {
	for a, want := range map[Action]string{
		ActionWait:       "Wait",
		ActionBidCST:     "BidCST",
		ActionBidETH:     "BidETH",
		ActionBidRWalk:   "BidRWalk",
		ActionClaimPrize: "ClaimPrize",
		Action(99):       "Unknown(99)",
	} {
		if got := a.String(); got != want {
			t.Errorf("Action(%d).String() = %q, want %q", int(a), got, want)
		}
	}
	for k, want := range map[pendingTxKind]string{
		pendingNone:      "None",
		pendingCSTBid:    "CSTBid",
		pendingETHBid:    "ETHBid",
		pendingRWalkMint: "RWalkMint",
		pendingRWalkBid:  "RWalkBid",
		pendingClaim:     "Claim",
		pendingTxKind(9): "Unknown(9)",
	} {
		if got := k.String(); got != want {
			t.Errorf("pendingTxKind(%d).String() = %q, want %q", int(k), got, want)
		}
	}
}
