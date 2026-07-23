//go:build integration

package apitest

import (
	"testing"
)

// TestErrorShapes pins the {"status":0,"error":...} envelope (including its
// HTTP status quirks) for representative invalid-parameter requests. These
// shapes are part of the frozen v1 contract: the frontend branches on them.
func TestErrorShapes(t *testing.T) {
	h := server(t)

	cases := []struct {
		name string
		path string
	}{
		// IsAddressValid writes the envelope with HTTP 200 (legacy quirk).
		{"bad_address_length", "/api/cosmicgame/user/info/zzz"},
		{"bad_address_hex", "/api/cosmicgame/user/info/0xzz00000000000000000000000000000000000021"},
		// Unknown addresses are valid but unregistered.
		{"unknown_address", "/api/cosmicgame/prizes/eth/raffle/by_user/0x9900000000000000000000000000000000000099"},
		// Numeric path params reject non-numeric input with HTTP 400.
		{"bad_offset", "/api/cosmicgame/rounds/list/abc/10"},
		{"bad_round", "/api/cosmicgame/statistics/claims/detail/xyz"},
		{"bad_order_by", "/api/randomwalk/current_offers/xyz"},
		{"bad_cosmic_market_order", "/api/cosmicgame/marketplace/current_offers/xyz"},
		{"bad_cosmic_market_offset", "/api/cosmicgame/marketplace/trading/sales/xyz/10"},
		{"bad_token_id", "/api/randomwalk/tokens/info/notanumber"},
		// Date params demand YYYYMMDD.
		{"bad_date_format", "/api/cosmicgame/ct/total_supply_history_by_date/2026-01-01/20260102"},
		{"reversed_date_range", "/api/cosmicgame/ct/total_supply_history_by_date/20260103/20260101"},
		// Missing rows behind valid params.
		{"missing_bid", "/api/cosmicgame/bid/info/999999"},
		{"missing_rw_token", "/api/randomwalk/tokens/info/9999"},
		{"missing_cst_token_metadata", "/api/cosmicgame/cst/metadata/9999"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := h.get(t, tc.path)
			compareGolden(t, "errors__"+tc.name, response{
				Status:      w.Code,
				ContentType: contentTypeOf(w),
				Body:        canonicalJSON(t, w.Body.Bytes()),
			})
		})
	}
}
