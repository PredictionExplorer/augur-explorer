# Architecture — Cosmic Signature FAQ Bot

This document describes the current FAQ bot architecture: a **generated knowledge base**, **Haystack BM25 retrieval**, **conditional live data fetches**, and **Codex MCP synthesis**, exposed through a Python FastAPI backend and optionally proxied by the Go `rwcg/websrv` and embedded in the Next.js Cosmic Signature frontend.

## System overview

```
┌──────────────────────────────────────────────────────────────────────────┐
│  User browser                                                             │
│  • Production: app.cosmicsignature.com (FaqBotWidget floating panel)     │
│  • Dev/test: backend/static/test-ui.html or local Next.js                │
└───────────────────────────────┬──────────────────────────────────────────┘
                                │ POST /api/cosmicgame/faq/query  (or direct)
                                ▼
┌──────────────────────────────────────────────────────────────────────────┐
│  Go websrv proxy (optional) — rwcg/websrv/api/faq/faq.go                 │
│  • AI_BOT_BACKEND_URL → Python FAQ bot (default http://127.0.0.1:8000)   │
│  • Routes: GET /health, POST /query, POST /reindex                       │
└───────────────────────────────┬──────────────────────────────────────────┘
                                │
                                ▼
┌──────────────────────────────────────────────────────────────────────────┐
│  FastAPI backend — rwcg/faq-bot/backend/app.py                           │
│  • Session store (TTL)                                                    │
│  • Orchestrator coordinates retrieval + live fetches + Codex             │
└───────┬───────────────────────┬───────────────────────┬──────────────────┘
        │                       │                       │
        ▼                       ▼                       ▼
┌───────────────┐   ┌───────────────────────┐   ┌─────────────────────────┐
│ Haystack KB   │   │ Live data layer       │   │ Codex MCP (stdio)       │
│ BM25 retrieve │   │ • cginfo (on-chain)   │   │ codex / codex-reply     │
│ tiered docs   │   │ • Cosmic Game REST API│   │ answer synthesis        │
│ + JSON facts  │   │   (indexed stats)     │   │                         │
└───────────────┘   └───────────────────────┘   └─────────────────────────┘
        │                       │
        ▼                       ▼
┌───────────────┐   ┌───────────────────────┐
│ KNOWLEDGE_BASE│   │ External services      │
│ (generated)   │   │ • RPC (Arbitrum)       │
│ docs + facts  │   │ • a1.cosmicsignature.com│
└───────────────┘   └───────────────────────┘
```

The bot does **not** answer from raw GitHub clones at query time. Repositories under `backend/data/repos/` are used only during **knowledge generation** (and optional code fallback chunks). Answers are synthesized by Codex from retrieved context plus any injected live blocks.

## Core design

| Concern | Implementation |
|--------|----------------|
| Static knowledge | Generated markdown, JSON facts, and source extracts under `KNOWLEDGE_BASE` |
| Retrieval | Haystack `InMemoryBM25Retriever`, tiered primary/secondary docs |
| Live prices, timers, on-chain state | `cginfo` CLI subprocess (`live/cginfo_client.py`) |
| Live round stats, stakers, champions | Cosmic Game REST API (`live/api_client.py`) |
| Answer generation | Codex MCP via long-lived stdio session (`llm/codex_mcp_client.py`) |
| Multi-turn chat | In-memory sessions + Codex `threadId` per session |
| Production UI | `FaqBotWidget` in cosmicgame-frontend, proxied through Go websrv |

Legacy modules (`rag_pipeline.py`, `answer_generator.py`, `github_fetcher.py`) remain for debugging but are **not** used by `app.py`.

---

## Query flow

```
1. User sends question (+ optional session_id)
2. Orchestrator loads session; rewrites question for retrieval if follow-up
3. Haystack retrieves tiered docs + structured facts (with query-type pinning)
4. Question classifiers decide optional live injections:
   a. needs_live_state      → cginfo on-chain dump
   b. needs_backend_api     → REST API fetches (dashboard, champions, etc.)
   c. needs_backend_url_info → runtime config block (FAQ_BOT_API_URL, etc.)
5. Context pack assembled (history + facts + docs + live blocks)
6. Codex MCP synthesizes answer (read-only sandbox, approval-policy never)
7. Response returned: { answer, sources, session_id }
8. Turn appended to conversation log (FAQ_BOT_LOG_DIR)
```

