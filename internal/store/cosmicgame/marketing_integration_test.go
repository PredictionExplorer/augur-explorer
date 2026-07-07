//go:build integration

package cosmicgame

import (
	"context"
	"testing"
)

func TestMarketingRewardHistoryGlobal(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "marketing_reward_history_global", func() any {
		recs, err := r.MarketingRewardHistoryGlobal(ctx, 0, 100)
		if err != nil {
			t.Fatalf("MarketingRewardHistoryGlobal: %v", err)
		}
		return recs
	})
	golden(t, "marketing_reward_history_global_paged", func() any {
		recs, err := r.MarketingRewardHistoryGlobal(ctx, 1, 1)
		if err != nil {
			t.Fatalf("MarketingRewardHistoryGlobal paged: %v", err)
		}
		return recs
	})
}

func TestMarketingRewardsByUser(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "marketing_rewards_by_user", func() any {
		recs, err := r.MarketingRewardsByUser(ctx, aidBob)
		if err != nil {
			t.Fatalf("MarketingRewardsByUser: %v", err)
		}
		return recs
	})
	got, err := r.MarketingRewardsByUser(ctx, aidZero)
	if err != nil {
		t.Fatalf("MarketingRewardsByUser(zero addr): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no marketing rewards for the zero address, got %d", len(got))
	}
	if got == nil {
		t.Error("empty result must be a non-nil slice (JSON [] parity)")
	}
}
