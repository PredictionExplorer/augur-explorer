package v2

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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

func TestListRoundPrizesPaginatesWithOpaqueCursor(t *testing.T) {
	t.Parallel()

	first := validRoundPrizeRecord(0)
	first.RoundNum, first.WinnerIndex = 9, 0
	second := validRoundPrizeRecord(1)
	second.RoundNum, second.WinnerIndex = 9, 0
	var gotExistsRound, gotPageRound int64
	var gotAfter *cgstore.PrizePageCursor
	var gotLimit int
	server := newPrizeTestServer(t, fakeRoundPrizeReader{
		exists: func(_ context.Context, round int64) (bool, error) {
			gotExistsRound = round
			return true, nil
		},
		page: func(_ context.Context, round int64, after *cgstore.PrizePageCursor, limit int) ([]cgprimitives.CGPrizeHistory, bool, error) {
			gotPageRound, gotAfter, gotLimit = round, after, limit
			return []cgprimitives.CGPrizeHistory{first, second}, true, nil
		},
	})

	response := serve(t, server, "/api/v2/cosmicgame/rounds/9/prizes?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotExistsRound != 9 || gotPageRound != 9 || gotAfter != nil || gotLimit != 2 {
		t.Fatalf("repository args = exists:%d page:(%d,%+v,%d)",
			gotExistsRound, gotPageRound, gotAfter, gotLimit)
	}

	var page RoundPrizePage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.Limit != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodePrizeCursor(*page.Meta.NextCursor, 9)
	if err != nil {
		t.Fatalf("decode next cursor: %v", err)
	}
	if cursor.PrizeType != 1 || cursor.WinnerIndex != 0 {
		t.Fatalf("next cursor = %+v", cursor)
	}
}

func TestListRoundPrizesDecodesContinuationCursor(t *testing.T) {
	t.Parallel()

	encoded, err := encodePrizeCursor(prizeCursor{
		Version:     prizeCursorVersion,
		Round:       3,
		PrizeType:   10,
		WinnerIndex: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	var gotAfter *cgstore.PrizePageCursor
	server := newPrizeTestServer(t, fakeRoundPrizeReader{
		page: func(_ context.Context, _ int64, after *cgstore.PrizePageCursor, limit int) ([]cgprimitives.CGPrizeHistory, bool, error) {
			gotAfter = after
			if limit != defaultPageLimit {
				t.Errorf("default limit = %d, want %d", limit, defaultPageLimit)
			}
			return []cgprimitives.CGPrizeHistory{}, false, nil
		},
	})

	response := serve(t, server, "/api/v2/cosmicgame/rounds/3/prizes?cursor="+encoded)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotAfter == nil || *gotAfter != (cgstore.PrizePageCursor{PrizeType: 10, WinnerIndex: 2}) {
		t.Fatalf("decoded repository cursor = %+v", gotAfter)
	}
	if !strings.Contains(response.Body.String(), `"data":[]`) {
		t.Fatalf("empty data was not encoded as []: %s", response.Body.String())
	}
	var page RoundPrizePage
	decodeResponse(t, response, &page)
	if page.Meta.NextCursor != nil {
		t.Fatalf("exhausted page has next cursor %q", *page.Meta.NextCursor)
	}
}

func TestListRoundPrizesRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	crossRound, err := encodePrizeCursor(prizeCursor{
		Version:     prizeCursorVersion,
		Round:       2,
		PrizeType:   1,
		WinnerIndex: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]string{
		"negative round":   "/api/v2/cosmicgame/rounds/-1/prizes",
		"zero limit":       "/api/v2/cosmicgame/rounds/1/prizes?limit=0",
		"excessive limit":  "/api/v2/cosmicgame/rounds/1/prizes?limit=201",
		"duplicate limit":  "/api/v2/cosmicgame/rounds/1/prizes?limit=1&limit=2",
		"malformed cursor": "/api/v2/cosmicgame/rounds/1/prizes?cursor=not-a-cursor",
		"oversized cursor": "/api/v2/cosmicgame/rounds/1/prizes?cursor=" + strings.Repeat("a", maxCursorLength+1),
		"cross-round":      "/api/v2/cosmicgame/rounds/1/prizes?cursor=" + crossRound,
		"bind round":       "/api/v2/cosmicgame/rounds/not-a-number/prizes",
		"bind limit":       "/api/v2/cosmicgame/rounds/1/prizes?limit=wat",
	}
	for name, path := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newPrizeTestServer(t, fakeRoundPrizeReader{}), path)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestListRoundPrizesReturnsNotFoundForUncompletedRound(t *testing.T) {
	t.Parallel()

	pageCalled := false
	server := newPrizeTestServer(t, fakeRoundPrizeReader{
		exists: func(context.Context, int64) (bool, error) {
			return false, nil
		},
		page: func(context.Context, int64, *cgstore.PrizePageCursor, int) ([]cgprimitives.CGPrizeHistory, bool, error) {
			pageCalled = true
			return nil, false, nil
		},
	})
	for _, round := range []int64{3, 999} {
		response := serve(t, server, "/api/v2/cosmicgame/rounds/"+itoaV2(round)+"/prizes")
		assertProblem(t, response, http.StatusNotFound)
	}
	if pageCalled {
		t.Fatal("prize page queried for an uncompleted round")
	}
}

func TestListRoundPrizesHidesInternalFailures(t *testing.T) {
	t.Parallel()

	tests := map[string]fakeRoundPrizeReader{
		"existence failure": {
			exists: func(context.Context, int64) (bool, error) {
				return false, errors.New("password=existence-secret")
			},
		},
		"page failure": {
			page: func(context.Context, int64, *cgstore.PrizePageCursor, int) ([]cgprimitives.CGPrizeHistory, bool, error) {
				return nil, false, errors.New("password=page-secret")
			},
		},
		"mapping failure": {
			page: func(context.Context, int64, *cgstore.PrizePageCursor, int) ([]cgprimitives.CGPrizeHistory, bool, error) {
				record := validRoundPrizeRecord(0)
				record.RoundNum = 1
				record.Amount = "password=mapping-secret"
				return []cgprimitives.CGPrizeHistory{record}, false, nil
			},
		},
	}
	for name, reader := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newPrizeTestServer(t, reader), "/api/v2/cosmicgame/rounds/1/prizes")
			assertProblem(t, response, http.StatusInternalServerError)
			for _, secret := range []string{"existence-secret", "page-secret", "mapping-secret", "password"} {
				if strings.Contains(response.Body.String(), secret) {
					t.Fatalf("500 leaked %q: %s", secret, response.Body.String())
				}
			}
		})
	}
}

