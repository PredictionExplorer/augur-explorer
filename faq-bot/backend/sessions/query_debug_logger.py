"""Verbose per-query debug logs alongside conversation transcripts."""
from __future__ import annotations

import json
import logging
import os
from datetime import datetime, timezone
from pathlib import Path
from typing import Any

from sessions.store import ChatSession

logger = logging.getLogger(__name__)

TURN_SEPARATOR = "=" * 80


class QueryDebugLogger:
    """Append-only debug trace: {session_id}-{date}.debug.log"""

    def __init__(self, log_dir: str | None = None):
        raw = (log_dir or os.getenv("FAQ_BOT_LOG_DIR") or "~/ae_logs/ai-faq-bot").strip()
        self.log_dir = Path(os.path.expanduser(raw))
        self.log_dir.mkdir(parents=True, exist_ok=True)

    def debug_path(self, session: ChatSession) -> Path:
        date = datetime.fromtimestamp(session.created_at, tz=timezone.utc).strftime("%Y-%m-%d")
        return self.log_dir / f"{session.session_id}-{date}.debug.log"

    def log_turn(self, session: ChatSession, data: dict[str, Any]) -> Path | None:
        path = self.debug_path(session)
        now = datetime.now(timezone.utc).isoformat()
        turn = data.get("turn", "?")
        lines = [
            TURN_SEPARATOR,
            f"QUERY DEBUG — turn {turn} — {now}",
            TURN_SEPARATOR,
            "",
        ]
        lines.extend(self._format_section("INPUT", data.get("input", {})))
        lines.extend(self._format_section("CLASSIFIERS", data.get("classifiers", {})))
        lines.extend(self._format_section("CGINFO", data.get("cginfo", {})))
        lines.extend(self._format_section("COSMIC GAME API", data.get("api", {})))
        lines.extend(self._format_section("HAYSTACK RETRIEVAL", data.get("retrieval", {})))
        lines.extend(self._format_section("PROMPT ASSEMBLY", data.get("prompt", {})))
        lines.extend(self._format_section("CODEX MCP", data.get("codex", {})))
        lines.extend(self._format_section("OUTPUT", data.get("output", {})))
        lines.extend(self._format_section("TIMING (ms)", data.get("timing", {})))
        if notes := data.get("notes"):
            lines.extend(self._format_section("NOTES", notes if isinstance(notes, dict) else {"info": notes}))
        lines.append("")
        block = "\n".join(lines)

        try:
            if not path.exists():
                started = datetime.fromtimestamp(session.created_at, tz=timezone.utc).isoformat()
                header = (
                    f"Session: {session.session_id}\n"
                    f"Started (UTC): {started}\n"
                    f"Debug log: {path.name}\n"
                    f"Transcript: {session.session_id}-{datetime.fromtimestamp(session.created_at, tz=timezone.utc).strftime('%Y-%m-%d')}.log\n\n"
                )
                path.write_text(header, encoding="utf-8")
            with path.open("a", encoding="utf-8") as fh:
                fh.write(block)
            logger.info("Query debug log appended: %s", path)
            return path
        except OSError as exc:
            logger.error("Failed to write query debug log %s: %s", path, exc)
            return None

    @staticmethod
    def _format_section(title: str, payload: dict[str, Any]) -> list[str]:
        if not payload:
            return [f"--- {title} ---", "(empty)", ""]
        lines = [f"--- {title} ---"]
        for key, value in payload.items():
            lines.append(f"{key}: {QueryDebugLogger._format_value(value)}")
        lines.append("")
        return lines

    @staticmethod
    def _format_value(value: Any) -> str:
        if value is None:
            return "(none)"
        if isinstance(value, bool):
            return str(value).lower()
        if isinstance(value, (int, float)):
            return str(value)
        if isinstance(value, str):
            if "\n" in value:
                indented = value.replace("\n", "\n  ")
                return f"\n  {indented}"
            return value
        if isinstance(value, (list, tuple)):
            if not value:
                return "[]"
            if all(isinstance(v, str) for v in value):
                return json.dumps(list(value), ensure_ascii=False)
            return json.dumps(value, ensure_ascii=False, indent=2, default=str)
        if isinstance(value, dict):
            return json.dumps(value, ensure_ascii=False, indent=2, default=str)
        return repr(value)

    @staticmethod
    def preview(text: str, max_chars: int = 4000) -> str:
        text = (text or "").strip()
        if len(text) <= max_chars:
            return text
        return f"{text[:max_chars]}\n... [truncated, total {len(text)} chars]"
