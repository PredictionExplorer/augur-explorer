//go:build integration

package randomwalk

// API v2 read-surface suite: every keyset page walks to exhaustion at
// several page sizes and is checked against the fixture events (and, where
// one exists, the legacy query it replaces); exact wei strings, filters,
// sorts and window series are pinned; cancellation and closed-pool paths
// degrade into errors instead of partial pages.

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Fixture wei constants (internal/testfixtures/seed/03_randomwalk.sql).
const (
	weiMint10   = "50000000000000000"
	weiMint11   = "55000000000000000"
	weiMint12   = "60000000000000000"
	weiMint13   = "65000000000000000"
	weiOffer1   = "1000000000000000000"
	weiOffer2   = "2000000000000000000"
	weiOffer3   = "2500000000000000000"
	weiOffer4   = "3000000000000000000"
	weiWithdraw = "30000000000000000"

	addrAlice = "0x2100000000000000000000000000000000000021"
	addrBob   = "0x2200000000000000000000000000000000000022"
	addrCarol = "0x2300000000000000000000000000000000000023"
	addrDave  = "0x2400000000000000000000000000000000000024"
)

func walkTokensPage(
	t *testing.T,
	r *Repo,
	filter TokenFilter,
	sort TokenSort,
	pageSize int,
) []TokenRecord {
	t.Helper()
	ctx := context.Background()
	var all []TokenRecord
	var after *TokenPageCursor
	for {
		page, hasMore, err := r.TokensPage(ctx, filter, sort, after, pageSize)
		if err != nil {
			t.Fatalf("TokensPage(%+v): %v", after, err)
		}
		if len(page) > pageSize {
			t.Fatalf("page of %d exceeds size %d", len(page), pageSize)
		}
		all = append(all, page...)
		if !hasMore {
			return all
		}
		if len(page) == 0 {
			t.Fatal("hasMore without rows")
		}
		last := page[len(page)-1]
		after = &TokenPageCursor{TokenID: last.TokenID, TradeCount: last.TradeCount}
	}
}

func tokenIDs(records []TokenRecord) []int64 {
	ids := make([]int64, 0, len(records))
	for _, record := range records {
		ids = append(ids, record.TokenID)
	}
	return ids
}

func TestTokensPageWalksMintDirectory(t *testing.T) {
	r := repo(t)
	for _, pageSize := range []int{1, 2, 50} {
		records := walkTokensPage(t, r, TokenFilter{}, TokenSortByID, pageSize)
		if !reflect.DeepEqual(tokenIDs(records), []int64{10, 11, 12, 13}) {
			t.Fatalf("pageSize=%d ids=%v", pageSize, tokenIDs(records))
		}
	}
	records := walkTokensPage(t, r, TokenFilter{}, TokenSortByID, 50)
	first := records[0]
	if first.MintTx.EvtLogID != 5080 || first.MintTx.BlockNum != 130 ||
		first.MintPriceWei != weiMint10 || first.SeedNum != "16" ||
		first.Seed != "aa00000000000000000000000000000000000000000000000000000000000010" ||
		first.MinterAddr != addrCarol || first.CurOwnerAddr != addrDave ||
		first.TokenName != "Wanderer" || first.TradeCount != 1 ||
		first.LastPriceWei != weiOffer1 || first.TradingVolumeWei != weiOffer1 {
		t.Fatalf("token 10 = %+v", first)
	}
	unsold := records[1]
	if unsold.TokenID != 11 || unsold.TradeCount != 0 ||
		unsold.LastPriceWei != weiMint11 || unsold.TradingVolumeWei != "0" ||
		unsold.TokenName != "" || unsold.CurOwnerAddr != addrDave {
		t.Fatalf("token 11 = %+v", unsold)
	}
}

func TestTokensPageFiltersAndSorts(t *testing.T) {
	r := repo(t)
	cases := map[string]struct {
		filter TokenFilter
		sort   TokenSort
		want   []int64
	}{
		"named only":            {TokenFilter{NamedOnly: true}, TokenSortByID, []int64{10}},
		"name case-insensitive": {TokenFilter{NameContains: "wand"}, TokenSortByID, []int64{10}},
		"name literal wildcard": {TokenFilter{NameContains: "%"}, TokenSortByID, nil},
		"minted window": {
			TokenFilter{MintedFrom: int64Pointer(1767228700), MintedUntil: int64Pointer(1767228900)},
			TokenSortByID,
			[]int64{11, 12},
		},
		"most traded": {TokenFilter{}, TokenSortByTrades, []int64{10, 11, 12, 13}},
		"most traded windowed": {
			TokenFilter{MintedFrom: int64Pointer(1767228700)},
			TokenSortByTrades,
			[]int64{11, 12, 13},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			for _, pageSize := range []int{1, 50} {
				got := tokenIDs(walkTokensPage(t, r, tc.filter, tc.sort, pageSize))
				want := tc.want
				if want == nil {
					want = []int64{}
				}
				if len(got) == 0 && len(want) == 0 {
					continue
				}
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("pageSize=%d ids=%v want=%v", pageSize, got, want)
				}
			}
		})
	}
}

func TestTokensPageRejectsInvalidInput(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	if _, _, err := r.TokensPage(ctx, TokenFilter{}, TokenSortByID, nil, 0); err == nil {
		t.Fatal("zero limit accepted")
	}
	if _, _, err := r.TokensPage(ctx, TokenFilter{}, TokenSort("bogus"), nil, 1); err == nil {
		t.Fatal("invalid sort accepted")
	}
	if _, _, err := r.TokensPage(ctx,
		TokenFilter{NamedOnly: true, NameContains: "x"}, TokenSortByID, nil, 1); err == nil {
		t.Fatal("contradictory filter accepted")
	}
	if _, _, err := r.TokensPage(ctx,
		TokenFilter{}, TokenSortByID, &TokenPageCursor{TokenID: -1}, 1); err == nil {
		t.Fatal("negative cursor accepted")
	}
}

