package store

import (
	"testing"
	"time"
)

func TestTimeText(t *testing.T) {
	var got string
	ts := time.Date(2026, 7, 7, 12, 30, 45, 0, time.UTC)
	if err := TimeText(&got).Scan(ts); err != nil {
		t.Fatalf("Scan(time.Time): %v", err)
	}
	// Must match database/sql's convertAssign formatting (RFC3339Nano):
	// no fractional seconds when zero, "Z" for UTC.
	if want := "2026-07-07T12:30:45Z"; got != want {
		t.Errorf("TimeText = %q, want %q", got, want)
	}

	withNanos := time.Date(2026, 7, 7, 12, 30, 45, 123456789, time.UTC)
	if err := TimeText(&got).Scan(withNanos); err != nil {
		t.Fatalf("Scan(time.Time with nanos): %v", err)
	}
	if want := "2026-07-07T12:30:45.123456789Z"; got != want {
		t.Errorf("TimeText = %q, want %q", got, want)
	}

	if err := TimeText(&got).Scan(nil); err == nil {
		t.Error("Scan(nil) succeeded, want error (legacy layer also failed on NULL)")
	}
	if err := TimeText(&got).Scan("2026-01-01"); err == nil {
		t.Error("Scan(string) succeeded, want error")
	}
}

func TestNullTimeText(t *testing.T) {
	got := "untouched"
	if err := NullTimeText(&got).Scan(nil); err != nil {
		t.Fatalf("Scan(nil): %v", err)
	}
	if got != "untouched" {
		t.Errorf("Scan(nil) modified dst to %q, want it left unchanged", got)
	}

	ts := time.Date(2026, 7, 7, 12, 30, 45, 0, time.UTC)
	if err := NullTimeText(&got).Scan(ts); err != nil {
		t.Fatalf("Scan(time.Time): %v", err)
	}
	if want := "2026-07-07T12:30:45Z"; got != want {
		t.Errorf("NullTimeText = %q, want %q", got, want)
	}

	if err := NullTimeText(&got).Scan(42); err == nil {
		t.Error("Scan(int) succeeded, want error")
	}
}
