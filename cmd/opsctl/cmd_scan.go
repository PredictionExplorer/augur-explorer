package main

import "github.com/spf13/cobra"

const scanDefaultFilterBatchBlocks uint64 = 100_000

// newScanCmd groups ad-hoc chain scanners that look for specific events
// directly on the node.
func newScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Ad-hoc chain scanners for specific events",
	}
	cmd.AddCommand(newScanCstAuctionLenCmd())
	return cmd
}
