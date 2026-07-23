package smoketest

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Suite identifies one API surface exercised by the smoke-test runner.
type Suite string

// Supported smoke-test suites.
const (
	SuiteV1          Suite = "v1"
	SuiteV2          Suite = "v2"
	SuiteBoth        Suite = "both"
	SuiteOperational Suite = "operational"
)

// ParseSuite validates a command-facing suite name.
func ParseSuite(value string) (Suite, error) {
	suite := Suite(strings.ToLower(strings.TrimSpace(value)))
	switch suite {
	case SuiteV1, SuiteV2, SuiteBoth, SuiteOperational:
		return suite, nil
	default:
		return "", fmt.Errorf(
			"invalid smoke-test suite %q (want v1, v2, both, or operational)",
			value,
		)
	}
}

// RequiresParameters reports whether the suite needs production-shaped URL
// values from the parameter source.
func (s Suite) RequiresParameters() bool {
	return s == SuiteV1 || s == SuiteV2 || s == SuiteBoth
}

func (s Suite) members() ([]Suite, error) {
	switch s {
	case SuiteV1:
		return []Suite{SuiteV1}, nil
	case SuiteV2:
		return []Suite{SuiteV2}, nil
	case SuiteBoth:
		// Canonical v2 runs first; the frozen compatibility surface follows.
		return []Suite{SuiteV2, SuiteV1}, nil
	case SuiteOperational:
		return []Suite{SuiteOperational}, nil
	default:
		return nil, errors.New("smoke-test suite is invalid")
	}
}

// Probe is one deterministic GET request in a smoke-test suite.
type Probe struct {
	Suite       Suite
	OperationID string
	Template    string
	Endpoint    string
	PathParams  map[string]string

	spec      *openapi3.T
	pathItem  *openapi3.PathItem
	operation *openapi3.Operation
}

func buildProbes(suite Suite, params Params) ([]Probe, error) {
	switch suite {
	case SuiteV1:
		endpoints := BuildEndpoints(params)
		probes := make([]Probe, 0, len(endpoints))
		for _, endpoint := range endpoints {
			probes = append(probes, Probe{
				Suite:    SuiteV1,
				Template: endpoint,
				Endpoint: endpoint,
			})
		}
		return probes, nil
	case SuiteV2:
		return BuildV2Probes(params)
	case SuiteOperational:
		return BuildOperationalProbes()
	default:
		return nil, fmt.Errorf("cannot build probes for suite %q", suite)
	}
}

func (p Probe) method() string {
	return http.MethodGet
}
