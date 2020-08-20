#!/bin/bash
. $HOME/configs/etl-config.env
LOG_DIR=$HOME/ae_logs
WD=`/bin/pwd`
while true
do
	nohup $WD/etl >> $LOG_DIR/etl_nohup.log
	sleep 10
done

