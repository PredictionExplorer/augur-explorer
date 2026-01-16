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

// INonfungiblePositionManagerCollectParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerCollectParams struct {
	TokenId    *big.Int
	Recipient  common.Address
	Amount0Max *big.Int
	Amount1Max *big.Int
}

// INonfungiblePositionManagerDecreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerDecreaseLiquidityParams struct {
	TokenId    *big.Int
	Liquidity  *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
	Deadline   *big.Int
}

// INonfungiblePositionManagerIncreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerIncreaseLiquidityParams struct {
	TokenId        *big.Int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Deadline       *big.Int
}

// INonfungiblePositionManagerMintParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerMintParams struct {
	Token0         common.Address
	Token1         common.Address
	Fee            *big.Int
	TickLower      *big.Int
	TickUpper      *big.Int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Recipient      common.Address
	Deadline       *big.Int
}

// IV3MigratorMigrateParams is an auto generated low-level Go binding around an user-defined struct.
type IV3MigratorMigrateParams struct {
	Pair                common.Address
	LiquidityToMigrate  *big.Int
	PercentageToMigrate uint8
	Token0              common.Address
	Token1              common.Address
	Fee                 *big.Int
	TickLower           *big.Int
	TickUpper           *big.Int
	Amount0Min          *big.Int
	Amount1Min          *big.Int
	Recipient           common.Address
	Deadline            *big.Int
	RefundAsETH         bool
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
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

// IERC721MetaData contains all meta data concerning the IERC721 contract.
var IERC721MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721MetaData.ABI instead.
var IERC721ABI = IERC721MetaData.ABI

// IERC721 is an auto generated Go binding around an Ethereum contract.
type IERC721 struct {
	IERC721Caller     // Read-only binding to the contract
	IERC721Transactor // Write-only binding to the contract
	IERC721Filterer   // Log filterer for contract events
}

// IERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721Session struct {
	Contract     *IERC721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721CallerSession struct {
	Contract *IERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721TransactorSession struct {
	Contract     *IERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721Raw struct {
	Contract *IERC721 // Generic contract binding to access the raw methods on
}

// IERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721CallerRaw struct {
	Contract *IERC721Caller // Generic read-only contract binding to access the raw methods on
}

// IERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721TransactorRaw struct {
	Contract *IERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	contract, err := bindIERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721{IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract}}, nil
}

// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	contract, err := bindIERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Caller{contract: contract}, nil
}

// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	contract, err := bindIERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Transactor{contract: contract}, nil
}

// NewIERC721Filterer creates a new log filterer instance of IERC721, bound to a specific deployed contract.
func NewIERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC721Filterer, error) {
	contract, err := bindIERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721Filterer{contract: contract}, nil
}

