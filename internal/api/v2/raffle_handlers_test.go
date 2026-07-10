package v2

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestListRoundRaffleEthDepositsPaginates(t *testing.T) {
	t.Parallel()

	first := validRaffleEthDepositRecord()
	first.RoundNum, first.WinnerIndex, first.Tx.EvtLogId = 9, 0, 100
	second := validRaffleEthDepositRecord()
	second.RoundNum, second.WinnerIndex, second.Tx.EvtLogId = 9, 1, 101
	var gotRound int64
	var gotAfter *cgstore.RaffleEthDepositPageCursor
	var gotLimit int
	server := newRaffleTestServer(t, fakeRoundRaffleReader{
		eth: func(_ context.Context, round int64, after *cgstore.RaffleEthDepositPageCursor, limit int) ([]cgstore.RaffleEthDepositRecord, bool, error) {
			gotRound, gotAfter, gotLimit = round, after, limit
			return []cgstore.RaffleEthDepositRecord{first, second}, true, nil
		},
	})

	response := serve(t, server, "/api/v2/cosmicgame/rounds/9/raffle-eth-deposits?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotRound != 9 || gotAfter != nil || gotLimit != 2 {
		t.Fatalf("repository args = (%d,%+v,%d)", gotRound, gotAfter, gotLimit)
	}
	var page RoundRaffleEthDepositPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.Limit != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeRaffleEthDepositCursor(*page.Meta.NextCursor, 9)
	if err != nil {
		t.Fatal(err)
	}
	if cursor.WinnerIndex != 1 || cursor.EventLogID != 101 {
		t.Fatalf("next cursor = %+v", cursor)
	}
}

func TestListRoundRaffleNftWinnersSelectsPoolAndPaginates(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		pool     RaffleNftPool
		isStaker bool
	}{
		{pool: Bidder, isStaker: false},
		{pool: RandomWalkStaker, isStaker: true},
	} {
		t.Run(string(tc.pool), func(t *testing.T) {
			t.Parallel()
			first := validRaffleNftWinnerRecord(tc.isStaker)
			first.RoundNum, first.WinnerIndex, first.Tx.EvtLogId = 9, 0, 100
			second := validRaffleNftWinnerRecord(tc.isStaker)
			second.RoundNum, second.WinnerIndex, second.Tx.EvtLogId = 9, 1, 101
			var gotStaker bool
			server := newRaffleTestServer(t, fakeRoundRaffleReader{
				nft: func(_ context.Context, _ int64, isStaker bool, _ *cgstore.RaffleNFTWinnerPageCursor, limit int) ([]cgprimitives.CGRaffleNFTWinnerRec, bool, error) {
					gotStaker = isStaker
					if limit != 2 {
						t.Errorf("limit = %d, want 2", limit)
					}
					return []cgprimitives.CGRaffleNFTWinnerRec{first, second}, true, nil
				},
			})
			path := "/api/v2/cosmicgame/rounds/9/raffle-nft-winners?pool=" + string(tc.pool) + "&limit=2"
			response := serve(t, server, path)
			if response.Code != http.StatusOK {
				t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
			}
			if gotStaker != tc.isStaker {
				t.Fatalf("isStaker = %v, want %v", gotStaker, tc.isStaker)
			}
			var page RoundRaffleNftWinnerPage
			decodeResponse(t, response, &page)
			if len(page.Data) != 2 || page.Meta.NextCursor == nil {
				t.Fatalf("page = %+v", page)
			}
			cursor, err := decodeRaffleNftWinnerCursor(*page.Meta.NextCursor, 9, tc.pool)
			if err != nil {
				t.Fatal(err)
			}
			if cursor.Pool != tc.pool || cursor.WinnerIndex != 1 || cursor.EventLogID != 101 {
				t.Fatalf("next cursor = %+v", cursor)
			}
		})
	}
}

