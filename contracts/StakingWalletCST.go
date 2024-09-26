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

// IStakingWalletCSTMetaData contains all meta data concerning the IStakingWalletCST contract.
var IStakingWalletCSTMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"ClaimRewardEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNFTs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"StakeActionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"UnstakeActionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"actions\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"deposits\",\"type\":\"uint256[]\"}],\"name\":\"claimManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isTokenStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"lastActionIdByTokenId\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokensStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"stakerByTokenId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"transferRemainingBalanceToCharity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ETHDepositId\",\"type\":\"uint256\"}],\"name\":\"unstakeClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"unstake_actions\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claim_actions\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claim_deposits\",\"type\":\"uint256[]\"}],\"name\":\"unstakeClaimMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"wasTokenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IStakingWalletCSTABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingWalletCSTMetaData.ABI instead.
var IStakingWalletCSTABI = IStakingWalletCSTMetaData.ABI

// IStakingWalletCST is an auto generated Go binding around an Ethereum contract.
type IStakingWalletCST struct {
	IStakingWalletCSTCaller     // Read-only binding to the contract
	IStakingWalletCSTTransactor // Write-only binding to the contract
	IStakingWalletCSTFilterer   // Log filterer for contract events
}

// IStakingWalletCSTCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingWalletCSTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletCSTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingWalletCSTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletCSTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingWalletCSTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletCSTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingWalletCSTSession struct {
	Contract     *IStakingWalletCST // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IStakingWalletCSTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingWalletCSTCallerSession struct {
	Contract *IStakingWalletCSTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IStakingWalletCSTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingWalletCSTTransactorSession struct {
	Contract     *IStakingWalletCSTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IStakingWalletCSTRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingWalletCSTRaw struct {
	Contract *IStakingWalletCST // Generic contract binding to access the raw methods on
}

// IStakingWalletCSTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingWalletCSTCallerRaw struct {
	Contract *IStakingWalletCSTCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingWalletCSTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingWalletCSTTransactorRaw struct {
	Contract *IStakingWalletCSTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakingWalletCST creates a new instance of IStakingWalletCST, bound to a specific deployed contract.
func NewIStakingWalletCST(address common.Address, backend bind.ContractBackend) (*IStakingWalletCST, error) {
	contract, err := bindIStakingWalletCST(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCST{IStakingWalletCSTCaller: IStakingWalletCSTCaller{contract: contract}, IStakingWalletCSTTransactor: IStakingWalletCSTTransactor{contract: contract}, IStakingWalletCSTFilterer: IStakingWalletCSTFilterer{contract: contract}}, nil
}

// NewIStakingWalletCSTCaller creates a new read-only instance of IStakingWalletCST, bound to a specific deployed contract.
func NewIStakingWalletCSTCaller(address common.Address, caller bind.ContractCaller) (*IStakingWalletCSTCaller, error) {
	contract, err := bindIStakingWalletCST(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCSTCaller{contract: contract}, nil
}

// NewIStakingWalletCSTTransactor creates a new write-only instance of IStakingWalletCST, bound to a specific deployed contract.
func NewIStakingWalletCSTTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingWalletCSTTransactor, error) {
	contract, err := bindIStakingWalletCST(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCSTTransactor{contract: contract}, nil
}

// NewIStakingWalletCSTFilterer creates a new log filterer instance of IStakingWalletCST, bound to a specific deployed contract.
func NewIStakingWalletCSTFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingWalletCSTFilterer, error) {
	contract, err := bindIStakingWalletCST(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCSTFilterer{contract: contract}, nil
}

// bindIStakingWalletCST binds a generic wrapper to an already deployed contract.
func bindIStakingWalletCST(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakingWalletCSTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletCST *IStakingWalletCSTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletCST.Contract.IStakingWalletCSTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletCST *IStakingWalletCSTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.IStakingWalletCSTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletCST *IStakingWalletCSTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.IStakingWalletCSTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletCST *IStakingWalletCSTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletCST.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletCST *IStakingWalletCSTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletCST *IStakingWalletCSTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.contract.Transact(opts, method, params...)
}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTCaller) IsTokenStaked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _IStakingWalletCST.contract.Call(opts, &out, "isTokenStaked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTSession) IsTokenStaked(tokenId *big.Int) (bool, error) {
	return _IStakingWalletCST.Contract.IsTokenStaked(&_IStakingWalletCST.CallOpts, tokenId)
}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTCallerSession) IsTokenStaked(tokenId *big.Int) (bool, error) {
	return _IStakingWalletCST.Contract.IsTokenStaked(&_IStakingWalletCST.CallOpts, tokenId)
}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_IStakingWalletCST *IStakingWalletCSTCaller) LastActionIdByTokenId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IStakingWalletCST.contract.Call(opts, &out, "lastActionIdByTokenId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_IStakingWalletCST *IStakingWalletCSTSession) LastActionIdByTokenId(tokenId *big.Int) (*big.Int, error) {
	return _IStakingWalletCST.Contract.LastActionIdByTokenId(&_IStakingWalletCST.CallOpts, tokenId)
}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_IStakingWalletCST *IStakingWalletCSTCallerSession) LastActionIdByTokenId(tokenId *big.Int) (*big.Int, error) {
	return _IStakingWalletCST.Contract.LastActionIdByTokenId(&_IStakingWalletCST.CallOpts, tokenId)
}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_IStakingWalletCST *IStakingWalletCSTCaller) NumTokensStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingWalletCST.contract.Call(opts, &out, "numTokensStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_IStakingWalletCST *IStakingWalletCSTSession) NumTokensStaked() (*big.Int, error) {
	return _IStakingWalletCST.Contract.NumTokensStaked(&_IStakingWalletCST.CallOpts)
}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_IStakingWalletCST *IStakingWalletCSTCallerSession) NumTokensStaked() (*big.Int, error) {
	return _IStakingWalletCST.Contract.NumTokensStaked(&_IStakingWalletCST.CallOpts)
}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_IStakingWalletCST *IStakingWalletCSTCaller) StakerByTokenId(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IStakingWalletCST.contract.Call(opts, &out, "stakerByTokenId", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_IStakingWalletCST *IStakingWalletCSTSession) StakerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _IStakingWalletCST.Contract.StakerByTokenId(&_IStakingWalletCST.CallOpts, tokenId)
}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_IStakingWalletCST *IStakingWalletCSTCallerSession) StakerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _IStakingWalletCST.Contract.StakerByTokenId(&_IStakingWalletCST.CallOpts, tokenId)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTCaller) WasTokenUsed(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _IStakingWalletCST.contract.Call(opts, &out, "wasTokenUsed", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTSession) WasTokenUsed(_tokenId *big.Int) (bool, error) {
	return _IStakingWalletCST.Contract.WasTokenUsed(&_IStakingWalletCST.CallOpts, _tokenId)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTCallerSession) WasTokenUsed(_tokenId *big.Int) (bool, error) {
	return _IStakingWalletCST.Contract.WasTokenUsed(&_IStakingWalletCST.CallOpts, _tokenId)
}

// ClaimManyRewards is a paid mutator transaction binding the contract method 0x6034eb5b.
//
// Solidity: function claimManyRewards(uint256[] actions, uint256[] deposits) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) ClaimManyRewards(opts *bind.TransactOpts, actions []*big.Int, deposits []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "claimManyRewards", actions, deposits)
}

// ClaimManyRewards is a paid mutator transaction binding the contract method 0x6034eb5b.
//
// Solidity: function claimManyRewards(uint256[] actions, uint256[] deposits) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) ClaimManyRewards(actions []*big.Int, deposits []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.ClaimManyRewards(&_IStakingWalletCST.TransactOpts, actions, deposits)
}

// ClaimManyRewards is a paid mutator transaction binding the contract method 0x6034eb5b.
//
// Solidity: function claimManyRewards(uint256[] actions, uint256[] deposits) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) ClaimManyRewards(actions []*big.Int, deposits []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.ClaimManyRewards(&_IStakingWalletCST.TransactOpts, actions, deposits)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x326c7b69.
//
// Solidity: function depositIfPossible() payable returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) DepositIfPossible(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "depositIfPossible")
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x326c7b69.
//
// Solidity: function depositIfPossible() payable returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) DepositIfPossible() (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.DepositIfPossible(&_IStakingWalletCST.TransactOpts)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x326c7b69.
//
// Solidity: function depositIfPossible() payable returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) DepositIfPossible() (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.DepositIfPossible(&_IStakingWalletCST.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) Stake(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "stake", _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.Stake(&_IStakingWalletCST.TransactOpts, _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.Stake(&_IStakingWalletCST.TransactOpts, _tokenId)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) StakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "stakeMany", ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.StakeMany(&_IStakingWalletCST.TransactOpts, ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.StakeMany(&_IStakingWalletCST.TransactOpts, ids)
}

// TransferRemainingBalanceToCharity is a paid mutator transaction binding the contract method 0xed614ffa.
//
// Solidity: function transferRemainingBalanceToCharity(address charityAddress_) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) TransferRemainingBalanceToCharity(opts *bind.TransactOpts, charityAddress_ common.Address) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "transferRemainingBalanceToCharity", charityAddress_)
}

// TransferRemainingBalanceToCharity is a paid mutator transaction binding the contract method 0xed614ffa.
//
// Solidity: function transferRemainingBalanceToCharity(address charityAddress_) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) TransferRemainingBalanceToCharity(charityAddress_ common.Address) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.TransferRemainingBalanceToCharity(&_IStakingWalletCST.TransactOpts, charityAddress_)
}

// TransferRemainingBalanceToCharity is a paid mutator transaction binding the contract method 0xed614ffa.
//
// Solidity: function transferRemainingBalanceToCharity(address charityAddress_) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) TransferRemainingBalanceToCharity(charityAddress_ common.Address) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.TransferRemainingBalanceToCharity(&_IStakingWalletCST.TransactOpts, charityAddress_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) Unstake(opts *bind.TransactOpts, stakeActionId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "unstake", stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.Unstake(&_IStakingWalletCST.TransactOpts, stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.Unstake(&_IStakingWalletCST.TransactOpts, stakeActionId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0x1969e066.
//
// Solidity: function unstakeClaim(uint256 stakeActionId, uint256 ETHDepositId) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) UnstakeClaim(opts *bind.TransactOpts, stakeActionId *big.Int, ETHDepositId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "unstakeClaim", stakeActionId, ETHDepositId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0x1969e066.
//
// Solidity: function unstakeClaim(uint256 stakeActionId, uint256 ETHDepositId) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) UnstakeClaim(stakeActionId *big.Int, ETHDepositId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.UnstakeClaim(&_IStakingWalletCST.TransactOpts, stakeActionId, ETHDepositId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0x1969e066.
//
// Solidity: function unstakeClaim(uint256 stakeActionId, uint256 ETHDepositId) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) UnstakeClaim(stakeActionId *big.Int, ETHDepositId *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.UnstakeClaim(&_IStakingWalletCST.TransactOpts, stakeActionId, ETHDepositId)
}

// UnstakeClaimMany is a paid mutator transaction binding the contract method 0xdb2b4bd8.
//
// Solidity: function unstakeClaimMany(uint256[] unstake_actions, uint256[] claim_actions, uint256[] claim_deposits) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) UnstakeClaimMany(opts *bind.TransactOpts, unstake_actions []*big.Int, claim_actions []*big.Int, claim_deposits []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "unstakeClaimMany", unstake_actions, claim_actions, claim_deposits)
}

// UnstakeClaimMany is a paid mutator transaction binding the contract method 0xdb2b4bd8.
//
// Solidity: function unstakeClaimMany(uint256[] unstake_actions, uint256[] claim_actions, uint256[] claim_deposits) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) UnstakeClaimMany(unstake_actions []*big.Int, claim_actions []*big.Int, claim_deposits []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.UnstakeClaimMany(&_IStakingWalletCST.TransactOpts, unstake_actions, claim_actions, claim_deposits)
}

