#!/bin/bash
. $HOME/configs/etl-config.env
LOG_DIR=$HOME/ae_logs
WD=`/bin/pwd`
while true
do
	nohup $WD/biddingwar >> $LOG_DIR/biddingwar_nohup.log
	sleep 2
done
