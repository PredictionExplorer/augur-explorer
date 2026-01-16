#!/bin/bash

function killeverything()
{
	echo "Interrupt caught, stopping servers..."
	kill -TERM $MESH_PID
	kill -TERM $ETL_PID
	kill -TERM $SRV_PID
	wait $SRV_PID
	wait $ETL_PID
	wait $MESH_PID
}

cd server
. config/local-testnet.env
./server 1>/var/tmp/backend-info.log 2>/var/tmp/backend-error.log &
SRV_PID=$!
cd ..

cd etl
. config/local-testnet.env
./etl 1>/var/tmp/etl-info.log 2>/var/tmp/etl-error.log &
ETL_PID=$!

trap killeverything SIGINT

cd dmesh
. config/local-testnet.env
./dmesh &
MESH_PID=$!
wait $MESH_PID
echo Run complete, you can find all the errors in $HOME/ae_logs
