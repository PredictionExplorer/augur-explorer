#!/bin/bash
# Root clean script for rwcg/
# Cleans all binaries hierarchically by calling child make-clean.sh scripts

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "========================================="
echo "Cleaning all rwcg binaries..."
echo "========================================="

# Clean root-level binaries
rm -f cosmicgame

# Call child make-clean.sh scripts hierarchically
if [ -x etl/cosmicgame/make-clean.sh ]; then
    etl/cosmicgame/make-clean.sh
fi

if [ -x etl/randomwalk/make-clean.sh ]; then
    etl/randomwalk/make-clean.sh
fi

if [ -x websrv/make-clean.sh ]; then
    websrv/make-clean.sh
fi

if [ -x freezer-scanner/freezer-scan/make-clean.sh ]; then
    freezer-scanner/freezer-scan/make-clean.sh
fi

if [ -x tools/make-clean.sh ]; then
    tools/make-clean.sh
fi

echo "========================================="
echo "All rwcg binaries cleaned!"
echo "========================================="
