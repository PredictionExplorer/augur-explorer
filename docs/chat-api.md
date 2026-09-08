# Cycle chat reads

The v2 chat API pages message bodies independently from the cycle's full gesture
records. Legacy endpoints and their response shapes are unchanged.

## Message pages

`GET /api/v2/cosmicgame/rounds/{round}/messages?limit=50`

The default is 50 messages and the maximum is 200. Initial pages are ordered by
`(occurredAt, eventLogId)` descending. Only messages with text after ECMAScript
`String.trim()` and without an active moderation ban are selected. Moderation uses
`cg_banned_bids.bid_id = cg_bid.id`, not the distinct event-log identifier. Both
filters apply before the limit, so blank or moderated rows cannot consume slots.

The response contains one slim `data` array with `eventLogId`, `round`, `position`,
`bidderAddress`, `occurredAt`, `message`, `bidType`, and `transactionHash`.
`ethPriceWei`, `cstPriceWei`, and `randomWalkTokenId` are optional; sentinel values
are omitted and amounts retain their exact decimal-string precision. Timestamps
are UTC RFC 3339 values. Reward and attached-asset fields are not read or returned.

`meta` contains:

- `limit`: requested page size.
- `nextCursor`: present only when an initial/older page has more older messages.
- `syncCursor`: always present, including an empty cycle.
- `hasMore`: true only when an `after` request has further catch-up pages.
- `revision`: a decimal-string history invalidation generation.

Use `?cursor={nextCursor}&limit=50` for older pages. Preserve the original
`syncCursor` while loading older pages: an older-page response also captures a
fresh latest boundary, and replacing an existing synchronization boundary with it
would skip intervening arrivals.

Use `?after={syncCursor}&limit=50` for new messages. These pages are ordered
ascending by the same unique key. Merge rows by event ID, advance to the returned
`syncCursor`, and continue while `hasMore` is true. The cursor advances only to the
last returned row, so bursts larger than one page are not skipped. An empty
catch-up page preserves its incoming boundary. The browser can continue its
existing visible-tab polling; no streaming connection is required.

Cursors are opaque, versioned, and bound to the cycle, direction, timestamp/event
boundary, and revision. `cursor` and `after` are mutually exclusive. Do not decode
or synthesize cursors in clients.

## Historical changes and snapshot validity

Migration 00030 installs a durable singleton revision and transactional row
triggers. Actual gesture updates/deletions and moderation inserts/updates/deletes
advance it. This includes reorganization cascades and legacy moderation writes.
The indexer's no-op delete-before-insert and ordinary appended events do not
invalidate older pages. The singleton intentionally invalidates all cycle chat
caches on a historical change; this trades rare extra reloads for inexpensive,
reliable invalidation without scanning the history on every request.

Revision, latest visible boundary, and message rows come from one SQL snapshot.
An older or synchronization cursor with a stale revision returns HTTP 409 with
problem type ending in `/feed-reset-required`. Clear the cycle's cached messages
and context, then obtain a fresh initial page. Never append to stale history after
this response. Do not use a row count, maximum ID, or latest-page hash as a
replacement for the revision: those can miss an old message's removal or unban.

## Minimal context compatibility snapshot

`GET /api/v2/cosmicgame/rounds/{round}/chat-context` returns all gestures' minimal
metadata in ascending `(occurredAt, eventLogId)` order. Each item has `eventLogId`,
`round`, `position`, `bidderAddress`, `occurredAt`, `bidType`, `prizeAt`, and optional
`cstDutchAuctionDurationSeconds`. `meta.revision` uses the same history generation.
The query does not select message bodies, rewards, or attached assets and retains
gestures without visible messages so participant counts remain correct.

This snapshot deliberately remains **O(cycle size)**. It preserves existing
client-derived milestones and counts without downloading every message body and
full gesture record. A future authoritative aggregate/milestone endpoint should
replace this remaining full metadata read. Compare its revision with message
responses before combining snapshots; ordinary append timing can still differ
between independent requests. Polling may temporarily show a newer snapshot in
one surface, as with other independently refreshed dashboard reads.

## Errors and rollout

Empty and not-yet-indexed cycles return HTTP 200 with `data: []`. Invalid input
returns 400, stale cursor revisions return 409, and server/store failures return
500 with an opaque RFC 9457 problem. A known route never uses 404 to represent an
empty feed. Clients may fall back to the unchanged legacy API only when the v2
route is unavailable (404/501); 400/409/429/5xx and invalid responses must not be
treated as permission to silently switch protocols.

Apply migration 00030, then the concurrently built partial index in 00031, before
deploying the new server. Old servers remain compatible with those migrations.
The frontend may ship first with capability fallback and discover the new routes
after deployment. Migration and server rollout are separate operator actions;
the implementation tests apply migrations only to disposable PostgreSQL instances.

The v2 OpenAPI document is authoritative. Regenerate the committed Go interfaces,
models, and embedded spec with `go generate ./internal/api/v2`. Coverage includes
strict input/scope validation, empty cycles, same-timestamp boundaries, exact
amounts, moderation identity, non-ASCII whitespace, a 55-message catch-up burst,
transactional revision rollback, reorganization deletion, and schema-validated API
responses with deterministic goldens.
