package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

const (
	randomWalkCursorAlice = "0x2100000000000000000000000000000000000021"
	randomWalkCursorBob   = "0x2200000000000000000000000000000000000022"
)

func TestRandomWalkTokenCursorScopesFilterAndSort(t *testing.T) {
	t.Parallel()
	from := int64(100)
	until := int64(200)
	filters := []randomWalkTokenFilterScope{
		{},
		{Named: true},
		{Name: "wanderer"},
		{MintedFrom: &from},
		{MintedFrom: &from, MintedUntil: &until},
	}
	sorts := []RandomWalkTokenSort{TokenId, MostTraded}
	for _, filter := range filters {
		for _, sort := range sorts {
			want := randomWalkTokenCursor{
				Version:  randomWalkCursorVersion,
				Resource: randomWalkResourceTokens,
				Filter:   filter,
				Sort:     sort,
				TokenID:  42,
			}
			if sort == MostTraded {
				want.TradeCount = 7
			}
			encoded, err := encodeRandomWalkTokenCursor(want)
			if err != nil {
				t.Fatalf("encode %+v: %v", want, err)
			}
			got, err := decodeRandomWalkTokenCursor(encoded, filter, sort)
			if err != nil || !got.Filter.equal(want.Filter) ||
				got.Sort != want.Sort || got.TokenID != want.TokenID ||
				got.TradeCount != want.TradeCount {
				t.Fatalf("round trip = %+v, %v; want %+v", got, err, want)
			}
			otherSort := TokenId
			if sort == TokenId {
				otherSort = MostTraded
			}
			if _, err := decodeRandomWalkTokenCursor(encoded, filter, otherSort); err == nil {
				t.Errorf("sort %q cursor decoded under %q", sort, otherSort)
			}
			if _, err := decodeRandomWalkTokenCursor(
				encoded, randomWalkTokenFilterScope{Named: !filter.Named}, sort); err == nil {
				t.Errorf("filter-scoped cursor decoded under another filter")
			}
		}
	}
	// A window filter differing only in bounds is another scope.
	windowCursor, err := encodeRandomWalkTokenCursor(randomWalkTokenCursor{
		Version:  randomWalkCursorVersion,
		Resource: randomWalkResourceTokens,
		Filter:   randomWalkTokenFilterScope{MintedFrom: &from},
		Sort:     TokenId,
		TokenID:  1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeRandomWalkTokenCursor(
		windowCursor, randomWalkTokenFilterScope{MintedFrom: &until}, TokenId); err == nil {
		t.Error("window cursor decoded under another window")
	}
}

func TestRandomWalkTokenCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	payload := func(body string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(body))
	}
	for _, value := range []string{
		"",
		"%%%",
		payload(`{"v":2,"k":"rwalkTokens","f":{},"s":"tokenId","t":1}`),
		payload(`{"v":1,"k":"rwalkUserTokens","f":{},"s":"tokenId","t":1}`),
		payload(`{"v":1,"k":"rwalkTokens","f":{},"s":"bogus","t":1}`),
		payload(`{"v":1,"k":"rwalkTokens","f":{},"s":"tokenId","t":-1}`),
		payload(`{"v":1,"k":"rwalkTokens","f":{},"s":"tokenId","c":3,"t":1}`),
		payload(`{"v":1,"k":"rwalkTokens","f":{"n":true,"q":"x"},"s":"tokenId","t":1}`),
		payload(`{"v":1,"k":"rwalkTokens","f":{},"s":"tokenId","t":1,"x":1}`),
		payload(`{"v":1,"k":"rwalkTokens","f":{},"s":"tokenId","t":1}{}`),
		payload(`{"v":1,"f":{},"t":1}`),
	} {
		if _, err := decodeRandomWalkTokenCursor(
			value, randomWalkTokenFilterScope{}, TokenId,
		); !errors.Is(err, errInvalidRandomWalkCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
	// The CosmicGame token directory cursor shares no resource marker and
	// must be rejected structurally.
	cgCursor, err := encodeGlobalTokenCursor(globalTokenCursor{
		Version: globalTokenCursorVersion,
		TokenID: 4,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeRandomWalkTokenCursor(
		cgCursor, randomWalkTokenFilterScope{}, TokenId,
	); !errors.Is(err, errInvalidRandomWalkCursor) {
		t.Errorf("CosmicGame token cursor accepted: %v", err)
	}
	for _, cursor := range []randomWalkTokenCursor{
		{Version: 2, Resource: randomWalkResourceTokens, Sort: TokenId, TokenID: 1},
		{Version: 1, Resource: randomWalkResourceTrades, Sort: TokenId, TokenID: 1},
		{Version: 1, Resource: randomWalkResourceTokens, Sort: "bogus", TokenID: 1},
		{Version: 1, Resource: randomWalkResourceTokens, Sort: TokenId, TokenID: 1, TradeCount: 5},
		{Version: 1, Resource: randomWalkResourceTokens, Sort: MostTraded, TokenID: -1},
	} {
		if _, err := encodeRandomWalkTokenCursor(cursor); !errors.Is(err, errInvalidRandomWalkCursor) {
			t.Errorf("encode %+v error = %v", cursor, err)
		}
	}
}

func TestRandomWalkTokenEventCursorScopesTokenAndResource(t *testing.T) {
	t.Parallel()
	for _, resource := range []randomWalkResource{
		randomWalkResourceNameHistory,
		randomWalkResourceTokenEvents,
	} {
		want := randomWalkTokenEventCursor{
			Version:    randomWalkCursorVersion,
			Resource:   resource,
			TokenID:    10,
			EventLogID: 5088,
		}
		encoded, err := encodeRandomWalkTokenEventCursor(want)
		if err != nil {
			t.Fatalf("encode %q: %v", resource, err)
		}
		got, err := decodeRandomWalkTokenEventCursor(encoded, 10, resource)
		if err != nil || got != want {
			t.Fatalf("round trip %q = %+v, %v", resource, got, err)
		}
		if _, err := decodeRandomWalkTokenEventCursor(encoded, 11, resource); err == nil {
			t.Errorf("%q cursor decoded under another token", resource)
		}
		other := randomWalkResourceNameHistory
		if resource == other {
			other = randomWalkResourceTokenEvents
		}
		if _, err := decodeRandomWalkTokenEventCursor(encoded, 10, other); err == nil {
			t.Errorf("%q cursor decoded as %q", resource, other)
		}
	}
	// The CosmicGame per-token cursor shape matches but its resource
	// values are disjoint by construction.
	cgCursor, err := encodeTokenEventCursor(tokenEventCursor{
		Version:    tokenEventCursorVersion,
		Resource:   tokenEventResourceNameHistory,
		TokenID:    10,
		EventLogID: 5088,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeRandomWalkTokenEventCursor(
		cgCursor, 10, randomWalkResourceNameHistory,
	); !errors.Is(err, errInvalidRandomWalkCursor) {
		t.Errorf("CosmicGame token event cursor accepted: %v", err)
	}
	if _, err := encodeRandomWalkTokenEventCursor(randomWalkTokenEventCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceTrades,
		TokenID:    10,
		EventLogID: 1,
	}); !errors.Is(err, errInvalidRandomWalkCursor) {
		t.Errorf("encode with ledger resource error = %v", err)
	}
}

func TestRandomWalkOfferBookCursorScopesSort(t *testing.T) {
	t.Parallel()
	priced := randomWalkOfferBookCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceOfferBook,
		Sort:       PriceAsc,
		PriceWei:   "2000000000000000000",
		EventLogID: 5091,
	}
	encoded, err := encodeRandomWalkOfferBookCursor(priced)
	if err != nil {
		t.Fatal(err)
	}
	got, err := decodeRandomWalkOfferBookCursor(encoded, PriceAsc)
	if err != nil || got != priced {
		t.Fatalf("round trip = %+v, %v", got, err)
	}
	for _, other := range []RandomWalkOfferSort{Newest, Oldest, PriceDesc} {
		if _, err := decodeRandomWalkOfferBookCursor(encoded, other); err == nil {
			t.Errorf("priceAsc cursor decoded under %q", other)
		}
	}

	eventOnly := randomWalkOfferBookCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceOfferBook,
		Sort:       Newest,
		EventLogID: 5092,
	}
	encoded, err = encodeRandomWalkOfferBookCursor(eventOnly)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeRandomWalkOfferBookCursor(encoded, Newest); err != nil {
		t.Fatal(err)
	}

	payload := func(body string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(body))
	}
	for _, value := range []string{
		payload(`{"v":1,"k":"rwalkOffers","s":"newest","p":"5","e":1}`),
		payload(`{"v":1,"k":"rwalkOffers","s":"priceAsc","e":1}`),
		payload(`{"v":1,"k":"rwalkOffers","s":"priceAsc","p":"1.5","e":1}`),
		payload(`{"v":1,"k":"rwalkOffers","s":"priceAsc","p":"-1","e":1}`),
		payload(`{"v":1,"k":"rwalkOffers","s":"newest","e":0}`),
		payload(`{"v":1,"k":"rwalkTrades","s":"newest","e":1}`),
	} {
		sort := Newest
		if _, err := decodeRandomWalkOfferBookCursor(value, sort); err == nil {
			if _, err := decodeRandomWalkOfferBookCursor(value, PriceAsc); err == nil {
				t.Errorf("cursor %q accepted under both sorts", value)
			}
		}
	}
	if _, err := encodeRandomWalkOfferBookCursor(randomWalkOfferBookCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceOfferBook,
		Sort:       PriceAsc,
		EventLogID: 1,
	}); !errors.Is(err, errInvalidRandomWalkCursor) {
		t.Errorf("encode price sort without price error = %v", err)
	}
}

func TestRandomWalkLedgerCursorScopesResourceAndWallet(t *testing.T) {
	t.Parallel()
	for _, resource := range []randomWalkResource{
		randomWalkResourceOfferHistory,
		randomWalkResourceTrades,
		randomWalkResourceWithdrawals,
	} {
		want := randomWalkLedgerCursor{
			Version:    randomWalkCursorVersion,
			Resource:   resource,
			EventLogID: 5090,
		}
		encoded, err := encodeRandomWalkLedgerCursor(want)
		if err != nil {
			t.Fatalf("encode %q: %v", resource, err)
		}
		got, err := decodeRandomWalkLedgerCursor(encoded, resource, "")
		if err != nil || got != want {
			t.Fatalf("round trip %q = %+v, %v", resource, got, err)
		}
		for _, other := range []randomWalkResource{
			randomWalkResourceOfferHistory,
			randomWalkResourceTrades,
			randomWalkResourceWithdrawals,
		} {
			if other == resource {
				continue
			}
			if _, err := decodeRandomWalkLedgerCursor(encoded, other, ""); err == nil {
				t.Errorf("%q cursor decoded as %q", resource, other)
			}
		}
	}

	wallet := randomWalkLedgerCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceUserOffers,
		Address:    randomWalkCursorAlice,
		EventLogID: 5093,
	}
	encoded, err := encodeRandomWalkLedgerCursor(wallet)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeRandomWalkLedgerCursor(
		encoded, randomWalkResourceUserOffers, randomWalkCursorAlice); err != nil {
		t.Fatal(err)
	}
	if _, err := decodeRandomWalkLedgerCursor(
		encoded, randomWalkResourceUserOffers, randomWalkCursorBob); err == nil {
		t.Error("wallet-scoped cursor decoded under another wallet")
	}
	if _, err := decodeRandomWalkLedgerCursor(
		encoded, randomWalkResourceOfferHistory, ""); err == nil {
		t.Error("wallet-scoped cursor decoded as the global ledger")
	}
	if _, err := encodeRandomWalkLedgerCursor(randomWalkLedgerCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceTrades,
		Address:    randomWalkCursorAlice,
		EventLogID: 1,
	}); !errors.Is(err, errInvalidRandomWalkCursor) {
		t.Errorf("encode global ledger with wallet error = %v", err)
	}
	if _, err := encodeRandomWalkLedgerCursor(randomWalkLedgerCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceUserOffers,
		EventLogID: 1,
	}); !errors.Is(err, errInvalidRandomWalkCursor) {
		t.Errorf("encode wallet ledger without wallet error = %v", err)
	}
}

