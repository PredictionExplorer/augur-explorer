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

// SportsLinkMarketFactoryMarketDetails is an auto generated low-level Go binding around an user-defined struct.
type SportsLinkMarketFactoryMarketDetails struct {
	EventId            *big.Int
	HomeTeamId         *big.Int
	AwayTeamId         *big.Int
	EstimatedStartTime *big.Int
	MarketType         uint8
	Value0             *big.Int
}

// SportsLinkMarketFactoryABI is the input ABI used to generate the binding from.
const SportsLinkMarketFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Full\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shareFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractFeePot\",\"name\":\"_feePot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creatorFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumSportsLinkMarketFactory.MarketType\",\"name\":\"marketType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"homeTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"awayTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"estimatedStarTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"score\",\"type\":\"int256\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"MarketResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"WinningsClaimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sharesToBurn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"burnShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"calcCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collateralIn\",\"type\":\"uint256\"}],\"name\":\"calcShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimManyWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"contractIERC20Full\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eventId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_homeTeamId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_awayTeamId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_homeSpread\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_overUnderTotal\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedStartTime\",\"type\":\"uint256\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_ids\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePot\",\"outputs\":[{\"internalType\":\"contractFeePot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMarket\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"contractOwnedERC20[]\",\"name\":\"shareTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"contractOwnedERC20\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"creatorFee\",\"type\":\"uint256\"}],\"internalType\":\"structAbstractMarketFactory.Market\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"homeTeamId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"awayTeamId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimatedStartTime\",\"type\":\"uint256\"},{\"internalType\":\"enumSportsLinkMarketFactory.MarketType\",\"name\":\"marketType\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value0\",\"type\":\"int256\"}],\"internalType\":\"structSportsLinkMarketFactory.MarketDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_shareToMint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mintShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"resolveMarket\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shareFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eventId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_homeScore\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_awayScore\",\"type\":\"int256\"}],\"name\":\"trustedResolveMarkets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SportsLinkMarketFactory is an auto generated Go binding around an Ethereum contract.
type SportsLinkMarketFactory struct {
	SportsLinkMarketFactoryCaller     // Read-only binding to the contract
	SportsLinkMarketFactoryTransactor // Write-only binding to the contract
	SportsLinkMarketFactoryFilterer   // Log filterer for contract events
}

// SportsLinkMarketFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SportsLinkMarketFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SportsLinkMarketFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SportsLinkMarketFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SportsLinkMarketFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SportsLinkMarketFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SportsLinkMarketFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SportsLinkMarketFactorySession struct {
	Contract     *SportsLinkMarketFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SportsLinkMarketFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SportsLinkMarketFactoryCallerSession struct {
	Contract *SportsLinkMarketFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// SportsLinkMarketFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SportsLinkMarketFactoryTransactorSession struct {
	Contract     *SportsLinkMarketFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// SportsLinkMarketFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SportsLinkMarketFactoryRaw struct {
	Contract *SportsLinkMarketFactory // Generic contract binding to access the raw methods on
}

// SportsLinkMarketFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SportsLinkMarketFactoryCallerRaw struct {
	Contract *SportsLinkMarketFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SportsLinkMarketFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SportsLinkMarketFactoryTransactorRaw struct {
	Contract *SportsLinkMarketFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSportsLinkMarketFactory creates a new instance of SportsLinkMarketFactory, bound to a specific deployed contract.
func NewSportsLinkMarketFactory(address common.Address, backend bind.ContractBackend) (*SportsLinkMarketFactory, error) {
	contract, err := bindSportsLinkMarketFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactory{SportsLinkMarketFactoryCaller: SportsLinkMarketFactoryCaller{contract: contract}, SportsLinkMarketFactoryTransactor: SportsLinkMarketFactoryTransactor{contract: contract}, SportsLinkMarketFactoryFilterer: SportsLinkMarketFactoryFilterer{contract: contract}}, nil
}

// NewSportsLinkMarketFactoryCaller creates a new read-only instance of SportsLinkMarketFactory, bound to a specific deployed contract.
func NewSportsLinkMarketFactoryCaller(address common.Address, caller bind.ContractCaller) (*SportsLinkMarketFactoryCaller, error) {
	contract, err := bindSportsLinkMarketFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryCaller{contract: contract}, nil
}

// NewSportsLinkMarketFactoryTransactor creates a new write-only instance of SportsLinkMarketFactory, bound to a specific deployed contract.
func NewSportsLinkMarketFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SportsLinkMarketFactoryTransactor, error) {
	contract, err := bindSportsLinkMarketFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryTransactor{contract: contract}, nil
}

