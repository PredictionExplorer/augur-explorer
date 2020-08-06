#!/bin/bash

echo Building ETL daemon ...
cd etl
go build ./

echo Building 0x Mesh order listener
cd dmesh
go build ./

cd ..
cd ..
echo Building Web Server ...
cd server
go build ./

echo Building Tools ...
cd ..
cd etl/tools
go build ./dai_balances.go
go build ./toprated.go
go build ./uniqueaddrs.go
go build ./augur_blocks.go
echo Suggested crontab configuration for production:
echo '*/5 * * * *	. $HOME/configs/etl-config.env; $HOME/augur-explorer/etl/tools/uniqueaddrs 5'
echo '*/10 * * * *	$HOME/augur-explorer/etl/tools/run-toprated.sh $HOME/augur-explorer/etl/tools $HOME/configs/etl-config.env'


echo Build complete , now execute: ./auto-run.sh


