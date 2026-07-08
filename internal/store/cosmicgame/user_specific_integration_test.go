//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestGetUserInfo(t *testing.T) {
	r := repo(t)
	golden(t, "user_info_alice", func() any {
		rec, err := r.UserInfo(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("expected user info for alice: %v", err)
		}
		return rec
	})
	golden(t, "user_info_bob", func() any {
		rec, err := r.UserInfo(context.Background(), aidBob)
		if err != nil {
			t.Fatalf("expected user info for bob: %v", err)
		}
		return rec
	})
	if _, err := r.UserInfo(context.Background(), 999999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("UserInfo(999999) = %v, want store.ErrNotFound", err)
	}
}

func TestGetPrizeClaimsByUser(t *testing.T) {
	r := repo(t)
	golden(t, "prize_claims_by_user_alice", func() any {
		recs, err := r.PrizeClaimsByUser(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("PrizeClaimsByUser(alice): %v", err)
		}
		return recs
	})
	got, err := r.PrizeClaimsByUser(context.Background(), aidBob)
	if err != nil {
		t.Fatalf("PrizeClaimsByUser(bob): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("bob never won a main prize; got %d claims", len(got))
	}
}

func TestGetBidsByUser(t *testing.T) {
	r := repo(t)
	golden(t, "bids_by_user_alice", func() any {
		recs, err := r.BidsByUser(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("BidsByUser(alice): %v", err)
		}
		return recs
	})
}

func TestGetUnclaimedDonatedNFTByUser(t *testing.T) {
	r := repo(t)
	// emma won round 2 whose donated NFT is unclaimed.
	golden(t, "unclaimed_donated_nft_by_user_emma", func() any {
		recs, err := r.UnclaimedDonatedNFTsByUser(context.Background(), aidEmma)
		if err != nil {
			t.Fatalf("UnclaimedDonatedNFTsByUser(emma): %v", err)
		}
		return recs
	})
	got, err := r.UnclaimedDonatedNFTsByUser(context.Background(), aidAlice)
	if err != nil {
		t.Fatalf("UnclaimedDonatedNFTsByUser(alice): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("alice claimed her donated NFT; got %d unclaimed", len(got))
	}
}

func TestGetRaffleNFTWinningsByUser(t *testing.T) {
	r := repo(t)
	golden(t, "raffle_nft_winnings_by_user_dave", func() any {
		recs, err := r.RaffleNFTWinningsByUser(context.Background(), aidDave)
		if err != nil {
			t.Fatalf("RaffleNFTWinningsByUser(dave): %v", err)
		}
		return recs
	})
}

func TestGetPrizeDepositsChronoWarriorByUser(t *testing.T) {
	r := repo(t)
	golden(t, "prize_deposits_chrono_warrior_by_user_alice", func() any {
		recs, err := r.PrizeDepositsChronoWarriorByUser(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("PrizeDepositsChronoWarriorByUser(alice): %v", err)
		}
		return recs
	})
}

func TestGetPrizeDepositsRaffleEthByUser(t *testing.T) {
	r := repo(t)
	golden(t, "prize_deposits_raffle_eth_by_user_carol", func() any {
		recs, err := r.PrizeDepositsRaffleEthByUser(context.Background(), aidCarol)
		if err != nil {
			t.Fatalf("PrizeDepositsRaffleEthByUser(carol): %v", err)
		}
		return recs
	})
}

func TestGetDonatedNFTClaimsByUser(t *testing.T) {
	r := repo(t)
	golden(t, "donated_nft_claims_by_user_alice", func() any {
		recs, err := r.DonatedNFTClaimsByUser(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("DonatedNFTClaimsByUser(alice): %v", err)
		}
		return recs
	})
}

func TestGetCosmicSignatureNFTListByUser(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_signature_nft_list_by_user_alice", func() any {
		recs, err := r.CosmicSignatureTokensByUser(context.Background(), aidAlice, 0, 100)
		if err != nil {
			t.Fatalf("CosmicSignatureTokensByUser(alice): %v", err)
		}
		return recs
	})
}

func TestGetCosmicTokenTransfersByUser(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_token_transfers_by_user_alice", func() any {
		recs, err := r.CosmicTokenTransfersByUser(context.Background(), aidAlice, 0, 100)
		if err != nil {
			t.Fatalf("CosmicTokenTransfersByUser(alice): %v", err)
		}
		return recs
	})
}

