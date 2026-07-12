//go:build integration

package apitest

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

// TestV1ReadFailuresAreOpaque exercises the real frozen-v1 router with an
// already-cancelled request. Each case must take its repository failure arm
// and return the shared opaque envelope without leaking storage details.
func TestV1ReadFailuresAreOpaque(t *testing.T) {
	h := server(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()

	paths := []string{
		"/api/cosmicgame/statistics/counters",
		"/api/cosmicgame/statistics/unique/bidders",
		"/api/cosmicgame/statistics/unique/winners",
		"/api/cosmicgame/statistics/leaderboard/roi",
		"/api/cosmicgame/statistics/claims/by_round",
		"/api/cosmicgame/statistics/claims/detail/0",
		"/api/cosmicgame/statistics/unique/donors",
		"/api/cosmicgame/statistics/unique/stakers/cst",
		"/api/cosmicgame/statistics/unique/stakers/randomwalk",
		"/api/cosmicgame/statistics/unique/stakers/both",
		"/api/cosmicgame/statistics/bidding/frequency/1767225600/1767230000/600",
		"/api/cosmicgame/statistics/bidding/time_bounds",
		"/api/cosmicgame/rounds/list/0/10",
		"/api/cosmicgame/rounds/info/0",
		"/api/cosmicgame/prizes/history/global/0/10",
		"/api/cosmicgame/prizes/eth/all/global/0/10",
		"/api/cosmicgame/prizes/eth/raffle/global/0/10",
		"/api/cosmicgame/prizes/eth/chronowarrior/global/0/10",
		"/api/cosmicgame/bid/list/all/0/10",
		"/api/cosmicgame/bid/info/5004",
		"/api/cosmicgame/bid/info_by_pos/0/1",
		"/api/cosmicgame/bid/with_message/by_round/0",
		"/api/cosmicgame/bid/list/by_round/0/0/0/10",
		"/api/cosmicgame/bid/bid_type_ratio",
		"/api/cosmicgame/bid/used_randomwalk_nfts",
		"/api/cosmicgame/get_banned_bids",
		"/api/cosmicgame/cst/list/all/0/10",
		"/api/cosmicgame/cst/info/1",
		"/api/cosmicgame/cst/names/history/1",
		"/api/cosmicgame/cst/names/search/alice",
		"/api/cosmicgame/cst/names/named_only",
		"/api/cosmicgame/cst/transfers/all/1/0/10",
		"/api/cosmicgame/cst/distribution",
		"/api/cosmicgame/ct/balances",
		"/api/cosmicgame/ct/total_supply_history_by_bid",
		"/api/cosmicgame/ct/total_supply_history_by_date/20260101/20260102",
		"/api/cosmicgame/donations/eth/simple/list/0/10",
		"/api/cosmicgame/donations/eth/simple/by_round/0",
		"/api/cosmicgame/donations/eth/with_info/list/0/10",
		"/api/cosmicgame/donations/eth/with_info/by_round/0",
		"/api/cosmicgame/donations/eth/with_info/info/1",
		"/api/cosmicgame/donations/eth/both/by_round/0",
		"/api/cosmicgame/donations/eth/both/all",
		"/api/cosmicgame/donations/charity/deposits",
		"/api/cosmicgame/donations/charity/cg_deposits",
		"/api/cosmicgame/donations/charity/voluntary",
		"/api/cosmicgame/donations/charity/withdrawals",
		"/api/cosmicgame/donations/nft/list/0/10",
		"/api/cosmicgame/donations/nft/info/1",
		"/api/cosmicgame/donations/nft/claims/0/10",
		"/api/cosmicgame/donations/nft/statistics",
		"/api/cosmicgame/donations/nft/by_round/0",
		"/api/cosmicgame/donations/nft/unclaimed/by_round/0",
		"/api/cosmicgame/donations/erc20/by_round/detailed/0",
		"/api/cosmicgame/donations/erc20/by_round/all/0",
		"/api/cosmicgame/donations/erc20/by_round/summarized/0",
		"/api/cosmicgame/donations/erc20/global/0/10",
		"/api/cosmicgame/donations/erc20/info/1",
		"/api/cosmicgame/donations/erc20/claims/0/10",
		"/api/cosmicgame/donations/erc20/claims/by_round/0",
		"/api/cosmicgame/raffle/deposits/list/0/10",
		"/api/cosmicgame/raffle/deposits/by_round/0",
		"/api/cosmicgame/raffle/nft/all/list/0/10",
		"/api/cosmicgame/raffle/nft/by_round/0",
		"/api/cosmicgame/staking/cst/staked_tokens/all",
		"/api/cosmicgame/staking/cst/actions/global/0/10",
		"/api/cosmicgame/staking/cst/actions/info/1",
		"/api/cosmicgame/staking/cst/rewards/global",
		"/api/cosmicgame/staking/cst/rewards/by_round/0",
		"/api/cosmicgame/staking/cst/mints/global/0/10",
		"/api/cosmicgame/staking/randomwalk/actions/info/1",
		"/api/cosmicgame/staking/randomwalk/actions/global/0/10",
		"/api/cosmicgame/staking/randomwalk/mints/global/0/10",
		"/api/cosmicgame/staking/randomwalk/staked_tokens/all",
		"/api/cosmicgame/marketing/rewards/global/0/10",
		"/api/cosmicgame/system/modelist/0/10",
		"/api/cosmicgame/system/admin_events/5001/5100",
	}

	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			response := h.do(t, request{path: path, ctx: cancelled})
			if response.Code != http.StatusInternalServerError {
				t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
			}
			var envelope struct {
				Status int    `json:"status"`
				Error  string `json:"error"`
			}
			if err := json.Unmarshal(response.Body.Bytes(), &envelope); err != nil {
				t.Fatalf("decode response: %v\n%s", err, response.Body.String())
			}
			if envelope.Status != 0 || envelope.Error != "Internal server error" {
				t.Fatalf("envelope = %+v", envelope)
			}
		})
	}
}

