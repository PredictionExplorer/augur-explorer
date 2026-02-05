#!/bin/bash
# Build script for rwcg/etl/randomwalk/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Building randomwalk ETL..."

go build -o rw_etl .

if [ $? -eq 0 ]; then
    echo "Build successful: rw_etl"
else
    echo "Build failed!"
    exit 1
fi
