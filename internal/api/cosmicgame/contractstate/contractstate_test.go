package contractstate

import (
	"context"
	"errors"
	"math"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	cg "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	cgp "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// Test fixture addresses and values served by the V1 game stub.
var (
	gameAddr    = ethcommon.HexToAddress("0x2000000000000000000000000000000000000002")
	tokenAddr   = ethcommon.HexToAddress("0x4000000000000000000000000000000000000004")
	charityAddr = ethcommon.HexToAddress("0xcccc000000000000000000000000000000000ccc")
	bidderAddr  = ethcommon.HexToAddress("0xaaaa000000000000000000000000000000000aaa")
	cstBidder   = ethcommon.HexToAddress("0xbbbb000000000000000000000000000000000bbb")
)

// fakeDB is an in-memory DataSource with settable results and errors.
type fakeDB struct {
	mu              sync.Mutex
	stats           cgp.CGStatistics
	statsErr        error
	roundStart      int64
	roundStartErr   error
	lastCst         int64
	lastCstErr      error
	lastCstMaxBlock int64
}

type blockingDB struct{}

func (blockingDB) CosmicGameStatistics(ctx context.Context) (cgp.CGStatistics, error) {
	<-ctx.Done()
	return cgp.CGStatistics{}, ctx.Err()
}

func (blockingDB) RoundStartTimestamp(ctx context.Context, _ uint64) (int64, error) {
	<-ctx.Done()
	return 0, ctx.Err()
}

func (blockingDB) LastCstBidEvtlogForBidderAtBlock(
	ctx context.Context,
	_ int64,
	_ string,
	_ int64,
) (int64, error) {
	<-ctx.Done()
	return 0, ctx.Err()
}

func (f *fakeDB) CosmicGameStatistics(context.Context) (cgp.CGStatistics, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.stats, f.statsErr
}

func (f *fakeDB) RoundStartTimestamp(context.Context, uint64) (int64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.roundStart, f.roundStartErr
}

func (f *fakeDB) LastCstBidEvtlogForBidderAtBlock(
	_ context.Context,
	_ int64,
	_ string,
	maxBlockNum int64,
) (int64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.lastCstMaxBlock = maxBlockNum
	return f.lastCst, f.lastCstErr
}

// newV1GameStub serves every V1 method the refresh cycles and the live
// special-winners fetch read, with fixed values the tests assert against.
func newV1GameStub() *testchain.ContractStub {
	stub := testchain.MustContractStub(cg.CosmicSignatureGameMetaData.ABI)

	// Constants.
	stub.Return("ethBidPriceIncreaseDivisor", big.NewInt(100))
	stub.Return("charityAddress", charityAddr)
	stub.Return("charityEthDonationAmountPercentage", big.NewInt(10))
	stub.Return("cstRewardAmountForBidding", mustBig("100000000000000000000")) // 100 CST
	stub.Return("mainEthPrizeAmountPercentage", big.NewInt(25))
	stub.Return("raffleTotalEthPrizeAmountForBiddersPercentage", big.NewInt(5))
	stub.Return("chronoWarriorEthPrizeAmountPercentage", big.NewInt(7))
	stub.Return("cosmicSignatureNftStakingTotalEthRewardAmountPercentage", big.NewInt(10))
	stub.Return("mainPrizeTimeIncrementIncreaseDivisor", big.NewInt(50))
	stub.Return("numRaffleEthPrizesForBidders", big.NewInt(3))
	stub.Return("numRaffleCosmicSignatureNftsForBidders", big.NewInt(5))
	stub.Return("numRaffleCosmicSignatureNftsForRandomWalkNftStakers", big.NewInt(4))

	// Variables.
	stub.Return("getNextEthBidPrice", mustBig("1010000000000000")) // 0.00101 ETH
	stub.Return("getNextCstBidPrice", mustBig("55000000000000000000"))
	stub.Return("getEthDutchAuctionDurations", big.NewInt(86400), big.NewInt(7200))
	stub.Return("getCstDutchAuctionDurations", big.NewInt(28800), big.NewInt(3600))
	stub.Return("getDurationUntilMainPrize", big.NewInt(3600))
	stub.Return("getMainEthPrizeAmount", mustBig("2000000000000000000"))                           // 2 ETH
	stub.Return("getCosmicSignatureNftStakingTotalEthRewardAmount", mustBig("300000000000000000")) // 0.3 ETH
	stub.Return("getRaffleTotalEthPrizeAmountForBidders", mustBig("150000000000000000"))           // 0.15 ETH
	stub.Return("roundNum", big.NewInt(7))
	stub.Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(3600000000))
	stub.Return("lastBidderAddress", bidderAddr)
	stub.Return("initialDurationUntilMainPrizeDivisor", big.NewInt(2))
	stub.Return("timeoutDurationToClaimMainPrize", big.NewInt(86400))
	stub.Return("cstDutchAuctionDurationDivisor", big.NewInt(11)) // V1 marker

	// Live special-winners reads.
	stub.Return("tryGetCurrentChampions", bidderAddr, big.NewInt(900), cstBidder, big.NewInt(1000))
	stub.Return("enduranceChampionStartTimeStamp", big.NewInt(1767229000))
	stub.Return("prevEnduranceChampionDuration", big.NewInt(200))
	stub.Return("enduranceChampionAddress", bidderAddr)
	stub.Return("enduranceChampionDuration", big.NewInt(900))
	stub.Return("chronoWarriorDuration", big.NewInt(1000))
	stub.Return("lastCstBidderAddress", cstBidder)
	stub.Handle("biddersInfo", func(args []any) ([]any, error) {
		addr, ok := args[1].(ethcommon.Address)
		if !ok {
			return nil, errors.New("biddersInfo: unexpected argument type")
		}
		ts := big.NewInt(1767230000) // last bidder
		if addr == cstBidder {
			ts = big.NewInt(1767229500)
		}
		return []any{big.NewInt(0), big.NewInt(0), ts}, nil
	})

	return stub
}

