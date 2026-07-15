package v2

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

func TestContractResourcesSuccess(t *testing.T) {
	t.Parallel()
	var snapshots atomic.Int64
	state := fakeContractState{snapshot: func() contractstate.Snapshot {
		snapshots.Add(1)
		return validContractSnapshot()
	}}
	server := newContractTestServer(t, fakeContractAddressReader{
		get: func(context.Context) (cgmodel.CosmicGameContractAddrs, error) {
			return validContractAddressRecord(), nil
		},
	}, state)
	tests := []struct {
		path   string
		target any
	}{
		{
			path:   contractAddressesInstance,
			target: &ContractAddressRegistry{},
		},
		{
			path:   contractConfigurationInstance,
			target: &ContractConfiguration{},
		},
		{
			path:   contractBalancesInstance,
			target: &ContractBalances{},
		},
		{
			path:   currentBidPricesInstance,
			target: &CurrentBidPrices{},
		},
		{
			path:   currentSpecialWinnersInstance,
			target: &CurrentSpecialWinners{},
		},
	}
	for _, test := range tests {
		response := serve(t, server, test.path)
		if response.Code != http.StatusOK {
			t.Fatalf("%s = %d %s", test.path, response.Code, response.Body.String())
		}
		decodeResponse(t, response, test.target)
	}
	if snapshots.Load() != 4 {
		t.Fatalf("Snapshot calls = %d, want one per cached resource", snapshots.Load())
	}
}

func TestCachedContractResourcesReturn503(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		path       string
		retryAfter string
		mutate     func(*contractstate.Snapshot)
	}{
		"configuration": {
			path:       contractConfigurationInstance,
			retryAfter: "5",
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.ConfigurationReady = false
			},
		},
		"configuration constants": {
			path:       contractConfigurationInstance,
			retryAfter: "300",
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.ConstantsReady = false
				snapshot.ConfigurationReady = false
			},
		},
		"balances": {
			path:       contractBalancesInstance,
			retryAfter: "5",
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.BalancesReady = false
			},
		},
		"bid prices": {
			path:       currentBidPricesInstance,
			retryAfter: "5",
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.BidPricesReady = false
			},
		},
		"special winners": {
			path:       currentSpecialWinnersInstance,
			retryAfter: "30",
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.SpecialWinnersReady = false
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			snapshot := validContractSnapshot()
			test.mutate(&snapshot)
			server := newContractTestServer(
				t,
				fakeContractAddressReader{},
				fakeContractState{snapshot: func() contractstate.Snapshot { return snapshot }},
			)
			response := serve(t, server, test.path)
			assertProblem(t, response, http.StatusServiceUnavailable)
			if got := response.Header().Get("Retry-After"); got != test.retryAfter {
				t.Fatalf("Retry-After = %q, want %q", got, test.retryAfter)
			}
		})
	}
}

func TestContractResourcesHideInternalErrors(t *testing.T) {
	t.Parallel()
	secret := errors.New("password=private")
	t.Run("repository", func(t *testing.T) {
		t.Parallel()
		server := newContractTestServer(t, fakeContractAddressReader{
			get: func(context.Context) (cgmodel.CosmicGameContractAddrs, error) {
				return cgmodel.CosmicGameContractAddrs{}, secret
			},
		}, fakeContractState{})
		response := serve(t, server, contractAddressesInstance)
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "private") {
			t.Fatalf("repository detail leaked: %s", response.Body.String())
		}
	})

	tests := map[string]struct {
		path   string
		mutate func(*contractstate.Snapshot)
	}{
		"configuration": {
			path: contractConfigurationInstance,
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.PriceIncrease = "bad"
			},
		},
		"balances": {
			path: contractBalancesInstance,
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.CosmicGameBalance = "bad"
			},
		},
		"bid prices": {
			path: currentBidPricesInstance,
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.ETHAuctionElapsed = snapshot.ETHAuctionDuration + 1
			},
		},
		"special winners": {
			path: currentSpecialWinnersInstance,
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.SpecialWinners.Err = secret
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			snapshot := validContractSnapshot()
			test.mutate(&snapshot)
			server := newContractTestServer(
				t,
				fakeContractAddressReader{},
				fakeContractState{snapshot: func() contractstate.Snapshot { return snapshot }},
			)
			response := serve(t, server, test.path)
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "private") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})
	}
}

func newContractTestServer(
	t *testing.T,
	addresses contractAddressReader,
	state contractStateReader,
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
		addresses,
		fakeParticipantReader{},
		fakeUserReader{},
		fakeUserHistoryReader{},
		fakeUserStakingReader{},
		fakeUserActivityReader{},
		state,
		slog.New(slog.NewTextHandler(io.Discard, nil)),
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}