func TestTokenDetailV2(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	named, err := r.TokenDetailV2(ctx, 10)
	if err != nil {
		t.Fatal(err)
	}
	if named.TokenName != "Wanderer" || named.NameChangeTs != 1767229000 ||
		named.NameChangeText == "" || named.LastPriceWei != weiOffer1 {
		t.Fatalf("token 10 detail = %+v", named)
	}

	unnamed, err := r.TokenDetailV2(ctx, 13)
	if err != nil {
		t.Fatal(err)
	}
	if unnamed.TokenName != "" || unnamed.NameChangeTs != 0 || unnamed.NameChangeText != "" {
		t.Fatalf("token 13 detail = %+v", unnamed)
	}

	if _, err := r.TokenDetailV2(ctx, 999); !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("missing token error = %v", err)
	}
	if _, err := r.TokenDetailV2(ctx, -1); err == nil {
		t.Fatal("negative token id accepted")
	}
}

func TestCollectionTokenExists(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	exists, err := r.CollectionTokenExists(ctx, 10)
	if err != nil || !exists {
		t.Fatalf("token 10 exists=(%v,%v)", exists, err)
	}
	exists, err = r.CollectionTokenExists(ctx, 999)
	if err != nil || exists {
		t.Fatalf("token 999 exists=(%v,%v)", exists, err)
	}
	if _, err := r.CollectionTokenExists(ctx, -5); err == nil {
		t.Fatal("negative token id accepted")
	}
}

func TestTokenNameChangesPageV2(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	records, hasMore, err := r.TokenNameChangesPageV2(ctx, 10, nil, 50)
	if err != nil || hasMore {
		t.Fatalf("name changes err=%v hasMore=%v", err, hasMore)
	}
	if len(records) != 1 || records[0].Tx.EvtLogID != 5087 ||
		records[0].NewName != "Wanderer" || records[0].Owner != addrCarol ||
		records[0].TokenID != 10 {
		t.Fatalf("name changes = %+v", records)
	}
	// The cursor positioned at the only rename exhausts the ledger.
	records, hasMore, err = r.TokenNameChangesPageV2(ctx, 10, &EventPageCursor{EventLogID: 5087}, 50)
	if err != nil || hasMore || len(records) != 0 {
		t.Fatalf("exhausted page = (%v,%v,%v)", records, hasMore, err)
	}
	records, hasMore, err = r.TokenNameChangesPageV2(ctx, 11, nil, 50)
	if err != nil || hasMore || len(records) != 0 {
		t.Fatalf("never-renamed page = (%v,%v,%v)", records, hasMore, err)
	}
}

// fixtureToken10Events is the expected newest-first provenance ledger of
// token 10: the mint transfer (5079) is represented by the mint event and
// the purchase transfer (5090) by the purchase event.
var fixtureToken10Events = []struct {
	evtlog int64
	kind   TokenEventKind
}{
	{5094, TokenEventOfferCanceled},
	{5093, TokenEventListed},
	{5089, TokenEventPurchase},
	{5088, TokenEventListed},
	{5087, TokenEventNameChange},
	{5080, TokenEventMint},
}

func walkTokenEvents(t *testing.T, r *Repo, tokenID int64, pageSize int) []TokenEventRecord {
	t.Helper()
	ctx := context.Background()
	var all []TokenEventRecord
	var after *EventPageCursor
	for {
		page, hasMore, err := r.TokenEventsPage(ctx, tokenID, after, pageSize)
		if err != nil {
			t.Fatalf("TokenEventsPage(%+v): %v", after, err)
		}
		all = append(all, page...)
		if !hasMore {
			return all
		}
		if len(page) == 0 {
			t.Fatal("hasMore without rows")
		}
		after = &EventPageCursor{EventLogID: page[len(page)-1].Tx.EvtLogID}
	}
}

func TestTokenEventsPageInterleavesSixSources(t *testing.T) {
	r := repo(t)
	for _, pageSize := range []int{1, 2, 50} {
		records := walkTokenEvents(t, r, 10, pageSize)
		if len(records) != len(fixtureToken10Events) {
			t.Fatalf("pageSize=%d events=%d want=%d", pageSize, len(records), len(fixtureToken10Events))
		}
		for i, want := range fixtureToken10Events {
			if records[i].Tx.EvtLogID != want.evtlog || records[i].Kind != want.kind {
				t.Fatalf("pageSize=%d event[%d]=(%d,%s) want (%d,%s)",
					pageSize, i, records[i].Tx.EvtLogID, records[i].Kind, want.evtlog, want.kind)
			}
		}
	}

	records := walkTokenEvents(t, r, 10, 50)
	canceled, relisted, purchase := records[0], records[1], records[2]
	firstListing, rename, mint := records[3], records[4], records[5]
	if canceled.OfferID != 4 || !canceled.HasOffer || canceled.OfferType != 1 ||
		canceled.PriceWei != weiOffer4 || canceled.MakerAddr != addrDave {
		t.Fatalf("canceled = %+v", canceled)
	}
	if relisted.OfferID != 4 || relisted.PriceWei != weiOffer4 || relisted.MakerAddr != addrDave {
		t.Fatalf("relisted = %+v", relisted)
	}
	if purchase.OfferID != 1 || purchase.PriceWei != weiOffer1 ||
		purchase.BuyerAddr != addrDave || purchase.SellerAddr != addrCarol {
		t.Fatalf("purchase = %+v", purchase)
	}
	if firstListing.OfferID != 1 || firstListing.PriceWei != weiOffer1 ||
		firstListing.MakerAddr != addrCarol {
		t.Fatalf("first listing = %+v", firstListing)
	}
	if !rename.HasNewName || rename.NewName != "Wanderer" {
		t.Fatalf("rename = %+v", rename)
	}
	if mint.MinterAddr != addrCarol || mint.PriceWei != weiMint10 || mint.SeedNum != "16" {
		t.Fatalf("mint = %+v", mint)
	}

	// A token without marketplace or naming activity is just its mint.
	quiet := walkTokenEvents(t, r, 12, 50)
	if len(quiet) != 1 || quiet[0].Kind != TokenEventMint || quiet[0].Tx.EvtLogID != 5084 {
		t.Fatalf("token 12 events = %+v", quiet)
	}
}

