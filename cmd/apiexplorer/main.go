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
	"html"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"math/big"
	"net/http"
	"os"
	"os/signal"
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
// route table lives in routes_gen.go (generatedPages).
var pages = []page{
	{"GET /cosmicsignature/{$}", "cosmicsignature/cg_index.html", "/api/cosmicgame/statistics/dashboard"},
	{"GET /randomwalk/{$}", "randomwalk/home.html", "/api/randomwalk/top5tokens"},
	// Short aliases kept from the prototype.
	{"GET /cosmicsignature/bids", "cosmicsignature/cg_bids.html", "/api/cosmicgame/bid/list/all/0/100000"},
	{"GET /cosmicsignature/rounds", "cosmicsignature/cg_rounds.html", "/api/cosmicgame/rounds/list/0/100000"},
}

// jsonNum replaces every float64 decoded from API JSON. Its String method
// makes bare {{.X}} template output render without scientific notation
// (plain %v turns 1735747446 into "1.735747446e+09"), while the underlying
// float64 kind keeps printf verbs and template comparisons (gt/eq with float
// constants) working.
type jsonNum float64

func (n jsonNum) String() string {
	return strconv.FormatFloat(float64(n), 'f', -1, 64)
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
	if f, ok := v.(float64); ok {
		return strconv.FormatFloat(f, 'f', -1, 64)
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
// strings) as an ETH value with 6 decimals.
func weiToEth(v any) string {
	f, _, err := big.ParseFloat(fmt.Sprint(v), 10, 128, big.ToNearestEven)
	if err != nil {
		return fmt.Sprint(v)
	}
	f.Quo(f, big.NewFloat(1e18))
	return f.Text('f', 6)
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

// ethAddrLink renders a user-info link for valid Ethereum addresses and
// escaped plain text otherwise (e.g. labels like "(All CS NFT Stakers)").
func ethAddrLink(addr string) template.HTML {
	return ethAddrLinkTo("/cosmicsignature/user/info/", addr)
}

// ethAddrLinkTo renders a pathPrefix+addr link for valid Ethereum addresses
// and escaped plain text otherwise.
func ethAddrLinkTo(pathPrefix, addr string) template.HTML {
	esc := html.EscapeString(addr)
	if !isEthAddr(addr) {
		return template.HTML(esc) // #nosec G203 -- esc is HTML-escaped above
	}
	// #nosec G203 -- pathPrefix is a compile-time constant from the template
	// and esc is HTML-escaped above.
	return template.HTML(fmt.Sprintf(`<a href="%s%s">%s</a>`, pathPrefix, esc, esc))
}

type server struct {
	apiBase   string
	templates *template.Template
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

func (s *server) render(w http.ResponseWriter, status int, name string, data any) {
	var buf strings.Builder
	if err := s.templates.ExecuteTemplate(&buf, name, data); err != nil {
		s.log.Error("template render failed", "template", name, "err", err)
		http.Error(w, "template render failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	_, _ = io.WriteString(w, buf.String())
}

func (s *server) renderError(w http.ResponseWriter, status int, descr string) {
	s.render(w, status, "error.html", map[string]any{"ErrDescr": descr})
}

// paramNames returns the {placeholder} names of an API path template.
func paramNames(api string) []string {
	var names []string
	for _, seg := range strings.Split(api, "/") {
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
	for _, part := range strings.Split(name, "_") {
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
		if p.API != "" {
			apiPath := p.API
			for _, name := range paramNames(p.API) {
				apiPath = strings.ReplaceAll(apiPath, "{"+name+"}", r.PathValue(name))
			}
			var err error
			data, err = s.fetch(r.Context(), apiPath)
			if err != nil {
				s.log.Error("API fetch failed", "api_path", apiPath, "err", err)
				s.renderError(w, http.StatusBadGateway, err.Error())
				return
			}
			injectParams(data, p.API, r.PathValue)
		}
		s.render(w, http.StatusOK, p.Template, data)
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

// loadTemplates parses every embedded template under its path relative to
// templates/ ("cosmicsignature/home.html"), because ParseFS would name the
// files by base name and the two project home.html files would collide.
func loadTemplates() (*template.Template, error) {
	root := template.New("").Funcs(template.FuncMap{
		"ethAddrLink":   ethAddrLink,
		"ethAddrLinkTo": ethAddrLinkTo,
		"isEthAddr":     isEthAddr,
		"num":           num,
		"weiToEth":      weiToEth,
		"addf":          addf,
		"nowTs":         nowTs,
		"unixDate":      unixDate,
	})
	err := fs.WalkDir(templatesFS, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		content, err := templatesFS.ReadFile(path)
		if err != nil {
			return err
		}
		_, err = root.New(strings.TrimPrefix(path, "templates/")).Parse(string(content))
		return err
	})
	return root, err
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

	tmpl, err := loadTemplates()
	if err != nil {
		return fmt.Errorf("parsing templates: %w", err)
	}

	s := &server{
		apiBase:   apiBase,
		templates: tmpl,
		client:    &http.Client{Timeout: apiTimeout},
		log:       logger,
	}

	mux := http.NewServeMux()
	mux.Handle("GET /res/", http.FileServerFS(resFS))
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		s.render(w, http.StatusOK, "index.html", nil)
	})
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
	logger.Info("apiexplorer listening", "port", port, "api_base", apiBase)

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
