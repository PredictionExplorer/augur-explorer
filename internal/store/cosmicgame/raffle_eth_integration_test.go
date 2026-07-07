//go:build integration

package cosmicgame

import (
	"context"
	"testing"
)

func TestGetUnclaimedPrizeEthDeposits(t *testing.T) {
	sw := wrapper(t)
	// carol never withdrew her round-0 raffle deposit.
	golden(t, "unclaimed_prize_eth_deposits_carol", func() any {
		return sw.Get_unclaimed_prize_eth_deposits(aidCarol, 0, 100)
	})
	// bob withdrew his round-0 deposit; anything left is pinned as-is.
	golden(t, "unclaimed_prize_eth_deposits_bob", func() any {
		return sw.Get_unclaimed_prize_eth_deposits(aidBob, 0, 100)
	})
}

func TestGetPrizeEthDepositsList(t *testing.T) {
	sw := wrapper(t)
	golden(t, "prize_eth_deposits_list", func() any {
		return sw.Get_prize_eth_deposits_list(0, 100)
	})
	golden(t, "prize_eth_deposits_list_paged", func() any {
		return sw.Get_prize_eth_deposits_list(2, 2)
	})
}

func TestPrizeDepositsByRound(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "prize_deposits_by_round_0", func() any {
		recs, err := r.PrizeDepositsByRound(ctx, 0)
		if err != nil {
			t.Fatalf("PrizeDepositsByRound(0): %v", err)
		}
		return recs
	})
	got, err := r.PrizeDepositsByRound(ctx, 999)
	if err != nil {
		t.Fatalf("PrizeDepositsByRound(999): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no deposits for round 999, got %d", len(got))
	}
}

func TestGetRaffleEthDepositsList(t *testing.T) {
	sw := wrapper(t)
	golden(t, "raffle_eth_deposits_list", func() any {
		return sw.Get_raffle_eth_deposits_list(0, 100)
	})
}

func TestGetChronowarriorEthDepositsList(t *testing.T) {
	sw := wrapper(t)
	golden(t, "chronowarrior_eth_deposits_list", func() any {
		return sw.Get_chronowarrior_eth_deposits_list(0, 100)
	})
}

func TestGetAllEthDepositsByUser(t *testing.T) {
	sw := wrapper(t)
	golden(t, "all_eth_deposits_by_user_alice", func() any {
		return sw.Get_all_eth_deposits_by_user(aidAlice)
	})
	golden(t, "all_eth_deposits_by_user_bob", func() any {
		return sw.Get_all_eth_deposits_by_user(aidBob)
	})
}

func TestGetRaffleEthDepositsByUser(t *testing.T) {
	sw := wrapper(t)
	golden(t, "raffle_eth_deposits_by_user_carol", func() any {
		return sw.Get_raffle_eth_deposits_by_user(aidCarol)
	})
}

func TestGetChronowarriorEthDepositsByUser(t *testing.T) {
	sw := wrapper(t)
	golden(t, "chronowarrior_eth_deposits_by_user_alice", func() any {
		return sw.Get_chronowarrior_eth_deposits_by_user(aidAlice)
	})
}
