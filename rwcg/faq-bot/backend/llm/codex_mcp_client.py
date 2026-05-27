"""Codex MCP client for FAQ answer synthesis."""
from __future__ import annotations

import json
import logging
import os
import shlex
from contextlib import AsyncExitStack
from typing import Any

from mcp import ClientSession, StdioServerParameters
from mcp.client.stdio import stdio_client

logger = logging.getLogger(__name__)

# `codex mcp-server` must expose only these MCP tools (text synthesis via codex / codex-reply).
EXPECTED_MCP_TOOLS = frozenset({"codex", "codex-reply"})


class _SuppressCodexNotificationWarnings(logging.Filter):
    """Codex MCP emits custom `codex/event` notifications the Python MCP SDK cannot parse."""

    def filter(self, record: logging.LogRecord) -> bool:
        return "Failed to validate notification" not in record.getMessage()


logging.getLogger().addFilter(_SuppressCodexNotificationWarnings())

SYSTEM_INSTRUCTIONS = """You are the Cosmic Signature FAQ assistant — a text synthesis model only.

The Python backend has already assembled the full CONTEXT in your prompt:
conversation history, Haystack retrieved documents, structured facts, live API blocks,
and on-chain cginfo blocks. Treat that prompt text as the only source of truth.

You must NOT run shell commands, read or write files, search the filesystem, use git,
curl, docker, or any other tools. Do not attempt to gather more data. Return only the
final answer text.

Answer using ONLY the provided CONTEXT.
If the context is insufficient, say exactly what information is missing — do not invent addresses, counts, or behavior.

Response length (important):
- Default to **short answers**: 1–3 sentences with only the facts the user asked for.
- Do **not** add UI walkthrough steps, background, or extra sections unless the user asked for them.
- Do **not** include a "Sources" section — sources are shown separately in the UI.
- Do **not** mention fetch timestamps or "from the live on-chain state you provided" unless the user asked about freshness.
- Give detailed step-by-step or developer-level answers only when the user clearly wants depth (explain, how to, steps, tutorial, tell me more).

Audience:
- Beginner how-to questions: plain language, steps only when requested.
- Expert/developer questions: precise terms and signatures when relevant, still concise unless detail was requested.

Time-range bid counts:
- When PERIOD_BID_COUNT is in LIVE BACKEND API DATA, state that number and the period label.
- Never replace a period count with the current-round TotalBids.
- If the user asked for a catch-up / status update, give the period bid count (if available) plus a brief current-round snapshot.
- If the question needed a date clarification, the assistant message may already be a clarifying question — do not invent dates or counts.

Chrono Warrior / Endurance Champion:
- When LIVE CHAMPIONS STATE is present, answer who (address) and when (ChronoSegmentStart UTC + live status) from those fields.
- Do not claim the data is unavailable if ChronoWarriorAddress or ChronoSegmentStart is in the context."""


class CodexMCPError(RuntimeError):
    pass


