//go:build integration

package cosmicgame

import "testing"

func TestGetGlobStatsCstRewardForBidding(t *testing.T) {
	sw := store(t)
	reward, err := sw.Get_glob_stats_cst_reward_for_bidding()
	if err != nil {
		t.Fatalf("Get_glob_stats_cst_reward_for_bidding: %v", err)
	}
	golden(t, "glob_stats_cst_reward_for_bidding", func() any {
		got, err := sw.Get_glob_stats_cst_reward_for_bidding()
		if err != nil {
			t.Fatalf("Get_glob_stats_cst_reward_for_bidding: %v", err)
		}
		return got
	})
	if reward == "" {
		t.Error("expected a non-empty cst_reward_for_bidding value")
	}
}

func TestGetLatestDecimalParam(t *testing.T) {
	sw := store(t)

	val, hasRow, err := sw.Get_latest_decimal_param("cg_adm_charity_pcent", "percentage")
	if err != nil {
		t.Fatalf("Get_latest_decimal_param(charity_pcent): %v", err)
	}
	if !hasRow {
		t.Fatal("expected cg_adm_charity_pcent to have a seeded row")
	}
	if val != "10" {
		t.Errorf("latest charity percentage: got %q, want %q", val, "10")
	}

	// Table exists but has no rows: reported as absent, not an error.
	_, hasRow, err = sw.Get_latest_decimal_param("cg_adm_msg_len", "new_length")
	if err != nil {
		t.Fatalf("Get_latest_decimal_param(msg_len): %v", err)
	}
	if hasRow {
		t.Error("expected cg_adm_msg_len to be empty in the fixture set")
	}

	// This function returns errors (no os.Exit): pin the error path too.
	if _, _, err := sw.Get_latest_decimal_param("cg_no_such_table", "x"); err == nil {
		t.Error("expected an error for a nonexistent table")
	}
}
