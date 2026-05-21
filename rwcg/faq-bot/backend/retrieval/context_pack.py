"""Build token-budgeted context packs for Codex."""
from __future__ import annotations

import json
from typing import Any

from haystack import Document


def _truncate(text: str, max_chars: int) -> str:
    if len(text) <= max_chars:
        return text
    return text[: max_chars - 20] + "\n...(truncated)"


def build_context_pack(
    question: str,
    documents: list[Document],
    facts: dict[str, Any] | None = None,
    history: list[dict[str, str]] | None = None,
    max_chars: int = 12000,
) -> tuple[str, list[str]]:
    """Return (prompt_text, source_labels)."""
    parts: list[str] = []
    sources: list[str] = []
    budget = max_chars

    if history:
        hist_lines = []
        for msg in history[-6:]:
            hist_lines.append(f"{msg['role'].upper()}: {msg['content'][:400]}")
        block = "CONVERSATION HISTORY\n" + "\n".join(hist_lines) + "\n\n"
        parts.append(block)
        budget -= len(block)

    if facts:
        facts_block = "STRUCTURED FACTS\n" + _truncate(json.dumps(facts, indent=2), min(3000, budget // 3)) + "\n\n"
        parts.append(facts_block)
        budget -= len(facts_block)

    parts.append("RETRIEVED DOCUMENTS\n")
    budget -= len(parts[-1])

    for i, doc in enumerate(documents, 1):
        tier = doc.meta.get("tier", "unknown")
        path = doc.meta.get("file_path", "unknown")
        source = doc.meta.get("source", "unknown")
        label = f"{source}:{path}"
        chunk_budget = max(400, budget // max(1, len(documents) - i + 1))
        section = (
            f"--- [{i}] tier={tier} source={source} path={path} ---\n"
            f"{_truncate(doc.content, chunk_budget)}\n\n"
        )
        if len(section) > budget:
            section = section[:budget]
        parts.append(section)
        sources.append(label)
        budget -= len(section)
        if budget <= 200:
            break

    parts.append(f"USER QUESTION\n{question}\n")
    return "".join(parts), sources
