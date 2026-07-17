# ADR-0009: API v2 bid moderation

- Status: Accepted
- Date: 2026-07-17

## Context

The v2 migration guide declared the v1 surface fully mapped after the ranking
slice, but three CosmicGame moderation behaviors were omitted:
`get_banned_bids`, `ban_bid`, and `unban_bid`. They are the final movable v1
surface; FAQ is an upstream proxy and the metadata path is pinned by deployed
contracts.

The v1 moderation API exposes an unbounded bare array, accepts a redundant
`user_addr` supplied by the administrator, permits duplicate rows at the
database boundary, and models both create and remove as operation-shaped
POSTs. V1 remains frozen while consumers migrate.

## Decision

### One addressable active-ban resource

API v2 exposes:

- `GET /api/v2/cosmicgame/moderation/banned-bids` as a newest-first,
  cursor-paginated collection;
- `POST /api/v2/cosmicgame/moderation/banned-bids` with `{bidId}` as a
  required body, returning `201` with the created active ban; and
- `DELETE /api/v2/cosmicgame/moderation/banned-bids/{bidId}`, returning
  `204` when the active ban is removed.

This narrows ADR-0008's POST rule: POST + typed `201` remains the convention
for commands that create records, while an individually addressable resource
is removed with DELETE + `204`. Missing bids and missing active bans are
typed `404` problems; a duplicate create is `409`.

The server derives `bidderAddress` from the indexed bid. Clients cannot submit
a contradictory address. Responses use checksummed addresses and RFC 3339 UTC
timestamps; internal row IDs appear only inside opaque versioned cursors.

### The database owns uniqueness

Migration 00025 collapses historical duplicates by `bid_id`, retaining the
newest row, then replaces the lookup index with a unique index. The constraint
makes concurrent duplicate creates race-safe and gives the item URL one
unambiguous identity.

### Existing admin deployment contract

The two writes declare an OpenAPI `AdminKey` scheme using `X-Admin-Key` and
`ADMIN_API_KEY`. Enforcement shares the generated strict middleware used by
the ranking scheme, hashes both values to a fixed width before constant-time
comparison, and fails closed with `503` when the secret is unset. Create and
delete each retain v1's 2 requests/second, burst-5 per-IP budget and expose
spec-declared `429` problems with `Retry-After`.

## Consequences

- The generated v2 inventory grows from 99 to 102 operations and every
  movable v1 behavior now has a documented v2 replacement.
- Frontends send only `bidId`, consume a bounded typed page, and use DELETE
  for unban.
- Existing v1 routes and response goldens remain unchanged until the
  ADR-0005 traffic-gated sunset; the v1 create handler maps the new uniqueness
  conflict to its historical idempotent `201 {"result":"success"}` response.
- The data migration intentionally removes redundant duplicate active bans;
  no moderation meaning is lost because v1 unban already deleted every row
  for a bid.
