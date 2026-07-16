//go:build integration

package apitest

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"

	apiv2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
)

const (
	v2ListRoundsPath     = "/api/v2/cosmicgame/rounds"
	v2CurrentRoundPath   = "/api/v2/cosmicgame/rounds/current"
	v2GetRoundPath       = "/api/v2/cosmicgame/rounds/{round}"
	v2ListRoundBidsPath  = "/api/v2/cosmicgame/rounds/{round}/bids"
	v2GetRoundBidPath    = "/api/v2/cosmicgame/rounds/{round}/bids/{position}"
	v2ListRoundPrizes    = "/api/v2/cosmicgame/rounds/{round}/prizes"
	v2ListRaffleEth      = "/api/v2/cosmicgame/rounds/{round}/raffle-eth-deposits"
	v2ListRaffleNft      = "/api/v2/cosmicgame/rounds/{round}/raffle-nft-winners"
	v2ListDonationEth    = "/api/v2/cosmicgame/rounds/{round}/eth-donations"
	v2ListDonationERC20  = "/api/v2/cosmicgame/rounds/{round}/erc20-donations"
	v2ListDonationNFT    = "/api/v2/cosmicgame/rounds/{round}/nft-donations"
	v2ContractAddresses  = "/api/v2/cosmicgame/contracts/addresses"
	v2ContractConfig     = "/api/v2/cosmicgame/contracts/configuration"
	v2ContractBalances   = "/api/v2/cosmicgame/contracts/balances"
	v2CurrentBidPrices   = "/api/v2/cosmicgame/rounds/current/bid-prices"
	v2CurrentWinners     = "/api/v2/cosmicgame/rounds/current/special-winners"
	v2GlobalStatistics   = "/api/v2/cosmicgame/statistics"
	v2StatisticsCounters = "/api/v2/cosmicgame/statistics/counters"
	v2BiddingActivity    = "/api/v2/cosmicgame/statistics/bidding/activity"
	v2BiddingFrequency   = "/api/v2/cosmicgame/statistics/bidding/frequency"
	v2BiddingTypeRatio   = "/api/v2/cosmicgame/statistics/bidding/type-ratio"
	v2BiddingTopPeriods  = "/api/v2/cosmicgame/statistics/bidding/top-active-periods"
	v2BiddingTimeBounds  = "/api/v2/cosmicgame/statistics/bidding/time-bounds"
	v2StatisticsROI      = "/api/v2/cosmicgame/statistics/leaderboard/roi"
	v2StatisticsClaims   = "/api/v2/cosmicgame/statistics/claims"
	v2ParticipantBidders = "/api/v2/cosmicgame/statistics/participants/bidders"
	v2ParticipantWinners = "/api/v2/cosmicgame/statistics/participants/winners"
	v2ParticipantDonors  = "/api/v2/cosmicgame/statistics/participants/donors"
	v2ParticipantCST     = "/api/v2/cosmicgame/statistics/participants/stakers/cst"
	v2ParticipantRWalk   = "/api/v2/cosmicgame/statistics/participants/stakers/random-walk"
	v2ParticipantBoth    = "/api/v2/cosmicgame/statistics/participants/stakers/both"
	v2RoundClaims        = "/api/v2/cosmicgame/rounds/{round}/claims"
	v2GetUser            = "/api/v2/cosmicgame/users/{address}"
	v2ListUserBids       = "/api/v2/cosmicgame/users/{address}/bids"
	v2ListUserPrizes     = "/api/v2/cosmicgame/users/{address}/prizes"
	v2ListUserDeposits   = "/api/v2/cosmicgame/users/{address}/raffle-eth-deposits"
	v2ListUserNftWins    = "/api/v2/cosmicgame/users/{address}/raffle-nft-wins"
	v2ListUserEthDons    = "/api/v2/cosmicgame/users/{address}/eth-donations"
	v2ListUserErc20Dons  = "/api/v2/cosmicgame/users/{address}/erc20-donations"
	v2ListUserNftDons    = "/api/v2/cosmicgame/users/{address}/nft-donations"
	v2ListUserDonatedNft = "/api/v2/cosmicgame/users/{address}/donated-nfts"
	v2ListUserDonated20  = "/api/v2/cosmicgame/users/{address}/donated-erc20"
	v2ListUserCstActions = "/api/v2/cosmicgame/users/{address}/staking/cst/actions"
	v2ListUserCstStaked  = "/api/v2/cosmicgame/users/{address}/staking/cst/staked-tokens"
	v2ListUserDeposits2  = "/api/v2/cosmicgame/users/{address}/staking/cst/deposits"
	v2ListUserDepRewards = "/api/v2/cosmicgame/users/{address}/staking/cst/deposits/{depositId}/rewards"
	v2ListUserTokRewards = "/api/v2/cosmicgame/users/{address}/staking/cst/token-rewards"
	v2ListUserTokDeps    = "/api/v2/cosmicgame/users/{address}/staking/cst/token-rewards/{nftTokenId}/deposits"
	v2ListUserRwActions  = "/api/v2/cosmicgame/users/{address}/staking/random-walk/actions"
	v2ListUserRwStaked   = "/api/v2/cosmicgame/users/{address}/staking/random-walk/staked-tokens"
	// #nosec G101 -- route templates, not credentials.
	v2ListUserCsTokens   = "/api/v2/cosmicgame/users/{address}/cosmic-signature-tokens"
	v2ListUserCsXfers    = "/api/v2/cosmicgame/users/{address}/cosmic-signature-transfers"
	v2GetUserCtSummary   = "/api/v2/cosmicgame/users/{address}/cosmic-token-summary"
	v2ListUserCtXfers    = "/api/v2/cosmicgame/users/{address}/cosmic-token-transfers"
	v2ListUserMktRewards = "/api/v2/cosmicgame/users/{address}/marketing-rewards"
	v2GetUserPendingWins = "/api/v2/cosmicgame/users/{address}/pending-winnings"
	// #nosec G101 -- route templates, not credentials.
	v2ListGlobalTokens = "/api/v2/cosmicgame/cosmic-signature-tokens"
	v2ListCsHolders    = "/api/v2/cosmicgame/cosmic-signature-tokens/holders"
	// #nosec G101 -- route template, not credentials.
	v2GetGlobalToken = "/api/v2/cosmicgame/cosmic-signature-tokens/{nftTokenId}"
	// #nosec G101 -- route template, not credentials.
	v2ListTokenNames = "/api/v2/cosmicgame/cosmic-signature-tokens/{nftTokenId}/name-history"
	// #nosec G101 -- route template, not credentials.
	v2ListTokenTransfers = "/api/v2/cosmicgame/cosmic-signature-tokens/{nftTokenId}/transfers"
	v2ListCtHolders      = "/api/v2/cosmicgame/cosmic-token/holders"
	v2GetCtStatistics    = "/api/v2/cosmicgame/cosmic-token/statistics"
	v2ListSupplyByBid    = "/api/v2/cosmicgame/cosmic-token/supply-history/by-bid"
	v2ListSupplyDaily    = "/api/v2/cosmicgame/cosmic-token/supply-history/daily"
	v2ListGlobalMkt      = "/api/v2/cosmicgame/marketing-rewards"
)

type v2GoldenCase struct {
	name       string
	target     string
	template   string
	pathParams map[string]string
	ctx        context.Context
}

func TestAPIV2CurrentRound(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []v2GoldenCase{
		{
			name:       "current_round_get",
			target:     v2CurrentRoundPath,
			template:   v2CurrentRoundPath,
			pathParams: map[string]string{},
		},
		{
			name:       "current_round_error_internal",
			target:     v2CurrentRoundPath,
			template:   v2CurrentRoundPath,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
	}

	runV2GoldenCases(t, h, spec, cases)

	h.gameStub.Handle("getNextEthBidPrice", func([]any) ([]any, error) {
		return nil, errors.New("forced bid-price failure")
	})
	h.state.LoadInitial(context.Background())
	t.Cleanup(func() {
		h.gameStub.Return("getNextEthBidPrice", wei("1010000000000000"))
		h.state.LoadInitial(context.Background())
	})
	unavailable := h.get(t, v2CurrentRoundPath)
	if unavailable.Code != http.StatusServiceUnavailable ||
		unavailable.Header().Get("Retry-After") != "5" {
		t.Fatalf("unavailable live state = status %d Retry-After %q, want 503/5",
			unavailable.Code, unavailable.Header().Get("Retry-After"))
	}
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{
			name:       "current_round_error_unavailable",
			target:     v2CurrentRoundPath,
			template:   v2CurrentRoundPath,
			pathParams: map[string]string{},
		},
	})
	h.gameStub.Return("getNextEthBidPrice", wei("1010000000000000"))
	h.state.LoadInitial(context.Background())
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "contracts_current_bid_prices_recovered", target: v2CurrentBidPrices, template: v2CurrentBidPrices, pathParams: map[string]string{}},
	})
}

