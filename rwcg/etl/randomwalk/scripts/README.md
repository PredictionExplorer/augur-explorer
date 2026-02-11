# RandomWalk Scripts

Scripts for interacting with RandomWalk NFT contracts and verifying data. Used for development, testing, and data checks.

## Setup

Transaction scripts (those that send transactions) require:

- **RPC_URL** – Ethereum RPC endpoint

```bash
export RPC_URL="http://localhost:8545"
# or
export RPC_URL="https://arb-sepolia.g.alchemy.com/v2/YOUR_KEY"
```

For scripts that send transactions (mint, approve, transfer, accept_offer, cancel_offer, new_offer, setname), also set:

- **PKEY_HEX** – Signer private key as 64 hex characters (no `0x` prefix). Set in the environment; do not pass on the command line.

```bash
export PKEY_HEX="your_64_char_hex_private_key_no_0x_prefix"
```

Database-backed scripts (scan_rwmints, scan_transfers, verify_owner, verify_erc20_transfers) also need PostgreSQL env vars: **PGSQL_HOST**, **PGSQL_USERNAME**, **PGSQL_PASSWORD**, **PGSQL_DATABASE**.

## Building

From this directory:

```bash
# Build all scripts
./make.sh

# Clean binaries
./make-clean.sh
```

To build a single script:

```bash
go build -o mint mint.go
```

## Scripts Overview

### Transaction Scripts (Write Operations)

| Script | Description |
|--------|-------------|
| `mint.go` | Mint a new RandomWalk token |
| `approve.go` | Set ERC721 approval for all (SetApprovalForAll) |
| `transfer.go` | Transfer a token to another address |
| `new_offer.go` | Create a buy or sell offer on the market |
| `accept_offer.go` | Accept an existing offer (buy or sell) |
| `cancel_offer.go` | Cancel an offer |
| `setname.go` | Set the name for a token |
| `withdrawal.go` | Withdraw proceeds from the market |

### Query Scripts (Read-Only)

| Script | Description |
|--------|-------------|
| `ownerof.go` | Get owner of a token by token ID |
| `price.go` | Get current price / token info |
| `status.go` | Get token or contract status |
| `statusmkt.go` | Get market status |
| `tokenuri.go` | Get token URI for a token ID |

### Verification / Scan Scripts

| Script | Description |
|--------|-------------|
| `scan_rwmints.go` | Scan chain for RWMint events and check DB presence |
| `scan_transfers.go` | Scan chain for Transfer events (e.g. for a given token) |
| `verify_owner.go` | Verify each token’s on-chain owner matches the DB |
| `verify_erc20_transfers.go` | Verify chain transfers match `rw_transfer` table |

## Usage Examples

### Mint a token

```bash
export RPC_URL="..." PKEY_HEX="..."
./mint
```

### Create and accept an offer

```bash
./new_offer    # create offer (see script for args)
./accept_offer # accept by offer ID (see script for args)
```

### Query token info

```bash
./ownerof [token_id]
./price [token_id]
./tokenuri [token_id]
```

### Scan / verify (require DB env)

```bash
./scan_rwmints [randomwalk_contract_addr]
./scan_transfers [randomwalk_contract_addr]
./verify_owner
./verify_erc20_transfers
```

## Build and Clean

- **Build all:** `./make.sh` – builds every script that compiles in this directory.
- **Clean all:** `./make-clean.sh` – removes all built binaries (and is invoked by the parent `make-clean.sh` hierarchy).
