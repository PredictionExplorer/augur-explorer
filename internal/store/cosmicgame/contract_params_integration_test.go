//go:build integration

package cosmicgame

import (
	"context"
	"testing"
)

func TestGlobStatsCstRewardForBidding(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	reward, err := r.GlobStatsCstRewardForBidding(ctx)
	if err != nil {
		t.Fatalf("GlobStatsCstRewardForBidding: %v", err)
	}
	golden(t, "glob_stats_cst_reward_for_bidding", func() any {
		got, err := r.GlobStatsCstRewardForBidding(ctx)
		if err != nil {
			t.Fatalf("GlobStatsCstRewardForBidding: %v", err)
		}
		return got
	})
	if reward == "" {
		t.Error("expected a non-empty cst_reward_for_bidding value")
	}
}

func TestLatestDecimalParam(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	val, hasRow, err := r.LatestDecimalParam(ctx, "cg_adm_charity_pcent", "percentage")
	if err != nil {
		t.Fatalf("LatestDecimalParam(charity_pcent): %v", err)
	}
	if !hasRow {
		t.Fatal("expected cg_adm_charity_pcent to have a seeded row")
	}
	if val != "10" {
		t.Errorf("latest charity percentage: got %q, want %q", val, "10")
	}

	// Table exists but has no rows: reported as absent, not an error.
	_, hasRow, err = r.LatestDecimalParam(ctx, "cg_adm_msg_len", "new_length")
	if err != nil {
		t.Fatalf("LatestDecimalParam(msg_len): %v", err)
	}
	if hasRow {
		t.Error("expected cg_adm_msg_len to be empty in the fixture set")
	}

	// This function returns errors (no os.Exit): pin the error path too.
	if _, _, err := r.LatestDecimalParam(ctx, "cg_no_such_table", "x"); err == nil {
		t.Error("expected an error for a nonexistent table")
	}
	// The identifier guard rejects anything that is not a plain lowercase
	// identifier before it can reach the interpolated SQL.
	if _, _, err := r.LatestDecimalParam(ctx, "cg_bid; DROP TABLE cg_bid", "x"); err == nil {
		t.Error("expected an error for a malformed table identifier")
	}
	if _, _, err := r.LatestDecimalParam(ctx, "cg_adm_charity_pcent", `percentage"`); err == nil {
		t.Error("expected an error for a malformed column identifier")
	}
}
