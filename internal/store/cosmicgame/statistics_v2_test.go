package cosmicgame

import (
	"strings"
	"testing"
)

func TestROILeaderboardV2OrderWhitelist(t *testing.T) {
	t.Parallel()
	for _, sort := range []ROILeaderboardSort{
		ROILeaderboardNetProfit,
		ROILeaderboardROI,
		ROILeaderboardWinRate,
		ROILeaderboardSpent,
		ROILeaderboardNFTs,
		ROILeaderboardBids,
	} {
		order, ok := roiLeaderboardV2Order(sort)
		if !ok || !strings.Contains(order, "bidder_aid ASC") ||
			roiLeaderboardV2Column(sort) == "" {
			t.Errorf("sort %q: order=%q column=%q ok=%v",
				sort, order, roiLeaderboardV2Column(sort), ok)
		}
	}
	if order, ok := roiLeaderboardV2Order("other"); ok || order != "" {
		t.Fatalf("unknown sort accepted: %q, %v", order, ok)
	}
}

func TestStatisticsV2PageMethodsRejectInvalidArguments(t *testing.T) {
	t.Parallel()
	var repo Repo
	if _, _, err := repo.ROILeaderboardPage(t.Context(), -1, ROILeaderboardNetProfit, nil, 1); err == nil {
		t.Error("ROI accepted negative minBids")
	}
	if _, _, err := repo.ROILeaderboardPage(t.Context(), 0, "other", nil, 1); err == nil {
		t.Error("ROI accepted invalid sort")
	}
	for _, limit := range []int{0, -1, maxStatisticsPageLimit + 1} {
		if _, _, err := repo.ROILeaderboardPage(t.Context(), 0, ROILeaderboardNetProfit, nil, limit); err == nil {
			t.Errorf("ROI accepted limit %d", limit)
		}
		if _, _, err := repo.ClaimsSummaryPage(t.Context(), nil, limit); err == nil {
			t.Errorf("claims accepted limit %d", limit)
		}
		if _, _, err := repo.ClaimTransactionsPage(t.Context(), 0, nil, limit); err == nil {
			t.Errorf("transactions accepted limit %d", limit)
		}
		if _, _, err := repo.AttachedTokensPage(t.Context(), 0, nil, limit); err == nil {
			t.Errorf("attached accepted limit %d", limit)
		}
		if _, _, err := repo.UnclaimedItemsPage(t.Context(), 0, nil, limit); err == nil {
			t.Errorf("unclaimed accepted limit %d", limit)
		}
	}
	if _, _, err := repo.ROILeaderboardPage(t.Context(), 0, ROILeaderboardNetProfit, &ROILeaderboardPageCursor{
		Sort: ROILeaderboardROI, MinBids: 0, SortValue: "1", BidderAid: 1,
	}, 1); err == nil {
		t.Error("ROI accepted cross-sort cursor")
	}
	if _, _, err := repo.ClaimsSummaryPage(t.Context(), &ClaimSummaryCursor{}, 1); err == nil {
		t.Error("claims accepted zero cursor")
	}
	if _, err := repo.ClaimSummaryByRound(t.Context(), -1); err == nil {
		t.Error("summary accepted negative round")
	}
	if _, _, err := repo.ClaimTransactionsPage(t.Context(), -1, nil, 1); err == nil {
		t.Error("transactions accepted negative round")
	}
	if _, _, err := repo.AttachedTokensPage(t.Context(), 0, &ClaimEventCursor{}, 1); err == nil {
		t.Error("attached accepted zero cursor")
	}
	if _, _, err := repo.UnclaimedItemsPage(t.Context(), 0, &UnclaimedItemCursor{Segment: 3, Key: 1}, 1); err == nil {
		t.Error("unclaimed accepted invalid segment")
	}
}

func TestClaimUnionQueriesApplyCursorToEveryBranch(t *testing.T) {
	t.Parallel()
	transactions := claimTransactionsPageSQL(true)
	if got := strings.Count(transactions, "evtlog_id>$2"); got != 3 {
		t.Fatalf("transaction cursor predicate count = %d, want 3", got)
	}
	attached := attachedTokensPageSQL(true)
	if got := strings.Count(attached, "evtlog_id>$2"); got != 2 {
		t.Fatalf("attached cursor predicate count = %d, want 2", got)
	}
}
