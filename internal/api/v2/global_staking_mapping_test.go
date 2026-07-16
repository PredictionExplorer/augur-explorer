package v2

import (
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func validGlobalStakingAction(kind cgstore.UserStakingActionKind) cgstore.GlobalStakingActionRecord {
	record := cgstore.GlobalStakingActionRecord{
		Tx:              validDonationTransaction(),
		Kind:            kind,
		StakerAid:       21,
		StakerAddress:   userCursorAlice,
		ActionID:        7,
		TokenID:         5,
		RoundNum:        2,
		TotalStakedNfts: 3,
	}
	if kind == cgstore.UserStakingActionUnstake {
		record.RewardWei = "100"
		record.RewardPerTokenWei = "25"
	}
	return record
}

func TestMapGlobalCstStakingAction(t *testing.T) {
	t.Parallel()
	stake, err := mapGlobalCstStakingAction(
		validGlobalStakingAction(cgstore.UserStakingActionStake),
	)
	if err != nil {
		t.Fatal(err)
	}
	if stake.ActionType != Stake || stake.RewardWei != nil ||
		stake.RewardPerTokenWei != nil || stake.Round != 2 ||
		stake.StakerAddress != userCursorAlice {
		t.Fatalf("stake = %+v", stake)
	}
	unstake, err := mapGlobalCstStakingAction(
		validGlobalStakingAction(cgstore.UserStakingActionUnstake),
	)
	if err != nil {
		t.Fatal(err)
	}
	if unstake.ActionType != Unstake || unstake.RewardWei == nil ||
		*unstake.RewardWei != "100" || unstake.RewardPerTokenWei == nil ||
		*unstake.RewardPerTokenWei != "25" {
		t.Fatalf("unstake = %+v", unstake)
	}

	invalid := map[string]func(*cgstore.GlobalStakingActionRecord){
		"negative action":    func(r *cgstore.GlobalStakingActionRecord) { r.ActionID = -1 },
		"negative token":     func(r *cgstore.GlobalStakingActionRecord) { r.TokenID = -1 },
		"negative round":     func(r *cgstore.GlobalStakingActionRecord) { r.RoundNum = -1 },
		"negative pool size": func(r *cgstore.GlobalStakingActionRecord) { r.TotalStakedNfts = -1 },
		"unknown kind":       func(r *cgstore.GlobalStakingActionRecord) { r.Kind = "restake" },
		"bad address":        func(r *cgstore.GlobalStakingActionRecord) { r.StakerAddress = "bad" },
		"bad transaction":    func(r *cgstore.GlobalStakingActionRecord) { r.Tx.TxHash = "bad" },
		"reward on stake":    func(r *cgstore.GlobalStakingActionRecord) { r.RewardWei = "1" },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validGlobalStakingAction(cgstore.UserStakingActionStake)
			corrupt(&record)
			if _, err := mapGlobalCstStakingAction(record); err == nil {
				t.Fatal("corrupt record mapped")
			}
		})
	}
	for name, corrupt := range map[string]func(*cgstore.GlobalStakingActionRecord){
		"empty reward":           func(r *cgstore.GlobalStakingActionRecord) { r.RewardWei = "" },
		"negative reward":        func(r *cgstore.GlobalStakingActionRecord) { r.RewardWei = "-1" },
		"empty per-token reward": func(r *cgstore.GlobalStakingActionRecord) { r.RewardPerTokenWei = "" },
		"bad per-token reward":   func(r *cgstore.GlobalStakingActionRecord) { r.RewardPerTokenWei = "1.5" },
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validGlobalStakingAction(cgstore.UserStakingActionUnstake)
			corrupt(&record)
			if _, err := mapGlobalCstStakingAction(record); err == nil {
				t.Fatal("corrupt unstake mapped")
			}
		})
	}
}

func TestMapGlobalRandomWalkStakingAction(t *testing.T) {
	t.Parallel()
	for _, kind := range []cgstore.UserStakingActionKind{
		cgstore.UserStakingActionStake,
		cgstore.UserStakingActionUnstake,
	} {
		record := validGlobalStakingAction(kind)
		record.RewardWei = ""
		record.RewardPerTokenWei = ""
		got, err := mapGlobalRandomWalkStakingAction(record)
		if err != nil {
			t.Fatalf("%s: %v", kind, err)
		}
		if string(got.ActionType) != string(kind) || got.NftTokenId != 5 {
			t.Fatalf("%s = %+v", kind, got)
		}
	}
	record := validGlobalStakingAction(cgstore.UserStakingActionUnstake)
	if _, err := mapGlobalRandomWalkStakingAction(record); err == nil {
		t.Fatal("RandomWalk reward fields accepted")
	}
	for name, corrupt := range map[string]func(*cgstore.GlobalStakingActionRecord){
		"invalid identity": func(r *cgstore.GlobalStakingActionRecord) { r.ActionID = -1 },
		"unknown kind":     func(r *cgstore.GlobalStakingActionRecord) { r.Kind = "restake" },
		"bad transaction":  func(r *cgstore.GlobalStakingActionRecord) { r.Tx.TxHash = "bad" },
		"bad address":      func(r *cgstore.GlobalStakingActionRecord) { r.StakerAddress = "bad" },
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validRwalkGlobalAction(cgstore.UserStakingActionStake)
			corrupt(&record)
			if _, err := mapGlobalRandomWalkStakingAction(record); err == nil {
				t.Fatal("corrupt RandomWalk action mapped")
			}
		})
	}
}

