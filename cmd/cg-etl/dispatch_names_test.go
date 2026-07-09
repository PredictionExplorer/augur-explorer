// Unit tests (no Docker) for the dispatch table's event names, which label
// the engine's rwcg_etl_events_total metric.
package main

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestDispatchTableNames(t *testing.T) {
	nameByTopic := make(map[common.Hash]string)
	for i, entry := range eventDispatchTable() {
		if entry.name == "" {
			t.Errorf("dispatch entry %d has an empty event name", i)
			continue
		}
		topic := common.BytesToHash(entry.topic0)
		if prev, ok := nameByTopic[topic]; ok && prev != entry.name {
			// Duplicate-topic events (CharityAddressChanged from two
			// contracts, FundsTransferredToCharity) must agree on one label,
			// otherwise the metric label depends on table order.
			t.Errorf("topic %s carries two names: %q and %q", topic.Hex(), prev, entry.name)
		}
		nameByTopic[topic] = entry.name
	}
}

func TestEventTopicName(t *testing.T) {
	bidTopic, err := hex.DecodeString(BID_EVENT)
	if err != nil {
		t.Fatalf("decoding BID_EVENT: %v", err)
	}
	if got := eventTopicName(common.BytesToHash(bidTopic)); got != "BidPlaced" {
		t.Errorf("eventTopicName(BID_EVENT) = %q, want BidPlaced", got)
	}
	if got := eventTopicName(common.HexToHash("0xdead")); got != "" {
		t.Errorf("eventTopicName(unknown) = %q, want empty", got)
	}
}