func TestRoundRaffleHandlersDecodeContinuationCursors(t *testing.T) {
	t.Parallel()

	ethCursor, err := encodeRaffleEthDepositCursor(raffleEthDepositCursor{
		Version:     raffleEthDepositCursorVersion,
		Round:       3,
		WinnerIndex: 7,
		EventLogID:  88,
	})
	if err != nil {
		t.Fatal(err)
	}
	var gotEthAfter *cgstore.RaffleEthDepositPageCursor
	ethServer := newRaffleTestServer(t, fakeRoundRaffleReader{
		eth: func(_ context.Context, _ int64, after *cgstore.RaffleEthDepositPageCursor, limit int) ([]cgstore.RaffleEthDepositRecord, bool, error) {
			gotEthAfter = after
			if limit != defaultPageLimit {
				t.Errorf("default limit = %d", limit)
			}
			return []cgstore.RaffleEthDepositRecord{}, false, nil
		},
	})
	ethResponse := serve(t, ethServer, "/api/v2/cosmicgame/rounds/3/raffle-eth-deposits?cursor="+ethCursor)
	if ethResponse.Code != http.StatusOK ||
		gotEthAfter == nil ||
		*gotEthAfter != (cgstore.RaffleEthDepositPageCursor{WinnerIndex: 7, EventLogID: 88}) {
		t.Fatalf("ETH response/after = %d/%+v", ethResponse.Code, gotEthAfter)
	}
	if !strings.Contains(ethResponse.Body.String(), `"data":[]`) {
		t.Fatalf("empty ETH page was not []: %s", ethResponse.Body.String())
	}

	nftCursor, err := encodeRaffleNftWinnerCursor(raffleNftWinnerCursor{
		Version:     raffleNftWinnerCursorVersion,
		Round:       3,
		Pool:        Bidder,
		WinnerIndex: 7,
		EventLogID:  88,
	})
	if err != nil {
		t.Fatal(err)
	}
	var gotNftAfter *cgstore.RaffleNFTWinnerPageCursor
	nftServer := newRaffleTestServer(t, fakeRoundRaffleReader{
		nft: func(_ context.Context, _ int64, _ bool, after *cgstore.RaffleNFTWinnerPageCursor, _ int) ([]cgprimitives.CGRaffleNFTWinnerRec, bool, error) {
			gotNftAfter = after
			return []cgprimitives.CGRaffleNFTWinnerRec{}, false, nil
		},
	})
	nftResponse := serve(t, nftServer,
		"/api/v2/cosmicgame/rounds/3/raffle-nft-winners?pool=bidder&cursor="+nftCursor)
	if nftResponse.Code != http.StatusOK ||
		gotNftAfter == nil ||
		*gotNftAfter != (cgstore.RaffleNFTWinnerPageCursor{WinnerIndex: 7, EventLogID: 88}) {
		t.Fatalf("NFT response/after = %d/%+v", nftResponse.Code, gotNftAfter)
	}
}

