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
	_ = abi.ConvertType
)

// IStakingWalletRWalkMetaData contains all meta data concerning the IStakingWalletRWalk contract.
var IStakingWalletRWalkMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"StakeActionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"UnstakeActionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isTokenStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"lastActionIdByTokenId\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokensStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"entropy\",\"type\":\"bytes32\"}],\"name\":\"pickRandomStaker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"stakerByTokenId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"wasTokenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IStakingWalletRWalkABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingWalletRWalkMetaData.ABI instead.
var IStakingWalletRWalkABI = IStakingWalletRWalkMetaData.ABI

// IStakingWalletRWalk is an auto generated Go binding around an Ethereum contract.
type IStakingWalletRWalk struct {
	IStakingWalletRWalkCaller     // Read-only binding to the contract
	IStakingWalletRWalkTransactor // Write-only binding to the contract
	IStakingWalletRWalkFilterer   // Log filterer for contract events
}

// IStakingWalletRWalkCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingWalletRWalkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRWalkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingWalletRWalkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRWalkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingWalletRWalkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRWalkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingWalletRWalkSession struct {
	Contract     *IStakingWalletRWalk // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IStakingWalletRWalkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingWalletRWalkCallerSession struct {
	Contract *IStakingWalletRWalkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IStakingWalletRWalkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingWalletRWalkTransactorSession struct {
	Contract     *IStakingWalletRWalkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IStakingWalletRWalkRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingWalletRWalkRaw struct {
	Contract *IStakingWalletRWalk // Generic contract binding to access the raw methods on
}

// IStakingWalletRWalkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingWalletRWalkCallerRaw struct {
	Contract *IStakingWalletRWalkCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingWalletRWalkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingWalletRWalkTransactorRaw struct {
	Contract *IStakingWalletRWalkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakingWalletRWalk creates a new instance of IStakingWalletRWalk, bound to a specific deployed contract.
func NewIStakingWalletRWalk(address common.Address, backend bind.ContractBackend) (*IStakingWalletRWalk, error) {
	contract, err := bindIStakingWalletRWalk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRWalk{IStakingWalletRWalkCaller: IStakingWalletRWalkCaller{contract: contract}, IStakingWalletRWalkTransactor: IStakingWalletRWalkTransactor{contract: contract}, IStakingWalletRWalkFilterer: IStakingWalletRWalkFilterer{contract: contract}}, nil
}

// NewIStakingWalletRWalkCaller creates a new read-only instance of IStakingWalletRWalk, bound to a specific deployed contract.
func NewIStakingWalletRWalkCaller(address common.Address, caller bind.ContractCaller) (*IStakingWalletRWalkCaller, error) {
	contract, err := bindIStakingWalletRWalk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRWalkCaller{contract: contract}, nil
}

// NewIStakingWalletRWalkTransactor creates a new write-only instance of IStakingWalletRWalk, bound to a specific deployed contract.
func NewIStakingWalletRWalkTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingWalletRWalkTransactor, error) {
	contract, err := bindIStakingWalletRWalk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRWalkTransactor{contract: contract}, nil
}

// NewIStakingWalletRWalkFilterer creates a new log filterer instance of IStakingWalletRWalk, bound to a specific deployed contract.
func NewIStakingWalletRWalkFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingWalletRWalkFilterer, error) {
	contract, err := bindIStakingWalletRWalk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRWalkFilterer{contract: contract}, nil
}

// bindIStakingWalletRWalk binds a generic wrapper to an already deployed contract.
func bindIStakingWalletRWalk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakingWalletRWalkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletRWalk *IStakingWalletRWalkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletRWalk.Contract.IStakingWalletRWalkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletRWalk *IStakingWalletRWalkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.IStakingWalletRWalkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletRWalk *IStakingWalletRWalkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.IStakingWalletRWalkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletRWalk *IStakingWalletRWalkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletRWalk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletRWalk *IStakingWalletRWalkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletRWalk *IStakingWalletRWalkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.contract.Transact(opts, method, params...)
}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_IStakingWalletRWalk *IStakingWalletRWalkCaller) IsTokenStaked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _IStakingWalletRWalk.contract.Call(opts, &out, "isTokenStaked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_IStakingWalletRWalk *IStakingWalletRWalkSession) IsTokenStaked(tokenId *big.Int) (bool, error) {
	return _IStakingWalletRWalk.Contract.IsTokenStaked(&_IStakingWalletRWalk.CallOpts, tokenId)
}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_IStakingWalletRWalk *IStakingWalletRWalkCallerSession) IsTokenStaked(tokenId *big.Int) (bool, error) {
	return _IStakingWalletRWalk.Contract.IsTokenStaked(&_IStakingWalletRWalk.CallOpts, tokenId)
}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_IStakingWalletRWalk *IStakingWalletRWalkCaller) LastActionIdByTokenId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IStakingWalletRWalk.contract.Call(opts, &out, "lastActionIdByTokenId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_IStakingWalletRWalk *IStakingWalletRWalkSession) LastActionIdByTokenId(tokenId *big.Int) (*big.Int, error) {
	return _IStakingWalletRWalk.Contract.LastActionIdByTokenId(&_IStakingWalletRWalk.CallOpts, tokenId)
}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_IStakingWalletRWalk *IStakingWalletRWalkCallerSession) LastActionIdByTokenId(tokenId *big.Int) (*big.Int, error) {
	return _IStakingWalletRWalk.Contract.LastActionIdByTokenId(&_IStakingWalletRWalk.CallOpts, tokenId)
}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_IStakingWalletRWalk *IStakingWalletRWalkCaller) NumTokensStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingWalletRWalk.contract.Call(opts, &out, "numTokensStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_IStakingWalletRWalk *IStakingWalletRWalkSession) NumTokensStaked() (*big.Int, error) {
	return _IStakingWalletRWalk.Contract.NumTokensStaked(&_IStakingWalletRWalk.CallOpts)
}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_IStakingWalletRWalk *IStakingWalletRWalkCallerSession) NumTokensStaked() (*big.Int, error) {
	return _IStakingWalletRWalk.Contract.NumTokensStaked(&_IStakingWalletRWalk.CallOpts)
}