### Question classifiers

Two modules route live fetches:

**`live/detector.py`** — on-chain / “current state” cues:

- Round end / countdown (`needs_round_end_time`)
- Chrono Warrior / Endurance Champion (`needs_champions_state`)
- Staker counts (`needs_staking_stats`)
- Backend/API URL meta questions (`needs_backend_url_info`)
- General live contract reads (`needs_live_state`: “current” + price/round/bid/etc.)

**`live/api_detector.py`** — indexed REST API data:

- Time-range bid counts (`needs_time_range_bids`)
- Status recap / catch-up (`needs_status_recap`)
- Donation and bid-count queries
- Delegates to the detectors above for round end, staking, champions

When a classifier matches, the orchestrator injects an authoritative block into the Codex prompt and adds prompt hints so the model uses those fields instead of saying data is missing.

### Retrieval pinning

`retrieval/pipeline.py` forces high-value documents for certain query types (addresses, ABIs, environment URLs) instead of relying on BM25 alone. Structured facts from `facts/*.json` are prioritized in the context pack for the same query types.

---

## Knowledge base

### Generation

Run from `backend/`:

```bash
python -m knowledge.generate.run_all
# or: scripts/generate-knowledge.sh
```

Generators under `backend/knowledge/generate/` produce:

| Generator | Output |
|-----------|--------|
| `extract_contracts.py` | Contract inventory facts |
| `extract_routes.py` | Frontend routes |
| `extract_api.py` | Backend API surface |
| `extract_deployments.py` | `deployed-addresses.json`, `network-environment.json`, expert deployment/env docs |
| `extract_abis.py` | Contract ABI summaries |
| `copy_source_docs.py` | Beginner/expert curated docs from monorepo sources |
| `extract_ui_pages.py` | UI page text extracts |

Seed files live in `backend/knowledge/deployments/` (e.g. `arbitrum-mainnet-environment.txt` with `NEXT_PUBLIC_API_URL`). Generated output defaults to `DATA_DIR/knowledge/generated` or an external path via `KNOWLEDGE_BASE`.

### Indexing

On startup, `KnowledgeRetriever.index_knowledge_base()` walks `KNOWLEDGE_BASE` for `.md`, `.txt`, `.json`, assigns tiers:

- **Primary:** `docs/beginner/`, `docs/expert/`, `facts/`
- **Secondary:** `docs/sources/` (+ optional Solidity/TS code fallback)

Re-index without restart: `POST /api/reindex`.

---

## Live data layer

### On-chain reads (`CginfoClient`)

- **Binary:** `rwcg/etl/cosmicgame/scripts/cginfo` (build from `cginfo.go`)
- **Config:** `RPC_URL`, `COSMIC_GAME_ADDR`, `CGINFO_PATH`
- **Used for:** current bid prices, round status, MainPrizeTime / countdown, contract parameters

### Indexed REST API (`CosmicGameApiClient`)

- **Config:** `FAQ_BOT_API_URL` (e.g. `https://a1.cosmicsignature.com/api/cosmicgame/`)
- **Fallback:** `NEXT_PUBLIC_API_URL` from `facts/network-environment.json`

Representative fetches:

| Question type | Endpoints |
|---------------|-----------|
| Round end / countdown | `rounds/current/time`, `time/until_prize`, `time/current`, `statistics/dashboard` |
| Staker counts | `statistics/dashboard` → `MainStats.StakeStatistics*` |
| Chrono Warrior / champions | `bid/current_special_winners` |
| Time-window bids | `statistics/bidding/frequency/{init}/{fin}/{interval}` |
| Donations / bid totals | `donations/*`, `bid/list/by_round/*`, `rounds/info/{n}` |

Live blocks are labeled in the prompt (e.g. `LIVE STAKING STATE`, `LIVE ROUND END STATE`) so Codex treats them as authoritative.

---

## LLM synthesis (Codex MCP)

`CodexMCPClient` spawns `codex mcp-server` over stdio and reuses one MCP session for the process lifetime. Per user session it stores a Codex `threadId` for multi-turn continuity.

