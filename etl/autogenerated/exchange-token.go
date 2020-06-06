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

// LibFillResultsBatchMatchedFillResults is an auto generated low-level Go binding around an user-defined struct.
type LibFillResultsBatchMatchedFillResults struct {
	Left                    []LibFillResultsFillResults
	Right                   []LibFillResultsFillResults
	ProfitInLeftMakerAsset  *big.Int
	ProfitInRightMakerAsset *big.Int
}

// LibFillResultsFillResults is an auto generated low-level Go binding around an user-defined struct.
type LibFillResultsFillResults struct {
	MakerAssetFilledAmount *big.Int
	TakerAssetFilledAmount *big.Int
	MakerFeePaid           *big.Int
	TakerFeePaid           *big.Int
	ProtocolFeePaid        *big.Int
}

// LibFillResultsMatchedFillResults is an auto generated low-level Go binding around an user-defined struct.
type LibFillResultsMatchedFillResults struct {
	Left                    LibFillResultsFillResults
	Right                   LibFillResultsFillResults
	ProfitInLeftMakerAsset  *big.Int
	ProfitInRightMakerAsset *big.Int
}

// LibOrderOrder is an auto generated low-level Go binding around an user-defined struct.
type LibOrderOrder struct {
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

// LibOrderOrderInfo is an auto generated low-level Go binding around an user-defined struct.
type LibOrderOrderInfo struct {
	OrderStatus                 uint8
	OrderHash                   [32]byte
	OrderTakerAssetFilledAmount *big.Int
}

// LibZeroExTransactionZeroExTransaction is an auto generated low-level Go binding around an user-defined struct.
type LibZeroExTransactionZeroExTransaction struct {
	Salt                  *big.Int
	ExpirationTimeSeconds *big.Int
	GasPrice              *big.Int
	SignerAddress         common.Address
	Data                  []byte
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"id\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetProxy\",\"type\":\"address\"}],\"name\":\"AssetProxyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"orderSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderEpoch\",\"type\":\"uint256\"}],\"name\":\"CancelUpTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"name\":\"Fill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProtocolFeeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedProtocolFeeCollector\",\"type\":\"address\"}],\"name\":\"ProtocolFeeCollectorAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProtocolFeeMultiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedProtocolFeeMultiplier\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"SignatureValidatorApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"TransactionExecution\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP1271_MAGIC_VALUE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_EXCHANGE_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedValidators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"batchCancelOrders\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structLibZeroExTransaction.ZeroExTransaction[]\",\"name\":\"transactions\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"batchExecuteTransactions\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"takerAssetFillAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"batchFillOrKillOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults[]\",\"name\":\"fillResults\",\"type\":\"tuple[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"takerAssetFillAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"batchFillOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults[]\",\"name\":\"fillResults\",\"type\":\"tuple[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"takerAssetFillAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"batchFillOrdersNoThrow\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults[]\",\"name\":\"fillResults\",\"type\":\"tuple[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"leftOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"rightOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"leftSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"rightSignatures\",\"type\":\"bytes[]\"}],\"name\":\"batchMatchOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults[]\",\"name\":\"left\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults[]\",\"name\":\"right\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"profitInLeftMakerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitInRightMakerAsset\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.BatchMatchedFillResults\",\"name\":\"batchMatchedFillResults\",\"type\":\"tuple\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"leftOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"rightOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"leftSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"rightSignatures\",\"type\":\"bytes[]\"}],\"name\":\"batchMatchOrdersWithMaximalFill\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults[]\",\"name\":\"left\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults[]\",\"name\":\"right\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"profitInLeftMakerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitInRightMakerAsset\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.BatchMatchedFillResults\",\"name\":\"batchMatchedFillResults\",\"type\":\"tuple\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetOrderEpoch\",\"type\":\"uint256\"}],\"name\":\"cancelOrdersUpTo\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentContextAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"detachProtocolFeeCollector\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structLibZeroExTransaction.ZeroExTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"fillOrKillOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"fillResults\",\"type\":\"tuple\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"fillOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"fillResults\",\"type\":\"tuple\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"assetProxyId\",\"type\":\"bytes4\"}],\"name\":\"getAssetProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"assetProxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.OrderStatus\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderTakerAssetFilledAmount\",\"type\":\"uint256\"}],\"internalType\":\"structLibOrder.OrderInfo\",\"name\":\"orderInfo\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidHashSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidOrderSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structLibZeroExTransaction.ZeroExTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidTransactionSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"marketBuyOrdersFillOrKill\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"fillResults\",\"type\":\"tuple\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"marketBuyOrdersNoThrow\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"fillResults\",\"type\":\"tuple\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"marketSellOrdersFillOrKill\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"fillResults\",\"type\":\"tuple\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"marketSellOrdersNoThrow\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"fillResults\",\"type\":\"tuple\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"leftOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"rightOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"leftSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rightSignature\",\"type\":\"bytes\"}],\"name\":\"matchOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"right\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"profitInLeftMakerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitInRightMakerAsset\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.MatchedFillResults\",\"name\":\"matchedFillResults\",\"type\":\"tuple\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"leftOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerFeeAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerFeeAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"rightOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"leftSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rightSignature\",\"type\":\"bytes\"}],\"name\":\"matchOrdersWithMaximalFill\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.FillResults\",\"name\":\"right\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"profitInLeftMakerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profitInRightMakerAsset\",\"type\":\"uint256\"}],\"internalType\":\"structLibFillResults.MatchedFillResults\",\"name\":\"matchedFillResults\",\"type\":\"tuple\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"orderEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"preSign\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"preSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolFeeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolFeeMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetProxy\",\"type\":\"address\"}],\"name\":\"registerAssetProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"updatedProtocolFeeCollector\",\"type\":\"address\"}],\"name\":\"setProtocolFeeCollectorAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"updatedProtocolFeeMultiplier\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeMultiplier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approval\",\"type\":\"bool\"}],\"name\":\"setSignatureValidatorApproval\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"assetData\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"fromAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"toAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"simulateDispatchTransferFromCalls\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactionsExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() view returns(bytes4)
func (_Token *TokenCaller) EIP1271MAGICVALUE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "EIP1271_MAGIC_VALUE")
	return *ret0, err
}

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() view returns(bytes4)
func (_Token *TokenSession) EIP1271MAGICVALUE() ([4]byte, error) {
	return _Token.Contract.EIP1271MAGICVALUE(&_Token.CallOpts)
}

// EIP1271MAGICVALUE is a free data retrieval call binding the contract method 0xdd885e2d.
//
// Solidity: function EIP1271_MAGIC_VALUE() view returns(bytes4)
func (_Token *TokenCallerSession) EIP1271MAGICVALUE() ([4]byte, error) {
	return _Token.Contract.EIP1271MAGICVALUE(&_Token.CallOpts)
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() view returns(bytes32)
func (_Token *TokenCaller) EIP712EXCHANGEDOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "EIP712_EXCHANGE_DOMAIN_HASH")
	return *ret0, err
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() view returns(bytes32)
func (_Token *TokenSession) EIP712EXCHANGEDOMAINHASH() ([32]byte, error) {
	return _Token.Contract.EIP712EXCHANGEDOMAINHASH(&_Token.CallOpts)
}

// EIP712EXCHANGEDOMAINHASH is a free data retrieval call binding the contract method 0xc26cfecd.
//
// Solidity: function EIP712_EXCHANGE_DOMAIN_HASH() view returns(bytes32)
func (_Token *TokenCallerSession) EIP712EXCHANGEDOMAINHASH() ([32]byte, error) {
	return _Token.Contract.EIP712EXCHANGEDOMAINHASH(&_Token.CallOpts)
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) view returns(bool)
func (_Token *TokenCaller) AllowedValidators(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowedValidators", arg0, arg1)
	return *ret0, err
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) view returns(bool)
func (_Token *TokenSession) AllowedValidators(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Token.Contract.AllowedValidators(&_Token.CallOpts, arg0, arg1)
}

// AllowedValidators is a free data retrieval call binding the contract method 0x7b8e3514.
//
// Solidity: function allowedValidators(address , address ) view returns(bool)
func (_Token *TokenCallerSession) AllowedValidators(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Token.Contract.AllowedValidators(&_Token.CallOpts, arg0, arg1)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Token *TokenCaller) Cancelled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "cancelled", arg0)
	return *ret0, err
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Token *TokenSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _Token.Contract.Cancelled(&_Token.CallOpts, arg0)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Token *TokenCallerSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _Token.Contract.Cancelled(&_Token.CallOpts, arg0)
}

// CurrentContextAddress is a free data retrieval call binding the contract method 0xeea086ba.
//
// Solidity: function currentContextAddress() view returns(address)
func (_Token *TokenCaller) CurrentContextAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "currentContextAddress")
	return *ret0, err
}

// CurrentContextAddress is a free data retrieval call binding the contract method 0xeea086ba.
//
// Solidity: function currentContextAddress() view returns(address)
func (_Token *TokenSession) CurrentContextAddress() (common.Address, error) {
	return _Token.Contract.CurrentContextAddress(&_Token.CallOpts)
}

// CurrentContextAddress is a free data retrieval call binding the contract method 0xeea086ba.
//
// Solidity: function currentContextAddress() view returns(address)
func (_Token *TokenCallerSession) CurrentContextAddress() (common.Address, error) {
	return _Token.Contract.CurrentContextAddress(&_Token.CallOpts)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Token *TokenCaller) Filled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "filled", arg0)
	return *ret0, err
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Token *TokenSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Token.Contract.Filled(&_Token.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Token *TokenCallerSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Token.Contract.Filled(&_Token.CallOpts, arg0)
}

