#!/bin/bash
# Clean binaries in rwcg/etl/cosmicgame/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Cleaning etl/cosmicgame binaries..."

rm -f cg_etl
rm -f cosmicgame

# Call child make-clean.sh
if [ -x scripts/make-clean.sh ]; then
    scripts/make-clean.sh
fi

echo "Done cleaning etl/cosmicgame."
