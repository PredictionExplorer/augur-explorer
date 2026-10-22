// The API explorer: a small HTML data viewer, separate from the API server.
// It renders the legacy "black" viewer look and feel, but instead of reading
// the database it fetches JSON from the API server (API_BASE) and maps the
// response onto the page templates. It needs no database credentials.
//
// URL scheme: /cosmicsignature/... and /randomwalk/... plus an index page
// with one button per project. Static assets are served under /res/.
package main

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"math"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

//go:embed templates
var templatesFS embed.FS

//go:embed res
var resFS embed.FS

const apiTimeout = 10 * time.Second

// page maps one explorer path to its template and the v1 API endpoint that
// feeds it. Path placeholders ({id} etc.) are substituted into the API URL.
type page struct {
	// Pattern is the ServeMux pattern ("GET /cosmicsignature/bids").
	Pattern string
	// Template is the file name under templates/.
	Template string
	// API is the API-server path; "{name}" segments are replaced with the
	// matching request path values.
	API string
}

// pages holds the project home pages and short aliases; the full legacy
// route table lives in routes_gen.go (generatedPages). The CosmicSignature
// dashboard aggregates several calls and has its own handler (dashboard.go).
var pages = []page{
	{"GET /randomwalk/{$}", "randomwalk/home.html", "/api/randomwalk/top5tokens"},
	// Short alias kept from the prototype (the /cosmicsignature/bids alias is
	// served by bidsHandler, see bids.go).
	{"GET /cosmicsignature/rounds", "cosmicsignature/cg_rounds.html", "/api/cosmicgame/rounds/list/0/100000"},
}

// jsonNum replaces every float64 decoded from API JSON. Its String method
// makes bare {{.X}} template output render without scientific notation
// (plain %v turns 1735747446 into "1.735747446e+09"), while the underlying
// float64 kind keeps printf verbs and template comparisons (gt/eq with float
// constants) working.
type jsonNum float64

func (n jsonNum) String() string {
	return formatJSONNum(float64(n))
}

