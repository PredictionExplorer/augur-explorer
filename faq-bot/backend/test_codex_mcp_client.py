"""Tests for Codex MCP synthesis-only tool surface validation."""
from __future__ import annotations

from llm.codex_mcp_client import CodexMCPClient, CodexMCPError, EXPECTED_MCP_TOOLS


def test_validate_tool_surface_accepts_synthesis_tools_only():
    CodexMCPClient._validate_tool_surface(set(EXPECTED_MCP_TOOLS))


def test_validate_tool_surface_rejects_extra_tools():
    try:
        CodexMCPClient._validate_tool_surface({"codex", "codex-reply", "read_file"})
    except CodexMCPError:
        return
    raise AssertionError("expected CodexMCPError for extra tools")


def test_validate_tool_surface_rejects_missing_tools():
    try:
        CodexMCPClient._validate_tool_surface({"codex"})
    except CodexMCPError:
        return
    raise AssertionError("expected CodexMCPError for missing tools")
