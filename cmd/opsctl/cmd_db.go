package main

import "github.com/spf13/cobra"

// newDBCmd groups database-to-database comparison tools used when migrating or
// re-syncing an indexer database.
func newDBCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "Database comparison tools (primary vs secondary)",
	}
	cmd.AddCommand(
		newDBEvtlogDiffCmd(),
		newDBVerifyCmd(),
	)
	return cmd
}
