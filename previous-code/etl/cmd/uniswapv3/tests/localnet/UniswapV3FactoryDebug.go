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

// BitMathMetaData contains all meta data concerning the BitMath contract.
var BitMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b4068228482c815c2f7d0c04c06284eeb77b50db177e6aa18acdde2fbc3c86a364736f6c63430007060033",
}

// BitMathABI is the input ABI used to generate the binding from.
// Deprecated: Use BitMathMetaData.ABI instead.
var BitMathABI = BitMathMetaData.ABI

// BitMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BitMathMetaData.Bin instead.
var BitMathBin = BitMathMetaData.Bin

// DeployBitMath deploys a new Ethereum contract, binding an instance of BitMath to it.
func DeployBitMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BitMath, error) {
	parsed, err := BitMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BitMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BitMath{BitMathCaller: BitMathCaller{contract: contract}, BitMathTransactor: BitMathTransactor{contract: contract}, BitMathFilterer: BitMathFilterer{contract: contract}}, nil
}

// BitMath is an auto generated Go binding around an Ethereum contract.
type BitMath struct {
	BitMathCaller     // Read-only binding to the contract
	BitMathTransactor // Write-only binding to the contract
	BitMathFilterer   // Log filterer for contract events
}

// BitMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type BitMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BitMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitMathSession struct {
	Contract     *BitMath          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BitMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitMathCallerSession struct {
	Contract *BitMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BitMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitMathTransactorSession struct {
	Contract     *BitMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BitMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type BitMathRaw struct {
	Contract *BitMath // Generic contract binding to access the raw methods on
}

// BitMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitMathCallerRaw struct {
	Contract *BitMathCaller // Generic read-only contract binding to access the raw methods on
}

// BitMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitMathTransactorRaw struct {
	Contract *BitMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBitMath creates a new instance of BitMath, bound to a specific deployed contract.
func NewBitMath(address common.Address, backend bind.ContractBackend) (*BitMath, error) {
	contract, err := bindBitMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BitMath{BitMathCaller: BitMathCaller{contract: contract}, BitMathTransactor: BitMathTransactor{contract: contract}, BitMathFilterer: BitMathFilterer{contract: contract}}, nil
}

// NewBitMathCaller creates a new read-only instance of BitMath, bound to a specific deployed contract.
func NewBitMathCaller(address common.Address, caller bind.ContractCaller) (*BitMathCaller, error) {
	contract, err := bindBitMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitMathCaller{contract: contract}, nil
}

// NewBitMathTransactor creates a new write-only instance of BitMath, bound to a specific deployed contract.
func NewBitMathTransactor(address common.Address, transactor bind.ContractTransactor) (*BitMathTransactor, error) {
	contract, err := bindBitMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitMathTransactor{contract: contract}, nil
}

// NewBitMathFilterer creates a new log filterer instance of BitMath, bound to a specific deployed contract.
func NewBitMathFilterer(address common.Address, filterer bind.ContractFilterer) (*BitMathFilterer, error) {
	contract, err := bindBitMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitMathFilterer{contract: contract}, nil
}

// bindBitMath binds a generic wrapper to an already deployed contract.
func bindBitMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BitMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitMath *BitMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitMath.Contract.BitMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitMath *BitMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitMath.Contract.BitMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitMath *BitMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitMath.Contract.BitMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BitMath *BitMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BitMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BitMath *BitMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BitMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BitMath *BitMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BitMath.Contract.contract.Transact(opts, method, params...)
}

// FixedPoint128MetaData contains all meta data concerning the FixedPoint128 contract.
var FixedPoint128MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fccfda7ab36329bf749c78ec01881056cd5a49d9625b41765a0540c6e3905bdf64736f6c63430007060033",
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

// Transact invokes the (paid) contract method with params as input values.
func (_FixedPoint128 *FixedPoint128TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FixedPoint128.Contract.contract.Transact(opts, method, params...)
}

// FixedPoint96MetaData contains all meta data concerning the FixedPoint96 contract.
var FixedPoint96MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220760d5300d5d601a0c67bc0869b30e566e2c181e7f008446833fdb1a560d9965e64736f6c63430007060033",
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

// FullMathMetaData contains all meta data concerning the FullMath contract.
var FullMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220323a5886215b3a45676230f1f09733acf0cfbe56758188931b0a6dcbe1f47e2764736f6c63430007060033",
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

// IERC20MinimalMetaData contains all meta data concerning the IERC20Minimal contract.
var IERC20MinimalMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20MinimalABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MinimalMetaData.ABI instead.
var IERC20MinimalABI = IERC20MinimalMetaData.ABI

// IERC20Minimal is an auto generated Go binding around an Ethereum contract.
type IERC20Minimal struct {
	IERC20MinimalCaller     // Read-only binding to the contract
	IERC20MinimalTransactor // Write-only binding to the contract
	IERC20MinimalFilterer   // Log filterer for contract events
}

