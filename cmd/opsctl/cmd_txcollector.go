package main

import "github.com/spf13/cobra"

// txCollectorCmd groups the RLP transaction/receipt backup tools. Both
// subcommands read the same JSON config (rpc_url, output_dir, start_block,
// contract addresses) — see tx-collector.example.json in this directory.
var txCollectorCmd = &cobra.Command{
	Use:   "tx-collector",
	Short: "RLP transaction/receipt backup collector and verifier",
}

func init() { register(txCollectorCmd) }
