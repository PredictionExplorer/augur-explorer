"""Append-only conversation logs per session."""
from __future__ import annotations

import logging
import os
from datetime import datetime, timezone
from pathlib import Path

from sessions.store import ChatSession

logger = logging.getLogger(__name__)

PROMPT_SEPARATOR = "============================================="
ANSWER_SEPARATOR = "---------------------------------------------"


class ConversationLogger:
    def __init__(self, log_dir: str | None = None):
        raw = (log_dir or os.getenv("FAQ_BOT_LOG_DIR") or "~/ae_logs/ai-faq-bot").strip()
        self.log_dir = Path(os.path.expanduser(raw))
        self.log_dir.mkdir(parents=True, exist_ok=True)

    def log_path(self, session: ChatSession) -> Path:
        date = datetime.fromtimestamp(session.created_at, tz=timezone.utc).strftime("%Y-%m-%d")
        return self.log_dir / f"{session.session_id}-{date}.log"

    def log_turn(self, session: ChatSession, prompt: str, response: str) -> Path | None:
        path = self.log_path(session)
        block = (
            f"{PROMPT_SEPARATOR}\n"
            f"{prompt.rstrip()}\n\n"
            f"{ANSWER_SEPARATOR}\n"
            f"{response.rstrip()}\n\n"
        )
        try:
            if not path.exists():
                started = datetime.fromtimestamp(session.created_at, tz=timezone.utc).isoformat()
                header = f"Session: {session.session_id}\nStarted (UTC): {started}\n\n"
                path.write_text(header, encoding="utf-8")
            with path.open("a", encoding="utf-8") as fh:
                fh.write(block)
            logger.info("Conversation log appended: %s", path)
            return path
        except OSError as exc:
            logger.error("Failed to write conversation log %s: %s", path, exc)
            return None
