from __future__ import annotations

from dataclasses import dataclass, field


@dataclass
class ApiFetchResult:
    """Result of live Cosmic Game API fetches for a user question."""

    block: str | None = None
    error: str | None = None
    sources: list[str] = field(default_factory=list)
    clarification_answer: str | None = None
