#!/bin/bash
. $HOME/configs/web-config.env
LOG_DIR=$HOME/ae_logs
while true
do
	nohup ./server >> $LOG_DIR/webserver_nohup.log
	sleep 1
done

