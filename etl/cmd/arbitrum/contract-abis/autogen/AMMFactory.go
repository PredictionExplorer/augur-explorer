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

// AMMFactoryABI is the input ABI used to generate the binding from.
const AMMFactoryABI = "[{\"inputs\":[{\"internalType\":\"contractBFactory\",\"name\":\"_bFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_hatchery\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_turboId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractITurboHatchery\",\"name\":\"_hatchery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_turboId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_collateralIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minLPTokensOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lpTokenRecipient\",\"type\":\"address\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bFactory\",\"outputs\":[{\"internalType\":\"contractBFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITurboHatchery\",\"name\":\"_hatchery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_turboId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_collateralIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minTokensOut\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITurboHatchery\",\"name\":\"_hatchery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_turboId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_weights\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_lpTokenRecipient\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITurboHatchery\",\"name\":\"_hatchery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_turboId\",\"type\":\"uint256\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITurboHatchery\",\"name\":\"_hatchery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_turboId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_lpTokensPerOutcome\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_minCollateralOut\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITurboHatchery\",\"name\":\"_hatchery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_turboId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_swaps\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_minCollateralOut\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITurboHatchery\",\"name\":\"_hatchery\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_turboId\",\"type\":\"uint256\"}],\"name\":\"tokenRatios\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AMMFactory is an auto generated Go binding around an Ethereum contract.
type AMMFactory struct {
	AMMFactoryCaller     // Read-only binding to the contract
	AMMFactoryTransactor // Write-only binding to the contract
	AMMFactoryFilterer   // Log filterer for contract events
}

// AMMFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AMMFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMMFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AMMFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMMFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AMMFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMMFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AMMFactorySession struct {
	Contract     *AMMFactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMMFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AMMFactoryCallerSession struct {
	Contract *AMMFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AMMFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AMMFactoryTransactorSession struct {
	Contract     *AMMFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AMMFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AMMFactoryRaw struct {
	Contract *AMMFactory // Generic contract binding to access the raw methods on
}

// AMMFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AMMFactoryCallerRaw struct {
	Contract *AMMFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AMMFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AMMFactoryTransactorRaw struct {
	Contract *AMMFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAMMFactory creates a new instance of AMMFactory, bound to a specific deployed contract.
func NewAMMFactory(address common.Address, backend bind.ContractBackend) (*AMMFactory, error) {
	contract, err := bindAMMFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AMMFactory{AMMFactoryCaller: AMMFactoryCaller{contract: contract}, AMMFactoryTransactor: AMMFactoryTransactor{contract: contract}, AMMFactoryFilterer: AMMFactoryFilterer{contract: contract}}, nil
}

// NewAMMFactoryCaller creates a new read-only instance of AMMFactory, bound to a specific deployed contract.
func NewAMMFactoryCaller(address common.Address, caller bind.ContractCaller) (*AMMFactoryCaller, error) {
	contract, err := bindAMMFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AMMFactoryCaller{contract: contract}, nil
}

// NewAMMFactoryTransactor creates a new write-only instance of AMMFactory, bound to a specific deployed contract.
func NewAMMFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AMMFactoryTransactor, error) {
	contract, err := bindAMMFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AMMFactoryTransactor{contract: contract}, nil
}

// NewAMMFactoryFilterer creates a new log filterer instance of AMMFactory, bound to a specific deployed contract.
func NewAMMFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AMMFactoryFilterer, error) {
	contract, err := bindAMMFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AMMFactoryFilterer{contract: contract}, nil
}

// bindAMMFactory binds a generic wrapper to an already deployed contract.
func bindAMMFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AMMFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMMFactory *AMMFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AMMFactory.Contract.AMMFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMMFactory *AMMFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMMFactory.Contract.AMMFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMMFactory *AMMFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMMFactory.Contract.AMMFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMMFactory *AMMFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AMMFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMMFactory *AMMFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMMFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMMFactory *AMMFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMMFactory.Contract.contract.Transact(opts, method, params...)
}

