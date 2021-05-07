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
const AMMFactoryABI = "[{\"inputs\":[{\"internalType\":\"contractBFactory\",\"name\":\"_bFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"collateral\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"lpTokens\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"sharesReturned\",\"type\":\"uint256[]\"}],\"name\":\"LiquidityChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lpTokenRecipient\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outcome\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"collateral\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"shares\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"inOutRatio\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SharesSwapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAbstractMarketFactory\",\"name\":\"_marketFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_collateralIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minLPTokensOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lpTokenRecipient\",\"type\":\"address\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bFactory\",\"outputs\":[{\"internalType\":\"contractBFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAbstractMarketFactory\",\"name\":\"_marketFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_collateralIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minTokensOut\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAbstractMarketFactory\",\"name\":\"_marketFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_weights\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_lpTokenRecipient\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAbstractMarketFactory\",\"name\":\"_marketFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAbstractMarketFactory\",\"name\":\"_marketFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"}],\"name\":\"getPoolBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAbstractMarketFactory\",\"name\":\"_marketFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"getPoolTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAbstractMarketFactory\",\"name\":\"_marketFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"}],\"name\":\"getPoolWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAbstractMarketFactory\",\"name\":\"_marketFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"}],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"contractBPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAbstractMarketFactory\",\"name\":\"_marketFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lpTokensIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minCollateralOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_collateralRecipient\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_collateralOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAbstractMarketFactory\",\"name\":\"_marketFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_shareTokensIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_setsOut\",\"type\":\"uint256\"}],\"name\":\"sellForCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAbstractMarketFactory\",\"name\":\"_marketFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"}],\"name\":\"tokenRatios\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "BONE")
	return *ret0, err
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_AMMFactory *AMMFactorySession) BONE() (*big.Int, error) {
	return _AMMFactory.Contract.BONE(&_AMMFactory.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) BONE() (*big.Int, error) {
	return _AMMFactory.Contract.BONE(&_AMMFactory.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "BPOW_PRECISION")
	return *ret0, err
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_AMMFactory *AMMFactorySession) BPOWPRECISION() (*big.Int, error) {
	return _AMMFactory.Contract.BPOWPRECISION(&_AMMFactory.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _AMMFactory.Contract.BPOWPRECISION(&_AMMFactory.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "EXIT_FEE")
	return *ret0, err
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_AMMFactory *AMMFactorySession) EXITFEE() (*big.Int, error) {
	return _AMMFactory.Contract.EXITFEE(&_AMMFactory.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) EXITFEE() (*big.Int, error) {
	return _AMMFactory.Contract.EXITFEE(&_AMMFactory.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "INIT_POOL_SUPPLY")
	return *ret0, err
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_AMMFactory *AMMFactorySession) INITPOOLSUPPLY() (*big.Int, error) {
	return _AMMFactory.Contract.INITPOOLSUPPLY(&_AMMFactory.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _AMMFactory.Contract.INITPOOLSUPPLY(&_AMMFactory.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MAX_BOUND_TOKENS")
	return *ret0, err
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _AMMFactory.Contract.MAXBOUNDTOKENS(&_AMMFactory.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _AMMFactory.Contract.MAXBOUNDTOKENS(&_AMMFactory.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MAX_BPOW_BASE")
	return *ret0, err
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MAXBPOWBASE() (*big.Int, error) {
	return _AMMFactory.Contract.MAXBPOWBASE(&_AMMFactory.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _AMMFactory.Contract.MAXBPOWBASE(&_AMMFactory.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MAX_FEE")
	return *ret0, err
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MAXFEE() (*big.Int, error) {
	return _AMMFactory.Contract.MAXFEE(&_AMMFactory.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MAXFEE() (*big.Int, error) {
	return _AMMFactory.Contract.MAXFEE(&_AMMFactory.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MAX_IN_RATIO")
	return *ret0, err
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MAXINRATIO() (*big.Int, error) {
	return _AMMFactory.Contract.MAXINRATIO(&_AMMFactory.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MAXINRATIO() (*big.Int, error) {
	return _AMMFactory.Contract.MAXINRATIO(&_AMMFactory.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MAX_OUT_RATIO")
	return *ret0, err
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MAXOUTRATIO() (*big.Int, error) {
	return _AMMFactory.Contract.MAXOUTRATIO(&_AMMFactory.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _AMMFactory.Contract.MAXOUTRATIO(&_AMMFactory.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MAX_TOTAL_WEIGHT")
	return *ret0, err
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _AMMFactory.Contract.MAXTOTALWEIGHT(&_AMMFactory.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _AMMFactory.Contract.MAXTOTALWEIGHT(&_AMMFactory.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MAX_WEIGHT")
	return *ret0, err
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MAXWEIGHT() (*big.Int, error) {
	return _AMMFactory.Contract.MAXWEIGHT(&_AMMFactory.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _AMMFactory.Contract.MAXWEIGHT(&_AMMFactory.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MIN_BALANCE")
	return *ret0, err
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MINBALANCE() (*big.Int, error) {
	return _AMMFactory.Contract.MINBALANCE(&_AMMFactory.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MINBALANCE() (*big.Int, error) {
	return _AMMFactory.Contract.MINBALANCE(&_AMMFactory.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MIN_BOUND_TOKENS")
	return *ret0, err
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MINBOUNDTOKENS() (*big.Int, error) {
	return _AMMFactory.Contract.MINBOUNDTOKENS(&_AMMFactory.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _AMMFactory.Contract.MINBOUNDTOKENS(&_AMMFactory.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MIN_BPOW_BASE")
	return *ret0, err
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MINBPOWBASE() (*big.Int, error) {
	return _AMMFactory.Contract.MINBPOWBASE(&_AMMFactory.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _AMMFactory.Contract.MINBPOWBASE(&_AMMFactory.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MIN_FEE")
	return *ret0, err
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MINFEE() (*big.Int, error) {
	return _AMMFactory.Contract.MINFEE(&_AMMFactory.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MINFEE() (*big.Int, error) {
	return _AMMFactory.Contract.MINFEE(&_AMMFactory.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "MIN_WEIGHT")
	return *ret0, err
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_AMMFactory *AMMFactorySession) MINWEIGHT() (*big.Int, error) {
	return _AMMFactory.Contract.MINWEIGHT(&_AMMFactory.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) MINWEIGHT() (*big.Int, error) {
	return _AMMFactory.Contract.MINWEIGHT(&_AMMFactory.CallOpts)
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

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() pure returns(bytes32)
func (_AMMFactory *AMMFactoryCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "getColor")
	return *ret0, err
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() pure returns(bytes32)
func (_AMMFactory *AMMFactorySession) GetColor() ([32]byte, error) {
	return _AMMFactory.Contract.GetColor(&_AMMFactory.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() pure returns(bytes32)
func (_AMMFactory *AMMFactoryCallerSession) GetColor() ([32]byte, error) {
	return _AMMFactory.Contract.GetColor(&_AMMFactory.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _marketFactory, uint256 _marketId) view returns(address)
func (_AMMFactory *AMMFactoryCaller) GetPool(opts *bind.CallOpts, _marketFactory common.Address, _marketId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "getPool", _marketFactory, _marketId)
	return *ret0, err
}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _marketFactory, uint256 _marketId) view returns(address)
func (_AMMFactory *AMMFactorySession) GetPool(_marketFactory common.Address, _marketId *big.Int) (common.Address, error) {
	return _AMMFactory.Contract.GetPool(&_AMMFactory.CallOpts, _marketFactory, _marketId)
}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _marketFactory, uint256 _marketId) view returns(address)
func (_AMMFactory *AMMFactoryCallerSession) GetPool(_marketFactory common.Address, _marketId *big.Int) (common.Address, error) {
	return _AMMFactory.Contract.GetPool(&_AMMFactory.CallOpts, _marketFactory, _marketId)
}

// GetPoolBalances is a free data retrieval call binding the contract method 0xd2364bf3.
//
// Solidity: function getPoolBalances(address _marketFactory, uint256 _marketId) view returns(uint256[])
func (_AMMFactory *AMMFactoryCaller) GetPoolBalances(opts *bind.CallOpts, _marketFactory common.Address, _marketId *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "getPoolBalances", _marketFactory, _marketId)
	return *ret0, err
}

// GetPoolBalances is a free data retrieval call binding the contract method 0xd2364bf3.
//
// Solidity: function getPoolBalances(address _marketFactory, uint256 _marketId) view returns(uint256[])
func (_AMMFactory *AMMFactorySession) GetPoolBalances(_marketFactory common.Address, _marketId *big.Int) ([]*big.Int, error) {
	return _AMMFactory.Contract.GetPoolBalances(&_AMMFactory.CallOpts, _marketFactory, _marketId)
}

// GetPoolBalances is a free data retrieval call binding the contract method 0xd2364bf3.
//
// Solidity: function getPoolBalances(address _marketFactory, uint256 _marketId) view returns(uint256[])
func (_AMMFactory *AMMFactoryCallerSession) GetPoolBalances(_marketFactory common.Address, _marketId *big.Int) ([]*big.Int, error) {
	return _AMMFactory.Contract.GetPoolBalances(&_AMMFactory.CallOpts, _marketFactory, _marketId)
}

// GetPoolTokenBalance is a free data retrieval call binding the contract method 0xa0147aa1.
//
// Solidity: function getPoolTokenBalance(address _marketFactory, uint256 _marketId, address whom) view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) GetPoolTokenBalance(opts *bind.CallOpts, _marketFactory common.Address, _marketId *big.Int, whom common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "getPoolTokenBalance", _marketFactory, _marketId, whom)
	return *ret0, err
}

// GetPoolTokenBalance is a free data retrieval call binding the contract method 0xa0147aa1.
//
// Solidity: function getPoolTokenBalance(address _marketFactory, uint256 _marketId, address whom) view returns(uint256)
func (_AMMFactory *AMMFactorySession) GetPoolTokenBalance(_marketFactory common.Address, _marketId *big.Int, whom common.Address) (*big.Int, error) {
	return _AMMFactory.Contract.GetPoolTokenBalance(&_AMMFactory.CallOpts, _marketFactory, _marketId, whom)
}

// GetPoolTokenBalance is a free data retrieval call binding the contract method 0xa0147aa1.
//
// Solidity: function getPoolTokenBalance(address _marketFactory, uint256 _marketId, address whom) view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) GetPoolTokenBalance(_marketFactory common.Address, _marketId *big.Int, whom common.Address) (*big.Int, error) {
	return _AMMFactory.Contract.GetPoolTokenBalance(&_AMMFactory.CallOpts, _marketFactory, _marketId, whom)
}

// GetPoolWeights is a free data retrieval call binding the contract method 0xd055da71.
//
// Solidity: function getPoolWeights(address _marketFactory, uint256 _marketId) view returns(uint256[])
func (_AMMFactory *AMMFactoryCaller) GetPoolWeights(opts *bind.CallOpts, _marketFactory common.Address, _marketId *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "getPoolWeights", _marketFactory, _marketId)
	return *ret0, err
}

// GetPoolWeights is a free data retrieval call binding the contract method 0xd055da71.
//
// Solidity: function getPoolWeights(address _marketFactory, uint256 _marketId) view returns(uint256[])
func (_AMMFactory *AMMFactorySession) GetPoolWeights(_marketFactory common.Address, _marketId *big.Int) ([]*big.Int, error) {
	return _AMMFactory.Contract.GetPoolWeights(&_AMMFactory.CallOpts, _marketFactory, _marketId)
}

// GetPoolWeights is a free data retrieval call binding the contract method 0xd055da71.
//
// Solidity: function getPoolWeights(address _marketFactory, uint256 _marketId) view returns(uint256[])
func (_AMMFactory *AMMFactoryCallerSession) GetPoolWeights(_marketFactory common.Address, _marketId *big.Int) ([]*big.Int, error) {
	return _AMMFactory.Contract.GetPoolWeights(&_AMMFactory.CallOpts, _marketFactory, _marketId)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xfa0de359.
//
// Solidity: function getSwapFee(address _marketFactory, uint256 _marketId) view returns(uint256)
func (_AMMFactory *AMMFactoryCaller) GetSwapFee(opts *bind.CallOpts, _marketFactory common.Address, _marketId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "getSwapFee", _marketFactory, _marketId)
	return *ret0, err
}

// GetSwapFee is a free data retrieval call binding the contract method 0xfa0de359.
//
// Solidity: function getSwapFee(address _marketFactory, uint256 _marketId) view returns(uint256)
func (_AMMFactory *AMMFactorySession) GetSwapFee(_marketFactory common.Address, _marketId *big.Int) (*big.Int, error) {
	return _AMMFactory.Contract.GetSwapFee(&_AMMFactory.CallOpts, _marketFactory, _marketId)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xfa0de359.
//
// Solidity: function getSwapFee(address _marketFactory, uint256 _marketId) view returns(uint256)
func (_AMMFactory *AMMFactoryCallerSession) GetSwapFee(_marketFactory common.Address, _marketId *big.Int) (*big.Int, error) {
	return _AMMFactory.Contract.GetSwapFee(&_AMMFactory.CallOpts, _marketFactory, _marketId)
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

// TokenRatios is a free data retrieval call binding the contract method 0xc7b4b6dd.
//
// Solidity: function tokenRatios(address _marketFactory, uint256 _marketId) view returns(uint256[])
func (_AMMFactory *AMMFactoryCaller) TokenRatios(opts *bind.CallOpts, _marketFactory common.Address, _marketId *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _AMMFactory.contract.Call(opts, out, "tokenRatios", _marketFactory, _marketId)
	return *ret0, err
}

// TokenRatios is a free data retrieval call binding the contract method 0xc7b4b6dd.
//
// Solidity: function tokenRatios(address _marketFactory, uint256 _marketId) view returns(uint256[])
func (_AMMFactory *AMMFactorySession) TokenRatios(_marketFactory common.Address, _marketId *big.Int) ([]*big.Int, error) {
	return _AMMFactory.Contract.TokenRatios(&_AMMFactory.CallOpts, _marketFactory, _marketId)
}

// TokenRatios is a free data retrieval call binding the contract method 0xc7b4b6dd.
//
// Solidity: function tokenRatios(address _marketFactory, uint256 _marketId) view returns(uint256[])
func (_AMMFactory *AMMFactoryCallerSession) TokenRatios(_marketFactory common.Address, _marketId *big.Int) ([]*big.Int, error) {
	return _AMMFactory.Contract.TokenRatios(&_AMMFactory.CallOpts, _marketFactory, _marketId)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x45fa6783.
//
// Solidity: function addLiquidity(address _marketFactory, uint256 _marketId, uint256 _collateralIn, uint256 _minLPTokensOut, address _lpTokenRecipient) returns(uint256 _poolAmountOut, uint256[] _balances)
func (_AMMFactory *AMMFactoryTransactor) AddLiquidity(opts *bind.TransactOpts, _marketFactory common.Address, _marketId *big.Int, _collateralIn *big.Int, _minLPTokensOut *big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.contract.Transact(opts, "addLiquidity", _marketFactory, _marketId, _collateralIn, _minLPTokensOut, _lpTokenRecipient)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x45fa6783.
//
// Solidity: function addLiquidity(address _marketFactory, uint256 _marketId, uint256 _collateralIn, uint256 _minLPTokensOut, address _lpTokenRecipient) returns(uint256 _poolAmountOut, uint256[] _balances)
func (_AMMFactory *AMMFactorySession) AddLiquidity(_marketFactory common.Address, _marketId *big.Int, _collateralIn *big.Int, _minLPTokensOut *big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.Contract.AddLiquidity(&_AMMFactory.TransactOpts, _marketFactory, _marketId, _collateralIn, _minLPTokensOut, _lpTokenRecipient)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x45fa6783.
//
// Solidity: function addLiquidity(address _marketFactory, uint256 _marketId, uint256 _collateralIn, uint256 _minLPTokensOut, address _lpTokenRecipient) returns(uint256 _poolAmountOut, uint256[] _balances)
func (_AMMFactory *AMMFactoryTransactorSession) AddLiquidity(_marketFactory common.Address, _marketId *big.Int, _collateralIn *big.Int, _minLPTokensOut *big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.Contract.AddLiquidity(&_AMMFactory.TransactOpts, _marketFactory, _marketId, _collateralIn, _minLPTokensOut, _lpTokenRecipient)
}

// Buy is a paid mutator transaction binding the contract method 0x72b60c30.
//
// Solidity: function buy(address _marketFactory, uint256 _marketId, uint256 _outcome, uint256 _collateralIn, uint256 _minTokensOut) returns(uint256)
func (_AMMFactory *AMMFactoryTransactor) Buy(opts *bind.TransactOpts, _marketFactory common.Address, _marketId *big.Int, _outcome *big.Int, _collateralIn *big.Int, _minTokensOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.contract.Transact(opts, "buy", _marketFactory, _marketId, _outcome, _collateralIn, _minTokensOut)
}

// Buy is a paid mutator transaction binding the contract method 0x72b60c30.
//
// Solidity: function buy(address _marketFactory, uint256 _marketId, uint256 _outcome, uint256 _collateralIn, uint256 _minTokensOut) returns(uint256)
func (_AMMFactory *AMMFactorySession) Buy(_marketFactory common.Address, _marketId *big.Int, _outcome *big.Int, _collateralIn *big.Int, _minTokensOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.Contract.Buy(&_AMMFactory.TransactOpts, _marketFactory, _marketId, _outcome, _collateralIn, _minTokensOut)
}

// Buy is a paid mutator transaction binding the contract method 0x72b60c30.
//
// Solidity: function buy(address _marketFactory, uint256 _marketId, uint256 _outcome, uint256 _collateralIn, uint256 _minTokensOut) returns(uint256)
func (_AMMFactory *AMMFactoryTransactorSession) Buy(_marketFactory common.Address, _marketId *big.Int, _outcome *big.Int, _collateralIn *big.Int, _minTokensOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.Contract.Buy(&_AMMFactory.TransactOpts, _marketFactory, _marketId, _outcome, _collateralIn, _minTokensOut)
}

// CreatePool is a paid mutator transaction binding the contract method 0xfd212cf6.
//
// Solidity: function createPool(address _marketFactory, uint256 _marketId, uint256 _initialLiquidity, uint256[] _weights, address _lpTokenRecipient) returns(uint256)
func (_AMMFactory *AMMFactoryTransactor) CreatePool(opts *bind.TransactOpts, _marketFactory common.Address, _marketId *big.Int, _initialLiquidity *big.Int, _weights []*big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.contract.Transact(opts, "createPool", _marketFactory, _marketId, _initialLiquidity, _weights, _lpTokenRecipient)
}

// CreatePool is a paid mutator transaction binding the contract method 0xfd212cf6.
//
// Solidity: function createPool(address _marketFactory, uint256 _marketId, uint256 _initialLiquidity, uint256[] _weights, address _lpTokenRecipient) returns(uint256)
func (_AMMFactory *AMMFactorySession) CreatePool(_marketFactory common.Address, _marketId *big.Int, _initialLiquidity *big.Int, _weights []*big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.Contract.CreatePool(&_AMMFactory.TransactOpts, _marketFactory, _marketId, _initialLiquidity, _weights, _lpTokenRecipient)
}

// CreatePool is a paid mutator transaction binding the contract method 0xfd212cf6.
//
// Solidity: function createPool(address _marketFactory, uint256 _marketId, uint256 _initialLiquidity, uint256[] _weights, address _lpTokenRecipient) returns(uint256)
func (_AMMFactory *AMMFactoryTransactorSession) CreatePool(_marketFactory common.Address, _marketId *big.Int, _initialLiquidity *big.Int, _weights []*big.Int, _lpTokenRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.Contract.CreatePool(&_AMMFactory.TransactOpts, _marketFactory, _marketId, _initialLiquidity, _weights, _lpTokenRecipient)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x59f842b2.
//
// Solidity: function removeLiquidity(address _marketFactory, uint256 _marketId, uint256 _lpTokensIn, uint256 _minCollateralOut, address _collateralRecipient) returns(uint256 _collateralOut, uint256[] _balances)
func (_AMMFactory *AMMFactoryTransactor) RemoveLiquidity(opts *bind.TransactOpts, _marketFactory common.Address, _marketId *big.Int, _lpTokensIn *big.Int, _minCollateralOut *big.Int, _collateralRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.contract.Transact(opts, "removeLiquidity", _marketFactory, _marketId, _lpTokensIn, _minCollateralOut, _collateralRecipient)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x59f842b2.
//
// Solidity: function removeLiquidity(address _marketFactory, uint256 _marketId, uint256 _lpTokensIn, uint256 _minCollateralOut, address _collateralRecipient) returns(uint256 _collateralOut, uint256[] _balances)
func (_AMMFactory *AMMFactorySession) RemoveLiquidity(_marketFactory common.Address, _marketId *big.Int, _lpTokensIn *big.Int, _minCollateralOut *big.Int, _collateralRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.Contract.RemoveLiquidity(&_AMMFactory.TransactOpts, _marketFactory, _marketId, _lpTokensIn, _minCollateralOut, _collateralRecipient)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x59f842b2.
//
// Solidity: function removeLiquidity(address _marketFactory, uint256 _marketId, uint256 _lpTokensIn, uint256 _minCollateralOut, address _collateralRecipient) returns(uint256 _collateralOut, uint256[] _balances)
func (_AMMFactory *AMMFactoryTransactorSession) RemoveLiquidity(_marketFactory common.Address, _marketId *big.Int, _lpTokensIn *big.Int, _minCollateralOut *big.Int, _collateralRecipient common.Address) (*types.Transaction, error) {
	return _AMMFactory.Contract.RemoveLiquidity(&_AMMFactory.TransactOpts, _marketFactory, _marketId, _lpTokensIn, _minCollateralOut, _collateralRecipient)
}

// SellForCollateral is a paid mutator transaction binding the contract method 0x3fec5cac.
//
// Solidity: function sellForCollateral(address _marketFactory, uint256 _marketId, uint256 _outcome, uint256 _shareTokensIn, uint256 _setsOut) returns(uint256)
func (_AMMFactory *AMMFactoryTransactor) SellForCollateral(opts *bind.TransactOpts, _marketFactory common.Address, _marketId *big.Int, _outcome *big.Int, _shareTokensIn *big.Int, _setsOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.contract.Transact(opts, "sellForCollateral", _marketFactory, _marketId, _outcome, _shareTokensIn, _setsOut)
}

// SellForCollateral is a paid mutator transaction binding the contract method 0x3fec5cac.
//
// Solidity: function sellForCollateral(address _marketFactory, uint256 _marketId, uint256 _outcome, uint256 _shareTokensIn, uint256 _setsOut) returns(uint256)
func (_AMMFactory *AMMFactorySession) SellForCollateral(_marketFactory common.Address, _marketId *big.Int, _outcome *big.Int, _shareTokensIn *big.Int, _setsOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.Contract.SellForCollateral(&_AMMFactory.TransactOpts, _marketFactory, _marketId, _outcome, _shareTokensIn, _setsOut)
}

// SellForCollateral is a paid mutator transaction binding the contract method 0x3fec5cac.
//
// Solidity: function sellForCollateral(address _marketFactory, uint256 _marketId, uint256 _outcome, uint256 _shareTokensIn, uint256 _setsOut) returns(uint256)
func (_AMMFactory *AMMFactoryTransactorSession) SellForCollateral(_marketFactory common.Address, _marketId *big.Int, _outcome *big.Int, _shareTokensIn *big.Int, _setsOut *big.Int) (*types.Transaction, error) {
	return _AMMFactory.Contract.SellForCollateral(&_AMMFactory.TransactOpts, _marketFactory, _marketId, _outcome, _shareTokensIn, _setsOut)
}

// AMMFactoryLiquidityChangedIterator is returned from FilterLiquidityChanged and is used to iterate over the raw logs and unpacked data for LiquidityChanged events raised by the AMMFactory contract.
type AMMFactoryLiquidityChangedIterator struct {
	Event *AMMFactoryLiquidityChanged // Event containing the contract specifics and raw log

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
func (it *AMMFactoryLiquidityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMMFactoryLiquidityChanged)
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
		it.Event = new(AMMFactoryLiquidityChanged)
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
func (it *AMMFactoryLiquidityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMMFactoryLiquidityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMMFactoryLiquidityChanged represents a LiquidityChanged event raised by the AMMFactory contract.
type AMMFactoryLiquidityChanged struct {
	MarketFactory  common.Address
	MarketId       *big.Int
	User           common.Address
	Recipient      common.Address
	Collateral     *big.Int
	LpTokens       *big.Int
	SharesReturned []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidityChanged is a free log retrieval operation binding the contract event 0x9a1dccf45b5053e827f262e45fbb5211c2bd99497d340eecaebbd245eb48f4bc.
//
// Solidity: event LiquidityChanged(address indexed marketFactory, uint256 indexed marketId, address indexed user, address recipient, int256 collateral, int256 lpTokens, uint256[] sharesReturned)
func (_AMMFactory *AMMFactoryFilterer) FilterLiquidityChanged(opts *bind.FilterOpts, marketFactory []common.Address, marketId []*big.Int, user []common.Address) (*AMMFactoryLiquidityChangedIterator, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AMMFactory.contract.FilterLogs(opts, "LiquidityChanged", marketFactoryRule, marketIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AMMFactoryLiquidityChangedIterator{contract: _AMMFactory.contract, event: "LiquidityChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidityChanged is a free log subscription operation binding the contract event 0x9a1dccf45b5053e827f262e45fbb5211c2bd99497d340eecaebbd245eb48f4bc.
//
// Solidity: event LiquidityChanged(address indexed marketFactory, uint256 indexed marketId, address indexed user, address recipient, int256 collateral, int256 lpTokens, uint256[] sharesReturned)
func (_AMMFactory *AMMFactoryFilterer) WatchLiquidityChanged(opts *bind.WatchOpts, sink chan<- *AMMFactoryLiquidityChanged, marketFactory []common.Address, marketId []*big.Int, user []common.Address) (event.Subscription, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AMMFactory.contract.WatchLogs(opts, "LiquidityChanged", marketFactoryRule, marketIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMMFactoryLiquidityChanged)
				if err := _AMMFactory.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
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

// ParseLiquidityChanged is a log parse operation binding the contract event 0x9a1dccf45b5053e827f262e45fbb5211c2bd99497d340eecaebbd245eb48f4bc.
//
// Solidity: event LiquidityChanged(address indexed marketFactory, uint256 indexed marketId, address indexed user, address recipient, int256 collateral, int256 lpTokens, uint256[] sharesReturned)
func (_AMMFactory *AMMFactoryFilterer) ParseLiquidityChanged(log types.Log) (*AMMFactoryLiquidityChanged, error) {
	event := new(AMMFactoryLiquidityChanged)
	if err := _AMMFactory.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
		return nil, err
	}
	return event, nil
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
	Pool             common.Address
	MarketFactory    common.Address
	MarketId         *big.Int
	Creator          common.Address
	LpTokenRecipient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xfb27591822deac7913ea973cf845992f243e2ba5634ef8fbe0b6554aca37a428.
//
// Solidity: event PoolCreated(address pool, address indexed marketFactory, uint256 indexed marketId, address indexed creator, address lpTokenRecipient)
func (_AMMFactory *AMMFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, marketFactory []common.Address, marketId []*big.Int, creator []common.Address) (*AMMFactoryPoolCreatedIterator, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AMMFactory.contract.FilterLogs(opts, "PoolCreated", marketFactoryRule, marketIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &AMMFactoryPoolCreatedIterator{contract: _AMMFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xfb27591822deac7913ea973cf845992f243e2ba5634ef8fbe0b6554aca37a428.
//
// Solidity: event PoolCreated(address pool, address indexed marketFactory, uint256 indexed marketId, address indexed creator, address lpTokenRecipient)
func (_AMMFactory *AMMFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *AMMFactoryPoolCreated, marketFactory []common.Address, marketId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AMMFactory.contract.WatchLogs(opts, "PoolCreated", marketFactoryRule, marketIdRule, creatorRule)
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

// ParsePoolCreated is a log parse operation binding the contract event 0xfb27591822deac7913ea973cf845992f243e2ba5634ef8fbe0b6554aca37a428.
//
// Solidity: event PoolCreated(address pool, address indexed marketFactory, uint256 indexed marketId, address indexed creator, address lpTokenRecipient)
func (_AMMFactory *AMMFactoryFilterer) ParsePoolCreated(log types.Log) (*AMMFactoryPoolCreated, error) {
	event := new(AMMFactoryPoolCreated)
	if err := _AMMFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AMMFactorySharesSwappedIterator is returned from FilterSharesSwapped and is used to iterate over the raw logs and unpacked data for SharesSwapped events raised by the AMMFactory contract.
type AMMFactorySharesSwappedIterator struct {
	Event *AMMFactorySharesSwapped // Event containing the contract specifics and raw log

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
func (it *AMMFactorySharesSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMMFactorySharesSwapped)
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
		it.Event = new(AMMFactorySharesSwapped)
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
func (it *AMMFactorySharesSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMMFactorySharesSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMMFactorySharesSwapped represents a SharesSwapped event raised by the AMMFactory contract.
type AMMFactorySharesSwapped struct {
	MarketFactory common.Address
	MarketId      *big.Int
	User          common.Address
	Outcome       *big.Int
	Collateral    *big.Int
	Shares        *big.Int
	InOutRatio    []*big.Int
	Price         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSharesSwapped is a free log retrieval operation binding the contract event 0x9a8518831e2d9c2ea12fd8df17781e97121a2d578e17a7c0ddf4462f89fdd5a0.
//
// Solidity: event SharesSwapped(address indexed marketFactory, uint256 indexed marketId, address indexed user, uint256 outcome, int256 collateral, int256 shares, uint256[] inOutRatio, uint256 price)
func (_AMMFactory *AMMFactoryFilterer) FilterSharesSwapped(opts *bind.FilterOpts, marketFactory []common.Address, marketId []*big.Int, user []common.Address) (*AMMFactorySharesSwappedIterator, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AMMFactory.contract.FilterLogs(opts, "SharesSwapped", marketFactoryRule, marketIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AMMFactorySharesSwappedIterator{contract: _AMMFactory.contract, event: "SharesSwapped", logs: logs, sub: sub}, nil
}

// WatchSharesSwapped is a free log subscription operation binding the contract event 0x9a8518831e2d9c2ea12fd8df17781e97121a2d578e17a7c0ddf4462f89fdd5a0.
//
// Solidity: event SharesSwapped(address indexed marketFactory, uint256 indexed marketId, address indexed user, uint256 outcome, int256 collateral, int256 shares, uint256[] inOutRatio, uint256 price)
func (_AMMFactory *AMMFactoryFilterer) WatchSharesSwapped(opts *bind.WatchOpts, sink chan<- *AMMFactorySharesSwapped, marketFactory []common.Address, marketId []*big.Int, user []common.Address) (event.Subscription, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AMMFactory.contract.WatchLogs(opts, "SharesSwapped", marketFactoryRule, marketIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMMFactorySharesSwapped)
				if err := _AMMFactory.contract.UnpackLog(event, "SharesSwapped", log); err != nil {
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

// ParseSharesSwapped is a log parse operation binding the contract event 0x9a8518831e2d9c2ea12fd8df17781e97121a2d578e17a7c0ddf4462f89fdd5a0.
//
// Solidity: event SharesSwapped(address indexed marketFactory, uint256 indexed marketId, address indexed user, uint256 outcome, int256 collateral, int256 shares, uint256[] inOutRatio, uint256 price)
func (_AMMFactory *AMMFactoryFilterer) ParseSharesSwapped(log types.Log) (*AMMFactorySharesSwapped, error) {
	event := new(AMMFactorySharesSwapped)
	if err := _AMMFactory.contract.UnpackLog(event, "SharesSwapped", log); err != nil {
		return nil, err
	}
	return event, nil
}
