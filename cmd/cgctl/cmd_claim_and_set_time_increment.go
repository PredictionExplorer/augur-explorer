package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/cmd/cgctl/internal/ethtx"
	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func init() {
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

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required; owner for admin calls;
            last bidder when claiming)`,
		Args: cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runClaimAndSetTimeIncrement(cmd.Context(), ethtx.NewPrinter(info), args)
		},
	}
	c.Flags().BoolVarP(&info, "info", "i", false, "print detailed output")
	register(c)
}

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

// refreshNetwork re-reads gas price, block info, and the account nonce.
func refreshNetwork(ctx context.Context, pkeyHex string) (*ethtx.NetworkInfo, *ethtx.AccountInfo, error) {
	net, err := ethtx.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}
	acc, err := net.PrepareAccount(ctx, pkeyHex)
	if err != nil {
		return nil, nil, err
	}
	return net, acc, nil
}

func sendDelay(
	ctx context.Context,
	out *ethtx.Printer,
	game *cgcontracts.CosmicSignatureGame,
	net *ethtx.NetworkInfo,
	acc *ethtx.AccountInfo,
	delaySeconds int64,
) error {
	out.TxSubmitting("SetDelayDurationBeforeRoundActivation", nil, ethtx.GasLimitAdminCall, net.GasPrice)
	tx, err := game.SetDelayDurationBeforeRoundActivation(
		net.TransactOpts(acc, nil, ethtx.GasLimitAdminCall),
		big.NewInt(delaySeconds),
	)
	if err != nil {
		return fmt.Errorf("setDelayDurationBeforeRoundActivation: %w", err)
	}
	return out.TxMined(ctx, net.Client, tx)
}

func sendDeferActivation(
	ctx context.Context,
	out *ethtx.Printer,
	game *cgcontracts.CosmicSignatureGame,
	net *ethtx.NetworkInfo,
	acc *ethtx.AccountInfo,
	delaySeconds int64,
) error {
	newActivation := big.NewInt(int64(net.BlockTime) + delaySeconds + 5)
	out.TxSubmitting("SetRoundActivationTime", nil, ethtx.GasLimitAdminCall, net.GasPrice)
	out.KeyValue("New activation time", newActivation.String())
	tx, err := game.SetRoundActivationTime(
		net.TransactOpts(acc, nil, ethtx.GasLimitAdminCall),
		newActivation,
	)
	if err != nil {
		return fmt.Errorf("setRoundActivationTime: %w", err)
	}
	return out.TxMined(ctx, net.Client, tx)
}

// openInactiveWindow defers round activation until the round reads as
// inactive, retrying a few times to absorb block-time races.
func openInactiveWindow(
	ctx context.Context,
	out *ethtx.Printer,
	game *cgcontracts.CosmicSignatureGame,
	pkeyHex string,
	delaySeconds int64,
) (*ethtx.NetworkInfo, *ethtx.AccountInfo, error) {
	const maxAttempts = 3
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		net, acc, err := refreshNetwork(ctx, pkeyHex)
		if err != nil {
			return nil, nil, err
		}
		state, err := readRoundState(game, net.BlockTime, big.NewInt(0), delaySeconds)
		if err != nil {
			return nil, nil, err
		}
		if state.roundInactive {
			return net, acc, nil
		}
		if attempt > 1 || out.Verbose {
			out.Section(fmt.Sprintf("DEFER ROUND ACTIVATION (attempt %d/%d)", attempt, maxAttempts))
		}
		if err := sendDeferActivation(ctx, out, game, net, acc, delaySeconds); err != nil {
			return nil, nil, fmt.Errorf("setRoundActivationTime failed on attempt %d: %w", attempt, err)
		}
	}
	net, acc, err := refreshNetwork(ctx, pkeyHex)
	if err != nil {
		return nil, nil, err
	}
	state, err := readRoundState(game, net.BlockTime, big.NewInt(0), delaySeconds)
	if err != nil {
		return nil, nil, err
	}
	if !state.roundInactive {
		return nil, nil, fmt.Errorf(
			"round is still active after %d defer attempts (activation=%s, block=%d)",
			maxAttempts,
			state.roundActivation.String(),
			net.BlockTime,
		)
	}
	return net, acc, nil
}

func sendClaim(
	ctx context.Context,
	out *ethtx.Printer,
	game *cgcontracts.CosmicSignatureGame,
	net *ethtx.NetworkInfo,
	acc *ethtx.AccountInfo,
) error {
	out.TxSubmitting("ClaimMainPrize", nil, ethtx.GasLimitClaimPrize, net.GasPrice)
	tx, err := game.ClaimMainPrize(net.TransactOpts(acc, nil, ethtx.GasLimitClaimPrize))
	if err != nil {
		return fmt.Errorf("claimMainPrize: %w", err)
	}
	return out.TxMined(ctx, net.Client, tx)
}

func sendIncrement(
	ctx context.Context,
	out *ethtx.Printer,
	game *cgcontracts.CosmicSignatureGame,
	net *ethtx.NetworkInfo,
	acc *ethtx.AccountInfo,
	newIncrementMicros *big.Int,
) error {
	out.TxSubmitting("SetMainPrizeTimeIncrementInMicroSeconds", nil, ethtx.GasLimitAdminCall, net.GasPrice)
	tx, err := game.SetMainPrizeTimeIncrementInMicroSeconds(
		net.TransactOpts(acc, nil, ethtx.GasLimitAdminCall),
		newIncrementMicros,
	)
	if err != nil {
		return fmt.Errorf("setMainPrizeTimeIncrementInMicroSeconds: %w", err)
	}
	return out.TxMined(ctx, net.Client, tx)
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
func maybeAdvanceForClaim(
	ctx context.Context,
	out *ethtx.Printer,
	game *cgcontracts.CosmicSignatureGame,
	net *ethtx.NetworkInfo,
	state *roundState,
) error {
	if state.durationUntilPrize.Int64() <= 0 {
		return nil
	}
	waitSec := state.durationUntilPrize.Int64() + 1
	if net.ChainID.Int64() == 31337 {
		out.Section("ADVANCE HARDHAT TIME FOR CLAIM")
		out.KeyValueDuration("Waiting for prize timer", waitSec)
		if err := net.AdvanceHardhatTime(ctx, waitSec); err != nil {
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

func runClaimAndSetTimeIncrement(ctx context.Context, out *ethtx.Printer, args []string) error {
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

	pkeyHex, err := ethtx.PrivateKeyHexFromEnv()
	if err != nil {
		return err
	}
	newIncrementMicros := new(big.Int).Mul(big.NewInt(timeIncrementSeconds), big.NewInt(1_000_000))

	net, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)

	acc, err := net.PrepareAccount(ctx, pkeyHex)
	if err != nil {
		return fmt.Errorf("account setup failed: %w", err)
	}
	out.AccountInfo(acc)

	game, err := cgcontracts.NewCosmicSignatureGame(gameAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate CosmicGame: %w", err)
	}
	out.ContractInfo("CosmicGame Address", gameAddr)

	state, err := readRoundState(game, net.BlockTime, newIncrementMicros, delaySeconds)
	if err != nil {
		return fmt.Errorf("failed to read contract state: %w", err)
	}

	claimOK, claimReason := state.canClaim(acc.Address)

	out.Section("PLAN")
	out.KeyValue("Round", state.roundNum.String())
	out.KeyValue("Round inactive", state.roundInactive)
	out.KeyValue("Bids this round", state.totalBids.String())
	out.KeyValue("Last bidder", state.lastBidder.String())
	out.KeyValueDuration("Duration until prize", state.durationUntilPrize.Int64())
	out.KeyValue("Increment already set", state.incrementAlreadySet)
	out.KeyValue("Delay already set", state.delayAlreadySet)
	out.KeyValue("Can claim now", fmt.Sprintf("%v (%s)", claimOK, claimReason))
	out.KeyValue("Target increment", fmt.Sprintf("%d sec (%s µs)", timeIncrementSeconds, newIncrementMicros))
	out.KeyValueDuration("Target delay", delaySeconds)

	if acc.Address != state.owner {
		out.Section("WARNING")
		out.KeyValue("Note", "PKEY_HEX is not the contract owner; admin setters will fail")
	}

	if state.incrementAlreadySet && state.delayAlreadySet {
		if !out.Verbose {
			fmt.Printf(
				"Success. Already configured (increment=%d sec, delay=%d sec, round=%s).\n",
				timeIncrementSeconds,
				delaySeconds,
				state.roundNum.String(),
			)
		} else {
			out.Section("SUMMARY")
			out.KeyValue("Status", "Nothing to do — increment and delay already match")
		}
		return nil
	}

	if acc.Address != state.owner {
		return fmt.Errorf("PKEY_HEX must be the contract owner to set time increment")
	}

	txCount := 0

	if !state.delayAlreadySet {
		out.Section("SET DELAY BEFORE NEXT ROUND")
		if err := sendDelay(ctx, out, game, net, acc, delaySeconds); err != nil {
			return err
		}
		txCount++
		net, acc, err = refreshNetwork(ctx, pkeyHex)
		if err != nil {
			return fmt.Errorf("network refresh failed after delay tx: %w", err)
		}
		state, err = readRoundState(game, net.BlockTime, newIncrementMicros, delaySeconds)
		if err != nil {
			return fmt.Errorf("failed to re-read contract state: %w", err)
		}
		claimOK, claimReason = state.canClaim(acc.Address)
	} else if out.Verbose {
		out.Section("SKIP")
		out.KeyValue("SetDelayDurationBeforeRoundActivation", "already set")
	}

	if state.incrementAlreadySet {
		if !out.Verbose {
			if txCount == 0 {
				fmt.Printf(
					"Success. Already configured (increment=%d sec, delay=%d sec, round=%s).\n",
					timeIncrementSeconds,
					delaySeconds,
					state.roundNum.String(),
				)
			}
		} else {
			out.Section("SUMMARY")
			out.KeyValue("Status", "Delay updated; increment already matches target")
		}
		return nil
	}

	if state.roundInactive {
		out.Section("SET TIME INCREMENT (INACTIVE ROUND)")
		net, acc, err = refreshNetwork(ctx, pkeyHex)
		if err != nil {
			return fmt.Errorf("network refresh failed: %w", err)
		}
		if err := ensureInactiveRound(game, net.BlockTime); err != nil {
			return err
		}
		if err := sendIncrement(ctx, out, game, net, acc, newIncrementMicros); err != nil {
			return err
		}
	} else if claimOK {
		out.Section("CLAIM MAIN PRIZE")
		if err := maybeAdvanceForClaim(ctx, out, game, net, state); err != nil {
			return err
		}
		net, acc, err = refreshNetwork(ctx, pkeyHex)
		if err != nil {
			return fmt.Errorf("network refresh failed before claim: %w", err)
		}
		if err := sendClaim(ctx, out, game, net, acc); err != nil {
			return err
		}
		txCount++
		net, acc, err = refreshNetwork(ctx, pkeyHex)
		if err != nil {
			return fmt.Errorf("network refresh failed after claim: %w", err)
		}
		if err := ensureInactiveRound(game, net.BlockTime); err != nil {
			return err
		}
		out.Section("SET TIME INCREMENT (POST-CLAIM INACTIVE WINDOW)")
		if err := sendIncrement(ctx, out, game, net, acc, newIncrementMicros); err != nil {
			return err
		}
	} else {
		if out.Verbose {
			out.Section("DEFER ROUND ACTIVATION (NO CLAIMABLE PRIZE)")
			out.KeyValue("Reason", claimReason)
		}
		net, acc, err = openInactiveWindow(ctx, out, game, pkeyHex, delaySeconds)
		if err != nil {
			return err
		}
		txCount++
		out.Section("SET TIME INCREMENT (DEFERRED INACTIVE WINDOW)")
		if err := sendIncrement(ctx, out, game, net, acc, newIncrementMicros); err != nil {
			return err
		}
	}

	out.Section("SUMMARY")
	out.KeyValue("Status", "Time increment updated")
	out.KeyValueDuration("Per-bid extension", timeIncrementSeconds)
	out.KeyValueDuration("Round delay setting", delaySeconds)
	if claimOK {
		out.KeyValue("Note", "Prize was claimed; next round activates after the delay")
	} else {
		out.KeyValue("Note", "Round activation deferred; bidding resumes after the delay")
	}
	return nil
}
