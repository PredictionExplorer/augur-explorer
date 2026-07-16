package v2

import (
	"strings"
	"testing"
	"time"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func validGlobalTokenRecord() cgstore.GlobalTokenRecord {
	return cgstore.GlobalTokenRecord{
		MintTx:       validDonationTransaction(),
		TokenID:      5,
		MintRound:    2,
		Seed:         "seed05",
		TokenName:    "Genesis",
		WinnerAddr:   userCursorAlice,
		CurOwnerAid:  7,
		CurOwnerAddr: userCursorBob,
		MintSource:   cgstore.MintSourceMainPrize,
		Staked:       false,
	}
}

func TestMapGlobalToken(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		got, err := mapGlobalToken(validGlobalTokenRecord())
		if err != nil {
			t.Fatalf("mapGlobalToken: %v", err)
		}
		if got.NftTokenId != 5 || got.MintRound != 2 || got.Seed != "seed05" ||
			got.MintType != CosmicSignatureMintTypeMainPrize || got.Staked ||
			got.TokenName == nil || *got.TokenName != "Genesis" ||
			got.EventLogId != 100 ||
			!strings.EqualFold(got.WinnerAddress, userCursorAlice) ||
			!strings.EqualFold(got.CurrentOwnerAddress, userCursorBob) ||
			got.MintedAt.UTC().Format(time.RFC3339) != "2026-01-01T00:01:40Z" {
			t.Fatalf("global token = %+v", got)
		}
	})

	t.Run("unnamed token omits the name", func(t *testing.T) {
		t.Parallel()
		record := validGlobalTokenRecord()
		record.TokenName = ""
		got, err := mapGlobalToken(record)
		if err != nil || got.TokenName != nil {
			t.Fatalf("unnamed token = %+v, err=%v", got, err)
		}
	})

	invalid := map[string]func(*cgstore.GlobalTokenRecord){
		"negative token id": func(r *cgstore.GlobalTokenRecord) { r.TokenID = -1 },
		"negative round":    func(r *cgstore.GlobalTokenRecord) { r.MintRound = -1 },
		"missing seed":      func(r *cgstore.GlobalTokenRecord) { r.Seed = "" },
		"bad winner":        func(r *cgstore.GlobalTokenRecord) { r.WinnerAddr = "nope" },
		"bad current owner": func(r *cgstore.GlobalTokenRecord) { r.CurOwnerAddr = "nope" },
		"unknown source":    func(r *cgstore.GlobalTokenRecord) { r.MintSource = "airdrop" },
		"invalid event id":  func(r *cgstore.GlobalTokenRecord) { r.MintTx.EvtLogId = 0 },
		"invalid hash":      func(r *cgstore.GlobalTokenRecord) { r.MintTx.TxHash = "0xnope" },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validGlobalTokenRecord()
			corrupt(&record)
			if _, err := mapGlobalToken(record); err == nil {
				t.Fatal("corrupt record mapped")
			}
		})
	}
}

func validGlobalTokenDetailRecord() cgstore.GlobalTokenDetailRecord {
	record := cgstore.GlobalTokenDetailRecord{
		GlobalTokenRecord: validGlobalTokenRecord(),
	}
	record.Staked = true
	record.CurrentStake = &cgstore.GlobalTokenStake{
		StakeActionID: 9,
		StakerAid:     7,
		StakerAddr:    userCursorBob,
		StakedAt:      1_767_226_500,
		StakedAtText:  "2026-01-01T01:08:20+01:00",
	}
	return record
}

