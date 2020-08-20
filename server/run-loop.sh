#!/bin/bash
. $HOME/configs/web-config.env
LOG_DIR=$HOME/ae_logs
WD=`/bin/pwd`
while true
do
	nohup $WD/server >> $LOG_DIR/webserver_nohup.log
	sleep 2
done

