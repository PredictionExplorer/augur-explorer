#!/usr/bin/env bash
# Start the FAQ bot Python backend (Haystack + Codex MCP).
set -euo pipefail
cd "$(dirname "$0")/../backend"

if [[ ! -d venv ]]; then
  echo "Creating venv..."
  python3 -m venv venv
  source venv/bin/activate
  pip install -r requirements.txt
else
  source venv/bin/activate
fi

if [[ ! -f .env ]]; then
  cp .env.example .env
  echo "Created .env from .env.example — review paths before production use."
fi

echo "FAQ bot API: http://127.0.0.1:${PORT:-8000}"
echo "Test UI:       http://127.0.0.1:${PORT:-8000}/test-ui"
exec python app.py
