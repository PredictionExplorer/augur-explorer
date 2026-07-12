package v2

import (
	"errors"
	"math"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

func TestMapContractAddressRegistry(t *testing.T) {
	t.Parallel()
	record := validContractAddressRecord()
	got, err := mapContractAddressRegistry(record)
	if err != nil {
		t.Fatalf("mapContractAddressRegistry: %v", err)
	}
	if got.CosmicGameAddress != ethcommon.HexToAddress(record.CosmicGameAddr).Hex() ||
		got.MarketplaceAddress != ethcommon.HexToAddress(record.MarketplaceAddr).Hex() {
		t.Fatalf("registry = %+v", got)
	}

	record.ImplementationAddr = ethcommon.Address{}.Hex()
	if _, err := mapContractAddressRegistry(record); err == nil {
		t.Fatal("zero implementation address accepted")
	}

	invalidFields := map[string]func(*cgprimitives.CosmicGameContractAddrs){
		"CosmicGame":         func(r *cgprimitives.CosmicGameContractAddrs) { r.CosmicGameAddr = "bad" },
		"CosmicSignature":    func(r *cgprimitives.CosmicGameContractAddrs) { r.CosmicSignatureAddr = "bad" },
		"CosmicToken":        func(r *cgprimitives.CosmicGameContractAddrs) { r.CosmicTokenAddr = "bad" },
		"CosmicDAO":          func(r *cgprimitives.CosmicGameContractAddrs) { r.CosmicDaoAddr = "bad" },
		"charity wallet":     func(r *cgprimitives.CosmicGameContractAddrs) { r.CharityWalletAddr = "bad" },
		"prizes wallet":      func(r *cgprimitives.CosmicGameContractAddrs) { r.PrizesWalletAddr = "bad" },
		"RandomWalk":         func(r *cgprimitives.CosmicGameContractAddrs) { r.RandomWalkAddr = "bad" },
		"CST staking":        func(r *cgprimitives.CosmicGameContractAddrs) { r.StakingWalletCSTAddr = "bad" },
		"RandomWalk staking": func(r *cgprimitives.CosmicGameContractAddrs) { r.StakingWalletRWalkAddr = "bad" },
		"marketing":          func(r *cgprimitives.CosmicGameContractAddrs) { r.MarketingWalletAddr = "bad" },
		"marketplace":        func(r *cgprimitives.CosmicGameContractAddrs) { r.MarketplaceAddr = "bad" },
		"implementation":     func(r *cgprimitives.CosmicGameContractAddrs) { r.ImplementationAddr = "bad" },
	}
	for name, mutate := range invalidFields {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validContractAddressRecord()
			mutate(&record)
			if _, err := mapContractAddressRegistry(record); err == nil {
				t.Fatal("invalid registry address accepted")
			}
		})
	}
}

func TestMapContractConfigurationMechanics(t *testing.T) {
	t.Parallel()
	v1Snapshot := validContractSnapshot()
	v1, err := mapContractConfiguration(v1Snapshot)
	if err != nil {
		t.Fatalf("map v1 configuration: %v", err)
	}
	if v1.MechanicsVersion != ContractMechanicsVersionV1 ||
		v1.CstRoundStartAuctionMode != CstRoundStartAuctionModeDivisor ||
		v1.CstBidRewardMode != CstBidRewardFixed ||
		v1.FixedCstBidRewardWei == nil || *v1.FixedCstBidRewardWei != "55000000000000000000" ||
		v1.CstBidRewardMultiplier != nil ||
		v1.CstDutchAuctionDurationChangeDivisor != nil ||
		v1.EthBidPriceIncreaseDivisor != "100" {
		t.Fatalf("v1 configuration = %+v", v1)
	}

	v2Snapshot := v1Snapshot
	v2Snapshot.MechanicsVersion = 2
	v2Snapshot.ConstantsMechanicsVersion = 2
	v2Snapshot.VariablesMechanicsVersion = 2
	v2Snapshot.RoundStartAuctionLength = 28800
	v2Snapshot.CSTAuctionDurationChangeDivisor = 33
	v2Snapshot.BidCSTRewardMultiplier = "7"
	v2Snapshot.TokenReward = "99"
	v2, err := mapContractConfiguration(v2Snapshot)
	if err != nil {
		t.Fatalf("map v2 configuration: %v", err)
	}
	if v2.MechanicsVersion != ContractMechanicsVersionV2 ||
		v2.CstRoundStartAuctionMode != CstRoundStartAuctionModeDurationSeconds ||
		v2.CstBidRewardMode != CstBidRewardDynamic ||
		v2.CstBidRewardMultiplier == nil || *v2.CstBidRewardMultiplier != "7" ||
		v2.FixedCstBidRewardWei != nil ||
		v2.CstDutchAuctionDurationChangeDivisor == nil ||
		*v2.CstDutchAuctionDurationChangeDivisor != 33 {
		t.Fatalf("v2 configuration = %+v", v2)
	}
}

