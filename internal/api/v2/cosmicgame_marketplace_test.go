package v2

import (
	"context"
	"errors"
	"log/slog"
	"math/big"
	"net/http"
	"strings"
	"testing"

	marketstore "github.com/PredictionExplorer/augur-explorer/internal/store/marketplace"
)

func newCosmicMarketplaceTestServer(
	t *testing.T,
	marketplace fakeGlobalDirectoryReader,
) *Server {
	t.Helper()
	server, err := newServer(
		nil,
		fakeBidReader{},
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
		fakeRoundRaffleReader{},
		fakeRoundDonationReader{},
		fakeStatisticsReader{},
		fakeBiddingAnalyticsReader{},
		fakeContractAddressReader{},
		fakeParticipantReader{},
		fakeUserReader{},
		fakeUserHistoryReader{},
		fakeUserStakingReader{},
		fakeUserActivityReader{},
		marketplace,
		fakeGlobalStakingReader{},
		fakeRandomWalkReader{},
		fakeRankingRepository{},
		fakeContractState{},
		slog.New(slog.DiscardHandler),
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func validCosmicMarketplaceEvent() marketstore.EventTx {
	return marketstore.EventTx{
		EventLogID: 5203,
		BlockNum:   132,
		TxID:       1046,
		TxHash:     "0xf000000000000000000000000000000000000000000000000000000000001046",
		TimeStamp:  1767228810,
		DateTime:   "2026-01-01T00:53:30Z",
	}
}

func validCosmicMarketplaceOffer() marketstore.OfferRecord {
	return marketstore.OfferRecord{
		ListTx:    validCosmicMarketplaceEvent(),
		OfferID:   102,
		OfferType: 1,
		TokenID:   2,
		PriceWei:  "2000000000000000000",
		MakerAid:  24,
		MakerAddr: rwTestDave,
	}
}

func TestCosmicMarketplaceOffersCursorScopesCollectionAndSort(t *testing.T) {
	t.Parallel()
	var gotAfter *marketstore.OfferPageCursor
	server := newCosmicMarketplaceTestServer(t, fakeGlobalDirectoryReader{
		marketOffers: func(
			_ context.Context,
			sort marketstore.OfferSort,
			after *marketstore.OfferPageCursor,
			limit int,
		) ([]marketstore.OfferRecord, bool, error) {
			if sort != marketstore.OfferSortPriceAsc || limit != 1 {
				t.Errorf("repository args = (%s,%d)", sort, limit)
			}
			gotAfter = after
			if after != nil {
				return []marketstore.OfferRecord{}, false, nil
			}
			return []marketstore.OfferRecord{validCosmicMarketplaceOffer()}, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/marketplace/offers?sort=priceAsc&limit=1")
	if response.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", response.Code, response.Body.String())
	}
	var page CosmicSignatureMarketplaceOfferPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 1 || page.Meta.NextCursor == nil ||
		page.Data[0].OfferId != 102 || page.Data[0].PriceWei != "2000000000000000000" {
		t.Fatalf("page = %+v", page)
	}
	continued := serve(t, server,
		"/api/v2/cosmicgame/marketplace/offers?sort=priceAsc&limit=1&cursor="+
			*page.Meta.NextCursor)
	if continued.Code != http.StatusOK || gotAfter == nil ||
		gotAfter.EventLogID != 5203 || gotAfter.PriceWei != "2000000000000000000" {
		t.Fatalf("continuation=%d after=%+v", continued.Code, gotAfter)
	}
	assertProblem(t, serve(t, server,
		"/api/v2/cosmicgame/marketplace/offers?sort=priceDesc&cursor="+
			*page.Meta.NextCursor), http.StatusBadRequest)

	randomWalkCursor, err := encodeRandomWalkOfferBookCursor(randomWalkOfferBookCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceOfferBook,
		Sort:       PriceAsc,
		PriceWei:   "2000000000000000000",
		EventLogID: 5203,
	})
	if err != nil {
		t.Fatal(err)
	}
	assertProblem(t, serve(t, server,
		"/api/v2/cosmicgame/marketplace/offers?sort=priceAsc&cursor="+randomWalkCursor),
		http.StatusBadRequest)
}

func TestCosmicMarketplaceLedgersAreEndpointScoped(t *testing.T) {
	t.Parallel()
	record := marketstore.OfferHistoryRecord{
		ListTx:    validCosmicMarketplaceEvent(),
		OfferID:   102,
		OfferType: 1,
		TokenID:   2,
		PriceWei:  "2000000000000000000",
		MakerAid:  24,
		MakerAddr: rwTestDave,
		Active:    true,
	}
	var gotAfter *marketstore.EventPageCursor
	server := newCosmicMarketplaceTestServer(t, fakeGlobalDirectoryReader{
		marketHistory: func(
			_ context.Context,
			after *marketstore.EventPageCursor,
			limit int,
		) ([]marketstore.OfferHistoryRecord, bool, error) {
			gotAfter = after
			if limit != 1 {
				t.Errorf("limit = %d", limit)
			}
			if after != nil {
				return []marketstore.OfferHistoryRecord{}, false, nil
			}
			return []marketstore.OfferHistoryRecord{record}, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/marketplace/offer-history?limit=1")
	if response.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", response.Code, response.Body.String())
	}
	var page CosmicSignatureMarketplaceOfferHistoryPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 1 || page.Meta.NextCursor == nil ||
		page.Data[0].Status != Active {
		t.Fatalf("page = %+v", page)
	}
	continued := serve(t, server,
		"/api/v2/cosmicgame/marketplace/offer-history?limit=1&cursor="+
			*page.Meta.NextCursor)
	if continued.Code != http.StatusOK || gotAfter == nil || gotAfter.EventLogID != 5203 {
		t.Fatalf("continuation=%d after=%+v", continued.Code, gotAfter)
	}

	tradesCursor, err := encodeCosmicMarketplaceLedgerCursor(cosmicMarketplaceLedgerCursor{
		Version:    cosmicMarketplaceCursorVersion,
		Resource:   cosmicMarketplaceResourceTrades,
		Collection: cosmicSignatureCollectionScope,
		EventLogID: 5202,
	})
	if err != nil {
		t.Fatal(err)
	}
	assertProblem(t, serve(t, server,
		"/api/v2/cosmicgame/marketplace/offer-history?cursor="+tradesCursor),
		http.StatusBadRequest)
}

func TestCosmicMarketplaceEmptyAndErrorBehavior(t *testing.T) {
	t.Parallel()
	empty := newCosmicMarketplaceTestServer(t, fakeGlobalDirectoryReader{})
	for _, path := range []string{
		"/api/v2/cosmicgame/marketplace/offers",
		"/api/v2/cosmicgame/marketplace/offer-history",
		"/api/v2/cosmicgame/marketplace/trades",
	} {
		response := serve(t, empty, path)
		if response.Code != http.StatusOK ||
			!strings.Contains(response.Body.String(), `"data":[]`) {
			t.Fatalf("%s = %d %s", path, response.Code, response.Body.String())
		}
	}
	floorResponse := serve(t, empty, "/api/v2/cosmicgame/marketplace/floor-price")
	if floorResponse.Code != http.StatusOK {
		t.Fatalf("floor = %d %s", floorResponse.Code, floorResponse.Body.String())
	}
	var floor CosmicSignatureMarketplaceFloorPrice
	decodeResponse(t, floorResponse, &floor)
	if floor.ActiveSellOfferCount != 0 || floor.Floor != nil {
		t.Fatalf("floor = %+v", floor)
	}

	failing := newCosmicMarketplaceTestServer(t, fakeGlobalDirectoryReader{
		marketTrades: func(
			context.Context,
			*marketstore.EventPageCursor,
			int,
		) ([]marketstore.TradeRecord, bool, error) {
			return nil, false, errors.New("password=secret")
		},
	})
	response := serve(t, failing, "/api/v2/cosmicgame/marketplace/trades")
	assertProblem(t, response, http.StatusInternalServerError)
	if strings.Contains(response.Body.String(), "secret") {
		t.Fatalf("internal error leaked: %s", response.Body.String())
	}

	for _, path := range []string{
		"/api/v2/cosmicgame/marketplace/offers?sort=bad",
		"/api/v2/cosmicgame/marketplace/offers?limit=0",
		"/api/v2/cosmicgame/marketplace/offers?cursor=bogus",
		"/api/v2/cosmicgame/marketplace/offer-history?limit=201",
		"/api/v2/cosmicgame/marketplace/trades?cursor=bogus",
	} {
		assertProblem(t, serve(t, empty, path), http.StatusBadRequest)
	}
}

func TestCosmicMarketplaceRejectsInconsistentRepositoryPages(t *testing.T) {
	t.Parallel()
	t.Run("unordered offers", func(t *testing.T) {
		t.Parallel()
		first := validCosmicMarketplaceOffer()
		second := validCosmicMarketplaceOffer()
		second.ListTx.EventLogID = first.ListTx.EventLogID + 1
		server := newCosmicMarketplaceTestServer(t, fakeGlobalDirectoryReader{
			marketOffers: func(
				context.Context,
				marketstore.OfferSort,
				*marketstore.OfferPageCursor,
				int,
			) ([]marketstore.OfferRecord, bool, error) {
				return []marketstore.OfferRecord{first, second}, false, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/cosmicgame/marketplace/offers?sort=newest"),
			http.StatusInternalServerError)
	})
	t.Run("over cardinality", func(t *testing.T) {
		t.Parallel()
		first := validCosmicMarketplaceOffer()
		second := validCosmicMarketplaceOffer()
		second.ListTx.EventLogID--
		server := newCosmicMarketplaceTestServer(t, fakeGlobalDirectoryReader{
			marketOffers: func(
				context.Context,
				marketstore.OfferSort,
				*marketstore.OfferPageCursor,
				int,
			) ([]marketstore.OfferRecord, bool, error) {
				return []marketstore.OfferRecord{first, second}, false, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/cosmicgame/marketplace/offers?limit=1"),
			http.StatusInternalServerError)
	})
	t.Run("inconsistent floor", func(t *testing.T) {
		t.Parallel()
		server := newCosmicMarketplaceTestServer(t, fakeGlobalDirectoryReader{
			marketFloor: func(context.Context) (marketstore.FloorPriceRecord, error) {
				return marketstore.FloorPriceRecord{ActiveSellOfferCount: 1}, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/cosmicgame/marketplace/floor-price"),
			http.StatusInternalServerError)
	})
	t.Run("continuation has more without rows", func(t *testing.T) {
		t.Parallel()
		cursor, err := encodeCosmicMarketplaceLedgerCursor(cosmicMarketplaceLedgerCursor{
			Version:    cosmicMarketplaceCursorVersion,
			Resource:   cosmicMarketplaceResourceOfferHistory,
			Collection: cosmicSignatureCollectionScope,
			EventLogID: 5203,
		})
		if err != nil {
			t.Fatal(err)
		}
		server := newCosmicMarketplaceTestServer(t, fakeGlobalDirectoryReader{
			marketHistory: func(
				context.Context,
				*marketstore.EventPageCursor,
				int,
			) ([]marketstore.OfferHistoryRecord, bool, error) {
				return []marketstore.OfferHistoryRecord{}, true, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/cosmicgame/marketplace/offer-history?cursor="+cursor),
			http.StatusInternalServerError)
	})
}

func TestCosmicMarketplaceOfferOrderingGuards(t *testing.T) {
	t.Parallel()
	record := validCosmicMarketplaceOffer()
	price := big.NewInt(20)

	record.ListTx.EventLogID = 0
	if cosmicMarketplaceOfferOrdered(Newest, nil, 10, record, nil) {
		t.Fatal("zero event-log id accepted")
	}
	record.ListTx.EventLogID = 11
	if !cosmicMarketplaceOfferOrdered(Oldest, nil, 10, record, nil) {
		t.Fatal("ascending event position rejected")
	}
	if !cosmicMarketplaceOfferOrdered(
		PriceAsc,
		big.NewInt(10),
		10,
		record,
		price,
	) {
		t.Fatal("ascending price position rejected")
	}
	if !cosmicMarketplaceOfferOrdered(
		PriceDesc,
		big.NewInt(30),
		10,
		record,
		price,
	) {
		t.Fatal("descending price position rejected")
	}
	if cosmicMarketplaceOfferOrdered(
		RandomWalkOfferSort("invalid"),
		nil,
		10,
		record,
		nil,
	) {
		t.Fatal("unknown sort accepted")
	}
}

func TestCosmicMarketplaceCursorEncodersRejectInvalidFields(t *testing.T) {
	t.Parallel()
	if _, err := encodeCosmicMarketplaceOfferCursor(cosmicMarketplaceOfferCursor{}); err == nil {
		t.Fatal("invalid offer cursor encoded")
	}
	if _, err := encodeCosmicMarketplaceLedgerCursor(cosmicMarketplaceLedgerCursor{}); err == nil {
		t.Fatal("invalid ledger cursor encoded")
	}
}
