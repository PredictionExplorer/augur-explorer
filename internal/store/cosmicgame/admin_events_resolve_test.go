package cosmicgame

import (
	"testing"
)

func TestAdminEventValueFormatters(t *testing.T) {
	t.Parallel()
	percentCases := map[int64]string{
		0: "", -1: "", 4: "25.0000%", 100: "1.0000%",
	}
	for input, want := range percentCases {
		if got := formatPercentFromDivisor(input); got != want {
			t.Errorf("formatPercentFromDivisor(%d) = %q, want %q", input, got, want)
		}
	}

	microsecondCases := map[int64]string{
		0: "0 sec", -1: "0 sec", 1_000_000: "1 sec", 1_500_001: "1.500001 sec",
	}
	for input, want := range microsecondCases {
		if got := formatMicrosecondsAsSeconds(input); got != want {
			t.Errorf("formatMicrosecondsAsSeconds(%d) = %q, want %q", input, got, want)
		}
	}

	durationCases := map[int64]string{
		0: "0 sec", -1: "0 sec", 119: "119 sec", 120: "2m",
		121: "2m 1s", 3600: "1h", 3660: "1h 1m", 3661: "1h 1m 1s",
	}
	for input, want := range durationCases {
		if got := formatDurationSeconds(input); got != want {
			t.Errorf("formatDurationSeconds(%d) = %q, want %q", input, got, want)
		}
	}

	ethCases := map[int64]string{
		0: "0 ETH", -1: "0 ETH", 1_000_000_000_000_000_000: "1.00000000 ETH",
		10_000_000_000_000: "0.000010000000 ETH",
	}
	for input, want := range ethCases {
		if got := formatEthFromWei(input); got != want {
			t.Errorf("formatEthFromWei(%d) = %q, want %q", input, got, want)
		}
	}
}

func TestLatestInt64ParamRejectsIdentifiersBeforeDatabaseAccess(t *testing.T) {
	t.Parallel()
	var repo Repo
	for _, test := range []struct {
		table  string
		column string
	}{
		{table: "bad-table", column: "value"},
		{table: "table", column: "bad column"},
		{table: "", column: "value"},
	} {
		if _, _, err := repo.latestInt64ParamAtEvtlog(
			t.Context(), test.table, test.column, 1, false,
		); err == nil {
			t.Errorf("identifiers %+v were accepted", test)
		}
	}
}

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
