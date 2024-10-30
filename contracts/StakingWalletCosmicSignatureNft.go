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

// IStakingWalletCosmicSignatureNftMetaData contains all meta data concerning the IStakingWalletCosmicSignatureNft contract.
var IStakingWalletCosmicSignatureNftMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"EthDepositReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicGameConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetState_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"tryPerformMaintenance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IStakingWalletCosmicSignatureNftABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingWalletCosmicSignatureNftMetaData.ABI instead.
var IStakingWalletCosmicSignatureNftABI = IStakingWalletCosmicSignatureNftMetaData.ABI

// IStakingWalletCosmicSignatureNft is an auto generated Go binding around an Ethereum contract.
type IStakingWalletCosmicSignatureNft struct {
	IStakingWalletCosmicSignatureNftCaller     // Read-only binding to the contract
	IStakingWalletCosmicSignatureNftTransactor // Write-only binding to the contract
	IStakingWalletCosmicSignatureNftFilterer   // Log filterer for contract events
}

// IStakingWalletCosmicSignatureNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingWalletCosmicSignatureNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletCosmicSignatureNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingWalletCosmicSignatureNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletCosmicSignatureNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingWalletCosmicSignatureNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletCosmicSignatureNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingWalletCosmicSignatureNftSession struct {
	Contract     *IStakingWalletCosmicSignatureNft // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IStakingWalletCosmicSignatureNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingWalletCosmicSignatureNftCallerSession struct {
	Contract *IStakingWalletCosmicSignatureNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// IStakingWalletCosmicSignatureNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingWalletCosmicSignatureNftTransactorSession struct {
	Contract     *IStakingWalletCosmicSignatureNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// IStakingWalletCosmicSignatureNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingWalletCosmicSignatureNftRaw struct {
	Contract *IStakingWalletCosmicSignatureNft // Generic contract binding to access the raw methods on
}

// IStakingWalletCosmicSignatureNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingWalletCosmicSignatureNftCallerRaw struct {
	Contract *IStakingWalletCosmicSignatureNftCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingWalletCosmicSignatureNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingWalletCosmicSignatureNftTransactorRaw struct {
	Contract *IStakingWalletCosmicSignatureNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakingWalletCosmicSignatureNft creates a new instance of IStakingWalletCosmicSignatureNft, bound to a specific deployed contract.
func NewIStakingWalletCosmicSignatureNft(address common.Address, backend bind.ContractBackend) (*IStakingWalletCosmicSignatureNft, error) {
	contract, err := bindIStakingWalletCosmicSignatureNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCosmicSignatureNft{IStakingWalletCosmicSignatureNftCaller: IStakingWalletCosmicSignatureNftCaller{contract: contract}, IStakingWalletCosmicSignatureNftTransactor: IStakingWalletCosmicSignatureNftTransactor{contract: contract}, IStakingWalletCosmicSignatureNftFilterer: IStakingWalletCosmicSignatureNftFilterer{contract: contract}}, nil
}

// NewIStakingWalletCosmicSignatureNftCaller creates a new read-only instance of IStakingWalletCosmicSignatureNft, bound to a specific deployed contract.
func NewIStakingWalletCosmicSignatureNftCaller(address common.Address, caller bind.ContractCaller) (*IStakingWalletCosmicSignatureNftCaller, error) {
	contract, err := bindIStakingWalletCosmicSignatureNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCosmicSignatureNftCaller{contract: contract}, nil
}

// NewIStakingWalletCosmicSignatureNftTransactor creates a new write-only instance of IStakingWalletCosmicSignatureNft, bound to a specific deployed contract.
func NewIStakingWalletCosmicSignatureNftTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingWalletCosmicSignatureNftTransactor, error) {
	contract, err := bindIStakingWalletCosmicSignatureNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCosmicSignatureNftTransactor{contract: contract}, nil
}

// NewIStakingWalletCosmicSignatureNftFilterer creates a new log filterer instance of IStakingWalletCosmicSignatureNft, bound to a specific deployed contract.
func NewIStakingWalletCosmicSignatureNftFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingWalletCosmicSignatureNftFilterer, error) {
	contract, err := bindIStakingWalletCosmicSignatureNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCosmicSignatureNftFilterer{contract: contract}, nil
}

