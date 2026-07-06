# cgctl — CosmicGame operator CLI

`cgctl` consolidates the former one-off scripts from `rwcg/etl/cosmicgame/scripts/`
into a single cobra-based binary with one subcommand per operation: bidding,
prize claiming, owner-only parameter changes, ERC-20/ERC-721 helpers, and
diagnostics.

## Building

```bash
go build -o cgctl ./cmd/cgctl
```

## Setup

Transaction subcommands (those that send transactions) require:

- **RPC_URL** – Ethereum RPC endpoint
- **PKEY_HEX** – Signer private key as 64 hex characters (no `0x` prefix). Never pass the private key on the command line; set it in the environment.
- **GAS_PRICE_MULTIPLIER** – (optional) Multiplier applied to the RPC-suggested gas price. Default is `2.0` so transactions stay above block base fee. Use e.g. `1.5` or `3` to tune.

```bash
export RPC_URL="http://localhost:8545"  # Local Hardhat/Anvil
# or
export RPC_URL="https://arb-sepolia.g.alchemy.com/v2/YOUR_KEY"  # Arbitrum Sepolia

# For transaction subcommands (bid, claim-prize, donate, etc.):
export PKEY_HEX="your_64_char_hex_private_key_no_0x_prefix"
```

Database-backed subcommands (`backfill-dao-evtlog`, `total-tokens`,
`token-seed`) additionally need the PostgreSQL environment:
`PGSQL_HOST`, `PGSQL_USERNAME`, `PGSQL_DATABASE`, `PGSQL_PASSWORD`.

## Architecture

All subcommands share the `internal/ethtx` package, which provides:

- **Network Connection**: chain ID and gas price are always read from the network (never hardcoded)
- **Account Management**: private key parsing, address derivation, nonce fetching
- **Transaction Creation**: standardized transaction options with EIP-155 signing
- **Output**: by default, transaction subcommands print only `Success. Tx hash = <hash>` or the error. Add **`-i`/`--info`** for full detailed output (network, account, round info, etc.). Read-only subcommands always print the detailed sections.

## Subcommands

Run `cgctl --help` (or `cgctl <subcommand> --help`) for details on flags,
positional arguments, and required environment variables.

### Game operations

| Subcommand | Description |
|------------|-------------|
| `bid <addr>` | Make an ETH bid in the current round |
| `claim-prize <addr> [--delay [secs]]` | Claim the main prize; with `--delay` sets `delayDurationBeforeRoundActivation` first (default 60s) |
| `claim-and-set-time-increment <addr> <secs> [delay]` | Set per-bid time increment, claiming/deferring as needed to open the inactive admin window |
| `donate <addr> <amount-wei>` | Donate ETH to the CosmicGame contract |
| `autobid` | Run the automated bidding bot (env-configured) |
| `deploy-erc20` | Deploy a sample ERC-20 token for donation testing |

### Diagnostics / read-only

| Subcommand | Description |
|------------|-------------|
| `info [addr]` | Comprehensive CosmicGame state dump |
| `owner <addr>` | Owner of an Ownable contract |
| `donation-records [addr]` | Dump EthDonationWithInfo records |
| `erc20 balance <token> <user>` | ERC-20 token balance |
| `erc20 allowance <token> <owner> <spender>` | ERC-20 allowance |
| `nft approved <token> <id>` | ERC-721 single-token approval |
| `nft is-approved-for-all <token> <owner> <operator>` | ERC-721 operator approval |
| `nft owner-of <token> <id>` | ERC-721 token owner |
| `total-tokens` | Total CosmicSignature tokens (from DB) |
| `token-seed <id>` | Seed of a CosmicSignature token (from DB) |

### Owner-only setters

| Subcommand | Contract parameter |
|------------|--------------------|
| `set-charity-percentage <addr> <pct>` | `charityEthDonationAmountPercentage` |
| `set-main-prize-percentage <addr> <pct>` | `mainEthPrizeAmountPercentage` |
| `set-raffle-percentage <addr> <pct>` | `raffleTotalEthPrizeAmountForBiddersPercentage` |
| `set-staking-percentage <addr> <pct>` | `cosmicSignatureNftStakingTotalEthRewardAmountPercentage` |
| `set-num-nft-winners <addr> <n>` | `numRaffleCosmicSignatureNftsForBidders` |
| `set-num-raffle-winners <addr> <n>` | `numRaffleEthPrizesForBidders` |
| `set-time-increment <addr> <secs>` | `mainPrizeTimeIncrementInMicroSeconds` |
| `set-initial-duration-divisor <addr> <div>` | `initialDurationUntilMainPrizeDivisor` |
| `set-activation-delay <addr> <secs>` | `delayDurationBeforeRoundActivation` |
| `set-round-activation <addr> <timestamp>` | `roundActivationTime` |