// BFactory is a free data retrieval call binding the contract method 0x0a165940.
//
// Solidity: function bFactory() view returns(address)
func (_AMMFactory *AMMFactoryCaller) BFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "bFactory")
	return *ret0, err
}

// BFactory is a free data retrieval call binding the contract method 0x0a165940.
//
// Solidity: function bFactory() view returns(address)
func (_AMMFactory *AMMFactorySession) BFactory() (common.Address, error) {
	return _AMMFactory.Contract.BFactory(&_AMMFactory.CallOpts)
}

// BFactory is a free data retrieval call binding the contract method 0x0a165940.
//
// Solidity: function bFactory() view returns(address)
func (_AMMFactory *AMMFactoryCallerSession) BFactory() (common.Address, error) {
	return _AMMFactory.Contract.BFactory(&_AMMFactory.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0x8f38a555.
//
// Solidity: function pools(address , uint256 ) view returns(address)
func (_AMMFactory *AMMFactoryCaller) Pools(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "pools", arg0, arg1)
	return *ret0, err
}

// Pools is a free data retrieval call binding the contract method 0x8f38a555.
//
// Solidity: function pools(address , uint256 ) view returns(address)
func (_AMMFactory *AMMFactorySession) Pools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _AMMFactory.Contract.Pools(&_AMMFactory.CallOpts, arg0, arg1)
}

// Pools is a free data retrieval call binding the contract method 0x8f38a555.
//
// Solidity: function pools(address , uint256 ) view returns(address)
func (_AMMFactory *AMMFactoryCallerSession) Pools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _AMMFactory.Contract.Pools(&_AMMFactory.CallOpts, arg0, arg1)
}

// Prices is a free data retrieval call binding the contract method 0xbbc492c0.
//
// Solidity: function prices(address _hatchery, uint256 _turboId) view returns(uint256[])
func (_AMMFactory *AMMFactoryCaller) Prices(opts *bind.CallOpts, _hatchery common.Address, _turboId *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "prices", _hatchery, _turboId)
	return *ret0, err
}

// Prices is a free data retrieval call binding the contract method 0xbbc492c0.
//
// Solidity: function prices(address _hatchery, uint256 _turboId) view returns(uint256[])
func (_AMMFactory *AMMFactorySession) Prices(_hatchery common.Address, _turboId *big.Int) ([]*big.Int, error) {
	return _AMMFactory.Contract.Prices(&_AMMFactory.CallOpts, _hatchery, _turboId)
}

// Prices is a free data retrieval call binding the contract method 0xbbc492c0.
//
// Solidity: function prices(address _hatchery, uint256 _turboId) view returns(uint256[])
func (_AMMFactory *AMMFactoryCallerSession) Prices(_hatchery common.Address, _turboId *big.Int) ([]*big.Int, error) {
	return _AMMFactory.Contract.Prices(&_AMMFactory.CallOpts, _hatchery, _turboId)
}

// TokenRatios is a free data retrieval call binding the contract method 0xc7b4b6dd.
//
// Solidity: function tokenRatios(address _hatchery, uint256 _turboId) view returns(uint256[])
func (_AMMFactory *AMMFactoryCaller) TokenRatios(opts *bind.CallOpts, _hatchery common.Address, _turboId *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "tokenRatios", _hatchery, _turboId)
	return *ret0, err
}

// TokenRatios is a free data retrieval call binding the contract method 0xc7b4b6dd.
//
// Solidity: function tokenRatios(address _hatchery, uint256 _turboId) view returns(uint256[])
func (_AMMFactory *AMMFactorySession) TokenRatios(_hatchery common.Address, _turboId *big.Int) ([]*big.Int, error) {
	return _AMMFactory.Contract.TokenRatios(&_AMMFactory.CallOpts, _hatchery, _turboId)
}