// NewSportsLinkMarketFactoryFilterer creates a new log filterer instance of SportsLinkMarketFactory, bound to a specific deployed contract.
func NewSportsLinkMarketFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SportsLinkMarketFactoryFilterer, error) {
	contract, err := bindSportsLinkMarketFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryFilterer{contract: contract}, nil
}

// bindSportsLinkMarketFactory binds a generic wrapper to an already deployed contract.
func bindSportsLinkMarketFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SportsLinkMarketFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SportsLinkMarketFactory.Contract.SportsLinkMarketFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SportsLinkMarketFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SportsLinkMarketFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SportsLinkMarketFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.contract.Transact(opts, method, params...)
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) CalcCost(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "calcCost", _shares)
	return *ret0, err
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) CalcCost(_shares *big.Int) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.CalcCost(&_SportsLinkMarketFactory.CallOpts, _shares)
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) CalcCost(_shares *big.Int) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.CalcCost(&_SportsLinkMarketFactory.CallOpts, _shares)
}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) CalcShares(opts *bind.CallOpts, _collateralIn *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "calcShares", _collateralIn)
	return *ret0, err
}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) CalcShares(_collateralIn *big.Int) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.CalcShares(&_SportsLinkMarketFactory.CallOpts, _collateralIn)
}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) CalcShares(_collateralIn *big.Int) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.CalcShares(&_SportsLinkMarketFactory.CallOpts, _collateralIn)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "collateral")
	return *ret0, err
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) Collateral() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.Collateral(&_SportsLinkMarketFactory.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) Collateral() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.Collateral(&_SportsLinkMarketFactory.CallOpts)
}

// CreatorFee is a free data retrieval call binding the contract method 0xe88958dc.
//
// Solidity: function creatorFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) CreatorFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "creatorFee")
	return *ret0, err
}

// CreatorFee is a free data retrieval call binding the contract method 0xe88958dc.
//
// Solidity: function creatorFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) CreatorFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.CreatorFee(&_SportsLinkMarketFactory.CallOpts)
}

