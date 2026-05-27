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

if [[ -z "${FAQ_BOT_VENV:-}" ]]; then
  echo "ERROR: FAQ_BOT_VENV is not set. Add it to ${CONFIG} and run scripts/setup-venv.sh" >&2
  exit 1
fi

PYTHON="${FAQ_BOT_VENV}/bin/python3"
if [[ ! -x "${PYTHON}" ]]; then
  echo "ERROR: venv python not found at ${PYTHON}" >&2
  echo "Run: scripts/setup-venv.sh" >&2
  exit 1
fi

if ! command -v git >/dev/null 2>&1; then
  echo "ERROR: system 'git' is not installed (needed to clone upstream repos)." >&2
  echo "Install with: sudo apt install git" >&2
  exit 1
fi

if ! "${PYTHON}" -c "import git" 2>/dev/null; then
  echo "ERROR: GitPython is not installed in ${FAQ_BOT_VENV}" >&2
  echo "Run: scripts/setup-venv.sh  (or: ${FAQ_BOT_VENV}/bin/pip install -r backend/requirements.txt)" >&2
  exit 1
fi

cd "$(dirname "$0")/../backend"
echo "Generating knowledge base → ${KNOWLEDGE_BASE}"
echo "Python: ${PYTHON}"
"${PYTHON}" -m knowledge.generate.run_all "$@"
echo "Re-index with: $(dirname "$0")/reindex-knowledge.sh"
echo "  (or restart backend / POST /api/reindex if the server is running)"
