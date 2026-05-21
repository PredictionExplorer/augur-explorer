"""In-memory per-user chat sessions."""
from __future__ import annotations

import time
import uuid
from dataclasses import dataclass, field
from typing import Any


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
    def __init__(self, ttl_seconds: int = 86400):
        self.ttl_seconds = ttl_seconds
        self._sessions: dict[str, ChatSession] = {}

    def _purge_expired(self) -> None:
        now = time.time()
        expired = [
            sid for sid, s in self._sessions.items() if now - s.last_active > self.ttl_seconds
        ]
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
        return self._sessions.get(session_id)

    def get_or_create(self, session_id: str | None) -> ChatSession:
        session = self.get(session_id)
        if session:
            return session
        return self.create()

    def stats(self) -> dict[str, Any]:
        self._purge_expired()
        return {"active_sessions": len(self._sessions)}
