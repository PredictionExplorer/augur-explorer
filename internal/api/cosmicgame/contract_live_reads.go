package cosmicgame

import (
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

const (
	contractMechanicsUnknown int64 = 0
	contractMechanicsV1    int64 = 1
	contractMechanicsV2    int64 = 2
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

func bindCosmicGameLiveReaders(addr ethcommon.Address, backend bind.ContractBackend) (*CosmicSignatureGame, *CosmicSignatureGameV2) {
	var v1 *CosmicSignatureGame
	var v2 *CosmicSignatureGameV2
	if backend != nil {
		v1, _ = NewCosmicSignatureGame(addr, backend)
		v2, _ = NewCosmicSignatureGameV2(addr, backend)
	}
	return v1, v2
}

// readRoundStartCSTAuctionSetting returns the live CST round-start auction parameter.
// V1: cstDutchAuctionDurationDivisor (divisor). V2: cstDutchAuctionDuration (seconds), same storage slot.
func readRoundStartCSTAuctionSetting(v1 *CosmicSignatureGame, v2 *CosmicSignatureGameV2, opts *bind.CallOpts) int64 {
	cached := getContractMechanicsVersion()
	if cached == contractMechanicsV2 && v2 != nil {
		if val, err := v2.CstDutchAuctionDuration(opts); err == nil {
			return val.Int64()
		}
		setContractMechanicsVersion(contractMechanicsUnknown)
		cached = contractMechanicsUnknown
	}
	if cached != contractMechanicsV2 && v1 != nil {
		if val, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
			setContractMechanicsVersion(contractMechanicsV1)
			return val.Int64()
		}
	}
	if v2 != nil {
		if val, err := v2.CstDutchAuctionDuration(opts); err == nil {
			setContractMechanicsVersion(contractMechanicsV2)
			return val.Int64()
		}
	}
	return -1
}

// readCSTAuctionDurationChangeDivisor returns V2 cstDutchAuctionDurationChangeDivisor, or -1 on V1 / error.
func readCSTAuctionDurationChangeDivisor(v1 *CosmicSignatureGame, v2 *CosmicSignatureGameV2, opts *bind.CallOpts) int64 {
	cached := getContractMechanicsVersion()
	if cached == contractMechanicsV1 {
		return -1
	}
	if v2 != nil {
		if val, err := v2.CstDutchAuctionDurationChangeDivisor(opts); err == nil {
			setContractMechanicsVersion(contractMechanicsV2)
			return val.Int64()
		}
	}
	if v1 != nil {
		if _, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
			setContractMechanicsVersion(contractMechanicsV1)
			return -1
		}
	}
	return -1
}

// readTokenReward returns the CST reward for bidding: fixed amount on V1, computed next-bid reward on V2.
func readTokenReward(v1 *CosmicSignatureGame, v2 *CosmicSignatureGameV2, opts *bind.CallOpts) (string, error) {
	cached := getContractMechanicsVersion()
	if cached == contractMechanicsV2 && v2 != nil {
		if val, err := v2.GetBidCstRewardAmount(opts); err == nil {
			return val.String(), nil
		}
		setContractMechanicsVersion(contractMechanicsUnknown)
		cached = contractMechanicsUnknown
	}
	if cached != contractMechanicsV2 && v1 != nil {
		if val, err := v1.CstRewardAmountForBidding(opts); err == nil {
			setContractMechanicsVersion(contractMechanicsV1)
			return val.String(), nil
		}
	}
	if v2 != nil {
		if val, err := v2.GetBidCstRewardAmount(opts); err == nil {
			setContractMechanicsVersion(contractMechanicsV2)
			return val.String(), nil
		}
	}
	return "", errors.New("can't read CST bid reward from contract")
}