**Environment (see `backend/.env.example`):**

- `CODEX_MCP_COMMAND`, `CODEX_MCP_ARGS`
- `CODEX_MCP_SANDBOX=read-only`
- `CODEX_MCP_APPROVAL_POLICY=never`
- Optional `CODEX_MCP_MODEL`

System instructions in `codex_mcp_client.py` enforce: answer only from context, brief by default, no inline Sources section, use live blocks when present.

Both Haystack **and** Codex MCP must be healthy for `/api/query` to succeed (`503` if either is down).

See also `codex-architecture.md` for MCP tool details (`codex`, `codex-reply`).

---

## Sessions and logging

**`sessions/store.py`** — in-memory chat sessions with idle TTL (`SESSION_TTL_SECONDS`, default 3600). Each session holds message history and `codex_thread_id`.

**`sessions/conversation_logger.py`** — append-only logs under `FAQ_BOT_LOG_DIR` (`{session_id}-{date}.log`).

---

## Frontend integration

### Production widget (cosmicgame-frontend)

Located in `components/faq-bot/`:

- `FaqBotWidget.tsx` — floating launcher + panel
- `FaqBotChat.tsx` — message list, markdown rendering
- `useFaqBotSession.ts` — session persistence, query calls
- `services/api/faqBot.ts` — API client

**API routing:**

1. **Default:** `getAPIUrl('faq/query')` → Go websrv `/api/cosmicgame/faq/query` → Python backend
2. **Direct dev:** `NEXT_PUBLIC_FAQ_BOT_URL=http://127.0.0.1:8000` bypasses Go proxy
3. **Next.js rewrites:** `AI_BOT_BACKEND_URL` in `next.config.ts` for local dev

Enable/disable: `NEXT_PUBLIC_ENABLE_FAQ_BOT` (default on).

### Built-in test UI

`backend/static/test-ui.html` — minimal chat page served by FastAPI for backend-only testing (`GET /` when static is mounted).

### Legacy Next.js app

`rwcg/faq-bot/app/` contains an older standalone chat UI (`ChatInterface.tsx`). The production path is the widget in cosmicgame-frontend.

---

## Go websrv proxy

**File:** `rwcg/websrv/api/faq/faq.go`

Registered from `rwcg/websrv/main.go` when FAQ routes are enabled.

| Client path | Upstream |
|-------------|----------|
| `GET /api/cosmicgame/faq/health` | `{AI_BOT_BACKEND_URL}/health` |
| `POST /api/cosmicgame/faq/query` | `{AI_BOT_BACKEND_URL}/api/query` |
| `POST /api/cosmicgame/faq/reindex` | `{AI_BOT_BACKEND_URL}/api/reindex` |

Env: `AI_BOT_BACKEND_URL` (legacy alias `FAQ_BOT_UPSTREAM_URL`). Default `http://127.0.0.1:8000`.

---

## Backend API

**File:** `backend/app.py`

| Method | Path | Purpose |
|--------|------|---------|
| GET | `/health` | Haystack, Codex, cginfo, API client, session stats |
| POST | `/api/query` | `{ question, session_id? }` → `{ answer, sources, session_id }` |
| POST | `/api/reindex` | Force Haystack re-index from `KNOWLEDGE_BASE` |
| GET | `/` | Test UI (static) |

Errors return `{ error, component }` with `component` ∈ `haystack | codex_mcp | validation`.

---

## Configuration reference

Typical production/dev env (`scripts/faq-bot.env.example`, `backend/.env`):

```bash
# Server
HOST=127.0.0.1
PORT=8000
KNOWLEDGE_BASE=/path/to/generated-kb
FAQ_BOT_LOG_DIR=~/ae_logs/ai-faq-bot

# Live indexed API (round stats, stakers, champions, donations)
FAQ_BOT_API_URL=https://a1.cosmicsignature.com/api/cosmicgame/

# On-chain reads (optional but needed for live prices/timers)
RPC_URL=https://a1.cosmicsignature.com:38546
CGINFO_PATH=/path/to/rwcg/etl/cosmicgame/scripts/cginfo

# Codex MCP (requires codex login / auth on host)
CODEX_MCP_COMMAND=codex
CODEX_MCP_ARGS=mcp-server
CODEX_MCP_SANDBOX=read-only
CODEX_MCP_APPROVAL_POLICY=never
```

