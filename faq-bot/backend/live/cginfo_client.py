"""Run cginfo to fetch live CosmicGame on-chain state."""
from __future__ import annotations

import asyncio
import json
import logging
import os
from datetime import datetime, timezone
from pathlib import Path

from knowledge.config import CURSOR_VREF_PATH, KNOWLEDGE_BASE

logger = logging.getLogger(__name__)

DEFAULT_CGINFO = (
    CURSOR_VREF_PATH / "rwcg" / "etl" / "cosmicgame" / "scripts" / "cginfo"
)


class CginfoClient:
    """Read-only wrapper around the cginfo CLI."""

    def __init__(
        self,
        cginfo_path: str | Path | None = None,
        rpc_url: str | None = None,
        game_address: str | None = None,
        timeout_seconds: int | None = None,
    ):
        raw_path = (cginfo_path or os.getenv("CGINFO_PATH") or str(DEFAULT_CGINFO)).strip()
        self.cginfo_path = Path(os.path.expanduser(raw_path))
        self.rpc_url = (rpc_url or os.getenv("FAQ_BOT_RPC_URL") or os.getenv("RPC_URL") or "").strip()
        self.game_address = (
            game_address or os.getenv("COSMIC_GAME_ADDR") or os.getenv("COSMIC_GAME_ADDRESS") or ""
        ).strip()
        self.timeout_seconds = timeout_seconds or int(os.getenv("CGINFO_TIMEOUT_SECONDS", "45"))
        self._apply_kb_defaults()

    def _apply_kb_defaults(self) -> None:
        if self.rpc_url and self.game_address:
            return
        env_facts = KNOWLEDGE_BASE / "facts" / "network-environment.json"
        addr_facts = KNOWLEDGE_BASE / "facts" / "deployed-addresses.json"
        try:
            if not self.rpc_url and env_facts.exists():
                data = json.loads(env_facts.read_text(encoding="utf-8"))
                mainnet = data.get("networks", {}).get("arbitrum_one_mainnet", {})
                self.rpc_url = (mainnet.get("frontend_env") or {}).get("NEXT_PUBLIC_RPC_URL", "").strip()
            if not self.game_address and addr_facts.exists():
                data = json.loads(addr_facts.read_text(encoding="utf-8"))
                addrs = (
                    data.get("networks", {})
                    .get("arbitrum_one_mainnet", {})
                    .get("addresses", {})
                )
                self.game_address = (
                    addrs.get("CosmicSignatureGame proxy")
                    or addrs.get("CosmicSignatureGame")
                    or ""
                ).strip()
        except Exception as exc:
            logger.warning("Failed to load KB defaults for cginfo: %s", exc)

    @property
    def is_configured(self) -> bool:
        return (
            self.cginfo_path.is_file()
            and os.access(self.cginfo_path, os.X_OK)
            and bool(self.rpc_url)
            and bool(self.game_address)
        )

    def config_status(self) -> dict[str, str | bool]:
        return {
            "configured": self.is_configured,
            "cginfo_path": str(self.cginfo_path),
            "cginfo_exists": self.cginfo_path.is_file(),
            "rpc_url_set": bool(self.rpc_url),
            "game_address": self.game_address or "",
        }

    async def fetch_state(self) -> tuple[str | None, str | None]:
        """Return (stdout, error_message)."""
        if not self.cginfo_path.is_file():
            return None, f"cginfo binary not found at {self.cginfo_path}"
        if not os.access(self.cginfo_path, os.X_OK):
            return None, f"cginfo is not executable: {self.cginfo_path}"
        if not self.rpc_url:
            return None, "RPC_URL (or FAQ_BOT_RPC_URL) is not set and no KB RPC default found"
        if not self.game_address:
            return None, "COSMIC_GAME_ADDR is not set and no KB proxy address found"

        env = os.environ.copy()
        env["RPC_URL"] = self.rpc_url

        try:
            proc = await asyncio.create_subprocess_exec(
                str(self.cginfo_path),
                self.game_address,
                stdout=asyncio.subprocess.PIPE,
                stderr=asyncio.subprocess.STDOUT,
                env=env,
            )
            stdout, _ = await asyncio.wait_for(
                proc.communicate(),
                timeout=self.timeout_seconds,
            )
        except asyncio.TimeoutError:
            return None, f"cginfo timed out after {self.timeout_seconds}s"
        except Exception as exc:
            return None, f"cginfo failed to start: {exc}"

        text = (stdout or b"").decode("utf-8", errors="replace").strip()
        if proc.returncode != 0:
            detail = text or f"exit code {proc.returncode}"
            return None, f"cginfo error: {detail[:500]}"

        fetched_at = datetime.now(timezone.utc).isoformat()
        header = (
            f"Fetched at (UTC): {fetched_at}\n"
            f"RPC URL: {self.rpc_url}\n"
            f"CosmicSignatureGame proxy: {self.game_address}\n"
        )
        return header + text, None
