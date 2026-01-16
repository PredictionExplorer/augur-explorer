// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// AbstractMarketFactoryMarket is an auto generated low-level Go binding around an user-defined struct.
type AbstractMarketFactoryMarket struct {
	SettlementAddress common.Address
	ShareTokens       []common.Address
	EndTime           *big.Int
	Winner            common.Address
	SettlementFee     *big.Int
	ProtocolFee       *big.Int
	StakerFee         *big.Int
	CreationTimestamp *big.Int
}

// TestPriceMarketFactoryMarketDetails is an auto generated low-level Go binding around an user-defined struct.
type TestPriceMarketFactoryMarketDetails struct {
	SpotPrice         *big.Int
	ResolvedSpotPrice *big.Int
}

// PriceMarketFactoryMetaData contains all meta data concerning the PriceMarketFactory contract.
var PriceMarketFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Full\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shareFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractFeePot\",\"name\":\"_feePot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_settlementFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_protocol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"contractBPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Full\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Full\",\"name\":\"_tokenOut\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"MarketResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"}],\"name\":\"ProtocolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SettlementFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"settlementAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SettlementFeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"StakerFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winningOutcome\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"WinningsClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accumulatedProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accumulatedSettlementFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sharesToBurn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"burnShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"calcCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collateralIn\",\"type\":\"uint256\"}],\"name\":\"calcShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimManyWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimProtocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimSettlementFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"contractIERC20Full\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_spotPrice\",\"type\":\"uint256\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePot\",\"outputs\":[{\"internalType\":\"contractFeePot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMarket\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"settlementAddress\",\"type\":\"address\"},{\"internalType\":\"contractOwnedERC20[]\",\"name\":\"shareTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"contractOwnedERC20\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"settlementFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structAbstractMarketFactory.Market\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMarketDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resolvedSpotPrice\",\"type\":\"uint256\"}],\"internalType\":\"structTestPriceMarketFactory.MarketDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isMarketResolved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listUnresolvedMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_shareToMint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mintShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocol\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"resolveMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newProtocol\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_claimFirst\",\"type\":\"bool\"}],\"name\":\"setProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setSettlementFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setStakerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shareFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PriceMarketFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceMarketFactoryMetaData.ABI instead.
var PriceMarketFactoryABI = PriceMarketFactoryMetaData.ABI

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
func (_PriceMarketFactory *PriceMarketFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_PriceMarketFactory *PriceMarketFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// AccumulatedProtocolFee is a free data retrieval call binding the contract method 0xa544a62c.
//
// Solidity: function accumulatedProtocolFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) AccumulatedProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "accumulatedProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedProtocolFee is a free data retrieval call binding the contract method 0xa544a62c.
//
// Solidity: function accumulatedProtocolFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) AccumulatedProtocolFee() (*big.Int, error) {
	return _PriceMarketFactory.Contract.AccumulatedProtocolFee(&_PriceMarketFactory.CallOpts)
}

// AccumulatedProtocolFee is a free data retrieval call binding the contract method 0xa544a62c.
//
// Solidity: function accumulatedProtocolFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) AccumulatedProtocolFee() (*big.Int, error) {
	return _PriceMarketFactory.Contract.AccumulatedProtocolFee(&_PriceMarketFactory.CallOpts)
}

// AccumulatedSettlementFees is a free data retrieval call binding the contract method 0x49a4d934.
//
// Solidity: function accumulatedSettlementFees(address ) view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) AccumulatedSettlementFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "accumulatedSettlementFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedSettlementFees is a free data retrieval call binding the contract method 0x49a4d934.
//
// Solidity: function accumulatedSettlementFees(address ) view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) AccumulatedSettlementFees(arg0 common.Address) (*big.Int, error) {
	return _PriceMarketFactory.Contract.AccumulatedSettlementFees(&_PriceMarketFactory.CallOpts, arg0)
}

