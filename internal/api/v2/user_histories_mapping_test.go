package v2

import (
	"strings"
	"testing"
	"time"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func validUserRaffleEthDepositRecord() cgstore.UserRaffleEthDepositRecord {
	return cgstore.UserRaffleEthDepositRecord{
		Tx:           validDonationTransaction(),
		RoundNum:     4,
		WinnerIndex:  2,
		WinnerAid:    21,
		WinnerAddr:   "0x2222222222222222222222222222222222222222",
		EthAmountWei: "50000000000000000",
	}
}

func validUserDepositWithdrawalRecord() cgstore.UserDepositWithdrawalRecord {
	return cgstore.UserDepositWithdrawalRecord{
		EventLogID:      160,
		TxHash:          "0x" + strings.Repeat("CD", 32),
		DateTime:        "2026-01-01T02:01:40+01:00",
		BeneficiaryAddr: "0x4444444444444444444444444444444444444444",
	}
}

func TestMapUserRaffleEthDeposit(t *testing.T) {
	t.Parallel()

	t.Run("unclaimed raffle deposit", func(t *testing.T) {
		t.Parallel()
		record := validUserRaffleEthDepositRecord()
		got, err := mapUserRaffleEthDeposit(record)
		if err != nil {
			t.Fatalf("mapUserRaffleEthDeposit: %v", err)
		}
		if got.Source != EthDepositSourceRaffle || got.Claimed || got.Withdrawal != nil ||
			got.EthAmountWei != "50000000000000000" || got.Round != 4 ||
			got.WinnerIndex != 2 || got.EventLogId != 100 ||
			got.TransactionHash != "0x"+strings.Repeat("ab", 32) ||
			!got.OccurredAt.Equal(time.Date(2026, 1, 1, 0, 1, 40, 0, time.UTC)) {
			t.Fatalf("deposit = %+v", got)
		}
	})

	t.Run("claimed chrono deposit with withdrawal", func(t *testing.T) {
		t.Parallel()
		record := validUserRaffleEthDepositRecord()
		record.IsChronoWarrior = true
		record.Claimed = true
		withdrawal := validUserDepositWithdrawalRecord()
		record.Withdrawal = &withdrawal
		got, err := mapUserRaffleEthDeposit(record)
		if err != nil {
			t.Fatalf("mapUserRaffleEthDeposit: %v", err)
		}
		if got.Source != EthDepositSourceChronoWarrior || !got.Claimed || got.Withdrawal == nil {
			t.Fatalf("deposit = %+v", got)
		}
		if got.Withdrawal.EventLogId != 160 ||
			got.Withdrawal.TransactionHash != "0x"+strings.Repeat("cd", 32) ||
			got.Withdrawal.BeneficiaryAddress != "0x4444444444444444444444444444444444444444" ||
			!got.Withdrawal.OccurredAt.Equal(time.Date(2026, 1, 1, 1, 1, 40, 0, time.UTC)) {
			t.Fatalf("withdrawal = %+v", got.Withdrawal)
		}
	})

	t.Run("claimed without withdrawal stays mappable", func(t *testing.T) {
		t.Parallel()
		record := validUserRaffleEthDepositRecord()
		record.Claimed = true
		got, err := mapUserRaffleEthDeposit(record)
		if err != nil || !got.Claimed || got.Withdrawal != nil {
			t.Fatalf("deposit = %+v, err=%v", got, err)
		}
	})

	rejections := map[string]func() error{
		"negative round": func() error {
			record := validUserRaffleEthDepositRecord()
			record.RoundNum = -1
			_, err := mapUserRaffleEthDeposit(record)
			return err
		},
		"negative winner index": func() error {
			record := validUserRaffleEthDepositRecord()
			record.WinnerIndex = -1
			_, err := mapUserRaffleEthDeposit(record)
			return err
		},
		"empty amount": func() error {
			record := validUserRaffleEthDepositRecord()
			record.EthAmountWei = ""
			_, err := mapUserRaffleEthDeposit(record)
			return err
		},
		"negative amount": func() error {
			record := validUserRaffleEthDepositRecord()
			record.EthAmountWei = "-1"
			_, err := mapUserRaffleEthDeposit(record)
			return err
		},
		"invalid transaction hash": func() error {
			record := validUserRaffleEthDepositRecord()
			record.Tx.TxHash = "not-a-hash"
			_, err := mapUserRaffleEthDeposit(record)
			return err
		},
		"unclaimed with withdrawal": func() error {
			record := validUserRaffleEthDepositRecord()
			withdrawal := validUserDepositWithdrawalRecord()
			record.Withdrawal = &withdrawal
			_, err := mapUserRaffleEthDeposit(record)
			return err
		},
		"withdrawal without event id": func() error {
			record := validUserRaffleEthDepositRecord()
			record.Claimed = true
			withdrawal := validUserDepositWithdrawalRecord()
			withdrawal.EventLogID = 0
			record.Withdrawal = &withdrawal
			_, err := mapUserRaffleEthDeposit(record)
			return err
		},
		"withdrawal with bad hash": func() error {
			record := validUserRaffleEthDepositRecord()
			record.Claimed = true
			withdrawal := validUserDepositWithdrawalRecord()
			withdrawal.TxHash = "bad"
			record.Withdrawal = &withdrawal
			_, err := mapUserRaffleEthDeposit(record)
			return err
		},
		"withdrawal with bad timestamp": func() error {
			record := validUserRaffleEthDepositRecord()
			record.Claimed = true
			withdrawal := validUserDepositWithdrawalRecord()
			withdrawal.DateTime = "yesterday"
			record.Withdrawal = &withdrawal
			_, err := mapUserRaffleEthDeposit(record)
			return err
		},
		"withdrawal with bad beneficiary": func() error {
			record := validUserRaffleEthDepositRecord()
			record.Claimed = true
			withdrawal := validUserDepositWithdrawalRecord()
			withdrawal.BeneficiaryAddr = "bad"
			record.Withdrawal = &withdrawal
			_, err := mapUserRaffleEthDeposit(record)
			return err
		},
	}
	for name, test := range rejections {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := test(); err == nil {
				t.Fatal("mapper accepted invalid record")
			}
		})
	}
}