func TestMapContractConfigurationFailures(t *testing.T) {
	t.Parallel()
	unavailable := validContractSnapshot()
	unavailable.ConfigurationReady = false
	if _, err := mapContractConfiguration(unavailable); !errors.Is(err, errCachedLiveUnavailable) {
		t.Fatalf("unavailable error = %v", err)
	}
	tests := map[string]func(*contractstate.Snapshot){
		"bad divisor":       func(s *contractstate.Snapshot) { s.PriceIncrease = "error" },
		"zero charity":      func(s *contractstate.Snapshot) { s.CharityAddr = ethcommon.Address{} },
		"bad percentage":    func(s *contractstate.Snapshot) { s.PrizePercentage = 101 },
		"unknown mechanics": func(s *contractstate.Snapshot) { s.MechanicsVersion = 0 },
		"v1 change divisor": func(s *contractstate.Snapshot) { s.CSTAuctionDurationChangeDivisor = 1 },
		"bad timeout":       func(s *contractstate.Snapshot) { s.TimeoutClaimPrize = -1 },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			snapshot := validContractSnapshot()
			mutate(&snapshot)
			if _, err := mapContractConfiguration(snapshot); err == nil {
				t.Fatal("invalid configuration accepted")
			}
		})
	}
}

func TestCachedContractMappersReturnExactInvariantErrors(t *testing.T) {
	t.Parallel()

	mapConfiguration := func(snapshot contractstate.Snapshot) error {
		_, err := mapContractConfiguration(snapshot)
		return err
	}
	mapBalances := func(snapshot contractstate.Snapshot) error {
		_, err := mapContractBalances(snapshot)
		return err
	}
	mapPrices := func(snapshot contractstate.Snapshot) error {
		_, err := mapCurrentBidPrices(snapshot)
		return err
	}
	mapWinners := func(snapshot contractstate.Snapshot) error {
		_, err := mapCurrentSpecialWinners(snapshot)
		return err
	}
	makeV2 := func(snapshot *contractstate.Snapshot) {
		snapshot.MechanicsVersion = 2
		snapshot.ConstantsMechanicsVersion = 2
		snapshot.VariablesMechanicsVersion = 2
		snapshot.CSTAuctionDurationChangeDivisor = 33
		snapshot.FixedCSTBidReward = ""
		snapshot.BidCSTRewardMultiplier = "7"
	}

	tests := []struct {
		name   string
		mapper func(contractstate.Snapshot) error
		mutate func(*contractstate.Snapshot)
		want   string
	}{
		{
			name:   "time divisor",
			mapper: mapConfiguration,
			mutate: func(snapshot *contractstate.Snapshot) { snapshot.TimeIncrease = "bad" },
			want:   `time-increment divisor: invalid non-negative decimal "bad"`,
		},
		{
			name:   "v1 fixed reward",
			mapper: mapConfiguration,
			mutate: func(snapshot *contractstate.Snapshot) { snapshot.FixedCSTBidReward = "" },
			want:   "fixed CST bid reward: amount is empty",
		},
		{
			name:   "v2 change divisor",
			mapper: mapConfiguration,
			mutate: func(snapshot *contractstate.Snapshot) {
				makeV2(snapshot)
				snapshot.CSTAuctionDurationChangeDivisor = 0
			},
			want: "v2 mechanics lacks an auction-change divisor",
		},
		{
			name:   "v2 reward multiplier",
			mapper: mapConfiguration,
			mutate: func(snapshot *contractstate.Snapshot) {
				makeV2(snapshot)
				snapshot.BidCSTRewardMultiplier = "bad"
			},
			want: `CST bid reward multiplier: invalid non-negative decimal "bad"`,
		},
		{
			name:   "game balance",
			mapper: mapBalances,
			mutate: func(snapshot *contractstate.Snapshot) { snapshot.CosmicGameBalance = "-1" },
			want:   `CosmicGame balance: invalid non-negative decimal "-1"`,
		},
		{
			name:   "charity balance generation",
			mapper: mapBalances,
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.BalanceCharityAddr = ethcommon.HexToAddress(
					"0x7000000000000000000000000000000000000007",
				)
			},
			want: "charity balance belongs to another address",
		},
		{
			name:   "CST bid price",
			mapper: mapPrices,
			mutate: func(snapshot *contractstate.Snapshot) { snapshot.NextCSTBidPrice = "bad" },
			want:   `next CST bid price: invalid non-negative decimal "bad"`,
		},
		{
			name:   "auction progress",
			mapper: mapPrices,
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.CSTAuctionElapsed = snapshot.CSTAuctionDuration + 1
			},
			want: "cached Dutch-auction progress is inconsistent",
		},
		{
			name:   "endurance address",
			mapper: mapWinners,
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.SpecialWinners.EnduranceChampionAddress = "bad"
			},
			want: "invalid endurance champion address",
		},
		{
			name:   "zero CST bidder attachments",
			mapper: mapWinners,
			mutate: func(snapshot *contractstate.Snapshot) {
				snapshot.SpecialWinners.LastCstBidderAddress = ethcommon.Address{}.Hex()
			},
			want: "zero last-CST bidder has attached values",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			snapshot := validContractSnapshot()
			test.mutate(&snapshot)
			err := test.mapper(snapshot)
			if err == nil || err.Error() != test.want {
				t.Fatalf("error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestMapContractBalancesAndBidPrices(t *testing.T) {
	t.Parallel()
	snapshot := validContractSnapshot()
	balances, err := mapContractBalances(snapshot)
	if err != nil {
		t.Fatalf("mapContractBalances: %v", err)
	}
	if balances.CosmicGameBalanceWei != "10" || balances.CharityBalanceWei != "20" {
		t.Fatalf("balances = %+v", balances)
	}
	prices, err := mapCurrentBidPrices(snapshot)
	if err != nil {
		t.Fatalf("mapCurrentBidPrices: %v", err)
	}
	if prices.NextEthBidPriceWei != "1000" || prices.NextCstBidPriceWei != "2000" ||
		prices.NextCstBidRewardWei != "55000000000000000000" ||
		prices.EthAuctionDurationSeconds != 100 || prices.CstAuctionElapsedSeconds != 20 {
		t.Fatalf("prices = %+v", prices)
	}

	snapshot.BalancesReady = false
	if _, err := mapContractBalances(snapshot); !errors.Is(err, errCachedLiveUnavailable) {
		t.Fatalf("balance availability error = %v", err)
	}
	snapshot = validContractSnapshot()
	snapshot.BalanceCharityAddr = ethcommon.HexToAddress(
		"0x7000000000000000000000000000000000000007",
	)
	if _, err := mapContractBalances(snapshot); err == nil {
		t.Fatal("cross-generation charity balance accepted")
	}
	snapshot = validContractSnapshot()
	snapshot.ETHAuctionElapsed = snapshot.ETHAuctionDuration + 1
	if _, err := mapCurrentBidPrices(snapshot); err == nil {
		t.Fatal("invalid auction progress accepted")
	}
}

func TestMapCurrentSpecialWinners(t *testing.T) {
	t.Parallel()
	snapshot := validContractSnapshot()
	got, err := mapCurrentSpecialWinners(snapshot)
	if err != nil {
		t.Fatalf("mapCurrentSpecialWinners: %v", err)
	}
	if got.Round != 3 || got.SourceBlockNumber != 100 ||
		got.EnduranceChampion == nil || got.ChronoWarrior == nil ||
		got.LastBidder == nil || got.LastCstBidder == nil ||
		got.LastCstBidder.EventLogId == nil || *got.LastCstBidder.EventLogId != 99 {
		t.Fatalf("special winners = %+v", got)
	}

	withoutCST := snapshot
	withoutCST.SpecialWinners.LastCstBidderAddress = ethcommon.Address{}.Hex()
	withoutCST.SpecialWinners.HasLastCstBidderLastBidTime = false
	withoutCST.SpecialWinners.HasLastCstBidEventLogId = false
	withoutCST.SpecialWinners.LastCstBidderLastBidTime = 0
	withoutCST.SpecialWinners.LastCstBidEventLogId = 0
	got, err = mapCurrentSpecialWinners(withoutCST)
	if err != nil || got.LastCstBidder != nil {
		t.Fatalf("optional last CST bidder = %+v, %v", got.LastCstBidder, err)
	}
}

func TestMapCurrentSpecialWinnersFailures(t *testing.T) {
	t.Parallel()
	unavailable := validContractSnapshot()
	unavailable.SpecialWinnersReady = false
	if _, err := mapCurrentSpecialWinners(unavailable); !errors.Is(err, errCachedLiveUnavailable) {
		t.Fatalf("unavailable error = %v", err)
	}
	tests := map[string]func(*contractstate.LiveSpecialWinners){
		"embedded error":  func(w *contractstate.LiveSpecialWinners) { w.Err = errors.New("bad") },
		"negative round":  func(w *contractstate.LiveSpecialWinners) { w.RoundNum = -1 },
		"overflow block":  func(w *contractstate.LiveSpecialWinners) { w.SourceBlockNumber = math.MaxUint64 },
		"negative source": func(w *contractstate.LiveSpecialWinners) { w.SourceBlockTimeStamp = -1 },
		"bad champion":    func(w *contractstate.LiveSpecialWinners) { w.EnduranceChampionDuration = -1 },
		"bad event ID":    func(w *contractstate.LiveSpecialWinners) { w.LastCstBidEventLogId = 0 },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			snapshot := validContractSnapshot()
			mutate(&snapshot.SpecialWinners)
			if _, err := mapCurrentSpecialWinners(snapshot); err == nil {
				t.Fatal("invalid special winners accepted")
			}
		})
	}
}

