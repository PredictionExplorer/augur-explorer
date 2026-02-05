#!/bin/bash
# Clean script binaries in rwcg/etl/randomwalk/scripts/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Cleaning randomwalk scripts binaries..."

rm -f mint
rm -f ownerof
rm -f price
rm -f status

echo "Done cleaning randomwalk scripts."
