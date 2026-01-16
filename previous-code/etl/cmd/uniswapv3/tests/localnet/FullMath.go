
package main

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// FullMathMetaData contains all meta data concerning the FullMath contract.
var FullMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d9c52e9c7e34e068308ad271f5f2baedafb3361f49887678f3b8c1bb3899b7f664736f6c63430007060033",
}

// FullMathABI is the input ABI used to generate the binding from.
// Deprecated: Use FullMathMetaData.ABI instead.
var FullMathABI = FullMathMetaData.ABI

// FullMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FullMathMetaData.Bin instead.
var FullMathBin = FullMathMetaData.Bin

// DeployFullMath deploys a new Ethereum contract, binding an instance of FullMath to it.
func DeployFullMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FullMath, error) {
	parsed, err := FullMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FullMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FullMath{FullMathCaller: FullMathCaller{contract: contract}, FullMathTransactor: FullMathTransactor{contract: contract}, FullMathFilterer: FullMathFilterer{contract: contract}}, nil
}

// FullMath is an auto generated Go binding around an Ethereum contract.
type FullMath struct {
	FullMathCaller     // Read-only binding to the contract
	FullMathTransactor // Write-only binding to the contract
	FullMathFilterer   // Log filterer for contract events
}

// FullMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type FullMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FullMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FullMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FullMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FullMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FullMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FullMathSession struct {
	Contract     *FullMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FullMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FullMathCallerSession struct {
	Contract *FullMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FullMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FullMathTransactorSession struct {
	Contract     *FullMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FullMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type FullMathRaw struct {
	Contract *FullMath // Generic contract binding to access the raw methods on
}

// FullMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FullMathCallerRaw struct {
	Contract *FullMathCaller // Generic read-only contract binding to access the raw methods on
}

// FullMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FullMathTransactorRaw struct {
	Contract *FullMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFullMath creates a new instance of FullMath, bound to a specific deployed contract.
func NewFullMath(address common.Address, backend bind.ContractBackend) (*FullMath, error) {
	contract, err := bindFullMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FullMath{FullMathCaller: FullMathCaller{contract: contract}, FullMathTransactor: FullMathTransactor{contract: contract}, FullMathFilterer: FullMathFilterer{contract: contract}}, nil
}

// NewFullMathCaller creates a new read-only instance of FullMath, bound to a specific deployed contract.
func NewFullMathCaller(address common.Address, caller bind.ContractCaller) (*FullMathCaller, error) {
	contract, err := bindFullMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FullMathCaller{contract: contract}, nil
}

// NewFullMathTransactor creates a new write-only instance of FullMath, bound to a specific deployed contract.
func NewFullMathTransactor(address common.Address, transactor bind.ContractTransactor) (*FullMathTransactor, error) {
	contract, err := bindFullMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FullMathTransactor{contract: contract}, nil
}

// NewFullMathFilterer creates a new log filterer instance of FullMath, bound to a specific deployed contract.
func NewFullMathFilterer(address common.Address, filterer bind.ContractFilterer) (*FullMathFilterer, error) {
	contract, err := bindFullMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FullMathFilterer{contract: contract}, nil
}

// bindFullMath binds a generic wrapper to an already deployed contract.
func bindFullMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FullMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FullMath *FullMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FullMath.Contract.FullMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FullMath *FullMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FullMath.Contract.FullMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FullMath *FullMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FullMath.Contract.FullMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FullMath *FullMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FullMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FullMath *FullMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FullMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FullMath *FullMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FullMath.Contract.contract.Transact(opts, method, params...)
}

