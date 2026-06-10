// Sets mainPrizeTimeIncrementInMicroSeconds, claiming the main prize first when needed.
//
// Contract rules (V1/V2 proxy):
//   - setMainPrizeTimeIncrementInMicroSeconds requires an inactive round (block.timestamp < roundActivationTime)
//   - claimMainPrize requires bids in the current round (lastBidderAddress != 0) and the prize timer or timeout
//   - setDelayDurationBeforeRoundActivation and setRoundActivationTime can run while the round is active
//
// This script picks a path from on-chain state:
//   A) increment and delay already match → exit success (idempotent)
//   B) round already inactive → set increment directly
//   C) round active, prize claimable → claim (creates inactive window) → set increment
//   D) round active, prize not claimable → defer activation (setRoundActivationTime) → set increment
package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

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
	cosmicGame *CosmicSignatureGame,
	blockTime uint64,
	newIncrementMicros *big.Int,
	delaySeconds int64,
) (*roundState, error) {
	copts := cutils.CreateCallOpts()

	roundNum, err := cosmicGame.RoundNum(copts)
	if err != nil {
		return nil, fmt.Errorf("roundNum: %w", err)
	}
	roundActivation, err := cosmicGame.RoundActivationTime(copts)
	if err != nil {
		return nil, fmt.Errorf("roundActivationTime: %w", err)
	}
	totalBids, err := cosmicGame.GetTotalNumBids(copts, roundNum)
	if err != nil {
		return nil, fmt.Errorf("getTotalNumBids: %w", err)
	}
	lastBidder, err := cosmicGame.LastBidderAddress(copts)
	if err != nil {
		return nil, fmt.Errorf("lastBidderAddress: %w", err)
	}
	durationUntilPrize, err := cosmicGame.GetDurationUntilMainPrize(copts)
	if err != nil {
		return nil, fmt.Errorf("getDurationUntilMainPrize: %w", err)
	}
	durationUntilPrizeRaw, err := cosmicGame.GetDurationUntilMainPrizeRaw(copts)
	if err != nil {
		return nil, fmt.Errorf("getDurationUntilMainPrizeRaw: %w", err)
	}
	timeoutClaim, err := cosmicGame.TimeoutDurationToClaimMainPrize(copts)
	if err != nil {
		return nil, fmt.Errorf("timeoutDurationToClaimMainPrize: %w", err)
	}
	currentIncrementMicros, err := cosmicGame.MainPrizeTimeIncrementInMicroSeconds(copts)
	if err != nil {
		return nil, fmt.Errorf("mainPrizeTimeIncrementInMicroSeconds: %w", err)
	}
	currentDelay, err := cosmicGame.DelayDurationBeforeRoundActivation(copts)
	if err != nil {
		return nil, fmt.Errorf("delayDurationBeforeRoundActivation: %w", err)
	}
	owner, err := cosmicGame.Owner(copts)
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

func refreshNetwork(pkeyHex string) (*cutils.NetworkInfo, *cutils.AccountInfo, error) {
	net, err := cutils.ConnectToRPC()
	if err != nil {
		return nil, nil, err
	}
	acc, err := cutils.PrepareAccount(net, pkeyHex)
	if err != nil {
		return nil, nil, err
	}
	return net, acc, nil
}

func sendDelay(
	cosmicGame *CosmicSignatureGame,
	net *cutils.NetworkInfo,
	acc *cutils.AccountInfo,
	delaySeconds int64,
) bool {
	cutils.PrintTxSubmitting("SetDelayDurationBeforeRoundActivation", nil, cutils.GasLimitAdminCall, net.GasPrice)
	tx, err := cosmicGame.SetDelayDurationBeforeRoundActivation(
		cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall),
		big.NewInt(delaySeconds),
	)
	return cutils.PrintTxResultAndWait(net.Client, tx, err)
}

func sendDeferActivation(
	cosmicGame *CosmicSignatureGame,
	net *cutils.NetworkInfo,
	acc *cutils.AccountInfo,
	delaySeconds int64,
) bool {
	newActivation := big.NewInt(int64(net.BlockTime) + delaySeconds + 5)
	cutils.PrintTxSubmitting("SetRoundActivationTime", nil, cutils.GasLimitAdminCall, net.GasPrice)
	cutils.PrintKeyValue("New activation time", newActivation.String())
	tx, err := cosmicGame.SetRoundActivationTime(
		cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall),
		newActivation,
	)
	return cutils.PrintTxResultAndWait(net.Client, tx, err)
}

