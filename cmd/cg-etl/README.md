# cg-etl — CosmicGame chain indexer

Indexes every CosmicGame-family contract event into PostgreSQL. The binary is
pure wiring: it connects the RPC node and the database, bootstraps the
contract addresses, runs the startup contract-parameter sync and hands
control to the shared indexing engine.

- Engine (polling loop, batch/retry policy, reorg handling, metrics):
  [`internal/indexer`](../../internal/indexer)
- Event handlers (decode/store pairs per event type):
  [`internal/indexer/cosmicgame`](../../internal/indexer/cosmicgame)
- Schema: goose migrations under [`db/migrations`](../../db/migrations)

## Build and run

```bash
make cg-etl          # builds bin/cg-etl
./bin/cg-etl
```

Required environment variables:

| Variable | Purpose |
|---|---|
| `RPC_URL` | Ethereum JSON-RPC endpoint |
| `PGSQL_HOST`, `PGSQL_USERNAME`, `PGSQL_PASSWORD`, `PGSQL_DATABASE` | PostgreSQL connection |
| `METRICS_ADDR` (optional) | private Prometheus `/metrics` + pprof listener |

Log files are written to `$HOME/ae_logs/`:
`cosmicgame_info.log`, `cosmicgame_error.log`, `cosmicgame_db.log`.

Shutdown: SIGINT/SIGTERM finishes the in-flight batch, persists the
watermark and exits 0. See [docs/operations.md](../../docs/operations.md)
for systemd deployment.
