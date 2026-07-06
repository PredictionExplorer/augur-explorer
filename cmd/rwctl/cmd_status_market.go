package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// newStatusMarketCmd builds the status-market subcommand (legacy statusmkt script).
func newStatusMarketCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status-market [market_addr]",
		Short: "Read status variables from the RandomWalk marketplace contract",
		Long:  "Reads variables from Market contract.\n\nEnvironment variables:\n  RPC_URL  Ethereum RPC endpoint (required)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			eclient, err := dialEthClient()
			if err != nil {
				return err
			}
			marketAddr := common.HexToAddress(args[0])
			fmt.Printf("Calling to contract at %v\n", marketAddr.String())

			market, err := rwcontracts.NewRWMarket(marketAddr, eclient)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWMarket contract: %w", err)
			}
			numOffers, err := market.NumOffers(callOpts())
			if err != nil {
				return fmt.Errorf("error at NumOffers(): %w", err)
			}
			fmt.Printf("NumOffers = %v\n", numOffers.Int64())
			return nil
		},
	}
}

func init() { register(newStatusMarketCmd()) }