Start script: `scripts/start-backend.sh` (loads `~/configs/faq-bot.env` if present).

---

## Key file map

```
rwcg/faq-bot/
├── ARCHITECTURE.md          ← this file
├── codex-architecture.md    ← Codex MCP protocol notes
├── scripts/
│   ├── start-backend.sh
│   ├── generate-knowledge.sh
│   └── faq-bot.env.example
└── backend/
    ├── app.py               ← FastAPI entrypoint
    ├── orchestrator.py      ← retrieval + live + Codex coordination
    ├── response_format.py   ← answer post-processing
    ├── retrieval/
    │   ├── pipeline.py      ← Haystack indexer + tiered retrieve + pinning
    │   └── context_pack.py  ← token-budgeted prompt assembly
    ├── live/
    │   ├── detector.py      ← live/cginfo question detection
    │   ├── api_detector.py  ← REST API question detection
    │   ├── api_client.py    ← Cosmic Game HTTP client
    │   ├── cginfo_client.py ← cginfo subprocess wrapper
    │   └── time_range.py    ← calendar bid-count parsing
    ├── llm/
    │   └── codex_mcp_client.py
    ├── sessions/
    │   ├── store.py
    │   └── conversation_logger.py
    └── knowledge/
        ├── generate/        ← KB generators
        └── deployments/     ← seed deployment/env .txt files

rwcg/websrv/api/faq/faq.go   ← Go reverse proxy

cosmicgame-frontend/
└── components/faq-bot/      ← production widget
```

---

## Design decisions

### Why Haystack BM25 (not embeddings)?

Fast to index, no embedding API, works well for exact protocol terms (contract names, routes, env vars). Query-type pinning compensates when keyword overlap is weak (e.g. “URL for the backend” vs `NEXT_PUBLIC_API_URL`).

### Why Codex MCP (not a custom answer generator)?

Natural-language synthesis over heterogeneous context (docs + JSON facts + live API blocks + on-chain dumps). Multi-turn via `threadId`. Constrained with read-only sandbox and strict system instructions.

### Why conditional live fetches?

Most FAQ answers come from static docs. Live calls (RPC, REST) are expensive and only run when classifiers detect the question needs current state — avoiding unnecessary latency and keeping answers authoritative when counts/timers matter.

### Why a separate knowledge base directory?

Generated content is large, versioned independently, and can live outside the git tree (`KNOWLEDGE_BASE`). Code changes do not require re-cloning repos on every backend restart — only re-run generators when sources change, then `POST /api/reindex`.

---

## Operations

### Startup checklist

1. Generate or update KB: `python -m knowledge.generate.run_all`
2. Ensure `KNOWLEDGE_BASE` points at generated output
3. Configure `FAQ_BOT_API_URL`, `RPC_URL`, Codex auth
4. Build `cginfo` if using on-chain live state
5. Start Python backend: `scripts/start-backend.sh`
6. (Optional) Start Go websrv with `AI_BOT_BACKEND_URL`
7. Verify: `GET /health` → `"status": "healthy"`

### Health/degraded modes

| Component missing | Effect |
|-------------------|--------|
| Knowledge base not indexed | Startup failure |
| Codex MCP down | `/api/query` → 503 |
| cginfo not configured | Live on-chain questions may lack data; static + API answers still work |
| FAQ_BOT_API_URL unset | Live stats questions fail; KB-only answers still work |

### Performance (typical)

- Retrieval: tens of ms
- Live API/cginfo: hundreds of ms to a few seconds
- Codex synthesis: several seconds to minutes (dominant latency)

---

## Future improvements

- Persistent vector or hybrid retrieval for semantic paraphrases
- Cache dashboard/champions responses with short TTL
- Structured output / tool calling for numeric answers (staker counts, timers)
- Horizontal scale: multiple Python workers + sticky sessions or stateless Codex threads per request
- Rate limiting and auth on `/api/cosmicgame/faq/*`

---

This architecture separates **static knowledge** (generated KB), **live truth** (on-chain + indexed API), and **language synthesis** (Codex), with explicit question routing so each layer is used when it adds value.