### Token / NFT transactions

| Subcommand | Description |
|------------|-------------|
| `erc20 approve <token> <spender>` | Approve MAX_UINT256 allowance |
| `erc20 revoke <token> <spender>` | Revoke allowance (set to 0) |
| `nft set-name <nft> <id> [name]` | Set a CosmicSignatureNft token name |

### ETL maintenance

| Subcommand | Description |
|------------|-------------|
| `backfill-dao-evtlog [--from-block N] [--to-block N]` | Backfill missing `cosmic_dao` evt_log rows (RPC + PGSQL) |

## Old script → new subcommand mapping

| Legacy script | cgctl subcommand |
|---------------|------------------|
| `autobid` | `autobid` |
| `backfill_dao_evtlog` | `backfill-dao-evtlog` (flags are now `--from-block`/`--to-block`) |
| `bid` | `bid` |
| `cginfo` | `info` |
| `claimprize` | `claim-prize` |
| `claimprize-delay60` | `claim-prize --delay [secs]` (default 60) |
| `claim-and-set-time-increment` | `claim-and-set-time-increment` |
| `deploy-erc20` | `deploy-erc20` |
| `donate` | `donate` |
| `dwi-records` | `donation-records` |
| `erc20bal` | `erc20 balance` |
| `erc20allowance` | `erc20 allowance` |
| `erc20approve` | `erc20 approve` |
| `erc20revoke` | `erc20 revoke` |
| `approved` | `nft approved` |
| `isapproved4all` | `nft is-approved-for-all` |
| `tokownerof` | `nft owner-of` |
| `set_token_name` | `nft set-name` |
| `owner` | `owner` |
| `set_charity_percentage` | `set-charity-percentage` |
| `set_main_prize_percentage` | `set-main-prize-percentage` |
| `set_num_nft_winners` | `set-num-nft-winners` |
| `set_num_raffle_winners` | `set-num-raffle-winners` |
| `set_raffle_percentage` | `set-raffle-percentage` |
| `set_staking_percentage` | `set-staking-percentage` |
| `set-initial-duration-divisor` | `set-initial-duration-divisor` |
| `set-time-increment` | `set-time-increment` |
| `setactivation` | `set-activation-delay` |
| `setroundactivation` | `set-round-activation` |
| `img_upload total_tokens` | `total-tokens` |
| `img_upload token_seed <id>` | `token-seed <id>` |

The `-i` flag of the old transaction scripts is preserved as `-i`/`--info`.

## Dev scripts

The `dev-scripts/` directory contains the shell/Hardhat helpers that used to
live next to the legacy scripts, rewritten to invoke `cgctl` (build it first
and put it on `PATH`, or set `CGCTL=/path/to/cgctl`):

- **claim-all.sh** – try to claim the main prize with every standard Hardhat dev account. (Historical note: this used to call a `claimnft` binary whose source was lost; `cgctl claim-prize` is the closest replacement.)
- **claim-and-configure.sh** – claim with a round-activation delay, then set `initialDurationUntilMainPrizeDivisor` during the inactive window.
- **claim-and-set-time-increment.sh** – wrapper around `cgctl claim-and-set-time-increment`.
- **backfill-dao-evtlog.sh** – source a config env file (default `~/configs/cg-prod.env`, override with `CG_ENV`) and run `cgctl backfill-dao-evtlog`.
- **imgcheck.sh** – check/regenerate/upload NFT PNG+MP4 artifacts using DB lookups via `cgctl total-tokens` / `cgctl token-seed`.
- **imgcheck_api.sh** – same idea but fetches token seeds from the public API (`GET /api/cosmicgame/cst/list/all/0/100000`); needs `curl`, `jq`, and SSH access, no DB. Prefer this one when the API is available.
- **deploy-and-populate.sh** + **deploy-dev-and-samp.js** + **populate.js** + **rpc-helpers.js** – deploy the CosmicSignatureGame stack with dev-friendly time settings plus two Samp ERC-20s from a Cosmic-Signature checkout, then populate the chain for testing.

### Dev: deploy + populate

```bash
./dev-scripts/deploy-and-populate.sh --network localhost /path/to/Cosmic-Signature
```

This script:

1. **Deploys** the full CosmicSignatureGame stack from the Cosmic-Signature repo with **dev-friendly time settings** (e.g. 1 minute claim timeout, 1 minute per bid, 1 hour withdraw timeout).
2. **Deploys** two Samp ERC20 contracts.
3. **Runs** `populate.js` with `CADDR`, `TSAMP1`, `TSAMP2` set so the chain is populated for testing.

Requirements: a Hardhat node (or Anvil) on `localhost:8545` (or set `NETWORK`),
and a Cosmic-Signature repo path (must contain `scripts/Deploy.js` and
`contracts/tests/Samp.sol`).