// formatJSONNum renders a JSON number without scientific notation. The API
// uses math.MaxInt64 as an open-ended sentinel (e.g. NextEvtLogId), which
// float64 rounds up to 2^63 (9223372036854776000); print the exact int64
// maximum instead so links back into integer-parsing API routes stay valid.
func formatJSONNum(f float64) string {
	if f >= math.MaxInt64 {
		return "9223372036854775807"
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// convertNums rewrites a decoded JSON tree in place, wrapping numbers.
func convertNums(v any) any {
	switch t := v.(type) {
	case map[string]any:
		for k, vv := range t {
			t[k] = convertNums(vv)
		}
	case []any:
		for i, vv := range t {
			t[i] = convertNums(vv)
		}
	case float64:
		return jsonNum(t)
	}
	return v
}

// num renders a JSON number without scientific notation; kept for templates
// that format numbers explicitly.
func num(v any) string {
	switch f := v.(type) {
	case float64:
		return formatJSONNum(f)
	case jsonNum:
		return f.String()
	}
	return fmt.Sprint(v)
}

// addf sums JSON numbers; used where the legacy backend computed a total
// server-side that the v1 API does not return.
func addf(vals ...any) float64 {
	var sum float64
	for _, v := range vals {
		switch t := v.(type) {
		case jsonNum:
			sum += float64(t)
		case float64:
			sum += t
		}
	}
	return sum
}

// nowTs returns the current unix time in seconds; float64 so it composes
// with addf in templates.
func nowTs() float64 {
	return float64(time.Now().Unix())
}

// unixDate formats a unix-seconds JSON number as a local date-time string.
func unixDate(v any) string {
	var ts int64
	switch t := v.(type) {
	case jsonNum:
		ts = int64(t)
	case float64:
		ts = int64(t)
	default:
		return fmt.Sprint(v)
	}
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05 MST")
}

// weiToEth formats a decimal wei string (the v1 API returns exact amounts as
// strings) as an ETH value with 6 decimals. CST amounts use the same 18-decimal
// encoding, so this helper also formats CST wei.
func weiToEth(v any) string {
	f, _, err := big.ParseFloat(fmt.Sprint(v), 10, 128, big.ToNearestEven)
	if err != nil {
		return fmt.Sprint(v)
	}
	f.Quo(f, big.NewFloat(1e18))
	return f.Text('f', 6)
}

func asFloat(v any) float64 {
	switch t := v.(type) {
	case jsonNum:
		return float64(t)
	case float64:
		return t
	case int:
		return float64(t)
	case int64:
		return float64(t)
	default:
		f, err := strconv.ParseFloat(fmt.Sprint(v), 64)
		if err != nil {
			return 0
		}
		return f
	}
}

// isWeiPerSec reports whether v looks like a V3.1 CST bid-price decline
// multiplier (wei/second) rather than a V2 Dutch-auction duration in seconds.
// Durations are at most days-to-weeks; the V3.1 multiplier is ~1e16 wei/s.
func isWeiPerSec(v any) bool {
	return asFloat(v) >= 1e9
}

// The contract stores the CST bid-price decline rate per second; the explorer
// shows it per minute, which gives readable magnitudes (~0.76 CST/min rather
// than ~0.0127 CST/s).
const cstRateUnit = "CST/min"

// weiPerSecToCstPerMin converts a wei-per-second rate (string or number) to
// CST per minute with 6 decimals. Unparseable input is echoed back.
func weiPerSecToCstPerMin(v any) string {
	f, _, err := big.ParseFloat(fmt.Sprint(v), 10, 128, big.ToNearestEven)
	if err != nil {
		return fmt.Sprint(v)
	}
	f.Mul(f, big.NewFloat(60))
	f.Quo(f, big.NewFloat(1e18))
	return f.Text('f', 6)
}

// perMin converts a per-second float (already in CST, e.g. an admin event
// FloatValue) to per minute with 6 decimals.
func perMin(v any) string {
	return strconv.FormatFloat(asFloat(v)*60, 'f', 6, 64)
}

// cstDecline is auctionParam without the unit suffix, for dense tables whose
// header already says CST/min. V2 rows keep their "N s" duration so the two
// mechanics are not confused.
func cstDecline(durationStr any, durationInt any) string {
	if asFloat(durationInt) < 0 {
		return "—"
	}
	if isWeiPerSec(durationInt) {
		return weiPerSecToCstPerMin(durationStr)
	}
	return num(durationInt) + " s"
}

// shortTime renders an RFC 3339 timestamp as "MM-DD HH:MM:SS" (UTC); other
// strings are returned unchanged.
func shortTime(v any) string {
	s := fmt.Sprint(v)
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return s
	}
	return t.UTC().Format("01-02 15:04:05")
}

// auctionParam formats BidPlaced field 8. V1 stores -1; V2 stores CST Dutch
// auction duration in seconds; V3.1 stores cstBidPriceDeclineMultiplier (wei/s)
// in the same JSON keys (CstDutchAuctionDuration / CstDutchAuctionDurationInt).
func auctionParam(durationStr any, durationInt any) string {
	if asFloat(durationInt) < 0 {
		return "—"
	}
	if isWeiPerSec(durationInt) {
		return weiPerSecToCstPerMin(durationStr) + " " + cstRateUnit
	}
	return num(durationInt) + " s"
}

// isEthAddr reports whether s is a 42-char 0x-prefixed hex string.
func isEthAddr(s string) bool {
	if len(s) != 42 || s[0] != '0' || (s[1] != 'x' && s[1] != 'X') {
		return false
	}
	for i := 2; i < 42; i++ {
		c := s[i]
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
			return false
		}
	}
	return true
}

func isTxHash(s string) bool {
	if len(s) != 66 || s[0] != '0' || (s[1] != 'x' && s[1] != 'X') {
		return false
	}
	for i := 2; i < 66; i++ {
		c := s[i]
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
			return false
		}
	}
	return true
}

type server struct {
	apiBase   string
	templates *template.Template
	liveDir   string // non-empty: re-read templates/CSS from this dir each request
	client    *http.Client
	log       *slog.Logger
}

// fetch GETs apiBase+path and decodes the JSON body. The v1 API reports
// request-level failures in an "error" field with HTTP 200, so both transport
// and payload errors surface here.
func (s *server) fetch(ctx context.Context, path string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(ctx, apiTimeout)
	defer cancel()
	// #nosec G704 -- proxying to the operator-configured API server is this
	// tool's purpose; the base URL comes from API_BASE, only path segments
	// vary with the request.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.apiBase+path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req) // #nosec G704 -- see request construction above
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("API answered %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}
	var data map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding API response: %w", err)
	}
	if msg, ok := data["error"].(string); ok && msg != "" {
		return nil, errors.New("API error: " + msg)
	}
	convertNums(data)
	return data, nil
}

func (s *server) templateSet() (*template.Template, error) {
	if s.liveDir != "" {
		return loadTemplatesFrom(os.DirFS(s.liveDir))
	}
	return s.templates, nil
}

