# CosmicGame Development Scripts

Scripts for managing CosmicGame transactions and queries, primarily for development and testing purposes.

## Setup

Transaction scripts (those that send transactions) require:

- **RPC_URL** – Ethereum RPC endpoint
- **PKEY_HEX** – Signer private key as 64 hex characters (no `0x` prefix). Never pass the private key on the command line; set it in the environment.

```bash
export RPC_URL="http://localhost:8545"  # Local Hardhat/Anvil
# or
export RPC_URL="https://arb-sepolia.g.alchemy.com/v2/YOUR_KEY"  # Arbitrum Sepolia

# For transaction scripts (bid, claimprize, donate, etc.):
export PKEY_HEX="your_64_char_hex_private_key_no_0x_prefix"
```

## Building

From the scripts directory:

```bash
# Build all scripts
go build ./...

# Build specific script
go build -o bid ./bid.go
```

## Architecture

All scripts use the shared `common` package which provides:

- **Network Connection**: ChainID and gas price are always read from the network (never hardcoded)
- **Account Management**: Private key parsing, address derivation, nonce fetching
- **Transaction Creation**: Standardized transaction options with EIP-155 signing
- **Output**: By default, transaction scripts print only `Success. Tx hash = <hash>` or the error. Add **-i** to the command line for full detailed output (network, account, round info, etc.).

### Common Package Structure

```
common/
├── txutils.go   # Core network/account/transaction utilities
├── format.go    # Formatting helpers (wei to eth, duration, etc.)
├── gas.go       # Gas limit constants by transaction type
└── output.go    # Standardized output helpers
```

## Dev: Deploy + Populate (single bash script)

For local development with fast execution (short time increments and timeouts), use:

```bash
./deploy-and-populate.sh --network localhost /path/to/Cosmic-Signature
```

This script:

1. **Deploys** the full CosmicSignatureGame stack from the Cosmic-Signature repo with **dev-friendly time settings** (e.g. 1 minute claim timeout, 1 minute per bid, 1 hour withdraw timeout).
2. **Deploys** two Samp ERC20 contracts.
3. **Runs** `populate.js` with `CADDR`, `TSAMP1`, `TSAMP2` set so the chain is populated for testing.

Requirements:

- Hardhat node (or Anvil) running on `localhost:8545` (or set `NETWORK`).
- Cosmic-Signature repo path (must contain `scripts/Deploy.js` and `contracts/tests/Samp.sol`).

Node scripts involved:

- **deploy-dev-and-samp.js** – Deploys via `basicDeployment`, sets short time params, deploys Samp; prints `CADDR`, `TSAMP1`, `TSAMP2`.
- **deploy-samp.js** – Standalone deploy of two Samp contracts only (if you already have a game deployed).
- **populate.js** – Populates the already-deployed game (requires `CADDR` and optionally `TSAMP1`/`TSAMP2`).

### Round-activation delay window

After deployment there is a short time interval before the first round becomes active. During this window the contract allows admin setters (e.g. `setCharityAddress`, `setNumRaffleEthPrizesForBidders`, timeouts for the next round). The deploy script sets a **5-second** delay so that:

1. Deploy sets round activation to `now + 5` seconds (and `delayDurationBeforeRoundActivation(5)` for later rounds).
2. Populate runs **admin config during that window** (when it detects the round is not yet active), then waits until the round is active (Hardhat: `evm_increaseTime` + `evm_mine`; real RPC: sleep).
3. Donate, bids, and the rest of the script then run with the round active.

This matches the pattern used by **claim-and-configure.sh** and **claim-and-set-time-increment.sh**: claim (or deploy) leaves a short delay before the next round activates; you run config during that window; a small timeout allows `sleep()` so the next transaction succeeds.

## Scripts Overview

### Transaction Scripts (Write Operations)

| Script | Description |
|--------|-------------|
| `bid.go` | Make a bid at the current CosmicGame round |
| `claimprize.go` | Claim the main prize |
| `donate.go` | Donate ETH to the CosmicGame contract |
| `sendeth.go` | Send ETH to an address |
| `erc20approve.go` | Approve MAX_UINT256 ERC20 allowance |
| `erc20revoke.go` | Revoke ERC20 allowance (set to 0) |
| `setactivation.go` | Set delay before round activation (admin) |
| `setroundactivation.go` | Set round activation time (admin) |
| `set-time-increment.go` | Set time increment per bid (admin) |
| `set-initial-duration-divisor.go` | Set initial duration divisor (admin) |

