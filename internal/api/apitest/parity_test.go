//go:build integration

package apitest

import (
	"net/http"
	"reflect"
	"slices"
	"strings"
	"testing"
)

// Fixture handles referenced by the parity cases (see
// internal/testfixtures/seed/*.sql); the participant addresses live in
// ethchain.go, shared with the chain stubs.
const (
	aidCarol = "23"
	aidDave  = "24"

	nftDonationContract = "0x2700000000000000000000000000000000000027"

	// Fixture time range covering all CosmicGame bid activity.
	cgFrom = "1767225600"
	cgTo   = "1767230000"
	cgStep = "600"
	// Fixture time range covering all RandomWalk trading activity.
	rwFrom = "1767228500"
	rwTo   = "1767229800"
)

// parityCase describes one snapshotted request against a registered route.
type parityCase struct {
	path   string   // concrete request path (route params substituted)
	host   string   // optional Host header override (metadata dispatch)
	redact redactor // optional volatile-field redaction
}

// redactStringFields replaces the named top-level fields with a fixed
// placeholder when present; used for legitimately random values.
func redactStringFields(fields ...string) redactor {
	return func(t *testing.T, body any) any {
		t.Helper()
		m, ok := body.(map[string]any)
		if !ok {
			t.Fatalf("expected JSON object for redaction, got %T", body)
		}
		for _, f := range fields {
			if _, ok := m[f]; ok {
				m[f] = "<redacted>"
			}
		}
		return m
	}
}

