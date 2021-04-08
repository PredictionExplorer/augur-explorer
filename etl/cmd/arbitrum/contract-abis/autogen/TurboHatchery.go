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

// HasTurboStructTurbo is an auto generated low-level Go binding around an user-defined struct.
type HasTurboStructTurbo struct {
	Creator     common.Address
	CreatorFee  *big.Int
	NumTicks    *big.Int
	Arbiter     common.Address
	ShareTokens []common.Address
	CreatorFees *big.Int
}

// TurboHatcheryABI is the input ABI used to generate the binding from.
const TurboHatcheryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"turboId\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"turboId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"CompleteSetsBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"turboId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"CompleteSetsMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creatorFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"outcomeSymbols\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"outcomeNames\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numTicks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIArbiter\",\"name\":\"arbiter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arbiterConfiguration\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"TurboCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"burnCompleteSets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"claimWinnings\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creatorFee\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_outcomeSymbols\",\"type\":\"string[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_outcomeNames\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_numTicks\",\"type\":\"uint256\"},{\"internalType\":\"contractIArbiter\",\"name\":\"_arbiter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_arbiterConfiguration\",\"type\":\"bytes\"}],\"name\":\"createTurbo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePot\",\"outputs\":[{\"internalType\":\"contractIFeePot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getShareTokens\",\"outputs\":[{\"internalType\":\"contractITurboShareToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTurboLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mintCompleteSets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"contractITurboShareTokenFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_turboId\",\"type\":\"uint256\"}],\"name\":\"turbos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"creatorFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numTicks\",\"type\":\"uint256\"},{\"internalType\":\"contractIArbiter\",\"name\":\"arbiter\",\"type\":\"address\"},{\"internalType\":\"contractITurboShareToken[]\",\"name\":\"shareTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"creatorFees\",\"type\":\"uint256\"}],\"internalType\":\"structHasTurboStruct.Turbo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"withdrawCreatorFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TurboHatchery is an auto generated Go binding around an Ethereum contract.
type TurboHatchery struct {
	TurboHatcheryCaller     // Read-only binding to the contract
	TurboHatcheryTransactor // Write-only binding to the contract
	TurboHatcheryFilterer   // Log filterer for contract events
}

// TurboHatcheryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TurboHatcheryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TurboHatcheryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TurboHatcheryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TurboHatcheryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TurboHatcheryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TurboHatcherySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TurboHatcherySession struct {
	Contract     *TurboHatchery    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TurboHatcheryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TurboHatcheryCallerSession struct {
	Contract *TurboHatcheryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TurboHatcheryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TurboHatcheryTransactorSession struct {
	Contract     *TurboHatcheryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TurboHatcheryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TurboHatcheryRaw struct {
	Contract *TurboHatchery // Generic contract binding to access the raw methods on
}

// TurboHatcheryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TurboHatcheryCallerRaw struct {
	Contract *TurboHatcheryCaller // Generic read-only contract binding to access the raw methods on
}

// TurboHatcheryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TurboHatcheryTransactorRaw struct {
	Contract *TurboHatcheryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTurboHatchery creates a new instance of TurboHatchery, bound to a specific deployed contract.
func NewTurboHatchery(address common.Address, backend bind.ContractBackend) (*TurboHatchery, error) {
	contract, err := bindTurboHatchery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TurboHatchery{TurboHatcheryCaller: TurboHatcheryCaller{contract: contract}, TurboHatcheryTransactor: TurboHatcheryTransactor{contract: contract}, TurboHatcheryFilterer: TurboHatcheryFilterer{contract: contract}}, nil
}

// NewTurboHatcheryCaller creates a new read-only instance of TurboHatchery, bound to a specific deployed contract.
func NewTurboHatcheryCaller(address common.Address, caller bind.ContractCaller) (*TurboHatcheryCaller, error) {
	contract, err := bindTurboHatchery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TurboHatcheryCaller{contract: contract}, nil
}

// NewTurboHatcheryTransactor creates a new write-only instance of TurboHatchery, bound to a specific deployed contract.
func NewTurboHatcheryTransactor(address common.Address, transactor bind.ContractTransactor) (*TurboHatcheryTransactor, error) {
	contract, err := bindTurboHatchery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TurboHatcheryTransactor{contract: contract}, nil
}

// NewTurboHatcheryFilterer creates a new log filterer instance of TurboHatchery, bound to a specific deployed contract.
func NewTurboHatcheryFilterer(address common.Address, filterer bind.ContractFilterer) (*TurboHatcheryFilterer, error) {
	contract, err := bindTurboHatchery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TurboHatcheryFilterer{contract: contract}, nil
}

// bindTurboHatchery binds a generic wrapper to an already deployed contract.
func bindTurboHatchery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TurboHatcheryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TurboHatchery *TurboHatcheryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TurboHatchery.Contract.TurboHatcheryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TurboHatchery *TurboHatcheryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TurboHatchery.Contract.TurboHatcheryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TurboHatchery *TurboHatcheryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TurboHatchery.Contract.TurboHatcheryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TurboHatchery *TurboHatcheryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TurboHatchery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TurboHatchery *TurboHatcheryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TurboHatchery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TurboHatchery *TurboHatcheryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TurboHatchery.Contract.contract.Transact(opts, method, params...)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_TurboHatchery *TurboHatcheryCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TurboHatchery.contract.Call(opts, out, "collateral")
	return *ret0, err
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_TurboHatchery *TurboHatcherySession) Collateral() (common.Address, error) {
	return _TurboHatchery.Contract.Collateral(&_TurboHatchery.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_TurboHatchery *TurboHatcheryCallerSession) Collateral() (common.Address, error) {
	return _TurboHatchery.Contract.Collateral(&_TurboHatchery.CallOpts)
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_TurboHatchery *TurboHatcheryCaller) FeePot(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TurboHatchery.contract.Call(opts, out, "feePot")
	return *ret0, err
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_TurboHatchery *TurboHatcherySession) FeePot() (common.Address, error) {
	return _TurboHatchery.Contract.FeePot(&_TurboHatchery.CallOpts)
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_TurboHatchery *TurboHatcheryCallerSession) FeePot() (common.Address, error) {
	return _TurboHatchery.Contract.FeePot(&_TurboHatchery.CallOpts)
}

// GetShareTokens is a free data retrieval call binding the contract method 0xcabae98f.
//
// Solidity: function getShareTokens(uint256 _id) view returns(address[])
func (_TurboHatchery *TurboHatcheryCaller) GetShareTokens(opts *bind.CallOpts, _id *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TurboHatchery.contract.Call(opts, out, "getShareTokens", _id)
	return *ret0, err
}

// GetShareTokens is a free data retrieval call binding the contract method 0xcabae98f.
//
// Solidity: function getShareTokens(uint256 _id) view returns(address[])
func (_TurboHatchery *TurboHatcherySession) GetShareTokens(_id *big.Int) ([]common.Address, error) {
	return _TurboHatchery.Contract.GetShareTokens(&_TurboHatchery.CallOpts, _id)
}

// GetShareTokens is a free data retrieval call binding the contract method 0xcabae98f.
//
// Solidity: function getShareTokens(uint256 _id) view returns(address[])
func (_TurboHatchery *TurboHatcheryCallerSession) GetShareTokens(_id *big.Int) ([]common.Address, error) {
	return _TurboHatchery.Contract.GetShareTokens(&_TurboHatchery.CallOpts, _id)
}

// GetTurboLength is a free data retrieval call binding the contract method 0x31540160.
//
// Solidity: function getTurboLength() view returns(uint256)
func (_TurboHatchery *TurboHatcheryCaller) GetTurboLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TurboHatchery.contract.Call(opts, out, "getTurboLength")
	return *ret0, err
}

// GetTurboLength is a free data retrieval call binding the contract method 0x31540160.
//
// Solidity: function getTurboLength() view returns(uint256)
func (_TurboHatchery *TurboHatcherySession) GetTurboLength() (*big.Int, error) {
	return _TurboHatchery.Contract.GetTurboLength(&_TurboHatchery.CallOpts)
}

