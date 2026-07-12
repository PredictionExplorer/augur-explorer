// Package covergate parses Go coverage profiles and enforces repository
// coverage policy against global and changed-code metrics.
package covergate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Policy is the versioned coverage policy shared by local hooks and CI.
type Policy struct {
	Version               int      `json:"version"`
	LegacyInternalFloor   float64  `json:"legacyInternalFloor"`
	InternalFloor         float64  `json:"internalFloor"`
	ProductionFloor       float64  `json:"productionFloor"`
	PatchFloor            float64  `json:"patchFloor"`
	InternalTarget        float64  `json:"internalTarget"`
	ProductionTarget      float64  `json:"productionTarget"`
	CriticalPackageTarget float64  `json:"criticalPackageTarget"`
	CommitGateEnabled     bool     `json:"commitGateEnabled"`
	CommitFloor           float64  `json:"commitFloor"`
	ExcludeFiles          []string `json:"excludeFiles"`
	ExcludePrefixes       []string `json:"excludePrefixes"`
}

// LoadPolicy reads and validates a policy JSON document from path.
func LoadPolicy(path string) (Policy, error) {
	// #nosec G304 -- the operator explicitly chooses the repository policy path.
	file, err := os.Open(path)
	if err != nil {
		return Policy{}, fmt.Errorf("open coverage policy: %w", err)
	}
	defer func() { _ = file.Close() }()
	policy, err := DecodePolicy(file)
	if err != nil {
		return Policy{}, fmt.Errorf("decode coverage policy %s: %w", path, err)
	}
	return policy, nil
}

// DecodePolicy decodes one strict policy JSON document.
func DecodePolicy(reader io.Reader) (Policy, error) {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	var policy Policy
	if err := decoder.Decode(&policy); err != nil {
		return Policy{}, err
	}
	var trailing any
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		if err == nil {
			return Policy{}, errors.New("coverage policy contains trailing JSON")
		}
		return Policy{}, fmt.Errorf("decode trailing coverage policy data: %w", err)
	}
	if err := policy.Validate(); err != nil {
		return Policy{}, err
	}
	slices.Sort(policy.ExcludeFiles)
	slices.Sort(policy.ExcludePrefixes)
	return policy, nil
}

// Validate checks policy ranges and exclusion path safety.
func (policy Policy) Validate() error {
	if policy.Version != 1 {
		return fmt.Errorf("unsupported coverage policy version %d", policy.Version)
	}
	thresholds := map[string]float64{
		"legacyInternalFloor":   policy.LegacyInternalFloor,
		"internalFloor":         policy.InternalFloor,
		"productionFloor":       policy.ProductionFloor,
		"patchFloor":            policy.PatchFloor,
		"internalTarget":        policy.InternalTarget,
		"productionTarget":      policy.ProductionTarget,
		"criticalPackageTarget": policy.CriticalPackageTarget,
		"commitFloor":           policy.CommitFloor,
	}
	for name, threshold := range thresholds {
		if threshold < 0 || threshold > 100 {
			return fmt.Errorf("%s must be between 0 and 100", name)
		}
	}
	if policy.InternalTarget < policy.InternalFloor {
		return errors.New("internal target is below its floor")
	}
	if policy.ProductionTarget < policy.ProductionFloor {
		return errors.New("production target is below its floor")
	}
	if policy.CommitGateEnabled && policy.InternalFloor < policy.CommitFloor {
		return errors.New("commit gate cannot activate below its commit floor")
	}
	for _, path := range append(slices.Clone(policy.ExcludeFiles), policy.ExcludePrefixes...) {
		if path == "" || strings.HasPrefix(path, "/") || strings.Contains(path, "..") ||
			strings.Contains(path, `\`) {
			return fmt.Errorf("unsafe coverage exclusion %q", path)
		}
	}
	return nil
}

func (policy Policy) excluded(path string) bool {
	if _, found := slices.BinarySearch(policy.ExcludeFiles, path); found {
		return true
	}
	return slices.ContainsFunc(policy.ExcludePrefixes, func(prefix string) bool {
		return strings.HasPrefix(path, prefix)
	})
}
