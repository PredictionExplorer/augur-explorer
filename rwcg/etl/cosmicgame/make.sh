#!/bin/bash
# Build script for rwcg/etl/cosmicgame/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Building cosmicgame ETL..."

go build -o cg_etl .

if [ $? -eq 0 ]; then
    echo "Build successful: cg_etl"
else
    echo "Build failed!"
    exit 1
fi
