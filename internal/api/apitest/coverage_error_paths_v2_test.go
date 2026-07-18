//go:build integration

package apitest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	apiv2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
)

type v2LaterStoreFailureCase struct {
	name       string
	target     string
	template   string
	pathParams map[string]string
	table      string
}

// TestAPIV2LaterStoreFailuresAreOpaque drives multi-stage v2 handlers past
// their successful resolver or existence query, then faults a later domain
// table. The cancelled-context golden cases cannot reach these branches
// because cancellation stops the first repository call.
func TestAPIV2LaterStoreFailuresAreOpaque(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	groups := []struct {
		name  string
		cases []v2LaterStoreFailureCase
	}{
		{
			name: "round resources",
			cases: []v2LaterStoreFailureCase{
				{
					name:       "current round bid count",
					target:     v2CurrentRoundPath,
					template:   v2CurrentRoundPath,
					pathParams: map[string]string{},
					table:      "cg_bid",
				},
				{
					name:       "round prizes after existence",
					target:     "/api/v2/cosmicgame/rounds/0/prizes?limit=1",
					template:   v2ListRoundPrizes,
					pathParams: map[string]string{"round": "0"},
					table:      "cg_prize",
				},
				{
					name:       "round raffle eth after existence",
					target:     "/api/v2/cosmicgame/rounds/0/raffle-eth-deposits?limit=1",
					template:   v2ListRaffleEth,
					pathParams: map[string]string{"round": "0"},
					table:      "cg_prize_deposit",
				},
				{
					name:       "round raffle nft after existence",
					target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners?pool=bidder",
					template:   v2ListRaffleNft,
					pathParams: map[string]string{"round": "0"},
					table:      "cg_raffle_nft_prize",
				},
				{
					name:       "round claims summary after existence",
					target:     "/api/v2/cosmicgame/rounds/0/claims?limit=1",
					template:   v2RoundClaims,
					pathParams: map[string]string{"round": "0"},
					table:      "cg_erc20_donation_stats",
				},
				{
					name:       "round attached tokens after summary",
					target:     "/api/v2/cosmicgame/rounds/0/claims?limit=1",
					template:   v2RoundClaims,
					pathParams: map[string]string{"round": "0"},
					table:      "cg_erc20_donation",
				},
			},
		},
		{
			name: "cosmicgame user resources",
			cases: []v2LaterStoreFailureCase{
				{
					name:       "profile after address resolution",
					target:     "/api/v2/cosmicgame/users/" + addrAlice,
					template:   v2GetUser,
					pathParams: map[string]string{"address": addrAlice},
					table:      "cg_transfer_stats",
				},
				{
					name:       "bids after address resolution",
					target:     "/api/v2/cosmicgame/users/" + addrAlice + "/bids?limit=2",
					template:   v2ListUserBids,
					pathParams: map[string]string{"address": addrAlice},
					table:      "cg_bid",
				},
				{
					name:       "prizes after address resolution",
					target:     "/api/v2/cosmicgame/users/" + addrAlice + "/prizes?limit=3",
					template:   v2ListUserPrizes,
					pathParams: map[string]string{"address": addrAlice},
					table:      "cg_prize",
				},
				{
					name:       "raffle eth after address resolution",
					target:     "/api/v2/cosmicgame/users/" + addrAlice + "/raffle-eth-deposits?limit=1",
					template:   v2ListUserDeposits,
					pathParams: map[string]string{"address": addrAlice},
					table:      "cg_prize_deposit",
				},
				{
					name:       "donated erc20 after address resolution",
					target:     "/api/v2/cosmicgame/users/" + addrAlice + "/donated-erc20",
					template:   v2ListUserDonated20,
					pathParams: map[string]string{"address": addrAlice},
					table:      "cg_erc20_donation_stats",
				},
				{
					name:       "staking actions after address resolution",
					target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/actions?limit=1",
					template:   v2ListUserCstActions,
					pathParams: map[string]string{"address": addrAlice},
					table:      "cg_nft_staked_cst",
				},
				{
					name:       "staked tokens after address resolution",
					target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/staked-tokens",
					template:   v2ListUserCstStaked,
					pathParams: map[string]string{"address": addrBob},
					table:      "cg_staked_token_cst",
				},
				{
					name:       "staking deposits after address resolution",
					target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/deposits",
					template:   v2ListUserDeposits2,
					pathParams: map[string]string{"address": addrBob},
					table:      "cg_staker_deposit",
				},
				{
					name:       "deposit rewards after deposit and address checks",
					target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/deposits/501/rewards",
					template:   v2ListUserDepRewards,
					pathParams: map[string]string{"address": addrAlice, "depositId": "501"},
					table:      "cg_st_reward",
				},
				{
					name:       "token reward deposits after token and address checks",
					target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/token-rewards/1/deposits",
					template:   v2ListUserTokDeps,
					pathParams: map[string]string{"address": addrAlice, "nftTokenId": "1"},
					table:      "cg_st_reward",
				},
				{
					name:       "signature tokens after address resolution",
					target:     "/api/v2/cosmicgame/users/" + addrBob + "/cosmic-signature-tokens?limit=2",
					template:   v2ListUserCsTokens,
					pathParams: map[string]string{"address": addrBob},
					table:      "cg_mint_event",
				},
				{
					name:       "pending winnings after address resolution",
					target:     "/api/v2/cosmicgame/users/" + addrAlice + "/pending-winnings",
					template:   v2GetUserPendingWins,
					pathParams: map[string]string{"address": addrAlice},
					table:      "cg_prize_deposit",
				},
			},
		},
		{
			name: "token directories and global staking",
			cases: []v2LaterStoreFailureCase{
				{
					name:       "token name history after existence",
					target:     "/api/v2/cosmicgame/cosmic-signature-tokens/1/name-history",
					template:   v2ListTokenNames,
					pathParams: map[string]string{"nftTokenId": "1"},
					table:      "cg_token_name",
				},
				{
					name:       "token transfers after existence",
					target:     "/api/v2/cosmicgame/cosmic-signature-tokens/2/transfers?limit=1",
					template:   v2ListTokenTransfers,
					pathParams: map[string]string{"nftTokenId": "2"},
					table:      "cg_erc721_transfer",
				},
				{
					name:       "round staking rewards after existence",
					target:     "/api/v2/cosmicgame/rounds/0/staking-rewards?limit=1",
					template:   v2RoundStakingReward,
					pathParams: map[string]string{"round": "0"},
					table:      "cg_staker_deposit",
				},
			},
		},
		{
			name: "randomwalk resources",
			cases: []v2LaterStoreFailureCase{
				{
					name:       "token name history after existence",
					target:     "/api/v2/randomwalk/tokens/10/name-history",
					template:   v2RwTokenNames,
					pathParams: map[string]string{"tokenId": "10"},
					table:      "rw_token_name",
				},
				{
					name:       "token events after existence",
					target:     "/api/v2/randomwalk/tokens/10/events?limit=3",
					template:   v2RwTokenEvents,
					pathParams: map[string]string{"tokenId": "10"},
					table:      "rw_transfer",
				},
				{
					name:       "profile after address resolution",
					target:     "/api/v2/randomwalk/users/" + addrCarol,
					template:   v2RwUser,
					pathParams: map[string]string{"address": addrCarol},
					table:      "rw_user_stats",
				},
				{
					name:       "tokens after address resolution",
					target:     "/api/v2/randomwalk/users/" + addrDave + "/tokens?limit=1",
					template:   v2RwUserTokens,
					pathParams: map[string]string{"address": addrDave},
					table:      "rw_token",
				},
				{
					name:       "offers after address resolution",
					target:     "/api/v2/randomwalk/users/" + addrDave + "/offers?limit=1",
					template:   v2RwUserOffers,
					pathParams: map[string]string{"address": addrDave},
					table:      "rw_new_offer",
				},
			},
		},
		{
			name: "ranking resources",
			cases: []v2LaterStoreFailureCase{
				{
					name:       "random tokens after contract lookup",
					target:     v2RankingRandomTokens,
					template:   v2RankingRandomTokens,
					pathParams: map[string]string{},
					table:      "rw_token",
				},
				{
					name:       "ratings after contract lookup",
					target:     v2RankingRatings + "?limit=2",
					template:   v2RankingRatings,
					pathParams: map[string]string{},
					table:      "rw_token_ranking",
				},
				{
					name:       "pair after address and contract lookups",
					target:     v2RankingPair + "?voter=" + addrAlice,
					template:   v2RankingPair,
					pathParams: map[string]string{},
					table:      "rw_token",
				},
			},
		},
	}

	for _, group := range groups {
		t.Run(group.name, func(t *testing.T) {
			for _, tc := range group.cases {
				t.Run(tc.name, func(t *testing.T) {
					faultTable(t, h, tc.table)
					response := h.get(t, tc.target)
					assertOpaqueV2InternalProblem(t, response, tc.target, tc.table)
					validateV2Response(t, spec, v2GoldenCase{
						name:       group.name + "/" + tc.name,
						target:     tc.target,
						template:   tc.template,
						pathParams: tc.pathParams,
					}, response)
				})
			}
		})
	}
}