// AccumulatedSettlementFees is a free data retrieval call binding the contract method 0x49a4d934.
//
// Solidity: function accumulatedSettlementFees(address ) view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) AccumulatedSettlementFees(arg0 common.Address) (*big.Int, error) {
	return _PriceMarketFactory.Contract.AccumulatedSettlementFees(&_PriceMarketFactory.CallOpts, arg0)
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) CalcCost(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "calcCost", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "calcShares", _collateralIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_PriceMarketFactory *PriceMarketFactoryCaller) FeePot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "feePot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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
// Solidity: function getMarket(uint256 _id) view returns((address,address[],uint256,address,uint256,uint256,uint256,uint256))
func (_PriceMarketFactory *PriceMarketFactoryCaller) GetMarket(opts *bind.CallOpts, _id *big.Int) (AbstractMarketFactoryMarket, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "getMarket", _id)

	if err != nil {
		return *new(AbstractMarketFactoryMarket), err
	}

	out0 := *abi.ConvertType(out[0], new(AbstractMarketFactoryMarket)).(*AbstractMarketFactoryMarket)

	return out0, err

}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns((address,address[],uint256,address,uint256,uint256,uint256,uint256))
func (_PriceMarketFactory *PriceMarketFactorySession) GetMarket(_id *big.Int) (AbstractMarketFactoryMarket, error) {
	return _PriceMarketFactory.Contract.GetMarket(&_PriceMarketFactory.CallOpts, _id)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns((address,address[],uint256,address,uint256,uint256,uint256,uint256))
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) GetMarket(_id *big.Int) (AbstractMarketFactoryMarket, error) {
	return _PriceMarketFactory.Contract.GetMarket(&_PriceMarketFactory.CallOpts, _id)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _id) view returns((uint256,uint256))
