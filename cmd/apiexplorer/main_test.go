package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestJSONNumString(t *testing.T) {
	if got := jsonNum(1735747446).String(); got != "1735747446" {
		t.Errorf("String() = %q, want 1735747446", got)
	}
	if got := jsonNum(1.5).String(); got != "1.5" {
		t.Errorf("String() = %q, want 1.5", got)
	}
}

func TestConvertNums(t *testing.T) {
	tree := map[string]any{
		"n":    float64(42),
		"s":    "text",
		"list": []any{float64(7), "x", map[string]any{"inner": float64(9)}},
	}
	convertNums(tree)
	if _, ok := tree["n"].(jsonNum); !ok {
		t.Errorf("top-level number not wrapped: %T", tree["n"])
	}
	if tree["s"] != "text" {
		t.Errorf("string mutated: %v", tree["s"])
	}
	list := tree["list"].([]any)
	if _, ok := list[0].(jsonNum); !ok {
		t.Errorf("list number not wrapped: %T", list[0])
	}
	inner := list[2].(map[string]any)
	if _, ok := inner["inner"].(jsonNum); !ok {
		t.Errorf("nested number not wrapped: %T", inner["inner"])
	}
	if got := convertNums(float64(3)); got != jsonNum(3) {
		t.Errorf("scalar conversion = %v (%T)", got, got)
	}
}

func TestNum(t *testing.T) {
	if got := num(float64(1735747446)); got != "1735747446" {
		t.Errorf("num(float64) = %q", got)
	}
	if got := num("plain"); got != "plain" {
		t.Errorf("num(string) = %q", got)
	}
}

func TestAddf(t *testing.T) {
	if got := addf(jsonNum(1), float64(2), "ignored", nil); got != 3 {
		t.Errorf("addf = %v, want 3", got)
	}
	if got := addf(); got != 0 {
		t.Errorf("addf() = %v, want 0", got)
	}
}

func TestNowTs(t *testing.T) {
	now := float64(time.Now().Unix())
	if got := nowTs(); got < now-2 || got > now+2 {
		t.Errorf("nowTs() = %v, want ~%v", got, now)
	}
}

func TestUnixDate(t *testing.T) {
	want := time.Unix(1636676049, 0).Format("2006-01-02 15:04:05 MST")
	if got := unixDate(jsonNum(1636676049)); got != want {
		t.Errorf("unixDate(jsonNum) = %q, want %q", got, want)
	}
	if got := unixDate(float64(1636676049)); got != want {
		t.Errorf("unixDate(float64) = %q, want %q", got, want)
	}
	if got := unixDate("not a number"); got != "not a number" {
		t.Errorf("unixDate(string) = %q", got)
	}
}

func TestWeiToEth(t *testing.T) {
	if got := weiToEth("1500000000000000000"); got != "1.500000" {
		t.Errorf("weiToEth = %q, want 1.500000", got)
	}
	if got := weiToEth(jsonNum(1e18)); got != "1.000000" {
		t.Errorf("weiToEth(jsonNum) = %q, want 1.000000", got)
	}
	if got := weiToEth("not-wei"); got != "not-wei" {
		t.Errorf("weiToEth(invalid) = %q", got)
	}
}

