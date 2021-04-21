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

// AbstractMarketFactoryMarket is an auto generated low-level Go binding around an user-defined struct.
type AbstractMarketFactoryMarket struct {
	Creator     common.Address
	ShareTokens []common.Address
	EndTime     *big.Int
	Winner      common.Address
	CreatorFee  *big.Int
}

// TrustedMarketFactoryMarketDetails is an auto generated low-level Go binding around an user-defined struct.
type TrustedMarketFactoryMarketDetails struct {
	Description string
}

// TrustedMarketFactoryABI is the input ABI used to generate the binding from.
const TrustedMarketFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Full\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shareFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractFeePot\",\"name\":\"_feePot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creatorFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"outcomes\",\"type\":\"string[]\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"MarketResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"WinningsClaimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sharesToBurn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"burnShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"calcCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collateralIn\",\"type\":\"uint256\"}],\"name\":\"calcShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimManyWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"contractIERC20Full\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_names\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_symbols\",\"type\":\"string[]\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePot\",\"outputs\":[{\"internalType\":\"contractFeePot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMarket\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"contractOwnedERC20[]\",\"name\":\"shareTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"contractOwnedERC20\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"creatorFee\",\"type\":\"uint256\"}],\"internalType\":\"structAbstractMarketFactory.Market\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMarketDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structTrustedMarketFactory.MarketDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_shareToMint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mintShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"resolveMarket\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shareFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_winningOutcome\",\"type\":\"uint256\"}],\"name\":\"trustedResolveMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TrustedMarketFactory is an auto generated Go binding around an Ethereum contract.
type TrustedMarketFactory struct {
	TrustedMarketFactoryCaller     // Read-only binding to the contract
	TrustedMarketFactoryTransactor // Write-only binding to the contract
	TrustedMarketFactoryFilterer   // Log filterer for contract events
}

// TrustedMarketFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrustedMarketFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustedMarketFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrustedMarketFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustedMarketFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrustedMarketFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustedMarketFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrustedMarketFactorySession struct {
	Contract     *TrustedMarketFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TrustedMarketFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrustedMarketFactoryCallerSession struct {
	Contract *TrustedMarketFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TrustedMarketFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrustedMarketFactoryTransactorSession struct {
	Contract     *TrustedMarketFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TrustedMarketFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrustedMarketFactoryRaw struct {
	Contract *TrustedMarketFactory // Generic contract binding to access the raw methods on
}

// TrustedMarketFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrustedMarketFactoryCallerRaw struct {
	Contract *TrustedMarketFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TrustedMarketFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrustedMarketFactoryTransactorRaw struct {
	Contract *TrustedMarketFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrustedMarketFactory creates a new instance of TrustedMarketFactory, bound to a specific deployed contract.
func NewTrustedMarketFactory(address common.Address, backend bind.ContractBackend) (*TrustedMarketFactory, error) {
	contract, err := bindTrustedMarketFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrustedMarketFactory{TrustedMarketFactoryCaller: TrustedMarketFactoryCaller{contract: contract}, TrustedMarketFactoryTransactor: TrustedMarketFactoryTransactor{contract: contract}, TrustedMarketFactoryFilterer: TrustedMarketFactoryFilterer{contract: contract}}, nil
}

// NewTrustedMarketFactoryCaller creates a new read-only instance of TrustedMarketFactory, bound to a specific deployed contract.
func NewTrustedMarketFactoryCaller(address common.Address, caller bind.ContractCaller) (*TrustedMarketFactoryCaller, error) {
	contract, err := bindTrustedMarketFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrustedMarketFactoryCaller{contract: contract}, nil
}

// NewTrustedMarketFactoryTransactor creates a new write-only instance of TrustedMarketFactory, bound to a specific deployed contract.
func NewTrustedMarketFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TrustedMarketFactoryTransactor, error) {
	contract, err := bindTrustedMarketFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrustedMarketFactoryTransactor{contract: contract}, nil
}

// NewTrustedMarketFactoryFilterer creates a new log filterer instance of TrustedMarketFactory, bound to a specific deployed contract.
func NewTrustedMarketFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TrustedMarketFactoryFilterer, error) {
	contract, err := bindTrustedMarketFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrustedMarketFactoryFilterer{contract: contract}, nil
}

// bindTrustedMarketFactory binds a generic wrapper to an already deployed contract.
func bindTrustedMarketFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrustedMarketFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustedMarketFactory *TrustedMarketFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TrustedMarketFactory.Contract.TrustedMarketFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustedMarketFactory *TrustedMarketFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.TrustedMarketFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustedMarketFactory *TrustedMarketFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.TrustedMarketFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustedMarketFactory *TrustedMarketFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TrustedMarketFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustedMarketFactory *TrustedMarketFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustedMarketFactory *TrustedMarketFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.contract.Transact(opts, method, params...)
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) CalcCost(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrustedMarketFactory.contract.Call(opts, out, "calcCost", _shares)
	return *ret0, err
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactorySession) CalcCost(_shares *big.Int) (*big.Int, error) {
	return _TrustedMarketFactory.Contract.CalcCost(&_TrustedMarketFactory.CallOpts, _shares)
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) CalcCost(_shares *big.Int) (*big.Int, error) {
	return _TrustedMarketFactory.Contract.CalcCost(&_TrustedMarketFactory.CallOpts, _shares)
}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) CalcShares(opts *bind.CallOpts, _collateralIn *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrustedMarketFactory.contract.Call(opts, out, "calcShares", _collateralIn)
	return *ret0, err
}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactorySession) CalcShares(_collateralIn *big.Int) (*big.Int, error) {
	return _TrustedMarketFactory.Contract.CalcShares(&_TrustedMarketFactory.CallOpts, _collateralIn)
}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) CalcShares(_collateralIn *big.Int) (*big.Int, error) {
	return _TrustedMarketFactory.Contract.CalcShares(&_TrustedMarketFactory.CallOpts, _collateralIn)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrustedMarketFactory.contract.Call(opts, out, "collateral")
	return *ret0, err
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_TrustedMarketFactory *TrustedMarketFactorySession) Collateral() (common.Address, error) {
	return _TrustedMarketFactory.Contract.Collateral(&_TrustedMarketFactory.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) Collateral() (common.Address, error) {
	return _TrustedMarketFactory.Contract.Collateral(&_TrustedMarketFactory.CallOpts)
}

// CreatorFee is a free data retrieval call binding the contract method 0xe88958dc.
//
// Solidity: function creatorFee() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) CreatorFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrustedMarketFactory.contract.Call(opts, out, "creatorFee")
	return *ret0, err
}

// CreatorFee is a free data retrieval call binding the contract method 0xe88958dc.
//
// Solidity: function creatorFee() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactorySession) CreatorFee() (*big.Int, error) {
	return _TrustedMarketFactory.Contract.CreatorFee(&_TrustedMarketFactory.CallOpts)
}

