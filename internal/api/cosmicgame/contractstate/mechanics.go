package contractstate

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	cg "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

// Contract mechanics generations. The deployed CosmicSignatureGame proxy has
// two implementation generations with different CST auction parameters; the
// version is detected lazily from which reads succeed and cached until a
// read contradicts it.
const (
	mechanicsUnknown int64 = 0
	mechanicsV1      int64 = 1
	mechanicsV2      int64 = 2
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

// bindLiveReaders instantiates the V1 and V2 read bindings on the game
// address. The constructors only fail on a nil backend, which New rules out.
func (s *State) bindLiveReaders() (*cg.CosmicSignatureGame, *cg.CosmicSignatureGameV2) {
	v1, _ := cg.NewCosmicSignatureGame(s.addrs.CosmicGame, s.client)
	v2, _ := cg.NewCosmicSignatureGameV2(s.addrs.CosmicGame, s.client)
	return v1, v2
}

// roundStartCSTAuctionSetting returns the live CST round-start auction
// parameter: the V1 divisor (cstDutchAuctionDurationDivisor) or the V2
// duration in seconds (cstDutchAuctionDuration), which share a storage slot.
// Returns -1 when neither generation answers.
func (s *State) roundStartCSTAuctionSetting(v1 *cg.CosmicSignatureGame, v2 *cg.CosmicSignatureGameV2, opts *bind.CallOpts) int64 {
	cached := s.mechanicsVersion()
	if cached == mechanicsV2 && v2 != nil {
		if val, err := v2.CstDutchAuctionDuration(opts); err == nil {
			parsed, ok := nonNegativeInt64(val)
			if !ok {
				s.setMechanicsVersion(mechanicsUnknown)
				return -1
			}
			return parsed
		}
		s.setMechanicsVersion(mechanicsUnknown)
		cached = mechanicsUnknown
	}
	if cached != mechanicsV2 && v1 != nil {
		if val, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
			parsed, ok := nonNegativeInt64(val)
			if !ok {
				return -1
			}
			s.setMechanicsVersion(mechanicsV1)
			return parsed
		}
	}
	if v2 != nil {
		if val, err := v2.CstDutchAuctionDuration(opts); err == nil {
			parsed, ok := nonNegativeInt64(val)
			if !ok {
				return -1
			}
			s.setMechanicsVersion(mechanicsV2)
			return parsed
		}
	}
	return -1
}

// cstAuctionDurationChangeDivisor returns the V2
// cstDutchAuctionDurationChangeDivisor, or -1 on V1 or error.
func (s *State) cstAuctionDurationChangeDivisor(
	v1 *cg.CosmicSignatureGame,
	v2 *cg.CosmicSignatureGameV2,
	opts *bind.CallOpts,
) (int64, bool) {
	if s.mechanicsVersion() == mechanicsV1 {
		return -1, true
	}
	if v2 != nil {
		if val, err := v2.CstDutchAuctionDurationChangeDivisor(opts); err == nil {
			parsed, ok := nonNegativeInt64(val)
			if !ok {
				return -1, false
			}
			s.setMechanicsVersion(mechanicsV2)
			return parsed, true
		}
	}
	if v1 != nil {
		if _, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
			s.setMechanicsVersion(mechanicsV1)
			return -1, true
		}
	}
	return -1, false
}

// tokenReward returns the CST reward for bidding: the fixed amount on V1,
// the computed next-bid reward on V2.
func (s *State) tokenReward(v1 *cg.CosmicSignatureGame, v2 *cg.CosmicSignatureGameV2, opts *bind.CallOpts) (string, error) {
	cached := s.mechanicsVersion()
	if cached == mechanicsV2 && v2 != nil {
		if val, err := v2.GetBidCstRewardAmount(opts); err == nil {
			return val.String(), nil
		}
		s.setMechanicsVersion(mechanicsUnknown)
		cached = mechanicsUnknown
	}
	if cached != mechanicsV2 && v1 != nil {
		if val, err := v1.CstRewardAmountForBidding(opts); err == nil {
			s.setMechanicsVersion(mechanicsV1)
			return val.String(), nil
		}
	}
	if v2 != nil {
		if val, err := v2.GetBidCstRewardAmount(opts); err == nil {
			s.setMechanicsVersion(mechanicsV2)
			return val.String(), nil
		}
	}
	return "", errors.New("can't read CST bid reward from contract")
}

// bidCSTRewardConfiguration returns the mechanics-specific static reward
// setting: a fixed wei amount on V1 or the V2 multiplier. The dynamic V2
// next-bid reward is refreshed with the live variable group instead.
func (s *State) bidCSTRewardConfiguration(
	v1 *cg.CosmicSignatureGame,
	v2 *cg.CosmicSignatureGameV2,
	opts *bind.CallOpts,
) (fixedAmount, multiplier string, err error) {
	cached := s.mechanicsVersion()
	if cached == mechanicsV2 && v2 != nil {
		if val, callErr := v2.BidCstRewardAmountMultiplier(opts); callErr == nil {
			return "", val.String(), nil
		}
		s.setMechanicsVersion(mechanicsUnknown)
		cached = mechanicsUnknown
	}
	if cached != mechanicsV2 && v1 != nil {
		if val, callErr := v1.CstRewardAmountForBidding(opts); callErr == nil {
			s.setMechanicsVersion(mechanicsV1)
			return val.String(), "", nil
		}
	}
	if v2 != nil {
		if val, callErr := v2.BidCstRewardAmountMultiplier(opts); callErr == nil {
			s.setMechanicsVersion(mechanicsV2)
			return "", val.String(), nil
		}
	}
	return "", "", errors.New("can't read CST bid reward configuration")
}
