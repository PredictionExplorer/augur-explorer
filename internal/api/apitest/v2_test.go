//go:build integration

package apitest

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

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
	v2GlobalStatistics   = "/api/v2/cosmicgame/statistics"
	v2StatisticsCounters = "/api/v2/cosmicgame/statistics/counters"
	v2StatisticsROI      = "/api/v2/cosmicgame/statistics/leaderboard/roi"
	v2StatisticsClaims   = "/api/v2/cosmicgame/statistics/claims"
	v2RoundClaims        = "/api/v2/cosmicgame/rounds/{round}/claims"
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

	original := h.state.Snapshot()
	h.state.SetBidPrice("error", 0)
	defer h.state.SetBidPrice(original.BidPrice, original.BidPriceEth)
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
