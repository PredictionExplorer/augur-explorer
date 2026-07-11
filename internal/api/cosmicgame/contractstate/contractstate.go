// Package contractstate owns the cached CosmicGame contract and database
// state shared by the JSON API handlers.
//
// The legacy implementation kept ~70 package-level mutable globals in
// internal/api/cosmicgame/state.go, refreshed by three unkillable
// `for { refresh; sleep }` goroutines. This package replaces that with one
// injected State component: LoadInitial performs the synchronous startup
// loads, Run drives the periodic refresh loops until its context is
// cancelled, and handlers read a consistent value copy via Snapshot.
package contractstate

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	cgp "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

// Default refresh intervals, matching the legacy
// CONTRACT_CONSTANTS_REFRESH_TIME / CONTRACT_VARIABLES_REFRESH_TIME values.
const (
	DefaultConstantsInterval      = 5 * time.Minute
	DefaultVariablesInterval      = 5 * time.Second
	DefaultDBStatsInterval        = 5 * time.Second
	DefaultSpecialWinnersInterval = 30 * time.Second
	DefaultRPCReadTimeout         = 10 * time.Second
	DefaultDBReadTimeout          = 10 * time.Second
)

// DataSource is the narrow slice of the CosmicGame repository the state
// component needs; *cosmicgame.Repo satisfies it.
type DataSource interface {
	// CosmicGameStatistics computes the global statistics aggregate.
	CosmicGameStatistics(ctx context.Context) (cgp.CGStatistics, error)
	// RoundStartTimestamp returns the first-bid timestamp of the round.
	RoundStartTimestamp(ctx context.Context, roundNum uint64) (int64, error)
	// LastCstBidEvtlogForBidderAtBlock returns the event log id of the
	// bidder's latest CST bid in the round at or before the source block.
	LastCstBidEvtlogForBidderAtBlock(
		ctx context.Context,
		roundNum int64,
		bidderAddr string,
		maxBlockNum int64,
	) (int64, error)
}

// Addresses are the deployed contract addresses, loaded once from the
// database contract registry at startup.
type Addresses struct {
	CosmicGame      ethcommon.Address
	CosmicSignature ethcommon.Address
	CosmicToken     ethcommon.Address
	CharityWallet   ethcommon.Address
	MarketingWallet ethcommon.Address
}

// Config carries the dependencies of a State.
type Config struct {
	// EthClient performs the contract reads. Required.
	EthClient *ethclient.Client
	// DB serves the database-backed aggregates. Required.
	DB DataSource
	// Addrs are the contract addresses from the registry table.
	Addrs Addresses
	// Info and Error receive the refresh diagnostics the legacy loggers
	// carried; nil loggers discard.
	Info  *log.Logger
	Error *log.Logger

	// Refresh intervals for Run; zero values pick the defaults above.
	ConstantsInterval      time.Duration
	VariablesInterval      time.Duration
	DBStatsInterval        time.Duration
	SpecialWinnersInterval time.Duration
	RPCReadTimeout         time.Duration
	DBReadTimeout          time.Duration
}

// Snapshot is one consistent value copy of the cached state. Field groups
// mirror the refresh cycles; the zero value of each field is what
// handlers observed before the first successful refresh, and the documented
// "error" / -1 / 0 sentinels are what a failed refresh leaves behind.
type Snapshot struct {
	// Addrs is the contract address registry (immutable after New).
	Addrs Addresses

	// Contract constants: owner-tunable parameters, refreshed every
	// ConstantsInterval.
	PriceIncrease                   string
	CharityAddr                     ethcommon.Address
	CharityPercentage               int64
	TokenReward                     string
	FixedCSTBidReward               string
	BidCSTRewardMultiplier          string
	PrizePercentage                 int64
	RafflePercentage                int64
	ChronoPercentage                int64
	StakingPercentage               int64
	TimeIncrease                    string
	RaffleEthWinnersBidding         int64
	RaffleNFTWinnersBidding         int64
	RaffleNFTWinnersStakingRWalk    int64
	CSTAuctionDurationChangeDivisor int64
	ConstantsReady                  bool
	ConfigurationReady              bool
	ConstantsMechanicsVersion       int64

	// Contract variables: per-round live state, refreshed every
	// VariablesInterval.
	BidPrice            string
	BidPriceEth         float64
	BlockPinnedBidPrice string
	// PrizeClaimTimestamp holds the contract's getDurationUntilMainPrize()
	// value; the dashboard has always rendered it as a Unix timestamp, and
	// that legacy quirk is pinned by the parity goldens.
	PrizeClaimTimestamp       int64
	PrizeAmount               string
	PrizeAmountEth            float64
	RaffleAmount              string
	RaffleAmountEth           float64
	StakingAmount             string
	StakingAmountEth          float64
	RoundNum                  int64
	MainPrizeTimeIncrement    string // microseconds, decimal string
	LastBidder                ethcommon.Address
	InitialSecondsUntilPrize  int64
	TimeoutClaimPrize         int64
	RoundStartAuctionLength   int64
	CharityBalance            string
	CharityBalanceEth         float64
	CosmicGameBalance         string
	NextCSTBidPrice           string
	NextCSTBidReward          string
	ETHAuctionDuration        int64
	ETHAuctionElapsed         int64
	CSTAuctionDuration        int64
	CSTAuctionElapsed         int64
	BidPricesReady            bool
	BalancesReady             bool
	VariablesMechanicsVersion int64
	BalanceCharityAddr        ethcommon.Address

	// Special-winner standings are refreshed independently so their
	// multi-call latest-block read cannot delay the other variable cache.
	SpecialWinners      LiveSpecialWinners
	SpecialWinnersReady bool

	// Database aggregates, refreshed every DBStatsInterval.
	Stats               cgp.CGStatistics
	RoundStartTimestamp int64

	// MechanicsVersion is the detected contract generation:
	// 0 unknown, 1 V1, 2 V2.
	MechanicsVersion int64
}

