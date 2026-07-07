//go:build integration

package cosmicgame

import "testing"

func TestGetStakeActionCstInfo(t *testing.T) {
	sw := store(t)
	// Action 1: alice staked and later unstaked token #1.
	golden(t, "stake_action_cst_info_1", func() any {
		found, rec := sw.Get_stake_action_cst_info(1)
		if !found {
			t.Fatal("expected CST stake action 1")
		}
		return rec
	})
	// Action 2: bob's stake, still open.
	golden(t, "stake_action_cst_info_2_open", func() any {
		found, rec := sw.Get_stake_action_cst_info(2)
		if !found {
			t.Fatal("expected CST stake action 2")
		}
		return rec
	})
	if found, _ := sw.Get_stake_action_cst_info(999); found {
		t.Error("expected no CST stake action 999")
	}
}

func TestGetStakeActionRwalkInfo(t *testing.T) {
	sw := store(t)
	golden(t, "stake_action_rwalk_info_101", func() any {
		found, rec := sw.Get_stake_action_rwalk_info(101)
		if !found {
			t.Fatal("expected RWalk stake action 101")
		}
		return rec
	})
	if found, _ := sw.Get_stake_action_rwalk_info(999); found {
		t.Error("expected no RWalk stake action 999")
	}
}

func TestGetStakingRewardsToBeClaimed(t *testing.T) {
	sw := store(t)
	// bob's token #5 sat staked through the round-0 deposit; his reward is unclaimed.
	golden(t, "staking_rewards_to_be_claimed_bob", func() any {
		return sw.Get_staking_rewards_to_be_claimed(aidBob)
	})
	golden(t, "staking_rewards_to_be_claimed_alice", func() any {
		return sw.Get_staking_rewards_to_be_claimed(aidAlice)
	})
}

func TestGetStakingRewardsCollected(t *testing.T) {
	sw := store(t)
	// alice collected 1 ETH when unstaking token #1.
	golden(t, "staking_rewards_collected_alice", func() any {
		return sw.Get_staking_rewards_collected(aidAlice, 0, 100)
	})
}

func TestGetStakedTokensCstGlobal(t *testing.T) {
	sw := store(t)
	golden(t, "staked_tokens_cst_global", func() any {
		return sw.Get_staked_tokens_cst_global()
	})
}

func TestGetStakedTokensRwalkGlobal(t *testing.T) {
	sw := store(t)
	golden(t, "staked_tokens_rwalk_global", func() any {
		return sw.Get_staked_tokens_rwalk_global()
	})
}

func TestGetActionIdsForDepositWithClaimInfo(t *testing.T) {
	sw := store(t)
	golden(t, "action_ids_for_deposit_501_alice", func() any {
		return sw.Get_action_ids_for_deposit_with_claim_info(501, aidAlice)
	})
	golden(t, "action_ids_for_deposit_501_bob", func() any {
		return sw.Get_action_ids_for_deposit_with_claim_info(501, aidBob)
	})
}

func TestGetGlobalStakingRewards(t *testing.T) {
	sw := store(t)
	golden(t, "global_staking_rewards", func() any {
		return sw.Get_global_staking_rewards()
	})
}

func TestGetStakingCstRewardsByRound(t *testing.T) {
	sw := store(t)
	golden(t, "staking_cst_rewards_by_round_0", func() any {
		return sw.Get_staking_cst_rewards_by_round(0)
	})
	if got := sw.Get_staking_cst_rewards_by_round(999); len(got) != 0 {
		t.Errorf("expected no staking rewards for round 999, got %d", len(got))
	}
}

func TestGetGlobalStakingCstHistory(t *testing.T) {
	sw := store(t)
	golden(t, "global_staking_cst_history", func() any {
		return sw.Get_global_staking_cst_history(0, 100)
	})
}

func TestGetGlobalStakingRwalkHistory(t *testing.T) {
	sw := store(t)
	golden(t, "global_staking_rwalk_history", func() any {
		return sw.Get_global_staking_rwalk_history(0, 100)
	})
}

func TestGetStakingRwalkMintsGlobal(t *testing.T) {
	sw := store(t)
	golden(t, "staking_rwalk_mints_global", func() any {
		return sw.Get_staking_rwalk_mints_global(0, 100)
	})
}

func TestGetStakingCstMintsGlobal(t *testing.T) {
	sw := store(t)
	golden(t, "staking_cst_mints_global", func() any {
		return sw.Get_staking_cst_mints_global(0, 100)
	})
}

func TestGetStakingCstByUserByDepositRewards(t *testing.T) {
	sw := store(t)
	golden(t, "staking_cst_by_user_by_deposit_rewards_alice", func() any {
		return sw.Get_staking_cst_by_user_by_deposit_rewards(aidAlice)
	})
}

func TestGetStakingCstByUserByTokenRewards(t *testing.T) {
	sw := store(t)
	golden(t, "staking_cst_by_user_by_token_rewards_alice", func() any {
		return sw.Get_staking_cst_by_user_by_token_rewards(aidAlice)
	})
}

func TestGetStakingCstByUserByTokenRewardsDetailsForToken(t *testing.T) {
	sw := store(t)
	golden(t, "staking_cst_token_rewards_details_alice_token_1", func() any {
		return sw.Get_staking_cst_by_user_by_token_rewards_details_for_token(aidAlice, 1)
	})
}
