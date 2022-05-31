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

// IApprovalReceiverMetaData contains all meta data concerning the IApprovalReceiver contract.
var IApprovalReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onTokenApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"00ba451f": "onTokenApproval(address,uint256,bytes)",
	},
}

// IApprovalReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IApprovalReceiverMetaData.ABI instead.
var IApprovalReceiverABI = IApprovalReceiverMetaData.ABI

// Deprecated: Use IApprovalReceiverMetaData.Sigs instead.
// IApprovalReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IApprovalReceiverFuncSigs = IApprovalReceiverMetaData.Sigs

// IApprovalReceiver is an auto generated Go binding around an Ethereum contract.
type IApprovalReceiver struct {
	IApprovalReceiverCaller     // Read-only binding to the contract
	IApprovalReceiverTransactor // Write-only binding to the contract
	IApprovalReceiverFilterer   // Log filterer for contract events
}

// IApprovalReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IApprovalReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApprovalReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IApprovalReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApprovalReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IApprovalReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApprovalReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IApprovalReceiverSession struct {
	Contract     *IApprovalReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IApprovalReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IApprovalReceiverCallerSession struct {
	Contract *IApprovalReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IApprovalReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IApprovalReceiverTransactorSession struct {
	Contract     *IApprovalReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IApprovalReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IApprovalReceiverRaw struct {
	Contract *IApprovalReceiver // Generic contract binding to access the raw methods on
}

// IApprovalReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IApprovalReceiverCallerRaw struct {
	Contract *IApprovalReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IApprovalReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IApprovalReceiverTransactorRaw struct {
	Contract *IApprovalReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIApprovalReceiver creates a new instance of IApprovalReceiver, bound to a specific deployed contract.
func NewIApprovalReceiver(address common.Address, backend bind.ContractBackend) (*IApprovalReceiver, error) {
	contract, err := bindIApprovalReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IApprovalReceiver{IApprovalReceiverCaller: IApprovalReceiverCaller{contract: contract}, IApprovalReceiverTransactor: IApprovalReceiverTransactor{contract: contract}, IApprovalReceiverFilterer: IApprovalReceiverFilterer{contract: contract}}, nil
}

// NewIApprovalReceiverCaller creates a new read-only instance of IApprovalReceiver, bound to a specific deployed contract.
func NewIApprovalReceiverCaller(address common.Address, caller bind.ContractCaller) (*IApprovalReceiverCaller, error) {
	contract, err := bindIApprovalReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IApprovalReceiverCaller{contract: contract}, nil
}

// NewIApprovalReceiverTransactor creates a new write-only instance of IApprovalReceiver, bound to a specific deployed contract.
func NewIApprovalReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IApprovalReceiverTransactor, error) {
	contract, err := bindIApprovalReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IApprovalReceiverTransactor{contract: contract}, nil
}

// NewIApprovalReceiverFilterer creates a new log filterer instance of IApprovalReceiver, bound to a specific deployed contract.
func NewIApprovalReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IApprovalReceiverFilterer, error) {
	contract, err := bindIApprovalReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IApprovalReceiverFilterer{contract: contract}, nil
}

// bindIApprovalReceiver binds a generic wrapper to an already deployed contract.
func bindIApprovalReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IApprovalReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApprovalReceiver *IApprovalReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApprovalReceiver.Contract.IApprovalReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApprovalReceiver *IApprovalReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.IApprovalReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApprovalReceiver *IApprovalReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.IApprovalReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApprovalReceiver *IApprovalReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApprovalReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApprovalReceiver *IApprovalReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApprovalReceiver *IApprovalReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnTokenApproval is a paid mutator transaction binding the contract method 0x00ba451f.
//
// Solidity: function onTokenApproval(address , uint256 , bytes ) returns(bool)
func (_IApprovalReceiver *IApprovalReceiverTransactor) OnTokenApproval(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _IApprovalReceiver.contract.Transact(opts, "onTokenApproval", arg0, arg1, arg2)
}

// OnTokenApproval is a paid mutator transaction binding the contract method 0x00ba451f.
//
// Solidity: function onTokenApproval(address , uint256 , bytes ) returns(bool)
func (_IApprovalReceiver *IApprovalReceiverSession) OnTokenApproval(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.OnTokenApproval(&_IApprovalReceiver.TransactOpts, arg0, arg1, arg2)
}

// OnTokenApproval is a paid mutator transaction binding the contract method 0x00ba451f.
//
// Solidity: function onTokenApproval(address , uint256 , bytes ) returns(bool)
func (_IApprovalReceiver *IApprovalReceiverTransactorSession) OnTokenApproval(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _IApprovalReceiver.Contract.OnTokenApproval(&_IApprovalReceiver.TransactOpts, arg0, arg1, arg2)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

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

// IERC2612MetaData contains all meta data concerning the IERC2612 contract.
var IERC2612MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3644e515": "DOMAIN_SEPARATOR()",
		"7ecebe00": "nonces(address)",
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
	},
}

// IERC2612ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC2612MetaData.ABI instead.
var IERC2612ABI = IERC2612MetaData.ABI

// Deprecated: Use IERC2612MetaData.Sigs instead.
// IERC2612FuncSigs maps the 4-byte function signature to its string representation.
var IERC2612FuncSigs = IERC2612MetaData.Sigs

// IERC2612 is an auto generated Go binding around an Ethereum contract.
type IERC2612 struct {
	IERC2612Caller     // Read-only binding to the contract
	IERC2612Transactor // Write-only binding to the contract
	IERC2612Filterer   // Log filterer for contract events
}

