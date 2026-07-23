# opsctl — data-operations utilities for the RWCG backend

`opsctl` bundles the former standalone tools from `rwcg/tools/` (plus the
notibot `verify-token-imgs` script) into a single cobra CLI with one
subcommand per utility.

Build from the repo root:

```sh
go build ./cmd/opsctl
```

Every subcommand supports `--help`. Flags use the standard cobra `--flag`
syntax (the old single-file tools used Go `flag` single-dash syntax; names and
defaults are unchanged).

`cmd/opsctl` contains only Cobra, environment and connection wiring. The
context-aware archive, assets, CST scan, DB verification, smoke-test and
transaction-collector engines live under `internal/ops`; all three node log
scanners share the adaptive range/retry implementation in
`internal/indexer/logscan`.

## Command overview

| Command | Replaces | Purpose |
|---|---|---|
| `opsctl archive export` | `rwcg/tools/archive_export.go` | Copy evt_log/transaction/block rows into `arch_*` tables of another DB |
| `opsctl archive verify` | `rwcg/tools/arch_verify.go` | Consistency-check live tables vs `arch_*` tables |
| `opsctl archive node-fill` | `rwcg/tools/arch_node_fill.go` | Backfill `arch_*` tables from an Ethereum node |
| `opsctl db verify` | `rwcg/tools/db_verify.go` | Compare evt_log/transaction/block between two DBs |
| `opsctl db evtlog-diff` | `rwcg/tools/evtlog_diff.go` | Field-level evt_log diff between two DBs |
| `opsctl tx-collector run` | `rwcg/tools/transaction_collector.go` | Back up tx + receipt RLP blobs for contract activity |
| `opsctl tx-collector verify` | `rwcg/tools/transaction_collector_verify.go` | Check RLP backup blobs against evt_log |
| `opsctl assets inventory` | `rwcg/tools/asset_inventory.go` | Inventory Cosmic Signature assets on disk vs DB seeds |
| `opsctl assets gen-thumbnails` | `rwcg/tools/gen_thumbnails.go` | Generate WebP thumbnails for per-seed packages |
| `opsctl assets verify-token-images` | `cmd/notibot/scripts/verify-token-imgs.go` | Check RandomWalk token image URLs for HTTP 403 |
| `opsctl smoketest` | `rwcg/tools/api_smoketest.go` | Validate the canonical v2 API, frozen v1 compatibility surface, or D6-safe operational probes |
| `opsctl scan cst-auction-len` | `rwcg/tools/scan_cst_auclen_chg_div.go` | Scan chain for `CstDutchAuctionDurationChangeDivisorChanged` events |

`dev-scripts/check_cs_images.sh` (formerly `rwcg/tools/check_cs_images.sh`)
remains a shell script: it checks the *public image URLs* served by the
frontend web server (via `psql` + `curl`), complementing the on-disk checks of
`assets inventory`.

## Common environment variables

| Variable | Used by | Meaning |
|---|---|---|
| `RPC_URL` | `archive node-fill`, `scan cst-auction-len` | Ethereum/Arbitrum JSON-RPC endpoint |
| `PGSQL_HOST` (`host[:port]`), `PGSQL_USERNAME`, `PGSQL_DATABASE`, `PGSQL_PASSWORD` | `tx-collector verify`, `assets inventory`, `assets gen-thumbnails`, `assets verify-token-images`, `smoketest` | PostgreSQL connection (used when no `--db` flag is given; `smoketest --suite=operational` does not need it) |
| `NFT_ASSETS_ROOT` | `assets gen-thumbnails` | Asset root; default base dir is `$NFT_ASSETS_ROOT/new/cosmicsignature` |
| `API_BASE` | `smoketest` | Overrides the API base URL |
| `HTTP_PORT` | `smoketest` | websrv port; base defaults to `http://127.0.0.1:$HTTP_PORT` (9090 when unset) |

## archive

Tools that maintain the `arch_evtlog` / `arch_tx` / `arch_block` archive
tables (see `db/layer1/archive_tables.sql`). `--project both` processes
cosmicgame first, then randomwalk; either order is safe because every tool
resumes per project.

### archive export

Copies evt_log rows for a project's contracts from a source (production) DB
into the destination's `arch_evtlog`, then fills `arch_tx` and `arch_block`
for every referenced transaction and block. `arch_evtlog` rows are keyed by
`(tx_hash, log_index)`; incremental export resumes from the per-contract
minimum `MAX(evt_id)` already archived.

| Flag | Default | Meaning |
|---|---|---|
| `--src` | (required) | Source DB connection string (production) |
| `--dst` | (required) | Destination DB connection string (dev) |
| `--project` | (required) | `randomwalk` \| `cosmicgame` \| `both` |

```sh
opsctl archive export --project both --src 'postgres://...' --dst 'postgres://...'
```

### archive verify

