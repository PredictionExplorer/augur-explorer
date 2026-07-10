//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestGetCosmicGameStatistics(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_game_statistics", func() any {
		stats, err := r.CosmicGameStatistics(context.Background())
		if err != nil {
			t.Fatalf("CosmicGameStatistics: %v", err)
		}
		return stats
	})
}

func TestCosmicGameGlobalStatistics(t *testing.T) {
	r := repo(t)
	golden(t, "global_statistics_v2", func() any {
		stats, err := r.CosmicGameGlobalStatistics(context.Background())
		if err != nil {
			t.Fatalf("CosmicGameGlobalStatistics: %v", err)
		}
		return stats
	})
}

func TestGetStakeStatisticsCst(t *testing.T) {
	r := repo(t)
	golden(t, "stake_statistics_cst", func() any {
		stats, err := r.StakeStatisticsCst(context.Background())
		if err != nil {
			t.Fatalf("StakeStatisticsCst: %v", err)
		}
		return stats
	})
}

func TestGetStakeStatisticsRwalk(t *testing.T) {
	r := repo(t)
	golden(t, "stake_statistics_rwalk", func() any {
		stats, err := r.StakeStatisticsRwalk(context.Background())
		if err != nil {
			t.Fatalf("StakeStatisticsRwalk: %v", err)
		}
		return stats
	})
}

func TestGetCosmicGameRoundStatistics(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_game_round_statistics_0", func() any {
		stats, err := r.CosmicGameRoundStatistics(context.Background(), 0)
		if err != nil {
			t.Fatalf("CosmicGameRoundStatistics(0): %v", err)
		}
		return stats
	})
	golden(t, "cosmic_game_round_statistics_3_open", func() any {
		stats, err := r.CosmicGameRoundStatistics(context.Background(), 3)
		if err != nil {
			t.Fatalf("CosmicGameRoundStatistics(3): %v", err)
		}
		return stats
	})
}

func TestGetUniqueBidders(t *testing.T) {
	r := repo(t)
	golden(t, "unique_bidders", func() any {
		recs, err := r.UniqueBidders(context.Background())
		if err != nil {
			t.Fatalf("UniqueBidders: %v", err)
		}
		return recs
	})
}

func TestGetUniqueWinners(t *testing.T) {
	r := repo(t)
	golden(t, "unique_winners", func() any {
		recs, err := r.UniqueWinners(context.Background())
		if err != nil {
			t.Fatalf("UniqueWinners: %v", err)
		}
		return recs
	})
}

func TestGetRoiLeaderboard(t *testing.T) {
	r := repo(t)
	// Every whitelisted sort column plus the default branch.
	for _, sortBy := range []string{"roi", "winrate", "spent", "nfts", "bids", "default"} {
		sortArg := sortBy
		if sortArg == "default" {
			sortArg = ""
		}
		golden(t, "roi_leaderboard_"+sortBy, func() any {
			recs, err := r.RoiLeaderboard(context.Background(), 0, sortArg, 0, 100)
			if err != nil {
				t.Fatalf("RoiLeaderboard(%q): %v", sortArg, err)
			}
			return recs
		})
	}
	golden(t, "roi_leaderboard_min_bids_3", func() any {
		recs, err := r.RoiLeaderboard(context.Background(), 3, "roi", 0, 100)
		if err != nil {
			t.Fatalf("RoiLeaderboard(min_bids=3): %v", err)
		}
		return recs
	})
}

func TestROILeaderboardPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	tests := []struct {
		sort   ROILeaderboardSort
		legacy string
	}{
		{ROILeaderboardNetProfit, ""},
		{ROILeaderboardROI, "roi"},
		{ROILeaderboardWinRate, "winrate"},
		{ROILeaderboardSpent, "spent"},
		{ROILeaderboardNFTs, "nfts"},
		{ROILeaderboardBids, "bids"},
	}
	for _, tc := range tests {
		t.Run(string(tc.sort), func(t *testing.T) {
			legacy, err := r.RoiLeaderboard(ctx, 0, tc.legacy, 0, 100)
			if err != nil {
				t.Fatalf("legacy ROI: %v", err)
			}
			var (
				after *ROILeaderboardPageCursor
				paged []ROILeaderboardRecord
			)
			for {
				page, hasMore, err := r.ROILeaderboardPage(ctx, 0, tc.sort, after, 1)
				if err != nil {
					t.Fatalf("ROI page after %+v: %v", after, err)
				}
				if page == nil {
					t.Fatal("ROI page encoded as nil")
				}
				paged = append(paged, page...)
				if !hasMore {
					break
				}
				last := page[len(page)-1]
				after = &ROILeaderboardPageCursor{
					Sort:      tc.sort,
					MinBids:   0,
					SortValue: ROILeaderboardSortValue(last, tc.sort),
					Secondary: func() int64 {
						if tc.sort == ROILeaderboardWinRate {
							return last.RoundsParticipated
						}
						return 0
					}(),
					BidderAid: last.BidderAid,
				}
			}
			if len(legacy) != len(paged) {
				t.Fatalf("legacy/page lengths = %d/%d", len(legacy), len(paged))
			}
			for i := range legacy {
				if legacy[i].BidderAid != paged[i].BidderAid ||
					legacy[i].TotalEthSpent != paged[i].TotalEthSpentWei ||
					legacy[i].EthWon != paged[i].EthWonWei {
					t.Fatalf("legacy/page row %d differs: %+v / %+v", i, legacy[i], paged[i])
				}
			}
		})
	}
	golden(t, "roi_leaderboard_page_v2", func() any {
		page, _, err := r.ROILeaderboardPage(ctx, 0, ROILeaderboardNetProfit, nil, 100)
		if err != nil {
			t.Fatalf("ROILeaderboardPage golden: %v", err)
		}
		return page
	})
}

func TestGetClaimsByRound(t *testing.T) {
	r := repo(t)
	golden(t, "claims_by_round", func() any {
		recs, err := r.ClaimsByRound(context.Background())
		if err != nil {
			t.Fatalf("ClaimsByRound: %v", err)
		}
		return recs
	})
}

func TestClaimsSummaryPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	legacy, err := r.ClaimsByRound(ctx)
	if err != nil {
		t.Fatalf("ClaimsByRound: %v", err)
	}
	var (
		after *ClaimSummaryCursor
		paged []ClaimSummaryRecord
	)
	for {
		page, hasMore, err := r.ClaimsSummaryPage(ctx, after, 1)
		if err != nil {
			t.Fatalf("claims page after %+v: %v", after, err)
		}
		if page == nil {
			t.Fatal("claims page encoded as nil")
		}
		paged = append(paged, page...)
		if !hasMore {
			break
		}
		last := page[len(page)-1]
		after = &ClaimSummaryCursor{RoundNum: last.RoundNum, EventLogID: last.EventLogID}
	}
	if len(legacy) != len(paged) {
		t.Fatalf("legacy/page lengths = %d/%d", len(legacy), len(paged))
	}
	for i := range legacy {
		if legacy[i].RoundNum != paged[i].RoundNum ||
			legacy[i].TotalAwarded != paged[i].TotalAwarded ||
			legacy[i].TotalUnclaimed != paged[i].TotalUnclaimed {
			t.Fatalf("legacy/page row %d differs: %+v / %+v", i, legacy[i], paged[i])
		}
	}
	for i := range paged {
		if paged[i].RoundNum == 0 && paged[i].UnclaimedEthAmountWei != "130000000000000000" {
			t.Fatalf("round-0 exact unclaimed ETH = %q", paged[i].UnclaimedEthAmountWei)
		}
	}
	golden(t, "claims_summary_page_v2", func() any {
		page, _, err := r.ClaimsSummaryPage(ctx, nil, 100)
		if err != nil {
			t.Fatalf("ClaimsSummaryPage golden: %v", err)
		}
		return page
	})
}

