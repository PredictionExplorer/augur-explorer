// Fuzz target: arbitrary event data through every RandomWalk ABI event decoder
// (MODERNIZATION.md §4.4, the decode-only half of FuzzEventDecodeRW — the full
// registry-driven fuzz over cmd/rw-etl handlers lands with the Phase 3 engine).
package randomwalk

import (
	"strings"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var rwEventABIs = sync.OnceValue(func() []abi.ABI {
	jsons := []string{
		RWalkMetaData.ABI,
		RWMarketMetaData.ABI,
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
	f.Add(make([]byte, 31))
	huge := make([]byte, 64)
	huge[23] = 0xff
	f.Add(huge)
	f.Fuzz(func(t *testing.T, data []byte) {
		for _, contractABI := range rwEventABIs() {
			for name, ev := range contractABI.Events {
				if _, err := ev.Inputs.NonIndexed().UnpackValues(data); err != nil {
					continue
				}
				_ = name
			}
		}
	})
}
