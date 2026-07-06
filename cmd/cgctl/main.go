// Command cgctl is the operator CLI for the CosmicGame contracts.
//
// It consolidates the former one-off scripts from rwcg/etl/cosmicgame/scripts
// into a single binary with one subcommand per operation (bidding, prize
// claiming, owner-only parameter changes, ERC-20/ERC-721 helpers, and
// diagnostics). See README.md for the full mapping of old scripts to
// subcommands.
//
// Common configuration comes from environment variables:
//
//	RPC_URL     Ethereum/Arbitrum JSON-RPC endpoint
//	PKEY_HEX    hex-encoded private key used to sign transactions
//	PGSQL_*     PostgreSQL connection (only for subcommands that read the DB)
//
// Run `cgctl --help` for the full list of subcommands.
package main

import (
	"os"

	"github.com/spf13/cobra"
)

// subcommands is populated by register calls in the per-command files.
var subcommands []*cobra.Command

// register adds a subcommand to the root command. Each subcommand file calls
// it from an init function so commands stay self-contained.
func register(c *cobra.Command) { subcommands = append(subcommands, c) }

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "cgctl",
		Short:         "Operator CLI for the CosmicGame contracts",
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
