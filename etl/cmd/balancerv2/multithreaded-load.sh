#!/bin/bash
# periodically kills the process to allow the update of 'last block' variable
ONE_HOUR=3600
while true
do
	./balancerv2 --numthreads=8 -rpc 'http://198.58.105.159:28545' -schema eth_bal_v2 -blockrcpts true &
	PROC_NUM=$!
	./delayed-kill.sh $ONE_HOUR $PROC_NUM
	wait $PROC_NUM
	sleep 60
done
