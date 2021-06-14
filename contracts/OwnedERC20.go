// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// OwnedERC20ABI is the input ABI used to generate the binding from.
const OwnedERC20ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"trustedBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"trustedBurnAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"trustedMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"trustedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnedERC20 is an auto generated Go binding around an Ethereum contract.
type OwnedERC20 struct {
	OwnedERC20Caller     // Read-only binding to the contract
	OwnedERC20Transactor // Write-only binding to the contract
	OwnedERC20Filterer   // Log filterer for contract events
}

// OwnedERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedERC20Session struct {
	Contract     *OwnedERC20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedERC20CallerSession struct {
	Contract *OwnedERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OwnedERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedERC20TransactorSession struct {
	Contract     *OwnedERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OwnedERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedERC20Raw struct {
	Contract *OwnedERC20 // Generic contract binding to access the raw methods on
}

// OwnedERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedERC20CallerRaw struct {
	Contract *OwnedERC20Caller // Generic read-only contract binding to access the raw methods on
}

// OwnedERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedERC20TransactorRaw struct {
	Contract *OwnedERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnedERC20 creates a new instance of OwnedERC20, bound to a specific deployed contract.
func NewOwnedERC20(address common.Address, backend bind.ContractBackend) (*OwnedERC20, error) {
	contract, err := bindOwnedERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnedERC20{OwnedERC20Caller: OwnedERC20Caller{contract: contract}, OwnedERC20Transactor: OwnedERC20Transactor{contract: contract}, OwnedERC20Filterer: OwnedERC20Filterer{contract: contract}}, nil
}

// NewOwnedERC20Caller creates a new read-only instance of OwnedERC20, bound to a specific deployed contract.
func NewOwnedERC20Caller(address common.Address, caller bind.ContractCaller) (*OwnedERC20Caller, error) {
	contract, err := bindOwnedERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedERC20Caller{contract: contract}, nil
}

// NewOwnedERC20Transactor creates a new write-only instance of OwnedERC20, bound to a specific deployed contract.
func NewOwnedERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*OwnedERC20Transactor, error) {
	contract, err := bindOwnedERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedERC20Transactor{contract: contract}, nil
}

// NewOwnedERC20Filterer creates a new log filterer instance of OwnedERC20, bound to a specific deployed contract.
func NewOwnedERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*OwnedERC20Filterer, error) {
	contract, err := bindOwnedERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedERC20Filterer{contract: contract}, nil
}

