package main

import (
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"strings"
)

// Hover cards ("balloons") keep tables narrow without hiding anything: the
// visible cell shows a short form, the card (CSS :hover, click to pin) holds
// the full value, a copy button and links to related pages.
//
// Markup contract shared with explorer.js / cosmicgame.css:
//
//	<span class="hx"><a ...><span class="hx-short">0xab…cd</span><span class="hx-full">0xab...cd</span></a>
//	  <span class="hx-card">...</span></span>
//
// body.full-hex swaps the short/full spans for the whole page.

// shortHex abbreviates a 0x-prefixed hex string to 0x + 6 + … + 4 chars.
// Other strings are returned unchanged.
func shortHex(s string) string {
	if len(s) <= 14 || !strings.HasPrefix(strings.ToLower(s), "0x") {
		return s
	}
	return s[:8] + "…" + s[len(s)-4:]
}

// copyBtn renders the clipboard button used inside hover cards.
func copyBtn(value string) string {
	return fmt.Sprintf(`<button type="button" class="hx-copy" data-copy="%s" title="copy">⧉</button>`,
		html.EscapeString(value))
}

// hexLabel renders the short/full pair for a hex string.
func hexLabel(s string) string {
	esc := html.EscapeString(s)
	return fmt.Sprintf(`<span class="hx-short">%s</span><span class="hx-full">%s</span>`,
		html.EscapeString(shortHex(s)), esc)
}

// addrCardLinks lists the per-user pages shown in an address card.
var addrCardLinks = []struct{ Label, Href string }{
	{"user info", "/cosmicsignature/user/info/%s"},
	{"balances", "/cosmicsignature/user/balances/%s"},
	{"prizes", "/cosmicsignature/prizes/history/by_user/%s/0/100000"},
	{"ETH prizes", "/cosmicsignature/prizes/eth/user/%s"},
	{"CS NFTs", "/cosmicsignature/cst/list/by_user/%s"},
	{"CT transfers", "/cosmicsignature/ct/transfers/by_user/%s"},
	{"staking CST", "/cosmicsignature/staking/cst/actions/by_user/%s"},
	{"staking RWalk", "/cosmicsignature/staking/randomwalk/actions/by_user/%s"},
	{"donations", "/cosmicsignature/donations/eth/by_user/%s"},
}

// ethAddrLink renders an address as a short link to the user-info page with
// a hover card (full address, copy, related pages). Non-addresses (labels
// like "(All CS NFT Stakers)") render as escaped plain text.
func ethAddrLink(addr string) template.HTML {
	return ethAddrLinkTo("/cosmicsignature/user/info/", addr)
}

// ethAddrLinkTo is ethAddrLink with a custom target page prefix.
func ethAddrLinkTo(pathPrefix, addr string) template.HTML {
	if !isEthAddr(addr) {
		return template.HTML(html.EscapeString(addr)) // #nosec G203 -- escaped above
	}
	esc := html.EscapeString(addr)
	var b strings.Builder
	b.WriteString(`<span class="hx hx-addr">`)
	fmt.Fprintf(&b, `<a href="%s%s">%s</a>`, pathPrefix, esc, hexLabel(addr))
	b.WriteString(`<span class="hx-card"><span class="hx-row mono">`)
	b.WriteString(esc)
	b.WriteString(copyBtn(addr))
	b.WriteString(`</span><span class="hx-links">`)
	for _, l := range addrCardLinks {
		fmt.Fprintf(&b, `<a href="`+l.Href+`">%s</a>`, esc, l.Label)
	}
	fmt.Fprintf(&b, `<a href="https://arbiscan.io/address/%s" target="_blank" rel="noopener">arbiscan ↗</a>`, esc)
	b.WriteString(`</span></span></span>`)
	// #nosec G203 -- every interpolated value is HTML-escaped above and
	// pathPrefix is a compile-time constant from the template.
	return template.HTML(b.String())
}

