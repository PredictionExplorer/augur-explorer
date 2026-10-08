package cosmicgame

import (
	"context"
	"errors"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

const (
	liveStateEnduranceChampionDuration = "endurance_champion_duration"
	liveStateChronoWarriorDuration     = "chrono_warrior_duration"
)

// legacyChampionDurationsABI is the pre-consolidation v3.1 getter that
// roundStats() replaced. Kept as a raw fragment (the current bindings no
// longer carry it) so recovery still works against chains deployed from
// older v3.1 code.
var legacyChampionDurationsABI = func() abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"uint256","name":"roundNum_","type":"uint256"}],"name":"championDurations","outputs":[{"internalType":"uint256","name":"enduranceChampion","type":"uint256"},{"internalType":"uint256","name":"chronoWarrior","type":"uint256"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		panic(err)
	}
	return parsed
}()

// checkedChampionPair validates the raw duration pair read from the contract.
func checkedChampionPair(enduranceRaw, chronoRaw *big.Int) (endurance, chrono int64, err error) {
	if enduranceRaw == nil || chronoRaw == nil ||
		!enduranceRaw.IsInt64() || !chronoRaw.IsInt64() ||
		enduranceRaw.Sign() < 0 || chronoRaw.Sign() < 0 {
		return 0, 0, errors.New("champion duration values exceed int64")
	}
	return enduranceRaw.Int64(), chronoRaw.Int64(), nil
}

func (h *Handlers) readChampionDurations(ctx context.Context, roundNum int64) (endurance, chrono int64, err error) {
	copts := &bind.CallOpts{Context: ctx}
	rn := big.NewInt(roundNum)

	// Current v3.1 contracts consolidated the per-round getters into
	// roundStats(); its 7th and 8th fields are the champion durations.
	caller, _ := cgc.NewCosmicSignatureGameV3Caller(h.c.Game, h.caller)
	if stats, statsErr := caller.RoundStats(copts, rn); statsErr == nil {
		return checkedChampionPair(stats.EnduranceChampionDuration, stats.ChronoWarriorDuration)
	}

	// Older v3.1 deployments expose championDurations(uint256) instead.
	legacy := bind.NewBoundContract(h.c.Game, legacyChampionDurationsABI, h.caller, nil, nil)
	var out []interface{}
	if err := legacy.Call(copts, &out, "championDurations", rn); err != nil {
		return 0, 0, err
	}
	if len(out) != 2 {
		return 0, 0, errors.New("championDurations returned unexpected shape")
	}
	enduranceRaw, _ := out[0].(*big.Int)
	chronoRaw, _ := out[1].(*big.Int)
	return checkedChampionPair(enduranceRaw, chronoRaw)
}

// captureChampionDurations reads and stores one V3 snapshot. Contract-call
// failure is deliberately non-fatal so MainPrizeClaimed remains indexed; a
// startup recovery pass retries missing rounds. Store failures still abort
// the surrounding block transaction.
func (h *Handlers) captureChampionDurations(
	ctx context.Context,
	roundNum, observedAtBlock, observedAtTime int64,
) (bool, error) {
	endurance, chrono, err := h.readChampionDurations(ctx, roundNum)
	if err != nil {
		h.log.Warn("champion durations read failed; startup recovery will retry",
			"round", roundNum, "err", err)
		return false, nil
	}
	if endurance == 0 && chrono == 0 {
		return true, nil
	}
	if err := h.repo.UpdateRoundChampionDurations(ctx, roundNum, endurance, chrono); err != nil {
		return true, err
	}
	for _, update := range []struct {
		name  string
		value int64
	}{
		{liveStateEnduranceChampionDuration, endurance},
		{liveStateChronoWarriorDuration, chrono},
	} {
		inserted, err := h.repo.InsertLiveStateUpdateIfChanged(
			ctx,
			update.name,
			h.c.GameAid,
			roundNum,
			observedAtBlock,
			observedAtTime,
			strconv.FormatInt(update.value, 10),
		)
		if err != nil {
			return true, err
		}
		if inserted {
			h.log.Info("live state update recorded",
				"variable", update.name,
				"round", roundNum,
				"value", update.value,
				"block", observedAtBlock)
		}
	}
	return true, nil
}

// RecoverChampionDurations backfills every claimed round whose duration pair
// is still zero. A selector/RPC failure stops the pass without failing ETL
// startup (V1/V2 contracts legitimately do not implement the getter).
func (h *Handlers) RecoverChampionDurations(ctx context.Context, observedAtBlock, observedAtTime int64) error {
	rounds, err := h.repo.RoundsMissingChampionDurations(ctx)
	if err != nil {
		return err
	}
	for _, roundNum := range rounds {
		read, err := h.captureChampionDurations(ctx, roundNum, observedAtBlock, observedAtTime)
		if err != nil {
			return err
		}
		if !read {
			h.log.Info("champion duration recovery stopped",
				"round", roundNum,
				"reason", "contract is pre-V3 or RPC read failed")
			return nil
		}
	}
	return nil
}
