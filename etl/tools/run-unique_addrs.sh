#!/bin/bash
if test -z "$1"
then
	echo "usage $0 [path] [configfile]"
	exit 1
fi
if test -z "$2"
then
	echo "usage $0 [path] [configfile]"
	exit 1
fi
cd $1
. $2
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
