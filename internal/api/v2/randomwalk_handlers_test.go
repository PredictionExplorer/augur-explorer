package v2

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"

	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

func newRandomWalkTestServer(t *testing.T, randomWalk randomWalkReader) *Server {
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
		fakeGlobalDirectoryReader{},
		fakeGlobalStakingReader{},
		randomWalk,
		fakeContractState{},
		slog.New(slog.NewTextHandler(io.Discard, nil)),
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func TestListRandomWalkTokensPaginatesAndScopesCursor(t *testing.T) {
	t.Parallel()

	first := validRandomWalkTokenRecord()
	second := validRandomWalkTokenRecord()
	second.TokenID = 11
	second.MintTx.EvtLogID = 5082
	second.TokenName = ""
	second.TradeCount = 0
	second.LastPriceWei = "55000000000000000"
	second.TradingVolumeWei = "0"

	var gotFilter rwstore.TokenFilter
	var gotSort rwstore.TokenSort
	var gotLimit int
	server := newRandomWalkTestServer(t, fakeRandomWalkReader{
		tokens: func(_ context.Context, filter rwstore.TokenFilter, sort rwstore.TokenSort, after *rwstore.TokenPageCursor, limit int) ([]rwstore.TokenRecord, bool, error) {
			gotFilter, gotSort, gotLimit = filter, sort, limit
			if after != nil {
				t.Errorf("first page carried a cursor: %+v", after)
			}
			return []rwstore.TokenRecord{first, second}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/randomwalk/tokens?limit=2&named=true")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if !gotFilter.NamedOnly || gotSort != rwstore.TokenSortByID || gotLimit != 2 {
		t.Fatalf("repository args = (%+v,%v,%d)", gotFilter, gotSort, gotLimit)
	}
	var page RandomWalkTokenPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeRandomWalkTokenCursor(
		*page.Meta.NextCursor, randomWalkTokenFilterScope{Named: true}, TokenId)
	if err != nil || cursor.TokenID != 11 || cursor.TradeCount != 0 {
		t.Fatalf("next cursor = %+v, %v", cursor, err)
	}
	// The cursor is bound to the filter: replaying it unfiltered is a 400.
	replayed := serve(t, server, "/api/v2/randomwalk/tokens?cursor="+*page.Meta.NextCursor)
	assertProblem(t, replayed, http.StatusBadRequest)
}

func TestListRandomWalkTokensMostTradedCursorCarriesRank(t *testing.T) {
	t.Parallel()

	first := validRandomWalkTokenRecord()
	first.TradeCount = 5
	var gotAfter *rwstore.TokenPageCursor
	server := newRandomWalkTestServer(t, fakeRandomWalkReader{
		tokens: func(_ context.Context, _ rwstore.TokenFilter, sort rwstore.TokenSort, after *rwstore.TokenPageCursor, _ int) ([]rwstore.TokenRecord, bool, error) {
			if sort != rwstore.TokenSortByTrades {
				t.Errorf("sort = %v", sort)
			}
			gotAfter = after
			if after != nil {
				return []rwstore.TokenRecord{}, false, nil
			}
			return []rwstore.TokenRecord{first}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/randomwalk/tokens?sort=mostTraded&limit=1")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var page RandomWalkTokenPage
	decodeResponse(t, response, &page)
	if page.Meta.NextCursor == nil {
		t.Fatal("no continuation cursor")
	}
	next := serve(t, server, "/api/v2/randomwalk/tokens?sort=mostTraded&limit=1&cursor="+*page.Meta.NextCursor)
	if next.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", next.Code, next.Body.String())
	}
	if gotAfter == nil || gotAfter.TokenID != 10 || gotAfter.TradeCount != 5 {
		t.Fatalf("continuation repository cursor = %+v", gotAfter)
	}
}

func TestListRandomWalkTokensRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	tests := map[string]string{
		"invalid sort":        "/api/v2/randomwalk/tokens?sort=priceAsc",
		"conflicting filters": "/api/v2/randomwalk/tokens?named=true&name=x",
		"empty name":          "/api/v2/randomwalk/tokens?name=",
		"oversized name":      "/api/v2/randomwalk/tokens?name=" + strings.Repeat("a", 65),
		"inverted window":     "/api/v2/randomwalk/tokens?mintedFrom=100&mintedUntil=100",
		"negative from":       "/api/v2/randomwalk/tokens?mintedFrom=-1",
		"zero limit":          "/api/v2/randomwalk/tokens?limit=0",
		"excessive limit":     "/api/v2/randomwalk/tokens?limit=201",
		"malformed cursor":    "/api/v2/randomwalk/tokens?cursor=bogus",
	}
	for name, path := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}), path)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestListRandomWalkTokensRejectsInconsistentRepository(t *testing.T) {
	t.Parallel()

	t.Run("unordered rows", func(t *testing.T) {
		t.Parallel()
		first := validRandomWalkTokenRecord()
		second := validRandomWalkTokenRecord()
		second.TokenID = 9
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokens: func(context.Context, rwstore.TokenFilter, rwstore.TokenSort, *rwstore.TokenPageCursor, int) ([]rwstore.TokenRecord, bool, error) {
				return []rwstore.TokenRecord{first, second}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens"), http.StatusInternalServerError)
	})

	t.Run("rank inversion under mostTraded", func(t *testing.T) {
		t.Parallel()
		first := validRandomWalkTokenRecord()
		first.TradeCount = 1
		second := validRandomWalkTokenRecord()
		second.TokenID = 11
		second.TradeCount = 2
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokens: func(context.Context, rwstore.TokenFilter, rwstore.TokenSort, *rwstore.TokenPageCursor, int) ([]rwstore.TokenRecord, bool, error) {
				return []rwstore.TokenRecord{first, second}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens?sort=mostTraded"),
			http.StatusInternalServerError)
	})

	t.Run("has more without rows", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokens: func(context.Context, rwstore.TokenFilter, rwstore.TokenSort, *rwstore.TokenPageCursor, int) ([]rwstore.TokenRecord, bool, error) {
				return []rwstore.TokenRecord{}, true, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens"), http.StatusInternalServerError)
	})

	t.Run("repository failure stays opaque", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokens: func(context.Context, rwstore.TokenFilter, rwstore.TokenSort, *rwstore.TokenPageCursor, int) ([]rwstore.TokenRecord, bool, error) {
				return nil, false, errors.New("password=super-secret")
			},
		})
		response := serve(t, server, "/api/v2/randomwalk/tokens")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "super-secret") {
			t.Fatalf("internal detail leaked: %s", response.Body.String())
		}
	})
}

