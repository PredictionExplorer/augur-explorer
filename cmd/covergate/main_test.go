package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/covergate"
)

const cliPolicy = `{
	"version": 1,
	"legacyInternalFloor": 80,
	"internalFloor": 60,
	"productionFloor": 20,
	"patchFloor": 30,
	"internalTarget": 90,
	"productionTarget": 90,
	"criticalPackageTarget": 95,
	"commitGateEnabled": false,
	"commitFloor": 90,
	"excludeFiles": ["internal/api/v2/api.gen.go"],
	"excludePrefixes": ["internal/testutil/"]
}`

const cliProfile = `mode: atomic
github.com/example/project/internal/core/core.go:10.1,12.2 2 1
github.com/example/project/internal/core/core.go:20.1,20.9 1 0
github.com/example/project/internal/api/v2/api.gen.go:1.1,3.2 3 1
github.com/example/project/internal/testutil/helper.go:1.1,2.2 2 1
github.com/example/project/cmd/tool/main.go:5.1,9.2 4 0
`

const cliDiff = `diff --git a/internal/core/core.go b/internal/core/core.go
--- a/internal/core/core.go
+++ b/internal/core/core.go
@@ -10 +10 @@
-old
+new
diff --git a/cmd/tool/main.go b/cmd/tool/main.go
--- a/cmd/tool/main.go
+++ b/cmd/tool/main.go
@@ -5 +5 @@
-old
+new
`

func TestRunPassesAndWritesReport(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	policy := writeTestFile(t, dir, "policy.json", cliPolicy)
	profile := writeTestFile(t, dir, "coverage.out", cliProfile)
	diff := writeTestFile(t, dir, "change.diff", cliDiff)
	report := filepath.Join(dir, "report.json")
	var stdout, stderr bytes.Buffer
	code := run([]string{
		"-policy", policy,
		"-profile", profile,
		"-diff", diff,
		"-json", report,
	}, &stdout, &stderr)
	if code != 0 || !strings.Contains(stdout.String(), "coverage gate: PASS") ||
		stderr.Len() != 0 {
		t.Fatalf("code=%d stdout=%q stderr=%q", code, stdout.String(), stderr.String())
	}
	// #nosec G304 -- report is inside the test's temporary directory.
	data, err := os.ReadFile(report)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(data, []byte(`"policyVersion": 1`)) {
		t.Fatalf("report = %s", data)
	}
}

func TestRunFailsThresholdAndSupportsReportOnly(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	policy := writeTestFile(t, dir, "policy.json",
		strings.Replace(cliPolicy, `"patchFloor": 30`, `"patchFloor": 95`, 1))
	profile := writeTestFile(t, dir, "coverage.out", cliProfile)
	diff := writeTestFile(t, dir, "change.diff", cliDiff)

	var stdout, stderr bytes.Buffer
	code := run([]string{
		"-policy", policy, "-profile", profile, "-diff", diff,
	}, &stdout, &stderr)
	if code != 1 || !strings.Contains(stdout.String(), "changed code coverage") {
		t.Fatalf("code=%d stdout=%q stderr=%q", code, stdout.String(), stderr.String())
	}

	stdout.Reset()
	stderr.Reset()
	code = run([]string{
		"-policy", policy, "-profile", profile, "-diff", diff, "-check=false",
	}, &stdout, &stderr)
	if code != 0 || !strings.Contains(stdout.String(), "coverage gate: FAIL") {
		t.Fatalf("report-only code=%d stdout=%q stderr=%q", code, stdout.String(), stderr.String())
	}
}

func TestRunRejectsUsageAndMalformedInputs(t *testing.T) {
	t.Parallel()
	var stdout, stderr bytes.Buffer
	if code := run(nil, &stdout, &stderr); code != 2 {
		t.Fatalf("missing profile code = %d", code)
	}
	if code := run([]string{"-unknown"}, &stdout, &stderr); code != 2 {
		t.Fatalf("unknown flag code = %d", code)
	}

	dir := t.TempDir()
	policy := writeTestFile(t, dir, "policy.json", cliPolicy)
	validProfile := writeTestFile(t, dir, "valid-coverage.out", cliProfile)
	profile := writeTestFile(t, dir, "coverage.out", "not coverage")
	stdout.Reset()
	stderr.Reset()
	if code := run([]string{"-policy", policy, "-profile", profile}, &stdout, &stderr); code != 2 {
		t.Fatalf("malformed profile code = %d", code)
	}
	if !strings.Contains(stderr.String(), "parse coverage profile") {
		t.Fatalf("stderr = %q", stderr.String())
	}

	stdout.Reset()
	stderr.Reset()
	if code := run([]string{
		"-policy", filepath.Join(dir, "missing-policy.json"),
		"-profile", validProfile,
	}, &stdout, &stderr); code != 2 {
		t.Fatalf("missing policy code = %d", code)
	}
	stdout.Reset()
	stderr.Reset()
	if code := run([]string{
		"-policy", policy,
		"-profile", validProfile,
		"-diff", filepath.Join(dir, "missing.diff"),
	}, &stdout, &stderr); code != 2 {
		t.Fatalf("missing diff code = %d", code)
	}
	stdout.Reset()
	stderr.Reset()
	if code := run([]string{
		"-policy", policy,
		"-profile", validProfile,
		"-json", filepath.Join(dir, "missing", "report.json"),
	}, &stdout, &stderr); code != 2 {
		t.Fatalf("invalid report path code = %d", code)
	}
}

