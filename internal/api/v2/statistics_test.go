package v2

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"

	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestStatisticsSingletons(t *testing.T) {
	t.Parallel()
	server := newStatisticsTestServer(t, fakeStatisticsReader{
		global: func(context.Context) (cgstore.GlobalStatisticsRecord, error) {
			return validGlobalStatisticsRecord(), nil
		},
		counters: func(context.Context) (cgprimitives.CGRecordCounters, error) {
			return cgprimitives.CGRecordCounters{
				TotalBids: 12, TotalPrizes: 3, TotalDonatedNFTs: 2,
			}, nil
		},
	})
	global := serve(t, server, "/api/v2/cosmicgame/statistics")
	if global.Code != http.StatusOK {
		t.Fatalf("global = %d %s", global.Code, global.Body.String())
	}
	var stats CosmicGameGlobalStatistics
	decodeResponse(t, global, &stats)
	if stats.TotalBids != 10 || stats.TotalPrizesPaidWei != "1000" {
		t.Fatalf("statistics = %+v", stats)
	}
	counters := serve(t, server, "/api/v2/cosmicgame/statistics/counters")
	if counters.Code != http.StatusOK {
		t.Fatalf("counters = %d %s", counters.Code, counters.Body.String())
	}
	var gotCounters CosmicGameCounters
	decodeResponse(t, counters, &gotCounters)
	if gotCounters.TotalBids != 12 || gotCounters.CompletedRounds != 3 || gotCounters.DonatedNfts != 2 {
		t.Fatalf("counters = %+v", gotCounters)
	}
}

func TestStatisticsSingletonsHideErrors(t *testing.T) {
	t.Parallel()
	secret := errors.New("password=private")
	server := newStatisticsTestServer(t, fakeStatisticsReader{
		global: func(context.Context) (cgstore.GlobalStatisticsRecord, error) {
			return cgstore.GlobalStatisticsRecord{}, secret
		},
		counters: func(context.Context) (cgprimitives.CGRecordCounters, error) {
			return cgprimitives.CGRecordCounters{}, secret
		},
	})
	for _, path := range []string{
		"/api/v2/cosmicgame/statistics",
		"/api/v2/cosmicgame/statistics/counters",
	} {
		response := serve(t, server, path)
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "private") {
			t.Fatalf("internal error leaked: %s", response.Body.String())
		}
	}
}

