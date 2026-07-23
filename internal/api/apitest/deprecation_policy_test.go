package apitest

import (
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/goccy/go-yaml"

	"github.com/PredictionExplorer/augur-explorer/internal/api/policy"
)

// TestDeprecationPolicyMatchesOpenAPI pins docs/openapi.yaml to the runtime
// deprecation policy in both directions: an operation carries
// `deprecated: true` exactly when policy.V1Deprecated matches its path, so
// the spec's deprecated flags and the Deprecation headers the router emits
// can never disagree. Together with TestRouteDriftAgainstOpenAPI (which
// pins spec ⇄ route table) this transitively pins the headers to the
// served routes.
func TestDeprecationPolicyMatchesOpenAPI(t *testing.T) {
	raw, err := os.ReadFile(openapiPath())
	if err != nil {
		t.Fatalf("reading OpenAPI spec: %v", err)
	}
	var doc struct {
		Paths map[string]map[string]struct {
			Deprecated bool `yaml:"deprecated"`
		} `yaml:"paths"`
	}
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("parsing OpenAPI spec: %v", err)
	}
	if len(doc.Paths) == 0 {
		t.Fatal("OpenAPI spec has no paths")
	}

	methods := map[string]bool{
		"get": true, "post": true, "put": true, "patch": true,
		"delete": true, "head": true, "options": true,
	}

	var deprecatedCount, activeCount int
	var problems []string
	for specPath, item := range doc.Paths {
		// Spec templates ({param}) never influence the prefix policy, so
		// paths compare directly.
		want := policy.V1Deprecated(specPath)
		for method, op := range item {
			if !methods[strings.ToLower(method)] {
				continue // parameters, summary, ...
			}
			if op.Deprecated {
				deprecatedCount++
			} else {
				activeCount++
			}
			if op.Deprecated != want {
				problems = append(problems, strings.ToUpper(method)+" "+specPath)
			}
		}
	}
	slices.Sort(problems)
	for _, p := range problems {
		t.Errorf("spec deprecated flag disagrees with policy.V1Deprecated: %s", p)
	}

	// The split itself is pinned so a future route lands on the right side
	// deliberately: 183 deprecated v1 operations; 8 exempt (healthz,
	// readyz, version, the two contract-pinned metadata routes and the
	// three FAQ proxy operations).
	if deprecatedCount != 183 || activeCount != 8 {
		t.Errorf("deprecated/active operations = %d/%d, want 183/8 — update this pin only with a deliberate policy change",
			deprecatedCount, activeCount)
	}
}
