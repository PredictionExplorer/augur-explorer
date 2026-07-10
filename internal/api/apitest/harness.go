//go:build integration

package apitest

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http/httptest"
	"os"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
	"github.com/PredictionExplorer/augur-explorer/internal/api/v2"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// harness is the fully wired API server under test: the real router with
// the real middleware stack, backed by the seeded test database and the
// deterministic in-memory Ethereum node (internal/testchain) serving
// fixture-coherent contract state.
type harness struct {
	router *httpx.Router
	db     *sql.DB
	store  *store.Store
	state  *contractstate.State

	ipCounter atomic.Uint64
}

// seedDatabase applies the shared fixture dataset (internal/testfixtures).
func seedDatabase(ctx context.Context, db *sql.DB) error {
	return testfixtures.Apply(ctx, db)
}

// newHarness initializes the API modules exactly like cmd/apiserver/main.go
// (common context, cosmicgame, randomwalk, faq) against the test database and
// fake chain, then assembles the production middleware chain and route table
// through the shared constructor (internal/api/routes). It must be called
// exactly once per process: the API packages hold their dependencies in
// package-level state until Phase 2 makes them injectable.
func newHarness(ctx context.Context, db *testdb.DB) (*harness, error) {
	discard := log.New(io.Discard, "", 0)

	chain, _ := testchain.Start() // lives for the whole test process
	registerChainState(chain)
	faqSrv := newFAQStub()
	if err := os.Setenv("AI_BOT_BACKEND_URL", faqSrv.URL); err != nil {
		return nil, err
	}

	rpcClient, err := ethrpc.DialContext(ctx, chain.URL())
	if err != nil {
		return nil, fmt.Errorf("dialing test chain: %w", err)
	}
	ethClient := ethclient.NewClient(rpcClient)

	// One Store over the container's pool backs every query, exactly like
	// cmd/apiserver.
	st := store.NewFromPool(db.Pool)

	common.InitContext(st, ethClient, discard, discard)

	// StartBackgroundRefresh is deliberately not called: the refresh loops
	// would mutate the contract-state snapshot on a timer and make golden
	// snapshots drift. Init's synchronous loads pin the state once.
	cgState, err := cosmicgame.Init(ctx, ethClient, rpcClient, discard, discard, true)
	if err != nil {
		return nil, fmt.Errorf("initializing cosmicgame module: %w", err)
	}
	randomwalk.Init(true)
	faq.Init(discard, discard, true)

	v2Server, err := v2.NewServer(st, cgState, slog.New(slog.NewTextHandler(io.Discard, nil)))
	if err != nil {
		return nil, fmt.Errorf("initializing api v2: %w", err)
	}

	// The zero Options build the production chain minus the stdout access
	// log (it would drown test output), metrics and static assets — the
	// same CORS, recovery and rate-limit middleware in the same order.
	r := routes.New(st, routes.Options{V2: v2Server})

	return &harness{router: r, db: db.SQL, store: st, state: cgState}, nil
}

// request performs one HTTP exchange through the real router. Every call uses
// a fresh client IP so the per-IP rate limiter behaves as it would across
// distinct production clients; remoteAddr pins one (rate-limit tests).
type request struct {
	method     string
	path       string
	body       io.Reader
	headers    map[string]string
	host       string // overrides the request Host (metadata dispatch)
	remoteAddr string // overrides the per-request unique client IP
	ctx        context.Context
}

func (h *harness) do(t *testing.T, req request) *httptest.ResponseRecorder {
	t.Helper()
	if req.method == "" {
		req.method = "GET"
	}
	httpReq := httptest.NewRequest(req.method, req.path, req.body)
	n := h.ipCounter.Add(1)
	httpReq.RemoteAddr = fmt.Sprintf("10.%d.%d.%d:4242", (n>>16)&0xff, (n>>8)&0xff, n&0xff)
	if req.remoteAddr != "" {
		httpReq.RemoteAddr = req.remoteAddr
	}
	if req.host != "" {
		httpReq.Host = req.host
	}
	for k, v := range req.headers {
		httpReq.Header.Set(k, v)
	}
	if req.body != nil && httpReq.Header.Get("Content-Type") == "" {
		httpReq.Header.Set("Content-Type", "application/json")
	}
	w := httptest.NewRecorder()
	if req.ctx != nil {
		httpReq = httpReq.WithContext(req.ctx)
	}
	h.router.ServeHTTP(w, httpReq)
	return w
}

// get performs a GET and returns the recorder.
func (h *harness) get(t *testing.T, path string) *httptest.ResponseRecorder {
	t.Helper()
	return h.do(t, request{path: path})
}

// contentTypeOf normalizes a Content-Type header for golden storage.
func contentTypeOf(w *httptest.ResponseRecorder) string {
	ct := w.Header().Get("Content-Type")
	if i := strings.IndexByte(ct, ';'); i >= 0 {
		ct = ct[:i]
	}
	return strings.TrimSpace(ct)
}
