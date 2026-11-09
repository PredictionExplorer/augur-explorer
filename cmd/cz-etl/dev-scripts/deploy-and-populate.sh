#!/bin/bash
# Deploy the ChaosZero local stack to a running local chain (hardhat node or
# anvil on 127.0.0.1:8545, chain id 31337) and populate it with sample data
# covering every event cz-etl indexes.
#
# Usage:
#   ./deploy-and-populate.sh [CHAOSZERO_DIR]
#
# CHAOSZERO_DIR may also come from the environment. The checkout must have
# node_modules (pnpm install) and compiled artifacts (pnpm build, or
# forge build). Optional: RPC_URL (default http://127.0.0.1:8545).
#
# Start a local chain first, e.g. from the ChaosZero checkout:
#   anvil
# or from any hardhat project:
#   npx hardhat node

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

CHAOSZERO_DIR="${1:-${CHAOSZERO_DIR:-}}"
if [ -z "$CHAOSZERO_DIR" ]; then
	echo "Usage: $0 CHAOSZERO_DIR   (or set the CHAOSZERO_DIR env variable)" >&2
	exit 1
fi
if [ ! -f "$CHAOSZERO_DIR/src/IndependentOutcomeMarket.sol" ]; then
	echo "ERROR: $CHAOSZERO_DIR does not look like the ChaosZero repo" >&2
	exit 1
fi
if [ ! -d "$CHAOSZERO_DIR/node_modules/ethers" ]; then
	echo "ERROR: $CHAOSZERO_DIR/node_modules/ethers missing; run 'pnpm install' there first" >&2
	exit 1
fi
if [ ! -d "$CHAOSZERO_DIR/.local/hardhat-artifacts" ] && [ ! -d "$CHAOSZERO_DIR/out" ]; then
	echo "ERROR: no compiled artifacts in $CHAOSZERO_DIR; run 'pnpm build' or 'forge build' there first" >&2
	exit 1
fi

export CHAOSZERO_DIR
export NODE_PATH="$CHAOSZERO_DIR/node_modules"

echo "=== deploying ChaosZero local stack ==="
node "$SCRIPT_DIR/deploy-local.js"

echo
echo "=== populating sample data ==="
node "$SCRIPT_DIR/populate.js"
