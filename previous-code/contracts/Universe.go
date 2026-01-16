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

// UniverseMetaData contains all meta data concerning the Universe contract.
var UniverseMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_parentPayoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"createChildUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feePerCashInAttoCash\",\"type\":\"uint256\"},{\"internalType\":\"contractIAffiliateValidator\",\"name\":\"_affiliateValidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_affiliateFeeDivisor\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_designatedReporterAddress\",\"type\":\"address\"},{\"internalType\":\"int256[]\",\"name\":\"_prices\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"_numTicks\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_extraInfo\",\"type\":\"string\"}],\"name\":\"createScalarMarket\",\"outputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_newMarket\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"decrementOpenInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"decrementOpenInterestFromMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fork\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"}],\"name\":\"getChildUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getCurrentDisputeWindow\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getDisputeRoundDurationInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDisputeThresholdForDisputePacing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDisputeThresholdForFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getDisputeWindowStartTimeAndDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkReputationGoal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkingMarket\",\"outputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialReportMinValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOpenInterestInAttoCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheDesignatedReportNoShowBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheDesignatedReportStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheMarketRepBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheReportingFeeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheValidityBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getOrCreateCurrentDisputeWindow\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getOrCreateDisputeWindowByTimestamp\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getOrCreateNextDisputeWindow\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getOrCreatePreviousDisputeWindow\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getParentPayoutDistributionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getParentUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getPayoutNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPayoutNumerators\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReportingFeeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReputationToken\",\"outputs\":[{\"internalType\":\"contractIV2ReputationToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetRepMarketCapInAttoCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getWinningChildPayoutNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWinningChildUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"incrementOpenInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"_shadyTarget\",\"type\":\"address\"}],\"name\":\"isContainerForDisputeWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_shadyTarget\",\"type\":\"address\"}],\"name\":\"isContainerForMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIReportingParticipant\",\"name\":\"_reportingParticipant\",\"type\":\"address\"}],\"name\":\"isContainerForReportingParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isForking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isForkingMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOpenInterestCash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_shadyChild\",\"type\":\"address\"}],\"name\":\"isParentOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cashBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_marketOI\",\"type\":\"uint256\"}],\"name\":\"migrateMarketIn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_destinationUniverse\",\"type\":\"address\"}],\"name\":\"migrateMarketOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sweepInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateForkValues\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"}],\"name\":\"updateTentativeWinningChildUniverse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniverseABI is the input ABI used to generate the binding from.
// Deprecated: Use UniverseMetaData.ABI instead.
var UniverseABI = UniverseMetaData.ABI

// Universe is an auto generated Go binding around an Ethereum contract.
type Universe struct {
	UniverseCaller     // Read-only binding to the contract
	UniverseTransactor // Write-only binding to the contract
	UniverseFilterer   // Log filterer for contract events
}

// UniverseCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniverseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniverseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniverseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniverseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniverseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniverseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniverseSession struct {
	Contract     *Universe         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniverseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniverseCallerSession struct {
	Contract *UniverseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UniverseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniverseTransactorSession struct {
	Contract     *UniverseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UniverseRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniverseRaw struct {
	Contract *Universe // Generic contract binding to access the raw methods on
}

// UniverseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniverseCallerRaw struct {
	Contract *UniverseCaller // Generic read-only contract binding to access the raw methods on
}

// UniverseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniverseTransactorRaw struct {
	Contract *UniverseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniverse creates a new instance of Universe, bound to a specific deployed contract.
func NewUniverse(address common.Address, backend bind.ContractBackend) (*Universe, error) {
	contract, err := bindUniverse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Universe{UniverseCaller: UniverseCaller{contract: contract}, UniverseTransactor: UniverseTransactor{contract: contract}, UniverseFilterer: UniverseFilterer{contract: contract}}, nil
}

// NewUniverseCaller creates a new read-only instance of Universe, bound to a specific deployed contract.
func NewUniverseCaller(address common.Address, caller bind.ContractCaller) (*UniverseCaller, error) {
	contract, err := bindUniverse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniverseCaller{contract: contract}, nil
}

// NewUniverseTransactor creates a new write-only instance of Universe, bound to a specific deployed contract.
func NewUniverseTransactor(address common.Address, transactor bind.ContractTransactor) (*UniverseTransactor, error) {
	contract, err := bindUniverse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniverseTransactor{contract: contract}, nil
}

// NewUniverseFilterer creates a new log filterer instance of Universe, bound to a specific deployed contract.
func NewUniverseFilterer(address common.Address, filterer bind.ContractFilterer) (*UniverseFilterer, error) {
	contract, err := bindUniverse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniverseFilterer{contract: contract}, nil
}

// bindUniverse binds a generic wrapper to an already deployed contract.
func bindUniverse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniverseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Universe *UniverseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Universe.Contract.UniverseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Universe *UniverseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.Contract.UniverseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Universe *UniverseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Universe.Contract.UniverseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Universe *UniverseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Universe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Universe *UniverseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Universe *UniverseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Universe.Contract.contract.Transact(opts, method, params...)
}

// CreationTime is a free data retrieval call binding the contract method 0xd8270dce.
//
// Solidity: function creationTime() view returns(uint256)
func (_Universe *UniverseCaller) CreationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "creationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreationTime is a free data retrieval call binding the contract method 0xd8270dce.
//
// Solidity: function creationTime() view returns(uint256)
func (_Universe *UniverseSession) CreationTime() (*big.Int, error) {
	return _Universe.Contract.CreationTime(&_Universe.CallOpts)
}

