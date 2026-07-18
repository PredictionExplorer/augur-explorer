# ADR-0010: Transactional ingestion — one commit per block, watermark included

- Status: Accepted
- Date: 2026-07-18

## Context

The indexing engine (`internal/indexer`) persisted every pipeline step —
block row, transaction row, `evt_log` row, each handler's domain writes and
the processing watermark — as its own autocommitted statement. Correctness
under crashes and mid-batch failures rested entirely on idempotent replay:
`InsertEventLog` deletes-then-inserts, handlers tolerate re-processing, and
the watermark only ever advances past completed block boundaries.

That design is sound but leaves two observable gaps. A failure (or crash)
partway through a block leaves partially applied rows behind — API readers
can see, say, a bid without its later side effects until the retry repairs
the block — and the watermark is written separately from the data it
acknowledges, so the two can transiently disagree. It also makes idempotent
replay *load-bearing*: a future non-idempotent handler would be a
correctness bug waiting for its first mid-block failure.
`docs/MODERNIZATION.md` §7 deferred the fix because it needed "a querier
abstraction across the repo methods and a tx-aware address cache". This ADR
delivers both.

## Decision

### One database transaction per block with events

`Engine.processBatch` groups the fetched logs by block and runs each
block's pipeline — `EnsureBlockExists` (including a chain-split rollback
when one is detected), `EnsureTransactionExists`, `InsertEventLog`, every
handler's store step, and the `Progress.SetLastBlock` watermark write —
inside one `store.InTx` transaction. A block now commits or vanishes
atomically; readers never observe a partial block, and the domain watermark
can never disagree with committed data. The engine's in-memory cursor
simply follows the committed boundary, so the old separate
"partial-progress acknowledgment" write is gone. Event-free tail ranges of
a batch are still acknowledged with a plain watermark update after the
batch — there is no data for them to be atomic with.

The unit is the *block*, not the batch: the watermark's only safe
granularity is a block boundary (a batch can be huge during catch-up, and
`EnsureTransactionExists` performs RPC reads that must not keep a
transaction open for minutes), and per-block commits preserve the engine's
established partial-progress behavior exactly while making each step
atomic. Metrics (`events_total`, the watermark gauge) are recorded only
after a successful commit, so a rolled-back block counts nothing; Begin and
Commit failures surface as the new `commit` stage of
`rwcg_etl_batch_failures_total`.

### Context-scoped querier (`store.Querier`, `Store.InTx`)

`Store.InTx(ctx, fn)` begins a `pgx.Tx` and passes fn a context carrying
it. Every base `Store` method and all 472 domain repo methods resolve their
query surface per call through `Store.Querier(ctx)` — the transaction when
the context carries one *begun by that same Store*, otherwise the shared
pool. That owner check means a foreign Store's methods (multi-Store test
harnesses, tools) can never run on another Store's transaction. Method
signatures are untouched: the handler interface (`Store(ctx, event)`), the
`Progress` adapters and all API read paths work unchanged, and API requests
— whose contexts never carry a transaction — keep running on the pool
exactly as before. Nested `InTx` calls join the open transaction (no
savepoints); the outermost call owns commit and rollback.

The alternative — explicit `pgx.Tx` parameters threaded through the
engine, the registry, 85 handlers and 472 repo methods — was rejected as a
signature-explosion with no added safety: the transaction is single-flight
state owned by one loop, which is exactly what a context value models. The
v2 ranking writes keep their existing explicit-transaction helpers; they
manage a different scope (nonce + match + rating) on the API path.

### Transaction-aware address cache

The Store's bounded address LRU must not learn ids created inside an
uncommitted transaction: after a rollback the id's row no longer exists,
and a poisoned cache entry would feed dangling foreign keys to every later
insert of that address. Inside `InTx`, resolved ids therefore go to a
per-transaction overlay (reads consult the shared LRU first, then the
overlay); `InTx` flushes the overlay into the shared LRU only after a
successful commit, and a rollback simply discards it.

`LookupOrCreateAddress`'s insert also changed from a plain `INSERT` (catch
unique violation, re-read) to `INSERT … ON CONFLICT (addr) DO NOTHING
RETURNING address_id` plus a re-read on conflict. The old pattern was
correct only under autocommit: inside a transaction the unique violation
aborts the transaction, and the recovery re-read itself fails with SQLSTATE
25P02. The new form loses the race gracefully in both modes and preserves
the winner's first-seen block. The regression test synchronizes two
sessions through `pg_stat_activity` so the conflict branch is hit
deterministically.

## Consequences

- A block's rows, its domain writes and its watermark acknowledgment are
  atomic. Idempotent delete-then-insert replay remains in place as defense
  in depth (and still covers operator tools like the evt_log backfill,
  which stay non-transactional), but no correctness claim rests on it for
  the polling loop anymore. The §7 revisit trigger — "if a non-idempotent
  handler ever appears" — is retired.
- Watermark writes happen once per block-with-events instead of once per
  batch. Total commit work *drops*: one fsync per block replaces one per
  statement. The container benchmark (`BenchmarkIngestBlock`) measures the
  transaction wrapper at ~9% per-block overhead versus the raw autocommit
  pipeline (~3.9ms vs ~3.6ms for a three-log block including all container
  round trips) — well under the RPC latency that dominates real ingestion.
- Locks acquired by a block's statements are now held until its commit.
  The two single-writer ETLs share a few hot rows (`last_block`, common
  `block`/`transaction`/`address` rows), so one ETL can briefly wait on the
  other's in-flight block; acquisition order is uniform (block row →
  `last_block` → per-log rows in log order), making deadlocks require two
  processes first-seeing the same two rows in opposite orders — vanishingly
  rare, detected by PostgreSQL within a second, and absorbed by the
  engine's existing retry-with-backoff as one failed batch.
- `NOW()` inside a block's triggers is the transaction timestamp, uniform
  across the block's statements, where autocommit gave each statement its
  own. Nothing golden-pinned depends on wall-clock values (the fixture
  suites are deterministic and stayed byte-identical).
- A chain-split rollback detected inside a block transaction rolls back
  and re-inserts atomically with that block; if the block later fails, the
  reorg metric may count the same split again on retry (the counter tracks
  detections, not distinct splits).
