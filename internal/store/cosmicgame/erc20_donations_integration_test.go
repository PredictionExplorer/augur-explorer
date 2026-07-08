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
