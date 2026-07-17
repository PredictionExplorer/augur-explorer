# ADR-0008: API v2 write conventions and the ranking slice

- Status: Accepted
- Date: 2026-07-16

## Context

Every one of the 92 generated v2 operations is a GET. The one v1 surface
still without a v2 replacement — deferred by D12 — is the RandomWalk
beauty-contest ranking mini-app, because it writes: an anonymous-wallet
voting flow (challenge nonce → EIP-191 `personal_sign` → Elo update) and an
admin match recorder. Porting it forces the decisions every future v2
mutation will inherit, so they are recorded here rather than improvised
per endpoint.

The v1 flow works and its cryptography is sound (one-time 15-minute nonce,
EIP-191 recovery, one vote per wallet per unordered pair, per-IP rate
limits). What v1 gets wrong is wire ergonomics: a side-effecting GET issues
the nonce, the vote body encodes the winner as an `nft1_win` 0/1 integer,
responses are untyped `{"result":"success"}` envelopes, and error shapes
are ad-hoc `{"error": ...}` objects with inconsistent statuses.

## Decision

### Record-creating mutations are POSTs

- Every record-creating v2 mutation in this slice is a `POST` with a typed
  request contract. No side-effecting GETs: the v1
  `ranking/sign-challenge` GET becomes `POST .../ranking/challenges`.
- Success is `201 Created` with a typed body describing the resulting
  state (the created challenge, the recorded vote with both new ratings).
  V2 mutations create ledger records that are not individually
  addressable, so no `Location` header is emitted.
- Handlers stay transport-agnostic: caching/compression middleware ignores
  POSTs by construction (`Compress` allows them, `ConditionalETag` only
  touches GET/HEAD), and mutation responses are marked
  `Cache-Control: no-store` by the handler layer's absence of caching —
  no additional middleware work is needed.
- ADR-0009 later clarifies the resource-lifecycle case: an individually
  addressable active bid ban is removed with `DELETE` + `204`; POST + typed
  `201` remains the record-creation convention.

### Errors are RFC 9457 problems, like every v2 read

- Malformed JSON bodies reuse the generated `RequestErrorHandlerFunc` seam
  and answer `400` with kind `invalid-request` (detail: "The request body
  is not valid JSON."). Field-level validation failures answer `400` with
  specific kinds: `invalid-pair`, `chain-not-allowed`,
  `invalid-signature`, `invalid-nonce`.
- Uniqueness conflicts answer `409` (`already-voted`) — v1 already chose
  409 here and it is correct.
- Store failures stay opaque `500 internal` problems; nonce and signature
  failures never reveal which check failed beyond their kind.

### Authentication

- **Admin mutations** declare an OpenAPI `apiKey` security scheme. The
  ranking match recorder keeps v1's exact deployment contract — header
  `X-Ranking-Admin-Key`, secret from `RANKING_ADMIN_KEY` falling back to
  `ADMIN_API_KEY` — so the operator bot migrates by changing only URL and
  body. Enforcement fails closed: with no key configured the operation
  answers `503` (`admin-disabled`); a wrong or missing key answers `401`
  (`unauthorized`) after a constant-time compare. Both are
  `application/problem+json`.
- **Wallet mutations** authenticate with EIP-191 `personal_sign` over a
  canonical message that binds chain id, one-time nonce and the exact
  action parameters. The v1 message format (`RandomWalk beauty vote /
  Version: 1 / chainId / nonce / nft1 / nft2 / winner`) is kept
  byte-identical so existing wallet UX ports without re-signing changes;
  the shared builder lives in `internal/beautyrank` and is pinned by
  goldens in both API generations.

### Per-operation rate limits as strict middleware

Write operations carry stricter per-IP token buckets than the global
50 rps/100 chain limiter: votes 1 rps/burst 10, admin matches 2/5,
challenges 2/20 (a challenge precedes every vote). They are enforced as
generated *strict middleware* keyed on the operation ID — the layer where
per-operation policy belongs in a generated router — and answer
`429` problems (kind `rate-limited`) with a `Retry-After: 1` header,
declared in the spec. The global limiter still wraps everything as the
outer net; its legacy 429 envelope stays undocumented transport behavior
exactly like v1.

### The ranking slice itself

Seven operations under `/api/v2/randomwalk/ranking/` replace the eight
remaining v1 read/write behaviors (ten registered paths):

- `GET random-tokens` (fewest-matches-first sample; explicit `sampleSize`
  bound), `GET pair` (voter-aware re-roll with `pairExhausted`),
  `GET ratings` (cursor-paginated `(rating, tokenId)` directory with match
  counts — replacing v1's two unbounded bare-array orders), and
  `GET statistics` (one-snapshot vote/voter/rated-token counts).
- `POST challenges`, `POST votes`, `POST matches` as described above. The
  vote request names the winner by token id (`winnerTokenId`) instead of
  v1's 0/1 `nft1_win` integer; the admin request gains the same shape.
- Elo math is unchanged (K decays from 250 by 0.00525 per recorded match,
  floor 1, zero-sum around the 400-point logistic expectation, default
  rating 1200) and now lives in `internal/beautyrank` with both API
  generations delegating to it.
- `rw_token_ranking`/`rw_ranking_match` stay the storage; the ratings
  directory and statistics queries are bounded by the frozen ~4k-token
  collection, so no new index or migration is required.

## Consequences

- The v2 endpoint inventory (§6.2.1) is complete; D6's sunset gates can
  start counting for the ranking group once the frontend migrates.
- Future v2 record creation has a settled pattern: POST + 201 typed body,
  problem-shaped 4xx/5xx, apiKey or wallet-signature auth, and strict
  middleware for per-operation limits. Addressable resource removal follows
  ADR-0009.
- The vote and challenge flows change shape (POST challenge, renamed
  fields), so the frontend migration for this group is a deliberate
  rewrite of a small mini-app, not a URL swap — documented in
  [api-v2-migration.md](../api-v2-migration.md).
- Rate-limited operations become spec-visible (`429` + `Retry-After`),
  which kin-openapi golden validation then enforces.
- v1 keeps answering all ten legacy paths unchanged until the D6 sunset.
