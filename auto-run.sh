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
. config/dev-config.env
rm /var/tmp/backend-info.log 2>/dev/null;
rm /var/tmp/backend-error.log 2>/dev/null;
rm /var/tmp/backend-db.log 2>/dev/null;
./server 1>/var/tmp/backend-info.log 2>/var/tmp/backend-error.log &
SRV_PID=$!
cd ..

cd etl
. config/dev-config.env
rm /var/tmp/etl-info.log 2>/dev/null;
rm /var/tmp/etl-error.log 2>/dev/null;
rm /var/tmp/db.log 2>/dev/null;
./etl 1>/var/tmp/etl-info.log 2>/var/tmp/etl-error.log &
ETL_PID=$!

trap killeverything SIGINT

cd mesh
. config/dev-config.env
./mesh &
MESH_PID=$!
wait $MESH_PID


