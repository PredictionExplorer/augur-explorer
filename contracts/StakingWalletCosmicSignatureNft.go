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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"EthDepositReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIStakingWalletNftBase.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStateResets\",\"type\":\"uint256\"}],\"name\":\"StateReset\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetState_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"tryPerformMaintenance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IStakingWalletCosmicSignatureNft.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
	return _IStakingWalletCosmicSignatureNft.Contract.WasNftUsed(&_IStakingWalletCosmicSignatureNft.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftCallerSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
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
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 indexed depositId, uint256 depositAmount, uint256 numStakedNfts)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) FilterEthDepositReceived(opts *bind.FilterOpts, roundNum []*big.Int, depositId []*big.Int) (*IStakingWalletCosmicSignatureNftEthDepositReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "EthDepositReceived", roundNumRule, depositIdRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCosmicSignatureNftEthDepositReceivedIterator{contract: _IStakingWalletCosmicSignatureNft.contract, event: "EthDepositReceived", logs: logs, sub: sub}, nil
}

// WatchEthDepositReceived is a free log subscription operation binding the contract event 0xb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c0913.
//
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 indexed depositId, uint256 depositAmount, uint256 numStakedNfts)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) WatchEthDepositReceived(opts *bind.WatchOpts, sink chan<- *IStakingWalletCosmicSignatureNftEthDepositReceived, roundNum []*big.Int, depositId []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "EthDepositReceived", roundNumRule, depositIdRule)
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
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 indexed depositId, uint256 depositAmount, uint256 numStakedNfts)
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
	ActionCounter            *big.Int
	StakeActionId            *big.Int
	NftId                    *big.Int
	StakerAddress            common.Address
	NumStakedNfts            *big.Int
	RewardAmount             *big.Int
	MaxUnpaidEthDepositIndex *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNftUnstaked is a free log retrieval operation binding the contract event 0xec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
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

// WatchNftUnstaked is a free log subscription operation binding the contract event 0xec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
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

// ParseNftUnstaked is a log parse operation binding the contract event 0xec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
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

