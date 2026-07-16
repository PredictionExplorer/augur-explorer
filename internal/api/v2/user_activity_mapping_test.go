package v2

import (
	"strings"
	"testing"
	"time"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func validUserOwnedTokenRecord() cgstore.UserOwnedTokenRecord {
	return cgstore.UserOwnedTokenRecord{
		MintTx:     validDonationTransaction(),
		OwnerAid:   1,
		TokenID:    5,
		MintRound:  2,
		Seed:       "seed05",
		TokenName:  "Genesis",
		WinnerAddr: userCursorAlice,
		MintSource: cgstore.MintSourceMainPrize,
		Staked:     true,
	}
}

func TestMapUserOwnedToken(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		got, err := mapUserOwnedToken(validUserOwnedTokenRecord())
		if err != nil {
			t.Fatalf("mapUserOwnedToken: %v", err)
		}
		if got.NftTokenId != 5 || got.MintRound != 2 || got.Seed != "seed05" ||
			got.MintType != CosmicSignatureMintTypeMainPrize || !got.Staked ||
			got.TokenName == nil || *got.TokenName != "Genesis" ||
			got.EventLogId != 100 ||
			!strings.EqualFold(got.WinnerAddress, userCursorAlice) ||
			got.MintedAt.UTC().Format(time.RFC3339) != "2026-01-01T00:01:40Z" {
			t.Fatalf("owned token = %+v", got)
		}
	})

	t.Run("unnamed token omits the name", func(t *testing.T) {
		t.Parallel()
		record := validUserOwnedTokenRecord()
		record.TokenName = ""
		got, err := mapUserOwnedToken(record)
		if err != nil || got.TokenName != nil {
			t.Fatalf("unnamed token = %+v, err=%v", got, err)
		}
	})

	t.Run("every mint source maps", func(t *testing.T) {
		t.Parallel()
		want := map[cgstore.CosmicSignatureMintSource]CosmicSignatureMintType{
			cgstore.MintSourceMainPrize:          CosmicSignatureMintTypeMainPrize,
			cgstore.MintSourceBidderRaffle:       CosmicSignatureMintTypeBidderRaffle,
			cgstore.MintSourceRandomWalkStaker:   CosmicSignatureMintTypeRandomWalkStakerRaffle,
			cgstore.MintSourceCosmicSigStaker:    CosmicSignatureMintTypeCosmicSignatureStakerRaffle,
			cgstore.MintSourceEnduranceChampion:  CosmicSignatureMintTypeEnduranceChampion,
			cgstore.MintSourceLastCstBidder:      CosmicSignatureMintTypeLastCstBidder,
			cgstore.MintSourceChronoWarriorPrize: CosmicSignatureMintTypeChronoWarrior,
		}
		for source, mintType := range want {
			record := validUserOwnedTokenRecord()
			record.MintSource = source
			got, err := mapUserOwnedToken(record)
			if err != nil || got.MintType != mintType || !got.MintType.Valid() {
				t.Errorf("source %s = %+v, err=%v", source, got.MintType, err)
			}
		}
		if CosmicSignatureMintType("airdrop").Valid() {
			t.Error("unknown mint type reported valid")
		}
	})

	invalid := map[string]func(*cgstore.UserOwnedTokenRecord){
		"negative token id":  func(r *cgstore.UserOwnedTokenRecord) { r.TokenID = -1 },
		"negative round":     func(r *cgstore.UserOwnedTokenRecord) { r.MintRound = -1 },
		"missing seed":       func(r *cgstore.UserOwnedTokenRecord) { r.Seed = "" },
		"bad winner address": func(r *cgstore.UserOwnedTokenRecord) { r.WinnerAddr = "nope" },
		"unknown source":     func(r *cgstore.UserOwnedTokenRecord) { r.MintSource = "airdrop" },
		"invalid event id":   func(r *cgstore.UserOwnedTokenRecord) { r.MintTx.EvtLogId = 0 },
		"invalid hash":       func(r *cgstore.UserOwnedTokenRecord) { r.MintTx.TxHash = "0xnope" },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validUserOwnedTokenRecord()
			corrupt(&record)
			if _, err := mapUserOwnedToken(record); err == nil {
				t.Fatalf("mapUserOwnedToken accepted %s", name)
			}
		})
	}
}

