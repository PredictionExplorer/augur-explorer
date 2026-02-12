# RWCG Backend Documentation

This document describes the backend architecture for the RWCG (RandomWalk + CosmicGame) platform, covering the database schema, ETL processes, contract-to-database mapping, and API layer.

## Table of Contents

1. [Architecture Overview](#architecture-overview)
2. [Database Schema](#database-schema)
3. [ETL Process](#etl-process)
4. [Contract Event Mapping](#contract-event-mapping)
5. [API Layer](#api-layer)

---

## Architecture Overview

The backend consists of three main components:

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   Blockchain    │────▶│   ETL Process   │────▶│   PostgreSQL    │
│   (Ethereum)    │     │   (Go service)  │     │   Database      │
└─────────────────┘     └─────────────────┘     └────────┬────────┘
                                                         │
                                                         ▼
                                                ┌─────────────────┐
                                                │   Web Server    │
                                                │   (Gin API)     │
                                                └─────────────────┘
```

### Components

| Component | Location | Purpose |
|-----------|----------|---------|
| ETL - CosmicGame | `etl/cosmicgame/` | Processes CosmicGame contract events |
| ETL - RandomWalk | `etl/randomwalk/` | Processes RandomWalk contract events |
| Web Server | `websrv/` | REST API serving data to frontend |
| Database Layer | `dbs/` | PostgreSQL connection and queries |
| SQL Schema | `sql/` | Database table definitions |
| Contracts | `contracts/` | ABI definitions for smart contracts |

---

## Database Schema

The database is organized into three schemas:

### 1. Layer1 Schema (Blockchain Infrastructure)

Core tables for storing raw blockchain data.

#### Core Tables

| Table | Purpose | Key Fields |
|-------|---------|------------|
| `block` | Block headers | `block_num`, `ts`, `block_hash`, `parent_hash` |
| `address` | Ethereum addresses | `address_id`, `addr` (unique) |
| `transaction` | Transaction records | `id`, `block_num`, `from_aid`, `to_aid`, `tx_hash` |
| `tx_input` | Transaction input data | `tx_id`, `data` (hex-encoded) |
| `evt_log` | Raw event logs (RLP-encoded) | `id`, `tx_id`, `topic0_sig`, `log_rlp` |
| `evt_topic` | Indexed event parameters | `evtlog_id`, `pos`, `value` |

#### Metadata Tables

| Table | Purpose |
|-------|---------|
| `abi_funcs` | Function signature registry |
| `abi_events` | Event signature registry |
| `last_block` | Last processed block number |
| `contract_addresses` | Platform contract addresses |
| `chain_reorg` | Chain reorganization tracking |

### 2. CosmicGame Schema

Game event and statistics tables. Most tables reference `evt_log(id)` for traceability.

#### Prize Tables

| Table | Purpose | Key Fields |
|-------|---------|------------|
| `cg_prize_claim` | Main prize claims | `round_num`, `winner_aid`, `amount`, `token_id` |
| `cg_prize` | Unified prize records | `round_num`, `winner_index`, `ptype` |
| `cg_raffle_nft_prize` | Raffle NFT prizes | `round_num`, `winner_aid`, `token_id`, `winner_idx` |
| `cg_raffle_eth_prize` | Raffle ETH prizes | `round_num`, `winner_aid`, `amount` |
| `cg_endurance_prize` | Endurance champion prizes | `round_num`, `erc721_token_id`, `erc20_amount` |
| `cg_lastcst_prize` | Last CST bidder prizes | `round_num`, `erc721_token_id`, `erc20_amount` |
| `cg_chrono_warrior_prize` | Chrono warrior prizes | `round_num`, `eth_amount`, `cst_amount`, `nft_id` |

#### Bidding Tables

| Table | Purpose | Key Fields |
|-------|---------|------------|
| `cg_bid` | Bid events | `bidder_aid`, `round_num`, `bid_type`, `eth_price`, `cst_price`, `msg` |
| `cg_first_bid` | First bid in each round | `round_num`, `start_ts` |

#### Donation Tables

| Table | Purpose | Key Fields |
|-------|---------|------------|
| `cg_eth_donated` | ETH donations | `donor_aid`, `round_num`, `amount` |
| `cg_eth_donated_wi` | ETH donations with info | `record_id`, `donor_aid`, `amount` |
| `cg_donation_json` | Donation metadata (JSON) | `record_id`, `data` |
| `cg_donation_received` | Charity wallet receipts | `donor_aid`, `amount` |
| `cg_donation_sent` | Charity wallet disbursements | `charity_aid`, `amount` |
| `cg_erc20_donation` | ERC20 token donations | `round_num`, `donor_aid`, `token_aid`, `amount` |
| `cg_nft_donation` | NFT donations | `round_num`, `donor_aid`, `token_aid`, `token_id` |

#### NFT & Token Tables

| Table | Purpose | Key Fields |
|-------|---------|------------|
| `cg_mint_event` | NFT mint events | `owner_aid`, `token_id`, `round_num`, `seed` |
| `cg_token_name` | NFT name changes | `token_id`, `token_name` |
| `cg_erc721_transfer` | ERC721 transfers | `token_id`, `from_aid`, `to_aid`, `otype` |
| `cg_erc20_transfer` | ERC20 transfers | `value`, `from_aid`, `to_aid`, `otype` |
| `cg_costok_owner` | CosmicToken balances | `owner_aid`, `cur_balance` |

#### Staking Tables (CST)

| Table | Purpose | Key Fields |
|-------|---------|------------|
| `cg_nft_staked_cst` | CST NFT stake events | `staker_aid`, `token_id`, `round_num` |
| `cg_nft_unstaked_cst` | CST NFT unstake events | `staker_aid`, `token_id`, `reward` |
| `cg_staking_eth_deposit` | ETH deposits for staking rewards | `round_num`, `deposit_amount`, `amount_per_token` |
| `cg_st_reward` | Individual staking rewards | `staker_aid`, `token_id`, `deposit_id`, `reward` |
| `cg_staker_cst` | Staker statistics | `staker_aid`, `total_tokens_staked`, `total_reward` |
| `cg_staked_token_cst` | Currently staked tokens | `staker_aid`, `token_id`, `stake_action_id` |

#### Staking Tables (RandomWalk)

| Table | Purpose | Key Fields |
|-------|---------|------------|
| `cg_nft_staked_rwalk` | RWalk NFT stake events | `staker_aid`, `token_id` |
| `cg_nft_unstaked_rwalk` | RWalk NFT unstake events | `staker_aid`, `token_id` |
| `cg_staker_rwalk` | RWalk staker statistics | `staker_aid`, `total_tokens_staked` |
| `cg_staked_token_rwalk` | Currently staked RWalk tokens | `staker_aid`, `token_id` |

#### Statistics Tables

| Table | Purpose | Key Fields |
|-------|---------|------------|
| `cg_round_stats` | Per-round statistics | `round_num`, `total_bids`, `total_eth_in_bids` |
| `cg_glob_stats` | Global statistics | `num_bids`, `num_wins`, `total_raffle_eth_deposits` |
| `cg_bidder` | Bidder statistics | `bidder_aid`, `num_bids`, `max_bid` |
| `cg_winner` | Winner statistics | `winner_aid`, `prizes_count`, `prizes_sum` |
| `cg_donor` | Donor statistics | `donor_aid`, `count_donations`, `total_eth_donated` |

#### Admin/Configuration Tables

Tables prefixed with `cg_adm_*` track configuration changes:
- `cg_adm_charity_pcent`, `cg_adm_main_prize_pcent`, `cg_adm_stake_pcent`
- `cg_adm_raffle_pcent`, `cg_adm_chrono_pcent`
- `cg_adm_charity_wallet`, `cg_adm_rwalk_addr`, `cg_adm_prizes_wallet_addr`
- `cg_adm_time_inc`, `cg_adm_price_inc`, `cg_adm_acttime`

### 3. RandomWalk Schema

NFT marketplace tables.

#### Core Tables

| Table | Purpose | Key Fields |
|-------|---------|------------|
| `rw_token` | Token information | `token_id`, `cur_owner_aid`, `seed_hex`, `last_price` |
| `rw_new_offer` | Marketplace offers | `offer_id`, `token_id`, `seller_aid`, `buyer_aid`, `price`, `active` |
| `rw_item_bought` | Purchase events | `offer_id`, `seller_aid`, `buyer_aid` |
| `rw_offer_canceled` | Cancelled offers | `offer_id` |
| `rw_withdrawal` | Withdrawal events | `token_id`, `aid`, `amount` |
| `rw_token_name` | Token name changes | `token_id`, `new_name` |
| `rw_mint_evt` | Mint events | `token_id`, `owner_aid`, `seed`, `price` |
| `rw_transfer` | Token transfers | `token_id`, `from_aid`, `to_aid`, `otype` |

#### Statistics Tables

| Table | Purpose | Key Fields |
|-------|---------|------------|
| `rw_stats` | Per-contract statistics | `rwalk_aid`, `total_vol`, `total_num_trades` |
| `rw_mkt_stats` | Marketplace statistics | `contract_aid`, `total_vol`, `total_buy_orders` |
| `rw_user_stats` | User statistics | `user_aid`, `total_vol`, `total_profit` |
| `rw_uranks` | User rankings | `aid`, `top_profit`, `top_trades`, `top_volume` |

### Key Relationships

```
evt_log (raw events)
    │
    ├──▶ cg_bid, cg_prize_claim, cg_mint_event, ... (CosmicGame events)
    │
    └──▶ rw_new_offer, rw_mint_evt, rw_transfer, ... (RandomWalk events)

address (address registry)
    │
    └──▶ *_aid foreign keys in all tables
```

---

## ETL Process

The ETL (Extract, Transform, Load) process synchronizes blockchain events to the database.

### Processing Flow

```
┌──────────────────────────────────────────────────────────────────┐
│                        ETL Main Loop                              │
├──────────────────────────────────────────────────────────────────┤
│  1. Get last processed block from cg_proc_status / rw_proc_status │
│  2. Fetch current block number from blockchain                    │
│  3. Calculate block range (adaptive batch sizing: 1K-1M blocks)   │
│  4. Fetch events using eth_getLogs (FilterLogs)                   │
│  5. For each event:                                               │
│     a. Ensure block exists in database                            │
│     b. Ensure transaction exists in database                      │
│     c. Insert raw event log (RLP-encoded)                         │
│     d. Decode and process event (domain-specific)                 │
│  6. Update processing status                                      │
│  7. Repeat                                                        │
└──────────────────────────────────────────────────────────────────┘
```

### Event Fetching

**Location**: `etl/common/eventfetcher.go`

```go
FetchEvents(client, fromBlock, toBlock, contracts []common.Address)
```

- Uses `ethclient.FilterLogs()` to fetch events
- Filters by contract addresses and block ranges
- Adaptive batch sizing based on event density

### Raw Event Storage

**Location**: `dbs/blockchain_insert.go`

Events are first stored in raw form:

1. Extract `topic0_sig` (first 4 bytes of `Topics[0]`)
2. RLP-encode the log
3. Look up/create contract address ID
4. Insert into `evt_log` table

```sql
INSERT INTO evt_log (block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
VALUES ($1, $2, $3, $4, $5, $6)
```

### Event Processing

**Location**: 
- `etl/cosmicgame/cosmicgame.go` → `process_single_event()`
- `etl/randomwalk/rwalk.go` → `process_single_event()`

Process flow:
1. Retrieve event log from `evt_log` by ID
2. RLP-decode back to `types.Log`
3. Match event signature (`Topics[0]`) to handler
4. Decode event data using contract ABI
5. Extract indexed parameters from `Topics[]`
6. Insert into domain-specific table

### Running the ETL

```bash
# Build
./etl/cosmicgame/make.sh
./etl/randomwalk/make.sh

# Run (requires environment variables)
./etl/cosmicgame/cg_etl
./etl/randomwalk/rw_etl
```

Required environment variables:
- `PGSQL_USERNAME`, `PGSQL_PASSWORD`, `PGSQL_DATABASE`, `PGSQL_HOST`
- `RPC_URL` (Ethereum RPC endpoint)

---

## Contract Event Mapping

This section describes how Solidity contract events map to database tables.

### Event Identification

Events are identified by the first 4 bytes of `keccak256(event_signature)`.

### CosmicGame Event Mappings

| Solidity Event | Database Table | Description |
|----------------|----------------|-------------|
| `BidPlaced(uint256 indexed roundNum, address indexed bidder, int256 indexed randomWalkTokenId, ...)` | `cg_bid` | User placed a bid |
| `MainPrizeClaimed(uint256 indexed roundNum, address indexed winner, ...)` | `cg_prize_claim` | Main prize claimed |
| `EthDonated(address indexed donor, uint256 indexed roundNum, uint256 amount)` | `cg_eth_donated` | ETH donated |
| `EthDonatedWithInfo(address indexed donor, uint256 indexed roundNum, uint256 recordId, uint256 amount)` | `cg_eth_donated_wi` | ETH donated with metadata |
| `NftMinted(address indexed owner, uint256 tokenId, uint256 roundNum, bytes32 seed)` | `cg_mint_event` | CosmicSignature NFT minted |
| `RaffleWinnerPrizePaid(uint256 indexed roundNum, address indexed winner, uint256 tokenId, ...)` | `cg_raffle_nft_prize` | Raffle NFT prize awarded |
| `RaffleWinnerBidderEthPrizeAllocated(uint256 indexed roundNum, address indexed winner, uint256 amount)` | `cg_raffle_eth_prize` | Raffle ETH prize awarded |
| `EnduranceChampionPrizePaid(uint256 indexed roundNum, address indexed winner, ...)` | `cg_endurance_prize` | Endurance champion prize |
| `LastCstBidderPrizePaid(uint256 indexed roundNum, address indexed winner, ...)` | `cg_lastcst_prize` | Last CST bidder prize |
| `ChronoWarriorPrizePaid(uint256 indexed roundNum, address indexed winner, ...)` | `cg_chrono_warrior_prize` | Chrono warrior prize |
| `NftStaked(uint256 indexed actionId, address indexed staker, uint256 tokenId, ...)` | `cg_nft_staked_cst` | CST NFT staked |
| `NftUnstaked(uint256 indexed actionId, address indexed staker, uint256 tokenId, ...)` | `cg_nft_unstaked_cst` | CST NFT unstaked |
| `EthDepositReceived(uint256 indexed roundNum, uint256 depositId, ...)` | `cg_staking_eth_deposit` | ETH deposited for staking rewards |
| `Transfer(address from, address to, uint256 tokenId)` | `cg_erc721_transfer` | ERC721 token transfer |
| `Transfer(address from, address to, uint256 value)` | `cg_erc20_transfer` | ERC20 token transfer |

### RandomWalk Event Mappings

| Solidity Event | Database Table | Description |
|----------------|----------------|-------------|
| `NewOffer(uint256 indexed offerId, uint256 indexed tokenId, address seller, ...)` | `rw_new_offer` | New marketplace offer |
| `ItemBought(uint256 indexed offerId, address indexed seller, address indexed buyer)` | `rw_item_bought` | Offer accepted |
| `OfferCanceled(uint256 indexed offerId)` | `rw_offer_canceled` | Offer cancelled |
| `WithdrawalEvent(uint256 indexed tokenId, address indexed to, uint256 amount)` | `rw_withdrawal` | Funds withdrawn |
| `TokenNameEvent(uint256 indexed tokenId, string newName)` | `rw_token_name` | Token renamed |
| `MintEvent(uint256 indexed tokenId, address indexed owner, bytes32 seed, uint256 price)` | `rw_mint_evt` | Token minted |
| `Transfer(address from, address to, uint256 tokenId)` | `rw_transfer` | Token transferred |

### Event Decoding Example

```go
// BidPlaced event processing (cosmicgame.go)
func process_bid_event(log types.Log) {
    // 1. Decode non-indexed data
    var evt struct {
        PaidEthPrice   *big.Int
        PaidCstPrice   *big.Int
        MainPrizeTime  *big.Int
        Message        string
    }
    cosmic_game_abi.UnpackIntoInterface(&evt, "BidPlaced", log.Data)
    
    // 2. Extract indexed parameters from Topics
    roundNum := new(big.Int).SetBytes(log.Topics[1].Bytes())
    bidder := common.BytesToAddress(log.Topics[2].Bytes()[12:])
    rwalkTokenId := new(big.Int).SetBytes(log.Topics[3].Bytes())
    
    // 3. Insert into database
    Insert_bid_event(evtlog_id, block_num, tx_id, timestamp,
        contract_aid, bidder_aid, rwalkTokenId, roundNum,
        evt.PaidEthPrice, evt.PaidCstPrice, evt.MainPrizeTime, evt.Message)
}
```

### Address Normalization

All Ethereum addresses are normalized to address IDs:

```go
aid := Lookup_or_create_address(addr)  // Returns BIGINT ID
```

This reduces storage and enables efficient joins.

---

## API Layer

The web server provides a REST API for accessing data.

### Server Architecture

**Location**: `websrv/`

- Framework: Gin
- Database: PostgreSQL via `dbs.SQLStorage`
- Response format: JSON

### Database Connection

```go
// Environment variables (PGSQL_ prefix only)
PGSQL_USERNAME, PGSQL_PASSWORD, PGSQL_DATABASE, PGSQL_HOST
```

### Response Format

```json
{
  "status": 1,      // 1 = success, 0 = error
  "error": "",      // Error message if status = 0
  // ... data fields
}
```

### API Endpoints

#### Statistics & Dashboard

| Endpoint | Description |
|----------|-------------|
| `GET /api/cosmicgame/statistics/dashboard` | Dashboard statistics |
| `GET /api/cosmicgame/statistics/counters` | Record counters |
| `GET /api/cosmicgame/statistics/unique/bidders` | Unique bidder count |
| `GET /api/cosmicgame/statistics/unique/winners` | Unique winner count |
| `GET /api/cosmicgame/statistics/unique/donors` | Unique donor count |

#### Rounds

| Endpoint | Description |
|----------|-------------|
| `GET /api/cosmicgame/rounds/list/:offset/:limit` | List prize rounds |
| `GET /api/cosmicgame/rounds/info/:prize_num` | Round details |
| `GET /api/cosmicgame/rounds/current/time` | Current round prize time |

#### Bids

| Endpoint | Description |
|----------|-------------|
| `GET /api/cosmicgame/bid/list/all/:offset/:limit` | All bids |
| `GET /api/cosmicgame/bid/info/:evtlog_id` | Bid details |
| `GET /api/cosmicgame/bid/list/by_round/:round_num/:sort/:offset/:limit` | Bids by round |
| `GET /api/cosmicgame/bid/cst_price` | Current CST bid price |
| `GET /api/cosmicgame/bid/eth_price` | Current ETH bid price |

#### Prizes

| Endpoint | Description |
|----------|-------------|
| `GET /api/cosmicgame/prizes/history/global/:offset/:limit` | Global prize history |
| `GET /api/cosmicgame/prizes/history/by_user/:user_addr/:offset/:limit` | User prize history |
| `GET /api/cosmicgame/prizes/eth/raffle/global` | Raffle ETH deposits |
| `GET /api/cosmicgame/prizes/eth/unclaimed/by_user/:user_addr/:offset/:limit` | User's unclaimed prizes |

#### User

| Endpoint | Description |
|----------|-------------|
| `GET /api/cosmicgame/user/info/:user_addr` | User information |
| `GET /api/cosmicgame/user/balances/:user_addr` | User balances |
| `GET /api/cosmicgame/user/notif_red_box/:user_addr` | User notifications |

#### CosmicSignature (CST) NFTs

| Endpoint | Description |
|----------|-------------|
| `GET /api/cosmicgame/cst/list/all/:offset/:limit` | All CST tokens |
| `GET /api/cosmicgame/cst/list/by_user/:user_addr/:offset/:limit` | User's CST tokens |
| `GET /api/cosmicgame/cst/info/:token_id` | Token details |
| `GET /api/cosmicgame/cst/transfers/all/:token_id/:offset/:limit` | Token transfer history |

#### Cosmic Token (CT - ERC20)

| Endpoint | Description |
|----------|-------------|
| `GET /api/cosmicgame/ct/balances` | All CT balances |
| `GET /api/cosmicgame/ct/statistics` | CT statistics |
| `GET /api/cosmicgame/ct/summary/by_user/:user_addr` | User's CT summary |

#### Donations

| Endpoint | Description |
|----------|-------------|
| `GET /api/cosmicgame/donations/eth/simple/list/:offset/:limit` | ETH donations |
| `GET /api/cosmicgame/donations/nft/list/:offset/:limit` | NFT donations |
| `GET /api/cosmicgame/donations/erc20/global/:offset/:limit` | ERC20 donations |
| `GET /api/cosmicgame/donations/charity/deposits` | Charity deposits |

#### Staking

| Endpoint | Description |
|----------|-------------|
| `GET /api/cosmicgame/staking/cst/staked_tokens/all` | All staked CST tokens |
| `GET /api/cosmicgame/staking/cst/staked_tokens/by_user/:user_addr` | User's staked tokens |
| `GET /api/cosmicgame/staking/cst/rewards/global` | Global staking rewards |
| `GET /api/cosmicgame/staking/cst/rewards/to_claim/by_user/:user_addr` | User's claimable rewards |
| `GET /api/cosmicgame/staking/rwalk/staked_tokens/all` | All staked RWalk tokens |

#### RandomWalk API

| Endpoint | Description |
|----------|-------------|
| `GET /api/rwalk/tokens/list/sequential/:rwalk_addr` | Token list |
| `GET /api/rwalk/tokens/info/:rwalk_addr/:token_id` | Token info |
| `GET /api/rwalk/current_offers/:rwalk_addr/:market_addr/:order_by` | Current offers |
| `GET /api/rwalk/floor_price/:rwalk_addr/:market_addr` | Floor price |
| `GET /api/rwalk/trading/history/:market_addr/:offset/:limit` | Trading history |
| `GET /api/rwalk/statistics/by_token/:rwalk_addr` | Token statistics |
| `GET /api/rwalk/user/info/:user_aid/:rwalk_addr` | User information |

### Running the Web Server

```bash
# Build
cd websrv && go build -o websrv .

# Run
./websrv
```

Required environment variables:
- `PGSQL_USERNAME`, `PGSQL_PASSWORD`, `PGSQL_DATABASE`, `PGSQL_HOST`
- `RPC_URL` (for live blockchain queries)
- `ENABLE_COSMICGAME` (optional, defaults to true)

---

## Directory Structure

```
rwcg/
├── contracts/              # Contract ABI definitions
│   ├── cosmicgame/         # CosmicGame ABIs
│   └── randomwalk/         # RandomWalk ABIs
├── dbs/                    # Database layer
│   ├── db.go               # Connection management
│   ├── blockchain_insert.go # Raw event insertion
│   ├── cosmicgame/         # CosmicGame queries
│   └── randomwalk/         # RandomWalk queries
├── etl/                    # ETL processes
│   ├── common/             # Shared ETL utilities
│   ├── cosmicgame/         # CosmicGame ETL
│   └── randomwalk/         # RandomWalk ETL
├── sql/                    # Database schema
│   ├── layer1/             # Blockchain tables
│   ├── cosmicgame/         # CosmicGame tables
│   └── randomwalk/         # RandomWalk tables
└── websrv/                 # Web API server
    ├── cosmicgame/         # CosmicGame handlers
    └── rwalk/              # RandomWalk handlers
```

---

## Configuration

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `PGSQL_USERNAME` | PostgreSQL username | Yes |
| `PGSQL_PASSWORD` | PostgreSQL password | Yes |
| `PGSQL_DATABASE` | PostgreSQL database name | Yes |
| `PGSQL_HOST` | PostgreSQL host | Yes |
| `RPC_URL` | Ethereum RPC endpoint | Yes |
| `ENABLE_COSMICGAME` | Enable CosmicGame API | No (default: true) |

### Database Initialization

Tables are created via SQL scripts in the `sql/` directory:

```bash
# Layer1 schema
psql -f sql/layer1/tables.sql
psql -f sql/layer1/indices.sql

# CosmicGame schema
psql -f sql/cosmicgame/tables-cosmicgame.sql
psql -f sql/cosmicgame/triggers-cosmicgame.sql
psql -f sql/cosmicgame/indices_cosmicgame.sql

# RandomWalk schema
psql -f sql/randomwalk/tables_randomwalk.sql
psql -f sql/randomwalk/triggers.sql
psql -f sql/randomwalk/indices_rwalk.sql
```