// arbiscanTx links label to the Arbitrum One explorer transaction page when
// hash looks like a 0x-prefixed 32-byte hex, with a hover card holding the
// full hash. Invalid/empty hashes (local Hardhat, missing fields) render as
// escaped plain text.
func arbiscanTx(hash any, label any) template.HTML {
	lab := html.EscapeString(fmt.Sprint(label))
	h := strings.TrimSpace(fmt.Sprint(hash))
	if !isTxHash(h) {
		return template.HTML(lab) // #nosec G203 -- lab is HTML-escaped above
	}
	return txLink(h, lab)
}

// txLink builds the transaction hover widget; h must be a valid tx hash and
// lab already HTML-escaped.
func txLink(h, lab string) template.HTML {
	esc := html.EscapeString(h)
	var b strings.Builder
	b.WriteString(`<span class="hx hx-tx">`)
	fmt.Fprintf(&b, `<a href="https://arbiscan.io/tx/%s" title="%s">%s</a>`, esc, esc, lab)
	fmt.Fprintf(&b, `<span class="hx-card"><span class="hx-row mono">%s%s</span>`, esc, copyBtn(h))
	fmt.Fprintf(&b, `<span class="hx-links"><a href="https://arbiscan.io/tx/%s" target="_blank" rel="noopener">arbiscan ↗</a></span>`, esc)
	b.WriteString(`</span></span>`)
	return template.HTML(b.String()) // #nosec G203 -- esc and lab are HTML-escaped
}

// txHashCell renders a bare transaction hash (short form, hover card) when
// no block number is available as the label.
func txHashCell(hash any) template.HTML {
	h := strings.TrimSpace(fmt.Sprint(hash))
	if !isTxHash(h) {
		return template.HTML(html.EscapeString(h)) // #nosec G203 -- escaped
	}
	return txLink(h, hexLabel(h))
}

// roundCardLinks lists the per-round pages shown in a round card.
var roundCardLinks = []struct{ Label, Href string }{
	{"info", "/cosmicsignature/round/info/%s"},
	{"bids", "/cosmicsignature/bid/list/all?round=%s"},
	{"claims", "/cosmicsignature/statistics/claims/detail/%s"},
	{"ETH prizes", "/cosmicsignature/prizes/eth/round/%s"},
	{"raffle NFTs", "/cosmicsignature/raffle/nft/by_round/%s"},
	{"staking", "/cosmicsignature/staking/cst/rewards/by_round/%s"},
	{"donated ETH", "/cosmicsignature/donations/eth/both/by_round/%s"},
	{"donated ERC20", "/cosmicsignature/donations/erc20/by_round/all/%s"},
	{"donated NFTs", "/cosmicsignature/donations/nft/by_round/%s"},
}

// roundLink renders a round number linking to the round page with a hover
// card listing every per-round view.
func roundLink(v any) template.HTML {
	n := html.EscapeString(num(v))
	var b strings.Builder
	b.WriteString(`<span class="hx hx-round">`)
	fmt.Fprintf(&b, `<a href="/cosmicsignature/round/info/%s">%s</a>`, n, n)
	b.WriteString(`<span class="hx-card"><span class="hx-links">`)
	for _, l := range roundCardLinks {
		fmt.Fprintf(&b, `<a href="`+l.Href+`">%s</a>`, n, l.Label)
	}
	b.WriteString(`</span></span></span>`)
	return template.HTML(b.String()) // #nosec G203 -- n is HTML-escaped
}

