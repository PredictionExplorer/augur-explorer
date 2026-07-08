//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestGetStakeActionCstInfo(t *testing.T) {
	r := repo(t)
	// Action 1: alice staked and later unstaked token #1.
	golden(t, "stake_action_cst_info_1", func() any {
		rec, err := r.StakeActionCstInfo(context.Background(), 1)
		if err != nil {
			t.Fatalf("expected CST stake action 1: %v", err)
		}
		return rec
	})
	// Action 2: bob's stake, still open.
	golden(t, "stake_action_cst_info_2_open", func() any {
		rec, err := r.StakeActionCstInfo(context.Background(), 2)
		if err != nil {
			t.Fatalf("expected CST stake action 2: %v", err)
		}
		return rec
	})
	if _, err := r.StakeActionCstInfo(context.Background(), 999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("StakeActionCstInfo(999) = %v, want store.ErrNotFound", err)
	}
}

func TestGetStakeActionRwalkInfo(t *testing.T) {
	r := repo(t)
	golden(t, "stake_action_rwalk_info_101", func() any {
		rec, err := r.StakeActionRwalkInfo(context.Background(), 101)
		if err != nil {
			t.Fatalf("expected RWalk stake action 101: %v", err)
		}
		return rec
	})
	if _, err := r.StakeActionRwalkInfo(context.Background(), 999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("StakeActionRwalkInfo(999) = %v, want store.ErrNotFound", err)
	}
}

func TestGetStakingRewardsToBeClaimed(t *testing.T) {
	r := repo(t)
	// bob's token #5 sat staked through the round-0 deposit; his reward is unclaimed.
	golden(t, "staking_rewards_to_be_claimed_bob", func() any {
		recs, err := r.StakingRewardsToBeClaimed(context.Background(), aidBob)
		if err != nil {
			t.Fatalf("StakingRewardsToBeClaimed(bob): %v", err)
		}
		return recs
	})
	golden(t, "staking_rewards_to_be_claimed_alice", func() any {
		recs, err := r.StakingRewardsToBeClaimed(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("StakingRewardsToBeClaimed(alice): %v", err)
		}
		return recs
	})
}

func TestGetStakingRewardsCollected(t *testing.T) {
	r := repo(t)
	// alice collected 1 ETH when unstaking token #1.
	golden(t, "staking_rewards_collected_alice", func() any {
		recs, err := r.StakingRewardsCollected(context.Background(), aidAlice, 0, 100)
		if err != nil {
			t.Fatalf("StakingRewardsCollected(alice): %v", err)
		}
		return recs
	})
}

func TestGetStakedTokensCstGlobal(t *testing.T) {
	r := repo(t)
	golden(t, "staked_tokens_cst_global", func() any {
		recs, err := r.StakedTokensCstGlobal(context.Background())
		if err != nil {
			t.Fatalf("StakedTokensCstGlobal: %v", err)
		}
		return recs
	})
}

func TestGetStakedTokensRwalkGlobal(t *testing.T) {
	r := repo(t)
	golden(t, "staked_tokens_rwalk_global", func() any {
		recs, err := r.StakedTokensRwalkGlobal(context.Background())
		if err != nil {
			t.Fatalf("StakedTokensRwalkGlobal: %v", err)
		}
		return recs
	})
}

func TestGetActionIdsForDepositWithClaimInfo(t *testing.T) {
	r := repo(t)
	golden(t, "action_ids_for_deposit_501_alice", func() any {
		recs, err := r.ActionIDsForDepositWithClaimInfo(context.Background(), 501, aidAlice)
		if err != nil {
			t.Fatalf("ActionIDsForDepositWithClaimInfo(alice): %v", err)
		}
		return recs
	})
	golden(t, "action_ids_for_deposit_501_bob", func() any {
		recs, err := r.ActionIDsForDepositWithClaimInfo(context.Background(), 501, aidBob)
		if err != nil {
			t.Fatalf("ActionIDsForDepositWithClaimInfo(bob): %v", err)
		}
		return recs
	})
}

