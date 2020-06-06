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

// DAICashABI is the input ABI used to generate the binding from.
const DAICashABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETERNAL_APPROVAL_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractIDaiJoin\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daiVat\",\"outputs\":[{\"internalType\":\"contractIDaiVat\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"faucet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"_augur\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"joinBurn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"joinMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DAICash is an auto generated Go binding around an Ethereum contract.
type DAICash struct {
	DAICashCaller     // Read-only binding to the contract
	DAICashTransactor // Write-only binding to the contract
	DAICashFilterer   // Log filterer for contract events
}

// DAICashCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAICashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAICashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAICashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAICashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAICashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAICashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAICashSession struct {
	Contract     *DAICash          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAICashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAICashCallerSession struct {
	Contract *DAICashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DAICashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAICashTransactorSession struct {
	Contract     *DAICashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DAICashRaw is an auto generated low-level Go binding around an Ethereum contract.
type DAICashRaw struct {
	Contract *DAICash // Generic contract binding to access the raw methods on
}

// DAICashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAICashCallerRaw struct {
	Contract *DAICashCaller // Generic read-only contract binding to access the raw methods on
}

// DAICashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAICashTransactorRaw struct {
	Contract *DAICashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAICash creates a new instance of DAICash, bound to a specific deployed contract.
func NewDAICash(address common.Address, backend bind.ContractBackend) (*DAICash, error) {
	contract, err := bindDAICash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAICash{DAICashCaller: DAICashCaller{contract: contract}, DAICashTransactor: DAICashTransactor{contract: contract}, DAICashFilterer: DAICashFilterer{contract: contract}}, nil
}

// NewDAICashCaller creates a new read-only instance of DAICash, bound to a specific deployed contract.
func NewDAICashCaller(address common.Address, caller bind.ContractCaller) (*DAICashCaller, error) {
	contract, err := bindDAICash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAICashCaller{contract: contract}, nil
}

// NewDAICashTransactor creates a new write-only instance of DAICash, bound to a specific deployed contract.
func NewDAICashTransactor(address common.Address, transactor bind.ContractTransactor) (*DAICashTransactor, error) {
	contract, err := bindDAICash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAICashTransactor{contract: contract}, nil
}

// NewDAICashFilterer creates a new log filterer instance of DAICash, bound to a specific deployed contract.
func NewDAICashFilterer(address common.Address, filterer bind.ContractFilterer) (*DAICashFilterer, error) {
	contract, err := bindDAICash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAICashFilterer{contract: contract}, nil
}

// bindDAICash binds a generic wrapper to an already deployed contract.
func bindDAICash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DAICashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAICash *DAICashRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DAICash.Contract.DAICashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAICash *DAICashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAICash.Contract.DAICashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAICash *DAICashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAICash.Contract.DAICashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAICash *DAICashCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DAICash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAICash *DAICashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAICash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAICash *DAICashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAICash.Contract.contract.Transact(opts, method, params...)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() view returns(uint256)
func (_DAICash *DAICashCaller) ETERNALAPPROVALVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "ETERNAL_APPROVAL_VALUE")
	return *ret0, err
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() view returns(uint256)
func (_DAICash *DAICashSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _DAICash.Contract.ETERNALAPPROVALVALUE(&_DAICash.CallOpts)
}

