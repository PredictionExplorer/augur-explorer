// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// ISwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// ISwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// ISwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// BlockTimestampMetaData contains all meta data concerning the BlockTimestamp contract.
var BlockTimestampMetaData = &bind.MetaData{
	ABI: "[]",
}

// BlockTimestampABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockTimestampMetaData.ABI instead.
var BlockTimestampABI = BlockTimestampMetaData.ABI

// BlockTimestamp is an auto generated Go binding around an Ethereum contract.
type BlockTimestamp struct {
	BlockTimestampCaller     // Read-only binding to the contract
	BlockTimestampTransactor // Write-only binding to the contract
	BlockTimestampFilterer   // Log filterer for contract events
}

// BlockTimestampCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockTimestampCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockTimestampTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockTimestampTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockTimestampFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockTimestampFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockTimestampSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockTimestampSession struct {
	Contract     *BlockTimestamp   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockTimestampCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockTimestampCallerSession struct {
	Contract *BlockTimestampCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BlockTimestampTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockTimestampTransactorSession struct {
	Contract     *BlockTimestampTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BlockTimestampRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockTimestampRaw struct {
	Contract *BlockTimestamp // Generic contract binding to access the raw methods on
}

// BlockTimestampCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockTimestampCallerRaw struct {
	Contract *BlockTimestampCaller // Generic read-only contract binding to access the raw methods on
}

// BlockTimestampTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockTimestampTransactorRaw struct {
	Contract *BlockTimestampTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockTimestamp creates a new instance of BlockTimestamp, bound to a specific deployed contract.
func NewBlockTimestamp(address common.Address, backend bind.ContractBackend) (*BlockTimestamp, error) {
	contract, err := bindBlockTimestamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockTimestamp{BlockTimestampCaller: BlockTimestampCaller{contract: contract}, BlockTimestampTransactor: BlockTimestampTransactor{contract: contract}, BlockTimestampFilterer: BlockTimestampFilterer{contract: contract}}, nil
}

// NewBlockTimestampCaller creates a new read-only instance of BlockTimestamp, bound to a specific deployed contract.
func NewBlockTimestampCaller(address common.Address, caller bind.ContractCaller) (*BlockTimestampCaller, error) {
	contract, err := bindBlockTimestamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockTimestampCaller{contract: contract}, nil
}

// NewBlockTimestampTransactor creates a new write-only instance of BlockTimestamp, bound to a specific deployed contract.
func NewBlockTimestampTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockTimestampTransactor, error) {
	contract, err := bindBlockTimestamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockTimestampTransactor{contract: contract}, nil
}

// NewBlockTimestampFilterer creates a new log filterer instance of BlockTimestamp, bound to a specific deployed contract.
func NewBlockTimestampFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockTimestampFilterer, error) {
	contract, err := bindBlockTimestamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockTimestampFilterer{contract: contract}, nil
}

// bindBlockTimestamp binds a generic wrapper to an already deployed contract.
func bindBlockTimestamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockTimestampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockTimestamp *BlockTimestampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockTimestamp.Contract.BlockTimestampCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockTimestamp *BlockTimestampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockTimestamp.Contract.BlockTimestampTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockTimestamp *BlockTimestampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockTimestamp.Contract.BlockTimestampTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockTimestamp *BlockTimestampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockTimestamp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockTimestamp *BlockTimestampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockTimestamp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockTimestamp *BlockTimestampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockTimestamp.Contract.contract.Transact(opts, method, params...)
}

// BytesLibMetaData contains all meta data concerning the BytesLib contract.
var BytesLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a961f87210d4d817bc3a750c1c33d1d22e25c0c4db5ae52e64207d4f3ae70c0464736f6c63430007060033",
}

// BytesLibABI is the input ABI used to generate the binding from.
// Deprecated: Use BytesLibMetaData.ABI instead.
var BytesLibABI = BytesLibMetaData.ABI

// BytesLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BytesLibMetaData.Bin instead.
var BytesLibBin = BytesLibMetaData.Bin

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := BytesLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// CallbackValidationMetaData contains all meta data concerning the CallbackValidation contract.
var CallbackValidationMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220009c4a187899687f5d0a13598b5a1ff8064087b5020d8bde497e43097ea0b59b64736f6c63430007060033",
}

// CallbackValidationABI is the input ABI used to generate the binding from.
// Deprecated: Use CallbackValidationMetaData.ABI instead.
var CallbackValidationABI = CallbackValidationMetaData.ABI

// CallbackValidationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CallbackValidationMetaData.Bin instead.
var CallbackValidationBin = CallbackValidationMetaData.Bin

// DeployCallbackValidation deploys a new Ethereum contract, binding an instance of CallbackValidation to it.
func DeployCallbackValidation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CallbackValidation, error) {
	parsed, err := CallbackValidationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CallbackValidationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CallbackValidation{CallbackValidationCaller: CallbackValidationCaller{contract: contract}, CallbackValidationTransactor: CallbackValidationTransactor{contract: contract}, CallbackValidationFilterer: CallbackValidationFilterer{contract: contract}}, nil
}

// CallbackValidation is an auto generated Go binding around an Ethereum contract.
type CallbackValidation struct {
	CallbackValidationCaller     // Read-only binding to the contract
	CallbackValidationTransactor // Write-only binding to the contract
	CallbackValidationFilterer   // Log filterer for contract events
}

// CallbackValidationCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallbackValidationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallbackValidationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallbackValidationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallbackValidationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallbackValidationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallbackValidationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallbackValidationSession struct {
	Contract     *CallbackValidation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CallbackValidationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallbackValidationCallerSession struct {
	Contract *CallbackValidationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CallbackValidationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallbackValidationTransactorSession struct {
	Contract     *CallbackValidationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CallbackValidationRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallbackValidationRaw struct {
	Contract *CallbackValidation // Generic contract binding to access the raw methods on
}

// CallbackValidationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallbackValidationCallerRaw struct {
	Contract *CallbackValidationCaller // Generic read-only contract binding to access the raw methods on
}

// CallbackValidationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallbackValidationTransactorRaw struct {
	Contract *CallbackValidationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallbackValidation creates a new instance of CallbackValidation, bound to a specific deployed contract.
func NewCallbackValidation(address common.Address, backend bind.ContractBackend) (*CallbackValidation, error) {
	contract, err := bindCallbackValidation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CallbackValidation{CallbackValidationCaller: CallbackValidationCaller{contract: contract}, CallbackValidationTransactor: CallbackValidationTransactor{contract: contract}, CallbackValidationFilterer: CallbackValidationFilterer{contract: contract}}, nil
}

// NewCallbackValidationCaller creates a new read-only instance of CallbackValidation, bound to a specific deployed contract.
func NewCallbackValidationCaller(address common.Address, caller bind.ContractCaller) (*CallbackValidationCaller, error) {
	contract, err := bindCallbackValidation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallbackValidationCaller{contract: contract}, nil
}

// NewCallbackValidationTransactor creates a new write-only instance of CallbackValidation, bound to a specific deployed contract.
func NewCallbackValidationTransactor(address common.Address, transactor bind.ContractTransactor) (*CallbackValidationTransactor, error) {
	contract, err := bindCallbackValidation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallbackValidationTransactor{contract: contract}, nil
}

// NewCallbackValidationFilterer creates a new log filterer instance of CallbackValidation, bound to a specific deployed contract.
func NewCallbackValidationFilterer(address common.Address, filterer bind.ContractFilterer) (*CallbackValidationFilterer, error) {
	contract, err := bindCallbackValidation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallbackValidationFilterer{contract: contract}, nil
}

// bindCallbackValidation binds a generic wrapper to an already deployed contract.
func bindCallbackValidation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CallbackValidationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallbackValidation *CallbackValidationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallbackValidation.Contract.CallbackValidationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallbackValidation *CallbackValidationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallbackValidation.Contract.CallbackValidationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallbackValidation *CallbackValidationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallbackValidation.Contract.CallbackValidationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallbackValidation *CallbackValidationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallbackValidation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallbackValidation *CallbackValidationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallbackValidation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallbackValidation *CallbackValidationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallbackValidation.Contract.contract.Transact(opts, method, params...)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20PermitMetaData contains all meta data concerning the IERC20Permit contract.
var IERC20PermitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20PermitABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20PermitMetaData.ABI instead.
var IERC20PermitABI = IERC20PermitMetaData.ABI

// IERC20Permit is an auto generated Go binding around an Ethereum contract.
type IERC20Permit struct {
	IERC20PermitCaller     // Read-only binding to the contract
	IERC20PermitTransactor // Write-only binding to the contract
	IERC20PermitFilterer   // Log filterer for contract events
}

// IERC20PermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20PermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20PermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20PermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20PermitSession struct {
	Contract     *IERC20Permit     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20PermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20PermitCallerSession struct {
	Contract *IERC20PermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IERC20PermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20PermitTransactorSession struct {
	Contract     *IERC20PermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20PermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20PermitRaw struct {
	Contract *IERC20Permit // Generic contract binding to access the raw methods on
}

// IERC20PermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20PermitCallerRaw struct {
	Contract *IERC20PermitCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20PermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20PermitTransactorRaw struct {
	Contract *IERC20PermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Permit creates a new instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20Permit(address common.Address, backend bind.ContractBackend) (*IERC20Permit, error) {
	contract, err := bindIERC20Permit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Permit{IERC20PermitCaller: IERC20PermitCaller{contract: contract}, IERC20PermitTransactor: IERC20PermitTransactor{contract: contract}, IERC20PermitFilterer: IERC20PermitFilterer{contract: contract}}, nil
}

// NewIERC20PermitCaller creates a new read-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitCaller(address common.Address, caller bind.ContractCaller) (*IERC20PermitCaller, error) {
	contract, err := bindIERC20Permit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitCaller{contract: contract}, nil
}

// NewIERC20PermitTransactor creates a new write-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20PermitTransactor, error) {
	contract, err := bindIERC20Permit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitTransactor{contract: contract}, nil
}

// NewIERC20PermitFilterer creates a new log filterer instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20PermitFilterer, error) {
	contract, err := bindIERC20Permit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitFilterer{contract: contract}, nil
}

// bindIERC20Permit binds a generic wrapper to an already deployed contract.
func bindIERC20Permit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20PermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.IERC20PermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// IERC20PermitAllowedMetaData contains all meta data concerning the IERC20PermitAllowed contract.
var IERC20PermitAllowedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20PermitAllowedABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20PermitAllowedMetaData.ABI instead.
var IERC20PermitAllowedABI = IERC20PermitAllowedMetaData.ABI

// IERC20PermitAllowed is an auto generated Go binding around an Ethereum contract.
type IERC20PermitAllowed struct {
	IERC20PermitAllowedCaller     // Read-only binding to the contract
	IERC20PermitAllowedTransactor // Write-only binding to the contract
	IERC20PermitAllowedFilterer   // Log filterer for contract events
}

// IERC20PermitAllowedCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20PermitAllowedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitAllowedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20PermitAllowedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitAllowedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20PermitAllowedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitAllowedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20PermitAllowedSession struct {
	Contract     *IERC20PermitAllowed // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IERC20PermitAllowedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20PermitAllowedCallerSession struct {
	Contract *IERC20PermitAllowedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IERC20PermitAllowedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20PermitAllowedTransactorSession struct {
	Contract     *IERC20PermitAllowedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IERC20PermitAllowedRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20PermitAllowedRaw struct {
	Contract *IERC20PermitAllowed // Generic contract binding to access the raw methods on
}

// IERC20PermitAllowedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20PermitAllowedCallerRaw struct {
	Contract *IERC20PermitAllowedCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20PermitAllowedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20PermitAllowedTransactorRaw struct {
	Contract *IERC20PermitAllowedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20PermitAllowed creates a new instance of IERC20PermitAllowed, bound to a specific deployed contract.
func NewIERC20PermitAllowed(address common.Address, backend bind.ContractBackend) (*IERC20PermitAllowed, error) {
	contract, err := bindIERC20PermitAllowed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitAllowed{IERC20PermitAllowedCaller: IERC20PermitAllowedCaller{contract: contract}, IERC20PermitAllowedTransactor: IERC20PermitAllowedTransactor{contract: contract}, IERC20PermitAllowedFilterer: IERC20PermitAllowedFilterer{contract: contract}}, nil
}

// NewIERC20PermitAllowedCaller creates a new read-only instance of IERC20PermitAllowed, bound to a specific deployed contract.
func NewIERC20PermitAllowedCaller(address common.Address, caller bind.ContractCaller) (*IERC20PermitAllowedCaller, error) {
	contract, err := bindIERC20PermitAllowed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitAllowedCaller{contract: contract}, nil
}

// NewIERC20PermitAllowedTransactor creates a new write-only instance of IERC20PermitAllowed, bound to a specific deployed contract.
func NewIERC20PermitAllowedTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20PermitAllowedTransactor, error) {
	contract, err := bindIERC20PermitAllowed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitAllowedTransactor{contract: contract}, nil
}

// NewIERC20PermitAllowedFilterer creates a new log filterer instance of IERC20PermitAllowed, bound to a specific deployed contract.
func NewIERC20PermitAllowedFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20PermitAllowedFilterer, error) {
	contract, err := bindIERC20PermitAllowed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitAllowedFilterer{contract: contract}, nil
}

// bindIERC20PermitAllowed binds a generic wrapper to an already deployed contract.
func bindIERC20PermitAllowed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20PermitAllowedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20PermitAllowed *IERC20PermitAllowedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20PermitAllowed.Contract.IERC20PermitAllowedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20PermitAllowed *IERC20PermitAllowedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20PermitAllowed.Contract.IERC20PermitAllowedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20PermitAllowed *IERC20PermitAllowedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20PermitAllowed.Contract.IERC20PermitAllowedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20PermitAllowed *IERC20PermitAllowedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20PermitAllowed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20PermitAllowed *IERC20PermitAllowedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20PermitAllowed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20PermitAllowed *IERC20PermitAllowedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20PermitAllowed.Contract.contract.Transact(opts, method, params...)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20PermitAllowed *IERC20PermitAllowedTransactor) Permit(opts *bind.TransactOpts, holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20PermitAllowed.contract.Transact(opts, "permit", holder, spender, nonce, expiry, allowed, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20PermitAllowed *IERC20PermitAllowedSession) Permit(holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20PermitAllowed.Contract.Permit(&_IERC20PermitAllowed.TransactOpts, holder, spender, nonce, expiry, allowed, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20PermitAllowed *IERC20PermitAllowedTransactorSession) Permit(holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20PermitAllowed.Contract.Permit(&_IERC20PermitAllowed.TransactOpts, holder, spender, nonce, expiry, allowed, v, r, s)
}

// IMulticallMetaData contains all meta data concerning the IMulticall contract.
var IMulticallMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IMulticallABI is the input ABI used to generate the binding from.
// Deprecated: Use IMulticallMetaData.ABI instead.
var IMulticallABI = IMulticallMetaData.ABI

// IMulticall is an auto generated Go binding around an Ethereum contract.
type IMulticall struct {
	IMulticallCaller     // Read-only binding to the contract
	IMulticallTransactor // Write-only binding to the contract
	IMulticallFilterer   // Log filterer for contract events
}

// IMulticallCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMulticallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMulticallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMulticallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMulticallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMulticallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMulticallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMulticallSession struct {
	Contract     *IMulticall       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMulticallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMulticallCallerSession struct {
	Contract *IMulticallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IMulticallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMulticallTransactorSession struct {
	Contract     *IMulticallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IMulticallRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMulticallRaw struct {
	Contract *IMulticall // Generic contract binding to access the raw methods on
}

// IMulticallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMulticallCallerRaw struct {
	Contract *IMulticallCaller // Generic read-only contract binding to access the raw methods on
}

// IMulticallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMulticallTransactorRaw struct {
	Contract *IMulticallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMulticall creates a new instance of IMulticall, bound to a specific deployed contract.
func NewIMulticall(address common.Address, backend bind.ContractBackend) (*IMulticall, error) {
	contract, err := bindIMulticall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMulticall{IMulticallCaller: IMulticallCaller{contract: contract}, IMulticallTransactor: IMulticallTransactor{contract: contract}, IMulticallFilterer: IMulticallFilterer{contract: contract}}, nil
}

// NewIMulticallCaller creates a new read-only instance of IMulticall, bound to a specific deployed contract.
func NewIMulticallCaller(address common.Address, caller bind.ContractCaller) (*IMulticallCaller, error) {
	contract, err := bindIMulticall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMulticallCaller{contract: contract}, nil
}

// NewIMulticallTransactor creates a new write-only instance of IMulticall, bound to a specific deployed contract.
func NewIMulticallTransactor(address common.Address, transactor bind.ContractTransactor) (*IMulticallTransactor, error) {
	contract, err := bindIMulticall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMulticallTransactor{contract: contract}, nil
}

// NewIMulticallFilterer creates a new log filterer instance of IMulticall, bound to a specific deployed contract.
func NewIMulticallFilterer(address common.Address, filterer bind.ContractFilterer) (*IMulticallFilterer, error) {
	contract, err := bindIMulticall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMulticallFilterer{contract: contract}, nil
}

// bindIMulticall binds a generic wrapper to an already deployed contract.
func bindIMulticall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMulticallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMulticall *IMulticallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMulticall.Contract.IMulticallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMulticall *IMulticallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMulticall.Contract.IMulticallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMulticall *IMulticallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMulticall.Contract.IMulticallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMulticall *IMulticallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMulticall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMulticall *IMulticallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMulticall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMulticall *IMulticallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMulticall.Contract.contract.Transact(opts, method, params...)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IMulticall *IMulticallTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _IMulticall.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IMulticall *IMulticallSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _IMulticall.Contract.Multicall(&_IMulticall.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IMulticall *IMulticallTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _IMulticall.Contract.Multicall(&_IMulticall.TransactOpts, data)
}

// IPeripheryImmutableStateMetaData contains all meta data concerning the IPeripheryImmutableState contract.
var IPeripheryImmutableStateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPeripheryImmutableStateABI is the input ABI used to generate the binding from.
// Deprecated: Use IPeripheryImmutableStateMetaData.ABI instead.
var IPeripheryImmutableStateABI = IPeripheryImmutableStateMetaData.ABI

// IPeripheryImmutableState is an auto generated Go binding around an Ethereum contract.
type IPeripheryImmutableState struct {
	IPeripheryImmutableStateCaller     // Read-only binding to the contract
	IPeripheryImmutableStateTransactor // Write-only binding to the contract
	IPeripheryImmutableStateFilterer   // Log filterer for contract events
}

// IPeripheryImmutableStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPeripheryImmutableStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPeripheryImmutableStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPeripheryImmutableStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPeripheryImmutableStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPeripheryImmutableStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPeripheryImmutableStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPeripheryImmutableStateSession struct {
	Contract     *IPeripheryImmutableState // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IPeripheryImmutableStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPeripheryImmutableStateCallerSession struct {
	Contract *IPeripheryImmutableStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IPeripheryImmutableStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPeripheryImmutableStateTransactorSession struct {
	Contract     *IPeripheryImmutableStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IPeripheryImmutableStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPeripheryImmutableStateRaw struct {
	Contract *IPeripheryImmutableState // Generic contract binding to access the raw methods on
}

// IPeripheryImmutableStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPeripheryImmutableStateCallerRaw struct {
	Contract *IPeripheryImmutableStateCaller // Generic read-only contract binding to access the raw methods on
}

// IPeripheryImmutableStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPeripheryImmutableStateTransactorRaw struct {
	Contract *IPeripheryImmutableStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPeripheryImmutableState creates a new instance of IPeripheryImmutableState, bound to a specific deployed contract.
func NewIPeripheryImmutableState(address common.Address, backend bind.ContractBackend) (*IPeripheryImmutableState, error) {
	contract, err := bindIPeripheryImmutableState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPeripheryImmutableState{IPeripheryImmutableStateCaller: IPeripheryImmutableStateCaller{contract: contract}, IPeripheryImmutableStateTransactor: IPeripheryImmutableStateTransactor{contract: contract}, IPeripheryImmutableStateFilterer: IPeripheryImmutableStateFilterer{contract: contract}}, nil
}

// NewIPeripheryImmutableStateCaller creates a new read-only instance of IPeripheryImmutableState, bound to a specific deployed contract.
func NewIPeripheryImmutableStateCaller(address common.Address, caller bind.ContractCaller) (*IPeripheryImmutableStateCaller, error) {
	contract, err := bindIPeripheryImmutableState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPeripheryImmutableStateCaller{contract: contract}, nil
}

// NewIPeripheryImmutableStateTransactor creates a new write-only instance of IPeripheryImmutableState, bound to a specific deployed contract.
func NewIPeripheryImmutableStateTransactor(address common.Address, transactor bind.ContractTransactor) (*IPeripheryImmutableStateTransactor, error) {
	contract, err := bindIPeripheryImmutableState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPeripheryImmutableStateTransactor{contract: contract}, nil
}

// NewIPeripheryImmutableStateFilterer creates a new log filterer instance of IPeripheryImmutableState, bound to a specific deployed contract.
func NewIPeripheryImmutableStateFilterer(address common.Address, filterer bind.ContractFilterer) (*IPeripheryImmutableStateFilterer, error) {
	contract, err := bindIPeripheryImmutableState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPeripheryImmutableStateFilterer{contract: contract}, nil
}

// bindIPeripheryImmutableState binds a generic wrapper to an already deployed contract.
func bindIPeripheryImmutableState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPeripheryImmutableStateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPeripheryImmutableState *IPeripheryImmutableStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPeripheryImmutableState.Contract.IPeripheryImmutableStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPeripheryImmutableState *IPeripheryImmutableStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPeripheryImmutableState.Contract.IPeripheryImmutableStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPeripheryImmutableState *IPeripheryImmutableStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPeripheryImmutableState.Contract.IPeripheryImmutableStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPeripheryImmutableState *IPeripheryImmutableStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPeripheryImmutableState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPeripheryImmutableState *IPeripheryImmutableStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPeripheryImmutableState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPeripheryImmutableState *IPeripheryImmutableStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPeripheryImmutableState.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_IPeripheryImmutableState *IPeripheryImmutableStateCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPeripheryImmutableState.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_IPeripheryImmutableState *IPeripheryImmutableStateSession) WETH9() (common.Address, error) {
	return _IPeripheryImmutableState.Contract.WETH9(&_IPeripheryImmutableState.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_IPeripheryImmutableState *IPeripheryImmutableStateCallerSession) WETH9() (common.Address, error) {
	return _IPeripheryImmutableState.Contract.WETH9(&_IPeripheryImmutableState.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IPeripheryImmutableState *IPeripheryImmutableStateCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPeripheryImmutableState.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IPeripheryImmutableState *IPeripheryImmutableStateSession) Factory() (common.Address, error) {
	return _IPeripheryImmutableState.Contract.Factory(&_IPeripheryImmutableState.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IPeripheryImmutableState *IPeripheryImmutableStateCallerSession) Factory() (common.Address, error) {
	return _IPeripheryImmutableState.Contract.Factory(&_IPeripheryImmutableState.CallOpts)
}

// IPeripheryPaymentsMetaData contains all meta data concerning the IPeripheryPayments contract.
var IPeripheryPaymentsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IPeripheryPaymentsABI is the input ABI used to generate the binding from.
// Deprecated: Use IPeripheryPaymentsMetaData.ABI instead.
var IPeripheryPaymentsABI = IPeripheryPaymentsMetaData.ABI

// IPeripheryPayments is an auto generated Go binding around an Ethereum contract.
type IPeripheryPayments struct {
	IPeripheryPaymentsCaller     // Read-only binding to the contract
	IPeripheryPaymentsTransactor // Write-only binding to the contract
	IPeripheryPaymentsFilterer   // Log filterer for contract events
}

// IPeripheryPaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPeripheryPaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPeripheryPaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPeripheryPaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPeripheryPaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPeripheryPaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPeripheryPaymentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPeripheryPaymentsSession struct {
	Contract     *IPeripheryPayments // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IPeripheryPaymentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPeripheryPaymentsCallerSession struct {
	Contract *IPeripheryPaymentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IPeripheryPaymentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPeripheryPaymentsTransactorSession struct {
	Contract     *IPeripheryPaymentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IPeripheryPaymentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPeripheryPaymentsRaw struct {
	Contract *IPeripheryPayments // Generic contract binding to access the raw methods on
}

// IPeripheryPaymentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPeripheryPaymentsCallerRaw struct {
	Contract *IPeripheryPaymentsCaller // Generic read-only contract binding to access the raw methods on
}

// IPeripheryPaymentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPeripheryPaymentsTransactorRaw struct {
	Contract *IPeripheryPaymentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPeripheryPayments creates a new instance of IPeripheryPayments, bound to a specific deployed contract.
func NewIPeripheryPayments(address common.Address, backend bind.ContractBackend) (*IPeripheryPayments, error) {
	contract, err := bindIPeripheryPayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPeripheryPayments{IPeripheryPaymentsCaller: IPeripheryPaymentsCaller{contract: contract}, IPeripheryPaymentsTransactor: IPeripheryPaymentsTransactor{contract: contract}, IPeripheryPaymentsFilterer: IPeripheryPaymentsFilterer{contract: contract}}, nil
}

// NewIPeripheryPaymentsCaller creates a new read-only instance of IPeripheryPayments, bound to a specific deployed contract.
func NewIPeripheryPaymentsCaller(address common.Address, caller bind.ContractCaller) (*IPeripheryPaymentsCaller, error) {
	contract, err := bindIPeripheryPayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPeripheryPaymentsCaller{contract: contract}, nil
}

// NewIPeripheryPaymentsTransactor creates a new write-only instance of IPeripheryPayments, bound to a specific deployed contract.
func NewIPeripheryPaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*IPeripheryPaymentsTransactor, error) {
	contract, err := bindIPeripheryPayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPeripheryPaymentsTransactor{contract: contract}, nil
}

// NewIPeripheryPaymentsFilterer creates a new log filterer instance of IPeripheryPayments, bound to a specific deployed contract.
func NewIPeripheryPaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*IPeripheryPaymentsFilterer, error) {
	contract, err := bindIPeripheryPayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPeripheryPaymentsFilterer{contract: contract}, nil
}

// bindIPeripheryPayments binds a generic wrapper to an already deployed contract.
func bindIPeripheryPayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPeripheryPaymentsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPeripheryPayments *IPeripheryPaymentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPeripheryPayments.Contract.IPeripheryPaymentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPeripheryPayments *IPeripheryPaymentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPeripheryPayments.Contract.IPeripheryPaymentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPeripheryPayments *IPeripheryPaymentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPeripheryPayments.Contract.IPeripheryPaymentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPeripheryPayments *IPeripheryPaymentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPeripheryPayments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPeripheryPayments *IPeripheryPaymentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPeripheryPayments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPeripheryPayments *IPeripheryPaymentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPeripheryPayments.Contract.contract.Transact(opts, method, params...)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_IPeripheryPayments *IPeripheryPaymentsTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPeripheryPayments.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_IPeripheryPayments *IPeripheryPaymentsSession) RefundETH() (*types.Transaction, error) {
	return _IPeripheryPayments.Contract.RefundETH(&_IPeripheryPayments.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_IPeripheryPayments *IPeripheryPaymentsTransactorSession) RefundETH() (*types.Transaction, error) {
	return _IPeripheryPayments.Contract.RefundETH(&_IPeripheryPayments.TransactOpts)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPayments *IPeripheryPaymentsTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPayments.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPayments *IPeripheryPaymentsSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPayments.Contract.SweepToken(&_IPeripheryPayments.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPayments *IPeripheryPaymentsTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPayments.Contract.SweepToken(&_IPeripheryPayments.TransactOpts, token, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPayments *IPeripheryPaymentsTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPayments.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPayments *IPeripheryPaymentsSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPayments.Contract.UnwrapWETH9(&_IPeripheryPayments.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPayments *IPeripheryPaymentsTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPayments.Contract.UnwrapWETH9(&_IPeripheryPayments.TransactOpts, amountMinimum, recipient)
}

// IPeripheryPaymentsWithFeeMetaData contains all meta data concerning the IPeripheryPaymentsWithFee contract.
var IPeripheryPaymentsWithFeeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IPeripheryPaymentsWithFeeABI is the input ABI used to generate the binding from.
// Deprecated: Use IPeripheryPaymentsWithFeeMetaData.ABI instead.
var IPeripheryPaymentsWithFeeABI = IPeripheryPaymentsWithFeeMetaData.ABI

// IPeripheryPaymentsWithFee is an auto generated Go binding around an Ethereum contract.
type IPeripheryPaymentsWithFee struct {
	IPeripheryPaymentsWithFeeCaller     // Read-only binding to the contract
	IPeripheryPaymentsWithFeeTransactor // Write-only binding to the contract
	IPeripheryPaymentsWithFeeFilterer   // Log filterer for contract events
}

// IPeripheryPaymentsWithFeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPeripheryPaymentsWithFeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPeripheryPaymentsWithFeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPeripheryPaymentsWithFeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPeripheryPaymentsWithFeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPeripheryPaymentsWithFeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPeripheryPaymentsWithFeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPeripheryPaymentsWithFeeSession struct {
	Contract     *IPeripheryPaymentsWithFee // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IPeripheryPaymentsWithFeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPeripheryPaymentsWithFeeCallerSession struct {
	Contract *IPeripheryPaymentsWithFeeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// IPeripheryPaymentsWithFeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPeripheryPaymentsWithFeeTransactorSession struct {
	Contract     *IPeripheryPaymentsWithFeeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// IPeripheryPaymentsWithFeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPeripheryPaymentsWithFeeRaw struct {
	Contract *IPeripheryPaymentsWithFee // Generic contract binding to access the raw methods on
}

// IPeripheryPaymentsWithFeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPeripheryPaymentsWithFeeCallerRaw struct {
	Contract *IPeripheryPaymentsWithFeeCaller // Generic read-only contract binding to access the raw methods on
}

// IPeripheryPaymentsWithFeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPeripheryPaymentsWithFeeTransactorRaw struct {
	Contract *IPeripheryPaymentsWithFeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPeripheryPaymentsWithFee creates a new instance of IPeripheryPaymentsWithFee, bound to a specific deployed contract.
func NewIPeripheryPaymentsWithFee(address common.Address, backend bind.ContractBackend) (*IPeripheryPaymentsWithFee, error) {
	contract, err := bindIPeripheryPaymentsWithFee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPeripheryPaymentsWithFee{IPeripheryPaymentsWithFeeCaller: IPeripheryPaymentsWithFeeCaller{contract: contract}, IPeripheryPaymentsWithFeeTransactor: IPeripheryPaymentsWithFeeTransactor{contract: contract}, IPeripheryPaymentsWithFeeFilterer: IPeripheryPaymentsWithFeeFilterer{contract: contract}}, nil
}

// NewIPeripheryPaymentsWithFeeCaller creates a new read-only instance of IPeripheryPaymentsWithFee, bound to a specific deployed contract.
func NewIPeripheryPaymentsWithFeeCaller(address common.Address, caller bind.ContractCaller) (*IPeripheryPaymentsWithFeeCaller, error) {
	contract, err := bindIPeripheryPaymentsWithFee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPeripheryPaymentsWithFeeCaller{contract: contract}, nil
}

// NewIPeripheryPaymentsWithFeeTransactor creates a new write-only instance of IPeripheryPaymentsWithFee, bound to a specific deployed contract.
func NewIPeripheryPaymentsWithFeeTransactor(address common.Address, transactor bind.ContractTransactor) (*IPeripheryPaymentsWithFeeTransactor, error) {
	contract, err := bindIPeripheryPaymentsWithFee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPeripheryPaymentsWithFeeTransactor{contract: contract}, nil
}

// NewIPeripheryPaymentsWithFeeFilterer creates a new log filterer instance of IPeripheryPaymentsWithFee, bound to a specific deployed contract.
func NewIPeripheryPaymentsWithFeeFilterer(address common.Address, filterer bind.ContractFilterer) (*IPeripheryPaymentsWithFeeFilterer, error) {
	contract, err := bindIPeripheryPaymentsWithFee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPeripheryPaymentsWithFeeFilterer{contract: contract}, nil
}

// bindIPeripheryPaymentsWithFee binds a generic wrapper to an already deployed contract.
func bindIPeripheryPaymentsWithFee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPeripheryPaymentsWithFeeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPeripheryPaymentsWithFee.Contract.IPeripheryPaymentsWithFeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.IPeripheryPaymentsWithFeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.IPeripheryPaymentsWithFeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPeripheryPaymentsWithFee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.contract.Transact(opts, method, params...)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeSession) RefundETH() (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.RefundETH(&_IPeripheryPaymentsWithFee.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactorSession) RefundETH() (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.RefundETH(&_IPeripheryPaymentsWithFee.TransactOpts)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.SweepToken(&_IPeripheryPaymentsWithFee.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.SweepToken(&_IPeripheryPaymentsWithFee.TransactOpts, token, amountMinimum, recipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.SweepTokenWithFee(&_IPeripheryPaymentsWithFee.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.SweepTokenWithFee(&_IPeripheryPaymentsWithFee.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.UnwrapWETH9(&_IPeripheryPaymentsWithFee.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.UnwrapWETH9(&_IPeripheryPaymentsWithFee.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactor) UnwrapWETH9WithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.contract.Transact(opts, "unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.UnwrapWETH9WithFee(&_IPeripheryPaymentsWithFee.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_IPeripheryPaymentsWithFee *IPeripheryPaymentsWithFeeTransactorSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _IPeripheryPaymentsWithFee.Contract.UnwrapWETH9WithFee(&_IPeripheryPaymentsWithFee.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// ISelfPermitMetaData contains all meta data concerning the ISelfPermit contract.
var ISelfPermitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ISelfPermitABI is the input ABI used to generate the binding from.
// Deprecated: Use ISelfPermitMetaData.ABI instead.
var ISelfPermitABI = ISelfPermitMetaData.ABI

// ISelfPermit is an auto generated Go binding around an Ethereum contract.
type ISelfPermit struct {
	ISelfPermitCaller     // Read-only binding to the contract
	ISelfPermitTransactor // Write-only binding to the contract
	ISelfPermitFilterer   // Log filterer for contract events
}

// ISelfPermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISelfPermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISelfPermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISelfPermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISelfPermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISelfPermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISelfPermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISelfPermitSession struct {
	Contract     *ISelfPermit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISelfPermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISelfPermitCallerSession struct {
	Contract *ISelfPermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ISelfPermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISelfPermitTransactorSession struct {
	Contract     *ISelfPermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ISelfPermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISelfPermitRaw struct {
	Contract *ISelfPermit // Generic contract binding to access the raw methods on
}

// ISelfPermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISelfPermitCallerRaw struct {
	Contract *ISelfPermitCaller // Generic read-only contract binding to access the raw methods on
}

// ISelfPermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISelfPermitTransactorRaw struct {
	Contract *ISelfPermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISelfPermit creates a new instance of ISelfPermit, bound to a specific deployed contract.
func NewISelfPermit(address common.Address, backend bind.ContractBackend) (*ISelfPermit, error) {
	contract, err := bindISelfPermit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISelfPermit{ISelfPermitCaller: ISelfPermitCaller{contract: contract}, ISelfPermitTransactor: ISelfPermitTransactor{contract: contract}, ISelfPermitFilterer: ISelfPermitFilterer{contract: contract}}, nil
}

// NewISelfPermitCaller creates a new read-only instance of ISelfPermit, bound to a specific deployed contract.
func NewISelfPermitCaller(address common.Address, caller bind.ContractCaller) (*ISelfPermitCaller, error) {
	contract, err := bindISelfPermit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISelfPermitCaller{contract: contract}, nil
}

// NewISelfPermitTransactor creates a new write-only instance of ISelfPermit, bound to a specific deployed contract.
func NewISelfPermitTransactor(address common.Address, transactor bind.ContractTransactor) (*ISelfPermitTransactor, error) {
	contract, err := bindISelfPermit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISelfPermitTransactor{contract: contract}, nil
}

// NewISelfPermitFilterer creates a new log filterer instance of ISelfPermit, bound to a specific deployed contract.
func NewISelfPermitFilterer(address common.Address, filterer bind.ContractFilterer) (*ISelfPermitFilterer, error) {
	contract, err := bindISelfPermit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISelfPermitFilterer{contract: contract}, nil
}

// bindISelfPermit binds a generic wrapper to an already deployed contract.
func bindISelfPermit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISelfPermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISelfPermit *ISelfPermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISelfPermit.Contract.ISelfPermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISelfPermit *ISelfPermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISelfPermit.Contract.ISelfPermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISelfPermit *ISelfPermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISelfPermit.Contract.ISelfPermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISelfPermit *ISelfPermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISelfPermit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISelfPermit *ISelfPermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISelfPermit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISelfPermit *ISelfPermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISelfPermit.Contract.contract.Transact(opts, method, params...)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.Contract.SelfPermit(&_ISelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.Contract.SelfPermit(&_ISelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.Contract.SelfPermitAllowed(&_ISelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.Contract.SelfPermitAllowed(&_ISelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.Contract.SelfPermitAllowedIfNecessary(&_ISelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.Contract.SelfPermitAllowedIfNecessary(&_ISelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.Contract.SelfPermitIfNecessary(&_ISelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_ISelfPermit *ISelfPermitTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISelfPermit.Contract.SelfPermitIfNecessary(&_ISelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// ISwapRouterMetaData contains all meta data concerning the ISwapRouter contract.
var ISwapRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISwapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwapRouterMetaData.ABI instead.
var ISwapRouterABI = ISwapRouterMetaData.ABI

// ISwapRouter is an auto generated Go binding around an Ethereum contract.
type ISwapRouter struct {
	ISwapRouterCaller     // Read-only binding to the contract
	ISwapRouterTransactor // Write-only binding to the contract
	ISwapRouterFilterer   // Log filterer for contract events
}

// ISwapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapRouterSession struct {
	Contract     *ISwapRouter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapRouterCallerSession struct {
	Contract *ISwapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ISwapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapRouterTransactorSession struct {
	Contract     *ISwapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ISwapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapRouterRaw struct {
	Contract *ISwapRouter // Generic contract binding to access the raw methods on
}

// ISwapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapRouterCallerRaw struct {
	Contract *ISwapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// ISwapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapRouterTransactorRaw struct {
	Contract *ISwapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwapRouter creates a new instance of ISwapRouter, bound to a specific deployed contract.
func NewISwapRouter(address common.Address, backend bind.ContractBackend) (*ISwapRouter, error) {
	contract, err := bindISwapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwapRouter{ISwapRouterCaller: ISwapRouterCaller{contract: contract}, ISwapRouterTransactor: ISwapRouterTransactor{contract: contract}, ISwapRouterFilterer: ISwapRouterFilterer{contract: contract}}, nil
}

// NewISwapRouterCaller creates a new read-only instance of ISwapRouter, bound to a specific deployed contract.
func NewISwapRouterCaller(address common.Address, caller bind.ContractCaller) (*ISwapRouterCaller, error) {
	contract, err := bindISwapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapRouterCaller{contract: contract}, nil
}

// NewISwapRouterTransactor creates a new write-only instance of ISwapRouter, bound to a specific deployed contract.
func NewISwapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwapRouterTransactor, error) {
	contract, err := bindISwapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapRouterTransactor{contract: contract}, nil
}

// NewISwapRouterFilterer creates a new log filterer instance of ISwapRouter, bound to a specific deployed contract.
func NewISwapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwapRouterFilterer, error) {
	contract, err := bindISwapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapRouterFilterer{contract: contract}, nil
}

// bindISwapRouter binds a generic wrapper to an already deployed contract.
func bindISwapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISwapRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapRouter *ISwapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapRouter.Contract.ISwapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapRouter *ISwapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapRouter.Contract.ISwapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapRouter *ISwapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapRouter.Contract.ISwapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapRouter *ISwapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapRouter *ISwapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapRouter *ISwapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapRouter.Contract.contract.Transact(opts, method, params...)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_ISwapRouter *ISwapRouterTransactor) ExactInput(opts *bind.TransactOpts, params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _ISwapRouter.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_ISwapRouter *ISwapRouterSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _ISwapRouter.Contract.ExactInput(&_ISwapRouter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_ISwapRouter *ISwapRouterTransactorSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _ISwapRouter.Contract.ExactInput(&_ISwapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_ISwapRouter *ISwapRouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_ISwapRouter *ISwapRouterSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter.Contract.ExactInputSingle(&_ISwapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_ISwapRouter *ISwapRouterTransactorSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter.Contract.ExactInputSingle(&_ISwapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_ISwapRouter *ISwapRouterTransactor) ExactOutput(opts *bind.TransactOpts, params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _ISwapRouter.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_ISwapRouter *ISwapRouterSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _ISwapRouter.Contract.ExactOutput(&_ISwapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_ISwapRouter *ISwapRouterTransactorSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _ISwapRouter.Contract.ExactOutput(&_ISwapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_ISwapRouter *ISwapRouterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_ISwapRouter *ISwapRouterSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter.Contract.ExactOutputSingle(&_ISwapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_ISwapRouter *ISwapRouterTransactorSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _ISwapRouter.Contract.ExactOutputSingle(&_ISwapRouter.TransactOpts, params)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_ISwapRouter *ISwapRouterTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _ISwapRouter.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_ISwapRouter *ISwapRouterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _ISwapRouter.Contract.UniswapV3SwapCallback(&_ISwapRouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_ISwapRouter *ISwapRouterTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _ISwapRouter.Contract.UniswapV3SwapCallback(&_ISwapRouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// IUniswapV3PoolMetaData contains all meta data concerning the IUniswapV3Pool contract.
var IUniswapV3PoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"CollectProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextOld\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextNew\",\"type\":\"uint16\"}],\"name\":\"IncreaseObservationCardinalityNext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0New\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1New\",\"type\":\"uint8\"}],\"name\":\"SetFeeProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collectProtocol\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal0X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal1X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"}],\"name\":\"increaseObservationCardinalityNext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidityPerTick\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"observations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityCumulativeX128\",\"type\":\"uint160\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"observe\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"tickCumulatives\",\"type\":\"int56[]\"},{\"internalType\":\"uint160[]\",\"name\":\"secondsPerLiquidityCumulativeX128s\",\"type\":\"uint160[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"token0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"token1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"feeProtocol0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol1\",\"type\":\"uint8\"}],\"name\":\"setFeeProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slot0\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"snapshotCumulativesInside\",\"outputs\":[{\"internalType\":\"int56\",\"name\":\"tickCumulativeInside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityInsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsInside\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountSpecified\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"wordPosition\",\"type\":\"int16\"}],\"name\":\"tickBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside0X128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside1X128\",\"type\":\"uint256\"},{\"internalType\":\"int56\",\"name\":\"tickCumulativeOutside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityOutsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsOutside\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IUniswapV3PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3PoolMetaData.ABI instead.
var IUniswapV3PoolABI = IUniswapV3PoolMetaData.ABI

// IUniswapV3Pool is an auto generated Go binding around an Ethereum contract.
type IUniswapV3Pool struct {
	IUniswapV3PoolCaller     // Read-only binding to the contract
	IUniswapV3PoolTransactor // Write-only binding to the contract
	IUniswapV3PoolFilterer   // Log filterer for contract events
}

// IUniswapV3PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3PoolSession struct {
	Contract     *IUniswapV3Pool   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUniswapV3PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3PoolCallerSession struct {
	Contract *IUniswapV3PoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IUniswapV3PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3PoolTransactorSession struct {
	Contract     *IUniswapV3PoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IUniswapV3PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3PoolRaw struct {
	Contract *IUniswapV3Pool // Generic contract binding to access the raw methods on
}

// IUniswapV3PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3PoolCallerRaw struct {
	Contract *IUniswapV3PoolCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3PoolTransactorRaw struct {
	Contract *IUniswapV3PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3Pool creates a new instance of IUniswapV3Pool, bound to a specific deployed contract.
func NewIUniswapV3Pool(address common.Address, backend bind.ContractBackend) (*IUniswapV3Pool, error) {
	contract, err := bindIUniswapV3Pool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3Pool{IUniswapV3PoolCaller: IUniswapV3PoolCaller{contract: contract}, IUniswapV3PoolTransactor: IUniswapV3PoolTransactor{contract: contract}, IUniswapV3PoolFilterer: IUniswapV3PoolFilterer{contract: contract}}, nil
}

// NewIUniswapV3PoolCaller creates a new read-only instance of IUniswapV3Pool, bound to a specific deployed contract.
func NewIUniswapV3PoolCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3PoolCaller, error) {
	contract, err := bindIUniswapV3Pool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolCaller{contract: contract}, nil
}

// NewIUniswapV3PoolTransactor creates a new write-only instance of IUniswapV3Pool, bound to a specific deployed contract.
func NewIUniswapV3PoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3PoolTransactor, error) {
	contract, err := bindIUniswapV3Pool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolTransactor{contract: contract}, nil
}

// NewIUniswapV3PoolFilterer creates a new log filterer instance of IUniswapV3Pool, bound to a specific deployed contract.
func NewIUniswapV3PoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3PoolFilterer, error) {
	contract, err := bindIUniswapV3Pool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolFilterer{contract: contract}, nil
}

// bindIUniswapV3Pool binds a generic wrapper to an already deployed contract.
func bindIUniswapV3Pool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3Pool *IUniswapV3PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3Pool.Contract.IUniswapV3PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3Pool *IUniswapV3PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.IUniswapV3PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3Pool *IUniswapV3PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.IUniswapV3PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3Pool *IUniswapV3PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3Pool *IUniswapV3PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3Pool *IUniswapV3PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Factory() (common.Address, error) {
	return _IUniswapV3Pool.Contract.Factory(&_IUniswapV3Pool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Factory() (common.Address, error) {
	return _IUniswapV3Pool.Contract.Factory(&_IUniswapV3Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Fee() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.Fee(&_IUniswapV3Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Fee() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.Fee(&_IUniswapV3Pool.CallOpts)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) FeeGrowthGlobal0X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "feeGrowthGlobal0X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_IUniswapV3Pool *IUniswapV3PoolSession) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.FeeGrowthGlobal0X128(&_IUniswapV3Pool.CallOpts)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.FeeGrowthGlobal0X128(&_IUniswapV3Pool.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) FeeGrowthGlobal1X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "feeGrowthGlobal1X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_IUniswapV3Pool *IUniswapV3PoolSession) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.FeeGrowthGlobal1X128(&_IUniswapV3Pool.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.FeeGrowthGlobal1X128(&_IUniswapV3Pool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Liquidity() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.Liquidity(&_IUniswapV3Pool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Liquidity() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.Liquidity(&_IUniswapV3Pool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) MaxLiquidityPerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "maxLiquidityPerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_IUniswapV3Pool *IUniswapV3PoolSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.MaxLiquidityPerTick(&_IUniswapV3Pool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.MaxLiquidityPerTick(&_IUniswapV3Pool.CallOpts)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Observations(opts *bind.CallOpts, index *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "observations", index)

	outstruct := new(struct {
		BlockTimestamp                    uint32
		TickCumulative                    *big.Int
		SecondsPerLiquidityCumulativeX128 *big.Int
		Initialized                       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockTimestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.TickCumulative = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityCumulativeX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Initialized = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Observations(index *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	return _IUniswapV3Pool.Contract.Observations(&_IUniswapV3Pool.CallOpts, index)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Observations(index *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	return _IUniswapV3Pool.Contract.Observations(&_IUniswapV3Pool.CallOpts, index)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Observe(opts *bind.CallOpts, secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "observe", secondsAgos)

	outstruct := new(struct {
		TickCumulatives                    []*big.Int
		SecondsPerLiquidityCumulativeX128s []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulatives = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SecondsPerLiquidityCumulativeX128s = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Observe(secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	return _IUniswapV3Pool.Contract.Observe(&_IUniswapV3Pool.CallOpts, secondsAgos)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Observe(secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	return _IUniswapV3Pool.Contract.Observe(&_IUniswapV3Pool.CallOpts, secondsAgos)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Positions(opts *bind.CallOpts, key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "positions", key)

	outstruct := new(struct {
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Positions(key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _IUniswapV3Pool.Contract.Positions(&_IUniswapV3Pool.CallOpts, key)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Positions(key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _IUniswapV3Pool.Contract.Positions(&_IUniswapV3Pool.CallOpts, key)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) ProtocolFees(opts *bind.CallOpts) (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "protocolFees")

	outstruct := new(struct {
		Token0 *big.Int
		Token1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Token1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_IUniswapV3Pool *IUniswapV3PoolSession) ProtocolFees() (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	return _IUniswapV3Pool.Contract.ProtocolFees(&_IUniswapV3Pool.CallOpts)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) ProtocolFees() (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	return _IUniswapV3Pool.Contract.ProtocolFees(&_IUniswapV3Pool.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "slot0")

	outstruct := new(struct {
		SqrtPriceX96               *big.Int
		Tick                       *big.Int
		ObservationIndex           uint16
		ObservationCardinality     uint16
		ObservationCardinalityNext uint16
		FeeProtocol                uint8
		Unlocked                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPriceX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationCardinality = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationCardinalityNext = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.FeeProtocol = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _IUniswapV3Pool.Contract.Slot0(&_IUniswapV3Pool.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _IUniswapV3Pool.Contract.Slot0(&_IUniswapV3Pool.CallOpts)
}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint32 secondsInside)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) SnapshotCumulativesInside(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "snapshotCumulativesInside", tickLower, tickUpper)

	outstruct := new(struct {
		TickCumulativeInside          *big.Int
		SecondsPerLiquidityInsideX128 *big.Int
		SecondsInside                 uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulativeInside = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityInsideX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondsInside = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint32 secondsInside)
func (_IUniswapV3Pool *IUniswapV3PoolSession) SnapshotCumulativesInside(tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	return _IUniswapV3Pool.Contract.SnapshotCumulativesInside(&_IUniswapV3Pool.CallOpts, tickLower, tickUpper)
}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint32 secondsInside)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) SnapshotCumulativesInside(tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	return _IUniswapV3Pool.Contract.SnapshotCumulativesInside(&_IUniswapV3Pool.CallOpts, tickLower, tickUpper)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) TickBitmap(opts *bind.CallOpts, wordPosition int16) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "tickBitmap", wordPosition)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_IUniswapV3Pool *IUniswapV3PoolSession) TickBitmap(wordPosition int16) (*big.Int, error) {
	return _IUniswapV3Pool.Contract.TickBitmap(&_IUniswapV3Pool.CallOpts, wordPosition)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) TickBitmap(wordPosition int16) (*big.Int, error) {
	return _IUniswapV3Pool.Contract.TickBitmap(&_IUniswapV3Pool.CallOpts, wordPosition)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_IUniswapV3Pool *IUniswapV3PoolSession) TickSpacing() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.TickSpacing(&_IUniswapV3Pool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) TickSpacing() (*big.Int, error) {
	return _IUniswapV3Pool.Contract.TickSpacing(&_IUniswapV3Pool.CallOpts)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Ticks(opts *bind.CallOpts, tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "ticks", tick)

	outstruct := new(struct {
		LiquidityGross                 *big.Int
		LiquidityNet                   *big.Int
		FeeGrowthOutside0X128          *big.Int
		FeeGrowthOutside1X128          *big.Int
		TickCumulativeOutside          *big.Int
		SecondsPerLiquidityOutsideX128 *big.Int
		SecondsOutside                 uint32
		Initialized                    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityGross = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityNet = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside0X128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside1X128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickCumulativeOutside = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityOutsideX128 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.SecondsOutside = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.Initialized = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Ticks(tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _IUniswapV3Pool.Contract.Ticks(&_IUniswapV3Pool.CallOpts, tick)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Ticks(tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _IUniswapV3Pool.Contract.Ticks(&_IUniswapV3Pool.CallOpts, tick)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Token0() (common.Address, error) {
	return _IUniswapV3Pool.Contract.Token0(&_IUniswapV3Pool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Token0() (common.Address, error) {
	return _IUniswapV3Pool.Contract.Token0(&_IUniswapV3Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV3Pool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Token1() (common.Address, error) {
	return _IUniswapV3Pool.Contract.Token1(&_IUniswapV3Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IUniswapV3Pool *IUniswapV3PoolCallerSession) Token1() (common.Address, error) {
	return _IUniswapV3Pool.Contract.Token1(&_IUniswapV3Pool.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolTransactor) Burn(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.contract.Transact(opts, "burn", tickLower, tickUpper, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Burn(tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Burn(&_IUniswapV3Pool.TransactOpts, tickLower, tickUpper, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolTransactorSession) Burn(tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Burn(&_IUniswapV3Pool.TransactOpts, tickLower, tickUpper, amount)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.contract.Transact(opts, "collect", recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Collect(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Collect(&_IUniswapV3Pool.TransactOpts, recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolTransactorSession) Collect(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Collect(&_IUniswapV3Pool.TransactOpts, recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolTransactor) CollectProtocol(opts *bind.TransactOpts, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.contract.Transact(opts, "collectProtocol", recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolSession) CollectProtocol(recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.CollectProtocol(&_IUniswapV3Pool.TransactOpts, recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolTransactorSession) CollectProtocol(recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.CollectProtocol(&_IUniswapV3Pool.TransactOpts, recipient, amount0Requested, amount1Requested)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IUniswapV3Pool *IUniswapV3PoolTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3Pool.contract.Transact(opts, "flash", recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IUniswapV3Pool *IUniswapV3PoolSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Flash(&_IUniswapV3Pool.TransactOpts, recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IUniswapV3Pool *IUniswapV3PoolTransactorSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Flash(&_IUniswapV3Pool.TransactOpts, recipient, amount0, amount1, data)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_IUniswapV3Pool *IUniswapV3PoolTransactor) IncreaseObservationCardinalityNext(opts *bind.TransactOpts, observationCardinalityNext uint16) (*types.Transaction, error) {
	return _IUniswapV3Pool.contract.Transact(opts, "increaseObservationCardinalityNext", observationCardinalityNext)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_IUniswapV3Pool *IUniswapV3PoolSession) IncreaseObservationCardinalityNext(observationCardinalityNext uint16) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.IncreaseObservationCardinalityNext(&_IUniswapV3Pool.TransactOpts, observationCardinalityNext)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_IUniswapV3Pool *IUniswapV3PoolTransactorSession) IncreaseObservationCardinalityNext(observationCardinalityNext uint16) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.IncreaseObservationCardinalityNext(&_IUniswapV3Pool.TransactOpts, observationCardinalityNext)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_IUniswapV3Pool *IUniswapV3PoolTransactor) Initialize(opts *bind.TransactOpts, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.contract.Transact(opts, "initialize", sqrtPriceX96)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_IUniswapV3Pool *IUniswapV3PoolSession) Initialize(sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Initialize(&_IUniswapV3Pool.TransactOpts, sqrtPriceX96)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_IUniswapV3Pool *IUniswapV3PoolTransactorSession) Initialize(sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Initialize(&_IUniswapV3Pool.TransactOpts, sqrtPriceX96)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3Pool.contract.Transact(opts, "mint", recipient, tickLower, tickUpper, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Mint(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Mint(&_IUniswapV3Pool.TransactOpts, recipient, tickLower, tickUpper, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolTransactorSession) Mint(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Mint(&_IUniswapV3Pool.TransactOpts, recipient, tickLower, tickUpper, amount, data)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x8206a4d1.
//
// Solidity: function setFeeProtocol(uint8 feeProtocol0, uint8 feeProtocol1) returns()
func (_IUniswapV3Pool *IUniswapV3PoolTransactor) SetFeeProtocol(opts *bind.TransactOpts, feeProtocol0 uint8, feeProtocol1 uint8) (*types.Transaction, error) {
	return _IUniswapV3Pool.contract.Transact(opts, "setFeeProtocol", feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x8206a4d1.
//
// Solidity: function setFeeProtocol(uint8 feeProtocol0, uint8 feeProtocol1) returns()
func (_IUniswapV3Pool *IUniswapV3PoolSession) SetFeeProtocol(feeProtocol0 uint8, feeProtocol1 uint8) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.SetFeeProtocol(&_IUniswapV3Pool.TransactOpts, feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x8206a4d1.
//
// Solidity: function setFeeProtocol(uint8 feeProtocol0, uint8 feeProtocol1) returns()
func (_IUniswapV3Pool *IUniswapV3PoolTransactorSession) SetFeeProtocol(feeProtocol0 uint8, feeProtocol1 uint8) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.SetFeeProtocol(&_IUniswapV3Pool.TransactOpts, feeProtocol0, feeProtocol1)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3Pool.contract.Transact(opts, "swap", recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Swap(&_IUniswapV3Pool.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolTransactorSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3Pool.Contract.Swap(&_IUniswapV3Pool.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// IUniswapV3PoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolBurnIterator struct {
	Event *IUniswapV3PoolBurn // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolBurn)
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
		it.Event = new(IUniswapV3PoolBurn)
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
func (it *IUniswapV3PoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolBurn represents a Burn event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolBurn struct {
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount    *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*IUniswapV3PoolBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolBurnIterator{contract: _IUniswapV3Pool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolBurn, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolBurn)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseBurn(log types.Log) (*IUniswapV3PoolBurn, error) {
	event := new(IUniswapV3PoolBurn)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolCollectIterator struct {
	Event *IUniswapV3PoolCollect // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolCollect)
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
		it.Event = new(IUniswapV3PoolCollect)
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
func (it *IUniswapV3PoolCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolCollect represents a Collect event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolCollect struct {
	Owner     common.Address
	Recipient common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*IUniswapV3PoolCollectIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolCollectIterator{contract: _IUniswapV3Pool.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolCollect, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolCollect)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "Collect", log); err != nil {
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

// ParseCollect is a log parse operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseCollect(log types.Log) (*IUniswapV3PoolCollect, error) {
	event := new(IUniswapV3PoolCollect)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolCollectProtocolIterator is returned from FilterCollectProtocol and is used to iterate over the raw logs and unpacked data for CollectProtocol events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolCollectProtocolIterator struct {
	Event *IUniswapV3PoolCollectProtocol // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolCollectProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolCollectProtocol)
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
		it.Event = new(IUniswapV3PoolCollectProtocol)
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
func (it *IUniswapV3PoolCollectProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolCollectProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolCollectProtocol represents a CollectProtocol event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolCollectProtocol struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectProtocol is a free log retrieval operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterCollectProtocol(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IUniswapV3PoolCollectProtocolIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolCollectProtocolIterator{contract: _IUniswapV3Pool.contract, event: "CollectProtocol", logs: logs, sub: sub}, nil
}

// WatchCollectProtocol is a free log subscription operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchCollectProtocol(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolCollectProtocol, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolCollectProtocol)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
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

// ParseCollectProtocol is a log parse operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseCollectProtocol(log types.Log) (*IUniswapV3PoolCollectProtocol, error) {
	event := new(IUniswapV3PoolCollectProtocol)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolFlashIterator struct {
	Event *IUniswapV3PoolFlash // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolFlash)
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
		it.Event = new(IUniswapV3PoolFlash)
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
func (it *IUniswapV3PoolFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolFlash represents a Flash event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolFlash struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Paid0     *big.Int
	Paid1     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IUniswapV3PoolFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolFlashIterator{contract: _IUniswapV3Pool.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolFlash)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "Flash", log); err != nil {
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

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseFlash(log types.Log) (*IUniswapV3PoolFlash, error) {
	event := new(IUniswapV3PoolFlash)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolIncreaseObservationCardinalityNextIterator is returned from FilterIncreaseObservationCardinalityNext and is used to iterate over the raw logs and unpacked data for IncreaseObservationCardinalityNext events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolIncreaseObservationCardinalityNextIterator struct {
	Event *IUniswapV3PoolIncreaseObservationCardinalityNext // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolIncreaseObservationCardinalityNextIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolIncreaseObservationCardinalityNext)
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
		it.Event = new(IUniswapV3PoolIncreaseObservationCardinalityNext)
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
func (it *IUniswapV3PoolIncreaseObservationCardinalityNextIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolIncreaseObservationCardinalityNextIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolIncreaseObservationCardinalityNext represents a IncreaseObservationCardinalityNext event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolIncreaseObservationCardinalityNext struct {
	ObservationCardinalityNextOld uint16
	ObservationCardinalityNextNew uint16
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterIncreaseObservationCardinalityNext is a free log retrieval operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterIncreaseObservationCardinalityNext(opts *bind.FilterOpts) (*IUniswapV3PoolIncreaseObservationCardinalityNextIterator, error) {

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolIncreaseObservationCardinalityNextIterator{contract: _IUniswapV3Pool.contract, event: "IncreaseObservationCardinalityNext", logs: logs, sub: sub}, nil
}

// WatchIncreaseObservationCardinalityNext is a free log subscription operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchIncreaseObservationCardinalityNext(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolIncreaseObservationCardinalityNext) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolIncreaseObservationCardinalityNext)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
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

// ParseIncreaseObservationCardinalityNext is a log parse operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseIncreaseObservationCardinalityNext(log types.Log) (*IUniswapV3PoolIncreaseObservationCardinalityNext, error) {
	event := new(IUniswapV3PoolIncreaseObservationCardinalityNext)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolInitializeIterator struct {
	Event *IUniswapV3PoolInitialize // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolInitialize)
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
		it.Event = new(IUniswapV3PoolInitialize)
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
func (it *IUniswapV3PoolInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolInitialize represents a Initialize event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolInitialize struct {
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterInitialize(opts *bind.FilterOpts) (*IUniswapV3PoolInitializeIterator, error) {

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolInitializeIterator{contract: _IUniswapV3Pool.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolInitialize) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolInitialize)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseInitialize(log types.Log) (*IUniswapV3PoolInitialize, error) {
	event := new(IUniswapV3PoolInitialize)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolMintIterator struct {
	Event *IUniswapV3PoolMint // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolMint)
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
		it.Event = new(IUniswapV3PoolMint)
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
func (it *IUniswapV3PoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolMint represents a Mint event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolMint struct {
	Sender    common.Address
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount    *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*IUniswapV3PoolMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolMintIterator{contract: _IUniswapV3Pool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolMint, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolMint)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseMint(log types.Log) (*IUniswapV3PoolMint, error) {
	event := new(IUniswapV3PoolMint)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolSetFeeProtocolIterator is returned from FilterSetFeeProtocol and is used to iterate over the raw logs and unpacked data for SetFeeProtocol events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolSetFeeProtocolIterator struct {
	Event *IUniswapV3PoolSetFeeProtocol // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolSetFeeProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolSetFeeProtocol)
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
		it.Event = new(IUniswapV3PoolSetFeeProtocol)
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
func (it *IUniswapV3PoolSetFeeProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolSetFeeProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolSetFeeProtocol represents a SetFeeProtocol event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolSetFeeProtocol struct {
	FeeProtocol0Old uint8
	FeeProtocol1Old uint8
	FeeProtocol0New uint8
	FeeProtocol1New uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeProtocol is a free log retrieval operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterSetFeeProtocol(opts *bind.FilterOpts) (*IUniswapV3PoolSetFeeProtocolIterator, error) {

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolSetFeeProtocolIterator{contract: _IUniswapV3Pool.contract, event: "SetFeeProtocol", logs: logs, sub: sub}, nil
}

// WatchSetFeeProtocol is a free log subscription operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchSetFeeProtocol(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolSetFeeProtocol) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolSetFeeProtocol)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
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

// ParseSetFeeProtocol is a log parse operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseSetFeeProtocol(log types.Log) (*IUniswapV3PoolSetFeeProtocol, error) {
	event := new(IUniswapV3PoolSetFeeProtocol)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolSwapIterator struct {
	Event *IUniswapV3PoolSwap // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolSwap)
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
		it.Event = new(IUniswapV3PoolSwap)
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
func (it *IUniswapV3PoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolSwap represents a Swap event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolSwap struct {
	Sender       common.Address
	Recipient    common.Address
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IUniswapV3PoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolSwapIterator{contract: _IUniswapV3Pool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolSwap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolSwap)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseSwap(log types.Log) (*IUniswapV3PoolSwap, error) {
	event := new(IUniswapV3PoolSwap)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolActionsMetaData contains all meta data concerning the IUniswapV3PoolActions contract.
var IUniswapV3PoolActionsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"}],\"name\":\"increaseObservationCardinalityNext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountSpecified\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IUniswapV3PoolActionsABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3PoolActionsMetaData.ABI instead.
var IUniswapV3PoolActionsABI = IUniswapV3PoolActionsMetaData.ABI

// IUniswapV3PoolActions is an auto generated Go binding around an Ethereum contract.
type IUniswapV3PoolActions struct {
	IUniswapV3PoolActionsCaller     // Read-only binding to the contract
	IUniswapV3PoolActionsTransactor // Write-only binding to the contract
	IUniswapV3PoolActionsFilterer   // Log filterer for contract events
}

// IUniswapV3PoolActionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3PoolActionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolActionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3PoolActionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolActionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3PoolActionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolActionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3PoolActionsSession struct {
	Contract     *IUniswapV3PoolActions // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IUniswapV3PoolActionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3PoolActionsCallerSession struct {
	Contract *IUniswapV3PoolActionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IUniswapV3PoolActionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3PoolActionsTransactorSession struct {
	Contract     *IUniswapV3PoolActionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IUniswapV3PoolActionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3PoolActionsRaw struct {
	Contract *IUniswapV3PoolActions // Generic contract binding to access the raw methods on
}

// IUniswapV3PoolActionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3PoolActionsCallerRaw struct {
	Contract *IUniswapV3PoolActionsCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3PoolActionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3PoolActionsTransactorRaw struct {
	Contract *IUniswapV3PoolActionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3PoolActions creates a new instance of IUniswapV3PoolActions, bound to a specific deployed contract.
func NewIUniswapV3PoolActions(address common.Address, backend bind.ContractBackend) (*IUniswapV3PoolActions, error) {
	contract, err := bindIUniswapV3PoolActions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolActions{IUniswapV3PoolActionsCaller: IUniswapV3PoolActionsCaller{contract: contract}, IUniswapV3PoolActionsTransactor: IUniswapV3PoolActionsTransactor{contract: contract}, IUniswapV3PoolActionsFilterer: IUniswapV3PoolActionsFilterer{contract: contract}}, nil
}

// NewIUniswapV3PoolActionsCaller creates a new read-only instance of IUniswapV3PoolActions, bound to a specific deployed contract.
func NewIUniswapV3PoolActionsCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3PoolActionsCaller, error) {
	contract, err := bindIUniswapV3PoolActions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolActionsCaller{contract: contract}, nil
}

// NewIUniswapV3PoolActionsTransactor creates a new write-only instance of IUniswapV3PoolActions, bound to a specific deployed contract.
func NewIUniswapV3PoolActionsTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3PoolActionsTransactor, error) {
	contract, err := bindIUniswapV3PoolActions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolActionsTransactor{contract: contract}, nil
}

// NewIUniswapV3PoolActionsFilterer creates a new log filterer instance of IUniswapV3PoolActions, bound to a specific deployed contract.
func NewIUniswapV3PoolActionsFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3PoolActionsFilterer, error) {
	contract, err := bindIUniswapV3PoolActions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolActionsFilterer{contract: contract}, nil
}

// bindIUniswapV3PoolActions binds a generic wrapper to an already deployed contract.
func bindIUniswapV3PoolActions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3PoolActionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolActions.Contract.IUniswapV3PoolActionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.IUniswapV3PoolActionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.IUniswapV3PoolActionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolActions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.contract.Transact(opts, method, params...)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactor) Burn(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.contract.Transact(opts, "burn", tickLower, tickUpper, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsSession) Burn(tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Burn(&_IUniswapV3PoolActions.TransactOpts, tickLower, tickUpper, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactorSession) Burn(tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Burn(&_IUniswapV3PoolActions.TransactOpts, tickLower, tickUpper, amount)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.contract.Transact(opts, "collect", recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsSession) Collect(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Collect(&_IUniswapV3PoolActions.TransactOpts, recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactorSession) Collect(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Collect(&_IUniswapV3PoolActions.TransactOpts, recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.contract.Transact(opts, "flash", recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Flash(&_IUniswapV3PoolActions.TransactOpts, recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactorSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Flash(&_IUniswapV3PoolActions.TransactOpts, recipient, amount0, amount1, data)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactor) IncreaseObservationCardinalityNext(opts *bind.TransactOpts, observationCardinalityNext uint16) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.contract.Transact(opts, "increaseObservationCardinalityNext", observationCardinalityNext)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsSession) IncreaseObservationCardinalityNext(observationCardinalityNext uint16) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.IncreaseObservationCardinalityNext(&_IUniswapV3PoolActions.TransactOpts, observationCardinalityNext)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactorSession) IncreaseObservationCardinalityNext(observationCardinalityNext uint16) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.IncreaseObservationCardinalityNext(&_IUniswapV3PoolActions.TransactOpts, observationCardinalityNext)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactor) Initialize(opts *bind.TransactOpts, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.contract.Transact(opts, "initialize", sqrtPriceX96)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsSession) Initialize(sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Initialize(&_IUniswapV3PoolActions.TransactOpts, sqrtPriceX96)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactorSession) Initialize(sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Initialize(&_IUniswapV3PoolActions.TransactOpts, sqrtPriceX96)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.contract.Transact(opts, "mint", recipient, tickLower, tickUpper, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsSession) Mint(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Mint(&_IUniswapV3PoolActions.TransactOpts, recipient, tickLower, tickUpper, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactorSession) Mint(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Mint(&_IUniswapV3PoolActions.TransactOpts, recipient, tickLower, tickUpper, amount, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.contract.Transact(opts, "swap", recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Swap(&_IUniswapV3PoolActions.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_IUniswapV3PoolActions *IUniswapV3PoolActionsTransactorSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3PoolActions.Contract.Swap(&_IUniswapV3PoolActions.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// IUniswapV3PoolDerivedStateMetaData contains all meta data concerning the IUniswapV3PoolDerivedState contract.
var IUniswapV3PoolDerivedStateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"observe\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"tickCumulatives\",\"type\":\"int56[]\"},{\"internalType\":\"uint160[]\",\"name\":\"secondsPerLiquidityCumulativeX128s\",\"type\":\"uint160[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"snapshotCumulativesInside\",\"outputs\":[{\"internalType\":\"int56\",\"name\":\"tickCumulativeInside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityInsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsInside\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IUniswapV3PoolDerivedStateABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3PoolDerivedStateMetaData.ABI instead.
var IUniswapV3PoolDerivedStateABI = IUniswapV3PoolDerivedStateMetaData.ABI

// IUniswapV3PoolDerivedState is an auto generated Go binding around an Ethereum contract.
type IUniswapV3PoolDerivedState struct {
	IUniswapV3PoolDerivedStateCaller     // Read-only binding to the contract
	IUniswapV3PoolDerivedStateTransactor // Write-only binding to the contract
	IUniswapV3PoolDerivedStateFilterer   // Log filterer for contract events
}

// IUniswapV3PoolDerivedStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3PoolDerivedStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolDerivedStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3PoolDerivedStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolDerivedStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3PoolDerivedStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolDerivedStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3PoolDerivedStateSession struct {
	Contract     *IUniswapV3PoolDerivedState // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IUniswapV3PoolDerivedStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3PoolDerivedStateCallerSession struct {
	Contract *IUniswapV3PoolDerivedStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IUniswapV3PoolDerivedStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3PoolDerivedStateTransactorSession struct {
	Contract     *IUniswapV3PoolDerivedStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IUniswapV3PoolDerivedStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3PoolDerivedStateRaw struct {
	Contract *IUniswapV3PoolDerivedState // Generic contract binding to access the raw methods on
}

// IUniswapV3PoolDerivedStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3PoolDerivedStateCallerRaw struct {
	Contract *IUniswapV3PoolDerivedStateCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3PoolDerivedStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3PoolDerivedStateTransactorRaw struct {
	Contract *IUniswapV3PoolDerivedStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3PoolDerivedState creates a new instance of IUniswapV3PoolDerivedState, bound to a specific deployed contract.
func NewIUniswapV3PoolDerivedState(address common.Address, backend bind.ContractBackend) (*IUniswapV3PoolDerivedState, error) {
	contract, err := bindIUniswapV3PoolDerivedState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDerivedState{IUniswapV3PoolDerivedStateCaller: IUniswapV3PoolDerivedStateCaller{contract: contract}, IUniswapV3PoolDerivedStateTransactor: IUniswapV3PoolDerivedStateTransactor{contract: contract}, IUniswapV3PoolDerivedStateFilterer: IUniswapV3PoolDerivedStateFilterer{contract: contract}}, nil
}

// NewIUniswapV3PoolDerivedStateCaller creates a new read-only instance of IUniswapV3PoolDerivedState, bound to a specific deployed contract.
func NewIUniswapV3PoolDerivedStateCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3PoolDerivedStateCaller, error) {
	contract, err := bindIUniswapV3PoolDerivedState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDerivedStateCaller{contract: contract}, nil
}

// NewIUniswapV3PoolDerivedStateTransactor creates a new write-only instance of IUniswapV3PoolDerivedState, bound to a specific deployed contract.
func NewIUniswapV3PoolDerivedStateTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3PoolDerivedStateTransactor, error) {
	contract, err := bindIUniswapV3PoolDerivedState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDerivedStateTransactor{contract: contract}, nil
}

// NewIUniswapV3PoolDerivedStateFilterer creates a new log filterer instance of IUniswapV3PoolDerivedState, bound to a specific deployed contract.
func NewIUniswapV3PoolDerivedStateFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3PoolDerivedStateFilterer, error) {
	contract, err := bindIUniswapV3PoolDerivedState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDerivedStateFilterer{contract: contract}, nil
}

// bindIUniswapV3PoolDerivedState binds a generic wrapper to an already deployed contract.
func bindIUniswapV3PoolDerivedState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3PoolDerivedStateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolDerivedState.Contract.IUniswapV3PoolDerivedStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolDerivedState.Contract.IUniswapV3PoolDerivedStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolDerivedState.Contract.IUniswapV3PoolDerivedStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolDerivedState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolDerivedState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolDerivedState.Contract.contract.Transact(opts, method, params...)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateCaller) Observe(opts *bind.CallOpts, secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	var out []interface{}
	err := _IUniswapV3PoolDerivedState.contract.Call(opts, &out, "observe", secondsAgos)

	outstruct := new(struct {
		TickCumulatives                    []*big.Int
		SecondsPerLiquidityCumulativeX128s []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulatives = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SecondsPerLiquidityCumulativeX128s = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateSession) Observe(secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	return _IUniswapV3PoolDerivedState.Contract.Observe(&_IUniswapV3PoolDerivedState.CallOpts, secondsAgos)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateCallerSession) Observe(secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	return _IUniswapV3PoolDerivedState.Contract.Observe(&_IUniswapV3PoolDerivedState.CallOpts, secondsAgos)
}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint32 secondsInside)
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateCaller) SnapshotCumulativesInside(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	var out []interface{}
	err := _IUniswapV3PoolDerivedState.contract.Call(opts, &out, "snapshotCumulativesInside", tickLower, tickUpper)

	outstruct := new(struct {
		TickCumulativeInside          *big.Int
		SecondsPerLiquidityInsideX128 *big.Int
		SecondsInside                 uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulativeInside = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityInsideX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondsInside = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint32 secondsInside)
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateSession) SnapshotCumulativesInside(tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	return _IUniswapV3PoolDerivedState.Contract.SnapshotCumulativesInside(&_IUniswapV3PoolDerivedState.CallOpts, tickLower, tickUpper)
}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint32 secondsInside)
func (_IUniswapV3PoolDerivedState *IUniswapV3PoolDerivedStateCallerSession) SnapshotCumulativesInside(tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	return _IUniswapV3PoolDerivedState.Contract.SnapshotCumulativesInside(&_IUniswapV3PoolDerivedState.CallOpts, tickLower, tickUpper)
}

// IUniswapV3PoolEventsMetaData contains all meta data concerning the IUniswapV3PoolEvents contract.
var IUniswapV3PoolEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"CollectProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextOld\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextNew\",\"type\":\"uint16\"}],\"name\":\"IncreaseObservationCardinalityNext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0New\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1New\",\"type\":\"uint8\"}],\"name\":\"SetFeeProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"}]",
}

// IUniswapV3PoolEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3PoolEventsMetaData.ABI instead.
var IUniswapV3PoolEventsABI = IUniswapV3PoolEventsMetaData.ABI

// IUniswapV3PoolEvents is an auto generated Go binding around an Ethereum contract.
type IUniswapV3PoolEvents struct {
	IUniswapV3PoolEventsCaller     // Read-only binding to the contract
	IUniswapV3PoolEventsTransactor // Write-only binding to the contract
	IUniswapV3PoolEventsFilterer   // Log filterer for contract events
}

// IUniswapV3PoolEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3PoolEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3PoolEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3PoolEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3PoolEventsSession struct {
	Contract     *IUniswapV3PoolEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IUniswapV3PoolEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3PoolEventsCallerSession struct {
	Contract *IUniswapV3PoolEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IUniswapV3PoolEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3PoolEventsTransactorSession struct {
	Contract     *IUniswapV3PoolEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IUniswapV3PoolEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3PoolEventsRaw struct {
	Contract *IUniswapV3PoolEvents // Generic contract binding to access the raw methods on
}

// IUniswapV3PoolEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3PoolEventsCallerRaw struct {
	Contract *IUniswapV3PoolEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3PoolEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3PoolEventsTransactorRaw struct {
	Contract *IUniswapV3PoolEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3PoolEvents creates a new instance of IUniswapV3PoolEvents, bound to a specific deployed contract.
func NewIUniswapV3PoolEvents(address common.Address, backend bind.ContractBackend) (*IUniswapV3PoolEvents, error) {
	contract, err := bindIUniswapV3PoolEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEvents{IUniswapV3PoolEventsCaller: IUniswapV3PoolEventsCaller{contract: contract}, IUniswapV3PoolEventsTransactor: IUniswapV3PoolEventsTransactor{contract: contract}, IUniswapV3PoolEventsFilterer: IUniswapV3PoolEventsFilterer{contract: contract}}, nil
}

// NewIUniswapV3PoolEventsCaller creates a new read-only instance of IUniswapV3PoolEvents, bound to a specific deployed contract.
func NewIUniswapV3PoolEventsCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3PoolEventsCaller, error) {
	contract, err := bindIUniswapV3PoolEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsCaller{contract: contract}, nil
}

// NewIUniswapV3PoolEventsTransactor creates a new write-only instance of IUniswapV3PoolEvents, bound to a specific deployed contract.
func NewIUniswapV3PoolEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3PoolEventsTransactor, error) {
	contract, err := bindIUniswapV3PoolEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsTransactor{contract: contract}, nil
}

// NewIUniswapV3PoolEventsFilterer creates a new log filterer instance of IUniswapV3PoolEvents, bound to a specific deployed contract.
func NewIUniswapV3PoolEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3PoolEventsFilterer, error) {
	contract, err := bindIUniswapV3PoolEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsFilterer{contract: contract}, nil
}

// bindIUniswapV3PoolEvents binds a generic wrapper to an already deployed contract.
func bindIUniswapV3PoolEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3PoolEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolEvents.Contract.IUniswapV3PoolEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolEvents.Contract.IUniswapV3PoolEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolEvents.Contract.IUniswapV3PoolEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolEvents.Contract.contract.Transact(opts, method, params...)
}

// IUniswapV3PoolEventsBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsBurnIterator struct {
	Event *IUniswapV3PoolEventsBurn // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsBurn)
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
		it.Event = new(IUniswapV3PoolEventsBurn)
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
func (it *IUniswapV3PoolEventsBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsBurn represents a Burn event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsBurn struct {
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount    *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*IUniswapV3PoolEventsBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsBurnIterator{contract: _IUniswapV3PoolEvents.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsBurn, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsBurn)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseBurn(log types.Log) (*IUniswapV3PoolEventsBurn, error) {
	event := new(IUniswapV3PoolEventsBurn)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsCollectIterator struct {
	Event *IUniswapV3PoolEventsCollect // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsCollect)
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
		it.Event = new(IUniswapV3PoolEventsCollect)
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
func (it *IUniswapV3PoolEventsCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsCollect represents a Collect event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsCollect struct {
	Owner     common.Address
	Recipient common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*IUniswapV3PoolEventsCollectIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsCollectIterator{contract: _IUniswapV3PoolEvents.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsCollect, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsCollect)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Collect", log); err != nil {
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

// ParseCollect is a log parse operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseCollect(log types.Log) (*IUniswapV3PoolEventsCollect, error) {
	event := new(IUniswapV3PoolEventsCollect)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsCollectProtocolIterator is returned from FilterCollectProtocol and is used to iterate over the raw logs and unpacked data for CollectProtocol events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsCollectProtocolIterator struct {
	Event *IUniswapV3PoolEventsCollectProtocol // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsCollectProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsCollectProtocol)
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
		it.Event = new(IUniswapV3PoolEventsCollectProtocol)
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
func (it *IUniswapV3PoolEventsCollectProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsCollectProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsCollectProtocol represents a CollectProtocol event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsCollectProtocol struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectProtocol is a free log retrieval operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterCollectProtocol(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IUniswapV3PoolEventsCollectProtocolIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsCollectProtocolIterator{contract: _IUniswapV3PoolEvents.contract, event: "CollectProtocol", logs: logs, sub: sub}, nil
}

// WatchCollectProtocol is a free log subscription operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchCollectProtocol(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsCollectProtocol, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsCollectProtocol)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
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

// ParseCollectProtocol is a log parse operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseCollectProtocol(log types.Log) (*IUniswapV3PoolEventsCollectProtocol, error) {
	event := new(IUniswapV3PoolEventsCollectProtocol)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsFlashIterator struct {
	Event *IUniswapV3PoolEventsFlash // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsFlash)
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
		it.Event = new(IUniswapV3PoolEventsFlash)
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
func (it *IUniswapV3PoolEventsFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsFlash represents a Flash event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsFlash struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Paid0     *big.Int
	Paid1     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IUniswapV3PoolEventsFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsFlashIterator{contract: _IUniswapV3PoolEvents.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsFlash)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Flash", log); err != nil {
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

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseFlash(log types.Log) (*IUniswapV3PoolEventsFlash, error) {
	event := new(IUniswapV3PoolEventsFlash)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator is returned from FilterIncreaseObservationCardinalityNext and is used to iterate over the raw logs and unpacked data for IncreaseObservationCardinalityNext events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator struct {
	Event *IUniswapV3PoolEventsIncreaseObservationCardinalityNext // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsIncreaseObservationCardinalityNext)
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
		it.Event = new(IUniswapV3PoolEventsIncreaseObservationCardinalityNext)
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
func (it *IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsIncreaseObservationCardinalityNext represents a IncreaseObservationCardinalityNext event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsIncreaseObservationCardinalityNext struct {
	ObservationCardinalityNextOld uint16
	ObservationCardinalityNextNew uint16
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterIncreaseObservationCardinalityNext is a free log retrieval operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterIncreaseObservationCardinalityNext(opts *bind.FilterOpts) (*IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator{contract: _IUniswapV3PoolEvents.contract, event: "IncreaseObservationCardinalityNext", logs: logs, sub: sub}, nil
}

// WatchIncreaseObservationCardinalityNext is a free log subscription operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchIncreaseObservationCardinalityNext(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsIncreaseObservationCardinalityNext) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsIncreaseObservationCardinalityNext)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
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

// ParseIncreaseObservationCardinalityNext is a log parse operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseIncreaseObservationCardinalityNext(log types.Log) (*IUniswapV3PoolEventsIncreaseObservationCardinalityNext, error) {
	event := new(IUniswapV3PoolEventsIncreaseObservationCardinalityNext)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsInitializeIterator struct {
	Event *IUniswapV3PoolEventsInitialize // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsInitialize)
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
		it.Event = new(IUniswapV3PoolEventsInitialize)
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
func (it *IUniswapV3PoolEventsInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsInitialize represents a Initialize event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsInitialize struct {
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterInitialize(opts *bind.FilterOpts) (*IUniswapV3PoolEventsInitializeIterator, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsInitializeIterator{contract: _IUniswapV3PoolEvents.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsInitialize) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsInitialize)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseInitialize(log types.Log) (*IUniswapV3PoolEventsInitialize, error) {
	event := new(IUniswapV3PoolEventsInitialize)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsMintIterator struct {
	Event *IUniswapV3PoolEventsMint // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsMint)
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
		it.Event = new(IUniswapV3PoolEventsMint)
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
func (it *IUniswapV3PoolEventsMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsMint represents a Mint event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsMint struct {
	Sender    common.Address
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount    *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*IUniswapV3PoolEventsMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsMintIterator{contract: _IUniswapV3PoolEvents.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsMint, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsMint)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseMint(log types.Log) (*IUniswapV3PoolEventsMint, error) {
	event := new(IUniswapV3PoolEventsMint)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsSetFeeProtocolIterator is returned from FilterSetFeeProtocol and is used to iterate over the raw logs and unpacked data for SetFeeProtocol events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsSetFeeProtocolIterator struct {
	Event *IUniswapV3PoolEventsSetFeeProtocol // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsSetFeeProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsSetFeeProtocol)
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
		it.Event = new(IUniswapV3PoolEventsSetFeeProtocol)
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
func (it *IUniswapV3PoolEventsSetFeeProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsSetFeeProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsSetFeeProtocol represents a SetFeeProtocol event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsSetFeeProtocol struct {
	FeeProtocol0Old uint8
	FeeProtocol1Old uint8
	FeeProtocol0New uint8
	FeeProtocol1New uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeProtocol is a free log retrieval operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterSetFeeProtocol(opts *bind.FilterOpts) (*IUniswapV3PoolEventsSetFeeProtocolIterator, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsSetFeeProtocolIterator{contract: _IUniswapV3PoolEvents.contract, event: "SetFeeProtocol", logs: logs, sub: sub}, nil
}

// WatchSetFeeProtocol is a free log subscription operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchSetFeeProtocol(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsSetFeeProtocol) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsSetFeeProtocol)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
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

// ParseSetFeeProtocol is a log parse operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseSetFeeProtocol(log types.Log) (*IUniswapV3PoolEventsSetFeeProtocol, error) {
	event := new(IUniswapV3PoolEventsSetFeeProtocol)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsSwapIterator struct {
	Event *IUniswapV3PoolEventsSwap // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsSwap)
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
		it.Event = new(IUniswapV3PoolEventsSwap)
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
func (it *IUniswapV3PoolEventsSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsSwap represents a Swap event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsSwap struct {
	Sender       common.Address
	Recipient    common.Address
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IUniswapV3PoolEventsSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsSwapIterator{contract: _IUniswapV3PoolEvents.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsSwap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsSwap)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseSwap(log types.Log) (*IUniswapV3PoolEventsSwap, error) {
	event := new(IUniswapV3PoolEventsSwap)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolImmutablesMetaData contains all meta data concerning the IUniswapV3PoolImmutables contract.
var IUniswapV3PoolImmutablesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidityPerTick\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IUniswapV3PoolImmutablesABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3PoolImmutablesMetaData.ABI instead.
var IUniswapV3PoolImmutablesABI = IUniswapV3PoolImmutablesMetaData.ABI

// IUniswapV3PoolImmutables is an auto generated Go binding around an Ethereum contract.
type IUniswapV3PoolImmutables struct {
	IUniswapV3PoolImmutablesCaller     // Read-only binding to the contract
	IUniswapV3PoolImmutablesTransactor // Write-only binding to the contract
	IUniswapV3PoolImmutablesFilterer   // Log filterer for contract events
}

// IUniswapV3PoolImmutablesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3PoolImmutablesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolImmutablesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3PoolImmutablesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolImmutablesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3PoolImmutablesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolImmutablesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3PoolImmutablesSession struct {
	Contract     *IUniswapV3PoolImmutables // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IUniswapV3PoolImmutablesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3PoolImmutablesCallerSession struct {
	Contract *IUniswapV3PoolImmutablesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IUniswapV3PoolImmutablesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3PoolImmutablesTransactorSession struct {
	Contract     *IUniswapV3PoolImmutablesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IUniswapV3PoolImmutablesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3PoolImmutablesRaw struct {
	Contract *IUniswapV3PoolImmutables // Generic contract binding to access the raw methods on
}

// IUniswapV3PoolImmutablesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3PoolImmutablesCallerRaw struct {
	Contract *IUniswapV3PoolImmutablesCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3PoolImmutablesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3PoolImmutablesTransactorRaw struct {
	Contract *IUniswapV3PoolImmutablesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3PoolImmutables creates a new instance of IUniswapV3PoolImmutables, bound to a specific deployed contract.
func NewIUniswapV3PoolImmutables(address common.Address, backend bind.ContractBackend) (*IUniswapV3PoolImmutables, error) {
	contract, err := bindIUniswapV3PoolImmutables(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolImmutables{IUniswapV3PoolImmutablesCaller: IUniswapV3PoolImmutablesCaller{contract: contract}, IUniswapV3PoolImmutablesTransactor: IUniswapV3PoolImmutablesTransactor{contract: contract}, IUniswapV3PoolImmutablesFilterer: IUniswapV3PoolImmutablesFilterer{contract: contract}}, nil
}

// NewIUniswapV3PoolImmutablesCaller creates a new read-only instance of IUniswapV3PoolImmutables, bound to a specific deployed contract.
func NewIUniswapV3PoolImmutablesCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3PoolImmutablesCaller, error) {
	contract, err := bindIUniswapV3PoolImmutables(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolImmutablesCaller{contract: contract}, nil
}

// NewIUniswapV3PoolImmutablesTransactor creates a new write-only instance of IUniswapV3PoolImmutables, bound to a specific deployed contract.
func NewIUniswapV3PoolImmutablesTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3PoolImmutablesTransactor, error) {
	contract, err := bindIUniswapV3PoolImmutables(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolImmutablesTransactor{contract: contract}, nil
}

// NewIUniswapV3PoolImmutablesFilterer creates a new log filterer instance of IUniswapV3PoolImmutables, bound to a specific deployed contract.
func NewIUniswapV3PoolImmutablesFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3PoolImmutablesFilterer, error) {
	contract, err := bindIUniswapV3PoolImmutables(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolImmutablesFilterer{contract: contract}, nil
}

// bindIUniswapV3PoolImmutables binds a generic wrapper to an already deployed contract.
func bindIUniswapV3PoolImmutables(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3PoolImmutablesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolImmutables.Contract.IUniswapV3PoolImmutablesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolImmutables.Contract.IUniswapV3PoolImmutablesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolImmutables.Contract.IUniswapV3PoolImmutablesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolImmutables.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolImmutables.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolImmutables.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV3PoolImmutables.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesSession) Factory() (common.Address, error) {
	return _IUniswapV3PoolImmutables.Contract.Factory(&_IUniswapV3PoolImmutables.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCallerSession) Factory() (common.Address, error) {
	return _IUniswapV3PoolImmutables.Contract.Factory(&_IUniswapV3PoolImmutables.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3PoolImmutables.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesSession) Fee() (*big.Int, error) {
	return _IUniswapV3PoolImmutables.Contract.Fee(&_IUniswapV3PoolImmutables.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCallerSession) Fee() (*big.Int, error) {
	return _IUniswapV3PoolImmutables.Contract.Fee(&_IUniswapV3PoolImmutables.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCaller) MaxLiquidityPerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3PoolImmutables.contract.Call(opts, &out, "maxLiquidityPerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _IUniswapV3PoolImmutables.Contract.MaxLiquidityPerTick(&_IUniswapV3PoolImmutables.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCallerSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _IUniswapV3PoolImmutables.Contract.MaxLiquidityPerTick(&_IUniswapV3PoolImmutables.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3PoolImmutables.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesSession) TickSpacing() (*big.Int, error) {
	return _IUniswapV3PoolImmutables.Contract.TickSpacing(&_IUniswapV3PoolImmutables.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCallerSession) TickSpacing() (*big.Int, error) {
	return _IUniswapV3PoolImmutables.Contract.TickSpacing(&_IUniswapV3PoolImmutables.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV3PoolImmutables.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesSession) Token0() (common.Address, error) {
	return _IUniswapV3PoolImmutables.Contract.Token0(&_IUniswapV3PoolImmutables.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCallerSession) Token0() (common.Address, error) {
	return _IUniswapV3PoolImmutables.Contract.Token0(&_IUniswapV3PoolImmutables.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV3PoolImmutables.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesSession) Token1() (common.Address, error) {
	return _IUniswapV3PoolImmutables.Contract.Token1(&_IUniswapV3PoolImmutables.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IUniswapV3PoolImmutables *IUniswapV3PoolImmutablesCallerSession) Token1() (common.Address, error) {
	return _IUniswapV3PoolImmutables.Contract.Token1(&_IUniswapV3PoolImmutables.CallOpts)
}

// IUniswapV3PoolOwnerActionsMetaData contains all meta data concerning the IUniswapV3PoolOwnerActions contract.
var IUniswapV3PoolOwnerActionsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collectProtocol\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"feeProtocol0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol1\",\"type\":\"uint8\"}],\"name\":\"setFeeProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IUniswapV3PoolOwnerActionsABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3PoolOwnerActionsMetaData.ABI instead.
var IUniswapV3PoolOwnerActionsABI = IUniswapV3PoolOwnerActionsMetaData.ABI

// IUniswapV3PoolOwnerActions is an auto generated Go binding around an Ethereum contract.
type IUniswapV3PoolOwnerActions struct {
	IUniswapV3PoolOwnerActionsCaller     // Read-only binding to the contract
	IUniswapV3PoolOwnerActionsTransactor // Write-only binding to the contract
	IUniswapV3PoolOwnerActionsFilterer   // Log filterer for contract events
}

// IUniswapV3PoolOwnerActionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3PoolOwnerActionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolOwnerActionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3PoolOwnerActionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolOwnerActionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3PoolOwnerActionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolOwnerActionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3PoolOwnerActionsSession struct {
	Contract     *IUniswapV3PoolOwnerActions // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IUniswapV3PoolOwnerActionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3PoolOwnerActionsCallerSession struct {
	Contract *IUniswapV3PoolOwnerActionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IUniswapV3PoolOwnerActionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3PoolOwnerActionsTransactorSession struct {
	Contract     *IUniswapV3PoolOwnerActionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IUniswapV3PoolOwnerActionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3PoolOwnerActionsRaw struct {
	Contract *IUniswapV3PoolOwnerActions // Generic contract binding to access the raw methods on
}

// IUniswapV3PoolOwnerActionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3PoolOwnerActionsCallerRaw struct {
	Contract *IUniswapV3PoolOwnerActionsCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3PoolOwnerActionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3PoolOwnerActionsTransactorRaw struct {
	Contract *IUniswapV3PoolOwnerActionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3PoolOwnerActions creates a new instance of IUniswapV3PoolOwnerActions, bound to a specific deployed contract.
func NewIUniswapV3PoolOwnerActions(address common.Address, backend bind.ContractBackend) (*IUniswapV3PoolOwnerActions, error) {
	contract, err := bindIUniswapV3PoolOwnerActions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolOwnerActions{IUniswapV3PoolOwnerActionsCaller: IUniswapV3PoolOwnerActionsCaller{contract: contract}, IUniswapV3PoolOwnerActionsTransactor: IUniswapV3PoolOwnerActionsTransactor{contract: contract}, IUniswapV3PoolOwnerActionsFilterer: IUniswapV3PoolOwnerActionsFilterer{contract: contract}}, nil
}

// NewIUniswapV3PoolOwnerActionsCaller creates a new read-only instance of IUniswapV3PoolOwnerActions, bound to a specific deployed contract.
func NewIUniswapV3PoolOwnerActionsCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3PoolOwnerActionsCaller, error) {
	contract, err := bindIUniswapV3PoolOwnerActions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolOwnerActionsCaller{contract: contract}, nil
}

// NewIUniswapV3PoolOwnerActionsTransactor creates a new write-only instance of IUniswapV3PoolOwnerActions, bound to a specific deployed contract.
func NewIUniswapV3PoolOwnerActionsTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3PoolOwnerActionsTransactor, error) {
	contract, err := bindIUniswapV3PoolOwnerActions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolOwnerActionsTransactor{contract: contract}, nil
}

// NewIUniswapV3PoolOwnerActionsFilterer creates a new log filterer instance of IUniswapV3PoolOwnerActions, bound to a specific deployed contract.
func NewIUniswapV3PoolOwnerActionsFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3PoolOwnerActionsFilterer, error) {
	contract, err := bindIUniswapV3PoolOwnerActions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolOwnerActionsFilterer{contract: contract}, nil
}

// bindIUniswapV3PoolOwnerActions binds a generic wrapper to an already deployed contract.
func bindIUniswapV3PoolOwnerActions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3PoolOwnerActionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolOwnerActions.Contract.IUniswapV3PoolOwnerActionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolOwnerActions.Contract.IUniswapV3PoolOwnerActionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolOwnerActions.Contract.IUniswapV3PoolOwnerActionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolOwnerActions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolOwnerActions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolOwnerActions.Contract.contract.Transact(opts, method, params...)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsTransactor) CollectProtocol(opts *bind.TransactOpts, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolOwnerActions.contract.Transact(opts, "collectProtocol", recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsSession) CollectProtocol(recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolOwnerActions.Contract.CollectProtocol(&_IUniswapV3PoolOwnerActions.TransactOpts, recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsTransactorSession) CollectProtocol(recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _IUniswapV3PoolOwnerActions.Contract.CollectProtocol(&_IUniswapV3PoolOwnerActions.TransactOpts, recipient, amount0Requested, amount1Requested)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x8206a4d1.
//
// Solidity: function setFeeProtocol(uint8 feeProtocol0, uint8 feeProtocol1) returns()
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsTransactor) SetFeeProtocol(opts *bind.TransactOpts, feeProtocol0 uint8, feeProtocol1 uint8) (*types.Transaction, error) {
	return _IUniswapV3PoolOwnerActions.contract.Transact(opts, "setFeeProtocol", feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x8206a4d1.
//
// Solidity: function setFeeProtocol(uint8 feeProtocol0, uint8 feeProtocol1) returns()
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsSession) SetFeeProtocol(feeProtocol0 uint8, feeProtocol1 uint8) (*types.Transaction, error) {
	return _IUniswapV3PoolOwnerActions.Contract.SetFeeProtocol(&_IUniswapV3PoolOwnerActions.TransactOpts, feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x8206a4d1.
//
// Solidity: function setFeeProtocol(uint8 feeProtocol0, uint8 feeProtocol1) returns()
func (_IUniswapV3PoolOwnerActions *IUniswapV3PoolOwnerActionsTransactorSession) SetFeeProtocol(feeProtocol0 uint8, feeProtocol1 uint8) (*types.Transaction, error) {
	return _IUniswapV3PoolOwnerActions.Contract.SetFeeProtocol(&_IUniswapV3PoolOwnerActions.TransactOpts, feeProtocol0, feeProtocol1)
}

// IUniswapV3PoolStateMetaData contains all meta data concerning the IUniswapV3PoolState contract.
var IUniswapV3PoolStateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"feeGrowthGlobal0X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal1X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"observations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityCumulativeX128\",\"type\":\"uint160\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"token0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"token1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slot0\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"wordPosition\",\"type\":\"int16\"}],\"name\":\"tickBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside0X128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside1X128\",\"type\":\"uint256\"},{\"internalType\":\"int56\",\"name\":\"tickCumulativeOutside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityOutsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsOutside\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IUniswapV3PoolStateABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3PoolStateMetaData.ABI instead.
var IUniswapV3PoolStateABI = IUniswapV3PoolStateMetaData.ABI

// IUniswapV3PoolState is an auto generated Go binding around an Ethereum contract.
type IUniswapV3PoolState struct {
	IUniswapV3PoolStateCaller     // Read-only binding to the contract
	IUniswapV3PoolStateTransactor // Write-only binding to the contract
	IUniswapV3PoolStateFilterer   // Log filterer for contract events
}

// IUniswapV3PoolStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3PoolStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3PoolStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3PoolStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3PoolStateSession struct {
	Contract     *IUniswapV3PoolState // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IUniswapV3PoolStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3PoolStateCallerSession struct {
	Contract *IUniswapV3PoolStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IUniswapV3PoolStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3PoolStateTransactorSession struct {
	Contract     *IUniswapV3PoolStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IUniswapV3PoolStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3PoolStateRaw struct {
	Contract *IUniswapV3PoolState // Generic contract binding to access the raw methods on
}

// IUniswapV3PoolStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3PoolStateCallerRaw struct {
	Contract *IUniswapV3PoolStateCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3PoolStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3PoolStateTransactorRaw struct {
	Contract *IUniswapV3PoolStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3PoolState creates a new instance of IUniswapV3PoolState, bound to a specific deployed contract.
func NewIUniswapV3PoolState(address common.Address, backend bind.ContractBackend) (*IUniswapV3PoolState, error) {
	contract, err := bindIUniswapV3PoolState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolState{IUniswapV3PoolStateCaller: IUniswapV3PoolStateCaller{contract: contract}, IUniswapV3PoolStateTransactor: IUniswapV3PoolStateTransactor{contract: contract}, IUniswapV3PoolStateFilterer: IUniswapV3PoolStateFilterer{contract: contract}}, nil
}

// NewIUniswapV3PoolStateCaller creates a new read-only instance of IUniswapV3PoolState, bound to a specific deployed contract.
func NewIUniswapV3PoolStateCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3PoolStateCaller, error) {
	contract, err := bindIUniswapV3PoolState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolStateCaller{contract: contract}, nil
}

// NewIUniswapV3PoolStateTransactor creates a new write-only instance of IUniswapV3PoolState, bound to a specific deployed contract.
func NewIUniswapV3PoolStateTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3PoolStateTransactor, error) {
	contract, err := bindIUniswapV3PoolState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolStateTransactor{contract: contract}, nil
}

// NewIUniswapV3PoolStateFilterer creates a new log filterer instance of IUniswapV3PoolState, bound to a specific deployed contract.
func NewIUniswapV3PoolStateFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3PoolStateFilterer, error) {
	contract, err := bindIUniswapV3PoolState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolStateFilterer{contract: contract}, nil
}

// bindIUniswapV3PoolState binds a generic wrapper to an already deployed contract.
func bindIUniswapV3PoolState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3PoolStateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolState *IUniswapV3PoolStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolState.Contract.IUniswapV3PoolStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolState *IUniswapV3PoolStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolState.Contract.IUniswapV3PoolStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolState *IUniswapV3PoolStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolState.Contract.IUniswapV3PoolStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolState *IUniswapV3PoolStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolState *IUniswapV3PoolStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolState *IUniswapV3PoolStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolState.Contract.contract.Transact(opts, method, params...)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCaller) FeeGrowthGlobal0X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3PoolState.contract.Call(opts, &out, "feeGrowthGlobal0X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_IUniswapV3PoolState *IUniswapV3PoolStateSession) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _IUniswapV3PoolState.Contract.FeeGrowthGlobal0X128(&_IUniswapV3PoolState.CallOpts)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCallerSession) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _IUniswapV3PoolState.Contract.FeeGrowthGlobal0X128(&_IUniswapV3PoolState.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCaller) FeeGrowthGlobal1X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3PoolState.contract.Call(opts, &out, "feeGrowthGlobal1X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_IUniswapV3PoolState *IUniswapV3PoolStateSession) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _IUniswapV3PoolState.Contract.FeeGrowthGlobal1X128(&_IUniswapV3PoolState.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCallerSession) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _IUniswapV3PoolState.Contract.FeeGrowthGlobal1X128(&_IUniswapV3PoolState.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCaller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3PoolState.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_IUniswapV3PoolState *IUniswapV3PoolStateSession) Liquidity() (*big.Int, error) {
	return _IUniswapV3PoolState.Contract.Liquidity(&_IUniswapV3PoolState.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCallerSession) Liquidity() (*big.Int, error) {
	return _IUniswapV3PoolState.Contract.Liquidity(&_IUniswapV3PoolState.CallOpts)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCaller) Observations(opts *bind.CallOpts, index *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	var out []interface{}
	err := _IUniswapV3PoolState.contract.Call(opts, &out, "observations", index)

	outstruct := new(struct {
		BlockTimestamp                    uint32
		TickCumulative                    *big.Int
		SecondsPerLiquidityCumulativeX128 *big.Int
		Initialized                       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockTimestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.TickCumulative = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityCumulativeX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Initialized = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_IUniswapV3PoolState *IUniswapV3PoolStateSession) Observations(index *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	return _IUniswapV3PoolState.Contract.Observations(&_IUniswapV3PoolState.CallOpts, index)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCallerSession) Observations(index *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	return _IUniswapV3PoolState.Contract.Observations(&_IUniswapV3PoolState.CallOpts, index)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCaller) Positions(opts *bind.CallOpts, key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _IUniswapV3PoolState.contract.Call(opts, &out, "positions", key)

	outstruct := new(struct {
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_IUniswapV3PoolState *IUniswapV3PoolStateSession) Positions(key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _IUniswapV3PoolState.Contract.Positions(&_IUniswapV3PoolState.CallOpts, key)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCallerSession) Positions(key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _IUniswapV3PoolState.Contract.Positions(&_IUniswapV3PoolState.CallOpts, key)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCaller) ProtocolFees(opts *bind.CallOpts) (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	var out []interface{}
	err := _IUniswapV3PoolState.contract.Call(opts, &out, "protocolFees")

	outstruct := new(struct {
		Token0 *big.Int
		Token1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Token1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_IUniswapV3PoolState *IUniswapV3PoolStateSession) ProtocolFees() (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	return _IUniswapV3PoolState.Contract.ProtocolFees(&_IUniswapV3PoolState.CallOpts)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCallerSession) ProtocolFees() (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	return _IUniswapV3PoolState.Contract.ProtocolFees(&_IUniswapV3PoolState.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCaller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _IUniswapV3PoolState.contract.Call(opts, &out, "slot0")

	outstruct := new(struct {
		SqrtPriceX96               *big.Int
		Tick                       *big.Int
		ObservationIndex           uint16
		ObservationCardinality     uint16
		ObservationCardinalityNext uint16
		FeeProtocol                uint8
		Unlocked                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPriceX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationCardinality = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationCardinalityNext = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.FeeProtocol = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_IUniswapV3PoolState *IUniswapV3PoolStateSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _IUniswapV3PoolState.Contract.Slot0(&_IUniswapV3PoolState.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _IUniswapV3PoolState.Contract.Slot0(&_IUniswapV3PoolState.CallOpts)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCaller) TickBitmap(opts *bind.CallOpts, wordPosition int16) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3PoolState.contract.Call(opts, &out, "tickBitmap", wordPosition)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_IUniswapV3PoolState *IUniswapV3PoolStateSession) TickBitmap(wordPosition int16) (*big.Int, error) {
	return _IUniswapV3PoolState.Contract.TickBitmap(&_IUniswapV3PoolState.CallOpts, wordPosition)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCallerSession) TickBitmap(wordPosition int16) (*big.Int, error) {
	return _IUniswapV3PoolState.Contract.TickBitmap(&_IUniswapV3PoolState.CallOpts, wordPosition)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCaller) Ticks(opts *bind.CallOpts, tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	var out []interface{}
	err := _IUniswapV3PoolState.contract.Call(opts, &out, "ticks", tick)

	outstruct := new(struct {
		LiquidityGross                 *big.Int
		LiquidityNet                   *big.Int
		FeeGrowthOutside0X128          *big.Int
		FeeGrowthOutside1X128          *big.Int
		TickCumulativeOutside          *big.Int
		SecondsPerLiquidityOutsideX128 *big.Int
		SecondsOutside                 uint32
		Initialized                    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityGross = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityNet = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside0X128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside1X128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickCumulativeOutside = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityOutsideX128 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.SecondsOutside = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.Initialized = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_IUniswapV3PoolState *IUniswapV3PoolStateSession) Ticks(tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _IUniswapV3PoolState.Contract.Ticks(&_IUniswapV3PoolState.CallOpts, tick)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_IUniswapV3PoolState *IUniswapV3PoolStateCallerSession) Ticks(tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _IUniswapV3PoolState.Contract.Ticks(&_IUniswapV3PoolState.CallOpts, tick)
}

// IUniswapV3SwapCallbackMetaData contains all meta data concerning the IUniswapV3SwapCallback contract.
var IUniswapV3SwapCallbackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IUniswapV3SwapCallbackABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3SwapCallbackMetaData.ABI instead.
var IUniswapV3SwapCallbackABI = IUniswapV3SwapCallbackMetaData.ABI

// IUniswapV3SwapCallback is an auto generated Go binding around an Ethereum contract.
type IUniswapV3SwapCallback struct {
	IUniswapV3SwapCallbackCaller     // Read-only binding to the contract
	IUniswapV3SwapCallbackTransactor // Write-only binding to the contract
	IUniswapV3SwapCallbackFilterer   // Log filterer for contract events
}

// IUniswapV3SwapCallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3SwapCallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3SwapCallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3SwapCallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3SwapCallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3SwapCallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3SwapCallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3SwapCallbackSession struct {
	Contract     *IUniswapV3SwapCallback // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IUniswapV3SwapCallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3SwapCallbackCallerSession struct {
	Contract *IUniswapV3SwapCallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IUniswapV3SwapCallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3SwapCallbackTransactorSession struct {
	Contract     *IUniswapV3SwapCallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IUniswapV3SwapCallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3SwapCallbackRaw struct {
	Contract *IUniswapV3SwapCallback // Generic contract binding to access the raw methods on
}

// IUniswapV3SwapCallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3SwapCallbackCallerRaw struct {
	Contract *IUniswapV3SwapCallbackCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3SwapCallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3SwapCallbackTransactorRaw struct {
	Contract *IUniswapV3SwapCallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3SwapCallback creates a new instance of IUniswapV3SwapCallback, bound to a specific deployed contract.
func NewIUniswapV3SwapCallback(address common.Address, backend bind.ContractBackend) (*IUniswapV3SwapCallback, error) {
	contract, err := bindIUniswapV3SwapCallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3SwapCallback{IUniswapV3SwapCallbackCaller: IUniswapV3SwapCallbackCaller{contract: contract}, IUniswapV3SwapCallbackTransactor: IUniswapV3SwapCallbackTransactor{contract: contract}, IUniswapV3SwapCallbackFilterer: IUniswapV3SwapCallbackFilterer{contract: contract}}, nil
}

// NewIUniswapV3SwapCallbackCaller creates a new read-only instance of IUniswapV3SwapCallback, bound to a specific deployed contract.
func NewIUniswapV3SwapCallbackCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3SwapCallbackCaller, error) {
	contract, err := bindIUniswapV3SwapCallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3SwapCallbackCaller{contract: contract}, nil
}

// NewIUniswapV3SwapCallbackTransactor creates a new write-only instance of IUniswapV3SwapCallback, bound to a specific deployed contract.
func NewIUniswapV3SwapCallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3SwapCallbackTransactor, error) {
	contract, err := bindIUniswapV3SwapCallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3SwapCallbackTransactor{contract: contract}, nil
}

// NewIUniswapV3SwapCallbackFilterer creates a new log filterer instance of IUniswapV3SwapCallback, bound to a specific deployed contract.
func NewIUniswapV3SwapCallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3SwapCallbackFilterer, error) {
	contract, err := bindIUniswapV3SwapCallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3SwapCallbackFilterer{contract: contract}, nil
}

// bindIUniswapV3SwapCallback binds a generic wrapper to an already deployed contract.
func bindIUniswapV3SwapCallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3SwapCallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3SwapCallback *IUniswapV3SwapCallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3SwapCallback.Contract.IUniswapV3SwapCallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3SwapCallback *IUniswapV3SwapCallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3SwapCallback.Contract.IUniswapV3SwapCallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3SwapCallback *IUniswapV3SwapCallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3SwapCallback.Contract.IUniswapV3SwapCallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3SwapCallback *IUniswapV3SwapCallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3SwapCallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3SwapCallback *IUniswapV3SwapCallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3SwapCallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3SwapCallback *IUniswapV3SwapCallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3SwapCallback.Contract.contract.Transact(opts, method, params...)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_IUniswapV3SwapCallback *IUniswapV3SwapCallbackTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3SwapCallback.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_IUniswapV3SwapCallback *IUniswapV3SwapCallbackSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3SwapCallback.Contract.UniswapV3SwapCallback(&_IUniswapV3SwapCallback.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_IUniswapV3SwapCallback *IUniswapV3SwapCallbackTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3SwapCallback.Contract.UniswapV3SwapCallback(&_IUniswapV3SwapCallback.TransactOpts, amount0Delta, amount1Delta, data)
}

// IWETH9MetaData contains all meta data concerning the IWETH9 contract.
var IWETH9MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IWETH9ABI is the input ABI used to generate the binding from.
// Deprecated: Use IWETH9MetaData.ABI instead.
var IWETH9ABI = IWETH9MetaData.ABI

// IWETH9 is an auto generated Go binding around an Ethereum contract.
type IWETH9 struct {
	IWETH9Caller     // Read-only binding to the contract
	IWETH9Transactor // Write-only binding to the contract
	IWETH9Filterer   // Log filterer for contract events
}

// IWETH9Caller is an auto generated read-only Go binding around an Ethereum contract.
type IWETH9Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETH9Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETH9Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETH9Session struct {
	Contract     *IWETH9           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETH9CallerSession struct {
	Contract *IWETH9Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWETH9TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETH9TransactorSession struct {
	Contract     *IWETH9Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9Raw is an auto generated low-level Go binding around an Ethereum contract.
type IWETH9Raw struct {
	Contract *IWETH9 // Generic contract binding to access the raw methods on
}

// IWETH9CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETH9CallerRaw struct {
	Contract *IWETH9Caller // Generic read-only contract binding to access the raw methods on
}

// IWETH9TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETH9TransactorRaw struct {
	Contract *IWETH9Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH9 creates a new instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9(address common.Address, backend bind.ContractBackend) (*IWETH9, error) {
	contract, err := bindIWETH9(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH9{IWETH9Caller: IWETH9Caller{contract: contract}, IWETH9Transactor: IWETH9Transactor{contract: contract}, IWETH9Filterer: IWETH9Filterer{contract: contract}}, nil
}

// NewIWETH9Caller creates a new read-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Caller(address common.Address, caller bind.ContractCaller) (*IWETH9Caller, error) {
	contract, err := bindIWETH9(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Caller{contract: contract}, nil
}

// NewIWETH9Transactor creates a new write-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Transactor(address common.Address, transactor bind.ContractTransactor) (*IWETH9Transactor, error) {
	contract, err := bindIWETH9(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Transactor{contract: contract}, nil
}

// NewIWETH9Filterer creates a new log filterer instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Filterer(address common.Address, filterer bind.ContractFilterer) (*IWETH9Filterer, error) {
	contract, err := bindIWETH9(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETH9Filterer{contract: contract}, nil
}

// bindIWETH9 binds a generic wrapper to an already deployed contract.
func bindIWETH9(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWETH9ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.IWETH9Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH9 *IWETH9Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH9 *IWETH9Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH9.Contract.Allowance(&_IWETH9.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH9.Contract.Allowance(&_IWETH9.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH9 *IWETH9Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH9 *IWETH9Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWETH9.Contract.BalanceOf(&_IWETH9.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWETH9.Contract.BalanceOf(&_IWETH9.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9Session) TotalSupply() (*big.Int, error) {
	return _IWETH9.Contract.TotalSupply(&_IWETH9.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) TotalSupply() (*big.Int, error) {
	return _IWETH9.Contract.TotalSupply(&_IWETH9.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Approve(&_IWETH9.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Approve(&_IWETH9.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Session) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9TransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Transfer(&_IWETH9.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Transfer(&_IWETH9.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.TransferFrom(&_IWETH9.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.TransferFrom(&_IWETH9.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH9 *IWETH9Transactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH9 *IWETH9Session) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH9 *IWETH9TransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, arg0)
}

// IWETH9ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IWETH9 contract.
type IWETH9ApprovalIterator struct {
	Event *IWETH9Approval // Event containing the contract specifics and raw log

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
func (it *IWETH9ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETH9Approval)
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
		it.Event = new(IWETH9Approval)
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
func (it *IWETH9ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETH9ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETH9Approval represents a Approval event raised by the IWETH9 contract.
type IWETH9Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH9 *IWETH9Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IWETH9ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH9.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IWETH9ApprovalIterator{contract: _IWETH9.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH9 *IWETH9Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IWETH9Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH9.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETH9Approval)
				if err := _IWETH9.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH9 *IWETH9Filterer) ParseApproval(log types.Log) (*IWETH9Approval, error) {
	event := new(IWETH9Approval)
	if err := _IWETH9.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWETH9TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IWETH9 contract.
type IWETH9TransferIterator struct {
	Event *IWETH9Transfer // Event containing the contract specifics and raw log

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
func (it *IWETH9TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETH9Transfer)
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
		it.Event = new(IWETH9Transfer)
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
func (it *IWETH9TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETH9TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETH9Transfer represents a Transfer event raised by the IWETH9 contract.
type IWETH9Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH9 *IWETH9Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IWETH9TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH9.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IWETH9TransferIterator{contract: _IWETH9.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH9 *IWETH9Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IWETH9Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH9.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETH9Transfer)
				if err := _IWETH9.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH9 *IWETH9Filterer) ParseTransfer(log types.Log) (*IWETH9Transfer, error) {
	event := new(IWETH9Transfer)
	if err := _IWETH9.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LowGasSafeMathMetaData contains all meta data concerning the LowGasSafeMath contract.
var LowGasSafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205f61f204c2ece836448b0163c5a5b8f737e2ad52551dd736f1789760fb933d3664736f6c63430007060033",
}

// LowGasSafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use LowGasSafeMathMetaData.ABI instead.
var LowGasSafeMathABI = LowGasSafeMathMetaData.ABI

// LowGasSafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LowGasSafeMathMetaData.Bin instead.
var LowGasSafeMathBin = LowGasSafeMathMetaData.Bin

// DeployLowGasSafeMath deploys a new Ethereum contract, binding an instance of LowGasSafeMath to it.
func DeployLowGasSafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LowGasSafeMath, error) {
	parsed, err := LowGasSafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LowGasSafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LowGasSafeMath{LowGasSafeMathCaller: LowGasSafeMathCaller{contract: contract}, LowGasSafeMathTransactor: LowGasSafeMathTransactor{contract: contract}, LowGasSafeMathFilterer: LowGasSafeMathFilterer{contract: contract}}, nil
}

// LowGasSafeMath is an auto generated Go binding around an Ethereum contract.
type LowGasSafeMath struct {
	LowGasSafeMathCaller     // Read-only binding to the contract
	LowGasSafeMathTransactor // Write-only binding to the contract
	LowGasSafeMathFilterer   // Log filterer for contract events
}

// LowGasSafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type LowGasSafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LowGasSafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LowGasSafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LowGasSafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LowGasSafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LowGasSafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LowGasSafeMathSession struct {
	Contract     *LowGasSafeMath   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LowGasSafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LowGasSafeMathCallerSession struct {
	Contract *LowGasSafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LowGasSafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LowGasSafeMathTransactorSession struct {
	Contract     *LowGasSafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LowGasSafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type LowGasSafeMathRaw struct {
	Contract *LowGasSafeMath // Generic contract binding to access the raw methods on
}

// LowGasSafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LowGasSafeMathCallerRaw struct {
	Contract *LowGasSafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// LowGasSafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LowGasSafeMathTransactorRaw struct {
	Contract *LowGasSafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLowGasSafeMath creates a new instance of LowGasSafeMath, bound to a specific deployed contract.
func NewLowGasSafeMath(address common.Address, backend bind.ContractBackend) (*LowGasSafeMath, error) {
	contract, err := bindLowGasSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LowGasSafeMath{LowGasSafeMathCaller: LowGasSafeMathCaller{contract: contract}, LowGasSafeMathTransactor: LowGasSafeMathTransactor{contract: contract}, LowGasSafeMathFilterer: LowGasSafeMathFilterer{contract: contract}}, nil
}

// NewLowGasSafeMathCaller creates a new read-only instance of LowGasSafeMath, bound to a specific deployed contract.
func NewLowGasSafeMathCaller(address common.Address, caller bind.ContractCaller) (*LowGasSafeMathCaller, error) {
	contract, err := bindLowGasSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LowGasSafeMathCaller{contract: contract}, nil
}

// NewLowGasSafeMathTransactor creates a new write-only instance of LowGasSafeMath, bound to a specific deployed contract.
func NewLowGasSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*LowGasSafeMathTransactor, error) {
	contract, err := bindLowGasSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LowGasSafeMathTransactor{contract: contract}, nil
}

// NewLowGasSafeMathFilterer creates a new log filterer instance of LowGasSafeMath, bound to a specific deployed contract.
func NewLowGasSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*LowGasSafeMathFilterer, error) {
	contract, err := bindLowGasSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LowGasSafeMathFilterer{contract: contract}, nil
}

// bindLowGasSafeMath binds a generic wrapper to an already deployed contract.
func bindLowGasSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LowGasSafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LowGasSafeMath *LowGasSafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LowGasSafeMath.Contract.LowGasSafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LowGasSafeMath *LowGasSafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LowGasSafeMath.Contract.LowGasSafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LowGasSafeMath *LowGasSafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LowGasSafeMath.Contract.LowGasSafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LowGasSafeMath *LowGasSafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LowGasSafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LowGasSafeMath *LowGasSafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LowGasSafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LowGasSafeMath *LowGasSafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LowGasSafeMath.Contract.contract.Transact(opts, method, params...)
}

// MulticallMetaData contains all meta data concerning the Multicall contract.
var MulticallMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MulticallABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallMetaData.ABI instead.
var MulticallABI = MulticallMetaData.ABI

// Multicall is an auto generated Go binding around an Ethereum contract.
type Multicall struct {
	MulticallCaller     // Read-only binding to the contract
	MulticallTransactor // Write-only binding to the contract
	MulticallFilterer   // Log filterer for contract events
}

// MulticallCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallSession struct {
	Contract     *Multicall        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallCallerSession struct {
	Contract *MulticallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MulticallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallTransactorSession struct {
	Contract     *MulticallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MulticallRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallRaw struct {
	Contract *Multicall // Generic contract binding to access the raw methods on
}

// MulticallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallCallerRaw struct {
	Contract *MulticallCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallTransactorRaw struct {
	Contract *MulticallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticall creates a new instance of Multicall, bound to a specific deployed contract.
func NewMulticall(address common.Address, backend bind.ContractBackend) (*Multicall, error) {
	contract, err := bindMulticall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicall{MulticallCaller: MulticallCaller{contract: contract}, MulticallTransactor: MulticallTransactor{contract: contract}, MulticallFilterer: MulticallFilterer{contract: contract}}, nil
}

// NewMulticallCaller creates a new read-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallCaller(address common.Address, caller bind.ContractCaller) (*MulticallCaller, error) {
	contract, err := bindMulticall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallCaller{contract: contract}, nil
}

// NewMulticallTransactor creates a new write-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallTransactor, error) {
	contract, err := bindMulticall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallTransactor{contract: contract}, nil
}

// NewMulticallFilterer creates a new log filterer instance of Multicall, bound to a specific deployed contract.
func NewMulticallFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallFilterer, error) {
	contract, err := bindMulticall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallFilterer{contract: contract}, nil
}

// bindMulticall binds a generic wrapper to an already deployed contract.
func bindMulticall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MulticallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.MulticallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transact(opts, method, params...)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Multicall *MulticallTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Multicall.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Multicall *MulticallSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Multicall.Contract.Multicall(&_Multicall.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Multicall *MulticallTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Multicall.Contract.Multicall(&_Multicall.TransactOpts, data)
}

// PathMetaData contains all meta data concerning the Path contract.
var PathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a176baea18ef9aa4509b1365ef652b92b6f3138fc8f6645b707ce78bd1eb4d9e64736f6c63430007060033",
}

// PathABI is the input ABI used to generate the binding from.
// Deprecated: Use PathMetaData.ABI instead.
var PathABI = PathMetaData.ABI

// PathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PathMetaData.Bin instead.
var PathBin = PathMetaData.Bin

// DeployPath deploys a new Ethereum contract, binding an instance of Path to it.
func DeployPath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Path, error) {
	parsed, err := PathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Path{PathCaller: PathCaller{contract: contract}, PathTransactor: PathTransactor{contract: contract}, PathFilterer: PathFilterer{contract: contract}}, nil
}

// Path is an auto generated Go binding around an Ethereum contract.
type Path struct {
	PathCaller     // Read-only binding to the contract
	PathTransactor // Write-only binding to the contract
	PathFilterer   // Log filterer for contract events
}

// PathCaller is an auto generated read-only Go binding around an Ethereum contract.
type PathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PathSession struct {
	Contract     *Path             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PathCallerSession struct {
	Contract *PathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PathTransactorSession struct {
	Contract     *PathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PathRaw is an auto generated low-level Go binding around an Ethereum contract.
type PathRaw struct {
	Contract *Path // Generic contract binding to access the raw methods on
}

// PathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PathCallerRaw struct {
	Contract *PathCaller // Generic read-only contract binding to access the raw methods on
}

// PathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PathTransactorRaw struct {
	Contract *PathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPath creates a new instance of Path, bound to a specific deployed contract.
func NewPath(address common.Address, backend bind.ContractBackend) (*Path, error) {
	contract, err := bindPath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Path{PathCaller: PathCaller{contract: contract}, PathTransactor: PathTransactor{contract: contract}, PathFilterer: PathFilterer{contract: contract}}, nil
}

// NewPathCaller creates a new read-only instance of Path, bound to a specific deployed contract.
func NewPathCaller(address common.Address, caller bind.ContractCaller) (*PathCaller, error) {
	contract, err := bindPath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PathCaller{contract: contract}, nil
}

// NewPathTransactor creates a new write-only instance of Path, bound to a specific deployed contract.
func NewPathTransactor(address common.Address, transactor bind.ContractTransactor) (*PathTransactor, error) {
	contract, err := bindPath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PathTransactor{contract: contract}, nil
}

// NewPathFilterer creates a new log filterer instance of Path, bound to a specific deployed contract.
func NewPathFilterer(address common.Address, filterer bind.ContractFilterer) (*PathFilterer, error) {
	contract, err := bindPath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PathFilterer{contract: contract}, nil
}

// bindPath binds a generic wrapper to an already deployed contract.
func bindPath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Path *PathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Path.Contract.PathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Path *PathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Path.Contract.PathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Path *PathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Path.Contract.PathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Path *PathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Path.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Path *PathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Path.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Path *PathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Path.Contract.contract.Transact(opts, method, params...)
}

// PeripheryImmutableStateMetaData contains all meta data concerning the PeripheryImmutableState contract.
var PeripheryImmutableStateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PeripheryImmutableStateABI is the input ABI used to generate the binding from.
// Deprecated: Use PeripheryImmutableStateMetaData.ABI instead.
var PeripheryImmutableStateABI = PeripheryImmutableStateMetaData.ABI

// PeripheryImmutableState is an auto generated Go binding around an Ethereum contract.
type PeripheryImmutableState struct {
	PeripheryImmutableStateCaller     // Read-only binding to the contract
	PeripheryImmutableStateTransactor // Write-only binding to the contract
	PeripheryImmutableStateFilterer   // Log filterer for contract events
}

// PeripheryImmutableStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeripheryImmutableStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryImmutableStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeripheryImmutableStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryImmutableStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeripheryImmutableStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryImmutableStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeripheryImmutableStateSession struct {
	Contract     *PeripheryImmutableState // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PeripheryImmutableStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeripheryImmutableStateCallerSession struct {
	Contract *PeripheryImmutableStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// PeripheryImmutableStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeripheryImmutableStateTransactorSession struct {
	Contract     *PeripheryImmutableStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// PeripheryImmutableStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeripheryImmutableStateRaw struct {
	Contract *PeripheryImmutableState // Generic contract binding to access the raw methods on
}

// PeripheryImmutableStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeripheryImmutableStateCallerRaw struct {
	Contract *PeripheryImmutableStateCaller // Generic read-only contract binding to access the raw methods on
}

// PeripheryImmutableStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeripheryImmutableStateTransactorRaw struct {
	Contract *PeripheryImmutableStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeripheryImmutableState creates a new instance of PeripheryImmutableState, bound to a specific deployed contract.
func NewPeripheryImmutableState(address common.Address, backend bind.ContractBackend) (*PeripheryImmutableState, error) {
	contract, err := bindPeripheryImmutableState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PeripheryImmutableState{PeripheryImmutableStateCaller: PeripheryImmutableStateCaller{contract: contract}, PeripheryImmutableStateTransactor: PeripheryImmutableStateTransactor{contract: contract}, PeripheryImmutableStateFilterer: PeripheryImmutableStateFilterer{contract: contract}}, nil
}

// NewPeripheryImmutableStateCaller creates a new read-only instance of PeripheryImmutableState, bound to a specific deployed contract.
func NewPeripheryImmutableStateCaller(address common.Address, caller bind.ContractCaller) (*PeripheryImmutableStateCaller, error) {
	contract, err := bindPeripheryImmutableState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryImmutableStateCaller{contract: contract}, nil
}

// NewPeripheryImmutableStateTransactor creates a new write-only instance of PeripheryImmutableState, bound to a specific deployed contract.
func NewPeripheryImmutableStateTransactor(address common.Address, transactor bind.ContractTransactor) (*PeripheryImmutableStateTransactor, error) {
	contract, err := bindPeripheryImmutableState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryImmutableStateTransactor{contract: contract}, nil
}

// NewPeripheryImmutableStateFilterer creates a new log filterer instance of PeripheryImmutableState, bound to a specific deployed contract.
func NewPeripheryImmutableStateFilterer(address common.Address, filterer bind.ContractFilterer) (*PeripheryImmutableStateFilterer, error) {
	contract, err := bindPeripheryImmutableState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeripheryImmutableStateFilterer{contract: contract}, nil
}

// bindPeripheryImmutableState binds a generic wrapper to an already deployed contract.
func bindPeripheryImmutableState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PeripheryImmutableStateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeripheryImmutableState *PeripheryImmutableStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeripheryImmutableState.Contract.PeripheryImmutableStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeripheryImmutableState *PeripheryImmutableStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryImmutableState.Contract.PeripheryImmutableStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeripheryImmutableState *PeripheryImmutableStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeripheryImmutableState.Contract.PeripheryImmutableStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeripheryImmutableState *PeripheryImmutableStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeripheryImmutableState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeripheryImmutableState *PeripheryImmutableStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryImmutableState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeripheryImmutableState *PeripheryImmutableStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeripheryImmutableState.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PeripheryImmutableState.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateSession) WETH9() (common.Address, error) {
	return _PeripheryImmutableState.Contract.WETH9(&_PeripheryImmutableState.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateCallerSession) WETH9() (common.Address, error) {
	return _PeripheryImmutableState.Contract.WETH9(&_PeripheryImmutableState.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PeripheryImmutableState.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateSession) Factory() (common.Address, error) {
	return _PeripheryImmutableState.Contract.Factory(&_PeripheryImmutableState.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateCallerSession) Factory() (common.Address, error) {
	return _PeripheryImmutableState.Contract.Factory(&_PeripheryImmutableState.CallOpts)
}

// PeripheryPaymentsMetaData contains all meta data concerning the PeripheryPayments contract.
var PeripheryPaymentsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PeripheryPaymentsABI is the input ABI used to generate the binding from.
// Deprecated: Use PeripheryPaymentsMetaData.ABI instead.
var PeripheryPaymentsABI = PeripheryPaymentsMetaData.ABI

// PeripheryPayments is an auto generated Go binding around an Ethereum contract.
type PeripheryPayments struct {
	PeripheryPaymentsCaller     // Read-only binding to the contract
	PeripheryPaymentsTransactor // Write-only binding to the contract
	PeripheryPaymentsFilterer   // Log filterer for contract events
}

// PeripheryPaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeripheryPaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryPaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeripheryPaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryPaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeripheryPaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryPaymentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeripheryPaymentsSession struct {
	Contract     *PeripheryPayments // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PeripheryPaymentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeripheryPaymentsCallerSession struct {
	Contract *PeripheryPaymentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PeripheryPaymentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeripheryPaymentsTransactorSession struct {
	Contract     *PeripheryPaymentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PeripheryPaymentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeripheryPaymentsRaw struct {
	Contract *PeripheryPayments // Generic contract binding to access the raw methods on
}

// PeripheryPaymentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeripheryPaymentsCallerRaw struct {
	Contract *PeripheryPaymentsCaller // Generic read-only contract binding to access the raw methods on
}

// PeripheryPaymentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeripheryPaymentsTransactorRaw struct {
	Contract *PeripheryPaymentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeripheryPayments creates a new instance of PeripheryPayments, bound to a specific deployed contract.
func NewPeripheryPayments(address common.Address, backend bind.ContractBackend) (*PeripheryPayments, error) {
	contract, err := bindPeripheryPayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PeripheryPayments{PeripheryPaymentsCaller: PeripheryPaymentsCaller{contract: contract}, PeripheryPaymentsTransactor: PeripheryPaymentsTransactor{contract: contract}, PeripheryPaymentsFilterer: PeripheryPaymentsFilterer{contract: contract}}, nil
}

// NewPeripheryPaymentsCaller creates a new read-only instance of PeripheryPayments, bound to a specific deployed contract.
func NewPeripheryPaymentsCaller(address common.Address, caller bind.ContractCaller) (*PeripheryPaymentsCaller, error) {
	contract, err := bindPeripheryPayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryPaymentsCaller{contract: contract}, nil
}

// NewPeripheryPaymentsTransactor creates a new write-only instance of PeripheryPayments, bound to a specific deployed contract.
func NewPeripheryPaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*PeripheryPaymentsTransactor, error) {
	contract, err := bindPeripheryPayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryPaymentsTransactor{contract: contract}, nil
}

// NewPeripheryPaymentsFilterer creates a new log filterer instance of PeripheryPayments, bound to a specific deployed contract.
func NewPeripheryPaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*PeripheryPaymentsFilterer, error) {
	contract, err := bindPeripheryPayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeripheryPaymentsFilterer{contract: contract}, nil
}

// bindPeripheryPayments binds a generic wrapper to an already deployed contract.
func bindPeripheryPayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PeripheryPaymentsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeripheryPayments *PeripheryPaymentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeripheryPayments.Contract.PeripheryPaymentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeripheryPayments *PeripheryPaymentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryPayments.Contract.PeripheryPaymentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeripheryPayments *PeripheryPaymentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeripheryPayments.Contract.PeripheryPaymentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeripheryPayments *PeripheryPaymentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeripheryPayments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeripheryPayments *PeripheryPaymentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryPayments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeripheryPayments *PeripheryPaymentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeripheryPayments.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryPayments *PeripheryPaymentsCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PeripheryPayments.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryPayments *PeripheryPaymentsSession) WETH9() (common.Address, error) {
	return _PeripheryPayments.Contract.WETH9(&_PeripheryPayments.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryPayments *PeripheryPaymentsCallerSession) WETH9() (common.Address, error) {
	return _PeripheryPayments.Contract.WETH9(&_PeripheryPayments.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryPayments *PeripheryPaymentsCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PeripheryPayments.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryPayments *PeripheryPaymentsSession) Factory() (common.Address, error) {
	return _PeripheryPayments.Contract.Factory(&_PeripheryPayments.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryPayments *PeripheryPaymentsCallerSession) Factory() (common.Address, error) {
	return _PeripheryPayments.Contract.Factory(&_PeripheryPayments.CallOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_PeripheryPayments *PeripheryPaymentsTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryPayments.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_PeripheryPayments *PeripheryPaymentsSession) RefundETH() (*types.Transaction, error) {
	return _PeripheryPayments.Contract.RefundETH(&_PeripheryPayments.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_PeripheryPayments *PeripheryPaymentsTransactorSession) RefundETH() (*types.Transaction, error) {
	return _PeripheryPayments.Contract.RefundETH(&_PeripheryPayments.TransactOpts)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPayments *PeripheryPaymentsTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPayments.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPayments *PeripheryPaymentsSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPayments.Contract.SweepToken(&_PeripheryPayments.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPayments *PeripheryPaymentsTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPayments.Contract.SweepToken(&_PeripheryPayments.TransactOpts, token, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPayments *PeripheryPaymentsTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPayments.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPayments *PeripheryPaymentsSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPayments.Contract.UnwrapWETH9(&_PeripheryPayments.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPayments *PeripheryPaymentsTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPayments.Contract.UnwrapWETH9(&_PeripheryPayments.TransactOpts, amountMinimum, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PeripheryPayments *PeripheryPaymentsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryPayments.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PeripheryPayments *PeripheryPaymentsSession) Receive() (*types.Transaction, error) {
	return _PeripheryPayments.Contract.Receive(&_PeripheryPayments.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PeripheryPayments *PeripheryPaymentsTransactorSession) Receive() (*types.Transaction, error) {
	return _PeripheryPayments.Contract.Receive(&_PeripheryPayments.TransactOpts)
}

// PeripheryPaymentsWithFeeMetaData contains all meta data concerning the PeripheryPaymentsWithFee contract.
var PeripheryPaymentsWithFeeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PeripheryPaymentsWithFeeABI is the input ABI used to generate the binding from.
// Deprecated: Use PeripheryPaymentsWithFeeMetaData.ABI instead.
var PeripheryPaymentsWithFeeABI = PeripheryPaymentsWithFeeMetaData.ABI

// PeripheryPaymentsWithFee is an auto generated Go binding around an Ethereum contract.
type PeripheryPaymentsWithFee struct {
	PeripheryPaymentsWithFeeCaller     // Read-only binding to the contract
	PeripheryPaymentsWithFeeTransactor // Write-only binding to the contract
	PeripheryPaymentsWithFeeFilterer   // Log filterer for contract events
}

// PeripheryPaymentsWithFeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeripheryPaymentsWithFeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryPaymentsWithFeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeripheryPaymentsWithFeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryPaymentsWithFeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeripheryPaymentsWithFeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryPaymentsWithFeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeripheryPaymentsWithFeeSession struct {
	Contract     *PeripheryPaymentsWithFee // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PeripheryPaymentsWithFeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeripheryPaymentsWithFeeCallerSession struct {
	Contract *PeripheryPaymentsWithFeeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// PeripheryPaymentsWithFeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeripheryPaymentsWithFeeTransactorSession struct {
	Contract     *PeripheryPaymentsWithFeeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PeripheryPaymentsWithFeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeripheryPaymentsWithFeeRaw struct {
	Contract *PeripheryPaymentsWithFee // Generic contract binding to access the raw methods on
}

// PeripheryPaymentsWithFeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeripheryPaymentsWithFeeCallerRaw struct {
	Contract *PeripheryPaymentsWithFeeCaller // Generic read-only contract binding to access the raw methods on
}

// PeripheryPaymentsWithFeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeripheryPaymentsWithFeeTransactorRaw struct {
	Contract *PeripheryPaymentsWithFeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeripheryPaymentsWithFee creates a new instance of PeripheryPaymentsWithFee, bound to a specific deployed contract.
func NewPeripheryPaymentsWithFee(address common.Address, backend bind.ContractBackend) (*PeripheryPaymentsWithFee, error) {
	contract, err := bindPeripheryPaymentsWithFee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PeripheryPaymentsWithFee{PeripheryPaymentsWithFeeCaller: PeripheryPaymentsWithFeeCaller{contract: contract}, PeripheryPaymentsWithFeeTransactor: PeripheryPaymentsWithFeeTransactor{contract: contract}, PeripheryPaymentsWithFeeFilterer: PeripheryPaymentsWithFeeFilterer{contract: contract}}, nil
}

// NewPeripheryPaymentsWithFeeCaller creates a new read-only instance of PeripheryPaymentsWithFee, bound to a specific deployed contract.
func NewPeripheryPaymentsWithFeeCaller(address common.Address, caller bind.ContractCaller) (*PeripheryPaymentsWithFeeCaller, error) {
	contract, err := bindPeripheryPaymentsWithFee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryPaymentsWithFeeCaller{contract: contract}, nil
}

// NewPeripheryPaymentsWithFeeTransactor creates a new write-only instance of PeripheryPaymentsWithFee, bound to a specific deployed contract.
func NewPeripheryPaymentsWithFeeTransactor(address common.Address, transactor bind.ContractTransactor) (*PeripheryPaymentsWithFeeTransactor, error) {
	contract, err := bindPeripheryPaymentsWithFee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryPaymentsWithFeeTransactor{contract: contract}, nil
}

// NewPeripheryPaymentsWithFeeFilterer creates a new log filterer instance of PeripheryPaymentsWithFee, bound to a specific deployed contract.
func NewPeripheryPaymentsWithFeeFilterer(address common.Address, filterer bind.ContractFilterer) (*PeripheryPaymentsWithFeeFilterer, error) {
	contract, err := bindPeripheryPaymentsWithFee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeripheryPaymentsWithFeeFilterer{contract: contract}, nil
}

// bindPeripheryPaymentsWithFee binds a generic wrapper to an already deployed contract.
func bindPeripheryPaymentsWithFee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PeripheryPaymentsWithFeeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeripheryPaymentsWithFee.Contract.PeripheryPaymentsWithFeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.PeripheryPaymentsWithFeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.PeripheryPaymentsWithFeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeripheryPaymentsWithFee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PeripheryPaymentsWithFee.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeSession) WETH9() (common.Address, error) {
	return _PeripheryPaymentsWithFee.Contract.WETH9(&_PeripheryPaymentsWithFee.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeCallerSession) WETH9() (common.Address, error) {
	return _PeripheryPaymentsWithFee.Contract.WETH9(&_PeripheryPaymentsWithFee.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PeripheryPaymentsWithFee.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeSession) Factory() (common.Address, error) {
	return _PeripheryPaymentsWithFee.Contract.Factory(&_PeripheryPaymentsWithFee.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeCallerSession) Factory() (common.Address, error) {
	return _PeripheryPaymentsWithFee.Contract.Factory(&_PeripheryPaymentsWithFee.CallOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeSession) RefundETH() (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.RefundETH(&_PeripheryPaymentsWithFee.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactorSession) RefundETH() (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.RefundETH(&_PeripheryPaymentsWithFee.TransactOpts)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.SweepToken(&_PeripheryPaymentsWithFee.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.SweepToken(&_PeripheryPaymentsWithFee.TransactOpts, token, amountMinimum, recipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.SweepTokenWithFee(&_PeripheryPaymentsWithFee.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.SweepTokenWithFee(&_PeripheryPaymentsWithFee.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.UnwrapWETH9(&_PeripheryPaymentsWithFee.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.UnwrapWETH9(&_PeripheryPaymentsWithFee.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactor) UnwrapWETH9WithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.contract.Transact(opts, "unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.UnwrapWETH9WithFee(&_PeripheryPaymentsWithFee.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactorSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.UnwrapWETH9WithFee(&_PeripheryPaymentsWithFee.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeSession) Receive() (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.Receive(&_PeripheryPaymentsWithFee.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PeripheryPaymentsWithFee *PeripheryPaymentsWithFeeTransactorSession) Receive() (*types.Transaction, error) {
	return _PeripheryPaymentsWithFee.Contract.Receive(&_PeripheryPaymentsWithFee.TransactOpts)
}

// PeripheryValidationMetaData contains all meta data concerning the PeripheryValidation contract.
var PeripheryValidationMetaData = &bind.MetaData{
	ABI: "[]",
}

// PeripheryValidationABI is the input ABI used to generate the binding from.
// Deprecated: Use PeripheryValidationMetaData.ABI instead.
var PeripheryValidationABI = PeripheryValidationMetaData.ABI

// PeripheryValidation is an auto generated Go binding around an Ethereum contract.
type PeripheryValidation struct {
	PeripheryValidationCaller     // Read-only binding to the contract
	PeripheryValidationTransactor // Write-only binding to the contract
	PeripheryValidationFilterer   // Log filterer for contract events
}

// PeripheryValidationCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeripheryValidationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryValidationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeripheryValidationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryValidationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeripheryValidationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryValidationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeripheryValidationSession struct {
	Contract     *PeripheryValidation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PeripheryValidationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeripheryValidationCallerSession struct {
	Contract *PeripheryValidationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PeripheryValidationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeripheryValidationTransactorSession struct {
	Contract     *PeripheryValidationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PeripheryValidationRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeripheryValidationRaw struct {
	Contract *PeripheryValidation // Generic contract binding to access the raw methods on
}

// PeripheryValidationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeripheryValidationCallerRaw struct {
	Contract *PeripheryValidationCaller // Generic read-only contract binding to access the raw methods on
}

// PeripheryValidationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeripheryValidationTransactorRaw struct {
	Contract *PeripheryValidationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeripheryValidation creates a new instance of PeripheryValidation, bound to a specific deployed contract.
func NewPeripheryValidation(address common.Address, backend bind.ContractBackend) (*PeripheryValidation, error) {
	contract, err := bindPeripheryValidation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PeripheryValidation{PeripheryValidationCaller: PeripheryValidationCaller{contract: contract}, PeripheryValidationTransactor: PeripheryValidationTransactor{contract: contract}, PeripheryValidationFilterer: PeripheryValidationFilterer{contract: contract}}, nil
}

// NewPeripheryValidationCaller creates a new read-only instance of PeripheryValidation, bound to a specific deployed contract.
func NewPeripheryValidationCaller(address common.Address, caller bind.ContractCaller) (*PeripheryValidationCaller, error) {
	contract, err := bindPeripheryValidation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryValidationCaller{contract: contract}, nil
}

// NewPeripheryValidationTransactor creates a new write-only instance of PeripheryValidation, bound to a specific deployed contract.
func NewPeripheryValidationTransactor(address common.Address, transactor bind.ContractTransactor) (*PeripheryValidationTransactor, error) {
	contract, err := bindPeripheryValidation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryValidationTransactor{contract: contract}, nil
}

// NewPeripheryValidationFilterer creates a new log filterer instance of PeripheryValidation, bound to a specific deployed contract.
func NewPeripheryValidationFilterer(address common.Address, filterer bind.ContractFilterer) (*PeripheryValidationFilterer, error) {
	contract, err := bindPeripheryValidation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeripheryValidationFilterer{contract: contract}, nil
}

// bindPeripheryValidation binds a generic wrapper to an already deployed contract.
func bindPeripheryValidation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PeripheryValidationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeripheryValidation *PeripheryValidationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeripheryValidation.Contract.PeripheryValidationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeripheryValidation *PeripheryValidationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryValidation.Contract.PeripheryValidationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeripheryValidation *PeripheryValidationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeripheryValidation.Contract.PeripheryValidationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeripheryValidation *PeripheryValidationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeripheryValidation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeripheryValidation *PeripheryValidationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryValidation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeripheryValidation *PeripheryValidationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeripheryValidation.Contract.contract.Transact(opts, method, params...)
}

// PoolAddressMetaData contains all meta data concerning the PoolAddress contract.
var PoolAddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220011eb8d76e96af7a779928ddc9a6f70a5bac8bfc3de5b91e106f6e47089da10e64736f6c63430007060033",
}

// PoolAddressABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolAddressMetaData.ABI instead.
var PoolAddressABI = PoolAddressMetaData.ABI

// PoolAddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolAddressMetaData.Bin instead.
var PoolAddressBin = PoolAddressMetaData.Bin

// DeployPoolAddress deploys a new Ethereum contract, binding an instance of PoolAddress to it.
func DeployPoolAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PoolAddress, error) {
	parsed, err := PoolAddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolAddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolAddress{PoolAddressCaller: PoolAddressCaller{contract: contract}, PoolAddressTransactor: PoolAddressTransactor{contract: contract}, PoolAddressFilterer: PoolAddressFilterer{contract: contract}}, nil
}

// PoolAddress is an auto generated Go binding around an Ethereum contract.
type PoolAddress struct {
	PoolAddressCaller     // Read-only binding to the contract
	PoolAddressTransactor // Write-only binding to the contract
	PoolAddressFilterer   // Log filterer for contract events
}

// PoolAddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolAddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolAddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolAddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolAddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolAddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolAddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolAddressSession struct {
	Contract     *PoolAddress      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolAddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolAddressCallerSession struct {
	Contract *PoolAddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PoolAddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolAddressTransactorSession struct {
	Contract     *PoolAddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoolAddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolAddressRaw struct {
	Contract *PoolAddress // Generic contract binding to access the raw methods on
}

// PoolAddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolAddressCallerRaw struct {
	Contract *PoolAddressCaller // Generic read-only contract binding to access the raw methods on
}

// PoolAddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolAddressTransactorRaw struct {
	Contract *PoolAddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolAddress creates a new instance of PoolAddress, bound to a specific deployed contract.
func NewPoolAddress(address common.Address, backend bind.ContractBackend) (*PoolAddress, error) {
	contract, err := bindPoolAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolAddress{PoolAddressCaller: PoolAddressCaller{contract: contract}, PoolAddressTransactor: PoolAddressTransactor{contract: contract}, PoolAddressFilterer: PoolAddressFilterer{contract: contract}}, nil
}

// NewPoolAddressCaller creates a new read-only instance of PoolAddress, bound to a specific deployed contract.
func NewPoolAddressCaller(address common.Address, caller bind.ContractCaller) (*PoolAddressCaller, error) {
	contract, err := bindPoolAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolAddressCaller{contract: contract}, nil
}

// NewPoolAddressTransactor creates a new write-only instance of PoolAddress, bound to a specific deployed contract.
func NewPoolAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolAddressTransactor, error) {
	contract, err := bindPoolAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolAddressTransactor{contract: contract}, nil
}

// NewPoolAddressFilterer creates a new log filterer instance of PoolAddress, bound to a specific deployed contract.
func NewPoolAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolAddressFilterer, error) {
	contract, err := bindPoolAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolAddressFilterer{contract: contract}, nil
}

// bindPoolAddress binds a generic wrapper to an already deployed contract.
func bindPoolAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolAddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolAddress *PoolAddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolAddress.Contract.PoolAddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolAddress *PoolAddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolAddress.Contract.PoolAddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolAddress *PoolAddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolAddress.Contract.PoolAddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolAddress *PoolAddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolAddress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolAddress *PoolAddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolAddress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolAddress *PoolAddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolAddress.Contract.contract.Transact(opts, method, params...)
}

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122026f3ac6a49506dbd2da53a242c371d9274f8f4c062ae7e7159094f9d792f840564736f6c63430007060033",
}

// SafeCastABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeCastMetaData.ABI instead.
var SafeCastABI = SafeCastMetaData.ABI

// SafeCastBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeCastMetaData.Bin instead.
var SafeCastBin = SafeCastMetaData.Bin

// DeploySafeCast deploys a new Ethereum contract, binding an instance of SafeCast to it.
func DeploySafeCast(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCast, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeCastBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	SafeCastCaller     // Read-only binding to the contract
	SafeCastTransactor // Write-only binding to the contract
	SafeCastFilterer   // Log filterer for contract events
}

// SafeCastCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCastCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCastTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCastFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCastSession struct {
	Contract     *SafeCast         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCastCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCastCallerSession struct {
	Contract *SafeCastCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeCastTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCastTransactorSession struct {
	Contract     *SafeCastTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeCastRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCastRaw struct {
	Contract *SafeCast // Generic contract binding to access the raw methods on
}

// SafeCastCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCastCallerRaw struct {
	Contract *SafeCastCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCastTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCastTransactorRaw struct {
	Contract *SafeCastTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCast creates a new instance of SafeCast, bound to a specific deployed contract.
func NewSafeCast(address common.Address, backend bind.ContractBackend) (*SafeCast, error) {
	contract, err := bindSafeCast(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// NewSafeCastCaller creates a new read-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastCaller(address common.Address, caller bind.ContractCaller) (*SafeCastCaller, error) {
	contract, err := bindSafeCast(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastCaller{contract: contract}, nil
}

// NewSafeCastTransactor creates a new write-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCastTransactor, error) {
	contract, err := bindSafeCast(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastTransactor{contract: contract}, nil
}

// NewSafeCastFilterer creates a new log filterer instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCastFilterer, error) {
	contract, err := bindSafeCast(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCastFilterer{contract: contract}, nil
}

// bindSafeCast binds a generic wrapper to an already deployed contract.
func bindSafeCast(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeCastABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.SafeCastCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transact(opts, method, params...)
}

// SelfPermitMetaData contains all meta data concerning the SelfPermit contract.
var SelfPermitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SelfPermitABI is the input ABI used to generate the binding from.
// Deprecated: Use SelfPermitMetaData.ABI instead.
var SelfPermitABI = SelfPermitMetaData.ABI

// SelfPermit is an auto generated Go binding around an Ethereum contract.
type SelfPermit struct {
	SelfPermitCaller     // Read-only binding to the contract
	SelfPermitTransactor // Write-only binding to the contract
	SelfPermitFilterer   // Log filterer for contract events
}

// SelfPermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type SelfPermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfPermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SelfPermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfPermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SelfPermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfPermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SelfPermitSession struct {
	Contract     *SelfPermit       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SelfPermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SelfPermitCallerSession struct {
	Contract *SelfPermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SelfPermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SelfPermitTransactorSession struct {
	Contract     *SelfPermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SelfPermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type SelfPermitRaw struct {
	Contract *SelfPermit // Generic contract binding to access the raw methods on
}

// SelfPermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SelfPermitCallerRaw struct {
	Contract *SelfPermitCaller // Generic read-only contract binding to access the raw methods on
}

// SelfPermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SelfPermitTransactorRaw struct {
	Contract *SelfPermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSelfPermit creates a new instance of SelfPermit, bound to a specific deployed contract.
func NewSelfPermit(address common.Address, backend bind.ContractBackend) (*SelfPermit, error) {
	contract, err := bindSelfPermit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SelfPermit{SelfPermitCaller: SelfPermitCaller{contract: contract}, SelfPermitTransactor: SelfPermitTransactor{contract: contract}, SelfPermitFilterer: SelfPermitFilterer{contract: contract}}, nil
}

// NewSelfPermitCaller creates a new read-only instance of SelfPermit, bound to a specific deployed contract.
func NewSelfPermitCaller(address common.Address, caller bind.ContractCaller) (*SelfPermitCaller, error) {
	contract, err := bindSelfPermit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SelfPermitCaller{contract: contract}, nil
}

// NewSelfPermitTransactor creates a new write-only instance of SelfPermit, bound to a specific deployed contract.
func NewSelfPermitTransactor(address common.Address, transactor bind.ContractTransactor) (*SelfPermitTransactor, error) {
	contract, err := bindSelfPermit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SelfPermitTransactor{contract: contract}, nil
}

// NewSelfPermitFilterer creates a new log filterer instance of SelfPermit, bound to a specific deployed contract.
func NewSelfPermitFilterer(address common.Address, filterer bind.ContractFilterer) (*SelfPermitFilterer, error) {
	contract, err := bindSelfPermit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SelfPermitFilterer{contract: contract}, nil
}

// bindSelfPermit binds a generic wrapper to an already deployed contract.
func bindSelfPermit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SelfPermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfPermit *SelfPermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfPermit.Contract.SelfPermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfPermit *SelfPermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfPermit *SelfPermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfPermit *SelfPermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfPermit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfPermit *SelfPermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfPermit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfPermit *SelfPermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfPermit.Contract.contract.Transact(opts, method, params...)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermit(&_SelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermit(&_SelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitAllowed(&_SelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitAllowed(&_SelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitAllowedIfNecessary(&_SelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitAllowedIfNecessary(&_SelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitIfNecessary(&_SelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitIfNecessary(&_SelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// SwapRouterMetaData contains all meta data concerning the SwapRouter contract.
var SwapRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c06040527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6000553480156200003557600080fd5b5060405162003d3238038062003d3283398181016040528101906200005b9190620000ec565b81818173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b81525050505050506200017b565b600081519050620000e68162000161565b92915050565b600080604083850312156200010057600080fd5b60006200011085828601620000d5565b92505060206200012385828601620000d5565b9150509250929050565b60006200013a8262000141565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200016c816200012d565b81146200017857600080fd5b50565b60805160601c60a05160601c613b5b620001d76000398061010952806106a952806107ec528061088752806108c75280610a0a528061207752806120d75280612158525080610fc6528061171452806122875250613b5b6000f3fe6080604052600436106101025760003560e01c8063c04b8d5911610095578063df2ab5bb11610064578063df2ab5bb1461037b578063e0e189a014610397578063f28c0498146103b3578063f3995c67146103e3578063fa461e33146103ff576101ca565b8063c04b8d59146102d4578063c2e3140a14610304578063c45a015514610320578063db3e21981461034b576101ca565b80634aa4a4fc116100d15780634aa4a4fc146102415780639b2c0a371461026c578063a4a78f0c14610288578063ac9650d8146102a4576101ca565b806312210e8a146101cf578063414bf389146101d95780634659a4941461020957806349404b7c14610225576101ca565b366101ca577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260098152602001807f4e6f74205745544839000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b005b600080fd5b6101d7610428565b005b6101f360048036038101906101ee9190613201565b61043e565b60405161020091906137d7565b60405180910390f35b610223600480360381019061021e9190613009565b6105c8565b005b61023f600480360381019061023a9190613329565b6106a5565b005b34801561024d57600080fd5b50610256610885565b60405161026391906136bc565b60405180910390f35b61028660048036038101906102819190613365565b6108a9565b005b6102a2600480360381019061029d9190613009565b610ade565b005b6102be60048036038101906102b99190613092565b610bdc565b6040516102cb9190613731565b60405180910390f35b6102ee60048036038101906102e991906131c0565b610d60565b6040516102fb91906137d7565b60405180910390f35b61031e60048036038101906103199190613009565b610ee6565b005b34801561032c57600080fd5b50610335610fc4565b60405161034291906136bc565b60405180910390f35b6103656004803603810190610360919061326c565b610fe8565b60405161037291906137d7565b60405180910390f35b61039560048036038101906103909190612f43565b611199565b005b6103b160048036038101906103ac9190612f92565b6112d0565b005b6103cd60048036038101906103c8919061322b565b61145d565b6040516103da91906137d7565b60405180910390f35b6103fd60048036038101906103f89190613009565b6115fa565b005b34801561040b57600080fd5b5061042660048036038101906104219190613113565b6116cc565b005b600047111561043c5761043b334761182e565b5b565b600081608001358061044e6119aa565b11156104c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f5472616e73616374696f6e20746f6f206f6c640000000000000000000000000081525060200191505060405180910390fd5b6105798360a001358460600160208101906104dd9190612f1a565b8560e00160208101906104f091906132d7565b604051806040016040528088600001602081019061050e9190612f1a565b8960400160208101906105219190613300565b8a60200160208101906105349190612f1a565b60405160200161054693929190613666565b60405160208183030381529060405281526020013373ffffffffffffffffffffffffffffffffffffffff168152506119b2565b91508260c001358210156105c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b990613795565b60405180910390fd5b50919050565b8573ffffffffffffffffffffffffffffffffffffffff16638fcbaf0c3330888860018989896040518963ffffffff1660e01b8152600401808973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200187815260200186815260200185151581526020018460ff16815260200183815260200182815260200198505050505050505050600060405180830381600087803b15801561068557600080fd5b505af1158015610699573d6000803e3d6000fd5b50505050505050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561072e57600080fd5b505afa158015610742573d6000803e3d6000fd5b505050506040513d602081101561075857600080fd5b81019080805190602001909291905050509050828110156107e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f496e73756666696369656e74205745544839000000000000000000000000000081525060200191505060405180910390fd5b6000811115610880577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561085d57600080fd5b505af1158015610871573d6000803e3d6000fd5b5050505061087f828261182e565b5b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000821180156108ba575060648211155b6108c357600080fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561094c57600080fd5b505afa158015610960573d6000803e3d6000fd5b505050506040513d602081101561097657600080fd5b81019080805190602001909291905050509050848110156109ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f496e73756666696369656e74205745544839000000000000000000000000000081525060200191505060405180910390fd5b6000811115610ad7577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015610a7b57600080fd5b505af1158015610a8f573d6000803e3d6000fd5b505050506000612710610aab8584611b7190919063ffffffff16565b81610ab257fe5b0490506000811115610ac957610ac8838261182e565b5b610ad58582840361182e565b505b5050505050565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b158015610b8457600080fd5b505afa158015610b98573d6000803e3d6000fd5b505050506040513d6020811015610bae57600080fd5b81019080805190602001909291905050501015610bd457610bd38686868686866105c8565b5b505050505050565b60608282905067ffffffffffffffff81118015610bf857600080fd5b50604051908082528060200260200182016040528015610c2c57816020015b6060815260200190600190039081610c175790505b50905060005b83839050811015610d59576000803073ffffffffffffffffffffffffffffffffffffffff16868685818110610c6357fe5b9050602002810190610c7591906137f2565b604051610c839291906136a3565b600060405180830381855af49150503d8060008114610cbe576040519150601f19603f3d011682016040523d82523d6000602084013e610cc3565b606091505b509150915081610d3257604481511015610cdc57600080fd5b60048101905080806020019051810190610cf6919061317f565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d299190613753565b60405180910390fd5b80848481518110610d3f57fe5b602002602001018190525050508080600101915050610c32565b5092915050565b6000816040015180610d706119aa565b1115610de4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f5472616e73616374696f6e20746f6f206f6c640000000000000000000000000081525060200191505060405180910390fd5b60003390505b600115610e98576000610e008560000151611b9d565b9050610e59856060015182610e19578660200151610e1b565b305b60006040518060400160405280610e358b60000151611bb8565b81526020018773ffffffffffffffffffffffffffffffffffffffff168152506119b2565b8560600181815250508015610e8557309150610e788560000151611bde565b8560000181905250610e92565b8460600151935050610e98565b50610dea565b8360800151831015610edf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed690613795565b60405180910390fd5b5050919050565b848673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b158015610f6c57600080fd5b505afa158015610f80573d6000803e3d6000fd5b505050506040513d6020811015610f9657600080fd5b81019080805190602001909291905050501015610fbc57610fbb8686868686866115fa565b5b505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000816080013580610ff86119aa565b111561106c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f5472616e73616374696f6e20746f6f206f6c640000000000000000000000000081525060200191505060405180910390fd5b6111238360a001358460600160208101906110879190612f1a565b8560e001602081019061109a91906132d7565b60405180604001604052808860200160208101906110b89190612f1a565b8960400160208101906110cb9190613300565b8a60000160208101906110de9190612f1a565b6040516020016110f093929190613666565b60405160208183030381529060405281526020013373ffffffffffffffffffffffffffffffffffffffff16815250611c07565b91508260c0013582111561116c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161116390613775565b60405180910390fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60008190555050919050565b60008373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561120257600080fd5b505afa158015611216573d6000803e3d6000fd5b505050506040513d602081101561122c57600080fd5b81019080805190602001909291905050509050828110156112b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f496e73756666696369656e7420746f6b656e000000000000000000000000000081525060200191505060405180910390fd5b60008111156112ca576112c9848383611e03565b5b50505050565b6000821180156112e1575060648211155b6112ea57600080fd5b60008573ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561135357600080fd5b505afa158015611367573d6000803e3d6000fd5b505050506040513d602081101561137d57600080fd5b8101908080519060200190929190505050905084811015611406576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f496e73756666696369656e7420746f6b656e000000000000000000000000000081525060200191505060405180910390fd5b60008111156114555760006127106114278584611b7190919063ffffffff16565b8161142e57fe5b049050600081111561144657611445878483611e03565b5b6114538786838503611e03565b505b505050505050565b600081604001358061146d6119aa565b11156114e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f5472616e73616374696f6e20746f6f206f6c640000000000000000000000000081525060200191505060405180910390fd5b61158083606001358460200160208101906114fc9190612f1a565b6000604051806040016040528088806000019061151991906137f2565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020013373ffffffffffffffffffffffffffffffffffffffff16815250611c07565b50600054915082608001358211156115cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115c490613775565b60405180910390fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60008190555050919050565b8573ffffffffffffffffffffffffffffffffffffffff1663d505accf333088888888886040518863ffffffff1660e01b8152600401808873ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018581526020018460ff168152602001838152602001828152602001975050505050505050600060405180830381600087803b1580156116ac57600080fd5b505af11580156116c0573d6000803e3d6000fd5b50505050505050505050565b60008413806116db5750600083135b6116e457600080fd5b600082828101906116f59190613296565b905060008060006117098460000151612004565b92509250925061173b7f0000000000000000000000000000000000000000000000000000000000000000848484612055565b5060008060008a1361177c578473ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610896117ad565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16108a5b9150915081156117cc576117c78587602001513384612075565b611822565b6117d98660000151611b9d565b15611806576117eb8660000151611bde565b86600001819052506118008133600089611c07565b50611821565b806000819055508394506118208587602001513384612075565b5b5b50505050505050505050565b60008273ffffffffffffffffffffffffffffffffffffffff1682600067ffffffffffffffff8111801561186057600080fd5b506040519080825280601f01601f1916602001820160405280156118935781602001600182028036833780820191505090505b506040518082805190602001908083835b602083106118c757805182526020820191506020810190506020830392506118a4565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611929576040519150601f19603f3d011682016040523d82523d6000602084013e61192e565b606091505b50509050806119a5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f535445000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b505050565b600042905090565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156119ec573093505b60008060006119fe8560000151612004565b92509250925060008273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16109050600080611a45868686612280565b73ffffffffffffffffffffffffffffffffffffffff1663128acb088b85611a6b8f6122bf565b60008e73ffffffffffffffffffffffffffffffffffffffff1614611a8f578d611abc565b87611ab157600173fffd8963efd1fc6a506488495d951d5263988d2603611abb565b60016401000276a3015b5b8d604051602001611acd91906137b5565b6040516020818303038152906040526040518663ffffffff1660e01b8152600401611afc9594939291906136d7565b6040805180830381600087803b158015611b1557600080fd5b505af1158015611b29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b4d91906130d7565b9150915082611b5c5781611b5e565b805b6000039650505050505050949350505050565b600080831480611b8e5750818383850292508281611b8b57fe5b04145b611b9757600080fd5b92915050565b60006003601401601460036014010101825110159050919050565b6060611bd760006014600360140101846122f59092919063ffffffff16565b9050919050565b6060611c0060036014016003601401845103846122f59092919063ffffffff16565b9050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415611c41573093505b6000806000611c538560000151612004565b92509250925060008373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16109050600080611c9a858786612280565b73ffffffffffffffffffffffffffffffffffffffff1663128acb088b85611cc08f6122bf565b60000360008e73ffffffffffffffffffffffffffffffffffffffff1614611ce7578d611d14565b87611d0957600173fffd8963efd1fc6a506488495d951d5263988d2603611d13565b60016401000276a3015b5b8d604051602001611d2591906137b5565b6040516020818303038152906040526040518663ffffffff1660e01b8152600401611d549594939291906136d7565b6040805180830381600087803b158015611d6d57600080fd5b505af1158015611d81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611da591906130d7565b91509150600083611dba578183600003611dc0565b82826000035b809250819950505060008a73ffffffffffffffffffffffffffffffffffffffff161415611df4578b8114611df357600080fd5b5b50505050505050949350505050565b6000808473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60e01b8585604051602401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518082805190602001908083835b60208310611ee45780518252602082019150602081019050602083039250611ec1565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611f46576040519150601f19603f3d011682016040523d82523d6000602084013e611f4b565b606091505b5091509150818015611f8b5750600081511480611f8a5750808060200190516020811015611f7857600080fd5b81019080805190602001909291905050505b5b611ffd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260028152602001807f535400000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b5050505050565b600080600061201d6000856124df90919063ffffffff16565b92506120336014856125f890919063ffffffff16565b905061204c6003601401856124df90919063ffffffff16565b91509193909250565b600061206b85612066868686612702565b61279e565b9050949350505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480156120d05750804710155b15612228577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b15801561213d57600080fd5b505af1158015612151573d6000803e3d6000fd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156121e757600080fd5b505af11580156121fb573d6000803e3d6000fd5b505050506040513d602081101561221157600080fd5b81019080805190602001909291905050505061227a565b3073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561226c57612267848383611e03565b612279565b612278848484846127ea565b5b5b50505050565b60006122b67f00000000000000000000000000000000000000000000000000000000000000006122b1868686612702565b612a0a565b90509392505050565b60007f800000000000000000000000000000000000000000000000000000000000000082106122ed57600080fd5b819050919050565b606081601f83011015612370576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f736c6963655f6f766572666c6f7700000000000000000000000000000000000081525060200191505060405180910390fd5b8282840110156123e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f736c6963655f6f766572666c6f7700000000000000000000000000000000000081525060200191505060405180910390fd5b81830184511015612461576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f736c6963655f6f75744f66426f756e647300000000000000000000000000000081525060200191505060405180910390fd5b606082156000811461248257604051915060008252602082016040526124d3565b6040519150601f8416801560200281840101858101878315602002848b0101015b818310156124c057805183526020830192506020810190506124a3565b50868552601f19601f8301166040525050505b50809150509392505050565b60008160148301101561255a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f746f416464726573735f6f766572666c6f77000000000000000000000000000081525060200191505060405180910390fd5b60148201835110156125d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f416464726573735f6f75744f66426f756e6473000000000000000000000081525060200191505060405180910390fd5b60006c01000000000000000000000000836020860101510490508091505092915050565b600081600383011015612673576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f746f55696e7432345f6f766572666c6f7700000000000000000000000000000081525060200191505060405180910390fd5b60038201835110156126ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f746f55696e7432345f6f75744f66426f756e647300000000000000000000000081525060200191505060405180910390fd5b60008260038501015190508091505092915050565b61270a612b65565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16111561274957828480945081955050505b60405180606001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018362ffffff1681525090509392505050565b60006127aa8383612a0a565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146127e457600080fd5b92915050565b6000808573ffffffffffffffffffffffffffffffffffffffff166323b872dd60e01b868686604051602401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518082805190602001908083835b602083106128e957805182526020820191506020810190506020830392506128c6565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461294b576040519150601f19603f3d011682016040523d82523d6000602084013e612950565b606091505b5091509150818015612990575060008151148061298f575080806020019051602081101561297d57600080fd5b81019080805190602001909291905050505b5b612a02576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f535446000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b505050505050565b6000816020015173ffffffffffffffffffffffffffffffffffffffff16826000015173ffffffffffffffffffffffffffffffffffffffff1610612a4c57600080fd5b82826000015183602001518460400151604051602001808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018262ffffff1681526020019350505050604051602081830303815290604052805190602001207fe34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b5460001b60405160200180807fff000000000000000000000000000000000000000000000000000000000000008152506001018473ffffffffffffffffffffffffffffffffffffffff1660601b815260140183815260200182815260200193505050506040516020818303038152906040528051906020012060001c905092915050565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600062ffffff1681525090565b6000612bca612bc58461387a565b613849565b905082815260208101848484011115612be257600080fd5b612bed8482856139df565b509392505050565b6000612c08612c03846138aa565b613849565b905082815260208101848484011115612c2057600080fd5b612c2b8482856139ee565b509392505050565b600081359050612c4281613a84565b92915050565b60008083601f840112612c5a57600080fd5b8235905067ffffffffffffffff811115612c7357600080fd5b602083019150836020820283011115612c8b57600080fd5b9250929050565b600081359050612ca181613a9b565b92915050565b60008083601f840112612cb957600080fd5b8235905067ffffffffffffffff811115612cd257600080fd5b602083019150836001820283011115612cea57600080fd5b9250929050565b600082601f830112612d0257600080fd5b8135612d12848260208601612bb7565b91505092915050565b600081359050612d2a81613ab2565b92915050565b600081519050612d3f81613ab2565b92915050565b600082601f830112612d5657600080fd5b8151612d66848260208601612bf5565b91505092915050565b600060a08284031215612d8157600080fd5b612d8b60a0613849565b9050600082013567ffffffffffffffff811115612da757600080fd5b612db384828501612cf1565b6000830152506020612dc784828501612c33565b6020830152506040612ddb84828501612ef0565b6040830152506060612def84828501612ef0565b6060830152506080612e0384828501612ef0565b60808301525092915050565b60006101008284031215612e2257600080fd5b81905092915050565b600060a08284031215612e3d57600080fd5b81905092915050565b60006101008284031215612e5957600080fd5b81905092915050565b600060408284031215612e7457600080fd5b612e7e6040613849565b9050600082013567ffffffffffffffff811115612e9a57600080fd5b612ea684828501612cf1565b6000830152506020612eba84828501612c33565b60208301525092915050565b600081359050612ed581613ac9565b92915050565b600081359050612eea81613ae0565b92915050565b600081359050612eff81613af7565b92915050565b600081359050612f1481613b0e565b92915050565b600060208284031215612f2c57600080fd5b6000612f3a84828501612c33565b91505092915050565b600080600060608486031215612f5857600080fd5b6000612f6686828701612c33565b9350506020612f7786828701612ef0565b9250506040612f8886828701612c33565b9150509250925092565b600080600080600060a08688031215612faa57600080fd5b6000612fb888828901612c33565b9550506020612fc988828901612ef0565b9450506040612fda88828901612c33565b9350506060612feb88828901612ef0565b9250506080612ffc88828901612c33565b9150509295509295909350565b60008060008060008060c0878903121561302257600080fd5b600061303089828a01612c33565b965050602061304189828a01612ef0565b955050604061305289828a01612ef0565b945050606061306389828a01612f05565b935050608061307489828a01612c92565b92505060a061308589828a01612c92565b9150509295509295509295565b600080602083850312156130a557600080fd5b600083013567ffffffffffffffff8111156130bf57600080fd5b6130cb85828601612c48565b92509250509250929050565b600080604083850312156130ea57600080fd5b60006130f885828601612d30565b925050602061310985828601612d30565b9150509250929050565b6000806000806060858703121561312957600080fd5b600061313787828801612d1b565b945050602061314887828801612d1b565b935050604085013567ffffffffffffffff81111561316557600080fd5b61317187828801612ca7565b925092505092959194509250565b60006020828403121561319157600080fd5b600082015167ffffffffffffffff8111156131ab57600080fd5b6131b784828501612d45565b91505092915050565b6000602082840312156131d257600080fd5b600082013567ffffffffffffffff8111156131ec57600080fd5b6131f884828501612d6f565b91505092915050565b6000610100828403121561321457600080fd5b600061322284828501612e0f565b91505092915050565b60006020828403121561323d57600080fd5b600082013567ffffffffffffffff81111561325757600080fd5b61326384828501612e2b565b91505092915050565b6000610100828403121561327f57600080fd5b600061328d84828501612e46565b91505092915050565b6000602082840312156132a857600080fd5b600082013567ffffffffffffffff8111156132c257600080fd5b6132ce84828501612e62565b91505092915050565b6000602082840312156132e957600080fd5b60006132f784828501612ec6565b91505092915050565b60006020828403121561331257600080fd5b600061332084828501612edb565b91505092915050565b6000806040838503121561333c57600080fd5b600061334a85828601612ef0565b925050602061335b85828601612c33565b9150509250929050565b6000806000806080858703121561337b57600080fd5b600061338987828801612ef0565b945050602061339a87828801612c33565b93505060406133ab87828801612ef0565b92505060606133bc87828801612c33565b91505092959194509250565b60006133d483836134ba565b905092915050565b6133e581613967565b82525050565b6133f481613967565b82525050565b61340b61340682613967565b613a21565b82525050565b600061341c826138ea565b6134268185613918565b935083602082028501613438856138da565b8060005b85811015613474578484038952815161345585826133c8565b94506134608361390b565b925060208a0199505060018101905061343c565b50829750879550505050505092915050565b61348f81613979565b82525050565b60006134a1838561394b565b93506134ae8385846139df565b82840190509392505050565b60006134c5826138f5565b6134cf8185613929565b93506134df8185602086016139ee565b6134e881613a59565b840191505092915050565b60006134fe826138f5565b613508818561393a565b93506135188185602086016139ee565b61352181613a59565b840191505092915050565b6135358161398f565b82525050565b600061354682613900565b6135508185613956565b93506135608185602086016139ee565b61356981613a59565b840191505092915050565b6000613581601283613956565b91507f546f6f206d7563682072657175657374656400000000000000000000000000006000830152602082019050919050565b60006135c1601383613956565b91507f546f6f206c6974746c65207265636569766564000000000000000000000000006000830152602082019050919050565b6000604083016000830151848203600086015261361182826134ba565b915050602083015161362660208601826133dc565b508091505092915050565b61363a81613999565b82525050565b61365161364c826139b9565b613a45565b82525050565b613660816139c8565b82525050565b600061367282866133fa565b6014820191506136828285613640565b60038201915061369282846133fa565b601482019150819050949350505050565b60006136b0828486613495565b91508190509392505050565b60006020820190506136d160008301846133eb565b92915050565b600060a0820190506136ec60008301886133eb565b6136f96020830187613486565b613706604083018661352c565b6137136060830185613631565b818103608083015261372581846134f3565b90509695505050505050565b6000602082019050818103600083015261374b8184613411565b905092915050565b6000602082019050818103600083015261376d818461353b565b905092915050565b6000602082019050818103600083015261378e81613574565b9050919050565b600060208201905081810360008301526137ae816135b4565b9050919050565b600060208201905081810360008301526137cf81846135f4565b905092915050565b60006020820190506137ec6000830184613657565b92915050565b6000808335600160200384360303811261380b57600080fd5b80840192508235915067ffffffffffffffff82111561382957600080fd5b60208301925060018202360383131561384157600080fd5b509250929050565b6000604051905081810181811067ffffffffffffffff821117156138705761386f613a57565b5b8060405250919050565b600067ffffffffffffffff82111561389557613894613a57565b5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff8211156138c5576138c4613a57565b5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061397282613999565b9050919050565b60008115159050919050565b6000819050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015613a0c5780820151818401526020810190506139f1565b83811115613a1b576000848401525b50505050565b6000613a2c82613a33565b9050919050565b6000613a3e82613a77565b9050919050565b6000613a5082613a6a565b9050919050565bfe5b6000601f19601f8301169050919050565b60008160e81b9050919050565b60008160601b9050919050565b613a8d81613967565b8114613a9857600080fd5b50565b613aa481613985565b8114613aaf57600080fd5b50565b613abb8161398f565b8114613ac657600080fd5b50565b613ad281613999565b8114613add57600080fd5b50565b613ae9816139b9565b8114613af457600080fd5b50565b613b00816139c8565b8114613b0b57600080fd5b50565b613b17816139d2565b8114613b2257600080fd5b5056fea2646970667358221220d8bd7f01a4876a4a2896d8c7c2080bfd5237bd5ee6ec7376fd02fc6bc18e9e5864736f6c63430007060033",
}

// SwapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapRouterMetaData.ABI instead.
var SwapRouterABI = SwapRouterMetaData.ABI

// SwapRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapRouterMetaData.Bin instead.
var SwapRouterBin = SwapRouterMetaData.Bin

// DeploySwapRouter deploys a new Ethereum contract, binding an instance of SwapRouter to it.
func DeploySwapRouter(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _WETH9 common.Address) (common.Address, *types.Transaction, *SwapRouter, error) {
	parsed, err := SwapRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapRouterBin), backend, _factory, _WETH9)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwapRouter{SwapRouterCaller: SwapRouterCaller{contract: contract}, SwapRouterTransactor: SwapRouterTransactor{contract: contract}, SwapRouterFilterer: SwapRouterFilterer{contract: contract}}, nil
}

// SwapRouter is an auto generated Go binding around an Ethereum contract.
type SwapRouter struct {
	SwapRouterCaller     // Read-only binding to the contract
	SwapRouterTransactor // Write-only binding to the contract
	SwapRouterFilterer   // Log filterer for contract events
}

// SwapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapRouterSession struct {
	Contract     *SwapRouter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapRouterCallerSession struct {
	Contract *SwapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SwapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapRouterTransactorSession struct {
	Contract     *SwapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SwapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRouterRaw struct {
	Contract *SwapRouter // Generic contract binding to access the raw methods on
}

// SwapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapRouterCallerRaw struct {
	Contract *SwapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// SwapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapRouterTransactorRaw struct {
	Contract *SwapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapRouter creates a new instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouter(address common.Address, backend bind.ContractBackend) (*SwapRouter, error) {
	contract, err := bindSwapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapRouter{SwapRouterCaller: SwapRouterCaller{contract: contract}, SwapRouterTransactor: SwapRouterTransactor{contract: contract}, SwapRouterFilterer: SwapRouterFilterer{contract: contract}}, nil
}

// NewSwapRouterCaller creates a new read-only instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouterCaller(address common.Address, caller bind.ContractCaller) (*SwapRouterCaller, error) {
	contract, err := bindSwapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapRouterCaller{contract: contract}, nil
}

// NewSwapRouterTransactor creates a new write-only instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapRouterTransactor, error) {
	contract, err := bindSwapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapRouterTransactor{contract: contract}, nil
}

// NewSwapRouterFilterer creates a new log filterer instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapRouterFilterer, error) {
	contract, err := bindSwapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapRouterFilterer{contract: contract}, nil
}

// bindSwapRouter binds a generic wrapper to an already deployed contract.
func bindSwapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapRouter *SwapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapRouter.Contract.SwapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapRouter *SwapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRouter.Contract.SwapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapRouter *SwapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapRouter.Contract.SwapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapRouter *SwapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapRouter *SwapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapRouter *SwapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapRouter.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_SwapRouter *SwapRouterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapRouter.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_SwapRouter *SwapRouterSession) WETH9() (common.Address, error) {
	return _SwapRouter.Contract.WETH9(&_SwapRouter.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_SwapRouter *SwapRouterCallerSession) WETH9() (common.Address, error) {
	return _SwapRouter.Contract.WETH9(&_SwapRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapRouter *SwapRouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapRouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapRouter *SwapRouterSession) Factory() (common.Address, error) {
	return _SwapRouter.Contract.Factory(&_SwapRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapRouter *SwapRouterCallerSession) Factory() (common.Address, error) {
	return _SwapRouter.Contract.Factory(&_SwapRouter.CallOpts)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterTransactor) ExactInput(opts *bind.TransactOpts, params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactInput(&_SwapRouter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterTransactorSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactInput(&_SwapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactInputSingle(&_SwapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterTransactorSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactInputSingle(&_SwapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterTransactor) ExactOutput(opts *bind.TransactOpts, params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactOutput(&_SwapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterTransactorSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactOutput(&_SwapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactOutputSingle(&_SwapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterTransactorSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactOutputSingle(&_SwapRouter.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SwapRouter *SwapRouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SwapRouter *SwapRouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.Multicall(&_SwapRouter.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SwapRouter *SwapRouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.Multicall(&_SwapRouter.TransactOpts, data)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SwapRouter *SwapRouterTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SwapRouter *SwapRouterSession) RefundETH() (*types.Transaction, error) {
	return _SwapRouter.Contract.RefundETH(&_SwapRouter.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SwapRouter *SwapRouterTransactorSession) RefundETH() (*types.Transaction, error) {
	return _SwapRouter.Contract.RefundETH(&_SwapRouter.TransactOpts)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermit(&_SwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermit(&_SwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitAllowed(&_SwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitAllowed(&_SwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitAllowedIfNecessary(&_SwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitAllowedIfNecessary(&_SwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitIfNecessary(&_SwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitIfNecessary(&_SwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.SweepToken(&_SwapRouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.SweepToken(&_SwapRouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterTransactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.SweepTokenWithFee(&_SwapRouter.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.SweepTokenWithFee(&_SwapRouter.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SwapRouter *SwapRouterTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SwapRouter *SwapRouterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.UniswapV3SwapCallback(&_SwapRouter.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SwapRouter *SwapRouterTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.UniswapV3SwapCallback(&_SwapRouter.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.UnwrapWETH9(&_SwapRouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.UnwrapWETH9(&_SwapRouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterTransactor) UnwrapWETH9WithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.UnwrapWETH9WithFee(&_SwapRouter.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.UnwrapWETH9WithFee(&_SwapRouter.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapRouter *SwapRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapRouter *SwapRouterSession) Receive() (*types.Transaction, error) {
	return _SwapRouter.Contract.Receive(&_SwapRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapRouter *SwapRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _SwapRouter.Contract.Receive(&_SwapRouter.TransactOpts)
}

// TickMathMetaData contains all meta data concerning the TickMath contract.
var TickMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203b28c0860c299f54ae704c77525da202fb2287719dddb7b1b8253bc16fea7d3064736f6c63430007060033",
}

// TickMathABI is the input ABI used to generate the binding from.
// Deprecated: Use TickMathMetaData.ABI instead.
var TickMathABI = TickMathMetaData.ABI

// TickMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TickMathMetaData.Bin instead.
var TickMathBin = TickMathMetaData.Bin

// DeployTickMath deploys a new Ethereum contract, binding an instance of TickMath to it.
func DeployTickMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TickMath, error) {
	parsed, err := TickMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TickMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TickMath{TickMathCaller: TickMathCaller{contract: contract}, TickMathTransactor: TickMathTransactor{contract: contract}, TickMathFilterer: TickMathFilterer{contract: contract}}, nil
}

// TickMath is an auto generated Go binding around an Ethereum contract.
type TickMath struct {
	TickMathCaller     // Read-only binding to the contract
	TickMathTransactor // Write-only binding to the contract
	TickMathFilterer   // Log filterer for contract events
}

// TickMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type TickMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TickMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TickMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TickMathSession struct {
	Contract     *TickMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TickMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TickMathCallerSession struct {
	Contract *TickMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TickMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TickMathTransactorSession struct {
	Contract     *TickMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TickMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type TickMathRaw struct {
	Contract *TickMath // Generic contract binding to access the raw methods on
}

// TickMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TickMathCallerRaw struct {
	Contract *TickMathCaller // Generic read-only contract binding to access the raw methods on
}

// TickMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TickMathTransactorRaw struct {
	Contract *TickMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTickMath creates a new instance of TickMath, bound to a specific deployed contract.
func NewTickMath(address common.Address, backend bind.ContractBackend) (*TickMath, error) {
	contract, err := bindTickMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TickMath{TickMathCaller: TickMathCaller{contract: contract}, TickMathTransactor: TickMathTransactor{contract: contract}, TickMathFilterer: TickMathFilterer{contract: contract}}, nil
}

// NewTickMathCaller creates a new read-only instance of TickMath, bound to a specific deployed contract.
func NewTickMathCaller(address common.Address, caller bind.ContractCaller) (*TickMathCaller, error) {
	contract, err := bindTickMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TickMathCaller{contract: contract}, nil
}

// NewTickMathTransactor creates a new write-only instance of TickMath, bound to a specific deployed contract.
func NewTickMathTransactor(address common.Address, transactor bind.ContractTransactor) (*TickMathTransactor, error) {
	contract, err := bindTickMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TickMathTransactor{contract: contract}, nil
}

// NewTickMathFilterer creates a new log filterer instance of TickMath, bound to a specific deployed contract.
func NewTickMathFilterer(address common.Address, filterer bind.ContractFilterer) (*TickMathFilterer, error) {
	contract, err := bindTickMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TickMathFilterer{contract: contract}, nil
}

// bindTickMath binds a generic wrapper to an already deployed contract.
func bindTickMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TickMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TickMath *TickMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TickMath.Contract.TickMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TickMath *TickMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TickMath.Contract.TickMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TickMath *TickMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TickMath.Contract.TickMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TickMath *TickMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TickMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TickMath *TickMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TickMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TickMath *TickMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TickMath.Contract.contract.Transact(opts, method, params...)
}

// TransferHelperMetaData contains all meta data concerning the TransferHelper contract.
var TransferHelperMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c609659cd9defab63faf55a38e3750ef828cd756777e6b000d994cdb383d87dc64736f6c63430007060033",
}

// TransferHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferHelperMetaData.ABI instead.
var TransferHelperABI = TransferHelperMetaData.ABI

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferHelperMetaData.Bin instead.
var TransferHelperBin = TransferHelperMetaData.Bin

// DeployTransferHelper deploys a new Ethereum contract, binding an instance of TransferHelper to it.
func DeployTransferHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferHelper, error) {
	parsed, err := TransferHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// TransferHelper is an auto generated Go binding around an Ethereum contract.
type TransferHelper struct {
	TransferHelperCaller     // Read-only binding to the contract
	TransferHelperTransactor // Write-only binding to the contract
	TransferHelperFilterer   // Log filterer for contract events
}

// TransferHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperSession struct {
	Contract     *TransferHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperCallerSession struct {
	Contract *TransferHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperTransactorSession struct {
	Contract     *TransferHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperRaw struct {
	Contract *TransferHelper // Generic contract binding to access the raw methods on
}

// TransferHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperCallerRaw struct {
	Contract *TransferHelperCaller // Generic read-only contract binding to access the raw methods on
}

// TransferHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperTransactorRaw struct {
	Contract *TransferHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferHelper creates a new instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelper(address common.Address, backend bind.ContractBackend) (*TransferHelper, error) {
	contract, err := bindTransferHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// NewTransferHelperCaller creates a new read-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperCaller, error) {
	contract, err := bindTransferHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperCaller{contract: contract}, nil
}

// NewTransferHelperTransactor creates a new write-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperTransactor, error) {
	contract, err := bindTransferHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperTransactor{contract: contract}, nil
}

// NewTransferHelperFilterer creates a new log filterer instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperFilterer, error) {
	contract, err := bindTransferHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferHelperFilterer{contract: contract}, nil
}

// bindTransferHelper binds a generic wrapper to an already deployed contract.
func bindTransferHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.TransferHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transact(opts, method, params...)
}
