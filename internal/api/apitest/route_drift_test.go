package apitest

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"

	"github.com/goccy/go-yaml"
)

// TestRouteDriftAgainstOpenAPI pins the route table to docs/openapi.yaml in
// both directions: every documented operation must be registered on the gin
// router, and every registered route must be documented. The spec is the v1
// contract for the rewrite — this test makes silent drift impossible.
func TestRouteDriftAgainstOpenAPI(t *testing.T) {
	specOps := loadSpecOperations(t)

	r := buildBareRouter()
	registered := make(map[string]bool)
	for _, route := range r.Routes() {
		registered[route.Method+" "+route.Path] = true
	}

	var missingFromRouter []string
	for op := range specOps {
		if !registered[op] {
			missingFromRouter = append(missingFromRouter, op)
		}
	}
	sort.Strings(missingFromRouter)
	for _, op := range missingFromRouter {
		t.Errorf("documented in openapi.yaml but not registered on the router: %s", op)
	}

	var missingFromSpec []string
	for op := range registered {
		if !specOps[op] {
			missingFromSpec = append(missingFromSpec, op)
		}
	}
	sort.Strings(missingFromSpec)
	for _, op := range missingFromSpec {
		t.Errorf("registered on the router but not documented in openapi.yaml: %s", op)
	}
}

// loadSpecOperations parses docs/openapi.yaml and returns the set of
// "METHOD /gin/style/:path" operation keys it documents.
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
		ginPath := openAPIPathToGin(specPath)
		for key := range item {
			lk := strings.ToLower(key)
			if !methods[lk] {
				continue // parameters, summary, etc.
			}
			ops[fmt.Sprintf("%s %s", strings.ToUpper(lk), ginPath)] = true
		}
	}
	return ops
}

// openAPIPathToGin converts /a/{b}/c template syntax to gin's /a/:b/c.
func openAPIPathToGin(p string) string {
	var b strings.Builder
	for _, seg := range strings.Split(p, "/") {
		if seg == "" {
			continue
		}
		b.WriteByte('/')
		if strings.HasPrefix(seg, "{") && strings.HasSuffix(seg, "}") {
			b.WriteByte(':')
			b.WriteString(seg[1 : len(seg)-1])
			continue
		}
		b.WriteString(seg)
	}
	if b.Len() == 0 {
		return "/"
	}
	return b.String()
}

// openapiPath locates docs/openapi.yaml relative to this source file.
func openapiPath() string {
	_, thisFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(thisFile), "..", "..", "..", "docs", "openapi.yaml")
}
