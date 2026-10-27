# ADR-0012: CosmicSignatureGame V3 — one indexer, mechanics-discriminated surfaces

- Status: Accepted
- Date: 2026-10-26

## Context

The deployed game is being upgraded to its third mechanics generation.
V3 adds a multi-NFT main prize, split bid rewards (current/previous bidder
CST shares), late-bid duration/premium parameters, and DAO governance, and
it *repurposes* existing event shapes rather than only adding new ones: the
eighth `BidPlaced` data word — same topic, same slot — carries the CST
Dutch auction duration in seconds under V2 but the CST bid price decline
multiplier in wei per second under V3, a value far beyond JavaScript's
safe-integer range.

The Solidity line itself moved while the backend tracked it. The
`v3-2026-07-24` branch introduced the decline multiplier, then reverted it
to exact V2 behavior (Solidity commit `0bc80af0`); `v3.1-2026-08-19`
reinstated it; `v3.1-2026-10-13` at commit `9d1a6447` is the Arbitrum One
deployment candidate, ABI-identical to the 08-19 compile with only creation
bytecode changed. The V3 reinitialize also changed the ETH split
percentages and the main-prize NFT count *without* emitting the
per-parameter admin events, so anything that reconstructs configuration
from events alone goes stale at the upgrade block.

## Decision

### Additive contract binding, pinned provenance

Mechanics upgrades are additive binding types: `CosmicSignatureGameV3`
joins V1/V2 in `contracts/cosmicgame`, compiled from `v3.1-2026-10-13` @
`9d1a6447` with the production settings recorded in `contracts/README.md`.
`v3_readiness_test.go` pins the getters each layer depends on.

### Dual topic dispatch, dynamic PrizesWallet sources

Where V3 changed an event's identity the indexer registers both topic
hashes and handles them in parallel (`TopicBidEvent`/`TopicBidEventV2`,
`TopicPrizeClaimEvent` for V1/V2 beside `TopicPrizeClaimEventV3`);
`TestRegistryConstantsMatchABIEventIDs` pins every constant to its
ABI-derived event ID. PrizesWallet-emitted events use a dynamic source
set: every wallet address announced on chain becomes a source, and
`BackfillPrizesWalletEvents` idempotently closes the window between a
wallet's announcement and the end of the batch that carried it.

### Schema tracks the Solidity line with reversible, empty-safe steps

Migration 00026 adds the V3 schema (multi-NFT claims, split bid rewards,
five evented configuration values, two event-less per-round champion
durations). The decline-multiplier tables then followed the contract
churn: 00029 created them for `v3-2026-07-24`, 00030 dropped them after
the Solidity revert (safe because V3 had not reached production, so the
tables were provably empty), and 00032 re-created them for v3.1 alongside
the `CstBidPriceDeclineMultiplierChanged` /
`CstBidPriceDeclineMultiplierChangeDivisorChanged` admin events. Each step
is reversible and states its contract-commit justification in the
migration header.

### Mechanics-discriminated API responses

v2 responses carry `mechanicsVersion`, derived from the game's
`Initialized` events preceding the row's block, and discriminate on it
rather than on topic: V2 rows serve `cstDutchAuctionDurationSeconds` as an
integer, V3 rows serve `cstBidPriceDeclineMultiplierWeiPerSecond` as a
decimal string, never both (the bids and chat-context handlers mirror the
same switch). `/contracts/configuration` is a per-generation `oneOf` whose
V3 variant requires the late-bid, premium, auction-floor and
decline-multiplier fields. Legacy V1/V2 rows keep their existing shape.

### Admin records and validators extended for V3

The admin-event record catalog extends through type 56: the V3 parameter
events (types 40–46, with 43 repurposed from the retired
`LastBidderBidCstRewardAmountPercentageChanged`), the DAO Governor
proposal/vote/settings events (47–55) and `DelegateVotesChanged` (56).
Because the reinitialize emitted no per-parameter events, `validate-db`
selects deploy-time defaults by comparing each round's claim block against
the V3 upgrade block (the `Initialized` version=3 event); explicit
`cg_adm_*` override rows still win when present.

## Consequences

- Decoding the shared `BidPlaced` slot depends on the mechanics version,
  not the topic, so the version derivation from `Initialized` events is
  load-bearing; serving the V3 multiplier through the V2 integer field
  would overflow consumers, which is why the API discriminates per row.
- Clients that predate V3 keep working: historical rows are unchanged and
  new fields are additive and version-gated.
- The schema's revert/re-add churn was the price of tracking a moving
  Solidity branch with reversible steps; the parallel work streams it
  spawned allocated duplicate goose versions, later resolved by
  renumbering the token-traits and chat migrations to 00036–00038.
- A future mechanics generation repeats the pattern: a new binding type,
  new topic constants pinned to the ABI, additive migrations, and a new
  `mechanicsVersion` arm — not an in-place rewrite of V3 paths.
