//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"math"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

// TestResolveAdminEventValuesFixture resolves the admin events actually
// present in the fixture set (charity percentage, activation time) end to end.
func TestResolveAdminEventValuesFixture(t *testing.T) {
	r := repo(t)
	golden(t, "resolved_admin_events_fixture", func() any {
		events, err := r.AdminEventsInRange(context.Background(), 0, math.MaxInt64)
		if err != nil {
			t.Fatalf("AdminEventsInRange: %v", err)
		}
		if err := r.ResolveAdminEventValues(context.Background(), events); err != nil {
			t.Fatalf("ResolveAdminEventValues: %v", err)
		}
		return events
	})
}

// TestResolveAdminEventValuesSQLPaths drives every SQL-backed resolver branch
// with synthetic events anchored after the last fixture evtlog (id 6000):
// the activation-span and last-ETH-price lookups hit real fixture rows, while
// the cg_adm_prize_microsec / cg_adm_cst_auclen lookups exercise the
// documented fallback shapes (those tables are empty in the fixture set).
func TestResolveAdminEventValuesSQLPaths(t *testing.T) {
	r := repo(t)

	mk := func(recordType, intValue int64) cgmodel.CGAdminEvent {
		return cgmodel.CGAdminEvent{RecordType: recordType, EvtLogId: 6000, IntegerValue: intValue}
	}
	events := []cgmodel.CGAdminEvent{
		mk(7, 90),                                // DelayDuration: plain seconds
		mk(18, 10100),                            // TimeIncrease: fallback (no microsec base) => x1.01 per bid
		mk(19, 3600),                             // TimeoutClaimPrize: 1h
		mk(20, 4),                                // PriceIncrease: percent-from-divisor
		mk(21, 1_500_000),                        // PrizeTimeIncrementMicros: 1.5 sec
		mk(22, 101),                              // InitialSecondsUntilPrize: fallback pct only
		mk(25, 45),                               // CstAuctionLength as duration
		mk(35, 7200),                             // TimeoutWithdrawPrizes: 2h
		mk(36, ethDutchAuctionDurationNumerator), // EthDutchAuctionDuration: span 500s => "8m 20s"
		mk(37, 10),                               // EthDutchAuctionEndingBidPrice: 0.06/10 ETH
		mk(39, 11),                               // CstDutchAuctionDurationChangeDivisor: fallback ""
		mk(37, 0),                                // guard: non-positive divisor
		{RecordType: 18, EvtLogId: 1, IntegerValue: 10100}, // before any bids/params
	}
	if err := r.ResolveAdminEventValues(context.Background(), events); err != nil {
		t.Fatalf("ResolveAdminEventValues: %v", err)
	}

	resolved := make([]string, len(events))
	for i := range events {
		resolved[i] = events[i].ResolvedValue
	}
	golden(t, "resolved_admin_events_sql_paths", func() any {
		again := make([]cgmodel.CGAdminEvent, len(events))
		copy(again, events)
		for i := range again {
			again[i].ResolvedValue = ""
		}
		if err := r.ResolveAdminEventValues(context.Background(), again); err != nil {
			t.Fatalf("ResolveAdminEventValues (again): %v", err)
		}
		out := make([]map[string]any, len(again))
		for i, ev := range again {
			out[i] = map[string]any{
				"recordType":    ev.RecordType,
				"evtlogId":      ev.EvtLogId,
				"integerValue":  ev.IntegerValue,
				"resolvedValue": ev.ResolvedValue,
			}
		}
		return out
	})

	// The activation-span resolver must have found the real fixture rows
	// (5067: 1767227800, 5075: 1767228300 => span 500s).
	for i, ev := range events {
		if ev.RecordType == 36 && ev.IntegerValue > 0 && resolved[i] == "" {
			t.Error("expected type-36 event to resolve against fixture activation times")
		}
		if ev.RecordType == 37 && ev.IntegerValue > 0 && resolved[i] == "" {
			t.Error("expected type-37 event to resolve against the fixture ETH bid price")
		}
	}
}

func TestResolveAdminEventValuesWithHistoricalParameters(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	for _, statement := range []string{
		`INSERT INTO cg_adm_prize_microsec
			(evtlog_id,block_num,tx_id,time_stamp,contract_aid,new_microseconds)
			VALUES (5094,140,1040,'2026-01-01T01:00:00Z',2,2000000)`,
		`INSERT INTO cg_adm_cst_auclen
			(evtlog_id,block_num,tx_id,time_stamp,contract_aid,new_len)
			VALUES (5094,140,1040,'2026-01-01T01:00:00Z',2,3600)`,
	} {
		if _, err := r.pool().Exec(ctx, statement); err != nil {
			t.Fatalf("seed historical parameter: %v", err)
		}
	}
	t.Cleanup(func() {
		if _, err := r.pool().Exec(context.Background(),
			"DELETE FROM cg_adm_prize_microsec WHERE evtlog_id=5094"); err != nil {
			t.Errorf("clean prize microseconds: %v", err)
		}
		if _, err := r.pool().Exec(context.Background(),
			"DELETE FROM cg_adm_cst_auclen WHERE evtlog_id=5094"); err != nil {
			t.Errorf("clean CST auction length: %v", err)
		}
	})

	events := []cgmodel.CGAdminEvent{
		{RecordType: 18, EvtLogId: 6000, IntegerValue: 200},
		{RecordType: 22, EvtLogId: 6000, IntegerValue: 2},
		{RecordType: 39, EvtLogId: 6000, IntegerValue: 10},
	}
	if err := r.ResolveAdminEventValues(ctx, events); err != nil {
		t.Fatal(err)
	}
	want := []string{
		"4 sec after next bid (×2.00)",
		"297h 30m (50.0000%)",
		"6m change per bid",
	}
	for i := range events {
		if events[i].ResolvedValue != want[i] {
			t.Errorf("event %d resolved = %q, want %q", i, events[i].ResolvedValue, want[i])
		}
	}

	if value, ok, err := r.latestInt64ParamAtEvtlog(
		ctx, "cg_adm_prize_microsec", "new_microseconds", 5094, false,
	); err != nil || ok || value != 0 {
		t.Fatalf("exclusive historical lookup = %d,%v,%v", value, ok, err)
	}
}

func TestAdminEventResolversPropagateCancellation(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	events := []cgmodel.CGAdminEvent{
		{RecordType: 18, EvtLogId: 6000, IntegerValue: 200},
		{RecordType: 22, EvtLogId: 6000, IntegerValue: 2},
		{RecordType: 36, EvtLogId: 6000, IntegerValue: 2},
		{RecordType: 37, EvtLogId: 6000, IntegerValue: 2},
		{RecordType: 39, EvtLogId: 6000, IntegerValue: 2},
	}
	for _, event := range events {
		if err := r.ResolveAdminEventValues(ctx, []cgmodel.CGAdminEvent{event}); !errors.Is(err, context.Canceled) {
			t.Errorf("record type %d cancellation = %v", event.RecordType, err)
		}
	}
}