// CreationTime is a free data retrieval call binding the contract method 0xd8270dce.
//
// Solidity: function creationTime() view returns(uint256)
func (_Universe *UniverseCallerSession) CreationTime() (*big.Int, error) {
	return _Universe.Contract.CreationTime(&_Universe.CallOpts)
}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(bytes32 _parentPayoutDistributionHash) view returns(address)
func (_Universe *UniverseCaller) GetChildUniverse(opts *bind.CallOpts, _parentPayoutDistributionHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getChildUniverse", _parentPayoutDistributionHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(bytes32 _parentPayoutDistributionHash) view returns(address)
func (_Universe *UniverseSession) GetChildUniverse(_parentPayoutDistributionHash [32]byte) (common.Address, error) {
	return _Universe.Contract.GetChildUniverse(&_Universe.CallOpts, _parentPayoutDistributionHash)
}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(bytes32 _parentPayoutDistributionHash) view returns(address)
func (_Universe *UniverseCallerSession) GetChildUniverse(_parentPayoutDistributionHash [32]byte) (common.Address, error) {
	return _Universe.Contract.GetChildUniverse(&_Universe.CallOpts, _parentPayoutDistributionHash)
}

// GetCurrentDisputeWindow is a free data retrieval call binding the contract method 0x8699d434.
//
// Solidity: function getCurrentDisputeWindow(bool _initial) view returns(address)
func (_Universe *UniverseCaller) GetCurrentDisputeWindow(opts *bind.CallOpts, _initial bool) (common.Address, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getCurrentDisputeWindow", _initial)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentDisputeWindow is a free data retrieval call binding the contract method 0x8699d434.
//
// Solidity: function getCurrentDisputeWindow(bool _initial) view returns(address)
func (_Universe *UniverseSession) GetCurrentDisputeWindow(_initial bool) (common.Address, error) {
	return _Universe.Contract.GetCurrentDisputeWindow(&_Universe.CallOpts, _initial)
}

// GetCurrentDisputeWindow is a free data retrieval call binding the contract method 0x8699d434.
//
// Solidity: function getCurrentDisputeWindow(bool _initial) view returns(address)
func (_Universe *UniverseCallerSession) GetCurrentDisputeWindow(_initial bool) (common.Address, error) {
	return _Universe.Contract.GetCurrentDisputeWindow(&_Universe.CallOpts, _initial)
}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x879eefa5.
//
// Solidity: function getDisputeRoundDurationInSeconds(bool _initial) view returns(uint256)
func (_Universe *UniverseCaller) GetDisputeRoundDurationInSeconds(opts *bind.CallOpts, _initial bool) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getDisputeRoundDurationInSeconds", _initial)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x879eefa5.
//
// Solidity: function getDisputeRoundDurationInSeconds(bool _initial) view returns(uint256)
func (_Universe *UniverseSession) GetDisputeRoundDurationInSeconds(_initial bool) (*big.Int, error) {
	return _Universe.Contract.GetDisputeRoundDurationInSeconds(&_Universe.CallOpts, _initial)
}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x879eefa5.
//
// Solidity: function getDisputeRoundDurationInSeconds(bool _initial) view returns(uint256)
func (_Universe *UniverseCallerSession) GetDisputeRoundDurationInSeconds(_initial bool) (*big.Int, error) {
	return _Universe.Contract.GetDisputeRoundDurationInSeconds(&_Universe.CallOpts, _initial)
}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() view returns(uint256)
func (_Universe *UniverseCaller) GetDisputeThresholdForDisputePacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getDisputeThresholdForDisputePacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() view returns(uint256)
func (_Universe *UniverseSession) GetDisputeThresholdForDisputePacing() (*big.Int, error) {
	return _Universe.Contract.GetDisputeThresholdForDisputePacing(&_Universe.CallOpts)
}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() view returns(uint256)
func (_Universe *UniverseCallerSession) GetDisputeThresholdForDisputePacing() (*big.Int, error) {
	return _Universe.Contract.GetDisputeThresholdForDisputePacing(&_Universe.CallOpts)
}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() view returns(uint256)
func (_Universe *UniverseCaller) GetDisputeThresholdForFork(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getDisputeThresholdForFork")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() view returns(uint256)
func (_Universe *UniverseSession) GetDisputeThresholdForFork() (*big.Int, error) {
	return _Universe.Contract.GetDisputeThresholdForFork(&_Universe.CallOpts)
}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() view returns(uint256)
func (_Universe *UniverseCallerSession) GetDisputeThresholdForFork() (*big.Int, error) {
	return _Universe.Contract.GetDisputeThresholdForFork(&_Universe.CallOpts)
}

// GetDisputeWindowStartTimeAndDuration is a free data retrieval call binding the contract method 0x5449aed5.
//
// Solidity: function getDisputeWindowStartTimeAndDuration(uint256 _timestamp, bool _initial) view returns(uint256, uint256)
func (_Universe *UniverseCaller) GetDisputeWindowStartTimeAndDuration(opts *bind.CallOpts, _timestamp *big.Int, _initial bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getDisputeWindowStartTimeAndDuration", _timestamp, _initial)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDisputeWindowStartTimeAndDuration is a free data retrieval call binding the contract method 0x5449aed5.
//
// Solidity: function getDisputeWindowStartTimeAndDuration(uint256 _timestamp, bool _initial) view returns(uint256, uint256)
func (_Universe *UniverseSession) GetDisputeWindowStartTimeAndDuration(_timestamp *big.Int, _initial bool) (*big.Int, *big.Int, error) {
	return _Universe.Contract.GetDisputeWindowStartTimeAndDuration(&_Universe.CallOpts, _timestamp, _initial)
}

// GetDisputeWindowStartTimeAndDuration is a free data retrieval call binding the contract method 0x5449aed5.
//
// Solidity: function getDisputeWindowStartTimeAndDuration(uint256 _timestamp, bool _initial) view returns(uint256, uint256)
func (_Universe *UniverseCallerSession) GetDisputeWindowStartTimeAndDuration(_timestamp *big.Int, _initial bool) (*big.Int, *big.Int, error) {
	return _Universe.Contract.GetDisputeWindowStartTimeAndDuration(&_Universe.CallOpts, _timestamp, _initial)
}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() view returns(uint256)
func (_Universe *UniverseCaller) GetForkEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getForkEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() view returns(uint256)
func (_Universe *UniverseSession) GetForkEndTime() (*big.Int, error) {
	return _Universe.Contract.GetForkEndTime(&_Universe.CallOpts)
}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() view returns(uint256)
func (_Universe *UniverseCallerSession) GetForkEndTime() (*big.Int, error) {
	return _Universe.Contract.GetForkEndTime(&_Universe.CallOpts)
}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() view returns(uint256)
func (_Universe *UniverseCaller) GetForkReputationGoal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getForkReputationGoal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() view returns(uint256)
func (_Universe *UniverseSession) GetForkReputationGoal() (*big.Int, error) {
	return _Universe.Contract.GetForkReputationGoal(&_Universe.CallOpts)
}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() view returns(uint256)
func (_Universe *UniverseCallerSession) GetForkReputationGoal() (*big.Int, error) {
	return _Universe.Contract.GetForkReputationGoal(&_Universe.CallOpts)
}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() view returns(address)
func (_Universe *UniverseCaller) GetForkingMarket(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getForkingMarket")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() view returns(address)
func (_Universe *UniverseSession) GetForkingMarket() (common.Address, error) {
	return _Universe.Contract.GetForkingMarket(&_Universe.CallOpts)
}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() view returns(address)
func (_Universe *UniverseCallerSession) GetForkingMarket() (common.Address, error) {
	return _Universe.Contract.GetForkingMarket(&_Universe.CallOpts)
}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() view returns(uint256)
func (_Universe *UniverseCaller) GetInitialReportMinValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getInitialReportMinValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() view returns(uint256)
func (_Universe *UniverseSession) GetInitialReportMinValue() (*big.Int, error) {
	return _Universe.Contract.GetInitialReportMinValue(&_Universe.CallOpts)
}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() view returns(uint256)
func (_Universe *UniverseCallerSession) GetInitialReportMinValue() (*big.Int, error) {
	return _Universe.Contract.GetInitialReportMinValue(&_Universe.CallOpts)
}

// GetOpenInterestInAttoCash is a free data retrieval call binding the contract method 0xc675f222.
//
// Solidity: function getOpenInterestInAttoCash() view returns(uint256)
func (_Universe *UniverseCaller) GetOpenInterestInAttoCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getOpenInterestInAttoCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOpenInterestInAttoCash is a free data retrieval call binding the contract method 0xc675f222.
//
// Solidity: function getOpenInterestInAttoCash() view returns(uint256)
func (_Universe *UniverseSession) GetOpenInterestInAttoCash() (*big.Int, error) {
	return _Universe.Contract.GetOpenInterestInAttoCash(&_Universe.CallOpts)
}

// GetOpenInterestInAttoCash is a free data retrieval call binding the contract method 0xc675f222.
//
// Solidity: function getOpenInterestInAttoCash() view returns(uint256)
func (_Universe *UniverseCallerSession) GetOpenInterestInAttoCash() (*big.Int, error) {
	return _Universe.Contract.GetOpenInterestInAttoCash(&_Universe.CallOpts)
}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() view returns(bytes32)
func (_Universe *UniverseCaller) GetParentPayoutDistributionHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getParentPayoutDistributionHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() view returns(bytes32)
func (_Universe *UniverseSession) GetParentPayoutDistributionHash() ([32]byte, error) {
	return _Universe.Contract.GetParentPayoutDistributionHash(&_Universe.CallOpts)
}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() view returns(bytes32)
func (_Universe *UniverseCallerSession) GetParentPayoutDistributionHash() ([32]byte, error) {
	return _Universe.Contract.GetParentPayoutDistributionHash(&_Universe.CallOpts)
}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() view returns(address)
func (_Universe *UniverseCaller) GetParentUniverse(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getParentUniverse")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() view returns(address)
func (_Universe *UniverseSession) GetParentUniverse() (common.Address, error) {
	return _Universe.Contract.GetParentUniverse(&_Universe.CallOpts)
}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() view returns(address)
func (_Universe *UniverseCallerSession) GetParentUniverse() (common.Address, error) {
	return _Universe.Contract.GetParentUniverse(&_Universe.CallOpts)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Universe *UniverseCaller) GetPayoutNumerator(opts *bind.CallOpts, _outcome *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getPayoutNumerator", _outcome)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Universe *UniverseSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _Universe.Contract.GetPayoutNumerator(&_Universe.CallOpts, _outcome)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Universe *UniverseCallerSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _Universe.Contract.GetPayoutNumerator(&_Universe.CallOpts, _outcome)
}

// GetPayoutNumerators is a free data retrieval call binding the contract method 0x6f84676e.
//
// Solidity: function getPayoutNumerators() view returns(uint256[])
func (_Universe *UniverseCaller) GetPayoutNumerators(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getPayoutNumerators")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPayoutNumerators is a free data retrieval call binding the contract method 0x6f84676e.
//
// Solidity: function getPayoutNumerators() view returns(uint256[])
func (_Universe *UniverseSession) GetPayoutNumerators() ([]*big.Int, error) {
	return _Universe.Contract.GetPayoutNumerators(&_Universe.CallOpts)
}

// GetPayoutNumerators is a free data retrieval call binding the contract method 0x6f84676e.
//
// Solidity: function getPayoutNumerators() view returns(uint256[])
func (_Universe *UniverseCallerSession) GetPayoutNumerators() ([]*big.Int, error) {
	return _Universe.Contract.GetPayoutNumerators(&_Universe.CallOpts)
}

// GetReportingFeeDivisor is a free data retrieval call binding the contract method 0x0dcde5f5.
//
// Solidity: function getReportingFeeDivisor() view returns(uint256)
func (_Universe *UniverseCaller) GetReportingFeeDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getReportingFeeDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportingFeeDivisor is a free data retrieval call binding the contract method 0x0dcde5f5.
//
// Solidity: function getReportingFeeDivisor() view returns(uint256)
func (_Universe *UniverseSession) GetReportingFeeDivisor() (*big.Int, error) {
	return _Universe.Contract.GetReportingFeeDivisor(&_Universe.CallOpts)
}

// GetReportingFeeDivisor is a free data retrieval call binding the contract method 0x0dcde5f5.
//
// Solidity: function getReportingFeeDivisor() view returns(uint256)
func (_Universe *UniverseCallerSession) GetReportingFeeDivisor() (*big.Int, error) {
	return _Universe.Contract.GetReportingFeeDivisor(&_Universe.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() view returns(address)
func (_Universe *UniverseCaller) GetReputationToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getReputationToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() view returns(address)
func (_Universe *UniverseSession) GetReputationToken() (common.Address, error) {
	return _Universe.Contract.GetReputationToken(&_Universe.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() view returns(address)
func (_Universe *UniverseCallerSession) GetReputationToken() (common.Address, error) {
	return _Universe.Contract.GetReputationToken(&_Universe.CallOpts)
}

// GetTargetRepMarketCapInAttoCash is a free data retrieval call binding the contract method 0xdf9fde7e.
//
// Solidity: function getTargetRepMarketCapInAttoCash() view returns(uint256)
func (_Universe *UniverseCaller) GetTargetRepMarketCapInAttoCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getTargetRepMarketCapInAttoCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetRepMarketCapInAttoCash is a free data retrieval call binding the contract method 0xdf9fde7e.
//
// Solidity: function getTargetRepMarketCapInAttoCash() view returns(uint256)
func (_Universe *UniverseSession) GetTargetRepMarketCapInAttoCash() (*big.Int, error) {
	return _Universe.Contract.GetTargetRepMarketCapInAttoCash(&_Universe.CallOpts)
}

// GetTargetRepMarketCapInAttoCash is a free data retrieval call binding the contract method 0xdf9fde7e.
//
// Solidity: function getTargetRepMarketCapInAttoCash() view returns(uint256)
func (_Universe *UniverseCallerSession) GetTargetRepMarketCapInAttoCash() (*big.Int, error) {
	return _Universe.Contract.GetTargetRepMarketCapInAttoCash(&_Universe.CallOpts)
}

// GetWinningChildPayoutNumerator is a free data retrieval call binding the contract method 0x7262f993.
//
// Solidity: function getWinningChildPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Universe *UniverseCaller) GetWinningChildPayoutNumerator(opts *bind.CallOpts, _outcome *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getWinningChildPayoutNumerator", _outcome)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWinningChildPayoutNumerator is a free data retrieval call binding the contract method 0x7262f993.
//
// Solidity: function getWinningChildPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Universe *UniverseSession) GetWinningChildPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _Universe.Contract.GetWinningChildPayoutNumerator(&_Universe.CallOpts, _outcome)
}

// GetWinningChildPayoutNumerator is a free data retrieval call binding the contract method 0x7262f993.
//
// Solidity: function getWinningChildPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Universe *UniverseCallerSession) GetWinningChildPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _Universe.Contract.GetWinningChildPayoutNumerator(&_Universe.CallOpts, _outcome)
}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() view returns(address)
func (_Universe *UniverseCaller) GetWinningChildUniverse(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "getWinningChildUniverse")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() view returns(address)
func (_Universe *UniverseSession) GetWinningChildUniverse() (common.Address, error) {
	return _Universe.Contract.GetWinningChildUniverse(&_Universe.CallOpts)
}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() view returns(address)
func (_Universe *UniverseCallerSession) GetWinningChildUniverse() (common.Address, error) {
	return _Universe.Contract.GetWinningChildUniverse(&_Universe.CallOpts)
}

// IsContainerForDisputeWindow is a free data retrieval call binding the contract method 0x01ba1fa3.
//
// Solidity: function isContainerForDisputeWindow(address _shadyTarget) view returns(bool)
func (_Universe *UniverseCaller) IsContainerForDisputeWindow(opts *bind.CallOpts, _shadyTarget common.Address) (bool, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "isContainerForDisputeWindow", _shadyTarget)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContainerForDisputeWindow is a free data retrieval call binding the contract method 0x01ba1fa3.
//
// Solidity: function isContainerForDisputeWindow(address _shadyTarget) view returns(bool)
func (_Universe *UniverseSession) IsContainerForDisputeWindow(_shadyTarget common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForDisputeWindow(&_Universe.CallOpts, _shadyTarget)
}

// IsContainerForDisputeWindow is a free data retrieval call binding the contract method 0x01ba1fa3.
//
// Solidity: function isContainerForDisputeWindow(address _shadyTarget) view returns(bool)
func (_Universe *UniverseCallerSession) IsContainerForDisputeWindow(_shadyTarget common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForDisputeWindow(&_Universe.CallOpts, _shadyTarget)
}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(address _shadyTarget) view returns(bool)
func (_Universe *UniverseCaller) IsContainerForMarket(opts *bind.CallOpts, _shadyTarget common.Address) (bool, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "isContainerForMarket", _shadyTarget)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(address _shadyTarget) view returns(bool)
func (_Universe *UniverseSession) IsContainerForMarket(_shadyTarget common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForMarket(&_Universe.CallOpts, _shadyTarget)
}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(address _shadyTarget) view returns(bool)
func (_Universe *UniverseCallerSession) IsContainerForMarket(_shadyTarget common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForMarket(&_Universe.CallOpts, _shadyTarget)
}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(address _reportingParticipant) view returns(bool)
func (_Universe *UniverseCaller) IsContainerForReportingParticipant(opts *bind.CallOpts, _reportingParticipant common.Address) (bool, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "isContainerForReportingParticipant", _reportingParticipant)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(address _reportingParticipant) view returns(bool)
func (_Universe *UniverseSession) IsContainerForReportingParticipant(_reportingParticipant common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForReportingParticipant(&_Universe.CallOpts, _reportingParticipant)
}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(address _reportingParticipant) view returns(bool)
func (_Universe *UniverseCallerSession) IsContainerForReportingParticipant(_reportingParticipant common.Address) (bool, error) {
	return _Universe.Contract.IsContainerForReportingParticipant(&_Universe.CallOpts, _reportingParticipant)
}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() view returns(bool)
func (_Universe *UniverseCaller) IsForking(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "isForking")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() view returns(bool)
func (_Universe *UniverseSession) IsForking() (bool, error) {
	return _Universe.Contract.IsForking(&_Universe.CallOpts)
}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() view returns(bool)
func (_Universe *UniverseCallerSession) IsForking() (bool, error) {
	return _Universe.Contract.IsForking(&_Universe.CallOpts)
}

// IsForkingMarket is a free data retrieval call binding the contract method 0xd372fbcd.
//
// Solidity: function isForkingMarket() view returns(bool)
func (_Universe *UniverseCaller) IsForkingMarket(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "isForkingMarket")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForkingMarket is a free data retrieval call binding the contract method 0xd372fbcd.
//
// Solidity: function isForkingMarket() view returns(bool)
func (_Universe *UniverseSession) IsForkingMarket() (bool, error) {
	return _Universe.Contract.IsForkingMarket(&_Universe.CallOpts)
}

// IsForkingMarket is a free data retrieval call binding the contract method 0xd372fbcd.
//
// Solidity: function isForkingMarket() view returns(bool)
func (_Universe *UniverseCallerSession) IsForkingMarket() (bool, error) {
	return _Universe.Contract.IsForkingMarket(&_Universe.CallOpts)
}

// IsOpenInterestCash is a free data retrieval call binding the contract method 0x47d20e3b.
//
// Solidity: function isOpenInterestCash(address ) view returns(bool)
func (_Universe *UniverseCaller) IsOpenInterestCash(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "isOpenInterestCash", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpenInterestCash is a free data retrieval call binding the contract method 0x47d20e3b.
//
// Solidity: function isOpenInterestCash(address ) view returns(bool)
func (_Universe *UniverseSession) IsOpenInterestCash(arg0 common.Address) (bool, error) {
	return _Universe.Contract.IsOpenInterestCash(&_Universe.CallOpts, arg0)
}

// IsOpenInterestCash is a free data retrieval call binding the contract method 0x47d20e3b.
//
// Solidity: function isOpenInterestCash(address ) view returns(bool)
func (_Universe *UniverseCallerSession) IsOpenInterestCash(arg0 common.Address) (bool, error) {
	return _Universe.Contract.IsOpenInterestCash(&_Universe.CallOpts, arg0)
}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(address _shadyChild) view returns(bool)
func (_Universe *UniverseCaller) IsParentOf(opts *bind.CallOpts, _shadyChild common.Address) (bool, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "isParentOf", _shadyChild)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(address _shadyChild) view returns(bool)
func (_Universe *UniverseSession) IsParentOf(_shadyChild common.Address) (bool, error) {
	return _Universe.Contract.IsParentOf(&_Universe.CallOpts, _shadyChild)
}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(address _shadyChild) view returns(bool)
func (_Universe *UniverseCallerSession) IsParentOf(_shadyChild common.Address) (bool, error) {
	return _Universe.Contract.IsParentOf(&_Universe.CallOpts, _shadyChild)
}

// MarketBalance is a free data retrieval call binding the contract method 0x9672e3ba.
//
// Solidity: function marketBalance(address ) view returns(uint256)
func (_Universe *UniverseCaller) MarketBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Universe.contract.Call(opts, &out, "marketBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketBalance is a free data retrieval call binding the contract method 0x9672e3ba.
//
// Solidity: function marketBalance(address ) view returns(uint256)
func (_Universe *UniverseSession) MarketBalance(arg0 common.Address) (*big.Int, error) {
	return _Universe.Contract.MarketBalance(&_Universe.CallOpts, arg0)
}

// MarketBalance is a free data retrieval call binding the contract method 0x9672e3ba.
//
// Solidity: function marketBalance(address ) view returns(uint256)
func (_Universe *UniverseCallerSession) MarketBalance(arg0 common.Address) (*big.Int, error) {
	return _Universe.Contract.MarketBalance(&_Universe.CallOpts, arg0)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x3a537176.
//
// Solidity: function createChildUniverse(uint256[] _parentPayoutNumerators) returns(address)
func (_Universe *UniverseTransactor) CreateChildUniverse(opts *bind.TransactOpts, _parentPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "createChildUniverse", _parentPayoutNumerators)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x3a537176.
//
// Solidity: function createChildUniverse(uint256[] _parentPayoutNumerators) returns(address)
func (_Universe *UniverseSession) CreateChildUniverse(_parentPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Universe.Contract.CreateChildUniverse(&_Universe.TransactOpts, _parentPayoutNumerators)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x3a537176.
//
// Solidity: function createChildUniverse(uint256[] _parentPayoutNumerators) returns(address)
func (_Universe *UniverseTransactorSession) CreateChildUniverse(_parentPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Universe.Contract.CreateChildUniverse(&_Universe.TransactOpts, _parentPayoutNumerators)
}

// CreateScalarMarket is a paid mutator transaction binding the contract method 0x11a80ffc.
//
// Solidity: function createScalarMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, int256[] _prices, uint256 _numTicks, string _extraInfo) returns(address _newMarket)
func (_Universe *UniverseTransactor) CreateScalarMarket(opts *bind.TransactOpts, _endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _prices []*big.Int, _numTicks *big.Int, _extraInfo string) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "createScalarMarket", _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _prices, _numTicks, _extraInfo)
}

// CreateScalarMarket is a paid mutator transaction binding the contract method 0x11a80ffc.
//
// Solidity: function createScalarMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, int256[] _prices, uint256 _numTicks, string _extraInfo) returns(address _newMarket)
func (_Universe *UniverseSession) CreateScalarMarket(_endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _prices []*big.Int, _numTicks *big.Int, _extraInfo string) (*types.Transaction, error) {
	return _Universe.Contract.CreateScalarMarket(&_Universe.TransactOpts, _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _prices, _numTicks, _extraInfo)
}

// CreateScalarMarket is a paid mutator transaction binding the contract method 0x11a80ffc.
//
// Solidity: function createScalarMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, int256[] _prices, uint256 _numTicks, string _extraInfo) returns(address _newMarket)
func (_Universe *UniverseTransactorSession) CreateScalarMarket(_endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _prices []*big.Int, _numTicks *big.Int, _extraInfo string) (*types.Transaction, error) {
	return _Universe.Contract.CreateScalarMarket(&_Universe.TransactOpts, _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _prices, _numTicks, _extraInfo)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(uint256 _amount) returns(bool)
func (_Universe *UniverseTransactor) DecrementOpenInterest(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "decrementOpenInterest", _amount)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(uint256 _amount) returns(bool)
func (_Universe *UniverseSession) DecrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.DecrementOpenInterest(&_Universe.TransactOpts, _amount)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(uint256 _amount) returns(bool)
func (_Universe *UniverseTransactorSession) DecrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.DecrementOpenInterest(&_Universe.TransactOpts, _amount)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x180ef158.
//
// Solidity: function decrementOpenInterestFromMarket(address _market) returns(bool)
func (_Universe *UniverseTransactor) DecrementOpenInterestFromMarket(opts *bind.TransactOpts, _market common.Address) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "decrementOpenInterestFromMarket", _market)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x180ef158.
//
// Solidity: function decrementOpenInterestFromMarket(address _market) returns(bool)
func (_Universe *UniverseSession) DecrementOpenInterestFromMarket(_market common.Address) (*types.Transaction, error) {
	return _Universe.Contract.DecrementOpenInterestFromMarket(&_Universe.TransactOpts, _market)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x180ef158.
//
// Solidity: function decrementOpenInterestFromMarket(address _market) returns(bool)
func (_Universe *UniverseTransactorSession) DecrementOpenInterestFromMarket(_market common.Address) (*types.Transaction, error) {
	return _Universe.Contract.DecrementOpenInterestFromMarket(&_Universe.TransactOpts, _market)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address _sender, uint256 _amount, address _market) returns(bool)
func (_Universe *UniverseTransactor) Deposit(opts *bind.TransactOpts, _sender common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "deposit", _sender, _amount, _market)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address _sender, uint256 _amount, address _market) returns(bool)
func (_Universe *UniverseSession) Deposit(_sender common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Universe.Contract.Deposit(&_Universe.TransactOpts, _sender, _amount, _market)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address _sender, uint256 _amount, address _market) returns(bool)
func (_Universe *UniverseTransactorSession) Deposit(_sender common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Universe.Contract.Deposit(&_Universe.TransactOpts, _sender, _amount, _market)
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_Universe *UniverseTransactor) Fork(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "fork")
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_Universe *UniverseSession) Fork() (*types.Transaction, error) {
	return _Universe.Contract.Fork(&_Universe.TransactOpts)
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_Universe *UniverseTransactorSession) Fork() (*types.Transaction, error) {
	return _Universe.Contract.Fork(&_Universe.TransactOpts)
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_Universe *UniverseTransactor) GetOrCacheDesignatedReportNoShowBond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCacheDesignatedReportNoShowBond")
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_Universe *UniverseSession) GetOrCacheDesignatedReportNoShowBond() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheDesignatedReportNoShowBond(&_Universe.TransactOpts)
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_Universe *UniverseTransactorSession) GetOrCacheDesignatedReportNoShowBond() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheDesignatedReportNoShowBond(&_Universe.TransactOpts)
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_Universe *UniverseTransactor) GetOrCacheDesignatedReportStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCacheDesignatedReportStake")
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_Universe *UniverseSession) GetOrCacheDesignatedReportStake() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheDesignatedReportStake(&_Universe.TransactOpts)
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_Universe *UniverseTransactorSession) GetOrCacheDesignatedReportStake() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheDesignatedReportStake(&_Universe.TransactOpts)
}

// GetOrCacheMarketRepBond is a paid mutator transaction binding the contract method 0xa7e8d762.
//
// Solidity: function getOrCacheMarketRepBond() returns(uint256)
func (_Universe *UniverseTransactor) GetOrCacheMarketRepBond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCacheMarketRepBond")
}

// GetOrCacheMarketRepBond is a paid mutator transaction binding the contract method 0xa7e8d762.
//
// Solidity: function getOrCacheMarketRepBond() returns(uint256)
func (_Universe *UniverseSession) GetOrCacheMarketRepBond() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheMarketRepBond(&_Universe.TransactOpts)
}

// GetOrCacheMarketRepBond is a paid mutator transaction binding the contract method 0xa7e8d762.
//
// Solidity: function getOrCacheMarketRepBond() returns(uint256)
func (_Universe *UniverseTransactorSession) GetOrCacheMarketRepBond() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheMarketRepBond(&_Universe.TransactOpts)
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_Universe *UniverseTransactor) GetOrCacheReportingFeeDivisor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCacheReportingFeeDivisor")
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_Universe *UniverseSession) GetOrCacheReportingFeeDivisor() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheReportingFeeDivisor(&_Universe.TransactOpts)
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_Universe *UniverseTransactorSession) GetOrCacheReportingFeeDivisor() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheReportingFeeDivisor(&_Universe.TransactOpts)
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_Universe *UniverseTransactor) GetOrCacheValidityBond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCacheValidityBond")
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_Universe *UniverseSession) GetOrCacheValidityBond() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheValidityBond(&_Universe.TransactOpts)
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_Universe *UniverseTransactorSession) GetOrCacheValidityBond() (*types.Transaction, error) {
	return _Universe.Contract.GetOrCacheValidityBond(&_Universe.TransactOpts)
}

