#!/usr/bin/env bash
# Rebuild Haystack index from $KNOWLEDGE_BASE (offline — no running backend required).
set -euo pipefail

CONFIG="${FAQ_BOT_ENV:-${HOME}/configs/faq-bot.env}"
if [[ -f "${CONFIG}" ]]; then
  # shellcheck disable=SC1090
  source "${CONFIG}"
fi

if [[ -z "${FAQ_BOT_VENV:-}" ]]; then
  echo "ERROR: FAQ_BOT_VENV is not set. Add it to ${CONFIG}" >&2
  exit 1
fi

PYTHON="${FAQ_BOT_VENV}/bin/python3"
if [[ ! -x "${PYTHON}" ]]; then
  echo "ERROR: venv python not found at ${PYTHON}. Run scripts/setup-venv.sh" >&2
  exit 1
fi

cd "$(dirname "$0")/../backend"
echo "Re-indexing Haystack from KNOWLEDGE_BASE=${KNOWLEDGE_BASE:-"(from backend/.env)"}"
echo "Python: ${PYTHON}"
"${PYTHON}" -m retrieval.pipeline --reindex
