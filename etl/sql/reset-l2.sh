#!/bin/bash
DBNAME=$1
if test -z "$1"
then
	echo usage: $0 '[dbname] [init_script.sql]'
	exit 1
fi
psql $DBNAME < drop-tokens.sql
psql $DBNAME < drop-augur.sql
psql $DBNAME < drop-balancer.sql
cat \
	tables_augur.sql \
	tables_tokens.sql \
	tables_balancer.sql \
	trigger-funcs.sql \
	triggers-augur.sql \
	triggers-balancer.sql \
	triggers-tokens.sql \
	indices_augur.sql \
	indices_tokens.sql \
| psql $DBNAME
