#!/bin/bash
. $HOME/configs/web-config.env
LOG_DIR=$HOME/ae_logs
WD=`/bin/pwd`
while true
do
	nohup $WD/imggen_monitor --regenerate=true > $LOG_DIR/imggen_monitor_latest.log
	sleep 3600
done