// CreatorFee is a free data retrieval call binding the contract method 0xe88958dc.
//
// Solidity: function creatorFee() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) CreatorFee() (*big.Int, error) {
	return _TrustedMarketFactory.Contract.CreatorFee(&_TrustedMarketFactory.CallOpts)
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) FeePot(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrustedMarketFactory.contract.Call(opts, out, "feePot")
	return *ret0, err
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_TrustedMarketFactory *TrustedMarketFactorySession) FeePot() (common.Address, error) {
	return _TrustedMarketFactory.Contract.FeePot(&_TrustedMarketFactory.CallOpts)
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) FeePot() (common.Address, error) {
	return _TrustedMarketFactory.Contract.FeePot(&_TrustedMarketFactory.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns(AbstractMarketFactoryMarket)
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) GetMarket(opts *bind.CallOpts, _id *big.Int) (AbstractMarketFactoryMarket, error) {
	var (
		ret0 = new(AbstractMarketFactoryMarket)
	)
	out := ret0
	err := _TrustedMarketFactory.contract.Call(opts, out, "getMarket", _id)
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns(AbstractMarketFactoryMarket)
func (_TrustedMarketFactory *TrustedMarketFactorySession) GetMarket(_id *big.Int) (AbstractMarketFactoryMarket, error) {
	return _TrustedMarketFactory.Contract.GetMarket(&_TrustedMarketFactory.CallOpts, _id)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns(AbstractMarketFactoryMarket)
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) GetMarket(_id *big.Int) (AbstractMarketFactoryMarket, error) {
	return _TrustedMarketFactory.Contract.GetMarket(&_TrustedMarketFactory.CallOpts, _id)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _id) view returns(TrustedMarketFactoryMarketDetails)
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) GetMarketDetails(opts *bind.CallOpts, _id *big.Int) (TrustedMarketFactoryMarketDetails, error) {
	var (
		ret0 = new(TrustedMarketFactoryMarketDetails)
	)
	out := ret0
	err := _TrustedMarketFactory.contract.Call(opts, out, "getMarketDetails", _id)
	return *ret0, err
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _id) view returns(TrustedMarketFactoryMarketDetails)
func (_TrustedMarketFactory *TrustedMarketFactorySession) GetMarketDetails(_id *big.Int) (TrustedMarketFactoryMarketDetails, error) {
	return _TrustedMarketFactory.Contract.GetMarketDetails(&_TrustedMarketFactory.CallOpts, _id)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _id) view returns(TrustedMarketFactoryMarketDetails)
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) GetMarketDetails(_id *big.Int) (TrustedMarketFactoryMarketDetails, error) {
	return _TrustedMarketFactory.Contract.GetMarketDetails(&_TrustedMarketFactory.CallOpts, _id)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrustedMarketFactory.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_TrustedMarketFactory *TrustedMarketFactorySession) GetOwner() (common.Address, error) {
	return _TrustedMarketFactory.Contract.GetOwner(&_TrustedMarketFactory.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) GetOwner() (common.Address, error) {
	return _TrustedMarketFactory.Contract.GetOwner(&_TrustedMarketFactory.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrustedMarketFactory.contract.Call(opts, out, "marketCount")
	return *ret0, err
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactorySession) MarketCount() (*big.Int, error) {
	return _TrustedMarketFactory.Contract.MarketCount(&_TrustedMarketFactory.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) MarketCount() (*big.Int, error) {
	return _TrustedMarketFactory.Contract.MarketCount(&_TrustedMarketFactory.CallOpts)
}

// ResolveMarket is a free data retrieval call binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 ) pure returns()
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) ResolveMarket(opts *bind.CallOpts, arg0 *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _TrustedMarketFactory.contract.Call(opts, out, "resolveMarket", arg0)
	return err
}

// ResolveMarket is a free data retrieval call binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 ) pure returns()
func (_TrustedMarketFactory *TrustedMarketFactorySession) ResolveMarket(arg0 *big.Int) error {
	return _TrustedMarketFactory.Contract.ResolveMarket(&_TrustedMarketFactory.CallOpts, arg0)
}

// ResolveMarket is a free data retrieval call binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 ) pure returns()
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) ResolveMarket(arg0 *big.Int) error {
	return _TrustedMarketFactory.Contract.ResolveMarket(&_TrustedMarketFactory.CallOpts, arg0)
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) ShareFactor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrustedMarketFactory.contract.Call(opts, out, "shareFactor")
	return *ret0, err
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactorySession) ShareFactor() (*big.Int, error) {
	return _TrustedMarketFactory.Contract.ShareFactor(&_TrustedMarketFactory.CallOpts)
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) ShareFactor() (*big.Int, error) {
	return _TrustedMarketFactory.Contract.ShareFactor(&_TrustedMarketFactory.CallOpts)
}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCaller) StakerFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrustedMarketFactory.contract.Call(opts, out, "stakerFee")
	return *ret0, err
}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactorySession) StakerFee() (*big.Int, error) {
	return _TrustedMarketFactory.Contract.StakerFee(&_TrustedMarketFactory.CallOpts)
}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryCallerSession) StakerFee() (*big.Int, error) {
	return _TrustedMarketFactory.Contract.StakerFee(&_TrustedMarketFactory.CallOpts)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryTransactor) BurnShares(opts *bind.TransactOpts, _id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.contract.Transact(opts, "burnShares", _id, _sharesToBurn, _receiver)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactorySession) BurnShares(_id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.BurnShares(&_TrustedMarketFactory.TransactOpts, _id, _sharesToBurn, _receiver)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryTransactorSession) BurnShares(_id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.BurnShares(&_TrustedMarketFactory.TransactOpts, _id, _sharesToBurn, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryTransactor) ClaimManyWinnings(opts *bind.TransactOpts, _ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.contract.Transact(opts, "claimManyWinnings", _ids, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactorySession) ClaimManyWinnings(_ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.ClaimManyWinnings(&_TrustedMarketFactory.TransactOpts, _ids, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryTransactorSession) ClaimManyWinnings(_ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.ClaimManyWinnings(&_TrustedMarketFactory.TransactOpts, _ids, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryTransactor) ClaimWinnings(opts *bind.TransactOpts, _id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.contract.Transact(opts, "claimWinnings", _id, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactorySession) ClaimWinnings(_id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.ClaimWinnings(&_TrustedMarketFactory.TransactOpts, _id, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryTransactorSession) ClaimWinnings(_id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.ClaimWinnings(&_TrustedMarketFactory.TransactOpts, _id, _receiver)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7cb9bc3a.
//
// Solidity: function createMarket(address _creator, uint256 _endTime, string _description, string[] _names, string[] _symbols) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryTransactor) CreateMarket(opts *bind.TransactOpts, _creator common.Address, _endTime *big.Int, _description string, _names []string, _symbols []string) (*types.Transaction, error) {
	return _TrustedMarketFactory.contract.Transact(opts, "createMarket", _creator, _endTime, _description, _names, _symbols)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7cb9bc3a.
//
// Solidity: function createMarket(address _creator, uint256 _endTime, string _description, string[] _names, string[] _symbols) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactorySession) CreateMarket(_creator common.Address, _endTime *big.Int, _description string, _names []string, _symbols []string) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.CreateMarket(&_TrustedMarketFactory.TransactOpts, _creator, _endTime, _description, _names, _symbols)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x7cb9bc3a.
//
// Solidity: function createMarket(address _creator, uint256 _endTime, string _description, string[] _names, string[] _symbols) returns(uint256)
func (_TrustedMarketFactory *TrustedMarketFactoryTransactorSession) CreateMarket(_creator common.Address, _endTime *big.Int, _description string, _names []string, _symbols []string) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.CreateMarket(&_TrustedMarketFactory.TransactOpts, _creator, _endTime, _description, _names, _symbols)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_TrustedMarketFactory *TrustedMarketFactoryTransactor) MintShares(opts *bind.TransactOpts, _id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.contract.Transact(opts, "mintShares", _id, _shareToMint, _receiver)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_TrustedMarketFactory *TrustedMarketFactorySession) MintShares(_id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.MintShares(&_TrustedMarketFactory.TransactOpts, _id, _shareToMint, _receiver)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_TrustedMarketFactory *TrustedMarketFactoryTransactorSession) MintShares(_id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.MintShares(&_TrustedMarketFactory.TransactOpts, _id, _shareToMint, _receiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_TrustedMarketFactory *TrustedMarketFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_TrustedMarketFactory *TrustedMarketFactorySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.TransferOwnership(&_TrustedMarketFactory.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_TrustedMarketFactory *TrustedMarketFactoryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.TransferOwnership(&_TrustedMarketFactory.TransactOpts, _newOwner)
}

// TrustedResolveMarket is a paid mutator transaction binding the contract method 0x3f8dd7aa.
//
// Solidity: function trustedResolveMarket(uint256 _id, uint256 _winningOutcome) returns()
func (_TrustedMarketFactory *TrustedMarketFactoryTransactor) TrustedResolveMarket(opts *bind.TransactOpts, _id *big.Int, _winningOutcome *big.Int) (*types.Transaction, error) {
	return _TrustedMarketFactory.contract.Transact(opts, "trustedResolveMarket", _id, _winningOutcome)
}

// TrustedResolveMarket is a paid mutator transaction binding the contract method 0x3f8dd7aa.
//
// Solidity: function trustedResolveMarket(uint256 _id, uint256 _winningOutcome) returns()
func (_TrustedMarketFactory *TrustedMarketFactorySession) TrustedResolveMarket(_id *big.Int, _winningOutcome *big.Int) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.TrustedResolveMarket(&_TrustedMarketFactory.TransactOpts, _id, _winningOutcome)
}

// TrustedResolveMarket is a paid mutator transaction binding the contract method 0x3f8dd7aa.
//
// Solidity: function trustedResolveMarket(uint256 _id, uint256 _winningOutcome) returns()
func (_TrustedMarketFactory *TrustedMarketFactoryTransactorSession) TrustedResolveMarket(_id *big.Int, _winningOutcome *big.Int) (*types.Transaction, error) {
	return _TrustedMarketFactory.Contract.TrustedResolveMarket(&_TrustedMarketFactory.TransactOpts, _id, _winningOutcome)
}

// TrustedMarketFactoryMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the TrustedMarketFactory contract.
type TrustedMarketFactoryMarketCreatedIterator struct {
	Event *TrustedMarketFactoryMarketCreated // Event containing the contract specifics and raw log

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
func (it *TrustedMarketFactoryMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedMarketFactoryMarketCreated)
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
		it.Event = new(TrustedMarketFactoryMarketCreated)
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
func (it *TrustedMarketFactoryMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedMarketFactoryMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedMarketFactoryMarketCreated represents a MarketCreated event raised by the TrustedMarketFactory contract.
type TrustedMarketFactoryMarketCreated struct {
	Id          *big.Int
	Creator     common.Address
	EndTime     *big.Int
	Description string
	Outcomes    []string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0xa1bb41461c32765a0cc838c35ce6b8e28985bb6a069dfe2af0873796438670d4.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 _endTime, string description, string[] outcomes)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) FilterMarketCreated(opts *bind.FilterOpts) (*TrustedMarketFactoryMarketCreatedIterator, error) {

	logs, sub, err := _TrustedMarketFactory.contract.FilterLogs(opts, "MarketCreated")
	if err != nil {
		return nil, err
	}
	return &TrustedMarketFactoryMarketCreatedIterator{contract: _TrustedMarketFactory.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0xa1bb41461c32765a0cc838c35ce6b8e28985bb6a069dfe2af0873796438670d4.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 _endTime, string description, string[] outcomes)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *TrustedMarketFactoryMarketCreated) (event.Subscription, error) {

	logs, sub, err := _TrustedMarketFactory.contract.WatchLogs(opts, "MarketCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedMarketFactoryMarketCreated)
				if err := _TrustedMarketFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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

// ParseMarketCreated is a log parse operation binding the contract event 0xa1bb41461c32765a0cc838c35ce6b8e28985bb6a069dfe2af0873796438670d4.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 _endTime, string description, string[] outcomes)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) ParseMarketCreated(log types.Log) (*TrustedMarketFactoryMarketCreated, error) {
	event := new(TrustedMarketFactoryMarketCreated)
	if err := _TrustedMarketFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrustedMarketFactoryMarketResolvedIterator is returned from FilterMarketResolved and is used to iterate over the raw logs and unpacked data for MarketResolved events raised by the TrustedMarketFactory contract.
type TrustedMarketFactoryMarketResolvedIterator struct {
	Event *TrustedMarketFactoryMarketResolved // Event containing the contract specifics and raw log

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
func (it *TrustedMarketFactoryMarketResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedMarketFactoryMarketResolved)
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
		it.Event = new(TrustedMarketFactoryMarketResolved)
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
func (it *TrustedMarketFactoryMarketResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedMarketFactoryMarketResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedMarketFactoryMarketResolved represents a MarketResolved event raised by the TrustedMarketFactory contract.
type TrustedMarketFactoryMarketResolved struct {
	Id     *big.Int
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketResolved is a free log retrieval operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) FilterMarketResolved(opts *bind.FilterOpts) (*TrustedMarketFactoryMarketResolvedIterator, error) {

	logs, sub, err := _TrustedMarketFactory.contract.FilterLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return &TrustedMarketFactoryMarketResolvedIterator{contract: _TrustedMarketFactory.contract, event: "MarketResolved", logs: logs, sub: sub}, nil
}

// WatchMarketResolved is a free log subscription operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) WatchMarketResolved(opts *bind.WatchOpts, sink chan<- *TrustedMarketFactoryMarketResolved) (event.Subscription, error) {

	logs, sub, err := _TrustedMarketFactory.contract.WatchLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedMarketFactoryMarketResolved)
				if err := _TrustedMarketFactory.contract.UnpackLog(event, "MarketResolved", log); err != nil {
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

// ParseMarketResolved is a log parse operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) ParseMarketResolved(log types.Log) (*TrustedMarketFactoryMarketResolved, error) {
	event := new(TrustedMarketFactoryMarketResolved)
	if err := _TrustedMarketFactory.contract.UnpackLog(event, "MarketResolved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrustedMarketFactorySharesBurnedIterator is returned from FilterSharesBurned and is used to iterate over the raw logs and unpacked data for SharesBurned events raised by the TrustedMarketFactory contract.
type TrustedMarketFactorySharesBurnedIterator struct {
	Event *TrustedMarketFactorySharesBurned // Event containing the contract specifics and raw log

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
func (it *TrustedMarketFactorySharesBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedMarketFactorySharesBurned)
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
		it.Event = new(TrustedMarketFactorySharesBurned)
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
func (it *TrustedMarketFactorySharesBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedMarketFactorySharesBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedMarketFactorySharesBurned represents a SharesBurned event raised by the TrustedMarketFactory contract.
type TrustedMarketFactorySharesBurned struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSharesBurned is a free log retrieval operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) FilterSharesBurned(opts *bind.FilterOpts) (*TrustedMarketFactorySharesBurnedIterator, error) {

	logs, sub, err := _TrustedMarketFactory.contract.FilterLogs(opts, "SharesBurned")
	if err != nil {
		return nil, err
	}
	return &TrustedMarketFactorySharesBurnedIterator{contract: _TrustedMarketFactory.contract, event: "SharesBurned", logs: logs, sub: sub}, nil
}

// WatchSharesBurned is a free log subscription operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) WatchSharesBurned(opts *bind.WatchOpts, sink chan<- *TrustedMarketFactorySharesBurned) (event.Subscription, error) {

	logs, sub, err := _TrustedMarketFactory.contract.WatchLogs(opts, "SharesBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedMarketFactorySharesBurned)
				if err := _TrustedMarketFactory.contract.UnpackLog(event, "SharesBurned", log); err != nil {
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

// ParseSharesBurned is a log parse operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) ParseSharesBurned(log types.Log) (*TrustedMarketFactorySharesBurned, error) {
	event := new(TrustedMarketFactorySharesBurned)
	if err := _TrustedMarketFactory.contract.UnpackLog(event, "SharesBurned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrustedMarketFactorySharesMintedIterator is returned from FilterSharesMinted and is used to iterate over the raw logs and unpacked data for SharesMinted events raised by the TrustedMarketFactory contract.
type TrustedMarketFactorySharesMintedIterator struct {
	Event *TrustedMarketFactorySharesMinted // Event containing the contract specifics and raw log

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
func (it *TrustedMarketFactorySharesMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedMarketFactorySharesMinted)
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
		it.Event = new(TrustedMarketFactorySharesMinted)
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
func (it *TrustedMarketFactorySharesMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedMarketFactorySharesMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedMarketFactorySharesMinted represents a SharesMinted event raised by the TrustedMarketFactory contract.
type TrustedMarketFactorySharesMinted struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSharesMinted is a free log retrieval operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) FilterSharesMinted(opts *bind.FilterOpts) (*TrustedMarketFactorySharesMintedIterator, error) {

	logs, sub, err := _TrustedMarketFactory.contract.FilterLogs(opts, "SharesMinted")
	if err != nil {
		return nil, err
	}
	return &TrustedMarketFactorySharesMintedIterator{contract: _TrustedMarketFactory.contract, event: "SharesMinted", logs: logs, sub: sub}, nil
}

// WatchSharesMinted is a free log subscription operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) WatchSharesMinted(opts *bind.WatchOpts, sink chan<- *TrustedMarketFactorySharesMinted) (event.Subscription, error) {

	logs, sub, err := _TrustedMarketFactory.contract.WatchLogs(opts, "SharesMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedMarketFactorySharesMinted)
				if err := _TrustedMarketFactory.contract.UnpackLog(event, "SharesMinted", log); err != nil {
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

// ParseSharesMinted is a log parse operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) ParseSharesMinted(log types.Log) (*TrustedMarketFactorySharesMinted, error) {
	event := new(TrustedMarketFactorySharesMinted)
	if err := _TrustedMarketFactory.contract.UnpackLog(event, "SharesMinted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrustedMarketFactoryWinningsClaimedIterator is returned from FilterWinningsClaimed and is used to iterate over the raw logs and unpacked data for WinningsClaimed events raised by the TrustedMarketFactory contract.
type TrustedMarketFactoryWinningsClaimedIterator struct {
	Event *TrustedMarketFactoryWinningsClaimed // Event containing the contract specifics and raw log

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
func (it *TrustedMarketFactoryWinningsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedMarketFactoryWinningsClaimed)
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
		it.Event = new(TrustedMarketFactoryWinningsClaimed)
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
func (it *TrustedMarketFactoryWinningsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedMarketFactoryWinningsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedMarketFactoryWinningsClaimed represents a WinningsClaimed event raised by the TrustedMarketFactory contract.
type TrustedMarketFactoryWinningsClaimed struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWinningsClaimed is a free log retrieval operation binding the contract event 0x2bdd7a5109193ce6270ec3b4afcf4ccd4a06c27742ba11f660498cb41433bb00.
//
// Solidity: event WinningsClaimed(uint256 id, uint256 amount, address receiver)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) FilterWinningsClaimed(opts *bind.FilterOpts) (*TrustedMarketFactoryWinningsClaimedIterator, error) {

	logs, sub, err := _TrustedMarketFactory.contract.FilterLogs(opts, "WinningsClaimed")
	if err != nil {
		return nil, err
	}
	return &TrustedMarketFactoryWinningsClaimedIterator{contract: _TrustedMarketFactory.contract, event: "WinningsClaimed", logs: logs, sub: sub}, nil
}

// WatchWinningsClaimed is a free log subscription operation binding the contract event 0x2bdd7a5109193ce6270ec3b4afcf4ccd4a06c27742ba11f660498cb41433bb00.
//
// Solidity: event WinningsClaimed(uint256 id, uint256 amount, address receiver)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) WatchWinningsClaimed(opts *bind.WatchOpts, sink chan<- *TrustedMarketFactoryWinningsClaimed) (event.Subscription, error) {

	logs, sub, err := _TrustedMarketFactory.contract.WatchLogs(opts, "WinningsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedMarketFactoryWinningsClaimed)
				if err := _TrustedMarketFactory.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
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

// ParseWinningsClaimed is a log parse operation binding the contract event 0x2bdd7a5109193ce6270ec3b4afcf4ccd4a06c27742ba11f660498cb41433bb00.
//
// Solidity: event WinningsClaimed(uint256 id, uint256 amount, address receiver)
func (_TrustedMarketFactory *TrustedMarketFactoryFilterer) ParseWinningsClaimed(log types.Log) (*TrustedMarketFactoryWinningsClaimed, error) {
	event := new(TrustedMarketFactoryWinningsClaimed)
	if err := _TrustedMarketFactory.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}