func TestRandomWalkReadFailuresAreOpaque(t *testing.T) {
	h := server(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	paths := []string{
		"/api/randomwalk/current_offers/price",
		"/api/randomwalk/floor_price",
		"/api/randomwalk/tokens/list/sequential",
		"/api/randomwalk/tokens/list/by_period/1767225600/1767230000",
		"/api/randomwalk/tokens/info/10",
		"/api/randomwalk/tokens/name_changes/10",
		"/api/randomwalk/trading/history/0/10",
		"/api/randomwalk/trading/by_user/23/0/10",
		"/api/randomwalk/trading/sales/0/10",
		"/api/randomwalk/tokens/history/10/0/10",
		"/api/randomwalk/tokens/by_user/23",
		"/api/randomwalk/statistics/by_token",
		"/api/randomwalk/statistics/by_market",
		"/api/randomwalk/statistics/trading_volume/1767225600/1767230000/600",
		"/api/randomwalk/statistics/mint_intervals",
		"/api/randomwalk/statistics/floor_price/1767225600/1767230000/600",
		"/api/randomwalk/statistics/withdrawal_chart",
		"/api/randomwalk/user/info/23",
		"/api/randomwalk/top5tokens",
		"/api/randomwalk/mint_report",
		"/api/randomwalk/contracts",
		"/api/randomwalk/explore/random",
		"/api/randomwalk/token-ranking/order",
		"/api/randomwalk/vote_count",
		"/api/randomwalk/ranking/sign-challenge",
		"/api/randomwalk/ranking/beauty-pair-ids",
	}
	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			response := h.do(t, request{path: path, ctx: cancelled})
			if response.Code != http.StatusInternalServerError {
				t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
			}
			var envelope struct {
				Status int    `json:"status"`
				Error  string `json:"error"`
			}
			if err := json.Unmarshal(response.Body.Bytes(), &envelope); err != nil {
				t.Fatalf("decode response: %v\n%s", err, response.Body.String())
			}
			if envelope.Status != 0 || envelope.Error != "Internal server error" {
				t.Fatalf("envelope = %+v", envelope)
			}
		})
	}
}

