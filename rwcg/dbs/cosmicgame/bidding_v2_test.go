package cosmicgame

import (
	"strings"
	"testing"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
)

func TestBuildBidSelectQueryIncludesV2Columns(t *testing.T) {
	sw := &SQLStorageWrapper{S: &SQLStorage{}}
	query := sw.buildBidSelectQuery("", "b.id DESC", "LIMIT 1")
	for _, col := range []string{
		// V3 bid CST reward 90/10 split (sourced from cg_bid_reward via a LATERAL join).
		"cg_bid_reward",
		"prev_reward",
		"this_reward",
		"cst_dutch_auction_duration",
		"cst_dutch_auction_duration >= 0",
	} {
		if !strings.Contains(query, col) {
			t.Fatalf("buildBidSelectQuery missing %q in:\n%s", col, query)
		}
	}
}