func TestGetRandomWalkTokenResponses(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokenDetail: func(_ context.Context, tokenID int64) (rwstore.TokenDetailRecord, error) {
				if tokenID != 10 {
					t.Fatalf("token id = %d", tokenID)
				}
				return rwstore.TokenDetailRecord{TokenRecord: validRandomWalkTokenRecord()}, nil
			},
		})
		response := serve(t, server, "/api/v2/randomwalk/tokens/10")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var detail RandomWalkTokenDetail
		decodeResponse(t, response, &detail)
		if detail.NftTokenId != 10 || detail.TokenName == nil {
			t.Fatalf("detail = %+v", detail)
		}
	})

	t.Run("not found", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/tokens/999")
		assertProblem(t, response, http.StatusNotFound)
	})

	t.Run("negative id", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/tokens/-1")
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("wrong token from repository", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokenDetail: func(context.Context, int64) (rwstore.TokenDetailRecord, error) {
				return rwstore.TokenDetailRecord{TokenRecord: validRandomWalkTokenRecord()}, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens/11"),
			http.StatusInternalServerError)
	})
}

func TestRandomWalkTokenScopedCollectionsGateOnExistence(t *testing.T) {
	t.Parallel()
	for _, path := range []string{
		"/api/v2/randomwalk/tokens/10/name-history",
		"/api/v2/randomwalk/tokens/10/events",
	} {
		t.Run(path, func(t *testing.T) {
			t.Parallel()
			missing := newRandomWalkTestServer(t, fakeRandomWalkReader{
				tokenExists: func(context.Context, int64) (bool, error) { return false, nil },
			})
			assertProblem(t, serve(t, missing, path), http.StatusNotFound)

			failing := newRandomWalkTestServer(t, fakeRandomWalkReader{
				tokenExists: func(context.Context, int64) (bool, error) {
					return false, errors.New("gate failure")
				},
			})
			assertProblem(t, serve(t, failing, path), http.StatusInternalServerError)

			empty := newRandomWalkTestServer(t, fakeRandomWalkReader{})
			response := serve(t, empty, path)
			if response.Code != http.StatusOK ||
				!strings.Contains(response.Body.String(), `"data":[]`) {
				t.Fatalf("existing token page = %d %s", response.Code, response.Body.String())
			}
		})
	}
}

func TestListRandomWalkTokenEventsPaginates(t *testing.T) {
	t.Parallel()
	first := validRandomWalkTokenEventRecord(rwstore.TokenEventPurchase)
	first.Tx.EvtLogID = 5089
	second := validRandomWalkTokenEventRecord(rwstore.TokenEventMint)
	second.Tx.EvtLogID = 5080

	var gotAfter *rwstore.EventPageCursor
	server := newRandomWalkTestServer(t, fakeRandomWalkReader{
		tokenEvents: func(_ context.Context, tokenID int64, after *rwstore.EventPageCursor, limit int) ([]rwstore.TokenEventRecord, bool, error) {
			if tokenID != 10 || limit != 2 {
				t.Errorf("repository args = (%d,%d)", tokenID, limit)
			}
			gotAfter = after
			if after != nil {
				return []rwstore.TokenEventRecord{}, false, nil
			}
			return []rwstore.TokenEventRecord{first, second}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/randomwalk/tokens/10/events?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var page RandomWalkTokenEventPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	if page.Data[0].EventType != RandomWalkTokenEventTypePurchase ||
		page.Data[1].EventType != RandomWalkTokenEventTypeMint {
		t.Fatalf("event types = %+v", page.Data)
	}

	continued := serve(t, server,
		"/api/v2/randomwalk/tokens/10/events?limit=2&cursor="+*page.Meta.NextCursor)
	if continued.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", continued.Code, continued.Body.String())
	}
	if gotAfter == nil || gotAfter.EventLogID != 5080 {
		t.Fatalf("continuation cursor = %+v", gotAfter)
	}

	// A cursor minted for another token is rejected.
	crossToken, err := encodeRandomWalkTokenEventCursor(randomWalkTokenEventCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceTokenEvents,
		TokenID:    11,
		EventLogID: 5080,
	})
	if err != nil {
		t.Fatal(err)
	}
	assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens/10/events?cursor="+crossToken),
		http.StatusBadRequest)
	// So is the name-history cursor of the same token.
	crossResource, err := encodeRandomWalkTokenEventCursor(randomWalkTokenEventCursor{
		Version:    randomWalkCursorVersion,
		Resource:   randomWalkResourceNameHistory,
		TokenID:    10,
		EventLogID: 5080,
	})
	if err != nil {
		t.Fatal(err)
	}
	assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens/10/events?cursor="+crossResource),
		http.StatusBadRequest)
}