// GetAssetProxy is a free data retrieval call binding the contract method 0x60704108.
//
// Solidity: function getAssetProxy(bytes4 assetProxyId) view returns(address assetProxy)
func (_Token *TokenCaller) GetAssetProxy(opts *bind.CallOpts, assetProxyId [4]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getAssetProxy", assetProxyId)
	return *ret0, err
}

// GetAssetProxy is a free data retrieval call binding the contract method 0x60704108.
//
// Solidity: function getAssetProxy(bytes4 assetProxyId) view returns(address assetProxy)
func (_Token *TokenSession) GetAssetProxy(assetProxyId [4]byte) (common.Address, error) {
	return _Token.Contract.GetAssetProxy(&_Token.CallOpts, assetProxyId)
}

// GetAssetProxy is a free data retrieval call binding the contract method 0x60704108.
//
// Solidity: function getAssetProxy(bytes4 assetProxyId) view returns(address assetProxy)
func (_Token *TokenCallerSession) GetAssetProxy(assetProxyId [4]byte) (common.Address, error) {
	return _Token.Contract.GetAssetProxy(&_Token.CallOpts, assetProxyId)
}

// GetOrderInfo is a free data retrieval call binding the contract method 0x9d3fa4b9.
//
// Solidity: function getOrderInfo(LibOrderOrder order) view returns(LibOrderOrderInfo orderInfo)
func (_Token *TokenCaller) GetOrderInfo(opts *bind.CallOpts, order LibOrderOrder) (LibOrderOrderInfo, error) {
	var (
		ret0 = new(LibOrderOrderInfo)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getOrderInfo", order)
	return *ret0, err
}

// GetOrderInfo is a free data retrieval call binding the contract method 0x9d3fa4b9.
//
// Solidity: function getOrderInfo(LibOrderOrder order) view returns(LibOrderOrderInfo orderInfo)
func (_Token *TokenSession) GetOrderInfo(order LibOrderOrder) (LibOrderOrderInfo, error) {
	return _Token.Contract.GetOrderInfo(&_Token.CallOpts, order)
}

// GetOrderInfo is a free data retrieval call binding the contract method 0x9d3fa4b9.
//
// Solidity: function getOrderInfo(LibOrderOrder order) view returns(LibOrderOrderInfo orderInfo)
func (_Token *TokenCallerSession) GetOrderInfo(order LibOrderOrder) (LibOrderOrderInfo, error) {
	return _Token.Contract.GetOrderInfo(&_Token.CallOpts, order)
}

// IsValidHashSignature is a free data retrieval call binding the contract method 0x8171c407.
//
// Solidity: function isValidHashSignature(bytes32 hash, address signerAddress, bytes signature) view returns(bool isValid)
func (_Token *TokenCaller) IsValidHashSignature(opts *bind.CallOpts, hash [32]byte, signerAddress common.Address, signature []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isValidHashSignature", hash, signerAddress, signature)
	return *ret0, err
}

// IsValidHashSignature is a free data retrieval call binding the contract method 0x8171c407.
//
// Solidity: function isValidHashSignature(bytes32 hash, address signerAddress, bytes signature) view returns(bool isValid)
func (_Token *TokenSession) IsValidHashSignature(hash [32]byte, signerAddress common.Address, signature []byte) (bool, error) {
	return _Token.Contract.IsValidHashSignature(&_Token.CallOpts, hash, signerAddress, signature)
}

// IsValidHashSignature is a free data retrieval call binding the contract method 0x8171c407.
//
// Solidity: function isValidHashSignature(bytes32 hash, address signerAddress, bytes signature) view returns(bool isValid)
func (_Token *TokenCallerSession) IsValidHashSignature(hash [32]byte, signerAddress common.Address, signature []byte) (bool, error) {
	return _Token.Contract.IsValidHashSignature(&_Token.CallOpts, hash, signerAddress, signature)
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature(LibOrderOrder order, bytes signature) view returns(bool isValid)
func (_Token *TokenCaller) IsValidOrderSignature(opts *bind.CallOpts, order LibOrderOrder, signature []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isValidOrderSignature", order, signature)
	return *ret0, err
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature(LibOrderOrder order, bytes signature) view returns(bool isValid)
func (_Token *TokenSession) IsValidOrderSignature(order LibOrderOrder, signature []byte) (bool, error) {
	return _Token.Contract.IsValidOrderSignature(&_Token.CallOpts, order, signature)
}

// IsValidOrderSignature is a free data retrieval call binding the contract method 0xa12dcc6f.
//
// Solidity: function isValidOrderSignature(LibOrderOrder order, bytes signature) view returns(bool isValid)
func (_Token *TokenCallerSession) IsValidOrderSignature(order LibOrderOrder, signature []byte) (bool, error) {
	return _Token.Contract.IsValidOrderSignature(&_Token.CallOpts, order, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x81d09ba1.
//
// Solidity: function isValidSignature(LibOrderOrder order, bytes32 orderHash, bytes signature) view returns(bool)
func (_Token *TokenCaller) IsValidSignature(opts *bind.CallOpts, order LibOrderOrder, orderHash [32]byte, signature []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isValidSignature", order, orderHash, signature)
	return *ret0, err
}

// IsValidSignature is a free data retrieval call binding the contract method 0x81d09ba1.
//
// Solidity: function isValidSignature(LibOrderOrder order, bytes32 orderHash, bytes signature) view returns(bool)
func (_Token *TokenSession) IsValidSignature(order LibOrderOrder, orderHash [32]byte, signature []byte) (bool, error) {
	return _Token.Contract.IsValidSignature(&_Token.CallOpts, order, orderHash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x81d09ba1.
//
// Solidity: function isValidSignature(LibOrderOrder order, bytes32 orderHash, bytes signature) view returns(bool)
func (_Token *TokenCallerSession) IsValidSignature(order LibOrderOrder, orderHash [32]byte, signature []byte) (bool, error) {
	return _Token.Contract.IsValidSignature(&_Token.CallOpts, order, orderHash, signature)
}

// IsValidTransactionSignature is a free data retrieval call binding the contract method 0x8d45cd23.
//
// Solidity: function isValidTransactionSignature(LibZeroExTransactionZeroExTransaction transaction, bytes signature) view returns(bool isValid)
func (_Token *TokenCaller) IsValidTransactionSignature(opts *bind.CallOpts, transaction LibZeroExTransactionZeroExTransaction, signature []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isValidTransactionSignature", transaction, signature)
	return *ret0, err
}

// IsValidTransactionSignature is a free data retrieval call binding the contract method 0x8d45cd23.
//
// Solidity: function isValidTransactionSignature(LibZeroExTransactionZeroExTransaction transaction, bytes signature) view returns(bool isValid)
func (_Token *TokenSession) IsValidTransactionSignature(transaction LibZeroExTransactionZeroExTransaction, signature []byte) (bool, error) {
	return _Token.Contract.IsValidTransactionSignature(&_Token.CallOpts, transaction, signature)
}

// IsValidTransactionSignature is a free data retrieval call binding the contract method 0x8d45cd23.
//
// Solidity: function isValidTransactionSignature(LibZeroExTransactionZeroExTransaction transaction, bytes signature) view returns(bool isValid)
func (_Token *TokenCallerSession) IsValidTransactionSignature(transaction LibZeroExTransactionZeroExTransaction, signature []byte) (bool, error) {
	return _Token.Contract.IsValidTransactionSignature(&_Token.CallOpts, transaction, signature)
}

// OrderEpoch is a free data retrieval call binding the contract method 0xd9bfa73e.
//
// Solidity: function orderEpoch(address , address ) view returns(uint256)
func (_Token *TokenCaller) OrderEpoch(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "orderEpoch", arg0, arg1)
	return *ret0, err
}

// OrderEpoch is a free data retrieval call binding the contract method 0xd9bfa73e.
//
// Solidity: function orderEpoch(address , address ) view returns(uint256)
func (_Token *TokenSession) OrderEpoch(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Token.Contract.OrderEpoch(&_Token.CallOpts, arg0, arg1)
}

// OrderEpoch is a free data retrieval call binding the contract method 0xd9bfa73e.
//
// Solidity: function orderEpoch(address , address ) view returns(uint256)
func (_Token *TokenCallerSession) OrderEpoch(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Token.Contract.OrderEpoch(&_Token.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) view returns(bool)
func (_Token *TokenCaller) PreSigned(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "preSigned", arg0, arg1)
	return *ret0, err
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) view returns(bool)
func (_Token *TokenSession) PreSigned(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Token.Contract.PreSigned(&_Token.CallOpts, arg0, arg1)
}

// PreSigned is a free data retrieval call binding the contract method 0x82c174d0.
//
// Solidity: function preSigned(bytes32 , address ) view returns(bool)
func (_Token *TokenCallerSession) PreSigned(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Token.Contract.PreSigned(&_Token.CallOpts, arg0, arg1)
}

// ProtocolFeeCollector is a free data retrieval call binding the contract method 0x850a1501.
//
// Solidity: function protocolFeeCollector() view returns(address)
func (_Token *TokenCaller) ProtocolFeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "protocolFeeCollector")
	return *ret0, err
}

// ProtocolFeeCollector is a free data retrieval call binding the contract method 0x850a1501.
//
// Solidity: function protocolFeeCollector() view returns(address)
func (_Token *TokenSession) ProtocolFeeCollector() (common.Address, error) {
	return _Token.Contract.ProtocolFeeCollector(&_Token.CallOpts)
}

// ProtocolFeeCollector is a free data retrieval call binding the contract method 0x850a1501.
//
// Solidity: function protocolFeeCollector() view returns(address)
func (_Token *TokenCallerSession) ProtocolFeeCollector() (common.Address, error) {
	return _Token.Contract.ProtocolFeeCollector(&_Token.CallOpts)
}

// ProtocolFeeMultiplier is a free data retrieval call binding the contract method 0x1ce4c78b.
//
// Solidity: function protocolFeeMultiplier() view returns(uint256)
func (_Token *TokenCaller) ProtocolFeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "protocolFeeMultiplier")
	return *ret0, err
}

// ProtocolFeeMultiplier is a free data retrieval call binding the contract method 0x1ce4c78b.
//
// Solidity: function protocolFeeMultiplier() view returns(uint256)
func (_Token *TokenSession) ProtocolFeeMultiplier() (*big.Int, error) {
	return _Token.Contract.ProtocolFeeMultiplier(&_Token.CallOpts)
}

// ProtocolFeeMultiplier is a free data retrieval call binding the contract method 0x1ce4c78b.
//
// Solidity: function protocolFeeMultiplier() view returns(uint256)
func (_Token *TokenCallerSession) ProtocolFeeMultiplier() (*big.Int, error) {
	return _Token.Contract.ProtocolFeeMultiplier(&_Token.CallOpts)
}

// TransactionsExecuted is a free data retrieval call binding the contract method 0x0228e168.
//
// Solidity: function transactionsExecuted(bytes32 ) view returns(bool)
func (_Token *TokenCaller) TransactionsExecuted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "transactionsExecuted", arg0)
	return *ret0, err
}