// GetOrCreateCurrentDisputeWindow is a paid mutator transaction binding the contract method 0xe3fa4b04.
//
// Solidity: function getOrCreateCurrentDisputeWindow(bool _initial) returns(address)
func (_Universe *UniverseTransactor) GetOrCreateCurrentDisputeWindow(opts *bind.TransactOpts, _initial bool) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCreateCurrentDisputeWindow", _initial)
}

// GetOrCreateCurrentDisputeWindow is a paid mutator transaction binding the contract method 0xe3fa4b04.
//
// Solidity: function getOrCreateCurrentDisputeWindow(bool _initial) returns(address)
func (_Universe *UniverseSession) GetOrCreateCurrentDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateCurrentDisputeWindow(&_Universe.TransactOpts, _initial)
}

// GetOrCreateCurrentDisputeWindow is a paid mutator transaction binding the contract method 0xe3fa4b04.
//
// Solidity: function getOrCreateCurrentDisputeWindow(bool _initial) returns(address)
func (_Universe *UniverseTransactorSession) GetOrCreateCurrentDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateCurrentDisputeWindow(&_Universe.TransactOpts, _initial)
}

// GetOrCreateDisputeWindowByTimestamp is a paid mutator transaction binding the contract method 0x8689526b.
//
// Solidity: function getOrCreateDisputeWindowByTimestamp(uint256 _timestamp, bool _initial) returns(address)
func (_Universe *UniverseTransactor) GetOrCreateDisputeWindowByTimestamp(opts *bind.TransactOpts, _timestamp *big.Int, _initial bool) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCreateDisputeWindowByTimestamp", _timestamp, _initial)
}

