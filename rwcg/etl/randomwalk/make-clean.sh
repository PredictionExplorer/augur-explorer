#!/bin/bash
# Clean binaries in rwcg/etl/randomwalk/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Cleaning etl/randomwalk binaries..."

rm -f rw_etl

# Call child make-clean.sh scripts
if [ -x scripts/make-clean.sh ]; then
    scripts/make-clean.sh
fi

if [ -x tools/make-clean.sh ]; then
    tools/make-clean.sh
fi

echo "Done cleaning etl/randomwalk."
