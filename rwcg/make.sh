#!/bin/bash
# Build all binaries under rwcg/
# Invokes child make.sh where present, otherwise runs go build in each directory that produces binaries.

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

FAILED=0

echo "========================================="
echo "Building all rwcg binaries..."
echo "========================================="

# ETL: cosmicgame
if [ -x etl/cosmicgame/make.sh ]; then
    if etl/cosmicgame/make.sh; then
        echo ""
    else
        FAILED=1
    fi
else
    echo "etl/cosmicgame/make.sh not found or not executable"
    FAILED=1
fi

# ETL: randomwalk
if [ -x etl/randomwalk/make.sh ]; then
    if etl/randomwalk/make.sh; then
        echo ""
    else
        FAILED=1
    fi
else
    echo "etl/randomwalk/make.sh not found or not executable"
    FAILED=1
fi

# websrv
echo "Building websrv..."
if (cd websrv && go build -o websrv .); then
    echo "Build successful: websrv/websrv"
else
    echo "Build failed: websrv"
    FAILED=1
fi
echo ""

# freezer-scanner
echo "Building freezer-scan..."
if (cd freezer-scanner/freezer-scan && go build -o freezer-scan .); then
    echo "Build successful: freezer-scanner/freezer-scan/freezer-scan"
else
    echo "Build failed: freezer-scanner/freezer-scan"
    FAILED=1
fi
echo ""

# tools (multiple mains: archive_export, db_verify, evtlog_diff)
echo "Building tools..."
if (cd tools && go build -o archive_export archive_export.go && go build -o db_verify db_verify.go && go build -o evtlog_diff evtlog_diff.go); then
    echo "Build successful: tools/archive_export, tools/db_verify, tools/evtlog_diff"
else
    echo "Build failed: tools"
    FAILED=1
fi
echo ""

echo "========================================="
if [ $FAILED -eq 0 ]; then
    echo "All rwcg binaries built successfully."
else
    echo "One or more builds failed."
    exit 1
fi
echo "========================================="
