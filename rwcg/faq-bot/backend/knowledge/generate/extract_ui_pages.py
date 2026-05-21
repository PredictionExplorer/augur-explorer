"""Extract UI page text chunks for Haystack indexing."""
from __future__ import annotations

from knowledge.config import DOCS_SOURCES_DIR, REPO_NAMES
from knowledge.generate.utils import read_text, relative, strip_tsx_to_text, write_text

PRIORITY_PAGES = [
    "src/app/game/how-it-works/page.tsx",
    "src/app/game/play/page.tsx",
    "src/app/game/prizes/page.tsx",
    "src/app/stake/page.tsx",
    "src/app/contracts/page.tsx",
    "src/app/about/page.tsx",
]


def run() -> None:
    frontend = REPO_NAMES["frontend"]
    out_dir = DOCS_SOURCES_DIR / "frontend-pages"
    for rel in PRIORITY_PAGES:
        path = frontend / rel
        if not path.exists():
            continue
        text = strip_tsx_to_text(read_text(path))
        name = rel.replace("/", "_").replace(".tsx", ".txt")
        write_text(
            out_dir / name,
            f"Source: {relative(path, frontend)}\nRoute hint: /{rel.split('/app/')[-1].replace('/page.tsx', '')}\n\n{text}",
        )
