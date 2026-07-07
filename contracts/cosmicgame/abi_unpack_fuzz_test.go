// Fuzz target: arbitrary event data through every CosmicGame ABI event decoder
// (MODERNIZATION.md §4.4, the decode-only half of FuzzEventDecodeCG — the full
// registry-driven fuzz over cmd/cg-etl handlers lands with the Phase 3 engine,
// which separates decoding from persistence).
package cosmicgame

import (
	"strings"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// cgEventABIs parses every binding ABI the ETL decodes events from.
var cgEventABIs = sync.OnceValue(func() []abi.ABI {
	jsons := []string{
		CosmicSignatureGameMetaData.ABI,
		CosmicSignatureGameV2MetaData.ABI,
		CosmicSignatureNftMetaData.ABI,
		CosmicSignatureTokenMetaData.ABI,
		PrizesWalletMetaData.ABI,
		CharityWalletMetaData.ABI,
		MarketingWalletMetaData.ABI,
		StakingWalletCosmicSignatureNftMetaData.ABI,
		StakingWalletRandomWalkNftMetaData.ABI,
		ERC20MetaData.ABI,
	}
	parsed := make([]abi.ABI, 0, len(jsons))
	for _, j := range jsons {
		a, err := abi.JSON(strings.NewReader(j))
		if err != nil {
			panic(err)
		}
		parsed = append(parsed, a)
	}
	return parsed
})

func FuzzABIEventUnpack(f *testing.F) {
	f.Add([]byte(nil))
	f.Add(make([]byte, 32))
	f.Add(make([]byte, 31)) // one byte short of a word
	// Offset pointing far outside the payload (classic dynamic-type trap).
	huge := make([]byte, 64)
	huge[23] = 0xff
	f.Add(huge)
	f.Fuzz(func(t *testing.T, data []byte) {
		for _, contractABI := range cgEventABIs() {
			for name, ev := range contractABI.Events {
				if _, err := ev.Inputs.NonIndexed().UnpackValues(data); err != nil {
					continue
				}
				_ = name // decode succeeded; nothing further to assert, must only not panic
			}
		}
	})
}
