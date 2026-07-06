package main

import (
	"sort"

	"github.com/spf13/cobra"

	rwprim "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// rankValue converts a position in a sorted leaderboard into the percentile
// rank value stored in the database (1.0 is the top).
func rankValue(position, total int) float64 {
	return (float64(position)/float64(total))*100.0 + 1.0
}

// updateProfitRanks writes the profit leaderboard ranks for all users.
func updateProfitRanks(storagew *rwstore.SQLStorageWrapper, records []rwprim.RankStats) {
	for i := range records {
		record := &records[i]
		storagew.Update_randomwalk_top_profit_rank(record.Aid, rankValue(i, len(records)), record.ProfitLoss)
	}
}

// updateTradeRanks writes the total-trades leaderboard ranks for all users.
func updateTradeRanks(storagew *rwstore.SQLStorageWrapper, records []rwprim.RankStats) {
	for i := range records {
		record := &records[i]
		storagew.Update_randomwalk_top_total_trades_rank(record.Aid, rankValue(i, len(records)), record.TotalTrades)
	}
}

// updateVolumeRanks writes the traded-volume leaderboard ranks for all users.
func updateVolumeRanks(storagew *rwstore.SQLStorageWrapper, records []rwprim.RankStats) {
	for i := range records {
		record := &records[i]
		storagew.Update_randomwalk_top_volume_rank(record.Aid, rankValue(i, len(records)), record.VolumeTraded)
	}
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
			storagew, err := connectRWStorage(newInfoLogger())
			if err != nil {
				return err
			}
			records := storagew.Get_randomwalk_ranking_data_for_all_users()

			sort.SliceStable(records, func(i, j int) bool {
				return records[j].TotalTrades < records[i].TotalTrades
			})
			updateTradeRanks(storagew, records)

			sort.SliceStable(records, func(i, j int) bool {
				return records[j].ProfitLoss < records[i].ProfitLoss
			})
			updateProfitRanks(storagew, records)

			sort.SliceStable(records, func(i, j int) bool {
				return records[j].VolumeTraded < records[i].VolumeTraded
			})
			updateVolumeRanks(storagew, records)
			return nil
		},
	}
}

func init() { register(newTopRatedCmd()) }
