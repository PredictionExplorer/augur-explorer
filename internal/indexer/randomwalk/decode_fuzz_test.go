// FuzzEventDecodeRW (MODERNIZATION.md §4.4): every registered RandomWalk
// handler's Decode must be total — arbitrary topics and data yield a typed
// event or an error, never a panic or a hang.
package randomwalk

import (
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func FuzzEventDecodeRW(f *testing.F) {
	handlers := newUnitHandlers(f).Registry().Handlers()

	word := make([]byte, 32)
	word[31] = 7
	f.Add(0, []byte{}, uint8(4))
	f.Add(3, word, uint8(2))
	f.Add(5, append(append([]byte{}, word...), word...), uint8(1))
	f.Add(6, []byte{0x01}, uint8(0))

	f.Fuzz(func(t *testing.T, handlerIdx int, data []byte, numTopics uint8) {
		if handlerIdx < 0 {
			handlerIdx = -handlerIdx
		}
		hd := handlers[handlerIdx%len(handlers)]

		topics := make([]ethcommon.Hash, int(numTopics)%6)
		if len(topics) > 0 {
			topics[0] = hd.Topic()
			for i := 1; i < len(topics); i++ {
				topics[i] = ethcommon.Hash{31: byte(i)}
			}
		}
		lg := &types.Log{
			Address: hd.Sources()[0],
			Topics:  topics,
			Data:    data,
		}
		elog := &store.EthereumEventLog{
			EvtId:    1,
			BlockNum: 2,
			TxId:     3,
			TxHash:   "0x00000000000000000000000000000000000000000000000000000000000000aa",
		}
		// Must never panic; errors are expected for malformed input.
		_, _ = hd.Decode(lg, elog)
	})
}
