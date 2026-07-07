//go:build integration

package cosmicgame

import "testing"

func TestGetUserInfo(t *testing.T) {
	sw := store(t)
	golden(t, "user_info_alice", func() any {
		found, rec := sw.Get_user_info(aidAlice)
		if !found {
			t.Fatal("expected user info for alice")
		}
		return rec
	})
	golden(t, "user_info_bob", func() any {
		found, rec := sw.Get_user_info(aidBob)
		if !found {
			t.Fatal("expected user info for bob")
		}
		return rec
	})
}

func TestGetPrizeClaimsByUser(t *testing.T) {
	sw := store(t)
	golden(t, "prize_claims_by_user_alice", func() any {
		return sw.Get_prize_claims_by_user(aidAlice)
	})
	if got := sw.Get_prize_claims_by_user(aidBob); len(got) != 0 {
		t.Errorf("bob never won a main prize; got %d claims", len(got))
	}
}

func TestGetBidsByUser(t *testing.T) {
	sw := store(t)
	golden(t, "bids_by_user_alice", func() any {
		return sw.Get_bids_by_user(aidAlice)
	})
}

func TestGetUnclaimedDonatedNFTByUser(t *testing.T) {
	sw := store(t)
	// emma won round 2 whose donated NFT is unclaimed.
	golden(t, "unclaimed_donated_nft_by_user_emma", func() any {
		return sw.Get_unclaimed_donated_nft_by_user(aidEmma)
	})
	if got := sw.Get_unclaimed_donated_nft_by_user(aidAlice); len(got) != 0 {
		t.Errorf("alice claimed her donated NFT; got %d unclaimed", len(got))
	}
}

func TestGetRaffleNFTWinningsByUser(t *testing.T) {
	sw := store(t)
	golden(t, "raffle_nft_winnings_by_user_dave", func() any {
		return sw.Get_raffle_nft_winnings_by_user(aidDave)
	})
}

func TestGetPrizeDepositsChronoWarriorByUser(t *testing.T) {
	sw := store(t)
	golden(t, "prize_deposits_chrono_warrior_by_user_alice", func() any {
		return sw.Get_prize_deposits_chrono_warrior_by_user(aidAlice)
	})
}

func TestGetPrizeDepositsRaffleEthByUser(t *testing.T) {
	sw := store(t)
	golden(t, "prize_deposits_raffle_eth_by_user_carol", func() any {
		return sw.Get_prize_deposits_raffle_eth_by_user(aidCarol)
	})
}

func TestGetDonatedNFTClaimsByUser(t *testing.T) {
	sw := store(t)
	golden(t, "donated_nft_claims_by_user_alice", func() any {
		return sw.Get_donated_nft_claims_by_user(aidAlice)
	})
}

func TestGetCosmicSignatureNFTListByUser(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_signature_nft_list_by_user_alice", func() any {
		return sw.Get_cosmic_signature_nft_list_by_user(aidAlice, 0, 100)
	})
}

func TestGetCosmicTokenTransfersByUser(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_token_transfers_by_user_alice", func() any {
		return sw.Get_cosmic_token_transfers_by_user(aidAlice, 0, 100)
	})
}

func TestGetCosmicSignatureTransfersByUser(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_signature_transfers_by_user_bob", func() any {
		return sw.Get_cosmic_signature_transfers_by_user(aidBob, 0, 100)
	})
}

func TestGetMarketingRewardHistoryByUser(t *testing.T) {
	sw := store(t)
	golden(t, "marketing_reward_history_by_user_emma", func() any {
		return sw.Get_marketing_reward_history_by_user(aidEmma, 0, 100)
	})
}

func TestGetStakedTokensCstByUser(t *testing.T) {
	sw := store(t)
	// bob's token #5 is still staked.
	golden(t, "staked_tokens_cst_by_user_bob", func() any {
		return sw.Get_staked_tokens_cst_by_user(aidBob)
	})
	if got := sw.Get_staked_tokens_cst_by_user(aidAlice); len(got) != 0 {
		t.Errorf("alice unstaked; got %d staked CST tokens", len(got))
	}
}

func TestGetStakedTokensRwalkByUser(t *testing.T) {
	sw := store(t)
	golden(t, "staked_tokens_rwalk_by_user_dave", func() any {
		return sw.Get_staked_tokens_rwalk_by_user(aidDave)
	})
}

func TestGetStakingRwalkMintsByUser(t *testing.T) {
	sw := store(t)
	// carol won the round-0 RandomWalk-staker raffle NFT.
	golden(t, "staking_rwalk_mints_by_user_carol", func() any {
		return sw.Get_staking_rwalk_mints_by_user(aidCarol)
	})
}

func TestGetStakingCstMintsByUser(t *testing.T) {
	sw := store(t)
	// bob won the round-2 CST-staker raffle NFT.
	golden(t, "staking_cst_mints_by_user_bob", func() any {
		return sw.Get_staking_cst_mints_by_user(aidBob)
	})
}

func TestGetStakingActionsCstByUser(t *testing.T) {
	sw := store(t)
	golden(t, "staking_actions_cst_by_user_alice", func() any {
		return sw.Get_staking_actions_cst_by_user(aidAlice, 0, 100)
	})
}

func TestGetStakingActionsRwalkByUser(t *testing.T) {
	sw := store(t)
	golden(t, "staking_actions_rwalk_by_user_carol", func() any {
		return sw.Get_staking_actions_rwalk_by_user(aidCarol, 0, 100)
	})
}

func TestGetUserNotifRedBoxRewards(t *testing.T) {
	sw := store(t)
	golden(t, "user_notif_red_box_rewards_alice", func() any {
		return sw.Get_user_notif_red_box_rewards(aidAlice)
	})
	golden(t, "user_notif_red_box_rewards_emma", func() any {
		return sw.Get_user_notif_red_box_rewards(aidEmma)
	})
}

func TestGetERC20DonatedPrizesERC20ByWinner(t *testing.T) {
	sw := store(t)
	golden(t, "erc20_donated_prizes_by_winner_alice", func() any {
		return sw.Get_erc20_donated_prizes_erc20_by_winner(aidAlice)
	})
}