// IERC20MinimalCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MinimalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MinimalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MinimalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MinimalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MinimalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MinimalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MinimalSession struct {
	Contract     *IERC20Minimal    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MinimalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MinimalCallerSession struct {
	Contract *IERC20MinimalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IERC20MinimalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MinimalTransactorSession struct {
	Contract     *IERC20MinimalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC20MinimalRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MinimalRaw struct {
	Contract *IERC20Minimal // Generic contract binding to access the raw methods on
}

// IERC20MinimalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MinimalCallerRaw struct {
	Contract *IERC20MinimalCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MinimalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MinimalTransactorRaw struct {
	Contract *IERC20MinimalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Minimal creates a new instance of IERC20Minimal, bound to a specific deployed contract.
func NewIERC20Minimal(address common.Address, backend bind.ContractBackend) (*IERC20Minimal, error) {
	contract, err := bindIERC20Minimal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Minimal{IERC20MinimalCaller: IERC20MinimalCaller{contract: contract}, IERC20MinimalTransactor: IERC20MinimalTransactor{contract: contract}, IERC20MinimalFilterer: IERC20MinimalFilterer{contract: contract}}, nil
}

// NewIERC20MinimalCaller creates a new read-only instance of IERC20Minimal, bound to a specific deployed contract.
func NewIERC20MinimalCaller(address common.Address, caller bind.ContractCaller) (*IERC20MinimalCaller, error) {
	contract, err := bindIERC20Minimal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MinimalCaller{contract: contract}, nil
}

// NewIERC20MinimalTransactor creates a new write-only instance of IERC20Minimal, bound to a specific deployed contract.
func NewIERC20MinimalTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MinimalTransactor, error) {
	contract, err := bindIERC20Minimal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MinimalTransactor{contract: contract}, nil
}

// NewIERC20MinimalFilterer creates a new log filterer instance of IERC20Minimal, bound to a specific deployed contract.
func NewIERC20MinimalFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MinimalFilterer, error) {
	contract, err := bindIERC20Minimal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MinimalFilterer{contract: contract}, nil
}

// bindIERC20Minimal binds a generic wrapper to an already deployed contract.
func bindIERC20Minimal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MinimalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Minimal *IERC20MinimalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Minimal.Contract.IERC20MinimalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Minimal *IERC20MinimalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Minimal.Contract.IERC20MinimalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Minimal *IERC20MinimalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Minimal.Contract.IERC20MinimalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Minimal *IERC20MinimalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Minimal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Minimal *IERC20MinimalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Minimal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Minimal *IERC20MinimalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Minimal.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Minimal *IERC20MinimalCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Minimal.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Minimal *IERC20MinimalSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Minimal.Contract.Allowance(&_IERC20Minimal.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Minimal *IERC20MinimalCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Minimal.Contract.Allowance(&_IERC20Minimal.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Minimal *IERC20MinimalCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Minimal.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Minimal *IERC20MinimalSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Minimal.Contract.BalanceOf(&_IERC20Minimal.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Minimal *IERC20MinimalCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Minimal.Contract.BalanceOf(&_IERC20Minimal.CallOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Minimal *IERC20MinimalTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minimal.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Minimal *IERC20MinimalSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minimal.Contract.Approve(&_IERC20Minimal.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Minimal *IERC20MinimalTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minimal.Contract.Approve(&_IERC20Minimal.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Minimal *IERC20MinimalTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minimal.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Minimal *IERC20MinimalSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minimal.Contract.Transfer(&_IERC20Minimal.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Minimal *IERC20MinimalTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minimal.Contract.Transfer(&_IERC20Minimal.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Minimal *IERC20MinimalTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minimal.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Minimal *IERC20MinimalSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minimal.Contract.TransferFrom(&_IERC20Minimal.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Minimal *IERC20MinimalTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minimal.Contract.TransferFrom(&_IERC20Minimal.TransactOpts, sender, recipient, amount)
}

// IERC20MinimalApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Minimal contract.
type IERC20MinimalApprovalIterator struct {
	Event *IERC20MinimalApproval // Event containing the contract specifics and raw log

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
func (it *IERC20MinimalApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MinimalApproval)
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
		it.Event = new(IERC20MinimalApproval)
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
func (it *IERC20MinimalApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MinimalApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MinimalApproval represents a Approval event raised by the IERC20Minimal contract.
type IERC20MinimalApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Minimal *IERC20MinimalFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MinimalApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Minimal.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MinimalApprovalIterator{contract: _IERC20Minimal.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Minimal *IERC20MinimalFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MinimalApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Minimal.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MinimalApproval)
				if err := _IERC20Minimal.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20Minimal *IERC20MinimalFilterer) ParseApproval(log types.Log) (*IERC20MinimalApproval, error) {
	event := new(IERC20MinimalApproval)
	if err := _IERC20Minimal.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MinimalTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Minimal contract.
type IERC20MinimalTransferIterator struct {
	Event *IERC20MinimalTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20MinimalTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MinimalTransfer)
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
		it.Event = new(IERC20MinimalTransfer)
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
func (it *IERC20MinimalTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MinimalTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MinimalTransfer represents a Transfer event raised by the IERC20Minimal contract.
type IERC20MinimalTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Minimal *IERC20MinimalFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MinimalTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Minimal.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MinimalTransferIterator{contract: _IERC20Minimal.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Minimal *IERC20MinimalFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MinimalTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Minimal.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MinimalTransfer)
				if err := _IERC20Minimal.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20Minimal *IERC20MinimalFilterer) ParseTransfer(log types.Log) (*IERC20MinimalTransfer, error) {
	event := new(IERC20MinimalTransfer)
	if err := _IERC20Minimal.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3FactoryMetaData contains all meta data concerning the IUniswapV3Factory contract.
var IUniswapV3FactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"FeeAmountEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"enableFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IUniswapV3FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3FactoryMetaData.ABI instead.
var IUniswapV3FactoryABI = IUniswapV3FactoryMetaData.ABI

// IUniswapV3Factory is an auto generated Go binding around an Ethereum contract.
type IUniswapV3Factory struct {
	IUniswapV3FactoryCaller     // Read-only binding to the contract
	IUniswapV3FactoryTransactor // Write-only binding to the contract
	IUniswapV3FactoryFilterer   // Log filterer for contract events
}

// IUniswapV3FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3FactorySession struct {
	Contract     *IUniswapV3Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IUniswapV3FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3FactoryCallerSession struct {
	Contract *IUniswapV3FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IUniswapV3FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3FactoryTransactorSession struct {
	Contract     *IUniswapV3FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IUniswapV3FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3FactoryRaw struct {
	Contract *IUniswapV3Factory // Generic contract binding to access the raw methods on
}

// IUniswapV3FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3FactoryCallerRaw struct {
	Contract *IUniswapV3FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3FactoryTransactorRaw struct {
	Contract *IUniswapV3FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3Factory creates a new instance of IUniswapV3Factory, bound to a specific deployed contract.
func NewIUniswapV3Factory(address common.Address, backend bind.ContractBackend) (*IUniswapV3Factory, error) {
	contract, err := bindIUniswapV3Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3Factory{IUniswapV3FactoryCaller: IUniswapV3FactoryCaller{contract: contract}, IUniswapV3FactoryTransactor: IUniswapV3FactoryTransactor{contract: contract}, IUniswapV3FactoryFilterer: IUniswapV3FactoryFilterer{contract: contract}}, nil
}

// NewIUniswapV3FactoryCaller creates a new read-only instance of IUniswapV3Factory, bound to a specific deployed contract.
func NewIUniswapV3FactoryCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3FactoryCaller, error) {
	contract, err := bindIUniswapV3Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3FactoryCaller{contract: contract}, nil
}

// NewIUniswapV3FactoryTransactor creates a new write-only instance of IUniswapV3Factory, bound to a specific deployed contract.
func NewIUniswapV3FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3FactoryTransactor, error) {
	contract, err := bindIUniswapV3Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3FactoryTransactor{contract: contract}, nil
}

// NewIUniswapV3FactoryFilterer creates a new log filterer instance of IUniswapV3Factory, bound to a specific deployed contract.
func NewIUniswapV3FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3FactoryFilterer, error) {
	contract, err := bindIUniswapV3Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3FactoryFilterer{contract: contract}, nil
}

// bindIUniswapV3Factory binds a generic wrapper to an already deployed contract.
func bindIUniswapV3Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3Factory *IUniswapV3FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3Factory.Contract.IUniswapV3FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3Factory *IUniswapV3FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3Factory.Contract.IUniswapV3FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3Factory *IUniswapV3FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3Factory.Contract.IUniswapV3FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3Factory *IUniswapV3FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3Factory *IUniswapV3FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3Factory *IUniswapV3FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3Factory.Contract.contract.Transact(opts, method, params...)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 fee) view returns(int24)
func (_IUniswapV3Factory *IUniswapV3FactoryCaller) FeeAmountTickSpacing(opts *bind.CallOpts, fee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapV3Factory.contract.Call(opts, &out, "feeAmountTickSpacing", fee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 fee) view returns(int24)
func (_IUniswapV3Factory *IUniswapV3FactorySession) FeeAmountTickSpacing(fee *big.Int) (*big.Int, error) {
	return _IUniswapV3Factory.Contract.FeeAmountTickSpacing(&_IUniswapV3Factory.CallOpts, fee)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 fee) view returns(int24)
func (_IUniswapV3Factory *IUniswapV3FactoryCallerSession) FeeAmountTickSpacing(fee *big.Int) (*big.Int, error) {
	return _IUniswapV3Factory.Contract.FeeAmountTickSpacing(&_IUniswapV3Factory.CallOpts, fee)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address pool)
func (_IUniswapV3Factory *IUniswapV3FactoryCaller) GetPool(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV3Factory.contract.Call(opts, &out, "getPool", tokenA, tokenB, fee)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address pool)
func (_IUniswapV3Factory *IUniswapV3FactorySession) GetPool(tokenA common.Address, tokenB common.Address, fee *big.Int) (common.Address, error) {
	return _IUniswapV3Factory.Contract.GetPool(&_IUniswapV3Factory.CallOpts, tokenA, tokenB, fee)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address pool)
func (_IUniswapV3Factory *IUniswapV3FactoryCallerSession) GetPool(tokenA common.Address, tokenB common.Address, fee *big.Int) (common.Address, error) {
	return _IUniswapV3Factory.Contract.GetPool(&_IUniswapV3Factory.CallOpts, tokenA, tokenB, fee)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IUniswapV3Factory *IUniswapV3FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV3Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IUniswapV3Factory *IUniswapV3FactorySession) Owner() (common.Address, error) {
	return _IUniswapV3Factory.Contract.Owner(&_IUniswapV3Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IUniswapV3Factory *IUniswapV3FactoryCallerSession) Owner() (common.Address, error) {
	return _IUniswapV3Factory.Contract.Owner(&_IUniswapV3Factory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_IUniswapV3Factory *IUniswapV3FactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Factory.contract.Transact(opts, "createPool", tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_IUniswapV3Factory *IUniswapV3FactorySession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Factory.Contract.CreatePool(&_IUniswapV3Factory.TransactOpts, tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_IUniswapV3Factory *IUniswapV3FactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Factory.Contract.CreatePool(&_IUniswapV3Factory.TransactOpts, tokenA, tokenB, fee)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_IUniswapV3Factory *IUniswapV3FactoryTransactor) EnableFeeAmount(opts *bind.TransactOpts, fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Factory.contract.Transact(opts, "enableFeeAmount", fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_IUniswapV3Factory *IUniswapV3FactorySession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Factory.Contract.EnableFeeAmount(&_IUniswapV3Factory.TransactOpts, fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_IUniswapV3Factory *IUniswapV3FactoryTransactorSession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _IUniswapV3Factory.Contract.EnableFeeAmount(&_IUniswapV3Factory.TransactOpts, fee, tickSpacing)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_IUniswapV3Factory *IUniswapV3FactoryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _IUniswapV3Factory.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_IUniswapV3Factory *IUniswapV3FactorySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _IUniswapV3Factory.Contract.SetOwner(&_IUniswapV3Factory.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_IUniswapV3Factory *IUniswapV3FactoryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _IUniswapV3Factory.Contract.SetOwner(&_IUniswapV3Factory.TransactOpts, _owner)
}

// IUniswapV3FactoryFeeAmountEnabledIterator is returned from FilterFeeAmountEnabled and is used to iterate over the raw logs and unpacked data for FeeAmountEnabled events raised by the IUniswapV3Factory contract.
type IUniswapV3FactoryFeeAmountEnabledIterator struct {
	Event *IUniswapV3FactoryFeeAmountEnabled // Event containing the contract specifics and raw log

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
func (it *IUniswapV3FactoryFeeAmountEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3FactoryFeeAmountEnabled)
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
		it.Event = new(IUniswapV3FactoryFeeAmountEnabled)
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
func (it *IUniswapV3FactoryFeeAmountEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3FactoryFeeAmountEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3FactoryFeeAmountEnabled represents a FeeAmountEnabled event raised by the IUniswapV3Factory contract.
type IUniswapV3FactoryFeeAmountEnabled struct {
	Fee         *big.Int
	TickSpacing *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeAmountEnabled is a free log retrieval operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_IUniswapV3Factory *IUniswapV3FactoryFilterer) FilterFeeAmountEnabled(opts *bind.FilterOpts, fee []*big.Int, tickSpacing []*big.Int) (*IUniswapV3FactoryFeeAmountEnabledIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _IUniswapV3Factory.contract.FilterLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3FactoryFeeAmountEnabledIterator{contract: _IUniswapV3Factory.contract, event: "FeeAmountEnabled", logs: logs, sub: sub}, nil
}

// WatchFeeAmountEnabled is a free log subscription operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_IUniswapV3Factory *IUniswapV3FactoryFilterer) WatchFeeAmountEnabled(opts *bind.WatchOpts, sink chan<- *IUniswapV3FactoryFeeAmountEnabled, fee []*big.Int, tickSpacing []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _IUniswapV3Factory.contract.WatchLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3FactoryFeeAmountEnabled)
				if err := _IUniswapV3Factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
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

// ParseFeeAmountEnabled is a log parse operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_IUniswapV3Factory *IUniswapV3FactoryFilterer) ParseFeeAmountEnabled(log types.Log) (*IUniswapV3FactoryFeeAmountEnabled, error) {
	event := new(IUniswapV3FactoryFeeAmountEnabled)
	if err := _IUniswapV3Factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3FactoryOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the IUniswapV3Factory contract.
type IUniswapV3FactoryOwnerChangedIterator struct {
	Event *IUniswapV3FactoryOwnerChanged // Event containing the contract specifics and raw log

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
func (it *IUniswapV3FactoryOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3FactoryOwnerChanged)
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
		it.Event = new(IUniswapV3FactoryOwnerChanged)
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
func (it *IUniswapV3FactoryOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3FactoryOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3FactoryOwnerChanged represents a OwnerChanged event raised by the IUniswapV3Factory contract.
type IUniswapV3FactoryOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_IUniswapV3Factory *IUniswapV3FactoryFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*IUniswapV3FactoryOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IUniswapV3Factory.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3FactoryOwnerChangedIterator{contract: _IUniswapV3Factory.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_IUniswapV3Factory *IUniswapV3FactoryFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *IUniswapV3FactoryOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IUniswapV3Factory.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3FactoryOwnerChanged)
				if err := _IUniswapV3Factory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_IUniswapV3Factory *IUniswapV3FactoryFilterer) ParseOwnerChanged(log types.Log) (*IUniswapV3FactoryOwnerChanged, error) {
	event := new(IUniswapV3FactoryOwnerChanged)
	if err := _IUniswapV3Factory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3FactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the IUniswapV3Factory contract.
type IUniswapV3FactoryPoolCreatedIterator struct {
	Event *IUniswapV3FactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *IUniswapV3FactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3FactoryPoolCreated)
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
		it.Event = new(IUniswapV3FactoryPoolCreated)
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
func (it *IUniswapV3FactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3FactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3FactoryPoolCreated represents a PoolCreated event raised by the IUniswapV3Factory contract.
type IUniswapV3FactoryPoolCreated struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Pool        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_IUniswapV3Factory *IUniswapV3FactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, fee []*big.Int) (*IUniswapV3FactoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _IUniswapV3Factory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3FactoryPoolCreatedIterator{contract: _IUniswapV3Factory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_IUniswapV3Factory *IUniswapV3FactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *IUniswapV3FactoryPoolCreated, token0 []common.Address, token1 []common.Address, fee []*big.Int) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _IUniswapV3Factory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3FactoryPoolCreated)
				if err := _IUniswapV3Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_IUniswapV3Factory *IUniswapV3FactoryFilterer) ParsePoolCreated(log types.Log) (*IUniswapV3FactoryPoolCreated, error) {
	event := new(IUniswapV3FactoryPoolCreated)
	if err := _IUniswapV3Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3FlashCallbackMetaData contains all meta data concerning the IUniswapV3FlashCallback contract.
var IUniswapV3FlashCallbackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3FlashCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IUniswapV3FlashCallbackABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3FlashCallbackMetaData.ABI instead.
var IUniswapV3FlashCallbackABI = IUniswapV3FlashCallbackMetaData.ABI

// IUniswapV3FlashCallback is an auto generated Go binding around an Ethereum contract.
type IUniswapV3FlashCallback struct {
	IUniswapV3FlashCallbackCaller     // Read-only binding to the contract
	IUniswapV3FlashCallbackTransactor // Write-only binding to the contract
	IUniswapV3FlashCallbackFilterer   // Log filterer for contract events
}

// IUniswapV3FlashCallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3FlashCallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3FlashCallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3FlashCallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3FlashCallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3FlashCallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3FlashCallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3FlashCallbackSession struct {
	Contract     *IUniswapV3FlashCallback // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IUniswapV3FlashCallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3FlashCallbackCallerSession struct {
	Contract *IUniswapV3FlashCallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// IUniswapV3FlashCallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3FlashCallbackTransactorSession struct {
	Contract     *IUniswapV3FlashCallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// IUniswapV3FlashCallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3FlashCallbackRaw struct {
	Contract *IUniswapV3FlashCallback // Generic contract binding to access the raw methods on
}

// IUniswapV3FlashCallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3FlashCallbackCallerRaw struct {
	Contract *IUniswapV3FlashCallbackCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3FlashCallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3FlashCallbackTransactorRaw struct {
	Contract *IUniswapV3FlashCallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3FlashCallback creates a new instance of IUniswapV3FlashCallback, bound to a specific deployed contract.
func NewIUniswapV3FlashCallback(address common.Address, backend bind.ContractBackend) (*IUniswapV3FlashCallback, error) {
	contract, err := bindIUniswapV3FlashCallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3FlashCallback{IUniswapV3FlashCallbackCaller: IUniswapV3FlashCallbackCaller{contract: contract}, IUniswapV3FlashCallbackTransactor: IUniswapV3FlashCallbackTransactor{contract: contract}, IUniswapV3FlashCallbackFilterer: IUniswapV3FlashCallbackFilterer{contract: contract}}, nil
}

// NewIUniswapV3FlashCallbackCaller creates a new read-only instance of IUniswapV3FlashCallback, bound to a specific deployed contract.
func NewIUniswapV3FlashCallbackCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3FlashCallbackCaller, error) {
	contract, err := bindIUniswapV3FlashCallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3FlashCallbackCaller{contract: contract}, nil
}

// NewIUniswapV3FlashCallbackTransactor creates a new write-only instance of IUniswapV3FlashCallback, bound to a specific deployed contract.
func NewIUniswapV3FlashCallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3FlashCallbackTransactor, error) {
	contract, err := bindIUniswapV3FlashCallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3FlashCallbackTransactor{contract: contract}, nil
}

// NewIUniswapV3FlashCallbackFilterer creates a new log filterer instance of IUniswapV3FlashCallback, bound to a specific deployed contract.
func NewIUniswapV3FlashCallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3FlashCallbackFilterer, error) {
	contract, err := bindIUniswapV3FlashCallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3FlashCallbackFilterer{contract: contract}, nil
}

// bindIUniswapV3FlashCallback binds a generic wrapper to an already deployed contract.
func bindIUniswapV3FlashCallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3FlashCallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3FlashCallback *IUniswapV3FlashCallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3FlashCallback.Contract.IUniswapV3FlashCallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3FlashCallback *IUniswapV3FlashCallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3FlashCallback.Contract.IUniswapV3FlashCallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3FlashCallback *IUniswapV3FlashCallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3FlashCallback.Contract.IUniswapV3FlashCallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3FlashCallback *IUniswapV3FlashCallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3FlashCallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3FlashCallback *IUniswapV3FlashCallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3FlashCallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3FlashCallback *IUniswapV3FlashCallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3FlashCallback.Contract.contract.Transact(opts, method, params...)
}

// UniswapV3FlashCallback is a paid mutator transaction binding the contract method 0xe9cbafb0.
//
// Solidity: function uniswapV3FlashCallback(uint256 fee0, uint256 fee1, bytes data) returns()
func (_IUniswapV3FlashCallback *IUniswapV3FlashCallbackTransactor) UniswapV3FlashCallback(opts *bind.TransactOpts, fee0 *big.Int, fee1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3FlashCallback.contract.Transact(opts, "uniswapV3FlashCallback", fee0, fee1, data)
}

// UniswapV3FlashCallback is a paid mutator transaction binding the contract method 0xe9cbafb0.
//
// Solidity: function uniswapV3FlashCallback(uint256 fee0, uint256 fee1, bytes data) returns()
func (_IUniswapV3FlashCallback *IUniswapV3FlashCallbackSession) UniswapV3FlashCallback(fee0 *big.Int, fee1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3FlashCallback.Contract.UniswapV3FlashCallback(&_IUniswapV3FlashCallback.TransactOpts, fee0, fee1, data)
}

// UniswapV3FlashCallback is a paid mutator transaction binding the contract method 0xe9cbafb0.
//
// Solidity: function uniswapV3FlashCallback(uint256 fee0, uint256 fee1, bytes data) returns()
func (_IUniswapV3FlashCallback *IUniswapV3FlashCallbackTransactorSession) UniswapV3FlashCallback(fee0 *big.Int, fee1 *big.Int, data []byte) (*types.Transaction, error) {
	return _IUniswapV3FlashCallback.Contract.UniswapV3FlashCallback(&_IUniswapV3FlashCallback.TransactOpts, fee0, fee1, data)
}

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

// IUniswapV3PoolMetaData contains all meta data concerning the IUniswapV3Pool contract.
var IUniswapV3PoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"CollectProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"slot0Tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"liquidityDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBefore\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"DBG_MOD_POS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtStartPrice\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimit\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"slot0sqrtpricex96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"DBG_SWAP_HEAD\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceStartX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceNextX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stepAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stepAmountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amountProcessed\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"exactInput\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobalX128\",\"type\":\"uint256\"}],\"name\":\"DBG_SWAP_LOOP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal0X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal1X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityDiff\",\"type\":\"uint128\"}],\"name\":\"DBG_SWAP_TAIL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"liquidityDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal0X128Before\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal1X128Before\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flippedLower\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flippedUpper\",\"type\":\"bool\"}],\"name\":\"DBG_UPD_POS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextOld\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextNew\",\"type\":\"uint16\"}],\"name\":\"IncreaseObservationCardinalityNext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0New\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1New\",\"type\":\"uint8\"}],\"name\":\"SetFeeProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collectProtocol\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal0X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal1X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"}],\"name\":\"increaseObservationCardinalityNext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidityPerTick\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"observations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityCumulativeX128\",\"type\":\"uint160\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"observe\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"tickCumulatives\",\"type\":\"int56[]\"},{\"internalType\":\"uint160[]\",\"name\":\"secondsPerLiquidityCumulativeX128s\",\"type\":\"uint160[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"token0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"token1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"feeProtocol0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol1\",\"type\":\"uint8\"}],\"name\":\"setFeeProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slot0\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"snapshotCumulativesInside\",\"outputs\":[{\"internalType\":\"int56\",\"name\":\"tickCumulativeInside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityInsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsInside\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountSpecified\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"wordPosition\",\"type\":\"int16\"}],\"name\":\"tickBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside0X128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside1X128\",\"type\":\"uint256\"},{\"internalType\":\"int56\",\"name\":\"tickCumulativeOutside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityOutsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsOutside\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// IUniswapV3PoolDBGMODPOSIterator is returned from FilterDBGMODPOS and is used to iterate over the raw logs and unpacked data for DBGMODPOS events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolDBGMODPOSIterator struct {
	Event *IUniswapV3PoolDBGMODPOS // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolDBGMODPOSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolDBGMODPOS)
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
		it.Event = new(IUniswapV3PoolDBGMODPOS)
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
func (it *IUniswapV3PoolDBGMODPOSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolDBGMODPOSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolDBGMODPOS represents a DBGMODPOS event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolDBGMODPOS struct {
	Owner           common.Address
	TickLower       *big.Int
	TickUpper       *big.Int
	Slot0Tick       *big.Int
	LiquidityDelta  *big.Int
	LiquidityBefore *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	SqrtPriceX96    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDBGMODPOS is a free log retrieval operation binding the contract event 0xd1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, int128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterDBGMODPOS(opts *bind.FilterOpts) (*IUniswapV3PoolDBGMODPOSIterator, error) {

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "DBG_MOD_POS")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDBGMODPOSIterator{contract: _IUniswapV3Pool.contract, event: "DBG_MOD_POS", logs: logs, sub: sub}, nil
}

// WatchDBGMODPOS is a free log subscription operation binding the contract event 0xd1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, int128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchDBGMODPOS(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolDBGMODPOS) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "DBG_MOD_POS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolDBGMODPOS)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "DBG_MOD_POS", log); err != nil {
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

// ParseDBGMODPOS is a log parse operation binding the contract event 0xd1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, int128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseDBGMODPOS(log types.Log) (*IUniswapV3PoolDBGMODPOS, error) {
	event := new(IUniswapV3PoolDBGMODPOS)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "DBG_MOD_POS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolDBGSWAPHEADIterator is returned from FilterDBGSWAPHEAD and is used to iterate over the raw logs and unpacked data for DBGSWAPHEAD events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolDBGSWAPHEADIterator struct {
	Event *IUniswapV3PoolDBGSWAPHEAD // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolDBGSWAPHEADIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolDBGSWAPHEAD)
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
		it.Event = new(IUniswapV3PoolDBGSWAPHEAD)
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
func (it *IUniswapV3PoolDBGSWAPHEADIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolDBGSWAPHEADIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolDBGSWAPHEAD represents a DBGSWAPHEAD event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolDBGSWAPHEAD struct {
	Pool              common.Address
	ZeroForOne        bool
	Amount            *big.Int
	SqrtStartPrice    *big.Int
	SqrtPriceLimit    *big.Int
	Slot0sqrtpricex96 *big.Int
	Tick              *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPHEAD is a free log retrieval operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterDBGSWAPHEAD(opts *bind.FilterOpts, pool []common.Address) (*IUniswapV3PoolDBGSWAPHEADIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "DBG_SWAP_HEAD", poolRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDBGSWAPHEADIterator{contract: _IUniswapV3Pool.contract, event: "DBG_SWAP_HEAD", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPHEAD is a free log subscription operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchDBGSWAPHEAD(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolDBGSWAPHEAD, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "DBG_SWAP_HEAD", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolDBGSWAPHEAD)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_HEAD", log); err != nil {
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

// ParseDBGSWAPHEAD is a log parse operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseDBGSWAPHEAD(log types.Log) (*IUniswapV3PoolDBGSWAPHEAD, error) {
	event := new(IUniswapV3PoolDBGSWAPHEAD)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_HEAD", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolDBGSWAPLOOPIterator is returned from FilterDBGSWAPLOOP and is used to iterate over the raw logs and unpacked data for DBGSWAPLOOP events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolDBGSWAPLOOPIterator struct {
	Event *IUniswapV3PoolDBGSWAPLOOP // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolDBGSWAPLOOPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolDBGSWAPLOOP)
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
		it.Event = new(IUniswapV3PoolDBGSWAPLOOP)
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
func (it *IUniswapV3PoolDBGSWAPLOOPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolDBGSWAPLOOPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolDBGSWAPLOOP represents a DBGSWAPLOOP event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolDBGSWAPLOOP struct {
	Pool                common.Address
	SqrtPriceX96        *big.Int
	SqrtPriceStartX96   *big.Int
	SqrtPriceNextX96    *big.Int
	Tick                *big.Int
	TickCumulative      *big.Int
	Initialized         bool
	StepAmountIn        *big.Int
	StepAmountOut       *big.Int
	AmountProcessed     *big.Int
	FeeAmount           *big.Int
	Liquidity           *big.Int
	ExactInput          bool
	FeeGrowthGlobalX128 *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPLOOP is a free log retrieval operation binding the contract event 0xf5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d5929.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int56 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, int256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterDBGSWAPLOOP(opts *bind.FilterOpts, pool []common.Address) (*IUniswapV3PoolDBGSWAPLOOPIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "DBG_SWAP_LOOP", poolRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDBGSWAPLOOPIterator{contract: _IUniswapV3Pool.contract, event: "DBG_SWAP_LOOP", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPLOOP is a free log subscription operation binding the contract event 0xf5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d5929.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int56 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, int256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchDBGSWAPLOOP(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolDBGSWAPLOOP, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "DBG_SWAP_LOOP", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolDBGSWAPLOOP)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_LOOP", log); err != nil {
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

// ParseDBGSWAPLOOP is a log parse operation binding the contract event 0xf5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d5929.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int56 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, int256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseDBGSWAPLOOP(log types.Log) (*IUniswapV3PoolDBGSWAPLOOP, error) {
	event := new(IUniswapV3PoolDBGSWAPLOOP)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_LOOP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolDBGSWAPTAILIterator is returned from FilterDBGSWAPTAIL and is used to iterate over the raw logs and unpacked data for DBGSWAPTAIL events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolDBGSWAPTAILIterator struct {
	Event *IUniswapV3PoolDBGSWAPTAIL // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolDBGSWAPTAILIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolDBGSWAPTAIL)
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
		it.Event = new(IUniswapV3PoolDBGSWAPTAIL)
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
func (it *IUniswapV3PoolDBGSWAPTAILIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolDBGSWAPTAILIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolDBGSWAPTAIL represents a DBGSWAPTAIL event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolDBGSWAPTAIL struct {
	Pool                 common.Address
	FeeGrowthGlobal0X128 *big.Int
	FeeGrowthGlobal1X128 *big.Int
	Tick                 *big.Int
	SqrtPriceX96         *big.Int
	Liquidity            *big.Int
	LiquidityDiff        *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPTAIL is a free log retrieval operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterDBGSWAPTAIL(opts *bind.FilterOpts, pool []common.Address) (*IUniswapV3PoolDBGSWAPTAILIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "DBG_SWAP_TAIL", poolRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDBGSWAPTAILIterator{contract: _IUniswapV3Pool.contract, event: "DBG_SWAP_TAIL", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPTAIL is a free log subscription operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchDBGSWAPTAIL(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolDBGSWAPTAIL, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "DBG_SWAP_TAIL", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolDBGSWAPTAIL)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_TAIL", log); err != nil {
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

// ParseDBGSWAPTAIL is a log parse operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseDBGSWAPTAIL(log types.Log) (*IUniswapV3PoolDBGSWAPTAIL, error) {
	event := new(IUniswapV3PoolDBGSWAPTAIL)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_TAIL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolDBGUPDPOSIterator is returned from FilterDBGUPDPOS and is used to iterate over the raw logs and unpacked data for DBGUPDPOS events raised by the IUniswapV3Pool contract.
type IUniswapV3PoolDBGUPDPOSIterator struct {
	Event *IUniswapV3PoolDBGUPDPOS // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolDBGUPDPOSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolDBGUPDPOS)
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
		it.Event = new(IUniswapV3PoolDBGUPDPOS)
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
func (it *IUniswapV3PoolDBGUPDPOSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolDBGUPDPOSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolDBGUPDPOS represents a DBGUPDPOS event raised by the IUniswapV3Pool contract.
type IUniswapV3PoolDBGUPDPOS struct {
	Owner                      common.Address
	TickLower                  *big.Int
	TickUpper                  *big.Int
	Tick                       *big.Int
	LiquidityDelta             *big.Int
	FeeGrowthGlobal0X128Before *big.Int
	FeeGrowthGlobal1X128Before *big.Int
	FeeGrowthInside0X128       *big.Int
	FeeGrowthInside1X128       *big.Int
	FlippedLower               bool
	FlippedUpper               bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterDBGUPDPOS is a free log retrieval operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) FilterDBGUPDPOS(opts *bind.FilterOpts) (*IUniswapV3PoolDBGUPDPOSIterator, error) {

	logs, sub, err := _IUniswapV3Pool.contract.FilterLogs(opts, "DBG_UPD_POS")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDBGUPDPOSIterator{contract: _IUniswapV3Pool.contract, event: "DBG_UPD_POS", logs: logs, sub: sub}, nil
}

// WatchDBGUPDPOS is a free log subscription operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) WatchDBGUPDPOS(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolDBGUPDPOS) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3Pool.contract.WatchLogs(opts, "DBG_UPD_POS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolDBGUPDPOS)
				if err := _IUniswapV3Pool.contract.UnpackLog(event, "DBG_UPD_POS", log); err != nil {
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

// ParseDBGUPDPOS is a log parse operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_IUniswapV3Pool *IUniswapV3PoolFilterer) ParseDBGUPDPOS(log types.Log) (*IUniswapV3PoolDBGUPDPOS, error) {
	event := new(IUniswapV3PoolDBGUPDPOS)
	if err := _IUniswapV3Pool.contract.UnpackLog(event, "DBG_UPD_POS", log); err != nil {
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

// IUniswapV3PoolDeployerMetaData contains all meta data concerning the IUniswapV3PoolDeployer contract.
var IUniswapV3PoolDeployerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IUniswapV3PoolDeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3PoolDeployerMetaData.ABI instead.
var IUniswapV3PoolDeployerABI = IUniswapV3PoolDeployerMetaData.ABI

// IUniswapV3PoolDeployer is an auto generated Go binding around an Ethereum contract.
type IUniswapV3PoolDeployer struct {
	IUniswapV3PoolDeployerCaller     // Read-only binding to the contract
	IUniswapV3PoolDeployerTransactor // Write-only binding to the contract
	IUniswapV3PoolDeployerFilterer   // Log filterer for contract events
}

// IUniswapV3PoolDeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3PoolDeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolDeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3PoolDeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolDeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3PoolDeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolDeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3PoolDeployerSession struct {
	Contract     *IUniswapV3PoolDeployer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IUniswapV3PoolDeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3PoolDeployerCallerSession struct {
	Contract *IUniswapV3PoolDeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IUniswapV3PoolDeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3PoolDeployerTransactorSession struct {
	Contract     *IUniswapV3PoolDeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IUniswapV3PoolDeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3PoolDeployerRaw struct {
	Contract *IUniswapV3PoolDeployer // Generic contract binding to access the raw methods on
}

// IUniswapV3PoolDeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3PoolDeployerCallerRaw struct {
	Contract *IUniswapV3PoolDeployerCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3PoolDeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3PoolDeployerTransactorRaw struct {
	Contract *IUniswapV3PoolDeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3PoolDeployer creates a new instance of IUniswapV3PoolDeployer, bound to a specific deployed contract.
func NewIUniswapV3PoolDeployer(address common.Address, backend bind.ContractBackend) (*IUniswapV3PoolDeployer, error) {
	contract, err := bindIUniswapV3PoolDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDeployer{IUniswapV3PoolDeployerCaller: IUniswapV3PoolDeployerCaller{contract: contract}, IUniswapV3PoolDeployerTransactor: IUniswapV3PoolDeployerTransactor{contract: contract}, IUniswapV3PoolDeployerFilterer: IUniswapV3PoolDeployerFilterer{contract: contract}}, nil
}

// NewIUniswapV3PoolDeployerCaller creates a new read-only instance of IUniswapV3PoolDeployer, bound to a specific deployed contract.
func NewIUniswapV3PoolDeployerCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3PoolDeployerCaller, error) {
	contract, err := bindIUniswapV3PoolDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDeployerCaller{contract: contract}, nil
}

// NewIUniswapV3PoolDeployerTransactor creates a new write-only instance of IUniswapV3PoolDeployer, bound to a specific deployed contract.
func NewIUniswapV3PoolDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3PoolDeployerTransactor, error) {
	contract, err := bindIUniswapV3PoolDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDeployerTransactor{contract: contract}, nil
}

// NewIUniswapV3PoolDeployerFilterer creates a new log filterer instance of IUniswapV3PoolDeployer, bound to a specific deployed contract.
func NewIUniswapV3PoolDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3PoolDeployerFilterer, error) {
	contract, err := bindIUniswapV3PoolDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolDeployerFilterer{contract: contract}, nil
}

// bindIUniswapV3PoolDeployer binds a generic wrapper to an already deployed contract.
func bindIUniswapV3PoolDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3PoolDeployerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolDeployer *IUniswapV3PoolDeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolDeployer.Contract.IUniswapV3PoolDeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolDeployer *IUniswapV3PoolDeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolDeployer.Contract.IUniswapV3PoolDeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolDeployer *IUniswapV3PoolDeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolDeployer.Contract.IUniswapV3PoolDeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolDeployer *IUniswapV3PoolDeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolDeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolDeployer *IUniswapV3PoolDeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolDeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolDeployer *IUniswapV3PoolDeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolDeployer.Contract.contract.Transact(opts, method, params...)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_IUniswapV3PoolDeployer *IUniswapV3PoolDeployerCaller) Parameters(opts *bind.CallOpts) (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	var out []interface{}
	err := _IUniswapV3PoolDeployer.contract.Call(opts, &out, "parameters")

	outstruct := new(struct {
		Factory     common.Address
		Token0      common.Address
		Token1      common.Address
		Fee         *big.Int
		TickSpacing *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Factory = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickSpacing = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_IUniswapV3PoolDeployer *IUniswapV3PoolDeployerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _IUniswapV3PoolDeployer.Contract.Parameters(&_IUniswapV3PoolDeployer.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_IUniswapV3PoolDeployer *IUniswapV3PoolDeployerCallerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _IUniswapV3PoolDeployer.Contract.Parameters(&_IUniswapV3PoolDeployer.CallOpts)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"CollectProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"slot0Tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"liquidityDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBefore\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"DBG_MOD_POS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtStartPrice\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimit\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"slot0sqrtpricex96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"DBG_SWAP_HEAD\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceStartX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceNextX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stepAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stepAmountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amountProcessed\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"exactInput\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobalX128\",\"type\":\"uint256\"}],\"name\":\"DBG_SWAP_LOOP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal0X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal1X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityDiff\",\"type\":\"uint128\"}],\"name\":\"DBG_SWAP_TAIL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"liquidityDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal0X128Before\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal1X128Before\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flippedLower\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flippedUpper\",\"type\":\"bool\"}],\"name\":\"DBG_UPD_POS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextOld\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextNew\",\"type\":\"uint16\"}],\"name\":\"IncreaseObservationCardinalityNext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0New\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1New\",\"type\":\"uint8\"}],\"name\":\"SetFeeProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"}]",
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

// IUniswapV3PoolEventsDBGMODPOSIterator is returned from FilterDBGMODPOS and is used to iterate over the raw logs and unpacked data for DBGMODPOS events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGMODPOSIterator struct {
	Event *IUniswapV3PoolEventsDBGMODPOS // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsDBGMODPOSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsDBGMODPOS)
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
		it.Event = new(IUniswapV3PoolEventsDBGMODPOS)
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
func (it *IUniswapV3PoolEventsDBGMODPOSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsDBGMODPOSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsDBGMODPOS represents a DBGMODPOS event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGMODPOS struct {
	Owner           common.Address
	TickLower       *big.Int
	TickUpper       *big.Int
	Slot0Tick       *big.Int
	LiquidityDelta  *big.Int
	LiquidityBefore *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	SqrtPriceX96    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDBGMODPOS is a free log retrieval operation binding the contract event 0xd1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, int128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterDBGMODPOS(opts *bind.FilterOpts) (*IUniswapV3PoolEventsDBGMODPOSIterator, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "DBG_MOD_POS")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsDBGMODPOSIterator{contract: _IUniswapV3PoolEvents.contract, event: "DBG_MOD_POS", logs: logs, sub: sub}, nil
}

// WatchDBGMODPOS is a free log subscription operation binding the contract event 0xd1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, int128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchDBGMODPOS(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsDBGMODPOS) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "DBG_MOD_POS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsDBGMODPOS)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_MOD_POS", log); err != nil {
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

// ParseDBGMODPOS is a log parse operation binding the contract event 0xd1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, int128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseDBGMODPOS(log types.Log) (*IUniswapV3PoolEventsDBGMODPOS, error) {
	event := new(IUniswapV3PoolEventsDBGMODPOS)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_MOD_POS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsDBGSWAPHEADIterator is returned from FilterDBGSWAPHEAD and is used to iterate over the raw logs and unpacked data for DBGSWAPHEAD events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPHEADIterator struct {
	Event *IUniswapV3PoolEventsDBGSWAPHEAD // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsDBGSWAPHEADIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsDBGSWAPHEAD)
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
		it.Event = new(IUniswapV3PoolEventsDBGSWAPHEAD)
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
func (it *IUniswapV3PoolEventsDBGSWAPHEADIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsDBGSWAPHEADIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsDBGSWAPHEAD represents a DBGSWAPHEAD event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPHEAD struct {
	Pool              common.Address
	ZeroForOne        bool
	Amount            *big.Int
	SqrtStartPrice    *big.Int
	SqrtPriceLimit    *big.Int
	Slot0sqrtpricex96 *big.Int
	Tick              *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPHEAD is a free log retrieval operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterDBGSWAPHEAD(opts *bind.FilterOpts, pool []common.Address) (*IUniswapV3PoolEventsDBGSWAPHEADIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "DBG_SWAP_HEAD", poolRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsDBGSWAPHEADIterator{contract: _IUniswapV3PoolEvents.contract, event: "DBG_SWAP_HEAD", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPHEAD is a free log subscription operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchDBGSWAPHEAD(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsDBGSWAPHEAD, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "DBG_SWAP_HEAD", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsDBGSWAPHEAD)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_HEAD", log); err != nil {
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

// ParseDBGSWAPHEAD is a log parse operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseDBGSWAPHEAD(log types.Log) (*IUniswapV3PoolEventsDBGSWAPHEAD, error) {
	event := new(IUniswapV3PoolEventsDBGSWAPHEAD)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_HEAD", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsDBGSWAPLOOPIterator is returned from FilterDBGSWAPLOOP and is used to iterate over the raw logs and unpacked data for DBGSWAPLOOP events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPLOOPIterator struct {
	Event *IUniswapV3PoolEventsDBGSWAPLOOP // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsDBGSWAPLOOPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsDBGSWAPLOOP)
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
		it.Event = new(IUniswapV3PoolEventsDBGSWAPLOOP)
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
func (it *IUniswapV3PoolEventsDBGSWAPLOOPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsDBGSWAPLOOPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsDBGSWAPLOOP represents a DBGSWAPLOOP event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPLOOP struct {
	Pool                common.Address
	SqrtPriceX96        *big.Int
	SqrtPriceStartX96   *big.Int
	SqrtPriceNextX96    *big.Int
	Tick                *big.Int
	TickCumulative      *big.Int
	Initialized         bool
	StepAmountIn        *big.Int
	StepAmountOut       *big.Int
	AmountProcessed     *big.Int
	FeeAmount           *big.Int
	Liquidity           *big.Int
	ExactInput          bool
	FeeGrowthGlobalX128 *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPLOOP is a free log retrieval operation binding the contract event 0xf5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d5929.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int56 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, int256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterDBGSWAPLOOP(opts *bind.FilterOpts, pool []common.Address) (*IUniswapV3PoolEventsDBGSWAPLOOPIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "DBG_SWAP_LOOP", poolRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsDBGSWAPLOOPIterator{contract: _IUniswapV3PoolEvents.contract, event: "DBG_SWAP_LOOP", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPLOOP is a free log subscription operation binding the contract event 0xf5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d5929.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int56 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, int256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchDBGSWAPLOOP(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsDBGSWAPLOOP, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "DBG_SWAP_LOOP", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsDBGSWAPLOOP)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_LOOP", log); err != nil {
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

// ParseDBGSWAPLOOP is a log parse operation binding the contract event 0xf5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d5929.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int56 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, int256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseDBGSWAPLOOP(log types.Log) (*IUniswapV3PoolEventsDBGSWAPLOOP, error) {
	event := new(IUniswapV3PoolEventsDBGSWAPLOOP)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_LOOP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsDBGSWAPTAILIterator is returned from FilterDBGSWAPTAIL and is used to iterate over the raw logs and unpacked data for DBGSWAPTAIL events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPTAILIterator struct {
	Event *IUniswapV3PoolEventsDBGSWAPTAIL // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsDBGSWAPTAILIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsDBGSWAPTAIL)
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
		it.Event = new(IUniswapV3PoolEventsDBGSWAPTAIL)
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
func (it *IUniswapV3PoolEventsDBGSWAPTAILIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsDBGSWAPTAILIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsDBGSWAPTAIL represents a DBGSWAPTAIL event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPTAIL struct {
	Pool                 common.Address
	FeeGrowthGlobal0X128 *big.Int
	FeeGrowthGlobal1X128 *big.Int
	Tick                 *big.Int
	SqrtPriceX96         *big.Int
	Liquidity            *big.Int
	LiquidityDiff        *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPTAIL is a free log retrieval operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterDBGSWAPTAIL(opts *bind.FilterOpts, pool []common.Address) (*IUniswapV3PoolEventsDBGSWAPTAILIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "DBG_SWAP_TAIL", poolRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsDBGSWAPTAILIterator{contract: _IUniswapV3PoolEvents.contract, event: "DBG_SWAP_TAIL", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPTAIL is a free log subscription operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchDBGSWAPTAIL(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsDBGSWAPTAIL, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "DBG_SWAP_TAIL", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsDBGSWAPTAIL)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_TAIL", log); err != nil {
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

// ParseDBGSWAPTAIL is a log parse operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseDBGSWAPTAIL(log types.Log) (*IUniswapV3PoolEventsDBGSWAPTAIL, error) {
	event := new(IUniswapV3PoolEventsDBGSWAPTAIL)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_TAIL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsDBGUPDPOSIterator is returned from FilterDBGUPDPOS and is used to iterate over the raw logs and unpacked data for DBGUPDPOS events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGUPDPOSIterator struct {
	Event *IUniswapV3PoolEventsDBGUPDPOS // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsDBGUPDPOSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsDBGUPDPOS)
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
		it.Event = new(IUniswapV3PoolEventsDBGUPDPOS)
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
func (it *IUniswapV3PoolEventsDBGUPDPOSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsDBGUPDPOSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsDBGUPDPOS represents a DBGUPDPOS event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGUPDPOS struct {
	Owner                      common.Address
	TickLower                  *big.Int
	TickUpper                  *big.Int
	Tick                       *big.Int
	LiquidityDelta             *big.Int
	FeeGrowthGlobal0X128Before *big.Int
	FeeGrowthGlobal1X128Before *big.Int
	FeeGrowthInside0X128       *big.Int
	FeeGrowthInside1X128       *big.Int
	FlippedLower               bool
	FlippedUpper               bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterDBGUPDPOS is a free log retrieval operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterDBGUPDPOS(opts *bind.FilterOpts) (*IUniswapV3PoolEventsDBGUPDPOSIterator, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "DBG_UPD_POS")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsDBGUPDPOSIterator{contract: _IUniswapV3PoolEvents.contract, event: "DBG_UPD_POS", logs: logs, sub: sub}, nil
}

// WatchDBGUPDPOS is a free log subscription operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchDBGUPDPOS(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsDBGUPDPOS) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "DBG_UPD_POS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsDBGUPDPOS)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_UPD_POS", log); err != nil {
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

// ParseDBGUPDPOS is a log parse operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseDBGUPDPOS(log types.Log) (*IUniswapV3PoolEventsDBGUPDPOS, error) {
	event := new(IUniswapV3PoolEventsDBGUPDPOS)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_UPD_POS", log); err != nil {
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

// LiquidityMathMetaData contains all meta data concerning the LiquidityMath contract.
var LiquidityMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bc0017534235b35ef4af663bd5444a112b945924ded5d4ead0636940a9fbf7ce64736f6c63430007060033",
}

// LiquidityMathABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidityMathMetaData.ABI instead.
var LiquidityMathABI = LiquidityMathMetaData.ABI

// LiquidityMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LiquidityMathMetaData.Bin instead.
var LiquidityMathBin = LiquidityMathMetaData.Bin

// DeployLiquidityMath deploys a new Ethereum contract, binding an instance of LiquidityMath to it.
func DeployLiquidityMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LiquidityMath, error) {
	parsed, err := LiquidityMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LiquidityMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LiquidityMath{LiquidityMathCaller: LiquidityMathCaller{contract: contract}, LiquidityMathTransactor: LiquidityMathTransactor{contract: contract}, LiquidityMathFilterer: LiquidityMathFilterer{contract: contract}}, nil
}

// LiquidityMath is an auto generated Go binding around an Ethereum contract.
type LiquidityMath struct {
	LiquidityMathCaller     // Read-only binding to the contract
	LiquidityMathTransactor // Write-only binding to the contract
	LiquidityMathFilterer   // Log filterer for contract events
}

// LiquidityMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidityMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidityMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidityMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidityMathSession struct {
	Contract     *LiquidityMath    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidityMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidityMathCallerSession struct {
	Contract *LiquidityMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LiquidityMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidityMathTransactorSession struct {
	Contract     *LiquidityMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LiquidityMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidityMathRaw struct {
	Contract *LiquidityMath // Generic contract binding to access the raw methods on
}

// LiquidityMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidityMathCallerRaw struct {
	Contract *LiquidityMathCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidityMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidityMathTransactorRaw struct {
	Contract *LiquidityMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidityMath creates a new instance of LiquidityMath, bound to a specific deployed contract.
func NewLiquidityMath(address common.Address, backend bind.ContractBackend) (*LiquidityMath, error) {
	contract, err := bindLiquidityMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidityMath{LiquidityMathCaller: LiquidityMathCaller{contract: contract}, LiquidityMathTransactor: LiquidityMathTransactor{contract: contract}, LiquidityMathFilterer: LiquidityMathFilterer{contract: contract}}, nil
}

// NewLiquidityMathCaller creates a new read-only instance of LiquidityMath, bound to a specific deployed contract.
func NewLiquidityMathCaller(address common.Address, caller bind.ContractCaller) (*LiquidityMathCaller, error) {
	contract, err := bindLiquidityMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityMathCaller{contract: contract}, nil
}

// NewLiquidityMathTransactor creates a new write-only instance of LiquidityMath, bound to a specific deployed contract.
func NewLiquidityMathTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidityMathTransactor, error) {
	contract, err := bindLiquidityMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityMathTransactor{contract: contract}, nil
}

// NewLiquidityMathFilterer creates a new log filterer instance of LiquidityMath, bound to a specific deployed contract.
func NewLiquidityMathFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidityMathFilterer, error) {
	contract, err := bindLiquidityMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidityMathFilterer{contract: contract}, nil
}

// bindLiquidityMath binds a generic wrapper to an already deployed contract.
func bindLiquidityMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidityMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityMath *LiquidityMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityMath.Contract.LiquidityMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityMath *LiquidityMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityMath.Contract.LiquidityMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityMath *LiquidityMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityMath.Contract.LiquidityMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityMath *LiquidityMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityMath *LiquidityMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityMath *LiquidityMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityMath.Contract.contract.Transact(opts, method, params...)
}

// LowGasSafeMathMetaData contains all meta data concerning the LowGasSafeMath contract.
var LowGasSafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220722f8fa6027baf175df53fc97cc41465739fa890307b44506c270415ee18807964736f6c63430007060033",
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

// NoDelegateCallMetaData contains all meta data concerning the NoDelegateCall contract.
var NoDelegateCallMetaData = &bind.MetaData{
	ABI: "[]",
}

// NoDelegateCallABI is the input ABI used to generate the binding from.
// Deprecated: Use NoDelegateCallMetaData.ABI instead.
var NoDelegateCallABI = NoDelegateCallMetaData.ABI

// NoDelegateCall is an auto generated Go binding around an Ethereum contract.
type NoDelegateCall struct {
	NoDelegateCallCaller     // Read-only binding to the contract
	NoDelegateCallTransactor // Write-only binding to the contract
	NoDelegateCallFilterer   // Log filterer for contract events
}

// NoDelegateCallCaller is an auto generated read-only Go binding around an Ethereum contract.
type NoDelegateCallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoDelegateCallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NoDelegateCallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoDelegateCallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NoDelegateCallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoDelegateCallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NoDelegateCallSession struct {
	Contract     *NoDelegateCall   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NoDelegateCallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NoDelegateCallCallerSession struct {
	Contract *NoDelegateCallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NoDelegateCallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NoDelegateCallTransactorSession struct {
	Contract     *NoDelegateCallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NoDelegateCallRaw is an auto generated low-level Go binding around an Ethereum contract.
type NoDelegateCallRaw struct {
	Contract *NoDelegateCall // Generic contract binding to access the raw methods on
}

// NoDelegateCallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NoDelegateCallCallerRaw struct {
	Contract *NoDelegateCallCaller // Generic read-only contract binding to access the raw methods on
}

// NoDelegateCallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NoDelegateCallTransactorRaw struct {
	Contract *NoDelegateCallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNoDelegateCall creates a new instance of NoDelegateCall, bound to a specific deployed contract.
func NewNoDelegateCall(address common.Address, backend bind.ContractBackend) (*NoDelegateCall, error) {
	contract, err := bindNoDelegateCall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoDelegateCall{NoDelegateCallCaller: NoDelegateCallCaller{contract: contract}, NoDelegateCallTransactor: NoDelegateCallTransactor{contract: contract}, NoDelegateCallFilterer: NoDelegateCallFilterer{contract: contract}}, nil
}

// NewNoDelegateCallCaller creates a new read-only instance of NoDelegateCall, bound to a specific deployed contract.
func NewNoDelegateCallCaller(address common.Address, caller bind.ContractCaller) (*NoDelegateCallCaller, error) {
	contract, err := bindNoDelegateCall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoDelegateCallCaller{contract: contract}, nil
}

// NewNoDelegateCallTransactor creates a new write-only instance of NoDelegateCall, bound to a specific deployed contract.
func NewNoDelegateCallTransactor(address common.Address, transactor bind.ContractTransactor) (*NoDelegateCallTransactor, error) {
	contract, err := bindNoDelegateCall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoDelegateCallTransactor{contract: contract}, nil
}

// NewNoDelegateCallFilterer creates a new log filterer instance of NoDelegateCall, bound to a specific deployed contract.
func NewNoDelegateCallFilterer(address common.Address, filterer bind.ContractFilterer) (*NoDelegateCallFilterer, error) {
	contract, err := bindNoDelegateCall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoDelegateCallFilterer{contract: contract}, nil
}

// bindNoDelegateCall binds a generic wrapper to an already deployed contract.
func bindNoDelegateCall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NoDelegateCallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoDelegateCall *NoDelegateCallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoDelegateCall.Contract.NoDelegateCallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoDelegateCall *NoDelegateCallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoDelegateCall.Contract.NoDelegateCallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoDelegateCall *NoDelegateCallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoDelegateCall.Contract.NoDelegateCallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoDelegateCall *NoDelegateCallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoDelegateCall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoDelegateCall *NoDelegateCallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoDelegateCall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoDelegateCall *NoDelegateCallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoDelegateCall.Contract.contract.Transact(opts, method, params...)
}

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201f97fc2c542f19b7b12d7d6cf6cdd9b349f88d80e86e112c9448877e44f5fcc864736f6c63430007060033",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// OracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleMetaData.Bin instead.
var OracleBin = OracleMetaData.Bin

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// PositionMetaData contains all meta data concerning the Position contract.
var PositionMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220be8b66dcf5cf0631afe97c25140d857e6b7b4f3f2f71652d14f583993e9160c264736f6c63430007060033",
}

// PositionABI is the input ABI used to generate the binding from.
// Deprecated: Use PositionMetaData.ABI instead.
var PositionABI = PositionMetaData.ABI

// PositionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PositionMetaData.Bin instead.
var PositionBin = PositionMetaData.Bin

// DeployPosition deploys a new Ethereum contract, binding an instance of Position to it.
func DeployPosition(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Position, error) {
	parsed, err := PositionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PositionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Position{PositionCaller: PositionCaller{contract: contract}, PositionTransactor: PositionTransactor{contract: contract}, PositionFilterer: PositionFilterer{contract: contract}}, nil
}

// Position is an auto generated Go binding around an Ethereum contract.
type Position struct {
	PositionCaller     // Read-only binding to the contract
	PositionTransactor // Write-only binding to the contract
	PositionFilterer   // Log filterer for contract events
}

// PositionCaller is an auto generated read-only Go binding around an Ethereum contract.
type PositionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PositionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PositionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PositionSession struct {
	Contract     *Position         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PositionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PositionCallerSession struct {
	Contract *PositionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PositionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PositionTransactorSession struct {
	Contract     *PositionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PositionRaw is an auto generated low-level Go binding around an Ethereum contract.
type PositionRaw struct {
	Contract *Position // Generic contract binding to access the raw methods on
}

// PositionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PositionCallerRaw struct {
	Contract *PositionCaller // Generic read-only contract binding to access the raw methods on
}

// PositionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PositionTransactorRaw struct {
	Contract *PositionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPosition creates a new instance of Position, bound to a specific deployed contract.
func NewPosition(address common.Address, backend bind.ContractBackend) (*Position, error) {
	contract, err := bindPosition(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Position{PositionCaller: PositionCaller{contract: contract}, PositionTransactor: PositionTransactor{contract: contract}, PositionFilterer: PositionFilterer{contract: contract}}, nil
}

// NewPositionCaller creates a new read-only instance of Position, bound to a specific deployed contract.
func NewPositionCaller(address common.Address, caller bind.ContractCaller) (*PositionCaller, error) {
	contract, err := bindPosition(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PositionCaller{contract: contract}, nil
}

// NewPositionTransactor creates a new write-only instance of Position, bound to a specific deployed contract.
func NewPositionTransactor(address common.Address, transactor bind.ContractTransactor) (*PositionTransactor, error) {
	contract, err := bindPosition(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PositionTransactor{contract: contract}, nil
}

// NewPositionFilterer creates a new log filterer instance of Position, bound to a specific deployed contract.
func NewPositionFilterer(address common.Address, filterer bind.ContractFilterer) (*PositionFilterer, error) {
	contract, err := bindPosition(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PositionFilterer{contract: contract}, nil
}

// bindPosition binds a generic wrapper to an already deployed contract.
func bindPosition(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PositionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Position *PositionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Position.Contract.PositionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Position *PositionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Position.Contract.PositionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Position *PositionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Position.Contract.PositionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Position *PositionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Position.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Position *PositionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Position.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Position *PositionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Position.Contract.contract.Transact(opts, method, params...)
}

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d065aefaad9b10271dd665f9f1ff0a887f07952f0bf91f46fa179b93364b4ea864736f6c63430007060033",
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

// SqrtPriceMathMetaData contains all meta data concerning the SqrtPriceMath contract.
var SqrtPriceMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122046ad3fc1b3e39b8c1667196da4133710b8d65aa0b6a046d9115312c2ce978f4764736f6c63430007060033",
}

// SqrtPriceMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SqrtPriceMathMetaData.ABI instead.
var SqrtPriceMathABI = SqrtPriceMathMetaData.ABI

// SqrtPriceMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SqrtPriceMathMetaData.Bin instead.
var SqrtPriceMathBin = SqrtPriceMathMetaData.Bin

// DeploySqrtPriceMath deploys a new Ethereum contract, binding an instance of SqrtPriceMath to it.
func DeploySqrtPriceMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SqrtPriceMath, error) {
	parsed, err := SqrtPriceMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SqrtPriceMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SqrtPriceMath{SqrtPriceMathCaller: SqrtPriceMathCaller{contract: contract}, SqrtPriceMathTransactor: SqrtPriceMathTransactor{contract: contract}, SqrtPriceMathFilterer: SqrtPriceMathFilterer{contract: contract}}, nil
}

// SqrtPriceMath is an auto generated Go binding around an Ethereum contract.
type SqrtPriceMath struct {
	SqrtPriceMathCaller     // Read-only binding to the contract
	SqrtPriceMathTransactor // Write-only binding to the contract
	SqrtPriceMathFilterer   // Log filterer for contract events
}

// SqrtPriceMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SqrtPriceMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SqrtPriceMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SqrtPriceMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SqrtPriceMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SqrtPriceMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SqrtPriceMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SqrtPriceMathSession struct {
	Contract     *SqrtPriceMath    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SqrtPriceMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SqrtPriceMathCallerSession struct {
	Contract *SqrtPriceMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SqrtPriceMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SqrtPriceMathTransactorSession struct {
	Contract     *SqrtPriceMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SqrtPriceMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SqrtPriceMathRaw struct {
	Contract *SqrtPriceMath // Generic contract binding to access the raw methods on
}

// SqrtPriceMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SqrtPriceMathCallerRaw struct {
	Contract *SqrtPriceMathCaller // Generic read-only contract binding to access the raw methods on
}

// SqrtPriceMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SqrtPriceMathTransactorRaw struct {
	Contract *SqrtPriceMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSqrtPriceMath creates a new instance of SqrtPriceMath, bound to a specific deployed contract.
func NewSqrtPriceMath(address common.Address, backend bind.ContractBackend) (*SqrtPriceMath, error) {
	contract, err := bindSqrtPriceMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SqrtPriceMath{SqrtPriceMathCaller: SqrtPriceMathCaller{contract: contract}, SqrtPriceMathTransactor: SqrtPriceMathTransactor{contract: contract}, SqrtPriceMathFilterer: SqrtPriceMathFilterer{contract: contract}}, nil
}

// NewSqrtPriceMathCaller creates a new read-only instance of SqrtPriceMath, bound to a specific deployed contract.
func NewSqrtPriceMathCaller(address common.Address, caller bind.ContractCaller) (*SqrtPriceMathCaller, error) {
	contract, err := bindSqrtPriceMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SqrtPriceMathCaller{contract: contract}, nil
}

// NewSqrtPriceMathTransactor creates a new write-only instance of SqrtPriceMath, bound to a specific deployed contract.
func NewSqrtPriceMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SqrtPriceMathTransactor, error) {
	contract, err := bindSqrtPriceMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SqrtPriceMathTransactor{contract: contract}, nil
}

// NewSqrtPriceMathFilterer creates a new log filterer instance of SqrtPriceMath, bound to a specific deployed contract.
func NewSqrtPriceMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SqrtPriceMathFilterer, error) {
	contract, err := bindSqrtPriceMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SqrtPriceMathFilterer{contract: contract}, nil
}

// bindSqrtPriceMath binds a generic wrapper to an already deployed contract.
func bindSqrtPriceMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SqrtPriceMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SqrtPriceMath *SqrtPriceMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SqrtPriceMath.Contract.SqrtPriceMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SqrtPriceMath *SqrtPriceMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SqrtPriceMath.Contract.SqrtPriceMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SqrtPriceMath *SqrtPriceMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SqrtPriceMath.Contract.SqrtPriceMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SqrtPriceMath *SqrtPriceMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SqrtPriceMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SqrtPriceMath *SqrtPriceMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SqrtPriceMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SqrtPriceMath *SqrtPriceMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SqrtPriceMath.Contract.contract.Transact(opts, method, params...)
}

// SwapMathMetaData contains all meta data concerning the SwapMath contract.
var SwapMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122058f20818fbc09f49902f7aee5e46c6146807c61b7923de06682e8ae7af59a6ad64736f6c63430007060033",
}

// SwapMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapMathMetaData.ABI instead.
var SwapMathABI = SwapMathMetaData.ABI

// SwapMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapMathMetaData.Bin instead.
var SwapMathBin = SwapMathMetaData.Bin

// DeploySwapMath deploys a new Ethereum contract, binding an instance of SwapMath to it.
func DeploySwapMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SwapMath, error) {
	parsed, err := SwapMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwapMath{SwapMathCaller: SwapMathCaller{contract: contract}, SwapMathTransactor: SwapMathTransactor{contract: contract}, SwapMathFilterer: SwapMathFilterer{contract: contract}}, nil
}

// SwapMath is an auto generated Go binding around an Ethereum contract.
type SwapMath struct {
	SwapMathCaller     // Read-only binding to the contract
	SwapMathTransactor // Write-only binding to the contract
	SwapMathFilterer   // Log filterer for contract events
}

// SwapMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapMathSession struct {
	Contract     *SwapMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapMathCallerSession struct {
	Contract *SwapMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SwapMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapMathTransactorSession struct {
	Contract     *SwapMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SwapMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapMathRaw struct {
	Contract *SwapMath // Generic contract binding to access the raw methods on
}

// SwapMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapMathCallerRaw struct {
	Contract *SwapMathCaller // Generic read-only contract binding to access the raw methods on
}

// SwapMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapMathTransactorRaw struct {
	Contract *SwapMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapMath creates a new instance of SwapMath, bound to a specific deployed contract.
func NewSwapMath(address common.Address, backend bind.ContractBackend) (*SwapMath, error) {
	contract, err := bindSwapMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapMath{SwapMathCaller: SwapMathCaller{contract: contract}, SwapMathTransactor: SwapMathTransactor{contract: contract}, SwapMathFilterer: SwapMathFilterer{contract: contract}}, nil
}

// NewSwapMathCaller creates a new read-only instance of SwapMath, bound to a specific deployed contract.
func NewSwapMathCaller(address common.Address, caller bind.ContractCaller) (*SwapMathCaller, error) {
	contract, err := bindSwapMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapMathCaller{contract: contract}, nil
}

// NewSwapMathTransactor creates a new write-only instance of SwapMath, bound to a specific deployed contract.
func NewSwapMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapMathTransactor, error) {
	contract, err := bindSwapMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapMathTransactor{contract: contract}, nil
}

// NewSwapMathFilterer creates a new log filterer instance of SwapMath, bound to a specific deployed contract.
func NewSwapMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapMathFilterer, error) {
	contract, err := bindSwapMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapMathFilterer{contract: contract}, nil
}

// bindSwapMath binds a generic wrapper to an already deployed contract.
func bindSwapMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapMath *SwapMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapMath.Contract.SwapMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapMath *SwapMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapMath.Contract.SwapMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapMath *SwapMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapMath.Contract.SwapMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapMath *SwapMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapMath *SwapMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapMath *SwapMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapMath.Contract.contract.Transact(opts, method, params...)
}

// TickMetaData contains all meta data concerning the Tick contract.
var TickMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122098f2f7abb8768e7d59a385620fa5e3ee31a17905c074b807a7d6f94eb848375464736f6c63430007060033",
}

// TickABI is the input ABI used to generate the binding from.
// Deprecated: Use TickMetaData.ABI instead.
var TickABI = TickMetaData.ABI

// TickBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TickMetaData.Bin instead.
var TickBin = TickMetaData.Bin

// DeployTick deploys a new Ethereum contract, binding an instance of Tick to it.
func DeployTick(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tick, error) {
	parsed, err := TickMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TickBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tick{TickCaller: TickCaller{contract: contract}, TickTransactor: TickTransactor{contract: contract}, TickFilterer: TickFilterer{contract: contract}}, nil
}

// Tick is an auto generated Go binding around an Ethereum contract.
type Tick struct {
	TickCaller     // Read-only binding to the contract
	TickTransactor // Write-only binding to the contract
	TickFilterer   // Log filterer for contract events
}

// TickCaller is an auto generated read-only Go binding around an Ethereum contract.
type TickCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TickTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TickFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TickSession struct {
	Contract     *Tick             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TickCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TickCallerSession struct {
	Contract *TickCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TickTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TickTransactorSession struct {
	Contract     *TickTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TickRaw is an auto generated low-level Go binding around an Ethereum contract.
type TickRaw struct {
	Contract *Tick // Generic contract binding to access the raw methods on
}

// TickCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TickCallerRaw struct {
	Contract *TickCaller // Generic read-only contract binding to access the raw methods on
}

// TickTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TickTransactorRaw struct {
	Contract *TickTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTick creates a new instance of Tick, bound to a specific deployed contract.
func NewTick(address common.Address, backend bind.ContractBackend) (*Tick, error) {
	contract, err := bindTick(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tick{TickCaller: TickCaller{contract: contract}, TickTransactor: TickTransactor{contract: contract}, TickFilterer: TickFilterer{contract: contract}}, nil
}

// NewTickCaller creates a new read-only instance of Tick, bound to a specific deployed contract.
func NewTickCaller(address common.Address, caller bind.ContractCaller) (*TickCaller, error) {
	contract, err := bindTick(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TickCaller{contract: contract}, nil
}

// NewTickTransactor creates a new write-only instance of Tick, bound to a specific deployed contract.
func NewTickTransactor(address common.Address, transactor bind.ContractTransactor) (*TickTransactor, error) {
	contract, err := bindTick(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TickTransactor{contract: contract}, nil
}

// NewTickFilterer creates a new log filterer instance of Tick, bound to a specific deployed contract.
func NewTickFilterer(address common.Address, filterer bind.ContractFilterer) (*TickFilterer, error) {
	contract, err := bindTick(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TickFilterer{contract: contract}, nil
}

// bindTick binds a generic wrapper to an already deployed contract.
func bindTick(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TickABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tick *TickRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tick.Contract.TickCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tick *TickRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tick.Contract.TickTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tick *TickRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tick.Contract.TickTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tick *TickCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tick.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tick *TickTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tick.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tick *TickTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tick.Contract.contract.Transact(opts, method, params...)
}

// TickBitmapMetaData contains all meta data concerning the TickBitmap contract.
var TickBitmapMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d8cd59c1cbdaccce3de0b833e4073414bed8d07de698d53efaa6a2703f5202e664736f6c63430007060033",
}

// TickBitmapABI is the input ABI used to generate the binding from.
// Deprecated: Use TickBitmapMetaData.ABI instead.
var TickBitmapABI = TickBitmapMetaData.ABI

// TickBitmapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TickBitmapMetaData.Bin instead.
var TickBitmapBin = TickBitmapMetaData.Bin

// DeployTickBitmap deploys a new Ethereum contract, binding an instance of TickBitmap to it.
func DeployTickBitmap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TickBitmap, error) {
	parsed, err := TickBitmapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TickBitmapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TickBitmap{TickBitmapCaller: TickBitmapCaller{contract: contract}, TickBitmapTransactor: TickBitmapTransactor{contract: contract}, TickBitmapFilterer: TickBitmapFilterer{contract: contract}}, nil
}

// TickBitmap is an auto generated Go binding around an Ethereum contract.
type TickBitmap struct {
	TickBitmapCaller     // Read-only binding to the contract
	TickBitmapTransactor // Write-only binding to the contract
	TickBitmapFilterer   // Log filterer for contract events
}

// TickBitmapCaller is an auto generated read-only Go binding around an Ethereum contract.
type TickBitmapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickBitmapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TickBitmapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickBitmapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TickBitmapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickBitmapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TickBitmapSession struct {
	Contract     *TickBitmap       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TickBitmapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TickBitmapCallerSession struct {
	Contract *TickBitmapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TickBitmapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TickBitmapTransactorSession struct {
	Contract     *TickBitmapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TickBitmapRaw is an auto generated low-level Go binding around an Ethereum contract.
type TickBitmapRaw struct {
	Contract *TickBitmap // Generic contract binding to access the raw methods on
}

// TickBitmapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TickBitmapCallerRaw struct {
	Contract *TickBitmapCaller // Generic read-only contract binding to access the raw methods on
}

// TickBitmapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TickBitmapTransactorRaw struct {
	Contract *TickBitmapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTickBitmap creates a new instance of TickBitmap, bound to a specific deployed contract.
func NewTickBitmap(address common.Address, backend bind.ContractBackend) (*TickBitmap, error) {
	contract, err := bindTickBitmap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TickBitmap{TickBitmapCaller: TickBitmapCaller{contract: contract}, TickBitmapTransactor: TickBitmapTransactor{contract: contract}, TickBitmapFilterer: TickBitmapFilterer{contract: contract}}, nil
}

// NewTickBitmapCaller creates a new read-only instance of TickBitmap, bound to a specific deployed contract.
func NewTickBitmapCaller(address common.Address, caller bind.ContractCaller) (*TickBitmapCaller, error) {
	contract, err := bindTickBitmap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TickBitmapCaller{contract: contract}, nil
}

// NewTickBitmapTransactor creates a new write-only instance of TickBitmap, bound to a specific deployed contract.
func NewTickBitmapTransactor(address common.Address, transactor bind.ContractTransactor) (*TickBitmapTransactor, error) {
	contract, err := bindTickBitmap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TickBitmapTransactor{contract: contract}, nil
}

// NewTickBitmapFilterer creates a new log filterer instance of TickBitmap, bound to a specific deployed contract.
func NewTickBitmapFilterer(address common.Address, filterer bind.ContractFilterer) (*TickBitmapFilterer, error) {
	contract, err := bindTickBitmap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TickBitmapFilterer{contract: contract}, nil
}

// bindTickBitmap binds a generic wrapper to an already deployed contract.
func bindTickBitmap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TickBitmapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TickBitmap *TickBitmapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TickBitmap.Contract.TickBitmapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TickBitmap *TickBitmapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TickBitmap.Contract.TickBitmapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TickBitmap *TickBitmapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TickBitmap.Contract.TickBitmapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TickBitmap *TickBitmapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TickBitmap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TickBitmap *TickBitmapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TickBitmap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TickBitmap *TickBitmapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TickBitmap.Contract.contract.Transact(opts, method, params...)
}

// TickMathMetaData contains all meta data concerning the TickMath contract.
var TickMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208d867c6e9269d2e81beaff09569922964c3f6bc4b3c1dc8c4944deb8774f1e9a64736f6c63430007060033",
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
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209d13388a67fff76fc4cff4a3304bbcba63839229c45c0b320d7a0e5d7e286ea764736f6c63430007060033",
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

// UniswapV3FactoryMetaData contains all meta data concerning the UniswapV3Factory contract.
var UniswapV3FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"FeeAmountEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"enableFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b503060601b608052600380546001600160a01b031916339081179091556040516000907fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c908290a36101f4600081815260046020527ffb8cf1d12598d1a039dd1d106665851a96aadf67d0d9ed76fceea282119208b7805462ffffff1916600a9081179091556040519092916000805160206171c783398151915291a3610bb8600081815260046020527f72dffa9b822156d9cf4b0090fa0b656bcb9cc2b2c60eb6acfc20a34f54b31743805462ffffff1916603c9081179091556040519092916000805160206171c783398151915291a3612710600081815260046020527f8cc740d51daa94ff54f33bd779c2d20149f524c340519b49181be5a08615f829805462ffffff191660c89081179091556040519092916000805160206171c783398151915291a360805160601c6170536101746000398061070352506170536000f3fe608060405234801561001057600080fd5b506004361061006d5760003560e01c806313af4035146100725780631698ee821461009a57806322afcccb146100f2578063890357301461012b5780638a7c195f146101755780638da5cb5b146101a0578063a1671295146101a8575b600080fd5b6100986004803603602081101561008857600080fd5b50356001600160a01b03166101e4565b005b6100d6600480360360608110156100b057600080fd5b5080356001600160a01b03908116916020810135909116906040013562ffffff16610289565b604080516001600160a01b039092168252519081900360200190f35b6101146004803603602081101561010857600080fd5b503562ffffff166102b5565b6040805160029290920b8252519081900360200190f35b6101336102ca565b604080516001600160a01b0396871681529486166020860152929094168383015262ffffff16606083015260029290920b608082015290519081900360a00190f35b6100986004803603604081101561018b57600080fd5b5062ffffff813516906020013560020b610300565b6100d661048b565b6100d6600480360360608110156101be57600080fd5b5080356001600160a01b03908116916020810135909116906040013562ffffff1661049a565b6003546001600160a01b0316331461022d5760405162461bcd60e51b815260040180806020018281038252602d815260200180616fc5602d913960400191505060405180910390fd5b6003546040516001600160a01b038084169216907fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c90600090a3600380546001600160a01b0319166001600160a01b0392909216919091179055565b60056020908152600093845260408085208252928452828420905282529020546001600160a01b031681565b60046020526000908152604090205460020b81565b600054600154600280546001600160a01b03938416939283169281169162ffffff600160a01b83041691600160b81b9004900b85565b6003546001600160a01b031633146103495760405162461bcd60e51b8152600401808060200182810382526035815260200180616e236035913960400191505060405180910390fd5b620f42408262ffffff161061038f5760405162461bcd60e51b815260040180806020018281038252602f815260200180616df4602f913960400191505060405180910390fd5b60008160020b1380156103a657506140008160020b125b6103e15760405162461bcd60e51b8152600401808060200182810382526045815260200180616f536045913960600191505060405180910390fd5b62ffffff8216600090815260046020526040902054600290810b900b156104395760405162461bcd60e51b8152600401808060200182810382526040815260200180616e586040913960400191505060405180910390fd5b62ffffff828116600081815260046020526040808220805462ffffff1916600287900b958616179055517fc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc9190a35050565b6003546001600160a01b031681565b60006104a46106f8565b826001600160a01b0316846001600160a01b031614156104f55760405162461bcd60e51b815260040180806020018281038252602c815260200180616ff2602c913960400191505060405180910390fd5b600080846001600160a01b0316866001600160a01b03161061051857848661051b565b85855b90925090506001600160a01b0382166105655760405162461bcd60e51b8152600401808060200182810382526030815260200180616ee66030913960400191505060405180910390fd5b62ffffff8416600090815260046020526040902054600290810b9081900b6105be5760405162461bcd60e51b815260040180806020018281038252602d815260200180616f98602d913960400191505060405180910390fd5b6001600160a01b0383811660009081526005602090815260408083208685168452825280832062ffffff8a168452909152902054161561062f5760405162461bcd60e51b815260040180806020018281038252604e815260200180616e98604e913960600191505060405180910390fd5b61063c3084848885610761565b6001600160a01b03808516600081815260056020818152604080842089871680865290835281852062ffffff8e168087529084528286208054988a166001600160a01b0319998a1681179091558287529484528286208787528452828620818752845294829020805490971684179096558051600289900b815291820192909252815195995091947f783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b71189281900390910190a45050509392505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461075f5760405162461bcd60e51b815260040180806020018281038252603d815260200180616f16603d913960400191505060405180910390fd5b565b6040805160a0810182526001600160a01b03878116808352878216602080850182905292881684860181905262ffffff888116606080880182905260028a810b6080998a01819052600080546001600160a01b03199081169099178155600180548a16891790558254909816861762ffffff60a01b1916600160a01b85021762ffffff60b81b1916600160b81b91830b90951602939093179092558751808701949094528388019290925282810191909152855180830390910181529301938490528251929091019190912090916108389061088f565b8190604051809103906000f5905080158015610858573d6000803e3d6000fd5b50600080546001600160a01b0319908116909155600180549091169055600280546001600160d01b03191690559695505050505050565b6165578061089d8339019056fe6101606040523480156200001257600080fd5b503060601b60805260408051630890357360e41b81529051600091339163890357309160048082019260a092909190829003018186803b1580156200005657600080fd5b505afa1580156200006b573d6000803e3d6000fd5b505050506040513d60a08110156200008257600080fd5b508051602080830151604084015160608086015160809096015160e896871b6001600160e81b0319166101005291811b6001600160601b031990811660e05292811b831660c0529390931b1660a052600282810b900b90921b610120529150620000f79082906200010f811b62002e1217901c565b60801b6001600160801b03191661014052506200017d565b60008082600281900b620d89e719816200012557fe5b05029050600083600281900b620d89e8816200013d57fe5b0502905060008460020b83830360020b816200015557fe5b0560010190508062ffffff166001600160801b038016816200017357fe5b0495945050505050565b60805160601c60a05160601c60c05160601c60e05160601c6101005160e81c6101205160e81c6101405160801c61630d6200024a600039806121e852806152195280615250525080610c305280612bce528061528452806152b6525080610d1f5280611bc55280611bfc5280612c165250806112ed5280611c7f52806120ee52806126b95280612bf2528061423452508061085a528061141b5280611c4e52806120885280612633528061407152508061227552806124c05280612baa525080612e89525061630d6000f3fe608060405234801561001057600080fd5b506004361061013e5760003560e01c80630dfe168114610143578063128acb08146101675780631a686502146102145780631ad8b03b14610238578063252c09d71461026f57806332148f67146102c65780633850c7bd146102e95780633c8a7d8d1461034257806346141319146103e2578063490e6cbc146103fc5780634f1eb3d814610486578063514ea4bf146104d75780635339c2961461053057806370cf754a146105505780638206a4d11461055857806385b6672914610580578063883bdbfd146105bd578063a34123a7146106c4578063a38807f2146106fe578063c45a015514610759578063d0c93a7c14610761578063d21220a714610780578063ddca3f4314610788578063f3058399146107a8578063f30dba93146107b0578063f637731d14610832575b600080fd5b61014b610858565b604080516001600160a01b039092168252519081900360200190f35b6101fb600480360360a081101561017d57600080fd5b6001600160a01b0382358116926020810135151592604082013592606083013516919081019060a081016080820135600160201b8111156101bd57600080fd5b8201836020820111156101cf57600080fd5b803590602001918460018302840111600160201b831117156101f057600080fd5b50909250905061087c565b6040805192835260208301919091528051918290030190f35b61021c611660565b604080516001600160801b039092168252519081900360200190f35b61024061166f565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b61028c6004803603602081101561028557600080fd5b5035611689565b6040805163ffffffff909516855260069390930b60208501526001600160a01b039091168383015215156060830152519081900360800190f35b6102e7600480360360208110156102dc57600080fd5b503561ffff166116ce565b005b6102f16117c8565b604080516001600160a01b03909816885260029690960b602088015261ffff9485168787015292841660608701529216608085015260ff90911660a0840152151560c0830152519081900360e00190f35b6101fb600480360360a081101561035857600080fd5b6001600160a01b03823516916020810135600290810b92604083013590910b916001600160801b036060820135169181019060a081016080820135600160201b8111156103a457600080fd5b8201836020820111156103b657600080fd5b803590602001918460018302840111600160201b831117156103d757600080fd5b509092509050611818565b6103ea611b1c565b60408051918252519081900360200190f35b6102e76004803603608081101561041257600080fd5b6001600160a01b038235169160208101359160408201359190810190608081016060820135600160201b81111561044857600080fd5b82018360208201111561045a57600080fd5b803590602001918460018302840111600160201b8311171561047b57600080fd5b509092509050611b22565b610240600480360360a081101561049c57600080fd5b506001600160a01b03813516906020810135600290810b91604081013590910b906001600160801b0360608201358116916080013516611f7d565b6104f4600480360360208110156104ed57600080fd5b5035612197565b604080516001600160801b0396871681526020810195909552848101939093529084166060840152909216608082015290519081900360a00190f35b6103ea6004803603602081101561054657600080fd5b503560010b6121d4565b61021c6121e6565b6102e76004803603604081101561056e57600080fd5b5060ff8135811691602001351661220a565b6102406004803603606081101561059657600080fd5b506001600160a01b03813516906001600160801b0360208201358116916040013516612452565b61062b600480360360208110156105d357600080fd5b810190602081018135600160201b8111156105ed57600080fd5b8201836020820111156105ff57600080fd5b803590602001918460208302840111600160201b8311171561062057600080fd5b509092509050612751565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561066f578181015183820152602001610657565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156106ae578181015183820152602001610696565b5050505090500194505050505060405180910390f35b6101fb600480360360608110156106da57600080fd5b508035600290810b91602081013590910b90604001356001600160801b03166127de565b6107286004803603604081101561071457600080fd5b508035600290810b9160200135900b612955565b6040805160069490940b84526001600160a01b03909216602084015263ffffffff1682820152519081900360600190f35b61014b612ba8565b610769612bcc565b6040805160029290920b8252519081900360200190f35b61014b612bf0565b610790612c14565b6040805162ffffff9092168252519081900360200190f35b6103ea612c38565b6107d0600480360360208110156107c657600080fd5b503560020b612c3e565b604080516001600160801b039099168952600f9790970b602089015287870195909552606087019390935260069190910b60808601526001600160a01b031660a085015263ffffffff1660c0840152151560e083015251908190036101000190f35b6102e76004803603602081101561084857600080fd5b50356001600160a01b0316612ca8565b7f000000000000000000000000000000000000000000000000000000000000000081565b600080610887612e7e565b856108be576040805162461bcd60e51b8152602060048201526002602482015261415360f01b604482015290519081900360640190fd5b6040805160e0810182526000546001600160a01b0381168252600160a01b8104600290810b810b900b602083015261ffff600160b81b8204811693830193909352600160c81b810483166060830152600160d81b8104909216608082015260ff600160e81b8304811660a0830152600160f01b909204909116151560c08201819052610977576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b876109c25780600001516001600160a01b0316866001600160a01b03161180156109bd575073fffd8963efd1fc6a506488495d951d5263988d266001600160a01b038716105b6109f4565b80600001516001600160a01b0316866001600160a01b03161080156109f457506401000276a36001600160a01b038716115b610a2b576040805162461bcd60e51b815260206004820152600360248201526214d41360ea1b604482015290519081900360640190fd5b6000805460ff60f01b191681556040805160c08101909152808a610a5a5760048460a0015160ff16901c610a6d565b60108460a0015160ff1681610a6b57fe5b065b60ff1681526004546001600160801b03166020820152604001610a8e612ee7565b63ffffffff168152602001600060060b815260200160006001600160a01b031681526020016000151581525090506000808913905060006040518060e001604052808b81526020016000815260200185600001516001600160a01b03168152602001856020015160020b81526020018c610b0a57600254610b0e565b6001545b815260200160006001600160801b0316815260200184602001516001600160801b03168152509050306001600160a01b03167f884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c2078c8c87600001518d6000800160009054906101000a90046001600160a01b03168760600151604051808715158152602001868152602001856001600160a01b03168152602001846001600160a01b03168152602001836001600160a01b031681526020018260020b8152602001965050505050505060405180910390a25b805115801590610c055750886001600160a01b031681604001516001600160a01b031614155b156110df57610c12615dd8565b60408201516001600160a01b031681526060820151610c55906006907f00000000000000000000000000000000000000000000000000000000000000008f612eeb565b15156040830152600290810b810b60208301819052620d89e719910b1215610c8657620d89e7196020820152610ca5565b6020810151620d89e860029190910b1315610ca557620d89e860208201525b610cb2816020015161302d565b6001600160a01b031660608201526040820151610d43908d610cec578b6001600160a01b031683606001516001600160a01b031611610d06565b8b6001600160a01b031683606001516001600160a01b0316105b610d14578260600151610d16565b8b5b60c085015185517f0000000000000000000000000000000000000000000000000000000000000000613354565b60c085015260a084015260808301526001600160a01b031660408301528215610da557610d798160c00151826080015101613546565b825103825260a0810151610d9b90610d9090613546565b6020840151906135a3565b6020830152610de0565b610db28160a00151613546565b825101825260c08101516080820151610dda91610dcf9101613546565b6020840151906135f1565b60208301525b835160ff1615610e26576000846000015160ff168260c0015181610e0057fe5b60c0840180519290910491829003905260a0840180519091016001600160801b03169052505b60c08201516001600160801b031615610e6557610e598160c00151600160801b8460c001516001600160801b0316613639565b60808301805190910190525b306001600160a01b03167ff5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d592983604001518360000151846060015186606001518960600151876040015188608001518960a001518b602001518b60c00151600460009054906101000a90046001600160801b03168f8f60800151604051808e6001600160a01b031681526020018d6001600160a01b031681526020018c6001600160a01b031681526020018b60020b81526020018a60060b81526020018915158152602001888152602001878152602001868152602001858152602001846001600160801b0316815260200183151581526020018281526020019d505050505050505050505050505060405180910390a280606001516001600160a01b031682604001516001600160a01b0316141561109e57806040015115611075578360a00151610fff57610fdd846040015160008760200151886040015188602001518a606001516008613763909695949392919063ffffffff16565b6001600160a01b03166080860152600690810b900b6060850152600160a08501525b600061104b82602001518e6110165760015461101c565b84608001515b8f61102b57856080015161102f565b6002545b608089015160608a015160408b015160059594939291906138f5565b90508c15611057576000035b6110658360c00151826139af565b6001600160801b031660c0840152505b8b61108457806020015161108d565b60018160200151035b600290810b900b60608301526110d9565b80600001516001600160a01b031682604001516001600160a01b0316146110d9576110cc8260400151613a65565b600290810b900b60608301525b50610bdf565b836020015160020b816060015160020b146111ad5760008061112d86604001518660400151886020015188602001518a606001518b608001516008613d50909695949392919063ffffffff16565b604085015160608601516000805461ffff60c81b1916600160c81b61ffff958616021761ffff60b81b1916600160b81b95909416949094029290921762ffffff60a01b1916600160a01b62ffffff60029490940b9390931692909202919091176001600160a01b0319166001600160a01b03909116179055506111d29050565b6040810151600080546001600160a01b0319166001600160a01b039092169190911790555b8060c001516001600160801b031683602001516001600160801b0316146112185760c0810151600480546001600160801b0319166001600160801b039092169190911790555b8a1561126857608081015160015560a08101516001600160801b0316156112635760a0810151600380546001600160801b031981166001600160801b03918216909301169190911790555b6112ae565b608081015160025560a08101516001600160801b0316156112ae5760a0810151600380546001600160801b03808216600160801b92839004821690940116029190911790555b8115158b1515146112c757602081015181518b036112d4565b80600001518a0381602001515b90965094508a1561140d576000851215611316576113167f00000000000000000000000000000000000000000000000000000000000000008d87600003613ed5565b6000611320614023565b9050336001600160a01b031663fa461e3388888c8c6040518563ffffffff1660e01b815260040180858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b1580156113a457600080fd5b505af11580156113b8573d6000803e3d6000fd5b505050506113c4614023565b6113ce828961418e565b1115611407576040805162461bcd60e51b815260206004820152600360248201526249494160e81b604482015290519081900360640190fd5b50611537565b6000861215611444576114447f00000000000000000000000000000000000000000000000000000000000000008d88600003613ed5565b600061144e6141e6565b9050336001600160a01b031663fa461e3388888c8c6040518563ffffffff1660e01b815260040180858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b1580156114d257600080fd5b505af11580156114e6573d6000803e3d6000fd5b505050506114f26141e6565b6114fc828861418e565b1115611535576040805162461bcd60e51b815260206004820152600360248201526249494160e81b604482015290519081900360640190fd5b505b6001546002805460608481015160005460c08088015160208b810151604080519a8b52918a01979097529390960b878401526001600160a01b03909116928601929092526001600160801b0380851660808701529290930390911660a0840152905130927f507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a58492908290030190a260408082015160c083015160608085015184518b8152602081018b90526001600160a01b03948516818701526001600160801b039093169183019190915260020b60808201529151908e169133917fc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca679181900360a00190a350506000805460ff60f01b1916600160f01b17905550919890975095505050505050565b6004546001600160801b031681565b6003546001600160801b0380821691600160801b90041682565b60088161ffff811061169a57600080fd5b015463ffffffff81169150600160201b810460060b90600160581b81046001600160a01b031690600160f81b900460ff1684565b600054600160f01b900460ff16611712576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b19169055611727612e7e565b60008054600160d81b900461ffff169061174360088385614333565b6000805461ffff808416600160d81b810261ffff60d81b19909316929092179092559192508316146117b0576040805161ffff80851682528316602082015281517fac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a929181900390910190a15b50506000805460ff60f01b1916600160f01b17905550565b6000546001600160a01b03811690600160a01b810460020b9061ffff600160b81b8204811691600160c81b8104821691600160d81b8204169060ff600160e81b8204811691600160f01b90041687565b600080548190600160f01b900460ff1661185f576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b191690556001600160801b0385166118c7576040805162461bcd60e51b815260206004820152601e60248201527f556e69737761705633506f6f6c3a6d696e7428293a20616d6f756e743e300000604482015290519081900360640190fd5b60008061191560405180608001604052808c6001600160a01b031681526020018b60020b81526020018a60020b815260200161190b8a6001600160801b03166143d6565b600f0b9052614419565b9250925050819350809250600080600086111561193757611934614023565b91505b8415611948576119456141e6565b90505b336001600160a01b031663d348799787878b8b6040518563ffffffff1660e01b815260040180858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b1580156119ca57600080fd5b505af11580156119de573d6000803e3d6000fd5b505050506000861115611a35576119f3614023565b6119fd838861418e565b1115611a35576040805162461bcd60e51b815260206004820152600260248201526104d360f41b604482015290519081900360640190fd5b8415611a8557611a436141e6565b611a4d828761418e565b1115611a85576040805162461bcd60e51b81526020600482015260026024820152614d3160f01b604482015290519081900360640190fd5b8960020b8b60020b8d6001600160a01b03167f7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde338d8b8b60405180856001600160a01b03168152602001846001600160801b0316815260200183815260200182815260200194505050505060405180910390a450506000805460ff60f01b1916600160f01b17905550919890975095505050505050565b60025481565b600054600160f01b900460ff16611b66576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b19169055611b7b612e7e565b6004546001600160801b031680611bbd576040805162461bcd60e51b81526020600482015260016024820152601360fa1b604482015290519081900360640190fd5b6000611bf2867f000000000000000000000000000000000000000000000000000000000000000062ffffff16620f4240614703565b90506000611c29867f000000000000000000000000000000000000000000000000000000000000000062ffffff16620f4240614703565b90506000611c35614023565b90506000611c416141e6565b90508815611c7457611c747f00000000000000000000000000000000000000000000000000000000000000008b8b613ed5565b8715611ca557611ca57f00000000000000000000000000000000000000000000000000000000000000008b8a613ed5565b336001600160a01b031663e9cbafb085858a8a6040518563ffffffff1660e01b815260040180858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b158015611d2757600080fd5b505af1158015611d3b573d6000803e3d6000fd5b505050506000611d49614023565b90506000611d556141e6565b905081611d62858861418e565b1115611d9a576040805162461bcd60e51b8152602060048201526002602482015261046360f41b604482015290519081900360640190fd5b80611da5848761418e565b1115611ddd576040805162461bcd60e51b8152602060048201526002602482015261463160f01b604482015290519081900360640190fd5b8382038382038115611e6c5760008054600160e81b9004600f16908115611e10578160ff168481611e0a57fe5b04611e13565b60005b90506001600160801b03811615611e4657600380546001600160801b038082168401166001600160801b03199091161790555b611e60818503600160801b8d6001600160801b0316613639565b60018054909101905550505b8015611ef75760008054600160e81b900460041c600f16908115611e9c578160ff168381611e9657fe5b04611e9f565b60005b90506001600160801b03811615611ed157600380546001600160801b03600160801b8083048216850182160291161790555b611eeb818403600160801b8d6001600160801b0316613639565b60028054909101905550505b8d6001600160a01b0316336001600160a01b03167fbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca6338f8f86866040518085815260200184815260200183815260200182815260200194505050505060405180910390a350506000805460ff60f01b1916600160f01b179055505050505050505050505050565b600080548190600160f01b900460ff16611fc4576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b19168155611fde600733898961476f565b60038101549091506001600160801b0390811690861611611fff578461200e565b60038101546001600160801b03165b60038201549093506001600160801b03600160801b909104811690851611612036578361204c565b6003810154600160801b90046001600160801b03165b91506001600160801b038316156120b1576003810180546001600160801b031981166001600160801b039182168690038216179091556120b1907f0000000000000000000000000000000000000000000000000000000000000000908a908616613ed5565b6001600160801b03821615612117576003810180546001600160801b03600160801b808304821686900382160291811691909117909155612117907f0000000000000000000000000000000000000000000000000000000000000000908a908516613ed5565b604080516001600160a01b038a1681526001600160801b0380861660208301528416818301529051600288810b92908a900b9133917f70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0919081900360600190a4506000805460ff60f01b1916600160f01b17905590969095509350505050565b60076020526000908152604090208054600182015460028301546003909301546001600160801b0392831693919281811691600160801b90041685565b60066020526000908152604090205481565b7f000000000000000000000000000000000000000000000000000000000000000081565b600054600160f01b900460ff1661224e576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b1916905560408051638da5cb5b60e01b815290516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691638da5cb5b916004808301926020929190829003018186803b1580156122bb57600080fd5b505afa1580156122cf573d6000803e3d6000fd5b505050506040513d60208110156122e557600080fd5b50516001600160a01b0316331461232d5760405162461bcd60e51b81526004018080602001828103825260508152602001806162576050913960600191505060405180910390fd5b60ff82161580612350575060048260ff16101580156123505750600a8260ff1611155b801561237a575060ff8116158061237a575060048160ff161015801561237a5750600a8160ff1611155b6123b55760405162461bcd60e51b81526004018080602001828103825260268152602001806161d56026913960400191505060405180910390fd5b60008054610ff0600484901b16840160ff908116600160e81b90810260ff60e81b19841617909355919004167f973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b1336010826040805160ff9390920683168252600f600486901c16602083015286831682820152918516606082015290519081900360800190a150506000805460ff60f01b1916600160f01b17905550565b600080548190600160f01b900460ff16612499576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b1916905560408051638da5cb5b60e01b815290516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691638da5cb5b916004808301926020929190829003018186803b15801561250657600080fd5b505afa15801561251a573d6000803e3d6000fd5b505050506040513d602081101561253057600080fd5b50516001600160a01b031633146125785760405162461bcd60e51b81526004018080602001828103825260508152602001806162576050913960600191505060405180910390fd5b6003546001600160801b039081169085161161259457836125a1565b6003546001600160801b03165b6003549092506001600160801b03600160801b9091048116908416116125c757826125db565b600354600160801b90046001600160801b03165b90506001600160801b0382161561265c576003546001600160801b038381169116141561260a57600019909101905b600380546001600160801b031981166001600160801b0391821685900382161790915561265c907f00000000000000000000000000000000000000000000000000000000000000009087908516613ed5565b6001600160801b038116156126e2576003546001600160801b03828116600160801b90920416141561268d57600019015b600380546001600160801b03600160801b8083048216859003821602918116919091179091556126e2907f00000000000000000000000000000000000000000000000000000000000000009087908416613ed5565b604080516001600160801b0380851682528316602082015281516001600160a01b0388169233927f596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151929081900390910190a36000805460ff60f01b1916600160f01b1790559094909350915050565b60608061275c612e7e565b6127d3612767612ee7565b858580806020026020016040519081016040528093929190818152602001838360200280828437600092018290525054600454600896959450600160a01b820460020b935061ffff600160b81b8304811693506001600160801b0390911691600160c81b9004166147ce565b915091509250929050565b600080548190600160f01b900460ff16612825576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b1916815560408051608081018252338152600288810b602083015287900b918101919091528190819061287e90606081016128716001600160801b038a166143d6565b600003600f0b9052614419565b925092509250816000039450806000039350600085118061289f5750600084115b156128de576003830180546001600160801b038082168089018216600160801b93849004831689019092169092029091176001600160801b0319161790555b604080516001600160801b0388168152602081018790528082018690529051600289810b92908b900b9133917f0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c919081900360600190a450506000805460ff60f01b1916600160f01b179055509094909350915050565b6000806000612962612e7e565b61296c8585614926565b600285810b810b60009081526005602052604080822087840b90930b825281206003830154600681900b93600160381b82046001600160a01b0316928492600160d81b810463ffffffff169284929091600160f81b900460ff1680612a025760405162461bcd60e51b815260040180806020018281038252603a815260200180615f4f603a913960400191505060405180910390fd5b6003820154600681900b9850600160381b81046001600160a01b03169650600160d81b810463ffffffff169450600160f81b900460ff1680612a755760405162461bcd60e51b815260040180806020018281038252603a81526020018061621d603a913960400191505060405180910390fd5b50506040805160e0810182526000546001600160a01b0381168252600160a01b8104600290810b810b810b6020840181905261ffff600160b81b8404811695850195909552600160c81b830485166060850152600160d81b8304909416608084015260ff600160e81b8304811660a0850152600160f01b909204909116151560c08301529093508e810b91900b12159050612b1e57509390940396509003935090039050612ba1565b8a60020b816020015160020b1215612b92576000612b3a612ee7565b6020830151604084015160045460608601519394506000938493612b70936008938893879392916001600160801b031690613763565b9a9003989098039b505094909603929092039650909103039250612ba1915050565b50949093039650039350900390505b9250925092565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60015481565b60056020526000908152604090208054600182015460028301546003909301546001600160801b03831693600160801b909304600f0b9290600681900b90600160381b81046001600160a01b031690600160d81b810463ffffffff1690600160f81b900460ff1688565b6000546001600160a01b031615612ceb576040805162461bcd60e51b8152602060048201526002602482015261414960f01b604482015290519081900360640190fd5b6000612cf682613a65565b9050600080612d0e612d06612ee7565b6008906149ef565b6040805160e0810182526001600160a01b038816808252600288810b6020808501829052600085870181905261ffff898116606088018190529089166080880181905260a08801839052600160c0909801979097528154600160f01b6001600160a01b0319909116871762ffffff60a01b1916600160a01b62ffffff9787900b97909716969096029590951763ffffffff60b81b1916600160c81b9091021761ffff60d81b1916600160d81b9096029590951761ffff60e81b191692909217909355835191825281019190915281519395509193507f98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c9592918290030190a150505050565b60008082600281900b620d89e71981612e2757fe5b05029050600083600281900b620d89e881612e3e57fe5b0502905060008460020b83830360020b81612e5557fe5b0560010190508062ffffff166001600160801b03801681612e7257fe5b0493505050505b919050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612ee55760405162461bcd60e51b815260040180806020018281038252603d815260200180616038603d913960400191505060405180910390fd5b565b4290565b60008060008460020b8660020b81612eff57fe5b05905060008660020b128015612f2657508460020b8660020b81612f1f57fe5b0760020b15155b15612f3057600019015b8315612fa557600080612f4283614a3b565b600182810b810b600090815260208d9052604090205460ff83169190911b80016000190190811680151597509294509092509085612f8757888360ff16860302612f9a565b88612f9182614a4d565b840360ff168603025b965050505050613023565b600080612fb483600101614a3b565b91509150600060018260ff166001901b031990506000818b60008660010b60010b815260200190815260200160002054169050806000141595508561300657888360ff0360ff1686600101010261301c565b888361301183614b19565b0360ff168660010101025b9650505050505b5094509492505050565b60008060008360020b12613044578260020b61304c565b8260020b6000035b9050620d89e881111561308a576040805162461bcd60e51b81526020600482015260016024820152601560fa1b604482015290519081900360640190fd5b60006001821661309e57600160801b6130b0565b6ffffcb933bd6fad37aa2d162d1a5940015b6001600160881b0316905060028216156130da576ffff97272373d413259a46990580e213a0260801c5b60048216156130f9576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615613118576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b6010821615613137576fffcb9843d60f6159c9db58835c9266440260801c5b6020821615613156576fff973b41fa98c081472e6896dfb254c00260801c5b6040821615613175576fff2ea16466c96a3843ec78b326b528610260801c5b6080821615613194576ffe5dee046a99a2a811c461f1969c30530260801c5b6101008216156131b4576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b6102008216156131d4576ff987a7253ac413176f2b074cf7815e540260801c5b6104008216156131f4576ff3392b0822b70005940c7a398e4b70f30260801c5b610800821615613214576fe7159475a2c29b7443b29c7fa6e889d90260801c5b611000821615613234576fd097f3bdfd2022b8845ad8f792aa58250260801c5b612000821615613254576fa9f746462d870fdf8a65dc1f90e061e50260801c5b614000821615613274576f70d869a156d2a1b890bb3df62baf32f70260801c5b618000821615613294576f31be135f97d08fd981231505542fcfa60260801c5b620100008216156132b5576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b620200008216156132d5576e5d6af8dedb81196699c329225ee6040260801c5b620400008216156132f4576d2216e584f5fa1ea926041bedfe980260801c5b62080000821615613311576b048a170391f7dc42444e8fa20260801c5b60008460020b131561332c57806000198161332857fe5b0490505b600160201b81061561333f576001613342565b60005b60ff16602082901c0192505050919050565b60008080806001600160a01b03808916908a1610158187128015906133d957600061338d8989620f42400362ffffff16620f4240613639565b9050826133a6576133a18c8c8c6001614c34565b6133b3565b6133b38b8d8c6001614caf565b95508581106133c4578a96506133d3565b6133d08c8b8386614d8c565b96505b50613423565b816133f0576133eb8b8b8b6000614caf565b6133fd565b6133fd8a8c8b6000614c34565b935083886000031061341157899550613423565b6134208b8a8a60000385614e3c565b95505b6001600160a01b038a8116908716148215613486578080156134425750815b61345857613453878d8c6001614caf565b61345a565b855b9550808015613467575081155b61347d57613478878d8c6000614c34565b61347f565b845b94506134d0565b8080156134905750815b6134a6576134a18c888c6001614c34565b6134a8565b855b95508080156134b5575081155b6134cb576134c68c888c6000614caf565b6134cd565b845b94505b811580156134e057508860000385115b156134ec578860000394505b81801561350b57508a6001600160a01b0316876001600160a01b031614155b1561351a578589039350613537565b613534868962ffffff168a620f42400362ffffff16614703565b93505b50505095509550955095915050565b6000600160ff1b821061359f576040805162461bcd60e51b815260206004820152601c60248201527b73616665636173743a746f496e7432353628293a20793c322a32353560201b604482015290519081900360640190fd5b5090565b808203828113156000831215146135eb5760405162461bcd60e51b81526004018080602001828103825260288152602001806160756028913960400191505060405180910390fd5b92915050565b818101828112156000831215146135eb5760405162461bcd60e51b81526004018080602001828103825260288152602001806160106028913960400191505060405180910390fd5b60008080600019858709868602925082811090839003039050806136b757600084116136ac576040805162461bcd60e51b815260206004820181905260248201527f66756c6c6d6174683a6d756c44697628293a2064656e6f6d696e61746f723e30604482015290519081900360640190fd5b50829004905061375c565b8084116136f55760405162461bcd60e51b815260040180806020018281038252602481526020018061609d6024913960400191505060405180910390fd5b6000848688096000868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b60008063ffffffff8716613809576000898661ffff1661ffff811061378457fe5b60408051608081018252919092015463ffffffff808216808452600160201b8304600690810b810b900b6020850152600160581b83046001600160a01b031694840194909452600160f81b90910460ff16151560608301529092508a16146137f5576137f2818a8988614eec565b90505b8060200151816040015192509250506138e9565b86880360008061381e8c8c858c8c8c8c614f8f565b91509150816000015163ffffffff168363ffffffff1614156138505781602001518260400151945094505050506138e9565b805163ffffffff848116911614156138785780602001518160400151945094505050506138e9565b8151815160208085015190840151918390039286039163ffffffff80841692908516910360060b816138a657fe5b05028460200151018263ffffffff168263ffffffff1686604001518660400151036001600160a01b031602816138d857fe5b048560400151019650965050505050505b97509795505050505050565b600295860b860b60009081526020979097526040909620600181018054909503909455938301805490920390915560038201805463ffffffff600160d81b6001600160a01b03600160381b808504821690960316909402600160381b600160d81b031990921691909117600681810b90960390950b66ffffffffffffff1666ffffffffffffff199095169490941782810485169095039093160263ffffffff60d81b1990931692909217905554600160801b9004600f0b90565b60008082600f0b1215613a1457826001600160801b03168260000384039150816001600160801b031610613a0f576040805162461bcd60e51b81526020600482015260026024820152614c5360f01b604482015290519081900360640190fd5b6135eb565b826001600160801b03168284019150816001600160801b031610156135eb576040805162461bcd60e51b81526020600482015260026024820152614c4160f01b604482015290519081900360640190fd5b60006401000276a36001600160a01b03831610801590613aa1575073fffd8963efd1fc6a506488495d951d5263988d266001600160a01b038316105b613ad6576040805162461bcd60e51b81526020600482015260016024820152602960f91b604482015290519081900360640190fd5b600160201b600160c01b03602083901b166001600160801b03811160071b81811c6001600160401b03811160061b90811c63ffffffff811160051b90811c61ffff811160041b90811c60ff8111600390811b91821c600f811160021b90811c918211600190811b92831c97908811961790941790921717909117171760808110613b6857607f810383901c9150613b72565b80607f0383901b91505b908002607f81811c60ff83811c9190911c800280831c81831c1c800280841c81841c1c800280851c81851c1c800280861c81861c1c800280871c81871c1c800280881c81881c1c800280891c81891c1c8002808a1c818a1c1c8002808b1c818b1c1c8002808c1c818c1c1c8002808d1c818d1c1c8002808e1c9c81901c9c909c1c80029c8d901c9e9d607f198f0160401b60c09190911c6001603f1b161760c19b909b1c6001603e1b169a909a1760c29990991c6001603d1b169890981760c39790971c6001603c1b169690961760c49590951c6001603b1b169490941760c59390931c6001603a1b169290921760c69190911c600160391b161760c79190911c600160381b161760c89190911c600160371b161760c99190911c600160361b161760ca9190911c600160351b161760cb9190911c600160341b161760cc9190911c600160331b161760cd9190911c600160321b1617693627a301d71055774c8581026f028f6481ab7f045a5af012a19d003aa9198101608090811d906fdb2df09e81959a81455e260799a0632f8301901d600281810b9083900b14613d4157886001600160a01b0316613d258261302d565b6001600160a01b03161115613d3a5781613d3c565b805b613d43565b815b9998505050505050505050565b6000806000898961ffff1661ffff8110613d6657fe5b60408051608081018252919092015463ffffffff808216808452600160201b8304600690810b810b900b6020850152600160581b83046001600160a01b031694840194909452600160f81b90910460ff161515606083015290925089161415613dd557888592509250506138e9565b8461ffff168461ffff16118015613df657506001850361ffff168961ffff16145b15613e0357839150613e07565b8491505b8161ffff168960010161ffff1681613e1b57fe5b069250613e2a81898989614eec565b8a8461ffff1661ffff8110613e3b57fe5b825191018054602084015160408501516060909501511515600160f81b026001600160f81b036001600160a01b03909616600160581b02600160581b600160f81b031960069390930b66ffffffffffffff16600160201b0266ffffffffffffff60201b1963ffffffff90971663ffffffff199095169490941795909516929092171692909217929092161790555097509795505050505050565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b1781529251825160009485949389169392918291908083835b60208310613f515780518252601f199092019160209182019101613f32565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114613fb3576040519150601f19603f3d011682016040523d82523d6000602084013e613fb8565b606091505b5091509150818015613fe6575080511580613fe65750808060200190516020811015613fe357600080fd5b50515b61401c576040805162461bcd60e51b81526020600482015260026024820152612a2360f11b604482015290519081900360640190fd5b5050505050565b604080513060248083019190915282518083039091018152604490910182526020810180516001600160e01b03166370a0823160e01b17815291518151600093849384936001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001693919290918291908083835b602083106140bc5780518252601f19909201916020918201910161409d565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d806000811461411c576040519150601f19603f3d011682016040523d82523d6000602084013e614121565b606091505b509150915081801561413557506020815110155b6141705760405162461bcd60e51b8152600401808060200182810382526031815260200180615eb26031913960400191505060405180910390fd5b80806020019051602081101561418557600080fd5b50519250505090565b808201828110156135eb576040805162461bcd60e51b815260206004820152601f60248201527f6c6f77676173736166656d6174683a616464282920287a3d782b79293e3d7800604482015290519081900360640190fd5b604080513060248083019190915282518083039091018152604490910182526020810180516001600160e01b03166370a0823160e01b17815291518151600093849384936001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001693919290918291908083835b6020831061427f5780518252601f199092019160209182019101614260565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d80600081146142df576040519150601f19603f3d011682016040523d82523d6000602084013e6142e4565b606091505b50915091508180156142f857506020815110155b6141705760405162461bcd60e51b81526004018080602001828103825260318152602001806162a76031913960400191505060405180910390fd5b6000808361ffff1611614371576040805162461bcd60e51b81526020600482015260016024820152604960f81b604482015290519081900360640190fd5b8261ffff168261ffff161161438757508161375c565b825b8261ffff168161ffff1610156143cd576001858261ffff1661ffff81106143ac57fe5b01805463ffffffff191663ffffffff92909216919091179055600101614389565b50909392505050565b80600f81900b8114612e795760405162461bcd60e51b81526004018080602001828103825260228152602001806161fb6022913960400191505060405180910390fd5b6000806000614426612e7e565b61443884602001518560400151614926565b6040805160e0810182526000546001600160a01b0381168252600160a01b8104600290810b810b900b602080840182905261ffff600160b81b8404811685870152600160c81b84048116606080870191909152600160d81b8504909116608086015260ff600160e81b8504811660a0870152600160f01b909404909316151560c0850152885190890151948901519289015193946144dc9491939092909190615189565b935060008560600151600f0b60001461465657856020015160020b826020015160020b12156145335761452c614515876020015161302d565b614522886040015161302d565b88606001516153cd565b9350614656565b856040015160020b826020015160020b121561462c575060045460408201516001600160801b039091169081906145889061456c612ee7565b6020860151606087015160808801516008949392918791613d50565b6000805461ffff60c81b1916600160c81b61ffff938416021761ffff60b81b1916600160b81b9390921692909202179055825160408801516145d891906145ce9061302d565b89606001516153cd565b94506145f66145ea886020015161302d565b845160608a0151615411565b93506146068188606001516139af565b600480546001600160801b0319166001600160801b039290921691909117905550614656565b61465361463c876020015161302d565b614649886040015161302d565b8860600151615411565b92505b85516020808801516040808a0151868401516060808d0151895185516001600160a01b03998a168152600297880b9881019890985293860b878601529190940b93850193909352600f9290920b60808401526001600160801b03851660a084015260c0830188905260e08301879052921661010082015290517fd1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f918190036101200190a150509193909250565b6000614710848484613639565b90506000828061471c57fe5b848609111561375c5760001981106147655760405162461bcd60e51b81526004018080602001828103825260298152602001806161ac6029913960400191505060405180910390fd5b6001019392505050565b6040805160609490941b6001600160601b031916602080860191909152600293840b60e890811b60348701529290930b90911b60378401528051808403601a018152603a90930181528251928201929092206000908152929052902090565b60608060008361ffff161161480e576040805162461bcd60e51b81526020600482015260016024820152604960f81b604482015290519081900360640190fd5b86516001600160401b038111801561482557600080fd5b5060405190808252806020026020018201604052801561484f578160200160208202803683370190505b50915086516001600160401b038111801561486957600080fd5b50604051908082528060200260200182016040528015614893578160200160208202803683370190505b50905060005b8751811015614919576148c48a8a8a84815181106148b357fe5b60200260200101518a8a8a8a613763565b8483815181106148d057fe5b602002602001018484815181106148e357fe5b60200260200101826001600160a01b03166001600160a01b03168152508260060b60060b81525050508080600101915050614899565b5097509795505050505050565b8060020b8260020b12614966576040805162461bcd60e51b8152602060048201526003602482015262544c5560e81b604482015290519081900360640190fd5b620d89e719600283900b12156149a9576040805162461bcd60e51b8152602060048201526003602482015262544c4d60e81b604482015290519081900360640190fd5b620d89e8600282900b13156149eb576040805162461bcd60e51b815260206004820152600360248201526254554d60e81b604482015290519081900360640190fd5b5050565b6040805160808101825263ffffffff9283168082526000602083018190529282019290925260016060909101819052835463ffffffff1916909117909116600160f81b17909155908190565b60020b600881901d9161010090910790565b6000808211614a8d5760405162461bcd60e51b8152600401808060200182810382526021815260200180615f2e6021913960400191505060405180910390fd5b600160801b8210614aa057608091821c91015b600160401b8210614ab357604091821c91015b600160201b8210614ac657602091821c91015b620100008210614ad857601091821c91015b6101008210614ae957600891821c91015b60108210614af957600491821c91015b60048210614b0957600291821c91015b60028210612e7957600101919050565b6000808211614b595760405162461bcd60e51b8152600401808060200182810382526022815260200180615fbf6022913960400191505060405180910390fd5b5060ff6001600160801b03821615614b7457607f1901614b7c565b608082901c91505b6001600160401b03821615614b9457603f1901614b9c565b604082901c91505b63ffffffff821615614bb157601f1901614bb9565b602082901c91505b61ffff821615614bcc57600f1901614bd4565b601082901c91505b60ff821615614be65760071901614bee565b600882901c91505b600f821615614c005760031901614c08565b600482901c91505b6003821615614c1a5760011901614c22565b600282901c91505b6001821615612e795760001901919050565b6000836001600160a01b0316856001600160a01b03161115614c54579293925b81614c8157614c7c836001600160801b03168686036001600160a01b0316600160601b613639565b614ca4565b614ca4836001600160801b03168686036001600160a01b0316600160601b614703565b90505b949350505050565b6000836001600160a01b0316856001600160a01b03161115614ccf579293925b600160601b600160e01b03606084901b166001600160a01b038686038116908716614d2b5760405162461bcd60e51b815260040180806020018281038252602f815260200180615fe1602f913960400191505060405180910390fd5b83614d5b57866001600160a01b0316614d4e8383896001600160a01b0316613639565b81614d5557fe5b04614d81565b614d81614d728383896001600160a01b0316614703565b886001600160a01b0316615440565b979650505050505050565b600080856001600160a01b031611614dd55760405162461bcd60e51b81526004018080602001828103825260358152602001806161196035913960400191505060405180910390fd5b6000846001600160801b031611614e1d5760405162461bcd60e51b8152600401808060200182810382526036815260200180615f896036913960400191505060405180910390fd5b81614e2f57614c7c858585600161544b565b614ca4858585600161555e565b600080856001600160a01b031611614e855760405162461bcd60e51b81526004018080602001828103825260378152602001806160c16037913960400191505060405180910390fd5b6000846001600160801b031611614ecd5760405162461bcd60e51b81526004018080602001828103825260388152602001806161746038913960400191505060405180910390fd5b81614edf57614c7c858585600061555e565b614ca4858585600061544b565b614ef4615e14565b600085600001518503905060405180608001604052808663ffffffff1681526020018263ffffffff168660020b0288602001510160060b81526020016000856001600160801b031611614f48576001614f4a565b845b6001600160801b031663ffffffff60801b608085901b1681614f6857fe5b048860400151016001600160a01b0316815260200160011515815250915050949350505050565b614f97615e14565b614f9f615e14565b888561ffff1661ffff8110614fb057fe5b60408051608081018252919092015463ffffffff8116808352600160201b8204600690810b810b900b6020840152600160581b82046001600160a01b031693830193909352600160f81b900460ff1615156060820152925061501490899089615673565b1561504c578663ffffffff16826000015163ffffffff161415615036576138e9565b8161504383898988614eec565b915091506138e9565b888361ffff168660010161ffff168161506157fe5b0661ffff1661ffff811061507157fe5b60408051608081018252929091015463ffffffff81168352600160201b8104600690810b810b900b60208401526001600160a01b03600160581b8204169183019190915260ff600160f81b9091041615156060820181905290925061512657604080516080810182528a5463ffffffff81168252600160201b8104600690810b810b900b6020830152600160581b81046001600160a01b031692820192909252600160f81b90910460ff161515606082015291505b61513588836000015189615673565b61516c576040805162461bcd60e51b815260206004820152600360248201526213d31160ea1b604482015290519081900360640190fd5b6151798989898887615734565b9150915097509795505050505050565b6000615198600787878761476f565b60015460025491925090600080600f87900b156152de5760006151b9612ee7565b60008054600454929350909182916152039160089186918591600160a01b810460020b9161ffff600160b81b83048116926001600160801b0390921691600160c81b900416613763565b909250905061523d60058d8b8d8b8b87898b60007f00000000000000000000000000000000000000000000000000000000000000006158d2565b945061527460058c8b8d8b8b87898b60017f00000000000000000000000000000000000000000000000000000000000000006158d2565b935084156152a8576152a860068d7f0000000000000000000000000000000000000000000000000000000000000000615a8b565b83156152da576152da60068c7f0000000000000000000000000000000000000000000000000000000000000000615a8b565b5050505b6000806152f060058c8c8b8a8a615b23565b9092509050615301878a8484615bcf565b600089600f0b121561532f57831561531e5761531e60058c615d64565b821561532f5761532f60058b615d64565b604080516001600160a01b038e16815260028d810b60208301528c810b828401528a900b6060820152600f8b900b608082015260a0810188905260c0810187905260e08101849052610100810183905285151561012082015284151561014082015290517f2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24918190036101600190a150505050505095945050505050565b60008082600f0b126153f3576153ee6153e98585856001614caf565b613546565b614ca7565b6154066153e98585856000036000614caf565b600003949350505050565b60008082600f0b1261542d576153ee6153e98585856001614c34565b6154066153e98585856000036000614c34565b808204910615150190565b600081156154be5760006001600160a01b038411156154815761547c84600160601b876001600160801b0316613639565b615499565b6001600160801b038516606085901b8161549757fe5b045b90506154b66154b16001600160a01b0388168361418e565b615d90565b915050614ca7565b60006001600160a01b038411156154ec576154e784600160601b876001600160801b0316614703565b615503565b615503606085901b6001600160801b038716615440565b905080866001600160a01b03161161554c5760405162461bcd60e51b815260040180806020018281038252604b815260200180615ee3604b913960600191505060405180910390fd5b6001600160a01b038616039050614ca7565b60008261556c575083614ca7565b600160601b600160e01b03606085901b1682156155fa576001600160a01b0386168481029085828161559a57fe5b0414156155cb578181018281106155c9576155bf83896001600160a01b031683614703565b9350505050614ca7565b505b6155f1826155ec878a6001600160a01b031686816155e557fe5b049061418e565b615440565b92505050614ca7565b6001600160a01b0386168481029085828161561157fe5b0414801561561e57508082115b6156595760405162461bcd60e51b8152600401808060200182810382526076815260200180615e3c6076913960800191505060405180910390fd5b8082036155bf6154b1846001600160a01b038b1684614703565b60008363ffffffff168363ffffffff161115801561569d57508363ffffffff168263ffffffff1611155b156156b9578163ffffffff168363ffffffff161115905061375c565b60008463ffffffff168463ffffffff16116156e0578363ffffffff16600160201b016156e8565b8363ffffffff165b64ffffffffff16905060008563ffffffff168463ffffffff1611615718578363ffffffff16600160201b01615720565b8363ffffffff165b64ffffffffff169091111595945050505050565b61573c615e14565b615744615e14565b60008361ffff168560010161ffff168161575a57fe5b0661ffff169050600060018561ffff16830103905060005b506002818301048961ffff8716828161578757fe5b0661ffff811061579357fe5b60408051608081018252929091015463ffffffff81168352600160201b8104600690810b810b900b60208401526001600160a01b03600160581b8204169183019190915260ff600160f81b909104161515606082018190529095506157fd57806001019250615772565b898661ffff16826001018161580e57fe5b0661ffff811061581a57fe5b60408051608081018252929091015463ffffffff81168352600160201b8104600690810b810b900b60208401526001600160a01b03600160581b8204169183019190915260ff600160f81b90910416151560608201528551909450600090615884908b908b615673565b905080801561589d575061589d8a8a8760000151615673565b156158a857506158c5565b806158b8576001820392506158bf565b8160010193505b50615772565b5050509550959350505050565b60028a810b900b600090815260208c90526040812080546001600160801b0316826158fd828d6139af565b9050846001600160801b0316816001600160801b0316111561594b576040805162461bcd60e51b81526020600482015260026024820152614c4f60f01b604482015290519081900360640190fd5b6001600160801b0382811615908216158114159450156159f0578c60020b8e60020b136159d857600183018b9055600283018a9055600383018054600160381b600160d81b031916600160381b6001600160a01b038c16021766ffffffffffffff191666ffffffffffffff60068b900b161763ffffffff60d81b1916600160d81b63ffffffff8a16021790555b6003830180546001600160f81b0316600160f81b1790555b82546001600160801b0319166001600160801b03821617835585615a39578254615a3490615a2f90600160801b9004600f90810b810b908f900b6135f1565b6143d6565b615a5a565b8254615a5a90615a2f90600160801b9004600f90810b810b908f900b6135a3565b8354600f9190910b6001600160801b03908116600160801b0291161790925550909c9b505050505050505050505050565b8060020b8260020b81615a9a57fe5b0760020b15615ada5760405162461bcd60e51b81526004018080602001828103825260218152602001806160f86021913960400191505060405180910390fd5b600080615af58360020b8560020b81615aef57fe5b05614a3b565b600191820b820b60009081526020979097526040909620805460ff9097169190911b90951890945550505050565b600285810b80820b60009081526020899052604080822088850b850b83529082209193849391929184918291908a900b12615b6957505060018201546002830154615b7c565b8360010154880391508360020154870390505b6000808b60020b8b60020b1215615b9e57505060018301546002840154615bb1565b84600101548a0391508460020154890390505b92909803979097039b96909503949094039850939650505050505050565b6040805160a08101825285546001600160801b0390811682526001870154602083015260028701549282019290925260038601548083166060830152600160801b900490911660808201526000600f85900b615c6e5781516001600160801b0316615c66576040805162461bcd60e51b815260206004820152600260248201526104e560f41b604482015290519081900360640190fd5b508051615c7d565b8151615c7a90866139af565b90505b6000615ca18360200151860384600001516001600160801b0316600160801b613639565b90506000615cc78460400151860385600001516001600160801b0316600160801b613639565b905086600f0b600014615cee5787546001600160801b0319166001600160801b0384161788555b60018801869055600288018590556001600160801b038216151580615d1c57506000816001600160801b0316115b15615d5a576003880180546001600160801b031981166001600160801b039182168501821617808216600160801b9182900483168501909216021790555b5050505050505050565b600290810b810b6000908152602092909252604082208281556001810183905590810182905560030155565b806001600160a01b0381168114612e795760405162461bcd60e51b815260040180806020018281038252602681526020018061614e6026913960400191505060405180910390fd5b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c081019190915290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe7371727470726963656d6174683a676574684e65787453717274507269636546726f6d416d6f756e7430526f756e64696e67557028293a202870726f647563743d616d6f756e742a7371727450783936292f616d6f756e743d3d737172745078393626266e756d657261746f72313e70726f64756374556e69737761705633506f6f6c3a62616c616e636530282920737563636573732626646174612e6c656e6774683e3d33327371727470726963656d6174683a6765744e65787453717274507269636546726f6d416d6f756e7431526f756e64696e67446f776e2829207371727450583936203e2071756f7469656e746269746d6174683a6d6f73745369676e6966696361746e42697428293a20783e30556e69737761705633506f6f6c3a736e617073686f7443756d756c6174697665496e7369646528293a20696e697469616c697a65644c6f7765727371727470726963656d6174683a6765744e65787453717274507269636546726f6d496e70757428293a206c69717569646974793e306269746d6174683a6c656173745369676e69666963616e7442697428293a20783e307371727470726963656d6174683a676574416d6f756e743044656c746128293a2073717274526174696f6e415839366c6f77676173736166656d6174683a61646428293a20287a3d782b79293e3d783d3d28793e3d30294e6f44656c656761746543616c6c3a636865636b4e6f7444656c656761746543616c6c282920616464726573732874686973293d3d6f726967696e616c6c6f77676173736166656d6174683a73756228293a20287a3d782d79293c3d783d3d28793e3d302966756c6c6d6174683a6d756c44697628293a2064656e6f6d696e61746f723e70726f64317371727470726963656d6174683a206765744e65787453717274507269636546726f6d4f757470757428293a2073717274505839363e305469636b4269746d61703a666c69705469636b28293a20666c69705469636b28297371727470726963656d6174683a6765744e65787453717274507269636546726f6d496e70757428293a2073717274505839363e3073616665636173743a746f55696e7431363028293a20287a3d75696e74313630287929293d797371727470726963656d6174683a206765744e65787453717274507269636546726f6d4f757470757428293a206c69717569646974793e3066756c6c6d6174683a6d756c446976526f756e64696e67557028293a20726573756c74203c206d6178556e69737761705633506f6f6c3a73657446656550726f746f636f6c3a206c696e652038343073616665636173743a746f496e7431323828293a20287a3d696e7431323828792929556e69737761705633506f6f6c3a736e617073686f7443756d756c6174697665496e7369646528293a20696e697469616c697a65645570706572556e69737761705633506f6f6c3a6f6e6c79466163746f72794f776e657228293a206d73672e73656e6465723d3d49556e69737761705633466163746f727928666163746f7279292e6f776e65722829556e69737761705633506f6f6c3a62616c616e636531282920737563636573732626646174612e6c656e6774683e3d3332a264697066735822122095074e04ee96156b634d25a3aa164b995f001ef9e273c3f07d788621156cd50f64736f6c63430007060033556e69737761705633466163746f72793a656e61626c65466565416d6f756e7428293a206665653c31303030303030556e69737761705633466163746f72793a656e61626c65466565416d6f756e7428293a206d73672e73656e6465723d3d6f776e6572556e69737761705633466163746f72793a656e61626c65466565416d6f756e7428293a20666565416d6f756e745469636b53706163696e675b6665655d3d3d30556e69737761705633466163746f72793a637265617465506f6f6c28293a207265717569726520676574506f6f6c5b746f6b656e305d5b746f6b656e315d5b6665655d3d61646472657373283029556e69737761705633466163746f72793a637265617465506f6f6c282920746f6b656e30213d616464726573732830294e6f44656c656761746543616c6c3a636865636b4e6f7444656c656761746543616c6c282920616464726573732874686973293d3d6f726967696e616c556e69737761705633466163746f72793a656e61626c65466565416d6f756e742829207469636b53706163696e673e30202626207469636b53706163696e673c3136333834556e69737761705633466163746f72793a637265617465506f6f6c28293a207469636b53706163696e67213d30556e69737761705633466163746f72793a7365744f776e657228293a206d73672e73656e6465723d6f776e6572556e69737761705633466163746f72793a637265617465506f6f6c282920746f6b656e41213d746f6b656e42a2646970667358221220a133227a1ff137fc75d785111e670fe3f5e9ef2fc2a2de7e167986a65eaeeafb64736f6c63430007060033c66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc",
}

// UniswapV3FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3FactoryMetaData.ABI instead.
var UniswapV3FactoryABI = UniswapV3FactoryMetaData.ABI

// UniswapV3FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniswapV3FactoryMetaData.Bin instead.
var UniswapV3FactoryBin = UniswapV3FactoryMetaData.Bin

// DeployUniswapV3Factory deploys a new Ethereum contract, binding an instance of UniswapV3Factory to it.
func DeployUniswapV3Factory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniswapV3Factory, error) {
	parsed, err := UniswapV3FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniswapV3FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniswapV3Factory{UniswapV3FactoryCaller: UniswapV3FactoryCaller{contract: contract}, UniswapV3FactoryTransactor: UniswapV3FactoryTransactor{contract: contract}, UniswapV3FactoryFilterer: UniswapV3FactoryFilterer{contract: contract}}, nil
}

// UniswapV3Factory is an auto generated Go binding around an Ethereum contract.
type UniswapV3Factory struct {
	UniswapV3FactoryCaller     // Read-only binding to the contract
	UniswapV3FactoryTransactor // Write-only binding to the contract
	UniswapV3FactoryFilterer   // Log filterer for contract events
}

// UniswapV3FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3FactorySession struct {
	Contract     *UniswapV3Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV3FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3FactoryCallerSession struct {
	Contract *UniswapV3FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UniswapV3FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3FactoryTransactorSession struct {
	Contract     *UniswapV3FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UniswapV3FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3FactoryRaw struct {
	Contract *UniswapV3Factory // Generic contract binding to access the raw methods on
}

// UniswapV3FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3FactoryCallerRaw struct {
	Contract *UniswapV3FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3FactoryTransactorRaw struct {
	Contract *UniswapV3FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3Factory creates a new instance of UniswapV3Factory, bound to a specific deployed contract.
func NewUniswapV3Factory(address common.Address, backend bind.ContractBackend) (*UniswapV3Factory, error) {
	contract, err := bindUniswapV3Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Factory{UniswapV3FactoryCaller: UniswapV3FactoryCaller{contract: contract}, UniswapV3FactoryTransactor: UniswapV3FactoryTransactor{contract: contract}, UniswapV3FactoryFilterer: UniswapV3FactoryFilterer{contract: contract}}, nil
}

// NewUniswapV3FactoryCaller creates a new read-only instance of UniswapV3Factory, bound to a specific deployed contract.
func NewUniswapV3FactoryCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3FactoryCaller, error) {
	contract, err := bindUniswapV3Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3FactoryCaller{contract: contract}, nil
}

// NewUniswapV3FactoryTransactor creates a new write-only instance of UniswapV3Factory, bound to a specific deployed contract.
func NewUniswapV3FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3FactoryTransactor, error) {
	contract, err := bindUniswapV3Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3FactoryTransactor{contract: contract}, nil
}

// NewUniswapV3FactoryFilterer creates a new log filterer instance of UniswapV3Factory, bound to a specific deployed contract.
func NewUniswapV3FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3FactoryFilterer, error) {
	contract, err := bindUniswapV3Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3FactoryFilterer{contract: contract}, nil
}

// bindUniswapV3Factory binds a generic wrapper to an already deployed contract.
func bindUniswapV3Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV3FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Factory *UniswapV3FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Factory.Contract.UniswapV3FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Factory *UniswapV3FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Factory.Contract.UniswapV3FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Factory *UniswapV3FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Factory.Contract.UniswapV3FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Factory *UniswapV3FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Factory *UniswapV3FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Factory *UniswapV3FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Factory.Contract.contract.Transact(opts, method, params...)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_UniswapV3Factory *UniswapV3FactoryCaller) FeeAmountTickSpacing(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3Factory.contract.Call(opts, &out, "feeAmountTickSpacing", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_UniswapV3Factory *UniswapV3FactorySession) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _UniswapV3Factory.Contract.FeeAmountTickSpacing(&_UniswapV3Factory.CallOpts, arg0)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_UniswapV3Factory *UniswapV3FactoryCallerSession) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _UniswapV3Factory.Contract.FeeAmountTickSpacing(&_UniswapV3Factory.CallOpts, arg0)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniswapV3Factory *UniswapV3FactoryCaller) GetPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Factory.contract.Call(opts, &out, "getPool", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniswapV3Factory *UniswapV3FactorySession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _UniswapV3Factory.Contract.GetPool(&_UniswapV3Factory.CallOpts, arg0, arg1, arg2)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniswapV3Factory *UniswapV3FactoryCallerSession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _UniswapV3Factory.Contract.GetPool(&_UniswapV3Factory.CallOpts, arg0, arg1, arg2)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3Factory *UniswapV3FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3Factory *UniswapV3FactorySession) Owner() (common.Address, error) {
	return _UniswapV3Factory.Contract.Owner(&_UniswapV3Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3Factory *UniswapV3FactoryCallerSession) Owner() (common.Address, error) {
	return _UniswapV3Factory.Contract.Owner(&_UniswapV3Factory.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniswapV3Factory *UniswapV3FactoryCaller) Parameters(opts *bind.CallOpts) (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	var out []interface{}
	err := _UniswapV3Factory.contract.Call(opts, &out, "parameters")

	outstruct := new(struct {
		Factory     common.Address
		Token0      common.Address
		Token1      common.Address
		Fee         *big.Int
		TickSpacing *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Factory = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickSpacing = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniswapV3Factory *UniswapV3FactorySession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _UniswapV3Factory.Contract.Parameters(&_UniswapV3Factory.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniswapV3Factory *UniswapV3FactoryCallerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _UniswapV3Factory.Contract.Parameters(&_UniswapV3Factory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_UniswapV3Factory *UniswapV3FactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _UniswapV3Factory.contract.Transact(opts, "createPool", tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_UniswapV3Factory *UniswapV3FactorySession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _UniswapV3Factory.Contract.CreatePool(&_UniswapV3Factory.TransactOpts, tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_UniswapV3Factory *UniswapV3FactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _UniswapV3Factory.Contract.CreatePool(&_UniswapV3Factory.TransactOpts, tokenA, tokenB, fee)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_UniswapV3Factory *UniswapV3FactoryTransactor) EnableFeeAmount(opts *bind.TransactOpts, fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _UniswapV3Factory.contract.Transact(opts, "enableFeeAmount", fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_UniswapV3Factory *UniswapV3FactorySession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _UniswapV3Factory.Contract.EnableFeeAmount(&_UniswapV3Factory.TransactOpts, fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_UniswapV3Factory *UniswapV3FactoryTransactorSession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _UniswapV3Factory.Contract.EnableFeeAmount(&_UniswapV3Factory.TransactOpts, fee, tickSpacing)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_UniswapV3Factory *UniswapV3FactoryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _UniswapV3Factory.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_UniswapV3Factory *UniswapV3FactorySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _UniswapV3Factory.Contract.SetOwner(&_UniswapV3Factory.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_UniswapV3Factory *UniswapV3FactoryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _UniswapV3Factory.Contract.SetOwner(&_UniswapV3Factory.TransactOpts, _owner)
}

// UniswapV3FactoryFeeAmountEnabledIterator is returned from FilterFeeAmountEnabled and is used to iterate over the raw logs and unpacked data for FeeAmountEnabled events raised by the UniswapV3Factory contract.
type UniswapV3FactoryFeeAmountEnabledIterator struct {
	Event *UniswapV3FactoryFeeAmountEnabled // Event containing the contract specifics and raw log

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
func (it *UniswapV3FactoryFeeAmountEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3FactoryFeeAmountEnabled)
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
		it.Event = new(UniswapV3FactoryFeeAmountEnabled)
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
func (it *UniswapV3FactoryFeeAmountEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3FactoryFeeAmountEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3FactoryFeeAmountEnabled represents a FeeAmountEnabled event raised by the UniswapV3Factory contract.
type UniswapV3FactoryFeeAmountEnabled struct {
	Fee         *big.Int
	TickSpacing *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeAmountEnabled is a free log retrieval operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_UniswapV3Factory *UniswapV3FactoryFilterer) FilterFeeAmountEnabled(opts *bind.FilterOpts, fee []*big.Int, tickSpacing []*big.Int) (*UniswapV3FactoryFeeAmountEnabledIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _UniswapV3Factory.contract.FilterLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3FactoryFeeAmountEnabledIterator{contract: _UniswapV3Factory.contract, event: "FeeAmountEnabled", logs: logs, sub: sub}, nil
}

// WatchFeeAmountEnabled is a free log subscription operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_UniswapV3Factory *UniswapV3FactoryFilterer) WatchFeeAmountEnabled(opts *bind.WatchOpts, sink chan<- *UniswapV3FactoryFeeAmountEnabled, fee []*big.Int, tickSpacing []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _UniswapV3Factory.contract.WatchLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3FactoryFeeAmountEnabled)
				if err := _UniswapV3Factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
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

// ParseFeeAmountEnabled is a log parse operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_UniswapV3Factory *UniswapV3FactoryFilterer) ParseFeeAmountEnabled(log types.Log) (*UniswapV3FactoryFeeAmountEnabled, error) {
	event := new(UniswapV3FactoryFeeAmountEnabled)
	if err := _UniswapV3Factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3FactoryOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the UniswapV3Factory contract.
type UniswapV3FactoryOwnerChangedIterator struct {
	Event *UniswapV3FactoryOwnerChanged // Event containing the contract specifics and raw log

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
func (it *UniswapV3FactoryOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3FactoryOwnerChanged)
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
		it.Event = new(UniswapV3FactoryOwnerChanged)
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
func (it *UniswapV3FactoryOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3FactoryOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3FactoryOwnerChanged represents a OwnerChanged event raised by the UniswapV3Factory contract.
type UniswapV3FactoryOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_UniswapV3Factory *UniswapV3FactoryFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*UniswapV3FactoryOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniswapV3Factory.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3FactoryOwnerChangedIterator{contract: _UniswapV3Factory.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_UniswapV3Factory *UniswapV3FactoryFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *UniswapV3FactoryOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniswapV3Factory.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3FactoryOwnerChanged)
				if err := _UniswapV3Factory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_UniswapV3Factory *UniswapV3FactoryFilterer) ParseOwnerChanged(log types.Log) (*UniswapV3FactoryOwnerChanged, error) {
	event := new(UniswapV3FactoryOwnerChanged)
	if err := _UniswapV3Factory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3FactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the UniswapV3Factory contract.
type UniswapV3FactoryPoolCreatedIterator struct {
	Event *UniswapV3FactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *UniswapV3FactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3FactoryPoolCreated)
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
		it.Event = new(UniswapV3FactoryPoolCreated)
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
func (it *UniswapV3FactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3FactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3FactoryPoolCreated represents a PoolCreated event raised by the UniswapV3Factory contract.
type UniswapV3FactoryPoolCreated struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Pool        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_UniswapV3Factory *UniswapV3FactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, fee []*big.Int) (*UniswapV3FactoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _UniswapV3Factory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3FactoryPoolCreatedIterator{contract: _UniswapV3Factory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_UniswapV3Factory *UniswapV3FactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *UniswapV3FactoryPoolCreated, token0 []common.Address, token1 []common.Address, fee []*big.Int) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _UniswapV3Factory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3FactoryPoolCreated)
				if err := _UniswapV3Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_UniswapV3Factory *UniswapV3FactoryFilterer) ParsePoolCreated(log types.Log) (*UniswapV3FactoryPoolCreated, error) {
	event := new(UniswapV3FactoryPoolCreated)
	if err := _UniswapV3Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolMetaData contains all meta data concerning the UniswapV3Pool contract.
var UniswapV3PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"CollectProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"slot0Tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"liquidityDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBefore\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"DBG_MOD_POS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtStartPrice\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimit\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"slot0sqrtpricex96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"DBG_SWAP_HEAD\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceStartX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceNextX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stepAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stepAmountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amountProcessed\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"exactInput\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobalX128\",\"type\":\"uint256\"}],\"name\":\"DBG_SWAP_LOOP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal0X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal1X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityDiff\",\"type\":\"uint128\"}],\"name\":\"DBG_SWAP_TAIL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"liquidityDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal0X128Before\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal1X128Before\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flippedLower\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flippedUpper\",\"type\":\"bool\"}],\"name\":\"DBG_UPD_POS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextOld\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextNew\",\"type\":\"uint16\"}],\"name\":\"IncreaseObservationCardinalityNext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0New\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1New\",\"type\":\"uint8\"}],\"name\":\"SetFeeProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collectProtocol\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal0X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal1X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"}],\"name\":\"increaseObservationCardinalityNext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidityPerTick\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"observations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityCumulativeX128\",\"type\":\"uint160\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"observe\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"tickCumulatives\",\"type\":\"int56[]\"},{\"internalType\":\"uint160[]\",\"name\":\"secondsPerLiquidityCumulativeX128s\",\"type\":\"uint160[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"token0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"token1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"feeProtocol0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol1\",\"type\":\"uint8\"}],\"name\":\"setFeeProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slot0\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"snapshotCumulativesInside\",\"outputs\":[{\"internalType\":\"int56\",\"name\":\"tickCumulativeInside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityInsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsInside\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountSpecified\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"name\":\"tickBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside0X128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside1X128\",\"type\":\"uint256\"},{\"internalType\":\"int56\",\"name\":\"tickCumulativeOutside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityOutsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsOutside\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101606040523480156200001257600080fd5b503060601b60805260408051630890357360e41b81529051600091339163890357309160048082019260a092909190829003018186803b1580156200005657600080fd5b505afa1580156200006b573d6000803e3d6000fd5b505050506040513d60a08110156200008257600080fd5b508051602080830151604084015160608086015160809096015160e896871b6001600160e81b0319166101005291811b6001600160601b031990811660e05292811b831660c0529390931b1660a052600282810b900b90921b610120529150620000f79082906200010f811b62002e1217901c565b60801b6001600160801b03191661014052506200017d565b60008082600281900b620d89e719816200012557fe5b05029050600083600281900b620d89e8816200013d57fe5b0502905060008460020b83830360020b816200015557fe5b0560010190508062ffffff166001600160801b038016816200017357fe5b0495945050505050565b60805160601c60a05160601c60c05160601c60e05160601c6101005160e81c6101205160e81c6101405160801c61630d6200024a600039806121e852806152195280615250525080610c305280612bce528061528452806152b6525080610d1f5280611bc55280611bfc5280612c165250806112ed5280611c7f52806120ee52806126b95280612bf2528061423452508061085a528061141b5280611c4e52806120885280612633528061407152508061227552806124c05280612baa525080612e89525061630d6000f3fe608060405234801561001057600080fd5b506004361061013e5760003560e01c80630dfe168114610143578063128acb08146101675780631a686502146102145780631ad8b03b14610238578063252c09d71461026f57806332148f67146102c65780633850c7bd146102e95780633c8a7d8d1461034257806346141319146103e2578063490e6cbc146103fc5780634f1eb3d814610486578063514ea4bf146104d75780635339c2961461053057806370cf754a146105505780638206a4d11461055857806385b6672914610580578063883bdbfd146105bd578063a34123a7146106c4578063a38807f2146106fe578063c45a015514610759578063d0c93a7c14610761578063d21220a714610780578063ddca3f4314610788578063f3058399146107a8578063f30dba93146107b0578063f637731d14610832575b600080fd5b61014b610858565b604080516001600160a01b039092168252519081900360200190f35b6101fb600480360360a081101561017d57600080fd5b6001600160a01b0382358116926020810135151592604082013592606083013516919081019060a081016080820135600160201b8111156101bd57600080fd5b8201836020820111156101cf57600080fd5b803590602001918460018302840111600160201b831117156101f057600080fd5b50909250905061087c565b6040805192835260208301919091528051918290030190f35b61021c611660565b604080516001600160801b039092168252519081900360200190f35b61024061166f565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b61028c6004803603602081101561028557600080fd5b5035611689565b6040805163ffffffff909516855260069390930b60208501526001600160a01b039091168383015215156060830152519081900360800190f35b6102e7600480360360208110156102dc57600080fd5b503561ffff166116ce565b005b6102f16117c8565b604080516001600160a01b03909816885260029690960b602088015261ffff9485168787015292841660608701529216608085015260ff90911660a0840152151560c0830152519081900360e00190f35b6101fb600480360360a081101561035857600080fd5b6001600160a01b03823516916020810135600290810b92604083013590910b916001600160801b036060820135169181019060a081016080820135600160201b8111156103a457600080fd5b8201836020820111156103b657600080fd5b803590602001918460018302840111600160201b831117156103d757600080fd5b509092509050611818565b6103ea611b1c565b60408051918252519081900360200190f35b6102e76004803603608081101561041257600080fd5b6001600160a01b038235169160208101359160408201359190810190608081016060820135600160201b81111561044857600080fd5b82018360208201111561045a57600080fd5b803590602001918460018302840111600160201b8311171561047b57600080fd5b509092509050611b22565b610240600480360360a081101561049c57600080fd5b506001600160a01b03813516906020810135600290810b91604081013590910b906001600160801b0360608201358116916080013516611f7d565b6104f4600480360360208110156104ed57600080fd5b5035612197565b604080516001600160801b0396871681526020810195909552848101939093529084166060840152909216608082015290519081900360a00190f35b6103ea6004803603602081101561054657600080fd5b503560010b6121d4565b61021c6121e6565b6102e76004803603604081101561056e57600080fd5b5060ff8135811691602001351661220a565b6102406004803603606081101561059657600080fd5b506001600160a01b03813516906001600160801b0360208201358116916040013516612452565b61062b600480360360208110156105d357600080fd5b810190602081018135600160201b8111156105ed57600080fd5b8201836020820111156105ff57600080fd5b803590602001918460208302840111600160201b8311171561062057600080fd5b509092509050612751565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561066f578181015183820152602001610657565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156106ae578181015183820152602001610696565b5050505090500194505050505060405180910390f35b6101fb600480360360608110156106da57600080fd5b508035600290810b91602081013590910b90604001356001600160801b03166127de565b6107286004803603604081101561071457600080fd5b508035600290810b9160200135900b612955565b6040805160069490940b84526001600160a01b03909216602084015263ffffffff1682820152519081900360600190f35b61014b612ba8565b610769612bcc565b6040805160029290920b8252519081900360200190f35b61014b612bf0565b610790612c14565b6040805162ffffff9092168252519081900360200190f35b6103ea612c38565b6107d0600480360360208110156107c657600080fd5b503560020b612c3e565b604080516001600160801b039099168952600f9790970b602089015287870195909552606087019390935260069190910b60808601526001600160a01b031660a085015263ffffffff1660c0840152151560e083015251908190036101000190f35b6102e76004803603602081101561084857600080fd5b50356001600160a01b0316612ca8565b7f000000000000000000000000000000000000000000000000000000000000000081565b600080610887612e7e565b856108be576040805162461bcd60e51b8152602060048201526002602482015261415360f01b604482015290519081900360640190fd5b6040805160e0810182526000546001600160a01b0381168252600160a01b8104600290810b810b900b602083015261ffff600160b81b8204811693830193909352600160c81b810483166060830152600160d81b8104909216608082015260ff600160e81b8304811660a0830152600160f01b909204909116151560c08201819052610977576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b876109c25780600001516001600160a01b0316866001600160a01b03161180156109bd575073fffd8963efd1fc6a506488495d951d5263988d266001600160a01b038716105b6109f4565b80600001516001600160a01b0316866001600160a01b03161080156109f457506401000276a36001600160a01b038716115b610a2b576040805162461bcd60e51b815260206004820152600360248201526214d41360ea1b604482015290519081900360640190fd5b6000805460ff60f01b191681556040805160c08101909152808a610a5a5760048460a0015160ff16901c610a6d565b60108460a0015160ff1681610a6b57fe5b065b60ff1681526004546001600160801b03166020820152604001610a8e612ee7565b63ffffffff168152602001600060060b815260200160006001600160a01b031681526020016000151581525090506000808913905060006040518060e001604052808b81526020016000815260200185600001516001600160a01b03168152602001856020015160020b81526020018c610b0a57600254610b0e565b6001545b815260200160006001600160801b0316815260200184602001516001600160801b03168152509050306001600160a01b03167f884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c2078c8c87600001518d6000800160009054906101000a90046001600160a01b03168760600151604051808715158152602001868152602001856001600160a01b03168152602001846001600160a01b03168152602001836001600160a01b031681526020018260020b8152602001965050505050505060405180910390a25b805115801590610c055750886001600160a01b031681604001516001600160a01b031614155b156110df57610c12615dd8565b60408201516001600160a01b031681526060820151610c55906006907f00000000000000000000000000000000000000000000000000000000000000008f612eeb565b15156040830152600290810b810b60208301819052620d89e719910b1215610c8657620d89e7196020820152610ca5565b6020810151620d89e860029190910b1315610ca557620d89e860208201525b610cb2816020015161302d565b6001600160a01b031660608201526040820151610d43908d610cec578b6001600160a01b031683606001516001600160a01b031611610d06565b8b6001600160a01b031683606001516001600160a01b0316105b610d14578260600151610d16565b8b5b60c085015185517f0000000000000000000000000000000000000000000000000000000000000000613354565b60c085015260a084015260808301526001600160a01b031660408301528215610da557610d798160c00151826080015101613546565b825103825260a0810151610d9b90610d9090613546565b6020840151906135a3565b6020830152610de0565b610db28160a00151613546565b825101825260c08101516080820151610dda91610dcf9101613546565b6020840151906135f1565b60208301525b835160ff1615610e26576000846000015160ff168260c0015181610e0057fe5b60c0840180519290910491829003905260a0840180519091016001600160801b03169052505b60c08201516001600160801b031615610e6557610e598160c00151600160801b8460c001516001600160801b0316613639565b60808301805190910190525b306001600160a01b03167ff5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d592983604001518360000151846060015186606001518960600151876040015188608001518960a001518b602001518b60c00151600460009054906101000a90046001600160801b03168f8f60800151604051808e6001600160a01b031681526020018d6001600160a01b031681526020018c6001600160a01b031681526020018b60020b81526020018a60060b81526020018915158152602001888152602001878152602001868152602001858152602001846001600160801b0316815260200183151581526020018281526020019d505050505050505050505050505060405180910390a280606001516001600160a01b031682604001516001600160a01b0316141561109e57806040015115611075578360a00151610fff57610fdd846040015160008760200151886040015188602001518a606001516008613763909695949392919063ffffffff16565b6001600160a01b03166080860152600690810b900b6060850152600160a08501525b600061104b82602001518e6110165760015461101c565b84608001515b8f61102b57856080015161102f565b6002545b608089015160608a015160408b015160059594939291906138f5565b90508c15611057576000035b6110658360c00151826139af565b6001600160801b031660c0840152505b8b61108457806020015161108d565b60018160200151035b600290810b900b60608301526110d9565b80600001516001600160a01b031682604001516001600160a01b0316146110d9576110cc8260400151613a65565b600290810b900b60608301525b50610bdf565b836020015160020b816060015160020b146111ad5760008061112d86604001518660400151886020015188602001518a606001518b608001516008613d50909695949392919063ffffffff16565b604085015160608601516000805461ffff60c81b1916600160c81b61ffff958616021761ffff60b81b1916600160b81b95909416949094029290921762ffffff60a01b1916600160a01b62ffffff60029490940b9390931692909202919091176001600160a01b0319166001600160a01b03909116179055506111d29050565b6040810151600080546001600160a01b0319166001600160a01b039092169190911790555b8060c001516001600160801b031683602001516001600160801b0316146112185760c0810151600480546001600160801b0319166001600160801b039092169190911790555b8a1561126857608081015160015560a08101516001600160801b0316156112635760a0810151600380546001600160801b031981166001600160801b03918216909301169190911790555b6112ae565b608081015160025560a08101516001600160801b0316156112ae5760a0810151600380546001600160801b03808216600160801b92839004821690940116029190911790555b8115158b1515146112c757602081015181518b036112d4565b80600001518a0381602001515b90965094508a1561140d576000851215611316576113167f00000000000000000000000000000000000000000000000000000000000000008d87600003613ed5565b6000611320614023565b9050336001600160a01b031663fa461e3388888c8c6040518563ffffffff1660e01b815260040180858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b1580156113a457600080fd5b505af11580156113b8573d6000803e3d6000fd5b505050506113c4614023565b6113ce828961418e565b1115611407576040805162461bcd60e51b815260206004820152600360248201526249494160e81b604482015290519081900360640190fd5b50611537565b6000861215611444576114447f00000000000000000000000000000000000000000000000000000000000000008d88600003613ed5565b600061144e6141e6565b9050336001600160a01b031663fa461e3388888c8c6040518563ffffffff1660e01b815260040180858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b1580156114d257600080fd5b505af11580156114e6573d6000803e3d6000fd5b505050506114f26141e6565b6114fc828861418e565b1115611535576040805162461bcd60e51b815260206004820152600360248201526249494160e81b604482015290519081900360640190fd5b505b6001546002805460608481015160005460c08088015160208b810151604080519a8b52918a01979097529390960b878401526001600160a01b03909116928601929092526001600160801b0380851660808701529290930390911660a0840152905130927f507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a58492908290030190a260408082015160c083015160608085015184518b8152602081018b90526001600160a01b03948516818701526001600160801b039093169183019190915260020b60808201529151908e169133917fc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca679181900360a00190a350506000805460ff60f01b1916600160f01b17905550919890975095505050505050565b6004546001600160801b031681565b6003546001600160801b0380821691600160801b90041682565b60088161ffff811061169a57600080fd5b015463ffffffff81169150600160201b810460060b90600160581b81046001600160a01b031690600160f81b900460ff1684565b600054600160f01b900460ff16611712576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b19169055611727612e7e565b60008054600160d81b900461ffff169061174360088385614333565b6000805461ffff808416600160d81b810261ffff60d81b19909316929092179092559192508316146117b0576040805161ffff80851682528316602082015281517fac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a929181900390910190a15b50506000805460ff60f01b1916600160f01b17905550565b6000546001600160a01b03811690600160a01b810460020b9061ffff600160b81b8204811691600160c81b8104821691600160d81b8204169060ff600160e81b8204811691600160f01b90041687565b600080548190600160f01b900460ff1661185f576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b191690556001600160801b0385166118c7576040805162461bcd60e51b815260206004820152601e60248201527f556e69737761705633506f6f6c3a6d696e7428293a20616d6f756e743e300000604482015290519081900360640190fd5b60008061191560405180608001604052808c6001600160a01b031681526020018b60020b81526020018a60020b815260200161190b8a6001600160801b03166143d6565b600f0b9052614419565b9250925050819350809250600080600086111561193757611934614023565b91505b8415611948576119456141e6565b90505b336001600160a01b031663d348799787878b8b6040518563ffffffff1660e01b815260040180858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b1580156119ca57600080fd5b505af11580156119de573d6000803e3d6000fd5b505050506000861115611a35576119f3614023565b6119fd838861418e565b1115611a35576040805162461bcd60e51b815260206004820152600260248201526104d360f41b604482015290519081900360640190fd5b8415611a8557611a436141e6565b611a4d828761418e565b1115611a85576040805162461bcd60e51b81526020600482015260026024820152614d3160f01b604482015290519081900360640190fd5b8960020b8b60020b8d6001600160a01b03167f7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde338d8b8b60405180856001600160a01b03168152602001846001600160801b0316815260200183815260200182815260200194505050505060405180910390a450506000805460ff60f01b1916600160f01b17905550919890975095505050505050565b60025481565b600054600160f01b900460ff16611b66576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b19169055611b7b612e7e565b6004546001600160801b031680611bbd576040805162461bcd60e51b81526020600482015260016024820152601360fa1b604482015290519081900360640190fd5b6000611bf2867f000000000000000000000000000000000000000000000000000000000000000062ffffff16620f4240614703565b90506000611c29867f000000000000000000000000000000000000000000000000000000000000000062ffffff16620f4240614703565b90506000611c35614023565b90506000611c416141e6565b90508815611c7457611c747f00000000000000000000000000000000000000000000000000000000000000008b8b613ed5565b8715611ca557611ca57f00000000000000000000000000000000000000000000000000000000000000008b8a613ed5565b336001600160a01b031663e9cbafb085858a8a6040518563ffffffff1660e01b815260040180858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b158015611d2757600080fd5b505af1158015611d3b573d6000803e3d6000fd5b505050506000611d49614023565b90506000611d556141e6565b905081611d62858861418e565b1115611d9a576040805162461bcd60e51b8152602060048201526002602482015261046360f41b604482015290519081900360640190fd5b80611da5848761418e565b1115611ddd576040805162461bcd60e51b8152602060048201526002602482015261463160f01b604482015290519081900360640190fd5b8382038382038115611e6c5760008054600160e81b9004600f16908115611e10578160ff168481611e0a57fe5b04611e13565b60005b90506001600160801b03811615611e4657600380546001600160801b038082168401166001600160801b03199091161790555b611e60818503600160801b8d6001600160801b0316613639565b60018054909101905550505b8015611ef75760008054600160e81b900460041c600f16908115611e9c578160ff168381611e9657fe5b04611e9f565b60005b90506001600160801b03811615611ed157600380546001600160801b03600160801b8083048216850182160291161790555b611eeb818403600160801b8d6001600160801b0316613639565b60028054909101905550505b8d6001600160a01b0316336001600160a01b03167fbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca6338f8f86866040518085815260200184815260200183815260200182815260200194505050505060405180910390a350506000805460ff60f01b1916600160f01b179055505050505050505050505050565b600080548190600160f01b900460ff16611fc4576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b19168155611fde600733898961476f565b60038101549091506001600160801b0390811690861611611fff578461200e565b60038101546001600160801b03165b60038201549093506001600160801b03600160801b909104811690851611612036578361204c565b6003810154600160801b90046001600160801b03165b91506001600160801b038316156120b1576003810180546001600160801b031981166001600160801b039182168690038216179091556120b1907f0000000000000000000000000000000000000000000000000000000000000000908a908616613ed5565b6001600160801b03821615612117576003810180546001600160801b03600160801b808304821686900382160291811691909117909155612117907f0000000000000000000000000000000000000000000000000000000000000000908a908516613ed5565b604080516001600160a01b038a1681526001600160801b0380861660208301528416818301529051600288810b92908a900b9133917f70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0919081900360600190a4506000805460ff60f01b1916600160f01b17905590969095509350505050565b60076020526000908152604090208054600182015460028301546003909301546001600160801b0392831693919281811691600160801b90041685565b60066020526000908152604090205481565b7f000000000000000000000000000000000000000000000000000000000000000081565b600054600160f01b900460ff1661224e576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b1916905560408051638da5cb5b60e01b815290516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691638da5cb5b916004808301926020929190829003018186803b1580156122bb57600080fd5b505afa1580156122cf573d6000803e3d6000fd5b505050506040513d60208110156122e557600080fd5b50516001600160a01b0316331461232d5760405162461bcd60e51b81526004018080602001828103825260508152602001806162576050913960600191505060405180910390fd5b60ff82161580612350575060048260ff16101580156123505750600a8260ff1611155b801561237a575060ff8116158061237a575060048160ff161015801561237a5750600a8160ff1611155b6123b55760405162461bcd60e51b81526004018080602001828103825260268152602001806161d56026913960400191505060405180910390fd5b60008054610ff0600484901b16840160ff908116600160e81b90810260ff60e81b19841617909355919004167f973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b1336010826040805160ff9390920683168252600f600486901c16602083015286831682820152918516606082015290519081900360800190a150506000805460ff60f01b1916600160f01b17905550565b600080548190600160f01b900460ff16612499576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b1916905560408051638da5cb5b60e01b815290516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691638da5cb5b916004808301926020929190829003018186803b15801561250657600080fd5b505afa15801561251a573d6000803e3d6000fd5b505050506040513d602081101561253057600080fd5b50516001600160a01b031633146125785760405162461bcd60e51b81526004018080602001828103825260508152602001806162576050913960600191505060405180910390fd5b6003546001600160801b039081169085161161259457836125a1565b6003546001600160801b03165b6003549092506001600160801b03600160801b9091048116908416116125c757826125db565b600354600160801b90046001600160801b03165b90506001600160801b0382161561265c576003546001600160801b038381169116141561260a57600019909101905b600380546001600160801b031981166001600160801b0391821685900382161790915561265c907f00000000000000000000000000000000000000000000000000000000000000009087908516613ed5565b6001600160801b038116156126e2576003546001600160801b03828116600160801b90920416141561268d57600019015b600380546001600160801b03600160801b8083048216859003821602918116919091179091556126e2907f00000000000000000000000000000000000000000000000000000000000000009087908416613ed5565b604080516001600160801b0380851682528316602082015281516001600160a01b0388169233927f596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151929081900390910190a36000805460ff60f01b1916600160f01b1790559094909350915050565b60608061275c612e7e565b6127d3612767612ee7565b858580806020026020016040519081016040528093929190818152602001838360200280828437600092018290525054600454600896959450600160a01b820460020b935061ffff600160b81b8304811693506001600160801b0390911691600160c81b9004166147ce565b915091509250929050565b600080548190600160f01b900460ff16612825576040805162461bcd60e51b81526020600482015260036024820152624c4f4b60e81b604482015290519081900360640190fd5b6000805460ff60f01b1916815560408051608081018252338152600288810b602083015287900b918101919091528190819061287e90606081016128716001600160801b038a166143d6565b600003600f0b9052614419565b925092509250816000039450806000039350600085118061289f5750600084115b156128de576003830180546001600160801b038082168089018216600160801b93849004831689019092169092029091176001600160801b0319161790555b604080516001600160801b0388168152602081018790528082018690529051600289810b92908b900b9133917f0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c919081900360600190a450506000805460ff60f01b1916600160f01b179055509094909350915050565b6000806000612962612e7e565b61296c8585614926565b600285810b810b60009081526005602052604080822087840b90930b825281206003830154600681900b93600160381b82046001600160a01b0316928492600160d81b810463ffffffff169284929091600160f81b900460ff1680612a025760405162461bcd60e51b815260040180806020018281038252603a815260200180615f4f603a913960400191505060405180910390fd5b6003820154600681900b9850600160381b81046001600160a01b03169650600160d81b810463ffffffff169450600160f81b900460ff1680612a755760405162461bcd60e51b815260040180806020018281038252603a81526020018061621d603a913960400191505060405180910390fd5b50506040805160e0810182526000546001600160a01b0381168252600160a01b8104600290810b810b810b6020840181905261ffff600160b81b8404811695850195909552600160c81b830485166060850152600160d81b8304909416608084015260ff600160e81b8304811660a0850152600160f01b909204909116151560c08301529093508e810b91900b12159050612b1e57509390940396509003935090039050612ba1565b8a60020b816020015160020b1215612b92576000612b3a612ee7565b6020830151604084015160045460608601519394506000938493612b70936008938893879392916001600160801b031690613763565b9a9003989098039b505094909603929092039650909103039250612ba1915050565b50949093039650039350900390505b9250925092565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60015481565b60056020526000908152604090208054600182015460028301546003909301546001600160801b03831693600160801b909304600f0b9290600681900b90600160381b81046001600160a01b031690600160d81b810463ffffffff1690600160f81b900460ff1688565b6000546001600160a01b031615612ceb576040805162461bcd60e51b8152602060048201526002602482015261414960f01b604482015290519081900360640190fd5b6000612cf682613a65565b9050600080612d0e612d06612ee7565b6008906149ef565b6040805160e0810182526001600160a01b038816808252600288810b6020808501829052600085870181905261ffff898116606088018190529089166080880181905260a08801839052600160c0909801979097528154600160f01b6001600160a01b0319909116871762ffffff60a01b1916600160a01b62ffffff9787900b97909716969096029590951763ffffffff60b81b1916600160c81b9091021761ffff60d81b1916600160d81b9096029590951761ffff60e81b191692909217909355835191825281019190915281519395509193507f98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c9592918290030190a150505050565b60008082600281900b620d89e71981612e2757fe5b05029050600083600281900b620d89e881612e3e57fe5b0502905060008460020b83830360020b81612e5557fe5b0560010190508062ffffff166001600160801b03801681612e7257fe5b0493505050505b919050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612ee55760405162461bcd60e51b815260040180806020018281038252603d815260200180616038603d913960400191505060405180910390fd5b565b4290565b60008060008460020b8660020b81612eff57fe5b05905060008660020b128015612f2657508460020b8660020b81612f1f57fe5b0760020b15155b15612f3057600019015b8315612fa557600080612f4283614a3b565b600182810b810b600090815260208d9052604090205460ff83169190911b80016000190190811680151597509294509092509085612f8757888360ff16860302612f9a565b88612f9182614a4d565b840360ff168603025b965050505050613023565b600080612fb483600101614a3b565b91509150600060018260ff166001901b031990506000818b60008660010b60010b815260200190815260200160002054169050806000141595508561300657888360ff0360ff1686600101010261301c565b888361301183614b19565b0360ff168660010101025b9650505050505b5094509492505050565b60008060008360020b12613044578260020b61304c565b8260020b6000035b9050620d89e881111561308a576040805162461bcd60e51b81526020600482015260016024820152601560fa1b604482015290519081900360640190fd5b60006001821661309e57600160801b6130b0565b6ffffcb933bd6fad37aa2d162d1a5940015b6001600160881b0316905060028216156130da576ffff97272373d413259a46990580e213a0260801c5b60048216156130f9576ffff2e50f5f656932ef12357cf3c7fdcc0260801c5b6008821615613118576fffe5caca7e10e4e61c3624eaa0941cd00260801c5b6010821615613137576fffcb9843d60f6159c9db58835c9266440260801c5b6020821615613156576fff973b41fa98c081472e6896dfb254c00260801c5b6040821615613175576fff2ea16466c96a3843ec78b326b528610260801c5b6080821615613194576ffe5dee046a99a2a811c461f1969c30530260801c5b6101008216156131b4576ffcbe86c7900a88aedcffc83b479aa3a40260801c5b6102008216156131d4576ff987a7253ac413176f2b074cf7815e540260801c5b6104008216156131f4576ff3392b0822b70005940c7a398e4b70f30260801c5b610800821615613214576fe7159475a2c29b7443b29c7fa6e889d90260801c5b611000821615613234576fd097f3bdfd2022b8845ad8f792aa58250260801c5b612000821615613254576fa9f746462d870fdf8a65dc1f90e061e50260801c5b614000821615613274576f70d869a156d2a1b890bb3df62baf32f70260801c5b618000821615613294576f31be135f97d08fd981231505542fcfa60260801c5b620100008216156132b5576f09aa508b5b7a84e1c677de54f3e99bc90260801c5b620200008216156132d5576e5d6af8dedb81196699c329225ee6040260801c5b620400008216156132f4576d2216e584f5fa1ea926041bedfe980260801c5b62080000821615613311576b048a170391f7dc42444e8fa20260801c5b60008460020b131561332c57806000198161332857fe5b0490505b600160201b81061561333f576001613342565b60005b60ff16602082901c0192505050919050565b60008080806001600160a01b03808916908a1610158187128015906133d957600061338d8989620f42400362ffffff16620f4240613639565b9050826133a6576133a18c8c8c6001614c34565b6133b3565b6133b38b8d8c6001614caf565b95508581106133c4578a96506133d3565b6133d08c8b8386614d8c565b96505b50613423565b816133f0576133eb8b8b8b6000614caf565b6133fd565b6133fd8a8c8b6000614c34565b935083886000031061341157899550613423565b6134208b8a8a60000385614e3c565b95505b6001600160a01b038a8116908716148215613486578080156134425750815b61345857613453878d8c6001614caf565b61345a565b855b9550808015613467575081155b61347d57613478878d8c6000614c34565b61347f565b845b94506134d0565b8080156134905750815b6134a6576134a18c888c6001614c34565b6134a8565b855b95508080156134b5575081155b6134cb576134c68c888c6000614caf565b6134cd565b845b94505b811580156134e057508860000385115b156134ec578860000394505b81801561350b57508a6001600160a01b0316876001600160a01b031614155b1561351a578589039350613537565b613534868962ffffff168a620f42400362ffffff16614703565b93505b50505095509550955095915050565b6000600160ff1b821061359f576040805162461bcd60e51b815260206004820152601c60248201527b73616665636173743a746f496e7432353628293a20793c322a32353560201b604482015290519081900360640190fd5b5090565b808203828113156000831215146135eb5760405162461bcd60e51b81526004018080602001828103825260288152602001806160756028913960400191505060405180910390fd5b92915050565b818101828112156000831215146135eb5760405162461bcd60e51b81526004018080602001828103825260288152602001806160106028913960400191505060405180910390fd5b60008080600019858709868602925082811090839003039050806136b757600084116136ac576040805162461bcd60e51b815260206004820181905260248201527f66756c6c6d6174683a6d756c44697628293a2064656e6f6d696e61746f723e30604482015290519081900360640190fd5b50829004905061375c565b8084116136f55760405162461bcd60e51b815260040180806020018281038252602481526020018061609d6024913960400191505060405180910390fd5b6000848688096000868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b60008063ffffffff8716613809576000898661ffff1661ffff811061378457fe5b60408051608081018252919092015463ffffffff808216808452600160201b8304600690810b810b900b6020850152600160581b83046001600160a01b031694840194909452600160f81b90910460ff16151560608301529092508a16146137f5576137f2818a8988614eec565b90505b8060200151816040015192509250506138e9565b86880360008061381e8c8c858c8c8c8c614f8f565b91509150816000015163ffffffff168363ffffffff1614156138505781602001518260400151945094505050506138e9565b805163ffffffff848116911614156138785780602001518160400151945094505050506138e9565b8151815160208085015190840151918390039286039163ffffffff80841692908516910360060b816138a657fe5b05028460200151018263ffffffff168263ffffffff1686604001518660400151036001600160a01b031602816138d857fe5b048560400151019650965050505050505b97509795505050505050565b600295860b860b60009081526020979097526040909620600181018054909503909455938301805490920390915560038201805463ffffffff600160d81b6001600160a01b03600160381b808504821690960316909402600160381b600160d81b031990921691909117600681810b90960390950b66ffffffffffffff1666ffffffffffffff199095169490941782810485169095039093160263ffffffff60d81b1990931692909217905554600160801b9004600f0b90565b60008082600f0b1215613a1457826001600160801b03168260000384039150816001600160801b031610613a0f576040805162461bcd60e51b81526020600482015260026024820152614c5360f01b604482015290519081900360640190fd5b6135eb565b826001600160801b03168284019150816001600160801b031610156135eb576040805162461bcd60e51b81526020600482015260026024820152614c4160f01b604482015290519081900360640190fd5b60006401000276a36001600160a01b03831610801590613aa1575073fffd8963efd1fc6a506488495d951d5263988d266001600160a01b038316105b613ad6576040805162461bcd60e51b81526020600482015260016024820152602960f91b604482015290519081900360640190fd5b600160201b600160c01b03602083901b166001600160801b03811160071b81811c6001600160401b03811160061b90811c63ffffffff811160051b90811c61ffff811160041b90811c60ff8111600390811b91821c600f811160021b90811c918211600190811b92831c97908811961790941790921717909117171760808110613b6857607f810383901c9150613b72565b80607f0383901b91505b908002607f81811c60ff83811c9190911c800280831c81831c1c800280841c81841c1c800280851c81851c1c800280861c81861c1c800280871c81871c1c800280881c81881c1c800280891c81891c1c8002808a1c818a1c1c8002808b1c818b1c1c8002808c1c818c1c1c8002808d1c818d1c1c8002808e1c9c81901c9c909c1c80029c8d901c9e9d607f198f0160401b60c09190911c6001603f1b161760c19b909b1c6001603e1b169a909a1760c29990991c6001603d1b169890981760c39790971c6001603c1b169690961760c49590951c6001603b1b169490941760c59390931c6001603a1b169290921760c69190911c600160391b161760c79190911c600160381b161760c89190911c600160371b161760c99190911c600160361b161760ca9190911c600160351b161760cb9190911c600160341b161760cc9190911c600160331b161760cd9190911c600160321b1617693627a301d71055774c8581026f028f6481ab7f045a5af012a19d003aa9198101608090811d906fdb2df09e81959a81455e260799a0632f8301901d600281810b9083900b14613d4157886001600160a01b0316613d258261302d565b6001600160a01b03161115613d3a5781613d3c565b805b613d43565b815b9998505050505050505050565b6000806000898961ffff1661ffff8110613d6657fe5b60408051608081018252919092015463ffffffff808216808452600160201b8304600690810b810b900b6020850152600160581b83046001600160a01b031694840194909452600160f81b90910460ff161515606083015290925089161415613dd557888592509250506138e9565b8461ffff168461ffff16118015613df657506001850361ffff168961ffff16145b15613e0357839150613e07565b8491505b8161ffff168960010161ffff1681613e1b57fe5b069250613e2a81898989614eec565b8a8461ffff1661ffff8110613e3b57fe5b825191018054602084015160408501516060909501511515600160f81b026001600160f81b036001600160a01b03909616600160581b02600160581b600160f81b031960069390930b66ffffffffffffff16600160201b0266ffffffffffffff60201b1963ffffffff90971663ffffffff199095169490941795909516929092171692909217929092161790555097509795505050505050565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b1781529251825160009485949389169392918291908083835b60208310613f515780518252601f199092019160209182019101613f32565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114613fb3576040519150601f19603f3d011682016040523d82523d6000602084013e613fb8565b606091505b5091509150818015613fe6575080511580613fe65750808060200190516020811015613fe357600080fd5b50515b61401c576040805162461bcd60e51b81526020600482015260026024820152612a2360f11b604482015290519081900360640190fd5b5050505050565b604080513060248083019190915282518083039091018152604490910182526020810180516001600160e01b03166370a0823160e01b17815291518151600093849384936001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001693919290918291908083835b602083106140bc5780518252601f19909201916020918201910161409d565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d806000811461411c576040519150601f19603f3d011682016040523d82523d6000602084013e614121565b606091505b509150915081801561413557506020815110155b6141705760405162461bcd60e51b8152600401808060200182810382526031815260200180615eb26031913960400191505060405180910390fd5b80806020019051602081101561418557600080fd5b50519250505090565b808201828110156135eb576040805162461bcd60e51b815260206004820152601f60248201527f6c6f77676173736166656d6174683a616464282920287a3d782b79293e3d7800604482015290519081900360640190fd5b604080513060248083019190915282518083039091018152604490910182526020810180516001600160e01b03166370a0823160e01b17815291518151600093849384936001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001693919290918291908083835b6020831061427f5780518252601f199092019160209182019101614260565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d80600081146142df576040519150601f19603f3d011682016040523d82523d6000602084013e6142e4565b606091505b50915091508180156142f857506020815110155b6141705760405162461bcd60e51b81526004018080602001828103825260318152602001806162a76031913960400191505060405180910390fd5b6000808361ffff1611614371576040805162461bcd60e51b81526020600482015260016024820152604960f81b604482015290519081900360640190fd5b8261ffff168261ffff161161438757508161375c565b825b8261ffff168161ffff1610156143cd576001858261ffff1661ffff81106143ac57fe5b01805463ffffffff191663ffffffff92909216919091179055600101614389565b50909392505050565b80600f81900b8114612e795760405162461bcd60e51b81526004018080602001828103825260228152602001806161fb6022913960400191505060405180910390fd5b6000806000614426612e7e565b61443884602001518560400151614926565b6040805160e0810182526000546001600160a01b0381168252600160a01b8104600290810b810b900b602080840182905261ffff600160b81b8404811685870152600160c81b84048116606080870191909152600160d81b8504909116608086015260ff600160e81b8504811660a0870152600160f01b909404909316151560c0850152885190890151948901519289015193946144dc9491939092909190615189565b935060008560600151600f0b60001461465657856020015160020b826020015160020b12156145335761452c614515876020015161302d565b614522886040015161302d565b88606001516153cd565b9350614656565b856040015160020b826020015160020b121561462c575060045460408201516001600160801b039091169081906145889061456c612ee7565b6020860151606087015160808801516008949392918791613d50565b6000805461ffff60c81b1916600160c81b61ffff938416021761ffff60b81b1916600160b81b9390921692909202179055825160408801516145d891906145ce9061302d565b89606001516153cd565b94506145f66145ea886020015161302d565b845160608a0151615411565b93506146068188606001516139af565b600480546001600160801b0319166001600160801b039290921691909117905550614656565b61465361463c876020015161302d565b614649886040015161302d565b8860600151615411565b92505b85516020808801516040808a0151868401516060808d0151895185516001600160a01b03998a168152600297880b9881019890985293860b878601529190940b93850193909352600f9290920b60808401526001600160801b03851660a084015260c0830188905260e08301879052921661010082015290517fd1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f918190036101200190a150509193909250565b6000614710848484613639565b90506000828061471c57fe5b848609111561375c5760001981106147655760405162461bcd60e51b81526004018080602001828103825260298152602001806161ac6029913960400191505060405180910390fd5b6001019392505050565b6040805160609490941b6001600160601b031916602080860191909152600293840b60e890811b60348701529290930b90911b60378401528051808403601a018152603a90930181528251928201929092206000908152929052902090565b60608060008361ffff161161480e576040805162461bcd60e51b81526020600482015260016024820152604960f81b604482015290519081900360640190fd5b86516001600160401b038111801561482557600080fd5b5060405190808252806020026020018201604052801561484f578160200160208202803683370190505b50915086516001600160401b038111801561486957600080fd5b50604051908082528060200260200182016040528015614893578160200160208202803683370190505b50905060005b8751811015614919576148c48a8a8a84815181106148b357fe5b60200260200101518a8a8a8a613763565b8483815181106148d057fe5b602002602001018484815181106148e357fe5b60200260200101826001600160a01b03166001600160a01b03168152508260060b60060b81525050508080600101915050614899565b5097509795505050505050565b8060020b8260020b12614966576040805162461bcd60e51b8152602060048201526003602482015262544c5560e81b604482015290519081900360640190fd5b620d89e719600283900b12156149a9576040805162461bcd60e51b8152602060048201526003602482015262544c4d60e81b604482015290519081900360640190fd5b620d89e8600282900b13156149eb576040805162461bcd60e51b815260206004820152600360248201526254554d60e81b604482015290519081900360640190fd5b5050565b6040805160808101825263ffffffff9283168082526000602083018190529282019290925260016060909101819052835463ffffffff1916909117909116600160f81b17909155908190565b60020b600881901d9161010090910790565b6000808211614a8d5760405162461bcd60e51b8152600401808060200182810382526021815260200180615f2e6021913960400191505060405180910390fd5b600160801b8210614aa057608091821c91015b600160401b8210614ab357604091821c91015b600160201b8210614ac657602091821c91015b620100008210614ad857601091821c91015b6101008210614ae957600891821c91015b60108210614af957600491821c91015b60048210614b0957600291821c91015b60028210612e7957600101919050565b6000808211614b595760405162461bcd60e51b8152600401808060200182810382526022815260200180615fbf6022913960400191505060405180910390fd5b5060ff6001600160801b03821615614b7457607f1901614b7c565b608082901c91505b6001600160401b03821615614b9457603f1901614b9c565b604082901c91505b63ffffffff821615614bb157601f1901614bb9565b602082901c91505b61ffff821615614bcc57600f1901614bd4565b601082901c91505b60ff821615614be65760071901614bee565b600882901c91505b600f821615614c005760031901614c08565b600482901c91505b6003821615614c1a5760011901614c22565b600282901c91505b6001821615612e795760001901919050565b6000836001600160a01b0316856001600160a01b03161115614c54579293925b81614c8157614c7c836001600160801b03168686036001600160a01b0316600160601b613639565b614ca4565b614ca4836001600160801b03168686036001600160a01b0316600160601b614703565b90505b949350505050565b6000836001600160a01b0316856001600160a01b03161115614ccf579293925b600160601b600160e01b03606084901b166001600160a01b038686038116908716614d2b5760405162461bcd60e51b815260040180806020018281038252602f815260200180615fe1602f913960400191505060405180910390fd5b83614d5b57866001600160a01b0316614d4e8383896001600160a01b0316613639565b81614d5557fe5b04614d81565b614d81614d728383896001600160a01b0316614703565b886001600160a01b0316615440565b979650505050505050565b600080856001600160a01b031611614dd55760405162461bcd60e51b81526004018080602001828103825260358152602001806161196035913960400191505060405180910390fd5b6000846001600160801b031611614e1d5760405162461bcd60e51b8152600401808060200182810382526036815260200180615f896036913960400191505060405180910390fd5b81614e2f57614c7c858585600161544b565b614ca4858585600161555e565b600080856001600160a01b031611614e855760405162461bcd60e51b81526004018080602001828103825260378152602001806160c16037913960400191505060405180910390fd5b6000846001600160801b031611614ecd5760405162461bcd60e51b81526004018080602001828103825260388152602001806161746038913960400191505060405180910390fd5b81614edf57614c7c858585600061555e565b614ca4858585600061544b565b614ef4615e14565b600085600001518503905060405180608001604052808663ffffffff1681526020018263ffffffff168660020b0288602001510160060b81526020016000856001600160801b031611614f48576001614f4a565b845b6001600160801b031663ffffffff60801b608085901b1681614f6857fe5b048860400151016001600160a01b0316815260200160011515815250915050949350505050565b614f97615e14565b614f9f615e14565b888561ffff1661ffff8110614fb057fe5b60408051608081018252919092015463ffffffff8116808352600160201b8204600690810b810b900b6020840152600160581b82046001600160a01b031693830193909352600160f81b900460ff1615156060820152925061501490899089615673565b1561504c578663ffffffff16826000015163ffffffff161415615036576138e9565b8161504383898988614eec565b915091506138e9565b888361ffff168660010161ffff168161506157fe5b0661ffff1661ffff811061507157fe5b60408051608081018252929091015463ffffffff81168352600160201b8104600690810b810b900b60208401526001600160a01b03600160581b8204169183019190915260ff600160f81b9091041615156060820181905290925061512657604080516080810182528a5463ffffffff81168252600160201b8104600690810b810b900b6020830152600160581b81046001600160a01b031692820192909252600160f81b90910460ff161515606082015291505b61513588836000015189615673565b61516c576040805162461bcd60e51b815260206004820152600360248201526213d31160ea1b604482015290519081900360640190fd5b6151798989898887615734565b9150915097509795505050505050565b6000615198600787878761476f565b60015460025491925090600080600f87900b156152de5760006151b9612ee7565b60008054600454929350909182916152039160089186918591600160a01b810460020b9161ffff600160b81b83048116926001600160801b0390921691600160c81b900416613763565b909250905061523d60058d8b8d8b8b87898b60007f00000000000000000000000000000000000000000000000000000000000000006158d2565b945061527460058c8b8d8b8b87898b60017f00000000000000000000000000000000000000000000000000000000000000006158d2565b935084156152a8576152a860068d7f0000000000000000000000000000000000000000000000000000000000000000615a8b565b83156152da576152da60068c7f0000000000000000000000000000000000000000000000000000000000000000615a8b565b5050505b6000806152f060058c8c8b8a8a615b23565b9092509050615301878a8484615bcf565b600089600f0b121561532f57831561531e5761531e60058c615d64565b821561532f5761532f60058b615d64565b604080516001600160a01b038e16815260028d810b60208301528c810b828401528a900b6060820152600f8b900b608082015260a0810188905260c0810187905260e08101849052610100810183905285151561012082015284151561014082015290517f2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24918190036101600190a150505050505095945050505050565b60008082600f0b126153f3576153ee6153e98585856001614caf565b613546565b614ca7565b6154066153e98585856000036000614caf565b600003949350505050565b60008082600f0b1261542d576153ee6153e98585856001614c34565b6154066153e98585856000036000614c34565b808204910615150190565b600081156154be5760006001600160a01b038411156154815761547c84600160601b876001600160801b0316613639565b615499565b6001600160801b038516606085901b8161549757fe5b045b90506154b66154b16001600160a01b0388168361418e565b615d90565b915050614ca7565b60006001600160a01b038411156154ec576154e784600160601b876001600160801b0316614703565b615503565b615503606085901b6001600160801b038716615440565b905080866001600160a01b03161161554c5760405162461bcd60e51b815260040180806020018281038252604b815260200180615ee3604b913960600191505060405180910390fd5b6001600160a01b038616039050614ca7565b60008261556c575083614ca7565b600160601b600160e01b03606085901b1682156155fa576001600160a01b0386168481029085828161559a57fe5b0414156155cb578181018281106155c9576155bf83896001600160a01b031683614703565b9350505050614ca7565b505b6155f1826155ec878a6001600160a01b031686816155e557fe5b049061418e565b615440565b92505050614ca7565b6001600160a01b0386168481029085828161561157fe5b0414801561561e57508082115b6156595760405162461bcd60e51b8152600401808060200182810382526076815260200180615e3c6076913960800191505060405180910390fd5b8082036155bf6154b1846001600160a01b038b1684614703565b60008363ffffffff168363ffffffff161115801561569d57508363ffffffff168263ffffffff1611155b156156b9578163ffffffff168363ffffffff161115905061375c565b60008463ffffffff168463ffffffff16116156e0578363ffffffff16600160201b016156e8565b8363ffffffff165b64ffffffffff16905060008563ffffffff168463ffffffff1611615718578363ffffffff16600160201b01615720565b8363ffffffff165b64ffffffffff169091111595945050505050565b61573c615e14565b615744615e14565b60008361ffff168560010161ffff168161575a57fe5b0661ffff169050600060018561ffff16830103905060005b506002818301048961ffff8716828161578757fe5b0661ffff811061579357fe5b60408051608081018252929091015463ffffffff81168352600160201b8104600690810b810b900b60208401526001600160a01b03600160581b8204169183019190915260ff600160f81b909104161515606082018190529095506157fd57806001019250615772565b898661ffff16826001018161580e57fe5b0661ffff811061581a57fe5b60408051608081018252929091015463ffffffff81168352600160201b8104600690810b810b900b60208401526001600160a01b03600160581b8204169183019190915260ff600160f81b90910416151560608201528551909450600090615884908b908b615673565b905080801561589d575061589d8a8a8760000151615673565b156158a857506158c5565b806158b8576001820392506158bf565b8160010193505b50615772565b5050509550959350505050565b60028a810b900b600090815260208c90526040812080546001600160801b0316826158fd828d6139af565b9050846001600160801b0316816001600160801b0316111561594b576040805162461bcd60e51b81526020600482015260026024820152614c4f60f01b604482015290519081900360640190fd5b6001600160801b0382811615908216158114159450156159f0578c60020b8e60020b136159d857600183018b9055600283018a9055600383018054600160381b600160d81b031916600160381b6001600160a01b038c16021766ffffffffffffff191666ffffffffffffff60068b900b161763ffffffff60d81b1916600160d81b63ffffffff8a16021790555b6003830180546001600160f81b0316600160f81b1790555b82546001600160801b0319166001600160801b03821617835585615a39578254615a3490615a2f90600160801b9004600f90810b810b908f900b6135f1565b6143d6565b615a5a565b8254615a5a90615a2f90600160801b9004600f90810b810b908f900b6135a3565b8354600f9190910b6001600160801b03908116600160801b0291161790925550909c9b505050505050505050505050565b8060020b8260020b81615a9a57fe5b0760020b15615ada5760405162461bcd60e51b81526004018080602001828103825260218152602001806160f86021913960400191505060405180910390fd5b600080615af58360020b8560020b81615aef57fe5b05614a3b565b600191820b820b60009081526020979097526040909620805460ff9097169190911b90951890945550505050565b600285810b80820b60009081526020899052604080822088850b850b83529082209193849391929184918291908a900b12615b6957505060018201546002830154615b7c565b8360010154880391508360020154870390505b6000808b60020b8b60020b1215615b9e57505060018301546002840154615bb1565b84600101548a0391508460020154890390505b92909803979097039b96909503949094039850939650505050505050565b6040805160a08101825285546001600160801b0390811682526001870154602083015260028701549282019290925260038601548083166060830152600160801b900490911660808201526000600f85900b615c6e5781516001600160801b0316615c66576040805162461bcd60e51b815260206004820152600260248201526104e560f41b604482015290519081900360640190fd5b508051615c7d565b8151615c7a90866139af565b90505b6000615ca18360200151860384600001516001600160801b0316600160801b613639565b90506000615cc78460400151860385600001516001600160801b0316600160801b613639565b905086600f0b600014615cee5787546001600160801b0319166001600160801b0384161788555b60018801869055600288018590556001600160801b038216151580615d1c57506000816001600160801b0316115b15615d5a576003880180546001600160801b031981166001600160801b039182168501821617808216600160801b9182900483168501909216021790555b5050505050505050565b600290810b810b6000908152602092909252604082208281556001810183905590810182905560030155565b806001600160a01b0381168114612e795760405162461bcd60e51b815260040180806020018281038252602681526020018061614e6026913960400191505060405180910390fd5b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c081019190915290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe7371727470726963656d6174683a676574684e65787453717274507269636546726f6d416d6f756e7430526f756e64696e67557028293a202870726f647563743d616d6f756e742a7371727450783936292f616d6f756e743d3d737172745078393626266e756d657261746f72313e70726f64756374556e69737761705633506f6f6c3a62616c616e636530282920737563636573732626646174612e6c656e6774683e3d33327371727470726963656d6174683a6765744e65787453717274507269636546726f6d416d6f756e7431526f756e64696e67446f776e2829207371727450583936203e2071756f7469656e746269746d6174683a6d6f73745369676e6966696361746e42697428293a20783e30556e69737761705633506f6f6c3a736e617073686f7443756d756c6174697665496e7369646528293a20696e697469616c697a65644c6f7765727371727470726963656d6174683a6765744e65787453717274507269636546726f6d496e70757428293a206c69717569646974793e306269746d6174683a6c656173745369676e69666963616e7442697428293a20783e307371727470726963656d6174683a676574416d6f756e743044656c746128293a2073717274526174696f6e415839366c6f77676173736166656d6174683a61646428293a20287a3d782b79293e3d783d3d28793e3d30294e6f44656c656761746543616c6c3a636865636b4e6f7444656c656761746543616c6c282920616464726573732874686973293d3d6f726967696e616c6c6f77676173736166656d6174683a73756228293a20287a3d782d79293c3d783d3d28793e3d302966756c6c6d6174683a6d756c44697628293a2064656e6f6d696e61746f723e70726f64317371727470726963656d6174683a206765744e65787453717274507269636546726f6d4f757470757428293a2073717274505839363e305469636b4269746d61703a666c69705469636b28293a20666c69705469636b28297371727470726963656d6174683a6765744e65787453717274507269636546726f6d496e70757428293a2073717274505839363e3073616665636173743a746f55696e7431363028293a20287a3d75696e74313630287929293d797371727470726963656d6174683a206765744e65787453717274507269636546726f6d4f757470757428293a206c69717569646974793e3066756c6c6d6174683a6d756c446976526f756e64696e67557028293a20726573756c74203c206d6178556e69737761705633506f6f6c3a73657446656550726f746f636f6c3a206c696e652038343073616665636173743a746f496e7431323828293a20287a3d696e7431323828792929556e69737761705633506f6f6c3a736e617073686f7443756d756c6174697665496e7369646528293a20696e697469616c697a65645570706572556e69737761705633506f6f6c3a6f6e6c79466163746f72794f776e657228293a206d73672e73656e6465723d3d49556e69737761705633466163746f727928666163746f7279292e6f776e65722829556e69737761705633506f6f6c3a62616c616e636531282920737563636573732626646174612e6c656e6774683e3d3332a264697066735822122095074e04ee96156b634d25a3aa164b995f001ef9e273c3f07d788621156cd50f64736f6c63430007060033",
}

// UniswapV3PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3PoolMetaData.ABI instead.
var UniswapV3PoolABI = UniswapV3PoolMetaData.ABI

// UniswapV3PoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniswapV3PoolMetaData.Bin instead.
var UniswapV3PoolBin = UniswapV3PoolMetaData.Bin

// DeployUniswapV3Pool deploys a new Ethereum contract, binding an instance of UniswapV3Pool to it.
func DeployUniswapV3Pool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniswapV3Pool, error) {
	parsed, err := UniswapV3PoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniswapV3PoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniswapV3Pool{UniswapV3PoolCaller: UniswapV3PoolCaller{contract: contract}, UniswapV3PoolTransactor: UniswapV3PoolTransactor{contract: contract}, UniswapV3PoolFilterer: UniswapV3PoolFilterer{contract: contract}}, nil
}

// UniswapV3Pool is an auto generated Go binding around an Ethereum contract.
type UniswapV3Pool struct {
	UniswapV3PoolCaller     // Read-only binding to the contract
	UniswapV3PoolTransactor // Write-only binding to the contract
	UniswapV3PoolFilterer   // Log filterer for contract events
}

// UniswapV3PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3PoolSession struct {
	Contract     *UniswapV3Pool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV3PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3PoolCallerSession struct {
	Contract *UniswapV3PoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UniswapV3PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3PoolTransactorSession struct {
	Contract     *UniswapV3PoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UniswapV3PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3PoolRaw struct {
	Contract *UniswapV3Pool // Generic contract binding to access the raw methods on
}

// UniswapV3PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3PoolCallerRaw struct {
	Contract *UniswapV3PoolCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3PoolTransactorRaw struct {
	Contract *UniswapV3PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3Pool creates a new instance of UniswapV3Pool, bound to a specific deployed contract.
func NewUniswapV3Pool(address common.Address, backend bind.ContractBackend) (*UniswapV3Pool, error) {
	contract, err := bindUniswapV3Pool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Pool{UniswapV3PoolCaller: UniswapV3PoolCaller{contract: contract}, UniswapV3PoolTransactor: UniswapV3PoolTransactor{contract: contract}, UniswapV3PoolFilterer: UniswapV3PoolFilterer{contract: contract}}, nil
}

// NewUniswapV3PoolCaller creates a new read-only instance of UniswapV3Pool, bound to a specific deployed contract.
func NewUniswapV3PoolCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3PoolCaller, error) {
	contract, err := bindUniswapV3Pool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolCaller{contract: contract}, nil
}

// NewUniswapV3PoolTransactor creates a new write-only instance of UniswapV3Pool, bound to a specific deployed contract.
func NewUniswapV3PoolTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3PoolTransactor, error) {
	contract, err := bindUniswapV3Pool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolTransactor{contract: contract}, nil
}

// NewUniswapV3PoolFilterer creates a new log filterer instance of UniswapV3Pool, bound to a specific deployed contract.
func NewUniswapV3PoolFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3PoolFilterer, error) {
	contract, err := bindUniswapV3Pool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolFilterer{contract: contract}, nil
}

// bindUniswapV3Pool binds a generic wrapper to an already deployed contract.
func bindUniswapV3Pool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV3PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Pool *UniswapV3PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Pool.Contract.UniswapV3PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Pool *UniswapV3PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.UniswapV3PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Pool *UniswapV3PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.UniswapV3PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Pool *UniswapV3PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Pool *UniswapV3PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Pool *UniswapV3PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV3Pool *UniswapV3PoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV3Pool *UniswapV3PoolSession) Factory() (common.Address, error) {
	return _UniswapV3Pool.Contract.Factory(&_UniswapV3Pool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) Factory() (common.Address, error) {
	return _UniswapV3Pool.Contract.Factory(&_UniswapV3Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_UniswapV3Pool *UniswapV3PoolCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_UniswapV3Pool *UniswapV3PoolSession) Fee() (*big.Int, error) {
	return _UniswapV3Pool.Contract.Fee(&_UniswapV3Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) Fee() (*big.Int, error) {
	return _UniswapV3Pool.Contract.Fee(&_UniswapV3Pool.CallOpts)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_UniswapV3Pool *UniswapV3PoolCaller) FeeGrowthGlobal0X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "feeGrowthGlobal0X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_UniswapV3Pool *UniswapV3PoolSession) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _UniswapV3Pool.Contract.FeeGrowthGlobal0X128(&_UniswapV3Pool.CallOpts)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _UniswapV3Pool.Contract.FeeGrowthGlobal0X128(&_UniswapV3Pool.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_UniswapV3Pool *UniswapV3PoolCaller) FeeGrowthGlobal1X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "feeGrowthGlobal1X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_UniswapV3Pool *UniswapV3PoolSession) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _UniswapV3Pool.Contract.FeeGrowthGlobal1X128(&_UniswapV3Pool.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _UniswapV3Pool.Contract.FeeGrowthGlobal1X128(&_UniswapV3Pool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_UniswapV3Pool *UniswapV3PoolCaller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_UniswapV3Pool *UniswapV3PoolSession) Liquidity() (*big.Int, error) {
	return _UniswapV3Pool.Contract.Liquidity(&_UniswapV3Pool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) Liquidity() (*big.Int, error) {
	return _UniswapV3Pool.Contract.Liquidity(&_UniswapV3Pool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_UniswapV3Pool *UniswapV3PoolCaller) MaxLiquidityPerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "maxLiquidityPerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_UniswapV3Pool *UniswapV3PoolSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _UniswapV3Pool.Contract.MaxLiquidityPerTick(&_UniswapV3Pool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _UniswapV3Pool.Contract.MaxLiquidityPerTick(&_UniswapV3Pool.CallOpts)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_UniswapV3Pool *UniswapV3PoolCaller) Observations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "observations", arg0)

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
// Solidity: function observations(uint256 ) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_UniswapV3Pool *UniswapV3PoolSession) Observations(arg0 *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	return _UniswapV3Pool.Contract.Observations(&_UniswapV3Pool.CallOpts, arg0)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) Observations(arg0 *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	return _UniswapV3Pool.Contract.Observations(&_UniswapV3Pool.CallOpts, arg0)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_UniswapV3Pool *UniswapV3PoolCaller) Observe(opts *bind.CallOpts, secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "observe", secondsAgos)

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
func (_UniswapV3Pool *UniswapV3PoolSession) Observe(secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	return _UniswapV3Pool.Contract.Observe(&_UniswapV3Pool.CallOpts, secondsAgos)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) Observe(secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	return _UniswapV3Pool.Contract.Observe(&_UniswapV3Pool.CallOpts, secondsAgos)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_UniswapV3Pool *UniswapV3PoolCaller) Positions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "positions", arg0)

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
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_UniswapV3Pool *UniswapV3PoolSession) Positions(arg0 [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _UniswapV3Pool.Contract.Positions(&_UniswapV3Pool.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) Positions(arg0 [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _UniswapV3Pool.Contract.Positions(&_UniswapV3Pool.CallOpts, arg0)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_UniswapV3Pool *UniswapV3PoolCaller) ProtocolFees(opts *bind.CallOpts) (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "protocolFees")

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
func (_UniswapV3Pool *UniswapV3PoolSession) ProtocolFees() (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	return _UniswapV3Pool.Contract.ProtocolFees(&_UniswapV3Pool.CallOpts)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) ProtocolFees() (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	return _UniswapV3Pool.Contract.ProtocolFees(&_UniswapV3Pool.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_UniswapV3Pool *UniswapV3PoolCaller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "slot0")

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
func (_UniswapV3Pool *UniswapV3PoolSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _UniswapV3Pool.Contract.Slot0(&_UniswapV3Pool.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _UniswapV3Pool.Contract.Slot0(&_UniswapV3Pool.CallOpts)
}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint32 secondsInside)
func (_UniswapV3Pool *UniswapV3PoolCaller) SnapshotCumulativesInside(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "snapshotCumulativesInside", tickLower, tickUpper)

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
func (_UniswapV3Pool *UniswapV3PoolSession) SnapshotCumulativesInside(tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	return _UniswapV3Pool.Contract.SnapshotCumulativesInside(&_UniswapV3Pool.CallOpts, tickLower, tickUpper)
}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint32 secondsInside)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) SnapshotCumulativesInside(tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	return _UniswapV3Pool.Contract.SnapshotCumulativesInside(&_UniswapV3Pool.CallOpts, tickLower, tickUpper)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 ) view returns(uint256)
func (_UniswapV3Pool *UniswapV3PoolCaller) TickBitmap(opts *bind.CallOpts, arg0 int16) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "tickBitmap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 ) view returns(uint256)
func (_UniswapV3Pool *UniswapV3PoolSession) TickBitmap(arg0 int16) (*big.Int, error) {
	return _UniswapV3Pool.Contract.TickBitmap(&_UniswapV3Pool.CallOpts, arg0)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 ) view returns(uint256)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) TickBitmap(arg0 int16) (*big.Int, error) {
	return _UniswapV3Pool.Contract.TickBitmap(&_UniswapV3Pool.CallOpts, arg0)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_UniswapV3Pool *UniswapV3PoolCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_UniswapV3Pool *UniswapV3PoolSession) TickSpacing() (*big.Int, error) {
	return _UniswapV3Pool.Contract.TickSpacing(&_UniswapV3Pool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) TickSpacing() (*big.Int, error) {
	return _UniswapV3Pool.Contract.TickSpacing(&_UniswapV3Pool.CallOpts)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_UniswapV3Pool *UniswapV3PoolCaller) Ticks(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _UniswapV3Pool.contract.Call(opts, &out, "ticks", arg0)

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
// Solidity: function ticks(int24 ) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_UniswapV3Pool *UniswapV3PoolSession) Ticks(arg0 *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _UniswapV3Pool.Contract.Ticks(&_UniswapV3Pool.CallOpts, arg0)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) Ticks(arg0 *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _UniswapV3Pool.Contract.Ticks(&_UniswapV3Pool.CallOpts, arg0)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapV3Pool *UniswapV3PoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapV3Pool *UniswapV3PoolSession) Token0() (common.Address, error) {
	return _UniswapV3Pool.Contract.Token0(&_UniswapV3Pool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) Token0() (common.Address, error) {
	return _UniswapV3Pool.Contract.Token0(&_UniswapV3Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapV3Pool *UniswapV3PoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Pool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapV3Pool *UniswapV3PoolSession) Token1() (common.Address, error) {
	return _UniswapV3Pool.Contract.Token1(&_UniswapV3Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapV3Pool *UniswapV3PoolCallerSession) Token1() (common.Address, error) {
	return _UniswapV3Pool.Contract.Token1(&_UniswapV3Pool.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_UniswapV3Pool *UniswapV3PoolTransactor) Burn(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.contract.Transact(opts, "burn", tickLower, tickUpper, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_UniswapV3Pool *UniswapV3PoolSession) Burn(tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Burn(&_UniswapV3Pool.TransactOpts, tickLower, tickUpper, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_UniswapV3Pool *UniswapV3PoolTransactorSession) Burn(tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Burn(&_UniswapV3Pool.TransactOpts, tickLower, tickUpper, amount)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_UniswapV3Pool *UniswapV3PoolTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.contract.Transact(opts, "collect", recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_UniswapV3Pool *UniswapV3PoolSession) Collect(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Collect(&_UniswapV3Pool.TransactOpts, recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_UniswapV3Pool *UniswapV3PoolTransactorSession) Collect(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Collect(&_UniswapV3Pool.TransactOpts, recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_UniswapV3Pool *UniswapV3PoolTransactor) CollectProtocol(opts *bind.TransactOpts, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.contract.Transact(opts, "collectProtocol", recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_UniswapV3Pool *UniswapV3PoolSession) CollectProtocol(recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.CollectProtocol(&_UniswapV3Pool.TransactOpts, recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_UniswapV3Pool *UniswapV3PoolTransactorSession) CollectProtocol(recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.CollectProtocol(&_UniswapV3Pool.TransactOpts, recipient, amount0Requested, amount1Requested)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_UniswapV3Pool *UniswapV3PoolTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3Pool.contract.Transact(opts, "flash", recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_UniswapV3Pool *UniswapV3PoolSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Flash(&_UniswapV3Pool.TransactOpts, recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_UniswapV3Pool *UniswapV3PoolTransactorSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Flash(&_UniswapV3Pool.TransactOpts, recipient, amount0, amount1, data)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_UniswapV3Pool *UniswapV3PoolTransactor) IncreaseObservationCardinalityNext(opts *bind.TransactOpts, observationCardinalityNext uint16) (*types.Transaction, error) {
	return _UniswapV3Pool.contract.Transact(opts, "increaseObservationCardinalityNext", observationCardinalityNext)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_UniswapV3Pool *UniswapV3PoolSession) IncreaseObservationCardinalityNext(observationCardinalityNext uint16) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.IncreaseObservationCardinalityNext(&_UniswapV3Pool.TransactOpts, observationCardinalityNext)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_UniswapV3Pool *UniswapV3PoolTransactorSession) IncreaseObservationCardinalityNext(observationCardinalityNext uint16) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.IncreaseObservationCardinalityNext(&_UniswapV3Pool.TransactOpts, observationCardinalityNext)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_UniswapV3Pool *UniswapV3PoolTransactor) Initialize(opts *bind.TransactOpts, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.contract.Transact(opts, "initialize", sqrtPriceX96)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_UniswapV3Pool *UniswapV3PoolSession) Initialize(sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Initialize(&_UniswapV3Pool.TransactOpts, sqrtPriceX96)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_UniswapV3Pool *UniswapV3PoolTransactorSession) Initialize(sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Initialize(&_UniswapV3Pool.TransactOpts, sqrtPriceX96)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_UniswapV3Pool *UniswapV3PoolTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3Pool.contract.Transact(opts, "mint", recipient, tickLower, tickUpper, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_UniswapV3Pool *UniswapV3PoolSession) Mint(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Mint(&_UniswapV3Pool.TransactOpts, recipient, tickLower, tickUpper, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_UniswapV3Pool *UniswapV3PoolTransactorSession) Mint(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Mint(&_UniswapV3Pool.TransactOpts, recipient, tickLower, tickUpper, amount, data)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x8206a4d1.
//
// Solidity: function setFeeProtocol(uint8 feeProtocol0, uint8 feeProtocol1) returns()
func (_UniswapV3Pool *UniswapV3PoolTransactor) SetFeeProtocol(opts *bind.TransactOpts, feeProtocol0 uint8, feeProtocol1 uint8) (*types.Transaction, error) {
	return _UniswapV3Pool.contract.Transact(opts, "setFeeProtocol", feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x8206a4d1.
//
// Solidity: function setFeeProtocol(uint8 feeProtocol0, uint8 feeProtocol1) returns()
func (_UniswapV3Pool *UniswapV3PoolSession) SetFeeProtocol(feeProtocol0 uint8, feeProtocol1 uint8) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.SetFeeProtocol(&_UniswapV3Pool.TransactOpts, feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x8206a4d1.
//
// Solidity: function setFeeProtocol(uint8 feeProtocol0, uint8 feeProtocol1) returns()
func (_UniswapV3Pool *UniswapV3PoolTransactorSession) SetFeeProtocol(feeProtocol0 uint8, feeProtocol1 uint8) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.SetFeeProtocol(&_UniswapV3Pool.TransactOpts, feeProtocol0, feeProtocol1)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_UniswapV3Pool *UniswapV3PoolTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3Pool.contract.Transact(opts, "swap", recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_UniswapV3Pool *UniswapV3PoolSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Swap(&_UniswapV3Pool.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_UniswapV3Pool *UniswapV3PoolTransactorSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _UniswapV3Pool.Contract.Swap(&_UniswapV3Pool.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// UniswapV3PoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the UniswapV3Pool contract.
type UniswapV3PoolBurnIterator struct {
	Event *UniswapV3PoolBurn // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolBurn)
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
		it.Event = new(UniswapV3PoolBurn)
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
func (it *UniswapV3PoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolBurn represents a Burn event raised by the UniswapV3Pool contract.
type UniswapV3PoolBurn struct {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*UniswapV3PoolBurnIterator, error) {

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

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolBurnIterator{contract: _UniswapV3Pool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolBurn, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolBurn)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "Burn", log); err != nil {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseBurn(log types.Log) (*UniswapV3PoolBurn, error) {
	event := new(UniswapV3PoolBurn)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the UniswapV3Pool contract.
type UniswapV3PoolCollectIterator struct {
	Event *UniswapV3PoolCollect // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolCollect)
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
		it.Event = new(UniswapV3PoolCollect)
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
func (it *UniswapV3PoolCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolCollect represents a Collect event raised by the UniswapV3Pool contract.
type UniswapV3PoolCollect struct {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*UniswapV3PoolCollectIterator, error) {

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

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolCollectIterator{contract: _UniswapV3Pool.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolCollect, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolCollect)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "Collect", log); err != nil {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseCollect(log types.Log) (*UniswapV3PoolCollect, error) {
	event := new(UniswapV3PoolCollect)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolCollectProtocolIterator is returned from FilterCollectProtocol and is used to iterate over the raw logs and unpacked data for CollectProtocol events raised by the UniswapV3Pool contract.
type UniswapV3PoolCollectProtocolIterator struct {
	Event *UniswapV3PoolCollectProtocol // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolCollectProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolCollectProtocol)
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
		it.Event = new(UniswapV3PoolCollectProtocol)
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
func (it *UniswapV3PoolCollectProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolCollectProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolCollectProtocol represents a CollectProtocol event raised by the UniswapV3Pool contract.
type UniswapV3PoolCollectProtocol struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectProtocol is a free log retrieval operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterCollectProtocol(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*UniswapV3PoolCollectProtocolIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolCollectProtocolIterator{contract: _UniswapV3Pool.contract, event: "CollectProtocol", logs: logs, sub: sub}, nil
}

// WatchCollectProtocol is a free log subscription operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchCollectProtocol(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolCollectProtocol, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolCollectProtocol)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseCollectProtocol(log types.Log) (*UniswapV3PoolCollectProtocol, error) {
	event := new(UniswapV3PoolCollectProtocol)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolDBGMODPOSIterator is returned from FilterDBGMODPOS and is used to iterate over the raw logs and unpacked data for DBGMODPOS events raised by the UniswapV3Pool contract.
type UniswapV3PoolDBGMODPOSIterator struct {
	Event *UniswapV3PoolDBGMODPOS // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolDBGMODPOSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolDBGMODPOS)
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
		it.Event = new(UniswapV3PoolDBGMODPOS)
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
func (it *UniswapV3PoolDBGMODPOSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolDBGMODPOSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolDBGMODPOS represents a DBGMODPOS event raised by the UniswapV3Pool contract.
type UniswapV3PoolDBGMODPOS struct {
	Owner           common.Address
	TickLower       *big.Int
	TickUpper       *big.Int
	Slot0Tick       *big.Int
	LiquidityDelta  *big.Int
	LiquidityBefore *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	SqrtPriceX96    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDBGMODPOS is a free log retrieval operation binding the contract event 0xd1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, int128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterDBGMODPOS(opts *bind.FilterOpts) (*UniswapV3PoolDBGMODPOSIterator, error) {

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "DBG_MOD_POS")
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolDBGMODPOSIterator{contract: _UniswapV3Pool.contract, event: "DBG_MOD_POS", logs: logs, sub: sub}, nil
}

// WatchDBGMODPOS is a free log subscription operation binding the contract event 0xd1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, int128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchDBGMODPOS(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolDBGMODPOS) (event.Subscription, error) {

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "DBG_MOD_POS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolDBGMODPOS)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "DBG_MOD_POS", log); err != nil {
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

// ParseDBGMODPOS is a log parse operation binding the contract event 0xd1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, int128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseDBGMODPOS(log types.Log) (*UniswapV3PoolDBGMODPOS, error) {
	event := new(UniswapV3PoolDBGMODPOS)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "DBG_MOD_POS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolDBGSWAPHEADIterator is returned from FilterDBGSWAPHEAD and is used to iterate over the raw logs and unpacked data for DBGSWAPHEAD events raised by the UniswapV3Pool contract.
type UniswapV3PoolDBGSWAPHEADIterator struct {
	Event *UniswapV3PoolDBGSWAPHEAD // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolDBGSWAPHEADIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolDBGSWAPHEAD)
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
		it.Event = new(UniswapV3PoolDBGSWAPHEAD)
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
func (it *UniswapV3PoolDBGSWAPHEADIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolDBGSWAPHEADIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolDBGSWAPHEAD represents a DBGSWAPHEAD event raised by the UniswapV3Pool contract.
type UniswapV3PoolDBGSWAPHEAD struct {
	Pool              common.Address
	ZeroForOne        bool
	Amount            *big.Int
	SqrtStartPrice    *big.Int
	SqrtPriceLimit    *big.Int
	Slot0sqrtpricex96 *big.Int
	Tick              *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPHEAD is a free log retrieval operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterDBGSWAPHEAD(opts *bind.FilterOpts, pool []common.Address) (*UniswapV3PoolDBGSWAPHEADIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "DBG_SWAP_HEAD", poolRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolDBGSWAPHEADIterator{contract: _UniswapV3Pool.contract, event: "DBG_SWAP_HEAD", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPHEAD is a free log subscription operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchDBGSWAPHEAD(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolDBGSWAPHEAD, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "DBG_SWAP_HEAD", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolDBGSWAPHEAD)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_HEAD", log); err != nil {
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

// ParseDBGSWAPHEAD is a log parse operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseDBGSWAPHEAD(log types.Log) (*UniswapV3PoolDBGSWAPHEAD, error) {
	event := new(UniswapV3PoolDBGSWAPHEAD)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_HEAD", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolDBGSWAPLOOPIterator is returned from FilterDBGSWAPLOOP and is used to iterate over the raw logs and unpacked data for DBGSWAPLOOP events raised by the UniswapV3Pool contract.
type UniswapV3PoolDBGSWAPLOOPIterator struct {
	Event *UniswapV3PoolDBGSWAPLOOP // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolDBGSWAPLOOPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolDBGSWAPLOOP)
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
		it.Event = new(UniswapV3PoolDBGSWAPLOOP)
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
func (it *UniswapV3PoolDBGSWAPLOOPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolDBGSWAPLOOPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolDBGSWAPLOOP represents a DBGSWAPLOOP event raised by the UniswapV3Pool contract.
type UniswapV3PoolDBGSWAPLOOP struct {
	Pool                common.Address
	SqrtPriceX96        *big.Int
	SqrtPriceStartX96   *big.Int
	SqrtPriceNextX96    *big.Int
	Tick                *big.Int
	TickCumulative      *big.Int
	Initialized         bool
	StepAmountIn        *big.Int
	StepAmountOut       *big.Int
	AmountProcessed     *big.Int
	FeeAmount           *big.Int
	Liquidity           *big.Int
	ExactInput          bool
	FeeGrowthGlobalX128 *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPLOOP is a free log retrieval operation binding the contract event 0xf5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d5929.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int56 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, int256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterDBGSWAPLOOP(opts *bind.FilterOpts, pool []common.Address) (*UniswapV3PoolDBGSWAPLOOPIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "DBG_SWAP_LOOP", poolRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolDBGSWAPLOOPIterator{contract: _UniswapV3Pool.contract, event: "DBG_SWAP_LOOP", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPLOOP is a free log subscription operation binding the contract event 0xf5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d5929.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int56 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, int256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchDBGSWAPLOOP(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolDBGSWAPLOOP, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "DBG_SWAP_LOOP", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolDBGSWAPLOOP)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_LOOP", log); err != nil {
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

// ParseDBGSWAPLOOP is a log parse operation binding the contract event 0xf5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d5929.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int56 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, int256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseDBGSWAPLOOP(log types.Log) (*UniswapV3PoolDBGSWAPLOOP, error) {
	event := new(UniswapV3PoolDBGSWAPLOOP)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_LOOP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolDBGSWAPTAILIterator is returned from FilterDBGSWAPTAIL and is used to iterate over the raw logs and unpacked data for DBGSWAPTAIL events raised by the UniswapV3Pool contract.
type UniswapV3PoolDBGSWAPTAILIterator struct {
	Event *UniswapV3PoolDBGSWAPTAIL // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolDBGSWAPTAILIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolDBGSWAPTAIL)
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
		it.Event = new(UniswapV3PoolDBGSWAPTAIL)
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
func (it *UniswapV3PoolDBGSWAPTAILIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolDBGSWAPTAILIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolDBGSWAPTAIL represents a DBGSWAPTAIL event raised by the UniswapV3Pool contract.
type UniswapV3PoolDBGSWAPTAIL struct {
	Pool                 common.Address
	FeeGrowthGlobal0X128 *big.Int
	FeeGrowthGlobal1X128 *big.Int
	Tick                 *big.Int
	SqrtPriceX96         *big.Int
	Liquidity            *big.Int
	LiquidityDiff        *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPTAIL is a free log retrieval operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterDBGSWAPTAIL(opts *bind.FilterOpts, pool []common.Address) (*UniswapV3PoolDBGSWAPTAILIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "DBG_SWAP_TAIL", poolRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolDBGSWAPTAILIterator{contract: _UniswapV3Pool.contract, event: "DBG_SWAP_TAIL", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPTAIL is a free log subscription operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchDBGSWAPTAIL(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolDBGSWAPTAIL, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "DBG_SWAP_TAIL", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolDBGSWAPTAIL)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_TAIL", log); err != nil {
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

// ParseDBGSWAPTAIL is a log parse operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseDBGSWAPTAIL(log types.Log) (*UniswapV3PoolDBGSWAPTAIL, error) {
	event := new(UniswapV3PoolDBGSWAPTAIL)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "DBG_SWAP_TAIL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolDBGUPDPOSIterator is returned from FilterDBGUPDPOS and is used to iterate over the raw logs and unpacked data for DBGUPDPOS events raised by the UniswapV3Pool contract.
type UniswapV3PoolDBGUPDPOSIterator struct {
	Event *UniswapV3PoolDBGUPDPOS // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolDBGUPDPOSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolDBGUPDPOS)
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
		it.Event = new(UniswapV3PoolDBGUPDPOS)
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
func (it *UniswapV3PoolDBGUPDPOSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolDBGUPDPOSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolDBGUPDPOS represents a DBGUPDPOS event raised by the UniswapV3Pool contract.
type UniswapV3PoolDBGUPDPOS struct {
	Owner                      common.Address
	TickLower                  *big.Int
	TickUpper                  *big.Int
	Tick                       *big.Int
	LiquidityDelta             *big.Int
	FeeGrowthGlobal0X128Before *big.Int
	FeeGrowthGlobal1X128Before *big.Int
	FeeGrowthInside0X128       *big.Int
	FeeGrowthInside1X128       *big.Int
	FlippedLower               bool
	FlippedUpper               bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterDBGUPDPOS is a free log retrieval operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterDBGUPDPOS(opts *bind.FilterOpts) (*UniswapV3PoolDBGUPDPOSIterator, error) {

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "DBG_UPD_POS")
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolDBGUPDPOSIterator{contract: _UniswapV3Pool.contract, event: "DBG_UPD_POS", logs: logs, sub: sub}, nil
}

// WatchDBGUPDPOS is a free log subscription operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchDBGUPDPOS(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolDBGUPDPOS) (event.Subscription, error) {

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "DBG_UPD_POS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolDBGUPDPOS)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "DBG_UPD_POS", log); err != nil {
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

// ParseDBGUPDPOS is a log parse operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseDBGUPDPOS(log types.Log) (*UniswapV3PoolDBGUPDPOS, error) {
	event := new(UniswapV3PoolDBGUPDPOS)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "DBG_UPD_POS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the UniswapV3Pool contract.
type UniswapV3PoolFlashIterator struct {
	Event *UniswapV3PoolFlash // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolFlash)
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
		it.Event = new(UniswapV3PoolFlash)
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
func (it *UniswapV3PoolFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolFlash represents a Flash event raised by the UniswapV3Pool contract.
type UniswapV3PoolFlash struct {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*UniswapV3PoolFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolFlashIterator{contract: _UniswapV3Pool.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolFlash)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "Flash", log); err != nil {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseFlash(log types.Log) (*UniswapV3PoolFlash, error) {
	event := new(UniswapV3PoolFlash)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolIncreaseObservationCardinalityNextIterator is returned from FilterIncreaseObservationCardinalityNext and is used to iterate over the raw logs and unpacked data for IncreaseObservationCardinalityNext events raised by the UniswapV3Pool contract.
type UniswapV3PoolIncreaseObservationCardinalityNextIterator struct {
	Event *UniswapV3PoolIncreaseObservationCardinalityNext // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolIncreaseObservationCardinalityNextIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolIncreaseObservationCardinalityNext)
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
		it.Event = new(UniswapV3PoolIncreaseObservationCardinalityNext)
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
func (it *UniswapV3PoolIncreaseObservationCardinalityNextIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolIncreaseObservationCardinalityNextIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolIncreaseObservationCardinalityNext represents a IncreaseObservationCardinalityNext event raised by the UniswapV3Pool contract.
type UniswapV3PoolIncreaseObservationCardinalityNext struct {
	ObservationCardinalityNextOld uint16
	ObservationCardinalityNextNew uint16
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterIncreaseObservationCardinalityNext is a free log retrieval operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterIncreaseObservationCardinalityNext(opts *bind.FilterOpts) (*UniswapV3PoolIncreaseObservationCardinalityNextIterator, error) {

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolIncreaseObservationCardinalityNextIterator{contract: _UniswapV3Pool.contract, event: "IncreaseObservationCardinalityNext", logs: logs, sub: sub}, nil
}

// WatchIncreaseObservationCardinalityNext is a free log subscription operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchIncreaseObservationCardinalityNext(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolIncreaseObservationCardinalityNext) (event.Subscription, error) {

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolIncreaseObservationCardinalityNext)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseIncreaseObservationCardinalityNext(log types.Log) (*UniswapV3PoolIncreaseObservationCardinalityNext, error) {
	event := new(UniswapV3PoolIncreaseObservationCardinalityNext)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the UniswapV3Pool contract.
type UniswapV3PoolInitializeIterator struct {
	Event *UniswapV3PoolInitialize // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolInitialize)
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
		it.Event = new(UniswapV3PoolInitialize)
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
func (it *UniswapV3PoolInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolInitialize represents a Initialize event raised by the UniswapV3Pool contract.
type UniswapV3PoolInitialize struct {
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterInitialize(opts *bind.FilterOpts) (*UniswapV3PoolInitializeIterator, error) {

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolInitializeIterator{contract: _UniswapV3Pool.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolInitialize) (event.Subscription, error) {

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolInitialize)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseInitialize(log types.Log) (*UniswapV3PoolInitialize, error) {
	event := new(UniswapV3PoolInitialize)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the UniswapV3Pool contract.
type UniswapV3PoolMintIterator struct {
	Event *UniswapV3PoolMint // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolMint)
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
		it.Event = new(UniswapV3PoolMint)
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
func (it *UniswapV3PoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolMint represents a Mint event raised by the UniswapV3Pool contract.
type UniswapV3PoolMint struct {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*UniswapV3PoolMintIterator, error) {

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

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolMintIterator{contract: _UniswapV3Pool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolMint, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolMint)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseMint(log types.Log) (*UniswapV3PoolMint, error) {
	event := new(UniswapV3PoolMint)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolSetFeeProtocolIterator is returned from FilterSetFeeProtocol and is used to iterate over the raw logs and unpacked data for SetFeeProtocol events raised by the UniswapV3Pool contract.
type UniswapV3PoolSetFeeProtocolIterator struct {
	Event *UniswapV3PoolSetFeeProtocol // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolSetFeeProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolSetFeeProtocol)
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
		it.Event = new(UniswapV3PoolSetFeeProtocol)
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
func (it *UniswapV3PoolSetFeeProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolSetFeeProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolSetFeeProtocol represents a SetFeeProtocol event raised by the UniswapV3Pool contract.
type UniswapV3PoolSetFeeProtocol struct {
	FeeProtocol0Old uint8
	FeeProtocol1Old uint8
	FeeProtocol0New uint8
	FeeProtocol1New uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeProtocol is a free log retrieval operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterSetFeeProtocol(opts *bind.FilterOpts) (*UniswapV3PoolSetFeeProtocolIterator, error) {

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolSetFeeProtocolIterator{contract: _UniswapV3Pool.contract, event: "SetFeeProtocol", logs: logs, sub: sub}, nil
}

// WatchSetFeeProtocol is a free log subscription operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchSetFeeProtocol(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolSetFeeProtocol) (event.Subscription, error) {

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolSetFeeProtocol)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseSetFeeProtocol(log types.Log) (*UniswapV3PoolSetFeeProtocol, error) {
	event := new(UniswapV3PoolSetFeeProtocol)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the UniswapV3Pool contract.
type UniswapV3PoolSwapIterator struct {
	Event *UniswapV3PoolSwap // Event containing the contract specifics and raw log

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
func (it *UniswapV3PoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3PoolSwap)
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
		it.Event = new(UniswapV3PoolSwap)
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
func (it *UniswapV3PoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3PoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3PoolSwap represents a Swap event raised by the UniswapV3Pool contract.
type UniswapV3PoolSwap struct {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*UniswapV3PoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolSwapIterator{contract: _UniswapV3Pool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_UniswapV3Pool *UniswapV3PoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *UniswapV3PoolSwap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _UniswapV3Pool.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3PoolSwap)
				if err := _UniswapV3Pool.contract.UnpackLog(event, "Swap", log); err != nil {
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
func (_UniswapV3Pool *UniswapV3PoolFilterer) ParseSwap(log types.Log) (*UniswapV3PoolSwap, error) {
	event := new(UniswapV3PoolSwap)
	if err := _UniswapV3Pool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV3PoolDeployerMetaData contains all meta data concerning the UniswapV3PoolDeployer contract.
var UniswapV3PoolDeployerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060e18061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80638903573014602d575b600080fd5b60336075565b604080516001600160a01b0396871681529486166020860152929094168383015262ffffff16606083015260029290920b608082015290519081900360a00190f35b600054600154600280546001600160a01b03938416939283169281169162ffffff600160a01b83041691600160b81b9004900b8556fea2646970667358221220acfe029b48c0f73fe97ec1aa295619f8225f63380ce12d9f74eda7b28103245864736f6c63430007060033",
}

// UniswapV3PoolDeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3PoolDeployerMetaData.ABI instead.
var UniswapV3PoolDeployerABI = UniswapV3PoolDeployerMetaData.ABI

// UniswapV3PoolDeployerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniswapV3PoolDeployerMetaData.Bin instead.
var UniswapV3PoolDeployerBin = UniswapV3PoolDeployerMetaData.Bin

// DeployUniswapV3PoolDeployer deploys a new Ethereum contract, binding an instance of UniswapV3PoolDeployer to it.
func DeployUniswapV3PoolDeployer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniswapV3PoolDeployer, error) {
	parsed, err := UniswapV3PoolDeployerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniswapV3PoolDeployerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniswapV3PoolDeployer{UniswapV3PoolDeployerCaller: UniswapV3PoolDeployerCaller{contract: contract}, UniswapV3PoolDeployerTransactor: UniswapV3PoolDeployerTransactor{contract: contract}, UniswapV3PoolDeployerFilterer: UniswapV3PoolDeployerFilterer{contract: contract}}, nil
}

// UniswapV3PoolDeployer is an auto generated Go binding around an Ethereum contract.
type UniswapV3PoolDeployer struct {
	UniswapV3PoolDeployerCaller     // Read-only binding to the contract
	UniswapV3PoolDeployerTransactor // Write-only binding to the contract
	UniswapV3PoolDeployerFilterer   // Log filterer for contract events
}

// UniswapV3PoolDeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3PoolDeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PoolDeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3PoolDeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PoolDeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3PoolDeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3PoolDeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3PoolDeployerSession struct {
	Contract     *UniswapV3PoolDeployer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UniswapV3PoolDeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3PoolDeployerCallerSession struct {
	Contract *UniswapV3PoolDeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// UniswapV3PoolDeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3PoolDeployerTransactorSession struct {
	Contract     *UniswapV3PoolDeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// UniswapV3PoolDeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3PoolDeployerRaw struct {
	Contract *UniswapV3PoolDeployer // Generic contract binding to access the raw methods on
}

// UniswapV3PoolDeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3PoolDeployerCallerRaw struct {
	Contract *UniswapV3PoolDeployerCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3PoolDeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3PoolDeployerTransactorRaw struct {
	Contract *UniswapV3PoolDeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3PoolDeployer creates a new instance of UniswapV3PoolDeployer, bound to a specific deployed contract.
func NewUniswapV3PoolDeployer(address common.Address, backend bind.ContractBackend) (*UniswapV3PoolDeployer, error) {
	contract, err := bindUniswapV3PoolDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolDeployer{UniswapV3PoolDeployerCaller: UniswapV3PoolDeployerCaller{contract: contract}, UniswapV3PoolDeployerTransactor: UniswapV3PoolDeployerTransactor{contract: contract}, UniswapV3PoolDeployerFilterer: UniswapV3PoolDeployerFilterer{contract: contract}}, nil
}

// NewUniswapV3PoolDeployerCaller creates a new read-only instance of UniswapV3PoolDeployer, bound to a specific deployed contract.
func NewUniswapV3PoolDeployerCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3PoolDeployerCaller, error) {
	contract, err := bindUniswapV3PoolDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolDeployerCaller{contract: contract}, nil
}

// NewUniswapV3PoolDeployerTransactor creates a new write-only instance of UniswapV3PoolDeployer, bound to a specific deployed contract.
func NewUniswapV3PoolDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3PoolDeployerTransactor, error) {
	contract, err := bindUniswapV3PoolDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolDeployerTransactor{contract: contract}, nil
}

// NewUniswapV3PoolDeployerFilterer creates a new log filterer instance of UniswapV3PoolDeployer, bound to a specific deployed contract.
func NewUniswapV3PoolDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3PoolDeployerFilterer, error) {
	contract, err := bindUniswapV3PoolDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolDeployerFilterer{contract: contract}, nil
}

// bindUniswapV3PoolDeployer binds a generic wrapper to an already deployed contract.
func bindUniswapV3PoolDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV3PoolDeployerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3PoolDeployer *UniswapV3PoolDeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3PoolDeployer.Contract.UniswapV3PoolDeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3PoolDeployer *UniswapV3PoolDeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PoolDeployer.Contract.UniswapV3PoolDeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3PoolDeployer *UniswapV3PoolDeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3PoolDeployer.Contract.UniswapV3PoolDeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3PoolDeployer *UniswapV3PoolDeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3PoolDeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3PoolDeployer *UniswapV3PoolDeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3PoolDeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3PoolDeployer *UniswapV3PoolDeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3PoolDeployer.Contract.contract.Transact(opts, method, params...)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniswapV3PoolDeployer *UniswapV3PoolDeployerCaller) Parameters(opts *bind.CallOpts) (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	var out []interface{}
	err := _UniswapV3PoolDeployer.contract.Call(opts, &out, "parameters")

	outstruct := new(struct {
		Factory     common.Address
		Token0      common.Address
		Token1      common.Address
		Fee         *big.Int
		TickSpacing *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Factory = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickSpacing = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniswapV3PoolDeployer *UniswapV3PoolDeployerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _UniswapV3PoolDeployer.Contract.Parameters(&_UniswapV3PoolDeployer.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniswapV3PoolDeployer *UniswapV3PoolDeployerCallerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _UniswapV3PoolDeployer.Contract.Parameters(&_UniswapV3PoolDeployer.CallOpts)
}

// UnsafeMathMetaData contains all meta data concerning the UnsafeMath contract.
var UnsafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204fd29ec9fcceb2df20ab5cca79becc3508e527ae00b6d1504d6c5f08a4550c6b64736f6c63430007060033",
}

// UnsafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use UnsafeMathMetaData.ABI instead.
var UnsafeMathABI = UnsafeMathMetaData.ABI

// UnsafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UnsafeMathMetaData.Bin instead.
var UnsafeMathBin = UnsafeMathMetaData.Bin

// DeployUnsafeMath deploys a new Ethereum contract, binding an instance of UnsafeMath to it.
func DeployUnsafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UnsafeMath, error) {
	parsed, err := UnsafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UnsafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnsafeMath{UnsafeMathCaller: UnsafeMathCaller{contract: contract}, UnsafeMathTransactor: UnsafeMathTransactor{contract: contract}, UnsafeMathFilterer: UnsafeMathFilterer{contract: contract}}, nil
}

// UnsafeMath is an auto generated Go binding around an Ethereum contract.
type UnsafeMath struct {
	UnsafeMathCaller     // Read-only binding to the contract
	UnsafeMathTransactor // Write-only binding to the contract
	UnsafeMathFilterer   // Log filterer for contract events
}

// UnsafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnsafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnsafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnsafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnsafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnsafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnsafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnsafeMathSession struct {
	Contract     *UnsafeMath       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnsafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnsafeMathCallerSession struct {
	Contract *UnsafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UnsafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnsafeMathTransactorSession struct {
	Contract     *UnsafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UnsafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnsafeMathRaw struct {
	Contract *UnsafeMath // Generic contract binding to access the raw methods on
}

// UnsafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnsafeMathCallerRaw struct {
	Contract *UnsafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// UnsafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnsafeMathTransactorRaw struct {
	Contract *UnsafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnsafeMath creates a new instance of UnsafeMath, bound to a specific deployed contract.
func NewUnsafeMath(address common.Address, backend bind.ContractBackend) (*UnsafeMath, error) {
	contract, err := bindUnsafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnsafeMath{UnsafeMathCaller: UnsafeMathCaller{contract: contract}, UnsafeMathTransactor: UnsafeMathTransactor{contract: contract}, UnsafeMathFilterer: UnsafeMathFilterer{contract: contract}}, nil
}

// NewUnsafeMathCaller creates a new read-only instance of UnsafeMath, bound to a specific deployed contract.
func NewUnsafeMathCaller(address common.Address, caller bind.ContractCaller) (*UnsafeMathCaller, error) {
	contract, err := bindUnsafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnsafeMathCaller{contract: contract}, nil
}

// NewUnsafeMathTransactor creates a new write-only instance of UnsafeMath, bound to a specific deployed contract.
func NewUnsafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*UnsafeMathTransactor, error) {
	contract, err := bindUnsafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnsafeMathTransactor{contract: contract}, nil
}

// NewUnsafeMathFilterer creates a new log filterer instance of UnsafeMath, bound to a specific deployed contract.
func NewUnsafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*UnsafeMathFilterer, error) {
	contract, err := bindUnsafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnsafeMathFilterer{contract: contract}, nil
}

// bindUnsafeMath binds a generic wrapper to an already deployed contract.
func bindUnsafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnsafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnsafeMath *UnsafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnsafeMath.Contract.UnsafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnsafeMath *UnsafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnsafeMath.Contract.UnsafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnsafeMath *UnsafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnsafeMath.Contract.UnsafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnsafeMath *UnsafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnsafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnsafeMath *UnsafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnsafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnsafeMath *UnsafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnsafeMath.Contract.contract.Transact(opts, method, params...)
}
