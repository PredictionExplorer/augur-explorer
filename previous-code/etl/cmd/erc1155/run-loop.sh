#!/bin/bash
. $HOME/configs/etl-config.env
LOG_DIR=$HOME/ae_logs
WD=`/bin/pwd`
while true
do
	nohup $WD/erc1155 >> $LOG_DIR/erc1155_nohup.log
	sleep 2
done

