package contractstate

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	cg "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// LiveSpecialWinners is the on-demand contract read behind the
// bid/current_special_winners route: the live endurance-champion and
// chrono-warrior standings at the latest block.
type LiveSpecialWinners struct {
	EnduranceChampionAddress        string
	EnduranceChampionDuration       int64
	EnduranceChampionStartTimeStamp int64
	PrevEnduranceChampionDuration   int64
	ChronoWarriorAddress            string
	ChronoWarriorDuration           int64
	ChronoWarriorIsLive             bool
	LastBidderAddress               string
	LastBidderLastBidTime           int64
	LastCstBidderAddress            string
	LastCstBidderLastBidTime        int64
	LastCstBidEventLogId            int64
	HasLastCstBidderLastBidTime     bool
	HasLastCstBidEventLogId         bool
	RoundNum                        int64
	SourceBlockNumber               uint64
	SourceBlockTimeStamp            int64
	Err                             error
}

// FetchLiveSpecialWinners performs the live special-winners reads against
// the latest block. Failures are reported in the Err field rather than a
// return value so the handler keeps its legacy single-error response shape.
func (s *State) FetchLiveSpecialWinners(ctx context.Context) LiveSpecialWinners {
	var out LiveSpecialWinners
	out.RoundNum = s.Snapshot().RoundNum

	header, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		out.Err = fmt.Errorf("failed to fetch latest block header: %w", err)
		return out
	}
	out.SourceBlockNumber = header.Number.Uint64()
	out.SourceBlockTimeStamp = int64(header.Time) // #nosec G115 -- block timestamps fit in int64 until year 292277026596

	contract, err := cg.NewCosmicSignatureGame(s.addrs.CosmicGame, s.client)
	if err != nil {
		out.Err = fmt.Errorf("failed to instantiate CosmicGame contract: %w", err)
		return out
	}

	copts := bind.CallOpts{Context: ctx, BlockNumber: header.Number}

	var (
		champs struct {
			EnduranceChampionAddress  ethcommon.Address
			EnduranceChampionDuration *big.Int
			ChronoWarriorAddress      ethcommon.Address
			ChronoWarriorDuration     *big.Int
		}
		enduranceStartTs            *big.Int
		prevEnduranceDuration       *big.Int
		storedEnduranceChampionAddr ethcommon.Address
		storedEnduranceChampionDur  *big.Int
		lastBidder                  ethcommon.Address
		lastCstBidder               ethcommon.Address
		storedChronoWarriorDur      *big.Int
		lastBidderLastBidTime       *big.Int
		lastCstBidderLastBidTime    *big.Int
	)

	var wg sync.WaitGroup
	var mu sync.Mutex
	recordErr := func(label string, err error) {
		if err == nil {
			return
		}
		mu.Lock()
		defer mu.Unlock()
		if out.Err == nil {
			out.Err = fmt.Errorf("%s: %w", label, err)
		}
	}

	wg.Add(8)
	go func() {
		defer wg.Done()
		result, err := contract.TryGetCurrentChampions(&copts)
		if err != nil {
			recordErr("TryGetCurrentChampions", err)
			return
		}
		mu.Lock()
		champs = result
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.EnduranceChampionStartTimeStamp(&copts)
		if err != nil {
			recordErr("EnduranceChampionStartTimeStamp", err)
			return
		}
		mu.Lock()
		enduranceStartTs = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.PrevEnduranceChampionDuration(&copts)
		if err != nil {
			recordErr("PrevEnduranceChampionDuration", err)
			return
		}
		mu.Lock()
		prevEnduranceDuration = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.LastBidderAddress(&copts)
		if err != nil {
			recordErr("LastBidderAddress", err)
			return
		}
		mu.Lock()
		lastBidder = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.LastCstBidderAddress(&copts)
		if err != nil {
			recordErr("LastCstBidderAddress", err)
			return
		}
		mu.Lock()
		lastCstBidder = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.ChronoWarriorDuration(&copts)
		if err != nil {
			recordErr("ChronoWarriorDuration", err)
			return
		}
		mu.Lock()
		storedChronoWarriorDur = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.EnduranceChampionAddress(&copts)
		if err != nil {
			recordErr("EnduranceChampionAddress", err)
			return
		}
		mu.Lock()
		storedEnduranceChampionAddr = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.EnduranceChampionDuration(&copts)
		if err != nil {
			recordErr("EnduranceChampionDuration", err)
			return
		}
		mu.Lock()
		storedEnduranceChampionDur = val
		mu.Unlock()
	}()
	wg.Wait()

	if out.Err != nil {
		return out
	}

	out.EnduranceChampionAddress = champs.EnduranceChampionAddress.String()
	out.EnduranceChampionDuration = champs.EnduranceChampionDuration.Int64()
	out.ChronoWarriorAddress = champs.ChronoWarriorAddress.String()
	out.ChronoWarriorDuration = champs.ChronoWarriorDuration.Int64()
	out.LastBidderAddress = lastBidder.String()
	out.LastCstBidderAddress = lastCstBidder.String()
	// The chrono-segment anchor (EnduranceChampionStartTimeStamp /
	// PrevEnduranceChampionDuration) and ChronoWarriorIsLive are computed
	// below, after the last bidder's lastBidTimeStamp is known, so the anchor
	// stays consistent with the LIVE champion from tryGetCurrentChampions().
	if out.RoundNum >= 0 {
		roundBig := big.NewInt(out.RoundNum)
		wg.Add(2)
		go func() {
			defer wg.Done()
			info, err := contract.BiddersInfo(&copts, roundBig, lastBidder)
			if err != nil {
				recordErr("BiddersInfo(lastBidder)", err)
				return
			}
			mu.Lock()
			lastBidderLastBidTime = info.LastBidTimeStamp
			mu.Unlock()
		}()
		go func() {
			defer wg.Done()
			if lastCstBidder == (ethcommon.Address{}) {
				return
			}
			info, err := contract.BiddersInfo(&copts, roundBig, lastCstBidder)
			if err != nil {
				recordErr("BiddersInfo(lastCstBidder)", err)
				return
			}
			mu.Lock()
			lastCstBidderLastBidTime = info.LastBidTimeStamp
			mu.Unlock()
		}()
		wg.Wait()

		if out.Err != nil {
			return out
		}

		if lastBidderLastBidTime != nil {
			out.LastBidderLastBidTime = lastBidderLastBidTime.Int64()
		}
		if lastCstBidderLastBidTime != nil {
			out.LastCstBidderLastBidTime = lastCstBidderLastBidTime.Int64()
			out.HasLastCstBidderLastBidTime = true
		}

		// Derive the LIVE endurance-champion anchor consistently with
		// tryGetCurrentChampions(). When the current last bidder has overtaken
		// the stored endurance record, the contract recomputes the champion
		// in-memory, but the enduranceChampionStartTimeStamp() and
		// prevEnduranceChampionDuration() storage getters still return the OLD
		// (stale) anchor. Mixing the live champion with that stale anchor
		// produced a wrong Chrono-Warrior "record-growing segment" and
		// "is live" status.
		liveEnduranceStart := enduranceStartTs.Int64()
		livePrevDuration := prevEnduranceDuration.Int64()
		if lastBidder != (ethcommon.Address{}) && lastBidderLastBidTime != nil {
			lastBidDuration := out.SourceBlockTimeStamp - lastBidderLastBidTime.Int64()
			if storedEnduranceChampionAddr == (ethcommon.Address{}) {
				// No champion recorded yet: the last bidder is the (live)
				// endurance champion.
				liveEnduranceStart = lastBidderLastBidTime.Int64()
			} else if storedEnduranceChampionDur != nil && lastBidDuration > storedEnduranceChampionDur.Int64() {
				// Last bidder overtook the stored record: champion start/prev
				// are recomputed live.
				livePrevDuration = storedEnduranceChampionDur.Int64()
				liveEnduranceStart = lastBidderLastBidTime.Int64()
			}
		}
		out.EnduranceChampionStartTimeStamp = liveEnduranceStart
		out.PrevEnduranceChampionDuration = livePrevDuration

		chronoSegmentStart := liveEnduranceStart + livePrevDuration
		currentChronoSegmentDuration := out.SourceBlockTimeStamp - chronoSegmentStart
		out.ChronoWarriorIsLive = currentChronoSegmentDuration > storedChronoWarriorDur.Int64()

		if lastCstBidder != (ethcommon.Address{}) {
			// Any error — including ErrNotFound — leaves the field unset,
			// matching the legacy (id, ok) behavior; real DB failures are
			// logged so they are not silently absorbed.
			evtlogID, err := s.db.LastCstBidEvtlogForBidder(ctx, out.RoundNum, lastCstBidder.String())
			switch {
			case err == nil:
				out.LastCstBidEventLogId = evtlogID
				out.HasLastCstBidEventLogId = true
			case !errors.Is(err, store.ErrNotFound):
				s.errlog.Printf("state refresh: last CST bid lookup: %v", err)
			}
		}
	}

	return out
}
