
package main

import (
	"strings"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)
// IUniswapV3MintCallbackMetaData contains all meta data concerning the IUniswapV3MintCallback contract.
var IUniswapV3MintCallbackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IUniswapV3MintCallbackABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3MintCallbackMetaData.ABI instead.
var IUniswapV3MintCallbackABI = IUniswapV3MintCallbackMetaData.ABI

// IUniswapV3MintCallback is an auto generated Go binding around an Ethereum contract.
type IUniswapV3MintCallback struct {
	IUniswapV3MintCallbackCaller     // Read-only binding to the contract
	IUniswapV3MintCallbackTransactor // Write-only binding to the contract
	IUniswapV3MintCallbackFilterer   // Log filterer for contract events
}

// IUniswapV3MintCallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3MintCallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3MintCallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3MintCallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3MintCallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3MintCallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3MintCallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3MintCallbackSession struct {
	Contract     *IUniswapV3MintCallback // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IUniswapV3MintCallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3MintCallbackCallerSession struct {
	Contract *IUniswapV3MintCallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IUniswapV3MintCallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3MintCallbackTransactorSession struct {
	Contract     *IUniswapV3MintCallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IUniswapV3MintCallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3MintCallbackRaw struct {
	Contract *IUniswapV3MintCallback // Generic contract binding to access the raw methods on
}

// IUniswapV3MintCallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3MintCallbackCallerRaw struct {
	Contract *IUniswapV3MintCallbackCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3MintCallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3MintCallbackTransactorRaw struct {
	Contract *IUniswapV3MintCallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3MintCallback creates a new instance of IUniswapV3MintCallback, bound to a specific deployed contract.
func NewIUniswapV3MintCallback(address common.Address, backend bind.ContractBackend) (*IUniswapV3MintCallback, error) {
	contract, err := bindIUniswapV3MintCallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3MintCallback{IUniswapV3MintCallbackCaller: IUniswapV3MintCallbackCaller{contract: contract}, IUniswapV3MintCallbackTransactor: IUniswapV3MintCallbackTransactor{contract: contract}, IUniswapV3MintCallbackFilterer: IUniswapV3MintCallbackFilterer{contract: contract}}, nil
}

// NewIUniswapV3MintCallbackCaller creates a new read-only instance of IUniswapV3MintCallback, bound to a specific deployed contract.
func NewIUniswapV3MintCallbackCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3MintCallbackCaller, error) {
	contract, err := bindIUniswapV3MintCallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3MintCallbackCaller{contract: contract}, nil
}

// NewIUniswapV3MintCallbackTransactor creates a new write-only instance of IUniswapV3MintCallback, bound to a specific deployed contract.
func NewIUniswapV3MintCallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3MintCallbackTransactor, error) {
	contract, err := bindIUniswapV3MintCallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3MintCallbackTransactor{contract: contract}, nil
}

// NewIUniswapV3MintCallbackFilterer creates a new log filterer instance of IUniswapV3MintCallback, bound to a specific deployed contract.
func NewIUniswapV3MintCallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3MintCallbackFilterer, error) {
	contract, err := bindIUniswapV3MintCallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3MintCallbackFilterer{contract: contract}, nil
}

// bindIUniswapV3MintCallback binds a generic wrapper to an already deployed contract.
func bindIUniswapV3MintCallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3MintCallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3MintCallback *IUniswapV3MintCallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3MintCallback.Contract.IUniswapV3MintCallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3MintCallback *IUniswapV3MintCallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3MintCallback.Contract.IUniswapV3MintCallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3MintCallback *IUniswapV3MintCallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3MintCallback.Contract.IUniswapV3MintCallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3MintCallback *IUniswapV3MintCallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3MintCallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3MintCallback *IUniswapV3MintCallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3MintCallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3MintCallback *IUniswapV3MintCallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3MintCallback.Contract.contract.Transact(opts, method, params...)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_IUniswapV3MintCallback *IUniswapV3MintCallbackTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3MintCallback.contract.Transact(opts, "uniswapV3MintCallback", amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_IUniswapV3MintCallback *IUniswapV3MintCallbackSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3MintCallback.Contract.UniswapV3MintCallback(&_IUniswapV3MintCallback.TransactOpts, amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_IUniswapV3MintCallback *IUniswapV3MintCallbackTransactorSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3MintCallback.Contract.UniswapV3MintCallback(&_IUniswapV3MintCallback.TransactOpts, amount0Owed, amount1Owed, data)
}