func TestAPIV2UserResources(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	const (
		indexedZeroAddress = "0x2800000000000000000000000000000000000028"
		unindexedAddress   = "0x9900000000000000000000000000000000000099"
	)
	firstPath := "/api/v2/cosmicgame/users/" + addrAlice + "/bids?limit=2"
	firstResponse := h.get(t, firstPath)
	if firstResponse.Code != http.StatusOK {
		t.Fatalf("first user-bid page: status=%d body=%s",
			firstResponse.Code, firstResponse.Body.String())
	}
	var firstPage apiv2.CosmicGameUserBidPage
	if err := json.Unmarshal(firstResponse.Body.Bytes(), &firstPage); err != nil {
		t.Fatalf("decoding first user-bid page: %v", err)
	}
	if firstPage.Meta.NextCursor == nil {
		t.Fatal("fixture first user-bid page did not return a continuation cursor")
	}

	afterLastAliceBid := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"v":1,"a":"` + addrAlice + `","e":5004}`),
	)
	bobCursor := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"v":1,"a":"` + addrBob + `","e":5029}`),
	)
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []v2GoldenCase{
		{
			name:       "user_profile_alice",
			target:     "/api/v2/cosmicgame/users/" + addrAlice,
			template:   v2GetUser,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_profile_indexed_zero",
			target:     "/api/v2/cosmicgame/users/" + indexedZeroAddress,
			template:   v2GetUser,
			pathParams: map[string]string{"address": indexedZeroAddress},
		},
		{
			name:       "user_profile_unindexed_zero",
			target:     "/api/v2/cosmicgame/users/" + unindexedAddress,
			template:   v2GetUser,
			pathParams: map[string]string{"address": unindexedAddress},
		},
		{
			name:       "user_profile_error_invalid_address",
			target:     "/api/v2/cosmicgame/users/not-an-address",
			template:   v2GetUser,
			pathParams: map[string]string{"address": "not-an-address"},
		},
		{
			name:       "user_profile_error_internal",
			target:     "/api/v2/cosmicgame/users/" + addrAlice,
			template:   v2GetUser,
			pathParams: map[string]string{"address": addrAlice},
			ctx:        cancelledCtx,
		},
		{
			name:       "user_bids_first_page",
			target:     firstPath,
			template:   v2ListUserBids,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_bids_next_page",
			target:     firstPath + "&cursor=" + *firstPage.Meta.NextCursor,
			template:   v2ListUserBids,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_bids_exhausted",
			target:     firstPath + "&cursor=" + afterLastAliceBid,
			template:   v2ListUserBids,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_bids_empty_indexed",
			target:     "/api/v2/cosmicgame/users/" + indexedZeroAddress + "/bids?limit=2",
			template:   v2ListUserBids,
			pathParams: map[string]string{"address": indexedZeroAddress},
		},
		{
			name:       "user_bids_empty_unindexed",
			target:     "/api/v2/cosmicgame/users/" + unindexedAddress + "/bids?limit=2",
			template:   v2ListUserBids,
			pathParams: map[string]string{"address": unindexedAddress},
		},
		{
			name:       "user_bids_error_invalid_address",
			target:     "/api/v2/cosmicgame/users/not-an-address/bids",
			template:   v2ListUserBids,
			pathParams: map[string]string{"address": "not-an-address"},
		},
		{
			name:       "user_bids_error_invalid_limit",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/bids?limit=201",
			template:   v2ListUserBids,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_bids_error_bind_limit",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/bids?limit=wat",
			template:   v2ListUserBids,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_bids_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/bids?cursor=bad",
			template:   v2ListUserBids,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_bids_error_cross_user_cursor",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/bids?cursor=" + bobCursor,
			template:   v2ListUserBids,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_bids_error_internal",
			target:     firstPath,
			template:   v2ListUserBids,
			pathParams: map[string]string{"address": addrAlice},
			ctx:        cancelledCtx,
		},
	}
	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2UserHistories(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	const unindexedAddress = "0x9900000000000000000000000000000000000099"
	userEventCursor := func(address, resource string, eventLogID int64) string {
		return base64.RawURLEncoding.EncodeToString([]byte(
			`{"v":1,"a":` + strconv.Quote(address) + `,"k":` + strconv.Quote(resource) +
				`,"e":` + strconv.FormatInt(eventLogID, 10) + `}`,
		))
	}
	prizeCursor := func(address string, round, prizeType, winnerIndex int64) string {
		return base64.RawURLEncoding.EncodeToString([]byte(
			`{"v":1,"a":` + strconv.Quote(address) +
				`,"r":` + strconv.FormatInt(round, 10) +
				`,"t":` + strconv.FormatInt(prizeType, 10) +
				`,"w":` + strconv.FormatInt(winnerIndex, 10) + `}`,
		))
	}

	prizesPath := "/api/v2/cosmicgame/users/" + addrAlice + "/prizes?limit=3"
	firstPrizes := h.get(t, prizesPath)
	if firstPrizes.Code != http.StatusOK {
		t.Fatalf("first prize page: status=%d body=%s", firstPrizes.Code, firstPrizes.Body.String())
	}
	var prizePage apiv2.CosmicGameUserPrizePage
	if err := json.Unmarshal(firstPrizes.Body.Bytes(), &prizePage); err != nil {
		t.Fatalf("decoding first prize page: %v", err)
	}
	if prizePage.Meta.NextCursor == nil {
		t.Fatal("fixture first prize page did not return a continuation cursor")
	}

	depositsPath := "/api/v2/cosmicgame/users/" + addrAlice + "/raffle-eth-deposits?limit=1"
	firstDeposits := h.get(t, depositsPath)
	if firstDeposits.Code != http.StatusOK {
		t.Fatalf("first deposit page: status=%d body=%s", firstDeposits.Code, firstDeposits.Body.String())
	}
	var depositPage apiv2.CosmicGameUserRaffleEthDepositPage
	if err := json.Unmarshal(firstDeposits.Body.Bytes(), &depositPage); err != nil {
		t.Fatalf("decoding first deposit page: %v", err)
	}
	if depositPage.Meta.NextCursor == nil {
		t.Fatal("fixture first deposit page did not return a continuation cursor")
	}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []v2GoldenCase{
		{
			name:       "user_prizes_first_page",
			target:     prizesPath,
			template:   v2ListUserPrizes,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_prizes_next_page",
			target:     prizesPath + "&cursor=" + *prizePage.Meta.NextCursor,
			template:   v2ListUserPrizes,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_prizes_exhausted",
			target:     prizesPath + "&cursor=" + prizeCursor(addrAlice, 0, 9, 2),
			template:   v2ListUserPrizes,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_prizes_empty_unindexed",
			target:     "/api/v2/cosmicgame/users/" + unindexedAddress + "/prizes",
			template:   v2ListUserPrizes,
			pathParams: map[string]string{"address": unindexedAddress},
		},
		{
			name:       "user_prizes_error_invalid_address",
			target:     "/api/v2/cosmicgame/users/not-an-address/prizes",
			template:   v2ListUserPrizes,
			pathParams: map[string]string{"address": "not-an-address"},
		},
		{
			name:       "user_prizes_error_invalid_limit",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/prizes?limit=201",
			template:   v2ListUserPrizes,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_prizes_error_bind_limit",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/prizes?limit=wat",
			template:   v2ListUserPrizes,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_prizes_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/prizes?cursor=bad",
			template:   v2ListUserPrizes,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_prizes_error_cross_user_cursor",
			target:     prizesPath + "&cursor=" + prizeCursor(addrBob, 0, 0, 0),
			template:   v2ListUserPrizes,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_prizes_error_internal",
			target:     prizesPath,
			template:   v2ListUserPrizes,
			pathParams: map[string]string{"address": addrAlice},
			ctx:        cancelledCtx,
		},
		{
			name:       "user_deposits_first_page",
			target:     depositsPath,
			template:   v2ListUserDeposits,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_deposits_next_page",
			target:     depositsPath + "&cursor=" + *depositPage.Meta.NextCursor,
			template:   v2ListUserDeposits,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_deposits_claimed_with_withdrawal",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/raffle-eth-deposits?claimed=true",
			template:   v2ListUserDeposits,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_deposits_unclaimed_filter_empty",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/raffle-eth-deposits?claimed=false",
			template:   v2ListUserDeposits,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_deposits_empty_unindexed",
			target:     "/api/v2/cosmicgame/users/" + unindexedAddress + "/raffle-eth-deposits",
			template:   v2ListUserDeposits,
			pathParams: map[string]string{"address": unindexedAddress},
		},
		{
			name:       "user_deposits_error_bind_claimed",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/raffle-eth-deposits?claimed=wat",
			template:   v2ListUserDeposits,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name: "user_deposits_error_cross_resource_cursor",
			target: depositsPath + "&cursor=" +
				userEventCursor(addrAlice, "nftDonations", 5040),
			template:   v2ListUserDeposits,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_deposits_error_internal",
			target:     depositsPath,
			template:   v2ListUserDeposits,
			pathParams: map[string]string{"address": addrAlice},
			ctx:        cancelledCtx,
		},
		{
			name:       "user_nft_wins_carol_randomwalk_pool",
			target:     "/api/v2/cosmicgame/users/" + addrCarol + "/raffle-nft-wins",
			template:   v2ListUserNftWins,
			pathParams: map[string]string{"address": addrCarol},
		},
		{
			name:       "user_nft_wins_bob_staker_pool",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/raffle-nft-wins",
			template:   v2ListUserNftWins,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_nft_wins_dave_bidder_pool",
			target:     "/api/v2/cosmicgame/users/" + addrDave + "/raffle-nft-wins",
			template:   v2ListUserNftWins,
			pathParams: map[string]string{"address": addrDave},
		},
		{
			name:       "user_nft_wins_empty_indexed",
			target:     "/api/v2/cosmicgame/users/" + addrEmma + "/raffle-nft-wins",
			template:   v2ListUserNftWins,
			pathParams: map[string]string{"address": addrEmma},
		},
		{
			name:       "user_nft_wins_error_internal",
			target:     "/api/v2/cosmicgame/users/" + addrCarol + "/raffle-nft-wins",
			template:   v2ListUserNftWins,
			pathParams: map[string]string{"address": addrCarol},
			ctx:        cancelledCtx,
		},
		{
			name:       "user_nft_wins_error_bind_limit",
			target:     "/api/v2/cosmicgame/users/" + addrCarol + "/raffle-nft-wins?limit=wat",
			template:   v2ListUserNftWins,
			pathParams: map[string]string{"address": addrCarol},
		},
		{
			name:       "user_eth_donations_dave_plain",
			target:     "/api/v2/cosmicgame/users/" + addrDave + "/eth-donations",
			template:   v2ListUserEthDons,
			pathParams: map[string]string{"address": addrDave},
		},
		{
			name:       "user_eth_donations_emma_with_info",
			target:     "/api/v2/cosmicgame/users/" + addrEmma + "/eth-donations",
			template:   v2ListUserEthDons,
			pathParams: map[string]string{"address": addrEmma},
		},
		{
			name:       "user_eth_donations_empty_indexed",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/eth-donations",
			template:   v2ListUserEthDons,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name: "user_eth_donations_error_malformed_cursor",
			target: "/api/v2/cosmicgame/users/" + addrDave +
				"/eth-donations?cursor=bad",
			template:   v2ListUserEthDons,
			pathParams: map[string]string{"address": addrDave},
		},
		{
			name:       "user_eth_donations_error_bind_limit",
			target:     "/api/v2/cosmicgame/users/" + addrDave + "/eth-donations?limit=wat",
			template:   v2ListUserEthDons,
			pathParams: map[string]string{"address": addrDave},
		},
		{
			name:       "user_erc20_donations_alice",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/erc20-donations",
			template:   v2ListUserErc20Dons,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_erc20_donations_empty_indexed",
			target:     "/api/v2/cosmicgame/users/" + addrDave + "/erc20-donations",
			template:   v2ListUserErc20Dons,
			pathParams: map[string]string{"address": addrDave},
		},
		{
			name:       "user_erc20_donations_error_internal",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/erc20-donations",
			template:   v2ListUserErc20Dons,
			pathParams: map[string]string{"address": addrAlice},
			ctx:        cancelledCtx,
		},
		{
			name:       "user_erc20_donations_error_bind_limit",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/erc20-donations?limit=wat",
			template:   v2ListUserErc20Dons,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_nft_donations_bob",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/nft-donations",
			template:   v2ListUserNftDons,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_nft_donations_emma",
			target:     "/api/v2/cosmicgame/users/" + addrEmma + "/nft-donations",
			template:   v2ListUserNftDons,
			pathParams: map[string]string{"address": addrEmma},
		},
		{
			name:       "user_nft_donations_error_invalid_limit",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/nft-donations?limit=0",
			template:   v2ListUserNftDons,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_nft_donations_error_bind_limit",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/nft-donations?limit=wat",
			template:   v2ListUserNftDons,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_donated_nfts_alice_claimed",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/donated-nfts",
			template:   v2ListUserDonatedNft,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_donated_nfts_emma_unclaimed",
			target:     "/api/v2/cosmicgame/users/" + addrEmma + "/donated-nfts",
			template:   v2ListUserDonatedNft,
			pathParams: map[string]string{"address": addrEmma},
		},
		{
			name: "user_donated_nfts_emma_unclaimed_filter",
			target: "/api/v2/cosmicgame/users/" + addrEmma +
				"/donated-nfts?status=unclaimed",
			template:   v2ListUserDonatedNft,
			pathParams: map[string]string{"address": addrEmma},
		},
		{
			name: "user_donated_nfts_emma_claimed_filter_empty",
			target: "/api/v2/cosmicgame/users/" + addrEmma +
				"/donated-nfts?status=claimed",
			template:   v2ListUserDonatedNft,
			pathParams: map[string]string{"address": addrEmma},
		},
		{
			name:       "user_donated_nfts_empty_indexed",
			target:     "/api/v2/cosmicgame/users/" + addrCarol + "/donated-nfts",
			template:   v2ListUserDonatedNft,
			pathParams: map[string]string{"address": addrCarol},
		},
		{
			name: "user_donated_nfts_error_invalid_status",
			target: "/api/v2/cosmicgame/users/" + addrAlice +
				"/donated-nfts?status=wat",
			template:   v2ListUserDonatedNft,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_donated_nfts_error_bind_limit",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/donated-nfts?limit=wat",
			template:   v2ListUserDonatedNft,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_donated_erc20_alice_fully_claimed",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/donated-erc20",
			template:   v2ListUserDonated20,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_donated_erc20_empty_indexed",
			target:     "/api/v2/cosmicgame/users/" + addrCarol + "/donated-erc20",
			template:   v2ListUserDonated20,
			pathParams: map[string]string{"address": addrCarol},
		},
		{
			name: "user_donated_erc20_error_malformed_cursor",
			target: "/api/v2/cosmicgame/users/" + addrAlice +
				"/donated-erc20?cursor=bad",
			template:   v2ListUserDonated20,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_donated_erc20_error_bind_limit",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/donated-erc20?limit=wat",
			template:   v2ListUserDonated20,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_donated_erc20_error_internal",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/donated-erc20",
			template:   v2ListUserDonated20,
			pathParams: map[string]string{"address": addrAlice},
			ctx:        cancelledCtx,
		},
	}
	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2UserStaking(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	const unindexedAddress = "0x9900000000000000000000000000000000000099"
	actionCursor := func(address, resource string, eventLogID int64) string {
		return base64.RawURLEncoding.EncodeToString([]byte(
			`{"v":1,"a":` + strconv.Quote(address) + `,"k":` + strconv.Quote(resource) +
				`,"e":` + strconv.FormatInt(eventLogID, 10) + `}`,
		))
	}
	tokenCursor := func(address, resource string, tokenID int64) string {
		return base64.RawURLEncoding.EncodeToString([]byte(
			`{"v":1,"a":` + strconv.Quote(address) + `,"k":` + strconv.Quote(resource) +
				`,"t":` + strconv.FormatInt(tokenID, 10) + `}`,
		))
	}
	depositCursor := func(address string, depositID int64) string {
		return base64.RawURLEncoding.EncodeToString([]byte(
			`{"v":1,"a":` + strconv.Quote(address) +
				`,"d":` + strconv.FormatInt(depositID, 10) + `}`,
		))
	}
	depositRewardCursor := func(address string, depositID, actionID int64) string {
		return base64.RawURLEncoding.EncodeToString([]byte(
			`{"v":1,"a":` + strconv.Quote(address) +
				`,"d":` + strconv.FormatInt(depositID, 10) +
				`,"s":` + strconv.FormatInt(actionID, 10) + `}`,
		))
	}
	tokenDepositCursor := func(address string, tokenID, depositID int64) string {
		return base64.RawURLEncoding.EncodeToString([]byte(
			`{"v":1,"a":` + strconv.Quote(address) +
				`,"t":` + strconv.FormatInt(tokenID, 10) +
				`,"d":` + strconv.FormatInt(depositID, 10) + `}`,
		))
	}

	// Alice's CST history is two events (stake 5051, unstake 5056); a
	// one-item page walk crosses a real continuation cursor.
	actionsPath := "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/actions?limit=1"
	firstActions := h.get(t, actionsPath)
	if firstActions.Code != http.StatusOK {
		t.Fatalf("first action page: status=%d body=%s", firstActions.Code, firstActions.Body.String())
	}
	var actionPage apiv2.CosmicGameUserCstStakingActionPage
	if err := json.Unmarshal(firstActions.Body.Bytes(), &actionPage); err != nil {
		t.Fatalf("decoding first action page: %v", err)
	}
	if actionPage.Meta.NextCursor == nil {
		t.Fatal("fixture first action page did not return a continuation cursor")
	}

	rwalkActionsPath := "/api/v2/cosmicgame/users/" + addrCarol + "/staking/random-walk/actions?limit=1"
	firstRwalkActions := h.get(t, rwalkActionsPath)
	if firstRwalkActions.Code != http.StatusOK {
		t.Fatalf("first rwalk action page: status=%d body=%s",
			firstRwalkActions.Code, firstRwalkActions.Body.String())
	}
	var rwalkActionPage apiv2.CosmicGameUserRandomWalkStakingActionPage
	if err := json.Unmarshal(firstRwalkActions.Body.Bytes(), &rwalkActionPage); err != nil {
		t.Fatalf("decoding first rwalk action page: %v", err)
	}
	if rwalkActionPage.Meta.NextCursor == nil {
		t.Fatal("fixture first rwalk action page did not return a continuation cursor")
	}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []v2GoldenCase{
		{
			name:       "user_staking_cst_actions_first_page",
			target:     actionsPath,
			template:   v2ListUserCstActions,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_staking_cst_actions_next_page",
			target:     actionsPath + "&cursor=" + *actionPage.Meta.NextCursor,
			template:   v2ListUserCstActions,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_staking_cst_actions_exhausted",
			target:     actionsPath + "&cursor=" + actionCursor(addrAlice, "cstStakingActions", 5051),
			template:   v2ListUserCstActions,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_staking_cst_actions_bob_stake_only",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/actions",
			template:   v2ListUserCstActions,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_staking_cst_actions_empty_unindexed",
			target:     "/api/v2/cosmicgame/users/" + unindexedAddress + "/staking/cst/actions",
			template:   v2ListUserCstActions,
			pathParams: map[string]string{"address": unindexedAddress},
		},
		{
			name:       "user_staking_cst_actions_error_invalid_address",
			target:     "/api/v2/cosmicgame/users/not-an-address/staking/cst/actions",
			template:   v2ListUserCstActions,
			pathParams: map[string]string{"address": "not-an-address"},
		},
		{
			name:       "user_staking_cst_actions_error_invalid_limit",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/actions?limit=201",
			template:   v2ListUserCstActions,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_staking_cst_actions_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/actions?cursor=bad",
			template:   v2ListUserCstActions,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_staking_cst_actions_error_cross_user_cursor",
			target:     actionsPath + "&cursor=" + actionCursor(addrBob, "cstStakingActions", 5052),
			template:   v2ListUserCstActions,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name: "user_staking_cst_actions_error_cross_resource_cursor",
			target: actionsPath + "&cursor=" +
				actionCursor(addrAlice, "randomWalkStakingActions", 5056),
			template:   v2ListUserCstActions,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_staking_cst_actions_error_internal",
			target:     actionsPath,
			template:   v2ListUserCstActions,
			pathParams: map[string]string{"address": addrAlice},
			ctx:        cancelledCtx,
		},
		{
			name:       "user_staking_rwalk_actions_first_page",
			target:     rwalkActionsPath,
			template:   v2ListUserRwActions,
			pathParams: map[string]string{"address": addrCarol},
		},
		{
			name:       "user_staking_rwalk_actions_next_page",
			target:     rwalkActionsPath + "&cursor=" + *rwalkActionPage.Meta.NextCursor,
			template:   v2ListUserRwActions,
			pathParams: map[string]string{"address": addrCarol},
		},
		{
			name:       "user_staking_rwalk_actions_bob_dual_staker",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/random-walk/actions",
			template:   v2ListUserRwActions,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_staking_rwalk_actions_error_cross_resource_cursor",
			target:     rwalkActionsPath + "&cursor=" + actionCursor(addrCarol, "cstStakingActions", 5057),
			template:   v2ListUserRwActions,
			pathParams: map[string]string{"address": addrCarol},
		},
		{
			name:       "user_staking_cst_staked_tokens_bob",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/staked-tokens",
			template:   v2ListUserCstStaked,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_staking_cst_staked_tokens_alice_after_unstake",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/staked-tokens",
			template:   v2ListUserCstStaked,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name: "user_staking_cst_staked_tokens_exhausted",
			target: "/api/v2/cosmicgame/users/" + addrBob +
				"/staking/cst/staked-tokens?cursor=" + tokenCursor(addrBob, "cstStakedTokens", 5),
			template:   v2ListUserCstStaked,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name: "user_staking_cst_staked_tokens_error_cross_resource_cursor",
			target: "/api/v2/cosmicgame/users/" + addrBob +
				"/staking/cst/staked-tokens?cursor=" + tokenCursor(addrBob, "cstTokenRewards", 5),
			template:   v2ListUserCstStaked,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_staking_rwalk_staked_tokens_dave",
			target:     "/api/v2/cosmicgame/users/" + addrDave + "/staking/random-walk/staked-tokens",
			template:   v2ListUserRwStaked,
			pathParams: map[string]string{"address": addrDave},
		},
		{
			name:       "user_staking_rwalk_staked_tokens_carol_after_unstake",
			target:     "/api/v2/cosmicgame/users/" + addrCarol + "/staking/random-walk/staked-tokens",
			template:   v2ListUserRwStaked,
			pathParams: map[string]string{"address": addrCarol},
		},
		{
			name:       "user_staking_deposits_alice_fully_claimed",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/deposits",
			template:   v2ListUserDeposits2,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_staking_deposits_bob_pending",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/deposits",
			template:   v2ListUserDeposits2,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_staking_deposits_alice_claimed_filter",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/deposits?claimed=true",
			template:   v2ListUserDeposits2,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_staking_deposits_bob_claimed_filter_empty",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/deposits?claimed=true",
			template:   v2ListUserDeposits2,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_staking_deposits_bob_pending_filter",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/deposits?claimed=false",
			template:   v2ListUserDeposits2,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name: "user_staking_deposits_exhausted",
			target: "/api/v2/cosmicgame/users/" + addrBob +
				"/staking/cst/deposits?cursor=" + depositCursor(addrBob, 501),
			template:   v2ListUserDeposits2,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_staking_deposits_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/deposits?cursor=bad",
			template:   v2ListUserDeposits2,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_staking_deposits_error_internal",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/deposits",
			template:   v2ListUserDeposits2,
			pathParams: map[string]string{"address": addrBob},
			ctx:        cancelledCtx,
		},
		{
			name:       "user_staking_deposit_rewards_alice_claimed",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/deposits/501/rewards",
			template:   v2ListUserDepRewards,
			pathParams: map[string]string{"address": addrAlice, "depositId": "501"},
		},
		{
			name:       "user_staking_deposit_rewards_bob_pending",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/deposits/501/rewards",
			template:   v2ListUserDepRewards,
			pathParams: map[string]string{"address": addrBob, "depositId": "501"},
		},
		{
			name:       "user_staking_deposit_rewards_carol_uninvolved",
			target:     "/api/v2/cosmicgame/users/" + addrCarol + "/staking/cst/deposits/501/rewards",
			template:   v2ListUserDepRewards,
			pathParams: map[string]string{"address": addrCarol, "depositId": "501"},
		},
		{
			name:       "user_staking_deposit_rewards_error_missing_deposit",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/deposits/999/rewards",
			template:   v2ListUserDepRewards,
			pathParams: map[string]string{"address": addrAlice, "depositId": "999"},
		},
		{
			name:       "user_staking_deposit_rewards_error_negative_deposit",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/deposits/-1/rewards",
			template:   v2ListUserDepRewards,
			pathParams: map[string]string{"address": addrAlice, "depositId": "-1"},
		},
		{
			name: "user_staking_deposit_rewards_error_cross_deposit_cursor",
			target: "/api/v2/cosmicgame/users/" + addrAlice +
				"/staking/cst/deposits/999/rewards?cursor=" + depositRewardCursor(addrAlice, 501, 1),
			template:   v2ListUserDepRewards,
			pathParams: map[string]string{"address": addrAlice, "depositId": "999"},
		},
		{
			name:       "user_staking_token_rewards_alice_collected",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/token-rewards",
			template:   v2ListUserTokRewards,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_staking_token_rewards_bob_pending",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/token-rewards",
			template:   v2ListUserTokRewards,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name: "user_staking_token_rewards_exhausted",
			target: "/api/v2/cosmicgame/users/" + addrAlice +
				"/staking/cst/token-rewards?cursor=" + tokenCursor(addrAlice, "cstTokenRewards", 1),
			template:   v2ListUserTokRewards,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_staking_token_rewards_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/token-rewards?cursor=bad",
			template:   v2ListUserTokRewards,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_staking_token_deposits_alice_token_1",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/token-rewards/1/deposits",
			template:   v2ListUserTokDeps,
			pathParams: map[string]string{"address": addrAlice, "nftTokenId": "1"},
		},
		{
			name:       "user_staking_token_deposits_bob_token_5",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/staking/cst/token-rewards/5/deposits",
			template:   v2ListUserTokDeps,
			pathParams: map[string]string{"address": addrBob, "nftTokenId": "5"},
		},
		{
			name:       "user_staking_token_deposits_alice_no_rewards_on_token_5",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/token-rewards/5/deposits",
			template:   v2ListUserTokDeps,
			pathParams: map[string]string{"address": addrAlice, "nftTokenId": "5"},
		},
		{
			name:       "user_staking_token_deposits_error_unminted_token",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/token-rewards/999/deposits",
			template:   v2ListUserTokDeps,
			pathParams: map[string]string{"address": addrAlice, "nftTokenId": "999"},
		},
		{
			name: "user_staking_token_deposits_error_cross_token_cursor",
			target: "/api/v2/cosmicgame/users/" + addrAlice +
				"/staking/cst/token-rewards/1/deposits?cursor=" + tokenDepositCursor(addrAlice, 5, 501),
			template:   v2ListUserTokDeps,
			pathParams: map[string]string{"address": addrAlice, "nftTokenId": "1"},
		},
		{
			name:       "user_staking_token_deposits_error_internal",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/staking/cst/token-rewards/1/deposits",
			template:   v2ListUserTokDeps,
			pathParams: map[string]string{"address": addrAlice, "nftTokenId": "1"},
			ctx:        cancelledCtx,
		},
	}
	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2UserActivity(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	const (
		indexedZeroAddress = "0x2800000000000000000000000000000000000028"
		unindexedAddress   = "0x9900000000000000000000000000000000000099"
	)
	eventCursor := func(address, resource string, eventLogID int64) string {
		return base64.RawURLEncoding.EncodeToString([]byte(
			`{"v":1,"a":` + strconv.Quote(address) + `,"k":` + strconv.Quote(resource) +
				`,"e":` + strconv.FormatInt(eventLogID, 10) + `}`,
		))
	}
	ownedTokenCursor := func(address string, tokenID int64) string {
		return base64.RawURLEncoding.EncodeToString([]byte(
			`{"v":1,"a":` + strconv.Quote(address) +
				`,"t":` + strconv.FormatInt(tokenID, 10) + `}`,
		))
	}

	// Bob owns three tokens (2, 5, 9); a two-item page crosses a real
	// continuation cursor.
	tokensPath := "/api/v2/cosmicgame/users/" + addrBob + "/cosmic-signature-tokens?limit=2"
	firstTokens := h.get(t, tokensPath)
	if firstTokens.Code != http.StatusOK {
		t.Fatalf("first token page: status=%d body=%s", firstTokens.Code, firstTokens.Body.String())
	}
	var tokenPage apiv2.CosmicGameUserCosmicSignatureTokenPage
	if err := json.Unmarshal(firstTokens.Body.Bytes(), &tokenPage); err != nil {
		t.Fatalf("decoding first token page: %v", err)
	}
	if tokenPage.Meta.NextCursor == nil {
		t.Fatal("fixture first token page did not return a continuation cursor")
	}

	// Alice's Cosmic Token ledger is four events; a two-item page crosses
	// a real continuation cursor.
	ctTransfersPath := "/api/v2/cosmicgame/users/" + addrAlice + "/cosmic-token-transfers?limit=2"
	firstCtTransfers := h.get(t, ctTransfersPath)
	if firstCtTransfers.Code != http.StatusOK {
		t.Fatalf("first ct-transfer page: status=%d body=%s",
			firstCtTransfers.Code, firstCtTransfers.Body.String())
	}
	var ctTransferPage apiv2.CosmicGameUserCosmicTokenTransferPage
	if err := json.Unmarshal(firstCtTransfers.Body.Bytes(), &ctTransferPage); err != nil {
		t.Fatalf("decoding first ct-transfer page: %v", err)
	}
	if ctTransferPage.Meta.NextCursor == nil {
		t.Fatal("fixture first ct-transfer page did not return a continuation cursor")
	}

	// Dave's NFT ledger is three events (mint in, transfer out, mint in).
	csTransfersPath := "/api/v2/cosmicgame/users/" + addrDave + "/cosmic-signature-transfers?limit=1"
	firstCsTransfers := h.get(t, csTransfersPath)
	if firstCsTransfers.Code != http.StatusOK {
		t.Fatalf("first cs-transfer page: status=%d body=%s",
			firstCsTransfers.Code, firstCsTransfers.Body.String())
	}
	var csTransferPage apiv2.CosmicGameUserCosmicSignatureTransferPage
	if err := json.Unmarshal(firstCsTransfers.Body.Bytes(), &csTransferPage); err != nil {
		t.Fatalf("decoding first cs-transfer page: %v", err)
	}
	if csTransferPage.Meta.NextCursor == nil {
		t.Fatal("fixture first cs-transfer page did not return a continuation cursor")
	}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []v2GoldenCase{
		{
			name:       "user_tokens_bob_first_page",
			target:     tokensPath,
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_tokens_bob_next_page",
			target:     tokensPath + "&cursor=" + *tokenPage.Meta.NextCursor,
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_tokens_bob_exhausted",
			target:     tokensPath + "&cursor=" + ownedTokenCursor(addrBob, 9),
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_tokens_alice_named_and_chrono",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/cosmic-signature-tokens",
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_tokens_carol_staker_and_lastcst",
			target:     "/api/v2/cosmicgame/users/" + addrCarol + "/cosmic-signature-tokens",
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": addrCarol},
		},
		{
			name:       "user_tokens_empty_unindexed",
			target:     "/api/v2/cosmicgame/users/" + unindexedAddress + "/cosmic-signature-tokens",
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": unindexedAddress},
		},
		{
			name:       "user_tokens_error_invalid_address",
			target:     "/api/v2/cosmicgame/users/not-an-address/cosmic-signature-tokens",
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": "not-an-address"},
		},
		{
			name:       "user_tokens_error_invalid_limit",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/cosmic-signature-tokens?limit=201",
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_tokens_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/cosmic-signature-tokens?cursor=bad",
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_tokens_error_cross_user_cursor",
			target:     tokensPath + "&cursor=" + ownedTokenCursor(addrAlice, 1),
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name: "user_tokens_error_cross_resource_cursor",
			target: tokensPath + "&cursor=" +
				base64.RawURLEncoding.EncodeToString([]byte(
					`{"v":1,"a":`+strconv.Quote(addrBob)+`,"k":"cstStakedTokens","t":5}`)),
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_tokens_error_internal",
			target:     tokensPath,
			template:   v2ListUserCsTokens,
			pathParams: map[string]string{"address": addrBob},
			ctx:        cancelledCtx,
		},
		{
			name:       "user_cs_transfers_dave_first_page",
			target:     csTransfersPath,
			template:   v2ListUserCsXfers,
			pathParams: map[string]string{"address": addrDave},
		},
		{
			name:       "user_cs_transfers_dave_next_page",
			target:     csTransfersPath + "&cursor=" + *csTransferPage.Meta.NextCursor,
			template:   v2ListUserCsXfers,
			pathParams: map[string]string{"address": addrDave},
		},
		{
			name:       "user_cs_transfers_bob_incoming",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/cosmic-signature-transfers",
			template:   v2ListUserCsXfers,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name: "user_cs_transfers_exhausted",
			target: "/api/v2/cosmicgame/users/" + addrDave +
				"/cosmic-signature-transfers?cursor=" +
				eventCursor(addrDave, "cosmicSignatureTransfers", 5027),
			template:   v2ListUserCsXfers,
			pathParams: map[string]string{"address": addrDave},
		},
		{
			name:       "user_cs_transfers_empty_unindexed",
			target:     "/api/v2/cosmicgame/users/" + unindexedAddress + "/cosmic-signature-transfers",
			template:   v2ListUserCsXfers,
			pathParams: map[string]string{"address": unindexedAddress},
		},
		{
			name: "user_cs_transfers_error_cross_resource_cursor",
			target: csTransfersPath + "&cursor=" +
				eventCursor(addrDave, "cosmicTokenTransfers", 5027),
			template:   v2ListUserCsXfers,
			pathParams: map[string]string{"address": addrDave},
		},
		{
			name:       "user_cs_transfers_error_internal",
			target:     csTransfersPath,
			template:   v2ListUserCsXfers,
			pathParams: map[string]string{"address": addrDave},
			ctx:        cancelledCtx,
		},
		{
			name:       "user_ct_transfers_alice_first_page",
			target:     ctTransfersPath,
			template:   v2ListUserCtXfers,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_ct_transfers_alice_next_page",
			target:     ctTransfersPath + "&cursor=" + *ctTransferPage.Meta.NextCursor,
			template:   v2ListUserCtXfers,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_ct_transfers_carol_burn",
			target:     "/api/v2/cosmicgame/users/" + addrCarol + "/cosmic-token-transfers",
			template:   v2ListUserCtXfers,
			pathParams: map[string]string{"address": addrCarol},
		},
		{
			name: "user_ct_transfers_exhausted",
			target: "/api/v2/cosmicgame/users/" + addrAlice +
				"/cosmic-token-transfers?cursor=" +
				eventCursor(addrAlice, "cosmicTokenTransfers", 5005),
			template:   v2ListUserCtXfers,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name: "user_ct_transfers_error_cross_user_cursor",
			target: ctTransfersPath + "&cursor=" +
				eventCursor(addrBob, "cosmicTokenTransfers", 5049),
			template:   v2ListUserCtXfers,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name: "user_ct_transfers_error_cross_resource_cursor",
			target: ctTransfersPath + "&cursor=" +
				eventCursor(addrAlice, "marketingRewards", 5049),
			template:   v2ListUserCtXfers,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_marketing_rewards_emma",
			target:     "/api/v2/cosmicgame/users/" + addrEmma + "/marketing-rewards",
			template:   v2ListUserMktRewards,
			pathParams: map[string]string{"address": addrEmma},
		},
		{
			name: "user_marketing_rewards_exhausted",
			target: "/api/v2/cosmicgame/users/" + addrEmma +
				"/marketing-rewards?cursor=" + eventCursor(addrEmma, "marketingRewards", 5017),
			template:   v2ListUserMktRewards,
			pathParams: map[string]string{"address": addrEmma},
		},
		{
			name:       "user_marketing_rewards_alice_none",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/marketing-rewards",
			template:   v2ListUserMktRewards,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_marketing_rewards_error_invalid_limit",
			target:     "/api/v2/cosmicgame/users/" + addrEmma + "/marketing-rewards?limit=0",
			template:   v2ListUserMktRewards,
			pathParams: map[string]string{"address": addrEmma},
		},
		{
			name:       "user_token_summary_alice",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/cosmic-token-summary",
			template:   v2GetUserCtSummary,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_token_summary_carol_negative_net",
			target:     "/api/v2/cosmicgame/users/" + addrCarol + "/cosmic-token-summary",
			template:   v2GetUserCtSummary,
			pathParams: map[string]string{"address": addrCarol},
		},
		{
			name:       "user_token_summary_indexed_zero",
			target:     "/api/v2/cosmicgame/users/" + indexedZeroAddress + "/cosmic-token-summary",
			template:   v2GetUserCtSummary,
			pathParams: map[string]string{"address": indexedZeroAddress},
		},
		{
			name:       "user_token_summary_unindexed_zero",
			target:     "/api/v2/cosmicgame/users/" + unindexedAddress + "/cosmic-token-summary",
			template:   v2GetUserCtSummary,
			pathParams: map[string]string{"address": unindexedAddress},
		},
		{
			name:       "user_token_summary_error_invalid_address",
			target:     "/api/v2/cosmicgame/users/not-an-address/cosmic-token-summary",
			template:   v2GetUserCtSummary,
			pathParams: map[string]string{"address": "not-an-address"},
		},
		{
			name:       "user_token_summary_error_internal",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/cosmic-token-summary",
			template:   v2GetUserCtSummary,
			pathParams: map[string]string{"address": addrAlice},
			ctx:        cancelledCtx,
		},
		{
			name:       "user_pending_winnings_alice_deposits",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/pending-winnings",
			template:   v2GetUserPendingWins,
			pathParams: map[string]string{"address": addrAlice},
		},
		{
			name:       "user_pending_winnings_bob_staking_reward",
			target:     "/api/v2/cosmicgame/users/" + addrBob + "/pending-winnings",
			template:   v2GetUserPendingWins,
			pathParams: map[string]string{"address": addrBob},
		},
		{
			name:       "user_pending_winnings_emma_donated_nft",
			target:     "/api/v2/cosmicgame/users/" + addrEmma + "/pending-winnings",
			template:   v2GetUserPendingWins,
			pathParams: map[string]string{"address": addrEmma},
		},
		{
			name:       "user_pending_winnings_unindexed_zero",
			target:     "/api/v2/cosmicgame/users/" + unindexedAddress + "/pending-winnings",
			template:   v2GetUserPendingWins,
			pathParams: map[string]string{"address": unindexedAddress},
		},
		{
			name:       "user_pending_winnings_error_invalid_address",
			target:     "/api/v2/cosmicgame/users/not-an-address/pending-winnings",
			template:   v2GetUserPendingWins,
			pathParams: map[string]string{"address": "not-an-address"},
		},
		{
			name:       "user_pending_winnings_error_internal",
			target:     "/api/v2/cosmicgame/users/" + addrAlice + "/pending-winnings",
			template:   v2GetUserPendingWins,
			pathParams: map[string]string{"address": addrAlice},
			ctx:        cancelledCtx,
		},
	}
	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2GlobalDirectories(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	globalTokenCursor := func(named bool, name string, tokenID int64) string {
		filter := "{}"
		switch {
		case named:
			filter = `{"n":true}`
		case name != "":
			filter = `{"q":` + strconv.Quote(name) + `}`
		}
		return base64.RawURLEncoding.EncodeToString([]byte(
			`{"v":1,"f":` + filter + `,"t":` + strconv.FormatInt(tokenID, 10) + `}`,
		))
	}
	tokenEventCursor := func(resource string, tokenID, eventLogID int64) string {
		return base64.RawURLEncoding.EncodeToString([]byte(
			`{"v":1,"k":` + strconv.Quote(resource) +
				`,"t":` + strconv.FormatInt(tokenID, 10) +
				`,"e":` + strconv.FormatInt(eventLogID, 10) + `}`,
		))
	}

	// Nine fixture mints: a four-item page crosses a real continuation
	// cursor twice.
	tokensPath := v2ListGlobalTokens + "?limit=4"
	var tokenPage apiv2.CosmicGameCosmicSignatureTokenPage
	decodeV2JSON(t, h.get(t, tokensPath), &tokenPage)
	if tokenPage.Meta.NextCursor == nil {
		t.Fatal("fixture first token page did not return a continuation cursor")
	}

	// Token 2 carries a mint plus a transfer; a one-item page crosses a
	// real continuation cursor.
	transfersPath := "/api/v2/cosmicgame/cosmic-signature-tokens/2/transfers?limit=1"
	var transferPage apiv2.CosmicGameCosmicSignatureTokenTransferPage
	decodeV2JSON(t, h.get(t, transfersPath), &transferPage)
	if transferPage.Meta.NextCursor == nil {
		t.Fatal("fixture first transfer page did not return a continuation cursor")
	}

	// Walk the two ranked holder directories for their real cursors.
	holdersPath := v2ListCsHolders + "?limit=2"
	var csHolderPage apiv2.CosmicGameCosmicSignatureHolderPage
	decodeV2JSON(t, h.get(t, holdersPath), &csHolderPage)
	if csHolderPage.Meta.NextCursor == nil {
		t.Fatal("fixture first holder page did not return a continuation cursor")
	}
	ctHoldersPath := v2ListCtHolders + "?limit=2"
	var ctHolderPage apiv2.CosmicGameCosmicTokenHolderPage
	decodeV2JSON(t, h.get(t, ctHoldersPath), &ctHolderPage)
	if ctHolderPage.Meta.NextCursor == nil {
		t.Fatal("fixture first balance page did not return a continuation cursor")
	}

	// The supply ledger and the marketing ledger provide their real
	// continuation and terminal cursors.
	supplyPath := v2ListSupplyByBid + "?limit=5"
	var supplyPage apiv2.CosmicGameCosmicTokenSupplyByBidPage
	decodeV2JSON(t, h.get(t, supplyPath), &supplyPage)
	if supplyPage.Meta.NextCursor == nil {
		t.Fatal("fixture first supply page did not return a continuation cursor")
	}
	lastSupplyEvent := int64(0)
	cursor := ""
	for {
		target := v2ListSupplyByBid + "?limit=200"
		if cursor != "" {
			target += "&cursor=" + cursor
		}
		var page apiv2.CosmicGameCosmicTokenSupplyByBidPage
		decodeV2JSON(t, h.get(t, target), &page)
		if len(page.Data) > 0 {
			lastSupplyEvent = page.Data[len(page.Data)-1].EventLogId
		}
		if page.Meta.NextCursor == nil {
			break
		}
		cursor = *page.Meta.NextCursor
	}
	if lastSupplyEvent == 0 {
		t.Fatal("fixture supply ledger is empty")
	}
	terminalSupplyCursor := base64.RawURLEncoding.EncodeToString([]byte(
		`{"v":1,"s":` + strconv.FormatInt(lastSupplyEvent, 10) + `}`))

	// The fixture pays exactly one marketing reward (event 5017), so the
	// ledger's terminal cursor is synthesized from it.
	marketingPath := v2ListGlobalMkt + "?limit=1"
	var marketingPage apiv2.CosmicGameMarketingRewardPage
	decodeV2JSON(t, h.get(t, marketingPath), &marketingPage)
	if len(marketingPage.Data) != 1 || marketingPage.Meta.NextCursor != nil {
		t.Fatalf("fixture marketing ledger changed: %+v", marketingPage)
	}
	terminalMarketingCursor := base64.RawURLEncoding.EncodeToString([]byte(
		`{"v":1,"m":` + strconv.FormatInt(marketingPage.Data[0].EventLogId, 10) + `}`))

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []v2GoldenCase{
		{
			name:       "global_tokens_first_page",
			target:     tokensPath,
			template:   v2ListGlobalTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "global_tokens_next_page",
			target:     tokensPath + "&cursor=" + *tokenPage.Meta.NextCursor,
			template:   v2ListGlobalTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "global_tokens_exhausted",
			target:     tokensPath + "&cursor=" + globalTokenCursor(false, "", 0),
			template:   v2ListGlobalTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "global_tokens_named_only",
			target:     v2ListGlobalTokens + "?named=true",
			template:   v2ListGlobalTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "global_tokens_name_search",
			target:     v2ListGlobalTokens + "?name=gene",
			template:   v2ListGlobalTokens,
			pathParams: map[string]string{},
		},
		{
			// A literal % matches nothing: ILIKE wildcards are escaped.
			name:       "global_tokens_name_search_escaped_wildcard",
			target:     v2ListGlobalTokens + "?name=%25",
			template:   v2ListGlobalTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "global_tokens_error_contradictory_filters",
			target:     v2ListGlobalTokens + "?named=true&name=x",
			template:   v2ListGlobalTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "global_tokens_error_invalid_limit",
			target:     v2ListGlobalTokens + "?limit=201",
			template:   v2ListGlobalTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "global_tokens_error_malformed_cursor",
			target:     v2ListGlobalTokens + "?cursor=bad",
			template:   v2ListGlobalTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "global_tokens_error_cross_filter_cursor",
			target:     v2ListGlobalTokens + "?cursor=" + globalTokenCursor(true, "", 5),
			template:   v2ListGlobalTokens,
			pathParams: map[string]string{},
		},
		{
			name:       "global_tokens_error_internal",
			target:     tokensPath,
			template:   v2ListGlobalTokens,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
		{
			name:       "global_token_detail_staked",
			target:     "/api/v2/cosmicgame/cosmic-signature-tokens/5",
			template:   v2GetGlobalToken,
			pathParams: map[string]string{"nftTokenId": "5"},
		},
		{
			name:       "global_token_detail_named_unstaked",
			target:     "/api/v2/cosmicgame/cosmic-signature-tokens/1",
			template:   v2GetGlobalToken,
			pathParams: map[string]string{"nftTokenId": "1"},
		},
		{
			name:       "global_token_detail_error_not_found",
			target:     "/api/v2/cosmicgame/cosmic-signature-tokens/424242",
			template:   v2GetGlobalToken,
			pathParams: map[string]string{"nftTokenId": "424242"},
		},
		{
			name:       "global_token_detail_error_negative",
			target:     "/api/v2/cosmicgame/cosmic-signature-tokens/-1",
			template:   v2GetGlobalToken,
			pathParams: map[string]string{"nftTokenId": "-1"},
		},
		{
			name:       "token_name_history_genesis",
			target:     "/api/v2/cosmicgame/cosmic-signature-tokens/1/name-history",
			template:   v2ListTokenNames,
			pathParams: map[string]string{"nftTokenId": "1"},
		},
		{
			name: "token_name_history_exhausted",
			target: "/api/v2/cosmicgame/cosmic-signature-tokens/1/name-history?cursor=" +
				tokenEventCursor("nameHistory", 1, 5047),
			template:   v2ListTokenNames,
			pathParams: map[string]string{"nftTokenId": "1"},
		},
		{
			name:       "token_name_history_empty_unnamed",
			target:     "/api/v2/cosmicgame/cosmic-signature-tokens/2/name-history",
			template:   v2ListTokenNames,
			pathParams: map[string]string{"nftTokenId": "2"},
		},
		{
			name:       "token_name_history_error_not_found",
			target:     "/api/v2/cosmicgame/cosmic-signature-tokens/424242/name-history",
			template:   v2ListTokenNames,
			pathParams: map[string]string{"nftTokenId": "424242"},
		},
		{
			name:       "token_transfers_first_page",
			target:     transfersPath,
			template:   v2ListTokenTransfers,
			pathParams: map[string]string{"nftTokenId": "2"},
		},
		{
			name:       "token_transfers_next_page",
			target:     transfersPath + "&cursor=" + *transferPage.Meta.NextCursor,
			template:   v2ListTokenTransfers,
			pathParams: map[string]string{"nftTokenId": "2"},
		},
		{
			name: "token_transfers_error_cross_resource_cursor",
			target: transfersPath + "&cursor=" +
				tokenEventCursor("nameHistory", 2, 5047),
			template:   v2ListTokenTransfers,
			pathParams: map[string]string{"nftTokenId": "2"},
		},
		{
			name:       "token_transfers_error_not_found",
			target:     "/api/v2/cosmicgame/cosmic-signature-tokens/424242/transfers",
			template:   v2ListTokenTransfers,
			pathParams: map[string]string{"nftTokenId": "424242"},
		},
		{
			name:       "cs_holders_first_page",
			target:     holdersPath,
			template:   v2ListCsHolders,
			pathParams: map[string]string{},
		},
		{
			name:       "cs_holders_next_page",
			target:     holdersPath + "&cursor=" + *csHolderPage.Meta.NextCursor,
			template:   v2ListCsHolders,
			pathParams: map[string]string{},
		},
		{
			name:       "cs_holders_error_cross_directory_cursor",
			target:     v2ListCsHolders + "?cursor=" + *ctHolderPage.Meta.NextCursor,
			template:   v2ListCsHolders,
			pathParams: map[string]string{},
		},
		{
			name:       "ct_holders_first_page",
			target:     ctHoldersPath,
			template:   v2ListCtHolders,
			pathParams: map[string]string{},
		},
		{
			name:       "ct_holders_next_page",
			target:     ctHoldersPath + "&cursor=" + *ctHolderPage.Meta.NextCursor,
			template:   v2ListCtHolders,
			pathParams: map[string]string{},
		},
		{
			name:       "ct_holders_error_internal",
			target:     ctHoldersPath,
			template:   v2ListCtHolders,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
		{
			name:       "ct_statistics",
			target:     v2GetCtStatistics,
			template:   v2GetCtStatistics,
			pathParams: map[string]string{},
		},
		{
			name:       "ct_statistics_error_internal",
			target:     v2GetCtStatistics,
			template:   v2GetCtStatistics,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
		{
			name:       "supply_by_bid_first_page",
			target:     supplyPath,
			template:   v2ListSupplyByBid,
			pathParams: map[string]string{},
		},
		{
			name:       "supply_by_bid_next_page",
			target:     supplyPath + "&cursor=" + *supplyPage.Meta.NextCursor,
			template:   v2ListSupplyByBid,
			pathParams: map[string]string{},
		},
		{
			name:       "supply_by_bid_exhausted",
			target:     v2ListSupplyByBid + "?cursor=" + terminalSupplyCursor,
			template:   v2ListSupplyByBid,
			pathParams: map[string]string{},
		},
		{
			name:       "supply_by_bid_error_malformed_cursor",
			target:     v2ListSupplyByBid + "?cursor=bad",
			template:   v2ListSupplyByBid,
			pathParams: map[string]string{},
		},
		{
			name:       "supply_by_bid_error_internal",
			target:     supplyPath,
			template:   v2ListSupplyByBid,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
		{
			name:       "supply_daily_fixture_era",
			target:     v2ListSupplyDaily + "?from=2026-01-01&to=2026-01-03",
			template:   v2ListSupplyDaily,
			pathParams: map[string]string{},
		},
		{
			name:       "supply_daily_empty_era",
			target:     v2ListSupplyDaily + "?from=2001-01-01&to=2001-01-03",
			template:   v2ListSupplyDaily,
			pathParams: map[string]string{},
		},
		{
			name:       "supply_daily_error_reversed_window",
			target:     v2ListSupplyDaily + "?from=2026-01-03&to=2026-01-01",
			template:   v2ListSupplyDaily,
			pathParams: map[string]string{},
		},
		{
			name:       "supply_daily_error_oversized_window",
			target:     v2ListSupplyDaily + "?from=2020-01-01&to=2026-01-01",
			template:   v2ListSupplyDaily,
			pathParams: map[string]string{},
		},
		{
			name:       "supply_daily_error_internal",
			target:     v2ListSupplyDaily + "?from=2026-01-01&to=2026-01-03",
			template:   v2ListSupplyDaily,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
		{
			name:       "global_marketing_ledger",
			target:     marketingPath,
			template:   v2ListGlobalMkt,
			pathParams: map[string]string{},
		},
		{
			name:       "global_marketing_exhausted",
			target:     v2ListGlobalMkt + "?cursor=" + terminalMarketingCursor,
			template:   v2ListGlobalMkt,
			pathParams: map[string]string{},
		},
		{
			name: "global_marketing_error_cross_resource_cursor",
			target: v2ListGlobalMkt + "?cursor=" + base64.RawURLEncoding.EncodeToString(
				[]byte(`{"v":1,"s":5017}`)),
			template:   v2ListGlobalMkt,
			pathParams: map[string]string{},
		},
		{
			name:       "global_marketing_error_internal",
			target:     marketingPath,
			template:   v2ListGlobalMkt,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
	}
	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2ContractResources(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "contracts_addresses", target: v2ContractAddresses, template: v2ContractAddresses, pathParams: map[string]string{}},
		{name: "contracts_configuration", target: v2ContractConfig, template: v2ContractConfig, pathParams: map[string]string{}},
		{name: "contracts_balances", target: v2ContractBalances, template: v2ContractBalances, pathParams: map[string]string{}},
		{name: "contracts_current_bid_prices", target: v2CurrentBidPrices, template: v2CurrentBidPrices, pathParams: map[string]string{}},
		{name: "contracts_current_special_winners", target: v2CurrentWinners, template: v2CurrentWinners, pathParams: map[string]string{}},
		{name: "contracts_addresses_error_internal", target: v2ContractAddresses, template: v2ContractAddresses, pathParams: map[string]string{}, ctx: cancelledCtx},
	})

	var marketplaceAddress string
	if err := h.db.QueryRowContext(
		context.Background(),
		"SELECT marketplace_addr FROM rw_contracts LIMIT 1",
	).Scan(&marketplaceAddress); err != nil {
		t.Fatalf("reading marketplace address: %v", err)
	}
	if _, err := h.db.ExecContext(
		context.Background(),
		"UPDATE rw_contracts SET marketplace_addr = ''",
	); err != nil {
		t.Fatalf("clearing marketplace address: %v", err)
	}
	t.Cleanup(func() {
		if _, err := h.db.ExecContext(
			context.Background(),
			"UPDATE rw_contracts SET marketplace_addr = $1",
			marketplaceAddress,
		); err != nil {
			t.Errorf("restoring marketplace address: %v", err)
		}
	})
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "contracts_addresses_error_invalid_registry", target: v2ContractAddresses, template: v2ContractAddresses, pathParams: map[string]string{}},
	})
	if _, err := h.db.ExecContext(
		context.Background(),
		"UPDATE rw_contracts SET marketplace_addr = $1",
		marketplaceAddress,
	); err != nil {
		t.Fatalf("restoring marketplace address: %v", err)
	}

	rpcFailure := func([]any) ([]any, error) {
		return nil, errors.New("forced contract read failure")
	}
	restoreV1 := func() {
		h.gameStub.Return("ethBidPriceIncreaseDivisor", big.NewInt(100))
		h.gameStub.Return("getNextEthBidPrice", wei("1010000000000000"))
		h.gameStub.Return("charityAddress", ethcommon.HexToAddress("0x6000000000000000000000000000000000000006"))
		h.gameStub.Return(
			"tryGetCurrentChampions",
			ethcommon.HexToAddress(addrBob),
			big.NewInt(600),
			ethcommon.HexToAddress(addrCarol),
			big.NewInt(800),
		)
		h.gameStub.Return("cstDutchAuctionDurationDivisor", big.NewInt(11))
		h.gameStub.Return("getCstDutchAuctionDurations", big.NewInt(28800), big.NewInt(3600))
		h.gameStub.Handle("cstDutchAuctionDuration", rpcFailure)
		h.state.LoadInitial(context.Background())
		h.gameStub.Return("cstDutchAuctionDuration", big.NewInt(28800))
	}
	t.Cleanup(restoreV1)

	h.gameStub.Handle("getNextEthBidPrice", rpcFailure)
	h.state.LoadInitial(context.Background())
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "contracts_current_bid_prices_error_unavailable", target: v2CurrentBidPrices, template: v2CurrentBidPrices, pathParams: map[string]string{}},
	})
	h.gameStub.Return("getNextEthBidPrice", wei("1010000000000000"))
	h.state.LoadInitial(context.Background())

	h.gameStub.Handle("ethBidPriceIncreaseDivisor", rpcFailure)
	h.state.LoadInitial(context.Background())
	configUnavailable := h.get(t, v2ContractConfig)
	if configUnavailable.Code != http.StatusServiceUnavailable ||
		configUnavailable.Header().Get("Retry-After") != "300" {
		t.Fatalf("configuration unavailable = %d Retry-After %q",
			configUnavailable.Code, configUnavailable.Header().Get("Retry-After"))
	}
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "contracts_configuration_error_unavailable", target: v2ContractConfig, template: v2ContractConfig, pathParams: map[string]string{}},
	})
	h.gameStub.Return("ethBidPriceIncreaseDivisor", big.NewInt(100))
	h.state.LoadInitial(context.Background())
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "contracts_configuration_recovered", target: v2ContractConfig, template: v2ContractConfig, pathParams: map[string]string{}},
	})

	h.gameStub.Return("charityAddress", ethcommon.Address{})
	h.state.LoadInitial(context.Background())
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "contracts_balances_error_unavailable", target: v2ContractBalances, template: v2ContractBalances, pathParams: map[string]string{}},
	})
	h.gameStub.Return("charityAddress", ethcommon.HexToAddress("0x6000000000000000000000000000000000000006"))
	h.state.LoadInitial(context.Background())
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "contracts_balances_recovered", target: v2ContractBalances, template: v2ContractBalances, pathParams: map[string]string{}},
	})

	h.gameStub.Handle("tryGetCurrentChampions", rpcFailure)
	h.state.LoadInitial(context.Background())
	specialUnavailable := h.get(t, v2CurrentWinners)
	if specialUnavailable.Code != http.StatusServiceUnavailable ||
		specialUnavailable.Header().Get("Retry-After") != "30" {
		t.Fatalf("special-winners unavailable = %d Retry-After %q",
			specialUnavailable.Code, specialUnavailable.Header().Get("Retry-After"))
	}
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "contracts_special_winners_error_unavailable", target: v2CurrentWinners, template: v2CurrentWinners, pathParams: map[string]string{}},
	})
	h.gameStub.Return(
		"tryGetCurrentChampions",
		ethcommon.HexToAddress(addrBob),
		big.NewInt(600),
		ethcommon.HexToAddress(addrCarol),
		big.NewInt(800),
	)
	h.state.LoadInitial(context.Background())
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "contracts_special_winners_recovered", target: v2CurrentWinners, template: v2CurrentWinners, pathParams: map[string]string{}},
	})

	h.gameStub.Handle("cstDutchAuctionDurationDivisor", rpcFailure)
	h.gameStub.Return(
		"getCstDutchAuctionDurations",
		big.NewInt(28800),
		big.NewInt(1767229000),
	)
	h.state.LoadInitial(context.Background())
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "contracts_configuration_v2", target: v2ContractConfig, template: v2ContractConfig, pathParams: map[string]string{}},
		{name: "contracts_current_bid_prices_v2", target: v2CurrentBidPrices, template: v2CurrentBidPrices, pathParams: map[string]string{}},
	})
	restoreV1()
}

