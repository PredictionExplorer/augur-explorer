package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
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

			s, err := newTxSession(cmd, verbose)
			if err != nil {
				return err
			}
			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, s.Net.Client)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}

			s.Out.Section("SET NAME INFO")
			s.Out.KeyValue("Token ID", tokenID)
			s.Out.KeyValue("New Name", newName)

			gasLimit := ethtx.GasLimitContractCall
			txopts := s.TransactOpts(big.NewInt(0), gasLimit)
			s.Out.TxSubmitting("SetTokenName", big.NewInt(0), gasLimit, ethtx.AdjustGasPrice(s.Net.GasPrice))

			tx, err := rwalk.SetTokenName(txopts, big.NewInt(tokenID), newName)
			return s.FinishTx(cmd.Context(), tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newSetNameCmd()) }
