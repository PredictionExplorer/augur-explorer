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
			return val.Int64()
		}
		s.setMechanicsVersion(mechanicsUnknown)
		cached = mechanicsUnknown
	}
	if cached != mechanicsV2 && v1 != nil {
		if val, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
			s.setMechanicsVersion(mechanicsV1)
			return val.Int64()
		}
	}
	if v2 != nil {
		if val, err := v2.CstDutchAuctionDuration(opts); err == nil {
			s.setMechanicsVersion(mechanicsV2)
			return val.Int64()
		}
	}
	return -1
}

// cstAuctionDurationChangeDivisor returns the V2
// cstDutchAuctionDurationChangeDivisor, or -1 on V1 or error.
func (s *State) cstAuctionDurationChangeDivisor(v1 *cg.CosmicSignatureGame, v2 *cg.CosmicSignatureGameV2, opts *bind.CallOpts) int64 {
	if s.mechanicsVersion() == mechanicsV1 {
		return -1
	}
	if v2 != nil {
		if val, err := v2.CstDutchAuctionDurationChangeDivisor(opts); err == nil {
			s.setMechanicsVersion(mechanicsV2)
			return val.Int64()
		}
	}
	if v1 != nil {
		if _, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
			s.setMechanicsVersion(mechanicsV1)
			return -1
		}
	}
	return -1
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