func TestRoundClaimDetailPages(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	legacy, err := r.ClaimDetailByRound(ctx, 0)
	if err != nil {
		t.Fatalf("ClaimDetailByRound(0): %v", err)
	}
	summary, err := r.ClaimSummaryByRound(ctx, 0)
	if err != nil || summary.RoundNum != 0 {
		t.Fatalf("ClaimSummaryByRound(0) = %+v, %v", summary, err)
	}
	golden(t, "claim_transactions_page_v2", func() any {
		page, _, err := r.ClaimTransactionsPage(ctx, 0, nil, 100)
		if err != nil {
			t.Fatalf("ClaimTransactionsPage: %v", err)
		}
		return page
	})
	golden(t, "attached_tokens_page_v2", func() any {
		page, _, err := r.AttachedTokensPage(ctx, 0, nil, 100)
		if err != nil {
			t.Fatalf("AttachedTokensPage: %v", err)
		}
		return page
	})
	golden(t, "unclaimed_items_page_v2", func() any {
		page, _, err := r.UnclaimedItemsPage(ctx, 0, nil, 100)
		if err != nil {
			t.Fatalf("UnclaimedItemsPage: %v", err)
		}
		return page
	})
	allTransactions, _, err := r.ClaimTransactionsPage(ctx, 0, nil, 100)
	if err != nil || len(allTransactions) != len(legacy.ClaimTransactions) {
		t.Fatalf("transaction parity lengths = %d/%d, err=%v",
			len(allTransactions), len(legacy.ClaimTransactions), err)
	}
	allAttached, _, err := r.AttachedTokensPage(ctx, 0, nil, 100)
	if err != nil || len(allAttached) != len(legacy.AttachedTokens) {
		t.Fatalf("attached parity lengths = %d/%d, err=%v",
			len(allAttached), len(legacy.AttachedTokens), err)
	}

	for name, fetch := range map[string]func(*ClaimEventCursor) (int64, bool, error){
		"transactions": func(after *ClaimEventCursor) (int64, bool, error) {
			page, more, err := r.ClaimTransactionsPage(ctx, 0, after, 1)
			if len(page) == 0 {
				return 0, more, err
			}
			return page[0].EventLogID, more, err
		},
		"attached": func(after *ClaimEventCursor) (int64, bool, error) {
			page, more, err := r.AttachedTokensPage(ctx, 0, after, 1)
			if len(page) == 0 {
				return 0, more, err
			}
			return page[0].EventLogID, more, err
		},
	} {
		t.Run(name, func(t *testing.T) {
			var (
				after *ClaimEventCursor
				ids   []int64
			)
			for {
				id, more, err := fetch(after)
				if err != nil {
					t.Fatalf("page: %v", err)
				}
				if id > 0 {
					ids = append(ids, id)
				}
				if !more {
					break
				}
				after = &ClaimEventCursor{EventLogID: id}
			}
			for i := 1; i < len(ids); i++ {
				if ids[i] <= ids[i-1] {
					t.Fatalf("event ids are not ascending: %v", ids)
				}
			}
		})
	}

	allUnclaimed, _, err := r.UnclaimedItemsPage(ctx, 0, nil, 100)
	if err != nil {
		t.Fatalf("all unclaimed items: %v", err)
	}
	var (
		unclaimedAfter *UnclaimedItemCursor
		unclaimedPaged []UnclaimedItemRecord
	)
	for {
		page, more, err := r.UnclaimedItemsPage(ctx, 0, unclaimedAfter, 1)
		if err != nil {
			t.Fatalf("unclaimed page: %v", err)
		}
		unclaimedPaged = append(unclaimedPaged, page...)
		if !more {
			break
		}
		last := page[len(page)-1]
		unclaimedAfter = &UnclaimedItemCursor{Segment: last.Segment, Key: last.Key}
	}
	if !reflect.DeepEqual(allUnclaimed, unclaimedPaged) {
		t.Fatalf("unclaimed full/page differ:\nfull=%#v\npage=%#v", allUnclaimed, unclaimedPaged)
	}

	emptyTransactions, more, err := r.ClaimTransactionsPage(ctx, 1, nil, 10)
	if err != nil || more || emptyTransactions == nil {
		t.Fatalf("round-1 transactions = %#v, more=%v, err=%v", emptyTransactions, more, err)
	}
}