// UnstakeClaimMany is a paid mutator transaction binding the contract method 0xdb2b4bd8.
//
// Solidity: function unstakeClaimMany(uint256[] unstake_actions, uint256[] claim_actions, uint256[] claim_deposits) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) UnstakeClaimMany(unstake_actions []*big.Int, claim_actions []*big.Int, claim_deposits []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.UnstakeClaimMany(&_IStakingWalletCST.TransactOpts, unstake_actions, claim_actions, claim_deposits)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) UnstakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "unstakeMany", ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.UnstakeMany(&_IStakingWalletCST.TransactOpts, ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.UnstakeMany(&_IStakingWalletCST.TransactOpts, ids)
}

// IStakingWalletCSTClaimRewardEventIterator is returned from FilterClaimRewardEvent and is used to iterate over the raw logs and unpacked data for ClaimRewardEvent events raised by the IStakingWalletCST contract.
type IStakingWalletCSTClaimRewardEventIterator struct {
	Event *IStakingWalletCSTClaimRewardEvent // Event containing the contract specifics and raw log

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
func (it *IStakingWalletCSTClaimRewardEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletCSTClaimRewardEvent)
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
		it.Event = new(IStakingWalletCSTClaimRewardEvent)
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
func (it *IStakingWalletCSTClaimRewardEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletCSTClaimRewardEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletCSTClaimRewardEvent represents a ClaimRewardEvent event raised by the IStakingWalletCST contract.
type IStakingWalletCSTClaimRewardEvent struct {
	ActionId  *big.Int
	DepositId *big.Int
	Reward    *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimRewardEvent is a free log retrieval operation binding the contract event 0xdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36.
//
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address indexed staker)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) FilterClaimRewardEvent(opts *bind.FilterOpts, actionId []*big.Int, depositId []*big.Int, staker []common.Address) (*IStakingWalletCSTClaimRewardEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingWalletCST.contract.FilterLogs(opts, "ClaimRewardEvent", actionIdRule, depositIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCSTClaimRewardEventIterator{contract: _IStakingWalletCST.contract, event: "ClaimRewardEvent", logs: logs, sub: sub}, nil
}

// WatchClaimRewardEvent is a free log subscription operation binding the contract event 0xdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36.
//
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address indexed staker)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) WatchClaimRewardEvent(opts *bind.WatchOpts, sink chan<- *IStakingWalletCSTClaimRewardEvent, actionId []*big.Int, depositId []*big.Int, staker []common.Address) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingWalletCST.contract.WatchLogs(opts, "ClaimRewardEvent", actionIdRule, depositIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletCSTClaimRewardEvent)
				if err := _IStakingWalletCST.contract.UnpackLog(event, "ClaimRewardEvent", log); err != nil {
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

// ParseClaimRewardEvent is a log parse operation binding the contract event 0xdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36.
//
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address indexed staker)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) ParseClaimRewardEvent(log types.Log) (*IStakingWalletCSTClaimRewardEvent, error) {
	event := new(IStakingWalletCSTClaimRewardEvent)
	if err := _IStakingWalletCST.contract.UnpackLog(event, "ClaimRewardEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletCSTEthDepositEventIterator is returned from FilterEthDepositEvent and is used to iterate over the raw logs and unpacked data for EthDepositEvent events raised by the IStakingWalletCST contract.
type IStakingWalletCSTEthDepositEventIterator struct {
	Event *IStakingWalletCSTEthDepositEvent // Event containing the contract specifics and raw log

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
func (it *IStakingWalletCSTEthDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletCSTEthDepositEvent)
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
		it.Event = new(IStakingWalletCSTEthDepositEvent)
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
func (it *IStakingWalletCSTEthDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletCSTEthDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletCSTEthDepositEvent represents a EthDepositEvent event raised by the IStakingWalletCST contract.
type IStakingWalletCSTEthDepositEvent struct {
	DepositTime   *big.Int
	DepositNum    *big.Int
	NumStakedNFTs *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthDepositEvent is a free log retrieval operation binding the contract event 0xd91db13dbdc68f51d0e522c3f29f3d5c0f5945fa377792cb335e42f78b82d1d8.
//
// Solidity: event EthDepositEvent(uint256 indexed depositTime, uint256 depositNum, uint256 numStakedNFTs, uint256 amount)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) FilterEthDepositEvent(opts *bind.FilterOpts, depositTime []*big.Int) (*IStakingWalletCSTEthDepositEventIterator, error) {

	var depositTimeRule []interface{}
	for _, depositTimeItem := range depositTime {
		depositTimeRule = append(depositTimeRule, depositTimeItem)
	}

	logs, sub, err := _IStakingWalletCST.contract.FilterLogs(opts, "EthDepositEvent", depositTimeRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCSTEthDepositEventIterator{contract: _IStakingWalletCST.contract, event: "EthDepositEvent", logs: logs, sub: sub}, nil
}

// WatchEthDepositEvent is a free log subscription operation binding the contract event 0xd91db13dbdc68f51d0e522c3f29f3d5c0f5945fa377792cb335e42f78b82d1d8.
//
// Solidity: event EthDepositEvent(uint256 indexed depositTime, uint256 depositNum, uint256 numStakedNFTs, uint256 amount)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) WatchEthDepositEvent(opts *bind.WatchOpts, sink chan<- *IStakingWalletCSTEthDepositEvent, depositTime []*big.Int) (event.Subscription, error) {

	var depositTimeRule []interface{}
	for _, depositTimeItem := range depositTime {
		depositTimeRule = append(depositTimeRule, depositTimeItem)
	}

	logs, sub, err := _IStakingWalletCST.contract.WatchLogs(opts, "EthDepositEvent", depositTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletCSTEthDepositEvent)
				if err := _IStakingWalletCST.contract.UnpackLog(event, "EthDepositEvent", log); err != nil {
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

// ParseEthDepositEvent is a log parse operation binding the contract event 0xd91db13dbdc68f51d0e522c3f29f3d5c0f5945fa377792cb335e42f78b82d1d8.
//
// Solidity: event EthDepositEvent(uint256 indexed depositTime, uint256 depositNum, uint256 numStakedNFTs, uint256 amount)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) ParseEthDepositEvent(log types.Log) (*IStakingWalletCSTEthDepositEvent, error) {
	event := new(IStakingWalletCSTEthDepositEvent)
	if err := _IStakingWalletCST.contract.UnpackLog(event, "EthDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletCSTStakeActionEventIterator is returned from FilterStakeActionEvent and is used to iterate over the raw logs and unpacked data for StakeActionEvent events raised by the IStakingWalletCST contract.
type IStakingWalletCSTStakeActionEventIterator struct {
	Event *IStakingWalletCSTStakeActionEvent // Event containing the contract specifics and raw log

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
func (it *IStakingWalletCSTStakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletCSTStakeActionEvent)
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
		it.Event = new(IStakingWalletCSTStakeActionEvent)
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
func (it *IStakingWalletCSTStakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletCSTStakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletCSTStakeActionEvent represents a StakeActionEvent event raised by the IStakingWalletCST contract.
type IStakingWalletCSTStakeActionEvent struct {
	ActionId  *big.Int
	TokenId   *big.Int
	TotalNFTs *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeActionEvent is a free log retrieval operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) FilterStakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*IStakingWalletCSTStakeActionEventIterator, error) {

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

	logs, sub, err := _IStakingWalletCST.contract.FilterLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCSTStakeActionEventIterator{contract: _IStakingWalletCST.contract, event: "StakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchStakeActionEvent is a free log subscription operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) WatchStakeActionEvent(opts *bind.WatchOpts, sink chan<- *IStakingWalletCSTStakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IStakingWalletCST.contract.WatchLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletCSTStakeActionEvent)
				if err := _IStakingWalletCST.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
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
func (_IStakingWalletCST *IStakingWalletCSTFilterer) ParseStakeActionEvent(log types.Log) (*IStakingWalletCSTStakeActionEvent, error) {
	event := new(IStakingWalletCSTStakeActionEvent)
	if err := _IStakingWalletCST.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletCSTUnstakeActionEventIterator is returned from FilterUnstakeActionEvent and is used to iterate over the raw logs and unpacked data for UnstakeActionEvent events raised by the IStakingWalletCST contract.
type IStakingWalletCSTUnstakeActionEventIterator struct {
	Event *IStakingWalletCSTUnstakeActionEvent // Event containing the contract specifics and raw log

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
func (it *IStakingWalletCSTUnstakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletCSTUnstakeActionEvent)
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
		it.Event = new(IStakingWalletCSTUnstakeActionEvent)
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
func (it *IStakingWalletCSTUnstakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletCSTUnstakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletCSTUnstakeActionEvent represents a UnstakeActionEvent event raised by the IStakingWalletCST contract.
type IStakingWalletCSTUnstakeActionEvent struct {
	ActionId  *big.Int
	TokenId   *big.Int
	TotalNFTs *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstakeActionEvent is a free log retrieval operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) FilterUnstakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*IStakingWalletCSTUnstakeActionEventIterator, error) {

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

	logs, sub, err := _IStakingWalletCST.contract.FilterLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCSTUnstakeActionEventIterator{contract: _IStakingWalletCST.contract, event: "UnstakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchUnstakeActionEvent is a free log subscription operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) WatchUnstakeActionEvent(opts *bind.WatchOpts, sink chan<- *IStakingWalletCSTUnstakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IStakingWalletCST.contract.WatchLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletCSTUnstakeActionEvent)
				if err := _IStakingWalletCST.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
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
func (_IStakingWalletCST *IStakingWalletCSTFilterer) ParseUnstakeActionEvent(log types.Log) (*IStakingWalletCSTUnstakeActionEvent, error) {
	event := new(IStakingWalletCSTUnstakeActionEvent)
	if err := _IStakingWalletCST.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SignedMathMetaData contains all meta data concerning the SignedMath contract.
var SignedMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x608060405234601d57600e6021565b603e602c823930815050603e90f35b6027565b60405190565b5f80fdfe60806040525f80fdfea26469706673582212203e074f928186aac89344dd0e5a1ac07c72a5216f4065fdcf7508afeed89d2d2b64736f6c634300081a0033",
}

// SignedMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SignedMathMetaData.ABI instead.
var SignedMathABI = SignedMathMetaData.ABI

// SignedMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SignedMathMetaData.Bin instead.
var SignedMathBin = SignedMathMetaData.Bin

// DeploySignedMath deploys a new Ethereum contract, binding an instance of SignedMath to it.
func DeploySignedMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignedMath, error) {
	parsed, err := SignedMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SignedMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignedMath{SignedMathCaller: SignedMathCaller{contract: contract}, SignedMathTransactor: SignedMathTransactor{contract: contract}, SignedMathFilterer: SignedMathFilterer{contract: contract}}, nil
}

// SignedMath is an auto generated Go binding around an Ethereum contract.
type SignedMath struct {
	SignedMathCaller     // Read-only binding to the contract
	SignedMathTransactor // Write-only binding to the contract
	SignedMathFilterer   // Log filterer for contract events
}

// SignedMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignedMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignedMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignedMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignedMathSession struct {
	Contract     *SignedMath       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignedMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignedMathCallerSession struct {
	Contract *SignedMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SignedMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignedMathTransactorSession struct {
	Contract     *SignedMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SignedMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignedMathRaw struct {
	Contract *SignedMath // Generic contract binding to access the raw methods on
}

// SignedMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignedMathCallerRaw struct {
	Contract *SignedMathCaller // Generic read-only contract binding to access the raw methods on
}

// SignedMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignedMathTransactorRaw struct {
	Contract *SignedMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignedMath creates a new instance of SignedMath, bound to a specific deployed contract.
func NewSignedMath(address common.Address, backend bind.ContractBackend) (*SignedMath, error) {
	contract, err := bindSignedMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignedMath{SignedMathCaller: SignedMathCaller{contract: contract}, SignedMathTransactor: SignedMathTransactor{contract: contract}, SignedMathFilterer: SignedMathFilterer{contract: contract}}, nil
}

// NewSignedMathCaller creates a new read-only instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathCaller(address common.Address, caller bind.ContractCaller) (*SignedMathCaller, error) {
	contract, err := bindSignedMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignedMathCaller{contract: contract}, nil
}

// NewSignedMathTransactor creates a new write-only instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SignedMathTransactor, error) {
	contract, err := bindSignedMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignedMathTransactor{contract: contract}, nil
}

// NewSignedMathFilterer creates a new log filterer instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SignedMathFilterer, error) {
	contract, err := bindSignedMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignedMathFilterer{contract: contract}, nil
}

// bindSignedMath binds a generic wrapper to an already deployed contract.
func bindSignedMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SignedMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedMath *SignedMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedMath.Contract.SignedMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedMath *SignedMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedMath.Contract.SignedMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedMath *SignedMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedMath.Contract.SignedMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedMath *SignedMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedMath *SignedMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedMath *SignedMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedMath.Contract.contract.Transact(opts, method, params...)
}

// StakingWalletCSTMetaData contains all meta data concerning the StakingWalletCST contract.
var StakingWalletCSTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"AccessError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"}],\"name\":\"DepositAlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"DepositFromUnauthorizedSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositDate\",\"type\":\"uint256\"}],\"name\":\"DepositOutsideStakingWindow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionsLen\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositsLen\",\"type\":\"uint256\"}],\"name\":\"IncorrectArrayArguments\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"InvalidActionId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"}],\"name\":\"InvalidDepositId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperationInCurrentState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OneTimeStaking\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenAlreadyDeleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"TokenAlreadyInserted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"TokenAlreadyUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"TokenNotUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"ClaimRewardEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNFTs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"}],\"name\":\"FundsTransferredToCharityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"StakeActionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"UnstakeActionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numStaked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"actions\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"deposits\",\"type\":\"uint256[]\"}],\"name\":\"claimManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isTokenStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"lastActionIdByTokenId\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastActionIds\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokensStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"stakerByTokenId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"transferRemainingBalanceToCharity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ETHDepositId\",\"type\":\"uint256\"}],\"name\":\"unstakeClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"unstake_actions\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claim_actions\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claim_deposits\",\"type\":\"uint256[]\"}],\"name\":\"unstakeClaimMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"wasTokenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052346100305761001a610014610133565b9061039c565b610022610035565b61333b610524823961333b90f35b61003b565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100679061003f565b810190811060018060401b0382111761007f57604052565b610049565b90610097610090610035565b928361005d565b565b5f80fd5b60018060a01b031690565b6100b19061009d565b90565b6100bd906100a8565b90565b6100c9816100b4565b036100d057565b5f80fd5b905051906100e1826100c0565b565b6100ec816100a8565b036100f357565b5f80fd5b90505190610104826100e3565b565b919060408382031261012e578061012261012b925f86016100d4565b936020016100f7565b90565b610099565b61015161385f8038038061014681610084565b928339810190610106565b9091565b90565b61016c6101676101719261009d565b610155565b61009d565b90565b61017d90610158565b90565b61018990610174565b90565b90565b6101a361019e6101a89261018c565b610155565b61009d565b90565b6101b49061018f565b90565b60209181520190565b60207f66742e0000000000000000000000000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f7220746865206e5f8201520152565b61021a60236040926101b7565b610223816101c0565b0190565b61023c9060208101905f81830391015261020d565b90565b1561024657565b61024e610035565b63eac0d38960e01b81528061026560048201610227565b0390fd5b60207f616d652e00000000000000000000000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f722074686520675f8201520152565b6102c360246040926101b7565b6102cc81610269565b0190565b6102e59060208101905f8183039101526102b6565b90565b156102ef57565b6102f7610035565b63eac0d38960e01b81528061030e600482016102d0565b0390fd5b5f1b90565b9061032860018060a01b0391610312565b9181191691161790565b61033b90610158565b90565b61034790610332565b90565b90565b9061036261035d6103699261033e565b61034a565b8254610317565b9055565b61037690610174565b90565b90565b9061039161038c6103989261036d565b610379565b8254610317565b9055565b9061040661040d926103ad3361040f565b6103da6103b982610180565b6103d36103cd6103c85f6101ab565b6100a8565b916100a8565b141561023f565b6103ff836103f86103f26103ed5f6101ab565b6100a8565b916100a8565b14156102e8565b600161034d565b600261037c565b565b6104189061043c565b565b610423906100a8565b9052565b919061043a905f6020850194019061041a565b565b8061045761045161044c5f6101ab565b6100a8565b916100a8565b1461046757610465906104c4565b565b61048a6104735f6101ab565b5f918291631e4fbdf760e01b835260048301610427565b0390fd5b5f1c90565b60018060a01b031690565b6104aa6104af9161048e565b610493565b90565b6104bc905461049e565b90565b5f0190565b6104cd5f6104b2565b6104d7825f61037c565b9061050b6105057f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09361036d565b9161036d565b91610514610035565b8061051e816104bf565b0390a356fe60806040526004361015610013575b610f20565b61001d5f356101dc565b80630d50c189146101d75780630f7ee879146101d257806317db6213146101cd5780631969e066146101c85780632a3247aa146101c35780632e17de78146101be578063326c7b69146101b957806344d110b9146101b4578063451f1adf146101af57806347ccca02146101aa57806355279fdb146101a55780635fda0acc146101a05780636034eb5b1461019b5780636427d9a914610196578063715018a614610191578063889d1e1a1461018c5780638da5cb5b14610187578063a2b136fb14610182578063a531aa861461017d578063a694fc3a14610178578063c065894e14610173578063c07885551461016e578063c3fe3e2814610169578063db2b4bd814610164578063ed614ffa1461015f578063f0a524241461015a578063f2fde38b146101555763fe939afc0361000e57610eed565b610eba565b610e85565b610e52565b610ddd565b610d27565b610cbd565b610c13565b610be0565b610bab565b610b63565b610a6e565b610a0b565b6109d8565b6109a3565b6108f2565b610860565b61082b565b6107e7565b61070e565b610639565b6105e1565b6105ae565b610579565b610545565b6104e3565b610445565b61034f565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610220906101f8565b810190811067ffffffffffffffff82111761023a57604052565b610202565b9061025261024b6101e2565b9283610216565b565b67ffffffffffffffff811161026c5760208091020190565b610202565b5f80fd5b90565b61028181610275565b0361028857565b5f80fd5b9050359061029982610278565b565b909291926102b06102ab82610254565b61023f565b93818552602080860192028301928184116102ed57915b8383106102d45750505050565b602080916102e2848661028c565b8152019201916102c7565b610271565b9080601f830112156103105781602061030d9335910161029b565b90565b6101f4565b90602082820312610345575f82013567ffffffffffffffff81116103405761033d92016102f2565b90565b6101f0565b6101ec565b5f0190565b3461037d57610367610362366004610315565b610f83565b61036f6101e2565b806103798161034a565b0390f35b6101e8565b9060208282031261039b57610398915f0161028c565b90565b6101ec565b90565b6103b76103b26103bc92610275565b6103a0565b610275565b90565b906103c9906103a3565b5f5260205260405f2090565b1c90565b60ff1690565b6103ef9060086103f493026103d5565b6103d9565b90565b9061040291546103df565b90565b61041b906104166005915f926103bf565b6103f7565b90565b151590565b61042c9061041e565b9052565b9190610443905f60208501940190610423565b565b346104755761047161046061045b366004610382565b610405565b6104686101e2565b91829182610430565b0390f35b6101e8565b5f91031261048457565b6101ec565b90565b61049c9060086104a193026103d5565b610489565b90565b906104af915461048c565b90565b6104be60075f906104a4565b90565b6104ca90610275565b9052565b91906104e1905f602085019401906104c1565b565b34610513576104f336600461047a565b61050f6104fe6104b2565b6105066101e2565b918291826104ce565b0390f35b6101e8565b9190604083820312610540578061053461053d925f860161028c565b9360200161028c565b90565b6101ec565b346105745761055e610558366004610518565b90611110565b6105666101e2565b806105708161034a565b0390f35b6101e8565b346105a9576105a561059461058f366004610382565b61119a565b61059c6101e2565b91829182610430565b0390f35b6101e8565b346105dc576105c66105c1366004610382565b61140c565b6105ce6101e2565b806105d88161034a565b0390f35b6101e8565b6105ec36600461047a565b6105f461170c565b6105fc6101e2565b806106068161034a565b0390f35b90610614906103a3565b5f5260205260405f2090565b610636906106316008915f9261060a565b6104a4565b90565b346106695761066561065461064f366004610382565b610620565b61065c6101e2565b918291826104ce565b0390f35b6101e8565b90610678906103a3565b5f5260205260405f2090565b5f1c90565b61069561069a91610684565b610489565b90565b6106a79054610689565b90565b6106b590600a61066e565b6106c05f820161069d565b916106d960026106d26001850161069d565b930161069d565b90565b60409061070561070c94969593966106fb60608401985f8501906104c1565b60208301906104c1565b01906104c1565b565b346107415761073d610729610724366004610382565b6106aa565b6107349391936101e2565b938493846106dc565b0390f35b6101e8565b60018060a01b031690565b61076190600861076693026103d5565b610746565b90565b906107749154610751565b90565b61078360015f90610769565b90565b60018060a01b031690565b6107a56107a06107aa92610786565b6103a0565b610786565b90565b6107b690610791565b90565b6107c2906107ad565b90565b6107ce906107b9565b9052565b91906107e5905f602085019401906107c5565b565b34610817576107f736600461047a565b610813610802610777565b61080a6101e2565b918291826107d2565b0390f35b6101e8565b610828600b5f906104a4565b90565b3461085b5761083b36600461047a565b61085761084661081c565b61084e6101e2565b918291826104ce565b0390f35b6101e8565b346108905761087036600461047a565b61088c61087b611962565b6108836101e2565b918291826104ce565b0390f35b6101e8565b9190916040818403126108ed575f81013567ffffffffffffffff81116108e857836108c19183016102f2565b92602082013567ffffffffffffffff81116108e3576108e092016102f2565b90565b6101f0565b6101f0565b6101ec565b346109215761090b610905366004610895565b90611b05565b6109136101e2565b8061091d8161034a565b0390f35b6101e8565b90610930906103a3565b5f5260205260405f2090565b90565b61094f90600861095493026103d5565b61093c565b90565b90610962915461093f565b90565b61097b906109766009915f92610926565b610957565b90565b90565b61098a9061097e565b9052565b91906109a1905f60208501940190610981565b565b346109d3576109cf6109be6109b9366004610382565b610965565b6109c66101e2565b9182918261098e565b0390f35b6101e8565b34610a06576109e836600461047a565b6109f0611c88565b6109f86101e2565b80610a028161034a565b0390f35b6101e8565b34610a3b57610a37610a26610a21366004610382565b611cd6565b610a2e6101e2565b9182918261098e565b0390f35b6101e8565b610a4990610786565b90565b610a5590610a40565b9052565b9190610a6c905f60208501940190610a4c565b565b34610a9e57610a7e36600461047a565b610a9a610a89611d33565b610a916101e2565b91829182610a59565b0390f35b6101e8565b90610aad906103a3565b5f5260205260405f2090565b60018060a01b031690565b610ad0610ad591610684565b610ab9565b90565b610ae29054610ac4565b90565b610af0906003610aa3565b90610afc5f830161069d565b91610b0960018201610ad8565b91610b226003610b1b6002850161069d565b930161069d565b90565b610b5a610b6194610b50606094989795610b46608086019a5f8701906104c1565b6020850190610a4c565b60408301906104c1565b01906104c1565b565b34610b9757610b93610b7e610b79366004610382565b610ae5565b90610b8a9492946101e2565b94859485610b25565b0390f35b6101e8565b610ba860045f906104a4565b90565b34610bdb57610bbb36600461047a565b610bd7610bc6610b9c565b610bce6101e2565b918291826104ce565b0390f35b6101e8565b34610c0e57610bf8610bf3366004610382565b611e85565b610c006101e2565b80610c0a8161034a565b0390f35b6101e8565b34610c4357610c3f610c2e610c29366004610382565b6120af565b610c366101e2565b91829182610a59565b0390f35b6101e8565b634e487b7160e01b5f52603260045260245ffd5b5490565b5f5260205f2090565b610c7281610c5c565b821015610c8c57610c84600191610c60565b910201905f90565b610c48565b6006610c9c81610c5c565b821015610cb957610cb691610cb091610c69565b906104a4565b90565b5f80fd5b34610ced57610ce9610cd8610cd3366004610382565b610c91565b610ce06101e2565b918291826104ce565b0390f35b6101e8565b610d02906008610d0793026103d5565b610ab9565b90565b90610d159154610cf2565b90565b610d2460025f90610d0a565b90565b34610d5757610d3736600461047a565b610d53610d42610d18565b610d4a6101e2565b91829182610a59565b0390f35b6101e8565b91606083830312610dd8575f83013567ffffffffffffffff8111610dd35782610d869185016102f2565b92602081013567ffffffffffffffff8111610dce5783610da79183016102f2565b92604082013567ffffffffffffffff8111610dc957610dc692016102f2565b90565b6101f0565b6101f0565b6101f0565b6101ec565b34610e0c57610df6610df0366004610d5c565b916121d1565b610dfe6101e2565b80610e088161034a565b0390f35b6101e8565b610e1a81610a40565b03610e2157565b5f80fd5b90503590610e3282610e11565b565b90602082820312610e4d57610e4a915f01610e25565b90565b6101ec565b34610e8057610e6a610e65366004610e34565b612584565b610e726101e2565b80610e7c8161034a565b0390f35b6101e8565b34610eb557610eb1610ea0610e9b366004610382565b61258f565b610ea86101e2565b91829182610430565b0390f35b6101e8565b34610ee857610ed2610ecd366004610e34565b612627565b610eda6101e2565b80610ee48161034a565b0390f35b6101e8565b34610f1b57610f05610f00366004610315565b612632565b610f0d6101e2565b80610f178161034a565b0390f35b6101e8565b5f80fd5b90565b610f3b610f36610f4092610f24565b6103a0565b610275565b90565b6001610f4f9101610275565b90565b5190565b90610f6082610f52565b811015610f71576020809102010190565b610c48565b610f809051610275565b90565b90610f8d5f610f27565b5b80610fa9610fa3610f9e86610f52565b610275565b91610275565b1015610fd857610fd390610fce610fc9610fc4868490610f56565b610f76565b61140c565b610f43565b610f8e565b509050565b905090565b610fed5f8092610fdd565b0190565b610ffa90610fe2565b90565b67ffffffffffffffff811161101b576110176020916101f8565b0190565b610202565b9061103261102d83610ffd565b61023f565b918252565b606090565b3d5f146110575761104c3d611020565b903d5f602084013e5b565b61105f611037565b90611055565b60209181520190565b5f7f526577617264207472616e73666572206661696c65642e000000000000000000910152565b6110a26017602092611065565b6110ab8161106e565b0190565b9160406110e09294936110d96110ce606083018381035f850152611095565b9660208301906104c1565b0190610a4c565b565b156110eb575050565b61110c6110f66101e2565b92839263310a0fbb60e21b8452600484016110af565b0390fd5b906111239161111e8161140c565b612b97565b806111366111305f610f27565b91610275565b1161113f575b50565b61116f905f80338361114f6101e2565b908161115a81610ff1565b03925af161116661103c565b509033916110e2565b5f61113c565b5f90565b61118561118a91610684565b6103d9565b90565b6111979054611179565b90565b6111b16111b6916111a9611175565b5060056103bf565b61118d565b90565b5f7f546f6b656e2068617320616c7265616479206265656e20756e7374616b65642e910152565b6111ec60208092611065565b6111f5816111b9565b0190565b919061121c906020611214604086018681035f8801526111e0565b9401906104c1565b565b156112265750565b611248906112326101e2565b91829163aed59e4f60e01b8352600483016111f9565b0390fd5b5f7f4f6e6c7920746865206f776e65722063616e20756e7374616b652e0000000000910152565b611280601b602092611065565b6112898161124c565b0190565b9160406112be9294936112b76112ac606083018381035f850152611273565b9660208301906104c1565b0190610a4c565b565b156112c9575050565b6112ea6112d46101e2565b9283926345c2e43b60e01b84526004840161128d565b0390fd5b5f1b90565b906112ff5f19916112ee565b9181191691161790565b90565b9061132161131c611328926103a3565b611309565b82546112f3565b9055565b634e487b7160e01b5f52601160045260245ffd5b61134990610275565b5f8114611357576001900390565b61132c565b61136861136d91610684565b610746565b90565b61137a905461135c565b90565b611386906107ad565b90565b5f80fd5b60e01b90565b5f91031261139d57565b6101ec565b6040906113cb6113d294969593966113c160608401985f850190610a4c565b6020830190610a4c565b01906104c1565b565b6113dc6101e2565b3d5f823e3d90fd5b6113ed90610275565b5f1981146113fb5760010190565b61132c565b611409906107ad565b90565b61143f611425600361141f818590610aa3565b0161069d565b6114376114315f610f27565b91610275565b14829061121e565b611474611459600161145360038590610aa3565b01610ad8565b61146b61146533610a40565b91610a40565b148233916112c0565b61148a5f61148460038490610aa3565b0161069d565b906114948261302b565b6114b46114a1600c61069d565b60036114ae818590610aa3565b0161130c565b6114d06114c96114c4600761069d565b611340565b600761130c565b6114e26114dd6001611370565b6107b9565b6323b872dd6114f03061137d565b33928592813b156115e4575f61151991611524829661150d6101e2565b9889978896879561138d565b8552600485016113a2565b03925af180156115df576115b3575b50611550611549611544600c61069d565b6113e4565b600c61130c565b61155a600761069d565b9133916115ae61159c6115966115907f33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8946103a3565b946103a3565b94611400565b946115a56101e2565b918291826104ce565b0390a4565b6115d2905f3d81116115d8575b6115ca8183610216565b810190611393565b5f611533565b503d6115c0565b6113d4565b611389565b60207f206465706f7369742e0000000000000000000000000000000000000000000000917f4f6e6c792074686520436f736d696347616d6520636f6e74726163742063616e5f8201520152565b6116436029604092611065565b61164c816115e9565b0190565b919061167390602061166b604086018681035f880152611636565b940190610a4c565b565b1561167d5750565b61169f906116896101e2565b918291637ed5977760e11b835260048301611650565b0390fd5b90565b6116ba6116b56116bf926116a3565b6103a0565b610275565b90565b6116d16116d791939293610275565b92610275565b82039182116116e257565b61132c565b6116f66116fc91939293610275565b92610275565b820180921161170757565b61132c565b6117333361172b6117256117206002610ad8565b610a40565b91610a40565b143390611675565b61173d600761069d565b61174f6117495f610f27565b91610275565b146119425761175e600b61069d565b61177061176a5f610f27565b91610275565b11806118ed575b5f14611858576117c3346117bd60016117ad600a6117a7611798600b61069d565b6117a1856116a6565b906116c2565b9061066e565b01916117b88361069d565b6116e7565b9061130c565b5b6117e06117d96117d4600c61069d565b6113e4565b600c61130c565b6117ea600c61069d565b6118076117f7600b61069d565b61180160016116a6565b906116c2565b90611812600761069d565b34926118536118417fd91db13dbdc68f51d0e522c3f29f3d5c0f5945fa377792cb335e42f78b82d1d8946103a3565b9461184a6101e2565b938493846106dc565b0390a2565b611881611865600c61069d565b5f61187b600a611875600b61069d565b9061066e565b0161130c565b6118a234600161189c600a611896600b61069d565b9061066e565b0161130c565b6118cc6118af600761069d565b60026118c6600a6118c0600b61069d565b9061066e565b0161130c565b6118e86118e16118dc600b61069d565b6113e4565b600b61130c565b6117c4565b50611921600261191b600a611915611905600b61069d565b61190f60016116a6565b906116c2565b9061066e565b0161069d565b61193c611936611931600761069d565b610275565b91610275565b14611777565b5f632d6ebd6f60e01b81528061195a6004820161034a565b0390fd5b5f90565b61196a61195e565b506119756006610c5c565b90565b60207f616d65206c656e6774682e000000000000000000000000000000000000000000917f417272617920617267756d656e7473206d757374206265206f662074686520735f8201520152565b6119d2602b604092611065565b6119db81611978565b0190565b916040611a10929493611a096119fe606083018381035f8501526119c5565b9660208301906104c1565b01906104c1565b565b15611a1b575050565b611a3c611a266101e2565b92839263672c0fa160e11b8452600484016119df565b0390fd5b611a54611a4f611a5992610275565b6103a0565b61097e565b90565b611a70611a6b611a75926116a3565b6103a0565b61097e565b90565b611a87611a8d9193929361097e565b9261097e565b91828103925f828512818312169285139112151617611aa857565b61132c565b611ab69061097e565b600160ff1b8114611ac8576001900390565b61132c565b611ae1611adc611ae692610f24565b6103a0565b61097e565b90565b611afd611af8611b029261097e565b6103a0565b610275565b90565b90611b45611b1283610f52565b611b2c611b26611b2185610f52565b610275565b91610275565b14611b3684610f52565b611b3f84610f52565b91611a12565b611b4e5f610f27565b91611b73611b63611b5e83610f52565b611a40565b611b6d6001611a5c565b90611a78565b925b83611b88611b825f611acd565b9161097e565b12611be557611bd9611bdf91611bd3611bb2611bad86611ba78a611ae9565b90610f56565b610f76565b611bcd611bc888611bc28b611ae9565b90610f56565b610f76565b90612b97565b906116e7565b93611aad565b92611b75565b9250505080611bfc611bf65f610f27565b91610275565b11611c05575b50565b611c35905f803383611c156101e2565b9081611c2081610ff1565b03925af1611c2c61103c565b509033916110e2565b5f611c02565b611c43613107565b611c4b611c75565b565b611c61611c5c611c6692610f24565b6103a0565b610786565b90565b611c7290611c4d565b90565b611c86611c815f611c69565b613155565b565b611c90611c3b565b565b5f90565b90565b611cad611ca8611cb292611c96565b6103a0565b61097e565b90565b611cc1611cc691610684565b61093c565b90565b611cd39054611cb5565b90565b611cde611c92565b50611cf3611cee6008839061060a565b61069d565b611d05611cff5f610f27565b91610275565b14611d2057611d18611d1d916009610926565b611cc9565b90565b50611d2c600119611c99565b90565b5f90565b611d3b611d2f565b50611d455f610ad8565b90565b60207f6564206f6e6c79206f6e63650000000000000000000000000000000000000000917f5374616b696e672f756e7374616b696e6720746f6b656e20697320616c6c6f775f8201520152565b611da2602c604092611065565b611dab81611d48565b0190565b9190611dd2906020611dca604086018681035f880152611d95565b9401906104c1565b565b15611ddc5750565b611dfe90611de86101e2565b918291632290948760e21b835260048301611daf565b0390fd5b90611e0e60ff916112ee565b9181191691161790565b611e219061041e565b90565b90565b90611e3c611e37611e4392611e18565b611e24565b8254611e02565b9055565b90611e5860018060a01b03916112ee565b9181191691161790565b90565b90611e7a611e75611e8192611400565b611e62565b8254611e47565b9055565b611eac611ea5611e9f611e9a600585906103bf565b61118d565b1561041e565b8290611dd4565b611ec26001611ebd600584906103bf565b611e27565b611ed681611ed0600461069d565b9061328b565b611ef6815f611ef06003611eea600461069d565b90610aa3565b0161130c565b611f17336001611f116003611f0b600461069d565b90610aa3565b01611e65565b611f41611f24600c61069d565b6002611f3b6003611f35600461069d565b90610aa3565b0161130c565b611f5d611f56611f51600461069d565b6113e4565b600461130c565b611f79611f72611f6d600761069d565b6113e4565b600761130c565b611f8b611f866001611370565b6107b9565b6323b872dd33611f9a3061137d565b928492813b156120aa575f611fc291611fcd8296611fb66101e2565b9889978896879561138d565b8552600485016113a2565b03925af180156120a557612079575b50611ff9611ff2611fed600c61069d565b6113e4565b600c61130c565b612016612006600461069d565b61201060016116a6565b906116c2565b612020600761069d565b91339161207461206261205c6120567fde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d946103a3565b946103a3565b94611400565b9461206b6101e2565b918291826104ce565b0390a4565b612098905f3d811161209e575b6120908183610216565b810190611393565b5f611fdc565b503d612086565b6113d4565b611389565b6120c1906120bb611d2f565b50611cd6565b806120d46120ce5f611acd565b9161097e565b126120fc5760016120f36120f9926120ed600391611ae9565b90610aa3565b01610ad8565b90565b506121065f611c69565b90565b60207f207468652073616d65206c656e6774682e000000000000000000000000000000917f436c61696d20617272617920617267756d656e7473206d757374206265206f665f8201520152565b6121636031604092611065565b61216c81612109565b0190565b9160406121a192949361219a61218f606083018381035f850152612156565b9660208301906104c1565b01906104c1565b565b156121ac575050565b6121cd6121b76101e2565b92839263672c0fa160e11b845260048401612170565b0390fd5b92906121dc5f610f27565b5b806121f86121f26121ed88610f52565b610275565b91610275565b1015612227576122229061221d612218612213888490610f56565b610f76565b61140c565b610f43565b6121dd565b509190925061226b61223883610f52565b61225261224c61224785610f52565b610275565b91610275565b1461225c84610f52565b61226584610f52565b916121a3565b6122745f610f27565b9161229961228961228483610f52565b611a40565b6122936001611a5c565b90611a78565b925b836122ae6122a85f611acd565b9161097e565b1261230b576122ff612305916122f96122d86122d3866122cd8a611ae9565b90610f56565b610f76565b6122f36122ee886122e88b611ae9565b90610f56565b610f76565b90612b97565b906116e7565b93611aad565b9261229b565b925050508061232261231c5f610f27565b91610275565b1161232b575b50565b61235b905f80338361233b6101e2565b908161234681610ff1565b03925af161235261103c565b509033916110e2565b5f612328565b6123729061236d613107565b6124bc565b565b1561237b57565b5f632d6ebd6f60e01b8152806123936004820161034a565b0390fd5b5f7f5a65726f2d616464726573732077617320676976656e2e000000000000000000910152565b6123cb6017602092611065565b6123d481612397565b0190565b6123ed9060208101905f8183039101526123be565b90565b156123f757565b6123ff6101e2565b63eac0d38960e01b815280612416600482016123d8565b0390fd5b5f7f5472616e7366657220746f2063686172697479206661696c65642e0000000000910152565b61244e601b602092611065565b6124578161241a565b0190565b91604061248c92949361248561247a606083018381035f850152612441565b9660208301906104c1565b0190610a4c565b565b15612497575050565b6124b86124a26101e2565b92839263310a0fbb60e21b84526004840161245b565b0390fd5b6124e16124c9600761069d565b6124db6124d55f610f27565b91610275565b14612374565b612506816124ff6124f96124f45f611c69565b610a40565b91610a40565b14156123f0565b61250f3061137d565b3161253f5f80848461251f6101e2565b908161252a81610ff1565b03925af161253661103c565b5082849161248e565b9061257f61256d7f80c1082d1fcf8195bbf5a158fbef654d58f69408bd2e339b466bbd7c9fd7f74e92611400565b926125766101e2565b918291826104ce565b0390a2565b61258d90612361565b565b6125a66125ab9161259e611175565b50600861060a565b61069d565b6125bd6125b75f610f27565b91610275565b141590565b6125d3906125ce613107565b6125d5565b565b806125f06125ea6125e55f611c69565b610a40565b91610a40565b14612600576125fe90613155565b565b61262361260c5f611c69565b5f918291631e4fbdf760e01b835260048301610a59565b0390fd5b612630906125c2565b565b9061263c5f610f27565b5b8061265861265261264d86610f52565b610275565b91610275565b1015612687576126829061267d612678612673868490610f56565b610f76565b611e85565b610f43565b61263d565b509050565b5f7f496e76616c6964207374616b65416374696f6e49642e00000000000000000000910152565b6126c06016602092611065565b6126c98161268c565b0190565b91906126f09060206126e8604086018681035f8801526126b3565b9401906104c1565b565b156126fa5750565b61271c906127066101e2565b9182916313b7e41160e01b8352600483016126cd565b0390fd5b5f7f496e76616c6964204554484465706f73697449642e0000000000000000000000910152565b6127546015602092611065565b61275d81612720565b0190565b919061278490602061277c604086018681035f880152612747565b9401906104c1565b565b1561278e5750565b6127b09061279a6101e2565b91829163500e431160e11b835260048301612761565b0390fd5b5f7f546f6b656e20686173206e6f74206265656e20756e7374616b65642e00000000910152565b6127e8601c602092611065565b6127f1816127b4565b0190565b9190612818906020612810604086018681035f8801526127db565b9401906104c1565b565b156128225750565b6128449061282e6101e2565b91829163495097bd60e01b8352600483016127f5565b0390fd5b60207f2e00000000000000000000000000000000000000000000000000000000000000917f54686973206465706f7369742077617320636c61696d656420616c72656164795f8201520152565b6128a26021604092611065565b6128ab81612848565b0190565b9160406128e09294936128d96128ce606083018381035f850152612895565b9660208301906104c1565b01906104c1565b565b156128eb575050565b61290c6128f66101e2565b928392636aaada4b60e01b8452600484016128af565b0390fd5b5f7f4f6e6c7920746865206f776e65722063616e20636c61696d207265776172642e910152565b61294360208092611065565b61294c81612910565b0190565b91604061298192949361297a61296f606083018381035f850152612937565b9660208301906104c1565b0190610a4c565b565b1561298c575050565b6129ad6129976101e2565b9283926345c2e43b60e01b845260048401612950565b0390fd5b5f7f596f752077657265206e6f74207374616b6564207965742e0000000000000000910152565b6129e56018602092611065565b6129ee816129b1565b0190565b90959492612a4d94612a3c612a4692612a3260a096612a28612a1d60c089018981035f8b01526129d8565b9c60208901906104c1565b60408701906104c1565b60608501906104c1565b60808301906104c1565b01906104c1565b565b94929093919415612a61575050505050565b90612a859291612a6f6101e2565b95869563618d37eb60e01b8752600487016129f2565b0390fd5b5f7f596f75207765726520616c726561647920756e7374616b65642e000000000000910152565b612abd601a602092611065565b612ac681612a89565b0190565b90959492612b2594612b14612b1e92612b0a60a096612b00612af560c089018981035f8b0152612ab0565b9c60208901906104c1565b60408701906104c1565b60608501906104c1565b60808301906104c1565b01906104c1565b565b94929093919415612b39575050505050565b90612b5d9291612b476101e2565b95869563618d37eb60e01b875260048701612aca565b0390fd5b634e487b7160e01b5f52601260045260245ffd5b612b81612b8791610275565b91610275565b908115612b92570490565b612b61565b90612ba061195e565b50612bc882612bc0612bba612bb5600461069d565b610275565b91610275565b1083906126f2565b612bef81612be7612be1612bdc600b61069d565b610275565b91610275565b108290612786565b612c22612c086003612c02818690610aa3565b0161069d565b612c1a612c145f610f27565b91610275565b11839061281a565b612c57612c4f612c49612c446004612c3c60038890610aa3565b0185906103bf565b61118d565b1561041e565b8383916128e2565b612c8c612c716001612c6b60038690610aa3565b01610ad8565b612c83612c7d33610a40565b91610a40565b14833391612983565b612d1b612ca66002612ca060038690610aa3565b0161069d565b612ccd612cc7612cc25f612cbc600a889061066e565b0161069d565b610275565b91610275565b108383612ce66003612ce0818990610aa3565b0161069d565b90612cfe6002612cf860038a90610aa3565b0161069d565b92612d155f612d0f600a8a9061066e565b0161069d565b94612a4f565b612da9612d346003612d2e818690610aa3565b0161069d565b612d5b612d55612d505f612d4a600a889061066e565b0161069d565b610275565b91610275565b118383612d746003612d6e818990610aa3565b0161069d565b90612d8c6002612d8660038a90610aa3565b0161069d565b92612da35f612d9d600a8a9061066e565b0161069d565b94612b27565b612dcc6001612dc76004612dbf60038790610aa3565b0184906103bf565b611e27565b612e03612de66001612de0600a859061066e565b0161069d565b612dfd6002612df7600a869061066e565b0161069d565b90612b75565b9182913391612e59612e47612e41612e3b7fdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36946103a3565b946103a3565b94611400565b94612e506101e2565b918291826104ce565b0390a490565b5f7f546f6b656e206973206e6f7420696e20746865206c6973742e00000000000000910152565b612e936019602092611065565b612e9c81612e5f565b0190565b9190612ec3906020612ebb604086018681035f880152612e86565b9401906104c1565b565b15612ecd5750565b612eef90612ed96101e2565b918291639aa6fa6560e01b835260048301612ea0565b0390fd5b1b90565b91906008612f12910291612f0c5f1984612ef3565b92612ef3565b9181191691161790565b9190612f32612f2d612f3a936103a3565b611309565b908354612ef7565b9055565b612f5091612f4a61195e565b91612f1c565b565b90565b634e487b7160e01b5f52603160045260245ffd5b5490565b5f5260205f2090565b612f7f81612f69565b821015612f9957612f91600191612f6d565b910201905f90565b610c48565b612fa781612f69565b8015612fc8576001900390612fc5612fbf8383612f76565b90612f3e565b55565b612f55565b90565b612fe4612fdf612fe992612fcd565b6103a0565b61097e565b90565b613000612ffb6130059261097e565b6103a0565b61097e565b90565b90565b9061302061301b61302792612fec565b613008565b82546112f3565b9055565b6131059061304261303b8261258f565b8290612ec5565b6130c46130596130546008849061060a565b61069d565b6130bf61308d61308760066130816130716006610c5c565b61307b60016116a6565b906116c2565b90610c69565b906104a4565b916130b7836130b160066130ab856130a560016116a6565b906116c2565b90610c69565b90612f1c565b91600861060a565b61130c565b6130d95f6130d46008849061060a565b612f3e565b6130eb6130e66006612f52565b612f9e565b6131006130f85f19612fd0565b916009610926565b61300b565b565b61310f611d33565b61312861312261311d6132f8565b610a40565b91610a40565b0361312f57565b61315161313a6132f8565b5f91829163118cdaa760e01b835260048301610a59565b0390fd5b61315e5f610ad8565b613168825f611e65565b9061319c6131967f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093611400565b91611400565b916131a56101e2565b806131af8161034a565b0390a3565b5f7f546f6b656e20616c726561647920696e20746865206c6973742e000000000000910152565b6131e8601a602092611065565b6131f1816131b4565b0190565b91604061322692949361321f613214606083018381035f8501526131db565b9660208301906104c1565b01906104c1565b565b15613231575050565b61325261323c6101e2565b92839263597558c560e11b8452600484016131f5565b0390fd5b9081549168010000000000000000831015613286578261327e91600161328495018155612f76565b90612f1c565b565b610202565b6132f16132e96132f6936132b26132aa6132a48661258f565b1561041e565b858391613228565b6132c66132bf6006612f52565b8590613256565b6132e46132d36006610c5c565b6132df6008879061060a565b61130c565b611a40565b916009610926565b61300b565b565b613300611d2f565b50339056fea26469706673582212208c54b4251ff4f425d74dda990717cd3db29acf4511b2983efbe6f5e0ef7d830664736f6c634300081a0033",
}

// StakingWalletCSTABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingWalletCSTMetaData.ABI instead.
var StakingWalletCSTABI = StakingWalletCSTMetaData.ABI

// StakingWalletCSTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingWalletCSTMetaData.Bin instead.
var StakingWalletCSTBin = StakingWalletCSTMetaData.Bin

// DeployStakingWalletCST deploys a new Ethereum contract, binding an instance of StakingWalletCST to it.
func DeployStakingWalletCST(auth *bind.TransactOpts, backend bind.ContractBackend, nft_ common.Address, game_ common.Address) (common.Address, *types.Transaction, *StakingWalletCST, error) {
	parsed, err := StakingWalletCSTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingWalletCSTBin), backend, nft_, game_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingWalletCST{StakingWalletCSTCaller: StakingWalletCSTCaller{contract: contract}, StakingWalletCSTTransactor: StakingWalletCSTTransactor{contract: contract}, StakingWalletCSTFilterer: StakingWalletCSTFilterer{contract: contract}}, nil
}

// StakingWalletCST is an auto generated Go binding around an Ethereum contract.
type StakingWalletCST struct {
	StakingWalletCSTCaller     // Read-only binding to the contract
	StakingWalletCSTTransactor // Write-only binding to the contract
	StakingWalletCSTFilterer   // Log filterer for contract events
}

// StakingWalletCSTCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingWalletCSTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletCSTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingWalletCSTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletCSTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingWalletCSTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletCSTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingWalletCSTSession struct {
	Contract     *StakingWalletCST // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingWalletCSTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingWalletCSTCallerSession struct {
	Contract *StakingWalletCSTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StakingWalletCSTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingWalletCSTTransactorSession struct {
	Contract     *StakingWalletCSTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StakingWalletCSTRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingWalletCSTRaw struct {
	Contract *StakingWalletCST // Generic contract binding to access the raw methods on
}

// StakingWalletCSTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingWalletCSTCallerRaw struct {
	Contract *StakingWalletCSTCaller // Generic read-only contract binding to access the raw methods on
}

// StakingWalletCSTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingWalletCSTTransactorRaw struct {
	Contract *StakingWalletCSTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingWalletCST creates a new instance of StakingWalletCST, bound to a specific deployed contract.
func NewStakingWalletCST(address common.Address, backend bind.ContractBackend) (*StakingWalletCST, error) {
	contract, err := bindStakingWalletCST(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCST{StakingWalletCSTCaller: StakingWalletCSTCaller{contract: contract}, StakingWalletCSTTransactor: StakingWalletCSTTransactor{contract: contract}, StakingWalletCSTFilterer: StakingWalletCSTFilterer{contract: contract}}, nil
}

// NewStakingWalletCSTCaller creates a new read-only instance of StakingWalletCST, bound to a specific deployed contract.
func NewStakingWalletCSTCaller(address common.Address, caller bind.ContractCaller) (*StakingWalletCSTCaller, error) {
	contract, err := bindStakingWalletCST(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTCaller{contract: contract}, nil
}

// NewStakingWalletCSTTransactor creates a new write-only instance of StakingWalletCST, bound to a specific deployed contract.
func NewStakingWalletCSTTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingWalletCSTTransactor, error) {
	contract, err := bindStakingWalletCST(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTTransactor{contract: contract}, nil
}

// NewStakingWalletCSTFilterer creates a new log filterer instance of StakingWalletCST, bound to a specific deployed contract.
func NewStakingWalletCSTFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingWalletCSTFilterer, error) {
	contract, err := bindStakingWalletCST(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTFilterer{contract: contract}, nil
}

// bindStakingWalletCST binds a generic wrapper to an already deployed contract.
func bindStakingWalletCST(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingWalletCSTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletCST *StakingWalletCSTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletCST.Contract.StakingWalletCSTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletCST *StakingWalletCSTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.StakingWalletCSTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletCST *StakingWalletCSTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.StakingWalletCSTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletCST *StakingWalletCSTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletCST.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletCST *StakingWalletCSTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletCST *StakingWalletCSTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.contract.Transact(opts, method, params...)
}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint256 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWalletCST *StakingWalletCSTCaller) ETHDeposits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DepositTime   *big.Int
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "ETHDeposits", arg0)

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
func (_StakingWalletCST *StakingWalletCSTSession) ETHDeposits(arg0 *big.Int) (struct {
	DepositTime   *big.Int
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	return _StakingWalletCST.Contract.ETHDeposits(&_StakingWalletCST.CallOpts, arg0)
}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint256 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWalletCST *StakingWalletCSTCallerSession) ETHDeposits(arg0 *big.Int) (struct {
	DepositTime   *big.Int
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	return _StakingWalletCST.Contract.ETHDeposits(&_StakingWalletCST.CallOpts, arg0)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWalletCST *StakingWalletCSTCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWalletCST *StakingWalletCSTSession) Game() (common.Address, error) {
	return _StakingWalletCST.Contract.Game(&_StakingWalletCST.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWalletCST *StakingWalletCSTCallerSession) Game() (common.Address, error) {
	return _StakingWalletCST.Contract.Game(&_StakingWalletCST.CallOpts)
}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTCaller) IsTokenStaked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "isTokenStaked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTSession) IsTokenStaked(tokenId *big.Int) (bool, error) {
	return _StakingWalletCST.Contract.IsTokenStaked(&_StakingWalletCST.CallOpts, tokenId)
}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTCallerSession) IsTokenStaked(tokenId *big.Int) (bool, error) {
	return _StakingWalletCST.Contract.IsTokenStaked(&_StakingWalletCST.CallOpts, tokenId)
}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_StakingWalletCST *StakingWalletCSTCaller) LastActionIdByTokenId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "lastActionIdByTokenId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_StakingWalletCST *StakingWalletCSTSession) LastActionIdByTokenId(tokenId *big.Int) (*big.Int, error) {
	return _StakingWalletCST.Contract.LastActionIdByTokenId(&_StakingWalletCST.CallOpts, tokenId)
}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_StakingWalletCST *StakingWalletCSTCallerSession) LastActionIdByTokenId(tokenId *big.Int) (*big.Int, error) {
	return _StakingWalletCST.Contract.LastActionIdByTokenId(&_StakingWalletCST.CallOpts, tokenId)
}

