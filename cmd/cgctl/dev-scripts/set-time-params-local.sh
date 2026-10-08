#!/usr/bin/env bash
# Applies quick dev timing parameters to the local hardhat game contract:
#   - per-bid main prize time bump:            2 minutes
#   - anyone-can-claim timeout:                5 minutes
#   - initial countdown armed by first bid:    5 minutes
#
# Usage:
#   ./set-time-params-local.bash
#
# Override any value via env, e.g.:
#   TIME_INCREMENT_SEC=300 ./set-time-params-local.bash
#
# Note: these are inactive-round-only setters. Run right after a main prize
# claim, BEFORE the first bid of the new round (set-time-params.js briefly
# deactivates/reactivates the round, which only works while it is bid-free).

set -e -o pipefail

# Solidity repo with hardhat.config.js (override with SOLIDITY_REPO env var).
SOLIDITY_REPO="${SOLIDITY_REPO:-$HOME/eth/dev/b/cg-v3/Cosmic-Signature-v3.1-2026-08-19}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

export CADDR="${CADDR:-0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512}"
export TIME_INCREMENT_SEC="${TIME_INCREMENT_SEC:-120}"
export TIMEOUT_CLAIM_SEC="${TIMEOUT_CLAIM_SEC:-300}"
export INITIAL_DURATION_SEC="${INITIAL_DURATION_SEC:-300}"

cd "$SOLIDITY_REPO"
npx hardhat run "$SCRIPT_DIR/set-time-params.js" --network hardhat_on_localhost
