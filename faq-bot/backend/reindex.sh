#!/usr/bin/env bash
# Re-index Haystack from KNOWLEDGE_BASE using the FAQ bot venv.
exec "$(cd "$(dirname "$0")/.." && pwd)/scripts/reindex-knowledge.sh" "$@"
