# RWCG - RandomWalk & CosmicGame Backend

Backend infrastructure for the RandomWalk NFT marketplace and CosmicGame platform on Ethereum/Arbitrum.

## What is RWCG?

**RandomWalk** is an NFT collection with a built-in marketplace for trading tokens.

**CosmicGame** is a blockchain-based game featuring:
- Bidding rounds with ETH and CST (Cosmic Token)
- Prize distribution (main prizes, raffles, endurance champions)
- NFT minting (CosmicSignature tokens)
- Staking rewards for NFT holders
- Charitable donations

## Architecture

```
Blockchain  ──►  ETL Process  ──►  PostgreSQL  ──►  REST API
```

The system indexes blockchain events into a PostgreSQL database and serves the data via a REST API.

## Project Structure

```
rwcg/
├── contracts/      # Smart contract ABI definitions
├── dbs/            # Database layer
├── etl/            # Event indexing (ETL) processes
│   ├── cosmicgame/ # CosmicGame indexer
│   └── randomwalk/ # RandomWalk indexer
├── sql/            # Database schema definitions
├── websrv/         # REST API server
├── tools/          # Utility tools
└── docs/           # Documentation
```

## Quick Start

### Prerequisites

- Go 1.21+
- PostgreSQL 14+
- Ethereum RPC endpoint (Alchemy, Infura, etc.)

### Build

```bash
# Build ETL processes
./etl/cosmicgame/make.sh
./etl/randomwalk/make.sh

# Build web server
cd websrv && go build -o websrv .
```

### Configure

```bash
export PGSQL_USERNAME="user"
export PGSQL_PASSWORD="pass"
export PGSQL_DATABASE="rwcg"
export PGSQL_HOST="localhost:5432"
export RPC_URL="https://arb-mainnet.g.alchemy.com/v2/YOUR_KEY"
```

### Run

```bash
# Start ETL (indexes blockchain events)
./etl/cosmicgame/cg_etl
./etl/randomwalk/rw_etl

# Start API server
./websrv/websrv
```

## Documentation

- **[Backend Technical Reference](docs/BACKEND.md)** - Database schema, ETL process, contract mappings, API endpoints

## License

Proprietary