func TestTokenEventsPageIncludesPlainTransfersAndBurns(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// A plain wallet-to-wallet transfer (no marketplace event in the tx) and
	// a burn must surface as transfer rows; the events reuse fixture
	// transactions under fresh log indexes. Cleanup restores fixture
	// ownership (the delete trigger deliberately keeps the newer owner, so
	// the update below puts the fixture owner back).
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		VALUES (8101, 142, 1042, 8, 'ddf252ad', 90, '\x00'), (8102, 129, 1043, 8, 'ddf252ad', 91, '\x00')`); err != nil {
		t.Fatal(err)
	}
	cleanup := func() {
		if _, err := r.pool().Exec(context.Background(),
			"DELETE FROM evt_log WHERE id IN (8101, 8102)"); err != nil {
			t.Errorf("cleaning transfer extension: %v", err)
		}
		if _, err := r.pool().Exec(context.Background(),
			"UPDATE rw_token SET cur_owner_aid=21 WHERE rwalk_aid=8 AND token_id=12"); err != nil {
			t.Errorf("restoring token 12 owner: %v", err)
		}
	}
	t.Cleanup(cleanup)
	if _, err := r.pool().Exec(ctx, `INSERT INTO rw_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, from_aid, to_aid, otype) VALUES
		(8101, 142, 1042, TO_TIMESTAMP(1767229800), 8, 12, 21, 22, 0),
		(8102, 129, 1043, TO_TIMESTAMP(1767229900), 8, 12, 22, 1, 2)`); err != nil {
		t.Fatal(err)
	}

	records := walkTokenEvents(t, r, 12, 50)
	if len(records) != 3 ||
		records[0].Kind != TokenEventTransfer || records[0].Tx.EvtLogID != 8102 ||
		records[1].Kind != TokenEventTransfer || records[1].Tx.EvtLogID != 8101 ||
		records[2].Kind != TokenEventMint {
		t.Fatalf("token 12 extended events = %+v", records)
	}
	if records[0].FromAddr != addrBob || records[0].ToAddr != "0x0000000000000000000000000000000000000000" {
		t.Fatalf("burn transfer = %+v", records[0])
	}
	if records[1].FromAddr != addrAlice || records[1].ToAddr != addrBob {
		t.Fatalf("plain transfer = %+v", records[1])
	}
}

func walkActiveOffers(t *testing.T, r *Repo, sort OfferSort, pageSize int) []OfferRecord {
	t.Helper()
	ctx := context.Background()
	var all []OfferRecord
	var after *OfferPageCursor
	for {
		page, hasMore, err := r.ActiveOffersPage(ctx, sort, after, pageSize)
		if err != nil {
			t.Fatalf("ActiveOffersPage(%+v): %v", after, err)
		}
		all = append(all, page...)
		if !hasMore {
			return all
		}
		if len(page) == 0 {
			t.Fatal("hasMore without rows")
		}
		last := page[len(page)-1]
		after = &OfferPageCursor{EventLogID: last.ListTx.EvtLogID}
		if sort == OfferSortPriceAsc || sort == OfferSortPriceDesc {
			after.PriceWei = last.PriceWei
		}
	}
}

func offerIDs(records []OfferRecord) []int64 {
	ids := make([]int64, 0, len(records))
	for _, record := range records {
		ids = append(ids, record.OfferID)
	}
	return ids
}

func TestActiveOffersPageSortsAndPaginates(t *testing.T) {
	r := repo(t)
	cases := map[OfferSort][]int64{
		OfferSortNewest:    {3, 2},
		OfferSortOldest:    {2, 3},
		OfferSortPriceAsc:  {2, 3},
		OfferSortPriceDesc: {3, 2},
	}
	for sort, want := range cases {
		for _, pageSize := range []int{1, 50} {
			got := offerIDs(walkActiveOffers(t, r, sort, pageSize))
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("sort=%s pageSize=%d ids=%v want=%v", sort, pageSize, got, want)
			}
		}
	}
	records := walkActiveOffers(t, r, OfferSortPriceAsc, 50)
	if records[0].PriceWei != weiOffer2 || records[0].MakerAddr != addrDave ||
		records[0].TokenID != 11 || records[0].OfferType != 1 {
		t.Fatalf("cheapest offer = %+v", records[0])
	}
}

func TestActiveOffersPriceTieBreaksOnEventLog(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Two active offers at the same price prove the (price, evtlog)
	// keyset cannot skip or repeat rows; the events reuse fixture
	// transactions under fresh log indexes. Buy-side offers also surface.
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		VALUES (8103, 138, 1038, 12, '55076e90', 92, '\x00'), (8104, 139, 1039, 12, '55076e90', 93, '\x00')`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := r.pool().Exec(context.Background(),
			"DELETE FROM evt_log WHERE id IN (8103, 8104)"); err != nil {
			t.Errorf("cleaning offer extension: %v", err)
		}
	})
	if _, err := r.pool().Exec(ctx, `INSERT INTO rw_new_offer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, rwalk_aid, offer_id, seller_aid, buyer_aid, otype, token_id, active, price) VALUES
		(8103, 138, 1038, TO_TIMESTAMP(1767230000), 12, 8, 5, 22, 1, 1, 13, TRUE, 2000000000000000000),
		(8104, 139, 1039, TO_TIMESTAMP(1767230100), 12, 8, 6, 1, 21, 0, 12, TRUE, 500000000000000000)`); err != nil {
		t.Fatal(err)
	}

	for _, pageSize := range []int{1, 2, 50} {
		got := offerIDs(walkActiveOffers(t, r, OfferSortPriceAsc, pageSize))
		if !reflect.DeepEqual(got, []int64{6, 2, 5, 3}) {
			t.Fatalf("pageSize=%d ids=%v", pageSize, got)
		}
	}
	records := walkActiveOffers(t, r, OfferSortPriceAsc, 50)
	if records[0].OfferType != 0 || records[0].MakerAddr != addrAlice {
		t.Fatalf("buy offer maker = %+v", records[0])
	}
}