// render executes a page template inside the layout; see renderPage.
func (s *server) render(w http.ResponseWriter, status int, name string, data any) {
	s.renderPage(w, status, name, data, pageMeta{})
}

func (s *server) renderError(w http.ResponseWriter, status int, descr string) {
	s.render(w, status, "error.html", map[string]any{"ErrDescr": descr})
}

// renderFields renders the generic all-fields view of an API response.
func (s *server) renderFields(w http.ResponseWriter, data map[string]any, meta pageMeta) {
	meta.View = "fields"
	s.renderPage(w, http.StatusOK, "fields.html", map[string]any{
		"Fields": buildFieldsView(meta.Project, data),
		"API":    meta.API,
	}, meta)
}

// paramNames returns the {placeholder} names of an API path template.
func paramNames(api string) []string {
	var names []string
	for seg := range strings.SplitSeq(api, "/") {
		if strings.HasPrefix(seg, "{") && strings.HasSuffix(seg, "}") {
			names = append(names, seg[1:len(seg)-1])
		}
	}
	return names
}

// paramKey converts a path placeholder name to the template-data key it is
// exposed under ("round_num" -> "RoundNum").
func paramKey(name string) string {
	var b strings.Builder
	for part := range strings.SplitSeq(name, "_") {
		if part == "" {
			continue
		}
		b.WriteString(strings.ToUpper(part[:1]) + part[1:])
	}
	return b.String()
}

// injectParams adds request path values to the template data (under their
// CamelCase names) so page headers can show them when the API response does
// not echo them back. Response fields always win.
func injectParams(data map[string]any, api string, get func(string) string) {
	for _, name := range paramNames(api) {
		key := paramKey(name)
		if _, ok := data[key]; !ok {
			data[key] = get(name)
		}
	}
}

// pageHandler builds the standard fetch -> decode -> render handler. Pages
// with an empty API render statically.
func (s *server) pageHandler(p page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]any{}
		meta := pageMeta{Project: projectOf(r.URL.Path), Path: r.URL.Path}
		if p.API != "" {
			apiPath := p.API
			for _, name := range paramNames(p.API) {
				apiPath = strings.ReplaceAll(apiPath, "{"+name+"}", r.PathValue(name))
			}
			meta.API = apiPath
			var err error
			data, err = s.fetch(r.Context(), apiPath)
			if err != nil {
				s.log.Error("API fetch failed", "api_path", apiPath, "err", err)
				s.renderError(w, http.StatusBadGateway, err.Error())
				return
			}
			if r.URL.Query().Get("view") == "fields" {
				s.renderFields(w, data, meta)
				return
			}
			injectParams(data, p.API, r.PathValue)
		}
		s.renderPage(w, http.StatusOK, p.Template, data, meta)
	}
}

// notImplemented answers project sub-paths that exist in the legacy viewer
// but have no explorer page yet (the prototype covers a handful of pages).
func (s *server) notImplemented(w http.ResponseWriter, r *http.Request) {
	s.renderError(w, http.StatusNotFound,
		fmt.Sprintf("page %q is not implemented yet in the API explorer prototype", r.URL.Path))
}

// logRequests wraps mux with one INFO line per request.
func logRequests(log *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		log.Info("request", "method", r.Method, "path", r.URL.Path,
			"status", rec.status, "duration_ms", float64(time.Since(start).Microseconds())/1000)
	})
}

type statusRecorder struct {
	http.ResponseWriter

	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

// loadTemplates parses the embedded templates. Tests and production binaries
// that have no source tree next to them use this path.
func loadTemplates() (*template.Template, error) {
	return loadTemplatesFrom(templatesFS)
}

func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"ethAddrLink":     ethAddrLink,
		"ethAddrLinkTo":   ethAddrLinkTo,
		"isEthAddr":       isEthAddr,
		"arbiscanTx":      arbiscanTx,
		"num":             num,
		"weiToEth":        weiToEth,
		"addf":            addf,
		"nowTs":           nowTs,
		"unixDate":        unixDate,
		"isWeiPerSec":     isWeiPerSec,
		"auctionParam":    auctionParam,
		"cstPerMin":       weiPerSecToCstPerMin,
		"perMin":          perMin,
		"cstDecline":      cstDecline,
		"shortTime":       shortTime,
		"adminEventHelp":  adminEventHelp,
		"roundRangeLabel": roundRangeLabel,
		// Dashboard-style helpers (hover.go).
		"shortHex":    shortHex,
		"hexCell":     hexCell,
		"txHashCell":  txHashCell,
		"roundLink":   roundLink,
		"tokenLink":   tokenLink,
		"rwTokenLink": rwTokenLink,
		"recordCard":  recordCard,
		"jsonPretty":  jsonPretty,
		"durationStr": durationStr,
		"hasKey":      hasKey,
		"pct":         pct,
		"asFloat":     asFloat,
		// Labels (labels.go).
		"prizeTypeBadge":   prizeTypeBadge,
		"bidTypeBadge":     bidTypeBadge,
		"mechanicsVersion": mechanicsVersion,
	}
}

