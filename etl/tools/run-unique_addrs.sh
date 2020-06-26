#!/bin/bash
cd /home/kovan/dev/augur-extractor/etl/tools
. ../config/dev-kovan-local.env
while true
do
	./uniqueaddrs
	RESULT=$?
	if test $RESULT -eq 1 ; then
		exit 1
	fi
	if test $RESULT -eq 2 ; then
		echo "no records to process, day isn't over"
		exit 0
	fi
done
