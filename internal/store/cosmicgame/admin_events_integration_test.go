//go:build integration

package cosmicgame

import (
	"math"
	"testing"
)

func TestGetSystemModeChangeEventList(t *testing.T) {
	sw := store(t)
	golden(t, "system_mode_change_event_list", func() any {
		return sw.Get_system_mode_change_event_list(0, 100)
	})
	// offset=-1 appends the synthetic "Deployment" row.
	golden(t, "system_mode_change_event_list_deployment", func() any {
		return sw.Get_system_mode_change_event_list(-1, 100)
	})
}

func TestGetAdminEventsInRange(t *testing.T) {
	sw := store(t)
	golden(t, "admin_events_in_range_all", func() any {
		return sw.Get_admin_events_in_range(0, math.MaxInt64)
	})
	if got := sw.Get_admin_events_in_range(0, 1); len(got) != 0 {
		t.Errorf("expected empty admin event range below first evtlog, got %d records", len(got))
	}
}
