package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// newStatusCmd builds the status subcommand (legacy status script).
func newStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status [rwalk_addr]",
		Short: "Read status variables from the RandomWalk contract",
		Long:  "Reads variables from RandomWalk contract.\n\nEnvironment variables:\n  RPC_URL  Ethereum RPC endpoint (required)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			eclient, err := dialEthClient()
			if err != nil {
				return err
			}
			rwalkAddr := common.HexToAddress(args[0])
			fmt.Printf("Calling to contract at %v\n", rwalkAddr.String())

			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, eclient)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}
			copts := callOpts()
			nextTokenID, err := rwalk.NextTokenId(copts)
			if err != nil {
				return fmt.Errorf("error at NextTokenId(): %w", err)
			}
			timeRemaining, err := rwalk.TimeUntilWithdrawal(copts)
			if err != nil {
				return fmt.Errorf("error at timeUntilWithdrawal(): %w", err)
			}
			withdrawalAmount, err := rwalk.WithdrawalAmount(copts)
			if err != nil {
				return fmt.Errorf("error at withdrawalAmount(): %w", err)
			}
			numWithdrawals, err := rwalk.NumWithdrawals(copts)
			if err != nil {
				return fmt.Errorf("error at numWithdrawals(): %w", err)
			}
			lastMinter, err := rwalk.LastMinter(copts)
			if err != nil {
				return fmt.Errorf("error at lastMinter(): %w", err)
			}
			baseURI, err := rwalk.TokenURI(copts, big.NewInt(0))
			if err != nil {
				return fmt.Errorf("error at TokenURI(): %w", err)
			}

			fmt.Printf("Next token ID = %v\n", nextTokenID.Int64())
			fmt.Printf("Time remaining: %v\n", timeRemaining.Int64())
			fmt.Printf("Withdrawal amount: %v\n", weiToEthText(withdrawalAmount))
			fmt.Printf("Num withdrawals: %v\n", numWithdrawals.Int64())
			fmt.Printf("Last minter: %v\n", lastMinter.Hex())
			fmt.Printf("Base uri: %v\n", baseURI)
			return nil
		},
	}
}

func init() { register(newStatusCmd()) }
