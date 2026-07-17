# API v2 migration guide

API v2 decomposes the frozen v1 mega-responses into typed resources. V1
remains available while consumers migrate and is removed only after the
sunset gates in [ADR-0005](adr/0005-api-v2.md) are met.

As of 2026-07-16 the v1 surface is formally deprecated: every v1 operation
is flagged `deprecated: true` in [openapi.yaml](openapi.yaml) and answers
with an RFC 9745 `Deprecation` header plus a `Link` to this guide. The
RFC 8594 `Sunset` header will announce the removal date once the gates are
met (`V1_SUNSET_AT`).

## CosmicGame dashboard

Replace `GET /api/cosmicgame/statistics/dashboard` with the following
resource composition:

- Current round identity, pools, countdown, last bidder and round statistics:
  `GET /api/v2/cosmicgame/rounds/current`.
- Global totals and staking summaries:
  `GET /api/v2/cosmicgame/statistics`.
- Canonical record counters:
  `GET /api/v2/cosmicgame/statistics/counters`.
- Ranked bidder, winner, donor and staker arrays:
  `GET /api/v2/cosmicgame/statistics/participants/*`.
- ROI and claim data:
  `GET /api/v2/cosmicgame/statistics/leaderboard/roi`,
  `GET /api/v2/cosmicgame/statistics/claims`, and
  `GET /api/v2/cosmicgame/rounds/{round}/claims`.
- Frequency, activity, type composition and top-bidder periods:
  `GET /api/v2/cosmicgame/statistics/bidding/*`.
- Complete deployed address registry:
  `GET /api/v2/cosmicgame/contracts/addresses`.
- Owner-tunable, mechanics-aware parameters:
  `GET /api/v2/cosmicgame/contracts/configuration`.
- Exact CosmicGame and charity balances:
  `GET /api/v2/cosmicgame/contracts/balances`.
- Next ETH/CST prices, next CST reward and normalized auction progress:
  `GET /api/v2/cosmicgame/rounds/current/bid-prices`.
- Endurance-champion, chrono-warrior and last-bidder standings:
  `GET /api/v2/cosmicgame/rounds/current/special-winners`.

## CosmicGame user profile

Replace the aggregate portion of
`GET /api/cosmicgame/user/info/{user_addr}` with:

- `GET /api/v2/cosmicgame/users/{address}` for the bounded activity profile.
- `GET /api/v2/cosmicgame/users/{address}/bids?limit=&cursor=` for bid history.

The profile fields map as follows:

- `UserInfo.Address` becomes the checksummed top-level `address`;
  `AddressId` is intentionally removed.
- `NumBids` and `MaxBidAmount` become `bidding.bidCount` and exact
  `bidding.maxEthBidWei`. Exact lifetime ETH/CST spend is added as
  `totalEthSpentWei` and `totalCstSpentWei`.
- `NumPrizes`, `MaxWinAmount`, `RewardNFTsCount`, `TotalCSTokensWon` and
  `UnclaimedNFTs` are replaced by canonical `prizes.*` counts and exact wei.
  The ambiguous legacy “CSTokens” name actually counted ERC-721 prizes and is
  not retained.
- Raffle ETH/NFT totals become `raffles.*`, reconstructed from allocation
  events rather than the mutable wallet-availability aggregate.
- `TotalDonatedCount` and `TotalDonatedAmountEth` become exact
  `ethDonations.donationCount` and `totalDonatedWei`.
- `CosmicSignatureNumTransfers` becomes
  `transfers.cosmicSignatureTransferCount`; the corresponding
  `cosmicTokenTransferCount` is now exposed alongside it.
- CST and RandomWalk staking counters are separate typed `cstStaking` and
  `randomWalkStaking` objects with exact reward wei.

A valid wallet absent from the index returns a zero-valued `200` profile and
an empty bid page, so clients do not need a special not-found state. The v1
mega-response's staking-action arrays are replaced by the user staking
resources below; its transfer, owned-NFT and marketing arrays by the user
activity resources below. With the activity group the `user/info`
mega-response is fully decomposed.

## CosmicGame user winnings

The winnings group replaces nine v1 routes with three cursor-paginated user
resources. Every resource answers a valid unindexed wallet with an empty
`200` page.

`GET /api/v2/cosmicgame/users/{address}/prizes` lists every typed prize the
wallet won, newest round first with the stable round-prize type and
winner-index order inside each round, reusing the round-prize item shape
plus the winning wallet:

