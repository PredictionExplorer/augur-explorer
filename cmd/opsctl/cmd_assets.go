package main

import (
	"github.com/spf13/cobra"
)

// newAssetsCmd groups the NFT asset utilities: on-disk inventory, thumbnail
// generation and token image URL verification.
func newAssetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assets",
		Short: "NFT asset utilities (inventory, thumbnails, image checks)",
	}
	cmd.AddCommand(
		newAssetsGenThumbnailsCmd(),
		newAssetsInventoryCmd(),
		newAssetsVerifyTokenImagesCmd(),
	)
	return cmd
}
