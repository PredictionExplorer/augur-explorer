#!/bin/bash
. $HOME/configs/etl-config.env
LOG_DIR=$HOME/ae_logs
WD=`/bin/pwd`
while true
do
	nohup $WD/erc20 >> $LOG_DIR/erc20_nohup.log
	sleep 2
done