func FuzzMapCachedContractConfiguration(f *testing.F) {
	f.Add("100", "50", int64(1), int64(10), int64(11), int64(-1))
	f.Add("100", "7", int64(2), int64(10), int64(28800), int64(33))
	f.Add("error", "-1", int64(0), int64(-1), int64(0), int64(0))
	f.Fuzz(func(
		t *testing.T,
		priceIncrease string,
		reward string,
		mechanics int64,
		percentage int64,
		auctionValue int64,
		changeDivisor int64,
	) {
		snapshot := validContractSnapshot()
		snapshot.PriceIncrease = priceIncrease
		snapshot.TokenReward = reward
		snapshot.MechanicsVersion = mechanics
		snapshot.ConstantsMechanicsVersion = mechanics
		snapshot.VariablesMechanicsVersion = mechanics
		switch mechanics {
		case 1:
			snapshot.FixedCSTBidReward = reward
			snapshot.BidCSTRewardMultiplier = ""
			snapshot.CSTAuctionDurationChangeDivisor = -1
		case 2:
			snapshot.FixedCSTBidReward = ""
			snapshot.BidCSTRewardMultiplier = reward
		}
		snapshot.PrizePercentage = percentage
		snapshot.RoundStartAuctionLength = auctionValue
		snapshot.CSTAuctionDurationChangeDivisor = changeDivisor
		got, err := mapContractConfiguration(snapshot)
		if err == nil && (!got.MechanicsVersion.Valid() ||
			!got.CstBidRewardMode.Valid() ||
			!got.CstRoundStartAuctionMode.Valid()) {
			t.Fatalf("successful mapping produced invalid enums: %+v", got)
		}
	})
}