func TestROILeaderboardPaginates(t *testing.T) {
	t.Parallel()
	first := validROILeaderboardRecord()
	first.BidderAid, first.NetProfitWei = 2, "50"
	second := validROILeaderboardRecord()
	second.BidderAid, second.NetProfitWei = 3, "40"
	var gotMinBids, gotLimit int
	var gotSort cgstore.ROILeaderboardSort
	server := newStatisticsTestServer(t, fakeStatisticsReader{
		roi: func(_ context.Context, minBids int, sort cgstore.ROILeaderboardSort, after *cgstore.ROILeaderboardPageCursor, limit int) ([]cgstore.ROILeaderboardRecord, bool, error) {
			gotMinBids, gotSort, gotLimit = minBids, sort, limit
			if after != nil {
				t.Fatalf("unexpected cursor: %+v", after)
			}
			return []cgstore.ROILeaderboardRecord{first, second}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/statistics/leaderboard/roi?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("response = %d %s", response.Code, response.Body.String())
	}
	if gotMinBids != 5 || gotSort != cgstore.ROILeaderboardNetProfit || gotLimit != 2 {
		t.Fatalf("repository args = %d,%s,%d", gotMinBids, gotSort, gotLimit)
	}
	var page RoiLeaderboardPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeROILeaderboardCursor(*page.Meta.NextCursor, NetProfit, 5)
	if err != nil || cursor.SortValue != "40" || cursor.BidderAid != 3 {
		t.Fatalf("cursor = %+v, %v", cursor, err)
	}
}

func TestROILeaderboardDecodesCursorAndSort(t *testing.T) {
	t.Parallel()
	encoded, err := encodeROILeaderboardCursor(roiLeaderboardCursor{
		Version: 1, Sort: WinRate, MinBids: 3, SortValue: "0.5",
		Secondary: 2, BidderAid: 21,
	})
	if err != nil {
		t.Fatal(err)
	}
	var got *cgstore.ROILeaderboardPageCursor
	server := newStatisticsTestServer(t, fakeStatisticsReader{
		roi: func(_ context.Context, minBids int, sort cgstore.ROILeaderboardSort, after *cgstore.ROILeaderboardPageCursor, _ int) ([]cgstore.ROILeaderboardRecord, bool, error) {
			if minBids != 3 || sort != cgstore.ROILeaderboardWinRate {
				t.Fatalf("scope = %d,%s", minBids, sort)
			}
			got = after
			return []cgstore.ROILeaderboardRecord{}, false, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/statistics/leaderboard/roi?sort=winRate&minBids=3&cursor="+encoded)
	if response.Code != http.StatusOK || got == nil ||
		got.SortValue != "0.5" || got.Secondary != 2 || got.BidderAid != 21 {
		t.Fatalf("response=%d cursor=%+v body=%s", response.Code, got, response.Body.String())
	}
	if !bytes.Contains(response.Body.Bytes(), []byte(`"data":[]`)) {
		t.Fatalf("empty data is not []: %s", response.Body.String())
	}
}

func TestROILeaderboardRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	crossSort, _ := encodeROILeaderboardCursor(roiLeaderboardCursor{
		Version: 1, Sort: Roi, MinBids: 5, SortValue: "1", BidderAid: 1,
	})
	tests := map[string]string{
		"invalid sort":     "/api/v2/cosmicgame/statistics/leaderboard/roi?sort=other",
		"negative filter":  "/api/v2/cosmicgame/statistics/leaderboard/roi?minBids=-1",
		"invalid limit":    "/api/v2/cosmicgame/statistics/leaderboard/roi?limit=201",
		"bind limit":       "/api/v2/cosmicgame/statistics/leaderboard/roi?limit=bad",
		"malformed cursor": "/api/v2/cosmicgame/statistics/leaderboard/roi?cursor=bad",
		"cross sort":       "/api/v2/cosmicgame/statistics/leaderboard/roi?sort=spent&cursor=" + crossSort,
	}
	for name, path := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assertProblem(t, serve(t, newStatisticsTestServer(t, fakeStatisticsReader{}), path), http.StatusBadRequest)
		})
	}
}

func TestROILeaderboardRejectsInconsistentPage(t *testing.T) {
	t.Parallel()
	first := validROILeaderboardRecord()
	first.BidderAid, first.NetProfitWei = 2, "40"
	second := validROILeaderboardRecord()
	second.BidderAid, second.NetProfitWei = 3, "50"
	server := newStatisticsTestServer(t, fakeStatisticsReader{
		roi: func(context.Context, int, cgstore.ROILeaderboardSort, *cgstore.ROILeaderboardPageCursor, int) ([]cgstore.ROILeaderboardRecord, bool, error) {
			return []cgstore.ROILeaderboardRecord{first, second}, false, nil
		},
	})
	assertProblem(t, serve(t, server, "/api/v2/cosmicgame/statistics/leaderboard/roi"),
		http.StatusInternalServerError)
}

func TestClaimsSummaryPaginates(t *testing.T) {
	t.Parallel()
	first := validClaimSummaryRecord(2)
	first.EventLogID = 200
	second := validClaimSummaryRecord(1)
	second.EventLogID = 100
	server := newStatisticsTestServer(t, fakeStatisticsReader{
		claims: func(context.Context, *cgstore.ClaimSummaryCursor, int) ([]cgstore.ClaimSummaryRecord, bool, error) {
			return []cgstore.ClaimSummaryRecord{first, second}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/statistics/claims?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("response = %d %s", response.Code, response.Body.String())
	}
	var page ClaimSummaryPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil ||
		page.Data[0].Round != 2 || page.Data[1].Round != 1 {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeClaimSummaryCursor(*page.Meta.NextCursor)
	if err != nil || cursor.Round != 1 || cursor.EventLogID != 100 {
		t.Fatalf("cursor = %+v, %v", cursor, err)
	}
}

func TestClaimsSummaryRejectsInvalidInputAndErrors(t *testing.T) {
	t.Parallel()
	for _, path := range []string{
		"/api/v2/cosmicgame/statistics/claims?limit=201",
		"/api/v2/cosmicgame/statistics/claims?limit=bad",
		"/api/v2/cosmicgame/statistics/claims?cursor=bad",
	} {
		assertProblem(t, serve(t, newStatisticsTestServer(t, fakeStatisticsReader{}), path),
			http.StatusBadRequest)
	}
	server := newStatisticsTestServer(t, fakeStatisticsReader{
		claims: func(context.Context, *cgstore.ClaimSummaryCursor, int) ([]cgstore.ClaimSummaryRecord, bool, error) {
			return nil, false, errors.New("secret")
		},
	})
	assertProblem(t, serve(t, server, "/api/v2/cosmicgame/statistics/claims"),
		http.StatusInternalServerError)
}

func TestRoundClaimsReturnsBoundedSections(t *testing.T) {
	t.Parallel()
	ethAmount := "10"
	tokenID := int64(7)
	transaction := validClaimTransactionRecord(cgstore.ClaimAssetETH, func(record *cgstore.ClaimTransactionRecord) {
		record.RoundNum, record.EventLogID = 0, 100
		record.EthAmountWei = &ethAmount
	})
	attached := validAttachedTokenRecord(cgstore.ClaimAssetERC721)
	attached.RoundNum, attached.EventLogID, attached.TokenID = 0, 101, &tokenID
	unclaimed := validUnclaimedItemRecord(cgstore.ClaimAssetETH, func(record *cgstore.UnclaimedItemRecord) {
		record.RoundNum, record.Segment, record.Key = 0, 0, 102
		record.EthAmountWei = &ethAmount
	})
	server := newStatisticsTestServer(t, fakeStatisticsReader{
		summary: func(context.Context, int64) (cgstore.ClaimSummaryRecord, error) {
			return validClaimSummaryRecord(0), nil
		},
		transactions: func(context.Context, int64, *cgstore.ClaimEventCursor, int) ([]cgstore.ClaimTransactionRecord, bool, error) {
			return []cgstore.ClaimTransactionRecord{transaction}, true, nil
		},
		attached: func(context.Context, int64, *cgstore.ClaimEventCursor, int) ([]cgstore.AttachedTokenRecord, bool, error) {
			return []cgstore.AttachedTokenRecord{attached}, true, nil
		},
		unclaimed: func(context.Context, int64, *cgstore.UnclaimedItemCursor, int) ([]cgstore.UnclaimedItemRecord, bool, error) {
			return []cgstore.UnclaimedItemRecord{unclaimed}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds/0/claims?limit=1")
	if response.Code != http.StatusOK {
		t.Fatalf("response = %d %s", response.Code, response.Body.String())
	}
	var detail RoundClaimsDetail
	decodeResponse(t, response, &detail)
	if len(detail.ClaimTransactions.Data) != 1 || len(detail.AttachedTokens.Data) != 1 ||
		len(detail.UnclaimedItems.Data) != 1 ||
		detail.ClaimTransactions.Meta.NextCursor == nil ||
		detail.AttachedTokens.Meta.NextCursor == nil ||
		detail.UnclaimedItems.Meta.NextCursor == nil {
		t.Fatalf("detail = %+v", detail)
	}
}

func TestRoundClaimsResponses(t *testing.T) {
	t.Parallel()
	t.Run("open or missing", func(t *testing.T) {
		t.Parallel()
		server := newStatisticsTestServer(t, fakeStatisticsReader{
			exists: func(context.Context, int64) (bool, error) { return false, nil },
		})
		assertProblem(t, serve(t, server, "/api/v2/cosmicgame/rounds/3/claims"), http.StatusNotFound)
	})
	t.Run("invalid cursor", func(t *testing.T) {
		t.Parallel()
		assertProblem(t, serve(t, newStatisticsTestServer(t, fakeStatisticsReader{}),
			"/api/v2/cosmicgame/rounds/0/claims?claimTransactionsCursor=bad"), http.StatusBadRequest)
	})
	t.Run("internal", func(t *testing.T) {
		t.Parallel()
		server := newStatisticsTestServer(t, fakeStatisticsReader{
			exists: func(context.Context, int64) (bool, error) {
				return false, errors.New("private")
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/0/claims")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "private") {
			t.Fatalf("error leaked: %s", response.Body.String())
		}
	})
}

func newStatisticsTestServer(t *testing.T, statistics statisticsReader) *Server {
	t.Helper()
	server, err := newServer(
		nil,
		fakeBidReader{},
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
		fakeRoundRaffleReader{},
		fakeRoundDonationReader{},
		statistics,
		fakeParticipantReader{},
		fakeContractState{},
		slog.New(slog.NewTextHandler(io.Discard, nil)),
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}
