# Knowledge Base Generation Guidelines

This document is the **plan for LLM-assisted doc regeneration**. Run deterministic
extractors first (`python -m knowledge.generate.run_all`), then optionally use Codex to
expand or refresh narrative docs.

## Goals

- Serve **beginners** (how to use the website) and **experts** (contracts, API, deployment)
- Minimize Codex token usage at query time by pre-generating stable reference text
- Keep all outputs **re-runnable** when repos update

## Generation pipeline

```bash
cd backend
python -m knowledge.generate.run_all
./scripts/reindex-knowledge.sh   # or: cd backend && python -m retrieval.pipeline --reindex
```

On a **fresh git clone**, `run_all` first shallow-clones three upstream GitHub repos into
`backend/data/repos/` (Cosmic-Signature, augur-explorer, cosmic-front-alternate). This
cache is gitignored and not part of the faq-bot repo. Requires `git` and network access.
Use `--skip-repo-sync` or `FAQ_BOT_SKIP_REPO_SYNC=1` only if that directory is already populated.

Local monorepo docs (e.g. `rwcg/docs/BACKEND.md`) are read via `CURSOR_VREF_PATH` when set;
contract ABIs still come from the **frontend** clone (`frontend/src/contracts/*.json`).

## Output layout

Set `KNOWLEDGE_BASE` in `backend/.env` (e.g. `/home/niko/eth/dev/ai-faq-kb`):

```
$KNOWLEDGE_BASE/
  facts/                    # machine-readable JSON
  docs/beginner/            # user-facing FAQ
  docs/expert/              # developer/operator FAQ
  docs/sources/             # copied upstream markdown + UI extracts
```

## Deterministic extractors (no LLM)

| Script | Output | Purpose |
|--------|--------|---------|
| `extract_contracts.py` | `facts/contracts-inventory.json`, expert contract doc | Counts, paths, roles |
| `extract_routes.py` | `facts/frontend-routes.json`, navigation docs | Site map |
| `extract_api.py` | `facts/api-endpoints.json`, API overview | Backend REST surface |
| `extract_deployments.py` | `facts/deployed-addresses.json`, expert deployment doc | On-chain addresses |
| `extract_abis.py` | `facts/contract-abis-summary.json`, `docs/sources/frontend-contracts/*.json`, expert integration doc | Contract ABIs + integration recipes |
| `copy_source_docs.py` | `docs/sources/*`, curated beginner/expert summaries | Import upstream docs |
| `extract_ui_pages.py` | `docs/sources/frontend-pages/*` | User-visible copy from TSX |

## LLM regeneration tasks (Codex)

Use Codex **offline** to refresh narrative docs when sources change. Each task includes
strict inputs and output paths.

### Task A — Beginner how-it-works synthesis

**Inputs:**
- `docs/sources/cursor-vref/docs/BIDDING_FLOW_AUDIT.md` (user flow sections only)
- `docs/sources/frontend-pages/src_app_game_how-it-works_page.txt`
- `docs/beginner/00-project-overview.md`

**Output:** `docs/beginner/01-how-it-works.md`

**Instructions for Codex:**
> Write for someone who has never used a dApp. No code blocks. Step-by-step. Mention
> wallet, network, `/game/play`, ETH vs CST, timer, last bidder wins. Max 800 words.

### Task B — Expert bidding contract flow

**Inputs:**
- `docs/sources/cursor-vref/docs/BIDDING_FLOW_AUDIT.md`
- `facts/contracts-inventory.json`

**Output:** `docs/expert/04-bidding-contract-flow.md`

**Instructions:**
> Write for a Solidity developer. Include entry points, revert conditions, ETH vs CST paths,
> donation approval target (PrizesWallet), and backend indexing notes. Cite file paths.

### Task C — Expert database/ETL overview

**Inputs:**
- `docs/sources/cursor-vref/rwcg/docs/BACKEND.md` (ETL + schema sections)

**Output:** `docs/expert/05-database-and-etl.md`

### Task D — Deployment FAQ refresh

**Inputs:**
- `docs/sources/smart_contracts/tasks/docs/deployment-and-registration.md`
- `facts/deployed-contract-roles.json`

**Output:** update `docs/expert/02-deployment-and-networks.md`

## Missing today (future extractors)

| Gap | Proposed extractor |
|-----|-------------------|
| Live Arbitrum addresses | Parse latest `deploy-cosmic-signature-contracts-report-*.json` if present |
| Production frontend env (RPC/API URLs) | `knowledge/deployments/*-environment.txt` → `facts/network-environment.json` |
| SQL/game mechanics | Summarize `cursor-vref/rwcg/sql/cosmicgame/*.sql` into expert doc |
| Error code catalog | Extract `CosmicSignatureErrors.sol` + frontend decoder |
| Prize math | Summarize `cosmic-signature-game-prizes.md` + `MainPrize.sol` comments |
| RandomWalk integration | Separate beginner + expert doc from RW CG code |

## Quality bar

Every generated markdown file must include:
- Audience tag in title `(Beginner)` or `(Expert)`
- Source references section at bottom
- No secrets (private keys, internal IPs unless public)
- Factual claims tied to a source path or JSON fact

## Query-time usage

Haystack indexes all files under `$KNOWLEDGE_BASE` with tiers:
- `primary` — beginner + expert curated docs + facts JSON
- `secondary` — sources + code fallback from repos (optional)

Codex receives only the retrieved context pack, not the full repo.
