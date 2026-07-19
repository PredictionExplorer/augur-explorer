//go:build integration

package apitest

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

var pathParameterPattern = regexp.MustCompile(`\{([^}]+)\}`)

func TestV1NumericPathParametersRejectMalformedValues(t *testing.T) {
	h := server(t)
	numeric := map[string]bool{
		"action_id":     true,
		"bid_position":  true,
		"deposit_id":    true,
		"evtlog_end":    true,
		"evtlog_id":     true,
		"evtlog_start":  true,
		"fin_ts":        true,
		"init_ts":       true,
		"interval_secs": true,
		"limit":         true,
		"offset":        true,
		"order_by":      true,
		"prize_num":     true,
		"record_id":     true,
		"round":         true,
		"round_num":     true,
		"sort":          true,
		"token_id":      true,
		"user_aid":      true,
	}
	legacyIgnored := map[string]map[string]bool{
		"/api/randomwalk/trading/by_user/{user_aid}/{offset}/{limit}": {
			"offset": true,
			"limit":  true,
		},
	}

	tested := 0
	for _, route := range h.router.Routes() {
		if route.Method != http.MethodGet ||
			strings.HasPrefix(route.Pattern, "/api/v2/") ||
			strings.HasPrefix(route.Pattern, "/metadata/") ||
			strings.HasPrefix(route.Pattern, "/cg/metadata/") {
			continue
		}
		matches := pathParameterPattern.FindAllStringSubmatch(route.Pattern, -1)
		for _, match := range matches {
			target := match[1]
			if !numeric[target] {
				continue
			}
			if legacyIgnored[route.Pattern][target] {
				continue
			}
			tested++
			name := route.Pattern + " invalid " + target
			t.Run(name, func(t *testing.T) {
				path := concreteV1Path(route.Pattern, target, "not-an-integer")
				response := h.get(t, path)
				if response.Code != http.StatusBadRequest {
					t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
				}
				var envelope struct {
					Status int    `json:"status"`
					Error  string `json:"error"`
				}
				if err := json.Unmarshal(response.Body.Bytes(), &envelope); err != nil {
					t.Fatalf("decode response: %v\n%s", err, response.Body.String())
				}
				if envelope.Status != 0 || strings.TrimSpace(envelope.Error) == "" ||
					strings.Contains(strings.ToLower(envelope.Error), "internal") {
					t.Fatalf("unexpected malformed-parameter envelope: %+v", envelope)
				}
			})
		}
	}
	if tested < 100 {
		t.Fatalf("tested only %d malformed numeric parameter positions", tested)
	}
}

func TestCosmicGameUserAddressBoundaries(t *testing.T) {
	h := server(t)
	const unknownAddress = "0x9900000000000000000000000000000000000099"

	tested := 0
	for _, route := range h.router.Routes() {
		if route.Method != http.MethodGet ||
			!strings.HasPrefix(route.Pattern, "/api/cosmicgame/") ||
			!strings.Contains(route.Pattern, "{user_addr}") {
			continue
		}
		tested++
		t.Run(route.Pattern+" malformed", func(t *testing.T) {
			response := h.get(t, concreteV1Path(route.Pattern, "user_addr", "not-an-address"))
			if response.Code >= http.StatusInternalServerError {
				t.Fatalf("malformed wallet caused server failure: status=%d body=%s",
					response.Code, response.Body.String())
			}
			var body any
			if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
				t.Fatalf("decode response: %v\n%s", err, response.Body.String())
			}
			assertNoInternalDetail(t, response.Body.String())
		})

		t.Run(route.Pattern+" unindexed", func(t *testing.T) {
			response := h.get(t, concreteV1Path(route.Pattern, "user_addr", unknownAddress))
			if response.Code >= http.StatusInternalServerError {
				t.Fatalf("unindexed wallet caused server failure: status=%d body=%s",
					response.Code, response.Body.String())
			}
			var body any
			if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
				t.Fatalf("decode response: %v\n%s", err, response.Body.String())
			}
			assertNoInternalDetail(t, response.Body.String())
		})
	}
	if tested < 25 {
		t.Fatalf("tested only %d user-address route templates", tested)
	}
}

