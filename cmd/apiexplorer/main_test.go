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
	"path/filepath"
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

func TestAuctionParam(t *testing.T) {
	if got := auctionParam("-1", jsonNum(-1)); got != "—" {
		t.Errorf("V1 sentinel = %q", got)
	}
	if got := auctionParam("1800", jsonNum(1800)); got != "1800 s" {
		t.Errorf("V2 duration = %q", got)
	}
	// 16666666666666666 wei/s = 0.016667 CST/s = 1 CST/min.
	if got := auctionParam("16666666666666666", jsonNum(16666666666666666)); got != "1.000000 CST/min" {
		t.Errorf("V3.1 decline multiplier = %q", got)
	}
	// Live value at the time of writing: 12740065401983209 wei/s.
	if got := weiPerSecToCstPerMin("12740065401983209"); got != "0.764404" {
		t.Errorf("weiPerSecToCstPerMin = %q, want 0.764404", got)
	}
	if got := weiPerSecToCstPerMin("not-wei"); got != "not-wei" {
		t.Errorf("weiPerSecToCstPerMin(invalid) = %q", got)
	}
	// Admin event FloatValue is already CST/s.
	if got := perMin(jsonNum(0.0127400654)); got != "0.764404" {
		t.Errorf("perMin = %q, want 0.764404", got)
	}
	// Unit-less variant for dense tables.
	for _, c := range []struct{ s, i, want string }{
		{"-1", "-1", "—"}, {"1800", "1800", "1800 s"}, {"16666666666666666", "16666666666666666", "1.000000"},
	} {
		if got := cstDecline(c.s, jsonNum(asFloat(c.i))); got != c.want {
			t.Errorf("cstDecline(%s) = %q, want %q", c.s, got, c.want)
		}
	}
}

func TestShortTime(t *testing.T) {
	if got := shortTime("2026-10-22T16:48:42Z"); got != "10-22 16:48:42" {
		t.Errorf("shortTime = %q", got)
	}
	if got := shortTime("2026-10-22T18:48:42+02:00"); got != "10-22 16:48:42" {
		t.Errorf("shortTime(offset) = %q, want UTC", got)
	}
	if got := shortTime("n/a"); got != "n/a" {
		t.Errorf("shortTime(invalid) = %q", got)
	}
	if isWeiPerSec(jsonNum(1800)) {
		t.Error("1800 s must not be treated as wei/s")
	}
	if !isWeiPerSec(jsonNum(1e16)) {
		t.Error("1e16 must be treated as wei/s")
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
	for _, want := range []string{
		fmt.Sprintf(`<a href="/cosmicsignature/user/info/%s">`, addr), // main link
		`<span class="hx-short">0x7BBF44…6975</span>`,                 // abbreviated label
		`<span class="hx-full">` + addr + `</span>`,                   // full label (body.full-hex)
		`data-copy="` + addr + `"`,                                    // copy button
		`/cosmicsignature/staking/cst/actions/by_user/` + addr,        // related-page link in the card
	} {
		if !strings.Contains(got, want) {
			t.Errorf("ethAddrLink missing %q in %q", want, got)
		}
	}
	if got := string(ethAddrLinkTo("/randomwalk/user/info/", "<label>")); got != "&lt;label&gt;" {
		t.Errorf("invalid address must render escaped text, got %q", got)
	}
}

func TestArbiscanTx(t *testing.T) {
	hash := "0xf000000000000000000000000000000000000000000000000000000000001009"
	got := string(arbiscanTx(hash, jsonNum(12345)))
	for _, want := range []string{
		`<a href="https://arbiscan.io/tx/` + hash + `" title="` + hash + `">12345</a>`,
		`data-copy="` + hash + `"`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("arbiscanTx missing %q in %q", want, got)
		}
	}
	if got := string(arbiscanTx("", 99)); got != "99" {
		t.Errorf("empty hash must render plain label, got %q", got)
	}
	if got := string(arbiscanTx("<script>", "<x>")); got != "&lt;x&gt;" {
		t.Errorf("invalid hash must escape label, got %q", got)
	}
	if got := string(txHashCell(hash)); !strings.Contains(got, `<span class="hx-short">0xf00000…1009</span>`) {
		t.Errorf("txHashCell must abbreviate the hash, got %q", got)
	}
	if got := string(txHashCell("")); got != "" {
		t.Errorf("txHashCell(empty) = %q", got)
	}
}