- `prizes/history/by_user/{user_addr}/{offset}/{limit}` maps to it directly.
  The union's non-prize rows are deliberately split off: third-party
  withdrawal records (legacy `RecordType` 18) surface as `withdrawal`
  beneficiaries on raffle-eth-deposits, and donated NFT/ERC-20 timeout
  claims (16/17) live on the donated-asset resources below.
- `prizes/deposits/raffle/by_user/{user_addr}` rows are the
  `bidderRaffleEth` prize items; `prizes/deposits/chrono_warrior/by_user`
  rows are the `chronoWarriorEth` items. Legacy `RecordType` tags (1/2) and
  float ETH amounts are removed.
- The round-wide `cosmicSignatureStakingEth` allocation has no single
  winner and never appears; per-wallet staking rewards remain a future
  staking-history resource.

`GET /api/v2/cosmicgame/users/{address}/raffle-eth-deposits` is the
PrizesWallet ETH ledger with exact `ethAmountWei`, a `source` discriminator
(`raffle`/`chronoWarrior`), per-deposit `claimed` state and the claiming
`withdrawal` (event, transaction, UTC time, beneficiary) when one exists:

- `prizes/eth/all/by_user/{user_addr}` maps to the unfiltered ledger; the
  legacy `RecordType` 7/10 tags become the `source` field.
- `prizes/eth/raffle/by_user` and `prizes/eth/chronowarrior/by_user` become
  client-side `source` filtering — deliberately not separate endpoints.
- `prizes/eth/unclaimed/by_user/{user_addr}/{offset}/{limit}` and its
  `prizes/deposits/unclaimed/by_user` alias map to `?claimed=false`; the
  claimed filter documents its weak membership semantics under live claims.

`GET /api/v2/cosmicgame/users/{address}/raffle-nft-wins` replaces
`raffle/nft/by_user/{user_addr}`: newest-first raffle NFT wins across all
three pools with exact `cstAmountWei` and both pool flags (`isStaker`,
`isRandomWalk`). The v1 envelope's inline `UserInfo` aggregate is not
retained — the user profile resource owns those numbers.

## CosmicGame user donations

The donations group replaces seven v1 routes with five user resources.
Donor-side histories mirror the round donation collections (same item
shapes, donor scope instead of round scope):

- `donations/eth/by_user/{user_addr}` becomes
  `GET /api/v2/cosmicgame/users/{address}/eth-donations` (plain/withInfo
  discriminator, exact wei; the legacy 0/1 `RecordType` and `-1`/empty
  sentinels are gone).
- `donations/erc20/donated/by_user/{user_addr}` becomes
  `GET /api/v2/cosmicgame/users/{address}/erc20-donations` with exact
  `amountBaseUnits` (arbitrary tokens need not use 18 decimals).
- `donations/nft/by_user/{user_addr}` becomes
  `GET /api/v2/cosmicgame/users/{address}/nft-donations`.

Winner-and-claimer entitlements are two dedicated resources:

- `GET /api/v2/cosmicgame/users/{address}/donated-nfts` replaces
  `donations/nft/claims/by_user` and `donations/nft/unclaimed/by_user`: one
  row per donated NFT from a round the wallet won plus each NFT the wallet
  timeout-claimed elsewhere, with per-item claim state, the claiming wallet
  and an optional `status=claimed|unclaimed` filter.
- `GET /api/v2/cosmicgame/users/{address}/donated-erc20` replaces
  `donations/erc20/by_user` and `donations/erc20/claims/by_user`: one
  summary per round and token with exact donated/claimed/remaining
  base-unit totals and the latest claim event. **Deliberate correction:**
  the v1 winner view reported the trigger-decremented remainder as the
  donated total (zero after a full claim, with a negative donate-claim
  difference); v2 reconstructs the true donated total, and
  `donated = claimed + remaining` always holds.

## CosmicGame user staking

The staking group replaces twelve v1 paths (fifteen registered routes: the
three `staking/randomwalk/*` handlers are also mounted under the legacy
`staking/rwalk/*` aliases) with eight cursor-paginated user resources under
`/api/v2/cosmicgame/users/{address}/staking/…`. Every wei amount is an exact
decimal string — the v1 float-ETH fields are removed — and every resource
answers a valid unindexed wallet with an empty `200` page.

Event histories:

- `staking/cst/actions/by_user/{user_addr}/{offset}/{limit}` becomes
  `GET …/staking/cst/actions`: one interleaved stake/unstake ledger, newest
  first by immutable event-log ID with an `actionType` discriminator.
  Unstake items carry the exact `rewardWei` that unstake transaction
  collected; stake items never do. **Deliberate correction:** v1 applied
  `OFFSET/LIMIT` inside each branch of its stake/unstake `UNION`, so pages
  beyond the first skipped and duplicated events; the v2 cursor pages the
  merged ledger.