func TestGetGlobalStakingRewards(t *testing.T) {
	r := repo(t)
	golden(t, "global_staking_rewards", func() any {
		recs, err := r.GlobalStakingRewards(context.Background())
		if err != nil {
			t.Fatalf("GlobalStakingRewards: %v", err)
		}
		return recs
	})
}

func TestGetStakingCstRewardsByRound(t *testing.T) {
	r := repo(t)
	golden(t, "staking_cst_rewards_by_round_0", func() any {
		recs, err := r.StakingCstRewardsByRound(context.Background(), 0)
		if err != nil {
			t.Fatalf("StakingCstRewardsByRound(0): %v", err)
		}
		return recs
	})
	got, err := r.StakingCstRewardsByRound(context.Background(), 999)
	if err != nil {
		t.Fatalf("StakingCstRewardsByRound(999): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no staking rewards for round 999, got %d", len(got))
	}
}

func TestGetGlobalStakingCstHistory(t *testing.T) {
	r := repo(t)
	golden(t, "global_staking_cst_history", func() any {
		recs, err := r.GlobalStakingCstHistory(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("GlobalStakingCstHistory: %v", err)
		}
		return recs
	})
}

func TestGetGlobalStakingRwalkHistory(t *testing.T) {
	r := repo(t)
	golden(t, "global_staking_rwalk_history", func() any {
		recs, err := r.GlobalStakingRwalkHistory(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("GlobalStakingRwalkHistory: %v", err)
		}
		return recs
	})
}

func TestGetStakingRwalkMintsGlobal(t *testing.T) {
	r := repo(t)
	golden(t, "staking_rwalk_mints_global", func() any {
		recs, err := r.StakingRwalkMintsGlobal(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("StakingRwalkMintsGlobal: %v", err)
		}
		return recs
	})
}

func TestGetStakingCstMintsGlobal(t *testing.T) {
	r := repo(t)
	golden(t, "staking_cst_mints_global", func() any {
		recs, err := r.StakingCstMintsGlobal(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("StakingCstMintsGlobal: %v", err)
		}
		return recs
	})
	// The legacy layer hardcoded IsRWalk=true on this list even though the
	// query filters is_rwalk=FALSE (copy-paste from the RWalk variant); the
	// conversion fixed it, so pin the corrected flag explicitly.
	recs, err := r.StakingCstMintsGlobal(context.Background(), 0, 100)
	if err != nil {
		t.Fatalf("StakingCstMintsGlobal: %v", err)
	}
	for _, rec := range recs {
		if rec.IsRWalk {
			t.Errorf("CST staking mint token %d reports IsRWalk=true, want false", rec.TokenId)
		}
		if !rec.IsStaker {
			t.Errorf("CST staking mint token %d reports IsStaker=false, want true", rec.TokenId)
		}
	}
}

func TestGetStakingCstByUserByDepositRewards(t *testing.T) {
	r := repo(t)
	golden(t, "staking_cst_by_user_by_deposit_rewards_alice", func() any {
		recs, err := r.StakingCstUserDepositRewards(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("StakingCstUserDepositRewards(alice): %v", err)
		}
		return recs
	})
}

func TestGetStakingCstByUserByTokenRewards(t *testing.T) {
	r := repo(t)
	golden(t, "staking_cst_by_user_by_token_rewards_alice", func() any {
		recs, err := r.StakingCstUserTokenRewards(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("StakingCstUserTokenRewards(alice): %v", err)
		}
		return recs
	})
}

func TestGetStakingCstByUserByTokenRewardsDetailsForToken(t *testing.T) {
	r := repo(t)
	golden(t, "staking_cst_token_rewards_details_alice_token_1", func() any {
		recs, err := r.StakingCstUserTokenRewardDetails(context.Background(), aidAlice, 1)
		if err != nil {
			t.Fatalf("StakingCstUserTokenRewardDetails(alice, 1): %v", err)
		}
		return recs
	})
}