// GetOrCreateDisputeWindowByTimestamp is a paid mutator transaction binding the contract method 0x8689526b.
//
// Solidity: function getOrCreateDisputeWindowByTimestamp(uint256 _timestamp, bool _initial) returns(address)
func (_Universe *UniverseSession) GetOrCreateDisputeWindowByTimestamp(_timestamp *big.Int, _initial bool) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateDisputeWindowByTimestamp(&_Universe.TransactOpts, _timestamp, _initial)
}

// GetOrCreateDisputeWindowByTimestamp is a paid mutator transaction binding the contract method 0x8689526b.
//
// Solidity: function getOrCreateDisputeWindowByTimestamp(uint256 _timestamp, bool _initial) returns(address)
func (_Universe *UniverseTransactorSession) GetOrCreateDisputeWindowByTimestamp(_timestamp *big.Int, _initial bool) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateDisputeWindowByTimestamp(&_Universe.TransactOpts, _timestamp, _initial)
}

// GetOrCreateNextDisputeWindow is a paid mutator transaction binding the contract method 0x92394f32.
//
// Solidity: function getOrCreateNextDisputeWindow(bool _initial) returns(address)
func (_Universe *UniverseTransactor) GetOrCreateNextDisputeWindow(opts *bind.TransactOpts, _initial bool) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCreateNextDisputeWindow", _initial)
}

