package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newClaimAndSetTimeIncrementCmd builds the claim-and-set-time-increment
// subcommand.
func newClaimAndSetTimeIncrementCmd() *cobra.Command {
	var info bool
	c := &cobra.Command{
		Use:   "claim-and-set-time-increment <cosmicgame-addr> <time-increment-seconds> [delay-seconds]",
		Short: "Set the per-bid time increment, claiming the prize first when needed",
		Long: `Set mainPrizeTimeIncrementInMicroSeconds, claiming the main prize first
when that is required to open the inactive admin window.

Contract rules (V1/V2 proxy):
  - setMainPrizeTimeIncrementInMicroSeconds requires an inactive round
  - claimMainPrize requires bids in the current round and the prize timer or timeout
  - setDelayDurationBeforeRoundActivation and setRoundActivationTime can run while active

The command picks a path from on-chain state:
  A) increment and delay already match: exit success (idempotent)
  B) round already inactive: set increment directly
  C) round active, prize claimable: claim (creates inactive window), then set increment
  D) round active, prize not claimable: defer activation, then set increment

delay-seconds defaults to 300.

` + txEnvHelp,
		Args: cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runClaimAndSetTimeIncrement(cmd, info, args)
		},
	}
	addInfoFlag(c, &info)
	return c
}

func init() { register(newClaimAndSetTimeIncrementCmd()) }

// roundState is a snapshot of the CosmicGame round used to plan which
// transactions claim-and-set-time-increment must send.
type roundState struct {
	roundNum               *big.Int
	roundActivation        *big.Int
	totalBids              *big.Int
	lastBidder             common.Address
	durationUntilPrize     *big.Int
	durationUntilPrizeRaw  *big.Int
	timeoutClaim           *big.Int
	currentIncrementMicros *big.Int
	currentDelay           *big.Int
	owner                  common.Address
	roundInactive          bool
	hasBids                bool
	incrementAlreadySet    bool
	delayAlreadySet        bool
}

func readRoundState(
	game *cgcontracts.CosmicSignatureGame,
	blockTime uint64,
	newIncrementMicros *big.Int,
	delaySeconds int64,
) (*roundState, error) {
	copts := ethtx.CallOpts()

	roundNum, err := game.RoundNum(copts)
	if err != nil {
		return nil, fmt.Errorf("roundNum: %w", err)
	}
	roundActivation, err := game.RoundActivationTime(copts)
	if err != nil {
		return nil, fmt.Errorf("roundActivationTime: %w", err)
	}
	totalBids, err := game.GetTotalNumBids(copts, roundNum)
	if err != nil {
		return nil, fmt.Errorf("getTotalNumBids: %w", err)
	}
	lastBidder, err := game.LastBidderAddress(copts)
	if err != nil {
		return nil, fmt.Errorf("lastBidderAddress: %w", err)
	}
	durationUntilPrize, err := game.GetDurationUntilMainPrize(copts)
	if err != nil {
		return nil, fmt.Errorf("getDurationUntilMainPrize: %w", err)
	}
	durationUntilPrizeRaw, err := game.GetDurationUntilMainPrizeRaw(copts)
	if err != nil {
		return nil, fmt.Errorf("getDurationUntilMainPrizeRaw: %w", err)
	}
	timeoutClaim, err := game.TimeoutDurationToClaimMainPrize(copts)
	if err != nil {
		return nil, fmt.Errorf("timeoutDurationToClaimMainPrize: %w", err)
	}
	currentIncrementMicros, err := game.MainPrizeTimeIncrementInMicroSeconds(copts)
	if err != nil {
		return nil, fmt.Errorf("mainPrizeTimeIncrementInMicroSeconds: %w", err)
	}
	currentDelay, err := game.DelayDurationBeforeRoundActivation(copts)
	if err != nil {
		return nil, fmt.Errorf("delayDurationBeforeRoundActivation: %w", err)
	}
	owner, err := game.Owner(copts)
	if err != nil {
		return nil, fmt.Errorf("owner: %w", err)
	}

	return &roundState{
		roundNum:               roundNum,
		roundActivation:        roundActivation,
		totalBids:              totalBids,
		lastBidder:             lastBidder,
		durationUntilPrize:     durationUntilPrize,
		durationUntilPrizeRaw:  durationUntilPrizeRaw,
		timeoutClaim:           timeoutClaim,
		currentIncrementMicros: currentIncrementMicros,
		currentDelay:           currentDelay,
		owner:                  owner,
		roundInactive:          blockTime < roundActivation.Uint64(),
		hasBids:                lastBidder != (common.Address{}),
		incrementAlreadySet:    currentIncrementMicros.Cmp(newIncrementMicros) == 0,
		delayAlreadySet:        currentDelay.Int64() == delaySeconds,
	}, nil
}

