package v2

import (
	"testing"
	"time"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func validUserStakingActionRecord(kind cgstore.UserStakingActionKind) cgstore.UserStakingActionRecord {
	record := cgstore.UserStakingActionRecord{
		Tx:              validDonationTransaction(),
		Kind:            kind,
		StakerAid:       1,
		ActionID:        2,
		TokenID:         5,
		TotalStakedNfts: 3,
	}
	if kind == cgstore.UserStakingActionUnstake {
		record.RewardWei = "1000000000000000000"
	}
	return record
}

func TestMapUserCstStakingAction(t *testing.T) {
	t.Parallel()

	t.Run("stake", func(t *testing.T) {
		t.Parallel()
		got, err := mapUserCstStakingAction(validUserStakingActionRecord(cgstore.UserStakingActionStake))
		if err != nil {
			t.Fatalf("mapUserCstStakingAction: %v", err)
		}
		if got.ActionType != Stake || got.RewardWei != nil || got.ActionId != 2 ||
			got.NftTokenId != 5 || got.TotalStakedNfts != 3 || got.EventLogId != 100 ||
			got.OccurredAt.UTC().Format(time.RFC3339) != "2026-01-01T00:01:40Z" {
			t.Fatalf("stake action = %+v", got)
		}
	})

	t.Run("unstake carries the exact reward", func(t *testing.T) {
		t.Parallel()
		got, err := mapUserCstStakingAction(validUserStakingActionRecord(cgstore.UserStakingActionUnstake))
		if err != nil {
			t.Fatalf("mapUserCstStakingAction: %v", err)
		}
		if got.ActionType != Unstake || got.RewardWei == nil || *got.RewardWei != "1000000000000000000" {
			t.Fatalf("unstake action = %+v", got)
		}
	})

	t.Run("zero-reward unstake is legal", func(t *testing.T) {
		t.Parallel()
		record := validUserStakingActionRecord(cgstore.UserStakingActionUnstake)
		record.RewardWei = "0"
		got, err := mapUserCstStakingAction(record)
		if err != nil || got.RewardWei == nil || *got.RewardWei != "0" {
			t.Fatalf("zero-reward unstake = %+v, err=%v", got, err)
		}
	})

	invalid := map[string]func(*cgstore.UserStakingActionRecord){
		"unknown kind":         func(r *cgstore.UserStakingActionRecord) { r.Kind = "burn" },
		"negative action id":   func(r *cgstore.UserStakingActionRecord) { r.ActionID = -1 },
		"negative token id":    func(r *cgstore.UserStakingActionRecord) { r.TokenID = -1 },
		"negative pool count":  func(r *cgstore.UserStakingActionRecord) { r.TotalStakedNfts = -1 },
		"stake with reward":    func(r *cgstore.UserStakingActionRecord) { r.RewardWei = "5" },
		"invalid event id":     func(r *cgstore.UserStakingActionRecord) { r.Tx.EvtLogId = 0 },
		"invalid hash":         func(r *cgstore.UserStakingActionRecord) { r.Tx.TxHash = "0xnope" },
		"unparseable datetime": func(r *cgstore.UserStakingActionRecord) { r.Tx.DateTime = "yesterday" },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validUserStakingActionRecord(cgstore.UserStakingActionStake)
			corrupt(&record)
			if _, err := mapUserCstStakingAction(record); err == nil {
				t.Fatalf("mapUserCstStakingAction accepted %s", name)
			}
		})
	}

	t.Run("unstake without reward", func(t *testing.T) {
		t.Parallel()
		record := validUserStakingActionRecord(cgstore.UserStakingActionUnstake)
		record.RewardWei = ""
		if _, err := mapUserCstStakingAction(record); err == nil {
			t.Fatal("mapUserCstStakingAction accepted a rewardless unstake")
		}
	})
	t.Run("unstake with malformed reward", func(t *testing.T) {
		t.Parallel()
		record := validUserStakingActionRecord(cgstore.UserStakingActionUnstake)
		record.RewardWei = "-3"
		if _, err := mapUserCstStakingAction(record); err == nil {
			t.Fatal("mapUserCstStakingAction accepted a negative reward")
		}
	})
}

