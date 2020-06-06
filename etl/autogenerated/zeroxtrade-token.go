// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP1271_ORDER_WITH_HASH_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"askBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augur\",\"outputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augurTrading\",\"outputs\":[{\"internalType\":\"contractIAugurTrading\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances_\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"bidBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order[]\",\"name\":\"_orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxProtocolFeeDai\",\"type\":\"uint256\"}],\"name\":\"cancelOrders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"contractICash\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"cashAvailableForTransferFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_attoshares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"createZeroXOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_zeroXOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_orderHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_attoshares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"createZeroXOrderFor\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_zeroXOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_orderHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"creatorHasFundsForTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_assetData\",\"type\":\"bytes\"}],\"name\":\"decodeAssetData\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"_assetProxyId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenValues\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_assetData\",\"type\":\"bytes\"}],\"name\":\"decodeTradeAssetData\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"_assetProxyId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenValues\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"encodeAssetData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_assetData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_zeroXOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_orderHash\",\"type\":\"bytes32\"}],\"name\":\"encodeEIP1271OrderWithHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numOrders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"estimateProtocolFeeCostInCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethExchange\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Exchange\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"contractIExchange\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"contractIFillOrder\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"getTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"getTokenIdFromOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTransferFromAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_assetData\",\"type\":\"bytes\"}],\"name\":\"getZeroXTradeTokenData\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"_augur\",\"type\":\"address\"},{\"internalType\":\"contractIAugurTrading\",\"name\":\"_augurTrading\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderAmount\",\"type\":\"uint256\"}],\"name\":\"isOrderAmountValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"parseOrderData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"marketAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"}],\"internalType\":\"structIZeroXTrade.AugurOrderData\",\"name\":\"_data\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shareToken\",\"outputs\":[{\"internalType\":\"contractIShareToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0IsCash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestedFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxProtocolFeeDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTrades\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIExchange.Order[]\",\"name\":\"_orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"unpackTokenId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_outcome\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// EIP1271ORDERWITHHASHSELECTOR is a free data retrieval call binding the contract method 0xad5c22c8.
//
// Solidity: function EIP1271_ORDER_WITH_HASH_SELECTOR() view returns(bytes4)
func (_Token *TokenCaller) EIP1271ORDERWITHHASHSELECTOR(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "EIP1271_ORDER_WITH_HASH_SELECTOR")
	return *ret0, err
}

// EIP1271ORDERWITHHASHSELECTOR is a free data retrieval call binding the contract method 0xad5c22c8.
//
// Solidity: function EIP1271_ORDER_WITH_HASH_SELECTOR() view returns(bytes4)
func (_Token *TokenSession) EIP1271ORDERWITHHASHSELECTOR() ([4]byte, error) {
	return _Token.Contract.EIP1271ORDERWITHHASHSELECTOR(&_Token.CallOpts)
}

// EIP1271ORDERWITHHASHSELECTOR is a free data retrieval call binding the contract method 0xad5c22c8.
//
// Solidity: function EIP1271_ORDER_WITH_HASH_SELECTOR() view returns(bytes4)
func (_Token *TokenCallerSession) EIP1271ORDERWITHHASHSELECTOR() ([4]byte, error) {
	return _Token.Contract.EIP1271ORDERWITHHASHSELECTOR(&_Token.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_Token *TokenCaller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "EIP712_DOMAIN_HASH")
	return *ret0, err
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_Token *TokenSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Token.Contract.EIP712DOMAINHASH(&_Token.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_Token *TokenCallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Token.Contract.EIP712DOMAINHASH(&_Token.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Token *TokenCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "WETH")
	return *ret0, err
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Token *TokenSession) WETH() (common.Address, error) {
	return _Token.Contract.WETH(&_Token.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Token *TokenCallerSession) WETH() (common.Address, error) {
	return _Token.Contract.WETH(&_Token.CallOpts)
}

// AskBalance is a free data retrieval call binding the contract method 0x5e03e6e4.
//
// Solidity: function askBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_Token *TokenCaller) AskBalance(opts *bind.CallOpts, _owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "askBalance", _owner, _market, _outcome, _price)
	return *ret0, err
}

// AskBalance is a free data retrieval call binding the contract method 0x5e03e6e4.
//
// Solidity: function askBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_Token *TokenSession) AskBalance(_owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	return _Token.Contract.AskBalance(&_Token.CallOpts, _owner, _market, _outcome, _price)
}