// LastActionIds is a free data retrieval call binding the contract method 0x6427d9a9.
//
// Solidity: function lastActionIds(uint256 ) view returns(int256)
func (_StakingWalletCST *StakingWalletCSTCaller) LastActionIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "lastActionIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastActionIds is a free data retrieval call binding the contract method 0x6427d9a9.
//
// Solidity: function lastActionIds(uint256 ) view returns(int256)
func (_StakingWalletCST *StakingWalletCSTSession) LastActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletCST.Contract.LastActionIds(&_StakingWalletCST.CallOpts, arg0)
}

// LastActionIds is a free data retrieval call binding the contract method 0x6427d9a9.
//
// Solidity: function lastActionIds(uint256 ) view returns(int256)
func (_StakingWalletCST *StakingWalletCSTCallerSession) LastActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletCST.Contract.LastActionIds(&_StakingWalletCST.CallOpts, arg0)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWalletCST *StakingWalletCSTCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWalletCST *StakingWalletCSTSession) Nft() (common.Address, error) {
	return _StakingWalletCST.Contract.Nft(&_StakingWalletCST.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWalletCST *StakingWalletCSTCallerSession) Nft() (common.Address, error) {
	return _StakingWalletCST.Contract.Nft(&_StakingWalletCST.CallOpts)
}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCaller) NumETHDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "numETHDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTSession) NumETHDeposits() (*big.Int, error) {
	return _StakingWalletCST.Contract.NumETHDeposits(&_StakingWalletCST.CallOpts)
}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCallerSession) NumETHDeposits() (*big.Int, error) {
	return _StakingWalletCST.Contract.NumETHDeposits(&_StakingWalletCST.CallOpts)
}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCaller) NumStakeActions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "numStakeActions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTSession) NumStakeActions() (*big.Int, error) {
	return _StakingWalletCST.Contract.NumStakeActions(&_StakingWalletCST.CallOpts)
}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCallerSession) NumStakeActions() (*big.Int, error) {
	return _StakingWalletCST.Contract.NumStakeActions(&_StakingWalletCST.CallOpts)
}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCaller) NumStakedNFTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "numStakedNFTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTSession) NumStakedNFTs() (*big.Int, error) {
	return _StakingWalletCST.Contract.NumStakedNFTs(&_StakingWalletCST.CallOpts)
}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCallerSession) NumStakedNFTs() (*big.Int, error) {
	return _StakingWalletCST.Contract.NumStakedNFTs(&_StakingWalletCST.CallOpts)
}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCaller) NumTokensStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "numTokensStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTSession) NumTokensStaked() (*big.Int, error) {
	return _StakingWalletCST.Contract.NumTokensStaked(&_StakingWalletCST.CallOpts)
}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCallerSession) NumTokensStaked() (*big.Int, error) {
	return _StakingWalletCST.Contract.NumTokensStaked(&_StakingWalletCST.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletCST *StakingWalletCSTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletCST *StakingWalletCSTSession) Owner() (common.Address, error) {
	return _StakingWalletCST.Contract.Owner(&_StakingWalletCST.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletCST *StakingWalletCSTCallerSession) Owner() (common.Address, error) {
	return _StakingWalletCST.Contract.Owner(&_StakingWalletCST.CallOpts)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address nftOwner, uint256 stakeTime, uint256 unstakeTime)
func (_StakingWalletCST *StakingWalletCSTCaller) StakeActions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId     *big.Int
	NftOwner    common.Address
	StakeTime   *big.Int
	UnstakeTime *big.Int
}, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "stakeActions", arg0)

	outstruct := new(struct {
		TokenId     *big.Int
		NftOwner    common.Address
		StakeTime   *big.Int
		UnstakeTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftOwner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StakeTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address nftOwner, uint256 stakeTime, uint256 unstakeTime)
func (_StakingWalletCST *StakingWalletCSTSession) StakeActions(arg0 *big.Int) (struct {
	TokenId     *big.Int
	NftOwner    common.Address
	StakeTime   *big.Int
	UnstakeTime *big.Int
}, error) {
	return _StakingWalletCST.Contract.StakeActions(&_StakingWalletCST.CallOpts, arg0)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address nftOwner, uint256 stakeTime, uint256 unstakeTime)
func (_StakingWalletCST *StakingWalletCSTCallerSession) StakeActions(arg0 *big.Int) (struct {
	TokenId     *big.Int
	NftOwner    common.Address
	StakeTime   *big.Int
	UnstakeTime *big.Int
}, error) {
	return _StakingWalletCST.Contract.StakeActions(&_StakingWalletCST.CallOpts, arg0)
}

// StakedTokens is a free data retrieval call binding the contract method 0xc0788555.
//
// Solidity: function stakedTokens(uint256 ) view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCaller) StakedTokens(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "stakedTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedTokens is a free data retrieval call binding the contract method 0xc0788555.
//
// Solidity: function stakedTokens(uint256 ) view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTSession) StakedTokens(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletCST.Contract.StakedTokens(&_StakingWalletCST.CallOpts, arg0)
}