### Query Scripts (Read-Only Operations)

| Script | Description |
|--------|-------------|
| `cginfo.go` | Comprehensive CosmicGame state dump |
| `owner.go` | Get contract owner |
| `erc20bal.go` | Get ERC20 token balance |
| `erc20allowance.go` | Get ERC20 token allowance |
| `isapproved4all.go` | Check ERC721 operator approval |
| `tokownerof.go` | Get owner of ERC721 token |

### NFT image / upload (no DB)

| Script | Description |
|--------|-------------|
| `imgcheck_api.sh` | Fetch all CosmicSignature token seeds from the **public API**, check which PNG/MP4 artifacts are missing on a remote host (`DST_DIR`), regenerate only missing via `run.py` (or `three_body_problem` fallback), and upload them. No SQL or `img_upload` helper. |
| `imgcheck.sh` | Legacy: same idea but uses SQL via `img_upload` (total_tokens, token_seed). Prefer `imgcheck_api.sh` when the API is available. |

**imgcheck_api.sh** usage (run on the machine where `run.py` and `./pics`/`./vids` live):

```bash
./imgcheck_api.sh <ssh-host> [api_base_url]
# Example:
./imgcheck_api.sh my.server.com
./imgcheck_api.sh my.server.com http://161.129.67.42:8353
```

- Requires `curl`, `jq`, and SSH key-based access to `frontend@<ssh-host>`.
- Destination directory on the remote host: `DST_DIR="/home/frontend/nft-assets/new/cosmicsignature"`.
- API used: `GET /api/cosmicgame/cst/list/all/0/100000` to obtain all token seeds.

## Usage Examples

### Make a Bid

```bash
export PKEY_HEX="your_64_char_hex_private_key"
./bid [cosmicgame_addr]

# Example output:
# ==================== NETWORK INFO ====================
# RPC URL             = http://localhost:8545
# Chain ID            = 31337
# Gas Price (gwei)    = 1.0000
# ...
# ==================== ROUND INFO ====================
# Round Number        = 5
# Next Bid Price      = 0.001000000000000000 ETH
# ...
# ==================== TRANSACTION RESULT ====================
# Status              = SUBMITTED
# Tx Hash             = 0x...
```

### Check Token Balance

```bash
./erc20bal [token_addr] [user_addr]
```

### Get Contract Info

```bash
./cginfo [cosmicgame_addr]
```

## Output Format

All scripts produce verbose, sectioned output:

```
==================== NETWORK INFO ====================
RPC URL             = http://localhost:8545
Chain ID            = 31337
Gas Price (gwei)    = 1.0000

==================== ACCOUNT INFO ====================
Address             = 0x...
Nonce               = 42
Balance (ETH)       = 10.000000000000000000

==================== TRANSACTION RESULT ====================
Status              = SUBMITTED
Tx Hash             = 0x...
```

## Key Design Principles

1. **Network-Derived Values**: ChainID and gas price are always fetched from the network
2. **Consistent Output**: All scripts use the same output format with clear sections
3. **Verbose Information**: Show all relevant context (balances, current state, etc.)
4. **Error Handling**: Clear error messages with context
5. **Balance Checks**: Verify sufficient funds before submitting transactions
6. **Ownership Warnings**: Warn when calling admin functions as non-owner

## Gas Limits

The common package defines standard gas limits:

| Constant | Value | Use Case |
|----------|-------|----------|
| `GasLimitSimpleTransfer` | 21,000 | ETH transfers |
| `GasLimitERC20Approve` | 100,000 | Token approvals |
| `GasLimitBid` | 500,000 | CosmicGame bids |
| `GasLimitClaimPrize` | 2,000,000 | Prize claims |
| `GasLimitDonate` | 300,000 | Donations |
| `GasLimitAdminCall` | 100,000 | Admin setters |

