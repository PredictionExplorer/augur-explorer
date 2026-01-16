// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IExchangeOrder is an auto generated low-level Go binding around an user-defined struct.
type IExchangeOrder struct {
	MakerAddress          common.Address
	TakerAddress          common.Address
	FeeRecipientAddress   common.Address
	SenderAddress         common.Address
	MakerAssetAmount      *big.Int
	TakerAssetAmount      *big.Int
	MakerFee              *big.Int
	TakerFee              *big.Int
	ExpirationTimeSeconds *big.Int
	Salt                  *big.Int
	MakerAssetData        []byte
	TakerAssetData        []byte
	MakerFeeAssetData     []byte
	TakerFeeAssetData     []byte
}

// IZeroXTradeAugurOrderData is an auto generated low-level Go binding around an user-defined struct.
type IZeroXTradeAugurOrderData struct {
	MarketAddress common.Address
	Price         *big.Int
	Outcome       uint8
	OrderType     uint8
}

// ZeroXMetaData contains all meta data concerning the ZeroX contract.
var ZeroXMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP1271_ORDER_WITH_HASH_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"askBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augur\",\"outputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augurTrading\",\"outputs\":[{\"internalType\":\"contractIAugurTrading\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances_\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"bidBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order[]\",\"name\":\"_orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxProtocolFeeDai\",\"type\":\"uint256\"}],\"name\":\"cancelOrders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"contractICash\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"cashAvailableForTransferFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_attoshares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"createZeroXOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_zeroXOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_orderHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_attoshares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"createZeroXOrderFor\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_zeroXOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_orderHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"creatorHasFundsForTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_assetData\",\"type\":\"bytes\"}],\"name\":\"decodeAssetData\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"_assetProxyId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenValues\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_assetData\",\"type\":\"bytes\"}],\"name\":\"decodeTradeAssetData\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"_assetProxyId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenValues\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"encodeAssetData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_assetData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_zeroXOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_orderHash\",\"type\":\"bytes32\"}],\"name\":\"encodeEIP1271OrderWithHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numOrders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"estimateProtocolFeeCostInCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethExchange\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Exchange\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"contractIExchange\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"contractIFillOrder\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"getTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"getTokenIdFromOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTransferFromAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_assetData\",\"type\":\"bytes\"}],\"name\":\"getZeroXTradeTokenData\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"_augur\",\"type\":\"address\"},{\"internalType\":\"contractIAugurTrading\",\"name\":\"_augurTrading\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderAmount\",\"type\":\"uint256\"}],\"name\":\"isOrderAmountValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"parseOrderData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"marketAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"}],\"internalType\":\"structIZeroXTrade.AugurOrderData\",\"name\":\"_data\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shareToken\",\"outputs\":[{\"internalType\":\"contractIShareToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0IsCash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestedFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxProtocolFeeDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTrades\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order[]\",\"name\":\"_orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"unpackTokenId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ZeroXABI is the input ABI used to generate the binding from.
// Deprecated: Use ZeroXMetaData.ABI instead.
var ZeroXABI = ZeroXMetaData.ABI

// ZeroX is an auto generated Go binding around an Ethereum contract.
type ZeroX struct {
	ZeroXCaller     // Read-only binding to the contract
	ZeroXTransactor // Write-only binding to the contract
	ZeroXFilterer   // Log filterer for contract events
}

// ZeroXCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroXSession struct {
	Contract     *ZeroX            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroXCallerSession struct {
	Contract *ZeroXCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZeroXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroXTransactorSession struct {
	Contract     *ZeroXTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroXRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroXRaw struct {
	Contract *ZeroX // Generic contract binding to access the raw methods on
}

// ZeroXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroXCallerRaw struct {
	Contract *ZeroXCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroXTransactorRaw struct {
	Contract *ZeroXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroX creates a new instance of ZeroX, bound to a specific deployed contract.
func NewZeroX(address common.Address, backend bind.ContractBackend) (*ZeroX, error) {
	contract, err := bindZeroX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroX{ZeroXCaller: ZeroXCaller{contract: contract}, ZeroXTransactor: ZeroXTransactor{contract: contract}, ZeroXFilterer: ZeroXFilterer{contract: contract}}, nil
}

// NewZeroXCaller creates a new read-only instance of ZeroX, bound to a specific deployed contract.
func NewZeroXCaller(address common.Address, caller bind.ContractCaller) (*ZeroXCaller, error) {
	contract, err := bindZeroX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroXCaller{contract: contract}, nil
}

// NewZeroXTransactor creates a new write-only instance of ZeroX, bound to a specific deployed contract.
func NewZeroXTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroXTransactor, error) {
	contract, err := bindZeroX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroXTransactor{contract: contract}, nil
}

// NewZeroXFilterer creates a new log filterer instance of ZeroX, bound to a specific deployed contract.
func NewZeroXFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroXFilterer, error) {
	contract, err := bindZeroX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroXFilterer{contract: contract}, nil
}

// bindZeroX binds a generic wrapper to an already deployed contract.
func bindZeroX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroXABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroX *ZeroXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroX.Contract.ZeroXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroX *ZeroXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroX.Contract.ZeroXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroX *ZeroXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroX.Contract.ZeroXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroX *ZeroXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroX *ZeroXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroX *ZeroXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroX.Contract.contract.Transact(opts, method, params...)
}