func TestMapGlobalTokenDetail(t *testing.T) {
	t.Parallel()

	t.Run("staked token carries its stake", func(t *testing.T) {
		t.Parallel()
		got, err := mapGlobalTokenDetail(validGlobalTokenDetailRecord())
		if err != nil {
			t.Fatalf("mapGlobalTokenDetail: %v", err)
		}
		if !got.Staked || got.CurrentStake == nil ||
			got.CurrentStake.StakeActionId != 9 ||
			!strings.EqualFold(got.CurrentStake.StakerAddress, userCursorBob) ||
			got.CurrentStake.StakedAt.UTC().Format(time.RFC3339) != "2026-01-01T00:08:20Z" {
			t.Fatalf("detail = %+v stake = %+v", got, got.CurrentStake)
		}
	})

	t.Run("unstaked token omits the stake", func(t *testing.T) {
		t.Parallel()
		record := validGlobalTokenDetailRecord()
		record.Staked = false
		record.CurrentStake = nil
		got, err := mapGlobalTokenDetail(record)
		if err != nil || got.Staked || got.CurrentStake != nil {
			t.Fatalf("unstaked detail = %+v, err=%v", got, err)
		}
	})

	invalid := map[string]func(*cgstore.GlobalTokenDetailRecord){
		"staked without stake action": func(r *cgstore.GlobalTokenDetailRecord) { r.CurrentStake = nil },
		"stake without staked flag": func(r *cgstore.GlobalTokenDetailRecord) {
			r.Staked = false
		},
		"negative stake action": func(r *cgstore.GlobalTokenDetailRecord) {
			r.CurrentStake.StakeActionID = -1
		},
		"bad staker address": func(r *cgstore.GlobalTokenDetailRecord) {
			r.CurrentStake.StakerAddr = "nope"
		},
		"bad stake timestamp": func(r *cgstore.GlobalTokenDetailRecord) {
			r.CurrentStake.StakedAtText = "yesterday"
		},
		"corrupt base record": func(r *cgstore.GlobalTokenDetailRecord) {
			r.Seed = ""
		},
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validGlobalTokenDetailRecord()
			corrupt(&record)
			if _, err := mapGlobalTokenDetail(record); err == nil {
				t.Fatal("corrupt record mapped")
			}
		})
	}
}

func validTokenNameChangeRecord() cgstore.TokenNameChangeRecord {
	return cgstore.TokenNameChangeRecord{
		Tx:           validDonationTransaction(),
		TokenID:      5,
		NewName:      "Aurora",
		ChangedByAid: 3,
		ChangedBy:    userCursorAlice,
	}
}

func TestMapTokenNameChange(t *testing.T) {
	t.Parallel()

	got, err := mapTokenNameChange(validTokenNameChangeRecord())
	if err != nil || got.NftTokenId != 5 || got.TokenName != "Aurora" ||
		!strings.EqualFold(got.ChangedByAddress, userCursorAlice) ||
		got.EventLogId != 100 {
		t.Fatalf("rename = %+v, err=%v", got, err)
	}

	cleared := validTokenNameChangeRecord()
	cleared.NewName = ""
	got, err = mapTokenNameChange(cleared)
	if err != nil || got.TokenName != "" {
		t.Fatalf("cleared rename = %+v, err=%v", got, err)
	}

	invalid := map[string]func(*cgstore.TokenNameChangeRecord){
		"negative token": func(r *cgstore.TokenNameChangeRecord) { r.TokenID = -1 },
		"bad author":     func(r *cgstore.TokenNameChangeRecord) { r.ChangedBy = "nope" },
		"bad tx hash":    func(r *cgstore.TokenNameChangeRecord) { r.Tx.TxHash = "0xnope" },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validTokenNameChangeRecord()
			corrupt(&record)
			if _, err := mapTokenNameChange(record); err == nil {
				t.Fatal("corrupt record mapped")
			}
		})
	}
}

func validTokenTransferRecord() cgstore.TokenTransferRecord {
	return cgstore.TokenTransferRecord{
		Tx:           validDonationTransaction(),
		TokenID:      5,
		FromAid:      1,
		FromAddr:     userCursorAlice,
		ToAid:        2,
		ToAddr:       userCursorBob,
		TransferType: 0,
	}
}

func TestMapTokenTransfer(t *testing.T) {
	t.Parallel()

	want := map[int64]TokenTransferType{0: TokenTransferTypeTransfer, 1: TokenTransferTypeMint, 2: TokenTransferTypeBurn}
	for otype, transferType := range want {
		record := validTokenTransferRecord()
		record.TransferType = otype
		got, err := mapTokenTransfer(record)
		if err != nil || got.TransferType != transferType || got.NftTokenId != 5 ||
			!strings.EqualFold(got.FromAddress, userCursorAlice) ||
			!strings.EqualFold(got.ToAddress, userCursorBob) {
			t.Errorf("otype %d = %+v, err=%v", otype, got, err)
		}
	}

	invalid := map[string]func(*cgstore.TokenTransferRecord){
		"negative token": func(r *cgstore.TokenTransferRecord) { r.TokenID = -1 },
		"bad from":       func(r *cgstore.TokenTransferRecord) { r.FromAddr = "nope" },
		"bad to":         func(r *cgstore.TokenTransferRecord) { r.ToAddr = "nope" },
		"unknown otype":  func(r *cgstore.TokenTransferRecord) { r.TransferType = 9 },
		"bad tx hash":    func(r *cgstore.TokenTransferRecord) { r.Tx.TxHash = "0xnope" },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validTokenTransferRecord()
			corrupt(&record)
			if _, err := mapTokenTransfer(record); err == nil {
				t.Fatal("corrupt record mapped")
			}
		})
	}
}

