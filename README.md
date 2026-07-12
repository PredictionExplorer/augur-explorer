# RWCG Backend

Backend for **CosmicGame** and **RandomWalk** — two blockchain applications on
Arbitrum. It indexes their smart-contract events into PostgreSQL and serves
the data through a JSON API, with notification bots and operational tooling
around the core.

- **CosmicGame** (cosmicsignature.com): a round-based bidding game — ETH/CST
  bids on a Dutch auction, a main prize clock extended by each bid, raffles,
  special prizes (Endurance Champion, Chrono Warrior), NFT minting, staking
  rewards, and charity donations.
- **RandomWalk** (randomwalknft.com): a generative NFT collection with a
  built-in marketplace, on-chain token naming, and an Elo-style community
  ranking game.

## Architecture

```
Arbitrum RPC ──► cg-etl / rw-etl ──► PostgreSQL ──► apiserver ──► frontends
                     (indexers)      (goose-managed)   (JSON API)
```

See [docs/architecture.md](docs/architecture.md) for the full picture,
[docs/openapi.yaml](docs/openapi.yaml) for the frozen v1 contract,
[docs/openapi-v2.yaml](docs/openapi-v2.yaml) for the incremental v2 contract,
and
[docs/adr/](docs/adr/) for the key design decisions.

## Quick start

Prerequisites: Go 1.26+, Docker.

```bash
# 1. Database (starts Postgres and applies migrations)
docker compose up -d db migrate

# 2. Configuration
cp .env.example .env          # set RPC_URL; see the file for all variables
export $(grep -v '^#' .env | xargs)

# 3. Build and run
make build
./bin/apiserver               # JSON API on :8080
./bin/cg-etl                  # CosmicGame indexer
./bin/rw-etl                  # RandomWalk indexer
```

Or everything in containers: `docker compose --profile etl up`.

## Repository layout

| Path | Purpose |
|------|---------|
| `cmd/` | One directory per binary: `apiserver`, `cg-etl`, `rw-etl`, `notibot`, `freezer-scan`, `srvmonitor`, CLIs (`cgctl`, `rwctl`, `opsctl`), ... |
| `internal/` | Shared packages: `api` (handlers), `store` (database), `etl`, `primitives`, `freezer`, `notify`, `testdb` |
| `contracts/` | Generated abigen contract bindings |
| `db/migrations/` | goose schema migrations (source of truth for the schema) |
| `deploy/` | Dockerfile and systemd units |
| `docs/` | Architecture, operations runbook, OpenAPI spec, ADRs |
| `faq-bot/` | Separate Python/Next.js RAG FAQ service (proxied by apiserver) |

## Development

```bash
make generate          # regenerate OpenAPI v2 models/server
make test              # unit tests (race detector)
make test-integration  # + real Postgres via testcontainers
make coverage-check    # enforced global + changed-code coverage policy
make hooks-install     # install the staged pre-commit policy hook
make lint              # golangci-lint
make help              # all targets
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for code standards and
[docs/operations.md](docs/operations.md) for deployment and maintenance.

The codebase is undergoing a tracked, test-first modernization to fully
idiomatic Go with a redesigned v2 API — the roadmap, checklists, and current
status live in [docs/MODERNIZATION.md](docs/MODERNIZATION.md).

## License

Proprietary.
