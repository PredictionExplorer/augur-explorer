#!/bin/bash

echo Building ETL daemon ...
cd etl
go build cmd/tokens/tokens.go
go build cmd/augur/augur.go cmd/augur/main.go cmd/augur/rpc.go cmd/augur/sequencer.go
go build cmd/layer1/layer1.go cmd/layer1/augur.go cmd/layer1/rpc.go cmd/layer1/blockchain.go
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
go build ./gas_usage.go
go build ./check_owner.go
go build ./check_wallet.go
echo Building Test scripts ...
cd ../tests
go build ./verif_dai_balances.go
go build ./verif_cash_flow.go
echo Suggested crontab configuration for production:
echo '*/5 * * * *	. $HOME/configs/etl-config.env; $HOME/augur-explorer/etl/tools/uniqueaddrs 5'
echo '*/10 * * * *	$HOME/augur-explorer/etl/tools/run-toprated.sh $HOME/augur-explorer/etl/tools $HOME/configs/etl-config.env'


echo Build complete , now execute: ./auto-run.sh


