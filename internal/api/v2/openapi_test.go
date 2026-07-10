package v2

import (
	"context"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func TestOpenAPISpecValid(t *testing.T) {
	t.Parallel()

	spec, err := GetSpec()
	if err != nil {
		t.Fatalf("GetSpec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("OpenAPI v2 document is invalid: %v", err)
	}
	if spec.OpenAPI != "3.0.3" {
		t.Fatalf("OpenAPI version = %q, want 3.0.3", spec.OpenAPI)
	}
}

func TestV2RouteDriftAgainstOpenAPI(t *testing.T) {
	t.Parallel()

	spec, err := GetSpec()
	if err != nil {
		t.Fatalf("GetSpec: %v", err)
	}
	specRoutes := make([]string, 0)
	for path, item := range spec.Paths.Map() {
		for method := range item.Operations() {
			specRoutes = append(specRoutes, strings.ToUpper(method)+" "+path)
		}
	}
	sort.Strings(specRoutes)

	router := httpx.NewRouter()
	newTestServer(t, fakeBidReader{}).RegisterRoutes(router)
	registered := make([]string, 0)
	for _, route := range router.Routes() {
		registered = append(registered, route.Method+" "+route.Pattern)
	}
	sort.Strings(registered)

	if !reflect.DeepEqual(registered, specRoutes) {
		t.Fatalf("generated route drift\nregistered: %v\nspec:       %v", registered, specRoutes)
	}
	for _, route := range router.Routes() {
		if route.Method != http.MethodGet {
			t.Errorf("unexpected method in initial v2 slice: %+v", route)
		}
	}
}