// parityCases maps every registered GET route template to at least one
// concrete request. TestAPIParity fails if a registered GET route has no
// entry, so a newly added route cannot dodge the parity suite.
var parityCases = map[string][]parityCase{
	// --- system -------------------------------------------------------------
	"/healthz": {{path: "/healthz"}},
	"/readyz":  {{path: "/readyz"}},

	// --- bare ERC-721 metadata (host dispatch) -------------------------------
	"/metadata/{token_id}": {
		{path: "/metadata/10", host: "api.randomwalknft.com"},
		{path: "/metadata/1", host: "nfts.cosmicsignature.com"},
	},

	// --- FAQ proxy ------------------------------------------------------------
	"/api/cosmicgame/faq/health": {{path: "/api/cosmicgame/faq/health"}},

	// --- CosmicGame: statistics ----------------------------------------------
	"/api/cosmicgame/statistics/dashboard":                 {{path: "/api/cosmicgame/statistics/dashboard"}},
	"/api/cosmicgame/statistics/counters":                  {{path: "/api/cosmicgame/statistics/counters"}},
	"/api/cosmicgame/statistics/unique/bidders":            {{path: "/api/cosmicgame/statistics/unique/bidders"}},
	"/api/cosmicgame/statistics/unique/winners":            {{path: "/api/cosmicgame/statistics/unique/winners"}},
	"/api/cosmicgame/statistics/leaderboard/roi":           {{path: "/api/cosmicgame/statistics/leaderboard/roi"}},
	"/api/cosmicgame/statistics/claims/by_round":           {{path: "/api/cosmicgame/statistics/claims/by_round"}},
	"/api/cosmicgame/statistics/claims/detail/{round_num}": {{path: "/api/cosmicgame/statistics/claims/detail/0"}},
	"/api/cosmicgame/statistics/unique/donors":             {{path: "/api/cosmicgame/statistics/unique/donors"}},
	"/api/cosmicgame/statistics/unique/stakers/cst":        {{path: "/api/cosmicgame/statistics/unique/stakers/cst"}},
	"/api/cosmicgame/statistics/unique/stakers/randomwalk": {
		{path: "/api/cosmicgame/statistics/unique/stakers/randomwalk"},
	},
	"/api/cosmicgame/statistics/unique/stakers/rwalk": {{path: "/api/cosmicgame/statistics/unique/stakers/rwalk"}},
	"/api/cosmicgame/statistics/unique/stakers/both":  {{path: "/api/cosmicgame/statistics/unique/stakers/both"}},
	"/api/cosmicgame/statistics/bidding/activity/{init_ts}/{fin_ts}/{interval_secs}": {
		{path: "/api/cosmicgame/statistics/bidding/activity/" + cgFrom + "/" + cgTo + "/" + cgStep},
	},
	"/api/cosmicgame/statistics/bidding/frequency/{init_ts}/{fin_ts}/{interval_secs}": {
		{path: "/api/cosmicgame/statistics/bidding/frequency/" + cgFrom + "/" + cgTo + "/" + cgStep},
	},
	"/api/cosmicgame/statistics/bidding/top_active_periods/{n}/{init_ts}/{fin_ts}": {
		{path: "/api/cosmicgame/statistics/bidding/top_active_periods/3/" + cgFrom + "/" + cgTo},
	},
	"/api/cosmicgame/statistics/bidding/time_bounds": {{path: "/api/cosmicgame/statistics/bidding/time_bounds"}},

	// --- CosmicGame: rounds ----------------------------------------------------
	"/api/cosmicgame/rounds/list/{offset}/{limit}": {{path: "/api/cosmicgame/rounds/list/0/10"}},
	"/api/cosmicgame/rounds/info/{prize_num}":      {{path: "/api/cosmicgame/rounds/info/0"}},
	"/api/cosmicgame/rounds/current/time":          {{path: "/api/cosmicgame/rounds/current/time"}},

	// --- CosmicGame: prizes ------------------------------------------------------
	"/api/cosmicgame/prizes/history/global/{offset}/{limit}": {{path: "/api/cosmicgame/prizes/history/global/0/10"}},
	"/api/cosmicgame/prizes/history/by_user/{user_addr}/{offset}/{limit}": {
		{path: "/api/cosmicgame/prizes/history/by_user/" + addrAlice + "/0/10"},
	},
	"/api/cosmicgame/prizes/eth/all/global":                  {{path: "/api/cosmicgame/prizes/eth/all/global"}},
	"/api/cosmicgame/prizes/eth/all/global/{offset}/{limit}": {{path: "/api/cosmicgame/prizes/eth/all/global/0/10"}},
	"/api/cosmicgame/prizes/eth/raffle/global":               {{path: "/api/cosmicgame/prizes/eth/raffle/global"}},
	"/api/cosmicgame/prizes/eth/raffle/global/{offset}/{limit}": {
		{path: "/api/cosmicgame/prizes/eth/raffle/global/0/10"},
	},
	"/api/cosmicgame/prizes/eth/chronowarrior/global": {{path: "/api/cosmicgame/prizes/eth/chronowarrior/global"}},
	"/api/cosmicgame/prizes/eth/chronowarrior/global/{offset}/{limit}": {
		{path: "/api/cosmicgame/prizes/eth/chronowarrior/global/0/10"},
	},
	"/api/cosmicgame/prizes/eth/all/by_user/{user_addr}":    {{path: "/api/cosmicgame/prizes/eth/all/by_user/" + addrBob}},
	"/api/cosmicgame/prizes/eth/raffle/by_user/{user_addr}": {{path: "/api/cosmicgame/prizes/eth/raffle/by_user/" + addrBob}},
	"/api/cosmicgame/prizes/eth/chronowarrior/by_user/{user_addr}": {
		{path: "/api/cosmicgame/prizes/eth/chronowarrior/by_user/" + addrAlice},
	},
	"/api/cosmicgame/prizes/eth/unclaimed/by_user/{user_addr}/{offset}/{limit}": {
		{path: "/api/cosmicgame/prizes/eth/unclaimed/by_user/" + addrCarol + "/0/10"},
	},
	"/api/cosmicgame/prizes/deposits/raffle/by_user/{user_addr}": {
		{path: "/api/cosmicgame/prizes/deposits/raffle/by_user/" + addrBob},
	},
	"/api/cosmicgame/prizes/deposits/chrono_warrior/by_user/{user_addr}": {
		{path: "/api/cosmicgame/prizes/deposits/chrono_warrior/by_user/" + addrAlice},
	},
	"/api/cosmicgame/prizes/deposits/unclaimed/by_user/{user_addr}/{offset}/{limit}": {
		{path: "/api/cosmicgame/prizes/deposits/unclaimed/by_user/" + addrCarol + "/0/10"},
	},

	// --- CosmicGame: bids ---------------------------------------------------------
	"/api/cosmicgame/bid/list/all/{offset}/{limit}": {{path: "/api/cosmicgame/bid/list/all/0/20"}},
	"/api/cosmicgame/bid/info/{evtlog_id}":          {{path: "/api/cosmicgame/bid/info/5004"}},
	"/api/cosmicgame/bid/info_by_pos/{round_num}/{bid_position}": {
		{path: "/api/cosmicgame/bid/info_by_pos/0/1"},
	},
	"/api/cosmicgame/bid/with_message/by_round/{round}": {{path: "/api/cosmicgame/bid/with_message/by_round/0"}},
	"/api/cosmicgame/bid/list/by_round/{round_num}/{sort}/{offset}/{limit}": {
		{path: "/api/cosmicgame/bid/list/by_round/0/0/0/10"},
		{path: "/api/cosmicgame/bid/list/by_round/0/1/0/10"},
	},
	// Explicit range: the parameterless default spans epoch 0..2^31-1 at daily
	// buckets — a deterministic but ~250k-line response not worth a golden.
	"/api/cosmicgame/bid/bid_type_ratio": {
		{path: "/api/cosmicgame/bid/bid_type_ratio?from_ts=" + cgFrom + "&to_ts=" + cgTo + "&interval_secs=" + cgStep},
	},
	"/api/cosmicgame/bid/used_randomwalk_nfts": {{path: "/api/cosmicgame/bid/used_randomwalk_nfts"}},
	"/api/cosmicgame/bid/used_rwalk_nfts":      {{path: "/api/cosmicgame/bid/used_rwalk_nfts"}},
	"/api/cosmicgame/bid/cst_price":            {{path: "/api/cosmicgame/bid/cst_price"}},
	"/api/cosmicgame/bid/eth_price":            {{path: "/api/cosmicgame/bid/eth_price"}},
	"/api/cosmicgame/bid/current_special_winners": {
		{path: "/api/cosmicgame/bid/current_special_winners"},
	},
	"/api/cosmicgame/get_banned_bids": {{path: "/api/cosmicgame/get_banned_bids"}},

	// --- CosmicGame: CS NFTs (cst/*) -----------------------------------------------
	"/api/cosmicgame/cst/list/all/{offset}/{limit}": {{path: "/api/cosmicgame/cst/list/all/0/20"}},
	"/api/cosmicgame/cst/list/by_user/{user_addr}/{offset}/{limit}": {
		{path: "/api/cosmicgame/cst/list/by_user/" + addrCarol + "/0/10"},
	},
	"/api/cosmicgame/cst/info/{token_id}":          {{path: "/api/cosmicgame/cst/info/1"}},
	"/api/cosmicgame/cst/metadata/{token_id}":      {{path: "/api/cosmicgame/cst/metadata/1"}},
	"/cg/metadata/{token_id}":                      {{path: "/cg/metadata/1"}},
	"/api/cosmicgame/cst/names/history/{token_id}": {{path: "/api/cosmicgame/cst/names/history/1"}},
	"/api/cosmicgame/cst/names/search/{name}":      {{path: "/api/cosmicgame/cst/names/search/Genesis"}},
	"/api/cosmicgame/cst/names/named_only":         {{path: "/api/cosmicgame/cst/names/named_only"}},
	"/api/cosmicgame/cst/transfers/all/{token_id}/{offset}/{limit}": {
		{path: "/api/cosmicgame/cst/transfers/all/2/0/10"},
	},
	"/api/cosmicgame/cst/transfers/by_user/{user_addr}/{offset}/{limit}": {
		{path: "/api/cosmicgame/cst/transfers/by_user/" + addrDave + "/0/10"},
	},
	"/api/cosmicgame/cst/distribution": {{path: "/api/cosmicgame/cst/distribution"}},

	// --- CosmicGame: CosmicToken (ct/*) ---------------------------------------------
	"/api/cosmicgame/ct/balances":   {{path: "/api/cosmicgame/ct/balances"}},
	"/api/cosmicgame/ct/statistics": {{path: "/api/cosmicgame/ct/statistics"}},
	"/api/cosmicgame/ct/summary/by_user/{user_addr}": {
		{path: "/api/cosmicgame/ct/summary/by_user/" + addrAlice},
	},
	"/api/cosmicgame/ct/transfers/by_user/{user_addr}/{offset}/{limit}": {
		{path: "/api/cosmicgame/ct/transfers/by_user/" + addrAlice + "/0/10"},
	},
	"/api/cosmicgame/ct/total_supply_history_by_bid": {{path: "/api/cosmicgame/ct/total_supply_history_by_bid"}},
	"/api/cosmicgame/ct/total_supply_history_by_date/{from_date}/{to_date}": {
		{path: "/api/cosmicgame/ct/total_supply_history_by_date/20260101/20260102"},
	},

	// --- CosmicGame: users ------------------------------------------------------------
	"/api/cosmicgame/user/info/{user_addr}":          {{path: "/api/cosmicgame/user/info/" + addrAlice}},
	"/api/cosmicgame/user/notif_red_box/{user_addr}": {{path: "/api/cosmicgame/user/notif_red_box/" + addrAlice}},
	"/api/cosmicgame/user/balances/{user_addr}":      {{path: "/api/cosmicgame/user/balances/" + addrAlice}},

	// --- CosmicGame: donations -----------------------------------------------------------
	"/api/cosmicgame/donations/eth/simple/list/{offset}/{limit}": {
		{path: "/api/cosmicgame/donations/eth/simple/list/0/10"},
	},
	"/api/cosmicgame/donations/eth/simple/by_round/{round_num}": {
		{path: "/api/cosmicgame/donations/eth/simple/by_round/0"},
	},
	"/api/cosmicgame/donations/eth/with_info/list/{offset}/{limit}": {
		{path: "/api/cosmicgame/donations/eth/with_info/list/0/10"},
	},
	"/api/cosmicgame/donations/eth/with_info/by_round/{round_num}": {
		{path: "/api/cosmicgame/donations/eth/with_info/by_round/0"},
	},
	"/api/cosmicgame/donations/eth/with_info/info/{record_id}": {
		{path: "/api/cosmicgame/donations/eth/with_info/info/0"},
	},
	"/api/cosmicgame/donations/eth/by_user/{user_addr}": {
		{path: "/api/cosmicgame/donations/eth/by_user/" + addrDave},
	},
	"/api/cosmicgame/donations/eth/both/by_round/{round_num}": {
		{path: "/api/cosmicgame/donations/eth/both/by_round/0"},
	},
	"/api/cosmicgame/donations/eth/both/all":        {{path: "/api/cosmicgame/donations/eth/both/all"}},
	"/api/cosmicgame/donations/charity/deposits":    {{path: "/api/cosmicgame/donations/charity/deposits"}},
	"/api/cosmicgame/donations/charity/cg_deposits": {{path: "/api/cosmicgame/donations/charity/cg_deposits"}},
	"/api/cosmicgame/donations/charity/voluntary":   {{path: "/api/cosmicgame/donations/charity/voluntary"}},
	"/api/cosmicgame/donations/charity/withdrawals": {{path: "/api/cosmicgame/donations/charity/withdrawals"}},
	"/api/cosmicgame/donations/nft/list/{offset}/{limit}": {
		{path: "/api/cosmicgame/donations/nft/list/0/10"},
	},
	"/api/cosmicgame/donations/nft/info/{record_id}": {{path: "/api/cosmicgame/donations/nft/info/1"}},
	"/api/cosmicgame/donations/nft/by_user/{user_addr}": {
		{path: "/api/cosmicgame/donations/nft/by_user/" + addrBob},
	},
	"/api/cosmicgame/donations/nft/claims": {{path: "/api/cosmicgame/donations/nft/claims"}},
	"/api/cosmicgame/donations/nft/claims/{offset}/{limit}": {
		{path: "/api/cosmicgame/donations/nft/claims/0/10"},
	},
	"/api/cosmicgame/donations/nft/claims/by_user/{user_addr}": {
		{path: "/api/cosmicgame/donations/nft/claims/by_user/" + addrAlice},
	},
	"/api/cosmicgame/donations/nft/statistics": {{path: "/api/cosmicgame/donations/nft/statistics"}},
	"/api/cosmicgame/donations/nft/by_round/{prize_num}": {
		{path: "/api/cosmicgame/donations/nft/by_round/0"},
	},
	"/api/cosmicgame/donations/nft/by_token/{token_addr}": {
		{path: "/api/cosmicgame/donations/nft/by_token/" + nftDonationContract},
	},
	"/api/cosmicgame/donations/nft/unclaimed/by_round/{prize_num}": {
		{path: "/api/cosmicgame/donations/nft/unclaimed/by_round/0"}, // fully claimed round
		{path: "/api/cosmicgame/donations/nft/unclaimed/by_round/2"}, // emma's donation pending
	},
	"/api/cosmicgame/donations/nft/unclaimed/by_user/{user_addr}": {
		{path: "/api/cosmicgame/donations/nft/unclaimed/by_user/" + addrEmma},
	},
	"/api/cosmicgame/donations/erc20/by_round/detailed/{round_num}": {
		{path: "/api/cosmicgame/donations/erc20/by_round/detailed/0"},
	},
	"/api/cosmicgame/donations/erc20/by_round/all/{round_num}": {
		{path: "/api/cosmicgame/donations/erc20/by_round/all/0"},
	},
	"/api/cosmicgame/donations/erc20/by_round/summarized/{round_num}": {
		{path: "/api/cosmicgame/donations/erc20/by_round/summarized/0"},
	},
	"/api/cosmicgame/donations/erc20/donated/by_user/{user_addr}": {
		{path: "/api/cosmicgame/donations/erc20/donated/by_user/" + addrAlice},
	},
	"/api/cosmicgame/donations/erc20/by_user/{user_addr}": {
		{path: "/api/cosmicgame/donations/erc20/by_user/" + addrAlice},
	},
	"/api/cosmicgame/donations/erc20/global/{offset}/{limit}": {
		{path: "/api/cosmicgame/donations/erc20/global/0/10"},
	},
	"/api/cosmicgame/donations/erc20/info/{record_id}": {{path: "/api/cosmicgame/donations/erc20/info/1"}},
	"/api/cosmicgame/donations/erc20/claims":           {{path: "/api/cosmicgame/donations/erc20/claims"}},
	"/api/cosmicgame/donations/erc20/claims/{offset}/{limit}": {
		{path: "/api/cosmicgame/donations/erc20/claims/0/10"},
	},
	"/api/cosmicgame/donations/erc20/claims/by_user/{user_addr}": {
		{path: "/api/cosmicgame/donations/erc20/claims/by_user/" + addrAlice},
	},
	"/api/cosmicgame/donations/erc20/claims/by_round/{round_num}": {
		{path: "/api/cosmicgame/donations/erc20/claims/by_round/0"},
	},

	// --- CosmicGame: raffles -----------------------------------------------------------------
	"/api/cosmicgame/raffle/deposits/list":                  {{path: "/api/cosmicgame/raffle/deposits/list"}},
	"/api/cosmicgame/raffle/deposits/list/{offset}/{limit}": {{path: "/api/cosmicgame/raffle/deposits/list/0/10"}},
	"/api/cosmicgame/raffle/deposits/by_round/{round_num}":  {{path: "/api/cosmicgame/raffle/deposits/by_round/0"}},
	"/api/cosmicgame/eth_deposits/all/list/{offset}/{limit}": {
		{path: "/api/cosmicgame/eth_deposits/all/list/0/10"},
	},
	"/api/cosmicgame/eth_deposits/raffle_eth/list/{offset}/{limit}": {
		{path: "/api/cosmicgame/eth_deposits/raffle_eth/list/0/10"},
	},
	"/api/cosmicgame/eth_deposits/chronowarrior_eth/list/{offset}/{limit}": {
		{path: "/api/cosmicgame/eth_deposits/chronowarrior_eth/list/0/10"},
	},
	"/api/cosmicgame/raffle/nft/all/list":                  {{path: "/api/cosmicgame/raffle/nft/all/list"}},
	"/api/cosmicgame/raffle/nft/all/list/{offset}/{limit}": {{path: "/api/cosmicgame/raffle/nft/all/list/0/10"}},
	"/api/cosmicgame/raffle/nft/by_round/{round_num}":      {{path: "/api/cosmicgame/raffle/nft/by_round/0"}},
	"/api/cosmicgame/raffle/nft/by_user/{user_addr}":       {{path: "/api/cosmicgame/raffle/nft/by_user/" + addrDave}},

	// --- CosmicGame: staking CST ------------------------------------------------------------------
	"/api/cosmicgame/staking/cst/staked_tokens/all": {{path: "/api/cosmicgame/staking/cst/staked_tokens/all"}},
	"/api/cosmicgame/staking/cst/staked_tokens/by_user/{user_addr}": {
		{path: "/api/cosmicgame/staking/cst/staked_tokens/by_user/" + addrBob},
	},
	"/api/cosmicgame/staking/cst/actions/global/{offset}/{limit}": {
		{path: "/api/cosmicgame/staking/cst/actions/global/0/10"},
	},
	"/api/cosmicgame/staking/cst/actions/by_user/{user_addr}/{offset}/{limit}": {
		{path: "/api/cosmicgame/staking/cst/actions/by_user/" + addrAlice + "/0/10"},
	},
	"/api/cosmicgame/staking/cst/actions/info/{action_id}": {
		{path: "/api/cosmicgame/staking/cst/actions/info/1"},
	},
	"/api/cosmicgame/staking/cst/rewards/global": {{path: "/api/cosmicgame/staking/cst/rewards/global"}},
	"/api/cosmicgame/staking/cst/rewards/to_claim/by_user/{user_addr}": {
		{path: "/api/cosmicgame/staking/cst/rewards/to_claim/by_user/" + addrBob},
	},
	"/api/cosmicgame/staking/cst/rewards/collected/by_user/{user_addr}/{offset}/{limit}": {
		{path: "/api/cosmicgame/staking/cst/rewards/collected/by_user/" + addrAlice + "/0/10"},
	},
	"/api/cosmicgame/staking/cst/rewards/action_ids_by_deposit/{user_addr}/{deposit_id}": {
		{path: "/api/cosmicgame/staking/cst/rewards/action_ids_by_deposit/" + addrBob + "/501"},
	},
	"/api/cosmicgame/staking/cst/rewards/by_user/by_token/summary/{user_addr}": {
		{path: "/api/cosmicgame/staking/cst/rewards/by_user/by_token/summary/" + addrBob},
	},
	"/api/cosmicgame/staking/cst/rewards/by_user/by_token/details/{user_addr}/{token_id}": {
		{path: "/api/cosmicgame/staking/cst/rewards/by_user/by_token/details/" + addrBob + "/5"},
	},
	"/api/cosmicgame/staking/cst/rewards/by_user/by_deposit/{user_addr}": {
		{path: "/api/cosmicgame/staking/cst/rewards/by_user/by_deposit/" + addrBob},
	},
	"/api/cosmicgame/staking/cst/rewards/by_round/{round_num}": {
		{path: "/api/cosmicgame/staking/cst/rewards/by_round/0"},
	},
	"/api/cosmicgame/staking/cst/mints/global/{offset}/{limit}": {
		{path: "/api/cosmicgame/staking/cst/mints/global/0/10"},
	},
	"/api/cosmicgame/staking/cst/mints/by_user/{user_addr}": {
		{path: "/api/cosmicgame/staking/cst/mints/by_user/" + addrBob},
	},

	// --- CosmicGame: staking RandomWalk (canonical + legacy aliases) ---------------------------------
	"/api/cosmicgame/staking/randomwalk/actions/info/{action_id}": {
		{path: "/api/cosmicgame/staking/randomwalk/actions/info/101"},
	},
	"/api/cosmicgame/staking/randomwalk/actions/global/{offset}/{limit}": {
		{path: "/api/cosmicgame/staking/randomwalk/actions/global/0/10"},
	},
	"/api/cosmicgame/staking/randomwalk/actions/by_user/{user_addr}/{offset}/{limit}": {
		{path: "/api/cosmicgame/staking/randomwalk/actions/by_user/" + addrCarol + "/0/10"},
	},
	"/api/cosmicgame/staking/randomwalk/mints/global/{offset}/{limit}": {
		{path: "/api/cosmicgame/staking/randomwalk/mints/global/0/10"},
	},
	"/api/cosmicgame/staking/randomwalk/mints/by_user/{user_addr}": {
		{path: "/api/cosmicgame/staking/randomwalk/mints/by_user/" + addrCarol},
	},
	"/api/cosmicgame/staking/randomwalk/staked_tokens/all": {
		{path: "/api/cosmicgame/staking/randomwalk/staked_tokens/all"},
	},
	"/api/cosmicgame/staking/randomwalk/staked_tokens/by_user/{user_addr}": {
		{path: "/api/cosmicgame/staking/randomwalk/staked_tokens/by_user/" + addrDave},
	},
	"/api/cosmicgame/staking/rwalk/actions/info/{action_id}": {
		{path: "/api/cosmicgame/staking/rwalk/actions/info/101"},
	},
	"/api/cosmicgame/staking/rwalk/actions/global/{offset}/{limit}": {
		{path: "/api/cosmicgame/staking/rwalk/actions/global/0/10"},
	},
	"/api/cosmicgame/staking/rwalk/actions/by_user/{user_addr}/{offset}/{limit}": {
		{path: "/api/cosmicgame/staking/rwalk/actions/by_user/" + addrCarol + "/0/10"},
	},
	"/api/cosmicgame/staking/rwalk/mints/global/{offset}/{limit}": {
		{path: "/api/cosmicgame/staking/rwalk/mints/global/0/10"},
	},
	"/api/cosmicgame/staking/rwalk/mints/by_user/{user_addr}": {
		{path: "/api/cosmicgame/staking/rwalk/mints/by_user/" + addrCarol},
	},
	"/api/cosmicgame/staking/rwalk/staked_tokens/all": {
		{path: "/api/cosmicgame/staking/rwalk/staked_tokens/all"},
	},
	"/api/cosmicgame/staking/rwalk/staked_tokens/by_user/{user_addr}": {
		{path: "/api/cosmicgame/staking/rwalk/staked_tokens/by_user/" + addrDave},
	},

	// --- CosmicGame: marketing ------------------------------------------------------------------------
	"/api/cosmicgame/marketing/rewards/global/{offset}/{limit}": {
		{path: "/api/cosmicgame/marketing/rewards/global/0/10"},
	},
	"/api/cosmicgame/marketing/rewards/by_user/{user_addr}/{offset}/{limit}": {
		{path: "/api/cosmicgame/marketing/rewards/by_user/" + addrEmma + "/0/10"},
	},
	"/api/cosmicgame/marketing/config/current": {{path: "/api/cosmicgame/marketing/config/current"}},

	// --- CosmicGame: time + system ----------------------------------------------------------------------
	"/api/cosmicgame/time/current":     {{path: "/api/cosmicgame/time/current"}},
	"/api/cosmicgame/time/until_prize": {{path: "/api/cosmicgame/time/until_prize"}},
	"/api/cosmicgame/system/modelist":  {{path: "/api/cosmicgame/system/modelist"}},
	"/api/cosmicgame/system/modelist/{offset}/{limit}": {
		{path: "/api/cosmicgame/system/modelist/0/10"},
	},
	"/api/cosmicgame/system/admin_events/{evtlog_start}/{evtlog_end}": {
		{path: "/api/cosmicgame/system/admin_events/5001/5100"},
	},

	// --- RandomWalk -----------------------------------------------------------------------------------------
	"/api/randomwalk/current_offers/{order_by}": {{path: "/api/randomwalk/current_offers/0"}},
	"/api/randomwalk/floor_price":               {{path: "/api/randomwalk/floor_price"}},
	"/api/randomwalk/tokens/list/sequential":    {{path: "/api/randomwalk/tokens/list/sequential"}},
	"/api/randomwalk/tokens/list/by_period/{init_ts}/{fin_ts}": {
		{path: "/api/randomwalk/tokens/list/by_period/" + rwFrom + "/" + rwTo},
	},
	"/api/randomwalk/tokens/info/{token_id}":            {{path: "/api/randomwalk/tokens/info/10"}},
	"/api/cosmicgame/randomwalk/tokens/info/{token_id}": {{path: "/api/cosmicgame/randomwalk/tokens/info/10"}},
	"/api/randomwalk/tokens/name_changes/{token_id}":    {{path: "/api/randomwalk/tokens/name_changes/10"}},
	"/api/randomwalk/trading/history/{offset}/{limit}":  {{path: "/api/randomwalk/trading/history/0/10"}},
	"/api/randomwalk/trading/by_user/{user_aid}/{offset}/{limit}": {
		{path: "/api/randomwalk/trading/by_user/" + aidCarol + "/0/10"},
	},
	"/api/randomwalk/trading/sales/{offset}/{limit}": {{path: "/api/randomwalk/trading/sales/0/10"}},
	"/api/randomwalk/tokens/history/{token_id}/{offset}/{limit}": {
		{path: "/api/randomwalk/tokens/history/10/0/10"},
	},
	"/api/randomwalk/tokens/by_user/{user_aid}": {{path: "/api/randomwalk/tokens/by_user/" + aidDave}},
	"/api/randomwalk/statistics/by_token":       {{path: "/api/randomwalk/statistics/by_token"}},
	"/api/randomwalk/statistics/by_market":      {{path: "/api/randomwalk/statistics/by_market"}},
	"/api/randomwalk/statistics/trading_volume/{init_ts}/{fin_ts}/{interval_secs}": {
		{path: "/api/randomwalk/statistics/trading_volume/" + rwFrom + "/" + rwTo + "/600"},
	},
	"/api/randomwalk/statistics/mint_intervals": {{path: "/api/randomwalk/statistics/mint_intervals"}},
	"/api/randomwalk/statistics/floor_price/{init_ts}/{fin_ts}/{interval_secs}": {
		{path: "/api/randomwalk/statistics/floor_price/" + rwFrom + "/" + rwTo + "/600"},
	},
	"/api/randomwalk/statistics/withdrawal_chart": {{path: "/api/randomwalk/statistics/withdrawal_chart"}},
	"/api/randomwalk/user/info/{user_aid}":        {{path: "/api/randomwalk/user/info/" + aidCarol}},
	"/api/randomwalk/top5tokens":                  {{path: "/api/randomwalk/top5tokens"}},
	"/api/randomwalk/mint_report":                 {{path: "/api/randomwalk/mint_report"}},
	"/api/randomwalk/contracts":                   {{path: "/api/randomwalk/contracts"}},
	"/api/randomwalk/explore/random":              {{path: "/api/randomwalk/explore/random"}},
	"/api/randomwalk/random":                      {{path: "/api/randomwalk/random?limit=3"}},
	"/api/randomwalk/token-ranking/order":         {{path: "/api/randomwalk/token-ranking/order"}},
	"/api/randomwalk/rating_order":                {{path: "/api/randomwalk/rating_order"}},
	"/api/randomwalk/vote_count":                  {{path: "/api/randomwalk/vote_count"}},
	"/api/randomwalk/ranking/sign-challenge": {
		{path: "/api/randomwalk/ranking/sign-challenge", redact: redactStringFields("nonce")},
	},
	"/api/randomwalk/ranking/beauty-pair-ids": {{path: "/api/randomwalk/ranking/beauty-pair-ids"}},
	"/api/randomwalk/metadata/{token_id}":     {{path: "/api/randomwalk/metadata/10"}},
}

