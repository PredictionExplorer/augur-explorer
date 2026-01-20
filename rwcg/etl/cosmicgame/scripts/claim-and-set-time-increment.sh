#!/bin/bash
# Claims prize, then sets the bid time increment

set -e

if [ "$#" -lt 3 ] || [ "$#" -gt 4 ]; then
    echo "Usage: $0 <private_key> <contract_addr> <time_increment_seconds> [rpc_url]"
    echo ""
    echo "  private_key            - 64 character hex private key (owner)"
    echo "  contract_addr          - CosmicSignatureGame contract address"
    echo "  time_increment_seconds - Time added per bid (e.g., 120 for 2 min)"
    echo "  rpc_url                - Optional: RPC URL (default: \$RPC_URL env var)"
    echo ""
    echo "Example:"
    echo "  $0 ac0974bec...f2ff80 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512 120"
    exit 1
fi

PRIV_KEY="$1"
CONTRACT_ADDR="$2"
TIME_INCREMENT="$3"
RPC_URL="${4:-$RPC_URL}"

if [ -z "$RPC_URL" ]; then
    echo "Error: RPC_URL not set. Provide as 4th argument or set RPC_URL env var."
    exit 1
fi

export RPC_URL

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
#./claimprize "$PRIV_KEY" "$CONTRACT_ADDR"

# Wait a bit for transaction to be mined
echo ""
echo "Waiting for claim transaction to be confirmed..."
sleep 3

# Step 2: Set time increment
echo ""
echo "Step 2: Setting time increment to $TIME_INCREMENT seconds per bid..."
echo "------------------------------------------------------------------------"
./set-time-increment "$PRIV_KEY" "$CONTRACT_ADDR" "$TIME_INCREMENT"

echo ""
echo "=========================================="
echo "Done!"
echo "- Prize claimed"
echo "- Each bid will extend time until main prize by $TIME_INCREMENT seconds"
echo "=========================================="
