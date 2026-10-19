package main

import (
	"strings"
	"testing"
)

func TestAdminEventHelp(t *testing.T) {
	// Record types arrive as jsonNum from decoded JSON.
	if got := adminEventHelp(jsonNum(21)); !strings.Contains(got, "main prize") {
		t.Errorf("adminEventHelp(jsonNum(21)) = %q, want main-prize-timer text", got)
	}
	if got := adminEventHelp(float64(7)); got == "" {
		t.Error("adminEventHelp(float64(7)) must return help text")
	}
	if got := adminEventHelp(9); got == "" {
		t.Error("adminEventHelp(int 9) must return help text")
	}
	// Unknown or retired types render no balloon.
	if got := adminEventHelp(jsonNum(0)); got != "" {
		t.Errorf("adminEventHelp(0) = %q, want empty", got)
	}
	if got := adminEventHelp(jsonNum(43)); !strings.Contains(got, "decline") {
		t.Errorf("adminEventHelp(43) = %q, want the decline-rate text", got)
	}
	if got := adminEventHelp(jsonNum(45)); !strings.Contains(got, "divisor") {
		t.Errorf("adminEventHelp(45) = %q, want the adjustment-divisor text", got)
	}
	if got := adminEventHelp("21"); got != "" {
		t.Errorf("adminEventHelp(string) = %q, want empty", got)
	}
}

func TestRoundRangeLabel(t *testing.T) {
	if got := roundRangeLabel(jsonNum(10), jsonNum(10)); got != "Round 10" {
		t.Errorf("same round = %q", got)
	}
	if got := roundRangeLabel(jsonNum(0), jsonNum(0)); got != "Round 0" {
		t.Errorf("round zero = %q", got)
	}
	if got := roundRangeLabel(jsonNum(4), jsonNum(7)); got != "Rounds 4 to 7" {
		t.Errorf("round span = %q", got)
	}
	// Older apiserver builds omit the fields; the label must vanish
	// instead of breaking the page.
	if got := roundRangeLabel(nil, nil); got != "" {
		t.Errorf("missing fields = %q, want empty", got)
	}
}

func TestAdminEventsPageRendersBalloon(t *testing.T) {
	tmpl, err := loadTemplates()
	if err != nil {
		t.Fatalf("loadTemplates: %v", err)
	}
	// RoundNumStart/RoundNumEnd are deliberately absent: an older apiserver
	// build omits them and the page must still render (without the label).
	data := map[string]any{
		"EvtLogIdStart": jsonNum(1), "EvtLogIdEnd": jsonNum(2),
		"AdminEvents": []any{
			map[string]any{
				"RecordType": jsonNum(21),
				"TxHash":     "0xabc", "BlockNum": jsonNum(1),
				"DateTime": "t", "IntegerValue": jsonNum(3600),
				"FloatValue": nil, "AddressValue": "",
				"StringValue": "", "ResolvedValue": "1 h",
			},
		},
	}
	var out strings.Builder
	name := "cosmicsignature/cg_system_admin_events_in_range.html"
	if err := tmpl.ExecuteTemplate(&out, name, data); err != nil {
		t.Fatalf("%s: %v", name, err)
	}
	body := out.String()
	if !strings.Contains(body, "balloon-text") {
		t.Error("rendered page is missing the help balloon")
	}
	if !strings.Contains(body, adminEventHelpTexts[21]) {
		t.Error("rendered page is missing the record-type 21 help text")
	}
}
