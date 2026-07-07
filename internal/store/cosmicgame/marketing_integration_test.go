//go:build integration

package cosmicgame

import "testing"

func TestGetMarketingRewardHistoryGlobal(t *testing.T) {
	sw := store(t)
	golden(t, "marketing_reward_history_global", func() any {
		return sw.Get_marketing_reward_history_global(0, 100)
	})
	golden(t, "marketing_reward_history_global_paged", func() any {
		return sw.Get_marketing_reward_history_global(1, 1)
	})
}

func TestGetMarketingRewardsByUser(t *testing.T) {
	sw := store(t)
	golden(t, "marketing_rewards_by_user", func() any {
		return sw.Get_marketing_rewards_by_user(aidBob)
	})
	if got := sw.Get_marketing_rewards_by_user(aidZero); len(got) != 0 {
		t.Errorf("expected no marketing rewards for the zero address, got %d", len(got))
	}
}