func TestMapUserRandomWalkStakingAction(t *testing.T) {
	t.Parallel()

	record := validUserStakingActionRecord(cgstore.UserStakingActionUnstake)
	record.RewardWei = ""
	got, err := mapUserRandomWalkStakingAction(record)
	if err != nil {
		t.Fatalf("mapUserRandomWalkStakingAction: %v", err)
	}
	if got.ActionType != Unstake || got.ActionId != 2 || got.NftTokenId != 5 {
		t.Fatalf("action = %+v", got)
	}

	withReward := validUserStakingActionRecord(cgstore.UserStakingActionUnstake)
	if _, err := mapUserRandomWalkStakingAction(withReward); err == nil {
		t.Fatal("mapUserRandomWalkStakingAction accepted a reward-carrying row")
	}
	unknown := record
	unknown.Kind = "burn"
	if _, err := mapUserRandomWalkStakingAction(unknown); err == nil {
		t.Fatal("mapUserRandomWalkStakingAction accepted an unknown kind")
	}
	badTx := record
	badTx.Tx.EvtLogId = 0
	if _, err := mapUserRandomWalkStakingAction(badTx); err == nil {
		t.Fatal("mapUserRandomWalkStakingAction accepted an invalid transaction")
	}
	negative := record
	negative.TokenID = -1
	if _, err := mapUserRandomWalkStakingAction(negative); err == nil {
		t.Fatal("mapUserRandomWalkStakingAction accepted a negative token id")
	}
}

func validUserStakedCstTokenRecord() cgstore.UserStakedCstTokenRecord {
	return cgstore.UserStakedCstTokenRecord{
		StakeTx:   validDonationTransaction(),
		StakerAid: 1,
		ActionID:  2,
		TokenID:   5,
		MintRound: 0,
		Seed:      "seed0000000000000000000000000000000000000000000000000000000005",
	}
}

func TestMapUserStakedCstToken(t *testing.T) {
	t.Parallel()

	t.Run("unnamed token", func(t *testing.T) {
		t.Parallel()
		got, err := mapUserStakedCstToken(validUserStakedCstTokenRecord())
		if err != nil {
			t.Fatalf("mapUserStakedCstToken: %v", err)
		}
		if got.NftTokenId != 5 || got.ActionId != 2 || got.MintRound != 0 ||
			got.TokenName != nil || got.Seed == "" || got.EventLogId != 100 {
			t.Fatalf("staked token = %+v", got)
		}
	})

	t.Run("named token", func(t *testing.T) {
		t.Parallel()
		record := validUserStakedCstTokenRecord()
		record.TokenName = "Genesis"
		got, err := mapUserStakedCstToken(record)
		if err != nil || got.TokenName == nil || *got.TokenName != "Genesis" {
			t.Fatalf("named token = %+v, err=%v", got, err)
		}
	})

	invalid := map[string]func(*cgstore.UserStakedCstTokenRecord){
		"negative action id": func(r *cgstore.UserStakedCstTokenRecord) { r.ActionID = -1 },
		"negative token id":  func(r *cgstore.UserStakedCstTokenRecord) { r.TokenID = -1 },
		"negative round":     func(r *cgstore.UserStakedCstTokenRecord) { r.MintRound = -1 },
		"missing seed":       func(r *cgstore.UserStakedCstTokenRecord) { r.Seed = "" },
		"invalid hash":       func(r *cgstore.UserStakedCstTokenRecord) { r.StakeTx.TxHash = "0xnope" },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validUserStakedCstTokenRecord()
			corrupt(&record)
			if _, err := mapUserStakedCstToken(record); err == nil {
				t.Fatalf("mapUserStakedCstToken accepted %s", name)
			}
		})
	}
}

func TestMapUserStakedRandomWalkToken(t *testing.T) {
	t.Parallel()

	record := cgstore.UserStakedRwalkTokenRecord{
		StakeTx:   validDonationTransaction(),
		StakerAid: 1,
		ActionID:  103,
		TokenID:   13,
	}
	got, err := mapUserStakedRandomWalkToken(record)
	if err != nil || got.NftTokenId != 13 || got.ActionId != 103 {
		t.Fatalf("staked token = %+v, err=%v", got, err)
	}

	negative := record
	negative.TokenID = -1
	if _, err := mapUserStakedRandomWalkToken(negative); err == nil {
		t.Fatal("mapUserStakedRandomWalkToken accepted a negative token id")
	}
	badTx := record
	badTx.StakeTx.EvtLogId = 0
	if _, err := mapUserStakedRandomWalkToken(badTx); err == nil {
		t.Fatal("mapUserStakedRandomWalkToken accepted an invalid transaction")
	}
}