func TestListRoundPrizesRejectsInconsistentPage(t *testing.T) {
	t.Parallel()

	tests := map[string]func() ([]cgprimitives.CGPrizeHistory, bool){
		"has more without row": func() ([]cgprimitives.CGPrizeHistory, bool) {
			return []cgprimitives.CGPrizeHistory{}, true
		},
		"foreign round": func() ([]cgprimitives.CGPrizeHistory, bool) {
			record := validRoundPrizeRecord(0)
			record.RoundNum = 2
			return []cgprimitives.CGPrizeHistory{record}, false
		},
		"duplicate key": func() ([]cgprimitives.CGPrizeHistory, bool) {
			first := validRoundPrizeRecord(0)
			first.RoundNum, first.WinnerIndex = 1, 0
			second := first
			return []cgprimitives.CGPrizeHistory{first, second}, false
		},
		"out of order": func() ([]cgprimitives.CGPrizeHistory, bool) {
			first := validRoundPrizeRecord(2)
			first.RoundNum, first.WinnerIndex = 1, 0
			second := validRoundPrizeRecord(1)
			second.RoundNum, second.WinnerIndex = 1, 0
			return []cgprimitives.CGPrizeHistory{first, second}, false
		},
	}
	for name, records := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			server := newPrizeTestServer(t, fakeRoundPrizeReader{
				page: func(context.Context, int64, *cgstore.PrizePageCursor, int) ([]cgprimitives.CGPrizeHistory, bool, error) {
					data, more := records()
					return data, more, nil
				},
			})
			response := serve(t, server, "/api/v2/cosmicgame/rounds/1/prizes")
			assertProblem(t, response, http.StatusInternalServerError)
		})
	}
}

func TestListRoundPrizesHonorsCancelledContext(t *testing.T) {
	t.Parallel()

	tests := map[string]fakeRoundPrizeReader{
		"existence": {
			exists: func(ctx context.Context, _ int64) (bool, error) {
				return false, ctx.Err()
			},
		},
		"page": {
			page: func(ctx context.Context, _ int64, _ *cgstore.PrizePageCursor, _ int) ([]cgprimitives.CGPrizeHistory, bool, error) {
				return nil, false, ctx.Err()
			},
		},
	}
	for name, reader := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			response := servePrizeContext(t, newPrizeTestServer(t, reader), ctx, "/api/v2/cosmicgame/rounds/1/prizes")
			assertProblem(t, response, http.StatusInternalServerError)
		})
	}
}

func TestListRoundPrizesResponseIsDeterministic(t *testing.T) {
	t.Parallel()

	server := newPrizeTestServer(t, fakeRoundPrizeReader{
		page: func(context.Context, int64, *cgstore.PrizePageCursor, int) ([]cgprimitives.CGPrizeHistory, bool, error) {
			record := validRoundPrizeRecord(0)
			record.RoundNum, record.WinnerIndex = 1, 0
			return []cgprimitives.CGPrizeHistory{record}, false, nil
		},
	})
	first := serve(t, server, "/api/v2/cosmicgame/rounds/1/prizes")
	second := serve(t, server, "/api/v2/cosmicgame/rounds/1/prizes")
	if first.Code != second.Code || !bytes.Equal(first.Body.Bytes(), second.Body.Bytes()) {
		t.Fatalf("nondeterministic response: first=%d %s second=%d %s",
			first.Code, first.Body.String(), second.Code, second.Body.String())
	}
}

func newPrizeTestServer(t *testing.T, prizes roundPrizeReader) *Server {
	t.Helper()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server, err := newServer(
		nil,
		fakeBidReader{},
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		prizes,
		fakeRoundRaffleReader{},
		fakeRoundDonationReader{},
		fakeStatisticsReader{},
		fakeBiddingAnalyticsReader{},
		fakeContractAddressReader{},
		fakeParticipantReader{},
		fakeUserReader{},
		fakeContractState{},
		logger,
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func servePrizeContext(
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

func itoaV2(value int64) string {
	return fmt.Sprintf("%d", value)
}
