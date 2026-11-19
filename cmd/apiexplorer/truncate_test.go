package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCapLists(t *testing.T) {
	long := make([]any, 250)
	nested := make([]any, 300)
	data := map[string]any{
		"Rows":    long,
		"Short":   []any{1, 2, 3},
		"Scalar":  "x",
		"Wrapped": map[string]any{"Inner": nested},
	}
	capped := capLists(data, 200)
	if len(capped) != 2 {
		t.Fatalf("capped = %+v", capped)
	}
	// Deterministic order (sorted by name) regardless of map iteration.
	if capped[0].Name != "Rows" || capped[0].Shown != 200 || capped[0].Total != 250 {
		t.Errorf("capped[0] = %+v", capped[0])
	}
	if capped[1].Name != "Wrapped.Inner" || capped[1].Total != 300 {
		t.Errorf("capped[1] = %+v", capped[1])
	}
	if len(data["Rows"].([]any)) != 200 ||
		len(data["Wrapped"].(map[string]any)["Inner"].([]any)) != 200 {
		t.Error("lists not truncated in place")
	}
	if len(data["Short"].([]any)) != 3 {
		t.Error("short list must stay untouched")
	}
	if again := capLists(data, 200); len(again) != 0 {
		t.Errorf("second pass capped %+v", again)
	}
}

func TestFullTableURL(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/cosmicsignature/bids?round=3", nil)
	if got := fullTableURL(r); got != "/cosmicsignature/bids?full=1&round=3" {
		t.Errorf("fullTableURL = %q", got)
	}
	if wantFull(r) {
		t.Error("wantFull without full=1")
	}
	if !wantFull(httptest.NewRequest(http.MethodGet, "/x?full=1", nil)) {
		t.Error("wantFull with full=1")
	}
}

// TestPageHandlerCapsLongLists renders a page whose API response has more
// rows than maxTableRows: the first fetch shows maxTableRows rows plus the
// banner; ?full=1 shows everything and no banner.
func TestPageHandlerCapsLongLists(t *testing.T) {
	const total = maxTableRows + 101
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rows := make([]string, total)
		for i := range rows {
			rows[i] = fmt.Sprintf(`{"BucketTs":%d,"EthBids":1,"RwalkBids":0,"CstBids":0,"TotalBids":1,"EthPct":100,"RwalkPct":0,"CstPct":0}`, i)
		}
		_, _ = io.WriteString(w, `{"error":"","status":1,"FromTs":0,"ToTs":1,"Interval":86400,"RatioHistory":[`+
			strings.Join(rows, ",")+`]}`)
	}))
	defer backend.Close()
	s := newTestServer(t, backend.URL)
	mux := http.NewServeMux()
	p := page{"GET /cosmicsignature/bid/bid_type_ratio", "cosmicsignature/cg_bid_type_ratio.html", "/api/cosmicgame/bid/bid_type_ratio"}
	mux.HandleFunc(p.Pattern, s.pageHandler(p))

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/cosmicsignature/bid/bid_type_ratio", nil))
	body := rec.Body.String()
	if got := strings.Count(body, `<td class="monospace">`); got != maxTableRows {
		t.Errorf("first fetch renders %d rows, want %d", got, maxTableRows)
	}
	wantBanner := fmt.Sprintf("Showing only the first %d of %d RatioHistory records", maxTableRows, total)
	if !strings.Contains(body, wantBanner) {
		t.Errorf("banner %q missing", wantBanner)
	}
	if !strings.Contains(body, `href="/cosmicsignature/bid/bid_type_ratio?full=1"`) {
		t.Error("show-all link missing")
	}

	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/cosmicsignature/bid/bid_type_ratio?full=1", nil))
	body = rec.Body.String()
	if got := strings.Count(body, `<td class="monospace">`); got != total {
		t.Errorf("full fetch renders %d rows, want %d", got, total)
	}
	if strings.Contains(body, "Showing only the first") {
		t.Error("full fetch must not show the cap banner")
	}

	// The generic fields view is capped too.
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/cosmicsignature/bid/bid_type_ratio?view=fields", nil))
	body = rec.Body.String()
	if !strings.Contains(body, wantBanner) {
		t.Error("fields view must show the cap banner")
	}
}
