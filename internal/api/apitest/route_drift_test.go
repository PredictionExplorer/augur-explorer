package apitest

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"

	"github.com/goccy/go-yaml"
)

// TestRouteDriftAgainstOpenAPI pins the route table to docs/openapi.yaml in
// both directions: every documented operation must be registered on the
// router, and every registered route must be documented. The spec is the v1
// contract for the rewrite — this test makes silent drift impossible.
//
// OpenAPI path templates ("/a/{b}") and ServeMux patterns share one syntax,
// so operations compare directly against the route registry.
func TestRouteDriftAgainstOpenAPI(t *testing.T) {
	specOps := loadSpecOperations(t)

	r := buildBareRouter()
	registered := make(map[string]bool)
	for _, route := range r.Routes() {
		registered[route.Method+" "+route.Pattern] = true
	}

	var missingFromRouter []string
	for op := range specOps {
		if !registered[op] {
			missingFromRouter = append(missingFromRouter, op)
		}
	}
	slices.Sort(missingFromRouter)
	for _, op := range missingFromRouter {
		t.Errorf("documented in openapi.yaml but not registered on the router: %s", op)
	}

	var missingFromSpec []string
	for op := range registered {
		if !specOps[op] {
			missingFromSpec = append(missingFromSpec, op)
		}
	}
	slices.Sort(missingFromSpec)
	for _, op := range missingFromSpec {
		t.Errorf("registered on the router but not documented in openapi.yaml: %s", op)
	}
}

// loadSpecOperations parses docs/openapi.yaml and returns the set of
// "METHOD /path/{param}" operation keys it documents.
func loadSpecOperations(t *testing.T) map[string]bool {
	t.Helper()

	raw, err := os.ReadFile(openapiPath())
	if err != nil {
		t.Fatalf("reading OpenAPI spec: %v", err)
	}

	var doc struct {
		Paths map[string]map[string]any `yaml:"paths"`
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

	ops := make(map[string]bool)
	for specPath, item := range doc.Paths {
		for key := range item {
			lk := strings.ToLower(key)
			if !methods[lk] {
				continue // parameters, summary, etc.
			}
			ops[fmt.Sprintf("%s %s", strings.ToUpper(lk), specPath)] = true
		}
	}
	return ops
}

// openapiPath locates docs/openapi.yaml relative to this source file.
func openapiPath() string {
	_, thisFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(thisFile), "..", "..", "..", "docs", "openapi.yaml")
}
