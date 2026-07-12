package covergate

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEvaluateCoveragePolicy(t *testing.T) {
	t.Parallel()
	policy, err := DecodePolicy(strings.NewReader(validPolicyJSON))
	if err != nil {
		t.Fatal(err)
	}
	passing := Analysis{
		LegacyInternal: Metric{Covered: 80, Total: 100, Percent: 80},
		Internal:       Metric{Covered: 80, Total: 100, Percent: 80},
		Production:     Metric{Covered: 60, Total: 100, Percent: 60},
	}
	patch := &PatchAnalysis{
		Metric:     Metric{Covered: 95, Total: 100, Percent: 95},
		Applicable: true,
	}
	if report := Evaluate(policy, passing, patch); !report.Passed() {
		t.Fatalf("passing report = %+v", report)
	}

	failing := passing
	failing.LegacyInternal = Metric{Covered: 75, Total: 100, Percent: 75}
	failing.Internal = Metric{Covered: 69, Total: 100, Percent: 69}
	failing.Production = Metric{Covered: 49, Total: 100, Percent: 49}
	patch.Metric = Metric{Covered: 94, Total: 100, Percent: 94}
	report := Evaluate(policy, failing, patch)
	if report.Passed() || len(report.Failures) != 4 {
		t.Fatalf("failing report = %+v", report)
	}
}

func TestEvaluateIgnoresNonApplicablePatch(t *testing.T) {
	t.Parallel()
	policy, err := DecodePolicy(strings.NewReader(validPolicyJSON))
	if err != nil {
		t.Fatal(err)
	}
	analysis := Analysis{
		LegacyInternal: Metric{Covered: 80, Total: 100, Percent: 80},
		Internal:       Metric{Covered: 80, Total: 100, Percent: 80},
		Production:     Metric{Covered: 60, Total: 100, Percent: 60},
	}
	report := Evaluate(policy, analysis, &PatchAnalysis{})
	if !report.Passed() {
		t.Fatalf("report = %+v", report)
	}
}

func TestEvaluateRejectsEmptyGlobalMetric(t *testing.T) {
	t.Parallel()
	policy, err := DecodePolicy(strings.NewReader(validPolicyJSON))
	if err != nil {
		t.Fatal(err)
	}
	report := Evaluate(policy, Analysis{}, nil)
	if report.Passed() || len(report.Failures) != 3 {
		t.Fatalf("report = %+v", report)
	}
}

func TestWriteJSON(t *testing.T) {
	t.Parallel()
	report := Report{
		PolicyVersion: 1,
		Analysis: Analysis{
			Internal: Metric{Covered: 1, Total: 2, Percent: 50},
			Packages: map[string]Metric{
				"internal/example": {Covered: 1, Total: 2, Percent: 50},
			},
		},
		Failures: []string{},
	}
	var output bytes.Buffer
	if err := WriteJSON(&output, report); err != nil {
		t.Fatal(err)
	}
	var decoded Report
	if err := json.Unmarshal(output.Bytes(), &decoded); err != nil {
		t.Fatalf("invalid JSON: %v\n%s", err, output.String())
	}
	if decoded.PolicyVersion != 1 ||
		decoded.Analysis.Packages["internal/example"].Covered != 1 ||
		decoded.Failures == nil {
		t.Fatalf("decoded report = %+v", decoded)
	}
}

func TestLoadPolicyProfileAndDiff(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	policyPath := writeCovergateFixture(t, dir, "policy.json", validPolicyJSON)
	profilePath := writeCovergateFixture(t, dir, "coverage.out", testProfile)
	diffPath := writeCovergateFixture(t, dir, "change.diff", testDiff)
	if _, err := LoadPolicy(policyPath); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadProfile(profilePath); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadDiff(diffPath); err != nil {
		t.Fatal(err)
	}

	for name, load := range map[string]func(string) error{
		"policy": func(path string) error {
			_, err := LoadPolicy(path)
			return err
		},
		"profile": func(path string) error {
			_, err := LoadProfile(path)
			return err
		},
		"diff": func(path string) error {
			_, err := LoadDiff(path)
			return err
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := load(filepath.Join(dir, "missing-"+name)); err == nil {
				t.Fatal("missing file was accepted")
			}
			malformedContent := map[string]string{
				"policy":  "{",
				"profile": "not coverage",
				"diff":    "@@ -1 +1 @@\n+line\n",
			}[name]
			malformed := writeCovergateFixture(t, t.TempDir(), name, malformedContent)
			if err := load(malformed); err == nil {
				t.Fatal("malformed file was accepted")
			}
		})
	}
}

func TestWriteJSONPropagatesWriterError(t *testing.T) {
	t.Parallel()
	if err := WriteJSON(errorWriter{}, Report{}); err == nil {
		t.Fatal("writer error was ignored")
	}
}

func writeCovergateFixture(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

type errorWriter struct{}

func (errorWriter) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}
