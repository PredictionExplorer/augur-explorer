//go:build integration

package dbverify_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/dbverify"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

func TestDatabaseComparisonsMatchingAndDivergent(t *testing.T) {
	primaryDB := testdb.New(t)
	secondaryDB := testdb.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	for name, db := range map[string]*testdb.DB{
		"primary":   primaryDB,
		"secondary": secondaryDB,
	} {
		if err := testfixtures.Apply(ctx, db.SQL); err != nil {
			t.Fatalf("applying %s fixtures: %v", name, err)
		}
	}

	contractAddressIDs, err := dbverify.LoadRandomWalkContractAddressIDs(ctx, primaryDB.Pool)
	if err != nil {
		t.Fatalf("loading contract ids: %v", err)
	}
	// db verify intentionally compares the project's primary subset with all
	// rows in a project-only secondary. Reduce the shared full fixture to that
	// shape while retaining the real migrated schema and data relationships.
	if _, err := secondaryDB.Pool.Exec(ctx,
		`DELETE FROM evt_log WHERE NOT (contract_aid = ANY($1))`,
		contractAddressIDs,
	); err != nil {
		t.Fatalf("reducing secondary events: %v", err)
	}
	if _, err := secondaryDB.Pool.Exec(ctx,
		`DELETE FROM transaction t WHERE NOT EXISTS (SELECT 1 FROM evt_log e WHERE e.tx_id = t.id)`,
	); err != nil {
		t.Fatalf("reducing secondary transactions: %v", err)
	}
	if _, err := secondaryDB.Pool.Exec(ctx,
		`DELETE FROM block b WHERE NOT EXISTS (SELECT 1 FROM transaction t WHERE t.block_num = b.block_num)`,
	); err != nil {
		t.Fatalf("reducing secondary blocks: %v", err)
	}

	primary := &dbverify.SQLLoader{DB: primaryDB.Pool}
	secondary := &dbverify.SQLLoader{DB: secondaryDB.Pool}
	matching, err := dbverify.VerifyDatabases(
		ctx,
		primary,
		secondary,
		contractAddressIDs,
		dbverify.DefaultVerifyReportLimit,
	)
	if err != nil {
		t.Fatalf("matching comparison: %v", err)
	}
	if !matching.Matched() {
		t.Fatalf("matching fixtures diverged: %+v", matching)
	}

	var duplicateEventID int64
	if err := primaryDB.Pool.QueryRow(ctx, `
		INSERT INTO evt_log (block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		SELECT block_num, tx_id, contract_aid, topic0_sig, log_index + 1000000, log_rlp
		FROM evt_log
		WHERE contract_aid = ANY($1)
		ORDER BY id
		LIMIT 1
		RETURNING id
	`, contractAddressIDs).Scan(&duplicateEventID); err != nil {
		t.Fatalf("inserting duplicate-RLP event: %v", err)
	}
	duplicateReport, err := dbverify.VerifyDatabases(
		ctx,
		primary,
		secondary,
		contractAddressIDs,
		dbverify.DefaultVerifyReportLimit,
	)
	if err != nil {
		t.Fatalf("comparing duplicate-RLP event: %v", err)
	}
	if duplicateReport.Matched() || duplicateReport.Events.MissingTotal != 1 {
		t.Fatalf("duplicate event multiplicity not detected: %+v", duplicateReport.Events)
	}
	if _, err := primaryDB.Pool.Exec(ctx, `DELETE FROM evt_log WHERE id = $1`, duplicateEventID); err != nil {
		t.Fatalf("removing duplicate-RLP event: %v", err)
	}

	if _, err := secondaryDB.Pool.Exec(ctx, `
		UPDATE transaction
		SET gas_used = gas_used + 1
		WHERE tx_hash = (
			SELECT t.tx_hash
			FROM transaction t
			JOIN evt_log e ON e.tx_id = t.id
			ORDER BY t.tx_hash
			LIMIT 1
		)
	`); err != nil {
		t.Fatalf("diverging transaction: %v", err)
	}
	divergent, err := dbverify.VerifyDatabases(
		ctx,
		primary,
		secondary,
		contractAddressIDs,
		dbverify.DefaultVerifyReportLimit,
	)
	if err != nil {
		t.Fatalf("divergent comparison: %v", err)
	}
	if divergent.Matched() || divergent.Transactions.MismatchTotal != 1 {
		t.Fatalf("transaction divergence not detected: %+v", divergent.Transactions)
	}

	if _, err := secondaryDB.Pool.Exec(ctx, `
		UPDATE evt_log
		SET topic0_sig = 'deadbeef'
		WHERE id = (SELECT MAX(id) FROM evt_log)
	`); err != nil {
		t.Fatalf("diverging event: %v", err)
	}
	diff, err := dbverify.DiffEventLogs(
		ctx,
		primary,
		secondary,
		contractAddressIDs,
		0,
		dbverify.DefaultDiffReportLimit,
	)
	if err != nil {
		t.Fatalf("informational event diff returned an error: %v", err)
	}
	if diff.Comparison.Matched() || diff.Comparison.MismatchTotal != 1 {
		t.Fatalf("event divergence not detected: %+v", diff.Comparison)
	}

	if _, err := primaryDB.Pool.Exec(ctx, `
		CREATE SCHEMA dbverify_bad;
		CREATE TABLE dbverify_bad.block (
			block_num TEXT,
			block_hash TEXT,
			parent_hash TEXT,
			num_tx BIGINT
		);
		INSERT INTO dbverify_bad.block VALUES ('not-a-block-number', 'hash', 'parent', 1);
	`); err != nil {
		t.Fatalf("creating deliberate scan-failure schema: %v", err)
	}
	separator := "?"
	if strings.Contains(primaryDB.ConnString, "?") {
		separator = "&"
	}
	badDB, err := pgxpool.New(ctx, primaryDB.ConnString+separator+"search_path=dbverify_bad")
	if err != nil {
		t.Fatalf("opening deliberate scan-failure database: %v", err)
	}
	t.Cleanup(badDB.Close)
	if _, err := (&dbverify.SQLLoader{DB: badDB}).LoadBlocks(ctx, nil); err == nil ||
		!strings.Contains(err.Error(), "scan block") {
		t.Fatalf("real PostgreSQL scan error = %v, want wrapped block scan failure", err)
	}

	if _, err := secondaryDB.Pool.Exec(ctx, `DROP TABLE evt_log CASCADE`); err != nil {
		t.Fatalf("dropping evt_log for deliberate query failure: %v", err)
	}
	if _, err := secondary.LoadDetailedEventLogs(ctx, nil, 0); err == nil {
		t.Fatal("real PostgreSQL query unexpectedly succeeded after dropping evt_log")
	}
}