// ETERNALAPPROVALVALUE is a free data retrieval call binding the contract method 0x634eaff1.
//
// Solidity: function ETERNAL_APPROVAL_VALUE() view returns(uint256)
func (_DAICash *DAICashCallerSession) ETERNALAPPROVALVALUE() (*big.Int, error) {
	return _DAICash.Contract.ETERNALAPPROVALVALUE(&_DAICash.CallOpts)
}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_DAICash *DAICashCaller) RAY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "RAY")
	return *ret0, err
}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_DAICash *DAICashSession) RAY() (*big.Int, error) {
	return _DAICash.Contract.RAY(&_DAICash.CallOpts)
}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_DAICash *DAICashCallerSession) RAY() (*big.Int, error) {
	return _DAICash.Contract.RAY(&_DAICash.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_DAICash *DAICashCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_DAICash *DAICashSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DAICash.Contract.Allowance(&_DAICash.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_DAICash *DAICashCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DAICash.Contract.Allowance(&_DAICash.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_DAICash *DAICashCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_DAICash *DAICashSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DAICash.Contract.BalanceOf(&_DAICash.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_DAICash *DAICashCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DAICash.Contract.BalanceOf(&_DAICash.CallOpts, _owner)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_DAICash *DAICashCaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "daiJoin")
	return *ret0, err
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_DAICash *DAICashSession) DaiJoin() (common.Address, error) {
	return _DAICash.Contract.DaiJoin(&_DAICash.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_DAICash *DAICashCallerSession) DaiJoin() (common.Address, error) {
	return _DAICash.Contract.DaiJoin(&_DAICash.CallOpts)
}

// DaiVat is a free data retrieval call binding the contract method 0x3bd831f4.
//
// Solidity: function daiVat() view returns(address)
func (_DAICash *DAICashCaller) DaiVat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "daiVat")
	return *ret0, err
}

// DaiVat is a free data retrieval call binding the contract method 0x3bd831f4.
//
// Solidity: function daiVat() view returns(address)
func (_DAICash *DAICashSession) DaiVat() (common.Address, error) {
	return _DAICash.Contract.DaiVat(&_DAICash.CallOpts)
}

// DaiVat is a free data retrieval call binding the contract method 0x3bd831f4.
//
// Solidity: function daiVat() view returns(address)
func (_DAICash *DAICashCallerSession) DaiVat() (common.Address, error) {
	return _DAICash.Contract.DaiVat(&_DAICash.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DAICash *DAICashCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DAICash *DAICashSession) Decimals() (uint8, error) {
	return _DAICash.Contract.Decimals(&_DAICash.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DAICash *DAICashCallerSession) Decimals() (uint8, error) {
	return _DAICash.Contract.Decimals(&_DAICash.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(bytes32)
func (_DAICash *DAICashCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(bytes32)
func (_DAICash *DAICashSession) GetTypeName() ([32]byte, error) {
	return _DAICash.Contract.GetTypeName(&_DAICash.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(bytes32)
func (_DAICash *DAICashCallerSession) GetTypeName() ([32]byte, error) {
	return _DAICash.Contract.GetTypeName(&_DAICash.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DAICash *DAICashCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DAICash *DAICashSession) Name() (string, error) {
	return _DAICash.Contract.Name(&_DAICash.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DAICash *DAICashCallerSession) Name() (string, error) {
	return _DAICash.Contract.Name(&_DAICash.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DAICash *DAICashCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DAICash *DAICashSession) Symbol() (string, error) {
	return _DAICash.Contract.Symbol(&_DAICash.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DAICash *DAICashCallerSession) Symbol() (string, error) {
	return _DAICash.Contract.Symbol(&_DAICash.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DAICash *DAICashCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DAICash *DAICashSession) TotalSupply() (*big.Int, error) {
	return _DAICash.Contract.TotalSupply(&_DAICash.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DAICash *DAICashCallerSession) TotalSupply() (*big.Int, error) {
	return _DAICash.Contract.TotalSupply(&_DAICash.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_DAICash *DAICashCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DAICash.contract.Call(opts, out, "wards", arg0)
	return *ret0, err
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_DAICash *DAICashSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _DAICash.Contract.Wards(&_DAICash.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_DAICash *DAICashCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _DAICash.Contract.Wards(&_DAICash.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_DAICash *DAICashTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DAICash.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_DAICash *DAICashSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.Approve(&_DAICash.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_DAICash *DAICashTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.Approve(&_DAICash.TransactOpts, _spender, _amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_DAICash *DAICashTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _DAICash.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_DAICash *DAICashSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.DecreaseApproval(&_DAICash.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_DAICash *DAICashTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.DecreaseApproval(&_DAICash.TransactOpts, _spender, _subtractedValue)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(uint256 _amount) returns(bool)
func (_DAICash *DAICashTransactor) Faucet(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _DAICash.contract.Transact(opts, "faucet", _amount)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(uint256 _amount) returns(bool)
func (_DAICash *DAICashSession) Faucet(_amount *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.Faucet(&_DAICash.TransactOpts, _amount)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(uint256 _amount) returns(bool)
func (_DAICash *DAICashTransactorSession) Faucet(_amount *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.Faucet(&_DAICash.TransactOpts, _amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_DAICash *DAICashTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _DAICash.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_DAICash *DAICashSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.IncreaseApproval(&_DAICash.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_DAICash *DAICashTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.IncreaseApproval(&_DAICash.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _augur) returns(bool)
func (_DAICash *DAICashTransactor) Initialize(opts *bind.TransactOpts, _augur common.Address) (*types.Transaction, error) {
	return _DAICash.contract.Transact(opts, "initialize", _augur)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _augur) returns(bool)
func (_DAICash *DAICashSession) Initialize(_augur common.Address) (*types.Transaction, error) {
	return _DAICash.Contract.Initialize(&_DAICash.TransactOpts, _augur)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _augur) returns(bool)
func (_DAICash *DAICashTransactorSession) Initialize(_augur common.Address) (*types.Transaction, error) {
	return _DAICash.Contract.Initialize(&_DAICash.TransactOpts, _augur)
}

// JoinBurn is a paid mutator transaction binding the contract method 0xf11f299e.
//
// Solidity: function joinBurn(address usr, uint256 wad) returns(bool)
func (_DAICash *DAICashTransactor) JoinBurn(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DAICash.contract.Transact(opts, "joinBurn", usr, wad)
}

// JoinBurn is a paid mutator transaction binding the contract method 0xf11f299e.
//
// Solidity: function joinBurn(address usr, uint256 wad) returns(bool)
func (_DAICash *DAICashSession) JoinBurn(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.JoinBurn(&_DAICash.TransactOpts, usr, wad)
}

// JoinBurn is a paid mutator transaction binding the contract method 0xf11f299e.
//
// Solidity: function joinBurn(address usr, uint256 wad) returns(bool)
func (_DAICash *DAICashTransactorSession) JoinBurn(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.JoinBurn(&_DAICash.TransactOpts, usr, wad)
}

// JoinMint is a paid mutator transaction binding the contract method 0xddc5972e.
//
// Solidity: function joinMint(address usr, uint256 wad) returns(bool)
func (_DAICash *DAICashTransactor) JoinMint(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DAICash.contract.Transact(opts, "joinMint", usr, wad)
}

// JoinMint is a paid mutator transaction binding the contract method 0xddc5972e.
//
// Solidity: function joinMint(address usr, uint256 wad) returns(bool)
func (_DAICash *DAICashSession) JoinMint(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.JoinMint(&_DAICash.TransactOpts, usr, wad)
}

// JoinMint is a paid mutator transaction binding the contract method 0xddc5972e.
//
// Solidity: function joinMint(address usr, uint256 wad) returns(bool)
func (_DAICash *DAICashTransactorSession) JoinMint(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.JoinMint(&_DAICash.TransactOpts, usr, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_DAICash *DAICashTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DAICash.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_DAICash *DAICashSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.Transfer(&_DAICash.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_DAICash *DAICashTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.Transfer(&_DAICash.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_DAICash *DAICashTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DAICash.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_DAICash *DAICashSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.TransferFrom(&_DAICash.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_DAICash *DAICashTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DAICash.Contract.TransferFrom(&_DAICash.TransactOpts, _from, _to, _amount)
}

// DAICashApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DAICash contract.
type DAICashApprovalIterator struct {
	Event *DAICashApproval // Event containing the contract specifics and raw log

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
func (it *DAICashApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAICashApproval)
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
		it.Event = new(DAICashApproval)
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
func (it *DAICashApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAICashApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAICashApproval represents a Approval event raised by the DAICash contract.
type DAICashApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DAICash *DAICashFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DAICashApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DAICash.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DAICashApprovalIterator{contract: _DAICash.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DAICash *DAICashFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DAICashApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DAICash.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAICashApproval)
				if err := _DAICash.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DAICash *DAICashFilterer) ParseApproval(log types.Log) (*DAICashApproval, error) {
	event := new(DAICashApproval)
	if err := _DAICash.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DAICashBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the DAICash contract.
type DAICashBurnIterator struct {
	Event *DAICashBurn // Event containing the contract specifics and raw log

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
func (it *DAICashBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAICashBurn)
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
		it.Event = new(DAICashBurn)
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
func (it *DAICashBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAICashBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAICashBurn represents a Burn event raised by the DAICash contract.
type DAICashBurn struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed target, uint256 value)
func (_DAICash *DAICashFilterer) FilterBurn(opts *bind.FilterOpts, target []common.Address) (*DAICashBurnIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _DAICash.contract.FilterLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return &DAICashBurnIterator{contract: _DAICash.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed target, uint256 value)
func (_DAICash *DAICashFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *DAICashBurn, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _DAICash.contract.WatchLogs(opts, "Burn", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAICashBurn)
				if err := _DAICash.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed target, uint256 value)
func (_DAICash *DAICashFilterer) ParseBurn(log types.Log) (*DAICashBurn, error) {
	event := new(DAICashBurn)
	if err := _DAICash.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DAICashMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the DAICash contract.
type DAICashMintIterator struct {
	Event *DAICashMint // Event containing the contract specifics and raw log

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
func (it *DAICashMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAICashMint)
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
		it.Event = new(DAICashMint)
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
func (it *DAICashMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAICashMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAICashMint represents a Mint event raised by the DAICash contract.
type DAICashMint struct {
	Target common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed target, uint256 value)
func (_DAICash *DAICashFilterer) FilterMint(opts *bind.FilterOpts, target []common.Address) (*DAICashMintIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _DAICash.contract.FilterLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return &DAICashMintIterator{contract: _DAICash.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed target, uint256 value)
func (_DAICash *DAICashFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *DAICashMint, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _DAICash.contract.WatchLogs(opts, "Mint", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAICashMint)
				if err := _DAICash.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed target, uint256 value)
func (_DAICash *DAICashFilterer) ParseMint(log types.Log) (*DAICashMint, error) {
	event := new(DAICashMint)
	if err := _DAICash.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DAICashTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DAICash contract.
type DAICashTransferIterator struct {
	Event *DAICashTransfer // Event containing the contract specifics and raw log

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
func (it *DAICashTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAICashTransfer)
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
		it.Event = new(DAICashTransfer)
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
func (it *DAICashTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAICashTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAICashTransfer represents a Transfer event raised by the DAICash contract.
type DAICashTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DAICash *DAICashFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DAICashTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DAICash.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DAICashTransferIterator{contract: _DAICash.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DAICash *DAICashFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DAICashTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DAICash.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAICashTransfer)
				if err := _DAICash.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DAICash *DAICashFilterer) ParseTransfer(log types.Log) (*DAICashTransfer, error) {
	event := new(DAICashTransfer)
	if err := _DAICash.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
