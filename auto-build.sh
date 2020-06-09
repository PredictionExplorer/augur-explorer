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

echo Build complete , now execute: ./auto-run.sh


