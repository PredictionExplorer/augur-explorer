#!/bin/bash
# Claims prize with delay, then configures initial duration divisor during the delay window

set -e

if [ "$#" -lt 2 ] || [ "$#" -gt 4 ]; then
    echo "Usage: $0 <contract_addr> <initial_duration_seconds> [delay_seconds] [rpc_url]"
    echo ""
    echo "  contract_addr          - CosmicSignatureGame contract address"
    echo "  initial_duration_seconds - Time until main prize after first bid (e.g., 300 for 5 min)"
    echo "  delay_seconds          - Optional: delay before round activation (default: 120)"
    echo "  rpc_url                - Optional: RPC URL (default: \$RPC_URL env var)"
    echo ""
    echo "Environment: PKEY_HEX (64-char hex private key) and RPC_URL must be set."
    echo ""
    echo "Example:"
    echo "  export PKEY_HEX=ac0974bec...f2ff80"
    echo "  $0 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512 300 120"
    exit 1
fi

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

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "=========================================="
echo "Claim and Configure Script"
echo "=========================================="
echo "Contract:          $CONTRACT_ADDR"
echo "Delay:             $DELAY_SECONDS seconds"
echo "Initial Duration:  $INITIAL_DURATION seconds"
echo "RPC URL:           $RPC_URL"
echo "=========================================="
echo ""

# Build scripts if needed
echo "Building scripts..."
cd "$SCRIPT_DIR"
go build -o claimprize-delay60 claimprize-delay60.go 2>/dev/null || true
go build -o set-initial-duration-divisor set-initial-duration-divisor.go 2>/dev/null || true

# Step 1: Set delay and claim prize
echo ""
echo "Step 1: Setting delay to $DELAY_SECONDS seconds and claiming prize..."
echo "------------------------------------------------------------------------"
./claimprize-delay60 "$CONTRACT_ADDR" "$DELAY_SECONDS"

# Wait a bit for transaction to be mined
echo ""
echo "Waiting for claim transaction to be confirmed..."
sleep 3

# Step 2: Set initial duration divisor (during the delay window)
echo ""
echo "Step 2: Setting initial duration to $INITIAL_DURATION seconds..."
echo "------------------------------------------------------------------------"
./set-initial-duration-divisor "$CONTRACT_ADDR" "$INITIAL_DURATION"

echo ""
echo "=========================================="
echo "Done!"
echo "- Prize claimed"
echo "- Round will activate in $DELAY_SECONDS seconds"
echo "- After first bid, time until main prize will be ~$INITIAL_DURATION seconds"
echo "=========================================="

