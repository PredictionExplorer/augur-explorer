package v2

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

const problemTypeBase = "https://api.cosmicsignature.com/problems/"

type bidReader interface {
	BidsByRoundPage(context.Context, int64, cgstore.BidPageCursor, int) ([]cgmodel.CGBidRec, bool, error)
	BidByRoundAndPosition(context.Context, int64, int64) (cgmodel.CGBidRec, error)
}

type roundReader interface {
	PrizeClaimsPage(context.Context, *cgstore.RoundPageCursor, int) ([]cgmodel.CGRoundRec, bool, error)
	RoundInfo(context.Context, int64) (cgmodel.CGRoundRec, error)
}

type currentRoundReader interface {
	CosmicGameRoundStatistics(context.Context, int64) (cgmodel.CGRoundStats, error)
	BidCountForRound(context.Context, int64) (int64, error)
}

type roundPrizeReader interface {
	CompletedRoundExists(context.Context, int64) (bool, error)
	AllPrizesForRoundPage(context.Context, int64, *cgstore.PrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error)
}

type roundRaffleReader interface {
	CompletedRoundExists(context.Context, int64) (bool, error)
	RaffleEthDepositsByRoundPage(context.Context, int64, *cgstore.RaffleEthDepositPageCursor, int) ([]cgstore.RaffleEthDepositRecord, bool, error)
	RaffleNFTWinnersByRoundPage(context.Context, int64, bool, *cgstore.RaffleNFTWinnerPageCursor, int) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error)
}

type roundDonationReader interface {
	EthDonationsByRoundPage(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error)
	ERC20DonationsByRoundPage(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundERC20DonationRecord, bool, error)
	NFTDonationsByRoundPage(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundNFTDonationRecord, bool, error)
}

type statisticsReader interface {
	CosmicGameGlobalStatistics(context.Context) (cgstore.GlobalStatisticsRecord, error)
	RecordCounters(context.Context) (cgmodel.CGRecordCounters, error)
	ROILeaderboardPage(context.Context, int, cgstore.ROILeaderboardSort, *cgstore.ROILeaderboardPageCursor, int) ([]cgstore.ROILeaderboardRecord, bool, error)
	ClaimsSummaryPage(context.Context, *cgstore.ClaimSummaryCursor, int) ([]cgstore.ClaimSummaryRecord, bool, error)
	CompletedRoundExists(context.Context, int64) (bool, error)
	ClaimSummaryByRound(context.Context, int64) (cgstore.ClaimSummaryRecord, error)
	ClaimTransactionsPage(context.Context, int64, *cgstore.ClaimEventCursor, int) ([]cgstore.ClaimTransactionRecord, bool, error)
	AttachedTokensPage(context.Context, int64, *cgstore.ClaimEventCursor, int) ([]cgstore.AttachedTokenRecord, bool, error)
	UnclaimedItemsPage(context.Context, int64, *cgstore.UnclaimedItemCursor, int) ([]cgstore.UnclaimedItemRecord, bool, error)
}

type biddingAnalyticsReader interface {
	BidFrequencyByPeriodBounded(context.Context, int, int, int) ([]cgmodel.CGBidFrequencyBucket, error)
	BidTypeRatioByPeriodBounded(context.Context, int, int, int) ([]cgmodel.CGBidTypeRatioBucket, error)
	TopBidderActivePeriodsBounded(context.Context, int, int, int, int, int) ([]cgmodel.CGTopBidderInfo, []cgmodel.CGBidderActivePeriod, bool, error)
	BidTimeBounds(context.Context) (int64, int64, error)
}

type contractAddressReader interface {
	ContractAddrs(context.Context) (cgmodel.CosmicGameContractAddrs, error)
}

type participantReader interface {
	BidderParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.BidderParticipantRecord, bool, error)
	WinnerParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.WinnerParticipantRecord, bool, error)
	DonorParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.DonorParticipantRecord, bool, error)
	CSTStakerParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CSTStakerParticipantRecord, bool, error)
	RandomWalkStakerParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.RandomWalkStakerParticipantRecord, bool, error)
	DualStakerParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.DualStakerParticipantRecord, bool, error)
}

type userReader interface {
	UserAddressID(context.Context, string) (int64, error)
	UserProfile(context.Context, int64) (cgstore.UserProfileRecord, error)
	BidsByUserPage(context.Context, int64, *cgstore.UserBidPageCursor, int) ([]cgmodel.CGBidRec, bool, error)
}

