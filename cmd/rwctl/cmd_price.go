package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// newPriceCmd builds the price subcommand (legacy price script).
func newPriceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "price [rwalk_addr]",
		Short: "Get the current mint price of the RandomWalk contract",
		Long:  "Gets latest price from RandomWalk contract.\n\nEnvironment variables:\n  RPC_URL  Ethereum RPC endpoint (required)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			eclient, err := dialEthClient()
			if err != nil {
				return err
			}
			rwalkAddr := common.HexToAddress(args[0])
			fmt.Fprintf(cmd.OutOrStdout(), "Calling to contract at %v\n", rwalkAddr.String())

			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, eclient)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}
			price, err := rwalk.GetMintPrice(callOpts())
			if err != nil {
				return fmt.Errorf("error at GetMintPrice(): %w", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Mint price = %v\n", price.String())
			return nil
		},
	}
}

func init() { register(newPriceCmd()) }