func TestRandomWalkUserTokenCursorScopesWallet(t *testing.T) {
	t.Parallel()
	want := randomWalkUserTokenCursor{
		Version:  randomWalkCursorVersion,
		Resource: randomWalkResourceUserTokens,
		Address:  randomWalkCursorAlice,
		TokenID:  11,
	}
	encoded, err := encodeRandomWalkUserTokenCursor(want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := decodeRandomWalkUserTokenCursor(encoded, randomWalkCursorAlice)
	if err != nil || got != want {
		t.Fatalf("round trip = %+v, %v", got, err)
	}
	if _, err := decodeRandomWalkUserTokenCursor(encoded, randomWalkCursorBob); err == nil {
		t.Error("wallet-scoped cursor decoded under another wallet")
	}
	// The CosmicGame owned-token cursor has the same field shapes but no
	// resource marker; it must fail structurally.
	cgCursor, err := encodeUserOwnedTokenCursor(userOwnedTokenCursor{
		Version: userOwnedTokenCursorVersion,
		Address: randomWalkCursorAlice,
		TokenID: 11,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeRandomWalkUserTokenCursor(
		cgCursor, randomWalkCursorAlice,
	); !errors.Is(err, errInvalidRandomWalkCursor) {
		t.Errorf("CosmicGame owned-token cursor accepted: %v", err)
	}
	// And symmetrically: the RandomWalk cursor cannot leak into the
	// CosmicGame owned-token collection.
	if _, err := decodeUserOwnedTokenCursor(
		encoded, randomWalkCursorAlice,
	); !errors.Is(err, errInvalidUserOwnedTokenCursor) {
		t.Errorf("RandomWalk owned-token cursor accepted by CosmicGame: %v", err)
	}
}

func TestRandomWalkTokenFilterScopeValid(t *testing.T) {
	t.Parallel()
	negative := int64(-1)
	zero := int64(0)
	if (randomWalkTokenFilterScope{Name: strings.Repeat("a", maxTokenNameSearchLength+1)}).valid() {
		t.Fatal("oversized name accepted")
	}
	if (randomWalkTokenFilterScope{MintedFrom: &negative}).valid() {
		t.Fatal("negative mintedFrom accepted")
	}
	if (randomWalkTokenFilterScope{MintedUntil: &zero}).valid() {
		t.Fatal("zero mintedUntil accepted")
	}
	if !(randomWalkTokenFilterScope{Name: "wanderer"}).valid() {
		t.Fatal("plain search rejected")
	}
}

func TestValidWeiAmountString(t *testing.T) {
	t.Parallel()
	for value, want := range map[string]bool{
		"0": true, "10": true, "": false, "1x": false, "-1": false, "1.5": false,
	} {
		if got := validWeiAmountString(value); got != want {
			t.Errorf("validWeiAmountString(%q) = %v, want %v", value, got, want)
		}
	}
}

func TestRandomWalkLedgerResourceRejectsNonLedgerValues(t *testing.T) {
	t.Parallel()
	for _, resource := range []randomWalkResource{
		randomWalkResourceTokens,
		randomWalkResourceNameHistory,
		randomWalkResourceTokenEvents,
		randomWalkResourceOfferBook,
		randomWalkResourceUserTokens,
		randomWalkResource("bogus"),
	} {
		if validRandomWalkLedgerResource(resource) {
			t.Errorf("%q accepted as a ledger resource", resource)
		}
	}
}

func TestEncodeRandomWalkUserTokenCursorRejectsInvalidFields(t *testing.T) {
	t.Parallel()
	if _, err := encodeRandomWalkUserTokenCursor(randomWalkUserTokenCursor{
		Version:  randomWalkCursorVersion,
		Resource: randomWalkResourceUserTokens,
		Address:  "not-an-address",
		TokenID:  1,
	}); !errors.Is(err, errInvalidRandomWalkCursor) {
		t.Fatalf("encode invalid user-token cursor error = %v", err)
	}
	if _, err := encodeRandomWalkUserTokenCursor(randomWalkUserTokenCursor{
		Version:  randomWalkCursorVersion,
		Resource: randomWalkResourceTokens,
		Address:  randomWalkCursorAlice,
		TokenID:  1,
	}); !errors.Is(err, errInvalidRandomWalkCursor) {
		t.Fatalf("encode wrong-resource user-token cursor error = %v", err)
	}
}

func FuzzDecodeRandomWalkTokenCursor(f *testing.F) {
	from := int64(100)
	valid, _ := encodeRandomWalkTokenCursor(randomWalkTokenCursor{
		Version:  randomWalkCursorVersion,
		Resource: randomWalkResourceTokens,
		Filter:   randomWalkTokenFilterScope{MintedFrom: &from},
		Sort:     MostTraded,
		TokenID:  4,
	})
	f.Add(valid, true, "wand", int64(100), int64(200), "mostTraded")
	f.Add("", false, "", int64(-1), int64(-1), "tokenId")
	f.Fuzz(func(t *testing.T, value string, named bool, name string, mintedFrom, mintedUntil int64, sort string) {
		filter := randomWalkTokenFilterScope{Named: named, Name: name}
		if mintedFrom >= 0 {
			filter.MintedFrom = &mintedFrom
		}
		if mintedUntil >= 0 {
			filter.MintedUntil = &mintedUntil
		}
		expectedSort := RandomWalkTokenSort(sort)
		cursor, err := decodeRandomWalkTokenCursor(value, filter, expectedSort)
		if err == nil && !validRandomWalkTokenCursor(cursor, filter, expectedSort) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}

func FuzzDecodeRandomWalkTokenEventCursor(f *testing.F) {
	valid, _ := encodeRandomWalkTokenEventCursor(randomWalkTokenEventCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceTokenEvents,
		TokenID:    10,
		EventLogID: 5088,
	})
	f.Add(valid, int64(10), "rwalkTokenEvents")
	f.Add("", int64(-1), "unknown")
	f.Fuzz(func(t *testing.T, value string, tokenID int64, resource string) {
		expected := randomWalkResource(resource)
		cursor, err := decodeRandomWalkTokenEventCursor(value, tokenID, expected)
		if err == nil && !validRandomWalkTokenEventCursor(cursor, tokenID, expected) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}

func FuzzDecodeRandomWalkOfferBookCursor(f *testing.F) {
	valid, _ := encodeRandomWalkOfferBookCursor(randomWalkOfferBookCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceOfferBook,
		Sort:       PriceDesc,
		PriceWei:   "10",
		EventLogID: 3,
	})
	f.Add(valid, "priceDesc")
	f.Add("", "newest")
	f.Fuzz(func(t *testing.T, value, sort string) {
		expected := RandomWalkOfferSort(sort)
		cursor, err := decodeRandomWalkOfferBookCursor(value, expected)
		if err == nil && !validRandomWalkOfferBookCursor(cursor, expected) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}

func FuzzDecodeRandomWalkLedgerCursor(f *testing.F) {
	valid, _ := encodeRandomWalkLedgerCursor(randomWalkLedgerCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceUserOffers,
		Address:    randomWalkCursorAlice,
		EventLogID: 9,
	})
	f.Add(valid, "rwalkUserOffers", randomWalkCursorAlice)
	f.Add("", "rwalkTrades", "")
	f.Fuzz(func(t *testing.T, value, resource, address string) {
		expected := randomWalkResource(resource)
		cursor, err := decodeRandomWalkLedgerCursor(value, expected, address)
		if err == nil && !validRandomWalkLedgerCursor(cursor, expected, address) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}

func FuzzDecodeRandomWalkUserTokenCursor(f *testing.F) {
	valid, _ := encodeRandomWalkUserTokenCursor(randomWalkUserTokenCursor{
		Version:  randomWalkCursorVersion,
		Resource: randomWalkResourceUserTokens,
		Address:  randomWalkCursorAlice,
		TokenID:  2,
	})
	f.Add(valid, randomWalkCursorAlice)
	f.Add("", "not-an-address")
	f.Fuzz(func(t *testing.T, value, address string) {
		cursor, err := decodeRandomWalkUserTokenCursor(value, address)
		if err == nil && !validRandomWalkUserTokenCursor(cursor, address) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}