func validCsTransferRecord() cgstore.UserCosmicSignatureTransferRecord {
	return cgstore.UserCosmicSignatureTransferRecord{
		Tx:           validDonationTransaction(),
		TokenID:      5,
		FromAid:      1,
		FromAddr:     userCursorAlice,
		ToAid:        2,
		ToAddr:       userCursorBob,
		TransferType: 0,
		Direction:    cgstore.UserTransferOut,
	}
}

func TestMapUserCosmicSignatureTransfer(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		got, err := mapUserCosmicSignatureTransfer(validCsTransferRecord())
		if err != nil {
			t.Fatalf("mapUserCosmicSignatureTransfer: %v", err)
		}
		if got.NftTokenId != 5 || got.TransferType != TokenTransferTypeTransfer || got.Direction != Out ||
			!strings.EqualFold(got.FromAddress, userCursorAlice) ||
			!strings.EqualFold(got.ToAddress, userCursorBob) ||
			got.EventLogId != 100 {
			t.Fatalf("transfer = %+v", got)
		}
	})

	t.Run("types and directions map", func(t *testing.T) {
		t.Parallel()
		types := map[int64]TokenTransferType{0: TokenTransferTypeTransfer, 1: TokenTransferTypeMint, 2: TokenTransferTypeBurn}
		directions := map[cgstore.UserTransferDirection]TokenTransferDirection{
			cgstore.UserTransferIn:   In,
			cgstore.UserTransferOut:  Out,
			cgstore.UserTransferSelf: Self,
		}
		for otype, wantType := range types {
			for direction, wantDirection := range directions {
				record := validCsTransferRecord()
				record.TransferType = otype
				record.Direction = direction
				got, err := mapUserCosmicSignatureTransfer(record)
				if err != nil || got.TransferType != wantType || got.Direction != wantDirection ||
					!got.TransferType.Valid() || !got.Direction.Valid() {
					t.Errorf("otype=%d direction=%s => %+v, err=%v", otype, direction, got, err)
				}
			}
		}
		if TokenTransferType("teleport").Valid() || TokenTransferDirection("sideways").Valid() {
			t.Error("unknown enum members reported valid")
		}
	})

	invalid := map[string]func(*cgstore.UserCosmicSignatureTransferRecord){
		"negative token id": func(r *cgstore.UserCosmicSignatureTransferRecord) { r.TokenID = -1 },
		"bad from address":  func(r *cgstore.UserCosmicSignatureTransferRecord) { r.FromAddr = "" },
		"bad to address":    func(r *cgstore.UserCosmicSignatureTransferRecord) { r.ToAddr = "nope" },
		"unknown otype":     func(r *cgstore.UserCosmicSignatureTransferRecord) { r.TransferType = 7 },
		"unknown direction": func(r *cgstore.UserCosmicSignatureTransferRecord) { r.Direction = "sideways" },
		"invalid event id":  func(r *cgstore.UserCosmicSignatureTransferRecord) { r.Tx.EvtLogId = 0 },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validCsTransferRecord()
			corrupt(&record)
			if _, err := mapUserCosmicSignatureTransfer(record); err == nil {
				t.Fatalf("mapUserCosmicSignatureTransfer accepted %s", name)
			}
		})
	}
}

func validCtTransferRecord() cgstore.UserCosmicTokenTransferRecord {
	return cgstore.UserCosmicTokenTransferRecord{
		Tx:           validDonationTransaction(),
		AmountWei:    "10000000000000000000",
		FromAid:      1,
		FromAddr:     userCursorAlice,
		ToAid:        2,
		ToAddr:       userCursorBob,
		TransferType: 0,
		Direction:    cgstore.UserTransferOut,
	}
}

