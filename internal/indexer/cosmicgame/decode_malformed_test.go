// Deterministic companion to FuzzEventDecodeCG: every registered handler's
// Decode is driven through the malformed-input shapes that fuzzing finds
// (missing indexed topics, truncated and misaligned data words), pinning the
// error branches in a regular `go test` run.
package cosmicgame

import (
	"fmt"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestEventDecodeMalformedInputs(t *testing.T) {
	h := newUnitHandlers(t)
	handlers := h.Registry().Handlers()
	if len(handlers) == 0 {
		t.Fatal("no handlers registered")
	}

	word := make([]byte, 32)
	word[31] = 7
	shapes := []struct {
		name      string
		numTopics int // including topic0
		data      []byte
	}{
		{"topic0 only, no data", 1, nil},
		{"topic0 only, one word", 1, word},
		{"misaligned data", 5, []byte{0x01}},
		{"one topic, truncated word", 2, word[:31]},
		{"five topics, two words", 5, append(append([]byte{}, word...), word...)},
	}

	elog := &store.EthereumEventLog{EvtID: 1, BlockNum: 2, TxID: 3}
	for _, hd := range handlers {
		for _, shape := range shapes {
			name := fmt.Sprintf("%s/%s", hd.Name(), shape.name)
			t.Run(name, func(t *testing.T) {
				topics := make([]ethcommon.Hash, shape.numTopics)
				topics[0] = hd.Topic()
				for i := 1; i < len(topics); i++ {
					topics[i] = ethcommon.Hash{31: byte(i)}
				}
				lg := &types.Log{
					Address: ethcommon.HexToAddress("0x1111111111111111111111111111111111111111"),
					Topics:  topics,
					Data:    shape.data,
				}
				// Decode must be total: a typed event or an error, never a
				// panic. Which of the two depends on the event's layout
				// (some events have no non-indexed body to misdecode).
				evt, err := hd.Decode(lg, elog)
				if err == nil && evt == nil {
					t.Errorf("Decode returned neither event nor error")
				}
			})
		}
	}
}