class CodexMCPClient:
    """Long-lived stdio connection to `codex mcp-server` (synthesis only)."""

    def __init__(self) -> None:
        command = os.getenv("CODEX_MCP_COMMAND", "codex")
        args_raw = os.getenv("CODEX_MCP_ARGS", "mcp-server")
        self._command = command
        self._args = shlex.split(args_raw)
        self._cwd = os.getenv("CODEX_MCP_CWD", os.getcwd())
        self._sandbox = os.getenv("CODEX_MCP_SANDBOX", "read-only")
        self._approval_policy = os.getenv("CODEX_MCP_APPROVAL_POLICY", "never")
        self._model = os.getenv("CODEX_MCP_MODEL", "")
        self._stack: AsyncExitStack | None = None
        self._session: ClientSession | None = None
        self._ready = False

    @staticmethod
    def _validate_tool_surface(tool_names: set[str]) -> None:
        if tool_names == EXPECTED_MCP_TOOLS:
            return
        unexpected = sorted(tool_names - EXPECTED_MCP_TOOLS)
        missing = sorted(EXPECTED_MCP_TOOLS - tool_names)
        raise CodexMCPError(
            "Codex MCP tool surface is not synthesis-only. "
            f"Expected exactly {sorted(EXPECTED_MCP_TOOLS)}; got {sorted(tool_names)}. "
            f"Unexpected: {unexpected or 'none'}; missing: {missing or 'none'}. "
            "Remove extra MCP servers from the Codex config and use codex-nsjail mcp-server."
        )

    def _build_server_env(self) -> dict[str, str]:
        """Subprocess env for the jailed Codex account. Visitor text never goes here."""
        env = os.environ.copy()
        home = os.getenv("CODEX_MCP_HOME", "").strip()
        codex_home = os.getenv("CODEX_MCP_CODEX_HOME", "").strip()
        if home:
            env["HOME"] = home
        if codex_home:
            env["CODEX_HOME"] = codex_home
        return env

    @property
    def is_ready(self) -> bool:
        return self._ready and self._session is not None

    async def start(self) -> None:
        if self._ready:
            return
        self._stack = AsyncExitStack()
        server_params = StdioServerParameters(
            command=self._command,
            args=self._args,
            cwd=self._cwd,
            env=self._build_server_env(),
        )
        read, write = await self._stack.enter_async_context(stdio_client(server_params))
        self._session = await self._stack.enter_async_context(ClientSession(read, write))
        await self._session.initialize()
        await self._assert_expected_tools()
        self._ready = True
        logger.info(
            "Codex MCP synthesis server ready (%s %s, cwd=%s, tools=%s)",
            self._command,
            " ".join(self._args),
            self._cwd,
            sorted(EXPECTED_MCP_TOOLS),
        )

    async def _assert_expected_tools(self) -> None:
        if self._session is None:
            raise CodexMCPError("Codex MCP session not initialized")
        tools = await self._session.list_tools()
        names = {t.name for t in tools.tools}
        self._validate_tool_surface(names)

    async def stop(self) -> None:
        if self._stack:
            await self._stack.aclose()
        self._stack = None
        self._session = None
        self._ready = False

    async def health_check(self) -> dict[str, Any]:
        if not self.is_ready:
            return {"ready": False, "error": "Codex MCP client not started"}
        try:
            await self._assert_expected_tools()
        except CodexMCPError as exc:
            return {"ready": False, "error": str(exc)}
        return {
            "ready": True,
            "tools": sorted(EXPECTED_MCP_TOOLS),
            "command": self._command,
            "args": self._args,
            "cwd": self._cwd,
            "sandbox": self._sandbox,
            "approval_policy": self._approval_policy,
        }

    def _extract_text(self, result: Any) -> tuple[str, str | None]:
        thread_id = None
        if hasattr(result, "structuredContent") and result.structuredContent:
            sc = result.structuredContent
            if isinstance(sc, dict):
                thread_id = sc.get("threadId") or sc.get("conversationId")
                if sc.get("content"):
                    return str(sc["content"]), thread_id
        # Fallback: parse content blocks
        texts = []
        for block in getattr(result, "content", []) or []:
            if getattr(block, "type", None) == "text":
                texts.append(block.text)
            elif isinstance(block, dict) and block.get("type") == "text":
                texts.append(block.get("text", ""))
        content = "\n".join(t for t in texts if t).strip()
        return content, thread_id

    async def ask(
        self,
        prompt: str,
        thread_id: str | None = None,
    ) -> tuple[str, str | None]:
        if not self.is_ready or self._session is None:
            raise CodexMCPError("Codex MCP is not connected. Ensure `codex mcp-server` can be started.")

        # Visitor text is passed only via MCP tool JSON (prompt field), never argv/env/cwd.
        tool_name = "codex-reply" if thread_id else "codex"
        arguments: dict[str, Any] = {
            "prompt": prompt,
            "sandbox": self._sandbox,
            "approval-policy": self._approval_policy,
            "base-instructions": SYSTEM_INSTRUCTIONS,
            "cwd": self._cwd,
        }
        if self._model:
            arguments["model"] = self._model
        if thread_id:
            arguments["threadId"] = thread_id

        try:
            result = await self._session.call_tool(tool_name, arguments=arguments)
        except Exception as exc:
            raise CodexMCPError(f"Codex MCP tool call failed: {exc}") from exc

        if getattr(result, "isError", False):
            raise CodexMCPError(f"Codex MCP returned error: {result}")

        content, new_thread_id = self._extract_text(result)
        if not content:
            # Some versions embed JSON in text content
            raw = json.dumps(result.model_dump() if hasattr(result, "model_dump") else str(result))
            raise CodexMCPError(f"Codex MCP returned empty content: {raw[:500]}")
        return content, new_thread_id or thread_id