- `staking/randomwalk/actions/by_user/{user_addr}/{offset}/{limit}` (and its
  `rwalk` alias) becomes `GET …/staking/random-walk/actions` — the same item
  shape without reward fields, because RandomWalk staking earns no ETH.

Live staked-token membership:

- `staking/cst/staked_tokens/by_user/{user_addr}` becomes
  `GET …/staking/cst/staked-tokens`: ascending token order with the locking
  stake action and the token's mint provenance (`mintRound`, `seed`,
  optional `tokenName`). The v1 inline mint-record envelope with
  winner/current-owner fields is not retained — token ownership belongs to
  the token resources.
- `staking/randomwalk/staked_tokens/by_user/{user_addr}` (and its `rwalk`
  alias) becomes `GET …/staking/random-walk/staked-tokens`.

Reward accounting (CST staking wallet only, as in v1):

- `staking/cst/rewards/to_claim/by_user/{user_addr}`,
  `staking/cst/rewards/collected/by_user/{user_addr}/{offset}/{limit}` and
  `staking/cst/rewards/by_user/by_deposit/{user_addr}` collapse into one
  `GET …/staking/cst/deposits` ledger: one row per staking ETH deposit the
  wallet had staked tokens in, newest deposit first, with the pool-wide
  deposit (`totalDepositWei`, `totalStakedNfts`, `amountPerTokenWei`) and
  the wallet's share (`rewardWei = claimedWei + pendingWei`, token counts,
  `fullyClaimed`). `?claimed=true|false` keeps fully claimed or still
  pending deposits — v1's `to_claim` view maps to `?claimed=false`.
- `staking/cst/rewards/action_ids_by_deposit/{user_addr}/{deposit_id}`
  becomes `GET …/staking/cst/deposits/{depositId}/rewards`: the smallest
  reward units (stake action, token, exact `rewardWei`, `claimed`) in
  ascending action order. An unknown deposit answers `404`. The v1 response
  decorated each row with the *deposit's* transaction mislabeled as claim
  data; rewards are collected by the token's unstake transaction, which
  lives on the actions ledger.
- `staking/cst/rewards/by_user/by_token/summary/{user_addr}` becomes
  `GET …/staking/cst/token-rewards`: per-token exact
  `totalWei = collectedWei + pendingWei` (v1 exposed only float ETH).
- `staking/cst/rewards/by_user/by_token/details/{user_addr}/{token_id}`
  becomes `GET …/staking/cst/token-rewards/{nftTokenId}/deposits`: the
  per-deposit rewards one staked token earned, ascending by deposit, with
  the deposit event identity. An unminted token answers `404`; a minted
  token the wallet never earned rewards with answers an empty page.

Staker raffle NFT mints need no new endpoints:
`staking/cst/mints/by_user/{user_addr}` and
`staking/randomwalk/mints/by_user/{user_addr}` (plus alias) map to the
existing `GET /api/v2/cosmicgame/users/{address}/raffle-nft-wins` — filter
client-side on `isStaker` and `isRandomWalk`. Their global counterparts are
mapped by the global staking resources below.

## CosmicGame user activity

The activity group replaces six v1 paths — and retires a seventh from v2
scope — with four cursor-paginated resources and two bounded summaries
under `/api/v2/cosmicgame/users/{address}/…`. Every amount is an exact
base-unit string and every resource answers a valid unindexed wallet with
an empty `200` page or the all-zero summary shape.

Token directory and transfer ledgers:

- `cst/list/by_user/{user_addr}/{offset}/{limit}` becomes
  `GET …/cosmic-signature-tokens`: the wallet's current Cosmic Signature
  NFTs in ascending token order with mint provenance (`mintRound`, `seed`,
  original `winnerAddress`, optional `tokenName`), a typed `mintType`
  (`mainPrize`, `bidderRaffle`, `randomWalkStakerRaffle`,
  `cosmicSignatureStakerRaffle`, `enduranceChampion`, `lastCstBidder`,
  `chronoWarrior`) and live `staked` membership. **Deliberate
  corrections:** v1 derived the staked flag by pairing stake/unstake event
  joins, which duplicated rows after repeated stake cycles — v2 reads the
  live membership table; v1 also had no chrono-warrior badge and showed
  those NFTs as main prizes.
- `cst/transfers/by_user/{user_addr}/{offset}/{limit}` becomes
  `GET …/cosmic-signature-transfers`: mints, burns and transfers newest
  first by immutable event-log ID with a typed `transferType`, both
  counterparties and the wallet-relative `direction` (`in`/`out`/`self`).
  V1 ordered by a surrogate row ID and paged with `OFFSET/LIMIT`.
