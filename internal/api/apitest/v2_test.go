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
