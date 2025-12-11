# RWCG - RandomWalk & CosmicGame Web Server

This is an extracted subproject containing the web server functionality for RandomWalk and CosmicGame.

## Structure

```
rwcg/
├── README.md
└── websrv/
    ├── main.go                           # Entry point
    ├── server.go                         # Server configuration and RWCGServer struct
    │
    ├── api/
    │   ├── common/                       # Shared utilities (package common)
    │   │   ├── context.go                # Shared server context
    │   │   ├── requests.go               # HTTP request utilities
    │   │   ├── api_utils.go              # Parameter parsing utilities
    │   │   └── jsbuild.go                # JavaScript chart building
    │   │
    │   ├── randomwalk/                   # RandomWalk handlers (package randomwalk)
    │   │   ├── randomwalk.go             # Package init and route registration
    │   │   ├── api_randomwalk.go         # JSON API handlers
    │   │   └── srv_randomwalk.go         # HTML page handlers
    │   │
    │   └── cosmicgame/                   # CosmicGame handlers (package cosmicgame)
    │       ├── cosmicgame.go             # Package init and route registration
    │       ├── api_cosmicgame.go         # Main JSON API handlers
    │       ├── api_cosmicgame_donations.go
    │       ├── api_cosmicgame_staking_cst.go
    │       ├── api_cosmicgame_staking_rwalk.go
    │       ├── api_cosmicgame_staking_both.go
    │       ├── srv_cosmicgame.go         # Main HTML handlers
    │       ├── srv_cosmicgame_donations.go
    │       ├── srv_cosmicgame_staking_cst.go
    │       ├── srv_cosmicgame_staking_rwalk.go
    │       └── srv_cosmicgame_staking_both.go
    │
    ├── html/                             # Static assets
    │   ├── imgs/                         # Images
    │   └── res/                          # CSS, JS resources
    │
    └── templates/                        # HTML templates
        ├── cosmicgame/                   # 116 CosmicGame templates
        ├── randomwalk/                   # 24 RandomWalk templates
        └── misc/                         # Common templates (error.html)
```

## Package Organization

- **`main`** (websrv/) - Entry point, server configuration
- **`common`** (api/common/) - Shared utilities and context
- **`randomwalk`** (api/randomwalk/) - RandomWalk business logic
- **`cosmicgame`** (api/cosmicgame/) - CosmicGame business logic

## Building

```bash
cd rwcg/websrv
go build .
```

## Running

### Environment Variables

```bash
# Required: RPC URL for Ethereum node (Arbitrum)
export RPC_URL="https://arb-mainnet.g.alchemy.com/v2/YOUR_KEY"

# Required: Database connection (primary database)
export EXTRACTOR_USERNAME="your_db_user"
export EXTRACTOR_PASSWORD="your_db_password"
export EXTRACTOR_DATABASE="your_db_name"
export EXTRACTOR_HOST="localhost:5432"

# Optional: Secondary Arbitrum database connection
export ARB_USERNAME="your_arb_user"
export ARB_PASSWORD="your_arb_password"
export ARB_DATABASE="your_arb_db"
export ARB_HOST="localhost:5432"

# Required: HTTP port
export HTTP_PORT="8080"

# Optional: HTTPS hostname for TLS
export HTTPS_HOSTNAME="api.example.com:443"
```

### Start the Server

```bash
cd rwcg/websrv
./websrv
```

## API Endpoints

### RandomWalk API (`/api/rwalk/...`)
- `/api/rwalk/current_offers/:rwalk_addr/:market_addr/:order_by`
- `/api/rwalk/floor_price/:rwalk_addr/:market_addr`
- `/api/rwalk/tokens/list/sequential/:rwalk_addr`
- `/api/rwalk/tokens/info/:rwalk_addr/:token_id`
- `/api/rwalk/trading/history/:market_addr/:offset/:limit`
- `/api/rwalk/statistics/by_token/:rwalk_addr`
- And more...

### CosmicGame API (`/api/cosmicgame/...`)
- `/api/cosmicgame/statistics/dashboard`
- `/api/cosmicgame/rounds/list/:offset/:limit`
- `/api/cosmicgame/bid/list/all/:offset/:limit`
- `/api/cosmicgame/cst/list/all/:offset/:limit`
- `/api/cosmicgame/user/info/:user_addr`
- `/api/cosmicgame/staking/cst/staked_tokens/all`
- And more...

### RandomWalk HTML (`/black/rwalk/...`)
- `/black/rwalk/` - Index page
- `/black/rwalk/tokens/info/:rwalk_addr/:token_id`
- `/black/rwalk/trading/history/:market_addr/:offset/:limit`
- And more...

### CosmicGame HTML (`/black/cosmicgame/...`)
- `/black/cosmicgame` - Dashboard
- `/black/cosmicgame/rounds/list`
- `/black/cosmicgame/bid/list/all`
- `/black/cosmicgame/user/info/:user_addr`
- And more...

## Dependencies

> **Note**: The import paths currently reference the parent repository. These will be updated
> when the ETL and database code are extracted in Phase 2 and 3 of the extraction plan.

Current external dependencies (to be extracted):
- `primitives` - Common data types
- `primitives/cosmicgame` - CosmicGame-specific types
- `dbs` - Database layer
- `dbs/cosmicgame` - CosmicGame database operations
- `contracts` - Smart contract bindings

## Extraction Status

- ✅ Phase 1: Web server code (complete)
- ⏳ Phase 2: ETL code (pending)
- ⏳ Phase 3: Database code (pending)

Once Phase 2 and 3 are complete, this project will have its own `go.mod` and can be 
completely separated from the parent repository.
