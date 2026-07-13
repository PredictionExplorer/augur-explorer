//go:build integration

package apitest

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
)

// faultTable renames a table for the duration of one subtest so a specific
// repository call fails while every earlier query in the handler succeeds.
func faultTable(t *testing.T, h *harness, table string) {
	t.Helper()
	backup := table + "_sprint4_backup"
	if _, err := h.db.Exec("ALTER TABLE " + table + " RENAME TO " + backup); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := h.db.Exec("ALTER TABLE " + backup + " RENAME TO " + table); err != nil {
			t.Errorf("restore %s: %v", table, err)
		}
	})
}

// TestRandomWalkSecondQueryFailuresAreOpaque drives each RandomWalk handler
// past its successful rw_contracts resolution and fails the handler's own
// repository query via a renamed table. This exercises the respondStoreError
// arms the cancelled-context matrix cannot reach (there, the shared contract
// lookup fails first).
func TestRandomWalkSecondQueryFailuresAreOpaque(t *testing.T) {
	h := server(t)
	tests := []struct {
		path  string
		table string
	}{
		{"/api/randomwalk/current_offers/0", "rw_new_offer"},
		{"/api/randomwalk/floor_price", "rw_new_offer"},
		{"/api/randomwalk/tokens/list/sequential", "rw_mint_evt"},
		{"/api/randomwalk/tokens/list/by_period/1767225600/1767230000", "rw_mint_evt"},
		{"/api/randomwalk/tokens/info/10", "rw_token"},
		{"/api/randomwalk/tokens/history/10/0/10", "rw_item_bought"},
		{"/api/randomwalk/tokens/by_user/23", "rw_token"},
		{"/api/randomwalk/trading/history/0/10", "rw_item_bought"},
		{"/api/randomwalk/trading/sales/0/10", "rw_item_bought"},
		{"/api/randomwalk/trading/by_user/23/0/10", "rw_new_offer"},
		{"/api/randomwalk/statistics/by_token", "rw_stats"},
		{"/api/randomwalk/statistics/by_market", "rw_mkt_stats"},
		{"/api/randomwalk/statistics/trading_volume/1767225600/1767230000/600", "rw_item_bought"},
		{"/api/randomwalk/statistics/mint_intervals", "rw_mint_evt"},
		{"/api/randomwalk/statistics/withdrawal_chart", "rw_mint_evt"},
		{"/api/randomwalk/statistics/withdrawal_chart", "rw_stats"},
		{"/api/randomwalk/statistics/floor_price/1767225600/1767230000/600", "rw_new_offer"},
		{"/api/randomwalk/statistics/floor_price/1767225600/1767230000/600", "rw_stats"},
		{"/api/randomwalk/user/info/23", "rw_new_offer"},
		{"/api/randomwalk/ranking/beauty-pair-ids?voter=" + addrAlice, "rw_ranking_match"},
	}
	for _, tt := range tests {
		t.Run(tt.table+" "+tt.path, func(t *testing.T) {
			faultTable(t, h, tt.table)
			response := h.get(t, tt.path)
			assertOpaqueInternalResponse(t, response.Code, response.Body.Bytes())
		})
	}
}

// TestCosmicGameSecondQueryFailuresAreOpaque does the same for the CosmicGame
// handlers whose failing query follows a successful address or registry
// resolution.
func TestCosmicGameSecondQueryFailuresAreOpaque(t *testing.T) {
	h := server(t)
	tests := []struct {
		path  string
		table string
	}{
		{"/api/cosmicgame/donations/charity/cg_deposits", "cg_donation_received"},
		{"/api/cosmicgame/donations/charity/voluntary", "cg_donation_received"},
		{"/api/cosmicgame/donations/charity/deposits", "cg_donation_received"},
		{"/api/cosmicgame/donations/nft/by_token/" + addrAlice, "cg_nft_donation"},
		{"/api/cosmicgame/user/notif_red_box/" + addrAlice, "cg_winner"},
		{"/api/cosmicgame/cst/metadata/1", "cg_mint_event"},
		{"/api/cosmicgame/staking/cst/rewards/collected/by_user/" + addrAlice + "/0/10", "cg_staking_eth_deposit"},
		{"/api/cosmicgame/staking/cst/rewards/action_ids_by_deposit/" + addrAlice + "/1", "cg_nft_staked_cst"},
		{"/api/cosmicgame/staking/cst/rewards/by_user/by_deposit/" + addrAlice, "cg_nft_staked_cst"},
		{"/api/cosmicgame/staking/cst/rewards/by_user/by_token/summary/" + addrAlice, "cg_st_reward"},
		{"/api/cosmicgame/staking/cst/rewards/by_user/by_token/details/" + addrAlice + "/1", "cg_nft_staked_cst"},
		{"/api/cosmicgame/statistics/bidding/activity/1767225600/1767230000/600", "cg_bid"},
		{"/api/cosmicgame/statistics/bidding/top_active_periods/5/1767225600/1767230000", "cg_bid"},
	}
	for _, tt := range tests {
		t.Run(tt.table+" "+tt.path, func(t *testing.T) {
			faultTable(t, h, tt.table)
			response := h.get(t, tt.path)
			assertOpaqueInternalResponse(t, response.Code, response.Body.Bytes())
		})
	}
}

