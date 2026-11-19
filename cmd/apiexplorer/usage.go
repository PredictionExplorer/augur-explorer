// The per-page usage footer. Every data page ends with a "How to use this
// request" section aimed at API developers: the explorer URL pattern, the
// current request as a working example, the backing API-server endpoint and
// a plain-language explanation of every path parameter. The content is
// derived from the route tables (routes_gen.go, pages, czPages), so new
// routes get a footer automatically; TestParamDocsCoverRoutes fails the
// build when a route introduces a placeholder that paramDocs does not
// explain.
package main

import (
	"net/http"
	"strings"
)

// usageParam documents one {placeholder} of an explorer route.
type usageParam struct {
	Name, Desc string
}

// paramDocs explains every path placeholder used by the route tables, for
// the usage footer. Keep the wording developer-facing: what to put there
// and where to find such a value.
var paramDocs = map[string]string{
	"user_addr":     "wallet address: 0x followed by 40 hex digits",
	"user_aid":      "numeric address id, as returned in *Aid fields of other responses",
	"round_num":     "bidding round number, starting at 0",
	"prize_num":     "prize number, starting at 0 (same numbering as rounds)",
	"bid_position":  "position of the bid inside its round, starting at 0",
	"evtlog_id":     "event-log id: unique per indexed on-chain event (EvtLogId fields); bids are identified by it",
	"evtlog_start":  "event-log id the range starts at (inclusive)",
	"evtlog_end":    "event-log id the range ends at (inclusive)",
	"token_id":      "NFT token id",
	"token_addr":    "NFT contract address: 0x followed by 40 hex digits",
	"name":          "token name (or a part of it) to search for",
	"record_id":     "record id, as returned by the matching list endpoint (RecordId fields)",
	"action_id":     "staking action id (ActionId fields)",
	"deposit_id":    "staking reward deposit id (DepositId fields)",
	"offset":        "how many records to skip; 0 starts at the first record",
	"limit":         "maximum number of records to return",
	"sort":          "sort direction: 0 oldest first, 1 newest first",
	"order_by":      "ordering: 1 price high to low, 2 price low to high, anything else insertion order",
	"init_ts":       "start of the period, as a Unix timestamp in seconds",
	"fin_ts":        "end of the period, as a Unix timestamp in seconds",
	"interval_secs": "length of one aggregation bucket in seconds (86400 = daily)",
}

// usageInfo feeds the footer in templates/layout.html.
type usageInfo struct {
	Origin     string       // scheme://host of the running explorer
	Pattern    string       // explorer path with {param} placeholders
	Example    string       // the current request: a working example
	APIPattern string       // backing API-server path with placeholders ("" = none)
	Params     []usageParam // explanation of each placeholder, in path order
	Notes      string       // optional extra, e.g. query-string parameters
}

// requestOrigin rebuilds scheme://host for the usage footer from the
// request, so examples are copy-pasteable wherever the explorer runs.
func requestOrigin(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}

// patternPath strips the method and the trailing {$} of a ServeMux pattern.
func patternPath(pattern string) string {
	if _, path, ok := strings.Cut(pattern, " "); ok {
		pattern = path
	}
	return strings.TrimSuffix(pattern, "{$}")
}

// usageFor derives the footer for a route-table page and the request being
// served. Static pages (no API behind them) get no footer.
func usageFor(p page, r *http.Request) *usageInfo {
	if p.API == "" {
		return nil
	}
	path := patternPath(p.Pattern)
	u := &usageInfo{
		Origin:     requestOrigin(r),
		Pattern:    path,
		Example:    r.URL.Path,
		APIPattern: p.API,
	}
	for _, name := range paramNames(path) {
		u.Params = append(u.Params, usageParam{Name: name, Desc: paramDocs[name]})
	}
	return u
}
