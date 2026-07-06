"""Format FAQ bot answers for API responses."""
from __future__ import annotations

import re

_SOURCES_TRAILER = re.compile(
    r"\n+\s*(?:#{1,3}\s*)?\*?\*?Sources:?\*?\*?\s*[\s\S]*$",
    re.IGNORECASE,
)


def wants_detail(question: str) -> bool:
    q = question.lower()
    triggers = (
        "explain",
        "how do i",
        "how to",
        "how can i",
        "step by step",
        "steps to",
        "in detail",
        "more detail",
        "tell me more",
        "walk me through",
        "why ",
        "what is the process",
        "guide me",
        "tutorial",
        "show me how",
    )
    return any(t in q for t in triggers)


def strip_sources_section(answer: str) -> str:
    """Remove trailing Sources block; sources are returned separately in the API."""
    cleaned = _SOURCES_TRAILER.sub("", answer.strip())
    return cleaned.strip()