func validUserRaffleNftWinRecord() cgstore.UserRaffleNftWinRecord {
	return cgstore.UserRaffleNftWinRecord{
		Tx:           validDonationTransaction(),
		RoundNum:     4,
		WinnerIndex:  1,
		WinnerAid:    21,
		WinnerAddr:   "0x2222222222222222222222222222222222222222",
		TokenID:      9,
		CstAmountWei: "30000000000000000000",
	}
}

func TestMapUserRaffleNftWin(t *testing.T) {
	t.Parallel()

	t.Run("all pools map", func(t *testing.T) {
		t.Parallel()
		pools := []struct {
			rwalk, staker bool
		}{
			{false, false}, // bidder raffle
			{true, true},   // RandomWalk staker raffle
			{false, true},  // Cosmic Signature staker raffle
		}
		for _, pool := range pools {
			record := validUserRaffleNftWinRecord()
			record.IsRWalk = pool.rwalk
			record.IsStaker = pool.staker
			got, err := mapUserRaffleNftWin(record)
			if err != nil {
				t.Fatalf("mapUserRaffleNftWin(%+v): %v", pool, err)
			}
			if got.IsRandomWalk != pool.rwalk || got.IsStaker != pool.staker ||
				got.NftTokenId != 9 || got.CstAmountWei != "30000000000000000000" ||
				got.Round != 4 || got.WinnerIndex != 1 {
				t.Fatalf("win = %+v", got)
			}
		}
	})

	rejections := map[string]func() error{
		"negative token": func() error {
			record := validUserRaffleNftWinRecord()
			record.TokenID = -1
			_, err := mapUserRaffleNftWin(record)
			return err
		},
		"negative round": func() error {
			record := validUserRaffleNftWinRecord()
			record.RoundNum = -1
			_, err := mapUserRaffleNftWin(record)
			return err
		},
		"malformed cst amount": func() error {
			record := validUserRaffleNftWinRecord()
			record.CstAmountWei = "1.5"
			_, err := mapUserRaffleNftWin(record)
			return err
		},
		"randomwalk flag without staker flag": func() error {
			record := validUserRaffleNftWinRecord()
			record.IsRWalk = true
			record.IsStaker = false
			_, err := mapUserRaffleNftWin(record)
			return err
		},
		"invalid transaction": func() error {
			record := validUserRaffleNftWinRecord()
			record.Tx.EvtLogId = 0
			_, err := mapUserRaffleNftWin(record)
			return err
		},
	}
	for name, test := range rejections {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := test(); err == nil {
				t.Fatal("mapper accepted invalid record")
			}
		})
	}
}

