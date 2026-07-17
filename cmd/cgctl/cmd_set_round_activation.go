package main

import (
	"fmt"
	"math/big"

	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newSetRoundActivationCmd builds the set-round-activation subcommand.
func newSetRoundActivationCmd() *cobra.Command {
	var info bool
	c := &cobra.Command{
		Use:   "set-round-activation <cosmicgame-addr> <timestamp>",
		Short: "Set roundActivationTime (owner only)",
		Long:  "Set the round activation time to the given Unix timestamp.\n\n" + txEnvHelp,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSetRoundActivation(cmd, info, args)
		},
	}
	addInfoFlag(c, &info)
	return c
}

func init() { register(newSetRoundActivationCmd()) }

func runSetRoundActivation(cmd *cobra.Command, verbose bool, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", args[0])
	if err != nil {
		return err
	}
	timestamp, err := parseInt64("timestamp", args[1])
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

	currentActivation, err := game.RoundActivationTime(copts)
	if err != nil {
		return fmt.Errorf("getting current activation time: %w", err)
	}
	owner, err := game.Owner(copts)
	if err != nil {
		return fmt.Errorf("getting contract owner: %w", err)
	}

	blockTime := int64(s.Net.BlockTime) // #nosec G115 -- real chain timestamps fit int64; display-only CLI
	secsUntilCurrent := currentActivation.Int64() - blockTime
	secsUntilNew := timestamp - blockTime

	s.Out.Section("CURRENT STATE")
	s.Out.KeyValue("Contract Owner", owner.String())
	s.Out.KeyValue("Current Block Time", s.Net.BlockTime)
	s.Out.KeyValue("Current Activation Time", currentActivation.String())
	s.Out.KeyValueDuration("Time Until Current Activation", secsUntilCurrent)

	s.Out.Section("NEW VALUES")
	s.Out.KeyValue("New Activation Time", timestamp)
	s.Out.KeyValueDuration("Time Until New Activation", secsUntilNew)

	if s.Acc.Address != owner {
		s.Out.Section("WARNING")
		s.Out.KeyValue("Your Address", s.Acc.Address.String())
		s.Out.KeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	s.Out.TxSubmitting("SetRoundActivationTime", nil, ethtx.GasLimitAdminCall, s.AdjustedGasPrice())
	tx, err := game.SetRoundActivationTime(s.TransactOpts(nil, ethtx.GasLimitAdminCall), big.NewInt(timestamp))
	return s.FinishTx(cmd.Context(), tx, err)
}
