## Augur Data Explorer

Extracts data from Augur Prediction Marketplace (http://augur.net) and stores it in an SQL database.
Exposes all extracted data via WEB. Demo: [URL pending]

### System components

 1. ETL (Extract-Transform-Load) process to extract data from the Blockchain and convert to SQL Database
 2. Web Server exposing the data in the DB over JSON API or HTML templates
 3. 0x Mesh Listener that feeds open orders from 0x to the DB
 4. Tools that generate statistics and run other useful processes


### Features

 * Converts Augur Platform trade data to SQL Database starting from the block Augur Platform was released
 * Does not need Etheeum archival node (Full node is enough)
 * Builds full trade history for all trading accounts
 * Keeps track of DAI balances of all User accounts without requiring archival Ethreum node
 * Detects chain splits and reverts trading history towards the correct chain, keeping integrity of the data intact
 * Full synchronization of 0x Mesh orders with the DB
 * Block & Transaction info
 * Search by: Block Number, Transaction Hash, 0x Order Hash, User Address/Wallet Address, Market Address
 * Web features:
   * Markets list
   * Market Info
       * Market Depth Chart
       * Market Price History chart
       * Market Trade History
       * Market Report History
   * User Info
       * Open Positions
       * Open Orders (un-filled)
       * Closed Positions
       * Profit/Loss chart
       * Report list
       * User ranking by Volume,Profit & Trade frequency
       * Deposit/Withdrawal list
       * Claim funds indicator
   * Categories
       * Markets per Category
   * Statistics
       * Unique Addresses chart
       * Cash flow chart
       * Trade statistics per User
       * Trade statistics per Market
       * Trade statistics per User-Market-Outcome
   * ToDo features
       * Market commisions
       * Affiliate commissions

### Requirements / Dependencies

 * Golang v 1.14
 * Augur Platform
 * 0x Mesh
 * PostgreSQL

### Build

This script will build everything, the only thing you need the OS to have `go` command
	
	./auto-build.sh

### Starting Augur in Development mode with Augur Local TestNet config

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
