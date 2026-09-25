// Unit tests (no Docker): every event the RandomWalkNFT and marketplace ABIs
// declare has a handler admitting that contract, and the inherited
// OpenZeppelin events decode to the expected rows.
package randomwalk

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	rwc "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func parseABI(t *testing.T, raw string) *abi.ABI {
	t.Helper()
	parsed, err := abi.JSON(strings.NewReader(raw))
	if err != nil {
		t.Fatalf("parsing ABI: %v", err)
	}
	return &parsed
}

func TestEveryABIEventHasHandler(t *testing.T) {
	h := newUnitHandlers(t)
	handled := make(map[ethcommon.Hash]map[ethcommon.Address]bool)
	for _, hd := range h.Registry().Handlers() {
		set := handled[hd.Topic()]
		if set == nil {
			set = make(map[ethcommon.Address]bool)
			handled[hd.Topic()] = set
		}
		for _, src := range hd.Sources() {
			set[src] = true
		}
	}
	for _, d := range []struct {
		name string
		abi  string
		addr ethcommon.Address
	}{
		{"RandomWalkNFT", rwc.RWalkABI, h.c.RandomWalk},
		{"RWMarket", rwc.RWMarketABI, h.c.Market},
	} {
		for name, ev := range parseABI(t, d.abi).Events {
			if !handled[ev.ID][d.addr] {
				t.Errorf("%s.%s (%s) has no handler admitting the %s contract", d.name, name, ev.ID.Hex(), d.name)
			}
		}
	}
	for name, constant := range map[string]string{
		"Approval":             TopicApproval,
		"ApprovalForAll":       TopicApprovalForAll,
		"OwnershipTransferred": TopicOwnershipTransferred,
	} {
		if got := parseABI(t, rwc.RWalkABI).Events[name].ID.Hex()[2:]; got != constant {
			t.Errorf("Topic constant for %s = %s, ABI event ID is %s", name, constant, got)
		}
	}
}

func TestDecodeInheritedEvents(t *testing.T) {
	h := newUnitHandlers(t)
	rwalk := parseABI(t, rwc.RWalkABI)
	alice := ethcommon.HexToAddress("0x2100000000000000000000000000000000000021")
	bob := ethcommon.HexToAddress("0x2200000000000000000000000000000000000022")
	meta := &store.EthereumEventLog{EvtID: 9, BlockNum: 8, TimeStamp: 7, TxID: 6}
	topic := func(a ethcommon.Address) ethcommon.Hash { return ethcommon.BytesToHash(a.Bytes()) }

	approval := &types.Log{Address: h.c.RandomWalk, Topics: []ethcommon.Hash{
		rwalk.Events["Approval"].ID, topic(alice), topic(bob), ethcommon.BigToHash(big.NewInt(1234)),
	}}
	a, err := h.decodeApproval(approval, meta)
	if err != nil {
		t.Fatal(err)
	}
	if a.Owner != alice.String() || a.Approved != bob.String() || a.TokenId != 1234 || a.EvtId != 9 {
		t.Fatalf("Approval decoded = %+v", a)
	}
	approval.Topics[3] = ethcommon.BigToHash(new(big.Int).Lsh(big.NewInt(1), 70))
	if _, err := h.decodeApproval(approval, meta); err == nil {
		t.Fatal("tokenId beyond int64 must fail the decode")
	}

	data, err := rwalk.Events["ApprovalForAll"].Inputs.NonIndexed().Pack(true)
	if err != nil {
		t.Fatal(err)
	}
	forAll := &types.Log{Address: h.c.RandomWalk, Topics: []ethcommon.Hash{
		rwalk.Events["ApprovalForAll"].ID, topic(alice), topic(bob),
	}, Data: data}
	fa, err := h.decodeApprovalForAll(forAll, meta)
	if err != nil {
		t.Fatal(err)
	}
	if fa.Owner != alice.String() || fa.Operator != bob.String() || !fa.Approved {
		t.Fatalf("ApprovalForAll decoded = %+v", fa)
	}

	owner := &types.Log{Address: h.c.RandomWalk, Topics: []ethcommon.Hash{
		rwalk.Events["OwnershipTransferred"].ID, topic(alice), topic(bob),
	}}
	o, err := h.decodeOwnershipTransferred(owner, meta)
	if err != nil {
		t.Fatal(err)
	}
	if o.PrevOwner != alice.String() || o.NewOwner != bob.String() {
		t.Fatalf("OwnershipTransferred decoded = %+v", o)
	}
	if _, err := h.decodeOwnershipTransferred(&types.Log{Topics: owner.Topics[:2]}, meta); err == nil {
		t.Fatal("two topics must fail, not panic")
	}
}
