// Command rwctl is the operator CLI for the RandomWalk NFT and its
// marketplace, plus the social-media utilities that accompany the project.
//
// It consolidates the former one-off scripts under etl/randomwalk/scripts and
// etl/randomwalk/tools into a single binary with one subcommand per operation
// (minting, offers, transfers, verification scans, Twitter/Discord tooling).
//
// Common configuration comes from environment variables:
//
//	RPC_URL             Ethereum/Arbitrum JSON-RPC endpoint
//	PGSQL_*             PostgreSQL connection (DB-backed subcommands only)
//	TWITTER_KEYS_FILE   JSON file with Twitter API credentials
//	DISGORD_TOKEN       Discord bot token
//
// Run `rwctl --help` for the full list of subcommands.
package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

// subcommands is populated by register calls in the per-command files.
var subcommands []*cobra.Command

// register adds a subcommand to the root command. Each subcommand file calls
// it from an init function so commands stay self-contained.
func register(c *cobra.Command) { subcommands = append(subcommands, c) }

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "rwctl",
		Short:         "Operator CLI for the RandomWalk NFT, marketplace and social tools",
		Version:       version.String(),
		SilenceUsage:  true,
		SilenceErrors: false,
	}
	root.AddCommand(subcommands...)
	return root
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