// TransactionsExecuted is a free data retrieval call binding the contract method 0x0228e168.
//
// Solidity: function transactionsExecuted(bytes32 ) view returns(bool)
func (_Token *TokenSession) TransactionsExecuted(arg0 [32]byte) (bool, error) {
	return _Token.Contract.TransactionsExecuted(&_Token.CallOpts, arg0)
}

// TransactionsExecuted is a free data retrieval call binding the contract method 0x0228e168.
//
// Solidity: function transactionsExecuted(bytes32 ) view returns(bool)
func (_Token *TokenCallerSession) TransactionsExecuted(arg0 [32]byte) (bool, error) {
	return _Token.Contract.TransactionsExecuted(&_Token.CallOpts, arg0)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xdedfc1f1.
//
// Solidity: function batchCancelOrders([]LibOrderOrder orders) payable returns()
func (_Token *TokenTransactor) BatchCancelOrders(opts *bind.TransactOpts, orders []LibOrderOrder) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchCancelOrders", orders)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xdedfc1f1.
//
// Solidity: function batchCancelOrders([]LibOrderOrder orders) payable returns()
func (_Token *TokenSession) BatchCancelOrders(orders []LibOrderOrder) (*types.Transaction, error) {
	return _Token.Contract.BatchCancelOrders(&_Token.TransactOpts, orders)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xdedfc1f1.
//
// Solidity: function batchCancelOrders([]LibOrderOrder orders) payable returns()
func (_Token *TokenTransactorSession) BatchCancelOrders(orders []LibOrderOrder) (*types.Transaction, error) {
	return _Token.Contract.BatchCancelOrders(&_Token.TransactOpts, orders)
}

// BatchExecuteTransactions is a paid mutator transaction binding the contract method 0xfc74896d.
//
// Solidity: function batchExecuteTransactions([]LibZeroExTransactionZeroExTransaction transactions, bytes[] signatures) payable returns(bytes[] returnData)
func (_Token *TokenTransactor) BatchExecuteTransactions(opts *bind.TransactOpts, transactions []LibZeroExTransactionZeroExTransaction, signatures [][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchExecuteTransactions", transactions, signatures)
}

// BatchExecuteTransactions is a paid mutator transaction binding the contract method 0xfc74896d.
//
// Solidity: function batchExecuteTransactions([]LibZeroExTransactionZeroExTransaction transactions, bytes[] signatures) payable returns(bytes[] returnData)
func (_Token *TokenSession) BatchExecuteTransactions(transactions []LibZeroExTransactionZeroExTransaction, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchExecuteTransactions(&_Token.TransactOpts, transactions, signatures)
}

// BatchExecuteTransactions is a paid mutator transaction binding the contract method 0xfc74896d.
//
// Solidity: function batchExecuteTransactions([]LibZeroExTransactionZeroExTransaction transactions, bytes[] signatures) payable returns(bytes[] returnData)
func (_Token *TokenTransactorSession) BatchExecuteTransactions(transactions []LibZeroExTransactionZeroExTransaction, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchExecuteTransactions(&_Token.TransactOpts, transactions, signatures)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0xbeee2e14.
//
// Solidity: function batchFillOrKillOrders([]LibOrderOrder orders, uint256[] takerAssetFillAmounts, bytes[] signatures) payable returns([]LibFillResultsFillResults fillResults)
func (_Token *TokenTransactor) BatchFillOrKillOrders(opts *bind.TransactOpts, orders []LibOrderOrder, takerAssetFillAmounts []*big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchFillOrKillOrders", orders, takerAssetFillAmounts, signatures)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0xbeee2e14.
//
// Solidity: function batchFillOrKillOrders([]LibOrderOrder orders, uint256[] takerAssetFillAmounts, bytes[] signatures) payable returns([]LibFillResultsFillResults fillResults)
func (_Token *TokenSession) BatchFillOrKillOrders(orders []LibOrderOrder, takerAssetFillAmounts []*big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchFillOrKillOrders(&_Token.TransactOpts, orders, takerAssetFillAmounts, signatures)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0xbeee2e14.
//
// Solidity: function batchFillOrKillOrders([]LibOrderOrder orders, uint256[] takerAssetFillAmounts, bytes[] signatures) payable returns([]LibFillResultsFillResults fillResults)
func (_Token *TokenTransactorSession) BatchFillOrKillOrders(orders []LibOrderOrder, takerAssetFillAmounts []*big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchFillOrKillOrders(&_Token.TransactOpts, orders, takerAssetFillAmounts, signatures)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0x9694a402.
//
// Solidity: function batchFillOrders([]LibOrderOrder orders, uint256[] takerAssetFillAmounts, bytes[] signatures) payable returns([]LibFillResultsFillResults fillResults)
func (_Token *TokenTransactor) BatchFillOrders(opts *bind.TransactOpts, orders []LibOrderOrder, takerAssetFillAmounts []*big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchFillOrders", orders, takerAssetFillAmounts, signatures)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0x9694a402.
//
// Solidity: function batchFillOrders([]LibOrderOrder orders, uint256[] takerAssetFillAmounts, bytes[] signatures) payable returns([]LibFillResultsFillResults fillResults)
func (_Token *TokenSession) BatchFillOrders(orders []LibOrderOrder, takerAssetFillAmounts []*big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchFillOrders(&_Token.TransactOpts, orders, takerAssetFillAmounts, signatures)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0x9694a402.
//
// Solidity: function batchFillOrders([]LibOrderOrder orders, uint256[] takerAssetFillAmounts, bytes[] signatures) payable returns([]LibFillResultsFillResults fillResults)
func (_Token *TokenTransactorSession) BatchFillOrders(orders []LibOrderOrder, takerAssetFillAmounts []*big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchFillOrders(&_Token.TransactOpts, orders, takerAssetFillAmounts, signatures)
}

// BatchFillOrdersNoThrow is a paid mutator transaction binding the contract method 0x8ea8dfe4.
//
// Solidity: function batchFillOrdersNoThrow([]LibOrderOrder orders, uint256[] takerAssetFillAmounts, bytes[] signatures) payable returns([]LibFillResultsFillResults fillResults)
func (_Token *TokenTransactor) BatchFillOrdersNoThrow(opts *bind.TransactOpts, orders []LibOrderOrder, takerAssetFillAmounts []*big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchFillOrdersNoThrow", orders, takerAssetFillAmounts, signatures)
}

// BatchFillOrdersNoThrow is a paid mutator transaction binding the contract method 0x8ea8dfe4.
//
// Solidity: function batchFillOrdersNoThrow([]LibOrderOrder orders, uint256[] takerAssetFillAmounts, bytes[] signatures) payable returns([]LibFillResultsFillResults fillResults)
func (_Token *TokenSession) BatchFillOrdersNoThrow(orders []LibOrderOrder, takerAssetFillAmounts []*big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchFillOrdersNoThrow(&_Token.TransactOpts, orders, takerAssetFillAmounts, signatures)
}

// BatchFillOrdersNoThrow is a paid mutator transaction binding the contract method 0x8ea8dfe4.
//
// Solidity: function batchFillOrdersNoThrow([]LibOrderOrder orders, uint256[] takerAssetFillAmounts, bytes[] signatures) payable returns([]LibFillResultsFillResults fillResults)
func (_Token *TokenTransactorSession) BatchFillOrdersNoThrow(orders []LibOrderOrder, takerAssetFillAmounts []*big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchFillOrdersNoThrow(&_Token.TransactOpts, orders, takerAssetFillAmounts, signatures)
}

// BatchMatchOrders is a paid mutator transaction binding the contract method 0x6fcf3e9e.
//
// Solidity: function batchMatchOrders([]LibOrderOrder leftOrders, []LibOrderOrder rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) payable returns(LibFillResultsBatchMatchedFillResults batchMatchedFillResults)
func (_Token *TokenTransactor) BatchMatchOrders(opts *bind.TransactOpts, leftOrders []LibOrderOrder, rightOrders []LibOrderOrder, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchMatchOrders", leftOrders, rightOrders, leftSignatures, rightSignatures)
}

// BatchMatchOrders is a paid mutator transaction binding the contract method 0x6fcf3e9e.
//
// Solidity: function batchMatchOrders([]LibOrderOrder leftOrders, []LibOrderOrder rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) payable returns(LibFillResultsBatchMatchedFillResults batchMatchedFillResults)
func (_Token *TokenSession) BatchMatchOrders(leftOrders []LibOrderOrder, rightOrders []LibOrderOrder, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchMatchOrders(&_Token.TransactOpts, leftOrders, rightOrders, leftSignatures, rightSignatures)
}

// BatchMatchOrders is a paid mutator transaction binding the contract method 0x6fcf3e9e.
//
// Solidity: function batchMatchOrders([]LibOrderOrder leftOrders, []LibOrderOrder rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) payable returns(LibFillResultsBatchMatchedFillResults batchMatchedFillResults)
func (_Token *TokenTransactorSession) BatchMatchOrders(leftOrders []LibOrderOrder, rightOrders []LibOrderOrder, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchMatchOrders(&_Token.TransactOpts, leftOrders, rightOrders, leftSignatures, rightSignatures)
}

// BatchMatchOrdersWithMaximalFill is a paid mutator transaction binding the contract method 0x6a1a80fd.
//
// Solidity: function batchMatchOrdersWithMaximalFill([]LibOrderOrder leftOrders, []LibOrderOrder rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) payable returns(LibFillResultsBatchMatchedFillResults batchMatchedFillResults)
func (_Token *TokenTransactor) BatchMatchOrdersWithMaximalFill(opts *bind.TransactOpts, leftOrders []LibOrderOrder, rightOrders []LibOrderOrder, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchMatchOrdersWithMaximalFill", leftOrders, rightOrders, leftSignatures, rightSignatures)
}

// BatchMatchOrdersWithMaximalFill is a paid mutator transaction binding the contract method 0x6a1a80fd.
//
// Solidity: function batchMatchOrdersWithMaximalFill([]LibOrderOrder leftOrders, []LibOrderOrder rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) payable returns(LibFillResultsBatchMatchedFillResults batchMatchedFillResults)
func (_Token *TokenSession) BatchMatchOrdersWithMaximalFill(leftOrders []LibOrderOrder, rightOrders []LibOrderOrder, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchMatchOrdersWithMaximalFill(&_Token.TransactOpts, leftOrders, rightOrders, leftSignatures, rightSignatures)
}

// BatchMatchOrdersWithMaximalFill is a paid mutator transaction binding the contract method 0x6a1a80fd.
//
// Solidity: function batchMatchOrdersWithMaximalFill([]LibOrderOrder leftOrders, []LibOrderOrder rightOrders, bytes[] leftSignatures, bytes[] rightSignatures) payable returns(LibFillResultsBatchMatchedFillResults batchMatchedFillResults)
func (_Token *TokenTransactorSession) BatchMatchOrdersWithMaximalFill(leftOrders []LibOrderOrder, rightOrders []LibOrderOrder, leftSignatures [][]byte, rightSignatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.BatchMatchOrdersWithMaximalFill(&_Token.TransactOpts, leftOrders, rightOrders, leftSignatures, rightSignatures)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder(LibOrderOrder order) payable returns()
func (_Token *TokenTransactor) CancelOrder(opts *bind.TransactOpts, order LibOrderOrder) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder(LibOrderOrder order) payable returns()
func (_Token *TokenSession) CancelOrder(order LibOrderOrder) (*types.Transaction, error) {
	return _Token.Contract.CancelOrder(&_Token.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2da62987.
//
// Solidity: function cancelOrder(LibOrderOrder order) payable returns()
func (_Token *TokenTransactorSession) CancelOrder(order LibOrderOrder) (*types.Transaction, error) {
	return _Token.Contract.CancelOrder(&_Token.TransactOpts, order)
}

// CancelOrdersUpTo is a paid mutator transaction binding the contract method 0x4f9559b1.
//
// Solidity: function cancelOrdersUpTo(uint256 targetOrderEpoch) payable returns()
func (_Token *TokenTransactor) CancelOrdersUpTo(opts *bind.TransactOpts, targetOrderEpoch *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "cancelOrdersUpTo", targetOrderEpoch)
}

// CancelOrdersUpTo is a paid mutator transaction binding the contract method 0x4f9559b1.
//
// Solidity: function cancelOrdersUpTo(uint256 targetOrderEpoch) payable returns()
func (_Token *TokenSession) CancelOrdersUpTo(targetOrderEpoch *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CancelOrdersUpTo(&_Token.TransactOpts, targetOrderEpoch)
}

// CancelOrdersUpTo is a paid mutator transaction binding the contract method 0x4f9559b1.
//
// Solidity: function cancelOrdersUpTo(uint256 targetOrderEpoch) payable returns()
func (_Token *TokenTransactorSession) CancelOrdersUpTo(targetOrderEpoch *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CancelOrdersUpTo(&_Token.TransactOpts, targetOrderEpoch)
}

// DetachProtocolFeeCollector is a paid mutator transaction binding the contract method 0x0efca185.
//
// Solidity: function detachProtocolFeeCollector() returns()
func (_Token *TokenTransactor) DetachProtocolFeeCollector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "detachProtocolFeeCollector")
}

// DetachProtocolFeeCollector is a paid mutator transaction binding the contract method 0x0efca185.
//
// Solidity: function detachProtocolFeeCollector() returns()
func (_Token *TokenSession) DetachProtocolFeeCollector() (*types.Transaction, error) {
	return _Token.Contract.DetachProtocolFeeCollector(&_Token.TransactOpts)
}

// DetachProtocolFeeCollector is a paid mutator transaction binding the contract method 0x0efca185.
//
// Solidity: function detachProtocolFeeCollector() returns()
func (_Token *TokenTransactorSession) DetachProtocolFeeCollector() (*types.Transaction, error) {
	return _Token.Contract.DetachProtocolFeeCollector(&_Token.TransactOpts)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2280c910.
//
// Solidity: function executeTransaction(LibZeroExTransactionZeroExTransaction transaction, bytes signature) payable returns(bytes)
func (_Token *TokenTransactor) ExecuteTransaction(opts *bind.TransactOpts, transaction LibZeroExTransactionZeroExTransaction, signature []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "executeTransaction", transaction, signature)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2280c910.
//
// Solidity: function executeTransaction(LibZeroExTransactionZeroExTransaction transaction, bytes signature) payable returns(bytes)
func (_Token *TokenSession) ExecuteTransaction(transaction LibZeroExTransactionZeroExTransaction, signature []byte) (*types.Transaction, error) {
	return _Token.Contract.ExecuteTransaction(&_Token.TransactOpts, transaction, signature)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2280c910.
//
// Solidity: function executeTransaction(LibZeroExTransactionZeroExTransaction transaction, bytes signature) payable returns(bytes)
func (_Token *TokenTransactorSession) ExecuteTransaction(transaction LibZeroExTransactionZeroExTransaction, signature []byte) (*types.Transaction, error) {
	return _Token.Contract.ExecuteTransaction(&_Token.TransactOpts, transaction, signature)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0xe14b58c4.
//
// Solidity: function fillOrKillOrder(LibOrderOrder order, uint256 takerAssetFillAmount, bytes signature) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactor) FillOrKillOrder(opts *bind.TransactOpts, order LibOrderOrder, takerAssetFillAmount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "fillOrKillOrder", order, takerAssetFillAmount, signature)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0xe14b58c4.
//
// Solidity: function fillOrKillOrder(LibOrderOrder order, uint256 takerAssetFillAmount, bytes signature) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenSession) FillOrKillOrder(order LibOrderOrder, takerAssetFillAmount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Token.Contract.FillOrKillOrder(&_Token.TransactOpts, order, takerAssetFillAmount, signature)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0xe14b58c4.
//
// Solidity: function fillOrKillOrder(LibOrderOrder order, uint256 takerAssetFillAmount, bytes signature) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactorSession) FillOrKillOrder(order LibOrderOrder, takerAssetFillAmount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Token.Contract.FillOrKillOrder(&_Token.TransactOpts, order, takerAssetFillAmount, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x9b44d556.
//
// Solidity: function fillOrder(LibOrderOrder order, uint256 takerAssetFillAmount, bytes signature) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactor) FillOrder(opts *bind.TransactOpts, order LibOrderOrder, takerAssetFillAmount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "fillOrder", order, takerAssetFillAmount, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x9b44d556.
//
// Solidity: function fillOrder(LibOrderOrder order, uint256 takerAssetFillAmount, bytes signature) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenSession) FillOrder(order LibOrderOrder, takerAssetFillAmount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Token.Contract.FillOrder(&_Token.TransactOpts, order, takerAssetFillAmount, signature)
}

// FillOrder is a paid mutator transaction binding the contract method 0x9b44d556.
//
// Solidity: function fillOrder(LibOrderOrder order, uint256 takerAssetFillAmount, bytes signature) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactorSession) FillOrder(order LibOrderOrder, takerAssetFillAmount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Token.Contract.FillOrder(&_Token.TransactOpts, order, takerAssetFillAmount, signature)
}

// MarketBuyOrdersFillOrKill is a paid mutator transaction binding the contract method 0x8bc8efb3.
//
// Solidity: function marketBuyOrdersFillOrKill([]LibOrderOrder orders, uint256 makerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactor) MarketBuyOrdersFillOrKill(opts *bind.TransactOpts, orders []LibOrderOrder, makerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "marketBuyOrdersFillOrKill", orders, makerAssetFillAmount, signatures)
}

// MarketBuyOrdersFillOrKill is a paid mutator transaction binding the contract method 0x8bc8efb3.
//
// Solidity: function marketBuyOrdersFillOrKill([]LibOrderOrder orders, uint256 makerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenSession) MarketBuyOrdersFillOrKill(orders []LibOrderOrder, makerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.MarketBuyOrdersFillOrKill(&_Token.TransactOpts, orders, makerAssetFillAmount, signatures)
}

// MarketBuyOrdersFillOrKill is a paid mutator transaction binding the contract method 0x8bc8efb3.
//
// Solidity: function marketBuyOrdersFillOrKill([]LibOrderOrder orders, uint256 makerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactorSession) MarketBuyOrdersFillOrKill(orders []LibOrderOrder, makerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.MarketBuyOrdersFillOrKill(&_Token.TransactOpts, orders, makerAssetFillAmount, signatures)
}

// MarketBuyOrdersNoThrow is a paid mutator transaction binding the contract method 0x78d29ac1.
//
// Solidity: function marketBuyOrdersNoThrow([]LibOrderOrder orders, uint256 makerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactor) MarketBuyOrdersNoThrow(opts *bind.TransactOpts, orders []LibOrderOrder, makerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "marketBuyOrdersNoThrow", orders, makerAssetFillAmount, signatures)
}

// MarketBuyOrdersNoThrow is a paid mutator transaction binding the contract method 0x78d29ac1.
//
// Solidity: function marketBuyOrdersNoThrow([]LibOrderOrder orders, uint256 makerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenSession) MarketBuyOrdersNoThrow(orders []LibOrderOrder, makerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.MarketBuyOrdersNoThrow(&_Token.TransactOpts, orders, makerAssetFillAmount, signatures)
}

// MarketBuyOrdersNoThrow is a paid mutator transaction binding the contract method 0x78d29ac1.
//
// Solidity: function marketBuyOrdersNoThrow([]LibOrderOrder orders, uint256 makerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactorSession) MarketBuyOrdersNoThrow(orders []LibOrderOrder, makerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.MarketBuyOrdersNoThrow(&_Token.TransactOpts, orders, makerAssetFillAmount, signatures)
}

// MarketSellOrdersFillOrKill is a paid mutator transaction binding the contract method 0xa6c3bf33.
//
// Solidity: function marketSellOrdersFillOrKill([]LibOrderOrder orders, uint256 takerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactor) MarketSellOrdersFillOrKill(opts *bind.TransactOpts, orders []LibOrderOrder, takerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "marketSellOrdersFillOrKill", orders, takerAssetFillAmount, signatures)
}

// MarketSellOrdersFillOrKill is a paid mutator transaction binding the contract method 0xa6c3bf33.
//
// Solidity: function marketSellOrdersFillOrKill([]LibOrderOrder orders, uint256 takerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenSession) MarketSellOrdersFillOrKill(orders []LibOrderOrder, takerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.MarketSellOrdersFillOrKill(&_Token.TransactOpts, orders, takerAssetFillAmount, signatures)
}

// MarketSellOrdersFillOrKill is a paid mutator transaction binding the contract method 0xa6c3bf33.
//
// Solidity: function marketSellOrdersFillOrKill([]LibOrderOrder orders, uint256 takerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactorSession) MarketSellOrdersFillOrKill(orders []LibOrderOrder, takerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.MarketSellOrdersFillOrKill(&_Token.TransactOpts, orders, takerAssetFillAmount, signatures)
}

// MarketSellOrdersNoThrow is a paid mutator transaction binding the contract method 0x369da099.
//
// Solidity: function marketSellOrdersNoThrow([]LibOrderOrder orders, uint256 takerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactor) MarketSellOrdersNoThrow(opts *bind.TransactOpts, orders []LibOrderOrder, takerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "marketSellOrdersNoThrow", orders, takerAssetFillAmount, signatures)
}