// IStakingWalletCosmicSignatureNftStateResetIterator is returned from FilterStateReset and is used to iterate over the raw logs and unpacked data for StateReset events raised by the IStakingWalletCosmicSignatureNft contract.
type IStakingWalletCosmicSignatureNftStateResetIterator struct {
	Event *IStakingWalletCosmicSignatureNftStateReset // Event containing the contract specifics and raw log

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
func (it *IStakingWalletCosmicSignatureNftStateResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletCosmicSignatureNftStateReset)
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
		it.Event = new(IStakingWalletCosmicSignatureNftStateReset)
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
func (it *IStakingWalletCosmicSignatureNftStateResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletCosmicSignatureNftStateResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletCosmicSignatureNftStateReset represents a StateReset event raised by the IStakingWalletCosmicSignatureNft contract.
type IStakingWalletCosmicSignatureNftStateReset struct {
	NumStateResets *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStateReset is a free log retrieval operation binding the contract event 0x8a77a30ae54ca88ceda5f17cd3679c1b868abd2c6d36b67bcbe5e9beac7ec648.
//
// Solidity: event StateReset(uint256 numStateResets)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) FilterStateReset(opts *bind.FilterOpts) (*IStakingWalletCosmicSignatureNftStateResetIterator, error) {

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "StateReset")
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCosmicSignatureNftStateResetIterator{contract: _IStakingWalletCosmicSignatureNft.contract, event: "StateReset", logs: logs, sub: sub}, nil
}

// WatchStateReset is a free log subscription operation binding the contract event 0x8a77a30ae54ca88ceda5f17cd3679c1b868abd2c6d36b67bcbe5e9beac7ec648.
//
// Solidity: event StateReset(uint256 numStateResets)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) WatchStateReset(opts *bind.WatchOpts, sink chan<- *IStakingWalletCosmicSignatureNftStateReset) (event.Subscription, error) {

	logs, sub, err := _IStakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "StateReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletCosmicSignatureNftStateReset)
				if err := _IStakingWalletCosmicSignatureNft.contract.UnpackLog(event, "StateReset", log); err != nil {
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

// ParseStateReset is a log parse operation binding the contract event 0x8a77a30ae54ca88ceda5f17cd3679c1b868abd2c6d36b67bcbe5e9beac7ec648.
//
// Solidity: event StateReset(uint256 numStateResets)
func (_IStakingWalletCosmicSignatureNft *IStakingWalletCosmicSignatureNftFilterer) ParseStateReset(log types.Log) (*IStakingWalletCosmicSignatureNftStateReset, error) {
	event := new(IStakingWalletCosmicSignatureNftStateReset)
	if err := _IStakingWalletCosmicSignatureNft.contract.UnpackLog(event, "StateReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletNftBaseMetaData contains all meta data concerning the IStakingWalletNftBase contract.
var IStakingWalletNftBaseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIStakingWalletNftBase.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_IStakingWalletNftBase *IStakingWalletNftBaseCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IStakingWalletNftBase.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_IStakingWalletNftBase *IStakingWalletNftBaseSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
	return _IStakingWalletNftBase.Contract.WasNftUsed(&_IStakingWalletNftBase.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_IStakingWalletNftBase *IStakingWalletNftBaseCallerSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignatureNft\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"InvalidOperationInCurrentState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftAlreadyUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"NftHasAlreadyBeenStaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftNotUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"NftStakeActionAccessDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftStakeActionInvalidId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NoStakedNfts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit\",\"type\":\"uint256\"}],\"name\":\"NumEthDepositsToEvaluateMaxLimitIsOutOfAllowedRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"UnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"EthDepositReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferredToCharity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIStakingWalletNftBase.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStateResets\",\"type\":\"uint256\"}],\"name\":\"StateReset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ethDeposits\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"depositId\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"rewardAmountPerStakedNft\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numEthDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStateResets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numUnpaidStakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetState_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"tryPerformMaintenance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052346100575761001a61001461015a565b906102ee565b61002261005c565b612ec36104ab823960805181818161047f01528181611b390152612a94015260a0518181816109420152610ca90152612ec390f35b610062565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061008e90610066565b810190811060018060401b038211176100a657604052565b610070565b906100be6100b761005c565b9283610084565b565b5f80fd5b60018060a01b031690565b6100d8906100c4565b90565b6100e4906100cf565b90565b6100f0816100db565b036100f757565b5f80fd5b90505190610108826100e7565b565b610113816100cf565b0361011a57565b5f80fd5b9050519061012b8261010a565b565b91906040838203126101555780610149610152925f86016100fb565b9360200161011e565b90565b6100c0565b61017861336e8038038061016d816100ab565b92833981019061012d565b9091565b90565b61019361018e610198926100c4565b61017c565b6100c4565b90565b6101a49061017f565b90565b6101b09061019b565b90565b90565b6101ca6101c56101cf926101b3565b61017c565b6100c4565b90565b6101db906101b6565b90565b60209181520190565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b61021b601d6020926101de565b610224816101e7565b0190565b61023d9060208101905f81830391015261020e565b90565b9061024a826101a7565b61026461025e6102595f6101d2565b6100cf565b916100cf565b146102745761027291610297565b565b61027c61005c565b63eac0d38960e01b81528061029360048201610228565b0390fd5b90806102b36102ad6102a85f6101d2565b6100cf565b916100cf565b146102c3576102c1916102e6565b565b6102cb61005c565b63eac0d38960e01b8152806102e260048201610228565b0390fd5b60805260a052565b90610308916103036102fe6103b9565b61030a565b610240565b565b61031390610315565b565b61031e90610320565b565b6103299061032b565b565b61033490610336565b565b61033f90610363565b565b61034a906100cf565b9052565b9190610361905f60208501940190610341565b565b8061037e6103786103735f6101d2565b6100cf565b916100cf565b1461038e5761038c9061044b565b565b6103b161039a5f6101d2565b5f918291631e4fbdf760e01b83526004830161034e565b0390fd5b5f90565b6103c16103b5565b503390565b5f1c90565b60018060a01b031690565b6103e26103e7916103c6565b6103cb565b90565b6103f490546103d6565b90565b5f1b90565b9061040d60018060a01b03916103f7565b9181191691161790565b6104209061019b565b90565b90565b9061043b61043661044292610417565b610423565b82546103fc565b9055565b5f0190565b6104545f6103ea565b61045e825f610426565b9061049261048c7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610417565b91610417565b9161049b61005c565b806104a581610446565b0390a356fe60806040526004361015610013575b610ba5565b61001d5f3561016c565b806315b4e68f146101675780632e53a259146101625780632ffb33041461015d578063400c89561461015857806347ccca021461015357806357951c741461014e578063715018a6146101495780637a025499146101445780637c843a461461013f5780638da5cb5b1461013a5780639e2c8a5b14610135578063a2b136fb14610130578063a694fc3a1461012b578063a6abc99f14610126578063c2ac898f14610121578063c3fe3e281461011c578063ca7c1f9214610117578063d8ee557314610112578063f2fde38b1461010d578063fdbd98b0146101085763fe939afc0361000e57610b71565b610b06565b610abc565b610a68565b610999565b610964565b61090c565b6108d7565b61088d565b610855565b610719565b6106b7565b610654565b610608565b610591565b61055c565b610505565b61043a565b610404565b610286565b6101c9565b60e01c90565b60405190565b5f80fd5b5f80fd5b90565b61018c81610180565b0361019357565b5f80fd5b905035906101a482610183565b565b906020828203126101bf576101bc915f01610197565b90565b610178565b5f0190565b6101dc6101d73660046101a6565b61135d565b6101e4610172565b806101ee816101c4565b0390f35b5f80fd5b5f80fd5b5f80fd5b5f80fd5b909182601f8301121561023c5781359167ffffffffffffffff831161023757602001926020830284011161023257565b6101fe565b6101fa565b6101f6565b91604083830312610281575f83013567ffffffffffffffff811161027c5761026e83610279928601610202565b939094602001610197565b90565b61017c565b610178565b346102b55761029f610299366004610241565b916114b8565b6102a7610172565b806102b1816101c4565b0390f35b6101f2565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b6102e8816102ce565b821015610302576102fa6001916102dc565b910201905f90565b6102ba565b5f1c90565b67ffffffffffffffff1690565b61032561032a91610307565b61030c565b90565b6103379054610319565b90565b60401c90565b60018060c01b031690565b61035761035c9161033a565b610340565b90565b610369905461034b565b90565b6804000000000000000690610380826102ce565b8110156103ab57610390916102df565b50906103a85f6103a181850161032d565b930161035f565b90565b5f80fd5b67ffffffffffffffff1690565b6103c5906103af565b9052565b60018060c01b031690565b6103dd906103c9565b9052565b9160206104029294936103fb60408201965f8301906103bc565b01906103d4565b565b346104355761041c6104173660046101a6565b61036c565b90610431610428610172565b928392836103e1565b0390f35b6101f2565b346104695761045361044d366004610241565b91611585565b61045b610172565b80610465816101c4565b0390f35b6101f2565b5f91031261047857565b610178565b7f000000000000000000000000000000000000000000000000000000000000000090565b60018060a01b031690565b90565b6104c36104be6104c8926104a1565b6104ac565b6104a1565b90565b6104d4906104af565b90565b6104e0906104cb565b90565b6104ec906104d7565b9052565b9190610503905f602085019401906104e3565b565b346105355761051536600461046e565b61053161052061047d565b610528610172565b918291826104f0565b0390f35b6101f2565b61054390610180565b9052565b919061055a905f6020850194019061053a565b565b3461058c576105886105776105723660046101a6565b61168b565b61057f610172565b91829182610547565b0390f35b6101f2565b346105bf576105a136600461046e565b6105a96116f8565b6105b1610172565b806105bb816101c4565b0390f35b6101f2565b1c90565b90565b6105db9060086105e093026105c4565b6105c8565b90565b906105ee91546105cb565b90565b610605680400000000000000055f906105e3565b90565b346106385761061836600461046e565b6106346106236105f1565b61062b610172565b91829182610547565b0390f35b6101f2565b610651680500000000000000065f906105e3565b90565b346106845761066436600461046e565b61068061066f61063d565b610677610172565b91829182610547565b0390f35b6101f2565b610692906104a1565b90565b61069e90610689565b9052565b91906106b5905f60208501940190610695565b565b346106e7576106c736600461046e565b6106e36106d2611706565b6106da610172565b918291826106a2565b0390f35b6101f2565b91906040838203126107145780610708610711925f8601610197565b93602001610197565b90565b610178565b346107485761073261072c3660046106ec565b9061171b565b61073a610172565b80610744816101c4565b0390f35b6101f2565b506801000000000000000090565b90565b6107678161074d565b8210156107815761077960039161075b565b910201905f90565b6102ba565b61079261079791610307565b6105c8565b90565b6107a49054610786565b90565b60018060a01b031690565b6107be6107c391610307565b6107a7565b90565b6107d090546107b2565b90565b68010000000000000004906107e78261074d565b81101561081f576107f79161075e565b506108035f820161079a565b9161081c6002610815600185016107c6565b930161079a565b90565b5f80fd5b60409061084c610853949695939661084260608401985f85019061053a565b6020830190610695565b019061053a565b565b346108885761088461087061086b3660046101a6565b6107d3565b61087b939193610172565b93849384610823565b0390f35b6101f2565b346108bb576108a56108a03660046101a6565b6119c6565b6108ad610172565b806108b7816101c4565b0390f35b6101f2565b6108d4680100000000000000035f906105e3565b90565b34610907576108e736600461046e565b6109036108f26108c0565b6108fa610172565b91829182610547565b0390f35b6101f2565b3461093b5761092561091f3660046106ec565b90611bef565b61092d610172565b80610937816101c4565b0390f35b6101f2565b7f000000000000000000000000000000000000000000000000000000000000000090565b346109945761097436600461046e565b61099061097f610940565b610987610172565b918291826106a2565b0390f35b6101f2565b346109c9576109a936600461046e565b6109c56109b4611c50565b6109bc610172565b91829182610547565b0390f35b6101f2565b151590565b6109dc816109ce565b036109e357565b5f80fd5b905035906109f4826109d3565b565b6109ff81610689565b03610a0657565b5f80fd5b90503590610a17826109f6565b565b9190604083820312610a415780610a35610a3e925f86016109e7565b93602001610a0a565b90565b610178565b610a4f906109ce565b9052565b9190610a66905f60208501940190610a46565b565b34610a9957610a95610a84610a7e366004610a19565b90612043565b610a8c610172565b91829182610a53565b0390f35b6101f2565b90602082820312610ab757610ab4915f01610a0a565b90565b610178565b34610aea57610ad4610acf366004610a9e565b6120bd565b610adc610172565b80610ae6816101c4565b0390f35b6101f2565b610b03680100000000000000025f906105e3565b90565b34610b3657610b1636600461046e565b610b32610b21610aef565b610b29610172565b91829182610547565b0390f35b6101f2565b90602082820312610b6c575f82013567ffffffffffffffff8111610b6757610b639201610202565b9091565b61017c565b610178565b34610ba057610b8a610b84366004610b3b565b906120c8565b610b92610172565b80610b9c816101c4565b0390f35b6101f2565b5f80fd5b60209181520190565b60407f6d6574686f642e00000000000000000000000000000000000000000000000000917f4f6e6c792074686520436f736d69635369676e617475726547616d6520636f6e5f8201527f7472616374206973207065726d697474656420746f2063616c6c20746869732060208201520152565b610c326047606092610ba9565b610c3b81610bb2565b0190565b9190610c62906020610c5a604086018681035f880152610c25565b940190610695565b565b15610c6c5750565b610c8e90610c78610172565b91829163ced50f6760e01b835260048301610c3f565b0390fd5b610ce790610ce2610ca1612128565b610cd3610ccd7f0000000000000000000000000000000000000000000000000000000000000000610689565b91610689565b14610cdc612128565b90610c64565b611155565b565b90565b610d00610cfb610d0592610ce9565b6104ac565b610180565b90565b5f7f546865726520617265206e6f207374616b6564204e4654732e00000000000000910152565b610d3c6019602092610ba9565b610d4581610d08565b0190565b610d5e9060208101905f818303910152610d2f565b90565b90565b610d78610d73610d7d92610d61565b6104ac565b610180565b90565b634e487b7160e01b5f52601160045260245ffd5b610da3610da991939293610180565b92610180565b8201809211610db457565b610d80565b5f1b90565b90610dca5f1991610db9565b9181191691161790565b610de8610de3610ded92610180565b6104ac565b610180565b90565b90565b90610e08610e03610e0f92610dd4565b610df0565b8254610dbe565b9055565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610e3b90610e13565b810190811067ffffffffffffffff821117610e5557604052565b610e1d565b90610e6d610e66610172565b9283610e31565b565b610e796040610e5a565b90565b5f90565b5f90565b610e8c610e6f565b9060208083610e99610e7c565b815201610ea4610e80565b81525050565b610eb2610e84565b90565b90610ebf906103af565b9052565b90610ecd906103c9565b9052565b90610f07610efe5f610ee1610e6f565b94610ef8610ef083830161032d565b838801610eb5565b0161035f565b60208401610ec3565b565b610f1290610ed1565b90565b634e487b7160e01b5f52601260045260245ffd5b610f35610f3b91610180565b91610180565b908115610f46570490565b610f15565b610f5f610f5a610f6492610180565b6104ac565b6103c9565b90565b610f7190516103c9565b90565b610f80610f86916103c9565b916103c9565b019060018060c01b038211610f9757565b610d80565b610fa590610180565b5f198114610fb35760010190565b610d80565b610fcc610fc7610fd192610180565b6104ac565b6103af565b90565b634e487b7160e01b5f525f60045260245ffd5b610ff190516103af565b90565b9061100767ffffffffffffffff91610db9565b9181191691161790565b61102561102061102a926103af565b6104ac565b6103af565b90565b90565b9061104561104061104c92611011565b61102d565b8254610ff4565b9055565b60401b90565b9061106a67ffffffffffffffff1991611050565b9181191691161790565b61108861108361108d926103c9565b6104ac565b6103c9565b90565b90565b906110a86110a36110af92611074565b611090565b8254611056565b9055565b906110dd60205f6110e3946110d58282016110cf848801610fe7565b90611030565b019201610f67565b90611093565b565b91906110f6576110f4916110b3565b565b610fd4565b61110f61110a611114926103af565b6104ac565b610180565b90565b61114c61115394611142606094989795611138608086019a5f87019061053a565b602085019061053a565b604083019061053a565b019061053a565b565b61115f600161079a565b908161117361116d5f610cec565b91610180565b1461133a5761119d61118d6801000000000000000261079a565b6111976001610d64565b90610d94565b6111b08168010000000000000002610df3565b6111c26804000000000000000561079a565b6111ca610eaa565b906111dd6804000000000000000461079a565b6111ef6111e95f610cec565b91610180565b14155f146112de575f61121f61128e9261121a61120b84610cec565b68040000000000000004610df3565b610f9c565b946112338668040000000000000005610df3565b61124761123f86610fb8565b838601610eb5565b61126661125d611258348a90610f29565b610f4b565b60208601610ec3565b5b6112858461127f6804000000000000000689906102df565b906110e5565b93949201610fe7565b926112d934956112c76112c17fb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c091396610dd4565b966110fb565b966112d0610172565b94859485611117565b0390a3565b92905061128e5f6113026112fc6804000000000000000687906102df565b50610f09565b92611335611319611314348a90610f29565b610f4b565b61132f602087019161132a83610f67565b610f74565b90610ec3565b611267565b611342610172565b630f13059960e41b81528061135960048201610d49565b0390fd5b61136690610c92565b565b90565b61137f61137a61138492611368565b6104ac565b610180565b90565b61139d5f1961139761010061136b565b90610f29565b90565b60207f5f206973206f7574206f662074686520616c6c6f7765642072616e67652e0000917f6e756d4574684465706f73697473546f4576616c756174654d61784c696d69745f8201520152565b6113fa603e604092610ba9565b611403816113a0565b0190565b919061142a906020611422604086018681035f8801526113ed565b94019061053a565b565b156114345750565b61145690611440610172565b91829163285cf1f560e01b835260048301611407565b0390fd5b5090565b61146d61147391939293610180565b92610180565b820391821161147e57565b610d80565b600161148f9101610180565b90565b5f90565b91908110156114a6576020020190565b6102ba565b356114b581610183565b90565b92916114f5906114e4816114db6114d56114d0611387565b610180565b91610180565b1115829061142c565b6114ef85849061145a565b9061145e565b906114ff5f610cec565b6115085f610cec565b925b8361152761152161151c89879061145a565b610180565b91610180565b10156115745761156761156161153f61156d93610f9c565b611547611492565b5061155c6115578a888a91611496565b6114ab565b61241c565b93610d94565b93611483565b929061150a565b509150506115839192506126f4565b565b92916115c2906115b1816115a86115a261159d611387565b610180565b91610180565b1115829061142c565b6115bc85849061145a565b9061145e565b906115cc5f610cec565b6115d55f610cec565b925b836115f46115ee6115e989879061145a565b610180565b91610180565b10156116415761163461162e61160c61163a93610f9c565b611614611492565b506116296116248a888a91611496565b6114ab565b61288a565b93610d94565b93611483565b92906115d7565b509150506116509192506126f4565b565b506801000000000000000090565b90565b61166c81611652565b8210156116865761167e600191611660565b910201905f90565b6102ba565b6116a26116a89161169a611492565b506002611663565b906105e3565b90565b6116b3612bf5565b6116bb6116e5565b565b6116d16116cc6116d692610ce9565b6104ac565b6104a1565b90565b6116e2906116bd565b90565b6116f66116f15f6116d9565b612c43565b565b6117006116ab565b565b5f90565b61170e611702565b506117185f6107c6565b90565b6117539161174d91816117366117305f610cec565b91610180565b1180611755575b61174890839061142c565b61288a565b506126f4565b565b506117488261177361176d611768611387565b610180565b91610180565b1115905061173d565b60407f746f206265207374616b6564206f6e6c79206f6e63652e000000000000000000917f54686973204e46542068617320616c7265616479206265656e207374616b65645f8201527f20696e2074686520706173742e20416e204e465420697320616c6c6f7765642060208201520152565b6117fc6057606092610ba9565b6118058161177c565b0190565b919061182c906020611824604086018681035f8801526117ef565b94019061053a565b565b156118365750565b61185890611842610172565b9182916315fdfd0160e31b835260048301611809565b0390fd5b1b90565b9190600861187b9102916118755f198461185c565b9261185c565b9181191691161790565b919061189b6118966118a393610dd4565b610df0565b908354611860565b9055565b90565b906118bb60018060a01b0391610db9565b9181191691161790565b6118ce906104cb565b90565b90565b906118e96118e46118f0926118c5565b6118d1565b82546118aa565b9055565b634e487b7160e01b5f52602160045260245ffd5b6003111561191257565b6118f4565b9061192182611908565b565b61192c90611917565b90565b61193890611923565b9052565b91602061195d92949361195660408201965f83019061192f565b019061053a565b565b611968906104cb565b90565b5f80fd5b60e01b90565b5f91031261197f57565b610178565b6040906119ad6119b494969593966119a360608401985f850190610695565b6020830190610695565b019061053a565b565b6119be610172565b3d5f823e3d90fd5b6119f86119de6119d860028490611663565b906105e3565b6119f06119ea5f610cec565b91610180565b14829061182e565b611a17611a056001610d64565b611a1160028490611663565b90611885565b611a3c611a2c6801000000000000000261079a565b611a366001610d64565b90610d94565b611a4f8168010000000000000002610df3565b611a8c611a6f611a6968010000000000000004849061075e565b506118a7565b611a7b845f8301610df3565b6001611a85612128565b91016118d4565b611aa9611a99600161079a565b611aa36001610d64565b90610d94565b90611ab5826001610df3565b611ad1611ac26001610d64565b68040000000000000004610df3565b906001918390611adf612128565b92611b1c611b16611b107fcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f0082994610dd4565b94610dd4565b946118c5565b94611b31611b28610172565b9283928361193c565b0390a4611b5d7f00000000000000000000000000000000000000000000000000000000000000006104d7565b6323b872dd90611b6b612128565b90611b753061195f565b9392813b15611bea575f611b9c91611ba78296611b90610172565b9889978896879561196f565b855260048501611984565b03925af18015611be557611bb9575b50565b611bd8905f3d8111611bde575b611bd08183610e31565b810190611975565b5f611bb6565b503d611bc6565b6119b6565b61196b565b611c2791611c219181611c0a611c045f610cec565b91610180565b1180611c29575b611c1c90839061142c565b61241c565b506126f4565b565b50611c1c82611c47611c41611c3c611387565b610180565b91610180565b11159050611c11565b611c58611492565b50611c63600161079a565b90565b5f90565b90611c7d9291611c78612bf5565b611e74565b90565b5f7f546865726520617265207374696c6c207374616b6564204e4654732e00000000910152565b611cb4601c602092610ba9565b611cbd81611c80565b0190565b611cd69060208101905f818303910152611ca7565b90565b15611ce057565b611ce8610172565b63a29f5c4d60e01b815280611cff60048201611cc1565b0390fd5b5f7f546865726520617265207374696c6c20756e7061696420726577617264732e00910152565b611d37601f602092610ba9565b611d4081611d03565b0190565b611d599060208101905f818303910152611d2a565b90565b15611d6357565b611d6b610172565b63a29f5c4d60e01b815280611d8260048201611d44565b0390fd5b905090565b611d965f8092611d86565b0190565b611da390611d8b565b90565b67ffffffffffffffff8111611dc457611dc0602091610e13565b0190565b610e1d565b90611ddb611dd683611da6565b610e5a565b918252565b606090565b3d5f14611e0057611df53d611dc9565b903d5f602084013e5b565b611e08611de0565b90611dfe565b5f7f455448207472616e7366657220746f2063686172697479206661696c65642e00910152565b611e42601f602092610ba9565b611e4b81611e0e565b0190565b9190611e72906020611e6a604086018681035f880152611e35565b94019061053a565b565b50611e9a611e82600161079a565b611e94611e8e5f610cec565b91610180565b14611cd9565b611ec7611eaf6801000000000000000361079a565b611ec1611ebb5f610cec565b91610180565b14611d5c565b611fbd575b60019080611eea611ee4611edf5f6116d9565b610689565b91610689565b03611ef4575b5090565b611efd3061195f565b31905f808284611f0b610172565b9081611f1681611d9a565b03925af1611f22611de5565b505f14611f7257611f68611f567f1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d926118c5565b92611f5f610172565b91829182610547565b0390a25b5f611ef0565b909150611fb4611fa27f154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a926118c5565b92611fab610172565b91829182611e4f565b0390a25f611f6c565b611fd8611fc95f610cec565b68040000000000000005610df3565b611ff2611fed6805000000000000000661079a565b610f9c565b6120058168050000000000000006610df3565b61203b7f8a77a30ae54ca88ceda5f17cd3679c1b868abd2c6d36b67bcbe5e9beac7ec64891612032610172565b91829182610547565b0390a1611ecc565b9061205591612050611c66565b611c6a565b90565b61206990612064612bf5565b61206b565b565b8061208661208061207b5f6116d9565b610689565b91610689565b146120965761209490612c43565b565b6120b96120a25f6116d9565b5f918291631e4fbdf760e01b8352600483016106a2565b0390fd5b6120c690612058565b565b9190916120d45f610cec565b5b806120f26120ec6120e785889061145a565b610180565b91610180565b10156121225761211d9061211861211361210e85888591611496565b6114ab565b6119c6565b611483565b6120d5565b50509050565b612130611702565b503390565b9061213f90610180565b9052565b9061214d90610689565b9052565b61215b6060610e5a565b90565b906121ad6121a4600261216f612151565b9461218661217e5f830161079a565b5f8801612135565b61219e612195600183016107c6565b60208801612143565b0161079a565b60408401612135565b565b6121b89061215e565b90565b6121c59051610689565b90565b5f7f496e76616c6964204e4654207374616b6520616374696f6e2049442e00000000910152565b6121fc601c602092610ba9565b612205816121c8565b0190565b919061222c906020612224604086018681035f8801526121ef565b94019061053a565b565b60207f656365697665207374616b696e67207265776172642e00000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20725f8201520152565b6122886036604092610ba9565b6122918161222e565b0190565b9160406122c69294936122bf6122b4606083018381035f85015261227b565b96602083019061053a565b0190610695565b565b6122d29051610180565b90565b5f7f4e465420686173206e6f74206265656e20756e7374616b65642e000000000000910152565b612309601a602092610ba9565b612312816122d5565b0190565b9190612339906020612331604086018681035f8801526122fc565b94019061053a565b565b156123435750565b6123659061234f610172565b9182916372c1673f60e01b835260048301612316565b0390fd5b61237b91612375611492565b91611885565b565b9190600861239d91029161239760018060a01b038461185c565b9261185c565b9181191691161790565b91906123bd6123b86123c5936118c5565b6118d1565b90835461237d565b9055565b6123db916123d5611702565b916123a7565b565b6123e690610180565b5f81146123f4576001900390565b610d80565b91602061241a92949361241360408201965f83019061053a565b019061053a565b565b91612425611492565b5061242e611492565b5061244c61244668010000000000000004859061075e565b506118a7565b91612456836121af565b9061245f612128565b61247c612476612471602086016121bb565b610689565b91610689565b036125d757936124c86124d794956124b3612499604086016122c8565b6124ab6124a55f610cec565b91610180565b11849061233b565b82906124c1604086016122c8565b9091612d4e565b60408597929701909690612135565b946124f06124e7604085016122c8565b60028301610df3565b6124fc604084016122c8565b61250e6125085f610cec565b91610180565b1461258c575b506125205f83016122c8565b612528612128565b9161253660408895016122c8565b61257261256c6125667ff9bac74bc321a00ef2afeb1f44684e20e22f567699a5840df47967ea88c5b44994610dd4565b94610dd4565b946118c5565b9461258761257e610172565b928392836123f9565b0390a4565b5f60018261259f83806125a59601612369565b016123c9565b6125d16125c26125bd6801000000000000000361079a565b6123dd565b68010000000000000003610df3565b5f612514565b846125e4602084016121bb565b6125fe6125f86125f35f6116d9565b610689565b91610689565b14155f146126345761260e612128565b9061263061261a610172565b9283926348aca7ef60e11b845260048401612295565b0390fd5b61265690612640610172565b91829163023df6b160e21b835260048301612209565b0390fd5b60207f61696c65642e0000000000000000000000000000000000000000000000000000917f4e4654207374616b696e672045544820726577617264207061796d656e7420665f8201520152565b6126b46026604092610ba9565b6126bd8161265a565b0190565b9160406126f29294936126eb6126e0606083018381035f8501526126a7565b966020830190610695565b019061053a565b565b6127285f80612701612128565b8461270a610172565b908161271581611d9a565b03925af1612721611de5565b50156109ce565b61272f5750565b612737612128565b612758612742610172565b928392630aa7db6360e11b8452600484016126c1565b0390fd5b60207f6e7374616b652069742e00000000000000000000000000000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20755f8201520152565b6127b6602a604092610ba9565b6127bf8161275c565b0190565b9160406127f49294936127ed6127e2606083018381035f8501526127a9565b96602083019061053a565b0190610695565b565b5f7f4e46542068617320616c7265616479206265656e20756e7374616b65642e0000910152565b61282a601e602092610ba9565b612833816127f6565b0190565b919061285a906020612852604086018681035f88015261281d565b94019061053a565b565b156128645750565b61288690612870610172565b91829163c339165960e01b835260048301612837565b0390fd5b91612893611492565b5061289c611492565b506128ba6128b468010000000000000004859061075e565b506118a7565b916128c4836121af565b906128cd612128565b6128ea6128e46128df602086016121bb565b610689565b91610689565b03612b72579361293c61294b9495612921612907604086016122c8565b6129196129135f610cec565b91610180565b14849061285c565b82906129356804000000000000000561079a565b9091612d4e565b60408597929701909690612135565b94612958604084016122c8565b61296a6129645f610cec565b91610180565b115f14612b545761298a906002612983604086016122c8565b9101610df3565b6129b66129a76129a26801000000000000000361079a565b610f9c565b68010000000000000003610df3565b5b6129d46129c4600161079a565b6129ce6001610d64565b9061145e565b6129df816001610df3565b612a046129f46801000000000000000261079a565b6129fe6001610d64565b90610d94565b612a178168010000000000000002610df3565b91612a235f85016122c8565b90612a2c612128565b928894612a8c612a3e604089016122c8565b612a7a612a74612a6e7fec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf97610dd4565b97610dd4565b976118c5565b97612a83610172565b94859485611117565b0390a4612ab87f00000000000000000000000000000000000000000000000000000000000000006104d7565b6323b872dd90612ac73061195f565b90612adb5f612ad4612128565b95016122c8565b92813b15612b4f575f612b0191612b0c8296612af5610172565b9889978896879561196f565b855260048501611984565b03925af18015612b4a57612b1e575b50565b612b3d905f3d8111612b43575b612b358183610e31565b810190611975565b5f612b1b565b503d612b2b565b6119b6565b61196b565b5f600182612b678380612b6d9601612369565b016123c9565b6129b7565b84612b7f602084016121bb565b612b99612b93612b8e5f6116d9565b610689565b91610689565b14155f14612bcf57612ba9612128565b90612bcb612bb5610172565b9283926348aca7ef60e11b8452600484016127c3565b0390fd5b612bf190612bdb610172565b91829163023df6b160e21b835260048301612209565b0390fd5b612bfd611706565b612c16612c10612c0b612128565b610689565b91610689565b03612c1d57565b612c3f612c28612128565b5f91829163118cdaa760e01b8352600483016106a2565b0390fd5b612c4c5f6107c6565b612c56825f6118d4565b90612c8a612c847f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936118c5565b916118c5565b91612c93610172565b80612c9d816101c4565b0390a3565b90565b612cb9612cb4612cbe92610180565b6104ac565b612ca2565b90565b90612ccc9103612ca2565b90565b612ce3612cde612ce892612ca2565b6104ac565b610180565b90565b612cff612cfa612d0492610ce9565b6104ac565b612ca2565b90565b6001612d139103610180565b90565b612d2a612d25612d2f926103c9565b6104ac565b610180565b90565b90612d3d9101610180565b90565b90612d4b9103610180565b90565b92612d57611492565b92612d60611492565b50612d69611492565b50612d8d612d88612d7985612ca5565b612d8284612ca5565b90612cc1565b612ccf565b94612d9786612ca5565b612db1612dab612da65f612ceb565b612ca2565b91612ca2565b12612e7c575b9483955b86612dce612dc884610180565b91610180565b115f14612e6b57612dfb612df5612def680400000000000000068a906102df565b50610f09565b97612d07565b96612e075f8201610fe7565b612e19612e1384610180565b916110fb565b11612e4457505050612e3b612e4192939495612e345f610cec565b955b612d40565b90612d40565b90565b95612e5e612e596020612e6494959901610f67565b612d16565b90612d32565b9490612dbb565b509394612e3b90612e419394612e36565b9450612e875f610cec565b94612db756fea2646970667358221220f1ba25035a60cc47b52fa71e93d11254307c000ede470b3739af61eec57b3dc464736f6c634300081c0033",
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

// NumStateResets is a free data retrieval call binding the contract method 0x7c843a46.
//
// Solidity: function numStateResets() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) NumStateResets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "numStateResets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStateResets is a free data retrieval call binding the contract method 0x7c843a46.
//
// Solidity: function numStateResets() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) NumStateResets() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.NumStateResets(&_StakingWalletCosmicSignatureNft.CallOpts)
}

// NumStateResets is a free data retrieval call binding the contract method 0x7c843a46.
//
// Solidity: function numStateResets() view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) NumStateResets() (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.NumStateResets(&_StakingWalletCosmicSignatureNft.CallOpts)
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
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletCosmicSignatureNft.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
	return _StakingWalletCosmicSignatureNft.Contract.WasNftUsed(&_StakingWalletCosmicSignatureNft.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftCallerSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
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
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 indexed depositId, uint256 depositAmount, uint256 numStakedNfts)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) FilterEthDepositReceived(opts *bind.FilterOpts, roundNum []*big.Int, depositId []*big.Int) (*StakingWalletCosmicSignatureNftEthDepositReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "EthDepositReceived", roundNumRule, depositIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftEthDepositReceivedIterator{contract: _StakingWalletCosmicSignatureNft.contract, event: "EthDepositReceived", logs: logs, sub: sub}, nil
}

// WatchEthDepositReceived is a free log subscription operation binding the contract event 0xb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c0913.
//
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 indexed depositId, uint256 depositAmount, uint256 numStakedNfts)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) WatchEthDepositReceived(opts *bind.WatchOpts, sink chan<- *StakingWalletCosmicSignatureNftEthDepositReceived, roundNum []*big.Int, depositId []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "EthDepositReceived", roundNumRule, depositIdRule)
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
// Solidity: event EthDepositReceived(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 indexed depositId, uint256 depositAmount, uint256 numStakedNfts)
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
	ActionCounter            *big.Int
	StakeActionId            *big.Int
	NftId                    *big.Int
	StakerAddress            common.Address
	NumStakedNfts            *big.Int
	RewardAmount             *big.Int
	MaxUnpaidEthDepositIndex *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNftUnstaked is a free log retrieval operation binding the contract event 0xec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
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

// WatchNftUnstaked is a free log subscription operation binding the contract event 0xec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
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

// ParseNftUnstaked is a log parse operation binding the contract event 0xec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts, uint256 rewardAmount, uint256 maxUnpaidEthDepositIndex)
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

// StakingWalletCosmicSignatureNftStateResetIterator is returned from FilterStateReset and is used to iterate over the raw logs and unpacked data for StateReset events raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftStateResetIterator struct {
	Event *StakingWalletCosmicSignatureNftStateReset // Event containing the contract specifics and raw log

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
func (it *StakingWalletCosmicSignatureNftStateResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCosmicSignatureNftStateReset)
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
		it.Event = new(StakingWalletCosmicSignatureNftStateReset)
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
func (it *StakingWalletCosmicSignatureNftStateResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCosmicSignatureNftStateResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCosmicSignatureNftStateReset represents a StateReset event raised by the StakingWalletCosmicSignatureNft contract.
type StakingWalletCosmicSignatureNftStateReset struct {
	NumStateResets *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStateReset is a free log retrieval operation binding the contract event 0x8a77a30ae54ca88ceda5f17cd3679c1b868abd2c6d36b67bcbe5e9beac7ec648.
//
// Solidity: event StateReset(uint256 numStateResets)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) FilterStateReset(opts *bind.FilterOpts) (*StakingWalletCosmicSignatureNftStateResetIterator, error) {

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.FilterLogs(opts, "StateReset")
	if err != nil {
		return nil, err
	}
	return &StakingWalletCosmicSignatureNftStateResetIterator{contract: _StakingWalletCosmicSignatureNft.contract, event: "StateReset", logs: logs, sub: sub}, nil
}

// WatchStateReset is a free log subscription operation binding the contract event 0x8a77a30ae54ca88ceda5f17cd3679c1b868abd2c6d36b67bcbe5e9beac7ec648.
//
// Solidity: event StateReset(uint256 numStateResets)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) WatchStateReset(opts *bind.WatchOpts, sink chan<- *StakingWalletCosmicSignatureNftStateReset) (event.Subscription, error) {

	logs, sub, err := _StakingWalletCosmicSignatureNft.contract.WatchLogs(opts, "StateReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCosmicSignatureNftStateReset)
				if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "StateReset", log); err != nil {
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

// ParseStateReset is a log parse operation binding the contract event 0x8a77a30ae54ca88ceda5f17cd3679c1b868abd2c6d36b67bcbe5e9beac7ec648.
//
// Solidity: event StateReset(uint256 numStateResets)
func (_StakingWalletCosmicSignatureNft *StakingWalletCosmicSignatureNftFilterer) ParseStateReset(log types.Log) (*StakingWalletCosmicSignatureNftStateReset, error) {
	event := new(StakingWalletCosmicSignatureNftStateReset)
	if err := _StakingWalletCosmicSignatureNft.contract.UnpackLog(event, "StateReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletNftBaseMetaData contains all meta data concerning the StakingWalletNftBase contract.
var StakingWalletNftBaseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIStakingWalletNftBase.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_StakingWalletNftBase *StakingWalletNftBaseCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletNftBase.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_StakingWalletNftBase *StakingWalletNftBaseSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
	return _StakingWalletNftBase.Contract.WasNftUsed(&_StakingWalletNftBase.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_StakingWalletNftBase *StakingWalletNftBaseCallerSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
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