func validUserStakingDepositRecord() cgstore.UserStakingDepositRecord {
	return cgstore.UserStakingDepositRecord{
		Tx:                 validDonationTransaction(),
		StakerAid:          1,
		DepositID:          501,
		RoundNum:           0,
		TotalDepositWei:    "2000000000000000000",
		TotalStakedNfts:    2,
		AmountPerTokenWei:  "1000000000000000000",
		StakedNftCount:     1,
		AmountDepositedWei: "1000000000000000000",
		AmountToClaimWei:   "0",
		ClaimedRewardWei:   "1000000000000000000",
		PendingRewardWei:   "0",
		ClaimedNftCount:    1,
		PendingNftCount:    0,
	}
}

func TestMapUserStakingDeposit(t *testing.T) {
	t.Parallel()

	t.Run("fully claimed", func(t *testing.T) {
		t.Parallel()
		got, err := mapUserStakingDeposit(validUserStakingDepositRecord())
		if err != nil {
			t.Fatalf("mapUserStakingDeposit: %v", err)
		}
		if got.DepositId != 501 || !got.FullyClaimed ||
			got.RewardWei != "1000000000000000000" ||
			got.ClaimedWei != "1000000000000000000" || got.PendingWei != "0" ||
			got.ClaimedNftCount != 1 || got.PendingNftCount != 0 ||
			got.TotalDepositWei != "2000000000000000000" ||
			got.AmountPerTokenWei != "1000000000000000000" {
			t.Fatalf("deposit = %+v", got)
		}
	})

	t.Run("partially pending", func(t *testing.T) {
		t.Parallel()
		record := validUserStakingDepositRecord()
		record.StakedNftCount = 2
		record.AmountDepositedWei = "2000000000000000000"
		record.AmountToClaimWei = "1000000000000000000"
		record.PendingRewardWei = "1000000000000000000"
		record.PendingNftCount = 1
		got, err := mapUserStakingDeposit(record)
		if err != nil || got.FullyClaimed || got.PendingWei != "1000000000000000000" {
			t.Fatalf("deposit = %+v, err=%v", got, err)
		}
	})

	invalid := map[string]func(*cgstore.UserStakingDepositRecord){
		"negative deposit id": func(r *cgstore.UserStakingDepositRecord) { r.DepositID = -1 },
		"negative round":      func(r *cgstore.UserStakingDepositRecord) { r.RoundNum = -1 },
		"empty pool":          func(r *cgstore.UserStakingDepositRecord) { r.TotalStakedNfts = 0 },
		"no staked tokens":    func(r *cgstore.UserStakingDepositRecord) { r.StakedNftCount = 0 },
		"count mismatch":      func(r *cgstore.UserStakingDepositRecord) { r.ClaimedNftCount = 2 },
		"totals diverge": func(r *cgstore.UserStakingDepositRecord) {
			r.ClaimedRewardWei = "999"
		},
		"pending accumulator diverges": func(r *cgstore.UserStakingDepositRecord) {
			r.AmountToClaimWei = "7"
		},
		"malformed reward": func(r *cgstore.UserStakingDepositRecord) {
			r.AmountDepositedWei = "not-a-number"
		},
		"malformed pool amount": func(r *cgstore.UserStakingDepositRecord) {
			r.TotalDepositWei = ""
		},
		"malformed per-token amount": func(r *cgstore.UserStakingDepositRecord) {
			r.AmountPerTokenWei = "-1"
		},
		"malformed claimed sum": func(r *cgstore.UserStakingDepositRecord) {
			r.ClaimedRewardWei = "x"
		},
		"malformed pending sum": func(r *cgstore.UserStakingDepositRecord) {
			r.PendingRewardWei = "x"
		},
		"malformed pending accumulator": func(r *cgstore.UserStakingDepositRecord) {
			r.AmountToClaimWei = "x"
		},
		"invalid transaction": func(r *cgstore.UserStakingDepositRecord) { r.Tx.EvtLogId = 0 },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validUserStakingDepositRecord()
			corrupt(&record)
			if _, err := mapUserStakingDeposit(record); err == nil {
				t.Fatalf("mapUserStakingDeposit accepted %s", name)
			}
		})
	}
}