func TestAPIV2ContractResourcesMatchV1Semantics(t *testing.T) {
	h := server(t)
	var dashboard map[string]any
	decodeV2JSON(t, h.get(t, "/api/cosmicgame/statistics/dashboard"), &dashboard)
	var addresses apiv2.ContractAddressRegistry
	decodeV2JSON(t, h.get(t, v2ContractAddresses), &addresses)
	legacyAddresses := dashboard["ContractAddrs"].(map[string]any)
	addressPairs := map[string]string{
		"CosmicGameAddr":         addresses.CosmicGameAddress,
		"CosmicSignatureAddr":    addresses.CosmicSignatureAddress,
		"CosmicTokenAddr":        addresses.CosmicTokenAddress,
		"CosmicDaoAddr":          addresses.CosmicDaoAddress,
		"CharityWalletAddr":      addresses.CharityWalletAddress,
		"PrizesWalletAddr":       addresses.PrizesWalletAddress,
		"RandomWalkAddr":         addresses.RandomWalkAddress,
		"StakingWalletCSTAddr":   addresses.CstStakingWalletAddress,
		"StakingWalletRWalkAddr": addresses.RandomWalkStakingWalletAddress,
		"MarketingWalletAddr":    addresses.MarketingWalletAddress,
		"MarketplaceAddr":        addresses.MarketplaceAddress,
		"ImplementationAddr":     addresses.ImplementationAddress,
	}
	for legacyField, v2Value := range addressPairs {
		if v2Value != legacyAddresses[legacyField] {
			t.Fatalf("%s diverged: v1=%v v2=%s", legacyField, legacyAddresses[legacyField], v2Value)
		}
	}

	var configuration apiv2.ContractConfiguration
	decodeV2JSON(t, h.get(t, v2ContractConfig), &configuration)
	if configuration.EthBidPriceIncreaseDivisor != dashboard["PriceIncrease"] ||
		configuration.CstBidRewardMode != apiv2.CstBidRewardFixed ||
		configuration.FixedCstBidRewardWei == nil ||
		*configuration.FixedCstBidRewardWei != dashboard["TokenReward"] ||
		configuration.CharityDonationPercentage != int64(dashboard["CharityPercentage"].(float64)) ||
		configuration.ChronoWarriorPercentage != int64(dashboard["ChronoWarriorPercentage"].(float64)) ||
		configuration.MainPrizePercentage != int64(dashboard["PrizePercentage"].(float64)) ||
		configuration.RaffleEthWinnerCount != int64(dashboard["NumRaffleEthWinnersBidding"].(float64)) ||
		configuration.RaffleNftBidderWinnerCount != int64(dashboard["NumRaffleNFTWinnersBidding"].(float64)) ||
		configuration.RaffleNftRandomWalkStakerWinnerCount != int64(dashboard["NumRaffleNFTWinnersStakingRWalk"].(float64)) {
		t.Fatalf("configuration split diverged: dashboard=%+v v2=%+v", dashboard, configuration)
	}
	var balances apiv2.ContractBalances
	decodeV2JSON(t, h.get(t, v2ContractBalances), &balances)
	if balances.CharityAddress != dashboard["CharityAddr"] ||
		balances.CharityBalanceWei != dashboard["CharityBalance"] ||
		balances.CosmicGameBalanceWei != "0" {
		t.Fatalf("balance split diverged: dashboard=%+v v2=%+v", dashboard, balances)
	}

	var cstPrice, ethPrice map[string]any
	decodeV2JSON(t, h.get(t, "/api/cosmicgame/bid/cst_price"), &cstPrice)
	decodeV2JSON(t, h.get(t, "/api/cosmicgame/bid/eth_price"), &ethPrice)
	var prices apiv2.CurrentBidPrices
	decodeV2JSON(t, h.get(t, v2CurrentBidPrices), &prices)
	if prices.NextCstBidPriceWei != cstPrice["CSTPrice"] ||
		prices.NextEthBidPriceWei != ethPrice["ETHPrice"] ||
		prices.NextCstBidRewardWei != dashboard["TokenReward"] ||
		strconv.FormatInt(prices.CstAuctionDurationSeconds, 10) != cstPrice["AuctionDuration"] ||
		strconv.FormatInt(prices.CstAuctionElapsedSeconds, 10) != cstPrice["SecondsElapsed"] ||
		strconv.FormatInt(prices.EthAuctionDurationSeconds, 10) != ethPrice["AuctionDuration"] ||
		strconv.FormatInt(prices.EthAuctionElapsedSeconds, 10) != ethPrice["SecondsElapsed"] {
		t.Fatalf("bid-price split diverged: cst=%+v eth=%+v v2=%+v", cstPrice, ethPrice, prices)
	}

	var legacyWinners map[string]any
	decodeV2JSON(t, h.get(t, "/api/cosmicgame/bid/current_special_winners"), &legacyWinners)
	var winners apiv2.CurrentSpecialWinners
	decodeV2JSON(t, h.get(t, v2CurrentWinners), &winners)
	if winners.EnduranceChampion == nil || winners.ChronoWarrior == nil ||
		winners.LastBidder == nil || winners.LastCstBidder == nil ||
		winners.EnduranceChampion.Address != legacyWinners["EnduranceChampionAddress"] ||
		winners.EnduranceChampion.DurationSeconds != int64(legacyWinners["EnduranceChampionDuration"].(float64)) ||
		winners.EnduranceChampion.StartedAt != int64(legacyWinners["EnduranceChampionStartTimeStamp"].(float64)) ||
		winners.EnduranceChampion.PreviousDurationSeconds != int64(legacyWinners["PrevEnduranceChampionDuration"].(float64)) ||
		winners.ChronoWarrior.Address != legacyWinners["ChronoWarriorAddress"] ||
		winners.ChronoWarrior.DurationSeconds != int64(legacyWinners["ChronoWarriorDuration"].(float64)) ||
		winners.ChronoWarrior.IsLive != legacyWinners["ChronoWarriorIsLive"] ||
		winners.LastBidder.Address != legacyWinners["LastBidderAddress"] ||
		winners.LastBidder.LastBidAt != int64(legacyWinners["LastBidderLastBidTime"].(float64)) ||
		winners.LastCstBidder.Address != legacyWinners["LastCstBidderAddress"] ||
		winners.Round != int64(legacyWinners["RoundNum"].(float64)) ||
		winners.SourceBlockNumber != int64(legacyWinners["SourceBlockNumber"].(float64)) ||
		winners.SourceBlockTimestamp != int64(legacyWinners["SourceBlockTimeStamp"].(float64)) {
		t.Fatalf("special-winner split diverged: v1=%+v v2=%+v", legacyWinners, winners)
	}
	// V2 intentionally omits the legacy epoch-zero CST timestamp sentinel.
	if winners.LastCstBidder.LastBidAt != nil ||
		legacyWinners["LastCstBidderLastBidTime"].(float64) != 0 {
		t.Fatalf("last-CST sentinel semantics changed: v1=%+v v2=%+v",
			legacyWinners["LastCstBidderLastBidTime"], winners.LastCstBidder)
	}
}

