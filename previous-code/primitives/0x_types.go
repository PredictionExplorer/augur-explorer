package primitives
import (
	"fmt"
	"math/big"
	"time"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common/math"
	//gethsigner "github.com/ethereum/go-ethereum/signer/core"
	ctypes "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"golang.org/x/crypto/sha3"
)
type ZxMeshOrderSpec struct {
	Market			common.Address
	Price			*big.Int
	Outcome			uint8
	Type			uint8
}
type EFill struct {
	MakerAddress           common.Address
	FeeRecipientAddress    common.Address
	MakerAssetData         []byte
	TakerAssetData         []byte
	MakerFeeAssetData      []byte
	TakerFeeAssetData      []byte
	OrderHash              [32]byte
	TakerAddress           common.Address
	SenderAddress          common.Address
	MakerAssetFilledAmount *big.Int
	TakerAssetFilledAmount *big.Int
	MakerFeePaid           *big.Int
	TakerFeePaid           *big.Int
	ProtocolFeePaid        *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// OrderInfo represents an fillable order and how much it could be filled for.
type OrderInfo0x struct {
	OrderHash                common.Hash         `json:"orderHash"`
	SignedOrder              *SignedOrder `json:"signedOrder"`
	FillableTakerAssetAmount *big.Int            `json:"fillableTakerAssetAmount"`
}

type orderInfoJSON struct {
	OrderHash                string              `json:"orderHash"`
	SignedOrder              *SignedOrder `json:"signedOrder"`
	FillableTakerAssetAmount string              `json:"fillableTakerAssetAmount"`
}

// Order represents an unsigned 0x order
type Order0x struct {
	ChainID               *big.Int       `json:"chainId"`
	ExchangeAddress       common.Address `json:"exchangeAddress"`
	MakerAddress          common.Address `json:"makerAddress"`
	MakerAssetData        []byte         `json:"makerAssetData"`
	MakerFeeAssetData     []byte         `json:"makerFeeAssetData"`
	MakerAssetAmount      *big.Int       `json:"makerAssetAmount"`
	MakerFee              *big.Int       `json:"makerFee"`
	TakerAddress          common.Address `json:"takerAddress"`
	TakerAssetData        []byte         `json:"takerAssetData"`
	TakerFeeAssetData     []byte         `json:"takerFeeAssetData"`
	TakerAssetAmount      *big.Int       `json:"takerAssetAmount"`
	TakerFee              *big.Int       `json:"takerFee"`
	SenderAddress         common.Address `json:"senderAddress"`
	FeeRecipientAddress   common.Address `json:"feeRecipientAddress"`
	ExpirationTimeSeconds *big.Int       `json:"expirationTimeSeconds"`
	Salt                  *big.Int       `json:"salt"`

	// Cache hash for performance
	hash *common.Hash
}

// SignedOrder represents a signed 0x order
type SignedOrder struct {
	Order0x
	Signature []byte `json:"signature"`
}

// SignatureType represents the type of 0x signature encountered
type SignatureType uint8

// SignatureType values
const (
	IllegalSignature SignatureType = iota
	InvalidSignature
	EIP712Signature
	EthSignSignature
	WalletSignature
	ValidatorSignature
	PreSignedSignature
	EIP1271WalletSignature
	NSignatureTypesSignature
)

// OrderStatus represents the status of an order as returned from the 0x smart contracts
// as part of OrderInfo
type OrderStatus uint8

// OrderStatus values
const (
	OSInvalid OrderStatus = iota
	OSInvalidMakerAssetAmount
	OSInvalidTakerAssetAmount
	OSFillable
	OSExpired
	OSFullyFilled
	OSCancelled
	OSSignatureInvalid
	OSInvalidMakerAssetData
	OSInvalidTakerAssetData
)

// ContractEvent is an event emitted by a smart contract
type ContractEvent struct {
	BlockHash  common.Hash
	TxHash     common.Hash
	TxIndex    uint
	LogIndex   uint
	IsRemoved  bool
	Address    common.Address
	Kind       string
	Parameters interface{}
}

type contractEventJSON struct {
	BlockHash  common.Hash
	TxHash     common.Hash
	TxIndex    uint
	LogIndex   uint
	IsRemoved  bool
	Address    common.Address
	Kind       string
	Parameters json.RawMessage
}

// OrderEvent is the order event emitted by the Mesh GraphQL API or the SubScribeToOrderEvents
// method in core.
type OrderEvent struct {
	// Timestamp is an order event timestamp that can be used for bookkeeping purposes.
	// If the OrderEvent represents a Mesh-specific event (e.g., ADDED, STOPPED_WATCHING),
	// the timestamp is when the event was generated. If the event was generated after
	// re-validating an order at the latest block height (e.g., FILLED, UNFUNDED, CANCELED),
	// then it is set to the latest block timestamp at which the order was re-validated.
	Timestamp time.Time `json:"timestamp"`
	// OrderHash is the EIP712 hash of the 0x order
	OrderHash common.Hash `json:"orderHash"`
	// SignedOrder is the signed 0x order struct
	SignedOrder *SignedOrder `json:"signedOrder"`
	// SignedOrder is the signed 0x order struct
	SignedOrderV4 *SignedOrderV4 `json:"signedOrderV4"`
	// EndState is the end state of this order at the time this event was generated
	EndState OrderEventEndState `json:"endState"`
	// FillableTakerAssetAmount is the amount for which this order is still fillable
	FillableTakerAssetAmount *big.Int `json:"fillableTakerAssetAmount"`
	// ContractEvents contains all the contract events that triggered this orders re-evaluation.
	// They did not all necessarily cause the orders state change itself, only it's re-evaluation.
	// Since it's state _did_ change, at least one of them did cause the actual state change.
	ContractEvents []*ContractEvent `json:"contractEvents"`
}

type orderEventJSON struct {
	Timestamp                time.Time            `json:"timestamp"`
	OrderHash                string               `json:"orderHash"`
	SignedOrder              *SignedOrder         `json:"signedOrder"`
	SignedOrderV4            *SignedOrderV4       `json:"signedOrderV4"`
	EndState                 string               `json:"endState"`
	FillableTakerAssetAmount string               `json:"fillableTakerAssetAmount"`
	ContractEvents           []*contractEventJSON `json:"contractEvents"`
}

// OrderV4 represents an unsigned 0x v4 limit order
// V4 Protocol also has RFQ orders, these are
// See <https://0xprotocol.readthedocs.io/en/latest/basics/orders.html#limit-orders>
type OrderV4 struct {
	// Domain information
	// TODO: These are constant within a chain context (mainnet/testnet/etc)
	// probably best to keep them out of the order struct, but this is how V3
	// does it.
	ChainID           *big.Int       `json:"chainId"`
	VerifyingContract common.Address `json:"verifyingContract"`

	// Limit order values
	MakerToken          common.Address `json:"makerToken"`
	TakerToken          common.Address `json:"takerToken"`
	MakerAmount         *big.Int       `json:"makerAmount"`         // uint128
	TakerAmount         *big.Int       `json:"takerAmount"`         // uint128
	TakerTokenFeeAmount *big.Int       `json:"takerTokenFeeAmount"` // uint128
	Maker               common.Address `json:"makerAddress"`
	Taker               common.Address `json:"takerAddress"`
	Sender              common.Address `json:"sender"`
	FeeRecipient        common.Address `json:"feeRecipient"`
	Pool                Bytes32        `json:"pool"`   // bytes32
	Expiry              *big.Int       `json:"expiry"` // uint64
	Salt                *big.Int       `json:"salt"`   // uint256

	// Cache hash for performance
	hash *common.Hash
}

type SignatureFieldV4 struct {
	SignatureType SignatureTypeV4 `json:"signatureType"`
	V             uint8           `json:"v"`
	R             Bytes32         `json:"r"`
	S             Bytes32         `json:"s"`
}

// SignatureTypeV4 represents the type of 0x signature encountered
type SignatureTypeV4 uint8

// SignedOrderV4 represents a signed 0x order
// See <https://0xprotocol.readthedocs.io/en/latest/basics/orders.html#how-to-sign>
type SignedOrderV4 struct {
	OrderV4   `json:"order"`
	Signature SignatureFieldV4 `json:"signature"`
}

// SignedOrderJSONV4 is an unmodified JSON representation of a SignedOrder
type SignedOrderJSONV4 struct {
	ChainID             int64  `json:"chainId"`
	VerifyingContract   string `json:"verifyingContract"`
	MakerToken          string `json:"makerToken"`
	TakerToken          string `json:"takerToken"`
	MakerAmount         string `json:"makerAmount"`
	TakerAmount         string `json:"takerAmount"`
	TakerTokenFeeAmount string `json:"takerTokenFeeAmount"`
	Maker               string `json:"maker"`
	Taker               string `json:"taker"`
	Sender              string `json:"sender"`
	FeeRecipient        string `json:"feeRecipient"`
	Pool                string `json:"pool"`
	Expiry              string `json:"expiry"`
	Salt                string `json:"salt"`
	SignatureType       string `json:"signatureType"`
	SignatureR          string `json:"signatureR"`
	SignatureV          string `json:"signatureV"`
	SignatureS          string `json:"signatureS"`
}

// SignatureType values
const (
	IllegalSignatureV4 SignatureTypeV4 = iota
	InvalidSignatureV4
	EIP712SignatureV4
	EthSignSignatureV4
)

// OrderStatusV4 represents the status of an order as returned from the 0x smart contracts
// as part of OrderInfo for v4 orders
type OrderStatusV4 uint8

// Order status values
// See <https://0xprotocol.readthedocs.io/en/latest/basics/functions.html?highlight=orderstatus#getlimitorderinfo>
// See <https://github.com/0xProject/protocol/blob/edda1edc507fbfceb6dcb02ef212ee4bdcb123a6/contracts/zero-ex/contracts/src/features/libs/LibNativeOrder.sol#L28>
const (
	OS4Invalid OrderStatusV4 = iota
	OS4Fillable
	OS4Filled
	OS4Cancelled
	OS4Expired
	OS4InvalidSignature // 0xMesh extension
)
// OrderEventEndState enumerates all the possible order event types. An OrderEventEndState describes the
// end state of a 0x order after revalidation
type OrderEventEndState string

// OrderEventEndState values
const (
	// ESInvalid is an event that is never emitted. It is here to discern between a declared but uninitialized OrderEventEndState
	ESInvalid = OrderEventEndState("INVALID")
	// ESOrderAdded means an order was successfully added to the Mesh node
	ESOrderAdded = OrderEventEndState("ADDED")
	// ESOrderFilled means an order was filled for a partial amount
	ESOrderFilled = OrderEventEndState("FILLED")
	// ESOrderFullyFilled means an order was fully filled such that it's remaining fillableTakerAssetAmount is 0
	ESOrderFullyFilled = OrderEventEndState("FULLY_FILLED")
	// ESOrderCancelled means an order was cancelled on-chain
	ESOrderCancelled = OrderEventEndState("CANCELLED")
	// ESOrderExpired means an order expired according to the latest block timestamp
	ESOrderExpired = OrderEventEndState("EXPIRED")
	// ESOrderUnexpired means an order is no longer expired. This can happen if a block re-org causes the latest
	// block timestamp to decline below the order's expirationTimestamp (rare and usually short-lived)
	ESOrderUnexpired = OrderEventEndState("UNEXPIRED")
	// ESOrderBecameUnfunded means an order has become unfunded. This happens if the maker transfers the balance /
	// changes their allowance backing an order
	ESOrderBecameUnfunded = OrderEventEndState("UNFUNDED")
	// ESOrderFillabilityIncreased means the fillability of an order has increased. Fillability for an order can
	// increase if a previously processed fill event gets reverted, or if a maker tops up their balance/allowance
	// backing an order
	ESOrderFillabilityIncreased = OrderEventEndState("FILLABILITY_INCREASED")
	// ESStoppedWatching means an order is potentially still valid but was removed for a different reason (e.g.
	// the database is full or the peer that sent the order was misbehaving). The order will no longer be watched
	// and no further events for this order will be emitted. In some cases, the order may be re-added in the
	// future.
	ESStoppedWatching = OrderEventEndState("STOPPED_WATCHING")
)

// Bytes32 represents the Solidity `bytes32` type.
// It is mostly copied from go-ethereum's Hash type.
// See <https://github.com/ethereum/go-ethereum/blob/053ed9cc847647a9b3ef707d0efe7104c4ab2a4c/common/types.go#L47>
type Bytes32 [32]byte



// ComputeOrderHash computes a 0x order hash
func (o *Order0x) ComputeOrderHash() (common.Hash, error) {
	if o.hash != nil {
		return *o.hash, nil
	}

	chainID := math.NewHexOrDecimal256(o.ChainID.Int64())
	var domain = ctypes.TypedDataDomain{
		Name:              "0x Protocol",
		Version:           "3.0.0",
		ChainId:           chainID,
		VerifyingContract: o.ExchangeAddress.Hex(),
	}

	var message = map[string]interface{}{
		"makerAddress":          o.MakerAddress.Hex(),
		"takerAddress":          o.TakerAddress.Hex(),
		"senderAddress":         o.SenderAddress.Hex(),
		"feeRecipientAddress":   o.FeeRecipientAddress.Hex(),
		"makerAssetData":        o.MakerAssetData,
		"makerFeeAssetData":     o.MakerFeeAssetData,
		"takerAssetData":        o.TakerAssetData,
		"takerFeeAssetData":     o.TakerFeeAssetData,
		"salt":                  o.Salt.String(),
		"makerFee":              o.MakerFee.String(),
		"takerFee":              o.TakerFee.String(),
		"makerAssetAmount":      o.MakerAssetAmount.String(),
		"takerAssetAmount":      o.TakerAssetAmount.String(),
		"expirationTimeSeconds": o.ExpirationTimeSeconds.String(),
	}

	var typedData = ctypes.TypedData{
		Types:       eip712OrderTypes,
		PrimaryType: "Order",
		Domain:      domain,
		Message:     message,
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return common.Hash{}, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return common.Hash{}, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hashBytes := keccak256(rawData)
	hash := common.BytesToHash(hashBytes)
	o.hash = &hash
	return hash, nil
}
// keccak256 calculates and returns the Keccak256 hash of the input data.
func keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		_, _ = d.Write(b)
	}
	return d.Sum(nil)
}
var eip712OrderTypes = ctypes.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},
	"Order": {
		{
			Name: "makerAddress",
			Type: "address",
		},
		{
			Name: "takerAddress",
			Type: "address",
		},
		{
			Name: "feeRecipientAddress",
			Type: "address",
		},
		{
			Name: "senderAddress",
			Type: "address",
		},
		{
			Name: "makerAssetAmount",
			Type: "uint256",
		},
		{
			Name: "takerAssetAmount",
			Type: "uint256",
		},
		{
			Name: "makerFee",
			Type: "uint256",
		},
		{
			Name: "takerFee",
			Type: "uint256",
		},
		{
			Name: "expirationTimeSeconds",
			Type: "uint256",
		},
		{
			Name: "salt",
			Type: "uint256",
		},
		{
			Name: "makerAssetData",
			Type: "bytes",
		},
		{
			Name: "takerAssetData",
			Type: "bytes",
		},
		{
			Name: "makerFeeAssetData",
			Type: "bytes",
		},
		{
			Name: "takerFeeAssetData",
			Type: "bytes",
		},
	},
}

