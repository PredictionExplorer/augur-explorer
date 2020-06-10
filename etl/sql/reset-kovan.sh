#!/bin/bash
psql dev < drop-tables.sql
cat tables.sql trigger-funcs.sql triggers.sql indices.sql kovan_dev_init.sql | psql dev
