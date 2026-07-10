//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestCharityDonations(t *testing.T) {
	r := repo(t)
	golden(t, "charity_donations", func() any {
		recs, err := r.CharityDonations(context.Background(), aidCosmicGame)
		if err != nil {
			t.Fatalf("CharityDonations: %v", err)
		}
		return recs
	})
}

func TestCharityDonationsFromCosmicGame(t *testing.T) {
	r := repo(t)
	golden(t, "charity_donations_from_cosmic_game", func() any {
		recs, err := r.CharityDonationsFromCosmicGame(context.Background(), aidCosmicGame)
		if err != nil {
			t.Fatalf("CharityDonationsFromCosmicGame: %v", err)
		}
		return recs
	})
}

func TestCharityDonationsVoluntary(t *testing.T) {
	r := repo(t)
	golden(t, "charity_donations_voluntary", func() any {
		recs, err := r.CharityDonationsVoluntary(context.Background(), aidCosmicGame)
		if err != nil {
			t.Fatalf("CharityDonationsVoluntary: %v", err)
		}
		return recs
	})
}

func TestCharityWalletWithdrawals(t *testing.T) {
	r := repo(t)
	golden(t, "charity_wallet_withdrawals", func() any {
		recs, err := r.CharityWalletWithdrawals(context.Background())
		if err != nil {
			t.Fatalf("CharityWalletWithdrawals: %v", err)
		}
		return recs
	})
}

func TestSimpleEthDonations(t *testing.T) {
	r := repo(t)
	golden(t, "donations_to_cosmic_game_simple_list", func() any {
		recs, err := r.SimpleEthDonations(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("SimpleEthDonations: %v", err)
		}
		return recs
	})
}

func TestSimpleEthDonationsByRound(t *testing.T) {
	r := repo(t)
	golden(t, "donations_to_cosmic_game_simple_by_round_0", func() any {
		recs, err := r.SimpleEthDonationsByRound(context.Background(), 0)
		if err != nil {
			t.Fatalf("SimpleEthDonationsByRound(0): %v", err)
		}
		return recs
	})
}

func TestEthDonationsWithInfo(t *testing.T) {
	r := repo(t)
	golden(t, "donations_to_cosmic_game_with_info_simple_list", func() any {
		recs, err := r.EthDonationsWithInfo(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("EthDonationsWithInfo: %v", err)
		}
		return recs
	})
}

func TestEthDonationsWithInfoByRound(t *testing.T) {
	r := repo(t)
	golden(t, "donations_to_cosmic_game_with_info_by_round_0", func() any {
		recs, err := r.EthDonationsWithInfoByRound(context.Background(), 0)
		if err != nil {
			t.Fatalf("EthDonationsWithInfoByRound(0): %v", err)
		}
		return recs
	})
}

func TestEthDonationsByUser(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "donations_to_cosmic_game_by_user_dave", func() any {
		recs, err := r.EthDonationsByUser(ctx, aidDave)
		if err != nil {
			t.Fatalf("EthDonationsByUser(dave): %v", err)
		}
		return recs
	})
	golden(t, "donations_to_cosmic_game_by_user_emma", func() any {
		recs, err := r.EthDonationsByUser(ctx, aidEmma)
		if err != nil {
			t.Fatalf("EthDonationsByUser(emma): %v", err)
		}
		return recs
	})
}

func TestEthDonationWithInfoRecord(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "donation_with_info_record_info_0", func() any {
		rec, err := r.EthDonationWithInfoRecord(ctx, 0)
		if err != nil {
			t.Fatalf("EthDonationWithInfoRecord(0): %v", err)
		}
		return rec
	})
	if _, err := r.EthDonationWithInfoRecord(ctx, 999_999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("EthDonationWithInfoRecord(999999) = %v, want ErrNotFound", err)
	}
}