func TestStatisticsV2PagesPropagateCancellation(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	checks := map[string]func() error{
		"global": func() error {
			_, err := r.CosmicGameGlobalStatistics(ctx)
			return err
		},
		"roi": func() error {
			_, _, err := r.ROILeaderboardPage(ctx, 0, ROILeaderboardNetProfit, nil, 1)
			return err
		},
		"claims": func() error {
			_, _, err := r.ClaimsSummaryPage(ctx, nil, 1)
			return err
		},
		"summary": func() error {
			_, err := r.ClaimSummaryByRound(ctx, 0)
			return err
		},
		"transactions": func() error {
			_, _, err := r.ClaimTransactionsPage(ctx, 0, nil, 1)
			return err
		},
		"attached": func() error {
			_, _, err := r.AttachedTokensPage(ctx, 0, nil, 1)
			return err
		},
		"unclaimed": func() error {
			_, _, err := r.UnclaimedItemsPage(ctx, 0, nil, 1)
			return err
		},
	}
	for name, check := range checks {
		t.Run(name, func(t *testing.T) {
			if err := check(); !errors.Is(err, context.Canceled) {
				t.Fatalf("error = %v, want context.Canceled", err)
			}
		})
	}
}

func TestGetClaimDetailByRound(t *testing.T) {
	r := repo(t)
	golden(t, "claim_detail_by_round_0", func() any {
		detail, err := r.ClaimDetailByRound(context.Background(), 0)
		if err != nil {
			t.Fatalf("ClaimDetailByRound(0): %v", err)
		}
		return detail
	})
	golden(t, "claim_detail_by_round_1", func() any {
		detail, err := r.ClaimDetailByRound(context.Background(), 1)
		if err != nil {
			t.Fatalf("ClaimDetailByRound(1): %v", err)
		}
		return detail
	})
}

func TestGetUniqueStakersCst(t *testing.T) {
	r := repo(t)
	golden(t, "unique_stakers_cst", func() any {
		recs, err := r.UniqueStakersCst(context.Background())
		if err != nil {
			t.Fatalf("UniqueStakersCst: %v", err)
		}
		return recs
	})
}

func TestGetUniqueStakersRwalk(t *testing.T) {
	r := repo(t)
	golden(t, "unique_stakers_rwalk", func() any {
		recs, err := r.UniqueStakersRwalk(context.Background())
		if err != nil {
			t.Fatalf("UniqueStakersRwalk: %v", err)
		}
		return recs
	})
}

func TestGetUniqueStakersBoth(t *testing.T) {
	r := repo(t)
	golden(t, "unique_stakers_both", func() any {
		recs, err := r.UniqueStakersBoth(context.Background())
		if err != nil {
			t.Fatalf("UniqueStakersBoth: %v", err)
		}
		return recs
	})
}

func TestGetUniqueDonors(t *testing.T) {
	r := repo(t)
	golden(t, "unique_donors", func() any {
		recs, err := r.UniqueDonors(context.Background())
		if err != nil {
			t.Fatalf("UniqueDonors: %v", err)
		}
		return recs
	})
}

func TestGetNFTDonationStats(t *testing.T) {
	r := repo(t)
	golden(t, "nft_donation_stats", func() any {
		recs, err := r.NFTDonationStats(context.Background())
		if err != nil {
			t.Fatalf("NFTDonationStats: %v", err)
		}
		return recs
	})
}

func TestGetRecordCounters(t *testing.T) {
	r := repo(t)
	golden(t, "record_counters", func() any {
		counters, err := r.RecordCounters(context.Background())
		if err != nil {
			t.Fatalf("RecordCounters: %v", err)
		}
		return counters
	})
}
