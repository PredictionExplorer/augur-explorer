#!/bin/bash
DBNAME=$1
if test -z "$1"
then
	echo usage: $0 '[dbname] [init_script.sql]'
	exit 1
fi
INIT_SCRIPT=$2
if test -z "$2"
then
	echo usage: $0 '[dbname] [init_script.sql]'
	exit 1
fi
psql $DBNAME < drop-tables.sql
cat tables.sql trigger-funcs.sql triggers.sql indices.sql $INIT_SCRIPT | psql $DBNAME