func TestMapUserCosmicTokenTransfer(t *testing.T) {
	t.Parallel()

	got, err := mapUserCosmicTokenTransfer(validCtTransferRecord())
	if err != nil {
		t.Fatalf("mapUserCosmicTokenTransfer: %v", err)
	}
	if got.AmountWei != "10000000000000000000" || got.TransferType != TokenTransferTypeTransfer ||
		got.Direction != Out || got.EventLogId != 100 {
		t.Fatalf("transfer = %+v", got)
	}

	invalid := map[string]func(*cgstore.UserCosmicTokenTransferRecord){
		"empty amount":      func(r *cgstore.UserCosmicTokenTransferRecord) { r.AmountWei = "" },
		"negative amount":   func(r *cgstore.UserCosmicTokenTransferRecord) { r.AmountWei = "-1" },
		"malformed amount":  func(r *cgstore.UserCosmicTokenTransferRecord) { r.AmountWei = "1.5e18" },
		"bad from address":  func(r *cgstore.UserCosmicTokenTransferRecord) { r.FromAddr = "nope" },
		"bad to address":    func(r *cgstore.UserCosmicTokenTransferRecord) { r.ToAddr = "" },
		"unknown otype":     func(r *cgstore.UserCosmicTokenTransferRecord) { r.TransferType = -1 },
		"unknown direction": func(r *cgstore.UserCosmicTokenTransferRecord) { r.Direction = "" },
		"invalid hash":      func(r *cgstore.UserCosmicTokenTransferRecord) { r.Tx.TxHash = "0xnope" },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validCtTransferRecord()
			corrupt(&record)
			if _, err := mapUserCosmicTokenTransfer(record); err == nil {
				t.Fatalf("mapUserCosmicTokenTransfer accepted %s", name)
			}
		})
	}
}

func TestMapUserMarketingReward(t *testing.T) {
	t.Parallel()

	record := cgstore.UserMarketingRewardRecord{
		Tx:          validDonationTransaction(),
		MarketerAid: 1,
		AmountWei:   "50000000000000000000",
	}
	got, err := mapUserMarketingReward(record)
	if err != nil || got.AmountWei != "50000000000000000000" || got.EventLogId != 100 {
		t.Fatalf("marketing reward = %+v, err=%v", got, err)
	}

	bad := record
	bad.AmountWei = "-5"
	if _, err := mapUserMarketingReward(bad); err == nil {
		t.Fatal("mapUserMarketingReward accepted a negative amount")
	}
	badTx := record
	badTx.Tx.DateTime = "yesterday"
	if _, err := mapUserMarketingReward(badTx); err == nil {
		t.Fatal("mapUserMarketingReward accepted an unparseable datetime")
	}
}

func validUserCosmicTokenSummaryRecord() cgstore.UserCosmicTokenSummaryRecord {
	return cgstore.UserCosmicTokenSummaryRecord{
		BalanceWei:                 "290000000000000000000",
		BiddingRewardsWei:          "500000000000000000000",
		MainPrizesWei:              "100000000000000000000",
		RafflePrizesWei:            "0",
		ChronoWarriorPrizesWei:     "35000000000000000000",
		EnduranceChampionPrizesWei: "0",
		LastCstBidderPrizesWei:     "0",
		MarketingRewardsWei:        "0",
		TotalEarnedWei:             "635000000000000000000",
		ConsumedInBidsWei:          "0",
		NetWei:                     "635000000000000000000",
		TransferCount:              4,
		MintCount:                  3,
		BurnCount:                  0,
	}
}

func TestMapUserCosmicTokenSummary(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		got, err := mapUserCosmicTokenSummary(userCursorAlice, validUserCosmicTokenSummaryRecord())
		if err != nil {
			t.Fatalf("mapUserCosmicTokenSummary: %v", err)
		}
		if got.Address != userCursorAlice || got.BalanceWei != "290000000000000000000" ||
			got.Earned.TotalWei != "635000000000000000000" ||
			got.NetWei != "635000000000000000000" ||
			got.Transfers.TotalCount != 4 || got.Transfers.MintCount != 3 {
			t.Fatalf("summary = %+v", got)
		}
	})

	t.Run("negative net flow is legal", func(t *testing.T) {
		t.Parallel()
		record := validUserCosmicTokenSummaryRecord()
		record.ConsumedInBidsWei = "700000000000000000000"
		record.NetWei = "-65000000000000000000"
		got, err := mapUserCosmicTokenSummary(userCursorAlice, record)
		if err != nil || got.NetWei != "-65000000000000000000" {
			t.Fatalf("negative net = %+v, err=%v", got, err)
		}
	})

	invalid := map[string]func(*cgstore.UserCosmicTokenSummaryRecord){
		"negative balance":       func(r *cgstore.UserCosmicTokenSummaryRecord) { r.BalanceWei = "-1" },
		"empty earning":          func(r *cgstore.UserCosmicTokenSummaryRecord) { r.MainPrizesWei = "" },
		"malformed earning":      func(r *cgstore.UserCosmicTokenSummaryRecord) { r.RafflePrizesWei = "1e18" },
		"sources exceed total":   func(r *cgstore.UserCosmicTokenSummaryRecord) { r.TotalEarnedWei = "1" },
		"malformed consumed":     func(r *cgstore.UserCosmicTokenSummaryRecord) { r.ConsumedInBidsWei = "wat" },
		"malformed net":          func(r *cgstore.UserCosmicTokenSummaryRecord) { r.NetWei = "wat" },
		"net diverges":           func(r *cgstore.UserCosmicTokenSummaryRecord) { r.NetWei = "5" },
		"negative transfers":     func(r *cgstore.UserCosmicTokenSummaryRecord) { r.TransferCount = -1 },
		"counts exceed the pool": func(r *cgstore.UserCosmicTokenSummaryRecord) { r.MintCount = 9 },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validUserCosmicTokenSummaryRecord()
			corrupt(&record)
			if _, err := mapUserCosmicTokenSummary(userCursorAlice, record); err == nil {
				t.Fatalf("mapUserCosmicTokenSummary accepted %s", name)
			}
		})
	}
}