// PickRandomStaker is a free data retrieval call binding the contract method 0x41810425.
//
// Solidity: function pickRandomStaker(bytes32 entropy) view returns(address)
func (_IStakingWalletRWalk *IStakingWalletRWalkCaller) PickRandomStaker(opts *bind.CallOpts, entropy [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IStakingWalletRWalk.contract.Call(opts, &out, "pickRandomStaker", entropy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PickRandomStaker is a free data retrieval call binding the contract method 0x41810425.
//
// Solidity: function pickRandomStaker(bytes32 entropy) view returns(address)
func (_IStakingWalletRWalk *IStakingWalletRWalkSession) PickRandomStaker(entropy [32]byte) (common.Address, error) {
	return _IStakingWalletRWalk.Contract.PickRandomStaker(&_IStakingWalletRWalk.CallOpts, entropy)
}

// PickRandomStaker is a free data retrieval call binding the contract method 0x41810425.
//
// Solidity: function pickRandomStaker(bytes32 entropy) view returns(address)
func (_IStakingWalletRWalk *IStakingWalletRWalkCallerSession) PickRandomStaker(entropy [32]byte) (common.Address, error) {
	return _IStakingWalletRWalk.Contract.PickRandomStaker(&_IStakingWalletRWalk.CallOpts, entropy)
}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_IStakingWalletRWalk *IStakingWalletRWalkCaller) StakerByTokenId(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IStakingWalletRWalk.contract.Call(opts, &out, "stakerByTokenId", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_IStakingWalletRWalk *IStakingWalletRWalkSession) StakerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _IStakingWalletRWalk.Contract.StakerByTokenId(&_IStakingWalletRWalk.CallOpts, tokenId)
}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_IStakingWalletRWalk *IStakingWalletRWalkCallerSession) StakerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _IStakingWalletRWalk.Contract.StakerByTokenId(&_IStakingWalletRWalk.CallOpts, tokenId)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_IStakingWalletRWalk *IStakingWalletRWalkCaller) WasTokenUsed(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _IStakingWalletRWalk.contract.Call(opts, &out, "wasTokenUsed", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_IStakingWalletRWalk *IStakingWalletRWalkSession) WasTokenUsed(_tokenId *big.Int) (bool, error) {
	return _IStakingWalletRWalk.Contract.WasTokenUsed(&_IStakingWalletRWalk.CallOpts, _tokenId)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_IStakingWalletRWalk *IStakingWalletRWalkCallerSession) WasTokenUsed(_tokenId *big.Int) (bool, error) {
	return _IStakingWalletRWalk.Contract.WasTokenUsed(&_IStakingWalletRWalk.CallOpts, _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkTransactor) Stake(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.contract.Transact(opts, "stake", _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.Stake(&_IStakingWalletRWalk.TransactOpts, _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkTransactorSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.Stake(&_IStakingWalletRWalk.TransactOpts, _tokenId)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkTransactor) StakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.contract.Transact(opts, "stakeMany", ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.StakeMany(&_IStakingWalletRWalk.TransactOpts, ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkTransactorSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.StakeMany(&_IStakingWalletRWalk.TransactOpts, ids)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkTransactor) Unstake(opts *bind.TransactOpts, stakeActionId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.contract.Transact(opts, "unstake", stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.Unstake(&_IStakingWalletRWalk.TransactOpts, stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkTransactorSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.Unstake(&_IStakingWalletRWalk.TransactOpts, stakeActionId)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkTransactor) UnstakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.contract.Transact(opts, "unstakeMany", ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.UnstakeMany(&_IStakingWalletRWalk.TransactOpts, ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_IStakingWalletRWalk *IStakingWalletRWalkTransactorSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRWalk.Contract.UnstakeMany(&_IStakingWalletRWalk.TransactOpts, ids)
}

// IStakingWalletRWalkStakeActionEventIterator is returned from FilterStakeActionEvent and is used to iterate over the raw logs and unpacked data for StakeActionEvent events raised by the IStakingWalletRWalk contract.
type IStakingWalletRWalkStakeActionEventIterator struct {
	Event *IStakingWalletRWalkStakeActionEvent // Event containing the contract specifics and raw log

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
func (it *IStakingWalletRWalkStakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletRWalkStakeActionEvent)
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
		it.Event = new(IStakingWalletRWalkStakeActionEvent)
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
func (it *IStakingWalletRWalkStakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletRWalkStakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletRWalkStakeActionEvent represents a StakeActionEvent event raised by the IStakingWalletRWalk contract.
type IStakingWalletRWalkStakeActionEvent struct {
	ActionId  *big.Int
	TokenId   *big.Int
	TotalNFTs *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeActionEvent is a free log retrieval operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_IStakingWalletRWalk *IStakingWalletRWalkFilterer) FilterStakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*IStakingWalletRWalkStakeActionEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingWalletRWalk.contract.FilterLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRWalkStakeActionEventIterator{contract: _IStakingWalletRWalk.contract, event: "StakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchStakeActionEvent is a free log subscription operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_IStakingWalletRWalk *IStakingWalletRWalkFilterer) WatchStakeActionEvent(opts *bind.WatchOpts, sink chan<- *IStakingWalletRWalkStakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingWalletRWalk.contract.WatchLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletRWalkStakeActionEvent)
				if err := _IStakingWalletRWalk.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
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

// ParseStakeActionEvent is a log parse operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_IStakingWalletRWalk *IStakingWalletRWalkFilterer) ParseStakeActionEvent(log types.Log) (*IStakingWalletRWalkStakeActionEvent, error) {
	event := new(IStakingWalletRWalkStakeActionEvent)
	if err := _IStakingWalletRWalk.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletRWalkUnstakeActionEventIterator is returned from FilterUnstakeActionEvent and is used to iterate over the raw logs and unpacked data for UnstakeActionEvent events raised by the IStakingWalletRWalk contract.
type IStakingWalletRWalkUnstakeActionEventIterator struct {
	Event *IStakingWalletRWalkUnstakeActionEvent // Event containing the contract specifics and raw log

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
func (it *IStakingWalletRWalkUnstakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletRWalkUnstakeActionEvent)
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
		it.Event = new(IStakingWalletRWalkUnstakeActionEvent)
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
func (it *IStakingWalletRWalkUnstakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletRWalkUnstakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletRWalkUnstakeActionEvent represents a UnstakeActionEvent event raised by the IStakingWalletRWalk contract.
type IStakingWalletRWalkUnstakeActionEvent struct {
	ActionId  *big.Int
	TokenId   *big.Int
	TotalNFTs *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstakeActionEvent is a free log retrieval operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_IStakingWalletRWalk *IStakingWalletRWalkFilterer) FilterUnstakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*IStakingWalletRWalkUnstakeActionEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingWalletRWalk.contract.FilterLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRWalkUnstakeActionEventIterator{contract: _IStakingWalletRWalk.contract, event: "UnstakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchUnstakeActionEvent is a free log subscription operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_IStakingWalletRWalk *IStakingWalletRWalkFilterer) WatchUnstakeActionEvent(opts *bind.WatchOpts, sink chan<- *IStakingWalletRWalkUnstakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingWalletRWalk.contract.WatchLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletRWalkUnstakeActionEvent)
				if err := _IStakingWalletRWalk.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
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

// ParseUnstakeActionEvent is a log parse operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_IStakingWalletRWalk *IStakingWalletRWalkFilterer) ParseUnstakeActionEvent(log types.Log) (*IStakingWalletRWalkUnstakeActionEvent, error) {
	event := new(IStakingWalletRWalkUnstakeActionEvent)
	if err := _IStakingWalletRWalk.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRWalkMetaData contains all meta data concerning the StakingWalletRWalk contract.
var StakingWalletRWalkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"rwalk_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"AccessError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NoTokensStaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OneTimeStaking\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenAlreadyDeleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"TokenAlreadyInserted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"TokenAlreadyUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"StakeActionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"UnstakeActionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numStaked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isTokenStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"lastActionIdByTokenId\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastActionIds\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokensStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"entropy\",\"type\":\"bytes32\"}],\"name\":\"pickRandomStaker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalk\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"stakerByTokenId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"wasTokenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523461002f57610019610014610100565b610290565b610021610034565b611eb36104178239611eb390f35b61003a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100669061003e565b810190811060018060401b0382111761007e57604052565b610048565b9061009661008f610034565b928361005c565b565b5f80fd5b60018060a01b031690565b6100b09061009c565b90565b6100bc906100a7565b90565b6100c8816100b3565b036100cf57565b5f80fd5b905051906100e0826100bf565b565b906020828203126100fb576100f8915f016100d3565b90565b610098565b61011e6122ca8038038061011381610083565b9283398101906100e2565b90565b90565b61013861013361013d9261009c565b610121565b61009c565b90565b61014990610124565b90565b61015590610140565b90565b90565b61016f61016a61017492610158565b610121565b61009c565b90565b6101809061015b565b90565b60209181520190565b60207f616e646f6d57616c6b20746f6b656e2e00000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f722074686520525f8201520152565b6101e66030604092610183565b6101ef8161018c565b0190565b6102089060208101905f8183039101526101d9565b90565b1561021257565b61021a610034565b63eac0d38960e01b815280610231600482016101f3565b0390fd5b5f1b90565b9061024b60018060a01b0391610235565b9181191691161790565b61025e90610124565b90565b61026a90610255565b90565b90565b9061028561028061028c92610261565b61026d565b825461023a565b9055565b6102d19061029d336102d3565b6102ca6102a98261014c565b6102c36102bd6102b85f610177565b6100a7565b916100a7565b141561020b565b600a610270565b565b6102dc90610300565b565b6102e7906100a7565b9052565b91906102fe905f602085019401906102de565b565b8061031b6103156103105f610177565b6100a7565b916100a7565b1461032b57610329906103b7565b565b61034e6103375f610177565b5f918291631e4fbdf760e01b8352600483016102eb565b0390fd5b5f1c90565b60018060a01b031690565b61036e61037391610352565b610357565b90565b6103809054610362565b90565b61038c90610140565b90565b90565b906103a76103a26103ae92610383565b61038f565b825461023a565b9055565b5f0190565b6103c05f610376565b6103ca825f610392565b906103fe6103f87f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610383565b91610383565b91610407610034565b80610411816103b2565b0390a356fe60806040526004361015610013575b610cdc565b61001d5f3561018c565b80630d50c189146101875780630f7ee8791461018257806317db62131461017d5780632a3247aa146101785780632e17de7814610173578063418104251461016e57806344d110b914610169578063451f1adf146101645780635111a2d61461015f57806355279fdb1461015a5780635fda0acc146101555780636427d9a914610150578063715018a61461014b578063889d1e1a146101465780638da5cb5b14610141578063a2b136fb1461013c578063a531aa8614610137578063a694fc3a14610132578063c065894e1461012d578063c078855514610128578063f0a5242414610123578063f2fde38b1461011e5763fe939afc0361000e57610ca9565b610c76565b610c00565b610bcb565b610b21565b610aee565b610ab9565b610a71565b61097c565b610947565b610914565b6108df565b61082d565b6107f8565b6107b4565b6106e6565b610611565b6105ad565b6104fd565b6104c8565b610493565b6103f5565b6102ff565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906101d0906101a8565b810190811067ffffffffffffffff8211176101ea57604052565b6101b2565b906102026101fb610192565b92836101c6565b565b67ffffffffffffffff811161021c5760208091020190565b6101b2565b5f80fd5b90565b61023181610225565b0361023857565b5f80fd5b9050359061024982610228565b565b9092919261026061025b82610204565b6101ef565b938185526020808601920283019281841161029d57915b8383106102845750505050565b60208091610292848661023c565b815201920191610277565b610221565b9080601f830112156102c0578160206102bd9335910161024b565b90565b6101a4565b906020828203126102f5575f82013567ffffffffffffffff81116102f0576102ed92016102a2565b90565b6101a0565b61019c565b5f0190565b3461032d576103176103123660046102c5565b610d3f565b61031f610192565b80610329816102fa565b0390f35b610198565b9060208282031261034b57610348915f0161023c565b90565b61019c565b90565b61036761036261036c92610225565b610350565b610225565b90565b9061037990610353565b5f5260205260405f2090565b1c90565b60ff1690565b61039f9060086103a49302610385565b610389565b90565b906103b2915461038f565b90565b6103cb906103c66003915f9261036f565b6103a7565b90565b151590565b6103dc906103ce565b9052565b91906103f3905f602085019401906103d3565b565b346104255761042161041061040b366004610332565b6103b5565b610418610192565b918291826103e0565b0390f35b610198565b5f91031261043457565b61019c565b90565b61044c9060086104519302610385565b610439565b90565b9061045f915461043c565b90565b61046e60095f90610454565b90565b61047a90610225565b9052565b9190610491905f60208501940190610471565b565b346104c3576104a336600461042a565b6104bf6104ae610462565b6104b6610192565b9182918261047e565b0390f35b610198565b346104f8576104f46104e36104de366004610332565b610dbe565b6104eb610192565b918291826103e0565b0390f35b610198565b3461052b57610515610510366004610332565b611045565b61051d610192565b80610527816102fa565b0390f35b610198565b90565b61053c81610530565b0361054357565b5f80fd5b9050359061055482610533565b565b9060208282031261056f5761056c915f01610547565b90565b61019c565b60018060a01b031690565b61058890610574565b90565b6105949061057f565b9052565b91906105ab905f6020850194019061058b565b565b346105dd576105d96105c86105c3366004610556565b61133c565b6105d0610192565b91829182610598565b0390f35b610198565b906105ec90610353565b5f5260205260405f2090565b61060e906106096005915f926105e2565b610454565b90565b346106415761063d61062c610627366004610332565b6105f8565b610634610192565b9182918261047e565b0390f35b610198565b9061065090610353565b5f5260205260405f2090565b5f1c90565b61066d6106729161065c565b610439565b90565b61067f9054610661565b90565b61068d906007610646565b6106985f8201610675565b916106b160026106aa60018501610675565b9301610675565b90565b6040906106dd6106e494969593966106d360608401985f850190610471565b6020830190610471565b0190610471565b565b34610719576107156107016106fc366004610332565b610682565b61070c939193610192565b938493846106b4565b0390f35b610198565b60018060a01b031690565b61073990600861073e9302610385565b61071e565b90565b9061074c9154610729565b90565b61075b600a5f90610741565b90565b61077261076d61077792610574565b610350565b610574565b90565b6107839061075e565b90565b61078f9061077a565b90565b61079b90610786565b9052565b91906107b2905f60208501940190610792565b565b346107e4576107c436600461042a565b6107e06107cf61074f565b6107d7610192565b9182918261079f565b0390f35b610198565b6107f560085f90610454565b90565b346108285761080836600461042a565b6108246108136107e9565b61081b610192565b9182918261047e565b0390f35b610198565b3461085d5761083d36600461042a565b6108596108486113d2565b610850610192565b9182918261047e565b0390f35b610198565b9061086c90610353565b5f5260205260405f2090565b90565b61088b9060086108909302610385565b610878565b90565b9061089e915461087b565b90565b6108b7906108b26006915f92610862565b610893565b90565b90565b6108c6906108ba565b9052565b91906108dd905f602085019401906108bd565b565b3461090f5761090b6108fa6108f5366004610332565b6108a1565b610902610192565b918291826108ca565b0390f35b610198565b346109425761092436600461042a565b61092c611435565b610934610192565b8061093e816102fa565b0390f35b610198565b346109775761097361096261095d366004610332565b611462565b61096a610192565b918291826108ca565b0390f35b610198565b346109ac5761098c36600461042a565b6109a86109976114bb565b61099f610192565b91829182610598565b0390f35b610198565b906109bb90610353565b5f5260205260405f2090565b60018060a01b031690565b6109de6109e39161065c565b6109c7565b90565b6109f090546109d2565b90565b6109fe9060016109b1565b90610a0a5f8301610675565b91610a17600182016109e6565b91610a306003610a2960028501610675565b9301610675565b90565b610a68610a6f94610a5e606094989795610a54608086019a5f870190610471565b602085019061058b565b6040830190610471565b0190610471565b565b34610aa557610aa1610a8c610a87366004610332565b6109f3565b90610a98949294610192565b94859485610a33565b0390f35b610198565b610ab660025f90610454565b90565b34610ae957610ac936600461042a565b610ae5610ad4610aaa565b610adc610192565b9182918261047e565b0390f35b610198565b34610b1c57610b06610b01366004610332565b611632565b610b0e610192565b80610b18816102fa565b0390f35b610198565b34610b5157610b4d610b3c610b37366004610332565b611865565b610b44610192565b91829182610598565b0390f35b610198565b634e487b7160e01b5f52603260045260245ffd5b5490565b5f5260205f2090565b610b8081610b6a565b821015610b9a57610b92600191610b6e565b910201905f90565b610b56565b6004610baa81610b6a565b821015610bc757610bc491610bbe91610b77565b90610454565b90565b5f80fd5b34610bfb57610bf7610be6610be1366004610332565b610b9f565b610bee610192565b9182918261047e565b0390f35b610198565b34610c3057610c2c610c1b610c16366004610332565b6118be565b610c23610192565b918291826103e0565b0390f35b610198565b610c3e8161057f565b03610c4557565b5f80fd5b90503590610c5682610c35565b565b90602082820312610c7157610c6e915f01610c49565b90565b61019c565b34610ca457610c8e610c89366004610c58565b611956565b610c96610192565b80610ca0816102fa565b0390f35b610198565b34610cd757610cc1610cbc3660046102c5565b611961565b610cc9610192565b80610cd3816102fa565b0390f35b610198565b5f80fd5b90565b610cf7610cf2610cfc92610ce0565b610350565b610225565b90565b6001610d0b9101610225565b90565b5190565b90610d1c82610d0e565b811015610d2d576020809102010190565b610b56565b610d3c9051610225565b90565b90610d495f610ce3565b5b80610d65610d5f610d5a86610d0e565b610225565b91610225565b1015610d9457610d8f90610d8a610d85610d80868490610d12565b610d32565b611045565b610cff565b610d4a565b509050565b5f90565b610da9610dae9161065c565b610389565b90565b610dbb9054610d9d565b90565b610dd5610dda91610dcd610d99565b50600361036f565b610db1565b90565b60209181520190565b5f7f546f6b656e2068617320616c7265616479206265656e20756e7374616b65642e910152565b610e1960208092610ddd565b610e2281610de6565b0190565b9190610e49906020610e41604086018681035f880152610e0d565b940190610471565b565b15610e535750565b610e7590610e5f610192565b91829163aed59e4f60e01b835260048301610e26565b0390fd5b5f7f4f6e6c7920746865206f776e65722063616e20756e7374616b652e0000000000910152565b610ead601b602092610ddd565b610eb681610e79565b0190565b916040610eeb929493610ee4610ed9606083018381035f850152610ea0565b966020830190610471565b019061058b565b565b15610ef6575050565b610f17610f01610192565b9283926345c2e43b60e01b845260048401610eba565b0390fd5b610f27610f2c9161065c565b61071e565b90565b610f399054610f1b565b90565b610f459061077a565b90565b5f80fd5b60e01b90565b5f910312610f5c57565b61019c565b604090610f8a610f919496959396610f8060608401985f85019061058b565b602083019061058b565b0190610471565b565b610f9b610192565b3d5f823e3d90fd5b5f1b90565b90610fb45f1991610fa3565b9181191691161790565b90565b90610fd6610fd1610fdd92610353565b610fbe565b8254610fa8565b9055565b90565b610ff8610ff3610ffd92610fe1565b610350565b610225565b90565b634e487b7160e01b5f52601160045260245ffd5b61102361102991939293610225565b92610225565b820391821161103457565b611000565b6110429061077a565b90565b61107961105f6003611059600185906109b1565b01610675565b61107161106b5f610ce3565b91610225565b148290610e4b565b6110ad611092600161108c8185906109b1565b016109e6565b6110a461109e3361057f565b9161057f565b14823391610eed565b6110c35f6110bd600184906109b1565b01610675565b906110cd82611b87565b6110df6110da600a610f2f565b610786565b6323b872dd6110ed30610f3c565b33928592813b15611203575f61111691611121829661110a610192565b98899788968795610f4c565b855260048501610f61565b03925af180156111fe576111d2575b50611149426003611143600185906109b1565b01610fc1565b61116f6111686111596001610fe4565b6111636009610675565b611014565b6009610fc1565b6111796009610675565b9133916111cd6111bb6111b56111af7f33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c894610353565b94610353565b94611039565b946111c4610192565b9182918261047e565b0390a4565b6111f1905f3d81116111f7575b6111e981836101c6565b810190610f52565b5f611130565b503d6111df565b610f93565b610f48565b5f90565b60207f74616b65642e0000000000000000000000000000000000000000000000000000917f546865726520617265206e6f2052616e646f6d57616c6b20746f6b656e7320735f8201520152565b6112666026604092610ddd565b61126f8161120c565b0190565b6112889060208101905f818303910152611259565b90565b1561129257565b61129a610192565b63bc8b155960e01b8152806112b160048201611273565b0390fd5b6112c16112c69161065c565b610353565b90565b634e487b7160e01b5f52601260045260245ffd5b6112e96112ef91610225565b91610225565b9081156112fa570690565b6112c9565b61130b6113109161065c565b610878565b90565b61131d90546112ff565b90565b61133461132f611339926108ba565b610350565b610225565b90565b60016113c56113b56113b06113a96113a36113cb96611359611208565b5061137f6113676004610b6a565b6113796113735f610ce3565b91610225565b1161128b565b61139d61138d6004926112b5565b6113976004610b6a565b906112dd565b90610b77565b90610454565b6006610862565b611313565b6113bf8391611320565b906109b1565b016109e6565b90565b5f90565b6113da6113ce565b506113e56004610b6a565b90565b6113f0611c63565b6113f8611422565b565b61140e61140961141392610ce0565b610350565b610574565b90565b61141f906113fa565b90565b61143361142e5f611416565b611cb1565b565b61143d6113e8565b565b5f90565b90565b61145a61145561145f92611443565b610350565b6108ba565b90565b61146a61143f565b5061147f61147a600583906105e2565b610675565b61149161148b5f610ce3565b91610225565b146114ac576114a46114a9916006610862565b611313565b90565b506114b8600119611446565b90565b6114c3611208565b506114cd5f6109e6565b90565b60207f6564206f6e6c79206f6e63650000000000000000000000000000000000000000917f5374616b696e672f756e7374616b696e6720746f6b656e20697320616c6c6f775f8201520152565b61152a602c604092610ddd565b611533816114d0565b0190565b919061155a906020611552604086018681035f88015261151d565b940190610471565b565b156115645750565b61158690611570610192565b918291632290948760e21b835260048301611537565b0390fd5b9061159660ff91610fa3565b9181191691161790565b6115a9906103ce565b90565b90565b906115c46115bf6115cb926115a0565b6115ac565b825461158a565b9055565b906115e060018060a01b0391610fa3565b9181191691161790565b90565b906116026115fd61160992611039565b6115ea565b82546115cf565b9055565b61161c61162291939293610225565b92610225565b820180921161162d57565b611000565b61165961165261164c6116476003859061036f565b610db1565b156103ce565b829061155c565b61166f600161166a6003849061036f565b6115af565b61168161167c600a610f2f565b610786565b6323b872dd3361169030610f3c565b928492813b15611844575f6116b8916116c382966116ac610192565b98899788968795610f4c565b855260048501610f61565b03925af1801561183f57611813575b506116e7816116e16002610675565b90611e03565b611707815f61170160016116fb6002610675565b906109b1565b01610fc1565b6117273360016117218161171b6002610675565b906109b1565b016115ed565b611747426002611741600161173b83610675565b906109b1565b01610fc1565b61176d6117666117576001610fe4565b6117616002610675565b61160d565b6002610fc1565b61179361178c61177d6001610fe4565b6117876009610675565b61160d565b6009610fc1565b6117b06117a06002610675565b6117aa6001610fe4565b90611014565b6117ba6009610675565b91339161180e6117fc6117f66117f07fde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d94610353565b94610353565b94611039565b94611805610192565b9182918261047e565b0390a4565b611832905f3d8111611838575b61182a81836101c6565b810190610f52565b5f6116d2565b503d611820565b610f93565b610f48565b61185d61185861186292610ce0565b610350565b6108ba565b90565b61187790611871611208565b50611462565b8061188a6118845f611849565b916108ba565b126118b15760016118a86118ae926118a28391611320565b906109b1565b016109e6565b90565b506118bb5f611416565b90565b6118d56118da916118cd610d99565b5060056105e2565b610675565b6118ec6118e65f610ce3565b91610225565b141590565b611902906118fd611c63565b611904565b565b8061191f6119196119145f611416565b61057f565b9161057f565b1461192f5761192d90611cb1565b565b61195261193b5f611416565b5f918291631e4fbdf760e01b835260048301610598565b0390fd5b61195f906118f1565b565b9061196b5f610ce3565b5b8061198761198161197c86610d0e565b610225565b91610225565b10156119b6576119b1906119ac6119a76119a2868490610d12565b610d32565b611632565b610cff565b61196c565b509050565b5f7f546f6b656e206973206e6f7420696e20746865206c6973742e00000000000000910152565b6119ef6019602092610ddd565b6119f8816119bb565b0190565b9190611a1f906020611a17604086018681035f8801526119e2565b940190610471565b565b15611a295750565b611a4b90611a35610192565b918291639aa6fa6560e01b8352600483016119fc565b0390fd5b1b90565b91906008611a6e910291611a685f1984611a4f565b92611a4f565b9181191691161790565b9190611a8e611a89611a9693610353565b610fbe565b908354611a53565b9055565b611aac91611aa66113ce565b91611a78565b565b90565b634e487b7160e01b5f52603160045260245ffd5b5490565b5f5260205f2090565b611adb81611ac5565b821015611af557611aed600191611ac9565b910201905f90565b610b56565b611b0381611ac5565b8015611b24576001900390611b21611b1b8383611ad2565b90611a9a565b55565b611ab1565b90565b611b40611b3b611b4592611b29565b610350565b6108ba565b90565b611b5c611b57611b61926108ba565b610350565b6108ba565b90565b90565b90611b7c611b77611b8392611b48565b611b64565b8254610fa8565b9055565b611c6190611b9e611b97826118be565b8290611a21565b611c20611bb5611bb0600584906105e2565b610675565b611c1b611be9611be36004611bdd611bcd6004610b6a565b611bd76001610fe4565b90611014565b90610b77565b90610454565b91611c1383611c0d6004611c0785611c016001610fe4565b90611014565b90610b77565b90611a78565b9160056105e2565b610fc1565b611c355f611c30600584906105e2565b611a9a565b611c47611c426004611aae565b611afa565b611c5c611c545f19611b2c565b916006610862565b611b67565b565b611c6b6114bb565b611c84611c7e611c79611e70565b61057f565b9161057f565b03611c8b57565b611cad611c96611e70565b5f91829163118cdaa760e01b835260048301610598565b0390fd5b611cba5f6109e6565b611cc4825f6115ed565b90611cf8611cf27f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093611039565b91611039565b91611d01610192565b80611d0b816102fa565b0390a3565b5f7f546f6b656e20616c726561647920696e20746865206c6973742e000000000000910152565b611d44601a602092610ddd565b611d4d81611d10565b0190565b916040611d82929493611d7b611d70606083018381035f850152611d37565b966020830190610471565b0190610471565b565b15611d8d575050565b611dae611d98610192565b92839263597558c560e11b845260048401611d51565b0390fd5b9081549168010000000000000000831015611de25782611dda916001611de095018155611ad2565b90611a78565b565b6101b2565b611dfb611df6611e0092610225565b610350565b6108ba565b90565b611e69611e61611e6e93611e2a611e22611e1c866118be565b156103ce565b858391611d84565b611e3e611e376004611aae565b8590611db2565b611e5c611e4b6004610b6a565b611e57600587906105e2565b610fc1565b611de7565b916006610862565b611b67565b565b611e78611208565b50339056fea26469706673582212209ef5baf213221be07cc7515bb64a2e9e53f2d38e59fe52dad7bd90eed616975164736f6c634300081a0033",
}

// StakingWalletRWalkABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingWalletRWalkMetaData.ABI instead.
var StakingWalletRWalkABI = StakingWalletRWalkMetaData.ABI

// StakingWalletRWalkBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingWalletRWalkMetaData.Bin instead.
var StakingWalletRWalkBin = StakingWalletRWalkMetaData.Bin

// DeployStakingWalletRWalk deploys a new Ethereum contract, binding an instance of StakingWalletRWalk to it.
func DeployStakingWalletRWalk(auth *bind.TransactOpts, backend bind.ContractBackend, rwalk_ common.Address) (common.Address, *types.Transaction, *StakingWalletRWalk, error) {
	parsed, err := StakingWalletRWalkMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingWalletRWalkBin), backend, rwalk_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingWalletRWalk{StakingWalletRWalkCaller: StakingWalletRWalkCaller{contract: contract}, StakingWalletRWalkTransactor: StakingWalletRWalkTransactor{contract: contract}, StakingWalletRWalkFilterer: StakingWalletRWalkFilterer{contract: contract}}, nil
}

// StakingWalletRWalk is an auto generated Go binding around an Ethereum contract.
type StakingWalletRWalk struct {
	StakingWalletRWalkCaller     // Read-only binding to the contract
	StakingWalletRWalkTransactor // Write-only binding to the contract
	StakingWalletRWalkFilterer   // Log filterer for contract events
}

// StakingWalletRWalkCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingWalletRWalkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRWalkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingWalletRWalkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRWalkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingWalletRWalkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRWalkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingWalletRWalkSession struct {
	Contract     *StakingWalletRWalk // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StakingWalletRWalkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingWalletRWalkCallerSession struct {
	Contract *StakingWalletRWalkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// StakingWalletRWalkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingWalletRWalkTransactorSession struct {
	Contract     *StakingWalletRWalkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// StakingWalletRWalkRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingWalletRWalkRaw struct {
	Contract *StakingWalletRWalk // Generic contract binding to access the raw methods on
}

// StakingWalletRWalkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingWalletRWalkCallerRaw struct {
	Contract *StakingWalletRWalkCaller // Generic read-only contract binding to access the raw methods on
}

// StakingWalletRWalkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingWalletRWalkTransactorRaw struct {
	Contract *StakingWalletRWalkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingWalletRWalk creates a new instance of StakingWalletRWalk, bound to a specific deployed contract.
func NewStakingWalletRWalk(address common.Address, backend bind.ContractBackend) (*StakingWalletRWalk, error) {
	contract, err := bindStakingWalletRWalk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalk{StakingWalletRWalkCaller: StakingWalletRWalkCaller{contract: contract}, StakingWalletRWalkTransactor: StakingWalletRWalkTransactor{contract: contract}, StakingWalletRWalkFilterer: StakingWalletRWalkFilterer{contract: contract}}, nil
}

// NewStakingWalletRWalkCaller creates a new read-only instance of StakingWalletRWalk, bound to a specific deployed contract.
func NewStakingWalletRWalkCaller(address common.Address, caller bind.ContractCaller) (*StakingWalletRWalkCaller, error) {
	contract, err := bindStakingWalletRWalk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkCaller{contract: contract}, nil
}

// NewStakingWalletRWalkTransactor creates a new write-only instance of StakingWalletRWalk, bound to a specific deployed contract.
func NewStakingWalletRWalkTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingWalletRWalkTransactor, error) {
	contract, err := bindStakingWalletRWalk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkTransactor{contract: contract}, nil
}

// NewStakingWalletRWalkFilterer creates a new log filterer instance of StakingWalletRWalk, bound to a specific deployed contract.
func NewStakingWalletRWalkFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingWalletRWalkFilterer, error) {
	contract, err := bindStakingWalletRWalk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkFilterer{contract: contract}, nil
}

// bindStakingWalletRWalk binds a generic wrapper to an already deployed contract.
func bindStakingWalletRWalk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingWalletRWalkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletRWalk *StakingWalletRWalkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletRWalk.Contract.StakingWalletRWalkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletRWalk *StakingWalletRWalkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.StakingWalletRWalkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletRWalk *StakingWalletRWalkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.StakingWalletRWalkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletRWalk *StakingWalletRWalkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletRWalk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletRWalk *StakingWalletRWalkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletRWalk *StakingWalletRWalkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.contract.Transact(opts, method, params...)
}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint256 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) ETHDeposits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DepositTime   *big.Int
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "ETHDeposits", arg0)

	outstruct := new(struct {
		DepositTime   *big.Int
		DepositAmount *big.Int
		NumStaked     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DepositAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NumStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint256 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWalletRWalk *StakingWalletRWalkSession) ETHDeposits(arg0 *big.Int) (struct {
	DepositTime   *big.Int
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	return _StakingWalletRWalk.Contract.ETHDeposits(&_StakingWalletRWalk.CallOpts, arg0)
}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint256 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) ETHDeposits(arg0 *big.Int) (struct {
	DepositTime   *big.Int
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	return _StakingWalletRWalk.Contract.ETHDeposits(&_StakingWalletRWalk.CallOpts, arg0)
}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) IsTokenStaked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "isTokenStaked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkSession) IsTokenStaked(tokenId *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.IsTokenStaked(&_StakingWalletRWalk.CallOpts, tokenId)
}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) IsTokenStaked(tokenId *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.IsTokenStaked(&_StakingWalletRWalk.CallOpts, tokenId)
}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) LastActionIdByTokenId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "lastActionIdByTokenId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) LastActionIdByTokenId(tokenId *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.LastActionIdByTokenId(&_StakingWalletRWalk.CallOpts, tokenId)
}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) LastActionIdByTokenId(tokenId *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.LastActionIdByTokenId(&_StakingWalletRWalk.CallOpts, tokenId)
}

// LastActionIds is a free data retrieval call binding the contract method 0x6427d9a9.
//
// Solidity: function lastActionIds(uint256 ) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) LastActionIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "lastActionIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastActionIds is a free data retrieval call binding the contract method 0x6427d9a9.
//
// Solidity: function lastActionIds(uint256 ) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) LastActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.LastActionIds(&_StakingWalletRWalk.CallOpts, arg0)
}