// TestAPIParity snapshots every registered GET route against the fixture
// database. Each case is fetched twice to prove the response is
// deterministic, then compared byte-for-byte against its golden file.
func TestAPIParity(t *testing.T) {
	h := server(t)

	routes := h.router.Routes()
	registeredGET := make(map[string]bool)
	for _, r := range routes {
		// V2 has a separate OpenAPI contract and golden suite. This table
		// remains the byte-for-byte freeze for v1 only. /version reports
		// build-dependent values, so it gets a shape test
		// (TestVersionEndpoint) instead of a golden.
		if r.Method == http.MethodGet && !strings.HasPrefix(r.Pattern, "/api/v2/") &&
			r.Pattern != "/version" {
			registeredGET[r.Pattern] = true
		}
	}

	// Table completeness in both directions.
	for path := range registeredGET {
		if _, ok := parityCases[path]; !ok {
			t.Errorf("registered GET route %q has no parity case; add one to parityCases", path)
		}
	}
	for path := range parityCases {
		if !registeredGET[path] {
			t.Errorf("parity case %q does not match any registered GET route", path)
		}
	}

	seenNames := make(map[string]string)
	templates := make([]string, 0, len(parityCases))
	for tmpl := range parityCases {
		templates = append(templates, tmpl)
	}
	slices.Sort(templates)

	for _, tmpl := range templates {
		for _, pc := range parityCases[tmpl] {
			name := goldenName("GET", pc.path, hostSuffix(pc.host))
			if prev, dup := seenNames[name]; dup {
				t.Fatalf("golden name collision: %q used by both %q and %q", name, prev, pc.path)
			}
			seenNames[name] = pc.path

			t.Run(name, func(t *testing.T) {
				first := h.do(t, request{path: pc.path, host: pc.host})
				second := h.do(t, request{path: pc.path, host: pc.host})

				fb := decodeAndRedact(t, pc, first.Body.Bytes())
				sb := decodeAndRedact(t, pc, second.Body.Bytes())
				if first.Code != second.Code || !reflect.DeepEqual(fb, sb) {
					t.Fatalf("nondeterministic response for %s: status %d vs %d", pc.path, first.Code, second.Code)
				}

				compareGolden(t, name, response{
					Status:      first.Code,
					ContentType: contentTypeOf(first),
					Body:        fb,
				})
			})
		}
	}
}

func decodeAndRedact(t *testing.T, pc parityCase, body []byte) any {
	t.Helper()
	decoded := canonicalJSON(t, body)
	if pc.redact != nil {
		decoded = pc.redact(t, decoded)
	}
	return decoded
}

func hostSuffix(host string) string {
	if host == "" {
		return ""
	}
	return strings.ReplaceAll(host, ".", "_")
}