func TestDonationReceivedEvtIDByTx(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	got, err := r.DonationReceivedEvtIDByTx(ctx, 1006, "dr000001")
	if err != nil {
		t.Fatalf("DonationReceivedEvtIDByTx: %v", err)
	}
	if got != 5013 {
		t.Errorf("donation received evtlog for tx 1006: got %d, want 5013", got)
	}
	got, err = r.DonationReceivedEvtIDByTx(ctx, 1006, "nosuchsig")
	if err != nil {
		t.Fatalf("DonationReceivedEvtIDByTx(miss): %v", err)
	}
	if got != 0 {
		t.Errorf("donation received evtlog for wrong sig: got %d, want 0", got)
	}
}

func TestEthDonationsByRound(t *testing.T) {
	r := repo(t)
	golden(t, "donations_to_cosmic_game_both_by_round_0", func() any {
		recs, err := r.EthDonationsByRound(context.Background(), 0)
		if err != nil {
			t.Fatalf("EthDonationsByRound(0): %v", err)
		}
		return recs
	})
}

func TestEthDonationsByRoundPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	first, hasMore, err := r.EthDonationsByRoundPage(ctx, 0, nil, 1)
	if err != nil {
		t.Fatalf("first page: %v", err)
	}
	if !hasMore || len(first) != 1 || first[0].Tx.EvtLogId != 5014 ||
		first[0].Kind != RoundEthDonationWithInfo ||
		first[0].ContractRecordID == nil || *first[0].ContractRecordID != 0 ||
		first[0].Data == nil {
		t.Fatalf("first page = %+v, hasMore=%v", first, hasMore)
	}

	second, hasMore, err := r.EthDonationsByRoundPage(ctx, 0, &DonationPageCursor{
		EventLogID: first[0].Tx.EvtLogId,
	}, 1)
	if err != nil {
		t.Fatalf("second page: %v", err)
	}
	if hasMore || len(second) != 1 || second[0].Tx.EvtLogId != 5012 ||
		second[0].Kind != RoundEthDonationPlain ||
		second[0].ContractRecordID != nil || second[0].Data != nil {
		t.Fatalf("second page = %+v, hasMore=%v", second, hasMore)
	}

	legacy, err := r.EthDonationsByRound(ctx, 0)
	if err != nil {
		t.Fatalf("EthDonationsByRound: %v", err)
	}
	paged := append(first, second...)
	if len(legacy) != len(paged) {
		t.Fatalf("legacy/page lengths = %d/%d", len(legacy), len(paged))
	}
	for i := range legacy {
		if legacy[i].Tx.EvtLogId != paged[i].Tx.EvtLogId ||
			legacy[i].RoundNum != paged[i].RoundNum ||
			legacy[i].DonorAddr != paged[i].DonorAddr ||
			legacy[i].Amount != paged[i].EthAmountWei {
			t.Fatalf("legacy/page record %d differs: %+v / %+v", i, legacy[i], paged[i])
		}
	}

	exhausted, hasMore, err := r.EthDonationsByRoundPage(ctx, 0, &DonationPageCursor{
		EventLogID: second[0].Tx.EvtLogId,
	}, 1)
	if err != nil {
		t.Fatalf("exhausted page: %v", err)
	}
	if hasMore || exhausted == nil || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %#v, hasMore=%v; want non-nil empty,false", exhausted, hasMore)
	}

	empty, hasMore, err := r.EthDonationsByRoundPage(ctx, 3, nil, 1)
	if err != nil {
		t.Fatalf("open-round page: %v", err)
	}
	if hasMore || empty == nil || len(empty) != 0 {
		t.Fatalf("open-round page = %#v, hasMore=%v; want non-nil empty,false", empty, hasMore)
	}
}

func TestEthDonationsByRoundPagePropagatesCancellation(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, _, err := r.EthDonationsByRoundPage(ctx, 0, nil, 1); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled page error = %v, want context.Canceled", err)
	}
}

func TestEthDonations(t *testing.T) {
	r := repo(t)
	golden(t, "donations_to_cosmic_game_both_all", func() any {
		recs, err := r.EthDonations(context.Background())
		if err != nil {
			t.Fatalf("EthDonations: %v", err)
		}
		return recs
	})
}