func FuzzMapCachedSpecialWinners(f *testing.F) {
	f.Add("0x2200000000000000000000000000000000000022", int64(3), uint64(100), int64(1000))
	f.Add("bad", int64(-1), uint64(0), int64(-1))
	f.Fuzz(func(t *testing.T, address string, round int64, block uint64, timestamp int64) {
		snapshot := validContractSnapshot()
		snapshot.SpecialWinners.EnduranceChampionAddress = address
		snapshot.SpecialWinners.RoundNum = round
		snapshot.SpecialWinners.SourceBlockNumber = block
		snapshot.SpecialWinners.SourceBlockTimeStamp = timestamp
		got, err := mapCurrentSpecialWinners(snapshot)
		if err == nil && (got.Round < 0 || got.SourceBlockNumber < 0 || got.SourceBlockTimestamp < 0) {
			t.Fatalf("successful mapping produced invalid source: %+v", got)
		}
	})
}

func validContractAddressRecord() cgprimitives.CosmicGameContractAddrs {
	return cgprimitives.CosmicGameContractAddrs{
		CosmicGameAddr:         "0x2000000000000000000000000000000000000002",
		CosmicSignatureAddr:    "0x3000000000000000000000000000000000000003",
		CosmicTokenAddr:        "0x4000000000000000000000000000000000000004",
		CosmicDaoAddr:          "0x5000000000000000000000000000000000000005",
		CharityWalletAddr:      "0x6000000000000000000000000000000000000006",
		PrizesWalletAddr:       "0x7000000000000000000000000000000000000007",
		RandomWalkAddr:         "0x8000000000000000000000000000000000000008",
		StakingWalletCSTAddr:   "0x9000000000000000000000000000000000000009",
		StakingWalletRWalkAddr: "0x1000000000000000000000000000000000000010",
		MarketingWalletAddr:    "0x1100000000000000000000000000000000000011",
		MarketplaceAddr:        "0x1200000000000000000000000000000000000012",
		ImplementationAddr:     "0x1300000000000000000000000000000000000013",
	}
}