func openInactiveWindow(
	cosmicGame *CosmicSignatureGame,
	pkeyHex string,
	delaySeconds int64,
) (*cutils.NetworkInfo, *cutils.AccountInfo, error) {
	const maxAttempts = 3
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		net, acc, err := refreshNetwork(pkeyHex)
		if err != nil {
			return nil, nil, err
		}
		state, err := readRoundState(cosmicGame, net.BlockTime, big.NewInt(0), delaySeconds)
		if err != nil {
			return nil, nil, err
		}
		if state.roundInactive {
			return net, acc, nil
		}
		if attempt > 1 || cutils.Verbose {
			cutils.Section(fmt.Sprintf("DEFER ROUND ACTIVATION (attempt %d/%d)", attempt, maxAttempts))
		}
		if !sendDeferActivation(cosmicGame, net, acc, delaySeconds) {
			return nil, nil, fmt.Errorf("setRoundActivationTime failed on attempt %d", attempt)
		}
	}
	net, acc, err := refreshNetwork(pkeyHex)
	if err != nil {
		return nil, nil, err
	}
	state, err := readRoundState(cosmicGame, net.BlockTime, big.NewInt(0), delaySeconds)
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
	cosmicGame *CosmicSignatureGame,
	net *cutils.NetworkInfo,
	acc *cutils.AccountInfo,
) bool {
	cutils.PrintTxSubmitting("ClaimMainPrize", nil, cutils.GasLimitClaimPrize, net.GasPrice)
	tx, err := cosmicGame.ClaimMainPrize(cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitClaimPrize))
	return cutils.PrintTxResultAndWait(net.Client, tx, err)
}

func sendIncrement(
	cosmicGame *CosmicSignatureGame,
	net *cutils.NetworkInfo,
	acc *cutils.AccountInfo,
	newIncrementMicros *big.Int,
) bool {
	cutils.PrintTxSubmitting("SetMainPrizeTimeIncrementInMicroSeconds", nil, cutils.GasLimitAdminCall, net.GasPrice)
	tx, err := cosmicGame.SetMainPrizeTimeIncrementInMicroSeconds(
		cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall),
		newIncrementMicros,
	)
	return cutils.PrintTxResultAndWait(net.Client, tx, err)
}