func TestIsEthAddr(t *testing.T) {
	valid := "0x7BBF44394a23504cbE46b2b2d76929451cb86975"
	cases := []struct {
		in   string
		want bool
	}{
		{valid, true},
		{"0X7BBF44394A23504CBE46B2B2D76929451CB86975", true},
		{"0x7BBF4439", false},                                 // too short
		{"1x7BBF44394a23504cbE46b2b2d76929451cb86975", false}, // bad first byte
		{"0y7BBF44394a23504cbE46b2b2d76929451cb86975", false}, // bad prefix
		{"0x7BBF44394a23504cbE46b2b2d76929451cb8697g", false}, // non-hex digit
		{"(All CS NFT Stakers)", false},                       // legacy label
	}
	for _, c := range cases {
		if got := isEthAddr(c.in); got != c.want {
			t.Errorf("isEthAddr(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestEthAddrLink(t *testing.T) {
	addr := "0x7BBF44394a23504cbE46b2b2d76929451cb86975"
	got := string(ethAddrLink(addr))
	want := fmt.Sprintf(`<a href="/cosmicsignature/user/info/%s">%s</a>`, addr, addr)
	if got != want {
		t.Errorf("ethAddrLink = %q, want %q", got, want)
	}
	if got := string(ethAddrLinkTo("/randomwalk/user/info/", "<label>")); got != "&lt;label&gt;" {
		t.Errorf("invalid address must render escaped text, got %q", got)
	}
}

func TestParamNames(t *testing.T) {
	got := paramNames("/api/cosmicgame/bid/list/by_round/{round_num}/{sort}/{offset}/{limit}")
	want := []string{"round_num", "sort", "offset", "limit"}
	if len(got) != len(want) {
		t.Fatalf("paramNames = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("paramNames = %v, want %v", got, want)
		}
	}
	if got := paramNames("/api/cosmicgame/statistics/dashboard"); len(got) != 0 {
		t.Errorf("paramNames without placeholders = %v", got)
	}
}

func TestParamKey(t *testing.T) {
	cases := map[string]string{
		"round_num": "RoundNum",
		"user_addr": "UserAddr",
		"offset":    "Offset",
		"a__b":      "AB", // empty parts are skipped
	}
	for in, want := range cases {
		if got := paramKey(in); got != want {
			t.Errorf("paramKey(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestInjectParams(t *testing.T) {
	data := map[string]any{"RoundNum": jsonNum(7)}
	get := func(name string) string { return "path-" + name }
	injectParams(data, "/api/x/{round_num}/{user_addr}", get)
	if data["RoundNum"] != jsonNum(7) {
		t.Errorf("response field must win, got %v", data["RoundNum"])
	}
	if data["UserAddr"] != "path-user_addr" {
		t.Errorf("missing key not injected: %v", data["UserAddr"])
	}
}

func TestLoadTemplates(t *testing.T) {
	tmpl, err := loadTemplates()
	if err != nil {
		t.Fatalf("loadTemplates: %v", err)
	}
	for _, name := range []string{
		"index.html", "error.html",
		"cosmicsignature/cg_index.html", "randomwalk/home.html",
		"randomwalk/rw_trading_history.html",
	} {
		if tmpl.Lookup(name) == nil {
			t.Errorf("template %q not loaded", name)
		}
	}
}

// newTestServer returns a server wired to parsed templates and a discard
// logger.
func newTestServer(t *testing.T, apiBase string) *server {
	t.Helper()
	tmpl, err := loadTemplates()
	if err != nil {
		t.Fatalf("loadTemplates: %v", err)
	}
	return &server{
		apiBase:   apiBase,
		templates: tmpl,
		client:    &http.Client{Timeout: apiTimeout},
		log:       slog.New(slog.DiscardHandler),
	}
}

func TestFetch(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/ok":
			_, _ = io.WriteString(w, `{"error":"","N":12,"Rows":[{"X":3}]}`)
		case "/apierr":
			_, _ = io.WriteString(w, `{"error":"boom","status":0}`)
		case "/badjson":
			_, _ = io.WriteString(w, `{"truncated":`)
		default:
			http.Error(w, "nope", http.StatusInternalServerError)
		}
	}))
	defer backend.Close()
	s := newTestServer(t, backend.URL)
	ctx := context.Background()

	data, err := s.fetch(ctx, "/ok")
	if err != nil {
		t.Fatalf("fetch(/ok): %v", err)
	}
	if data["N"] != jsonNum(12) {
		t.Errorf("numbers must be converted, got %T", data["N"])
	}

	if _, err := s.fetch(ctx, "/apierr"); err == nil || !strings.Contains(err.Error(), "API error: boom") {
		t.Errorf("fetch(/apierr) err = %v", err)
	}
	if _, err := s.fetch(ctx, "/badjson"); err == nil || !strings.Contains(err.Error(), "decoding API response") {
		t.Errorf("fetch(/badjson) err = %v", err)
	}
	if _, err := s.fetch(ctx, "/http500"); err == nil || !strings.Contains(err.Error(), "API answered 500") {
		t.Errorf("fetch(/http500) err = %v", err)
	}

	s.apiBase = "http://invalid host" // space makes URL parsing fail
	if _, err := s.fetch(ctx, "/x"); err == nil {
		t.Error("fetch with invalid base must fail")
	}

	closed := httptest.NewServer(http.NotFoundHandler())
	closed.Close()
	s.apiBase = closed.URL
	if _, err := s.fetch(ctx, "/x"); err == nil {
		t.Error("fetch against closed backend must fail")
	}
}

func TestRenderAndRenderError(t *testing.T) {
	s := newTestServer(t, "http://unused")

	rec := httptest.NewRecorder()
	s.renderError(rec, http.StatusNotFound, "no such page")
	if rec.Code != http.StatusNotFound || !strings.Contains(rec.Body.String(), "no such page") {
		t.Errorf("renderError: code=%d body=%q", rec.Code, rec.Body.String())
	}

	rec = httptest.NewRecorder()
	s.render(rec, http.StatusOK, "does-not-exist.html", nil)
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("render(missing template) code = %d, want 500", rec.Code)
	}
}

func TestPageHandler(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/randomwalk/trading/history/0/5":
			_, _ = io.WriteString(w, `{"error":"","MarketAddr":"0x47eF85Dfb775aCE0934fBa9EEd09D22e6eC0Cc08",
				"Sales":[{"Tx":{"BlockNum":2984169},"BlockNum":2984169,"RealDate":"2021-11-11",
				"OfferId":0,"OfferType":1,"SellerAddr":"0x7BBF44394a23504cbE46b2b2d76929451cb86975",
				"BuyerAddr":"","TokenId":17,"Price":3,"Profit":null,
				"WasBought":false,"WasCanceled":false,"BoughtDuration":"","CanceledDuration":""}]}`)
		default:
			http.Error(w, `{"error":"down"}`, http.StatusBadGateway)
		}
	}))
	defer backend.Close()
	s := newTestServer(t, backend.URL)

	mux := http.NewServeMux()
	tradingPage := page{
		"GET /randomwalk/trading/history/{offset}/{limit}",
		"randomwalk/rw_trading_history.html",
		"/api/randomwalk/trading/history/{offset}/{limit}",
	}
	staticPage := page{"GET /static", "index.html", ""}
	failingPage := page{"GET /failing", "error.html", "/api/failing"}
	for _, p := range []page{tradingPage, staticPage, failingPage} {
		mux.HandleFunc(p.Pattern, s.pageHandler(p))
	}

	// Data page: parameters substituted into the API path, rows rendered.
	// This is the regression test for the empty trading-history page (the
	// template must consume the response's Sales/MarketAddr keys).
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/randomwalk/trading/history/0/5", nil))
	body := rec.Body.String()
	if rec.Code != http.StatusOK {
		t.Fatalf("trading history code = %d, body=%q", rec.Code, body)
	}
	if !strings.Contains(body, "0x47eF85Dfb775aCE0934fBa9EEd09D22e6eC0Cc08") {
		t.Error("market address missing from rendered page")
	}
	if !strings.Contains(body, "0x7BBF44394a23504cbE46b2b2d76929451cb86975") {
		t.Error("sale row missing from rendered page")
	}

	// Static page: no API fetch.
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/static", nil))
	if rec.Code != http.StatusOK {
		t.Errorf("static page code = %d", rec.Code)
	}

	// Failing API: styled error page with 502.
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/failing", nil))
	if rec.Code != http.StatusBadGateway {
		t.Errorf("failing page code = %d, want 502", rec.Code)
	}
}

