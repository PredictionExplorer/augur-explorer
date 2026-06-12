package cosmicgame

import "testing"

func TestFormatPercentFromDivisor(t *testing.T) {
	got := formatPercentFromDivisor(101)
	if got != "0.9901%" {
		t.Fatalf("got %q want 0.9901%%", got)
	}
	got = formatPercentFromDivisor(202)
	if got != "0.4950%" {
		t.Fatalf("got %q want 0.4950%%", got)
	}
}

func TestFormatDurationSeconds(t *testing.T) {
	if got := formatDurationSeconds(47268); got != "13h 7m 48s" {
		t.Fatalf("got %q", got)
	}
	if got := formatDurationSeconds(150); got != "2m 30s" {
		t.Fatalf("got %q", got)
	}
}

func TestEthDutchAuctionDurationFormula(t *testing.T) {
	span := int64(86399)
	divisor := int64(21041)
	secs := span * ethDutchAuctionDurationNumerator / divisor
	if secs < 47680 || secs > 47690 {
		t.Fatalf("expected ~47684 sec, got %d", secs)
	}
}

func TestInitialDurationFormula(t *testing.T) {
	micros := int64(918090000)
	divisor := int64(1000000)
	raw := micros * initialDurationRawNumerator / 1000
	secs := raw / divisor
	if secs < 980 || secs > 990 {
		t.Fatalf("expected ~983 sec, got %d", secs)
	}
}
