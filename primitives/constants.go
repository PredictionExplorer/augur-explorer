package primitives
import (
	"errors"
)
const (
	MAX_BLOCKS_CHAIN_SPLIT = 128
	OWNER_FIELD_OFFSET int = 2	// offset to the 'owner' field in WalletContract in EVM (contract storage)
	CATEGORICAL_MULTIPLIER int = 1000
//	SCALAR_MULTIPLIER int = 10
	ENS_NOT_PUBLIC			= "ENS Name is not public"
)
const (
	MktTypeYesNo		= iota
	MktTypeCategorical
	MktTypeScalar
)
const (
	COINTYPE_ETHEREUM   int = 60
	COINTYPE_BITCOIN		=  0
	COINTYPE_LITECOIN		=  2
	COINTYPE_DOGECOIN		=  3
	COINTYPE_MONACOIN		= 22
	COINTYPE_ETHEREUM_CLASSIC = 61
	COINTYPE_ROOTSTOCK		= 137
	COINTYPE_RIPPLE			= 144
	COINTYPE_BITCOIN_CASH	= 145
	COINTYPE_BINANCE		= 714
)
const (
	OOOpCodeNone= iota
	OOOpCodeCreated
	OOOpCodeFill
	OOOpCodeCancelledByUser
	OOOpCodeExpired
	OOOpCodeSyncProcess		// when no other reason exist, this one is used (this is a kind of a bugfix)
)
type MeshEvtCode int
const (
	MeshEvtGetOrders MeshEvtCode = iota
	MeshEvtInvalid
	MeshEvtAdded				// 2
	MeshEvtFilled				// 3
	MeshEvtFullyFilled			// 4
	MeshEvtCancelled			// 5
	MeshEvtExpired				// 6
	MeshEvtUnexpired			// 7
	MeshEvtBecameUnfunded		// 8
	MeshEvtFillabilityIncreased
	MeshEvtStoppedWatching
)
type SearchResultType int
const (
	SR_Unknown SearchResultType = iota
	SR_MarketOrders				// 1
	SR_Address					// 2
	SR_Hash						// 3
	SR_Transaction				// 4
	SR_Block					// 5
	SR_UserInfo					// 6
	SR_WalletContractInfo		// 7
	SR_AugurMarketInfo			// 8
	SR_AugurUniverseInfo		// 9
	SR_ShareTokenWrapper		// 10
	SR_BalancerPool				// 11
	SR_UniswapPair				// 12
	SR_TextSearchResults		// 13
)
type WhatsNewAugurCode int
const (
	WNA_6Hours WhatsNewAugurCode = iota
	WNA_12Hours					// 1
	WNA_1Day					// 2
	WNA_2Days					// 3
	WNA_3Days					// 4
	WNA_1Week					// 5
	WNA_2Weeks					// 6
)
const (
	RecTypeBalancer = iota
	RecTypeMint
	RecTypeBurn
	RecTypeERC20
)
var (
	ErrChainSplit error = errors.New("Chainsplit detected")
)

type OrderType uint8
const (
	OrderTypeBid		OrderType = 0
	OrderTypeAsk		OrderType = 1
)
type OrderAction uint8
const(
	OrderActionCreate	OrderAction = 0
	OrderActionCancel	OrderAction = 1
	OrderActionFill		OrderAction = 2
)
type TokenType uint8
const(
	ReputationToken		TokenType = 0
	DisputeCrowdsourcer TokenType = 1
	ParticipationToken	TokenType = 2
)
type MarketStatus uint8
const (
	MktStatusTraded		MarketStatus = 0
	MktStatusReporting	MarketStatus = 1
	MktStatusReported	MarketStatus = 2
	MktStatusDisputing	MarketStatus = 3
	MktStatusFinalized	MarketStatus = 4
	MktStatusFinInvalid	MarketStatus = 5
)
