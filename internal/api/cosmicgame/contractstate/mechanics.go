package contractstate

import (
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	cg "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

const (
	mechanicsUnknown int64 = 0
	mechanicsV1      int64 = 1
	mechanicsV2      int64 = 2
	mechanicsV3      int64 = 3
)

func (s *State) mechanicsVersion() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.snap.MechanicsVersion
}

func (s *State) setMechanicsVersion(v int64) {
	s.mu.Lock()
	if s.snap.MechanicsVersion != v {
		s.snap.ConfigurationReady = false
	}
	s.snap.MechanicsVersion = v
	s.mu.Unlock()
}

func (s *State) bindLiveReaders() (
	*cg.CosmicSignatureGame,
	*cg.CosmicSignatureGameV2,
	*cg.CosmicSignatureGameV3,
) {
	v1, _ := cg.NewCosmicSignatureGame(s.addrs.CosmicGame, s.client)
	v2, _ := cg.NewCosmicSignatureGameV2(s.addrs.CosmicGame, s.client)
	v3, _ := cg.NewCosmicSignatureGameV3(s.addrs.CosmicGame, s.client)
	return v1, v2, v3
}

// resolveMechanicsVersion returns the deployed contract generation.
//
// The full newest-first probe calls getters that do not exist on older
// generations, so on (say) a V2 deployment every probe shows up as an
// "execution reverted" eth_call warning in the RPC node's log. To keep that
// noise off the steady-state refresh ticks, the cached version is normally
// just re-confirmed with a getter that exists on that generation — a
// successful, non-reverting call. The full probe runs only when:
//   - there is no cached version yet (startup, or a previous probe failed),
//   - the confirmation call fails (proxy upgrade that removed the old
//     getter, or an RPC error), or
//   - MechanicsProbeInterval has elapsed since the last full probe, which
//     catches upgrades where the old getter still answers (e.g. V2 -> V3).
func (s *State) resolveMechanicsVersion(
	v1 *cg.CosmicSignatureGame,
	v2 *cg.CosmicSignatureGameV2,
	v3 *cg.CosmicSignatureGameV3,
	opts *bind.CallOpts,
) int64 {
	cached := s.mechanicsVersion()
	if cached != mechanicsUnknown && !s.mechanicsFullProbeDue() &&
		s.confirmMechanicsVersion(cached, v1, v2, v3, opts) {
		return cached
	}

	// Full newest-first probe.
	if v3 != nil {
		if _, err := v3.MainPrizeNumCosmicSignatureNfts(opts); err == nil {
			s.setMechanicsVersion(mechanicsV3)
			s.recordMechanicsFullProbe()
			return mechanicsV3
		}
	}
	if v2 != nil {
		if _, err := v2.CstDutchAuctionDurationChangeDivisor(opts); err == nil {
			s.setMechanicsVersion(mechanicsV2)
			s.recordMechanicsFullProbe()
			return mechanicsV2
		}
	}
	if v1 != nil {
		if _, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
			s.setMechanicsVersion(mechanicsV1)
			s.recordMechanicsFullProbe()
			return mechanicsV1
		}
	}
	// Not recorded as a completed probe: with no version detected the next
	// resolve attempt should probe again immediately.
	s.setMechanicsVersion(mechanicsUnknown)
	return mechanicsUnknown
}

// confirmMechanicsVersion checks that the contract still answers a getter
// specific to the cached generation. Unlike the full probe, a match here
// produces no reverting calls.
func (s *State) confirmMechanicsVersion(
	cached int64,
	v1 *cg.CosmicSignatureGame,
	v2 *cg.CosmicSignatureGameV2,
	v3 *cg.CosmicSignatureGameV3,
	opts *bind.CallOpts,
) bool {
	switch cached {
	case mechanicsV3:
		if v3 != nil {
			_, err := v3.MainPrizeNumCosmicSignatureNfts(opts)
			return err == nil
		}
	case mechanicsV2:
		if v2 != nil {
			_, err := v2.CstDutchAuctionDurationChangeDivisor(opts)
			return err == nil
		}
	case mechanicsV1:
		if v1 != nil {
			_, err := v1.CstDutchAuctionDurationDivisor(opts)
			return err == nil
		}
	}
	return false
}

func (s *State) mechanicsFullProbeDue() bool {
	s.mechanicsProbeMu.Lock()
	defer s.mechanicsProbeMu.Unlock()
	return time.Since(s.lastMechanicsFullProbe) >= s.mechanicsProbeInterval
}

func (s *State) recordMechanicsFullProbe() {
	s.mechanicsProbeMu.Lock()
	s.lastMechanicsFullProbe = time.Now()
	s.mechanicsProbeMu.Unlock()
}

func (s *State) roundStartCSTAuctionSetting(
	v1 *cg.CosmicSignatureGame,
	v2 *cg.CosmicSignatureGameV2,
	v3 *cg.CosmicSignatureGameV3,
	opts *bind.CallOpts,
) int64 {
	switch s.resolveMechanicsVersion(v1, v2, v3, opts) {
	case mechanicsV3:
		if v3 != nil {
			duration, _, err := v3.GetCstDutchAuctionDurations(opts)
			if err == nil {
				if parsed, ok := nonNegativeInt64(duration); ok {
					return parsed
				}
			}
		}
	case mechanicsV2:
		if v2 != nil {
			value, err := v2.CstDutchAuctionDuration(opts)
			if err == nil {
				if parsed, ok := nonNegativeInt64(value); ok {
					return parsed
				}
			}
		}
	case mechanicsV1:
		if v1 != nil {
			value, err := v1.CstDutchAuctionDurationDivisor(opts)
			if err == nil {
				if parsed, ok := nonNegativeInt64(value); ok {
					return parsed
				}
			}
		}
	}
	return -1
}

