//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestERC20DonationsByRoundDetailed(t *testing.T) {
	r := repo(t)
	golden(t, "erc20_donations_by_round_detailed_0", func() any {
		recs, err := r.ERC20DonationsByRoundDetailed(context.Background(), 0)
		if err != nil {
			t.Fatalf("ERC20DonationsByRoundDetailed(0): %v", err)
		}
		return recs
	})
}

func TestERC20DonationsByRoundAll(t *testing.T) {
	r := repo(t)
	golden(t, "erc20_donations_by_round_all_0", func() any {
		recs, err := r.ERC20DonationsByRoundAll(context.Background(), 0)
		if err != nil {
			t.Fatalf("ERC20DonationsByRoundAll(0): %v", err)
		}
		return recs
	})
}

func TestERC20DonationsByRoundPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	page, hasMore, err := r.ERC20DonationsByRoundPage(ctx, 0, nil, 1)
	if err != nil {
		t.Fatalf("first page: %v", err)
	}
	if hasMore || len(page) != 1 || page[0].Tx.EvtLogId != 5015 ||
		page[0].AmountBaseUnits != "500000000000000000000" {
		t.Fatalf("page = %+v, hasMore=%v", page, hasMore)
	}

	legacy, err := r.ERC20DonationsByRoundAll(ctx, 0)
	if err != nil {
		t.Fatalf("ERC20DonationsByRoundAll: %v", err)
	}
	if len(legacy) != len(page) ||
		legacy[0].Tx.EvtLogId != page[0].Tx.EvtLogId ||
		legacy[0].RoundNum != page[0].RoundNum ||
		legacy[0].DonorAddr != page[0].DonorAddr ||
		legacy[0].TokenAddr != page[0].TokenAddr ||
		legacy[0].Amount != page[0].AmountBaseUnits {
		t.Fatalf("legacy/page records differ: %+v / %+v", legacy, page)
	}

	exhausted, hasMore, err := r.ERC20DonationsByRoundPage(ctx, 0, &DonationPageCursor{
		EventLogID: page[0].Tx.EvtLogId,
	}, 1)
	if err != nil {
		t.Fatalf("exhausted page: %v", err)
	}
	if hasMore || exhausted == nil || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %#v, hasMore=%v; want non-nil empty,false", exhausted, hasMore)
	}

	empty, hasMore, err := r.ERC20DonationsByRoundPage(ctx, 3, nil, 1)
	if err != nil {
		t.Fatalf("open-round page: %v", err)
	}
	if hasMore || empty == nil || len(empty) != 0 {
		t.Fatalf("open-round page = %#v, hasMore=%v; want non-nil empty,false", empty, hasMore)
	}
}

func TestERC20DonationsByRoundPagePropagatesCancellation(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, _, err := r.ERC20DonationsByRoundPage(ctx, 0, nil, 1); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled page error = %v, want context.Canceled", err)
	}
}

func TestERC20DonationsByRoundSummarized(t *testing.T) {
	r := repo(t)
	golden(t, "erc20_donations_by_round_summarized_0", func() any {
		recs, err := r.ERC20DonationsByRoundSummarized(context.Background(), 0)
		if err != nil {
			t.Fatalf("ERC20DonationsByRoundSummarized(0): %v", err)
		}
		return recs
	})
}

func TestERC20Donations(t *testing.T) {
	r := repo(t)
	golden(t, "erc20_donations_global", func() any {
		recs, err := r.ERC20Donations(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("ERC20Donations: %v", err)
		}
		return recs
	})
}

func TestERC20DonationInfo(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "erc20_donation_info_1", func() any {
		rec, err := r.ERC20DonationInfo(ctx, 1)
		if err != nil {
			t.Fatalf("ERC20DonationInfo(1): %v", err)
		}
		return rec
	})
	if _, err := r.ERC20DonationInfo(ctx, 999_999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("ERC20DonationInfo(999999) = %v, want ErrNotFound", err)
	}
}

func TestERC20DonationsByUser(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "erc20_donations_by_user_alice", func() any {
		recs, err := r.ERC20DonationsByUser(ctx, aidAlice)
		if err != nil {
			t.Fatalf("ERC20DonationsByUser(alice): %v", err)
		}
		return recs
	})
	got, err := r.ERC20DonationsByUser(ctx, aidZero)
	if err != nil {
		t.Fatalf("ERC20DonationsByUser(zero): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no ERC20 donations from the zero address, got %d", len(got))
	}
}

func TestERC20DonationClaims(t *testing.T) {
	r := repo(t)
	golden(t, "erc20_donated_token_claims_global", func() any {
		recs, err := r.ERC20DonationClaims(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("ERC20DonationClaims: %v", err)
		}
		return recs
	})
}

func TestERC20DonationClaimsByUser(t *testing.T) {
	r := repo(t)
	golden(t, "erc20_donated_token_claims_by_user_alice", func() any {
		recs, err := r.ERC20DonationClaimsByUser(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("ERC20DonationClaimsByUser(alice): %v", err)
		}
		return recs
	})
}

func TestERC20DonationClaimsByRound(t *testing.T) {
	r := repo(t)
	golden(t, "erc20_donated_token_claims_by_round_0", func() any {
		recs, err := r.ERC20DonationClaimsByRound(context.Background(), 0)
		if err != nil {
			t.Fatalf("ERC20DonationClaimsByRound(0): %v", err)
		}
		return recs
	})
}
