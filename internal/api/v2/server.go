package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const problemTypeBase = "https://api.cosmicsignature.com/problems/"

type bidReader interface {
	BidsByRoundPage(context.Context, int64, cgstore.BidPageCursor, int) ([]cgprimitives.CGBidRec, bool, error)
	BidByRoundAndPosition(context.Context, int64, int64) (cgprimitives.CGBidRec, error)
}

type roundReader interface {
	PrizeClaimsPage(context.Context, *cgstore.RoundPageCursor, int) ([]cgprimitives.CGRoundRec, bool, error)
	RoundInfo(context.Context, int64) (cgprimitives.CGRoundRec, error)
}

type currentRoundReader interface {
	CosmicGameRoundStatistics(context.Context, int64) (cgprimitives.CGRoundStats, error)
	BidCountForRound(context.Context, int64) (int64, error)
}

type roundPrizeReader interface {
	CompletedRoundExists(context.Context, int64) (bool, error)
	AllPrizesForRoundPage(context.Context, int64, *cgstore.PrizePageCursor, int) ([]cgprimitives.CGPrizeHistory, bool, error)
}

type roundRaffleReader interface {
	CompletedRoundExists(context.Context, int64) (bool, error)
	RaffleEthDepositsByRoundPage(context.Context, int64, *cgstore.RaffleEthDepositPageCursor, int) ([]cgstore.RaffleEthDepositRecord, bool, error)
	RaffleNFTWinnersByRoundPage(context.Context, int64, bool, *cgstore.RaffleNFTWinnerPageCursor, int) ([]cgprimitives.CGRaffleNFTWinnerRec, bool, error)
}

type roundDonationReader interface {
	EthDonationsByRoundPage(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error)
	ERC20DonationsByRoundPage(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundERC20DonationRecord, bool, error)
	NFTDonationsByRoundPage(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundNFTDonationRecord, bool, error)
}

type statisticsReader interface {
	CosmicGameGlobalStatistics(context.Context) (cgstore.GlobalStatisticsRecord, error)
	RecordCounters(context.Context) (cgprimitives.CGRecordCounters, error)
	ROILeaderboardPage(context.Context, int, cgstore.ROILeaderboardSort, *cgstore.ROILeaderboardPageCursor, int) ([]cgstore.ROILeaderboardRecord, bool, error)
	ClaimsSummaryPage(context.Context, *cgstore.ClaimSummaryCursor, int) ([]cgstore.ClaimSummaryRecord, bool, error)
	CompletedRoundExists(context.Context, int64) (bool, error)
	ClaimSummaryByRound(context.Context, int64) (cgstore.ClaimSummaryRecord, error)
	ClaimTransactionsPage(context.Context, int64, *cgstore.ClaimEventCursor, int) ([]cgstore.ClaimTransactionRecord, bool, error)
	AttachedTokensPage(context.Context, int64, *cgstore.ClaimEventCursor, int) ([]cgstore.AttachedTokenRecord, bool, error)
	UnclaimedItemsPage(context.Context, int64, *cgstore.UnclaimedItemCursor, int) ([]cgstore.UnclaimedItemRecord, bool, error)
}

type participantReader interface {
	BidderParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.BidderParticipantRecord, bool, error)
	WinnerParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.WinnerParticipantRecord, bool, error)
	DonorParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.DonorParticipantRecord, bool, error)
	CSTStakerParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CSTStakerParticipantRecord, bool, error)
	RandomWalkStakerParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.RandomWalkStakerParticipantRecord, bool, error)
	DualStakerParticipantsPage(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.DualStakerParticipantRecord, bool, error)
}

type contractStateReader interface {
	Snapshot() contractstate.Snapshot
}

// Server implements the generated v2 strict-server contract. Every runtime
// dependency is injected once at construction; handlers do not read package
// globals.
type Server struct {
	store         *store.Store
	bids          bidReader
	rounds        roundReader
	currentRounds currentRoundReader
	prizes        roundPrizeReader
	raffles       roundRaffleReader
	donations     roundDonationReader
	statistics    statisticsReader
	participants  participantReader
	contractState contractStateReader
	logger        *slog.Logger
}

var _ StrictServerInterface = (*Server)(nil)

// NewServer constructs the production v2 server over the shared store and
// contract-state cache.
func NewServer(st *store.Store, state *contractstate.State, logger *slog.Logger) (*Server, error) {
	if st == nil {
		return nil, errors.New("api v2: store is required")
	}
	if state == nil {
		return nil, errors.New("api v2: contract state is required")
	}
	repo := cgstore.NewRepo(st)
	return newServer(st, repo, repo, repo, repo, repo, repo, repo, repo, state, logger)
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
	participants participantReader,
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
	if participants == nil {
		return nil, errors.New("api v2: participant repository is required")
	}
	if state == nil {
		return nil, errors.New("api v2: contract state is required")
	}
	if logger == nil {
		logger = slog.Default()
	}
	return &Server{
		store:         st,
		bids:          bids,
		rounds:        rounds,
		currentRounds: currentRounds,
		prizes:        prizes,
		raffles:       raffles,
		donations:     donations,
		statistics:    statistics,
		participants:  participants,
		contractState: state,
		logger:        logger,
	}, nil
}

// RegisterRoutes installs every generated v2 operation on the shared router.
// Custom error hooks ensure generated parameter-binding and response failures
// use the same RFC 9457 representation as handler-level errors.
func (s *Server) RegisterRoutes(r *httpx.Router) {
	strict := NewStrictHandlerWithOptions(s, nil, StrictHTTPServerOptions{
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