func TestShortHex(t *testing.T) {
	cases := map[string]string{
		"0x7BBF44394a23504cbE46b2b2d76929451cb86975": "0x7BBF44…6975",
		"0x1234":     "0x1234",
		"plain text": "plain text",
	}
	for in, want := range cases {
		if got := shortHex(in); got != want {
			t.Errorf("shortHex(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestHoverHelpers(t *testing.T) {
	if got := string(roundLink(jsonNum(7))); !strings.Contains(got, `href="/cosmicsignature/round/info/7"`) ||
		!strings.Contains(got, `/cosmicsignature/bid/list/all?round=7`) {
		t.Errorf("roundLink = %q", got)
	}
	if got := string(tokenLink(jsonNum(-1))); got != `<span class="empty">—</span>` {
		t.Errorf("tokenLink(-1) = %q", got)
	}
	if got := string(tokenLink(jsonNum(12))); !strings.Contains(got, `href="/cosmicsignature/cst/info/12"`) {
		t.Errorf("tokenLink = %q", got)
	}
	if got := string(rwTokenLink(jsonNum(3))); !strings.Contains(got, `href="/randomwalk/tokens/history/3"`) {
		t.Errorf("rwTokenLink = %q", got)
	}
	if got := string(hexCell("0xabcdef0123456789")); !strings.Contains(got, `0xabcdef…6789`) {
		t.Errorf("hexCell = %q", got)
	}
	if got := string(hexCell("short")); got != "short" {
		t.Errorf("hexCell(short) = %q", got)
	}
	if got := string(recordCard(map[string]any{"A": 1})); !strings.Contains(got, `&#34;A&#34;: 1`) {
		t.Errorf("recordCard = %q", got)
	}
	if got := durationStr(jsonNum(90061)); got != "1d 01h 01m 01s" {
		t.Errorf("durationStr = %q", got)
	}
	if got := durationStr(jsonNum(-65)); got != "-1m 05s" {
		t.Errorf("durationStr(neg) = %q", got)
	}
	if got := durationStr(jsonNum(3700)); got != "1h 01m 40s" {
		t.Errorf("durationStr(h) = %q", got)
	}
	if got := durationStr(jsonNum(5)); got != "5s" {
		t.Errorf("durationStr(s) = %q", got)
	}
	if !hasKey(map[string]any{"x": nil}, "x") || hasKey(map[string]any{}, "x") {
		t.Error("hasKey")
	}
	if got := pct(jsonNum(0.256)); got != "25.6%" {
		t.Errorf("pct = %q", got)
	}
	if got := string(prizeTypeBadge(jsonNum(7))); !strings.Contains(got, "Chrono Warrior ETH") {
		t.Errorf("prizeTypeBadge = %q", got)
	}
	if got := string(prizeTypeBadge(jsonNum(99))); !strings.Contains(got, "type 99") {
		t.Errorf("prizeTypeBadge(unknown) = %q", got)
	}
	if got := string(bidTypeBadge(jsonNum(2))); !strings.Contains(got, "CST") {
		t.Errorf("bidTypeBadge = %q", got)
	}
	for in, want := range map[float64]string{1: "V1", 2: "V2", 3: "V3.1", 9: "unknown"} {
		if got := mechanicsVersion(jsonNum(in)); got != want {
			t.Errorf("mechanicsVersion(%v) = %q", in, got)
		}
	}
}

func TestSplitTitle(t *testing.T) {
	title, rest := splitTitle("\n<title>Bids &amp; Rounds</title>\n<div>x</div>")
	if title != "Bids & Rounds" || rest != "<div>x</div>" {
		t.Errorf("splitTitle = %q, %q", title, rest)
	}
	title, rest = splitTitle("<div>no title</div>")
	if title != "" || rest != "<div>no title</div>" {
		t.Errorf("splitTitle(no title) = %q, %q", title, rest)
	}
}

func TestProjectOf(t *testing.T) {
	cases := map[string]string{
		"cosmicsignature/cg_bids.html": "cosmicsignature",
		"/randomwalk/tokens/info/1":    "randomwalk",
		"index.html":                   "",
		"/":                            "",
	}
	for in, want := range cases {
		if got := projectOf(in); got != want {
			t.Errorf("projectOf(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestJumpTarget(t *testing.T) {
	addr := "0x7BBF44394a23504cbE46b2b2d76929451cb86975"
	hash := "0xf000000000000000000000000000000000000000000000000000000000001009"
	cases := map[string]string{
		addr:         "/cosmicsignature/user/info/" + addr,
		" " + hash:   "https://arbiscan.io/tx/" + hash,
		"5":          "/cosmicsignature/round/info/5",
		"#5":         "/cosmicsignature/round/info/5",
		"r 5":        "/cosmicsignature/round/info/5",
		"b 31603":    "/cosmicsignature/bid/info/31603",
		"t #12":      "/cosmicsignature/cst/info/12",
		"rw 3":       "/randomwalk/tokens/history/3",
		"a 9":        "/cosmicsignature/staking/cst/actions/info/9",
		"e 1820":     "/cosmicsignature/system/admin_events/1820/9223372036854775807",
		"":           "",
		"hello":      "",
		"b x":        "",
		"zz 1":       "",
		"round five": "",
	}
	for in, want := range cases {
		if got := jumpTarget(in); got != want {
			t.Errorf("jumpTarget(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestJumpHandler(t *testing.T) {
	s := newTestServer(t, "http://unused")
	rec := httptest.NewRecorder()
	s.jumpHandler(rec, httptest.NewRequest(http.MethodGet, "/jump?q=7", nil))
	if rec.Code != http.StatusFound || rec.Header().Get("Location") != "/cosmicsignature/round/info/7" {
		t.Errorf("jump: code=%d location=%q", rec.Code, rec.Header().Get("Location"))
	}
	rec = httptest.NewRecorder()
	s.jumpHandler(rec, httptest.NewRequest(http.MethodGet, "/jump?q=nonsense", nil))
	if rec.Code != http.StatusBadRequest || !strings.Contains(rec.Body.String(), "quick-jump") {
		t.Errorf("bad jump: code=%d body=%q", rec.Code, rec.Body.String())
	}
}

func TestRenderWrapsLayout(t *testing.T) {
	s := newTestServer(t, "http://api.example")
	rec := httptest.NewRecorder()
	s.renderPage(rec, http.StatusOK, "cosmicsignature/cg_bids.html",
		map[string]any{"Bids": []any{}, "AllRounds": true, "SelectedRound": 3, "CurRoundNum": 3, "RoundChoices": []int{3, 2, 1, 0}},
		pageMeta{Project: "cosmicsignature", Path: "/cosmicsignature/bid/list/all", API: "/api/cosmicgame/bid/list/all/0/100000"})
	body := rec.Body.String()
	for _, want := range []string{
		"<title>Bids · API Explorer</title>",                  // lifted page title
		`class="sidebar"`,                                     // nav rendered for the project
		`href="/cosmicsignature/bid/list/all" class="active"`, // current page highlighted
		`href="/cosmicsignature/bid/list/all?view=fields"`,    // fields toggle
		`data-api="/api/cosmicgame/bid/list/all/0/100000"`,    // JSON drawer target
		`<select name="round"`,                                // page content (round picker)
	} {
		if !strings.Contains(body, want) {
			t.Errorf("layout output missing %q", want)
		}
	}
	if strings.Count(body, "<title>") != 1 {
		t.Errorf("page <title> must be removed from the content, got %d title tags", strings.Count(body, "<title>"))
	}
	// Index page: no project, no sidebar.
	rec = httptest.NewRecorder()
	s.render(rec, http.StatusOK, "index.html", nil)
	if strings.Contains(rec.Body.String(), `class="sidebar"`) {
		t.Error("index page must not render a project sidebar")
	}
}

func TestBuildFieldsView(t *testing.T) {
	addr := "0x7BBF44394a23504cbE46b2b2d76929451cb86975"
	hash := "0xf000000000000000000000000000000000000000000000000000000000001009"
	data := map[string]any{
		"error":  "",
		"status": jsonNum(1),
		"Offset": jsonNum(0),
		"Stats":  map[string]any{"TotalBids": jsonNum(3), "Addr": addr},
		"Rows": []any{
			map[string]any{
				"Tx":       map[string]any{"EvtLogId": jsonNum(10), "BlockNum": jsonNum(5), "TxHash": hash, "TimeStamp": jsonNum(1700000000)},
				"RoundNum": jsonNum(2), "WinnerAddr": addr, "TokenId": jsonNum(4), "Claimed": true,
				"Message": "", "Nested": []any{jsonNum(1)}, "Zebra": "z",
				"Long": strings.Repeat("x", 80),
			},
			map[string]any{"Tx": map[string]any{"EvtLogId": jsonNum(11)}, "Extra": nil, "Claimed": false},
		},
		"Scalars": []any{jsonNum(1), jsonNum(2)},
	}
	fv := buildFieldsView("cosmicsignature", data)
	if len(fv.Scalars) != 2 { // Offset + Scalars (array of scalars)
		t.Fatalf("scalars = %+v", fv.Scalars)
	}
	if len(fv.Objects) != 1 || fv.Objects[0].Name != "Stats" || len(fv.Objects[0].KV) != 2 {
		t.Fatalf("objects = %+v", fv.Objects)
	}
	if len(fv.Tables) != 1 {
		t.Fatalf("tables = %+v", fv.Tables)
	}
	tbl := fv.Tables[0]
	// Identity columns first, then alphabetical within the same priority.
	wantPrefix := []string{"Tx.EvtLogId", "Tx.BlockNum", "Tx.TxHash", "Tx.TimeStamp", "RoundNum", "TokenId", "WinnerAddr"}
	for i, w := range wantPrefix {
		if i >= len(tbl.Cols) || tbl.Cols[i] != w {
			t.Fatalf("cols = %v, want prefix %v", tbl.Cols, wantPrefix)
		}
	}
	if tbl.Cols[len(tbl.Cols)-1] != "Zebra" {
		t.Errorf("last column = %q, want Zebra", tbl.Cols[len(tbl.Cols)-1])
	}
	if len(tbl.Rows) != 2 || len(tbl.Rows[0].Cells) != len(tbl.Cols) || len(tbl.Rows[1].Cells) != len(tbl.Cols) {
		t.Fatalf("rows malformed: %+v", tbl.Rows)
	}
	joined := ""
	for _, c := range tbl.Rows[0].Cells {
		joined += string(c)
	}
	for _, want := range []string{
		`href="/cosmicsignature/round/info/2"`,
		`href="/cosmicsignature/cst/info/4"`,
		`/cosmicsignature/user/info/` + addr,
		`https://arbiscan.io/tx/` + hash,
		`<span class="ok">✓</span>`,
		`<span class="empty">—</span>`,
		`class="trunc"`,
	} {
		if !strings.Contains(joined, want) {
			t.Errorf("row cells missing %q", want)
		}
	}
	if !strings.Contains(string(tbl.Rows[0].Card), "WinnerAddr") {
		t.Error("record card must contain the full row JSON")
	}
	// Missing keys in the second row render as empty cells; null renders as "null".
	second := ""
	for _, c := range tbl.Rows[1].Cells {
		second += string(c)
	}
	if !strings.Contains(second, `<span class="empty">null</span>`) || !strings.Contains(second, `<span class="empty">✗</span>`) {
		t.Errorf("second row = %q", second)
	}
	// RandomWalk project: token ids go to the RandomWalk token page and
	// addresses are not user links.
	rw := buildFieldsView("randomwalk", map[string]any{"Rows": []any{map[string]any{"TokenId": jsonNum(1), "Addr": addr}}})
	rwJoined := string(rw.Tables[0].Rows[0].Cells[0]) + string(rw.Tables[0].Rows[0].Cells[1])
	if !strings.Contains(rwJoined, "/randomwalk/tokens/history/1") || strings.Contains(rwJoined, "/user/info/") {
		t.Errorf("randomwalk cells = %q", rwJoined)
	}
	// Arrays of non-objects are not tables.
	if _, ok := buildTable("cosmicsignature", "x", []any{jsonNum(1)}); ok {
		t.Error("scalar array must not become a table")
	}
	if tbl, ok := buildTable("cosmicsignature", "x", nil); !ok || tbl.Name != "x" {
		t.Error("empty array must become an empty table")
	}
}

func TestColPriority(t *testing.T) {
	order := []string{"EvtLogId", "BlockNum", "TxHash", "DateTime", "TimeStamp", "TxId", "RoundNum", "BidderAid", "WinnerAddr", "AmountEth", "NumBids", "Message"}
	for i := 1; i < len(order); i++ {
		if colPriority(order[i-1]) > colPriority(order[i]) {
			t.Errorf("%s must sort before %s", order[i-1], order[i])
		}
	}
}

func TestFieldsViewHandler(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{"error":"","status":1,"Rows":[{"EvtLogId":1,"Note":"hello"}]}`)
	}))
	defer backend.Close()
	s := newTestServer(t, backend.URL)
	mux := http.NewServeMux()
	p := page{"GET /cosmicsignature/x", "cosmicsignature/cg_bids.html", "/api/x"}
	mux.HandleFunc(p.Pattern, s.pageHandler(p))

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/cosmicsignature/x?view=fields", nil))
	body := rec.Body.String()
	if rec.Code != http.StatusOK {
		t.Fatalf("fields view code = %d body=%q", rec.Code, body)
	}
	for _, want := range []string{"<th>EvtLogId</th>", "<th>Note</th>", "hello", `href="/cosmicsignature/x"`} {
		if !strings.Contains(body, want) {
			t.Errorf("fields view missing %q", want)
		}
	}
	if strings.Contains(body, "<h1>All Bids</h1>") {
		t.Error("fields view must not render the page template")
	}
}

func TestAPIProxy(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/cosmicgame/x" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"ok":true}`)
	}))
	defer backend.Close()
	s := newTestServer(t, backend.URL)

	rec := httptest.NewRecorder()
	s.apiProxy(rec, httptest.NewRequest(http.MethodGet, "/api/cosmicgame/x", nil))
	if rec.Code != http.StatusOK || rec.Body.String() != `{"ok":true}` || rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("proxy: code=%d body=%q ct=%q", rec.Code, rec.Body.String(), rec.Header().Get("Content-Type"))
	}
	rec = httptest.NewRecorder()
	s.apiProxy(rec, httptest.NewRequest(http.MethodGet, "/api/missing", nil))
	if rec.Code != http.StatusNotFound {
		t.Errorf("proxy must pass the upstream status, got %d", rec.Code)
	}
	closed := httptest.NewServer(http.NotFoundHandler())
	closed.Close()
	s.apiBase = closed.URL
	rec = httptest.NewRecorder()
	s.apiProxy(rec, httptest.NewRequest(http.MethodGet, "/api/x", nil))
	if rec.Code != http.StatusBadGateway {
		t.Errorf("proxy against closed backend code = %d, want 502", rec.Code)
	}
}

func TestDashboardHandler(t *testing.T) {
	dash := `{"error":"","CurRoundNum":3,"BidPriceEth":0.004,"PrizeAmountEth":8.06,"CurNumBids":356,
		"LastBidderAddr":"0xefAEBf3801bBfdd2aC0bFF7c72058eE2E7e5787f","ContractMechanicsVersion":3,
		"CgPrizeRowCount":161,"TsRoundStart":1792627209,"PrizeClaimDate":"x","TimeoutClaimPrize":172800,
		"RaffleAmountEth":1,"StakingAmountEth":1,"CosmicGameBalanceEth":1,"CharityBalance":"1","CharityBalanceEth":1,
		"TokenReward":"5273539803740201488","PrizePercentage":25,"RafflePercentage":4,"ChronoWarriorPercentage":7,
		"StakingPercentage":10,"CharityPercentage":10,"PriceIncrease":"100","TimeIncrease":"100",
		"MainPrizeTimeIncrementInMicroSeconds":3600,"InitialSecondsUntilPrize":41667,"RoundStartCSTAuctionLength":1250,
		"CstDutchAuctionDurationChangeDivisor":250,"NumRaffleEthWinnersBidding":3,"NumRaffleNFTWinnersBidding":10,
		"NumRaffleNFTWinnersStakingRWalk":10,"CharityAddr":"0x96bB0ADB414d5350f435E52f94946B6C7A0760a9",
		"CurRoundStats":{"ActivationTime":1792627200,"TotalEthInBidsEth":0.3,"TotalCstInBidsEth":1413,"TotalDonatedCount":1,"TotalDonatedAmountEth":20},
		"V3Config":{"IsV3":true,"MainPrizeNumCosmicSignatureNfts":3,"CstBidPriceDeclineMultiplier":"13523836144179803",
		"CstBidPriceDeclineMultiplierChangeDivisor":"100","CstAuctionPriceMinLimit":"1000000000000000000",
		"CstRewardPerTimeIncrement":"60000000000000000000","RoundLateBidDurationSeconds":1236,"RoundLateBidDurationDivisor":"3000000",
		"RoundLateBidPremiumBaseMultiplier":"29228998656","RoundLateBidPremiumExponent":8},
		"MainStats":{"TotalBids":3651,"NumBidsCST":1505,"NumUniqueBidders":43,"NumUniqueWinners":14,"TotalPrizes":3,"TotalPrizeAwards":159,
		"TotalPrizesPaidAmountEth":36.5,"TotalCSTGivenInPrizes":72000,"TotalCSTConsumedEth":280522,"NumCSTokenMints":72,"TotalNamedTokens":3,
		"TotalRaffleEthDeposits":4.3,"TotalRaffleEthWithdrawn":13,"TotalChronoWarriorEthDeposits":8.7,"NumCosmicGameDonations":3,
		"SumCosmicGameDonationsEth":7.6,"NumDirectDonations":3,"DirectDonationsEth":50,"NumVoluntaryDonations":0,"SumVoluntaryDonationsEth":0,
		"NumDonatedNFTs":4,"NumUniqueDonors":3,"NumRwalkTokensUsed":8,"NumUniqueStakersCST":8,"NumUniqueStakersRWalk":10,"NumUniqueStakersBoth":5,
		"StakeStatisticsCST":{"TotalTokensStaked":39,"NumDeposits":2,"TotalRewardEth":5.04},
		"StakeStatisticsRWalk":{"TotalTokensStaked":85,"TotalTokensMinted":30},"NumMktRewards":8,"TotalMktRewardsEth":5999},
		"ContractAddrs":{"CosmicGameAddr":"0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2","ImplementationAddr":"0x7739148013777c485AD9f3d971e1005Eca686661",
		"CosmicSignatureAddr":"0xbb84Be3500A63581d3F2d5AC3bdF8685AAedad25","CosmicTokenAddr":"0xAD91843e6A58Ba560F577E676986AFb1dba6FBA0",
		"CosmicDaoAddr":"0xF3D52E1c681949be7E624778dB13DaD7F8c729db","PrizesWalletAddr":"0xE1b619e9B39ea4109D2F429Ea5eAA307759b0011",
		"CharityWalletAddr":"0x96bB0ADB414d5350f435E52f94946B6C7A0760a9","StakingWalletCSTAddr":"0x6308A405B4FF1eA890870Efe2a6D036750B81F7C",
		"StakingWalletRWalkAddr":"0x5EB3396092841E6c5b0b51141699F6711E830529","MarketingWalletAddr":"0xa3802c799f5e3D3D3562A9B513a41C6aAF92e25e",
		"RandomWalkAddr":"0x895a6F444BE4ba9d124F61DF736605792B35D66b","MarketplaceAddr":"0x47eF85Dfb775aCE0934fBa9EEd09D22e6eC0Cc08"}}`
	// Two rows, newest first: the second exists only to give the first a
	// predecessor (SinceLastBid = 457 s = 7m 37s).
	bids := `{"error":"","Bids":[{"Tx":{"EvtLogId":31603,"BlockNum":517281980,"TxHash":"0x2a9e8143367df4fdbbb3fb9924e25764f6d834bfa687d3b6e252b0e5f73a13d5","DateTime":"2026-10-22T16:48:42Z","TimeStamp":1792687722},
		"BidderAddr":"0xefAEBf3801bBfdd2aC0bFF7c72058eE2E7e5787f","EthPriceEth":-1e-18,"CstPriceEth":8.19,"RWalkNFTId":-1,"RoundNum":3,"BidType":2,"BidPosition":358,"BidCstRewardAmountEth":7.3,
		"CstDutchAuctionDuration":"12740065401983209","CstDutchAuctionDurationInt":12740065401983209},
		{"Tx":{"EvtLogId":31600,"BlockNum":517281000,"TxHash":"0x2a9e8143367df4fdbbb3fb9924e25764f6d834bfa687d3b6e252b0e5f73a13d5","DateTime":"2026-10-22T16:41:05Z","TimeStamp":1792687265},
		"BidderAddr":"0xefAEBf3801bBfdd2aC0bFF7c72058eE2E7e5787f","EthPriceEth":-1e-18,"CstPriceEth":8.1,"RWalkNFTId":-1,"RoundNum":3,"BidType":2,"BidPosition":357,"BidCstRewardAmountEth":7.3,
		"CstDutchAuctionDuration":"12740065401983209","CstDutchAuctionDurationInt":12740065401983209}]}`
	prizes := `{"error":"","GlobalPrizeHistory":[{"Tx":{"EvtLogId":30131,"BlockNum":511892555,"TxHash":"","DateTime":"2026-10-05T10:14:37Z"},
		"RecordType":0,"RoundNum":2,"AmountEth":9.97,"TokenId":-1,"Claimed":true,"IsTimeoutClaim":false,"WinnerAddr":"0xA169574D0d353E3010997A3E64846b7D1B2a63B6"}]}`
	modes := `{"error":"","SystemModeChanges":[{"EvtLogId":30131,"BlockNum":511892555,"TxHash":"","DateTime":"2026-10-05T10:14:37Z","NextEvtLogId":30591,"RoundNum":3}]}`

	var failTails bool
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if failTails && r.URL.Path != dashboardAPI {
			http.Error(w, "down", http.StatusInternalServerError)
			return
		}
		switch r.URL.Path {
		case dashboardAPI:
			_, _ = io.WriteString(w, dash)
		case recentBidsAPI:
			_, _ = io.WriteString(w, bids)
		case "/api/cosmicgame/prizes/history/global/0/10":
			_, _ = io.WriteString(w, prizes)
		case "/api/cosmicgame/system/modelist/0/5":
			_, _ = io.WriteString(w, modes)
		case "/api/cosmicgame/time/until_prize":
			_, _ = io.WriteString(w, `{"TimeUntilPrize":1352548,"error":"","status":1}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer backend.Close()
	s := newTestServer(t, backend.URL)

	rec := httptest.NewRecorder()
	s.dashboardHandler(rec, httptest.NewRequest(http.MethodGet, "/cosmicsignature/", nil))
	body := rec.Body.String()
	if rec.Code != http.StatusOK {
		t.Fatalf("dashboard code = %d body=%q", rec.Code, body)
	}
	for _, want := range []string{
		"15d 15h 42m 28s",                        // UntilPrize tail
		`href="/cosmicsignature/bid/info/31603"`, // bids tail
		"Main Prize ETH",                         // prizes tail
		"30591",                                  // modes tail
		"0x6a714A…63F2",                          // contracts panel (abbreviated)
		"V3.1",
		`>10-22 16:48:42</a>`, // bids tail: short time links to the record
		`title="2026-10-22T16:48:42Z · evtlog 31603"`,
		`<td class="num">0.764404</td>`,                 // per-bid CST decline, CST/min
		`<td class="num">7m 37s</td>`,                   // since previous bid
		`0.811430</span><span class="u">CST/min</span>`, // headline: 13523836144179803 wei/s × 60
	} {
		if !strings.Contains(body, want) {
			t.Errorf("dashboard missing %q", want)
		}
	}

	// Tail failures degrade to "unavailable" notes instead of an error page.
	failTails = true
	rec = httptest.NewRecorder()
	s.dashboardHandler(rec, httptest.NewRequest(http.MethodGet, "/cosmicsignature/", nil))
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "bid tail unavailable") {
		t.Errorf("dashboard with failing tails: code=%d", rec.Code)
	}

	// Fields view of the dashboard payload.
	rec = httptest.NewRecorder()
	s.dashboardHandler(rec, httptest.NewRequest(http.MethodGet, "/cosmicsignature/?view=fields", nil))
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "<h2>MainStats</h2>") {
		t.Errorf("dashboard fields view: code=%d", rec.Code)
	}

	// Dashboard endpoint down: error page.
	s.apiBase = "http://127.0.0.1:1"
	rec = httptest.NewRecorder()
	s.dashboardHandler(rec, httptest.NewRequest(http.MethodGet, "/cosmicsignature/", nil))
	if rec.Code != http.StatusBadGateway {
		t.Errorf("dashboard with API down code = %d, want 502", rec.Code)
	}
}

func TestBidsHandler(t *testing.T) {
	row := func(evt, round int) string {
		return fmt.Sprintf(`{"Tx":{"EvtLogId":%d,"BlockNum":517281980,"TxHash":"0x2a9e8143367df4fdbbb3fb9924e25764f6d834bfa687d3b6e252b0e5f73a13d5","DateTime":"2026-10-22T16:48:42Z"},
		"BidderAddr":"0xefAEBf3801bBfdd2aC0bFF7c72058eE2E7e5787f","EthPriceEth":0.01,"CstPriceEth":8.19,"CSTRewardEth":0,"RWalkNFTId":-1,"RoundNum":%d,"BidType":0,"BidPosition":1,"BidCstRewardAmountEth":7.3,"Message":""}`, evt, round)
	}
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case dashboardAPI:
			_, _ = io.WriteString(w, `{"error":"","CurRoundNum":3}`)
		case "/api/cosmicgame/bid/list/by_round/3/1/0/100000":
			_, _ = io.WriteString(w, `{"error":"","TotalRows":1,"BidsByRound":[`+row(31603, 3)+`]}`)
		case "/api/cosmicgame/bid/list/by_round/1/1/0/100000":
			_, _ = io.WriteString(w, `{"error":"","TotalRows":1,"BidsByRound":[`+row(20001, 1)+`]}`)
		case "/api/cosmicgame/bid/list/all/0/100000":
			_, _ = io.WriteString(w, `{"error":"","Bids":[`+row(31603, 3)+`,`+row(20001, 1)+`]}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer backend.Close()
	s := newTestServer(t, backend.URL)

	get := func(q string) (int, string) {
		rec := httptest.NewRecorder()
		s.bidsHandler(rec, httptest.NewRequest(http.MethodGet, "/cosmicsignature/bid/list/all"+q, nil))
		return rec.Code, rec.Body.String()
	}

	// Default: the current round, picker shows it selected plus an "all" option.
	code, body := get("")
	if code != http.StatusOK {
		t.Fatalf("bids code = %d body=%q", code, body)
	}
	for _, want := range []string{
		`<option value="3" selected>3 (current)</option>`,
		`<option value="0">0</option>`,
		`<option value="all">all rounds</option>`,
		`<title>Bids · round 3 · API Explorer</title>`,
		`href="/cosmicsignature/bid/info/31603"`,
		`CST Δsupply`,
		`>+7.3000</td>`,         // ETH bid: reward minted, nothing burned
		`+7.3000</span> minted`, // totals line
		`−0.0000</span> burned`,
		`net <b class="ok">+7.3000</b> CST`,
	} {
		if !strings.Contains(body, want) {
			t.Errorf("bids page missing %q", want)
		}
	}
	if strings.Contains(body, "20001") {
		t.Error("bids page for round 3 must not list round-1 bids")
	}

	// Explicit round.
	code, body = get("?round=1")
	if code != http.StatusOK || !strings.Contains(body, `<option value="1" selected>1</option>`) ||
		!strings.Contains(body, "20001") || strings.Contains(body, "31603") {
		t.Errorf("bids ?round=1: code=%d", code)
	}

	// All rounds.
	code, body = get("?round=all")
	if code != http.StatusOK || !strings.Contains(body, `<option value="all" selected>`) ||
		!strings.Contains(body, "20001") || !strings.Contains(body, "31603") ||
		strings.Contains(body, `<option value="3" selected>`) {
		t.Errorf("bids ?round=all: code=%d", code)
	}

	// Junk / out-of-range rounds fall back to the current round.
	for _, q := range []string{"?round=99", "?round=-1", "?round=abc"} {
		if _, body := get(q); !strings.Contains(body, `<option value="3" selected>`) {
			t.Errorf("%s must fall back to the current round", q)
		}
	}

	// Fields view uses the by-round payload.
	if code, body := get("?round=1&view=fields"); code != http.StatusOK || !strings.Contains(body, "<h2>BidsByRound <span") {
		t.Errorf("bids fields view: code=%d", code)
	}

	// Backend down: error page.
	s.apiBase = "http://127.0.0.1:1"
	if code, _ := get(""); code != http.StatusBadGateway {
		t.Errorf("bids with API down: code=%d, want 502", code)
	}
}

func TestAnnotateBidGaps(t *testing.T) {
	row := func(ts float64) map[string]any {
		return map[string]any{"Tx": map[string]any{"TimeStamp": jsonNum(ts)}}
	}
	gap := func(m map[string]any) string {
		v, ok := m["SinceLastBid"]
		if !ok {
			return "none"
		}
		return num(v)
	}

	// Newest first (API order): each row pairs with the next; the oldest has none.
	desc := []any{row(1000), row(940), row(100)}
	annotateBidGaps(desc)
	if got := []string{gap(desc[0].(map[string]any)), gap(desc[1].(map[string]any)), gap(desc[2].(map[string]any))}; strings.Join(got, ",") != "60,840,none" {
		t.Errorf("desc gaps = %v", got)
	}

	// Oldest first works too.
	asc := []any{row(100), row(940), row(1000)}
	annotateBidGaps(asc)
	if got := []string{gap(asc[0].(map[string]any)), gap(asc[1].(map[string]any)), gap(asc[2].(map[string]any))}; strings.Join(got, ",") != "none,840,60" {
		t.Errorf("asc gaps = %v", got)
	}

	// Rows without timestamps are skipped, not a panic.
	odd := []any{row(10), map[string]any{"Tx": map[string]any{}}, "junk"}
	annotateBidGaps(odd)
	if gap(odd[0].(map[string]any)) != "none" {
		t.Error("row without a timestamped neighbour must have no gap")
	}
}

func TestAnnotateCstSupply(t *testing.T) {
	rows := []any{
		// CST bid: mints 7.5 to the previous bidder, burns the 8.2 price.
		map[string]any{"BidType": jsonNum(2), "BidCstRewardAmountEth": jsonNum(7.5), "CstPriceEth": jsonNum(8.2)},
		// ETH bid: mints only (CstPriceEth is the -1e-18 sentinel).
		map[string]any{"BidType": jsonNum(0), "BidCstRewardAmountEth": jsonNum(10), "CstPriceEth": jsonNum(-1e-18)},
		// V1 bid: BidCstRewardAmount is -1, fall back to CSTRewardEth.
		map[string]any{"BidType": jsonNum(0), "BidCstRewardAmountEth": jsonNum(-1), "CSTRewardEth": jsonNum(100), "CstPriceEth": jsonNum(-1e-18)},
		// First bid of a round: nothing minted.
		map[string]any{"BidType": jsonNum(0), "BidCstRewardAmountEth": jsonNum(0), "CstPriceEth": jsonNum(-1e-18)},
		"junk",
	}
	tot := annotateCstSupply(rows)
	if tot.Minted != 117.5 || tot.Burned != 8.2 || fmt.Sprintf("%.4f", tot.Net()) != "109.3000" {
		t.Errorf("totals = %+v net %.4f", tot, tot.Net())
	}
	want := []string{"-0.7000", "+10.0000", "+100.0000", "+0.0000"}
	for i, w := range want {
		if got := string(signedCst(rows[i].(map[string]any)["CstSupplyDelta"])); got != w {
			t.Errorf("row %d delta = %s, want %s", i, got, w)
		}
	}
}

func TestBidsRound(t *testing.T) {
	cases := []struct {
		raw  string
		cur  int
		want int
	}{
		{"", 3, 3}, {"2", 3, 2}, {"0", 3, 0}, {"3", 3, 3},
		{"4", 3, 3}, {"-1", 3, 3}, {"x", 3, 3}, {"1", 0, 0},
	}
	for _, c := range cases {
		if got := bidsRound(c.raw, c.cur); got != c.want {
			t.Errorf("bidsRound(%q, %d) = %d, want %d", c.raw, c.cur, got, c.want)
		}
	}
}

func TestProjectNav(t *testing.T) {
	if projectNav("") != nil {
		t.Error("index must have no nav")
	}
	seen := map[string]bool{}
	for _, project := range []string{"cosmicsignature", "randomwalk"} {
		for _, g := range projectNav(project) {
			if g.Title == "" || len(g.Items) == 0 {
				t.Errorf("%s: empty nav group %+v", project, g)
			}
			for _, it := range g.Items {
				if !strings.HasPrefix(it.Href, "/"+project+"/") {
					t.Errorf("%s: item %q links outside the project: %s", project, it.Label, it.Href)
				}
				if g.Param != "" && !strings.Contains(it.Href, "{"+g.Param+"}") {
					t.Errorf("%s: item %q lacks {%s}", project, it.Label, g.Param)
				}
				if g.Param == "" && strings.Contains(it.Href, "{") {
					t.Errorf("%s: plain item %q has a placeholder: %s", project, it.Label, it.Href)
				}
				seen[it.Href] = true
			}
		}
	}
	if !seen["/cosmicsignature/bid/list/all"] || !seen["/randomwalk/trading/history"] {
		t.Error("expected core tables in the nav")
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
		"cosmicsignature/cg_roi_leaderboard.html",
		"cosmicsignature/cg_claims_by_round.html",
		"cosmicsignature/cg_claim_detail_by_round.html",
		"cosmicsignature/cg_bid_info.html",
		"cosmicsignature/cg_cosmic_sig_token_list.html",
		"cosmicsignature/cg_current_special_winners.html",
	} {
		if tmpl.Lookup(name) == nil {
			t.Errorf("template %q not loaded", name)
		}
	}
}

func TestLiveAssetsReload(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "templates"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "res"), 0o755); err != nil {
		t.Fatal(err)
	}
	page := filepath.Join(dir, "templates", "index.html")
	if err := os.WriteFile(page, []byte("before"), 0o644); err != nil {
		t.Fatal(err)
	}
	s := &server{liveDir: dir, log: slog.New(slog.DiscardHandler)}
	rec := httptest.NewRecorder()
	s.render(rec, http.StatusOK, "index.html", nil)
	if rec.Body.String() != "before" {
		t.Fatalf("first render = %q", rec.Body.String())
	}
	if err := os.WriteFile(page, []byte("after"), 0o644); err != nil {
		t.Fatal(err)
	}
	rec = httptest.NewRecorder()
	s.render(rec, http.StatusOK, "index.html", nil)
	if rec.Body.String() != "after" {
		t.Fatalf("reload render = %q", rec.Body.String())
	}
}

func TestIsLiveAssetsDir(t *testing.T) {
	if !isLiveAssetsDir(".") {
		t.Error("package dir should contain templates/index.html")
	}
	if isLiveAssetsDir(t.TempDir()) {
		t.Error("empty dir must not count as live assets")
	}
}

func TestV31TemplatesExecute(t *testing.T) {
	tmpl, err := loadTemplates()
	if err != nil {
		t.Fatalf("loadTemplates: %v", err)
	}
	tmpl.Option("missingkey=error")

	cases := []struct {
		name string
		data map[string]any
	}{
		{
			name: "cosmicsignature/cg_bid_info.html",
			data: map[string]any{
				"BidInfo": map[string]any{
					"Tx": map[string]any{
						"EvtLogId": jsonNum(10), "BlockNum": jsonNum(1), "TxId": jsonNum(2),
						"TxHash": "0xabc", "TimeStamp": jsonNum(1), "DateTime": "t",
					},
					"RoundNum": jsonNum(0), "BidderAddr": "0x1", "BidType": jsonNum(0),
					"BidPosition": jsonNum(1), "EthPriceEth": jsonNum(0.1),
					"PreviousBidderAddr":         "0x2",
					"PreviousCstRewardAmountEth": jsonNum(1.5),
					"ThisCstRewardAmountEth":     jsonNum(0),
					"RWalkNFTId":                 jsonNum(-1), "CstPriceEth": jsonNum(-1),
					"BidCstRewardAmountEth":      jsonNum(1.5),
					"CstDutchAuctionDuration":    "16666666666666666",
					"CstDutchAuctionDurationInt": jsonNum(16666666666666666),
					"PrizeTimeDate":              "soon", "Message": "hi",
					"NFTDonationTokenAddr": "", "DonatedERC20TokenAddr": "",
				},
			},
		},
		{
			name: "cosmicsignature/cg_cosmic_sig_token_list.html",
			data: map[string]any{
				"CosmicSignatureTokenList": []any{
					map[string]any{
						"TokenId": jsonNum(1), "WinnerAddr": "0x1", "CurOwnerAddr": "0x1",
						"RoundNum": jsonNum(0), "RecordType": jsonNum(3),
						"Tx": map[string]any{"DateTime": "t"},
					},
				},
			},
		},
		{
			name: "cosmicsignature/cg_system_admin_events_in_range.html",
			data: map[string]any{
				"EvtLogIdStart": jsonNum(1820), "EvtLogIdEnd": jsonNum(1904),
				"RoundNumStart": jsonNum(4), "RoundNumEnd": jsonNum(5),
				"AdminEvents": []any{
					map[string]any{
						"RecordType": jsonNum(21), // has help text
						"TxHash":     "0xabc", "BlockNum": jsonNum(1),
						"DateTime": "t", "IntegerValue": jsonNum(3600),
						"FloatValue": nil, "AddressValue": "",
						"StringValue": "", "ResolvedValue": "1 h",
					},
					map[string]any{
						"RecordType": jsonNum(0), // no help text: no balloon
						"TxHash":     "0xdef", "BlockNum": jsonNum(2),
						"DateTime": "t", "IntegerValue": nil,
						"FloatValue": nil, "AddressValue": "",
						"StringValue": "x", "ResolvedValue": nil,
					},
				},
			},
		},
		{
			name: "cosmicsignature/cg_current_special_winners.html",
			data: map[string]any{
				"LastBidderAddress": "0x1", "LastBidderLastBidTime": jsonNum(1),
				"EnduranceChampionAddress": "0x2", "EnduranceChampionDuration": jsonNum(10),
				"EnduranceChampionStartTimeStamp": jsonNum(1),
				"PrevEnduranceChampionDuration":   jsonNum(5),
				"ChronoWarriorAddress":            "0x3", "ChronoWarriorDuration": jsonNum(8),
				"ChronoWarriorIsLive":  true,
				"LastCstBidderAddress": "0x4",
				"RoundNum":             jsonNum(1), "SourceBlockNumber": jsonNum(9),
				"SourceBlockTimeStamp": jsonNum(1),
			},
		},
	}
	for _, tc := range cases {
		if err := tmpl.ExecuteTemplate(io.Discard, tc.name, tc.data); err != nil {
			t.Errorf("%s: %v", tc.name, err)
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
	// run() binds the wildcard address, so the port has to be occupied
	// there too. A loopback-only listener does not collide with a wildcard
	// bind on macOS, which let the server start and leave run() blocked in
	// its signal select until the test timed out.
	l, err := net.Listen("tcp", ":0")
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