// canClaim reports whether caller may claim the main prize now, with a
// human-readable reason.
func (s *roundState) canClaim(caller common.Address) (bool, string) {
	if !s.hasBids {
		return false, "no bids in current round"
	}
	if s.durationUntilPrize.Int64() <= 0 && caller == s.lastBidder {
		return true, "last bidder, prize timer expired"
	}
	anyoneDeadline := new(big.Int).Add(s.durationUntilPrizeRaw, s.timeoutClaim)
	if anyoneDeadline.Sign() <= 0 {
		return true, "claim timeout expired (anyone may claim)"
	}
	return false, "prize not claimable yet"
}

func sendDelay(ctx context.Context, s *ethtx.Session, game *cgcontracts.CosmicSignatureGame, delaySeconds int64) error {
	s.Out.TxSubmitting("SetDelayDurationBeforeRoundActivation", nil, ethtx.GasLimitAdminCall, s.AdjustedGasPrice())
	tx, err := game.SetDelayDurationBeforeRoundActivation(
		s.TransactOpts(nil, ethtx.GasLimitAdminCall),
		big.NewInt(delaySeconds),
	)
	if err = s.FinishTx(ctx, tx, err); err != nil {
		return fmt.Errorf("setDelayDurationBeforeRoundActivation: %w", err)
	}
	return nil
}

func sendDeferActivation(ctx context.Context, s *ethtx.Session, game *cgcontracts.CosmicSignatureGame, delaySeconds int64) error {
	newActivation := big.NewInt(int64(s.Net.BlockTime) + delaySeconds + 5)
	s.Out.TxSubmitting("SetRoundActivationTime", nil, ethtx.GasLimitAdminCall, s.AdjustedGasPrice())
	s.Out.KeyValue("New activation time", newActivation.String())
	tx, err := game.SetRoundActivationTime(
		s.TransactOpts(nil, ethtx.GasLimitAdminCall),
		newActivation,
	)
	if err = s.FinishTx(ctx, tx, err); err != nil {
		return fmt.Errorf("setRoundActivationTime: %w", err)
	}
	return nil
}

// openInactiveWindow defers round activation until the round reads as
// inactive, retrying a few times to absorb block-time races.
func openInactiveWindow(ctx context.Context, s *ethtx.Session, game *cgcontracts.CosmicSignatureGame, delaySeconds int64) error {
	const maxAttempts = 3
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if err := s.Refresh(ctx); err != nil {
			return err
		}
		state, err := readRoundState(game, s.Net.BlockTime, big.NewInt(0), delaySeconds)
		if err != nil {
			return err
		}
		if state.roundInactive {
			return nil
		}
		if attempt > 1 || s.Out.Verbose {
			s.Out.Section(fmt.Sprintf("DEFER ROUND ACTIVATION (attempt %d/%d)", attempt, maxAttempts))
		}
		if err := sendDeferActivation(ctx, s, game, delaySeconds); err != nil {
			return fmt.Errorf("setRoundActivationTime failed on attempt %d: %w", attempt, err)
		}
	}
	if err := s.Refresh(ctx); err != nil {
		return err
	}
	state, err := readRoundState(game, s.Net.BlockTime, big.NewInt(0), delaySeconds)
	if err != nil {
		return err
	}
	if !state.roundInactive {
		return fmt.Errorf(
			"round is still active after %d defer attempts (activation=%s, block=%d)",
			maxAttempts,
			state.roundActivation.String(),
			s.Net.BlockTime,
		)
	}
	return nil
}

func sendClaim(ctx context.Context, s *ethtx.Session, game *cgcontracts.CosmicSignatureGame) error {
	s.Out.TxSubmitting("ClaimMainPrize", nil, ethtx.GasLimitClaimPrize, s.AdjustedGasPrice())
	tx, err := game.ClaimMainPrize(s.TransactOpts(nil, ethtx.GasLimitClaimPrize))
	if err = s.FinishTx(ctx, tx, err); err != nil {
		return fmt.Errorf("claimMainPrize: %w", err)
	}
	return nil
}

