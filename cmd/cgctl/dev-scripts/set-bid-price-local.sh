#!/usr/bin/env bash
# Sets the next ETH bid price on the local hardhat game contract by writing
# contract storage (hardhat_setStorageAt). Default: 1 ETH.
#
# Usage:
#   ./set-bid-price-local.sh              # 1 ETH
#   BID_PRICE_ETH=0.5 ./set-bid-price-local.sh
#
# Works both mid-round (sets nextEthBidPrice) and before the first bid of a
# round (calibrates the ETH Dutch auction beginning price so the effective
# price equals the target at calibration time).

set -e -o pipefail

# Solidity repo with hardhat.config.js (override with SOLIDITY_REPO env var).
SOLIDITY_REPO="${SOLIDITY_REPO:-$HOME/eth/dev/b/cg-v3/Cosmic-Signature-v3.1-2026-08-19}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

export CADDR="${CADDR:-0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512}"
export BID_PRICE_ETH="${BID_PRICE_ETH:-1}"

cd "$SOLIDITY_REPO"
npx hardhat run "$SCRIPT_DIR/set-bid-price.js" --network hardhat_on_localhost
