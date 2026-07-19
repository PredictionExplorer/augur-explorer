package cosmicgame

// Benchmarks for the ETL hot path (§4.5 of docs/MODERNIZATION.md): decoding
// the most common event, BidPlaced. Baselines live in docs/benchmarks.md;
// re-run after each rewrite phase with:
//
//	go test ./cmd/cg-etl/ -bench BenchmarkEventDecode -benchmem -count=6

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

// The v1 BidPlaced log carries (roundNum, bidder, randomWalkNftId) as indexed
// topics and (paidEthPrice, paidCstPrice, message, mainPrizeTime) as data.

// buildBidPlacedLog packs a realistic v1 BidPlaced log (ETH bid with a short
// message), mirroring what the fake chain produces for the fixture suite.
func buildBidPlacedLog(b *testing.B, cgABI *abi.ABI) *types.Log {
	b.Helper()
	event, ok := cgABI.Events["BidPlaced"]
	if !ok {
		b.Fatal("BidPlaced not found in CosmicSignatureGame ABI")
	}
	data, err := event.Inputs.NonIndexed().Pack(
		big.NewInt(100000000000000000), // paidEthPrice: 0.1 ETH
		big.NewInt(-1),                 // paidCstPrice: not a CST bid
		"benchmark bid message",        // message
		big.NewInt(1767226700),         // mainPrizeTime
	)
	if err != nil {
		b.Fatalf("packing BidPlaced data: %v", err)
	}
	topicRows, err := abi.MakeTopics(
		[]any{big.NewInt(3)}, // roundNum
		[]any{ethcommon.HexToAddress("0x2100000000000000000000000000000000000021")}, // bidder
		[]any{big.NewInt(-1)}, // randomWalkNftId: -1 => plain ETH bid
	)
	if err != nil {
		b.Fatalf("packing BidPlaced topics: %v", err)
	}
	topics := make([]ethcommon.Hash, 0, 1+len(topicRows))
	topics = append(topics, event.ID)
	for _, row := range topicRows {
		topics = append(topics, row[0])
	}
	return &types.Log{
		Address: ethcommon.HexToAddress("0x2000000000000000000000000000000000000002"),
		Topics:  topics,
		Data:    data,
	}
}

// BenchmarkEventDecode measures the decode half of proc_bid_event_v1: ABI
// unpack of the data segment plus the indexed-topic extractions. Persistence
// is excluded (Phase 3 separates decode from store).
func BenchmarkEventDecode(b *testing.B) {
	parsed, err := abi.JSON(strings.NewReader(cgc.CosmicSignatureGameABI))
	if err != nil {
		b.Fatalf("parsing CosmicSignatureGame ABI: %v", err)
	}
	lg := buildBidPlacedLog(b, &parsed)

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		var ethEvt cgc.CosmicSignatureGameBidPlaced
		if err := parsed.UnpackIntoInterface(&ethEvt, "BidPlaced", lg.Data); err != nil {
			b.Fatalf("unpack: %v", err)
		}
		bidder := ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
		roundNum := lg.Topics[1].Big().Int64()
		rwalkID := lg.Topics[3].Big().Int64()
		if bidder == "" || roundNum != 3 || rwalkID >= 0 {
			b.Fatal("decoded values drifted")
		}
	}
}