func (_PriceMarketFactory *PriceMarketFactoryCaller) GetMarketDetails(opts *bind.CallOpts, _id *big.Int) (TestPriceMarketFactoryMarketDetails, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "getMarketDetails", _id)

	if err != nil {
		return *new(TestPriceMarketFactoryMarketDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(TestPriceMarketFactoryMarketDetails)).(*TestPriceMarketFactoryMarketDetails)

	return out0, err

}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _id) view returns((uint256,uint256))
func (_PriceMarketFactory *PriceMarketFactorySession) GetMarketDetails(_id *big.Int) (TestPriceMarketFactoryMarketDetails, error) {
	return _PriceMarketFactory.Contract.GetMarketDetails(&_PriceMarketFactory.CallOpts, _id)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _id) view returns((uint256,uint256))
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) GetMarketDetails(_id *big.Int) (TestPriceMarketFactoryMarketDetails, error) {
	return _PriceMarketFactory.Contract.GetMarketDetails(&_PriceMarketFactory.CallOpts, _id)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_PriceMarketFactory *PriceMarketFactoryCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_PriceMarketFactory *PriceMarketFactorySession) GetOwner() (common.Address, error) {
	return _PriceMarketFactory.Contract.GetOwner(&_PriceMarketFactory.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) GetOwner() (common.Address, error) {
	return _PriceMarketFactory.Contract.GetOwner(&_PriceMarketFactory.CallOpts)
}

// IsMarketResolved is a free data retrieval call binding the contract method 0x53ac55f5.
//
// Solidity: function isMarketResolved(uint256 _id) view returns(bool)
func (_PriceMarketFactory *PriceMarketFactoryCaller) IsMarketResolved(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "isMarketResolved", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketResolved is a free data retrieval call binding the contract method 0x53ac55f5.
//
// Solidity: function isMarketResolved(uint256 _id) view returns(bool)
func (_PriceMarketFactory *PriceMarketFactorySession) IsMarketResolved(_id *big.Int) (bool, error) {
	return _PriceMarketFactory.Contract.IsMarketResolved(&_PriceMarketFactory.CallOpts, _id)
}

// IsMarketResolved is a free data retrieval call binding the contract method 0x53ac55f5.
//
// Solidity: function isMarketResolved(uint256 _id) view returns(bool)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) IsMarketResolved(_id *big.Int) (bool, error) {
	return _PriceMarketFactory.Contract.IsMarketResolved(&_PriceMarketFactory.CallOpts, _id)
}

// ListUnresolvedMarkets is a free data retrieval call binding the contract method 0xd9113f0d.
//
// Solidity: function listUnresolvedMarkets() view returns(uint256[])
func (_PriceMarketFactory *PriceMarketFactoryCaller) ListUnresolvedMarkets(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "listUnresolvedMarkets")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ListUnresolvedMarkets is a free data retrieval call binding the contract method 0xd9113f0d.
//
// Solidity: function listUnresolvedMarkets() view returns(uint256[])
func (_PriceMarketFactory *PriceMarketFactorySession) ListUnresolvedMarkets() ([]*big.Int, error) {
	return _PriceMarketFactory.Contract.ListUnresolvedMarkets(&_PriceMarketFactory.CallOpts)
}

// ListUnresolvedMarkets is a free data retrieval call binding the contract method 0xd9113f0d.
//
// Solidity: function listUnresolvedMarkets() view returns(uint256[])
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) ListUnresolvedMarkets() ([]*big.Int, error) {
	return _PriceMarketFactory.Contract.ListUnresolvedMarkets(&_PriceMarketFactory.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "marketCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(address)
func (_PriceMarketFactory *PriceMarketFactoryCaller) Protocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "protocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(address)
func (_PriceMarketFactory *PriceMarketFactorySession) Protocol() (common.Address, error) {
	return _PriceMarketFactory.Contract.Protocol(&_PriceMarketFactory.CallOpts)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(address)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) Protocol() (common.Address, error) {
	return _PriceMarketFactory.Contract.Protocol(&_PriceMarketFactory.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) ProtocolFee() (*big.Int, error) {
	return _PriceMarketFactory.Contract.ProtocolFee(&_PriceMarketFactory.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) ProtocolFee() (*big.Int, error) {
	return _PriceMarketFactory.Contract.ProtocolFee(&_PriceMarketFactory.CallOpts)
}

// SettlementFee is a free data retrieval call binding the contract method 0x7d1d7fb8.
//
// Solidity: function settlementFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) SettlementFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "settlementFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SettlementFee is a free data retrieval call binding the contract method 0x7d1d7fb8.
//
// Solidity: function settlementFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) SettlementFee() (*big.Int, error) {
	return _PriceMarketFactory.Contract.SettlementFee(&_PriceMarketFactory.CallOpts)
}

// SettlementFee is a free data retrieval call binding the contract method 0x7d1d7fb8.
//
// Solidity: function settlementFee() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCallerSession) SettlementFee() (*big.Int, error) {
	return _PriceMarketFactory.Contract.SettlementFee(&_PriceMarketFactory.CallOpts)
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryCaller) ShareFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "shareFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _PriceMarketFactory.contract.Call(opts, &out, "stakerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x4a7d0369.
//
// Solidity: function claimProtocolFees() returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactor) ClaimProtocolFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "claimProtocolFees")
}

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x4a7d0369.
//
// Solidity: function claimProtocolFees() returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) ClaimProtocolFees() (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.ClaimProtocolFees(&_PriceMarketFactory.TransactOpts)
}

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x4a7d0369.
//
// Solidity: function claimProtocolFees() returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) ClaimProtocolFees() (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.ClaimProtocolFees(&_PriceMarketFactory.TransactOpts)
}

// ClaimSettlementFees is a paid mutator transaction binding the contract method 0x8e0ed193.
//
// Solidity: function claimSettlementFees(address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactor) ClaimSettlementFees(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "claimSettlementFees", _receiver)
}

// ClaimSettlementFees is a paid mutator transaction binding the contract method 0x8e0ed193.
//
// Solidity: function claimSettlementFees(address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactorySession) ClaimSettlementFees(_receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.ClaimSettlementFees(&_PriceMarketFactory.TransactOpts, _receiver)
}

// ClaimSettlementFees is a paid mutator transaction binding the contract method 0x8e0ed193.
//
// Solidity: function claimSettlementFees(address _receiver) returns(uint256)
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) ClaimSettlementFees(_receiver common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.ClaimSettlementFees(&_PriceMarketFactory.TransactOpts, _receiver)
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

// SetProtocol is a paid mutator transaction binding the contract method 0x32ecabe9.
//
// Solidity: function setProtocol(address _newProtocol, bool _claimFirst) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactor) SetProtocol(opts *bind.TransactOpts, _newProtocol common.Address, _claimFirst bool) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "setProtocol", _newProtocol, _claimFirst)
}

