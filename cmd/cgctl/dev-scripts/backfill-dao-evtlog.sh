#!/bin/bash
# Backfill cosmic_dao evt_log rows on cgprod (or any env with RPC_URL + PGSQL_*).
#
# Usage:
#   ./backfill-dao-evtlog.sh
#   ./backfill-dao-evtlog.sh --from-block 455767500 --to-block 470000000
#
# Requires the cgctl binary on PATH (or set CGCTL to its location):
#   go build -o cgctl ./cmd/cgctl

set -euo pipefail

CGCTL="${CGCTL:-cgctl}"
CONFIG="${CG_ENV:-$HOME/configs/cg-prod.env}"

if [[ ! -f "$CONFIG" ]]; then
	echo "config not found: $CONFIG (set CG_ENV to override)" >&2
	exit 1
fi
# shellcheck source=/dev/null
source "$CONFIG"

if ! command -v "$CGCTL" >/dev/null 2>&1; then
	echo "cgctl not found; build it with: go build -o cgctl ./cmd/cgctl (or set CGCTL)" >&2
	exit 1
fi

exec "$CGCTL" backfill-dao-evtlog "$@"
