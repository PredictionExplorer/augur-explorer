#!/bin/bash
# Clean binaries in rwcg/freezer-scanner/freezer-scan/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Cleaning freezer-scan binaries..."

rm -f freezer-scan

echo "Done cleaning freezer-scan."