func TestMapHolders(t *testing.T) {
	t.Parallel()

	holder, err := mapCosmicSignatureHolder(cgstore.CosmicSignatureHolderRecord{
		OwnerAid: 1, Address: userCursorAlice, TokenCount: 4,
	})
	if err != nil || holder.TokenCount != 4 ||
		!strings.EqualFold(holder.OwnerAddress, userCursorAlice) {
		t.Fatalf("cs holder = %+v, err=%v", holder, err)
	}
	if _, err := mapCosmicSignatureHolder(cgstore.CosmicSignatureHolderRecord{
		OwnerAid: 1, Address: userCursorAlice, TokenCount: 0,
	}); err == nil {
		t.Fatal("holder without tokens mapped")
	}
	if _, err := mapCosmicSignatureHolder(cgstore.CosmicSignatureHolderRecord{
		OwnerAid: 1, Address: "nope", TokenCount: 1,
	}); err == nil {
		t.Fatal("holder with bad address mapped")
	}

	balance, err := mapCosmicTokenHolder(cgstore.CosmicTokenHolderRecord{
		OwnerAid: 1, Address: userCursorAlice, BalanceWei: "1000000000000000000000",
	})
	if err != nil || balance.BalanceWei != "1000000000000000000000" ||
		!strings.EqualFold(balance.OwnerAddress, userCursorAlice) {
		t.Fatalf("ct holder = %+v, err=%v", balance, err)
	}
	for name, record := range map[string]cgstore.CosmicTokenHolderRecord{
		"zero balance":     {OwnerAid: 1, Address: userCursorAlice, BalanceWei: "0"},
		"negative balance": {OwnerAid: 1, Address: userCursorAlice, BalanceWei: "-1"},
		"empty balance":    {OwnerAid: 1, Address: userCursorAlice, BalanceWei: ""},
		"bad address":      {OwnerAid: 1, Address: "nope", BalanceWei: "1"},
	} {
		if _, err := mapCosmicTokenHolder(record); err == nil {
			t.Errorf("%s mapped", name)
		}
	}
}

func TestSupplySharePercentage(t *testing.T) {
	t.Parallel()
	cases := []struct {
		balance string
		supply  string
		want    string
	}{
		{"1", "3", "33.33"},
		{"1", "1", "100"},
		{"1", "10000", "0.01"},
		{"1", "1000000", "0"},
		{"2", "3", "66.67"},
		{"25", "100", "25"},
	}
	for _, test := range cases {
		got, err := supplySharePercentage(test.balance, test.supply)
		if err != nil || got != test.want {
			t.Errorf("share(%s/%s) = %q, %v; want %q",
				test.balance, test.supply, got, err, test.want)
		}
	}
	for name, pair := range map[string][2]string{
		"balance above supply": {"2", "1"},
		"zero supply":          {"1", "0"},
		"bad balance":          {"x", "1"},
		"bad supply":           {"1", "x"},
	} {
		if _, err := supplySharePercentage(pair[0], pair[1]); err == nil {
			t.Errorf("%s accepted", name)
		}
	}
}

func validCosmicTokenStatisticsRecord() cgstore.CosmicTokenStatisticsRecord {
	return cgstore.CosmicTokenStatisticsRecord{
		TotalSupplyWei:             "1000",
		HolderCount:                3,
		BiddingRewardsWei:          "700",
		MainPrizesWei:              "100",
		RafflePrizesWei:            "80",
		ChronoWarriorPrizesWei:     "60",
		EnduranceChampionPrizesWei: "40",
		LastCstBidderPrizesWei:     "15",
		MarketingRewardsWei:        "5",
		TotalEarnedWei:             "1000",
		ConsumedInBidsWei:          "300",
		NetWei:                     "700",
		TransferCount:              9,
		MintCount:                  6,
		BurnCount:                  2,
		TopHolders: []cgstore.CosmicTokenTopHolderRecord{
			{OwnerAid: 1, Address: userCursorAlice, BalanceWei: "600"},
			{OwnerAid: 2, Address: userCursorBob, BalanceWei: "300"},
			{OwnerAid: 5, Address: "0x2300000000000000000000000000000000000023", BalanceWei: "100"},
		},
	}
}

