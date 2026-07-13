package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/archive"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

type archiveExportDeps struct {
	resolveProjects func(string) ([]string, error)
	openDB          func(string, string) (*sql.DB, error)
	newExporter     func(*sql.DB, *sql.DB, archive.Logger) archive.ProjectExporter
	exportProjects  func(context.Context, []string, archive.ProjectExporter) ([]archive.ExportResult, error)
}

func defaultArchiveExportDeps() archiveExportDeps {
	return archiveExportDeps{
		resolveProjects: archive.ResolveProjects,
		openDB:          sql.Open,
		newExporter: func(source, destination *sql.DB, logger archive.Logger) archive.ProjectExporter {
			return &archive.SQLExporter{
				Source:      source,
				Destination: destination,
				BatchSize:   archive.DefaultExportBatchSize,
				Logger:      logger,
			}
		},
		exportProjects: archive.ExportProjects,
	}
}

// newArchiveExportCmd builds `opsctl archive export`, the replacement for the
// standalone archive_export tool.
func newArchiveExportCmd() *cobra.Command {
	return newArchiveExportCmdWithDeps(defaultArchiveExportDeps())
}

func newArchiveExportCmdWithDeps(deps archiveExportDeps) *cobra.Command {
	var (
		srcConn     string
		dstConn     string
		projectType string
	)
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Copy evt_log / transaction / block rows for project contracts into arch_* tables",
		Long: `Copies evt_log rows for a project's contracts from a source (production)
database into the destination's arch_evtlog table, then fills arch_tx and
arch_block for every referenced transaction and block.

arch_evtlog rows are keyed by (tx_hash, log_index) so archives stay valid if
evt_log ids change; incremental export resumes from the per-contract minimum
MAX(evt_id) already archived on the destination.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			projects, err := deps.resolveProjects(projectType)
			if err != nil {
				return err
			}

			srcDB, err := deps.openDB("postgres", srcConn)
			if err != nil {
				return fmt.Errorf("connect to source: %w", err)
			}
			defer func() { _ = srcDB.Close() }()
			srcDB.SetMaxOpenConns(2)

			dstDB, err := deps.openDB("postgres", dstConn)
			if err != nil {
				return fmt.Errorf("connect to destination: %w", err)
			}
			defer func() { _ = dstDB.Close() }()
			dstDB.SetMaxOpenConns(2)

			logger := log.New(cmd.ErrOrStderr(), "", log.LstdFlags)
			exporter := deps.newExporter(srcDB, dstDB, logger)
			if _, err := deps.exportProjects(cmd.Context(), projects, exporter); err != nil {
				return err
			}
			logger.Println("=== All exports complete ===")
			return nil
		},
	}
	cmd.Flags().StringVar(&srcConn, "src", "", "Source DB connection string (production)")
	cmd.Flags().StringVar(&dstConn, "dst", "", "Destination DB connection string (dev)")
	cmd.Flags().StringVar(&projectType, "project", "", "Project: randomwalk | cosmicgame | both (runs cosmicgame then randomwalk)")
	_ = cmd.MarkFlagRequired("src")
	_ = cmd.MarkFlagRequired("dst")
	_ = cmd.MarkFlagRequired("project")
	return cmd
}