// loadTemplatesFrom parses every template under templates/ using names
// relative to that directory ("cosmicsignature/home.html"), because ParseFS
// would name the files by base name and the two project home.html files
// would collide.
func loadTemplatesFrom(rootFS fs.FS) (*template.Template, error) {
	root := template.New("").Funcs(templateFuncs())
	err := fs.WalkDir(rootFS, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		content, err := fs.ReadFile(rootFS, path)
		if err != nil {
			return err
		}
		_, err = root.New(strings.TrimPrefix(path, "templates/")).Parse(string(content))
		return err
	})
	return root, err
}

func isLiveAssetsDir(dir string) bool {
	st, err := os.Stat(filepath.Join(dir, "templates", "index.html"))
	return err == nil && !st.IsDir()
}

// liveAssetsDir returns a directory that contains templates/ and res/ to
// re-read on each request. Empty means use the embedded copies. Look next
// to the process (repo root or cmd/apiexplorer) so `go run` and a local
// binary pick up HTML/CSS edits without a restart.
func liveAssetsDir() string {
	if d := strings.TrimSpace(os.Getenv("APIEXPLORER_LIVE_DIR")); d != "" {
		if isLiveAssetsDir(d) {
			return d
		}
	}
	candidates := []string{".", "cmd/apiexplorer"}
	if exe, err := os.Executable(); err == nil {
		candidates = append(candidates, filepath.Dir(exe))
	}
	for _, c := range candidates {
		if isLiveAssetsDir(c) {
			return c
		}
	}
	return ""
}

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func run() error {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	apiBase := strings.TrimSuffix(os.Getenv("API_BASE"), "/")
	if apiBase == "" {
		apiBase = "http://127.0.0.1:9090"
	}
	port := os.Getenv("APIEXPLORER_PORT")
	if port == "" {
		port = "9091"
	}

	liveDir := liveAssetsDir()
	tmpl, err := loadTemplates()
	if err != nil {
		return fmt.Errorf("parsing templates: %w", err)
	}

	s := &server{
		apiBase:   apiBase,
		templates: tmpl,
		liveDir:   liveDir,
		client:    &http.Client{Timeout: apiTimeout},
		log:       logger,
	}

	mux := http.NewServeMux()
	if liveDir != "" {
		mux.Handle("GET /res/", noCache(http.FileServerFS(os.DirFS(liveDir))))
	} else {
		mux.Handle("GET /res/", http.FileServerFS(resFS))
	}
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		s.render(w, http.StatusOK, "index.html", nil)
	})
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, resFS, "res/favicon.ico")
	})
	mux.HandleFunc("GET /jump", s.jumpHandler)
	mux.HandleFunc("GET /api/", s.apiProxy)
	mux.HandleFunc("GET /cosmicsignature/{$}", s.dashboardHandler)
	mux.HandleFunc("GET /cosmicsignature/bid/list/all", s.bidsHandler)
	mux.HandleFunc("GET /cosmicsignature/bids", s.bidsHandler)
	for _, p := range pages {
		mux.HandleFunc(p.Pattern, s.pageHandler(p))
	}
	for _, p := range generatedPages {
		mux.HandleFunc(p.Pattern, s.pageHandler(p))
	}
	// Everything else under the two project prefixes is a legacy page that
	// is not ported yet; answer with the styled placeholder.
	mux.HandleFunc("GET /cosmicsignature/", s.notImplemented)
	mux.HandleFunc("GET /randomwalk/", s.notImplemented)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           logRequests(logger, mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	errCh := make(chan error, 1)
	go func() { errCh <- srv.ListenAndServe() }()
	if liveDir != "" {
		logger.Info("apiexplorer listening", "port", port, "api_base", apiBase, "live_assets", liveDir)
	} else {
		logger.Info("apiexplorer listening", "port", port, "api_base", apiBase, "live_assets", "embedded")
	}

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return srv.Shutdown(shutdownCtx)
	}
}

func main() {
	if err := run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintf(os.Stderr, "apiexplorer: %v\n", err)
		os.Exit(1)
	}
}
