#!/usr/bin/env bash
# Donates ETH to the CosmicSignatureGame prize pool (EthDonations.donateEth)
# on the local hardhat node, signing with hardhat account 0.
#
# Usage:
#   ./donate.sh            # donates 10 ETH
#   ./donate.sh 25         # donates 25 ETH (decimals OK, e.g. 0.5)
#
# Overrides via env:
#   RPC_URL   (default http://127.0.0.1:8545)
#   PKEY_HEX  (default hardhat account 0)
#   CADDR     (default 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512)

set -e -o pipefail

AMOUNT_ETH="${1:-10}"

export RPC_URL="${RPC_URL:-http://127.0.0.1:8545}"
# Hardhat account 0 (0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266), a well-known dev key.
export PKEY_HEX="${PKEY_HEX:-ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}"
CADDR="${CADDR:-0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CGCTL="$SCRIPT_DIR/../cgctl"
if [ ! -x "$CGCTL" ]; then
    echo "cgctl binary not found at $CGCTL — build it with: (cd $SCRIPT_DIR/.. && go build)" >&2
    exit 1
fi

AMOUNT_WEI="$(python3 -c "from decimal import Decimal; print(int(Decimal('$AMOUNT_ETH') * 10**18))")"

echo "Donating $AMOUNT_ETH ETH ($AMOUNT_WEI wei) to game $CADDR via $RPC_URL"
"$CGCTL" donate "$CADDR" "$AMOUNT_WEI"
