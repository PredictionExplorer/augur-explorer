#!/bin/bash
# Clean binaries in rwcg/etl/cosmicgame/imggen_monitor/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Cleaning imggen_monitor binaries..."

rm -f imggen_exec
rm -f imggen_monitor

echo "Done cleaning imggen_monitor."
