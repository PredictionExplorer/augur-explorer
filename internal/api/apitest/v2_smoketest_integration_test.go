//go:build integration

package apitest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/smoketest"
)

type smokeRouterClient struct {
	harness *harness
	nextIP  atomic.Uint64
}

func (c *smokeRouterClient) Do(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.RemoteAddr = fmt.Sprintf("10.240.%d.%d:4242",
		(c.nextIP.Load()/250)%250,
		c.nextIP.Add(1)%250+1,
	)
	recorder := httptest.NewRecorder()
	c.harness.router.ServeHTTP(recorder, clone)
	return recorder.Result(), nil
}

// TestV2SmoketestCatalogThroughProductionRouter is the out-of-process
// operational contract brought into the seeded integration harness: every
// safe GET generated from OpenAPI must answer 200, validate against that same
// contract, and remain outside the v1 deprecation surface.
func TestV2SmoketestCatalogThroughProductionRouter(t *testing.T) {
	h := server(t)
	summary, err := smoketest.Run(t.Context(), smoketest.Options{
		Source:        smoketest.SQLParameterSource{DB: h.store.Pool()},
		Client:        &smokeRouterClient{harness: h},
		BaseURL:       "http://apitest.local",
		Output:        io.Discard,
		Suite:         smoketest.SuiteV2,
		DisablePacing: true,
	})
	if err != nil {
		t.Fatalf("v2 smoke run failed: %v\nsummary: %#v", err, summary)
	}
	if summary.Total != 101 || summary.OK != 101 || len(summary.Failures) != 0 {
		t.Fatalf("summary = %#v", summary)
	}
}

func TestV1SmoketestCatalogThroughProductionRouter(t *testing.T) {
	h := server(t)
	summary, err := smoketest.Run(t.Context(), smoketest.Options{
		Source:        smoketest.SQLParameterSource{DB: h.store.Pool()},
		Client:        &smokeRouterClient{harness: h},
		BaseURL:       "http://apitest.local",
		Output:        io.Discard,
		Suite:         smoketest.SuiteV1,
		DisablePacing: true,
	})
	if err != nil {
		t.Fatalf("v1 smoke run failed: %v\nsummary: %#v", err, summary)
	}
	if summary.Total != 145 || summary.OK != 145 || len(summary.Failures) != 0 {
		t.Fatalf("summary = %#v", summary)
	}
}