func TestRankingMutationFailureIsOpaque(t *testing.T) {
	h := server(t)
	t.Setenv("RANKING_ADMIN_KEY", adminKey)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	response := h.do(t, request{
		method: http.MethodPost,
		path:   "/api/randomwalk/token-ranking/match",
		body: strings.NewReader(
			`{"nft1":12,"nft2":13,"nft1_won":true}`,
		),
		headers: map[string]string{"X-Ranking-Admin-Key": adminKey},
		ctx:     cancelled,
	})
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var envelope struct {
		Status int    `json:"status"`
		Error  string `json:"error"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &envelope); err != nil {
		t.Fatal(err)
	}
	if envelope.Status != 0 || envelope.Error != "Internal server error" {
		t.Fatalf("envelope = %+v", envelope)
	}
}

func TestRankingReadFailuresAfterPartialSuccessAreOpaque(t *testing.T) {
	h := server(t)
	tests := []struct {
		table  string
		backup string
		path   string
	}{
		{
			table: "rw_token_ranking", backup: "rw_token_ranking_coverage_backup",
			path: "/api/randomwalk/token-ranking/order",
		},
		{
			table: "rw_ranking_match", backup: "rw_ranking_match_coverage_backup",
			path: "/api/randomwalk/ranking/beauty-pair-ids?voter=0x2100000000000000000000000000000000000021",
		},
	}
	for _, test := range tests {
		t.Run(test.table+" "+test.path, func(t *testing.T) {
			if _, err := h.db.Exec("ALTER TABLE " + test.table + " RENAME TO " + test.backup); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if _, err := h.db.Exec("ALTER TABLE " + test.backup + " RENAME TO " + test.table); err != nil {
					t.Errorf("restore %s: %v", test.table, err)
				}
			})
			response := h.get(t, test.path)
			if response.Code != http.StatusInternalServerError {
				t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
			}
			var envelope struct {
				Status int    `json:"status"`
				Error  string `json:"error"`
			}
			if err := json.Unmarshal(response.Body.Bytes(), &envelope); err != nil {
				t.Fatal(err)
			}
			if envelope.Status != 0 || envelope.Error != "Internal server error" {
				t.Fatalf("envelope = %+v", envelope)
			}
		})
	}
}

func TestUserInfoFailuresAfterAddressResolutionAreOpaque(t *testing.T) {
	h := server(t)
	tests := []struct {
		stage string
		table string
	}{
		{stage: "profile", table: "cg_bidder"},
		{stage: "bids", table: "cg_bid"},
		{stage: "main prize claims", table: "cg_prize_claim"},
		{stage: "prize history", table: "cg_prize"},
		{stage: "NFT donations", table: "cg_nft_donation"},
		{stage: "ERC-20 donations", table: "cg_erc20_donation"},
		{stage: "ETH donations", table: "cg_eth_donated"},
		{stage: "marketing rewards", table: "cg_mkt_reward"},
		{stage: "unclaimed ETH", table: "cg_prize_deposit"},
		{stage: "unclaimed donated NFTs", table: "cg_donated_nft_claimed"},
		{stage: "claimed donated tokens", table: "cg_erc20_donation_stats"},
		{stage: "staked CST tokens", table: "cg_staked_token_cst"},
		{stage: "staked RandomWalk tokens", table: "cg_staked_token_rwalk"},
		{stage: "CST unstake history", table: "cg_nft_unstaked_cst"},
		{stage: "RandomWalk unstake history", table: "cg_nft_unstaked_rwalk"},
		{stage: "CosmicToken transfers", table: "cg_erc20_transfer"},
		{stage: "CosmicSignature transfers", table: "cg_erc721_transfer"},
	}
	for _, test := range tests {
		t.Run(test.stage, func(t *testing.T) {
			backup := test.table + "_user_info_failure_backup"
			if _, err := h.db.Exec("ALTER TABLE " + test.table + " RENAME TO " + backup); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if _, err := h.db.Exec("ALTER TABLE " + backup + " RENAME TO " + test.table); err != nil {
					t.Errorf("restore %s: %v", test.table, err)
				}
			})

			response := h.get(t, "/api/cosmicgame/user/info/"+addrAlice)
			assertOpaqueInternalResponse(t, response.Code, response.Body.Bytes())
		})
	}
}

func TestUserScopedReadFailuresAfterAddressResolutionAreOpaque(t *testing.T) {
	h := server(t)
	tests := []struct {
		name  string
		table string
		path  string
	}{
		{name: "prize history", table: "cg_prize", path: "/api/cosmicgame/prizes/history/by_user/" + addrAlice + "/0/10"},
		{name: "all ETH deposits", table: "cg_prize_deposit", path: "/api/cosmicgame/prizes/eth/all/by_user/" + addrBob},
		{name: "raffle ETH deposits", table: "cg_prize_deposit", path: "/api/cosmicgame/prizes/eth/raffle/by_user/" + addrBob},
		{name: "chrono ETH deposits", table: "cg_chrono_warrior_prize", path: "/api/cosmicgame/prizes/eth/chronowarrior/by_user/" + addrAlice},
		{name: "unclaimed ETH deposits", table: "cg_prize_deposit", path: "/api/cosmicgame/prizes/eth/unclaimed/by_user/" + addrCarol + "/0/10"},
		{name: "raffle allocation history", table: "cg_raffle_eth_prize", path: "/api/cosmicgame/prizes/deposits/raffle/by_user/" + addrBob},
		{name: "chrono allocation history", table: "cg_chrono_warrior_prize", path: "/api/cosmicgame/prizes/deposits/chrono_warrior/by_user/" + addrAlice},
		{name: "raffle NFT winnings", table: "cg_raffle_nft_prize", path: "/api/cosmicgame/raffle/nft/by_user/" + addrDave},
		{name: "owned CosmicSignature tokens", table: "cg_mint_event", path: "/api/cosmicgame/cst/list/by_user/" + addrCarol + "/0/10"},
		{name: "CosmicSignature transfers", table: "cg_erc721_transfer", path: "/api/cosmicgame/cst/transfers/by_user/" + addrDave + "/0/10"},
		{name: "CosmicToken transfers", table: "cg_erc20_transfer", path: "/api/cosmicgame/ct/transfers/by_user/" + addrAlice + "/0/10"},
		{name: "ETH donations", table: "cg_eth_donated", path: "/api/cosmicgame/donations/eth/by_user/" + addrDave},
		{name: "NFT donations", table: "cg_nft_donation", path: "/api/cosmicgame/donations/nft/by_user/" + addrBob},
		{name: "NFT claims", table: "cg_donated_nft_claimed", path: "/api/cosmicgame/donations/nft/claims/by_user/" + addrAlice},
		{name: "unclaimed NFTs", table: "cg_donated_nft_claimed", path: "/api/cosmicgame/donations/nft/unclaimed/by_user/" + addrEmma},
		{name: "ERC20 donations", table: "cg_erc20_donation", path: "/api/cosmicgame/donations/erc20/donated/by_user/" + addrAlice},
		{name: "ERC20 prize summary", table: "cg_erc20_donation_stats", path: "/api/cosmicgame/donations/erc20/by_user/" + addrAlice},
		{name: "ERC20 claims", table: "cg_donated_tok_claimed", path: "/api/cosmicgame/donations/erc20/claims/by_user/" + addrAlice},
		{name: "staked CST tokens", table: "cg_staked_token_cst", path: "/api/cosmicgame/staking/cst/staked_tokens/by_user/" + addrBob},
		{name: "CST action history", table: "cg_nft_unstaked_cst", path: "/api/cosmicgame/staking/cst/actions/by_user/" + addrAlice + "/0/10"},
		{name: "CST claimable rewards", table: "cg_st_reward", path: "/api/cosmicgame/staking/cst/rewards/to_claim/by_user/" + addrBob},
		{name: "CST raffle mints", table: "cg_raffle_nft_prize", path: "/api/cosmicgame/staking/cst/mints/by_user/" + addrBob},
		{name: "RandomWalk action history", table: "cg_nft_unstaked_rwalk", path: "/api/cosmicgame/staking/randomwalk/actions/by_user/" + addrCarol + "/0/10"},
		{name: "RandomWalk raffle mints", table: "cg_raffle_nft_prize", path: "/api/cosmicgame/staking/randomwalk/mints/by_user/" + addrCarol},
		{name: "staked RandomWalk tokens", table: "cg_staked_token_rwalk", path: "/api/cosmicgame/staking/randomwalk/staked_tokens/by_user/" + addrDave},
		{name: "marketing rewards", table: "cg_mkt_reward", path: "/api/cosmicgame/marketing/rewards/by_user/" + addrEmma + "/0/10"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backup := test.table + "_scoped_failure_backup"
			if _, err := h.db.Exec("ALTER TABLE " + test.table + " RENAME TO " + backup); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if _, err := h.db.Exec("ALTER TABLE " + backup + " RENAME TO " + test.table); err != nil {
					t.Errorf("restore %s: %v", test.table, err)
				}
			})

			response := h.get(t, test.path)
			assertOpaqueInternalResponse(t, response.Code, response.Body.Bytes())
		})
	}
}

func TestDashboardRepositoryFailuresAreOpaque(t *testing.T) {
	h := server(t)
	tests := []struct {
		stage string
		table string
	}{
		{stage: "contract registry", table: "cg_contracts"},
		{stage: "round statistics", table: "cg_round_stats"},
		{stage: "bid count", table: "cg_bid"},
	}
	for _, test := range tests {
		t.Run(test.stage, func(t *testing.T) {
			backup := test.table + "_dashboard_failure_backup"
			if _, err := h.db.Exec("ALTER TABLE " + test.table + " RENAME TO " + backup); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if _, err := h.db.Exec("ALTER TABLE " + backup + " RENAME TO " + test.table); err != nil {
					t.Errorf("restore %s: %v", test.table, err)
				}
			})

			response := h.get(t, "/api/cosmicgame/statistics/dashboard")
			assertOpaqueInternalResponse(t, response.Code, response.Body.Bytes())
		})
	}
}

func assertOpaqueInternalResponse(t *testing.T, status int, body []byte) {
	t.Helper()
	if status != http.StatusInternalServerError {
		t.Fatalf("status = %d, body=%s", status, body)
	}
	var envelope struct {
		Status int    `json:"status"`
		Error  string `json:"error"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		t.Fatalf("decode response: %v\n%s", err, body)
	}
	if envelope.Status != 0 || envelope.Error != "Internal server error" {
		t.Fatalf("envelope = %+v", envelope)
	}
}
