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

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

func TestGetCurrentRoundUsesOneSnapshotAndAuthoritativeBidCount(t *testing.T) {
	t.Parallel()

	snapshotCalls := 0
	statisticsCalls := 0
	bidCountCalls := 0
	server := newCurrentRoundTestServer(t,
		fakeContractState{snapshot: func() contractstate.Snapshot {
			snapshotCalls++
			return validCurrentRoundSnapshot()
		}},
		fakeCurrentRoundReader{
			statistics: func(_ context.Context, round int64) (cgmodel.CGRoundStats, error) {
				statisticsCalls++
				if round != 3 {
					t.Fatalf("statistics round = %d, want 3", round)
				}
				return validCurrentRoundStats(), nil
			},
			bidCount: func(_ context.Context, round int64) (int64, error) {
				bidCountCalls++
				if round != 3 {
					t.Fatalf("bid-count round = %d, want 3", round)
				}
				return 7, nil
			},
		},
	)

	response := serve(t, server, currentRoundInstance)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if snapshotCalls != 1 || statisticsCalls != 1 || bidCountCalls != 1 {
		t.Fatalf("dependency calls = snapshot:%d statistics:%d bidCount:%d, want one each",
			snapshotCalls, statisticsCalls, bidCountCalls)
	}

	var current CosmicGameCurrentRound
	decodeResponse(t, response, &current)
	if current.Round != 3 || current.Status != Open || current.Statistics.TotalBids != 7 {
		t.Fatalf("current round = %+v", current)
	}
}

func TestGetCurrentRoundReturnsServiceUnavailableForCacheSentinel(t *testing.T) {
	t.Parallel()

	tests := map[string]func(*contractstate.Snapshot){
		"round":       func(s *contractstate.Snapshot) { s.RoundNum = -1 },
		"countdown":   func(s *contractstate.Snapshot) { s.PrizeClaimTimestamp = -1 },
		"bid price":   func(s *contractstate.Snapshot) { s.BlockPinnedBidPrice = "error" },
		"prize pool":  func(s *contractstate.Snapshot) { s.PrizeAmount = "error" },
		"time config": func(s *contractstate.Snapshot) { s.MainPrizeTimeIncrement = "error" },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			repositoryCalled := false
			snapshot := validCurrentRoundSnapshot()
			mutate(&snapshot)
			server := newCurrentRoundTestServer(t,
				fakeContractState{snapshot: func() contractstate.Snapshot { return snapshot }},
				fakeCurrentRoundReader{
					statistics: func(context.Context, int64) (cgmodel.CGRoundStats, error) {
						repositoryCalled = true
						return cgmodel.CGRoundStats{}, nil
					},
					bidCount: func(context.Context, int64) (int64, error) {
						repositoryCalled = true
						return 0, nil
					},
				},
			)

			response := serve(t, server, currentRoundInstance)
			assertProblem(t, response, http.StatusServiceUnavailable)
			if got := response.Header().Get("Retry-After"); got != "5" {
				t.Fatalf("Retry-After = %q, want 5", got)
			}
			if repositoryCalled {
				t.Fatal("repository was called for an unavailable snapshot")
			}
			if strings.Contains(response.Body.String(), "BidPrice") ||
				strings.Contains(response.Body.String(), "error sentinel") {
				t.Fatalf("503 leaked cache details: %s", response.Body.String())
			}
		})
	}
}

func TestGetCurrentRoundTreatsUninitializedLastBidderAsUnavailable(t *testing.T) {
	t.Parallel()

	snapshot := validCurrentRoundSnapshot()
	snapshot.LastBidder = [20]byte{}
	server := newCurrentRoundTestServer(t,
		fakeContractState{snapshot: func() contractstate.Snapshot { return snapshot }},
		fakeCurrentRoundReader{
			statistics: func(context.Context, int64) (cgmodel.CGRoundStats, error) {
				return validCurrentRoundStats(), nil
			},
			bidCount: func(context.Context, int64) (int64, error) {
				return 3, nil
			},
		},
	)

	response := serve(t, server, currentRoundInstance)
	assertProblem(t, response, http.StatusServiceUnavailable)
	if got := response.Header().Get("Retry-After"); got != "5" {
		t.Fatalf("Retry-After = %q, want 5", got)
	}
}

