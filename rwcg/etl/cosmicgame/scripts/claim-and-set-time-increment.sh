#!/bin/bash
# Claims prize, then sets the bid time increment

set -e

if [ "$#" -lt 2 ] || [ "$#" -gt 3 ]; then
    echo "Usage: $0 <contract_addr> <time_increment_seconds> [rpc_url]"
    echo ""
    echo "  contract_addr          - CosmicSignatureGame contract address"
    echo "  time_increment_seconds - Time added per bid (e.g., 120 for 2 min)"
    echo "  rpc_url                - Optional: RPC URL (default: \$RPC_URL env var)"
    echo ""
    echo "Environment: PKEY_HEX (64-char hex private key) and RPC_URL must be set."
    echo ""
    echo "Example:"
    echo "  export PKEY_HEX=ac0974bec...f2ff80"
    echo "  $0 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512 120"
    exit 1
fi

CONTRACT_ADDR="$1"
TIME_INCREMENT="$2"
RPC_URL="${3:-$RPC_URL}"

if [ -z "$RPC_URL" ]; then
    echo "Error: RPC_URL not set. Provide as 3rd argument or set RPC_URL env var."
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
echo "Claim and Set Time Increment Script"
echo "=========================================="
echo "Contract:          $CONTRACT_ADDR"
echo "Time Increment:    $TIME_INCREMENT seconds per bid"
echo "RPC URL:           $RPC_URL"
echo "=========================================="
echo ""

# Build scripts if needed
echo "Building scripts..."
cd "$SCRIPT_DIR"
go build -o claimprize claimprize.go 2>/dev/null || true
go build -o set-time-increment set-time-increment.go 2>/dev/null || true

# Step 1: Claim prize
echo ""
echo "Step 1: Claiming prize..."
echo "------------------------------------------------------------------------"
#./claimprize "$CONTRACT_ADDR"

# Wait a bit for transaction to be mined
echo ""
echo "Waiting for claim transaction to be confirmed..."
sleep 3

# Step 2: Set time increment
echo ""
echo "Step 2: Setting time increment to $TIME_INCREMENT seconds per bid..."
echo "------------------------------------------------------------------------"
./set-time-increment "$CONTRACT_ADDR" "$TIME_INCREMENT"

echo ""
echo "=========================================="
echo "Done!"
echo "- Prize claimed"
echo "- Each bid will extend time until main prize by $TIME_INCREMENT seconds"
echo "=========================================="