// AskBalance is a free data retrieval call binding the contract method 0x5e03e6e4.
//
// Solidity: function askBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_Token *TokenCallerSession) AskBalance(_owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	return _Token.Contract.AskBalance(&_Token.CallOpts, _owner, _market, _outcome, _price)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Token *TokenCaller) Augur(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "augur")
	return *ret0, err
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Token *TokenSession) Augur() (common.Address, error) {
	return _Token.Contract.Augur(&_Token.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Token *TokenCallerSession) Augur() (common.Address, error) {
	return _Token.Contract.Augur(&_Token.CallOpts)
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_Token *TokenCaller) AugurTrading(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "augurTrading")
	return *ret0, err
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_Token *TokenSession) AugurTrading() (common.Address, error) {
	return _Token.Contract.AugurTrading(&_Token.CallOpts)
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_Token *TokenCallerSession) AugurTrading() (common.Address, error) {
	return _Token.Contract.AugurTrading(&_Token.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", owner, id)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_Token *TokenSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, owner, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, owner, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances_)
func (_Token *TokenCaller) BalanceOfBatch(opts *bind.CallOpts, owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOfBatch", owners, ids)
	return *ret0, err
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances_)
func (_Token *TokenSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Token.Contract.BalanceOfBatch(&_Token.CallOpts, owners, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances_)
func (_Token *TokenCallerSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Token.Contract.BalanceOfBatch(&_Token.CallOpts, owners, ids)
}

// BidBalance is a free data retrieval call binding the contract method 0x2c52071c.
//
// Solidity: function bidBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_Token *TokenCaller) BidBalance(opts *bind.CallOpts, _owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "bidBalance", _owner, _market, _outcome, _price)
	return *ret0, err
}

// BidBalance is a free data retrieval call binding the contract method 0x2c52071c.
//
// Solidity: function bidBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_Token *TokenSession) BidBalance(_owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	return _Token.Contract.BidBalance(&_Token.CallOpts, _owner, _market, _outcome, _price)
}

// BidBalance is a free data retrieval call binding the contract method 0x2c52071c.
//
// Solidity: function bidBalance(address _owner, address _market, uint8 _outcome, uint256 _price) view returns(uint256)
func (_Token *TokenCallerSession) BidBalance(_owner common.Address, _market common.Address, _outcome uint8, _price *big.Int) (*big.Int, error) {
	return _Token.Contract.BidBalance(&_Token.CallOpts, _owner, _market, _outcome, _price)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Token *TokenCaller) Cash(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "cash")
	return *ret0, err
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Token *TokenSession) Cash() (common.Address, error) {
	return _Token.Contract.Cash(&_Token.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Token *TokenCallerSession) Cash() (common.Address, error) {
	return _Token.Contract.Cash(&_Token.CallOpts)
}

// CashAvailableForTransferFrom is a free data retrieval call binding the contract method 0x1d4b44bb.
//
// Solidity: function cashAvailableForTransferFrom(address _owner, address _sender) view returns(uint256)
func (_Token *TokenCaller) CashAvailableForTransferFrom(opts *bind.CallOpts, _owner common.Address, _sender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "cashAvailableForTransferFrom", _owner, _sender)
	return *ret0, err
}

// CashAvailableForTransferFrom is a free data retrieval call binding the contract method 0x1d4b44bb.
//
// Solidity: function cashAvailableForTransferFrom(address _owner, address _sender) view returns(uint256)
func (_Token *TokenSession) CashAvailableForTransferFrom(_owner common.Address, _sender common.Address) (*big.Int, error) {
	return _Token.Contract.CashAvailableForTransferFrom(&_Token.CallOpts, _owner, _sender)
}

// CashAvailableForTransferFrom is a free data retrieval call binding the contract method 0x1d4b44bb.
//
// Solidity: function cashAvailableForTransferFrom(address _owner, address _sender) view returns(uint256)
func (_Token *TokenCallerSession) CashAvailableForTransferFrom(_owner common.Address, _sender common.Address) (*big.Int, error) {
	return _Token.Contract.CashAvailableForTransferFrom(&_Token.CallOpts, _owner, _sender)
}

// CreateZeroXOrder is a free data retrieval call binding the contract method 0xf94515f0.
//
// Solidity: function createZeroXOrder(uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns(IExchangeOrder _zeroXOrder, bytes32 _orderHash)
func (_Token *TokenCaller) CreateZeroXOrder(opts *bind.CallOpts, _type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	ret := new(struct {
		ZeroXOrder IExchangeOrder
		OrderHash  [32]byte
	})
	out := ret
	err := _Token.contract.Call(opts, out, "createZeroXOrder", _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)
	return *ret, err
}

// CreateZeroXOrder is a free data retrieval call binding the contract method 0xf94515f0.
//
// Solidity: function createZeroXOrder(uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns(IExchangeOrder _zeroXOrder, bytes32 _orderHash)
func (_Token *TokenSession) CreateZeroXOrder(_type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	return _Token.Contract.CreateZeroXOrder(&_Token.CallOpts, _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)
}

// CreateZeroXOrder is a free data retrieval call binding the contract method 0xf94515f0.
//
// Solidity: function createZeroXOrder(uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns(IExchangeOrder _zeroXOrder, bytes32 _orderHash)
func (_Token *TokenCallerSession) CreateZeroXOrder(_type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	return _Token.Contract.CreateZeroXOrder(&_Token.CallOpts, _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)
}

