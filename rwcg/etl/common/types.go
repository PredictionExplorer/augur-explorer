// Package common provides shared ETL functionality for event fetching and chain management
package common

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
)

// ETLContext holds shared context for ETL operations
type ETLContext struct {
	Storage   *dbs.SQLStorage
	EthClient *ethclient.Client
	RpcClient *rpc.Client
	Info      *log.Logger
	Error     *log.Logger
}
