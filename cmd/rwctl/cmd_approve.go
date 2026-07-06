package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
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

			s, err := newTxSession(verbose)
			if err != nil {
				return err
			}
			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, s.net.client)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}

			s.out.section("APPROVAL INFO")
			s.out.keyValue("Operator", operatorAddr.String())

			gasLimit := gasLimitERC721Approve
			txopts := transactOpts(s.net, s.acc, big.NewInt(0), gasLimit)
			s.out.txSubmitting("SetApprovalForAll", big.NewInt(0), gasLimit, adjustGasPrice(s.net.gasPrice))

			tx, err := rwalk.SetApprovalForAll(txopts, operatorAddr, true)
			return s.finishTx(tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newApproveCmd()) }