type userHistoryReader interface {
	UserAddressID(context.Context, string) (int64, error)
	UserPrizesPage(context.Context, int64, *cgstore.UserPrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error)
	UserRaffleEthDepositsPage(context.Context, int64, *bool, *cgstore.UserEventPageCursor, int) ([]cgstore.UserRaffleEthDepositRecord, bool, error)
	UserRaffleNftWinsPage(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserRaffleNftWinRecord, bool, error)
	EthDonationsByUserPage(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error)
	ERC20DonationsByUserPage(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.RoundERC20DonationRecord, bool, error)
	NFTDonationsByUserPage(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.RoundNFTDonationRecord, bool, error)
	UserDonatedNftsPage(context.Context, int64, *bool, *cgstore.UserEventPageCursor, int) ([]cgstore.UserDonatedNftRecord, bool, error)
	UserDonatedErc20Page(context.Context, int64, *cgstore.UserDonatedErc20PageCursor, int) ([]cgstore.UserDonatedErc20Record, bool, error)
}

type userStakingReader interface {
	UserAddressID(context.Context, string) (int64, error)
	UserCstStakingActionsPage(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserStakingActionRecord, bool, error)
	UserRwalkStakingActionsPage(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserStakingActionRecord, bool, error)
	UserStakedCstTokensPage(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedCstTokenRecord, bool, error)
	UserStakedRwalkTokensPage(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedRwalkTokenRecord, bool, error)
	UserStakingDepositsPage(context.Context, int64, *bool, *cgstore.UserStakingDepositPageCursor, int) ([]cgstore.UserStakingDepositRecord, bool, error)
	StakingDepositExists(context.Context, int64) (bool, error)
	UserStakingDepositRewardsPage(context.Context, int64, int64, *cgstore.UserStakingRewardPageCursor, int) ([]cgstore.UserStakingDepositRewardRecord, bool, error)
	UserStakingTokenRewardsPage(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakingTokenRewardRecord, bool, error)
	CosmicSignatureTokenExists(context.Context, int64) (bool, error)
	UserStakingTokenRewardDepositsPage(context.Context, int64, int64, *cgstore.UserStakingTokenDepositPageCursor, int) ([]cgstore.UserStakingTokenRewardDepositRecord, bool, error)
}

type userActivityReader interface {
	UserAddressID(context.Context, string) (int64, error)
	UserCosmicSignatureTokensPage(context.Context, int64, *cgstore.UserTokenPageCursor, int) ([]cgstore.UserOwnedTokenRecord, bool, error)
	UserCosmicSignatureTransfersPage(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserCosmicSignatureTransferRecord, bool, error)
	UserCosmicTokenTransfersPage(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserCosmicTokenTransferRecord, bool, error)
	UserMarketingRewardsPage(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserMarketingRewardRecord, bool, error)
	UserCosmicTokenSummaryV2(context.Context, int64) (cgstore.UserCosmicTokenSummaryRecord, error)
	UserPendingWinnings(context.Context, int64) (cgstore.UserPendingWinningsRecord, error)
}

type globalDirectoryReader interface {
	CosmicSignatureTokensGlobalPage(context.Context, cgstore.GlobalTokenFilter, *cgstore.GlobalTokenPageCursor, int) ([]cgstore.GlobalTokenRecord, bool, error)
	CosmicSignatureTokenDetailV2(context.Context, int64) (cgstore.GlobalTokenDetailRecord, error)
	CosmicSignatureTokenExists(context.Context, int64) (bool, error)
	TokenNameHistoryPage(context.Context, int64, *cgstore.TokenEventPageCursor, int) ([]cgstore.TokenNameChangeRecord, bool, error)
	TokenTransfersPage(context.Context, int64, *cgstore.TokenEventPageCursor, int) ([]cgstore.TokenTransferRecord, bool, error)
	CosmicSignatureHoldersPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CosmicSignatureHolderRecord, bool, error)
	CosmicTokenHoldersPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CosmicTokenHolderRecord, bool, error)
	CosmicTokenStatisticsV2(context.Context) (cgstore.CosmicTokenStatisticsRecord, error)
	CosmicTokenSupplyByBidPage(context.Context, *cgstore.SupplyChangePageCursor, int) ([]cgstore.SupplyChangeRecord, bool, error)
	CosmicTokenSupplyDaily(context.Context, time.Time, time.Time) ([]cgstore.DailySupplyRecord, error)
	MarketingRewardsGlobalPage(context.Context, *cgstore.UserEventPageCursor, int) ([]cgstore.MarketingRewardRecord, bool, error)
}

