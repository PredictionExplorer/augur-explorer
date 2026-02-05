#!/bin/bash
# Clean tool binaries in rwcg/etl/randomwalk/tools/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Cleaning randomwalk tools binaries..."

rm -f rw_toprated

echo "Done cleaning randomwalk tools."
