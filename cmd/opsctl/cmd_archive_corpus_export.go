package main

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/archive"
)

type archiveCorpusExportDeps struct {
	openDB func(context.Context, string) (opsDB, error)
	export func(
		context.Context,
		archive.Querier,
		archive.CorpusExportOptions,
		io.Writer,
	) (archive.CorpusExportStats, error)
}

func defaultArchiveCorpusExportDeps() archiveCorpusExportDeps {
	return archiveCorpusExportDeps{
		openDB: openOpsDB(archiveMaxConns),
		export: archive.ExportCorpus,
	}
}

// newArchiveCorpusExportCmd builds `opsctl archive corpus-export`, which
// renders complete selected archive transactions as strict RLP corpus JSONL.
func newArchiveCorpusExportCmd() *cobra.Command {
	return newArchiveCorpusExportCmdWithDeps(defaultArchiveCorpusExportDeps())
}

func newArchiveCorpusExportCmdWithDeps(deps archiveCorpusExportDeps) *cobra.Command {
	var (
		connString string
		project    string
		txHashes   []string
	)
	cmd := &cobra.Command{
		Use:   "corpus-export",
		Short: "Write selected complete arch_evtlog transactions as strict RLP corpus JSONL",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			db, err := deps.openDB(cmd.Context(), connString)
			if err != nil {
				return fmt.Errorf("connect to archive database: %w", err)
			}
			defer db.Close()

			stats, err := deps.export(
				cmd.Context(),
				db,
				archive.CorpusExportOptions{
					Project:  project,
					TxHashes: txHashes,
				},
				cmd.OutOrStdout(),
			)
			if err != nil {
				return err
			}
			fmt.Fprintf(
				cmd.ErrOrStderr(),
				"exported %d event logs from %d complete transactions\n",
				stats.EventLogs,
				stats.Transactions,
			)
			return nil
		},
	}
	cmd.Flags().StringVar(&connString, "db", "", "Archive database connection string")
	cmd.Flags().StringVar(&project, "project", "", "Project label: randomwalk | cosmicgame")
	cmd.Flags().StringArrayVar(
		&txHashes,
		"tx-hash",
		nil,
		"Transaction hash to include with every sibling event log (repeatable)",
	)
	_ = cmd.MarkFlagRequired("db")
	_ = cmd.MarkFlagRequired("project")
	_ = cmd.MarkFlagRequired("tx-hash")
	return cmd
}