// MarketSellOrdersNoThrow is a paid mutator transaction binding the contract method 0x369da099.
//
// Solidity: function marketSellOrdersNoThrow([]LibOrderOrder orders, uint256 takerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenSession) MarketSellOrdersNoThrow(orders []LibOrderOrder, takerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.MarketSellOrdersNoThrow(&_Token.TransactOpts, orders, takerAssetFillAmount, signatures)
}

// MarketSellOrdersNoThrow is a paid mutator transaction binding the contract method 0x369da099.
//
// Solidity: function marketSellOrdersNoThrow([]LibOrderOrder orders, uint256 takerAssetFillAmount, bytes[] signatures) payable returns(LibFillResultsFillResults fillResults)
func (_Token *TokenTransactorSession) MarketSellOrdersNoThrow(orders []LibOrderOrder, takerAssetFillAmount *big.Int, signatures [][]byte) (*types.Transaction, error) {
	return _Token.Contract.MarketSellOrdersNoThrow(&_Token.TransactOpts, orders, takerAssetFillAmount, signatures)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders(LibOrderOrder leftOrder, LibOrderOrder rightOrder, bytes leftSignature, bytes rightSignature) payable returns(LibFillResultsMatchedFillResults matchedFillResults)
func (_Token *TokenTransactor) MatchOrders(opts *bind.TransactOpts, leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "matchOrders", leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders(LibOrderOrder leftOrder, LibOrderOrder rightOrder, bytes leftSignature, bytes rightSignature) payable returns(LibFillResultsMatchedFillResults matchedFillResults)
func (_Token *TokenSession) MatchOrders(leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _Token.Contract.MatchOrders(&_Token.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x88ec79fb.
//
// Solidity: function matchOrders(LibOrderOrder leftOrder, LibOrderOrder rightOrder, bytes leftSignature, bytes rightSignature) payable returns(LibFillResultsMatchedFillResults matchedFillResults)
func (_Token *TokenTransactorSession) MatchOrders(leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _Token.Contract.MatchOrders(&_Token.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrdersWithMaximalFill is a paid mutator transaction binding the contract method 0xb718e292.
//
// Solidity: function matchOrdersWithMaximalFill(LibOrderOrder leftOrder, LibOrderOrder rightOrder, bytes leftSignature, bytes rightSignature) payable returns(LibFillResultsMatchedFillResults matchedFillResults)
func (_Token *TokenTransactor) MatchOrdersWithMaximalFill(opts *bind.TransactOpts, leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "matchOrdersWithMaximalFill", leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrdersWithMaximalFill is a paid mutator transaction binding the contract method 0xb718e292.
//
// Solidity: function matchOrdersWithMaximalFill(LibOrderOrder leftOrder, LibOrderOrder rightOrder, bytes leftSignature, bytes rightSignature) payable returns(LibFillResultsMatchedFillResults matchedFillResults)
func (_Token *TokenSession) MatchOrdersWithMaximalFill(leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _Token.Contract.MatchOrdersWithMaximalFill(&_Token.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// MatchOrdersWithMaximalFill is a paid mutator transaction binding the contract method 0xb718e292.
//
// Solidity: function matchOrdersWithMaximalFill(LibOrderOrder leftOrder, LibOrderOrder rightOrder, bytes leftSignature, bytes rightSignature) payable returns(LibFillResultsMatchedFillResults matchedFillResults)
func (_Token *TokenTransactorSession) MatchOrdersWithMaximalFill(leftOrder LibOrderOrder, rightOrder LibOrderOrder, leftSignature []byte, rightSignature []byte) (*types.Transaction, error) {
	return _Token.Contract.MatchOrdersWithMaximalFill(&_Token.TransactOpts, leftOrder, rightOrder, leftSignature, rightSignature)
}

// PreSign is a paid mutator transaction binding the contract method 0x46c02d7a.
//
// Solidity: function preSign(bytes32 hash) payable returns()
func (_Token *TokenTransactor) PreSign(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "preSign", hash)
}

// PreSign is a paid mutator transaction binding the contract method 0x46c02d7a.
//
// Solidity: function preSign(bytes32 hash) payable returns()
func (_Token *TokenSession) PreSign(hash [32]byte) (*types.Transaction, error) {
	return _Token.Contract.PreSign(&_Token.TransactOpts, hash)
}

// PreSign is a paid mutator transaction binding the contract method 0x46c02d7a.
//
// Solidity: function preSign(bytes32 hash) payable returns()
func (_Token *TokenTransactorSession) PreSign(hash [32]byte) (*types.Transaction, error) {
	return _Token.Contract.PreSign(&_Token.TransactOpts, hash)
}

// RegisterAssetProxy is a paid mutator transaction binding the contract method 0xc585bb93.
//
// Solidity: function registerAssetProxy(address assetProxy) returns()
func (_Token *TokenTransactor) RegisterAssetProxy(opts *bind.TransactOpts, assetProxy common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "registerAssetProxy", assetProxy)
}

// RegisterAssetProxy is a paid mutator transaction binding the contract method 0xc585bb93.
//
// Solidity: function registerAssetProxy(address assetProxy) returns()
func (_Token *TokenSession) RegisterAssetProxy(assetProxy common.Address) (*types.Transaction, error) {
	return _Token.Contract.RegisterAssetProxy(&_Token.TransactOpts, assetProxy)
}

// RegisterAssetProxy is a paid mutator transaction binding the contract method 0xc585bb93.
//
// Solidity: function registerAssetProxy(address assetProxy) returns()
func (_Token *TokenTransactorSession) RegisterAssetProxy(assetProxy common.Address) (*types.Transaction, error) {
	return _Token.Contract.RegisterAssetProxy(&_Token.TransactOpts, assetProxy)
}

// SetProtocolFeeCollectorAddress is a paid mutator transaction binding the contract method 0xc0fa16cc.
//
// Solidity: function setProtocolFeeCollectorAddress(address updatedProtocolFeeCollector) returns()
func (_Token *TokenTransactor) SetProtocolFeeCollectorAddress(opts *bind.TransactOpts, updatedProtocolFeeCollector common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setProtocolFeeCollectorAddress", updatedProtocolFeeCollector)
}

// SetProtocolFeeCollectorAddress is a paid mutator transaction binding the contract method 0xc0fa16cc.
//
// Solidity: function setProtocolFeeCollectorAddress(address updatedProtocolFeeCollector) returns()
func (_Token *TokenSession) SetProtocolFeeCollectorAddress(updatedProtocolFeeCollector common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetProtocolFeeCollectorAddress(&_Token.TransactOpts, updatedProtocolFeeCollector)
}

// SetProtocolFeeCollectorAddress is a paid mutator transaction binding the contract method 0xc0fa16cc.
//
// Solidity: function setProtocolFeeCollectorAddress(address updatedProtocolFeeCollector) returns()
func (_Token *TokenTransactorSession) SetProtocolFeeCollectorAddress(updatedProtocolFeeCollector common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetProtocolFeeCollectorAddress(&_Token.TransactOpts, updatedProtocolFeeCollector)
}

// SetProtocolFeeMultiplier is a paid mutator transaction binding the contract method 0x9331c742.
//
// Solidity: function setProtocolFeeMultiplier(uint256 updatedProtocolFeeMultiplier) returns()
func (_Token *TokenTransactor) SetProtocolFeeMultiplier(opts *bind.TransactOpts, updatedProtocolFeeMultiplier *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setProtocolFeeMultiplier", updatedProtocolFeeMultiplier)
}

// SetProtocolFeeMultiplier is a paid mutator transaction binding the contract method 0x9331c742.
//
// Solidity: function setProtocolFeeMultiplier(uint256 updatedProtocolFeeMultiplier) returns()
func (_Token *TokenSession) SetProtocolFeeMultiplier(updatedProtocolFeeMultiplier *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetProtocolFeeMultiplier(&_Token.TransactOpts, updatedProtocolFeeMultiplier)
}

// SetProtocolFeeMultiplier is a paid mutator transaction binding the contract method 0x9331c742.
//
// Solidity: function setProtocolFeeMultiplier(uint256 updatedProtocolFeeMultiplier) returns()
func (_Token *TokenTransactorSession) SetProtocolFeeMultiplier(updatedProtocolFeeMultiplier *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetProtocolFeeMultiplier(&_Token.TransactOpts, updatedProtocolFeeMultiplier)
}

// SetSignatureValidatorApproval is a paid mutator transaction binding the contract method 0x77fcce68.
//
// Solidity: function setSignatureValidatorApproval(address validatorAddress, bool approval) payable returns()
func (_Token *TokenTransactor) SetSignatureValidatorApproval(opts *bind.TransactOpts, validatorAddress common.Address, approval bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setSignatureValidatorApproval", validatorAddress, approval)
}

// SetSignatureValidatorApproval is a paid mutator transaction binding the contract method 0x77fcce68.
//
// Solidity: function setSignatureValidatorApproval(address validatorAddress, bool approval) payable returns()
func (_Token *TokenSession) SetSignatureValidatorApproval(validatorAddress common.Address, approval bool) (*types.Transaction, error) {
	return _Token.Contract.SetSignatureValidatorApproval(&_Token.TransactOpts, validatorAddress, approval)
}

// SetSignatureValidatorApproval is a paid mutator transaction binding the contract method 0x77fcce68.
//
// Solidity: function setSignatureValidatorApproval(address validatorAddress, bool approval) payable returns()
func (_Token *TokenTransactorSession) SetSignatureValidatorApproval(validatorAddress common.Address, approval bool) (*types.Transaction, error) {
	return _Token.Contract.SetSignatureValidatorApproval(&_Token.TransactOpts, validatorAddress, approval)
}

// SimulateDispatchTransferFromCalls is a paid mutator transaction binding the contract method 0xb04fbddd.
//
// Solidity: function simulateDispatchTransferFromCalls(bytes[] assetData, address[] fromAddresses, address[] toAddresses, uint256[] amounts) returns()
func (_Token *TokenTransactor) SimulateDispatchTransferFromCalls(opts *bind.TransactOpts, assetData [][]byte, fromAddresses []common.Address, toAddresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "simulateDispatchTransferFromCalls", assetData, fromAddresses, toAddresses, amounts)
}

// SimulateDispatchTransferFromCalls is a paid mutator transaction binding the contract method 0xb04fbddd.
//
// Solidity: function simulateDispatchTransferFromCalls(bytes[] assetData, address[] fromAddresses, address[] toAddresses, uint256[] amounts) returns()
func (_Token *TokenSession) SimulateDispatchTransferFromCalls(assetData [][]byte, fromAddresses []common.Address, toAddresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.SimulateDispatchTransferFromCalls(&_Token.TransactOpts, assetData, fromAddresses, toAddresses, amounts)
}

// SimulateDispatchTransferFromCalls is a paid mutator transaction binding the contract method 0xb04fbddd.
//
// Solidity: function simulateDispatchTransferFromCalls(bytes[] assetData, address[] fromAddresses, address[] toAddresses, uint256[] amounts) returns()
func (_Token *TokenTransactorSession) SimulateDispatchTransferFromCalls(assetData [][]byte, fromAddresses []common.Address, toAddresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.SimulateDispatchTransferFromCalls(&_Token.TransactOpts, assetData, fromAddresses, toAddresses, amounts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TokenAssetProxyRegisteredIterator is returned from FilterAssetProxyRegistered and is used to iterate over the raw logs and unpacked data for AssetProxyRegistered events raised by the Token contract.
type TokenAssetProxyRegisteredIterator struct {
	Event *TokenAssetProxyRegistered // Event containing the contract specifics and raw log

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
func (it *TokenAssetProxyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAssetProxyRegistered)
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
		it.Event = new(TokenAssetProxyRegistered)
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
func (it *TokenAssetProxyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAssetProxyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAssetProxyRegistered represents a AssetProxyRegistered event raised by the Token contract.
type TokenAssetProxyRegistered struct {
	Id         [4]byte
	AssetProxy common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAssetProxyRegistered is a free log retrieval operation binding the contract event 0xd2c6b762299c609bdb96520b58a49bfb80186934d4f71a86a367571a15c03194.
//
// Solidity: event AssetProxyRegistered(bytes4 id, address assetProxy)
func (_Token *TokenFilterer) FilterAssetProxyRegistered(opts *bind.FilterOpts) (*TokenAssetProxyRegisteredIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "AssetProxyRegistered")
	if err != nil {
		return nil, err
	}
	return &TokenAssetProxyRegisteredIterator{contract: _Token.contract, event: "AssetProxyRegistered", logs: logs, sub: sub}, nil
}

// WatchAssetProxyRegistered is a free log subscription operation binding the contract event 0xd2c6b762299c609bdb96520b58a49bfb80186934d4f71a86a367571a15c03194.
//
// Solidity: event AssetProxyRegistered(bytes4 id, address assetProxy)
func (_Token *TokenFilterer) WatchAssetProxyRegistered(opts *bind.WatchOpts, sink chan<- *TokenAssetProxyRegistered) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "AssetProxyRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAssetProxyRegistered)
				if err := _Token.contract.UnpackLog(event, "AssetProxyRegistered", log); err != nil {
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

// ParseAssetProxyRegistered is a log parse operation binding the contract event 0xd2c6b762299c609bdb96520b58a49bfb80186934d4f71a86a367571a15c03194.
//
// Solidity: event AssetProxyRegistered(bytes4 id, address assetProxy)
func (_Token *TokenFilterer) ParseAssetProxyRegistered(log types.Log) (*TokenAssetProxyRegistered, error) {
	event := new(TokenAssetProxyRegistered)
	if err := _Token.contract.UnpackLog(event, "AssetProxyRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenCancelIterator is returned from FilterCancel and is used to iterate over the raw logs and unpacked data for Cancel events raised by the Token contract.
type TokenCancelIterator struct {
	Event *TokenCancel // Event containing the contract specifics and raw log

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
func (it *TokenCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCancel)
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
		it.Event = new(TokenCancel)
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
func (it *TokenCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCancel represents a Cancel event raised by the Token contract.
type TokenCancel struct {
	MakerAddress        common.Address
	FeeRecipientAddress common.Address
	MakerAssetData      []byte
	TakerAssetData      []byte
	SenderAddress       common.Address
	OrderHash           [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCancel is a free log retrieval operation binding the contract event 0x02c310a9a43963ff31a754a4099cc435ed498049687539d72d7818d9b093415c.
//
// Solidity: event Cancel(address indexed makerAddress, address indexed feeRecipientAddress, bytes makerAssetData, bytes takerAssetData, address senderAddress, bytes32 indexed orderHash)
func (_Token *TokenFilterer) FilterCancel(opts *bind.FilterOpts, makerAddress []common.Address, feeRecipientAddress []common.Address, orderHash [][32]byte) (*TokenCancelIterator, error) {

	var makerAddressRule []interface{}
	for _, makerAddressItem := range makerAddress {
		makerAddressRule = append(makerAddressRule, makerAddressItem)
	}
	var feeRecipientAddressRule []interface{}
	for _, feeRecipientAddressItem := range feeRecipientAddress {
		feeRecipientAddressRule = append(feeRecipientAddressRule, feeRecipientAddressItem)
	}

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Cancel", makerAddressRule, feeRecipientAddressRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &TokenCancelIterator{contract: _Token.contract, event: "Cancel", logs: logs, sub: sub}, nil
}

// WatchCancel is a free log subscription operation binding the contract event 0x02c310a9a43963ff31a754a4099cc435ed498049687539d72d7818d9b093415c.
//
// Solidity: event Cancel(address indexed makerAddress, address indexed feeRecipientAddress, bytes makerAssetData, bytes takerAssetData, address senderAddress, bytes32 indexed orderHash)
func (_Token *TokenFilterer) WatchCancel(opts *bind.WatchOpts, sink chan<- *TokenCancel, makerAddress []common.Address, feeRecipientAddress []common.Address, orderHash [][32]byte) (event.Subscription, error) {

	var makerAddressRule []interface{}
	for _, makerAddressItem := range makerAddress {
		makerAddressRule = append(makerAddressRule, makerAddressItem)
	}
	var feeRecipientAddressRule []interface{}
	for _, feeRecipientAddressItem := range feeRecipientAddress {
		feeRecipientAddressRule = append(feeRecipientAddressRule, feeRecipientAddressItem)
	}

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Cancel", makerAddressRule, feeRecipientAddressRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCancel)
				if err := _Token.contract.UnpackLog(event, "Cancel", log); err != nil {
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

// ParseCancel is a log parse operation binding the contract event 0x02c310a9a43963ff31a754a4099cc435ed498049687539d72d7818d9b093415c.
//
// Solidity: event Cancel(address indexed makerAddress, address indexed feeRecipientAddress, bytes makerAssetData, bytes takerAssetData, address senderAddress, bytes32 indexed orderHash)
func (_Token *TokenFilterer) ParseCancel(log types.Log) (*TokenCancel, error) {
	event := new(TokenCancel)
	if err := _Token.contract.UnpackLog(event, "Cancel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenCancelUpToIterator is returned from FilterCancelUpTo and is used to iterate over the raw logs and unpacked data for CancelUpTo events raised by the Token contract.
type TokenCancelUpToIterator struct {
	Event *TokenCancelUpTo // Event containing the contract specifics and raw log

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
func (it *TokenCancelUpToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCancelUpTo)
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
		it.Event = new(TokenCancelUpTo)
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
func (it *TokenCancelUpToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCancelUpToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCancelUpTo represents a CancelUpTo event raised by the Token contract.
type TokenCancelUpTo struct {
	MakerAddress       common.Address
	OrderSenderAddress common.Address
	OrderEpoch         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCancelUpTo is a free log retrieval operation binding the contract event 0x82af639571738f4ebd4268fb0363d8957ebe1bbb9e78dba5ebd69eed39b154f0.
//
// Solidity: event CancelUpTo(address indexed makerAddress, address indexed orderSenderAddress, uint256 orderEpoch)
func (_Token *TokenFilterer) FilterCancelUpTo(opts *bind.FilterOpts, makerAddress []common.Address, orderSenderAddress []common.Address) (*TokenCancelUpToIterator, error) {

	var makerAddressRule []interface{}
	for _, makerAddressItem := range makerAddress {
		makerAddressRule = append(makerAddressRule, makerAddressItem)
	}
	var orderSenderAddressRule []interface{}
	for _, orderSenderAddressItem := range orderSenderAddress {
		orderSenderAddressRule = append(orderSenderAddressRule, orderSenderAddressItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "CancelUpTo", makerAddressRule, orderSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenCancelUpToIterator{contract: _Token.contract, event: "CancelUpTo", logs: logs, sub: sub}, nil
}

// WatchCancelUpTo is a free log subscription operation binding the contract event 0x82af639571738f4ebd4268fb0363d8957ebe1bbb9e78dba5ebd69eed39b154f0.
//
// Solidity: event CancelUpTo(address indexed makerAddress, address indexed orderSenderAddress, uint256 orderEpoch)
func (_Token *TokenFilterer) WatchCancelUpTo(opts *bind.WatchOpts, sink chan<- *TokenCancelUpTo, makerAddress []common.Address, orderSenderAddress []common.Address) (event.Subscription, error) {

	var makerAddressRule []interface{}
	for _, makerAddressItem := range makerAddress {
		makerAddressRule = append(makerAddressRule, makerAddressItem)
	}
	var orderSenderAddressRule []interface{}
	for _, orderSenderAddressItem := range orderSenderAddress {
		orderSenderAddressRule = append(orderSenderAddressRule, orderSenderAddressItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "CancelUpTo", makerAddressRule, orderSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCancelUpTo)
				if err := _Token.contract.UnpackLog(event, "CancelUpTo", log); err != nil {
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

// ParseCancelUpTo is a log parse operation binding the contract event 0x82af639571738f4ebd4268fb0363d8957ebe1bbb9e78dba5ebd69eed39b154f0.
//
// Solidity: event CancelUpTo(address indexed makerAddress, address indexed orderSenderAddress, uint256 orderEpoch)
func (_Token *TokenFilterer) ParseCancelUpTo(log types.Log) (*TokenCancelUpTo, error) {
	event := new(TokenCancelUpTo)
	if err := _Token.contract.UnpackLog(event, "CancelUpTo", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenFillIterator is returned from FilterFill and is used to iterate over the raw logs and unpacked data for Fill events raised by the Token contract.
type TokenFillIterator struct {
	Event *TokenFill // Event containing the contract specifics and raw log

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
func (it *TokenFillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFill)
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
		it.Event = new(TokenFill)
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
func (it *TokenFillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFill represents a Fill event raised by the Token contract.
type TokenFill struct {
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

// FilterFill is a free log retrieval operation binding the contract event 0x6869791f0a34781b29882982cc39e882768cf2c96995c2a110c577c53bc932d5.
//
// Solidity: event Fill(address indexed makerAddress, address indexed feeRecipientAddress, bytes makerAssetData, bytes takerAssetData, bytes makerFeeAssetData, bytes takerFeeAssetData, bytes32 indexed orderHash, address takerAddress, address senderAddress, uint256 makerAssetFilledAmount, uint256 takerAssetFilledAmount, uint256 makerFeePaid, uint256 takerFeePaid, uint256 protocolFeePaid)
func (_Token *TokenFilterer) FilterFill(opts *bind.FilterOpts, makerAddress []common.Address, feeRecipientAddress []common.Address, orderHash [][32]byte) (*TokenFillIterator, error) {

	var makerAddressRule []interface{}
	for _, makerAddressItem := range makerAddress {
		makerAddressRule = append(makerAddressRule, makerAddressItem)
	}
	var feeRecipientAddressRule []interface{}
	for _, feeRecipientAddressItem := range feeRecipientAddress {
		feeRecipientAddressRule = append(feeRecipientAddressRule, feeRecipientAddressItem)
	}

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Fill", makerAddressRule, feeRecipientAddressRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &TokenFillIterator{contract: _Token.contract, event: "Fill", logs: logs, sub: sub}, nil
}

// WatchFill is a free log subscription operation binding the contract event 0x6869791f0a34781b29882982cc39e882768cf2c96995c2a110c577c53bc932d5.
//
// Solidity: event Fill(address indexed makerAddress, address indexed feeRecipientAddress, bytes makerAssetData, bytes takerAssetData, bytes makerFeeAssetData, bytes takerFeeAssetData, bytes32 indexed orderHash, address takerAddress, address senderAddress, uint256 makerAssetFilledAmount, uint256 takerAssetFilledAmount, uint256 makerFeePaid, uint256 takerFeePaid, uint256 protocolFeePaid)
func (_Token *TokenFilterer) WatchFill(opts *bind.WatchOpts, sink chan<- *TokenFill, makerAddress []common.Address, feeRecipientAddress []common.Address, orderHash [][32]byte) (event.Subscription, error) {

	var makerAddressRule []interface{}
	for _, makerAddressItem := range makerAddress {
		makerAddressRule = append(makerAddressRule, makerAddressItem)
	}
	var feeRecipientAddressRule []interface{}
	for _, feeRecipientAddressItem := range feeRecipientAddress {
		feeRecipientAddressRule = append(feeRecipientAddressRule, feeRecipientAddressItem)
	}

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Fill", makerAddressRule, feeRecipientAddressRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFill)
				if err := _Token.contract.UnpackLog(event, "Fill", log); err != nil {
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

// ParseFill is a log parse operation binding the contract event 0x6869791f0a34781b29882982cc39e882768cf2c96995c2a110c577c53bc932d5.
//
// Solidity: event Fill(address indexed makerAddress, address indexed feeRecipientAddress, bytes makerAssetData, bytes takerAssetData, bytes makerFeeAssetData, bytes takerFeeAssetData, bytes32 indexed orderHash, address takerAddress, address senderAddress, uint256 makerAssetFilledAmount, uint256 takerAssetFilledAmount, uint256 makerFeePaid, uint256 takerFeePaid, uint256 protocolFeePaid)
func (_Token *TokenFilterer) ParseFill(log types.Log) (*TokenFill, error) {
	event := new(TokenFill)
	if err := _Token.contract.UnpackLog(event, "Fill", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token contract.
type TokenOwnershipTransferredIterator struct {
	Event *TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnershipTransferred)
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
		it.Event = new(TokenOwnershipTransferred)
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
func (it *TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnershipTransferred represents a OwnershipTransferred event raised by the Token contract.
type TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnershipTransferredIterator{contract: _Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnershipTransferred)
				if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) ParseOwnershipTransferred(log types.Log) (*TokenOwnershipTransferred, error) {
	event := new(TokenOwnershipTransferred)
	if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenProtocolFeeCollectorAddressIterator is returned from FilterProtocolFeeCollectorAddress and is used to iterate over the raw logs and unpacked data for ProtocolFeeCollectorAddress events raised by the Token contract.
type TokenProtocolFeeCollectorAddressIterator struct {
	Event *TokenProtocolFeeCollectorAddress // Event containing the contract specifics and raw log

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
func (it *TokenProtocolFeeCollectorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenProtocolFeeCollectorAddress)
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
		it.Event = new(TokenProtocolFeeCollectorAddress)
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
func (it *TokenProtocolFeeCollectorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenProtocolFeeCollectorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenProtocolFeeCollectorAddress represents a ProtocolFeeCollectorAddress event raised by the Token contract.
type TokenProtocolFeeCollectorAddress struct {
	OldProtocolFeeCollector     common.Address
	UpdatedProtocolFeeCollector common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeCollectorAddress is a free log retrieval operation binding the contract event 0xe1a5430ebec577336427f40f15822f1f36c5e3509ff209d6db9e6c9e6941cb0b.
//
// Solidity: event ProtocolFeeCollectorAddress(address oldProtocolFeeCollector, address updatedProtocolFeeCollector)
func (_Token *TokenFilterer) FilterProtocolFeeCollectorAddress(opts *bind.FilterOpts) (*TokenProtocolFeeCollectorAddressIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "ProtocolFeeCollectorAddress")
	if err != nil {
		return nil, err
	}
	return &TokenProtocolFeeCollectorAddressIterator{contract: _Token.contract, event: "ProtocolFeeCollectorAddress", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeCollectorAddress is a free log subscription operation binding the contract event 0xe1a5430ebec577336427f40f15822f1f36c5e3509ff209d6db9e6c9e6941cb0b.
//
// Solidity: event ProtocolFeeCollectorAddress(address oldProtocolFeeCollector, address updatedProtocolFeeCollector)
func (_Token *TokenFilterer) WatchProtocolFeeCollectorAddress(opts *bind.WatchOpts, sink chan<- *TokenProtocolFeeCollectorAddress) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "ProtocolFeeCollectorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenProtocolFeeCollectorAddress)
				if err := _Token.contract.UnpackLog(event, "ProtocolFeeCollectorAddress", log); err != nil {
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

// ParseProtocolFeeCollectorAddress is a log parse operation binding the contract event 0xe1a5430ebec577336427f40f15822f1f36c5e3509ff209d6db9e6c9e6941cb0b.
//
// Solidity: event ProtocolFeeCollectorAddress(address oldProtocolFeeCollector, address updatedProtocolFeeCollector)
func (_Token *TokenFilterer) ParseProtocolFeeCollectorAddress(log types.Log) (*TokenProtocolFeeCollectorAddress, error) {
	event := new(TokenProtocolFeeCollectorAddress)
	if err := _Token.contract.UnpackLog(event, "ProtocolFeeCollectorAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenProtocolFeeMultiplierIterator is returned from FilterProtocolFeeMultiplier and is used to iterate over the raw logs and unpacked data for ProtocolFeeMultiplier events raised by the Token contract.
type TokenProtocolFeeMultiplierIterator struct {
	Event *TokenProtocolFeeMultiplier // Event containing the contract specifics and raw log

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
func (it *TokenProtocolFeeMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenProtocolFeeMultiplier)
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
		it.Event = new(TokenProtocolFeeMultiplier)
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
func (it *TokenProtocolFeeMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenProtocolFeeMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenProtocolFeeMultiplier represents a ProtocolFeeMultiplier event raised by the Token contract.
type TokenProtocolFeeMultiplier struct {
	OldProtocolFeeMultiplier     *big.Int
	UpdatedProtocolFeeMultiplier *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeMultiplier is a free log retrieval operation binding the contract event 0x3a3e76d7a75e198aef1f53137e4f2a8a2ec74e2e9526db8404d08ccc9f1e621d.
//
// Solidity: event ProtocolFeeMultiplier(uint256 oldProtocolFeeMultiplier, uint256 updatedProtocolFeeMultiplier)
func (_Token *TokenFilterer) FilterProtocolFeeMultiplier(opts *bind.FilterOpts) (*TokenProtocolFeeMultiplierIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "ProtocolFeeMultiplier")
	if err != nil {
		return nil, err
	}
	return &TokenProtocolFeeMultiplierIterator{contract: _Token.contract, event: "ProtocolFeeMultiplier", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeMultiplier is a free log subscription operation binding the contract event 0x3a3e76d7a75e198aef1f53137e4f2a8a2ec74e2e9526db8404d08ccc9f1e621d.
//
// Solidity: event ProtocolFeeMultiplier(uint256 oldProtocolFeeMultiplier, uint256 updatedProtocolFeeMultiplier)
func (_Token *TokenFilterer) WatchProtocolFeeMultiplier(opts *bind.WatchOpts, sink chan<- *TokenProtocolFeeMultiplier) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "ProtocolFeeMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenProtocolFeeMultiplier)
				if err := _Token.contract.UnpackLog(event, "ProtocolFeeMultiplier", log); err != nil {
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

// ParseProtocolFeeMultiplier is a log parse operation binding the contract event 0x3a3e76d7a75e198aef1f53137e4f2a8a2ec74e2e9526db8404d08ccc9f1e621d.
//
// Solidity: event ProtocolFeeMultiplier(uint256 oldProtocolFeeMultiplier, uint256 updatedProtocolFeeMultiplier)
func (_Token *TokenFilterer) ParseProtocolFeeMultiplier(log types.Log) (*TokenProtocolFeeMultiplier, error) {
	event := new(TokenProtocolFeeMultiplier)
	if err := _Token.contract.UnpackLog(event, "ProtocolFeeMultiplier", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenSignatureValidatorApprovalIterator is returned from FilterSignatureValidatorApproval and is used to iterate over the raw logs and unpacked data for SignatureValidatorApproval events raised by the Token contract.
type TokenSignatureValidatorApprovalIterator struct {
	Event *TokenSignatureValidatorApproval // Event containing the contract specifics and raw log

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
func (it *TokenSignatureValidatorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSignatureValidatorApproval)
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
		it.Event = new(TokenSignatureValidatorApproval)
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
func (it *TokenSignatureValidatorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSignatureValidatorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSignatureValidatorApproval represents a SignatureValidatorApproval event raised by the Token contract.
type TokenSignatureValidatorApproval struct {
	SignerAddress    common.Address
	ValidatorAddress common.Address
	IsApproved       bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSignatureValidatorApproval is a free log retrieval operation binding the contract event 0xa8656e308026eeabce8f0bc18048433252318ab80ac79da0b3d3d8697dfba891.
//
// Solidity: event SignatureValidatorApproval(address indexed signerAddress, address indexed validatorAddress, bool isApproved)
func (_Token *TokenFilterer) FilterSignatureValidatorApproval(opts *bind.FilterOpts, signerAddress []common.Address, validatorAddress []common.Address) (*TokenSignatureValidatorApprovalIterator, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SignatureValidatorApproval", signerAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenSignatureValidatorApprovalIterator{contract: _Token.contract, event: "SignatureValidatorApproval", logs: logs, sub: sub}, nil
}

// WatchSignatureValidatorApproval is a free log subscription operation binding the contract event 0xa8656e308026eeabce8f0bc18048433252318ab80ac79da0b3d3d8697dfba891.
//
// Solidity: event SignatureValidatorApproval(address indexed signerAddress, address indexed validatorAddress, bool isApproved)
func (_Token *TokenFilterer) WatchSignatureValidatorApproval(opts *bind.WatchOpts, sink chan<- *TokenSignatureValidatorApproval, signerAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SignatureValidatorApproval", signerAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSignatureValidatorApproval)
				if err := _Token.contract.UnpackLog(event, "SignatureValidatorApproval", log); err != nil {
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

// ParseSignatureValidatorApproval is a log parse operation binding the contract event 0xa8656e308026eeabce8f0bc18048433252318ab80ac79da0b3d3d8697dfba891.
//
// Solidity: event SignatureValidatorApproval(address indexed signerAddress, address indexed validatorAddress, bool isApproved)
func (_Token *TokenFilterer) ParseSignatureValidatorApproval(log types.Log) (*TokenSignatureValidatorApproval, error) {
	event := new(TokenSignatureValidatorApproval)
	if err := _Token.contract.UnpackLog(event, "SignatureValidatorApproval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenTransactionExecutionIterator is returned from FilterTransactionExecution and is used to iterate over the raw logs and unpacked data for TransactionExecution events raised by the Token contract.
type TokenTransactionExecutionIterator struct {
	Event *TokenTransactionExecution // Event containing the contract specifics and raw log

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
func (it *TokenTransactionExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransactionExecution)
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
		it.Event = new(TokenTransactionExecution)
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
func (it *TokenTransactionExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransactionExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransactionExecution represents a TransactionExecution event raised by the Token contract.
type TokenTransactionExecution struct {
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransactionExecution is a free log retrieval operation binding the contract event 0xa4a7329f1dd821363067e07d359e347b4af9b1efe4b6cccf13240228af3c800d.
//
// Solidity: event TransactionExecution(bytes32 indexed transactionHash)
func (_Token *TokenFilterer) FilterTransactionExecution(opts *bind.FilterOpts, transactionHash [][32]byte) (*TokenTransactionExecutionIterator, error) {

	var transactionHashRule []interface{}
	for _, transactionHashItem := range transactionHash {
		transactionHashRule = append(transactionHashRule, transactionHashItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "TransactionExecution", transactionHashRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransactionExecutionIterator{contract: _Token.contract, event: "TransactionExecution", logs: logs, sub: sub}, nil
}

// WatchTransactionExecution is a free log subscription operation binding the contract event 0xa4a7329f1dd821363067e07d359e347b4af9b1efe4b6cccf13240228af3c800d.
//
// Solidity: event TransactionExecution(bytes32 indexed transactionHash)
func (_Token *TokenFilterer) WatchTransactionExecution(opts *bind.WatchOpts, sink chan<- *TokenTransactionExecution, transactionHash [][32]byte) (event.Subscription, error) {

	var transactionHashRule []interface{}
	for _, transactionHashItem := range transactionHash {
		transactionHashRule = append(transactionHashRule, transactionHashItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "TransactionExecution", transactionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransactionExecution)
				if err := _Token.contract.UnpackLog(event, "TransactionExecution", log); err != nil {
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

// ParseTransactionExecution is a log parse operation binding the contract event 0xa4a7329f1dd821363067e07d359e347b4af9b1efe4b6cccf13240228af3c800d.
//
// Solidity: event TransactionExecution(bytes32 indexed transactionHash)
func (_Token *TokenFilterer) ParseTransactionExecution(log types.Log) (*TokenTransactionExecution, error) {
	event := new(TokenTransactionExecution)
	if err := _Token.contract.UnpackLog(event, "TransactionExecution", log); err != nil {
		return nil, err
	}
	return event, nil
}