// StakedTokens is a free data retrieval call binding the contract method 0xc0788555.
//
// Solidity: function stakedTokens(uint256 ) view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCallerSession) StakedTokens(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletCST.Contract.StakedTokens(&_StakingWalletCST.CallOpts, arg0)
}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_StakingWalletCST *StakingWalletCSTCaller) StakerByTokenId(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "stakerByTokenId", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_StakingWalletCST *StakingWalletCSTSession) StakerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _StakingWalletCST.Contract.StakerByTokenId(&_StakingWalletCST.CallOpts, tokenId)
}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_StakingWalletCST *StakingWalletCSTCallerSession) StakerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _StakingWalletCST.Contract.StakerByTokenId(&_StakingWalletCST.CallOpts, tokenId)
}

// TokenIndices is a free data retrieval call binding the contract method 0x44d110b9.
//
// Solidity: function tokenIndices(uint256 ) view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCaller) TokenIndices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "tokenIndices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIndices is a free data retrieval call binding the contract method 0x44d110b9.
//
// Solidity: function tokenIndices(uint256 ) view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTSession) TokenIndices(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletCST.Contract.TokenIndices(&_StakingWalletCST.CallOpts, arg0)
}

// TokenIndices is a free data retrieval call binding the contract method 0x44d110b9.
//
// Solidity: function tokenIndices(uint256 ) view returns(uint256)
func (_StakingWalletCST *StakingWalletCSTCallerSession) TokenIndices(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletCST.Contract.TokenIndices(&_StakingWalletCST.CallOpts, arg0)
}

