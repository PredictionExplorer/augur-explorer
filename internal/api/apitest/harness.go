//go:build integration

package apitest

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log/slog"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
	v2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// adminKey is the shared secret configured for admin-guarded routes in the
// harness modules (the success branches of the auth matrices use it; the
// fail-closed 503 branch builds a keyless router).
const adminKey = "apitest-admin-key"

// harness is the fully wired API server under test: the real router with
// the real middleware stack, backed by the seeded test database and the
// deterministic in-memory Ethereum node (internal/testchain) serving
// fixture-coherent contract state.
type harness struct {
	router        *httpx.Router
	db            *sql.DB
	store         *store.Store
	state         *contractstate.State
	cg            *cosmicgame.API
	rw            *randomwalk.API
	ethClient     *ethclient.Client
	rpcClient     *ethrpc.Client
	chain         *testchain.Chain
	faqStubURL    string
	gameStub      *testchain.ContractStub
	tokenStub     *testchain.ContractStub
	marketingStub *testchain.ContractStub

	ipCounter atomic.Uint64
}

// seedDatabase applies the shared fixture dataset (internal/testfixtures).
func seedDatabase(ctx context.Context, db *sql.DB) error {
	return testfixtures.Apply(ctx, db)
}

// newHarness constructs the API modules exactly like cmd/apiserver/main.go
// (cosmicgame, randomwalk, faq) against the test database and fake chain,
// then assembles the production middleware chain and route table through the
// shared constructor (internal/api/routes). Modules are injected values, so
// harnesses and tests can build as many independent routers as they need.
func newHarness(ctx context.Context, db *testdb.DB) (*harness, error) {
	chain, _ := testchain.Start() // lives for the whole test process
	stubs := registerChainState(chain)
	faqSrv := newFAQStub()

	rpcClient, err := ethrpc.DialContext(ctx, chain.URL())
	if err != nil {
		return nil, fmt.Errorf("dialing test chain: %w", err)
	}
	ethClient := ethclient.NewClient(rpcClient)

	// One Store over the container's pool backs every query, exactly like
	// cmd/apiserver.
	st := store.NewFromPool(db.Pool)

	// StartBackgroundRefresh is deliberately not called: the refresh loops
	// would mutate the contract-state snapshot on a timer and make golden
	// snapshots drift. New's synchronous loads pin the state once. The
	// admin keys mirror a production deployment; the fail-closed matrices
	// build a keyless router next to this one.
	cgAPI, err := cosmicgame.New(ctx, cosmicgame.Config{
		Store:       st,
		EthClient:   ethClient,
		RPCClient:   rpcClient,
		AdminAPIKey: adminKey,
	})
	if err != nil {
		return nil, fmt.Errorf("initializing cosmicgame module: %w", err)
	}
	rwAPI := randomwalk.New(st, randomwalk.Options{
		AdminAPIKey:     adminKey,
		RankingAdminKey: adminKey,
	})
	faqProxy := faq.New(faq.Options{UpstreamURL: faqSrv.URL})

	v2Server, err := v2.NewServer(
		st,
		cgAPI.ContractState(),
		slog.New(slog.DiscardHandler),
		// Pin relative-time fields so HTTP goldens do not age.
		v2.WithClock(func() time.Time { return time.Unix(1767230000, 0) }),
		v2.WithAdmin(v2.AdminConfig{
			AdminKeys: []common.AdminKey{
				{Name: "ADMIN_API_KEY", Value: adminKey},
			},
		}),
		// Production ranking wiring: the same admin key as the v1 module,
		// a configured chain allowlist (so the chain-rejection problem is
		// reachable) and the default write rate limits (each harness
		// request uses a fresh client IP, so per-IP buckets only trip when
		// a test pins remoteAddr on purpose).
		v2.WithRanking(v2.RankingConfig{
			AdminKeys: []common.AdminKey{
				{Name: "RANKING_ADMIN_KEY", Value: adminKey},
				{Name: "ADMIN_API_KEY", Value: adminKey},
			},
			VoteChainIDs: []int64{1, 42161},
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("initializing api v2: %w", err)
	}

	// The Options build the production chain minus the stdout access log
	// (it would drown test output), metrics and static assets — the same
	// CORS, recovery and rate-limit middleware in the same order.
	r := routes.New(st, routes.Options{
		CosmicGame: cgAPI,
		RandomWalk: rwAPI,
		FAQ:        faqProxy,
		V2:         v2Server,
	})

	return &harness{
		router:        r,
		db:            db.SQL,
		store:         st,
		state:         cgAPI.ContractState(),
		cg:            cgAPI,
		rw:            rwAPI,
		ethClient:     ethClient,
		rpcClient:     rpcClient,
		chain:         chain,
		faqStubURL:    faqSrv.URL,
		gameStub:      stubs.game,
		tokenStub:     stubs.token,
		marketingStub: stubs.marketing,
	}, nil
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
	httpReq := httptest.NewRequestWithContext(context.Background(), req.method, req.path, req.body)
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
