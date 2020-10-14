#!/bin/bash
DBNAME=$1
if test -z "$1"
then
	echo usage: $0 '[dbname] [init_script.sql]'
	exit 1
fi
psql $DBNAME < drop-tokens.sql
psql $DBNAME < drop-augur.sql
cat tables_augur.sql tables_tokens.sql trigger-funcs.sql triggers.sql indices.sql indices_tokens.sql | psql $DBNAME