func TestListRandomWalkMarketplaceOffersSortsAndValidates(t *testing.T) {
	t.Parallel()

	offer := func(evtlog int64, price string) rwstore.OfferRecord {
		record := rwstore.OfferRecord{
			ListTx:    validRandomWalkEventTx(),
			OfferID:   evtlog % 100,
			OfferType: 1,
			TokenID:   11,
			PriceWei:  price,
			MakerAid:  24,
			MakerAddr: rwTestDave,
		}
		record.ListTx.EvtLogID = evtlog
		return record
	}

	t.Run("price sort cursor round trip", func(t *testing.T) {
		t.Parallel()
		var gotSort rwstore.OfferSort
		var gotAfter *rwstore.OfferPageCursor
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			activeOffers: func(_ context.Context, sort rwstore.OfferSort, after *rwstore.OfferPageCursor, _ int) ([]rwstore.OfferRecord, bool, error) {
				gotSort, gotAfter = sort, after
				if after != nil {
					return []rwstore.OfferRecord{}, false, nil
				}
				return []rwstore.OfferRecord{offer(5091, "2000000000000000000")}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/randomwalk/marketplace/offers?sort=priceAsc&limit=1")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		if gotSort != rwstore.OfferSortPriceAsc {
			t.Fatalf("sort = %v", gotSort)
		}
		var page RandomWalkMarketplaceOfferPage
		decodeResponse(t, response, &page)
		if page.Meta.NextCursor == nil {
			t.Fatal("no continuation cursor")
		}
		continued := serve(t, server,
			"/api/v2/randomwalk/marketplace/offers?sort=priceAsc&limit=1&cursor="+*page.Meta.NextCursor)
		if continued.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", continued.Code, continued.Body.String())
		}
		if gotAfter == nil || gotAfter.PriceWei != "2000000000000000000" || gotAfter.EventLogID != 5091 {
			t.Fatalf("continuation cursor = %+v", gotAfter)
		}
		// The same cursor under another sort is rejected.
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/marketplace/offers?sort=priceDesc&cursor="+*page.Meta.NextCursor),
			http.StatusBadRequest)
	})

	t.Run("unordered price page", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			activeOffers: func(context.Context, rwstore.OfferSort, *rwstore.OfferPageCursor, int) ([]rwstore.OfferRecord, bool, error) {
				return []rwstore.OfferRecord{
					offer(5091, "2000000000000000000"),
					offer(5092, "1000000000000000000"),
				}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/marketplace/offers?sort=priceAsc"),
			http.StatusInternalServerError)
	})

	t.Run("unordered event page", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			activeOffers: func(context.Context, rwstore.OfferSort, *rwstore.OfferPageCursor, int) ([]rwstore.OfferRecord, bool, error) {
				return []rwstore.OfferRecord{
					offer(5091, "2000000000000000000"),
					offer(5092, "1000000000000000000"),
				}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/marketplace/offers?sort=newest"),
			http.StatusInternalServerError)
	})

	t.Run("invalid sort", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/marketplace/offers?sort=cheapest"), http.StatusBadRequest)
	})
}

func TestRandomWalkLedgersPaginateAndScope(t *testing.T) {
	t.Parallel()

	t.Run("offer history", func(t *testing.T) {
		t.Parallel()
		record := validRandomWalkOfferHistoryRecord()
		var gotAfter *rwstore.EventPageCursor
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			offerHistory: func(_ context.Context, after *rwstore.EventPageCursor, limit int) ([]rwstore.OfferHistoryRecord, bool, error) {
				gotAfter = after
				if limit != 1 {
					t.Errorf("limit = %d", limit)
				}
				if after != nil {
					return []rwstore.OfferHistoryRecord{}, false, nil
				}
				return []rwstore.OfferHistoryRecord{record}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/randomwalk/marketplace/offer-history?limit=1")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var page RandomWalkOfferHistoryPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 1 || page.Data[0].Status != Active || page.Meta.NextCursor == nil {
			t.Fatalf("page = %+v", page)
		}
		continued := serve(t, server,
			"/api/v2/randomwalk/marketplace/offer-history?limit=1&cursor="+*page.Meta.NextCursor)
		if continued.Code != http.StatusOK || gotAfter == nil ||
			gotAfter.EventLogID != record.ListTx.EvtLogID {
			t.Fatalf("continuation = %d %+v", continued.Code, gotAfter)
		}
		// A trades cursor cannot continue the offer ledger.
		tradesCursor, err := encodeRandomWalkLedgerCursor(randomWalkLedgerCursor{
			Version:    randomWalkCursorVersion,
			Resource:   randomWalkResourceTrades,
			EventLogID: 5089,
		})
		if err != nil {
			t.Fatal(err)
		}
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/marketplace/offer-history?cursor="+tradesCursor),
			http.StatusBadRequest)
	})

	t.Run("trades", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			trades: func(context.Context, *rwstore.EventPageCursor, int) ([]rwstore.TradeRecord, bool, error) {
				return nil, false, errors.New("database detail")
			},
		})
		response := serve(t, server, "/api/v2/randomwalk/marketplace/trades")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "database detail") {
			t.Fatal("internal detail leaked")
		}
	})

	t.Run("withdrawals unordered", func(t *testing.T) {
		t.Parallel()
		first := rwstore.WithdrawalRecord{
			Tx:             validRandomWalkEventTx(),
			WithdrawerAid:  23,
			WithdrawerAddr: rwTestCarol,
			TokenID:        10,
			AmountWei:      "1",
		}
		second := first
		second.Tx.EvtLogID = first.Tx.EvtLogID + 1
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			withdrawals: func(context.Context, *rwstore.EventPageCursor, int) ([]rwstore.WithdrawalRecord, bool, error) {
				return []rwstore.WithdrawalRecord{first, second}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/withdrawals"),
			http.StatusInternalServerError)
	})
}

