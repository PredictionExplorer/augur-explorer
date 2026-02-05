#!/bin/bash
# Build script for rwcg/etl/cosmicgame/scripts/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Building cosmicgame scripts..."

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

build approved
build autobid
build bid
build cginfo
build claimprize
build claimprize-delay60
build deploy-erc20
build donate
build dwi-records
build erc20allowance
build erc20approve
build erc20bal
build erc20revoke
build isapproved4all
build owner
build set_charity_percentage
build set_main_prize_percentage
build set_num_nft_winners
build set_num_raffle_winners
build set_raffle_percentage
build set_staking_percentage
build set_token_name
build set-initial-duration-divisor
build set-time-increment
build setactivation
build setroundactivation
build tokownerof

if [ $FAILED -eq 0 ]; then
    echo "All cosmicgame scripts built successfully!"
else
    echo "Some builds failed!"
    exit 1
fi