// GetOrCreateNextDisputeWindow is a paid mutator transaction binding the contract method 0x92394f32.
//
// Solidity: function getOrCreateNextDisputeWindow(bool _initial) returns(address)
func (_Universe *UniverseSession) GetOrCreateNextDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateNextDisputeWindow(&_Universe.TransactOpts, _initial)
}

// GetOrCreateNextDisputeWindow is a paid mutator transaction binding the contract method 0x92394f32.
//
// Solidity: function getOrCreateNextDisputeWindow(bool _initial) returns(address)
func (_Universe *UniverseTransactorSession) GetOrCreateNextDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreateNextDisputeWindow(&_Universe.TransactOpts, _initial)
}

// GetOrCreatePreviousDisputeWindow is a paid mutator transaction binding the contract method 0xe2d8edaf.
//
// Solidity: function getOrCreatePreviousDisputeWindow(bool _initial) returns(address)
func (_Universe *UniverseTransactor) GetOrCreatePreviousDisputeWindow(opts *bind.TransactOpts, _initial bool) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "getOrCreatePreviousDisputeWindow", _initial)
}

// GetOrCreatePreviousDisputeWindow is a paid mutator transaction binding the contract method 0xe2d8edaf.
//
// Solidity: function getOrCreatePreviousDisputeWindow(bool _initial) returns(address)
func (_Universe *UniverseSession) GetOrCreatePreviousDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreatePreviousDisputeWindow(&_Universe.TransactOpts, _initial)
}