type globalStakingReader interface {
	GlobalCstStakingActionsPage(context.Context, *cgstore.GlobalStakingActionPageCursor, int) ([]cgstore.GlobalStakingActionRecord, bool, error)
	GlobalRwalkStakingActionsPage(context.Context, *cgstore.GlobalStakingActionPageCursor, int) ([]cgstore.GlobalStakingActionRecord, bool, error)
	StakeActionCstInfo(context.Context, int64) (cgmodel.CGStakeUnstakeCombined, error)
	StakeActionRwalkInfo(context.Context, int64) (cgmodel.CGStakeUnstakeCombined, error)
	GlobalStakedCstTokensPage(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedCstTokenRecord, bool, error)
	GlobalStakedRwalkTokensPage(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedRwalkTokenRecord, bool, error)
	GlobalStakingDepositsPage(context.Context, *cgstore.GlobalStakingDepositPageCursor, int) ([]cgstore.GlobalStakingDepositRecord, bool, error)
	CompletedRoundExists(context.Context, int64) (bool, error)
	RoundStakingRewardsPage(context.Context, int64, *cgstore.RoundStakingRewardPageCursor, int) ([]cgstore.RoundStakingRewardRecord, bool, error)
	GlobalStakerRaffleNftWinsPage(context.Context, bool, *cgstore.GlobalStakerRafflePageCursor, int) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error)
}

type randomWalkReader interface {
	TokensPage(context.Context, rwstore.TokenFilter, rwstore.TokenSort, *rwstore.TokenPageCursor, int) ([]rwstore.TokenRecord, bool, error)
	TokenDetailV2(context.Context, int64) (rwstore.TokenDetailRecord, error)
	CollectionTokenExists(context.Context, int64) (bool, error)
	TokenNameChangesPageV2(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.TokenNameChangeRecord, bool, error)
	TokenEventsPage(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.TokenEventRecord, bool, error)
	ActiveOffersPage(context.Context, rwstore.OfferSort, *rwstore.OfferPageCursor, int) ([]rwstore.OfferRecord, bool, error)
	OfferHistoryPage(context.Context, *rwstore.EventPageCursor, int) ([]rwstore.OfferHistoryRecord, bool, error)
	UserOffersPage(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.OfferHistoryRecord, bool, error)
	TradesPage(context.Context, *rwstore.EventPageCursor, int) ([]rwstore.TradeRecord, bool, error)
	FloorPriceV2(context.Context) (rwstore.FloorPriceRecord, error)
	UserAddressID(context.Context, string) (int64, error)
	UserProfileV2(context.Context, int64) (rwstore.UserProfileRecord, error)
	UserTokensPage(context.Context, int64, *rwstore.TokenPageCursor, int) ([]rwstore.OwnedTokenRecord, bool, error)
	StatisticsV2(context.Context) (rwstore.StatisticsRecord, error)
	TradingVolumeSeries(context.Context, int, int, int) (string, []rwstore.VolumeBucketRecord, error)
	ListingFloorSeries(context.Context, int, int, int) ([]rwstore.FloorPointRecord, error)
	MintReportV2(context.Context) ([]rwstore.MonthlyMintRecord, error)
	WithdrawalsPage(context.Context, *rwstore.EventPageCursor, int) ([]rwstore.WithdrawalRecord, bool, error)
	ContractAddrs(context.Context) (rwmodel.ContractAddresses, error)
}

type rankingRepository interface {
	ContractAddrs(context.Context) (rwmodel.ContractAddresses, error)
	ExploreRandomTokenIDs(context.Context, int64, int64, int) ([]int64, error)
	UserAddressID(context.Context, string) (int64, error)
	HasRankingVoteForVoterPair(context.Context, int64, int64, int64) (bool, error)
	RankingRatingsPage(context.Context, int64, *rwstore.RankingRatingPageCursor, int) ([]rwstore.RankingRatingRecord, bool, error)
	RankingStatistics(context.Context) (rwstore.RankingStatisticsRecord, error)
	CountRankingMatches(context.Context) (int64, error)
	RatingPair(context.Context, int64, int64) (float64, float64, error)
	EnsureVoterAddress(context.Context, string) (int64, error)
	CreateRankingVoteNonce(context.Context, string, time.Duration) (time.Time, error)
	RecordRankingMatch(context.Context, int64, int64, bool, float64, float64) error
	RecordSignedRankingVote(context.Context, string, int64, int64, bool, float64, float64, int64) error
}

type contractStateReader interface {
	Snapshot() contractstate.Snapshot
}

