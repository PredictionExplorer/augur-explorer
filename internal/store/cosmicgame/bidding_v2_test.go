package cosmicgame

import (
	"strings"
	"testing"

	. "github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestBuildBidSelectQueryIncludesV2Columns(t *testing.T) {
	sw := &SQLStorageWrapper{S: &SQLStorage{}}
	query := sw.buildBidSelectQuery("", "b.id DESC", "LIMIT 1")
	for _, col := range []string{
		"bid_cst_reward_amount",
		"cst_dutch_auction_duration",
		"bid_cst_reward_amount >= 0",
		"cst_dutch_auction_duration >= 0",
	} {
		if !strings.Contains(query, col) {
			t.Fatalf("buildBidSelectQuery missing %q in:\n%s", col, query)
		}
	}
}