func TestRoundRaffleHandlersRejectInvalidInput(t *testing.T) {
	t.Parallel()

	crossRoundEth, err := encodeRaffleEthDepositCursor(raffleEthDepositCursor{
		Version: raffleEthDepositCursorVersion, Round: 2, WinnerIndex: 0, EventLogID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	crossPoolNft, err := encodeRaffleNftWinnerCursor(raffleNftWinnerCursor{
		Version: raffleNftWinnerCursorVersion, Round: 1, Pool: RandomWalkStaker, WinnerIndex: 0, EventLogID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	crossRoundNft, err := encodeRaffleNftWinnerCursor(raffleNftWinnerCursor{
		Version: raffleNftWinnerCursorVersion, Round: 2, Pool: Bidder, WinnerIndex: 0, EventLogID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]string{
		"eth negative round":   "/api/v2/cosmicgame/rounds/-1/raffle-eth-deposits",
		"eth zero limit":       "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits?limit=0",
		"eth invalid limit":    "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits?limit=201",
		"eth malformed cursor": "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits?cursor=bad",
		"eth oversized cursor": "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits?cursor=" + strings.Repeat("a", maxCursorLength+1),
		"eth cross round":      "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits?cursor=" + crossRoundEth,
		"nft missing pool":     "/api/v2/cosmicgame/rounds/1/raffle-nft-winners",
		"nft invalid pool":     "/api/v2/cosmicgame/rounds/1/raffle-nft-winners?pool=other",
		"nft invalid limit":    "/api/v2/cosmicgame/rounds/1/raffle-nft-winners?pool=bidder&limit=0",
		"nft malformed cursor": "/api/v2/cosmicgame/rounds/1/raffle-nft-winners?pool=bidder&cursor=bad",
		"nft cross pool":       "/api/v2/cosmicgame/rounds/1/raffle-nft-winners?pool=bidder&cursor=" + crossPoolNft,
		"nft cross round":      "/api/v2/cosmicgame/rounds/1/raffle-nft-winners?pool=bidder&cursor=" + crossRoundNft,
	}
	for name, path := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newRaffleTestServer(t, fakeRoundRaffleReader{}), path)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestRoundRaffleHandlersReturnNotFoundBeforePage(t *testing.T) {
	t.Parallel()

	pageCalled := false
	server := newRaffleTestServer(t, fakeRoundRaffleReader{
		exists: func(context.Context, int64) (bool, error) { return false, nil },
		eth: func(context.Context, int64, *cgstore.RaffleEthDepositPageCursor, int) ([]cgstore.RaffleEthDepositRecord, bool, error) {
			pageCalled = true
			return nil, false, nil
		},
		nft: func(context.Context, int64, bool, *cgstore.RaffleNFTWinnerPageCursor, int) ([]cgprimitives.CGRaffleNFTWinnerRec, bool, error) {
			pageCalled = true
			return nil, false, nil
		},
	})
	paths := []string{
		"/api/v2/cosmicgame/rounds/3/raffle-eth-deposits",
		"/api/v2/cosmicgame/rounds/999/raffle-nft-winners?pool=bidder",
	}
	for _, path := range paths {
		response := serve(t, server, path)
		assertProblem(t, response, http.StatusNotFound)
	}
	if pageCalled {
		t.Fatal("raffle page queried for an uncompleted round")
	}
}

func TestRoundRaffleHandlersHideInternalFailures(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		path   string
		reader fakeRoundRaffleReader
	}{
		"existence": {
			path: "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits",
			reader: fakeRoundRaffleReader{exists: func(context.Context, int64) (bool, error) {
				return false, errors.New("password=existence-secret")
			}},
		},
		"ETH page": {
			path: "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits",
			reader: fakeRoundRaffleReader{eth: func(context.Context, int64, *cgstore.RaffleEthDepositPageCursor, int) ([]cgstore.RaffleEthDepositRecord, bool, error) {
				return nil, false, errors.New("password=eth-secret")
			}},
		},
		"NFT page": {
			path: "/api/v2/cosmicgame/rounds/1/raffle-nft-winners?pool=bidder",
			reader: fakeRoundRaffleReader{nft: func(context.Context, int64, bool, *cgstore.RaffleNFTWinnerPageCursor, int) ([]cgprimitives.CGRaffleNFTWinnerRec, bool, error) {
				return nil, false, errors.New("password=nft-secret")
			}},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newRaffleTestServer(t, tc.reader), tc.path)
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "password") ||
				strings.Contains(response.Body.String(), "secret") {
				t.Fatalf("500 leaked internal detail: %s", response.Body.String())
			}
		})
	}
}

func TestRoundRaffleHandlersRejectInconsistentPages(t *testing.T) {
	t.Parallel()

	ethCases := map[string]func() ([]cgstore.RaffleEthDepositRecord, bool){
		"has more without row": func() ([]cgstore.RaffleEthDepositRecord, bool) {
			return []cgstore.RaffleEthDepositRecord{}, true
		},
		"foreign round": func() ([]cgstore.RaffleEthDepositRecord, bool) {
			record := validRaffleEthDepositRecord()
			record.RoundNum = 2
			return []cgstore.RaffleEthDepositRecord{record}, false
		},
		"duplicate key": func() ([]cgstore.RaffleEthDepositRecord, bool) {
			record := validRaffleEthDepositRecord()
			record.RoundNum, record.WinnerIndex, record.Tx.EvtLogId = 1, 0, 1
			return []cgstore.RaffleEthDepositRecord{record, record}, false
		},
		"out of order": func() ([]cgstore.RaffleEthDepositRecord, bool) {
			first := validRaffleEthDepositRecord()
			first.RoundNum, first.WinnerIndex, first.Tx.EvtLogId = 1, 1, 2
			second := validRaffleEthDepositRecord()
			second.RoundNum, second.WinnerIndex, second.Tx.EvtLogId = 1, 0, 1
			return []cgstore.RaffleEthDepositRecord{first, second}, false
		},
	}
	for name, records := range ethCases {
		t.Run("eth/"+name, func(t *testing.T) {
			t.Parallel()
			server := newRaffleTestServer(t, fakeRoundRaffleReader{
				eth: func(context.Context, int64, *cgstore.RaffleEthDepositPageCursor, int) ([]cgstore.RaffleEthDepositRecord, bool, error) {
					data, more := records()
					return data, more, nil
				},
			})
			assertProblem(t, serve(t, server, "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits"), http.StatusInternalServerError)
		})
	}

	nftCases := map[string]func() ([]cgprimitives.CGRaffleNFTWinnerRec, bool){
		"has more without row": func() ([]cgprimitives.CGRaffleNFTWinnerRec, bool) {
			return []cgprimitives.CGRaffleNFTWinnerRec{}, true
		},
		"wrong pool": func() ([]cgprimitives.CGRaffleNFTWinnerRec, bool) {
			record := validRaffleNftWinnerRecord(true)
			record.RoundNum = 1
			return []cgprimitives.CGRaffleNFTWinnerRec{record}, false
		},
		"out of order": func() ([]cgprimitives.CGRaffleNFTWinnerRec, bool) {
			first := validRaffleNftWinnerRecord(false)
			first.RoundNum, first.WinnerIndex, first.Tx.EvtLogId = 1, 1, 2
			second := validRaffleNftWinnerRecord(false)
			second.RoundNum, second.WinnerIndex, second.Tx.EvtLogId = 1, 0, 1
			return []cgprimitives.CGRaffleNFTWinnerRec{first, second}, false
		},
	}
	for name, records := range nftCases {
		t.Run("nft/"+name, func(t *testing.T) {
			t.Parallel()
			server := newRaffleTestServer(t, fakeRoundRaffleReader{
				nft: func(context.Context, int64, bool, *cgstore.RaffleNFTWinnerPageCursor, int) ([]cgprimitives.CGRaffleNFTWinnerRec, bool, error) {
					data, more := records()
					return data, more, nil
				},
			})
			assertProblem(t, serve(t, server,
				"/api/v2/cosmicgame/rounds/1/raffle-nft-winners?pool=bidder"), http.StatusInternalServerError)
		})
	}
}

func TestRoundRaffleHandlersHonorCancellation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		path   string
		reader fakeRoundRaffleReader
	}{
		"ETH existence": {
			path: "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits",
			reader: fakeRoundRaffleReader{exists: func(ctx context.Context, _ int64) (bool, error) {
				return false, ctx.Err()
			}},
		},
		"NFT page": {
			path: "/api/v2/cosmicgame/rounds/1/raffle-nft-winners?pool=bidder",
			reader: fakeRoundRaffleReader{nft: func(ctx context.Context, _ int64, _ bool, _ *cgstore.RaffleNFTWinnerPageCursor, _ int) ([]cgprimitives.CGRaffleNFTWinnerRec, bool, error) {
				return nil, false, ctx.Err()
			}},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			response := serveRaffleContext(t, newRaffleTestServer(t, tc.reader), ctx, tc.path)
			assertProblem(t, response, http.StatusInternalServerError)
		})
	}
}

