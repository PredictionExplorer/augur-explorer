#!/bin/bash
. $HOME/configs/etl-config.env
LOG_DIR=$HOME/ae_logs
WD=`/bin/pwd`
while true
do
	nohup $WD/layer1 >> $LOG_DIR/layer1_nohup.log
	sleep 2
done