func validUserDonatedNftRecord() cgstore.UserDonatedNftRecord {
	// #nosec G101 -- deterministic chain fixture values, not credentials.
	return cgstore.UserDonatedNftRecord{
		Tx:             validDonationTransaction(),
		RoundNum:       4,
		DonorAddr:      "0x2222222222222222222222222222222222222222",
		TokenAddr:      "0x3333333333333333333333333333333333333333",
		TokenID:        888,
		DonationIndex:  1,
		TokenURI:       "https://nft.example/888",
		RoundWinnerAid: 25,
	}
}

func validUserDonatedNftClaimRecord() cgstore.UserDonatedNftClaimRecord {
	return cgstore.UserDonatedNftClaimRecord{
		EventLogID:  170,
		TxHash:      "0x" + strings.Repeat("EF", 32),
		DateTime:    "2026-01-01T03:01:40+01:00",
		ClaimerAid:  22,
		ClaimerAddr: "0x5555555555555555555555555555555555555555",
	}
}

func TestMapUserDonatedNft(t *testing.T) {
	t.Parallel()

	t.Run("unclaimed", func(t *testing.T) {
		t.Parallel()
		got, err := mapUserDonatedNft(validUserDonatedNftRecord())
		if err != nil {
			t.Fatalf("mapUserDonatedNft: %v", err)
		}
		if got.Claimed || got.Claim != nil || got.TokenId != 888 ||
			got.DonationIndex != 1 || got.TokenUri != "https://nft.example/888" ||
			got.DonorAddress != "0x2222222222222222222222222222222222222222" {
			t.Fatalf("donated NFT = %+v", got)
		}
	})

	t.Run("claimed", func(t *testing.T) {
		t.Parallel()
		record := validUserDonatedNftRecord()
		record.Claimed = true
		claim := validUserDonatedNftClaimRecord()
		record.Claim = &claim
		got, err := mapUserDonatedNft(record)
		if err != nil {
			t.Fatalf("mapUserDonatedNft: %v", err)
		}
		if !got.Claimed || got.Claim == nil ||
			got.Claim.EventLogId != 170 ||
			got.Claim.ClaimerAddress != "0x5555555555555555555555555555555555555555" ||
			got.Claim.TransactionHash != "0x"+strings.Repeat("ef", 32) ||
			!got.Claim.OccurredAt.Equal(time.Date(2026, 1, 1, 2, 1, 40, 0, time.UTC)) {
			t.Fatalf("claim = %+v", got.Claim)
		}
	})

	rejections := map[string]func() error{
		"claimed without claim": func() error {
			record := validUserDonatedNftRecord()
			record.Claimed = true
			_, err := mapUserDonatedNft(record)
			return err
		},
		"invalid donation transaction": func() error {
			record := validUserDonatedNftRecord()
			record.Tx.EvtLogId = 0
			_, err := mapUserDonatedNft(record)
			return err
		},
		"claim without event id": func() error {
			record := validUserDonatedNftRecord()
			record.Claimed = true
			claim := validUserDonatedNftClaimRecord()
			claim.EventLogID = 0
			record.Claim = &claim
			_, err := mapUserDonatedNft(record)
			return err
		},
		"claim with invalid hash": func() error {
			record := validUserDonatedNftRecord()
			record.Claimed = true
			claim := validUserDonatedNftClaimRecord()
			claim.TxHash = "bad"
			record.Claim = &claim
			_, err := mapUserDonatedNft(record)
			return err
		},
		"claim without claimed flag": func() error {
			record := validUserDonatedNftRecord()
			claim := validUserDonatedNftClaimRecord()
			record.Claim = &claim
			_, err := mapUserDonatedNft(record)
			return err
		},
		"invalid donor": func() error {
			record := validUserDonatedNftRecord()
			record.DonorAddr = "bad"
			_, err := mapUserDonatedNft(record)
			return err
		},
		"invalid token contract": func() error {
			record := validUserDonatedNftRecord()
			record.TokenAddr = "bad"
			_, err := mapUserDonatedNft(record)
			return err
		},
		"negative donation index": func() error {
			record := validUserDonatedNftRecord()
			record.DonationIndex = -1
			_, err := mapUserDonatedNft(record)
			return err
		},
		"claim with invalid claimer": func() error {
			record := validUserDonatedNftRecord()
			record.Claimed = true
			claim := validUserDonatedNftClaimRecord()
			claim.ClaimerAddr = "bad"
			record.Claim = &claim
			_, err := mapUserDonatedNft(record)
			return err
		},
		"claim with invalid timestamp": func() error {
			record := validUserDonatedNftRecord()
			record.Claimed = true
			claim := validUserDonatedNftClaimRecord()
			claim.DateTime = "bad"
			record.Claim = &claim
			_, err := mapUserDonatedNft(record)
			return err
		},
	}
	for name, test := range rejections {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := test(); err == nil {
				t.Fatal("mapper accepted invalid record")
			}
		})
	}
}

