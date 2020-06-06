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

// RepTokABI is the input ABI used to generate the binding from.
const RepTokABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_repBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"FundedAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerLookupName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"faucet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIsMigratingFromLegacy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLegacyRepToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalMigrated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalTheoreticalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_spenders\",\"type\":\"address[]\"}],\"name\":\"migrateAllowancesFromLegacyRep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_holders\",\"type\":\"address[]\"}],\"name\":\"migrateBalancesFromLegacyRep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateIn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIReputationToken\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateOutByPayout\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountMigrated\",\"type\":\"uint256\"}],\"name\":\"mintForReportingParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedDisputeWindowTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedMarketTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedReportingParticipantTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedUniverseTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateParentTotalTheoreticalSupply\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIReputationToken\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"updateSiblingMigrationTotal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RepTok is an auto generated Go binding around an Ethereum contract.
type RepTok struct {
	RepTokCaller     // Read-only binding to the contract
	RepTokTransactor // Write-only binding to the contract
	RepTokFilterer   // Log filterer for contract events
}

// RepTokCaller is an auto generated read-only Go binding around an Ethereum contract.
type RepTokCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepTokTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RepTokTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepTokFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RepTokFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepTokSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RepTokSession struct {
	Contract     *RepTok           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RepTokCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RepTokCallerSession struct {
	Contract *RepTokCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RepTokTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RepTokTransactorSession struct {
	Contract     *RepTokTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RepTokRaw is an auto generated low-level Go binding around an Ethereum contract.
type RepTokRaw struct {
	Contract *RepTok // Generic contract binding to access the raw methods on
}

// RepTokCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RepTokCallerRaw struct {
	Contract *RepTokCaller // Generic read-only contract binding to access the raw methods on
}

// RepTokTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RepTokTransactorRaw struct {
	Contract *RepTokTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRepTok creates a new instance of RepTok, bound to a specific deployed contract.
func NewRepTok(address common.Address, backend bind.ContractBackend) (*RepTok, error) {
	contract, err := bindRepTok(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RepTok{RepTokCaller: RepTokCaller{contract: contract}, RepTokTransactor: RepTokTransactor{contract: contract}, RepTokFilterer: RepTokFilterer{contract: contract}}, nil
}

// NewRepTokCaller creates a new read-only instance of RepTok, bound to a specific deployed contract.
func NewRepTokCaller(address common.Address, caller bind.ContractCaller) (*RepTokCaller, error) {
	contract, err := bindRepTok(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RepTokCaller{contract: contract}, nil
}

// NewRepTokTransactor creates a new write-only instance of RepTok, bound to a specific deployed contract.
func NewRepTokTransactor(address common.Address, transactor bind.ContractTransactor) (*RepTokTransactor, error) {
	contract, err := bindRepTok(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RepTokTransactor{contract: contract}, nil
}

// NewRepTokFilterer creates a new log filterer instance of RepTok, bound to a specific deployed contract.
func NewRepTokFilterer(address common.Address, filterer bind.ContractFilterer) (*RepTokFilterer, error) {
	contract, err := bindRepTok(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RepTokFilterer{contract: contract}, nil
}

// bindRepTok binds a generic wrapper to an already deployed contract.
func bindRepTok(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RepTokABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepTok *RepTokRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RepTok.Contract.RepTokCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepTok *RepTokRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepTok.Contract.RepTokTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepTok *RepTokRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepTok.Contract.RepTokTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepTok *RepTokCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RepTok.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepTok *RepTokTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepTok.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepTok *RepTokTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepTok.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_RepTok *RepTokCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_RepTok *RepTokSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RepTok.Contract.Allowance(&_RepTok.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_RepTok *RepTokCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RepTok.Contract.Allowance(&_RepTok.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_RepTok *RepTokCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "allowances", arg0, arg1)
	return *ret0, err
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_RepTok *RepTokSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RepTok.Contract.Allowances(&_RepTok.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_RepTok *RepTokCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RepTok.Contract.Allowances(&_RepTok.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_RepTok *RepTokCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "balanceOf", _account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_RepTok *RepTokSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _RepTok.Contract.BalanceOf(&_RepTok.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_RepTok *RepTokCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _RepTok.Contract.BalanceOf(&_RepTok.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RepTok *RepTokCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RepTok *RepTokSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RepTok.Contract.Balances(&_RepTok.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RepTok *RepTokCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RepTok.Contract.Balances(&_RepTok.CallOpts, arg0)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() view returns(bytes32)
func (_RepTok *RepTokCaller) ControllerLookupName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "controllerLookupName")
	return *ret0, err
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() view returns(bytes32)
func (_RepTok *RepTokSession) ControllerLookupName() ([32]byte, error) {
	return _RepTok.Contract.ControllerLookupName(&_RepTok.CallOpts)
}

// ControllerLookupName is a free data retrieval call binding the contract method 0xbef72fa2.
//
// Solidity: function controllerLookupName() view returns(bytes32)
func (_RepTok *RepTokCallerSession) ControllerLookupName() ([32]byte, error) {
	return _RepTok.Contract.ControllerLookupName(&_RepTok.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RepTok *RepTokCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RepTok *RepTokSession) Decimals() (uint8, error) {
	return _RepTok.Contract.Decimals(&_RepTok.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RepTok *RepTokCallerSession) Decimals() (uint8, error) {
	return _RepTok.Contract.Decimals(&_RepTok.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_RepTok *RepTokCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_RepTok *RepTokSession) GetInitialized() (bool, error) {
	return _RepTok.Contract.GetInitialized(&_RepTok.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_RepTok *RepTokCallerSession) GetInitialized() (bool, error) {
	return _RepTok.Contract.GetInitialized(&_RepTok.CallOpts)
}

// GetIsMigratingFromLegacy is a free data retrieval call binding the contract method 0x7cf99c33.
//
// Solidity: function getIsMigratingFromLegacy() view returns(bool)
func (_RepTok *RepTokCaller) GetIsMigratingFromLegacy(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "getIsMigratingFromLegacy")
	return *ret0, err
}

// GetIsMigratingFromLegacy is a free data retrieval call binding the contract method 0x7cf99c33.
//
// Solidity: function getIsMigratingFromLegacy() view returns(bool)
func (_RepTok *RepTokSession) GetIsMigratingFromLegacy() (bool, error) {
	return _RepTok.Contract.GetIsMigratingFromLegacy(&_RepTok.CallOpts)
}

// GetIsMigratingFromLegacy is a free data retrieval call binding the contract method 0x7cf99c33.
//
// Solidity: function getIsMigratingFromLegacy() view returns(bool)
func (_RepTok *RepTokCallerSession) GetIsMigratingFromLegacy() (bool, error) {
	return _RepTok.Contract.GetIsMigratingFromLegacy(&_RepTok.CallOpts)
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() pure returns(address)
func (_RepTok *RepTokCaller) GetLegacyRepToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "getLegacyRepToken")
	return *ret0, err
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() pure returns(address)
func (_RepTok *RepTokSession) GetLegacyRepToken() (common.Address, error) {
	return _RepTok.Contract.GetLegacyRepToken(&_RepTok.CallOpts)
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() pure returns(address)
func (_RepTok *RepTokCallerSession) GetLegacyRepToken() (common.Address, error) {
	return _RepTok.Contract.GetLegacyRepToken(&_RepTok.CallOpts)
}

// GetTargetSupply is a free data retrieval call binding the contract method 0xdea6aec7.
//
// Solidity: function getTargetSupply() view returns(uint256)
func (_RepTok *RepTokCaller) GetTargetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "getTargetSupply")
	return *ret0, err
}

// GetTargetSupply is a free data retrieval call binding the contract method 0xdea6aec7.
//
// Solidity: function getTargetSupply() view returns(uint256)
func (_RepTok *RepTokSession) GetTargetSupply() (*big.Int, error) {
	return _RepTok.Contract.GetTargetSupply(&_RepTok.CallOpts)
}

// GetTargetSupply is a free data retrieval call binding the contract method 0xdea6aec7.
//
// Solidity: function getTargetSupply() view returns(uint256)
func (_RepTok *RepTokCallerSession) GetTargetSupply() (*big.Int, error) {
	return _RepTok.Contract.GetTargetSupply(&_RepTok.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256)
func (_RepTok *RepTokCaller) GetTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "getTimestamp")
	return *ret0, err
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256)
func (_RepTok *RepTokSession) GetTimestamp() (*big.Int, error) {
	return _RepTok.Contract.GetTimestamp(&_RepTok.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256)
func (_RepTok *RepTokCallerSession) GetTimestamp() (*big.Int, error) {
	return _RepTok.Contract.GetTimestamp(&_RepTok.CallOpts)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() view returns(uint256)
func (_RepTok *RepTokCaller) GetTotalMigrated(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "getTotalMigrated")
	return *ret0, err
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() view returns(uint256)
func (_RepTok *RepTokSession) GetTotalMigrated() (*big.Int, error) {
	return _RepTok.Contract.GetTotalMigrated(&_RepTok.CallOpts)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() view returns(uint256)
func (_RepTok *RepTokCallerSession) GetTotalMigrated() (*big.Int, error) {
	return _RepTok.Contract.GetTotalMigrated(&_RepTok.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() view returns(uint256)
func (_RepTok *RepTokCaller) GetTotalTheoreticalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "getTotalTheoreticalSupply")
	return *ret0, err
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() view returns(uint256)
func (_RepTok *RepTokSession) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _RepTok.Contract.GetTotalTheoreticalSupply(&_RepTok.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() view returns(uint256)
func (_RepTok *RepTokCallerSession) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _RepTok.Contract.GetTotalTheoreticalSupply(&_RepTok.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(bytes32)
func (_RepTok *RepTokCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(bytes32)
func (_RepTok *RepTokSession) GetTypeName() ([32]byte, error) {
	return _RepTok.Contract.GetTypeName(&_RepTok.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(bytes32)
func (_RepTok *RepTokCallerSession) GetTypeName() ([32]byte, error) {
	return _RepTok.Contract.GetTypeName(&_RepTok.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() view returns(address)
func (_RepTok *RepTokCaller) GetUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "getUniverse")
	return *ret0, err
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() view returns(address)
func (_RepTok *RepTokSession) GetUniverse() (common.Address, error) {
	return _RepTok.Contract.GetUniverse(&_RepTok.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() view returns(address)
func (_RepTok *RepTokCallerSession) GetUniverse() (common.Address, error) {
	return _RepTok.Contract.GetUniverse(&_RepTok.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RepTok *RepTokCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RepTok *RepTokSession) Name() (string, error) {
	return _RepTok.Contract.Name(&_RepTok.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RepTok *RepTokCallerSession) Name() (string, error) {
	return _RepTok.Contract.Name(&_RepTok.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RepTok *RepTokCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RepTok *RepTokSession) Symbol() (string, error) {
	return _RepTok.Contract.Symbol(&_RepTok.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RepTok *RepTokCallerSession) Symbol() (string, error) {
	return _RepTok.Contract.Symbol(&_RepTok.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RepTok *RepTokCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepTok.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RepTok *RepTokSession) TotalSupply() (*big.Int, error) {
	return _RepTok.Contract.TotalSupply(&_RepTok.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RepTok *RepTokCallerSession) TotalSupply() (*big.Int, error) {
	return _RepTok.Contract.TotalSupply(&_RepTok.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_RepTok *RepTokTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_RepTok *RepTokSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.Approve(&_RepTok.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_RepTok *RepTokTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.Approve(&_RepTok.TransactOpts, _spender, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_RepTok *RepTokTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_RepTok *RepTokSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.DecreaseAllowance(&_RepTok.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_RepTok *RepTokTransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.DecreaseAllowance(&_RepTok.TransactOpts, _spender, _subtractedValue)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(uint256 _amount) returns(bool)
func (_RepTok *RepTokTransactor) Faucet(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "faucet", _amount)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(uint256 _amount) returns(bool)
func (_RepTok *RepTokSession) Faucet(_amount *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.Faucet(&_RepTok.TransactOpts, _amount)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(uint256 _amount) returns(bool)
func (_RepTok *RepTokTransactorSession) Faucet(_amount *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.Faucet(&_RepTok.TransactOpts, _amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_RepTok *RepTokTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_RepTok *RepTokSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.IncreaseAllowance(&_RepTok.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_RepTok *RepTokTransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.IncreaseAllowance(&_RepTok.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _universe) returns()
func (_RepTok *RepTokTransactor) Initialize(opts *bind.TransactOpts, _universe common.Address) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "initialize", _universe)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _universe) returns()
func (_RepTok *RepTokSession) Initialize(_universe common.Address) (*types.Transaction, error) {
	return _RepTok.Contract.Initialize(&_RepTok.TransactOpts, _universe)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _universe) returns()
func (_RepTok *RepTokTransactorSession) Initialize(_universe common.Address) (*types.Transaction, error) {
	return _RepTok.Contract.Initialize(&_RepTok.TransactOpts, _universe)
}

// MigrateAllowancesFromLegacyRep is a paid mutator transaction binding the contract method 0xde4c0574.
//
// Solidity: function migrateAllowancesFromLegacyRep(address[] _owners, address[] _spenders) returns(bool)
func (_RepTok *RepTokTransactor) MigrateAllowancesFromLegacyRep(opts *bind.TransactOpts, _owners []common.Address, _spenders []common.Address) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "migrateAllowancesFromLegacyRep", _owners, _spenders)
}

// MigrateAllowancesFromLegacyRep is a paid mutator transaction binding the contract method 0xde4c0574.
//
// Solidity: function migrateAllowancesFromLegacyRep(address[] _owners, address[] _spenders) returns(bool)
func (_RepTok *RepTokSession) MigrateAllowancesFromLegacyRep(_owners []common.Address, _spenders []common.Address) (*types.Transaction, error) {
	return _RepTok.Contract.MigrateAllowancesFromLegacyRep(&_RepTok.TransactOpts, _owners, _spenders)
}

// MigrateAllowancesFromLegacyRep is a paid mutator transaction binding the contract method 0xde4c0574.
//
// Solidity: function migrateAllowancesFromLegacyRep(address[] _owners, address[] _spenders) returns(bool)
func (_RepTok *RepTokTransactorSession) MigrateAllowancesFromLegacyRep(_owners []common.Address, _spenders []common.Address) (*types.Transaction, error) {
	return _RepTok.Contract.MigrateAllowancesFromLegacyRep(&_RepTok.TransactOpts, _owners, _spenders)
}

// MigrateBalancesFromLegacyRep is a paid mutator transaction binding the contract method 0x7f686259.
//
// Solidity: function migrateBalancesFromLegacyRep(address[] _holders) returns(bool)
func (_RepTok *RepTokTransactor) MigrateBalancesFromLegacyRep(opts *bind.TransactOpts, _holders []common.Address) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "migrateBalancesFromLegacyRep", _holders)
}

// MigrateBalancesFromLegacyRep is a paid mutator transaction binding the contract method 0x7f686259.
//
// Solidity: function migrateBalancesFromLegacyRep(address[] _holders) returns(bool)
func (_RepTok *RepTokSession) MigrateBalancesFromLegacyRep(_holders []common.Address) (*types.Transaction, error) {
	return _RepTok.Contract.MigrateBalancesFromLegacyRep(&_RepTok.TransactOpts, _holders)
}

// MigrateBalancesFromLegacyRep is a paid mutator transaction binding the contract method 0x7f686259.
//
// Solidity: function migrateBalancesFromLegacyRep(address[] _holders) returns(bool)
func (_RepTok *RepTokTransactorSession) MigrateBalancesFromLegacyRep(_holders []common.Address) (*types.Transaction, error) {
	return _RepTok.Contract.MigrateBalancesFromLegacyRep(&_RepTok.TransactOpts, _holders)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(address _reporter, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactor) MigrateIn(opts *bind.TransactOpts, _reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "migrateIn", _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(address _reporter, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokSession) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.MigrateIn(&_RepTok.TransactOpts, _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(address _reporter, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactorSession) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.MigrateIn(&_RepTok.TransactOpts, _reporter, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactor) MigrateOut(opts *bind.TransactOpts, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "migrateOut", _destination, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokSession) MigrateOut(_destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.MigrateOut(&_RepTok.TransactOpts, _destination, _attotokens)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactorSession) MigrateOut(_destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.MigrateOut(&_RepTok.TransactOpts, _destination, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x5b66876a.
//
// Solidity: function migrateOutByPayout(uint256[] _payoutNumerators, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactor) MigrateOutByPayout(opts *bind.TransactOpts, _payoutNumerators []*big.Int, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "migrateOutByPayout", _payoutNumerators, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x5b66876a.
//
// Solidity: function migrateOutByPayout(uint256[] _payoutNumerators, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokSession) MigrateOutByPayout(_payoutNumerators []*big.Int, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.MigrateOutByPayout(&_RepTok.TransactOpts, _payoutNumerators, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x5b66876a.
//
// Solidity: function migrateOutByPayout(uint256[] _payoutNumerators, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactorSession) MigrateOutByPayout(_payoutNumerators []*big.Int, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.MigrateOutByPayout(&_RepTok.TransactOpts, _payoutNumerators, _attotokens)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(uint256 _amountMigrated) returns(bool)
func (_RepTok *RepTokTransactor) MintForReportingParticipant(opts *bind.TransactOpts, _amountMigrated *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "mintForReportingParticipant", _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(uint256 _amountMigrated) returns(bool)
func (_RepTok *RepTokSession) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.MintForReportingParticipant(&_RepTok.TransactOpts, _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(uint256 _amountMigrated) returns(bool)
func (_RepTok *RepTokTransactorSession) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.MintForReportingParticipant(&_RepTok.TransactOpts, _amountMigrated)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_RepTok *RepTokTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_RepTok *RepTokSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.Transfer(&_RepTok.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_RepTok *RepTokTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.Transfer(&_RepTok.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_RepTok *RepTokTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_RepTok *RepTokSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.TransferFrom(&_RepTok.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_RepTok *RepTokTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.TransferFrom(&_RepTok.TransactOpts, _from, _to, _value)
}

// TrustedDisputeWindowTransfer is a paid mutator transaction binding the contract method 0x721bb170.
//
// Solidity: function trustedDisputeWindowTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactor) TrustedDisputeWindowTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "trustedDisputeWindowTransfer", _source, _destination, _attotokens)
}

// TrustedDisputeWindowTransfer is a paid mutator transaction binding the contract method 0x721bb170.
//
// Solidity: function trustedDisputeWindowTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokSession) TrustedDisputeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.TrustedDisputeWindowTransfer(&_RepTok.TransactOpts, _source, _destination, _attotokens)
}

// TrustedDisputeWindowTransfer is a paid mutator transaction binding the contract method 0x721bb170.
//
// Solidity: function trustedDisputeWindowTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactorSession) TrustedDisputeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.TrustedDisputeWindowTransfer(&_RepTok.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactor) TrustedMarketTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "trustedMarketTransfer", _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokSession) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.TrustedMarketTransfer(&_RepTok.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactorSession) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.TrustedMarketTransfer(&_RepTok.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactor) TrustedReportingParticipantTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "trustedReportingParticipantTransfer", _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokSession) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.TrustedReportingParticipantTransfer(&_RepTok.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactorSession) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.TrustedReportingParticipantTransfer(&_RepTok.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactor) TrustedUniverseTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "trustedUniverseTransfer", _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokSession) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.TrustedUniverseTransfer(&_RepTok.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_RepTok *RepTokTransactorSession) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _RepTok.Contract.TrustedUniverseTransfer(&_RepTok.TransactOpts, _source, _destination, _attotokens)
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_RepTok *RepTokTransactor) UpdateParentTotalTheoreticalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "updateParentTotalTheoreticalSupply")
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_RepTok *RepTokSession) UpdateParentTotalTheoreticalSupply() (*types.Transaction, error) {
	return _RepTok.Contract.UpdateParentTotalTheoreticalSupply(&_RepTok.TransactOpts)
}

// UpdateParentTotalTheoreticalSupply is a paid mutator transaction binding the contract method 0xa819515d.
//
// Solidity: function updateParentTotalTheoreticalSupply() returns(bool)
func (_RepTok *RepTokTransactorSession) UpdateParentTotalTheoreticalSupply() (*types.Transaction, error) {
	return _RepTok.Contract.UpdateParentTotalTheoreticalSupply(&_RepTok.TransactOpts)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(address _token) returns(bool)
func (_RepTok *RepTokTransactor) UpdateSiblingMigrationTotal(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _RepTok.contract.Transact(opts, "updateSiblingMigrationTotal", _token)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(address _token) returns(bool)
func (_RepTok *RepTokSession) UpdateSiblingMigrationTotal(_token common.Address) (*types.Transaction, error) {
	return _RepTok.Contract.UpdateSiblingMigrationTotal(&_RepTok.TransactOpts, _token)
}

// UpdateSiblingMigrationTotal is a paid mutator transaction binding the contract method 0xd9d3e07d.
//
// Solidity: function updateSiblingMigrationTotal(address _token) returns(bool)
func (_RepTok *RepTokTransactorSession) UpdateSiblingMigrationTotal(_token common.Address) (*types.Transaction, error) {
	return _RepTok.Contract.UpdateSiblingMigrationTotal(&_RepTok.TransactOpts, _token)
}

// RepTokApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RepTok contract.
type RepTokApprovalIterator struct {
	Event *RepTokApproval // Event containing the contract specifics and raw log

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
func (it *RepTokApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepTokApproval)
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
		it.Event = new(RepTokApproval)
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
func (it *RepTokApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepTokApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepTokApproval represents a Approval event raised by the RepTok contract.
type RepTokApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RepTok *RepTokFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RepTokApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RepTok.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RepTokApprovalIterator{contract: _RepTok.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RepTok *RepTokFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RepTokApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RepTok.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepTokApproval)
				if err := _RepTok.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_RepTok *RepTokFilterer) ParseApproval(log types.Log) (*RepTokApproval, error) {
	event := new(RepTokApproval)
	if err := _RepTok.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RepTokFundedAccountIterator is returned from FilterFundedAccount and is used to iterate over the raw logs and unpacked data for FundedAccount events raised by the RepTok contract.
type RepTokFundedAccountIterator struct {
	Event *RepTokFundedAccount // Event containing the contract specifics and raw log

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
func (it *RepTokFundedAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepTokFundedAccount)
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
		it.Event = new(RepTokFundedAccount)
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
func (it *RepTokFundedAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepTokFundedAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepTokFundedAccount represents a FundedAccount event raised by the RepTok contract.
type RepTokFundedAccount struct {
	Universe   common.Address
	Sender     common.Address
	RepBalance *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFundedAccount is a free log retrieval operation binding the contract event 0xbf88b5a3158512dc8ad44dadf221f46eb1df762a2e7fd1de86aab34b0af455cf.
//
// Solidity: event FundedAccount(address indexed _universe, address indexed _sender, uint256 _repBalance, uint256 _timestamp)
func (_RepTok *RepTokFilterer) FilterFundedAccount(opts *bind.FilterOpts, _universe []common.Address, _sender []common.Address) (*RepTokFundedAccountIterator, error) {

	var _universeRule []interface{}
	for _, _universeItem := range _universe {
		_universeRule = append(_universeRule, _universeItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _RepTok.contract.FilterLogs(opts, "FundedAccount", _universeRule, _senderRule)
	if err != nil {
		return nil, err
	}
	return &RepTokFundedAccountIterator{contract: _RepTok.contract, event: "FundedAccount", logs: logs, sub: sub}, nil
}

// WatchFundedAccount is a free log subscription operation binding the contract event 0xbf88b5a3158512dc8ad44dadf221f46eb1df762a2e7fd1de86aab34b0af455cf.
//
// Solidity: event FundedAccount(address indexed _universe, address indexed _sender, uint256 _repBalance, uint256 _timestamp)
func (_RepTok *RepTokFilterer) WatchFundedAccount(opts *bind.WatchOpts, sink chan<- *RepTokFundedAccount, _universe []common.Address, _sender []common.Address) (event.Subscription, error) {

	var _universeRule []interface{}
	for _, _universeItem := range _universe {
		_universeRule = append(_universeRule, _universeItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _RepTok.contract.WatchLogs(opts, "FundedAccount", _universeRule, _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepTokFundedAccount)
				if err := _RepTok.contract.UnpackLog(event, "FundedAccount", log); err != nil {
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

// ParseFundedAccount is a log parse operation binding the contract event 0xbf88b5a3158512dc8ad44dadf221f46eb1df762a2e7fd1de86aab34b0af455cf.
//
// Solidity: event FundedAccount(address indexed _universe, address indexed _sender, uint256 _repBalance, uint256 _timestamp)
func (_RepTok *RepTokFilterer) ParseFundedAccount(log types.Log) (*RepTokFundedAccount, error) {
	event := new(RepTokFundedAccount)
	if err := _RepTok.contract.UnpackLog(event, "FundedAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RepTokTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RepTok contract.
type RepTokTransferIterator struct {
	Event *RepTokTransfer // Event containing the contract specifics and raw log

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
func (it *RepTokTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepTokTransfer)
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
		it.Event = new(RepTokTransfer)
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
func (it *RepTokTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepTokTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepTokTransfer represents a Transfer event raised by the RepTok contract.
type RepTokTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RepTok *RepTokFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RepTokTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RepTok.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RepTokTransferIterator{contract: _RepTok.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RepTok *RepTokFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RepTokTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RepTok.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepTokTransfer)
				if err := _RepTok.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_RepTok *RepTokFilterer) ParseTransfer(log types.Log) (*RepTokTransfer, error) {
	event := new(RepTokTransfer)
	if err := _RepTok.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