func TestZeroUserCosmicTokenSummaryMatchesStoreZeroShape(t *testing.T) {
	t.Parallel()
	mapped, err := mapUserCosmicTokenSummary(userCursorAlice, zeroUserCosmicTokenSummaryRecord())
	if err != nil {
		t.Fatalf("mapping the store zero shape: %v", err)
	}
	if mapped != zeroUserCosmicTokenSummary(userCursorAlice) {
		t.Fatalf("zero shapes diverge: mapped=%+v static=%+v",
			mapped, zeroUserCosmicTokenSummary(userCursorAlice))
	}
}

func TestMapUserPendingWinnings(t *testing.T) {
	t.Parallel()

	record := cgstore.UserPendingWinningsRecord{
		RaffleEthWei:           "60000000000000000",
		ChronoWarriorEthWei:    "80000000000000000",
		DonatedNftCount:        1,
		StakingRewardWei:       "1000000000000000000",
		DonatedErc20TokenCount: 2,
	}
	got, err := mapUserPendingWinnings(userCursorAlice, record)
	if err != nil {
		t.Fatalf("mapUserPendingWinnings: %v", err)
	}
	if got.Address != userCursorAlice || got.RaffleEthWei != "60000000000000000" ||
		got.ChronoWarriorEthWei != "80000000000000000" || got.DonatedNftCount != 1 ||
		got.StakingRewardWei != "1000000000000000000" || got.DonatedErc20TokenCount != 2 {
		t.Fatalf("pending winnings = %+v", got)
	}

	invalid := map[string]func(*cgstore.UserPendingWinningsRecord){
		"negative raffle":     func(r *cgstore.UserPendingWinningsRecord) { r.RaffleEthWei = "-1" },
		"empty chrono":        func(r *cgstore.UserPendingWinningsRecord) { r.ChronoWarriorEthWei = "" },
		"malformed staking":   func(r *cgstore.UserPendingWinningsRecord) { r.StakingRewardWei = "wat" },
		"negative nft count":  func(r *cgstore.UserPendingWinningsRecord) { r.DonatedNftCount = -1 },
		"negative token rows": func(r *cgstore.UserPendingWinningsRecord) { r.DonatedErc20TokenCount = -1 },
	}
	for name, corrupt := range invalid {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			bad := record
			corrupt(&bad)
			if _, err := mapUserPendingWinnings(userCursorAlice, bad); err == nil {
				t.Fatalf("mapUserPendingWinnings accepted %s", name)
			}
		})
	}

	zero := zeroUserPendingWinnings(userCursorAlice)
	mapped, err := mapUserPendingWinnings(userCursorAlice, cgstore.UserPendingWinningsRecord{
		RaffleEthWei:        "0",
		ChronoWarriorEthWei: "0",
		StakingRewardWei:    "0",
	})
	if err != nil || mapped != zero {
		t.Fatalf("zero shapes diverge: mapped=%+v static=%+v err=%v", mapped, zero, err)
	}
}

func TestSignedAmount(t *testing.T) {
	t.Parallel()
	for input, want := range map[string]string{
		"0":    "0",
		"-1":   "-1",
		"0042": "42",
		"9000": "9000",
	} {
		got, err := signedAmount(input)
		if err != nil || got != want {
			t.Errorf("signedAmount(%q) = %q, %v; want %q", input, got, err, want)
		}
	}
	for _, input := range []string{"", "wat", "1.5", "0x10"} {
		if _, err := signedAmount(input); err == nil {
			t.Errorf("signedAmount(%q) succeeded", input)
		}
	}
}