// EIP1271ORDERWITHHASHSELECTOR is a free data retrieval call binding the contract method 0xad5c22c8.
//
// Solidity: function EIP1271_ORDER_WITH_HASH_SELECTOR() view returns(bytes4)
func (_ZeroX *ZeroXCaller) EIP1271ORDERWITHHASHSELECTOR(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "EIP1271_ORDER_WITH_HASH_SELECTOR")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// EIP1271ORDERWITHHASHSELECTOR is a free data retrieval call binding the contract method 0xad5c22c8.
//
// Solidity: function EIP1271_ORDER_WITH_HASH_SELECTOR() view returns(bytes4)
func (_ZeroX *ZeroXSession) EIP1271ORDERWITHHASHSELECTOR() ([4]byte, error) {
	return _ZeroX.Contract.EIP1271ORDERWITHHASHSELECTOR(&_ZeroX.CallOpts)
}

// EIP1271ORDERWITHHASHSELECTOR is a free data retrieval call binding the contract method 0xad5c22c8.
//
// Solidity: function EIP1271_ORDER_WITH_HASH_SELECTOR() view returns(bytes4)
func (_ZeroX *ZeroXCallerSession) EIP1271ORDERWITHHASHSELECTOR() ([4]byte, error) {
	return _ZeroX.Contract.EIP1271ORDERWITHHASHSELECTOR(&_ZeroX.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_ZeroX *ZeroXCaller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "EIP712_DOMAIN_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_ZeroX *ZeroXSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _ZeroX.Contract.EIP712DOMAINHASH(&_ZeroX.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_ZeroX *ZeroXCallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _ZeroX.Contract.EIP712DOMAINHASH(&_ZeroX.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_ZeroX *ZeroXCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_ZeroX *ZeroXSession) WETH() (common.Address, error) {
	return _ZeroX.Contract.WETH(&_ZeroX.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_ZeroX *ZeroXCallerSession) WETH() (common.Address, error) {
	return _ZeroX.Contract.WETH(&_ZeroX.CallOpts)
}

// AskBalance is a free data retrieval call binding the contract method 0x5e03e6e4.
//
// Solidity: function askBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_ZeroX *ZeroXCaller) AskBalance(opts *bind.CallOpts, _owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "askBalance", _owner, _market, _outcome, _price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AskBalance is a free data retrieval call binding the contract method 0x5e03e6e4.
//
// Solidity: function askBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_ZeroX *ZeroXSession) AskBalance(_owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.AskBalance(&_ZeroX.CallOpts, _owner, _market, _outcome, _price)
}

// AskBalance is a free data retrieval call binding the contract method 0x5e03e6e4.
//
// Solidity: function askBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_ZeroX *ZeroXCallerSession) AskBalance(_owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.AskBalance(&_ZeroX.CallOpts, _owner, _market, _outcome, _price)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_ZeroX *ZeroXCaller) Augur(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "augur")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_ZeroX *ZeroXSession) Augur() (common.Address, error) {
	return _ZeroX.Contract.Augur(&_ZeroX.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_ZeroX *ZeroXCallerSession) Augur() (common.Address, error) {
	return _ZeroX.Contract.Augur(&_ZeroX.CallOpts)
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_ZeroX *ZeroXCaller) AugurTrading(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "augurTrading")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_ZeroX *ZeroXSession) AugurTrading() (common.Address, error) {
	return _ZeroX.Contract.AugurTrading(&_ZeroX.CallOpts)
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_ZeroX *ZeroXCallerSession) AugurTrading() (common.Address, error) {
	return _ZeroX.Contract.AugurTrading(&_ZeroX.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_ZeroX *ZeroXCaller) BalanceOf(opts *bind.CallOpts, owner common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "balanceOf", owner, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_ZeroX *ZeroXSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.BalanceOf(&_ZeroX.CallOpts, owner, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_ZeroX *ZeroXCallerSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.BalanceOf(&_ZeroX.CallOpts, owner, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances_)
func (_ZeroX *ZeroXCaller) BalanceOfBatch(opts *bind.CallOpts, owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "balanceOfBatch", owners, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances_)
func (_ZeroX *ZeroXSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ZeroX.Contract.BalanceOfBatch(&_ZeroX.CallOpts, owners, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances_)
func (_ZeroX *ZeroXCallerSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ZeroX.Contract.BalanceOfBatch(&_ZeroX.CallOpts, owners, ids)
}

// BidBalance is a free data retrieval call binding the contract method 0x2c52071c.
//
// Solidity: function bidBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_ZeroX *ZeroXCaller) BidBalance(opts *bind.CallOpts, _owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "bidBalance", _owner, _market, _outcome, _price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidBalance is a free data retrieval call binding the contract method 0x2c52071c.
//
// Solidity: function bidBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_ZeroX *ZeroXSession) BidBalance(_owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.BidBalance(&_ZeroX.CallOpts, _owner, _market, _outcome, _price)
}

// BidBalance is a free data retrieval call binding the contract method 0x2c52071c.
//
// Solidity: function bidBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_ZeroX *ZeroXCallerSession) BidBalance(_owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.BidBalance(&_ZeroX.CallOpts, _owner, _market, _outcome, _price)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_ZeroX *ZeroXCaller) Cash(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_ZeroX *ZeroXSession) Cash() (common.Address, error) {
	return _ZeroX.Contract.Cash(&_ZeroX.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_ZeroX *ZeroXCallerSession) Cash() (common.Address, error) {
	return _ZeroX.Contract.Cash(&_ZeroX.CallOpts)
}

// CashAvailableForTransferFrom is a free data retrieval call binding the contract method 0x1d4b44bb.
//
// Solidity: function cashAvailableForTransferFrom(address _owner, address _sender) view returns(uint256)
func (_ZeroX *ZeroXCaller) CashAvailableForTransferFrom(opts *bind.CallOpts, _owner common.Address, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "cashAvailableForTransferFrom", _owner, _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CashAvailableForTransferFrom is a free data retrieval call binding the contract method 0x1d4b44bb.
//
// Solidity: function cashAvailableForTransferFrom(address _owner, address _sender) view returns(uint256)
func (_ZeroX *ZeroXSession) CashAvailableForTransferFrom(_owner common.Address, _sender common.Address) (*big.Int, error) {
	return _ZeroX.Contract.CashAvailableForTransferFrom(&_ZeroX.CallOpts, _owner, _sender)
}

// CashAvailableForTransferFrom is a free data retrieval call binding the contract method 0x1d4b44bb.
//
// Solidity: function cashAvailableForTransferFrom(address _owner, address _sender) view returns(uint256)
func (_ZeroX *ZeroXCallerSession) CashAvailableForTransferFrom(_owner common.Address, _sender common.Address) (*big.Int, error) {
	return _ZeroX.Contract.CashAvailableForTransferFrom(&_ZeroX.CallOpts, _owner, _sender)
}

// CreateZeroXOrder is a free data retrieval call binding the contract method 0xf94515f0.
//
// Solidity: function createZeroXOrder(uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _zeroXOrder, bytes32 _orderHash)
func (_ZeroX *ZeroXCaller) CreateZeroXOrder(opts *bind.CallOpts, _type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "createZeroXOrder", _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)

	outstruct := new(struct {
		ZeroXOrder IExchangeOrder
		OrderHash  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ZeroXOrder = *abi.ConvertType(out[0], new(IExchangeOrder)).(*IExchangeOrder)
	outstruct.OrderHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// CreateZeroXOrder is a free data retrieval call binding the contract method 0xf94515f0.
//
// Solidity: function createZeroXOrder(uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _zeroXOrder, bytes32 _orderHash)
func (_ZeroX *ZeroXSession) CreateZeroXOrder(_type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	return _ZeroX.Contract.CreateZeroXOrder(&_ZeroX.CallOpts, _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)
}

// CreateZeroXOrder is a free data retrieval call binding the contract method 0xf94515f0.
//
// Solidity: function createZeroXOrder(uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _zeroXOrder, bytes32 _orderHash)
func (_ZeroX *ZeroXCallerSession) CreateZeroXOrder(_type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	return _ZeroX.Contract.CreateZeroXOrder(&_ZeroX.CallOpts, _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)
}

// CreateZeroXOrderFor is a free data retrieval call binding the contract method 0xe186955b.
//
// Solidity: function createZeroXOrderFor(address _maker, uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _zeroXOrder, bytes32 _orderHash)
func (_ZeroX *ZeroXCaller) CreateZeroXOrderFor(opts *bind.CallOpts, _maker common.Address, _type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "createZeroXOrderFor", _maker, _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)

	outstruct := new(struct {
		ZeroXOrder IExchangeOrder
		OrderHash  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ZeroXOrder = *abi.ConvertType(out[0], new(IExchangeOrder)).(*IExchangeOrder)
	outstruct.OrderHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// CreateZeroXOrderFor is a free data retrieval call binding the contract method 0xe186955b.
//
// Solidity: function createZeroXOrderFor(address _maker, uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _zeroXOrder, bytes32 _orderHash)
func (_ZeroX *ZeroXSession) CreateZeroXOrderFor(_maker common.Address, _type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	return _ZeroX.Contract.CreateZeroXOrderFor(&_ZeroX.CallOpts, _maker, _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)
}

// CreateZeroXOrderFor is a free data retrieval call binding the contract method 0xe186955b.
//
// Solidity: function createZeroXOrderFor(address _maker, uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _zeroXOrder, bytes32 _orderHash)
func (_ZeroX *ZeroXCallerSession) CreateZeroXOrderFor(_maker common.Address, _type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	return _ZeroX.Contract.CreateZeroXOrderFor(&_ZeroX.CallOpts, _maker, _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)
}

// CreatorHasFundsForTrade is a free data retrieval call binding the contract method 0xe264cff1.
//
// Solidity: function creatorHasFundsForTrade((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _order, uint256 _amount) view returns(bool)
func (_ZeroX *ZeroXCaller) CreatorHasFundsForTrade(opts *bind.CallOpts, _order IExchangeOrder, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "creatorHasFundsForTrade", _order, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreatorHasFundsForTrade is a free data retrieval call binding the contract method 0xe264cff1.
//
// Solidity: function creatorHasFundsForTrade((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _order, uint256 _amount) view returns(bool)
func (_ZeroX *ZeroXSession) CreatorHasFundsForTrade(_order IExchangeOrder, _amount *big.Int) (bool, error) {
	return _ZeroX.Contract.CreatorHasFundsForTrade(&_ZeroX.CallOpts, _order, _amount)
}

// CreatorHasFundsForTrade is a free data retrieval call binding the contract method 0xe264cff1.
//
// Solidity: function creatorHasFundsForTrade((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _order, uint256 _amount) view returns(bool)
func (_ZeroX *ZeroXCallerSession) CreatorHasFundsForTrade(_order IExchangeOrder, _amount *big.Int) (bool, error) {
	return _ZeroX.Contract.CreatorHasFundsForTrade(&_ZeroX.CallOpts, _order, _amount)
}

// DecodeAssetData is a free data retrieval call binding the contract method 0x7cb11125.
//
// Solidity: function decodeAssetData(bytes _assetData) view returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_ZeroX *ZeroXCaller) DecodeAssetData(opts *bind.CallOpts, _assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "decodeAssetData", _assetData)

	outstruct := new(struct {
		AssetProxyId [4]byte
		TokenAddress common.Address
		TokenIds     []*big.Int
		TokenValues  []*big.Int
		CallbackData []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AssetProxyId = *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	outstruct.TokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenIds = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.TokenValues = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	outstruct.CallbackData = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// DecodeAssetData is a free data retrieval call binding the contract method 0x7cb11125.
//
// Solidity: function decodeAssetData(bytes _assetData) view returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_ZeroX *ZeroXSession) DecodeAssetData(_assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	return _ZeroX.Contract.DecodeAssetData(&_ZeroX.CallOpts, _assetData)
}

// DecodeAssetData is a free data retrieval call binding the contract method 0x7cb11125.
//
// Solidity: function decodeAssetData(bytes _assetData) view returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_ZeroX *ZeroXCallerSession) DecodeAssetData(_assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	return _ZeroX.Contract.DecodeAssetData(&_ZeroX.CallOpts, _assetData)
}

// DecodeTradeAssetData is a free data retrieval call binding the contract method 0xb3ca8e25.
//
// Solidity: function decodeTradeAssetData(bytes _assetData) pure returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_ZeroX *ZeroXCaller) DecodeTradeAssetData(opts *bind.CallOpts, _assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "decodeTradeAssetData", _assetData)

	outstruct := new(struct {
		AssetProxyId [4]byte
		TokenAddress common.Address
		TokenIds     []*big.Int
		TokenValues  []*big.Int
		CallbackData []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AssetProxyId = *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	outstruct.TokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenIds = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.TokenValues = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	outstruct.CallbackData = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// DecodeTradeAssetData is a free data retrieval call binding the contract method 0xb3ca8e25.
//
// Solidity: function decodeTradeAssetData(bytes _assetData) pure returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_ZeroX *ZeroXSession) DecodeTradeAssetData(_assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	return _ZeroX.Contract.DecodeTradeAssetData(&_ZeroX.CallOpts, _assetData)
}

// DecodeTradeAssetData is a free data retrieval call binding the contract method 0xb3ca8e25.
//
// Solidity: function decodeTradeAssetData(bytes _assetData) pure returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_ZeroX *ZeroXCallerSession) DecodeTradeAssetData(_assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	return _ZeroX.Contract.DecodeTradeAssetData(&_ZeroX.CallOpts, _assetData)
}

// EncodeAssetData is a free data retrieval call binding the contract method 0x235f9e08.
//
// Solidity: function encodeAssetData(address _market, uint256 _price, uint8 _outcome, uint8 _type) view returns(bytes _assetData)
func (_ZeroX *ZeroXCaller) EncodeAssetData(opts *bind.CallOpts, _market common.Address, _price *big.Int, _outcome uint8, _type uint8) ([]byte, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "encodeAssetData", _market, _price, _outcome, _type)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeAssetData is a free data retrieval call binding the contract method 0x235f9e08.
//
// Solidity: function encodeAssetData(address _market, uint256 _price, uint8 _outcome, uint8 _type) view returns(bytes _assetData)
func (_ZeroX *ZeroXSession) EncodeAssetData(_market common.Address, _price *big.Int, _outcome uint8, _type uint8) ([]byte, error) {
	return _ZeroX.Contract.EncodeAssetData(&_ZeroX.CallOpts, _market, _price, _outcome, _type)
}

// EncodeAssetData is a free data retrieval call binding the contract method 0x235f9e08.
//
// Solidity: function encodeAssetData(address _market, uint256 _price, uint8 _outcome, uint8 _type) view returns(bytes _assetData)
func (_ZeroX *ZeroXCallerSession) EncodeAssetData(_market common.Address, _price *big.Int, _outcome uint8, _type uint8) ([]byte, error) {
	return _ZeroX.Contract.EncodeAssetData(&_ZeroX.CallOpts, _market, _price, _outcome, _type)
}

// EncodeEIP1271OrderWithHash is a free data retrieval call binding the contract method 0x4bf1ee87.
//
// Solidity: function encodeEIP1271OrderWithHash((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _zeroXOrder, bytes32 _orderHash) pure returns(bytes encoded)
func (_ZeroX *ZeroXCaller) EncodeEIP1271OrderWithHash(opts *bind.CallOpts, _zeroXOrder IExchangeOrder, _orderHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "encodeEIP1271OrderWithHash", _zeroXOrder, _orderHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeEIP1271OrderWithHash is a free data retrieval call binding the contract method 0x4bf1ee87.
//
// Solidity: function encodeEIP1271OrderWithHash((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _zeroXOrder, bytes32 _orderHash) pure returns(bytes encoded)
func (_ZeroX *ZeroXSession) EncodeEIP1271OrderWithHash(_zeroXOrder IExchangeOrder, _orderHash [32]byte) ([]byte, error) {
	return _ZeroX.Contract.EncodeEIP1271OrderWithHash(&_ZeroX.CallOpts, _zeroXOrder, _orderHash)
}

// EncodeEIP1271OrderWithHash is a free data retrieval call binding the contract method 0x4bf1ee87.
//
// Solidity: function encodeEIP1271OrderWithHash((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _zeroXOrder, bytes32 _orderHash) pure returns(bytes encoded)
func (_ZeroX *ZeroXCallerSession) EncodeEIP1271OrderWithHash(_zeroXOrder IExchangeOrder, _orderHash [32]byte) ([]byte, error) {
	return _ZeroX.Contract.EncodeEIP1271OrderWithHash(&_ZeroX.CallOpts, _zeroXOrder, _orderHash)
}

// EstimateProtocolFeeCostInCash is a free data retrieval call binding the contract method 0x3fc6f627.
//
// Solidity: function estimateProtocolFeeCostInCash(uint256 _numOrders, uint256 _gasPrice) view returns(uint256)
func (_ZeroX *ZeroXCaller) EstimateProtocolFeeCostInCash(opts *bind.CallOpts, _numOrders *big.Int, _gasPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "estimateProtocolFeeCostInCash", _numOrders, _gasPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateProtocolFeeCostInCash is a free data retrieval call binding the contract method 0x3fc6f627.
//
// Solidity: function estimateProtocolFeeCostInCash(uint256 _numOrders, uint256 _gasPrice) view returns(uint256)
func (_ZeroX *ZeroXSession) EstimateProtocolFeeCostInCash(_numOrders *big.Int, _gasPrice *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.EstimateProtocolFeeCostInCash(&_ZeroX.CallOpts, _numOrders, _gasPrice)
}

// EstimateProtocolFeeCostInCash is a free data retrieval call binding the contract method 0x3fc6f627.
//
// Solidity: function estimateProtocolFeeCostInCash(uint256 _numOrders, uint256 _gasPrice) view returns(uint256)
func (_ZeroX *ZeroXCallerSession) EstimateProtocolFeeCostInCash(_numOrders *big.Int, _gasPrice *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.EstimateProtocolFeeCostInCash(&_ZeroX.CallOpts, _numOrders, _gasPrice)
}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_ZeroX *ZeroXCaller) EthExchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "ethExchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_ZeroX *ZeroXSession) EthExchange() (common.Address, error) {
	return _ZeroX.Contract.EthExchange(&_ZeroX.CallOpts)
}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_ZeroX *ZeroXCallerSession) EthExchange() (common.Address, error) {
	return _ZeroX.Contract.EthExchange(&_ZeroX.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_ZeroX *ZeroXCaller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "exchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_ZeroX *ZeroXSession) Exchange() (common.Address, error) {
	return _ZeroX.Contract.Exchange(&_ZeroX.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_ZeroX *ZeroXCallerSession) Exchange() (common.Address, error) {
	return _ZeroX.Contract.Exchange(&_ZeroX.CallOpts)
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_ZeroX *ZeroXCaller) FillOrder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "fillOrder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_ZeroX *ZeroXSession) FillOrder() (common.Address, error) {
	return _ZeroX.Contract.FillOrder(&_ZeroX.CallOpts)
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_ZeroX *ZeroXCallerSession) FillOrder() (common.Address, error) {
	return _ZeroX.Contract.FillOrder(&_ZeroX.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_ZeroX *ZeroXCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_ZeroX *ZeroXSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.GetAmountIn(&_ZeroX.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_ZeroX *ZeroXCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.GetAmountIn(&_ZeroX.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_ZeroX *ZeroXCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "getInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_ZeroX *ZeroXSession) GetInitialized() (bool, error) {
	return _ZeroX.Contract.GetInitialized(&_ZeroX.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_ZeroX *ZeroXCallerSession) GetInitialized() (bool, error) {
	return _ZeroX.Contract.GetInitialized(&_ZeroX.CallOpts)
}

// GetTokenId is a free data retrieval call binding the contract method 0x1f981bde.
//
// Solidity: function getTokenId(address _market, uint256 _price, uint8 _outcome, uint8 _type) pure returns(uint256 _tokenId)
func (_ZeroX *ZeroXCaller) GetTokenId(opts *bind.CallOpts, _market common.Address, _price *big.Int, _outcome uint8, _type uint8) (*big.Int, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "getTokenId", _market, _price, _outcome, _type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenId is a free data retrieval call binding the contract method 0x1f981bde.
//
// Solidity: function getTokenId(address _market, uint256 _price, uint8 _outcome, uint8 _type) pure returns(uint256 _tokenId)
func (_ZeroX *ZeroXSession) GetTokenId(_market common.Address, _price *big.Int, _outcome uint8, _type uint8) (*big.Int, error) {
	return _ZeroX.Contract.GetTokenId(&_ZeroX.CallOpts, _market, _price, _outcome, _type)
}

// GetTokenId is a free data retrieval call binding the contract method 0x1f981bde.
//
// Solidity: function getTokenId(address _market, uint256 _price, uint8 _outcome, uint8 _type) pure returns(uint256 _tokenId)
func (_ZeroX *ZeroXCallerSession) GetTokenId(_market common.Address, _price *big.Int, _outcome uint8, _type uint8) (*big.Int, error) {
	return _ZeroX.Contract.GetTokenId(&_ZeroX.CallOpts, _market, _price, _outcome, _type)
}

// GetTokenIdFromOrder is a free data retrieval call binding the contract method 0xafc6d8ba.
//
// Solidity: function getTokenIdFromOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _order) view returns(uint256 _tokenId)
func (_ZeroX *ZeroXCaller) GetTokenIdFromOrder(opts *bind.CallOpts, _order IExchangeOrder) (*big.Int, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "getTokenIdFromOrder", _order)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenIdFromOrder is a free data retrieval call binding the contract method 0xafc6d8ba.
//
// Solidity: function getTokenIdFromOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _order) view returns(uint256 _tokenId)
func (_ZeroX *ZeroXSession) GetTokenIdFromOrder(_order IExchangeOrder) (*big.Int, error) {
	return _ZeroX.Contract.GetTokenIdFromOrder(&_ZeroX.CallOpts, _order)
}

// GetTokenIdFromOrder is a free data retrieval call binding the contract method 0xafc6d8ba.
//
// Solidity: function getTokenIdFromOrder((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _order) view returns(uint256 _tokenId)
func (_ZeroX *ZeroXCallerSession) GetTokenIdFromOrder(_order IExchangeOrder) (*big.Int, error) {
	return _ZeroX.Contract.GetTokenIdFromOrder(&_ZeroX.CallOpts, _order)
}

// GetTransferFromAllowed is a free data retrieval call binding the contract method 0x09c1df9f.
//
// Solidity: function getTransferFromAllowed() view returns(bool)
func (_ZeroX *ZeroXCaller) GetTransferFromAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "getTransferFromAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetTransferFromAllowed is a free data retrieval call binding the contract method 0x09c1df9f.
//
// Solidity: function getTransferFromAllowed() view returns(bool)
func (_ZeroX *ZeroXSession) GetTransferFromAllowed() (bool, error) {
	return _ZeroX.Contract.GetTransferFromAllowed(&_ZeroX.CallOpts)
}

// GetTransferFromAllowed is a free data retrieval call binding the contract method 0x09c1df9f.
//
// Solidity: function getTransferFromAllowed() view returns(bool)
func (_ZeroX *ZeroXCallerSession) GetTransferFromAllowed() (bool, error) {
	return _ZeroX.Contract.GetTransferFromAllowed(&_ZeroX.CallOpts)
}

// GetZeroXTradeTokenData is a free data retrieval call binding the contract method 0x00ee3542.
//
// Solidity: function getZeroXTradeTokenData(bytes _assetData) view returns(address _token, uint256 _tokenId)
func (_ZeroX *ZeroXCaller) GetZeroXTradeTokenData(opts *bind.CallOpts, _assetData []byte) (struct {
	Token   common.Address
	TokenId *big.Int
}, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "getZeroXTradeTokenData", _assetData)

	outstruct := new(struct {
		Token   common.Address
		TokenId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetZeroXTradeTokenData is a free data retrieval call binding the contract method 0x00ee3542.
//
// Solidity: function getZeroXTradeTokenData(bytes _assetData) view returns(address _token, uint256 _tokenId)
func (_ZeroX *ZeroXSession) GetZeroXTradeTokenData(_assetData []byte) (struct {
	Token   common.Address
	TokenId *big.Int
}, error) {
	return _ZeroX.Contract.GetZeroXTradeTokenData(&_ZeroX.CallOpts, _assetData)
}

// GetZeroXTradeTokenData is a free data retrieval call binding the contract method 0x00ee3542.
//
// Solidity: function getZeroXTradeTokenData(bytes _assetData) view returns(address _token, uint256 _tokenId)
func (_ZeroX *ZeroXCallerSession) GetZeroXTradeTokenData(_assetData []byte) (struct {
	Token   common.Address
	TokenId *big.Int
}, error) {
	return _ZeroX.Contract.GetZeroXTradeTokenData(&_ZeroX.CallOpts, _assetData)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ZeroX *ZeroXCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ZeroX *ZeroXSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ZeroX.Contract.IsApprovedForAll(&_ZeroX.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ZeroX *ZeroXCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ZeroX.Contract.IsApprovedForAll(&_ZeroX.CallOpts, owner, operator)
}

// IsOrderAmountValid is a free data retrieval call binding the contract method 0x9c6fbce5.
//
// Solidity: function isOrderAmountValid(address _market, uint256 _orderAmount) view returns(bool)
func (_ZeroX *ZeroXCaller) IsOrderAmountValid(opts *bind.CallOpts, _market common.Address, _orderAmount *big.Int) (bool, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "isOrderAmountValid", _market, _orderAmount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderAmountValid is a free data retrieval call binding the contract method 0x9c6fbce5.
//
// Solidity: function isOrderAmountValid(address _market, uint256 _orderAmount) view returns(bool)
func (_ZeroX *ZeroXSession) IsOrderAmountValid(_market common.Address, _orderAmount *big.Int) (bool, error) {
	return _ZeroX.Contract.IsOrderAmountValid(&_ZeroX.CallOpts, _market, _orderAmount)
}

// IsOrderAmountValid is a free data retrieval call binding the contract method 0x9c6fbce5.
//
// Solidity: function isOrderAmountValid(address _market, uint256 _orderAmount) view returns(bool)
func (_ZeroX *ZeroXCallerSession) IsOrderAmountValid(_market common.Address, _orderAmount *big.Int) (bool, error) {
	return _ZeroX.Contract.IsOrderAmountValid(&_ZeroX.CallOpts, _market, _orderAmount)
}

// ParseOrderData is a free data retrieval call binding the contract method 0x61a4760b.
//
// Solidity: function parseOrderData((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _order) view returns((address,uint256,uint8,uint8) _data)
func (_ZeroX *ZeroXCaller) ParseOrderData(opts *bind.CallOpts, _order IExchangeOrder) (IZeroXTradeAugurOrderData, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "parseOrderData", _order)

	if err != nil {
		return *new(IZeroXTradeAugurOrderData), err
	}

	out0 := *abi.ConvertType(out[0], new(IZeroXTradeAugurOrderData)).(*IZeroXTradeAugurOrderData)

	return out0, err

}

// ParseOrderData is a free data retrieval call binding the contract method 0x61a4760b.
//
// Solidity: function parseOrderData((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _order) view returns((address,uint256,uint8,uint8) _data)
func (_ZeroX *ZeroXSession) ParseOrderData(_order IExchangeOrder) (IZeroXTradeAugurOrderData, error) {
	return _ZeroX.Contract.ParseOrderData(&_ZeroX.CallOpts, _order)
}

// ParseOrderData is a free data retrieval call binding the contract method 0x61a4760b.
//
// Solidity: function parseOrderData((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes) _order) view returns((address,uint256,uint8,uint8) _data)
func (_ZeroX *ZeroXCallerSession) ParseOrderData(_order IExchangeOrder) (IZeroXTradeAugurOrderData, error) {
	return _ZeroX.Contract.ParseOrderData(&_ZeroX.CallOpts, _order)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_ZeroX *ZeroXCaller) ShareToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "shareToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_ZeroX *ZeroXSession) ShareToken() (common.Address, error) {
	return _ZeroX.Contract.ShareToken(&_ZeroX.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_ZeroX *ZeroXCallerSession) ShareToken() (common.Address, error) {
	return _ZeroX.Contract.ShareToken(&_ZeroX.CallOpts)
}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_ZeroX *ZeroXCaller) Token0IsCash(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "token0IsCash")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_ZeroX *ZeroXSession) Token0IsCash() (bool, error) {
	return _ZeroX.Contract.Token0IsCash(&_ZeroX.CallOpts)
}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_ZeroX *ZeroXCallerSession) Token0IsCash() (bool, error) {
	return _ZeroX.Contract.Token0IsCash(&_ZeroX.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ZeroX *ZeroXCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ZeroX *ZeroXSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.TotalSupply(&_ZeroX.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ZeroX *ZeroXCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _ZeroX.Contract.TotalSupply(&_ZeroX.CallOpts, id)
}

// UnpackTokenId is a free data retrieval call binding the contract method 0x26afd2e8.
//
// Solidity: function unpackTokenId(uint256 _tokenId) pure returns(address _market, uint256 _price, uint8 _outcome, uint8 _type)
/*func (_ZeroX *ZeroXCaller) UnpackTokenId(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Market  common.Address
	Price   *big.Int
	Outcome uint8
	Type    uint8
}, error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "unpackTokenId", _tokenId)

	outstruct := new(struct {
		Market  common.Address
		Price   *big.Int
		Outcome uint8
		Type    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Market = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Outcome = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Type = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}*/
func (_ZeroX *ZeroXCaller) UnpackTokenId(opts *bind.CallOpts, _tokenId *big.Int) (p.ZxMeshOrderSpec,error) {
	var out []interface{}
	err := _ZeroX.contract.Call(opts, &out, "unpackTokenId", _tokenId)

	outstruct := new(p.ZxMeshOrderSpec)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Market = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Outcome = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Type = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}
// UnpackTokenId is a free data retrieval call binding the contract method 0x26afd2e8.
//
// Solidity: function unpackTokenId(uint256 _tokenId) pure returns(address _market, uint256 _price, uint8 _outcome, uint8 _type)
func (_ZeroX *ZeroXSession) UnpackTokenId(_tokenId *big.Int) (struct {
	Market  common.Address
	Price   *big.Int
	Outcome uint8
	Type    uint8
}, error) {
	return _ZeroX.Contract.UnpackTokenId(&_ZeroX.CallOpts, _tokenId)
}

// UnpackTokenId is a free data retrieval call binding the contract method 0x26afd2e8.
//
// Solidity: function unpackTokenId(uint256 _tokenId) pure returns(address _market, uint256 _price, uint8 _outcome, uint8 _type)
func (_ZeroX *ZeroXCallerSession) UnpackTokenId(_tokenId *big.Int) (struct {
	Market  common.Address
	Price   *big.Int
	Outcome uint8
	Type    uint8
}, error) {
	return _ZeroX.Contract.UnpackTokenId(&_ZeroX.CallOpts, _tokenId)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x4ea96c30.
//
// Solidity: function cancelOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] _orders, bytes[] _signatures, uint256 _maxProtocolFeeDai) returns(bool)
func (_ZeroX *ZeroXTransactor) CancelOrders(opts *bind.TransactOpts, _orders []IExchangeOrder, _signatures [][]byte, _maxProtocolFeeDai *big.Int) (*types.Transaction, error) {
	return _ZeroX.contract.Transact(opts, "cancelOrders", _orders, _signatures, _maxProtocolFeeDai)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x4ea96c30.
//
// Solidity: function cancelOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] _orders, bytes[] _signatures, uint256 _maxProtocolFeeDai) returns(bool)
func (_ZeroX *ZeroXSession) CancelOrders(_orders []IExchangeOrder, _signatures [][]byte, _maxProtocolFeeDai *big.Int) (*types.Transaction, error) {
	return _ZeroX.Contract.CancelOrders(&_ZeroX.TransactOpts, _orders, _signatures, _maxProtocolFeeDai)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x4ea96c30.
//
// Solidity: function cancelOrders((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] _orders, bytes[] _signatures, uint256 _maxProtocolFeeDai) returns(bool)
func (_ZeroX *ZeroXTransactorSession) CancelOrders(_orders []IExchangeOrder, _signatures [][]byte, _maxProtocolFeeDai *big.Int) (*types.Transaction, error) {
	return _ZeroX.Contract.CancelOrders(&_ZeroX.TransactOpts, _orders, _signatures, _maxProtocolFeeDai)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) returns()
func (_ZeroX *ZeroXTransactor) Initialize(opts *bind.TransactOpts, _augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _ZeroX.contract.Transact(opts, "initialize", _augur, _augurTrading)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) returns()
func (_ZeroX *ZeroXSession) Initialize(_augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _ZeroX.Contract.Initialize(&_ZeroX.TransactOpts, _augur, _augurTrading)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) returns()
func (_ZeroX *ZeroXTransactorSession) Initialize(_augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _ZeroX.Contract.Initialize(&_ZeroX.TransactOpts, _augur, _augurTrading)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ZeroX *ZeroXTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ZeroX.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ZeroX *ZeroXSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ZeroX.Contract.SafeBatchTransferFrom(&_ZeroX.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ZeroX *ZeroXTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ZeroX.Contract.SafeBatchTransferFrom(&_ZeroX.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ZeroX *ZeroXTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ZeroX.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ZeroX *ZeroXSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ZeroX.Contract.SafeTransferFrom(&_ZeroX.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ZeroX *ZeroXTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ZeroX.Contract.SafeTransferFrom(&_ZeroX.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ZeroX *ZeroXTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ZeroX.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ZeroX *ZeroXSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ZeroX.Contract.SetApprovalForAll(&_ZeroX.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ZeroX *ZeroXTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ZeroX.Contract.SetApprovalForAll(&_ZeroX.TransactOpts, operator, approved)
}

// Trade is a paid mutator transaction binding the contract method 0x2f562016.
//
// Solidity: function trade(uint256 _requestedFillAmount, bytes32 _fingerprint, bytes32 _tradeGroupId, uint256 _maxProtocolFeeDai, uint256 _maxTrades, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] _orders, bytes[] _signatures) payable returns(uint256)
func (_ZeroX *ZeroXTransactor) Trade(opts *bind.TransactOpts, _requestedFillAmount *big.Int, _fingerprint [32]byte, _tradeGroupId [32]byte, _maxProtocolFeeDai *big.Int, _maxTrades *big.Int, _orders []IExchangeOrder, _signatures [][]byte) (*types.Transaction, error) {
	return _ZeroX.contract.Transact(opts, "trade", _requestedFillAmount, _fingerprint, _tradeGroupId, _maxProtocolFeeDai, _maxTrades, _orders, _signatures)
}

// Trade is a paid mutator transaction binding the contract method 0x2f562016.
//
// Solidity: function trade(uint256 _requestedFillAmount, bytes32 _fingerprint, bytes32 _tradeGroupId, uint256 _maxProtocolFeeDai, uint256 _maxTrades, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] _orders, bytes[] _signatures) payable returns(uint256)
func (_ZeroX *ZeroXSession) Trade(_requestedFillAmount *big.Int, _fingerprint [32]byte, _tradeGroupId [32]byte, _maxProtocolFeeDai *big.Int, _maxTrades *big.Int, _orders []IExchangeOrder, _signatures [][]byte) (*types.Transaction, error) {
	return _ZeroX.Contract.Trade(&_ZeroX.TransactOpts, _requestedFillAmount, _fingerprint, _tradeGroupId, _maxProtocolFeeDai, _maxTrades, _orders, _signatures)
}

// Trade is a paid mutator transaction binding the contract method 0x2f562016.
//
// Solidity: function trade(uint256 _requestedFillAmount, bytes32 _fingerprint, bytes32 _tradeGroupId, uint256 _maxProtocolFeeDai, uint256 _maxTrades, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes,bytes,bytes)[] _orders, bytes[] _signatures) payable returns(uint256)
func (_ZeroX *ZeroXTransactorSession) Trade(_requestedFillAmount *big.Int, _fingerprint [32]byte, _tradeGroupId [32]byte, _maxProtocolFeeDai *big.Int, _maxTrades *big.Int, _orders []IExchangeOrder, _signatures [][]byte) (*types.Transaction, error) {
	return _ZeroX.Contract.Trade(&_ZeroX.TransactOpts, _requestedFillAmount, _fingerprint, _tradeGroupId, _maxProtocolFeeDai, _maxTrades, _orders, _signatures)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZeroX *ZeroXTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ZeroX.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZeroX *ZeroXSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ZeroX.Contract.Fallback(&_ZeroX.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZeroX *ZeroXTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ZeroX.Contract.Fallback(&_ZeroX.TransactOpts, calldata)
}

// ZeroXApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ZeroX contract.
type ZeroXApprovalForAllIterator struct {
	Event *ZeroXApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZeroXApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroXApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZeroXApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZeroXApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroXApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroXApprovalForAll represents a ApprovalForAll event raised by the ZeroX contract.
type ZeroXApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ZeroX *ZeroXFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ZeroXApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ZeroX.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ZeroXApprovalForAllIterator{contract: _ZeroX.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ZeroX *ZeroXFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ZeroXApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ZeroX.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroXApprovalForAll)
				if err := _ZeroX.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ZeroX *ZeroXFilterer) ParseApprovalForAll(log types.Log) (*ZeroXApprovalForAll, error) {
	event := new(ZeroXApprovalForAll)
	if err := _ZeroX.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZeroXTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ZeroX contract.
type ZeroXTransferBatchIterator struct {
	Event *ZeroXTransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZeroXTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroXTransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZeroXTransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZeroXTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroXTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroXTransferBatch represents a TransferBatch event raised by the ZeroX contract.
type ZeroXTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ZeroX *ZeroXFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ZeroXTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZeroX.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZeroXTransferBatchIterator{contract: _ZeroX.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ZeroX *ZeroXFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ZeroXTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZeroX.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroXTransferBatch)
				if err := _ZeroX.contract.UnpackLog(event, "TransferBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ZeroX *ZeroXFilterer) ParseTransferBatch(log types.Log) (*ZeroXTransferBatch, error) {
	event := new(ZeroXTransferBatch)
	if err := _ZeroX.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZeroXTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ZeroX contract.
type ZeroXTransferSingleIterator struct {
	Event *ZeroXTransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZeroXTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroXTransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZeroXTransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZeroXTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroXTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroXTransferSingle represents a TransferSingle event raised by the ZeroX contract.
type ZeroXTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ZeroX *ZeroXFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ZeroXTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZeroX.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZeroXTransferSingleIterator{contract: _ZeroX.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ZeroX *ZeroXFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ZeroXTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZeroX.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroXTransferSingle)
				if err := _ZeroX.contract.UnpackLog(event, "TransferSingle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ZeroX *ZeroXFilterer) ParseTransferSingle(log types.Log) (*ZeroXTransferSingle, error) {
	event := new(ZeroXTransferSingle)
	if err := _ZeroX.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZeroXURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ZeroX contract.
type ZeroXURIIterator struct {
	Event *ZeroXURI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZeroXURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroXURI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZeroXURI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZeroXURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroXURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroXURI represents a URI event raised by the ZeroX contract.
type ZeroXURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ZeroX *ZeroXFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ZeroXURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ZeroX.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ZeroXURIIterator{contract: _ZeroX.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ZeroX *ZeroXFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ZeroXURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ZeroX.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroXURI)
				if err := _ZeroX.contract.UnpackLog(event, "URI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ZeroX *ZeroXFilterer) ParseURI(log types.Log) (*ZeroXURI, error) {
	event := new(ZeroXURI)
	if err := _ZeroX.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
