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

// AbstractMarketFactoryMarket is an auto generated low-level Go binding around an user-defined struct.
/*
type AbstractMarketFactoryMarket struct {
	Creator     common.Address
	ShareTokens []common.Address
	EndTime     *big.Int
	Winner      common.Address
	CreatorFee  *big.Int
}
*/
// PriceMarketFactoryMarketDetails is an auto generated low-level Go binding around an user-defined struct.
type PriceMarketFactoryMarketDetails struct {
	SpotPrice         *big.Int
	ResolvedSpotPrice *big.Int
}

// PriceMarketFactoryABI is the input ABI used to generate the binding from.
const PriceMarketFactoryABI = "[{\"inputs\":[{\"internalType\":\"contractBPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Full\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Full\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Full\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shareFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractFeePot\",\"name\":\"_feePot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creatorFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"MarketResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"WinningsClaimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sharesToBurn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"burnShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"calcCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collateralIn\",\"type\":\"uint256\"}],\"name\":\"calcShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimManyWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"contractIERC20Full\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_spotPrice\",\"type\":\"uint256\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePot\",\"outputs\":[{\"internalType\":\"contractFeePot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMarket\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"contractOwnedERC20[]\",\"name\":\"shareTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"contractOwnedERC20\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"creatorFee\",\"type\":\"uint256\"}],\"internalType\":\"structAbstractMarketFactory.Market\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMarketDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resolvedSpotPrice\",\"type\":\"uint256\"}],\"internalType\":\"structPriceMarketFactory.MarketDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_shareToMint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mintShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"resolveMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shareFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PriceMarketFactory is an auto generated Go binding around an Ethereum contract.
type PriceMarketFactory struct {
	PriceMarketFactoryCaller     // Read-only binding to the contract
	PriceMarketFactoryTransactor // Write-only binding to the contract
	PriceMarketFactoryFilterer   // Log filterer for contract events
}

// PriceMarketFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceMarketFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceMarketFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceMarketFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceMarketFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceMarketFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceMarketFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceMarketFactorySession struct {
	Contract     *PriceMarketFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PriceMarketFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceMarketFactoryCallerSession struct {
	Contract *PriceMarketFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PriceMarketFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceMarketFactoryTransactorSession struct {
	Contract     *PriceMarketFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PriceMarketFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceMarketFactoryRaw struct {
	Contract *PriceMarketFactory // Generic contract binding to access the raw methods on
}

// PriceMarketFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceMarketFactoryCallerRaw struct {
	Contract *PriceMarketFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PriceMarketFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceMarketFactoryTransactorRaw struct {
	Contract *PriceMarketFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceMarketFactory creates a new instance of PriceMarketFactory, bound to a specific deployed contract.
func NewPriceMarketFactory(address common.Address, backend bind.ContractBackend) (*PriceMarketFactory, error) {
	contract, err := bindPriceMarketFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactory{PriceMarketFactoryCaller: PriceMarketFactoryCaller{contract: contract}, PriceMarketFactoryTransactor: PriceMarketFactoryTransactor{contract: contract}, PriceMarketFactoryFilterer: PriceMarketFactoryFilterer{contract: contract}}, nil
}

// NewPriceMarketFactoryCaller creates a new read-only instance of PriceMarketFactory, bound to a specific deployed contract.
func NewPriceMarketFactoryCaller(address common.Address, caller bind.ContractCaller) (*PriceMarketFactoryCaller, error) {
	contract, err := bindPriceMarketFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactoryCaller{contract: contract}, nil
}

// NewPriceMarketFactoryTransactor creates a new write-only instance of PriceMarketFactory, bound to a specific deployed contract.
func NewPriceMarketFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceMarketFactoryTransactor, error) {
	contract, err := bindPriceMarketFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactoryTransactor{contract: contract}, nil
}

// NewPriceMarketFactoryFilterer creates a new log filterer instance of PriceMarketFactory, bound to a specific deployed contract.
func NewPriceMarketFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceMarketFactoryFilterer, error) {
	contract, err := bindPriceMarketFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactoryFilterer{contract: contract}, nil
}

// bindPriceMarketFactory binds a generic wrapper to an already deployed contract.
func bindPriceMarketFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceMarketFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceMarketFactory *PriceMarketFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceMarketFactory.Contract.PriceMarketFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceMarketFactory *PriceMarketFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.PriceMarketFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceMarketFactory *PriceMarketFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.PriceMarketFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceMarketFactory *PriceMarketFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceMarketFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceMarketFactory *PriceMarketFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceMarketFactory *PriceMarketFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.contract.Transact(opts, method, params...)
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) CalcCost(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceMarketFactory.contract.Call(opts, out, "calcCost", _shares)
	return *ret0, err
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) CalcCost(_shares *big.Int) (*big.Int, error) {
	return _PriceMarketFactory.Contract.CalcCost(&_PriceMarketFactory.CallOpts, _shares)
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) CalcCost(_shares *big.Int) (*big.Int, error) {
	return _PriceMarketFactory.Contract.CalcCost(&_PriceMarketFactory.CallOpts, _shares)
}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) CalcShares(opts *bind.CallOpts, _collateralIn *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceMarketFactory.contract.Call(opts, out, "calcShares", _collateralIn)
	return *ret0, err
}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) CalcShares(_collateralIn *big.Int) (*big.Int, error) {
	return _PriceMarketFactory.Contract.CalcShares(&_PriceMarketFactory.CallOpts, _collateralIn)
}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) CalcShares(_collateralIn *big.Int) (*big.Int, error) {
	return _PriceMarketFactory.Contract.CalcShares(&_PriceMarketFactory.CallOpts, _collateralIn)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_PriceMarketFactory *PriceMarketFactoryCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PriceMarketFactory.contract.Call(opts, out, "collateral")
	return *ret0, err
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_PriceMarketFactory *PriceMarketFactorySession) Collateral() (common.Address, error) {
	return _PriceMarketFactory.Contract.Collateral(&_PriceMarketFactory.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) Collateral() (common.Address, error) {
	return _PriceMarketFactory.Contract.Collateral(&_PriceMarketFactory.CallOpts)
}

// CreatorFee is a free data retrieval call binding the contract method 0xe88958dc.
//
// Solidity: function creatorFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) CreatorFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceMarketFactory.contract.Call(opts, out, "creatorFee")
	return *ret0, err
}