#### Round-activation delay window

After deployment there is a short time interval before the first round becomes
active. During this window the contract allows admin setters (e.g.
`setCharityAddress`, `setNumRaffleEthPrizesForBidders`, timeouts for the next
round). The deploy script sets a **5-second** delay so that:

1. Deploy sets round activation to `now + 5` seconds (and `delayDurationBeforeRoundActivation(5)` for later rounds).
2. Populate runs **admin config during that window** (when it detects the round is not yet active), then waits until the round is active (Hardhat: `evm_increaseTime` + `evm_mine`; real RPC: sleep).
3. Donate, bids, and the rest of the script then run with the round active.

This matches the pattern used by **claim-and-configure.sh** and
**claim-and-set-time-increment.sh**: claim (or deploy) leaves a short delay
before the next round activates; you run config during that window.

## Usage examples

```bash
# Make a bid (quiet output)
cgctl bid 0x5FbDB2315678afecb367f032d93F642f64180aa3
# Success. Tx hash = 0x...

# Make a bid with detailed output
cgctl bid -i 0x5FbDB2315678afecb367f032d93F642f64180aa3

# Check a token balance
cgctl erc20 balance <token_addr> <user_addr>

# Full contract state
cgctl info <cosmicgame_addr>
```

Detailed (`-i`) output uses the same sectioned format as the old scripts:

```
==================== NETWORK INFO ====================
RPC URL             = http://localhost:8545
Chain ID            = 31337
...
==================== TRANSACTION RESULT ====================
Status              = SUBMITTED
Tx Hash             = 0x...
```

## Key design principles

1. **Network-Derived Values**: chain ID and gas price are always fetched from the network
2. **Consistent Output**: all subcommands use the same output format with clear sections
3. **Verbose Information**: show all relevant context (balances, current state, etc.)
4. **Error Handling**: errors are wrapped with context and reported once via cobra
5. **Balance Checks**: verify sufficient funds before submitting transactions
6. **Ownership Warnings**: warn when calling admin functions as non-owner

## Gas limits

The `internal/ethtx` package defines standard gas limits:

| Constant | Value | Use Case |
|----------|-------|----------|
| `GasLimitERC20Approve` | 100,000 | Token approvals |
| `GasLimitBid` | 500,000 | CosmicGame bids |
| `GasLimitClaimPrize` | 3,500,000 | Prize claims (V2 needs ~3M) |
| `GasLimitDonate` | 300,000 | Donations |
| `GasLimitAdminCall` | 100,000 | Admin setters |

The `autobid` bot uses its own larger limits (1M for bids, 5M for claims).

## ERC721 NFTs minted per round (production contracts)

Source: `production/` (Cosmic-Signature contracts). Verified against
[cosmic-signature-game-prizes.md](https://github.com/PredictionExplorer/Cosmic-Signature/blob/main/docs/cosmic-signature-game-prizes.md)
(Group 2 prizes, CS NFT rows). Each round mints Cosmic Signature (CS) NFTs from
these prize types:

| Source | NFTs per round |
|--------|-----------------|
| Main prize | 1 |
| Last CST bidder | 0 or 1 (only if someone bid with CST) |
| Endurance champion | 1 |
| Chrono warrior | 1 |
| Raffle (bidders) | `numRaffleCosmicSignatureNftsForBidders` (default **10**) |
| Raffle (Random Walk stakers) | 0 or `numRaffleCosmicSignatureNftsForRandomWalkNftStakers` (default **10**) |

With **default constants** (`CosmicSignatureConstants.sol`):

- **Minimum**: 1 + 0 + 1 + 1 + 10 + 0 = **13** (no CST bidder, no RW staker raffle winners)
- **Maximum**: 1 + 1 + 1 + 1 + 10 + 10 = **24** (CST bidder present, RW staker raffle full)

If the owner sets both raffle NFT counts to 0 (`setNumRaffleCosmicSignatureNftsForBidders(0)`
and `setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(0)` when the round
is inactive), the range is **3** (no CST bidder) to **4** (with CST bidder).
The ETL field `cg_round_stats.total_nfts_minted` stores the actual count per round.

## SQL (no schema names)

Use table names only; do not qualify with a schema (e.g. avoid
`cosmicgame.cg_nft_donation`). Rely on the DB search path.

Example – list all unclaimed donated ERC721s (round number, token ID, contract address):

```sql
SELECT
  d.round_num,
  d.token_id,
  a.addr AS contract_address
FROM cg_nft_donation d
LEFT JOIN cg_donated_nft_claimed c
  ON d.round_num = c.round_num AND d.idx = c.idx
JOIN address a ON d.token_aid = a.address_id
WHERE c.idx IS NULL
ORDER BY d.round_num, d.idx;
```