func TestGetRandomWalkUserProfiles(t *testing.T) {
	t.Parallel()

	t.Run("unindexed wallet gets the zero shape", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			addressID: func(context.Context, string) (int64, error) {
				return 0, store.ErrNotFound
			},
		})
		response := serve(t, server, "/api/v2/randomwalk/users/"+randomWalkCursorAlice)
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var profile RandomWalkUserProfile
		decodeResponse(t, response, &profile)
		if profile.Address != randomWalkCursorAlice || profile.TradingVolumeWei != "0" ||
			profile.ProfitWei != "0" || profile.OwnedTokenCount != 0 {
			t.Fatalf("profile = %+v", profile)
		}
	})

	t.Run("profile for another address is internal", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			profile: func(context.Context, int64) (rwstore.UserProfileRecord, error) {
				return rwstore.UserProfileRecord{
					Address:          rwTestDave,
					TradingVolumeWei: "0",
					ProfitWei:        "0",
				}, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/users/"+randomWalkCursorAlice),
			http.StatusInternalServerError)
	})

	t.Run("malformed address", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/users/0x123"), http.StatusBadRequest)
	})
}

func TestListRandomWalkUserTokens(t *testing.T) {
	t.Parallel()

	t.Run("pagination and wallet-scoped cursor", func(t *testing.T) {
		t.Parallel()
		record := rwstore.OwnedTokenRecord{
			TokenID:          10,
			LastPriceWei:     "1",
			TradingVolumeWei: "0",
		}
		var gotAfter *rwstore.TokenPageCursor
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			userTokens: func(_ context.Context, userAid int64, after *rwstore.TokenPageCursor, _ int) ([]rwstore.OwnedTokenRecord, bool, error) {
				if userAid != 1 {
					t.Errorf("user aid = %d", userAid)
				}
				gotAfter = after
				if after != nil {
					return []rwstore.OwnedTokenRecord{}, false, nil
				}
				return []rwstore.OwnedTokenRecord{record}, true, nil
			},
		})
		target := "/api/v2/randomwalk/users/" + randomWalkCursorAlice + "/tokens?limit=1"
		response := serve(t, server, target)
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var page RandomWalkOwnedTokenPage
		decodeResponse(t, response, &page)
		if page.Meta.NextCursor == nil {
			t.Fatal("no continuation cursor")
		}
		continued := serve(t, server, target+"&cursor="+*page.Meta.NextCursor)
		if continued.Code != http.StatusOK || gotAfter == nil || gotAfter.TokenID != 10 {
			t.Fatalf("continuation = %d %+v", continued.Code, gotAfter)
		}
		// Another wallet cannot reuse the cursor.
		other := serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorBob+"/tokens?cursor="+*page.Meta.NextCursor)
		assertProblem(t, other, http.StatusBadRequest)
	})

	t.Run("unindexed wallet gets an empty page", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			addressID: func(context.Context, string) (int64, error) {
				return 0, store.ErrNotFound
			},
		})
		response := serve(t, server, "/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/tokens")
		if response.Code != http.StatusOK ||
			!strings.Contains(response.Body.String(), `"data":[]`) {
			t.Fatalf("page = %d %s", response.Code, response.Body.String())
		}
	})

	t.Run("descending rows are rejected", func(t *testing.T) {
		t.Parallel()
		first := rwstore.OwnedTokenRecord{TokenID: 11, LastPriceWei: "1", TradingVolumeWei: "0"}
		second := rwstore.OwnedTokenRecord{TokenID: 10, LastPriceWei: "1", TradingVolumeWei: "0"}
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			userTokens: func(context.Context, int64, *rwstore.TokenPageCursor, int) ([]rwstore.OwnedTokenRecord, bool, error) {
				return []rwstore.OwnedTokenRecord{first, second}, false, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/tokens"),
			http.StatusInternalServerError)
	})
}

