#!/bin/bash
# Periodically scan the websrv access log for anomalies (HTTP 5xx + panics)
# and refresh $HOME/ae_logs/webserver_anomalies.log, which srvmonitor fetches
# via scp. Run this on the production host (e.g. cosmic1) as the app user.
#
# If the loganomaly binary disappears - e.g. a redeploy replaced or removed the
# directory out from under a running loop, leaving a stale "(deleted)" cwd - the
# script exits instead of looping forever against a missing/old binary (which
# would let the anomaly feed go stale silently). Under a supervisor with
# Restart=always (systemd), it is then restarted cleanly against the new files.
#
# Usage: ./run-loop.sh   (build the binary first: go build -o loganomaly .)

set -u

# Resolve the directory this script lives in so paths don't depend on the cwd.
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" 2>/dev/null && pwd -P)"

BIN="${LOGANOMALY_BIN:-${SCRIPT_DIR:-.}/loganomaly}"
IN="${LOGANOMALY_IN:-$HOME/ae_logs/webserver_cosmic_nohup.log}"
OUT="${LOGANOMALY_OUT:-$HOME/ae_logs/webserver_anomalies.log}"
MIN_STATUS="${LOGANOMALY_MIN_STATUS:-500}"
KEEP="${LOGANOMALY_KEEP:-50}"
INTERVAL="${LOGANOMALY_INTERVAL:-300}"

while true; do
	# Exit if the binary is gone (deleted, or its directory was replaced by a
	# redeploy). Continuing would just spin against a missing/stale binary.
	if [ ! -x "$BIN" ]; then
		echo "loganomaly: binary '$BIN' missing or not executable - exiting at $(date)" >&2
		exit 1
	fi

	"$BIN" -in "$IN" -out "$OUT" -min-status "$MIN_STATUS" -keep "$KEEP" \
		|| echo "loganomaly: run failed at $(date)" >&2
	sleep "$INTERVAL"
done