func TestMapCosmicTokenStatistics(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		got, err := mapCosmicTokenStatistics(validCosmicTokenStatisticsRecord())
		if err != nil {
			t.Fatalf("mapCosmicTokenStatistics: %v", err)
		}
		if got.TotalSupplyWei != "1000" || got.HolderCount != 3 ||
			got.Earned.TotalWei != "1000" || got.ConsumedInBidsWei != "300" ||
			got.NetWei != "700" || got.Transfers.TotalCount != 9 ||
			len(got.TopHolders) != 3 {
			t.Fatalf("statistics = %+v", got)
		}
		if got.TopHolders[0].ShareOfSupply != "60" ||
			got.TopHolders[1].ShareOfSupply != "30" ||
			got.TopHolders[2].ShareOfSupply != "10" {
			t.Fatalf("shares = %+v", got.TopHolders)
		}
	})

	t.Run("negative net flow maps", func(t *testing.T) {
		t.Parallel()
		record := validCosmicTokenStatisticsRecord()
		record.ConsumedInBidsWei = "1500"
		record.NetWei = "-500"
		got, err := mapCosmicTokenStatistics(record)
		if err != nil || got.NetWei != "-500" {
			t.Fatalf("negative net = %+v, err=%v", got, err)
		}
	})

	invalid := map[string]func(*cgstore.CosmicTokenStatisticsRecord){
		"sources do not add up": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.BiddingRewardsWei = "1"
		},
		"net diverges": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.NetWei = "1"
		},
		"malformed source": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.BiddingRewardsWei = "x"
		},
		"malformed consumed": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.ConsumedInBidsWei = "x"
		},
		"malformed net": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.NetWei = "x"
		},
		"malformed top-holder balance": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.TopHolders[0].BalanceWei = "x"
		},
		"negative holder count": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.HolderCount = -1
		},
		"inconsistent transfer counters": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.MintCount = 8
		},
		"bad supply": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.TotalSupplyWei = "x"
		},
		"too many top holders": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.TopHolders = make([]cgstore.CosmicTokenTopHolderRecord, 11)
		},
		"holder count below list": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.HolderCount = 2
		},
		"unordered top holders": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.TopHolders[0], r.TopHolders[1] = r.TopHolders[1], r.TopHolders[0]
		},
		"tied holders with unordered ids": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.TopHolders[1].BalanceWei = "600"
			r.TopHolders[1].OwnerAid = 0
		},
		"top holder without balance": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.TopHolders[2].BalanceWei = "0"
		},
		"top holder above supply": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.TopHolders[0].BalanceWei = "1001"
		},
		"bad top holder address": func(r *cgstore.CosmicTokenStatisticsRecord) {
			r.TopHolders[0].Address = "nope"
		},
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validCosmicTokenStatisticsRecord()
			corrupt(&record)
			if _, err := mapCosmicTokenStatistics(record); err == nil {
				t.Fatal("corrupt record mapped")
			}
		})
	}
}

func validSupplyChangeRecord() cgstore.SupplyChangeRecord {
	return cgstore.SupplyChangeRecord{
		Tx:             validDonationTransaction(),
		BidderAid:      1,
		BidderAddr:     userCursorAlice,
		BidType:        2,
		MintedWei:      "100",
		BurnedWei:      "30",
		NetWei:         "70",
		TotalSupplyWei: "570",
	}
}

func TestMapSupplyChange(t *testing.T) {
	t.Parallel()

	got, err := mapSupplyChange(validSupplyChangeRecord())
	if err != nil || got.MintedWei != "100" || got.BurnedWei != "30" ||
		got.NetWei != "70" || got.TotalSupplyWei != "570" ||
		got.BidType != Cst || !strings.EqualFold(got.BidderAddress, userCursorAlice) {
		t.Fatalf("supply change = %+v, err=%v", got, err)
	}

	negative := validSupplyChangeRecord()
	negative.MintedWei = "0"
	negative.BurnedWei = "30"
	negative.NetWei = "-30"
	negative.TotalSupplyWei = "470"
	if got, err := mapSupplyChange(negative); err != nil || got.NetWei != "-30" {
		t.Fatalf("negative net = %+v, err=%v", got, err)
	}

	invalid := map[string]func(*cgstore.SupplyChangeRecord){
		"net diverges":            func(r *cgstore.SupplyChangeRecord) { r.NetWei = "71" },
		"malformed net":           func(r *cgstore.SupplyChangeRecord) { r.NetWei = "x" },
		"negative running supply": func(r *cgstore.SupplyChangeRecord) { r.TotalSupplyWei = "-1" },
		"bad minted":              func(r *cgstore.SupplyChangeRecord) { r.MintedWei = "x" },
		"bad burned":              func(r *cgstore.SupplyChangeRecord) { r.BurnedWei = "-2" },
		"bad bidder":              func(r *cgstore.SupplyChangeRecord) { r.BidderAddr = "nope" },
		"bad tx":                  func(r *cgstore.SupplyChangeRecord) { r.Tx.TxHash = "0xnope" },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validSupplyChangeRecord()
			corrupt(&record)
			if _, err := mapSupplyChange(record); err == nil {
				t.Fatal("corrupt record mapped")
			}
		})
	}
}

