package main

import (
	"fmt"
	"math/big"

	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newClaimPrizeCmd builds the claim-prize subcommand.
func newClaimPrizeCmd() *cobra.Command {
	var info bool
	var delaySeconds int64 = -1
	c := &cobra.Command{
		Use:   "claim-prize <cosmicgame-addr>",
		Short: "Claim the CosmicGame main prize",
		Long: `Claim the main prize from CosmicGame.

With --delay, the command first sets delayDurationBeforeRoundActivation to the
given number of seconds and then claims, so the next round activates only after
that delay (this replaces the old claimprize-delay60 script; --delay without a
value uses 60 seconds).

` + txEnvHelp,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Flags().Changed("delay") {
				return runClaimPrizeWithDelay(cmd, info, args[0], delaySeconds)
			}
			return runClaimPrize(cmd, info, args[0])
		},
	}
	addInfoFlag(c, &info)
	c.Flags().Int64Var(&delaySeconds, "delay", 60, "set delayDurationBeforeRoundActivation to this many seconds before claiming")
	c.Flags().Lookup("delay").NoOptDefVal = "60"
	return c
}

func init() { register(newClaimPrizeCmd()) }

func runClaimPrize(cmd *cobra.Command, verbose bool, addrArg string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", addrArg)
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

	roundNum, err := game.RoundNum(copts)
	if err != nil {
		return fmt.Errorf("getting round number: %w", err)
	}
	lastBidder, err := game.LastBidderAddress(copts)
	if err != nil {
		return fmt.Errorf("getting last bidder: %w", err)
	}
	prizeAmount, err := game.GetMainEthPrizeAmount(copts)
	if err != nil {
		return fmt.Errorf("getting prize amount: %w", err)
	}
	durationUntilPrize, err := game.GetDurationUntilMainPrize(copts)
	if err != nil {
		return fmt.Errorf("getting duration until prize: %w", err)
	}

	s.Out.Section("PRIZE INFO")
	s.Out.KeyValue("Round Number", roundNum.String())
	s.Out.KeyValue("Last Bidder", lastBidder.String())
	s.Out.KeyValueEth("Prize Amount", prizeAmount)
	s.Out.KeyValueDuration("Time Until Prize", durationUntilPrize.Int64())

	if s.Acc.Address != lastBidder {
		s.Out.Section("WARNING")
		s.Out.KeyValue("Your Address", s.Acc.Address.String())
		s.Out.KeyValue("Last Bidder", lastBidder.String())
		s.Out.KeyValue("Note", "You are NOT the last bidder. Claim may fail unless timeout has passed.")
	}
	if durationUntilPrize.Int64() > 0 {
		s.Out.Section("WARNING")
		s.Out.KeyValue("Status", "Prize is NOT yet claimable")
		s.Out.KeyValueDuration("Wait Time Remaining", durationUntilPrize.Int64())
	}

	s.Out.TxSubmitting("ClaimMainPrize", nil, ethtx.GasLimitClaimPrize, s.AdjustedGasPrice())
	tx, err := game.ClaimMainPrize(s.TransactOpts(nil, ethtx.GasLimitClaimPrize))
	return s.FinishTx(cmd.Context(), tx, err)
}

// runClaimPrizeWithDelay sets delayDurationBeforeRoundActivation, waits for
// that transaction to be mined, and then claims the prize, so the new round
// activates after the requested delay. The legacy script fired the claim
// after a blind two-second sleep; both steps now wait for their receipts.
func runClaimPrizeWithDelay(cmd *cobra.Command, verbose bool, addrArg string, delaySeconds int64) error {
	gameAddr, err := parseAddress("cosmicgame-addr", addrArg)
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
	prizeAmount, err := game.GetMainEthPrizeAmount(copts)
	if err != nil {
		return fmt.Errorf("getting prize amount: %w", err)
	}

	s.Out.Section("CURRENT STATE")
	s.Out.KeyValueDuration("Current Delay", currentDelay.Int64())
	s.Out.KeyValueDuration("New Delay To Set", delaySeconds)
	s.Out.KeyValueEth("Prize Amount", prizeAmount)

	s.Out.Section("STEP 1: SET DELAY")
	s.Out.TxSubmitting("SetDelayDurationBeforeRoundActivation", nil, ethtx.GasLimitAdminCall, s.AdjustedGasPrice())
	tx1, err := game.SetDelayDurationBeforeRoundActivation(
		s.TransactOpts(nil, ethtx.GasLimitAdminCall),
		big.NewInt(delaySeconds),
	)
	if err = s.FinishTx(cmd.Context(), tx1, err); err != nil {
		return fmt.Errorf("failed to set delay, aborting: %w", err)
	}

	// Refresh the nonce and gas price for the second transaction.
	if err := s.Refresh(cmd.Context()); err != nil {
		return fmt.Errorf("network refresh failed: %w", err)
	}

	s.Out.Section("STEP 2: CLAIM PRIZE")
	s.Out.TxSubmitting("ClaimMainPrize", nil, ethtx.GasLimitClaimPrize, s.AdjustedGasPrice())
	tx2, err := game.ClaimMainPrize(s.TransactOpts(nil, ethtx.GasLimitClaimPrize))
	if err = s.FinishTx(cmd.Context(), tx2, err); err != nil {
		return err
	}

	s.Out.Section("SUMMARY")
	s.Out.KeyValueDuration("Delay set to", delaySeconds)
	s.Out.KeyValue("Status", "Prize claimed successfully")
	s.Out.KeyValue("Note", fmt.Sprintf("New round will activate %d seconds after the claim", delaySeconds))
	return nil
}