// Server implements the generated v2 strict-server contract. Every runtime
// dependency is injected once at construction; handlers do not read package
// globals.
type Server struct {
	store             *store.Store
	bids              bidReader
	rounds            roundReader
	currentRounds     currentRoundReader
	prizes            roundPrizeReader
	raffles           roundRaffleReader
	donations         roundDonationReader
	statistics        statisticsReader
	analytics         biddingAnalyticsReader
	contractAddresses contractAddressReader
	participants      participantReader
	users             userReader
	userHistories     userHistoryReader
	userStaking       userStakingReader
	userActivity      userActivityReader
	globalDirectories globalDirectoryReader
	globalStaking     globalStakingReader
	randomWalk        randomWalkReader
	ranking           rankingRepository
	contractState     contractStateReader
	logger            *slog.Logger
	now               func() time.Time
	entropy           io.Reader
	rankingConfig     RankingConfig
}

// ServerOption customizes a Server at construction.
type ServerOption func(*Server)

// WithClock replaces the server clock. It is primarily useful for
// deterministic tests of time-relative response fields.
func WithClock(now func() time.Time) ServerOption {
	return func(server *Server) {
		server.now = now
	}
}

// WithEntropy replaces the randomness source behind challenge nonces. It is
// primarily useful for deterministic tests and entropy-failure tests.
func WithEntropy(entropy io.Reader) ServerOption {
	return func(server *Server) {
		server.entropy = entropy
	}
}

// WithRanking installs the ranking slice's deployment configuration (admin
// keys, the vote chain allowlist, the exploration bound and optional write
// rate-limit overrides). Servers built without it fail closed on the admin
// operation and allow any positive vote chain id.
func WithRanking(config RankingConfig) ServerOption {
	return func(server *Server) {
		server.rankingConfig = config
	}
}

var _ StrictServerInterface = (*Server)(nil)

// NewServer constructs the production v2 server over the shared store and
// contract-state cache.
func NewServer(
	st *store.Store,
	state *contractstate.State,
	logger *slog.Logger,
	options ...ServerOption,
) (*Server, error) {
	if st == nil {
		return nil, errors.New("api v2: store is required")
	}
	if state == nil {
		return nil, errors.New("api v2: contract state is required")
	}
	repo := cgstore.NewRepo(st)
	rwRepo := rwstore.NewRepo(st)
	server, err := newServer(st, repo, repo, repo, repo, repo, repo, repo, repo, repo, repo, repo, repo, repo, repo, repo, repo, rwRepo, rwRepo, state, logger)
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		if option == nil {
			return nil, errors.New("api v2: server option is nil")
		}
		option(server)
	}
	if server.now == nil {
		return nil, errors.New("api v2: clock is required")
	}
	if server.entropy == nil {
		return nil, errors.New("api v2: entropy source is required")
	}
	return server, nil
}

func newServer(
	st *store.Store,
	bids bidReader,
	rounds roundReader,
	currentRounds currentRoundReader,
	prizes roundPrizeReader,
	raffles roundRaffleReader,
	donations roundDonationReader,
	statistics statisticsReader,
	analytics biddingAnalyticsReader,
	contractAddresses contractAddressReader,
	participants participantReader,
	users userReader,
	userHistories userHistoryReader,
	userStaking userStakingReader,
	userActivity userActivityReader,
	globalDirectories globalDirectoryReader,
	globalStaking globalStakingReader,
	randomWalk randomWalkReader,
	ranking rankingRepository,
	state contractStateReader,
	logger *slog.Logger,
) (*Server, error) {
	if bids == nil {
		return nil, errors.New("api v2: bid repository is required")
	}
	if rounds == nil {
		return nil, errors.New("api v2: round repository is required")
	}
	if currentRounds == nil {
		return nil, errors.New("api v2: current-round repository is required")
	}
	if prizes == nil {
		return nil, errors.New("api v2: round-prize repository is required")
	}
	if raffles == nil {
		return nil, errors.New("api v2: round-raffle repository is required")
	}
	if donations == nil {
		return nil, errors.New("api v2: round-donation repository is required")
	}
	if statistics == nil {
		return nil, errors.New("api v2: statistics repository is required")
	}
	if analytics == nil {
		return nil, errors.New("api v2: bidding analytics repository is required")
	}
	if contractAddresses == nil {
		return nil, errors.New("api v2: contract-address repository is required")
	}
	if participants == nil {
		return nil, errors.New("api v2: participant repository is required")
	}
	if users == nil {
		return nil, errors.New("api v2: user repository is required")
	}
	if userHistories == nil {
		return nil, errors.New("api v2: user-history repository is required")
	}
	if userStaking == nil {
		return nil, errors.New("api v2: user-staking repository is required")
	}
	if userActivity == nil {
		return nil, errors.New("api v2: user-activity repository is required")
	}
	if globalDirectories == nil {
		return nil, errors.New("api v2: global-directory repository is required")
	}
	if globalStaking == nil {
		return nil, errors.New("api v2: global-staking repository is required")
	}
	if randomWalk == nil {
		return nil, errors.New("api v2: random-walk repository is required")
	}
	if ranking == nil {
		return nil, errors.New("api v2: ranking repository is required")
	}
	if state == nil {
		return nil, errors.New("api v2: contract state is required")
	}
	if logger == nil {
		logger = slog.Default()
	}
	return &Server{
		store:             st,
		bids:              bids,
		rounds:            rounds,
		currentRounds:     currentRounds,
		prizes:            prizes,
		raffles:           raffles,
		donations:         donations,
		statistics:        statistics,
		analytics:         analytics,
		contractAddresses: contractAddresses,
		participants:      participants,
		users:             users,
		userHistories:     userHistories,
		userStaking:       userStaking,
		userActivity:      userActivity,
		globalDirectories: globalDirectories,
		globalStaking:     globalStaking,
		randomWalk:        randomWalk,
		ranking:           ranking,
		contractState:     state,
		logger:            logger,
		now:               time.Now,
		entropy:           rand.Reader,
	}, nil
}

