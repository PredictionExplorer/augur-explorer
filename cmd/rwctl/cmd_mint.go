package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
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

			s, err := newTxSession(cmd, verbose)
			if err != nil {
				return err
			}
			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, s.Net.Client)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}

			s.Out.Section("MINT INFO")
			s.Out.KeyValueEth("Amount (wei)", amount)

			gasLimit := ethtx.GasLimitHighComplexity
			totalNeeded := new(big.Int).Mul(s.Net.GasPrice, big.NewInt(int64(gasLimit)))
			totalNeeded.Add(totalNeeded, amount)
			if s.Acc.Balance.Cmp(totalNeeded) < 0 {
				return fmt.Errorf("insufficient balance: need %s ETH (amount + gas), have %s ETH",
					ethtx.WeiToEthText(totalNeeded), ethtx.WeiToEthText(s.Acc.Balance))
			}

			txopts := s.TransactOpts(amount, gasLimit)
			s.Out.TxSubmitting("Mint", amount, gasLimit, ethtx.AdjustGasPrice(s.Net.GasPrice))

			tx, err := rwalk.Mint(txopts)
			return s.FinishTx(cmd.Context(), tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newMintCmd()) }
