#!/bin/bash
. $HOME/configs/etl-config.env
LOG_DIR=$HOME/ae_logs
while true
do
	nohup ./etl >> $LOG_DIR/etl_nohup.log
	sleep 1
done

