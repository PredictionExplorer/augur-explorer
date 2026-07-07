//go:build integration

package cosmicgame

import (
	"math"
	"testing"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

// TestResolveAdminEventValuesFixture resolves the admin events actually
// present in the fixture set (charity percentage, activation time) end to end.
func TestResolveAdminEventValuesFixture(t *testing.T) {
	sw := store(t)
	golden(t, "resolved_admin_events_fixture", func() any {
		events := sw.Get_admin_events_in_range(0, math.MaxInt64)
		sw.Resolve_admin_event_values(events)
		return events
	})
}

// TestResolveAdminEventValuesSQLPaths drives every SQL-backed resolver branch
// with synthetic events anchored after the last fixture evtlog (id 6000):
// the activation-span and last-ETH-price lookups hit real fixture rows, while
// the cg_adm_prize_microsec / cg_adm_cst_auclen lookups exercise the
// documented fallback shapes (those tables are empty in the fixture set).
func TestResolveAdminEventValuesSQLPaths(t *testing.T) {
	sw := store(t)

	mk := func(recordType, intValue int64) p.CGAdminEvent {
		return p.CGAdminEvent{RecordType: recordType, EvtLogId: 6000, IntegerValue: intValue}
	}
	events := []p.CGAdminEvent{
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
	sw.Resolve_admin_event_values(events)

	resolved := make([]string, len(events))
	for i := range events {
		resolved[i] = events[i].ResolvedValue
	}
	golden(t, "resolved_admin_events_sql_paths", func() any {
		again := make([]p.CGAdminEvent, len(events))
		copy(again, events)
		for i := range again {
			again[i].ResolvedValue = ""
		}
		sw.Resolve_admin_event_values(again)
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
