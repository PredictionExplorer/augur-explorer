package cosmicgame

import (
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
)

const (
	contractMechanicsUnknown int64 = 0
	contractMechanicsV1    int64 = 1
	contractMechanicsV2    int64 = 2
	contractMechanicsV3    int64 = 3
)

var (
	contractMechanicsMu sync.RWMutex
	contractMechanicsVersion int64 = contractMechanicsUnknown
)

func getContractMechanicsVersion() int64 {
	contractMechanicsMu.RLock()
	defer contractMechanicsMu.RUnlock()
	return contractMechanicsVersion
}

func setContractMechanicsVersion(v int64) {
	contractMechanicsMu.Lock()
	contractMechanicsVersion = v
	contractMechanicsMu.Unlock()
}

func bindCosmicGameLiveReaders(addr ethcommon.Address, backend bind.ContractBackend) (*CosmicSignatureGame, *CosmicSignatureGameV2, *CosmicSignatureGameV3) {
	var v1 *CosmicSignatureGame
	var v2 *CosmicSignatureGameV2
	var v3 *CosmicSignatureGameV3
	if backend != nil {
		v1, _ = NewCosmicSignatureGame(addr, backend)
		v2, _ = NewCosmicSignatureGameV2(addr, backend)
		v3, _ = NewCosmicSignatureGameV3(addr, backend)
	}
	return v1, v2, v3
}

// resolveMechanicsVersion probes the live contract to determine whether it is running V1, V2 or V3
// mechanics, and caches the result (readable via getContractMechanicsVersion()).
//
// It always re-probes so a live V2 -> V3 (or V1 -> V2) upgrade is detected on the next constants refresh
// without a server restart. Detection is cheap (one eth_call per version tier) and runs on a timer, not
// per HTTP request. Probing order matters: V3 is an ABI superset of V2 (which is a superset of V1), so we
// must test the most-derived version-only getter first.
func resolveMechanicsVersion(v1 *CosmicSignatureGame, v2 *CosmicSignatureGameV2, v3 *CosmicSignatureGameV3, opts *bind.CallOpts) int64 {
	// `mainPrizeNumCosmicSignatureNfts` exists only in V3.
	if v3 != nil {
		if _, err := v3.MainPrizeNumCosmicSignatureNfts(opts); err == nil {
			setContractMechanicsVersion(contractMechanicsV3)
			return contractMechanicsV3
		}
	}
	// `cstDutchAuctionDurationChangeDivisor` exists in V2 (and V3) but not V1.
	if v2 != nil {
		if _, err := v2.CstDutchAuctionDurationChangeDivisor(opts); err == nil {
			setContractMechanicsVersion(contractMechanicsV2)
			return contractMechanicsV2
		}
	}
	// `cstDutchAuctionDurationDivisor` is a V1 public getter (the slot was repurposed in V2).
	if v1 != nil {
		if _, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
			setContractMechanicsVersion(contractMechanicsV1)
			return contractMechanicsV1
		}
	}
	setContractMechanicsVersion(contractMechanicsUnknown)
	return contractMechanicsUnknown
}

// readRoundStartCSTAuctionSetting returns the live CST round-start auction parameter.
// V1: cstDutchAuctionDurationDivisor (divisor). V2/V3: cstDutchAuctionDuration (seconds), same storage slot.
func readRoundStartCSTAuctionSetting(v1 *CosmicSignatureGame, v2 *CosmicSignatureGameV2, v3 *CosmicSignatureGameV3, opts *bind.CallOpts) int64 {
	switch getContractMechanicsVersion() {
	case contractMechanicsV3, contractMechanicsV2:
		// V3 keeps the V2 selector, so the V2 binding reads it fine on a V3 contract.
		if v2 != nil {
			if val, err := v2.CstDutchAuctionDuration(opts); err == nil {
				return val.Int64()
			}
		}
	case contractMechanicsV1:
		if v1 != nil {
			if val, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
				return val.Int64()
			}
		}
	}
	return -1
}

