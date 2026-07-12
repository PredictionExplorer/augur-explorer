package main

import (
	"fmt"
	"math/big"

	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newSetActivationDelayCmd builds the set-activation-delay subcommand.
func newSetActivationDelayCmd() *cobra.Command {
	var info bool
	c := &cobra.Command{
		Use:   "set-activation-delay <cosmicgame-addr> <seconds>",
		Short: "Set delayDurationBeforeRoundActivation (owner only)",
		Long:  "Set the delay duration before the next round activates after a prize claim.\n\n" + txEnvHelp,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSetActivationDelay(cmd, info, args)
		},
	}
	addInfoFlag(c, &info)
	return c
}

func init() { register(newSetActivationDelayCmd()) }

func runSetActivationDelay(cmd *cobra.Command, verbose bool, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", args[0])
	if err != nil {
		return err
	}
	seconds, err := parseInt64("seconds", args[1])
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

	copts := ethtx.CallOpts()

	currentDelay, err := game.DelayDurationBeforeRoundActivation(copts)
	if err != nil {
		return fmt.Errorf("getting current delay: %w", err)
	}
	owner, err := game.Owner(copts)
	if err != nil {
		return fmt.Errorf("getting contract owner: %w", err)
	}

	s.Out.Section("CURRENT STATE")
	s.Out.KeyValue("Contract Owner", owner.String())
	s.Out.KeyValueDuration("Current Delay", currentDelay.Int64())
	s.Out.KeyValueDuration("New Delay", seconds)

	if s.Acc.Address != owner {
		s.Out.Section("WARNING")
		s.Out.KeyValue("Your Address", s.Acc.Address.String())
		s.Out.KeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	s.Out.TxSubmitting("SetDelayDurationBeforeRoundActivation", nil, ethtx.GasLimitAdminCall, s.AdjustedGasPrice())
	tx, err := game.SetDelayDurationBeforeRoundActivation(s.TransactOpts(nil, ethtx.GasLimitAdminCall), big.NewInt(seconds))
	return s.FinishTx(cmd.Context(), tx, err)
}