// IERC2612Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC2612Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC2612Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC2612Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC2612Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC2612Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC2612Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC2612Session struct {
	Contract     *IERC2612         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC2612CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC2612CallerSession struct {
	Contract *IERC2612Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC2612TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC2612TransactorSession struct {
	Contract     *IERC2612Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC2612Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC2612Raw struct {
	Contract *IERC2612 // Generic contract binding to access the raw methods on
}

// IERC2612CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC2612CallerRaw struct {
	Contract *IERC2612Caller // Generic read-only contract binding to access the raw methods on
}

// IERC2612TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC2612TransactorRaw struct {
	Contract *IERC2612Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC2612 creates a new instance of IERC2612, bound to a specific deployed contract.
func NewIERC2612(address common.Address, backend bind.ContractBackend) (*IERC2612, error) {
	contract, err := bindIERC2612(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC2612{IERC2612Caller: IERC2612Caller{contract: contract}, IERC2612Transactor: IERC2612Transactor{contract: contract}, IERC2612Filterer: IERC2612Filterer{contract: contract}}, nil
}

// NewIERC2612Caller creates a new read-only instance of IERC2612, bound to a specific deployed contract.
func NewIERC2612Caller(address common.Address, caller bind.ContractCaller) (*IERC2612Caller, error) {
	contract, err := bindIERC2612(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC2612Caller{contract: contract}, nil
}

// NewIERC2612Transactor creates a new write-only instance of IERC2612, bound to a specific deployed contract.
func NewIERC2612Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC2612Transactor, error) {
	contract, err := bindIERC2612(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC2612Transactor{contract: contract}, nil
}

// NewIERC2612Filterer creates a new log filterer instance of IERC2612, bound to a specific deployed contract.
func NewIERC2612Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC2612Filterer, error) {
	contract, err := bindIERC2612(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC2612Filterer{contract: contract}, nil
}

// bindIERC2612 binds a generic wrapper to an already deployed contract.
func bindIERC2612(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC2612ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC2612 *IERC2612Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC2612.Contract.IERC2612Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC2612 *IERC2612Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC2612.Contract.IERC2612Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC2612 *IERC2612Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC2612.Contract.IERC2612Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC2612 *IERC2612CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC2612.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC2612 *IERC2612TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC2612.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC2612 *IERC2612TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC2612.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC2612 *IERC2612Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC2612.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC2612 *IERC2612Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC2612.Contract.DOMAINSEPARATOR(&_IERC2612.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC2612 *IERC2612CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC2612.Contract.DOMAINSEPARATOR(&_IERC2612.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC2612 *IERC2612Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC2612.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC2612 *IERC2612Session) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC2612.Contract.Nonces(&_IERC2612.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC2612 *IERC2612CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC2612.Contract.Nonces(&_IERC2612.CallOpts, owner)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC2612 *IERC2612Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC2612.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC2612 *IERC2612Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC2612.Contract.Permit(&_IERC2612.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC2612 *IERC2612TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC2612.Contract.Permit(&_IERC2612.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// IERC3156FlashBorrowerMetaData contains all meta data concerning the IERC3156FlashBorrower contract.
var IERC3156FlashBorrowerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onFlashLoan\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"23e30c8b": "onFlashLoan(address,address,uint256,uint256,bytes)",
	},
}

// IERC3156FlashBorrowerABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC3156FlashBorrowerMetaData.ABI instead.
var IERC3156FlashBorrowerABI = IERC3156FlashBorrowerMetaData.ABI

// Deprecated: Use IERC3156FlashBorrowerMetaData.Sigs instead.
// IERC3156FlashBorrowerFuncSigs maps the 4-byte function signature to its string representation.
var IERC3156FlashBorrowerFuncSigs = IERC3156FlashBorrowerMetaData.Sigs

// IERC3156FlashBorrower is an auto generated Go binding around an Ethereum contract.
type IERC3156FlashBorrower struct {
	IERC3156FlashBorrowerCaller     // Read-only binding to the contract
	IERC3156FlashBorrowerTransactor // Write-only binding to the contract
	IERC3156FlashBorrowerFilterer   // Log filterer for contract events
}

// IERC3156FlashBorrowerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashBorrowerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashBorrowerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC3156FlashBorrowerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashBorrowerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC3156FlashBorrowerSession struct {
	Contract     *IERC3156FlashBorrower // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IERC3156FlashBorrowerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC3156FlashBorrowerCallerSession struct {
	Contract *IERC3156FlashBorrowerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IERC3156FlashBorrowerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC3156FlashBorrowerTransactorSession struct {
	Contract     *IERC3156FlashBorrowerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IERC3156FlashBorrowerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC3156FlashBorrowerRaw struct {
	Contract *IERC3156FlashBorrower // Generic contract binding to access the raw methods on
}

// IERC3156FlashBorrowerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerCallerRaw struct {
	Contract *IERC3156FlashBorrowerCaller // Generic read-only contract binding to access the raw methods on
}

// IERC3156FlashBorrowerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerTransactorRaw struct {
	Contract *IERC3156FlashBorrowerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC3156FlashBorrower creates a new instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrower(address common.Address, backend bind.ContractBackend) (*IERC3156FlashBorrower, error) {
	contract, err := bindIERC3156FlashBorrower(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrower{IERC3156FlashBorrowerCaller: IERC3156FlashBorrowerCaller{contract: contract}, IERC3156FlashBorrowerTransactor: IERC3156FlashBorrowerTransactor{contract: contract}, IERC3156FlashBorrowerFilterer: IERC3156FlashBorrowerFilterer{contract: contract}}, nil
}

// NewIERC3156FlashBorrowerCaller creates a new read-only instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrowerCaller(address common.Address, caller bind.ContractCaller) (*IERC3156FlashBorrowerCaller, error) {
	contract, err := bindIERC3156FlashBorrower(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrowerCaller{contract: contract}, nil
}

// NewIERC3156FlashBorrowerTransactor creates a new write-only instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrowerTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC3156FlashBorrowerTransactor, error) {
	contract, err := bindIERC3156FlashBorrower(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrowerTransactor{contract: contract}, nil
}

// NewIERC3156FlashBorrowerFilterer creates a new log filterer instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrowerFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC3156FlashBorrowerFilterer, error) {
	contract, err := bindIERC3156FlashBorrower(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrowerFilterer{contract: contract}, nil
}

// bindIERC3156FlashBorrower binds a generic wrapper to an already deployed contract.
func bindIERC3156FlashBorrower(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC3156FlashBorrowerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashBorrower.Contract.IERC3156FlashBorrowerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.IERC3156FlashBorrowerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.IERC3156FlashBorrowerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashBorrower.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.contract.Transact(opts, method, params...)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactor) OnFlashLoan(opts *bind.TransactOpts, initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.contract.Transact(opts, "onFlashLoan", initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerSession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.OnFlashLoan(&_IERC3156FlashBorrower.TransactOpts, initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactorSession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.OnFlashLoan(&_IERC3156FlashBorrower.TransactOpts, initiator, token, amount, fee, data)
}

// IERC3156FlashLenderMetaData contains all meta data concerning the IERC3156FlashLender contract.
var IERC3156FlashLenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d9d98ce4": "flashFee(address,uint256)",
		"5cffe9de": "flashLoan(address,address,uint256,bytes)",
		"613255ab": "maxFlashLoan(address)",
	},
}

// IERC3156FlashLenderABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC3156FlashLenderMetaData.ABI instead.
var IERC3156FlashLenderABI = IERC3156FlashLenderMetaData.ABI

// Deprecated: Use IERC3156FlashLenderMetaData.Sigs instead.
// IERC3156FlashLenderFuncSigs maps the 4-byte function signature to its string representation.
var IERC3156FlashLenderFuncSigs = IERC3156FlashLenderMetaData.Sigs

// IERC3156FlashLender is an auto generated Go binding around an Ethereum contract.
type IERC3156FlashLender struct {
	IERC3156FlashLenderCaller     // Read-only binding to the contract
	IERC3156FlashLenderTransactor // Write-only binding to the contract
	IERC3156FlashLenderFilterer   // Log filterer for contract events
}

// IERC3156FlashLenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC3156FlashLenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashLenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC3156FlashLenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashLenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC3156FlashLenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashLenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC3156FlashLenderSession struct {
	Contract     *IERC3156FlashLender // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IERC3156FlashLenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC3156FlashLenderCallerSession struct {
	Contract *IERC3156FlashLenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IERC3156FlashLenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC3156FlashLenderTransactorSession struct {
	Contract     *IERC3156FlashLenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IERC3156FlashLenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC3156FlashLenderRaw struct {
	Contract *IERC3156FlashLender // Generic contract binding to access the raw methods on
}

// IERC3156FlashLenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC3156FlashLenderCallerRaw struct {
	Contract *IERC3156FlashLenderCaller // Generic read-only contract binding to access the raw methods on
}

// IERC3156FlashLenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC3156FlashLenderTransactorRaw struct {
	Contract *IERC3156FlashLenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC3156FlashLender creates a new instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLender(address common.Address, backend bind.ContractBackend) (*IERC3156FlashLender, error) {
	contract, err := bindIERC3156FlashLender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLender{IERC3156FlashLenderCaller: IERC3156FlashLenderCaller{contract: contract}, IERC3156FlashLenderTransactor: IERC3156FlashLenderTransactor{contract: contract}, IERC3156FlashLenderFilterer: IERC3156FlashLenderFilterer{contract: contract}}, nil
}

// NewIERC3156FlashLenderCaller creates a new read-only instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLenderCaller(address common.Address, caller bind.ContractCaller) (*IERC3156FlashLenderCaller, error) {
	contract, err := bindIERC3156FlashLender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLenderCaller{contract: contract}, nil
}

// NewIERC3156FlashLenderTransactor creates a new write-only instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLenderTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC3156FlashLenderTransactor, error) {
	contract, err := bindIERC3156FlashLender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLenderTransactor{contract: contract}, nil
}

// NewIERC3156FlashLenderFilterer creates a new log filterer instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLenderFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC3156FlashLenderFilterer, error) {
	contract, err := bindIERC3156FlashLender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLenderFilterer{contract: contract}, nil
}

// bindIERC3156FlashLender binds a generic wrapper to an already deployed contract.
func bindIERC3156FlashLender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC3156FlashLenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashLender *IERC3156FlashLenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashLender.Contract.IERC3156FlashLenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashLender *IERC3156FlashLenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.IERC3156FlashLenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashLender *IERC3156FlashLenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.IERC3156FlashLenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashLender *IERC3156FlashLenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashLender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashLender *IERC3156FlashLenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashLender *IERC3156FlashLenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.contract.Transact(opts, method, params...)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCaller) FlashFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC3156FlashLender.contract.Call(opts, &out, "flashFee", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.FlashFee(&_IERC3156FlashLender.CallOpts, token, amount)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCallerSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.FlashFee(&_IERC3156FlashLender.CallOpts, token, amount)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCaller) MaxFlashLoan(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC3156FlashLender.contract.Call(opts, &out, "maxFlashLoan", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.MaxFlashLoan(&_IERC3156FlashLender.CallOpts, token)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCallerSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.MaxFlashLoan(&_IERC3156FlashLender.CallOpts, token)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IERC3156FlashLender *IERC3156FlashLenderTransactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashLender.contract.Transact(opts, "flashLoan", receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IERC3156FlashLender *IERC3156FlashLenderSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.FlashLoan(&_IERC3156FlashLender.TransactOpts, receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IERC3156FlashLender *IERC3156FlashLenderTransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.FlashLoan(&_IERC3156FlashLender.TransactOpts, receiver, token, amount, data)
}

// ITransferReceiverMetaData contains all meta data concerning the ITransferReceiver contract.
var ITransferReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a4c0ed36": "onTokenTransfer(address,uint256,bytes)",
	},
}

// ITransferReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use ITransferReceiverMetaData.ABI instead.
var ITransferReceiverABI = ITransferReceiverMetaData.ABI

// Deprecated: Use ITransferReceiverMetaData.Sigs instead.
// ITransferReceiverFuncSigs maps the 4-byte function signature to its string representation.
var ITransferReceiverFuncSigs = ITransferReceiverMetaData.Sigs

// ITransferReceiver is an auto generated Go binding around an Ethereum contract.
type ITransferReceiver struct {
	ITransferReceiverCaller     // Read-only binding to the contract
	ITransferReceiverTransactor // Write-only binding to the contract
	ITransferReceiverFilterer   // Log filterer for contract events
}

// ITransferReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITransferReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITransferReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITransferReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITransferReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITransferReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITransferReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITransferReceiverSession struct {
	Contract     *ITransferReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ITransferReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITransferReceiverCallerSession struct {
	Contract *ITransferReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ITransferReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITransferReceiverTransactorSession struct {
	Contract     *ITransferReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ITransferReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITransferReceiverRaw struct {
	Contract *ITransferReceiver // Generic contract binding to access the raw methods on
}

// ITransferReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITransferReceiverCallerRaw struct {
	Contract *ITransferReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ITransferReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITransferReceiverTransactorRaw struct {
	Contract *ITransferReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITransferReceiver creates a new instance of ITransferReceiver, bound to a specific deployed contract.
func NewITransferReceiver(address common.Address, backend bind.ContractBackend) (*ITransferReceiver, error) {
	contract, err := bindITransferReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITransferReceiver{ITransferReceiverCaller: ITransferReceiverCaller{contract: contract}, ITransferReceiverTransactor: ITransferReceiverTransactor{contract: contract}, ITransferReceiverFilterer: ITransferReceiverFilterer{contract: contract}}, nil
}

// NewITransferReceiverCaller creates a new read-only instance of ITransferReceiver, bound to a specific deployed contract.
func NewITransferReceiverCaller(address common.Address, caller bind.ContractCaller) (*ITransferReceiverCaller, error) {
	contract, err := bindITransferReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITransferReceiverCaller{contract: contract}, nil
}

// NewITransferReceiverTransactor creates a new write-only instance of ITransferReceiver, bound to a specific deployed contract.
func NewITransferReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ITransferReceiverTransactor, error) {
	contract, err := bindITransferReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITransferReceiverTransactor{contract: contract}, nil
}

// NewITransferReceiverFilterer creates a new log filterer instance of ITransferReceiver, bound to a specific deployed contract.
func NewITransferReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ITransferReceiverFilterer, error) {
	contract, err := bindITransferReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITransferReceiverFilterer{contract: contract}, nil
}

// bindITransferReceiver binds a generic wrapper to an already deployed contract.
func bindITransferReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITransferReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITransferReceiver *ITransferReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITransferReceiver.Contract.ITransferReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITransferReceiver *ITransferReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.ITransferReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITransferReceiver *ITransferReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.ITransferReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITransferReceiver *ITransferReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITransferReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITransferReceiver *ITransferReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITransferReceiver *ITransferReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes ) returns(bool)
func (_ITransferReceiver *ITransferReceiverTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _ITransferReceiver.contract.Transact(opts, "onTokenTransfer", arg0, arg1, arg2)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes ) returns(bool)
func (_ITransferReceiver *ITransferReceiverSession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.OnTokenTransfer(&_ITransferReceiver.TransactOpts, arg0, arg1, arg2)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes ) returns(bool)
func (_ITransferReceiver *ITransferReceiverTransactorSession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _ITransferReceiver.Contract.OnTokenTransfer(&_ITransferReceiver.TransactOpts, arg0, arg1, arg2)
}

// IWETH10MetaData contains all meta data concerning the IWETH10 contract.
var IWETH10MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"depositToAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3644e515": "DOMAIN_SEPARATOR()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"cae9ca51": "approveAndCall(address,uint256,bytes)",
		"70a08231": "balanceOf(address)",
		"d0e30db0": "deposit()",
		"b760faf9": "depositTo(address)",
		"5ddb7d7e": "depositToAndCall(address,bytes)",
		"d9d98ce4": "flashFee(address,uint256)",
		"5cffe9de": "flashLoan(address,address,uint256,bytes)",
		"8b28d32f": "flashMinted()",
		"613255ab": "maxFlashLoan(address)",
		"7ecebe00": "nonces(address)",
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"4000aea0": "transferAndCall(address,uint256,bytes)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"2e1a7d4d": "withdraw(uint256)",
		"9555a942": "withdrawFrom(address,address,uint256)",
		"205c2878": "withdrawTo(address,uint256)",
	},
}

// IWETH10ABI is the input ABI used to generate the binding from.
// Deprecated: Use IWETH10MetaData.ABI instead.
var IWETH10ABI = IWETH10MetaData.ABI

// Deprecated: Use IWETH10MetaData.Sigs instead.
// IWETH10FuncSigs maps the 4-byte function signature to its string representation.
var IWETH10FuncSigs = IWETH10MetaData.Sigs

// IWETH10 is an auto generated Go binding around an Ethereum contract.
type IWETH10 struct {
	IWETH10Caller     // Read-only binding to the contract
	IWETH10Transactor // Write-only binding to the contract
	IWETH10Filterer   // Log filterer for contract events
}

// IWETH10Caller is an auto generated read-only Go binding around an Ethereum contract.
type IWETH10Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH10Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETH10Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH10Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETH10Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH10Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETH10Session struct {
	Contract     *IWETH10          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH10CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETH10CallerSession struct {
	Contract *IWETH10Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IWETH10TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETH10TransactorSession struct {
	Contract     *IWETH10Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IWETH10Raw is an auto generated low-level Go binding around an Ethereum contract.
type IWETH10Raw struct {
	Contract *IWETH10 // Generic contract binding to access the raw methods on
}

// IWETH10CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETH10CallerRaw struct {
	Contract *IWETH10Caller // Generic read-only contract binding to access the raw methods on
}

// IWETH10TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETH10TransactorRaw struct {
	Contract *IWETH10Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH10 creates a new instance of IWETH10, bound to a specific deployed contract.
func NewIWETH10(address common.Address, backend bind.ContractBackend) (*IWETH10, error) {
	contract, err := bindIWETH10(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH10{IWETH10Caller: IWETH10Caller{contract: contract}, IWETH10Transactor: IWETH10Transactor{contract: contract}, IWETH10Filterer: IWETH10Filterer{contract: contract}}, nil
}

// NewIWETH10Caller creates a new read-only instance of IWETH10, bound to a specific deployed contract.
func NewIWETH10Caller(address common.Address, caller bind.ContractCaller) (*IWETH10Caller, error) {
	contract, err := bindIWETH10(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH10Caller{contract: contract}, nil
}

// NewIWETH10Transactor creates a new write-only instance of IWETH10, bound to a specific deployed contract.
func NewIWETH10Transactor(address common.Address, transactor bind.ContractTransactor) (*IWETH10Transactor, error) {
	contract, err := bindIWETH10(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH10Transactor{contract: contract}, nil
}

// NewIWETH10Filterer creates a new log filterer instance of IWETH10, bound to a specific deployed contract.
func NewIWETH10Filterer(address common.Address, filterer bind.ContractFilterer) (*IWETH10Filterer, error) {
	contract, err := bindIWETH10(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETH10Filterer{contract: contract}, nil
}

// bindIWETH10 binds a generic wrapper to an already deployed contract.
func bindIWETH10(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWETH10ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH10 *IWETH10Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH10.Contract.IWETH10Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH10 *IWETH10Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH10.Contract.IWETH10Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH10 *IWETH10Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH10.Contract.IWETH10Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH10 *IWETH10CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH10.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH10 *IWETH10TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH10.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH10 *IWETH10TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH10.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IWETH10 *IWETH10Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IWETH10 *IWETH10Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _IWETH10.Contract.DOMAINSEPARATOR(&_IWETH10.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IWETH10 *IWETH10CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IWETH10.Contract.DOMAINSEPARATOR(&_IWETH10.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH10 *IWETH10Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH10 *IWETH10Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH10.Contract.Allowance(&_IWETH10.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH10.Contract.Allowance(&_IWETH10.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH10 *IWETH10Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH10 *IWETH10Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWETH10.Contract.BalanceOf(&_IWETH10.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWETH10.Contract.BalanceOf(&_IWETH10.CallOpts, account)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IWETH10 *IWETH10Caller) FlashFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "flashFee", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IWETH10 *IWETH10Session) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IWETH10.Contract.FlashFee(&_IWETH10.CallOpts, token, amount)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IWETH10.Contract.FlashFee(&_IWETH10.CallOpts, token, amount)
}

// FlashMinted is a free data retrieval call binding the contract method 0x8b28d32f.
//
// Solidity: function flashMinted() view returns(uint256)
func (_IWETH10 *IWETH10Caller) FlashMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "flashMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashMinted is a free data retrieval call binding the contract method 0x8b28d32f.
//
// Solidity: function flashMinted() view returns(uint256)
func (_IWETH10 *IWETH10Session) FlashMinted() (*big.Int, error) {
	return _IWETH10.Contract.FlashMinted(&_IWETH10.CallOpts)
}

// FlashMinted is a free data retrieval call binding the contract method 0x8b28d32f.
//
// Solidity: function flashMinted() view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) FlashMinted() (*big.Int, error) {
	return _IWETH10.Contract.FlashMinted(&_IWETH10.CallOpts)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IWETH10 *IWETH10Caller) MaxFlashLoan(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "maxFlashLoan", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IWETH10 *IWETH10Session) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IWETH10.Contract.MaxFlashLoan(&_IWETH10.CallOpts, token)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IWETH10.Contract.MaxFlashLoan(&_IWETH10.CallOpts, token)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IWETH10 *IWETH10Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IWETH10 *IWETH10Session) Nonces(owner common.Address) (*big.Int, error) {
	return _IWETH10.Contract.Nonces(&_IWETH10.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IWETH10.Contract.Nonces(&_IWETH10.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH10 *IWETH10Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWETH10.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH10 *IWETH10Session) TotalSupply() (*big.Int, error) {
	return _IWETH10.Contract.TotalSupply(&_IWETH10.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH10 *IWETH10CallerSession) TotalSupply() (*big.Int, error) {
	return _IWETH10.Contract.TotalSupply(&_IWETH10.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Approve(&_IWETH10.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Approve(&_IWETH10.TransactOpts, spender, amount)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10Transactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "approveAndCall", spender, value, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10Session) ApproveAndCall(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.ApproveAndCall(&_IWETH10.TransactOpts, spender, value, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) ApproveAndCall(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.ApproveAndCall(&_IWETH10.TransactOpts, spender, value, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH10 *IWETH10Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH10 *IWETH10Session) Deposit() (*types.Transaction, error) {
	return _IWETH10.Contract.Deposit(&_IWETH10.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH10 *IWETH10TransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH10.Contract.Deposit(&_IWETH10.TransactOpts)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_IWETH10 *IWETH10Transactor) DepositTo(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "depositTo", to)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_IWETH10 *IWETH10Session) DepositTo(to common.Address) (*types.Transaction, error) {
	return _IWETH10.Contract.DepositTo(&_IWETH10.TransactOpts, to)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_IWETH10 *IWETH10TransactorSession) DepositTo(to common.Address) (*types.Transaction, error) {
	return _IWETH10.Contract.DepositTo(&_IWETH10.TransactOpts, to)
}

// DepositToAndCall is a paid mutator transaction binding the contract method 0x5ddb7d7e.
//
// Solidity: function depositToAndCall(address to, bytes data) payable returns(bool)
func (_IWETH10 *IWETH10Transactor) DepositToAndCall(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "depositToAndCall", to, data)
}

// DepositToAndCall is a paid mutator transaction binding the contract method 0x5ddb7d7e.
//
// Solidity: function depositToAndCall(address to, bytes data) payable returns(bool)
func (_IWETH10 *IWETH10Session) DepositToAndCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.DepositToAndCall(&_IWETH10.TransactOpts, to, data)
}

// DepositToAndCall is a paid mutator transaction binding the contract method 0x5ddb7d7e.
//
// Solidity: function depositToAndCall(address to, bytes data) payable returns(bool)
func (_IWETH10 *IWETH10TransactorSession) DepositToAndCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.DepositToAndCall(&_IWETH10.TransactOpts, to, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IWETH10 *IWETH10Transactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "flashLoan", receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IWETH10 *IWETH10Session) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.FlashLoan(&_IWETH10.TransactOpts, receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.FlashLoan(&_IWETH10.TransactOpts, receiver, token, amount, data)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IWETH10 *IWETH10Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IWETH10 *IWETH10Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IWETH10.Contract.Permit(&_IWETH10.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IWETH10 *IWETH10TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IWETH10.Contract.Permit(&_IWETH10.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Transfer(&_IWETH10.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Transfer(&_IWETH10.TransactOpts, recipient, amount)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10Transactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "transferAndCall", to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10Session) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.TransferAndCall(&_IWETH10.TransactOpts, to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IWETH10.Contract.TransferAndCall(&_IWETH10.TransactOpts, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.TransferFrom(&_IWETH10.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH10 *IWETH10TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.TransferFrom(&_IWETH10.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_IWETH10 *IWETH10Transactor) Withdraw(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "withdraw", value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_IWETH10 *IWETH10Session) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Withdraw(&_IWETH10.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_IWETH10 *IWETH10TransactorSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.Withdraw(&_IWETH10.TransactOpts, value)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9555a942.
//
// Solidity: function withdrawFrom(address from, address to, uint256 value) returns()
func (_IWETH10 *IWETH10Transactor) WithdrawFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "withdrawFrom", from, to, value)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9555a942.
//
// Solidity: function withdrawFrom(address from, address to, uint256 value) returns()
func (_IWETH10 *IWETH10Session) WithdrawFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.WithdrawFrom(&_IWETH10.TransactOpts, from, to, value)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9555a942.
//
// Solidity: function withdrawFrom(address from, address to, uint256 value) returns()
func (_IWETH10 *IWETH10TransactorSession) WithdrawFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.WithdrawFrom(&_IWETH10.TransactOpts, from, to, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_IWETH10 *IWETH10Transactor) WithdrawTo(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.contract.Transact(opts, "withdrawTo", to, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_IWETH10 *IWETH10Session) WithdrawTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.WithdrawTo(&_IWETH10.TransactOpts, to, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_IWETH10 *IWETH10TransactorSession) WithdrawTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH10.Contract.WithdrawTo(&_IWETH10.TransactOpts, to, value)
}

// IWETH10ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IWETH10 contract.
type IWETH10ApprovalIterator struct {
	Event *IWETH10Approval // Event containing the contract specifics and raw log

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
func (it *IWETH10ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETH10Approval)
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
		it.Event = new(IWETH10Approval)
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
func (it *IWETH10ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETH10ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETH10Approval represents a Approval event raised by the IWETH10 contract.
type IWETH10Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH10 *IWETH10Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IWETH10ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH10.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IWETH10ApprovalIterator{contract: _IWETH10.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH10 *IWETH10Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IWETH10Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH10.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETH10Approval)
				if err := _IWETH10.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IWETH10 *IWETH10Filterer) ParseApproval(log types.Log) (*IWETH10Approval, error) {
	event := new(IWETH10Approval)
	if err := _IWETH10.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWETH10TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IWETH10 contract.
type IWETH10TransferIterator struct {
	Event *IWETH10Transfer // Event containing the contract specifics and raw log

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
func (it *IWETH10TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETH10Transfer)
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
		it.Event = new(IWETH10Transfer)
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
func (it *IWETH10TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETH10TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETH10Transfer represents a Transfer event raised by the IWETH10 contract.
type IWETH10Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH10 *IWETH10Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IWETH10TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH10.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IWETH10TransferIterator{contract: _IWETH10.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH10 *IWETH10Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IWETH10Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH10.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETH10Transfer)
				if err := _IWETH10.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IWETH10 *IWETH10Filterer) ParseTransfer(log types.Log) (*IWETH10Transfer, error) {
	event := new(IWETH10Transfer)
	if err := _IWETH10.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WETH10MetaData contains all meta data concerning the WETH10 contract.
var WETH10MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CALLBACK_SUCCESS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deploymentChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"depositToAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"8237e538": "CALLBACK_SUCCESS()",
		"3644e515": "DOMAIN_SEPARATOR()",
		"30adf81f": "PERMIT_TYPEHASH()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"cae9ca51": "approveAndCall(address,uint256,bytes)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"cd0d0096": "deploymentChainId()",
		"d0e30db0": "deposit()",
		"b760faf9": "depositTo(address)",
		"5ddb7d7e": "depositToAndCall(address,bytes)",
		"d9d98ce4": "flashFee(address,uint256)",
		"5cffe9de": "flashLoan(address,address,uint256,bytes)",
		"8b28d32f": "flashMinted()",
		"613255ab": "maxFlashLoan(address)",
		"06fdde03": "name()",
		"7ecebe00": "nonces(address)",
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"4000aea0": "transferAndCall(address,uint256,bytes)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"2e1a7d4d": "withdraw(uint256)",
		"9555a942": "withdrawFrom(address,address,uint256)",
		"205c2878": "withdrawTo(address,uint256)",
	},
	Bin: "0x6101006040527f439148f0bbc682ca079e46d6e2c2f0c1e3b820f1a291b069d8882abf8cf18dd96080527f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c960a05234801561005957600080fd5b504660c081905261006981610072565b60e0525061014c565b60408051808201825260118152700577261707065642045746865722076313607c1b6020918201528151808301835260018152603160f81b9082015281517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818301527f5afefa023603a382ac22bd36644e004cefabbfaf61d4180000d8dc2b10c168b8818401527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6606082015260808101939093523060a0808501919091528251808503909101815260c0909301909152815191012090565b60805160a05160c05160e05161214161019a60003980610e7a5280611d53525080610e455280611be85280611d1e525080610e195280611cb8525080611275528061168652506121416000f3fe6080604052600436106101a05760003560e01c806370a08231116100ec578063b760faf91161008a578063d0e30db011610064578063d0e30db0146107a8578063d505accf146107b0578063d9d98ce41461080e578063dd62ed3e14610847576101e0565b8063b760faf9146106dd578063cae9ca5114610703578063cd0d009614610793576101e0565b80638b28d32f116100c65780638b28d32f146106375780639555a9421461064c57806395d89b411461068f578063a9059cbb146106a4576101e0565b806370a08231146105bc5780637ecebe00146105ef5780638237e53814610622576101e0565b806330adf81f116101595780634000aea0116101335780634000aea0146103e05780635cffe9de146104705780635ddb7d7e1461050b578063613255ab14610589576101e0565b806330adf81f1461038b578063313ce567146103a05780633644e515146103cb576101e0565b806306fdde03146101e5578063095ea7b31461026f57806318160ddd146102bc578063205c2878146102e357806323b872dd1461031e5780632e1a7d4d14610361576101e0565b366101e05733600081815260208181526040808320805434908101909155815190815290516000805160206120ab833981519152929181900390910190a3005b600080fd5b3480156101f157600080fd5b506101fa610882565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561023457818101518382015260200161021c565b50505050905090810190601f1680156102615780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561027b57600080fd5b506102a86004803603604081101561029257600080fd5b506001600160a01b0381351690602001356108af565b604080519115158252519081900360200190f35b3480156102c857600080fd5b506102d1610903565b60408051918252519081900360200190f35b3480156102ef57600080fd5b5061031c6004803603604081101561030657600080fd5b506001600160a01b03813516906020013561090b565b005b34801561032a57600080fd5b506102a86004803603606081101561034157600080fd5b506001600160a01b03813581169160208101359091169060400135610a2c565b34801561036d57600080fd5b5061031c6004803603602081101561038457600080fd5b5035610d00565b34801561039757600080fd5b506102d1610e17565b3480156103ac57600080fd5b506103b5610e3b565b6040805160ff9092168252519081900360200190f35b3480156103d757600080fd5b506102d1610e40565b3480156103ec57600080fd5b506102a86004803603606081101561040357600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561043257600080fd5b82018360208201111561044457600080fd5b803590602001918460018302840111600160201b8311171561046557600080fd5b509092509050610ea0565b34801561047c57600080fd5b506102a86004803603608081101561049357600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b8111156104cd57600080fd5b8201836020820111156104df57600080fd5b803590602001918460018302840111600160201b8311171561050057600080fd5b50909250905061112a565b6102a86004803603604081101561052157600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561054b57600080fd5b82018360208201111561055d57600080fd5b803590602001918460018302840111600160201b8311171561057e57600080fd5b509092509050611538565b34801561059557600080fd5b506102d1600480360360208110156105ac57600080fd5b50356001600160a01b0316611634565b3480156105c857600080fd5b506102d1600480360360208110156105df57600080fd5b50356001600160a01b0316611660565b3480156105fb57600080fd5b506102d16004803603602081101561061257600080fd5b50356001600160a01b0316611672565b34801561062e57600080fd5b506102d1611684565b34801561064357600080fd5b506102d16116a8565b34801561065857600080fd5b5061031c6004803603606081101561066f57600080fd5b506001600160a01b038135811691602081013590911690604001356116ae565b34801561069b57600080fd5b506101fa6118d7565b3480156106b057600080fd5b506102a8600480360360408110156106c757600080fd5b506001600160a01b0381351690602001356118f9565b61031c600480360360208110156106f357600080fd5b50356001600160a01b0316611acc565b34801561070f57600080fd5b506102a86004803603606081101561072657600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561075557600080fd5b82018360208201111561076757600080fd5b803590602001918460018302840111600160201b8311171561078857600080fd5b509092509050611b11565b34801561079f57600080fd5b506102d1611be6565b61031c611c0a565b3480156107bc57600080fd5b5061031c600480360360e08110156107d357600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060ff6080820135169060a08101359060c00135611c45565b34801561081a57600080fd5b506102d16004803603604081101561083157600080fd5b506001600160a01b038135169060200135611ee2565b34801561085357600080fd5b506102d16004803603604081101561086a57600080fd5b506001600160a01b0381358116916020013516611f4a565b604051806040016040528060118152602001700577261707065642045746865722076313607c1b81525081565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390926000805160206120cb833981519152928290030190a350600192915050565b600354470190565b336000908152602081905260409020548181101561095a5760405162461bcd60e51b81526004018080602001828103825260218152602001806120eb6021913960400191505060405180910390fd5b336000818152602081815260408083208686039055805186815290519293926000805160206120ab833981519152929181900390910190a36040516000906001600160a01b0385169084908381818185875af1925050503d80600081146109dd576040519150601f19603f3d011682016040523d82523d6000602084013e6109e2565b606091505b5050905080610a26576040805162461bcd60e51b81526020600482015260196024820152600080516020612067833981519152604482015290519081900360640190fd5b50505050565b60006001600160a01b0384163314610b11576001600160a01b03841660009081526002602090815260408083203384529091529020546000198114610b0f5782811015610ac0576040805162461bcd60e51b815260206004820152601f60248201527f574554483a2072657175657374206578636565647320616c6c6f77616e636500604482015290519081900360640190fd5b6001600160a01b0385166000818152600260209081526040808320338085529083529281902087860390819055815181815291519094926000805160206120cb833981519152928290030190a3505b505b6001600160a01b03831615610bcf576001600160a01b03841660009081526020819052604090205482811015610b785760405162461bcd60e51b81526004018080602001828103825260258152602001806120426025913960400191505060405180910390fd5b6001600160a01b038086166000818152602081815260408083208887039055938816808352918490208054880190558351878152935191936000805160206120ab833981519152929081900390910190a350610cf6565b6001600160a01b03841660009081526020819052604090205482811015610c275760405162461bcd60e51b81526004018080602001828103825260218152602001806120eb6021913960400191505060405180910390fd5b6001600160a01b0385166000818152602081815260408083208786039055805187815290519293926000805160206120ab833981519152929181900390910190a3604051600090339085908381818185875af1925050503d8060008114610caa576040519150601f19603f3d011682016040523d82523d6000602084013e610caf565b606091505b5050905080610cf3576040805162461bcd60e51b81526020600482015260196024820152600080516020612067833981519152604482015290519081900360640190fd5b50505b5060019392505050565b3360009081526020819052604090205481811015610d4f5760405162461bcd60e51b81526004018080602001828103825260218152602001806120eb6021913960400191505060405180910390fd5b336000818152602081815260408083208686039055805186815290519293926000805160206120ab833981519152929181900390910190a3604051600090339084908381818185875af1925050503d8060008114610dc9576040519150601f19603f3d011682016040523d82523d6000602084013e610dce565b606091505b5050905080610e12576040805162461bcd60e51b81526020600482015260196024820152600080516020612067833981519152604482015290519081900360640190fd5b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b601281565b6000467f00000000000000000000000000000000000000000000000000000000000000008114610e7857610e7381611f67565b610e9a565b7f00000000000000000000000000000000000000000000000000000000000000005b91505090565b60006001600160a01b03851615610f55573360009081526020819052604090205484811015610f005760405162461bcd60e51b81526004018080602001828103825260258152602001806120426025913960400191505060405180910390fd5b3360008181526020818152604080832089860390556001600160a01b038a168084529281902080548a019055805189815290519293926000805160206120ab833981519152929181900390910190a35061106a565b3360009081526020819052604090205484811015610fa45760405162461bcd60e51b81526004018080602001828103825260218152602001806120eb6021913960400191505060405180910390fd5b336000818152602081815260408083208986039055805189815290519293926000805160206120ab833981519152929181900390910190a3604051600090339087908381818185875af1925050503d806000811461101e576040519150601f19603f3d011682016040523d82523d6000602084013e611023565b606091505b5050905080611067576040805162461bcd60e51b81526020600482015260196024820152600080516020612067833981519152604482015290519081900360640190fd5b50505b846001600160a01b031663a4c0ed36338686866040518563ffffffff1660e01b815260040180856001600160a01b03168152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505095505050505050602060405180830381600087803b1580156110f557600080fd5b505af1158015611109573d6000803e3d6000fd5b505050506040513d602081101561111f57600080fd5b505195945050505050565b60006001600160a01b0385163014611189576040805162461bcd60e51b815260206004820152601c60248201527f574554483a20666c617368206d696e74206f6e6c792057455448313000000000604482015290519081900360640190fd5b6001600160701b038411156111cf5760405162461bcd60e51b81526004018080602001828103825260248152602001806120876024913960400191505060405180910390fd5b600380548501908190556001600160701b031015611234576040805162461bcd60e51b815260206004820152601f60248201527f574554483a20746f74616c206c6f616e206c696d697420657863656564656400604482015290519081900360640190fd5b6001600160a01b038616600081815260208181526040808320805489019055805188815290516000805160206120ab833981519152929181900390910190a37f0000000000000000000000000000000000000000000000000000000000000000866001600160a01b03166323e30c8b333088600089896040518763ffffffff1660e01b815260040180876001600160a01b03168152602001866001600160a01b03168152602001858152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050975050505050505050602060405180830381600087803b15801561133957600080fd5b505af115801561134d573d6000803e3d6000fd5b505050506040513d602081101561136357600080fd5b5051146113b7576040805162461bcd60e51b815260206004820152601760248201527f574554483a20666c617368206c6f616e206661696c6564000000000000000000604482015290519081900360640190fd5b6001600160a01b038616600090815260026020908152604080832030845290915290205460001981146114885784811015611439576040805162461bcd60e51b815260206004820152601f60248201527f574554483a2072657175657374206578636565647320616c6c6f77616e636500604482015290519081900360640190fd5b6001600160a01b0387166000818152600260209081526040808320308085529083529281902089860390819055815181815291519094926000805160206120cb833981519152928290030190a3505b6001600160a01b038716600090815260208190526040902054858110156114e05760405162461bcd60e51b81526004018080602001828103825260218152602001806120eb6021913960400191505060405180910390fd5b6001600160a01b0388166000818152602081815260408083208a8603905580518a815290519293926000805160206120ab833981519152929181900390910190a3505060038054859003905550600195945050505050565b6001600160a01b0383166000818152602081815260408083208054349081019091558151908152905192939284926000805160206120ab833981519152928290030190a3604051635260769b60e11b815233600482018181523460248401819052606060448501908152606485018790526001600160a01b0389169463a4c0ed36949389928992608401848480828437600081840152601f19601f82011690508083019250505095505050505050602060405180830381600087803b15801561160057600080fd5b505af1158015611614573d6000803e3d6000fd5b505050506040513d602081101561162a57600080fd5b5051949350505050565b60006001600160a01b038216301461164d57600061165a565b6003546001600160701b03035b92915050565b60006020819052908152604090205481565b60016020526000908152604090205481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60035481565b6001600160a01b0383163314611791576001600160a01b0383166000908152600260209081526040808320338452909152902054600019811461178f5781811015611740576040805162461bcd60e51b815260206004820152601f60248201527f574554483a2072657175657374206578636565647320616c6c6f77616e636500604482015290519081900360640190fd5b6001600160a01b0384166000818152600260209081526040808320338085529083529281902086860390819055815181815291519094926000805160206120cb833981519152928290030190a3505b505b6001600160a01b038316600090815260208190526040902054818110156117e95760405162461bcd60e51b81526004018080602001828103825260218152602001806120eb6021913960400191505060405180910390fd5b6001600160a01b0384166000818152602081815260408083208686039055805186815290519293926000805160206120ab833981519152929181900390910190a36040516000906001600160a01b0385169084908381818185875af1925050503d8060008114611875576040519150601f19603f3d011682016040523d82523d6000602084013e61187a565b606091505b50509050806118d0576040805162461bcd60e51b815260206004820152601b60248201527f574554483a204574686572207472616e73666572206661696c65640000000000604482015290519081900360640190fd5b5050505050565b6040518060400160405280600681526020016505745544831360d41b81525081565b60006001600160a01b038316156119ae5733600090815260208190526040902054828110156119595760405162461bcd60e51b81526004018080602001828103825260258152602001806120426025913960400191505060405180910390fd5b3360008181526020818152604080832087860390556001600160a01b03881680845292819020805488019055805187815290519293926000805160206120ab833981519152929181900390910190a350611ac3565b33600090815260208190526040902054828110156119fd5760405162461bcd60e51b81526004018080602001828103825260218152602001806120eb6021913960400191505060405180910390fd5b336000818152602081815260408083208786039055805187815290519293926000805160206120ab833981519152929181900390910190a3604051600090339085908381818185875af1925050503d8060008114611a77576040519150601f19603f3d011682016040523d82523d6000602084013e611a7c565b606091505b5050905080611ac0576040805162461bcd60e51b81526020600482015260196024820152600080516020612067833981519152604482015290519081900360640190fd5b50505b50600192915050565b6001600160a01b038116600081815260208181526040808320805434908101909155815190815290516000805160206120ab833981519152929181900390910190a350565b3360008181526002602090815260408083206001600160a01b038916808552908352818420889055815188815291519394909390926000805160206120cb833981519152928290030190a3846001600160a01b031662ba451f338686866040518563ffffffff1660e01b815260040180856001600160a01b03168152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505095505050505050602060405180830381600087803b1580156110f557600080fd5b7f000000000000000000000000000000000000000000000000000000000000000081565b33600081815260208181526040808320805434908101909155815190815290516000805160206120ab833981519152929181900390910190a3565b83421115611c91576040805162461bcd60e51b815260206004820152601460248201527315d155120e88115e1c1a5c9959081c195c9b5a5d60621b604482015290519081900360640190fd5b6001600160a01b0380881660008181526001602081815260408084208054938401905580517f00000000000000000000000000000000000000000000000000000000000000008184015280820195909552948b166060850152608084018a905260a084019190915260c08084018990528451808503909101815260e09093019093528151919092012046917f00000000000000000000000000000000000000000000000000000000000000008314611d5157611d4c83611f67565b611d73565b7f00000000000000000000000000000000000000000000000000000000000000005b82604051602001808061190160f01b81525060020183815260200182815260200192505050604051602081830303815290604052805190602001209050600060018288888860405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611e0c573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811615801590611e4257508a6001600160a01b0316816001600160a01b0316145b611e8a576040805162461bcd60e51b815260206004820152601460248201527315d155120e881a5b9d985b1a59081c195c9b5a5d60621b604482015290519081900360640190fd5b6001600160a01b03808c166000818152600260209081526040808320948f16808452948252918290208d905581518d815291516000805160206120cb8339815191529281900390910190a35050505050505050505050565b60006001600160a01b0383163014611f41576040805162461bcd60e51b815260206004820152601c60248201527f574554483a20666c617368206d696e74206f6e6c792057455448313000000000604482015290519081900360640190fd5b50600092915050565b600260209081526000928352604080842090915290825290205481565b60408051808201825260118152700577261707065642045746865722076313607c1b6020918201528151808301835260018152603160f81b9082015281517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818301527f5afefa023603a382ac22bd36644e004cefabbfaf61d4180000d8dc2b10c168b8818401527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6606082015260808101939093523060a0808501919091528251808503909101815260c090930190915281519101209056fe574554483a207472616e7366657220616d6f756e7420657863656564732062616c616e6365574554483a20455448207472616e73666572206661696c656400000000000000574554483a20696e646976696475616c206c6f616e206c696d6974206578636565646564ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925574554483a206275726e20616d6f756e7420657863656564732062616c616e6365a2646970667358221220e497f60997fd2507bfb8fa6439811f8ef6cd16d7649567fb9312dc8b97437f2f64736f6c63430007060033",
}

// WETH10ABI is the input ABI used to generate the binding from.
// Deprecated: Use WETH10MetaData.ABI instead.
var WETH10ABI = WETH10MetaData.ABI

// Deprecated: Use WETH10MetaData.Sigs instead.
// WETH10FuncSigs maps the 4-byte function signature to its string representation.
var WETH10FuncSigs = WETH10MetaData.Sigs

// WETH10Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WETH10MetaData.Bin instead.
var WETH10Bin = WETH10MetaData.Bin

// DeployWETH10 deploys a new Ethereum contract, binding an instance of WETH10 to it.
func DeployWETH10(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WETH10, error) {
	parsed, err := WETH10MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WETH10Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WETH10{WETH10Caller: WETH10Caller{contract: contract}, WETH10Transactor: WETH10Transactor{contract: contract}, WETH10Filterer: WETH10Filterer{contract: contract}}, nil
}

// WETH10 is an auto generated Go binding around an Ethereum contract.
type WETH10 struct {
	WETH10Caller     // Read-only binding to the contract
	WETH10Transactor // Write-only binding to the contract
	WETH10Filterer   // Log filterer for contract events
}

// WETH10Caller is an auto generated read-only Go binding around an Ethereum contract.
type WETH10Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETH10Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WETH10Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETH10Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WETH10Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETH10Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WETH10Session struct {
	Contract     *WETH10           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETH10CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WETH10CallerSession struct {
	Contract *WETH10Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WETH10TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WETH10TransactorSession struct {
	Contract     *WETH10Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETH10Raw is an auto generated low-level Go binding around an Ethereum contract.
type WETH10Raw struct {
	Contract *WETH10 // Generic contract binding to access the raw methods on
}

// WETH10CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WETH10CallerRaw struct {
	Contract *WETH10Caller // Generic read-only contract binding to access the raw methods on
}

// WETH10TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WETH10TransactorRaw struct {
	Contract *WETH10Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWETH10 creates a new instance of WETH10, bound to a specific deployed contract.
func NewWETH10(address common.Address, backend bind.ContractBackend) (*WETH10, error) {
	contract, err := bindWETH10(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETH10{WETH10Caller: WETH10Caller{contract: contract}, WETH10Transactor: WETH10Transactor{contract: contract}, WETH10Filterer: WETH10Filterer{contract: contract}}, nil
}

// NewWETH10Caller creates a new read-only instance of WETH10, bound to a specific deployed contract.
func NewWETH10Caller(address common.Address, caller bind.ContractCaller) (*WETH10Caller, error) {
	contract, err := bindWETH10(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETH10Caller{contract: contract}, nil
}

// NewWETH10Transactor creates a new write-only instance of WETH10, bound to a specific deployed contract.
func NewWETH10Transactor(address common.Address, transactor bind.ContractTransactor) (*WETH10Transactor, error) {
	contract, err := bindWETH10(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETH10Transactor{contract: contract}, nil
}

// NewWETH10Filterer creates a new log filterer instance of WETH10, bound to a specific deployed contract.
func NewWETH10Filterer(address common.Address, filterer bind.ContractFilterer) (*WETH10Filterer, error) {
	contract, err := bindWETH10(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETH10Filterer{contract: contract}, nil
}

// bindWETH10 binds a generic wrapper to an already deployed contract.
func bindWETH10(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WETH10ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETH10 *WETH10Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH10.Contract.WETH10Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETH10 *WETH10Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH10.Contract.WETH10Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETH10 *WETH10Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH10.Contract.WETH10Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETH10 *WETH10CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH10.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETH10 *WETH10TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH10.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETH10 *WETH10TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH10.Contract.contract.Transact(opts, method, params...)
}

// CALLBACKSUCCESS is a free data retrieval call binding the contract method 0x8237e538.
//
// Solidity: function CALLBACK_SUCCESS() view returns(bytes32)
func (_WETH10 *WETH10Caller) CALLBACKSUCCESS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "CALLBACK_SUCCESS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CALLBACKSUCCESS is a free data retrieval call binding the contract method 0x8237e538.
//
// Solidity: function CALLBACK_SUCCESS() view returns(bytes32)
func (_WETH10 *WETH10Session) CALLBACKSUCCESS() ([32]byte, error) {
	return _WETH10.Contract.CALLBACKSUCCESS(&_WETH10.CallOpts)
}

// CALLBACKSUCCESS is a free data retrieval call binding the contract method 0x8237e538.
//
// Solidity: function CALLBACK_SUCCESS() view returns(bytes32)
func (_WETH10 *WETH10CallerSession) CALLBACKSUCCESS() ([32]byte, error) {
	return _WETH10.Contract.CALLBACKSUCCESS(&_WETH10.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WETH10 *WETH10Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WETH10 *WETH10Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _WETH10.Contract.DOMAINSEPARATOR(&_WETH10.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WETH10 *WETH10CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WETH10.Contract.DOMAINSEPARATOR(&_WETH10.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_WETH10 *WETH10Caller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_WETH10 *WETH10Session) PERMITTYPEHASH() ([32]byte, error) {
	return _WETH10.Contract.PERMITTYPEHASH(&_WETH10.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_WETH10 *WETH10CallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _WETH10.Contract.PERMITTYPEHASH(&_WETH10.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WETH10 *WETH10Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WETH10 *WETH10Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WETH10.Contract.Allowance(&_WETH10.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WETH10 *WETH10CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WETH10.Contract.Allowance(&_WETH10.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WETH10 *WETH10Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WETH10 *WETH10Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WETH10.Contract.BalanceOf(&_WETH10.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WETH10 *WETH10CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WETH10.Contract.BalanceOf(&_WETH10.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETH10 *WETH10Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETH10 *WETH10Session) Decimals() (uint8, error) {
	return _WETH10.Contract.Decimals(&_WETH10.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETH10 *WETH10CallerSession) Decimals() (uint8, error) {
	return _WETH10.Contract.Decimals(&_WETH10.CallOpts)
}

// DeploymentChainId is a free data retrieval call binding the contract method 0xcd0d0096.
//
// Solidity: function deploymentChainId() view returns(uint256)
func (_WETH10 *WETH10Caller) DeploymentChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "deploymentChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeploymentChainId is a free data retrieval call binding the contract method 0xcd0d0096.
//
// Solidity: function deploymentChainId() view returns(uint256)
func (_WETH10 *WETH10Session) DeploymentChainId() (*big.Int, error) {
	return _WETH10.Contract.DeploymentChainId(&_WETH10.CallOpts)
}

// DeploymentChainId is a free data retrieval call binding the contract method 0xcd0d0096.
//
// Solidity: function deploymentChainId() view returns(uint256)
func (_WETH10 *WETH10CallerSession) DeploymentChainId() (*big.Int, error) {
	return _WETH10.Contract.DeploymentChainId(&_WETH10.CallOpts)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 ) view returns(uint256)
func (_WETH10 *WETH10Caller) FlashFee(opts *bind.CallOpts, token common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "flashFee", token, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 ) view returns(uint256)
func (_WETH10 *WETH10Session) FlashFee(token common.Address, arg1 *big.Int) (*big.Int, error) {
	return _WETH10.Contract.FlashFee(&_WETH10.CallOpts, token, arg1)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 ) view returns(uint256)
func (_WETH10 *WETH10CallerSession) FlashFee(token common.Address, arg1 *big.Int) (*big.Int, error) {
	return _WETH10.Contract.FlashFee(&_WETH10.CallOpts, token, arg1)
}

// FlashMinted is a free data retrieval call binding the contract method 0x8b28d32f.
//
// Solidity: function flashMinted() view returns(uint256)
func (_WETH10 *WETH10Caller) FlashMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "flashMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashMinted is a free data retrieval call binding the contract method 0x8b28d32f.
//
// Solidity: function flashMinted() view returns(uint256)
func (_WETH10 *WETH10Session) FlashMinted() (*big.Int, error) {
	return _WETH10.Contract.FlashMinted(&_WETH10.CallOpts)
}

// FlashMinted is a free data retrieval call binding the contract method 0x8b28d32f.
//
// Solidity: function flashMinted() view returns(uint256)
func (_WETH10 *WETH10CallerSession) FlashMinted() (*big.Int, error) {
	return _WETH10.Contract.FlashMinted(&_WETH10.CallOpts)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_WETH10 *WETH10Caller) MaxFlashLoan(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "maxFlashLoan", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_WETH10 *WETH10Session) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _WETH10.Contract.MaxFlashLoan(&_WETH10.CallOpts, token)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_WETH10 *WETH10CallerSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _WETH10.Contract.MaxFlashLoan(&_WETH10.CallOpts, token)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETH10 *WETH10Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETH10 *WETH10Session) Name() (string, error) {
	return _WETH10.Contract.Name(&_WETH10.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETH10 *WETH10CallerSession) Name() (string, error) {
	return _WETH10.Contract.Name(&_WETH10.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_WETH10 *WETH10Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_WETH10 *WETH10Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _WETH10.Contract.Nonces(&_WETH10.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_WETH10 *WETH10CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _WETH10.Contract.Nonces(&_WETH10.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETH10 *WETH10Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETH10 *WETH10Session) Symbol() (string, error) {
	return _WETH10.Contract.Symbol(&_WETH10.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETH10 *WETH10CallerSession) Symbol() (string, error) {
	return _WETH10.Contract.Symbol(&_WETH10.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH10 *WETH10Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETH10.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH10 *WETH10Session) TotalSupply() (*big.Int, error) {
	return _WETH10.Contract.TotalSupply(&_WETH10.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH10 *WETH10CallerSession) TotalSupply() (*big.Int, error) {
	return _WETH10.Contract.TotalSupply(&_WETH10.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WETH10 *WETH10Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WETH10 *WETH10Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.Approve(&_WETH10.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WETH10 *WETH10TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.Approve(&_WETH10.TransactOpts, spender, value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_WETH10 *WETH10Transactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "approveAndCall", spender, value, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_WETH10 *WETH10Session) ApproveAndCall(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WETH10.Contract.ApproveAndCall(&_WETH10.TransactOpts, spender, value, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_WETH10 *WETH10TransactorSession) ApproveAndCall(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WETH10.Contract.ApproveAndCall(&_WETH10.TransactOpts, spender, value, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETH10 *WETH10Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETH10 *WETH10Session) Deposit() (*types.Transaction, error) {
	return _WETH10.Contract.Deposit(&_WETH10.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETH10 *WETH10TransactorSession) Deposit() (*types.Transaction, error) {
	return _WETH10.Contract.Deposit(&_WETH10.TransactOpts)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_WETH10 *WETH10Transactor) DepositTo(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "depositTo", to)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_WETH10 *WETH10Session) DepositTo(to common.Address) (*types.Transaction, error) {
	return _WETH10.Contract.DepositTo(&_WETH10.TransactOpts, to)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_WETH10 *WETH10TransactorSession) DepositTo(to common.Address) (*types.Transaction, error) {
	return _WETH10.Contract.DepositTo(&_WETH10.TransactOpts, to)
}

// DepositToAndCall is a paid mutator transaction binding the contract method 0x5ddb7d7e.
//
// Solidity: function depositToAndCall(address to, bytes data) payable returns(bool success)
func (_WETH10 *WETH10Transactor) DepositToAndCall(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "depositToAndCall", to, data)
}

// DepositToAndCall is a paid mutator transaction binding the contract method 0x5ddb7d7e.
//
// Solidity: function depositToAndCall(address to, bytes data) payable returns(bool success)
func (_WETH10 *WETH10Session) DepositToAndCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _WETH10.Contract.DepositToAndCall(&_WETH10.TransactOpts, to, data)
}

// DepositToAndCall is a paid mutator transaction binding the contract method 0x5ddb7d7e.
//
// Solidity: function depositToAndCall(address to, bytes data) payable returns(bool success)
func (_WETH10 *WETH10TransactorSession) DepositToAndCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _WETH10.Contract.DepositToAndCall(&_WETH10.TransactOpts, to, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 value, bytes data) returns(bool)
func (_WETH10 *WETH10Transactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "flashLoan", receiver, token, value, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 value, bytes data) returns(bool)
func (_WETH10 *WETH10Session) FlashLoan(receiver common.Address, token common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WETH10.Contract.FlashLoan(&_WETH10.TransactOpts, receiver, token, value, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 value, bytes data) returns(bool)
func (_WETH10 *WETH10TransactorSession) FlashLoan(receiver common.Address, token common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WETH10.Contract.FlashLoan(&_WETH10.TransactOpts, receiver, token, value, data)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WETH10 *WETH10Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WETH10 *WETH10Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WETH10.Contract.Permit(&_WETH10.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WETH10 *WETH10TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WETH10.Contract.Permit(&_WETH10.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WETH10 *WETH10Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WETH10 *WETH10Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.Transfer(&_WETH10.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WETH10 *WETH10TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.Transfer(&_WETH10.TransactOpts, to, value)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_WETH10 *WETH10Transactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "transferAndCall", to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_WETH10 *WETH10Session) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WETH10.Contract.TransferAndCall(&_WETH10.TransactOpts, to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_WETH10 *WETH10TransactorSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WETH10.Contract.TransferAndCall(&_WETH10.TransactOpts, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WETH10 *WETH10Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WETH10 *WETH10Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.TransferFrom(&_WETH10.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WETH10 *WETH10TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.TransferFrom(&_WETH10.TransactOpts, from, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_WETH10 *WETH10Transactor) Withdraw(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "withdraw", value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_WETH10 *WETH10Session) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.Withdraw(&_WETH10.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_WETH10 *WETH10TransactorSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.Withdraw(&_WETH10.TransactOpts, value)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9555a942.
//
// Solidity: function withdrawFrom(address from, address to, uint256 value) returns()
func (_WETH10 *WETH10Transactor) WithdrawFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "withdrawFrom", from, to, value)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9555a942.
//
// Solidity: function withdrawFrom(address from, address to, uint256 value) returns()
func (_WETH10 *WETH10Session) WithdrawFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.WithdrawFrom(&_WETH10.TransactOpts, from, to, value)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9555a942.
//
// Solidity: function withdrawFrom(address from, address to, uint256 value) returns()
func (_WETH10 *WETH10TransactorSession) WithdrawFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.WithdrawFrom(&_WETH10.TransactOpts, from, to, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_WETH10 *WETH10Transactor) WithdrawTo(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.contract.Transact(opts, "withdrawTo", to, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_WETH10 *WETH10Session) WithdrawTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.WithdrawTo(&_WETH10.TransactOpts, to, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_WETH10 *WETH10TransactorSession) WithdrawTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WETH10.Contract.WithdrawTo(&_WETH10.TransactOpts, to, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETH10 *WETH10Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH10.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETH10 *WETH10Session) Receive() (*types.Transaction, error) {
	return _WETH10.Contract.Receive(&_WETH10.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETH10 *WETH10TransactorSession) Receive() (*types.Transaction, error) {
	return _WETH10.Contract.Receive(&_WETH10.TransactOpts)
}

// WETH10ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WETH10 contract.
type WETH10ApprovalIterator struct {
	Event *WETH10Approval // Event containing the contract specifics and raw log

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
func (it *WETH10ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETH10Approval)
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
		it.Event = new(WETH10Approval)
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
func (it *WETH10ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETH10ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETH10Approval represents a Approval event raised by the WETH10 contract.
type WETH10Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETH10 *WETH10Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WETH10ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WETH10.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WETH10ApprovalIterator{contract: _WETH10.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETH10 *WETH10Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WETH10Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WETH10.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETH10Approval)
				if err := _WETH10.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WETH10 *WETH10Filterer) ParseApproval(log types.Log) (*WETH10Approval, error) {
	event := new(WETH10Approval)
	if err := _WETH10.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WETH10TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WETH10 contract.
type WETH10TransferIterator struct {
	Event *WETH10Transfer // Event containing the contract specifics and raw log

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
func (it *WETH10TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETH10Transfer)
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
		it.Event = new(WETH10Transfer)
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
func (it *WETH10TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETH10TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETH10Transfer represents a Transfer event raised by the WETH10 contract.
type WETH10Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETH10 *WETH10Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WETH10TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WETH10.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WETH10TransferIterator{contract: _WETH10.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETH10 *WETH10Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WETH10Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WETH10.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETH10Transfer)
				if err := _WETH10.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WETH10 *WETH10Filterer) ParseTransfer(log types.Log) (*WETH10Transfer, error) {
	event := new(WETH10Transfer)
	if err := _WETH10.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
