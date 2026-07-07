package testutil

import (
	"encoding/json"
	"strings"
	"testing"
)

func row(pairs ...any) Row {
	r := make(Row)
	for i := 0; i+1 < len(pairs); i += 2 {
		r[pairs[i].(string)] = pairs[i+1]
	}
	return r
}

func TestDiffSnapshotsEmptyForIdenticalStates(t *testing.T) {
	snap := Snapshot{"cg_bid": {row("round_num", json.Number("1"))}}
	diff, err := DiffSnapshots(snap, snap)
	if err != nil {
		t.Fatalf("DiffSnapshots: %v", err)
	}
	if strings.TrimSpace(string(diff)) != "{}" {
		t.Errorf("diff of identical snapshots = %s, want {}", diff)
	}
}

func TestDiffSnapshotsReportsAddsAndRemoves(t *testing.T) {
	before := Snapshot{
		"stats": {row("num_bids", json.Number("5"))},
		"gone":  {row("x", json.Number("1"))},
	}
	after := Snapshot{
		"stats": {row("num_bids", json.Number("6"))},
		"fresh": {row("y", json.Number("2"))},
	}
	rendered, err := DiffSnapshots(before, after)
	if err != nil {
		t.Fatalf("DiffSnapshots: %v", err)
	}
	var diff map[string]TableDiff
	if err := json.Unmarshal(rendered, &diff); err != nil {
		t.Fatalf("diff output is not valid JSON: %v\n%s", err, rendered)
	}
	if len(diff["stats"].Added) != 1 || len(diff["stats"].Removed) != 1 {
		t.Errorf("stats diff = %+v, want one add and one remove (update)", diff["stats"])
	}
	if len(diff["gone"].Removed) != 1 || len(diff["gone"].Added) != 0 {
		t.Errorf("gone diff = %+v, want one removal", diff["gone"])
	}
	if len(diff["fresh"].Added) != 1 {
		t.Errorf("fresh diff = %+v, want one addition", diff["fresh"])
	}
}

func TestDiffSnapshotsMultisetSemantics(t *testing.T) {
	dup := row("v", json.Number("1"))
	before := Snapshot{"t": {dup}}
	after := Snapshot{"t": {dup, dup}}
	rendered, err := DiffSnapshots(before, after)
	if err != nil {
		t.Fatalf("DiffSnapshots: %v", err)
	}
	var diff map[string]TableDiff
	if err := json.Unmarshal(rendered, &diff); err != nil {
		t.Fatalf("unmarshaling diff: %v", err)
	}
	if len(diff["t"].Added) != 1 {
		t.Errorf("duplicate-row diff = %+v, want exactly one addition", diff["t"])
	}
}

func TestDiffSnapshotsDeterministicOrder(t *testing.T) {
	before := Snapshot{}
	after := Snapshot{"t": {
		row("v", json.Number("2")),
		row("v", json.Number("1")),
		row("v", json.Number("3")),
	}}
	first, err := DiffSnapshots(before, after)
	if err != nil {
		t.Fatalf("DiffSnapshots: %v", err)
	}
	// Reversed input order must render identically.
	after2 := Snapshot{"t": {
		row("v", json.Number("3")),
		row("v", json.Number("1")),
		row("v", json.Number("2")),
	}}
	second, err := DiffSnapshots(before, after2)
	if err != nil {
		t.Fatalf("DiffSnapshots: %v", err)
	}
	if string(first) != string(second) {
		t.Errorf("diff not order-independent:\n%s\nvs\n%s", first, second)
	}
}

func TestResolveValueMapsForeignKeys(t *testing.T) {
	res := &resolvers{
		addr:   map[int64]string{7: "0xabc"},
		txHash: map[int64]string{9: "0xtx"},
		evtKey: map[int64]string{3: "evt:100/0"},
		bidKey: map[int64]string{4: "evt:100/1"},
	}
	cases := []struct {
		col  string
		val  any
		want any
	}{
		{"bidder_aid", json.Number("7"), "0xabc"},
		{"aid", json.Number("7"), "0xabc"},
		{"tx_id", json.Number("9"), "0xtx"},
		{"evtlog_id", json.Number("3"), "evt:100/0"},
		{"withdrawal_id", json.Number("3"), "evt:100/0"},
		{"bid_id", json.Number("4"), "evt:100/1"},
		// Unresolvable ids pass through: 0 means "none" in several tables.
		{"evtlog_id", json.Number("0"), json.Number("0")},
		{"withdrawal_id", json.Number("0"), json.Number("0")},
		// Non-FK columns never change even when numeric.
		{"round_num", json.Number("7"), json.Number("7")},
		{"amount", "123", "123"},
	}
	for _, c := range cases {
		if got := resolveValue(c.col, c.val, res); got != c.want {
			t.Errorf("resolveValue(%q, %v) = %v, want %v", c.col, c.val, got, c.want)
		}
	}
}

func TestTransformRowDropsVolatileColumns(t *testing.T) {
	res := &resolvers{addr: map[int64]string{}, txHash: map[int64]string{}, evtKey: map[int64]string{}, bidKey: map[int64]string{}}
	out := transformRow("evt_log", row(
		"id", json.Number("55"),
		"address_id", json.Number("2"),
		"log_rlp", "\\xdeadbeef",
		"round_num", json.Number("1"),
	), res)
	for _, dropped := range []string{"id", "address_id", "log_rlp"} {
		if _, ok := out[dropped]; ok {
			t.Errorf("column %q not dropped: %+v", dropped, out)
		}
	}
	if _, ok := out["round_num"]; !ok {
		t.Errorf("column round_num missing: %+v", out)
	}
}

func TestTransformRowDropsTableScopedColumns(t *testing.T) {
	res := &resolvers{addr: map[int64]string{}, txHash: map[int64]string{}, evtKey: map[int64]string{}, bidKey: map[int64]string{}}
	// address.tx_id is first-seen bookkeeping and dropped only on that table.
	withDrop := transformRow("address", row("tx_id", json.Number("9"), "addr", "0xabc"), res)
	if _, ok := withDrop["tx_id"]; ok {
		t.Errorf("address.tx_id not dropped: %+v", withDrop)
	}
	kept := transformRow("evt_log", row("tx_id", json.Number("9")), res)
	if _, ok := kept["tx_id"]; !ok {
		t.Errorf("evt_log.tx_id wrongly dropped: %+v", kept)
	}
}

func TestCompareGoldenRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/sample.json"
	// Write via the -update path, then compare via the normal path.
	*updateGolden = true
	CompareGolden(t, path, []byte(`{"a":1}`))
	*updateGolden = false
	CompareGolden(t, path, []byte(`{"a":1}`))
}
