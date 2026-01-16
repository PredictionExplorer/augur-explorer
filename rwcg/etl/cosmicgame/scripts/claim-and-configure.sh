#!/bin/bash
# Claims prize with delay, then configures initial duration divisor during the delay window

set -e

if [ "$#" -lt 3 ] || [ "$#" -gt 5 ]; then
    echo "Usage: $0 <private_key> <contract_addr> <initial_duration_seconds> [delay_seconds] [rpc_url]"
    echo ""
    echo "  private_key            - 64 character hex private key (owner)"
    echo "  contract_addr          - CosmicSignatureGame contract address"
    echo "  initial_duration_seconds - Time until main prize after first bid (e.g., 300 for 5 min)"
    echo "  delay_seconds          - Optional: delay before round activation (default: 120)"
    echo "  rpc_url                - Optional: RPC URL (default: \$RPC_URL env var)"
    echo ""
    echo "Example:"
    echo "  $0 ac0974bec...f2ff80 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512 300 120"
    exit 1
fi

PRIV_KEY="$1"
CONTRACT_ADDR="$2"
INITIAL_DURATION="$3"
DELAY_SECONDS="${4:-120}"
RPC_URL="${5:-$RPC_URL}"

if [ -z "$RPC_URL" ]; then
    echo "Error: RPC_URL not set. Provide as 5th argument or set RPC_URL env var."
    exit 1
fi

export RPC_URL

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
./claimprize-delay60 "$PRIV_KEY" "$CONTRACT_ADDR" "$DELAY_SECONDS"

# Wait a bit for transaction to be mined
echo ""
echo "Waiting for claim transaction to be confirmed..."
sleep 3

# Step 2: Set initial duration divisor (during the delay window)
echo ""
echo "Step 2: Setting initial duration to $INITIAL_DURATION seconds..."
echo "------------------------------------------------------------------------"
./set-initial-duration-divisor "$PRIV_KEY" "$CONTRACT_ADDR" "$INITIAL_DURATION"

echo ""
echo "=========================================="
echo "Done!"
echo "- Prize claimed"
echo "- Round will activate in $DELAY_SECONDS seconds"
echo "- After first bid, time until main prize will be ~$INITIAL_DURATION seconds"
echo "=========================================="

