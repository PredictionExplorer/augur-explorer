package cosmicgame

import (
	"strings"
	"testing"
)

// TestBidSelectQueryWhitelists proves the bid query builder composes every
// whitelisted clause combination into well-formed SQL and rejects anything
// outside the whitelists — the guard that keeps request-derived strings from
// ever reaching ORDER BY (the classic injection site of dynamic builders).
func TestBidSelectQueryWhitelists(t *testing.T) {
	for where := range bidWhereWhitelist {
		for orderBy := range bidOrderWhitelist {
			for paging := range bidPagingWhitelist {
				query, err := bidSelectQuery(where, orderBy, paging)
				if err != nil {
					t.Fatalf("bidSelectQuery(%q, %q, %q): %v", where, orderBy, paging, err)
				}
				if !strings.HasPrefix(query, "SELECT b.evtlog_id") {
					t.Fatalf("query does not start with the unified SELECT:\n%s", query)
				}
				if where != "" && !strings.Contains(query, "WHERE "+where) {
					t.Errorf("missing WHERE %q in query", where)
				}
				if orderBy != "" && !strings.Contains(query, "ORDER BY "+orderBy) {
					t.Errorf("missing ORDER BY %q in query", orderBy)
				}
				if paging != "" && !strings.HasSuffix(query, paging) {
					t.Errorf("missing paging %q at end of query", paging)
				}
			}
		}
	}

	rejected := []struct{ where, orderBy, paging string }{
		{"b.round_num=1; DROP TABLE cg_bid", "b.id DESC", ""},
		{"", "b.id DESC, (SELECT 1)", ""},
		{"", "b.id desc", ""}, // even case differences are rejected
		{"", "", "OFFSET 0 LIMIT 10"},
		{"b.evtlog_id=$1 OR 1=1", "", ""},
	}
	for _, tc := range rejected {
		if _, err := bidSelectQuery(tc.where, tc.orderBy, tc.paging); err == nil {
			t.Errorf("bidSelectQuery(%q, %q, %q) accepted a non-whitelisted clause",
				tc.where, tc.orderBy, tc.paging)
		}
	}
}

// TestBidSelectQueryIncludesV2Columns pins the IBiddingV2 columns in the
// pgx-native builder (the legacy twin is pinned by bidding_v2_test.go until
// user-specific.go converts).
func TestBidSelectQueryIncludesV2Columns(t *testing.T) {
	query, err := bidSelectQuery("", "b.id DESC", "")
	if err != nil {
		t.Fatalf("bidSelectQuery: %v", err)
	}
	for _, col := range []string{
		"bid_cst_reward_amount",
		"cst_dutch_auction_duration",
		"bid_cst_reward_amount >= 0",
		"cst_dutch_auction_duration >= 0",
	} {
		if !strings.Contains(query, col) {
			t.Fatalf("bidSelectQuery missing %q in:\n%s", col, query)
		}
	}
}

func TestBidsByRoundPageRejectsNonPositiveLimit(t *testing.T) {
	t.Parallel()

	var repo Repo
	for _, limit := range []int{0, -1} {
		if _, _, err := repo.BidsByRoundPage(t.Context(), 0, BidPageCursor{}, limit); err == nil {
			t.Errorf("BidsByRoundPage(limit=%d) succeeded", limit)
		}
	}
}
