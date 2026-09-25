// Unit test (no Docker) enforcing that the ETL dispatches every event a
// deployed CosmicGame-family contract can emit, including the events the
// contracts inherit from OpenZeppelin (approvals, votes, governor,
// ownership, initializer, proxy). A new event in a regenerated ABI without a
// handler fails here instead of being silently left in evt_log.
package cosmicgame

import (
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

// TestEventErrorNameCollisions pins the event names that share a name with a
// custom error in the same ABI. abi.UnpackIntoInterface / abi.Unpack resolve
// such a name to the *error* (all inputs non-indexed), which silently shifts
// the decode of the event's data: FundTransferFailed stored the errStr
// length as the amount until decodeFundTransferFailed switched to the
// event's NonIndexed inputs. Any new collision must decode the same way.
func TestEventErrorNameCollisions(t *testing.T) {
	want := map[string]bool{"FundTransferFailed": true}
	for name, raw := range map[string]string{
		"CosmicSignatureGame":             cgc.CosmicSignatureGameABI,
		"CosmicSignatureGameV2":           cgc.CosmicSignatureGameV2ABI,
		"CosmicSignatureGameV3":           cgc.CosmicSignatureGameV3ABI,
		"CosmicSignatureNft":              cgc.CosmicSignatureNftABI,
		"CosmicSignatureToken":            cgc.CosmicSignatureTokenABI,
		"CosmicSignatureDao":              cgc.CosmicSignatureDaoABI,
		"CharityWallet":                   cgc.CharityWalletABI,
		"PrizesWallet":                    cgc.PrizesWalletABI,
		"StakingWalletCosmicSignatureNft": cgc.StakingWalletCosmicSignatureNftABI,
		"StakingWalletRandomWalkNft":      cgc.StakingWalletRandomWalkNftABI,
		"MarketingWallet":                 cgc.MarketingWalletABI,
	} {
		parsed := mustABI(t, raw)
		for ev := range parsed.Events {
			if _, ok := parsed.Errors[ev]; ok && !want[ev] {
				t.Errorf("%s: event %s shares its name with a custom error; its decoder must unpack via Events[%q].Inputs.NonIndexed()", name, ev, ev)
			}
		}
	}
}

func TestEveryABIEventHasHandler(t *testing.T) {
	h := newUnitHandlers(t)
	c := h.c

	// handled[topic0] = contracts whose logs the registry admits for it.
	handled := make(map[ethcommon.Hash]map[ethcommon.Address]bool)
	for _, hd := range h.Registry().Handlers() {
		set := handled[hd.Topic()]
		if set == nil {
			set = make(map[ethcommon.Address]bool)
			handled[hd.Topic()] = set
		}
		for _, src := range hd.Sources() {
			set[src] = true
		}
	}

	deployed := []struct {
		name string
		abi  string
		addr ethcommon.Address
	}{
		// The proxy address carries every game generation.
		{"CosmicSignatureGame", cgc.CosmicSignatureGameABI, c.Game},
		{"CosmicSignatureGameV2", cgc.CosmicSignatureGameV2ABI, c.Game},
		{"CosmicSignatureGameV3", cgc.CosmicSignatureGameV3ABI, c.Game},
		{"CosmicSignatureNft", cgc.CosmicSignatureNftABI, c.Signature},
		{"CosmicSignatureToken", cgc.CosmicSignatureTokenABI, c.Token},
		{"CosmicSignatureDao", cgc.CosmicSignatureDaoABI, c.Dao},
		{"CharityWallet", cgc.CharityWalletABI, c.CharityWallet},
		{"PrizesWallet", cgc.PrizesWalletABI, c.PrizesWallet},
		{"StakingWalletCosmicSignatureNft", cgc.StakingWalletCosmicSignatureNftABI, c.StakingCST},
		{"StakingWalletRandomWalkNft", cgc.StakingWalletRandomWalkNftABI, c.StakingRWalk},
		{"MarketingWallet", cgc.MarketingWalletABI, c.MarketingWallet},
		{"DonatedTokenHolder", cgc.DonatedTokenHolderABI, ethcommon.Address{}},
	}

	for _, d := range deployed {
		parsed := mustABI(t, d.abi)
		for name, ev := range parsed.Events {
			if d.addr == (ethcommon.Address{}) {
				t.Errorf("%s declares %s but the ETL does not watch that contract", d.name, name)
				continue
			}
			if !handled[ev.ID][d.addr] {
				t.Errorf("%s.%s (%s) has no handler admitting the %s contract",
					d.name, name, ev.ID.Hex(), d.name)
			}
		}
	}
}