// newV2GameStub serves a V2-generation contract: the V1-only CST auction
// methods are absent from its ABI, the V2 ones answer.
func newV2GameStub() *testchain.ContractStub {
	stub := testchain.MustContractStub(cg.CosmicSignatureGameV2MetaData.ABI)
	stub.Return("ethBidPriceIncreaseDivisor", big.NewInt(100))
	stub.Return("charityAddress", charityAddr)
	stub.Return("charityEthDonationAmountPercentage", big.NewInt(10))
	stub.Return("mainEthPrizeAmountPercentage", big.NewInt(25))
	stub.Return("raffleTotalEthPrizeAmountForBiddersPercentage", big.NewInt(5))
	stub.Return("chronoWarriorEthPrizeAmountPercentage", big.NewInt(7))
	stub.Return("cosmicSignatureNftStakingTotalEthRewardAmountPercentage", big.NewInt(10))
	stub.Return("mainPrizeTimeIncrementIncreaseDivisor", big.NewInt(50))
	stub.Return("numRaffleEthPrizesForBidders", big.NewInt(3))
	stub.Return("numRaffleCosmicSignatureNftsForBidders", big.NewInt(5))
	stub.Return("numRaffleCosmicSignatureNftsForRandomWalkNftStakers", big.NewInt(4))
	stub.Return("cstDutchAuctionDuration", big.NewInt(28800))
	stub.Return("cstDutchAuctionDurationChangeDivisor", big.NewInt(33))
	stub.Return("getBidCstRewardAmount", mustBig("99000000000000000000")) // 99 CST
	stub.Return("bidCstRewardAmountMultiplier", big.NewInt(7))
	stub.Return("getNextEthBidPrice", mustBig("1010000000000000"))
	stub.Return("getNextCstBidPrice", mustBig("55000000000000000000"))
	stub.Return("getEthDutchAuctionDurations", big.NewInt(86400), big.NewInt(7200))
	stub.Return(
		"getCstDutchAuctionDurations",
		big.NewInt(28800),
		big.NewInt(int64(testchain.BlockTime(0))-100),
	)
	stub.Return("getDurationUntilMainPrize", big.NewInt(3600))
	stub.Return("getMainEthPrizeAmount", mustBig("2000000000000000000"))
	stub.Return("getCosmicSignatureNftStakingTotalEthRewardAmount", mustBig("300000000000000000"))
	stub.Return("getRaffleTotalEthPrizeAmountForBidders", mustBig("150000000000000000"))
	stub.Return("roundNum", big.NewInt(7))
	stub.Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(3600000000))
	stub.Return("lastBidderAddress", bidderAddr)
	stub.Return("initialDurationUntilMainPrizeDivisor", big.NewInt(2))
	stub.Return("timeoutDurationToClaimMainPrize", big.NewInt(86400))
	return stub
}

func mustBig(s string) *big.Int {
	v, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("bad big.Int literal: " + s)
	}
	return v
}

func defaultFakeDB() *fakeDB {
	return &fakeDB{
		stats:      cgp.CGStatistics{TotalPrizes: 3, TotalBids: 12},
		roundStart: 1767225700,
		lastCst:    5099,
	}
}

// newTestState dials the chain and builds a State around it.
func newTestState(t *testing.T, chain *testchain.Chain, db DataSource) *State {
	t.Helper()
	chain.EnsureBlock(0)
	client, err := ethclient.Dial(chain.URL())
	if err != nil {
		t.Fatalf("dialing test chain: %v", err)
	}
	t.Cleanup(client.Close)
	return newStateForClient(t, client, db)
}