func walkOfferHistory(t *testing.T, r *Repo, userAid *int64, pageSize int) []OfferHistoryRecord {
	t.Helper()
	ctx := context.Background()
	var all []OfferHistoryRecord
	var after *EventPageCursor
	for {
		var page []OfferHistoryRecord
		var hasMore bool
		var err error
		if userAid == nil {
			page, hasMore, err = r.OfferHistoryPage(ctx, after, pageSize)
		} else {
			page, hasMore, err = r.UserOffersPage(ctx, *userAid, after, pageSize)
		}
		if err != nil {
			t.Fatalf("offer history (%+v): %v", after, err)
		}
		all = append(all, page...)
		if !hasMore {
			return all
		}
		if len(page) == 0 {
			t.Fatal("hasMore without rows")
		}
		after = &EventPageCursor{EventLogID: page[len(page)-1].ListTx.EvtLogID}
	}
}

func TestOfferHistoryPageCarriesOutcomes(t *testing.T) {
	r := repo(t)
	for _, pageSize := range []int{1, 2, 50} {
		records := walkOfferHistory(t, r, nil, pageSize)
		if len(records) != 4 {
			t.Fatalf("pageSize=%d offers=%d", pageSize, len(records))
		}
	}
	records := walkOfferHistory(t, r, nil, 50)

	canceled := records[0]
	if canceled.OfferID != 4 || canceled.Active || canceled.Cancellation == nil ||
		canceled.Purchase != nil || canceled.Cancellation.EvtLogID != 5094 {
		t.Fatalf("canceled offer = %+v", canceled)
	}
	active := records[1]
	if active.OfferID != 3 || !active.Active || active.Purchase != nil ||
		active.Cancellation != nil || active.PriceWei != weiOffer3 {
		t.Fatalf("active offer = %+v", active)
	}
	bought := records[3]
	if bought.OfferID != 1 || bought.Active || bought.Purchase == nil ||
		bought.Purchase.Tx.EvtLogID != 5089 ||
		bought.Purchase.BuyerAddr != addrDave || bought.Purchase.SellerAddr != addrCarol {
		t.Fatalf("bought offer = %+v", bought)
	}
	// The trigger tracked carol's position: sold for 1 ETH, minted at 0.05.
	if bought.ProfitWei != "950000000000000000" {
		t.Fatalf("seller profit = %q", bought.ProfitWei)
	}
}

func TestUserOffersPageScopesToWallet(t *testing.T) {
	r := repo(t)
	carol := int64(aidCarol)
	dave := int64(aidDave)
	bob := int64(aidBob)
	alice := int64(aidAlice)

	carolOffers := walkOfferHistory(t, r, &carol, 50)
	if len(carolOffers) != 1 || carolOffers[0].OfferID != 1 {
		t.Fatalf("carol offers = %+v", carolOffers)
	}
	daveOffers := walkOfferHistory(t, r, &dave, 1)
	if len(daveOffers) != 3 {
		t.Fatalf("dave offers = %+v", daveOffers)
	}
	bobOffers := walkOfferHistory(t, r, &bob, 50)
	if len(bobOffers) != 1 || bobOffers[0].OfferID != 3 {
		t.Fatalf("bob offers = %+v", bobOffers)
	}
	// Alice never traded: the unfilled-offer buyer placeholder (aid 1) must
	// not leak other wallets into her ledger.
	aliceOffers := walkOfferHistory(t, r, &alice, 50)
	if len(aliceOffers) != 0 {
		t.Fatalf("alice offers = %+v", aliceOffers)
	}
}

func TestTradesPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	records, hasMore, err := r.TradesPage(ctx, nil, 50)
	if err != nil || hasMore {
		t.Fatalf("trades err=%v hasMore=%v", err, hasMore)
	}
	if len(records) != 1 {
		t.Fatalf("trades = %+v", records)
	}
	trade := records[0]
	if trade.Tx.EvtLogID != 5089 || trade.OfferID != 1 || trade.TokenID != 10 ||
		trade.PriceWei != weiOffer1 || trade.BuyerAddr != addrDave ||
		trade.SellerAddr != addrCarol || trade.ProfitWei != "950000000000000000" {
		t.Fatalf("trade = %+v", trade)
	}
	exhausted, hasMore, err := r.TradesPage(ctx, &EventPageCursor{EventLogID: 5089}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted trades = (%v,%v,%v)", exhausted, hasMore, err)
	}
}

