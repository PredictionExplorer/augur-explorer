#!/bin/bash
# Periodically scan the websrv access log for anomalies (HTTP 5xx + panics)
# and refresh $HOME/ae_logs/webserver_anomalies.log, which srvmonitor fetches
# via scp. Run this on the production host (e.g. cosmic1) as the app user.
#
# Usage: ./run-loop.sh   (build the binary first: go build -o loganomaly .)

set -u

BIN="${LOGANOMALY_BIN:-./loganomaly}"
IN="${LOGANOMALY_IN:-$HOME/ae_logs/webserver_cosmic_nohup.log}"
OUT="${LOGANOMALY_OUT:-$HOME/ae_logs/webserver_anomalies.log}"
MIN_STATUS="${LOGANOMALY_MIN_STATUS:-500}"
KEEP="${LOGANOMALY_KEEP:-50}"
INTERVAL="${LOGANOMALY_INTERVAL:-300}"

while true; do
	"$BIN" -in "$IN" -out "$OUT" -min-status "$MIN_STATUS" -keep "$KEEP" \
		|| echo "loganomaly: run failed at $(date)" >&2
	sleep "$INTERVAL"
done
