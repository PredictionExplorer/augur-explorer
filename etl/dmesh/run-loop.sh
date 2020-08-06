#!/bin/bash
. $HOME/configs/mesh-config.env
LOG_DIR=$HOME/ae_logs
while true
do
	nohup ./dmesh >> $LOG_DIR/mesh_nohup.log
	sleep 1
done

