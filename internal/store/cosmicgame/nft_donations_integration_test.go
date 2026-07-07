//go:build integration

package cosmicgame

import "testing"

func TestGetNFTDonations(t *testing.T) {
	sw := wrapper(t)
	golden(t, "nft_donations", func() any {
		return sw.Get_NFT_donations(0, 100)
	})
}

func TestGetNFTDonationInfo(t *testing.T) {
	sw := wrapper(t)
	golden(t, "nft_donation_info_1", func() any {
		found, rec := sw.Get_NFT_donation_info(1)
		if !found {
			t.Fatal("expected NFT donation record 1 to exist")
		}
		return rec
	})
	if found, _ := sw.Get_NFT_donation_info(999_999); found {
		t.Error("expected NFT donation record 999999 to be missing")
	}
}

func TestGetDonatedNFTClaims(t *testing.T) {
	sw := wrapper(t)
	golden(t, "donated_nft_claims", func() any {
		return sw.Get_donated_nft_claims(0, 100)
	})
}

func TestGetNFTDonationsByPrize(t *testing.T) {
	sw := wrapper(t)
	golden(t, "nft_donations_by_prize_round_0", func() any {
		return sw.Get_nft_donations_by_prize(0)
	})
	golden(t, "nft_donations_by_prize_round_2", func() any {
		return sw.Get_nft_donations_by_prize(2)
	})
}

func TestGetNFTDonationsByTokenAid(t *testing.T) {
	sw := wrapper(t)
	golden(t, "nft_donations_by_token_aid", func() any {
		return sw.Get_nft_donations_by_token_aid(aidDonatedNFT)
	})
}

func TestGetUnclaimedDonatedNFTsByPrize(t *testing.T) {
	sw := wrapper(t)
	// Round 0's donated NFT was claimed by alice; round 2's is still unclaimed.
	if got := sw.Get_unclaimed_donated_nfts_by_prize(0); len(got) != 0 {
		t.Errorf("expected round 0 donated NFTs to be claimed, got %d unclaimed", len(got))
	}
	golden(t, "unclaimed_donated_nfts_by_prize_round_2", func() any {
		return sw.Get_unclaimed_donated_nfts_by_prize(2)
	})
}

func TestGetDonatedTokenDistribution(t *testing.T) {
	sw := wrapper(t)
	golden(t, "donated_token_distribution", func() any {
		return sw.Get_donated_token_distribution()
	})
}

func TestGetNFTDonationsByUser(t *testing.T) {
	sw := wrapper(t)
	golden(t, "nft_donations_by_user_bob", func() any {
		return sw.Get_nft_donations_by_user(aidBob)
	})
}
