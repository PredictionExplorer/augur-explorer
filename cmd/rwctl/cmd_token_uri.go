package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// newTokenURICmd builds the token-uri subcommand (legacy tokenuri script).
func newTokenURICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "token-uri [rwalk_addr] [token_id]",
		Short: "Get the token URI of a RandomWalk token",
		Long:  "Gets token URI from RandomWalk contract.\n\nEnvironment variables:\n  RPC_URL  Ethereum RPC endpoint (required)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			eclient, err := dialEthClient()
			if err != nil {
				return err
			}
			rwalkAddr := common.HexToAddress(args[0])
			tokenID, err := parseInt64("token_id", args[1])
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Calling to contract at %v\n", rwalkAddr.String())

			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, eclient)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}
			uri, err := rwalk.TokenURI(callOpts(), big.NewInt(tokenID))
			if err != nil {
				return fmt.Errorf("error at TokenURI(): %w", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "URI = %v\n", uri)
			return nil
		},
	}
}

func init() { register(newTokenURICmd()) }
