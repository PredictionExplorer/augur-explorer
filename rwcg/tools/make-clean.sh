#!/bin/bash
# Clean binaries in rwcg/tools/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Cleaning tools binaries..."

rm -f archive_export
rm -f db_verify
rm -f evtlog_diff

echo "Done cleaning tools."