func validStakeUnstakeCombined(closed bool, withReward bool) cgmodel.CGStakeUnstakeCombined {
	stakeTx := validDonationTransaction()
	stake := cgmodel.CGStakeActionInfoRec{
		Tx:            stakeTx,
		ActionId:      7,
		TokenId:       5,
		RoundNum:      2,
		NumStakedNFTs: 3,
		StakerAid:     21,
		StakerAddr:    userCursorAlice,
	}
	result := cgmodel.CGStakeUnstakeCombined{Stake: stake}
	if closed {
		unstakeTx := stakeTx
		unstakeTx.EvtLogId++
		unstakeTx.TimeStamp++
		result.Unstake = cgmodel.CGUnstakeActionInfoRec{
			Tx:             unstakeTx,
			ActionId:       7,
			TokenId:        5,
			RoundNum:       3,
			NumStakedNFTs:  2,
			StakerAid:      21,
			StakerAddr:     userCursorAlice,
			RewardAmount:   "100",
			RewardPerToken: "25",
		}
		if !withReward {
			result.Unstake.RewardAmount = ""
			result.Unstake.RewardPerToken = ""
		}
	}
	return result
}

func TestMapGlobalStakingActionDetails(t *testing.T) {
	t.Parallel()
	open, err := mapGlobalCstStakingActionDetail(validStakeUnstakeCombined(false, true))
	if err != nil || open.Unstake != nil {
		t.Fatalf("open CST detail = %+v, %v", open, err)
	}
	closed, err := mapGlobalCstStakingActionDetail(validStakeUnstakeCombined(true, true))
	if err != nil || closed.Unstake == nil || closed.Unstake.RewardWei == nil {
		t.Fatalf("closed CST detail = %+v, %v", closed, err)
	}
	rwalk, err := mapGlobalRandomWalkStakingActionDetail(validStakeUnstakeCombined(true, false))
	if err != nil || rwalk.Unstake == nil {
		t.Fatalf("closed RandomWalk detail = %+v, %v", rwalk, err)
	}

	for name, corrupt := range map[string]func(*cgmodel.CGStakeUnstakeCombined){
		"action mismatch":  func(r *cgmodel.CGStakeUnstakeCombined) { r.Unstake.ActionId++ },
		"token mismatch":   func(r *cgmodel.CGStakeUnstakeCombined) { r.Unstake.TokenId++ },
		"staker mismatch":  func(r *cgmodel.CGStakeUnstakeCombined) { r.Unstake.StakerAid++ },
		"address mismatch": func(r *cgmodel.CGStakeUnstakeCombined) { r.Unstake.StakerAddr = userCursorBob },
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validStakeUnstakeCombined(true, true)
			corrupt(&record)
			if _, err := mapGlobalCstStakingActionDetail(record); err == nil {
				t.Fatal("inconsistent lifecycle mapped")
			}
		})
	}
	t.Run("invalid CST stake", func(t *testing.T) {
		t.Parallel()
		record := validStakeUnstakeCombined(false, true)
		record.Stake.StakerAddr = "bad"
		if _, err := mapGlobalCstStakingActionDetail(record); err == nil {
			t.Fatal("invalid CST stake mapped")
		}
	})
	t.Run("invalid CST unstake", func(t *testing.T) {
		t.Parallel()
		record := validStakeUnstakeCombined(true, true)
		record.Unstake.RewardAmount = "bad"
		if _, err := mapGlobalCstStakingActionDetail(record); err == nil {
			t.Fatal("invalid CST unstake mapped")
		}
	})
	t.Run("invalid RandomWalk stake", func(t *testing.T) {
		t.Parallel()
		record := validStakeUnstakeCombined(false, false)
		record.Stake.StakerAddr = "bad"
		if _, err := mapGlobalRandomWalkStakingActionDetail(record); err == nil {
			t.Fatal("invalid RandomWalk stake mapped")
		}
	})
	t.Run("inconsistent RandomWalk unstake", func(t *testing.T) {
		t.Parallel()
		record := validStakeUnstakeCombined(true, false)
		record.Unstake.TokenId++
		if _, err := mapGlobalRandomWalkStakingActionDetail(record); err == nil {
			t.Fatal("inconsistent RandomWalk unstake mapped")
		}
	})
	t.Run("invalid RandomWalk unstake", func(t *testing.T) {
		t.Parallel()
		record := validStakeUnstakeCombined(true, false)
		record.Unstake.Tx.TxHash = "bad"
		if _, err := mapGlobalRandomWalkStakingActionDetail(record); err == nil {
			t.Fatal("invalid RandomWalk unstake mapped")
		}
	})
}

