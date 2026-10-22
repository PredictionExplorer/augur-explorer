package main

import (
	"context"
	"html"
	"html/template"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Page templates render only their content; render() wraps the result in
// templates/layout.html (sidebar, top bar, quick-jump, JSON drawer). A page
// template may start with <title>…</title>; the tag is lifted into the
// layout's <head>. When no layout template exists in the set (bare template
// directories in tests) the content is written as-is.

// pageMeta carries request facts into the layout.
type pageMeta struct {
	Project string // "cosmicsignature", "randomwalk" or ""
	Path    string // request path, highlights the sidebar entry
	API     string // resolved API path ("" for static pages)
	View    string // "" (page template) or "fields" (generic field view)
}

type layoutData struct {
	pageMeta
	Title   string
	Content template.HTML
	Nav     []navGroup
	APIBase string
	Now     string
}

var titleRe = regexp.MustCompile(`(?s)^\s*<title>(.*?)</title>\s*`)

// splitTitle lifts a leading <title> element out of rendered page content.
func splitTitle(content string) (title, rest string) {
	m := titleRe.FindStringSubmatchIndex(content)
	if m == nil {
		return "", content
	}
	title = html.UnescapeString(strings.TrimSpace(content[m[2]:m[3]]))
	return title, content[m[1]:]
}

// projectOf derives the project from a template name or request path
// ("cosmicsignature/cg_bids.html", "/randomwalk/...").
func projectOf(s string) string {
	s = strings.TrimPrefix(s, "/")
	if i := strings.IndexByte(s, '/'); i >= 0 {
		s = s[:i]
	}
	switch s {
	case "cosmicsignature", "randomwalk":
		return s
	}
	return ""
}

// renderPage executes the page template and wraps it in the layout.
func (s *server) renderPage(w http.ResponseWriter, status int, name string, data any, meta pageMeta) {
	tmpl, err := s.templateSet()
	if err != nil {
		s.log.Error("template parse failed", "template", name, "err", err)
		http.Error(w, "template parse failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var buf strings.Builder
	if err := tmpl.ExecuteTemplate(&buf, name, data); err != nil {
		s.log.Error("template render failed", "template", name, "err", err)
		http.Error(w, "template render failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	out := buf.String()
	if layout := tmpl.Lookup("layout.html"); layout != nil {
		title, content := splitTitle(out)
		if meta.Project == "" {
			meta.Project = projectOf(name)
		}
		ld := layoutData{
			pageMeta: meta,
			Title:    title,
			Content:  template.HTML(content), // #nosec G203 -- output of html/template
			Nav:      projectNav(meta.Project),
			APIBase:  s.apiBase,
			Now:      time.Now().UTC().Format("2006-01-02 15:04:05 UTC"),
		}
		var page strings.Builder
		if err := layout.Execute(&page, ld); err != nil {
			s.log.Error("layout render failed", "template", name, "err", err)
			http.Error(w, "layout render failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		out = page.String()
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	_, _ = io.WriteString(w, out)
}

// jumpTarget resolves the quick-jump query into an explorer (or Arbiscan)
// URL. Accepted forms:
//
//	0x + 40 hex   user page
//	0x + 64 hex   Arbiscan transaction
//	N             round info
//	b N           bid by evtlog id
//	t N           CosmicSignature token
//	rw N          RandomWalk token history
//	a N           CST stake action
//	e N           admin events from evtlog id N
//
// The empty string means "not understood".
func jumpTarget(q string) string {
	q = strings.TrimSpace(q)
	if isEthAddr(q) {
		return "/cosmicsignature/user/info/" + q
	}
	if isTxHash(q) {
		return "https://arbiscan.io/tx/" + q
	}
	q = strings.TrimPrefix(q, "#")
	if _, err := strconv.ParseInt(q, 10, 64); err == nil {
		return "/cosmicsignature/round/info/" + q
	}
	prefix, rest, ok := strings.Cut(q, " ")
	if !ok {
		return ""
	}
	rest = strings.TrimSpace(strings.TrimPrefix(rest, "#"))
	if _, err := strconv.ParseInt(rest, 10, 64); err != nil {
		return ""
	}
	switch strings.ToLower(prefix) {
	case "r", "round":
		return "/cosmicsignature/round/info/" + rest
	case "b", "bid":
		return "/cosmicsignature/bid/info/" + rest
	case "t", "tok", "token":
		return "/cosmicsignature/cst/info/" + rest
	case "rw", "rwalk":
		return "/randomwalk/tokens/history/" + rest
	case "a", "action":
		return "/cosmicsignature/staking/cst/actions/info/" + rest
	case "e", "evt", "events":
		return "/cosmicsignature/system/admin_events/" + rest + "/9223372036854775807"
	}
	return ""
}

// jumpHandler redirects the quick-jump form.
func (s *server) jumpHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	target := jumpTarget(q)
	if target == "" {
		s.renderError(w, http.StatusBadRequest,
			"quick-jump: cannot interpret "+strconv.Quote(q)+
				" (use an address, a tx hash, a round number, or b/t/rw/a/e followed by an id)")
		return
	}
	http.Redirect(w, r, target, http.StatusFound)
}

// apiProxy forwards GET /api/... to the API server so the JSON drawer and
// the raw-JSON links stay same-origin. Only the path is forwarded.
func (s *server) apiProxy(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), apiTimeout)
	defer cancel()
	// #nosec G704 -- apiBase is operator configuration; only the request
	// path (validated by the mux to start with /api/) is appended.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.apiBase+r.URL.Path, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := s.client.Do(req) // #nosec G704 -- see above
	if err != nil {
		http.Error(w, "API unreachable: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer func() { _ = resp.Body.Close() }()
	ct := resp.Header.Get("Content-Type")
	if ct == "" {
		ct = "application/json; charset=utf-8"
	}
	w.Header().Set("Content-Type", ct)
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}