// UsedTokens is a free data retrieval call binding the contract method 0x0f7ee879.
//
// Solidity: function usedTokens(uint256 ) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTCaller) UsedTokens(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "usedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedTokens is a free data retrieval call binding the contract method 0x0f7ee879.
//
// Solidity: function usedTokens(uint256 ) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTSession) UsedTokens(arg0 *big.Int) (bool, error) {
	return _StakingWalletCST.Contract.UsedTokens(&_StakingWalletCST.CallOpts, arg0)
}

// UsedTokens is a free data retrieval call binding the contract method 0x0f7ee879.
//
// Solidity: function usedTokens(uint256 ) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTCallerSession) UsedTokens(arg0 *big.Int) (bool, error) {
	return _StakingWalletCST.Contract.UsedTokens(&_StakingWalletCST.CallOpts, arg0)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTCaller) WasTokenUsed(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "wasTokenUsed", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTSession) WasTokenUsed(_tokenId *big.Int) (bool, error) {
	return _StakingWalletCST.Contract.WasTokenUsed(&_StakingWalletCST.CallOpts, _tokenId)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTCallerSession) WasTokenUsed(_tokenId *big.Int) (bool, error) {
	return _StakingWalletCST.Contract.WasTokenUsed(&_StakingWalletCST.CallOpts, _tokenId)
}

