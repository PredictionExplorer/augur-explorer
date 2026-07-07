package cosmicgame

import "testing"

// roiLeaderboardWhitelist is the closed set of clauses Get_roi_leaderboard may
// ever interpolate into its ORDER BY.
var roiLeaderboardWhitelist = map[string]bool{
	"roi DESC": true,
	"win_rate DESC, rounds_participated DESC": true,
	"b.total_eth_spent DESC":                  true,
	"nft_prizes_count DESC":                   true,
	"b.num_bids DESC":                         true,
	"net_pl_eth DESC":                         true,
}

func TestRoiLeaderboardOrderClause(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"roi", "roi DESC"},
		{"winrate", "win_rate DESC, rounds_participated DESC"},
		{"spent", "b.total_eth_spent DESC"},
		{"nfts", "nft_prizes_count DESC"},
		{"bids", "b.num_bids DESC"},
		{"", "net_pl_eth DESC"},
		{"pl", "net_pl_eth DESC"},
		{"ROI", "net_pl_eth DESC"}, // matching is case-sensitive
		{"roi; DROP TABLE cg_bid;--", "net_pl_eth DESC"},
	}
	for _, tc := range cases {
		if got := roiLeaderboardOrderClause(tc.in); got != tc.want {
			t.Errorf("roiLeaderboardOrderClause(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func FuzzRoiLeaderboardOrderClause(f *testing.F) {
	for _, seed := range []string{"roi", "winrate", "spent", "nfts", "bids", "", "1; SELECT pg_sleep(10)--"} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, sortBy string) {
		got := roiLeaderboardOrderClause(sortBy)
		if !roiLeaderboardWhitelist[got] {
			t.Fatalf("roiLeaderboardOrderClause(%q) = %q: not in the whitelist", sortBy, got)
		}
	})
}
