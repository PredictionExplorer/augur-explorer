"""Coordinates Haystack retrieval and Codex MCP synthesis."""
from __future__ import annotations

import json
import logging
import time
from typing import Any

from knowledge.config import KNOWLEDGE_BASE
from llm.codex_mcp_client import CodexMCPClient, CodexMCPError
from live.api_client import CosmicGameApiClient
from live.api_detector import (
    needs_activation_window_config,
    needs_backend_api,
    needs_time_range_bids,
)
from live.cginfo_client import CginfoClient
from live.detector import (
    needs_backend_url_info,
    needs_champions_state,
    needs_live_state,
    needs_round_end_time,
    needs_staking_stats,
    normalize_question,
)
from response_format import strip_sources_section, wants_detail
from retrieval.pipeline import KnowledgeRetriever
from sessions.conversation_logger import ConversationLogger
from sessions.query_debug_logger import QueryDebugLogger
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
        query_debug_logger: QueryDebugLogger | None = None,
        cginfo: CginfoClient | None = None,
        api_client: CosmicGameApiClient | None = None,
    ):
        self.retriever = retriever
        self.codex = codex
        self.sessions = sessions
        self.conversation_logger = conversation_logger or ConversationLogger()
        self.query_debug_logger = query_debug_logger or QueryDebugLogger()
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

    def _build_runtime_config_block(self, question: str) -> tuple[str, list[str]]:
        """Inject configured API/RPC URLs for meta questions about backend endpoints."""
        if not needs_backend_url_info(question):
            return "", []

        api_cfg = self.api_client.config_status()
        cg_cfg = self.cginfo.config_status()
        lines = [
            "BOT RUNTIME CONFIG (authoritative for backend/API URL questions)",
            f"Cosmic Game REST API base (FAQ_BOT_API_URL): {api_cfg.get('base_url') or '(not configured)'}",
            f"Cosmic Game API configured: {api_cfg.get('configured')}",
        ]

        facts = self.retriever._facts_cache or self.retriever.load_facts()
        net_env = (facts.get("network-environment") or {}).get("networks") or {}
        for network, info in sorted(net_env.items()):
            frontend_env = info.get("frontend_env") or {}
            api_url = frontend_env.get("NEXT_PUBLIC_API_URL")
            rpc_url = frontend_env.get("NEXT_PUBLIC_RPC_URL")
            if api_url:
                label = info.get("network_label") or network
                lines.append(f"Production frontend API URL ({label}, NEXT_PUBLIC_API_URL): {api_url}")
            if rpc_url:
                label = info.get("network_label") or network
                lines.append(f"Production frontend RPC URL ({label}, NEXT_PUBLIC_RPC_URL): {rpc_url}")

        if cg_cfg.get("rpc_url_set"):
            lines.append("On-chain reads RPC (RPC_URL / FAQ_BOT_RPC_URL): configured")
        else:
            lines.append("On-chain reads RPC (RPC_URL / FAQ_BOT_RPC_URL): not configured")
        if cg_cfg.get("game_address"):
            lines.append(f"CosmicSignatureGame proxy: {cg_cfg['game_address']}")

        lines.append(
            "The REST API base path is /api/cosmicgame/ (e.g. statistics/dashboard, rounds/info, bid/list)."
        )
        return "\n".join(lines) + "\n\n", ["runtime:config"]

    @staticmethod
    def _classifier_snapshot(question: str) -> dict[str, bool]:
        return {
            "needs_live_state": needs_live_state(question),
            "needs_backend_api": needs_backend_api(question),
            "needs_backend_url_info": needs_backend_url_info(question),
            "needs_round_end_time": needs_round_end_time(question),
            "needs_champions_state": needs_champions_state(question),
            "needs_staking_stats": needs_staking_stats(question),
            "needs_time_range_bids": needs_time_range_bids(question),
            "needs_activation_window_config": needs_activation_window_config(question),
            "wants_detail": wants_detail(question),
            "normalized_question": normalize_question(question),
        }

    @staticmethod
    def _kb_address_diagnosis() -> dict[str, Any]:
        addr_path = KNOWLEDGE_BASE / "facts" / "deployed-addresses.json"
        diagnosis: dict[str, Any] = {
            "knowledge_base": str(KNOWLEDGE_BASE),
            "deployed_addresses_path": str(addr_path),
            "deployed_addresses_exists": addr_path.is_file(),
        }
        if not addr_path.is_file():
            diagnosis["error"] = "deployed-addresses.json missing from knowledge base"
            return diagnosis
        try:
            data = json.loads(addr_path.read_text(encoding="utf-8"))
            networks = data.get("networks") or {}
            diagnosis["network_keys"] = sorted(networks.keys())
            mainnet = networks.get("arbitrum_one_mainnet") or {}
            addresses = mainnet.get("addresses") or {}
            diagnosis["mainnet_address_keys"] = sorted(addresses.keys())
            diagnosis["CosmicSignatureGame proxy"] = addresses.get("CosmicSignatureGame proxy", "")
            diagnosis["CosmicSignatureGame"] = addresses.get("CosmicSignatureGame", "")
            if not addresses:
                diagnosis["error"] = (
                    "arbitrum_one_mainnet.addresses is empty — re-run knowledge generation "
                    "or set COSMIC_GAME_ADDR in faq-bot.env"
                )
            elif not (
                addresses.get("CosmicSignatureGame proxy") or addresses.get("CosmicSignatureGame")
            ):
                diagnosis["error"] = (
                    "No CosmicSignatureGame proxy address in KB — set COSMIC_GAME_ADDR "
                    "or add deployment facts"
                )
        except Exception as exc:
            diagnosis["error"] = f"Failed to parse deployed-addresses.json: {exc}"
        return diagnosis

    def _emit_debug(self, session: ChatSession, debug: dict[str, Any]) -> None:
        self.query_debug_logger.log_turn(session, debug)

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

        t_start = time.monotonic()
        session = self.sessions.require(session_id)
        turn = len(session.messages) // 2 + 1
        classifiers = self._classifier_snapshot(question)
        retrieval_query = self._rewrite_question(question, session)
        thread_id_before = session.codex_thread_id
        prompt_hints: list[str] = []

        debug: dict[str, Any] = {
            "turn": turn,
            "input": {
                "session_id": session.session_id,
                "question": question,
                "message_count_before": len(session.messages),
                "retrieval_query": retrieval_query,
                "codex_thread_id_before": thread_id_before,
            },
            "classifiers": classifiers,
            "timing": {},
        }

        t_retrieval = time.monotonic()
        try:
            context_text, sources, docs = self.retriever.build_context(
                question=question,
                history=session.messages,
                retrieval_query=retrieval_query,
            )
        except Exception as exc:
            debug["retrieval"] = {"error": str(exc)}
            debug["output"] = {"error": f"haystack: {exc}"}
            debug["timing"]["total_ms"] = round((time.monotonic() - t_start) * 1000, 1)
            self._emit_debug(session, debug)
            self.conversation_logger.log_turn(session, question, f"ERROR [haystack]: {exc}")
            raise HaystackUnavailableError(f"Haystack retrieval failed: {exc}") from exc
        debug["timing"]["haystack_ms"] = round((time.monotonic() - t_retrieval) * 1000, 1)

        doc_meta = [
            {
                "source": getattr(d, "meta", {}).get("source"),
                "file_path": getattr(d, "meta", {}).get("file_path"),
                "tier": getattr(d, "meta", {}).get("tier"),
            }
            for d in (docs or [])
        ]
        debug["retrieval"] = {
            "sources": list(sources),
            "document_count": len(docs or []),
            "documents": doc_meta,
            "context_chars": len(context_text),
            "context_preview": QueryDebugLogger.preview(context_text),
        }

        sources = list(sources)
        live_block = ""
        api_block = ""
        config_block, config_sources = self._build_runtime_config_block(question)
        sources.extend(config_sources)

        cginfo_cfg = self.cginfo.config_status()
        cginfo_debug: dict[str, Any] = {
            "fetch_attempted": classifiers["needs_live_state"],
            "config_status": cginfo_cfg,
            "kb_address_diagnosis": self._kb_address_diagnosis(),
        }

        t_cginfo = time.monotonic()
        if classifiers["needs_live_state"]:
            live_output, live_err = await self.cginfo.fetch_state()
            cginfo_debug["fetch_success"] = bool(live_output)
            cginfo_debug["fetch_error"] = live_err
            cginfo_debug["live_output_chars"] = len(live_output or "")
            cginfo_debug["live_output_preview"] = QueryDebugLogger.preview(live_output or "")
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
        else:
            cginfo_debug["skipped_reason"] = "needs_live_state=false"
        debug["timing"]["cginfo_ms"] = round((time.monotonic() - t_cginfo) * 1000, 1)
        debug["cginfo"] = cginfo_debug

        api_cfg = self.api_client.config_status()
        api_debug: dict[str, Any] = {
            "fetch_attempted": classifiers["needs_backend_api"],
            "config_status": api_cfg,
        }

        t_api = time.monotonic()
        api_result = None
        if classifiers["needs_backend_api"]:
            api_result = await self.api_client.fetch_for_question(question)
            api_debug["fetch_sources"] = api_result.sources
            api_debug["clarification_answer"] = api_result.clarification_answer
            api_debug["block_chars"] = len(api_result.block or "")
            api_debug["block_preview"] = QueryDebugLogger.preview(api_result.block or "")
            api_debug["error"] = api_result.error
            sources.extend(api_result.sources)
            if api_result.clarification_answer:
                answer = api_result.clarification_answer
                session.add_message("user", question)
                session.add_message("assistant", answer)
                debug["api"] = api_debug
                debug["output"] = {
                    "early_return": "api_clarification",
                    "sources": sources,
                    "answer": answer,
                }
                debug["timing"]["api_ms"] = round((time.monotonic() - t_api) * 1000, 1)
                debug["timing"]["total_ms"] = round((time.monotonic() - t_start) * 1000, 1)
                self._emit_debug(session, debug)
                self.conversation_logger.log_turn(session, question, answer)
                return {
                    "answer": answer,
                    "sources": sources,
                    "session_id": session.session_id,
                }
            if api_result.block:
                api_block = f"{api_result.block}\n\n"
                logger.info("Injected live API data (%d chars)", len(api_result.block))
            elif api_result.error:
                api_block = f"LIVE BACKEND API DATA UNAVAILABLE\n{api_result.error}\n\n"
                sources.append("live:api-unavailable")
                logger.warning("Backend API fetch failed: %s", api_result.error)
        else:
            api_debug["skipped_reason"] = "needs_backend_api=false"
        debug["timing"]["api_ms"] = round((time.monotonic() - t_api) * 1000, 1)
        debug["api"] = api_debug

        prompt = (
            "Use the following CONTEXT to answer the user.\n\n"
            f"{config_block}"
            f"{live_block}"
            f"{api_block}"
            f"{context_text}\n\n"
        )
        if config_block:
            prompt_hints.append("runtime_config")
            prompt += (
                "The BOT RUNTIME CONFIG section lists the Cosmic Game REST API base URL this bot uses. "
                "Answer backend/API URL questions with FAQ_BOT_API_URL and/or NEXT_PUBLIC_API_URL from "
                "that section. Do not say the URL is missing when it is present.\n\n"
            )
        if live_block.startswith("LIVE ON-CHAIN STATE (via cginfo"):
            prompt_hints.append("live_on_chain_authoritative")
            prompt += (
                "The LIVE ON-CHAIN STATE section contains a fresh contract read. "
                "Use those exact values when answering about current bid prices, round "
                "status, countdowns, bidders, or prize amounts. Do not say you cannot "
                "know the current price if that data is present.\n\n"
            )
        elif live_block.startswith("LIVE ON-CHAIN STATE UNAVAILABLE"):
            prompt_hints.append("live_on_chain_unavailable")
        if api_block.startswith("LIVE BACKEND API DATA (indexed"):
            prompt_hints.append("live_api_authoritative")
            prompt += (
                "The LIVE BACKEND API DATA section contains fresh indexed statistics from "
                "the Cosmic Game backend. Use NumERC20Donations, TotalDonatedNFTs, "
                "TotalRows, PERIOD_BID_COUNT, and CURRENT_ROUND_STATUS fields as authoritative. "
                "PERIOD_BID_COUNT is for calendar/time-window bid questions; round TotalBids is "
                "the entire current round — do not confuse them. "
                "For catch-up questions, answer the time-window count first (if present), then "
                "summarize current round status in 1–2 sentences. "
                "Do not say data is missing if those fields are present.\n\n"
            )
        elif api_block.startswith("LIVE BACKEND API DATA UNAVAILABLE"):
            prompt_hints.append("live_api_unavailable")
        if api_block.startswith("LIVE ACTIVATION WINDOW CONFIG"):
            prompt_hints.append("live_activation_window_config")
            prompt += (
                "The LIVE ACTIVATION WINDOW CONFIG section lists indexed admin configuration "
                "changes during the round's param/activation window (between prize claim / "
                "param window start and round activation). Summarize what changed using those "
                "event names and values. Do not say config history is missing when that list "
                "is present.\n\n"
            )
        if "LIVE CHAMPIONS STATE" in api_block:
            prompt_hints.append("live_champions")
            prompt += (
                "The LIVE CHAMPIONS STATE section has the current Chrono Warrior and "
                "Endurance Champion addresses and timestamps. Use ChronoWarriorAddress and "
                "ChronoSegmentStart (UTC) for who/when questions. Do not say the data is missing "
                "when those fields are present.\n\n"
            )
        if "LIVE STAKING STATE" in api_block:
            prompt_hints.append("live_staking")
            prompt += (
                "The LIVE STAKING STATE section has current and historical staker counts. "
                "For 'how many stakers currently', use NumActiveStakersCST and "
                "NumActiveStakersRandomWalk. Do not say staker data is missing when those "
                "fields are present.\n\n"
            )
        if "LIVE ROUND END STATE" in api_block or "ROUND_END_UTC" in api_block:
            prompt_hints.append("live_round_end")
            prompt += (
                "The LIVE ROUND END STATE section has the projected cycle end time. "
                "Answer with ROUND_END_UTC (and TIME_UNTIL_ROUND_END if helpful). "
                "If ROUND_STATUS is pre-activation, use ROUND_OPENS_UTC instead. "
                "If waiting for the first gesture, say the countdown has not started yet. "
                "Mention that new gestures extend the timer when relevant. "
                "Do not say the end time is unknown when those fields are present.\n\n"
            )
        if live_block.startswith("LIVE ON-CHAIN STATE (via cginfo") and "MainPrizeTime" in live_block:
            prompt_hints.append("main_prize_time_hint")
            prompt += (
                "For round/cycle end time, prefer MainPrizeTime (timestamp) and "
                "'Duration until prize' from LIVE ON-CHAIN STATE when LIVE ROUND END STATE "
                "is not present.\n\n"
            )
        if classifiers["wants_detail"]:
            prompt_hints.append("wants_detail")
            prompt += (
                "The user wants a detailed answer. You may use steps, background, and "
                "technical detail as needed.\n\n"
            )
        else:
            prompt_hints.append("brief_answer")
            prompt += (
                "Keep the answer **brief** (1–3 sentences). Answer only what was asked. "
                "No Sources section, no /game/play steps, no fetch timestamps.\n\n"
            )
        prompt += "Write the answer now."

        debug["prompt"] = {
            "hints": prompt_hints,
            "config_block_chars": len(config_block),
            "config_block_preview": QueryDebugLogger.preview(config_block),
            "live_block_chars": len(live_block),
            "live_block_preview": QueryDebugLogger.preview(live_block),
            "api_block_chars": len(api_block),
            "api_block_preview": QueryDebugLogger.preview(api_block),
            "context_text_chars": len(context_text),
            "full_prompt_chars": len(prompt),
            "full_prompt": prompt,
        }

        t_codex = time.monotonic()
        try:
            answer, thread_id = await self.codex.ask(
                prompt=prompt,
                thread_id=session.codex_thread_id,
            )
        except CodexMCPError as exc:
            debug["codex"] = {
                "success": False,
                "error": str(exc),
                "thread_id_before": thread_id_before,
            }
            debug["output"] = {"error": f"codex_mcp: {exc}", "sources": sources}
            debug["timing"]["codex_ms"] = round((time.monotonic() - t_codex) * 1000, 1)
            debug["timing"]["total_ms"] = round((time.monotonic() - t_start) * 1000, 1)
            self._emit_debug(session, debug)
            self.conversation_logger.log_turn(session, question, f"ERROR [codex_mcp]: {exc}")
            raise
        except Exception as exc:
            err = CodexMCPError(f"Codex MCP synthesis failed: {exc}")
            debug["codex"] = {
                "success": False,
                "error": str(err),
                "thread_id_before": thread_id_before,
            }
            debug["output"] = {"error": str(err), "sources": sources}
            debug["timing"]["codex_ms"] = round((time.monotonic() - t_codex) * 1000, 1)
            debug["timing"]["total_ms"] = round((time.monotonic() - t_start) * 1000, 1)
            self._emit_debug(session, debug)
            self.conversation_logger.log_turn(session, question, f"ERROR [codex_mcp]: {err}")
            raise err from exc
        debug["timing"]["codex_ms"] = round((time.monotonic() - t_codex) * 1000, 1)

        session.codex_thread_id = thread_id
        answer = strip_sources_section(answer)
        session.add_message("user", question)
        session.add_message("assistant", answer)

        debug["codex"] = {
            "success": True,
            "thread_id_before": thread_id_before,
            "thread_id_after": thread_id,
        }
        debug["output"] = {
            "sources": sources,
            "answer_chars": len(answer),
            "answer": answer,
        }
        if classifiers["needs_live_state"] and "live:cginfo-unavailable" in sources:
            debug["notes"] = {
                "live_data_failure": (
                    "Question required on-chain live data but cginfo fetch failed or is "
                    "unconfigured. Codex only sees LIVE ON-CHAIN STATE UNAVAILABLE — expect "
                    "answers that say balance/price is missing."
                ),
                "likely_fix": (
                    "Set COSMIC_GAME_ADDR in faq-bot.env and ensure deployed-addresses.json "
                    "exists in KNOWLEDGE_BASE, or re-run knowledge generation."
                ),
            }
        debug["timing"]["total_ms"] = round((time.monotonic() - t_start) * 1000, 1)
        self._emit_debug(session, debug)
        self.conversation_logger.log_turn(session, question, answer)

        return {
            "answer": answer,
            "sources": sources,
            "session_id": session.session_id,
        }