func TestV1OptionalQueryParametersUseDocumentedBounds(t *testing.T) {
	h := server(t)

	t.Run("ROI clamps negative and excessive values", func(t *testing.T) {
		response := h.get(t, "/api/cosmicgame/statistics/leaderboard/roi?min_bids=-5&offset=-9&limit=1001&sort_by=not-a-sort")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var body struct {
			Status  int `json:"status"`
			MinBids int `json:"MinBids"`
			Offset  int `json:"Offset"`
			Limit   int `json:"Limit"`
		}
		if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
			t.Fatal(err)
		}
		if body.Status != 1 || body.MinBids != 0 || body.Offset != 0 || body.Limit != 100 {
			t.Fatalf("clamped ROI parameters = %+v", body)
		}
	})

	t.Run("bid ratio rejects non-positive interval by using daily buckets", func(t *testing.T) {
		response := h.get(t, "/api/cosmicgame/bid/bid_type_ratio?from_ts="+cgFrom+"&to_ts="+cgTo+"&interval_secs=0")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var body struct {
			Status   int `json:"status"`
			Interval int `json:"Interval"`
		}
		if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
			t.Fatal(err)
		}
		if body.Status != 1 || body.Interval != 86400 {
			t.Fatalf("bid-ratio defaults = %+v", body)
		}
	})

	t.Run("top periods use defaults for malformed optional filters", func(t *testing.T) {
		response := h.get(t, "/api/cosmicgame/statistics/bidding/top_active_periods/invalid/"+cgFrom+"/"+cgTo+
			"?gap_hours=invalid&min_bids=invalid")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var body struct {
			Status   int `json:"status"`
			TopN     int `json:"TopN"`
			GapHours int `json:"GapHours"`
			MinBids  int `json:"MinBids"`
		}
		if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
			t.Fatal(err)
		}
		if body.Status != 1 || body.TopN != 20 || body.GapHours != 6 || body.MinBids != 2 {
			t.Fatalf("top-period defaults = %+v", body)
		}
	})
}

func TestV1MissingRecordsUsePublicNotFoundShapes(t *testing.T) {
	h := server(t)
	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantError  string
	}{
		{name: "round", path: "/api/cosmicgame/rounds/info/999999", wantStatus: http.StatusBadRequest, wantError: "record not found"},
		{name: "bid position", path: "/api/cosmicgame/bid/info_by_pos/999999/1", wantStatus: http.StatusBadRequest, wantError: "record not found"},
		{name: "NFT donation", path: "/api/cosmicgame/donations/nft/info/999999", wantStatus: http.StatusBadRequest, wantError: "Record not found"},
		{name: "ERC20 donation", path: "/api/cosmicgame/donations/erc20/info/999999", wantStatus: http.StatusBadRequest, wantError: "Record not found"},
		{name: "CST stake action", path: "/api/cosmicgame/staking/cst/actions/info/999999", wantStatus: http.StatusBadRequest, wantError: "record not found"},
		{name: "RandomWalk stake action", path: "/api/cosmicgame/staking/randomwalk/actions/info/999999", wantStatus: http.StatusBadRequest, wantError: "record not found"},
		{name: "CosmicSignature token", path: "/api/cosmicgame/cst/info/999999", wantStatus: http.StatusBadRequest, wantError: "record not found"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := h.get(t, test.path)
			if response.Code != test.wantStatus {
				t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
			}
			var envelope struct {
				Status int    `json:"status"`
				Error  string `json:"error"`
			}
			if err := json.Unmarshal(response.Body.Bytes(), &envelope); err != nil {
				t.Fatalf("decode response: %v\n%s", err, response.Body.String())
			}
			if envelope.Status != 0 || envelope.Error != test.wantError {
				t.Fatalf("envelope = %+v", envelope)
			}
		})
	}
}

func assertNoInternalDetail(t *testing.T, body string) {
	t.Helper()
	lowerBody := strings.ToLower(body)
	for _, internal := range []string{"context canceled", "pgx", "sqlstate", "password"} {
		if strings.Contains(lowerBody, internal) {
			t.Fatalf("response leaked internal detail %q: %s", internal, body)
		}
	}
}

func concreteV1Path(template, invalidName, invalidValue string) string {
	return pathParameterPattern.ReplaceAllStringFunc(template, func(placeholder string) string {
		name := strings.TrimSuffix(strings.TrimPrefix(placeholder, "{"), "}")
		if name == invalidName {
			return invalidValue
		}
		switch name {
		case "user_addr":
			return addrAlice
		case "token_addr":
			return nftDonationContract
		case "name":
			return "Genesis"
		case "from_date":
			return "20260101"
		case "to_date":
			return "20260102"
		case "init_ts":
			return cgFrom
		case "fin_ts":
			return cgTo
		case "interval_secs":
			return cgStep
		case "limit":
			return "10"
		case "sort", "offset":
			return "0"
		case "user_aid":
			return aidCarol
		default:
			return "1"
		}
	})
}