func TestFloorPriceV2(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	record, err := r.FloorPriceV2(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if record.ActiveSellOfferCount != 2 || record.Floor == nil ||
		record.Floor.OfferID != 2 || record.Floor.TokenID != 11 ||
		record.Floor.PriceWei != weiOffer2 || record.Floor.ListedAtTs != 1767229300 {
		t.Fatalf("floor = %+v", record)
	}

	// An empty book is a valid zero result: hide the two active sell
	// offers, then restore them.
	if _, err := r.pool().Exec(ctx,
		"UPDATE rw_new_offer SET active=FALSE WHERE offer_id IN (2, 3) AND contract_aid=12"); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := r.pool().Exec(context.Background(),
			"UPDATE rw_new_offer SET active=TRUE WHERE offer_id IN (2, 3) AND contract_aid=12"); err != nil {
			t.Errorf("restoring active offers: %v", err)
		}
	})
	empty, err := r.FloorPriceV2(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if empty.ActiveSellOfferCount != 0 || empty.Floor != nil {
		t.Fatalf("empty floor = %+v", empty)
	}
}

func TestUserProfileV2(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	carol, err := r.UserProfileV2(ctx, aidCarol)
	if err != nil {
		t.Fatal(err)
	}
	if carol.Address != addrCarol || carol.MintedTokenCount != 1 ||
		carol.OwnedTokenCount != 0 || carol.TradeCount != 1 ||
		carol.TradingVolumeWei != weiOffer1 ||
		carol.ProfitWei != "950000000000000000" || carol.WithdrawalCount != 1 {
		t.Fatalf("carol profile = %+v", carol)
	}

	// Alice minted token 12 and never traded. The legacy accumulator has no
	// row for her (the mint trigger lacks an insert fallback), so the exact
	// mint count must come from the ledger.
	alice, err := r.UserProfileV2(ctx, aidAlice)
	if err != nil {
		t.Fatal(err)
	}
	if alice.MintedTokenCount != 1 || alice.OwnedTokenCount != 1 ||
		alice.TradeCount != 0 || alice.TradingVolumeWei != "0" ||
		alice.ProfitWei != "0" || alice.WithdrawalCount != 0 {
		t.Fatalf("alice profile = %+v", alice)
	}

	// Emma is indexed but has zero RandomWalk activity.
	emma, err := r.UserProfileV2(ctx, 25)
	if err != nil {
		t.Fatal(err)
	}
	if emma.MintedTokenCount != 0 || emma.OwnedTokenCount != 0 ||
		emma.TradingVolumeWei != "0" || emma.ProfitWei != "0" {
		t.Fatalf("emma profile = %+v", emma)
	}

	if _, err := r.UserProfileV2(ctx, 0); err == nil {
		t.Fatal("zero address id accepted")
	}
	if _, err := r.UserProfileV2(ctx, 999999); !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("unknown address id error = %v", err)
	}
}

func TestUserTokensPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	var all []OwnedTokenRecord
	var after *TokenPageCursor
	for {
		page, hasMore, err := r.UserTokensPage(ctx, aidDave, after, 1)
		if err != nil {
			t.Fatal(err)
		}
		all = append(all, page...)
		if !hasMore {
			break
		}
		after = &TokenPageCursor{TokenID: page[len(page)-1].TokenID}
	}
	if len(all) != 2 || all[0].TokenID != 10 || all[1].TokenID != 11 {
		t.Fatalf("dave tokens = %+v", all)
	}
	bought := all[0]
	if bought.TokenName != "Wanderer" || bought.LastPriceWei != weiOffer1 ||
		bought.TradeCount != 1 || bought.TradingVolumeWei != weiOffer1 ||
		!bought.HasMint || bought.MintPriceWei != weiMint10 || bought.MintTs != 1767228600 {
		t.Fatalf("dave token 10 = %+v", bought)
	}

	empty, hasMore, err := r.UserTokensPage(ctx, 25, nil, 50)
	if err != nil || hasMore || len(empty) != 0 {
		t.Fatalf("emma tokens = (%v,%v,%v)", empty, hasMore, err)
	}
	if _, _, err := r.UserTokensPage(ctx, 0, nil, 1); err == nil {
		t.Fatal("zero address id accepted")
	}
}