- `ct/transfers/by_user/{user_addr}/{offset}/{limit}` becomes
  `GET …/cosmic-token-transfers` — the same ledger shape with exact
  `amountWei` (the v1 float mirror is removed).
- `marketing/rewards/by_user/{user_addr}/{offset}/{limit}` becomes
  `GET …/marketing-rewards` with exact `amountWei`.

Bounded summaries:

- `ct/summary/by_user/{user_addr}` becomes `GET …/cosmic-token-summary`:
  the indexed `balanceWei`, earnings by source, `consumedInBidsWei` and a
  signed `netWei`, all computed in one query so every field reflects the
  same database snapshot (v1 issued eight sequential queries and returned
  floats). **Deliberate correction:** v1's earnings breakdown omitted the
  endurance-champion and last-CST-bidder CST prizes; v2 adds
  `enduranceChampionPrizesWei` and `lastCstBidderPrizesWei` and includes
  them in `totalWei`.
- `user/notif_red_box/{user_addr}` becomes `GET …/pending-winnings`:
  unclaimed raffle and chrono-warrior ETH (split by the chrono-warrior
  registration join rather than v1's hardcoded winner-index threshold),
  the pending donated-NFT count, the exact unclaimed staking `rewardWei`
  (v1 exposed float ETH) and the count of donated ERC-20 entitlements with
  a remaining balance. The v1 inline donated-ERC-20 array is unbounded and
  is not retained — itemized views live on `raffle-eth-deposits`,
  `donated-nfts`, `donated-erc20` and `staking/cst/deposits`.

Live balances are retired from v2 scope (decision D10): v2 handlers never
perform request-time RPC, so `user/balances/{user_addr}` has no v2
replacement. The indexed Cosmic Token balance is `cosmic-token-summary`'s
`balanceWei`; wallets read live ETH balances from the chain directly. The
v1 route keeps serving until the v1 sunset.

## CosmicGame global token and marketing directories

The global-directories group replaces thirteen v1 paths with ten v2
operations. Every amount is an exact base-unit string; cursor-paginated
collections use the standard `limit`/`cursor` parameters.

Cosmic Signature token directory:

- `cst/list/all/{offset}/{limit}` becomes
  `GET /api/v2/cosmicgame/cosmic-signature-tokens`: every minted token,
  newest mint first by immutable token ID, with mint provenance (typed
  `mintType`, `mintRound`, `seed`, original `winnerAddress`), the live
  `currentOwnerAddress`, optional `tokenName` and live `staked`
  membership. **Deliberate corrections:** the v1 list's stake-event joins
  duplicated rows after repeated stake cycles, mislabeled chrono-warrior
  NFTs as main prizes, and had no chrono-warrior record type — v2 reads
  the live membership table and derives one of seven `mintType` values
  with exactly-one-source enforcement.
- `cst/names/named_only` becomes the `?named=true` filter on the same
  collection; `cst/names/search/{name}` becomes the `?name=` filter.
  Both page on the immutable token keyset instead of returning unbounded
  arrays, and the search term is matched literally — ILIKE wildcards in
  the term are escaped, unlike v1. The v1 name-based sort of
  `named_only` is not retained; clients sort the bounded pages they
  need.
- `cst/info/{token_id}` becomes
  `GET /api/v2/cosmicgame/cosmic-signature-tokens/{nftTokenId}`: the
  lean detail with provenance, current owner, name and live staking
  state (`currentStake` carries the locking action, staker and
  timestamp while staked). The v1 response embedded the full round
  mega-record for main-prize tokens; round data lives on
  `/rounds/{round}` in v2.
- `cst/names/history/{token_id}` becomes
  `GET …/cosmic-signature-tokens/{nftTokenId}/name-history`: the
  token's renames newest first by immutable event-log ID with the
  renaming wallet; an empty `tokenName` cleared the name. An unknown
  token answers `404`.
- `cst/transfers/all/{token_id}/{offset}/{limit}` becomes
  `GET …/cosmic-signature-tokens/{nftTokenId}/transfers`: the token's
  full ownership history — mint included — newest first by immutable
  event-log ID with a typed `transferType`. V1 ordered by a surrogate
  row ID and paged with `OFFSET/LIMIT`.
- `cst/distribution` becomes
  `GET …/cosmic-signature-tokens/holders`: one row per current holder,
  largest holding first with a stable internal tie-break and cursor
  pagination instead of v1's unbounded array. Like the participant
  directories, ranks are weakly consistent while holdings change.

Cosmic Token (ERC-20) views:

- `ct/balances` becomes `GET /api/v2/cosmicgame/cosmic-token/holders`:
  wallets with a positive indexed balance, largest first with a stable
  tie-break, exact `balanceWei` and cursor pagination (v1 returned every
  row unbounded with float mirrors).
- `ct/statistics` becomes `GET /api/v2/cosmicgame/cosmic-token/statistics`:
  the game-wide position from one consistent snapshot query — supply,
  holder count, earnings by source, `consumedInBidsWei`, a signed
  `netWei`, transfer counters and the top-10 holders with deterministic
  decimal-string `shareOfSupply`. **Deliberate corrections:** v1 issued
  eight sequential queries (torn reads), returned float mirrors and
  omitted the endurance-champion and last-CST-bidder CST prize sources;
  v2 adds both and proves the sources sum to `earned.totalWei`.
  V1's live token name/symbol/decimals reads are retired with the
  request-time RPC (decision D11): they are immutable public chain data
  clients read once from the chain; the token address is
  `/contracts/addresses.cosmicTokenAddress`.
- `ct/total_supply_history_by_bid` becomes
  `GET /api/v2/cosmicgame/cosmic-token/supply-history/by-bid`: one row
  per bid, oldest first by immutable event-log ID, with exact minted,
  burned, net and running-total base units — cursor-paginated instead of
  v1's unbounded full-table response.
- `ct/total_supply_history_by_date/{from}/{to}` becomes
  `GET /api/v2/cosmicgame/cosmic-token/supply-history/daily?from=&to=`:
  daily aggregates over a half-open `[from,to)` window of RFC 3339
  dates (v1 used inclusive `YYYYMMDD` path segments), capped at 2,000
  days.

Marketing:

- `marketing/rewards/global/{offset}/{limit}` becomes
  `GET /api/v2/cosmicgame/marketing-rewards`: the global reward ledger,
  newest first by immutable event-log ID, with `marketerAddress` and
  exact `amountWei` on each row.
- `marketing/config/current` is decomposed (decision D11): the
  owner-tunable `treasurerAddress` joins `GET
  /api/v2/cosmicgame/contracts/configuration` (refreshed with the
  five-minute constants group — v2 still performs no request-time RPC),
  and the wallet and token addresses were already in
  `/contracts/addresses`. The MarketingWallet's `owner()` is chain
  governance data, not game state, and is retired from v2 scope; the v1
  route keeps serving until the v1 sunset.

## CosmicGame global staking

The global-staking group replaces ten v1 behaviors (fourteen registered
paths because all four RandomWalk handlers also have `staking/rwalk/*`
aliases) with nine cursor-paginated v2 operations. Every amount is an exact
wei string; every collection replaces `OFFSET/LIMIT` or an unbounded array
with an endpoint-scoped keyset cursor.

Global action ledgers and lifecycle details:

- `staking/cst/actions/global/{offset}/{limit}` becomes
  `GET /api/v2/cosmicgame/staking/cst/actions`: the interleaved global
  stake/unstake ledger newest first by immutable event-log ID, carrying
  the staker and round. Unstake rows add exact `rewardWei` and
  `rewardPerTokenWei`; stake rows omit both.
- `staking/randomwalk/actions/global/{offset}/{limit}` and its `rwalk`
  alias become `GET …/staking/random-walk/actions`, with the same event
  shape but no reward fields because RandomWalk staking earns no ETH.
- The two `staking/*/actions/info/{action_id}` handlers (plus the
  RandomWalk alias) become `GET …/staking/{cst|random-walk}/actions/{actionId}`:
  one lifecycle resource containing the required stake event and an
  optional matching unstake event. Unknown actions answer `404`; mapper
  checks ensure the two events agree on action, token and staker.

Live global membership:

- `staking/cst/staked_tokens/all` becomes
  `GET /api/v2/cosmicgame/staking/cst/staked-tokens`: ascending live token
  membership with the staker, locking action and Cosmic Signature mint
  provenance (`mintRound`, `seed`, optional `tokenName`).
- `staking/randomwalk/staked_tokens/all` and its alias become
  `GET …/staking/random-walk/staked-tokens`. Both resources document that
  live unstakes can move rows while a client traverses pages.

Reward accounting:

- `staking/cst/rewards/global` becomes
  `GET /api/v2/cosmicgame/staking/cst/deposits`: one row per staking-wallet
  ETH deposit, newest first, with exact total, per-token, collected,
  pending and integer-division remainder wei plus reward counts.
  **Deliberate correction:** v1 calculated pending as
  `deposit - collected`, silently folding the non-claimable division
  remainder into pending. V2 exposes `remainderWei` separately and proves
  `totalDepositWei = collectedWei + pendingWei + remainderWei`.
- `staking/cst/rewards/by_round/{round_num}` becomes
  `GET /api/v2/cosmicgame/rounds/{round}/staking-rewards`: bounded
  per-staker allocations ordered by deposit and wallet, with
  `rewardWei = collectedWei + pendingWei`. Missing and open rounds answer
  `404`.

Staker-raffle mints:

- `staking/cst/mints/global/{offset}/{limit}` and
  `staking/randomwalk/mints/global/{offset}/{limit}` (plus alias) become one
  `GET /api/v2/cosmicgame/staking/raffle-nft-wins?pool=cst|randomWalk`
  collection. The selected pool is embedded in the cursor, preventing
  cross-pool continuation; rows reuse the exact typed raffle NFT winner
  shape.

## Contract field mapping

The dashboard's `ContractAddrs` object maps one-for-one to
`/contracts/addresses`, with camelCase names and checksummed addresses.

The following dashboard fields move to `/contracts/configuration`:

- `ContractMechanicsVersion` becomes `mechanicsVersion`.
- `PriceIncrease` becomes `ethBidPriceIncreaseDivisor`.
- `CharityAddr` and the five percentage fields become typed address/integer
  fields.
- `TimeIncrease`, raffle winner counts, initial-duration divisor and claim
  timeout retain their contract units with explicit names.
- V1 exposes `fixedCstBidRewardWei` and a CST auction divisor.
- V2 exposes `cstBidRewardMultiplier`, a CST auction duration and its
  duration-change divisor.
- The MarketingWallet's owner-tunable treasurer is `treasurerAddress`
  (from `marketing/config/current`, decision D11).

`BidPrice` moves to `/rounds/current/bid-prices` as
`nextEthBidPriceWei`. The v1 `bid/eth_price` and `bid/cst_price` routes map to
the same resource. V2 auction progress is block-pinned and clamped to
`0..duration`; a V2 CST auction start timestamp is normalized to elapsed
seconds.

The dashboard's charity balance moves to `/contracts/balances`. The v1
floating-point CosmicGame balance is replaced by exact
`cosmicGameBalanceWei`.

The v1 `bid/current_special_winners` route maps to
`/rounds/current/special-winners`. Contract reads and the optional CST bid
event lookup are pinned to the same source block. The legacy epoch-zero CST
last-bid timestamp sentinel is omitted.

## Intentional v2 behavior changes

- V2 never rebuilds the dashboard mega-response.
- Wei and divisors are decimal strings; float ETH fields are removed.
- Live handlers perform no request-time RPC. A failed or uninitialized cache
  returns RFC 9457 `503` with `Retry-After`.
- Configuration is served only when constant and variable fields belong to
  the same detected contract mechanics generation.
- V2 does not reproduce the dashboard's request-time bid-price write-back or
  activation-time RPC override.
- Empty collections are `[]`; optional values are omitted instead of using
  magic sentinels.

## RandomWalk explorer, marketplace and statistics

Seventeen `/api/v2/randomwalk/*` operations replace twenty-one v1 read
behaviors (twenty-three registered paths counting the
`cosmicgame/randomwalk/tokens/info` and `rwalk` conveniences). Every amount
is an exact wei string, wallets are addressed by checksummed address rather
than v1's internal `user_aid`, and every collection replaces `OFFSET/LIMIT`
segments or unbounded arrays with an endpoint-scoped keyset cursor.

Token directory and per-token histories:

- `tokens/list/sequential` — which fetched every mint with a hardcoded
  ten-billion row limit — becomes `GET /api/v2/randomwalk/tokens`: the
  cursor-paginated mint directory in immutable ascending token order with
  mint provenance (minter, seed, exact `mintPriceWei`) plus the live owner,
  name and exact trading state. `tokens/list/by_period/{init_ts}/{fin_ts}`
  becomes the `mintedFrom`/`mintedUntil` half-open window filter;
  `named`/`name` filters match current names with ILIKE wildcards escaped.
- `top5tokens` becomes `GET …/tokens?sort=mostTraded&limit=5`: the
  most-traded ranking is a directory sort whose cursor carries the live
  trade-count rank.
- `tokens/info/{token_id}` (and its `/api/cosmicgame/randomwalk/…` alias)
  becomes `GET …/tokens/{tokenId}`, adding the rename recency as
  `nameChangedAt`. Unknown tokens answer `404` instead of v1's HTTP 200
  envelope carrying `"sql: no rows in result set"`.
- `tokens/name_changes/{token_id}` becomes `GET …/tokens/{tokenId}/name-history`,
  newest first by immutable event-log ID with a `404` gate.
- `tokens/history/{token_id}/{offset}/{limit}` becomes
  `GET …/tokens/{tokenId}/events`: one newest-first provenance ledger over
  six event sources with a typed `eventType` (`mint`, `transfer`,
  `nameChange`, `listed`, `offerCanceled`, `purchase`). Transfers that
  merely mirror the mint or a purchase in the same transaction are
  represented once, by the mint or purchase event; burns surface as
  transfers to the zero address.

Marketplace order book and ledgers:

- `current_offers/{order_by}` becomes
  `GET /api/v2/randomwalk/marketplace/offers`: the live book under a named
  `sort` (`newest`, `oldest`, `priceAsc`, `priceDesc` — replacing v1's
  numeric selector, whose unknown values silently fell back to insertion
  order). Price ranks page on an exact `(price, eventLogId)` keyset bound
  to the sort.
- `trading/history/{offset}/{limit}` becomes `GET …/marketplace/offer-history`:
  the immutable offer-creation ledger, each row carrying its outcome
  (`active`, `bought` with the purchase event and parties, or `canceled`
  with the cancellation event) and the signed exact `sellerProfitWei` when
  the marketplace tracked the seller's position.
- `trading/sales/{offset}/{limit}` becomes `GET …/marketplace/trades`: the
  purchase ledger newest first by the purchase event.
- `trading/by_user/{user_aid}/{offset}/{limit}` becomes
  `GET …/users/{address}/offers`. **Deliberate correction:** v1 parsed the
  offset/limit segments and then ignored them, returning the wallet's
  entire closed-offer history; v2 pages on a wallet-scoped cursor and also
  includes the wallet's still-active offers.
- `floor_price` becomes `GET …/marketplace/floor-price` with the active
  sell-offer count and the cheapest listing; an empty book omits `floor`
  instead of reproducing v1's `DBError: "sql: no rows in result set"`
  sentinel.

Users:

- `user/info/{user_aid}` becomes `GET /api/v2/randomwalk/users/{address}`.
  Exact `tradingVolumeWei` and signed `profitWei` replace float ETH, and a
  live `ownedTokenCount` is added. **Deliberate corrections:** the mint
  count is computed from mint events because the legacy accumulator
  silently dropped mints made before a wallet's first trade, and the
  withdrawal count comes from the withdrawal ledger because the legacy
  `total_withdrawals` accumulators were never written by anything. The
  `IsMarketPlaceContract` flag is not retained — the marketplace address
  is served by the contract registry. Valid unindexed wallets return the
  zero-activity `200` shape.
- `tokens/by_user/{user_aid}` becomes `GET …/users/{address}/tokens`:
  ascending live ownership with mint provenance and exact trading state,
  paged on a wallet-scoped cursor.

Statistics and ledgers:

- `statistics/by_token` and `statistics/by_market` become one
  `GET /api/v2/randomwalk/statistics` snapshot: exact collection aggregates
  (mint count, live unique-owner count, trading volume, `mintFundsWei`,
  the last mint), exact marketplace aggregates with live active-offer
  counts, and the withdrawal count with the latest withdrawal.
  **Deliberate corrections:** v1's `MaximumTradedPrice`, `CurOwnerAid` and
  `CurOwnerAddr` fields were never populated by the query and are not
  retained; v1's `WithdrawalAmount` (half the mint-price sum — an
  approximation that ignores prior withdrawals) is retired, since the live
  withdrawable prize is a chain read; the unique-user count is the live
  distinct-owner count instead of the row count of the cron-maintained
  `rw_uranks` table.
