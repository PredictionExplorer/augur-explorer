#!/bin/bash

echo Building ETL daemon ...
cd etl
go build ./

echo Building 0x Mesh order listener
cd mesh
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
echo Build complete , now execute: ./auto-run.sh


