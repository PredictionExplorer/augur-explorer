package randomwalk

// No-database tests for the v2 read surface: cursor/filter/sort validation,
// page truncation, ILIKE escaping and the compile-time shape of the
// six-branch token event merge.

import (
	"strings"
	"testing"
)

func TestTruncatePage(t *testing.T) {
	t.Parallel()
	records, hasMore := truncatePage([]int{1, 2, 3}, 2)
	if len(records) != 2 || !hasMore {
		t.Fatalf("truncate = (%v,%v)", records, hasMore)
	}
	records, hasMore = truncatePage([]int{1, 2}, 2)
	if len(records) != 2 || hasMore {
		t.Fatalf("exact page = (%v,%v)", records, hasMore)
	}
	records, hasMore = truncatePage([]int{}, 2)
	if len(records) != 0 || hasMore {
		t.Fatalf("empty page = (%v,%v)", records, hasMore)
	}
}

func TestTokenFilterValid(t *testing.T) {
	t.Parallel()
	negative := int64(-1)
	zero := int64(0)
	one := int64(1)
	cases := map[string]struct {
		filter TokenFilter
		want   bool
	}{
		"empty":              {TokenFilter{}, true},
		"named":              {TokenFilter{NamedOnly: true}, true},
		"search":             {TokenFilter{NameContains: "wanderer"}, true},
		"contradictory":      {TokenFilter{NamedOnly: true, NameContains: "x"}, false},
		"window":             {TokenFilter{MintedFrom: &zero, MintedUntil: &one}, true},
		"negative from":      {TokenFilter{MintedFrom: &negative}, false},
		"zero until":         {TokenFilter{MintedUntil: &zero}, false},
		"negative until":     {TokenFilter{MintedUntil: &negative}, false},
		"from without until": {TokenFilter{MintedFrom: &one}, true},
	}
	for name, tc := range cases {
		if got := tc.filter.valid(); got != tc.want {
			t.Errorf("%s: valid() = %v, want %v", name, got, tc.want)
		}
	}
}

func TestTokenSortValid(t *testing.T) {
	t.Parallel()
	if !TokenSortByID.valid() || !TokenSortByTrades.valid() {
		t.Fatal("canonical sorts rejected")
	}
	if TokenSort("").valid() || TokenSort("priceAsc").valid() {
		t.Fatal("unknown sort accepted")
	}
}

func TestCursorValidation(t *testing.T) {
	t.Parallel()
	var nilToken *TokenPageCursor
	if !nilToken.valid() {
		t.Fatal("nil token cursor rejected")
	}
	if (&TokenPageCursor{TokenID: -1}).valid() ||
		(&TokenPageCursor{TradeCount: -1}).valid() {
		t.Fatal("negative token cursor accepted")
	}
	var nilEvent *EventPageCursor
	if !nilEvent.valid() {
		t.Fatal("nil event cursor rejected")
	}
	if (&EventPageCursor{EventLogID: 0}).valid() {
		t.Fatal("zero event cursor accepted")
	}
	if !(&EventPageCursor{EventLogID: 1}).valid() {
		t.Fatal("valid event cursor rejected")
	}
}

func TestOfferPageCursorValidFor(t *testing.T) {
	t.Parallel()
	var nilCursor *OfferPageCursor
	for _, sort := range []OfferSort{
		OfferSortNewest, OfferSortOldest, OfferSortPriceAsc, OfferSortPriceDesc,
	} {
		if !nilCursor.validFor(sort) {
			t.Fatalf("nil cursor rejected for %s", sort)
		}
	}
	eventOnly := &OfferPageCursor{EventLogID: 7}
	if !eventOnly.validFor(OfferSortNewest) || !eventOnly.validFor(OfferSortOldest) {
		t.Fatal("event-only cursor rejected for event sorts")
	}
	if eventOnly.validFor(OfferSortPriceAsc) {
		t.Fatal("price sort accepted a cursor without a price")
	}
	priced := &OfferPageCursor{EventLogID: 7, PriceWei: "1000"}
	if !priced.validFor(OfferSortPriceAsc) || !priced.validFor(OfferSortPriceDesc) {
		t.Fatal("priced cursor rejected for price sorts")
	}
	if priced.validFor(OfferSortNewest) {
		t.Fatal("event sort accepted a priced cursor")
	}
	if (&OfferPageCursor{EventLogID: 0, PriceWei: "1"}).validFor(OfferSortPriceAsc) {
		t.Fatal("zero event-log cursor accepted")
	}
	if (&OfferPageCursor{EventLogID: 1, PriceWei: "1.5"}).validFor(OfferSortPriceAsc) {
		t.Fatal("fractional price cursor accepted")
	}
	if (&OfferPageCursor{EventLogID: 1, PriceWei: "-1"}).validFor(OfferSortPriceAsc) {
		t.Fatal("negative price cursor accepted")
	}
	if (&OfferPageCursor{EventLogID: 1}).validFor(OfferSort("bogus")) {
		t.Fatal("unknown sort accepted")
	}
}

func TestValidWeiString(t *testing.T) {
	t.Parallel()
	for value, want := range map[string]bool{
		"0":                           true,
		"1000":                        true,
		"":                            false,
		"-1":                          false,
		"1.5":                         false,
		"1e18":                        false,
		"0x10":                        false,
		" 1":                          false, //nolint:gocritic // deliberate: leading whitespace must be rejected
		"9" + strings.Repeat("0", 77): true,
	} {
		if got := validWeiString(value); got != want {
			t.Errorf("validWeiString(%q) = %v, want %v", value, got, want)
		}
	}
}

func TestEscapeLikePattern(t *testing.T) {
	t.Parallel()
	for input, want := range map[string]string{
		"plain":   "plain",
		"100%":    `100\%`,
		"a_b":     `a\_b`,
		`back\`:   `back\\`,
		`%_\mix`:  `\%\_\\mix`,
		"":        "",
		"Wander%": `Wander\%`,
	} {
		if got := escapeLikePattern(input); got != want {
			t.Errorf("escapeLikePattern(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestTokenEventsPageSQLShape(t *testing.T) {
	t.Parallel()
	for _, withCursor := range []bool{false, true} {
		query := tokenEventsPageSQL(withCursor)
		for _, kind := range []TokenEventKind{
			TokenEventMint, TokenEventTransfer, TokenEventNameChange,
			TokenEventListed, TokenEventOfferCanceled, TokenEventPurchase,
		} {
			if !strings.Contains(query, "'"+string(kind)+"'::TEXT") {
				t.Fatalf("withCursor=%v: branch %q missing", withCursor, kind)
			}
		}
		if got := strings.Count(query, "UNION ALL"); got != 5 {
			t.Fatalf("withCursor=%v: %d UNION ALL joins, want 5", withCursor, got)
		}
		limitArg, cursorFilters := "$4", 0
		if withCursor {
			limitArg = "$5"
			cursorFilters = 6
		}
		if got := strings.Count(query, "e.evtlog_id < $4"); got != cursorFilters {
			t.Fatalf("withCursor=%v: %d cursor filters, want %d", withCursor, got, cursorFilters)
		}
		if got := strings.Count(query, "LIMIT "+limitArg); got != 7 {
			t.Fatalf("withCursor=%v: %d limits on %s, want 7", withCursor, got, limitArg)
		}
		// The mint-mirroring transfers are excluded structurally.
		if !strings.Contains(query, "e.otype <> 1") {
			t.Fatalf("withCursor=%v: transfer branch keeps mint transfers", withCursor)
		}
	}
}
