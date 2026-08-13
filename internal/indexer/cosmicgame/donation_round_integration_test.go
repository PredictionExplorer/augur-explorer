//go:build integration

package cosmicgame

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
)

// TestDonationRoundResolvedWhenDonationPrecedesClaim reproduces the
// production log order of claimMainPrize(): the CharityWallet emits
// DonationReceived before the game emits MainPrizeClaimed, so the donation
// handler cannot resolve the round at insert time (it stores -1). The
// prize-claim handler must back-patch the row once the round is known.
func TestDonationRoundResolvedWhenDonationPrecedesClaim(t *testing.T) {
	resetDB(t)
	const (
		blockNum = int64(6100)
		roundNum = int64(7)
	)

	ingestTx(t, blockNum, addr(fxGameAddr), 0, []*types.Log{
		buildLog(t, charityWalletABI, "DonationReceived", addr(fxCharityAddr),
			[]any{addr(fxGameAddr)}, []any{eth(1)}),
		buildLog(t, gameABI, "MainPrizeClaimed", addr(fxGameAddr),
			[]any{bigInt(roundNum), addr(fxAlice), bigInt(1)},
			[]any{eth(3), eth(10), bigInt(600)}),
	})

	var donationRound int64
	if err := testDB.SQL.QueryRow(
		`SELECT d.round_num FROM cg_donation_received d
		 JOIN cg_prize_claim pc ON pc.tx_id = d.tx_id
		 WHERE pc.round_num = $1`, roundNum,
	).Scan(&donationRound); err != nil {
		t.Fatalf("reading claim-tx donation: %v", err)
	}
	if donationRound != roundNum {
		t.Fatalf("claim-tx donation round_num = %d, want %d", donationRound, roundNum)
	}

	// A standalone donation in a different transaction must keep -1.
	ingestTx(t, blockNum+1, addr(fxCharityAddr), 0, []*types.Log{
		buildLog(t, charityWalletABI, "DonationReceived", addr(fxCharityAddr),
			[]any{addr(fxAlice)}, []any{eth(2)}),
	})
	var standaloneRound int64
	if err := testDB.SQL.QueryRow(
		`SELECT round_num FROM cg_donation_received WHERE block_num = $1`, blockNum+1,
	).Scan(&standaloneRound); err != nil {
		t.Fatalf("reading standalone donation: %v", err)
	}
	if standaloneRound != -1 {
		t.Fatalf("standalone donation round_num = %d, want -1", standaloneRound)
	}
}
