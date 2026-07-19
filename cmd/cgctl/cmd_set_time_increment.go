package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newSetTimeIncrementCmd builds the set-time-increment subcommand.
func newSetTimeIncrementCmd() *cobra.Command {
	var info bool
	c := &cobra.Command{
		Use:   "set-time-increment <cosmicgame-addr> <time-increment-seconds>",
		Short: "Set mainPrizeTimeIncrementInMicroSeconds (owner only)",
		Long: `Set mainPrizeTimeIncrementInMicroSeconds so that each bid extends the time
until the main prize by the given number of seconds. Requires an inactive
round; see claim-and-set-time-increment for a variant that opens the inactive
window automatically.

` + txEnvHelp,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSetTimeIncrement(cmd, info, args)
		},
	}
	addInfoFlag(c, &info)
	return c
}

func init() { register(newSetTimeIncrementCmd()) }

func runSetTimeIncrement(cmd *cobra.Command, verbose bool, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", args[0])
	if err != nil {
		return err
	}
	desiredSeconds, err := parseInt64("time_increment_seconds", args[1])
	if err != nil {
		return err
	}
	if desiredSeconds <= 0 {
		return errors.New("time_increment_seconds must be positive")
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

	currentMicroseconds, err := game.MainPrizeTimeIncrementInMicroSeconds(copts)
	if err != nil {
		return fmt.Errorf("reading mainPrizeTimeIncrementInMicroSeconds: %w", err)
	}
	currentSeconds := new(big.Int).Div(currentMicroseconds, big.NewInt(1000000))

	owner, err := game.Owner(copts)
	if err != nil {
		return fmt.Errorf("getting contract owner: %w", err)
	}

	newMicroseconds := new(big.Int).Mul(big.NewInt(desiredSeconds), big.NewInt(1000000))

	s.Out.Section("CURRENT STATE")
	s.Out.KeyValue("Contract Owner", owner.String())
	s.Out.KeyValue("Current Microseconds", currentMicroseconds.String())
	s.Out.KeyValueDuration("Current Time Increment", currentSeconds.Int64())

	s.Out.Section("NEW VALUES")
	s.Out.KeyValue("New Microseconds", newMicroseconds.String())
	s.Out.KeyValueDuration("New Time Increment", desiredSeconds)
	s.Out.KeyValue("Formula", "timeIncrement (seconds) = microseconds / 1,000,000")

	if s.Acc.Address != owner {
		s.Out.Section("WARNING")
		s.Out.KeyValue("Your Address", s.Acc.Address.String())
		s.Out.KeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	s.Out.TxSubmitting("SetMainPrizeTimeIncrementInMicroSeconds", nil, ethtx.GasLimitAdminCall, s.AdjustedGasPrice())
	tx, err := game.SetMainPrizeTimeIncrementInMicroSeconds(s.TransactOpts(nil, ethtx.GasLimitAdminCall), newMicroseconds)
	return s.FinishTx(cmd.Context(), tx, err)
}
