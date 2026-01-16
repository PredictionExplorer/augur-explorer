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

// GetSwapFeeMetaData contains all meta data concerning the GetSwapFee contract.
var GetSwapFeeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getSwapFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GetSwapFeeABI is the input ABI used to generate the binding from.
// Deprecated: Use GetSwapFeeMetaData.ABI instead.
var GetSwapFeeABI = GetSwapFeeMetaData.ABI

// GetSwapFee is an auto generated Go binding around an Ethereum contract.
type GetSwapFee struct {
	GetSwapFeeCaller     // Read-only binding to the contract
	GetSwapFeeTransactor // Write-only binding to the contract
	GetSwapFeeFilterer   // Log filterer for contract events
}

// GetSwapFeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type GetSwapFeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetSwapFeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GetSwapFeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetSwapFeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GetSwapFeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GetSwapFeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GetSwapFeeSession struct {
	Contract     *GetSwapFee       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GetSwapFeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GetSwapFeeCallerSession struct {
	Contract *GetSwapFeeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GetSwapFeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GetSwapFeeTransactorSession struct {
	Contract     *GetSwapFeeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GetSwapFeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type GetSwapFeeRaw struct {
	Contract *GetSwapFee // Generic contract binding to access the raw methods on
}

// GetSwapFeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GetSwapFeeCallerRaw struct {
	Contract *GetSwapFeeCaller // Generic read-only contract binding to access the raw methods on
}

// GetSwapFeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GetSwapFeeTransactorRaw struct {
	Contract *GetSwapFeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGetSwapFee creates a new instance of GetSwapFee, bound to a specific deployed contract.
func NewGetSwapFee(address common.Address, backend bind.ContractBackend) (*GetSwapFee, error) {
	contract, err := bindGetSwapFee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GetSwapFee{GetSwapFeeCaller: GetSwapFeeCaller{contract: contract}, GetSwapFeeTransactor: GetSwapFeeTransactor{contract: contract}, GetSwapFeeFilterer: GetSwapFeeFilterer{contract: contract}}, nil
}

// NewGetSwapFeeCaller creates a new read-only instance of GetSwapFee, bound to a specific deployed contract.
func NewGetSwapFeeCaller(address common.Address, caller bind.ContractCaller) (*GetSwapFeeCaller, error) {
	contract, err := bindGetSwapFee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GetSwapFeeCaller{contract: contract}, nil
}

// NewGetSwapFeeTransactor creates a new write-only instance of GetSwapFee, bound to a specific deployed contract.
func NewGetSwapFeeTransactor(address common.Address, transactor bind.ContractTransactor) (*GetSwapFeeTransactor, error) {
	contract, err := bindGetSwapFee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GetSwapFeeTransactor{contract: contract}, nil
}

// NewGetSwapFeeFilterer creates a new log filterer instance of GetSwapFee, bound to a specific deployed contract.
func NewGetSwapFeeFilterer(address common.Address, filterer bind.ContractFilterer) (*GetSwapFeeFilterer, error) {
	contract, err := bindGetSwapFee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GetSwapFeeFilterer{contract: contract}, nil
}

// bindGetSwapFee binds a generic wrapper to an already deployed contract.
func bindGetSwapFee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GetSwapFeeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetSwapFee *GetSwapFeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetSwapFee.Contract.GetSwapFeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetSwapFee *GetSwapFeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetSwapFee.Contract.GetSwapFeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetSwapFee *GetSwapFeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetSwapFee.Contract.GetSwapFeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GetSwapFee *GetSwapFeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GetSwapFee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GetSwapFee *GetSwapFeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GetSwapFee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GetSwapFee *GetSwapFeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GetSwapFee.Contract.contract.Transact(opts, method, params...)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_GetSwapFee *GetSwapFeeCaller) GetSwapFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GetSwapFee.contract.Call(opts, &out, "getSwapFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_GetSwapFee *GetSwapFeeSession) GetSwapFeePercentage() (*big.Int, error) {
	return _GetSwapFee.Contract.GetSwapFeePercentage(&_GetSwapFee.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_GetSwapFee *GetSwapFeeCallerSession) GetSwapFeePercentage() (*big.Int, error) {
	return _GetSwapFee.Contract.GetSwapFeePercentage(&_GetSwapFee.CallOpts)
}
