#!/bin/bash
# Deploy CosmicSignatureGame with dev time settings + Samp tokens, then run populate.js.
# Usage:
#   ./deploy-and-populate.sh [--network NETWORK] [COSMIC_SIGNATURE_DIR]
#
# Defaults: network=localhost, COSMIC_SIGNATURE_DIR=$COSMIC_SIGNATURE_DIR or parent/todays/Cosmic-Signature.
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
POPULATE_JS="$SCRIPT_DIR/populate.js"
DEPLOY_DEV_JS="$SCRIPT_DIR/deploy-dev-and-samp.js"
NETWORK="${NETWORK:-localhost}"

# Parse args: [--network NETWORK] [COSMIC_SIGNATURE_DIR]
COSMIC_SIGNATURE_DIR=""
while [ $# -gt 0 ]; do
    case "$1" in
        --network)
            NETWORK="$2"
            shift 2
            ;;
        *)
            COSMIC_SIGNATURE_DIR="$1"
            shift
            ;;
    esac
done

if [ -z "$COSMIC_SIGNATURE_DIR" ]; then
    COSMIC_SIGNATURE_DIR="${COSMIC_SIGNATURE_DIR:-}"
fi
if [ -z "$COSMIC_SIGNATURE_DIR" ]; then
    echo "Error: COSMIC_SIGNATURE_DIR required (path to Cosmic-Signature repo)."
    echo "Usage: $0 [--network NETWORK] COSMIC_SIGNATURE_DIR"
    echo "Example: $0 --network localhost /home/niko/eth/dev/b/todays/Cosmic-Signature"
    exit 1
fi

if [ ! -f "$COSMIC_SIGNATURE_DIR/scripts/Deploy.js" ]; then
    echo "Error: Not a Cosmic-Signature repo (missing scripts/Deploy.js): $COSMIC_SIGNATURE_DIR"
    exit 1
fi

if [ ! -f "$POPULATE_JS" ]; then
    echo "Error: populate.js not found: $POPULATE_JS"
    exit 1
fi

echo "=========================================="
echo "Deploy (dev times) + Samp + Populate"
echo "=========================================="
echo "Cosmic-Signature: $COSMIC_SIGNATURE_DIR"
echo "Network:           $NETWORK"
echo "=========================================="

# 1) Deploy game with dev times and Samp from Cosmic-Signature cwd
echo ""
echo "Step 1: Deploy CosmicSignatureGame (dev times) and Samp..."
echo "------------------------------------------------------------------------"
cd "$COSMIC_SIGNATURE_DIR"
OUTPUT=$(NODE_PATH=$PWD/node_modules npx hardhat run "$DEPLOY_DEV_JS" --network "$NETWORK" 2>&1)
echo "$OUTPUT"

CADDR=$(echo "$OUTPUT" | sed -n 's/^CADDR=\(0x[0-9a-fA-F]\{40\}\)$/\1/p' | tail -1)
TSAMP1=$(echo "$OUTPUT" | sed -n 's/^TSAMP1=\(0x[0-9a-fA-F]\{40\}\)$/\1/p' | tail -1)
TSAMP2=$(echo "$OUTPUT" | sed -n 's/^TSAMP2=\(0x[0-9a-fA-F]\{40\}\)$/\1/p' | tail -1)

if [ -z "$CADDR" ] || [ -z "$TSAMP1" ] || [ -z "$TSAMP2" ]; then
    echo "Error: Could not parse CADDR, TSAMP1, TSAMP2 from deploy output."
    echo "Last 30 lines of deploy output:"
    echo "$OUTPUT" | tail -30
    exit 1
fi

echo ""
echo "Parsed: CADDR=$CADDR TSAMP1=$TSAMP1 TSAMP2=$TSAMP2"

# 2) Run populate.js with env set
echo ""
echo "Step 2: Running populate.js..."
echo "------------------------------------------------------------------------"
export CADDR
export TSAMP1
export TSAMP2
NODE_PATH="$COSMIC_SIGNATURE_DIR/node_modules" npx hardhat run "$POPULATE_JS" --network "$NETWORK"

echo ""
echo "=========================================="
echo "Done: deployed and populated."
echo "=========================================="