// LastActionIds is a free data retrieval call binding the contract method 0x6427d9a9.
//
// Solidity: function lastActionIds(uint256 ) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) LastActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.LastActionIds(&_StakingWalletRWalk.CallOpts, arg0)
}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) NumETHDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "numETHDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) NumETHDeposits() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumETHDeposits(&_StakingWalletRWalk.CallOpts)
}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) NumETHDeposits() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumETHDeposits(&_StakingWalletRWalk.CallOpts)
}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) NumStakeActions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "numStakeActions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) NumStakeActions() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumStakeActions(&_StakingWalletRWalk.CallOpts)
}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) NumStakeActions() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumStakeActions(&_StakingWalletRWalk.CallOpts)
}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) NumStakedNFTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "numStakedNFTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) NumStakedNFTs() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumStakedNFTs(&_StakingWalletRWalk.CallOpts)
}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) NumStakedNFTs() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumStakedNFTs(&_StakingWalletRWalk.CallOpts)
}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) NumTokensStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "numTokensStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) NumTokensStaked() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumTokensStaked(&_StakingWalletRWalk.CallOpts)
}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) NumTokensStaked() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumTokensStaked(&_StakingWalletRWalk.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkSession) Owner() (common.Address, error) {
	return _StakingWalletRWalk.Contract.Owner(&_StakingWalletRWalk.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) Owner() (common.Address, error) {
	return _StakingWalletRWalk.Contract.Owner(&_StakingWalletRWalk.CallOpts)
}

// PickRandomStaker is a free data retrieval call binding the contract method 0x41810425.
//
// Solidity: function pickRandomStaker(bytes32 entropy) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) PickRandomStaker(opts *bind.CallOpts, entropy [32]byte) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "pickRandomStaker", entropy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PickRandomStaker is a free data retrieval call binding the contract method 0x41810425.
//
// Solidity: function pickRandomStaker(bytes32 entropy) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkSession) PickRandomStaker(entropy [32]byte) (common.Address, error) {
	return _StakingWalletRWalk.Contract.PickRandomStaker(&_StakingWalletRWalk.CallOpts, entropy)
}

