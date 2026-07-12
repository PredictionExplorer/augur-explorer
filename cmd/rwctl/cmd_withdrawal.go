package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// newWithdrawalCmd builds the withdrawal subcommand (legacy withdrawal script).
func newWithdrawalCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "withdrawal [rwalk_addr]",
		Short: "Get the current withdrawal amount of the RandomWalk contract",
		Long:  "Gets withdrawal amount from RandomWalk contract.\n\nEnvironment variables:\n  RPC_URL  Ethereum RPC endpoint (required)",
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
			amount, err := rwalk.WithdrawalAmount(callOpts())
			if err != nil {
				return fmt.Errorf("error at WithdrawalAmount(): %w", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Withdrawal amount = %v\n", amount.String())
			return nil
		},
	}
}

func init() { register(newWithdrawalCmd()) }