// CreatorFee is a free data retrieval call binding the contract method 0xe88958dc.
//
// Solidity: function creatorFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) CreatorFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.CreatorFee(&_SportsLinkMarketFactory.CallOpts)
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) FeePot(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "feePot")
	return *ret0, err
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) FeePot() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.FeePot(&_SportsLinkMarketFactory.CallOpts)
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) FeePot() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.FeePot(&_SportsLinkMarketFactory.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns(AbstractMarketFactoryMarket)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) GetMarket(opts *bind.CallOpts, _id *big.Int) (AbstractMarketFactoryMarket, error) {
	var (
		ret0 = new(AbstractMarketFactoryMarket)
	)
	out := ret0
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "getMarket", _id)
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns(AbstractMarketFactoryMarket)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) GetMarket(_id *big.Int) (AbstractMarketFactoryMarket, error) {
	return _SportsLinkMarketFactory.Contract.GetMarket(&_SportsLinkMarketFactory.CallOpts, _id)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns(AbstractMarketFactoryMarket)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) GetMarket(_id *big.Int) (AbstractMarketFactoryMarket, error) {
	return _SportsLinkMarketFactory.Contract.GetMarket(&_SportsLinkMarketFactory.CallOpts, _id)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _marketId) view returns(SportsLinkMarketFactoryMarketDetails)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) GetMarketDetails(opts *bind.CallOpts, _marketId *big.Int) (SportsLinkMarketFactoryMarketDetails, error) {
	var (
		ret0 = new(SportsLinkMarketFactoryMarketDetails)
	)
	out := ret0
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "getMarketDetails", _marketId)
	return *ret0, err
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _marketId) view returns(SportsLinkMarketFactoryMarketDetails)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) GetMarketDetails(_marketId *big.Int) (SportsLinkMarketFactoryMarketDetails, error) {
	return _SportsLinkMarketFactory.Contract.GetMarketDetails(&_SportsLinkMarketFactory.CallOpts, _marketId)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _marketId) view returns(SportsLinkMarketFactoryMarketDetails)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) GetMarketDetails(_marketId *big.Int) (SportsLinkMarketFactoryMarketDetails, error) {
	return _SportsLinkMarketFactory.Contract.GetMarketDetails(&_SportsLinkMarketFactory.CallOpts, _marketId)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) GetOwner() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.GetOwner(&_SportsLinkMarketFactory.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) GetOwner() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.GetOwner(&_SportsLinkMarketFactory.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "marketCount")
	return *ret0, err
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) MarketCount() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.MarketCount(&_SportsLinkMarketFactory.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) MarketCount() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.MarketCount(&_SportsLinkMarketFactory.CallOpts)
}

// ResolveMarket is a free data retrieval call binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 ) pure returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) ResolveMarket(opts *bind.CallOpts, arg0 *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "resolveMarket", arg0)
	return err
}

// ResolveMarket is a free data retrieval call binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 ) pure returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ResolveMarket(arg0 *big.Int) error {
	return _SportsLinkMarketFactory.Contract.ResolveMarket(&_SportsLinkMarketFactory.CallOpts, arg0)
}

