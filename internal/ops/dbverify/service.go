package dbverify

import (
	"context"
	"errors"
	"fmt"
)

const (
	DefaultVerifyReportLimit = 10
	DefaultDiffReportLimit   = 20
)

// VerifyReport is the complete db verify result.
type VerifyReport struct {
	Events             Comparison
	Transactions       Comparison
	Blocks             Comparison
	PrimaryTxHashCount int
	PrimaryBlockCount  int
}

// Matched reports whether all three compared tables match.
func (r VerifyReport) Matched() bool {
	return r.Events.Matched() && r.Transactions.Matched() && r.Blocks.Matched()
}

// VerifyDatabases loads the primary project scope and the complete secondary
// tables, then applies pure comparisons. A mismatch is data in the returned
// report; only loader failures are returned as errors.
func VerifyDatabases(
	ctx context.Context,
	primary Loader,
	secondary Loader,
	contractAddressIDs []int64,
	reportLimit int,
) (VerifyReport, error) {
	var report VerifyReport
	if primary == nil || secondary == nil {
		return report, errors.New("db verify: primary and secondary loaders are required")
	}
	if reportLimit <= 0 {
		reportLimit = DefaultVerifyReportLimit
	}

	primaryEvents, err := primary.LoadEventRecords(ctx, contractAddressIDs)
	if err != nil {
		return report, fmt.Errorf("query evt_log (primary): %w", err)
	}
	if err := ctx.Err(); err != nil {
		return report, err
	}
	secondaryEvents, err := secondary.LoadEventRecords(ctx, nil)
	if err != nil {
		return report, fmt.Errorf("query evt_log (secondary): %w", err)
	}
	report.Events = CompareEvents(primaryEvents, secondaryEvents, reportLimit)

	if err := ctx.Err(); err != nil {
		return report, err
	}
	txHashes, err := primary.TransactionHashesFromEvents(ctx, contractAddressIDs)
	if err != nil {
		return report, fmt.Errorf("get tx_hashes: %w", err)
	}
	report.PrimaryTxHashCount = len(txHashes)
	if err := ctx.Err(); err != nil {
		return report, err
	}
	primaryTransactions, err := primary.LoadTransactions(ctx, txHashes)
	if err != nil {
		return report, fmt.Errorf("query transactions (primary): %w", err)
	}
	if err := ctx.Err(); err != nil {
		return report, err
	}
	secondaryTransactions, err := secondary.LoadTransactions(ctx, nil)
	if err != nil {
		return report, fmt.Errorf("query transactions (secondary): %w", err)
	}
	report.Transactions = CompareTransactions(primaryTransactions, secondaryTransactions, reportLimit)

	if err := ctx.Err(); err != nil {
		return report, err
	}
	blockNumbers, err := primary.BlockNumbersFromEvents(ctx, contractAddressIDs)
	if err != nil {
		return report, fmt.Errorf("get block_nums: %w", err)
	}
	report.PrimaryBlockCount = len(blockNumbers)
	if err := ctx.Err(); err != nil {
		return report, err
	}
	primaryBlocks, err := primary.LoadBlocks(ctx, blockNumbers)
	if err != nil {
		return report, fmt.Errorf("query blocks (primary): %w", err)
	}
	if err := ctx.Err(); err != nil {
		return report, err
	}
	secondaryBlocks, err := secondary.LoadBlocks(ctx, nil)
	if err != nil {
		return report, fmt.Errorf("query blocks (secondary): %w", err)
	}
	if err := ctx.Err(); err != nil {
		return report, err
	}
	report.Blocks = CompareBlocks(primaryBlocks, secondaryBlocks, reportLimit)
	return report, nil
}

// FormatVerifyReport returns stable log lines for opsctl db verify.
func FormatVerifyReport(report VerifyReport) []string {
	var lines []string
	lines = append(lines,
		"",
		"=== Comparing evt_log tables ===",
		fmt.Sprintf("Primary: %d events", report.Events.PrimaryCount),
		fmt.Sprintf("Secondary: %d events", report.Events.SecondaryCount),
	)
	lines = appendComparison(lines, "evt_log", report.Events)

	lines = append(lines,
		"",
		"=== Comparing transaction tables ===",
		fmt.Sprintf("Primary has %d unique transactions for our contracts", report.PrimaryTxHashCount),
		fmt.Sprintf("Primary: %d transactions loaded", report.Transactions.PrimaryCount),
		fmt.Sprintf("Secondary: %d transactions", report.Transactions.SecondaryCount),
	)
	lines = appendComparison(lines, "transaction", report.Transactions)

	lines = append(lines,
		"",
		"=== Comparing block tables ===",
		fmt.Sprintf("Primary has %d unique blocks for our contracts", report.PrimaryBlockCount),
		fmt.Sprintf("Primary: %d blocks loaded", report.Blocks.PrimaryCount),
		fmt.Sprintf("Secondary: %d blocks", report.Blocks.SecondaryCount),
	)
	lines = appendComparison(lines, "block", report.Blocks)

	lines = append(lines, "", "=== FINAL SUMMARY ===")
	if report.Matched() {
		return append(lines, "✓ All tables match perfectly!")
	}
	if !report.Events.Matched() {
		lines = append(lines, "✗ evt_log: MISMATCH")
	}
	if !report.Transactions.Matched() {
		lines = append(lines, "✗ transaction: MISMATCH")
	}
	if !report.Blocks.Matched() {
		lines = append(lines, "✗ block: MISMATCH")
	}
	return lines
}