func TestRoundRaffleHandlersAreDeterministic(t *testing.T) {
	t.Parallel()

	server := newRaffleTestServer(t, fakeRoundRaffleReader{
		eth: func(context.Context, int64, *cgstore.RaffleEthDepositPageCursor, int) ([]cgstore.RaffleEthDepositRecord, bool, error) {
			record := validRaffleEthDepositRecord()
			record.RoundNum = 1
			return []cgstore.RaffleEthDepositRecord{record}, false, nil
		},
	})
	first := serve(t, server, "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits")
	second := serve(t, server, "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits")
	if first.Code != second.Code || !bytes.Equal(first.Body.Bytes(), second.Body.Bytes()) {
		t.Fatalf("nondeterministic response: first=%d %s second=%d %s",
			first.Code, first.Body.String(), second.Code, second.Body.String())
	}
}

func newRaffleTestServer(t *testing.T, raffles roundRaffleReader) *Server {
	t.Helper()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server, err := newServer(
		nil,
		fakeBidReader{},
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
		raffles,
		fakeRoundDonationReader{},
		fakeStatisticsReader{},
		fakeBiddingAnalyticsReader{},
		fakeParticipantReader{},
		fakeContractState{},
		logger,
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func serveRaffleContext(
	t *testing.T,
	server *Server,
	ctx context.Context,
	target string,
) *httptest.ResponseRecorder {
	t.Helper()
	router := httpx.NewRouter()
	server.RegisterRoutes(router)
	request := httptest.NewRequest(http.MethodGet, target, nil).WithContext(ctx)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}
