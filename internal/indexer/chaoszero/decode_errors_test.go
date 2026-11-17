package chaoszero

import (
	"math/big"
	"strings"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	czstore "github.com/PredictionExplorer/augur-explorer/internal/store/chaoszero"
)

// The decode layer must reject malformed logs with an error (failing the
// batch) instead of storing a half-decoded event. These tests need no
// database: decoding happens before any store call.

// decodeHandlers builds a handler set good enough for decoding (the nil
// store is never touched by decode*).
func decodeHandlers(t *testing.T) *Handlers {
	t.Helper()
	h, err := New(Config{Repo: czstore.NewRepo(nil)})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return h
}

func TestNewRejectsMissingRepo(t *testing.T) {
	if _, err := New(Config{}); err == nil || !strings.Contains(err.Error(), "Config.Repo is required") {
		t.Fatalf("New(Config{}) err = %v, want the Repo requirement", err)
	}
}

// overflowTopic is a uint256 topic beyond int64 (and uint8).
var overflowTopic = ethcommon.HexToHash("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")

// topic builds a small-number topic.
func topic(n int64) ethcommon.Hash {
	return ethcommon.BigToHash(big.NewInt(n))
}

// packEventData packs the non-indexed arguments of a market/vault event.
func packEventData(t *testing.T, h *Handlers, vault bool, event string, args ...any) []byte {
	t.Helper()
	contractABI := h.marketABI
	if vault {
		contractABI = h.vaultABI
	}
	data, err := contractABI.Events[event].Inputs.NonIndexed().Pack(args...)
	if err != nil {
		t.Fatalf("packing %s: %v", event, err)
	}
	return data
}