- `statistics/trading_volume/{init_ts}/{fin_ts}/{interval_secs}` becomes
  `GET …/statistics/trading-volume?from=&to=&intervalSeconds=`: zero-filled
  buckets bounded to 2,000, exact per-bucket `volumeWei` and an exact
  `cumulativeVolumeWei` continued from the exact `baseVolumeWei` before the
  window (v1 accumulated float ETH).
- `statistics/floor_price/{init_ts}/{fin_ts}/{interval_secs}` becomes
  `GET …/statistics/listing-floor-history`, named for what the query has
  always measured: the minimum price among sell offers *created* in each
  bucket — listing pressure, not an order-book reconstruction. Buckets
  without new sell listings are omitted, as in v1.
- `mint_report` becomes `GET …/statistics/mint-report`. **Deliberate
  correction:** v1 hardcoded the November 2021 – December 2022 window; v2
  aggregates every month with at least one mint and carries exact monthly
  and cumulative wei. The float `RedeemAmount` column (the same half-sum
  approximation as above) is not retained.
- `statistics/withdrawal_chart` is replaced by the exact withdrawal ledger
  `GET /api/v2/randomwalk/withdrawals` plus `statistics.tokens.mintFundsWei`;
  the chart's cumulative half-mint-price curve is the same approximation
  retired above and clients that want it can derive it from `/tokens` pages.