func ensureInactiveRound(cosmicGame *CosmicSignatureGame, client *ethclient.Client, blockTime uint64) error {
	roundActivation, err := cosmicGame.RoundActivationTime(cutils.CreateCallOpts())
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

func maybeAdvanceForClaim(
	cosmicGame *CosmicSignatureGame,
	net *cutils.NetworkInfo,
	state *roundState,
) error {
	if state.durationUntilPrize.Int64() <= 0 {
		return nil
	}
	waitSec := state.durationUntilPrize.Int64() + 1
	if net.ChainID.Int64() == 31337 {
		cutils.Section("ADVANCE HARDHAT TIME FOR CLAIM")
		cutils.PrintKeyValueDuration("Waiting for prize timer", waitSec)
		if err := cutils.AdvanceHardhatTime(net, waitSec); err != nil {
			return fmt.Errorf("advance hardhat time: %w", err)
		}
		durationUntilPrize, err := cosmicGame.GetDurationUntilMainPrize(cutils.CreateCallOpts())
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

func main() {
	cutils.ParseInfoFlag()

	if len(os.Args) < 3 || len(os.Args) > 4 {
		cutils.PrintUsage(os.Args[0],
			"[cosmicgame_contract_addr] [time_increment_seconds] [delay_seconds]",
			"Sets per-bid time increment. Claims the prize first when that is required to open the inactive admin window. delay_seconds defaults to 300",
			map[string]string{
				"RPC_URL":  "Ethereum RPC endpoint (required)",
				"PKEY_HEX": "64-char hex private key, no 0x prefix (required; owner for admin calls; last bidder when claiming)",
			},
		)
		os.Exit(1)
	}

	timeIncrementSeconds, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil || timeIncrementSeconds <= 0 {
		cutils.Fatal("time_increment_seconds must be a positive integer")
	}

	delaySeconds := int64(300)
	if len(os.Args) == 4 {
		delaySeconds, err = strconv.ParseInt(os.Args[3], 10, 64)
		if err != nil || delaySeconds <= 0 {
			cutils.Fatal("delay_seconds must be a positive integer")
		}
	}

	pkeyHex := cutils.MustGetPkeyHex()
	cosmicGameAddr := common.HexToAddress(os.Args[1])
	newIncrementMicros := new(big.Int).Mul(big.NewInt(timeIncrementSeconds), big.NewInt(1_000_000))

	net, err := cutils.ConnectToRPC()
	if err != nil {
		cutils.Fatal("Network connection failed: %v", err)
	}
	cutils.PrintNetworkInfo(net)

	acc, err := cutils.PrepareAccount(net, pkeyHex)
	if err != nil {
		cutils.Fatal("Account setup failed: %v", err)
	}
	cutils.PrintAccountInfo(acc)

	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame: %v", err)
	}
	cutils.PrintContractInfo("CosmicGame Address", cosmicGameAddr)

	state, err := readRoundState(cosmicGame, net.BlockTime, newIncrementMicros, delaySeconds)
	if err != nil {
		cutils.Fatal("Failed to read contract state: %v", err)
	}

	claimOK, claimReason := state.canClaim(acc.Address)

	cutils.Section("PLAN")
	cutils.PrintKeyValue("Round", state.roundNum.String())
	cutils.PrintKeyValue("Round inactive", state.roundInactive)
	cutils.PrintKeyValue("Bids this round", state.totalBids.String())
	cutils.PrintKeyValue("Last bidder", state.lastBidder.String())
	cutils.PrintKeyValueDuration("Duration until prize", state.durationUntilPrize.Int64())
	cutils.PrintKeyValue("Increment already set", state.incrementAlreadySet)
	cutils.PrintKeyValue("Delay already set", state.delayAlreadySet)
	cutils.PrintKeyValue("Can claim now", fmt.Sprintf("%v (%s)", claimOK, claimReason))
	cutils.PrintKeyValue("Target increment", fmt.Sprintf("%d sec (%s µs)", timeIncrementSeconds, newIncrementMicros))
	cutils.PrintKeyValueDuration("Target delay", delaySeconds)

	if acc.Address != state.owner {
		cutils.Section("WARNING")
		cutils.PrintKeyValue("Note", "PKEY_HEX is not the contract owner; admin setters will fail")
	}

	if state.incrementAlreadySet && state.delayAlreadySet {
		if !cutils.Verbose {
			fmt.Printf(
				"Success. Already configured (increment=%d sec, delay=%d sec, round=%s).\n",
				timeIncrementSeconds,
				delaySeconds,
				state.roundNum.String(),
			)
		} else {
			cutils.Section("SUMMARY")
			cutils.PrintKeyValue("Status", "Nothing to do — increment and delay already match")
		}
		return
	}

	if acc.Address != state.owner {
		cutils.Fatal("PKEY_HEX must be the contract owner to set time increment")
	}

	txCount := 0

	if !state.delayAlreadySet {
		cutils.Section("SET DELAY BEFORE NEXT ROUND")
		if !sendDelay(cosmicGame, net, acc, delaySeconds) {
			os.Exit(1)
		}
		txCount++
		net, acc, err = refreshNetwork(pkeyHex)
		cutils.FatalIf(err, "Network refresh failed after delay tx")
		state, err = readRoundState(cosmicGame, net.BlockTime, newIncrementMicros, delaySeconds)
		cutils.FatalIf(err, "Failed to re-read contract state")
		claimOK, claimReason = state.canClaim(acc.Address)
	} else if cutils.Verbose {
		cutils.Section("SKIP")
		cutils.PrintKeyValue("SetDelayDurationBeforeRoundActivation", "already set")
	}

	if state.incrementAlreadySet {
		if !cutils.Verbose {
			if txCount == 0 {
				fmt.Printf(
					"Success. Already configured (increment=%d sec, delay=%d sec, round=%s).\n",
					timeIncrementSeconds,
					delaySeconds,
					state.roundNum.String(),
				)
			}
		} else {
			cutils.Section("SUMMARY")
			cutils.PrintKeyValue("Status", "Delay updated; increment already matches target")
		}
		return
	}

	if state.roundInactive {
		cutils.Section("SET TIME INCREMENT (INACTIVE ROUND)")
		net, acc, err = refreshNetwork(pkeyHex)
		cutils.FatalIf(err, "Network refresh failed")
		if err := ensureInactiveRound(cosmicGame, net.Client, net.BlockTime); err != nil {
			cutils.Fatal("%v", err)
		}
		if !sendIncrement(cosmicGame, net, acc, newIncrementMicros) {
			os.Exit(1)
		}
	} else if claimOK {
		cutils.Section("CLAIM MAIN PRIZE")
		if err := maybeAdvanceForClaim(cosmicGame, net, state); err != nil {
			cutils.Fatal("%v", err)
		}
		net, acc, err = refreshNetwork(pkeyHex)
		cutils.FatalIf(err, "Network refresh failed before claim")
		if !sendClaim(cosmicGame, net, acc) {
			os.Exit(1)
		}
		txCount++
		net, acc, err = refreshNetwork(pkeyHex)
		cutils.FatalIf(err, "Network refresh failed after claim")
		if err := ensureInactiveRound(cosmicGame, net.Client, net.BlockTime); err != nil {
			cutils.Fatal("%v", err)
		}
		cutils.Section("SET TIME INCREMENT (POST-CLAIM INACTIVE WINDOW)")
		if !sendIncrement(cosmicGame, net, acc, newIncrementMicros) {
			os.Exit(1)
		}
	} else {
		if cutils.Verbose {
			cutils.Section("DEFER ROUND ACTIVATION (NO CLAIMABLE PRIZE)")
			cutils.PrintKeyValue("Reason", claimReason)
		}
		var err error
		net, acc, err = openInactiveWindow(cosmicGame, pkeyHex, delaySeconds)
		if err != nil {
			cutils.Fatal("%v", err)
		}
		txCount++
		cutils.Section("SET TIME INCREMENT (DEFERRED INACTIVE WINDOW)")
		if !sendIncrement(cosmicGame, net, acc, newIncrementMicros) {
			os.Exit(1)
		}
	}

	cutils.Section("SUMMARY")
	cutils.PrintKeyValue("Status", "Time increment updated")
	cutils.PrintKeyValueDuration("Per-bid extension", timeIncrementSeconds)
	cutils.PrintKeyValueDuration("Round delay setting", delaySeconds)
	if claimOK {
		cutils.PrintKeyValue("Note", "Prize was claimed; next round activates after the delay")
	} else {
		cutils.PrintKeyValue("Note", "Round activation deferred; bidding resumes after the delay")
	}
}
