# Operations

How to run, deploy, and maintain the RWCG backend.

## Local development

Prerequisites: Go 1.26+, Docker (for Postgres and integration tests).

```bash
# Start a migrated database
docker compose up db migrate

# Environment
cp .env.example .env   # fill in RPC_URL at minimum
export $(grep -v '^#' .env | xargs)
# The database can be one DATABASE_URL instead of the PGSQL_* variables.
# Services log structured records to stdout: text by default (dev), JSON
# with LOG_FORMAT=json (production); LOG_LEVEL selects the minimum level.

# Build everything into ./bin
make build

# Run the API server
./bin/apiserver

# Run the indexers (need RPC_URL for Arbitrum)
./bin/cg-etl
./bin/rw-etl
```

Or run the full stack in containers: `docker compose --profile etl up`.

## Testing

```bash
make test              # unit tests, race detector, shuffled
make test-integration  # + testcontainers Postgres (needs Docker)
make lint              # golangci-lint
make vuln              # govulncheck dependency scan
```

Integration tests get a fresh PostgreSQL container with all migrations
applied via `internal/testdb`; they skip automatically when Docker is not
available.

## Database migrations

The schema is owned by goose migrations in `db/migrations`.

```bash
# Apply pending migrations
GOOSE_DBSTRING="postgres://user:pass@host:5432/db" make migrate-up

# Create a new migration
goose -dir db/migrations create add_my_table sql
```

### Adopting migrations on the existing production database

The baseline migrations (00001–00003) describe the schema production already
has. On first rollout, record them as applied without executing them:

```sql
-- run once against the production database
CREATE TABLE IF NOT EXISTS goose_db_version (
    id SERIAL PRIMARY KEY,
    version_id BIGINT NOT NULL,
    is_applied BOOLEAN NOT NULL,
    tstamp TIMESTAMP DEFAULT now()
);
INSERT INTO goose_db_version (version_id, is_applied)
VALUES (0, true), (1, true), (2, true), (3, true);
```

All subsequent migrations are applied normally with `goose up`.

## Production deployment (systemd)

Unit files live in `deploy/systemd/`. Layout on the host:

- Binaries in `/opt/rwcg/bin` (`make build`, then copy `bin/*`).
- Config in `/etc/rwcg/*.env` (verified against the code by
  `internal/config`'s `.env.example` test; set `LOG_FORMAT=json` here).
- Logs go to stdout and journald owns persistence — the legacy
  `$HOME/ae_logs` file logs (`webserver_*.log`, `*_info/_error/_db.log`) are
  gone since the §8.3 slog migration.

```bash
sudo cp deploy/systemd/*.service deploy/systemd/*.timer /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now rwcg-apiserver@cosmic rwcg-apiserver@rwalk
sudo systemctl enable --now rwcg-etl@cg rwcg-etl@rw
sudo systemctl enable --now rwcg-notibot rwcg-imggen-monitor.timer
```

These replace the legacy `run-loop-*.sh` nohup scripts (still present next to
each command for the transition; delete them once systemd is live).

### Reading logs

```bash
journalctl -u rwcg-apiserver@cosmic -f          # follow the API server
journalctl -u rwcg-etl@cg --since "1 hour ago"  # ETL records
journalctl -u rwcg-apiserver@cosmic -o cat | jq 'select(.msg=="request")'
```

Every service emits one JSON record per line (`LOG_FORMAT=json`): access
lines are `msg="request"` with `route`, `status`, `duration_ms` and `ip`
fields; database traces carry `component=db`; startup prints an
`effective configuration` record with secrets redacted.

## Security posture

- Read API is public by design; every route is rate limited per client IP.
- Mutating routes (`ban_bid`, `unban_bid`, ranking `match`) require
  `X-Admin-Key` / `X-Ranking-Admin-Key` and fail closed when the key env var
  is missing. Generate keys with `openssl rand -hex 32`.
- `/metrics` and `/debug/pprof` bind only to `METRICS_ADDR` — keep it on
  localhost or a private interface.
- Secrets come exclusively from environment files; nothing is committed.

## Monitoring

- `srvmonitor` — terminal dashboard checking RPC nodes, DB, web APIs, disk,
  image server; optional WhatsApp/Android alerts (see `cmd/srvmonitor/README.md`).
- `loganomaly` — scans the API access log for 5xx/panics, feeding srvmonitor.
  It understands slog JSON records (production), slog text lines and the
  legacy `[GIN]` lines still present in older capture files. Under systemd,
  export the journal for it on the production host, e.g. a cron entry:

  ```bash
  journalctl -u rwcg-apiserver@cosmic --since -2d -o cat \
    > ~/ae_logs/webserver_cosmic_journal.log \
    && loganomaly -in ~/ae_logs/webserver_cosmic_journal.log
  ```

  The anomalies file and the srvmonitor scp pipeline are unchanged.
- Prometheus can scrape `METRICS_ADDR/metrics`; alert on
  `rwcg_http_requests_total{status="5xx"}` and request-duration percentiles.
  Since the stdlib-router migration the `route` label uses ServeMux syntax
  (`/api/.../{param}` instead of `/api/.../:param`) — update dashboards that
  filter on route values.
- The ETLs (`cg-etl`, `rw-etl`) honor `METRICS_ADDR` too (each process needs
  its own port on a shared host) and expose `rwcg_etl_last_block` (the
  processing watermark — alert when it stops advancing),
  `rwcg_etl_events_total{type}`, `rwcg_etl_batch_duration_seconds`,
  `rwcg_etl_reorgs_total` and `rwcg_etl_batch_failures_total{stage}`.
- ETL failure behavior (since the indexer-engine migration): a failed batch —
  RPC or database — is retried in-process with exponential backoff instead of
  crashing; the process exits non-zero only after 10 consecutive failures
  (systemd restarts it). A restart never skips events: the watermark only
  advances past fully processed blocks, and the engine re-reads it from
  `cg_proc_status`/`rw_proc_status` at startup (rewind it with the ETL
  stopped if you need a manual replay).
- `GET /readyz` returns 503 whenever the database is unreachable — wire it
  into your load balancer health checks.

## Routine tasks

| Task | Command |
|------|---------|
| Verify DB contents against chain | `opsctl db verify`, `opsctl db evtlog-diff` |
| Export/verify RLP archives | `opsctl archive export` / `verify` / `node-fill` |
| Check NFT asset presence | `opsctl assets inventory`, `imggen-monitor` |
| Regenerate thumbnails | `opsctl assets gen-thumbnails` |
| API smoke test | `opsctl smoketest` |
| Contract admin (owner ops) | `cgctl --help` (bid, claim-prize, set-* commands) |
| Social/notification tools | `rwctl --help` (tweet-mints, notify-bot, ...) |
