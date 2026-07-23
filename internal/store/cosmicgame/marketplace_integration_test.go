//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"reflect"
	"testing"

	marketstore "github.com/PredictionExplorer/augur-explorer/internal/store/marketplace"
)

func marketplaceOfferIDs(records []marketstore.OfferRecord) []int64 {
	ids := make([]int64, 0, len(records))
	for i := range records {
		ids = append(ids, records[i].OfferID)
	}
	return ids
}

func walkCosmicMarketplaceOffers(
	t *testing.T,
	r *Repo,
	sort marketstore.OfferSort,
	pageSize int,
) []marketstore.OfferRecord {
	t.Helper()
	var all []marketstore.OfferRecord
	var after *marketstore.OfferPageCursor
	for {
		page, hasMore, err := r.CosmicSignatureMarketplaceOffersPage(
			context.Background(),
			sort,
			after,
			pageSize,
		)
		if err != nil {
			t.Fatalf("offers page: %v", err)
		}
		all = append(all, page...)
		if !hasMore {
			return all
		}
		if len(page) == 0 {
			t.Fatal("hasMore without a cursor row")
		}
		last := page[len(page)-1]
		after = &marketstore.OfferPageCursor{EventLogID: last.ListTx.EventLogID}
		if sort == marketstore.OfferSortPriceAsc || sort == marketstore.OfferSortPriceDesc {
			after.PriceWei = last.PriceWei
		}
	}
}

func TestCosmicSignatureMarketplaceScopeAndLegacyReads(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	scope, err := r.CosmicSignatureMarketplaceScope(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if scope.MarketplaceAid != 12 || scope.CollectionAid != 3 {
		t.Fatalf("scope = %+v", scope)
	}

	offers, gotScope, err := r.CosmicSignatureMarketplaceActiveOffers(ctx, 2)
	if err != nil {
		t.Fatal(err)
	}
	if gotScope != scope || len(offers) != 3 ||
		offers[0].OfferID != 103 || offers[1].OfferID != 102 ||
		offers[2].OfferID != 105 {
		t.Fatalf("active offers = %+v scope=%+v", offers, gotScope)
	}
	if offers[0].CollectionAid != 3 || offers[0].PriceETH != 0.75 {
		t.Fatalf("cheapest offer = %+v", offers[0])
	}

	noOffers, floor, offerID, tokenID, floorScope, err := r.CosmicSignatureMarketplaceFloorPriceETH(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if noOffers || floor != 2 || offerID != 102 || tokenID != 2 || floorScope != scope {
		t.Fatalf("floor = (%v,%v,%v,%v,%+v)", noOffers, floor, offerID, tokenID, floorScope)
	}

	sales, saleScope, err := r.CosmicSignatureMarketplaceSales(ctx, 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	if saleScope != scope || len(sales) != 1 || sales[0].OfferID != 101 ||
		sales[0].TokenID != 1 || sales[0].PriceETH != 1.5 ||
		sales[0].SellerAddress != "0x2100000000000000000000000000000000000021" ||
		sales[0].BuyerAddress != "0x2200000000000000000000000000000000000022" {
		t.Fatalf("sales = %+v scope=%+v", sales, saleScope)
	}
}

func TestCosmicSignatureMarketplaceV2PagesAreCollectionScoped(t *testing.T) {
	r := repo(t)
	cases := map[marketstore.OfferSort][]int64{
		marketstore.OfferSortNewest:    {105, 103, 102},
		marketstore.OfferSortOldest:    {102, 103, 105},
		marketstore.OfferSortPriceAsc:  {103, 102, 105},
		marketstore.OfferSortPriceDesc: {102, 105, 103},
	}
	for sort, want := range cases {
		for _, size := range []int{1, 2, 50} {
			got := marketplaceOfferIDs(walkCosmicMarketplaceOffers(t, r, sort, size))
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("sort=%s size=%d offers=%v want=%v", sort, size, got, want)
			}
		}
	}

	ctx := context.Background()
	history, hasMore, err := r.CosmicSignatureMarketplaceOfferHistoryPage(ctx, nil, 50)
	if err != nil || hasMore {
		t.Fatalf("history = (%+v,%v,%v)", history, hasMore, err)
	}
	if len(history) != 5 ||
		history[0].OfferID != 105 ||
		history[1].OfferID != 104 || history[1].Cancellation == nil ||
		history[4].OfferID != 101 || history[4].Purchase == nil {
		t.Fatalf("history = %+v", history)
	}

	trades, hasMore, err := r.CosmicSignatureMarketplaceTradesPage(ctx, nil, 50)
	if err != nil || hasMore || len(trades) != 1 ||
		trades[0].OfferID != 101 || trades[0].Tx.EventLogID != 5202 ||
		trades[0].PriceWei != "1500000000000000000" {
		t.Fatalf("trades = (%+v,%v,%v)", trades, hasMore, err)
	}
	floor, err := r.CosmicSignatureMarketplaceFloorPrice(ctx)
	if err != nil || floor.ActiveSellOfferCount != 2 || floor.Floor == nil ||
		floor.Floor.OfferID != 102 || floor.Floor.PriceWei != "2000000000000000000" {
		t.Fatalf("floor = (%+v,%v)", floor, err)
	}
}

func TestCosmicSignatureMarketplaceErrorsAndIndexes(t *testing.T) {
	r := repo(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	calls := func(repo *Repo, ctx context.Context) map[string]func() error {
		return map[string]func() error{
			"offers": func() error {
				_, _, err := repo.CosmicSignatureMarketplaceOffersPage(
					ctx, marketstore.OfferSortNewest, nil, 1)
				return err
			},
			"history": func() error {
				_, _, err := repo.CosmicSignatureMarketplaceOfferHistoryPage(ctx, nil, 1)
				return err
			},
			"trades": func() error {
				_, _, err := repo.CosmicSignatureMarketplaceTradesPage(ctx, nil, 1)
				return err
			},
			"floor": func() error {
				_, err := repo.CosmicSignatureMarketplaceFloorPrice(ctx)
				return err
			},
		}
	}
	for name, call := range calls(r, cancelled) {
		t.Run("canceled "+name, func(t *testing.T) {
			if err := call(); !errors.Is(err, context.Canceled) {
				t.Fatalf("error = %v, want context.Canceled", err)
			}
		})
	}

	st, err := spareStore(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	closedRepo := NewRepo(st)
	st.Close()
	for name, call := range calls(closedRepo, context.Background()) {
		t.Run("closed pool "+name, func(t *testing.T) {
			if err := call(); err == nil {
				t.Fatal("closed-pool marketplace read succeeded")
			}
		})
	}

	for _, name := range []string{
		"rw_new_offer_collection_evt_idx",
		"rw_new_offer_collection_active_price_idx",
	} {
		var exists bool
		if err := r.pool().QueryRow(context.Background(),
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