// readCSTAuctionDurationChangeDivisor returns V2/V3 cstDutchAuctionDurationChangeDivisor, or -1 on V1 / error.
func readCSTAuctionDurationChangeDivisor(v1 *CosmicSignatureGame, v2 *CosmicSignatureGameV2, v3 *CosmicSignatureGameV3, opts *bind.CallOpts) int64 {
	switch getContractMechanicsVersion() {
	case contractMechanicsV3, contractMechanicsV2:
		if v2 != nil {
			if val, err := v2.CstDutchAuctionDurationChangeDivisor(opts); err == nil {
				return val.Int64()
			}
		}
	}
	return -1
}

// readTokenReward returns the CST reward for bidding: fixed amount on V1, computed next-bid reward on V2/V3.
func readTokenReward(v1 *CosmicSignatureGame, v2 *CosmicSignatureGameV2, v3 *CosmicSignatureGameV3, opts *bind.CallOpts) (string, error) {
	switch getContractMechanicsVersion() {
	case contractMechanicsV3, contractMechanicsV2:
		if v2 != nil {
			if val, err := v2.GetBidCstRewardAmount(opts); err == nil {
				return val.String(), nil
			}
		}
	case contractMechanicsV1:
		if v1 != nil {
			if val, err := v1.CstRewardAmountForBidding(opts); err == nil {
				return val.String(), nil
			}
		}
	}
	return "", errors.New("can't read CST bid reward from contract")
}

// V3LiveConfig holds the live values of the 5 V3-only configuration parameters (ISystemEventsV3),
// plus the derived late-bid window duration. IsV3 is false when the contract is still V1/V2, in which
// case the other fields are unset and the dashboard hides the V3 section.
type V3LiveConfig struct {
	IsV3                              bool
	RoundLateBidDurationDivisor       string // roundLateBidDurationDivisor
	RoundLateBidDurationSeconds       int64  // getRoundLateBidDuration() — derived late-bid window, in seconds
	RoundLateBidPremiumBaseMultiplier string // roundLateBidPricePremiumAmountBaseMultiplier
	RoundLateBidPremiumExponent       int64  // roundLateBidPricePremiumAmountExponent
	BidCstRewardAmountPerMinute       string // bidCstRewardAmountPerMinute (CST wei per minute)
	MainPrizeNumCosmicSignatureNfts   int64  // mainPrizeNumCosmicSignatureNfts
}

// readV3Config reads the 5 new V3 configuration getters (and the derived late-bid window). Returns
// IsV3=false when the contract is not (yet) V3, so callers can hide the V3 panel.
func readV3Config(v3 *CosmicSignatureGameV3, opts *bind.CallOpts) V3LiveConfig {
	var cfg V3LiveConfig
	if v3 == nil || getContractMechanicsVersion() != contractMechanicsV3 {
		return cfg
	}
	cfg.IsV3 = true
	if val, err := v3.RoundLateBidDurationDivisor(opts); err == nil {
		cfg.RoundLateBidDurationDivisor = val.String()
	}
	if val, err := v3.GetRoundLateBidDuration(opts); err == nil {
		cfg.RoundLateBidDurationSeconds = val.Int64()
	}
	if val, err := v3.RoundLateBidPricePremiumAmountBaseMultiplier(opts); err == nil {
		cfg.RoundLateBidPremiumBaseMultiplier = val.String()
	}
	if val, err := v3.RoundLateBidPricePremiumAmountExponent(opts); err == nil {
		cfg.RoundLateBidPremiumExponent = val.Int64()
	}
	if val, err := v3.BidCstRewardAmountPerMinute(opts); err == nil {
		cfg.BidCstRewardAmountPerMinute = val.String()
	}
	if val, err := v3.MainPrizeNumCosmicSignatureNfts(opts); err == nil {
		cfg.MainPrizeNumCosmicSignatureNfts = val.Int64()
	}
	return cfg
}

// bidCstRewardAmountPerMinuteEth converts the wei-per-minute reward to a float (CST per minute).
func (c V3LiveConfig) BidCstRewardAmountPerMinuteEth() float64 {
	if c.BidCstRewardAmountPerMinute == "" {
		return 0
	}
	v, ok := new(big.Int).SetString(c.BidCstRewardAmountPerMinute, 10)
	if !ok {
		return 0
	}
	f := new(big.Float).Quo(new(big.Float).SetInt(v), big.NewFloat(1e18))
	out, _ := f.Float64()
	return out
}
