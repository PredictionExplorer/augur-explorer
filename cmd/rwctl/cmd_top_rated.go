package main

import (
	"cmp"
	"context"
	"fmt"
	"slices"

	"github.com/spf13/cobra"

	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// rankValue converts a position in a sorted leaderboard into the percentile
// rank value stored in the database (1.0 is the top).
func rankValue(position, total int) float64 {
	return (float64(position)/float64(total))*100.0 + 1.0
}

// updateProfitRanks writes the profit leaderboard ranks for all users.
func updateProfitRanks(ctx context.Context, repo *rwstore.Repo, records []rwmodel.RankStats) error {
	for i := range records {
		record := &records[i]
		if err := repo.UpdateTopProfitRank(ctx, record.Aid, rankValue(i, len(records)), record.ProfitLoss); err != nil {
			return fmt.Errorf("profit rank for aid %d: %w", record.Aid, err)
		}
	}
	return nil
}

// updateTradeRanks writes the total-trades leaderboard ranks for all users.
func updateTradeRanks(ctx context.Context, repo *rwstore.Repo, records []rwmodel.RankStats) error {
	for i := range records {
		record := &records[i]
		if err := repo.UpdateTopTotalTradesRank(ctx, record.Aid, rankValue(i, len(records)), record.TotalTrades); err != nil {
			return fmt.Errorf("trade rank for aid %d: %w", record.Aid, err)
		}
	}
	return nil
}

// updateVolumeRanks writes the traded-volume leaderboard ranks for all users.
func updateVolumeRanks(ctx context.Context, repo *rwstore.Repo, records []rwmodel.RankStats) error {
	for i := range records {
		record := &records[i]
		if err := repo.UpdateTopVolumeRank(ctx, record.Aid, rankValue(i, len(records)), record.VolumeTraded); err != nil {
			return fmt.Errorf("volume rank for aid %d: %w", record.Aid, err)
		}
	}
	return nil
}

// newTopRatedCmd builds the top-rated subcommand (legacy rw_toprated tool).
//
// Design notes carried over from the original tool: statistics are not a
// critical task and run once a day, so production is not loaded with a heavy
// update query. All updates go through the primary key and the sorting is
// done in the client's memory. The process is single threaded and
// light-weight but takes a long time to finish.
func newTopRatedCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "top-rated",
		Short: "Recompute top-100 trader/profit/volume rankings",
		Long: "Generates top-rated-100 statistics (top100 traders, top100 profit makers, top100 volume makers).\n\n" +
			"Environment variables:\n  PGSQL_*  PostgreSQL connection (required)",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			repo, _, err := connectRWStorage(newInfoLogger())
			if err != nil {
				return err
			}
			records, err := repo.RankingDataForAllUsers(ctx)
			if err != nil {
				return err
			}

			slices.SortStableFunc(records, func(a, b rwmodel.RankStats) int {
				return cmp.Compare(b.TotalTrades, a.TotalTrades)
			})
			if err := updateTradeRanks(ctx, repo, records); err != nil {
				return err
			}

			slices.SortStableFunc(records, func(a, b rwmodel.RankStats) int {
				return cmp.Compare(b.ProfitLoss, a.ProfitLoss)
			})
			if err := updateProfitRanks(ctx, repo, records); err != nil {
				return err
			}

			slices.SortStableFunc(records, func(a, b rwmodel.RankStats) int {
				return cmp.Compare(b.VolumeTraded, a.VolumeTraded)
			})
			return updateVolumeRanks(ctx, repo, records)
		},
	}
}

func init() { register(newTopRatedCmd()) }