func TestRunWritesJSONToStdout(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	policy := writeTestFile(t, dir, "policy.json", cliPolicy)
	profile := writeTestFile(t, dir, "coverage.out", cliProfile)
	var stdout, stderr bytes.Buffer
	code := run([]string{
		"-policy", policy, "-profile", profile, "-json=-",
	}, &stdout, &stderr)
	if code != 0 || !strings.Contains(stdout.String(), `"policyVersion": 1`) {
		t.Fatalf("code=%d stdout=%q stderr=%q", code, stdout.String(), stderr.String())
	}
}

func TestRunReportsDeferredAndEnabledCommitPolicy(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	deferred := writeTestFile(t, dir, "deferred.json", cliPolicy)
	var stdout, stderr bytes.Buffer
	if code := run([]string{
		"-policy", deferred, "-commit-status",
	}, &stdout, &stderr); code != 0 ||
		strings.TrimSpace(stdout.String()) != "deferred 90.00" {
		t.Fatalf("deferred code=%d stdout=%q stderr=%q", code, stdout.String(), stderr.String())
	}

	enabledPolicy := strings.NewReplacer(
		`"commitGateEnabled": false`, `"commitGateEnabled": true`,
		`"internalFloor": 60`, `"internalFloor": 90`,
	).Replace(cliPolicy)
	enabled := writeTestFile(t, dir, "enabled.json", enabledPolicy)
	stdout.Reset()
	stderr.Reset()
	if code := run([]string{
		"-policy", enabled, "-commit-status",
	}, &stdout, &stderr); code != 0 ||
		strings.TrimSpace(stdout.String()) != "enabled 90.00" {
		t.Fatalf("enabled code=%d stdout=%q stderr=%q", code, stdout.String(), stderr.String())
	}

	stdout.Reset()
	stderr.Reset()
	if code := run([]string{
		"-policy", deferred, "-commit-status", "extra",
	}, &stdout, &stderr); code != 2 {
		t.Fatalf("commit status positional code = %d", code)
	}
}

func TestPrintReportPropagatesEveryWriterFailure(t *testing.T) {
	t.Parallel()
	policy := mustCLIPolicy(t)
	uncovered := make([]covergate.UncoveredBlock, 25)
	for index := range uncovered {
		uncovered[index] = covergate.UncoveredBlock{
			File: "internal/example/example.go", StartLine: index + 1,
			EndLine: index + 1, Statements: 1,
		}
	}
	report := covergate.Report{
		Analysis: covergate.Analysis{
			LegacyInternal: covergate.Metric{Covered: 8, Total: 10, Percent: 80},
			Internal:       covergate.Metric{Covered: 8, Total: 10, Percent: 80},
			Production:     covergate.Metric{Covered: 6, Total: 10, Percent: 60},
		},
		Patch: &covergate.PatchAnalysis{
			Applicable: true,
			Metric:     covergate.Metric{Covered: 9, Total: 10, Percent: 90},
			Uncovered:  uncovered,
		},
		Failures: []string{"first", "second"},
	}
	var success bytes.Buffer
	if err := printReport(&success, policy, report); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(success.String(), "5 more uncovered") ||
		!strings.Contains(success.String(), "coverage gate: FAIL") {
		t.Fatalf("report = %s", success.String())
	}
	for failureAt := 1; failureAt <= 28; failureAt++ {
		writer := &failWriter{failAt: failureAt}
		if err := printReport(writer, policy, report); err == nil {
			t.Errorf("writer failure %d was ignored", failureAt)
		}
	}

	nothingChanged := report
	nothingChanged.Patch = &covergate.PatchAnalysis{ChangedFiles: 1, ChangedLines: 2}
	nothingChanged.Failures = nil
	success.Reset()
	if err := printReport(&success, policy, nothingChanged); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(success.String(), "N/A") ||
		!strings.Contains(success.String(), "coverage gate: PASS") {
		t.Fatalf("N/A report = %s", success.String())
	}
}

func TestRunRejectsBrokenOutputAndJSONWriter(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	policy := writeTestFile(t, dir, "policy.json", cliPolicy)
	profile := writeTestFile(t, dir, "coverage.out", cliProfile)
	writer := &failWriter{failAt: 1}
	var stderr bytes.Buffer
	if code := run([]string{"-policy", policy, "-profile", profile}, writer, &stderr); code != 2 {
		t.Fatalf("broken output code = %d", code)
	}
	if err := writeJSONReport("-", covergate.Report{}, &failWriter{failAt: 1}); err == nil {
		t.Fatal("broken JSON writer was ignored")
	}
	if err := writeJSONReportToCloser(
		&failWriteCloser{writeErr: errors.New("write failed")},
		covergate.Report{},
	); err == nil {
		t.Fatal("JSON file write error was ignored")
	}
	if err := writeJSONReportToCloser(
		&failWriteCloser{closeErr: errors.New("close failed")},
		covergate.Report{},
	); err == nil || !strings.Contains(err.Error(), "close JSON report") {
		t.Fatalf("JSON close error = %v", err)
	}
}

func mustCLIPolicy(t *testing.T) covergate.Policy {
	t.Helper()
	policy, err := covergate.DecodePolicy(strings.NewReader(cliPolicy))
	if err != nil {
		t.Fatal(err)
	}
	return policy
}

type failWriter struct {
	writes int
	failAt int
}

func (writer *failWriter) Write(data []byte) (int, error) {
	writer.writes++
	if writer.writes == writer.failAt {
		return 0, errors.New("write failed")
	}
	return len(data), nil
}

type failWriteCloser struct {
	writeErr error
	closeErr error
}

func (writer *failWriteCloser) Write(data []byte) (int, error) {
	if writer.writeErr != nil {
		return 0, writer.writeErr
	}
	return len(data), nil
}

func (writer *failWriteCloser) Close() error {
	return writer.closeErr
}

func writeTestFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}