// CreateZeroXOrderFor is a free data retrieval call binding the contract method 0xe186955b.
//
// Solidity: function createZeroXOrderFor(address _maker, uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns(IExchangeOrder _zeroXOrder, bytes32 _orderHash)
func (_Token *TokenCaller) CreateZeroXOrderFor(opts *bind.CallOpts, _maker common.Address, _type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	ret := new(struct {
		ZeroXOrder IExchangeOrder
		OrderHash  [32]byte
	})
	out := ret
	err := _Token.contract.Call(opts, out, "createZeroXOrderFor", _maker, _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)
	return *ret, err
}

// CreateZeroXOrderFor is a free data retrieval call binding the contract method 0xe186955b.
//
// Solidity: function createZeroXOrderFor(address _maker, uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns(IExchangeOrder _zeroXOrder, bytes32 _orderHash)
func (_Token *TokenSession) CreateZeroXOrderFor(_maker common.Address, _type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	return _Token.Contract.CreateZeroXOrderFor(&_Token.CallOpts, _maker, _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)
}

// CreateZeroXOrderFor is a free data retrieval call binding the contract method 0xe186955b.
//
// Solidity: function createZeroXOrderFor(address _maker, uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) view returns(IExchangeOrder _zeroXOrder, bytes32 _orderHash)
func (_Token *TokenCallerSession) CreateZeroXOrderFor(_maker common.Address, _type uint8, _attoshares *big.Int, _price *big.Int, _market common.Address, _outcome uint8, _expirationTimeSeconds *big.Int, _salt *big.Int) (struct {
	ZeroXOrder IExchangeOrder
	OrderHash  [32]byte
}, error) {
	return _Token.Contract.CreateZeroXOrderFor(&_Token.CallOpts, _maker, _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt)
}

// CreatorHasFundsForTrade is a free data retrieval call binding the contract method 0xe264cff1.
//
// Solidity: function creatorHasFundsForTrade(IExchangeOrder _order, uint256 _amount) view returns(bool)
func (_Token *TokenCaller) CreatorHasFundsForTrade(opts *bind.CallOpts, _order IExchangeOrder, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "creatorHasFundsForTrade", _order, _amount)
	return *ret0, err
}

// CreatorHasFundsForTrade is a free data retrieval call binding the contract method 0xe264cff1.
//
// Solidity: function creatorHasFundsForTrade(IExchangeOrder _order, uint256 _amount) view returns(bool)
func (_Token *TokenSession) CreatorHasFundsForTrade(_order IExchangeOrder, _amount *big.Int) (bool, error) {
	return _Token.Contract.CreatorHasFundsForTrade(&_Token.CallOpts, _order, _amount)
}

// CreatorHasFundsForTrade is a free data retrieval call binding the contract method 0xe264cff1.
//
// Solidity: function creatorHasFundsForTrade(IExchangeOrder _order, uint256 _amount) view returns(bool)
func (_Token *TokenCallerSession) CreatorHasFundsForTrade(_order IExchangeOrder, _amount *big.Int) (bool, error) {
	return _Token.Contract.CreatorHasFundsForTrade(&_Token.CallOpts, _order, _amount)
}

// DecodeAssetData is a free data retrieval call binding the contract method 0x7cb11125.
//
// Solidity: function decodeAssetData(bytes _assetData) view returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_Token *TokenCaller) DecodeAssetData(opts *bind.CallOpts, _assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	ret := new(struct {
		AssetProxyId [4]byte
		TokenAddress common.Address
		TokenIds     []*big.Int
		TokenValues  []*big.Int
		CallbackData []byte
	})
	out := ret
	err := _Token.contract.Call(opts, out, "decodeAssetData", _assetData)
	return *ret, err
}

// DecodeAssetData is a free data retrieval call binding the contract method 0x7cb11125.
//
// Solidity: function decodeAssetData(bytes _assetData) view returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_Token *TokenSession) DecodeAssetData(_assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	return _Token.Contract.DecodeAssetData(&_Token.CallOpts, _assetData)
}

// DecodeAssetData is a free data retrieval call binding the contract method 0x7cb11125.
//
// Solidity: function decodeAssetData(bytes _assetData) view returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_Token *TokenCallerSession) DecodeAssetData(_assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	return _Token.Contract.DecodeAssetData(&_Token.CallOpts, _assetData)
}

// DecodeTradeAssetData is a free data retrieval call binding the contract method 0xb3ca8e25.
//
// Solidity: function decodeTradeAssetData(bytes _assetData) pure returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_Token *TokenCaller) DecodeTradeAssetData(opts *bind.CallOpts, _assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	ret := new(struct {
		AssetProxyId [4]byte
		TokenAddress common.Address
		TokenIds     []*big.Int
		TokenValues  []*big.Int
		CallbackData []byte
	})
	out := ret
	err := _Token.contract.Call(opts, out, "decodeTradeAssetData", _assetData)
	return *ret, err
}

// DecodeTradeAssetData is a free data retrieval call binding the contract method 0xb3ca8e25.
//
// Solidity: function decodeTradeAssetData(bytes _assetData) pure returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_Token *TokenSession) DecodeTradeAssetData(_assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	return _Token.Contract.DecodeTradeAssetData(&_Token.CallOpts, _assetData)
}

