// Command covergate evaluates a Go coverage profile against the repository
// coverage policy (ADR-0006): it deduplicates profile blocks, excludes
// generated and test-only code from the canonical metrics, intersects diffs
// with executable lines and fails when a configured floor is missed.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/PredictionExplorer/augur-explorer/internal/covergate"
	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

func main() {
	// Before the flag set parses: --version must win over flag validation.
	if version.HandleFlag(os.Args[1:], os.Stdout) {
		return
	}
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("covergate", flag.ContinueOnError)
	flags.SetOutput(stderr)
	policyPath := flags.String("policy", "coverage/policy.json", "coverage policy JSON")
	profilePath := flags.String("profile", "", "Go coverage profile")
	diffPath := flags.String("diff", "", "optional unified Git diff")
	jsonPath := flags.String("json", "", "optional JSON report path or - for stdout")
	check := flags.Bool("check", true, "fail when a configured floor is missed")
	commitStatus := flags.Bool("commit-status", false, "print whether the 90% commit gate is active")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	policy, err := covergate.LoadPolicy(*policyPath)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "covergate: %v\n", err)
		return 2
	}
	if *commitStatus {
		if flags.NArg() != 0 {
			_, _ = fmt.Fprintln(stderr, "covergate: positional arguments are not accepted")
			return 2
		}
		if policy.CommitGateEnabled {
			_, _ = fmt.Fprintf(stdout, "enabled %.2f\n", policy.CommitFloor)
		} else {
			_, _ = fmt.Fprintf(stdout, "deferred %.2f\n", policy.CommitFloor)
		}
		return 0
	}
	if *profilePath == "" || flags.NArg() != 0 {
		_, _ = fmt.Fprintln(stderr, "covergate: -profile is required and positional arguments are not accepted")
		return 2
	}
	profile, err := covergate.LoadProfile(*profilePath)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "covergate: %v\n", err)
		return 2
	}
	analysis := covergate.Analyze(profile, policy)

	var patch *covergate.PatchAnalysis
	if *diffPath != "" {
		changed, err := covergate.LoadDiff(*diffPath)
		if err != nil {
			_, _ = fmt.Fprintf(stderr, "covergate: %v\n", err)
			return 2
		}
		result := covergate.AnalyzePatch(profile, policy, changed)
		patch = &result
	}
	report := covergate.Evaluate(policy, analysis, patch)
	if err := printReport(stdout, policy, report); err != nil {
		_, _ = fmt.Fprintf(stderr, "covergate: write report: %v\n", err)
		return 2
	}
	if *jsonPath != "" {
		if err := writeJSONReport(*jsonPath, report, stdout); err != nil {
			_, _ = fmt.Fprintf(stderr, "covergate: %v\n", err)
			return 2
		}
	}
	if *check && !report.Passed() {
		return 1
	}
	return 0
}

func printReport(writer io.Writer, policy covergate.Policy, report covergate.Report) error {
	if err := printMetric(writer, "legacy internal", report.Analysis.LegacyInternal, policy.LegacyInternalFloor); err != nil {
		return err
	}
	if err := printMetric(writer, "handwritten internal", report.Analysis.Internal, policy.InternalFloor); err != nil {
		return err
	}
	if err := printMetric(writer, "all production", report.Analysis.Production, policy.ProductionFloor); err != nil {
		return err
	}
	if report.Patch != nil {
		if report.Patch.Applicable {
			if err := printMetric(writer, "changed code", report.Patch.Metric, policy.PatchFloor); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(writer, "%-22s N/A (%d changed lines across %d files)\n",
				"changed code",
				report.Patch.ChangedLines,
				report.Patch.ChangedFiles,
			); err != nil {
				return err
			}
		}
		const maxUncovered = 20
		for index, block := range report.Patch.Uncovered {
			if index == maxUncovered {
				if _, err := fmt.Fprintf(writer, "  ... %d more uncovered changed blocks\n",
					len(report.Patch.Uncovered)-maxUncovered); err != nil {
					return err
				}
				break
			}
			if _, err := fmt.Fprintf(writer, "  uncovered %s:%d-%d (%d statements)\n",
				block.File, block.StartLine, block.EndLine, block.Statements); err != nil {
				return err
			}
		}
	}
	if len(report.Failures) == 0 {
		_, err := fmt.Fprintln(writer, "coverage gate: PASS")
		return err
	}
	if _, err := fmt.Fprintln(writer, "coverage gate: FAIL"); err != nil {
		return err
	}
	for _, failure := range report.Failures {
		if _, err := fmt.Fprintf(writer, "  - %s\n", failure); err != nil {
			return err
		}
	}
	return nil
}

func printMetric(writer io.Writer, name string, metric covergate.Metric, floor float64) error {
	_, err := fmt.Fprintf(writer, "%-22s %6.2f%% (%d/%d; floor %.2f%%)\n",
		name,
		metric.Percent,
		metric.Covered,
		metric.Total,
		floor,
	)
	return err
}

func writeJSONReport(path string, report covergate.Report, stdout io.Writer) error {
	if path == "-" {
		return covergate.WriteJSON(stdout, report)
	}
	// #nosec G304 -- the operator explicitly chooses the local report path.
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create JSON report: %w", err)
	}
	return writeJSONReportToCloser(file, report)
}

func writeJSONReportToCloser(writer io.WriteCloser, report covergate.Report) error {
	if err := covergate.WriteJSON(writer, report); err != nil {
		_ = writer.Close()
		return err
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("close JSON report: %w", err)
	}
	return nil
}
