package main

import "github.com/spf13/cobra"

// newArchiveCmd groups the tools that maintain the arch_evtlog / arch_tx /
// arch_block archive tables: export from a live database, render strict RLP
// corpora, verify consistency, and backfill from an Ethereum node.
func newArchiveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "archive",
		Short: "Event-log archival tools (arch_evtlog / arch_tx / arch_block)",
	}
	cmd.AddCommand(
		newArchiveCorpusExportCmd(),
		newArchiveExportCmd(),
		newArchiveNodeFillCmd(),
		newArchiveVerifyCmd(),
	)
	return cmd
}
