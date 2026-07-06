// Command opsctl bundles the data-operations utilities for the RWCG backend:
// event-log archival (archive export / verify / node-fill), database
// comparison (db verify / evtlog-diff), transaction backup (tx-collector
// run / verify), NFT asset utilities (assets inventory / gen-thumbnails /
// verify-token-images), API smoke testing (smoketest) and ad-hoc chain
// scanners (scan cst-auction-len).
//
// It consolidates the former standalone tools from rwcg/tools/ (plus the
// notibot verify-token-imgs script) into a single binary with one subcommand
// per utility.
//
// Common configuration comes from environment variables:
//
//	RPC_URL     Ethereum/Arbitrum JSON-RPC endpoint (chain-reading subcommands)
//	PGSQL_*     PostgreSQL connection
//
// Run `opsctl --help` for the full list of subcommands; see README.md in this
// directory for per-command flags and environment variables.
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
		Use:           "opsctl",
		Short:         "Data-operations utilities for the RWCG backend",
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
