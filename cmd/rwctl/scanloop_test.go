package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// fakeScanClient scripts the logScanClient surface.
type fakeScanClient struct {
	tip       int64
	headerErr error
	filterErr error
	// logs holds the answer per FilterLogs call, keyed by FromBlock.
	logs    map[int64][]types.Log
	queries []ethereum.FilterQuery
}

func (f *fakeScanClient) HeaderByNumber(_ context.Context, _ *big.Int) (*types.Header, error) {
	if f.headerErr != nil {
		return nil, f.headerErr
	}
	return &types.Header{Number: big.NewInt(f.tip)}, nil
}

func (f *fakeScanClient) FilterLogs(_ context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	f.queries = append(f.queries, q)
	if f.filterErr != nil {
		return nil, f.filterErr
	}
	return f.logs[q.FromBlock.Int64()], nil
}

func TestScanLogsByRangeWalksRangesAndSkipsRemoved(t *testing.T) {
	addr := common.HexToAddress("0xaa00000000000000000000000000000000000001")
	topic := common.HexToHash("0x01")
	client := &fakeScanClient{
		tip: 250,
		logs: map[int64][]types.Log{
			0:   {{BlockNumber: 5}, {BlockNumber: 7, Removed: true}},
			100: {{BlockNumber: 150}},
			200: {},
		},
	}

	var handled []uint64
	var ranges []string
	err := scanLogsByRange(context.Background(), client, addr, topic, 0, 100,
		func(from, to int64) { ranges = append(ranges, fmt.Sprintf("%d-%d", from, to)) },
		func(lg *types.Log) error {
			handled = append(handled, lg.BlockNumber)
			return nil
		})
	if err != nil {
		t.Fatalf("scanLogsByRange: %v", err)
	}
	if len(handled) != 2 || handled[0] != 5 || handled[1] != 150 {
		t.Errorf("handled = %v, want [5 150] (removed log skipped)", handled)
	}
	if len(ranges) != 3 || ranges[0] != "0-100" || ranges[2] != "200-300" {
		t.Errorf("ranges = %v, want three 100-block steps", ranges)
	}
	// Filter carries the topic + address on every query.
	for _, q := range client.queries {
		if len(q.Addresses) != 1 || q.Addresses[0] != addr {
			t.Errorf("query addresses = %v", q.Addresses)
		}
		if len(q.Topics) != 1 || len(q.Topics[0]) != 1 || q.Topics[0][0] != topic {
			t.Errorf("query topics = %v", q.Topics)
		}
	}
}

func TestScanLogsByRangeStartBlockOffset(t *testing.T) {
	client := &fakeScanClient{tip: 10, logs: map[int64][]types.Log{}}
	err := scanLogsByRange(context.Background(), client, common.Address{}, common.Hash{}, 8, 5, nil,
		func(*types.Log) error { return nil })
	if err != nil {
		t.Fatalf("scanLogsByRange: %v", err)
	}
	if len(client.queries) != 1 || client.queries[0].FromBlock.Int64() != 8 {
		t.Errorf("queries = %+v, want a single range starting at 8", client.queries)
	}
}

func TestScanLogsByRangeErrors(t *testing.T) {
	t.Run("header error", func(t *testing.T) {
		client := &fakeScanClient{headerErr: errors.New("node down")}
		err := scanLogsByRange(context.Background(), client, common.Address{}, common.Hash{}, 0, 10, nil,
			func(*types.Log) error { return nil })
		if err == nil || !errors.Is(err, client.headerErr) {
			t.Errorf("err = %v, want header failure", err)
		}
	})
	t.Run("filter error", func(t *testing.T) {
		client := &fakeScanClient{tip: 5, filterErr: errors.New("query too wide")}
		err := scanLogsByRange(context.Background(), client, common.Address{}, common.Hash{}, 0, 10, nil,
			func(*types.Log) error { return nil })
		if err == nil || !errors.Is(err, client.filterErr) {
			t.Errorf("err = %v, want filter failure", err)
		}
	})
	t.Run("handler error stops the scan", func(t *testing.T) {
		client := &fakeScanClient{tip: 100, logs: map[int64][]types.Log{
			0:  {{BlockNumber: 1}},
			50: {{BlockNumber: 60}},
		}}
		boom := errors.New("db exploded")
		calls := 0
		err := scanLogsByRange(context.Background(), client, common.Address{}, common.Hash{}, 0, 50, nil,
			func(*types.Log) error { calls++; return boom })
		if !errors.Is(err, boom) {
			t.Errorf("err = %v, want handler error", err)
		}
		if calls != 1 {
			t.Errorf("handler calls = %d, want scan aborted after the first failure", calls)
		}
	})
}
