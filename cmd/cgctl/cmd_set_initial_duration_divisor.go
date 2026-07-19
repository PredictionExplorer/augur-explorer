package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newSetInitialDurationDivisorCmd builds the set-initial-duration-divisor
// subcommand.
func newSetInitialDurationDivisorCmd() *cobra.Command {
	var info bool
	c := &cobra.Command{
		Use:   "set-initial-duration-divisor <cosmicgame-addr> <divisor>",
		Short: "Set initialDurationUntilMainPrizeDivisor (owner only)",
		Long: `Set the initial duration until main prize divisor (e.g. 100 = 1% bump on
the first bid). Initial timer after the first bid equals
mainPrizeTimeIncrementInMicroSeconds / divisor.

` + txEnvHelp,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSetInitialDurationDivisor(cmd, info, args)
		},
	}
	addInfoFlag(c, &info)
	return c
}

func init() { register(newSetInitialDurationDivisorCmd()) }

func runSetInitialDurationDivisor(cmd *cobra.Command, verbose bool, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", args[0])
	if err != nil {
		return err
	}
	divisor, err := parseInt64("divisor", args[1])
	if err != nil {
		return err
	}
	if divisor <= 0 {
		return errors.New("divisor must be positive")
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

	copts := ethtx.CallOpts()

	currentDivisor, err := game.InitialDurationUntilMainPrizeDivisor(copts)
	if err != nil {
		return fmt.Errorf("getting current divisor: %w", err)
	}
	owner, err := game.Owner(copts)
	if err != nil {
		return fmt.Errorf("getting contract owner: %w", err)
	}

	s.Out.Section("CURRENT STATE")
	s.Out.KeyValue("Contract Owner", owner.String())
	s.Out.KeyValue("Current Divisor", currentDivisor.String())
	s.Out.KeyValue("Current Percentage", ethtx.ConvertToPercentage(currentDivisor))

	s.Out.Section("NEW VALUES")
	s.Out.KeyValue("New Divisor", divisor)
	s.Out.KeyValue("New Percentage", ethtx.ConvertToPercentage(big.NewInt(divisor)))
	s.Out.KeyValue("Formula", "percentage = 100 / divisor")

	if s.Acc.Address != owner {
		s.Out.Section("WARNING")
		s.Out.KeyValue("Your Address", s.Acc.Address.String())
		s.Out.KeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	s.Out.TxSubmitting("SetInitialDurationUntilMainPrizeDivisor", nil, ethtx.GasLimitAdminCall, s.AdjustedGasPrice())
	tx, err := game.SetInitialDurationUntilMainPrizeDivisor(s.TransactOpts(nil, ethtx.GasLimitAdminCall), big.NewInt(divisor))
	return s.FinishTx(cmd.Context(), tx, err)
}
