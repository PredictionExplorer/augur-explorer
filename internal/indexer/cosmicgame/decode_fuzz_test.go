// FuzzEventDecodeCG (MODERNIZATION.md §4.4): every registered CosmicGame
// handler's Decode must be total — arbitrary topics and data yield a typed
// event or an error, never a panic or a hang. Decode is pure (no DB, no
// RPC), which is exactly what makes this fuzzable; the registry guarantees
// handlers only ever see their own topic0/source in production, but the
// decoder itself must not rely on that for memory safety.
package cosmicgame

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/primitives"
)

// inertCaller satisfies bind.ContractCaller without any chain behind it
// (Decode never performs contract calls; this proves it).
type inertCaller struct{}

func (inertCaller) CodeAt(context.Context, ethcommon.Address, *big.Int) ([]byte, error) {
	return nil, fmt.Errorf("inertCaller: no chain")
}

func (inertCaller) CallContract(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error) {
	return nil, fmt.Errorf("inertCaller: no chain")
}

var _ bind.ContractCaller = inertCaller{}

func FuzzEventDecodeCG(f *testing.F) {
	h := newUnitHandlers(f)
	handlers := h.Registry().Handlers()

	// Seeds: a realistic 4-topic log body, an empty body, a single word and
	// a truncated word.
	word := make([]byte, 32)
	word[31] = 7
	f.Add(3, []byte{}, uint8(4))
	f.Add(7, word, uint8(2))
	f.Add(11, append(append([]byte{}, word...), word...), uint8(1))
	f.Add(0, []byte{0x01}, uint8(0))

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
		elog := &primitives.EthereumEventLog{
			EvtId:    1,
			BlockNum: 2,
			TxId:     3,
			TxHash:   "0x00000000000000000000000000000000000000000000000000000000000000aa",
		}
		// Must never panic; errors are expected for malformed input.
		_, _ = hd.Decode(lg, elog)
	})
}
