//go:build integration

package cosmicgame

import "testing"

func TestGetCharityDonations(t *testing.T) {
	sw := store(t)
	golden(t, "charity_donations", func() any {
		return sw.Get_charity_donations(aidCosmicGame)
	})
}

func TestGetCharityDonationsFromCosmicGame(t *testing.T) {
	sw := store(t)
	golden(t, "charity_donations_from_cosmic_game", func() any {
		return sw.Get_charity_donations_from_cosmic_game(aidCosmicGame)
	})
}

func TestGetCharityDonationsVoluntary(t *testing.T) {
	sw := store(t)
	golden(t, "charity_donations_voluntary", func() any {
		return sw.Get_charity_donations_voluntary(aidCosmicGame)
	})
}

func TestGetCharityWalletWithdrawals(t *testing.T) {
	sw := store(t)
	golden(t, "charity_wallet_withdrawals", func() any {
		return sw.Get_charity_wallet_withdrawals()
	})
}

func TestGetDonationsToCosmicGameSimpleList(t *testing.T) {
	sw := store(t)
	golden(t, "donations_to_cosmic_game_simple_list", func() any {
		return sw.Get_donations_to_cosmic_game_simple_list(0, 100)
	})
}

func TestGetDonationsToCosmicGameSimpleByRound(t *testing.T) {
	sw := store(t)
	golden(t, "donations_to_cosmic_game_simple_by_round_0", func() any {
		return sw.Get_donations_to_cosmic_game_simple_by_round(0)
	})
}

func TestGetDonationsToCosmicGameWithInfoSimpleList(t *testing.T) {
	sw := store(t)
	golden(t, "donations_to_cosmic_game_with_info_simple_list", func() any {
		return sw.Get_donations_to_cosmic_game_with_info_simple_list(0, 100)
	})
}

func TestGetDonationsToCosmicGameWithInfoByRound(t *testing.T) {
	sw := store(t)
	golden(t, "donations_to_cosmic_game_with_info_by_round_0", func() any {
		return sw.Get_donations_to_cosmic_game_with_info_by_round(0)
	})
}

func TestGetDonationsToCosmicGameByUser(t *testing.T) {
	sw := store(t)
	golden(t, "donations_to_cosmic_game_by_user_dave", func() any {
		return sw.Get_donations_to_cosmic_game_by_user(aidDave)
	})
	golden(t, "donations_to_cosmic_game_by_user_emma", func() any {
		return sw.Get_donations_to_cosmic_game_by_user(aidEmma)
	})
}

func TestGetDonationWithInfoRecordInfo(t *testing.T) {
	sw := store(t)
	golden(t, "donation_with_info_record_info_0", func() any {
		return sw.Get_donation_with_info_record_info(0)
	})
	// Missing record: zero-value struct, no error.
	rec := sw.Get_donation_with_info_record_info(999_999)
	if rec.Tx.EvtLogId != 0 || rec.DataJson != "" {
		t.Errorf("expected zero-value record for missing donation, got %+v", rec)
	}
}

func TestGetDonationReceivedEvtIdByTxId(t *testing.T) {
	sw := store(t)
	// tx 1006 carries dave's voluntary DonationReceived at evtlog 5013.
	if got := sw.Get_donation_received_evt_id_by_tx_id(1006, "dr000001"); got != 5013 {
		t.Errorf("donation received evt id: got %d, want 5013", got)
	}
	if got := sw.Get_donation_received_evt_id_by_tx_id(1006, "nosuchsig"); got != 0 {
		t.Errorf("donation received evt id for missing sig: got %d, want 0", got)
	}
}

func TestGetDonationsToCosmicGameBothByRound(t *testing.T) {
	sw := store(t)
	golden(t, "donations_to_cosmic_game_both_by_round_0", func() any {
		return sw.Get_donations_to_cosmic_game_both_by_round(0)
	})
}

func TestGetDonationsToCosmicGameBothAll(t *testing.T) {
	sw := store(t)
	golden(t, "donations_to_cosmic_game_both_all", func() any {
		return sw.Get_donations_to_cosmic_game_both_all()
	})
}
