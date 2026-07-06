"""In-memory per-user chat sessions."""
from __future__ import annotations

import time
import uuid
from dataclasses import dataclass, field
from typing import Any


class SessionExpiredError(RuntimeError):
    """Raised when a client reuses a session_id that has expired due to idle time."""


@dataclass
class ChatSession:
    session_id: str
    codex_thread_id: str | None = None
    messages: list[dict[str, str]] = field(default_factory=list)
    created_at: float = field(default_factory=time.time)
    last_active: float = field(default_factory=time.time)

    def touch(self) -> None:
        self.last_active = time.time()

    def add_message(self, role: str, content: str) -> None:
        self.messages.append({"role": role, "content": content})
        self.touch()


class SessionStore:
    def __init__(self, ttl_seconds: int = 3600):
        self.ttl_seconds = ttl_seconds
        self._sessions: dict[str, ChatSession] = {}

    def _is_expired(self, session: ChatSession) -> bool:
        return time.time() - session.last_active > self.ttl_seconds

    def _purge_expired(self) -> None:
        expired = [sid for sid, s in self._sessions.items() if self._is_expired(s)]
        for sid in expired:
            del self._sessions[sid]

    def create(self) -> ChatSession:
        self._purge_expired()
        session = ChatSession(session_id=str(uuid.uuid4()))
        self._sessions[session.session_id] = session
        return session

    def get(self, session_id: str | None) -> ChatSession | None:
        self._purge_expired()
        if not session_id:
            return None
        session = self._sessions.get(session_id)
        if session and self._is_expired(session):
            del self._sessions[session_id]
            return None
        return session

    def get_or_create(self, session_id: str | None) -> ChatSession:
        """Use only when starting a new conversation (no session_id)."""
        session = self.get(session_id)
        if session:
            return session
        return self.create()

    def require(self, session_id: str | None) -> ChatSession:
        """Return an existing session or raise SessionExpiredError."""
        if not session_id:
            return self.create()
        session = self.get(session_id)
        if session is None:
            raise SessionExpiredError(
                "error: context expired, please reload window"
            )
        return session

    def stats(self) -> dict[str, Any]:
        self._purge_expired()
        return {
            "active_sessions": len(self._sessions),
            "ttl_seconds": self.ttl_seconds,
        }
