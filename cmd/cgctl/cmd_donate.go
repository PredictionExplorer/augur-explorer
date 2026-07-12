package main

import (
	"fmt"
	"math/big"

	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newDonateCmd builds the donate subcommand.
func newDonateCmd() *cobra.Command {
	var info bool
	c := &cobra.Command{
		Use:   "donate <cosmicgame-addr> <amount-wei>",
		Short: "Donate ETH to the CosmicGame contract",
		Long:  "Donate ETH to the CosmicGame contract (DonateEth).\n\n" + txEnvHelp,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDonate(cmd, info, args)
		},
	}
	addInfoFlag(c, &info)
	return c
}

func init() { register(newDonateCmd()) }

func runDonate(cmd *cobra.Command, verbose bool, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", args[0])
	if err != nil {
		return err
	}
	donationAmount, err := parseBigInt("amount", args[1])
	if err != nil {
		return err
	}

	s, err := newTxSession(cmd, verbose)
	if err != nil {
		return err
	}
	s.Out.ContractInfo("CosmicGame Address", gameAddr)
	game, err := cgcontracts.NewCosmicSignatureGame(gameAddr, s.Net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate CosmicGame: %w", err)
	}

	contractBalance, err := s.Net.Balance(cmd.Context(), gameAddr)
	if err != nil {
		return fmt.Errorf("getting contract balance: %w", err)
	}

	s.Out.Section("DONATION INFO")
	s.Out.KeyValueEth("Contract Current Balance", contractBalance)
	s.Out.KeyValueEth("Donation Amount", donationAmount)
	s.Out.KeyValueEth("Contract Balance After", new(big.Int).Add(contractBalance, donationAmount))

	if s.Acc.Balance.Cmp(donationAmount) < 0 {
		return fmt.Errorf("insufficient balance: need %s ETH, have %s ETH",
			ethtx.WeiToEthText(donationAmount), ethtx.WeiToEthText(s.Acc.Balance))
	}

	s.Out.TxSubmitting("DonateEth", donationAmount, ethtx.GasLimitDonate, s.AdjustedGasPrice())
	tx, err := game.DonateEth(s.TransactOpts(donationAmount, ethtx.GasLimitDonate))
	return s.FinishTx(cmd.Context(), tx, err)
}
