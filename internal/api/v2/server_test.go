package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

type fakeBidReader struct {
	page func(context.Context, int64, cgstore.BidPageCursor, int) ([]cgprimitives.CGBidRec, bool, error)
	item func(context.Context, int64, int64) (cgprimitives.CGBidRec, error)
}

type fakeRoundReader struct {
	page func(context.Context, *cgstore.RoundPageCursor, int) ([]cgprimitives.CGRoundRec, bool, error)
	item func(context.Context, int64) (cgprimitives.CGRoundRec, error)
}

type fakeCurrentRoundReader struct {
	statistics func(context.Context, int64) (cgprimitives.CGRoundStats, error)
	bidCount   func(context.Context, int64) (int64, error)
}

type fakeRoundPrizeReader struct {
	exists func(context.Context, int64) (bool, error)
	page   func(context.Context, int64, *cgstore.PrizePageCursor, int) ([]cgprimitives.CGPrizeHistory, bool, error)
}

type fakeRoundRaffleReader struct {
	exists func(context.Context, int64) (bool, error)
	eth    func(context.Context, int64, *cgstore.RaffleEthDepositPageCursor, int) ([]cgstore.RaffleEthDepositRecord, bool, error)
	nft    func(context.Context, int64, bool, *cgstore.RaffleNFTWinnerPageCursor, int) ([]cgprimitives.CGRaffleNFTWinnerRec, bool, error)
}

type fakeRoundDonationReader struct {
	eth   func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error)
	erc20 func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundERC20DonationRecord, bool, error)
	nft   func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundNFTDonationRecord, bool, error)
}

type fakeContractState struct {
	snapshot func() contractstate.Snapshot
}

func (f fakeRoundReader) PrizeClaimsPage(ctx context.Context, after *cgstore.RoundPageCursor, limit int) ([]cgprimitives.CGRoundRec, bool, error) {
	if f.page == nil {
		return []cgprimitives.CGRoundRec{}, false, nil
	}
	return f.page(ctx, after, limit)
}

func (f fakeRoundReader) RoundInfo(ctx context.Context, round int64) (cgprimitives.CGRoundRec, error) {
	if f.item == nil {
		return cgprimitives.CGRoundRec{}, store.ErrNotFound
	}
	return f.item(ctx, round)
}

func (f fakeBidReader) BidsByRoundPage(ctx context.Context, round int64, after cgstore.BidPageCursor, limit int) ([]cgprimitives.CGBidRec, bool, error) {
	if f.page == nil {
		return []cgprimitives.CGBidRec{}, false, nil
	}
	return f.page(ctx, round, after, limit)
}

func (f fakeBidReader) BidByRoundAndPosition(ctx context.Context, round, position int64) (cgprimitives.CGBidRec, error) {
	if f.item == nil {
		return cgprimitives.CGBidRec{}, store.ErrNotFound
	}
	return f.item(ctx, round, position)
}

func (f fakeCurrentRoundReader) CosmicGameRoundStatistics(ctx context.Context, round int64) (cgprimitives.CGRoundStats, error) {
	if f.statistics == nil {
		return cgprimitives.CGRoundStats{RoundNum: round}, nil
	}
	return f.statistics(ctx, round)
}

func (f fakeCurrentRoundReader) BidCountForRound(ctx context.Context, round int64) (int64, error) {
	if f.bidCount == nil {
		return 0, nil
	}
	return f.bidCount(ctx, round)
}

func (f fakeRoundPrizeReader) CompletedRoundExists(ctx context.Context, round int64) (bool, error) {
	if f.exists == nil {
		return true, nil
	}
	return f.exists(ctx, round)
}

func (f fakeRoundPrizeReader) AllPrizesForRoundPage(
	ctx context.Context,
	round int64,
	after *cgstore.PrizePageCursor,
	limit int,
) ([]cgprimitives.CGPrizeHistory, bool, error) {
	if f.page == nil {
		return []cgprimitives.CGPrizeHistory{}, false, nil
	}
	return f.page(ctx, round, after, limit)
}

