// Unit tests (no Docker) for the handler registry: construction succeeds
// (NewRegistry validates the per-topic metric names), the labels feeding
// rwcg_etl_events_total resolve, and every handler declares its emitting
// contracts.
package cosmicgame

import (
	"log/slog"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// unitContracts returns a contract set with distinct synthetic addresses, so
// source-filter registrations stay distinguishable in unit tests.
func unitContracts() Contracts {
	a := func(b byte) ethcommon.Address { return ethcommon.Address{19: b} }
	return Contracts{
		Game:            a(1),
		Signature:       a(2),
		Token:           a(3),
		Dao:             a(4),
		CharityWallet:   a(5),
		PrizesWallet:    a(6),
		StakingCST:      a(7),
		StakingRWalk:    a(8),
		MarketingWallet: a(9),
		Implementation:  a(10),
		CosmicTokenAid:  3,
	}
}

// newUnitHandlers builds a Handlers set over inert dependencies: good enough
// for registry construction and pure decode calls, which touch neither the
// database nor the chain.
func newUnitHandlers(t testing.TB) *Handlers {
	t.Helper()
	st := &store.Store{}
	h, err := New(Config{
		Repo:      cgstore.NewRepo(st),
		Store:     st,
		Caller:    inertCaller{},
		Contracts: unitContracts(),
		Logger:    slog.New(slog.DiscardHandler),
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return h
}

func TestRegistryBuildsAndResolvesNames(t *testing.T) {
	h := newUnitHandlers(t)
	reg := h.Registry()

	// One registration per legacy dispatch row (75), plus one: the single
	// Transfer row split into the ERC721 and ERC20 handlers. The two
	// CharityAddressChanged rows were already separate registrations.
	if got := len(reg.Handlers()); got != 76 {
		t.Errorf("registered handlers = %d, want 76", got)
	}

	for _, c := range []struct{ topic, want string }{
		{TopicBidEvent, "BidPlaced"},
		{TopicBidEventV2, "BidPlacedV2"},
		{TopicMintEvent, "NftMinted"},
		{TopicTransferEvt, "Transfer"},
		// The two shared-signature pairs resolve to one label each
		// (TopicCharityWalletChanged == TopicCharityReceiverChanged and
		// TopicFundsToCharity == TopicDonationSentEvent by construction).
		{TopicCharityWalletChanged, "CharityAddressChanged"},
		{TopicFundsToCharity, "FundsTransferredToCharity"},
	} {
		if got := reg.TopicName(topicHash(c.topic)); got != c.want {
			t.Errorf("TopicName(%s) = %q, want %q", c.topic, got, c.want)
		}
	}
	if got := reg.TopicName(ethcommon.HexToHash("0xdead")); got != "" {
		t.Errorf("TopicName(unknown) = %q, want empty", got)
	}
}

func TestEveryHandlerDeclaresSources(t *testing.T) {
	// FilterLogs fetches whole contracts, so every CosmicGame handler must
	// scope itself to at least one emitting contract; an empty source set
	// (accept-any) would process foreign events.
	for _, hd := range newUnitHandlers(t).Registry().Handlers() {
		if len(hd.Sources()) == 0 {
			t.Errorf("handler %s (%s) declares no source contracts", hd.Name(), hd.Topic().Hex())
		}
		for _, src := range hd.Sources() {
			if src == (ethcommon.Address{}) {
				t.Errorf("handler %s (%s) declares the zero address as a source", hd.Name(), hd.Topic().Hex())
			}
		}
	}
}

func TestOwnershipContractCodes(t *testing.T) {
	h := newUnitHandlers(t)
	c := h.c
	want := map[ethcommon.Address]int64{
		c.Game:            1,
		c.Signature:       2,
		c.Token:           3,
		c.CharityWallet:   4,
		c.PrizesWallet:    5,
		c.StakingCST:      6,
		c.StakingRWalk:    7,
		c.MarketingWallet: 8,
		c.Dao:             9,
		c.Implementation:  0, // not in the ownership set
	}
	for addr, code := range want {
		if got := h.ownershipContractCode(addr); got != code {
			t.Errorf("ownershipContractCode(%v) = %d, want %d", addr, got, code)
		}
	}
	if got := len(h.ownershipSources()); got != 9 {
		t.Errorf("ownership sources = %d, want 9", got)
	}
	if got := len(h.initializedSources()); got != 10 {
		t.Errorf("initialized sources = %d, want 10", got)
	}
}
