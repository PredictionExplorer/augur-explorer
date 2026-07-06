# ADR-0004: Fail-closed admin auth and per-IP rate limiting

Date: 2026-07-06. Status: accepted.

## Context

The mutating endpoints (`POST /api/cosmicgame/ban_bid`, `unban_bid`,
`POST /api/randomwalk/token-ranking/match`) previously had no authentication
(or, for the ranking endpoint, authentication that silently switched off when
its environment variable was unset). There was no rate limiting anywhere.

## Decision

- Mutating endpoints require a shared-secret header (`X-Admin-Key`, or
  `X-Ranking-Admin-Key` for the ranking route) validated with a constant-time
  compare against `ADMIN_API_KEY` / `RANKING_ADMIN_KEY`.
- **Fail closed:** if no key is configured the endpoint returns 503 rather
  than allowing anonymous access.
- Public voting continues through `POST /api/randomwalk/add_game`, which
  verifies an EIP-191 wallet signature per vote and carries its own strict
  rate limit.
- Every route is behind a per-client-IP token bucket (50 rps, burst 100
  globally; 1–2 rps on mutating routes), returning 429 when exceeded.

## Consequences

Bid moderation and ranking administration can no longer be invoked
anonymously; a missing deployment variable disables the endpoint instead of
exposing it. Rate limits bound the blast radius of scraping and abuse without
affecting legitimate frontend traffic.