func validUserDonatedErc20Record() cgstore.UserDonatedErc20Record {
	return cgstore.UserDonatedErc20Record{
		RoundNum:           4,
		TokenAid:           26,
		TokenAddr:          "0x3333333333333333333333333333333333333333",
		DonatedBaseUnits:   "100",
		ClaimedBaseUnits:   "40",
		RemainingBaseUnits: "60",
	}
}

func validUserDonatedErc20ClaimRecord() cgstore.UserDonatedErc20ClaimRecord {
	return cgstore.UserDonatedErc20ClaimRecord{
		EventLogID:      180,
		TxHash:          "0x" + strings.Repeat("AA", 32),
		DateTime:        "2026-01-01T04:01:40+01:00",
		ClaimerAid:      25,
		ClaimerAddr:     "0x6666666666666666666666666666666666666666",
		AmountBaseUnits: "40",
	}
}

func TestMapUserDonatedErc20(t *testing.T) {
	t.Parallel()

	t.Run("partial entitlement with last claim", func(t *testing.T) {
		t.Parallel()
		record := validUserDonatedErc20Record()
		claim := validUserDonatedErc20ClaimRecord()
		record.LastClaim = &claim
		got, err := mapUserDonatedErc20(record)
		if err != nil {
			t.Fatalf("mapUserDonatedErc20: %v", err)
		}
		if got.DonatedBaseUnits != "100" || got.ClaimedBaseUnits != "40" ||
			got.RemainingBaseUnits != "60" || got.Round != 4 ||
			got.TokenAddress != "0x3333333333333333333333333333333333333333" {
			t.Fatalf("summary = %+v", got)
		}
		if got.LastClaim == nil || got.LastClaim.AmountBaseUnits != "40" ||
			got.LastClaim.ClaimerAddress != "0x6666666666666666666666666666666666666666" ||
			!got.LastClaim.OccurredAt.Equal(time.Date(2026, 1, 1, 3, 1, 40, 0, time.UTC)) {
			t.Fatalf("last claim = %+v", got.LastClaim)
		}
	})

	t.Run("untouched entitlement", func(t *testing.T) {
		t.Parallel()
		record := validUserDonatedErc20Record()
		record.ClaimedBaseUnits = "0"
		record.RemainingBaseUnits = "100"
		got, err := mapUserDonatedErc20(record)
		if err != nil || got.LastClaim != nil || got.ClaimedBaseUnits != "0" {
			t.Fatalf("summary = %+v, err=%v", got, err)
		}
	})

	rejections := map[string]func() error{
		"inconsistent totals": func() error {
			record := validUserDonatedErc20Record()
			record.RemainingBaseUnits = "61"
			_, err := mapUserDonatedErc20(record)
			return err
		},
		"negative remaining": func() error {
			record := validUserDonatedErc20Record()
			record.ClaimedBaseUnits = "140"
			record.RemainingBaseUnits = "-40"
			_, err := mapUserDonatedErc20(record)
			return err
		},
		"invalid token address": func() error {
			record := validUserDonatedErc20Record()
			record.TokenAddr = "bad"
			_, err := mapUserDonatedErc20(record)
			return err
		},
		"negative round": func() error {
			record := validUserDonatedErc20Record()
			record.RoundNum = -1
			_, err := mapUserDonatedErc20(record)
			return err
		},
		"malformed donated total": func() error {
			record := validUserDonatedErc20Record()
			record.DonatedBaseUnits = "1e18"
			_, err := mapUserDonatedErc20(record)
			return err
		},
		"claim with malformed amount": func() error {
			record := validUserDonatedErc20Record()
			claim := validUserDonatedErc20ClaimRecord()
			claim.AmountBaseUnits = ""
			record.LastClaim = &claim
			_, err := mapUserDonatedErc20(record)
			return err
		},
		"claim with invalid hash": func() error {
			record := validUserDonatedErc20Record()
			claim := validUserDonatedErc20ClaimRecord()
			claim.TxHash = "bad"
			record.LastClaim = &claim
			_, err := mapUserDonatedErc20(record)
			return err
		},
		"malformed claimed total": func() error {
			record := validUserDonatedErc20Record()
			record.ClaimedBaseUnits = "forty"
			_, err := mapUserDonatedErc20(record)
			return err
		},
		"claim without event id": func() error {
			record := validUserDonatedErc20Record()
			claim := validUserDonatedErc20ClaimRecord()
			claim.EventLogID = 0
			record.LastClaim = &claim
			_, err := mapUserDonatedErc20(record)
			return err
		},
		"claim with invalid claimer": func() error {
			record := validUserDonatedErc20Record()
			claim := validUserDonatedErc20ClaimRecord()
			claim.ClaimerAddr = "bad"
			record.LastClaim = &claim
			_, err := mapUserDonatedErc20(record)
			return err
		},
		"claim with invalid timestamp": func() error {
			record := validUserDonatedErc20Record()
			claim := validUserDonatedErc20ClaimRecord()
			claim.DateTime = "bad"
			record.LastClaim = &claim
			_, err := mapUserDonatedErc20(record)
			return err
		},
	}
	for name, test := range rejections {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := test(); err == nil {
				t.Fatal("mapper accepted invalid record")
			}
		})
	}
}

func TestDonatedErc20TotalsConsistent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		donated, claimed, remaining string
		want                        bool
	}{
		{"100", "40", "60", true},
		{"0", "0", "0", true},
		{"100", "40", "61", false},
		{"junk", "40", "60", false},
		{"100", "junk", "60", false},
		{"100", "40", "junk", false},
	}
	for _, test := range tests {
		if got := donatedErc20TotalsConsistent(test.donated, test.claimed, test.remaining); got != test.want {
			t.Errorf("donatedErc20TotalsConsistent(%q,%q,%q) = %v, want %v",
				test.donated, test.claimed, test.remaining, got, test.want)
		}
	}
}
