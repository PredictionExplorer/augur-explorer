#!/bin/bash
. $HOME/configs/etl-config.env
LOG_DIR=$HOME/ae_logs
WD=`/bin/pwd`
while true
do
	nohup $WD/randomwalk >> $LOG_DIR/randomwalk_nohup.log
	sleep 2
done

