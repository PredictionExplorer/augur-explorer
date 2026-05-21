"""Haystack tiered retrieval over the generated knowledge base."""
from __future__ import annotations

import json
import logging
from pathlib import Path
from typing import Any

from haystack import Document
from haystack.components.retrievers.in_memory import InMemoryBM25Retriever
from haystack.document_stores.in_memory import InMemoryDocumentStore

from knowledge.config import KNOWLEDGE_BASE, REPOS_DIR
from retrieval.context_pack import build_context_pack

logger = logging.getLogger(__name__)

TIER_PRIMARY = "primary"
TIER_SECONDARY = "secondary"

PRIMARY_PREFIXES = (
    "docs/beginner/",
    "docs/expert/",
    "facts/",
)

SECONDARY_PREFIXES = (
    "docs/sources/",
)


class KnowledgeRetriever:
    """Indexes generated knowledge and optional code fallback."""

    def __init__(self, knowledge_dir: Path | None = None, include_code_fallback: bool = True):
        self.knowledge_dir = Path(knowledge_dir) if knowledge_dir else KNOWLEDGE_BASE
        self.include_code_fallback = include_code_fallback
        self.document_store = InMemoryDocumentStore()
        self.retriever: InMemoryBM25Retriever | None = None
        self._facts_cache: dict[str, Any] = {}

    @property
    def is_ready(self) -> bool:
        return self.document_store.count_documents() > 0 and self.retriever is not None

    def load_facts(self) -> dict[str, Any]:
        facts_dir = self.knowledge_dir / "facts"
        combined: dict[str, Any] = {}
        if not facts_dir.exists():
            return combined
        for path in sorted(facts_dir.glob("*.json")):
            try:
                combined[path.stem] = json.loads(path.read_text(encoding="utf-8"))
            except Exception as exc:
                logger.warning("Failed to load facts %s: %s", path, exc)
        self._facts_cache = combined
        return combined

    def _tier_for_path(self, rel: str) -> str:
        rel_norm = rel.replace("\\", "/")
        if rel_norm.startswith(PRIMARY_PREFIXES):
            return TIER_PRIMARY
        if rel_norm.startswith(SECONDARY_PREFIXES):
            return TIER_SECONDARY
        return TIER_SECONDARY

    def _content_type_for_path(self, rel: str) -> str:
        rel_norm = rel.replace("\\", "/")
        if rel_norm.startswith("facts/"):
            return "facts"
        if rel_norm.endswith(".md"):
            return "doc"
        if rel_norm.endswith(".txt"):
            return "extracted_ui"
        return "other"

    def _chunk_markdown(self, content: str, max_chars: int = 2500) -> list[str]:
        if len(content) <= max_chars:
            return [content]
        chunks = []
        start = 0
        while start < len(content):
            end = min(len(content), start + max_chars)
            if end < len(content):
                nl = content.rfind("\n", start, end)
                if nl > start + 500:
                    end = nl
            chunks.append(content[start:end])
            start = end
        return chunks

    def _documents_from_file(self, path: Path, source: str) -> list[Document]:
        rel = str(path.relative_to(self.knowledge_dir)).replace("\\", "/")
        tier = self._tier_for_path(rel)
        ctype = self._content_type_for_path(rel)
        try:
            text = path.read_text(encoding="utf-8", errors="replace")
        except Exception as exc:
            logger.warning("Skip unreadable %s: %s", path, exc)
            return []
        if not text.strip():
            return []

        docs = []
        chunks = self._chunk_markdown(text) if path.suffix == ".md" else [text]
        for idx, chunk in enumerate(chunks):
            docs.append(
                Document(
                    content=chunk,
                    meta={
                        "source": source,
                        "file_path": rel,
                        "tier": tier,
                        "content_type": ctype,
                        "chunk": idx,
                    },
                )
            )
        return docs

    def _code_fallback_documents(self) -> list[Document]:
        if not self.include_code_fallback:
            return []
        docs: list[Document] = []
        code_roots = [
            (REPOS_DIR / "smart_contracts" / "contracts" / "production", "smart_contracts", [".sol"]),
            (REPOS_DIR / "frontend" / "src" / "app" / "game", "frontend", [".tsx"]),
        ]
        repo_roots = {
            "smart_contracts": REPOS_DIR / "smart_contracts",
            "frontend": REPOS_DIR / "frontend",
        }
        for root, source, suffixes in code_roots:
            if not root.exists():
                continue
            repo_root = repo_roots[source]
            for path in root.rglob("*"):
                if not path.is_file() or path.suffix not in suffixes:
                    continue
                try:
                    text = path.read_text(encoding="utf-8", errors="replace")
                except Exception:
                    continue
                if len(text) > 8000:
                    text = text[:8000] + "\n...(truncated)"
                rel = str(path.relative_to(repo_root))
                docs.append(
                    Document(
                        content=text,
                        meta={
                            "source": source,
                            "file_path": rel,
                            "tier": TIER_SECONDARY,
                            "content_type": "code",
                        },
                    )
                )
        return docs

    def index_knowledge_base(self, force: bool = False) -> int:
        if self.is_ready and not force:
            return self.document_store.count_documents()

        if force:
            self.document_store = InMemoryDocumentStore()

        if not self.knowledge_dir.exists():
            raise FileNotFoundError(
                f"Knowledge base not found at {self.knowledge_dir}. "
                "Set KNOWLEDGE_BASE in backend/.env, run `python -m knowledge.generate.run_all`, "
                "then POST /api/reindex or restart the backend."
            )

        documents: list[Document] = []
        for path in sorted(self.knowledge_dir.rglob("*")):
            if not path.is_file():
                continue
            if path.suffix not in {".md", ".txt", ".json"}:
                continue
            documents.extend(self._documents_from_file(path, source="generated"))

        documents.extend(self._code_fallback_documents())
        if not documents:
            raise RuntimeError(f"No documents indexed from {self.knowledge_dir}")

        self.document_store.write_documents(documents)
        self.retriever = InMemoryBM25Retriever(document_store=self.document_store)
        self.load_facts()
        logger.info("Indexed %d knowledge documents", len(documents))
        return len(documents)

    def _retrieve_tier(self, query: str, tier: str, top_k: int) -> list[Document]:
        if not self.retriever:
            raise RuntimeError("Retriever not initialized")
        result = self.retriever.run(
            query=query,
            filters={"field": "tier", "operator": "==", "value": tier},
            top_k=top_k,
        )
        return result.get("documents", [])

    def retrieve(self, query: str, top_k_primary: int = 4, top_k_secondary: int = 2) -> list[Document]:
        if not self.is_ready:
            raise RuntimeError("Knowledge retriever is not ready")

        primary = self._retrieve_tier(query, TIER_PRIMARY, top_k_primary)
        secondary = self._retrieve_tier(query, TIER_SECONDARY, top_k_secondary)

        seen = set()
        merged: list[Document] = []
        for doc in primary + secondary:
            key = (doc.meta.get("file_path"), doc.meta.get("chunk", 0))
            if key in seen:
                continue
            seen.add(key)
            merged.append(doc)
        return merged

    def build_context(
        self,
        question: str,
        history: list[dict[str, str]] | None = None,
    ) -> tuple[str, list[str], list[Document]]:
        documents = self.retrieve(question)
        facts = self._facts_cache or self.load_facts()
        context_text, sources = build_context_pack(
            question=question,
            documents=documents,
            facts=facts,
            history=history,
        )
        return context_text, sources, documents
