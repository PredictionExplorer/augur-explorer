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
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "opsctl",
		Short:         "Data-operations utilities for the RWCG backend",
		Version:       version.String(),
		SilenceUsage:  true,
		SilenceErrors: false,
	}
	root.AddCommand(
		newArchiveCmd(),
		newAssetsCmd(),
		newDBCmd(),
		newScanCmd(),
		newSmoketestCmd(),
		newTxCollectorCmd(),
	)
	return root
}

// osExit is stubbed by tests that drive main through its failure arm.
var osExit = os.Exit

func main() {
	if err := runMain(os.Args[1:]); err != nil {
		osExit(1)
	}
}

// runMain owns the signal-scoped context so its deferred cleanup always runs
// before main decides the exit code (os.Exit skips deferred calls).
func runMain(args []string) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	root := newRootCmd()
	root.SetArgs(args)
	return root.ExecuteContext(ctx)
}
