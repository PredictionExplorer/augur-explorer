// First-fetch row capping. Some API endpoints return the full history
// (tens of thousands of rows); rendering all of it makes the page slow,
// heavy and pushes the usage footer out of reach. Every list of a decoded
// API response is therefore capped at maxTableRows before the template
// runs, and the layout shows a banner with the real row count and a link
// to the same page with ?full=1, which disables the cap.
package main

import (
	"net/http"
	"sort"
)

// maxTableRows is how many rows of each list a page shows on the first
// fetch; ?full=1 shows everything.
const maxTableRows = 200

// listCap records one capped list for the layout banner.
type listCap struct {
	Name         string // response field holding the list ("RatioHistory")
	Shown, Total int
}

// capLists truncates every list of the decoded API response that is longer
// than limit and reports what was capped. It recurses into nested objects
// (not into rows), so templates keep complete rows and real totals computed
// before the cap stay correct.
func capLists(data map[string]any, limit int) []listCap {
	capped := capListsIn("", data, limit)
	sort.Slice(capped, func(i, j int) bool { return capped[i].Name < capped[j].Name })
	return capped
}

func capListsIn(prefix string, m map[string]any, limit int) []listCap {
	var out []listCap
	for key, value := range m {
		name := key
		if prefix != "" {
			name = prefix + "." + key
		}
		switch t := value.(type) {
		case []any:
			if len(t) > limit {
				m[key] = t[:limit]
				out = append(out, listCap{Name: name, Shown: limit, Total: len(t)})
			}
		case map[string]any:
			out = append(out, capListsIn(name, t, limit)...)
		}
	}
	return out
}

// wantFull reports whether the request asks for the uncapped response.
func wantFull(r *http.Request) bool {
	return r.URL.Query().Get("full") != ""
}

// fullTableURL is the current request with full=1 added (other query
// parameters, e.g. the Bids page round picker, are preserved).
func fullTableURL(r *http.Request) string {
	q := r.URL.Query()
	q.Set("full", "1")
	return r.URL.Path + "?" + q.Encode()
}

// capResponse applies the first-fetch cap unless ?full=1 and stores the
// banner facts on the page metadata.
func capResponse(meta *pageMeta, data map[string]any, r *http.Request) {
	if wantFull(r) {
		return
	}
	if capped := capLists(data, maxTableRows); len(capped) > 0 {
		meta.Truncated = capped
		meta.FullURL = fullTableURL(r)
	}
}