func TestDecodeRejectsMalformedLogs(t *testing.T) {
	h := decodeHandlers(t)
	elog := &store.EthereumEventLog{EvtID: 1, BlockNum: 1, TxID: 1, TimeStamp: 1767225600}
	bigOne := big.NewInt(1)
	big256 := new(big.Int).Lsh(bigOne, 256)
	_ = big256
	// A position id whose round (id >> 4) exceeds int64.
	hugePositionID := new(big.Int).Lsh(bigOne, 70)

	// Shorthand builders over the decode methods.
	cases := []struct {
		name    string
		decode  func(lg *types.Log) error
		topics  []ethcommon.Hash
		data    []byte
		wantErr string
	}{
		{"round_initialized topics", func(lg *types.Log) error { _, err := h.decodeRoundInitialized(lg, elog); return err }, nil, nil, "topics"},
		{"round_initialized overflow", func(lg *types.Log) error { _, err := h.decodeRoundInitialized(lg, elog); return err },
			[]ethcommon.Hash{{}, overflowTopic}, nil, "roundId exceeds int64"},

		{"round_funded topics", func(lg *types.Log) error { _, err := h.decodeRoundFunded(lg, elog); return err }, nil, nil, "topics"},
		{"round_funded bad data", func(lg *types.Log) error { _, err := h.decodeRoundFunded(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2)}, []byte{0xde}, "length"},
		{"round_funded overflow", func(lg *types.Log) error { _, err := h.decodeRoundFunded(lg, elog); return err },
			[]ethcommon.Hash{{}, overflowTopic, topic(2)},
			packEventData(t, h, false, "RoundFunded", bigOne, bigOne, bigOne, [8]*big.Int{bigOne, bigOne, bigOne, bigOne, bigOne, bigOne, bigOne, bigOne}),
			"roundId exceeds int64"},

		{"trade topics", func(lg *types.Log) error { _, err := h.decodeTrade(lg, elog); return err }, nil, nil, "topics"},
		{"trade bad data", func(lg *types.Log) error { _, err := h.decodeTrade(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2), topic(0)}, []byte{0xde}, "length"},
		{"trade round overflow", func(lg *types.Log) error { _, err := h.decodeTrade(lg, elog); return err },
			[]ethcommon.Hash{{}, overflowTopic, topic(2), topic(0)},
			packEventData(t, h, false, "Trade", true, true, bigOne, bigOne), "roundId exceeds int64"},
		{"trade question overflow", func(lg *types.Log) error { _, err := h.decodeTrade(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2), overflowTopic},
			packEventData(t, h, false, "Trade", true, true, bigOne, bigOne), "question exceeds uint8"},

		{"split topics", func(lg *types.Log) error { _, err := h.decodePositionSplit(lg, elog); return err }, nil, nil, "topics"},
		{"split bad data", func(lg *types.Log) error { _, err := h.decodePositionSplit(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2), topic(0)}, []byte{0xde}, "length"},
		{"split round overflow", func(lg *types.Log) error { _, err := h.decodePositionSplit(lg, elog); return err },
			[]ethcommon.Hash{{}, overflowTopic, topic(2), topic(0)},
			packEventData(t, h, false, "PositionSplit", bigOne), "roundId exceeds int64"},
		{"split question overflow", func(lg *types.Log) error { _, err := h.decodePositionSplit(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2), overflowTopic},
			packEventData(t, h, false, "PositionSplit", bigOne), "question exceeds uint8"},

		{"merged topics", func(lg *types.Log) error { _, err := h.decodePositionsMerged(lg, elog); return err }, nil, nil, "topics"},
		{"merged bad data", func(lg *types.Log) error { _, err := h.decodePositionsMerged(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2)}, []byte{0xde}, "length"},
		{"merged round overflow", func(lg *types.Log) error { _, err := h.decodePositionsMerged(lg, elog); return err },
			[]ethcommon.Hash{{}, overflowTopic, topic(2)},
			packEventData(t, h, false, "PositionsMerged", [8]*big.Int{bigOne, bigOne, bigOne, bigOne, bigOne, bigOne, bigOne, bigOne}, bigOne),
			"roundId exceeds int64"},

		{"settled topics", func(lg *types.Log) error { _, err := h.decodeRoundSettled(lg, elog); return err }, nil, nil, "topics"},
		{"settled bad data", func(lg *types.Log) error { _, err := h.decodeRoundSettled(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1)}, []byte{0xde}, "length"},
		{"settled round overflow", func(lg *types.Log) error { _, err := h.decodeRoundSettled(lg, elog); return err },
			[]ethcommon.Hash{{}, overflowTopic},
			packEventData(t, h, false, "RoundSettled", uint8(1), [8]*big.Int{bigOne, bigOne, bigOne, bigOne, bigOne, bigOne, bigOne, bigOne}, bigOne),
			"roundId exceeds int64"},

		{"claimed topics", func(lg *types.Log) error { _, err := h.decodeClaimed(lg, elog); return err }, nil, nil, "topics"},
		{"claimed bad data", func(lg *types.Log) error { _, err := h.decodeClaimed(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2)}, []byte{0xde}, "length"},
		{"claimed round overflow", func(lg *types.Log) error { _, err := h.decodeClaimed(lg, elog); return err },
			[]ethcommon.Hash{{}, overflowTopic, topic(2)},
			packEventData(t, h, false, "Claimed", [8]*big.Int{bigOne, bigOne, bigOne, bigOne, bigOne, bigOne, bigOne, bigOne}, bigOne),
			"roundId exceeds int64"},

		{"answers topics", func(lg *types.Log) error { _, err := h.decodeAnswersObserved(lg, elog); return err }, nil, nil, "topics"},
		{"answers bad data", func(lg *types.Log) error { _, err := h.decodeAnswersObserved(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1)}, []byte{0xde}, "length"},
		{"answers round overflow", func(lg *types.Log) error { _, err := h.decodeAnswersObserved(lg, elog); return err },
			[]ethcommon.Hash{{}, overflowTopic},
			packEventData(t, h, false, "AnswersObserved", uint8(1), uint8(1)), "roundId exceeds int64"},

		{"escrowed topics", func(lg *types.Log) error { _, err := h.decodeCycleEscrowed(lg, elog); return err }, nil, nil, "topics"},
		{"escrowed bad data", func(lg *types.Log) error { _, err := h.decodeCycleEscrowed(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1)}, []byte{0xde}, "length"},
		{"escrowed round overflow", func(lg *types.Log) error { _, err := h.decodeCycleEscrowed(lg, elog); return err },
			[]ethcommon.Hash{{}, overflowTopic},
			packEventData(t, h, true, "CycleEscrowed", [8]*big.Int{bigOne, bigOne, bigOne, bigOne, bigOne, bigOne, bigOne, bigOne}),
			"roundId exceeds int64"},

		{"paid topics", func(lg *types.Log) error { _, err := h.decodePaid(lg, elog); return err }, nil, nil, "topics"},
		{"paid bad data", func(lg *types.Log) error { _, err := h.decodePaid(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(0), topic(2)}, []byte{0xde}, "length"},
		{"paid round overflow", func(lg *types.Log) error { _, err := h.decodePaid(lg, elog); return err },
			[]ethcommon.Hash{{}, overflowTopic, topic(0), topic(2)},
			packEventData(t, h, true, "Paid", bigOne), "roundId exceeds int64"},
		{"paid question overflow", func(lg *types.Log) error { _, err := h.decodePaid(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), overflowTopic, topic(2)},
			packEventData(t, h, true, "Paid", bigOne), "question exceeds uint8"},

		{"transfer_single topics", func(lg *types.Log) error { _, err := h.decodeTransferSingle(lg, elog); return err }, nil, nil, "topics"},
		{"transfer_single bad data", func(lg *types.Log) error { _, err := h.decodeTransferSingle(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2), topic(3)}, []byte{0xde}, "length"},
		{"transfer_single position overflow", func(lg *types.Log) error { _, err := h.decodeTransferSingle(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2), topic(3)},
			packEventData(t, h, false, "TransferSingle", hugePositionID, bigOne), "roundId exceeds int64"},

		{"transfer_batch topics", func(lg *types.Log) error { _, err := h.decodeTransferBatch(lg, elog); return err }, nil, nil, "topics"},
		{"transfer_batch bad data", func(lg *types.Log) error { _, err := h.decodeTransferBatch(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2), topic(3)}, []byte{0xde}, "length"},
		{"transfer_batch length mismatch", func(lg *types.Log) error { _, err := h.decodeTransferBatch(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2), topic(3)},
			packEventData(t, h, false, "TransferBatch", []*big.Int{bigOne}, []*big.Int{bigOne, bigOne}),
			"length mismatch"},
		{"transfer_batch position overflow", func(lg *types.Log) error { _, err := h.decodeTransferBatch(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2), topic(3)},
			packEventData(t, h, false, "TransferBatch", []*big.Int{hugePositionID}, []*big.Int{bigOne}),
			"roundId exceeds int64"},

		{"approval topics", func(lg *types.Log) error { _, err := h.decodeApprovalForAll(lg, elog); return err }, nil, nil, "topics"},
		{"approval bad data", func(lg *types.Log) error { _, err := h.decodeApprovalForAll(lg, elog); return err },
			[]ethcommon.Hash{{}, topic(1), topic(2)}, []byte{0xde}, "length"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			lg := &types.Log{Topics: tc.topics, Data: tc.data}
			err := tc.decode(lg)
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("decode err = %v, want substring %q", err, tc.wantErr)
			}
		})
	}
}