// bindIERC721 binds a generic wrapper to an already deployed contract.
func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.IERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Session) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721TransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// IERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721 contract.
type IERC721ApprovalIterator struct {
	Event *IERC721Approval // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Approval)
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
		it.Event = new(IERC721Approval)
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
func (it *IERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Approval)
				if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseApproval(log types.Log) (*IERC721Approval, error) {
	event := new(IERC721Approval)
	if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721 contract.
type IERC721ApprovalForAllIterator struct {
	Event *IERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721ApprovalForAll)
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
		it.Event = new(IERC721ApprovalForAll)
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
func (it *IERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721ApprovalForAll)
				if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) ParseApprovalForAll(log types.Log) (*IERC721ApprovalForAll, error) {
	event := new(IERC721ApprovalForAll)
	if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721 contract.
type IERC721TransferIterator struct {
	Event *IERC721Transfer // Event containing the contract specifics and raw log

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
func (it *IERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Transfer)
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
		it.Event = new(IERC721Transfer)
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
func (it *IERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Transfer)
				if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
	event := new(IERC721Transfer)
	if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721EnumerableMetaData contains all meta data concerning the IERC721Enumerable contract.
var IERC721EnumerableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721EnumerableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721EnumerableMetaData.ABI instead.
var IERC721EnumerableABI = IERC721EnumerableMetaData.ABI

// IERC721Enumerable is an auto generated Go binding around an Ethereum contract.
type IERC721Enumerable struct {
	IERC721EnumerableCaller     // Read-only binding to the contract
	IERC721EnumerableTransactor // Write-only binding to the contract
	IERC721EnumerableFilterer   // Log filterer for contract events
}

// IERC721EnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721EnumerableSession struct {
	Contract     *IERC721Enumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721EnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721EnumerableCallerSession struct {
	Contract *IERC721EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IERC721EnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721EnumerableTransactorSession struct {
	Contract     *IERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IERC721EnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721EnumerableRaw struct {
	Contract *IERC721Enumerable // Generic contract binding to access the raw methods on
}

// IERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721EnumerableCallerRaw struct {
	Contract *IERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721EnumerableTransactorRaw struct {
	Contract *IERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Enumerable creates a new instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721Enumerable(address common.Address, backend bind.ContractBackend) (*IERC721Enumerable, error) {
	contract, err := bindIERC721Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Enumerable{IERC721EnumerableCaller: IERC721EnumerableCaller{contract: contract}, IERC721EnumerableTransactor: IERC721EnumerableTransactor{contract: contract}, IERC721EnumerableFilterer: IERC721EnumerableFilterer{contract: contract}}, nil
}

// NewIERC721EnumerableCaller creates a new read-only instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*IERC721EnumerableCaller, error) {
	contract, err := bindIERC721Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableCaller{contract: contract}, nil
}

// NewIERC721EnumerableTransactor creates a new write-only instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721EnumerableTransactor, error) {
	contract, err := bindIERC721Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableTransactor{contract: contract}, nil
}

// NewIERC721EnumerableFilterer creates a new log filterer instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721EnumerableFilterer, error) {
	contract, err := bindIERC721Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableFilterer{contract: contract}, nil
}

// bindIERC721Enumerable binds a generic wrapper to an already deployed contract.
func bindIERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Enumerable *IERC721EnumerableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Enumerable.Contract.IERC721EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Enumerable *IERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.IERC721EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Enumerable *IERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.IERC721EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Enumerable *IERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Enumerable *IERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Enumerable *IERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Enumerable.Contract.BalanceOf(&_IERC721Enumerable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Enumerable.Contract.BalanceOf(&_IERC721Enumerable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.GetApproved(&_IERC721Enumerable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.GetApproved(&_IERC721Enumerable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Enumerable.Contract.IsApprovedForAll(&_IERC721Enumerable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Enumerable.Contract.IsApprovedForAll(&_IERC721Enumerable.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.OwnerOf(&_IERC721Enumerable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.OwnerOf(&_IERC721Enumerable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Enumerable.Contract.SupportsInterface(&_IERC721Enumerable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Enumerable.Contract.SupportsInterface(&_IERC721Enumerable.CallOpts, interfaceId)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenByIndex(&_IERC721Enumerable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenByIndex(&_IERC721Enumerable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenOfOwnerByIndex(&_IERC721Enumerable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenOfOwnerByIndex(&_IERC721Enumerable.CallOpts, owner, index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableSession) TotalSupply() (*big.Int, error) {
	return _IERC721Enumerable.Contract.TotalSupply(&_IERC721Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC721Enumerable.Contract.TotalSupply(&_IERC721Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.Approve(&_IERC721Enumerable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.Approve(&_IERC721Enumerable.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom0(&_IERC721Enumerable.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom0(&_IERC721Enumerable.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SetApprovalForAll(&_IERC721Enumerable.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SetApprovalForAll(&_IERC721Enumerable.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.TransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.TransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// IERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalIterator struct {
	Event *IERC721EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableApproval)
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
		it.Event = new(IERC721EnumerableApproval)
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
func (it *IERC721EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableApproval represents a Approval event raised by the IERC721Enumerable contract.
type IERC721EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721EnumerableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableApprovalIterator{contract: _IERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableApproval)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseApproval(log types.Log) (*IERC721EnumerableApproval, error) {
	event := new(IERC721EnumerableApproval)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalForAllIterator struct {
	Event *IERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableApprovalForAll)
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
		it.Event = new(IERC721EnumerableApprovalForAll)
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
func (it *IERC721EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721EnumerableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableApprovalForAllIterator{contract: _IERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableApprovalForAll)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseApprovalForAll(log types.Log) (*IERC721EnumerableApprovalForAll, error) {
	event := new(IERC721EnumerableApprovalForAll)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Enumerable contract.
type IERC721EnumerableTransferIterator struct {
	Event *IERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableTransfer)
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
		it.Event = new(IERC721EnumerableTransfer)
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
func (it *IERC721EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableTransfer represents a Transfer event raised by the IERC721Enumerable contract.
type IERC721EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721EnumerableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableTransferIterator{contract: _IERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableTransfer)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseTransfer(log types.Log) (*IERC721EnumerableTransfer, error) {
	event := new(IERC721EnumerableTransfer)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataMetaData contains all meta data concerning the IERC721Metadata contract.
var IERC721MetadataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721MetadataMetaData.ABI instead.
var IERC721MetadataABI = IERC721MetadataMetaData.ABI

// IERC721Metadata is an auto generated Go binding around an Ethereum contract.
type IERC721Metadata struct {
	IERC721MetadataCaller     // Read-only binding to the contract
	IERC721MetadataTransactor // Write-only binding to the contract
	IERC721MetadataFilterer   // Log filterer for contract events
}

// IERC721MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721MetadataSession struct {
	Contract     *IERC721Metadata  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721MetadataCallerSession struct {
	Contract *IERC721MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721MetadataTransactorSession struct {
	Contract     *IERC721MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721MetadataRaw struct {
	Contract *IERC721Metadata // Generic contract binding to access the raw methods on
}

// IERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721MetadataCallerRaw struct {
	Contract *IERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721MetadataTransactorRaw struct {
	Contract *IERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Metadata creates a new instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721Metadata(address common.Address, backend bind.ContractBackend) (*IERC721Metadata, error) {
	contract, err := bindIERC721Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Metadata{IERC721MetadataCaller: IERC721MetadataCaller{contract: contract}, IERC721MetadataTransactor: IERC721MetadataTransactor{contract: contract}, IERC721MetadataFilterer: IERC721MetadataFilterer{contract: contract}}, nil
}

// NewIERC721MetadataCaller creates a new read-only instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC721MetadataCaller, error) {
	contract, err := bindIERC721Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataCaller{contract: contract}, nil
}

// NewIERC721MetadataTransactor creates a new write-only instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721MetadataTransactor, error) {
	contract, err := bindIERC721Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataTransactor{contract: contract}, nil
}

// NewIERC721MetadataFilterer creates a new log filterer instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721MetadataFilterer, error) {
	contract, err := bindIERC721Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataFilterer{contract: contract}, nil
}

// bindIERC721Metadata binds a generic wrapper to an already deployed contract.
func bindIERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Metadata *IERC721MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Metadata.Contract.IERC721MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Metadata *IERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Metadata *IERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Metadata *IERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) Name() (string, error) {
	return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) Name() (string, error) {
	return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) Symbol() (string, error) {
	return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) Symbol() (string, error) {
	return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// IERC721MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Metadata contract.
type IERC721MetadataApprovalIterator struct {
	Event *IERC721MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataApproval)
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
		it.Event = new(IERC721MetadataApproval)
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
func (it *IERC721MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataApproval represents a Approval event raised by the IERC721Metadata contract.
type IERC721MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataApprovalIterator{contract: _IERC721Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataApproval)
				if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) ParseApproval(log types.Log) (*IERC721MetadataApproval, error) {
	event := new(IERC721MetadataApproval)
	if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Metadata contract.
type IERC721MetadataApprovalForAllIterator struct {
	Event *IERC721MetadataApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataApprovalForAll)
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
		it.Event = new(IERC721MetadataApprovalForAll)
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
func (it *IERC721MetadataApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the IERC721Metadata contract.
type IERC721MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721MetadataApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataApprovalForAllIterator{contract: _IERC721Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataApprovalForAll)
				if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) ParseApprovalForAll(log types.Log) (*IERC721MetadataApprovalForAll, error) {
	event := new(IERC721MetadataApprovalForAll)
	if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Metadata contract.
type IERC721MetadataTransferIterator struct {
	Event *IERC721MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataTransfer)
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
		it.Event = new(IERC721MetadataTransfer)
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
func (it *IERC721MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataTransfer represents a Transfer event raised by the IERC721Metadata contract.
type IERC721MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataTransferIterator{contract: _IERC721Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721MetadataTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataTransfer)
				if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) ParseTransfer(log types.Log) (*IERC721MetadataTransfer, error) {
	event := new(IERC721MetadataTransfer)
	if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721PermitMetaData contains all meta data concerning the IERC721Permit contract.
var IERC721PermitMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721PermitABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721PermitMetaData.ABI instead.
var IERC721PermitABI = IERC721PermitMetaData.ABI

// IERC721Permit is an auto generated Go binding around an Ethereum contract.
type IERC721Permit struct {
	IERC721PermitCaller     // Read-only binding to the contract
	IERC721PermitTransactor // Write-only binding to the contract
	IERC721PermitFilterer   // Log filterer for contract events
}

// IERC721PermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721PermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721PermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721PermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721PermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721PermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721PermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721PermitSession struct {
	Contract     *IERC721Permit    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721PermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721PermitCallerSession struct {
	Contract *IERC721PermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IERC721PermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721PermitTransactorSession struct {
	Contract     *IERC721PermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC721PermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721PermitRaw struct {
	Contract *IERC721Permit // Generic contract binding to access the raw methods on
}

// IERC721PermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721PermitCallerRaw struct {
	Contract *IERC721PermitCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721PermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721PermitTransactorRaw struct {
	Contract *IERC721PermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Permit creates a new instance of IERC721Permit, bound to a specific deployed contract.
func NewIERC721Permit(address common.Address, backend bind.ContractBackend) (*IERC721Permit, error) {
	contract, err := bindIERC721Permit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Permit{IERC721PermitCaller: IERC721PermitCaller{contract: contract}, IERC721PermitTransactor: IERC721PermitTransactor{contract: contract}, IERC721PermitFilterer: IERC721PermitFilterer{contract: contract}}, nil
}

// NewIERC721PermitCaller creates a new read-only instance of IERC721Permit, bound to a specific deployed contract.
func NewIERC721PermitCaller(address common.Address, caller bind.ContractCaller) (*IERC721PermitCaller, error) {
	contract, err := bindIERC721Permit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721PermitCaller{contract: contract}, nil
}

// NewIERC721PermitTransactor creates a new write-only instance of IERC721Permit, bound to a specific deployed contract.
func NewIERC721PermitTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721PermitTransactor, error) {
	contract, err := bindIERC721Permit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721PermitTransactor{contract: contract}, nil
}

// NewIERC721PermitFilterer creates a new log filterer instance of IERC721Permit, bound to a specific deployed contract.
func NewIERC721PermitFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721PermitFilterer, error) {
	contract, err := bindIERC721Permit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721PermitFilterer{contract: contract}, nil
}

// bindIERC721Permit binds a generic wrapper to an already deployed contract.
func bindIERC721Permit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721PermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Permit *IERC721PermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Permit.Contract.IERC721PermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Permit *IERC721PermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Permit.Contract.IERC721PermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Permit *IERC721PermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Permit.Contract.IERC721PermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Permit *IERC721PermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Permit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Permit *IERC721PermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Permit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Permit *IERC721PermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Permit.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC721Permit *IERC721PermitCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC721Permit.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC721Permit *IERC721PermitSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC721Permit.Contract.DOMAINSEPARATOR(&_IERC721Permit.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC721Permit *IERC721PermitCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC721Permit.Contract.DOMAINSEPARATOR(&_IERC721Permit.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IERC721Permit *IERC721PermitCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC721Permit.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IERC721Permit *IERC721PermitSession) PERMITTYPEHASH() ([32]byte, error) {
	return _IERC721Permit.Contract.PERMITTYPEHASH(&_IERC721Permit.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_IERC721Permit *IERC721PermitCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _IERC721Permit.Contract.PERMITTYPEHASH(&_IERC721Permit.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Permit *IERC721PermitCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Permit.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Permit *IERC721PermitSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Permit.Contract.BalanceOf(&_IERC721Permit.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Permit *IERC721PermitCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Permit.Contract.BalanceOf(&_IERC721Permit.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Permit *IERC721PermitCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Permit.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Permit *IERC721PermitSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Permit.Contract.GetApproved(&_IERC721Permit.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Permit *IERC721PermitCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Permit.Contract.GetApproved(&_IERC721Permit.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Permit *IERC721PermitCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721Permit.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Permit *IERC721PermitSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Permit.Contract.IsApprovedForAll(&_IERC721Permit.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Permit *IERC721PermitCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Permit.Contract.IsApprovedForAll(&_IERC721Permit.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Permit *IERC721PermitCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Permit.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Permit *IERC721PermitSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Permit.Contract.OwnerOf(&_IERC721Permit.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Permit *IERC721PermitCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Permit.Contract.OwnerOf(&_IERC721Permit.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Permit *IERC721PermitCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721Permit.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Permit *IERC721PermitSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Permit.Contract.SupportsInterface(&_IERC721Permit.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Permit *IERC721PermitCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Permit.Contract.SupportsInterface(&_IERC721Permit.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Permit *IERC721PermitTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Permit.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Permit *IERC721PermitSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Permit.Contract.Approve(&_IERC721Permit.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Permit *IERC721PermitTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Permit.Contract.Approve(&_IERC721Permit.TransactOpts, to, tokenId)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IERC721Permit *IERC721PermitTransactor) Permit(opts *bind.TransactOpts, spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC721Permit.contract.Transact(opts, "permit", spender, tokenId, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IERC721Permit *IERC721PermitSession) Permit(spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC721Permit.Contract.Permit(&_IERC721Permit.TransactOpts, spender, tokenId, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IERC721Permit *IERC721PermitTransactorSession) Permit(spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC721Permit.Contract.Permit(&_IERC721Permit.TransactOpts, spender, tokenId, deadline, v, r, s)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Permit *IERC721PermitTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Permit.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Permit *IERC721PermitSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Permit.Contract.SafeTransferFrom(&_IERC721Permit.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Permit *IERC721PermitTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Permit.Contract.SafeTransferFrom(&_IERC721Permit.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Permit *IERC721PermitTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Permit.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Permit *IERC721PermitSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Permit.Contract.SafeTransferFrom0(&_IERC721Permit.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Permit *IERC721PermitTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Permit.Contract.SafeTransferFrom0(&_IERC721Permit.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Permit *IERC721PermitTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Permit.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Permit *IERC721PermitSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Permit.Contract.SetApprovalForAll(&_IERC721Permit.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Permit *IERC721PermitTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Permit.Contract.SetApprovalForAll(&_IERC721Permit.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Permit *IERC721PermitTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Permit.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Permit *IERC721PermitSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Permit.Contract.TransferFrom(&_IERC721Permit.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Permit *IERC721PermitTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Permit.Contract.TransferFrom(&_IERC721Permit.TransactOpts, from, to, tokenId)
}

// IERC721PermitApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Permit contract.
type IERC721PermitApprovalIterator struct {
	Event *IERC721PermitApproval // Event containing the contract specifics and raw log

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
func (it *IERC721PermitApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721PermitApproval)
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
		it.Event = new(IERC721PermitApproval)
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
func (it *IERC721PermitApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721PermitApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721PermitApproval represents a Approval event raised by the IERC721Permit contract.
type IERC721PermitApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Permit *IERC721PermitFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721PermitApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Permit.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721PermitApprovalIterator{contract: _IERC721Permit.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Permit *IERC721PermitFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721PermitApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Permit.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721PermitApproval)
				if err := _IERC721Permit.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Permit *IERC721PermitFilterer) ParseApproval(log types.Log) (*IERC721PermitApproval, error) {
	event := new(IERC721PermitApproval)
	if err := _IERC721Permit.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721PermitApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Permit contract.
type IERC721PermitApprovalForAllIterator struct {
	Event *IERC721PermitApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721PermitApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721PermitApprovalForAll)
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
		it.Event = new(IERC721PermitApprovalForAll)
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
func (it *IERC721PermitApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721PermitApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721PermitApprovalForAll represents a ApprovalForAll event raised by the IERC721Permit contract.
type IERC721PermitApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Permit *IERC721PermitFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721PermitApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Permit.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721PermitApprovalForAllIterator{contract: _IERC721Permit.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Permit *IERC721PermitFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721PermitApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Permit.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721PermitApprovalForAll)
				if err := _IERC721Permit.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Permit *IERC721PermitFilterer) ParseApprovalForAll(log types.Log) (*IERC721PermitApprovalForAll, error) {
	event := new(IERC721PermitApprovalForAll)
	if err := _IERC721Permit.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721PermitTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Permit contract.
type IERC721PermitTransferIterator struct {
	Event *IERC721PermitTransfer // Event containing the contract specifics and raw log

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
func (it *IERC721PermitTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721PermitTransfer)
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
		it.Event = new(IERC721PermitTransfer)
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
func (it *IERC721PermitTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721PermitTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721PermitTransfer represents a Transfer event raised by the IERC721Permit contract.
type IERC721PermitTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Permit *IERC721PermitFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721PermitTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Permit.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721PermitTransferIterator{contract: _IERC721Permit.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Permit *IERC721PermitFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721PermitTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Permit.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721PermitTransfer)
				if err := _IERC721Permit.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Permit *IERC721PermitFilterer) ParseTransfer(log types.Log) (*IERC721PermitTransfer, error) {
	event := new(IERC721PermitTransfer)
	if err := _IERC721Permit.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// INonfungiblePositionManagerMetaData contains all meta data concerning the INonfungiblePositionManager contract.
var INonfungiblePositionManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"DecreaseLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"IncreaseLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Max\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Max\",\"type\":\"uint128\"}],\"internalType\":\"structINonfungiblePositionManager.CollectParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"createAndInitializePoolIfNecessary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.DecreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"decreaseLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.IncreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"increaseLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.MintParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// INonfungiblePositionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use INonfungiblePositionManagerMetaData.ABI instead.
var INonfungiblePositionManagerABI = INonfungiblePositionManagerMetaData.ABI

// INonfungiblePositionManager is an auto generated Go binding around an Ethereum contract.
type INonfungiblePositionManager struct {
	INonfungiblePositionManagerCaller     // Read-only binding to the contract
	INonfungiblePositionManagerTransactor // Write-only binding to the contract
	INonfungiblePositionManagerFilterer   // Log filterer for contract events
}

// INonfungiblePositionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type INonfungiblePositionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonfungiblePositionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INonfungiblePositionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonfungiblePositionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INonfungiblePositionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonfungiblePositionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INonfungiblePositionManagerSession struct {
	Contract     *INonfungiblePositionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// INonfungiblePositionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INonfungiblePositionManagerCallerSession struct {
	Contract *INonfungiblePositionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// INonfungiblePositionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INonfungiblePositionManagerTransactorSession struct {
	Contract     *INonfungiblePositionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// INonfungiblePositionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type INonfungiblePositionManagerRaw struct {
	Contract *INonfungiblePositionManager // Generic contract binding to access the raw methods on
}

// INonfungiblePositionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INonfungiblePositionManagerCallerRaw struct {
	Contract *INonfungiblePositionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// INonfungiblePositionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INonfungiblePositionManagerTransactorRaw struct {
	Contract *INonfungiblePositionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINonfungiblePositionManager creates a new instance of INonfungiblePositionManager, bound to a specific deployed contract.
func NewINonfungiblePositionManager(address common.Address, backend bind.ContractBackend) (*INonfungiblePositionManager, error) {
	contract, err := bindINonfungiblePositionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INonfungiblePositionManager{INonfungiblePositionManagerCaller: INonfungiblePositionManagerCaller{contract: contract}, INonfungiblePositionManagerTransactor: INonfungiblePositionManagerTransactor{contract: contract}, INonfungiblePositionManagerFilterer: INonfungiblePositionManagerFilterer{contract: contract}}, nil
}

// NewINonfungiblePositionManagerCaller creates a new read-only instance of INonfungiblePositionManager, bound to a specific deployed contract.
func NewINonfungiblePositionManagerCaller(address common.Address, caller bind.ContractCaller) (*INonfungiblePositionManagerCaller, error) {
	contract, err := bindINonfungiblePositionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INonfungiblePositionManagerCaller{contract: contract}, nil
}

// NewINonfungiblePositionManagerTransactor creates a new write-only instance of INonfungiblePositionManager, bound to a specific deployed contract.
func NewINonfungiblePositionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*INonfungiblePositionManagerTransactor, error) {
	contract, err := bindINonfungiblePositionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INonfungiblePositionManagerTransactor{contract: contract}, nil
}

// NewINonfungiblePositionManagerFilterer creates a new log filterer instance of INonfungiblePositionManager, bound to a specific deployed contract.
func NewINonfungiblePositionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*INonfungiblePositionManagerFilterer, error) {
	contract, err := bindINonfungiblePositionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INonfungiblePositionManagerFilterer{contract: contract}, nil
}

// bindINonfungiblePositionManager binds a generic wrapper to an already deployed contract.
func bindINonfungiblePositionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INonfungiblePositionManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INonfungiblePositionManager *INonfungiblePositionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INonfungiblePositionManager.Contract.INonfungiblePositionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INonfungiblePositionManager *INonfungiblePositionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.INonfungiblePositionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INonfungiblePositionManager *INonfungiblePositionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.INonfungiblePositionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INonfungiblePositionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _INonfungiblePositionManager.Contract.DOMAINSEPARATOR(&_INonfungiblePositionManager.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _INonfungiblePositionManager.Contract.DOMAINSEPARATOR(&_INonfungiblePositionManager.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _INonfungiblePositionManager.Contract.PERMITTYPEHASH(&_INonfungiblePositionManager.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _INonfungiblePositionManager.Contract.PERMITTYPEHASH(&_INonfungiblePositionManager.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) WETH9() (common.Address, error) {
	return _INonfungiblePositionManager.Contract.WETH9(&_INonfungiblePositionManager.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) WETH9() (common.Address, error) {
	return _INonfungiblePositionManager.Contract.WETH9(&_INonfungiblePositionManager.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _INonfungiblePositionManager.Contract.BalanceOf(&_INonfungiblePositionManager.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _INonfungiblePositionManager.Contract.BalanceOf(&_INonfungiblePositionManager.CallOpts, owner)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) Factory() (common.Address, error) {
	return _INonfungiblePositionManager.Contract.Factory(&_INonfungiblePositionManager.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) Factory() (common.Address, error) {
	return _INonfungiblePositionManager.Contract.Factory(&_INonfungiblePositionManager.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _INonfungiblePositionManager.Contract.GetApproved(&_INonfungiblePositionManager.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _INonfungiblePositionManager.Contract.GetApproved(&_INonfungiblePositionManager.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _INonfungiblePositionManager.Contract.IsApprovedForAll(&_INonfungiblePositionManager.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _INonfungiblePositionManager.Contract.IsApprovedForAll(&_INonfungiblePositionManager.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) Name() (string, error) {
	return _INonfungiblePositionManager.Contract.Name(&_INonfungiblePositionManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) Name() (string, error) {
	return _INonfungiblePositionManager.Contract.Name(&_INonfungiblePositionManager.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _INonfungiblePositionManager.Contract.OwnerOf(&_INonfungiblePositionManager.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _INonfungiblePositionManager.Contract.OwnerOf(&_INonfungiblePositionManager.CallOpts, tokenId)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) Positions(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "positions", tokenId)

	outstruct := new(struct {
		Nonce                    *big.Int
		Operator                 common.Address
		Token0                   common.Address
		Token1                   common.Address
		Fee                      *big.Int
		TickLower                *big.Int
		TickUpper                *big.Int
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TickLower = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TickUpper = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) Positions(tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _INonfungiblePositionManager.Contract.Positions(&_INonfungiblePositionManager.CallOpts, tokenId)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 tokenId) view returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) Positions(tokenId *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _INonfungiblePositionManager.Contract.Positions(&_INonfungiblePositionManager.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _INonfungiblePositionManager.Contract.SupportsInterface(&_INonfungiblePositionManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _INonfungiblePositionManager.Contract.SupportsInterface(&_INonfungiblePositionManager.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) Symbol() (string, error) {
	return _INonfungiblePositionManager.Contract.Symbol(&_INonfungiblePositionManager.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) Symbol() (string, error) {
	return _INonfungiblePositionManager.Contract.Symbol(&_INonfungiblePositionManager.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _INonfungiblePositionManager.Contract.TokenByIndex(&_INonfungiblePositionManager.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _INonfungiblePositionManager.Contract.TokenByIndex(&_INonfungiblePositionManager.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _INonfungiblePositionManager.Contract.TokenOfOwnerByIndex(&_INonfungiblePositionManager.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _INonfungiblePositionManager.Contract.TokenOfOwnerByIndex(&_INonfungiblePositionManager.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _INonfungiblePositionManager.Contract.TokenURI(&_INonfungiblePositionManager.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _INonfungiblePositionManager.Contract.TokenURI(&_INonfungiblePositionManager.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INonfungiblePositionManager.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) TotalSupply() (*big.Int, error) {
	return _INonfungiblePositionManager.Contract.TotalSupply(&_INonfungiblePositionManager.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_INonfungiblePositionManager *INonfungiblePositionManagerCallerSession) TotalSupply() (*big.Int, error) {
	return _INonfungiblePositionManager.Contract.TotalSupply(&_INonfungiblePositionManager.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.Approve(&_INonfungiblePositionManager.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.Approve(&_INonfungiblePositionManager.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.Burn(&_INonfungiblePositionManager.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.Burn(&_INonfungiblePositionManager.TransactOpts, tokenId)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) Collect(opts *bind.TransactOpts, params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "collect", params)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) Collect(params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.Collect(&_INonfungiblePositionManager.TransactOpts, params)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) payable returns(uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) Collect(params INonfungiblePositionManagerCollectParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.Collect(&_INonfungiblePositionManager.TransactOpts, params)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) CreateAndInitializePoolIfNecessary(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "createAndInitializePoolIfNecessary", token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.CreateAndInitializePoolIfNecessary(&_INonfungiblePositionManager.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.CreateAndInitializePoolIfNecessary(&_INonfungiblePositionManager.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) DecreaseLiquidity(opts *bind.TransactOpts, params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "decreaseLiquidity", params)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) DecreaseLiquidity(params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.DecreaseLiquidity(&_INonfungiblePositionManager.TransactOpts, params)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x0c49ccbe.
//
// Solidity: function decreaseLiquidity((uint256,uint128,uint256,uint256,uint256) params) payable returns(uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) DecreaseLiquidity(params INonfungiblePositionManagerDecreaseLiquidityParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.DecreaseLiquidity(&_INonfungiblePositionManager.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) IncreaseLiquidity(opts *bind.TransactOpts, params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "increaseLiquidity", params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) IncreaseLiquidity(params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.IncreaseLiquidity(&_INonfungiblePositionManager.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x219f5d17.
//
// Solidity: function increaseLiquidity((uint256,uint256,uint256,uint256,uint256,uint256) params) payable returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) IncreaseLiquidity(params INonfungiblePositionManagerIncreaseLiquidityParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.IncreaseLiquidity(&_INonfungiblePositionManager.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) Mint(opts *bind.TransactOpts, params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "mint", params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.Mint(&_INonfungiblePositionManager.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.Mint(&_INonfungiblePositionManager.TransactOpts, params)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) Permit(opts *bind.TransactOpts, spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "permit", spender, tokenId, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) Permit(spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.Permit(&_INonfungiblePositionManager.TransactOpts, spender, tokenId, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x7ac2ff7b.
//
// Solidity: function permit(address spender, uint256 tokenId, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) Permit(spender common.Address, tokenId *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.Permit(&_INonfungiblePositionManager.TransactOpts, spender, tokenId, deadline, v, r, s)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) RefundETH() (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.RefundETH(&_INonfungiblePositionManager.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) RefundETH() (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.RefundETH(&_INonfungiblePositionManager.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.SafeTransferFrom(&_INonfungiblePositionManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.SafeTransferFrom(&_INonfungiblePositionManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.SafeTransferFrom0(&_INonfungiblePositionManager.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.SafeTransferFrom0(&_INonfungiblePositionManager.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.SetApprovalForAll(&_INonfungiblePositionManager.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.SetApprovalForAll(&_INonfungiblePositionManager.TransactOpts, operator, _approved)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.SweepToken(&_INonfungiblePositionManager.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.SweepToken(&_INonfungiblePositionManager.TransactOpts, token, amountMinimum, recipient)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.TransferFrom(&_INonfungiblePositionManager.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.TransferFrom(&_INonfungiblePositionManager.TransactOpts, from, to, tokenId)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _INonfungiblePositionManager.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.UnwrapWETH9(&_INonfungiblePositionManager.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_INonfungiblePositionManager *INonfungiblePositionManagerTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _INonfungiblePositionManager.Contract.UnwrapWETH9(&_INonfungiblePositionManager.TransactOpts, amountMinimum, recipient)
}

// INonfungiblePositionManagerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerApprovalIterator struct {
	Event *INonfungiblePositionManagerApproval // Event containing the contract specifics and raw log

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
func (it *INonfungiblePositionManagerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INonfungiblePositionManagerApproval)
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
		it.Event = new(INonfungiblePositionManagerApproval)
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
func (it *INonfungiblePositionManagerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INonfungiblePositionManagerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INonfungiblePositionManagerApproval represents a Approval event raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*INonfungiblePositionManagerApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &INonfungiblePositionManagerApprovalIterator{contract: _INonfungiblePositionManager.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *INonfungiblePositionManagerApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INonfungiblePositionManagerApproval)
				if err := _INonfungiblePositionManager.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) ParseApproval(log types.Log) (*INonfungiblePositionManagerApproval, error) {
	event := new(INonfungiblePositionManagerApproval)
	if err := _INonfungiblePositionManager.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INonfungiblePositionManagerApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerApprovalForAllIterator struct {
	Event *INonfungiblePositionManagerApprovalForAll // Event containing the contract specifics and raw log

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
func (it *INonfungiblePositionManagerApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INonfungiblePositionManagerApprovalForAll)
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
		it.Event = new(INonfungiblePositionManagerApprovalForAll)
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
func (it *INonfungiblePositionManagerApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INonfungiblePositionManagerApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INonfungiblePositionManagerApprovalForAll represents a ApprovalForAll event raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*INonfungiblePositionManagerApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &INonfungiblePositionManagerApprovalForAllIterator{contract: _INonfungiblePositionManager.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *INonfungiblePositionManagerApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INonfungiblePositionManagerApprovalForAll)
				if err := _INonfungiblePositionManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) ParseApprovalForAll(log types.Log) (*INonfungiblePositionManagerApprovalForAll, error) {
	event := new(INonfungiblePositionManagerApprovalForAll)
	if err := _INonfungiblePositionManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INonfungiblePositionManagerCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerCollectIterator struct {
	Event *INonfungiblePositionManagerCollect // Event containing the contract specifics and raw log

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
func (it *INonfungiblePositionManagerCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INonfungiblePositionManagerCollect)
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
		it.Event = new(INonfungiblePositionManagerCollect)
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
func (it *INonfungiblePositionManagerCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INonfungiblePositionManagerCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INonfungiblePositionManagerCollect represents a Collect event raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerCollect struct {
	TokenId   *big.Int
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) FilterCollect(opts *bind.FilterOpts, tokenId []*big.Int) (*INonfungiblePositionManagerCollectIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.FilterLogs(opts, "Collect", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &INonfungiblePositionManagerCollectIterator{contract: _INonfungiblePositionManager.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *INonfungiblePositionManagerCollect, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.WatchLogs(opts, "Collect", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INonfungiblePositionManagerCollect)
				if err := _INonfungiblePositionManager.contract.UnpackLog(event, "Collect", log); err != nil {
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

// ParseCollect is a log parse operation binding the contract event 0x40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01.
//
// Solidity: event Collect(uint256 indexed tokenId, address recipient, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) ParseCollect(log types.Log) (*INonfungiblePositionManagerCollect, error) {
	event := new(INonfungiblePositionManagerCollect)
	if err := _INonfungiblePositionManager.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INonfungiblePositionManagerDecreaseLiquidityIterator is returned from FilterDecreaseLiquidity and is used to iterate over the raw logs and unpacked data for DecreaseLiquidity events raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerDecreaseLiquidityIterator struct {
	Event *INonfungiblePositionManagerDecreaseLiquidity // Event containing the contract specifics and raw log

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
func (it *INonfungiblePositionManagerDecreaseLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INonfungiblePositionManagerDecreaseLiquidity)
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
		it.Event = new(INonfungiblePositionManagerDecreaseLiquidity)
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
func (it *INonfungiblePositionManagerDecreaseLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INonfungiblePositionManagerDecreaseLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INonfungiblePositionManagerDecreaseLiquidity represents a DecreaseLiquidity event raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerDecreaseLiquidity struct {
	TokenId   *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDecreaseLiquidity is a free log retrieval operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) FilterDecreaseLiquidity(opts *bind.FilterOpts, tokenId []*big.Int) (*INonfungiblePositionManagerDecreaseLiquidityIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.FilterLogs(opts, "DecreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &INonfungiblePositionManagerDecreaseLiquidityIterator{contract: _INonfungiblePositionManager.contract, event: "DecreaseLiquidity", logs: logs, sub: sub}, nil
}

// WatchDecreaseLiquidity is a free log subscription operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) WatchDecreaseLiquidity(opts *bind.WatchOpts, sink chan<- *INonfungiblePositionManagerDecreaseLiquidity, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.WatchLogs(opts, "DecreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INonfungiblePositionManagerDecreaseLiquidity)
				if err := _INonfungiblePositionManager.contract.UnpackLog(event, "DecreaseLiquidity", log); err != nil {
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

// ParseDecreaseLiquidity is a log parse operation binding the contract event 0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4.
//
// Solidity: event DecreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) ParseDecreaseLiquidity(log types.Log) (*INonfungiblePositionManagerDecreaseLiquidity, error) {
	event := new(INonfungiblePositionManagerDecreaseLiquidity)
	if err := _INonfungiblePositionManager.contract.UnpackLog(event, "DecreaseLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INonfungiblePositionManagerIncreaseLiquidityIterator is returned from FilterIncreaseLiquidity and is used to iterate over the raw logs and unpacked data for IncreaseLiquidity events raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerIncreaseLiquidityIterator struct {
	Event *INonfungiblePositionManagerIncreaseLiquidity // Event containing the contract specifics and raw log

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
func (it *INonfungiblePositionManagerIncreaseLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INonfungiblePositionManagerIncreaseLiquidity)
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
		it.Event = new(INonfungiblePositionManagerIncreaseLiquidity)
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
func (it *INonfungiblePositionManagerIncreaseLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INonfungiblePositionManagerIncreaseLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INonfungiblePositionManagerIncreaseLiquidity represents a IncreaseLiquidity event raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerIncreaseLiquidity struct {
	TokenId   *big.Int
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreaseLiquidity is a free log retrieval operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) FilterIncreaseLiquidity(opts *bind.FilterOpts, tokenId []*big.Int) (*INonfungiblePositionManagerIncreaseLiquidityIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.FilterLogs(opts, "IncreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &INonfungiblePositionManagerIncreaseLiquidityIterator{contract: _INonfungiblePositionManager.contract, event: "IncreaseLiquidity", logs: logs, sub: sub}, nil
}

// WatchIncreaseLiquidity is a free log subscription operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) WatchIncreaseLiquidity(opts *bind.WatchOpts, sink chan<- *INonfungiblePositionManagerIncreaseLiquidity, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.WatchLogs(opts, "IncreaseLiquidity", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INonfungiblePositionManagerIncreaseLiquidity)
				if err := _INonfungiblePositionManager.contract.UnpackLog(event, "IncreaseLiquidity", log); err != nil {
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

// ParseIncreaseLiquidity is a log parse operation binding the contract event 0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f.
//
// Solidity: event IncreaseLiquidity(uint256 indexed tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) ParseIncreaseLiquidity(log types.Log) (*INonfungiblePositionManagerIncreaseLiquidity, error) {
	event := new(INonfungiblePositionManagerIncreaseLiquidity)
	if err := _INonfungiblePositionManager.contract.UnpackLog(event, "IncreaseLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INonfungiblePositionManagerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerTransferIterator struct {
	Event *INonfungiblePositionManagerTransfer // Event containing the contract specifics and raw log

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
func (it *INonfungiblePositionManagerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INonfungiblePositionManagerTransfer)
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
		it.Event = new(INonfungiblePositionManagerTransfer)
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
func (it *INonfungiblePositionManagerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INonfungiblePositionManagerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INonfungiblePositionManagerTransfer represents a Transfer event raised by the INonfungiblePositionManager contract.
type INonfungiblePositionManagerTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*INonfungiblePositionManagerTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &INonfungiblePositionManagerTransferIterator{contract: _INonfungiblePositionManager.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *INonfungiblePositionManagerTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _INonfungiblePositionManager.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INonfungiblePositionManagerTransfer)
				if err := _INonfungiblePositionManager.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_INonfungiblePositionManager *INonfungiblePositionManagerFilterer) ParseTransfer(log types.Log) (*INonfungiblePositionManagerTransfer, error) {
	event := new(INonfungiblePositionManagerTransfer)
	if err := _INonfungiblePositionManager.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// IPoolInitializerMetaData contains all meta data concerning the IPoolInitializer contract.
var IPoolInitializerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"createAndInitializePoolIfNecessary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IPoolInitializerABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoolInitializerMetaData.ABI instead.
var IPoolInitializerABI = IPoolInitializerMetaData.ABI

// IPoolInitializer is an auto generated Go binding around an Ethereum contract.
type IPoolInitializer struct {
	IPoolInitializerCaller     // Read-only binding to the contract
	IPoolInitializerTransactor // Write-only binding to the contract
	IPoolInitializerFilterer   // Log filterer for contract events
}

// IPoolInitializerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolInitializerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolInitializerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolInitializerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolInitializerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolInitializerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolInitializerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolInitializerSession struct {
	Contract     *IPoolInitializer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolInitializerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolInitializerCallerSession struct {
	Contract *IPoolInitializerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IPoolInitializerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolInitializerTransactorSession struct {
	Contract     *IPoolInitializerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IPoolInitializerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolInitializerRaw struct {
	Contract *IPoolInitializer // Generic contract binding to access the raw methods on
}

// IPoolInitializerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolInitializerCallerRaw struct {
	Contract *IPoolInitializerCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolInitializerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolInitializerTransactorRaw struct {
	Contract *IPoolInitializerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPoolInitializer creates a new instance of IPoolInitializer, bound to a specific deployed contract.
func NewIPoolInitializer(address common.Address, backend bind.ContractBackend) (*IPoolInitializer, error) {
	contract, err := bindIPoolInitializer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPoolInitializer{IPoolInitializerCaller: IPoolInitializerCaller{contract: contract}, IPoolInitializerTransactor: IPoolInitializerTransactor{contract: contract}, IPoolInitializerFilterer: IPoolInitializerFilterer{contract: contract}}, nil
}

// NewIPoolInitializerCaller creates a new read-only instance of IPoolInitializer, bound to a specific deployed contract.
func NewIPoolInitializerCaller(address common.Address, caller bind.ContractCaller) (*IPoolInitializerCaller, error) {
	contract, err := bindIPoolInitializer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolInitializerCaller{contract: contract}, nil
}

// NewIPoolInitializerTransactor creates a new write-only instance of IPoolInitializer, bound to a specific deployed contract.
func NewIPoolInitializerTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolInitializerTransactor, error) {
	contract, err := bindIPoolInitializer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolInitializerTransactor{contract: contract}, nil
}

// NewIPoolInitializerFilterer creates a new log filterer instance of IPoolInitializer, bound to a specific deployed contract.
func NewIPoolInitializerFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolInitializerFilterer, error) {
	contract, err := bindIPoolInitializer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolInitializerFilterer{contract: contract}, nil
}

// bindIPoolInitializer binds a generic wrapper to an already deployed contract.
func bindIPoolInitializer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPoolInitializerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolInitializer *IPoolInitializerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolInitializer.Contract.IPoolInitializerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolInitializer *IPoolInitializerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.IPoolInitializerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolInitializer *IPoolInitializerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.IPoolInitializerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolInitializer *IPoolInitializerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolInitializer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolInitializer *IPoolInitializerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolInitializer *IPoolInitializerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.contract.Transact(opts, method, params...)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_IPoolInitializer *IPoolInitializerTransactor) CreateAndInitializePoolIfNecessary(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IPoolInitializer.contract.Transact(opts, "createAndInitializePoolIfNecessary", token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_IPoolInitializer *IPoolInitializerSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.CreateAndInitializePoolIfNecessary(&_IPoolInitializer.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_IPoolInitializer *IPoolInitializerTransactorSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.CreateAndInitializePoolIfNecessary(&_IPoolInitializer.TransactOpts, token0, token1, fee, sqrtPriceX96)
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

// IUniswapV3PoolMetaData contains all meta data concerning the IUniswapV3Pool contract.
var IUniswapV3PoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"CollectProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextOld\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextNew\",\"type\":\"uint16\"}],\"name\":\"IncreaseObservationCardinalityNext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0New\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1New\",\"type\":\"uint8\"}],\"name\":\"SetFeeProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collectProtocol\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal0X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal1X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"}],\"name\":\"increaseObservationCardinalityNext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidityPerTick\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"observations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityCumulativeX128\",\"type\":\"uint160\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"observe\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"tickCumulatives\",\"type\":\"int56[]\"},{\"internalType\":\"uint160[]\",\"name\":\"secondsPerLiquidityCumulativeX128s\",\"type\":\"uint160[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"token0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"token1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"feeProtocol0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol1\",\"type\":\"uint8\"}],\"name\":\"setFeeProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slot0\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"snapshotCumulativesInside\",\"outputs\":[{\"internalType\":\"int56\",\"name\":\"tickCumulativeInside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityInsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsInside\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountSpecified\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"wordPosition\",\"type\":\"int16\"}],\"name\":\"tickBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside0X128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside1X128\",\"type\":\"uint256\"},{\"internalType\":\"int56\",\"name\":\"tickCumulativeOutside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityOutsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsOutside\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"CollectProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextOld\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextNew\",\"type\":\"uint16\"}],\"name\":\"IncreaseObservationCardinalityNext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0New\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1New\",\"type\":\"uint8\"}],\"name\":\"SetFeeProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"}]",
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

// IV3MigratorMetaData contains all meta data concerning the IV3Migrator contract.
var IV3MigratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"createAndInitializePoolIfNecessary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidityToMigrate\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"percentageToMigrate\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"refundAsETH\",\"type\":\"bool\"}],\"internalType\":\"structIV3Migrator.MigrateParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IV3MigratorABI is the input ABI used to generate the binding from.
// Deprecated: Use IV3MigratorMetaData.ABI instead.
var IV3MigratorABI = IV3MigratorMetaData.ABI

// IV3Migrator is an auto generated Go binding around an Ethereum contract.
type IV3Migrator struct {
	IV3MigratorCaller     // Read-only binding to the contract
	IV3MigratorTransactor // Write-only binding to the contract
	IV3MigratorFilterer   // Log filterer for contract events
}

// IV3MigratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IV3MigratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IV3MigratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IV3MigratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IV3MigratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IV3MigratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IV3MigratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IV3MigratorSession struct {
	Contract     *IV3Migrator      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IV3MigratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IV3MigratorCallerSession struct {
	Contract *IV3MigratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IV3MigratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IV3MigratorTransactorSession struct {
	Contract     *IV3MigratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IV3MigratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IV3MigratorRaw struct {
	Contract *IV3Migrator // Generic contract binding to access the raw methods on
}

// IV3MigratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IV3MigratorCallerRaw struct {
	Contract *IV3MigratorCaller // Generic read-only contract binding to access the raw methods on
}

// IV3MigratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IV3MigratorTransactorRaw struct {
	Contract *IV3MigratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIV3Migrator creates a new instance of IV3Migrator, bound to a specific deployed contract.
func NewIV3Migrator(address common.Address, backend bind.ContractBackend) (*IV3Migrator, error) {
	contract, err := bindIV3Migrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IV3Migrator{IV3MigratorCaller: IV3MigratorCaller{contract: contract}, IV3MigratorTransactor: IV3MigratorTransactor{contract: contract}, IV3MigratorFilterer: IV3MigratorFilterer{contract: contract}}, nil
}

// NewIV3MigratorCaller creates a new read-only instance of IV3Migrator, bound to a specific deployed contract.
func NewIV3MigratorCaller(address common.Address, caller bind.ContractCaller) (*IV3MigratorCaller, error) {
	contract, err := bindIV3Migrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IV3MigratorCaller{contract: contract}, nil
}

// NewIV3MigratorTransactor creates a new write-only instance of IV3Migrator, bound to a specific deployed contract.
func NewIV3MigratorTransactor(address common.Address, transactor bind.ContractTransactor) (*IV3MigratorTransactor, error) {
	contract, err := bindIV3Migrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IV3MigratorTransactor{contract: contract}, nil
}

// NewIV3MigratorFilterer creates a new log filterer instance of IV3Migrator, bound to a specific deployed contract.
func NewIV3MigratorFilterer(address common.Address, filterer bind.ContractFilterer) (*IV3MigratorFilterer, error) {
	contract, err := bindIV3Migrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IV3MigratorFilterer{contract: contract}, nil
}

// bindIV3Migrator binds a generic wrapper to an already deployed contract.
func bindIV3Migrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IV3MigratorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IV3Migrator *IV3MigratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IV3Migrator.Contract.IV3MigratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IV3Migrator *IV3MigratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IV3Migrator.Contract.IV3MigratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IV3Migrator *IV3MigratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IV3Migrator.Contract.IV3MigratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IV3Migrator *IV3MigratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IV3Migrator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IV3Migrator *IV3MigratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IV3Migrator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IV3Migrator *IV3MigratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IV3Migrator.Contract.contract.Transact(opts, method, params...)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_IV3Migrator *IV3MigratorTransactor) CreateAndInitializePoolIfNecessary(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IV3Migrator.contract.Transact(opts, "createAndInitializePoolIfNecessary", token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_IV3Migrator *IV3MigratorSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IV3Migrator.Contract.CreateAndInitializePoolIfNecessary(&_IV3Migrator.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_IV3Migrator *IV3MigratorTransactorSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IV3Migrator.Contract.CreateAndInitializePoolIfNecessary(&_IV3Migrator.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// Migrate is a paid mutator transaction binding the contract method 0xd44f2bf2.
//
// Solidity: function migrate((address,uint256,uint8,address,address,uint24,int24,int24,uint256,uint256,address,uint256,bool) params) returns()
func (_IV3Migrator *IV3MigratorTransactor) Migrate(opts *bind.TransactOpts, params IV3MigratorMigrateParams) (*types.Transaction, error) {
	return _IV3Migrator.contract.Transact(opts, "migrate", params)
}

// Migrate is a paid mutator transaction binding the contract method 0xd44f2bf2.
//
// Solidity: function migrate((address,uint256,uint8,address,address,uint24,int24,int24,uint256,uint256,address,uint256,bool) params) returns()
func (_IV3Migrator *IV3MigratorSession) Migrate(params IV3MigratorMigrateParams) (*types.Transaction, error) {
	return _IV3Migrator.Contract.Migrate(&_IV3Migrator.TransactOpts, params)
}

// Migrate is a paid mutator transaction binding the contract method 0xd44f2bf2.
//
// Solidity: function migrate((address,uint256,uint8,address,address,uint24,int24,int24,uint256,uint256,address,uint256,bool) params) returns()
func (_IV3Migrator *IV3MigratorTransactorSession) Migrate(params IV3MigratorMigrateParams) (*types.Transaction, error) {
	return _IV3Migrator.Contract.Migrate(&_IV3Migrator.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IV3Migrator *IV3MigratorTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _IV3Migrator.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IV3Migrator *IV3MigratorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _IV3Migrator.Contract.Multicall(&_IV3Migrator.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IV3Migrator *IV3MigratorTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _IV3Migrator.Contract.Multicall(&_IV3Migrator.TransactOpts, data)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.Contract.SelfPermit(&_IV3Migrator.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.Contract.SelfPermit(&_IV3Migrator.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.Contract.SelfPermitAllowed(&_IV3Migrator.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.Contract.SelfPermitAllowed(&_IV3Migrator.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.Contract.SelfPermitAllowedIfNecessary(&_IV3Migrator.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.Contract.SelfPermitAllowedIfNecessary(&_IV3Migrator.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.Contract.SelfPermitIfNecessary(&_IV3Migrator.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_IV3Migrator *IV3MigratorTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IV3Migrator.Contract.SelfPermitIfNecessary(&_IV3Migrator.TransactOpts, token, value, deadline, v, r, s)
}

// IWETH9MetaData contains all meta data concerning the IWETH9 contract.
var IWETH9MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IWETH9ABI is the input ABI used to generate the binding from.
// Deprecated: Use IWETH9MetaData.ABI instead.
var IWETH9ABI = IWETH9MetaData.ABI

// IWETH9 is an auto generated Go binding around an Ethereum contract.
type IWETH9 struct {
	IWETH9Caller     // Read-only binding to the contract
	IWETH9Transactor // Write-only binding to the contract
	IWETH9Filterer   // Log filterer for contract events
}

// IWETH9Caller is an auto generated read-only Go binding around an Ethereum contract.
type IWETH9Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETH9Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETH9Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETH9Session struct {
	Contract     *IWETH9           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETH9CallerSession struct {
	Contract *IWETH9Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWETH9TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETH9TransactorSession struct {
	Contract     *IWETH9Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9Raw is an auto generated low-level Go binding around an Ethereum contract.
type IWETH9Raw struct {
	Contract *IWETH9 // Generic contract binding to access the raw methods on
}

// IWETH9CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETH9CallerRaw struct {
	Contract *IWETH9Caller // Generic read-only contract binding to access the raw methods on
}

// IWETH9TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETH9TransactorRaw struct {
	Contract *IWETH9Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH9 creates a new instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9(address common.Address, backend bind.ContractBackend) (*IWETH9, error) {
	contract, err := bindIWETH9(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH9{IWETH9Caller: IWETH9Caller{contract: contract}, IWETH9Transactor: IWETH9Transactor{contract: contract}, IWETH9Filterer: IWETH9Filterer{contract: contract}}, nil
}

// NewIWETH9Caller creates a new read-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Caller(address common.Address, caller bind.ContractCaller) (*IWETH9Caller, error) {
	contract, err := bindIWETH9(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Caller{contract: contract}, nil
}

// NewIWETH9Transactor creates a new write-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Transactor(address common.Address, transactor bind.ContractTransactor) (*IWETH9Transactor, error) {
	contract, err := bindIWETH9(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Transactor{contract: contract}, nil
}

// NewIWETH9Filterer creates a new log filterer instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Filterer(address common.Address, filterer bind.ContractFilterer) (*IWETH9Filterer, error) {
	contract, err := bindIWETH9(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETH9Filterer{contract: contract}, nil
}

// bindIWETH9 binds a generic wrapper to an already deployed contract.
func bindIWETH9(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWETH9ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.IWETH9Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH9 *IWETH9Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH9 *IWETH9Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH9.Contract.Allowance(&_IWETH9.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH9.Contract.Allowance(&_IWETH9.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH9 *IWETH9Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH9 *IWETH9Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWETH9.Contract.BalanceOf(&_IWETH9.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWETH9.Contract.BalanceOf(&_IWETH9.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9Session) TotalSupply() (*big.Int, error) {
	return _IWETH9.Contract.TotalSupply(&_IWETH9.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) TotalSupply() (*big.Int, error) {
	return _IWETH9.Contract.TotalSupply(&_IWETH9.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Approve(&_IWETH9.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Approve(&_IWETH9.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Session) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9TransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Transfer(&_IWETH9.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Transfer(&_IWETH9.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.TransferFrom(&_IWETH9.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.TransferFrom(&_IWETH9.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH9 *IWETH9Transactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH9 *IWETH9Session) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH9 *IWETH9TransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, arg0)
}

// IWETH9ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IWETH9 contract.
type IWETH9ApprovalIterator struct {
	Event *IWETH9Approval // Event containing the contract specifics and raw log

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
func (it *IWETH9ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETH9Approval)
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
		it.Event = new(IWETH9Approval)
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
func (it *IWETH9ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETH9ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETH9Approval represents a Approval event raised by the IWETH9 contract.
type IWETH9Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH9 *IWETH9Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IWETH9ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH9.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IWETH9ApprovalIterator{contract: _IWETH9.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH9 *IWETH9Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IWETH9Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH9.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETH9Approval)
				if err := _IWETH9.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IWETH9 *IWETH9Filterer) ParseApproval(log types.Log) (*IWETH9Approval, error) {
	event := new(IWETH9Approval)
	if err := _IWETH9.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWETH9TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IWETH9 contract.
type IWETH9TransferIterator struct {
	Event *IWETH9Transfer // Event containing the contract specifics and raw log

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
func (it *IWETH9TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETH9Transfer)
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
		it.Event = new(IWETH9Transfer)
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
func (it *IWETH9TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETH9TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETH9Transfer represents a Transfer event raised by the IWETH9 contract.
type IWETH9Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH9 *IWETH9Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IWETH9TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH9.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IWETH9TransferIterator{contract: _IWETH9.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH9 *IWETH9Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IWETH9Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH9.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETH9Transfer)
				if err := _IWETH9.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IWETH9 *IWETH9Filterer) ParseTransfer(log types.Log) (*IWETH9Transfer, error) {
	event := new(IWETH9Transfer)
	if err := _IWETH9.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LowGasSafeMathMetaData contains all meta data concerning the LowGasSafeMath contract.
var LowGasSafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d2eaa6fb15818f513d54f617902047f3c8009f6e0cccfb26a75dd8354f03ca5164736f6c63430007060033",
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

// MulticallMetaData contains all meta data concerning the Multicall contract.
var MulticallMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MulticallABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallMetaData.ABI instead.
var MulticallABI = MulticallMetaData.ABI

// Multicall is an auto generated Go binding around an Ethereum contract.
type Multicall struct {
	MulticallCaller     // Read-only binding to the contract
	MulticallTransactor // Write-only binding to the contract
	MulticallFilterer   // Log filterer for contract events
}

// MulticallCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallSession struct {
	Contract     *Multicall        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallCallerSession struct {
	Contract *MulticallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MulticallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallTransactorSession struct {
	Contract     *MulticallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MulticallRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallRaw struct {
	Contract *Multicall // Generic contract binding to access the raw methods on
}

// MulticallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallCallerRaw struct {
	Contract *MulticallCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallTransactorRaw struct {
	Contract *MulticallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticall creates a new instance of Multicall, bound to a specific deployed contract.
func NewMulticall(address common.Address, backend bind.ContractBackend) (*Multicall, error) {
	contract, err := bindMulticall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicall{MulticallCaller: MulticallCaller{contract: contract}, MulticallTransactor: MulticallTransactor{contract: contract}, MulticallFilterer: MulticallFilterer{contract: contract}}, nil
}

// NewMulticallCaller creates a new read-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallCaller(address common.Address, caller bind.ContractCaller) (*MulticallCaller, error) {
	contract, err := bindMulticall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallCaller{contract: contract}, nil
}

// NewMulticallTransactor creates a new write-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallTransactor, error) {
	contract, err := bindMulticall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallTransactor{contract: contract}, nil
}

// NewMulticallFilterer creates a new log filterer instance of Multicall, bound to a specific deployed contract.
func NewMulticallFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallFilterer, error) {
	contract, err := bindMulticall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallFilterer{contract: contract}, nil
}

// bindMulticall binds a generic wrapper to an already deployed contract.
func bindMulticall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MulticallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.MulticallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transact(opts, method, params...)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Multicall *MulticallTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Multicall.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Multicall *MulticallSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Multicall.Contract.Multicall(&_Multicall.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Multicall *MulticallTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Multicall.Contract.Multicall(&_Multicall.TransactOpts, data)
}

// PeripheryImmutableStateMetaData contains all meta data concerning the PeripheryImmutableState contract.
var PeripheryImmutableStateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PeripheryImmutableStateABI is the input ABI used to generate the binding from.
// Deprecated: Use PeripheryImmutableStateMetaData.ABI instead.
var PeripheryImmutableStateABI = PeripheryImmutableStateMetaData.ABI

// PeripheryImmutableState is an auto generated Go binding around an Ethereum contract.
type PeripheryImmutableState struct {
	PeripheryImmutableStateCaller     // Read-only binding to the contract
	PeripheryImmutableStateTransactor // Write-only binding to the contract
	PeripheryImmutableStateFilterer   // Log filterer for contract events
}

// PeripheryImmutableStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeripheryImmutableStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryImmutableStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeripheryImmutableStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryImmutableStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeripheryImmutableStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryImmutableStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeripheryImmutableStateSession struct {
	Contract     *PeripheryImmutableState // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PeripheryImmutableStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeripheryImmutableStateCallerSession struct {
	Contract *PeripheryImmutableStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// PeripheryImmutableStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeripheryImmutableStateTransactorSession struct {
	Contract     *PeripheryImmutableStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// PeripheryImmutableStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeripheryImmutableStateRaw struct {
	Contract *PeripheryImmutableState // Generic contract binding to access the raw methods on
}

// PeripheryImmutableStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeripheryImmutableStateCallerRaw struct {
	Contract *PeripheryImmutableStateCaller // Generic read-only contract binding to access the raw methods on
}

// PeripheryImmutableStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeripheryImmutableStateTransactorRaw struct {
	Contract *PeripheryImmutableStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeripheryImmutableState creates a new instance of PeripheryImmutableState, bound to a specific deployed contract.
func NewPeripheryImmutableState(address common.Address, backend bind.ContractBackend) (*PeripheryImmutableState, error) {
	contract, err := bindPeripheryImmutableState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PeripheryImmutableState{PeripheryImmutableStateCaller: PeripheryImmutableStateCaller{contract: contract}, PeripheryImmutableStateTransactor: PeripheryImmutableStateTransactor{contract: contract}, PeripheryImmutableStateFilterer: PeripheryImmutableStateFilterer{contract: contract}}, nil
}

// NewPeripheryImmutableStateCaller creates a new read-only instance of PeripheryImmutableState, bound to a specific deployed contract.
func NewPeripheryImmutableStateCaller(address common.Address, caller bind.ContractCaller) (*PeripheryImmutableStateCaller, error) {
	contract, err := bindPeripheryImmutableState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryImmutableStateCaller{contract: contract}, nil
}

// NewPeripheryImmutableStateTransactor creates a new write-only instance of PeripheryImmutableState, bound to a specific deployed contract.
func NewPeripheryImmutableStateTransactor(address common.Address, transactor bind.ContractTransactor) (*PeripheryImmutableStateTransactor, error) {
	contract, err := bindPeripheryImmutableState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryImmutableStateTransactor{contract: contract}, nil
}

// NewPeripheryImmutableStateFilterer creates a new log filterer instance of PeripheryImmutableState, bound to a specific deployed contract.
func NewPeripheryImmutableStateFilterer(address common.Address, filterer bind.ContractFilterer) (*PeripheryImmutableStateFilterer, error) {
	contract, err := bindPeripheryImmutableState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeripheryImmutableStateFilterer{contract: contract}, nil
}

// bindPeripheryImmutableState binds a generic wrapper to an already deployed contract.
func bindPeripheryImmutableState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PeripheryImmutableStateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeripheryImmutableState *PeripheryImmutableStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeripheryImmutableState.Contract.PeripheryImmutableStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeripheryImmutableState *PeripheryImmutableStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryImmutableState.Contract.PeripheryImmutableStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeripheryImmutableState *PeripheryImmutableStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeripheryImmutableState.Contract.PeripheryImmutableStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PeripheryImmutableState *PeripheryImmutableStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PeripheryImmutableState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PeripheryImmutableState *PeripheryImmutableStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PeripheryImmutableState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PeripheryImmutableState *PeripheryImmutableStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PeripheryImmutableState.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PeripheryImmutableState.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateSession) WETH9() (common.Address, error) {
	return _PeripheryImmutableState.Contract.WETH9(&_PeripheryImmutableState.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateCallerSession) WETH9() (common.Address, error) {
	return _PeripheryImmutableState.Contract.WETH9(&_PeripheryImmutableState.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PeripheryImmutableState.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateSession) Factory() (common.Address, error) {
	return _PeripheryImmutableState.Contract.Factory(&_PeripheryImmutableState.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PeripheryImmutableState *PeripheryImmutableStateCallerSession) Factory() (common.Address, error) {
	return _PeripheryImmutableState.Contract.Factory(&_PeripheryImmutableState.CallOpts)
}

// PinitMetaData contains all meta data concerning the Pinit contract.
var PinitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nonfungiblePositionManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"createAndInitializePoolIfNecessary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonfungiblePositionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610d05380380610d0583398101604081905261002f91610072565b6001600160601b0319606093841b811660805291831b821660a05290911b1660c0526100b4565b80516001600160a01b038116811461006d57600080fd5b919050565b600080600060608486031215610086578283fd5b61008f84610056565b925061009d60208501610056565b91506100ab60408501610056565b90509250925092565b60805160601c60a05160601c60c05160601c610c0e6100f760003980610748525080609a528061054b5250806101dd52806102a452806107fb5250610c0e6000f3fe60806040526004361061008a5760003560e01c8063ac9650d811610059578063ac9650d81461014b578063b44a27221461016b578063c2e3140a14610180578063c45a015514610193578063f3995c67146101a8576100e2565b806313ead562146100e75780634659a494146101105780634aa4a4fc14610123578063a4a78f0c14610138576100e2565b366100e257336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146100e05760405162461bcd60e51b81526004016100d790610b21565b60405180910390fd5b005b600080fd5b6100fa6100f536600461088f565b6101bb565b6040516101079190610a93565b60405180910390f35b6100e061011e3660046108f1565b6104af565b34801561012f57600080fd5b506100fa610549565b6100e06101463660046108f1565b61056d565b61015e610159366004610951565b610606565b6040516101079190610aa7565b34801561017757600080fd5b506100fa610746565b6100e061018e3660046108f1565b61076a565b34801561019f57600080fd5b506100fa6107f9565b6100e06101b63660046108f1565b61081d565b6000836001600160a01b0316856001600160a01b0316106101db57600080fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631698ee828686866040518463ffffffff1660e01b815260040180846001600160a01b03168152602001836001600160a01b031681526020018262ffffff168152602001935050505060206040518083038186803b15801561026657600080fd5b505afa15801561027a573d6000803e3d6000fd5b505050506040513d602081101561029057600080fd5b505190506001600160a01b0381166103c6577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a16712958686866040518463ffffffff1660e01b815260040180846001600160a01b03168152602001836001600160a01b031681526020018262ffffff1681526020019350505050602060405180830381600087803b15801561032f57600080fd5b505af1158015610343573d6000803e3d6000fd5b505050506040513d602081101561035957600080fd5b50516040805163f637731d60e01b81526001600160a01b03858116600483015291519293509083169163f637731d9160248082019260009290919082900301818387803b1580156103a957600080fd5b505af11580156103bd573d6000803e3d6000fd5b505050506104a7565b6000816001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b15801561040157600080fd5b505afa158015610415573d6000803e3d6000fd5b505050506040513d60e081101561042b57600080fd5b505190506001600160a01b0381166104a557816001600160a01b031663f637731d846040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561048c57600080fd5b505af11580156104a0573d6000803e3d6000fd5b505050505b505b949350505050565b604080516323f2ebc360e21b815233600482015230602482015260448101879052606481018690526001608482015260ff851660a482015260c4810184905260e4810183905290516001600160a01b03881691638fcbaf0c9161010480830192600092919082900301818387803b15801561052957600080fd5b505af115801561053d573d6000803e3d6000fd5b50505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60408051636eb1769f60e11b81523360048201523060248201529051600019916001600160a01b0389169163dd62ed3e91604480820192602092909190829003018186803b1580156105be57600080fd5b505afa1580156105d2573d6000803e3d6000fd5b505050506040513d60208110156105e857600080fd5b505110156105fe576105fe8686868686866104af565b505050505050565b60608167ffffffffffffffff8111801561061f57600080fd5b5060405190808252806020026020018201604052801561065357816020015b606081526020019060019003908161063e5790505b50905060005b8281101561073f576000803086868581811061067157fe5b90506020028101906106839190610b44565b604051610691929190610a83565b600060405180830381855af49150503d80600081146106cc576040519150601f19603f3d011682016040523d82523d6000602084013e6106d1565b606091505b50915091508161071d576044815110156106ea57600080fd5b6004810190508080602001905181019061070491906109c0565b60405162461bcd60e51b81526004016100d79190610b07565b8084848151811061072a57fe5b60209081029190910101525050600101610659565b5092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60408051636eb1769f60e11b8152336004820152306024820152905186916001600160a01b0389169163dd62ed3e91604480820192602092909190829003018186803b1580156107b957600080fd5b505afa1580156107cd573d6000803e3d6000fd5b505050506040513d60208110156107e357600080fd5b505110156105fe576105fe86868686868661081d565b7f000000000000000000000000000000000000000000000000000000000000000081565b6040805163d505accf60e01b8152336004820152306024820152604481018790526064810186905260ff8516608482015260a4810184905260c4810183905290516001600160a01b0388169163d505accf9160e480830192600092919082900301818387803b15801561052957600080fd5b600080600080608085870312156108a4578384fd5b84356108af81610bc0565b935060208501356108bf81610bc0565b9250604085013562ffffff811681146108d6578283fd5b915060608501356108e681610bc0565b939692955090935050565b60008060008060008060c08789031215610909578182fd5b863561091481610bc0565b95506020870135945060408701359350606087013560ff81168114610937578283fd5b9598949750929560808101359460a0909101359350915050565b60008060208385031215610963578182fd5b823567ffffffffffffffff8082111561097a578384fd5b818501915085601f83011261098d578384fd5b81358181111561099b578485fd5b86602080830285010111156109ae578485fd5b60209290920196919550909350505050565b6000602082840312156109d1578081fd5b815167ffffffffffffffff808211156109e8578283fd5b818401915084601f8301126109fb578283fd5b815181811115610a0757fe5b604051601f8201601f191681016020018381118282101715610a2557fe5b604052818152838201602001871015610a3c578485fd5b610a4d826020830160208701610b90565b9695505050505050565b60008151808452610a6f816020860160208601610b90565b601f01601f19169290920160200192915050565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b6000602080830181845280855180835260408601915060408482028701019250838701855b82811015610afa57603f19888603018452610ae8858351610a57565b94509285019290850190600101610acc565b5092979650505050505050565b600060208252610b1a6020830184610a57565b9392505050565b6020808252600990820152684e6f7420574554483960b81b604082015260600190565b6000808335601e19843603018112610b5a578283fd5b83018035915067ffffffffffffffff821115610b74578283fd5b602001915036819003821315610b8957600080fd5b9250929050565b60005b83811015610bab578181015183820152602001610b93565b83811115610bba576000848401525b50505050565b6001600160a01b0381168114610bd557600080fd5b5056fea2646970667358221220f95cd94acb6dd5a6054900c74c7858dc2968da762c0cccee19926878c16ca53a64736f6c63430007060033",
}

// PinitABI is the input ABI used to generate the binding from.
// Deprecated: Use PinitMetaData.ABI instead.
var PinitABI = PinitMetaData.ABI

// PinitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PinitMetaData.Bin instead.
var PinitBin = PinitMetaData.Bin

// DeployPinit deploys a new Ethereum contract, binding an instance of Pinit to it.
func DeployPinit(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _WETH9 common.Address, _nonfungiblePositionManager common.Address) (common.Address, *types.Transaction, *Pinit, error) {
	parsed, err := PinitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PinitBin), backend, _factory, _WETH9, _nonfungiblePositionManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pinit{PinitCaller: PinitCaller{contract: contract}, PinitTransactor: PinitTransactor{contract: contract}, PinitFilterer: PinitFilterer{contract: contract}}, nil
}

// Pinit is an auto generated Go binding around an Ethereum contract.
type Pinit struct {
	PinitCaller     // Read-only binding to the contract
	PinitTransactor // Write-only binding to the contract
	PinitFilterer   // Log filterer for contract events
}

// PinitCaller is an auto generated read-only Go binding around an Ethereum contract.
type PinitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PinitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PinitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PinitSession struct {
	Contract     *Pinit            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PinitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PinitCallerSession struct {
	Contract *PinitCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PinitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PinitTransactorSession struct {
	Contract     *PinitTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PinitRaw is an auto generated low-level Go binding around an Ethereum contract.
type PinitRaw struct {
	Contract *Pinit // Generic contract binding to access the raw methods on
}

// PinitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PinitCallerRaw struct {
	Contract *PinitCaller // Generic read-only contract binding to access the raw methods on
}

// PinitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PinitTransactorRaw struct {
	Contract *PinitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPinit creates a new instance of Pinit, bound to a specific deployed contract.
func NewPinit(address common.Address, backend bind.ContractBackend) (*Pinit, error) {
	contract, err := bindPinit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pinit{PinitCaller: PinitCaller{contract: contract}, PinitTransactor: PinitTransactor{contract: contract}, PinitFilterer: PinitFilterer{contract: contract}}, nil
}

// NewPinitCaller creates a new read-only instance of Pinit, bound to a specific deployed contract.
func NewPinitCaller(address common.Address, caller bind.ContractCaller) (*PinitCaller, error) {
	contract, err := bindPinit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PinitCaller{contract: contract}, nil
}

// NewPinitTransactor creates a new write-only instance of Pinit, bound to a specific deployed contract.
func NewPinitTransactor(address common.Address, transactor bind.ContractTransactor) (*PinitTransactor, error) {
	contract, err := bindPinit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PinitTransactor{contract: contract}, nil
}

// NewPinitFilterer creates a new log filterer instance of Pinit, bound to a specific deployed contract.
func NewPinitFilterer(address common.Address, filterer bind.ContractFilterer) (*PinitFilterer, error) {
	contract, err := bindPinit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PinitFilterer{contract: contract}, nil
}

// bindPinit binds a generic wrapper to an already deployed contract.
func bindPinit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PinitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pinit *PinitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pinit.Contract.PinitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pinit *PinitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pinit.Contract.PinitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pinit *PinitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pinit.Contract.PinitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pinit *PinitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pinit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pinit *PinitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pinit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pinit *PinitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pinit.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Pinit *PinitCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pinit.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Pinit *PinitSession) WETH9() (common.Address, error) {
	return _Pinit.Contract.WETH9(&_Pinit.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Pinit *PinitCallerSession) WETH9() (common.Address, error) {
	return _Pinit.Contract.WETH9(&_Pinit.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pinit *PinitCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pinit.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pinit *PinitSession) Factory() (common.Address, error) {
	return _Pinit.Contract.Factory(&_Pinit.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pinit *PinitCallerSession) Factory() (common.Address, error) {
	return _Pinit.Contract.Factory(&_Pinit.CallOpts)
}

// NonfungiblePositionManager is a free data retrieval call binding the contract method 0xb44a2722.
//
// Solidity: function nonfungiblePositionManager() view returns(address)
func (_Pinit *PinitCaller) NonfungiblePositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pinit.contract.Call(opts, &out, "nonfungiblePositionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonfungiblePositionManager is a free data retrieval call binding the contract method 0xb44a2722.
//
// Solidity: function nonfungiblePositionManager() view returns(address)
func (_Pinit *PinitSession) NonfungiblePositionManager() (common.Address, error) {
	return _Pinit.Contract.NonfungiblePositionManager(&_Pinit.CallOpts)
}

// NonfungiblePositionManager is a free data retrieval call binding the contract method 0xb44a2722.
//
// Solidity: function nonfungiblePositionManager() view returns(address)
func (_Pinit *PinitCallerSession) NonfungiblePositionManager() (common.Address, error) {
	return _Pinit.Contract.NonfungiblePositionManager(&_Pinit.CallOpts)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_Pinit *PinitTransactor) CreateAndInitializePoolIfNecessary(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _Pinit.contract.Transact(opts, "createAndInitializePoolIfNecessary", token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_Pinit *PinitSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _Pinit.Contract.CreateAndInitializePoolIfNecessary(&_Pinit.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_Pinit *PinitTransactorSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _Pinit.Contract.CreateAndInitializePoolIfNecessary(&_Pinit.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Pinit *PinitTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Pinit.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Pinit *PinitSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Pinit.Contract.Multicall(&_Pinit.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Pinit *PinitTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Pinit.Contract.Multicall(&_Pinit.TransactOpts, data)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.Contract.SelfPermit(&_Pinit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.Contract.SelfPermit(&_Pinit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.Contract.SelfPermitAllowed(&_Pinit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.Contract.SelfPermitAllowed(&_Pinit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.Contract.SelfPermitAllowedIfNecessary(&_Pinit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.Contract.SelfPermitAllowedIfNecessary(&_Pinit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.Contract.SelfPermitIfNecessary(&_Pinit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Pinit *PinitTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pinit.Contract.SelfPermitIfNecessary(&_Pinit.TransactOpts, token, value, deadline, v, r, s)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pinit *PinitTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pinit.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pinit *PinitSession) Receive() (*types.Transaction, error) {
	return _Pinit.Contract.Receive(&_Pinit.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pinit *PinitTransactorSession) Receive() (*types.Transaction, error) {
	return _Pinit.Contract.Receive(&_Pinit.TransactOpts)
}

// PoolAddressMetaData contains all meta data concerning the PoolAddress contract.
var PoolAddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122077dbd08890c31664d133f8f2d84c30e80824600f1b58281a904356aa21ddfd2764736f6c63430007060033",
}

// PoolAddressABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolAddressMetaData.ABI instead.
var PoolAddressABI = PoolAddressMetaData.ABI

// PoolAddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolAddressMetaData.Bin instead.
var PoolAddressBin = PoolAddressMetaData.Bin

// DeployPoolAddress deploys a new Ethereum contract, binding an instance of PoolAddress to it.
func DeployPoolAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PoolAddress, error) {
	parsed, err := PoolAddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolAddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolAddress{PoolAddressCaller: PoolAddressCaller{contract: contract}, PoolAddressTransactor: PoolAddressTransactor{contract: contract}, PoolAddressFilterer: PoolAddressFilterer{contract: contract}}, nil
}

// PoolAddress is an auto generated Go binding around an Ethereum contract.
type PoolAddress struct {
	PoolAddressCaller     // Read-only binding to the contract
	PoolAddressTransactor // Write-only binding to the contract
	PoolAddressFilterer   // Log filterer for contract events
}

// PoolAddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolAddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolAddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolAddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolAddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolAddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolAddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolAddressSession struct {
	Contract     *PoolAddress      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolAddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolAddressCallerSession struct {
	Contract *PoolAddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PoolAddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolAddressTransactorSession struct {
	Contract     *PoolAddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoolAddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolAddressRaw struct {
	Contract *PoolAddress // Generic contract binding to access the raw methods on
}

// PoolAddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolAddressCallerRaw struct {
	Contract *PoolAddressCaller // Generic read-only contract binding to access the raw methods on
}

// PoolAddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolAddressTransactorRaw struct {
	Contract *PoolAddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolAddress creates a new instance of PoolAddress, bound to a specific deployed contract.
func NewPoolAddress(address common.Address, backend bind.ContractBackend) (*PoolAddress, error) {
	contract, err := bindPoolAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolAddress{PoolAddressCaller: PoolAddressCaller{contract: contract}, PoolAddressTransactor: PoolAddressTransactor{contract: contract}, PoolAddressFilterer: PoolAddressFilterer{contract: contract}}, nil
}

// NewPoolAddressCaller creates a new read-only instance of PoolAddress, bound to a specific deployed contract.
func NewPoolAddressCaller(address common.Address, caller bind.ContractCaller) (*PoolAddressCaller, error) {
	contract, err := bindPoolAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolAddressCaller{contract: contract}, nil
}

// NewPoolAddressTransactor creates a new write-only instance of PoolAddress, bound to a specific deployed contract.
func NewPoolAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolAddressTransactor, error) {
	contract, err := bindPoolAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolAddressTransactor{contract: contract}, nil
}

// NewPoolAddressFilterer creates a new log filterer instance of PoolAddress, bound to a specific deployed contract.
func NewPoolAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolAddressFilterer, error) {
	contract, err := bindPoolAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolAddressFilterer{contract: contract}, nil
}

// bindPoolAddress binds a generic wrapper to an already deployed contract.
func bindPoolAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolAddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolAddress *PoolAddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolAddress.Contract.PoolAddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolAddress *PoolAddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolAddress.Contract.PoolAddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolAddress *PoolAddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolAddress.Contract.PoolAddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolAddress *PoolAddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolAddress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolAddress *PoolAddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolAddress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolAddress *PoolAddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolAddress.Contract.contract.Transact(opts, method, params...)
}

// PoolInitializerMetaData contains all meta data concerning the PoolInitializer contract.
var PoolInitializerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"createAndInitializePoolIfNecessary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PoolInitializerABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolInitializerMetaData.ABI instead.
var PoolInitializerABI = PoolInitializerMetaData.ABI

// PoolInitializer is an auto generated Go binding around an Ethereum contract.
type PoolInitializer struct {
	PoolInitializerCaller     // Read-only binding to the contract
	PoolInitializerTransactor // Write-only binding to the contract
	PoolInitializerFilterer   // Log filterer for contract events
}

// PoolInitializerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolInitializerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolInitializerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolInitializerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolInitializerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolInitializerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolInitializerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolInitializerSession struct {
	Contract     *PoolInitializer  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolInitializerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolInitializerCallerSession struct {
	Contract *PoolInitializerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PoolInitializerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolInitializerTransactorSession struct {
	Contract     *PoolInitializerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PoolInitializerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolInitializerRaw struct {
	Contract *PoolInitializer // Generic contract binding to access the raw methods on
}

// PoolInitializerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolInitializerCallerRaw struct {
	Contract *PoolInitializerCaller // Generic read-only contract binding to access the raw methods on
}

// PoolInitializerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolInitializerTransactorRaw struct {
	Contract *PoolInitializerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolInitializer creates a new instance of PoolInitializer, bound to a specific deployed contract.
func NewPoolInitializer(address common.Address, backend bind.ContractBackend) (*PoolInitializer, error) {
	contract, err := bindPoolInitializer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolInitializer{PoolInitializerCaller: PoolInitializerCaller{contract: contract}, PoolInitializerTransactor: PoolInitializerTransactor{contract: contract}, PoolInitializerFilterer: PoolInitializerFilterer{contract: contract}}, nil
}

// NewPoolInitializerCaller creates a new read-only instance of PoolInitializer, bound to a specific deployed contract.
func NewPoolInitializerCaller(address common.Address, caller bind.ContractCaller) (*PoolInitializerCaller, error) {
	contract, err := bindPoolInitializer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolInitializerCaller{contract: contract}, nil
}

// NewPoolInitializerTransactor creates a new write-only instance of PoolInitializer, bound to a specific deployed contract.
func NewPoolInitializerTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolInitializerTransactor, error) {
	contract, err := bindPoolInitializer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolInitializerTransactor{contract: contract}, nil
}

// NewPoolInitializerFilterer creates a new log filterer instance of PoolInitializer, bound to a specific deployed contract.
func NewPoolInitializerFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolInitializerFilterer, error) {
	contract, err := bindPoolInitializer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolInitializerFilterer{contract: contract}, nil
}

// bindPoolInitializer binds a generic wrapper to an already deployed contract.
func bindPoolInitializer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolInitializerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolInitializer *PoolInitializerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolInitializer.Contract.PoolInitializerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolInitializer *PoolInitializerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolInitializer.Contract.PoolInitializerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolInitializer *PoolInitializerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolInitializer.Contract.PoolInitializerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolInitializer *PoolInitializerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolInitializer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolInitializer *PoolInitializerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolInitializer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolInitializer *PoolInitializerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolInitializer.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PoolInitializer *PoolInitializerCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolInitializer.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PoolInitializer *PoolInitializerSession) WETH9() (common.Address, error) {
	return _PoolInitializer.Contract.WETH9(&_PoolInitializer.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PoolInitializer *PoolInitializerCallerSession) WETH9() (common.Address, error) {
	return _PoolInitializer.Contract.WETH9(&_PoolInitializer.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PoolInitializer *PoolInitializerCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolInitializer.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PoolInitializer *PoolInitializerSession) Factory() (common.Address, error) {
	return _PoolInitializer.Contract.Factory(&_PoolInitializer.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PoolInitializer *PoolInitializerCallerSession) Factory() (common.Address, error) {
	return _PoolInitializer.Contract.Factory(&_PoolInitializer.CallOpts)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_PoolInitializer *PoolInitializerTransactor) CreateAndInitializePoolIfNecessary(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _PoolInitializer.contract.Transact(opts, "createAndInitializePoolIfNecessary", token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_PoolInitializer *PoolInitializerSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _PoolInitializer.Contract.CreateAndInitializePoolIfNecessary(&_PoolInitializer.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_PoolInitializer *PoolInitializerTransactorSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _PoolInitializer.Contract.CreateAndInitializePoolIfNecessary(&_PoolInitializer.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// SelfPermitMetaData contains all meta data concerning the SelfPermit contract.
var SelfPermitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SelfPermitABI is the input ABI used to generate the binding from.
// Deprecated: Use SelfPermitMetaData.ABI instead.
var SelfPermitABI = SelfPermitMetaData.ABI

// SelfPermit is an auto generated Go binding around an Ethereum contract.
type SelfPermit struct {
	SelfPermitCaller     // Read-only binding to the contract
	SelfPermitTransactor // Write-only binding to the contract
	SelfPermitFilterer   // Log filterer for contract events
}

// SelfPermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type SelfPermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfPermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SelfPermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfPermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SelfPermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfPermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SelfPermitSession struct {
	Contract     *SelfPermit       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SelfPermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SelfPermitCallerSession struct {
	Contract *SelfPermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SelfPermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SelfPermitTransactorSession struct {
	Contract     *SelfPermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SelfPermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type SelfPermitRaw struct {
	Contract *SelfPermit // Generic contract binding to access the raw methods on
}

// SelfPermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SelfPermitCallerRaw struct {
	Contract *SelfPermitCaller // Generic read-only contract binding to access the raw methods on
}

// SelfPermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SelfPermitTransactorRaw struct {
	Contract *SelfPermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSelfPermit creates a new instance of SelfPermit, bound to a specific deployed contract.
func NewSelfPermit(address common.Address, backend bind.ContractBackend) (*SelfPermit, error) {
	contract, err := bindSelfPermit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SelfPermit{SelfPermitCaller: SelfPermitCaller{contract: contract}, SelfPermitTransactor: SelfPermitTransactor{contract: contract}, SelfPermitFilterer: SelfPermitFilterer{contract: contract}}, nil
}

// NewSelfPermitCaller creates a new read-only instance of SelfPermit, bound to a specific deployed contract.
func NewSelfPermitCaller(address common.Address, caller bind.ContractCaller) (*SelfPermitCaller, error) {
	contract, err := bindSelfPermit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SelfPermitCaller{contract: contract}, nil
}

// NewSelfPermitTransactor creates a new write-only instance of SelfPermit, bound to a specific deployed contract.
func NewSelfPermitTransactor(address common.Address, transactor bind.ContractTransactor) (*SelfPermitTransactor, error) {
	contract, err := bindSelfPermit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SelfPermitTransactor{contract: contract}, nil
}

// NewSelfPermitFilterer creates a new log filterer instance of SelfPermit, bound to a specific deployed contract.
func NewSelfPermitFilterer(address common.Address, filterer bind.ContractFilterer) (*SelfPermitFilterer, error) {
	contract, err := bindSelfPermit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SelfPermitFilterer{contract: contract}, nil
}

// bindSelfPermit binds a generic wrapper to an already deployed contract.
func bindSelfPermit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SelfPermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfPermit *SelfPermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfPermit.Contract.SelfPermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfPermit *SelfPermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfPermit *SelfPermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfPermit *SelfPermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfPermit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfPermit *SelfPermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfPermit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfPermit *SelfPermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfPermit.Contract.contract.Transact(opts, method, params...)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermit(&_SelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermit(&_SelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitAllowed(&_SelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitAllowed(&_SelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitAllowedIfNecessary(&_SelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitAllowedIfNecessary(&_SelfPermit.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitIfNecessary(&_SelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SelfPermit *SelfPermitTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SelfPermit.Contract.SelfPermitIfNecessary(&_SelfPermit.TransactOpts, token, value, deadline, v, r, s)
}

// TransferHelperMetaData contains all meta data concerning the TransferHelper contract.
var TransferHelperMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220515d61339156c1a27013030047c239d79c94a8cd1b4c6d6f1c235e6653a21af764736f6c63430007060033",
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