// GetTurboLength is a free data retrieval call binding the contract method 0x31540160.
//
// Solidity: function getTurboLength() view returns(uint256)
func (_TurboHatchery *TurboHatcheryCallerSession) GetTurboLength() (*big.Int, error) {
	return _TurboHatchery.Contract.GetTurboLength(&_TurboHatchery.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_TurboHatchery *TurboHatcheryCaller) TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TurboHatchery.contract.Call(opts, out, "tokenFactory")
	return *ret0, err
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_TurboHatchery *TurboHatcherySession) TokenFactory() (common.Address, error) {
	return _TurboHatchery.Contract.TokenFactory(&_TurboHatchery.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_TurboHatchery *TurboHatcheryCallerSession) TokenFactory() (common.Address, error) {
	return _TurboHatchery.Contract.TokenFactory(&_TurboHatchery.CallOpts)
}

// Turbos is a free data retrieval call binding the contract method 0x4185c085.
//
// Solidity: function turbos(uint256 _turboId) view returns(HasTurboStructTurbo)
func (_TurboHatchery *TurboHatcheryCaller) Turbos(opts *bind.CallOpts, _turboId *big.Int) (HasTurboStructTurbo, error) {
	var (
		ret0 = new(HasTurboStructTurbo)
	)
	out := ret0
	err := _TurboHatchery.contract.Call(opts, out, "turbos", _turboId)
	return *ret0, err
}

// Turbos is a free data retrieval call binding the contract method 0x4185c085.
//
// Solidity: function turbos(uint256 _turboId) view returns(HasTurboStructTurbo)
func (_TurboHatchery *TurboHatcherySession) Turbos(_turboId *big.Int) (HasTurboStructTurbo, error) {
	return _TurboHatchery.Contract.Turbos(&_TurboHatchery.CallOpts, _turboId)
}

// Turbos is a free data retrieval call binding the contract method 0x4185c085.
//
// Solidity: function turbos(uint256 _turboId) view returns(HasTurboStructTurbo)
func (_TurboHatchery *TurboHatcheryCallerSession) Turbos(_turboId *big.Int) (HasTurboStructTurbo, error) {
	return _TurboHatchery.Contract.Turbos(&_TurboHatchery.CallOpts, _turboId)
}

// BurnCompleteSets is a paid mutator transaction binding the contract method 0xdefbde95.
//
// Solidity: function burnCompleteSets(uint256 _id, uint256 _amount, address _receiver) returns(bool)
func (_TurboHatchery *TurboHatcheryTransactor) BurnCompleteSets(opts *bind.TransactOpts, _id *big.Int, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TurboHatchery.contract.Transact(opts, "burnCompleteSets", _id, _amount, _receiver)
}

// BurnCompleteSets is a paid mutator transaction binding the contract method 0xdefbde95.
//
// Solidity: function burnCompleteSets(uint256 _id, uint256 _amount, address _receiver) returns(bool)
func (_TurboHatchery *TurboHatcherySession) BurnCompleteSets(_id *big.Int, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TurboHatchery.Contract.BurnCompleteSets(&_TurboHatchery.TransactOpts, _id, _amount, _receiver)
}

// BurnCompleteSets is a paid mutator transaction binding the contract method 0xdefbde95.
//
// Solidity: function burnCompleteSets(uint256 _id, uint256 _amount, address _receiver) returns(bool)
func (_TurboHatchery *TurboHatcheryTransactorSession) BurnCompleteSets(_id *big.Int, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TurboHatchery.Contract.BurnCompleteSets(&_TurboHatchery.TransactOpts, _id, _amount, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x677bd9ff.
//
// Solidity: function claimWinnings(uint256 _id) returns(bool)
func (_TurboHatchery *TurboHatcheryTransactor) ClaimWinnings(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _TurboHatchery.contract.Transact(opts, "claimWinnings", _id)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x677bd9ff.
//
// Solidity: function claimWinnings(uint256 _id) returns(bool)
func (_TurboHatchery *TurboHatcherySession) ClaimWinnings(_id *big.Int) (*types.Transaction, error) {
	return _TurboHatchery.Contract.ClaimWinnings(&_TurboHatchery.TransactOpts, _id)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x677bd9ff.
//
// Solidity: function claimWinnings(uint256 _id) returns(bool)
func (_TurboHatchery *TurboHatcheryTransactorSession) ClaimWinnings(_id *big.Int) (*types.Transaction, error) {
	return _TurboHatchery.Contract.ClaimWinnings(&_TurboHatchery.TransactOpts, _id)
}

// CreateTurbo is a paid mutator transaction binding the contract method 0x363bfc92.
//
// Solidity: function createTurbo(uint256 _index, uint256 _creatorFee, string[] _outcomeSymbols, bytes32[] _outcomeNames, uint256 _numTicks, address _arbiter, bytes _arbiterConfiguration) returns(uint256)
func (_TurboHatchery *TurboHatcheryTransactor) CreateTurbo(opts *bind.TransactOpts, _index *big.Int, _creatorFee *big.Int, _outcomeSymbols []string, _outcomeNames [][32]byte, _numTicks *big.Int, _arbiter common.Address, _arbiterConfiguration []byte) (*types.Transaction, error) {
	return _TurboHatchery.contract.Transact(opts, "createTurbo", _index, _creatorFee, _outcomeSymbols, _outcomeNames, _numTicks, _arbiter, _arbiterConfiguration)
}

// CreateTurbo is a paid mutator transaction binding the contract method 0x363bfc92.
//
// Solidity: function createTurbo(uint256 _index, uint256 _creatorFee, string[] _outcomeSymbols, bytes32[] _outcomeNames, uint256 _numTicks, address _arbiter, bytes _arbiterConfiguration) returns(uint256)
func (_TurboHatchery *TurboHatcherySession) CreateTurbo(_index *big.Int, _creatorFee *big.Int, _outcomeSymbols []string, _outcomeNames [][32]byte, _numTicks *big.Int, _arbiter common.Address, _arbiterConfiguration []byte) (*types.Transaction, error) {
	return _TurboHatchery.Contract.CreateTurbo(&_TurboHatchery.TransactOpts, _index, _creatorFee, _outcomeSymbols, _outcomeNames, _numTicks, _arbiter, _arbiterConfiguration)
}

// CreateTurbo is a paid mutator transaction binding the contract method 0x363bfc92.
//
// Solidity: function createTurbo(uint256 _index, uint256 _creatorFee, string[] _outcomeSymbols, bytes32[] _outcomeNames, uint256 _numTicks, address _arbiter, bytes _arbiterConfiguration) returns(uint256)
func (_TurboHatchery *TurboHatcheryTransactorSession) CreateTurbo(_index *big.Int, _creatorFee *big.Int, _outcomeSymbols []string, _outcomeNames [][32]byte, _numTicks *big.Int, _arbiter common.Address, _arbiterConfiguration []byte) (*types.Transaction, error) {
	return _TurboHatchery.Contract.CreateTurbo(&_TurboHatchery.TransactOpts, _index, _creatorFee, _outcomeSymbols, _outcomeNames, _numTicks, _arbiter, _arbiterConfiguration)
}

// MintCompleteSets is a paid mutator transaction binding the contract method 0x758f779f.
//
// Solidity: function mintCompleteSets(uint256 _id, uint256 _amount, address _receiver) returns(bool)
func (_TurboHatchery *TurboHatcheryTransactor) MintCompleteSets(opts *bind.TransactOpts, _id *big.Int, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TurboHatchery.contract.Transact(opts, "mintCompleteSets", _id, _amount, _receiver)
}

// MintCompleteSets is a paid mutator transaction binding the contract method 0x758f779f.
//
// Solidity: function mintCompleteSets(uint256 _id, uint256 _amount, address _receiver) returns(bool)
func (_TurboHatchery *TurboHatcherySession) MintCompleteSets(_id *big.Int, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TurboHatchery.Contract.MintCompleteSets(&_TurboHatchery.TransactOpts, _id, _amount, _receiver)
}

// MintCompleteSets is a paid mutator transaction binding the contract method 0x758f779f.
//
// Solidity: function mintCompleteSets(uint256 _id, uint256 _amount, address _receiver) returns(bool)
func (_TurboHatchery *TurboHatcheryTransactorSession) MintCompleteSets(_id *big.Int, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TurboHatchery.Contract.MintCompleteSets(&_TurboHatchery.TransactOpts, _id, _amount, _receiver)
}

// WithdrawCreatorFees is a paid mutator transaction binding the contract method 0x73962620.
//
// Solidity: function withdrawCreatorFees(uint256 _id) returns(bool)
func (_TurboHatchery *TurboHatcheryTransactor) WithdrawCreatorFees(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _TurboHatchery.contract.Transact(opts, "withdrawCreatorFees", _id)
}

// WithdrawCreatorFees is a paid mutator transaction binding the contract method 0x73962620.
//
// Solidity: function withdrawCreatorFees(uint256 _id) returns(bool)
func (_TurboHatchery *TurboHatcherySession) WithdrawCreatorFees(_id *big.Int) (*types.Transaction, error) {
	return _TurboHatchery.Contract.WithdrawCreatorFees(&_TurboHatchery.TransactOpts, _id)
}

// WithdrawCreatorFees is a paid mutator transaction binding the contract method 0x73962620.
//
// Solidity: function withdrawCreatorFees(uint256 _id) returns(bool)
func (_TurboHatchery *TurboHatcheryTransactorSession) WithdrawCreatorFees(_id *big.Int) (*types.Transaction, error) {
	return _TurboHatchery.Contract.WithdrawCreatorFees(&_TurboHatchery.TransactOpts, _id)
}

// TurboHatcheryClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the TurboHatchery contract.
type TurboHatcheryClaimIterator struct {
	Event *TurboHatcheryClaim // Event containing the contract specifics and raw log

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
func (it *TurboHatcheryClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TurboHatcheryClaim)
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
		it.Event = new(TurboHatcheryClaim)
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
func (it *TurboHatcheryClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TurboHatcheryClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TurboHatcheryClaim represents a Claim event raised by the TurboHatchery contract.
type TurboHatcheryClaim struct {
	TurboId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f7.
//
// Solidity: event Claim(uint256 turboId)
func (_TurboHatchery *TurboHatcheryFilterer) FilterClaim(opts *bind.FilterOpts) (*TurboHatcheryClaimIterator, error) {

	logs, sub, err := _TurboHatchery.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &TurboHatcheryClaimIterator{contract: _TurboHatchery.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f7.
//
// Solidity: event Claim(uint256 turboId)
func (_TurboHatchery *TurboHatcheryFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *TurboHatcheryClaim) (event.Subscription, error) {

	logs, sub, err := _TurboHatchery.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TurboHatcheryClaim)
				if err := _TurboHatchery.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f7.
//
// Solidity: event Claim(uint256 turboId)
func (_TurboHatchery *TurboHatcheryFilterer) ParseClaim(log types.Log) (*TurboHatcheryClaim, error) {
	event := new(TurboHatcheryClaim)
	if err := _TurboHatchery.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TurboHatcheryCompleteSetsBurnedIterator is returned from FilterCompleteSetsBurned and is used to iterate over the raw logs and unpacked data for CompleteSetsBurned events raised by the TurboHatchery contract.
type TurboHatcheryCompleteSetsBurnedIterator struct {
	Event *TurboHatcheryCompleteSetsBurned // Event containing the contract specifics and raw log

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
func (it *TurboHatcheryCompleteSetsBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TurboHatcheryCompleteSetsBurned)
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
		it.Event = new(TurboHatcheryCompleteSetsBurned)
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
func (it *TurboHatcheryCompleteSetsBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TurboHatcheryCompleteSetsBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TurboHatcheryCompleteSetsBurned represents a CompleteSetsBurned event raised by the TurboHatchery contract.
type TurboHatcheryCompleteSetsBurned struct {
	TurboId *big.Int
	Amount  *big.Int
	Target  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCompleteSetsBurned is a free log retrieval operation binding the contract event 0x2df8f390c89a8c8e8b89875f61085269c64b16b81e7745b844ba42a40a3dde27.
//
// Solidity: event CompleteSetsBurned(uint256 turboId, uint256 amount, address target)
func (_TurboHatchery *TurboHatcheryFilterer) FilterCompleteSetsBurned(opts *bind.FilterOpts) (*TurboHatcheryCompleteSetsBurnedIterator, error) {

	logs, sub, err := _TurboHatchery.contract.FilterLogs(opts, "CompleteSetsBurned")
	if err != nil {
		return nil, err
	}
	return &TurboHatcheryCompleteSetsBurnedIterator{contract: _TurboHatchery.contract, event: "CompleteSetsBurned", logs: logs, sub: sub}, nil
}

// WatchCompleteSetsBurned is a free log subscription operation binding the contract event 0x2df8f390c89a8c8e8b89875f61085269c64b16b81e7745b844ba42a40a3dde27.
//
// Solidity: event CompleteSetsBurned(uint256 turboId, uint256 amount, address target)
func (_TurboHatchery *TurboHatcheryFilterer) WatchCompleteSetsBurned(opts *bind.WatchOpts, sink chan<- *TurboHatcheryCompleteSetsBurned) (event.Subscription, error) {

	logs, sub, err := _TurboHatchery.contract.WatchLogs(opts, "CompleteSetsBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TurboHatcheryCompleteSetsBurned)
				if err := _TurboHatchery.contract.UnpackLog(event, "CompleteSetsBurned", log); err != nil {
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

// ParseCompleteSetsBurned is a log parse operation binding the contract event 0x2df8f390c89a8c8e8b89875f61085269c64b16b81e7745b844ba42a40a3dde27.
//
// Solidity: event CompleteSetsBurned(uint256 turboId, uint256 amount, address target)
func (_TurboHatchery *TurboHatcheryFilterer) ParseCompleteSetsBurned(log types.Log) (*TurboHatcheryCompleteSetsBurned, error) {
	event := new(TurboHatcheryCompleteSetsBurned)
	if err := _TurboHatchery.contract.UnpackLog(event, "CompleteSetsBurned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TurboHatcheryCompleteSetsMintedIterator is returned from FilterCompleteSetsMinted and is used to iterate over the raw logs and unpacked data for CompleteSetsMinted events raised by the TurboHatchery contract.
type TurboHatcheryCompleteSetsMintedIterator struct {
	Event *TurboHatcheryCompleteSetsMinted // Event containing the contract specifics and raw log

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
func (it *TurboHatcheryCompleteSetsMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TurboHatcheryCompleteSetsMinted)
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
		it.Event = new(TurboHatcheryCompleteSetsMinted)
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
func (it *TurboHatcheryCompleteSetsMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TurboHatcheryCompleteSetsMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TurboHatcheryCompleteSetsMinted represents a CompleteSetsMinted event raised by the TurboHatchery contract.
type TurboHatcheryCompleteSetsMinted struct {
	TurboId *big.Int
	Amount  *big.Int
	Target  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCompleteSetsMinted is a free log retrieval operation binding the contract event 0x51b2bca5bb2f65b2670950591ce7b54cfc4d99b2db85abfea36b8b92d10ac380.
//
// Solidity: event CompleteSetsMinted(uint256 turboId, uint256 amount, address target)
func (_TurboHatchery *TurboHatcheryFilterer) FilterCompleteSetsMinted(opts *bind.FilterOpts) (*TurboHatcheryCompleteSetsMintedIterator, error) {

	logs, sub, err := _TurboHatchery.contract.FilterLogs(opts, "CompleteSetsMinted")
	if err != nil {
		return nil, err
	}
	return &TurboHatcheryCompleteSetsMintedIterator{contract: _TurboHatchery.contract, event: "CompleteSetsMinted", logs: logs, sub: sub}, nil
}

// WatchCompleteSetsMinted is a free log subscription operation binding the contract event 0x51b2bca5bb2f65b2670950591ce7b54cfc4d99b2db85abfea36b8b92d10ac380.
//
// Solidity: event CompleteSetsMinted(uint256 turboId, uint256 amount, address target)
func (_TurboHatchery *TurboHatcheryFilterer) WatchCompleteSetsMinted(opts *bind.WatchOpts, sink chan<- *TurboHatcheryCompleteSetsMinted) (event.Subscription, error) {

	logs, sub, err := _TurboHatchery.contract.WatchLogs(opts, "CompleteSetsMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TurboHatcheryCompleteSetsMinted)
				if err := _TurboHatchery.contract.UnpackLog(event, "CompleteSetsMinted", log); err != nil {
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

// ParseCompleteSetsMinted is a log parse operation binding the contract event 0x51b2bca5bb2f65b2670950591ce7b54cfc4d99b2db85abfea36b8b92d10ac380.
//
// Solidity: event CompleteSetsMinted(uint256 turboId, uint256 amount, address target)
func (_TurboHatchery *TurboHatcheryFilterer) ParseCompleteSetsMinted(log types.Log) (*TurboHatcheryCompleteSetsMinted, error) {
	event := new(TurboHatcheryCompleteSetsMinted)
	if err := _TurboHatchery.contract.UnpackLog(event, "CompleteSetsMinted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TurboHatcheryTurboCreatedIterator is returned from FilterTurboCreated and is used to iterate over the raw logs and unpacked data for TurboCreated events raised by the TurboHatchery contract.
type TurboHatcheryTurboCreatedIterator struct {
	Event *TurboHatcheryTurboCreated // Event containing the contract specifics and raw log

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
func (it *TurboHatcheryTurboCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TurboHatcheryTurboCreated)
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
		it.Event = new(TurboHatcheryTurboCreated)
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
func (it *TurboHatcheryTurboCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TurboHatcheryTurboCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TurboHatcheryTurboCreated represents a TurboCreated event raised by the TurboHatchery contract.
type TurboHatcheryTurboCreated struct {
	Id                   *big.Int
	CreatorFee           *big.Int
	OutcomeSymbols       []string
	OutcomeNames         [][32]byte
	NumTicks             *big.Int
	Arbiter              common.Address
	ArbiterConfiguration []byte
	Index                *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTurboCreated is a free log retrieval operation binding the contract event 0x2c4d919a4805caed2e2fdd9bb8a122413c2a643b61e08b957445484bbbfd8f4f.
//
// Solidity: event TurboCreated(uint256 id, uint256 creatorFee, string[] outcomeSymbols, bytes32[] outcomeNames, uint256 numTicks, address arbiter, bytes arbiterConfiguration, uint256 indexed index)
func (_TurboHatchery *TurboHatcheryFilterer) FilterTurboCreated(opts *bind.FilterOpts, index []*big.Int) (*TurboHatcheryTurboCreatedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _TurboHatchery.contract.FilterLogs(opts, "TurboCreated", indexRule)
	if err != nil {
		return nil, err
	}
	return &TurboHatcheryTurboCreatedIterator{contract: _TurboHatchery.contract, event: "TurboCreated", logs: logs, sub: sub}, nil
}

// WatchTurboCreated is a free log subscription operation binding the contract event 0x2c4d919a4805caed2e2fdd9bb8a122413c2a643b61e08b957445484bbbfd8f4f.
//
// Solidity: event TurboCreated(uint256 id, uint256 creatorFee, string[] outcomeSymbols, bytes32[] outcomeNames, uint256 numTicks, address arbiter, bytes arbiterConfiguration, uint256 indexed index)
func (_TurboHatchery *TurboHatcheryFilterer) WatchTurboCreated(opts *bind.WatchOpts, sink chan<- *TurboHatcheryTurboCreated, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _TurboHatchery.contract.WatchLogs(opts, "TurboCreated", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TurboHatcheryTurboCreated)
				if err := _TurboHatchery.contract.UnpackLog(event, "TurboCreated", log); err != nil {
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

// ParseTurboCreated is a log parse operation binding the contract event 0x2c4d919a4805caed2e2fdd9bb8a122413c2a643b61e08b957445484bbbfd8f4f.
//
// Solidity: event TurboCreated(uint256 id, uint256 creatorFee, string[] outcomeSymbols, bytes32[] outcomeNames, uint256 numTicks, address arbiter, bytes arbiterConfiguration, uint256 indexed index)
func (_TurboHatchery *TurboHatcheryFilterer) ParseTurboCreated(log types.Log) (*TurboHatcheryTurboCreated, error) {
	event := new(TurboHatcheryTurboCreated)
	if err := _TurboHatchery.contract.UnpackLog(event, "TurboCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}