func assertOpaqueV2InternalProblem(
	t *testing.T,
	response *httptest.ResponseRecorder,
	target string,
	faultedTable string,
) {
	t.Helper()

	if response.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500\n%s", response.Code, response.Body.String())
	}
	if got := contentTypeOf(response); got != "application/problem+json" {
		t.Fatalf("content type = %q, want application/problem+json", got)
	}

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(response.Body.Bytes(), &fields); err != nil {
		t.Fatalf("decode problem fields: %v\n%s", err, response.Body.String())
	}
	wantFields := []string{"type", "title", "status", "detail", "instance"}
	if len(fields) != len(wantFields) {
		t.Fatalf("problem fields = %v, want exactly %v", fields, wantFields)
	}
	for _, field := range wantFields {
		if _, ok := fields[field]; !ok {
			t.Fatalf("problem is missing %q: %s", field, response.Body.String())
		}
	}

	var problem struct {
		Type     string `json:"type"`
		Title    string `json:"title"`
		Status   int    `json:"status"`
		Detail   string `json:"detail"`
		Instance string `json:"instance"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &problem); err != nil {
		t.Fatalf("decode problem: %v\n%s", err, response.Body.String())
	}
	wantInstance := target
	if query := strings.IndexByte(wantInstance, '?'); query >= 0 {
		wantInstance = wantInstance[:query]
	}
	if problem.Type != "https://api.cosmicsignature.com/problems/internal" ||
		problem.Title != "Internal server error" ||
		problem.Status != http.StatusInternalServerError ||
		problem.Detail != "The request could not be completed." ||
		problem.Instance != wantInstance {
		t.Fatalf("problem = %+v, want canonical internal problem for %q", problem, wantInstance)
	}

	body := strings.ToLower(response.Body.String())
	for _, leaked := range []string{
		faultedTable,
		faultedTable + faultTableBackupSuffix,
		"sqlstate",
		"postgres",
		"password",
		"database",
		"relation",
		"pgx",
		"pq:",
	} {
		if strings.Contains(body, strings.ToLower(leaked)) {
			t.Fatalf("problem leaks %q: %s", leaked, response.Body.String())
		}
	}
}
