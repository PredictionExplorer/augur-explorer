"""Fetch live indexed data from the Cosmic Game backend REST API."""
from __future__ import annotations

import json
import logging
import os
from datetime import datetime, timezone
from typing import Any

import aiohttp

from knowledge.config import KNOWLEDGE_BASE
from live.api_detector import (
    bid_count_fetch_path,
    donation_fetch_paths,
    extract_round_num,
    mentions_erc20_tokens,
    mentions_nft,
    needs_backend_api,
    wants_current_round,
)

logger = logging.getLogger(__name__)


class CosmicGameApiClient:
    """Read-only HTTP client for /api/cosmicgame/* endpoints."""

    def __init__(self, base_url: str | None = None, timeout_seconds: int | None = None):
        self.base_url = (base_url or os.getenv("FAQ_BOT_API_URL") or os.getenv("COSMIC_GAME_API_URL") or "").strip()
        self.timeout_seconds = timeout_seconds or int(os.getenv("FAQ_BOT_API_TIMEOUT_SECONDS", "20"))
        self._apply_kb_defaults()

    def _apply_kb_defaults(self) -> None:
        if self.base_url:
            return
        env_facts = KNOWLEDGE_BASE / "facts" / "network-environment.json"
        try:
            if env_facts.exists():
                data = json.loads(env_facts.read_text(encoding="utf-8"))
                mainnet = data.get("networks", {}).get("arbitrum_one_mainnet", {})
                self.base_url = (mainnet.get("frontend_env") or {}).get("NEXT_PUBLIC_API_URL", "").strip()
        except Exception as exc:
            logger.warning("Failed to load KB API URL default: %s", exc)

    @property
    def is_configured(self) -> bool:
        return bool(self.base_url)

    def config_status(self) -> dict[str, str | bool]:
        return {
            "configured": self.is_configured,
            "base_url": self.base_url or "",
        }

    def _url(self, path: str) -> str:
        base = self.base_url.rstrip("/") + "/"
        return base + path.lstrip("/")

    async def _get_json(self, path: str) -> tuple[dict[str, Any] | None, str | None]:
        url = self._url(path)
        timeout = aiohttp.ClientTimeout(total=self.timeout_seconds)
        try:
            async with aiohttp.ClientSession(timeout=timeout) as session:
                async with session.get(url) as resp:
                    text = await resp.text()
                    if resp.status != 200:
                        return None, f"HTTP {resp.status} for {path}: {text[:300]}"
                    try:
                        return json.loads(text), None
                    except json.JSONDecodeError as exc:
                        return None, f"Invalid JSON from {path}: {exc}"
        except aiohttp.ClientError as exc:
            return None, f"Request failed for {path}: {exc}"
        except TimeoutError:
            return None, f"Timeout after {self.timeout_seconds}s for {path}"

    @staticmethod
    def _summarize_dashboard(data: dict[str, Any], round_num: int | None) -> list[str]:
        lines: list[str] = []
        cur_round = data.get("CurRoundNum")
        lines.append(f"Current round (CurRoundNum): {cur_round}")
        stats = data.get("CurRoundStats") or {}
        if stats and (round_num is None or stats.get("RoundNum") == round_num):
            lines.append(f"Round {stats.get('RoundNum')} TotalBids: {stats.get('TotalBids')}")
            lines.append(f"Round {stats.get('RoundNum')} NumERC20Donations: {stats.get('NumERC20Donations')}")
            lines.append(f"Round {stats.get('RoundNum')} TotalDonatedNFTs: {stats.get('TotalDonatedNFTs')}")
            lines.append(
                f"Round {stats.get('RoundNum')} TotalDonatedCount (ETH donations): "
                f"{stats.get('TotalDonatedCount')} "
                f"({stats.get('TotalDonatedAmountEth')} ETH total)"
            )
        return lines

    @staticmethod
    def _summarize_payload(label: str, path: str, data: dict[str, Any]) -> list[str]:
        lines = [f"[{label}] GET /api/cosmicgame/{path}"]

        if "DonationsERC20ByRoundSummarized" in data:
            items = data["DonationsERC20ByRoundSummarized"] or []
            lines.append(f"ERC20 donation summary rows: {len(items)}")
        if "DonationsERC20ByRoundAll" in data:
            items = data["DonationsERC20ByRoundAll"] or []
            lines.append(f"ERC20 donation events (all rows): {len(items)}")
        if "NFTDonations" in data:
            items = data["NFTDonations"] or []
            lines.append(f"NFT donations: {len(items)}")
        if "DirectCGDonations" in data:
            items = data["DirectCGDonations"] or []
            lines.append(f"Direct ETH donations to game: {len(items)}")
        if "CosmicGameDonations" in data:
            items = data["CosmicGameDonations"] or []
            lines.append(f"ETH donations with info: {len(items)}")
        if "TotalRows" in data:
            lines.append(f"Total bid rows for round {data.get('RoundNum')}: {data.get('TotalRows')}")
        if "RoundInfo" in data:
            info = data["RoundInfo"] or {}
            round_stats = info.get("RoundStats") or info
            lines.append(f"RoundInfo RoundNum: {round_stats.get('RoundNum')}")
            lines.append(f"RoundInfo TotalBids: {round_stats.get('TotalBids')}")
            lines.append(f"RoundInfo NumERC20Donations: {round_stats.get('NumERC20Donations')}")
            lines.append(f"RoundInfo TotalDonatedNFTs: {round_stats.get('TotalDonatedNFTs')}")

        if data.get("status") == 0 and data.get("error"):
            lines.append(f"API error: {data.get('error')}")

        compact = json.dumps(data, indent=2)
        if len(compact) > 6000:
            compact = compact[:6000] + "\n...(truncated)"
        lines.append(compact)
        return lines

    async def fetch_for_question(self, question: str) -> tuple[str | None, str | None, list[str]]:
        """Return (prompt_block, error_message, source_labels)."""
        if not needs_backend_api(question):
            return None, None, []

        if not self.is_configured:
            return None, "FAQ_BOT_API_URL is not set and no KB API default found", []

        q_lower = question.lower()
        round_num = extract_round_num(question)
        wants_bids = any(p in q_lower for p in ("how many bid", "number of bid", "total bid", "count bid", "num bid"))
        wants_donations = any(w in q_lower for w in ("donat", "donor")) or (
            "how many" in q_lower and any(w in q_lower for w in ("token", "nft"))
        )

        parts: list[str] = [
            f"Fetched at (UTC): {datetime.now(timezone.utc).isoformat()}",
            f"API base: {self.base_url}",
        ]
        sources: list[str] = []
        errors: list[str] = []

        async def fetch_one(label: str, path: str) -> dict[str, Any] | None:
            sources.append(f"live:api:{path}")
            data, err = await self._get_json(path)
            if err:
                errors.append(f"{path}: {err}")
                parts.append(f"[{label}] ERROR: {err}")
                return None
            return data

        dashboard = await fetch_one("dashboard", "statistics/dashboard")
        if dashboard and round_num is None:
            round_num = dashboard.get("CurRoundNum")

        if round_num is None:
            return None, "Could not determine round number", sources

        parts.append(f"Resolved round: {round_num}")
        if dashboard:
            parts.extend(self._summarize_dashboard(dashboard, round_num))

        if wants_donations or any(w in q_lower for w in ("token", "nft", "donat")):
            for label, path in donation_fetch_paths(round_num, question):
                data = await fetch_one(label, path)
                if data:
                    parts.extend(self._summarize_payload(label, path, data))

        if wants_bids:
            label, path = bid_count_fetch_path(round_num)
            data = await fetch_one(label, path)
            if data:
                parts.extend(self._summarize_payload(label, path, data))

        if not wants_donations and not wants_bids:
            label, path = f"round_info_{round_num}", f"rounds/info/{round_num}"
            data = await fetch_one(label, path)
            if data:
                parts.extend(self._summarize_payload(label, path, data))

        if errors and dashboard is None:
            return None, "; ".join(errors), sources

        block = (
            "LIVE BACKEND API DATA (indexed Cosmic Game REST API — authoritative for "
            "donation counts, bid totals, and round statistics)\n"
            + "\n".join(parts)
        )
        return block, None, sources
