#!/bin/bash
# Claims prize with a round-activation delay, then sets initialDurationUntilMainPrizeDivisor.
# Works on the CosmicSignatureGame proxy after V1 or V2 upgrade.
#
# IMPORTANT: the second argument to set-initial-duration-divisor is a DIVISOR, not seconds.
# Initial timer after first bid (microseconds) = mainPrizeTimeIncrementInMicroSeconds / divisor.
# Example: increment 240s (240000000 µs) and divisor 1000000 → ~240s initial duration.
#
# Requires cgctl on PATH (or set CGCTL): go build -o cgctl ./cmd/cgctl

set -e

if [ "$#" -lt 2 ] || [ "$#" -gt 4 ]; then
    echo "Usage: $0 <contract_addr> <initial_duration_seconds> [delay_seconds] [rpc_url]"
    echo ""
    echo "  contract_addr            - CosmicSignatureGame proxy address"
    echo "  initial_duration_divisor - initialDurationUntilMainPrizeDivisor (e.g. 1000000)"
    echo "  delay_seconds            - Optional: delay before round activation (default: 120)"
    echo "  rpc_url                - Optional: RPC URL (default: \$RPC_URL env var)"
    echo ""
    echo "Environment: PKEY_HEX (64-char hex private key) and RPC_URL must be set."
    echo ""
    echo "Example:"
    echo "  export PKEY_HEX=ac0974bec...f2ff80"
    echo "  $0 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512 1000000 120"
    exit 1
fi

CGCTL="${CGCTL:-cgctl}"
CONTRACT_ADDR="$1"
INITIAL_DURATION="$2"
DELAY_SECONDS="${3:-120}"
RPC_URL="${4:-$RPC_URL}"

if [ -z "$RPC_URL" ]; then
    echo "Error: RPC_URL not set. Provide as 4th argument or set RPC_URL env var."
    exit 1
fi
if [ -z "$PKEY_HEX" ]; then
    echo "Error: PKEY_HEX not set (64-char hex private key)."
    exit 1
fi

export RPC_URL
export PKEY_HEX

echo "=========================================="
echo "Claim and Configure Script"
echo "=========================================="
echo "Contract:          $CONTRACT_ADDR"
echo "Delay:             $DELAY_SECONDS seconds"
echo "Initial Duration:  $INITIAL_DURATION seconds"
echo "RPC URL:           $RPC_URL"
echo "=========================================="
echo ""

# Step 1: Set delay and claim prize
echo ""
echo "Step 1: Setting delay to $DELAY_SECONDS seconds and claiming prize..."
echo "------------------------------------------------------------------------"
"$CGCTL" claim-prize "$CONTRACT_ADDR" --delay "$DELAY_SECONDS"

# Wait a bit for transaction to be mined
echo ""
echo "Waiting for claim transaction to be confirmed..."
sleep 3

# Step 2: Set initial duration divisor (during the delay window)
echo ""
echo "Step 2: Setting initial duration to $INITIAL_DURATION seconds..."
echo "------------------------------------------------------------------------"
"$CGCTL" set-initial-duration-divisor "$CONTRACT_ADDR" "$INITIAL_DURATION"

echo ""
echo "=========================================="
echo "Done!"
echo "- Prize claimed"
echo "- Round will activate in $DELAY_SECONDS seconds"
echo "- After first bid, time until main prize will be ~$INITIAL_DURATION seconds"
echo "=========================================="
