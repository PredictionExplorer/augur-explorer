package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/dbverify"
)

type dbVerifyDeps struct {
	openDB    func(context.Context, string) (opsDB, error)
	loadIDs   func(context.Context, dbverify.Querier) ([]int64, error)
	newLoader func(dbverify.Querier) dbverify.Loader
	verify    func(context.Context, dbverify.Loader, dbverify.Loader, []int64, int) (dbverify.VerifyReport, error)
}

func defaultDBVerifyDeps() dbVerifyDeps {
	return dbVerifyDeps{
		openDB:  openOpsDB(dbCompareMaxConns),
		loadIDs: dbverify.LoadRandomWalkContractAddressIDs,
		newLoader: func(db dbverify.Querier) dbverify.Loader {
			return &dbverify.SQLLoader{DB: db}
		},
		verify: dbverify.VerifyDatabases,
	}
}

// newDBVerifyCmd builds `opsctl db verify`, the replacement for the
// standalone db_verify tool.
func newDBVerifyCmd() *cobra.Command {
	return newDBVerifyCmdWithDeps(defaultDBVerifyDeps())
}

func newDBVerifyCmdWithDeps(deps dbVerifyDeps) *cobra.Command {
	var (
		primaryConn   string
		secondaryConn string
	)
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Compare evt_log / transaction / block between two databases (randomwalk contracts)",
		Long: `Loads the RandomWalk-contract event logs, transactions and blocks from the
primary (gold standard) database and checks that the secondary database holds
exactly the same records.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.New(cmd.ErrOrStderr(), "", log.LstdFlags)
			primaryDB, err := deps.openDB(cmd.Context(), primaryConn)
			if err != nil {
				return fmt.Errorf("connect to primary: %w", err)
			}
			defer primaryDB.Close()
			logger.Println("Connected to primary database")

			secondaryDB, err := deps.openDB(cmd.Context(), secondaryConn)
			if err != nil {
				return fmt.Errorf("connect to secondary: %w", err)
			}
			defer secondaryDB.Close()
			logger.Println("Connected to secondary database")

			contractAddressIDs, err := deps.loadIDs(cmd.Context(), primaryDB)
			if err != nil {
				return err
			}
			logger.Printf("Found %d contract address IDs: %v", len(contractAddressIDs), contractAddressIDs)

			report, err := deps.verify(
				cmd.Context(),
				deps.newLoader(primaryDB),
				deps.newLoader(secondaryDB),
				contractAddressIDs,
				dbverify.DefaultVerifyReportLimit,
			)
			if err != nil {
				return err
			}
			for _, line := range dbverify.FormatVerifyReport(report) {
				logger.Print(line)
			}
			if !report.Matched() {
				return errors.New("verification FAILED")
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&primaryConn, "primary", "", "Primary DB connection string (production - gold standard)")
	cmd.Flags().StringVar(&secondaryConn, "secondary", "", "Secondary DB connection string (new rwcg)")
	_ = cmd.MarkFlagRequired("primary")
	_ = cmd.MarkFlagRequired("secondary")
	return cmd
}