// TokenRatios is a free data retrieval call binding the contract method 0xc7b4b6dd.
//
// Solidity: function tokenRatios(address _hatchery, uint256 _turboId) view returns(uint256[])
func (_AMMFactory *AMMFactoryCallerSession) TokenRatios(_hatchery common.Address, _turboId *big.Int) ([]*big.Int, error) {
	return _AMMFactory.Contract.TokenRatios(&_AMMFactory.CallOpts, _hatchery, _turboId)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x45fa6783.
//
// Solidity: function addLiquidity(address _hatchery, uint256 _turboId, uint256 _collateralIn, uint256 _minLPTokensOut, address _lpTokenRecipient) returns(uint256)
func (_AMMFactory *AMMFactoryTransactor) AddLiquidity(opts *bind.TransactOpts, _hatchery common.Address, _turboId *big.Int, _collateralIn *big.Int, _minLPTokensOut *big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.contract.Transact(opts, "addLiquidity", _hatchery, _turboId, _collateralIn, _minLPTokensOut, _lpTokenRecipient)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x45fa6783.
//
// Solidity: function addLiquidity(address _hatchery, uint256 _turboId, uint256 _collateralIn, uint256 _minLPTokensOut, address _lpTokenRecipient) returns(uint256)
func (_AMMFactory *AMMFactorySession) AddLiquidity(_hatchery common.Address, _turboId *big.Int, _collateralIn *big.Int, _minLPTokensOut *big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.Contract.AddLiquidity(&_AMMFactory.TransactOpts, _hatchery, _turboId, _collateralIn, _minLPTokensOut, _lpTokenRecipient)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x45fa6783.
//
// Solidity: function addLiquidity(address _hatchery, uint256 _turboId, uint256 _collateralIn, uint256 _minLPTokensOut, address _lpTokenRecipient) returns(uint256)
func (_AMMFactory *AMMFactoryTransactorSession) AddLiquidity(_hatchery common.Address, _turboId *big.Int, _collateralIn *big.Int, _minLPTokensOut *big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.Contract.AddLiquidity(&_AMMFactory.TransactOpts, _hatchery, _turboId, _collateralIn, _minLPTokensOut, _lpTokenRecipient)
}

// Buy is a paid mutator transaction binding the contract method 0x72b60c30.
//
// Solidity: function buy(address _hatchery, uint256 _turboId, uint256 _outcome, uint256 _collateralIn, uint256 _minTokensOut) returns(uint256)
func (_AMMFactory *AMMFactoryTransactor) Buy(opts *bind.TransactOpts, _hatchery common.Address, _turboId *big.Int, _outcome *big.Int, _collateralIn *big.Int, _minTokensOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.contract.Transact(opts, "buy", _hatchery, _turboId, _outcome, _collateralIn, _minTokensOut)
}

// Buy is a paid mutator transaction binding the contract method 0x72b60c30.
//
// Solidity: function buy(address _hatchery, uint256 _turboId, uint256 _outcome, uint256 _collateralIn, uint256 _minTokensOut) returns(uint256)
func (_AMMFactory *AMMFactorySession) Buy(_hatchery common.Address, _turboId *big.Int, _outcome *big.Int, _collateralIn *big.Int, _minTokensOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.Contract.Buy(&_AMMFactory.TransactOpts, _hatchery, _turboId, _outcome, _collateralIn, _minTokensOut)
}

// Buy is a paid mutator transaction binding the contract method 0x72b60c30.
//
// Solidity: function buy(address _hatchery, uint256 _turboId, uint256 _outcome, uint256 _collateralIn, uint256 _minTokensOut) returns(uint256)
func (_AMMFactory *AMMFactoryTransactorSession) Buy(_hatchery common.Address, _turboId *big.Int, _outcome *big.Int, _collateralIn *big.Int, _minTokensOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.Contract.Buy(&_AMMFactory.TransactOpts, _hatchery, _turboId, _outcome, _collateralIn, _minTokensOut)
}

// CreatePool is a paid mutator transaction binding the contract method 0xfd212cf6.
//
// Solidity: function createPool(address _hatchery, uint256 _turboId, uint256 _initialLiquidity, uint256[] _weights, address _lpTokenRecipient) returns(address)
func (_AMMFactory *AMMFactoryTransactor) CreatePool(opts *bind.TransactOpts, _hatchery common.Address, _turboId *big.Int, _initialLiquidity *big.Int, _weights []*big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.contract.Transact(opts, "createPool", _hatchery, _turboId, _initialLiquidity, _weights, _lpTokenRecipient)
}

// CreatePool is a paid mutator transaction binding the contract method 0xfd212cf6.
//
// Solidity: function createPool(address _hatchery, uint256 _turboId, uint256 _initialLiquidity, uint256[] _weights, address _lpTokenRecipient) returns(address)
func (_AMMFactory *AMMFactorySession) CreatePool(_hatchery common.Address, _turboId *big.Int, _initialLiquidity *big.Int, _weights []*big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.Contract.CreatePool(&_AMMFactory.TransactOpts, _hatchery, _turboId, _initialLiquidity, _weights, _lpTokenRecipient)
}

// CreatePool is a paid mutator transaction binding the contract method 0xfd212cf6.
//
// Solidity: function createPool(address _hatchery, uint256 _turboId, uint256 _initialLiquidity, uint256[] _weights, address _lpTokenRecipient) returns(address)
func (_AMMFactory *AMMFactoryTransactorSession) CreatePool(_hatchery common.Address, _turboId *big.Int, _initialLiquidity *big.Int, _weights []*big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.Contract.CreatePool(&_AMMFactory.TransactOpts, _hatchery, _turboId, _initialLiquidity, _weights, _lpTokenRecipient)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd3b7634f.
//
// Solidity: function removeLiquidity(address _hatchery, uint256 _turboId, uint256[] _lpTokensPerOutcome, uint256 _minCollateralOut) returns(uint256)
func (_AMMFactory *AMMFactoryTransactor) RemoveLiquidity(opts *bind.TransactOpts, _hatchery common.Address, _turboId *big.Int, _lpTokensPerOutcome []*big.Int, _minCollateralOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.contract.Transact(opts, "removeLiquidity", _hatchery, _turboId, _lpTokensPerOutcome, _minCollateralOut)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd3b7634f.
//
// Solidity: function removeLiquidity(address _hatchery, uint256 _turboId, uint256[] _lpTokensPerOutcome, uint256 _minCollateralOut) returns(uint256)
func (_AMMFactory *AMMFactorySession) RemoveLiquidity(_hatchery common.Address, _turboId *big.Int, _lpTokensPerOutcome []*big.Int, _minCollateralOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.Contract.RemoveLiquidity(&_AMMFactory.TransactOpts, _hatchery, _turboId, _lpTokensPerOutcome, _minCollateralOut)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd3b7634f.
//
// Solidity: function removeLiquidity(address _hatchery, uint256 _turboId, uint256[] _lpTokensPerOutcome, uint256 _minCollateralOut) returns(uint256)
func (_AMMFactory *AMMFactoryTransactorSession) RemoveLiquidity(_hatchery common.Address, _turboId *big.Int, _lpTokensPerOutcome []*big.Int, _minCollateralOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.Contract.RemoveLiquidity(&_AMMFactory.TransactOpts, _hatchery, _turboId, _lpTokensPerOutcome, _minCollateralOut)
}

// Sell is a paid mutator transaction binding the contract method 0xdcd80dcf.
//
// Solidity: function sell(address _hatchery, uint256 _turboId, uint256 _outcome, uint256[] _swaps, uint256 _minCollateralOut) returns(uint256)
func (_AMMFactory *AMMFactoryTransactor) Sell(opts *bind.TransactOpts, _hatchery common.Address, _turboId *big.Int, _outcome *big.Int, _swaps []*big.Int, _minCollateralOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.contract.Transact(opts, "sell", _hatchery, _turboId, _outcome, _swaps, _minCollateralOut)
}

// Sell is a paid mutator transaction binding the contract method 0xdcd80dcf.
//
// Solidity: function sell(address _hatchery, uint256 _turboId, uint256 _outcome, uint256[] _swaps, uint256 _minCollateralOut) returns(uint256)
func (_AMMFactory *AMMFactorySession) Sell(_hatchery common.Address, _turboId *big.Int, _outcome *big.Int, _swaps []*big.Int, _minCollateralOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.Contract.Sell(&_AMMFactory.TransactOpts, _hatchery, _turboId, _outcome, _swaps, _minCollateralOut)
}

// Sell is a paid mutator transaction binding the contract method 0xdcd80dcf.
//
// Solidity: function sell(address _hatchery, uint256 _turboId, uint256 _outcome, uint256[] _swaps, uint256 _minCollateralOut) returns(uint256)
func (_AMMFactory *AMMFactoryTransactorSession) Sell(_hatchery common.Address, _turboId *big.Int, _outcome *big.Int, _swaps []*big.Int, _minCollateralOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.Contract.Sell(&_AMMFactory.TransactOpts, _hatchery, _turboId, _outcome, _swaps, _minCollateralOut)
}

// AMMFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the AMMFactory contract.
type AMMFactoryPoolCreatedIterator struct {
	Event *AMMFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *AMMFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMMFactoryPoolCreated)
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
		it.Event = new(AMMFactoryPoolCreated)
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
func (it *AMMFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMMFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMMFactoryPoolCreated represents a PoolCreated event raised by the AMMFactory contract.
type AMMFactoryPoolCreated struct {
	Pool     common.Address
	Hatchery common.Address
	TurboId  *big.Int
	Creator  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x0537d3e5d88810bbfb16874b28bc0f95856d7bb24c8f29511fe463c5b1d27c6b.
//
// Solidity: event PoolCreated(address _pool, address indexed _hatchery, uint256 indexed _turboId, address indexed _creator)
func (_AMMFactory *AMMFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, _hatchery []common.Address, _turboId []*big.Int, _creator []common.Address) (*AMMFactoryPoolCreatedIterator, error) {

	var _hatcheryRule []interface{}
	for _, _hatcheryItem := range _hatchery {
		_hatcheryRule = append(_hatcheryRule, _hatcheryItem)
	}
	var _turboIdRule []interface{}
	for _, _turboIdItem := range _turboId {
		_turboIdRule = append(_turboIdRule, _turboIdItem)
	}
	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _AMMFactory.contract.FilterLogs(opts, "PoolCreated", _hatcheryRule, _turboIdRule, _creatorRule)
	if err != nil {
		return nil, err
	}
	return &AMMFactoryPoolCreatedIterator{contract: _AMMFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x0537d3e5d88810bbfb16874b28bc0f95856d7bb24c8f29511fe463c5b1d27c6b.
//
// Solidity: event PoolCreated(address _pool, address indexed _hatchery, uint256 indexed _turboId, address indexed _creator)
func (_AMMFactory *AMMFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *AMMFactoryPoolCreated, _hatchery []common.Address, _turboId []*big.Int, _creator []common.Address) (event.Subscription, error) {

	var _hatcheryRule []interface{}
	for _, _hatcheryItem := range _hatchery {
		_hatcheryRule = append(_hatcheryRule, _hatcheryItem)
	}
	var _turboIdRule []interface{}
	for _, _turboIdItem := range _turboId {
		_turboIdRule = append(_turboIdRule, _turboIdItem)
	}
	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _AMMFactory.contract.WatchLogs(opts, "PoolCreated", _hatcheryRule, _turboIdRule, _creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMMFactoryPoolCreated)
				if err := _AMMFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x0537d3e5d88810bbfb16874b28bc0f95856d7bb24c8f29511fe463c5b1d27c6b.
//
// Solidity: event PoolCreated(address _pool, address indexed _hatchery, uint256 indexed _turboId, address indexed _creator)
func (_AMMFactory *AMMFactoryFilterer) ParsePoolCreated(log types.Log) (*AMMFactoryPoolCreated, error) {
	event := new(AMMFactoryPoolCreated)
	if err := _AMMFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}
