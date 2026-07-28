#!/bin/bash
# Legacy nohup run loop. Uses a binary built in place (`go build` in this
# directory) when present, otherwise the repository bin/ (`make apiserver`).
. $HOME/configs/cosmic-api-config.env
LOG_DIR=$HOME/ae_logs
SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
BIN="$SCRIPT_DIR/apiserver"
[ -x "$BIN" ] || BIN="$SCRIPT_DIR/../../bin/apiserver"
while true
do
	nohup "$BIN" >> $LOG_DIR/webserver_cosmic_nohup.log 2>&1
	sleep 2
done