// ResolveMarket is a free data retrieval call binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 ) pure returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) ResolveMarket(arg0 *big.Int) error {
	return _SportsLinkMarketFactory.Contract.ResolveMarket(&_SportsLinkMarketFactory.CallOpts, arg0)
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) ShareFactor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "shareFactor")
	return *ret0, err
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ShareFactor() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ShareFactor(&_SportsLinkMarketFactory.CallOpts)
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) ShareFactor() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ShareFactor(&_SportsLinkMarketFactory.CallOpts)
}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) StakerFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SportsLinkMarketFactory.contract.Call(opts, out, "stakerFee")
	return *ret0, err
}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) StakerFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.StakerFee(&_SportsLinkMarketFactory.CallOpts)
}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) StakerFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.StakerFee(&_SportsLinkMarketFactory.CallOpts)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) BurnShares(opts *bind.TransactOpts, _id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "burnShares", _id, _sharesToBurn, _receiver)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) BurnShares(_id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.BurnShares(&_SportsLinkMarketFactory.TransactOpts, _id, _sharesToBurn, _receiver)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) BurnShares(_id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.BurnShares(&_SportsLinkMarketFactory.TransactOpts, _id, _sharesToBurn, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) ClaimManyWinnings(opts *bind.TransactOpts, _ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "claimManyWinnings", _ids, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ClaimManyWinnings(_ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimManyWinnings(&_SportsLinkMarketFactory.TransactOpts, _ids, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) ClaimManyWinnings(_ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimManyWinnings(&_SportsLinkMarketFactory.TransactOpts, _ids, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) ClaimWinnings(opts *bind.TransactOpts, _id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "claimWinnings", _id, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ClaimWinnings(_id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimWinnings(&_SportsLinkMarketFactory.TransactOpts, _id, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) ClaimWinnings(_id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimWinnings(&_SportsLinkMarketFactory.TransactOpts, _id, _receiver)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x97104223.
//
// Solidity: function createMarket(address _creator, uint256 _endTime, uint256 _eventId, uint256 _homeTeamId, uint256 _awayTeamId, int256 _homeSpread, int256 _overUnderTotal, uint256 _estimatedStartTime) returns(uint256[3] _ids)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) CreateMarket(opts *bind.TransactOpts, _creator common.Address, _endTime *big.Int, _eventId *big.Int, _homeTeamId *big.Int, _awayTeamId *big.Int, _homeSpread *big.Int, _overUnderTotal *big.Int, _estimatedStartTime *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "createMarket", _creator, _endTime, _eventId, _homeTeamId, _awayTeamId, _homeSpread, _overUnderTotal, _estimatedStartTime)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x97104223.
//
// Solidity: function createMarket(address _creator, uint256 _endTime, uint256 _eventId, uint256 _homeTeamId, uint256 _awayTeamId, int256 _homeSpread, int256 _overUnderTotal, uint256 _estimatedStartTime) returns(uint256[3] _ids)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) CreateMarket(_creator common.Address, _endTime *big.Int, _eventId *big.Int, _homeTeamId *big.Int, _awayTeamId *big.Int, _homeSpread *big.Int, _overUnderTotal *big.Int, _estimatedStartTime *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.CreateMarket(&_SportsLinkMarketFactory.TransactOpts, _creator, _endTime, _eventId, _homeTeamId, _awayTeamId, _homeSpread, _overUnderTotal, _estimatedStartTime)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x97104223.
//
// Solidity: function createMarket(address _creator, uint256 _endTime, uint256 _eventId, uint256 _homeTeamId, uint256 _awayTeamId, int256 _homeSpread, int256 _overUnderTotal, uint256 _estimatedStartTime) returns(uint256[3] _ids)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) CreateMarket(_creator common.Address, _endTime *big.Int, _eventId *big.Int, _homeTeamId *big.Int, _awayTeamId *big.Int, _homeSpread *big.Int, _overUnderTotal *big.Int, _estimatedStartTime *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.CreateMarket(&_SportsLinkMarketFactory.TransactOpts, _creator, _endTime, _eventId, _homeTeamId, _awayTeamId, _homeSpread, _overUnderTotal, _estimatedStartTime)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) MintShares(opts *bind.TransactOpts, _id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "mintShares", _id, _shareToMint, _receiver)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) MintShares(_id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.MintShares(&_SportsLinkMarketFactory.TransactOpts, _id, _shareToMint, _receiver)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) MintShares(_id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.MintShares(&_SportsLinkMarketFactory.TransactOpts, _id, _shareToMint, _receiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.TransferOwnership(&_SportsLinkMarketFactory.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.TransferOwnership(&_SportsLinkMarketFactory.TransactOpts, _newOwner)
}

// TrustedResolveMarkets is a paid mutator transaction binding the contract method 0xd93674d7.
//
// Solidity: function trustedResolveMarkets(uint256 _eventId, int256 _homeScore, int256 _awayScore) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) TrustedResolveMarkets(opts *bind.TransactOpts, _eventId *big.Int, _homeScore *big.Int, _awayScore *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "trustedResolveMarkets", _eventId, _homeScore, _awayScore)
}

// TrustedResolveMarkets is a paid mutator transaction binding the contract method 0xd93674d7.
//
// Solidity: function trustedResolveMarkets(uint256 _eventId, int256 _homeScore, int256 _awayScore) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) TrustedResolveMarkets(_eventId *big.Int, _homeScore *big.Int, _awayScore *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.TrustedResolveMarkets(&_SportsLinkMarketFactory.TransactOpts, _eventId, _homeScore, _awayScore)
}

// TrustedResolveMarkets is a paid mutator transaction binding the contract method 0xd93674d7.
//
// Solidity: function trustedResolveMarkets(uint256 _eventId, int256 _homeScore, int256 _awayScore) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) TrustedResolveMarkets(_eventId *big.Int, _homeScore *big.Int, _awayScore *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.TrustedResolveMarkets(&_SportsLinkMarketFactory.TransactOpts, _eventId, _homeScore, _awayScore)
}

// SportsLinkMarketFactoryMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryMarketCreatedIterator struct {
	Event *SportsLinkMarketFactoryMarketCreated // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactoryMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactoryMarketCreated)
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
		it.Event = new(SportsLinkMarketFactoryMarketCreated)
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
func (it *SportsLinkMarketFactoryMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactoryMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactoryMarketCreated represents a MarketCreated event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryMarketCreated struct {
	Id                *big.Int
	Creator           common.Address
	EndTime           *big.Int
	MarketType        uint8
	EventId           *big.Int
	HomeTeamId        *big.Int
	AwayTeamId        *big.Int
	EstimatedStarTime *big.Int
	Score             *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0xafad6545e5200f9fdf4db34dfca61a9e7d72518593fd5155d11fd69c20e1555f.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 endTime, uint8 marketType, uint256 indexed eventId, uint256 homeTeamId, uint256 awayTeamId, uint256 estimatedStarTime, int256 score)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterMarketCreated(opts *bind.FilterOpts, eventId []*big.Int) (*SportsLinkMarketFactoryMarketCreatedIterator, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "MarketCreated", eventIdRule)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryMarketCreatedIterator{contract: _SportsLinkMarketFactory.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0xafad6545e5200f9fdf4db34dfca61a9e7d72518593fd5155d11fd69c20e1555f.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 endTime, uint8 marketType, uint256 indexed eventId, uint256 homeTeamId, uint256 awayTeamId, uint256 estimatedStarTime, int256 score)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactoryMarketCreated, eventId []*big.Int) (event.Subscription, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "MarketCreated", eventIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactoryMarketCreated)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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

// ParseMarketCreated is a log parse operation binding the contract event 0xafad6545e5200f9fdf4db34dfca61a9e7d72518593fd5155d11fd69c20e1555f.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 endTime, uint8 marketType, uint256 indexed eventId, uint256 homeTeamId, uint256 awayTeamId, uint256 estimatedStarTime, int256 score)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseMarketCreated(log types.Log) (*SportsLinkMarketFactoryMarketCreated, error) {
	event := new(SportsLinkMarketFactoryMarketCreated)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SportsLinkMarketFactoryMarketResolvedIterator is returned from FilterMarketResolved and is used to iterate over the raw logs and unpacked data for MarketResolved events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryMarketResolvedIterator struct {
	Event *SportsLinkMarketFactoryMarketResolved // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactoryMarketResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactoryMarketResolved)
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
		it.Event = new(SportsLinkMarketFactoryMarketResolved)
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
func (it *SportsLinkMarketFactoryMarketResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactoryMarketResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactoryMarketResolved represents a MarketResolved event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryMarketResolved struct {
	Id     *big.Int
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketResolved is a free log retrieval operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterMarketResolved(opts *bind.FilterOpts) (*SportsLinkMarketFactoryMarketResolvedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryMarketResolvedIterator{contract: _SportsLinkMarketFactory.contract, event: "MarketResolved", logs: logs, sub: sub}, nil
}

// WatchMarketResolved is a free log subscription operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchMarketResolved(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactoryMarketResolved) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactoryMarketResolved)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "MarketResolved", log); err != nil {
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
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseMarketResolved(log types.Log) (*SportsLinkMarketFactoryMarketResolved, error) {
	event := new(SportsLinkMarketFactoryMarketResolved)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "MarketResolved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SportsLinkMarketFactorySharesBurnedIterator is returned from FilterSharesBurned and is used to iterate over the raw logs and unpacked data for SharesBurned events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySharesBurnedIterator struct {
	Event *SportsLinkMarketFactorySharesBurned // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactorySharesBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactorySharesBurned)
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
		it.Event = new(SportsLinkMarketFactorySharesBurned)
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
func (it *SportsLinkMarketFactorySharesBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactorySharesBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactorySharesBurned represents a SharesBurned event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySharesBurned struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSharesBurned is a free log retrieval operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterSharesBurned(opts *bind.FilterOpts) (*SportsLinkMarketFactorySharesBurnedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "SharesBurned")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactorySharesBurnedIterator{contract: _SportsLinkMarketFactory.contract, event: "SharesBurned", logs: logs, sub: sub}, nil
}

// WatchSharesBurned is a free log subscription operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchSharesBurned(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactorySharesBurned) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "SharesBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactorySharesBurned)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SharesBurned", log); err != nil {
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
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseSharesBurned(log types.Log) (*SportsLinkMarketFactorySharesBurned, error) {
	event := new(SportsLinkMarketFactorySharesBurned)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SharesBurned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SportsLinkMarketFactorySharesMintedIterator is returned from FilterSharesMinted and is used to iterate over the raw logs and unpacked data for SharesMinted events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySharesMintedIterator struct {
	Event *SportsLinkMarketFactorySharesMinted // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactorySharesMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactorySharesMinted)
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
		it.Event = new(SportsLinkMarketFactorySharesMinted)
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
func (it *SportsLinkMarketFactorySharesMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactorySharesMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactorySharesMinted represents a SharesMinted event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySharesMinted struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSharesMinted is a free log retrieval operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterSharesMinted(opts *bind.FilterOpts) (*SportsLinkMarketFactorySharesMintedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "SharesMinted")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactorySharesMintedIterator{contract: _SportsLinkMarketFactory.contract, event: "SharesMinted", logs: logs, sub: sub}, nil
}

// WatchSharesMinted is a free log subscription operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchSharesMinted(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactorySharesMinted) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "SharesMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactorySharesMinted)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SharesMinted", log); err != nil {
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
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseSharesMinted(log types.Log) (*SportsLinkMarketFactorySharesMinted, error) {
	event := new(SportsLinkMarketFactorySharesMinted)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SharesMinted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SportsLinkMarketFactoryWinningsClaimedIterator is returned from FilterWinningsClaimed and is used to iterate over the raw logs and unpacked data for WinningsClaimed events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryWinningsClaimedIterator struct {
	Event *SportsLinkMarketFactoryWinningsClaimed // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactoryWinningsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactoryWinningsClaimed)
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
		it.Event = new(SportsLinkMarketFactoryWinningsClaimed)
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
func (it *SportsLinkMarketFactoryWinningsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactoryWinningsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactoryWinningsClaimed represents a WinningsClaimed event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryWinningsClaimed struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWinningsClaimed is a free log retrieval operation binding the contract event 0x2bdd7a5109193ce6270ec3b4afcf4ccd4a06c27742ba11f660498cb41433bb00.
//
// Solidity: event WinningsClaimed(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterWinningsClaimed(opts *bind.FilterOpts) (*SportsLinkMarketFactoryWinningsClaimedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "WinningsClaimed")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryWinningsClaimedIterator{contract: _SportsLinkMarketFactory.contract, event: "WinningsClaimed", logs: logs, sub: sub}, nil
}

// WatchWinningsClaimed is a free log subscription operation binding the contract event 0x2bdd7a5109193ce6270ec3b4afcf4ccd4a06c27742ba11f660498cb41433bb00.
//
// Solidity: event WinningsClaimed(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchWinningsClaimed(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactoryWinningsClaimed) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "WinningsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactoryWinningsClaimed)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
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
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseWinningsClaimed(log types.Log) (*SportsLinkMarketFactoryWinningsClaimed, error) {
	event := new(SportsLinkMarketFactoryWinningsClaimed)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}