// bindIStakingWalletCosmicSignatureNft binds a generic wrapper to an already deployed contract.
func bindIStakingWalletCosmicSignatureNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakingWalletCosmicSignatureNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletCosmicSignatureNft.Contract.IStakingWalletCosmicSignatureNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.IStakingWalletCosmicSignatureNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.IStakingWalletCosmicSignatureNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletCosmicSignatureNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.contract.Transact(opts, method, params...)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftCaller) NumStakedNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingWalletCosmicSignatureNft.contract.Call(opts, &out, "numStakedNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftSession) NumStakedNfts() (*big.Int, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.NumStakedNfts(&_IStakingWalletCosmicSignatureNft.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftCallerSession) NumStakedNfts() (*big.Int, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.NumStakedNfts(&_IStakingWalletCosmicSignatureNft.CallOpts)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _IStakingWalletCosmicSignatureNft.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.WasNftUsed(&_IStakingWalletCosmicSignatureNft.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftCallerSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.WasNftUsed(&_IStakingWalletCosmicSignatureNft.CallOpts, nftId_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactor) DepositIfPossible(opts *bind.TransactOpts, roundNum_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.contract.Transact(opts, "depositIfPossible", roundNum_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftSession) DepositIfPossible(roundNum_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.DepositIfPossible(&_IStakingWalletCosmicSignatureNft.TransactOpts, roundNum_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactorSession) DepositIfPossible(roundNum_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.DepositIfPossible(&_IStakingWalletCosmicSignatureNft.TransactOpts, roundNum_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0x2e53a259.
//
// Solidity: function payManyRewards(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactor) PayManyRewards(opts *bind.TransactOpts, stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.contract.Transact(opts, "payManyRewards", stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0x2e53a259.
//
// Solidity: function payManyRewards(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftSession) PayManyRewards(stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.PayManyRewards(&_IStakingWalletCosmicSignatureNft.TransactOpts, stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0x2e53a259.
//
// Solidity: function payManyRewards(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactorSession) PayManyRewards(stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.PayManyRewards(&_IStakingWalletCosmicSignatureNft.TransactOpts, stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// PayReward is a paid mutator transaction binding the contract method 0xc2ac898f.
//
// Solidity: function payReward(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactor) PayReward(opts *bind.TransactOpts, stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.contract.Transact(opts, "payReward", stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// PayReward is a paid mutator transaction binding the contract method 0xc2ac898f.
//
// Solidity: function payReward(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftSession) PayReward(stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.PayReward(&_IStakingWalletCosmicSignatureNft.TransactOpts, stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// PayReward is a paid mutator transaction binding the contract method 0xc2ac898f.
//
// Solidity: function payReward(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactorSession) PayReward(stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.PayReward(&_IStakingWalletCosmicSignatureNft.TransactOpts, stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactor) Stake(opts *bind.TransactOpts, nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.contract.Transact(opts, "stake", nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.Stake(&_IStakingWalletCosmicSignatureNft.TransactOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactorSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.Stake(&_IStakingWalletCosmicSignatureNft.TransactOpts, nftId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactor) StakeMany(opts *bind.TransactOpts, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.contract.Transact(opts, "stakeMany", nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.StakeMany(&_IStakingWalletCosmicSignatureNft.TransactOpts, nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactorSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.StakeMany(&_IStakingWalletCosmicSignatureNft.TransactOpts, nftIds_)
}

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactor) TryPerformMaintenance(opts *bind.TransactOpts, resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.contract.Transact(opts, "tryPerformMaintenance", resetState_, charityAddress_)
}

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftSession) TryPerformMaintenance(resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.TryPerformMaintenance(&_IStakingWalletCosmicSignatureNft.TransactOpts, resetState_, charityAddress_)
}

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactorSession) TryPerformMaintenance(resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.TryPerformMaintenance(&_IStakingWalletCosmicSignatureNft.TransactOpts, resetState_, charityAddress_)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactor) Unstake(opts *bind.TransactOpts, stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.contract.Transact(opts, "unstake", stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftSession) Unstake(stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.Unstake(&_IStakingWalletCosmicSignatureNft.TransactOpts, stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactorSession) Unstake(stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.Unstake(&_IStakingWalletCosmicSignatureNft.TransactOpts, stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x400c8956.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactor) UnstakeMany(opts *bind.TransactOpts, stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.contract.Transact(opts, "unstakeMany", stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x400c8956.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftSession) UnstakeMany(stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.UnstakeMany(&_IStakingWalletCosmicSignatureNft.TransactOpts, stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x400c8956.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftTransactorSession) UnstakeMany(stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.UnstakeMany(&_IStakingWalletCosmicSignatureNft.TransactOpts, stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// IStakingWalletCosmicSignatureNftEthDepositReceivedIterator is returned from FilterEthDepositReceived and is used to iterate over the raw logs and unpacked data for EthDepositReceived events raised by the IStakingWalletCosmicSignatureNft contract.
type IStakingWalletCosmicSignatureNftEthDepositReceivedIterator struct {
	Event *IStakingWalletCosmicSignatureNftEthDepositReceived // Event containing the contract specifics and raw log

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
func (it *IStakingWalletCosmicSignatureNftEthDepositReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletCosmicSignatureNftEthDepositReceived)
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
		it.Event = new(IStakingWalletCosmicSignatureNftEthDepositReceived)
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
func (it *IStakingWalletCosmicSignatureNftEthDepositReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletCosmicSignatureNftEthDepositReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletCosmicSignatureNftEthDepositReceived represents a EthDepositReceived event raised by the IStakingWalletCosmicSignatureNft contract.
type IStakingWalletCosmicSignatureNftEthDepositReceived struct {
	RoundNum      *big.Int
	ActionCounter *big.Int
	DepositIndex  *big.Int
	DepositId     *big.Int
	DepositAmount *big.Int
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthDepositReceived is a free log retrieval operation binding the contract event 0xb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c0913.
//
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNfts)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) FilterEthDepositReceived(opts *bind.FilterOpts, roundNum []*big.Int) (*IStakingWalletCosmicSignatureNftEthDepositReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "EthDepositReceived", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCosmicSignatureNftEthDepositReceivedIterator{contract: _IStakingWalletCosmicSignatureNft.contract, event: "EthDepositReceived", logs: logs, sub: sub}, nil
}

// WatchEthDepositReceived is a free log subscription operation binding the contract event 0xb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c0913.
//
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNfts)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) WatchEthDepositReceived(opts *bind.WatchOpts, sink chan<- *IStakingWalletCosmicSignatureNftEthDepositReceived, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "EthDepositReceived", roundNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletCosmicSignatureNftEthDepositReceived)
				if err := _IStakingWalletCosmicSignatureNft.contract.UnpackLog(event, "EthDepositReceived", log); err != nil {
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

// ParseEthDepositReceived is a log parse operation binding the contract event 0xb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c0913.
//
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNfts)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) ParseEthDepositReceived(log types.Log) (*IStakingWalletCosmicSignatureNftEthDepositReceived, error) {
	event := new(IStakingWalletCosmicSignatureNftEthDepositReceived)
	if err := _IStakingWalletCosmicSignatureNft.contract.UnpackLog(event, "EthDepositReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletCosmicSignatureNftNftStakedIterator is returned from FilterNftStaked and is used to iterate over the raw logs and unpacked data for NftStaked events raised by the IStakingWalletCosmicSignatureNft contract.
type IStakingWalletCosmicSignatureNftNftStakedIterator struct {
	Event *IStakingWalletCosmicSignatureNftNftStaked // Event containing the contract specifics and raw log

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
func (it *IStakingWalletCosmicSignatureNftNftStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletCosmicSignatureNftNftStaked)
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
		it.Event = new(IStakingWalletCosmicSignatureNftNftStaked)
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
func (it *IStakingWalletCosmicSignatureNftNftStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletCosmicSignatureNftNftStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletCosmicSignatureNftNftStaked represents a NftStaked event raised by the IStakingWalletCosmicSignatureNft contract.
type IStakingWalletCosmicSignatureNftNftStaked struct {
	StakeActionId *big.Int
	NftTypeCode   uint8
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftStaked is a free log retrieval operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) FilterNftStaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*IStakingWalletCosmicSignatureNftNftStakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCosmicSignatureNftNftStakedIterator{contract: _IStakingWalletCosmicSignatureNft.contract, event: "NftStaked", logs: logs, sub: sub}, nil
}

// WatchNftStaked is a free log subscription operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) WatchNftStaked(opts *bind.WatchOpts, sink chan<- *IStakingWalletCosmicSignatureNftNftStaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletCosmicSignatureNftNftStaked)
				if err := _IStakingWalletCosmicSignatureNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
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

// ParseNftStaked is a log parse operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) ParseNftStaked(log types.Log) (*IStakingWalletCosmicSignatureNftNftStaked, error) {
	event := new(IStakingWalletCosmicSignatureNftNftStaked)
	if err := _IStakingWalletCosmicSignatureNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletCosmicSignatureNftNftUnstakedIterator is returned from FilterNftUnstaked and is used to iterate over the raw logs and unpacked data for NftUnstaked events raised by the IStakingWalletCosmicSignatureNft contract.
type IStakingWalletCosmicSignatureNftNftUnstakedIterator struct {
	Event *IStakingWalletCosmicSignatureNftNftUnstaked // Event containing the contract specifics and raw log

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
func (it *IStakingWalletCosmicSignatureNftNftUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletCosmicSignatureNftNftUnstaked)
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
		it.Event = new(IStakingWalletCosmicSignatureNftNftUnstaked)
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
func (it *IStakingWalletCosmicSignatureNftNftUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletCosmicSignatureNftNftUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletCosmicSignatureNftNftUnstaked represents a NftUnstaked event raised by the IStakingWalletCosmicSignatureNft contract.
type IStakingWalletCosmicSignatureNftNftUnstaked struct {
	StakeActionId            *big.Int
	NftId                    *big.Int
	StakerAddress            common.Address
	NumStakedNfts            *big.Int
	RewardAmount             *big.Int
	MaxUnpaidEthDepositIndex *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNftUnstaked is a free log retrieval operation binding the contract event 0xaf32ebf13e3cd2cc7fa22886a8a6fcfd5b5c6ed9455d773bde63c1dbdc1f64c8.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) FilterNftUnstaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*IStakingWalletCosmicSignatureNftNftUnstakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCosmicSignatureNftNftUnstakedIterator{contract: _IStakingWalletCosmicSignatureNft.contract, event: "NftUnstaked", logs: logs, sub: sub}, nil
}

// WatchNftUnstaked is a free log subscription operation binding the contract event 0xaf32ebf13e3cd2cc7fa22886a8a6fcfd5b5c6ed9455d773bde63c1dbdc1f64c8.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) WatchNftUnstaked(opts *bind.WatchOpts, sink chan<- *IStakingWalletCosmicSignatureNftNftUnstaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletCosmicSignatureNftNftUnstaked)
				if err := _IStakingWalletCosmicSignatureNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
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

// ParseNftUnstaked is a log parse operation binding the contract event 0xaf32ebf13e3cd2cc7fa22886a8a6fcfd5b5c6ed9455d773bde63c1dbdc1f64c8.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) ParseNftUnstaked(log types.Log) (*IStakingWalletCosmicSignatureNftNftUnstaked, error) {
	event := new(IStakingWalletCosmicSignatureNftNftUnstaked)
	if err := _IStakingWalletCosmicSignatureNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletCosmicSignatureNftRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the IStakingWalletCosmicSignatureNft contract.
type IStakingWalletCosmicSignatureNftRewardPaidIterator struct {
	Event *IStakingWalletCosmicSignatureNftRewardPaid // Event containing the contract specifics and raw log

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
func (it *IStakingWalletCosmicSignatureNftRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletCosmicSignatureNftRewardPaid)
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
		it.Event = new(IStakingWalletCosmicSignatureNftRewardPaid)
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
func (it *IStakingWalletCosmicSignatureNftRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletCosmicSignatureNftRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletCosmicSignatureNftRewardPaid represents a RewardPaid event raised by the IStakingWalletCosmicSignatureNft contract.
type IStakingWalletCosmicSignatureNftRewardPaid struct {
	StakeActionId            *big.Int
	NftId                    *big.Int
	StakerAddress            common.Address
	RewardAmount             *big.Int
	MaxUnpaidEthDepositIndex *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xf9bac74bc321a00ef2afeb1f44684e20e22f567699a5840df47967ea88c5b449.
//
// Solidity: event RewardPaid(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) FilterRewardPaid(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*IStakingWalletCosmicSignatureNftRewardPaidIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "RewardPaid", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCosmicSignatureNftRewardPaidIterator{contract: _IStakingWalletCosmicSignatureNft.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xf9bac74bc321a00ef2afeb1f44684e20e22f567699a5840df47967ea88c5b449.
//
// Solidity: event RewardPaid(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *IStakingWalletCosmicSignatureNftRewardPaid, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "RewardPaid", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletCosmicSignatureNftRewardPaid)
				if err := _IStakingWalletCosmicSignatureNft.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xf9bac74bc321a00ef2afeb1f44684e20e22f567699a5840df47967ea88c5b449.
//
// Solidity: event RewardPaid(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) ParseRewardPaid(log types.Log) (*IStakingWalletCosmicSignatureNftRewardPaid, error) {
	event := new(IStakingWalletCosmicSignatureNftRewardPaid)
	if err := _IStakingWalletCosmicSignatureNft.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletNftBaseMetaData contains all meta data concerning the IStakingWalletNftBase contract.
var IStakingWalletNftBaseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicGameConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IStakingWalletNftBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingWalletNftBaseMetaData.ABI instead.
var IStakingWalletNftBaseABI = IStakingWalletNftBaseMetaData.ABI

// IStakingWalletNftBase is an auto generated Go binding around an Ethereum contract.
type IStakingWalletNftBase struct {
	IStakingWalletNftBaseCaller     // Read-only binding to the contract
	IStakingWalletNftBaseTransactor // Write-only binding to the contract
	IStakingWalletNftBaseFilterer   // Log filterer for contract events
}

// IStakingWalletNftBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingWalletNftBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletNftBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingWalletNftBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletNftBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingWalletNftBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletNftBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingWalletNftBaseSession struct {
	Contract     *IStakingWalletNftBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IStakingWalletNftBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingWalletNftBaseCallerSession struct {
	Contract *IStakingWalletNftBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IStakingWalletNftBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingWalletNftBaseTransactorSession struct {
	Contract     *IStakingWalletNftBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IStakingWalletNftBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingWalletNftBaseRaw struct {
	Contract *IStakingWalletNftBase // Generic contract binding to access the raw methods on
}

// IStakingWalletNftBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingWalletNftBaseCallerRaw struct {
	Contract *IStakingWalletNftBaseCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingWalletNftBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingWalletNftBaseTransactorRaw struct {
	Contract *IStakingWalletNftBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakingWalletNftBase creates a new instance of IStakingWalletNftBase, bound to a specific deployed contract.
func NewIStakingWalletNftBase(address common.Address, backend bind.ContractBackend) (*IStakingWalletNftBase, error) {
	contract, err := bindIStakingWalletNftBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletNftBase{IStakingWalletNftBaseCaller: IStakingWalletNftBaseCaller{contract: contract}, IStakingWalletNftBaseTransactor: IStakingWalletNftBaseTransactor{contract: contract}, IStakingWalletNftBaseFilterer: IStakingWalletNftBaseFilterer{contract: contract}}, nil
}

// NewIStakingWalletNftBaseCaller creates a new read-only instance of IStakingWalletNftBase, bound to a specific deployed contract.
func NewIStakingWalletNftBaseCaller(address common.Address, caller bind.ContractCaller) (*IStakingWalletNftBaseCaller, error) {
	contract, err := bindIStakingWalletNftBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletNftBaseCaller{contract: contract}, nil
}

// NewIStakingWalletNftBaseTransactor creates a new write-only instance of IStakingWalletNftBase, bound to a specific deployed contract.
func NewIStakingWalletNftBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingWalletNftBaseTransactor, error) {
	contract, err := bindIStakingWalletNftBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletNftBaseTransactor{contract: contract}, nil
}

// NewIStakingWalletNftBaseFilterer creates a new log filterer instance of IStakingWalletNftBase, bound to a specific deployed contract.
func NewIStakingWalletNftBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingWalletNftBaseFilterer, error) {
	contract, err := bindIStakingWalletNftBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletNftBaseFilterer{contract: contract}, nil
}

// bindIStakingWalletNftBase binds a generic wrapper to an already deployed contract.
func bindIStakingWalletNftBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakingWalletNftBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletNftBase *IStakingWalletNftBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletNftBase.Contract.IStakingWalletNftBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletNftBase *IStakingWalletNftBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletNftBase.Contract.IStakingWalletNftBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletNftBase *IStakingWalletNftBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletNftBase.Contract.IStakingWalletNftBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletNftBase *IStakingWalletNftBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletNftBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletNftBase *IStakingWalletNftBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletNftBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletNftBase *IStakingWalletNftBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletNftBase.Contract.contract.Transact(opts, method, params...)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletNftBase *IStakingWalletNftBaseCaller) NumStakedNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingWalletNftBase.contract.Call(opts, &out, "numStakedNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletNftBase *IStakingWalletNftBaseSession) NumStakedNfts() (*big.Int, error) {
	return _IStakingWalletNftBase.Contract.NumStakedNfts(&_IStakingWalletNftBase.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletNftBase *IStakingWalletNftBaseCallerSession) NumStakedNfts() (*big.Int, error) {
	return _IStakingWalletNftBase.Contract.NumStakedNfts(&_IStakingWalletNftBase.CallOpts)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_IStakingWalletNftBase *IStakingWalletNftBaseCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _IStakingWalletNftBase.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_IStakingWalletNftBase *IStakingWalletNftBaseSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _IStakingWalletNftBase.Contract.WasNftUsed(&_IStakingWalletNftBase.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_IStakingWalletNftBase *IStakingWalletNftBaseCallerSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _IStakingWalletNftBase.Contract.WasNftUsed(&_IStakingWalletNftBase.CallOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletNftBase *IStakingWalletNftBaseTransactor) Stake(opts *bind.TransactOpts, nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletNftBase.contract.Transact(opts, "stake", nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletNftBase *IStakingWalletNftBaseSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletNftBase.Contract.Stake(&_IStakingWalletNftBase.TransactOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletNftBase *IStakingWalletNftBaseTransactorSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletNftBase.Contract.Stake(&_IStakingWalletNftBase.TransactOpts, nftId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletNftBase *IStakingWalletNftBaseTransactor) StakeMany(opts *bind.TransactOpts, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletNftBase.contract.Transact(opts, "stakeMany", nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletNftBase *IStakingWalletNftBaseSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletNftBase.Contract.StakeMany(&_IStakingWalletNftBase.TransactOpts, nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletNftBase *IStakingWalletNftBaseTransactorSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletNftBase.Contract.StakeMany(&_IStakingWalletNftBase.TransactOpts, nftIds_)
}

// IStakingWalletNftBaseNftStakedIterator is returned from FilterNftStaked and is used to iterate over the raw logs and unpacked data for NftStaked events raised by the IStakingWalletNftBase contract.
type IStakingWalletNftBaseNftStakedIterator struct {
	Event *IStakingWalletNftBaseNftStaked // Event containing the contract specifics and raw log

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
func (it *IStakingWalletNftBaseNftStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletNftBaseNftStaked)
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
		it.Event = new(IStakingWalletNftBaseNftStaked)
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
func (it *IStakingWalletNftBaseNftStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletNftBaseNftStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletNftBaseNftStaked represents a NftStaked event raised by the IStakingWalletNftBase contract.
type IStakingWalletNftBaseNftStaked struct {
	StakeActionId *big.Int
	NftTypeCode   uint8
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftStaked is a free log retrieval operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletNftBase *IStakingWalletNftBaseFilterer) FilterNftStaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*IStakingWalletNftBaseNftStakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletNftBase.contract.FilterLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletNftBaseNftStakedIterator{contract: _IStakingWalletNftBase.contract, event: "NftStaked", logs: logs, sub: sub}, nil
}

// WatchNftStaked is a free log subscription operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletNftBase *IStakingWalletNftBaseFilterer) WatchNftStaked(opts *bind.WatchOpts, sink chan<- *IStakingWalletNftBaseNftStaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletNftBase.contract.WatchLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletNftBaseNftStaked)
				if err := _IStakingWalletNftBase.contract.UnpackLog(event, "NftStaked", log); err != nil {
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

// ParseNftStaked is a log parse operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletNftBase *IStakingWalletNftBaseFilterer) ParseNftStaked(log types.Log) (*IStakingWalletNftBaseNftStaked, error) {
	event := new(IStakingWalletNftBaseNftStaked)
	if err := _IStakingWalletNftBase.contract.UnpackLog(event, "NftStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCosmicSignatureNftMetaData contains all meta data concerning the StakingWalletCosmicSignatureNft contract.
var StakingWalletCosmicSignatureNftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"DepositFromUnauthorizedSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"InvalidOperationInCurrentState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftAlreadyUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftNotUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"NftOneTimeStaking\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"NftStakeActionAccessDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftStakeActionInvalidId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NoStakedNfts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit\",\"type\":\"uint256\"}],\"name\":\"NumEthDepositsToEvaluateMaxLimitIsOutOfAllowedRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"EthDepositReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferredToCharity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicGameConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_ETH_DEPOSITS_TO_EVALUATE_HARD_MAX_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"actionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ethDeposits\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"depositId\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"rewardAmountPerStakedNft\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numEthDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numUnpaidStakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetState_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"tryPerformMaintenance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052346100305761001a610014610133565b9061039c565b610022610035565b612fae61054a8239612fae90f35b61003b565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100679061003f565b810190811060018060401b0382111761007f57604052565b610049565b90610097610090610035565b928361005d565b565b5f80fd5b60018060a01b031690565b6100b19061009d565b90565b6100bd906100a8565b90565b6100c9816100b4565b036100d057565b5f80fd5b905051906100e1826100c0565b565b6100ec816100a8565b036100f357565b5f80fd5b90505190610104826100e3565b565b919060408382031261012e578061012261012b925f86016100d4565b936020016100f7565b90565b610099565b6101516134f88038038061014681610084565b928339810190610106565b9091565b90565b61016c6101676101719261009d565b610155565b61009d565b90565b61017d90610158565b90565b61018990610174565b90565b90565b6101a361019e6101a89261018c565b610155565b61009d565b90565b6101b49061018f565b90565b60209181520190565b60207f66745f2e00000000000000000000000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f7220746865206e5f8201520152565b61021a60246040926101b7565b610223816101c0565b0190565b61023c9060208101905f81830391015261020d565b90565b1561024657565b61024e610035565b63eac0d38960e01b81528061026560048201610227565b0390fd5b60207f616d655f2e000000000000000000000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f722074686520675f8201520152565b6102c360256040926101b7565b6102cc81610269565b0190565b6102e59060208101905f8183039101526102b6565b90565b156102ef57565b6102f7610035565b63eac0d38960e01b81528061030e600482016102d0565b0390fd5b5f1b90565b9061032860018060a01b0391610312565b9181191691161790565b61033b90610158565b90565b61034790610332565b90565b90565b9061036261035d6103699261033e565b61034a565b8254610317565b9055565b61037690610174565b90565b90565b9061039161038c6103989261036d565b610379565b8254610317565b9055565b9061040e61041d926103ad3361041f565b6103da6103b982610180565b6103d36103cd6103c85f6101ab565b6100a8565b916100a8565b141561023f565b6103ff836103f86103f26103ed5f6101ab565b6100a8565b916100a8565b14156102e8565b6801000000000000000361034d565b6801000000000000000461037c565b565b6104289061042a565b565b61043390610435565b565b61043e90610462565b565b610449906100a8565b9052565b9190610460905f60208501940190610440565b565b8061047d6104776104725f6101ab565b6100a8565b916100a8565b1461048d5761048b906104ea565b565b6104b06104995f6101ab565b5f918291631e4fbdf760e01b83526004830161044d565b0390fd5b5f1c90565b60018060a01b031690565b6104d06104d5916104b4565b6104b9565b90565b6104e290546104c4565b90565b5f0190565b6104f35f6104d8565b6104fd825f61037c565b9061053161052b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09361036d565b9161036d565b9161053a610035565b80610544816104e5565b0390a356fe60806040526004361015610013575b610c58565b61001d5f3561016c565b806315b4e68f146101675780632e53a259146101625780632ffb33041461015d578063400c89561461015857806347ccca021461015357806357951c741461014e578063715018a6146101495780637a025499146101445780638da5cb5b1461013f5780639e2c8a5b1461013a578063a2b136fb14610135578063a694fc3a14610130578063a6abc99f1461012b578063c2ac898f14610126578063c3fe3e2814610121578063ca7c1f921461011c578063d8ee557314610117578063e0c10d6a14610112578063f2fde38b1461010d578063fdbd98b0146101085763fe939afc0361000e57610c24565b610bb9565b610b6f565b610b1c565b610a59565b6109b1565b61097c565b61090b565b6108d6565b61088c565b610854565b610718565b6106b6565b610653565b6105be565b610589565b61052d565b61043a565b610404565b610286565b6101c9565b60e01c90565b60405190565b5f80fd5b5f80fd5b90565b61018c81610180565b0361019357565b5f80fd5b905035906101a482610183565b565b906020828203126101bf576101bc915f01610197565b90565b610178565b5f0190565b6101dc6101d73660046101a6565b61117c565b6101e4610172565b806101ee816101c4565b0390f35b5f80fd5b5f80fd5b5f80fd5b5f80fd5b909182601f8301121561023c5781359167ffffffffffffffff831161023757602001926020830284011161023257565b6101fe565b6101fa565b6101f6565b91604083830312610281575f83013567ffffffffffffffff811161027c5761026e83610279928601610202565b939094602001610197565b90565b61017c565b610178565b346102b55761029f610299366004610241565b916114c7565b6102a7610172565b806102b1816101c4565b0390f35b6101f2565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b6102e8816102ce565b821015610302576102fa6001916102dc565b910201905f90565b6102ba565b5f1c90565b67ffffffffffffffff1690565b61032561032a91610307565b61030c565b90565b6103379054610319565b90565b60401c90565b60018060c01b031690565b61035761035c9161033a565b610340565b90565b610369905461034b565b90565b6804000000000000000790610380826102ce565b8110156103ab57610390916102df565b50906103a85f6103a181850161032d565b930161035f565b90565b5f80fd5b67ffffffffffffffff1690565b6103c5906103af565b9052565b60018060c01b031690565b6103dd906103c9565b9052565b9160206104029294936103fb60408201965f8301906103bc565b01906103d4565b565b346104355761041c6104173660046101a6565b61036c565b90610431610428610172565b928392836103e1565b0390f35b6101f2565b346104695761045361044d366004610241565b91611594565b61045b610172565b80610465816101c4565b0390f35b6101f2565b5f91031261047857565b610178565b1c90565b60018060a01b031690565b61049c9060086104a1930261047d565b610481565b90565b906104af915461048c565b90565b6104c6680100000000000000035f906104a4565b90565b60018060a01b031690565b90565b6104eb6104e66104f0926104c9565b6104d4565b6104c9565b90565b6104fc906104d7565b90565b610508906104f3565b90565b610514906104ff565b9052565b919061052b905f6020850194019061050b565b565b3461055d5761053d36600461046e565b6105596105486104b2565b610550610172565b91829182610518565b0390f35b6101f2565b151590565b61057090610562565b9052565b9190610587905f60208501940190610567565b565b346105b9576105b56105a461059f3660046101a6565b6116c5565b6105ac610172565b91829182610574565b0390f35b6101f2565b346105ec576105ce36600461046e565b6105d6611734565b6105de610172565b806105e8816101c4565b0390f35b6101f2565b90565b610604906008610609930261047d565b6105f1565b90565b9061061791546105f4565b90565b61062e680500000000000000075f9061060c565b90565b61063a90610180565b9052565b9190610651905f60208501940190610631565b565b346106835761066336600461046e565b61067f61066e61061a565b610676610172565b9182918261063e565b0390f35b6101f2565b610691906104c9565b90565b61069d90610688565b9052565b91906106b4905f60208501940190610694565b565b346106e6576106c636600461046e565b6106e26106d1611742565b6106d9610172565b918291826106a1565b0390f35b6101f2565b91906040838203126107135780610707610710925f8601610197565b93602001610197565b90565b610178565b346107475761073161072b3660046106eb565b90611757565b610739610172565b80610743816101c4565b0390f35b6101f2565b506801000000000000000090565b90565b6107668161074c565b8210156107805761077860039161075a565b910201905f90565b6102ba565b61079161079691610307565b6105f1565b90565b6107a39054610785565b90565b60018060a01b031690565b6107bd6107c291610307565b6107a6565b90565b6107cf90546107b1565b90565b68010000000000000005906107e68261074c565b81101561081e576107f69161075d565b506108025f8201610799565b9161081b6002610814600185016107c5565b9301610799565b90565b5f80fd5b60409061084b610852949695939661084160608401985f850190610631565b6020830190610694565b0190610631565b565b346108875761088361086f61086a3660046101a6565b6107d2565b61087a939193610172565b93849384610822565b0390f35b6101f2565b346108ba576108a461089f3660046101a6565b611b2b565b6108ac610172565b806108b6816101c4565b0390f35b6101f2565b6108d3680400000000000000055f9061060c565b90565b34610906576108e636600461046e565b6109026108f16108bf565b6108f9610172565b9182918261063e565b0390f35b6101f2565b3461093a5761092461091e3660046106eb565b90611d46565b61092c610172565b80610936816101c4565b0390f35b6101f2565b61094f906008610954930261047d565b6107a6565b90565b90610962915461093f565b90565b610979680100000000000000045f90610957565b90565b346109ac5761098c36600461046e565b6109a8610997610965565b61099f610172565b918291826106a1565b0390f35b6101f2565b346109e1576109c136600461046e565b6109dd6109cc611da7565b6109d4610172565b9182918261063e565b0390f35b6101f2565b6109ef81610562565b036109f657565b5f80fd5b90503590610a07826109e6565b565b610a1281610688565b03610a1957565b5f80fd5b90503590610a2a82610a09565b565b9190604083820312610a545780610a48610a51925f86016109fa565b93602001610a1d565b90565b610178565b34610a8a57610a86610a75610a6f366004610a2c565b90612130565b610a7d610172565b91829182610574565b0390f35b6101f2565b90565b610aa6610aa1610aab92610a8f565b6104d4565b610180565b90565b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b610ae2610ae891610180565b91610180565b908115610af3570490565b610aae565b610b0e5f19610b08610100610a92565b90610ad6565b90565b610b19610af8565b90565b34610b4c57610b2c36600461046e565b610b48610b37610b11565b610b3f610172565b9182918261063e565b0390f35b6101f2565b90602082820312610b6a57610b67915f01610a1d565b90565b610178565b34610b9d57610b87610b82366004610b51565b6121aa565b610b8f610172565b80610b99816101c4565b0390f35b6101f2565b610bb6680100000000000000025f9061060c565b90565b34610be957610bc936600461046e565b610be5610bd4610ba2565b610bdc610172565b9182918261063e565b0390f35b6101f2565b90602082820312610c1f575f82013567ffffffffffffffff8111610c1a57610c169201610202565b9091565b61017c565b610178565b34610c5357610c3d610c37366004610bee565b906121b5565b610c45610172565b80610c4f816101c4565b0390f35b6101f2565b5f80fd5b60209181520190565b60207f7065726d697474656420746f206d616b652061206465706f7369742e00000000917f4f6e6c792074686520436f736d696347616d6520636f6e7472616374206973205f8201520152565b610cbf603c604092610c5c565b610cc881610c65565b0190565b9190610cef906020610ce7604086018681035f880152610cb2565b940190610694565b565b15610cf95750565b610d1b90610d05610172565b918291637ed5977760e11b835260048301610ccc565b0390fd5b90565b610d36610d31610d3b92610d1f565b6104d4565b610180565b90565b5f7f546865726520617265206e6f207374616b6564204e4654732e00000000000000910152565b610d726019602092610c5c565b610d7b81610d3e565b0190565b610d949060208101905f818303910152610d65565b90565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610dbf90610d97565b810190811067ffffffffffffffff821117610dd957604052565b610da1565b90610df1610dea610172565b9283610db5565b565b610dfd6040610dde565b90565b5f90565b5f90565b610e10610df3565b9060208083610e1d610e00565b815201610e28610e04565b81525050565b610e36610e08565b90565b90565b610e50610e4b610e5592610e39565b6104d4565b610180565b90565b610e67610e6d91939293610180565b92610180565b8201809211610e7857565b610ac2565b5f1b90565b90610e8e5f1991610e7d565b9181191691161790565b610eac610ea7610eb192610180565b6104d4565b610180565b90565b90565b90610ecc610ec7610ed392610e98565b610eb4565b8254610e82565b9055565b90565b610eee610ee9610ef392610ed7565b6104d4565b610180565b90565b90610f00906103af565b9052565b90610f0e906103c9565b9052565b90610f48610f3f5f610f22610df3565b94610f39610f3183830161032d565b838801610ef6565b0161035f565b60208401610f04565b565b610f5390610f12565b90565b610f6a610f65610f6f92610180565b6104d4565b6103c9565b90565b610f7c90516103c9565b90565b610f8b610f91916103c9565b916103c9565b019060018060c01b038211610fa257565b610ac2565b610fb090610180565b5f198114610fbe5760010190565b610ac2565b610fd7610fd2610fdc92610180565b6104d4565b6103af565b90565b634e487b7160e01b5f525f60045260245ffd5b610ffc90516103af565b90565b9061101267ffffffffffffffff91610e7d565b9181191691161790565b61103061102b611035926103af565b6104d4565b6103af565b90565b90565b9061105061104b6110579261101c565b611038565b8254610fff565b9055565b60401b90565b9061107567ffffffffffffffff199161105b565b9181191691161790565b61109361108e611098926103c9565b6104d4565b6103c9565b90565b90565b906110b36110ae6110ba9261107f565b61109b565b8254611061565b9055565b906110e860205f6110ee946110e08282016110da848801610ff2565b9061103b565b019201610f72565b9061109e565b565b9190611101576110ff916110be565b565b610fdf565b61111a61111561111f926103af565b6104d4565b610180565b90565b61112b90611106565b9052565b9095949261117a946111696111739261115f60809661115560a088019c5f890190610631565b6020870190610631565b6040850190611122565b6060830190610631565b0190610631565b565b6111ab336111a361119d611198680100000000000000046107c5565b610688565b91610688565b143390610cf1565b6111b56001610799565b90816111c96111c35f610d22565b91610180565b1461138c576111d6610e2e565b6111e868050000000000000007610799565b9161120e6111fe68010000000000000002610799565b6112086001610e3c565b90610e58565b926112228468010000000000000002610eb7565b61123468040000000000000006610799565b6112476112416002610eda565b91610180565b10155f1461132d576112ea5f61127b611328936112766112676001610e3c565b68040000000000000006610eb7565b610fa7565b9361128f8568050000000000000007610eb7565b6112a361129b88610fc3565b838801610ef6565b6112c26112b96112b4348b90610ad6565b610f56565b60208801610f04565b5b6112e1866112db6804000000000000000788906102df565b906110f0565b95939401610ff2565b94346113167fb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c091396610e98565b9661131f610172565b9586958661112f565b0390a2565b9091506113286112ea5f61135461134e6804000000000000000786906102df565b50610f4a565b9461138761136b611366348b90610ad6565b610f56565b611381602089019161137c83610f72565b610f7f565b90610f04565b6112c3565b611394610172565b630f13059960e41b8152806113ab60048201610d7f565b0390fd5b60207f5f206973206f7574206f662074686520616c6c6f7765642072616e67652e0000917f6e756d4574684465706f73697473546f4576616c756174654d61784c696d69745f8201520152565b611409603e604092610c5c565b611412816113af565b0190565b9190611439906020611431604086018681035f8801526113fc565b940190610631565b565b156114435750565b6114659061144f610172565b91829163285cf1f560e01b835260048301611416565b0390fd5b5090565b61147c61148291939293610180565b92610180565b820391821161148d57565b610ac2565b600161149e9101610180565b90565b5f90565b91908110156114b5576020020190565b6102ba565b356114c481610183565b90565b9291611504906114f3816114ea6114e46114df610af8565b610180565b91610180565b1115829061143b565b6114fe858490611469565b9061146d565b9061150e5f610d22565b6115175f610d22565b925b8361153661153061152b898790611469565b610180565b91610180565b10156115835761157661157061154e61157c93610fa7565b6115566114a1565b5061156b6115668a888a916114a5565b6114ba565b612547565b93610e58565b93611492565b9290611519565b50915050611592919250612838565b565b92916115d1906115c0816115b76115b16115ac610af8565b610180565b91610180565b1115829061143b565b6115cb858490611469565b9061146d565b906115db5f610d22565b6115e45f610d22565b925b836116036115fd6115f8898790611469565b610180565b91610180565b10156116505761164361163d61161b61164993610fa7565b6116236114a1565b506116386116338a888a916114a5565b6114ba565b6129cb565b93610e58565b93611492565b92906115e6565b5091505061165f919250612838565b565b5f90565b506801000000000000000090565b90565b61167f81611665565b82101561169957611691600191611673565b910201905f90565b6102ba565b60ff1690565b6116b06116b591610307565b61169e565b90565b6116c290546116a4565b90565b5f6116dd6116e4926116d5611661565b506002611676565b50016116b8565b90565b6116ef612cd3565b6116f7611721565b565b61170d61170861171292610d1f565b6104d4565b6104c9565b90565b61171e906116f9565b90565b61173261172d5f611715565b612d21565b565b61173c6116e7565b565b5f90565b61174a61173e565b506117545f6107c5565b90565b61178f91611789918161177261176c5f610d22565b91610180565b1180611791575b61178490839061143b565b6129cb565b50612838565b565b50611784826117af6117a96117a4610af8565b610180565b91610180565b11159050611779565b60407f206f6e6c79206f6e63652e000000000000000000000000000000000000000000917f54686973204e46542068617320616c7265616479206265656e207374616b65645f8201527f2e20416e204e465420697320616c6c6f77656420746f206265207374616b656460208201520152565b611838604b606092610c5c565b611841816117b8565b0190565b9190611868906020611860604086018681035f88015261182b565b940190610631565b565b156118725750565b6118949061187e610172565b918291633b471d1f60e21b835260048301611845565b0390fd5b90565b906118ac60018060a01b0391610e7d565b9181191691161790565b6118bf906104f3565b90565b90565b906118da6118d56118e1926118b6565b6118c2565b825461189b565b9055565b6118ef6040610dde565b90565b906118fc90610562565b9052565b60018060f81b031690565b61191f61191a61192492610d1f565b6104d4565b611900565b90565b9061193190611900565b9052565b61193f9051610562565b90565b9061194e60ff91610e7d565b9181191691161790565b61196190610562565b90565b90565b9061197c61197761198392611958565b611964565b8254611942565b9055565b6119919051611900565b90565b60081b90565b906119a760ff1991611994565b9181191691161790565b6119c56119c06119ca92611900565b6104d4565b611900565b90565b90565b906119e56119e06119ec926119b1565b6119cd565b825461199a565b9055565b90611a1a60205f611a2094611a12828201611a0c848801611935565b90611967565b019201611987565b906119d0565b565b9190611a3357611a31916119f0565b565b610fdf565b611a44611a4991610307565b610481565b90565b611a569054611a38565b90565b611a62906104f3565b90565b5f80fd5b60e01b90565b5f910312611a7957565b610178565b604090611aa7611aae9496959396611a9d60608401985f850190610694565b6020830190610694565b0190610631565b565b611ab8610172565b3d5f823e3d90fd5b634e487b7160e01b5f52602160045260245ffd5b60031115611ade57565b611ac0565b90611aed82611ad4565b565b611af890611ae3565b90565b611b0490611aef565b9052565b916020611b29929493611b2260408201965f830190611afb565b0190610631565b565b611b55611b4e611b485f611b4160028690611676565b50016116b8565b15610562565b829061186a565b611b7a611b6a68010000000000000002610799565b611b746001610e3c565b90610e58565b611b8d8168010000000000000002610eb7565b611bc3611bad611ba768010000000000000005849061075d565b50611898565b611bb9845f8301610eb7565b60013391016118c5565b611be0611bd06001610799565b611bda6001610e3c565b90610e58565b90611bec826001610eb7565b611c08611bf96002610eda565b68040000000000000006610eb7565b611c486001611c36611c2d5f611c28611c1f6118e5565b945f86016118f2565b61190b565b60208301611927565b611c4260028690611676565b90611a22565b611c62611c5d68010000000000000003611a4c565b6104ff565b6323b872dd33611c7130611a59565b928692813b15611d41575f611c9991611ca48296611c8d610172565b98899788968795611a69565b855260048501611a7e565b03925af18015611d3c57611d10575b50600192903392611cf6611cf0611cea7fcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f0082994610e98565b94610e98565b946118b6565b94611d0b611d02610172565b92839283611b08565b0390a4565b611d2f905f3d8111611d35575b611d278183610db5565b810190611a6f565b5f611cb3565b503d611d1d565b611ab0565b611a65565b611d7e91611d789181611d61611d5b5f610d22565b91610180565b1180611d80575b611d7390839061143b565b612547565b50612838565b565b50611d7382611d9e611d98611d93610af8565b610180565b91610180565b11159050611d68565b611daf6114a1565b50611dba6001610799565b90565b90611dd09291611dcb612cd3565b611fc7565b90565b5f7f546865726520617265207374696c6c207374616b6564204e4654732e00000000910152565b611e07601c602092610c5c565b611e1081611dd3565b0190565b611e299060208101905f818303910152611dfa565b90565b15611e3357565b611e3b610172565b63a29f5c4d60e01b815280611e5260048201611e14565b0390fd5b5f7f546865726520617265207374696c6c20756e7061696420726577617264732e00910152565b611e8a601f602092610c5c565b611e9381611e56565b0190565b611eac9060208101905f818303910152611e7d565b90565b15611eb657565b611ebe610172565b63a29f5c4d60e01b815280611ed560048201611e97565b0390fd5b905090565b611ee95f8092611ed9565b0190565b611ef690611ede565b90565b67ffffffffffffffff8111611f1757611f13602091610d97565b0190565b610da1565b90611f2e611f2983611ef9565b610dde565b918252565b606090565b3d5f14611f5357611f483d611f1c565b903d5f602084013e5b565b611f5b611f33565b90611f51565b5f7f5472616e7366657220746f2063686172697479206661696c65642e0000000000910152565b611f95601b602092610c5c565b611f9e81611f61565b0190565b9190611fc5906020611fbd604086018681035f880152611f88565b940190610631565b565b50611fed611fd56001610799565b611fe7611fe15f610d22565b91610180565b14611e2c565b61201a61200268040000000000000005610799565b61201461200e5f610d22565b91610180565b14611eaf565b612110575b6001908061203d6120376120325f611715565b610688565b91610688565b03612047575b5090565b61205030611a59565b31905f80828461205e610172565b908161206981611eed565b03925af1612075611f38565b505f146120c5576120bb6120a97f1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d926118b6565b926120b2610172565b9182918261063e565b0390a25b5f612043565b9091506121076120f57f154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a926118b6565b926120fe610172565b91829182611fa2565b0390a25f6120bf565b61212b61211c5f610d22565b68050000000000000007610eb7565b61201f565b906121429161213d611661565b611dbd565b90565b61215690612151612cd3565b612158565b565b8061217361216d6121685f611715565b610688565b91610688565b146121835761218190612d21565b565b6121a661218f5f611715565b5f918291631e4fbdf760e01b8352600483016106a1565b0390fd5b6121b390612145565b565b9190916121c15f610d22565b5b806121df6121d96121d4858890611469565b610180565b91610180565b101561220f5761220a906122056122006121fb858885916114a5565b6114ba565b611b2b565b611492565b6121c2565b50509050565b9061221f90610180565b9052565b9061222d90610688565b9052565b61223b6060610dde565b90565b9061228d612284600261224f612231565b9461226661225e5f8301610799565b5f8801612215565b61227e612275600183016107c5565b60208801612223565b01610799565b60408401612215565b565b6122989061223e565b90565b6122a59051610688565b90565b5f7f496e76616c6964204e4654207374616b6520616374696f6e2049442e00000000910152565b6122dc601c602092610c5c565b6122e5816122a8565b0190565b919061230c906020612304604086018681035f8801526122cf565b940190610631565b565b60207f656365697665207374616b696e67207265776172642e00000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20725f8201520152565b6123686036604092610c5c565b6123718161230e565b0190565b9160406123a692949361239f612394606083018381035f85015261235b565b966020830190610631565b0190610694565b565b6123b29051610180565b90565b5f7f4e465420686173206e6f74206265656e20756e7374616b65642e000000000000910152565b6123e9601a602092610c5c565b6123f2816123b5565b0190565b9190612419906020612411604086018681035f8801526123dc565b940190610631565b565b156124235750565b6124459061242f610172565b9182916372c1673f60e01b8352600483016123f6565b0390fd5b1b90565b919060086124689102916124625f1984612449565b92612449565b9181191691161790565b919061248861248361249093610e98565b610eb4565b90835461244d565b9055565b6124a6916124a06114a1565b91612472565b565b919060086124c89102916124c260018060a01b0384612449565b92612449565b9181191691161790565b91906124e86124e36124f0936118b6565b6118c2565b9083546124a8565b9055565b6125069161250061173e565b916124d2565b565b61251190610180565b5f811461251f576001900390565b610ac2565b91602061254592949361253e60408201965f830190610631565b0190610631565b565b916125506114a1565b506125596114a1565b5061257761257168010000000000000005859061075d565b50611898565b916125818361228f565b90336125a061259a6125956020860161229b565b610688565b91610688565b036126f457936125ec6125fb94956125d76125bd604086016123a8565b6125cf6125c95f610d22565b91610180565b11849061241b565b82906125e5604086016123a8565b9091612e2c565b60408597929701909690612215565b9461261461260b604085016123a8565b60028301610eb7565b612620604084016123a8565b61263261262c5f610d22565b91610180565b146126a9575b506126445f83016123a8565b339161265360408895016123a8565b61268f6126896126837ff9bac74bc321a00ef2afeb1f44684e20e22f567699a5840df47967ea88c5b44994610e98565b94610e98565b946118b6565b946126a461269b610172565b92839283612524565b0390a4565b5f6001826126bc83806126c29601612494565b016124f4565b6126ee6126df6126da68040000000000000005610799565b612508565b68040000000000000005610eb7565b5f612638565b846127016020840161229b565b61271b6127156127105f611715565b610688565b91610688565b14155f1461274a573390612746612730610172565b9283926348aca7ef60e11b845260048401612375565b0390fd5b61276c90612756610172565b91829163023df6b160e21b8352600483016122e9565b0390fd5b60207f642e000000000000000000000000000000000000000000000000000000000000917f4e4654207374616b696e6720726577617264207061796d656e74206661696c655f8201520152565b6127ca6022604092610c5c565b6127d381612770565b0190565b9160406128089294936128016127f6606083018381035f8501526127bd565b966020830190610694565b0190610631565b565b15612813575050565b61283461281e610172565b928392630aa7db6360e11b8452600484016127d7565b0390fd5b612869905f803383612848610172565b908161285381611eed565b03925af161285f611f38565b509033909161280a565b565b60207f6e7374616b652069742e00000000000000000000000000000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20755f8201520152565b6128c5602a604092610c5c565b6128ce8161286b565b0190565b9160406129039294936128fc6128f1606083018381035f8501526128b8565b966020830190610631565b0190610694565b565b5f7f4e46542068617320616c7265616479206265656e20756e7374616b65642e0000910152565b612939601e602092610c5c565b61294281612905565b0190565b9190612969906020612961604086018681035f88015261292c565b940190610631565b565b156129735750565b6129959061297f610172565b91829163c339165960e01b835260048301612946565b0390fd5b6040906129c26129c994969593966129b860608401985f850190610631565b6020830190610631565b0190610631565b565b916129d46114a1565b506129dd6114a1565b506129fb6129f568010000000000000005859061075d565b50611898565b91612a058361228f565b9033612a24612a1e612a196020860161229b565b610688565b91610688565b03612c575793612a76612a859495612a5b612a41604086016123a8565b612a53612a4d5f610d22565b91610180565b14849061296b565b8290612a6f68050000000000000007610799565b9091612e2c565b60408597929701909690612215565b94612a92604084016123a8565b612aa4612a9e5f610d22565b91610180565b115f14612c3957612ac4906002612abd604086016123a8565b9101610eb7565b612af0612ae1612adc68040000000000000005610799565b610fa7565b68040000000000000005610eb7565b5b612b0e612afe6001610799565b612b086001610e3c565b9061146d565b90612b1a826001610eb7565b612b34612b2f68010000000000000003611a4c565b6104ff565b6323b872dd612b4230611a59565b3392612b4f5f88016123a8565b92813b15612c34575f612b7591612b808296612b69610172565b98899788968795611a69565b855260048501611a7e565b03925af18015612c2f57612c03575b50612b9b5f84016123a8565b90339293612bac60408992016123a8565b94612bfe612bec612be6612be07faf32ebf13e3cd2cc7fa22886a8a6fcfd5b5c6ed9455d773bde63c1dbdc1f64c896610e98565b96610e98565b966118b6565b96612bf5610172565b93849384612999565b0390a4565b612c22905f3d8111612c28575b612c1a8183610db5565b810190611a6f565b5f612b8f565b503d612c10565b611ab0565b611a65565b5f600182612c4c8380612c529601612494565b016124f4565b612af1565b84612c646020840161229b565b612c7e612c78612c735f611715565b610688565b91610688565b14155f14612cad573390612ca9612c93610172565b9283926348aca7ef60e11b8452600484016128d2565b0390fd5b612ccf90612cb9610172565b91829163023df6b160e21b8352600483016122e9565b0390fd5b612cdb611742565b612cf4612cee612ce9612f6b565b610688565b91610688565b03612cfb57565b612d1d612d06612f6b565b5f91829163118cdaa760e01b8352600483016106a1565b0390fd5b612d2a5f6107c5565b612d34825f6118c5565b90612d68612d627f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936118b6565b916118b6565b91612d71610172565b80612d7b816101c4565b0390a3565b90565b612d97612d92612d9c92610180565b6104d4565b612d80565b90565b90612daa9103612d80565b90565b612dc1612dbc612dc692612d80565b6104d4565b610180565b90565b612ddd612dd8612de292610d1f565b6104d4565b612d80565b90565b6001612df19103610180565b90565b612e08612e03612e0d926103c9565b6104d4565b610180565b90565b90612e1b9101610180565b90565b90612e299103610180565b90565b92612e356114a1565b92612e3e6114a1565b50612e476114a1565b50612e6b612e66612e5785612d83565b612e6084612d83565b90612d9f565b612dad565b94612e7586612d83565b612e8f612e89612e845f612dc9565b612d80565b91612d80565b12612f5a575b9483955b86612eac612ea684610180565b91610180565b115f14612f4957612ed9612ed3612ecd680400000000000000078a906102df565b50610f4a565b97612de5565b96612ee55f8201610ff2565b612ef7612ef184610180565b91611106565b11612f2257505050612f19612f1f92939495612f125f610d22565b955b612e1e565b90612e1e565b90565b95612f3c612f376020612f4294959901610f72565b612df4565b90612e10565b9490612e99565b509394612f1990612f1f9394612f14565b9450612f655f610d22565b94612e95565b612f7361173e565b50339056fea2646970667358221220191fc65ebd605698d11dbaf0c99d1613ceafb4be120160d2db2126717d5fcd8064736f6c634300081b0033",
}

// StakingWalletCosmicSignatureNftABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingWalletCosmicSignatureNftMetaData.ABI instead.
var StakingWalletCosmicSignatureNftABI = StakingWalletCosmicSignatureNftMetaData.ABI

// StakingWalletCosmicSignatureNftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingWalletCosmicSignatureNftMetaData.Bin instead.
var StakingWalletCosmicSignatureNftBin = StakingWalletCosmicSignatureNftMetaData.Bin

// DeployStakingWalletCosmicSignatureNft deploys a new Ethereum contract, binding an instance of StakingWalletCosmicSignatureNft to it.
func DeployStakingWalletCosmicSignatureNft(auth *bind.TransactOpts, backend bind.ContractBackend, nft_ common.Address, game_ common.Address) (common.Address, *types.Transaction, *StakingWalletCosmicSignatureNft, error) {
	parsed, err := StakingWalletCosmicSignatureNftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingWalletCosmicSignatureNftBin), backend, nft_, game_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingWalletCosmicSignatureNft{StakingWalletCosmicSignatureNftCaller: StakingWalletCosmicSignatureNftCaller{contract: contract}, StakingWalletCosmicSignatureNftTransactor: StakingWalletCosmicSignatureNftTransactor{contract: contract}, StakingWalletCosmicSignatureNftFilterer: StakingWalletCosmicSignatureNftFilterer{contract: contract}}, nil
}

// StakingWalletCosmicSignatureNft is an auto generated Go binding around an Ethereum contract.
type StakingWalletCosmicSignatureNft struct {
	StakingWalletCosmicSignatureNftCaller     // Read-only binding to the contract
	StakingWalletCosmicSignatureNftTransactor // Write-only binding to the contract
	StakingWalletCosmicSignatureNftFilterer   // Log filterer for contract events
}

// StakingWalletCosmicSignatureNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingWalletCosmicSignatureNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletCosmicSignatureNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingWalletCosmicSignatureNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletCosmicSignatureNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingWalletCosmicSignatureNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletCosmicSignatureNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingWalletCosmicSignatureNftSession struct {
	Contract     *StakingWalletCosmicSignatureNft // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// StakingWalletCosmicSignatureNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingWalletCosmicSignatureNftCallerSession struct {
	Contract *StakingWalletCosmicSignatureNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// StakingWalletCosmicSignatureNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingWalletCosmicSignatureNftTransactorSession struct {
	Contract     *StakingWalletCosmicSignatureNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// StakingWalletCosmicSignatureNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingWalletCosmicSignatureNftRaw struct {
	Contract *StakingWalletCosmicSignatureNft // Generic contract binding to access the raw methods on
}

// StakingWalletCosmicSignatureNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingWalletCosmicSignatureNftCallerRaw struct {
	Contract *StakingWalletCosmicSignatureNftCaller // Generic read-only contract binding to access the raw methods on
}

// StakingWalletCosmicSignatureNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingWalletCosmicSignatureNftTransactorRaw struct {
	Contract *StakingWalletCosmicSignatureNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingWalletCosmicSignatureNft creates a new instance of StakingWalletCosmicSignatureNft, bound to a specific deployed contract.
func NewStakingWalletCosmicSignatureNft(address common.Address, backend bind.ContractBackend) (*StakingWalletCosmicSignatureNft, error) {
	contract, err := bindStakingWalletCosmicSignatureNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNft{StakingWalletCosmicSignatureNftCaller: StakingWalletCosmicSignatureNftCaller{contract: contract}, StakingWalletCosmicSignatureNftTransactor: StakingWalletCosmicSignatureNftTransactor{contract: contract}, StakingWalletCosmicSignatureNftFilterer: StakingWalletCosmicSignatureNftFilterer{contract: contract}}, nil
}

// NewStakingWalletCosmicSignatureNftCaller creates a new read-only instance of StakingWalletCosmicSignatureNft, bound to a specific deployed contract.
func NewStakingWalletCosmicSignatureNftCaller(address common.Address, caller bind.ContractCaller) (*StakingWalletCosmicSignatureNftCaller, error) {
	contract, err := bindStakingWalletCosmicSignatureNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftCaller{contract: contract}, nil
}

// NewStakingWalletCosmicSignatureNftTransactor creates a new write-only instance of StakingWalletCosmicSignatureNft, bound to a specific deployed contract.
func NewStakingWalletCosmicSignatureNftTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingWalletCosmicSignatureNftTransactor, error) {
	contract, err := bindStakingWalletCosmicSignatureNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftTransactor{contract: contract}, nil
}

// NewStakingWalletCosmicSignatureNftFilterer creates a new log filterer instance of StakingWalletCosmicSignatureNft, bound to a specific deployed contract.
func NewStakingWalletCosmicSignatureNftFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingWalletCosmicSignatureNftFilterer, error) {
	contract, err := bindStakingWalletCosmicSignatureNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftFilterer{contract: contract}, nil
}

// bindStakingWalletCosmicSignatureNft binds a generic wrapper to an already deployed contract.
func bindStakingWalletCosmicSignatureNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingWalletCosmicSignatureNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletCosmicSignatureNft.Contract.StakingWalletCosmicSignatureNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.StakingWalletCosmicSignatureNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.StakingWalletCosmicSignatureNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletCosmicSignatureNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.contract.Transact(opts, method, params...)
}

// NUMETHDEPOSITSTOEVALUATEHARDMAXLIMIT is a free data retrieval call binding the contract method 0xe0c10d6a.
//
// Solidity: function NUM_ETH_DEPOSITS_TO_EVALUATE_HARD_MAX_LIMIT() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) NUMETHDEPOSITSTOEVALUATEHARDMAXLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "NUM_ETH_DEPOSITS_TO_EVALUATE_HARD_MAX_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NUMETHDEPOSITSTOEVALUATEHARDMAXLIMIT is a free data retrieval call binding the contract method 0xe0c10d6a.
//
// Solidity: function NUM_ETH_DEPOSITS_TO_EVALUATE_HARD_MAX_LIMIT() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) NUMETHDEPOSITSTOEVALUATEHARDMAXLIMIT() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.NUMETHDEPOSITSTOEVALUATEHARDMAXLIMIT(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// NUMETHDEPOSITSTOEVALUATEHARDMAXLIMIT is a free data retrieval call binding the contract method 0xe0c10d6a.
//
// Solidity: function NUM_ETH_DEPOSITS_TO_EVALUATE_HARD_MAX_LIMIT() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) NUMETHDEPOSITSTOEVALUATEHARDMAXLIMIT() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.NUMETHDEPOSITSTOEVALUATEHARDMAXLIMIT(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) ActionCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "actionCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) ActionCounter() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.ActionCounter(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) ActionCounter() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.ActionCounter(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// EthDeposits is a free data retrieval call binding the contract method 0x2ffb3304.
//
// Solidity: function ethDeposits(uint256 ) view returns(uint64 depositId, uint192 rewardAmountPerStakedNft)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) EthDeposits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DepositId                uint64
	RewardAmountPerStakedNft *big.Int
}, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "ethDeposits", arg0)

	outstruct := new(struct {
		DepositId                uint64
		RewardAmountPerStakedNft *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.RewardAmountPerStakedNft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EthDeposits is a free data retrieval call binding the contract method 0x2ffb3304.
//
// Solidity: function ethDeposits(uint256 ) view returns(uint64 depositId, uint192 rewardAmountPerStakedNft)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) EthDeposits(arg0 *big.Int) (struct {
	DepositId                uint64
	RewardAmountPerStakedNft *big.Int
}, error) {
	return _StakingWalletCosmicSignatureNft.Contract.EthDeposits(&_StakingWalletCosmicSignatureNft.CallOpts, arg0)
}

// EthDeposits is a free data retrieval call binding the contract method 0x2ffb3304.
//
// Solidity: function ethDeposits(uint256 ) view returns(uint64 depositId, uint192 rewardAmountPerStakedNft)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) EthDeposits(arg0 *big.Int) (struct {
	DepositId                uint64
	RewardAmountPerStakedNft *big.Int
}, error) {
	return _StakingWalletCosmicSignatureNft.Contract.EthDeposits(&_StakingWalletCosmicSignatureNft.CallOpts, arg0)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) Game() (common.Address, error) {
	return _StakingWalletCosmicSignatureNft.Contract.Game(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) Game() (common.Address, error) {
	return _StakingWalletCosmicSignatureNft.Contract.Game(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) Nft() (common.Address, error) {
	return _StakingWalletCosmicSignatureNft.Contract.Nft(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) Nft() (common.Address, error) {
	return _StakingWalletCosmicSignatureNft.Contract.Nft(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// NumEthDeposits is a free data retrieval call binding the contract method 0x7a025499.
//
// Solidity: function numEthDeposits() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) NumEthDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "numEthDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumEthDeposits is a free data retrieval call binding the contract method 0x7a025499.
//
// Solidity: function numEthDeposits() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) NumEthDeposits() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.NumEthDeposits(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// NumEthDeposits is a free data retrieval call binding the contract method 0x7a025499.
//
// Solidity: function numEthDeposits() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) NumEthDeposits() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.NumEthDeposits(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) NumStakedNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "numStakedNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) NumStakedNfts() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.NumStakedNfts(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) NumStakedNfts() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.NumStakedNfts(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// NumUnpaidStakeActions is a free data retrieval call binding the contract method 0xa6abc99f.
//
// Solidity: function numUnpaidStakeActions() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) NumUnpaidStakeActions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "numUnpaidStakeActions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumUnpaidStakeActions is a free data retrieval call binding the contract method 0xa6abc99f.
//
// Solidity: function numUnpaidStakeActions() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) NumUnpaidStakeActions() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.NumUnpaidStakeActions(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// NumUnpaidStakeActions is a free data retrieval call binding the contract method 0xa6abc99f.
//
// Solidity: function numUnpaidStakeActions() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) NumUnpaidStakeActions() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.NumUnpaidStakeActions(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) Owner() (common.Address, error) {
	return _StakingWalletCosmicSignatureNft.Contract.Owner(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) Owner() (common.Address, error) {
	return _StakingWalletCosmicSignatureNft.Contract.Owner(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 nftId, address nftOwnerAddress, uint256 maxUnpaidEthDepositIndex)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) StakeActions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NftId                    *big.Int
	NftOwnerAddress          common.Address
	MaxUnpaidEthDepositIndex *big.Int
}, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "stakeActions", arg0)

	outstruct := new(struct {
		NftId                    *big.Int
		NftOwnerAddress          common.Address
		MaxUnpaidEthDepositIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NftId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftOwnerAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.MaxUnpaidEthDepositIndex = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 nftId, address nftOwnerAddress, uint256 maxUnpaidEthDepositIndex)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) StakeActions(arg0 *big.Int) (struct {
	NftId                    *big.Int
	NftOwnerAddress          common.Address
	MaxUnpaidEthDepositIndex *big.Int
}, error) {
	return _StakingWalletCosmicSignatureNft.Contract.StakeActions(&_StakingWalletCosmicSignatureNft.CallOpts, arg0)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 nftId, address nftOwnerAddress, uint256 maxUnpaidEthDepositIndex)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) StakeActions(arg0 *big.Int) (struct {
	NftId                    *big.Int
	NftOwnerAddress          common.Address
	MaxUnpaidEthDepositIndex *big.Int
}, error) {
	return _StakingWalletCosmicSignatureNft.Contract.StakeActions(&_StakingWalletCosmicSignatureNft.CallOpts, arg0)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _StakingWalletCosmicSignatureNft.Contract.WasNftUsed(&_StakingWalletCosmicSignatureNft.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _StakingWalletCosmicSignatureNft.Contract.WasNftUsed(&_StakingWalletCosmicSignatureNft.CallOpts, nftId_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactor) DepositIfPossible(opts *bind.TransactOpts, roundNum_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.contract.Transact(opts, "depositIfPossible", roundNum_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) DepositIfPossible(roundNum_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.DepositIfPossible(&_StakingWalletCosmicSignatureNft.TransactOpts, roundNum_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorSession) DepositIfPossible(roundNum_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.DepositIfPossible(&_StakingWalletCosmicSignatureNft.TransactOpts, roundNum_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0x2e53a259.
//
// Solidity: function payManyRewards(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactor) PayManyRewards(opts *bind.TransactOpts, stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.contract.Transact(opts, "payManyRewards", stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0x2e53a259.
//
// Solidity: function payManyRewards(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) PayManyRewards(stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.PayManyRewards(&_StakingWalletCosmicSignatureNft.TransactOpts, stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0x2e53a259.
//
// Solidity: function payManyRewards(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorSession) PayManyRewards(stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.PayManyRewards(&_StakingWalletCosmicSignatureNft.TransactOpts, stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// PayReward is a paid mutator transaction binding the contract method 0xc2ac898f.
//
// Solidity: function payReward(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactor) PayReward(opts *bind.TransactOpts, stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.contract.Transact(opts, "payReward", stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// PayReward is a paid mutator transaction binding the contract method 0xc2ac898f.
//
// Solidity: function payReward(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) PayReward(stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.PayReward(&_StakingWalletCosmicSignatureNft.TransactOpts, stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// PayReward is a paid mutator transaction binding the contract method 0xc2ac898f.
//
// Solidity: function payReward(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorSession) PayReward(stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.PayReward(&_StakingWalletCosmicSignatureNft.TransactOpts, stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.RenounceOwnership(&_StakingWalletCosmicSignatureNft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.RenounceOwnership(&_StakingWalletCosmicSignatureNft.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactor) Stake(opts *bind.TransactOpts, nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.contract.Transact(opts, "stake", nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.Stake(&_StakingWalletCosmicSignatureNft.TransactOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.Stake(&_StakingWalletCosmicSignatureNft.TransactOpts, nftId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactor) StakeMany(opts *bind.TransactOpts, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.contract.Transact(opts, "stakeMany", nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.StakeMany(&_StakingWalletCosmicSignatureNft.TransactOpts, nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.StakeMany(&_StakingWalletCosmicSignatureNft.TransactOpts, nftIds_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.TransferOwnership(&_StakingWalletCosmicSignatureNft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.TransferOwnership(&_StakingWalletCosmicSignatureNft.TransactOpts, newOwner)
}

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactor) TryPerformMaintenance(opts *bind.TransactOpts, resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.contract.Transact(opts, "tryPerformMaintenance", resetState_, charityAddress_)
}

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) TryPerformMaintenance(resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.TryPerformMaintenance(&_StakingWalletCosmicSignatureNft.TransactOpts, resetState_, charityAddress_)
}

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorSession) TryPerformMaintenance(resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.TryPerformMaintenance(&_StakingWalletCosmicSignatureNft.TransactOpts, resetState_, charityAddress_)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactor) Unstake(opts *bind.TransactOpts, stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.contract.Transact(opts, "unstake", stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) Unstake(stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.Unstake(&_StakingWalletCosmicSignatureNft.TransactOpts, stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 stakeActionId_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorSession) Unstake(stakeActionId_ *big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.Unstake(&_StakingWalletCosmicSignatureNft.TransactOpts, stakeActionId_, numEthDepositsToEvaluateMaxLimit_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x400c8956.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactor) UnstakeMany(opts *bind.TransactOpts, stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.contract.Transact(opts, "unstakeMany", stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x400c8956.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) UnstakeMany(stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.UnstakeMany(&_StakingWalletCosmicSignatureNft.TransactOpts, stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x400c8956.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_, uint256 numEthDepositsToEvaluateMaxLimit_) returns()
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftTransactorSession) UnstakeMany(stakeActionIds_ []*big.Int, numEthDepositsToEvaluateMaxLimit_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCosmicSignatureNft.Contract.UnstakeMany(&_StakingWalletCosmicSignatureNft.TransactOpts, stakeActionIds_, numEthDepositsToEvaluateMaxLimit_)
}

// StakingWalletCosmicSignatureNftEthDepositReceivedIterator is returned from FilterEthDepositReceived and is used to iterate over the raw logs and unpacked data for EthDepositReceived events raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftEthDepositReceivedIterator struct {
	Event *StakingWalletCosmicSignatureNftEthDepositReceived // Event containing the contract specifics and raw log

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
func (it *StakingWalletCosmicSignatureNftEthDepositReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCosmicSignatureNftEthDepositReceived)
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
		it.Event = new(StakingWalletCosmicSignatureNftEthDepositReceived)
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
func (it *StakingWalletCosmicSignatureNftEthDepositReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCosmicSignatureNftEthDepositReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCosmicSignatureNftEthDepositReceived represents a EthDepositReceived event raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftEthDepositReceived struct {
	RoundNum      *big.Int
	ActionCounter *big.Int
	DepositIndex  *big.Int
	DepositId     *big.Int
	DepositAmount *big.Int
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthDepositReceived is a free log retrieval operation binding the contract event 0xb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c0913.
//
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNfts)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) FilterEthDepositReceived(opts *bind.FilterOpts, roundNum []*big.Int) (*StakingWalletCosmicSignatureNftEthDepositReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "EthDepositReceived", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftEthDepositReceivedIterator{contract: _StakingWalletCosmicSignatureNft.contract, event: "EthDepositReceived", logs: logs, sub: sub}, nil
}

// WatchEthDepositReceived is a free log subscription operation binding the contract event 0xb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c0913.
//
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNfts)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) WatchEthDepositReceived(opts *bind.WatchOpts, sink chan<- *StakingWalletCosmicSignatureNftEthDepositReceived, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "EthDepositReceived", roundNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCosmicSignatureNftEthDepositReceived)
				if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "EthDepositReceived", log); err != nil {
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

// ParseEthDepositReceived is a log parse operation binding the contract event 0xb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c0913.
//
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNfts)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) ParseEthDepositReceived(log types.Log) (*StakingWalletCosmicSignatureNftEthDepositReceived, error) {
	event := new(StakingWalletCosmicSignatureNftEthDepositReceived)
	if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "EthDepositReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCosmicSignatureNftFundTransferFailedIterator is returned from FilterFundTransferFailed and is used to iterate over the raw logs and unpacked data for FundTransferFailed events raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftFundTransferFailedIterator struct {
	Event *StakingWalletCosmicSignatureNftFundTransferFailed // Event containing the contract specifics and raw log

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
func (it *StakingWalletCosmicSignatureNftFundTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCosmicSignatureNftFundTransferFailed)
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
		it.Event = new(StakingWalletCosmicSignatureNftFundTransferFailed)
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
func (it *StakingWalletCosmicSignatureNftFundTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCosmicSignatureNftFundTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCosmicSignatureNftFundTransferFailed represents a FundTransferFailed event raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftFundTransferFailed struct {
	ErrStr             string
	DestinationAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFundTransferFailed is a free log retrieval operation binding the contract event 0x154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a.
//
// Solidity: event FundTransferFailed(string errStr, address indexed destinationAddress, uint256 amount)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) FilterFundTransferFailed(opts *bind.FilterOpts, destinationAddress []common.Address) (*StakingWalletCosmicSignatureNftFundTransferFailedIterator, error) {

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "FundTransferFailed", destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftFundTransferFailedIterator{contract: _StakingWalletCosmicSignatureNft.contract, event: "FundTransferFailed", logs: logs, sub: sub}, nil
}

// WatchFundTransferFailed is a free log subscription operation binding the contract event 0x154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a.
//
// Solidity: event FundTransferFailed(string errStr, address indexed destinationAddress, uint256 amount)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) WatchFundTransferFailed(opts *bind.WatchOpts, sink chan<- *StakingWalletCosmicSignatureNftFundTransferFailed, destinationAddress []common.Address) (event.Subscription, error) {

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "FundTransferFailed", destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCosmicSignatureNftFundTransferFailed)
				if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "FundTransferFailed", log); err != nil {
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

// ParseFundTransferFailed is a log parse operation binding the contract event 0x154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a.
//
// Solidity: event FundTransferFailed(string errStr, address indexed destinationAddress, uint256 amount)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) ParseFundTransferFailed(log types.Log) (*StakingWalletCosmicSignatureNftFundTransferFailed, error) {
	event := new(StakingWalletCosmicSignatureNftFundTransferFailed)
	if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "FundTransferFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCosmicSignatureNftFundsTransferredToCharityIterator is returned from FilterFundsTransferredToCharity and is used to iterate over the raw logs and unpacked data for FundsTransferredToCharity events raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftFundsTransferredToCharityIterator struct {
	Event *StakingWalletCosmicSignatureNftFundsTransferredToCharity // Event containing the contract specifics and raw log

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
func (it *StakingWalletCosmicSignatureNftFundsTransferredToCharityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCosmicSignatureNftFundsTransferredToCharity)
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
		it.Event = new(StakingWalletCosmicSignatureNftFundsTransferredToCharity)
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
func (it *StakingWalletCosmicSignatureNftFundsTransferredToCharityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCosmicSignatureNftFundsTransferredToCharityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCosmicSignatureNftFundsTransferredToCharity represents a FundsTransferredToCharity event raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftFundsTransferredToCharity struct {
	CharityAddress common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFundsTransferredToCharity is a free log retrieval operation binding the contract event 0x1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d.
//
// Solidity: event FundsTransferredToCharity(address indexed charityAddress, uint256 amount)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) FilterFundsTransferredToCharity(opts *bind.FilterOpts, charityAddress []common.Address) (*StakingWalletCosmicSignatureNftFundsTransferredToCharityIterator, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "FundsTransferredToCharity", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftFundsTransferredToCharityIterator{contract: _StakingWalletCosmicSignatureNft.contract, event: "FundsTransferredToCharity", logs: logs, sub: sub}, nil
}

// WatchFundsTransferredToCharity is a free log subscription operation binding the contract event 0x1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d.
//
// Solidity: event FundsTransferredToCharity(address indexed charityAddress, uint256 amount)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) WatchFundsTransferredToCharity(opts *bind.WatchOpts, sink chan<- *StakingWalletCosmicSignatureNftFundsTransferredToCharity, charityAddress []common.Address) (event.Subscription, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "FundsTransferredToCharity", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCosmicSignatureNftFundsTransferredToCharity)
				if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "FundsTransferredToCharity", log); err != nil {
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

// ParseFundsTransferredToCharity is a log parse operation binding the contract event 0x1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d.
//
// Solidity: event FundsTransferredToCharity(address indexed charityAddress, uint256 amount)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) ParseFundsTransferredToCharity(log types.Log) (*StakingWalletCosmicSignatureNftFundsTransferredToCharity, error) {
	event := new(StakingWalletCosmicSignatureNftFundsTransferredToCharity)
	if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "FundsTransferredToCharity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCosmicSignatureNftNftStakedIterator is returned from FilterNftStaked and is used to iterate over the raw logs and unpacked data for NftStaked events raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftNftStakedIterator struct {
	Event *StakingWalletCosmicSignatureNftNftStaked // Event containing the contract specifics and raw log

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
func (it *StakingWalletCosmicSignatureNftNftStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCosmicSignatureNftNftStaked)
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
		it.Event = new(StakingWalletCosmicSignatureNftNftStaked)
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
func (it *StakingWalletCosmicSignatureNftNftStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCosmicSignatureNftNftStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCosmicSignatureNftNftStaked represents a NftStaked event raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftNftStaked struct {
	StakeActionId *big.Int
	NftTypeCode   uint8
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftStaked is a free log retrieval operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) FilterNftStaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*StakingWalletCosmicSignatureNftNftStakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftNftStakedIterator{contract: _StakingWalletCosmicSignatureNft.contract, event: "NftStaked", logs: logs, sub: sub}, nil
}

// WatchNftStaked is a free log subscription operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) WatchNftStaked(opts *bind.WatchOpts, sink chan<- *StakingWalletCosmicSignatureNftNftStaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCosmicSignatureNftNftStaked)
				if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
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

// ParseNftStaked is a log parse operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) ParseNftStaked(log types.Log) (*StakingWalletCosmicSignatureNftNftStaked, error) {
	event := new(StakingWalletCosmicSignatureNftNftStaked)
	if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCosmicSignatureNftNftUnstakedIterator is returned from FilterNftUnstaked and is used to iterate over the raw logs and unpacked data for NftUnstaked events raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftNftUnstakedIterator struct {
	Event *StakingWalletCosmicSignatureNftNftUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingWalletCosmicSignatureNftNftUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCosmicSignatureNftNftUnstaked)
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
		it.Event = new(StakingWalletCosmicSignatureNftNftUnstaked)
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
func (it *StakingWalletCosmicSignatureNftNftUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCosmicSignatureNftNftUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCosmicSignatureNftNftUnstaked represents a NftUnstaked event raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftNftUnstaked struct {
	StakeActionId            *big.Int
	NftId                    *big.Int
	StakerAddress            common.Address
	NumStakedNfts            *big.Int
	RewardAmount             *big.Int
	MaxUnpaidEthDepositIndex *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNftUnstaked is a free log retrieval operation binding the contract event 0xaf32ebf13e3cd2cc7fa22886a8a6fcfd5b5c6ed9455d773bde63c1dbdc1f64c8.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) FilterNftUnstaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*StakingWalletCosmicSignatureNftNftUnstakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftNftUnstakedIterator{contract: _StakingWalletCosmicSignatureNft.contract, event: "NftUnstaked", logs: logs, sub: sub}, nil
}

// WatchNftUnstaked is a free log subscription operation binding the contract event 0xaf32ebf13e3cd2cc7fa22886a8a6fcfd5b5c6ed9455d773bde63c1dbdc1f64c8.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) WatchNftUnstaked(opts *bind.WatchOpts, sink chan<- *StakingWalletCosmicSignatureNftNftUnstaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCosmicSignatureNftNftUnstaked)
				if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
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

// ParseNftUnstaked is a log parse operation binding the contract event 0xaf32ebf13e3cd2cc7fa22886a8a6fcfd5b5c6ed9455d773bde63c1dbdc1f64c8.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) ParseNftUnstaked(log types.Log) (*StakingWalletCosmicSignatureNftNftUnstaked, error) {
	event := new(StakingWalletCosmicSignatureNftNftUnstaked)
	if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCosmicSignatureNftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftOwnershipTransferredIterator struct {
	Event *StakingWalletCosmicSignatureNftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingWalletCosmicSignatureNftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCosmicSignatureNftOwnershipTransferred)
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
		it.Event = new(StakingWalletCosmicSignatureNftOwnershipTransferred)
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
func (it *StakingWalletCosmicSignatureNftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCosmicSignatureNftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCosmicSignatureNftOwnershipTransferred represents a OwnershipTransferred event raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingWalletCosmicSignatureNftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftOwnershipTransferredIterator{contract: _StakingWalletCosmicSignatureNft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingWalletCosmicSignatureNftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCosmicSignatureNftOwnershipTransferred)
				if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) ParseOwnershipTransferred(log types.Log) (*StakingWalletCosmicSignatureNftOwnershipTransferred, error) {
	event := new(StakingWalletCosmicSignatureNftOwnershipTransferred)
	if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCosmicSignatureNftRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftRewardPaidIterator struct {
	Event *StakingWalletCosmicSignatureNftRewardPaid // Event containing the contract specifics and raw log

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
func (it *StakingWalletCosmicSignatureNftRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCosmicSignatureNftRewardPaid)
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
		it.Event = new(StakingWalletCosmicSignatureNftRewardPaid)
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
func (it *StakingWalletCosmicSignatureNftRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCosmicSignatureNftRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCosmicSignatureNftRewardPaid represents a RewardPaid event raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftRewardPaid struct {
	StakeActionId            *big.Int
	NftId                    *big.Int
	StakerAddress            common.Address
	RewardAmount             *big.Int
	MaxUnpaidEthDepositIndex *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xf9bac74bc321a00ef2afeb1f44684e20e22f567699a5840df47967ea88c5b449.
//
// Solidity: event RewardPaid(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) FilterRewardPaid(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*StakingWalletCosmicSignatureNftRewardPaidIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "RewardPaid", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftRewardPaidIterator{contract: _StakingWalletCosmicSignatureNft.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xf9bac74bc321a00ef2afeb1f44684e20e22f567699a5840df47967ea88c5b449.
//
// Solidity: event RewardPaid(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *StakingWalletCosmicSignatureNftRewardPaid, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "RewardPaid", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCosmicSignatureNftRewardPaid)
				if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xf9bac74bc321a00ef2afeb1f44684e20e22f567699a5840df47967ea88c5b449.
//
// Solidity: event RewardPaid(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) ParseRewardPaid(log types.Log) (*StakingWalletCosmicSignatureNftRewardPaid, error) {
	event := new(StakingWalletCosmicSignatureNftRewardPaid)
	if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletNftBaseMetaData contains all meta data concerning the StakingWalletNftBase contract.
var StakingWalletNftBaseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicGameConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingWalletNftBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingWalletNftBaseMetaData.ABI instead.
var StakingWalletNftBaseABI = StakingWalletNftBaseMetaData.ABI

// StakingWalletNftBase is an auto generated Go binding around an Ethereum contract.
type StakingWalletNftBase struct {
	StakingWalletNftBaseCaller     // Read-only binding to the contract
	StakingWalletNftBaseTransactor // Write-only binding to the contract
	StakingWalletNftBaseFilterer   // Log filterer for contract events
}

// StakingWalletNftBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingWalletNftBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletNftBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingWalletNftBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletNftBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingWalletNftBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletNftBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingWalletNftBaseSession struct {
	Contract     *StakingWalletNftBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StakingWalletNftBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingWalletNftBaseCallerSession struct {
	Contract *StakingWalletNftBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// StakingWalletNftBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingWalletNftBaseTransactorSession struct {
	Contract     *StakingWalletNftBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// StakingWalletNftBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingWalletNftBaseRaw struct {
	Contract *StakingWalletNftBase // Generic contract binding to access the raw methods on
}

// StakingWalletNftBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingWalletNftBaseCallerRaw struct {
	Contract *StakingWalletNftBaseCaller // Generic read-only contract binding to access the raw methods on
}

// StakingWalletNftBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingWalletNftBaseTransactorRaw struct {
	Contract *StakingWalletNftBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingWalletNftBase creates a new instance of StakingWalletNftBase, bound to a specific deployed contract.
func NewStakingWalletNftBase(address common.Address, backend bind.ContractBackend) (*StakingWalletNftBase, error) {
	contract, err := bindStakingWalletNftBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingWalletNftBase{StakingWalletNftBaseCaller: StakingWalletNftBaseCaller{contract: contract}, StakingWalletNftBaseTransactor: StakingWalletNftBaseTransactor{contract: contract}, StakingWalletNftBaseFilterer: StakingWalletNftBaseFilterer{contract: contract}}, nil
}

// NewStakingWalletNftBaseCaller creates a new read-only instance of StakingWalletNftBase, bound to a specific deployed contract.
func NewStakingWalletNftBaseCaller(address common.Address, caller bind.ContractCaller) (*StakingWalletNftBaseCaller, error) {
	contract, err := bindStakingWalletNftBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletNftBaseCaller{contract: contract}, nil
}

// NewStakingWalletNftBaseTransactor creates a new write-only instance of StakingWalletNftBase, bound to a specific deployed contract.
func NewStakingWalletNftBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingWalletNftBaseTransactor, error) {
	contract, err := bindStakingWalletNftBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletNftBaseTransactor{contract: contract}, nil
}

// NewStakingWalletNftBaseFilterer creates a new log filterer instance of StakingWalletNftBase, bound to a specific deployed contract.
func NewStakingWalletNftBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingWalletNftBaseFilterer, error) {
	contract, err := bindStakingWalletNftBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingWalletNftBaseFilterer{contract: contract}, nil
}

// bindStakingWalletNftBase binds a generic wrapper to an already deployed contract.
func bindStakingWalletNftBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingWalletNftBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletNftBase *StakingWalletNftBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletNftBase.Contract.StakingWalletNftBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletNftBase *StakingWalletNftBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletNftBase.Contract.StakingWalletNftBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletNftBase *StakingWalletNftBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletNftBase.Contract.StakingWalletNftBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletNftBase *StakingWalletNftBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletNftBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletNftBase *StakingWalletNftBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletNftBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletNftBase *StakingWalletNftBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletNftBase.Contract.contract.Transact(opts, method, params...)
}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletNftBase *StakingWalletNftBaseCaller) ActionCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletNftBase.contract.Call(opts, &out, "actionCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletNftBase *StakingWalletNftBaseSession) ActionCounter() (*big.Int, error) {
	return _StakingWalletNftBase.Contract.ActionCounter(&_StakingWalletNftBase.CallOpts)
}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletNftBase *StakingWalletNftBaseCallerSession) ActionCounter() (*big.Int, error) {
	return _StakingWalletNftBase.Contract.ActionCounter(&_StakingWalletNftBase.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletNftBase *StakingWalletNftBaseCaller) NumStakedNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletNftBase.contract.Call(opts, &out, "numStakedNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletNftBase *StakingWalletNftBaseSession) NumStakedNfts() (*big.Int, error) {
	return _StakingWalletNftBase.Contract.NumStakedNfts(&_StakingWalletNftBase.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletNftBase *StakingWalletNftBaseCallerSession) NumStakedNfts() (*big.Int, error) {
	return _StakingWalletNftBase.Contract.NumStakedNfts(&_StakingWalletNftBase.CallOpts)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_StakingWalletNftBase *StakingWalletNftBaseCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletNftBase.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_StakingWalletNftBase *StakingWalletNftBaseSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _StakingWalletNftBase.Contract.WasNftUsed(&_StakingWalletNftBase.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_StakingWalletNftBase *StakingWalletNftBaseCallerSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _StakingWalletNftBase.Contract.WasNftUsed(&_StakingWalletNftBase.CallOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletNftBase *StakingWalletNftBaseTransactor) Stake(opts *bind.TransactOpts, nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletNftBase.contract.Transact(opts, "stake", nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletNftBase *StakingWalletNftBaseSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletNftBase.Contract.Stake(&_StakingWalletNftBase.TransactOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletNftBase *StakingWalletNftBaseTransactorSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletNftBase.Contract.Stake(&_StakingWalletNftBase.TransactOpts, nftId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletNftBase *StakingWalletNftBaseTransactor) StakeMany(opts *bind.TransactOpts, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletNftBase.contract.Transact(opts, "stakeMany", nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletNftBase *StakingWalletNftBaseSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletNftBase.Contract.StakeMany(&_StakingWalletNftBase.TransactOpts, nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletNftBase *StakingWalletNftBaseTransactorSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletNftBase.Contract.StakeMany(&_StakingWalletNftBase.TransactOpts, nftIds_)
}

// StakingWalletNftBaseNftStakedIterator is returned from FilterNftStaked and is used to iterate over the raw logs and unpacked data for NftStaked events raised by the StakingWalletNftBase contract.
type StakingWalletNftBaseNftStakedIterator struct {
	Event *StakingWalletNftBaseNftStaked // Event containing the contract specifics and raw log

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
func (it *StakingWalletNftBaseNftStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletNftBaseNftStaked)
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
		it.Event = new(StakingWalletNftBaseNftStaked)
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
func (it *StakingWalletNftBaseNftStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletNftBaseNftStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletNftBaseNftStaked represents a NftStaked event raised by the StakingWalletNftBase contract.
type StakingWalletNftBaseNftStaked struct {
	StakeActionId *big.Int
	NftTypeCode   uint8
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftStaked is a free log retrieval operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletNftBase *StakingWalletNftBaseFilterer) FilterNftStaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*StakingWalletNftBaseNftStakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletNftBase.contract.FilterLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletNftBaseNftStakedIterator{contract: _StakingWalletNftBase.contract, event: "NftStaked", logs: logs, sub: sub}, nil
}

// WatchNftStaked is a free log subscription operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletNftBase *StakingWalletNftBaseFilterer) WatchNftStaked(opts *bind.WatchOpts, sink chan<- *StakingWalletNftBaseNftStaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletNftBase.contract.WatchLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletNftBaseNftStaked)
				if err := _StakingWalletNftBase.contract.UnpackLog(event, "NftStaked", log); err != nil {
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

// ParseNftStaked is a log parse operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletNftBase *StakingWalletNftBaseFilterer) ParseNftStaked(log types.Log) (*StakingWalletNftBaseNftStaked, error) {
	event := new(StakingWalletNftBaseNftStaked)
	if err := _StakingWalletNftBase.contract.UnpackLog(event, "NftStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
