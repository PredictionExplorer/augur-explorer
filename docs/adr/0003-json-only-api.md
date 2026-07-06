# ADR-0003: JSON-only API; HTML explorer removed

Date: 2026-07-06. Status: accepted.

## Context

The web server historically served two parallel stacks: a JSON API under
`/api/*` and a server-rendered HTML explorer under `/black/*`. The HTML stack
duplicated nearly every JSON handler (~5,700 lines of near-identical code plus
150 template/asset files) and had drifted behind the JSON API.

## Decision

Remove the HTML explorer entirely. The server exposes:

- the JSON API under `/api/*` (contract: [docs/openapi.yaml](../openapi.yaml)),
- ERC-721 metadata at `/metadata/:token_id` and `/cg/metadata/:token_id`
  (referenced by on-chain `baseURI` values — these routes must stay stable
  forever),
- the NFT asset mirror at `/images/*`,
- health/observability endpoints.

Frontends own all presentation.

## Consequences

About 4,200 lines of handler code and 150 template files deleted; every data
fix now lands in exactly one handler. Anyone who depended on the `/black/*`
pages must move to the JSON API (accepted: those pages were an internal
explorer, and the decision to allow breaking changes was made explicitly).