func TestGetCurrentRoundHidesInternalFailures(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		snapshot   func() contractstate.Snapshot
		statistics func(context.Context, int64) (cgmodel.CGRoundStats, error)
		bidCount   func(context.Context, int64) (int64, error)
	}{
		"malformed snapshot": {
			snapshot: func() contractstate.Snapshot {
				value := validCurrentRoundSnapshot()
				value.PrizeAmount = "password=cache-secret"
				return value
			},
		},
		"statistics failure": {
			statistics: func(context.Context, int64) (cgmodel.CGRoundStats, error) {
				return cgmodel.CGRoundStats{}, errors.New("password=database-secret")
			},
		},
		"bid count failure": {
			bidCount: func(context.Context, int64) (int64, error) {
				return 0, errors.New("rpc-token=database-secret")
			},
		},
		"statistics identity mismatch": {
			statistics: func(context.Context, int64) (cgmodel.CGRoundStats, error) {
				stats := validCurrentRoundStats()
				stats.RoundNum = 2
				return stats, nil
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			snapshot := tc.snapshot
			if snapshot == nil {
				snapshot = func() contractstate.Snapshot { return validCurrentRoundSnapshot() }
			}
			server := newCurrentRoundTestServer(t,
				fakeContractState{snapshot: snapshot},
				fakeCurrentRoundReader{
					statistics: tc.statistics,
					bidCount:   tc.bidCount,
				},
			)
			response := serve(t, server, currentRoundInstance)
			assertProblem(t, response, http.StatusInternalServerError)
			for _, secret := range []string{"cache-secret", "database-secret", "rpc-token", "password"} {
				if strings.Contains(response.Body.String(), secret) {
					t.Fatalf("500 leaked %q: %s", secret, response.Body.String())
				}
			}
		})
	}
}

func TestGetCurrentRoundHonorsCancelledContext(t *testing.T) {
	t.Parallel()

	tests := map[string]fakeCurrentRoundReader{
		"statistics": {
			statistics: func(ctx context.Context, _ int64) (cgmodel.CGRoundStats, error) {
				return cgmodel.CGRoundStats{}, ctx.Err()
			},
		},
		"bid count": {
			statistics: func(context.Context, int64) (cgmodel.CGRoundStats, error) {
				return validCurrentRoundStats(), nil
			},
			bidCount: func(ctx context.Context, _ int64) (int64, error) {
				return 0, ctx.Err()
			},
		},
	}
	for name, currentRounds := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			server := newCurrentRoundTestServer(t,
				fakeContractState{snapshot: func() contractstate.Snapshot { return validCurrentRoundSnapshot() }},
				currentRounds,
			)
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			response := serveCurrentRoundContext(t, server, ctx)
			assertProblem(t, response, http.StatusInternalServerError)
		})
	}
}

func TestGetCurrentRoundResponseIsDeterministic(t *testing.T) {
	t.Parallel()

	server := newCurrentRoundTestServer(t,
		fakeContractState{snapshot: func() contractstate.Snapshot { return validCurrentRoundSnapshot() }},
		fakeCurrentRoundReader{
			statistics: func(context.Context, int64) (cgmodel.CGRoundStats, error) {
				return validCurrentRoundStats(), nil
			},
			bidCount: func(context.Context, int64) (int64, error) {
				return 3, nil
			},
		},
	)
	first := serve(t, server, currentRoundInstance)
	second := serve(t, server, currentRoundInstance)
	if first.Code != second.Code || !bytes.Equal(first.Body.Bytes(), second.Body.Bytes()) {
		t.Fatalf("nondeterministic response: first=%d %s second=%d %s",
			first.Code, first.Body.String(), second.Code, second.Body.String())
	}
}

func newCurrentRoundTestServer(
	t *testing.T,
	state contractStateReader,
	currentRounds currentRoundReader,
) *Server {
	t.Helper()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server, err := newServer(
		nil,
		fakeBidReader{},
		fakeRoundReader{},
		currentRounds,
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
		fakeRandomWalkReader{},
		state,
		logger,
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func serveCurrentRoundContext(
	t *testing.T,
	server *Server,
	ctx context.Context,
) *httptest.ResponseRecorder {
	t.Helper()
	router := httpx.NewRouter()
	server.RegisterRoutes(router)
	request := httptest.NewRequest(http.MethodGet, currentRoundInstance, nil).WithContext(ctx)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}