func (f fakeRoundRaffleReader) CompletedRoundExists(ctx context.Context, round int64) (bool, error) {
	if f.exists == nil {
		return true, nil
	}
	return f.exists(ctx, round)
}

func (f fakeRoundRaffleReader) RaffleEthDepositsByRoundPage(
	ctx context.Context,
	round int64,
	after *cgstore.RaffleEthDepositPageCursor,
	limit int,
) ([]cgstore.RaffleEthDepositRecord, bool, error) {
	if f.eth == nil {
		return []cgstore.RaffleEthDepositRecord{}, false, nil
	}
	return f.eth(ctx, round, after, limit)
}

func (f fakeRoundRaffleReader) RaffleNFTWinnersByRoundPage(
	ctx context.Context,
	round int64,
	isStaker bool,
	after *cgstore.RaffleNFTWinnerPageCursor,
	limit int,
) ([]cgprimitives.CGRaffleNFTWinnerRec, bool, error) {
	if f.nft == nil {
		return []cgprimitives.CGRaffleNFTWinnerRec{}, false, nil
	}
	return f.nft(ctx, round, isStaker, after, limit)
}

func (f fakeRoundDonationReader) EthDonationsByRoundPage(
	ctx context.Context,
	round int64,
	after *cgstore.DonationPageCursor,
	limit int,
) ([]cgstore.RoundEthDonationRecord, bool, error) {
	if f.eth == nil {
		return []cgstore.RoundEthDonationRecord{}, false, nil
	}
	return f.eth(ctx, round, after, limit)
}

func (f fakeRoundDonationReader) ERC20DonationsByRoundPage(
	ctx context.Context,
	round int64,
	after *cgstore.DonationPageCursor,
	limit int,
) ([]cgstore.RoundERC20DonationRecord, bool, error) {
	if f.erc20 == nil {
		return []cgstore.RoundERC20DonationRecord{}, false, nil
	}
	return f.erc20(ctx, round, after, limit)
}

func (f fakeRoundDonationReader) NFTDonationsByRoundPage(
	ctx context.Context,
	round int64,
	after *cgstore.DonationPageCursor,
	limit int,
) ([]cgstore.RoundNFTDonationRecord, bool, error) {
	if f.nft == nil {
		return []cgstore.RoundNFTDonationRecord{}, false, nil
	}
	return f.nft(ctx, round, after, limit)
}

func (f fakeContractState) Snapshot() contractstate.Snapshot {
	if f.snapshot == nil {
		return contractstate.Snapshot{}
	}
	return f.snapshot()
}

