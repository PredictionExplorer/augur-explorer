# API v2 migration guide

API v2 decomposes the frozen v1 mega-responses into typed resources. V1
remains available while consumers migrate and is removed only after the
sunset gates in [ADR-0005](adr/0005-api-v2.md) are met.

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
resources below; its transfer and owned-NFT arrays remain on v1 until
bounded user sub-resources land. Live `user/balances`, CosmicToken summary,
marketing history and red-box claim status are also intentionally deferred.

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
client-side on `isStaker` and `isRandomWalk`. The global `staking/*/mints`
and reward views, `staking/cst/rewards/global` and per-round staking
statistics remain v1-only until a statistics slice replaces them.

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

## Remaining endpoint groups

Remaining user-scoped transfer, token, marketing and live-balance histories,
global staking statistics, RandomWalk resources, CosmicToken statistics and
marketing-wallet configuration stay on v1 until their dedicated v2 sprints
land. Their presence does not require continued use of the v1 dashboard, the
v1 user mega-response, or the v1 winnings, donations and staking routes.
