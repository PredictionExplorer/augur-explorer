package main

import "github.com/spf13/cobra"

// scanCmd groups ad-hoc chain scanners that look for specific events directly
// on the node.
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Ad-hoc chain scanners for specific events",
}

func init() { register(scanCmd) }