Archival consistency checks: live `evt_log` / `transaction` / `block` vs the
`arch_*` tables in the same database. Exits non-zero on blocking mismatches.

| Flag | Default | Meaning |
|---|---|---|
| `--db` | (required) | Connection string (same DB holds live + `arch_*` tables) |
| `--project` | (required) | `randomwalk` \| `cosmicgame` \| `both` |
| `--strict-arch-block-metadata` | `false` | Also fail on `num_tx`/`ts`/`cash_flow` drift (default: only hashes must match) |
| `--strict-arch-tx-num-logs` | `false` | Also fail on `arch_tx.num_logs` drift |

```sh
opsctl archive verify --project both --db 'postgres://user:pass@host:5432/dbname?sslmode=disable'
```

### archive node-fill

Backfills `arch_evtlog` / `arch_tx` / `arch_block` from an Ethereum node via
`FilterLogs`, inserting only rows missing from `arch_evtlog`. Requires
`RPC_URL` and `arch_evtlog` keyed by `(tx_hash, log_index)`.
Retries also repair partial dependencies and reconcile stale project rows
against canonical blocks, including reorg replacements with no matching logs.
The command completes best-effort scanning but exits non-zero if any row-level
RPC or database error remains unresolved.

| Flag | Default | Meaning |
|---|---|---|
| `--db` | (required) | PostgreSQL connection string |
| `--project` | (required) | `randomwalk` \| `cosmicgame` \| `both` |
| `--from` | `0` | Start block (0 = auto-detect from address/evt_log metadata) |
| `--start-block` | `0` | Same as `--from`; overrides `--from` when both are set |
| `--to` | `0` | End block inclusive (0 = chain head) |
| `--batch` | `100000` | FilterLogs block range size (halved on RPC errors) |
| `--dry-run` | `false` | Scan and report only; do not insert |

```sh
RPC_URL=https://... opsctl archive node-fill --project both --db 'postgres://...' --start-block 9292392
```

## db

Database-to-database comparison tools for RandomWalk contract data (contract
set is read from `rw_contracts` on the primary).

### db verify

Compares `evt_log`, `transaction` and `block` between a primary (gold
standard) and a secondary DB; the secondary is expected to contain only the
project's data. Exits non-zero on any mismatch.

| Flag | Default | Meaning |
|---|---|---|
| `--primary` | (required) | Primary DB connection string (production) |
| `--secondary` | (required) | Secondary DB connection string (new rwcg) |

### db evtlog-diff

Field-level `evt_log` diff keyed by `log_rlp` content: reports missing/extra
records and per-field mismatches. Informational — always exits 0 unless a
query fails.

| Flag | Default | Meaning |
|---|---|---|
| `--primary` | (required) | Primary DB connection string (gold standard) |
| `--secondary` | (required) | Secondary DB connection string (to verify) |
| `--limit` | `0` | Limit comparison to first N records (0 = all) |

## tx-collector

Backs up RLP-encoded transactions and receipts for contract activity, and
verifies those backups against `evt_log`. Both subcommands read the same JSON
config — see `tx-collector.example.json`:

```json
{
  "rpc_url": "https://...",
  "output_dir": "/var/backups/cosmicgame-transactions",
  "start_block": 0,
  "contracts": { "cosmic_game_addr": "0x...", ... }
}
```

Blobs are stored as `<output_dir>/<block_num>/<tx_hash>_tx.rlp` and
`..._receipt.rlp`. New directories use mode `0750` and blobs use `0640`;
writes use a unique same-directory temporary file followed by atomic rename,
with file and directory `fsync`, so interrupted or concurrent runs never
expose or acknowledge a partial blob. Existing blobs are decoded and their
transaction identity checked before they are skipped; unrepaired
encoding/filesystem failures make the command exit non-zero.

### tx-collector run

| Flag | Default | Meaning |
|---|---|---|
| `--config` | (required) | Path to the JSON config |
| `--start-block` | `0` | Override config `start_block` |
| `--to` | `0` | End block inclusive (0 = chain head) |
| `--batch` | `100000` | FilterLogs block range size |

```sh
opsctl tx-collector run --config ~/configs/transaction-collector.cosmicgame.json
```

### tx-collector verify

For each `evt_log` row (scoped by config contracts and start block) checks the
tx hash, log index and RLP bytes against the backup blobs, and cross-checks
disk/SQL coverage in both directions. Exits non-zero on failures.

| Flag | Default | Meaning |
|---|---|---|
| `--config` | (required) | The same collector JSON config |
| `--db` | `""` | PostgreSQL URL (default: built from `PGSQL_*` env) |
| `--start-block` | `0` | Override config `start_block` for the evt_log filter |
| `--max-report` | `50` | Max detailed mismatch lines per category (0 = unlimited) |

## assets

### assets inventory

