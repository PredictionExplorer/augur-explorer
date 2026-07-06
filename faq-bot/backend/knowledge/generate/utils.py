"""Shared helpers for knowledge extraction."""
from __future__ import annotations

import json
import re
from pathlib import Path
from typing import Any, Iterable


def write_json(path: Path, data: Any) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(data, indent=2, sort_keys=True) + "\n", encoding="utf-8")


def write_text(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content.rstrip() + "\n", encoding="utf-8")


def read_text(path: Path) -> str:
    return path.read_text(encoding="utf-8", errors="replace")


def strip_tsx_to_text(content: str) -> str:
    """Best-effort extraction of user-visible text from TSX/JSX."""
    text = content
    text = re.sub(r"^\s*import .+$", "", text, flags=re.MULTILINE)
    text = re.sub(r"^\s*export default .+$", "", text, flags=re.MULTILINE)
    text = re.sub(r"\{/\*.*?\*/\}", " ", text, flags=re.DOTALL)
    text = re.sub(r"/\*.*?\*/", " ", text, flags=re.DOTALL)
    text = re.sub(r"//.*", " ", text)
    text = re.sub(r"className=\{[^}]+\}", " ", text)
    text = re.sub(r"className=\"[^\"]*\"", " ", text)
    text = re.sub(r"<\s*[A-Z][A-Za-z0-9]*[^>]*?/>", " ", text)
    text = re.sub(r"<\s*/?\s*[a-zA-Z][^>]*>", " ", text)
    strings = re.findall(r"['\"`]([^'\"`\n]{3,})['\"`]", text)
    cleaned = []
    for s in strings:
        s = s.strip()
        if not s:
            continue
        if s.startswith(("http://", "https://", "0x", "/")):
            cleaned.append(s)
            continue
        if re.fullmatch(r"[a-z0-9_\-./]+", s):
            continue
        if any(token in s for token in ("{", "}", "=>", "function", "const ", "return")):
            continue
        cleaned.append(s)
    return "\n".join(dict.fromkeys(cleaned))


def iter_files(root: Path, suffixes: Iterable[str]) -> list[Path]:
    if not root.exists():
        return []
    suffix_set = set(suffixes)
    return sorted(
        p for p in root.rglob("*") if p.is_file() and p.suffix in suffix_set
    )


def relative(path: Path, base: Path) -> str:
    try:
        return str(path.relative_to(base))
    except ValueError:
        return str(path)