// GetOrCreatePreviousDisputeWindow is a paid mutator transaction binding the contract method 0xe2d8edaf.
//
// Solidity: function getOrCreatePreviousDisputeWindow(bool _initial) returns(address)
func (_Universe *UniverseTransactorSession) GetOrCreatePreviousDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Universe.Contract.GetOrCreatePreviousDisputeWindow(&_Universe.TransactOpts, _initial)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(uint256 _amount) returns(bool)
func (_Universe *UniverseTransactor) IncrementOpenInterest(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "incrementOpenInterest", _amount)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(uint256 _amount) returns(bool)
func (_Universe *UniverseSession) IncrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.IncrementOpenInterest(&_Universe.TransactOpts, _amount)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(uint256 _amount) returns(bool)
func (_Universe *UniverseTransactorSession) IncrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.IncrementOpenInterest(&_Universe.TransactOpts, _amount)
}

// MigrateMarketIn is a paid mutator transaction binding the contract method 0x8d2ecfba.
//
// Solidity: function migrateMarketIn(address _market, uint256 _cashBalance, uint256 _marketOI) returns(bool)
func (_Universe *UniverseTransactor) MigrateMarketIn(opts *bind.TransactOpts, _market common.Address, _cashBalance *big.Int, _marketOI *big.Int) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "migrateMarketIn", _market, _cashBalance, _marketOI)
}

