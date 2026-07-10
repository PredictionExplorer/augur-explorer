//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestNFTDonations(t *testing.T) {
	r := repo(t)
	golden(t, "nft_donations", func() any {
		recs, err := r.NFTDonations(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("NFTDonations: %v", err)
		}
		return recs
	})
}

func TestNFTDonationInfo(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "nft_donation_info_1", func() any {
		rec, err := r.NFTDonationInfo(ctx, 1)
		if err != nil {
			t.Fatalf("NFTDonationInfo(1): %v", err)
		}
		return rec
	})
	if _, err := r.NFTDonationInfo(ctx, 999_999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("NFTDonationInfo(999999) = %v, want ErrNotFound", err)
	}
}

func TestDonatedNFTClaims(t *testing.T) {
	r := repo(t)
	golden(t, "donated_nft_claims", func() any {
		recs, err := r.DonatedNFTClaims(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("DonatedNFTClaims: %v", err)
		}
		return recs
	})
}

func TestNFTDonationsByRound(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "nft_donations_by_prize_round_0", func() any {
		recs, err := r.NFTDonationsByRound(ctx, 0)
		if err != nil {
			t.Fatalf("NFTDonationsByRound(0): %v", err)
		}
		return recs
	})
	golden(t, "nft_donations_by_prize_round_2", func() any {
		recs, err := r.NFTDonationsByRound(ctx, 2)
		if err != nil {
			t.Fatalf("NFTDonationsByRound(2): %v", err)
		}
		return recs
	})
}

func TestNFTDonationsByRoundPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	page, hasMore, err := r.NFTDonationsByRoundPage(ctx, 0, nil, 1)
	if err != nil {
		t.Fatalf("first page: %v", err)
	}
	if hasMore || len(page) != 1 || page[0].Tx.EvtLogId != 5016 ||
		page[0].TokenID != 777 || page[0].DonationIndex != 0 {
		t.Fatalf("page = %+v, hasMore=%v", page, hasMore)
	}

	legacy, err := r.NFTDonationsByRound(ctx, 0)
	if err != nil {
		t.Fatalf("NFTDonationsByRound: %v", err)
	}
	if len(legacy) != len(page) ||
		legacy[0].Tx.EvtLogId != page[0].Tx.EvtLogId ||
		legacy[0].RoundNum != page[0].RoundNum ||
		legacy[0].DonorAddr != page[0].DonorAddr ||
		legacy[0].TokenAddr != page[0].TokenAddr ||
		legacy[0].NFTTokenId != page[0].TokenID ||
		legacy[0].Index != page[0].DonationIndex ||
		legacy[0].NFTTokenURI != page[0].TokenURI {
		t.Fatalf("legacy/page records differ: %+v / %+v", legacy, page)
	}

	exhausted, hasMore, err := r.NFTDonationsByRoundPage(ctx, 0, &DonationPageCursor{
		EventLogID: page[0].Tx.EvtLogId,
	}, 1)
	if err != nil {
		t.Fatalf("exhausted page: %v", err)
	}
	if hasMore || exhausted == nil || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %#v, hasMore=%v; want non-nil empty,false", exhausted, hasMore)
	}

	empty, hasMore, err := r.NFTDonationsByRoundPage(ctx, 3, nil, 1)
	if err != nil {
		t.Fatalf("open-round page: %v", err)
	}
	if hasMore || empty == nil || len(empty) != 0 {
		t.Fatalf("open-round page = %#v, hasMore=%v; want non-nil empty,false", empty, hasMore)
	}
}

func TestNFTDonationsByRoundPagePropagatesCancellation(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, _, err := r.NFTDonationsByRoundPage(ctx, 0, nil, 1); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled page error = %v, want context.Canceled", err)
	}
}

func TestNFTDonationsByToken(t *testing.T) {
	r := repo(t)
	golden(t, "nft_donations_by_token_aid", func() any {
		recs, err := r.NFTDonationsByToken(context.Background(), aidDonatedNFT)
		if err != nil {
			t.Fatalf("NFTDonationsByToken: %v", err)
		}
		return recs
	})
}

func TestUnclaimedDonatedNFTsByRound(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// Round 0's donated NFT was claimed by alice; round 2's is still unclaimed.
	got, err := r.UnclaimedDonatedNFTsByRound(ctx, 0)
	if err != nil {
		t.Fatalf("UnclaimedDonatedNFTsByRound(0): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected round 0 donated NFTs to be claimed, got %d unclaimed", len(got))
	}
	golden(t, "unclaimed_donated_nfts_by_prize_round_2", func() any {
		recs, err := r.UnclaimedDonatedNFTsByRound(ctx, 2)
		if err != nil {
			t.Fatalf("UnclaimedDonatedNFTsByRound(2): %v", err)
		}
		return recs
	})
}

func TestDonatedTokenDistribution(t *testing.T) {
	r := repo(t)
	golden(t, "donated_token_distribution", func() any {
		recs, err := r.DonatedTokenDistribution(context.Background())
		if err != nil {
			t.Fatalf("DonatedTokenDistribution: %v", err)
		}
		return recs
	})
}

func TestNFTDonationsByUser(t *testing.T) {
	r := repo(t)
	golden(t, "nft_donations_by_user_bob", func() any {
		recs, err := r.NFTDonationsByUser(context.Background(), aidBob)
		if err != nil {
			t.Fatalf("NFTDonationsByUser(bob): %v", err)
		}
		return recs
	})
}