func appendComparison(lines []string, name string, comparison Comparison) []string {
	for _, issue := range comparison.Missing {
		lines = append(lines, "MISSING in secondary: "+issue.Summary)
	}
	if remaining := comparison.MissingTotal - len(comparison.Missing); remaining > 0 {
		lines = append(lines, fmt.Sprintf("... and %d more missing records", remaining))
	}
	for _, issue := range comparison.Extra {
		lines = append(lines, "EXTRA in secondary: "+issue.Summary)
	}
	if remaining := comparison.ExtraTotal - len(comparison.Extra); remaining > 0 {
		lines = append(lines, fmt.Sprintf("... and %d more extra records", remaining))
	}
	for _, mismatch := range comparison.Mismatched {
		lines = append(lines, fmt.Sprintf("MISMATCH in secondary: %s: %s", mismatch.Summary, formatFields(mismatch.Fields)))
	}
	if remaining := comparison.MismatchTotal - len(comparison.Mismatched); remaining > 0 {
		lines = append(lines, fmt.Sprintf("... and %d more mismatched records", remaining))
	}
	return append(lines, fmt.Sprintf(
		"%s: Missing=%d, Extra=%d, Mismatched=%d",
		name,
		comparison.MissingTotal,
		comparison.ExtraTotal,
		comparison.MismatchTotal,
	))
}

// EventLogDiffReport is the informational evtlog-diff result.
type EventLogDiffReport struct {
	PrimaryTotal    int64
	SecondaryTotal  int64
	PrimaryLoaded   int
	SecondaryLoaded int
	Comparison      Comparison
}

// DiffEventLogs loads and compares detailed event records. Differences are
// informational and therefore never returned as errors.
func DiffEventLogs(
	ctx context.Context,
	primary Loader,
	secondary Loader,
	contractAddressIDs []int64,
	loadLimit int,
	reportLimit int,
) (EventLogDiffReport, error) {
	var report EventLogDiffReport
	if primary == nil || secondary == nil {
		return report, errors.New("evtlog diff: primary and secondary loaders are required")
	}
	if reportLimit <= 0 {
		reportLimit = DefaultDiffReportLimit
	}
	var err error
	report.PrimaryTotal, err = primary.CountEventLogs(ctx, contractAddressIDs)
	if err != nil {
		return report, fmt.Errorf("count primary events: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return report, err
	}
	report.SecondaryTotal, err = secondary.CountEventLogs(ctx, nil)
	if err != nil {
		return report, fmt.Errorf("count secondary events: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return report, err
	}

	primaryEvents, err := primary.LoadDetailedEventLogs(ctx, contractAddressIDs, loadLimit)
	if err != nil {
		return report, fmt.Errorf("query primary events: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return report, err
	}
	secondaryEvents, err := secondary.LoadDetailedEventLogs(ctx, nil, loadLimit)
	if err != nil {
		return report, fmt.Errorf("query secondary events: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return report, err
	}
	report.PrimaryLoaded = len(primaryEvents)
	report.SecondaryLoaded = len(secondaryEvents)
	report.Comparison = CompareDetailedEventLogs(primaryEvents, secondaryEvents, reportLimit)
	return report, nil
}

// FormatEventLogDiffReport returns stable log lines for opsctl db evtlog-diff.
func FormatEventLogDiffReport(report EventLogDiffReport) []string {
	lines := []string{
		fmt.Sprintf(
			"Primary has %d records for our contracts, Secondary has %d records total",
			report.PrimaryTotal,
			report.SecondaryTotal,
		),
		"",
		"=== Loading events from primary (filtered by contract) ===",
		fmt.Sprintf("Loaded %d events from primary", report.PrimaryLoaded),
		"",
		"=== Loading events from secondary ===",
		fmt.Sprintf("Loaded %d events from secondary", report.SecondaryLoaded),
		"",
		"=== Comparing events ===",
	}
	for _, issue := range report.Comparison.Missing {
		lines = append(lines, "ERROR: Secondary missing event: "+issue.Summary)
	}
	if remaining := report.Comparison.MissingTotal - len(report.Comparison.Missing); remaining > 0 {
		lines = append(lines, fmt.Sprintf("... (showing first %d missing; %d more)", len(report.Comparison.Missing), remaining))
	}
	for _, issue := range report.Comparison.Extra {
		lines = append(lines, "ERROR: Secondary has extra event: "+issue.Summary)
	}
	if remaining := report.Comparison.ExtraTotal - len(report.Comparison.Extra); remaining > 0 {
		lines = append(lines, fmt.Sprintf("... (showing first %d extra; %d more)", len(report.Comparison.Extra), remaining))
	}
	for _, mismatch := range report.Comparison.Mismatched {
		lines = append(lines, fmt.Sprintf(
			"ERROR: Mismatch: %s: %s",
			mismatch.Summary,
			formatFields(mismatch.Fields),
		))
	}
	if remaining := report.Comparison.MismatchTotal - len(report.Comparison.Mismatched); remaining > 0 {
		lines = append(lines, fmt.Sprintf("... (showing first %d mismatches; %d more)", len(report.Comparison.Mismatched), remaining))
	}

	lines = append(lines,
		"",
		"=== SUMMARY ===",
		fmt.Sprintf("Primary events (for our contracts): %d", report.PrimaryLoaded),
		fmt.Sprintf("Secondary events: %d", report.SecondaryLoaded),
		fmt.Sprintf("Missing in secondary: %d records", report.Comparison.MissingTotal),
		fmt.Sprintf("Extra in secondary: %d records", report.Comparison.ExtraTotal),
		fmt.Sprintf("Field mismatches: %d records", report.Comparison.MismatchTotal),
	)
	if report.Comparison.Matched() {
		return append(lines, "✓ Databases match perfectly!")
	}
	return append(lines, "✗ Databases have differences!")
}
