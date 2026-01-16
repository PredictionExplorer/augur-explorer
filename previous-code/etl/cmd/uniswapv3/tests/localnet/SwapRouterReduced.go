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