func validContractSnapshot() contractstate.Snapshot {
	charity := ethcommon.HexToAddress("0x6000000000000000000000000000000000000006")
	return contractstate.Snapshot{
		PriceIncrease:                   "0100",
		CharityAddr:                     charity,
		CharityPercentage:               10,
		TokenReward:                     "055000000000000000000",
		FixedCSTBidReward:               "055000000000000000000",
		PrizePercentage:                 25,
		RafflePercentage:                5,
		ChronoPercentage:                7,
		StakingPercentage:               10,
		TimeIncrease:                    "050",
		RaffleEthWinnersBidding:         3,
		RaffleNFTWinnersBidding:         5,
		RaffleNFTWinnersStakingRWalk:    4,
		CSTAuctionDurationChangeDivisor: -1,
		ConstantsReady:                  true,
		ConfigurationReady:              true,
		ConstantsMechanicsVersion:       1,
		BidPrice:                        "01000",
		BlockPinnedBidPrice:             "01000",
		NextCSTBidPrice:                 "02000",
		NextCSTBidReward:                "055000000000000000000",
		ETHAuctionDuration:              100,
		ETHAuctionElapsed:               10,
		CSTAuctionDuration:              200,
		CSTAuctionElapsed:               20,
		BidPricesReady:                  true,
		CharityBalance:                  "020",
		CosmicGameBalance:               "010",
		BalancesReady:                   true,
		VariablesMechanicsVersion:       1,
		BalanceCharityAddr:              charity,
		InitialSecondsUntilPrize:        2,
		TimeoutClaimPrize:               86400,
		RoundStartAuctionLength:         11,
		MechanicsVersion:                1,
		SpecialWinnersReady:             true,
		SpecialWinners: contractstate.LiveSpecialWinners{
			EnduranceChampionAddress:        "0x2200000000000000000000000000000000000022",
			EnduranceChampionDuration:       600,
			EnduranceChampionStartTimeStamp: 1000,
			PrevEnduranceChampionDuration:   300,
			ChronoWarriorAddress:            "0x2300000000000000000000000000000000000023",
			ChronoWarriorDuration:           800,
			ChronoWarriorIsLive:             true,
			LastBidderAddress:               "0x2100000000000000000000000000000000000021",
			LastBidderLastBidTime:           1200,
			LastCstBidderAddress:            "0x2300000000000000000000000000000000000023",
			LastCstBidderLastBidTime:        1100,
			LastCstBidEventLogId:            99,
			HasLastCstBidderLastBidTime:     true,
			HasLastCstBidEventLogId:         true,
			RoundNum:                        3,
			SourceBlockNumber:               100,
			SourceBlockTimeStamp:            1300,
		},
	}
}
