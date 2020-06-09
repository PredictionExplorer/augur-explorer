## Augur Data Extractor

Extracts data from Augur Prediction Marketplace (http://augur.net) and stores it in an SQL database.

### Requirements

 * Golang v 1.14
 * Augur platform with local test net setup
 * PostgreSQL

### Build

This script will build everything, the only thing you need to have os `go` command
	
	./auto-build.sh

### Starting Augur

First, start Augur. Instructions are located here: https://github.com/AugurProject/augur

But just in case, here is the command list to run:

	docker kill $(docker ps -a -q);
	docker system prune -af
	yarn clean
	yarn
	yarn build
	yarn docker:all

on another terminal:

	yarn workspace @augurproject/ui dev


### Running

There are 3 executables to run. This script will run all of them:

	./auto-run.sh

### Database initialization

Create user on Ubuntu

	useradd -m aedev
	passwd aedev	# we are giong to set password to 123

Enter Postgres as superuser

	su - postgres
	psql

	postgres-#  CREATE ROLE aedev WITH LOGIN CREATEDB ENCRYPTED PASSWORD '123';
	\q

Enter Postgres as development user

	su - aedev
	createdb dev

Init DB

	cd etl/sql
	./reset-db.sh

### Database Schema and documentation

	cat etl/sql/tables.sql