// PickRandomStaker is a free data retrieval call binding the contract method 0x41810425.
//
// Solidity: function pickRandomStaker(bytes32 entropy) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) PickRandomStaker(entropy [32]byte) (common.Address, error) {
	return _StakingWalletRWalk.Contract.PickRandomStaker(&_StakingWalletRWalk.CallOpts, entropy)
}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) RandomWalk(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "randomWalk")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkSession) RandomWalk() (common.Address, error) {
	return _StakingWalletRWalk.Contract.RandomWalk(&_StakingWalletRWalk.CallOpts)
}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) RandomWalk() (common.Address, error) {
	return _StakingWalletRWalk.Contract.RandomWalk(&_StakingWalletRWalk.CallOpts)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address owner, uint256 stakeTime, uint256 unstakeTime)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) StakeActions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId     *big.Int
	Owner       common.Address
	StakeTime   *big.Int
	UnstakeTime *big.Int
}, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "stakeActions", arg0)

	outstruct := new(struct {
		TokenId     *big.Int
		Owner       common.Address
		StakeTime   *big.Int
		UnstakeTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StakeTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address owner, uint256 stakeTime, uint256 unstakeTime)
func (_StakingWalletRWalk *StakingWalletRWalkSession) StakeActions(arg0 *big.Int) (struct {
	TokenId     *big.Int
	Owner       common.Address
	StakeTime   *big.Int
	UnstakeTime *big.Int
}, error) {
	return _StakingWalletRWalk.Contract.StakeActions(&_StakingWalletRWalk.CallOpts, arg0)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address owner, uint256 stakeTime, uint256 unstakeTime)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) StakeActions(arg0 *big.Int) (struct {
	TokenId     *big.Int
	Owner       common.Address
	StakeTime   *big.Int
	UnstakeTime *big.Int
}, error) {
	return _StakingWalletRWalk.Contract.StakeActions(&_StakingWalletRWalk.CallOpts, arg0)
}

