//go:build integration

package cosmicgame

import (
	"context"
	"math"
	"testing"
)

func TestSystemModeChanges(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "system_mode_change_event_list", func() any {
		recs, err := r.SystemModeChanges(ctx, 0, 100)
		if err != nil {
			t.Fatalf("SystemModeChanges: %v", err)
		}
		return recs
	})
	// offset=-1 appends the synthetic "Deployment" row.
	golden(t, "system_mode_change_event_list_deployment", func() any {
		recs, err := r.SystemModeChanges(ctx, -1, 100)
		if err != nil {
			t.Fatalf("SystemModeChanges(deployment): %v", err)
		}
		return recs
	})
}

func TestAdminEventsInRange(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "admin_events_in_range_all", func() any {
		recs, err := r.AdminEventsInRange(ctx, 0, math.MaxInt64)
		if err != nil {
			t.Fatalf("AdminEventsInRange: %v", err)
		}
		return recs
	})
	got, err := r.AdminEventsInRange(ctx, 0, 1)
	if err != nil {
		t.Fatalf("AdminEventsInRange(empty range): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected empty admin event range below first evtlog, got %d records", len(got))
	}
}

// TestAdminEventsQueryCoversEveryBranch guards the branch registry: every
// live record type must appear in the generated UNION exactly once, so a
// registry edit can never silently drop an admin event type from the API.
// Retired codes are listed explicitly and must stay absent — their events
// cannot be emitted any more and their meaning must not be recycled.
func TestAdminEventsQueryCoversEveryBranch(t *testing.T) {
	const highestRecordType = 46
	// 43 and 45 held the CstBidPriceDecline* events that migration 00030
	// reverted along with the contract change.
	retired := map[int]bool{43: true, 45: true}

	seen := make(map[int]bool, len(adminEventBranches))
	for _, b := range adminEventBranches {
		if seen[b.recordType] {
			t.Errorf("record type %d listed twice", b.recordType)
		}
		if retired[b.recordType] {
			t.Errorf("record type %d is retired and must not be reused", b.recordType)
		}
		seen[b.recordType] = true
	}
	for want := 1; want <= highestRecordType; want++ {
		if !seen[want] && !retired[want] {
			t.Errorf("record type %d missing from adminEventBranches", want)
		}
	}
	if want := highestRecordType - len(retired); len(adminEventBranches) != want {
		t.Errorf("expected %d branches, got %d", want, len(adminEventBranches))
	}
}

// TestAdminEventsQueryValidSQL proves the assembled UNION is executable
// against the real schema even for event tables the fixture set leaves empty
// (a wrong table or column name in the registry would fail here, not in
// production).
func TestAdminEventsQueryValidSQL(t *testing.T) {
	r := repo(t)
	rows, err := r.pool().Query(context.Background(), adminEventsQuery(), 0, math.MaxInt64)
	if err != nil {
		t.Fatalf("admin events query rejected by PostgreSQL: %v", err)
	}
	rows.Close()
	if rows.Err() != nil {
		t.Fatalf("admin events query failed: %v", rows.Err())
	}
}
