//go:build integration

package apitest

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	dbs "github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// harness is the fully wired API server under test: the real gin router with
// the real middleware stack, backed by the seeded test database and the
// deterministic Ethereum stub.
type harness struct {
	router  *gin.Engine
	db      *sql.DB
	storage *dbs.SQLStorage

	ipCounter atomic.Uint64
}

// seedDatabase applies the shared fixture dataset (internal/testfixtures).
func seedDatabase(ctx context.Context, db *sql.DB) error {
	return testfixtures.Apply(ctx, db)
}

// newHarness initializes the API modules exactly like cmd/apiserver/main.go
// (common context, cosmicgame, randomwalk, faq) against the test database and
// eth stub, then assembles the production middleware chain and route table.
// It must be called exactly once per process: the API packages hold their
// dependencies in package-level state until Phase 2 makes them injectable.
func newHarness(ctx context.Context, db *sql.DB) (*harness, error) {
	gin.SetMode(gin.TestMode)
	discard := log.New(io.Discard, "", 0)

	ethSrv := newEthStub() // lives for the whole test process
	faqSrv := newFAQStub()
	if err := os.Setenv("AI_BOT_BACKEND_URL", faqSrv.URL); err != nil {
		return nil, err
	}

	rpcClient, err := ethrpc.DialContext(ctx, ethSrv.URL)
	if err != nil {
		return nil, fmt.Errorf("dialing eth stub: %w", err)
	}
	ethClient := ethclient.NewClient(rpcClient)

	// Store-layer errors precede an os.Exit(1) in the legacy query methods
	// (Phase 1 removes those); keep them visible so a fixture/query problem
	// that kills the test binary is diagnosable from the output.
	storeLog := log.New(os.Stderr, "store: ", 0)
	storage := dbs.NewSQLStorageFromDB(db, storeLog)

	common.InitContext(storage, ethClient, discard, discard)

	// The reload_* refresh goroutines mutate package state on a timer, which
	// would race with request handling and make snapshots drift; the initial
	// synchronous loads still run inside Init.
	cosmicgame.DisableBackgroundRefresh()
	cosmicgame.Init(ethClient, rpcClient, discard, discard, true)
	randomwalk.Init(true)
	faq.Init(discard, discard, true)

	r := gin.New()
	if err := r.SetTrustedProxies(nil); err != nil {
		return nil, err
	}
	// CORS middleware replicated from cmd/apiserver/main.go.
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
	r.Use(gin.Recovery())
	// gin.Logger() is omitted: it only writes to stdout and would drown the
	// test output. The production rate limiter stays in place.
	r.Use(common.RateLimit(50, 100))

	registerAllRoutes(r, storage)

	return &harness{router: r, db: db, storage: storage}, nil
}

// request performs one HTTP exchange through the real router. Every call uses
// a fresh client IP so the per-IP rate limiter behaves as it would across
// distinct production clients.
type request struct {
	method  string
	path    string
	body    io.Reader
	headers map[string]string
	host    string // overrides the request Host (metadata dispatch)
}

func (h *harness) do(t *testing.T, req request) *httptest.ResponseRecorder {
	t.Helper()
	if req.method == "" {
		req.method = "GET"
	}
	httpReq := httptest.NewRequest(req.method, req.path, req.body)
	n := h.ipCounter.Add(1)
	httpReq.RemoteAddr = fmt.Sprintf("10.%d.%d.%d:4242", (n>>16)&0xff, (n>>8)&0xff, n&0xff)
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