func TestListRandomWalkUserOffers(t *testing.T) {
	t.Parallel()

	t.Run("wallet scope is enforced on rows", func(t *testing.T) {
		t.Parallel()
		foreign := validRandomWalkOfferHistoryRecord() // maker dave
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			userOffers: func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.OfferHistoryRecord, bool, error) {
				return []rwstore.OfferHistoryRecord{foreign}, false, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/offers"),
			http.StatusInternalServerError)
	})

	t.Run("maker rows map", func(t *testing.T) {
		t.Parallel()
		record := validRandomWalkOfferHistoryRecord()
		record.MakerAddr = randomWalkCursorAlice
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			userOffers: func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.OfferHistoryRecord, bool, error) {
				return []rwstore.OfferHistoryRecord{record}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/offers")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
	})

	t.Run("unindexed wallet still validates cursors", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			addressID: func(context.Context, string) (int64, error) {
				return 0, store.ErrNotFound
			},
		})
		empty := serve(t, server, "/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/offers")
		if empty.Code != http.StatusOK || !strings.Contains(empty.Body.String(), `"data":[]`) {
			t.Fatalf("page = %d %s", empty.Code, empty.Body.String())
		}
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/offers?cursor=bogus"),
			http.StatusBadRequest)
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/offers?limit=0"),
			http.StatusBadRequest)
	})
}

func TestGetRandomWalkFloorPriceResponses(t *testing.T) {
	t.Parallel()

	t.Run("empty book", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/marketplace/floor-price")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var floor RandomWalkFloorPrice
		decodeResponse(t, response, &floor)
		if floor.ActiveSellOfferCount != 0 || floor.Floor != nil {
			t.Fatalf("floor = %+v", floor)
		}
	})

	t.Run("inconsistent snapshot is internal", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			floorPrice: func(context.Context) (rwstore.FloorPriceRecord, error) {
				return rwstore.FloorPriceRecord{ActiveSellOfferCount: 2}, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/marketplace/floor-price"),
			http.StatusInternalServerError)
	})
}

func TestRandomWalkSeriesWindows(t *testing.T) {
	t.Parallel()

	t.Run("volume happy path", func(t *testing.T) {
		t.Parallel()
		var gotFrom, gotTo, gotInterval int
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			volumeSeries: func(_ context.Context, from, to, interval int) (string, []rwstore.VolumeBucketRecord, error) {
				gotFrom, gotTo, gotInterval = from, to, interval
				return "5", []rwstore.VolumeBucketRecord{
					{BucketStart: 100, TradeCount: 1, VolumeWei: "7"},
				}, nil
			},
		})
		response := serve(t, server,
			"/api/v2/randomwalk/statistics/trading-volume?from=100&to=200&intervalSeconds=100")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		if gotFrom != 100 || gotTo != 200 || gotInterval != 100 {
			t.Fatalf("window = (%d,%d,%d)", gotFrom, gotTo, gotInterval)
		}
		var volume RandomWalkTradingVolume
		decodeResponse(t, response, &volume)
		if volume.BaseVolumeWei != "5" || len(volume.Buckets) != 1 ||
			volume.Buckets[0].CumulativeVolumeWei != "12" {
			t.Fatalf("volume = %+v", volume)
		}
	})

	t.Run("window validation", func(t *testing.T) {
		t.Parallel()
		for name, target := range map[string]string{
			"inverted":         "/api/v2/randomwalk/statistics/trading-volume?from=200&to=100",
			"equal":            "/api/v2/randomwalk/statistics/listing-floor-history?from=100&to=100",
			"zero interval":    "/api/v2/randomwalk/statistics/trading-volume?from=0&to=100&intervalSeconds=0",
			"too many buckets": "/api/v2/randomwalk/statistics/trading-volume?from=0&to=100000000&intervalSeconds=1",
			"negative from":    "/api/v2/randomwalk/statistics/listing-floor-history?from=-1&to=100",
		} {
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}), target),
					http.StatusBadRequest)
			})
		}
	})

	t.Run("out-of-window bucket is internal", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			floorSeries: func(context.Context, int, int, int) ([]rwstore.FloorPointRecord, error) {
				return []rwstore.FloorPointRecord{{BucketStart: 999999, FloorWei: "1"}}, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/statistics/listing-floor-history?from=100&to=200"),
			http.StatusInternalServerError)
	})
}

func TestGetRandomWalkStatisticsAndMintReport(t *testing.T) {
	t.Parallel()

	t.Run("statistics", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			statistics: func(context.Context) (rwstore.StatisticsRecord, error) {
				return validRandomWalkStatisticsRecord(), nil
			},
		})
		response := serve(t, server, "/api/v2/randomwalk/statistics")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var statistics RandomWalkStatistics
		decodeResponse(t, response, &statistics)
		if statistics.Tokens.MintedCount != 4 || statistics.Withdrawals.Latest == nil {
			t.Fatalf("statistics = %+v", statistics)
		}
	})

	t.Run("mint report cumulative totals", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			mintReport: func(context.Context) ([]rwstore.MonthlyMintRecord, error) {
				return []rwstore.MonthlyMintRecord{
					{Year: 2021, Month: 11, MintCount: 3, MintedWei: "30"},
					{Year: 2021, Month: 12, MintCount: 1, MintedWei: "12"},
				}, nil
			},
		})
		response := serve(t, server, "/api/v2/randomwalk/statistics/mint-report")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var report RandomWalkMintReport
		decodeResponse(t, response, &report)
		if len(report.Months) != 2 || report.Months[1].CumulativeMintedWei != "42" {
			t.Fatalf("report = %+v", report)
		}
	})

	t.Run("unordered mint report is internal", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			mintReport: func(context.Context) ([]rwstore.MonthlyMintRecord, error) {
				return []rwstore.MonthlyMintRecord{
					{Year: 2022, Month: 1, MintCount: 1, MintedWei: "1"},
					{Year: 2021, Month: 12, MintCount: 1, MintedWei: "1"},
				}, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/statistics/mint-report"),
			http.StatusInternalServerError)
	})
}