// ClaimManyRewards is a paid mutator transaction binding the contract method 0x6034eb5b.
//
// Solidity: function claimManyRewards(uint256[] actions, uint256[] deposits) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) ClaimManyRewards(opts *bind.TransactOpts, actions []*big.Int, deposits []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "claimManyRewards", actions, deposits)
}

// ClaimManyRewards is a paid mutator transaction binding the contract method 0x6034eb5b.
//
// Solidity: function claimManyRewards(uint256[] actions, uint256[] deposits) returns()
func (_StakingWalletCST *StakingWalletCSTSession) ClaimManyRewards(actions []*big.Int, deposits []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.ClaimManyRewards(&_StakingWalletCST.TransactOpts, actions, deposits)
}

// ClaimManyRewards is a paid mutator transaction binding the contract method 0x6034eb5b.
//
// Solidity: function claimManyRewards(uint256[] actions, uint256[] deposits) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) ClaimManyRewards(actions []*big.Int, deposits []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.ClaimManyRewards(&_StakingWalletCST.TransactOpts, actions, deposits)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x326c7b69.
//
// Solidity: function depositIfPossible() payable returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) DepositIfPossible(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "depositIfPossible")
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x326c7b69.
//
// Solidity: function depositIfPossible() payable returns()
func (_StakingWalletCST *StakingWalletCSTSession) DepositIfPossible() (*types.Transaction, error) {
	return _StakingWalletCST.Contract.DepositIfPossible(&_StakingWalletCST.TransactOpts)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x326c7b69.
//
// Solidity: function depositIfPossible() payable returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) DepositIfPossible() (*types.Transaction, error) {
	return _StakingWalletCST.Contract.DepositIfPossible(&_StakingWalletCST.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletCST *StakingWalletCSTSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWalletCST.Contract.RenounceOwnership(&_StakingWalletCST.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWalletCST.Contract.RenounceOwnership(&_StakingWalletCST.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) Stake(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "stake", _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWalletCST *StakingWalletCSTSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.Stake(&_StakingWalletCST.TransactOpts, _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.Stake(&_StakingWalletCST.TransactOpts, _tokenId)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) StakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "stakeMany", ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWalletCST *StakingWalletCSTSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.StakeMany(&_StakingWalletCST.TransactOpts, ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.StakeMany(&_StakingWalletCST.TransactOpts, ids)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletCST *StakingWalletCSTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.TransferOwnership(&_StakingWalletCST.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.TransferOwnership(&_StakingWalletCST.TransactOpts, newOwner)
}

// TransferRemainingBalanceToCharity is a paid mutator transaction binding the contract method 0xed614ffa.
//
// Solidity: function transferRemainingBalanceToCharity(address charityAddress_) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) TransferRemainingBalanceToCharity(opts *bind.TransactOpts, charityAddress_ common.Address) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "transferRemainingBalanceToCharity", charityAddress_)
}

// TransferRemainingBalanceToCharity is a paid mutator transaction binding the contract method 0xed614ffa.
//
// Solidity: function transferRemainingBalanceToCharity(address charityAddress_) returns()
func (_StakingWalletCST *StakingWalletCSTSession) TransferRemainingBalanceToCharity(charityAddress_ common.Address) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.TransferRemainingBalanceToCharity(&_StakingWalletCST.TransactOpts, charityAddress_)
}

// TransferRemainingBalanceToCharity is a paid mutator transaction binding the contract method 0xed614ffa.
//
// Solidity: function transferRemainingBalanceToCharity(address charityAddress_) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) TransferRemainingBalanceToCharity(charityAddress_ common.Address) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.TransferRemainingBalanceToCharity(&_StakingWalletCST.TransactOpts, charityAddress_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) Unstake(opts *bind.TransactOpts, stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "unstake", stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWalletCST *StakingWalletCSTSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.Unstake(&_StakingWalletCST.TransactOpts, stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.Unstake(&_StakingWalletCST.TransactOpts, stakeActionId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0x1969e066.
//
// Solidity: function unstakeClaim(uint256 stakeActionId, uint256 ETHDepositId) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) UnstakeClaim(opts *bind.TransactOpts, stakeActionId *big.Int, ETHDepositId *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "unstakeClaim", stakeActionId, ETHDepositId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0x1969e066.
//
// Solidity: function unstakeClaim(uint256 stakeActionId, uint256 ETHDepositId) returns()
func (_StakingWalletCST *StakingWalletCSTSession) UnstakeClaim(stakeActionId *big.Int, ETHDepositId *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.UnstakeClaim(&_StakingWalletCST.TransactOpts, stakeActionId, ETHDepositId)
}

// UnstakeClaim is a paid mutator transaction binding the contract method 0x1969e066.
//
// Solidity: function unstakeClaim(uint256 stakeActionId, uint256 ETHDepositId) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) UnstakeClaim(stakeActionId *big.Int, ETHDepositId *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.UnstakeClaim(&_StakingWalletCST.TransactOpts, stakeActionId, ETHDepositId)
}

// UnstakeClaimMany is a paid mutator transaction binding the contract method 0xdb2b4bd8.
//
// Solidity: function unstakeClaimMany(uint256[] unstake_actions, uint256[] claim_actions, uint256[] claim_deposits) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) UnstakeClaimMany(opts *bind.TransactOpts, unstake_actions []*big.Int, claim_actions []*big.Int, claim_deposits []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "unstakeClaimMany", unstake_actions, claim_actions, claim_deposits)
}

// UnstakeClaimMany is a paid mutator transaction binding the contract method 0xdb2b4bd8.
//
// Solidity: function unstakeClaimMany(uint256[] unstake_actions, uint256[] claim_actions, uint256[] claim_deposits) returns()
func (_StakingWalletCST *StakingWalletCSTSession) UnstakeClaimMany(unstake_actions []*big.Int, claim_actions []*big.Int, claim_deposits []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.UnstakeClaimMany(&_StakingWalletCST.TransactOpts, unstake_actions, claim_actions, claim_deposits)
}

// UnstakeClaimMany is a paid mutator transaction binding the contract method 0xdb2b4bd8.
//
// Solidity: function unstakeClaimMany(uint256[] unstake_actions, uint256[] claim_actions, uint256[] claim_deposits) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) UnstakeClaimMany(unstake_actions []*big.Int, claim_actions []*big.Int, claim_deposits []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.UnstakeClaimMany(&_StakingWalletCST.TransactOpts, unstake_actions, claim_actions, claim_deposits)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) UnstakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "unstakeMany", ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWalletCST *StakingWalletCSTSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.UnstakeMany(&_StakingWalletCST.TransactOpts, ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.UnstakeMany(&_StakingWalletCST.TransactOpts, ids)
}

