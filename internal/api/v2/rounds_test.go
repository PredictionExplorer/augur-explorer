package v2

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestListRoundsPaginatesWithOpaqueCursor(t *testing.T) {
	t.Parallel()

	var gotAfter *cgstore.RoundPageCursor
	var gotLimit int
	first := validRoundRecord()
	first.RoundNum, first.ClaimPrizeTx.Tx.EvtLogId = 2, 30
	second := validRoundRecord()
	second.RoundNum, second.ClaimPrizeTx.Tx.EvtLogId = 1, 20

	server := newRoundTestServer(t, fakeRoundReader{
		page: func(_ context.Context, after *cgstore.RoundPageCursor, limit int) ([]cgmodel.CGRoundRec, bool, error) {
			gotAfter, gotLimit = after, limit
			return []cgmodel.CGRoundRec{first, second}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotAfter != nil || gotLimit != 2 {
		t.Fatalf("repository args = (%+v,%d)", gotAfter, gotLimit)
	}

	var page RoundPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.Limit != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeRoundCursor(*page.Meta.NextCursor)
	if err != nil {
		t.Fatalf("decode next cursor: %v", err)
	}
	if cursor.RoundNum != 1 || cursor.EventLogID != 20 {
		t.Fatalf("next cursor = %+v", cursor)
	}
}

func TestListRoundsDecodesContinuationCursor(t *testing.T) {
	t.Parallel()

	encoded, err := encodeRoundCursor(roundCursor{
		Version:    roundCursorVersion,
		RoundNum:   7,
		EventLogID: 88,
	})
	if err != nil {
		t.Fatal(err)
	}

	var gotAfter *cgstore.RoundPageCursor
	server := newRoundTestServer(t, fakeRoundReader{
		page: func(_ context.Context, after *cgstore.RoundPageCursor, limit int) ([]cgmodel.CGRoundRec, bool, error) {
			gotAfter = after
			if limit != defaultPageLimit {
				t.Errorf("default limit = %d, want %d", limit, defaultPageLimit)
			}
			return []cgmodel.CGRoundRec{}, false, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds?cursor="+encoded)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotAfter == nil || *gotAfter != (cgstore.RoundPageCursor{RoundNum: 7, EventLogID: 88}) {
		t.Fatalf("decoded repository cursor = %+v", gotAfter)
	}
	if !strings.Contains(response.Body.String(), `"data":[]`) {
		t.Fatalf("empty data was not encoded as []: %s", response.Body.String())
	}
}

func TestListRoundsRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"zero limit":       "/api/v2/cosmicgame/rounds?limit=0",
		"excessive limit":  "/api/v2/cosmicgame/rounds?limit=201",
		"duplicate limit":  "/api/v2/cosmicgame/rounds?limit=1&limit=2",
		"malformed cursor": "/api/v2/cosmicgame/rounds?cursor=not-a-cursor",
		"oversized cursor": "/api/v2/cosmicgame/rounds?cursor=" + strings.Repeat("a", maxCursorLength+1),
		"bind limit":       "/api/v2/cosmicgame/rounds?limit=wat",
	}
	for name, path := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newRoundTestServer(t, fakeRoundReader{}), path)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestListRoundsHidesRepositoryErrors(t *testing.T) {
	t.Parallel()

	server := newRoundTestServer(t, fakeRoundReader{
		page: func(context.Context, *cgstore.RoundPageCursor, int) ([]cgmodel.CGRoundRec, bool, error) {
			return nil, false, errors.New("password=super-secret")
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds")
	assertProblem(t, response, http.StatusInternalServerError)
	if strings.Contains(response.Body.String(), "super-secret") {
		t.Fatalf("internal error leaked: %s", response.Body.String())
	}
}

func TestListRoundsRejectsInconsistentRepositoryPage(t *testing.T) {
	t.Parallel()

	t.Run("has more without row", func(t *testing.T) {
		t.Parallel()
		server := newRoundTestServer(t, fakeRoundReader{
			page: func(context.Context, *cgstore.RoundPageCursor, int) ([]cgmodel.CGRoundRec, bool, error) {
				return []cgmodel.CGRoundRec{}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("out of order", func(t *testing.T) {
		t.Parallel()
		first := validRoundRecord()
		first.RoundNum, first.ClaimPrizeTx.Tx.EvtLogId = 1, 10
		second := validRoundRecord()
		second.RoundNum, second.ClaimPrizeTx.Tx.EvtLogId = 2, 20
		server := newRoundTestServer(t, fakeRoundReader{
			page: func(context.Context, *cgstore.RoundPageCursor, int) ([]cgmodel.CGRoundRec, bool, error) {
				return []cgmodel.CGRoundRec{first, second}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds")
		assertProblem(t, response, http.StatusInternalServerError)
	})
}

func TestGetRoundResponses(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		record := validRoundRecord()
		record.RoundNum = 4
		server := newRoundTestServer(t, fakeRoundReader{
			item: func(_ context.Context, round int64) (cgmodel.CGRoundRec, error) {
				if round != 4 {
					t.Fatalf("repository round = %d", round)
				}
				return record, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/4")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var round CosmicGameRound
		decodeResponse(t, response, &round)
		if round.Round != 4 || round.Status != Completed {
			t.Fatalf("round = %+v", round)
		}
	})

	t.Run("not found", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newRoundTestServer(t, fakeRoundReader{}), "/api/v2/cosmicgame/rounds/999")
		assertProblem(t, response, http.StatusNotFound)
	})

	t.Run("invalid round", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newRoundTestServer(t, fakeRoundReader{}), "/api/v2/cosmicgame/rounds/-1")
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("bind round", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newRoundTestServer(t, fakeRoundReader{}), "/api/v2/cosmicgame/rounds/not-a-number")
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("repository failure", func(t *testing.T) {
		t.Parallel()
		server := newRoundTestServer(t, fakeRoundReader{
			item: func(context.Context, int64) (cgmodel.CGRoundRec, error) {
				return cgmodel.CGRoundRec{}, errors.New("private database detail")
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/4")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "private database detail") {
			t.Fatalf("internal error leaked: %s", response.Body.String())
		}
	})

	t.Run("identity mismatch", func(t *testing.T) {
		t.Parallel()
		record := validRoundRecord()
		record.RoundNum = 5
		server := newRoundTestServer(t, fakeRoundReader{
			item: func(context.Context, int64) (cgmodel.CGRoundRec, error) {
				return record, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/4")
		assertProblem(t, response, http.StatusInternalServerError)
	})
}

func newRoundTestServer(t *testing.T, rounds roundReader) *Server {
	t.Helper()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server, err := newServer(
		nil,
		fakeBidReader{},
		rounds,
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
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
