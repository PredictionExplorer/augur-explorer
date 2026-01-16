package main

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)
// FixedPoint128MetaData contains all meta data concerning the FixedPoint128 contract.
var FixedPoint128MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f4d9c2a1841b99bdf0af786e49218c51b05798226cd2b06e38c48e0540c162a364736f6c63430007060033",
}

// FixedPoint128ABI is the input ABI used to generate the binding from.
// Deprecated: Use FixedPoint128MetaData.ABI instead.
var FixedPoint128ABI = FixedPoint128MetaData.ABI

// FixedPoint128Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FixedPoint128MetaData.Bin instead.
var FixedPoint128Bin = FixedPoint128MetaData.Bin

// DeployFixedPoint128 deploys a new Ethereum contract, binding an instance of FixedPoint128 to it.
func DeployFixedPoint128(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FixedPoint128, error) {
	parsed, err := FixedPoint128MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FixedPoint128Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FixedPoint128{FixedPoint128Caller: FixedPoint128Caller{contract: contract}, FixedPoint128Transactor: FixedPoint128Transactor{contract: contract}, FixedPoint128Filterer: FixedPoint128Filterer{contract: contract}}, nil
}

// FixedPoint128 is an auto generated Go binding around an Ethereum contract.
type FixedPoint128 struct {
	FixedPoint128Caller     // Read-only binding to the contract
	FixedPoint128Transactor // Write-only binding to the contract
	FixedPoint128Filterer   // Log filterer for contract events
}

// FixedPoint128Caller is an auto generated read-only Go binding around an Ethereum contract.
type FixedPoint128Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedPoint128Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FixedPoint128Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedPoint128Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FixedPoint128Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedPoint128Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FixedPoint128Session struct {
	Contract     *FixedPoint128    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FixedPoint128CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FixedPoint128CallerSession struct {
	Contract *FixedPoint128Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FixedPoint128TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FixedPoint128TransactorSession struct {
	Contract     *FixedPoint128Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FixedPoint128Raw is an auto generated low-level Go binding around an Ethereum contract.
type FixedPoint128Raw struct {
	Contract *FixedPoint128 // Generic contract binding to access the raw methods on
}

// FixedPoint128CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FixedPoint128CallerRaw struct {
	Contract *FixedPoint128Caller // Generic read-only contract binding to access the raw methods on
}

// FixedPoint128TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FixedPoint128TransactorRaw struct {
	Contract *FixedPoint128Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFixedPoint128 creates a new instance of FixedPoint128, bound to a specific deployed contract.
func NewFixedPoint128(address common.Address, backend bind.ContractBackend) (*FixedPoint128, error) {
	contract, err := bindFixedPoint128(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FixedPoint128{FixedPoint128Caller: FixedPoint128Caller{contract: contract}, FixedPoint128Transactor: FixedPoint128Transactor{contract: contract}, FixedPoint128Filterer: FixedPoint128Filterer{contract: contract}}, nil
}

// NewFixedPoint128Caller creates a new read-only instance of FixedPoint128, bound to a specific deployed contract.
func NewFixedPoint128Caller(address common.Address, caller bind.ContractCaller) (*FixedPoint128Caller, error) {
	contract, err := bindFixedPoint128(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FixedPoint128Caller{contract: contract}, nil
}

// NewFixedPoint128Transactor creates a new write-only instance of FixedPoint128, bound to a specific deployed contract.
func NewFixedPoint128Transactor(address common.Address, transactor bind.ContractTransactor) (*FixedPoint128Transactor, error) {
	contract, err := bindFixedPoint128(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FixedPoint128Transactor{contract: contract}, nil
}

// NewFixedPoint128Filterer creates a new log filterer instance of FixedPoint128, bound to a specific deployed contract.
func NewFixedPoint128Filterer(address common.Address, filterer bind.ContractFilterer) (*FixedPoint128Filterer, error) {
	contract, err := bindFixedPoint128(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FixedPoint128Filterer{contract: contract}, nil
}

// bindFixedPoint128 binds a generic wrapper to an already deployed contract.
func bindFixedPoint128(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FixedPoint128ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FixedPoint128 *FixedPoint128Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FixedPoint128.Contract.FixedPoint128Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FixedPoint128 *FixedPoint128Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FixedPoint128.Contract.FixedPoint128Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FixedPoint128 *FixedPoint128Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FixedPoint128.Contract.FixedPoint128Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FixedPoint128 *FixedPoint128CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FixedPoint128.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FixedPoint128 *FixedPoint128TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FixedPoint128.Contract.contract.Transfer(opts)
}

