// Unit tests (no Docker) for the ChaosZero handler registry: the topic-hash
// constants are pinned to their ABI-derived event IDs, construction succeeds
// (NewRegistry validates the per-topic metric names), the labels feeding
// rwcg_etl_events_total resolve, every handler declares its emitting
// contract, and the decoders round-trip ABI-packed payloads.
package chaoszero

import (
	"log/slog"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	czc "github.com/PredictionExplorer/augur-explorer/contracts/chaoszero"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	czstore "github.com/PredictionExplorer/augur-explorer/internal/store/chaoszero"
)

func mustABI(tb testing.TB, abiJSON string) *abi.ABI {
	tb.Helper()
	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		tb.Fatalf("parsing ABI: %v", err)
	}
	return &parsed
}

// newUnitHandlers builds a Handlers set over inert dependencies: good enough
// for registry construction and pure decode calls.
func newUnitHandlers(tb testing.TB) *Handlers {
	tb.Helper()
	h, err := New(Config{
		Repo: czstore.NewRepo(nil),
		Contracts: Contracts{
			Market: ethcommon.Address{19: 1},
			Vault:  ethcommon.Address{19: 2},
		},
		Logger: slog.New(slog.DiscardHandler),
	})
	if err != nil {
		tb.Fatalf("New: %v", err)
	}
	return h
}

func TestRegistryConstantsMatchABIEventIDs(t *testing.T) {
	market := mustABI(t, czc.MarketABI)
	vault := mustABI(t, czc.VaultABI)

	cases := []struct {
		constant string
		abi      *abi.ABI
		event    string
	}{
		{TopicRoundInitialized, market, "RoundInitialized"},
		{TopicRoundFunded, market, "RoundFunded"},
		{TopicTrade, market, "Trade"},
		{TopicPositionSplit, market, "PositionSplit"},
		{TopicPositionsMerged, market, "PositionsMerged"},
		{TopicRoundSettled, market, "RoundSettled"},
		{TopicClaimed, market, "Claimed"},
		{TopicAnswersObserved, market, "AnswersObserved"},
		{TopicTransferSingle, market, "TransferSingle"},
		{TopicTransferBatch, market, "TransferBatch"},
		{TopicApprovalForAll, market, "ApprovalForAll"},
		{TopicCycleEscrowed, vault, "CycleEscrowed"},
		{TopicPaid, vault, "Paid"},
	}
	for _, c := range cases {
		ev, ok := c.abi.Events[c.event]
		if !ok {
			t.Errorf("event %s missing from ABI", c.event)
			continue
		}
		if got := ev.ID.Hex()[2:]; got != c.constant {
			t.Errorf("event %s: ABI ID %s, constant %s", c.event, got, c.constant)
		}
	}
}

