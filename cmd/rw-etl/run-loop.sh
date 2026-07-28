#!/bin/bash
# Legacy nohup run loop. Uses a binary built in place (`go build` in this
# directory) when present, otherwise the repository bin/ (`make rw-etl`).
. $HOME/configs/rwalk-etl-config.env
LOG_DIR=$HOME/ae_logs
SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
BIN="$SCRIPT_DIR/rw-etl"
[ -x "$BIN" ] || BIN="$SCRIPT_DIR/../../bin/rw-etl"
while true
do
	nohup "$BIN" >> $LOG_DIR/randomwalk_nohup.log 2>&1
	sleep 2
done
