"""Coordinates Haystack retrieval and Codex MCP synthesis."""
from __future__ import annotations

import logging
from typing import Any

from llm.codex_mcp_client import CodexMCPClient, CodexMCPError
from live.api_client import CosmicGameApiClient
from live.api_detector import needs_backend_api
from live.cginfo_client import CginfoClient
from live.detector import needs_live_state
from response_format import strip_sources_section, wants_detail
from retrieval.pipeline import KnowledgeRetriever
from sessions.conversation_logger import ConversationLogger
from sessions.store import ChatSession, SessionExpiredError, SessionStore

logger = logging.getLogger(__name__)


class HaystackUnavailableError(RuntimeError):
    pass


class Orchestrator:
    def __init__(
        self,
        retriever: KnowledgeRetriever,
        codex: CodexMCPClient,
        sessions: SessionStore,
        conversation_logger: ConversationLogger | None = None,
        cginfo: CginfoClient | None = None,
        api_client: CosmicGameApiClient | None = None,
    ):
        self.retriever = retriever
        self.codex = codex
        self.sessions = sessions
        self.conversation_logger = conversation_logger or ConversationLogger()
        self.cginfo = cginfo or CginfoClient()
        self.api_client = api_client or CosmicGameApiClient()

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
            "cginfo": self.cginfo.config_status(),
            "cosmic_game_api": self.api_client.config_status(),
            "sessions": self.sessions.stats(),
        }

    def _rewrite_question(self, question: str, session: ChatSession) -> str:
        """Expand follow-ups using recent history for better retrieval."""
        q_lower = question.lower()
        if len(session.messages) < 2:
            return question
        # Follow-ups about addresses should inherit deployment context from prior turn.
        if self._needs_address_context(q_lower, session):
            return (
                f"{question}\n"
                "Context: user previously asked about Cosmic Signature contracts deployed on Arbitrum. "
                "Return the concrete mainnet contract addresses from deployment records."
            )
        prior = session.messages[-4:-1]
        if not prior:
            return question
        context = " | ".join(f"{m['role']}: {m['content'][:120]}" for m in prior)
        return f"Conversation context: {context}\n\nCurrent question: {question}"

    @staticmethod
    def _needs_address_context(question_lower: str, session: ChatSession) -> bool:
        address_words = ("address", "addresses", "0x", "where are they deployed")
        if not any(w in question_lower for w in address_words):
            return False
        for msg in reversed(session.messages):
            if msg["role"] != "user":
                continue
            text = msg["content"].lower()
            if any(w in text for w in ("contract", "arbitrum", "deploy", "mainnet")):
                return True
        return False

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

        session = self.sessions.require(session_id)
        retrieval_query = self._rewrite_question(question, session)

        try:
            context_text, sources, _docs = self.retriever.build_context(
                question=question,
                history=session.messages,
                retrieval_query=retrieval_query,
            )
        except Exception as exc:
            self.conversation_logger.log_turn(session, question, f"ERROR [haystack]: {exc}")
            raise HaystackUnavailableError(f"Haystack retrieval failed: {exc}") from exc

        sources = list(sources)
        live_block = ""
        api_block = ""

        if needs_live_state(question):
            live_output, live_err = await self.cginfo.fetch_state()
            if live_output:
                live_block = (
                    "LIVE ON-CHAIN STATE (via cginfo — authoritative for current prices, "
                    "round status, timers, and bidder info)\n"
                    f"{live_output}\n\n"
                )
                sources.append("live:cginfo")
                logger.info("Injected live cginfo state (%d chars)", len(live_output))
            else:
                live_block = (
                    "LIVE ON-CHAIN STATE UNAVAILABLE\n"
                    f"{live_err}\n\n"
                )
                sources.append("live:cginfo-unavailable")
                logger.warning("Live state requested but cginfo failed: %s", live_err)

        if needs_backend_api(question):
            api_output, api_err, api_sources = await self.api_client.fetch_for_question(question)
            sources.extend(api_sources)
            if api_output:
                api_block = f"{api_output}\n\n"
                logger.info("Injected live API data (%d chars)", len(api_output))
            elif api_err:
                api_block = f"LIVE BACKEND API DATA UNAVAILABLE\n{api_err}\n\n"
                sources.append("live:api-unavailable")
                logger.warning("Backend API fetch failed: %s", api_err)

        prompt = (
            "Use the following CONTEXT to answer the user.\n\n"
            f"{live_block}"
            f"{api_block}"
            f"{context_text}\n\n"
        )
        if live_block.startswith("LIVE ON-CHAIN STATE (via cginfo"):
            prompt += (
                "The LIVE ON-CHAIN STATE section contains a fresh contract read. "
                "Use those exact values when answering about current bid prices, round "
                "status, countdowns, bidders, or prize amounts. Do not say you cannot "
                "know the current price if that data is present.\n\n"
            )
        if api_block.startswith("LIVE BACKEND API DATA (indexed"):
            prompt += (
                "The LIVE BACKEND API DATA section contains fresh indexed statistics from "
                "the Cosmic Game backend. Use NumERC20Donations, TotalDonatedNFTs, "
                "TotalRows, and related counts as authoritative when answering about "
                "donations, bid totals, or round statistics. Do not say the data is "
                "missing if those fields are present.\n\n"
            )
        if wants_detail(question):
            prompt += (
                "The user wants a detailed answer. You may use steps, background, and "
                "technical detail as needed.\n\n"
            )
        else:
            prompt += (
                "Keep the answer **brief** (1–3 sentences). Answer only what was asked. "
                "No Sources section, no /game/play steps, no fetch timestamps.\n\n"
            )
        prompt += "Write the answer now."

        try:
            answer, thread_id = await self.codex.ask(
                prompt=prompt,
                thread_id=session.codex_thread_id,
            )
        except CodexMCPError as exc:
            self.conversation_logger.log_turn(session, question, f"ERROR [codex_mcp]: {exc}")
            raise
        except Exception as exc:
            err = CodexMCPError(f"Codex MCP synthesis failed: {exc}")
            self.conversation_logger.log_turn(session, question, f"ERROR [codex_mcp]: {err}")
            raise err from exc

        session.codex_thread_id = thread_id
        answer = strip_sources_section(answer)
        session.add_message("user", question)
        session.add_message("assistant", answer)
        self.conversation_logger.log_turn(session, question, answer)

        return {
            "answer": answer,
            "sources": sources,
            "session_id": session.session_id,
        }