// SetProtocol is a paid mutator transaction binding the contract method 0x32ecabe9.
//
// Solidity: function setProtocol(address _newProtocol, bool _claimFirst) returns()
func (_PriceMarketFactory *PriceMarketFactorySession) SetProtocol(_newProtocol common.Address, _claimFirst bool) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.SetProtocol(&_PriceMarketFactory.TransactOpts, _newProtocol, _claimFirst)
}

// SetProtocol is a paid mutator transaction binding the contract method 0x32ecabe9.
//
// Solidity: function setProtocol(address _newProtocol, bool _claimFirst) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) SetProtocol(_newProtocol common.Address, _claimFirst bool) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.SetProtocol(&_PriceMarketFactory.TransactOpts, _newProtocol, _claimFirst)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _newFee) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactor) SetProtocolFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "setProtocolFee", _newFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _newFee) returns()
func (_PriceMarketFactory *PriceMarketFactorySession) SetProtocolFee(_newFee *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.SetProtocolFee(&_PriceMarketFactory.TransactOpts, _newFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _newFee) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) SetProtocolFee(_newFee *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.SetProtocolFee(&_PriceMarketFactory.TransactOpts, _newFee)
}

// SetSettlementFee is a paid mutator transaction binding the contract method 0xd5da4f1d.
//
// Solidity: function setSettlementFee(uint256 _newFee) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactor) SetSettlementFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "setSettlementFee", _newFee)
}

// SetSettlementFee is a paid mutator transaction binding the contract method 0xd5da4f1d.
//
// Solidity: function setSettlementFee(uint256 _newFee) returns()
func (_PriceMarketFactory *PriceMarketFactorySession) SetSettlementFee(_newFee *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.SetSettlementFee(&_PriceMarketFactory.TransactOpts, _newFee)
}

// SetSettlementFee is a paid mutator transaction binding the contract method 0xd5da4f1d.
//
// Solidity: function setSettlementFee(uint256 _newFee) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) SetSettlementFee(_newFee *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.SetSettlementFee(&_PriceMarketFactory.TransactOpts, _newFee)
}

// SetStakerFee is a paid mutator transaction binding the contract method 0x97eef187.
//
// Solidity: function setStakerFee(uint256 _newFee) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactor) SetStakerFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "setStakerFee", _newFee)
}

// SetStakerFee is a paid mutator transaction binding the contract method 0x97eef187.
//
// Solidity: function setStakerFee(uint256 _newFee) returns()
func (_PriceMarketFactory *PriceMarketFactorySession) SetStakerFee(_newFee *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.SetStakerFee(&_PriceMarketFactory.TransactOpts, _newFee)
}

