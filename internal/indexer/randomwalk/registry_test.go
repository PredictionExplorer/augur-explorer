// Unit tests (no Docker) for the handler registry: construction succeeds
// (NewRegistry validates the per-topic metric names), the labels feeding
// rwcg_etl_events_total resolve, and every handler declares its emitting
// contract.
package randomwalk

import (
	"log/slog"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"

	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// newUnitHandlers builds a Handlers set over inert dependencies: good enough
// for registry construction and pure decode calls.
func newUnitHandlers(t testing.TB) *Handlers {
	t.Helper()
	h, err := New(Config{
		Repo: rwstore.NewRepo(nil),
		Contracts: Contracts{
			Market:     ethcommon.Address{19: 1},
			RandomWalk: ethcommon.Address{19: 2},
		},
		Logger: slog.New(slog.DiscardHandler),
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return h
}

func TestRegistryBuildsAndResolvesNames(t *testing.T) {
	reg := newUnitHandlers(t).Registry()

	if got := len(reg.Handlers()); got != 7 {
		t.Errorf("registered handlers = %d, want 7", got)
	}
	for topic, want := range map[string]string{
		TopicNewOffer:      "NewOffer",
		TopicItemBought:    "ItemBought",
		TopicOfferCanceled: "OfferCanceled",
		TopicWithdrawalEvt: "WithdrawalEvent",
		TopicTokenNameEvt:  "TokenNameEvent",
		TopicTransferEvt:   "Transfer",
		TopicMintEvent:     "MintEvent",
	} {
		if got := reg.TopicName(topicHash(topic)); got != want {
			t.Errorf("TopicName(%s) = %q, want %q", topic, got, want)
		}
	}
	if got := reg.TopicName(ethcommon.HexToHash("0xdead")); got != "" {
		t.Errorf("TopicName(unknown) = %q, want empty", got)
	}
}

func TestEveryHandlerDeclaresSources(t *testing.T) {
	for _, hd := range newUnitHandlers(t).Registry().Handlers() {
		if len(hd.Sources()) == 0 {
			t.Errorf("handler %s (%s) declares no source contracts", hd.Name(), hd.Topic().Hex())
		}
	}
}
