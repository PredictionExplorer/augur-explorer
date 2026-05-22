#!/usr/bin/env bash
# Generate FAQ knowledge base into $KNOWLEDGE_BASE (see ~/configs/faq-bot.env).
set -euo pipefail

CONFIG="${FAQ_BOT_ENV:-${HOME}/configs/faq-bot.env}"
if [[ -f "${CONFIG}" ]]; then
  # shellcheck disable=SC1090
  source "${CONFIG}"
fi

if [[ -z "${KNOWLEDGE_BASE:-}" ]]; then
  echo "ERROR: KNOWLEDGE_BASE is not set. Add it to ${CONFIG}" >&2
  exit 1
fi

if [[ -n "${FAQ_BOT_VENV:-}" && -f "${FAQ_BOT_VENV}/bin/activate" ]]; then
  # shellcheck disable=SC1091
  source "${FAQ_BOT_VENV}/bin/activate"
fi

cd "$(dirname "$0")/../backend"
echo "Generating knowledge base → ${KNOWLEDGE_BASE}"
python3 -m knowledge.generate.run_all "$@"
echo "Re-index with: curl -X POST http://127.0.0.1:8000/api/reindex"
