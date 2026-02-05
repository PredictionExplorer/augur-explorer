#!/bin/bash
# Clean binaries in rwcg/websrv/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Cleaning websrv binaries..."

rm -f websrv

echo "Done cleaning websrv."
