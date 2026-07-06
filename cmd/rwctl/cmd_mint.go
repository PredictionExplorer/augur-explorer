package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// newMintCmd builds the mint subcommand (legacy mint script).
func newMintCmd() *cobra.Command {
	var verbose bool
	c := &cobra.Command{
		Use:   "mint [rwalk_addr] [amount_wei]",
		Short: "Mint a RandomWalk token",
		Long:  "Mints a RandomWalk token by sending amount_wei to the contract.\n\n" + txEnvHelp,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			rwalkAddr := common.HexToAddress(args[0])
			amount, err := parseBigInt("amount_wei", args[1])
			if err != nil {
				return err
			}

			s, err := newTxSession(verbose)
			if err != nil {
				return err
			}
			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, s.net.client)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}

			s.out.section("MINT INFO")
			s.out.keyValueEth("Amount (wei)", amount)

			gasLimit := gasLimitHighComplexity
			totalNeeded := new(big.Int).Mul(s.net.gasPrice, big.NewInt(int64(gasLimit)))
			totalNeeded.Add(totalNeeded, amount)
			if s.acc.balance.Cmp(totalNeeded) < 0 {
				return fmt.Errorf("insufficient balance: need %s ETH (amount + gas), have %s ETH",
					weiToEthText(totalNeeded), weiToEthText(s.acc.balance))
			}

			txopts := transactOpts(s.net, s.acc, amount, gasLimit)
			s.out.txSubmitting("Mint", amount, gasLimit, adjustGasPrice(s.net.gasPrice))

			tx, err := rwalk.Mint(txopts)
			return s.finishTx(tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newMintCmd()) }
