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
	return newServer(st, repo, repo, repo, state, logger)
}

func newServer(
	st *store.Store,
	bids bidReader,
	rounds roundReader,
	currentRounds currentRoundReader,
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
