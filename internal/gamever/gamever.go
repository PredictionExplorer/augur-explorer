// Package gamever detects which version of the CosmicSignatureGame contract
// is deployed at an address by probing version-specific view selectors.
//
// The game is a UUPS proxy upgraded in place (V1 -> V2 -> V3), so tools that
// talk to it cannot assume a fixed ABI: the V2 upgrade changed the bidWithEth
// and bidWithCst signatures (a bidCstRewardAmountMinLimit argument was
// added), and the V3.1 upgrade renamed the bidderAddresses getter to
// bidsInfo. Probing distinguishes the versions:
//
//   - cstBidPriceDeclineMultiplier() exists only on V3 (v3.1 decline-rate
//     CST pricing).
//   - cstDutchAuctionDuration() exists on V2 and V3 (V2 replaced the V1
//     divisor-derived CST auction duration with a directly-set one).
//   - Neither exists on V1.
package gamever

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

// Version is the major version of a deployed CosmicSignatureGame contract.
type Version int

const (
	V1 Version = 1
	V2 Version = 2
	V3 Version = 3
)

func (v Version) String() string { return fmt.Sprintf("V%d", int(v)) }

// BidsTakeMinLimit reports whether bidWithEth/bidWithCst take the trailing
// bidCstRewardAmountMinLimit argument (V2 and later).
func (v Version) BidsTakeMinLimit() bool { return v >= V2 }

// Detect probes the game contract at gameAddr for version-specific
// selectors. A probe that reverts (unrecognized selector) means the getter
// is absent from that deployment. When neither V3 nor V2 getters answer,
// lastBidderAddress() — present and stateless on every version — is used as
// a sanity check that a game contract is deployed at all, so RPC failures
// and wrong addresses surface as errors instead of being misreported as V1.
func Detect(copts *bind.CallOpts, gameAddr common.Address, client bind.ContractCaller) (Version, error) {
	v3, err := cgcontracts.NewCosmicSignatureGameV3Caller(gameAddr, client)
	if err != nil {
		return 0, fmt.Errorf("instantiating V3 caller: %w", err)
	}
	if _, err := v3.CstBidPriceDeclineMultiplier(copts); err == nil {
		return V3, nil
	}
	if _, err := v3.CstDutchAuctionDuration(copts); err == nil {
		return V2, nil
	}
	v1, err := cgcontracts.NewCosmicSignatureGameCaller(gameAddr, client)
	if err != nil {
		return 0, fmt.Errorf("instantiating game caller: %w", err)
	}
	if _, err := v1.LastBidderAddress(copts); err != nil {
		return 0, fmt.Errorf("no CosmicSignatureGame found at %s (lastBidderAddress() failed: %w)", gameAddr, err)
	}
	return V1, nil
}
