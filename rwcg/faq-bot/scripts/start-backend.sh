#!/usr/bin/env bash
# Start the FAQ bot Python backend (Haystack + Codex MCP).
#
# Bind address and port come from ~/configs/faq-bot.env (HOST, PORT) and can be
# overridden per invocation: ./scripts/start-backend.sh --host 127.0.0.1 --port 8001
set -euo pipefail

usage() {
  cat <<EOF
Usage: $(basename "$0") [options]

Options:
  --host HOST   Listen address (default: 0.0.0.0, or HOST from faq-bot.env)
  --port PORT   Listen port (default: 8000, or PORT from faq-bot.env)
  -h, --help    Show this help

Config file: \${FAQ_BOT_ENV:-\$HOME/configs/faq-bot.env}
EOF
}

CONFIG="${FAQ_BOT_ENV:-${HOME}/configs/faq-bot.env}"
if [[ -f "${CONFIG}" ]]; then
  # shellcheck disable=SC1090
  source "${CONFIG}"
fi

CLI_HOST=""
CLI_PORT=""
while [[ $# -gt 0 ]]; do
  case "$1" in
    --host)
      CLI_HOST="${2:-}"
      shift 2
      ;;
    --port)
      CLI_PORT="${2:-}"
      shift 2
      ;;
    -h | --help)
      usage
      exit 0
      ;;
    *)
      echo "ERROR: Unknown option: $1" >&2
      usage >&2
      exit 1
      ;;
  esac
done

if [[ -n "${CLI_HOST}" ]]; then
  HOST="${CLI_HOST}"
fi
if [[ -n "${CLI_PORT}" ]]; then
  PORT="${CLI_PORT}"
fi

HOST="${HOST:-0.0.0.0}"
PORT="${PORT:-8000}"
export HOST PORT

if [[ -z "${FAQ_BOT_VENV:-}" ]]; then
  echo "ERROR: FAQ_BOT_VENV is not set." >&2
  echo "Add to ${CONFIG}:" >&2
  echo "  export FAQ_BOT_VENV=/path/to/venv" >&2
  echo "Then run: scripts/setup-venv.sh" >&2
  exit 1
fi

if [[ ! -f "${FAQ_BOT_VENV}/bin/activate" ]]; then
  echo "ERROR: venv not found at FAQ_BOT_VENV=${FAQ_BOT_VENV}" >&2
  echo "Run: scripts/setup-venv.sh" >&2
  exit 1
fi

# shellcheck disable=SC1091
source "${FAQ_BOT_VENV}/bin/activate"

cd "$(dirname "$0")/../backend"

# Config comes from ~/configs/faq-bot.env (sourced above). Do not require backend/.env.
if [[ -f .env ]]; then
  echo "Note: backend/.env exists but faq-bot.env exports take precedence (dotenv does not override)."
fi

# Codex nsjail: Python spawns the wrapper from a cwd cgprod can enter (not /home/aijail).
if [[ -f llm/codex_mcp_client.py ]] && ! grep -q '_spawn_cwd' llm/codex_mcp_client.py; then
  echo "ERROR: llm/codex_mcp_client.py is outdated (missing CODEX_MCP_SPAWN_CWD support)." >&2
  echo "Run: git pull   (in rwcg/faq-bot), then restart." >&2
  exit 1
fi
CODEX_MCP_SPAWN_CWD="${CODEX_MCP_SPAWN_CWD:-$(pwd)}"
if [[ ! -r "${CODEX_MCP_SPAWN_CWD}" || ! -x "${CODEX_MCP_SPAWN_CWD}" ]]; then
  echo "ERROR: CODEX_MCP_SPAWN_CWD is not accessible: ${CODEX_MCP_SPAWN_CWD}" >&2
  echo "Set in ${CONFIG}, e.g. export CODEX_MCP_SPAWN_CWD=$(pwd)" >&2
  exit 1
fi
export CODEX_MCP_SPAWN_CWD

display_host="${HOST}"
if [[ "${display_host}" == "0.0.0.0" || "${display_host}" == "::" ]]; then
  display_host="127.0.0.1"
fi

echo "Config:  ${CONFIG}"
echo "Venv:    ${FAQ_BOT_VENV}"
echo "KB:      ${KNOWLEDGE_BASE:-"(set KNOWLEDGE_BASE in faq-bot.env)"}"
echo "Codex:   ${CODEX_MCP_COMMAND:-codex} (spawn cwd ${CODEX_MCP_SPAWN_CWD})"
echo "Listen:  ${HOST}:${PORT}"
echo "API:     http://${display_host}:${PORT}"
echo "Test UI: http://${display_host}:${PORT}/test-ui"
exec python app.py
