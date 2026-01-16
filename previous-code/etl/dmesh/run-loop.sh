#!/bin/bash
. $HOME/configs/mesh-config.env
LOG_DIR=$HOME/ae_logs
WD=`/bin/pwd`
while true
do
	nohup $WD/dmesh >> $LOG_DIR/mesh_nohup.log
	sleep 1
done