func TestRegistryBuildsAndResolvesNames(t *testing.T) {
	reg := newUnitHandlers(t).Registry()

	// Eight market events, two vault events and the three OpenZeppelin
	// ERC-1155 events the market inherits.
	if got := len(reg.Handlers()); got != 13 {
		t.Errorf("registered handlers = %d, want 13", got)
	}
	for topic, want := range map[string]string{
		TopicRoundInitialized: "RoundInitialized",
		TopicRoundFunded:      "RoundFunded",
		TopicTrade:            "Trade",
		TopicPositionSplit:    "PositionSplit",
		TopicPositionsMerged:  "PositionsMerged",
		TopicRoundSettled:     "RoundSettled",
		TopicClaimed:          "Claimed",
		TopicAnswersObserved:  "AnswersObserved",
		TopicCycleEscrowed:    "CycleEscrowed",
		TopicPaid:             "Paid",
		TopicTransferSingle:   "TransferSingle",
		TopicTransferBatch:    "TransferBatch",
		TopicApprovalForAll:   "ApprovalForAll",
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

// packEvent ABI-packs the non-indexed inputs of one event.
func packEvent(tb testing.TB, contractABI *abi.ABI, event string, vals ...any) []byte {
	tb.Helper()
	data, err := contractABI.Events[event].Inputs.NonIndexed().Pack(vals...)
	if err != nil {
		tb.Fatalf("packing %s: %v", event, err)
	}
	return data
}

func unitElog() *store.EthereumEventLog {
	return &store.EthereumEventLog{EvtID: 7, BlockNum: 517037400, TxID: 9, TimeStamp: 1760000000}
}

func TestDecodeTrade(t *testing.T) {
	h := newUnitHandlers(t)
	user := ethcommon.Address{19: 0xaa}
	lg := &types.Log{
		Address: h.c.Market,
		Topics: []ethcommon.Hash{
			topicHash(TopicTrade),
			ethcommon.BigToHash(big.NewInt(42)), // roundId
			ethcommon.BytesToHash(user.Bytes()), // user
			ethcommon.BigToHash(big.NewInt(5)),  // question
		},
		Data: packEvent(t, h.marketABI, "Trade", true, false, big.NewInt(1000), big.NewInt(900)),
	}
	evt, err := h.decodeTrade(lg, unitElog())
	if err != nil {
		t.Fatalf("decodeTrade: %v", err)
	}
	if evt.RoundId != 42 || evt.Question != 5 || !evt.Yes || evt.Buy ||
		evt.Cash != "1000" || evt.Units != "900" || evt.User != user.String() {
		t.Errorf("decoded Trade = %+v", evt)
	}
}

func TestDecodeRoundSettledUnpacksMask(t *testing.T) {
	h := newUnitHandlers(t)
	escrowed := [8]*big.Int{}
	for i := range escrowed {
		escrowed[i] = big.NewInt(int64(i + 1))
	}
	lg := &types.Log{
		Address: h.c.Market,
		Topics: []ethcommon.Hash{
			topicHash(TopicRoundSettled),
			ethcommon.BigToHash(big.NewInt(3)),
		},
		Data: packEvent(t, h.marketABI, "RoundSettled", uint8(0b10100101), escrowed, big.NewInt(777)),
	}
	evt, err := h.decodeRoundSettled(lg, unitElog())
	if err != nil {
		t.Fatalf("decodeRoundSettled: %v", err)
	}
	if evt.ValuesMask != 0b10100101 || evt.Carried != "777" {
		t.Errorf("decoded RoundSettled = %+v", evt)
	}
	wantAnswers := []bool{true, false, true, false, false, true, false, true}
	for q, want := range wantAnswers {
		if evt.Answers[q] != want {
			t.Errorf("Answers[%d] = %v, want %v", q, evt.Answers[q], want)
		}
	}
	if evt.Escrowed[0] != "1" || evt.Escrowed[7] != "8" {
		t.Errorf("Escrowed = %v", evt.Escrowed)
	}
}

func TestDecodeTransferBatchDecodesPositionIds(t *testing.T) {
	h := newUnitHandlers(t)
	// positionId = (roundId << 4) | (question << 1) | yes.
	ids := []*big.Int{
		big.NewInt(42<<4 | 3<<1 | 1), // round 42, question 3, YES
		big.NewInt(42<<4 | 7<<1 | 0), // round 42, question 7, NO
	}
	values := []*big.Int{big.NewInt(100), big.NewInt(200)}
	operator := ethcommon.Address{19: 0x01}
	from := ethcommon.Address{}
	to := ethcommon.Address{19: 0x02}
	lg := &types.Log{
		Address: h.c.Market,
		Topics: []ethcommon.Hash{
			topicHash(TopicTransferBatch),
			ethcommon.BytesToHash(operator.Bytes()),
			ethcommon.BytesToHash(from.Bytes()),
			ethcommon.BytesToHash(to.Bytes()),
		},
		Data: packEvent(t, h.marketABI, "TransferBatch", ids, values),
	}
	evt, err := h.decodeTransferBatch(lg, unitElog())
	if err != nil {
		t.Fatalf("decodeTransferBatch: %v", err)
	}
	if len(evt.Items) != 2 {
		t.Fatalf("items = %d, want 2", len(evt.Items))
	}
	first, second := evt.Items[0], evt.Items[1]
	if first.RoundId != 42 || first.Question != 3 || !first.Yes || first.Amount != "100" {
		t.Errorf("item 0 = %+v", first)
	}
	if second.RoundId != 42 || second.Question != 7 || second.Yes || second.Amount != "200" {
		t.Errorf("item 1 = %+v", second)
	}
}

func TestDecodeRejectsOverflowingRoundID(t *testing.T) {
	h := newUnitHandlers(t)
	huge := new(big.Int).Lsh(big.NewInt(1), 100)
	lg := &types.Log{
		Address: h.c.Market,
		Topics: []ethcommon.Hash{
			topicHash(TopicRoundInitialized),
			ethcommon.BigToHash(huge),
		},
	}
	if _, err := h.decodeRoundInitialized(lg, unitElog()); err == nil {
		t.Error("decodeRoundInitialized accepted an int64-overflowing roundId")
	}
}
