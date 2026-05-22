#!/usr/bin/env bash
# Start the FAQ bot Python backend (Haystack + Codex MCP).
set -euo pipefail

CONFIG="${FAQ_BOT_ENV:-${HOME}/configs/faq-bot.env}"
if [[ -f "${CONFIG}" ]]; then
  # shellcheck disable=SC1090
  source "${CONFIG}"
fi

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

if [[ ! -f .env ]]; then
  cp .env.example .env
  echo "Created backend/.env from .env.example — review before production use."
fi

PORT="${PORT:-8000}"
echo "Config:  ${CONFIG}"
echo "Venv:    ${FAQ_BOT_VENV}"
echo "KB:      ${KNOWLEDGE_BASE:-"(set in backend/.env)"}"
echo "API:     http://127.0.0.1:${PORT}"
echo "Test UI: http://127.0.0.1:${PORT}/test-ui"
exec python app.py
