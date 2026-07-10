package v2

import (
	"encoding/json"
	"strings"
	"testing"

	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestMapRoundRaffleEthDeposit(t *testing.T) {
	t.Parallel()

	for _, claimed := range []bool{false, true} {
		record := validRaffleEthDepositRecord()
		record.Claimed = claimed
		got, err := mapRoundRaffleEthDeposit(record)
		if err != nil {
			t.Fatalf("mapRoundRaffleEthDeposit: %v", err)
		}
		if got.Round != 3 ||
			got.WinnerIndex != 2 ||
			got.EventLogId != 7001 ||
			got.BlockNumber != 123 ||
			got.EthAmountWei != "50000000000000000" ||
			got.Claimed != claimed {
			t.Fatalf("mapped deposit = %+v", got)
		}
		if got.WinnerAddress != "0x2100000000000000000000000000000000000021" {
			t.Errorf("WinnerAddress = %q", got.WinnerAddress)
		}
		if got.OccurredAt.String() != "2026-01-01 05:41:40 +0000 UTC" {
			t.Errorf("OccurredAt = %s", got.OccurredAt)
		}
	}
}

func TestMapRoundRaffleNftWinnerPools(t *testing.T) {
	t.Parallel()

	for _, staker := range []bool{false, true} {
		record := validRaffleNftWinnerRecord(staker)
		got, err := mapRoundRaffleNftWinner(record, staker)
		if err != nil {
			t.Fatalf("mapRoundRaffleNftWinner(staker=%v): %v", staker, err)
		}
		if got.Round != 3 ||
			got.WinnerIndex != 2 ||
			got.EventLogId != 7002 ||
			got.NftTokenId != 7 ||
			got.CstAmountWei != "30000000000000000000" ||
			got.IsRandomWalk != staker {
			t.Fatalf("mapped winner = %+v", got)
		}
		if got.WinnerAddress != "0x2100000000000000000000000000000000000021" {
			t.Errorf("WinnerAddress = %q", got.WinnerAddress)
		}
	}
}

func TestRaffleMappingsOmitLegacyFields(t *testing.T) {
	t.Parallel()

	deposit, err := mapRoundRaffleEthDeposit(validRaffleEthDepositRecord())
	if err != nil {
		t.Fatal(err)
	}
	winner, err := mapRoundRaffleNftWinner(validRaffleNftWinnerRecord(false), false)
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal([]any{deposit, winner})
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"RecordId", "recordId", "WinnerAid", "winnerAid", "AmountEth",
		"amountEth", "CstAmountEth", "cstAmountEth", "isStaker",
	} {
		if strings.Contains(string(encoded), forbidden) {
			t.Errorf("raffle JSON leaked %q: %s", forbidden, encoded)
		}
	}
}

func TestMapRoundRaffleEthDepositRejectsMalformedData(t *testing.T) {
	t.Parallel()

	tests := map[string]func(*cgstore.RaffleEthDepositRecord){
		"negative round":  func(r *cgstore.RaffleEthDepositRecord) { r.RoundNum = -1 },
		"negative winner": func(r *cgstore.RaffleEthDepositRecord) { r.WinnerIndex = -1 },
		"event identity":  func(r *cgstore.RaffleEthDepositRecord) { r.Tx.EvtLogId = 0 },
		"hash":            func(r *cgstore.RaffleEthDepositRecord) { r.Tx.TxHash = "bad" },
		"timestamp":       func(r *cgstore.RaffleEthDepositRecord) { r.Tx.DateTime = "bad" },
		"address":         func(r *cgstore.RaffleEthDepositRecord) { r.WinnerAddr = "bad" },
		"empty amount":    func(r *cgstore.RaffleEthDepositRecord) { r.EthAmountWei = "" },
		"negative amount": func(r *cgstore.RaffleEthDepositRecord) { r.EthAmountWei = "-1" },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validRaffleEthDepositRecord()
			mutate(&record)
			if _, err := mapRoundRaffleEthDeposit(record); err == nil {
				t.Fatal("mapping accepted malformed deposit")
			}
		})
	}
}

func TestMapRoundRaffleNftWinnerRejectsMalformedData(t *testing.T) {
	t.Parallel()

	tests := map[string]func(*cgprimitives.CGRaffleNFTWinnerRec){
		"negative round":  func(r *cgprimitives.CGRaffleNFTWinnerRec) { r.RoundNum = -1 },
		"negative winner": func(r *cgprimitives.CGRaffleNFTWinnerRec) { r.WinnerIndex = -1 },
		"negative token":  func(r *cgprimitives.CGRaffleNFTWinnerRec) { r.TokenId = -1 },
		"wrong pool":      func(r *cgprimitives.CGRaffleNFTWinnerRec) { r.IsStaker = true },
		"event identity":  func(r *cgprimitives.CGRaffleNFTWinnerRec) { r.Tx.EvtLogId = 0 },
		"hash":            func(r *cgprimitives.CGRaffleNFTWinnerRec) { r.Tx.TxHash = "bad" },
		"timestamp":       func(r *cgprimitives.CGRaffleNFTWinnerRec) { r.Tx.DateTime = "bad" },
		"address":         func(r *cgprimitives.CGRaffleNFTWinnerRec) { r.WinnerAddr = "bad" },
		"empty amount":    func(r *cgprimitives.CGRaffleNFTWinnerRec) { r.CstAmount = "" },
		"fractional":      func(r *cgprimitives.CGRaffleNFTWinnerRec) { r.CstAmount = "1.5" },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validRaffleNftWinnerRecord(false)
			mutate(&record)
			if _, err := mapRoundRaffleNftWinner(record, false); err == nil {
				t.Fatal("mapping accepted malformed winner")
			}
		})
	}
}

func validRaffleEthDepositRecord() cgstore.RaffleEthDepositRecord {
	return cgstore.RaffleEthDepositRecord{
		Tx: cgprimitives.Transaction{
			EvtLogId: 7001,
			BlockNum: 123,
			TxHash:   "0xabcdef0000000000000000000000000000000000000000000000000000000001",
			DateTime: "2026-01-01T00:41:40-05:00",
		},
		RoundNum:     3,
		WinnerIndex:  2,
		WinnerAddr:   "0x2100000000000000000000000000000000000021",
		EthAmountWei: "050000000000000000",
	}
}

func validRaffleNftWinnerRecord(staker bool) cgprimitives.CGRaffleNFTWinnerRec {
	return cgprimitives.CGRaffleNFTWinnerRec{
		Tx: cgprimitives.Transaction{
			EvtLogId: 7002,
			BlockNum: 123,
			TxHash:   "0xabcdef0000000000000000000000000000000000000000000000000000000002",
			DateTime: "2026-01-01T00:41:40-05:00",
		},
		WinnerAddr:  "0x2100000000000000000000000000000000000021",
		RoundNum:    3,
		TokenId:     7,
		CstAmount:   "030000000000000000000",
		WinnerIndex: 2,
		IsRWalk:     staker,
		IsStaker:    staker,
	}
}
