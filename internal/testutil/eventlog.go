package testutil

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BuildEventLog constructs a types.Log for an ABI event: topic0 is the event
// ID, indexed arguments are packed into the remaining topics (in declaration
// order), non-indexed arguments into the data segment.
func BuildEventLog(t *testing.T, contractABI *abi.ABI, eventName string, address common.Address, indexed []any, nonIndexed []any) *types.Log {
	t.Helper()
	event, ok := contractABI.Events[eventName]
	if !ok {
		t.Fatalf("event %q not found in ABI", eventName)
	}

	if want, got := len(event.Inputs)-len(event.Inputs.NonIndexed()), len(indexed); want != got {
		t.Fatalf("event %s has %d indexed inputs, fixture provides %d", eventName, want, got)
	}

	data, err := event.Inputs.NonIndexed().Pack(nonIndexed...)
	if err != nil {
		t.Fatalf("packing %s data: %v", eventName, err)
	}

	topics := []common.Hash{event.ID}
	if len(indexed) > 0 {
		query := make([][]any, len(indexed))
		for i, v := range indexed {
			query[i] = []any{v}
		}
		packed, err := abi.MakeTopics(query...)
		if err != nil {
			t.Fatalf("packing %s topics: %v", eventName, err)
		}
		for _, row := range packed {
			topics = append(topics, row[0])
		}
	}

	return &types.Log{Address: address, Topics: topics, Data: data}
}

// BuildRawLog constructs a types.Log from a raw topic0 hex constant (without
// 0x prefix) and 32-byte data words, for events that no current ABI defines.
func BuildRawLog(t *testing.T, topic0Hex string, address common.Address, indexedTopics []common.Hash, dataWords ...*big.Int) *types.Log {
	t.Helper()
	topics := []common.Hash{common.HexToHash("0x" + topic0Hex)}
	topics = append(topics, indexedTopics...)
	var data []byte
	for _, w := range dataWords {
		data = append(data, common.BigToHash(w).Bytes()...)
	}
	return &types.Log{Address: address, Topics: topics, Data: data}
}