// StakedTokens is a free data retrieval call binding the contract method 0xc0788555.
//
// Solidity: function stakedTokens(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) StakedTokens(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "stakedTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedTokens is a free data retrieval call binding the contract method 0xc0788555.
//
// Solidity: function stakedTokens(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) StakedTokens(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.StakedTokens(&_StakingWalletRWalk.CallOpts, arg0)
}

// StakedTokens is a free data retrieval call binding the contract method 0xc0788555.
//
// Solidity: function stakedTokens(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) StakedTokens(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.StakedTokens(&_StakingWalletRWalk.CallOpts, arg0)
}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) StakerByTokenId(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "stakerByTokenId", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkSession) StakerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _StakingWalletRWalk.Contract.StakerByTokenId(&_StakingWalletRWalk.CallOpts, tokenId)
}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) StakerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _StakingWalletRWalk.Contract.StakerByTokenId(&_StakingWalletRWalk.CallOpts, tokenId)
}

// TokenIndices is a free data retrieval call binding the contract method 0x44d110b9.
//
// Solidity: function tokenIndices(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) TokenIndices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "tokenIndices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIndices is a free data retrieval call binding the contract method 0x44d110b9.
//
// Solidity: function tokenIndices(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) TokenIndices(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.TokenIndices(&_StakingWalletRWalk.CallOpts, arg0)
}