// MigrateMarketIn is a paid mutator transaction binding the contract method 0x8d2ecfba.
//
// Solidity: function migrateMarketIn(address _market, uint256 _cashBalance, uint256 _marketOI) returns(bool)
func (_Universe *UniverseSession) MigrateMarketIn(_market common.Address, _cashBalance *big.Int, _marketOI *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.MigrateMarketIn(&_Universe.TransactOpts, _market, _cashBalance, _marketOI)
}

// MigrateMarketIn is a paid mutator transaction binding the contract method 0x8d2ecfba.
//
// Solidity: function migrateMarketIn(address _market, uint256 _cashBalance, uint256 _marketOI) returns(bool)
func (_Universe *UniverseTransactorSession) MigrateMarketIn(_market common.Address, _cashBalance *big.Int, _marketOI *big.Int) (*types.Transaction, error) {
	return _Universe.Contract.MigrateMarketIn(&_Universe.TransactOpts, _market, _cashBalance, _marketOI)
}

// MigrateMarketOut is a paid mutator transaction binding the contract method 0x11be56d7.
//
// Solidity: function migrateMarketOut(address _destinationUniverse) returns(bool)
func (_Universe *UniverseTransactor) MigrateMarketOut(opts *bind.TransactOpts, _destinationUniverse common.Address) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "migrateMarketOut", _destinationUniverse)
}

