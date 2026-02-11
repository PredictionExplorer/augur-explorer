#!/bin/bash
# Build script for rwcg/etl/cosmicgame/imggen_monitor/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Building imggen_monitor binaries..."

FAILED=0

echo -n "  imggen_exec... "
if go build -o imggen_exec imggen_exec.go; then
    echo "OK"
else
    echo "FAILED"
    FAILED=1
fi

echo -n "  imggen_monitor... "
if go build -o imggen_monitor imggen_monitor.go; then
    echo "OK"
else
    echo "FAILED"
    FAILED=1
fi

if [ $FAILED -eq 0 ]; then
    echo "Build successful: imggen_exec, imggen_monitor"
else
    echo "Build failed!"
    exit 1
fi