// tokenLink renders a CosmicSignature NFT id linking to the token page with
// a hover card (name history, transfers). Negative ids (no token) render as
// an empty marker.
func tokenLink(v any) template.HTML {
	if asFloat(v) < 0 {
		return template.HTML(`<span class="empty">—</span>`)
	}
	n := html.EscapeString(num(v))
	var b strings.Builder
	b.WriteString(`<span class="hx hx-token">`)
	fmt.Fprintf(&b, `<a href="/cosmicsignature/cst/info/%s">%s</a>`, n, n)
	fmt.Fprintf(&b, `<span class="hx-card"><span class="hx-links"><a href="/cosmicsignature/cst/info/%s">info</a>`, n)
	fmt.Fprintf(&b, `<a href="/cosmicsignature/cst/names/history/%s">names</a>`, n)
	fmt.Fprintf(&b, `<a href="/cosmicsignature/cst/transfers/all/%s">transfers</a></span></span></span>`, n)
	return template.HTML(b.String()) // #nosec G203 -- n is HTML-escaped
}

// rwTokenLink renders a RandomWalk NFT id linking to its history page.
func rwTokenLink(v any) template.HTML {
	if asFloat(v) < 0 {
		return template.HTML(`<span class="empty">—</span>`)
	}
	n := html.EscapeString(num(v))
	var b strings.Builder
	b.WriteString(`<span class="hx hx-token">`)
	fmt.Fprintf(&b, `<a href="/randomwalk/tokens/history/%s">%s</a>`, n, n)
	fmt.Fprintf(&b, `<span class="hx-card"><span class="hx-links"><a href="/randomwalk/tokens/info/%s">info</a>`, n)
	fmt.Fprintf(&b, `<a href="/randomwalk/tokens/history/%s">history</a>`, n)
	fmt.Fprintf(&b, `<a href="/randomwalk/tokens/name_changes/%s">names</a></span></span></span>`, n)
	return template.HTML(b.String()) // #nosec G203 -- n is HTML-escaped
}

// hexCell renders any 0x hex string (address, hash, seed) shortened with a
// copy-able hover card but without links.
func hexCell(v any) template.HTML {
	s := strings.TrimSpace(fmt.Sprint(v))
	if len(s) <= 14 || !strings.HasPrefix(s, "0x") {
		return template.HTML(html.EscapeString(s)) // #nosec G203 -- escaped
	}
	esc := html.EscapeString(s)
	out := fmt.Sprintf(`<span class="hx mono">%s<span class="hx-card"><span class="hx-row mono">%s%s</span></span></span>`,
		hexLabel(s), esc, copyBtn(s))
	return template.HTML(out) // #nosec G203 -- escaped above
}

// jsonPretty renders v as indented JSON (for record-detail cards).
func jsonPretty(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprint(v)
	}
	return string(b)
}

// recordCard renders a "⋯" button whose card shows the full JSON of one
// record, so a table can show a few columns without losing any field.
func recordCard(v any) template.HTML {
	out := fmt.Sprintf(`<span class="hx hx-rec"><span class="hx-dots" title="all fields">⋯</span><span class="hx-card"><pre>%s</pre></span></span>`,
		html.EscapeString(jsonPretty(v)))
	return template.HTML(out) // #nosec G203 -- JSON is HTML-escaped
}

// durationStr formats a number of seconds as "1d 02h 03m 04s".
func durationStr(v any) string {
	secs := int64(asFloat(v))
	neg := secs < 0
	if neg {
		secs = -secs
	}
	d, h, m, s := secs/86400, secs%86400/3600, secs%3600/60, secs%60
	var out string
	switch {
	case d > 0:
		out = fmt.Sprintf("%dd %02dh %02dm %02ds", d, h, m, s)
	case h > 0:
		out = fmt.Sprintf("%dh %02dm %02ds", h, m, s)
	case m > 0:
		out = fmt.Sprintf("%dm %02ds", m, s)
	default:
		out = fmt.Sprintf("%ds", s)
	}
	if neg {
		return "-" + out
	}
	return out
}

// hasKey reports whether m contains key k (templates use it to guard
// optional response fields under missingkey=error).
func hasKey(m map[string]any, k string) bool {
	_, ok := m[k]
	return ok
}

// pct formats a ratio 0..1 as a percentage with one decimal.
func pct(v any) string {
	return fmt.Sprintf("%.1f%%", asFloat(v)*100)
}
