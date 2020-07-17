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
./toprated