- `statistics/mint_intervals` is retired from v2 scope (decision D12): the
  mint set is frozen and inter-mint intervals are derivable from the
  `mintedAt` column while paging `/tokens`.
- `contracts` becomes `GET /api/v2/randomwalk/contracts/addresses` with
  checksummed `nftAddress` and `marketplaceAddress`; the internal address
  IDs are not exposed.

Staying on v1 (decision D12):

- `metadata/{token_id}` (and the host-dispatched `/metadata/{token_id}`)
  is the ERC-721 `tokenURI` target hardcoded in the deployed contract; the
  path can never move and stays permanently outside `/api/v2`.

## RandomWalk beauty-contest ranking

The ranking mini-app is the first v2 write surface (decision D13,
[ADR-0008](adr/0008-api-v2-writes.md)). Seven operations under
`/api/v2/randomwalk/ranking/` replace ten registered v1 paths. Because the
challenge issuance moves from a side-effecting GET to a POST and the vote
request names the winner by token id, migrating this group is a small
deliberate frontend rewrite rather than a URL swap — the wallet signature
itself is unchanged.

Reads:

- `explore/random` and `random` become
  `GET /api/v2/randomwalk/ranking/random-tokens?sampleSize=` (1–100,
  default 2). Same fewest-matches-then-lowest-rating selection; the reply
  is a typed `{"tokenIds": [...]}` object instead of a bare array. The v1
  `ORDER BY RANDOM()` fallback for half-migrated databases is not retained:
  a failing ranked query is an opaque 500.
