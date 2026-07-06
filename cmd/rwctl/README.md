# rwctl

`rwctl` is the operator CLI for the RandomWalk NFT, its marketplace, and the
social-media utilities that accompany the project. It consolidates the former
one-off programs from `rwcg/etl/randomwalk/scripts/` and
`rwcg/etl/randomwalk/tools/` into a single binary with one subcommand per
operation.

## Building

```bash
go build ./cmd/rwctl
```

Run `rwctl --help` for the full list of subcommands and `rwctl <cmd> --help`
for per-command usage.

## Environment variables

| Variable | Used by | Description |
|----------|---------|-------------|
| `RPC_URL` | all chain subcommands except `notify-bot`/`tweet-mints` | Ethereum/Arbitrum JSON-RPC endpoint |
| `PKEY_HEX` | transaction subcommands | Signer private key, 64 hex chars, no `0x` prefix. Set in the environment; never pass on the command line |
| `PGSQL_HOST`, `PGSQL_USERNAME`, `PGSQL_PASSWORD`, `PGSQL_DATABASE` | DB-backed subcommands | PostgreSQL connection (empty `PGSQL_HOST` selects the local Unix socket) |
| `AUGUR_ETH_NODE_RPC_URL` | `notify-bot`, `tweet-mints` | Ethereum RPC endpoint (legacy variable name, preserved as-is) |
| `TWITTER_KEYS_FILE` | `notify-bot`, `tweet-mints`, `tweet-reply-image` | Name of the JSON file under `$HOME/configs/` holding `ApiKey`, `ApiSecret`, `TokenKey`, `TokenSecret` |

## Transaction subcommands (write operations)

All transaction subcommands require `RPC_URL` and `PKEY_HEX`. The chain ID is
always fetched from the network, and the suggested gas price is doubled for
faster inclusion. By default they print only `Success. Tx hash = <hash>` or
the error; pass `-i`/`--info` for detailed output (network, account and
transaction sections).

| Subcommand | Description |
|------------|-------------|
| `mint [rwalk_addr] [amount_wei]` | Mint a RandomWalk token by sending `amount_wei` to the contract |
| `approve [rwalk_addr] [operator_addr]` | Set ERC-721 approval-for-all for the operator |
| `transfer [rwalk_addr] [token_id] [new_owner_addr]` | Transfer a token to a new owner |
| `new-offer [BUY\|SELL] [market_addr] [nft_addr] [token_id] [price_wei]` | Create a buy or sell offer on the marketplace |
| `accept-offer [market_addr] [offer_id]` | Accept an existing buy or sell offer |
| `cancel-offer [market_addr] [offer_id]` | Cancel an existing buy or sell offer |
| `set-name [rwalk_addr] [token_id] [new_name]` | Set the display name of a token |

## Query subcommands (read-only)

Require only `RPC_URL`.

| Subcommand | Description |
|------------|-------------|
| `owner-of [rwalk_addr] [token_id]` | Get the owner of a token |
| `price [rwalk_addr]` | Get the current mint price |
| `status [rwalk_addr]` | Read status variables (next token ID, withdrawal info, last minter, base URI) |
| `status-market [market_addr]` | Read marketplace status (number of offers) |
| `token-uri [rwalk_addr] [token_id]` | Get the token URI of a token |
| `withdrawal [rwalk_addr]` | Get the current withdrawal amount (wei) |

## Scan / verification subcommands

Require `RPC_URL`; those marked DB also require the `PGSQL_*` variables.

| Subcommand | Description |
|------------|-------------|
| `scan-mints [randomwalk_addr]` (DB) | Scan the chain for MintEvent logs and report tokens missing from the database |
| `scan-transfers [randomwalk_addr]` | Scan the chain for ERC-721 Transfer logs of one token (`--token-id`, default 3601 as in the legacy script) |
| `verify-owner` (DB) | Verify each token's on-chain owner matches the database |
| `verify-erc20-transfers` (DB) | Verify chain Transfer logs match the `rw_transfer` table (name kept from the legacy script; it checks the NFT's ERC-721 transfers) |

## Statistics and social subcommands

| Subcommand | Description |
|------------|-------------|
| `top-rated` (DB) | Recompute top-100 trader / profit / volume rankings (run once a day) |
| `notify-bot` | Monitor mint/offer/purchase events and floor-price changes, announce them on Twitter (uses `AUGUR_ETH_NODE_RPC_URL`, `TWITTER_KEYS_FILE`, `PGSQL_*`) |
| `tweet-mints` | Same monitor as `notify-bot` (the two legacy tools were identical copies) |
| `tweet-send [api_key] [api_secret] [access_token] [token_secret] [nonce] [message]` | Send a tweet using raw OAuth 1.0a credentials; `message` is optional and defaults to the legacy test message |
| `tweet-reply-image [reply_to_id] [media_file] [message]` | Send a tweet with attached media (image or video) as a reply (uses `TWITTER_KEYS_FILE`) |
| `twitter-auth [--config config.json]` | OAuth 1.0a out-of-band (PIN) flow to provision the token credentials stored in the `TWITTER_KEYS_FILE` config |

## Reference contract addresses (Arbitrum One)

| Contract | Address |
|----------|---------|
| RandomWalk NFT | `0x895a6F444BE4ba9d124F61DF736605792B35D66b` |
| Marketplace | `0x47eF85Dfb775aCE0934fBa9EEd09D22e6eC0Cc08` |