func TestMapGlobalStakedTokens(t *testing.T) {
	t.Parallel()
	cstRecord := cgstore.GlobalStakedCstTokenRecord{
		StakeTx:       validDonationTransaction(),
		StakerAid:     21,
		StakerAddress: userCursorAlice,
		ActionID:      7,
		TokenID:       5,
		MintRound:     2,
		Seed:          "seed05",
		TokenName:     "Genesis",
	}
	cst, err := mapGlobalStakedCstToken(cstRecord)
	if err != nil || cst.TokenName == nil || *cst.TokenName != "Genesis" ||
		cst.StakerAddress != userCursorAlice {
		t.Fatalf("CST token = %+v, %v", cst, err)
	}
	cstRecord.TokenName = ""
	cst, err = mapGlobalStakedCstToken(cstRecord)
	if err != nil || cst.TokenName != nil {
		t.Fatalf("unnamed CST token = %+v, %v", cst, err)
	}
	cstRecord.Seed = ""
	if _, err := mapGlobalStakedCstToken(cstRecord); err == nil {
		t.Fatal("missing CST seed accepted")
	}
	for name, corrupt := range map[string]func(*cgstore.GlobalStakedCstTokenRecord){
		"invalid identity": func(r *cgstore.GlobalStakedCstTokenRecord) { r.ActionID = -1 },
		"bad transaction":  func(r *cgstore.GlobalStakedCstTokenRecord) { r.StakeTx.TxHash = "bad" },
		"bad address":      func(r *cgstore.GlobalStakedCstTokenRecord) { r.StakerAddress = "bad" },
	} {
		t.Run("CST "+name, func(t *testing.T) {
			t.Parallel()
			record := cgstore.GlobalStakedCstTokenRecord{
				StakeTx:       validDonationTransaction(),
				StakerAid:     21,
				StakerAddress: userCursorAlice,
				ActionID:      7,
				TokenID:       5,
				MintRound:     2,
				Seed:          "seed05",
			}
			corrupt(&record)
			if _, err := mapGlobalStakedCstToken(record); err == nil {
				t.Fatal("corrupt CST token mapped")
			}
		})
	}

	rwalkRecord := cgstore.GlobalStakedRwalkTokenRecord{
		StakeTx:       validDonationTransaction(),
		StakerAid:     21,
		StakerAddress: userCursorAlice,
		ActionID:      7,
		TokenID:       5,
	}
	rwalk, err := mapGlobalStakedRandomWalkToken(rwalkRecord)
	if err != nil || rwalk.NftTokenId != 5 || rwalk.StakerAddress != userCursorAlice {
		t.Fatalf("RandomWalk token = %+v, %v", rwalk, err)
	}
	rwalkRecord.StakerAddress = "bad"
	if _, err := mapGlobalStakedRandomWalkToken(rwalkRecord); err == nil {
		t.Fatal("bad RandomWalk staker accepted")
	}
	rwalkRecord.StakerAddress = userCursorAlice
	rwalkRecord.ActionID = -1
	if _, err := mapGlobalStakedRandomWalkToken(rwalkRecord); err == nil {
		t.Fatal("invalid RandomWalk identity accepted")
	}
	rwalkRecord.ActionID = 7
	rwalkRecord.StakeTx.TxHash = "bad"
	if _, err := mapGlobalStakedRandomWalkToken(rwalkRecord); err == nil {
		t.Fatal("invalid RandomWalk transaction accepted")
	}
}

func validGlobalStakingDepositRecord() cgstore.GlobalStakingDepositRecord {
	return cgstore.GlobalStakingDepositRecord{
		Tx:                 validDonationTransaction(),
		DepositID:          8,
		RoundNum:           2,
		TotalDepositWei:    "101",
		TotalStakedNfts:    4,
		AmountPerTokenWei:  "25",
		CollectedWei:       "50",
		PendingWei:         "50",
		RemainderWei:       "1",
		RewardCount:        4,
		PendingRewardCount: 2,
	}
}

