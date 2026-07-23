package smoketest

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/PredictionExplorer/augur-explorer/internal/api/policy"
	apiv2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
)

var operationalTemplates = []string{
	"/api/v2/cosmicgame/rounds",
	"/api/v2/cosmicgame/statistics/counters",
	"/api/v2/randomwalk/statistics",
}

// BuildV2Probes constructs one strict probe for every GET operation in the
// canonical embedded v2 OpenAPI document. A newly required parameter without
// an explicit production-shaped binding makes construction fail instead of
// silently dropping coverage.
func BuildV2Probes(input Params) ([]Probe, error) {
	params := WithDefaults(input)
	spec, err := apiv2.GetSpec()
	if err != nil {
		return nil, fmt.Errorf("load embedded OpenAPI v2 document: %w", err)
	}

	templates := make([]string, 0, spec.Paths.Len())
	for template, item := range spec.Paths.Map() {
		if item.Get != nil {
			templates = append(templates, template)
		}
	}
	slices.Sort(templates)

	probes := make([]Probe, 0, len(templates))
	for _, template := range templates {
		item := spec.Paths.Value(template)
		if item == nil || item.Get == nil {
			return nil, fmt.Errorf("OpenAPI GET operation disappeared at %s", template)
		}
		probe, err := buildV2Probe(spec, template, item, params)
		if err != nil {
			return nil, err
		}
		probes = append(probes, probe)
	}
	return probes, nil
}

func buildV2Probe(
	spec *openapi3.T,
	template string,
	item *openapi3.PathItem,
	params Params,
) (Probe, error) {
	operation := item.Get
	if operation.OperationID == "" {
		return Probe{}, fmt.Errorf("OpenAPI GET %s has no operationId", template)
	}
	if policy.V1Deprecated(template) {
		return Probe{}, fmt.Errorf("OpenAPI v2 operation %s uses deprecated path %s", operation.OperationID, template)
	}

	endpointPath := template
	pathParams := make(map[string]string)
	query := make(url.Values)
	for _, parameterRef := range operationParameters(item, operation) {
		parameter := parameterRef.Value
		if parameter == nil {
			return Probe{}, fmt.Errorf("%s %s has an unresolved parameter", operation.OperationID, template)
		}
		switch parameter.In {
		case openapi3.ParameterInPath:
			value, ok := v2PathValue(template, parameter.Name, params)
			if !ok {
				return Probe{}, fmt.Errorf(
					"%s %s has unmapped path parameter %q",
					operation.OperationID,
					template,
					parameter.Name,
				)
			}
			pathParams[parameter.Name] = value
			endpointPath = strings.ReplaceAll(
				endpointPath,
				"{"+parameter.Name+"}",
				url.PathEscape(value),
			)
		case openapi3.ParameterInQuery:
			value, include, err := v2QueryValue(template, parameter, params)
			if err != nil {
				return Probe{}, fmt.Errorf("%s %s: %w", operation.OperationID, template, err)
			}
			if include {
				query.Set(parameter.Name, value)
			}
		}
	}
	if strings.Contains(endpointPath, "{") {
		return Probe{}, fmt.Errorf("%s %s contains an unbound path parameter", operation.OperationID, template)
	}
	endpoint := endpointPath
	if encoded := query.Encode(); encoded != "" {
		endpoint += "?" + encoded
	}
	return Probe{
		Suite:       SuiteV2,
		OperationID: operation.OperationID,
		Template:    template,
		Endpoint:    endpoint,
		PathParams:  pathParams,
		spec:        spec,
		pathItem:    item,
		operation:   operation,
	}, nil
}

func operationParameters(item *openapi3.PathItem, operation *openapi3.Operation) openapi3.Parameters {
	parameters := make(openapi3.Parameters, 0, len(item.Parameters)+len(operation.Parameters))
	parameters = append(parameters, item.Parameters...)
	parameters = append(parameters, operation.Parameters...)
	return parameters
}

func v2PathValue(template, name string, params Params) (string, bool) {
	switch name {
	case "address":
		if strings.Contains(template, "/staking/cst/deposits/{depositId}/rewards") {
			return params.CSTStakerAddress, true
		}
		return params.UserAddress, true
	case "round":
		if strings.Contains(template, "/bids/{position}") {
			return params.BidRound, true
		}
		return params.RoundNumber, true
	case "position":
		return params.BidPosition, true
	case "depositId":
		return params.DepositID, true
	case "actionId":
		if strings.Contains(template, "/random-walk/") {
			return params.RandomWalkActionID, true
		}
		return params.CSTActionID, true
	case "nftTokenId":
		return params.TokenID, true
	case "tokenId":
		return params.RandomWalkTokenID, true
	default:
		return "", false
	}
}

func v2QueryValue(template string, parameter *openapi3.Parameter, params Params) (string, bool, error) {
	if parameter.Name == "limit" {
		return "1", true, nil
	}
	if !parameter.Required {
		return "", false, nil
	}

	switch parameter.Name {
	case "from", "to":
		var schema *openapi3.Schema
		if parameter.Schema != nil {
			schema = parameter.Schema.Value
		}
		if schema != nil && schema.Format == "date" {
			value := params.FromDate
			if parameter.Name == "to" {
				value = params.ToDate
			}
			iso, err := compactDateToISO(value)
			if err != nil {
				return "", false, fmt.Errorf("bind required query parameter %q: %w", parameter.Name, err)
			}
			return iso, true, nil
		}
		if parameter.Name == "from" {
			return params.TimestampMin, true, nil
		}
		return params.TimestampMax, true, nil
	case "pool":
		var schema *openapi3.Schema
		if parameter.Schema != nil {
			schema = parameter.Schema.Value
		}
		if schema != nil {
			for _, candidate := range schema.Enum {
				if value, ok := candidate.(string); ok && value != "" {
					return value, true, nil
				}
			}
		}
		return "", false, fmt.Errorf("required pool parameter on %s has no string enum", template)
	default:
		return "", false, fmt.Errorf("unmapped required query parameter %q", parameter.Name)
	}
}

func compactDateToISO(value string) (string, error) {
	parsed, err := time.Parse("20060102", value)
	if err != nil {
		return "", fmt.Errorf("date %q is not YYYYMMDD: %w", value, err)
	}
	return parsed.Format(time.DateOnly), nil
}

// BuildOperationalProbes returns the D6-safe, no-parameter-discovery probe
// set used by routine monitoring and Compose. Live contract-cache resources
// are deliberately absent; the full v2 suite checks those as strict 200s.
func BuildOperationalProbes() ([]Probe, error) {
	v2Probes, err := BuildV2Probes(DefaultParams())
	if err != nil {
		return nil, err
	}
	byTemplate := make(map[string]Probe, len(v2Probes))
	for _, probe := range v2Probes {
		byTemplate[probe.Template] = probe
	}

	probes := []Probe{
		{Suite: SuiteOperational, OperationID: "healthz", Template: "/healthz", Endpoint: "/healthz"},
		{Suite: SuiteOperational, OperationID: "readyz", Template: "/readyz", Endpoint: "/readyz"},
		{Suite: SuiteOperational, OperationID: "version", Template: "/version", Endpoint: "/version"},
	}
	for _, template := range operationalTemplates {
		probe, ok := byTemplate[template]
		if !ok {
			return nil, fmt.Errorf("operational probe template %s is absent from OpenAPI v2", template)
		}
		probe.Suite = SuiteOperational
		probes = append(probes, probe)
	}
	return probes, nil
}
