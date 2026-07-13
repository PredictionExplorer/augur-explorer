// Package cosmicgame serves the frozen v1 CosmicGame JSON API. The module
// is an injected API value (no package-level state); construction loads the
// contract registry and the synchronous contract-state snapshot, and route
// registration is a method the shared router constructor calls.
package cosmicgame

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgdb "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// API carries the injected dependencies of the v1 CosmicGame module: the
// context-first repository and base store, the contract-state cache the
// handlers snapshot per request, the Ethereum clients for the few live
// reads, and the legacy file loggers.
type API struct {
	repo      *cgdb.Repo
	store     *store.Store
	state     *contractstate.State
	ethClient *ethclient.Client
	rpc       *ethrpc.Client
	info      *log.Logger
	errlog    *log.Logger
}

// Config carries the dependencies for New. Store must be non-nil; the
// clients may be nil in tests that never exercise live contract reads.
type Config struct {
	Store     *store.Store
	EthClient *ethclient.Client
	RPCClient *ethrpc.Client
	Info      *log.Logger
	Error     *log.Logger
}

// New builds the module and performs the synchronous contract-state loads
// the legacy Init did: it reads the contract registry from the database,
// constructs the contractstate cache and pins the initial snapshot. A
// non-nil error means the module cannot serve (missing database link or
// unreadable contract registry); the caller decides whether that is fatal.
// New does not start the periodic refresh loops — call
// StartBackgroundRefresh for that.
func New(ctx context.Context, cfg Config) (*API, error) {
	a := &API{
		store:     cfg.Store,
		ethClient: cfg.EthClient,
		rpc:       cfg.RPCClient,
		info:      orDiscardLogger(cfg.Info),
		errlog:    orDiscardLogger(cfg.Error),
	}
	if cfg.Store == nil {
		return nil, errors.New("cosmicgame: database link wasn't configured")
	}
	a.repo = cgdb.NewRepo(cfg.Store)

	caddrs, err := a.repo.ContractAddrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("cosmicgame: reading contract addresses: %w", err)
	}
	st, err := contractstate.New(contractstate.Config{
		EthClient: a.ethClient,
		DB:        a.repo,
		Addrs: contractstate.Addresses{
			CosmicGame:      ethcommon.HexToAddress(caddrs.CosmicGameAddr),
			CosmicSignature: ethcommon.HexToAddress(caddrs.CosmicSignatureAddr),
			CosmicToken:     ethcommon.HexToAddress(caddrs.CosmicTokenAddr),
			CharityWallet:   ethcommon.HexToAddress(caddrs.CharityWalletAddr),
			MarketingWallet: ethcommon.HexToAddress(caddrs.MarketingWalletAddr),
		},
		Info:  a.info,
		Error: a.errlog,
	})
	if err != nil {
		return nil, fmt.Errorf("cosmicgame: building contract state: %w", err)
	}
	a.state = st
	a.state.LoadInitial(ctx)
	return a, nil
}

// NewBare returns an unloaded module: handlers answer the legacy
// "Database link wasn't configured" / "module not available" envelopes and
// no contract state exists. The route-drift test uses it to enumerate the
// route table without a database, and the guard tests use it to pin the
// pre-initialization failure shapes.
func NewBare() *API {
	discard := log.New(io.Discard, "", 0)
	return &API{info: discard, errlog: discard}
}

// orDiscardLogger keeps the module loggers non-nil so handlers can log
// unconditionally.
func orDiscardLogger(l *log.Logger) *log.Logger {
	if l == nil {
		return log.New(io.Discard, "", 0)
	}
	return l
}

// ContractState exposes the module's contract-state cache so cmd/apiserver
// can share it with the v2 server (one cache, one set of refresh loops).
func (a *API) ContractState() *contractstate.State {
	return a.state
}

// StartBackgroundRefresh launches the periodic contract/database state
// refresh loops and returns immediately; cancelling ctx stops them. Call it
// after a successful New. The API parity test harness deliberately never
// calls it, keeping snapshots deterministic.
func (a *API) StartBackgroundRefresh(ctx context.Context) {
	if a.state == nil {
		return
	}
	go a.state.Run(ctx)
}