// State is the refreshable contract/database state cache. Create it with
// New, perform the synchronous startup loads with LoadInitial, then either
// run the refresh loops via Run (production) or refresh on demand (tests).
type State struct {
	client *ethclient.Client
	db     DataSource
	addrs  Addresses
	info   *log.Logger
	errlog *log.Logger

	constantsInterval      time.Duration
	variablesInterval      time.Duration
	dbStatsInterval        time.Duration
	specialWinnersInterval time.Duration
	rpcReadTimeout         time.Duration
	dbReadTimeout          time.Duration

	contractRefreshMu       sync.Mutex
	dbStatsRefreshMu        sync.Mutex
	specialWinnersRefreshMu sync.Mutex

	mu   sync.RWMutex
	snap Snapshot // Addrs left zero; filled from s.addrs on read
}

// New validates cfg and builds a State. No I/O happens until LoadInitial.
func New(cfg Config) (*State, error) {
	if cfg.EthClient == nil {
		return nil, errors.New("contractstate: Config.EthClient is required")
	}
	if cfg.DB == nil {
		return nil, errors.New("contractstate: Config.DB is required")
	}
	if cfg.Info == nil {
		cfg.Info = discardLogger()
	}
	if cfg.Error == nil {
		cfg.Error = discardLogger()
	}
	if cfg.ConstantsInterval <= 0 {
		cfg.ConstantsInterval = DefaultConstantsInterval
	}
	if cfg.VariablesInterval <= 0 {
		cfg.VariablesInterval = DefaultVariablesInterval
	}
	if cfg.DBStatsInterval <= 0 {
		cfg.DBStatsInterval = DefaultDBStatsInterval
	}
	if cfg.SpecialWinnersInterval <= 0 {
		cfg.SpecialWinnersInterval = DefaultSpecialWinnersInterval
	}
	if cfg.RPCReadTimeout <= 0 {
		cfg.RPCReadTimeout = DefaultRPCReadTimeout
	}
	if cfg.DBReadTimeout <= 0 {
		cfg.DBReadTimeout = DefaultDBReadTimeout
	}
	return &State{
		client:                 cfg.EthClient,
		db:                     cfg.DB,
		addrs:                  cfg.Addrs,
		info:                   cfg.Info,
		errlog:                 cfg.Error,
		constantsInterval:      cfg.ConstantsInterval,
		variablesInterval:      cfg.VariablesInterval,
		dbStatsInterval:        cfg.DBStatsInterval,
		specialWinnersInterval: cfg.SpecialWinnersInterval,
		rpcReadTimeout:         cfg.RPCReadTimeout,
		dbReadTimeout:          cfg.DBReadTimeout,
	}, nil
}

func discardLogger() *log.Logger {
	return log.New(discardWriter{}, "", 0)
}

type discardWriter struct{}

func (discardWriter) Write(p []byte) (int, error) { return len(p), nil }

// LoadInitial performs the synchronous startup loads in the legacy order
// (contract variables, database aggregates, contract constants). Failures
// are logged and leave the documented sentinel values; startup proceeds so
// the DB-backed routes stay available while the RPC node is down.
func (s *State) LoadInitial(ctx context.Context) {
	s.refreshVariables(ctx)
	s.refreshSpecialWinners(ctx)
	s.refreshDBStats(ctx)
	s.refreshConstants(ctx)
	s.refreshBalances(ctx)
}

// Run drives the four periodic refresh loops until ctx is cancelled. The
// initial synchronous loads are LoadInitial's job; each loop here waits one
// interval before its first refresh.
func (s *State) Run(ctx context.Context) {
	var wg sync.WaitGroup
	loop := func(interval time.Duration, refresh func(context.Context)) {
		defer wg.Done()
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				refresh(ctx)
			}
		}
	}
	wg.Add(4)
	go loop(s.constantsInterval, s.refreshConstants)
	go loop(s.variablesInterval, func(ctx context.Context) {
		s.refreshVariables(ctx)
		snapshot := s.Snapshot()
		if snapshot.VariablesMechanicsVersion != mechanicsUnknown &&
			snapshot.ConstantsMechanicsVersion != snapshot.VariablesMechanicsVersion {
			s.refreshConstants(ctx)
		}
		s.refreshBalances(ctx)
	})
	go loop(s.dbStatsInterval, s.refreshDBStats)
	go loop(s.specialWinnersInterval, s.refreshSpecialWinners)
	wg.Wait()
}

// Snapshot returns a consistent value copy of the cached state.
func (s *State) Snapshot() Snapshot {
	s.mu.RLock()
	out := s.snap
	s.mu.RUnlock()
	out.Addrs = s.addrs
	return out
}

// SetBidPrice overwrites the cached next-bid price. The dashboard handler
// uses it to write back a live read that succeeded while the cache held the
// "error" sentinel, so subsequent requests see the recovered value early.
func (s *State) SetBidPrice(price string, priceEth float64) {
	s.mu.Lock()
	s.snap.BidPrice = price
	s.snap.BidPriceEth = priceEth
	s.snap.BidPricesReady = bidPricesReady(s.snap)
	s.mu.Unlock()
}

// logf records a refresh diagnostic on both legacy log streams, matching the
// Error.Print + Info.Print pairs the previous implementation emitted.
func (s *State) logf(format string, args ...any) {
	s.errlog.Printf(format, args...)
	s.info.Printf(format, args...)
}
