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

// TestAdminCorrectionRoundTrip inserts a correction row through the Repo
// primitives (as the cg-etl startup sync would) and deletes it again,
// asserting LatestDecimalParam sees the new value while it exists. Fixture
// state is restored so test order cannot affect the goldens.
func TestAdminCorrectionRoundTrip(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	before, hasRow, err := r.LatestDecimalParam(ctx, "cg_adm_charity_pcent", "percentage")
	if err != nil || !hasRow {
		t.Fatalf("LatestDecimalParam(before): val=%q hasRow=%v err=%v", before, hasRow, err)
	}

	// The fixture set has evt_log rows; reference one to satisfy the FK and
	// mark the row so cleanup can target it precisely.
	meta := &AdminCorrectionMeta{
		EvtId:       5001,
		BlockNum:    424242,
		TxId:        1001,
		TimeStamp:   1767230000,
		ContractAid: aidCosmicGame,
	}
	if err := r.InsertAdminCorrectionDecimal(ctx, "cg_adm_charity_pcent", "percentage", "25", meta, 0); err != nil {
		t.Fatalf("InsertAdminCorrectionDecimal: %v", err)
	}
	t.Cleanup(func() {
		_, err := r.store.Pool().Exec(context.Background(),
			"DELETE FROM cg_adm_charity_pcent WHERE block_num=$1", meta.BlockNum)
		if err != nil {
			t.Fatalf("cleanup of correction row failed: %v", err)
		}
	})

	got, hasRow, err := r.LatestDecimalParam(ctx, "cg_adm_charity_pcent", "percentage")
	if err != nil || !hasRow {
		t.Fatalf("LatestDecimalParam(after insert): hasRow=%v err=%v", hasRow, err)
	}
	if got != "25" {
		t.Errorf("latest charity percentage after correction: got %q, want %q", got, "25")
	}

	// Malformed identifiers must be rejected before touching SQL.
	if err := r.InsertAdminCorrectionDecimal(ctx, "cg_adm_charity_pcent", "percentage; --", "1", meta, 0); err == nil {
		t.Error("expected an error for a malformed column identifier")
	}
}