- `ranking/beauty-pair-ids` becomes
  `GET /api/v2/randomwalk/ranking/pair?voter=`. The reply names
  `firstTokenId`/`secondTokenId` and keeps `pairExhausted`. v1's
  `skip_pair_filter=1` is expressed by omitting `voter`; a collection
  without two tokens answers a 404 problem instead of a short array.
- `token-ranking/order` and `rating_order` become the cursor-paginated
  `GET /api/v2/randomwalk/ranking/ratings` — ascending `(rating, tokenId)`
  like v1's order, but each row carries the rating value and the match
  count, and the unbounded all-token id array is retired.
- `vote_count` becomes `GET /api/v2/randomwalk/ranking/statistics`:
  `total_count` maps to `totalVotes`, alongside the new `walletVotes`,
  `distinctVoters` and `ratedTokens` counters.

Writes (both flows keep v1's per-IP rate limits; over-limit requests
answer a spec-declared 429 problem with `Retry-After`):

- `GET ranking/sign-challenge` becomes
  `POST /api/v2/randomwalk/ranking/challenges` → `201 {nonce, expiresAt}`.
  The nonce contract is unchanged (one-time, 15 minutes); `expiresAt` is
  new and comes from the database clock that also judges expiry.
- `POST add_game` becomes `POST /api/v2/randomwalk/ranking/votes`. The
  request renames `nft1`/`nft2` to `firstTokenId`/`secondTokenId`,
  replaces the 0/1 `nft1_win` integer with `winnerTokenId`, and renames
  `sign_nonce` to `nonce`. The signed message is **byte-identical** to
  v1 (`RandomWalk beauty vote\nVersion: 1\n…`), so existing wallet
  prompts port as-is. Success is `201` with both new ratings and the
  recovered `voterAddress` instead of `{"result":"success"}`; the
  duplicate-pair rejection stays `409`, and invalid pairs, disallowed
  chains, bad signatures and unknown/expired nonces answer `400`
  problems with specific types (`invalid-pair`, `chain-not-allowed`,
  `invalid-signature`, `invalid-nonce`).
- `POST token-ranking/match` becomes
  `POST /api/v2/randomwalk/ranking/matches` with the same
  `X-Ranking-Admin-Key` header contract (fail-closed 503 when
  unconfigured, 401 on a wrong key — now as RFC 9457 problems). The body
  uses the same `firstTokenId`/`secondTokenId`/`winnerTokenId` shape as
  votes instead of v1's `nft1_won` boolean; success is `201` with both
  new ratings.

## Remaining endpoint groups

Every v1 surface is now mapped: the CosmicGame groups above, the
RandomWalk explorer/marketplace/statistics group, and the ranking
mini-app (reads and writes). The only v1 route without a v2 replacement
is the contract-pinned metadata route (documented under D12 above).