func TestGetCosmicSignatureTransfersByUser(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_signature_transfers_by_user_bob", func() any {
		recs, err := r.CosmicSignatureTransfersByUser(context.Background(), aidBob, 0, 100)
		if err != nil {
			t.Fatalf("CosmicSignatureTransfersByUser(bob): %v", err)
		}
		return recs
	})
}

func TestGetMarketingRewardHistoryByUser(t *testing.T) {
	r := repo(t)
	golden(t, "marketing_reward_history_by_user_emma", func() any {
		recs, err := r.MarketingRewardHistoryByUser(context.Background(), aidEmma, 0, 100)
		if err != nil {
			t.Fatalf("MarketingRewardHistoryByUser(emma): %v", err)
		}
		return recs
	})
}

func TestGetStakedTokensCstByUser(t *testing.T) {
	r := repo(t)
	// bob's token #5 is still staked.
	golden(t, "staked_tokens_cst_by_user_bob", func() any {
		recs, err := r.StakedTokensCstByUser(context.Background(), aidBob)
		if err != nil {
			t.Fatalf("StakedTokensCstByUser(bob): %v", err)
		}
		return recs
	})
	got, err := r.StakedTokensCstByUser(context.Background(), aidAlice)
	if err != nil {
		t.Fatalf("StakedTokensCstByUser(alice): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("alice unstaked; got %d staked CST tokens", len(got))
	}
}

func TestGetStakedTokensRwalkByUser(t *testing.T) {
	r := repo(t)
	golden(t, "staked_tokens_rwalk_by_user_dave", func() any {
		recs, err := r.StakedTokensRwalkByUser(context.Background(), aidDave)
		if err != nil {
			t.Fatalf("StakedTokensRwalkByUser(dave): %v", err)
		}
		return recs
	})
}

func TestGetStakingRwalkMintsByUser(t *testing.T) {
	r := repo(t)
	// carol won the round-0 RandomWalk-staker raffle NFT.
	golden(t, "staking_rwalk_mints_by_user_carol", func() any {
		recs, err := r.StakingRwalkMintsByUser(context.Background(), aidCarol)
		if err != nil {
			t.Fatalf("StakingRwalkMintsByUser(carol): %v", err)
		}
		return recs
	})
}

func TestGetStakingCstMintsByUser(t *testing.T) {
	r := repo(t)
	// bob won the round-2 CST-staker raffle NFT.
	golden(t, "staking_cst_mints_by_user_bob", func() any {
		recs, err := r.StakingCstMintsByUser(context.Background(), aidBob)
		if err != nil {
			t.Fatalf("StakingCstMintsByUser(bob): %v", err)
		}
		return recs
	})
}

func TestGetStakingActionsCstByUser(t *testing.T) {
	r := repo(t)
	golden(t, "staking_actions_cst_by_user_alice", func() any {
		recs, err := r.StakingActionsCstByUser(context.Background(), aidAlice, 0, 100)
		if err != nil {
			t.Fatalf("StakingActionsCstByUser(alice): %v", err)
		}
		return recs
	})
}

func TestGetStakingActionsRwalkByUser(t *testing.T) {
	r := repo(t)
	golden(t, "staking_actions_rwalk_by_user_carol", func() any {
		recs, err := r.StakingActionsRwalkByUser(context.Background(), aidCarol, 0, 100)
		if err != nil {
			t.Fatalf("StakingActionsRwalkByUser(carol): %v", err)
		}
		return recs
	})
}

func TestGetUserNotifRedBoxRewards(t *testing.T) {
	r := repo(t)
	golden(t, "user_notif_red_box_rewards_alice", func() any {
		info, err := r.UserNotifRedBoxRewards(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("UserNotifRedBoxRewards(alice): %v", err)
		}
		return info
	})
	golden(t, "user_notif_red_box_rewards_emma", func() any {
		info, err := r.UserNotifRedBoxRewards(context.Background(), aidEmma)
		if err != nil {
			t.Fatalf("UserNotifRedBoxRewards(emma): %v", err)
		}
		return info
	})
}

func TestGetERC20DonatedPrizesERC20ByWinner(t *testing.T) {
	r := repo(t)
	golden(t, "erc20_donated_prizes_by_winner_alice", func() any {
		recs, err := r.ERC20DonatedPrizesByWinner(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("ERC20DonatedPrizesByWinner(alice): %v", err)
		}
		return recs
	})
}
