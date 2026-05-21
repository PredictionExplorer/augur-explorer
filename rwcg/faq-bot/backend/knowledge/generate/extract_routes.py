"""Extract frontend routes and user-facing page summaries."""
from __future__ import annotations

import re
from pathlib import Path

from knowledge.config import DOCS_BEGINNER_DIR, DOCS_EXPERT_DIR, FACTS_DIR, REPO_NAMES
from knowledge.generate.utils import read_text, relative, write_json, write_text


APP_DIR = REPO_NAMES["frontend"] / "src" / "app"

ROUTE_DESCRIPTIONS = {
    "/": "Landing page",
    "/game/play": "Place bids during an active round",
    "/game/prizes": "Prize types, distribution, and rewards",
    "/game/how-it-works": "Beginner explanation of the game",
    "/game/leaderboard": "Top players and rankings",
    "/game/history/rounds": "Past rounds",
    "/game/statistics": "Game analytics and charts",
    "/stake": "Stake Cosmic Signature or RandomWalk NFTs",
    "/account": "Account overview",
    "/account/nfts": "View owned NFTs",
    "/account/winnings": "Prize winnings and claims",
    "/account/activity": "Bid and activity history",
    "/gallery": "Browse Cosmic Signature NFT gallery",
    "/contracts": "On-chain contract addresses and info",
    "/about": "About the project",
    "/contributions": "Community contributions",
}


def _page_path_to_route(page_path: Path) -> str:
    rel = page_path.parent.relative_to(APP_DIR)
    parts = list(rel.parts)
    if parts == ["."]:
        return "/"
    route_parts = []
    for part in parts:
        if part.startswith("[") and part.endswith("]"):
            route_parts.append(f":{part[1:-1]}")
        else:
            route_parts.append(part)
    return "/" + "/".join(route_parts)


def extract_routes() -> dict:
    frontend = REPO_NAMES["frontend"]
    routes = []
    if not APP_DIR.exists():
        return {"routes": [], "warning": f"Frontend app dir missing: {APP_DIR}"}

    for page in sorted(APP_DIR.rglob("page.tsx")):
        route = _page_path_to_route(page)
        text = strip_page_text(read_text(page))
        routes.append(
            {
                "route": route,
                "file": relative(page, frontend),
                "description": ROUTE_DESCRIPTIONS.get(route, "Application page"),
                "extracted_text_preview": text[:500],
            }
        )
    return {"routes": routes, "total_routes": len(routes)}


def strip_page_text(content: str) -> str:
    from knowledge.generate.utils import strip_tsx_to_text

    return strip_tsx_to_text(content)


def generate_beginner_navigation_doc(routes_data: dict) -> str:
    lines = [
        "# Site Navigation (Beginner)",
        "",
        "Cosmic Signature is a web app for a blockchain bidding game. These are the main pages:",
        "",
    ]
    priority = [
        "/game/how-it-works",
        "/game/play",
        "/game/prizes",
        "/stake",
        "/account/nfts",
        "/account/winnings",
        "/gallery",
        "/contracts",
    ]
    by_route = {r["route"]: r for r in routes_data.get("routes", [])}
    lines.append("## Start here")
    for route in priority:
        item = by_route.get(route)
        if not item:
            continue
        lines.append(f"- **{route}** — {item['description']}")
    lines += ["", "## All routes"]
    for item in routes_data.get("routes", []):
        lines.append(f"- `{item['route']}` — {item['description']}")
    return "\n".join(lines)


def generate_expert_routes_doc(routes_data: dict) -> str:
    lines = [
        "# Frontend Routes (Expert)",
        "",
        f"Total Next.js app routes discovered: **{routes_data.get('total_routes', 0)}**",
        "",
    ]
    for item in routes_data.get("routes", []):
        lines += [
            f"## `{item['route']}`",
            f"- File: `{item['file']}`",
            f"- Description: {item['description']}",
            "",
        ]
    return "\n".join(lines)


def run() -> None:
    routes_data = extract_routes()
    write_json(FACTS_DIR / "frontend-routes.json", routes_data)
    write_text(DOCS_BEGINNER_DIR / "05-site-navigation.md", generate_beginner_navigation_doc(routes_data))
    write_text(DOCS_EXPERT_DIR / "06-frontend-routes.md", generate_expert_routes_doc(routes_data))
