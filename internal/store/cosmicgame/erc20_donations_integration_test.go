//go:build integration

package cosmicgame

import "testing"

func TestGetERC20DonationsByRoundDetailed(t *testing.T) {
	sw := wrapper(t)
	golden(t, "erc20_donations_by_round_detailed_0", func() any {
		return sw.Get_erc20_donations_by_round_detailed(0)
	})
}

func TestGetERC20DonationsByRoundAll(t *testing.T) {
	sw := wrapper(t)
	golden(t, "erc20_donations_by_round_all_0", func() any {
		return sw.Get_erc20_donations_by_round_all(0)
	})
}

func TestGetERC20DonationsByRoundSummarized(t *testing.T) {
	sw := wrapper(t)
	golden(t, "erc20_donations_by_round_summarized_0", func() any {
		return sw.Get_erc20_donations_by_round_summarized(0)
	})
}

func TestGetERC20DonationsGlobal(t *testing.T) {
	sw := wrapper(t)
	golden(t, "erc20_donations_global", func() any {
		return sw.Get_erc20_donations_global(0, 100)
	})
}

func TestGetERC20DonationInfo(t *testing.T) {
	sw := wrapper(t)
	golden(t, "erc20_donation_info_1", func() any {
		found, rec := sw.Get_erc20_donation_info(1)
		if !found {
			t.Fatal("expected ERC20 donation record 1 to exist")
		}
		return rec
	})
	if found, _ := sw.Get_erc20_donation_info(999_999); found {
		t.Error("expected ERC20 donation record 999999 to be missing")
	}
}

func TestGetERC20DonationsByUser(t *testing.T) {
	sw := wrapper(t)
	golden(t, "erc20_donations_by_user_alice", func() any {
		return sw.Get_erc20_donations_by_user(aidAlice)
	})
	if got := sw.Get_erc20_donations_by_user(aidZero); len(got) != 0 {
		t.Errorf("expected no ERC20 donations from the zero address, got %d", len(got))
	}
}

func TestGetERC20DonatedTokenClaimsGlobal(t *testing.T) {
	sw := wrapper(t)
	golden(t, "erc20_donated_token_claims_global", func() any {
		return sw.Get_erc20_donated_token_claims_global(0, 100)
	})
}

func TestGetERC20DonatedTokenClaimsByUser(t *testing.T) {
	sw := wrapper(t)
	golden(t, "erc20_donated_token_claims_by_user_alice", func() any {
		return sw.Get_erc20_donated_token_claims_by_user(aidAlice)
	})
}

func TestGetERC20DonatedTokenClaimsByRound(t *testing.T) {
	sw := wrapper(t)
	golden(t, "erc20_donated_token_claims_by_round_0", func() any {
		return sw.Get_erc20_donated_token_claims_by_round(0)
	})
}
