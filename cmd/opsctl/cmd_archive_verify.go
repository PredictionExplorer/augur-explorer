package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/archive"
)

type archiveVerifyDeps struct {
	resolveProjects func(string) ([]string, error)
	openDB          func(context.Context, string) (opsDB, error)
	newVerifier     func(archive.Querier, archive.Logger) archive.ProjectVerifier
	verifyProjects  func(context.Context, []string, archive.VerifyOptions, archive.ProjectVerifier) (archive.VerificationReport, error)
}

func defaultArchiveVerifyDeps() archiveVerifyDeps {
	return archiveVerifyDeps{
		resolveProjects: archive.ResolveProjects,
		openDB:          openOpsDB(archiveMaxConns),
		newVerifier: func(db archive.Querier, logger archive.Logger) archive.ProjectVerifier {
			return &archive.SQLVerifier{DB: db, Logger: logger}
		},
		verifyProjects: archive.VerifyProjects,
	}
}

// newArchiveVerifyCmd builds `opsctl archive verify`, the replacement for the
// standalone arch_verify tool.
func newArchiveVerifyCmd() *cobra.Command {
	return newArchiveVerifyCmdWithDeps(defaultArchiveVerifyDeps())
}

func newArchiveVerifyCmdWithDeps(deps archiveVerifyDeps) *cobra.Command {
	var (
		dbConn          string
		projectType     string
		strictBlockMeta bool
		strictTxNumLogs bool
	)
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Check live evt_log / transaction / block against the arch_* tables",
		Long: `Runs archival consistency checks between the live tables and the arch_*
tables in the same database, per project. Exits non-zero when a blocking
mismatch is found; metadata drift is a warning unless a --strict-* flag is set.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			projects, err := deps.resolveProjects(projectType)
			if err != nil {
				return err
			}
			db, err := deps.openDB(cmd.Context(), dbConn)
			if err != nil {
				return fmt.Errorf("connect: %w", err)
			}
			defer db.Close()

			options := archive.VerifyOptions{
				StrictBlockMetadata: strictBlockMeta,
				StrictTxNumLogs:     strictTxNumLogs,
			}
			logger := log.New(cmd.ErrOrStderr(), "", log.LstdFlags)
			report, err := deps.verifyProjects(
				cmd.Context(),
				projects,
				options,
				deps.newVerifier(db, logger),
			)
			if err != nil {
				return err
			}

			logger.Println("")
			logger.Println("=== SUMMARY ===")
			if report.Passed {
				logger.Println("OK — no blocking mismatches for selected project(s). Review any warnings above.")
				return nil
			}
			logger.Println("FAILED — see details above.")
			return errors.New("archive verification failed")
		},
	}
	cmd.Flags().StringVar(&dbConn, "db", "", "PostgreSQL connection string (same DB holds live + arch_* tables)")
	cmd.Flags().StringVar(&projectType, "project", "", "Project: randomwalk | cosmicgame | both (same order as archive export: cosmicgame then randomwalk)")
	cmd.Flags().BoolVar(&strictBlockMeta, "strict-arch-block-metadata", false,
		"If set, require arch_block num_tx, ts, cash_flow to match live block (default: only block_hash and parent_hash must match).")
	cmd.Flags().BoolVar(&strictTxNumLogs, "strict-arch-tx-num-logs", false,
		"If set, require arch_tx.num_logs to match transaction.num_logs (default: ignore num_logs drift; indexer may refresh it after archival).")
	_ = cmd.MarkFlagRequired("db")
	_ = cmd.MarkFlagRequired("project")
	return cmd
}