func sendIncrement(ctx context.Context, s *ethtx.Session, game *cgcontracts.CosmicSignatureGame, newIncrementMicros *big.Int) error {
	s.Out.TxSubmitting("SetMainPrizeTimeIncrementInMicroSeconds", nil, ethtx.GasLimitAdminCall, s.AdjustedGasPrice())
	tx, err := game.SetMainPrizeTimeIncrementInMicroSeconds(
		s.TransactOpts(nil, ethtx.GasLimitAdminCall),
		newIncrementMicros,
	)
	if err = s.FinishTx(ctx, tx, err); err != nil {
		return fmt.Errorf("setMainPrizeTimeIncrementInMicroSeconds: %w", err)
	}
	return nil
}

func ensureInactiveRound(game *cgcontracts.CosmicSignatureGame, blockTime uint64) error {
	roundActivation, err := game.RoundActivationTime(ethtx.CallOpts())
	if err != nil {
		return err
	}
	if blockTime >= roundActivation.Uint64() {
		return fmt.Errorf(
			"round is still active (activation=%s, block=%d); cannot set increment",
			roundActivation.String(),
			blockTime,
		)
	}
	return nil
}

// maybeAdvanceForClaim waits out the prize timer on Hardhat by advancing block
// time; on real networks it fails when the prize is not yet claimable.
func maybeAdvanceForClaim(ctx context.Context, s *ethtx.Session, game *cgcontracts.CosmicSignatureGame, state *roundState) error {
	if state.durationUntilPrize.Int64() <= 0 {
		return nil
	}
	waitSec := state.durationUntilPrize.Int64() + 1
	if s.Net.IsDevChain() {
		s.Out.Section("ADVANCE HARDHAT TIME FOR CLAIM")
		s.Out.KeyValueDuration("Waiting for prize timer", waitSec)
		if err := s.Net.AdvanceDevChainTime(ctx, waitSec); err != nil {
			return fmt.Errorf("advance hardhat time: %w", err)
		}
		durationUntilPrize, err := game.GetDurationUntilMainPrize(ethtx.CallOpts())
		if err != nil {
			return err
		}
		if durationUntilPrize.Int64() > 0 {
			return fmt.Errorf("prize still not claimable after advancing %d seconds (remaining: %d)", waitSec, durationUntilPrize.Int64())
		}
		state.durationUntilPrize = durationUntilPrize
		return nil
	}
	return fmt.Errorf(
		"prize not claimable yet (%d seconds remaining); wait until the timer expires",
		state.durationUntilPrize.Int64(),
	)
}