// dbInitialized reports whether the module was constructed with a database
// link (NewBare and nil-store modules answer the legacy guard envelopes).
func (a *API) dbInitialized() bool {
	return a.repo != nil && a.store != nil
}

// RegisterRoutes registers all CosmicGame JSON API routes.
func (a *API) RegisterRoutes(r *httpx.Router) {
	// Statistics
	r.GET("/api/cosmicgame/statistics/dashboard", a.handleDashboard)
	r.GET("/api/cosmicgame/statistics/counters", a.handleRecordCounters)
	r.GET("/api/cosmicgame/statistics/unique/bidders", a.handleUserUniqueBidders)
	r.GET("/api/cosmicgame/statistics/unique/winners", a.handleUserUniqueWinners)
	r.GET("/api/cosmicgame/statistics/leaderboard/roi", a.handleRoiLeaderboard)
	r.GET("/api/cosmicgame/statistics/claims/by_round", a.handleClaimsByRound)
	r.GET("/api/cosmicgame/statistics/claims/detail/{round_num}", a.handleClaimDetailByRound)
	r.GET("/api/cosmicgame/statistics/unique/donors", a.handleUserUniqueDonors)
	r.GET("/api/cosmicgame/statistics/unique/stakers/cst", a.handleUserUniqueStakersCst)
	r.GET("/api/cosmicgame/statistics/unique/stakers/randomwalk", a.handleUserUniqueStakersRwalk)
	r.GET("/api/cosmicgame/statistics/unique/stakers/rwalk", a.handleUserUniqueStakersRwalk) // legacy alias
	r.GET("/api/cosmicgame/statistics/unique/stakers/both", a.handleUserUniqueStakersBoth)
	r.GET("/api/cosmicgame/statistics/bidding/activity/{init_ts}/{fin_ts}/{interval_secs}", a.handleBiddingActivity)
	r.GET("/api/cosmicgame/statistics/bidding/frequency/{init_ts}/{fin_ts}/{interval_secs}", a.handleBiddingFrequency)
	r.GET("/api/cosmicgame/statistics/bidding/top_active_periods/{n}/{init_ts}/{fin_ts}", a.handleBiddingTopActivePeriods)
	r.GET("/api/cosmicgame/statistics/bidding/time_bounds", a.handleBiddingTimeBounds)

	// Rounds
	r.GET("/api/cosmicgame/rounds/list/{offset}/{limit}", a.handlePrizeList)
	r.GET("/api/cosmicgame/rounds/info/{prize_num}", a.handleRoundInfo)
	r.GET("/api/cosmicgame/rounds/current/time", a.handlePrizeCurRoundTime)

	// Prizes
	r.GET("/api/cosmicgame/prizes/history/global/{offset}/{limit}", a.handleGlobalClaimHistoryDetail)
	r.GET("/api/cosmicgame/prizes/history/by_user/{user_addr}/{offset}/{limit}", a.handlePrizeHistoryDetailByUser)
	r.GET("/api/cosmicgame/prizes/eth/all/global", a.handleAllEthDepositsList)
	r.GET("/api/cosmicgame/prizes/eth/all/global/{offset}/{limit}", a.handleAllEthDepositsList)
	r.GET("/api/cosmicgame/prizes/eth/raffle/global", a.handleRaffleEthDepositsList)
	r.GET("/api/cosmicgame/prizes/eth/raffle/global/{offset}/{limit}", a.handleRaffleEthDepositsList)
	r.GET("/api/cosmicgame/prizes/eth/chronowarrior/global", a.handleChronowarriorEthDepositsList)
	r.GET("/api/cosmicgame/prizes/eth/chronowarrior/global/{offset}/{limit}", a.handleChronowarriorEthDepositsList)
	r.GET("/api/cosmicgame/prizes/eth/all/by_user/{user_addr}", a.handleUnifiedEthAllByUser)
	r.GET("/api/cosmicgame/prizes/eth/raffle/by_user/{user_addr}", a.handleUnifiedEthRaffleByUser)
	r.GET("/api/cosmicgame/prizes/eth/chronowarrior/by_user/{user_addr}", a.handleUnifiedEthChronowarriorByUser)
	r.GET("/api/cosmicgame/prizes/eth/unclaimed/by_user/{user_addr}/{offset}/{limit}", a.handleUnclaimedPrizeDepositsByUser)
	r.GET("/api/cosmicgame/prizes/deposits/raffle/by_user/{user_addr}", a.handlePrizeDepositsRaffleEthByUser)
	r.GET("/api/cosmicgame/prizes/deposits/chrono_warrior/by_user/{user_addr}", a.handlePrizeDepositsChronoWarriorByUser)
	r.GET("/api/cosmicgame/prizes/deposits/unclaimed/by_user/{user_addr}/{offset}/{limit}", a.handleUnclaimedPrizeDepositsByUser)

	// Bids
	r.GET("/api/cosmicgame/bid/list/all/{offset}/{limit}", a.handleBidList)
	r.GET("/api/cosmicgame/bid/info/{evtlog_id}", a.handleBidInfo)
	r.GET("/api/cosmicgame/bid/info_by_pos/{round_num}/{bid_position}", a.handleBidInfoByPos)
	r.GET("/api/cosmicgame/bid/with_message/by_round/{round}", a.handleBidWithMessageByRound)
	r.GET("/api/cosmicgame/bid/list/by_round/{round_num}/{sort}/{offset}/{limit}", a.handleBidListByRound)
	r.GET("/api/cosmicgame/bid/bid_type_ratio", a.handleBidTypeRatio)
	r.GET("/api/cosmicgame/bid/used_randomwalk_nfts", a.handleUsedRwalkNfts)
	r.GET("/api/cosmicgame/bid/used_rwalk_nfts", a.handleUsedRwalkNfts) // legacy path (same handler)
	r.GET("/api/cosmicgame/bid/cst_price", a.handleGetCstPrice)
	r.GET("/api/cosmicgame/bid/eth_price", a.handleGetEthPrice)
	r.GET("/api/cosmicgame/bid/current_special_winners", a.handleBidSpecialWinners)
	r.GET("/api/cosmicgame/get_banned_bids", a.handleGetBannedBids)
	// Bid moderation is admin-only: requires X-Admin-Key matching ADMIN_API_KEY
	// (503 when unset — fails closed) and is strictly rate limited.
	adminGuard := common.RequireAdminKey("X-Admin-Key", "ADMIN_API_KEY")
	adminRate := common.RateLimit(2, 5)
	r.POST("/api/cosmicgame/ban_bid", a.handleBanBid, adminRate, adminGuard)
	r.POST("/api/cosmicgame/unban_bid", a.handleUnbanBid, adminRate, adminGuard)

	// CST Tokens
	r.GET("/api/cosmicgame/cst/list/all/{offset}/{limit}", a.handleCosmicSignatureTokenList)
	r.GET("/api/cosmicgame/cst/list/by_user/{user_addr}/{offset}/{limit}", a.handleCosmicSignatureTokenListByUser)
	r.GET("/api/cosmicgame/cst/info/{token_id}", a.handleCosmicSignatureTokenInfo)
	r.GET("/api/cosmicgame/cst/metadata/{token_id}", a.handleCstMetadata)
	r.GET("/cg/metadata/{token_id}", a.handleCstMetadata) // legacy Python path
	r.GET("/api/cosmicgame/cst/names/history/{token_id}", a.handleTokenNameHistory)
	r.GET("/api/cosmicgame/cst/names/search/{name}", a.handleTokenNameSearch)
	r.GET("/api/cosmicgame/cst/names/named_only", a.handleNamedTokensOnly)
	r.GET("/api/cosmicgame/cst/transfers/all/{token_id}/{offset}/{limit}", a.handleTokenOwnershipTransfers)
	r.GET("/api/cosmicgame/cst/transfers/by_user/{user_addr}/{offset}/{limit}", a.handleCosmicSignatureTransfersByUser)
	r.GET("/api/cosmicgame/cst/distribution", a.handleCsTokenDistribution)

	// Cosmic Token (CT)
	r.GET("/api/cosmicgame/ct/balances", a.handleCosmicTokenBalances)
	r.GET("/api/cosmicgame/ct/statistics", a.handleCosmicTokenStatistics)
	r.GET("/api/cosmicgame/ct/summary/by_user/{user_addr}", a.handleCosmicTokenSummaryByUser)
	r.GET("/api/cosmicgame/ct/transfers/by_user/{user_addr}/{offset}/{limit}", a.handleCosmicTokenTransfersByUser)
	r.GET("/api/cosmicgame/ct/total_supply_history_by_bid", a.handleCosmicTokenTotalSupplyHistoryByBid)
	r.GET("/api/cosmicgame/ct/total_supply_history_by_date/{from_date}/{to_date}", a.handleCosmicTokenTotalSupplyHistoryByDate)

	// User
	r.GET("/api/cosmicgame/user/info/{user_addr}", a.handleUserInfo)
	r.GET("/api/cosmicgame/user/notif_red_box/{user_addr}", a.handleUserGlobalWinnings)
	r.GET("/api/cosmicgame/user/balances/{user_addr}", a.handleUserBalances)

	// Donations
	r.GET("/api/cosmicgame/donations/eth/simple/list/{offset}/{limit}", a.handleDonationsCgSimpleList)
	r.GET("/api/cosmicgame/donations/eth/simple/by_round/{round_num}", a.handleDonationsCgSimpleByRound)
	r.GET("/api/cosmicgame/donations/eth/with_info/list/{offset}/{limit}", a.handleDonationsCgWithInfoList)
	r.GET("/api/cosmicgame/donations/eth/with_info/by_round/{round_num}", a.handleDonationsCgWithInfoByRound)
	r.GET("/api/cosmicgame/donations/eth/with_info/info/{record_id}", a.handleDonationsCgWithInfoRecordInfo)
	r.GET("/api/cosmicgame/donations/eth/by_user/{user_addr}", a.handleDonationsByUser)
	r.GET("/api/cosmicgame/donations/eth/both/by_round/{round_num}", a.handleDonationsCgBothByRound)
	r.GET("/api/cosmicgame/donations/eth/both/all", a.handleDonationsCgBothAll)
	r.GET("/api/cosmicgame/donations/charity/deposits", a.handleCharityDonationsDeposits)
	r.GET("/api/cosmicgame/donations/charity/cg_deposits", a.handleCharityCosmicgameDeposits)
	r.GET("/api/cosmicgame/donations/charity/voluntary", a.handleCharityVoluntaryDeposits)
	r.GET("/api/cosmicgame/donations/charity/withdrawals", a.handleCharityDonationsWithdrawals)
	r.GET("/api/cosmicgame/donations/nft/list/{offset}/{limit}", a.handleDonationsNftList)
	r.GET("/api/cosmicgame/donations/nft/info/{record_id}", a.handleDonatedNftInfo)
	r.GET("/api/cosmicgame/donations/nft/by_user/{user_addr}", a.handleNftDonationsByUser)
	r.GET("/api/cosmicgame/donations/nft/claims", a.handleDonatedNftClaimsAll)
	r.GET("/api/cosmicgame/donations/nft/claims/{offset}/{limit}", a.handleDonatedNftClaimsAll)
	r.GET("/api/cosmicgame/donations/nft/claims/by_user/{user_addr}", a.handleDonatedNftClaimsByUser)
	r.GET("/api/cosmicgame/donations/nft/statistics", a.handleNftDonationStats)
	r.GET("/api/cosmicgame/donations/nft/by_round/{prize_num}", a.handleNftDonationsByPrize)
	r.GET("/api/cosmicgame/donations/nft/by_token/{token_addr}", a.handleNftDonationsByToken)
	r.GET("/api/cosmicgame/donations/nft/unclaimed/by_round/{prize_num}", a.handleUnclaimedDonatedNftsByPrize)
	r.GET("/api/cosmicgame/donations/nft/unclaimed/by_user/{user_addr}", a.handleUnclaimedDonatedNftsByUser)
	r.GET("/api/cosmicgame/donations/erc20/by_round/detailed/{round_num}", a.handleDonationsErc20ByRoundDetailed)
	r.GET("/api/cosmicgame/donations/erc20/by_round/all/{round_num}", a.handleDonationsErc20ByRoundAll)
	r.GET("/api/cosmicgame/donations/erc20/by_round/summarized/{round_num}", a.handleDonationsErc20ByRoundSummarized)
	r.GET("/api/cosmicgame/donations/erc20/donated/by_user/{user_addr}", a.handleDonationsErc20DonatedByUser)
	r.GET("/api/cosmicgame/donations/erc20/by_user/{user_addr}", a.handleDonationsErc20ByUser)
	r.GET("/api/cosmicgame/donations/erc20/global/{offset}/{limit}", a.handleDonationsErc20Global)
	r.GET("/api/cosmicgame/donations/erc20/info/{record_id}", a.handleDonatedErc20Info)
	r.GET("/api/cosmicgame/donations/erc20/claims", a.handleErc20ClaimsGlobal)
	r.GET("/api/cosmicgame/donations/erc20/claims/{offset}/{limit}", a.handleErc20ClaimsGlobal)
	r.GET("/api/cosmicgame/donations/erc20/claims/by_user/{user_addr}", a.handleErc20ClaimsByUser)
	r.GET("/api/cosmicgame/donations/erc20/claims/by_round/{round_num}", a.handleErc20ClaimsByRound)

	// Raffle
	r.GET("/api/cosmicgame/raffle/deposits/list", a.handlePrizeDepositsList)
	r.GET("/api/cosmicgame/raffle/deposits/list/{offset}/{limit}", a.handlePrizeDepositsList)
	r.GET("/api/cosmicgame/raffle/deposits/by_round/{round_num}", a.handlePrizeDepositsByRound)
	r.GET("/api/cosmicgame/eth_deposits/all/list/{offset}/{limit}", a.handleAllEthDepositsList)
	r.GET("/api/cosmicgame/eth_deposits/raffle_eth/list/{offset}/{limit}", a.handleRaffleEthDepositsList)
	r.GET("/api/cosmicgame/eth_deposits/chronowarrior_eth/list/{offset}/{limit}", a.handleChronowarriorEthDepositsList)
	r.GET("/api/cosmicgame/raffle/nft/all/list", a.handleRaffleNftWinnersList)
	r.GET("/api/cosmicgame/raffle/nft/all/list/{offset}/{limit}", a.handleRaffleNftWinnersList)
	r.GET("/api/cosmicgame/raffle/nft/by_round/{round_num}", a.handleRaffleNftWinnersByRound)
	r.GET("/api/cosmicgame/raffle/nft/by_user/{user_addr}", a.handleUserRaffleNftWinnings)

	// Staking CST
	r.GET("/api/cosmicgame/staking/cst/staked_tokens/all", a.handleStakedTokensCstGlobal)
	r.GET("/api/cosmicgame/staking/cst/staked_tokens/by_user/{user_addr}", a.handleStakedTokensCstByUser)
	r.GET("/api/cosmicgame/staking/cst/actions/global/{offset}/{limit}", a.handleStakingActionsCstGlobal)
	r.GET("/api/cosmicgame/staking/cst/actions/by_user/{user_addr}/{offset}/{limit}", a.handleStakingCstActionsByUser)
	r.GET("/api/cosmicgame/staking/cst/actions/info/{action_id}", a.handleStakingActionCstInfo)
	r.GET("/api/cosmicgame/staking/cst/rewards/global", a.handleStakingCstRewardsGlobal)
	r.GET("/api/cosmicgame/staking/cst/rewards/to_claim/by_user/{user_addr}", a.handleStakingCstRewardsToClaimByUser)
	r.GET("/api/cosmicgame/staking/cst/rewards/collected/by_user/{user_addr}/{offset}/{limit}", a.handleStakingCstRewardsCollectedByUser)
	r.GET("/api/cosmicgame/staking/cst/rewards/action_ids_by_deposit/{user_addr}/{deposit_id}", a.handleStakingCstRewardsActionIDsByDeposit)
	r.GET("/api/cosmicgame/staking/cst/rewards/by_user/by_token/summary/{user_addr}", a.handleStakingCstByUserByTokenRewards)
	r.GET("/api/cosmicgame/staking/cst/rewards/by_user/by_token/details/{user_addr}/{token_id}", a.handleStakingCstByUserByTokenRewardsDetails)
	r.GET("/api/cosmicgame/staking/cst/rewards/by_user/by_deposit/{user_addr}", a.handleStakingCstByUserByDepositRewards)
	r.GET("/api/cosmicgame/staking/cst/rewards/by_round/{round_num}", a.handleStakingCstRewardsByRound)
	r.GET("/api/cosmicgame/staking/cst/mints/global/{offset}/{limit}", a.handleStakingCstMintsGlobal)
	r.GET("/api/cosmicgame/staking/cst/mints/by_user/{user_addr}", a.handleStakingCstMintsByUser)

	// Staking RWalk (canonical: .../randomwalk/...; legacy: .../rwalk/... — same handlers)
	r.GET("/api/cosmicgame/staking/randomwalk/actions/info/{action_id}", a.handleStakingActionRwalkInfo)
	r.GET("/api/cosmicgame/staking/randomwalk/actions/global/{offset}/{limit}", a.handleStakingActionsRwalkGlobal)
	r.GET("/api/cosmicgame/staking/randomwalk/actions/by_user/{user_addr}/{offset}/{limit}", a.handleStakingActionsRwalkByUser)
	r.GET("/api/cosmicgame/staking/randomwalk/mints/global/{offset}/{limit}", a.handleStakingRwalkMintsGlobal)
	r.GET("/api/cosmicgame/staking/randomwalk/mints/by_user/{user_addr}", a.handleStakingRwalkMintsByUser)
	r.GET("/api/cosmicgame/staking/randomwalk/staked_tokens/all", a.handleStakedTokensRwalkGlobal)
	r.GET("/api/cosmicgame/staking/randomwalk/staked_tokens/by_user/{user_addr}", a.handleStakedTokensRwalkByUser)
	r.GET("/api/cosmicgame/staking/rwalk/actions/info/{action_id}", a.handleStakingActionRwalkInfo)
	r.GET("/api/cosmicgame/staking/rwalk/actions/global/{offset}/{limit}", a.handleStakingActionsRwalkGlobal)
	r.GET("/api/cosmicgame/staking/rwalk/actions/by_user/{user_addr}/{offset}/{limit}", a.handleStakingActionsRwalkByUser)
	r.GET("/api/cosmicgame/staking/rwalk/mints/global/{offset}/{limit}", a.handleStakingRwalkMintsGlobal)
	r.GET("/api/cosmicgame/staking/rwalk/mints/by_user/{user_addr}", a.handleStakingRwalkMintsByUser)
	r.GET("/api/cosmicgame/staking/rwalk/staked_tokens/all", a.handleStakedTokensRwalkGlobal)
	r.GET("/api/cosmicgame/staking/rwalk/staked_tokens/by_user/{user_addr}", a.handleStakedTokensRwalkByUser)

	// Marketing
	r.GET("/api/cosmicgame/marketing/rewards/global/{offset}/{limit}", a.handleMarketingRewardsGlobal)
	r.GET("/api/cosmicgame/marketing/rewards/by_user/{user_addr}/{offset}/{limit}", a.handleMarketingRewardsByUser)
	r.GET("/api/cosmicgame/marketing/config/current", a.handleMarketingConfigCurrent)

	// Time
	r.GET("/api/cosmicgame/time/current", a.handleTimeCurrent)
	r.GET("/api/cosmicgame/time/until_prize", a.handleTimeUntilPrize)

	// System
	r.GET("/api/cosmicgame/system/modelist", a.handleSysmodeChanges)
	r.GET("/api/cosmicgame/system/modelist/{offset}/{limit}", a.handleSysmodeChanges)
	r.GET("/api/cosmicgame/system/admin_events/{evtlog_start}/{evtlog_end}", a.handleAdminEventsInRange)
}