// SetStakerFee is a paid mutator transaction binding the contract method 0x97eef187.
//
// Solidity: function setStakerFee(uint256 _newFee) returns()
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) SetStakerFee(_newFee *big.Int) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.SetStakerFee(&_PriceMarketFactory.TransactOpts, _newFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_PriceMarketFactory *PriceMarketFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_PriceMarketFactory *PriceMarketFactorySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.TransferOwnership(&_PriceMarketFactory.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_PriceMarketFactory *PriceMarketFactoryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PriceMarketFactory.Contract.TransferOwnership(&_PriceMarketFactory.TransactOpts, _newOwner)
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
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// PriceMarketFactoryProtocolChangedIterator is returned from FilterProtocolChanged and is used to iterate over the raw logs and unpacked data for ProtocolChanged events raised by the PriceMarketFactory contract.
type PriceMarketFactoryProtocolChangedIterator struct {
	Event *PriceMarketFactoryProtocolChanged // Event containing the contract specifics and raw log

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
func (it *PriceMarketFactoryProtocolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceMarketFactoryProtocolChanged)
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
		it.Event = new(PriceMarketFactoryProtocolChanged)
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
func (it *PriceMarketFactoryProtocolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceMarketFactoryProtocolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceMarketFactoryProtocolChanged represents a ProtocolChanged event raised by the PriceMarketFactory contract.
type PriceMarketFactoryProtocolChanged struct {
	Protocol common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProtocolChanged is a free log retrieval operation binding the contract event 0x15b84596b3c567ae2998116949ae5f2d47f3055c12d9053db4d6e50f4c794dd9.
//
// Solidity: event ProtocolChanged(address protocol)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterProtocolChanged(opts *bind.FilterOpts) (*PriceMarketFactoryProtocolChangedIterator, error) {

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "ProtocolChanged")
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactoryProtocolChangedIterator{contract: _PriceMarketFactory.contract, event: "ProtocolChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolChanged is a free log subscription operation binding the contract event 0x15b84596b3c567ae2998116949ae5f2d47f3055c12d9053db4d6e50f4c794dd9.
//
// Solidity: event ProtocolChanged(address protocol)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchProtocolChanged(opts *bind.WatchOpts, sink chan<- *PriceMarketFactoryProtocolChanged) (event.Subscription, error) {

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "ProtocolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceMarketFactoryProtocolChanged)
				if err := _PriceMarketFactory.contract.UnpackLog(event, "ProtocolChanged", log); err != nil {
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

// ParseProtocolChanged is a log parse operation binding the contract event 0x15b84596b3c567ae2998116949ae5f2d47f3055c12d9053db4d6e50f4c794dd9.
//
// Solidity: event ProtocolChanged(address protocol)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseProtocolChanged(log types.Log) (*PriceMarketFactoryProtocolChanged, error) {
	event := new(PriceMarketFactoryProtocolChanged)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "ProtocolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceMarketFactoryProtocolFeeChangedIterator is returned from FilterProtocolFeeChanged and is used to iterate over the raw logs and unpacked data for ProtocolFeeChanged events raised by the PriceMarketFactory contract.
type PriceMarketFactoryProtocolFeeChangedIterator struct {
	Event *PriceMarketFactoryProtocolFeeChanged // Event containing the contract specifics and raw log

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
func (it *PriceMarketFactoryProtocolFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceMarketFactoryProtocolFeeChanged)
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
		it.Event = new(PriceMarketFactoryProtocolFeeChanged)
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
func (it *PriceMarketFactoryProtocolFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceMarketFactoryProtocolFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceMarketFactoryProtocolFeeChanged represents a ProtocolFeeChanged event raised by the PriceMarketFactory contract.
type PriceMarketFactoryProtocolFeeChanged struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeChanged is a free log retrieval operation binding the contract event 0xada2cde3c4a561f5c23e2fdbfb223e1f0d1ec7109b9811b32644e6e974d6631f.
//
// Solidity: event ProtocolFeeChanged(uint256 fee)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterProtocolFeeChanged(opts *bind.FilterOpts) (*PriceMarketFactoryProtocolFeeChangedIterator, error) {

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "ProtocolFeeChanged")
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactoryProtocolFeeChangedIterator{contract: _PriceMarketFactory.contract, event: "ProtocolFeeChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeChanged is a free log subscription operation binding the contract event 0xada2cde3c4a561f5c23e2fdbfb223e1f0d1ec7109b9811b32644e6e974d6631f.
//
// Solidity: event ProtocolFeeChanged(uint256 fee)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchProtocolFeeChanged(opts *bind.WatchOpts, sink chan<- *PriceMarketFactoryProtocolFeeChanged) (event.Subscription, error) {

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "ProtocolFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceMarketFactoryProtocolFeeChanged)
				if err := _PriceMarketFactory.contract.UnpackLog(event, "ProtocolFeeChanged", log); err != nil {
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

// ParseProtocolFeeChanged is a log parse operation binding the contract event 0xada2cde3c4a561f5c23e2fdbfb223e1f0d1ec7109b9811b32644e6e974d6631f.
//
// Solidity: event ProtocolFeeChanged(uint256 fee)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseProtocolFeeChanged(log types.Log) (*PriceMarketFactoryProtocolFeeChanged, error) {
	event := new(PriceMarketFactoryProtocolFeeChanged)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "ProtocolFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceMarketFactoryProtocolFeeClaimedIterator is returned from FilterProtocolFeeClaimed and is used to iterate over the raw logs and unpacked data for ProtocolFeeClaimed events raised by the PriceMarketFactory contract.
type PriceMarketFactoryProtocolFeeClaimedIterator struct {
	Event *PriceMarketFactoryProtocolFeeClaimed // Event containing the contract specifics and raw log

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
func (it *PriceMarketFactoryProtocolFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceMarketFactoryProtocolFeeClaimed)
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
		it.Event = new(PriceMarketFactoryProtocolFeeClaimed)
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
func (it *PriceMarketFactoryProtocolFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceMarketFactoryProtocolFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceMarketFactoryProtocolFeeClaimed represents a ProtocolFeeClaimed event raised by the PriceMarketFactory contract.
type PriceMarketFactoryProtocolFeeClaimed struct {
	Protocol common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeClaimed is a free log retrieval operation binding the contract event 0x0f7f5b155b0b0ac6890709a2c7bf1b8bb3f675fff1e7840b4dd3c9acde59048b.
//
// Solidity: event ProtocolFeeClaimed(address protocol, uint256 amount)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterProtocolFeeClaimed(opts *bind.FilterOpts) (*PriceMarketFactoryProtocolFeeClaimedIterator, error) {

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "ProtocolFeeClaimed")
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactoryProtocolFeeClaimedIterator{contract: _PriceMarketFactory.contract, event: "ProtocolFeeClaimed", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeClaimed is a free log subscription operation binding the contract event 0x0f7f5b155b0b0ac6890709a2c7bf1b8bb3f675fff1e7840b4dd3c9acde59048b.
//
// Solidity: event ProtocolFeeClaimed(address protocol, uint256 amount)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchProtocolFeeClaimed(opts *bind.WatchOpts, sink chan<- *PriceMarketFactoryProtocolFeeClaimed) (event.Subscription, error) {

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "ProtocolFeeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceMarketFactoryProtocolFeeClaimed)
				if err := _PriceMarketFactory.contract.UnpackLog(event, "ProtocolFeeClaimed", log); err != nil {
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

// ParseProtocolFeeClaimed is a log parse operation binding the contract event 0x0f7f5b155b0b0ac6890709a2c7bf1b8bb3f675fff1e7840b4dd3c9acde59048b.
//
// Solidity: event ProtocolFeeClaimed(address protocol, uint256 amount)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseProtocolFeeClaimed(log types.Log) (*PriceMarketFactoryProtocolFeeClaimed, error) {
	event := new(PriceMarketFactoryProtocolFeeClaimed)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "ProtocolFeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceMarketFactorySettlementFeeChangedIterator is returned from FilterSettlementFeeChanged and is used to iterate over the raw logs and unpacked data for SettlementFeeChanged events raised by the PriceMarketFactory contract.
type PriceMarketFactorySettlementFeeChangedIterator struct {
	Event *PriceMarketFactorySettlementFeeChanged // Event containing the contract specifics and raw log

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
func (it *PriceMarketFactorySettlementFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceMarketFactorySettlementFeeChanged)
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
		it.Event = new(PriceMarketFactorySettlementFeeChanged)
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
func (it *PriceMarketFactorySettlementFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceMarketFactorySettlementFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceMarketFactorySettlementFeeChanged represents a SettlementFeeChanged event raised by the PriceMarketFactory contract.
type PriceMarketFactorySettlementFeeChanged struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSettlementFeeChanged is a free log retrieval operation binding the contract event 0x92d395c429898992f8532ee7145901513e524c2085fd7fd1da39b8badcd6df31.
//
// Solidity: event SettlementFeeChanged(uint256 fee)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterSettlementFeeChanged(opts *bind.FilterOpts) (*PriceMarketFactorySettlementFeeChangedIterator, error) {

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "SettlementFeeChanged")
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactorySettlementFeeChangedIterator{contract: _PriceMarketFactory.contract, event: "SettlementFeeChanged", logs: logs, sub: sub}, nil
}

// WatchSettlementFeeChanged is a free log subscription operation binding the contract event 0x92d395c429898992f8532ee7145901513e524c2085fd7fd1da39b8badcd6df31.
//
// Solidity: event SettlementFeeChanged(uint256 fee)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchSettlementFeeChanged(opts *bind.WatchOpts, sink chan<- *PriceMarketFactorySettlementFeeChanged) (event.Subscription, error) {

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "SettlementFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceMarketFactorySettlementFeeChanged)
				if err := _PriceMarketFactory.contract.UnpackLog(event, "SettlementFeeChanged", log); err != nil {
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

// ParseSettlementFeeChanged is a log parse operation binding the contract event 0x92d395c429898992f8532ee7145901513e524c2085fd7fd1da39b8badcd6df31.
//
// Solidity: event SettlementFeeChanged(uint256 fee)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseSettlementFeeChanged(log types.Log) (*PriceMarketFactorySettlementFeeChanged, error) {
	event := new(PriceMarketFactorySettlementFeeChanged)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "SettlementFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceMarketFactorySettlementFeeClaimedIterator is returned from FilterSettlementFeeClaimed and is used to iterate over the raw logs and unpacked data for SettlementFeeClaimed events raised by the PriceMarketFactory contract.
type PriceMarketFactorySettlementFeeClaimedIterator struct {
	Event *PriceMarketFactorySettlementFeeClaimed // Event containing the contract specifics and raw log

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
func (it *PriceMarketFactorySettlementFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceMarketFactorySettlementFeeClaimed)
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
		it.Event = new(PriceMarketFactorySettlementFeeClaimed)
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
func (it *PriceMarketFactorySettlementFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceMarketFactorySettlementFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceMarketFactorySettlementFeeClaimed represents a SettlementFeeClaimed event raised by the PriceMarketFactory contract.
type PriceMarketFactorySettlementFeeClaimed struct {
	SettlementAddress common.Address
	Amount            *big.Int
	Receiver          common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSettlementFeeClaimed is a free log retrieval operation binding the contract event 0xc9985ad824d943d66367ce5feea26e18979b3e1c9273742926d87e2b0d747387.
//
// Solidity: event SettlementFeeClaimed(address settlementAddress, uint256 amount, address indexed receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterSettlementFeeClaimed(opts *bind.FilterOpts, receiver []common.Address) (*PriceMarketFactorySettlementFeeClaimedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "SettlementFeeClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactorySettlementFeeClaimedIterator{contract: _PriceMarketFactory.contract, event: "SettlementFeeClaimed", logs: logs, sub: sub}, nil
}

// WatchSettlementFeeClaimed is a free log subscription operation binding the contract event 0xc9985ad824d943d66367ce5feea26e18979b3e1c9273742926d87e2b0d747387.
//
// Solidity: event SettlementFeeClaimed(address settlementAddress, uint256 amount, address indexed receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchSettlementFeeClaimed(opts *bind.WatchOpts, sink chan<- *PriceMarketFactorySettlementFeeClaimed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "SettlementFeeClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceMarketFactorySettlementFeeClaimed)
				if err := _PriceMarketFactory.contract.UnpackLog(event, "SettlementFeeClaimed", log); err != nil {
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

// ParseSettlementFeeClaimed is a log parse operation binding the contract event 0xc9985ad824d943d66367ce5feea26e18979b3e1c9273742926d87e2b0d747387.
//
// Solidity: event SettlementFeeClaimed(address settlementAddress, uint256 amount, address indexed receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseSettlementFeeClaimed(log types.Log) (*PriceMarketFactorySettlementFeeClaimed, error) {
	event := new(PriceMarketFactorySettlementFeeClaimed)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "SettlementFeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// PriceMarketFactoryStakerFeeChangedIterator is returned from FilterStakerFeeChanged and is used to iterate over the raw logs and unpacked data for StakerFeeChanged events raised by the PriceMarketFactory contract.
type PriceMarketFactoryStakerFeeChangedIterator struct {
	Event *PriceMarketFactoryStakerFeeChanged // Event containing the contract specifics and raw log

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
func (it *PriceMarketFactoryStakerFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceMarketFactoryStakerFeeChanged)
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
		it.Event = new(PriceMarketFactoryStakerFeeChanged)
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
func (it *PriceMarketFactoryStakerFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceMarketFactoryStakerFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceMarketFactoryStakerFeeChanged represents a StakerFeeChanged event raised by the PriceMarketFactory contract.
type PriceMarketFactoryStakerFeeChanged struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakerFeeChanged is a free log retrieval operation binding the contract event 0xcc4df50442ac32f0142ba4853f617661e0823be0e92148e7e5f36ce56c139825.
//
// Solidity: event StakerFeeChanged(uint256 fee)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterStakerFeeChanged(opts *bind.FilterOpts) (*PriceMarketFactoryStakerFeeChangedIterator, error) {

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "StakerFeeChanged")
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactoryStakerFeeChangedIterator{contract: _PriceMarketFactory.contract, event: "StakerFeeChanged", logs: logs, sub: sub}, nil
}

// WatchStakerFeeChanged is a free log subscription operation binding the contract event 0xcc4df50442ac32f0142ba4853f617661e0823be0e92148e7e5f36ce56c139825.
//
// Solidity: event StakerFeeChanged(uint256 fee)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchStakerFeeChanged(opts *bind.WatchOpts, sink chan<- *PriceMarketFactoryStakerFeeChanged) (event.Subscription, error) {

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "StakerFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceMarketFactoryStakerFeeChanged)
				if err := _PriceMarketFactory.contract.UnpackLog(event, "StakerFeeChanged", log); err != nil {
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

// ParseStakerFeeChanged is a log parse operation binding the contract event 0xcc4df50442ac32f0142ba4853f617661e0823be0e92148e7e5f36ce56c139825.
//
// Solidity: event StakerFeeChanged(uint256 fee)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseStakerFeeChanged(log types.Log) (*PriceMarketFactoryStakerFeeChanged, error) {
	event := new(PriceMarketFactoryStakerFeeChanged)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "StakerFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	Id             *big.Int
	WinningOutcome common.Address
	Amount         *big.Int
	SettlementFee  *big.Int
	Payout         *big.Int
	Receiver       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWinningsClaimed is a free log retrieval operation binding the contract event 0xe67bd0100cd3289557430d36098901ba18161e6279c9711d8650b8af10552104.
//
// Solidity: event WinningsClaimed(uint256 id, address winningOutcome, uint256 amount, uint256 settlementFee, uint256 payout, address indexed receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) FilterWinningsClaimed(opts *bind.FilterOpts, receiver []common.Address) (*PriceMarketFactoryWinningsClaimedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PriceMarketFactory.contract.FilterLogs(opts, "WinningsClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &PriceMarketFactoryWinningsClaimedIterator{contract: _PriceMarketFactory.contract, event: "WinningsClaimed", logs: logs, sub: sub}, nil
}

// WatchWinningsClaimed is a free log subscription operation binding the contract event 0xe67bd0100cd3289557430d36098901ba18161e6279c9711d8650b8af10552104.
//
// Solidity: event WinningsClaimed(uint256 id, address winningOutcome, uint256 amount, uint256 settlementFee, uint256 payout, address indexed receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) WatchWinningsClaimed(opts *bind.WatchOpts, sink chan<- *PriceMarketFactoryWinningsClaimed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PriceMarketFactory.contract.WatchLogs(opts, "WinningsClaimed", receiverRule)
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

// ParseWinningsClaimed is a log parse operation binding the contract event 0xe67bd0100cd3289557430d36098901ba18161e6279c9711d8650b8af10552104.
//
// Solidity: event WinningsClaimed(uint256 id, address winningOutcome, uint256 amount, uint256 settlementFee, uint256 payout, address indexed receiver)
func (_PriceMarketFactory *PriceMarketFactoryFilterer) ParseWinningsClaimed(log types.Log) (*PriceMarketFactoryWinningsClaimed, error) {
	event := new(PriceMarketFactoryWinningsClaimed)
	if err := _PriceMarketFactory.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
