package cosmicgame

import (
	"context"
	"errors"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

const (
	liveStateEnduranceChampionDuration = "endurance_champion_duration"
	liveStateChronoWarriorDuration     = "chrono_warrior_duration"
)

func (h *Handlers) readChampionDurations(ctx context.Context, roundNum int64) (endurance, chrono int64, err error) {
	caller, _ := cgc.NewCosmicSignatureGameV3Caller(h.c.Game, h.caller)
	out, err := caller.ChampionDurations(&bind.CallOpts{Context: ctx}, big.NewInt(roundNum))
	if err != nil {
		return 0, 0, err
	}
	if out.EnduranceChampion == nil || out.ChronoWarrior == nil ||
		!out.EnduranceChampion.IsInt64() || !out.ChronoWarrior.IsInt64() ||
		out.EnduranceChampion.Sign() < 0 || out.ChronoWarrior.Sign() < 0 {
		return 0, 0, errors.New("championDurations values exceed int64")
	}
	return out.EnduranceChampion.Int64(), out.ChronoWarrior.Int64(), nil
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
