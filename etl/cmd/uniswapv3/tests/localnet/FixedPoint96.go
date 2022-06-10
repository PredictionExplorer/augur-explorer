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

// FixedPoint96MetaData contains all meta data concerning the FixedPoint96 contract.
var FixedPoint96MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122002788048011064e1124b787e1525c4cb36e3776c1fa6cde03b9425a0d5f12a9264736f6c63430007060033",
}

// FixedPoint96ABI is the input ABI used to generate the binding from.
// Deprecated: Use FixedPoint96MetaData.ABI instead.
var FixedPoint96ABI = FixedPoint96MetaData.ABI

// FixedPoint96Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FixedPoint96MetaData.Bin instead.
var FixedPoint96Bin = FixedPoint96MetaData.Bin

// DeployFixedPoint96 deploys a new Ethereum contract, binding an instance of FixedPoint96 to it.
func DeployFixedPoint96(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FixedPoint96, error) {
	parsed, err := FixedPoint96MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FixedPoint96Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FixedPoint96{FixedPoint96Caller: FixedPoint96Caller{contract: contract}, FixedPoint96Transactor: FixedPoint96Transactor{contract: contract}, FixedPoint96Filterer: FixedPoint96Filterer{contract: contract}}, nil
}

// FixedPoint96 is an auto generated Go binding around an Ethereum contract.
type FixedPoint96 struct {
	FixedPoint96Caller     // Read-only binding to the contract
	FixedPoint96Transactor // Write-only binding to the contract
	FixedPoint96Filterer   // Log filterer for contract events
}

// FixedPoint96Caller is an auto generated read-only Go binding around an Ethereum contract.
type FixedPoint96Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedPoint96Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FixedPoint96Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedPoint96Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FixedPoint96Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedPoint96Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FixedPoint96Session struct {
	Contract     *FixedPoint96     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FixedPoint96CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FixedPoint96CallerSession struct {
	Contract *FixedPoint96Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FixedPoint96TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FixedPoint96TransactorSession struct {
	Contract     *FixedPoint96Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FixedPoint96Raw is an auto generated low-level Go binding around an Ethereum contract.
type FixedPoint96Raw struct {
	Contract *FixedPoint96 // Generic contract binding to access the raw methods on
}

// FixedPoint96CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FixedPoint96CallerRaw struct {
	Contract *FixedPoint96Caller // Generic read-only contract binding to access the raw methods on
}

// FixedPoint96TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FixedPoint96TransactorRaw struct {
	Contract *FixedPoint96Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFixedPoint96 creates a new instance of FixedPoint96, bound to a specific deployed contract.
func NewFixedPoint96(address common.Address, backend bind.ContractBackend) (*FixedPoint96, error) {
	contract, err := bindFixedPoint96(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FixedPoint96{FixedPoint96Caller: FixedPoint96Caller{contract: contract}, FixedPoint96Transactor: FixedPoint96Transactor{contract: contract}, FixedPoint96Filterer: FixedPoint96Filterer{contract: contract}}, nil
}

// NewFixedPoint96Caller creates a new read-only instance of FixedPoint96, bound to a specific deployed contract.
func NewFixedPoint96Caller(address common.Address, caller bind.ContractCaller) (*FixedPoint96Caller, error) {
	contract, err := bindFixedPoint96(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FixedPoint96Caller{contract: contract}, nil
}

// NewFixedPoint96Transactor creates a new write-only instance of FixedPoint96, bound to a specific deployed contract.
func NewFixedPoint96Transactor(address common.Address, transactor bind.ContractTransactor) (*FixedPoint96Transactor, error) {
	contract, err := bindFixedPoint96(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FixedPoint96Transactor{contract: contract}, nil
}

// NewFixedPoint96Filterer creates a new log filterer instance of FixedPoint96, bound to a specific deployed contract.
func NewFixedPoint96Filterer(address common.Address, filterer bind.ContractFilterer) (*FixedPoint96Filterer, error) {
	contract, err := bindFixedPoint96(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FixedPoint96Filterer{contract: contract}, nil
}

// bindFixedPoint96 binds a generic wrapper to an already deployed contract.
func bindFixedPoint96(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FixedPoint96ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FixedPoint96 *FixedPoint96Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FixedPoint96.Contract.FixedPoint96Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FixedPoint96 *FixedPoint96Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FixedPoint96.Contract.FixedPoint96Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FixedPoint96 *FixedPoint96Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FixedPoint96.Contract.FixedPoint96Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FixedPoint96 *FixedPoint96CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FixedPoint96.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FixedPoint96 *FixedPoint96TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FixedPoint96.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FixedPoint96 *FixedPoint96TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FixedPoint96.Contract.contract.Transact(opts, method, params...)
}