// CreatorFee is a free data retrieval call binding the contract method 0xe88958dc.
//
// Solidity: function creatorFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) CreatorFee() (*big.Int, error) {
	return _PriceMarketFactory.Contract.CreatorFee(&_PriceMarketFactory.CallOpts)
}

// CreatorFee is a free data retrieval call binding the contract method 0xe88958dc.
//
// Solidity: function creatorFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) CreatorFee() (*big.Int, error) {
	return _PriceMarketFactory.Contract.CreatorFee(&_PriceMarketFactory.CallOpts)
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_PriceMarketFactory *PriceMarketFactoryCaller) FeePot(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PriceMarketFactory.contract.Call(opts, out, "feePot")
	return *ret0, err
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_PriceMarketFactory *PriceMarketFactorySession) FeePot() (common.Address, error) {
	return _PriceMarketFactory.Contract.FeePot(&_PriceMarketFactory.CallOpts)
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) FeePot() (common.Address, error) {
	return _PriceMarketFactory.Contract.FeePot(&_PriceMarketFactory.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns(AbstractMarketFactoryMarket)
func (_PriceMarketFactory *PriceMarketFactoryCaller) GetMarket(opts *bind.CallOpts, _id *big.Int) (AbstractMarketFactoryMarket, error) {
	var (
		ret0 = new(AbstractMarketFactoryMarket)
	)
	out := ret0
	err := _PriceMarketFactory.contract.Call(opts, out, "getMarket", _id)
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns(AbstractMarketFactoryMarket)
func (_PriceMarketFactory *PriceMarketFactorySession) GetMarket(_id *big.Int) (AbstractMarketFactoryMarket, error) {
	return _PriceMarketFactory.Contract.GetMarket(&_PriceMarketFactory.CallOpts, _id)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns(AbstractMarketFactoryMarket)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) GetMarket(_id *big.Int) (AbstractMarketFactoryMarket, error) {
	return _PriceMarketFactory.Contract.GetMarket(&_PriceMarketFactory.CallOpts, _id)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _id) view returns(PriceMarketFactoryMarketDetails)
func (_PriceMarketFactory *PriceMarketFactoryCaller) GetMarketDetails(opts *bind.CallOpts, _id *big.Int) (PriceMarketFactoryMarketDetails, error) {
	var (
		ret0 = new(PriceMarketFactoryMarketDetails)
	)
	out := ret0
	err := _PriceMarketFactory.contract.Call(opts, out, "getMarketDetails", _id)
	return *ret0, err
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _id) view returns(PriceMarketFactoryMarketDetails)
func (_PriceMarketFactory *PriceMarketFactorySession) GetMarketDetails(_id *big.Int) (PriceMarketFactoryMarketDetails, error) {
	return _PriceMarketFactory.Contract.GetMarketDetails(&_PriceMarketFactory.CallOpts, _id)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _id) view returns(PriceMarketFactoryMarketDetails)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) GetMarketDetails(_id *big.Int) (PriceMarketFactoryMarketDetails, error) {
	return _PriceMarketFactory.Contract.GetMarketDetails(&_PriceMarketFactory.CallOpts, _id)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceMarketFactory.contract.Call(opts, out, "marketCount")
	return *ret0, err
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) MarketCount() (*big.Int, error) {
	return _PriceMarketFactory.Contract.MarketCount(&_PriceMarketFactory.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) MarketCount() (*big.Int, error) {
	return _PriceMarketFactory.Contract.MarketCount(&_PriceMarketFactory.CallOpts)
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) ShareFactor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceMarketFactory.contract.Call(opts, out, "shareFactor")
	return *ret0, err
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) ShareFactor() (*big.Int, error) {
	return _PriceMarketFactory.Contract.ShareFactor(&_PriceMarketFactory.CallOpts)
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) ShareFactor() (*big.Int, error) {
	return _PriceMarketFactory.Contract.ShareFactor(&_PriceMarketFactory.CallOpts)
}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) StakerFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceMarketFactory.contract.Call(opts, out, "stakerFee")
	return *ret0, err
}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) StakerFee() (*big.Int, error) {
	return _PriceMarketFactory.Contract.StakerFee(&_PriceMarketFactory.CallOpts)
}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) StakerFee() (*big.Int, error) {
	return _PriceMarketFactory.Contract.StakerFee(&_PriceMarketFactory.CallOpts)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactor) BurnShares(opts *bind.TransactOpts, _id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "burnShares", _id, _sharesToBurn, _receiver)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) BurnShares(_id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.BurnShares(&_PriceMarketFactory.TransactOpts, _id, _sharesToBurn, _receiver)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) BurnShares(_id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.BurnShares(&_PriceMarketFactory.TransactOpts, _id, _sharesToBurn, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactor) ClaimManyWinnings(opts *bind.TransactOpts, _ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "claimManyWinnings", _ids, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) ClaimManyWinnings(_ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.ClaimManyWinnings(&_PriceMarketFactory.TransactOpts, _ids, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) ClaimManyWinnings(_ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.ClaimManyWinnings(&_PriceMarketFactory.TransactOpts, _ids, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactor) ClaimWinnings(opts *bind.TransactOpts, _id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "claimWinnings", _id, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) ClaimWinnings(_id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.ClaimWinnings(&_PriceMarketFactory.TransactOpts, _id, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) ClaimWinnings(_id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.ClaimWinnings(&_PriceMarketFactory.TransactOpts, _id, _receiver)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xbf605d8e.
//
// Solidity: function createMarket(address _creator, uint256 _endTime, uint256 _spotPrice) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactor) CreateMarket(opts *bind.TransactOpts, _creator common.Address, _endTime *big.Int, _spotPrice *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "createMarket", _creator, _endTime, _spotPrice)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xbf605d8e.
//
// Solidity: function createMarket(address _creator, uint256 _endTime, uint256 _spotPrice) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) CreateMarket(_creator common.Address, _endTime *big.Int, _spotPrice *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.CreateMarket(&_PriceMarketFactory.TransactOpts, _creator, _endTime, _spotPrice)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xbf605d8e.
//
// Solidity: function createMarket(address _creator, uint256 _endTime, uint256 _spotPrice) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) CreateMarket(_creator common.Address, _endTime *big.Int, _spotPrice *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.CreateMarket(&_PriceMarketFactory.TransactOpts, _creator, _endTime, _spotPrice)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactor) MintShares(opts *bind.TransactOpts, _id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "mintShares", _id, _shareToMint, _receiver)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_PriceMarketFactory *PriceMarketFactorySession) MintShares(_id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.MintShares(&_PriceMarketFactory.TransactOpts, _id, _shareToMint, _receiver)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) MintShares(_id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.MintShares(&_PriceMarketFactory.TransactOpts, _id, _shareToMint, _receiver)
}

// ResolveMarket is a paid mutator transaction binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 _id) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactor) ResolveMarket(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "resolveMarket", _id)
}

// ResolveMarket is a paid mutator transaction binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 _id) returns()
func (_PriceMarketFactory *PriceMarketFactorySession) ResolveMarket(_id *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.ResolveMarket(&_PriceMarketFactory.TransactOpts, _id)
}

// ResolveMarket is a paid mutator transaction binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 _id) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) ResolveMarket(_id *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.ResolveMarket(&_PriceMarketFactory.TransactOpts, _id)
}

// PriceMarketFactoryMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the PriceMarketFactory contract.
type PriceMarketFactoryMarketCreatedIterator struct {
	Event *PriceMarketFactoryMarketCreated // Event containing the contract specifics and raw log

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
func (it *PriceMarketFactoryMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceMarketFactoryMarketCreated)
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
		it.Event = new(PriceMarketFactoryMarketCreated)
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
func (it *PriceMarketFactoryMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceMarketFactoryMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceMarketFactoryMarketCreated represents a MarketCreated event raised by the PriceMarketFactory contract.
type PriceMarketFactoryMarketCreated struct {
	Id        *big.Int
	Creator   common.Address
	EndTime   *big.Int
	SpotPrice *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0x28c8de42a10b7bcc4a65ea3618bc8ada3e24cd7394886eae1b4f5f7440477080.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 endTime, uint256 spotPrice)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterMarketCreated(opts *bind.FilterOpts) (*PriceMarketFactoryMarketCreatedIterator, error) {

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "MarketCreated")
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactoryMarketCreatedIterator{contract: _PriceMarketFactory.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0x28c8de42a10b7bcc4a65ea3618bc8ada3e24cd7394886eae1b4f5f7440477080.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 endTime, uint256 spotPrice)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *PriceMarketFactoryMarketCreated) (event.Subscription, error) {

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "MarketCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceMarketFactoryMarketCreated)
				if err := _PriceMarketFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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

// ParseMarketCreated is a log parse operation binding the contract event 0x28c8de42a10b7bcc4a65ea3618bc8ada3e24cd7394886eae1b4f5f7440477080.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 endTime, uint256 spotPrice)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseMarketCreated(log types.Log) (*PriceMarketFactoryMarketCreated, error) {
	event := new(PriceMarketFactoryMarketCreated)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceMarketFactoryMarketResolvedIterator is returned from FilterMarketResolved and is used to iterate over the raw logs and unpacked data for MarketResolved events raised by the PriceMarketFactory contract.
type PriceMarketFactoryMarketResolvedIterator struct {
	Event *PriceMarketFactoryMarketResolved // Event containing the contract specifics and raw log

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
func (it *PriceMarketFactoryMarketResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceMarketFactoryMarketResolved)
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
		it.Event = new(PriceMarketFactoryMarketResolved)
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
func (it *PriceMarketFactoryMarketResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceMarketFactoryMarketResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceMarketFactoryMarketResolved represents a MarketResolved event raised by the PriceMarketFactory contract.
type PriceMarketFactoryMarketResolved struct {
	Id     *big.Int
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketResolved is a free log retrieval operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterMarketResolved(opts *bind.FilterOpts) (*PriceMarketFactoryMarketResolvedIterator, error) {

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactoryMarketResolvedIterator{contract: _PriceMarketFactory.contract, event: "MarketResolved", logs: logs, sub: sub}, nil
}

// WatchMarketResolved is a free log subscription operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchMarketResolved(opts *bind.WatchOpts, sink chan<- *PriceMarketFactoryMarketResolved) (event.Subscription, error) {

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceMarketFactoryMarketResolved)
				if err := _PriceMarketFactory.contract.UnpackLog(event, "MarketResolved", log); err != nil {
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
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseMarketResolved(log types.Log) (*PriceMarketFactoryMarketResolved, error) {
	event := new(PriceMarketFactoryMarketResolved)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "MarketResolved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceMarketFactorySharesBurnedIterator is returned from FilterSharesBurned and is used to iterate over the raw logs and unpacked data for SharesBurned events raised by the PriceMarketFactory contract.
type PriceMarketFactorySharesBurnedIterator struct {
	Event *PriceMarketFactorySharesBurned // Event containing the contract specifics and raw log

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
func (it *PriceMarketFactorySharesBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceMarketFactorySharesBurned)
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
		it.Event = new(PriceMarketFactorySharesBurned)
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
func (it *PriceMarketFactorySharesBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceMarketFactorySharesBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceMarketFactorySharesBurned represents a SharesBurned event raised by the PriceMarketFactory contract.
type PriceMarketFactorySharesBurned struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSharesBurned is a free log retrieval operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterSharesBurned(opts *bind.FilterOpts) (*PriceMarketFactorySharesBurnedIterator, error) {

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "SharesBurned")
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactorySharesBurnedIterator{contract: _PriceMarketFactory.contract, event: "SharesBurned", logs: logs, sub: sub}, nil
}

// WatchSharesBurned is a free log subscription operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchSharesBurned(opts *bind.WatchOpts, sink chan<- *PriceMarketFactorySharesBurned) (event.Subscription, error) {

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "SharesBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceMarketFactorySharesBurned)
				if err := _PriceMarketFactory.contract.UnpackLog(event, "SharesBurned", log); err != nil {
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
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseSharesBurned(log types.Log) (*PriceMarketFactorySharesBurned, error) {
	event := new(PriceMarketFactorySharesBurned)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "SharesBurned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceMarketFactorySharesMintedIterator is returned from FilterSharesMinted and is used to iterate over the raw logs and unpacked data for SharesMinted events raised by the PriceMarketFactory contract.
type PriceMarketFactorySharesMintedIterator struct {
	Event *PriceMarketFactorySharesMinted // Event containing the contract specifics and raw log

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
func (it *PriceMarketFactorySharesMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceMarketFactorySharesMinted)
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
		it.Event = new(PriceMarketFactorySharesMinted)
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
func (it *PriceMarketFactorySharesMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceMarketFactorySharesMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceMarketFactorySharesMinted represents a SharesMinted event raised by the PriceMarketFactory contract.
type PriceMarketFactorySharesMinted struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSharesMinted is a free log retrieval operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterSharesMinted(opts *bind.FilterOpts) (*PriceMarketFactorySharesMintedIterator, error) {

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "SharesMinted")
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactorySharesMintedIterator{contract: _PriceMarketFactory.contract, event: "SharesMinted", logs: logs, sub: sub}, nil
}

// WatchSharesMinted is a free log subscription operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchSharesMinted(opts *bind.WatchOpts, sink chan<- *PriceMarketFactorySharesMinted) (event.Subscription, error) {

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "SharesMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceMarketFactorySharesMinted)
				if err := _PriceMarketFactory.contract.UnpackLog(event, "SharesMinted", log); err != nil {
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
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseSharesMinted(log types.Log) (*PriceMarketFactorySharesMinted, error) {
	event := new(PriceMarketFactorySharesMinted)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "SharesMinted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceMarketFactoryWinningsClaimedIterator is returned from FilterWinningsClaimed and is used to iterate over the raw logs and unpacked data for WinningsClaimed events raised by the PriceMarketFactory contract.
type PriceMarketFactoryWinningsClaimedIterator struct {
	Event *PriceMarketFactoryWinningsClaimed // Event containing the contract specifics and raw log

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
func (it *PriceMarketFactoryWinningsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceMarketFactoryWinningsClaimed)
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
		it.Event = new(PriceMarketFactoryWinningsClaimed)
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
func (it *PriceMarketFactoryWinningsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceMarketFactoryWinningsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceMarketFactoryWinningsClaimed represents a WinningsClaimed event raised by the PriceMarketFactory contract.
type PriceMarketFactoryWinningsClaimed struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWinningsClaimed is a free log retrieval operation binding the contract event 0x2bdd7a5109193ce6270ec3b4afcf4ccd4a06c27742ba11f660498cb41433bb00.
//
// Solidity: event WinningsClaimed(uint256 id, uint256 amount, address receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterWinningsClaimed(opts *bind.FilterOpts) (*PriceMarketFactoryWinningsClaimedIterator, error) {

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "WinningsClaimed")
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactoryWinningsClaimedIterator{contract: _PriceMarketFactory.contract, event: "WinningsClaimed", logs: logs, sub: sub}, nil
}

// WatchWinningsClaimed is a free log subscription operation binding the contract event 0x2bdd7a5109193ce6270ec3b4afcf4ccd4a06c27742ba11f660498cb41433bb00.
//
// Solidity: event WinningsClaimed(uint256 id, uint256 amount, address receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchWinningsClaimed(opts *bind.WatchOpts, sink chan<- *PriceMarketFactoryWinningsClaimed) (event.Subscription, error) {

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "WinningsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceMarketFactoryWinningsClaimed)
				if err := _PriceMarketFactory.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
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
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseWinningsClaimed(log types.Log) (*PriceMarketFactoryWinningsClaimed, error) {
	event := new(PriceMarketFactoryWinningsClaimed)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}
