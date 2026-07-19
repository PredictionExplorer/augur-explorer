package covergate

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Report is the machine-readable result emitted by the coverage gate.
type Report struct {
	PolicyVersion int            `json:"policyVersion"`
	Analysis      Analysis       `json:"analysis"`
	Patch         *PatchAnalysis `json:"patch,omitempty"`
	Failures      []string       `json:"failures"`
}

// Passed reports whether every configured floor was satisfied.
func (report *Report) Passed() bool {
	return len(report.Failures) == 0
}

// Evaluate applies policy floors to global and optional patch metrics.
func Evaluate(policy Policy, analysis Analysis, patch *PatchAnalysis) Report {
	report := Report{
		PolicyVersion: policy.Version,
		Analysis:      analysis,
		Patch:         patch,
		Failures:      []string{},
	}
	report.checkMetric("legacy internal", analysis.LegacyInternal, policy.LegacyInternalFloor)
	report.checkMetric("handwritten internal", analysis.Internal, policy.InternalFloor)
	report.checkMetric("all production", analysis.Production, policy.ProductionFloor)
	if patch != nil && patch.Applicable {
		report.checkMetric("changed code", patch.Metric, policy.PatchFloor)
	}
	return report
}

func (report *Report) checkMetric(name string, metric Metric, floor float64) {
	if metric.Total == 0 {
		report.Failures = append(report.Failures, name+" coverage has no statements")
		return
	}
	if metric.Percent+1e-9 < floor {
		report.Failures = append(report.Failures, fmt.Sprintf(
			"%s coverage %.2f%% is below %.2f%%",
			name,
			metric.Percent,
			floor,
		))
	}
}

// LoadProfile reads and parses a coverage profile from path.
func LoadProfile(path string) (Profile, error) {
	// #nosec G304 -- the operator explicitly chooses the local profile path.
	file, err := os.Open(path)
	if err != nil {
		return Profile{}, fmt.Errorf("open coverage profile: %w", err)
	}
	defer func() { _ = file.Close() }()
	profile, err := ParseProfile(file)
	if err != nil {
		return Profile{}, fmt.Errorf("parse coverage profile %s: %w", path, err)
	}
	return profile, nil
}

// LoadDiff reads and parses a unified diff from path.
func LoadDiff(path string) (ChangedLines, error) {
	// #nosec G304 -- the operator explicitly chooses the local diff path.
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open coverage diff: %w", err)
	}
	defer func() { _ = file.Close() }()
	changed, err := ParseUnifiedDiff(file)
	if err != nil {
		return nil, fmt.Errorf("parse coverage diff %s: %w", path, err)
	}
	return changed, nil
}

// WriteJSON writes report as stable indented JSON.
func WriteJSON(writer io.Writer, report Report) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(report); err != nil {
		return fmt.Errorf("encode coverage report: %w", err)
	}
	return nil
}
