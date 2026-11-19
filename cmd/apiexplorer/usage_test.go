package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestParamDocsCoverRoutes fails when any route table uses a placeholder
// that paramDocs does not explain, so the usage footer can never show an
// undocumented parameter.
func TestParamDocsCoverRoutes(t *testing.T) {
	tables := map[string][]page{
		"pages":          pages,
		"generatedPages": generatedPages,
		"czPages":        czPages,
	}
	for table, list := range tables {
		for _, p := range list {
			for _, src := range []string{patternPath(p.Pattern), p.API} {
				for _, name := range paramNames(src) {
					if strings.TrimSpace(paramDocs[name]) == "" {
						t.Errorf("%s: %s: placeholder {%s} has no paramDocs entry",
							table, p.Pattern, name)
					}
				}
			}
		}
	}
}

func TestUsageFor(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet,
		"http://localhost:7272/cosmicsignature/bid/info/905", nil)
	p := page{
		"GET /cosmicsignature/bid/info/{evtlog_id}",
		"cosmicsignature/cg_bid_info.html",
		"/api/cosmicgame/bid/info/{evtlog_id}",
	}
	u := usageFor(p, req)
	if u == nil {
		t.Fatal("usageFor returned nil for an API-backed page")
	}
	if u.Origin != "http://localhost:7272" {
		t.Errorf("Origin = %q", u.Origin)
	}
	if u.Pattern != "/cosmicsignature/bid/info/{evtlog_id}" {
		t.Errorf("Pattern = %q", u.Pattern)
	}
	if u.Example != "/cosmicsignature/bid/info/905" {
		t.Errorf("Example = %q", u.Example)
	}
	if u.APIPattern != "/api/cosmicgame/bid/info/{evtlog_id}" {
		t.Errorf("APIPattern = %q", u.APIPattern)
	}
	if len(u.Params) != 1 || u.Params[0].Name != "evtlog_id" || u.Params[0].Desc == "" {
		t.Errorf("Params = %+v", u.Params)
	}

	// Static pages (no API behind them) get no footer.
	if got := usageFor(page{"GET /cosmicsignature/api_docs", "x.html", ""}, req); got != nil {
		t.Errorf("static page usage = %+v, want nil", got)
	}

	// The {$} exact-match marker is not part of the shown pattern, and
	// parameters keep their path order.
	u = usageFor(page{
		"GET /chaoszero/trades/by_round/{round_num}/{offset}/{limit}",
		"chaoszero/cz_trades.html",
		"/api/chaoszero/trades/by_round/{round_num}/{offset}/{limit}",
	}, req)
	if len(u.Params) != 3 || u.Params[0].Name != "round_num" ||
		u.Params[1].Name != "offset" || u.Params[2].Name != "limit" {
		t.Errorf("Params = %+v", u.Params)
	}
	u = usageFor(page{"GET /chaoszero/{$}", "chaoszero/home.html", "/api/chaoszero/statistics"}, req)
	if u.Pattern != "/chaoszero/" || len(u.Params) != 0 {
		t.Errorf("home usage = %+v", u)
	}
}

// TestUsageFooterRendered checks the layout renders the footer: pattern,
// example link, backing API endpoint and the parameter explanations.
func TestUsageFooterRendered(t *testing.T) {
	s := newTestServer(t, "http://api.example")
	req := httptest.NewRequest(http.MethodGet,
		"http://localhost:7272/cosmicsignature/bid/list/by_round/3/1/0/25", nil)
	p := page{
		"GET /cosmicsignature/bid/list/by_round/{round_num}/{sort}/{offset}/{limit}",
		"cosmicsignature/cg_bids_by_round_num.html",
		"/api/cosmicgame/bid/list/by_round/{round_num}/{sort}/{offset}/{limit}",
	}
	meta := pageMeta{
		Project: "cosmicsignature",
		Path:    req.URL.Path,
		API:     "/api/cosmicgame/bid/list/by_round/3/1/0/25",
		Usage:   usageFor(p, req),
	}
	rec := httptest.NewRecorder()
	s.renderPage(rec, http.StatusOK, "cosmicsignature/cg_bids.html",
		map[string]any{"Bids": []any{}, "AllRounds": true, "SelectedRound": 3,
			"CurRoundNum": 3, "RoundChoices": []int{3, 2, 1, 0}}, meta)
	body := rec.Body.String()
	for _, want := range []string{
		"How to use this request",
		"GET http://localhost:7272/cosmicsignature/bid/list/by_round/{round_num}/{sort}/{offset}/{limit}",
		`<a href="/cosmicsignature/bid/list/by_round/3/1/0/25">`,
		"GET http://api.example/api/cosmicgame/bid/list/by_round/{round_num}/{sort}/{offset}/{limit}",
		`<a href="/api/cosmicgame/bid/list/by_round/3/1/0/25"`,
		"round_num", paramDocs["round_num"],
		"sort", paramDocs["sort"],
		"offset", paramDocs["offset"],
		"limit", paramDocs["limit"],
	} {
		if !strings.Contains(body, want) {
			t.Errorf("usage footer missing %q", want)
		}
	}
	// No footer without usage metadata (error pages, index).
	rec = httptest.NewRecorder()
	s.render(rec, http.StatusOK, "index.html", nil)
	if strings.Contains(rec.Body.String(), "How to use this request") {
		t.Error("index page must not render a usage footer")
	}
}
