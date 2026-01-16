#!/bin/bash
. $HOME/configs/etl-config.env
. $HOME/configs/twitter-prod.env
. $HOME/configs/discord-prod.env
LOG_DIR=$HOME/ae_logs
WD=`/bin/pwd`
while true
do
	nohup $WD/notibot --twitter --discord >> $LOG_DIR/notibot_nohup.log
	sleep 2
done

