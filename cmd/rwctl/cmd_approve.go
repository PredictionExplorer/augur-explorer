package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newApproveCmd builds the approve subcommand (legacy approve script).
func newApproveCmd() *cobra.Command {
	var verbose bool
	c := &cobra.Command{
		Use:   "approve [rwalk_addr] [operator_addr]",
		Short: "Set ERC-721 approval-for-all on the RandomWalk token",
		Long:  "Sets ERC721 approval for all (SetApprovalForAll) for the given operator.\n\n" + txEnvHelp,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			rwalkAddr := common.HexToAddress(args[0])
			operatorAddr := common.HexToAddress(args[1])

			s, err := newTxSession(cmd, verbose)
			if err != nil {
				return err
			}
			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, s.Net.Client)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}

			s.Out.Section("APPROVAL INFO")
			s.Out.KeyValue("Operator", operatorAddr.String())

			gasLimit := ethtx.GasLimitApprove
			txopts := s.TransactOpts(big.NewInt(0), gasLimit)
			s.Out.TxSubmitting("SetApprovalForAll", big.NewInt(0), gasLimit, ethtx.AdjustGasPrice(s.Net.GasPrice))

			tx, err := rwalk.SetApprovalForAll(txopts, operatorAddr, true)
			return s.FinishTx(cmd.Context(), tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newApproveCmd()) }
