//go:build integration

package apitest

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

// TestRandomWalkUserIdentifierValidation covers the RandomWalk user-scoped
// routes' identifier resolution: malformed ids, unindexed ids, hex-address
// resolution and the legacy 200-with-DBError shape for a wallet that exists
// in the address table but has no RandomWalk statistics.
func TestRandomWalkUserIdentifierValidation(t *testing.T) {
	h := server(t)

	badRequests := []struct {
		path    string
		wantErr string
	}{
		{"/api/randomwalk/user/info/notanumber", ""},
		{"/api/randomwalk/user/info/999999", "Address lookup on user_aid failed"},
		{"/api/randomwalk/tokens/by_user/notanumber", "Can't resolve user identifier"},
		{"/api/randomwalk/tokens/by_user/0x9999000000000000000000000000000000009999", "Cant find provided user"},
		{"/api/randomwalk/tokens/by_user/999999", "Address lookup on user_aid failed"},
		{"/api/randomwalk/trading/by_user/notanumber/0/10", ""},
		{"/api/randomwalk/statistics/floor_price/notanumber/0/0", ""},
	}
	for _, tc := range badRequests {
		t.Run(tc.path, func(t *testing.T) {
			w := h.get(t, tc.path)
			if w.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400\n%s", w.Code, w.Body.String())
			}
			status, errStr := decodeEnvelope(t, w.Body.Bytes())
			if status != 0 || errStr == "" {
				t.Fatalf("envelope = (%d, %q)", status, errStr)
			}
			if tc.wantErr != "" && !strings.Contains(errStr, tc.wantErr) {
				t.Fatalf("error = %q, want %q", errStr, tc.wantErr)
			}
		})
	}

	t.Run("wallet without randomwalk stats keeps the legacy DBError shape", func(t *testing.T) {
		// Emma (aid 25) exists in the address table (CosmicGame donor) but
		// has no RandomWalk statistics rows.
		w := h.get(t, "/api/randomwalk/user/info/25")
		if w.Code != http.StatusOK {
			t.Fatalf("status = %d, want the legacy 200 shape\n%s", w.Code, w.Body.String())
		}
		var resp struct {
			Status  int    `json:"status"`
			DBError string `json:"DBError"`
			UserAid int64  `json:"UserAid"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if resp.Status != 1 || resp.UserAid != 25 || resp.DBError == "" {
			t.Fatalf("response = %+v, want status 1 with the legacy DBError text", resp)
		}
	})

	t.Run("hex address resolves to the indexed user", func(t *testing.T) {
		w := h.get(t, "/api/randomwalk/tokens/by_user/"+addrCarol)
		if w.Code != http.StatusOK {
			t.Fatalf("status = %d\n%s", w.Code, w.Body.String())
		}
		var resp struct {
			Status  int   `json:"status"`
			UserAid int64 `json:"UserAid"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if resp.Status != 1 || resp.UserAid != 23 {
			t.Fatalf("response = %+v, want carol resolved to aid 23", resp)
		}
	})
}

// TestRandomWalkFloorPriceDefaults pins the zero-parameter defaults of the
// floor-price time series: zero timestamps select the deployment epoch and
// "now", and the legacy INT_MAX interval sentinel selects one day.
func TestRandomWalkFloorPriceDefaults(t *testing.T) {
	h := server(t)

	w := h.get(t, "/api/randomwalk/statistics/floor_price/0/0/2147483647")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d\n%s", w.Code, w.Body.String())
	}
	var resp struct {
		Status   int   `json:"status"`
		InitTs   int64 `json:"InitTs"`
		FinTs    int64 `json:"FinTs"`
		Interval int64 `json:"Interval"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if resp.Status != 1 || resp.InitTs != 1636676049 || resp.Interval != 86400 {
		t.Fatalf("defaults = %+v", resp)
	}
	if resp.FinTs == 0 {
		t.Fatal("zero fin_ts must default to the current time")
	}
}

// TestRandomWalkBeautyPairVoterFilter covers the beauty-pair selection's
// voter-scoped re-roll: a voter with recorded votes re-rolls to an unvoted
// pair, unknown/malformed voters fall back to the unfiltered pick, and
// skip_pair_filter bypasses the filter entirely.
func TestRandomWalkBeautyPairVoterFilter(t *testing.T) {
	h := server(t)

	paths := []string{
		"/api/randomwalk/ranking/beauty-pair-ids?skip_pair_filter=1",
		"/api/randomwalk/ranking/beauty-pair-ids?voter=" + addrCarol,
		"/api/randomwalk/ranking/beauty-pair-ids?voter=0x9999000000000000000000000000000000009999",
		"/api/randomwalk/ranking/beauty-pair-ids?voter=not-an-address",
		"/api/randomwalk/ranking/beauty-pair-ids",
	}
	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			w := h.get(t, path)
			if w.Code != http.StatusOK {
				t.Fatalf("status = %d\n%s", w.Code, w.Body.String())
			}
			var resp struct {
				TokenIDs      []int64 `json:"token_ids"`
				PairExhausted *bool   `json:"pair_exhausted"`
			}
			if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
				t.Fatalf("decode: %v", err)
			}
			if len(resp.TokenIDs) != 2 || resp.TokenIDs[0] == resp.TokenIDs[1] {
				t.Fatalf("token_ids = %v, want two distinct ids", resp.TokenIDs)
			}
			if resp.PairExhausted == nil {
				t.Fatal("pair_exhausted missing from response")
			}
		})
	}
}