func runClaimAndSetTimeIncrement(cmd *cobra.Command, verbose bool, args []string) error {
	ctx := cmd.Context()
	gameAddr, err := parseAddress("cosmicgame-addr", args[0])
	if err != nil {
		return err
	}
	timeIncrementSeconds, err := parseInt64("time_increment_seconds", args[1])
	if err != nil || timeIncrementSeconds <= 0 {
		return fmt.Errorf("time_increment_seconds must be a positive integer")
	}
	delaySeconds := int64(300)
	if len(args) == 3 {
		delaySeconds, err = parseInt64("delay_seconds", args[2])
		if err != nil || delaySeconds <= 0 {
			return fmt.Errorf("delay_seconds must be a positive integer")
		}
	}

	newIncrementMicros := new(big.Int).Mul(big.NewInt(timeIncrementSeconds), big.NewInt(1_000_000))

	s, err := newTxSession(cmd, verbose)
	if err != nil {
		return err
	}

	game, err := cgcontracts.NewCosmicSignatureGame(gameAddr, s.Net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate CosmicGame: %w", err)
	}
	s.Out.ContractInfo("CosmicGame Address", gameAddr)

	state, err := readRoundState(game, s.Net.BlockTime, newIncrementMicros, delaySeconds)
	if err != nil {
		return fmt.Errorf("failed to read contract state: %w", err)
	}

	claimOK, claimReason := state.canClaim(s.Acc.Address)

	s.Out.Section("PLAN")
	s.Out.KeyValue("Round", state.roundNum.String())
	s.Out.KeyValue("Round inactive", state.roundInactive)
	s.Out.KeyValue("Bids this round", state.totalBids.String())
	s.Out.KeyValue("Last bidder", state.lastBidder.String())
	s.Out.KeyValueDuration("Duration until prize", state.durationUntilPrize.Int64())
	s.Out.KeyValue("Increment already set", state.incrementAlreadySet)
	s.Out.KeyValue("Delay already set", state.delayAlreadySet)
	s.Out.KeyValue("Can claim now", fmt.Sprintf("%v (%s)", claimOK, claimReason))
	s.Out.KeyValue("Target increment", fmt.Sprintf("%d sec (%s µs)", timeIncrementSeconds, newIncrementMicros))
	s.Out.KeyValueDuration("Target delay", delaySeconds)

	if s.Acc.Address != state.owner {
		s.Out.Section("WARNING")
		s.Out.KeyValue("Note", "PKEY_HEX is not the contract owner; admin setters will fail")
	}

	if state.incrementAlreadySet && state.delayAlreadySet {
		if !s.Out.Verbose {
			fmt.Fprintf(cmd.OutOrStdout(),
				"Success. Already configured (increment=%d sec, delay=%d sec, round=%s).\n",
				timeIncrementSeconds,
				delaySeconds,
				state.roundNum.String(),
			)
		} else {
			s.Out.Section("SUMMARY")
			s.Out.KeyValue("Status", "Nothing to do — increment and delay already match")
		}
		return nil
	}

	if s.Acc.Address != state.owner {
		return fmt.Errorf("PKEY_HEX must be the contract owner to set time increment")
	}

	if !state.delayAlreadySet {
		s.Out.Section("SET DELAY BEFORE NEXT ROUND")
		if err := sendDelay(ctx, s, game, delaySeconds); err != nil {
			return err
		}
		if err := s.Refresh(ctx); err != nil {
			return fmt.Errorf("network refresh failed after delay tx: %w", err)
		}
		state, err = readRoundState(game, s.Net.BlockTime, newIncrementMicros, delaySeconds)
		if err != nil {
			return fmt.Errorf("failed to re-read contract state: %w", err)
		}
		claimOK, claimReason = state.canClaim(s.Acc.Address)
	} else if s.Out.Verbose {
		s.Out.Section("SKIP")
		s.Out.KeyValue("SetDelayDurationBeforeRoundActivation", "already set")
	}

	// Reaching this with the increment matching means the delay was just
	// updated (a matching increment plus a matching delay exits above), so
	// quiet mode already printed the delay transaction's success line.
	if state.incrementAlreadySet {
		if s.Out.Verbose {
			s.Out.Section("SUMMARY")
			s.Out.KeyValue("Status", "Delay updated; increment already matches target")
		}
		return nil
	}

	if state.roundInactive {
		s.Out.Section("SET TIME INCREMENT (INACTIVE ROUND)")
		if err := s.Refresh(ctx); err != nil {
			return fmt.Errorf("network refresh failed: %w", err)
		}
		if err := ensureInactiveRound(game, s.Net.BlockTime); err != nil {
			return err
		}
		if err := sendIncrement(ctx, s, game, newIncrementMicros); err != nil {
			return err
		}
	} else if claimOK {
		s.Out.Section("CLAIM MAIN PRIZE")
		if err := maybeAdvanceForClaim(ctx, s, game, state); err != nil {
			return err
		}
		if err := s.Refresh(ctx); err != nil {
			return fmt.Errorf("network refresh failed before claim: %w", err)
		}
		if err := sendClaim(ctx, s, game); err != nil {
			return err
		}
		if err := s.Refresh(ctx); err != nil {
			return fmt.Errorf("network refresh failed after claim: %w", err)
		}
		if err := ensureInactiveRound(game, s.Net.BlockTime); err != nil {
			return err
		}
		s.Out.Section("SET TIME INCREMENT (POST-CLAIM INACTIVE WINDOW)")
		if err := sendIncrement(ctx, s, game, newIncrementMicros); err != nil {
			return err
		}
	} else {
		if s.Out.Verbose {
			s.Out.Section("DEFER ROUND ACTIVATION (NO CLAIMABLE PRIZE)")
			s.Out.KeyValue("Reason", claimReason)
		}
		if err := openInactiveWindow(ctx, s, game, delaySeconds); err != nil {
			return err
		}
		s.Out.Section("SET TIME INCREMENT (DEFERRED INACTIVE WINDOW)")
		if err := sendIncrement(ctx, s, game, newIncrementMicros); err != nil {
			return err
		}
	}

	s.Out.Section("SUMMARY")
	s.Out.KeyValue("Status", "Time increment updated")
	s.Out.KeyValueDuration("Per-bid extension", timeIncrementSeconds)
	s.Out.KeyValueDuration("Round delay setting", delaySeconds)
	if claimOK {
		s.Out.KeyValue("Note", "Prize was claimed; next round activates after the delay")
	} else {
		s.Out.KeyValue("Note", "Round activation deferred; bidding resumes after the delay")
	}
	return nil
}
