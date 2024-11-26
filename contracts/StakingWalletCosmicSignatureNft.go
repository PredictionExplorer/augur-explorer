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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"EthDepositReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicSignatureConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStateResets\",\"type\":\"uint256\"}],\"name\":\"StateReset\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetState_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"tryPerformMaintenance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicSignatureConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignatureNft\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"CallDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"InvalidOperationInCurrentState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftAlreadyUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftNotUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"NftOneTimeStaking\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requesterAddress\",\"type\":\"address\"}],\"name\":\"NftStakeActionAccessDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftStakeActionInvalidId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NoStakedNfts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit\",\"type\":\"uint256\"}],\"name\":\"NumEthDepositsToEvaluateMaxLimitIsOutOfAllowedRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"EthDepositReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferredToCharity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicSignatureConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStateResets\",\"type\":\"uint256\"}],\"name\":\"StateReset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_ETH_DEPOSITS_TO_EVALUATE_HARD_MAX_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"actionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ethDeposits\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"depositId\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"rewardAmountPerStakedNft\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numEthDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStateResets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numUnpaidStakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxUnpaidEthDepositIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetState_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"tryPerformMaintenance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"numEthDepositsToEvaluateMaxLimit_\",\"type\":\"uint256\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052346100575761001a61001461015a565b90610339565b61002261005c565b612f0d610516823960805181818161048f01528181611b840152612adf015260a0518181816109520152610d740152612f0d90f35b610062565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061008e90610066565b810190811060018060401b038211176100a657604052565b610070565b906100be6100b761005c565b9283610084565b565b5f80fd5b60018060a01b031690565b6100d8906100c4565b90565b6100e4906100cf565b90565b6100f0816100db565b036100f757565b5f80fd5b90505190610108826100e7565b565b610113816100cf565b0361011a57565b5f80fd5b9050519061012b8261010a565b565b91906040838203126101555780610149610152925f86016100fb565b9360200161011e565b90565b6100c0565b6101786134238038038061016d816100ab565b92833981019061012d565b9091565b90565b61019361018e610198926100c4565b61017c565b6100c4565b90565b6101a49061017f565b90565b6101b09061019b565b90565b90565b6101ca6101c56101cf926101b3565b61017c565b6100c4565b90565b6101db906101b6565b90565b60209181520190565b60207f66745f2e00000000000000000000000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f7220746865206e5f8201520152565b61024160246040926101de565b61024a816101e7565b0190565b6102639060208101905f818303910152610234565b90565b1561026d57565b61027561005c565b63eac0d38960e01b81528061028c6004820161024e565b0390fd5b60207f616d655f2e000000000000000000000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f722074686520675f8201520152565b6102ea60256040926101de565b6102f381610290565b0190565b61030c9060208101905f8183039101526102dd565b90565b1561031657565b61031e61005c565b63eac0d38960e01b815280610335600482016102f7565b0390fd5b6103423361039c565b61036f61034e826101a7565b61036861036261035d5f6101d2565b6100cf565b916100cf565b1415610266565b6103948261038d6103876103825f6101d2565b6100cf565b916100cf565b141561030f565b60805260a052565b6103a5906103a7565b565b6103b0906103b2565b565b6103bb906103df565b565b6103c6906100cf565b9052565b91906103dd905f602085019401906103bd565b565b806103fa6103f46103ef5f6101d2565b6100cf565b916100cf565b1461040a57610408906104b6565b565b61042d6104165f6101d2565b5f918291631e4fbdf760e01b8352600483016103ca565b0390fd5b5f1c90565b60018060a01b031690565b61044d61045291610431565b610436565b90565b61045f9054610441565b90565b5f1b90565b9061047860018060a01b0391610462565b9181191691161790565b61048b9061019b565b90565b90565b906104a66104a16104ad92610482565b61048e565b8254610467565b9055565b5f0190565b6104bf5f610455565b6104c9825f610491565b906104fd6104f77f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610482565b91610482565b9161050661005c565b80610510816104b1565b0390a356fe60806040526004361015610013575b610c77565b61001d5f3561017c565b806315b4e68f146101775780632e53a259146101725780632ffb33041461016d578063400c89561461016857806347ccca021461016357806357951c741461015e578063715018a6146101595780637a025499146101545780637c843a461461014f5780638da5cb5b1461014a5780639e2c8a5b14610145578063a2b136fb14610140578063a694fc3a1461013b578063a6abc99f14610136578063c2ac898f14610131578063c3fe3e281461012c578063ca7c1f9214610127578063d8ee557314610122578063e0c10d6a1461011d578063f2fde38b14610118578063fdbd98b0146101135763fe939afc0361000e57610c43565b610bd8565b610b8e565b610b3b565b610a78565b6109a9565b610974565b61091c565b6108e7565b61089d565b610865565b610729565b6106c7565b610664565b610618565b6105a1565b61056c565b610515565b61044a565b610414565b610296565b6101d9565b60e01c90565b60405190565b5f80fd5b5f80fd5b90565b61019c81610190565b036101a357565b5f80fd5b905035906101b482610193565b565b906020828203126101cf576101cc915f016101a7565b90565b610188565b5f0190565b6101ec6101e73660046101b6565b6113ee565b6101f4610182565b806101fe816101d4565b0390f35b5f80fd5b5f80fd5b5f80fd5b5f80fd5b909182601f8301121561024c5781359167ffffffffffffffff831161024757602001926020830284011161024257565b61020e565b61020a565b610206565b91604083830312610291575f83013567ffffffffffffffff811161028c5761027e83610289928601610212565b9390946020016101a7565b90565b61018c565b610188565b346102c5576102af6102a9366004610251565b91611511565b6102b7610182565b806102c1816101d4565b0390f35b610202565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b6102f8816102de565b8210156103125761030a6001916102ec565b910201905f90565b6102ca565b5f1c90565b67ffffffffffffffff1690565b61033561033a91610317565b61031c565b90565b6103479054610329565b90565b60401c90565b60018060c01b031690565b61036761036c9161034a565b610350565b90565b610379905461035b565b90565b6804000000000000000690610390826102de565b8110156103bb576103a0916102ef565b50906103b85f6103b181850161033d565b930161036f565b90565b5f80fd5b67ffffffffffffffff1690565b6103d5906103bf565b9052565b60018060c01b031690565b6103ed906103d9565b9052565b91602061041292949361040b60408201965f8301906103cc565b01906103e4565b565b346104455761042c6104273660046101b6565b61037c565b90610441610438610182565b928392836103f1565b0390f35b610202565b346104795761046361045d366004610251565b916115de565b61046b610182565b80610475816101d4565b0390f35b610202565b5f91031261048857565b610188565b7f000000000000000000000000000000000000000000000000000000000000000090565b60018060a01b031690565b90565b6104d36104ce6104d8926104b1565b6104bc565b6104b1565b90565b6104e4906104bf565b90565b6104f0906104db565b90565b6104fc906104e7565b9052565b9190610513905f602085019401906104f3565b565b346105455761052536600461047e565b61054161053061048d565b610538610182565b91829182610500565b0390f35b610202565b61055390610190565b9052565b919061056a905f6020850194019061054a565b565b3461059c576105986105876105823660046101b6565b6116e4565b61058f610182565b91829182610557565b0390f35b610202565b346105cf576105b136600461047e565b6105b9611751565b6105c1610182565b806105cb816101d4565b0390f35b610202565b1c90565b90565b6105eb9060086105f093026105d4565b6105d8565b90565b906105fe91546105db565b90565b610615680400000000000000055f906105f3565b90565b346106485761062836600461047e565b610644610633610601565b61063b610182565b91829182610557565b0390f35b610202565b610661680500000000000000065f906105f3565b90565b346106945761067436600461047e565b61069061067f61064d565b610687610182565b91829182610557565b0390f35b610202565b6106a2906104b1565b90565b6106ae90610699565b9052565b91906106c5905f602085019401906106a5565b565b346106f7576106d736600461047e565b6106f36106e261175f565b6106ea610182565b918291826106b2565b0390f35b610202565b91906040838203126107245780610718610721925f86016101a7565b936020016101a7565b90565b610188565b346107585761074261073c3660046106fc565b90611774565b61074a610182565b80610754816101d4565b0390f35b610202565b506801000000000000000090565b90565b6107778161075d565b8210156107915761078960039161076b565b910201905f90565b6102ca565b6107a26107a791610317565b6105d8565b90565b6107b49054610796565b90565b60018060a01b031690565b6107ce6107d391610317565b6107b7565b90565b6107e090546107c2565b90565b68010000000000000004906107f78261075d565b81101561082f576108079161076e565b506108135f82016107aa565b9161082c6002610825600185016107d6565b93016107aa565b90565b5f80fd5b60409061085c610863949695939661085260608401985f85019061054a565b60208301906106a5565b019061054a565b565b346108985761089461088061087b3660046101b6565b6107e3565b61088b939193610182565b93849384610833565b0390f35b610202565b346108cb576108b56108b03660046101b6565b611a1f565b6108bd610182565b806108c7816101d4565b0390f35b610202565b6108e4680100000000000000035f906105f3565b90565b34610917576108f736600461047e565b6109136109026108d0565b61090a610182565b91829182610557565b0390f35b610202565b3461094b5761093561092f3660046106fc565b90611c33565b61093d610182565b80610947816101d4565b0390f35b610202565b7f000000000000000000000000000000000000000000000000000000000000000090565b346109a45761098436600461047e565b6109a061098f610950565b610997610182565b918291826106b2565b0390f35b610202565b346109d9576109b936600461047e565b6109d56109c4611c94565b6109cc610182565b91829182610557565b0390f35b610202565b151590565b6109ec816109de565b036109f357565b5f80fd5b90503590610a04826109e3565b565b610a0f81610699565b03610a1657565b5f80fd5b90503590610a2782610a06565b565b9190604083820312610a515780610a45610a4e925f86016109f7565b93602001610a1a565b90565b610188565b610a5f906109de565b9052565b9190610a76905f60208501940190610a56565b565b34610aa957610aa5610a94610a8e366004610a29565b90612087565b610a9c610182565b91829182610a63565b0390f35b610202565b90565b610ac5610ac0610aca92610aae565b6104bc565b610190565b90565b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b610b01610b0791610190565b91610190565b908115610b12570490565b610acd565b610b2d5f19610b27610100610ab1565b90610af5565b90565b610b38610b17565b90565b34610b6b57610b4b36600461047e565b610b67610b56610b30565b610b5e610182565b91829182610557565b0390f35b610202565b90602082820312610b8957610b86915f01610a1a565b90565b610188565b34610bbc57610ba6610ba1366004610b70565b612101565b610bae610182565b80610bb8816101d4565b0390f35b610202565b610bd5680100000000000000025f906105f3565b90565b34610c0857610be836600461047e565b610c04610bf3610bc1565b610bfb610182565b91829182610557565b0390f35b610202565b90602082820312610c3e575f82013567ffffffffffffffff8111610c3957610c359201610212565b9091565b61018c565b610188565b34610c7257610c5c610c56366004610c0d565b9061210c565b610c64610182565b80610c6e816101d4565b0390f35b610202565b5f80fd5b60209181520190565b60407f6d6574686f642e00000000000000000000000000000000000000000000000000917f4f6e6c792074686520436f736d69635369676e617475726547616d6520636f6e5f8201527f7472616374206973207065726d697474656420746f2063616c6c20746869732060208201520152565b610d046047606092610c7b565b610d0d81610c84565b0190565b9190610d34906020610d2c604086018681035f880152610cf7565b9401906106a5565b565b15610d3e5750565b610d6090610d4a610182565b91829163336e7de360e11b835260048301610d11565b0390fd5b610dab90610da633610d9e610d987f0000000000000000000000000000000000000000000000000000000000000000610699565b91610699565b143390610d36565b6111eb565b565b90565b610dc4610dbf610dc992610dad565b6104bc565b610190565b90565b5f7f546865726520617265206e6f207374616b6564204e4654732e00000000000000910152565b610e006019602092610c7b565b610e0981610dcc565b0190565b610e229060208101905f818303910152610df3565b90565b90565b610e3c610e37610e4192610e25565b6104bc565b610190565b90565b610e53610e5991939293610190565b92610190565b8201809211610e6457565b610ae1565b5f1b90565b90610e7a5f1991610e69565b9181191691161790565b610e98610e93610e9d92610190565b6104bc565b610190565b90565b90565b90610eb8610eb3610ebf92610e84565b610ea0565b8254610e6e565b9055565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610eeb90610ec3565b810190811067ffffffffffffffff821117610f0557604052565b610ecd565b90610f1d610f16610182565b9283610ee1565b565b610f296040610f0a565b90565b5f90565b5f90565b610f3c610f1f565b9060208083610f49610f2c565b815201610f54610f30565b81525050565b610f62610f34565b90565b90610f6f906103bf565b9052565b90610f7d906103d9565b9052565b90610fb7610fae5f610f91610f1f565b94610fa8610fa083830161033d565b838801610f65565b0161036f565b60208401610f73565b565b610fc290610f81565b90565b610fd9610fd4610fde92610190565b6104bc565b6103d9565b90565b610feb90516103d9565b90565b610ffa611000916103d9565b916103d9565b019060018060c01b03821161101157565b610ae1565b61101f90610190565b5f19811461102d5760010190565b610ae1565b61104661104161104b92610190565b6104bc565b6103bf565b90565b634e487b7160e01b5f525f60045260245ffd5b61106b90516103bf565b90565b9061108167ffffffffffffffff91610e69565b9181191691161790565b61109f61109a6110a4926103bf565b6104bc565b6103bf565b90565b90565b906110bf6110ba6110c69261108b565b6110a7565b825461106e565b9055565b60401b90565b906110e467ffffffffffffffff19916110ca565b9181191691161790565b6111026110fd611107926103d9565b6104bc565b6103d9565b90565b90565b9061112261111d611129926110ee565b61110a565b82546110d0565b9055565b9061115760205f61115d9461114f828201611149848801611061565b906110aa565b019201610fe1565b9061110d565b565b91906111705761116e9161112d565b565b61104e565b61118961118461118e926103bf565b6104bc565b610190565b90565b61119a90611175565b9052565b909594926111e9946111d86111e2926111ce6080966111c460a088019c5f89019061054a565b602087019061054a565b6040850190611191565b606083019061054a565b019061054a565b565b6111f560016107aa565b90816112096112035f610db0565b91610190565b146113cb57611233611223680100000000000000026107aa565b61122d6001610e28565b90610e44565b906112478268010000000000000002610ea3565b611259680400000000000000056107aa565b90611262610f5a565b91611275680400000000000000046107aa565b6112876112815f610db0565b91610190565b14155f1461136c576113295f6112ba611367936112b56112a684610db0565b68040000000000000004610ea3565b611016565b936112ce8568040000000000000005610ea3565b6112e26112da88611032565b838801610f65565b6113016112f86112f3348b90610af5565b610fc5565b60208801610f73565b5b6113208661131a6804000000000000000688906102ef565b9061115f565b95939401611061565b94346113557fb71b1087ee7f659cf742c29d8095c562e6e832337190e10dbe81db89955c091396610e84565b9661135e610182565b9586958661119e565b0390a2565b9091506113676113295f61139361138d6804000000000000000686906102ef565b50610fb9565b946113c66113aa6113a5348b90610af5565b610fc5565b6113c060208901916113bb83610fe1565b610fee565b90610f73565b611302565b6113d3610182565b630f13059960e41b8152806113ea60048201610e0d565b0390fd5b6113f790610d64565b565b60207f5f206973206f7574206f662074686520616c6c6f7765642072616e67652e0000917f6e756d4574684465706f73697473546f4576616c756174654d61784c696d69745f8201520152565b611453603e604092610c7b565b61145c816113f9565b0190565b919061148390602061147b604086018681035f880152611446565b94019061054a565b565b1561148d5750565b6114af90611499610182565b91829163285cf1f560e01b835260048301611460565b0390fd5b5090565b6114c66114cc91939293610190565b92610190565b82039182116114d757565b610ae1565b60016114e89101610190565b90565b5f90565b91908110156114ff576020020190565b6102ca565b3561150e81610193565b90565b929161154e9061153d8161153461152e611529610b17565b610190565b91610190565b11158290611485565b6115488584906114b3565b906114b7565b906115585f610db0565b6115615f610db0565b925b8361158061157a6115758987906114b3565b610190565b91610190565b10156115cd576115c06115ba6115986115c693611016565b6115a06114eb565b506115b56115b08a888a916114ef565b611504565b612453565b93610e44565b936114dc565b9290611563565b509150506115dc919250612744565b565b929161161b9061160a816116016115fb6115f6610b17565b610190565b91610190565b11158290611485565b6116158584906114b3565b906114b7565b906116255f610db0565b61162e5f610db0565b925b8361164d6116476116428987906114b3565b610190565b91610190565b101561169a5761168d61168761166561169393611016565b61166d6114eb565b5061168261167d8a888a916114ef565b611504565b6128e3565b93610e44565b936114dc565b9290611630565b509150506116a9919250612744565b565b506801000000000000000090565b90565b6116c5816116ab565b8210156116df576116d76001916116b9565b910201905f90565b6102ca565b6116fb611701916116f36114eb565b5060026116bc565b906105f3565b90565b61170c612c32565b61171461173e565b565b61172a61172561172f92610dad565b6104bc565b6104b1565b90565b61173b90611716565b90565b61174f61174a5f611732565b612c80565b565b611759611704565b565b5f90565b61176761175b565b506117715f6107d6565b90565b6117ac916117a6918161178f6117895f610db0565b91610190565b11806117ae575b6117a1908390611485565b6128e3565b50612744565b565b506117a1826117cc6117c66117c1610b17565b610190565b91610190565b11159050611796565b60407f206f6e6c79206f6e63652e000000000000000000000000000000000000000000917f54686973204e46542068617320616c7265616479206265656e207374616b65645f8201527f2e20416e204e465420697320616c6c6f77656420746f206265207374616b656460208201520152565b611855604b606092610c7b565b61185e816117d5565b0190565b919061188590602061187d604086018681035f880152611848565b94019061054a565b565b1561188f5750565b6118b19061189b610182565b918291633b471d1f60e21b835260048301611862565b0390fd5b1b90565b919060086118d49102916118ce5f19846118b5565b926118b5565b9181191691161790565b91906118f46118ef6118fc93610e84565b610ea0565b9083546118b9565b9055565b90565b9061191460018060a01b0391610e69565b9181191691161790565b611927906104db565b90565b90565b9061194261193d6119499261191e565b61192a565b8254611903565b9055565b634e487b7160e01b5f52602160045260245ffd5b6003111561196b57565b61194d565b9061197a82611961565b565b61198590611970565b90565b6119919061197c565b9052565b9160206119b69294936119af60408201965f830190611988565b019061054a565b565b6119c1906104db565b90565b5f80fd5b60e01b90565b5f9103126119d857565b610188565b604090611a06611a0d94969593966119fc60608401985f8501906106a5565b60208301906106a5565b019061054a565b565b611a17610182565b3d5f823e3d90fd5b611a51611a37611a31600284906116bc565b906105f3565b611a49611a435f610db0565b91610190565b148290611887565b611a70611a5e6001610e28565b611a6a600284906116bc565b906118de565b611a95611a85680100000000000000026107aa565b611a8f6001610e28565b90610e44565b611aa88168010000000000000002610ea3565b611ade611ac8611ac268010000000000000004849061076e565b50611900565b611ad4845f8301610ea3565b600133910161192d565b611afb611aeb60016107aa565b611af56001610e28565b90610e44565b90611b07826001610ea3565b611b23611b146001610e28565b68040000000000000004610ea3565b9060019183903392611b67611b61611b5b7fcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f0082994610e84565b94610e84565b9461191e565b94611b7c611b73610182565b92839283611995565b0390a4611ba87f00000000000000000000000000000000000000000000000000000000000000006104e7565b6323b872dd903390611bb9306119b8565b9392813b15611c2e575f611be091611beb8296611bd4610182565b988997889687956119c8565b8552600485016119dd565b03925af18015611c2957611bfd575b50565b611c1c905f3d8111611c22575b611c148183610ee1565b8101906119ce565b5f611bfa565b503d611c0a565b611a0f565b6119c4565b611c6b91611c659181611c4e611c485f610db0565b91610190565b1180611c6d575b611c60908390611485565b612453565b50612744565b565b50611c6082611c8b611c85611c80610b17565b610190565b91610190565b11159050611c55565b611c9c6114eb565b50611ca760016107aa565b90565b5f90565b90611cc19291611cbc612c32565b611eb8565b90565b5f7f546865726520617265207374696c6c207374616b6564204e4654732e00000000910152565b611cf8601c602092610c7b565b611d0181611cc4565b0190565b611d1a9060208101905f818303910152611ceb565b90565b15611d2457565b611d2c610182565b63a29f5c4d60e01b815280611d4360048201611d05565b0390fd5b5f7f546865726520617265207374696c6c20756e7061696420726577617264732e00910152565b611d7b601f602092610c7b565b611d8481611d47565b0190565b611d9d9060208101905f818303910152611d6e565b90565b15611da757565b611daf610182565b63a29f5c4d60e01b815280611dc660048201611d88565b0390fd5b905090565b611dda5f8092611dca565b0190565b611de790611dcf565b90565b67ffffffffffffffff8111611e0857611e04602091610ec3565b0190565b610ecd565b90611e1f611e1a83611dea565b610f0a565b918252565b606090565b3d5f14611e4457611e393d611e0d565b903d5f602084013e5b565b611e4c611e24565b90611e42565b5f7f5472616e7366657220746f2063686172697479206661696c65642e0000000000910152565b611e86601b602092610c7b565b611e8f81611e52565b0190565b9190611eb6906020611eae604086018681035f880152611e79565b94019061054a565b565b50611ede611ec660016107aa565b611ed8611ed25f610db0565b91610190565b14611d1d565b611f0b611ef3680100000000000000036107aa565b611f05611eff5f610db0565b91610190565b14611da0565b612001575b60019080611f2e611f28611f235f611732565b610699565b91610699565b03611f38575b5090565b611f41306119b8565b31905f808284611f4f610182565b9081611f5a81611dde565b03925af1611f66611e29565b505f14611fb657611fac611f9a7f1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d9261191e565b92611fa3610182565b91829182610557565b0390a25b5f611f34565b909150611ff8611fe67f154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a9261191e565b92611fef610182565b91829182611e93565b0390a25f611fb0565b61201c61200d5f610db0565b68040000000000000005610ea3565b612036612031680500000000000000066107aa565b611016565b6120498168050000000000000006610ea3565b61207f7f8a77a30ae54ca88ceda5f17cd3679c1b868abd2c6d36b67bcbe5e9beac7ec64891612076610182565b91829182610557565b0390a1611f10565b9061209991612094611caa565b611cae565b90565b6120ad906120a8612c32565b6120af565b565b806120ca6120c46120bf5f611732565b610699565b91610699565b146120da576120d890612c80565b565b6120fd6120e65f611732565b5f918291631e4fbdf760e01b8352600483016106b2565b0390fd5b61210a9061209c565b565b9190916121185f610db0565b5b8061213661213061212b8588906114b3565b610190565b91610190565b1015612166576121619061215c612157612152858885916114ef565b611504565b611a1f565b6114dc565b612119565b50509050565b9061217690610190565b9052565b9061218490610699565b9052565b6121926060610f0a565b90565b906121e46121db60026121a6612188565b946121bd6121b55f83016107aa565b5f880161216c565b6121d56121cc600183016107d6565b6020880161217a565b016107aa565b6040840161216c565b565b6121ef90612195565b90565b6121fc9051610699565b90565b5f7f496e76616c6964204e4654207374616b6520616374696f6e2049442e00000000910152565b612233601c602092610c7b565b61223c816121ff565b0190565b919061226390602061225b604086018681035f880152612226565b94019061054a565b565b60207f656365697665207374616b696e67207265776172642e00000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20725f8201520152565b6122bf6036604092610c7b565b6122c881612265565b0190565b9160406122fd9294936122f66122eb606083018381035f8501526122b2565b96602083019061054a565b01906106a5565b565b6123099051610190565b90565b5f7f4e465420686173206e6f74206265656e20756e7374616b65642e000000000000910152565b612340601a602092610c7b565b6123498161230c565b0190565b9190612370906020612368604086018681035f880152612333565b94019061054a565b565b1561237a5750565b61239c90612386610182565b9182916372c1673f60e01b83526004830161234d565b0390fd5b6123b2916123ac6114eb565b916118de565b565b919060086123d49102916123ce60018060a01b03846118b5565b926118b5565b9181191691161790565b91906123f46123ef6123fc9361191e565b61192a565b9083546123b4565b9055565b6124129161240c61175b565b916123de565b565b61241d90610190565b5f811461242b576001900390565b610ae1565b91602061245192949361244a60408201965f83019061054a565b019061054a565b565b9161245c6114eb565b506124656114eb565b5061248361247d68010000000000000004859061076e565b50611900565b9161248d836121e6565b90336124ac6124a66124a1602086016121f2565b610699565b91610699565b0361260057936124f861250794956124e36124c9604086016122ff565b6124db6124d55f610db0565b91610190565b118490612372565b82906124f1604086016122ff565b9091612d8b565b6040859792970190969061216c565b94612520612517604085016122ff565b60028301610ea3565b61252c604084016122ff565b61253e6125385f610db0565b91610190565b146125b5575b506125505f83016122ff565b339161255f60408895016122ff565b61259b61259561258f7ff9bac74bc321a00ef2afeb1f44684e20e22f567699a5840df47967ea88c5b44994610e84565b94610e84565b9461191e565b946125b06125a7610182565b92839283612430565b0390a4565b5f6001826125c883806125ce96016123a0565b01612400565b6125fa6125eb6125e6680100000000000000036107aa565b612414565b68010000000000000003610ea3565b5f612544565b8461260d602084016121f2565b61262761262161261c5f611732565b610699565b91610699565b14155f1461265657339061265261263c610182565b9283926348aca7ef60e11b8452600484016122cc565b0390fd5b61267890612662610182565b91829163023df6b160e21b835260048301612240565b0390fd5b60207f642e000000000000000000000000000000000000000000000000000000000000917f4e4654207374616b696e6720726577617264207061796d656e74206661696c655f8201520152565b6126d66022604092610c7b565b6126df8161267c565b0190565b91604061271492949361270d612702606083018381035f8501526126c9565b9660208301906106a5565b019061054a565b565b1561271f575050565b61274061272a610182565b928392630aa7db6360e11b8452600484016126e3565b0390fd5b612775905f803383612754610182565b908161275f81611dde565b03925af161276b611e29565b5090339091612716565b565b60207f6e7374616b652069742e00000000000000000000000000000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20755f8201520152565b6127d1602a604092610c7b565b6127da81612777565b0190565b91604061280f9294936128086127fd606083018381035f8501526127c4565b96602083019061054a565b01906106a5565b565b5f7f4e46542068617320616c7265616479206265656e20756e7374616b65642e0000910152565b612845601e602092610c7b565b61284e81612811565b0190565b919061287590602061286d604086018681035f880152612838565b94019061054a565b565b1561287f5750565b6128a19061288b610182565b91829163c339165960e01b835260048301612852565b0390fd5b6128da6128e1946128d06060949897956128c6608086019a5f87019061054a565b602085019061054a565b604083019061054a565b019061054a565b565b916128ec6114eb565b506128f56114eb565b5061291361290d68010000000000000004859061076e565b50611900565b9161291d836121e6565b903361293c612936612931602086016121f2565b610699565b91610699565b03612bb6579361298e61299d9495612973612959604086016122ff565b61296b6129655f610db0565b91610190565b148490612877565b8290612987680400000000000000056107aa565b9091612d8b565b6040859792970190969061216c565b946129aa604084016122ff565b6129bc6129b65f610db0565b91610190565b115f14612b98576129dc9060026129d5604086016122ff565b9101610ea3565b612a086129f96129f4680100000000000000036107aa565b611016565b68010000000000000003610ea3565b5b612a26612a1660016107aa565b612a206001610e28565b906114b7565b612a31816001610ea3565b612a56612a46680100000000000000026107aa565b612a506001610e28565b90610e44565b612a698168010000000000000002610ea3565b91612a755f85016122ff565b9033928894612ad7612a89604089016122ff565b612ac5612abf612ab97fec478a78e4e3931ff728a54eeb6875304c891fa5fa253337b60d37fdc5e1feaf97610e84565b97610e84565b9761191e565b97612ace610182565b948594856128a5565b0390a4612b037f00000000000000000000000000000000000000000000000000000000000000006104e7565b6323b872dd90612b12306119b8565b90612b1f5f3395016122ff565b92813b15612b93575f612b4591612b508296612b39610182565b988997889687956119c8565b8552600485016119dd565b03925af18015612b8e57612b62575b50565b612b81905f3d8111612b87575b612b798183610ee1565b8101906119ce565b5f612b5f565b503d612b6f565b611a0f565b6119c4565b5f600182612bab8380612bb196016123a0565b01612400565b612a09565b84612bc3602084016121f2565b612bdd612bd7612bd25f611732565b610699565b91610699565b14155f14612c0c573390612c08612bf2610182565b9283926348aca7ef60e11b8452600484016127de565b0390fd5b612c2e90612c18610182565b91829163023df6b160e21b835260048301612240565b0390fd5b612c3a61175f565b612c53612c4d612c48612eca565b610699565b91610699565b03612c5a57565b612c7c612c65612eca565b5f91829163118cdaa760e01b8352600483016106b2565b0390fd5b612c895f6107d6565b612c93825f61192d565b90612cc7612cc17f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09361191e565b9161191e565b91612cd0610182565b80612cda816101d4565b0390a3565b90565b612cf6612cf1612cfb92610190565b6104bc565b612cdf565b90565b90612d099103612cdf565b90565b612d20612d1b612d2592612cdf565b6104bc565b610190565b90565b612d3c612d37612d4192610dad565b6104bc565b612cdf565b90565b6001612d509103610190565b90565b612d67612d62612d6c926103d9565b6104bc565b610190565b90565b90612d7a9101610190565b90565b90612d889103610190565b90565b92612d946114eb565b92612d9d6114eb565b50612da66114eb565b50612dca612dc5612db685612ce2565b612dbf84612ce2565b90612cfe565b612d0c565b94612dd486612ce2565b612dee612de8612de35f612d28565b612cdf565b91612cdf565b12612eb9575b9483955b86612e0b612e0584610190565b91610190565b115f14612ea857612e38612e32612e2c680400000000000000068a906102ef565b50610fb9565b97612d44565b96612e445f8201611061565b612e56612e5084610190565b91611175565b11612e8157505050612e78612e7e92939495612e715f610db0565b955b612d7d565b90612d7d565b90565b95612e9b612e966020612ea194959901610fe1565b612d53565b90612d6f565b9490612df8565b509394612e7890612e7e9394612e73565b9450612ec45f610db0565b94612df4565b612ed261175b565b50339056fea2646970667358221220318e7c1df0c0f7bef4aa247bbb7bc17732e20ab7cbbff155bdae94555b15920664736f6c634300081b0033",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicSignatureConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