// DecodeTradeAssetData is a free data retrieval call binding the contract method 0xb3ca8e25.
//
// Solidity: function decodeTradeAssetData(bytes _assetData) pure returns(bytes4 _assetProxyId, address _tokenAddress, uint256[] _tokenIds, uint256[] _tokenValues, bytes _callbackData)
func (_Token *TokenCallerSession) DecodeTradeAssetData(_assetData []byte) (struct {
	AssetProxyId [4]byte
	TokenAddress common.Address
	TokenIds     []*big.Int
	TokenValues  []*big.Int
	CallbackData []byte
}, error) {
	return _Token.Contract.DecodeTradeAssetData(&_Token.CallOpts, _assetData)
}

// EncodeAssetData is a free data retrieval call binding the contract method 0x235f9e08.
//
// Solidity: function encodeAssetData(address _market, uint256 _price, uint8 _outcome, uint8 _type) view returns(bytes _assetData)
func (_Token *TokenCaller) EncodeAssetData(opts *bind.CallOpts, _market common.Address, _price *big.Int, _outcome uint8, _type uint8) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "encodeAssetData", _market, _price, _outcome, _type)
	return *ret0, err
}

// EncodeAssetData is a free data retrieval call binding the contract method 0x235f9e08.
//
// Solidity: function encodeAssetData(address _market, uint256 _price, uint8 _outcome, uint8 _type) view returns(bytes _assetData)
func (_Token *TokenSession) EncodeAssetData(_market common.Address, _price *big.Int, _outcome uint8, _type uint8) ([]byte, error) {
	return _Token.Contract.EncodeAssetData(&_Token.CallOpts, _market, _price, _outcome, _type)
}

// EncodeAssetData is a free data retrieval call binding the contract method 0x235f9e08.
//
// Solidity: function encodeAssetData(address _market, uint256 _price, uint8 _outcome, uint8 _type) view returns(bytes _assetData)
func (_Token *TokenCallerSession) EncodeAssetData(_market common.Address, _price *big.Int, _outcome uint8, _type uint8) ([]byte, error) {
	return _Token.Contract.EncodeAssetData(&_Token.CallOpts, _market, _price, _outcome, _type)
}

// EncodeEIP1271OrderWithHash is a free data retrieval call binding the contract method 0x4bf1ee87.
//
// Solidity: function encodeEIP1271OrderWithHash(IExchangeOrder _zeroXOrder, bytes32 _orderHash) pure returns(bytes encoded)
func (_Token *TokenCaller) EncodeEIP1271OrderWithHash(opts *bind.CallOpts, _zeroXOrder IExchangeOrder, _orderHash [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "encodeEIP1271OrderWithHash", _zeroXOrder, _orderHash)
	return *ret0, err
}

// EncodeEIP1271OrderWithHash is a free data retrieval call binding the contract method 0x4bf1ee87.
//
// Solidity: function encodeEIP1271OrderWithHash(IExchangeOrder _zeroXOrder, bytes32 _orderHash) pure returns(bytes encoded)
func (_Token *TokenSession) EncodeEIP1271OrderWithHash(_zeroXOrder IExchangeOrder, _orderHash [32]byte) ([]byte, error) {
	return _Token.Contract.EncodeEIP1271OrderWithHash(&_Token.CallOpts, _zeroXOrder, _orderHash)
}

// EncodeEIP1271OrderWithHash is a free data retrieval call binding the contract method 0x4bf1ee87.
//
// Solidity: function encodeEIP1271OrderWithHash(IExchangeOrder _zeroXOrder, bytes32 _orderHash) pure returns(bytes encoded)
func (_Token *TokenCallerSession) EncodeEIP1271OrderWithHash(_zeroXOrder IExchangeOrder, _orderHash [32]byte) ([]byte, error) {
	return _Token.Contract.EncodeEIP1271OrderWithHash(&_Token.CallOpts, _zeroXOrder, _orderHash)
}

// EstimateProtocolFeeCostInCash is a free data retrieval call binding the contract method 0x3fc6f627.
//
// Solidity: function estimateProtocolFeeCostInCash(uint256 _numOrders, uint256 _gasPrice) view returns(uint256)
func (_Token *TokenCaller) EstimateProtocolFeeCostInCash(opts *bind.CallOpts, _numOrders *big.Int, _gasPrice *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "estimateProtocolFeeCostInCash", _numOrders, _gasPrice)
	return *ret0, err
}

// EstimateProtocolFeeCostInCash is a free data retrieval call binding the contract method 0x3fc6f627.
//
// Solidity: function estimateProtocolFeeCostInCash(uint256 _numOrders, uint256 _gasPrice) view returns(uint256)
func (_Token *TokenSession) EstimateProtocolFeeCostInCash(_numOrders *big.Int, _gasPrice *big.Int) (*big.Int, error) {
	return _Token.Contract.EstimateProtocolFeeCostInCash(&_Token.CallOpts, _numOrders, _gasPrice)
}

