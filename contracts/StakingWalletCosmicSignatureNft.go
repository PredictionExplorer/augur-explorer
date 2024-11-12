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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"EthDepositReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicGameConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStateResets\",\"type\":\"uint256\"}],\"name\":\"StateReset\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetState_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"tryPerformMaintenance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"InvalidOperationInCurrentState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftAlreadyUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftNotUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"NftOneTimeStaking\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"NftStakeActionAccessDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftStakeActionInvalidId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NoStakedNfts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit\",\"type\":\"uint256\"}],\"name\":\"NumEthDepositsToEvaluateMaxLimitIsOutOfAllowedRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"EthDepositReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferredToCharity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicGameConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStateResets\",\"type\":\"uint256\"}],\"name\":\"StateReset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_ETH_DEPOSITS_TO_EVALUATE_HARD_MAX_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"actionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ethDeposits\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"depositId\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"rewardAmountPerStakedNft\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numEthDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStateResets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numUnpaidStakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetState_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"tryPerformMaintenance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052346100305761001a610014610133565b9061039c565b610022610035565b6130ca61054a82396130ca90f35b61003b565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100679061003f565b810190811060018060401b0382111761007f57604052565b610049565b90610097610090610035565b928361005d565b565b5f80fd5b60018060a01b031690565b6100b19061009d565b90565b6100bd906100a8565b90565b6100c9816100b4565b036100d057565b5f80fd5b905051906100e1826100c0565b565b6100ec816100a8565b036100f357565b5f80fd5b90505190610104826100e3565b565b919060408382031261012e578061012261012b925f86016100d4565b936020016100f7565b90565b610099565b6101516136148038038061014681610084565b928339810190610106565b9091565b90565b61016c6101676101719261009d565b610155565b61009d565b90565b61017d90610158565b90565b61018990610174565b90565b90565b6101a361019e6101a89261018c565b610155565b61009d565b90565b6101b49061018f565b90565b60209181520190565b60207f66745f2e00000000000000000000000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f7220746865206e5f8201520152565b61021a60246040926101b7565b610223816101c0565b0190565b61023c9060208101905f81830391015261020d565b90565b1561024657565b61024e610035565b63eac0d38960e01b81528061026560048201610227565b0390fd5b60207f616d655f2e000000000000000000000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f722074686520675f8201520152565b6102c360256040926101b7565b6102cc81610269565b0190565b6102e59060208101905f8183039101526102b6565b90565b156102ef57565b6102f7610035565b63eac0d38960e01b81528061030e600482016102d0565b0390fd5b5f1b90565b9061032860018060a01b0391610312565b9181191691161790565b61033b90610158565b90565b61034790610332565b90565b90565b9061036261035d6103699261033e565b61034a565b8254610317565b9055565b61037690610174565b90565b90565b9061039161038c6103989261036d565b610379565b8254610317565b9055565b9061040e61041d926103ad3361041f565b6103da6103b982610180565b6103d36103cd6103c85f6101ab565b6100a8565b916100a8565b141561023f565b6103ff836103f86103f26103ed5f6101ab565b6100a8565b916100a8565b14156102e8565b6801000000000000000361034d565b6801000000000000000461037c565b565b6104289061042a565b565b61043390610435565b565b61043e90610462565b565b610449906100a8565b9052565b9190610460905f60208501940190610440565b565b8061047d6104776104725f6101ab565b6100a8565b916100a8565b1461048d5761048b906104ea565b565b6104b06104995f6101ab565b5f918291631e4fbdf760e01b83526004830161044d565b0390fd5b5f1c90565b60018060a01b031690565b6104d06104d5916104b4565b6104b9565b90565b6104e290546104c4565b90565b5f0190565b6104f35f6104d8565b6104fd825f61037c565b9061053161052b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09361036d565b9161036d565b9161053a610035565b80610544816104e5565b0390a356fe60806040526004361015610013575b610cb4565b61001d5f3561017c565b806315b4e68f146101775780632e53a259146101725780632ffb33041461016d578063400c89561461016857806347ccca021461016357806357951c741461015e578063715018a6146101595780637a025499146101545780637c843a461461014f5780638da5cb5b1461014a5780639e2c8a5b14610145578063a2b136fb14610140578063a694fc3a1461013b578063a6abc99f14610136578063c2ac898f14610131578063c3fe3e281461012c578063ca7c1f9214610127578063d8ee557314610122578063e0c10d6a1461011d578063f2fde38b14610118578063fdbd98b0146101135763fe939afc0361000e57610c80565b610c15565b610bcb565b610b78565b610ab5565b610a0d565b6109d8565b610967565b610932565b6108e8565b6108b0565b610774565b610712565b6106af565b610663565b6105ce565b610599565b61053d565b61044a565b610414565b610296565b6101d9565b60e01c90565b60405190565b5f80fd5b5f80fd5b90565b61019c81610190565b036101a357565b5f80fd5b905035906101b482610193565b565b906020828203126101cf576101cc915f016101a7565b90565b610188565b5f0190565b6101ec6101e73660046101b6565b611416565b6101f4610182565b806101fe816101d4565b0390f35b5f80fd5b5f80fd5b5f80fd5b5f80fd5b909182601f8301121561024c5781359167ffffffffffffffff831161024757602001926020830284011161024257565b61020e565b61020a565b610206565b91604083830312610291575f83013567ffffffffffffffff811161028c5761027e83610289928601610212565b9390946020016101a7565b90565b61018c565b610188565b346102c5576102af6102a9366004610251565b91611539565b6102b7610182565b806102c1816101d4565b0390f35b610202565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b6102f8816102de565b8210156103125761030a6001916102ec565b910201905f90565b6102ca565b5f1c90565b67ffffffffffffffff1690565b61033561033a91610317565b61031c565b90565b6103479054610329565b90565b60401c90565b60018060c01b031690565b61036761036c9161034a565b610350565b90565b610379905461035b565b90565b6804000000000000000790610390826102de565b8110156103bb576103a0916102ef565b50906103b85f6103b181850161033d565b930161036f565b90565b5f80fd5b67ffffffffffffffff1690565b6103d5906103bf565b9052565b60018060c01b031690565b6103ed906103d9565b9052565b91602061041292949361040b60408201965f8301906103cc565b01906103e4565b565b346104455761042c6104273660046101b6565b61037c565b90610441610438610182565b928392836103f1565b0390f35b610202565b346104795761046361045d366004610251565b91611606565b61046b610182565b80610475816101d4565b0390f35b610202565b5f91031261048857565b610188565b1c90565b60018060a01b031690565b6104ac9060086104b1930261048d565b610491565b90565b906104bf915461049c565b90565b6104d6680100000000000000035f906104b4565b90565b60018060a01b031690565b90565b6104fb6104f6610500926104d9565b6104e4565b6104d9565b90565b61050c906104e7565b90565b61051890610503565b90565b6105249061050f565b9052565b919061053b905f6020850194019061051b565b565b3461056d5761054d36600461047e565b6105696105586104c2565b610560610182565b91829182610528565b0390f35b610202565b151590565b61058090610572565b9052565b9190610597905f60208501940190610577565b565b346105c9576105c56105b46105af3660046101b6565b611737565b6105bc610182565b91829182610584565b0390f35b610202565b346105fc576105de36600461047e565b6105e66117a6565b6105ee610182565b806105f8816101d4565b0390f35b610202565b90565b610614906008610619930261048d565b610601565b90565b906106279154610604565b90565b61063e680500000000000000075f9061061c565b90565b61064a90610190565b9052565b9190610661905f60208501940190610641565b565b346106935761067336600461047e565b61068f61067e61062a565b610686610182565b9182918261064e565b0390f35b610202565b6106ac680500000000000000085f9061061c565b90565b346106df576106bf36600461047e565b6106db6106ca610698565b6106d2610182565b9182918261064e565b0390f35b610202565b6106ed906104d9565b90565b6106f9906106e4565b9052565b9190610710905f602085019401906106f0565b565b346107425761072236600461047e565b61073e61072d6117b4565b610735610182565b918291826106fd565b0390f35b610202565b919060408382031261076f578061076361076c925f86016101a7565b936020016101a7565b90565b610188565b346107a35761078d610787366004610747565b906117c9565b610795610182565b8061079f816101d4565b0390f35b610202565b506801000000000000000090565b90565b6107c2816107a8565b8210156107dc576107d46003916107b6565b910201905f90565b6102ca565b6107ed6107f291610317565b610601565b90565b6107ff90546107e1565b90565b60018060a01b031690565b61081961081e91610317565b610802565b90565b61082b905461080d565b90565b6801000000000000000590610842826107a8565b81101561087a57610852916107b9565b5061085e5f82016107f5565b91610877600261087060018501610821565b93016107f5565b90565b5f80fd5b6040906108a76108ae949695939661089d60608401985f850190610641565b60208301906106f0565b0190610641565b565b346108e3576108df6108cb6108c63660046101b6565b61082e565b6108d6939193610182565b9384938461087e565b0390f35b610202565b34610916576109006108fb3660046101b6565b611b9d565b610908610182565b80610912816101d4565b0390f35b610202565b61092f680400000000000000055f9061061c565b90565b346109625761094236600461047e565b61095e61094d61091b565b610955610182565b9182918261064e565b0390f35b610202565b346109965761098061097a366004610747565b90611db8565b610988610182565b80610992816101d4565b0390f35b610202565b6109ab9060086109b0930261048d565b610802565b90565b906109be915461099b565b90565b6109d5680100000000000000045f906109b3565b90565b34610a08576109e836600461047e565b610a046109f36109c1565b6109fb610182565b918291826106fd565b0390f35b610202565b34610a3d57610a1d36600461047e565b610a39610a28611e19565b610a30610182565b9182918261064e565b0390f35b610202565b610a4b81610572565b03610a5257565b5f80fd5b90503590610a6382610a42565b565b610a6e816106e4565b03610a7557565b5f80fd5b90503590610a8682610a65565b565b9190604083820312610ab05780610aa4610aad925f8601610a56565b93602001610a79565b90565b610188565b34610ae657610ae2610ad1610acb366004610a88565b90612208565b610ad9610182565b91829182610584565b0390f35b610202565b90565b610b02610afd610b0792610aeb565b6104e4565b610190565b90565b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b610b3e610b4491610190565b91610190565b908115610b4f570490565b610b0a565b610b6a5f19610b64610100610aee565b90610b32565b90565b610b75610b54565b90565b34610ba857610b8836600461047e565b610ba4610b93610b6d565b610b9b610182565b9182918261064e565b0390f35b610202565b90602082820312610bc657610bc3915f01610a79565b90565b610188565b34610bf957610be3610bde366004610bad565b612282565b610beb610182565b80610bf5816101d4565b0390f35b610202565b610c12680100000000000000025f9061061c565b90565b34610c4557610c2536600461047e565b610c41610c30610bfe565b610c38610182565b9182918261064e565b0390f35b610202565b90602082820312610c7b575f82013567ffffffffffffffff8111610c7657610c729201610212565b9091565b61018c565b610188565b34610caf57610c99610c93366004610c4a565b9061228d565b610ca1610182565b80610cab816101d4565b0390f35b610202565b5f80fd5b60209181520190565b60207f7065726d697474656420746f2063616c6c2074686973206d6574686f642e0000917f4f6e6c792074686520436f736d696347616d6520636f6e7472616374206973205f8201520152565b610d1b603e604092610cb8565b610d2481610cc1565b0190565b9190610d4b906020610d43604086018681035f880152610d0e565b9401906106f0565b565b15610d555750565b610d7790610d61610182565b91829163336e7de360e11b835260048301610d28565b0390fd5b610db390610dae33610da6610da0610d9b68010000000000000004610821565b6106e4565b916106e4565b143390610d4d565b611212565b565b90565b610dcc610dc7610dd192610db5565b6104e4565b610190565b90565b5f7f546865726520617265206e6f207374616b6564204e4654732e00000000000000910152565b610e086019602092610cb8565b610e1181610dd4565b0190565b610e2a9060208101905f818303910152610dfb565b90565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610e5590610e2d565b810190811067ffffffffffffffff821117610e6f57604052565b610e37565b90610e87610e80610182565b9283610e4b565b565b610e936040610e74565b90565b5f90565b5f90565b610ea6610e89565b9060208083610eb3610e96565b815201610ebe610e9a565b81525050565b610ecc610e9e565b90565b90565b610ee6610ee1610eeb92610ecf565b6104e4565b610190565b90565b610efd610f0391939293610190565b92610190565b8201809211610f0e57565b610b1e565b5f1b90565b90610f245f1991610f13565b9181191691161790565b610f42610f3d610f4792610190565b6104e4565b610190565b90565b90565b90610f62610f5d610f6992610f2e565b610f4a565b8254610f18565b9055565b90565b610f84610f7f610f8992610f6d565b6104e4565b610190565b90565b90610f96906103bf565b9052565b90610fa4906103d9565b9052565b90610fde610fd55f610fb8610e89565b94610fcf610fc783830161033d565b838801610f8c565b0161036f565b60208401610f9a565b565b610fe990610fa8565b90565b611000610ffb61100592610190565b6104e4565b6103d9565b90565b61101290516103d9565b90565b611021611027916103d9565b916103d9565b019060018060c01b03821161103857565b610b1e565b61104690610190565b5f1981146110545760010190565b610b1e565b61106d61106861107292610190565b6104e4565b6103bf565b90565b634e487b7160e01b5f525f60045260245ffd5b61109290516103bf565b90565b906110a867ffffffffffffffff91610f13565b9181191691161790565b6110c66110c16110cb926103bf565b6104e4565b6103bf565b90565b90565b906110e66110e16110ed926110b2565b6110ce565b8254611095565b9055565b60401b90565b9061110b67ffffffffffffffff19916110f1565b9181191691161790565b61112961112461112e926103d9565b6104e4565b6103d9565b90565b90565b9061114961114461115092611115565b611131565b82546110f7565b9055565b9061117e60205f61118494611176828201611170848801611088565b906110d1565b019201611008565b90611134565b565b91906111975761119591611154565b565b611075565b6111b06111ab6111b5926103bf565b6104e4565b610190565b90565b6111c19061119c565b9052565b90959492611210946111ff611209926111f56080966111eb60a088019c5f890190610641565b6020870190610641565b60408501906111b8565b6060830190610641565b0190610641565b565b61121c60016107f5565b908161123061122a5f610db8565b91610190565b146113f35761123d610ec4565b61124f680500000000000000076107f5565b91611275611265680100000000000000026107f5565b61126f6001610ed2565b90610eee565b926112898468010000000000000002610f4d565b61129b680400000000000000066107f5565b6112ae6112a86002610f70565b91610190565b10155f14611394576113515f6112e261138f936112dd6112ce6001610ed2565b68040000000000000006610f4d565b61103d565b936112f68568050000000000000007610f4d565b61130a61130288611059565b838801610f8c565b61132961132061131b348b90610b32565b610fec565b60208801610f9a565b5b611348866113426804000000000000000788906102ef565b90611186565b95939401611088565b943461137d7fb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c091396610f2e565b96611386610182565b958695866111c5565b0390a2565b90915061138f6113515f6113bb6113b56804000000000000000786906102ef565b50610fe0565b946113ee6113d26113cd348b90610b32565b610fec565b6113e860208901916113e383611008565b611015565b90610f9a565b61132a565b6113fb610182565b630f13059960e41b81528061141260048201610e15565b0390fd5b61141f90610d7b565b565b60207f5f206973206f7574206f662074686520616c6c6f7765642072616e67652e0000917f6e756d4574684465706f73697473546f4576616c756174654d61784c696d69745f8201520152565b61147b603e604092610cb8565b61148481611421565b0190565b91906114ab9060206114a3604086018681035f88015261146e565b940190610641565b565b156114b55750565b6114d7906114c1610182565b91829163285cf1f560e01b835260048301611488565b0390fd5b5090565b6114ee6114f491939293610190565b92610190565b82039182116114ff57565b610b1e565b60016115109101610190565b90565b5f90565b9190811015611527576020020190565b6102ca565b3561153681610193565b90565b9291611576906115658161155c611556611551610b54565b610190565b91610190565b111582906114ad565b6115708584906114db565b906114df565b906115805f610db8565b6115895f610db8565b925b836115a86115a261159d8987906114db565b610190565b91610190565b10156115f5576115e86115e26115c06115ee9361103d565b6115c8611513565b506115dd6115d88a888a91611517565b61152c565b61261f565b93610eee565b93611504565b929061158b565b50915050611604919250612910565b565b9291611643906116328161162961162361161e610b54565b610190565b91610190565b111582906114ad565b61163d8584906114db565b906114df565b9061164d5f610db8565b6116565f610db8565b925b8361167561166f61166a8987906114db565b610190565b91610190565b10156116c2576116b56116af61168d6116bb9361103d565b611695611513565b506116aa6116a58a888a91611517565b61152c565b612aaf565b93610eee565b93611504565b9290611658565b509150506116d1919250612910565b565b5f90565b506801000000000000000090565b90565b6116f1816116d7565b82101561170b576117036001916116e5565b910201905f90565b6102ca565b60ff1690565b61172261172791610317565b611710565b90565b6117349054611716565b90565b5f61174f611756926117476116d3565b5060026116e8565b500161172a565b90565b611761612def565b611769611793565b565b61177f61177a61178492610db5565b6104e4565b6104d9565b90565b6117909061176b565b90565b6117a461179f5f611787565b612e3d565b565b6117ae611759565b565b5f90565b6117bc6117b0565b506117c65f610821565b90565b611801916117fb91816117e46117de5f610db8565b91610190565b1180611803575b6117f69083906114ad565b612aaf565b50612910565b565b506117f68261182161181b611816610b54565b610190565b91610190565b111590506117eb565b60407f206f6e6c79206f6e63652e000000000000000000000000000000000000000000917f54686973204e46542068617320616c7265616479206265656e207374616b65645f8201527f2e20416e204e465420697320616c6c6f77656420746f206265207374616b656460208201520152565b6118aa604b606092610cb8565b6118b38161182a565b0190565b91906118da9060206118d2604086018681035f88015261189d565b940190610641565b565b156118e45750565b611906906118f0610182565b918291633b471d1f60e21b8352600483016118b7565b0390fd5b90565b9061191e60018060a01b0391610f13565b9181191691161790565b61193190610503565b90565b90565b9061194c61194761195392611928565b611934565b825461190d565b9055565b6119616040610e74565b90565b9061196e90610572565b9052565b60018060f81b031690565b61199161198c61199692610db5565b6104e4565b611972565b90565b906119a390611972565b9052565b6119b19051610572565b90565b906119c060ff91610f13565b9181191691161790565b6119d390610572565b90565b90565b906119ee6119e96119f5926119ca565b6119d6565b82546119b4565b9055565b611a039051611972565b90565b60081b90565b90611a1960ff1991611a06565b9181191691161790565b611a37611a32611a3c92611972565b6104e4565b611972565b90565b90565b90611a57611a52611a5e92611a23565b611a3f565b8254611a0c565b9055565b90611a8c60205f611a9294611a84828201611a7e8488016119a7565b906119d9565b0192016119f9565b90611a42565b565b9190611aa557611aa391611a62565b565b611075565b611ab6611abb91610317565b610491565b90565b611ac89054611aaa565b90565b611ad490610503565b90565b5f80fd5b60e01b90565b5f910312611aeb57565b610188565b604090611b19611b209496959396611b0f60608401985f8501906106f0565b60208301906106f0565b0190610641565b565b611b2a610182565b3d5f823e3d90fd5b634e487b7160e01b5f52602160045260245ffd5b60031115611b5057565b611b32565b90611b5f82611b46565b565b611b6a90611b55565b90565b611b7690611b61565b9052565b916020611b9b929493611b9460408201965f830190611b6d565b0190610641565b565b611bc7611bc0611bba5f611bb3600286906116e8565b500161172a565b15610572565b82906118dc565b611bec611bdc680100000000000000026107f5565b611be66001610ed2565b90610eee565b611bff8168010000000000000002610f4d565b611c35611c1f611c196801000000000000000584906107b9565b5061190a565b611c2b845f8301610f4d565b6001339101611937565b611c52611c4260016107f5565b611c4c6001610ed2565b90610eee565b90611c5e826001610f4d565b611c7a611c6b6002610f70565b68040000000000000006610f4d565b611cba6001611ca8611c9f5f611c9a611c91611957565b945f8601611964565b61197d565b60208301611999565b611cb4600286906116e8565b90611a94565b611cd4611ccf68010000000000000003611abe565b61050f565b6323b872dd33611ce330611acb565b928692813b15611db3575f611d0b91611d168296611cff610182565b98899788968795611adb565b855260048501611af0565b03925af18015611dae57611d82575b50600192903392611d68611d62611d5c7fcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f0082994610f2e565b94610f2e565b94611928565b94611d7d611d74610182565b92839283611b7a565b0390a4565b611da1905f3d8111611da7575b611d998183610e4b565b810190611ae1565b5f611d25565b503d611d8f565b611b22565b611ad7565b611df091611dea9181611dd3611dcd5f610db8565b91610190565b1180611df2575b611de59083906114ad565b61261f565b50612910565b565b50611de582611e10611e0a611e05610b54565b610190565b91610190565b11159050611dda565b611e21611513565b50611e2c60016107f5565b90565b90611e429291611e3d612def565b612039565b90565b5f7f546865726520617265207374696c6c207374616b6564204e4654732e00000000910152565b611e79601c602092610cb8565b611e8281611e45565b0190565b611e9b9060208101905f818303910152611e6c565b90565b15611ea557565b611ead610182565b63a29f5c4d60e01b815280611ec460048201611e86565b0390fd5b5f7f546865726520617265207374696c6c20756e7061696420726577617264732e00910152565b611efc601f602092610cb8565b611f0581611ec8565b0190565b611f1e9060208101905f818303910152611eef565b90565b15611f2857565b611f30610182565b63a29f5c4d60e01b815280611f4760048201611f09565b0390fd5b905090565b611f5b5f8092611f4b565b0190565b611f6890611f50565b90565b67ffffffffffffffff8111611f8957611f85602091610e2d565b0190565b610e37565b90611fa0611f9b83611f6b565b610e74565b918252565b606090565b3d5f14611fc557611fba3d611f8e565b903d5f602084013e5b565b611fcd611fa5565b90611fc3565b5f7f5472616e7366657220746f2063686172697479206661696c65642e0000000000910152565b612007601b602092610cb8565b61201081611fd3565b0190565b919061203790602061202f604086018681035f880152611ffa565b940190610641565b565b5061205f61204760016107f5565b6120596120535f610db8565b91610190565b14611e9e565b61208c612074680400000000000000056107f5565b6120866120805f610db8565b91610190565b14611f21565b612182575b600190806120af6120a96120a45f611787565b6106e4565b916106e4565b036120b9575b5090565b6120c230611acb565b31905f8082846120d0610182565b90816120db81611f5f565b03925af16120e7611faa565b505f146121375761212d61211b7f1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d92611928565b92612124610182565b9182918261064e565b0390a25b5f6120b5565b9091506121796121677f154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a92611928565b92612170610182565b91829182612014565b0390a25f612131565b61219d61218e5f610db8565b68050000000000000007610f4d565b6121b76121b2680500000000000000086107f5565b61103d565b6121ca8168050000000000000008610f4d565b6122007f8a77a30ae54ca88ceda5f17cd3679c1b868abd2c6d36b67bcbe5e9beac7ec648916121f7610182565b9182918261064e565b0390a1612091565b9061221a916122156116d3565b611e2f565b90565b61222e90612229612def565b612230565b565b8061224b6122456122405f611787565b6106e4565b916106e4565b1461225b5761225990612e3d565b565b61227e6122675f611787565b5f918291631e4fbdf760e01b8352600483016106fd565b0390fd5b61228b9061221d565b565b9190916122995f610db8565b5b806122b76122b16122ac8588906114db565b610190565b91610190565b10156122e7576122e2906122dd6122d86122d385888591611517565b61152c565b611b9d565b611504565b61229a565b50509050565b906122f790610190565b9052565b90612305906106e4565b9052565b6123136060610e74565b90565b9061236561235c6002612327612309565b9461233e6123365f83016107f5565b5f88016122ed565b61235661234d60018301610821565b602088016122fb565b016107f5565b604084016122ed565b565b61237090612316565b90565b61237d90516106e4565b90565b5f7f496e76616c6964204e4654207374616b6520616374696f6e2049442e00000000910152565b6123b4601c602092610cb8565b6123bd81612380565b0190565b91906123e49060206123dc604086018681035f8801526123a7565b940190610641565b565b60207f656365697665207374616b696e67207265776172642e00000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20725f8201520152565b6124406036604092610cb8565b612449816123e6565b0190565b91604061247e92949361247761246c606083018381035f850152612433565b966020830190610641565b01906106f0565b565b61248a9051610190565b90565b5f7f4e465420686173206e6f74206265656e20756e7374616b65642e000000000000910152565b6124c1601a602092610cb8565b6124ca8161248d565b0190565b91906124f19060206124e9604086018681035f8801526124b4565b940190610641565b565b156124fb5750565b61251d90612507610182565b9182916372c1673f60e01b8352600483016124ce565b0390fd5b1b90565b9190600861254091029161253a5f1984612521565b92612521565b9181191691161790565b919061256061255b61256893610f2e565b610f4a565b908354612525565b9055565b61257e91612578611513565b9161254a565b565b919060086125a091029161259a60018060a01b0384612521565b92612521565b9181191691161790565b91906125c06125bb6125c893611928565b611934565b908354612580565b9055565b6125de916125d86117b0565b916125aa565b565b6125e990610190565b5f81146125f7576001900390565b610b1e565b91602061261d92949361261660408201965f830190610641565b0190610641565b565b91612628611513565b50612631611513565b5061264f6126496801000000000000000585906107b9565b5061190a565b9161265983612367565b903361267861267261266d60208601612373565b6106e4565b916106e4565b036127cc57936126c46126d394956126af61269560408601612480565b6126a76126a15f610db8565b91610190565b1184906124f3565b82906126bd60408601612480565b9091612f48565b604085979297019096906122ed565b946126ec6126e360408501612480565b60028301610f4d565b6126f860408401612480565b61270a6127045f610db8565b91610190565b14612781575b5061271c5f8301612480565b339161272b6040889501612480565b61276761276161275b7ff9bac74bc321a00ef2afeb1f44684e20e22f567699a5840df47967ea88c5b44994610f2e565b94610f2e565b94611928565b9461277c612773610182565b928392836125fc565b0390a4565b5f600182612794838061279a960161256c565b016125cc565b6127c66127b76127b2680400000000000000056107f5565b6125e0565b68040000000000000005610f4d565b5f612710565b846127d960208401612373565b6127f36127ed6127e85f611787565b6106e4565b916106e4565b14155f1461282257339061281e612808610182565b9283926348aca7ef60e11b84526004840161244d565b0390fd5b6128449061282e610182565b91829163023df6b160e21b8352600483016123c1565b0390fd5b60207f642e000000000000000000000000000000000000000000000000000000000000917f4e4654207374616b696e6720726577617264207061796d656e74206661696c655f8201520152565b6128a26022604092610cb8565b6128ab81612848565b0190565b9160406128e09294936128d96128ce606083018381035f850152612895565b9660208301906106f0565b0190610641565b565b156128eb575050565b61290c6128f6610182565b928392630aa7db6360e11b8452600484016128af565b0390fd5b612941905f803383612920610182565b908161292b81611f5f565b03925af1612937611faa565b50903390916128e2565b565b60207f6e7374616b652069742e00000000000000000000000000000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20755f8201520152565b61299d602a604092610cb8565b6129a681612943565b0190565b9160406129db9294936129d46129c9606083018381035f850152612990565b966020830190610641565b01906106f0565b565b5f7f4e46542068617320616c7265616479206265656e20756e7374616b65642e0000910152565b612a11601e602092610cb8565b612a1a816129dd565b0190565b9190612a41906020612a39604086018681035f880152612a04565b940190610641565b565b15612a4b5750565b612a6d90612a57610182565b91829163c339165960e01b835260048301612a1e565b0390fd5b612aa6612aad94612a9c606094989795612a92608086019a5f870190610641565b6020850190610641565b6040830190610641565b0190610641565b565b91612ab8611513565b50612ac1611513565b50612adf612ad96801000000000000000585906107b9565b5061190a565b91612ae983612367565b9033612b08612b02612afd60208601612373565b6106e4565b916106e4565b03612d735793612b5a612b699495612b3f612b2560408601612480565b612b37612b315f610db8565b91610190565b148490612a43565b8290612b53680500000000000000076107f5565b9091612f48565b604085979297019096906122ed565b94612b7660408401612480565b612b88612b825f610db8565b91610190565b115f14612d5557612ba8906002612ba160408601612480565b9101610f4d565b612bd4612bc5612bc0680400000000000000056107f5565b61103d565b68040000000000000005610f4d565b5b612bf2612be260016107f5565b612bec6001610ed2565b906114df565b90612bfe826001610f4d565b612c18612c1368010000000000000003611abe565b61050f565b6323b872dd612c2630611acb565b3392612c335f8801612480565b92813b15612d50575f612c5991612c648296612c4d610182565b98899788968795611adb565b855260048501611af0565b03925af18015612d4b57612d1f575b50612c99612c89680100000000000000026107f5565b612c936001610ed2565b90610eee565b612cac8168010000000000000002610f4d565b90612cb85f8501612480565b91339390612d1a612ccc60408b9801612480565b612d08612d02612cfc7fec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf97610f2e565b97610f2e565b97611928565b97612d11610182565b94859485612a71565b0390a4565b612d3e905f3d8111612d44575b612d368183610e4b565b810190611ae1565b5f612c73565b503d612d2c565b611b22565b611ad7565b5f600182612d688380612d6e960161256c565b016125cc565b612bd5565b84612d8060208401612373565b612d9a612d94612d8f5f611787565b6106e4565b916106e4565b14155f14612dc9573390612dc5612daf610182565b9283926348aca7ef60e11b8452600484016129aa565b0390fd5b612deb90612dd5610182565b91829163023df6b160e21b8352600483016123c1565b0390fd5b612df76117b4565b612e10612e0a612e05613087565b6106e4565b916106e4565b03612e1757565b612e39612e22613087565b5f91829163118cdaa760e01b8352600483016106fd565b0390fd5b612e465f610821565b612e50825f611937565b90612e84612e7e7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093611928565b91611928565b91612e8d610182565b80612e97816101d4565b0390a3565b90565b612eb3612eae612eb892610190565b6104e4565b612e9c565b90565b90612ec69103612e9c565b90565b612edd612ed8612ee292612e9c565b6104e4565b610190565b90565b612ef9612ef4612efe92610db5565b6104e4565b612e9c565b90565b6001612f0d9103610190565b90565b612f24612f1f612f29926103d9565b6104e4565b610190565b90565b90612f379101610190565b90565b90612f459103610190565b90565b92612f51611513565b92612f5a611513565b50612f63611513565b50612f87612f82612f7385612e9f565b612f7c84612e9f565b90612ebb565b612ec9565b94612f9186612e9f565b612fab612fa5612fa05f612ee5565b612e9c565b91612e9c565b12613076575b9483955b86612fc8612fc284610190565b91610190565b115f1461306557612ff5612fef612fe9680400000000000000078a906102ef565b50610fe0565b97612f01565b966130015f8201611088565b61301361300d84610190565b9161119c565b1161303e5750505061303561303b9293949561302e5f610db8565b955b612f3a565b90612f3a565b90565b95613058613053602061305e94959901611008565b612f10565b90612f2c565b9490612fb5565b5093946130359061303b9394613030565b94506130815f610db8565b94612fb1565b61308f6117b0565b50339056fea26469706673582212202caed7a32b032525ab29f62e494c43a647a72cdb6ef97a34a4d3b0cb4ad9bf5964736f6c634300081b0033",
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