func TestStatisticsV2(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	record, err := r.StatisticsV2(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if record.MintedCount != 4 || record.UniqueOwnerCount != 3 ||
		record.TokenTradeCount != 1 || record.TokenTradingVolumeWei != weiOffer1 ||
		record.MintFundsWei != "230000000000000000" {
		t.Fatalf("token statistics = %+v", record)
	}
	if record.MarketTradeCount != 1 || record.MarketTradingVolumeWei != weiOffer1 ||
		record.ActiveSellOfferCount != 2 || record.ActiveBuyOfferCount != 0 {
		t.Fatalf("marketplace statistics = %+v", record)
	}
	if record.WithdrawalCount != 1 || record.LatestWithdrawal == nil ||
		record.LatestWithdrawal.AmountWei != weiWithdraw ||
		record.LatestWithdrawal.WithdrawerAddr != addrCarol ||
		record.LatestWithdrawal.TokenID != 10 {
		t.Fatalf("withdrawal statistics = %+v", record)
	}
	if record.LastMint == nil || record.LastMint.TokenID != 13 ||
		record.LastMint.PriceWei != weiMint13 || record.LastMint.MinterAddr != addrBob {
		t.Fatalf("last mint = %+v", record.LastMint)
	}
}

func TestTradingVolumeSeries(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// The only purchase happened at 1767229200; a window starting after the
	// fixture mint era carries it as base volume.
	base, buckets, err := r.TradingVolumeSeries(ctx, 1767229200, 1767229500, 100)
	if err != nil {
		t.Fatal(err)
	}
	if base != "0" || len(buckets) != 3 {
		t.Fatalf("base=%q buckets=%+v", base, buckets)
	}
	if buckets[0].BucketStart != 1767229200 || buckets[0].TradeCount != 1 ||
		buckets[0].VolumeWei != weiOffer1 {
		t.Fatalf("bucket[0] = %+v", buckets[0])
	}
	if buckets[1].TradeCount != 0 || buckets[1].VolumeWei != "0" ||
		buckets[2].TradeCount != 0 {
		t.Fatalf("zero-filled buckets = %+v", buckets[1:])
	}

	base, buckets, err = r.TradingVolumeSeries(ctx, 1767229300, 1767229400, 100)
	if err != nil {
		t.Fatal(err)
	}
	if base != weiOffer1 || len(buckets) != 1 || buckets[0].TradeCount != 0 {
		t.Fatalf("post-trade window base=%q buckets=%+v", base, buckets)
	}

	if _, _, err := r.TradingVolumeSeries(ctx, 100, 100, 60); err == nil {
		t.Fatal("empty window accepted")
	}
}

func TestListingFloorSeries(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Sell listings at 1767229100 (1 ETH), 1767229300 (2), 1767229400
	// (2.5), 1767229500 (3). 200-second buckets anchored at 1767229100.
	points, err := r.ListingFloorSeries(ctx, 1767229100, 1767229700, 200)
	if err != nil {
		t.Fatal(err)
	}
	if len(points) != 3 ||
		points[0].BucketStart != 1767229100 || points[0].FloorWei != weiOffer1 ||
		points[1].BucketStart != 1767229300 || points[1].FloorWei != weiOffer2 ||
		points[2].BucketStart != 1767229500 || points[2].FloorWei != weiOffer4 {
		t.Fatalf("floor points = %+v", points)
	}

	empty, err := r.ListingFloorSeries(ctx, 1767229100, 1767229101, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(empty) != 1 || empty[0].FloorWei != weiOffer1 {
		t.Fatalf("single-second window = %+v", empty)
	}
	if _, err := r.ListingFloorSeries(ctx, -1, 100, 60); err == nil {
		t.Fatal("negative window accepted")
	}
}

func TestMintReportV2(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	months, err := r.MintReportV2(ctx)
	if err != nil {
		t.Fatal(err)
	}
	// All four fixture mints land in January 2026 (UTC).
	if len(months) != 1 || months[0].Year != 2026 || months[0].Month != 1 ||
		months[0].MintCount != 4 || months[0].MintedWei != "230000000000000000" {
		t.Fatalf("mint report = %+v", months)
	}
}

func TestWithdrawalsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	records, hasMore, err := r.WithdrawalsPage(ctx, nil, 50)
	if err != nil || hasMore {
		t.Fatalf("withdrawals err=%v hasMore=%v", err, hasMore)
	}
	if len(records) != 1 {
		t.Fatalf("withdrawals = %+v", records)
	}
	withdrawal := records[0]
	if withdrawal.Tx.EvtLogID != 5095 || withdrawal.WithdrawerAddr != addrCarol ||
		withdrawal.TokenID != 10 || withdrawal.AmountWei != weiWithdraw {
		t.Fatalf("withdrawal = %+v", withdrawal)
	}
	exhausted, hasMore, err := r.WithdrawalsPage(ctx, &EventPageCursor{EventLogID: 5095}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted withdrawals = (%v,%v,%v)", exhausted, hasMore, err)
	}
}

func TestRandomWalkV2RejectsInvalidArguments(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	calls := map[string]func() error{
		"token detail negative id": func() error {
			_, err := r.TokenDetailV2(ctx, -1)
			return err
		},
		"token exists negative id": func() error {
			_, err := r.CollectionTokenExists(ctx, -1)
			return err
		},
		"name changes zero limit": func() error {
			_, _, err := r.TokenNameChangesPageV2(ctx, 10, nil, 0)
			return err
		},
		"name changes invalid cursor": func() error {
			_, _, err := r.TokenNameChangesPageV2(ctx, 10, &EventPageCursor{}, 1)
			return err
		},
		"token events zero limit": func() error {
			_, _, err := r.TokenEventsPage(ctx, 10, nil, 0)
			return err
		},
		"token events invalid cursor": func() error {
			_, _, err := r.TokenEventsPage(ctx, 10, &EventPageCursor{}, 1)
			return err
		},
		"offers zero limit": func() error {
			_, _, err := r.ActiveOffersPage(ctx, OfferSortNewest, nil, 0)
			return err
		},
		"offers invalid sort": func() error {
			_, _, err := r.ActiveOffersPage(ctx, OfferSort("bogus"), nil, 1)
			return err
		},
		"offers cursor without price": func() error {
			_, _, err := r.ActiveOffersPage(ctx, OfferSortPriceAsc, &OfferPageCursor{EventLogID: 1}, 1)
			return err
		},
		"offer history zero limit": func() error {
			_, _, err := r.OfferHistoryPage(ctx, nil, 0)
			return err
		},
		"offer history invalid cursor": func() error {
			_, _, err := r.OfferHistoryPage(ctx, &EventPageCursor{}, 1)
			return err
		},
		"user offers zero aid": func() error {
			_, _, err := r.UserOffersPage(ctx, 0, nil, 1)
			return err
		},
		"trades zero limit": func() error {
			_, _, err := r.TradesPage(ctx, nil, 0)
			return err
		},
		"trades invalid cursor": func() error {
			_, _, err := r.TradesPage(ctx, &EventPageCursor{}, 1)
			return err
		},
		"user tokens invalid cursor": func() error {
			_, _, err := r.UserTokensPage(ctx, aidDave, &TokenPageCursor{TokenID: -1}, 1)
			return err
		},
		"volume series inverted": func() error {
			_, _, err := r.TradingVolumeSeries(ctx, 100, 50, 10)
			return err
		},
		"volume series zero interval": func() error {
			_, _, err := r.TradingVolumeSeries(ctx, 0, 100, 0)
			return err
		},
		"floor series zero interval": func() error {
			_, err := r.ListingFloorSeries(ctx, 0, 100, 0)
			return err
		},
		"withdrawals zero limit": func() error {
			_, _, err := r.WithdrawalsPage(ctx, nil, 0)
			return err
		},
		"withdrawals invalid cursor": func() error {
			_, _, err := r.WithdrawalsPage(ctx, &EventPageCursor{}, 1)
			return err
		},
	}
	for name, call := range calls {
		t.Run(name, func(t *testing.T) {
			if err := call(); err == nil {
				t.Fatal("invalid arguments accepted")
			}
		})
	}
}

// TestRandomWalkV2RequiresContractRegistry proves that every v2 read
// propagates the missing-registry error instead of inventing scope: with the
// rw_contracts row gone, ContractAddrs yields store.ErrNotFound.
func TestRandomWalkV2RequiresContractRegistry(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	var marketplace, randomwalk string
	if err := r.pool().QueryRow(ctx,
		"SELECT marketplace_addr, randomwalk_addr FROM rw_contracts").Scan(&marketplace, &randomwalk); err != nil {
		t.Fatal(err)
	}
	if _, err := r.pool().Exec(ctx, "DELETE FROM rw_contracts"); err != nil {
		t.Fatal(err)
	}
	restore := func() {
		if _, err := r.pool().Exec(context.Background(),
			"INSERT INTO rw_contracts(marketplace_addr, randomwalk_addr) VALUES ($1, $2)",
			marketplace, randomwalk); err != nil {
			t.Errorf("restoring rw_contracts: %v", err)
		}
	}
	defer restore()

	for name, call := range randomWalkV2Calls(r, ctx) {
		t.Run(name, func(t *testing.T) {
			if err := call(); !errors.Is(err, store.ErrNotFound) {
				t.Fatalf("error = %v, want store.ErrNotFound", err)
			}
		})
	}
}

// TestRandomWalkV2PropagatesQueryFailures drives the error arms between the
// successful registry read and each page/snapshot query: hiding the queried
// table makes the registry read succeed and the follow-up query fail, and
// every method must surface that error instead of a partial result.
func TestRandomWalkV2PropagatesQueryFailures(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	hide := func(t *testing.T, table string) {
		t.Helper()
		if _, err := r.pool().Exec(ctx,
			"ALTER TABLE "+table+" RENAME TO "+table+"_hidden"); err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			if _, err := r.pool().Exec(context.Background(),
				"ALTER TABLE "+table+"_hidden RENAME TO "+table); err != nil {
				t.Errorf("restoring %s: %v", table, err)
			}
		})
	}

	t.Run("token reads", func(t *testing.T) {
		hide(t, "rw_mint_evt")
		if _, _, err := r.TokensPage(ctx, TokenFilter{}, TokenSortByID, nil, 1); err == nil {
			t.Error("TokensPage succeeded without its table")
		}
		if _, err := r.TokenDetailV2(ctx, 10); err == nil {
			t.Error("TokenDetailV2 succeeded without its table")
		}
		if _, err := r.CollectionTokenExists(ctx, 10); err == nil {
			t.Error("CollectionTokenExists succeeded without its table")
		}
		if _, _, err := r.TokenEventsPage(ctx, 10, nil, 1); err == nil {
			t.Error("TokenEventsPage succeeded without its table")
		}
		if _, err := r.MintReportV2(ctx); err == nil {
			t.Error("MintReportV2 succeeded without its table")
		}
		if _, err := r.StatisticsV2(ctx); err == nil {
			t.Error("StatisticsV2 succeeded without its table")
		}
		if _, _, err := r.UserTokensPage(ctx, aidDave, nil, 1); err == nil {
			t.Error("UserTokensPage succeeded without its table")
		}
		if _, err := r.UserProfileV2(ctx, aidDave); err == nil {
			t.Error("UserProfileV2 succeeded without its table")
		}
	})

	t.Run("token rename reads", func(t *testing.T) {
		hide(t, "rw_token_name")
		if _, _, err := r.TokenNameChangesPageV2(ctx, 10, nil, 1); err == nil {
			t.Error("TokenNameChangesPageV2 succeeded without its table")
		}
		if _, err := r.TokenDetailV2(ctx, 10); err == nil {
			t.Error("TokenDetailV2 succeeded without the rename table")
		}
	})

	t.Run("marketplace reads", func(t *testing.T) {
		hide(t, "rw_new_offer")
		if _, _, err := r.ActiveOffersPage(ctx, OfferSortNewest, nil, 1); err == nil {
			t.Error("ActiveOffersPage succeeded without its table")
		}
		if _, _, err := r.OfferHistoryPage(ctx, nil, 1); err == nil {
			t.Error("OfferHistoryPage succeeded without its table")
		}
		if _, _, err := r.UserOffersPage(ctx, aidDave, nil, 1); err == nil {
			t.Error("UserOffersPage succeeded without its table")
		}
		if _, _, err := r.TradesPage(ctx, nil, 1); err == nil {
			t.Error("TradesPage succeeded without its table")
		}
		if _, err := r.FloorPriceV2(ctx); err == nil {
			t.Error("FloorPriceV2 succeeded without its table")
		}
		if _, _, err := r.TradingVolumeSeries(ctx, 0, 100, 50); err == nil {
			t.Error("TradingVolumeSeries succeeded without its table")
		}
		if _, err := r.ListingFloorSeries(ctx, 0, 100, 50); err == nil {
			t.Error("ListingFloorSeries succeeded without its table")
		}
	})

	t.Run("withdrawal reads", func(t *testing.T) {
		hide(t, "rw_withdrawal")
		if _, _, err := r.WithdrawalsPage(ctx, nil, 1); err == nil {
			t.Error("WithdrawalsPage succeeded without its table")
		}
	})
}

