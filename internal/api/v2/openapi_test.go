package v2

import (
	"context"
	"net/http"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

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
	slices.Sort(specRoutes)
	if len(specRoutes) != 102 {
		t.Fatalf("OpenAPI v2 operation count = %d, want 102", len(specRoutes))
	}

	router := httpx.NewRouter()
	newTestServer(t, fakeBidReader{}).RegisterRoutes(router)
	registered := make([]string, 0, len(router.Routes()))
	for _, route := range router.Routes() {
		registered = append(registered, route.Method+" "+route.Pattern)
	}
	slices.Sort(registered)

	if !reflect.DeepEqual(registered, specRoutes) {
		t.Fatalf("generated route drift\nregistered: %v\nspec:       %v", registered, specRoutes)
	}
	// The write surface is deliberately small (ADR-0008/0009): the three
	// ranking mutations plus bid-ban creation and deletion. Any other non-GET
	// registration is drift.
	writeOperations := map[string]bool{
		"POST /api/v2/cosmicgame/moderation/banned-bids":           true,
		"DELETE /api/v2/cosmicgame/moderation/banned-bids/{bidId}": true,
		"POST /api/v2/randomwalk/ranking/challenges":               true,
		"POST /api/v2/randomwalk/ranking/votes":                    true,
		"POST /api/v2/randomwalk/ranking/matches":                  true,
	}
	seenWrites := 0
	for _, route := range router.Routes() {
		if route.Method == http.MethodGet {
			continue
		}
		if !writeOperations[route.Method+" "+route.Pattern] {
			t.Errorf("unexpected non-GET v2 route: %+v", route)
			continue
		}
		seenWrites++
	}
	if seenWrites != len(writeOperations) {
		t.Errorf("registered %d of the %d documented write operations", seenWrites, len(writeOperations))
	}
}

// TestEveryOperationAnswersStableBindingProblems drives every documented
// operation's generated parameter-binding arms: duplicated and (for typed
// parameters) malformed query values must always produce the shared RFC 9457
// problem shape naming the parameter and must never echo the client's value.
// This pins that no operation can be registered without the custom binding
// error handler.
func TestEveryOperationAnswersStableBindingProblems(t *testing.T) {
	t.Parallel()

	spec, err := GetSpec()
	if err != nil {
		t.Fatalf("GetSpec: %v", err)
	}
	pathValues := map[string]string{
		"address":    userCursorAlice,
		"round":      "1",
		"position":   "1",
		"depositId":  "501",
		"actionId":   "1",
		"bidId":      "2001",
		"nftTokenId": "5",
		"tokenId":    "10",
	}

	for specPath, item := range spec.Paths.Map() {
		operation := item.Get
		if operation == nil {
			continue
		}
		target := specPath
		for name, value := range pathValues {
			target = strings.ReplaceAll(target, "{"+name+"}", value)
		}
		if strings.Contains(target, "{") {
			t.Fatalf("unmapped path parameter in %s; extend pathValues", specPath)
		}

		// Required sibling parameters must bind cleanly so the arm under
		// test is the one that fires. Date-formatted windows need date
		// values where the analytics windows use Unix seconds.
		requiredDefaults := map[string]string{
			"pool":      "bidder",
			"from":      "1767225600",
			"to":        "1767226600",
			"from:date": "2026-01-01",
			"to:date":   "2026-01-10",
		}
		requiredQuery := func(except string) string {
			parts := make([]string, 0, 1)
			for _, parameterRef := range operation.Parameters {
				parameter := parameterRef.Value
				if parameter == nil || parameter.In != openapi3.ParameterInQuery ||
					!parameter.Required || parameter.Name == except {
					continue
				}
				key := parameter.Name
				if schema := parameter.Schema; schema != nil && schema.Value != nil &&
					schema.Value.Format == "date" {
					key = parameter.Name + ":date"
				}
				value, known := requiredDefaults[key]
				if parameter.Name == "pool" &&
					strings.Contains(specPath, "/staking/raffle-nft-wins") {
					value, known = "cst", true
				}
				if !known {
					t.Fatalf("no default for required parameter %q of %s; extend requiredDefaults",
						parameter.Name, specPath)
				}
				parts = append(parts, parameter.Name+"="+value)
			}
			return strings.Join(parts, "&")
		}

		for _, parameterRef := range operation.Parameters {
			parameter := parameterRef.Value
			if parameter == nil || parameter.In != openapi3.ParameterInQuery {
				continue
			}
			queries := map[string]string{
				"duplicated": parameter.Name + "=password-super-secret&" + parameter.Name + "=b",
			}
			if schema := parameter.Schema; schema != nil && schema.Value != nil &&
				!schema.Value.Type.Is(openapi3.TypeString) {
				queries["malformed"] = parameter.Name + "=password-super-secret"
			}
			if parameter.Required {
				queries["missing"] = ""
			}
			siblings := requiredQuery(parameter.Name)
			for kind, query := range queries {
				if siblings != "" {
					if query != "" {
						query += "&"
					}
					query += siblings
				}
				name := kind + " " + parameter.Name + " on " + specPath
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					server := newTestServer(t, fakeBidReader{})
					response := serve(t, server, target+"?"+query)
					assertProblem(t, response, http.StatusBadRequest)

					var problem Problem
					decodeResponse(t, response, &problem)
					if problem.Type != problemTypeBase+"invalid-request" ||
						problem.Detail == nil ||
						!strings.Contains(*problem.Detail, `"`+parameter.Name+`"`) {
						t.Fatalf("problem = %+v", problem)
					}
					if strings.Contains(response.Body.String(), "password-super-secret") {
						t.Fatalf("parameter value leaked: %s", response.Body.String())
					}
				})
			}
		}
	}
}