// bindOwnedERC20 binds a generic wrapper to an already deployed contract.
func bindOwnedERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnedERC20 *OwnedERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OwnedERC20.Contract.OwnedERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnedERC20 *OwnedERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnedERC20.Contract.OwnedERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnedERC20 *OwnedERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnedERC20.Contract.OwnedERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnedERC20 *OwnedERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OwnedERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnedERC20 *OwnedERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnedERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnedERC20 *OwnedERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnedERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OwnedERC20 *OwnedERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OwnedERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OwnedERC20 *OwnedERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _OwnedERC20.Contract.Allowance(&_OwnedERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OwnedERC20 *OwnedERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _OwnedERC20.Contract.Allowance(&_OwnedERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OwnedERC20 *OwnedERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OwnedERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OwnedERC20 *OwnedERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _OwnedERC20.Contract.BalanceOf(&_OwnedERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OwnedERC20 *OwnedERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _OwnedERC20.Contract.BalanceOf(&_OwnedERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OwnedERC20 *OwnedERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _OwnedERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OwnedERC20 *OwnedERC20Session) Decimals() (uint8, error) {
	return _OwnedERC20.Contract.Decimals(&_OwnedERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OwnedERC20 *OwnedERC20CallerSession) Decimals() (uint8, error) {
	return _OwnedERC20.Contract.Decimals(&_OwnedERC20.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_OwnedERC20 *OwnedERC20Caller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OwnedERC20.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_OwnedERC20 *OwnedERC20Session) GetOwner() (common.Address, error) {
	return _OwnedERC20.Contract.GetOwner(&_OwnedERC20.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_OwnedERC20 *OwnedERC20CallerSession) GetOwner() (common.Address, error) {
	return _OwnedERC20.Contract.GetOwner(&_OwnedERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OwnedERC20 *OwnedERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OwnedERC20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OwnedERC20 *OwnedERC20Session) Name() (string, error) {
	return _OwnedERC20.Contract.Name(&_OwnedERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OwnedERC20 *OwnedERC20CallerSession) Name() (string, error) {
	return _OwnedERC20.Contract.Name(&_OwnedERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OwnedERC20 *OwnedERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OwnedERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OwnedERC20 *OwnedERC20Session) Symbol() (string, error) {
	return _OwnedERC20.Contract.Symbol(&_OwnedERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OwnedERC20 *OwnedERC20CallerSession) Symbol() (string, error) {
	return _OwnedERC20.Contract.Symbol(&_OwnedERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OwnedERC20 *OwnedERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OwnedERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OwnedERC20 *OwnedERC20Session) TotalSupply() (*big.Int, error) {
	return _OwnedERC20.Contract.TotalSupply(&_OwnedERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OwnedERC20 *OwnedERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _OwnedERC20.Contract.TotalSupply(&_OwnedERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_OwnedERC20 *OwnedERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_OwnedERC20 *OwnedERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.Approve(&_OwnedERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_OwnedERC20 *OwnedERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.Approve(&_OwnedERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OwnedERC20 *OwnedERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OwnedERC20 *OwnedERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.DecreaseAllowance(&_OwnedERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OwnedERC20 *OwnedERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.DecreaseAllowance(&_OwnedERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OwnedERC20 *OwnedERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OwnedERC20 *OwnedERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.IncreaseAllowance(&_OwnedERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OwnedERC20 *OwnedERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.IncreaseAllowance(&_OwnedERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_OwnedERC20 *OwnedERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_OwnedERC20 *OwnedERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.Transfer(&_OwnedERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_OwnedERC20 *OwnedERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.Transfer(&_OwnedERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_OwnedERC20 *OwnedERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_OwnedERC20 *OwnedERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TransferFrom(&_OwnedERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_OwnedERC20 *OwnedERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TransferFrom(&_OwnedERC20.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_OwnedERC20 *OwnedERC20Transactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _OwnedERC20.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_OwnedERC20 *OwnedERC20Session) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TransferOwnership(&_OwnedERC20.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_OwnedERC20 *OwnedERC20TransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TransferOwnership(&_OwnedERC20.TransactOpts, _newOwner)
}

// TrustedBurn is a paid mutator transaction binding the contract method 0x42986e13.
//
// Solidity: function trustedBurn(address _target, uint256 _amount) returns()
func (_OwnedERC20 *OwnedERC20Transactor) TrustedBurn(opts *bind.TransactOpts, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.contract.Transact(opts, "trustedBurn", _target, _amount)
}

// TrustedBurn is a paid mutator transaction binding the contract method 0x42986e13.
//
// Solidity: function trustedBurn(address _target, uint256 _amount) returns()
func (_OwnedERC20 *OwnedERC20Session) TrustedBurn(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TrustedBurn(&_OwnedERC20.TransactOpts, _target, _amount)
}

// TrustedBurn is a paid mutator transaction binding the contract method 0x42986e13.
//
// Solidity: function trustedBurn(address _target, uint256 _amount) returns()
func (_OwnedERC20 *OwnedERC20TransactorSession) TrustedBurn(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TrustedBurn(&_OwnedERC20.TransactOpts, _target, _amount)
}

// TrustedBurnAll is a paid mutator transaction binding the contract method 0x71297784.
//
// Solidity: function trustedBurnAll(address _target) returns(uint256)
func (_OwnedERC20 *OwnedERC20Transactor) TrustedBurnAll(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _OwnedERC20.contract.Transact(opts, "trustedBurnAll", _target)
}

// TrustedBurnAll is a paid mutator transaction binding the contract method 0x71297784.
//
// Solidity: function trustedBurnAll(address _target) returns(uint256)
func (_OwnedERC20 *OwnedERC20Session) TrustedBurnAll(_target common.Address) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TrustedBurnAll(&_OwnedERC20.TransactOpts, _target)
}

// TrustedBurnAll is a paid mutator transaction binding the contract method 0x71297784.
//
// Solidity: function trustedBurnAll(address _target) returns(uint256)
func (_OwnedERC20 *OwnedERC20TransactorSession) TrustedBurnAll(_target common.Address) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TrustedBurnAll(&_OwnedERC20.TransactOpts, _target)
}

// TrustedMint is a paid mutator transaction binding the contract method 0xc024cd26.
//
// Solidity: function trustedMint(address _target, uint256 _amount) returns()
func (_OwnedERC20 *OwnedERC20Transactor) TrustedMint(opts *bind.TransactOpts, _target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.contract.Transact(opts, "trustedMint", _target, _amount)
}

// TrustedMint is a paid mutator transaction binding the contract method 0xc024cd26.
//
// Solidity: function trustedMint(address _target, uint256 _amount) returns()
func (_OwnedERC20 *OwnedERC20Session) TrustedMint(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TrustedMint(&_OwnedERC20.TransactOpts, _target, _amount)
}

// TrustedMint is a paid mutator transaction binding the contract method 0xc024cd26.
//
// Solidity: function trustedMint(address _target, uint256 _amount) returns()
func (_OwnedERC20 *OwnedERC20TransactorSession) TrustedMint(_target common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TrustedMint(&_OwnedERC20.TransactOpts, _target, _amount)
}

// TrustedTransfer is a paid mutator transaction binding the contract method 0x0fb66557.
//
// Solidity: function trustedTransfer(address _from, address _to, uint256 _amount) returns()
func (_OwnedERC20 *OwnedERC20Transactor) TrustedTransfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.contract.Transact(opts, "trustedTransfer", _from, _to, _amount)
}

// TrustedTransfer is a paid mutator transaction binding the contract method 0x0fb66557.
//
// Solidity: function trustedTransfer(address _from, address _to, uint256 _amount) returns()
func (_OwnedERC20 *OwnedERC20Session) TrustedTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TrustedTransfer(&_OwnedERC20.TransactOpts, _from, _to, _amount)
}

// TrustedTransfer is a paid mutator transaction binding the contract method 0x0fb66557.
//
// Solidity: function trustedTransfer(address _from, address _to, uint256 _amount) returns()
func (_OwnedERC20 *OwnedERC20TransactorSession) TrustedTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OwnedERC20.Contract.TrustedTransfer(&_OwnedERC20.TransactOpts, _from, _to, _amount)
}

// OwnedERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the OwnedERC20 contract.
type OwnedERC20ApprovalIterator struct {
	Event *OwnedERC20Approval // Event containing the contract specifics and raw log

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
func (it *OwnedERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnedERC20Approval)
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
		it.Event = new(OwnedERC20Approval)
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
func (it *OwnedERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnedERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnedERC20Approval represents a Approval event raised by the OwnedERC20 contract.
type OwnedERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OwnedERC20 *OwnedERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OwnedERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OwnedERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OwnedERC20ApprovalIterator{contract: _OwnedERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OwnedERC20 *OwnedERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OwnedERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OwnedERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnedERC20Approval)
				if err := _OwnedERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_OwnedERC20 *OwnedERC20Filterer) ParseApproval(log types.Log) (*OwnedERC20Approval, error) {
	event := new(OwnedERC20Approval)
	if err := _OwnedERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnedERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OwnedERC20 contract.
type OwnedERC20TransferIterator struct {
	Event *OwnedERC20Transfer // Event containing the contract specifics and raw log

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
func (it *OwnedERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnedERC20Transfer)
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
		it.Event = new(OwnedERC20Transfer)
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
func (it *OwnedERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnedERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnedERC20Transfer represents a Transfer event raised by the OwnedERC20 contract.
type OwnedERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OwnedERC20 *OwnedERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnedERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnedERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OwnedERC20TransferIterator{contract: _OwnedERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OwnedERC20 *OwnedERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OwnedERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnedERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnedERC20Transfer)
				if err := _OwnedERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_OwnedERC20 *OwnedERC20Filterer) ParseTransfer(log types.Log) (*OwnedERC20Transfer, error) {
	event := new(OwnedERC20Transfer)
	if err := _OwnedERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