func newStateForClient(t *testing.T, client *ethclient.Client, db DataSource) *State {
	t.Helper()
	s, err := New(Config{
		EthClient: client,
		DB:        db,
		Addrs: Addresses{
			CosmicGame:      gameAddr,
			CosmicSignature: ethcommon.HexToAddress("0x3000000000000000000000000000000000000003"),
			CosmicToken:     tokenAddr,
			CharityWallet:   ethcommon.HexToAddress("0x6000000000000000000000000000000000000006"),
			MarketingWallet: ethcommon.HexToAddress("0x1100000000000000000000000000000000000011"),
		},
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return s
}

func TestNewValidation(t *testing.T) {
	if _, err := New(Config{DB: defaultFakeDB()}); err == nil {
		t.Fatal("expected error for missing EthClient")
	}
	chain := testchain.New(t)
	client, err := ethclient.Dial(chain.URL())
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer client.Close()
	if _, err := New(Config{EthClient: client}); err == nil {
		t.Fatal("expected error for missing DB")
	}
}

func TestLoadInitialHappyPath(t *testing.T) {
	chain := testchain.New(t)
	chain.RegisterCall(gameAddr, newV1GameStub().Handler())
	s := newTestState(t, chain, defaultFakeDB())

	s.LoadInitial(context.Background())
	snap := s.Snapshot()

	// Addresses pass through.
	if snap.Addrs.CosmicGame != gameAddr || snap.Addrs.CosmicToken != tokenAddr {
		t.Fatalf("unexpected addresses: %+v", snap.Addrs)
	}

	// Constants.
	if snap.PriceIncrease != "100" {
		t.Errorf("PriceIncrease = %q, want 100", snap.PriceIncrease)
	}
	if snap.CharityAddr != charityAddr {
		t.Errorf("CharityAddr = %v, want %v", snap.CharityAddr, charityAddr)
	}
	if snap.CharityPercentage != 10 {
		t.Errorf("CharityPercentage = %d, want 10", snap.CharityPercentage)
	}
	if snap.TokenReward != "100000000000000000000" {
		t.Errorf("TokenReward = %q", snap.TokenReward)
	}
	if snap.FixedCSTBidReward != snap.TokenReward || snap.BidCSTRewardMultiplier != "" {
		t.Errorf("V1 reward configuration = fixed:%q multiplier:%q",
			snap.FixedCSTBidReward, snap.BidCSTRewardMultiplier)
	}
	if snap.PrizePercentage != 25 || snap.RafflePercentage != 5 || snap.ChronoPercentage != 7 || snap.StakingPercentage != 10 {
		t.Errorf("percentages = %d/%d/%d/%d", snap.PrizePercentage, snap.RafflePercentage, snap.ChronoPercentage, snap.StakingPercentage)
	}
	if snap.TimeIncrease != "50" {
		t.Errorf("TimeIncrease = %q, want 50", snap.TimeIncrease)
	}
	if snap.RaffleEthWinnersBidding != 3 || snap.RaffleNFTWinnersBidding != 5 || snap.RaffleNFTWinnersStakingRWalk != 4 {
		t.Errorf("raffle winner counts = %d/%d/%d", snap.RaffleEthWinnersBidding, snap.RaffleNFTWinnersBidding, snap.RaffleNFTWinnersStakingRWalk)
	}

	// Variables.
	if snap.BidPrice != "1010000000000000" {
		t.Errorf("BidPrice = %q", snap.BidPrice)
	}
	if snap.BlockPinnedBidPrice != snap.BidPrice {
		t.Errorf("BlockPinnedBidPrice = %q, want %q", snap.BlockPinnedBidPrice, snap.BidPrice)
	}
	if math.Abs(snap.BidPriceEth-0.00101) > 1e-12 {
		t.Errorf("BidPriceEth = %v, want 0.00101", snap.BidPriceEth)
	}
	if snap.NextCSTBidPrice != "55000000000000000000" ||
		snap.NextCSTBidReward != "100000000000000000000" ||
		snap.ETHAuctionDuration != 86400 || snap.ETHAuctionElapsed != 7200 ||
		snap.CSTAuctionDuration != 28800 || snap.CSTAuctionElapsed != 3600 ||
		!snap.BidPricesReady {
		t.Errorf("cached bid prices = %+v", snap)
	}
	if snap.PrizeClaimTimestamp != 3600 {
		t.Errorf("PrizeClaimTimestamp = %d, want 3600", snap.PrizeClaimTimestamp)
	}
	if snap.PrizeAmount != "2000000000000000000" || snap.PrizeAmountEth != 2.0 {
		t.Errorf("PrizeAmount = %q / %v", snap.PrizeAmount, snap.PrizeAmountEth)
	}
	if snap.StakingAmount != "300000000000000000" || math.Abs(snap.StakingAmountEth-0.3) > 1e-12 {
		t.Errorf("StakingAmount = %q / %v", snap.StakingAmount, snap.StakingAmountEth)
	}
	if snap.RaffleAmount != "150000000000000000" || math.Abs(snap.RaffleAmountEth-0.15) > 1e-12 {
		t.Errorf("RaffleAmount = %q / %v", snap.RaffleAmount, snap.RaffleAmountEth)
	}
	if snap.RoundNum != 7 {
		t.Errorf("RoundNum = %d, want 7", snap.RoundNum)
	}
	if snap.MainPrizeTimeIncrement != "3600000000" {
		t.Errorf("MainPrizeTimeIncrement = %q", snap.MainPrizeTimeIncrement)
	}
	if snap.LastBidder != bidderAddr {
		t.Errorf("LastBidder = %v", snap.LastBidder)
	}
	if snap.InitialSecondsUntilPrize != 2 || snap.TimeoutClaimPrize != 86400 {
		t.Errorf("InitialSecondsUntilPrize/TimeoutClaimPrize = %d/%d", snap.InitialSecondsUntilPrize, snap.TimeoutClaimPrize)
	}
	if snap.RoundStartAuctionLength != 11 {
		t.Errorf("RoundStartAuctionLength = %d, want 11 (V1 divisor)", snap.RoundStartAuctionLength)
	}
	if snap.CharityBalance != "0" || snap.CharityBalanceEth != 0 {
		t.Errorf("CharityBalance = %q / %v", snap.CharityBalance, snap.CharityBalanceEth)
	}
	if snap.CosmicGameBalance != "0" || !snap.BalancesReady {
		t.Errorf("CosmicGameBalance/BalancesReady = %q/%v", snap.CosmicGameBalance, snap.BalancesReady)
	}

	// Mechanics: the V1 divisor answered, the V2 change divisor did not.
	if snap.MechanicsVersion != mechanicsV1 {
		t.Errorf("MechanicsVersion = %d, want V1", snap.MechanicsVersion)
	}
	if snap.CSTAuctionDurationChangeDivisor != -1 {
		t.Errorf("CSTAuctionDurationChangeDivisor = %d, want -1 on V1", snap.CSTAuctionDurationChangeDivisor)
	}
	if !snap.ConstantsReady || !snap.ConfigurationReady {
		t.Errorf("constant/configuration readiness = %v/%v", snap.ConstantsReady, snap.ConfigurationReady)
	}

	// Database aggregates.
	if snap.Stats.TotalPrizes != 3 || snap.Stats.TotalBids != 12 {
		t.Errorf("Stats = %+v", snap.Stats)
	}
	if snap.RoundStartTimestamp != 1767225700 {
		t.Errorf("RoundStartTimestamp = %d", snap.RoundStartTimestamp)
	}

	// Live balance of the game contract (testchain serves 0 wei).
	if got := s.CosmicGameBalanceEth(context.Background()); got != 0 {
		t.Errorf("CosmicGameBalanceEth = %v, want 0", got)
	}
}

// TestLoadInitialRPCUnavailable pins the failure sentinels: with eth_call
// failing (no registered handler) but the node reachable, every field takes
// the documented degraded value — the shape the dashboard renders when the
// contract is unreachable.
func TestLoadInitialRPCUnavailable(t *testing.T) {
	chain := testchain.New(t) // no call handler registered
	s := newTestState(t, chain, defaultFakeDB())

	s.LoadInitial(context.Background())
	snap := s.Snapshot()

	if snap.BidPrice != "error" || snap.BidPriceEth != 0 {
		t.Errorf("BidPrice = %q / %v", snap.BidPrice, snap.BidPriceEth)
	}
	if snap.BlockPinnedBidPrice != "error" || snap.NextCSTBidReward != "error" {
		t.Errorf("pinned price/reward = %q/%q", snap.BlockPinnedBidPrice, snap.NextCSTBidReward)
	}
	if snap.ConstantsReady || snap.BidPricesReady || snap.ConfigurationReady {
		t.Errorf("cache readiness = constants:%v prices:%v configuration:%v",
			snap.ConstantsReady, snap.BidPricesReady, snap.ConfigurationReady)
	}
	if snap.PrizeClaimTimestamp != -1 {
		t.Errorf("PrizeClaimTimestamp = %d, want -1", snap.PrizeClaimTimestamp)
	}
	if snap.PrizeAmount != "error" || snap.RaffleAmount != "error" || snap.StakingAmount != "error" {
		t.Errorf("amounts = %q/%q/%q", snap.PrizeAmount, snap.RaffleAmount, snap.StakingAmount)
	}
	if snap.RoundNum != -1 {
		t.Errorf("RoundNum = %d, want -1", snap.RoundNum)
	}
	if snap.MainPrizeTimeIncrement != "error" || snap.PriceIncrease != "error" || snap.TimeIncrease != "error" || snap.TokenReward != "error" {
		t.Errorf("strings = %q/%q/%q/%q", snap.MainPrizeTimeIncrement, snap.PriceIncrease, snap.TimeIncrease, snap.TokenReward)
	}
	if snap.LastBidder != (ethcommon.Address{}) || snap.CharityAddr != (ethcommon.Address{}) {
		t.Errorf("addresses should stay zero: %v / %v", snap.LastBidder, snap.CharityAddr)
	}
	if snap.PrizePercentage != -1 || snap.RafflePercentage != -1 || snap.ChronoPercentage != -1 || snap.StakingPercentage != -1 {
		t.Errorf("percentages = %d/%d/%d/%d, want -1", snap.PrizePercentage, snap.RafflePercentage, snap.ChronoPercentage, snap.StakingPercentage)
	}
	if snap.CharityPercentage != 0 {
		t.Errorf("CharityPercentage = %d, want 0", snap.CharityPercentage)
	}
	if snap.InitialSecondsUntilPrize != -1 || snap.TimeoutClaimPrize != -1 || snap.RoundStartAuctionLength != -1 {
		t.Errorf("durations = %d/%d/%d, want -1", snap.InitialSecondsUntilPrize, snap.TimeoutClaimPrize, snap.RoundStartAuctionLength)
	}
	if snap.RaffleEthWinnersBidding != -1 || snap.RaffleNFTWinnersBidding != -1 || snap.RaffleNFTWinnersStakingRWalk != -1 {
		t.Errorf("winner counts should be -1")
	}
	if snap.CSTAuctionDurationChangeDivisor != -1 {
		t.Errorf("CSTAuctionDurationChangeDivisor = %d, want -1", snap.CSTAuctionDurationChangeDivisor)
	}
	if snap.MechanicsVersion != mechanicsUnknown {
		t.Errorf("MechanicsVersion = %d, want unknown", snap.MechanicsVersion)
	}
	// eth_getBalance still answers on the reachable node.
	if snap.CharityBalance != "0" {
		t.Errorf("CharityBalance = %q, want 0", snap.CharityBalance)
	}
	if snap.BalancesReady {
		t.Error("BalancesReady = true without a charity address")
	}
	// The DB half is independent of the RPC node.
	if snap.Stats.TotalPrizes != 3 || snap.RoundStartTimestamp != 1767225700 {
		t.Errorf("DB fields = %+v / %d", snap.Stats, snap.RoundStartTimestamp)
	}
}

// TestLoadInitialNodeDown covers the fully dead node: balance reads fail too.
func TestLoadInitialNodeDown(t *testing.T) {
	chain, stop := testchain.Start()
	url := chain.URL()
	stop() // server is gone; the client dials lazily so calls fail

	client, err := ethclient.Dial(url)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer client.Close()
	s := newStateForClient(t, client, defaultFakeDB())

	s.LoadInitial(context.Background())
	snap := s.Snapshot()

	if snap.CharityBalance != "error" || snap.CharityBalanceEth != 0 {
		t.Errorf("CharityBalance = %q / %v, want error / 0", snap.CharityBalance, snap.CharityBalanceEth)
	}
	if snap.CosmicGameBalance != "error" || snap.BalancesReady {
		t.Errorf("CosmicGameBalance/BalancesReady = %q/%v", snap.CosmicGameBalance, snap.BalancesReady)
	}
	if snap.BidPrice != "error" || snap.RoundNum != -1 {
		t.Errorf("BidPrice/RoundNum = %q/%d", snap.BidPrice, snap.RoundNum)
	}
	if got := s.CosmicGameBalanceEth(context.Background()); !math.IsNaN(got) {
		t.Errorf("CosmicGameBalanceEth = %v, want NaN", got)
	}
}

func TestLiveCacheReadinessRecovers(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(50)
	s := newTestState(t, chain, defaultFakeDB())
	s.LoadInitial(context.Background())
	initial := s.Snapshot()
	if initial.ConfigurationReady || initial.BidPricesReady || initial.BalancesReady {
		t.Fatalf("unavailable cache marked ready: %+v", initial)
	}

	chain.RegisterCall(gameAddr, newV1GameStub().Handler())
	s.refreshVariables(context.Background())
	s.refreshConstants(context.Background())
	s.refreshBalances(context.Background())
	recovered := s.Snapshot()
	if !recovered.ConfigurationReady || !recovered.BidPricesReady || !recovered.BalancesReady {
		t.Fatalf("cache did not recover: configuration=%v prices=%v balances=%v",
			recovered.ConfigurationReady, recovered.BidPricesReady, recovered.BalancesReady)
	}
}

func TestConfigurationReadinessRequiresBothRefreshGroups(t *testing.T) {
	chain := testchain.New(t)
	stub := newV1GameStub()
	stub.Handle("timeoutDurationToClaimMainPrize", func([]any) ([]any, error) {
		return nil, errors.New("variable read failed")
	})
	chain.RegisterCall(gameAddr, stub.Handler())
	s := newTestState(t, chain, defaultFakeDB())

	s.LoadInitial(context.Background())
	snap := s.Snapshot()
	if !snap.ConstantsReady || snap.ConfigurationReady {
		t.Fatalf("readiness with variable failure = constants:%v configuration:%v",
			snap.ConstantsReady, snap.ConfigurationReady)
	}
	if !snap.BidPricesReady {
		t.Fatal("independent bid-price cache should remain ready")
	}
}

func TestConstantsFailureKeepsValueButMarksConfigurationUnavailable(t *testing.T) {
	chain := testchain.New(t)
	stub := newV1GameStub()
	chain.RegisterCall(gameAddr, stub.Handler())
	s := newTestState(t, chain, defaultFakeDB())
	s.LoadInitial(context.Background())
	originalCharity := s.Snapshot().CharityAddr

	stub.Handle("charityAddress", func([]any) ([]any, error) {
		return nil, errors.New("constant read failed")
	})
	s.refreshConstants(context.Background())
	snap := s.Snapshot()
	if snap.CharityAddr != originalCharity {
		t.Fatalf("charity address changed on failed refresh: %v", snap.CharityAddr)
	}
	if snap.ConstantsReady || snap.ConfigurationReady {
		t.Fatalf("failed constant refresh stayed ready: constants:%v configuration:%v",
			snap.ConstantsReady, snap.ConfigurationReady)
	}
}

func TestCharityAddressChangeInvalidatesBalanceGeneration(t *testing.T) {
	chain := testchain.New(t)
	stub := newV1GameStub()
	chain.RegisterCall(gameAddr, stub.Handler())
	s := newTestState(t, chain, defaultFakeDB())
	s.LoadInitial(context.Background())
	if !s.Snapshot().BalancesReady {
		t.Fatal("initial balances are unavailable")
	}

	newCharity := ethcommon.HexToAddress("0xdddd000000000000000000000000000000000ddd")
	stub.Return("charityAddress", newCharity)
	s.refreshConstants(context.Background())
	changed := s.Snapshot()
	if changed.CharityAddr != newCharity || changed.BalancesReady ||
		changed.BalanceCharityAddr == changed.CharityAddr {
		t.Fatalf("changed charity generation = addr:%v balanceAddr:%v ready:%v",
			changed.CharityAddr, changed.BalanceCharityAddr, changed.BalancesReady)
	}

	s.refreshBalances(context.Background())
	refreshed := s.Snapshot()
	if !refreshed.BalancesReady || refreshed.BalanceCharityAddr != newCharity {
		t.Fatalf("refreshed charity balance generation = addr:%v ready:%v",
			refreshed.BalanceCharityAddr, refreshed.BalancesReady)
	}
}

func TestRewardConfigurationFailureDoesNotCorruptLiveBidPrices(t *testing.T) {
	chain := testchain.New(t)
	stub := newV1GameStub()
	chain.RegisterCall(gameAddr, stub.Handler())
	s := newTestState(t, chain, defaultFakeDB())
	s.LoadInitial(context.Background())
	originalReward := s.Snapshot().NextCSTBidReward

	stub.Handle("cstRewardAmountForBidding", func([]any) ([]any, error) {
		return nil, errors.New("fixed reward read failed")
	})
	s.refreshConstants(context.Background())
	s.refreshVariables(context.Background())
	snap := s.Snapshot()
	if snap.ConfigurationReady || !snap.BidPricesReady ||
		snap.NextCSTBidReward != originalReward || snap.TokenReward != "error" {
		t.Fatalf("reward failure state = configuration:%v prices:%v next:%q legacy:%q",
			snap.ConfigurationReady, snap.BidPricesReady, snap.NextCSTBidReward, snap.TokenReward)
	}
}

func TestDBStatsFailureKeepsPreviousValues(t *testing.T) {
	chain := testchain.New(t)
	chain.RegisterCall(gameAddr, newV1GameStub().Handler())
	db := defaultFakeDB()
	s := newTestState(t, chain, db)
	s.LoadInitial(context.Background())

	// A failed statistics read keeps both DB fields.
	db.mu.Lock()
	db.statsErr = errors.New("db down")
	db.mu.Unlock()
	s.refreshDBStats(context.Background())
	snap := s.Snapshot()
	if snap.Stats.TotalPrizes != 3 || snap.RoundStartTimestamp != 1767225700 {
		t.Fatalf("failed stats refresh must keep previous values, got %+v / %d", snap.Stats, snap.RoundStartTimestamp)
	}

	// A failed round-start read keeps the previous timestamp while the fresh
	// statistics stand (legacy behavior).
	db.mu.Lock()
	db.statsErr = nil
	db.stats = cgp.CGStatistics{TotalPrizes: 4}
	db.roundStartErr = errors.New("db down")
	db.mu.Unlock()
	s.refreshDBStats(context.Background())
	snap = s.Snapshot()
	if snap.Stats.TotalPrizes != 4 {
		t.Fatalf("stats should have refreshed, got %+v", snap.Stats)
	}
	if snap.RoundStartTimestamp != 1767225700 {
		t.Fatalf("round start timestamp should keep previous value, got %d", snap.RoundStartTimestamp)
	}
}

func TestDBRefreshDeadline(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(0)
	client, err := ethclient.Dial(chain.URL())
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer client.Close()
	s, err := New(Config{
		EthClient:     client,
		DB:            blockingDB{},
		DBReadTimeout: 5 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	start := time.Now()
	s.refreshDBStats(context.Background())
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Fatalf("DB refresh ignored timeout: %v", elapsed)
	}
}

func TestMechanicsV2Detection(t *testing.T) {
	chain := testchain.New(t)
	chain.RegisterCall(gameAddr, newV2GameStub().Handler())
	s := newTestState(t, chain, defaultFakeDB())

	s.LoadInitial(context.Background())
	snap := s.Snapshot()

	if snap.MechanicsVersion != mechanicsV2 {
		t.Fatalf("MechanicsVersion = %d, want V2", snap.MechanicsVersion)
	}
	if snap.RoundStartAuctionLength != 28800 {
		t.Errorf("RoundStartAuctionLength = %d, want 28800 (V2 duration)", snap.RoundStartAuctionLength)
	}
	if snap.CSTAuctionDurationChangeDivisor != 33 {
		t.Errorf("CSTAuctionDurationChangeDivisor = %d, want 33", snap.CSTAuctionDurationChangeDivisor)
	}
	if snap.TokenReward != "99000000000000000000" {
		t.Errorf("TokenReward = %q, want the V2 computed reward", snap.TokenReward)
	}
	if snap.NextCSTBidReward != snap.TokenReward ||
		snap.FixedCSTBidReward != "" ||
		snap.BidCSTRewardMultiplier != "7" ||
		!snap.ConfigurationReady || !snap.BidPricesReady {
		t.Errorf("V2 reward/cache state = next:%q fixed:%q multiplier:%q configuration:%v prices:%v",
			snap.NextCSTBidReward, snap.FixedCSTBidReward, snap.BidCSTRewardMultiplier,
			snap.ConfigurationReady, snap.BidPricesReady)
	}
}

func TestV2ChangeDivisorFailureMarksConstantsUnavailable(t *testing.T) {
	chain := testchain.New(t)
	stub := newV2GameStub()
	stub.Handle("cstDutchAuctionDurationChangeDivisor", func([]any) ([]any, error) {
		return nil, errors.New("divisor read failed")
	})
	chain.RegisterCall(gameAddr, stub.Handler())
	s := newTestState(t, chain, defaultFakeDB())

	s.LoadInitial(context.Background())
	snap := s.Snapshot()
	if snap.ConstantsReady || snap.ConfigurationReady {
		t.Fatalf("failed V2 divisor stayed ready: constants=%v configuration=%v",
			snap.ConstantsReady, snap.ConfigurationReady)
	}
}

// TestMechanicsUpgradeFlip simulates the proxy upgrading V1 -> V2 between
// refreshes: the cached version must flip once the V1 reads stop answering.
func TestMechanicsUpgradeFlip(t *testing.T) {
	chain := testchain.New(t)
	chain.RegisterCall(gameAddr, newV1GameStub().Handler())
	s := newTestState(t, chain, defaultFakeDB())
	s.LoadInitial(context.Background())
	if got := s.Snapshot().MechanicsVersion; got != mechanicsV1 {
		t.Fatalf("MechanicsVersion = %d, want V1 before upgrade", got)
	}

	chain.RegisterCall(gameAddr, newV2GameStub().Handler()) // proxy upgraded
	s.refreshConstants(context.Background())
	if snap := s.Snapshot(); snap.ConfigurationReady ||
		snap.ConstantsMechanicsVersion != mechanicsV2 ||
		snap.VariablesMechanicsVersion != mechanicsV1 {
		t.Fatalf("mixed mechanics snapshot was exposed as ready: %+v", snap)
	}
	s.refreshVariables(context.Background())

	snap := s.Snapshot()
	if snap.MechanicsVersion != mechanicsV2 {
		t.Fatalf("MechanicsVersion = %d, want V2 after upgrade", snap.MechanicsVersion)
	}
	if snap.TokenReward != "99000000000000000000" {
		t.Errorf("TokenReward = %q, want V2 value", snap.TokenReward)
	}
	if snap.RoundStartAuctionLength != 28800 {
		t.Errorf("RoundStartAuctionLength = %d, want V2 duration", snap.RoundStartAuctionLength)
	}
	if !snap.ConfigurationReady || snap.BidCSTRewardMultiplier != "7" {
		t.Errorf("upgraded configuration = ready:%v multiplier:%q",
			snap.ConfigurationReady, snap.BidCSTRewardMultiplier)
	}
}

func TestLiveReadHelpersWithoutBindings(t *testing.T) {
	chain := testchain.New(t)
	s := newTestState(t, chain, defaultFakeDB())
	opts := &bind.CallOpts{}

	if got := s.roundStartCSTAuctionSetting(nil, nil, opts); got != -1 {
		t.Errorf("roundStartCSTAuctionSetting(nil, nil) = %d, want -1", got)
	}
	if got, ok := s.cstAuctionDurationChangeDivisor(nil, nil, opts); got != -1 || ok {
		t.Errorf("cstAuctionDurationChangeDivisor(nil, nil) = %d,%v, want -1,false", got, ok)
	}
	if _, err := s.tokenReward(nil, nil, opts); err == nil {
		t.Error("tokenReward(nil, nil) should error")
	}
}

func TestNormalizeAuctionProgress(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		duration     *big.Int
		second       *big.Int
		blockTime    int64
		mayBeStart   bool
		wantDuration int64
		wantElapsed  int64
		wantOK       bool
	}{
		{"elapsed", big.NewInt(100), big.NewInt(25), 1000, false, 100, 25, true},
		{"before activation", big.NewInt(100), big.NewInt(-10), 1000, false, 100, 0, true},
		{"past floor", big.NewInt(100), big.NewInt(150), 1000, false, 100, 100, true},
		{"start timestamp", big.NewInt(100), big.NewInt(1_700_000_000), 1_700_000_025, true, 100, 25, true},
		{"future start", big.NewInt(100), big.NewInt(1_700_000_100), 1_700_000_025, true, 100, 0, true},
		{"overflow", new(big.Int).Lsh(big.NewInt(1), 80), big.NewInt(1), 1000, false, 0, 0, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			duration, elapsed, ok := normalizeAuctionProgress(
				test.duration,
				test.second,
				test.blockTime,
				test.mayBeStart,
			)
			if duration != test.wantDuration || elapsed != test.wantElapsed || ok != test.wantOK {
				t.Fatalf("normalizeAuctionProgress = %d,%d,%v; want %d,%d,%v",
					duration, elapsed, ok, test.wantDuration, test.wantElapsed, test.wantOK)
			}
		})
	}
}

func TestSetBidPriceWriteBack(t *testing.T) {
	chain := testchain.New(t)
	s := newTestState(t, chain, defaultFakeDB())

	s.SetBidPrice("42", 42e-18)
	snap := s.Snapshot()
	if snap.BidPrice != "42" || snap.BidPriceEth != 42e-18 {
		t.Fatalf("SetBidPrice not visible in snapshot: %q / %v", snap.BidPrice, snap.BidPriceEth)
	}
	if snap.BlockPinnedBidPrice != "" || snap.BidPricesReady {
		t.Fatalf("SetBidPrice mutated v2 cache: pinned=%q ready=%v",
			snap.BlockPinnedBidPrice, snap.BidPricesReady)
	}
}

func TestRunRefreshesAndStops(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(0)
	stub := newV1GameStub()
	chain.RegisterCall(gameAddr, stub.Handler())

	client, err := ethclient.Dial(chain.URL())
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer client.Close()
	s, err := New(Config{
		EthClient:              client,
		DB:                     defaultFakeDB(),
		Addrs:                  Addresses{CosmicGame: gameAddr},
		ConstantsInterval:      time.Millisecond,
		VariablesInterval:      time.Millisecond,
		DBStatsInterval:        time.Millisecond,
		SpecialWinnersInterval: time.Millisecond,
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	s.LoadInitial(context.Background())
	if got := s.Snapshot().RoundNum; got != 7 {
		t.Fatalf("RoundNum after LoadInitial = %d, want 7", got)
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		s.Run(ctx)
		close(done)
	}()

	// A new round starts on chain; Run must pick it up.
	stub.Return("roundNum", big.NewInt(8))
	deadline := time.After(10 * time.Second)
	for s.Snapshot().RoundNum != 8 {
		select {
		case <-deadline:
			t.Fatal("Run never refreshed RoundNum to 8")
		case <-time.After(2 * time.Millisecond):
		}
	}

	cancel()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("Run did not stop on context cancellation")
	}
}

// TestSnapshotConcurrency exercises concurrent readers, the write-back
// setter and the refresh loops under -race.
func TestSnapshotConcurrency(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(0)
	chain.RegisterCall(gameAddr, newV1GameStub().Handler())

	client, err := ethclient.Dial(chain.URL())
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer client.Close()
	s, err := New(Config{
		EthClient:              client,
		DB:                     defaultFakeDB(),
		Addrs:                  Addresses{CosmicGame: gameAddr},
		ConstantsInterval:      time.Millisecond,
		VariablesInterval:      time.Millisecond,
		DBStatsInterval:        time.Millisecond,
		SpecialWinnersInterval: time.Millisecond,
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	s.LoadInitial(context.Background())

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Run(ctx)
	}()
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 200; j++ {
				snap := s.Snapshot()
				if snap.BidPrice == "" && snap.RoundNum == 7 {
					// Refresh batches must never publish a half-written
					// variables group.
					t.Error("torn snapshot: RoundNum set while BidPrice empty")
					return
				}
				s.SetBidPrice(snap.BidPrice, snap.BidPriceEth)
			}
		}()
	}
	time.Sleep(50 * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestFetchLiveSpecialWinnersHappyPath(t *testing.T) {
	chain := testchain.New(t)
	chain.RegisterCall(gameAddr, newV1GameStub().Handler())
	chain.EnsureBlock(50) // tip: timestamp BaseTime + 5000
	db := defaultFakeDB()
	s := newTestState(t, chain, db)
	s.LoadInitial(context.Background())

	cached := s.Snapshot()
	if !cached.SpecialWinnersReady || cached.SpecialWinners.SourceBlockNumber != 50 {
		t.Fatalf("cached special winners = %+v ready=%v",
			cached.SpecialWinners, cached.SpecialWinnersReady)
	}

	out := s.FetchLiveSpecialWinners(context.Background())
	if out.Err != nil {
		t.Fatalf("unexpected error: %v", out.Err)
	}
	if out.RoundNum != 7 {
		t.Errorf("RoundNum = %d, want 7", out.RoundNum)
	}
	if out.SourceBlockNumber != 50 || out.SourceBlockTimeStamp != int64(testchain.BlockTime(50)) {
		t.Errorf("source block = %d @ %d", out.SourceBlockNumber, out.SourceBlockTimeStamp)
	}
	if out.LastBidderAddress != bidderAddr.String() || out.LastBidderLastBidTime != 1767230000 {
		t.Errorf("last bidder = %s @ %d", out.LastBidderAddress, out.LastBidderLastBidTime)
	}
	if out.EnduranceChampionAddress != bidderAddr.String() || out.EnduranceChampionDuration != 900 {
		t.Errorf("endurance champion = %s / %d", out.EnduranceChampionAddress, out.EnduranceChampionDuration)
	}
	// lastBidDuration = 1767230600-1767230000 = 600 < stored 900: no live
	// overtake, the stored anchor stands.
	if out.EnduranceChampionStartTimeStamp != 1767229000 || out.PrevEnduranceChampionDuration != 200 {
		t.Errorf("anchor = %d / %d", out.EnduranceChampionStartTimeStamp, out.PrevEnduranceChampionDuration)
	}
	// chrono segment: start 1767229200, current duration 1400 > stored 1000.
	if !out.ChronoWarriorIsLive {
		t.Error("ChronoWarriorIsLive = false, want true")
	}
	if out.ChronoWarriorAddress != cstBidder.String() || out.ChronoWarriorDuration != 1000 {
		t.Errorf("chrono warrior = %s / %d", out.ChronoWarriorAddress, out.ChronoWarriorDuration)
	}
	if !out.HasLastCstBidderLastBidTime || out.LastCstBidderLastBidTime != 1767229500 {
		t.Errorf("cst bidder last bid = %v / %d", out.HasLastCstBidderLastBidTime, out.LastCstBidderLastBidTime)
	}
	if !out.HasLastCstBidEventLogId || out.LastCstBidEventLogId != 5099 {
		t.Errorf("cst bid evtlog = %v / %d", out.HasLastCstBidEventLogId, out.LastCstBidEventLogId)
	}
	db.mu.Lock()
	maxBlock := db.lastCstMaxBlock
	db.mu.Unlock()
	if maxBlock != 50 {
		t.Errorf("CST bid lookup max block = %d, want 50", maxBlock)
	}
}

func TestFetchLiveSpecialWinnersCstBidNotFound(t *testing.T) {
	chain := testchain.New(t)
	chain.RegisterCall(gameAddr, newV1GameStub().Handler())
	chain.EnsureBlock(50)
	db := defaultFakeDB()
	db.lastCstErr = store.ErrNotFound
	s := newTestState(t, chain, db)
	s.LoadInitial(context.Background())

	out := s.FetchLiveSpecialWinners(context.Background())
	if out.Err != nil {
		t.Fatalf("unexpected error: %v", out.Err)
	}
	if out.HasLastCstBidEventLogId {
		t.Error("HasLastCstBidEventLogId should be false on ErrNotFound")
	}
}

func TestSpecialWinnersCacheKeepsOptionalDBFailureNonFatal(t *testing.T) {
	chain := testchain.New(t)
	chain.RegisterCall(gameAddr, newV1GameStub().Handler())
	chain.EnsureBlock(50)
	db := defaultFakeDB()
	db.lastCstErr = errors.New("database unavailable")
	s := newTestState(t, chain, db)

	s.LoadInitial(context.Background())
	snap := s.Snapshot()
	if !snap.SpecialWinnersReady || snap.SpecialWinners.Err != nil {
		t.Fatalf("special-winners cache = %+v ready=%v",
			snap.SpecialWinners, snap.SpecialWinnersReady)
	}
	if snap.SpecialWinners.HasLastCstBidEventLogId {
		t.Fatal("optional CST event-log ID survived a DB failure")
	}
}

func TestSpecialWinnersCacheRequiresBlockPinnedRound(t *testing.T) {
	chain := testchain.New(t)
	stub := newV1GameStub()
	stub.Handle("roundNum", func([]any) ([]any, error) {
		return nil, errors.New("round read failed")
	})
	chain.RegisterCall(gameAddr, stub.Handler())
	chain.EnsureBlock(50)
	s := newTestState(t, chain, defaultFakeDB())

	s.LoadInitial(context.Background())
	if s.Snapshot().SpecialWinnersReady {
		t.Fatal("special-winners cache is ready without a live round number")
	}
}

func TestFetchLiveSpecialWinnersRPCFailure(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(50) // header fetch works, contract reads fail
	s := newTestState(t, chain, defaultFakeDB())

	out := s.FetchLiveSpecialWinners(context.Background())
	if out.Err == nil {
		t.Fatal("expected error when contract reads fail")
	}
}

func TestSpecialWinnersCacheRecovers(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(50)
	s := newTestState(t, chain, defaultFakeDB())

	s.refreshVariables(context.Background())
	s.refreshSpecialWinners(context.Background())
	if s.Snapshot().SpecialWinnersReady {
		t.Fatal("special-winners cache is ready after failed contract reads")
	}

	chain.RegisterCall(gameAddr, newV1GameStub().Handler())
	s.refreshSpecialWinners(context.Background())
	snap := s.Snapshot()
	if !snap.SpecialWinnersReady || snap.SpecialWinners.Err != nil ||
		snap.SpecialWinners.SourceBlockNumber != 50 {
		t.Fatalf("recovered special-winners cache = %+v ready=%v",
			snap.SpecialWinners, snap.SpecialWinnersReady)
	}
}
