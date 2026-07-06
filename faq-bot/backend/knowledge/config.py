"""Paths and settings for the knowledge base generator."""
from __future__ import annotations

import os
from pathlib import Path

from dotenv import load_dotenv

BACKEND_ROOT = Path(__file__).resolve().parents[1]
FAQ_BOT_ROOT = BACKEND_ROOT.parent
RWCG_ROOT = FAQ_BOT_ROOT.parent  # .../cursor-vref/rwcg

# Load backend/.env so KNOWLEDGE_BASE is available to generators and retrieval.
load_dotenv(BACKEND_ROOT / ".env")

DATA_DIR = Path(os.getenv("DATA_DIR", BACKEND_ROOT / "data"))
REPOS_DIR = DATA_DIR / "repos"

# Generated FAQ text files (facts, docs) — separate from application code.
# Set KNOWLEDGE_BASE to an absolute path outside the faq-bot tree.
_raw_kb = os.getenv("KNOWLEDGE_BASE", "").strip()
KNOWLEDGE_BASE = Path(_raw_kb) if _raw_kb else DATA_DIR / "knowledge" / "generated"

# Alias kept for existing imports
GENERATED_DIR = KNOWLEDGE_BASE

FACTS_DIR = KNOWLEDGE_BASE / "facts"
DOCS_BEGINNER_DIR = KNOWLEDGE_BASE / "docs" / "beginner"
DOCS_EXPERT_DIR = KNOWLEDGE_BASE / "docs" / "expert"
DOCS_SOURCES_DIR = KNOWLEDGE_BASE / "docs" / "sources"
DEPLOYMENTS_SOURCE_DIR = DOCS_SOURCES_DIR / "deployments"

CURSOR_VREF_PATH = Path(
    os.getenv("CURSOR_VREF_PATH", str(RWCG_ROOT.parent))
)

REPO_NAMES = {
    "smart_contracts": REPOS_DIR / "smart_contracts",
    "backend_api": REPOS_DIR / "backend_api",
    "frontend": REPOS_DIR / "frontend",
}

GITHUB_REPOS = {
    "smart_contracts": "https://github.com/PredictionExplorer/Cosmic-Signature",
    "backend_api": "https://github.com/PredictionExplorer/augur-explorer",
    "frontend": "https://github.com/PredictionExplorer/cosmic-front-alternate",
}


def ensure_output_dirs() -> None:
    for path in (
        KNOWLEDGE_BASE,
        FACTS_DIR,
        DOCS_BEGINNER_DIR,
        DOCS_EXPERT_DIR,
        DOCS_SOURCES_DIR,
        DEPLOYMENTS_SOURCE_DIR,
    ):
        path.mkdir(parents=True, exist_ok=True)
