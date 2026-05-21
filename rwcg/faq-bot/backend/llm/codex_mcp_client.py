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

SYSTEM_INSTRUCTIONS = """You are the Cosmic Signature FAQ assistant.

Answer using ONLY the provided CONTEXT (facts + retrieved documents + conversation history).
If the context is insufficient, say exactly what information is missing — do not invent addresses, counts, or behavior.

Audience rules:
- For beginner questions: plain language, step-by-step, mention pages like /game/play when relevant.
- For expert/developer questions: precise terminology, cite source paths, include counts and contract names when present in facts.

Always end with a short "Sources:" list referencing paths from the context."""


class CodexMCPError(RuntimeError):
    pass


class CodexMCPClient:
    """Long-lived stdio connection to `codex mcp-server`."""

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
        )
        read, write = await self._stack.enter_async_context(stdio_client(server_params))
        self._session = await self._stack.enter_async_context(ClientSession(read, write))
        await self._session.initialize()
        self._ready = True
        logger.info("Codex MCP client connected (%s %s)", self._command, " ".join(self._args))

    async def stop(self) -> None:
        if self._stack:
            await self._stack.aclose()
        self._stack = None
        self._session = None
        self._ready = False

    async def health_check(self) -> dict[str, Any]:
        if not self.is_ready:
            return {"ready": False, "error": "Codex MCP client not started"}
        assert self._session is not None
        tools = await self._session.list_tools()
        names = [t.name for t in tools.tools]
        return {"ready": True, "tools": names}

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