func TestRandomWalkV2PagesPropagateCancellation(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	calls := randomWalkV2Calls(r, ctx)
	for name, call := range calls {
		t.Run(name, func(t *testing.T) {
			if err := call(); !errors.Is(err, context.Canceled) {
				t.Fatalf("error=%v, want context.Canceled", err)
			}
		})
	}
}

func TestRandomWalkV2PagesFailOnClosedPool(t *testing.T) {
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	ctx := context.Background()
	st, err := spareStore(ctx)
	if err != nil {
		t.Fatal(err)
	}
	r := NewRepo(st)
	st.Close()
	for name, call := range randomWalkV2Calls(r, ctx) {
		t.Run(name, func(t *testing.T) {
			if err := call(); err == nil {
				t.Fatal("closed pool query succeeded")
			}
		})
	}
}

func randomWalkV2Calls(r *Repo, ctx context.Context) map[string]func() error {
	return map[string]func() error{
		"tokens": func() error {
			_, _, err := r.TokensPage(ctx, TokenFilter{}, TokenSortByID, nil, 1)
			return err
		},
		"token detail": func() error {
			_, err := r.TokenDetailV2(ctx, 10)
			return err
		},
		"token exists": func() error {
			_, err := r.CollectionTokenExists(ctx, 10)
			return err
		},
		"name changes": func() error {
			_, _, err := r.TokenNameChangesPageV2(ctx, 10, nil, 1)
			return err
		},
		"token events": func() error {
			_, _, err := r.TokenEventsPage(ctx, 10, nil, 1)
			return err
		},
		"active offers": func() error {
			_, _, err := r.ActiveOffersPage(ctx, OfferSortNewest, nil, 1)
			return err
		},
		"offer history": func() error {
			_, _, err := r.OfferHistoryPage(ctx, nil, 1)
			return err
		},
		"user offers": func() error {
			_, _, err := r.UserOffersPage(ctx, aidDave, nil, 1)
			return err
		},
		"trades": func() error {
			_, _, err := r.TradesPage(ctx, nil, 1)
			return err
		},
		"floor price": func() error {
			_, err := r.FloorPriceV2(ctx)
			return err
		},
		"user profile": func() error {
			_, err := r.UserProfileV2(ctx, aidDave)
			return err
		},
		"user tokens": func() error {
			_, _, err := r.UserTokensPage(ctx, aidDave, nil, 1)
			return err
		},
		"statistics": func() error {
			_, err := r.StatisticsV2(ctx)
			return err
		},
		"volume series": func() error {
			_, _, err := r.TradingVolumeSeries(ctx, 0, 100, 50)
			return err
		},
		"floor series": func() error {
			_, err := r.ListingFloorSeries(ctx, 0, 100, 50)
			return err
		},
		"mint report": func() error {
			_, err := r.MintReportV2(ctx)
			return err
		},
		"withdrawals": func() error {
			_, _, err := r.WithdrawalsPage(ctx, nil, 1)
			return err
		},
	}
}

