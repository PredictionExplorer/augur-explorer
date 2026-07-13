package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/dbverify"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

type dbEvtlogDiffDeps struct {
	openDB    func(string, string) (*sql.DB, error)
	loadIDs   func(context.Context, *sql.DB) ([]int64, error)
	newLoader func(*sql.DB) dbverify.Loader
	diff      func(context.Context, dbverify.Loader, dbverify.Loader, []int64, int, int) (dbverify.EventLogDiffReport, error)
}

func defaultDBEvtlogDiffDeps() dbEvtlogDiffDeps {
	return dbEvtlogDiffDeps{
		openDB:  sql.Open,
		loadIDs: dbverify.LoadRandomWalkContractAddressIDs,
		newLoader: func(db *sql.DB) dbverify.Loader {
			return &dbverify.SQLLoader{DB: db}
		},
		diff: dbverify.DiffEventLogs,
	}
}

// newDBEvtlogDiffCmd builds `opsctl db evtlog-diff`, the replacement for the
// standalone evtlog_diff tool.
func newDBEvtlogDiffCmd() *cobra.Command {
	return newDBEvtlogDiffCmdWithDeps(defaultDBEvtlogDiffDeps())
}

func newDBEvtlogDiffCmdWithDeps(deps dbEvtlogDiffDeps) *cobra.Command {
	var (
		primaryConn   string
		secondaryConn string
		limit         int
	)
	cmd := &cobra.Command{
		Use:   "evtlog-diff",
		Short: "Field-level diff of evt_log between two databases (randomwalk contracts)",
		Long: `Loads the RandomWalk-contract event logs from both databases, indexes them
by log_rlp content, and reports records missing from or extra in the
secondary plus any field mismatches on matching records.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.New(cmd.ErrOrStderr(), "", log.LstdFlags)
			return runEvtlogDiffWithDeps(cmd.Context(), logger, primaryConn, secondaryConn, limit, deps)
		},
	}
	cmd.Flags().StringVar(&primaryConn, "primary", "", "Primary DB connection string (gold standard)")
	cmd.Flags().StringVar(&secondaryConn, "secondary", "", "Secondary DB connection string (to verify)")
	cmd.Flags().IntVar(&limit, "limit", 0, "Limit comparison to first N records (0 = all)")
	_ = cmd.MarkFlagRequired("primary")
	_ = cmd.MarkFlagRequired("secondary")
	return cmd
}

func runEvtlogDiff(ctx context.Context, logger *log.Logger, primaryConn, secondaryConn string, limit int) error {
	return runEvtlogDiffWithDeps(ctx, logger, primaryConn, secondaryConn, limit, defaultDBEvtlogDiffDeps())
}

func runEvtlogDiffWithDeps(
	ctx context.Context,
	logger *log.Logger,
	primaryConn, secondaryConn string,
	limit int,
	deps dbEvtlogDiffDeps,
) error {
	primaryDB, err := deps.openDB("postgres", primaryConn)
	if err != nil {
		return fmt.Errorf("connect to primary: %w", err)
	}
	defer func() { _ = primaryDB.Close() }()
	logger.Println("Connected to primary database")

	secondaryDB, err := deps.openDB("postgres", secondaryConn)
	if err != nil {
		return fmt.Errorf("connect to secondary: %w", err)
	}
	defer func() { _ = secondaryDB.Close() }()
	logger.Println("Connected to secondary database")

	contractAddressIDs, err := deps.loadIDs(ctx, primaryDB)
	if err != nil {
		return err
	}
	logger.Printf("Found %d contract address IDs: %v", len(contractAddressIDs), contractAddressIDs)

	report, err := deps.diff(
		ctx,
		deps.newLoader(primaryDB),
		deps.newLoader(secondaryDB),
		contractAddressIDs,
		limit,
		dbverify.DefaultDiffReportLimit,
	)
	if err != nil {
		return err
	}
	for _, line := range dbverify.FormatEventLogDiffReport(report) {
		logger.Print(line)
	}
	return nil
}