func TestAmountsAddUp(t *testing.T) {
	t.Parallel()

	cases := []struct {
		total, left, right string
		want               bool
	}{
		{"4", "1", "3", true},
		{"0", "0", "0", true},
		{"4", "2", "3", false},
		{"x", "1", "3", false},
		{"4", "x", "3", false},
		{"4", "1", "x", false},
	}
	for _, test := range cases {
		if got := amountsAddUp(test.total, test.left, test.right); got != test.want {
			t.Errorf("amountsAddUp(%q,%q,%q) = %v, want %v",
				test.total, test.left, test.right, got, test.want)
		}
	}
}

func TestMapUserStakingDepositReward(t *testing.T) {
	t.Parallel()

	record := cgstore.UserStakingDepositRewardRecord{
		StakerAid: 1,
		ActionID:  2,
		TokenID:   5,
		RewardWei: "1000000000000000000",
		Claimed:   false,
	}
	got, err := mapUserStakingDepositReward(record)
	if err != nil || got.ActionId != 2 || got.NftTokenId != 5 ||
		got.RewardWei != "1000000000000000000" || got.Claimed {
		t.Fatalf("reward = %+v, err=%v", got, err)
	}

	negative := record
	negative.ActionID = -1
	if _, err := mapUserStakingDepositReward(negative); err == nil {
		t.Fatal("mapUserStakingDepositReward accepted a negative action id")
	}
	malformed := record
	malformed.RewardWei = ""
	if _, err := mapUserStakingDepositReward(malformed); err == nil {
		t.Fatal("mapUserStakingDepositReward accepted an empty reward")
	}
}

func TestMapUserStakingTokenReward(t *testing.T) {
	t.Parallel()

	record := cgstore.UserStakingTokenRewardRecord{
		TokenID:      5,
		TotalWei:     "4000000000000000000",
		CollectedWei: "1000000000000000000",
		PendingWei:   "3000000000000000000",
	}
	got, err := mapUserStakingTokenReward(record)
	if err != nil || got.NftTokenId != 5 || got.TotalWei != "4000000000000000000" ||
		got.CollectedWei != "1000000000000000000" || got.PendingWei != "3000000000000000000" {
		t.Fatalf("token reward = %+v, err=%v", got, err)
	}

	invalid := map[string]func(*cgstore.UserStakingTokenRewardRecord){
		"negative token id":   func(r *cgstore.UserStakingTokenRewardRecord) { r.TokenID = -1 },
		"inconsistent totals": func(r *cgstore.UserStakingTokenRewardRecord) { r.TotalWei = "5" },
		"malformed total":     func(r *cgstore.UserStakingTokenRewardRecord) { r.TotalWei = "x" },
		"malformed collected": func(r *cgstore.UserStakingTokenRewardRecord) { r.CollectedWei = "" },
		"malformed pending":   func(r *cgstore.UserStakingTokenRewardRecord) { r.PendingWei = "-1" },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			bad := record
			corrupt(&bad)
			if _, err := mapUserStakingTokenReward(bad); err == nil {
				t.Fatalf("mapUserStakingTokenReward accepted %s", name)
			}
		})
	}
}

func TestMapUserStakingTokenRewardDeposit(t *testing.T) {
	t.Parallel()

	record := cgstore.UserStakingTokenRewardDepositRecord{
		Tx:        validDonationTransaction(),
		StakerAid: 1,
		DepositID: 501,
		RoundNum:  0,
		RewardWei: "1000000000000000000",
		Claimed:   true,
	}
	got, err := mapUserStakingTokenRewardDeposit(record)
	if err != nil || got.DepositId != 501 || got.Round != 0 || !got.Claimed ||
		got.RewardWei != "1000000000000000000" || got.EventLogId != 100 {
		t.Fatalf("token deposit = %+v, err=%v", got, err)
	}

	negative := record
	negative.DepositID = -1
	if _, err := mapUserStakingTokenRewardDeposit(negative); err == nil {
		t.Fatal("mapUserStakingTokenRewardDeposit accepted a negative deposit id")
	}
	badTx := record
	badTx.Tx.DateTime = "yesterday"
	if _, err := mapUserStakingTokenRewardDeposit(badTx); err == nil {
		t.Fatal("mapUserStakingTokenRewardDeposit accepted an invalid transaction")
	}
	malformed := record
	malformed.RewardWei = "x"
	if _, err := mapUserStakingTokenRewardDeposit(malformed); err == nil {
		t.Fatal("mapUserStakingTokenRewardDeposit accepted a malformed reward")
	}
}