func TestMapGlobalStakingDeposit(t *testing.T) {
	t.Parallel()
	got, err := mapGlobalStakingDeposit(validGlobalStakingDepositRecord())
	if err != nil {
		t.Fatal(err)
	}
	if got.TotalDepositWei != "101" || got.CollectedWei != "50" ||
		got.PendingWei != "50" || got.RemainderWei != "1" || got.FullyClaimed {
		t.Fatalf("deposit = %+v", got)
	}
	claimed := validGlobalStakingDepositRecord()
	claimed.CollectedWei = "100"
	claimed.PendingWei = "0"
	claimed.PendingRewardCount = 0
	got, err = mapGlobalStakingDeposit(claimed)
	if err != nil || !got.FullyClaimed {
		t.Fatalf("claimed deposit = %+v, %v", got, err)
	}
	for name, corrupt := range map[string]func(*cgstore.GlobalStakingDepositRecord){
		"negative id":        func(r *cgstore.GlobalStakingDepositRecord) { r.DepositID = -1 },
		"zero token count":   func(r *cgstore.GlobalStakingDepositRecord) { r.TotalStakedNfts = 0 },
		"bad reward count":   func(r *cgstore.GlobalStakingDepositRecord) { r.PendingRewardCount = 5 },
		"bad transaction":    func(r *cgstore.GlobalStakingDepositRecord) { r.Tx.TxHash = "bad" },
		"bad amount":         func(r *cgstore.GlobalStakingDepositRecord) { r.TotalDepositWei = "x" },
		"bad per-token":      func(r *cgstore.GlobalStakingDepositRecord) { r.AmountPerTokenWei = "x" },
		"bad collected":      func(r *cgstore.GlobalStakingDepositRecord) { r.CollectedWei = "x" },
		"negative amount":    func(r *cgstore.GlobalStakingDepositRecord) { r.PendingWei = "-1" },
		"bad remainder":      func(r *cgstore.GlobalStakingDepositRecord) { r.RemainderWei = "x" },
		"closure divergence": func(r *cgstore.GlobalStakingDepositRecord) { r.RemainderWei = "2" },
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validGlobalStakingDepositRecord()
			corrupt(&record)
			if _, err := mapGlobalStakingDeposit(record); err == nil {
				t.Fatal("corrupt deposit mapped")
			}
		})
	}
}

func TestMapRoundStakingReward(t *testing.T) {
	t.Parallel()
	record := cgstore.RoundStakingRewardRecord{
		DepositID:      8,
		RoundNum:       2,
		StakerAid:      21,
		StakerAddress:  userCursorAlice,
		StakedNftCount: 2,
		RewardWei:      "100",
		CollectedWei:   "60",
		PendingWei:     "40",
	}
	got, err := mapRoundStakingReward(record)
	if err != nil || got.FullyClaimed || got.StakerAddress != userCursorAlice {
		t.Fatalf("reward = %+v, %v", got, err)
	}
	record.CollectedWei = "100"
	record.PendingWei = "0"
	got, err = mapRoundStakingReward(record)
	if err != nil || !got.FullyClaimed {
		t.Fatalf("claimed reward = %+v, %v", got, err)
	}
	for name, corrupt := range map[string]func(*cgstore.RoundStakingRewardRecord){
		"negative deposit":   func(r *cgstore.RoundStakingRewardRecord) { r.DepositID = -1 },
		"zero staker":        func(r *cgstore.RoundStakingRewardRecord) { r.StakerAid = 0 },
		"zero token count":   func(r *cgstore.RoundStakingRewardRecord) { r.StakedNftCount = 0 },
		"bad address":        func(r *cgstore.RoundStakingRewardRecord) { r.StakerAddress = "bad" },
		"bad amount":         func(r *cgstore.RoundStakingRewardRecord) { r.RewardWei = "x" },
		"bad collected":      func(r *cgstore.RoundStakingRewardRecord) { r.CollectedWei = "x" },
		"bad pending":        func(r *cgstore.RoundStakingRewardRecord) { r.PendingWei = "x" },
		"closure divergence": func(r *cgstore.RoundStakingRewardRecord) { r.PendingWei = "41" },
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			corruptRecord := record
			corruptRecord.CollectedWei = "60"
			corruptRecord.PendingWei = "40"
			corrupt(&corruptRecord)
			if _, err := mapRoundStakingReward(corruptRecord); err == nil {
				t.Fatal("corrupt reward mapped")
			}
		})
	}
}

func TestAmountEqualsSum(t *testing.T) {
	t.Parallel()
	if !amountEqualsSum("10", "3", "2", "5") {
		t.Fatal("valid sum rejected")
	}
	for _, parts := range [][]string{
		{"10", "3", "8"},
		{"x", "1"},
		{"10", "x"},
	} {
		if amountEqualsSum(parts[0], parts[1:]...) {
			t.Fatalf("invalid sum accepted: %v", parts)
		}
	}
}
