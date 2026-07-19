# ADR-0002: Database layer — pgx driver, goose migrations, incremental sqlc

Date: 2026-07-06. Status: accepted; migration completed 2026-07-08 (the
whole store is context-first, error-returning, pgx-native — see
MODERNIZATION.md §5; the interim sqlc scaffolding was retired, D7).
2026-07-18: the driver replacement is total — the ops toolchain
(`internal/ops/*`, `opsctl`) moved off `database/sql`+`lib/pq` onto
pgx-native queries, `lib/pq` left `go.mod`, and a `depguard` rule keeps it
out (MODERNIZATION.md D16).

## Context

The legacy store (`internal/store`) had three structural problems:

1. It used `lib/pq`, which is in maintenance mode; the community-standard
   driver is `jackc/pgx/v5`.
2. Schema DDL lived in loose `.sql` files applied by hand-run shell scripts,
   with "upgrade" files whose applied-state was untracked.
3. Roughly 500 query helpers called `os.Exit(1)` on any database error,
   killing whichever process was running them (including the API server
   mid-request) and making the layer untestable.

## Decision

- **Migrations:** all schema DDL is expressed as goose migrations in
  `db/migrations`, applied by `goose up` (locally via `make migrate-up`,
  in tests via `internal/testdb`). The legacy DDL files were converted into
  the baseline migrations 00001–00003. On the production database, the
  baseline is recorded as already-applied (see docs/operations.md).
- **Driver:** `pgx/v5` replaces `lib/pq` (via the `stdlib` adapter where the
  `database/sql` interface is still needed).
- **Queries:** new queries are written as SQL in `internal/store/queries/`
  and generated into type-safe Go by sqlc (`make generate`). The legacy
  hand-written query functions are being converted incrementally; dynamic
  queries that sqlc cannot express (whitelisted ORDER BY, optional filters)
  remain hand-written pgx code.
- **Error handling:** store functions return errors; `os.Exit` in library
  code is prohibited (enforced by review and lint configuration). Callers
  decide whether an error is fatal.

## Consequences

The schema has a single source of truth exercised by CI on every run; the
store becomes unit- and integration-testable; database failures degrade
gracefully instead of killing services. The incremental sqlc conversion means
two query styles coexist during the transition — the tradeoff accepted to
avoid a big-bang rewrite of ~200 query functions.