Fetches every minted Cosmic Signature token seed from `cg_mint_event` and
checks the local filesystem for each token's image, preview and video,
accepting the current per-seed package layout plus legacy/flat fallbacks
(mirroring the web server's resolution priority). Run it from the asset
directory or pass `--base`.

| Flag | Default | Meaning |
|---|---|---|
| `--db` | `""` | Connection string (default: built from `PGSQL_*` env) |
| `--base` | `.` | Base directory holding the per-seed assets |
| `--schema` | `public` | Schema holding `cg_mint_event` |
| `--missing-only` | `false` | Only list tokens with at least one missing asset |
| `--all` | `false` | List every token with its status |

### assets gen-thumbnails

Generates `thumb_card.webp` (640px) and `thumb_micro.webp` (160px) next to
each on-disk `0x<seed>/image.png`. Requires ImageMagick (7 `magick` or 6
`convert`, auto-detected). By default only missing/stale thumbnails are
regenerated; exits non-zero when any generation fails.

| Flag | Default | Meaning |
|---|---|---|
| `--db` | `""` | Connection string (default: built from `PGSQL_*` env) |
| `--base` | `""` | Asset directory (default `$NFT_ASSETS_ROOT/new/cosmicsignature`) |
| `--schema` | `public` | Schema holding `cg_mint_event` |
| `--force` | `false` | Regenerate every thumbnail |
| `--magick` | `""` | ImageMagick binary override |
| `--min-age` | `10` | Skip tokens whose `image.png` changed within N seconds |

Cron (every minute), wrapped in `flock` so runs never overlap:

```cron
* * * * * flock -n /tmp/cs_thumbs.lock nice -n 19 ionice -c3 bash -lc 'source ~/configs/cosmic-api-config.env && opsctl assets gen-thumbnails >> ~/ae_logs/thumbnails.log 2>&1'
```

### assets verify-token-images

Fetches every minted RandomWalk token image from the public image server
(`https://api.randomwalknft.com:1443/images/randomwalk`) and reports tokens
answered with HTTP 403 (a known RandomWalk webserver bug for early token ids).
It checks HTTP status without writing temporary image files. No flags; the DB
connection comes from the `PGSQL_*` environment variables.

## smoketest

The default `v2` suite discovers every GET operation from the embedded
canonical OpenAPI document (currently 97), binds production-shaped path/query
values from PostgreSQL, requires HTTP 200, validates each response against the
operation schema, and fails if a v2 response carries `Deprecation` or `Sunset`.
Requests are paced at 50 rps so the test cannot trip the API's own global
limiter.

| `--suite` | Database | Purpose |
|---|---|---|
| `v2` (default) | required | Full canonical v2 read-surface and dependency-degradation check |
| `v1` | required | Frozen 142-request v1 compatibility regression |
| `both` | required | v2 followed by v1, with per-suite and aggregate failures |
| `operational` | not used | `/healthz`, `/readyz`, `/version`, and three stable DB-backed v2 reads for uptime/Compose |

The operational set deliberately excludes contract-state cache resources:
those return a correct 503 when RPC-backed state is unavailable and the full
v2 suite must report that degradation. The v1 suite retains its legacy
top-level `error` / numeric `"status":0` detection until v1 is retired.

```sh
source ~/configs/cg-prod.env
opsctl smoketest
opsctl smoketest --suite=both
API_BASE=http://127.0.0.1:9090 opsctl smoketest --suite=operational

# Whole-stack, D6-safe local smoke:
docker compose --profile smoke up --build --abort-on-container-exit \
  --exit-code-from smoketest smoketest
```

## scan

### scan cst-auction-len

Scans the chain for `CstDutchAuctionDurationChangeDivisorChanged(uint256)`
events on the CosmicSignatureGame proxy (topic0
`0xacbc6b...f3e6f`) and, with `--db`, reports which on-chain occurrences are
missing from the `cg_adm_cst_auclen_chg_div` table. Requires `RPC_URL`.

| Flag | Default | Meaning |
|---|---|---|
| `--contract` | `0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2` | CosmicSignatureGame proxy address |
| `--from-block` | `455767589` | Start block (contract deployment) |
| `--to-block` | `0` | End block (0 = latest) |
| `--batch` | `100000` | FilterLogs block range size |
| `--db` | `""` | Optional postgres conn string for the DB cross-check |

```sh
RPC_URL=https://arb1.arbitrum.io/rpc opsctl scan cst-auction-len --db 'postgres://cgprod@localhost/cgprod?sslmode=disable'
```

## dev-scripts

`dev-scripts/check_cs_images.sh` verifies that the Cosmic Signature token
image URLs served by the frontend (`thumb_micro.webp`, `thumb_card.webp`,
full `.png`) return HTTP 200 for every minted token. Requires `psql` and
`curl`; DB connection from `PGSQL_*` (optionally sourced via `--env`). See the
header comment for options (`--base-url`, `--limit`, `--jobs`, `--fail-only`,
`--schema`, `--timeout`).
