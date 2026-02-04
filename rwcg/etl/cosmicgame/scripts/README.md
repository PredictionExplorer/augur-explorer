# CosmicGame Development Scripts

Scripts for managing CosmicGame transactions and queries, primarily for development and testing purposes.

## Setup

All scripts require the `RPC_URL` environment variable to be set:

```bash
export RPC_URL="http://localhost:8545"  # Local Hardhat/Anvil
# or
export RPC_URL="https://arb-sepolia.g.alchemy.com/v2/YOUR_KEY"  # Arbitrum Sepolia
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
- **Verbose Output**: Consistent, sectioned output format across all scripts

### Common Package Structure

```
common/
├── txutils.go   # Core network/account/transaction utilities
├── format.go    # Formatting helpers (wei to eth, duration, etc.)
├── gas.go       # Gas limit constants by transaction type
└── output.go    # Standardized output helpers
```

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

## Usage Examples

### Make a Bid

```bash
./bid [private_key] [cosmicgame_addr]

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