// EstimateProtocolFeeCostInCash is a free data retrieval call binding the contract method 0x3fc6f627.
//
// Solidity: function estimateProtocolFeeCostInCash(uint256 _numOrders, uint256 _gasPrice) view returns(uint256)
func (_Token *TokenCallerSession) EstimateProtocolFeeCostInCash(_numOrders *big.Int, _gasPrice *big.Int) (*big.Int, error) {
	return _Token.Contract.EstimateProtocolFeeCostInCash(&_Token.CallOpts, _numOrders, _gasPrice)
}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_Token *TokenCaller) EthExchange(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "ethExchange")
	return *ret0, err
}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_Token *TokenSession) EthExchange() (common.Address, error) {
	return _Token.Contract.EthExchange(&_Token.CallOpts)
}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_Token *TokenCallerSession) EthExchange() (common.Address, error) {
	return _Token.Contract.EthExchange(&_Token.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Token *TokenCaller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "exchange")
	return *ret0, err
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Token *TokenSession) Exchange() (common.Address, error) {
	return _Token.Contract.Exchange(&_Token.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Token *TokenCallerSession) Exchange() (common.Address, error) {
	return _Token.Contract.Exchange(&_Token.CallOpts)
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_Token *TokenCaller) FillOrder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "fillOrder")
	return *ret0, err
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_Token *TokenSession) FillOrder() (common.Address, error) {
	return _Token.Contract.FillOrder(&_Token.CallOpts)
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_Token *TokenCallerSession) FillOrder() (common.Address, error) {
	return _Token.Contract.FillOrder(&_Token.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Token *TokenCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getAmountIn", amountOut, reserveIn, reserveOut)
	return *ret0, err
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Token *TokenSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Token.Contract.GetAmountIn(&_Token.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Token *TokenCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Token.Contract.GetAmountIn(&_Token.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_Token *TokenCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_Token *TokenSession) GetInitialized() (bool, error) {
	return _Token.Contract.GetInitialized(&_Token.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_Token *TokenCallerSession) GetInitialized() (bool, error) {
	return _Token.Contract.GetInitialized(&_Token.CallOpts)
}

// GetTokenId is a free data retrieval call binding the contract method 0x1f981bde.
//
// Solidity: function getTokenId(address _market, uint256 _price, uint8 _outcome, uint8 _type) pure returns(uint256 _tokenId)
func (_Token *TokenCaller) GetTokenId(opts *bind.CallOpts, _market common.Address, _price *big.Int, _outcome uint8, _type uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getTokenId", _market, _price, _outcome, _type)
	return *ret0, err
}

// GetTokenId is a free data retrieval call binding the contract method 0x1f981bde.
//
// Solidity: function getTokenId(address _market, uint256 _price, uint8 _outcome, uint8 _type) pure returns(uint256 _tokenId)
func (_Token *TokenSession) GetTokenId(_market common.Address, _price *big.Int, _outcome uint8, _type uint8) (*big.Int, error) {
	return _Token.Contract.GetTokenId(&_Token.CallOpts, _market, _price, _outcome, _type)
}

// GetTokenId is a free data retrieval call binding the contract method 0x1f981bde.
//
// Solidity: function getTokenId(address _market, uint256 _price, uint8 _outcome, uint8 _type) pure returns(uint256 _tokenId)
func (_Token *TokenCallerSession) GetTokenId(_market common.Address, _price *big.Int, _outcome uint8, _type uint8) (*big.Int, error) {
	return _Token.Contract.GetTokenId(&_Token.CallOpts, _market, _price, _outcome, _type)
}

// GetTokenIdFromOrder is a free data retrieval call binding the contract method 0xafc6d8ba.
//
// Solidity: function getTokenIdFromOrder(IExchangeOrder _order) view returns(uint256 _tokenId)
func (_Token *TokenCaller) GetTokenIdFromOrder(opts *bind.CallOpts, _order IExchangeOrder) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getTokenIdFromOrder", _order)
	return *ret0, err
}

// GetTokenIdFromOrder is a free data retrieval call binding the contract method 0xafc6d8ba.
//
// Solidity: function getTokenIdFromOrder(IExchangeOrder _order) view returns(uint256 _tokenId)
func (_Token *TokenSession) GetTokenIdFromOrder(_order IExchangeOrder) (*big.Int, error) {
	return _Token.Contract.GetTokenIdFromOrder(&_Token.CallOpts, _order)
}

// GetTokenIdFromOrder is a free data retrieval call binding the contract method 0xafc6d8ba.
//
// Solidity: function getTokenIdFromOrder(IExchangeOrder _order) view returns(uint256 _tokenId)
func (_Token *TokenCallerSession) GetTokenIdFromOrder(_order IExchangeOrder) (*big.Int, error) {
	return _Token.Contract.GetTokenIdFromOrder(&_Token.CallOpts, _order)
}

// GetTransferFromAllowed is a free data retrieval call binding the contract method 0x09c1df9f.
//
// Solidity: function getTransferFromAllowed() view returns(bool)
func (_Token *TokenCaller) GetTransferFromAllowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getTransferFromAllowed")
	return *ret0, err
}

// GetTransferFromAllowed is a free data retrieval call binding the contract method 0x09c1df9f.
//
// Solidity: function getTransferFromAllowed() view returns(bool)
func (_Token *TokenSession) GetTransferFromAllowed() (bool, error) {
	return _Token.Contract.GetTransferFromAllowed(&_Token.CallOpts)
}

// GetTransferFromAllowed is a free data retrieval call binding the contract method 0x09c1df9f.
//
// Solidity: function getTransferFromAllowed() view returns(bool)
func (_Token *TokenCallerSession) GetTransferFromAllowed() (bool, error) {
	return _Token.Contract.GetTransferFromAllowed(&_Token.CallOpts)
}

// GetZeroXTradeTokenData is a free data retrieval call binding the contract method 0x00ee3542.
//
// Solidity: function getZeroXTradeTokenData(bytes _assetData) view returns(address _token, uint256 _tokenId)
func (_Token *TokenCaller) GetZeroXTradeTokenData(opts *bind.CallOpts, _assetData []byte) (struct {
	Token   common.Address
	TokenId *big.Int
}, error) {
	ret := new(struct {
		Token   common.Address
		TokenId *big.Int
	})
	out := ret
	err := _Token.contract.Call(opts, out, "getZeroXTradeTokenData", _assetData)
	return *ret, err
}

// GetZeroXTradeTokenData is a free data retrieval call binding the contract method 0x00ee3542.
//
// Solidity: function getZeroXTradeTokenData(bytes _assetData) view returns(address _token, uint256 _tokenId)
func (_Token *TokenSession) GetZeroXTradeTokenData(_assetData []byte) (struct {
	Token   common.Address
	TokenId *big.Int
}, error) {
	return _Token.Contract.GetZeroXTradeTokenData(&_Token.CallOpts, _assetData)
}

// GetZeroXTradeTokenData is a free data retrieval call binding the contract method 0x00ee3542.
//
// Solidity: function getZeroXTradeTokenData(bytes _assetData) view returns(address _token, uint256 _tokenId)
func (_Token *TokenCallerSession) GetZeroXTradeTokenData(_assetData []byte) (struct {
	Token   common.Address
	TokenId *big.Int
}, error) {
	return _Token.Contract.GetZeroXTradeTokenData(&_Token.CallOpts, _assetData)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Token.Contract.IsApprovedForAll(&_Token.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Token.Contract.IsApprovedForAll(&_Token.CallOpts, owner, operator)
}

// IsOrderAmountValid is a free data retrieval call binding the contract method 0x9c6fbce5.
//
// Solidity: function isOrderAmountValid(address _market, uint256 _orderAmount) view returns(bool)
func (_Token *TokenCaller) IsOrderAmountValid(opts *bind.CallOpts, _market common.Address, _orderAmount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isOrderAmountValid", _market, _orderAmount)
	return *ret0, err
}

// IsOrderAmountValid is a free data retrieval call binding the contract method 0x9c6fbce5.
//
// Solidity: function isOrderAmountValid(address _market, uint256 _orderAmount) view returns(bool)
func (_Token *TokenSession) IsOrderAmountValid(_market common.Address, _orderAmount *big.Int) (bool, error) {
	return _Token.Contract.IsOrderAmountValid(&_Token.CallOpts, _market, _orderAmount)
}

// IsOrderAmountValid is a free data retrieval call binding the contract method 0x9c6fbce5.
//
// Solidity: function isOrderAmountValid(address _market, uint256 _orderAmount) view returns(bool)
func (_Token *TokenCallerSession) IsOrderAmountValid(_market common.Address, _orderAmount *big.Int) (bool, error) {
	return _Token.Contract.IsOrderAmountValid(&_Token.CallOpts, _market, _orderAmount)
}

// ParseOrderData is a free data retrieval call binding the contract method 0x61a4760b.
//
// Solidity: function parseOrderData(IExchangeOrder _order) view returns(IZeroXTradeAugurOrderData _data)
func (_Token *TokenCaller) ParseOrderData(opts *bind.CallOpts, _order IExchangeOrder) (IZeroXTradeAugurOrderData, error) {
	var (
		ret0 = new(IZeroXTradeAugurOrderData)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "parseOrderData", _order)
	return *ret0, err
}

// ParseOrderData is a free data retrieval call binding the contract method 0x61a4760b.
//
// Solidity: function parseOrderData(IExchangeOrder _order) view returns(IZeroXTradeAugurOrderData _data)
func (_Token *TokenSession) ParseOrderData(_order IExchangeOrder) (IZeroXTradeAugurOrderData, error) {
	return _Token.Contract.ParseOrderData(&_Token.CallOpts, _order)
}

// ParseOrderData is a free data retrieval call binding the contract method 0x61a4760b.
//
// Solidity: function parseOrderData(IExchangeOrder _order) view returns(IZeroXTradeAugurOrderData _data)
func (_Token *TokenCallerSession) ParseOrderData(_order IExchangeOrder) (IZeroXTradeAugurOrderData, error) {
	return _Token.Contract.ParseOrderData(&_Token.CallOpts, _order)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_Token *TokenCaller) ShareToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "shareToken")
	return *ret0, err
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_Token *TokenSession) ShareToken() (common.Address, error) {
	return _Token.Contract.ShareToken(&_Token.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_Token *TokenCallerSession) ShareToken() (common.Address, error) {
	return _Token.Contract.ShareToken(&_Token.CallOpts)
}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_Token *TokenCaller) Token0IsCash(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "token0IsCash")
	return *ret0, err
}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_Token *TokenSession) Token0IsCash() (bool, error) {
	return _Token.Contract.Token0IsCash(&_Token.CallOpts)
}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_Token *TokenCallerSession) Token0IsCash() (bool, error) {
	return _Token.Contract.Token0IsCash(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply", id)
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Token *TokenSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts, id)
}

