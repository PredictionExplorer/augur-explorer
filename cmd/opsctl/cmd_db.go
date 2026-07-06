package main

import "github.com/spf13/cobra"

// dbCmd groups database-to-database comparison tools used when migrating or
// re-syncing an indexer database.
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database comparison tools (primary vs secondary)",
}

func init() { register(dbCmd) }