func TestRandomWalkV2ReadIndexesExist(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	for _, name := range []string{
		"rw_mint_evt_contract_token_idx",
		"rw_mint_evt_owner_idx",
		"rw_token_trades_rank_idx",
		"rw_token_owner_idx",
		"rw_token_name_token_evt_idx",
		"rw_transfer_token_evt_idx",
		"rw_new_offer_token_evt_idx",
		"rw_new_offer_contract_evt_idx",
		"rw_new_offer_active_price_idx",
		"rw_new_offer_seller_evt_idx",
		"rw_new_offer_buyer_evt_idx",
		"rw_new_offer_sell_time_idx",
		"rw_item_bought_offer_idx",
		"rw_item_bought_contract_evt_idx",
		"rw_item_bought_time_idx",
		"rw_offer_canceled_offer_idx",
		"rw_withdrawal_contract_evt_idx",
	} {
		var exists bool
		if err := r.pool().QueryRow(ctx,
			`SELECT EXISTS(SELECT 1 FROM pg_indexes WHERE schemaname='public' AND indexname=$1)`,
			name,
		).Scan(&exists); err != nil {
			t.Fatal(err)
		}
		if !exists {
			t.Errorf("index %s does not exist", name)
		}
	}
}

func int64Pointer(value int64) *int64 { return &value }
