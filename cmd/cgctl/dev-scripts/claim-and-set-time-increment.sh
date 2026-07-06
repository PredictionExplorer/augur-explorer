#!/bin/bash
# Sets mainPrizeTimeIncrementInMicroSeconds in one run. Claims the prize first only when
# the contract requires it to open the inactive admin window.
# V1/V2 CosmicSignatureGame proxy compatible.
#
# Contract rules:
#   - setMainPrizeTimeIncrementInMicroSeconds needs an inactive round
#   - claimMainPrize needs bids in the current round (reverts with NoBidsPlacedInCurrentRound otherwise)
#   - If the round is active with no claimable prize, the script defers activation instead of claiming
#
# Requires PKEY_HEX for the contract owner (Hardhat #0 on local: 0xf39F...).
# When claiming, the same key must also be the last bidder (or wait for the claim timeout).
# Requires cgctl on PATH (or set CGCTL): go build -o cgctl ./cmd/cgctl
#
# Usage:
#   export RPC_URL=http://127.0.0.1:8545
#   export PKEY_HEX=<64-char hex, no 0x>
#   ./claim-and-set-time-increment.sh <contract_addr> <time_increment_seconds> [delay_seconds]

set -e

if [ "$#" -lt 2 ] || [ "$#" -gt 4 ]; then
    echo "Usage: $0 <contract_addr> <time_increment_seconds> [delay_seconds] [rpc_url]"
    echo ""
    echo "  contract_addr            - CosmicSignatureGame proxy address"
    echo "  time_increment_seconds   - Seconds added to main prize time per bid (e.g. 240)"
    echo "  delay_seconds            - Optional delay before next round (default: 300)"
    echo "  rpc_url                  - Optional RPC URL (default: \$RPC_URL env var)"
    echo ""
    echo "Environment: PKEY_HEX and RPC_URL must be set (unless rpc_url is passed)."
    echo ""
    echo "Example:"
    echo "  export RPC_URL=http://127.0.0.1:8545"
    echo "  export PKEY_HEX=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
    echo "  $0 0x5c74c94173F05dA1720953407cbb920F3DF9f887 240 300"
    exit 1
fi

CGCTL="${CGCTL:-cgctl}"
CONTRACT_ADDR="$1"
TIME_INCREMENT="$2"
DELAY_SECONDS="${3:-300}"
RPC_URL="${4:-$RPC_URL}"

if [ -z "$RPC_URL" ]; then
    echo "Error: RPC_URL not set. Provide as 4th argument or set RPC_URL env var."
    exit 1
fi
if [ -z "$PKEY_HEX" ]; then
    echo "Error: PKEY_HEX not set (64-char hex private key, no 0x prefix)."
    exit 1
fi

export RPC_URL
export PKEY_HEX

echo "=========================================="
echo "Claim + Set Time Increment"
echo "=========================================="
echo "Contract:        $CONTRACT_ADDR"
echo "Time increment:  $TIME_INCREMENT seconds per bid"
echo "Round delay:     $DELAY_SECONDS seconds before next round"
echo "RPC URL:         $RPC_URL"
echo "=========================================="
echo ""

exec "$CGCTL" claim-and-set-time-increment "$CONTRACT_ADDR" "$TIME_INCREMENT" "$DELAY_SECONDS"
