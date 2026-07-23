package smoketest

import (
	"slices"
	"strings"
	"testing"
)

func TestBuildEndpointsCompletenessAndStableOrder(t *testing.T) {
	t.Parallel()
	// #nosec G101 -- deterministic URL placeholders, not credentials.
	params := Params{
		UserAddress:        "0xUser",
		CSTStakerAddress:   "0xStaker",
		RoundNumber:        "77",
		BidEventLogID:      "901",
		BidRound:           "76",
		BidPosition:        "12",
		TokenID:            "333",
		TokenName:          "Alpha/Beta ?",
		ETHDonationID:      "41",
		NFTDonationID:      "42",
		ERC20DonationID:    "43",
		CSTActionID:        "51",
		RandomWalkActionID: "52",
		DepositID:          "61",
		NFTTokenAddress:    "0xNFT",
		TimestampMin:       "100",
		TimestampMax:       "200",
		FromDate:           "20260101",
		ToDate:             "20260131",
	}
	first := BuildEndpoints(params)
	second := BuildEndpoints(params)
	if len(first) != 145 {
		t.Fatalf("endpoint count = %d, want 145", len(first))
	}
	if !slices.Equal(first, second) {
		t.Fatal("endpoint order changed between calls")
	}
	if first[0] != "/api/cosmicgame/statistics/dashboard" {
		t.Fatalf("first endpoint = %q", first[0])
	}
	if first[len(first)-1] != "/api/cosmicgame/system/admin_events/0/9223372036854775807" {
		t.Fatalf("last endpoint = %q", first[len(first)-1])
	}

	want := []string{
		"/api/cosmicgame/statistics/bidding/activity/100/200/3600",
		"/api/cosmicgame/rounds/info/77",
		"/api/cosmicgame/bid/info/901",
		"/api/cosmicgame/bid/info_by_pos/76/12",
		"/api/cosmicgame/marketplace/current_offers/0",
		"/api/cosmicgame/marketplace/floor_price",
		"/api/cosmicgame/marketplace/trading/sales/0/1000000",
		"/api/cosmicgame/cst/names/search/Alpha%2FBeta%20%3F",
		"/api/cosmicgame/ct/total_supply_history_by_date/20260101/20260131",
		"/api/cosmicgame/donations/eth/with_info/info/41",
		"/api/cosmicgame/donations/nft/info/42",
		"/api/cosmicgame/donations/nft/by_token/0xNFT",
		"/api/cosmicgame/donations/erc20/info/43",
		"/api/cosmicgame/staking/cst/actions/info/51",
		"/api/cosmicgame/staking/cst/rewards/action_ids_by_deposit/0xStaker/61",
		"/api/cosmicgame/staking/randomwalk/actions/info/52",
		"/api/cosmicgame/staking/rwalk/actions/info/52",
	}
	for _, endpoint := range want {
		if !slices.Contains(first, endpoint) {
			t.Errorf("missing endpoint %q", endpoint)
		}
	}

	seen := make(map[string]struct{}, len(first))
	for _, endpoint := range first {
		if !strings.HasPrefix(endpoint, "/api/cosmicgame/") {
			t.Errorf("endpoint has wrong prefix: %q", endpoint)
		}
		if _, duplicate := seen[endpoint]; duplicate {
			t.Errorf("duplicate endpoint: %q", endpoint)
		}
		seen[endpoint] = struct{}{}
	}
}

func TestBuildEndpointsFillsPartialParams(t *testing.T) {
	t.Parallel()
	endpoints := BuildEndpoints(Params{UserAddress: "0xCustom", RoundNumber: "9"})
	for _, want := range []string{
		"/api/cosmicgame/bid/info/1",
		"/api/cosmicgame/bid/info_by_pos/9/1",
		"/api/cosmicgame/cst/info/0",
		"/api/cosmicgame/staking/cst/rewards/action_ids_by_deposit/0xCustom/0",
	} {
		if !slices.Contains(endpoints, want) {
			t.Errorf("missing default-backed endpoint %q", want)
		}
	}
}
