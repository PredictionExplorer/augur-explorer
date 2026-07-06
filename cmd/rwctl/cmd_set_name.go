package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// newSetNameCmd builds the set-name subcommand (legacy setname script).
func newSetNameCmd() *cobra.Command {
	var verbose bool
	c := &cobra.Command{
		Use:   "set-name [rwalk_addr] [token_id] [new_name]",
		Short: "Set the display name of a RandomWalk token",
		Long:  "Sets the display name for a RandomWalk token.\n\n" + txEnvHelp,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			rwalkAddr := common.HexToAddress(args[0])
			tokenID, err := parseInt64("token_id", args[1])
			if err != nil {
				return err
			}
			newName := args[2]

			s, err := newTxSession(verbose)
			if err != nil {
				return err
			}
			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, s.net.client)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}

			s.out.section("SET NAME INFO")
			s.out.keyValue("Token ID", tokenID)
			s.out.keyValue("New Name", newName)

			gasLimit := gasLimitContractCall
			txopts := transactOpts(s.net, s.acc, big.NewInt(0), gasLimit)
			s.out.txSubmitting("SetTokenName", big.NewInt(0), gasLimit, adjustGasPrice(s.net.gasPrice))

			tx, err := rwalk.SetTokenName(txopts, big.NewInt(tokenID), newName)
			return s.finishTx(tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newSetNameCmd()) }
