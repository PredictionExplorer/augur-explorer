package main

import "github.com/spf13/cobra"

const txCollectorDefaultFilterBatchBlocks uint64 = 100_000

// newTxCollectorCmd groups the RLP transaction/receipt backup tools. Both
// subcommands read the same JSON config (rpc_url, output_dir, start_block,
// contract addresses) — see tx-collector.example.json in this directory.
func newTxCollectorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx-collector",
		Short: "RLP transaction/receipt backup collector and verifier",
	}
	cmd.AddCommand(
		newTxCollectorRunCmd(),
		newTxCollectorVerifyCmd(),
	)
	return cmd
}
