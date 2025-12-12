#!/bin/bash
. $HOME/configs/etl-config.env
LOG_DIR=$HOME/ae_logs
WD=`/bin/pwd`
while true
do
	nohup $WD/cosmicgame >> $LOG_DIR/cosmicgame_nohup.log
	sleep 2
done
