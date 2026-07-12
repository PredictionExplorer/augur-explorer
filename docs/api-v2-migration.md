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
mega-response's prize, donation, claim, staking-action, transfer and owned-NFT
arrays remain on v1 until bounded user sub-resources land. Live
`user/balances`, CosmicToken summary, marketing history and red-box claim
status are also intentionally deferred.

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

Remaining user-scoped histories, RandomWalk resources, CosmicToken statistics
and marketing-wallet configuration stay on v1 until their dedicated v2
sprints land. Their presence does not require continued use of either the v1
dashboard or the v1 user mega-response for profile and bid data.