// MigrateMarketOut is a paid mutator transaction binding the contract method 0x11be56d7.
//
// Solidity: function migrateMarketOut(address _destinationUniverse) returns(bool)
func (_Universe *UniverseSession) MigrateMarketOut(_destinationUniverse common.Address) (*types.Transaction, error) {
	return _Universe.Contract.MigrateMarketOut(&_Universe.TransactOpts, _destinationUniverse)
}

// MigrateMarketOut is a paid mutator transaction binding the contract method 0x11be56d7.
//
// Solidity: function migrateMarketOut(address _destinationUniverse) returns(bool)
func (_Universe *UniverseTransactorSession) MigrateMarketOut(_destinationUniverse common.Address) (*types.Transaction, error) {
	return _Universe.Contract.MigrateMarketOut(&_Universe.TransactOpts, _destinationUniverse)
}

// SweepInterest is a paid mutator transaction binding the contract method 0x3342f689.
//
// Solidity: function sweepInterest() returns(bool)
func (_Universe *UniverseTransactor) SweepInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "sweepInterest")
}

// SweepInterest is a paid mutator transaction binding the contract method 0x3342f689.
//
// Solidity: function sweepInterest() returns(bool)
func (_Universe *UniverseSession) SweepInterest() (*types.Transaction, error) {
	return _Universe.Contract.SweepInterest(&_Universe.TransactOpts)
}

// SweepInterest is a paid mutator transaction binding the contract method 0x3342f689.
//
// Solidity: function sweepInterest() returns(bool)
func (_Universe *UniverseTransactorSession) SweepInterest() (*types.Transaction, error) {
	return _Universe.Contract.SweepInterest(&_Universe.TransactOpts)
}

// UpdateForkValues is a paid mutator transaction binding the contract method 0x9ab448d9.
//
// Solidity: function updateForkValues() returns(bool)
func (_Universe *UniverseTransactor) UpdateForkValues(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "updateForkValues")
}

// UpdateForkValues is a paid mutator transaction binding the contract method 0x9ab448d9.
//
// Solidity: function updateForkValues() returns(bool)
func (_Universe *UniverseSession) UpdateForkValues() (*types.Transaction, error) {
	return _Universe.Contract.UpdateForkValues(&_Universe.TransactOpts)
}

// UpdateForkValues is a paid mutator transaction binding the contract method 0x9ab448d9.
//
// Solidity: function updateForkValues() returns(bool)
func (_Universe *UniverseTransactorSession) UpdateForkValues() (*types.Transaction, error) {
	return _Universe.Contract.UpdateForkValues(&_Universe.TransactOpts)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(bytes32 _parentPayoutDistributionHash) returns(bool)
func (_Universe *UniverseTransactor) UpdateTentativeWinningChildUniverse(opts *bind.TransactOpts, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "updateTentativeWinningChildUniverse", _parentPayoutDistributionHash)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(bytes32 _parentPayoutDistributionHash) returns(bool)
func (_Universe *UniverseSession) UpdateTentativeWinningChildUniverse(_parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Universe.Contract.UpdateTentativeWinningChildUniverse(&_Universe.TransactOpts, _parentPayoutDistributionHash)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(bytes32 _parentPayoutDistributionHash) returns(bool)
func (_Universe *UniverseTransactorSession) UpdateTentativeWinningChildUniverse(_parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Universe.Contract.UpdateTentativeWinningChildUniverse(&_Universe.TransactOpts, _parentPayoutDistributionHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address _recipient, uint256 _amount, address _market) returns(bool)
func (_Universe *UniverseTransactor) Withdraw(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Universe.contract.Transact(opts, "withdraw", _recipient, _amount, _market)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address _recipient, uint256 _amount, address _market) returns(bool)
func (_Universe *UniverseSession) Withdraw(_recipient common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Universe.Contract.Withdraw(&_Universe.TransactOpts, _recipient, _amount, _market)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address _recipient, uint256 _amount, address _market) returns(bool)
func (_Universe *UniverseTransactorSession) Withdraw(_recipient common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Universe.Contract.Withdraw(&_Universe.TransactOpts, _recipient, _amount, _market)
}
