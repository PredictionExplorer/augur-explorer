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

