//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"
)

func TestUnclaimedPrizeEthDeposits(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// carol never withdrew her round-0 raffle deposit.
	golden(t, "unclaimed_prize_eth_deposits_carol", func() any {
		recs, err := r.UnclaimedPrizeEthDeposits(ctx, aidCarol, 0, 100)
		if err != nil {
			t.Fatalf("UnclaimedPrizeEthDeposits(carol): %v", err)
		}
		return recs
	})
	// bob withdrew his round-0 deposit; anything left is pinned as-is.
	golden(t, "unclaimed_prize_eth_deposits_bob", func() any {
		recs, err := r.UnclaimedPrizeEthDeposits(ctx, aidBob, 0, 100)
		if err != nil {
			t.Fatalf("UnclaimedPrizeEthDeposits(bob): %v", err)
		}
		return recs
	})
}

func TestPrizeEthDeposits(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "prize_eth_deposits_list", func() any {
		recs, err := r.PrizeEthDeposits(ctx, 0, 100)
		if err != nil {
			t.Fatalf("PrizeEthDeposits: %v", err)
		}
		return recs
	})
	golden(t, "prize_eth_deposits_list_paged", func() any {
		recs, err := r.PrizeEthDeposits(ctx, 2, 2)
		if err != nil {
			t.Fatalf("PrizeEthDeposits(paged): %v", err)
		}
		return recs
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

func TestRaffleEthDepositsByRoundPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	first, hasMore, err := r.RaffleEthDepositsByRoundPage(ctx, 0, nil, 1)
	if err != nil {
		t.Fatalf("first page: %v", err)
	}
	if !hasMore || len(first) != 1 {
		t.Fatalf("first page = %d records, hasMore=%v; want 1,true", len(first), hasMore)
	}
	if first[0].WinnerIndex != 0 ||
		first[0].Tx.EvtLogId != 5022 ||
		first[0].EthAmountWei != "50000000000000000" ||
		!first[0].Claimed {
		t.Fatalf("first deposit = %+v", first[0])
	}

	second, hasMore, err := r.RaffleEthDepositsByRoundPage(ctx, 0, &RaffleEthDepositPageCursor{
		WinnerIndex: first[0].WinnerIndex,
		EventLogID:  first[0].Tx.EvtLogId,
	}, 1)
	if err != nil {
		t.Fatalf("second page: %v", err)
	}
	if hasMore || len(second) != 1 {
		t.Fatalf("second page = %d records, hasMore=%v; want 1,false", len(second), hasMore)
	}
	if second[0].WinnerIndex != 1 ||
		second[0].Tx.EvtLogId != 5024 ||
		second[0].EthAmountWei != "50000000000000000" ||
		second[0].Claimed {
		t.Fatalf("second deposit = %+v", second[0])
	}

	legacy, err := r.PrizeDepositsByRound(ctx, 0)
	if err != nil {
		t.Fatalf("PrizeDepositsByRound: %v", err)
	}
	paged := append(first, second...)
	if len(legacy) != len(paged) {
		t.Fatalf("legacy/page lengths = %d/%d", len(legacy), len(paged))
	}
	for i := range legacy {
		if legacy[i].Tx.EvtLogId != paged[i].Tx.EvtLogId ||
			legacy[i].WinnerIndex != paged[i].WinnerIndex ||
			legacy[i].WinnerAddr != paged[i].WinnerAddr ||
			legacy[i].Claimed != paged[i].Claimed {
			t.Fatalf("legacy/page record %d differs: %+v / %+v", i, legacy[i], paged[i])
		}
	}

	exhausted, hasMore, err := r.RaffleEthDepositsByRoundPage(ctx, 0, &RaffleEthDepositPageCursor{
		WinnerIndex: second[0].WinnerIndex,
		EventLogID:  second[0].Tx.EvtLogId,
	}, 1)
	if err != nil {
		t.Fatalf("exhausted page: %v", err)
	}
	if hasMore || exhausted == nil || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %#v, hasMore=%v; want non-nil empty,false", exhausted, hasMore)
	}
}

func TestRaffleEthDepositsByRoundPagePropagatesCancellation(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, _, err := r.RaffleEthDepositsByRoundPage(ctx, 0, nil, 1); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled page error = %v, want context.Canceled", err)
	}
}

func TestRaffleEthDeposits(t *testing.T) {
	r := repo(t)
	golden(t, "raffle_eth_deposits_list", func() any {
		recs, err := r.RaffleEthDeposits(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("RaffleEthDeposits: %v", err)
		}
		return recs
	})
}

func TestChronoWarriorEthDeposits(t *testing.T) {
	r := repo(t)
	golden(t, "chronowarrior_eth_deposits_list", func() any {
		recs, err := r.ChronoWarriorEthDeposits(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("ChronoWarriorEthDeposits: %v", err)
		}
		return recs
	})
}

func TestEthDepositsByUser(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "all_eth_deposits_by_user_alice", func() any {
		recs, err := r.EthDepositsByUser(ctx, aidAlice)
		if err != nil {
			t.Fatalf("EthDepositsByUser(alice): %v", err)
		}
		return recs
	})
	golden(t, "all_eth_deposits_by_user_bob", func() any {
		recs, err := r.EthDepositsByUser(ctx, aidBob)
		if err != nil {
			t.Fatalf("EthDepositsByUser(bob): %v", err)
		}
		return recs
	})
}

func TestRaffleEthDepositsByUser(t *testing.T) {
	r := repo(t)
	golden(t, "raffle_eth_deposits_by_user_carol", func() any {
		recs, err := r.RaffleEthDepositsByUser(context.Background(), aidCarol)
		if err != nil {
			t.Fatalf("RaffleEthDepositsByUser(carol): %v", err)
		}
		return recs
	})
}

func TestChronoWarriorEthDepositsByUser(t *testing.T) {
	r := repo(t)
	golden(t, "chronowarrior_eth_deposits_by_user_alice", func() any {
		recs, err := r.ChronoWarriorEthDepositsByUser(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("ChronoWarriorEthDepositsByUser(alice): %v", err)
		}
		return recs
	})
}
