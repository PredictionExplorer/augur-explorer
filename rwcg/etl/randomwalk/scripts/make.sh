#!/bin/bash
# Build script for rwcg/etl/randomwalk/scripts/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Building randomwalk scripts..."

FAILED=0

build() {
    local name=$1
    echo -n "  Building $name... "
    if go build -o "$name" "${name}.go" 2>/dev/null; then
        echo "OK"
    else
        echo "FAILED"
        FAILED=1
    fi
}

build accept_offer
build approve
build cancel_offer
build mint
build new_offer
build ownerof
build price
build scan_rwmints
build scan_transfers
build setname
build status
build statusmkt
build tokenuri
build transfer
build verify_erc20_transfers
build verify_owner
build withdrawal

if [ $FAILED -eq 0 ]; then
    echo "All randomwalk scripts built successfully!"
else
    echo "Some builds failed!"
    exit 1
fi
