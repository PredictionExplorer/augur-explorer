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
- Config in `/etc/rwcg/*.env` (see `.env.example` for every variable).
- Logs in `/var/log/rwcg` (plus the legacy `$HOME/ae_logs` file logs until
  the slog migration completes).

```bash
sudo cp deploy/systemd/*.service deploy/systemd/*.timer /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now rwcg-apiserver@cosmic rwcg-apiserver@rwalk
sudo systemctl enable --now rwcg-etl@cg rwcg-etl@rw
sudo systemctl enable --now rwcg-notibot rwcg-imggen-monitor.timer
```

These replace the legacy `run-loop-*.sh` nohup scripts (still present next to
each command for the transition; delete them once systemd is live).

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
- Prometheus can scrape `METRICS_ADDR/metrics`; alert on
  `rwcg_http_requests_total{status="5xx"}` and request-duration percentiles.
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