func TestAPIV2BiddingAnalytics(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	var originalRoundStart time.Time
	if err := h.db.QueryRowContext(
		context.Background(),
		"SELECT round_start_time FROM cg_round_stats WHERE round_num = 0",
	).Scan(&originalRoundStart); err != nil {
		t.Fatalf("reading fixture round start: %v", err)
	}
	if _, err := h.db.ExecContext(
		context.Background(),
		"UPDATE cg_round_stats SET round_start_time = TO_TIMESTAMP($1) WHERE round_num = 0",
		1767222000,
	); err != nil {
		t.Fatalf("moving fixture round start: %v", err)
	}
	const syntheticBidID = 9_000_001
	t.Cleanup(func() {
		if _, err := h.db.ExecContext(
			context.Background(),
			"DELETE FROM cg_bid WHERE id = $1",
			syntheticBidID,
		); err != nil {
			t.Errorf("deleting synthetic analytics bid: %v", err)
		}
		if _, err := h.db.ExecContext(
			context.Background(),
			"UPDATE cg_round_stats SET round_start_time = $1 WHERE round_num = 0",
			originalRoundStart,
		); err != nil {
			t.Errorf("restoring fixture round start: %v", err)
		}
	})
	if _, err := h.db.ExecContext(
		context.Background(),
		`INSERT INTO cg_bid(
			id, evtlog_id, block_num, tx_id, time_stamp, contract_aid, bidder_aid,
			rwalk_nft_id, round_num, bid_type, bid_position, prize_time,
			eth_price, cst_price, cst_reward, bid_cst_reward_amount,
			cst_dutch_auction_duration, msg
		) VALUES (
			$1, NULL, 105, 1006, TO_TIMESTAMP(1767226100), 2, 24,
			-1, 0, 0, 5, TO_TIMESTAMP(1767229700),
			130000000000000000, -1, 100000000000000000000, -1, -1,
			'analytics spike fixture'
		)`,
		syntheticBidID,
	); err != nil {
		t.Fatalf("inserting synthetic analytics bid: %v", err)
	}

	const window = "?from=1767225600&to=1767230000&intervalSeconds=600"
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []v2GoldenCase{
		{name: "bidding_activity", target: v2BiddingActivity + window, template: v2BiddingActivity, pathParams: map[string]string{}},
		{name: "bidding_activity_default_interval", target: v2BiddingActivity + "?from=1767225600&to=1767230000", template: v2BiddingActivity, pathParams: map[string]string{}},
		{name: "bidding_frequency", target: v2BiddingFrequency + window, template: v2BiddingFrequency, pathParams: map[string]string{}},
		{name: "bidding_frequency_default_interval", target: v2BiddingFrequency + "?from=1767225600&to=1767230000", template: v2BiddingFrequency, pathParams: map[string]string{}},
		{name: "bidding_frequency_exact_boundary", target: v2BiddingFrequency + "?from=1767225600&to=1767229200&intervalSeconds=600", template: v2BiddingFrequency, pathParams: map[string]string{}},
		{name: "bidding_frequency_empty_window", target: v2BiddingFrequency + "?from=100&to=200&intervalSeconds=60", template: v2BiddingFrequency, pathParams: map[string]string{}},
		{name: "bidding_type_ratio", target: v2BiddingTypeRatio + window, template: v2BiddingTypeRatio, pathParams: map[string]string{}},
		{name: "bidding_top_active_periods", target: v2BiddingTopPeriods + "?from=1767225600&to=1767230000&top=3&gapHours=1&minBids=1", template: v2BiddingTopPeriods, pathParams: map[string]string{}},
		{name: "bidding_time_bounds", target: v2BiddingTimeBounds, template: v2BiddingTimeBounds, pathParams: map[string]string{}},
		{name: "bidding_error_missing_from", target: v2BiddingFrequency + "?to=2", template: v2BiddingFrequency, pathParams: map[string]string{}},
		{name: "bidding_error_inverted_range", target: v2BiddingFrequency + "?from=2&to=2", template: v2BiddingFrequency, pathParams: map[string]string{}},
		{name: "bidding_error_window_limit", target: v2BiddingFrequency + "?from=0&to=158112001", template: v2BiddingFrequency, pathParams: map[string]string{}},
		{name: "bidding_error_timestamp_limit", target: v2BiddingFrequency + "?from=253402300799&to=253402300800", template: v2BiddingFrequency, pathParams: map[string]string{}},
		{name: "bidding_error_bucket_limit", target: v2BiddingFrequency + "?from=0&to=2001&intervalSeconds=1", template: v2BiddingFrequency, pathParams: map[string]string{}},
		{name: "bidding_error_bind_interval", target: v2BiddingFrequency + "?from=0&to=2&intervalSeconds=bad", template: v2BiddingFrequency, pathParams: map[string]string{}},
		{name: "bidding_error_top_limit", target: v2BiddingTopPeriods + "?from=0&to=2&top=101", template: v2BiddingTopPeriods, pathParams: map[string]string{}},
		{name: "bidding_activity_error_internal", target: v2BiddingActivity + window, template: v2BiddingActivity, pathParams: map[string]string{}, ctx: cancelledCtx},
		{name: "bidding_frequency_error_internal", target: v2BiddingFrequency + window, template: v2BiddingFrequency, pathParams: map[string]string{}, ctx: cancelledCtx},
		{name: "bidding_type_ratio_error_internal", target: v2BiddingTypeRatio + window, template: v2BiddingTypeRatio, pathParams: map[string]string{}, ctx: cancelledCtx},
		{name: "bidding_top_active_periods_error_internal", target: v2BiddingTopPeriods + "?from=1767225600&to=1767230000&top=3&gapHours=1&minBids=1", template: v2BiddingTopPeriods, pathParams: map[string]string{}, ctx: cancelledCtx},
		{name: "bidding_time_bounds_error_internal", target: v2BiddingTimeBounds, template: v2BiddingTimeBounds, pathParams: map[string]string{}, ctx: cancelledCtx},
	}
	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2RoundPrizes(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	firstPath := "/api/v2/cosmicgame/rounds/0/prizes?limit=5"
	firstResponse := h.get(t, firstPath)
	if firstResponse.Code != http.StatusOK {
		t.Fatalf("first page: status=%d body=%s", firstResponse.Code, firstResponse.Body.String())
	}
	var firstPage apiv2.RoundPrizePage
	if err := json.Unmarshal(firstResponse.Body.Bytes(), &firstPage); err != nil {
		t.Fatalf("decoding first page: %v", err)
	}
	if firstPage.Meta.NextCursor == nil {
		t.Fatal("fixture first page did not return a continuation cursor")
	}

	afterChrono := base64.RawURLEncoding.EncodeToString([]byte(`{"v":1,"r":0,"t":9,"w":2}`))
	rafflePath := "/api/v2/cosmicgame/rounds/0/prizes?limit=5&cursor=" + afterChrono
	raffleResponse := h.get(t, rafflePath)
	if raffleResponse.Code != http.StatusOK {
		t.Fatalf("raffle page: status=%d body=%s", raffleResponse.Code, raffleResponse.Body.String())
	}
	var rafflePage apiv2.RoundPrizePage
	if err := json.Unmarshal(raffleResponse.Body.Bytes(), &rafflePage); err != nil {
		t.Fatalf("decoding raffle page: %v", err)
	}
	if rafflePage.Meta.NextCursor == nil {
		t.Fatal("fixture raffle page did not return a continuation cursor")
	}
	afterLastPrize := base64.RawURLEncoding.EncodeToString([]byte(`{"v":1,"r":0,"t":15,"w":0}`))
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []v2GoldenCase{
		{
			name:       "prizes_list_first_page",
			target:     firstPath,
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "prizes_list_next_page",
			target:     "/api/v2/cosmicgame/rounds/0/prizes?limit=5&cursor=" + *firstPage.Meta.NextCursor,
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "prizes_list_raffle_page",
			target:     rafflePath,
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "prizes_list_raffle_next_page",
			target:     "/api/v2/cosmicgame/rounds/0/prizes?limit=5&cursor=" + *rafflePage.Meta.NextCursor,
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "prizes_list_empty_page",
			target:     "/api/v2/cosmicgame/rounds/0/prizes?limit=5&cursor=" + afterLastPrize,
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "prizes_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/prizes?cursor=not-a-cursor",
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "prizes_error_invalid_limit",
			target:     "/api/v2/cosmicgame/rounds/0/prizes?limit=201",
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "prizes_error_bind_limit",
			target:     "/api/v2/cosmicgame/rounds/0/prizes?limit=wat",
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "prizes_error_cross_round_cursor",
			target:     "/api/v2/cosmicgame/rounds/1/prizes?cursor=" + *firstPage.Meta.NextCursor,
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "1"},
		},
		{
			name:       "prizes_error_open_round",
			target:     "/api/v2/cosmicgame/rounds/3/prizes",
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "3"},
		},
		{
			name:       "prizes_error_not_found",
			target:     "/api/v2/cosmicgame/rounds/999/prizes",
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "999"},
		},
		{
			name:       "prizes_error_internal",
			target:     firstPath,
			template:   v2ListRoundPrizes,
			pathParams: map[string]string{"round": "0"},
			ctx:        cancelledCtx,
		},
	}

	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2RoundRaffles(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	firstEthPath := "/api/v2/cosmicgame/rounds/0/raffle-eth-deposits?limit=1"
	firstEthResponse := h.get(t, firstEthPath)
	if firstEthResponse.Code != http.StatusOK {
		t.Fatalf("first ETH page: status=%d body=%s", firstEthResponse.Code, firstEthResponse.Body.String())
	}
	var firstEthPage apiv2.RoundRaffleEthDepositPage
	if err := json.Unmarshal(firstEthResponse.Body.Bytes(), &firstEthPage); err != nil {
		t.Fatalf("decoding first ETH page: %v", err)
	}
	if firstEthPage.Meta.NextCursor == nil {
		t.Fatal("fixture first ETH page did not return a continuation cursor")
	}

	afterLastEth := base64.RawURLEncoding.EncodeToString([]byte(`{"v":1,"r":0,"w":1,"e":5024}`))
	bidderNftCursor := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"v":1,"r":0,"p":"bidder","w":0,"e":5025}`),
	)
	stakerNftCursor := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"v":1,"r":0,"p":"randomWalkStaker","w":0,"e":5028}`),
	)
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	cases := []v2GoldenCase{
		{
			name:       "raffle_eth_list_first_page",
			target:     firstEthPath,
			template:   v2ListRaffleEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_eth_list_next_page",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-eth-deposits?limit=1&cursor=" + *firstEthPage.Meta.NextCursor,
			template:   v2ListRaffleEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_eth_list_empty_page",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-eth-deposits?limit=1&cursor=" + afterLastEth,
			template:   v2ListRaffleEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_eth_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-eth-deposits?cursor=bad",
			template:   v2ListRaffleEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_eth_error_cross_round_cursor",
			target:     "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits?cursor=" + *firstEthPage.Meta.NextCursor,
			template:   v2ListRaffleEth,
			pathParams: map[string]string{"round": "1"},
		},
		{
			name:       "raffle_eth_error_invalid_limit",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-eth-deposits?limit=201",
			template:   v2ListRaffleEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_eth_error_bind_limit",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-eth-deposits?limit=wat",
			template:   v2ListRaffleEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_eth_error_open_round",
			target:     "/api/v2/cosmicgame/rounds/3/raffle-eth-deposits",
			template:   v2ListRaffleEth,
			pathParams: map[string]string{"round": "3"},
		},
		{
			name:       "raffle_eth_error_not_found",
			target:     "/api/v2/cosmicgame/rounds/999/raffle-eth-deposits",
			template:   v2ListRaffleEth,
			pathParams: map[string]string{"round": "999"},
		},
		{
			name:       "raffle_eth_error_internal",
			target:     firstEthPath,
			template:   v2ListRaffleEth,
			pathParams: map[string]string{"round": "0"},
			ctx:        cancelledCtx,
		},
		{
			name:       "raffle_nft_bidder_page",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners?pool=bidder",
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_nft_staker_page",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners?pool=randomWalkStaker",
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_nft_empty_page",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners?pool=bidder&cursor=" + bidderNftCursor,
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_nft_staker_empty_page",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners?pool=randomWalkStaker&cursor=" + stakerNftCursor,
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_nft_error_missing_pool",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners",
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_nft_error_invalid_pool",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners?pool=other",
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_nft_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners?pool=bidder&cursor=bad",
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_nft_error_cross_pool_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners?pool=bidder&cursor=" + stakerNftCursor,
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_nft_error_cross_round_cursor",
			target:     "/api/v2/cosmicgame/rounds/1/raffle-nft-winners?pool=bidder&cursor=" + bidderNftCursor,
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "1"},
		},
		{
			name:       "raffle_nft_error_invalid_limit",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners?pool=bidder&limit=201",
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_nft_error_bind_limit",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners?pool=bidder&limit=wat",
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "raffle_nft_error_open_round",
			target:     "/api/v2/cosmicgame/rounds/3/raffle-nft-winners?pool=bidder",
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "3"},
		},
		{
			name:       "raffle_nft_error_not_found",
			target:     "/api/v2/cosmicgame/rounds/999/raffle-nft-winners?pool=bidder",
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "999"},
		},
		{
			name:       "raffle_nft_error_internal",
			target:     "/api/v2/cosmicgame/rounds/0/raffle-nft-winners?pool=bidder",
			template:   v2ListRaffleNft,
			pathParams: map[string]string{"round": "0"},
			ctx:        cancelledCtx,
		},
	}

	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2RoundDonations(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	firstEthPath := "/api/v2/cosmicgame/rounds/0/eth-donations?limit=1"
	firstEthResponse := h.get(t, firstEthPath)
	if firstEthResponse.Code != http.StatusOK {
		t.Fatalf("first ETH page: status=%d body=%s", firstEthResponse.Code, firstEthResponse.Body.String())
	}
	var firstEthPage apiv2.RoundEthDonationPage
	if err := json.Unmarshal(firstEthResponse.Body.Bytes(), &firstEthPage); err != nil {
		t.Fatalf("decoding first ETH page: %v", err)
	}
	if firstEthPage.Meta.NextCursor == nil {
		t.Fatal("fixture first ETH donation page did not return a continuation cursor")
	}

	afterLastEth := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"v":1,"r":0,"k":"eth","e":5012}`),
	)
	afterLastERC20 := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"v":1,"r":0,"k":"erc20","e":5015}`),
	)
	afterLastNFT := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"v":1,"r":0,"k":"nft","e":5016}`),
	)
	crossRoundETH := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"v":1,"r":1,"k":"eth","e":5012}`),
	)
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	cases := []v2GoldenCase{
		{
			name:       "donation_eth_list_first_page",
			target:     firstEthPath,
			template:   v2ListDonationEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_eth_list_next_page",
			target:     "/api/v2/cosmicgame/rounds/0/eth-donations?limit=1&cursor=" + *firstEthPage.Meta.NextCursor,
			template:   v2ListDonationEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_eth_list_empty_page",
			target:     "/api/v2/cosmicgame/rounds/0/eth-donations?limit=1&cursor=" + afterLastEth,
			template:   v2ListDonationEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_eth_list_open_round",
			target:     "/api/v2/cosmicgame/rounds/3/eth-donations",
			template:   v2ListDonationEth,
			pathParams: map[string]string{"round": "3"},
		},
		{
			name:       "donation_eth_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/eth-donations?cursor=bad",
			template:   v2ListDonationEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_eth_error_cross_round_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/eth-donations?cursor=" + crossRoundETH,
			template:   v2ListDonationEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_eth_error_cross_resource_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/eth-donations?cursor=" + afterLastNFT,
			template:   v2ListDonationEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_eth_error_invalid_limit",
			target:     "/api/v2/cosmicgame/rounds/0/eth-donations?limit=201",
			template:   v2ListDonationEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_eth_error_bind_limit",
			target:     "/api/v2/cosmicgame/rounds/0/eth-donations?limit=wat",
			template:   v2ListDonationEth,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_eth_error_internal",
			target:     firstEthPath,
			template:   v2ListDonationEth,
			pathParams: map[string]string{"round": "0"},
			ctx:        cancelledCtx,
		},
		{
			name:       "donation_erc20_list_page",
			target:     "/api/v2/cosmicgame/rounds/0/erc20-donations?limit=1",
			template:   v2ListDonationERC20,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_erc20_list_empty_page",
			target:     "/api/v2/cosmicgame/rounds/0/erc20-donations?limit=1&cursor=" + afterLastERC20,
			template:   v2ListDonationERC20,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_erc20_list_open_round",
			target:     "/api/v2/cosmicgame/rounds/3/erc20-donations",
			template:   v2ListDonationERC20,
			pathParams: map[string]string{"round": "3"},
		},
		{
			name:       "donation_erc20_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/erc20-donations?cursor=bad",
			template:   v2ListDonationERC20,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_erc20_error_cross_round_cursor",
			target:     "/api/v2/cosmicgame/rounds/1/erc20-donations?cursor=" + afterLastERC20,
			template:   v2ListDonationERC20,
			pathParams: map[string]string{"round": "1"},
		},
		{
			name:       "donation_erc20_error_cross_resource_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/erc20-donations?cursor=" + afterLastNFT,
			template:   v2ListDonationERC20,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_erc20_error_invalid_limit",
			target:     "/api/v2/cosmicgame/rounds/0/erc20-donations?limit=201",
			template:   v2ListDonationERC20,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_erc20_error_bind_limit",
			target:     "/api/v2/cosmicgame/rounds/0/erc20-donations?limit=wat",
			template:   v2ListDonationERC20,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_erc20_error_internal",
			target:     "/api/v2/cosmicgame/rounds/0/erc20-donations",
			template:   v2ListDonationERC20,
			pathParams: map[string]string{"round": "0"},
			ctx:        cancelledCtx,
		},
		{
			name:       "donation_nft_list_page",
			target:     "/api/v2/cosmicgame/rounds/0/nft-donations?limit=1",
			template:   v2ListDonationNFT,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_nft_list_empty_page",
			target:     "/api/v2/cosmicgame/rounds/0/nft-donations?limit=1&cursor=" + afterLastNFT,
			template:   v2ListDonationNFT,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_nft_list_open_round",
			target:     "/api/v2/cosmicgame/rounds/3/nft-donations",
			template:   v2ListDonationNFT,
			pathParams: map[string]string{"round": "3"},
		},
		{
			name:       "donation_nft_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/nft-donations?cursor=bad",
			template:   v2ListDonationNFT,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_nft_error_cross_round_cursor",
			target:     "/api/v2/cosmicgame/rounds/1/nft-donations?cursor=" + afterLastNFT,
			template:   v2ListDonationNFT,
			pathParams: map[string]string{"round": "1"},
		},
		{
			name:       "donation_nft_error_cross_resource_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/nft-donations?cursor=" + afterLastERC20,
			template:   v2ListDonationNFT,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_nft_error_invalid_limit",
			target:     "/api/v2/cosmicgame/rounds/0/nft-donations?limit=201",
			template:   v2ListDonationNFT,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_nft_error_bind_limit",
			target:     "/api/v2/cosmicgame/rounds/0/nft-donations?limit=wat",
			template:   v2ListDonationNFT,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "donation_nft_error_internal",
			target:     "/api/v2/cosmicgame/rounds/0/nft-donations",
			template:   v2ListDonationNFT,
			pathParams: map[string]string{"round": "0"},
			ctx:        cancelledCtx,
		},
	}

	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2StatisticsAndClaims(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	firstROIPath := "/api/v2/cosmicgame/statistics/leaderboard/roi?minBids=0&limit=2"
	firstROIResponse := h.get(t, firstROIPath)
	if firstROIResponse.Code != http.StatusOK {
		t.Fatalf("first ROI page: status=%d body=%s", firstROIResponse.Code, firstROIResponse.Body.String())
	}
	var firstROIPage apiv2.RoiLeaderboardPage
	if err := json.Unmarshal(firstROIResponse.Body.Bytes(), &firstROIPage); err != nil {
		t.Fatalf("decoding first ROI page: %v", err)
	}
	if firstROIPage.Meta.NextCursor == nil {
		t.Fatal("fixture first ROI page did not return a continuation cursor")
	}
	secondROIPath := firstROIPath + "&cursor=" + *firstROIPage.Meta.NextCursor
	secondROIResponse := h.get(t, secondROIPath)
	if secondROIResponse.Code != http.StatusOK {
		t.Fatalf("second ROI page: status=%d body=%s", secondROIResponse.Code, secondROIResponse.Body.String())
	}
	var secondROIPage apiv2.RoiLeaderboardPage
	if err := json.Unmarshal(secondROIResponse.Body.Bytes(), &secondROIPage); err != nil {
		t.Fatalf("decoding second ROI page: %v", err)
	}
	if secondROIPage.Meta.NextCursor == nil {
		t.Fatal("fixture second ROI page did not return a continuation cursor")
	}
	lastROICursor := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"v":1,"s":"netProfit","m":0,"k":"-155000000000000000","x":0,"a":22}`),
	)

	firstClaimsPath := "/api/v2/cosmicgame/statistics/claims?limit=1"
	firstClaimsResponse := h.get(t, firstClaimsPath)
	if firstClaimsResponse.Code != http.StatusOK {
		t.Fatalf("first claims page: status=%d body=%s", firstClaimsResponse.Code, firstClaimsResponse.Body.String())
	}
	var firstClaimsPage apiv2.ClaimSummaryPage
	if err := json.Unmarshal(firstClaimsResponse.Body.Bytes(), &firstClaimsPage); err != nil {
		t.Fatalf("decoding first claims page: %v", err)
	}
	if firstClaimsPage.Meta.NextCursor == nil {
		t.Fatal("fixture first claims page did not return a continuation cursor")
	}
	afterLastClaim := base64.RawURLEncoding.EncodeToString(
		[]byte(`{"v":1,"r":0,"e":5018}`),
	)

	firstDetailPath := "/api/v2/cosmicgame/rounds/0/claims?limit=1"
	firstDetailResponse := h.get(t, firstDetailPath)
	if firstDetailResponse.Code != http.StatusOK {
		t.Fatalf("first claim detail: status=%d body=%s", firstDetailResponse.Code, firstDetailResponse.Body.String())
	}
	var firstDetail apiv2.RoundClaimsDetail
	if err := json.Unmarshal(firstDetailResponse.Body.Bytes(), &firstDetail); err != nil {
		t.Fatalf("decoding first claim detail: %v", err)
	}
	if firstDetail.ClaimTransactions.Meta.NextCursor == nil ||
		firstDetail.AttachedTokens.Meta.NextCursor == nil ||
		firstDetail.UnclaimedItems.Meta.NextCursor == nil {
		t.Fatal("fixture claim detail did not return all section cursors")
	}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []v2GoldenCase{
		{name: "statistics_global_get", target: v2GlobalStatistics, template: v2GlobalStatistics, pathParams: map[string]string{}},
		{name: "statistics_global_error_internal", target: v2GlobalStatistics, template: v2GlobalStatistics, pathParams: map[string]string{}, ctx: cancelledCtx},
		{name: "statistics_counters_get", target: v2StatisticsCounters, template: v2StatisticsCounters, pathParams: map[string]string{}},
		{name: "statistics_counters_error_internal", target: v2StatisticsCounters, template: v2StatisticsCounters, pathParams: map[string]string{}, ctx: cancelledCtx},
		{name: "statistics_roi_first_page", target: firstROIPath, template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_next_page", target: secondROIPath, template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_final_page", target: firstROIPath + "&cursor=" + *secondROIPage.Meta.NextCursor, template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_empty_page", target: firstROIPath + "&cursor=" + lastROICursor, template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_sort_roi", target: "/api/v2/cosmicgame/statistics/leaderboard/roi?minBids=0&sort=roi&limit=2", template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_sort_win_rate", target: "/api/v2/cosmicgame/statistics/leaderboard/roi?minBids=0&sort=winRate&limit=2", template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_sort_spent", target: "/api/v2/cosmicgame/statistics/leaderboard/roi?minBids=0&sort=spent&limit=2", template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_sort_nfts", target: "/api/v2/cosmicgame/statistics/leaderboard/roi?minBids=0&sort=nfts&limit=2", template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_sort_bids", target: "/api/v2/cosmicgame/statistics/leaderboard/roi?minBids=0&sort=bids&limit=2", template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_error_malformed_cursor", target: "/api/v2/cosmicgame/statistics/leaderboard/roi?cursor=bad", template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_error_cross_sort_cursor", target: "/api/v2/cosmicgame/statistics/leaderboard/roi?minBids=0&sort=roi&cursor=" + *firstROIPage.Meta.NextCursor, template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_error_cross_filter_cursor", target: "/api/v2/cosmicgame/statistics/leaderboard/roi?minBids=1&cursor=" + *firstROIPage.Meta.NextCursor, template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_error_invalid_sort", target: "/api/v2/cosmicgame/statistics/leaderboard/roi?sort=other", template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_error_invalid_limit", target: "/api/v2/cosmicgame/statistics/leaderboard/roi?limit=201", template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_error_bind_limit", target: "/api/v2/cosmicgame/statistics/leaderboard/roi?limit=bad", template: v2StatisticsROI, pathParams: map[string]string{}},
		{name: "statistics_roi_error_internal", target: firstROIPath, template: v2StatisticsROI, pathParams: map[string]string{}, ctx: cancelledCtx},
		{name: "statistics_claims_first_page", target: firstClaimsPath, template: v2StatisticsClaims, pathParams: map[string]string{}},
		{name: "statistics_claims_next_page", target: firstClaimsPath + "&cursor=" + *firstClaimsPage.Meta.NextCursor, template: v2StatisticsClaims, pathParams: map[string]string{}},
		{name: "statistics_claims_empty_page", target: firstClaimsPath + "&cursor=" + afterLastClaim, template: v2StatisticsClaims, pathParams: map[string]string{}},
		{name: "statistics_claims_error_malformed_cursor", target: "/api/v2/cosmicgame/statistics/claims?cursor=bad", template: v2StatisticsClaims, pathParams: map[string]string{}},
		{name: "statistics_claims_error_invalid_limit", target: "/api/v2/cosmicgame/statistics/claims?limit=201", template: v2StatisticsClaims, pathParams: map[string]string{}},
		{name: "statistics_claims_error_bind_limit", target: "/api/v2/cosmicgame/statistics/claims?limit=bad", template: v2StatisticsClaims, pathParams: map[string]string{}},
		{name: "statistics_claims_error_internal", target: firstClaimsPath, template: v2StatisticsClaims, pathParams: map[string]string{}, ctx: cancelledCtx},
		{name: "round_claims_detail", target: "/api/v2/cosmicgame/rounds/0/claims", template: v2RoundClaims, pathParams: map[string]string{"round": "0"}},
		{name: "round_claims_detail_first_pages", target: firstDetailPath, template: v2RoundClaims, pathParams: map[string]string{"round": "0"}},
		{name: "round_claims_detail_next_pages", target: firstDetailPath + "&claimTransactionsCursor=" + *firstDetail.ClaimTransactions.Meta.NextCursor + "&attachedTokensCursor=" + *firstDetail.AttachedTokens.Meta.NextCursor + "&unclaimedItemsCursor=" + *firstDetail.UnclaimedItems.Meta.NextCursor, template: v2RoundClaims, pathParams: map[string]string{"round": "0"}},
		{name: "round_claims_detail_round_1", target: "/api/v2/cosmicgame/rounds/1/claims", template: v2RoundClaims, pathParams: map[string]string{"round": "1"}},
		{name: "round_claims_error_open_round", target: "/api/v2/cosmicgame/rounds/3/claims", template: v2RoundClaims, pathParams: map[string]string{"round": "3"}},
		{name: "round_claims_error_not_found", target: "/api/v2/cosmicgame/rounds/999/claims", template: v2RoundClaims, pathParams: map[string]string{"round": "999"}},
		{name: "round_claims_error_malformed_cursor", target: "/api/v2/cosmicgame/rounds/0/claims?claimTransactionsCursor=bad", template: v2RoundClaims, pathParams: map[string]string{"round": "0"}},
		{name: "round_claims_error_cross_section_cursor", target: "/api/v2/cosmicgame/rounds/0/claims?attachedTokensCursor=" + *firstDetail.ClaimTransactions.Meta.NextCursor, template: v2RoundClaims, pathParams: map[string]string{"round": "0"}},
		{name: "round_claims_error_cross_round_cursor", target: "/api/v2/cosmicgame/rounds/1/claims?claimTransactionsCursor=" + *firstDetail.ClaimTransactions.Meta.NextCursor, template: v2RoundClaims, pathParams: map[string]string{"round": "1"}},
		{name: "round_claims_error_invalid_limit", target: "/api/v2/cosmicgame/rounds/0/claims?limit=201", template: v2RoundClaims, pathParams: map[string]string{"round": "0"}},
		{name: "round_claims_error_bind_limit", target: "/api/v2/cosmicgame/rounds/0/claims?limit=bad", template: v2RoundClaims, pathParams: map[string]string{"round": "0"}},
		{name: "round_claims_error_internal", target: "/api/v2/cosmicgame/rounds/0/claims", template: v2RoundClaims, pathParams: map[string]string{"round": "0"}, ctx: cancelledCtx},
	}
	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2ParticipantDirectories(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	nextCursor := func(path string) string {
		response := h.get(t, path)
		if response.Code != http.StatusOK {
			t.Fatalf("%s: status=%d body=%s", path, response.Code, response.Body.String())
		}
		var page struct {
			Meta apiv2.PageMeta `json:"meta"`
		}
		if err := json.Unmarshal(response.Body.Bytes(), &page); err != nil {
			t.Fatalf("decoding %s: %v", path, err)
		}
		if page.Meta.NextCursor == nil {
			t.Fatalf("%s did not return a continuation cursor", path)
		}
		return *page.Meta.NextCursor
	}

	bidderFirst := v2ParticipantBidders + "?limit=2"
	bidderCursor := nextCursor(bidderFirst)
	winnerFirst := v2ParticipantWinners + "?limit=2"
	winnerCursor := nextCursor(winnerFirst)
	donorFirst := v2ParticipantDonors + "?limit=1"
	donorCursor := nextCursor(donorFirst)
	cstFirst := v2ParticipantCST + "?limit=1"
	cstCursor := nextCursor(cstFirst)
	rwalkFirst := v2ParticipantRWalk + "?limit=1"
	rwalkCursor := nextCursor(rwalkFirst)

	afterLast := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	cases := []v2GoldenCase{
		{name: "participants_bidders_first_page", target: bidderFirst, template: v2ParticipantBidders, pathParams: map[string]string{}},
		{name: "participants_bidders_next_page", target: bidderFirst + "&cursor=" + bidderCursor, template: v2ParticipantBidders, pathParams: map[string]string{}},
		{name: "participants_bidders_empty_page", target: bidderFirst + "&cursor=" + afterLast(`{"v":1,"k":"bidders","s":"1","a":25}`), template: v2ParticipantBidders, pathParams: map[string]string{}},
		{name: "participants_winners_first_page", target: winnerFirst, template: v2ParticipantWinners, pathParams: map[string]string{}},
		{name: "participants_winners_next_page", target: winnerFirst + "&cursor=" + winnerCursor, template: v2ParticipantWinners, pathParams: map[string]string{}},
		{name: "participants_winners_empty_page", target: winnerFirst + "&cursor=" + afterLast(`{"v":1,"k":"winners","s":"3","a":25}`), template: v2ParticipantWinners, pathParams: map[string]string{}},
		{name: "participants_donors_first_page", target: donorFirst, template: v2ParticipantDonors, pathParams: map[string]string{}},
		{name: "participants_donors_next_page", target: donorFirst + "&cursor=" + donorCursor, template: v2ParticipantDonors, pathParams: map[string]string{}},
		{name: "participants_donors_empty_page", target: donorFirst + "&cursor=" + afterLast(`{"v":1,"k":"donors","s":"200000000000000000","a":24}`), template: v2ParticipantDonors, pathParams: map[string]string{}},
		{name: "participants_cst_stakers_first_page", target: cstFirst, template: v2ParticipantCST, pathParams: map[string]string{}},
		{name: "participants_cst_stakers_next_page", target: cstFirst + "&cursor=" + cstCursor, template: v2ParticipantCST, pathParams: map[string]string{}},
		{name: "participants_cst_stakers_empty_page", target: cstFirst + "&cursor=" + afterLast(`{"v":1,"k":"cstStakers","s":"1000000000000000000","a":22}`), template: v2ParticipantCST, pathParams: map[string]string{}},
		{name: "participants_randomwalk_stakers_first_page", target: rwalkFirst, template: v2ParticipantRWalk, pathParams: map[string]string{}},
		{name: "participants_randomwalk_stakers_next_page", target: rwalkFirst + "&cursor=" + rwalkCursor, template: v2ParticipantRWalk, pathParams: map[string]string{}},
		{name: "participants_randomwalk_stakers_empty_page", target: rwalkFirst + "&cursor=" + afterLast(`{"v":1,"k":"randomWalkStakers","s":"0","a":23}`), template: v2ParticipantRWalk, pathParams: map[string]string{}},
		{name: "participants_dual_stakers_page", target: v2ParticipantBoth + "?limit=1", template: v2ParticipantBoth, pathParams: map[string]string{}},
		{name: "participants_dual_stakers_empty_page", target: v2ParticipantBoth + "?limit=1&cursor=" + afterLast(`{"v":1,"k":"dualStakers","s":"2","a":22}`), template: v2ParticipantBoth, pathParams: map[string]string{}},
		{name: "participants_error_cross_directory_cursor", target: winnerFirst + "&cursor=" + bidderCursor, template: v2ParticipantWinners, pathParams: map[string]string{}},
		{name: "participants_error_invalid_limit", target: v2ParticipantBidders + "?limit=201", template: v2ParticipantBidders, pathParams: map[string]string{}},
		{name: "participants_error_bind_limit", target: v2ParticipantWinners + "?limit=bad", template: v2ParticipantWinners, pathParams: map[string]string{}},
	}
	for name, path := range map[string]string{
		"bidders":            v2ParticipantBidders,
		"winners":            v2ParticipantWinners,
		"donors":             v2ParticipantDonors,
		"cst_stakers":        v2ParticipantCST,
		"randomwalk_stakers": v2ParticipantRWalk,
		"dual_stakers":       v2ParticipantBoth,
	} {
		cases = append(cases,
			v2GoldenCase{
				name:   "participants_" + name + "_error_malformed_cursor",
				target: path + "?cursor=bad", template: path, pathParams: map[string]string{},
			},
			v2GoldenCase{
				name:   "participants_" + name + "_error_internal",
				target: path, template: path, pathParams: map[string]string{}, ctx: cancelledCtx,
			},
		)
	}
	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2RoundBids(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	firstPath := "/api/v2/cosmicgame/rounds/0/bids?limit=2"
	firstResponse := h.get(t, firstPath)
	if firstResponse.Code != http.StatusOK {
		t.Fatalf("first page: status=%d body=%s", firstResponse.Code, firstResponse.Body.String())
	}
	var firstPage apiv2.RoundBidPage
	if err := json.Unmarshal(firstResponse.Body.Bytes(), &firstPage); err != nil {
		t.Fatalf("decoding first page: %v", err)
	}
	if firstPage.Meta.NextCursor == nil {
		t.Fatal("fixture first page did not return a continuation cursor")
	}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	cases := []v2GoldenCase{
		{
			name:       "list_first_page",
			target:     firstPath,
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "list_next_page",
			target:     "/api/v2/cosmicgame/rounds/0/bids?limit=2&cursor=" + *firstPage.Meta.NextCursor,
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "list_empty_page",
			target:     "/api/v2/cosmicgame/rounds/999/bids?limit=2",
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "999"},
		},
		{
			name:       "get_bid",
			target:     "/api/v2/cosmicgame/rounds/0/bids/3",
			template:   v2GetRoundBidPath,
			pathParams: map[string]string{"round": "0", "position": "3"},
		},
		{
			name:       "error_malformed_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/bids?cursor=not-a-cursor",
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "error_invalid_limit",
			target:     "/api/v2/cosmicgame/rounds/0/bids?limit=201",
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "error_bind_limit",
			target:     "/api/v2/cosmicgame/rounds/0/bids?limit=wat",
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "error_cross_round_cursor",
			target:     "/api/v2/cosmicgame/rounds/1/bids?cursor=" + *firstPage.Meta.NextCursor,
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "1"},
		},
		{
			name:       "error_bid_not_found",
			target:     "/api/v2/cosmicgame/rounds/0/bids/99",
			template:   v2GetRoundBidPath,
			pathParams: map[string]string{"round": "0", "position": "99"},
		},
		{
			name:       "error_internal",
			target:     "/api/v2/cosmicgame/rounds/0/bids?limit=2",
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
			ctx:        cancelledCtx,
		},
	}

	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2Rounds(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	firstPath := "/api/v2/cosmicgame/rounds?limit=2"
	firstResponse := h.get(t, firstPath)
	if firstResponse.Code != http.StatusOK {
		t.Fatalf("first page: status=%d body=%s", firstResponse.Code, firstResponse.Body.String())
	}
	var firstPage apiv2.RoundPage
	if err := json.Unmarshal(firstResponse.Body.Bytes(), &firstPage); err != nil {
		t.Fatalf("decoding first page: %v", err)
	}
	if firstPage.Meta.NextCursor == nil {
		t.Fatal("fixture first page did not return a continuation cursor")
	}

	afterRoundZero := base64.RawURLEncoding.EncodeToString([]byte(`{"v":1,"r":0,"e":5018}`))
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	cases := []v2GoldenCase{
		{
			name:       "rounds_list_first_page",
			target:     firstPath,
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_list_next_page",
			target:     "/api/v2/cosmicgame/rounds?limit=2&cursor=" + *firstPage.Meta.NextCursor,
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_list_empty_page",
			target:     "/api/v2/cosmicgame/rounds?limit=2&cursor=" + afterRoundZero,
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_get_0",
			target:     "/api/v2/cosmicgame/rounds/0",
			template:   v2GetRoundPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "rounds_get_2",
			target:     "/api/v2/cosmicgame/rounds/2",
			template:   v2GetRoundPath,
			pathParams: map[string]string{"round": "2"},
		},
		{
			name:       "rounds_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/rounds?cursor=not-a-cursor",
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_error_invalid_limit",
			target:     "/api/v2/cosmicgame/rounds?limit=201",
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_error_bind_limit",
			target:     "/api/v2/cosmicgame/rounds?limit=wat",
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_error_not_found",
			target:     "/api/v2/cosmicgame/rounds/999",
			template:   v2GetRoundPath,
			pathParams: map[string]string{"round": "999"},
		},
		{
			name:       "rounds_error_open_round",
			target:     "/api/v2/cosmicgame/rounds/3",
			template:   v2GetRoundPath,
			pathParams: map[string]string{"round": "3"},
		},
		{
			name:       "rounds_error_internal",
			target:     "/api/v2/cosmicgame/rounds?limit=2",
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
	}

	runV2GoldenCases(t, h, spec, cases)
}

func runV2GoldenCases(t *testing.T, h *harness, spec *openapi3.T, cases []v2GoldenCase) {
	t.Helper()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			request := request{path: tc.target, ctx: tc.ctx}
			first := h.do(t, request)
			second := h.do(t, request)

			firstBody := canonicalJSON(t, first.Body.Bytes())
			secondBody := canonicalJSON(t, second.Body.Bytes())
			if first.Code != second.Code || !reflect.DeepEqual(firstBody, secondBody) {
				t.Fatalf("nondeterministic v2 response: first=%d %#v second=%d %#v",
					first.Code, firstBody, second.Code, secondBody)
			}

			validateV2Response(t, spec, tc, first)
			compareV2Golden(t, tc.name, response{
				Status:      first.Code,
				ContentType: contentTypeOf(first),
				Body:        firstBody,
			})
		})
	}
}

func decodeV2JSON(t *testing.T, response *httptest.ResponseRecorder, target any) {
	t.Helper()
	if response.Code != http.StatusOK {
		t.Fatalf("response = %d %s", response.Code, response.Body.String())
	}
	if err := json.Unmarshal(response.Body.Bytes(), target); err != nil {
		t.Fatalf("decoding response: %v\n%s", err, response.Body.String())
	}
}

func validateV2Response(t *testing.T, spec *openapi3.T, tc v2GoldenCase, response *httptest.ResponseRecorder) {
	t.Helper()

	pathItem := spec.Paths.Value(tc.template)
	if pathItem == nil || pathItem.Get == nil {
		t.Fatalf("spec has no GET operation for %s", tc.template)
	}
	request := httptest.NewRequest(http.MethodGet, tc.target, nil)
	route := &routers.Route{
		Spec:      spec,
		Path:      tc.template,
		PathItem:  pathItem,
		Method:    http.MethodGet,
		Operation: pathItem.Get,
	}
	requestInput := &openapi3filter.RequestValidationInput{
		Request:    request,
		PathParams: tc.pathParams,
		Route:      route,
	}
	if response.Code < http.StatusBadRequest {
		if err := openapi3filter.ValidateRequest(context.Background(), requestInput); err != nil {
			t.Fatalf("%s request violates OpenAPI v2: %v", tc.name, err)
		}
	}
	responseInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestInput,
		Status:                 response.Code,
		Header:                 response.Header(),
	}
	responseInput.SetBodyBytes(response.Body.Bytes())
	if err := openapi3filter.ValidateResponse(context.Background(), responseInput); err != nil {
		t.Fatalf("%s response violates OpenAPI v2: %v\n%s", tc.name, err, response.Body.String())
	}
}