func (s *State) cstAuctionDurationChangeDivisor(
	v1 *cg.CosmicSignatureGame,
	v2 *cg.CosmicSignatureGameV2,
	v3 *cg.CosmicSignatureGameV3,
	opts *bind.CallOpts,
) (int64, bool) {
	switch s.resolveMechanicsVersion(v1, v2, v3, opts) {
	case mechanicsV1:
		return -1, true
	case mechanicsV2:
		if v2 != nil {
			if value, err := v2.CstDutchAuctionDurationChangeDivisor(opts); err == nil {
				parsed, ok := nonNegativeInt64(value)
				return parsed, ok
			}
		}
	case mechanicsV3:
		if v3 != nil {
			if value, err := v3.CstDutchAuctionDurationChangeDivisor(opts); err == nil {
				parsed, ok := nonNegativeInt64(value)
				return parsed, ok
			}
		}
	}
	_ = v1
	return -1, false
}

func (s *State) tokenReward(
	v1 *cg.CosmicSignatureGame,
	v2 *cg.CosmicSignatureGameV2,
	v3 *cg.CosmicSignatureGameV3,
	opts *bind.CallOpts,
) (string, error) {
	switch s.resolveMechanicsVersion(v1, v2, v3, opts) {
	case mechanicsV3:
		if v3 != nil {
			if value, err := v3.GetBidCstRewardAmount(opts); err == nil {
				return value.String(), nil
			}
		}
	case mechanicsV2:
		if v2 != nil {
			if value, err := v2.GetBidCstRewardAmount(opts); err == nil {
				return value.String(), nil
			}
		}
	case mechanicsV1:
		if v1 != nil {
			if value, err := v1.CstRewardAmountForBidding(opts); err == nil {
				return value.String(), nil
			}
		}
	}
	return "", errors.New("cannot read CST bid reward from contract")
}

func (s *State) bidCSTRewardConfiguration(
	v1 *cg.CosmicSignatureGame,
	v2 *cg.CosmicSignatureGameV2,
	v3 *cg.CosmicSignatureGameV3,
	opts *bind.CallOpts,
) (fixedAmount, multiplier string, err error) {
	switch s.resolveMechanicsVersion(v1, v2, v3, opts) {
	case mechanicsV3:
		if v3 != nil {
			if value, callErr := v3.BidCstRewardAmountMultiplier(opts); callErr == nil {
				return "", value.String(), nil
			}
		}
	case mechanicsV2:
		if v2 != nil {
			if value, callErr := v2.BidCstRewardAmountMultiplier(opts); callErr == nil {
				return "", value.String(), nil
			}
		}
	case mechanicsV1:
		if v1 != nil {
			if value, callErr := v1.CstRewardAmountForBidding(opts); callErr == nil {
				return value.String(), "", nil
			}
		}
	}
	return "", "", errors.New("cannot read CST bid reward configuration")
}

func readV3Configuration(
	v3 *cg.CosmicSignatureGameV3,
	opts *bind.CallOpts,
) (V3Configuration, error) {
	if v3 == nil {
		return V3Configuration{}, errors.New("V3 binding unavailable")
	}
	var out V3Configuration
	value, err := v3.RoundLateBidDurationDivisor(opts)
	if err != nil {
		return out, fmt.Errorf("roundLateBidDurationDivisor: %w", err)
	}
	out.RoundLateBidDurationDivisor = value.String()
	value, err = v3.GetRoundLateBidDuration(opts)
	if err != nil {
		return out, fmt.Errorf("getRoundLateBidDuration: %w", err)
	}
	if out.RoundLateBidDurationSeconds, _ = nonNegativeInt64(value); out.RoundLateBidDurationSeconds <= 0 {
		return out, errors.New("getRoundLateBidDuration is not positive")
	}
	value, err = v3.RoundLateBidPricePremiumAmountBaseMultiplier(opts)
	if err != nil {
		return out, fmt.Errorf("roundLateBidPricePremiumAmountBaseMultiplier: %w", err)
	}
	out.RoundLateBidPricePremiumAmountBaseMultiplier = value.String()
	value, err = v3.RoundLateBidPricePremiumAmountExponent(opts)
	if err != nil {
		return out, fmt.Errorf("roundLateBidPricePremiumAmountExponent: %w", err)
	}
	var ok bool
	out.RoundLateBidPricePremiumAmountExponent, ok = nonNegativeInt64(value)
	if !ok {
		return out, errors.New("roundLateBidPricePremiumAmountExponent exceeds int64")
	}
	value, err = v3.MainPrizeNumCosmicSignatureNfts(opts)
	if err != nil {
		return out, fmt.Errorf("mainPrizeNumCosmicSignatureNfts: %w", err)
	}
	out.MainPrizeNumCosmicSignatureNfts, ok = nonNegativeInt64(value)
	if !ok || out.MainPrizeNumCosmicSignatureNfts <= 0 {
		return out, errors.New("mainPrizeNumCosmicSignatureNfts is invalid")
	}
	value, err = v3.CstDutchAuctionBeginningBidPriceMinLimit(opts)
	if err != nil {
		return out, fmt.Errorf("cstDutchAuctionBeginningBidPriceMinLimit: %w", err)
	}
	out.CstDutchAuctionBeginningBidPriceMinLimit = value.String()
	return out, nil
}