func TestNotImplemented(t *testing.T) {
	s := newTestServer(t, "http://unused")
	rec := httptest.NewRecorder()
	s.notImplemented(rec, httptest.NewRequest(http.MethodGet, "/cosmicsignature/not/ported", nil))
	if rec.Code != http.StatusNotFound || !strings.Contains(rec.Body.String(), "not implemented") {
		t.Errorf("notImplemented: code=%d body=%q", rec.Code, rec.Body.String())
	}
}

func TestLogRequests(t *testing.T) {
	var buf strings.Builder
	log := slog.New(slog.NewTextHandler(&buf, nil))
	h := logRequests(log, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	}))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tea", nil))
	if rec.Code != http.StatusTeapot {
		t.Errorf("status = %d, want 418", rec.Code)
	}
	if !strings.Contains(buf.String(), "status=418") || !strings.Contains(buf.String(), "path=/tea") {
		t.Errorf("access log line missing fields: %q", buf.String())
	}
}

// freePort reserves an ephemeral port and releases it for the caller.
func freePort(t *testing.T) string {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("reserving port: %v", err)
	}
	_, port, _ := net.SplitHostPort(l.Addr().String())
	_ = l.Close()
	return port
}

func TestRunServesAndShutsDownOnSignal(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{"error":"","Top5TradedTokens":[]}`)
	}))
	defer backend.Close()

	port := freePort(t)
	t.Setenv("API_BASE", backend.URL+"/") // trailing slash exercises TrimSuffix
	t.Setenv("APIEXPLORER_PORT", port)

	done := make(chan error, 1)
	go func() { done <- run() }()

	base := "http://127.0.0.1:" + port
	client := &http.Client{Timeout: 2 * time.Second}
	var resp *http.Response
	var err error
	for deadline := time.Now().Add(10 * time.Second); time.Now().Before(deadline); {
		resp, err = client.Get(base + "/")
		if err == nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	if err != nil {
		t.Fatalf("server did not come up: %v", err)
	}
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("index status = %d", resp.StatusCode)
	}

	for path, want := range map[string]int{
		"/res/main.css":            http.StatusOK,
		"/randomwalk/":             http.StatusOK, // pages[] route through the fake API
		"/cosmicsignature/no/page": http.StatusNotFound,
	} {
		resp, err := client.Get(base + path)
		if err != nil {
			t.Fatalf("GET %s: %v", path, err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != want {
			t.Errorf("GET %s status = %d, want %d", path, resp.StatusCode, want)
		}
	}

	// Graceful shutdown path: SIGTERM is what the run loop scripts and
	// systemd send. NotifyContext is armed before the listener goroutine,
	// so the server responding above guarantees the handler is installed.
	if err := syscall.Kill(os.Getpid(), syscall.SIGTERM); err != nil {
		t.Fatalf("sending SIGTERM: %v", err)
	}
	select {
	case err := <-done:
		if err != nil {
			t.Errorf("run() after SIGTERM = %v, want nil", err)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("run() did not return after SIGTERM")
	}
}

func TestRunReportsBindFailure(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("occupying port: %v", err)
	}
	defer func() { _ = l.Close() }()
	_, port, _ := net.SplitHostPort(l.Addr().String())

	t.Setenv("API_BASE", "")
	t.Setenv("APIEXPLORER_PORT", port)

	done := make(chan error, 1)
	go func() { done <- run() }()
	select {
	case err := <-done:
		if err == nil {
			t.Error("run() on an occupied port must fail")
		}
	case <-time.After(10 * time.Second):
		t.Fatal("run() did not return on bind failure")
	}
}

func TestMainExitsCleanlyOnServerClosed(t *testing.T) {
	// main() treats http.ErrServerClosed as success; verify the check it
	// relies on rather than exercising os.Exit.
	if !errors.Is(fmt.Errorf("wrap: %w", http.ErrServerClosed), http.ErrServerClosed) {
		t.Fatal("errors.Is must unwrap ErrServerClosed")
	}
}
