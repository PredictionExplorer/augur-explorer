#!/bin/bash
HOST="http://198.58.105.159:38546"
TXHASH=$1
if test -z "$TXHASH"
then
	echo usage: $0 '[tx_hash]'
	exit 1
fi
while true
do
	curl $HOST\
       -X POST\
       --header 'Content-type: application/json'\
       --data '{"jsonrpc":"2.0", "method":"debug_traceTransaction","params":["'"$TXHASH"'", {}], "id":1}'
	sleep 10
done