// UnpackTokenId is a free data retrieval call binding the contract method 0x26afd2e8.
//
// Solidity: function unpackTokenId(uint256 _tokenId) pure returns(address _market, uint256 _price, uint8 _outcome, uint8 _type)
func (_Token *TokenCaller) UnpackTokenId(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Market  common.Address
	Price   *big.Int
	Outcome uint8
	Type    uint8
}, error) {
	ret := new(struct {
		Market  common.Address
		Price   *big.Int
		Outcome uint8
		Type    uint8
	})
	out := ret
	err := _Token.contract.Call(opts, out, "unpackTokenId", _tokenId)
	return *ret, err
}

// UnpackTokenId is a free data retrieval call binding the contract method 0x26afd2e8.
//
// Solidity: function unpackTokenId(uint256 _tokenId) pure returns(address _market, uint256 _price, uint8 _outcome, uint8 _type)
func (_Token *TokenSession) UnpackTokenId(_tokenId *big.Int) (struct {
	Market  common.Address
	Price   *big.Int
	Outcome uint8
	Type    uint8
}, error) {
	return _Token.Contract.UnpackTokenId(&_Token.CallOpts, _tokenId)
}

// UnpackTokenId is a free data retrieval call binding the contract method 0x26afd2e8.
//
// Solidity: function unpackTokenId(uint256 _tokenId) pure returns(address _market, uint256 _price, uint8 _outcome, uint8 _type)
func (_Token *TokenCallerSession) UnpackTokenId(_tokenId *big.Int) (struct {
	Market  common.Address
	Price   *big.Int
	Outcome uint8
	Type    uint8
}, error) {
	return _Token.Contract.UnpackTokenId(&_Token.CallOpts, _tokenId)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x4ea96c30.
//
// Solidity: function cancelOrders([]IExchangeOrder _orders, bytes[] _signatures, uint256 _maxProtocolFeeDai) returns(bool)
func (_Token *TokenTransactor) CancelOrders(opts *bind.TransactOpts, _orders []IExchangeOrder, _signatures [][]byte, _maxProtocolFeeDai *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "cancelOrders", _orders, _signatures, _maxProtocolFeeDai)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x4ea96c30.
//
// Solidity: function cancelOrders([]IExchangeOrder _orders, bytes[] _signatures, uint256 _maxProtocolFeeDai) returns(bool)
func (_Token *TokenSession) CancelOrders(_orders []IExchangeOrder, _signatures [][]byte, _maxProtocolFeeDai *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CancelOrders(&_Token.TransactOpts, _orders, _signatures, _maxProtocolFeeDai)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x4ea96c30.
//
// Solidity: function cancelOrders([]IExchangeOrder _orders, bytes[] _signatures, uint256 _maxProtocolFeeDai) returns(bool)
func (_Token *TokenTransactorSession) CancelOrders(_orders []IExchangeOrder, _signatures [][]byte, _maxProtocolFeeDai *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CancelOrders(&_Token.TransactOpts, _orders, _signatures, _maxProtocolFeeDai)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) returns()
func (_Token *TokenTransactor) Initialize(opts *bind.TransactOpts, _augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "initialize", _augur, _augurTrading)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) returns()
func (_Token *TokenSession) Initialize(_augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, _augur, _augurTrading)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) returns()
func (_Token *TokenTransactorSession) Initialize(_augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, _augur, _augurTrading)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_Token *TokenTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_Token *TokenSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.SafeBatchTransferFrom(&_Token.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_Token *TokenTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.SafeBatchTransferFrom(&_Token.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_Token *TokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_Token *TokenSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom(&_Token.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_Token *TokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom(&_Token.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.Contract.SetApprovalForAll(&_Token.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.Contract.SetApprovalForAll(&_Token.TransactOpts, operator, approved)
}

// Trade is a paid mutator transaction binding the contract method 0x2f562016.
//
// Solidity: function trade(uint256 _requestedFillAmount, bytes32 _fingerprint, bytes32 _tradeGroupId, uint256 _maxProtocolFeeDai, uint256 _maxTrades, []IExchangeOrder _orders, bytes[] _signatures) payable returns(uint256)
func (_Token *TokenTransactor) Trade(opts *bind.TransactOpts, _requestedFillAmount *big.Int, _fingerprint [32]byte, _tradeGroupId [32]byte, _maxProtocolFeeDai *big.Int, _maxTrades *big.Int, _orders []IExchangeOrder, _signatures [][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "trade", _requestedFillAmount, _fingerprint, _tradeGroupId, _maxProtocolFeeDai, _maxTrades, _orders, _signatures)
}

// Trade is a paid mutator transaction binding the contract method 0x2f562016.
//
// Solidity: function trade(uint256 _requestedFillAmount, bytes32 _fingerprint, bytes32 _tradeGroupId, uint256 _maxProtocolFeeDai, uint256 _maxTrades, []IExchangeOrder _orders, bytes[] _signatures) payable returns(uint256)
func (_Token *TokenSession) Trade(_requestedFillAmount *big.Int, _fingerprint [32]byte, _tradeGroupId [32]byte, _maxProtocolFeeDai *big.Int, _maxTrades *big.Int, _orders []IExchangeOrder, _signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.Trade(&_Token.TransactOpts, _requestedFillAmount, _fingerprint, _tradeGroupId, _maxProtocolFeeDai, _maxTrades, _orders, _signatures)
}

// Trade is a paid mutator transaction binding the contract method 0x2f562016.
//
// Solidity: function trade(uint256 _requestedFillAmount, bytes32 _fingerprint, bytes32 _tradeGroupId, uint256 _maxProtocolFeeDai, uint256 _maxTrades, []IExchangeOrder _orders, bytes[] _signatures) payable returns(uint256)
func (_Token *TokenTransactorSession) Trade(_requestedFillAmount *big.Int, _fingerprint [32]byte, _tradeGroupId [32]byte, _maxProtocolFeeDai *big.Int, _maxTrades *big.Int, _orders []IExchangeOrder, _signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.Trade(&_Token.TransactOpts, _requestedFillAmount, _fingerprint, _tradeGroupId, _maxProtocolFeeDai, _maxTrades, _orders, _signatures)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Token.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Token.Contract.Fallback(&_Token.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Token.Contract.Fallback(&_Token.TransactOpts, calldata)
}

// TokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Token contract.
type TokenApprovalForAllIterator struct {
	Event *TokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApprovalForAll)
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
		it.Event = new(TokenApprovalForAll)
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
func (it *TokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApprovalForAll represents a ApprovalForAll event raised by the Token contract.
type TokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Token *TokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalForAllIterator{contract: _Token.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Token *TokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApprovalForAll)
				if err := _Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Token *TokenFilterer) ParseApprovalForAll(log types.Log) (*TokenApprovalForAll, error) {
	event := new(TokenApprovalForAll)
	if err := _Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Token contract.
type TokenTransferBatchIterator struct {
	Event *TokenTransferBatch // Event containing the contract specifics and raw log

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
func (it *TokenTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransferBatch)
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
		it.Event = new(TokenTransferBatch)
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
func (it *TokenTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransferBatch represents a TransferBatch event raised by the Token contract.
type TokenTransferBatch struct {
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
func (_Token *TokenFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TokenTransferBatchIterator, error) {

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

	logs, sub, err := _Token.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferBatchIterator{contract: _Token.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Token *TokenFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *TokenTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Token.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransferBatch)
				if err := _Token.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
func (_Token *TokenFilterer) ParseTransferBatch(log types.Log) (*TokenTransferBatch, error) {
	event := new(TokenTransferBatch)
	if err := _Token.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Token contract.
type TokenTransferSingleIterator struct {
	Event *TokenTransferSingle // Event containing the contract specifics and raw log

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
func (it *TokenTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransferSingle)
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
		it.Event = new(TokenTransferSingle)
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
func (it *TokenTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransferSingle represents a TransferSingle event raised by the Token contract.
type TokenTransferSingle struct {
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
func (_Token *TokenFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TokenTransferSingleIterator, error) {

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

	logs, sub, err := _Token.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferSingleIterator{contract: _Token.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Token *TokenFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *TokenTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Token.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransferSingle)
				if err := _Token.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
func (_Token *TokenFilterer) ParseTransferSingle(log types.Log) (*TokenTransferSingle, error) {
	event := new(TokenTransferSingle)
	if err := _Token.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Token contract.
type TokenURIIterator struct {
	Event *TokenURI // Event containing the contract specifics and raw log

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
func (it *TokenURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenURI)
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
		it.Event = new(TokenURI)
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
func (it *TokenURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenURI represents a URI event raised by the Token contract.
type TokenURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Token *TokenFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*TokenURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &TokenURIIterator{contract: _Token.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Token *TokenFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *TokenURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenURI)
				if err := _Token.contract.UnpackLog(event, "URI", log); err != nil {
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
func (_Token *TokenFilterer) ParseURI(log types.Log) (*TokenURI, error) {
	event := new(TokenURI)
	if err := _Token.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	return event, nil
}
