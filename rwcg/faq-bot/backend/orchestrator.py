"""Coordinates Haystack retrieval and Codex MCP synthesis."""
from __future__ import annotations

import logging
from typing import Any

from llm.codex_mcp_client import CodexMCPClient, CodexMCPError
from retrieval.pipeline import KnowledgeRetriever
from sessions.store import ChatSession, SessionStore

logger = logging.getLogger(__name__)


class HaystackUnavailableError(RuntimeError):
    pass


class Orchestrator:
    def __init__(
        self,
        retriever: KnowledgeRetriever,
        codex: CodexMCPClient,
        sessions: SessionStore,
    ):
        self.retriever = retriever
        self.codex = codex
        self.sessions = sessions

    async def health(self) -> dict[str, Any]:
        haystack = {
            "ready": self.retriever.is_ready,
            "documents": self.retriever.document_store.count_documents() if self.retriever.is_ready else 0,
            "knowledge_base": str(self.retriever.knowledge_dir),
        }
        try:
            codex = await self.codex.health_check()
        except Exception as exc:
            codex = {"ready": False, "error": str(exc)}
        ok = haystack["ready"] and codex.get("ready", False)
        return {
            "status": "healthy" if ok else "degraded",
            "haystack": haystack,
            "codex_mcp": codex,
            "sessions": self.sessions.stats(),
        }

    def _rewrite_question(self, question: str, session: ChatSession) -> str:
        """Expand follow-ups using recent history for better retrieval."""
        if len(session.messages) < 2:
            return question
        prior = session.messages[-4:-1]
        if not prior:
            return question
        context = " | ".join(f"{m['role']}: {m['content'][:120]}" for m in prior)
        return f"Conversation context: {context}\n\nCurrent question: {question}"

    async def query(self, question: str, session_id: str | None = None) -> dict[str, Any]:
        question = question.strip()
        if not question:
            raise ValueError("Question cannot be empty")

        if not self.retriever.is_ready:
            raise HaystackUnavailableError(
                f"Haystack knowledge base is not indexed (KNOWLEDGE_BASE={self.retriever.knowledge_dir}). "
                "Run `python -m knowledge.generate.run_all` then POST /api/reindex."
            )
        if not self.codex.is_ready:
            raise CodexMCPError(
                "Codex MCP is not connected. Start the backend with access to `codex mcp-server` "
                "or ensure CODEX_MCP_COMMAND is configured."
            )

        session = self.sessions.get_or_create(session_id)
        retrieval_query = self._rewrite_question(question, session)

        try:
            context_text, sources, _docs = self.retriever.build_context(
                question=retrieval_query,
                history=session.messages,
            )
        except Exception as exc:
            raise HaystackUnavailableError(f"Haystack retrieval failed: {exc}") from exc

        prompt = (
            "Use the following CONTEXT to answer the user.\n\n"
            f"{context_text}\n\n"
            "Write a helpful answer. Include a Sources section with paths listed above."
        )

        try:
            answer, thread_id = await self.codex.ask(
                prompt=prompt,
                thread_id=session.codex_thread_id,
            )
        except CodexMCPError:
            raise
        except Exception as exc:
            raise CodexMCPError(f"Codex MCP synthesis failed: {exc}") from exc

        session.codex_thread_id = thread_id
        session.add_message("user", question)
        session.add_message("assistant", answer)

        return {
            "answer": answer,
            "sources": sources,
            "session_id": session.session_id,
        }
