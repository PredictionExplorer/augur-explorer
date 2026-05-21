#!/usr/bin/env bash
# Generate FAQ knowledge base into $KNOWLEDGE_BASE (see backend/.env).
set -euo pipefail
cd "$(dirname "$0")/../backend"

if [[ -f .env ]]; then
  set -a
  # shellcheck disable=SC1091
  source .env
  set +a
fi

if [[ -z "${KNOWLEDGE_BASE:-}" ]]; then
  echo "ERROR: KNOWLEDGE_BASE is not set. Add it to backend/.env or export it." >&2
  exit 1
fi

echo "Generating knowledge base → ${KNOWLEDGE_BASE}"
python3 -m knowledge.generate.run_all "$@"
echo "Re-index with: curl -X POST http://127.0.0.1:8000/api/reindex"
