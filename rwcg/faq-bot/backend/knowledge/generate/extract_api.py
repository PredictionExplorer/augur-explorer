"""Extract API endpoint summaries from BACKEND.md."""
from __future__ import annotations

import re

from knowledge.config import CURSOR_VREF_PATH, DOCS_EXPERT_DIR, FACTS_DIR, REPO_NAMES
from knowledge.generate.utils import read_text, write_json, write_text


def _find_backend_md() -> str | None:
    candidates = [
        CURSOR_VREF_PATH / "rwcg/docs/BACKEND.md",
        REPO_NAMES["backend_api"] / "rwcg/docs/BACKEND.md",
    ]
    for path in candidates:
        if path.exists():
            return read_text(path)
    return None


def extract_api_endpoints() -> dict:
    content = _find_backend_md()
    if not content:
        return {"endpoints": [], "warning": "BACKEND.md not found"}

    endpoints = []
    # Match markdown headings or bullet lines that look like REST routes
    for line in content.splitlines():
        m = re.search(r"`?(GET|POST|PUT|DELETE|PATCH)\s+(/[^`\s]+)`?", line, re.I)
        if m:
            endpoints.append({"method": m.group(1).upper(), "path": m.group(2), "context": line.strip()[:200]})
        m2 = re.search(r"(/api/[a-zA-Z0-9_./\-{}:]+)", line)
        if m2 and not any(e["path"] == m2.group(1) for e in endpoints):
            endpoints.append({"method": "GET", "path": m2.group(1), "context": line.strip()[:200]})

    # Also capture handler function names from audit-style docs
    handler_names = sorted(set(re.findall(r"api_cosmic_game_[a-z_]+", content)))
    return {
        "endpoints": endpoints[:200],
        "handler_functions": handler_names,
        "source": "rwcg/docs/BACKEND.md",
        "total_endpoints_found": len(endpoints),
    }


def generate_expert_api_doc(api_data: dict) -> str:
    lines = [
        "# Backend API Overview (Expert)",
        "",
        f"Discovered **{api_data.get('total_endpoints_found', 0)}** route-like entries from BACKEND.md.",
        "",
        "## Notable handlers",
    ]
    for name in api_data.get("handler_functions", [])[:40]:
        lines.append(f"- `{name}`")
    lines += ["", "## Routes"]
    for ep in api_data.get("endpoints", [])[:60]:
        lines.append(f"- `{ep['method']} {ep['path']}`")
    lines += [
        "",
        "## Bidding note",
        "Bids are submitted on-chain only. Backend exposes dashboard, prices, bid history, balances, and statistics.",
    ]
    return "\n".join(lines)


def run() -> None:
    api_data = extract_api_endpoints()
    write_json(FACTS_DIR / "api-endpoints.json", api_data)
    write_text(DOCS_EXPERT_DIR / "03-backend-api.md", generate_expert_api_doc(api_data))