func TestRandomWalkOfferBookOrderedGuards(t *testing.T) {
	t.Parallel()
	record := rwstore.OfferRecord{ListTx: validRandomWalkEventTx()}
	record.ListTx.EvtLogID = 0
	if randomWalkOfferBookOrdered(Newest, nil, 10, record, nil) {
		t.Fatal("zero event-log offer accepted")
	}
	record.ListTx.EvtLogID = 5
	if randomWalkOfferBookOrdered(RandomWalkOfferSort("bogus"), nil, 10, record, nil) {
		t.Fatal("unknown sort accepted")
	}
}

func TestListRandomWalkMarketplaceOffersInputAndRepositoryGuards(t *testing.T) {
	t.Parallel()

	validOffer := func(evtlog int64) rwstore.OfferRecord {
		record := rwstore.OfferRecord{
			ListTx:    validRandomWalkEventTx(),
			OfferID:   2,
			OfferType: 1,
			TokenID:   11,
			PriceWei:  "5",
			MakerAid:  24,
			MakerAddr: rwTestDave,
		}
		record.ListTx.EvtLogID = evtlog
		return record
	}

	t.Run("invalid limit", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/marketplace/offers?limit=0"), http.StatusBadRequest)
	})
	t.Run("malformed cursor", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/marketplace/offers?cursor=bogus"), http.StatusBadRequest)
	})
	t.Run("over-cardinality page", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			activeOffers: func(context.Context, rwstore.OfferSort, *rwstore.OfferPageCursor, int) ([]rwstore.OfferRecord, bool, error) {
				return []rwstore.OfferRecord{validOffer(3), validOffer(2)}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/marketplace/offers?limit=1"),
			http.StatusInternalServerError)
	})
	t.Run("mapper failure", func(t *testing.T) {
		t.Parallel()
		broken := validOffer(3)
		broken.PriceWei = "x"
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			activeOffers: func(context.Context, rwstore.OfferSort, *rwstore.OfferPageCursor, int) ([]rwstore.OfferRecord, bool, error) {
				return []rwstore.OfferRecord{broken}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/marketplace/offers"),
			http.StatusInternalServerError)
	})
	t.Run("has more without rows", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			activeOffers: func(context.Context, rwstore.OfferSort, *rwstore.OfferPageCursor, int) ([]rwstore.OfferRecord, bool, error) {
				return []rwstore.OfferRecord{}, true, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/marketplace/offers"),
			http.StatusInternalServerError)
	})
}

func TestRandomWalkLedgerRepositoryGuards(t *testing.T) {
	t.Parallel()

	t.Run("invalid limit", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/withdrawals?limit=201"), http.StatusBadRequest)
	})
	t.Run("over-cardinality page", func(t *testing.T) {
		t.Parallel()
		first := validRandomWalkOfferHistoryRecord()
		second := validRandomWalkOfferHistoryRecord()
		second.ListTx.EvtLogID = first.ListTx.EvtLogID - 1
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			offerHistory: func(context.Context, *rwstore.EventPageCursor, int) ([]rwstore.OfferHistoryRecord, bool, error) {
				return []rwstore.OfferHistoryRecord{first, second}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/marketplace/offer-history?limit=1"),
			http.StatusInternalServerError)
	})
	t.Run("has more without rows", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			trades: func(context.Context, *rwstore.EventPageCursor, int) ([]rwstore.TradeRecord, bool, error) {
				return []rwstore.TradeRecord{}, true, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/marketplace/trades"),
			http.StatusInternalServerError)
	})
	t.Run("mapper failure", func(t *testing.T) {
		t.Parallel()
		record := rwstore.WithdrawalRecord{
			Tx:             validRandomWalkEventTx(),
			WithdrawerAid:  23,
			WithdrawerAddr: rwTestCarol,
			TokenID:        10,
			AmountWei:      "x",
		}
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			withdrawals: func(context.Context, *rwstore.EventPageCursor, int) ([]rwstore.WithdrawalRecord, bool, error) {
				return []rwstore.WithdrawalRecord{record}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/withdrawals"),
			http.StatusInternalServerError)
	})
}

func TestListRandomWalkTokensRepositoryGuards(t *testing.T) {
	t.Parallel()

	t.Run("window upper bound", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/tokens?mintedUntil=999999999999"), http.StatusBadRequest)
	})
	t.Run("negative repository row", func(t *testing.T) {
		t.Parallel()
		record := validRandomWalkTokenRecord()
		record.TokenID = -1
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokens: func(context.Context, rwstore.TokenFilter, rwstore.TokenSort, *rwstore.TokenPageCursor, int) ([]rwstore.TokenRecord, bool, error) {
				return []rwstore.TokenRecord{record}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens"),
			http.StatusInternalServerError)
	})
	t.Run("over-cardinality page", func(t *testing.T) {
		t.Parallel()
		first := validRandomWalkTokenRecord()
		second := validRandomWalkTokenRecord()
		second.TokenID = 11
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokens: func(context.Context, rwstore.TokenFilter, rwstore.TokenSort, *rwstore.TokenPageCursor, int) ([]rwstore.TokenRecord, bool, error) {
				return []rwstore.TokenRecord{first, second}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens?limit=1"),
			http.StatusInternalServerError)
	})
	t.Run("mapper failure", func(t *testing.T) {
		t.Parallel()
		record := validRandomWalkTokenRecord()
		record.Seed = ""
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokens: func(context.Context, rwstore.TokenFilter, rwstore.TokenSort, *rwstore.TokenPageCursor, int) ([]rwstore.TokenRecord, bool, error) {
				return []rwstore.TokenRecord{record}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens"),
			http.StatusInternalServerError)
	})
	t.Run("detail mapper failure", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokenDetail: func(context.Context, int64) (rwstore.TokenDetailRecord, error) {
				return rwstore.TokenDetailRecord{
					TokenRecord:    validRandomWalkTokenRecord(),
					NameChangeText: "bogus",
				}, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens/10"),
			http.StatusInternalServerError)
	})
}