// StakingWalletCSTClaimRewardEventIterator is returned from FilterClaimRewardEvent and is used to iterate over the raw logs and unpacked data for ClaimRewardEvent events raised by the StakingWalletCST contract.
type StakingWalletCSTClaimRewardEventIterator struct {
	Event *StakingWalletCSTClaimRewardEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletCSTClaimRewardEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCSTClaimRewardEvent)
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
		it.Event = new(StakingWalletCSTClaimRewardEvent)
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
func (it *StakingWalletCSTClaimRewardEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCSTClaimRewardEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCSTClaimRewardEvent represents a ClaimRewardEvent event raised by the StakingWalletCST contract.
type StakingWalletCSTClaimRewardEvent struct {
	ActionId  *big.Int
	DepositId *big.Int
	Reward    *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimRewardEvent is a free log retrieval operation binding the contract event 0xdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36.
//
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address indexed staker)
func (_StakingWalletCST *StakingWalletCSTFilterer) FilterClaimRewardEvent(opts *bind.FilterOpts, actionId []*big.Int, depositId []*big.Int, staker []common.Address) (*StakingWalletCSTClaimRewardEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWalletCST.contract.FilterLogs(opts, "ClaimRewardEvent", actionIdRule, depositIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTClaimRewardEventIterator{contract: _StakingWalletCST.contract, event: "ClaimRewardEvent", logs: logs, sub: sub}, nil
}

// WatchClaimRewardEvent is a free log subscription operation binding the contract event 0xdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36.
//
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address indexed staker)
func (_StakingWalletCST *StakingWalletCSTFilterer) WatchClaimRewardEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletCSTClaimRewardEvent, actionId []*big.Int, depositId []*big.Int, staker []common.Address) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWalletCST.contract.WatchLogs(opts, "ClaimRewardEvent", actionIdRule, depositIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCSTClaimRewardEvent)
				if err := _StakingWalletCST.contract.UnpackLog(event, "ClaimRewardEvent", log); err != nil {
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

// ParseClaimRewardEvent is a log parse operation binding the contract event 0xdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36.
//
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address indexed staker)
func (_StakingWalletCST *StakingWalletCSTFilterer) ParseClaimRewardEvent(log types.Log) (*StakingWalletCSTClaimRewardEvent, error) {
	event := new(StakingWalletCSTClaimRewardEvent)
	if err := _StakingWalletCST.contract.UnpackLog(event, "ClaimRewardEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCSTEthDepositEventIterator is returned from FilterEthDepositEvent and is used to iterate over the raw logs and unpacked data for EthDepositEvent events raised by the StakingWalletCST contract.
type StakingWalletCSTEthDepositEventIterator struct {
	Event *StakingWalletCSTEthDepositEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletCSTEthDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCSTEthDepositEvent)
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
		it.Event = new(StakingWalletCSTEthDepositEvent)
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
func (it *StakingWalletCSTEthDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCSTEthDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCSTEthDepositEvent represents a EthDepositEvent event raised by the StakingWalletCST contract.
type StakingWalletCSTEthDepositEvent struct {
	DepositTime   *big.Int
	DepositNum    *big.Int
	NumStakedNFTs *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthDepositEvent is a free log retrieval operation binding the contract event 0xd91db13dbdc68f51d0e522c3f29f3d5c0f5945fa377792cb335e42f78b82d1d8.
//
// Solidity: event EthDepositEvent(uint256 indexed depositTime, uint256 depositNum, uint256 numStakedNFTs, uint256 amount)
func (_StakingWalletCST *StakingWalletCSTFilterer) FilterEthDepositEvent(opts *bind.FilterOpts, depositTime []*big.Int) (*StakingWalletCSTEthDepositEventIterator, error) {

	var depositTimeRule []interface{}
	for _, depositTimeItem := range depositTime {
		depositTimeRule = append(depositTimeRule, depositTimeItem)
	}

	logs, sub, err := _StakingWalletCST.contract.FilterLogs(opts, "EthDepositEvent", depositTimeRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTEthDepositEventIterator{contract: _StakingWalletCST.contract, event: "EthDepositEvent", logs: logs, sub: sub}, nil
}

// WatchEthDepositEvent is a free log subscription operation binding the contract event 0xd91db13dbdc68f51d0e522c3f29f3d5c0f5945fa377792cb335e42f78b82d1d8.
//
// Solidity: event EthDepositEvent(uint256 indexed depositTime, uint256 depositNum, uint256 numStakedNFTs, uint256 amount)
func (_StakingWalletCST *StakingWalletCSTFilterer) WatchEthDepositEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletCSTEthDepositEvent, depositTime []*big.Int) (event.Subscription, error) {

	var depositTimeRule []interface{}
	for _, depositTimeItem := range depositTime {
		depositTimeRule = append(depositTimeRule, depositTimeItem)
	}

	logs, sub, err := _StakingWalletCST.contract.WatchLogs(opts, "EthDepositEvent", depositTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCSTEthDepositEvent)
				if err := _StakingWalletCST.contract.UnpackLog(event, "EthDepositEvent", log); err != nil {
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

// ParseEthDepositEvent is a log parse operation binding the contract event 0xd91db13dbdc68f51d0e522c3f29f3d5c0f5945fa377792cb335e42f78b82d1d8.
//
// Solidity: event EthDepositEvent(uint256 indexed depositTime, uint256 depositNum, uint256 numStakedNFTs, uint256 amount)
func (_StakingWalletCST *StakingWalletCSTFilterer) ParseEthDepositEvent(log types.Log) (*StakingWalletCSTEthDepositEvent, error) {
	event := new(StakingWalletCSTEthDepositEvent)
	if err := _StakingWalletCST.contract.UnpackLog(event, "EthDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCSTFundsTransferredToCharityEventIterator is returned from FilterFundsTransferredToCharityEvent and is used to iterate over the raw logs and unpacked data for FundsTransferredToCharityEvent events raised by the StakingWalletCST contract.
type StakingWalletCSTFundsTransferredToCharityEventIterator struct {
	Event *StakingWalletCSTFundsTransferredToCharityEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletCSTFundsTransferredToCharityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCSTFundsTransferredToCharityEvent)
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
		it.Event = new(StakingWalletCSTFundsTransferredToCharityEvent)
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
func (it *StakingWalletCSTFundsTransferredToCharityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCSTFundsTransferredToCharityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCSTFundsTransferredToCharityEvent represents a FundsTransferredToCharityEvent event raised by the StakingWalletCST contract.
type StakingWalletCSTFundsTransferredToCharityEvent struct {
	Amount         *big.Int
	CharityAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFundsTransferredToCharityEvent is a free log retrieval operation binding the contract event 0x80c1082d1fcf8195bbf5a158fbef654d58f69408bd2e339b466bbd7c9fd7f74e.
//
// Solidity: event FundsTransferredToCharityEvent(uint256 amount, address indexed charityAddress)
func (_StakingWalletCST *StakingWalletCSTFilterer) FilterFundsTransferredToCharityEvent(opts *bind.FilterOpts, charityAddress []common.Address) (*StakingWalletCSTFundsTransferredToCharityEventIterator, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _StakingWalletCST.contract.FilterLogs(opts, "FundsTransferredToCharityEvent", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTFundsTransferredToCharityEventIterator{contract: _StakingWalletCST.contract, event: "FundsTransferredToCharityEvent", logs: logs, sub: sub}, nil
}

// WatchFundsTransferredToCharityEvent is a free log subscription operation binding the contract event 0x80c1082d1fcf8195bbf5a158fbef654d58f69408bd2e339b466bbd7c9fd7f74e.
//
// Solidity: event FundsTransferredToCharityEvent(uint256 amount, address indexed charityAddress)
func (_StakingWalletCST *StakingWalletCSTFilterer) WatchFundsTransferredToCharityEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletCSTFundsTransferredToCharityEvent, charityAddress []common.Address) (event.Subscription, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _StakingWalletCST.contract.WatchLogs(opts, "FundsTransferredToCharityEvent", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCSTFundsTransferredToCharityEvent)
				if err := _StakingWalletCST.contract.UnpackLog(event, "FundsTransferredToCharityEvent", log); err != nil {
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

// ParseFundsTransferredToCharityEvent is a log parse operation binding the contract event 0x80c1082d1fcf8195bbf5a158fbef654d58f69408bd2e339b466bbd7c9fd7f74e.
//
// Solidity: event FundsTransferredToCharityEvent(uint256 amount, address indexed charityAddress)
func (_StakingWalletCST *StakingWalletCSTFilterer) ParseFundsTransferredToCharityEvent(log types.Log) (*StakingWalletCSTFundsTransferredToCharityEvent, error) {
	event := new(StakingWalletCSTFundsTransferredToCharityEvent)
	if err := _StakingWalletCST.contract.UnpackLog(event, "FundsTransferredToCharityEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCSTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingWalletCST contract.
type StakingWalletCSTOwnershipTransferredIterator struct {
	Event *StakingWalletCSTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingWalletCSTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCSTOwnershipTransferred)
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
		it.Event = new(StakingWalletCSTOwnershipTransferred)
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
func (it *StakingWalletCSTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCSTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCSTOwnershipTransferred represents a OwnershipTransferred event raised by the StakingWalletCST contract.
type StakingWalletCSTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWalletCST *StakingWalletCSTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingWalletCSTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWalletCST.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTOwnershipTransferredIterator{contract: _StakingWalletCST.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWalletCST *StakingWalletCSTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingWalletCSTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWalletCST.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCSTOwnershipTransferred)
				if err := _StakingWalletCST.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingWalletCST *StakingWalletCSTFilterer) ParseOwnershipTransferred(log types.Log) (*StakingWalletCSTOwnershipTransferred, error) {
	event := new(StakingWalletCSTOwnershipTransferred)
	if err := _StakingWalletCST.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCSTStakeActionEventIterator is returned from FilterStakeActionEvent and is used to iterate over the raw logs and unpacked data for StakeActionEvent events raised by the StakingWalletCST contract.
type StakingWalletCSTStakeActionEventIterator struct {
	Event *StakingWalletCSTStakeActionEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletCSTStakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCSTStakeActionEvent)
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
		it.Event = new(StakingWalletCSTStakeActionEvent)
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
func (it *StakingWalletCSTStakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCSTStakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCSTStakeActionEvent represents a StakeActionEvent event raised by the StakingWalletCST contract.
type StakingWalletCSTStakeActionEvent struct {
	ActionId  *big.Int
	TokenId   *big.Int
	TotalNFTs *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeActionEvent is a free log retrieval operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletCST *StakingWalletCSTFilterer) FilterStakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*StakingWalletCSTStakeActionEventIterator, error) {

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

	logs, sub, err := _StakingWalletCST.contract.FilterLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTStakeActionEventIterator{contract: _StakingWalletCST.contract, event: "StakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchStakeActionEvent is a free log subscription operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletCST *StakingWalletCSTFilterer) WatchStakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletCSTStakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingWalletCST.contract.WatchLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCSTStakeActionEvent)
				if err := _StakingWalletCST.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
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
func (_StakingWalletCST *StakingWalletCSTFilterer) ParseStakeActionEvent(log types.Log) (*StakingWalletCSTStakeActionEvent, error) {
	event := new(StakingWalletCSTStakeActionEvent)
	if err := _StakingWalletCST.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCSTUnstakeActionEventIterator is returned from FilterUnstakeActionEvent and is used to iterate over the raw logs and unpacked data for UnstakeActionEvent events raised by the StakingWalletCST contract.
type StakingWalletCSTUnstakeActionEventIterator struct {
	Event *StakingWalletCSTUnstakeActionEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletCSTUnstakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCSTUnstakeActionEvent)
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
		it.Event = new(StakingWalletCSTUnstakeActionEvent)
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
func (it *StakingWalletCSTUnstakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCSTUnstakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCSTUnstakeActionEvent represents a UnstakeActionEvent event raised by the StakingWalletCST contract.
type StakingWalletCSTUnstakeActionEvent struct {
	ActionId  *big.Int
	TokenId   *big.Int
	TotalNFTs *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstakeActionEvent is a free log retrieval operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletCST *StakingWalletCSTFilterer) FilterUnstakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*StakingWalletCSTUnstakeActionEventIterator, error) {

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

	logs, sub, err := _StakingWalletCST.contract.FilterLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTUnstakeActionEventIterator{contract: _StakingWalletCST.contract, event: "UnstakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchUnstakeActionEvent is a free log subscription operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletCST *StakingWalletCSTFilterer) WatchUnstakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletCSTUnstakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingWalletCST.contract.WatchLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCSTUnstakeActionEvent)
				if err := _StakingWalletCST.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
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
func (_StakingWalletCST *StakingWalletCSTFilterer) ParseUnstakeActionEvent(log types.Log) (*StakingWalletCSTUnstakeActionEvent, error) {
	event := new(StakingWalletCSTUnstakeActionEvent)
	if err := _StakingWalletCST.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