func TestMapDailySupply(t *testing.T) {
	t.Parallel()

	got, err := mapDailySupply(cgstore.DailySupplyRecord{
		Date:           "2026-01-05",
		BidCount:       4,
		MintedWei:      "400",
		BurnedWei:      "100",
		NetWei:         "300",
		TotalSupplyWei: "900",
	})
	if err != nil || got.BidCount != 4 ||
		got.Date.Format("2006-01-02") != "2026-01-05" ||
		got.NetWei != "300" || got.TotalSupplyWei != "900" {
		t.Fatalf("daily supply = %+v, err=%v", got, err)
	}

	invalid := map[string]cgstore.DailySupplyRecord{
		"no bids": {
			Date: "2026-01-05", BidCount: 0,
			MintedWei: "0", BurnedWei: "0", NetWei: "0", TotalSupplyWei: "0",
		},
		"bad date": {
			Date: "Jan 5", BidCount: 1,
			MintedWei: "1", BurnedWei: "0", NetWei: "1", TotalSupplyWei: "1",
		},
		"net diverges": {
			Date: "2026-01-05", BidCount: 1,
			MintedWei: "1", BurnedWei: "0", NetWei: "2", TotalSupplyWei: "1",
		},
		"negative running supply": {
			Date: "2026-01-05", BidCount: 1,
			MintedWei: "1", BurnedWei: "0", NetWei: "1", TotalSupplyWei: "-1",
		},
		"malformed minted": {
			Date: "2026-01-05", BidCount: 1,
			MintedWei: "x", BurnedWei: "0", NetWei: "1", TotalSupplyWei: "1",
		},
		"malformed burned": {
			Date: "2026-01-05", BidCount: 1,
			MintedWei: "1", BurnedWei: "x", NetWei: "1", TotalSupplyWei: "1",
		},
		"malformed net": {
			Date: "2026-01-05", BidCount: 1,
			MintedWei: "1", BurnedWei: "0", NetWei: "x", TotalSupplyWei: "1",
		},
	}
	for name, record := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := mapDailySupply(record); err == nil {
				t.Fatal("corrupt record mapped")
			}
		})
	}
}

func TestMapMarketingReward(t *testing.T) {
	t.Parallel()

	got, err := mapMarketingReward(cgstore.MarketingRewardRecord{
		Tx:           validDonationTransaction(),
		MarketerAid:  1,
		MarketerAddr: userCursorAlice,
		AmountWei:    "5000000000000000000",
	})
	if err != nil || got.AmountWei != "5000000000000000000" ||
		!strings.EqualFold(got.MarketerAddress, userCursorAlice) || got.EventLogId != 100 {
		t.Fatalf("marketing reward = %+v, err=%v", got, err)
	}

	for name, corrupt := range map[string]func(*cgstore.MarketingRewardRecord){
		"bad marketer": func(r *cgstore.MarketingRewardRecord) { r.MarketerAddr = "nope" },
		"bad amount":   func(r *cgstore.MarketingRewardRecord) { r.AmountWei = "-1" },
		"bad tx":       func(r *cgstore.MarketingRewardRecord) { r.Tx.TxHash = "0xnope" },
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := cgstore.MarketingRewardRecord{
				Tx:           validDonationTransaction(),
				MarketerAid:  1,
				MarketerAddr: userCursorAlice,
				AmountWei:    "1",
			}
			corrupt(&record)
			if _, err := mapMarketingReward(record); err == nil {
				t.Fatal("corrupt record mapped")
			}
		})
	}
}