func TestRandomWalkTokenScopedRepositoryGuards(t *testing.T) {
	t.Parallel()

	t.Run("negative token id", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/tokens/-1/events"), http.StatusBadRequest)
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/tokens/-1/name-history"), http.StatusBadRequest)
	})
	t.Run("invalid limit", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/tokens/10/events?limit=0"), http.StatusBadRequest)
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/tokens/10/name-history?limit=0"), http.StatusBadRequest)
	})
	t.Run("fetch failure", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokenEvents: func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.TokenEventRecord, bool, error) {
				return nil, false, errors.New("page failure")
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens/10/events"),
			http.StatusInternalServerError)
	})
	t.Run("over-cardinality page", func(t *testing.T) {
		t.Parallel()
		first := validRandomWalkTokenEventRecord(rwstore.TokenEventMint)
		second := validRandomWalkTokenEventRecord(rwstore.TokenEventMint)
		second.Tx.EvtLogID = first.Tx.EvtLogID - 1
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokenEvents: func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.TokenEventRecord, bool, error) {
				return []rwstore.TokenEventRecord{first, second}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens/10/events?limit=1"),
			http.StatusInternalServerError)
	})
	t.Run("unordered rows", func(t *testing.T) {
		t.Parallel()
		first := validRandomWalkTokenEventRecord(rwstore.TokenEventMint)
		second := validRandomWalkTokenEventRecord(rwstore.TokenEventNameChange)
		second.Tx.EvtLogID = first.Tx.EvtLogID + 1
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokenEvents: func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.TokenEventRecord, bool, error) {
				return []rwstore.TokenEventRecord{first, second}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens/10/events"),
			http.StatusInternalServerError)
	})
	t.Run("mapper failure", func(t *testing.T) {
		t.Parallel()
		record := validRandomWalkTokenEventRecord(rwstore.TokenEventMint)
		record.Seed = ""
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			tokenEvents: func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.TokenEventRecord, bool, error) {
				return []rwstore.TokenEventRecord{record}, false, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens/10/events"),
			http.StatusInternalServerError)
	})
	t.Run("has more without rows", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			nameHistory: func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.TokenNameChangeRecord, bool, error) {
				return []rwstore.TokenNameChangeRecord{}, true, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/tokens/10/name-history"),
			http.StatusInternalServerError)
	})
}

func TestRandomWalkUserRepositoryGuards(t *testing.T) {
	t.Parallel()

	t.Run("profile resolution failure", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			addressID: func(context.Context, string) (int64, error) {
				return 0, errors.New("resolver down")
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/users/"+randomWalkCursorAlice),
			http.StatusInternalServerError)
	})
	t.Run("user tokens invalid address", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/users/0x123/tokens"), http.StatusBadRequest)
	})
	t.Run("user tokens invalid limit", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/tokens?limit=0"),
			http.StatusBadRequest)
	})
	t.Run("user tokens resolution failure", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			addressID: func(context.Context, string) (int64, error) {
				return 0, errors.New("resolver down")
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/tokens"),
			http.StatusInternalServerError)
	})
	t.Run("user tokens fetch failure", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			userTokens: func(context.Context, int64, *rwstore.TokenPageCursor, int) ([]rwstore.OwnedTokenRecord, bool, error) {
				return nil, false, errors.New("page failure")
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/tokens"),
			http.StatusInternalServerError)
	})
	t.Run("user tokens over-cardinality", func(t *testing.T) {
		t.Parallel()
		first := rwstore.OwnedTokenRecord{TokenID: 10, LastPriceWei: "1", TradingVolumeWei: "0"}
		second := rwstore.OwnedTokenRecord{TokenID: 11, LastPriceWei: "1", TradingVolumeWei: "0"}
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			userTokens: func(context.Context, int64, *rwstore.TokenPageCursor, int) ([]rwstore.OwnedTokenRecord, bool, error) {
				return []rwstore.OwnedTokenRecord{first, second}, false, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/tokens?limit=1"),
			http.StatusInternalServerError)
	})
	t.Run("user tokens mapper failure", func(t *testing.T) {
		t.Parallel()
		record := rwstore.OwnedTokenRecord{TokenID: 10, LastPriceWei: "x", TradingVolumeWei: "0"}
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			userTokens: func(context.Context, int64, *rwstore.TokenPageCursor, int) ([]rwstore.OwnedTokenRecord, bool, error) {
				return []rwstore.OwnedTokenRecord{record}, false, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/tokens"),
			http.StatusInternalServerError)
	})
	t.Run("user tokens has more without rows", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			userTokens: func(context.Context, int64, *rwstore.TokenPageCursor, int) ([]rwstore.OwnedTokenRecord, bool, error) {
				return []rwstore.OwnedTokenRecord{}, true, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/tokens"),
			http.StatusInternalServerError)
	})
	t.Run("user offers invalid address", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/users/0x123/offers"), http.StatusBadRequest)
	})
	t.Run("user offers resolution failure", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			addressID: func(context.Context, string) (int64, error) {
				return 0, errors.New("resolver down")
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/offers"),
			http.StatusInternalServerError)
	})
	t.Run("user offers mapper failure", func(t *testing.T) {
		t.Parallel()
		record := validRandomWalkOfferHistoryRecord()
		record.MakerAddr = randomWalkCursorAlice
		record.PriceWei = "x"
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			userOffers: func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.OfferHistoryRecord, bool, error) {
				return []rwstore.OfferHistoryRecord{record}, false, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/offers"),
			http.StatusInternalServerError)
	})
	t.Run("user offers purchase-side rows map", func(t *testing.T) {
		t.Parallel()
		purchaseTx := validRandomWalkEventTx()
		purchaseTx.EvtLogID = 5089
		record := validRandomWalkOfferHistoryRecord()
		record.Active = false
		record.Purchase = &rwstore.OfferOutcomePurchase{
			Tx:         purchaseTx,
			BuyerAid:   21,
			BuyerAddr:  randomWalkCursorAlice,
			SellerAid:  24,
			SellerAddr: rwTestDave,
		}
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			userOffers: func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.OfferHistoryRecord, bool, error) {
				return []rwstore.OfferHistoryRecord{record}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/randomwalk/users/"+randomWalkCursorAlice+"/offers")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
	})
}

