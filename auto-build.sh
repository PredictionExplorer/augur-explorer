#!/bin/bash

GOLANGPATH=$HOME/go

echo Getting Golang dependencies for Augur Extractor
go get -v github.com/gin-gonic/gin
go get -v github.com/plaid/go-envvar/envvar
go get -v github.com/stretchr/testify/require
go get -v github.com/libp2p/go-libp2p-peerstore
go get -v github.com/jpillora/backoff
go get -v github.com/0xProject/0x-mesh/zeroex
go get -v github.com/ethereum/go-ethereum/ethclient
go get -v github.com/ethereum/go-ethereum/accounts
go get -v github.com/ethereum/go-ethereum/common
go get -v github.com/ethereum/go-ethereum/core
go get -v github.com/lib/pq
echo All dependencies were downloaded
echo Getting Augur-Extractor package
go get -v github.com/afterether/augur-extractor
cd $GOLANGPATH/src/github.com/afterether/augur-extractor/etl
go build ./
cd $GOLANGPATH/src/github.com/afterether/augur-extractor/etl/mesh
go build ./
cd $GOLANGPATH/src/github.com/afterether/augur-extractor/server
go build ./

echo Build complete , now execute: ./auto-run.sh


