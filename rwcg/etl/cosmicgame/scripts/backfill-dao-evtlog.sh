#!/bin/bash
# Backfill cosmic_dao evt_log rows on cgprod (or any env with RPC_URL + PGSQL_*).
#
# Usage:
#   ./backfill-dao-evtlog.sh
#   ./backfill-dao-evtlog.sh -from-block 455767500 -to-block 470000000

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CONFIG="${CG_ENV:-$HOME/configs/cg-prod.env}"

if [[ ! -f "$CONFIG" ]]; then
	echo "config not found: $CONFIG (set CG_ENV to override)" >&2
	exit 1
fi
# shellcheck source=/dev/null
source "$CONFIG"

BIN="$SCRIPT_DIR/backfill_dao_evtlog"
if [[ ! -x "$BIN" ]]; then
	echo "binary missing; run: cd $SCRIPT_DIR && ./make.sh" >&2
	exit 1
fi

exec "$BIN" "$@"