func TestRandomWalkStatisticsAndSeriesRepositoryGuards(t *testing.T) {
	t.Parallel()

	t.Run("statistics mapper failure", func(t *testing.T) {
		t.Parallel()
		record := validRandomWalkStatisticsRecord()
		record.MintedCount = -1
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			statistics: func(context.Context) (rwstore.StatisticsRecord, error) {
				return record, nil
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/statistics"),
			http.StatusInternalServerError)
	})
	t.Run("volume mapper failure", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			volumeSeries: func(context.Context, int, int, int) (string, []rwstore.VolumeBucketRecord, error) {
				return "0", []rwstore.VolumeBucketRecord{{BucketStart: 999999, VolumeWei: "0"}}, nil
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/statistics/trading-volume?from=100&to=200"),
			http.StatusInternalServerError)
	})
	t.Run("floor series fetch failure", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			floorSeries: func(context.Context, int, int, int) ([]rwstore.FloorPointRecord, error) {
				return nil, errors.New("series failure")
			},
		})
		assertProblem(t, serve(t, server,
			"/api/v2/randomwalk/statistics/listing-floor-history?from=100&to=200"),
			http.StatusInternalServerError)
	})
	t.Run("interval above the window cap", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
			"/api/v2/randomwalk/statistics/trading-volume?from=100&to=200&intervalSeconds=158112001"),
			http.StatusBadRequest)
	})
	t.Run("mint report fetch failure", func(t *testing.T) {
		t.Parallel()
		server := newRandomWalkTestServer(t, fakeRandomWalkReader{
			mintReport: func(context.Context) ([]rwstore.MonthlyMintRecord, error) {
				return nil, errors.New("report failure")
			},
		})
		assertProblem(t, serve(t, server, "/api/v2/randomwalk/statistics/mint-report"),
			http.StatusInternalServerError)
	})
	t.Run("invalid registry addresses", func(t *testing.T) {
		t.Parallel()
		badNft := newRandomWalkTestServer(t, fakeRandomWalkReader{
			addrs: func(context.Context) (rwmodel.ContractAddresses, error) {
				return rwmodel.ContractAddresses{
					RandomWalk:  "not-hex",
					MarketPlace: "0x1200000000000000000000000000000000000012",
				}, nil
			},
		})
		assertProblem(t, serve(t, badNft, "/api/v2/randomwalk/contracts/addresses"),
			http.StatusInternalServerError)
		badMarketplace := newRandomWalkTestServer(t, fakeRandomWalkReader{
			addrs: func(context.Context) (rwmodel.ContractAddresses, error) {
				return rwmodel.ContractAddresses{
					RandomWalk:  "0x8000000000000000000000000000000000000008",
					MarketPlace: "0x0000000000000000000000000000000000000000",
				}, nil
			},
		})
		assertProblem(t, serve(t, badMarketplace, "/api/v2/randomwalk/contracts/addresses"),
			http.StatusInternalServerError)
	})
}

func TestGetRandomWalkContractAddresses(t *testing.T) {
	t.Parallel()

	response := serve(t, newRandomWalkTestServer(t, fakeRandomWalkReader{}),
		"/api/v2/randomwalk/contracts/addresses")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var addresses RandomWalkContractAddresses
	decodeResponse(t, response, &addresses)
	if addresses.NftAddress != "0x8000000000000000000000000000000000000008" ||
		addresses.MarketplaceAddress != "0x1200000000000000000000000000000000000012" {
		t.Fatalf("addresses = %+v", addresses)
	}

	failing := newRandomWalkTestServer(t, fakeRandomWalkReader{
		addrs: func(context.Context) (rwmodel.ContractAddresses, error) {
			return rwmodel.ContractAddresses{}, errors.New("registry unavailable")
		},
	})
	assertProblem(t, serve(t, failing, "/api/v2/randomwalk/contracts/addresses"),
		http.StatusInternalServerError)
}
