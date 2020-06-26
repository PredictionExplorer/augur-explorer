#!/bin/bash
psql kovan < drop-tables.sql
cat tables.sql trigger-funcs.sql triggers.sql indices.sql kovan_dev_init.sql | psql kovan