// TestUnindexedWalletRoutesKeepEmptyShapes pins the ErrNotFound arms of the
// wallet-scoped v1 routes: a syntactically valid wallet the indexer has
// never seen answers HTTP 200 with the legacy empty payload.
func TestUnindexedWalletRoutesKeepEmptyShapes(t *testing.T) {
	h := server(t)
	const fresh = "0x9900000000000000000000000000000000000099"
	paths := []string{
		"/api/cosmicgame/prizes/history/by_user/" + fresh + "/0/10",
		"/api/cosmicgame/donations/nft/claims/by_user/" + fresh,
		"/api/cosmicgame/donations/nft/unclaimed/by_user/" + fresh,
		"/api/cosmicgame/prizes/eth/unclaimed/by_user/" + fresh + "/0/10",
		"/api/cosmicgame/cst/list/by_user/" + fresh + "/0/10",
	}
	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			response := h.get(t, path)
			if response.Code != http.StatusOK {
				t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
			}
			var envelope struct {
				Status int    `json:"status"`
				Error  string `json:"error"`
			}
			if err := json.Unmarshal(response.Body.Bytes(), &envelope); err != nil {
				t.Fatalf("decode response: %v\n%s", err, response.Body.String())
			}
			if envelope.Status != 1 || envelope.Error != "" {
				t.Fatalf("envelope = %+v", envelope)
			}
		})
	}
}

// TestUserBalancesEthBalanceFailure drives the live BalanceAt error arm
// through the real router: the node refuses eth_getBalance once.
func TestUserBalancesEthBalanceFailure(t *testing.T) {
	h := server(t)
	h.chain.FailNextRPC("eth_getBalance", "balance read refused")
	w := h.get(t, "/api/cosmicgame/user/balances/"+addrAlice)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400\n%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "Error at BalanceAt() call") {
		t.Fatalf("body = %s", w.Body.String())
	}
}

// TestCosmicGameModuleConstruction pins the constructor error arms and the
// refresh-loop lifecycle of the injected module (Phase 4 DI).
func TestCosmicGameModuleConstruction(t *testing.T) {
	h := server(t)
	discard := log.New(io.Discard, "", 0)

	t.Run("nil store is the legacy database-link error", func(t *testing.T) {
		_, err := cosmicgame.New(context.Background(), cosmicgame.Config{})
		if err == nil || !strings.Contains(err.Error(), "database link wasn't configured") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("contract registry read failure propagates", func(t *testing.T) {
		cancelled, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := cosmicgame.New(cancelled, cosmicgame.Config{
			Store:     h.store,
			EthClient: h.ethClient,
			Info:      discard,
			Error:     discard,
		})
		if err == nil || !strings.Contains(err.Error(), "reading contract addresses") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("missing eth client fails contract-state construction", func(t *testing.T) {
		_, err := cosmicgame.New(context.Background(), cosmicgame.Config{Store: h.store})
		if err == nil || !strings.Contains(err.Error(), "building contract state") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("background refresh starts and stops via ctx", func(t *testing.T) {
		module, err := cosmicgame.New(context.Background(), cosmicgame.Config{
			Store:     h.store,
			EthClient: h.ethClient,
			RPCClient: h.rpcClient,
			Info:      discard,
			Error:     discard,
		})
		if err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Run exits promptly on an already-cancelled context
		module.StartBackgroundRefresh(ctx)

		// A bare module has no state; the call is a documented no-op.
		cosmicgame.NewBare().StartBackgroundRefresh(ctx)
	})
}