// TokenIndices is a free data retrieval call binding the contract method 0x44d110b9.
//
// Solidity: function tokenIndices(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) TokenIndices(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.TokenIndices(&_StakingWalletRWalk.CallOpts, arg0)
}

// UsedTokens is a free data retrieval call binding the contract method 0x0f7ee879.
//
// Solidity: function usedTokens(uint256 ) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) UsedTokens(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "usedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedTokens is a free data retrieval call binding the contract method 0x0f7ee879.
//
// Solidity: function usedTokens(uint256 ) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkSession) UsedTokens(arg0 *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.UsedTokens(&_StakingWalletRWalk.CallOpts, arg0)
}

// UsedTokens is a free data retrieval call binding the contract method 0x0f7ee879.
//
// Solidity: function usedTokens(uint256 ) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) UsedTokens(arg0 *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.UsedTokens(&_StakingWalletRWalk.CallOpts, arg0)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) WasTokenUsed(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "wasTokenUsed", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkSession) WasTokenUsed(_tokenId *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.WasTokenUsed(&_StakingWalletRWalk.CallOpts, _tokenId)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) WasTokenUsed(_tokenId *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.WasTokenUsed(&_StakingWalletRWalk.CallOpts, _tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.RenounceOwnership(&_StakingWalletRWalk.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.RenounceOwnership(&_StakingWalletRWalk.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) Stake(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "stake", _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.Stake(&_StakingWalletRWalk.TransactOpts, _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.Stake(&_StakingWalletRWalk.TransactOpts, _tokenId)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) StakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "stakeMany", ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.StakeMany(&_StakingWalletRWalk.TransactOpts, ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.StakeMany(&_StakingWalletRWalk.TransactOpts, ids)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.TransferOwnership(&_StakingWalletRWalk.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.TransferOwnership(&_StakingWalletRWalk.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) Unstake(opts *bind.TransactOpts, stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "unstake", stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.Unstake(&_StakingWalletRWalk.TransactOpts, stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.Unstake(&_StakingWalletRWalk.TransactOpts, stakeActionId)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) UnstakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "unstakeMany", ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.UnstakeMany(&_StakingWalletRWalk.TransactOpts, ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.UnstakeMany(&_StakingWalletRWalk.TransactOpts, ids)
}

// StakingWalletRWalkOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingWalletRWalk contract.
type StakingWalletRWalkOwnershipTransferredIterator struct {
	Event *StakingWalletRWalkOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingWalletRWalkOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRWalkOwnershipTransferred)
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
		it.Event = new(StakingWalletRWalkOwnershipTransferred)
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
func (it *StakingWalletRWalkOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRWalkOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRWalkOwnershipTransferred represents a OwnershipTransferred event raised by the StakingWalletRWalk contract.
type StakingWalletRWalkOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingWalletRWalkOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWalletRWalk.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkOwnershipTransferredIterator{contract: _StakingWalletRWalk.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingWalletRWalkOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWalletRWalk.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRWalkOwnershipTransferred)
				if err := _StakingWalletRWalk.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) ParseOwnershipTransferred(log types.Log) (*StakingWalletRWalkOwnershipTransferred, error) {
	event := new(StakingWalletRWalkOwnershipTransferred)
	if err := _StakingWalletRWalk.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRWalkStakeActionEventIterator is returned from FilterStakeActionEvent and is used to iterate over the raw logs and unpacked data for StakeActionEvent events raised by the StakingWalletRWalk contract.
type StakingWalletRWalkStakeActionEventIterator struct {
	Event *StakingWalletRWalkStakeActionEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletRWalkStakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRWalkStakeActionEvent)
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
		it.Event = new(StakingWalletRWalkStakeActionEvent)
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
func (it *StakingWalletRWalkStakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRWalkStakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRWalkStakeActionEvent represents a StakeActionEvent event raised by the StakingWalletRWalk contract.
type StakingWalletRWalkStakeActionEvent struct {
	ActionId  *big.Int
	TokenId   *big.Int
	TotalNFTs *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeActionEvent is a free log retrieval operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) FilterStakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*StakingWalletRWalkStakeActionEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWalletRWalk.contract.FilterLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkStakeActionEventIterator{contract: _StakingWalletRWalk.contract, event: "StakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchStakeActionEvent is a free log subscription operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) WatchStakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletRWalkStakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWalletRWalk.contract.WatchLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRWalkStakeActionEvent)
				if err := _StakingWalletRWalk.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
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

// ParseStakeActionEvent is a log parse operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) ParseStakeActionEvent(log types.Log) (*StakingWalletRWalkStakeActionEvent, error) {
	event := new(StakingWalletRWalkStakeActionEvent)
	if err := _StakingWalletRWalk.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRWalkUnstakeActionEventIterator is returned from FilterUnstakeActionEvent and is used to iterate over the raw logs and unpacked data for UnstakeActionEvent events raised by the StakingWalletRWalk contract.
type StakingWalletRWalkUnstakeActionEventIterator struct {
	Event *StakingWalletRWalkUnstakeActionEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletRWalkUnstakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRWalkUnstakeActionEvent)
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
		it.Event = new(StakingWalletRWalkUnstakeActionEvent)
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
func (it *StakingWalletRWalkUnstakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRWalkUnstakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRWalkUnstakeActionEvent represents a UnstakeActionEvent event raised by the StakingWalletRWalk contract.
type StakingWalletRWalkUnstakeActionEvent struct {
	ActionId  *big.Int
	TokenId   *big.Int
	TotalNFTs *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstakeActionEvent is a free log retrieval operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) FilterUnstakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*StakingWalletRWalkUnstakeActionEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWalletRWalk.contract.FilterLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkUnstakeActionEventIterator{contract: _StakingWalletRWalk.contract, event: "UnstakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchUnstakeActionEvent is a free log subscription operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) WatchUnstakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletRWalkUnstakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWalletRWalk.contract.WatchLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRWalkUnstakeActionEvent)
				if err := _StakingWalletRWalk.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
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

// ParseUnstakeActionEvent is a log parse operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) ParseUnstakeActionEvent(log types.Log) (*StakingWalletRWalkUnstakeActionEvent, error) {
	event := new(StakingWalletRWalkUnstakeActionEvent)
	if err := _StakingWalletRWalk.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
