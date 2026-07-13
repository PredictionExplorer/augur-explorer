package v2

import (
	"encoding/json"
	"strings"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

func TestMapRoundPrizeTaxonomy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		recordType int64
		wantType   RoundPrizeType
		asset      roundPrizeAsset
	}{
		{0, MainPrizeEth, roundPrizeAssetEth},
		{1, MainPrizeCst, roundPrizeAssetCst},
		{2, MainPrizeNft, roundPrizeAssetNft},
		{3, LastCstBidderNft, roundPrizeAssetNft},
		{4, LastCstBidderCst, roundPrizeAssetCst},
		{5, EnduranceChampionNft, roundPrizeAssetNft},
		{6, EnduranceChampionCst, roundPrizeAssetCst},
		{7, ChronoWarriorEth, roundPrizeAssetEth},
		{8, ChronoWarriorCst, roundPrizeAssetCst},
		{9, ChronoWarriorNft, roundPrizeAssetNft},
		{10, BidderRaffleEth, roundPrizeAssetEth},
		{11, BidderRaffleCst, roundPrizeAssetCst},
		{12, BidderRaffleNft, roundPrizeAssetNft},
		{13, RandomWalkStakerRaffleCst, roundPrizeAssetCst},
		{14, RandomWalkStakerRaffleNft, roundPrizeAssetNft},
		{15, CosmicSignatureStakingEth, roundPrizeAssetEth},
	}
	for _, tc := range tests {
		t.Run(string(tc.wantType), func(t *testing.T) {
			t.Parallel()
			record := validRoundPrizeRecord(tc.recordType)
			got, err := mapRoundPrize(record)
			if err != nil {
				t.Fatalf("mapRoundPrize(type %d): %v", tc.recordType, err)
			}
			if got.Type != tc.wantType ||
				got.Round != 3 ||
				got.WinnerIndex != 2 ||
				got.EventLogId != 7001 ||
				got.BlockNumber != 123 {
				t.Fatalf("mapped identity = %+v", got)
			}
			if got.OccurredAt.String() != "2026-01-01 05:41:40 +0000 UTC" {
				t.Errorf("OccurredAt = %s", got.OccurredAt)
			}
			if got.TransactionHash != "0xabcdef0000000000000000000000000000000000000000000000000000000001" {
				t.Errorf("TransactionHash = %q", got.TransactionHash)
			}
			if tc.recordType == 15 {
				if got.WinnerAddress != nil {
					t.Fatalf("staking winner = %q, want omitted", *got.WinnerAddress)
				}
			} else if got.WinnerAddress == nil ||
				*got.WinnerAddress != "0x2100000000000000000000000000000000000021" {
				t.Fatalf("WinnerAddress = %v", got.WinnerAddress)
			}

			switch tc.asset {
			case roundPrizeAssetEth:
				if got.EthAmountWei == nil || *got.EthAmountWei != "123" ||
					got.CstAmountWei != nil || got.NftTokenId != nil {
					t.Fatalf("ETH asset fields = %+v", got)
				}
			case roundPrizeAssetCst:
				if got.CstAmountWei == nil || *got.CstAmountWei != "123" ||
					got.EthAmountWei != nil || got.NftTokenId != nil {
					t.Fatalf("CST asset fields = %+v", got)
				}
			case roundPrizeAssetNft:
				if got.NftTokenId == nil || *got.NftTokenId != 7 ||
					got.EthAmountWei != nil || got.CstAmountWei != nil {
					t.Fatalf("NFT asset fields = %+v", got)
				}
			}
		})
	}
}

func TestMapRoundPrizeOmitsLegacyFields(t *testing.T) {
	t.Parallel()

	got, err := mapRoundPrize(validRoundPrizeRecord(0))
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"AmountEth", "amountEth", "WinnerAid", "winnerAid", "claimed",
		"isTimeoutClaim", "tokenAddress", "tokenUri",
	} {
		if strings.Contains(string(encoded), forbidden) {
			t.Errorf("round prize JSON leaked %q: %s", forbidden, encoded)
		}
	}
}

func TestMapRoundPrizeRejectsMalformedData(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		recordType int64
		mutate     func(*cgmodel.CGPrizeHistory)
	}{
		"negative round": {
			recordType: 0,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.RoundNum = -1 },
		},
		"negative winner index": {
			recordType: 0,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.WinnerIndex = -1 },
		},
		"transaction identity": {
			recordType: 0,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.Tx.EvtLogId = 0 },
		},
		"transaction hash": {
			recordType: 0,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.Tx.TxHash = "bad" },
		},
		"transaction timestamp": {
			recordType: 0,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.Tx.DateTime = "bad" },
		},
		"unknown prize type": {
			recordType: 16,
		},
		"invalid winner": {
			recordType: 0,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.WinnerAddr = "bad" },
		},
		"staking has winner": {
			recordType: 15,
			mutate: func(r *cgmodel.CGPrizeHistory) {
				r.WinnerAddr = "0x2100000000000000000000000000000000000021"
			},
		},
		"ETH amount": {
			recordType: 0,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.Amount = "-1" },
		},
		"ETH token id": {
			recordType: 0,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.TokenId = 1 },
		},
		"CST amount": {
			recordType: 1,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.Amount = "1.2" },
		},
		"CST token id": {
			recordType: 1,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.TokenId = 1 },
		},
		"NFT token id": {
			recordType: 2,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.TokenId = -1 },
		},
		"NFT fungible amount": {
			recordType: 2,
			mutate:     func(r *cgmodel.CGPrizeHistory) { r.Amount = "1" },
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validRoundPrizeRecord(tc.recordType)
			if tc.mutate != nil {
				tc.mutate(&record)
			}
			if _, err := mapRoundPrize(record); err == nil {
				t.Fatal("mapRoundPrize accepted malformed data")
			}
		})
	}
}

func validRoundPrizeRecord(recordType int64) cgmodel.CGPrizeHistory {
	record := cgmodel.CGPrizeHistory{
		Tx: cgmodel.Transaction{
			EvtLogId: 7001,
			BlockNum: 123,
			TxHash:   "0xABCDEF0000000000000000000000000000000000000000000000000000000001",
			DateTime: "2026-01-01T00:41:40-05:00",
		},
		RecordType:  recordType,
		RoundNum:    3,
		Amount:      "000123",
		WinnerIndex: 2,
		TokenId:     -1,
		WinnerAddr:  "0x2100000000000000000000000000000000000021",
	}
	switch recordType {
	case 2, 3, 5, 9, 12, 14:
		record.Amount = "000"
		record.TokenId = 7
	case 15:
		record.WinnerAddr = "(All CS NFT Stakers)"
	}
	return record
}