// RegisterRoutes installs every generated v2 operation on the shared router.
// Custom error hooks ensure generated parameter-binding and response failures
// use the same RFC 9457 representation as handler-level errors; the strict
// middleware chain enforces the per-operation write rate limits and the
// spec-declared admin authentication before any handler runs.
func (s *Server) RegisterRoutes(r *httpx.Router) {
	middlewares := []StrictMiddlewareFunc{
		s.rankingWriteRateLimitMiddleware(),
		s.adminKeyMiddleware(),
	}
	strict := NewStrictHandlerWithOptions(s, middlewares, StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, req *http.Request, err error) {
			s.writeRequestError(w, req, err)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, req *http.Request, err error) {
			s.logger.ErrorContext(req.Context(), "api v2 response failure",
				"method", req.Method,
				"path", req.URL.Path,
				"error", err)
			s.writeProblem(w, internalProblem(req.URL.Path))
		},
	})
	_ = HandlerWithOptions(strict, StdHTTPServerOptions{
		BaseRouter: r,
		ErrorHandlerFunc: func(w http.ResponseWriter, req *http.Request, err error) {
			s.writeRequestError(w, req, err)
		},
	})
}

func (s *Server) writeRequestError(w http.ResponseWriter, req *http.Request, err error) {
	detail := "A path or query parameter has an invalid value."
	var invalid *InvalidParamFormatError
	var required *RequiredParamError
	var tooMany *TooManyValuesForParamError
	switch {
	case errors.As(err, &invalid):
		detail = fmt.Sprintf("Parameter %q has an invalid value.", invalid.ParamName)
	case errors.As(err, &required):
		detail = fmt.Sprintf("Parameter %q is required.", required.ParamName)
	case errors.As(err, &tooMany):
		detail = fmt.Sprintf("Parameter %q must be provided once.", tooMany.ParamName)
	case strings.Contains(err.Error(), "decode JSON body"):
		// The generated strict wrapper wraps request-body decode failures
		// with this text; the client's malformed payload is never echoed.
		detail = "The request body is not valid JSON matching the documented schema."
	}
	s.writeProblem(w, newProblem(
		http.StatusBadRequest,
		"invalid-request",
		"Invalid request",
		detail,
		req.URL.Path,
	))
}

func (s *Server) writeProblem(w http.ResponseWriter, problem Problem) {
	body, err := json.Marshal(problem)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(problem.Status)
	_, _ = w.Write(append(body, '\n'))
}

func newProblem(status int, kind, title, detail, instance string) Problem {
	problemType := problemTypeBase + kind
	return Problem{
		Type:     problemType,
		Title:    title,
		Status:   status,
		Detail:   &detail,
		Instance: &instance,
	}
}

func internalProblem(instance string) Problem {
	return newProblem(
		http.StatusInternalServerError,
		"internal",
		"Internal server error",
		"The request could not be completed.",
		instance,
	)
}

func roundNotFoundProblem(instance string) Problem {
	return newProblem(
		http.StatusNotFound,
		"round-not-found",
		"Round not found",
		"No completed round exists with that number.",
		instance,
	)
}