func TestListRoundBidsPaginatesWithOpaqueCursor(t *testing.T) {
	t.Parallel()

	var gotRound int64
	var gotAfter cgstore.BidPageCursor
	var gotLimit int
	first := validBidRecord()
	first.RoundNum, first.BidPosition, first.Tx.EvtLogId = 9, 1, 100
	second := validBidRecord()
	second.RoundNum, second.BidPosition, second.Tx.EvtLogId = 9, 2, 101

	server := newTestServer(t, fakeBidReader{
		page: func(_ context.Context, round int64, after cgstore.BidPageCursor, limit int) ([]cgprimitives.CGBidRec, bool, error) {
			gotRound, gotAfter, gotLimit = round, after, limit
			return []cgprimitives.CGBidRec{first, second}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds/9/bids?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotRound != 9 || gotAfter != (cgstore.BidPageCursor{}) || gotLimit != 2 {
		t.Fatalf("repository args = (%d,%+v,%d)", gotRound, gotAfter, gotLimit)
	}

	var page RoundBidPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.Limit != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeBidCursor(*page.Meta.NextCursor, 9)
	if err != nil {
		t.Fatalf("decode next cursor: %v", err)
	}
	if cursor.BidPosition != 2 || cursor.EventLogID != 101 {
		t.Fatalf("next cursor = %+v", cursor)
	}
}

func TestListRoundBidsDecodesContinuationCursor(t *testing.T) {
	t.Parallel()

	encoded, err := encodeBidCursor(bidCursor{
		Version:     bidCursorVersion,
		Round:       3,
		BidPosition: 7,
		EventLogID:  88,
	})
	if err != nil {
		t.Fatal(err)
	}

	var gotAfter cgstore.BidPageCursor
	server := newTestServer(t, fakeBidReader{
		page: func(_ context.Context, _ int64, after cgstore.BidPageCursor, limit int) ([]cgprimitives.CGBidRec, bool, error) {
			gotAfter = after
			if limit != defaultPageLimit {
				t.Errorf("default limit = %d, want %d", limit, defaultPageLimit)
			}
			return []cgprimitives.CGBidRec{}, false, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds/3/bids?cursor="+encoded)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotAfter != (cgstore.BidPageCursor{BidPosition: 7, EventLogID: 88}) {
		t.Fatalf("decoded repository cursor = %+v", gotAfter)
	}
	if !bytes.Contains(response.Body.Bytes(), []byte(`"data":[]`)) {
		t.Fatalf("empty data was not encoded as []: %s", response.Body.String())
	}
	var page RoundBidPage
	decodeResponse(t, response, &page)
	if page.Meta.NextCursor != nil {
		t.Fatalf("exhausted page has next cursor %q", *page.Meta.NextCursor)
	}
}

func TestListRoundBidsRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	crossRound, err := encodeBidCursor(bidCursor{
		Version:     bidCursorVersion,
		Round:       2,
		BidPosition: 1,
		EventLogID:  1,
	})
	if err != nil {
		t.Fatal(err)
	}

	tests := map[string]string{
		"negative round":   "/api/v2/cosmicgame/rounds/-1/bids",
		"zero limit":       "/api/v2/cosmicgame/rounds/1/bids?limit=0",
		"excessive limit":  "/api/v2/cosmicgame/rounds/1/bids?limit=201",
		"duplicate limit":  "/api/v2/cosmicgame/rounds/1/bids?limit=1&limit=2",
		"malformed cursor": "/api/v2/cosmicgame/rounds/1/bids?cursor=not-a-cursor",
		"oversized cursor": "/api/v2/cosmicgame/rounds/1/bids?cursor=" + strings.Repeat("a", maxCursorLength+1),
		"cross-round":      "/api/v2/cosmicgame/rounds/1/bids?cursor=" + crossRound,
		"bind round":       "/api/v2/cosmicgame/rounds/not-a-number/bids",
		"bind limit":       "/api/v2/cosmicgame/rounds/1/bids?limit=wat",
	}
	for name, path := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newTestServer(t, fakeBidReader{}), path)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestListRoundBidsHidesRepositoryErrors(t *testing.T) {
	t.Parallel()

	server := newTestServer(t, fakeBidReader{
		page: func(context.Context, int64, cgstore.BidPageCursor, int) ([]cgprimitives.CGBidRec, bool, error) {
			return nil, false, errors.New("password=super-secret")
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds/1/bids")
	assertProblem(t, response, http.StatusInternalServerError)
	if strings.Contains(response.Body.String(), "super-secret") {
		t.Fatalf("internal error leaked: %s", response.Body.String())
	}
}

func TestListRoundBidsRejectsInconsistentRepositoryPage(t *testing.T) {
	t.Parallel()

	t.Run("has more without row", func(t *testing.T) {
		t.Parallel()
		server := newTestServer(t, fakeBidReader{
			page: func(context.Context, int64, cgstore.BidPageCursor, int) ([]cgprimitives.CGBidRec, bool, error) {
				return []cgprimitives.CGBidRec{}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/1/bids")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("out of order", func(t *testing.T) {
		t.Parallel()
		first := validBidRecord()
		first.RoundNum, first.BidPosition, first.Tx.EvtLogId = 1, 2, 20
		second := validBidRecord()
		second.RoundNum, second.BidPosition, second.Tx.EvtLogId = 1, 1, 10
		server := newTestServer(t, fakeBidReader{
			page: func(context.Context, int64, cgstore.BidPageCursor, int) ([]cgprimitives.CGBidRec, bool, error) {
				return []cgprimitives.CGBidRec{first, second}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/1/bids")
		assertProblem(t, response, http.StatusInternalServerError)
	})
}

func TestGetRoundBidResponses(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		record := validBidRecord()
		record.RoundNum, record.BidPosition = 4, 2
		server := newTestServer(t, fakeBidReader{
			item: func(_ context.Context, round, position int64) (cgprimitives.CGBidRec, error) {
				if round != 4 || position != 2 {
					t.Fatalf("repository args = (%d,%d)", round, position)
				}
				return record, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/4/bids/2")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var bid Bid
		decodeResponse(t, response, &bid)
		if bid.Round != 4 || bid.Position != 2 {
			t.Fatalf("bid = %+v", bid)
		}
	})

	t.Run("not found", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newTestServer(t, fakeBidReader{}), "/api/v2/cosmicgame/rounds/4/bids/99")
		assertProblem(t, response, http.StatusNotFound)
	})

	t.Run("invalid position", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newTestServer(t, fakeBidReader{}), "/api/v2/cosmicgame/rounds/4/bids/0")
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("repository failure", func(t *testing.T) {
		t.Parallel()
		server := newTestServer(t, fakeBidReader{
			item: func(context.Context, int64, int64) (cgprimitives.CGBidRec, error) {
				return cgprimitives.CGBidRec{}, errors.New("private database detail")
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/4/bids/2")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "private database detail") {
			t.Fatalf("internal error leaked: %s", response.Body.String())
		}
	})
}

func TestNewServerValidatesDependencies(t *testing.T) {
	t.Parallel()

	if _, err := NewServer(nil, nil, nil); err == nil {
		t.Fatal("NewServer accepted nil dependencies")
	}
	if _, err := NewServer(&store.Store{}, nil, nil); err == nil {
		t.Fatal("NewServer accepted a nil contract state")
	}
	if _, err := newServer(nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil bid repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil round repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil current-round repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil round-prize repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil round-raffle repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil round-donation repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil contract state")
	}
	server, err := newServer(
		nil,
		fakeBidReader{},
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
		fakeRoundRaffleReader{},
		fakeRoundDonationReader{},
		fakeContractState{},
		nil,
	)
	if err != nil {
		t.Fatalf("newServer rejected test dependencies: %v", err)
	}
	if server.logger == nil {
		t.Fatal("newServer did not install a default logger")
	}
}

func newTestServer(t *testing.T, bids bidReader) *Server {
	t.Helper()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server, err := newServer(
		nil,
		bids,
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
		fakeRoundRaffleReader{},
		fakeRoundDonationReader{},
		fakeContractState{},
		logger,
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func serve(t *testing.T, server *Server, target string) *httptest.ResponseRecorder {
	t.Helper()
	router := httpx.NewRouter()
	server.RegisterRoutes(router)
	request := httptest.NewRequest(http.MethodGet, target, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}

func decodeResponse(t *testing.T, response *httptest.ResponseRecorder, target any) {
	t.Helper()
	if err := json.Unmarshal(response.Body.Bytes(), target); err != nil {
		t.Fatalf("decode response: %v\n%s", err, response.Body.String())
	}
}

func assertProblem(t *testing.T, response *httptest.ResponseRecorder, status int) {
	t.Helper()
	if response.Code != status {
		t.Fatalf("status = %d, want %d; body=%s", response.Code, status, response.Body.String())
	}
	if got := response.Header().Get("Content-Type"); got != "application/problem+json" {
		t.Fatalf("Content-Type = %q", got)
	}
	var problem Problem
	decodeResponse(t, response, &problem)
	if problem.Status != status || problem.Type == "" || problem.Title == "" {
		t.Fatalf("problem = %+v", problem)
	}
}
