package layer1

import (
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)

type ETL_Manager interface {
	Process_transaction(tx *AugurTx,rcpt *types.Receipt)
}
type ETL_Layer1 struct {
	Storage						*SQLStorage
	UseBlockReceiptsCall		bool
	NoRollbackBlocks			bool
	UpdateLastBlock				bool
	NoChainSplitCheck			bool
	SingleBlockNum				int64
	NumThreads					int64
	SchemaName					string
	RPC_Url						string
	AppName						string
	Manager						ETL_Manager
	Error						*log.Logger
	Info						*log.Logger
}
